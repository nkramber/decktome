// Package generate turns a filled session into a deck. The model writes
// the list from the shortlist, and the code checks every name and every
// rule before the deck reaches the user (roadmap PR-8).
package generate

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
)

// A model names a card that exists and is not the card meant. "Ajani's
// Pridemate" and "Ajani's Welcome" are one letter apart in intent and
// far apart in play. A fuzzy match hides that, so the normalizer matches
// exact names only, and it reports every miss (F-13).
//
// The match is case-insensitive and it ignores the spaces around the
// name. It is exact in every other way. A near name is a miss, and a
// miss is never a silent substitution.

// Miss is one card name the model wrote that the shortlist does not
// hold.
type Miss struct {
	// Name is the name the model wrote.
	Name string
	// Near lists shortlist names that a person may confuse with Name.
	// The repair turn shows these, and the code never picks one (F-13).
	Near []string
}

// Normalized is the result of one normalize pass.
type Normalized struct {
	// Cards are the deck cards whose names the shortlist holds. Each one
	// carries the oracle id and the owned counts of the shortlist entry.
	Cards []*mtgv1.DeckCard
	// Misses are the names the shortlist does not hold, in the order the
	// model wrote them.
	Misses []Miss
}

// Pool is the shortlist the model may write from. Build makes one from a
// candidates.List, and the normalizer reads nothing else: a card outside
// the pool is a miss even when the card index knows it.
type Pool struct {
	byName   map[string]*mtgv1.Card
	byOracle map[string]*mtgv1.Card
	owned    map[string]int32
	names    []string
}

// NewPool indexes the cards the model may name. A later card with the
// same name replaces an earlier one, so the caller passes the commander
// list first and the shortlist after it.
func NewPool(cards []*mtgv1.Card, owned map[string]int32) *Pool {
	p := &Pool{
		byName:   make(map[string]*mtgv1.Card, len(cards)),
		byOracle: make(map[string]*mtgv1.Card, len(cards)),
		owned:    owned,
	}
	for _, c := range cards {
		if c.GetName() == "" {
			continue
		}
		key := candidates.FoldName(c.GetName())
		if _, seen := p.byName[key]; !seen {
			p.names = append(p.names, c.GetName())
		}
		p.byName[key] = c
		if id := c.GetOracleId(); id != "" {
			p.byOracle[id] = c
		}
	}
	sort.Strings(p.names)
	return p
}

// Size is how many distinct names the pool holds.
func (p *Pool) Size() int { return len(p.byName) }

// Filter returns the pool without the cards keep refuses. A revision
// drops the cards the user wants out and the cards over the cap, so the
// model can not name them (PR-12B).
func (p *Pool) Filter(keep func(*mtgv1.Card) bool) *Pool {
	var cards []*mtgv1.Card
	for _, name := range p.names {
		c := p.byName[candidates.FoldName(name)]
		if keep(c) {
			cards = append(cards, c)
		}
	}
	return NewPool(cards, p.owned)
}

// Names lists the pool names in sort order. The prompt writes this list,
// and the repair turn reads it again.
func (p *Pool) Names() []string { return append([]string(nil), p.names...) }

// Card returns the pool card of an exact name.
func (p *Pool) Card(name string) (*mtgv1.Card, bool) {
	c, ok := p.byName[candidates.FoldName(name)]
	return c, ok
}

// ByOracleID returns the pool card of an oracle id. The builder reads
// the pool by id when it puts a card back or names a locked card, and a
// scan of every name for each id was the slow way to do that.
func (p *Pool) ByOracleID(id string) (*mtgv1.Card, bool) {
	c, ok := p.byOracle[id]
	return c, ok
}

// OwnedCount is how many copies of a card the collection holds, zero
// with no collection. Every card the builder inserts reads it, so a
// card the user owns is never charged as a purchase.
func (p *Pool) OwnedCount(id string) int32 { return p.owned[id] }

// Entry is one line of the model's deck list.
type Entry struct {
	Name   string `json:"name"`
	Count  int32  `json:"count"`
	Role   string `json:"role"`
	Reason string `json:"reason"`
}

// Normalize matches every entry against the pool. A matched entry
// becomes a DeckCard with the oracle id and the owned counts filled in.
// An unmatched entry becomes a Miss, and it reaches no deck. The owned
// flag reads the count of an oracle id across every entry of the list.
func Normalize(p *Pool, entries []Entry) Normalized {
	var out Normalized
	for _, e := range entries {
		name := strings.TrimSpace(e.Name)
		if name == "" {
			continue
		}
		c, ok := p.Card(name)
		if !ok {
			out.Misses = append(out.Misses, Miss{Name: name, Near: p.near(name)})
			continue
		}
		owned := p.owned[c.GetOracleId()]
		out.Cards = append(out.Cards, &mtgv1.DeckCard{
			OracleId:   c.GetOracleId(),
			Name:       c.GetName(),
			Count:      e.Count,
			Role:       cardRole(e.Role),
			Reason:     strings.TrimSpace(e.Reason),
			OwnedCount: owned,
			// The display price of the card, per copy. It follows the
			// paper printing the index shows (D-231).
			PriceUsd: c.GetPriceUsd(),
		})
	}
	markOwned(out.Cards)
	return out
}

// markOwned sets the owned flag of every entry from the count of its
// oracle id across the whole list, so two entries of one card are owned
// only when the collection covers both (D-37).
func markOwned(cards []*mtgv1.DeckCard) {
	need := map[string]int32{}
	for _, c := range cards {
		need[c.GetOracleId()] += c.GetCount()
	}
	for _, c := range cards {
		c.Owned = c.GetOwnedCount() >= need[c.GetOracleId()]
	}
}

// nearLimit is how many near names one miss reports. The repair turn
// reads them, and a long list costs tokens and helps nobody.
const nearLimit = 3

// nearStem is the shortest first word that may report a near name. A
// two-letter stem matches half the pool and helps nobody.
const nearStem = 4

// near lists pool names whose first word starts like the miss. "Ajani
// Welcome" reports "Ajani's Welcome" and "Ajani's Pridemate", which is
// the pair F-13 names. It is a hint for the repair turn, and the code
// never picks one of them.
func (p *Pool) near(name string) []string {
	first := firstWord(name)
	if len(first) < nearStem {
		return nil
	}
	var out []string
	for _, n := range p.names {
		w := firstWord(n)
		if len(w) < nearStem {
			continue
		}
		if strings.HasPrefix(w, first) || strings.HasPrefix(first, w) {
			out = append(out, n)
			if len(out) == nearLimit {
				break
			}
		}
	}
	return out
}

// firstWord is the first word of a name, with the letters kept and
// everything else removed. The apostrophe of "Ajani's" must not hide the
// match with "Ajani".
func firstWord(s string) string {
	s = candidates.FoldName(s)
	if i := strings.IndexAny(s, " ,"); i >= 0 {
		s = s[:i]
	}
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

var roleNames = map[string]mtgv1.CardRole{
	"land": mtgv1.CardRole_CARD_ROLE_LAND, "ramp": mtgv1.CardRole_CARD_ROLE_RAMP,
	"draw": mtgv1.CardRole_CARD_ROLE_DRAW, "removal": mtgv1.CardRole_CARD_ROLE_REMOVAL,
	"wipe": mtgv1.CardRole_CARD_ROLE_WIPE, "threat": mtgv1.CardRole_CARD_ROLE_THREAT,
	"interaction": mtgv1.CardRole_CARD_ROLE_INTERACTION, "synergy": mtgv1.CardRole_CARD_ROLE_SYNERGY,
	"wincon": mtgv1.CardRole_CARD_ROLE_WINCON, "other": mtgv1.CardRole_CARD_ROLE_OTHER,
}

// cardRole reads the role word the model wrote. An unknown word is the
// other role, which is a report and never an error: the engine checks
// the deck, and no rule turns on this field.
func cardRole(s string) mtgv1.CardRole {
	if r, ok := roleNames[candidates.FoldName(s)]; ok {
		return r
	}
	return mtgv1.CardRole_CARD_ROLE_OTHER
}

// MissNote is the user-visible line for a name that missed twice. The
// roadmap gives the model one repair turn, and a second miss becomes a
// note the user reads (PR-8).
func MissNote(m Miss) string {
	if len(m.Near) == 0 {
		return fmt.Sprintf("I could not place %q, so it is not in the deck.", m.Name)
	}
	return fmt.Sprintf("I could not place %q, so it is not in the deck. The shortlist holds %s.",
		m.Name, strings.Join(m.Near, ", "))
}
