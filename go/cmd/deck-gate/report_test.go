package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// A bar that has never failed may not be a bar at all, so every bar is
// fired here on purpose (D-234).

func goodResult() result {
	return result{
		prompt: prompt{ID: 1, Name: "a deck", Format: "commander", Theme: "lifegain"},
		deck: &mtgv1.Deck{
			Summary:    "a plain deck summary",
			Cards:      []*mtgv1.DeckCard{{Name: "Ajani's Welcome", Count: 1}},
			Validation: &mtgv1.ValidationResult{},
		},
		judged: &generate.Judgement{Verdict: "clean"},
	}
}

// render writes the document once and returns the verdict word with it.
func render(t *testing.T, rs []result) (string, string) {
	t.Helper()
	var b bytes.Buffer
	idx := cards.NewIndex(nil, nil, nil, time.Time{})
	pass := report(&b, rs, llm.NewAccumulator(nil), idx, time.Second)
	doc := b.String()
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	if !strings.Contains(doc, "Verdict: "+verdict) {
		t.Fatalf("report returned %v and the document says otherwise: %s", pass, doc)
	}
	return verdict, doc
}

func verdictOf(t *testing.T, rs []result) string {
	t.Helper()
	v, _ := render(t, rs)
	return v
}

// TestCleanRunPasses is the control. Without it a bar that always fails
// would look like a working bar.
func TestCleanRunPasses(t *testing.T) {
	if got := verdictOf(t, []result{goodResult()}); got != "PASS" {
		t.Errorf("verdict = %s, want PASS on a clean run", got)
	}
}

// TestBlockFindingFailsTheRun fires the block bar.
func TestBlockFindingFailsTheRun(t *testing.T) {
	r := goodResult()
	r.deck.Validation.Findings = []*mtgv1.Finding{{
		Code: "deck_size", Severity: mtgv1.Severity_SEVERITY_BLOCK,
		Message: "deck has 99 cards, the format needs exactly 100",
	}}
	if got := verdictOf(t, []result{r}); got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on a block finding", got)
	}
}

// TestInventedNameFailsTheRun fires the invented-name bar. A name that
// misses twice reaches the user as a note, and no invented name may
// reach the user.
func TestInventedNameFailsTheRun(t *testing.T) {
	r := goodResult()
	r.notes = []string{`I could not place "Craterhoof Behemoth", so it is not in the deck.`}
	if got := verdictOf(t, []result{r}); got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on an invented name", got)
	}
}

// TestFalseRuleFailsTheRun fires the F-26 bar: a summary that states a
// false rule of the game fails the gate whatever the linter says (D-229).
func TestFalseRuleFailsTheRun(t *testing.T) {
	r := goodResult()
	r.deck.Summary = "Grist, the Hunger Tide can not lead a deck, so it sits in the 99."
	r.judged = &generate.Judgement{
		Verdict: "states_a_false_rule",
		Claims: []generate.Claim{{
			Text: "Grist, the Hunger Tide can not lead a deck", Truth: "false",
			Why: "Grist is a creature card outside the battlefield and can be a commander",
		}},
	}
	if got := verdictOf(t, []result{r}); got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on a false rule of the game", got)
	}
}

// TestFalseClaimUnderASoftVerdictFailsTheRun keeps the two judge signals
// from disagreeing in the gate's favour.
func TestFalseClaimUnderASoftVerdictFailsTheRun(t *testing.T) {
	r := goodResult()
	r.judged = &generate.Judgement{
		Verdict: "states_a_rule",
		Claims:  []generate.Claim{{Text: "Sol Ring is banned", Truth: "false", Why: "it is legal"}},
	}
	if got := verdictOf(t, []result{r}); got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL: a false claim fails whatever the verdict says", got)
	}
}

// TestATrueRuleDoesNotFailTheRun keeps the bar from being too strict. A
// true statement of a rule must not fail the gate.
func TestATrueRuleDoesNotFailTheRun(t *testing.T) {
	r := goodResult()
	r.judged = &generate.Judgement{
		Verdict: "states_a_rule",
		Claims:  []generate.Claim{{Text: "Urza leads a dense artifact shell", Truth: "true", Why: "Urza is legendary"}},
	}
	if got := verdictOf(t, []result{r}); got != "PASS" {
		t.Errorf("verdict = %s, want PASS on a true statement", got)
	}
}

// TestAnErrorFailsTheRun covers the prompt that never produced a deck. A
// gate can not pass with an error.
func TestAnErrorFailsTheRun(t *testing.T) {
	r := goodResult()
	r.err = errTest
	r.deck = nil
	if got := verdictOf(t, []result{r}); got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on an error", got)
	}
}

var errTest = &testError{}

type testError struct{}

func (*testError) Error() string { return "the prompt failed" }

// TestJudgeErrorFailsTheRun is T-17. A deck the judge could not read has
// no verdict on F-26, so it can not pass that bar, and the row says so.
func TestJudgeErrorFailsTheRun(t *testing.T) {
	cases := []struct {
		name     string
		judgeErr error
		want     string
		row      string
	}{
		{"no judge error", nil, "PASS", "| Judge errors | 0 |"},
		{"a judge error", fmt.Errorf("the provider timed out"), "FAIL", "| Judge errors | 1 |"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := goodResult()
			if tc.judgeErr != nil {
				r.judged, r.judgeErr = nil, tc.judgeErr
			}
			got, doc := render(t, []result{r})
			if got != tc.want {
				t.Errorf("verdict = %s, want %s", got, tc.want)
			}
			if !strings.Contains(doc, tc.row) {
				t.Errorf("the report lacks %q:\n%s", tc.row, doc)
			}
			if tc.judgeErr != nil && !strings.Contains(doc, "JUDGE ERROR: the provider timed out") {
				t.Errorf("the deck does not name its judge error:\n%s", doc)
			}
		})
	}
}

// TestEmptyRunFailsTheRun covers a -only list that names no prompt. Zero
// of zero is not a pass.
func TestEmptyRunFailsTheRun(t *testing.T) {
	got, doc := render(t, nil)
	if got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on an empty run", got)
	}
	if !strings.Contains(doc, "no prompt") {
		t.Errorf("the document does not say the run was empty:\n%s", doc)
	}
}

// TestSelectPromptsRefusesAnEmptyMatch covers the -only flag.
func TestSelectPromptsRefusesAnEmptyMatch(t *testing.T) {
	all := []prompt{{ID: 1, Name: "a"}, {ID: 2, Name: "b"}, {ID: 3, Name: "c"}}
	got, err := selectPrompts(all, "")
	if err != nil || len(got) != 3 {
		t.Errorf("empty -only = %d prompts, %v; want all", len(got), err)
	}
	got, err = selectPrompts(all, "3, 1")
	if err != nil || len(got) != 2 || got[0].ID != 1 || got[1].ID != 3 {
		t.Errorf("-only 3,1 = %v, %v", got, err)
	}
	if _, err := selectPrompts(all, "9"); err == nil {
		t.Error("-only 9 matched nothing and was accepted")
	}
	if _, err := selectPrompts(all, "x"); err == nil {
		t.Error("-only x was accepted")
	}
}

// TestEmptySummaryIsAJudgeError covers F-26 on a deck with no summary.
// The judge has nothing to read, so the deck has no verdict on that bar.
func TestEmptySummaryIsAJudgeError(t *testing.T) {
	j, err := judge(context.Background(), nil, "a deck", &mtgv1.Deck{}, nil)
	if err == nil || j != nil {
		t.Fatalf("judge = %v, %v; want the empty-summary error", j, err)
	}
	r := goodResult()
	r.deck.Summary = ""
	r.judged, r.judgeErr = nil, err
	got, doc := render(t, []result{r})
	if got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on an empty summary", got)
	}
	if !strings.Contains(doc, "| Judge errors | 1 |") {
		t.Errorf("the report does not count the empty summary as a judge error:\n%s", doc)
	}
}

// TestCostRowSaysUnpriced covers M-1: a nil cost is not $0.
func TestCostRowSaysUnpriced(t *testing.T) {
	_, doc := render(t, []result{goodResult()})
	if !strings.Contains(doc, "| Cost | unpriced |") {
		t.Errorf("the report prints a nil cost as a number:\n%s", doc)
	}
}

// TestExcludedPreconCardsAreCounted: the PR-24 block names the products
// and counts the deck cards that belong to them. Such a card is a block
// on the deck, so the verdict reads it (D-408). A run with no exclusion
// prompt writes no block.
func TestExcludedPreconCardsAreCounted(t *testing.T) {
	slipped := goodResult()
	slipped.products = []string{"Avengers Assemble"}
	slipped.excluded = map[string]bool{"o-avenge": true, "o-jarvis": true}
	slipped.spare = 3
	slipped.deck.Cards = []*mtgv1.DeckCard{{Name: "Avenge", Count: 1, OracleId: "o-avenge"}, {Name: "Sol Ring", Count: 1, OracleId: "o-sol"}}
	slipped.deck.Validation.Findings = []*mtgv1.Finding{{
		Code: rules.CodeExcludedPrecon, Severity: mtgv1.Severity_SEVERITY_BLOCK, Message: "Avenge is a card of Avengers Assemble",
	}}
	verdict, doc := render(t, []result{slipped})
	if verdict != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on a card of an excluded precon", verdict)
	}
	if want := "| 1 | a deck | Avengers Assemble | 2 | 3 | 1 | 1 |"; !strings.Contains(doc, want) {
		t.Errorf("the block lacks %q:\n%s", want, doc)
	}

	clean := goodResult()
	clean.products = []string{"Avengers Assemble"}
	clean.excluded = map[string]bool{"o-avenge": true}
	clean.deck.Cards = []*mtgv1.DeckCard{{Name: "Sol Ring", Count: 1, OracleId: "o-sol"}}
	verdict, doc = render(t, []result{clean})
	if verdict != "PASS" {
		t.Errorf("verdict = %s, want PASS when no excluded card is in the deck", verdict)
	}
	if want := "| 1 | a deck | Avengers Assemble | 1 | 0 | 0 | 0 |"; !strings.Contains(doc, want) {
		t.Errorf("the block lacks %q:\n%s", want, doc)
	}

	if _, doc := render(t, []result{goodResult()}); strings.Contains(doc, "## The precon exclusion") {
		t.Error("a run with no exclusion prompt wrote the block")
	}
}
