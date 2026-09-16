package generate

import (
	"fmt"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The PR-52 tests of the mana pass (F-139, F-146, D-720, D-733, D-734).
// Each land text is the Oracle text of the card.

// swapCards are the cards of a blue-black deck and the lands of its pool.
type swapCards struct {
	island, swamp, blue, black        *mtgv1.Card
	grave, delta, hollow, isle, flats *mtgv1.Card
}

func newSwapCards() swapCards {
	U, B := mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B
	land := func(oid, name, text string, subtypes []string, produced ...mtgv1.Color) *mtgv1.Card {
		c := manaCard(oid, name, 0, []string{"Land"}, text, produced...)
		c.Subtypes = subtypes
		return c
	}
	spell := func(oid, name, cost string, color mtgv1.Color) *mtgv1.Card {
		c := manaCard(oid, name, 2, []string{"Creature"}, "")
		c.ManaCost, c.Colors, c.ColorIdentity = cost, []mtgv1.Color{color}, []mtgv1.Color{color}
		return c
	}
	s := swapCards{
		island: land("o-island", "Island", "({T}: Add {U}.)", []string{"Island"}, U),
		swamp:  land("o-swamp", "Swamp", "({T}: Add {B}.)", []string{"Swamp"}, B),
		blue:   spell("o-blue", "Blue Spell", "{1}{U}", U),
		black:  spell("o-black", "Black Spell", "{1}{B}", B),
		grave: land("o-grave", "Watery Grave", "({T}: Add {U} or {B}.)\nAs this land enters, you may pay 2 life. If you don't, it enters tapped.",
			[]string{"Island", "Swamp"}, U, B),
		delta: land("o-delta", "Polluted Delta", "{T}, Pay 1 life, Sacrifice this land: Search your library for an Island or Swamp card, put it onto the battlefield, then shuffle.", nil),
		hollow: land("o-hollow", "Sunken Hollow", "({T}: Add {U} or {B}.)\nThis land enters tapped unless you control two or more basic lands.",
			[]string{"Island", "Swamp"}, U, B),
		isle: land("o-isle", "Thriving Isle", "This land enters tapped. As it enters, choose a color other than blue.\n{T}: Add {U} or one mana of the chosen color.",
			nil, mtgv1.Color_COLOR_W, U, B, mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G),
		flats: land("o-flats", "Marsh Flats", "{T}, Pay 1 life, Sacrifice this land: Search your library for a Plains or Swamp card, put it onto the battlefield, then shuffle.", nil),
	}
	s.island.Supertypes = []string{"Basic"}
	s.swamp.Supertypes = []string{"Basic"}
	return s
}

func (s swapCards) all() []*mtgv1.Card {
	return []*mtgv1.Card{s.island, s.swamp, s.blue, s.black, s.grave, s.delta, s.hollow, s.isle, s.flats}
}

// deck is 17 Islands, 18 Swamps, and 64 spells of two mana, at a bracket.
func (s swapCards) deck(bracket int32) *mtgv1.Deck {
	entry := func(c *mtgv1.Card, n int32, role mtgv1.CardRole) *mtgv1.DeckCard {
		return &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: n, Role: role}
	}
	return &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:  &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bracket}},
		Cards: []*mtgv1.DeckCard{
			entry(s.island, 17, mtgv1.CardRole_CARD_ROLE_LAND), entry(s.swamp, 18, mtgv1.CardRole_CARD_ROLE_LAND),
			entry(s.blue, 32, mtgv1.CardRole_CARD_ROLE_THREAT), entry(s.black, 32, mtgv1.CardRole_CARD_ROLE_THREAT),
		},
	}
}

// fixed runs the pass over the deck at a bracket, with a pool rule, the
// owned counts, and a budget.
func (s swapCards) fixed(t *testing.T, bracket int32, rule mtgv1.PoolRule, owned map[string]int32, budget float64) *mtgv1.Deck {
	t.Helper()
	src := source{}
	for _, c := range s.all() {
		src[c.GetOracleId()] = c
	}
	b := manaBuilder(t, src)
	deck := s.deck(bracket)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(), PoolRule: rule,
		Pool: NewPool(s.all(), owned), OracleCounts: owned, BudgetUSD: budget}
	b.fixMana(req, deck)
	if n := deckCount(deck); n != 99 {
		t.Errorf("the deck holds %d cards after the pass, want 99", n)
	}
	return deck
}

func ownedAll(cs []*mtgv1.Card) map[string]int32 {
	out := map[string]int32{}
	for _, c := range cs {
		out[c.GetOracleId()] = 40
	}
	return out
}

// TestTheManaPassSwapsABasicLandForABetterLand is F-139, D-720, and
// D-734. A bracket 4 deck of basic lands takes the untapped dual, the
// fetch land, and the check land of its pool. It never takes a tapped
// dual, and never a fetch land that finds one deck color.
func TestTheManaPassSwapsABasicLandForABetterLand(t *testing.T) {
	s := newSwapCards()
	deck := s.fixed(t, 4, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, ownedAll(s.all()), 0)
	for _, c := range []*mtgv1.Card{s.grave, s.delta, s.hollow} {
		if countOf(deck, c.GetOracleId()) != 1 {
			t.Errorf("the pass did not trade a basic land for %s", c.GetName())
		}
	}
	for _, c := range []*mtgv1.Card{s.isle, s.flats} {
		if countOf(deck, c.GetOracleId()) != 0 {
			t.Errorf("the pass traded a basic land for %s, a worse land", c.GetName())
		}
	}
	for _, dc := range deck.GetCards() {
		if dc.GetOracleId() == s.grave.GetOracleId() && (dc.GetReason() != swapReason || dc.GetOwnedCount() != 40 || !dc.GetOwned()) {
			t.Errorf("Watery Grave reads reason %q, %d owned, and owned %v", dc.GetReason(), dc.GetOwnedCount(), dc.GetOwned())
		}
	}
}

// TestTheSwapWaitsForBracketFour is D-734. The land rank runs at every
// bracket, and the swap at brackets 4 and 5 alone.
func TestTheSwapWaitsForBracketFour(t *testing.T) {
	s := newSwapCards()
	deck := s.fixed(t, 3, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, ownedAll(s.all()), 0)
	for _, c := range []*mtgv1.Card{s.grave, s.delta, s.hollow} {
		if countOf(deck, c.GetOracleId()) != 0 {
			t.Errorf("the pass traded a basic land for %s at bracket 3", c.GetName())
		}
	}
}

// TestTheSwapKeepsThePoolRuleAndTheBudget: under owned-first a swap adds
// a land the reader owns alone, and under any card no swap takes the cards
// to buy over the budget.
func TestTheSwapKeepsThePoolRuleAndTheBudget(t *testing.T) {
	s := newSwapCards()
	owned := map[string]int32{"o-island": 40, "o-swamp": 40, "o-blue": 32, "o-black": 32, "o-hollow": 1}
	deck := s.fixed(t, 5, mtgv1.PoolRule_POOL_RULE_OWNED_FIRST, owned, 0)
	if countOf(deck, s.grave.GetOracleId()) != 0 || countOf(deck, s.delta.GetOracleId()) != 0 {
		t.Error("owned-first traded a basic land for a land the reader does not own")
	}
	if countOf(deck, s.hollow.GetOracleId()) != 1 {
		t.Error("owned-first did not trade a basic land for the owned Sunken Hollow")
	}

	s.grave.PriceUsd, s.delta.PriceUsd, s.hollow.PriceUsd = 20, 30, 0.5
	delete(owned, "o-hollow")
	deck = s.fixed(t, 5, mtgv1.PoolRule_POOL_RULE_ANY_CARD, owned, 1)
	if countOf(deck, s.grave.GetOracleId()) != 0 || countOf(deck, s.delta.GetOracleId()) != 0 {
		t.Error("a swap took the cards to buy over a budget of $1")
	}
	if countOf(deck, s.hollow.GetOracleId()) != 1 {
		t.Error("the swap skipped Sunken Hollow, and it fits a budget of $1")
	}
}

// TestTheUntappedLandStepRanksThePool is F-146. The step read the first
// eight untapped lands of the pool, and the pool lists its names by the
// alphabet. So nine lands whose names come first kept Watery Grave out.
func TestTheUntappedLandStepRanksThePool(t *testing.T) {
	s := newSwapCards()
	gate := manaCard("o-gate", "Dimir Guildgate", 0, []string{"Land"}, "This land enters tapped.\n{T}: Add {U} or {B}.",
		mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B)
	pool := append(s.all(), gate)
	for i := range 9 {
		pool = append(pool, manaCard(fmt.Sprintf("o-hold-%d", i), fmt.Sprintf("Aether Hold %d", i), 0, []string{"Land"},
			"{T}: Add {U}.", mtgv1.Color_COLOR_U))
	}
	src := source{}
	for _, c := range pool {
		src[c.GetOracleId()] = c
	}
	b := manaBuilder(t, src)
	deck := s.deck(4)
	deck.Cards = append(deck.Cards, &mtgv1.DeckCard{OracleId: gate.GetOracleId(), Name: gate.GetName(), Count: 1, Role: mtgv1.CardRole_CARD_ROLE_LAND})
	_, untapped := b.landsOf(deck, Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(), Pool: NewPool(pool, nil)})
	var got []string
	for _, c := range untapped {
		got = append(got, c.GetName())
	}
	if len(got) == 0 || got[0] != "Watery Grave" {
		t.Errorf("the step reads %v, want Watery Grave first", got)
	}
}

// TestTheSwapDropsTheBasicOfTheColorWithMostToSpare is PR-52. When every
// color meets its need, the capped ratios tie. The swap then drops a basic
// land of the color with the most sources to spare, and not the first
// basic land by name.
func TestTheSwapDropsTheBasicOfTheColorWithMostToSpare(t *testing.T) {
	s := newSwapCards()
	// In 99 cards a {1}{U} spell needs 19 blue sources, and a {3}{B} spell
	// needs 16 black sources (Karsten 2022). So 19 Islands meet the blue
	// need exactly, and 18 Swamps leave 2 black sources to spare.
	s.black.ManaCost, s.black.ManaValue = "{3}{B}", 4
	src := source{}
	for _, c := range s.all() {
		src[c.GetOracleId()] = c
	}
	b := manaBuilder(t, src)
	deck := s.deck(4)
	for _, dc := range deck.GetCards() {
		switch dc.GetOracleId() {
		case s.island.GetOracleId():
			dc.Count = 19
		case s.swamp.GetOracleId():
			dc.Count = 18
		case s.blue.GetOracleId():
			dc.Count = 30
		}
	}
	owned := ownedAll(s.all())
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(), PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		Pool: NewPool(s.all(), owned), OracleCounts: owned}
	if !b.swapOne(req, deck) {
		t.Fatal("the swap made no step")
	}
	if countOf(deck, s.grave.GetOracleId()) != 1 {
		t.Fatal("the first swap did not take Watery Grave")
	}
	if islands, swamps := countOf(deck, s.island.GetOracleId()), countOf(deck, s.swamp.GetOracleId()); islands != 19 || swamps != 17 {
		t.Errorf("the swap left %d Islands and %d Swamps, want 19 and 17", islands, swamps)
	}
}

// TestTheSwapKeepsABasicLandForEachLandThatReadsOne is PR-52. Fabled
// Passage finds a basic land card, and Sunken Hollow enters untapped with
// two basic lands. A swap never leaves fewer basic lands than the lands
// that read them.
func TestTheSwapKeepsABasicLandForEachLandThatReadsOne(t *testing.T) {
	s := newSwapCards()
	passage := manaCard("o-passage", "Fabled Passage", 0, []string{"Land"},
		"{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle. Then if you control four or more lands, untap that land.")
	dual := manaCard("o-dual", "Open Channel", 0, []string{"Land"}, "{T}: Add {U} or {B}.", mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B)
	all := append(s.all(), passage, dual)
	src := source{}
	for _, c := range all {
		src[c.GetOracleId()] = c
	}
	b := manaBuilder(t, src)
	deck := s.deck(4)
	land, threat := mtgv1.CardRole_CARD_ROLE_LAND, mtgv1.CardRole_CARD_ROLE_THREAT
	deck.Cards = []*mtgv1.DeckCard{
		{OracleId: s.island.GetOracleId(), Name: s.island.GetName(), Count: 2, Role: land},
		{OracleId: s.swamp.GetOracleId(), Name: s.swamp.GetName(), Count: 2, Role: land},
		{OracleId: s.hollow.GetOracleId(), Name: s.hollow.GetName(), Count: 3, Role: land},
		{OracleId: dual.GetOracleId(), Name: dual.GetName(), Count: 28, Role: land},
		{OracleId: s.blue.GetOracleId(), Name: s.blue.GetName(), Count: 32, Role: threat},
		{OracleId: s.black.GetOracleId(), Name: s.black.GetName(), Count: 32, Role: threat},
	}
	owned := ownedAll(all)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(), PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		Pool: NewPool([]*mtgv1.Card{s.island, s.swamp, s.blue, s.black, s.hollow, dual, passage}, owned), OracleCounts: owned}
	if b.swapOne(req, deck) || countOf(deck, passage.GetOracleId()) != 0 {
		t.Error("the swap took Fabled Passage and left 3 basic lands for 4 lands that read them")
	}
}
