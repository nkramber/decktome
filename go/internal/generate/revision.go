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
//
// A land swap the deck did not make is a block too (D-448). The user
// asked for a number of basic lands replaced, and a deck that kept them
// answered the ask with a shrug.
const (
	CodeRevisionRemovedPresent = "revision_removed_present"
	CodeRevisionKeptMissing    = "revision_kept_missing"
	CodeRevisionOverManaValue  = "revision_over_mana_value"
	CodeRevisionLandsKept      = "revision_lands_kept"
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
	if r.SwapBasics > 0 && cards != nil {
		before := nonbasicLands(r.Base, cards)
		after := nonbasicLands(deck.GetCards(), cards)
		if rise := after - before; rise < r.SwapBasics {
			out = append(out, &mtgv1.Finding{Code: CodeRevisionLandsKept, Severity: mtgv1.Severity_SEVERITY_BLOCK,
				Message: fmt.Sprintf("the deck holds %d more nonbasic lands than before, and the change asked for %d basic lands replaced", rise, r.SwapBasics)})
		}
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

// nonbasicLands sums the copies of every nonbasic land in a card list.
// A card the source does not know counts as no land.
func nonbasicLands(list []*mtgv1.DeckCard, cards rules.CardSource) int {
	n := 0
	for _, dc := range list {
		c, ok := cards.ByOracleID(dc.GetOracleId())
		if ok && isLand(c) && !candidates.IsBasicLand(c) {
			n += int(dc.GetCount())
		}
	}
	return n
}

// FitSwapBasics lowers the swap count to what the deck and the pool can
// give: the basic lands the base deck holds, and the nonbasic lands the
// pool offers that the base deck does not hold yet. The check of
// CheckRevision holds the deck to the fitted count, so the model is
// never blocked for a land it can not name (D-448). It returns the
// fitted count.
func FitSwapBasics(r *Revision, pool *Pool) int {
	if r == nil || r.SwapBasics <= 0 || pool == nil {
		return 0
	}
	basics, held := 0, map[string]bool{}
	for _, dc := range r.Base {
		if c, ok := pool.ByOracleID(dc.GetOracleId()); ok && candidates.IsBasicLand(c) {
			basics += int(dc.GetCount())
		}
		held[dc.GetOracleId()] = true
	}
	offered := 0
	for _, name := range pool.Names() {
		c, ok := pool.Card(name)
		if ok && isLand(c) && !candidates.IsBasicLand(c) && !held[c.GetOracleId()] {
			offered++
		}
	}
	r.SwapBasics = min(r.SwapBasics, basics, offered)
	return r.SwapBasics
}

func isLand(c *mtgv1.Card) bool {
	for _, t := range c.GetCardTypes() {
		if t == "Land" {
			return true
		}
	}
	return false
}
