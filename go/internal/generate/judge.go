package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// The judge lane is the real check for F-26. The deterministic net of
// LintSummary reads the shape of a rules claim and can not read its
// truth.
//
// The judge runs on another provider than the generator, so it never
// rates its own work (D-22). It is cheap enough to run beside every gate
// (D-229).

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

// The bracket judge is the second bar of the bracket gate (PR-14A). It
// reads a deck with the bracket definitions in hand and names the
// bracket it would play at. The gate asks it to agree with the bracket
// the deck was built for in eight of ten.

const bracketJudgeInstructions = `You read one Commander deck and name the bracket it plays at.

The Commander brackets, from the Commander Format Panel (2025-02-11, revised 2025-10-21):
- Bracket 1, Exhibition: an ultra-casual deck, games of nine turns or more. No Game Changers, no mass land denial, no extra turns, no two-card infinite combos.
- Bracket 2, Core: near the strength of a preconstructed deck, games of eight turns or more. No Game Changers, no mass land denial, few extra turns and never chained, no two-card infinite combos.
- Bracket 3, Upgraded: souped up beyond a precon, games of six turns or more. Up to three Game Changers, no mass land denial, few extra turns and never chained, no cheap two-card infinite combo in about the first six turns.
- Bracket 4, Optimized: the strongest cards and decks, games can end from turn four. Only the ban list applies.
- Bracket 5, cEDH: competitive and metagame-focused, a game can end on any turn. Only the ban list applies. The deck plays the best strategy and not a theme.

Read the card list for its speed, its mana base, its fast mana and tutors, its interaction, its combos, and its Game Changers. Name one bracket, 1 to 5, and say why in two or three sentences. Judge the deck as it is, and not the bracket the builder may have aimed at.`

// The bracket is a string enum and not a bounded integer: the Anthropic
// structured-output schema refuses minimum and maximum on an integer
// (bracket gate run 1, 2026-09-02).
const bracketJudgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["bracket", "why"],
  "properties": {
    "bracket": {"type": "string", "enum": ["1", "2", "3", "4", "5"]},
    "why": {"type": "string"}
  }
}`

// BracketJudgement is the judge's bracket for one deck.
type BracketJudgement struct {
	Bracket int32  `json:"-"`
	Why     string `json:"why"`
}

// bracketOut is the judge's answer as the schema shapes it.
type bracketOut struct {
	Bracket string `json:"bracket"`
	Why     string `json:"why"`
}

// JudgeBracket asks the judge role which bracket a deck plays at. The
// deck goes out as a card list with the commander first, and the judge
// never sees the bracket the deck was built for.
func JudgeBracket(ctx context.Context, c *llm.Client, deck *mtgv1.Deck, cards rules.CardSource, acc *llm.Accumulator) (*BracketJudgement, error) {
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: bracketJudgeInstructions,
		Input:        DeckText(deck, cards),
		SchemaName:   "bracket_check",
		Schema:       json.RawMessage(bracketJudgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge bracket: %w", err)
	}
	var out bracketOut
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge bracket output: %w", err)
	}
	n, err := strconv.Atoi(out.Bracket)
	if err != nil || n < 1 || n > 5 {
		return nil, fmt.Errorf("judge bracket output: bracket %q", out.Bracket)
	}
	return &BracketJudgement{Bracket: int32(n), Why: out.Why}, nil
}

// DeckText writes a deck as the judge reads it: the commander, then one
// line per card with its count and its job, when the deck names one. No
// bracket, no summary, and no finding.
func DeckText(deck *mtgv1.Deck, cards rules.CardSource) string {
	var s strings.Builder
	for _, id := range deck.GetCommanderOracleIds() {
		if c, ok := cards.ByOracleID(id); ok {
			fmt.Fprintf(&s, "Commander: %s\n", c.GetName())
		}
	}
	s.WriteString("\nCards:\n")
	for _, dc := range deck.GetCards() {
		if dc.GetRole() == mtgv1.CardRole_CARD_ROLE_UNSPECIFIED {
			fmt.Fprintf(&s, "%d %s\n", dc.GetCount(), dc.GetName())
			continue
		}
		fmt.Fprintf(&s, "%d %s (%s)\n", dc.GetCount(), dc.GetName(), roleWord(dc.GetRole()))
	}
	return s.String()
}

// The tier judge of PR-14B: the judge names the rung of the ladder of
// D-414 a deck sits on, and the gate reads whether it agrees with the
// quality model's grade in eight of ten.
const tierJudgeInstructions = `You read one deck list and name its quality tier.

The ladder, from the top:
- great: a list that would place in the top 8 of a Challenge or the top cut of a cEDH tournament. Tuned, efficient, the cards and the pairs the best lists of the format play.
- good: a list that would post a winning record in a league or finish mid-field at a Challenge. Sound, on plan, a step below the top.
- typical: the average deck of its commander or archetype as the community builds it. Reasonable choices, no sharp edge.
- baseline: the strength of a preconstructed deck as sold. Playable, slow, generic cards, few of the format's best.
- bad: below a precon. A broken mana base, a curve that never lands its spells, playsets turned to singletons, or cards that do not work together.

Read the card list for its card quality against the format's best lists, its synergy, its mana base, its curve, and its interaction. Name one tier word and say why in two or three sentences. Judge the deck as it is.`

const tierJudgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["tier", "why"],
  "properties": {
    "tier": {"type": "string", "enum": ["great", "good", "typical", "baseline", "bad"]},
    "why": {"type": "string"}
  }
}`

// TierJudgement is the judge's tier for one deck.
type TierJudgement struct {
	Tier string `json:"tier"`
	Why  string `json:"why"`
}

// JudgeTier asks the judge role for the tier of a deck. The judge sees
// the format and the card list, and never the grade.
func JudgeTier(ctx context.Context, c *llm.Client, deck *mtgv1.Deck, cards rules.CardSource, acc *llm.Accumulator) (*TierJudgement, error) {
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: tierJudgeInstructions,
		Input:        "Format: " + FormatWord(deck.GetFormat().GetId()) + "\n" + DeckText(deck, cards),
		SchemaName:   "tier_check",
		Schema:       json.RawMessage(tierJudgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge tier: %w", err)
	}
	var out TierJudgement
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge tier output: %w", err)
	}
	if out.Tier == "" {
		return nil, fmt.Errorf("judge tier output: no tier")
	}
	return &out, nil
}
