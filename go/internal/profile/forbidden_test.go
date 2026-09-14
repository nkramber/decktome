package profile

import (
	"context"
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/spellbook"
)

// TestReadForbiddenNamesWhatTheBracketForbids is PR-45a: the content check
// names each combo, mass land denial card, and extra-turn card that the
// bracket forbids, by the same rules as its warnings.
func TestReadForbiddenNamesWhatTheBracketForbids(t *testing.T) {
	res := &spellbook.Result{
		Cards: []spellbook.ClassifiedCard{
			{Card: spellbook.CardRef{Name: "Armageddon"}, Quantity: 1, MassLandDenial: true},
			{Card: spellbook.CardRef{Name: "Time Warp"}, Quantity: 1, ExtraTurn: true},
			{Card: spellbook.CardRef{Name: "Temporal Manipulation"}, Quantity: 1, ExtraTurn: true},
		},
		Combos: []spellbook.ClassifiedCombo{{
			Combo: spellbook.ComboRef{ID: "1", Uses: []spellbook.ComboUse{
				{Card: spellbook.CardRef{Name: "Demonic Consultation"}}, {Card: spellbook.CardRef{Name: "Thassa's Oracle"}}}},
			Relevant: true, DefinitelyTwoCard: true, Speed: 4,
		}},
	}
	p := newProfiler(t, &fakeClassifier{res: res})
	_, _, f := p.ReadForbidden(context.Background(), deckOf(3, shaped()...), source(testCards))
	if len(f.Combos) != 1 || strings.Join(f.Combos[0], " + ") != "Demonic Consultation + Thassa's Oracle" {
		t.Errorf("combos %v", f.Combos)
	}
	if strings.Join(f.MassLandDenial, ",") != "Armageddon" {
		t.Errorf("mass land denial %v", f.MassLandDenial)
	}
	if len(f.ExtraTurns) != 2 || f.ExtraTurnCap != 1 {
		t.Errorf("extra turns %v, cap %d", f.ExtraTurns, f.ExtraTurnCap)
	}

	// Bracket 4 forbids none of it.
	p4 := newProfiler(t, &fakeClassifier{res: res})
	if _, _, f4 := p4.ReadForbidden(context.Background(), deckOf(4, shaped()...), source(testCards)); !f4.Empty() {
		t.Errorf("bracket 4 forbids %+v", f4)
	}

	// A near two-card combo at speed 4 reads as 3, which bracket 3 allows.
	near := &spellbook.Result{Combos: []spellbook.ClassifiedCombo{{
		Combo: spellbook.ComboRef{ID: "2", Uses: []spellbook.ComboUse{
			{Card: spellbook.CardRef{Name: "Sanguine Bond"}}, {Card: spellbook.CardRef{Name: "Exquisite Blood"}}}},
		Relevant: true, ArguablyTwoCard: true, Speed: 4,
	}}}
	pn := newProfiler(t, &fakeClassifier{res: near})
	if _, _, fn := pn.ReadForbidden(context.Background(), deckOf(3, shaped()...), source(testCards)); !fn.Empty() {
		t.Errorf("bracket 3 allows a near combo at speed 4: %+v", fn)
	}
}
