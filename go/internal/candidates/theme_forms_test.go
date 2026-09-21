package candidates

import (
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestThemeWordsFindTheirRow is D-724. The match read a row by the exact
// word, so "milling" matched no card, and the Gríma shortlist held staple
// roles alone (F-142). Each word of the first list finds its row through
// the row name, an alias, or a base form. No word of the second list finds
// a row. A singular creature type is one of them: it keeps the generic rule,
// whose text needle finds type cards the type row does not hold (D-731).
func TestThemeWordsFindTheirRow(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	rows := map[string]string{
		"mill": "mill", "milling": "mill", "blinking": "blink", "flickering": "flicker",
		"reanimate": "reanimator", "reanimating": "reanimator", "reanimation": "reanimator",
		"sacrificing": "sacrifice", "aristocrat": "aristocrats", "steal": "theft", "stealing": "theft",
		"life-gain": "lifegain", "gaining-life": "lifegain", "wheeling": "wheel", "wheels": "wheel",
		"ramping": "ramp", "storming": "storm", "controlling": "control", "counterspells": "control",
		"counter-spell": "control", "proliferating": "proliferate", "clones": "copy", "flyers": "flying",
		"fliers": "flying", "combos": "combo", "token": "tokens", "treasures": "treasure", "spell": "spells",
		"extra-turn": "extra-turns",
		"self-mill":  "graveyard", "prison": "stax", "weenie": "aggro", "aggressive": "aggro",
		"superfriends": "superfriends", "planeswalkers": "superfriends", "going-wide": "go-wide",
		"spell-slinger": "spellslinger", "land-destruction": "land-destruction",
		"heroes": "heroes", "superhero": "heroes", "superheroes": "heroes", "avengers": "heroes",
	}
	for word, want := range rows {
		if got, ok := b.themes.rowOf(word); !ok || got != want {
			t.Errorf("rowOf(%q) = %q, %v, want %q", word, got, ok, want)
		}
	}
	for _, word := range []string{"opponent", "zzzz", "tempo", "king", "string", "draining", "hobbits", "turtles", "elf", "zombie", "dragon", "hero"} {
		if got, ok := b.themes.rowOf(word); ok {
			t.Errorf("rowOf(%q) = %q, want no row", word, got)
		}
	}
}

// TestTwoWordsJoinThroughAForm is D-724. "extra turn" names no row, and
// its hyphenated form is the singular of extra-turns. "land destruction"
// must join too, or "land" alone finds the lands row. Two words that find
// one row are one word of the match.
func TestTwoWordsJoinThroughAForm(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string]string{
		"extra turn":             "extra-turn",
		"gaining life":           "gaining-life",
		"life gain":              "life-gain",
		"land destruction":       "land-destruction",
		"counter spells":         "counter-spells",
		"opponent milling cards": "opponent milling",
	}
	for theme, want := range cases {
		if got := strings.Join(b.themes.words(theme), " "); got != want {
			t.Errorf("words(%q) = %q, want %q", theme, got, want)
		}
	}
	idx := fixture(t, testCards())
	if m := b.themes.match("mill milling", idx.Tags()); len(m.Words) != 1 {
		t.Errorf("two words of one row read as %v, want one word", m.Words)
	}
}

// TestThemeFormsMatchCards is F-142 on a built pool. Before D-724 every
// word of each theme came back unmatched, and the superfriends theme found
// no planeswalker (D-729).
func TestThemeFormsMatchCards(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	R, U := mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_U
	idx := fixture(t, append(testCards(),
		tc{id: "millstone", name: "Millstone", typeLine: "Artifact", text: "{2}, {T}: Target player mills two cards.", mv: 2, rank: 800, tags: []string{"mill-opponent"}},
		tc{id: "ephemerate", name: "Ephemerate", typeLine: "Instant", text: "Exile target creature you control, then return it to the battlefield under its owner's control.", identity: []mtgv1.Color{W}, mv: 1, rank: 90, tags: []string{"flicker-creature"}},
		tc{id: "treason", name: "Act of Treason", typeLine: "Sorcery", text: "Gain control of target creature until end of turn.", identity: []mtgv1.Color{R}, mv: 3, rank: 700, tags: []string{"theft"}},
		tc{id: "stonerain", name: "Stone Rain", typeLine: "Sorcery", text: "Destroy target land.", identity: []mtgv1.Color{R}, mv: 3, rank: 900, tags: []string{"removal-land"}},
		tc{id: "jace", name: "Jace Beleren", typeLine: "Legendary Planeswalker — Jace", text: "+2: Each player draws a card.", identity: []mtgv1.Color{U}, mv: 3, rank: 600},
	))
	for _, theme := range []string{"milling", "blinking", "stealing", "gaining life", "land destruction", "superfriends"} {
		list, err := b.Build(idx, Request{Format: cmdr, Theme: theme})
		if err != nil {
			t.Fatal(err)
		}
		if len(list.Theme.Unmatched) > 0 {
			t.Errorf("%q left words unmatched: %s", theme, list.Theme.Describe())
		}
		if theme != "superfriends" {
			continue
		}
		jace, ok := find(list.Candidates, "Jace Beleren")
		if !ok || !slices.Contains(jace.Signals, "type:Planeswalker") || jace.Role != mtgv1.CardRole_CARD_ROLE_THREAT {
			t.Errorf("superfriends must list Jace as a threat by its type: %+v, %v", jace, ok)
		}
	}
}

// TestParseThemesRefusesABadRow holds the load checks of D-724 and D-729.
// A bad alias or type would find a wrong row, or no card, and say nothing.
func TestParseThemesRefusesABadRow(t *testing.T) {
	cases := map[string]string{
		"an alias that names a row":   `{"verified_at":"x","themes":{"mill":{"aliases":["blink"]},"blink":{}}}`,
		"an alias of two rows":        `{"verified_at":"x","themes":{"mill":{"aliases":["grind"]},"blink":{"aliases":["grind"]}}}`,
		"an alias of two words":       `{"verified_at":"x","themes":{"mill":{"aliases":["mill cards"]}}}`,
		"an alias with a capital":     `{"verified_at":"x","themes":{"mill":{"aliases":["Grind"]}}}`,
		"a type that is no card type": `{"verified_at":"x","themes":{"superfriends":{"types":["Walker"]}}}`,
		"two rows of one singular":    `{"verified_at":"x","themes":{"elves":{},"elfs":{}}}`,
	}
	for name, raw := range cases {
		if _, err := parseThemes([]byte(raw)); err == nil {
			t.Errorf("%s: parseThemes accepted %s", name, raw)
		}
	}
	if _, err := parseThemes(themesJSON); err != nil {
		t.Errorf("the embedded table: %v", err)
	}
}
