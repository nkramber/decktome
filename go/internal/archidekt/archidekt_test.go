package archidekt

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/collections"
	"github.com/nkramber/decktome/go/internal/decklist"
)

func TestParseURL(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		id   int64
		want error
	}{
		{"https://archidekt.com/decks/3045150/modern_artifact_burn", 3045150, nil},
		{"https://www.archidekt.com/decks/3045150", 3045150, nil},
		{"  archidekt.com/decks/3045150/  ", 3045150, nil},
		{"http://Archidekt.com/decks/7031486#main", 7031486, nil},
		{"https://archidekt.com/decks/3045150?tab=stats", 3045150, nil},
		{"https://moxfield.com/decks/VaR9P-HceECgmgm55DC7ow", 0, ErrNotDeckURL},
		{"https://archidekt.com.evil.example/decks/1", 0, ErrNotDeckURL},
		{"https://evil.example/archidekt.com/decks/1", 0, ErrNotDeckURL},
		{"https://archidekt.com:8443/decks/1", 0, ErrNotDeckURL},
		{"https://user@archidekt.com/decks/1", 0, ErrNotDeckURL},
		{"ftp://archidekt.com/decks/1", 0, ErrNotDeckURL},
		{"javascript:alert(1)", 0, ErrNotDeckURL},
		{"https://archidekt.com/folders/1", 0, ErrNotDeckURL},
		{"https://archidekt.com/decks/", 0, ErrNotDeckURL},
		{"https://archidekt.com/decks/abc", 0, ErrNotDeckURL},
		{"https://archidekt.com/decks/0", 0, ErrNotDeckURL},
		{"https://archidekt.com/decks/007", 0, ErrNotDeckURL},
		{"https://archidekt.com/decks/-5", 0, ErrNotDeckURL},
		{"", 0, ErrNotDeckURL},
		{"https://archidekt.com/decks/1/" + strings.Repeat("a", maxURLBytes), 0, ErrNotDeckURL},
	} {
		id, err := ParseURL(tc.raw)
		if !errors.Is(err, tc.want) || id != tc.id {
			t.Errorf("ParseURL(%q) = %d, %v; want %d, %v", tc.raw, id, err, tc.id, tc.want)
		}
	}
}

func TestDeckURL(t *testing.T) {
	if got := DeckURL(3045150); got != "https://archidekt.com/decks/3045150" {
		t.Fatalf("DeckURL = %q", got)
	}
}

// serve answers the fixture at the API path of the id, and records the
// agent of the request.
func serve(t *testing.T, id, fixture string, agent *string) *Client {
	t.Helper()
	body, err := os.ReadFile("testdata/" + fixture)
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*agent = r.Header.Get("User-Agent")
		if r.URL.Path != "/api/decks/"+id+"/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	t.Cleanup(srv.Close)
	return New().WithBaseURL(srv.URL)
}

// counts parses the text as the import does, and sums the copies of each
// section.
func counts(t *testing.T, text string) map[decklist.Section]int {
	t.Helper()
	list, err := decklist.Parse(strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Bad) > 0 {
		t.Fatalf("lines that read as no card: %v", list.Bad)
	}
	out := map[decklist.Section]int{}
	for _, l := range list.Lines {
		out[l.Section] += l.Quantity
	}
	return out
}

// TestFetchSixtyWithSideboard reads a real Modern deck: 60 cards, a
// sideboard of 15, and a card whose first category is Sideboard and
// whose second is Instant.
func TestFetchSixtyWithSideboard(t *testing.T) {
	var agent string
	deck, err := serve(t, "3045150", "modern-burn.json", &agent).Fetch(context.Background(), 3045150)
	if err != nil {
		t.Fatal(err)
	}
	if agent != UserAgent {
		t.Errorf("agent = %q, want %q", agent, UserAgent)
	}
	if deck.Name != "Modern Artifact Burn" || deck.LeftOut != 0 {
		t.Errorf("name %q, left out %d", deck.Name, deck.LeftOut)
	}
	got := counts(t, deck.Text)
	if got[decklist.Main] != 60 || got[decklist.Sideboard] != 15 || got[decklist.Commander] != 0 {
		t.Errorf("sections = %v, want 60 main and 15 sideboard", got)
	}
	for _, want := range []string{"Deck\n", "\nSideboard\n", "2 Unlicensed Disintegration (f17) 5 *F*\n", "3 Roiling Vortex (znr) 156\n"} {
		if !strings.Contains(deck.Text, want) {
			t.Errorf("text holds no %q:\n%s", want, deck.Text)
		}
	}
}

// TestFetchCommanderLeavesOutTheMaybeboard reads a real Commander deck:
// the premier category leads, the maybeboard stays out, and an etched
// card keeps its mark.
func TestFetchCommanderLeavesOutTheMaybeboard(t *testing.T) {
	var agent string
	deck, err := serve(t, "4406647", "commander.json", &agent).Fetch(context.Background(), 4406647)
	if err != nil {
		t.Fatal(err)
	}
	if deck.LeftOut != 5 {
		t.Errorf("left out = %d, want the 5 maybeboard copies", deck.LeftOut)
	}
	got := counts(t, deck.Text)
	if got[decklist.Commander] != 1 || got[decklist.Main] != 102 {
		t.Errorf("sections = %v, want 1 commander and 102 main", got)
	}
	if !strings.HasPrefix(deck.Text, "Commander\n1 Urabrask // The Great Work (mom) 299 *F*\n\nDeck\n") {
		t.Errorf("text opens with:\n%.120s", deck.Text)
	}
	if n := strings.Count(deck.Text, " *E*\n"); n != 1 {
		t.Errorf("etched lines = %d, want 1", n)
	}
}

// TestFetchCompanionAndDeleted reads the flags that no real fixture
// carries: a companion, a deleted card, a card with no category, and a
// name that holds a line break.
func TestFetchCompanionAndDeleted(t *testing.T) {
	d := apiDeck{Name: "Two\nLines", Categories: []apiCategory{{Name: "Ramp", IncludedInDeck: true}}}
	card := func(name string, qty int, cats []string, companion bool, deleted bool) apiCard {
		c := apiCard{Quantity: qty, Categories: cats, Companion: companion, Modifier: "Normal"}
		c.Card.OracleCard.Name = name
		if deleted {
			s := "2026-10-03T00:00:00Z"
			c.DeletedAt = &s
		}
		return c
	}
	d.Cards = []apiCard{
		card("Lurrus of the Dream-Den", 1, []string{"Ramp"}, true, false),
		card("Sol Ring", 1, []string{"Ramp"}, false, true),
		card("Island", 7, nil, false, false),
		card("Ghost", 0, nil, false, false),
	}
	out := d.list()
	want := "Companion\n1 Lurrus of the Dream-Den\n\nDeck\n7 Island\n"
	if out.Text != want || out.Name != "Two Lines" {
		t.Fatalf("list = %q, name %q; want %q", out.Text, out.Name, want)
	}
}

func TestFetchStatus(t *testing.T) {
	for _, tc := range []struct {
		status int
		body   string
		want   error
	}{
		{http.StatusNotFound, `{"error":"Deck not found."}`, ErrNotFound},
		{http.StatusForbidden, `{}`, ErrNotFound},
		{http.StatusUnauthorized, `{}`, ErrNotFound},
		{http.StatusInternalServerError, `oops`, ErrUnavailable},
		{http.StatusTooManyRequests, ``, ErrUnavailable},
		{http.StatusOK, `<html>`, ErrUnavailable},
	} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(tc.status)
			_, _ = w.Write([]byte(tc.body))
		}))
		_, err := New().WithBaseURL(srv.URL).Fetch(context.Background(), 1)
		srv.Close()
		if !errors.Is(err, tc.want) {
			t.Errorf("status %d: err = %v, want %v", tc.status, err, tc.want)
		}
	}
}

// TestFetchFollowsNoRedirect proves that an answer can not send the read
// to another host.
func TestFetchFollowsNoRedirect(t *testing.T) {
	other := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Error("the client followed the redirect")
	}))
	defer other.Close()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, other.URL, http.StatusFound)
	}))
	defer srv.Close()
	if _, err := New().WithBaseURL(srv.URL).Fetch(context.Background(), 1); !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v, want ErrUnavailable", err)
	}
}

func TestFetchCapsTheAnswer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"name":"` + strings.Repeat("a", maxBodyBytes) + `"}`))
	}))
	defer srv.Close()
	if _, err := New().WithBaseURL(srv.URL).Fetch(context.Background(), 1); !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v, want ErrUnavailable", err)
	}
}

// TestFixturesResolve reads the text of both real decks against the card
// snapshot: every line names a card of the index, except a card of a set
// newer than the snapshot. The fixture of 2026-10-03 holds Artifist
// Acumen of Reality Fracture, released 2026-10-02. The test skips without
// CARDS_SNAPSHOT_DIR, as the snapshot tests of decklist do.
func TestFixturesResolve(t *testing.T) {
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		t.Skip("set CARDS_SNAPSHOT_DIR to run the snapshot tests")
	}
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatal(err)
	}
	// newer holds each card of a fixture that a snapshot before its
	// release does not hold.
	newer := map[string]bool{"Artifist Acumen": true}
	for id, fixture := range map[string]string{"3045150": "modern-burn.json", "4406647": "commander.json"} {
		var agent string
		n, _ := strconv.ParseInt(id, 10, 64)
		deck, err := serve(t, id, fixture, &agent).Fetch(context.Background(), n)
		if err != nil {
			t.Fatal(err)
		}
		list, err := decklist.Parse(strings.NewReader(deck.Text))
		if err != nil {
			t.Fatal(err)
		}
		_, bad := decklist.Resolve(list, idx)
		for _, row := range bad {
			if line, _ := collections.ParseLine(row.GetRaw()); newer[line.Name] {
				t.Logf("%s: %q can be newer than the snapshot", fixture, row.GetRaw())
				continue
			}
			t.Errorf("%s: unresolved %v", fixture, row)
		}
	}
}
