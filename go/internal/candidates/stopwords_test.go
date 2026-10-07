package candidates

import (
	"strings"
	"testing"
)

// fillerThemes are the theme phrasings of D-1188, each with the theme words
// that must survive. The first is the prompt of a live eval, and the rest
// say the same thing in other words. Before D-1188 each one kept a word
// that matched no card, and the theme row asked about it.
var fillerThemes = map[string]string{
	"planeswalkers and other themes from this set":             "planeswalkers",
	"planeswalkers and other mechanics from this set":          "planeswalkers",
	"planeswalkers plus whatever else the set does well":       "planeswalkers",
	"planeswalkers and the main mechanics of the set":          "planeswalkers",
	"planeswalkers along with similar themes":                  "planeswalkers",
	"planeswalkers and related archetypes from this expansion": "planeswalkers",
	"planeswalkers and the other strategies of the block":      "planeswalkers",
	"planeswalkers and additional synergies":                   "planeswalkers",
	"dragons and anything else that fits":                      "dragons",
	"dragons and whatever works with them":                     "dragons",
	"elves and similar tribal stuff":                           "elves tribal",
	"tokens and other things":                                  "tokens",
	"lifegain etc":                                             "lifegain",
	"zombies, maybe some graveyard things too":                 "zombies graveyard",
	"vampires and also lifegain if possible":                   "vampires lifegain",
	"artifacts and related synergies":                          "artifacts",
	"goblins mostly":                                           "goblins",
	"burn, sideboard":                                          "burn",
}

func TestFillerLeavesTheTheme(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	for theme, want := range fillerThemes {
		if got := strings.Join(b.themes.words(theme), " "); got != want {
			t.Errorf("words(%q) = %q, want %q", theme, got, want)
		}
	}
}

// TestStopWordsNameNoTheme guards the lists of D-1188: a stop word that
// names a theme row or a card type would drop a theme the reader asked
// for.
func TestStopWordsNameNoTheme(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	for _, list := range []map[string]bool{englishStopWords, requestFiller} {
		for w := range list {
			if row, ok := b.themes.rowOf(w); ok {
				t.Errorf("stop word %q finds the row %q", w, row)
			}
			if cardTypes[title(singular(w))] {
				t.Errorf("stop word %q names a card type", w)
			}
		}
	}
}
