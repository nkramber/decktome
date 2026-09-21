package profile

import (
	"context"
	"errors"
	"testing"

	"github.com/nkramber/decktome/go/internal/spellbook"
)

// countingClassifier counts the calls, so a test proves that Floor asks
// Commander Spellbook one time.
type countingClassifier struct {
	fakeClassifier
	calls int
}

func (c *countingClassifier) EstimateBracket(ctx context.Context, commanders, main []string) (*spellbook.Result, error) {
	c.calls++
	return c.fakeClassifier.EstimateBracket(ctx, commanders, main)
}

func twoCard(speed int, sure bool) *spellbook.Result {
	return &spellbook.Result{Combos: []spellbook.ClassifiedCombo{{
		Combo: spellbook.ComboRef{ID: "1", Uses: []spellbook.ComboUse{
			{Card: spellbook.CardRef{Name: "Lightning Runner"}}, {Card: spellbook.CardRef{Name: "Aetherwind Basker"}}}},
		Relevant: true, DefinitelyTwoCard: sure, ArguablyTwoCard: !sure, Speed: speed,
	}}}
}

// TestFloorReadsTheLowestBracketTheRulesAllow is F-162: a precon with a
// two-card infinite of speed 5 passes the rules of bracket 4 alone, and
// a deck with no combo passes bracket 1.
func TestFloorReadsTheLowestBracketTheRulesAllow(t *testing.T) {
	cases := []struct {
		name string
		res  *spellbook.Result
		want int32
	}{
		{"no combo", &spellbook.Result{}, 1},
		{"sure speed 5", twoCard(5, true), 4},
		{"sure speed 2", twoCard(2, true), 2},
		{"sure speed 3", twoCard(3, true), 3},
		{"near speed 3 reads 2", twoCard(3, false), 2},
		{"mass land denial", &spellbook.Result{Cards: []spellbook.ClassifiedCard{
			{Card: spellbook.CardRef{Name: "Armageddon"}, Quantity: 1, MassLandDenial: true}}}, 4},
		{"two extra-turn cards", &spellbook.Result{Cards: []spellbook.ClassifiedCard{
			{Card: spellbook.CardRef{Name: "Time Warp"}, Quantity: 1, ExtraTurn: true},
			{Card: spellbook.CardRef{Name: "Temporal Manipulation"}, Quantity: 1, ExtraTurn: true}}}, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cl := &countingClassifier{fakeClassifier: fakeClassifier{res: c.res}}
			p := newProfiler(t, cl)
			got, err := p.Floor(context.Background(), deckOf(2, shaped()...), source(testCards))
			if err != nil {
				t.Fatal(err)
			}
			if got != c.want {
				t.Errorf("floor %d, want %d", got, c.want)
			}
			if cl.calls != 1 {
				t.Errorf("%d classifier calls, want 1", cl.calls)
			}
		})
	}
}

// TestFloorRefusesAnUncheckedDeck: with no content check, the rules give
// no floor, and the calibration must not guess one.
func TestFloorRefusesAnUncheckedDeck(t *testing.T) {
	p := newProfiler(t, &fakeClassifier{err: errors.New("502")})
	if _, err := p.Floor(context.Background(), deckOf(2, shaped()...), source(testCards)); err == nil {
		t.Error("floor with a failed content check, want an error")
	}
	if _, err := newProfiler(t, nil).Floor(context.Background(), deckOf(2, shaped()...), source(testCards)); err == nil {
		t.Error("floor with no classifier, want an error")
	}
}
