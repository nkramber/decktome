package candidates

import (
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
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

// TestInflectedFillerLeavesTheMatch is F-230. The stop lists hold "make",
// "build", and "want", and the theme row asked about "making" in "making
// lots of treasure tokens and using them to win". An -s, -ing, or -ed form
// of a stop word that fires on no card is filler, so it is no theme word
// and the theme row never names it. A form that a card holds stays, such
// as "blocking" of the filler word "block".
func TestInflectedFillerLeavesTheMatch(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, append(testCards(),
		tc{id: "wall", name: "Wall of Denial", typeLine: "Creature — Wall", text: "Whenever this creature blocks, it gets +0/+2 until end of turn for each blocking creature you control.", identity: []mtgv1.Color{W}, mv: 3, rank: 900},
	))
	cases := map[string][]string{
		"making lots of lifegain":                 {"lifegain"},
		"building around cats":                    {"cats"},
		"a deck that wanted lots of lifegain":     {"lifegain"},
		"lifegain that makes and builds counters": {"lifegain", "counters"},
		"needing, liking, and focusing on cats":   {"cats"},
		"blocking lifegain":                       {"blocking", "lifegain"},
		"making":                                  nil,
	}
	for theme, want := range cases {
		list, err := b.Build(idx, Request{Format: cmdr, Theme: theme, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD})
		if err != nil {
			t.Fatal(err)
		}
		if len(list.Theme.Unmatched) > 0 {
			t.Errorf("%q leaves %v unmatched, and the theme row would ask", theme, list.Theme.Unmatched)
		}
		if !slices.Equal(list.Theme.Words, want) {
			t.Errorf("%q has the theme words %v, want %v", theme, list.Theme.Words, want)
		}
	}
	// A word that is no form of a stop word stays unmatched (D-1116).
	list, err := b.Build(idx, Request{Format: cmdr, Theme: "making zzzz lifegain", PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD})
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(list.Theme.Unmatched, []string{"zzzz"}) {
		t.Errorf("unmatched %v, want [zzzz]", list.Theme.Unmatched)
	}
}

// TestFillerFormReadsTheInflections guards the forms of F-230.
func TestFillerFormReadsTheInflections(t *testing.T) {
	for _, w := range []string{"making", "makes", "building", "builds", "wanted", "wanting", "needed", "liked", "giving", "gives", "focusing", "focuses", "fitting", "blocking"} {
		if !fillerForm(w) {
			t.Errorf("fillerForm(%q) = false, want true", w)
		}
	}
	for _, w := range []string{"make", "treasure", "tokens", "milling", "sacrificed", "zzzz", "used", "red", "wishing"} {
		if fillerForm(w) {
			t.Errorf("fillerForm(%q) = true, want false", w)
		}
	}
}
