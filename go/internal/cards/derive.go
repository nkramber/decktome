package cards

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// derive fills the computed fields: parsed types, copy-limit exceptions,
// and commander eligibility. The rules engine (PR-5) reads these fields,
// never the raw strings.
func derive(c *mtgv1.Card) {
	c.Supertypes, c.CardTypes, c.Subtypes = parseTypeLine(c.TypeLine)
	// The sentence is anchored (CR 113.6n). A card that only talks about
	// "any number of cards named" in another context must not match.
	c.AnyCountInDeck = strings.Contains(c.OracleText, "A deck can have any number of cards named")
	c.MaxCopiesOverride = maxCopiesOverride(c.OracleText)
	c.IsBackground = slices.Contains(c.Subtypes, "Background")
	c.IsCompanion = slices.Contains(c.Keywords, "Companion")
	c.Partner = partnerKind(c)
	c.PartnerText = ""
	switch c.Partner {
	case mtgv1.PartnerKind_PARTNER_KIND_WITH:
		c.PartnerWithName = partnerWithName(c)
	case mtgv1.PartnerKind_PARTNER_KIND_PARTNER:
		c.PartnerText = partnerText(c.OracleText)
	}
	c.CanBeCommander = canBeCommander(c)
}

// upToCopiesRe matches "A deck can have up to seven cards named ...".
// Scryfall spells the number as a word (Seven Dwarves, Nazgûl, checked
// 2026-08-24). Digits are accepted as well.
var upToCopiesRe = regexp.MustCompile(`A deck can have up to ([A-Za-z0-9]+) cards named`)

var numberWords = map[string]int32{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6,
	"seven": 7, "eight": 8, "nine": 9, "ten": 10, "eleven": 11, "twelve": 12,
}

// maxCopiesOverride reads the per-card copy limit from the rules text.
// Zero means the format limit applies.
func maxCopiesOverride(text string) int32 {
	m := upToCopiesRe.FindStringSubmatch(text)
	if m == nil {
		return 0
	}
	word := strings.ToLower(m[1])
	if n, ok := numberWords[word]; ok {
		return n
	}
	n, err := strconv.Atoi(word)
	if err != nil || n <= 0 {
		return 0
	}
	return int32(n)
}

// partnerText reads the variant name after "Partner—" (CR 702.124i).
// Example: "Partner—Survivors (You can have...)" gives "Survivors".
// Plain "Partner" gives "".
func partnerText(text string) string {
	i := strings.Index(text, "Partner—")
	if i < 0 {
		return ""
	}
	rest := text[i+len("Partner—"):]
	if j := strings.IndexAny(rest, "(\n"); j >= 0 {
		rest = rest[:j]
	}
	return strings.TrimSpace(rest)
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
// Scryfall prints Friends forever as "Partner—Friends forever" with the
// keyword "Partner", so that text check runs before the plain keyword.
// Other "Partner—[text]" variants (Survivors, Father & son, Character
// select) stay PARTNER_KIND_PARTNER and carry the text in PartnerText.
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
	case slices.Contains(c.Keywords, "Partner"), strings.Contains(c.OracleText, "Partner—"):
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

// canBeCommander applies CR 903.3 (text of 2026-08-07): a legendary
// creature, a legendary Vehicle or Spacecraft with a power/toughness box,
// or a card whose text allows it. Only the front face counts (CR 712.8a),
// so a card such as Bloodline Keeper does not qualify.
//
// Known gap: Grist, the Hunger Tide is a creature outside the battlefield
// by a characteristic-defining ability and is a legal commander. Scryfall
// card JSON carries no "commander-eligible" signal, and a name match is
// not acceptable, so Grist reads as not eligible here (2026-08-24).
func canBeCommander(c *mtgv1.Card) bool {
	if strings.Contains(c.OracleText, "can be your commander") &&
		!strings.Contains(c.OracleText, "can't be your commander") {
		return true
	}
	line, power := c.TypeLine, c.Power
	if len(c.Faces) > 0 {
		line, power = c.Faces[0].TypeLine, c.Faces[0].Power
	}
	supers, types, subs := parseTypeLine(line)
	if !slices.Contains(supers, "Legendary") {
		return false
	}
	if slices.Contains(types, "Creature") {
		return true
	}
	if slices.Contains(subs, "Vehicle") || slices.Contains(subs, "Spacecraft") {
		return power != ""
	}
	return false
}
