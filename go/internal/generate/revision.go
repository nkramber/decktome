package generate

import (
	"fmt"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
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
// kept card is absent.
func CheckRevision(deck *mtgv1.Deck, r *Revision, cards rules.CardSource) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	present := map[string]*mtgv1.DeckCard{}
	for _, c := range deck.GetCards() {
		present[fold(c.GetName())] = c
	}
	var stayed []string
	for _, name := range r.Remove {
		if _, ok := present[fold(name)]; ok {
			stayed = append(stayed, name)
		}
	}
	if len(stayed) > 0 {
		out = append(out, &mtgv1.Finding{Code: CodeRevisionRemovedPresent, Severity: mtgv1.Severity_SEVERITY_BLOCK,
			Message: fmt.Sprintf("the deck still holds %s, which you asked to remove", strings.Join(stayed, ", "))})
	}
	var gone []string
	for _, name := range r.Keep {
		if _, ok := present[fold(name)]; !ok {
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
// brief forbids never reaches it.
func AllowedByRevision(r *Revision, card *mtgv1.Card) bool {
	if r == nil {
		return true
	}
	for _, name := range r.Remove {
		if fold(name) == fold(card.GetName()) {
			return false
		}
	}
	if r.MaxManaValue > 0 && !isLand(card) && card.GetManaValue() > r.MaxManaValue {
		return false
	}
	return true
}

func isLand(c *mtgv1.Card) bool {
	for _, t := range c.GetCardTypes() {
		if t == "Land" {
			return true
		}
	}
	return false
}

func fold(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
