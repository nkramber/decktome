package main

import (
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/quality"
)

const judgedDoc = `# PR-14B quality judge lane

Run date: 2026-09-10. Card snapshot: 2026-09-04. Source: /somewhere/docs/reference/pr8-deck-gate-run16.md.

## Summary

| Measure | Value |
|---|---|
| Decks | 3 |

## Decks

| # | Deck | Format | Grade | Document | Judge | Agree |
|---|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | Commander | bad | bad | typical | no |
| 2 | Modern burn, casual | Modern | baseline | bad | bad | no |
| 3 | a deck the judge failed on | Standard | bad | bad | error |  |
`

func TestReadJudged(t *testing.T) {
	source, answers, err := readJudged(strings.NewReader(judgedDoc))
	if err != nil {
		t.Fatal(err)
	}
	if source != "/somewhere/docs/reference/pr8-deck-gate-run16.md" {
		t.Errorf("source = %q", source)
	}
	if len(answers) != 2 || answers[1] != (judgeAnswer{title: "lifegain Commander, any card", tier: meta.TierTypical}) || answers[2].tier != meta.TierBad {
		t.Errorf("answers = %+v, want two, and the judge error out", answers)
	}
	if _, _, err := readJudged(strings.NewReader("# x\n\n| 1 | a | Commander | bad | bad | good | no |\n")); err == nil {
		t.Errorf("a judge document with no source must fail the read")
	}
}

// TestSameDecks: the agreement compares a grade with a judge answer about
// the same deck, so another document or another title fails.
func TestSameDecks(t *testing.T) {
	decks := []judged{{id: 1, title: "lifegain Commander, any card"}}
	answers := map[int]judgeAnswer{1: {title: "lifegain Commander, any card", tier: meta.TierTypical}}
	if err := sameDecks("/a/pr8-deck-gate-run16.md", "/b/pr8-deck-gate-run16.md", decks, answers); err != nil {
		t.Errorf("the same document in two folders: %v", err)
	}
	if err := sameDecks("/a/pr8-deck-gate-run18.md", "/b/pr8-deck-gate-run16.md", decks, answers); err == nil {
		t.Errorf("another deck gate document must fail")
	}
	answers[1] = judgeAnswer{title: "another deck", tier: meta.TierGood}
	if err := sameDecks("/a/pr8-deck-gate-run16.md", "/b/pr8-deck-gate-run16.md", decks, answers); err == nil {
		t.Errorf("another title must fail")
	}
}

func TestReaderCounts(t *testing.T) {
	r := &readerRead{
		decks:   []judged{{id: 1, modelGrade: meta.TierBad}, {id: 2, modelGrade: meta.TierBad}, {id: 3, modelGrade: meta.TierTypical}},
		answers: map[int]judgeAnswer{1: {tier: meta.TierBad}, 3: {tier: meta.TierGood}},
	}
	bad, judgedCount, agreed := r.counts()
	if bad != 2 || judgedCount != 2 || agreed != 1 {
		t.Errorf("counts = %d bad, %d judged, %d agreed, want 2, 2, 1", bad, judgedCount, agreed)
	}
}

// TestBrokenCopiesGradedBad: the count of D-674 reads the bad row of the
// confusion table.
func TestBrokenCopiesGradedBad(t *testing.T) {
	h := quality.Holdout{Confusion: [][]int{{40, 5, 3, 2, 0}, {1, 9, 0, 0, 0}}}
	if graded, copies := brokenCopiesGradedBad(h); graded != 40 || copies != 50 {
		t.Errorf("broken copies graded bad = %d of %d, want 40 of 50", graded, copies)
	}
	if graded, copies := brokenCopiesGradedBad(quality.Holdout{}); graded != 0 || copies != 0 {
		t.Errorf("an empty holdout = %d of %d, want 0 of 0", graded, copies)
	}
}
