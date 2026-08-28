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
			"A sideboard of up to 15 cards, under the same copy limit across both."
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

// these reads "this card" or "these cards", for a sentence about a list.
func these(n int) string {
	if n == 1 {
		return "this card"
	}
	return "these cards"
}

// CodeThinCommanderPool reports a library with no commander for the
// theme, in an owned mode. The retired weak-pool row asked about this
// before the build, and a delegated commander silenced it (D-232).
const CodeThinCommanderPool = "thin_commander_pool"

// CodeOverBudget reports a deck that costs more than the user allowed.
// It is a warning and never a block: the price is a daily estimate and
// not a rule of the game, and a deck the user can trim is more use than
// no deck (D-17, D-236).
const CodeOverBudget = "over_budget"

// BuyCost is what the user must buy: every copy the collection does not
// cover, at the card's display price. A deck built from an owned pool
// costs nothing to buy.
func BuyCost(deck *mtgv1.Deck) float64 {
	total := 0.0
	for _, c := range allCards(deck) {
		short := c.GetCount() - c.GetOwnedCount()
		if short > 0 && c.GetPriceUsd() > 0 {
			total += float64(short) * c.GetPriceUsd()
		}
	}
	return total
}

// DeckCost is what the whole deck is worth, owned copies included. The
// budget-scope question of D-77 asks the user which of the two they mean.
func DeckCost(deck *mtgv1.Deck) float64 {
	total := 0.0
	for _, c := range allCards(deck) {
		total += float64(c.GetCount()) * c.GetPriceUsd()
	}
	return total
}

func allCards(deck *mtgv1.Deck) []*mtgv1.DeckCard {
	return append(append([]*mtgv1.DeckCard(nil), deck.GetCards()...), deck.GetSideboard()...)
}

// CodeLockedCardMissing reports a card the user said to keep that the
// deck does not hold. It blocks, so the repair turn gets one chance to
// put the card back (D-70, D-242).
const CodeLockedCardMissing = "locked_card_missing"

// PreconSharePercent is how much of a named precon a built deck keeps.
// The owner set it on 2026-08-26, and called it a start and not a settled
// figure (D-218, answers OQ-21). On 2026-08-28 the owner set the base:
// the nonbasic names of the precon, with basic lands free to swap. The
// rest is a ceiling and not a target (A-5 of the 2026-08-28 audit).
const PreconSharePercent = 85

// CodePreconShare is the finding an upgrade gets when it drops too much
// of the precon it was asked to upgrade.
const CodePreconShare = "precon_share"

// PreconKeepCount is how many of a precon's nonbasic names a built deck
// must keep (D-218, A-5). The share is a percentage, and the prompt
// states a count: a model asked for a percentage must do arithmetic
// against a list it is still writing, and deck gate prompts 17 and 18
// kept 68 and 29 percent of theirs (D-248).
func PreconKeepCount(total int) int {
	if total <= 0 {
		return 0
	}
	// Round up, so the share is met and not approached.
	return (total*PreconSharePercent + 99) / 100
}

// CodePreconCardsRestored reports the precon cards the builder put back
// to meet the share. It is an INFO: the deck is what the user asked for,
// and they should know the builder finished the job (D-250).
const CodePreconCardsRestored = "precon_cards_restored"
