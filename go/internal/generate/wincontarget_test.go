package generate

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestTheWinconTargetOfEachBracket is D-726. The job targets held no
// wincon, so no prompt ever asked for a card that wins the game. The
// finisher places come out of the synergy pieces, so the deck still
// counts 99 cards.
func TestTheWinconTargetOfEachBracket(t *testing.T) {
	for bracket, want := range map[int32]int{1: 3, 2: 3, 3: 3, 4: 3, 5: 1} {
		power := &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bracket}}
		targets := TargetsFor(mtgv1.FormatId_FORMAT_ID_COMMANDER, power)
		if got := targets["wincon"]; got != want {
			t.Errorf("bracket %d reads a wincon target of %d, want %d", bracket, got, want)
		}
		sum := 0
		for _, n := range targets {
			sum += n
		}
		if sum != 99 {
			t.Errorf("bracket %d asks for %d cards, want 99: %v", bracket, sum, targets)
		}
	}
	// A 60-card format holds no bracket and no finisher target.
	modern := TargetsFor(mtgv1.FormatId_FORMAT_ID_MODERN, &mtgv1.PowerLevel{})
	if _, ok := modern["wincon"]; ok {
		t.Errorf("a Modern deck reads a wincon target: %v", modern)
	}
}
