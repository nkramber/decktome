package quality

import (
	"strings"
	"testing"
)

// TestCEDHSignalReadsBracketFiveAlone reads a commander that places in
// no cEDH event and leads few decks. A deck of bracket 3 or 4 reads the
// mean of the model for both, so neither names a reason. A deck of
// bracket 5 and a deck with no bracket keep both signals (D-1123).
func TestCEDHSignalReadsBracketFiveAlone(t *testing.T) {
	w := newCommanderWorld(t)
	fm := &FormatModel{
		Keys:    []string{KeyCEDHSignal, KeyCommanderDecks},
		Means:   []float64{0.2, 6},
		Stds:    []float64{0.1, 4},
		Weights: []float64{0.5, -0.4},
	}
	words := []string{"the commander does not place in cEDH events", "few decks lead with the commander"}
	names := func(got []string, w string) bool { return strings.Contains(strings.Join(got, "|"), w) }
	read := func(bracket int32) []string {
		d := w.landDeck(36)
		d.Power = nil
		if bracket > 0 {
			d.Power = bracketPower(bracket)
		}
		in := w.input(d)
		return reasons(fm, fm.vector(Features(in, fm)), ruleChecks(in), in.Profile)
	}
	for _, want := range words {
		for _, b := range []int32{3, 4} {
			if got := read(b); names(got, want) {
				t.Errorf("bracket %d names %q", b, got)
			}
		}
		for _, b := range []int32{5, 0} {
			if got := read(b); !names(got, want) {
				t.Errorf("bracket %d names %q, want %q", b, got, want)
			}
		}
	}
}
