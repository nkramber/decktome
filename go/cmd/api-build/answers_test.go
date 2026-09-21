package main

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func TestParsePlanReadsEveryForm(t *testing.T) {
	p, err := parsePlan("commander=Karlov of the Ghost Council; power=#2 ;budget=decline")
	if err != nil {
		t.Fatalf("parsePlan: %v", err)
	}
	if got := p["commander"].text; got != "Karlov of the Ghost Council" {
		t.Errorf("commander text = %q", got)
	}
	if got := p["power"].option; got != 2 {
		t.Errorf("power option = %d, want 2", got)
	}
	if !p["budget"].decline {
		t.Error("budget is not a decline")
	}
	if len(p) != 3 {
		t.Errorf("plan holds %d slots, want 3", len(p))
	}
}

func TestParsePlanRefusesBadInput(t *testing.T) {
	for _, in := range []string{
		"commander",
		"=text",
		"power=#0",
		"power=#x",
		"budget=",
		"power=#1;power=#2",
	} {
		if _, err := parsePlan(in); err == nil {
			t.Errorf("parsePlan(%q) took a bad answer", in)
		}
	}
}

func TestParsePlanEmptyIsEmpty(t *testing.T) {
	p, err := parsePlan("  ")
	if err != nil {
		t.Fatalf("parsePlan: %v", err)
	}
	if len(p) != 0 {
		t.Errorf("plan holds %d slots, want 0", len(p))
	}
}

// A closed question takes the whole answer space in its options (D-295),
// so the run must never answer it with free text.
func TestAnswerForClosedQuestionTakesFirstOption(t *testing.T) {
	q := &mtgv1.Question{Id: "q1", Slot: "format", Options: []string{"Commander", "Modern"}, Closed: true}
	a, err := answerFor(q, plan{})
	if err != nil {
		t.Fatalf("answerFor: %v", err)
	}
	if a.OptionIndex == nil || a.GetOptionIndex() != 0 {
		t.Fatalf("closed question did not take option 0: %+v", a)
	}
	if a.GetDeclined() {
		t.Error("a closed question was declined")
	}
}

// The commander row shows no decline control (D-690), so the run must
// not hand that choice back.
func TestAnswerForNoDeclineNeverDeclines(t *testing.T) {
	q := &mtgv1.Question{Id: "q2", Slot: "commander", Options: []string{"Karlov of the Ghost Council"}, NoDecline: true}
	a, err := answerFor(q, plan{})
	if err != nil {
		t.Fatalf("answerFor: %v", err)
	}
	if a.GetDeclined() {
		t.Error("a no_decline question was declined")
	}
	if _, err := answerFor(q, plan{"commander": {decline: true}}); err == nil {
		t.Error("a planned decline on a no_decline question was allowed")
	}
}

func TestAnswerForDeclinesAnOpenQuestion(t *testing.T) {
	q := &mtgv1.Question{Id: "q3", Slot: "budget", Text: "What is your budget?"}
	a, err := answerFor(q, plan{})
	if err != nil {
		t.Fatalf("answerFor: %v", err)
	}
	if !a.GetDeclined() {
		t.Errorf("an open question was not declined: %+v", a)
	}
	if a.OptionIndex != nil || a.GetText() != "" {
		t.Error("a declined answer carries an option or a text (D-353)")
	}
}

// The plan counts options from one, and the wire counts from zero.
func TestAnswerForOptionIsOneBasedOnTheCommandLine(t *testing.T) {
	q := &mtgv1.Question{Id: "q4", Slot: "power", Options: []string{"1", "2", "3"}}
	a, err := answerFor(q, plan{"power": {option: 3}})
	if err != nil {
		t.Fatalf("answerFor: %v", err)
	}
	if a.GetOptionIndex() != 2 {
		t.Errorf("option #3 became index %d, want 2", a.GetOptionIndex())
	}
	if _, err := answerFor(q, plan{"power": {option: 4}}); err == nil {
		t.Error("an option past the offered ones was allowed")
	}
}

func TestAnswerForRefusesFreeTextOnAClosedQuestion(t *testing.T) {
	q := &mtgv1.Question{Id: "q5", Slot: "format", Options: []string{"Commander"}, Closed: true}
	if _, err := answerFor(q, plan{"format": {text: "Pauper"}}); err == nil {
		t.Error("free text on a closed question was allowed")
	}
}

func TestAnswerForRefusesAQuestionItCanNotAnswer(t *testing.T) {
	q := &mtgv1.Question{Id: "q6", Slot: "commander", NoDecline: true}
	_, err := answerFor(q, plan{})
	if err == nil {
		t.Fatal("a question with no option and no decline was answered")
	}
	if !strings.Contains(err.Error(), "commander") {
		t.Errorf("the error names no slot: %v", err)
	}
}

func TestAnswerAllKeepsTheOrderOfTheQuestions(t *testing.T) {
	qs := []*mtgv1.Question{
		{Id: "a", Slot: "one"},
		{Id: "b", Slot: "two"},
		{Id: "c", Slot: "three"},
	}
	as, err := answerAll(qs, plan{})
	if err != nil {
		t.Fatalf("answerAll: %v", err)
	}
	if len(as) != 3 {
		t.Fatalf("answerAll gave %d answers, want 3", len(as))
	}
	for i, a := range as {
		if a.GetQuestionId() != qs[i].GetId() {
			t.Errorf("answer %d replies to %q, want %q", i, a.GetQuestionId(), qs[i].GetId())
		}
	}
}

func TestParsePoolRule(t *testing.T) {
	cases := map[string]mtgv1.PoolRule{
		"owned-first": mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
		"owned-only":  mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		"any":         mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		"ask":         mtgv1.PoolRule_POOL_RULE_UNSPECIFIED,
		"":            mtgv1.PoolRule_POOL_RULE_UNSPECIFIED,
	}
	for in, want := range cases {
		got, err := parsePoolRule(in)
		if err != nil {
			t.Fatalf("parsePoolRule(%q): %v", in, err)
		}
		if got != want {
			t.Errorf("parsePoolRule(%q) = %v, want %v", in, got, want)
		}
	}
	if _, err := parsePoolRule("everything"); err == nil {
		t.Error("an unknown pool rule was allowed")
	}
}

func TestDescribeNamesTheChosenOption(t *testing.T) {
	q := &mtgv1.Question{Slot: "power", Options: []string{"bracket 2", "bracket 3"}}
	i := int32(1)
	if got := describe(q, &mtgv1.Answer{OptionIndex: &i}); got != `option #2 "bracket 3"` {
		t.Errorf("describe = %q", got)
	}
	if got := describe(q, &mtgv1.Answer{Declined: true}); got != "declined" {
		t.Errorf("describe = %q", got)
	}
	if got := describe(q, &mtgv1.Answer{Text: "any"}); got != `text "any"` {
		t.Errorf("describe = %q", got)
	}
}
