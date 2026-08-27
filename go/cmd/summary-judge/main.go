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
	"encoding/json"
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

const judgeInstructions = `You check one paragraph from a Magic: The Gathering deck builder.

The paragraph is a deck summary written for the person who will play the deck. It must describe the deck and state no rule of the game. The rules engine reports the rules, and the summary must not.

Report two things.

First, every statement in the paragraph that asserts a rule of the game. A rule of the game is anything about what a card may do, what a format allows, what is banned or legal, what counts toward a limit, how many copies a deck may hold, or whether a card can lead a deck. A description of what the deck does on the table is not a rule.

Second, for each such statement, whether it is true. Judge it against the real rules of Magic: The Gathering. Say "unknown" when you can not tell.

Be strict about what counts as a rules claim and honest about truth. A summary with no rules claim is the expected result.`

const judgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["claims", "verdict"],
  "properties": {
    "claims": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["text", "truth", "why"],
        "properties": {
          "text": {"type": "string"},
          "truth": {"type": "string", "enum": ["true", "false", "unknown"]},
          "why": {"type": "string"}
        }
      }
    },
    "verdict": {"type": "string", "enum": ["clean", "states_a_rule", "states_a_falseRulerule"]}
  }
}`

type judgeOut struct {
	Claims []struct {
		Text  string `json:"text"`
		Truth string `json:"truth"`
		Why   string `json:"why"`
	} `json:"claims"`
	Verdict string `json:"verdict"`
}

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
		res, err := client.Complete(context.Background(), llm.RoleJudge, llm.Request{
			Instructions: judgeInstructions,
			Input:        fmt.Sprintf("Deck: %s\n\nSummary:\n%s", m[2], summary),
			SchemaName:   "summary_check",
			Schema:       json.RawMessage(judgeSchema),
		}, acc)
		if err != nil {
			return fmt.Errorf("judge deck %s: %w", m[1], err)
		}
		var out judgeOut
		if err := json.Unmarshal(res.Output, &out); err != nil {
			return err
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
