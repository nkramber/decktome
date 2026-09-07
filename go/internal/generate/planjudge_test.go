package generate

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
)

func planClient(t *testing.T, output string) *llm.Client {
	t.Helper()
	sc := llm.NewScript(llm.Step{Output: json.RawMessage(output)})
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	return c
}

// TestJudgePlanReadsTheFourGrades: the four words parse, the values
// read no as 0, partly as 0.5, and yes as 1, and the score is the mean.
func TestJudgePlanReadsTheFourGrades(t *testing.T) {
	c := planClient(t, `{
		"plan_coherent": {"grade": "yes", "why": "one plan"},
		"theme_fit": {"grade": "partly", "why": "half the theme"},
		"useful_as_built": {"grade": "no", "why": "no lands"},
		"summary_honest": {"grade": "yes", "why": "plain"}
	}`)
	deck := &mtgv1.Deck{Summary: "a deck", Cards: []*mtgv1.DeckCard{{Name: "Sol Ring", Count: 1}}}
	j, err := JudgePlan(context.Background(), c, "a lifegain deck", deck, source{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if j.ThemeFit.Value() != 0.5 || j.UsefulAsBuilt.Value() != 0 || j.Grade("summary_honest").Why != "plain" {
		t.Errorf("judgement = %+v", j)
	}
	if j.Score() != 0.625 {
		t.Errorf("score = %v, want the mean of 1, 0.5, 0, 1", j.Score())
	}
	if len(PlanFields) != 4 || PlanFields[0] != "plan_coherent" {
		t.Errorf("fields = %v", PlanFields)
	}
}

// TestJudgePlanRefusesAWordOutsideTheScale: a grade the rubric does not
// hold is an error and never a zero. The client checks the schema and
// retries once, so the script serves the bad answer twice.
func TestJudgePlanRefusesAWordOutsideTheScale(t *testing.T) {
	bad := `{
		"plan_coherent": {"grade": "excellent", "why": "x"},
		"theme_fit": {"grade": "yes", "why": "x"},
		"useful_as_built": {"grade": "yes", "why": "x"},
		"summary_honest": {"grade": "yes", "why": "x"}
	}`
	sc := llm.NewScript(llm.Step{Output: json.RawMessage(bad)}, llm.Step{Output: json.RawMessage(bad)})
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	_, err = JudgePlan(context.Background(), c, "a deck", &mtgv1.Deck{}, source{}, nil)
	if err == nil || !strings.Contains(err.Error(), "judge plan") {
		t.Errorf("err = %v, want a judge plan error", err)
	}
	if j := (PlanJudgement{PlanCoherent: PlanGrade{Grade: "excellent"}}); j.PlanCoherent.Value() != 0 {
		t.Error("a word outside the scale reads as 0, never more")
	}
}

// TestPlanReasonEmpty: a blank reason and the word "placeholder" count as
// no reason, and a sentence counts as one.
func TestPlanReasonEmpty(t *testing.T) {
	for _, why := range []string{"", "  ", "placeholder", "Placeholder"} {
		if !(PlanGrade{Grade: "yes", Why: why}).ReasonEmpty() {
			t.Errorf("%q must read as no reason", why)
		}
	}
	if (PlanGrade{Grade: "no", Why: "the deck has no lands"}).ReasonEmpty() {
		t.Error("a sentence is a reason")
	}
	j := PlanJudgement{
		PlanCoherent:  PlanGrade{Grade: "yes", Why: "one plan"},
		ThemeFit:      PlanGrade{Grade: "yes", Why: "placeholder"},
		UsefulAsBuilt: PlanGrade{Grade: "yes", Why: ""},
		SummaryHonest: PlanGrade{Grade: "yes", Why: "plain"},
	}
	if n := j.EmptyReasons(); n != 2 {
		t.Errorf("EmptyReasons = %d, want 2", n)
	}
}
