package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// report writes the gate document and returns the verdict. The bars come
// from the roadmap: every returned deck passes the block checks, no
// invented name reaches the user, and no summary states a false rule. A
// run of zero prompts fails: there is nothing to pass.
func report(w io.Writer, rs []result, acc *llm.Accumulator, idx *cards.Index, took time.Duration) bool {
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
	_, _ = fmt.Fprintf(w, "| Errors | %d |\n", errs)
	_, _ = fmt.Fprintf(w, "| Prompt version | %d |\n", generate.PromptVersion)
	rep := acc.Report()
	_, _ = fmt.Fprintf(w, "| Calls | %d |\n", rep.Calls)
	_, _ = fmt.Fprintf(w, "| Cost | %s |\n", gatekit.CostWord(rep))
	_, _ = fmt.Fprintf(w, "| Time | %.0f seconds |\n\n", took.Seconds())

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
	_, _ = fmt.Fprintf(w, "\n\n## Decks\n\n")
	for _, r := range rs {
		writeDeck(w, r)
	}
	return pass
}

func writeDeck(w io.Writer, r result) {
	p := r.prompt
	_, _ = fmt.Fprintf(w, "### %d. %s\n\n", p.ID, p.Name)
	_, _ = fmt.Fprintf(w, "Format: %s. Theme: %s. Pool: %s. Shortlist: %d names.\n\n",
		generate.FormatWord(gatekit.FormatID(p.Format)), p.Theme, p.Pool, r.poolSize)
	if r.err != nil {
		_, _ = fmt.Fprintf(w, "ERROR: %v\n\n", r.err)
		return
	}
	d := r.deck
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
