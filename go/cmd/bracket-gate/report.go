package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/profile"
)

// JudgeAgreementPercent is the third bar: the judge must name the
// bracket the deck was built for in eight of ten decks (roadmap PR-14A).
const JudgeAgreementPercent = 80

// offBand lists the features of a deck outside their band.
func offBand(d *mtgv1.Deck) []*mtgv1.ProfileFeature {
	var out []*mtgv1.ProfileFeature
	for _, f := range d.GetProfile().GetFeatures() {
		if f.GetOffBand() {
			out = append(out, f)
		}
	}
	return out
}

// contentFindings lists the content-rule findings of a deck.
func contentFindings(d *mtgv1.Deck) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	for _, f := range d.GetValidation().GetFindings() {
		switch f.GetCode() {
		case profile.CodeMassLandDenial, profile.CodeExtraTurns, profile.CodeTwoCardCombo:
			out = append(out, f)
		}
	}
	return out
}

// report writes the gate document and returns the verdict. A run of
// zero prompts fails, and so does a deck with no profile or no judge
// verdict: a bar with nothing to read can not pass.
func report(w io.Writer, rs []result, acc *llm.Accumulator, idx *cards.Index, took time.Duration, run *evalrun.Run) bool {
	built, clean, inBand, noContent, unchecked, repaired, errs := 0, 0, 0, 0, 0, 0, 0
	judged, agreed, judgeErrs, noProfile := 0, 0, 0, 0
	for _, r := range rs {
		if r.judgeErr != nil {
			judgeErrs++
		}
		switch {
		case r.err != nil:
			errs++
			continue
		case r.deck == nil:
			continue
		}
		built++
		if len(gatekit.BlockFindings(r.deck)) == 0 {
			clean++
		}
		if r.repaired {
			repaired++
		}
		if r.deck.GetProfile() == nil {
			noProfile++
		} else {
			if len(offBand(r.deck)) == 0 {
				inBand++
			}
			if len(contentFindings(r.deck)) == 0 {
				noContent++
			}
			if !r.deck.GetProfile().GetContent().GetChecked() {
				unchecked++
			}
		}
		if r.judged != nil {
			judged++
			if r.judged.Bracket == r.prompt.Bracket {
				agreed++
			}
		}
	}
	agreement := 0
	if judged > 0 {
		agreement = agreed * 100 / judged
	}
	pass := len(rs) > 0 && errs == 0 && built == len(rs) && clean == built && noProfile == 0 &&
		inBand == built && noContent == built && unchecked == 0 && judgeErrs == 0 && judged == built &&
		agreement >= JudgeAgreementPercent
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	recordRows(run, rs)
	run.Gate("suite", "judge_agreement", float64(agreement), fmt.Sprintf("%d of %d, the bar is %d", agreed, judged, JudgeAgreementPercent))
	run.Finish(acc.Report(), took, verdict)
	_, _ = fmt.Fprintf(w, "# PR-14A bracket gate\n\n")
	_, _ = fmt.Fprintf(w, "Run date: %s. Card snapshot: %s.\n\n", time.Now().UTC().Format("2006-01-02"), idx.AsOf.Format("2006-01-02"))
	_, _ = fmt.Fprintf(w, "Verdict: %s. %d of %d decks passed every block check, %d sat in every band, %d held no content violation, and the judge agreed with the bracket on %d of %d (%d percent, the bar is %d).\n\n",
		verdict, clean, len(rs), inBand, noContent, agreed, judged, agreement, JudgeAgreementPercent)
	if len(rs) == 0 {
		_, _ = fmt.Fprintf(w, "The run held no prompt. A gate can not pass with nothing to measure.\n\n")
	}
	if errs > 0 {
		_, _ = fmt.Fprintf(w, "%d prompts failed before a deck existed. A gate can not pass with an error.\n\n", errs)
	}
	if noProfile > 0 {
		_, _ = fmt.Fprintf(w, "%d decks carry no profile. The band bar can not pass without one.\n\n", noProfile)
	}
	if unchecked > 0 {
		_, _ = fmt.Fprintf(w, "%d decks got no content check, because the endpoint did not answer. The content bar can not pass without one.\n\n", unchecked)
	}
	if judgeErrs > 0 {
		_, _ = fmt.Fprintf(w, "%d decks got no judge verdict, because the judge call failed. The judge bar can not pass without one.\n\n", judgeErrs)
	}

	_, _ = fmt.Fprintf(w, "## Summary\n\n| Measure | Value |\n|---|---|\n")
	_, _ = fmt.Fprintf(w, "| Prompts | %d |\n", len(rs))
	_, _ = fmt.Fprintf(w, "| Decks returned | %d |\n", built)
	_, _ = fmt.Fprintf(w, "| Decks with no block finding | %d |\n", clean)
	_, _ = fmt.Fprintf(w, "| Decks in every band | %d |\n", inBand)
	_, _ = fmt.Fprintf(w, "| Decks with no content violation | %d |\n", noContent)
	_, _ = fmt.Fprintf(w, "| Decks the endpoint did not check | %d |\n", unchecked)
	_, _ = fmt.Fprintf(w, "| Decks that needed a repair turn | %d |\n", repaired)
	_, _ = fmt.Fprintf(w, "| Decks judged | %d |\n", judged)
	_, _ = fmt.Fprintf(w, "| Judge agreed with the bracket | %d |\n", agreed)
	_, _ = fmt.Fprintf(w, "| Judge errors | %d |\n", judgeErrs)
	_, _ = fmt.Fprintf(w, "| Errors | %d |\n", errs)
	_, _ = fmt.Fprintf(w, "| Prompt version | %d |\n", generate.PromptVersion)
	rep := acc.Report()
	_, _ = fmt.Fprintf(w, "| Calls | %d |\n", rep.Calls)
	_, _ = fmt.Fprintf(w, "| Cost | %s |\n", gatekit.CostWord(rep))
	_, _ = fmt.Fprintf(w, "| Time | %.0f seconds |\n\n", took.Seconds())
	run.Markdown(w)
	_, _ = fmt.Fprintf(w, "\n")

	_, _ = fmt.Fprintf(w, "## Off-band features by key\n\nEach row counts the decks whose feature sat outside its band after the repair turns.\n\n| Feature | Decks off band |\n|---|---|\n")
	byKey := map[string]int{}
	var keys []string
	for _, r := range rs {
		for _, f := range offBand(r.deck) {
			if byKey[f.GetKey()] == 0 {
				keys = append(keys, f.GetKey())
			}
			byKey[f.GetKey()]++
		}
	}
	for _, k := range keys {
		_, _ = fmt.Fprintf(w, "| `%s` | %d |\n", k, byKey[k])
	}
	if len(keys) == 0 {
		_, _ = fmt.Fprintf(w, "| none | 0 |\n")
	}
	_, _ = fmt.Fprintf(w, "\n## Decks\n\n")
	for _, r := range rs {
		writeDeck(w, r)
	}
	return pass
}

func writeDeck(w io.Writer, r result) {
	p := r.prompt
	_, _ = fmt.Fprintf(w, "### %d. Bracket %d, %s, %s\n\n", p.ID, p.Bracket, p.Commander, p.Theme)
	if r.err != nil {
		_, _ = fmt.Fprintf(w, "ERROR: %s\n\n", r.err.Error())
		return
	}
	if r.deck == nil {
		_, _ = fmt.Fprintf(w, "No deck.\n\n")
		return
	}
	d := r.deck
	_, _ = fmt.Fprintf(w, "Pool %d. %d cards, %d block findings, repaired %v", r.poolSize, gatekit.CountCards(d), len(gatekit.BlockFindings(d)), r.repaired)
	if r.repairReason != "" {
		_, _ = fmt.Fprintf(w, " (%s)", r.repairReason)
	}
	_, _ = fmt.Fprintf(w, ".\n\n")
	if r.judged != nil {
		word := "agrees"
		if r.judged.Bracket != p.Bracket {
			word = "disagrees"
		}
		_, _ = fmt.Fprintf(w, "Judge: bracket %d, %s. %s\n\n", r.judged.Bracket, word, strings.TrimSpace(r.judged.Why))
	} else if r.judgeErr != nil {
		_, _ = fmt.Fprintf(w, "Judge: error, %s\n\n", r.judgeErr.Error())
	}
	prof := d.GetProfile()
	if prof == nil {
		_, _ = fmt.Fprintf(w, "No profile.\n\n")
	} else {
		_, _ = fmt.Fprintf(w, "| Feature | Value | Band | Off band | Note |\n|---|---|---|---|---|\n")
		for _, f := range prof.GetFeatures() {
			band := ""
			switch {
			case f.GetHasHigh():
				band = fmt.Sprintf("%g to %g", f.GetLow(), f.GetHigh())
			case f.GetLow() != 0:
				band = fmt.Sprintf("%g or more", f.GetLow())
			default:
				band = "none"
			}
			off := ""
			if f.GetOffBand() {
				off = "yes"
			}
			_, _ = fmt.Fprintf(w, "| `%s` | %g | %s | %s | %s |\n", f.GetKey(), f.GetValue(), band, off, f.GetNote())
		}
		if g := prof.GetGoldfish(); g != nil {
			_, _ = fmt.Fprintf(w, "\nGoldfish over %d hands: the commander on turn %g, %g mana on turn four, and %g of first hands hold two to four lands.\n",
				g.GetHands(), g.GetCommanderTurn(), g.GetManaTurnFour(), g.GetShareTwoToFourLands())
		}
		if c := prof.GetContent(); c != nil {
			if !c.GetChecked() {
				_, _ = fmt.Fprintf(w, "\nContent: unchecked, %s.\n", c.GetError())
			} else {
				_, _ = fmt.Fprintf(w, "\nContent: Spellbook tag %s. Game Changers %s. Mass land denial %s. Extra turns %s. Combos %d.\n",
					gatekit.OrNone(c.GetSourceTag()), gatekit.OrNone(strings.Join(c.GetGameChangers(), ", ")),
					gatekit.OrNone(strings.Join(c.GetMassLandDenial(), ", ")), gatekit.OrNone(strings.Join(c.GetExtraTurns(), ", ")), len(c.GetCombos()))
				for _, combo := range c.GetCombos() {
					two := ""
					if combo.GetTwoCard() {
						two = ", two cards"
					}
					_, _ = fmt.Fprintf(w, "- %s (speed %d%s)\n", strings.Join(combo.GetCards(), " + "), combo.GetSpeed(), two)
				}
			}
		}
		_, _ = fmt.Fprintf(w, "\n")
	}
	if fs := d.GetValidation().GetFindings(); len(fs) > 0 {
		_, _ = fmt.Fprintf(w, "Findings:\n\n")
		for _, f := range fs {
			_, _ = fmt.Fprintf(w, "- %s `%s`: %s\n", strings.TrimPrefix(f.GetSeverity().String(), "SEVERITY_"), f.GetCode(), f.GetMessage())
		}
		_, _ = fmt.Fprintf(w, "\n")
	}
	for _, n := range r.notes {
		_, _ = fmt.Fprintf(w, "- Note: %s\n", n)
	}
	_, _ = fmt.Fprintf(w, "Cards:\n\n")
	for _, c := range d.GetCards() {
		_, _ = fmt.Fprintf(w, "- %d %s\n", c.GetCount(), c.GetName())
	}
	_, _ = fmt.Fprintf(w, "\n")
}

// recordRows writes one row per bar per deck into the run (PR-15). The
// judge agreement is a bar over the whole run, so the per-deck row is
// information and the suite row is the gate.
func recordRows(run *evalrun.Run, rs []result) {
	for _, r := range rs {
		item := fmt.Sprintf("%d", r.prompt.ID)
		if r.err != nil {
			run.Gate(item, "built", 0, r.err.Error())
			continue
		}
		if r.deck == nil {
			run.Gate(item, "built", 0, "no deck")
			continue
		}
		run.Gate(item, "built", 1, "")
		var codes []string
		for _, f := range gatekit.BlockFindings(r.deck) {
			codes = append(codes, f.GetCode())
		}
		run.Gate(item, "blocks", float64(len(codes)), strings.Join(codes, ", "))
		if r.deck.GetProfile() == nil {
			run.Gate(item, "profile", 0, "no profile")
		} else {
			run.Gate(item, "profile", 1, "")
			var keys []string
			for _, f := range offBand(r.deck) {
				keys = append(keys, f.GetKey())
			}
			run.Gate(item, "off_band", float64(len(keys)), strings.Join(keys, ", "))
			var content []string
			for _, f := range contentFindings(r.deck) {
				content = append(content, f.GetCode())
			}
			run.Gate(item, "content_violations", float64(len(content)), strings.Join(content, ", "))
			checked := 0.0
			if r.deck.GetProfile().GetContent().GetChecked() {
				checked = 1
			}
			run.Gate(item, "content_checked", checked, "")
		}
		switch {
		case r.judgeErr != nil:
			run.Gate(item, "judge_error", 1, r.judgeErr.Error())
		case r.judged != nil:
			agrees := 0.0
			if r.judged.Bracket == r.prompt.Bracket {
				agrees = 1
			}
			run.Info(item, "judge_agrees", agrees, fmt.Sprintf("built for %d, judged %d", r.prompt.Bracket, r.judged.Bracket))
		}
		repaired := 0.0
		if r.repaired {
			repaired = 1
		}
		run.Info(item, "repaired", repaired, r.repairReason)
		run.Info(item, "pool", float64(r.poolSize), "")
	}
}
