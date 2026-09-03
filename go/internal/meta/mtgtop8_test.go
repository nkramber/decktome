package meta

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func readTestdata(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func sumCopies(cards []Card) int {
	n := 0
	for _, c := range cards {
		n += c.Count
	}
	return n
}

// TestParseMTGTop8Format reads the Modern format page of 2026-09-03: the
// major events and the last twenty, with the paper mark, the stars, and
// the date of each, and the hrefs of the other pages.
func TestParseMTGTop8Format(t *testing.T) {
	events, pages := ParseMTGTop8Format(readTestdata(t, "mtgtop8-format-mo.html"))
	// The fixture holds the two event tables of the page, 21 events once
	// each. A major event sits in both tables and counts once.
	if len(events) != 21 {
		t.Errorf("events = %d, want 21 unique ids", len(events))
	}
	byID := map[string]MTGTop8Event{}
	for _, ev := range events {
		byID[ev.ID] = ev
	}
	spotlight := byID["90258"]
	if !spotlight.Paper || spotlight.Stars != 3 || spotlight.Date != "2026-08-29" || !strings.Contains(spotlight.Name, "Magic Spotlight") {
		t.Errorf("the Brisbane Spotlight reads %+v", spotlight)
	}
	showcase := byID["90275"]
	if showcase.Paper || showcase.Stars != 3 || showcase.Name != "MTGO Showcase Challenge" {
		t.Errorf("the Showcase Challenge reads %+v", showcase)
	}
	league := byID["90377"]
	if league.Paper || league.Stars != 1 || league.Date != "2026-09-02" || league.Name != "MTGO League" {
		t.Errorf("the league reads %+v", league)
	}
	if len(pages) != 3 || pages[0] != "?f=MO&meta=54&cp=2" {
		t.Errorf("pages = %v, want the three later pages", pages)
	}
}

// TestParseMTGTop8Event reads three event pages of 2026-09-03: a paper
// Spotlight with 570 players, an MTGO challenge, and an MTGO league.
func TestParseMTGTop8Event(t *testing.T) {
	info, rows, err := ParseMTGTop8Event(readTestdata(t, "mtgtop8-event-90258.html"))
	if err != nil {
		t.Fatal(err)
	}
	if info.Online || info.Players != 570 || info.Date != "2026-08-29" || info.Format != "MO" || !strings.Contains(info.Name, "Magic Spotlight") {
		t.Errorf("info = %+v", info)
	}
	if len(rows) != 64 {
		t.Fatalf("rows = %d, want 64", len(rows))
	}
	if rows[0] != (MTGTop8Row{Placement: 1, DeckID: "885065", Name: "UR Cutter Prowess", Player: "Jefferson Lee"}) {
		t.Errorf("row 1 = %+v", rows[0])
	}
	if rows[2].Placement != 3 || rows[3].Placement != 3 {
		t.Errorf("a shared place reads its first number: %+v %+v", rows[2], rows[3])
	}
	info, rows, err = ParseMTGTop8Event(readTestdata(t, "mtgtop8-event-90275.html"))
	if err != nil {
		t.Fatal(err)
	}
	if !info.Online || info.Players != 289 || len(rows) != 16 || rows[15].Placement != 16 {
		t.Errorf("the challenge reads %+v with %d rows", info, len(rows))
	}
	info, rows, err = ParseMTGTop8Event(readTestdata(t, "mtgtop8-event-90377.html"))
	if err != nil {
		t.Fatal(err)
	}
	if !info.Online || info.Players != 0 || len(rows) != 8 {
		t.Errorf("the league reads %+v with %d rows", info, len(rows))
	}
	if _, _, err := ParseMTGTop8Event([]byte("<html></html>")); err == nil {
		t.Error("a page with no row must fail")
	}
}

// TestParseMTGTop8Deck reads the text export: a Modern list with its
// sideboard, and a cEDH list whose commander sits under the sideboard.
func TestParseMTGTop8Deck(t *testing.T) {
	main, side := ParseMTGTop8Deck(readTestdata(t, "mtgtop8-deck-885253.txt"))
	if sumCopies(main) != 60 || sumCopies(side) != 15 {
		t.Errorf("modern export: %d main copies, %d side copies", sumCopies(main), sumCopies(side))
	}
	main, side = ParseMTGTop8Deck(readTestdata(t, "mtgtop8-deck-714543.txt"))
	if sumCopies(main) != 99 || len(side) != 1 || side[0].Name != "Atraxa, Grand Unifier" {
		t.Errorf("cedh export: %d main copies, side %+v", sumCopies(main), side)
	}
}

func TestMTGTop8Tier(t *testing.T) {
	for _, tc := range []struct {
		placement, players, stars int
		want                      string
	}{
		{1, 570, 3, TierGreat},
		{8, 32, 1, TierGreat},
		{9, 570, 3, TierGood},
		{1, 12, 1, TierGood},
		{1, 0, 2, TierGreat},
		{1, 0, 1, TierGood},
		{0, 570, 3, TierGood},
	} {
		if got := mtgtop8Tier(tc.placement, tc.players, tc.stars); got != tc.want {
			t.Errorf("tier(%d, %d, %d) = %s, want %s", tc.placement, tc.players, tc.stars, got, tc.want)
		}
	}
}

// TestMTGTop8ListMovesTheCommander: a cEDH list carries its commander as
// the sideboard of the export, and the list moves it to its own field.
func TestMTGTop8ListMovesTheCommander(t *testing.T) {
	text := readTestdata(t, "mtgtop8-deck-714543.txt")
	ev := MTGTop8Event{ID: "68000", Name: "April Competitive EDH 1K", Date: "2025-04-26", Stars: 1}
	info := MTGTop8EventInfo{Players: 26, Format: "cEDH"}
	l, ok := MTGTop8List("cEDH", ev, info, MTGTop8Row{Placement: 1, DeckID: "714543"}, text)
	if !ok {
		t.Fatal("the cEDH list did not build")
	}
	if l.Format != FormatCommander || l.ID != "68000/714543" || len(l.Commanders) != 1 || l.Commanders[0] != "Atraxa, Grand Unifier" || l.Sideboard != nil {
		t.Errorf("list = %+v", l)
	}
	if l.Tier != TierGood || l.Event != "April Competitive EDH 1K" || l.Date != "2025-04-26" {
		t.Errorf("a top finish in a field of 26 is good: %s, %s, %s", l.Tier, l.Event, l.Date)
	}
	if _, ok := MTGTop8List("EDH", ev, info, MTGTop8Row{DeckID: "1"}, text); ok {
		t.Error("Duel Commander is not a format the model covers")
	}
}
