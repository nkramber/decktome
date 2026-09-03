package meta

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const topdeckAnswer = `[
 {"TID": "t1", "tournamentName": "Weekly cEDH", "startDate": 1756771200, "topCut": 0, "standings": [
   {"name": "A", "id": "a", "wins": 4, "losses": 0, "draws": 1, "winRate": 0.8,
    "deckObj": {"Commanders": {"Kinnan, Bonder Prodigy": {"id": "x", "count": 1}}, "Mainboard": {"Sol Ring": {"count": 1}, "Forest": 3}, "metadata": {"format": "EDH"}}},
   {"name": "B", "id": "b", "deckObj": "None", "decklist": "~~Commanders~~\n1 Thrasios, Triton Hero\n1 Tymna the Weaver\n~~Mainboard~~\n1 Sol Ring\n2x Island\n~~Sideboard~~\n1 Swamp"},
   {"name": "C", "id": "c", "deckObj": "None", "decklist": "None"},
   {"name": "D", "id": "d", "deckObj": "None", "decklist": "https://moxfield.com/decks/x"},
   {"name": "E", "id": "e", "deckObj": "None", "decklist": "None"},
   {"name": "F", "id": "f", "deckObj": "None", "decklist": "None"},
   {"name": "G", "id": "g", "deckObj": "None", "decklist": "None"},
   {"name": "H", "id": "h", "deckObj": "None", "decklist": "None"},
   {"name": "I", "id": "i", "decklist": "~~Mainboard~~\n1 Sol Ring"}
 ]},
 {"TID": "test", "tournamentName": "Test event", "startDate": 1919260800, "standings": [
   {"name": "Z", "id": "z", "decklist": "~~Mainboard~~\n1 Sol Ring"}
 ]},
 {"TID": "small", "tournamentName": "Pod", "startDate": 1756771200, "standings": [
   {"name": "Y", "id": "y", "decklist": "~~Mainboard~~\n1 Sol Ring"}
 ]}
]`

func TestNewTopdeckNeedsKey(t *testing.T) {
	if _, err := NewTopdeck(nil, "", "", nil); err == nil {
		t.Fatal("an empty key must refuse the client")
	}
}

func TestTopdeckTournaments(t *testing.T) {
	var gotAuth, gotBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		body, _ := io.ReadAll(r.Body)
		gotBody = string(body)
		if r.URL.Path != "/v2/tournaments" || r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write([]byte(topdeckAnswer))
	}))
	defer srv.Close()
	fetch := NewFetcher(srv.Client(), time.Millisecond, nil)
	td, err := NewTopdeck(fetch, "secret", srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	td.MinPlayers = 9
	td.Now = func() time.Time { return time.Date(2026, 9, 2, 12, 0, 0, 0, time.UTC) }
	ts, err := td.Tournaments(context.Background(), TopdeckFormatEDH, 20)
	if err != nil {
		t.Fatal(err)
	}
	if gotAuth != "secret" {
		t.Errorf("authorization = %q", gotAuth)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(gotBody), &body); err != nil {
		t.Fatal(err)
	}
	if body["format"] != "EDH" || body["game"] != TopdeckGame || body["participantMin"] != float64(9) || body["start"] == nil || body["end"] == nil {
		t.Errorf("body = %s", gotBody)
	}
	// Three weekly slices serve the same answer, so the tournament
	// counts once. The future-dated test event and the small pod are out.
	if len(ts) != 1 || ts[0].Players != 9 || ts[0].Date() != "2025-09-02" || ts[0].TID != "t1" {
		t.Fatalf("tournaments = %+v", ts)
	}
	lists := ts[0].Lists()
	if len(lists) != 3 {
		t.Fatalf("lists = %d, want 3 (the rows with no deck and the link are out)", len(lists))
	}
	a := lists[0]
	if a.Tier != TierGreat || len(a.Commanders) != 1 || a.Commanders[0] != "Kinnan, Bonder Prodigy" {
		t.Errorf("a = tier %s, commanders %v", a.Tier, a.Commanders)
	}
	if len(a.Cards) != 2 || a.Size() != 4 || a.Wins != 4 {
		t.Errorf("a cards = %+v, wins %d", a.Cards, a.Wins)
	}
	b := lists[1]
	if b.Tier != TierGreat || len(b.Commanders) != 2 || b.Size() != 3 || len(b.Sideboard) != 0 {
		t.Errorf("b = tier %s, commanders %v, cards %+v", b.Tier, b.Commanders, b.Cards)
	}
	c := lists[2]
	if c.Tier != TierGood || c.Placement != 9 {
		t.Errorf("c = tier %s, placement %d", c.Tier, c.Placement)
	}
	if fetch.MaxBody != TopdeckMaxBody {
		t.Errorf("max body = %d", fetch.MaxBody)
	}
	if a.ID != "t1/a" || a.Event != "Weekly cEDH" || a.Format != FormatCommander {
		t.Errorf("a = id %s, event %s, format %s", a.ID, a.Event, a.Format)
	}
}

func TestTopCutSize(t *testing.T) {
	tests := []struct {
		players, topCut, want int
	}{
		{8, 0, 4}, {16, 0, 4}, {17, 0, 8}, {64, 0, 8}, {65, 0, 16}, {100, 32, 32},
	}
	for _, tt := range tests {
		tr := Tournament{Players: tt.players, TopCut: tt.topCut}
		if got := tr.TopCutSize(); got != tt.want {
			t.Errorf("%d players, topCut %d: %d, want %d", tt.players, tt.topCut, got, tt.want)
		}
	}
}

func TestParseTopdeckDeckNone(t *testing.T) {
	if cs, cards := ParseTopdeckDeck(json.RawMessage(`"None"`), "None"); cs != nil || cards != nil {
		t.Errorf("None read as a deck: %v %v", cs, cards)
	}
	if _, cards := ParseTopdeckDeck(nil, "https://moxfield.com/decks/x"); cards != nil {
		t.Errorf("a link read as a deck: %v", cards)
	}
}
