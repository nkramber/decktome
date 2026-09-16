package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The PR-52 tests of the land cap (F-139, D-733, D-734). Each land text
// and rank is the Oracle text and the EDHREC rank of the card.

var (
	blueBlack  = []mtgv1.Color{mtgv1.Color_COLOR_U, B}
	fiveColors = []mtgv1.Color{W, mtgv1.Color_COLOR_U, B, mtgv1.Color_COLOR_R, G}
)

func landRankCards() []tc {
	return []tc{
		{id: "grave", name: "Watery Grave", typeLine: "Land — Island Swamp", subtypes: []string{"Island", "Swamp"}, produced: blueBlack, rank: 50,
			text: "({T}: Add {U} or {B}.)\nAs this land enters, you may pay 2 life. If you don't, it enters tapped."},
		{id: "delta", name: "Polluted Delta", typeLine: "Land", rank: 36,
			text: "{T}, Pay 1 life, Sacrifice this land: Search your library for an Island or Swamp card, put it onto the battlefield, then shuffle."},
		{id: "hollow", name: "Sunken Hollow", typeLine: "Land — Island Swamp", subtypes: []string{"Island", "Swamp"}, produced: blueBlack, rank: 86,
			text: "({T}: Add {U} or {B}.)\nThis land enters tapped unless you control two or more basic lands."},
		{id: "orchard", name: "Exotic Orchard", typeLine: "Land", produced: fiveColors, rank: 9,
			text: "{T}: Add one mana of any color that a land an opponent controls could produce."},
		{id: "path", name: "Path of Ancestry", typeLine: "Land", produced: fiveColors, rank: 14,
			text: "This land enters tapped.\n{T}: Add one mana of any color in your commander's color identity. When that mana is spent to cast a creature spell that shares a creature type with your commander, scry 1."},
		{id: "wilds", name: "Evolving Wilds", typeLine: "Land", rank: 19,
			text: "{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle."},
		{id: "flats", name: "Marsh Flats", typeLine: "Land", rank: 52,
			text: "{T}, Pay 1 life, Sacrifice this land: Search your library for a Plains or Swamp card, put it onto the battlefield, then shuffle."},
	}
}

func landCap(n int) Limits {
	return Limits{PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_LAND: n}}
}

func sameNames(t *testing.T, what string, got []Candidate, want ...string) {
	t.Helper()
	have := map[string]bool{}
	for _, c := range got {
		have[c.Card.GetName()] = true
	}
	if len(got) != len(want) {
		t.Errorf("%s: %v, want %v", what, names(got), want)
		return
	}
	for _, name := range want {
		if !have[name] {
			t.Errorf("%s: %v, want %v", what, names(got), want)
			return
		}
	}
}

// TestCommanderLandsRankByClass is D-733 and D-734. The land cap ranked
// the deck colors a land makes, then play. So a fetch land ranked under
// every dual, and Exotic Orchard and Path of Ancestry, the most played,
// ranked over Sunken Hollow. A Commander land reads its class now, at
// every bracket, and a Standard land keeps the order of D-450.
func TestCommanderLandsRankByClass(t *testing.T) {
	b, _ := New()
	idx := fixture(t, landRankCards())
	for _, bracket := range []int32{0, 2, 5} {
		list, err := b.Build(idx, Request{Format: cmdr, Colors: blueBlack, Bracket: bracket, Limits: landCap(3)})
		if err != nil {
			t.Fatal(err)
		}
		sameNames(t, "a Commander land cap of 3", list.Candidates, "Watery Grave", "Polluted Delta", "Sunken Hollow")
	}
	std, err := b.Build(idx, Request{Format: mtgv1.FormatId_FORMAT_ID_STANDARD, Colors: blueBlack, Limits: landCap(3)})
	if err != nil {
		t.Fatal(err)
	}
	sameNames(t, "a Standard land cap of 3", std.Candidates, "Exotic Orchard", "Path of Ancestry", "Watery Grave")
}

// TestAnUnownedLandOverTheBudgetRanksLast is D-736. The rank reads the
// function of a land alone, so a request with any card and a budget once
// listed a land the reader can never buy.
func TestAnUnownedLandOverTheBudgetRanksLast(t *testing.T) {
	b, _ := New()
	idx := fixture(t, landRankCards())
	for name, price := range map[string]float64{"Watery Grave": 300, "Polluted Delta": 30} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("no card named %q", name)
		}
		c.PriceUsd = price
	}
	req := Request{Format: cmdr, Colors: blueBlack, Limits: landCap(2), BudgetUSD: 100}
	list, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	sameNames(t, "a land cap of 2 under a budget of $100", list.Candidates, "Polluted Delta", "Sunken Hollow")
	// The same request with no budget takes the dearest land first.
	req.BudgetUSD = 0
	if list, err = b.Build(idx, req); err != nil {
		t.Fatal(err)
	}
	sameNames(t, "a land cap of 2 with no budget", list.Candidates, "Watery Grave", "Polluted Delta")
	// A land the reader owns keeps its class, whatever it costs.
	req.BudgetUSD, req.PoolRule, req.Owned = 100, mtgv1.PoolRule_POOL_RULE_ANY_CARD, map[string]int32{"grave": 1}
	if list, err = b.Build(idx, req); err != nil {
		t.Fatal(err)
	}
	sameNames(t, "a land cap of 2 with the dear land owned", list.Candidates, "Watery Grave", "Polluted Delta")
}

// TestAnOwnedLandLeadsItsClass is D-733. An owned land comes before an
// unowned land of its class, and never before a better class.
func TestAnOwnedLandLeadsItsClass(t *testing.T) {
	b, _ := New()
	idx := fixture(t, append(landRankCards(), tc{id: "sea", name: "Underground Sea", typeLine: "Land — Island Swamp",
		subtypes: []string{"Island", "Swamp"}, produced: blueBlack, rank: 298, text: "({T}: Add {U} or {B}.)"}))
	req := Request{Format: cmdr, Colors: blueBlack, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Owned: map[string]int32{"sea": 1, "path": 1}}
	req.Limits = landCap(1)
	list, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	sameNames(t, "a land cap of 1 with Underground Sea owned", list.Candidates, "Underground Sea")
	req.Limits = landCap(3)
	if list, err = b.Build(idx, req); err != nil {
		t.Fatal(err)
	}
	sameNames(t, "a land cap of 3 with Path of Ancestry owned", list.Candidates, "Underground Sea", "Watery Grave", "Polluted Delta")
}
