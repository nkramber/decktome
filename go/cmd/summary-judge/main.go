// Command summary-judge reads a deck gate document and asks the judge
// role whether any deck summary states a rule of the game, and whether
// that statement is true.
//
// This is the real check for F-26. The deterministic linter of
// internal/generate reads the shape of a rules claim and can not read its
// truth (D-229). The judge runs on another provider than the generator,
// so it never rates its own work (D-22).
//
// CAUTION: this calls a real provider and it costs money. SUMMARY_JUDGE=1
// is required, so it can not run by accident.
//
// Usage:
//
//	SUMMARY_JUDGE=1 go run ./cmd/summary-judge -in ../docs/reference/pr8-deck-gate-run3.md
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
)

var (
	deckRe    = regexp.MustCompile(`(?m)^### (\d+)\. (.+)$`)
	summaryRe = regexp.MustCompile(`(?m)^\*\*Summary:\*\* (.+)$`)
)

// deckBlock is one deck of the gate document: its heading and the
// summary under it, or an empty summary when the block holds none.
type deckBlock struct {
	ID      string
	Name    string
	Summary string
}

// deckBlocks reads the document deck by deck. Each block runs from one
// deck heading to the next, and the summary is read inside the block
// alone, so a deck with no summary can not take the next deck's.
func deckBlocks(text string) []deckBlock {
	heads := deckRe.FindAllStringSubmatchIndex(text, -1)
	out := make([]deckBlock, 0, len(heads))
	for i, h := range heads {
		end := len(text)
		if i+1 < len(heads) {
			end = heads[i+1][0]
		}
		block := deckBlock{ID: text[h[2]:h[3]], Name: text[h[4]:h[5]]}
		if m := summaryRe.FindStringSubmatch(text[h[1]:end]); m != nil {
			block.Summary = strings.TrimSpace(m[1])
		}
		out = append(out, block)
	}
	return out
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	in := flag.String("in", "", "a deck gate document")
	flag.Parse()
	if *in == "" {
		return fmt.Errorf("give -in, a deck gate document")
	}
	if err := gatekit.SpendGuard("SUMMARY_JUDGE"); err != nil {
		return err
	}
	raw, err := os.ReadFile(*in) // #nosec G304 -- the operator names the file.
	if err != nil {
		return err
	}
	blocks := deckBlocks(string(raw))
	if len(blocks) == 0 {
		return fmt.Errorf("%s holds no deck heading", *in)
	}

	client, err := llm.NewFromEnv(gatekit.Env, gatekit.Quiet())
	if err != nil {
		return err
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)

	clean, rules, falseRule, skipped := 0, 0, 0, 0
	for _, b := range blocks {
		if b.Summary == "" {
			skipped++
			fmt.Printf("%-3s %-42s no summary, skipped\n", b.ID+".", trunc(b.Name, 42))
			continue
		}
		out, err := generate.JudgeSummary(context.Background(), client, b.Name, b.Summary, acc)
		if err != nil {
			return fmt.Errorf("judge deck %s: %w", b.ID, err)
		}
		// The false-rule count reads the same helper the deck gate reads,
		// so a claim marked false counts whatever the verdict word says
		// (T-3).
		switch {
		case out.StatesAFalseRule():
			falseRule++
		case out.Verdict == "states_a_rule":
			rules++
		case out.Verdict == "clean":
			clean++
		}
		fmt.Printf("%-3s %-42s %s\n", b.ID+".", trunc(b.Name, 42), out.Verdict)
		for _, c := range out.Claims {
			fmt.Printf("      [%s] %q\n            %s\n", c.Truth, trunc(c.Text, 90), trunc(c.Why, 130))
		}
		// The deterministic net of D-224, for comparison.
		if shapes := generate.LintSummary(b.Summary); len(shapes) > 0 {
			fmt.Printf("      linter also flagged: %s\n", strings.Join(shapes, ", "))
		}
	}
	rep := acc.Report()
	fmt.Printf("\n%d decks: %d clean, %d state a rule, %d state a FALSE rule, %d with no summary\n",
		len(blocks), clean, rules, falseRule, skipped)
	fmt.Printf("calls: %d. Cost: %s\n", rep.Calls, gatekit.CostWord(rep))
	if falseRule > 0 {
		return fmt.Errorf("F-26: %d summaries state a false rule", falseRule)
	}
	return nil
}

func trunc(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}
