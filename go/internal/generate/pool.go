package generate

import (
	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
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
//
// The always cards carry no owned count here. A caller with a collection
// uses FromListOwned, or the user's own precon cards count as purchases
// (G-3 of the 2026-08-28 audit).
func FromList(l *candidates.List, always []*mtgv1.Card, buyList bool) *Pool {
	return FromListOwned(l, always, nil, buyList)
}

// FromListOwned is FromList with the collection's owned count per oracle
// id. The shortlist rows carry their own counts, and the always cards
// read theirs from owned, so every card in the pool knows whether the
// user holds it.
func FromListOwned(l *candidates.List, always []*mtgv1.Card, ownedCounts map[string]int32, buyList bool) *Pool {
	cards := append([]*mtgv1.Card(nil), always...)
	owned := map[string]int32{}
	for _, c := range always {
		if n := ownedCounts[c.GetOracleId()]; n > 0 {
			owned[c.GetOracleId()] = n
		}
	}
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
				if n := ownedCounts[c.Partner.GetOracleId()]; n > 0 {
					owned[c.Partner.GetOracleId()] = n
				}
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

// AllColors is the five colors in color order. A 60-card session with no
// color choice gets every basic (G-7 of the 2026-08-28 audit).
var AllColors = []mtgv1.Color{
	mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B,
	mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G,
}

// BasicLands returns the basic lands of a color identity, in color order.
// find is the card lookup, so this package needs no card index.
//
// A colorless deck gets no basic. Wastes is the colorless basic, and the
// app builds no deck that needs it today.
func BasicLands(find func(string) (*mtgv1.Card, bool), colors []mtgv1.Color) []*mtgv1.Card {
	var out []*mtgv1.Card
	for _, col := range AllColors {
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

// deckCard makes the entry for a card the builder inserts itself. The
// owned count and the price come from the pool, so a card the user owns
// is never charged as a purchase (G-3 of the 2026-08-28 audit).
func deckCard(pool *Pool, c *mtgv1.Card, count int32, role mtgv1.CardRole, reason string) *mtgv1.DeckCard {
	owned := pool.OwnedCount(c.GetOracleId())
	return &mtgv1.DeckCard{
		OracleId:   c.GetOracleId(),
		Name:       c.GetName(),
		Count:      count,
		Role:       role,
		Reason:     reason,
		Owned:      owned >= count,
		OwnedCount: owned,
		PriceUsd:   c.GetPriceUsd(),
	}
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
			dc.Owned = dc.OwnedCount >= dc.Count
			continue
		}
		dc := deckCard(req.Pool, b, 1, mtgv1.CardRole_CARD_ROLE_LAND,
			"the builder added this basic land to reach the deck size")
		deck.Cards = append(deck.Cards, dc)
		byOracle[b.GetOracleId()] = dc
	}
	return short
}

// missingLocked names the cards the user said to keep that the deck does
// not hold. A commander counts as held: it is in the deck, in the command
// zone. The names come from the pool, so the message reads as the user
// wrote them (D-242).
func missingLocked(deck *mtgv1.Deck, req Request, cards rules.CardSource) []string {
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
		// The pool may have dropped the card, for example when a revision
		// removed it, so the card index names it then (D-301).
		name := id
		if c, ok := req.Pool.ByOracleID(id); ok {
			name = c.GetName()
		} else if cards != nil {
			if c, ok := cards.ByOracleID(id); ok {
				name = c.GetName()
			}
		}
		out = append(out, name)
	}
	return out
}

// preconNonbasics is the set of the precon's nonbasic oracle ids. The
// share rule of D-218 measures these: basic lands swap free, so a deck
// that trades a Swamp for a Forest has kept the precon (A-5 of the
// 2026-08-28 audit). A precon card the pool does not hold still counts,
// because the deck can not keep what the model can not name, and the
// share must say so.
func preconNonbasics(req Request) map[string]bool {
	out := make(map[string]bool, len(req.PreconOracleIDs))
	for _, id := range req.PreconOracleIDs {
		if c, ok := req.Pool.ByOracleID(id); ok && IsBasic(c) {
			continue
		}
		out[id] = true
	}
	return out
}

// MaxPreconSwap is the largest precon shortfall the builder closes
// itself. The model chose which cards to drop, and putting one back is
// what the user asked for. A larger gap means the model built a different
// deck, and that stays a finding for the user to see (D-250).
const MaxPreconSwap = 3

// swapBackPrecon puts precon cards back until the share is met, and
// returns how many it moved. It trades a card the precon does not hold
// for one it does, so the deck size does not change.
//
// The model repaired prompt 17 of 2026-08-28 to 67 of the 68 it needed,
// read the finding that said so, and returned 67 again. It can not count
// its own list reliably, so the builder finishes the job.
func swapBackPrecon(deck *mtgv1.Deck, req Request) int {
	in := preconNonbasics(req)
	want := len(in)
	if want == 0 {
		return 0
	}
	held := map[string]bool{}
	for _, c := range deck.GetCards() {
		if in[c.GetOracleId()] {
			held[c.GetOracleId()] = true
		}
	}
	for _, id := range deck.GetCommanderOracleIds() {
		if in[id] {
			held[id] = true
		}
	}
	short := PreconKeepCount(want) - len(held)
	if short <= 0 || short > MaxPreconSwap {
		return 0
	}
	// The cards to put back, in the precon's own order so the choice is
	// stable. A card the pool does not hold can not go back.
	var missing []*mtgv1.Card
	for _, id := range req.PreconOracleIDs {
		if !in[id] || held[id] {
			continue
		}
		if c, ok := req.Pool.ByOracleID(id); ok {
			missing = append(missing, c)
		}
	}
	locked := map[string]bool{}
	for _, id := range req.Locked {
		locked[id] = true
	}
	moved := 0
	for i := len(deck.Cards) - 1; i >= 0 && moved < short && moved < len(missing); i-- {
		c := deck.Cards[i]
		// A card the precon holds stays. So does a locked card, a basic
		// land, and any entry of more than one copy: those are the mana
		// base, and swapping one would break the deck size.
		if in[c.GetOracleId()] || locked[c.GetOracleId()] || c.GetCount() != 1 {
			continue
		}
		if pc, ok := req.Pool.ByOracleID(c.GetOracleId()); ok && IsBasic(pc) {
			continue
		}
		deck.Cards[i] = deckCard(req.Pool, missing[moved], 1, c.GetRole(),
			"the deck upgrades a precon, and this card is one the precon holds")
		moved++
	}
	return moved
}
