package profile

import (
	"math"
	"strings"
	"testing"
)

// TestColorSourcesReadsTheSameNumbersAsTheFeature holds the export of
// F-102 to the feature it came out of: the worst ratio of ColorSources is
// the value of color_sources, and each color carries the sources and the
// need the note names.
func TestColorSourcesReadsTheSameNumbersAsTheFeature(t *testing.T) {
	p := newProfiler(t, nil)
	deck := deckOf(3, shaped()...)
	f := feature(t, p.Measure(deck, source(testCards)), KeyColorSources)
	rows := ColorSources(deck, source(testCards))
	if len(rows) == 0 {
		t.Fatal("ColorSources read no color of the shaped deck")
	}
	worst := math.Inf(1)
	for _, r := range rows {
		worst = math.Min(worst, r.Ratio())
		want := colorLetters[r.Color] + " " + num(r.Have) + " of "
		if !strings.Contains(f.GetNote(), want) {
			t.Errorf("the note %q holds no %q", f.GetNote(), want)
		}
	}
	if got := math.Round(worst*100) / 100; got != f.GetValue() {
		t.Errorf("the worst ratio of ColorSources reads %v, and color_sources reads %v", got, f.GetValue())
	}
	if (ColorSource{Have: 3}).Ratio() != 1 {
		t.Error("a color no spell asks for reads a ratio other than 1")
	}
}
