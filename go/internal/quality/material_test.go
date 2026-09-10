package quality

import (
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/meta"
)

// TestDropImmaterial holds the check of D-652. A synergy copy stays when
// its break lowered the synergy feature by the floor, in standard
// deviations of the real training lists. The real lists and the copies of
// other axes always stay, a break that lowered nothing never does, and a
// negative floor keeps every row.
func TestDropImmaterial(t *testing.T) {
	realList := func(id string, synergy float64, hold bool) sample {
		l := &meta.List{Source: meta.SourceMTGJSON, ID: id, Tier: meta.TierBaseline}
		return sample{r: &Resolved{List: l}, features: map[string]float64{KeySynergy: synergy}, hold: hold}
	}
	copyOf := func(base sample, requested, defect string, synergy float64) sample {
		l := &meta.List{Source: meta.SourceSynthetic, ID: base.r.List.Key() + "/" + requested, Tier: meta.TierBad, Defect: defect}
		return sample{r: &Resolved{List: l}, features: map[string]float64{KeySynergy: synergy}, hold: base.hold}
	}
	// The real training lists read 0 and 2, so the unit is 1. The holdout
	// list takes no part in the unit.
	low, high, held := realList("low", 0, false), realList("high", 2, false), realList("held", 9, true)
	samples := []sample{
		low, high, held,
		copyOf(high, DefectSynergy, DefectSynergy, 1.5),
		copyOf(high, DefectCopies, DefectSynergy, 1.9),
		copyOf(low, DefectSynergy, DefectSynergy, 0),
		copyOf(held, DefectColors, DefectSynergy, 8.8),
		copyOf(held, DefectSynergy, DefectSynergy, 9.2),
		copyOf(low, DefectLands, DefectLands, 0),
	}
	keys := func(ss []sample) string {
		var out []string
		for _, s := range ss {
			out = append(out, s.r.List.Key())
		}
		return strings.Join(out, ",")
	}

	kept, dropped, unit := dropImmaterial(samples, 0.25)
	want := "mtgjson:low,mtgjson:high,mtgjson:held,synthetic:mtgjson:high/synergy,synthetic:mtgjson:low/lands"
	if got := keys(kept); got != want {
		t.Errorf("kept at 0.25 = %s, want %s", got, want)
	}
	if unit != 1 || len(dropped) != 2 || dropped[DefectColors] != 1 || dropped[DefectSynergy] != 1 {
		t.Errorf("unit %.2f, dropped %v, want 1 and one holdout copy each of colors and synergy", unit, dropped)
	}

	// At 0.15 the colors copy of the holdout falls far enough, and the copy
	// that rose still leaves.
	kept, dropped, _ = dropImmaterial(samples, 0.15)
	if len(kept) != 6 || len(dropped) != 1 || dropped[DefectSynergy] != 1 {
		t.Errorf("at 0.15: kept %s, dropped %v", keys(kept), dropped)
	}

	kept, dropped, unit = dropImmaterial(samples, -1)
	if len(kept) != len(samples) || len(dropped) != 0 || unit != 0 {
		t.Errorf("a negative floor kept %d of %d rows and dropped %v", len(kept), len(samples), dropped)
	}
	if got := keys(samples); !strings.HasPrefix(got, "mtgjson:low,mtgjson:high,mtgjson:held,synthetic:mtgjson:high/synergy,synthetic:mtgjson:high/copies") {
		t.Errorf("the check changed its input: %s", got)
	}
}
