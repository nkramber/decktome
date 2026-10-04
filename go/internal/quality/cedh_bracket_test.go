package quality

import (
	"strings"
	"testing"
)

// TestCEDHSignalReadsBracketFiveAlone reads a commander that places in
// no cEDH event. A deck of bracket 3 or 4 reads the mean of the model,
// so the signal names no reason. A deck of bracket 5 and a deck with no
// bracket keep the signal (D-1123).
func TestCEDHSignalReadsBracketFiveAlone(t *testing.T) {
	w := newCommanderWorld(t)
	fm := &FormatModel{
		Keys:    []string{KeyCEDHSignal},
		Means:   []float64{0.2},
		Stds:    []float64{0.1},
		Weights: []float64{0.5},
	}
	const words = "the commander does not place in cEDH events"
	read := func(bracket int32) []string {
		d := w.landDeck(36)
		d.Power = nil
		if bracket > 0 {
			d.Power = bracketPower(bracket)
		}
		in := w.input(d)
		return reasons(fm, fm.vector(Features(in, fm)), ruleChecks(in), in.Profile)
	}
	for _, b := range []int32{3, 4} {
		if got := read(b); strings.Contains(strings.Join(got, "|"), words) {
			t.Errorf("bracket %d names %q", b, got)
		}
	}
	for _, b := range []int32{5, 0} {
		if got := read(b); !strings.Contains(strings.Join(got, "|"), words) {
			t.Errorf("bracket %d names %q, want the cEDH reason", b, got)
		}
	}
}
