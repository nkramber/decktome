package generate

import (
	"fmt"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// The revision findings (PR-12B). A removed card that stayed and a kept
// card that left are blocks, like a locked card that is missing (D-242):
// the deck answers the user's own instruction with silence otherwise. A
// card over the mana cap warns and buys the repair turn, like a budget.
const (
	CodeRevisionRemovedPresent = "revision_removed_present"
	CodeRevisionKeptMissing    = "revision_kept_missing"
	CodeRevisionOverManaValue  = "revision_over_mana_value"
)

// CheckRevision reads the brief against the deck the model returned. The
// pool already dropped the removed cards and the cards over the cap, so
// a finding here means the model named a card outside the pool, or a
// kept card is absent. The sideboard counts as the deck, and a commander
// counts as held: it is in the deck, in the command zone.
func CheckRevision(deck *mtgv1.Deck, r *Revision, cards rules.CardSource) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	present := map[string]bool{}
	for _, c := range allCards(deck) {
		present[candidates.FoldName(c.GetName())] = true
	}
	var stayed []string
	for _, name := range r.Remove {
		if present[candidates.FoldName(name)] {
			stayed = append(stayed, name)
		}
	}
	if len(stayed) > 0 {
		out = append(out, &mtgv1.Finding{Code: CodeRevisionRemovedPresent, Severity: mtgv1.Severity_SEVERITY_BLOCK,
			Message: fmt.Sprintf("the deck still holds %s, which you asked to remove", strings.Join(stayed, ", "))})
	}
	if cards != nil {
		for _, id := range deck.GetCommanderOracleIds() {
			if c, ok := cards.ByOracleID(id); ok {
				present[candidates.FoldName(c.GetName())] = true
			}
		}
	}
	var gone []string
	for _, name := range r.Keep {
		if !present[candidates.FoldName(name)] {
			gone = append(gone, name)
		}
	}
	if len(gone) > 0 {
		out = append(out, &mtgv1.Finding{Code: CodeRevisionKeptMissing, Severity: mtgv1.Severity_SEVERITY_BLOCK,
			Message: fmt.Sprintf("the deck does not hold %s, which you asked to keep", strings.Join(gone, ", "))})
	}
	if r.MaxManaValue > 0 && cards != nil {
		var over []string
		for _, c := range deck.GetCards() {
			card, ok := cards.ByOracleID(c.GetOracleId())
			if !ok || isLand(card) {
				continue
			}
			if card.GetManaValue() > r.MaxManaValue {
				over = append(over, fmt.Sprintf("%s (%g)", card.GetName(), card.GetManaValue()))
			}
		}
		if len(over) > 0 {
			out = append(out, &mtgv1.Finding{Code: CodeRevisionOverManaValue, Severity: mtgv1.Severity_SEVERITY_WARN,
				Message: fmt.Sprintf("the deck holds %s above mana value %g, and you asked for none", strings.Join(over, ", "), r.MaxManaValue)})
		}
	}
	return out
}

// AllowedByRevision says whether a card may sit in the pool of a revised
// build: not one the user wants out, and not a nonland over the cap.
// The pool is the whole contract with the model (D-222), so a card the
// brief forbids never reaches it. A kept card and an exempt card pass
// the cap, or the deck must hold a card the model can not name.
func AllowedByRevision(r *Revision, card *mtgv1.Card) bool {
	if r == nil {
		return true
	}
	name := candidates.FoldName(card.GetName())
	for _, n := range r.Remove {
		if candidates.FoldName(n) == name {
			return false
		}
	}
	if r.MaxManaValue <= 0 || isLand(card) || card.GetManaValue() <= r.MaxManaValue {
		return true
	}
	for _, n := range r.Keep {
		if candidates.FoldName(n) == name {
			return true
		}
	}
	for _, id := range r.Exempt {
		if id == card.GetOracleId() {
			return true
		}
	}
	return false
}

func isLand(c *mtgv1.Card) bool {
	for _, t := range c.GetCardTypes() {
		if t == "Land" {
			return true
		}
	}
	return false
}
