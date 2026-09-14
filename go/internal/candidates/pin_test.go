package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestPinnedPowerCardSkipsItsRoleCap is D-710. A bracket 5 list caps ramp
// at one card, and an on-theme ramp card takes that place. Sol Ring is
// fast mana at the keep rate: with no pin the cap drops it, and with a pin
// it stays beside the on-theme card. Under a total of one, the pinned card
// holds its place.
func TestPinnedPowerCardSkipsItsRoleCap(t *testing.T) {
	b, _ := New()
	idx := fixture(t, append(powerCards(), tc{id: "stone", name: "Lifegain Stone", typeLine: "Artifact",
		text: "Whenever you gain life, add {W}.", identity: []mtgv1.Color{W}, mv: 2, rank: 5000, tags: []string{"ramp", "lifegain"}}))
	req := Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain", Bracket: 5,
		MetaBoost: fixedRates(map[string]float64{"solring": 0.9}),
		Limits:    Limits{PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_RAMP: 1}}}
	for _, pin := range []bool{false, true} {
		req.PowerRate = PowerRate{Weight: 0.1, Keep: 0.5, Pin: pin}
		list, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		sol, hasSol := find(list.Candidates, "Sol Ring")
		_, hasStone := find(list.Candidates, "Lifegain Stone")
		if hasSol != pin || sol.Pinned != pin || !hasStone {
			t.Errorf("pin %v: Sol Ring listed %v and pinned %v, Lifegain Stone listed %v", pin, hasSol, sol.Pinned, hasStone)
		}
	}
	req.Limits = Limits{Total: 1}
	req.PowerRate = PowerRate{Weight: 0.1, Keep: 0.5, Pin: true}
	list, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if got := names(list.Candidates); len(got) != 1 || got[0] != "Sol Ring" {
		t.Errorf("under a total of one, the list holds %v, want the pinned Sol Ring alone", got)
	}
}
