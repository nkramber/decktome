package cards

import (
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// derive fills the computed fields: parsed types, copy-limit exceptions,
// and commander eligibility. The rules engine (PR-5) reads these fields,
// never the raw strings.
func derive(c *mtgv1.Card) {
	c.Supertypes, c.CardTypes, c.Subtypes = parseTypeLine(c.TypeLine)
	c.AnyCountInDeck = strings.Contains(c.OracleText, "any number of cards named")
	c.IsBackground = slices.Contains(c.Subtypes, "Background")
	c.IsCompanion = slices.Contains(c.Keywords, "Companion")
	c.Partner = partnerKind(c)
	if c.Partner == mtgv1.PartnerKind_PARTNER_KIND_WITH {
		c.PartnerWithName = partnerWithName(c)
	}
	c.CanBeCommander = canBeCommander(c)
}

var knownSupertypes = map[string]bool{
	"Basic": true, "Legendary": true, "Snow": true, "World": true,
	"Elite": true, "Ongoing": true, "Token": true,
}

// parseTypeLine splits "Legendary Creature - Elk // Land" style lines.
// A multi-face type line is the union of the face type lines.
func parseTypeLine(line string) (supers, types, subs []string) {
	seenSuper := map[string]bool{}
	seenType := map[string]bool{}
	seenSub := map[string]bool{}
	for _, part := range strings.Split(line, "//") {
		// The dash between types and subtypes is an em dash on Scryfall.
		var left, right string
		if i := strings.Index(part, "—"); i >= 0 {
			left, right = part[:i], part[i+len("—"):]
		} else {
			left = part
		}
		for _, w := range strings.Fields(left) {
			if knownSupertypes[w] {
				if !seenSuper[w] {
					seenSuper[w] = true
					supers = append(supers, w)
				}
			} else if !seenType[w] {
				seenType[w] = true
				types = append(types, w)
			}
		}
		for _, w := range strings.Fields(right) {
			if !seenSub[w] {
				seenSub[w] = true
				subs = append(subs, w)
			}
		}
	}
	return supers, types, subs
}

// partnerKind reads the keywords first. Some variants (Choose a Background,
// Doctor's companion) are not in the Scryfall keywords array, so the rules
// text is the fallback signal. Verified against the 2026-08-24 fixture.
func partnerKind(c *mtgv1.Card) mtgv1.PartnerKind {
	switch {
	case hasKeywordPrefix(c.Keywords, "Partner with"), strings.Contains(c.OracleText, "Partner with "):
		return mtgv1.PartnerKind_PARTNER_KIND_WITH
	case slices.Contains(c.Keywords, "Friends forever"), strings.Contains(c.OracleText, "Friends forever"):
		return mtgv1.PartnerKind_PARTNER_KIND_FRIENDS_FOREVER
	case hasKeywordFold(c.Keywords, "Doctor's companion"), strings.Contains(c.OracleText, "Doctor's companion"):
		// The companion card carries this, not the Doctor.
		return mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION
	case hasKeywordFold(c.Keywords, "Choose a Background"), strings.Contains(c.OracleText, "Choose a Background"):
		// The keyword is "Choose a background" (lowercase b) in the data.
		return mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND
	case slices.Contains(c.Keywords, "Partner"):
		return mtgv1.PartnerKind_PARTNER_KIND_PARTNER
	default:
		return mtgv1.PartnerKind_PARTNER_KIND_NONE
	}
}

func hasKeywordFold(keywords []string, want string) bool {
	for _, k := range keywords {
		if strings.EqualFold(k, want) {
			return true
		}
	}
	return false
}

func hasKeywordPrefix(keywords []string, prefix string) bool {
	for _, k := range keywords {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}

// partnerWithName reads the partner name from "Partner with <name>" text.
// The reminder text in parentheses is not part of the name.
func partnerWithName(c *mtgv1.Card) string {
	text := c.OracleText
	i := strings.Index(text, "Partner with ")
	if i < 0 {
		return ""
	}
	rest := text[i+len("Partner with "):]
	if j := strings.IndexAny(rest, "(\n"); j >= 0 {
		rest = rest[:j]
	}
	return strings.TrimSpace(rest)
}

// canBeCommander: a legendary creature, or a card whose text allows it.
// The face check covers transform commanders (example: Tovolar).
func canBeCommander(c *mtgv1.Card) bool {
	if strings.Contains(c.OracleText, "can be your commander") {
		return true
	}
	if !slices.Contains(c.Supertypes, "Legendary") {
		return false
	}
	if slices.Contains(c.CardTypes, "Creature") {
		// The front face must be the creature for a multi-face card.
		if len(c.Faces) > 1 {
			return strings.Contains(c.Faces[0].TypeLine, "Creature")
		}
		return true
	}
	return false
}
