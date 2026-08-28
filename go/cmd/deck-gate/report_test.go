package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// A bar that has never failed may not be a bar at all. The block bar is
// proved: deck gate run 1 of 2026-08-27 failed on it. The other two had
// never fired once, so these cases fire them on purpose (D-234).

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

func verdictOf(t *testing.T, rs []result) string {
	t.Helper()
	var b bytes.Buffer
	idx := cards.NewIndex(nil, nil, nil, time.Time{})
	report(&b, rs, llm.NewAccumulator(nil), idx, time.Second)
	line := b.String()
	switch {
	case strings.Contains(line, "Verdict: PASS"):
		return "PASS"
	case strings.Contains(line, "Verdict: FAIL"):
		return "FAIL"
	}
	t.Fatalf("no verdict in the document: %s", line)
	return ""
}

// TestCleanRunPasses is the control. Without it a bar that always fails
// would look like a working bar.
func TestCleanRunPasses(t *testing.T) {
	if got := verdictOf(t, []result{goodResult()}); got != "PASS" {
		t.Errorf("verdict = %s, want PASS on a clean run", got)
	}
}

// TestBlockFindingFailsTheRun is the bar run 1 proved.
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

// TestInventedNameFailsTheRun fires the bar that had never fired. A name
// that misses twice reaches the user as a note, and no invented name may
// reach the user.
func TestInventedNameFailsTheRun(t *testing.T) {
	r := goodResult()
	r.notes = []string{`I could not place "Craterhoof Behemoth", so it is not in the deck.`}
	if got := verdictOf(t, []result{r}); got != "FAIL" {
		t.Errorf("verdict = %s, want FAIL on an invented name", got)
	}
}

// TestFalseRuleFailsTheRun fires the F-26 bar, which had never fired.
// This is the failure that reached a user twice and passed both the gate
// and the deterministic linter.
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

// TestATrueRuleDoesNotFailTheRun keeps the bar from being too strict. The
// judge rated two summaries of run 4 as stating a rule, and both were
// true. A true statement must not fail the gate.
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
			var b bytes.Buffer
			report(&b, []result{r}, llm.NewAccumulator(nil), cards.NewIndex(nil, nil, nil, time.Time{}), time.Second)
			if got := verdictOf(t, []result{r}); got != tc.want {
				t.Errorf("verdict = %s, want %s", got, tc.want)
			}
			if !strings.Contains(b.String(), tc.row) {
				t.Errorf("the report lacks %q:\n%s", tc.row, b.String())
			}
			if tc.judgeErr != nil && !strings.Contains(b.String(), "JUDGE ERROR: the provider timed out") {
				t.Errorf("the deck does not name its judge error:\n%s", b.String())
			}
		})
	}
}
