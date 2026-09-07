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

// searchesLand reports text that searches the library for a land: "a
// land card", "a basic land card", or a basic land type such as "a
// Plains or Swamp card".
func searchesLand(text string) bool {
	i := strings.Index(text, "search your library for")
	if i < 0 {
		return false
	}
	clause := text[i:]
	if j := strings.Index(clause, "card"); j >= 0 {
		clause = clause[:j]
	}
	for _, word := range []string{"land", "plains", "island", "swamp", "mountain", "forest"} {
		if strings.Contains(clause, word) {
			return true
		}
	}
	return false
}

// isFetch reports a land that searches for a land card, which counts as
// a source of every color the deck's basics make (Karsten 2022).
func isFetch(c *mtgv1.Card) bool {
	return isLand(c) && searchesLand(strings.ToLower(c.GetOracleText()))
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
