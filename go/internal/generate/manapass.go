package generate

import (
	"math"
	"sort"

	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/profile"
)

// The deterministic mana pass of PR-33 (F-78, Part 3).
//
// Every band the profile measures is a number this app computes itself,
// and the goldfish simulation the mana bands rest on takes about 10
// milliseconds. Before this the app moved those numbers with a repair
// call of 60 to 90 seconds, twice, and the numbers did not converge:
// session 833r7UccvAqFsyYJzHfz read two findings before the repair and
// two after, and the third call met the build cap.
//
// The pass makes one step at a time, measures again, and keeps the step
// only when the deck moved closer to its bands. So a step never trades
// one off-band feature for another.

// maxManaSteps bounds the pass. Each step measures the whole deck, so
// the bound is what keeps a pathological pool from a long loop. Twelve
// steps move about a third of a Commander mana base.
const maxManaSteps = 12

// fixMana moves the mana base of a built deck until every mana feature
// sits in band. It calls no model. It returns the count of steps it
// kept, and the deck it changed is the one the caller holds.
//
// It runs before the profile is read, so the stored profile and every
// finding describe the deck after the pass.
func (b *Builder) fixMana(req Request, deck *mtgv1.Deck) int {
	// A revision and an upgrade keep the deck they were given. The
	// reader asked for a change, or the precon is a working deck, and a
	// mana pass over either one is a second author (D-249, PR-12B).
	if b.profiler == nil || req.Revision != nil || req.Precon != "" {
		return 0
	}
	basics := b.basicsOf(deck)
	if len(basics) == 0 {
		return 0
	}
	score := b.manaScore(deck)
	kept := 0
	for range maxManaSteps {
		if score == 0 {
			break
		}
		next, ok := b.bestManaStep(req, deck, basics, score)
		if !ok {
			break
		}
		next.apply(deck)
		score = b.manaScore(deck)
		kept++
	}
	if kept > 0 {
		b.log.Info("the mana pass moved the deck toward its bands with no model call",
			"session", req.SessionID, "steps", kept, "off_band", score)
	}
	return kept
}

// FixManaForCheck runs the pass from the free lane of `deck-gate
// -manapass`, which reads the stored decks of a gate document. The lane
// is how PR-33 proves the pass against real decks, and it calls no
// provider (F-78, Part 3).
func (b *Builder) FixManaForCheck(req Request, deck *mtgv1.Deck) int {
	return b.fixMana(req, deck)
}

// manaScore is how far the deck sits outside its bands, summed over
// every feature. Zero means every band holds. A step is kept only when
// this number falls, so a fix of one band never breaks another.
func (b *Builder) manaScore(deck *mtgv1.Deck) float64 {
	prof := b.profiler.Measure(deck, b.cards)
	total := 0.0
	for _, f := range prof.GetFeatures() {
		total += bandDistance(f)
	}
	return total
}

// bandDistance is how far one feature sits outside its band, as a share
// of the band's own scale. A feature inside its band is zero.
//
// The share is what makes the features comparable: the land count runs
// to 40 and the opening-hand share runs to 1, and a raw distance would
// let one land outweigh a broken mana base.
func bandDistance(f *mtgv1.ProfileFeature) float64 {
	if !f.GetOffBand() {
		return 0
	}
	scale := math.Max(math.Abs(f.GetLow()), 1)
	if f.GetHasHigh() {
		scale = math.Max(math.Abs(f.GetHigh()), scale)
	}
	switch {
	case f.GetValue() < f.GetLow():
		return (f.GetLow() - f.GetValue()) / scale
	case f.GetHasHigh() && f.GetValue() > f.GetHigh():
		return (f.GetValue() - f.GetHigh()) / scale
	}
	return 0
}

// manaStep is one change to the deck: a card to add and a card to drop,
// by Oracle id. Either half may be empty, which makes the step a plain
// addition or a plain cut.
type manaStep struct {
	add  *mtgv1.Card
	drop string
	role mtgv1.CardRole
}

// apply writes the step into the deck.
func (s manaStep) apply(deck *mtgv1.Deck) {
	if s.drop != "" {
		deck.Cards = dropOne(deck.GetCards(), s.drop)
	}
	if s.add != nil {
		deck.Cards = addOne(deck.GetCards(), s.add, s.role)
	}
}

// bestManaStep measures every candidate step and returns the one that
// moves the deck closest to its bands. ok is false when no step helps.
func (b *Builder) bestManaStep(req Request, deck *mtgv1.Deck, basics []*mtgv1.Card, score float64) (manaStep, bool) {
	best, bestScore, found := manaStep{}, score, false
	for _, step := range b.manaCandidates(req, deck, basics) {
		undo := snapshot(deck)
		step.apply(deck)
		got := b.manaScore(deck)
		restore(deck, undo)
		if got < bestScore {
			best, bestScore, found = step, got, true
		}
	}
	return best, found
}

// manaCandidates lists the steps the pass may take, in the order it
// prefers them. Each one keeps the deck size, so no step can make a deck
// the engine refuses.
//
// Three levers. An untapped land of the pool replaces a tapped land of
// the deck. A basic land of one color replaces a basic land of another,
// which is how the pass answers a color the sources do not cover. A
// basic land replaces the costliest nonland card, or the reverse, which
// is how it moves the land count inside its band.
func (b *Builder) manaCandidates(req Request, deck *mtgv1.Deck, basics []*mtgv1.Card) []manaStep {
	var out []manaStep
	inDeck := map[string]bool{}
	for _, dc := range deck.GetCards() {
		inDeck[dc.GetOracleId()] = true
	}
	tapped, untapped := b.landsOf(deck, req)
	for _, drop := range tapped {
		for _, add := range untapped {
			if inDeck[add.GetOracleId()] {
				continue
			}
			out = append(out, manaStep{add: add, drop: drop, role: mtgv1.CardRole_CARD_ROLE_LAND})
		}
	}
	// A basic for another basic. The deck holds more than one of each,
	// so neither half needs to be absent from the deck.
	for _, add := range basics {
		for _, other := range basics {
			if add.GetOracleId() == other.GetOracleId() || !inDeck[other.GetOracleId()] {
				continue
			}
			out = append(out, manaStep{add: add, drop: other.GetOracleId(), role: mtgv1.CardRole_CARD_ROLE_LAND})
		}
	}
	// A basic land for the costliest spell, and the reverse. The land
	// band bounds both: a step that leaves the band scores worse and the
	// caller refuses it, so no rule here repeats the band.
	dear, dearRole, dearMV := costliestSpell(deck, b.cards)
	if dear != "" {
		out = append(out, manaStep{add: basics[0], drop: dear, role: mtgv1.CardRole_CARD_ROLE_LAND})
	}
	if cut := b.spareBasic(deck, basics); cut != "" {
		if add := b.cheapestSpell(req, deck); add != nil {
			out = append(out, manaStep{add: add, drop: cut, role: mtgv1.CardRole_CARD_ROLE_SYNERGY})
		}
	}
	// A cheaper card of the same job for the costliest one. This is the
	// lever of the curve: the average mana value of a high bracket runs
	// low, and no land moves it. The free lane of PR-33 read the first
	// pass closing 2 of 10 off-band features without this step.
	for _, add := range b.cheaperOfRole(req, inDeck, dearRole, dearMV) {
		out = append(out, manaStep{add: add, drop: dear, role: dearRole})
	}
	// Ramp of a low mana value for the costliest card. The mana of turn
	// four reads ramp more than it reads one more land, and a deck at
	// the top of its land band has no land step left.
	for _, add := range b.rampOf(req, inDeck) {
		out = append(out, manaStep{add: add, drop: dear, role: mtgv1.CardRole_CARD_ROLE_RAMP})
	}
	return out
}

// cheaperOfRole is the pool's cheapest cards of one job that cost less
// than the card the step drops. The deck keeps its shape: a removal
// spell leaves for a cheaper removal spell.
func (b *Builder) cheaperOfRole(req Request, inDeck map[string]bool,
	role mtgv1.CardRole, mv float64,
) []*mtgv1.Card {
	if role == mtgv1.CardRole_CARD_ROLE_UNSPECIFIED || mv <= 1 {
		return nil
	}
	var out []*mtgv1.Card
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || profile.IsLand(c) || inDeck[c.GetOracleId()] || c.GetManaValue() >= mv {
			continue
		}
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GetManaValue() < out[j].GetManaValue() })
	if len(out) > maxSpellCandidates {
		out = out[:maxSpellCandidates]
	}
	return out
}

// rampOf is the pool's cheapest nonland mana producers the deck does not
// hold.
func (b *Builder) rampOf(req Request, inDeck map[string]bool) []*mtgv1.Card {
	var out []*mtgv1.Card
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || profile.IsLand(c) || inDeck[c.GetOracleId()] || !profile.ProducesMana(c) {
			continue
		}
		if c.GetManaValue() > maxRampManaValue {
			continue
		}
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GetManaValue() < out[j].GetManaValue() })
	if len(out) > maxSpellCandidates {
		out = out[:maxSpellCandidates]
	}
	return out
}

// maxSpellCandidates bounds the spells one step measures, and
// maxRampManaValue is the cost above which a mana producer no longer
// helps the mana of turn four.
const (
	maxSpellCandidates = 6
	maxRampManaValue   = 3
)

// landsOf splits the deck's lands from the pool's: the tapped lands the
// deck holds, by Oracle id, and the untapped lands of the pool that
// could replace one.
func (b *Builder) landsOf(deck *mtgv1.Deck, req Request) (tapped []string, untapped []*mtgv1.Card) {
	for _, dc := range deck.GetCards() {
		c, ok := b.cards.ByOracleID(dc.GetOracleId())
		if !ok || !profile.IsLand(c) || profile.IsBasic(c) {
			continue
		}
		if profile.EntersTapped(c) {
			tapped = append(tapped, dc.GetOracleId())
		}
	}
	if len(tapped) == 0 {
		return nil, nil
	}
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || !profile.IsLand(c) || profile.IsBasic(c) || profile.EntersTapped(c) {
			continue
		}
		untapped = append(untapped, c)
	}
	// The pool is long, so the pass reads the untapped lands it ranks
	// first. Every candidate costs one measurement of the whole deck.
	if len(untapped) > maxLandCandidates {
		untapped = untapped[:maxLandCandidates]
	}
	return tapped, untapped
}

// maxLandCandidates bounds the untapped lands one step measures. The
// pool holds hundreds, and each candidate costs a measurement.
const maxLandCandidates = 8

// basicsOf is the basic lands of the deck's colors, from the pool the
// build used. The always list of the pool holds them (D-225).
func (b *Builder) basicsOf(deck *mtgv1.Deck) []*mtgv1.Card {
	var out []*mtgv1.Card
	seen := map[string]bool{}
	for _, dc := range deck.GetCards() {
		c, ok := b.cards.ByOracleID(dc.GetOracleId())
		if !ok || !profile.IsBasic(c) || seen[c.GetOracleId()] {
			continue
		}
		seen[c.GetOracleId()] = true
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GetName() < out[j].GetName() })
	return out
}

// costliestSpell is the Oracle id of the deck's dearest nonland card. It
// is the card a step trades for a land, because the curve band and the
// mana bands both read it.
func costliestSpell(deck *mtgv1.Deck, src cardSource) (id string, role mtgv1.CardRole, mv float64) {
	mv = -1
	for _, dc := range deck.GetCards() {
		c, ok := src.ByOracleID(dc.GetOracleId())
		if !ok || profile.IsLand(c) {
			continue
		}
		if v := c.GetManaValue(); v > mv {
			mv, id, role = v, dc.GetOracleId(), dc.GetRole()
		}
	}
	return id, role, mv
}

// cheapestSpell is the pool's cheapest card the deck does not hold. A
// step adds it when the land count sits above its band.
func (b *Builder) cheapestSpell(req Request, deck *mtgv1.Deck) *mtgv1.Card {
	inDeck := map[string]bool{}
	for _, dc := range deck.GetCards() {
		inDeck[dc.GetOracleId()] = true
	}
	var best *mtgv1.Card
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || profile.IsLand(c) || inDeck[c.GetOracleId()] {
			continue
		}
		if best == nil || c.GetManaValue() < best.GetManaValue() {
			best = c
		}
	}
	return best
}

// spareBasic is a basic land the deck holds more than one of. A step
// cuts one of those, so no color loses its last source to the cut.
func (b *Builder) spareBasic(deck *mtgv1.Deck, basics []*mtgv1.Card) string {
	for _, c := range basics {
		for _, dc := range deck.GetCards() {
			if dc.GetOracleId() == c.GetOracleId() && dc.GetCount() > 1 {
				return dc.GetOracleId()
			}
		}
	}
	return ""
}

// dropOne removes one copy of a card. An entry of one copy leaves.
func dropOne(list []*mtgv1.DeckCard, id string) []*mtgv1.DeckCard {
	out := make([]*mtgv1.DeckCard, 0, len(list))
	done := false
	for _, dc := range list {
		if !done && dc.GetOracleId() == id {
			done = true
			if dc.GetCount() > 1 {
				cp := cloneCard(dc)
				cp.Count--
				out = append(out, cp)
			}
			continue
		}
		out = append(out, dc)
	}
	return out
}

// addOne adds one copy of a card, as a new entry or on the entry the
// deck holds.
func addOne(list []*mtgv1.DeckCard, c *mtgv1.Card, role mtgv1.CardRole) []*mtgv1.DeckCard {
	out := make([]*mtgv1.DeckCard, 0, len(list)+1)
	done := false
	for _, dc := range list {
		if !done && dc.GetOracleId() == c.GetOracleId() {
			done = true
			cp := cloneCard(dc)
			cp.Count++
			out = append(out, cp)
			continue
		}
		out = append(out, dc)
	}
	if done {
		return out
	}
	return append(out, &mtgv1.DeckCard{
		OracleId: c.GetOracleId(), Name: c.GetName(), Count: 1, Role: role,
		Reason: "the mana pass added it to bring the mana base inside the power level",
	})
}

// cloneCard copies one entry, so a trial step never writes on the entry
// the deck holds. A proto message carries a lock, so the copy goes
// through proto.Clone and never through an assignment.
func cloneCard(dc *mtgv1.DeckCard) *mtgv1.DeckCard {
	return proto.Clone(dc).(*mtgv1.DeckCard) //nolint:errcheck,forcetypeassert // Clone of a DeckCard is a DeckCard
}

// snapshot and restore hold the card list of a trial step. Every step
// runs on the deck itself, because the profiler reads a whole deck.
func snapshot(deck *mtgv1.Deck) []*mtgv1.DeckCard {
	return append([]*mtgv1.DeckCard(nil), deck.GetCards()...)
}

func restore(deck *mtgv1.Deck, list []*mtgv1.DeckCard) { deck.Cards = list }

// cardSource is the card reader the pass needs. rules.CardSource is the
// same shape, and this names the one method the pass uses.
type cardSource interface {
	ByOracleID(id string) (*mtgv1.Card, bool)
}
