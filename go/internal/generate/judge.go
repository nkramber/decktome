package generate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// The judge lane is the real check for F-26. The deterministic net of
// LintSummary reads the shape of a rules claim and can not read its
// truth, and it found neither of the two false claims that reached a
// user in gate run 14 of 2026-08-26.
//
// The judge runs on another provider than the generator, so it never
// rates its own work (D-22). It costs about $0.0034 a deck, which is
// cheap enough to run beside every gate (D-229).

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
    "verdict": {"type": "string", "enum": ["clean", "states_a_rule", "states_a_false_rule"]}
  }
}`

// Claim is one rules statement the judge found in a summary.
type Claim struct {
	Text  string `json:"text"`
	Truth string `json:"truth"`
	Why   string `json:"why"`
}

// Judgement is the judge's answer for one summary.
type Judgement struct {
	Claims  []Claim `json:"claims"`
	Verdict string  `json:"verdict"`
}

// StatesAFalseRule reports the F-26 failure: the summary asserts a rule
// of the game and the rule is wrong.
func (j Judgement) StatesAFalseRule() bool {
	if j.Verdict == "states_a_false_rule" {
		return true
	}
	for _, c := range j.Claims {
		if c.Truth == "false" {
			return true
		}
	}
	return false
}

// JudgeSummary asks the judge role whether one summary states a rule of
// the game. deck names the deck, for the judge's context.
func JudgeSummary(ctx context.Context, c *llm.Client, deck, summary string, acc *llm.Accumulator) (*Judgement, error) {
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: judgeInstructions,
		Input:        fmt.Sprintf("Deck: %s\n\nSummary:\n%s", deck, summary),
		SchemaName:   "summary_check",
		Schema:       json.RawMessage(judgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge summary: %w", err)
	}
	var out Judgement
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge summary output: %w", err)
	}
	return &out, nil
}
