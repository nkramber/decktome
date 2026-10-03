package agentsvc

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/archidekt"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/deckreads"
)

// archidektDeck is an answer of the API in the shape of 2026-10-03: a
// commander in the premier category, a land, and a maybeboard card.
const archidektDeck = `{"name":"Karlov Lifegain","categories":[
 {"name":"Commander","isPremier":true,"includedInDeck":true},
 {"name":"Land","isPremier":false,"includedInDeck":true},
 {"name":"Maybeboard","isPremier":false,"includedInDeck":false}],
"cards":[
 {"quantity":1,"categories":["Commander"],"companion":false,"modifier":"Normal","deletedAt":null,
  "card":{"collectorNumber":"1","edition":{"editioncode":"mkm"},"oracleCard":{"name":"Karlov of the Ghost Council"}}},
 {"quantity":99,"categories":["Land"],"companion":false,"modifier":"Normal","deletedAt":null,
  "card":{"collectorNumber":"294","edition":{"editioncode":"neo"},"oracleCard":{"name":"Plains"}}},
 {"quantity":2,"categories":["Maybeboard"],"companion":false,"modifier":"Normal","deletedAt":null,
  "card":{"collectorNumber":"1","edition":{"editioncode":"m19"},"oracleCard":{"name":"Ajani's Welcome"}}}]}`

// fakeReads is a ReadLog in memory. It keeps the rules of deckreads, and
// fail makes each call fail.
type fakeReads struct {
	mu    sync.Mutex
	reads map[string][]deckreads.Read
	fail  error
}

func (f *fakeReads) Add(_ context.Context, uid string, deckID int64, text string, now time.Time) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.fail != nil {
		return f.fail
	}
	if f.reads == nil {
		f.reads = map[string][]deckreads.Read{}
	}
	f.reads[uid] = deckreads.Keep(f.reads[uid], &deckreads.Read{DeckID: deckID, TextHash: deckreads.Hash(text), ReadAt: now}, now)
	return nil
}

func (f *fakeReads) Has(_ context.Context, uid string, deckID int64, text string, now time.Time) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.fail != nil {
		return false, f.fail
	}
	return deckreads.Matches(f.reads[uid], deckID, text, now), nil
}

// fetchServer wires an agent server to a fake Archidekt that answers
// status with body, and counts its reads.
func fetchServer(t *testing.T, status int, body string) (mtgv1connect.AgentServiceClient, *atomic.Int32) {
	t.Helper()
	return fetchServerWith(t, status, body, &fakeReads{})
}

// fetchServerWith is fetchServer with a read log of the test.
func fetchServerWith(t *testing.T, status int, body string, log ReadLog) (mtgv1connect.AgentServiceClient, *atomic.Int32) {
	t.Helper()
	var reads atomic.Int32
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		reads.Add(1)
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(api.Close)
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	opts := []Option{
		WithDecks(&fakeDecks{}), WithCandidates(fixedIndex{importCardIndex()}, cb), WithDeckStore(&fakeDeckStore{}), WithUsers(&fakeNoter{}),
		WithArchidekt(archidekt.New().WithBaseURL(api.URL)), WithReadLog(log),
	}
	client, _ := testServerOpts(t, newFakeStore(), opts)
	return client, &reads
}

func fetchList(c mtgv1connect.AgentServiceClient, url string) (*mtgv1.FetchDeckListResponse, error) {
	res, err := c.FetchDeckList(context.Background(), connect.NewRequest(&mtgv1.FetchDeckListRequest{Url: url}))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// TestFetchDeckListFeedsTheImport is D-1100 and D-1101: the read answers
// the list as text, the import reads that text, and the deck keeps the
// link that the server made again from the id.
func TestFetchDeckListFeedsTheImport(t *testing.T) {
	client, reads := fetchServer(t, http.StatusOK, archidektDeck)
	got, err := fetchList(client, "  www.archidekt.com/decks/42/karlov_lifegain?tab=stats ")
	if err != nil {
		t.Fatal(err)
	}
	want := "Commander\n1 Karlov of the Ghost Council (mkm) 1\n\nDeck\n99 Plains (neo) 294\n"
	if got.GetText() != want || got.GetName() != "Karlov Lifegain" || got.GetLeftOut() != 2 || got.GetSourceUrl() != "https://archidekt.com/decks/42" {
		t.Fatalf("fetch = %v", got)
	}
	if reads.Load() != 1 {
		t.Errorf("reads = %d", reads.Load())
	}
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: got.GetText(), Name: got.GetName(), SourceUrl: got.GetSourceUrl()})
	deck := res.GetDeck()
	if deck.GetSourceUrl() != "https://archidekt.com/decks/42" {
		t.Errorf("source = %q", deck.GetSourceUrl())
	}
	if ids := deck.GetCommanderOracleIds(); len(ids) != 1 || ids[0] != "o-karlov" {
		t.Errorf("commanders = %v", ids)
	}
}

// TestImportKeepsOnlyTheRemadeLink is D-1101 and D-1107: a pasted list
// keeps no link, a link keeps only the deck page, a link needs a read of
// the same deck and text, and a link to another host fails before the
// list reads.
func TestImportKeepsOnlyTheRemadeLink(t *testing.T) {
	client, _ := fetchServer(t, http.StatusOK, archidektDeck)
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList})
	if res.GetDeck().GetSourceUrl() != "" {
		t.Errorf("a pasted list keeps %q", res.GetDeck().GetSourceUrl())
	}
	_, err := client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: markedList, SourceUrl: "https://archidekt.com/decks/42"}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("a link with no read: err = %v, want InvalidArgument", err)
	}
	got, err := fetchList(client, "https://archidekt.com/decks/42")
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: markedList, SourceUrl: "https://archidekt.com/decks/42"}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("another text than the read: err = %v, want InvalidArgument", err)
	}
	_, err = client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: got.GetText(), SourceUrl: "https://archidekt.com/decks/43"}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("another deck than the read: err = %v, want InvalidArgument", err)
	}
	res = importList(t, client, &mtgv1.ImportDeckRequest{Text: got.GetText(), SourceUrl: "archidekt.com/decks/42/a\"><script>"})
	if res.GetDeck().GetSourceUrl() != "https://archidekt.com/decks/42" {
		t.Errorf("source = %q", res.GetDeck().GetSourceUrl())
	}
	for _, bad := range []string{"https://evil.example/decks/42", "javascript:alert(1)", "https://moxfield.com/decks/abc"} {
		_, err := client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: markedList, SourceUrl: bad}))
		if connect.CodeOf(err) != connect.CodeInvalidArgument {
			t.Errorf("%q: err = %v, want InvalidArgument", bad, err)
		}
	}
}

// TestFetchDeckListGivesTheStepsOfASite is D-1103 and D-1104: Moxfield
// gets its exact steps, another site gets the general steps and reads as
// unknown, text that is no link fails, and the server sends no request.
func TestFetchDeckListGivesTheStepsOfASite(t *testing.T) {
	client, reads := fetchServer(t, http.StatusOK, archidektDeck)
	mox, err := fetchList(client, "https://moxfield.com/decks/VaR9P-HceECgmgm55DC7ow")
	if err != nil {
		t.Fatal(err)
	}
	if mox.GetSite() != "Moxfield" || !mox.GetKnownSite() || mox.GetText() != "" || !strings.Contains(strings.Join(mox.GetExportSteps(), " "), "Copy for Arena") {
		t.Errorf("moxfield = %v", mox)
	}
	other, err := fetchList(client, "https://tappedout.net/mtg-decks/some-deck/")
	if err != nil {
		t.Fatal(err)
	}
	if other.GetSite() != "tappedout.net" || other.GetKnownSite() || len(other.GetExportSteps()) == 0 {
		t.Errorf("other = %v", other)
	}
	if _, err := fetchList(client, "1 Sol Ring"); connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("no link: err = %v", err)
	}
	if reads.Load() != 0 {
		t.Errorf("reads = %d, want none", reads.Load())
	}
}

func TestFetchDeckListAnswers(t *testing.T) {
	for _, tc := range []struct {
		status int
		want   connect.Code
	}{
		{http.StatusNotFound, connect.CodeNotFound},
		{http.StatusForbidden, connect.CodeNotFound},
		{http.StatusBadGateway, connect.CodeUnavailable},
	} {
		client, _ := fetchServer(t, tc.status, `{}`)
		if _, err := fetchList(client, "https://archidekt.com/decks/42"); connect.CodeOf(err) != tc.want {
			t.Errorf("status %d: err = %v, want %v", tc.status, err, tc.want)
		}
	}
}

// TestFetchDeckListLimitsEachUser is D-1100: the read after the cap of
// one minute fails, and Archidekt sees no request for it.
func TestFetchDeckListLimitsEachUser(t *testing.T) {
	client, reads := fetchServer(t, http.StatusOK, archidektDeck)
	for i := range FetchesPerMinute {
		if _, err := fetchList(client, "https://archidekt.com/decks/42"); err != nil {
			t.Fatalf("read %d: %v", i+1, err)
		}
	}
	if _, err := fetchList(client, "https://archidekt.com/decks/42"); connect.CodeOf(err) != connect.CodeResourceExhausted {
		t.Fatalf("err = %v, want ResourceExhausted", err)
	}
	if reads.Load() != FetchesPerMinute {
		t.Errorf("reads = %d, want %d", reads.Load(), FetchesPerMinute)
	}
}

// TestAReadLogFaultFailsClosed is D-1107: when the record of a read
// fails, the read answers Unavailable, and an import with a link does too.
func TestAReadLogFaultFailsClosed(t *testing.T) {
	log := &fakeReads{}
	client, _ := fetchServerWith(t, http.StatusOK, archidektDeck, log)
	got, err := fetchList(client, "https://archidekt.com/decks/42")
	if err != nil {
		t.Fatal(err)
	}
	log.fail = errors.New("firestore down")
	if _, err := fetchList(client, "https://archidekt.com/decks/42"); connect.CodeOf(err) != connect.CodeUnavailable {
		t.Errorf("read: err = %v, want Unavailable", err)
	}
	_, err = client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: got.GetText(), SourceUrl: got.GetSourceUrl()}))
	if connect.CodeOf(err) != connect.CodeUnavailable {
		t.Errorf("import: err = %v, want Unavailable", err)
	}
}

func TestFetchDeckListNeedsAReader(t *testing.T) {
	client, _ := testServerOpts(t, newFakeStore(), nil)
	if _, err := fetchList(client, "https://archidekt.com/decks/42"); connect.CodeOf(err) != connect.CodeUnimplemented {
		t.Fatalf("err = %v, want Unimplemented", err)
	}
}
