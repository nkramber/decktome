package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

func formatID(s string) mtgv1.FormatId {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "commander":
		return mtgv1.FormatId_FORMAT_ID_COMMANDER
	case "standard":
		return mtgv1.FormatId_FORMAT_ID_STANDARD
	case "modern":
		return mtgv1.FormatId_FORMAT_ID_MODERN
	}
	return mtgv1.FormatId_FORMAT_ID_UNSPECIFIED
}

var colorNames = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G,
}

func colorList(in []string) []mtgv1.Color {
	var out []mtgv1.Color
	for _, s := range in {
		if c, ok := colorNames[strings.ToUpper(strings.TrimSpace(s))]; ok {
			out = append(out, c)
		}
	}
	return out
}

func poolRuleID(s string) mtgv1.PoolRule {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "owned_first":
		return mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
	case "owned_only":
		return mtgv1.PoolRule_POOL_RULE_OWNED_ONLY
	}
	return mtgv1.PoolRule_POOL_RULE_ANY_CARD
}

var sixtySteps = map[string]mtgv1.SixtyStep{
	"casual": mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
	"fnm":    mtgv1.SixtyStep_SIXTY_STEP_FNM,
	// The corpus calls the top step tournament-meta.
	"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}

func power(p prompt) *mtgv1.PowerLevel {
	if p.Bracket > 0 {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: p.Bracket}}
	}
	if step, ok := sixtySteps[strings.ToLower(p.Power)]; ok {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: step}}
	}
	return nil
}

func blocks(d *mtgv1.Deck) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	for _, f := range d.GetValidation().GetFindings() {
		if f.GetSeverity() == mtgv1.Severity_SEVERITY_BLOCK {
			out = append(out, f)
		}
	}
	return out
}

func countCards(d *mtgv1.Deck) int {
	n := 0
	for _, c := range d.GetCards() {
		n += int(c.GetCount())
	}
	return n
}

// report writes the gate document. The two bars come from the roadmap:
// every returned deck passes the block checks, and no invented name
// reaches the user.
func report(w io.Writer, rs []result, acc *llm.Accumulator, idx *cards.Index, took time.Duration) {
	built, clean, notes, repaired, errs := 0, 0, 0, 0, 0
	for _, r := range rs {
		switch {
		case r.err != nil:
			errs++
		case r.deck != nil:
			built++
			if len(blocks(r.deck)) == 0 {
				clean++
			}
			notes += len(r.notes)
			if r.repaired {
				repaired++
			}
		}
	}
	pass := errs == 0 && built == len(rs) && clean == built && notes == 0
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	_, _ = fmt.Fprintf(w, "# PR-8 deck gate\n\n")
	_, _ = fmt.Fprintf(w, "Run date: %s. Card snapshot: %s.\n\n", time.Now().UTC().Format("2006-01-02"), idx.AsOf.Format("2006-01-02"))
	_, _ = fmt.Fprintf(w, "Verdict: %s. %d of %d decks passed every block check, and %d invented names reached the user. Both bars are zero tolerance.\n\n",
		verdict, clean, len(rs), notes)
	if errs > 0 {
		_, _ = fmt.Fprintf(w, "%d prompts failed before a deck existed. A gate can not pass with an error.\n\n", errs)
	}

	_, _ = fmt.Fprintf(w, "## Summary\n\n| Measure | Value |\n|---|---|\n")
	_, _ = fmt.Fprintf(w, "| Prompts | %d |\n", len(rs))
	_, _ = fmt.Fprintf(w, "| Decks returned | %d |\n", built)
	_, _ = fmt.Fprintf(w, "| Decks with no block finding | %d |\n", clean)
	_, _ = fmt.Fprintf(w, "| Invented names that reached the user | %d |\n", notes)
	_, _ = fmt.Fprintf(w, "| Decks that needed the repair turn | %d |\n", repaired)
	_, _ = fmt.Fprintf(w, "| Errors | %d |\n", errs)
	_, _ = fmt.Fprintf(w, "| Prompt version | %d |\n", generate.PromptVersion)
	rep := acc.Report()
	_, _ = fmt.Fprintf(w, "| Calls | %d |\n", rep.Calls)
	cost := 0.0
	if rep.CostUSD != nil {
		cost = *rep.CostUSD
	}
	_, _ = fmt.Fprintf(w, "| Cost | $%.4f |\n", cost)
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
}

func writeDeck(w io.Writer, r result) {
	p := r.prompt
	_, _ = fmt.Fprintf(w, "### %d. %s\n\n", p.ID, p.Name)
	_, _ = fmt.Fprintf(w, "Format: %s. Theme: %s. Pool: %s. Shortlist: %d names.\n\n",
		generate.FormatWord(formatID(p.Format)), p.Theme, p.Pool, r.poolSize)
	if r.err != nil {
		_, _ = fmt.Fprintf(w, "ERROR: %v\n\n", r.err)
		return
	}
	d := r.deck
	_, _ = fmt.Fprintf(w, "Cards: %d. Repair turn: %v. Block findings: %d.\n\n",
		countCards(d), r.repaired, len(blocks(d)))
	_, _ = fmt.Fprintf(w, "**Summary:** %s\n\n", d.GetSummary())
	if claims := generate.LintSummary(d.GetSummary()); len(claims) > 0 {
		_, _ = fmt.Fprintf(w, "Summary rules claims (F-26): %s\n\n", strings.Join(claims, ", "))
	}
	for _, n := range r.notes {
		_, _ = fmt.Fprintf(w, "- NOTE: %s\n", n)
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
