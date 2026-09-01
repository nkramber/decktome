package generate

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	cardsets "github.com/nkramber/mtg-deck-builder/go/internal/cards"
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
// (D-37).
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
// nothing (candidates.go). A deck still needs them, so they join the pool
// as always cards (D-225).
var basicNames = map[mtgv1.Color]string{
	mtgv1.Color_COLOR_W: "Plains",
	mtgv1.Color_COLOR_U: "Island",
	mtgv1.Color_COLOR_B: "Swamp",
	mtgv1.Color_COLOR_R: "Mountain",
	mtgv1.Color_COLOR_G: "Forest",
}

// AllColors is the five colors in color order. A 60-card session with no
// color choice gets every basic (D-225).
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

// deckCard makes the entry for a card the builder inserts itself. The
// owned count and the price come from the pool, so a card the user owns
// is never charged as a purchase (D-37).
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
		if ok && candidates.IsBasicLand(c) {
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
// that trades a Swamp for a Forest has kept the precon. A precon card the
// pool does not hold still counts, because the deck can not keep what
// the model can not name, and the share must say so. The card index
// names a basic the pool does not hold, so a Wastes or a snow basic is
// never counted as a nonbasic.
func preconNonbasics(req Request, cards rules.CardSource) map[string]bool {
	out := make(map[string]bool, len(req.PreconOracleIDs))
	for _, id := range req.PreconOracleIDs {
		if c, ok := lookup(req.Pool, cards, id); ok && candidates.IsBasicLand(c) {
			continue
		}
		out[id] = true
	}
	return out
}

// lookup reads a card from the pool first and the card index second.
func lookup(pool *Pool, cards rules.CardSource, id string) (*mtgv1.Card, bool) {
	if pool != nil {
		if c, ok := pool.ByOracleID(id); ok {
			return c, true
		}
	}
	if cards != nil {
		return cards.ByOracleID(id)
	}
	return nil, false
}

// heldPrecon is the set of precon ids the deck holds, the commanders
// included: a commander is in the deck, in the command zone. It is a set,
// so two entries of one card count once.
func heldPrecon(deck *mtgv1.Deck, in map[string]bool) map[string]bool {
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
	return held
}

// MaxPreconSwap is the largest precon shortfall the builder closes
// itself. The model chose which cards to drop, and putting one back is
// what the user asked for. A larger gap means the model built a different
// deck, and that stays a finding for the user to see (D-250).
const MaxPreconSwap = 3

// swapBackPrecon puts precon cards back until the share is met, and
// returns how many it moved. It trades a card the precon does not hold
// for one it does, so the deck size does not change. The model can not
// count its own list reliably, so the builder finishes the job (D-250).
func swapBackPrecon(deck *mtgv1.Deck, req Request, cards rules.CardSource) int {
	in := preconNonbasics(req, cards)
	want := len(in)
	if want == 0 {
		return 0
	}
	held := heldPrecon(deck, in)
	short := PreconKeepCount(want) - len(held)
	if short <= 0 || short > MaxPreconSwap {
		return 0
	}
	// The cards to put back, in the precon's own order so the choice is
	// stable. A card the pool does not hold can not go back, and a card
	// outside the chosen commanders' color identity can not either (CR
	// 903.5c).
	identity := commanderIdentity(deck, req.Pool, cards)
	var missing []*mtgv1.Card
	for _, id := range req.PreconOracleIDs {
		if !in[id] || held[id] {
			continue
		}
		c, ok := req.Pool.ByOracleID(id)
		if !ok || (identity != nil && !candidates.IdentityFits(c.GetColorIdentity(), identity)) {
			continue
		}
		missing = append(missing, c)
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
		if pc, ok := req.Pool.ByOracleID(c.GetOracleId()); ok && candidates.IsBasicLand(pc) {
			continue
		}
		deck.Cards[i] = deckCard(req.Pool, missing[moved], 1, c.GetRole(),
			"the deck upgrades a precon, and this card is one the precon holds")
		moved++
	}
	return moved
}

// commanderIdentity is the union of the chosen commanders' color
// identities, as an allowed-color set. Nil means no commander is chosen,
// so no color test applies.
func commanderIdentity(deck *mtgv1.Deck, pool *Pool, cards rules.CardSource) map[mtgv1.Color]bool {
	var colors []mtgv1.Color
	found := false
	for _, id := range deck.GetCommanderOracleIds() {
		if c, ok := lookup(pool, cards, id); ok {
			found = true
			colors = append(colors, c.GetColorIdentity()...)
		}
	}
	if !found {
		return nil
	}
	set := candidates.ColorSet(colors)
	if set == nil {
		set = map[mtgv1.Color]bool{}
	}
	return set
}

// markOutsideSets sets DeckCard.outside_requested_sets on every card the
// reader's sets do not hold, and adds one warning that counts them
// (D-383). A deck with no set limit is left alone.
//
// The mark is written once, at build time, so it says what was true then
// and a later snapshot does not move it. A basic land is never marked:
// the mana base is out of a set limit (D-378).
func markOutsideSets(deck *mtgv1.Deck, req Request, cards rules.CardSource) {
	codes := cardsets.CodeSet(req.SetCodes)
	if codes == nil {
		return
	}
	var outside []string
	for _, list := range [][]*mtgv1.DeckCard{deck.GetCards(), deck.GetSideboard()} {
		for _, dc := range list {
			c, ok := lookup(req.Pool, cards, dc.GetOracleId())
			if !ok || candidates.IsBasicLand(c) {
				continue
			}
			if cardsets.InSets(c, codes) {
				continue
			}
			dc.OutsideRequestedSets = true
			outside = append(outside, dc.GetName())
		}
	}
	if len(outside) == 0 {
		return
	}
	sort.Strings(outside)
	addFinding(deck, CodeOutsideSet, mtgv1.Severity_SEVERITY_WARN,
		fmt.Sprintf("the sets you named do not hold %s: %s",
			plural(len(outside), "card"), strings.Join(outside, ", ")))
}

// trimToSize drops cards until the deck holds the exact size, and it
// returns the names it cut (D-391). It is the mirror of padWithBasics.
//
// A deck one or two cards over is a counting slip, the same slip D-225
// pads when it falls the other way. Revise gate run 3 saw the model
// remove six cards and add seven under a mana cap, and the engine
// blocked the whole deck for one card.
//
// It never cuts a commander, a locked card, a card the reader named, or
// an entry of more than one copy: those are the mana base and the
// instructions. Among the rest it drops the dearest card first, because
// a revision that lowers the curve is the case this arose in, and the
// cheapest cards are the ones the deck can least afford to lose.
func trimToSize(deck *mtgv1.Deck, req Request) []string {
	size := DeckSize(req.Format)
	if size == 0 {
		return nil
	}
	have := len(deck.GetCommanderOracleIds())
	for _, c := range deck.GetCards() {
		have += int(c.GetCount())
	}
	over := have - size
	if over <= 0 || over > MaxTrim {
		return nil
	}
	keep := map[string]bool{}
	for _, id := range req.Commanders {
		keep[id] = true
	}
	for _, id := range req.Locked {
		keep[id] = true
	}
	if req.Revision != nil {
		for _, id := range req.Revision.Exempt {
			keep[id] = true
		}
	}
	// The candidates, dearest first. A single copy only: an entry of two
	// or more is the mana base, and cutting one copy of it is a change
	// the reader did not ask for.
	type row struct {
		index int
		mv    float64
	}
	var rows []row
	for i, dc := range deck.GetCards() {
		if keep[dc.GetOracleId()] || dc.GetCount() != 1 {
			continue
		}
		c, ok := req.Pool.ByOracleID(dc.GetOracleId())
		if !ok || candidates.IsBasicLand(c) {
			continue
		}
		rows = append(rows, row{index: i, mv: c.GetManaValue()})
	}
	if len(rows) < over {
		return nil
	}
	sort.SliceStable(rows, func(i, j int) bool { return rows[i].mv > rows[j].mv })
	cut := map[int]bool{}
	var names []string
	for _, r := range rows[:over] {
		cut[r.index] = true
		names = append(names, deck.GetCards()[r.index].GetName())
	}
	kept := make([]*mtgv1.DeckCard, 0, len(deck.GetCards())-over)
	for i, dc := range deck.GetCards() {
		if !cut[i] {
			kept = append(kept, dc)
		}
	}
	deck.Cards = kept
	sort.Strings(names)
	return names
}
