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

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
)

// fakeRepo is an in-memory Repo. It records what the service stored.
type fakeRepo struct {
	byHash  map[string]string
	stored  map[string]*mtgv1.Collection
	findErr error
	putErr  error
	getErr  error
	puts    int
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{byHash: map[string]string{}, stored: map[string]*mtgv1.Collection{}}
}

func (f *fakeRepo) Put(_ context.Context, _ string, col *mtgv1.Collection) (string, error) {
	if f.putErr != nil {
		return "", f.putErr
	}
	f.puts++
	id := col.Id
	if id == "" {
		id = "doc" + string(rune('0'+len(f.stored)))
	}
	f.stored[id] = col
	f.byHash[col.ContentHash] = id
	return id, nil
}

func (f *fakeRepo) Get(_ context.Context, _, id string) (*mtgv1.Collection, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}
	c, ok := f.stored[id]
	if !ok {
		// The real repo passes the Firestore status through.
		return nil, status.Error(codes.NotFound, "not found")
	}
	return c, nil
}

func (f *fakeRepo) List(context.Context, string) ([]*mtgv1.Collection, error) {
	var out []*mtgv1.Collection
	for _, c := range f.stored {
		out = append(out, c)
	}
	return out, nil
}

func (f *fakeRepo) FindByHash(_ context.Context, _, hash string) (string, error) {
	if f.findErr != nil {
		return "", f.findErr
	}
	return f.byHash[hash], nil
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
		{"bad source", importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED, goodCSV), connect.CodeInvalidArgument},
		{"bad header", importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, "Foo,Bar\n1,2\n"), connect.CodeInvalidArgument},
		{"nothing resolved", importReq("n", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, "Name,Set code\nNope,zzz\n"), connect.CodeInvalidArgument},
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

// TestImportCollectionDedup covers D-16 and C-6: an identical
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

// TestGetCollectionCodes is L-13: only a Firestore NotFound is NotFound.
func TestGetCollectionCodes(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		getErr error
		want   connect.Code
	}{
		{name: "empty id", id: "", want: connect.CodeInvalidArgument},
		{name: "blank id", id: "  ", want: connect.CodeInvalidArgument},
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
