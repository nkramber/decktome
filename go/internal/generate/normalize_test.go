package generate

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

func card(oid, name string) *mtgv1.Card {
	return &mtgv1.Card{OracleId: oid, Name: name}
}

func testPool() *Pool {
	return NewPool([]*mtgv1.Card{
		card("o-welcome", "Ajani's Welcome"),
		card("o-pridemate", "Ajani's Pridemate"),
		card("o-solring", "Sol Ring"),
		card("o-karlov", "Karlov of the Ghost Council"),
	}, map[string]int32{"o-solring": 1, "o-welcome": 4})
}

// TestNormalizeMatchesExactNamesOnly is F-13. A model can name a card
// that exists and is not the card meant, and a fuzzy match hides it.
func TestNormalizeMatchesExactNamesOnly(t *testing.T) {
	p := testPool()
	got := Normalize(p, []Entry{
		{Name: "Ajani's Welcome", Count: 1, Role: "synergy", Reason: "gains life"},
		{Name: "  sol ring  ", Count: 1, Role: "ramp"},
		{Name: "Ajani Welcome", Count: 1, Role: "synergy"},
		{Name: "Craterhoof Behemoth", Count: 1, Role: "wincon"},
	})
	if len(got.Cards) != 2 {
		t.Fatalf("cards = %d, want 2: %+v", len(got.Cards), got.Cards)
	}
	if got.Cards[0].GetOracleId() != "o-welcome" || got.Cards[0].GetRole() != mtgv1.CardRole_CARD_ROLE_SYNERGY {
		t.Errorf("first card = %+v", got.Cards[0])
	}
	// The pool spelling wins, so the deck never shows the model's casing.
	if got.Cards[1].GetName() != "Sol Ring" {
		t.Errorf("name = %q, want the pool spelling", got.Cards[1].GetName())
	}
	if len(got.Misses) != 2 {
		t.Fatalf("misses = %+v, want two", got.Misses)
	}
	if got.Misses[0].Name != "Ajani Welcome" {
		t.Errorf("miss = %q", got.Misses[0].Name)
	}
	if len(got.Misses[0].Near) == 0 {
		t.Error("the near names are empty, so the repair turn has no hint")
	}
	// A near name is reported and never substituted.
	for _, c := range got.Cards {
		if c.GetName() == "Ajani Welcome" {
			t.Fatal("a near name was substituted into the deck")
		}
	}
}

// TestNormalizeReadsTheOwnedCount is D-2. The deck shows what the
// collection covers, and a count above the owned count is not covered.
func TestNormalizeReadsTheOwnedCount(t *testing.T) {
	p := testPool()
	got := Normalize(p, []Entry{
		{Name: "Ajani's Welcome", Count: 4, Role: "synergy"},
		{Name: "Sol Ring", Count: 2, Role: "ramp"},
		{Name: "Karlov of the Ghost Council", Count: 1, Role: "threat"},
	})
	want := []struct {
		owned bool
		count int32
	}{{true, 4}, {false, 1}, {false, 0}}
	for i, w := range want {
		c := got.Cards[i]
		if c.GetOwned() != w.owned || c.GetOwnedCount() != w.count {
			t.Errorf("%s: owned = %v/%d, want %v/%d", c.GetName(), c.GetOwned(), c.GetOwnedCount(), w.owned, w.count)
		}
	}
}

// TestMissNoteNamesTheCard covers the user-visible note. The roadmap
// gives the model one repair turn, and a second miss reaches the user.
func TestMissNoteNamesTheCard(t *testing.T) {
	p := testPool()
	got := Normalize(p, []Entry{{Name: "Ajani Welcome", Count: 1}})
	note := MissNote(got.Misses[0])
	if !strings.Contains(note, "Ajani Welcome") {
		t.Errorf("note = %q, want the card name in it", note)
	}
	if !strings.Contains(note, "Ajani's Welcome") {
		t.Errorf("note = %q, want the near name in it", note)
	}
}

// TestPadWithBasicsFinishesAShortList is D-225. Gate run 1 of 2026-08-27
// returned two Commander decks of 99 cards against a size of 100, and one
// of them asked for a Plains the pool did not hold.
func TestPadWithBasicsFinishesAShortList(t *testing.T) {
	plains := &mtgv1.Card{OracleId: "o-plains", Name: "Plains", Supertypes: []string{"Basic"}}
	swamp := &mtgv1.Card{OracleId: "o-swamp", Name: "Swamp", Supertypes: []string{"Basic"}}
	pool := NewPool([]*mtgv1.Card{card("o-karlov", "Karlov of the Ghost Council"), plains, swamp}, nil)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: pool}

	// 97 cards plus one commander is two short of 100.
	deck := &mtgv1.Deck{
		CommanderOracleIds: []string{"o-karlov"},
		Cards:              []*mtgv1.DeckCard{{OracleId: "o-x", Name: "X", Count: 97}},
	}
	if got := padWithBasics(deck, req); got != 2 {
		t.Fatalf("padded = %d, want 2", got)
	}
	total := 1
	for _, c := range deck.GetCards() {
		total += int(c.GetCount())
	}
	if total != 100 {
		t.Errorf("deck size = %d, want 100", total)
	}
	// The pad spreads across the colors, one of each before a second.
	if len(deck.GetCards()) != 3 {
		t.Errorf("lines = %d, want the two basics on their own lines", len(deck.GetCards()))
	}
}

// TestPadRefusesALargeShortfall keeps a real failure visible. A deck far
// from its size is not a counting slip, and padding would hide it.
func TestPadRefusesALargeShortfall(t *testing.T) {
	plains := &mtgv1.Card{OracleId: "o-plains", Name: "Plains", Supertypes: []string{"Basic"}}
	pool := NewPool([]*mtgv1.Card{plains}, nil)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: pool}
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-x", Name: "X", Count: 40}}}
	if got := padWithBasics(deck, req); got != 0 {
		t.Errorf("padded = %d, want 0: a 60-card gap is a failure and not a slip", got)
	}
}

// TestBasicLandsFollowTheColorIdentity keeps an off-color basic out.
func TestBasicLandsFollowTheColorIdentity(t *testing.T) {
	all := map[string]*mtgv1.Card{
		"Plains": {OracleId: "o-p", Name: "Plains"}, "Island": {OracleId: "o-i", Name: "Island"},
		"Swamp": {OracleId: "o-s", Name: "Swamp"}, "Mountain": {OracleId: "o-m", Name: "Mountain"},
		"Forest": {OracleId: "o-f", Name: "Forest"},
	}
	find := func(n string) (*mtgv1.Card, bool) { c, ok := all[n]; return c, ok }
	got := BasicLands(find, []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B})
	if len(got) != 2 || got[0].GetName() != "Plains" || got[1].GetName() != "Swamp" {
		t.Errorf("basics = %+v, want Plains and Swamp", got)
	}
	if len(BasicLands(find, nil)) != 0 {
		t.Error("a colorless deck got a basic land")
	}
}

// TestDeckCardsCarryTheirPrice is D-236. The deck showed no price at all,
// so a budget deck could not be checked or reported.
func TestDeckCardsCarryTheirPrice(t *testing.T) {
	pricey := &mtgv1.Card{OracleId: "o-dv", Name: "Diamond Valley", PriceUsd: 693.33}
	cheap := &mtgv1.Card{OracleId: "o-sr", Name: "Sol Ring", PriceUsd: 1.50}
	pool := NewPool([]*mtgv1.Card{pricey, cheap}, map[string]int32{"o-sr": 1})
	got := Normalize(pool, []Entry{
		{Name: "Diamond Valley", Count: 1, Role: "land"},
		{Name: "Sol Ring", Count: 2, Role: "ramp"},
	})
	if got.Cards[0].GetPriceUsd() != 693.33 {
		t.Errorf("price = %v, want the card's price", got.Cards[0].GetPriceUsd())
	}
	deck := &mtgv1.Deck{Cards: got.Cards}
	// The buy cost counts only what the collection does not cover: one
	// Diamond Valley, and one of the two Sol Rings.
	if want := 693.33 + 1.50; BuyCost(deck) != want {
		t.Errorf("buy cost = %v, want %v", BuyCost(deck), want)
	}
	// The whole deck counts every copy.
	if want := 693.33 + 3.00; DeckCost(deck) != want {
		t.Errorf("deck cost = %v, want %v", DeckCost(deck), want)
	}
}

// TestOverBudgetWarnsAndNeverBlocks is D-236. The price is a daily
// estimate and not a rule of the game, and a deck the user can trim is
// more use than no deck.
func TestOverBudgetWarnsAndNeverBlocks(t *testing.T) {
	b, _, _ := testBuilder(t, step(t, deckOut{Summary: "s", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 1, Role: "synergy"},
	}}), step(t, deckOut{Summary: "s", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 1, Role: "synergy"},
	}}))
	req := testRequest()
	req.BudgetUSD = 10
	req.Pool = NewPool([]*mtgv1.Card{{OracleId: "o-w", Name: "Ajani's Welcome", PriceUsd: 50}}, nil)
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	var found *mtgv1.Finding
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeOverBudget {
			found = f
		}
	}
	if found == nil {
		t.Fatal("a deck over budget carried no finding")
	}
	if found.GetSeverity() != mtgv1.Severity_SEVERITY_WARN {
		t.Errorf("severity = %v, want WARN: a price estimate must not block a deck", found.GetSeverity())
	}
}
