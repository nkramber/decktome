package profile

import (
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The card readers of the profile. Each one reads the parsed types, the
// produced mana, and the Oracle text of one card. They are guesses at
// what a card does, and the tags decide where a tag exists.

func isLand(c *mtgv1.Card) bool { return slices.Contains(c.GetCardTypes(), "Land") }

func isBasic(c *mtgv1.Card) bool {
	return isLand(c) && slices.Contains(c.GetSupertypes(), "Basic")
}

func isCreature(c *mtgv1.Card) bool { return slices.Contains(c.GetCardTypes(), "Creature") }

// producesMana reports a card that adds mana: Scryfall lists the colors
// it produces, or its text says "add {".
func producesMana(c *mtgv1.Card) bool {
	if len(c.GetProducedMana()) > 0 {
		return true
	}
	return strings.Contains(strings.ToLower(c.GetOracleText()), "add {")
}

// sourceColorsOf answers the colors a nonland card counts as a source of
// (F-103). It counts only through a tap mana ability whose cost
// sacrifices nothing, so a ritual, a Treasure maker, and a sacrifice altar
// count as no source. It counts as no source of a color its own cost
// needs, because the card is not in play before that color is.
func sourceColorsOf(c *mtgv1.Card) []mtgv1.Color {
	types := c.GetCardTypes()
	if slices.Contains(types, "Instant") || slices.Contains(types, "Sorcery") || !tapManaAbility(c.GetOracleText()) {
		return nil
	}
	own := colorPips(c.GetManaCost())
	var out []mtgv1.Color
	for _, col := range c.GetProducedMana() {
		if col == mtgv1.Color_COLOR_C || own[col] > 0 {
			continue
		}
		out = append(out, col)
	}
	return out
}

// tapManaAbility reports text that holds a mana ability with {T} in its
// cost and no sacrifice. A cost reads back from ": Add" to the last quote,
// period, or parenthesis, so the reminder text of a Treasure token,
// "{T}, Sacrifice this token: Add", reads as a sacrifice.
func tapManaAbility(text string) bool {
	for _, line := range strings.Split(text, "\n") {
		rest := line
		for {
			i := strings.Index(rest, ": Add")
			if i < 0 {
				break
			}
			cost := rest[:i]
			if j := strings.LastIndexAny(cost, "\".("); j >= 0 {
				cost = cost[j+1:]
			}
			if strings.Contains(cost, "{T}") && !strings.Contains(strings.ToLower(cost), "sacrifice") {
				return true
			}
			rest = rest[i+len(": Add"):]
		}
	}
	return false
}

// manaMade is how much mana one activation of a rock or a dork adds:
// two for a card that adds {C}{C}, as Sol Ring does, and one otherwise.
func manaMade(c *mtgv1.Card) int {
	if strings.Contains(c.GetOracleText(), "{C}{C}") {
		return 2
	}
	return 1
}

// entersTapped reports a land that enters tapped with no choice. A land
// that says "unless" enters untapped when its condition holds, and a
// shock land that offers two life is an untapped land at the price.
func entersTapped(c *mtgv1.Card) bool {
	text := strings.ToLower(c.GetOracleText())
	if !strings.Contains(text, "enters tapped") && !strings.Contains(text, "enters the battlefield tapped") {
		return false
	}
	return !strings.Contains(text, "unless") && !strings.Contains(text, "pay 2 life")
}

// The readers the mana pass of PR-33 needs. It moves the mana base of a
// built deck until every mana feature sits in band, and it must read a
// land the way the profile that grades it does (F-78, Part 3). One
// definition serves both.

// IsLand reports a land card.
func IsLand(c *mtgv1.Card) bool { return isLand(c) }

// IsBasic reports a basic land.
func IsBasic(c *mtgv1.Card) bool { return isBasic(c) }

// EntersTapped reports a land that enters tapped with no choice.
func EntersTapped(c *mtgv1.Card) bool { return entersTapped(c) }

// ProducesMana reports a nonland card that adds mana. The pass reads it
// to raise the mana of turn four, which ramp of a low mana value moves
// more than any land does.
func ProducesMana(c *mtgv1.Card) bool { return producesMana(c) }

// IsColorlessLand reports a nonbasic land whose only mana is colorless.
// Every bracket caps them, and the pass trades one for a land that makes
// a color the deck needs.
func IsColorlessLand(c *mtgv1.Card) bool { return isColorlessLand(c) }

// isColorlessLand reports a nonbasic land whose only mana is colorless.
// A land that produces nothing, such as a fetch land, is not one.
func isColorlessLand(c *mtgv1.Card) bool {
	if !isLand(c) || isBasic(c) {
		return false
	}
	pm := c.GetProducedMana()
	return len(pm) == 1 && pm[0] == mtgv1.Color_COLOR_C
}

// isFastMana reports a nonland, noncreature mana producer of mana value
// one or less: Sol Ring, a ritual, a Mox. A signet costs two and a dork
// is a creature, so neither counts.
func isFastMana(c *mtgv1.Card) bool {
	return !isLand(c) && !isCreature(c) && c.GetManaValue() <= 1 && producesMana(c)
}

// searchClause is the lower-case text of a library search up to the card
// it finds, such as "search your library for a plains or swamp " on
// Marsh Flats. ok is false when the text holds no search.
func searchClause(text string) (clause string, ok bool) {
	i := strings.Index(text, "search your library for")
	if i < 0 {
		return "", false
	}
	clause = text[i:]
	if j := strings.Index(clause, "card"); j >= 0 {
		clause = clause[:j]
	}
	return clause, true
}

// basicType is one basic land type, as a search names it, and the color
// of the basic land that carries it.
type basicType struct {
	word  string
	color mtgv1.Color
}

// basicTypes lists the basic land types in color order.
var basicTypes = []basicType{
	{"plains", mtgv1.Color_COLOR_W}, {"island", mtgv1.Color_COLOR_U}, {"swamp", mtgv1.Color_COLOR_B},
	{"mountain", mtgv1.Color_COLOR_R}, {"forest", mtgv1.Color_COLOR_G},
}

// searchesLand reports text that searches the library for a land: "a
// land card", "a basic land card", or a basic land type such as "a
// Plains or Swamp card".
func searchesLand(text string) bool {
	clause, ok := searchClause(text)
	if !ok {
		return false
	}
	if strings.Contains(clause, "land") {
		return true
	}
	return slices.ContainsFunc(basicTypes, func(bt basicType) bool { return strings.Contains(clause, bt.word) })
}

// isFetch reports a land that searches for a land card. It counts as a
// source of each deck color it can find (Karsten 2022, F-147).
func isFetch(c *mtgv1.Card) bool {
	return isLand(c) && searchesLand(strings.ToLower(c.GetOracleText()))
}

// fetchSources lists the deck colors a fetch land counts as a source of,
// in color order (F-147, D-735). A search for a basic land card or a land
// card finds every deck color. A search that names basic land types finds
// the color of each type, and the colors of each deck land of those types:
// Marsh Flats finds Watery Grave, an Island Swamp. So Marsh Flats is no
// blue source in a blue-black deck of basic lands.
func fetchSources(c *mtgv1.Card, colors []mtgv1.Color, typed map[string]map[mtgv1.Color]bool) []mtgv1.Color {
	clause, _ := searchClause(strings.ToLower(c.GetOracleText()))
	found := map[mtgv1.Color]bool{}
	named := false
	for _, bt := range basicTypes {
		if !strings.Contains(clause, bt.word) {
			continue
		}
		named = true
		found[bt.color] = true
		for col := range typed[bt.word] {
			found[col] = true
		}
	}
	var out []mtgv1.Color
	for _, col := range colors {
		if !named || found[col] {
			out = append(out, col)
		}
	}
	return out
}

// typedLandColors maps each basic land type to the colors of the nonbasic
// lands of a deck that carry it, as Watery Grave carries Island and Swamp.
func typedLandColors(entries []entry) map[string]map[mtgv1.Color]bool {
	out := map[string]map[mtgv1.Color]bool{}
	for _, e := range entries {
		if !isLand(e.card) || isBasic(e.card) {
			continue
		}
		for _, sub := range e.card.GetSubtypes() {
			word := strings.ToLower(sub)
			if !slices.ContainsFunc(basicTypes, func(bt basicType) bool { return bt.word == word }) {
				continue
			}
			if out[word] == nil {
				out[word] = map[mtgv1.Color]bool{}
			}
			for _, col := range e.card.GetProducedMana() {
				if col != mtgv1.Color_COLOR_C {
					out[word][col] = true
				}
			}
		}
	}
	return out
}

// LandClass is the class of a land for the colors of one deck, and the
// lower class is the better land (PR-52, F-139, D-733). The order follows
// real lists: at the median a TopDeck list plays more untapped duals than
// any other class, and no tapped dual (M-17).
type LandClass int

const (
	// LandUntappedDual makes two or more deck colors and enters untapped.
	// A shock land counts, because it enters untapped for 2 life.
	LandUntappedDual LandClass = iota
	// LandFetch searches for a land of two or more deck colors and puts it
	// onto the battlefield untapped, as Polluted Delta does.
	LandFetch
	// LandUntappedOnCondition makes two or more deck colors and enters
	// untapped when a condition holds, as Sunken Hollow does. Fabled
	// Passage counts here, because it untaps its land at four lands.
	LandUntappedOnCondition
	// LandManaOnCondition makes two or more deck colors, and its colored
	// mana has a condition or a small cost, as Exotic Orchard, Opal
	// Palace, and Fetid Heath do.
	LandManaOnCondition
	// LandTappedDual makes two or more deck colors and enters tapped, as
	// Thriving Isle does. Evolving Wilds counts here, because its land
	// enters tapped. So does a land whose condition most decks miss, as
	// Secluded Glen does.
	LandTappedDual
	// LandOther is every other land: a land of one deck color or none, as
	// a basic land and Ash Barrens are, and a fetch land that finds one
	// deck color, as Marsh Flats does in a blue-black deck.
	LandOther
)

// manaConditions are the words of a mana ability with a condition, on the
// line of the ability, as on Plaza of Heroes, Spire of Industry, Exotic
// Orchard, Hidden Lair, and Paliano, the High City.
var manaConditions = []string{"spend this mana only", "activate only if", "could produce", "among", "chosen"}

// tapState is how a land enters the battlefield.
type tapState int

const (
	tapNever tapState = iota
	tapOnCondition
	tapAlways
)

// entersTappedWhen reads the sentence that makes a land enter tapped. A
// shock land enters untapped for 2 life. A condition that a deck meets,
// such as a basic land type or a count of lands, makes the land enter
// untapped on a condition. A condition that most decks miss makes it a
// tapped land.
func entersTappedWhen(text string) tapState {
	sentences := strings.FieldsFunc(text, func(r rune) bool { return r == '\n' || r == '.' })
	for i, s := range sentences {
		if !strings.Contains(s, "enters tapped") && !strings.Contains(s, "enters the battlefield tapped") {
			continue
		}
		var condition string
		unless := strings.Index(s, "unless")
		switch {
		case unless >= 0:
			condition = s[unless:]
		case strings.HasPrefix(strings.TrimSpace(s), "if you don't") && i > 0:
			condition = sentences[i-1]
		default:
			return tapAlways
		}
		switch {
		case strings.Contains(condition, "pay 2 life"):
			return tapNever
		case weakCondition(strings.ReplaceAll(condition, "this land", "")):
			return tapAlways
		}
		return tapOnCondition
	}
	return tapNever
}

// weakCondition reports a condition that most decks miss: a creature type
// to reveal, behold, or control, as on Secluded Glen, or a life total of
// 13 or less, as on Murky Sewer. A basic land type or a count of lands is
// a condition that a deck meets.
func weakCondition(condition string) bool {
	if strings.Contains(condition, "or less life") {
		return true
	}
	if !strings.Contains(condition, "reveal") && !strings.Contains(condition, "behold") && !strings.Contains(condition, "you control a") {
		return false
	}
	return !strings.Contains(condition, "land") &&
		!slices.ContainsFunc(basicTypes, func(bt basicType) bool { return strings.Contains(condition, bt.word) })
}

// fetchSearch is the search clause of a fetch land: an ability that costs
// no mana and puts a land onto the battlefield. Ash Barrens puts its land
// into the hand, and Demolition Field pays {2} and destroys a land first.
func fetchSearch(text string) (string, bool) {
	i := strings.Index(text, "search your library for")
	if i < 0 || !searchesLand(text) || !strings.Contains(text[i:], "onto the battlefield") {
		return "", false
	}
	line := text[:i]
	if j := strings.LastIndex(line, "\n"); j >= 0 {
		line = line[j+1:]
	}
	if k := strings.Index(line, ":"); k >= 0 && strings.Contains(strings.ReplaceAll(line[:k], "{t}", ""), "{") {
		return "", false
	}
	clause, _ := searchClause(text)
	return clause, true
}

// landDrawbacks are the words of a land that costs more than a tap: a
// land it returns, control it loses, or an untap step it skips. A counter
// that a mana ability removes is a cost, and cheapManaCost reads it. Opal
// Palace names counters that its commander gets, and that is no drawback.
var landDrawbacks = []string{
	"gains control of this land", "doesn't untap", "return this land", "land you control to its owner's hand", "non-lair land",
}

// LandClassOf answers the class of a land for the deck colors. With no
// deck color every color counts, as the land cap of D-450 reads it. A
// nonland reads LandOther.
//
// The class reads the mana abilities and not the produced mana alone.
// Scryfall lists every color for Lotus Vale, Gemstone Mine, and Survivors'
// Encampment, and none of them is a dual.
func LandClassOf(c *mtgv1.Card, colors map[mtgv1.Color]bool) LandClass {
	if !isLand(c) || isBasic(c) {
		return LandOther
	}
	text := strings.ToLower(c.GetOracleText())
	if clause, ok := fetchSearch(text); ok {
		return fetchClass(c, text, clause, colors)
	}
	clean, cheap := manaColors(c, text, colors)
	if cheap < 2 {
		return LandOther
	}
	class := LandUntappedDual
	switch entersTappedWhen(text) {
	case tapAlways:
		return LandTappedDual
	case tapOnCondition:
		class = LandUntappedOnCondition
	case tapNever:
	}
	switch {
	case landDrawback(text):
		return LandOther
	case clean < 2:
		return LandManaOnCondition
	}
	return class
}

// manaColors counts the deck colors a land adds with its own mana
// abilities. clean counts the abilities that cost {T}, or {T} and life,
// with no condition on their line. cheap counts the abilities that cost
// {1} or one hybrid mana at most beside the tap, as on Opal Palace and
// Fetid Heath, whatever the condition. A color that costs more, as on
// Cascading Cataracts, is no fixing. An ability in quotes belongs to
// another permanent, as on Forgotten Monument.
func manaColors(c *mtgv1.Card, text string, colors map[mtgv1.Color]bool) (clean, cheap int) {
	cleanSet, cheapSet := map[mtgv1.Color]bool{}, map[mtgv1.Color]bool{}
	for _, line := range strings.Split(text, "\n") {
		rest := line
		for {
			i := strings.Index(rest, ": add")
			if i < 0 {
				break
			}
			cost, effect := rest[:i], rest[i+len(": add"):]
			rest = effect
			j := strings.LastIndexAny(cost, ".(\"")
			if j >= 0 && cost[j] == '"' {
				continue
			}
			if cost = cost[j+1:]; !cheapManaCost(cost) {
				continue
			}
			sentence := effect
			if k := strings.Index(effect, "."); k >= 0 {
				sentence = effect[:k]
			}
			cols := addedColors(c, sentence, colors)
			for col := range cols {
				cheapSet[col] = true
			}
			if cleanManaCost(cost) && !slices.ContainsFunc(manaConditions, func(w string) bool { return strings.Contains(effect, w) }) {
				for col := range cols {
					cleanSet[col] = true
				}
			}
		}
	}
	return len(cleanSet), len(cheapSet)
}

// cheapManaCost reports a mana cost of {T} with life, {1}, or one hybrid
// mana at most, as on Opal Palace and Fetid Heath.
func cheapManaCost(cost string) bool {
	tap := false
	for _, part := range strings.Split(cost, ",") {
		switch p := strings.TrimSpace(part); {
		case p == "{t}":
			tap = true
		case p == "{1}", strings.HasPrefix(p, "pay ") && strings.HasSuffix(p, " life"):
		case len(p) == len("{w/b}") && p[0] == '{' && p[2] == '/' && p[4] == '}':
		default:
			return false
		}
	}
	return tap
}

// addedColors is the set of deck colors one add sentence names: a color
// symbol, or a word such as "any color", which reads the produced mana.
func addedColors(c *mtgv1.Card, sentence string, colors map[mtgv1.Color]bool) map[mtgv1.Color]bool {
	out := map[mtgv1.Color]bool{}
	for _, col := range colorOrder {
		if (colors == nil || colors[col]) && strings.Contains(sentence, "{"+strings.ToLower(colorLetters[col])+"}") {
			out[col] = true
		}
	}
	if strings.Contains(strings.ReplaceAll(sentence, "colorless", ""), "color") {
		for _, col := range c.GetProducedMana() {
			if col != mtgv1.Color_COLOR_C && (colors == nil || colors[col]) {
				out[col] = true
			}
		}
	}
	return out
}

// cleanManaCost reports a mana cost of {T} alone, or {T} and life.
func cleanManaCost(cost string) bool {
	tap := false
	for _, part := range strings.Split(cost, ",") {
		switch p := strings.TrimSpace(part); {
		case p == "{t}":
			tap = true
		case strings.HasPrefix(p, "pay ") && strings.HasSuffix(p, " life"):
		default:
			return false
		}
	}
	return tap
}

// landDrawback reports a land whose text costs more than a tap, such as
// Glimmervoid, which sacrifices itself with no artifact. A sacrifice in a
// cost, as on Horizon Canopy, is no drawback: a colon ends the cost.
func landDrawback(text string) bool {
	if slices.ContainsFunc(landDrawbacks, func(w string) bool { return strings.Contains(text, w) }) {
		return true
	}
	rest := text
	for {
		i := strings.Index(rest, "sacrifice")
		if i < 0 {
			return false
		}
		rest = rest[i+len("sacrifice"):]
		colon, stop := strings.Index(rest, ":"), strings.IndexAny(rest, ".\n")
		if colon < 0 || (stop >= 0 && stop < colon) {
			return true
		}
	}
}

// fetchClass is the class of a land that searches for a land. A search
// that names no basic land type finds every deck color.
func fetchClass(c *mtgv1.Card, text, clause string, colors map[mtgv1.Color]bool) LandClass {
	found, named := 0, false
	for _, bt := range basicTypes {
		if strings.Contains(clause, bt.word) {
			named = true
			if colors == nil || colors[bt.color] {
				found++
			}
		}
	}
	if !named {
		found = len(basicTypes)
		if colors != nil {
			found = len(colors)
		}
	}
	intoTapped := strings.Contains(text, "onto the battlefield tapped")
	switch {
	case found < 2:
		return LandOther
	case entersTapped(c):
		return LandTappedDual
	// Fabled Passage untaps its land at four lands, which a deck meets.
	// Elven Passage untaps it for an Elf, which most decks miss.
	case intoTapped && strings.Contains(text, "untap that land") && strings.Contains(text, "or more lands"):
		return LandUntappedOnCondition
	case intoTapped:
		return LandTappedDual
	}
	return LandFetch
}

// ReadsBasicLands reports a nonbasic land whose text reads basic lands: a
// fetch land for a basic land card, or a land that enters untapped with
// two basic lands. The swap of the mana pass keeps a basic land for each
// one (PR-52).
func ReadsBasicLands(c *mtgv1.Card) bool {
	return isLand(c) && !isBasic(c) && strings.Contains(strings.ToLower(c.GetOracleText()), "basic land")
}

// isLandRamp reports a spell that puts a land onto the battlefield from
// the library, such as Rampant Growth or Cultivate.
func isLandRamp(c *mtgv1.Card) bool {
	if isLand(c) {
		return false
	}
	text := strings.ToLower(c.GetOracleText())
	return searchesLand(text) && strings.Contains(text, "onto the battlefield")
}

// colorPips counts the colored mana symbols of a cost per color. A
// hybrid or Phyrexian symbol can be paid another way, so it is not a
// pip of either color.
func colorPips(manaCost string) map[mtgv1.Color]int {
	out := map[mtgv1.Color]int{}
	for _, sym := range strings.Split(manaCost, "{") {
		sym = strings.TrimSuffix(strings.TrimSpace(sym), "}")
		if c, ok := pipColors[sym]; ok {
			out[c]++
		}
	}
	return out
}

var pipColors = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G,
}

var colorLetters = map[mtgv1.Color]string{
	mtgv1.Color_COLOR_W: "W", mtgv1.Color_COLOR_U: "U", mtgv1.Color_COLOR_B: "B",
	mtgv1.Color_COLOR_R: "R", mtgv1.Color_COLOR_G: "G",
}

// colorOrder is WUBRG, the order every report lists colors in.
var colorOrder = []mtgv1.Color{
	mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G,
}
