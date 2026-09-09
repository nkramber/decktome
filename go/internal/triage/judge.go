package triage

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nkramber/decktome/go/internal/harvest"
	"github.com/nkramber/decktome/go/internal/llm"
)

// The judge lane reads the verdicts the reason keys can not place
// (D-643). It runs on the judge role, which is another provider than the
// one that made the thing the reader judged, so it never rates its own
// work (D-22).
//
// It answers one of two things: the class of the fault, or an owner
// question. A reader who argues with a decision the owner made raises a
// question and never a fix (D-558). The prompt carries the decisions a
// reader most often argues with, and the judge names one by its id.

// Verdict is the judge's answer for one item.
type Verdict struct {
	Class         string `json:"class"`
	OwnerQuestion bool   `json:"owner_question"`
	Decision      string `json:"decision"`
	Why           string `json:"why"`
}

// Judger reads one verdict and answers its class. The command passes the
// live judge, and a test passes its own.
type Judger func(ctx context.Context, rec harvest.Record, r Route) (Verdict, error)

// disputedDecisions are the rules a reader argues with, in the words of
// the decision. The judge picks one by its id, and it may pick none.
var disputedDecisions = []struct{ ID, Says string }{
	{"D-37", "Owned-first means the deck prefers cards the reader owns and still buys the rest. Only owned-only refuses every card the reader does not own."},
	{"D-459", "The Commander bracket rules come from the Commander Format Panel, and this app follows them. What a bracket allows is not this app's choice."},
	{"D-63", "The agent may offer a commander before it asks about the card pool."},
	{"D-2", "The app reads the reader's ManaBox collection export, and it prices what the reader must buy."},
}

const judgeInstructions = `You triage one piece of feedback about a Magic: The Gathering deck builder.

A reader gave a thumbs down on a question the app asked, a deck summary it wrote, one card it picked, a whole deck it built, or the chat itself. Read what the reader said and name the class of the fault, from the list below.

Name exactly one class id. Pick the class that matches what the reader is unhappy about, and not the kind of object alone.

Some complaints are not faults. They argue with a rule the product's owner already decided. When the reader is asking for a different rule, and not reporting a thing that went wrong, set owner_question to true and name the decision id. Do not set it for a reader who reports that the app broke its own rule: that is a fault and it takes a class.

Say why in one or two sentences.`

// judgeSchema keeps class as a free string. The class list changes with
// the reason keys, and an enum in the schema would need a second place
// to hold it.
const judgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["class", "owner_question", "decision", "why"],
  "properties": {
    "class": {"type": "string"},
    "owner_question": {"type": "boolean"},
    "decision": {"type": "string"},
    "why": {"type": "string"}
  }
}`

// JudgePrompt writes the instructions the judge reads. It names every
// class of the kind at hand and every decision a reader argues with, so
// the judge never invents an id.
func JudgePrompt(kind string) string {
	var b strings.Builder
	b.WriteString(judgeInstructions)
	b.WriteString("\n\nThe classes:\n")
	for _, c := range classes {
		if kind != "" && c.Kind != kind {
			continue
		}
		fmt.Fprintf(&b, "- %s, %s: the reader means %q\n", c.ID, c.Name, c.Reason)
	}
	b.WriteString("\nThe decisions a reader may be arguing with:\n")
	for _, d := range disputedDecisions {
		fmt.Fprintf(&b, "- %s: %s\n", d.ID, d.Says)
	}
	b.WriteString("\nAnswer with an empty class when you set owner_question. Answer with an empty decision otherwise.")
	return b.String()
}

// JudgeInput writes what the judge reads about one verdict. It carries
// the reader's own words, which is the point of the lane, and it carries
// no uid: the judge needs no user (D-559).
func JudgeInput(rec harvest.Record) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Kind: %s\n", rec.Kind)
	if len(rec.Reasons) > 0 {
		fmt.Fprintf(&b, "Reasons checked: %s\n", strings.Join(rec.Reasons, ", "))
	}
	if rec.Question != "" {
		fmt.Fprintf(&b, "The question the app asked: %s\n", rec.Question)
	}
	if rec.Answer != "" {
		fmt.Fprintf(&b, "What the reader had answered: %s\n", rec.Answer)
	}
	if rec.Text != "" {
		fmt.Fprintf(&b, "What the reader wrote: %s\n", rec.Text)
	}
	return b.String()
}

// LiveJudger asks the judge role. The caller owns the client and the
// accumulator, so the cost of the lane joins the run's ledger.
func LiveJudger(c *llm.Client, acc *llm.Accumulator) Judger {
	return func(ctx context.Context, rec harvest.Record, _ Route) (Verdict, error) {
		res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
			Instructions: JudgePrompt(rec.Kind),
			Input:        JudgeInput(rec),
			SchemaName:   "feedback_triage",
			Schema:       json.RawMessage(judgeSchema),
		}, acc)
		if err != nil {
			return Verdict{}, fmt.Errorf("judge triage: %w", err)
		}
		var out Verdict
		if err := json.Unmarshal(res.Output, &out); err != nil {
			return Verdict{}, fmt.Errorf("judge triage output: %w", err)
		}
		return out, nil
	}
}
