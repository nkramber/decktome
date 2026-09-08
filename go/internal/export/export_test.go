package export

import (
	"os"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/collections"
)

func loadIndex(t *testing.T) *cards.Index {
	t.Helper()
	f, err := os.Open("../cards/testdata/cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	list, err := cards.LoadCards(f, "cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	return cards.NewIndex(list, nil, nil, time.Date(2026, 8, 24, 9, 1, 0, 0, time.UTC))
}

func card(t *testing.T, idx *cards.Index, name string) *mtgv1.Card {
	t.Helper()
	c, ok := idx.ByName(name)
	if !ok {
		t.Fatalf("fixture card %q missing", name)
	}
	return c
}

// twoFaced finds a paper card with two faces in the fixture, so the
// round trip covers a full name with " // ".
func twoFaced(t *testing.T, idx *cards.Index) *mtgv1.Card {
	t.Helper()
	for _, c := range idx.All() {
		if len(c.GetFaces()) == 2 && !strings.HasPrefix(c.GetName(), "A-") && !c.GetDefaultPrinting().GetDigital() {
			return c
		}
	}
	t.Skip("the fixture holds no paper two-faced card")
	return nil
}

func entry(c *mtgv1.Card, n int32) *mtgv1.DeckCard {
	return &mtgv1.DeckCard{OracleId: c.OracleId, Name: c.Name, Count: n}
}

// TestArenaTextRoundTrip: the PR-13 gate. A deck exported as Arena text
// and read back by ParseArenaText loses nothing (D-15).
func TestArenaTextRoundTrip(t *testing.T) {
	idx := loadIndex(t)
	sol := card(t, idx, "Sol Ring")
	sw := card(t, idx, "Soul Warden")
	pl := card(t, idx, "Plains")
	dfc := twoFaced(t, idx)
	cmd := card(t, idx, "Anikthea, Hand of Erebos")
	d := &mtgv1.Deck{
		Name:               "Anikthea test deck",
		CommanderOracleIds: []string{cmd.OracleId},
		Cards:              []*mtgv1.DeckCard{entry(cmd, 1), entry(sol, 1), entry(dfc, 1), entry(pl, 30)},
		Sideboard:          []*mtgv1.DeckCard{entry(sw, 2)},
	}
	text := ArenaText(d, idx)
	if !strings.HasPrefix(text, "Commander\n1 "+cmd.Name+" (") {
		t.Fatalf("commander header missing:\n%s", text)
	}
	if !strings.Contains(text, "\nDeck\n") || !strings.Contains(text, "\nSideboard\n2 Soul Warden (") {
		t.Fatalf("section headers wrong:\n%s", text)
	}
	if strings.Count(text, cmd.Name) != 1 {
		t.Fatalf("the commander must appear once:\n%s", text)
	}

	rows, bad, err := collections.ParseArenaText(strings.NewReader(text))
	if err != nil || len(bad) != 0 {
		t.Fatalf("parse: err=%v bad=%v", err, bad)
	}
	want := map[string]int{cmd.OracleId: 1, sol.OracleId: 1, dfc.OracleId: 1, pl.OracleId: 30, sw.OracleId: 2}
	got := map[string]int{}
	for _, r := range rows {
		c, ok := idx.ByName(r.Name)
		if !ok {
			t.Fatalf("row %q does not resolve by name", r.Name)
		}
		got[c.OracleId] += r.Quantity
		dp := c.GetDefaultPrinting()
		if !strings.EqualFold(r.SetCode, dp.GetSetCode()) || r.Collector != dp.GetCollectorNumber() {
			t.Errorf("%s: printing %s/%s, want %s/%s", r.Name, r.SetCode, r.Collector, dp.GetSetCode(), dp.GetCollectorNumber())
		}
	}
	if len(got) != len(want) {
		t.Fatalf("got %d cards, want %d", len(got), len(want))
	}
	for id, n := range want {
		if got[id] != n {
			t.Errorf("%s: count %d, want %d", id, got[id], n)
		}
	}
}

func TestArenaTextNamesTheOwnedPrinting(t *testing.T) {
	idx := loadIndex(t)
	sol := card(t, idx, "Sol Ring")
	owned := entry(sol, 1)
	owned.Owned = true
	owned.OwnedPrinting = &mtgv1.Printing{SetCode: "c21", CollectorNumber: "263"}
	d := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{owned}}
	if got := ArenaText(d, idx); got != "Deck\n1 Sol Ring (C21) 263\n" {
		t.Fatalf("got %q", got)
	}
	// An unowned card keeps the default paper printing, even when a
	// stale owned printing is on the entry.
	owned.Owned = false
	dp := sol.GetDefaultPrinting()
	want := "Deck\n1 Sol Ring (" + strings.ToUpper(dp.GetSetCode()) + ") " + dp.GetCollectorNumber() + "\n"
	if got := ArenaText(d, idx); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestArenaTextWithoutCardData(t *testing.T) {
	d := &mtgv1.Deck{
		CommanderOracleIds: []string{"o-unknown"},
		Cards:              []*mtgv1.DeckCard{{OracleId: "o-x", Name: "Some Card", Count: 3}},
	}
	got := ArenaText(d, emptyLookup{})
	if got != "Commander\n1 \n\nDeck\n3 Some Card\n" {
		t.Fatalf("got %q", got)
	}
}

type emptyLookup struct{}

func (emptyLookup) ByOracleID(string) (*mtgv1.Card, bool) { return nil, false }

func TestBuyListSumsTheShortfall(t *testing.T) {
	idx := loadIndex(t)
	sol := card(t, idx, "Sol Ring")
	sw := card(t, idx, "Soul Warden")
	pl := card(t, idx, "Plains")
	cmd := card(t, idx, "Anikthea, Hand of Erebos")
	main := entry(sw, 4)
	main.OwnedCount = 3
	side := entry(sw, 2)
	side.OwnedCount = 0
	ownedSol := entry(sol, 1)
	ownedSol.Owned = true
	plains := entry(pl, 30)
	plains.Owned = true
	d := &mtgv1.Deck{
		CommanderOracleIds: []string{cmd.OracleId},
		Cards:              []*mtgv1.DeckCard{main, ownedSol, plains},
		Sideboard:          []*mtgv1.DeckCard{side},
		Upgrades:           []*mtgv1.DeckCard{entry(sol, 1)},
	}
	needed, upgrades := BuyList(d, idx)
	// F-76: a deck that carries no entry for its commander carries no
	// ownership fact either, so the commander reaches no buy list. The
	// invented entry read "not owned", and every deck named its
	// commander as a card to buy (D-604).
	if len(needed) != 1 {
		t.Fatalf("needed %+v", needed)
	}
	if needed[0].Name != "Soul Warden" || needed[0].Count != 3 {
		t.Errorf("Soul Warden shortfall 1 main + 2 side = 3: %+v", needed[0])
	}
	if needed[0].SetCode == "" || needed[0].CollectorNumber == "" {
		t.Errorf("a row carries the default printing for its link: %+v", needed[0])
	}
	if len(upgrades) != 1 || upgrades[0].Name != "Sol Ring" {
		t.Errorf("upgrades %+v", upgrades)
	}
	text := BuyListText(d, idx)
	if !strings.HasSuffix(text, "3 Soul Warden\n\nUpgrades\n1 Sol Ring\n") {
		t.Fatalf("text %q", text)
	}
	// The deck carries the fact now (D-608). An unowned commander is a
	// row, and an owned one is not.
	unowned := entry(cmd, 1)
	d.Commanders = []*mtgv1.DeckCard{unowned}
	needed, _ = BuyList(d, idx)
	if len(needed) != 2 || needed[0].Name != cmd.Name || needed[0].Count != 1 {
		t.Errorf("an unowned commander is a buy row of one: %+v", needed)
	}
	owned := entry(cmd, 1)
	owned.Owned, owned.OwnedCount = true, 1
	d.Commanders = []*mtgv1.DeckCard{owned}
	needed, _ = BuyList(d, idx)
	if len(needed) != 1 || needed[0].Name != "Soul Warden" {
		t.Errorf("an owned commander reached the buy list: %+v", needed)
	}
}

func TestBuyListTextWhenNothingIsNeeded(t *testing.T) {
	idx := loadIndex(t)
	pl := card(t, idx, "Plains")
	e := entry(pl, 30)
	e.Owned = true
	got := BuyListText(&mtgv1.Deck{Cards: []*mtgv1.DeckCard{e}}, idx)
	if got != "// Nothing to buy: every card is owned.\n" {
		t.Fatalf("got %q", got)
	}
}

func TestRenderAndFileName(t *testing.T) {
	d := &mtgv1.Deck{Name: "Ghalta's Big Stompy Deck!!"}
	text, name, err := Render(d, emptyLookup{}, mtgv1.ExportFormat_EXPORT_FORMAT_UNSPECIFIED)
	if err != nil || name != "ghalta-s-big-stompy-deck.txt" || !strings.HasPrefix(text, "Deck\n") {
		t.Fatalf("arena: %q %q %v", text, name, err)
	}
	_, name, err = Render(d, emptyLookup{}, mtgv1.ExportFormat_EXPORT_FORMAT_BUY_LIST_TEXT)
	if err != nil || name != "ghalta-s-big-stompy-deck-buy-list.txt" {
		t.Fatalf("buy list: %q %v", name, err)
	}
	if got := FileName(&mtgv1.Deck{}, mtgv1.ExportFormat_EXPORT_FORMAT_ARENA_TEXT); got != "deck.txt" {
		t.Fatalf("empty name: %q", got)
	}
	if _, _, err := Render(d, emptyLookup{}, mtgv1.ExportFormat(99)); err == nil {
		t.Fatal("an unknown format must fail")
	}
}
