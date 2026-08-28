package generate

import (
	"context"
	"fmt"
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

// TestBudgetScopeChoosesTheCost is D-238. The budget-scope row asks
// whether a cap covers the cards to buy or the whole deck, and nothing
// stored the answer, so the agent asked and discarded it.
func TestBudgetScopeChoosesTheCost(t *testing.T) {
	// One owned copy and one to buy, each $50. Buying costs $50, and the
	// whole deck is worth $100.
	pool := NewPool([]*mtgv1.Card{{OracleId: "o-w", Name: "Ajani's Welcome", PriceUsd: 50}},
		map[string]int32{"o-w": 1})
	entry := []Entry{{Name: "Ajani's Welcome", Count: 2, Role: "synergy"}}
	for _, tc := range []struct {
		name  string
		whole bool
		cap   float64
		want  bool
	}{
		{"buy scope, under the cap", false, 75, false},
		{"whole-deck scope, over the cap", true, 75, true},
		{"whole-deck scope, under the cap", true, 150, false},
	} {
		one := step(t, deckOut{Summary: "s", Cards: entry})
		b, _, _ := testBuilder(t, one, one)
		req := testRequest()
		req.Pool, req.BudgetUSD, req.BudgetWholeDeck = pool, tc.cap, tc.whole
		got, err := b.Build(context.Background(), req, nil)
		if err != nil {
			t.Fatalf("%s: build: %v", tc.name, err)
		}
		var over bool
		for _, f := range got.Deck.GetValidation().GetFindings() {
			if f.GetCode() == CodeOverBudget {
				over = true
			}
		}
		if over != tc.want {
			t.Errorf("%s: over budget = %v, want %v", tc.name, over, tc.want)
		}
	}
}

// TestBudgetReachesThePrompt is D-244. Deck gate run 5 spent $268.37
// against a $100.00 cap, on a shortlist whose cheapest 99 cards cost
// $25.66. No shortlist line carried a price, so the model was guessing.
func TestBudgetReachesThePrompt(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-cheap", Name: "Cheap Card", PriceUsd: 0.25},
		{OracleId: "o-dear", Name: "Dear Card", PriceUsd: 53.68},
	}, map[string]int32{"o-cheap": 1})
	b := &Builder{}
	req := testRequest()
	req.Pool, req.BudgetUSD, req.OracleCounts = pool, 100, map[string]int32{"o-cheap": 1}
	in := b.input(req, nil, nil)
	for _, want := range []string{"$0.25", "$53.68", "## Budget", "$100.00", "owned"} {
		if !strings.Contains(in, want) {
			t.Errorf("the prompt does not carry %q", want)
		}
	}
	// With no budget the prices stay out, so an ordinary build pays no
	// tokens for them.
	req.BudgetUSD = 0
	if got := b.input(req, nil, nil); strings.Contains(got, "$53.68") {
		t.Error("a build with no budget still carried prices")
	}
}

// TestOverBudgetBuysTheRepairTurn is D-244. The finding stays a warning,
// because a price is an estimate and not a rule, and the model still gets
// one chance to come under the cap.
func TestOverBudgetBuysTheRepairTurn(t *testing.T) {
	dear := deckOut{Summary: "s", Cards: []Entry{{Name: "Dear Card", Count: 1, Role: "synergy"}}}
	cheap := deckOut{Summary: "s", Cards: []Entry{{Name: "Cheap Card", Count: 1, Role: "synergy"}}}
	b, _, sc := testBuilder(t, step(t, dear), step(t, cheap))
	req := testRequest()
	req.Pool = NewPool([]*mtgv1.Card{
		{OracleId: "o-cheap", Name: "Cheap Card", PriceUsd: 0.25},
		{OracleId: "o-dear", Name: "Dear Card", PriceUsd: 53.68},
	}, nil)
	req.BudgetUSD = 10
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(sc.Calls) != 2 {
		t.Fatalf("provider calls = %d, want 2: over budget buys the repair turn", len(sc.Calls))
	}
	if !got.Repaired {
		t.Error("the repair turn did not run on an over-budget deck")
	}
	// The repair came under the cap, so no warning survives.
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeOverBudget {
			t.Errorf("the repaired deck is still over budget: %s", f.GetMessage())
		}
	}
	// The repair input must name the cost, or the model cannot fix it.
	if !strings.Contains(sc.Calls[1].Input, "cost") {
		t.Error("the repair input did not name the cost")
	}
}

// TestPreconKeepCountRoundsUp is D-248. The share is met and not
// approached, so a fraction of a card rounds up.
func TestPreconKeepCountRoundsUp(t *testing.T) {
	for _, tc := range []struct{ total, want int }{
		{0, 0}, {100, 85}, {79, 68}, {93, 80}, {87, 74}, {10, 9},
	} {
		if got := PreconKeepCount(tc.total); got != tc.want {
			t.Errorf("PreconKeepCount(%d) = %d, want %d", tc.total, got, tc.want)
		}
	}
}

// TestShortPreconBuysTheRepairTurn is D-248. An upgrade that keeps too
// little of the precon is not the deck the user asked for, and the model
// gets one chance to put the cards back.
func TestShortPreconBuysTheRepairTurn(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-a", Name: "Card A"}, {OracleId: "o-b", Name: "Card B"},
		{OracleId: "o-c", Name: "Card C"}, {OracleId: "o-d", Name: "Card D"},
	}, nil)
	// Two of the four precon cards is under the keep count of four.
	few := deckOut{Summary: "s", Cards: []Entry{
		{Name: "Card A", Count: 1, Role: "synergy"}, {Name: "Card B", Count: 1, Role: "synergy"},
	}}
	all := deckOut{Summary: "s", Cards: []Entry{
		{Name: "Card A", Count: 1, Role: "synergy"}, {Name: "Card B", Count: 1, Role: "synergy"},
		{Name: "Card C", Count: 1, Role: "synergy"}, {Name: "Card D", Count: 1, Role: "synergy"},
	}}
	b, _, sc := testBuilder(t, step(t, few), step(t, all))
	req := testRequest()
	req.Pool = pool
	req.Precon = "Test Precon"
	req.PreconOracleIDs = []string{"o-a", "o-b", "o-c", "o-d"}
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(sc.Calls) != 2 {
		t.Fatalf("provider calls = %d, want 2: a short precon buys the repair turn", len(sc.Calls))
	}
	// The repair put the cards back, so no finding survives.
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodePreconShare {
			t.Errorf("the repaired deck is still short: %s", f.GetMessage())
		}
	}
	// The prompt must state the count, or the model is guessing.
	if !strings.Contains(sc.Calls[0].Input, "Keep at least 4 of them") {
		t.Error("the prompt did not state how many precon cards to keep")
	}
}

// TestAnUpgradeGetsNoJobTargets is D-249. The targets prescribe the whole
// deck and the share demands most of its slots, so the two fought.
func TestAnUpgradeGetsNoJobTargets(t *testing.T) {
	b := &Builder{}
	req := testRequest()
	req.Targets = map[string]int{"land": 36, "ramp": 10}
	if got := b.input(req, nil, nil); !strings.Contains(got, "Job targets") {
		t.Error("an ordinary build lost its job targets")
	}
	req.Precon = "Goblin Storm"
	req.PreconOracleIDs = []string{"o-a", "o-b"}
	// The list states the land count, so nothing guesses at it (D-251).
	req.PreconLands = 34
	got := b.input(req, nil, nil)
	// No job target reaches an upgrade: each one is a quota against the
	// share. The mana base comes from the precon's own count instead, so
	// the two do not fight (D-251).
	if strings.Contains(got, "Job targets") {
		t.Error("an upgrade was given job targets, which fight the share")
	}
	if !strings.Contains(got, "The precon holds 34 lands") {
		t.Error("the upgrade prompt does not name the precon's land count")
	}
	if !strings.Contains(got, "Keep at least 2 of them") {
		t.Error("the upgrade prompt does not state the keep count")
	}
}

// TestSwapBackPreconClosesASmallShortfall is D-250. The model repaired a
// deck to 67 of the 68 precon cards it needed, read the finding that said
// so, and returned 67 again. It can not count its own list reliably.
func TestSwapBackPreconClosesASmallShortfall(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-p1", Name: "Precon One"}, {OracleId: "o-p2", Name: "Precon Two"},
		{OracleId: "o-p3", Name: "Precon Three"}, {OracleId: "o-p4", Name: "Precon Four"},
		{OracleId: "o-x", Name: "Outsider"},
		{OracleId: "o-plains", Name: "Plains", Supertypes: []string{"Basic"}},
	}, nil)
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: pool,
		Precon:          "Test Precon",
		PreconOracleIDs: []string{"o-p1", "o-p2", "o-p3", "o-p4"},
	}
	// Three of four precon cards, one outsider, and a basic. The keep
	// count is four, so one card is missing.
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-p1", Name: "Precon One", Count: 1},
		{OracleId: "o-p2", Name: "Precon Two", Count: 1},
		{OracleId: "o-p3", Name: "Precon Three", Count: 1},
		{OracleId: "o-x", Name: "Outsider", Count: 1},
		{OracleId: "o-plains", Name: "Plains", Count: 30},
	}}
	if got := swapBackPrecon(deck, req); got != 1 {
		t.Fatalf("swapped = %d, want 1", got)
	}
	names := map[string]bool{}
	size := 0
	for _, c := range deck.GetCards() {
		names[c.GetName()] = true
		size += int(c.GetCount())
	}
	if !names["Precon Four"] {
		t.Error("the missing precon card was not put back")
	}
	if names["Outsider"] {
		t.Error("the outsider was not the card traded out")
	}
	// The basic land and the deck size must not move.
	if !names["Plains"] || size != 34 {
		t.Errorf("the mana base moved: names %v size %d", names, size)
	}
}

// TestSwapBackRefusesALargeShortfall keeps a real failure visible. A deck
// far from the share is a different deck, not a counting slip.
func TestSwapBackRefusesALargeShortfall(t *testing.T) {
	var ids []string
	var cards []*mtgv1.Card
	for i := 0; i < 20; i++ {
		id := fmt.Sprintf("o-p%d", i)
		ids = append(ids, id)
		cards = append(cards, &mtgv1.Card{OracleId: id, Name: fmt.Sprintf("Precon %d", i)})
	}
	req := Request{Pool: NewPool(cards, nil), Precon: "Big", PreconOracleIDs: ids}
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-p0", Name: "Precon 0", Count: 1}}}
	if got := swapBackPrecon(deck, req); got != 0 {
		t.Errorf("swapped = %d, want 0: a 16-card gap is a different deck", got)
	}
}
