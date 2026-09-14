package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestPinnedPowerLandSkipsTheLandCap is the review of PR-45b. A Game
// Changer land counts toward the Game Changer floor, so the pin reads it
// as a power card. The land cap once read no pin, and a list that caps its
// lands at one dropped the pinned land for a more played land.
func TestPinnedPowerLandSkipsTheLandCap(t *testing.T) {
	b, _ := New()
	idx := fixture(t, append(pinCards(), tc{id: "tomb", name: "Ancient Tomb", typeLine: "Land",
		text: "{T}: Add {C}{C}. Ancient Tomb deals 2 damage to you.", rank: 60, gameChanger: true}))
	req := pinRequest()
	req.MetaBoost = fixedRates(map[string]float64{"tomb": 0.9})
	req.Limits = Limits{PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_LAND: 1}}
	for _, pin := range []bool{false, true} {
		req.PowerRate = PowerRate{Weight: 0.1, Keep: 0.5, Pin: pin}
		list, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		tomb, hasTomb := find(list.Candidates, "Ancient Tomb")
		_, hasTower := find(list.Candidates, "Command Tower")
		if hasTomb != pin || tomb.Pinned != pin || !hasTower {
			t.Errorf("pin %v: Ancient Tomb listed %v and pinned %v, Command Tower listed %v", pin, hasTomb, tomb.Pinned, hasTower)
		}
	}
}
