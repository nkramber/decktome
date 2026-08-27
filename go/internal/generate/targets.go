package generate

import (
	"fmt"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The role targets and the limits block come from the corpus role guide
// (mtg-corpus section 6). They are guide numbers and not rules, so the
// prompt asks the model to come as close as the shortlist allows, and no
// check enforces them.

// TargetsFor is the wanted count per job. Commander counts the 99, and a
// 60-card format counts the main deck.
func TargetsFor(format mtgv1.FormatId, power *mtgv1.PowerLevel) map[string]int {
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		return map[string]int{
			"land": 36, "ramp": 10, "draw": 10, "removal": 8,
			"wipe": 3, "threat": 12, "interaction": 6, "synergy": 14,
		}
	}
	// A 60-card deck's land count follows the archetype, and the guide
	// range is 20 to 27. The middle of that range is the default, and a
	// tournament target leans on interaction.
	lands := 24
	if power.GetSixtyStep() == mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT {
		return map[string]int{
			"land": lands, "draw": 6, "removal": 8, "threat": 12,
			"interaction": 6, "synergy": 4,
		}
	}
	return map[string]int{
		"land": lands, "ramp": 2, "draw": 6, "removal": 6,
		"threat": 14, "synergy": 8,
	}
}

// LimitsFor is the deck-building limits block the prompt reads. It states
// the size, the copy limit, and the color rule, in the words the model
// must build to (mtg-corpus sections 2.1 and 2.2).
func LimitsFor(format mtgv1.FormatId) string {
	switch format {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		return "Exactly 100 cards, the commander included. List exactly 99 cards, and do not list the commander. " +
			"One copy of each name, basic lands excepted. Every card must fit the commander's color identity. " +
			"No sideboard."
	case mtgv1.FormatId_FORMAT_ID_STANDARD, mtgv1.FormatId_FORMAT_ID_MODERN:
		return "At least 60 cards in the main deck. At most four copies of each name, basic lands excepted. " +
			"A sideboard of exactly 15 cards, under the same copy limit across both."
	default:
		return "At least 60 cards in the main deck. At most four copies of each name, basic lands excepted."
	}
}

// FormatWord names the format for a log line or a status message.
func FormatWord(f mtgv1.FormatId) string {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		return "Commander"
	case mtgv1.FormatId_FORMAT_ID_STANDARD:
		return "Standard"
	case mtgv1.FormatId_FORMAT_ID_MODERN:
		return "Modern"
	default:
		return fmt.Sprintf("format %d", int(f))
	}
}

// DeckSize is the exact card count a format needs, the commander
// included. Zero means the format states a minimum and not an exact size,
// so no padding applies.
func DeckSize(format mtgv1.FormatId) int {
	switch format {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		return 100
	case mtgv1.FormatId_FORMAT_ID_STANDARD, mtgv1.FormatId_FORMAT_ID_MODERN:
		return 60
	}
	return 0
}

// MaxPad is the largest shortfall the builder fills with basic lands. A
// deck one or two cards short is a counting slip, and a basic land is
// always a legal answer. A larger gap is a real failure, and it stays a
// block finding for the user to see (D-225).
const MaxPad = 5

// CodeBasicsAdded reports the basic lands the builder added to reach the
// deck size. It is an INFO, because the deck is legal and the user should
// still know the builder finished the list.
const CodeBasicsAdded = "basics_added"

// plural writes a count and its noun. A finding reaches the user, and
// "1 cards" is a mistake the user sees.
func plural(n int, noun string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, noun)
	}
	return fmt.Sprintf("%d %ss", n, noun)
}
