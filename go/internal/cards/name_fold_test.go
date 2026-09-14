package cards

import (
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The D-716 tests of the index. A reader typed "Grima, Saruman's
// Footman", and the lookup missed, because the card spells it "Gríma".

// foldIndex holds two real card names with a special mark, an invented
// split card with accented faces, and the extra cards a test adds.
func foldIndex(extra ...*mtgv1.Card) *Index {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	list := []*mtgv1.Card{
		{OracleId: "o-footman", Name: "Gríma, Saruman's Footman", Legalities: legal},
		{OracleId: "o-raton", Name: "Ratonhnhaké꞉ton", Legalities: legal},
		{OracleId: "o-split", Name: "Lórien Gate // Lórien Path", Legalities: legal,
			Faces: []*mtgv1.CardFace{{Name: "Lórien Gate"}, {Name: "Lórien Path"}}},
	}
	for _, c := range extra {
		c.Legalities = legal
	}
	return NewIndex(append(list, extra...), nil, nil, time.Time{})
}

// TestByNameFoldsWhatAReaderTypes: a name without its accent or its
// special mark finds the card, and so does a face name.
func TestByNameFoldsWhatAReaderTypes(t *testing.T) {
	idx := foldIndex()
	for _, tc := range []struct{ look, want string }{
		{"Gríma, Saruman's Footman", "o-footman"},
		{"Grima, Saruman's Footman", "o-footman"},
		{"  grima, saruman’s footman ", "o-footman"},
		{"Ratonhnhake:ton", "o-raton"},
		{"Lorien Path", "o-split"},
	} {
		c, ok := idx.ByName(tc.look)
		if !ok || c.GetOracleId() != tc.want {
			t.Errorf("ByName(%q) = %q, %v, want %s", tc.look, c.GetOracleId(), ok, tc.want)
		}
	}
	// The fold reads marks and not words, so a short name finds nothing.
	if _, ok := idx.ByName("Grima"); ok {
		t.Error("a short name found a card")
	}
}

// TestByNameReadsTheExactNameBeforeTheFold: two invented card names that
// differ by an accent alone keep their exact names, and a folded name
// that fits both finds nothing (guardrail 4).
func TestByNameReadsTheExactNameBeforeTheFold(t *testing.T) {
	idx := foldIndex(
		&mtgv1.Card{OracleId: "o-plain", Name: "Cafe Sign"},
		&mtgv1.Card{OracleId: "o-accent", Name: "Café Sign"},
	)
	for _, tc := range []struct{ look, want string }{
		{"Cafe Sign", "o-plain"},
		{"CAFÉ SIGN", "o-accent"},
		{"Cafè Sign", ""},
	} {
		c, ok := idx.ByName(tc.look)
		if ok != (tc.want != "") || (ok && c.GetOracleId() != tc.want) {
			t.Errorf("ByName(%q) = %q, %v, want %q", tc.look, c.GetOracleId(), ok, tc.want)
		}
	}
	if got := idx.Collisions().FoldNames; got != 1 {
		t.Errorf("fold name collisions = %d, want 1", got)
	}
}
