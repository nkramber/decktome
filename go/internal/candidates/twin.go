package candidates

import (
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// An owned twin is a card the user owns that does what another card does
// at the same cost (F-232, D-1226). Grizzly Bears, Balduvian Bears, and
// Runeclaw Bear are one card under three names, and so are Llanowar
// Elves and Fyndhorn Elves. A deck that buys one while the collection
// holds another spends money on no gap, and the gap question of D-1011
// lets the deck buy a card to fill a gap alone.
//
// The rule reads the printed card and no tag: the same mana cost, the
// same type line, and the same rules text with the own name folded, and
// power and toughness as high or higher. A card with two faces has no
// twin, because each face must agree, and one face can hide a drawback.

// twinName stands in for the own name of a card in its rules text.
const twinName = "CARDNAME"

// TwinKey is what two twins share, and false for a card that has no twin:
// a basic land, a card with two faces, or a card with no type line.
func TwinKey(c *mtgv1.Card) (string, bool) {
	if c == nil || IsBasicLand(c) || len(c.GetFaces()) > 1 || strings.TrimSpace(c.GetTypeLine()) == "" {
		return "", false
	}
	return strings.Join([]string{c.GetManaCost(), c.GetTypeLine(), twinText(c), c.GetLoyalty()}, "\x00"), true
}

// twinText is the rules text with the own name folded. A legendary card
// calls itself by the part of its name before the comma, as "Ayula"
// does for Ayula, Queen Among Bears.
func twinText(c *mtgv1.Card) string {
	text := c.GetOracleText()
	name := c.GetName()
	text = strings.ReplaceAll(text, name, twinName)
	if short, _, ok := strings.Cut(name, ","); ok && short != "" {
		text = strings.ReplaceAll(text, short, twinName)
	}
	return strings.TrimSpace(text)
}

// TwinBeats says whether owned is a twin of bought: the two share a key,
// and owned has power and toughness as high or higher. A value that is
// not a number, such as "*", must be the same on both cards.
func TwinBeats(owned, bought *mtgv1.Card) bool {
	if owned.GetOracleId() == bought.GetOracleId() {
		return false
	}
	ko, ok := TwinKey(owned)
	if !ok {
		return false
	}
	kb, ok := TwinKey(bought)
	if !ok || ko != kb {
		return false
	}
	return statAtLeast(owned.GetPower(), bought.GetPower()) && statAtLeast(owned.GetToughness(), bought.GetToughness())
}

func statAtLeast(have, want string) bool {
	if have == want {
		return true
	}
	h, err1 := strconv.ParseFloat(have, 64)
	w, err2 := strconv.ParseFloat(want, 64)
	return err1 == nil && err2 == nil && h >= w
}

// addOwnedTwins puts on the main list each owned card that is a twin of
// a card to buy, so the model can name it, and the builder can play it in
// place of that card (D-1226). A card to buy is an upgrade, or an unowned
// card that the staple floor put on the main list. The twin skips the
// caps and keeps the role and the score of the card to buy, as a pin
// does. Without it the floor put Llanowar Elves on the list and left the
// owned Fyndhorn Elves off, and the model bought Llanowar Elves in each
// base replay of F-232.
func addOwnedTwins(main, upgrades, scored []Candidate) ([]Candidate, int) {
	var toBuy []Candidate
	for _, c := range main {
		if c.Owned <= 0 {
			toBuy = append(toBuy, c)
		}
	}
	toBuy = append(toBuy, upgrades...)
	if len(toBuy) == 0 {
		return main, 0
	}
	held := map[string]bool{}
	for _, c := range main {
		held[c.Card.GetOracleId()] = true
	}
	owned := map[string][]Candidate{}
	for _, c := range scored {
		if c.Owned <= 0 {
			continue
		}
		if k, ok := TwinKey(c.Card); ok {
			owned[k] = append(owned[k], c)
		}
	}
	added := 0
	for _, up := range toBuy {
		k, ok := TwinKey(up.Card)
		if !ok {
			continue
		}
		for _, tw := range owned[k] {
			id := tw.Card.GetOracleId()
			if held[id] || !TwinBeats(tw.Card, up.Card) {
				continue
			}
			held[id] = true
			tw.Role, tw.Score = up.Role, up.Score
			tw.Signals = append(append([]string(nil), tw.Signals...), "owned twin of "+up.Card.GetName())
			main = append(main, tw)
			added++
		}
	}
	return main, added
}
