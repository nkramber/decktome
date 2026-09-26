package quality

import (
	"errors"
	"fmt"

	"github.com/nkramber/decktome/go/internal/meta"
)

// The bars of the quality gate (D-414, D-573). The gate and the daily
// refit read the same values, so a fit that the gate fails never serves
// (D-927).
//
// A precon beats its own broken copies of the lands, curve, colors, and
// copies axes at BarBaselineOverOwn, and its own synergy copies at
// BarSynergyOverOwn. The synergy axis reads 0.80 to 0.85 over the real
// lists and 0.70 with no synergy check, so one bar over all the axes
// failed on the share of synergy copies alone (D-957).
const (
	BarGreatOverBaseline = 0.90
	BarBaselineOverOwn   = 0.95
	BarSynergyOverOwn    = 0.75
)

// BarFormats are the formats the bars read.
var BarFormats = []string{meta.FormatCommander, meta.FormatStandard, meta.FormatModern}

// ErrBarFailed is the refusal of a fit that fails a bar. Refit then
// keeps the stored model (D-927).
var ErrBarFailed = errors.New("quality: the fit fails a bar")

// BarFailures names each bar the fit fails, and nil when it passes
// all of them. A format with no lists or no fit fails.
func BarFailures(model *Model, rep *FitReport) []string {
	var out []string
	for _, word := range BarFormats {
		var fr *FormatReport
		if rep != nil {
			fr = rep.Formats[word]
		}
		if fr == nil {
			out = append(out, word+": no lists")
			continue
		}
		if model == nil || model.Formats[word] == nil {
			out = append(out, word+": no fit")
			continue
		}
		gb := fr.Holdout.GreatOverBaseline
		if gb.Pairs == 0 || gb.Share() < BarGreatOverBaseline {
			out = append(out, fmt.Sprintf("%s: great over precon %.2f of %d, the bar is %.2f", word, gb.Share(), gb.Pairs, BarGreatOverBaseline))
		}
		out = append(out, OwnBarFailures(word, fr.Holdout)...)
	}
	return out
}

// OwnShares splits the own-copy pairs of h into the synergy axis and
// every other axis (D-957).
func OwnShares(h Holdout) (axes, synergy PairShare) {
	for axis, ps := range h.OwnByDefect {
		if axis == DefectSynergy {
			synergy = ps
			continue
		}
		axes.Pairs += ps.Pairs
		axes.Wins += ps.Wins
	}
	return axes, synergy
}

// OwnBarFailures names each own-copy bar of the format that h fails. A
// bar with no pairs fails.
func OwnBarFailures(word string, h Holdout) []string {
	var out []string
	axes, synergy := OwnShares(h)
	if axes.Pairs == 0 || axes.Share() < BarBaselineOverOwn {
		out = append(out, fmt.Sprintf("%s: precon over own copy, synergy aside, %.2f of %d, the bar is %.2f", word, axes.Share(), axes.Pairs, BarBaselineOverOwn))
	}
	if synergy.Pairs == 0 || synergy.Share() < BarSynergyOverOwn {
		out = append(out, fmt.Sprintf("%s: precon over own synergy copy %.2f of %d, the bar is %.2f", word, synergy.Share(), synergy.Pairs, BarSynergyOverOwn))
	}
	return out
}
