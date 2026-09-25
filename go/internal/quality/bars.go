package quality

import (
	"errors"
	"fmt"

	"github.com/nkramber/decktome/go/internal/meta"
)

// The bars of the quality gate (D-414, D-573). The gate and the daily
// refit read the same values, so a fit that the gate fails never serves
// (D-927).
const (
	BarGreatOverBaseline = 0.90
	BarBaselineOverOwn   = 0.95
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
		gb, bb := fr.Holdout.GreatOverBaseline, fr.Holdout.BaselineOverOwn
		if gb.Pairs == 0 || gb.Share() < BarGreatOverBaseline {
			out = append(out, fmt.Sprintf("%s: great over precon %.2f of %d, the bar is %.2f", word, gb.Share(), gb.Pairs, BarGreatOverBaseline))
		}
		if bb.Pairs == 0 || bb.Share() < BarBaselineOverOwn {
			out = append(out, fmt.Sprintf("%s: precon over own copy %.2f of %d, the bar is %.2f", word, bb.Share(), bb.Pairs, BarBaselineOverOwn))
		}
	}
	return out
}
