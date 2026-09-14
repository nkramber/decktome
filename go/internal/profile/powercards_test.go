package profile

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestPowerCardsReadTheProfileRules is the counter of the gate dry run
// (D-706). Fast mana is a nonland, noncreature mana source of mana value
// one or less, so a mana creature does not count, and a Game Changer
// counts by its flag.
func TestPowerCardsReadTheProfileRules(t *testing.T) {
	sol := &mtgv1.Card{OracleId: "o-sol", Name: "Sol Ring", CardTypes: []string{"Artifact"}, ManaValue: 1,
		ProducedMana: []mtgv1.Color{mtgv1.Color_COLOR_C}}
	elves := &mtgv1.Card{OracleId: "o-elves", Name: "Llanowar Elves", CardTypes: []string{"Creature"}, ManaValue: 1,
		ProducedMana: []mtgv1.Color{mtgv1.Color_COLOR_G}}
	study := &mtgv1.Card{OracleId: "o-study", Name: "Rhystic Study", CardTypes: []string{"Enchantment"}, ManaValue: 3,
		GameChanger: true}
	tutors, fast, changers := PowerCards([]*mtgv1.Card{sol, elves, study}, nil)
	if tutors != 0 || fast != 1 || changers != 1 {
		t.Errorf("tutors %d, fast mana %d, Game Changers %d; want 0, 1, 1", tutors, fast, changers)
	}
}
