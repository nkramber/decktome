package candidates

import (
	"fmt"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// TestPruneNoisyDropsANeedleMostCardsHold is D-411. A generic text
// needle that sits in the text of more than a tenth of the cards ranks
// the pool on chance, so it goes. A needle few cards hold stays.
func TestPruneNoisyDropsANeedleMostCardsHold(t *testing.T) {
	var all []*mtgv1.Card
	for i := 0; i < 20; i++ {
		text := "Whenever a creature enters, you gain 1 life."
		if i == 0 {
			text = "Landfall - Whenever a land you control enters, draw a card."
		}
		all = append(all, &mtgv1.Card{OracleId: fmt.Sprint("o", i), OracleText: text})
	}
	m := ThemeMatch{Text: []string{"whenever", "landfall"}, generic: []string{"whenever", "landfall"}}
	m.pruneNoisy(all)
	if len(m.Text) != 1 || m.Text[0] != "landfall" {
		t.Errorf("text needles = %v, want landfall alone", m.Text)
	}
}

// TestRequestWordsNameNoTheme: the words of "build me the best deck you
// can" are stop words, so the request reaches the unthemed pool and the
// offer ranks on popularity (D-411).
func TestRequestWordsNameNoTheme(t *testing.T) {
	tbl := &themeTable{Themes: map[string]themeRow{}}
	for _, theme := range []string{"Build me the best deck you can", "the best possible deck you can", "make something really powerful"} {
		if got := tbl.words(theme); len(got) != 0 {
			t.Errorf("%q gave theme words %v, want none", theme, got)
		}
	}
	if got := tbl.words("the best landfall deck you can"); len(got) != 1 || got[0] != "landfall" {
		t.Errorf("a real theme word must survive: %v", got)
	}
}
