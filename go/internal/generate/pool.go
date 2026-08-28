package generate

import (
	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
)

// FromList makes the pool the model may write from out of a shortlist.
// always holds the cards that join whatever the shortlist ranked: the
// commanders, the locked cards, and the basic lands. They go in first, so
// a card that also appears in the shortlist keeps one entry.
//
// The pool is the whole contract with the model: a card outside it is a
// miss, whatever the card index knows (F-13). Upgrades join it only when
// the session may buy cards, because a deck must not name a card the
// user can neither own nor buy.
func FromList(l *candidates.List, always []*mtgv1.Card, buyList bool) *Pool {
	cards := append([]*mtgv1.Card(nil), always...)
	owned := map[string]int32{}
	add := func(cs []candidates.Candidate) {
		for _, c := range cs {
			if c.Card == nil {
				continue
			}
			cards = append(cards, c.Card)
			if c.Owned > 0 {
				owned[c.Card.GetOracleId()] = c.Owned
			}
			// A pair carries its partner, and the partner must be
			// nameable or the pair can not be built (D-154).
			if c.Partner != nil {
				cards = append(cards, c.Partner)
			}
		}
	}
	add(l.Candidates)
	if buyList {
		add(l.Upgrades)
	}
	return NewPool(cards, owned)
}

// Roles is the job word per card, for the shortlist block. The prompt
// reads the job from the shortlist, so a card the model picks carries the
// job the builder gave it.
func Roles(l *candidates.List) map[string]string {
	out := make(map[string]string, len(l.Candidates))
	for _, c := range l.Candidates {
		if c.Card == nil {
			continue
		}
		out[c.Card.GetOracleId()] = roleWord(c.Role)
	}
	return out
}

var roleWords = map[mtgv1.CardRole]string{
	mtgv1.CardRole_CARD_ROLE_LAND: "land", mtgv1.CardRole_CARD_ROLE_RAMP: "ramp",
	mtgv1.CardRole_CARD_ROLE_DRAW: "draw", mtgv1.CardRole_CARD_ROLE_REMOVAL: "removal",
	mtgv1.CardRole_CARD_ROLE_WIPE: "wipe", mtgv1.CardRole_CARD_ROLE_THREAT: "threat",
	mtgv1.CardRole_CARD_ROLE_INTERACTION: "interaction", mtgv1.CardRole_CARD_ROLE_SYNERGY: "synergy",
	mtgv1.CardRole_CARD_ROLE_WINCON: "wincon",
}

func roleWord(r mtgv1.CardRole) string {
	if w, ok := roleWords[r]; ok {
		return w
	}
	return "other"
}

// basicNames are the five basic lands, one per color. The shortlist
// leaves them out on purpose: ranking a Plains against a real card means
// nothing (candidates.go). A deck still needs them, and a Commander deck
// that runs 36 lands and no basic is a deck nobody can afford.
//
// Gate run 1 of 2026-08-27 returned 12 decks and not one basic land. Two
// of them came up one card short, and one asked for a Plains the pool did
// not hold (D-225).
var basicNames = map[mtgv1.Color]string{
	mtgv1.Color_COLOR_W: "Plains",
	mtgv1.Color_COLOR_U: "Island",
	mtgv1.Color_COLOR_B: "Swamp",
	mtgv1.Color_COLOR_R: "Mountain",
	mtgv1.Color_COLOR_G: "Forest",
}

// BasicLands returns the basic lands of a color identity, in color order.
// find is the card lookup, so this package needs no card index.
//
// A colorless deck gets no basic. Wastes is the colorless basic, and the
// app builds no deck that needs it today.
func BasicLands(find func(string) (*mtgv1.Card, bool), colors []mtgv1.Color) []*mtgv1.Card {
	var out []*mtgv1.Card
	for _, col := range []mtgv1.Color{
		mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B,
		mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G,
	} {
		if !hasColor(colors, col) {
			continue
		}
		if c, ok := find(basicNames[col]); ok {
			out = append(out, c)
		}
	}
	return out
}

func hasColor(list []mtgv1.Color, want mtgv1.Color) bool {
	for _, c := range list {
		if c == want {
			return true
		}
	}
	return false
}

// IsBasic reports whether a card is a basic land.
func IsBasic(c *mtgv1.Card) bool {
	for _, s := range c.GetSupertypes() {
		if s == "Basic" {
			return true
		}
	}
	return false
}

// padWithBasics fills a small shortfall with basic lands and returns how
// many it added. It reads the pool, so a basic the pool does not hold is
// never added, and a colorless deck is never padded.
//
// The cards spread across the colors in turn, so a two-color deck gets
// one of each rather than two of one.
func padWithBasics(deck *mtgv1.Deck, req Request) int {
	size := DeckSize(req.Format)
	if size == 0 {
		return 0
	}
	have := len(deck.GetCommanderOracleIds())
	for _, c := range deck.GetCards() {
		have += int(c.GetCount())
	}
	short := size - have
	if short <= 0 || short > MaxPad {
		return 0
	}
	var basics []*mtgv1.Card
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if ok && IsBasic(c) {
			basics = append(basics, c)
		}
	}
	if len(basics) == 0 {
		return 0
	}
	// An existing entry takes the count, so the deck holds one line per
	// basic and not one line per copy.
	byOracle := map[string]*mtgv1.DeckCard{}
	for _, c := range deck.Cards {
		byOracle[c.GetOracleId()] = c
	}
	for i := 0; i < short; i++ {
		b := basics[i%len(basics)]
		if dc, ok := byOracle[b.GetOracleId()]; ok {
			dc.Count++
			continue
		}
		dc := &mtgv1.DeckCard{
			OracleId: b.GetOracleId(), Name: b.GetName(), Count: 1,
			Role:   mtgv1.CardRole_CARD_ROLE_LAND,
			Reason: "the builder added this basic land to reach the deck size",
		}
		deck.Cards = append(deck.Cards, dc)
		byOracle[b.GetOracleId()] = dc
	}
	return short
}

// missingLocked names the cards the user said to keep that the deck does
// not hold. A commander counts as held: it is in the deck, in the command
// zone. The names come from the pool, so the message reads as the user
// wrote them (D-242).
func missingLocked(deck *mtgv1.Deck, req Request) []string {
	if len(req.Locked) == 0 {
		return nil
	}
	have := map[string]bool{}
	for _, c := range deck.GetCards() {
		have[c.GetOracleId()] = true
	}
	for _, c := range deck.GetSideboard() {
		have[c.GetOracleId()] = true
	}
	for _, id := range deck.GetCommanderOracleIds() {
		have[id] = true
	}
	var out []string
	for _, id := range req.Locked {
		if have[id] {
			continue
		}
		name := id
		for _, n := range req.Pool.Names() {
			if c, ok := req.Pool.Card(n); ok && c.GetOracleId() == id {
				name = c.GetName()
				break
			}
		}
		out = append(out, name)
	}
	return out
}
