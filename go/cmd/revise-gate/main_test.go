package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

func render(t *testing.T, outcomes []outcome) (bool, string) {
	t.Helper()
	var b bytes.Buffer
	pass := report(&b, outcomes, llm.NewAccumulator(nil), cards.NewIndex(nil, nil, nil, time.Time{}), time.Second, evalrun.New("revise", "test"))
	doc := b.String()
	want := "Verdict: FAIL"
	if pass {
		want = "Verdict: PASS"
	}
	if !strings.Contains(doc, want) {
		t.Fatalf("report returned %v and the document says otherwise:\n%s", pass, doc)
	}
	return pass, doc
}

func good() outcome {
	return outcome{
		base: base{ID: 1, Name: "a base"}, rev: revision{ID: 1, Message: "swap a card"},
		deck: &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{Name: "Sol Ring", Count: 1}}}, kept: 1,
	}
}

// TestReportVerdict covers the verdict the exit code reads: a clean run
// passes, a failure or an error fails, and an empty run fails.
func TestReportVerdict(t *testing.T) {
	if pass, _ := render(t, []outcome{good()}); !pass {
		t.Error("a clean run failed")
	}
	failed := good()
	failed.failures = []string{"kept 50% of the untouched cards, the bar is 80%"}
	if pass, doc := render(t, []outcome{good(), failed}); pass || !strings.Contains(doc, "1 of 2 turns") {
		t.Errorf("a failure passed:\n%s", doc)
	}
	errored := good()
	errored.err = errors.New("base build: the provider timed out")
	if pass, _ := render(t, []outcome{errored}); pass {
		t.Error("an error passed")
	}
	if pass, _ := render(t, nil); pass {
		t.Error("an empty run passed")
	}
}

// TestCostSaysUnpriced covers M-1: a nil cost is not $0.
func TestCostSaysUnpriced(t *testing.T) {
	if _, doc := render(t, []outcome{good()}); !strings.Contains(doc, "0 calls, unpriced.") {
		t.Errorf("the report prints a nil cost as a number:\n%s", doc)
	}
}

// TestSelectBasesRefusesAnEmptyMatch covers the -only flag.
func TestSelectBasesRefusesAnEmptyMatch(t *testing.T) {
	all := []base{{ID: 1}, {ID: 2}, {ID: 3}}
	if got, err := selectBases(all, ""); err != nil || len(got) != 3 {
		t.Errorf("empty -only = %d bases, %v; want all", len(got), err)
	}
	if got, err := selectBases(all, "2"); err != nil || len(got) != 1 || got[0].ID != 2 {
		t.Errorf("-only 2 = %v, %v", got, err)
	}
	if _, err := selectBases(all, "9"); err == nil {
		t.Error("-only 9 matched nothing and was accepted")
	}
}

// TestReportMarksTheAnswerTurn is D-448: the second turn of an unclear
// request reads as the answer in the table and in its section.
func TestReportMarksTheAnswerTurn(t *testing.T) {
	answered := good()
	answered.answered = true
	answered.message = "Q: Which lands?\nA: A mix."
	_, doc := render(t, []outcome{good(), answered})
	for _, want := range []string{"| 1 | 1 | (the answer) Q: Which lands? A: A mix. |", "## Revision 1, the answer, base 1: a base"} {
		if !strings.Contains(doc, want) {
			t.Errorf("doc lacks %q:\n%s", want, doc)
		}
	}
}
