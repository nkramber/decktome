// Command quality-gate fits the deck quality model over the stored
// lists and writes the PR-14B gate document.
//
// The bars come from the roadmap (PR-14B). On the holdout of each
// format a great list scores above a precon in 90 percent of the
// pairs, and a precon above a synthetic bad deck in 95 percent. A
// bracket 5 request offers three commanders from the top cuts of the
// Topdeck.gg cEDH tournaments of the last 90 days. The judge bar over
// the golden decks reads the next deck gate run, whose summaries carry
// the tier, and this command does not run it.
//
// The run is free: it calls no provider. It reads the meta store beside
// the card snapshot, which `make meta-refresh` fills. -write stores the
// fitted model as a new version, and without it the run leaves no
// model behind. The exit code is 1 on a FAIL verdict, and the document
// is written first.
//
// The judge lane is the one paid mode: -judge <deck gate document> asks
// the judge role for the tier of every graded deck of the document, and
// it needs QUALITY_JUDGE=1 (`make quality-judge`).
//
// Usage:
//
//	CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall go run ./cmd/quality-gate
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/gcpenv"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/quality"
)

// The bars.
const (
	barGreatOverBaseline = 0.90
	barBaselineOverBad   = 0.95
	commanderOffer       = 3
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	write := flag.Bool("write", false, "store the fitted model as a new version")
	judge := flag.String("judge", "", "run the tier judge lane over this deck gate document (CAUTION: costs money, needs QUALITY_JUDGE=1)")
	prompts := flag.String("prompts", "cmd/deck-gate/prompts.json", "the deck gate prompts, for the commander of a deck the document names none for")
	explain := flag.String("explain", "", "print how the stored model grades every deck of this deck gate document (free)")
	runOut := flag.String("run-out", "", "write the run header and the rows as JSONL here (PR-15)")
	flag.Parse()
	if *judge != "" {
		return runJudge(*judge, *prompts, *runOut)
	}
	if *explain != "" {
		return runExplain(*explain, *prompts)
	}
	// The run file is never overwritten (D-65).
	if err := gatekit.RefuseExisting(*runOut); err != nil {
		return err
	}
	run := evalrun.New("quality", evalrun.RunID(*runOut))
	ctx := context.Background()
	log := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(ctx, log)
	if err != nil {
		return err
	}
	run.SetSnapshot(idx.AsOf)
	store, err := gcpenv.MetaStore(ctx, gcpenv.LocalProject, nil)
	if err != nil {
		return err
	}
	start := time.Now()
	var model *quality.Model
	var rep *quality.FitReport
	if *write {
		model, rep, err = quality.Refit(ctx, store, idx, start.UTC(), log)
	} else {
		model, rep, err = quality.FitStore(ctx, store, idx, start.UTC(), log)
	}
	if err != nil && model == nil {
		return fmt.Errorf("quality gate: %w", err)
	}
	day, err := meta.LatestCommandersDay(ctx, store)
	if err != nil {
		return err
	}
	var reads []meta.Commander
	if day != "" {
		if reads, err = meta.ReadCommanders(ctx, store, day); err != nil {
			return err
		}
	}
	offer, offerErr := offerAtBracketFive(idx, model)
	if model != nil {
		run.Header.Versions["quality_model"] = model.Version
	}
	run.Header.Versions["commanders_day"] = gatekit.OrNone(day)
	pass := report(os.Stdout, idx, model, rep, reads, day, offer, offerErr, time.Since(start), run)
	if err := evalrun.WriteFile(*runOut, run); err != nil {
		return err
	}
	if !pass {
		return fmt.Errorf("quality gate: FAIL")
	}
	return nil
}

// offered is one commander of the bracket 5 offer with its signal.
type offered struct {
	name   string
	signal float64
}

// offerAtBracketFive asks the commander pool for the three commanders a
// bracket 5 request with no theme gets, with the model's signal wired.
func offerAtBracketFive(idx *cards.Index, model *quality.Model) ([]offered, error) {
	builder, err := candidates.New()
	if err != nil {
		return nil, err
	}
	scorer := quality.NewScorer(model)
	pool, err := builder.Commanders(idx, candidates.Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Bracket: 5, CommanderSignal: scorer.CommanderSignal(),
	}, commanderOffer)
	if err != nil {
		return nil, err
	}
	var out []offered
	for _, c := range pool {
		o := offered{name: c.DisplayName()}
		if fm := model.Format(mtgv1.FormatId_FORMAT_ID_COMMANDER); fm != nil {
			ids := []string{c.Card.GetOracleId()}
			if c.Partner != nil {
				ids = append(ids, c.Partner.GetOracleId())
			}
			sig, _ := fm.Commander(ids...)
			o.signal = sig.CEDH
		}
		out = append(out, o)
	}
	return out, nil
}

// report writes the document and answers the verdict.
func report(w io.Writer, idx *cards.Index, model *quality.Model, rep *quality.FitReport, reads []meta.Commander, day string, offer []offered, offerErr error, took time.Duration, run *evalrun.Run) bool {
	// The document goes to stdout, and a write error there ends the run
	// with a short document, which the verdict check refuses.
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	now := time.Now().UTC().Format("2006-01-02")
	pass := true
	var fails []string
	p("# PR-14B quality gate\n\nRun date: %s. Card snapshot: %s.", now, idx.AsOf.Format("2006-01-02"))
	if model != nil {
		p(" Model: %s.", model.Version)
	}
	p("\n")
	p("\n")

	// The pair bars per format.
	words := []string{meta.FormatCommander, meta.FormatStandard, meta.FormatModern}
	p("## Separation on the holdout" + "\n")
	p("\n")
	p("| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |" + "\n")
	p("|---|---|---|---|---|---|---|---|" + "\n")
	for _, word := range words {
		fr := rep.Formats[word]
		if fr == nil {
			p("| %s | 0 | 0 | 0 | 0 | no fit | no fit | no fit |\n", word)
			fails = append(fails, word+": no lists")
			pass = false
			run.Gate(word, "fit", 0, "no lists")
			continue
		}
		fm := (*quality.FormatModel)(nil)
		if model != nil {
			fm = model.Formats[word]
		}
		if fm == nil {
			p("| %s | %d | %d | %d | 0 | no fit | no fit | no fit |\n", word, fr.Read, fr.Used, fr.Synthetic)
			fails = append(fails, word+": no fit")
			pass = false
			run.Gate(word, "fit", 0, "no fit")
			continue
		}
		h := fm.Holdout
		gb, bb := h.GreatOverBaseline, h.BaselineOverBad
		run.Gate(word, "fit", 1, "")
		run.Gate(word, "great_over_precon", gb.Share(), fmt.Sprintf("%d pairs, the bar is %.2f", gb.Pairs, barGreatOverBaseline))
		run.Gate(word, "precon_over_bad", bb.Share(), fmt.Sprintf("%d pairs, the bar is %.2f", bb.Pairs, barBaselineOverBad))
		run.Info(word, "accuracy", h.Accuracy, "")
		run.Info(word, "lists", float64(fr.Read), fmt.Sprintf("%d used, %d synthetic, %d holdout", fr.Used, fr.Synthetic, h.Lists))
		p("| %s | %d | %d | %d | %d | %s | %s | %.2f |\n", word, fr.Read, fr.Used, fr.Synthetic, h.Lists,
			pairWord(gb, barGreatOverBaseline), pairWord(bb, barBaselineOverBad), h.Accuracy)
		if gb.Pairs == 0 || gb.Share() < barGreatOverBaseline {
			fails = append(fails, fmt.Sprintf("%s: great over precon %s", word, pairWord(gb, barGreatOverBaseline)))
			pass = false
		}
		if bb.Pairs == 0 || bb.Share() < barBaselineOverBad {
			fails = append(fails, fmt.Sprintf("%s: precon over bad %s", word, pairWord(bb, barBaselineOverBad)))
			pass = false
		}
	}
	p("\n")

	// The commander offer.
	p("## The bracket 5 offer" + "\n")
	p("\n")
	bySlug := map[string]meta.Commander{}
	topdeckRows := 0
	for _, c := range reads {
		bySlug[c.Slug] = c
		if c.Entries > 0 {
			topdeckRows++
		}
	}
	switch {
	case offerErr != nil:
		p("The offer failed: %v\n", offerErr)
		fails = append(fails, "offer: "+offerErr.Error())
		pass = false
	case len(offer) < commanderOffer:
		p("The pool offered %d commanders, and the bar is %d.\n", len(offer), commanderOffer)
		fails = append(fails, "offer: too few")
		pass = false
	default:
		p("Commander reads of %s: %d rows, %d with Topdeck.gg entries.\n\n", gatekit.OrNone(day), len(reads), topdeckRows)
		p("| Commander | cEDH signal | Top cuts | Entries |" + "\n")
		p("|---|---|---|---|" + "\n")
		offerPass := true
		for _, o := range offer {
			c := bySlug[meta.EDHRECSlug(strings.Split(o.name, " + ")...)]
			p("| %s | %.2f | %d | %d |\n", o.name, o.signal, c.TopCuts, c.Entries)
			if c.TopCuts == 0 {
				offerPass = false
			}
		}
		if !offerPass && topdeckRows == 0 {
			fails = append(fails, "offer: no Topdeck.gg read, so no commander comes from a top cut ("+meta.TopdeckKeyEnv+")")
		} else if !offerPass {
			fails = append(fails, "offer: a commander comes from no top cut")
		}
		pass = pass && offerPass
	}
	p("\n")

	// The model, per format.
	if model != nil {
		for _, word := range words {
			fm := model.Formats[word]
			fr := rep.Formats[word]
			if fm == nil {
				continue
			}
			p("## The %s model\n\n", word)
			p("Tiers, worst first: %s. Train counts: %s. Holdout counts: %s.\n\n",
				strings.Join(fm.Tiers, ", "), countWords(fm.TrainCounts), countWords(fm.HoldoutCounts))
			p("Cards with a rate: %d. Pairs that lift: %d. Commanders with a signal: %d. Top lists shape: %.1f lands at %.2f over %d lists.\n\n",
				len(fm.CardRates), len(fm.Pairs), len(fm.Commanders), fm.Shape.Lands, fm.Shape.AvgManaValue, fm.Shape.Lists)
			if len(fr.Dropped) > 0 {
				p("Features with no spread, dropped: %s.\n\n", strings.Join(fr.Dropped, ", "))
			}
			p("| Feature | Mean | Std | Weight |" + "\n")
			p("|---|---|---|---|" + "\n")
			for i, k := range fm.Keys {
				p("| `%s` | %.3f | %.3f | %+.3f |\n", k, fm.Means[i], fm.Stds[i], fm.Weights[i])
			}
			p("\n")
			p("Thresholds: %s. Loss: %.3f over %d iterations. Defect detector accuracy on the holdout: %.2f, at a cut of %.2f.\n\n", floats(fm.Thresholds), fr.Loss, fr.Iterations, fm.Holdout.DefectAccuracy, fm.DefectThreshold)
			if len(fm.Holdout.BadByDefect) > 0 {
				p("The precon bar per broken axis:" + "\n")
				p("\n")
				for _, axis := range quality.Defects {
					if ps, ok := fm.Holdout.BadByDefect[axis]; ok {
						p("- %s: %.2f of %d\n", axis, ps.Share(), ps.Pairs)
					}
				}
				p("\n")
			}
			p("Confusion on the holdout, rows are the label and columns the grade, worst first:" + "\n")
			p("\n")
			for i, row := range fm.Holdout.Confusion {
				p("- %s: %v\n", fm.Tiers[i], row)
			}
			p("\n")
			if len(fm.TopCards) > 0 {
				p("The cards the top lists hold most: %s.\n\n", strings.Join(fm.TopCards, ", "))
			}
		}
	}

	offerOK := 1.0
	if offerErr != nil || len(offer) < commanderOffer || !pass && len(fails) > 0 && strings.HasPrefix(fails[len(fails)-1], "offer:") {
		offerOK = 0
	}
	run.Gate("suite", "offer", offerOK, fmt.Sprintf("%d commanders, the bar is %d", len(offer), commanderOffer))

	verdict := "PASS"
	if !pass {
		verdict = "FAIL"
	}
	run.Finish(llm.Report{}, took, verdict)
	run.Markdown(w)
	p("\n")

	// The verdict, last in the file and first for a reader.
	p("## Verdict" + "\n")
	p("\n")
	p("Verdict: %s. Time: %.0f seconds, no provider call.", verdict, took.Seconds())
	if len(fails) > 0 {
		p(" Failed bars: %s.", strings.Join(fails, "; "))
	}
	p("\n")
	return pass
}

func pairWord(p quality.PairShare, bar float64) string {
	if p.Pairs == 0 {
		return "no pair"
	}
	return fmt.Sprintf("%.2f of %d (bar %.2f)", p.Share(), p.Pairs, bar)
}

func countWords(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s %d", k, m[k]))
	}
	if len(parts) == 0 {
		return "none"
	}
	return strings.Join(parts, ", ")
}

func floats(v []float64) string {
	parts := make([]string, 0, len(v))
	for _, f := range v {
		parts = append(parts, fmt.Sprintf("%.3f", f))
	}
	return strings.Join(parts, ", ")
}
