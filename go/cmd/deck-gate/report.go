package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// report writes the gate document and returns the verdict. The bars come
// from the roadmap: every returned deck passes the block checks, no
// invented name reaches the user, and no summary states a false rule. A
// run of zero prompts fails: there is nothing to pass. The run gets one
// row per bar per deck, and the document carries its fingerprint (PR-15).
func report(w io.Writer, rs []result, acc *llm.Accumulator, idx *cards.Index, took time.Duration, run *evalrun.Run) bool {
	built, clean, notes, repaired, errs := 0, 0, 0, 0, 0
	judged, falseRules, statesRule, judgeErrs := 0, 0, 0, 0
	for _, r := range rs {
		if r.judgeErr != nil {
			judgeErrs++
		}
		switch {
		case r.err != nil:
			errs++
		case r.deck != nil:
			built++
			if len(gatekit.BlockFindings(r.deck)) == 0 {
				clean++
			}
			notes += len(r.notes)
			if r.repaired {
				repaired++
			}
		}
		if r.judged != nil {
			judged++
			if r.judged.StatesAFalseRule() {
				falseRules++
			} else if len(r.judged.Claims) > 0 {
				statesRule++
			}
		}
	}
	// F-26 is a bar and not a footnote: the deterministic linter can not
	// read the truth of a rules claim (D-229). A deck the judge could not
	// read has no verdict on that bar, so it can not pass it (T-17).
	pass := len(rs) > 0 && errs == 0 && built == len(rs) && clean == built && notes == 0 && falseRules == 0 && judgeErrs == 0
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	rep := acc.Report()
	recordRows(run, rs)
	run.Finish(rep, took, verdict)
	_, _ = fmt.Fprintf(w, "# PR-8 deck gate\n\n")
	_, _ = fmt.Fprintf(w, "Run date: %s. Card snapshot: %s.\n\n", time.Now().UTC().Format("2006-01-02"), idx.AsOf.Format("2006-01-02"))
	_, _ = fmt.Fprintf(w, "Verdict: %s. %d of %d decks passed every block check, %d invented names reached the user, and %d summaries stated a false rule of the game. All three bars are zero tolerance.\n\n",
		verdict, clean, len(rs), notes, falseRules)
	if len(rs) == 0 {
		_, _ = fmt.Fprintf(w, "The run held no prompt. A gate can not pass with nothing to measure.\n\n")
	}
	if errs > 0 {
		_, _ = fmt.Fprintf(w, "%d prompts failed before a deck existed. A gate can not pass with an error.\n\n", errs)
	}
	if judgeErrs > 0 {
		_, _ = fmt.Fprintf(w, "%d decks got no judge verdict, because the judge call failed. The F-26 bar can not pass without one.\n\n", judgeErrs)
	}

	_, _ = fmt.Fprintf(w, "## Summary\n\n| Measure | Value |\n|---|---|\n")
	_, _ = fmt.Fprintf(w, "| Prompts | %d |\n", len(rs))
	_, _ = fmt.Fprintf(w, "| Decks returned | %d |\n", built)
	_, _ = fmt.Fprintf(w, "| Decks with no block finding | %d |\n", clean)
	_, _ = fmt.Fprintf(w, "| Invented names that reached the user | %d |\n", notes)
	_, _ = fmt.Fprintf(w, "| Decks that needed the repair turn | %d |\n", repaired)
	_, _ = fmt.Fprintf(w, "| Summaries judged (F-26) | %d |\n", judged)
	_, _ = fmt.Fprintf(w, "| Summaries that state a rule of the game | %d |\n", statesRule)
	_, _ = fmt.Fprintf(w, "| Summaries that state a FALSE rule | %d |\n", falseRules)
	_, _ = fmt.Fprintf(w, "| Judge errors | %d |\n", judgeErrs)
	planJudged, planScore, planEmpty := planTotals(rs)
	_, _ = fmt.Fprintf(w, "| Decks the plan judge read (PR-15, information) | %d |\n", planJudged)
	if planJudged > 0 {
		_, _ = fmt.Fprintf(w, "| Mean plan score, 0 to 1 | %.2f |\n", planScore)
		_, _ = fmt.Fprintf(w, "| Plan reasons the judge left empty | %d |\n", planEmpty)
	}
	_, _ = fmt.Fprintf(w, "| Errors | %d |\n", errs)
	_, _ = fmt.Fprintf(w, "| Prompt version | %d |\n", generate.PromptVersion)
	_, _ = fmt.Fprintf(w, "| Calls | %d |\n", rep.Calls)
	_, _ = fmt.Fprintf(w, "| Cost | %s |\n", gatekit.CostWord(rep))
	_, _ = fmt.Fprintf(w, "| Time | %.0f seconds |\n\n", took.Seconds())
	run.Markdown(w)
	_, _ = fmt.Fprintf(w, "\n")

	_, _ = fmt.Fprintf(w, "## Findings by code\n\nA block stops the deck. A warning and a note are reports.\n\n")
	byCode := map[string]int{}
	bySev := map[string]int{}
	for _, r := range rs {
		for _, f := range r.deck.GetValidation().GetFindings() {
			byCode[f.GetCode()]++
			bySev[strings.TrimPrefix(f.GetSeverity().String(), "SEVERITY_")]++
		}
	}
	codes := make([]string, 0, len(byCode))
	for c := range byCode {
		codes = append(codes, c)
	}
	sort.Slice(codes, func(i, j int) bool { return byCode[codes[i]] > byCode[codes[j]] })
	_, _ = fmt.Fprintf(w, "| Code | Count |\n|---|---|\n")
	for _, c := range codes {
		_, _ = fmt.Fprintf(w, "| `%s` | %d |\n", c, byCode[c])
	}
	_, _ = fmt.Fprintf(w, "\nBy severity: ")
	for _, s := range []string{"BLOCK", "WARN", "INFO"} {
		_, _ = fmt.Fprintf(w, "%s %d. ", s, bySev[s])
	}
	_, _ = fmt.Fprintf(w, "\n\n")
	setReport(w, rs)
	exclusionReport(w, rs)
	_, _ = fmt.Fprintf(w, "## Decks\n\n")
	for _, r := range rs {
		writeDeck(w, r, idx)
	}
	return pass
}

func writeDeck(w io.Writer, r result, idx *cards.Index) {
	p := r.prompt
	_, _ = fmt.Fprintf(w, "### %d. %s\n\n", p.ID, p.Name)
	_, _ = fmt.Fprintf(w, "Format: %s. Theme: %s. Pool: %s. Shortlist: %d names.\n\n",
		generate.FormatWord(gatekit.FormatID(p.Format)), p.Theme, p.Pool, r.poolSize)
	if r.err != nil {
		_, _ = fmt.Fprintf(w, "ERROR: %v\n\n", r.err)
		return
	}
	d := r.deck
	// The commander and the grade go out as lines of their own, so the
	// tier judge lane of PR-14B reads the deck back whole.
	var commanders []string
	for _, id := range d.GetCommanderOracleIds() {
		if c, ok := idx.ByOracleID(id); ok {
			commanders = append(commanders, c.GetName())
		}
	}
	if len(commanders) > 0 {
		_, _ = fmt.Fprintf(w, "Commander: %s.\n\n", strings.Join(commanders, ", "))
	}
	if q := d.GetQuality(); q != nil {
		_, _ = fmt.Fprintf(w, "Grade: %s, score %.2f, model %s.\n\n", q.GetTier(), q.GetScore(), q.GetModelVersion())
	}
	repair := "no"
	if r.repaired {
		repair = "yes, for " + r.repairReason
	}
	// The sideboard count is printed because the engine refuses only a
	// sideboard that is too big, and a deck with none passes (D-233).
	_, _ = fmt.Fprintf(w, "Cards: %d main, %d sideboard. Repair turn: %s. Block findings: %d.\n\n",
		gatekit.CountCards(d), gatekit.CountSideboard(d), repair, len(gatekit.BlockFindings(d)))
	_, _ = fmt.Fprintf(w, "Cost: $%.2f to buy, $%.2f the whole deck.\n\n",
		generate.BuyCost(d), generate.DeckCost(d))
	_, _ = fmt.Fprintf(w, "**Summary:** %s\n\n", d.GetSummary())
	if claims := generate.LintSummary(d.GetSummary()); len(claims) > 0 {
		_, _ = fmt.Fprintf(w, "Summary rules claims (F-26): %s\n\n", strings.Join(claims, ", "))
	}
	for _, n := range r.notes {
		_, _ = fmt.Fprintf(w, "- NOTE: %s\n", n)
	}
	if r.judgeErr != nil {
		_, _ = fmt.Fprintf(w, "- JUDGE ERROR: %v\n", r.judgeErr)
	}
	if r.judged != nil {
		for _, c := range r.judged.Claims {
			_, _ = fmt.Fprintf(w, "- JUDGE [%s]: %q. %s\n", c.Truth, c.Text, c.Why)
		}
	}
	if r.planErr != nil {
		_, _ = fmt.Fprintf(w, "- PLAN JUDGE ERROR: %v\n", r.planErr)
	}
	if r.plan != nil {
		for _, f := range generate.PlanFields {
			g := r.plan.Grade(f)
			_, _ = fmt.Fprintf(w, "- PLAN %s=%s: %s\n", f, g.Grade, g.Why)
		}
	}
	for _, f := range d.GetValidation().GetFindings() {
		_, _ = fmt.Fprintf(w, "- [%s] `%s`: %s\n", strings.TrimPrefix(f.GetSeverity().String(), "SEVERITY_"), f.GetCode(), f.GetMessage())
	}
	_, _ = fmt.Fprintf(w, "\n<details><summary>The deck list</summary>\n\n")
	for _, c := range d.GetCards() {
		_, _ = fmt.Fprintf(w, "- %d %s | %s | %s\n", c.GetCount(), c.GetName(),
			strings.ToLower(strings.TrimPrefix(c.GetRole().String(), "CARD_ROLE_")), c.GetReason())
	}
	_, _ = fmt.Fprintf(w, "\n</details>\n\n")
}

// setReport writes the PR-17B block: which sets each limited prompt
// applied, how many cards they held, and how many cards of the deck the
// sets do not hold (D-373, D-383). A run with no set-limited prompt
// writes nothing.
func setReport(w io.Writer, rs []result) {
	var limited []result
	for _, r := range rs {
		if len(r.setCodes) > 0 {
			limited = append(limited, r)
		}
	}
	if len(limited) == 0 {
		return
	}
	_, _ = fmt.Fprintf(w, "## The set filter (PR-17B)\n\n")
	_, _ = fmt.Fprintf(w, "A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).\n\n")
	_, _ = fmt.Fprintf(w, "| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |\n|---|---|---|---|---|---|---|\n")
	for _, r := range limited {
		outside, marked := 0, 0
		for _, list := range [][]*mtgv1.DeckCard{r.deck.GetCards(), r.deck.GetSideboard()} {
			for _, dc := range list {
				if dc.GetOutsideRequestedSets() {
					marked++
				}
			}
		}
		outside = marked
		_, _ = fmt.Fprintf(w, "| %d | %s | `%s` | %d | %d | %d | %d |\n",
			r.prompt.ID, r.prompt.Name, strings.Join(r.setCodes, ","), r.inSet, r.outside, outside, marked)
	}
	_, _ = fmt.Fprintf(w, "\n")
}

// exclusionReport writes the PR-24 block: which precons each prompt
// excluded, how many cards left the pool, and how many cards of the deck
// belong to them (D-408). A card that slips through is a block, and the
// verdict reads it. A run with no such prompt writes nothing.
func exclusionReport(w io.Writer, rs []result) {
	var excluding []result
	for _, r := range rs {
		if len(r.products) > 0 {
			excluding = append(excluding, r)
		}
	}
	if len(excluding) == 0 {
		return
	}
	_, _ = fmt.Fprintf(w, "## The precon exclusion (PR-24)\n\n")
	_, _ = fmt.Fprintf(w, "A deck asked to use no card of a precon holds none of its cards. The products' copies leave the owned counts, so a card with a spare copy in the binder stays usable, and a basic land never leaves (D-408, D-37). An excluded card in the deck is a block.\n\n")
	_, _ = fmt.Fprintf(w, "| # | Prompt | Products | Cards excluded | Cards spare | Excluded cards in the deck | Blocked |\n|---|---|---|---|---|---|---|\n")
	for _, r := range excluding {
		inDeck := 0
		for _, id := range r.deck.GetCommanderOracleIds() {
			if r.excluded[id] {
				inDeck++
			}
		}
		for _, list := range [][]*mtgv1.DeckCard{r.deck.GetCards(), r.deck.GetSideboard()} {
			for _, dc := range list {
				if r.excluded[dc.GetOracleId()] {
					inDeck++
				}
			}
		}
		blocked := 0
		for _, f := range r.deck.GetValidation().GetFindings() {
			if f.GetCode() == rules.CodeExcludedPrecon {
				blocked++
			}
		}
		_, _ = fmt.Fprintf(w, "| %d | %s | %s | %d | %d | %d | %d |\n",
			r.prompt.ID, r.prompt.Name, strings.Join(r.products, ", "), len(r.excluded), r.spare, inDeck, blocked)
	}
	_, _ = fmt.Fprintf(w, "\n")
}

// recordRows writes one row per bar per deck into the run, and the
// information rows beside them (PR-15). A gate row carries a bar of the
// verdict, so a compare names the deck that flipped and never a count.
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
		run.Gate(item, "invented_names", float64(len(r.notes)), strings.Join(r.notes, " | "))
		switch {
		case r.judgeErr != nil:
			run.Gate(item, "judge_error", 1, r.judgeErr.Error())
		case r.judged != nil:
			falseRule := 0.0
			if r.judged.StatesAFalseRule() {
				falseRule = 1
			}
			run.Gate(item, "false_rules", falseRule, "")
			run.Info(item, "rules_claims", float64(len(r.judged.Claims)), "")
		}
		repaired := 0.0
		if r.repaired {
			repaired = 1
		}
		run.Info(item, "repaired", repaired, r.repairReason)
		run.Info(item, "cards", float64(gatekit.CountCards(r.deck)), "")
		run.Info(item, "pool", float64(r.poolSize), "")
		warnings := 0
		for _, f := range r.deck.GetValidation().GetFindings() {
			if f.GetSeverity() == mtgv1.Severity_SEVERITY_WARN {
				warnings++
			}
		}
		run.Info(item, "warnings", float64(warnings), "")
		run.Info(item, "buy_cost", generate.BuyCost(r.deck), "")
		run.Info(item, "deck_cost", generate.DeckCost(r.deck), "")
		if r.plan != nil {
			for _, f := range generate.PlanFields {
				g := r.plan.Grade(f)
				run.Info(item, "plan_"+f, g.Value(), g.Grade+": "+g.Why)
			}
			run.Info(item, "plan_score", r.plan.Score(), "")
			run.Info(item, "plan_reasons_empty", float64(r.plan.EmptyReasons()), "")
		} else if r.planErr != nil {
			run.Info(item, "plan_judge_error", 1, r.planErr.Error())
		}
		if q := r.deck.GetQuality(); q != nil {
			run.Info(item, "grade", float64(q.GetScore()), q.GetTier())
			if run.Header.Versions["quality_model"] == "" {
				run.Header.Versions["quality_model"] = q.GetModelVersion()
			}
		}
		if len(r.setCodes) > 0 {
			marked := 0
			for _, list := range [][]*mtgv1.DeckCard{r.deck.GetCards(), r.deck.GetSideboard()} {
				for _, dc := range list {
					if dc.GetOutsideRequestedSets() {
						marked++
					}
				}
			}
			run.Info(item, "outside_set", float64(marked), strings.Join(r.setCodes, ","))
		}
		if len(r.products) > 0 {
			inDeck := 0
			for _, id := range r.deck.GetCommanderOracleIds() {
				if r.excluded[id] {
					inDeck++
				}
			}
			for _, list := range [][]*mtgv1.DeckCard{r.deck.GetCards(), r.deck.GetSideboard()} {
				for _, dc := range list {
					if r.excluded[dc.GetOracleId()] {
						inDeck++
					}
				}
			}
			run.Gate(item, "excluded_in_deck", float64(inDeck), strings.Join(r.products, ", "))
			run.Info(item, "excluded", float64(len(r.excluded)), "")
			run.Info(item, "spare", float64(r.spare), "")
		}
	}
}

// planTotals counts the decks the plan judge read, their mean score, and
// the reasons it left empty.
func planTotals(rs []result) (judged int, mean float64, empty int) {
	sum := 0.0
	for _, r := range rs {
		if r.plan != nil {
			judged++
			sum += r.plan.Score()
			empty += r.plan.EmptyReasons()
		}
	}
	if judged > 0 {
		mean = sum / float64(judged)
	}
	return judged, mean, empty
}
