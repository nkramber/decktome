package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestCommanderRateOrdersTheLandCap is D-843 (F-168). The mana half of the
// land cap read the land rank at bracket 5 too, so a shock land and a fetch
// land that no list of the commander plays took the place of lands that
// most of its lists play. A bracket 5 request with a commander rate fills
// the whole land cap in score order now. Below bracket 5 the cap keeps the
// order of D-733.
func TestCommanderRateOrdersTheLandCap(t *testing.T) {
	b, _ := New()
	idx := fixture(t, landRankCards())
	rate := fixedRates(map[string]float64{"path": 0.7, "orchard": 0.64, "wilds": 0.4})
	build := func(bracket int32, pool mtgv1.PoolRule, lim Limits) []Candidate {
		t.Helper()
		req := Request{Format: cmdr, Colors: blueBlack, Bracket: bracket, PoolRule: pool, CommanderRate: rate, Limits: lim}
		if pool == mtgv1.PoolRule_POOL_RULE_OWNED_FIRST {
			req.Owned = map[string]int32{"grave": 1}
		}
		list, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		return list.Candidates
	}
	anyCard, first := mtgv1.PoolRule_POOL_RULE_ANY_CARD, mtgv1.PoolRule_POOL_RULE_OWNED_FIRST

	sameNames(t, "a bracket 5 land cap of 3", build(5, anyCard, landCap(3)), "Path of Ancestry", "Exotic Orchard", "Evolving Wilds")
	// The fill of an owned-first list caps the lands it adds in the same
	// order. The collection owns Watery Grave, so the fill caps two lands.
	sameNames(t, "an owned-first land cap of 3", build(5, first, landCap(3)), "Watery Grave", "Path of Ancestry", "Exotic Orchard")
	// The total cut keeps the mana half first, in the same order.
	cut := landCap(3)
	cut.Total = 1
	sameNames(t, "a total of 1", build(5, anyCard, cut), "Path of Ancestry")
	// Bracket 4 reads no commander rate, so its cap keeps the class order.
	sameNames(t, "a bracket 4 land cap of 3", build(4, anyCard, landCap(3)), "Watery Grave", "Polluted Delta", "Sunken Hollow")
}
