package collectionsvc

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/collections"
	"github.com/nkramber/decktome/go/internal/users"
)

// fakeRepo is an in-memory Repo. It records what the service stored.
type fakeRepo struct {
	stored  map[string]*mtgv1.Collection
	findErr error
	putErr  error
	getErr  error
	puts    int
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{stored: map[string]*mtgv1.Collection{}}
}

func (f *fakeRepo) Delete(_ context.Context, _ string, id string) error {
	if f.getErr != nil {
		return f.getErr
	}
	if _, ok := f.stored[id]; !ok {
		return status.Error(codes.NotFound, "no such collection")
	}
	delete(f.stored, id)
	return nil
}

// GetHead answers the binder head, which reads no entry (D-392).
func (f *fakeRepo) GetHead(ctx context.Context, uid, id string) (*mtgv1.Collection, error) {
	col, err := f.Get(ctx, uid, id)
	if err != nil {
		return nil, err
	}
	head := proto.Clone(col).(*mtgv1.Collection)
	head.Entries = nil
	if head.GetSummary() == nil {
		head.Summary = collections.Summarize(col.GetEntries(), nil)
	}
	return head, nil
}

// Rename writes the name and keeps everything else.
func (f *fakeRepo) Rename(ctx context.Context, uid, id, name string) (*mtgv1.Collection, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}
	col, ok := f.stored[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "no such collection")
	}
	col.Name = name
	return f.GetHead(ctx, uid, id)
}

// Replace writes new entries into one collection and keeps its id.
func (f *fakeRepo) Replace(ctx context.Context, uid, id string, col *mtgv1.Collection) (*mtgv1.Collection, error) {
	old, ok := f.stored[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "no such collection")
	}
	col.Id, col.Name = id, old.GetName()
	if _, err := f.Put(ctx, uid, col); err != nil {
		return nil, err
	}
	return f.GetHead(ctx, uid, id)
}

// Put mirrors the id rule of the real repo (D-399). A collection with no
// id takes the id its hash derives, unless that document exists and
// holds another hash, which is a collection a Replace moved on.
func (f *fakeRepo) Put(_ context.Context, _ string, col *mtgv1.Collection) (string, error) {
	if f.putErr != nil {
		return "", f.putErr
	}
	f.puts++
	id := col.Id
	if id == "" && col.ContentHash != "" {
		derived := collections.DocID(col.ContentHash)
		if have, ok := f.stored[derived]; !ok || have.GetContentHash() == col.ContentHash {
			id = derived
		}
	}
	if id == "" {
		id = "doc" + string(rune('0'+len(f.stored)))
	}
	f.stored[id] = col
	return id, nil
}

// Get answers a collection the caller owns, as the Repo contract says.
// The real repo unmarshals a fresh object on every call, and a fake that
// shares one hands the next reader a collection the last page
// truncated.
func (f *fakeRepo) Get(_ context.Context, _, id string) (*mtgv1.Collection, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}
	c, ok := f.stored[id]
	if !ok {
		// The real repo passes the Firestore status through.
		return nil, status.Error(codes.NotFound, "not found")
	}
	return proto.Clone(c).(*mtgv1.Collection), nil
}

func (f *fakeRepo) List(context.Context, string) ([]*mtgv1.Collection, error) {
	var out []*mtgv1.Collection
	for _, c := range f.stored {
		out = append(out, c)
	}
	return out, nil
}

// FindByHash reads the hash each document holds now, as the Firestore
// query does. A document a Replace moved on answers to its new hash
// alone.
func (f *fakeRepo) FindByHash(_ context.Context, _, hash string) (string, error) {
	if f.findErr != nil {
		return "", f.findErr
	}
	if hash == "" {
		return "", nil
	}
	for id, col := range f.stored {
		if col.GetContentHash() == hash {
			return id, nil
		}
	}
	return "", nil
}

type staticIndex struct{ idx *cards.Index }

func (s staticIndex) Current() *cards.Index { return s.idx }

const scryfallID = "aaaaaaaa-0000-0000-0000-000000000001"

func testIndex() *cards.Index {
	c := &mtgv1.Card{
		OracleId:        "oracle-1",
		Name:            "Pawpatch Recruit",
		DefaultPrinting: &mtgv1.Printing{ScryfallId: scryfallID, SetCode: "blb", CollectorNumber: "187"},
	}
	printings := []cards.Printing{
		{ScryfallID: scryfallID, OracleID: c.OracleId, Name: c.Name, SetCode: "blb", CollectorNumber: "187", Layout: "normal"},
		{ScryfallID: "7cdd8679-93a7-4e58-a6f3-b48897697e89", OracleID: "oracle-token", Name: c.Name, SetCode: "tblb", CollectorNumber: "21", Layout: "token"},
	}
	return cards.NewIndex([]*mtgv1.Card{c}, printings, nil, time.Now())
}

func newServer(repo Repo, idx *cards.Index) *Server {
	return New(repo, staticIndex{idx}, func(context.Context) string { return "user1" })
}

// goodCSV holds five rows: two resolve to one merged entry, three fail.
const goodCSV = "Name,Set code,Collector number,Quantity,Scryfall ID,Foil,Condition,Language\n" +
	"Pawpatch Recruit,BLB,187,3," + scryfallID + ",foil,near_mint,en\n" +
	"Pawpatch Recruit,BLB,187,2," + scryfallID + ",foil,near_mint,en\n" +
	"Pawpatch Recruit,TBLB,21,1,7cdd8679-93a7-4e58-a6f3-b48897697e89,normal,near_mint,en\n" +
	"Nope,ZZZ,1,1,,normal,near_mint,en\n" +
	"Pawpatch Recruit,BLB,187,1," + scryfallID + ",shiny,near_mint,en\n"

func importReq(name string, src mtgv1.ImportSource, content string) *connect.Request[mtgv1.ImportCollectionRequest] {
	return connect.NewRequest(&mtgv1.ImportCollectionRequest{Name: name, Source: src, Content: []byte(content)})
}

func TestImportCollectionReport(t *testing.T) {
	repo := newFakeRepo()
	s := newServer(repo, testIndex())
	resp, err := s.ImportCollection(context.Background(), importReq("Binder", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if err != nil {
		t.Fatal(err)
	}
	col, rep := resp.Msg.Collection, resp.Msg.Report
	if col.Id == "" || col.CardCount != 5 || len(col.Entries) != 1 {
		t.Errorf("collection = %v", col)
	}
	// ResolvedCount counts rows, so resolved plus unresolved is the
	// input row count of five.
	if rep.ResolvedCount != 2 || len(rep.Unresolved) != 3 {
		t.Errorf("report = %v", rep)
	}
	want := map[string]int32{
		"UNRESOLVED_REASON_NOT_PLAYABLE":  1,
		"UNRESOLVED_REASON_UNKNOWN_CARD":  1,
		"UNRESOLVED_REASON_UNKNOWN_VALUE": 1,
	}
	for k, v := range want {
		if rep.UnresolvedByReason[k] != v {
			t.Errorf("by reason %s = %d, want %d (%v)", k, rep.UnresolvedByReason[k], v, rep.UnresolvedByReason)
		}
	}
}

func TestImportCollectionRejects(t *testing.T) {
	s := newServer(newFakeRepo(), testIndex())
	tests := []struct {
		name string
		req  *connect.Request[mtgv1.ImportCollectionRequest]
		code connect.Code
	}{
		{"empty body", importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, ""), connect.CodeInvalidArgument},
		{"too large", importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, strings.Repeat("a", maxUpload+1)), connect.CodeInvalidArgument},
		// An unnamed source reads the format out of the file now (D-647),
		// so a file no format claims is the refusal and not the unnamed
		// source.
		{"a file no format claims", importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED, "First,Last\nAnn,Lee\n"), connect.CodeInvalidArgument},
		{"bad header", importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, "Foo,Bar\n1,2\n"), connect.CodeInvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.ImportCollection(context.Background(), tt.req)
			if connect.CodeOf(err) != tt.code {
				t.Errorf("code = %v (%v), want %v", connect.CodeOf(err), err, tt.code)
			}
		})
	}
}

func TestImportCollectionNoIndex(t *testing.T) {
	s := newServer(newFakeRepo(), nil)
	_, err := s.ImportCollection(context.Background(), importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if connect.CodeOf(err) != connect.CodeUnavailable {
		t.Errorf("err = %v", err)
	}
}

// TestImportCollectionDedup covers D-16: an identical
// re-upload reuses the document and takes the new name.
func TestImportCollectionDedup(t *testing.T) {
	repo := newFakeRepo()
	s := newServer(repo, testIndex())
	first, err := s.ImportCollection(context.Background(), importReq("Old", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if err != nil {
		t.Fatal(err)
	}
	second, err := s.ImportCollection(context.Background(), importReq("New", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if err != nil {
		t.Fatal(err)
	}
	if first.Msg.Collection.Id != second.Msg.Collection.Id {
		t.Errorf("ids differ: %s %s", first.Msg.Collection.Id, second.Msg.Collection.Id)
	}
	if len(repo.stored) != 1 || repo.stored[second.Msg.Collection.Id].Name != "New" {
		t.Errorf("stored = %v", repo.stored)
	}
}

func TestImportCollectionErrorMapping(t *testing.T) {
	repo := newFakeRepo()
	repo.findErr = errors.New("firestore down")
	s := newServer(repo, testIndex())
	_, err := s.ImportCollection(context.Background(), importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if connect.CodeOf(err) != connect.CodeInternal || repo.puts != 0 {
		t.Errorf("find error: %v, puts %d", err, repo.puts)
	}

	repo = newFakeRepo()
	repo.putErr = collections.ErrTooLarge
	s = newServer(repo, testIndex())
	_, err = s.ImportCollection(context.Background(), importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if connect.CodeOf(err) != connect.CodeResourceExhausted {
		t.Errorf("too large: %v", err)
	}

	repo = newFakeRepo()
	repo.putErr = errors.New("write failed")
	s = newServer(repo, testIndex())
	_, err = s.ImportCollection(context.Background(), importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if connect.CodeOf(err) != connect.CodeInternal {
		t.Errorf("put error: %v", err)
	}
}

// TestGetCollectionCodes: only a Firestore NotFound is NotFound, and an
// id that names another path is refused.
func TestGetCollectionCodes(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		getErr error
		want   connect.Code
	}{
		{name: "empty id", id: "", want: connect.CodeInvalidArgument},
		{name: "blank id", id: "  ", want: connect.CodeInvalidArgument},
		{name: "path id", id: "../other/doc", want: connect.CodeInvalidArgument},
		{name: "dot id", id: ".", want: connect.CodeInvalidArgument},
		{name: "long id", id: strings.Repeat("x", 1501), want: connect.CodeInvalidArgument},
		{name: "missing document", id: "missing", want: connect.CodeNotFound},
		{name: "store failure", id: "col-1", getErr: errors.New("firestore down"), want: connect.CodeInternal},
		{name: "permission denied is not not-found", id: "col-1", getErr: status.Error(codes.PermissionDenied, "no"), want: connect.CodeInternal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newFakeRepo()
			repo.getErr = tt.getErr
			s := newServer(repo, testIndex())
			_, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{CollectionId: tt.id}))
			if connect.CodeOf(err) != tt.want {
				t.Errorf("code = %v, want %v: %v", connect.CodeOf(err), tt.want, err)
			}
		})
	}
}

// TestImportNothingResolvedAnswersTheReport: an upload where no row
// resolves stores nothing and returns the report, not an error.
func TestImportNothingResolvedAnswersTheReport(t *testing.T) {
	repo := newFakeRepo()
	s := newServer(repo, testIndex())
	resp, err := s.ImportCollection(context.Background(), importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, "Name,Set code\nNope,zzz\n"))
	if err != nil {
		t.Fatalf("err = %v, want a response", err)
	}
	if repo.puts != 0 {
		t.Error("an empty collection was stored")
	}
	col, rep := resp.Msg.GetCollection(), resp.Msg.GetReport()
	if col.GetId() != "" || len(col.GetEntries()) != 0 || col.GetContentHash() == "" {
		t.Errorf("collection = %v", col)
	}
	if rep.GetResolvedCount() != 0 || len(rep.GetUnresolved()) != 1 {
		t.Errorf("report = %v", rep)
	}
}

// TestImportName: the name is trimmed and capped.
func TestImportName(t *testing.T) {
	repo := newFakeRepo()
	s := newServer(repo, testIndex())
	resp, err := s.ImportCollection(context.Background(), importReq("  Binder  ", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if err != nil {
		t.Fatal(err)
	}
	if resp.Msg.GetCollection().GetName() != "Binder" {
		t.Errorf("name = %q", resp.Msg.GetCollection().GetName())
	}
	_, err = s.ImportCollection(context.Background(), importReq(strings.Repeat("n", MaxNameBytes+1), mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("a long name gave %v", err)
	}
}

// TestNoUserIsUnauthenticated covers every RPC.
func TestNoUserIsUnauthenticated(t *testing.T) {
	s := New(newFakeRepo(), staticIndex{testIndex()}, func(context.Context) string { return "" })
	ctx := context.Background()
	if _, err := s.ImportCollection(ctx, importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV)); connect.CodeOf(err) != connect.CodeUnauthenticated {
		t.Errorf("import: %v", err)
	}
	if _, err := s.GetCollection(ctx, connect.NewRequest(&mtgv1.GetCollectionRequest{CollectionId: "c1"})); connect.CodeOf(err) != connect.CodeUnauthenticated {
		t.Errorf("get: %v", err)
	}
	if _, err := s.ListCollections(ctx, connect.NewRequest(&mtgv1.ListCollectionsRequest{})); connect.CodeOf(err) != connect.CodeUnauthenticated {
		t.Errorf("list: %v", err)
	}
}

// DeleteCollection removes one collection for good (D-347).
func TestDeleteCollection(t *testing.T) {
	ctx := context.Background()
	del := func(repo *fakeRepo, id string) error {
		_, err := newServer(repo, nil).DeleteCollection(ctx, connect.NewRequest(&mtgv1.DeleteCollectionRequest{CollectionId: id}))
		return err
	}

	t.Run("it removes the collection", func(t *testing.T) {
		repo := newFakeRepo()
		repo.stored["doc0"] = &mtgv1.Collection{Id: "doc0", Name: "binder.csv"}
		if err := del(repo, "doc0"); err != nil {
			t.Fatal(err)
		}
		if _, ok := repo.stored["doc0"]; ok {
			t.Error("the collection is still stored")
		}
	})

	t.Run("a collection that is gone is NotFound", func(t *testing.T) {
		if err := del(newFakeRepo(), "doc0"); connect.CodeOf(err) != connect.CodeNotFound {
			t.Errorf("code = %v", connect.CodeOf(err))
		}
	})

	t.Run("an empty id is an invalid argument", func(t *testing.T) {
		if err := del(newFakeRepo(), " "); connect.CodeOf(err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", connect.CodeOf(err))
		}
	})
}

// fakeNoter records what the service counted on the user record (D-638).
type fakeNoter struct{ counts []users.Counter }

func (f *fakeNoter) Note(_ context.Context, _, _ string, c users.Counter, _ time.Time) error {
	f.counts = append(f.counts, c)
	return nil
}

// TestOnlyANewCollectionCounts locks the dedupe rule of D-638. A repeat
// upload of the same content answers the collection that holds it, so
// it makes no new collection and raises no counter. The emulator lane
// proves the store, and CI never runs it, so this proves the branch.
func TestOnlyANewCollectionCounts(t *testing.T) {
	repo := newFakeRepo()
	noter := &fakeNoter{}
	s := New(repo, staticIndex{testIndex()}, func(context.Context) string { return "user1" },
		WithUsers(noter))

	if _, err := s.ImportCollection(context.Background(),
		importReq("Binder", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV)); err != nil {
		t.Fatalf("first import: %v", err)
	}
	if len(noter.counts) != 1 || noter.counts[0] != users.CollectionsUpload {
		t.Fatalf("the first upload counted %v, want one collection", noter.counts)
	}

	// The same bytes again. The repo answers the collection that already
	// holds the hash, so no collection is made.
	if _, err := s.ImportCollection(context.Background(),
		importReq("Binder again", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV)); err != nil {
		t.Fatalf("second import: %v", err)
	}
	if len(noter.counts) != 1 {
		t.Errorf("a repeat upload counted %v, and it made no collection", noter.counts)
	}
}

// TestImportDetectsTheFormat holds D-647: the reader drops a file and
// never names the app it came from. The web app sent MANABOX_CSV for
// every upload before this, so an Arena list uploaded there failed
// every row.
func TestImportDetectsTheFormat(t *testing.T) {
	for name, content := range map[string]string{
		"a ManaBox export": goodCSV,
		"an Arena list":    "4 Sol Ring (C21) 263\n",
	} {
		t.Run(name, func(t *testing.T) {
			s := newServer(newFakeRepo(), testIndex())
			resp, err := s.ImportCollection(context.Background(),
				importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED, content))
			if err != nil {
				t.Fatalf("an unnamed source did not detect: %v", err)
			}
			if resp.Msg.GetCollection() == nil {
				t.Fatal("no collection came back")
			}
		})
	}
}

// TestANamedSourceBeatsTheDetector holds the other half. A caller that
// names a format means it, and the reader's own word beats a guess
// (D-591).
func TestANamedSourceBeatsTheDetector(t *testing.T) {
	s := newServer(newFakeRepo(), testIndex())
	// An Arena list named as a ManaBox CSV fails on the header, and it
	// never falls back to the detector.
	_, err := s.ImportCollection(context.Background(),
		importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, "4 Sol Ring (C21) 263\n"))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("err = %v, want the named parser to refuse it", err)
	}
}
