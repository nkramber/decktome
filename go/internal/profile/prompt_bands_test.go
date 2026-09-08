package profile

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestEveryBandReachesTheModel is the gate of F-78, Part 1. A feature
// the profile can call off band, and no prompt names, is a number the
// model answers blind. Before this the prompt named five of the fifteen
// features `bands.json` holds. 12 of the 26 off-band findings of deck
// gate run 16 and bracket gate run 1 were on `mana_turn_four` and
// `color_sources`, and no prompt ever stated either one.
//
// A feature reaches the model in one of two places: the deck shape block
// writes it, or the job block sends it as a count with its band.
func TestEveryBandReachesTheModel(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	// jobKeys are the features the job target block carries. generate
	// owns that block, and TestJobTargetsCarryTheirBand there proves the
	// band goes with the count.
	jobKeys := map[string]bool{
		KeyLand: true, KeyRamp: true, KeyDraw: true,
		KeyRemoval: true, KeyWipe: true, KeyInteraction: true,
	}
	// The rules engine reports the Game Changer limit as a block, so the
	// profile marks the band and raises no off-band finding for it.
	skip := map[string]bool{KeyGameChanger: true}

	for _, tc := range []struct {
		name   string
		format mtgv1.FormatId
		power  *mtgv1.PowerLevel
	}{
		{"bracket 1", mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(1)},
		{"bracket 3", mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(3)},
		{"bracket 4", mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(4)},
		{"bracket 5", mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(5)},
		{"60-card casual", mtgv1.FormatId_FORMAT_ID_MODERN, step(mtgv1.SixtyStep_SIXTY_STEP_CASUAL)},
		{"60-card tournament", mtgv1.FormatId_FORMAT_ID_MODERN, step(mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			table, _ := b.For(tc.format, tc.power)
			lines := strings.Join(b.Lines(tc.format, tc.power), "\n")
			for key := range table {
				switch {
				case skip[key]:
					continue
				case jobKeys[key]:
					if b.Words(tc.format, tc.power, key) == "" {
						t.Errorf("the job block can not state the band of %q", key)
					}
					continue
				}
				if !strings.Contains(lines, promptWords[key]) && derivedLine(key, table[key], tc.format) == "" {
					t.Errorf("no prompt line names %q, and the profile calls it off band", key)
				}
			}
		})
	}
}

// The four derived lines name the number the check reads, so a model
// that follows them aims at the band and not past it.
func TestDerivedLinesNameTheirNumber(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Join(b.Lines(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(4)), "\n")
	for _, want := range []string{
		"4.8 mana on turn four", // mana_turn_four, bracket 4
		"70 percent",            // hands_two_to_four_lands
		"0.5 of a turn",         // commander_turn_over_mv
		"about 19 sources",      // color_sources, the 99-card table
		"about 26",              // the two-pip row of that table
	} {
		if !strings.Contains(lines, want) {
			t.Errorf("the deck shape block does not name %q:\n%s", want, lines)
		}
	}
	// A 60-card deck reads its own source counts, and no commander line.
	sixty := strings.Join(b.Lines(mtgv1.FormatId_FORMAT_ID_MODERN, step(mtgv1.SixtyStep_SIXTY_STEP_FNM)), "\n")
	if !strings.Contains(sixty, "about 13 sources") {
		t.Errorf("the 60-card block reads the 99-card source table:\n%s", sixty)
	}
	if strings.Contains(sixty, "the commander comes down") {
		t.Errorf("a 60-card block names the commander line:\n%s", sixty)
	}
}

func step(s mtgv1.SixtyStep) *mtgv1.PowerLevel {
	return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: s}}
}
