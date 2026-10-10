package candidates

import (
	"regexp"
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// castTrigger finds the kind of spell that a cast trigger counts (F-229).
// "Whenever you cast a legendary spell" gives "legendary", and "Whenever
// you cast your first instant or sorcery spell each turn" gives "instant
// or sorcery". The lazy group stops at the first "spell", and a phrase
// with a word that names no kind of spell gives no class.
var castTrigger = regexp.MustCompile(`whenever you cast (?:a|an|another|your first|your second) ([a-z][a-z ,'-]*?) spells?\b`)

// spellTest reads one word of a cast trigger on a card.
type spellTest func(c *mtgv1.Card) bool

// SpellClass is one kind of spell that a cast trigger of a commander
// counts (F-229). A card fits the class when it fits one alternative, and
// it fits an alternative when each test of it holds. "Aura, Equipment, or
// Vehicle" is three alternatives of one test, and "Dragon creature" is one
// alternative of two tests.
type SpellClass struct {
	Phrase string
	alts   [][]spellTest
}

// fits reports whether a card is a spell of the class. A land is never
// cast, so it never fits.
func (s SpellClass) fits(c *mtgv1.Card) bool {
	if slices.Contains(c.GetCardTypes(), "Land") {
		return false
	}
	for _, alt := range s.alts {
		all := true
		for _, test := range alt {
			if !test(c) {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}
	return false
}

// commanderSpells reads the cast triggers of each commander (F-229). The
// spells that a trigger counts are the theme of the commander, and a
// reader who gives no theme word for them still wants them: Jodah, the
// Unifier counts legendary spells, and its shortlist held 20 of the 570
// owned legendary cards before this rule.
func commanderSpells(commanders []*mtgv1.Card, idx *cards.Index) []SpellClass {
	var set map[string]bool
	isSubtype := func(s string) bool {
		if set == nil {
			set = map[string]bool{}
			for _, c := range idx.All() {
				for _, st := range c.GetSubtypes() {
					set[st] = true
				}
			}
		}
		return set[s]
	}
	var out []SpellClass
	for _, c := range commanders {
		for _, cl := range castSpells(c.GetOracleText(), isSubtype) {
			if !slices.ContainsFunc(out, func(o SpellClass) bool { return o.Phrase == cl.Phrase }) {
				out = append(out, cl)
			}
		}
	}
	return out
}

// castSpells reads each cast trigger of one rules text. isSubtype says
// whether a word, in title case, is a subtype that a card holds.
func castSpells(text string, isSubtype func(string) bool) []SpellClass {
	var out []SpellClass
	for _, m := range castTrigger.FindAllStringSubmatch(strings.ToLower(text), -1) {
		if cl, ok := parseSpellClass(m[1], isSubtype); ok {
			out = append(out, cl)
		}
	}
	return out
}

// parseSpellClass turns the words of one cast trigger into a class. Each
// word must name a kind of spell, or the phrase is no class: "kicked"
// and "modal" name a way to cast a spell, and the type line holds no
// fact of either.
func parseSpellClass(phrase string, isSubtype func(string) bool) (SpellClass, bool) {
	cl := SpellClass{Phrase: phrase}
	list := strings.ReplaceAll(strings.ReplaceAll(phrase, ", or ", ", "), " or ", ", ")
	for _, alt := range strings.Split(list, ",") {
		words := strings.Fields(alt)
		if len(words) == 0 {
			return SpellClass{}, false
		}
		var tests []spellTest
		for _, w := range words {
			test, ok := spellWord(w, isSubtype)
			if !ok {
				return SpellClass{}, false
			}
			tests = append(tests, test)
		}
		cl.alts = append(cl.alts, tests)
	}
	return cl, len(cl.alts) > 0
}

// spellColors are the color words of a cast trigger.
var spellColors = map[string]mtgv1.Color{
	"white": mtgv1.Color_COLOR_W, "blue": mtgv1.Color_COLOR_U, "black": mtgv1.Color_COLOR_B,
	"red": mtgv1.Color_COLOR_R, "green": mtgv1.Color_COLOR_G,
}

// spellWord reads one word of a cast trigger. It knows the card types,
// the supertypes, the subtypes that a card holds, the colors, the words
// of a color count, and historic, which the reminder text of Jhoira,
// Weatherlight Captain defines: "Artifacts, legendaries, and Sagas are
// historic." A word that starts with "non" negates the rest.
func spellWord(w string, isSubtype func(string) bool) (spellTest, bool) {
	w = strings.Trim(w, "'-")
	if rest, ok := strings.CutPrefix(w, "non"); ok && rest != "" {
		test, ok := spellWord(rest, isSubtype)
		if !ok {
			return nil, false
		}
		return func(c *mtgv1.Card) bool { return !test(c) }, true
	}
	if col, ok := spellColors[w]; ok {
		return func(c *mtgv1.Card) bool { return slices.Contains(c.GetColors(), col) }, true
	}
	switch w {
	case "historic":
		return func(c *mtgv1.Card) bool {
			return slices.Contains(c.GetCardTypes(), "Artifact") || slices.Contains(c.GetSupertypes(), "Legendary") ||
				slices.Contains(c.GetSubtypes(), "Saga")
		}, true
	case "multicolored":
		return func(c *mtgv1.Card) bool { return len(c.GetColors()) > 1 }, true
	case "monocolored":
		return func(c *mtgv1.Card) bool { return len(c.GetColors()) == 1 }, true
	case "colorless":
		return func(c *mtgv1.Card) bool { return len(c.GetColors()) == 0 }, true
	case "legendary", "snow":
		super := title(w)
		return func(c *mtgv1.Card) bool { return slices.Contains(c.GetSupertypes(), super) }, true
	}
	if ty := title(w); cardTypes[ty] {
		return func(c *mtgv1.Card) bool { return slices.Contains(c.GetCardTypes(), ty) }, true
	}
	if st := title(w); isSubtype(st) {
		return func(c *mtgv1.Card) bool { return slices.Contains(c.GetSubtypes(), st) }, true
	}
	return nil, false
}
