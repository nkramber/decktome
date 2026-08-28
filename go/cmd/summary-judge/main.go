// Command summary-judge reads a deck gate document and asks the judge
// role whether any deck summary states a rule of the game, and whether
// that statement is true.
//
// This is the real check for F-26. The deterministic linter of
// internal/generate reads the shape of a rules claim and can not read its
// truth, and it found neither of the two false claims that reached a user
// in gate run 14. The judge runs on another provider than the generator,
// so it never rates its own work (D-22).
//
// CAUTION: this calls a real provider and it costs money.
//
// Usage:
//
//	go run ./cmd/summary-judge -in ../docs/reference/pr8-deck-gate-run3.md
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

var (
	deckRe    = regexp.MustCompile(`(?m)^### (\d+)\. (.+)$`)
	summaryRe = regexp.MustCompile(`(?m)^\*\*Summary:\*\* (.+)$`)
)

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
	raw, err := os.ReadFile(*in) // #nosec G304 -- the operator names the file.
	if err != nil {
		return err
	}
	text := string(raw)
	names := deckRe.FindAllStringSubmatch(text, -1)
	sums := summaryRe.FindAllStringSubmatch(text, -1)
	if len(names) != len(sums) {
		return fmt.Errorf("the document holds %d decks and %d summaries", len(names), len(sums))
	}

	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	env := func(k string) string {
		if k == llm.EnvRequireKeys {
			return "1"
		}
		return os.Getenv(k)
	}
	client, err := llm.NewFromEnv(env, quiet)
	if err != nil {
		return err
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)

	clean, rules, falseRule := 0, 0, 0
	for i, m := range names {
		summary := strings.TrimSpace(sums[i][1])
		out, err := generate.JudgeSummary(context.Background(), client, m[2], summary, acc)
		if err != nil {
			return fmt.Errorf("judge deck %s: %w", m[1], err)
		}
		switch out.Verdict {
		case "clean":
			clean++
		case "states_a_rule":
			rules++
		case "states_a_falseRulerule":
			falseRule++
		}
		fmt.Printf("%-3s %-42s %s\n", m[1]+".", trunc(m[2], 42), out.Verdict)
		for _, c := range out.Claims {
			fmt.Printf("      [%s] %q\n            %s\n", c.Truth, trunc(c.Text, 90), trunc(c.Why, 130))
		}
		// The deterministic net of D-224, for comparison.
		if shapes := generate.LintSummary(summary); len(shapes) > 0 {
			fmt.Printf("      linter also flagged: %s\n", strings.Join(shapes, ", "))
		}
	}
	rep := acc.Report()
	cost := 0.0
	if rep.CostUSD != nil {
		cost = *rep.CostUSD
	}
	fmt.Printf("\n%d summaries: %d clean, %d state a rule, %d state a FALSE rule\n", len(names), clean, rules, falseRule)
	fmt.Printf("calls: %d. Cost: $%.4f\n", rep.Calls, cost)
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
