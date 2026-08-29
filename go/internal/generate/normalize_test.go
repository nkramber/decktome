package generate

import (
	"context"
	"fmt"
	"math"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
)

func card(oid, name string) *mtgv1.Card {
	return &mtgv1.Card{OracleId: oid, Name: name}
}

func basic(oid, name string) *mtgv1.Card {
	return &mtgv1.Card{OracleId: oid, Name: name, Supertypes: []string{"Basic"}, CardTypes: []string{"Land"}}
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

// TestPadWithBasicsFinishesAShortList is D-225. A deck one or two cards
// short is a counting slip, and a basic land is always a legal answer.
func TestPadWithBasicsFinishesAShortList(t *testing.T) {
	plains := basic("o-plains", "Plains")
	swamp := basic("o-swamp", "Swamp")
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

// TestBudgetReachesThePrompt is D-244. Every shortlist line carries a
// price when a budget applies, or the model is guessing.
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
	if !strings.Contains(got.RepairReason, CodeOverBudget) {
		t.Errorf("repair reason = %q, want it to name %s", got.RepairReason, CodeOverBudget)
	}
	// The repair came under the cap, so no warning survives.
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeOverBudget {
			t.Errorf("the repaired deck is still over budget: %s", f.GetMessage())
		}
	}
	// The repair input must carry the finding, or the model cannot fix
	// it. It is a warning, and it buys the repair turn (D-244).
	for _, want := range []string{CodeOverBudget, "cost about $53.68, and the budget is $10.00"} {
		if !strings.Contains(sc.Calls[1].Input, want) {
			t.Errorf("the repair input does not carry %q", want)
		}
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
	if !strings.Contains(sc.Calls[0].Input, "Keep at least 4 of those names") {
		t.Error("the prompt did not state how many precon cards to keep")
	}
	// The repair input must carry the finding, or the model cannot fix
	// it (D-248).
	for _, want := range []string{CodePreconShare, "keeps 2 of the 4 nonbasic Test Precon precon names"} {
		if !strings.Contains(sc.Calls[1].Input, want) {
			t.Errorf("the repair input does not carry %q", want)
		}
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
	if !strings.Contains(got, "Keep at least 2 of those names") {
		t.Error("the upgrade prompt does not state the keep count")
	}
	// The change count is a ceiling and not a target (A-5).
	if !strings.Contains(got, "a ceiling and not a target") {
		t.Error("the upgrade prompt does not say the change count is a ceiling")
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
	if got := swapBackPrecon(deck, req, nil); got != 1 {
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
	if got := swapBackPrecon(deck, req, nil); got != 0 {
		t.Errorf("swapped = %d, want 0: a 16-card gap is a different deck", got)
	}
}

// TestPreconShareCountsNonbasicNames is D-218. The share is 85 percent of
// the precon's nonbasic names, and a basic land swapped for another
// basic has not dropped the precon.
func TestPreconShareCountsNonbasicNames(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-p1", Name: "Precon One"}, {OracleId: "o-p2", Name: "Precon Two"},
		{OracleId: "o-p3", Name: "Precon Three"}, {OracleId: "o-p4", Name: "Precon Four"},
		basic("o-swamp", "Swamp"), basic("o-forest", "Forest"),
	}, nil)
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: pool, Precon: "Test Precon",
		PreconOracleIDs: []string{"o-p1", "o-p2", "o-p3", "o-p4", "o-swamp"},
	}
	cases := []struct {
		name  string
		cards []*mtgv1.DeckCard
		short bool
	}{
		{"every name and the basics swapped", []*mtgv1.DeckCard{
			{OracleId: "o-p1", Count: 1}, {OracleId: "o-p2", Count: 1},
			{OracleId: "o-p3", Count: 1}, {OracleId: "o-p4", Count: 1},
			{OracleId: "o-forest", Count: 30},
		}, false},
		{"one name short with every basic kept", []*mtgv1.DeckCard{
			{OracleId: "o-p1", Count: 1}, {OracleId: "o-p2", Count: 1},
			{OracleId: "o-p3", Count: 1}, {OracleId: "o-swamp", Count: 30},
		}, true},
	}
	for _, tc := range cases {
		deck := &mtgv1.Deck{Cards: tc.cards, Validation: &mtgv1.ValidationResult{Passed: true}}
		if got := checkPreconShare(deck, req, nil); got != tc.short {
			t.Errorf("%s: short = %v, want %v", tc.name, got, tc.short)
		}
		if got := len(preconNonbasics(req, nil)); got != 4 {
			t.Errorf("%s: nonbasic names = %d, want 4", tc.name, got)
		}
	}
	// The prompt states the nonbasic count and not the card count.
	b := &Builder{}
	in := b.input(req, nil, nil)
	for _, want := range []string{"nonbasic cards are 4 names", "Keep at least 4 of those names", "Basic lands are not counted"} {
		if !strings.Contains(in, want) {
			t.Errorf("the prompt does not say %q", want)
		}
	}
}

// TestAddFindingKeepsPassedTrue: a BLOCK the builder adds after the
// engine ran clears Passed (D-242).
func TestAddFindingKeepsPassedTrue(t *testing.T) {
	without := deckOut{Summary: "s", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 1, Role: "synergy", Reason: "gains life"},
	}}
	b, _, _ := testBuilder(t, step(t, without), step(t, without))
	req := testRequest()
	req.Locked = []string{"o-solring"}
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	v := got.Deck.GetValidation()
	if v.GetPassed() {
		t.Error("a deck with a locked_card_missing block reports passed = true")
	}
	// The pure helper, over every severity.
	for _, tc := range []struct {
		sev  mtgv1.Severity
		pass bool
	}{
		{mtgv1.Severity_SEVERITY_INFO, true},
		{mtgv1.Severity_SEVERITY_WARN, true},
		{mtgv1.Severity_SEVERITY_BLOCK, false},
	} {
		deck := &mtgv1.Deck{Validation: &mtgv1.ValidationResult{Passed: true}}
		addFinding(deck, "x", tc.sev, "m")
		if deck.GetValidation().GetPassed() != tc.pass || len(deck.GetValidation().GetFindings()) != 1 {
			t.Errorf("%v: passed = %v, want %v", tc.sev, deck.GetValidation().GetPassed(), tc.pass)
		}
	}
}

// TestInsertedCardsReadTheOwnedCount is D-37. A card the builder inserts
// reads the owned count from the pool, so the user's own precon cards
// never count as purchases. The commander is priced from the card index
// and owned from the collection.
func TestInsertedCardsReadTheOwnedCount(t *testing.T) {
	plains := basic("o-plains", "Plains")
	plains.PriceUsd = 0.10
	p4 := &mtgv1.Card{OracleId: "o-p4", Name: "Precon Four", PriceUsd: 12}
	commander := &mtgv1.Card{OracleId: "o-c", Name: "The Commander", PriceUsd: 30}
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-p1", Name: "Precon One"}, {OracleId: "o-p2", Name: "Precon Two"},
		{OracleId: "o-p3", Name: "Precon Three"}, p4,
		{OracleId: "o-x", Name: "Outsider"}, plains,
	}, map[string]int32{"o-p4": 1, "o-plains": 2})
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: pool, Precon: "Test Precon",
		PreconOracleIDs: []string{"o-p1", "o-p2", "o-p3", "o-p4"},
	}
	// 96 cards and one commander: three short, so three Plains are added.
	deck := &mtgv1.Deck{CommanderOracleIds: []string{"o-c"}, Cards: []*mtgv1.DeckCard{
		{OracleId: "o-p1", Name: "Precon One", Count: 1},
		{OracleId: "o-p2", Name: "Precon Two", Count: 1},
		{OracleId: "o-p3", Name: "Precon Three", Count: 1},
		{OracleId: "o-x", Name: "Outsider", Count: 1},
		{OracleId: "o-y", Name: "Filler", Count: 92},
	}}
	if got := padWithBasics(deck, req); got != 3 {
		t.Fatalf("padded = %d, want 3", got)
	}
	if got := swapBackPrecon(deck, req, nil); got != 1 {
		t.Fatalf("swapped = %d, want 1", got)
	}
	by := map[string]*mtgv1.DeckCard{}
	for _, c := range deck.GetCards() {
		by[c.GetName()] = c
	}
	for _, tc := range []struct {
		name  string
		owned bool
		count int32
		price float64
	}{
		{"Precon Four", true, 1, 12},
		// Two owned of three needed: not fully owned, and the count says so.
		{"Plains", false, 2, 0.10},
	} {
		c := by[tc.name]
		if c == nil {
			t.Fatalf("%s is not in the deck", tc.name)
		}
		if c.GetOwned() != tc.owned || c.GetOwnedCount() != tc.count || c.GetPriceUsd() != tc.price {
			t.Errorf("%s: owned %v/%d at $%.2f, want %v/%d at $%.2f",
				tc.name, c.GetOwned(), c.GetOwnedCount(), c.GetPriceUsd(), tc.owned, tc.count, tc.price)
		}
	}
	// The owned precon card costs nothing to buy. One Plains does, and so
	// does the commander until the collection holds it.
	if got := BuyCost(deck); got != 0.10 {
		t.Errorf("BuyCost = %.2f, want 0.10 without a card index", got)
	}
	cards := cardMap{"o-c": commander}
	if got := BuyCostWith(deck, cards, nil); math.Abs(got-30.10) > 1e-9 {
		t.Errorf("BuyCostWith = %.2f, want 30.10 with the commander unowned", got)
	}
	if got := BuyCostWith(deck, cards, map[string]int32{"o-c": 1}); math.Abs(got-0.10) > 1e-9 {
		t.Errorf("BuyCostWith = %.2f, want 0.10 with the commander owned", got)
	}
	// The whole deck counts the commander and the owned precon card.
	if got := DeckCostWith(deck, cards); math.Abs(got-42.30) > 1e-9 {
		t.Errorf("DeckCostWith = %.2f, want 42.30", got)
	}
}

// TestBuyCostSumsCopiesPerOracleId: four main and two sideboard copies
// of a card the user owns four of cost two copies, not none (D-37).
func TestBuyCostSumsCopiesPerOracleId(t *testing.T) {
	deck := &mtgv1.Deck{
		Cards:     []*mtgv1.DeckCard{{OracleId: "o-w", Name: "Ajani's Welcome", Count: 4, OwnedCount: 4, PriceUsd: 0.50}},
		Sideboard: []*mtgv1.DeckCard{{OracleId: "o-w", Name: "Ajani's Welcome", Count: 2, OwnedCount: 4, PriceUsd: 0.50}},
	}
	if got := BuyCost(deck); math.Abs(got-1.00) > 1e-9 {
		t.Errorf("BuyCost = %.2f, want 1.00 for the two copies over the four owned", got)
	}
	if got := DeckCost(deck); math.Abs(got-3.00) > 1e-9 {
		t.Errorf("DeckCost = %.2f, want 3.00", got)
	}
}

// TestOwnedFlagReadsTheWholeList: the owned flag reads the copies of an
// oracle id across every entry, so two entries of one card are owned
// only when the collection covers both (D-37).
func TestOwnedFlagReadsTheWholeList(t *testing.T) {
	got := Normalize(testPool(), []Entry{
		{Name: "Ajani's Welcome", Count: 3, Role: "synergy"},
		{Name: "Ajani's Welcome", Count: 2, Role: "synergy"},
		{Name: "Sol Ring", Count: 1, Role: "ramp"},
	})
	if len(got.Cards) != 3 {
		t.Fatalf("cards = %d, want 3", len(got.Cards))
	}
	for _, c := range got.Cards[:2] {
		if c.GetOwned() || c.GetOwnedCount() != 4 {
			t.Errorf("%s x%d: owned %v/%d, want false/4: five copies against four owned", c.GetName(), c.GetCount(), c.GetOwned(), c.GetOwnedCount())
		}
	}
	if !got.Cards[2].GetOwned() {
		t.Error("one Sol Ring against one owned is not owned")
	}
	// Across the main deck and the sideboard, assemble applies the same
	// rule.
	b, _, _ := testBuilder(t)
	req := testRequest()
	req.OracleCounts = map[string]int32{"o-welcome": 4}
	pass := b.assemble(req, &deckOut{Summary: "s",
		Cards:     []Entry{{Name: "Ajani's Welcome", Count: 4, Role: "synergy"}},
		Sideboard: []Entry{{Name: "Ajani's Welcome", Count: 2, Role: "synergy"}},
	})
	for _, c := range allCards(pass.deck) {
		if c.GetOwned() {
			t.Errorf("%s x%d in a deck of six copies against four owned is marked owned", c.GetName(), c.GetCount())
		}
	}
}

// TestPreconShareCountsDistinctIds: two entries of one precon card keep
// one name, not two (D-218).
func TestPreconShareCountsDistinctIds(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-p1", Name: "Precon One"}, {OracleId: "o-p2", Name: "Precon Two"},
		{OracleId: "o-p3", Name: "Precon Three"}, {OracleId: "o-p4", Name: "Precon Four"},
	}, nil)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: pool, Precon: "Test Precon",
		PreconOracleIDs: []string{"o-p1", "o-p2", "o-p3", "o-p4"}}
	// Four names need four kept. Three distinct names and a duplicate
	// entry is three, so the share is short.
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-p1", Name: "Precon One", Count: 1}, {OracleId: "o-p1", Name: "Precon One", Count: 1},
		{OracleId: "o-p2", Name: "Precon Two", Count: 1}, {OracleId: "o-p3", Name: "Precon Three", Count: 1},
		{OracleId: "o-x", Name: "Outsider", Count: 1},
	}}
	if !checkPreconShare(deck, req, nil) {
		t.Error("a duplicate entry counted as a second precon name")
	}
}

// TestSwapBackKeepsTheCommanderColors: a precon card outside the chosen
// commander's color identity never goes back (CR 903.5c).
func TestSwapBackKeepsTheCommanderColors(t *testing.T) {
	W, G := mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_G
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-c", Name: "White Commander", ColorIdentity: []mtgv1.Color{W}},
		{OracleId: "o-p1", Name: "Precon One", ColorIdentity: []mtgv1.Color{W}},
		{OracleId: "o-p2", Name: "Precon Two", ColorIdentity: []mtgv1.Color{W}},
		{OracleId: "o-p3", Name: "Precon Three", ColorIdentity: []mtgv1.Color{W}},
		{OracleId: "o-green", Name: "Green Precon Card", ColorIdentity: []mtgv1.Color{G}},
		{OracleId: "o-p5", Name: "Precon Five", ColorIdentity: nil},
		{OracleId: "o-x", Name: "Outsider"}, {OracleId: "o-y", Name: "Outsider Two"},
	}, nil)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: pool, Precon: "Test Precon",
		PreconOracleIDs: []string{"o-p1", "o-p2", "o-p3", "o-green", "o-p5"}}
	// Five names need five kept. Three are held, and the green card is
	// first in precon order, so a color-blind swap would put it back.
	deck := &mtgv1.Deck{CommanderOracleIds: []string{"o-c"}, Cards: []*mtgv1.DeckCard{
		{OracleId: "o-p1", Name: "Precon One", Count: 1}, {OracleId: "o-p2", Name: "Precon Two", Count: 1},
		{OracleId: "o-p3", Name: "Precon Three", Count: 1},
		{OracleId: "o-x", Name: "Outsider", Count: 1}, {OracleId: "o-y", Name: "Outsider Two", Count: 1},
	}}
	if got := swapBackPrecon(deck, req, nil); got != 1 {
		t.Errorf("swapped = %d, want 1: the colorless card and not the green one", got)
	}
	for _, c := range deck.GetCards() {
		if c.GetOracleId() == "o-green" {
			t.Error("an off-color precon card went back into the deck")
		}
	}
}

// TestPreconNonbasicsReadsTheCardIndex: a precon basic the pool does not
// hold, Wastes or a snow basic, is a basic and not a nonbasic (D-218).
func TestPreconNonbasicsReadsTheCardIndex(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{{OracleId: "o-p1", Name: "Precon One"}}, nil)
	req := Request{Pool: pool, PreconOracleIDs: []string{"o-p1", "o-wastes", "o-snow"}}
	cards := cardMap{"o-wastes": basic("o-wastes", "Wastes"), "o-snow": basic("o-snow", "Snow-Covered Forest")}
	cards["o-snow"].Supertypes = []string{"Basic", "Snow"}
	if got := preconNonbasics(req, cards); len(got) != 1 || !got["o-p1"] {
		t.Errorf("nonbasics = %v, want the one real card", got)
	}
	if got := preconNonbasics(req, nil); len(got) != 3 {
		t.Errorf("nonbasics without a card index = %v, want every id the pool can not name", got)
	}
}

// TestFromListOwnedReadsTheAlwaysCards covers the pool side of G-3. The
// always cards are the commanders, the basics, the locked cards, and the
// precon, and they read their owned count from the collection.
func TestFromListOwnedReadsTheAlwaysCards(t *testing.T) {
	always := []*mtgv1.Card{card("o-karlov", "Karlov of the Ghost Council"), card("o-p1", "Precon One")}
	owned := map[string]int32{"o-karlov": 1, "o-p1": 3}
	list := &candidates.List{Candidates: []candidates.Candidate{{Card: card("o-w", "Ajani's Welcome"), Owned: 2}}}
	for _, tc := range []struct {
		name string
		pool *Pool
		want map[string]int32
	}{
		{"with the collection", FromListOwned(list, always, owned, false),
			map[string]int32{"o-karlov": 1, "o-p1": 3, "o-w": 2}},
		{"without", FromList(list, always, false),
			map[string]int32{"o-karlov": 0, "o-p1": 0, "o-w": 2}},
	} {
		for id, n := range tc.want {
			if got := tc.pool.OwnedCount(id); got != n {
				t.Errorf("%s: owned %s = %d, want %d", tc.name, id, got, n)
			}
		}
		if _, ok := tc.pool.ByOracleID("o-p1"); !ok {
			t.Errorf("%s: ByOracleID lost the precon card", tc.name)
		}
	}
}

// TestFindingsDescribeTheFinalDeck is D-250. The builder puts the precon
// cards back before the engine runs, so every finding describes the deck
// the user gets.
func TestFindingsDescribeTheFinalDeck(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{
		{OracleId: "o-p1", Name: "Precon One"}, {OracleId: "o-p2", Name: "Precon Two"},
		{OracleId: "o-p3", Name: "Precon Three"}, {OracleId: "o-p4", Name: "Precon Four"},
		{OracleId: "o-x", Name: "Outsider"},
	}, nil)
	// Three of four and an outsider: one short, and the builder swaps.
	three := deckOut{Summary: "s", Cards: []Entry{
		{Name: "Precon One", Count: 1, Role: "synergy"}, {Name: "Precon Two", Count: 1, Role: "synergy"},
		{Name: "Precon Three", Count: 1, Role: "synergy"}, {Name: "Outsider", Count: 1, Role: "synergy"},
	}}
	b, _, sc := testBuilder(t, step(t, three), step(t, three))
	req := testRequest()
	req.Pool, req.Precon, req.PreconOracleIDs = pool, "Test Precon", []string{"o-p1", "o-p2", "o-p3", "o-p4"}
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	codes := map[string]int{}
	for _, f := range got.Deck.GetValidation().GetFindings() {
		codes[f.GetCode()]++
	}
	if codes[CodePreconCardsRestored] != 1 || codes[CodePreconShare] != 0 {
		t.Errorf("findings = %v, want one restore and no share warning", codes)
	}
	// The swap met the share, so no repair turn ran for it. The size
	// block still buys one, and the engine's findings name the swapped
	// deck: the outsider is gone.
	if strings.Contains(sc.Calls[1].Input, "Outsider is") {
		t.Error("a finding names a card the swap removed")
	}
	for _, c := range got.Deck.GetCards() {
		if c.GetName() == "Outsider" {
			t.Error("the outsider survived the swap")
		}
	}
}

// TestCommanderSentenceIsCommanderOnly is D-233. Only a Commander session
// hears "the commander is X".
func TestCommanderSentenceIsCommanderOnly(t *testing.T) {
	b, _, _ := testBuilder(t)
	for _, tc := range []struct {
		format mtgv1.FormatId
		want   bool
	}{
		{mtgv1.FormatId_FORMAT_ID_COMMANDER, true},
		{mtgv1.FormatId_FORMAT_ID_MODERN, false},
		{mtgv1.FormatId_FORMAT_ID_STANDARD, false},
	} {
		req := testRequest()
		req.Format, req.Commanders = tc.format, []string{"o-karlov"}
		got := strings.Contains(b.input(req, nil, nil), "The commander is Karlov")
		if got != tc.want {
			t.Errorf("%s: names the commander = %v, want %v", tc.format, got, tc.want)
		}
	}
}
