package profile

import (
	"math/rand/v2"
	"sort"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The goldfish simulation (D-453). It deals opening hands, mulligans by
// the London rule, and plays the lands and the mana producers on curve.
// No opponent, no spell but the ramp, and no color. It answers the turn
// the deck casts its commander, the mana on turn four, and the share of
// first hands with two to four lands.
//
// The mulligan rule follows Karsten 2022: in Commander a first seven with
// zero, one, two, six, or seven lands goes back for free, and a regular
// seven with zero, one, six, or seven lands goes back. The second
// mulligan keeps seven and puts one card on the bottom.

// DefaultHands is how many games one profile deals (D-453).
const DefaultHands = 10000

// lastTurn is the turn the simulation stops at. A commander not cast by
// then reads as cast on the turn after.
const lastTurn = 10

type simKind uint8

const (
	kindSpell simKind = iota
	kindLand
	kindRock
	kindDork
	kindLandRamp
)

// simCard is what the simulation knows of one card.
type simCard struct {
	kind   simKind
	mv     int
	tapped bool
	// mana is what a rock or a dork adds per turn.
	mana int
}

func simCardOf(c *mtgv1.Card) simCard {
	mv := int(c.GetManaValue())
	switch {
	case isLand(c):
		return simCard{kind: kindLand, tapped: entersTapped(c)}
	case isLandRamp(c):
		return simCard{kind: kindLandRamp, mv: mv, mana: 1}
	case isCreature(c) && producesMana(c):
		return simCard{kind: kindDork, mv: mv, mana: manaMade(c)}
	case producesMana(c):
		return simCard{kind: kindRock, mv: mv, mana: manaMade(c)}
	}
	return simCard{kind: kindSpell, mv: mv}
}

// simInput is one deck for the simulation.
type simInput struct {
	cards []simCard
	// commander is true for a Commander deck: the free mulligan, the
	// draw on turn one, and the commander in the command zone.
	commander bool
	// commanderMV is the mana value of the commander, the cheaper of
	// two.
	commanderMV int
	hands       int
	seed        uint64
}

// simResult is what the simulation measured.
type simResult struct {
	hands          int
	commanderTurn  float64
	manaTurnFour   float64
	shareTwoToFour float64
}

func simulate(in simInput) simResult {
	if in.hands <= 0 || len(in.cards) < 7 {
		return simResult{}
	}
	r := rand.New(rand.NewPCG(in.seed, in.seed^0x9e3779b97f4a7c15))
	deck := make([]simCard, len(in.cards))
	var turnSum, manaSum float64
	var goodHands int
	for range in.hands {
		copy(deck, in.cards)
		r.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
		hand, next, good := openingHand(deck, in.commander)
		if good {
			goodHands++
		}
		turn, mana := play(hand, deck[next:], in)
		turnSum += float64(turn)
		manaSum += float64(mana)
	}
	n := float64(in.hands)
	return simResult{
		hands:          in.hands,
		commanderTurn:  turnSum / n,
		manaTurnFour:   manaSum / n,
		shareTwoToFour: float64(goodHands) / n,
	}
}

func countLands(hand []simCard) int {
	n := 0
	for _, c := range hand {
		if c.kind == kindLand {
			n++
		}
	}
	return n
}

// openingHand deals the hand under the mulligan rule. It returns the
// hand, the index of the next card of the library, and whether the very
// first seven held two to four lands.
func openingHand(deck []simCard, commander bool) (hand []simCard, next int, good bool) {
	first := deck[:7]
	lands := countLands(first)
	good = lands >= 2 && lands <= 4
	keep := func(l int, lo, hi int) bool { return l >= lo && l <= hi }
	// Each mulligan draws a fresh seven from the top of a new shuffle.
	// The library is one shuffle, so a fresh seven is the next seven.
	if commander {
		if keep(lands, 3, 5) {
			return append([]simCard(nil), first...), 7, good
		}
		second := deck[7:14]
		if keep(countLands(second), 2, 5) {
			return append([]simCard(nil), second...), 14, good
		}
		third := deck[14:21]
		return bottomOne(third), 21, good
	}
	if keep(lands, 2, 5) {
		return append([]simCard(nil), first...), 7, good
	}
	second := deck[7:14]
	return bottomOne(second), 14, good
}

// bottomOne puts one card of a seven on the bottom: a land when the hand
// holds more than four, and the dearest spell otherwise.
func bottomOne(seven []simCard) []simCard {
	hand := append([]simCard(nil), seven...)
	drop := -1
	if countLands(hand) > 4 {
		for i, c := range hand {
			if c.kind == kindLand {
				drop = i
				break
			}
		}
	} else {
		for i, c := range hand {
			if c.kind != kindLand && (drop < 0 || c.mv > hand[drop].mv) {
				drop = i
			}
		}
	}
	if drop < 0 {
		drop = 0
	}
	return append(hand[:drop], hand[drop+1:]...)
}

// play runs the turns. It returns the turn the commander was castable
// and the mana available on turn four.
func play(hand, library []simCard, in simInput) (commanderTurn, manaFour int) {
	var landsInPlay, rocks, dorksReady, dorksNew, rampLands, rampNew int
	commanderTurn = lastTurn + 1
	cast := false
	for turn := 1; turn <= lastTurn; turn++ {
		if (in.commander || turn > 1) && len(library) > 0 {
			hand = append(hand, library[0])
			library = library[1:]
		}
		// Play a land, an untapped one first.
		tappedNow := 0
		if i := pickLand(hand); i >= 0 {
			if hand[i].tapped {
				tappedNow = 1
			}
			hand = append(hand[:i], hand[i+1:]...)
			landsInPlay++
		}
		mana := landsInPlay - tappedNow + rocks + dorksReady + rampLands
		if turn == 4 {
			manaFour = mana
		}
		if in.commander && !cast && mana >= in.commanderMV {
			cast = true
			commanderTurn = turn
			mana -= in.commanderMV
		}
		// Cast the ramp the mana allows, cheapest first. A rock adds its
		// mana at once, and a dork or a land search adds from the next
		// turn.
		sort.SliceStable(hand, func(i, j int) bool { return hand[i].mv < hand[j].mv })
		for i := 0; i < len(hand); {
			c := hand[i]
			if c.kind == kindSpell || c.kind == kindLand || c.mv > mana {
				i++
				continue
			}
			mana -= c.mv
			switch c.kind {
			case kindRock:
				rocks += c.mana
				mana += c.mana
			case kindDork:
				dorksNew += c.mana
			case kindLandRamp:
				rampNew += c.mana
			}
			hand = append(hand[:i], hand[i+1:]...)
		}
		// A rock cast this turn can pay for the commander this turn.
		if in.commander && !cast && mana >= in.commanderMV {
			cast = true
			commanderTurn = turn
		}
		dorksReady += dorksNew
		dorksNew = 0
		rampLands += rampNew
		rampNew = 0
	}
	return commanderTurn, manaFour
}

// pickLand returns the index of the land to play: an untapped one when
// the hand holds one, else any land, else -1.
func pickLand(hand []simCard) int {
	first := -1
	for i, c := range hand {
		if c.kind != kindLand {
			continue
		}
		if !c.tapped {
			return i
		}
		if first < 0 {
			first = i
		}
	}
	return first
}
