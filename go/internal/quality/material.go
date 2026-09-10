package quality

import (
	"math"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/meta"
)

// DefaultSynergyFloor is the fall of the synergy feature a synergy break
// needs before its copy joins the fit, in standard deviations of the real
// training lists (D-652).
//
// A break must be material, or it is no defect and the label lies
// (D-485). The copies and the colors axes read that from the deck before
// the break. A pairing break can not, so the fit makes the copy and reads
// how far the feature moved. A weak precon's cards do not pair in the
// corpus, and a break that removes half its pairs removes half of nothing
// (F-95).
//
// The pairs a floor of 0.10 drops were won at chance, and each higher
// floor drops pairs the model reads. A higher floor lifts the bar with
// real defects (D-653).
const DefaultSynergyFloor = 0.10

// DefaultSynergyFormats names the formats the synergy check reads:
// Commander alone (D-653). In Modern and Standard the synergy break moves
// the synergy feature by nothing in most pairs, and the bar reads those
// copies through the shape features (F-96).
var DefaultSynergyFormats = []mtgv1.FormatId{mtgv1.FormatId_FORMAT_ID_COMMANDER}

// dropImmaterial drops every synergy copy whose break lowered the synergy
// feature by less than the floor (D-652). It answers the rows it keeps,
// the holdout copies it dropped by the axis the fit asked for, and the
// unit it read.
//
// The move reads the corpus of the fold, so the check reads the feature
// the model reads, and no holdout list decides a training row. The unit is
// the standard deviation of the feature over the real training lists. No
// copy takes part in it, so the unit stays the same at every floor.
//
// A break that lowered nothing is no break at any floor. A request whose
// own check fails and whose synergy copy fails this one makes no copy at
// all, so the bad rung loses the row. A negative floor keeps every copy,
// which is the fit before the check.
func dropImmaterial(samples []sample, floor float64) (kept []sample, dropped map[string]int, unit float64) {
	dropped = map[string]int{}
	if floor < 0 {
		return samples, dropped, 0
	}
	realSynergy := map[string]float64{}
	var n, sum, sq float64
	for _, s := range samples {
		if s.r.List.Source == meta.SourceSynthetic {
			continue
		}
		v := s.features[KeySynergy]
		realSynergy[s.r.List.Key()] = v
		if !s.hold {
			n++
			sum += v
			sq += v * v
		}
	}
	if n > 0 {
		mean := sum / n
		unit = math.Sqrt(math.Max(sq/n-mean*mean, 0))
	}
	kept = samples[:0:0]
	for _, s := range samples {
		l := s.r.List
		if l.Source == meta.SourceSynthetic && l.Defect == DefectSynergy {
			fall := realSynergy[baseKey(l)] - s.features[KeySynergy]
			if fall <= 0 || fall < floor*unit {
				if s.hold {
					dropped[requestedAxis(l)]++
				}
				continue
			}
		}
		kept = append(kept, s)
	}
	return kept, dropped, unit
}

// requestedAxis is the axis the fit asked a copy for: the part of its id
// after the last slash (F-55). A copy that fell to synergy names its first
// axis here and synergy in its defect.
func requestedAxis(l *meta.List) string {
	if i := strings.LastIndex(l.ID, "/"); i >= 0 {
		return l.ID[i+1:]
	}
	return l.Defect
}

// ImmaterialCount sums the copies the synergy check dropped over the
// folds (D-652).
func (fr *FormatReport) ImmaterialCount() int {
	n := 0
	for _, v := range fr.Immaterial {
		n += v
	}
	return n
}
