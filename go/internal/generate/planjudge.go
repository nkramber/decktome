package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// The plan judge of PR-15 reads a built deck with the request it answers
// and grades it on a fixed rubric of four fields, each on a three-point
// scale as D-66 shaped the M-5 rubric. Its rows are information until a
// baseline exists, and a bar reads a number and never the prose (lesson
// 11). The fields wait on the owner (OQ-58).

// PlanRubricVersion changes when a field or its words change. A change
// starts a new epoch of the plan rows.
const PlanRubricVersion = 2

const planJudgeInstructions = `You read one deck a deck builder made for a person, with the request the person wrote and the summary the builder wrote for them. You grade the deck on four fields. Each field takes one of three words: no, partly, or yes.

- plan_coherent: the cards serve one plan. A deck whose pieces pull in different directions, or whose summary names a plan the cards do not carry, reads no.
- theme_fit: the deck is the deck the person asked for. Read the request: the theme, the format, the commander, the colors, the power, and any limit the person named. A deck that answers a different request reads no.
- useful_as_built: a person can pick this deck up and play it as it stands. A deck with a broken mana base, a curve that never lands its spells, or too few ways to win reads no.
- summary_honest: the summary claims nothing the deck lacks and hides nothing the deck does. A summary that names a strategy the list does not carry reads no.

Give one sentence of reason per field. Judge the deck as it is. Do not grade the power of the deck against tournament lists: that is the job of another judge. Do not grade the format legality, the color identity, or the card counts: code checked them against the card data of the run date before you read the deck, and your knowledge of the card pool can be older than that data.`

// The reason comes before the grade in the schema. With the grade first
// the judge wrote "placeholder" as the reason of the last field in 10 of
// 25 decks (deck gate run 16), and a reason written first is a reason.
const planGradeSchema = `{"type": "object", "additionalProperties": false, "required": ["why", "grade"], "properties": {"why": {"type": "string"}, "grade": {"type": "string", "enum": ["no", "partly", "yes"]}}}`

var planJudgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["plan_coherent", "theme_fit", "useful_as_built", "summary_honest"],
  "properties": {
    "plan_coherent": ` + planGradeSchema + `,
    "theme_fit": ` + planGradeSchema + `,
    "useful_as_built": ` + planGradeSchema + `,
    "summary_honest": ` + planGradeSchema + `
  }
}`

// PlanGrade is one field of the rubric: the word and the reason.
type PlanGrade struct {
	Grade string `json:"grade"`
	Why   string `json:"why"`
}

// Value reads the word as a number: no is 0, partly is 0.5, yes is 1.
func (g PlanGrade) Value() float64 {
	switch g.Grade {
	case "yes":
		return 1
	case "partly":
		return 0.5
	}
	return 0
}

// ReasonEmpty reads a reason the judge did not write: blank, or the one
// word "placeholder". A grade with no reason still counts, because a bar
// reads the number and never the prose (lesson 11), and the count of
// empty reasons is a number of its own.
func (g PlanGrade) ReasonEmpty() bool {
	w := strings.ToLower(strings.TrimSpace(g.Why))
	return w == "" || w == "placeholder"
}

// PlanJudgement is the judge's four grades for one deck.
type PlanJudgement struct {
	PlanCoherent  PlanGrade `json:"plan_coherent"`
	ThemeFit      PlanGrade `json:"theme_fit"`
	UsefulAsBuilt PlanGrade `json:"useful_as_built"`
	SummaryHonest PlanGrade `json:"summary_honest"`
}

// PlanFields names the fields of the rubric, in rubric order.
var PlanFields = []string{"plan_coherent", "theme_fit", "useful_as_built", "summary_honest"}

// Grade reads one field by name.
func (j PlanJudgement) Grade(field string) PlanGrade {
	switch field {
	case "plan_coherent":
		return j.PlanCoherent
	case "theme_fit":
		return j.ThemeFit
	case "useful_as_built":
		return j.UsefulAsBuilt
	case "summary_honest":
		return j.SummaryHonest
	}
	return PlanGrade{}
}

// EmptyReasons counts the fields whose reason the judge did not write.
func (j PlanJudgement) EmptyReasons() int {
	n := 0
	for _, f := range PlanFields {
		if j.Grade(f).ReasonEmpty() {
			n++
		}
	}
	return n
}

// Score is the mean of the four values, 0 to 1.
func (j PlanJudgement) Score() float64 {
	sum := 0.0
	for _, f := range PlanFields {
		sum += j.Grade(f).Value()
	}
	return sum / float64(len(PlanFields))
}

// JudgePlan asks the judge role to grade a deck on the plan rubric. The
// judge sees the request, the summary, and the card list.
func JudgePlan(ctx context.Context, c *llm.Client, request string, deck *mtgv1.Deck, cards rules.CardSource, acc *llm.Accumulator) (*PlanJudgement, error) {
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: planJudgeInstructions,
		Input:        "Request: " + request + "\n\nFormat: " + FormatWord(deck.GetFormat().GetId()) + "\n\nSummary:\n" + deck.GetSummary() + "\n\n" + DeckText(deck, cards),
		SchemaName:   "plan_check",
		Schema:       json.RawMessage(planJudgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge plan: %w", err)
	}
	var out PlanJudgement
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge plan output: %w", err)
	}
	for _, f := range PlanFields {
		if g := out.Grade(f).Grade; g != "no" && g != "partly" && g != "yes" {
			return nil, fmt.Errorf("judge plan output: %s reads %q", f, g)
		}
	}
	return &out, nil
}
