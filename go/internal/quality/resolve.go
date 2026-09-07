package quality

import (
	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
)

// Roles answers the role a card fills. candidates.Builder.Roles is the
// one implementation.
type Roles func(c *mtgv1.Card) mtgv1.CardRole

// Resolved is a published list read as a deck the profiler and the
// scorer can measure.
type Resolved struct {
	List *meta.List
	Deck *mtgv1.Deck
	// Missing names the rows the card index could not answer.
	Missing []string
}

// Resolve reads a list against the card index. The format is the one
// the fit reads the list as, which a 60-card product leaves open.
func Resolve(l *meta.List, idx *cards.Index, roles Roles, format mtgv1.FormatId) *Resolved {
	r := &Resolved{List: l, Deck: &mtgv1.Deck{Format: &mtgv1.Format{Id: format}}}
	find := func(c meta.Card) (*mtgv1.Card, bool) {
		if c.OracleID != "" {
			if card, ok := idx.ByOracleID(c.OracleID); ok {
				return card, true
			}
		}
		return idx.ByName(c.Name)
	}
	for _, name := range l.Commanders {
		card, ok := idx.ByName(name)
		if !ok {
			r.Missing = append(r.Missing, name)
			continue
		}
		r.Deck.CommanderOracleIds = append(r.Deck.CommanderOracleIds, card.GetOracleId())
	}
	for _, c := range l.Cards {
		card, ok := find(c)
		if !ok {
			r.Missing = append(r.Missing, c.Name)
			continue
		}
		var role mtgv1.CardRole
		if roles != nil {
			role = roles(card)
		}
		r.Deck.Cards = append(r.Deck.Cards, &mtgv1.DeckCard{
			OracleId: card.GetOracleId(), Name: card.GetName(), Count: int32(c.Count), Role: role,
		})
	}
	return r
}

// Usable says whether a resolved list is whole enough to learn from: a
// deck of the format's size, with at most three rows unresolved.
func (r *Resolved) Usable() bool {
	if len(r.Missing) > 3 {
		return false
	}
	n := 0
	for _, dc := range r.Deck.GetCards() {
		n += int(dc.GetCount())
	}
	if r.Deck.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		n += len(r.Deck.GetCommanderOracleIds())
		return n >= 95 && n <= 101 && len(r.Deck.GetCommanderOracleIds()) > 0
	}
	return n >= 55 && n <= 80
}
