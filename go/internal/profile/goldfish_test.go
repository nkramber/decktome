package profile

import (
	"testing"
)

func deckSim(lands, tappedLands, rocks, spells int, spellMV int) []simCard {
	var out []simCard
	for range lands {
		out = append(out, simCard{kind: kindLand})
	}
	for range tappedLands {
		out = append(out, simCard{kind: kindLand, tapped: true})
	}
	for range rocks {
		out = append(out, simCard{kind: kindRock, mv: 2, mana: 1})
	}
	for range spells {
		out = append(out, simCard{kind: kindSpell, mv: spellMV})
	}
	return out
}

func TestSimulateIsDeterministicAndReadsTheShape(t *testing.T) {
	in := simInput{cards: deckSim(37, 0, 10, 52, 3), commander: true, commanderMV: 4, hands: 2000, seed: 7}
	a := simulate(in)
	b := simulate(in)
	if a != b {
		t.Errorf("two runs differ: %+v and %+v", a, b)
	}
	// 37 lands and 10 rocks: the four-drop commander comes down on turn
	// three or four on average, four or more mana on turn four, and most
	// first hands hold two to four lands.
	if a.commanderTurn < 3 || a.commanderTurn > 4.2 {
		t.Errorf("commander turn %v", a.commanderTurn)
	}
	if a.manaTurnFour < 4 || a.manaTurnFour > 6 {
		t.Errorf("mana on turn four %v", a.manaTurnFour)
	}
	if a.shareTwoToFour < 0.7 || a.shareTwoToFour > 0.9 {
		t.Errorf("share of two to four lands %v", a.shareTwoToFour)
	}
	if a.hands != 2000 {
		t.Errorf("hands %d", a.hands)
	}
}

func TestSimulateSeesFewerLandsAndTappedLands(t *testing.T) {
	full := simulate(simInput{cards: deckSim(37, 0, 10, 52, 3), commander: true, commanderMV: 4, hands: 2000, seed: 1})
	thin := simulate(simInput{cards: deckSim(28, 0, 10, 61, 3), commander: true, commanderMV: 4, hands: 2000, seed: 1})
	tapped := simulate(simInput{cards: deckSim(17, 20, 10, 52, 3), commander: true, commanderMV: 4, hands: 2000, seed: 1})
	if thin.manaTurnFour >= full.manaTurnFour || thin.commanderTurn <= full.commanderTurn {
		t.Errorf("a thin mana base must be slower: full %+v thin %+v", full, thin)
	}
	if tapped.manaTurnFour >= full.manaTurnFour {
		t.Errorf("tapped lands must cost mana: full %+v tapped %+v", full, tapped)
	}
	if thin.shareTwoToFour >= full.shareTwoToFour {
		t.Errorf("fewer lands give fewer good hands: full %+v thin %+v", full, thin)
	}
}

func TestSimulateSixtyCardHasNoCommanderAndNoFreeMulligan(t *testing.T) {
	res := simulate(simInput{cards: deckSim(24, 0, 0, 36, 3), commander: false, hands: 1000, seed: 3})
	if res.commanderTurn != lastTurn+1 {
		t.Errorf("a 60-card deck casts no commander, got turn %v", res.commanderTurn)
	}
	if res.manaTurnFour < 3 || res.manaTurnFour > 4 {
		t.Errorf("mana on turn four %v", res.manaTurnFour)
	}
}

func TestSimulateRefusesATinyDeck(t *testing.T) {
	if res := simulate(simInput{cards: deckSim(3, 0, 0, 2, 1), hands: 10}); res.hands != 0 {
		t.Errorf("a deck under seven cards must not run, got %+v", res)
	}
	if res := simulate(simInput{cards: deckSim(30, 0, 0, 30, 1), hands: 0}); res.hands != 0 {
		t.Errorf("zero hands must not run, got %+v", res)
	}
}

func TestOpeningHandMulliganRules(t *testing.T) {
	land := simCard{kind: kindLand}
	spell := simCard{kind: kindSpell, mv: 3}
	// First seven holds one land: a Commander deck goes to the second
	// seven, which holds two lands and is kept.
	deck := []simCard{land, spell, spell, spell, spell, spell, spell,
		land, land, spell, spell, spell, spell, spell}
	hand, next, good := openingHand(deck, true)
	if len(hand) != 7 || next != 14 || countLands(hand) != 2 || good {
		t.Errorf("commander mulligan: %d cards, next %d, lands %d, good %v", len(hand), next, countLands(hand), good)
	}
	// A 60-card deck keeps a two-land seven at once.
	deck60 := append([]simCard{land, land, spell, spell, spell, spell, spell}, deck...)
	hand, next, good = openingHand(deck60, false)
	if len(hand) != 7 || next != 7 || !good {
		t.Errorf("sixty keep: %d cards, next %d, good %v", len(hand), next, good)
	}
	// The third seven of a Commander deck loses one card: a land when
	// it holds five or more, else the dearest spell.
	deck = append([]simCard{spell, spell, spell, spell, spell, spell, spell,
		spell, spell, spell, spell, spell, spell, spell},
		land, land, land, land, land, spell, spell)
	hand, next, _ = openingHand(deck, true)
	if len(hand) != 6 || next != 21 || countLands(hand) != 4 {
		t.Errorf("third seven: %d cards, next %d, lands %d", len(hand), next, countLands(hand))
	}
	six := bottomOne([]simCard{land, land, spell, {kind: kindSpell, mv: 6}, spell, spell, spell})
	for _, c := range six {
		if c.mv == 6 {
			t.Error("the dearest spell must go to the bottom")
		}
	}
}

func TestPlayRampAddsMana(t *testing.T) {
	land := simCard{kind: kindLand}
	// Three lands and a Sol Ring in hand: turn one plays a land and Sol
	// Ring, so turn four reads four lands and two from the ring.
	hand := []simCard{land, land, land, land, {kind: kindRock, mv: 1, mana: 2}, {kind: kindSpell, mv: 5}, {kind: kindSpell, mv: 5}}
	library := make([]simCard, 20)
	for i := range library {
		library[i] = simCard{kind: kindSpell, mv: 4}
	}
	turn, mana := play(hand, library, simInput{commander: true, commanderMV: 2})
	if mana != 6 {
		t.Errorf("mana on turn four %d, want 6", mana)
	}
	if turn != 1 {
		t.Errorf("a two-drop commander with a land and Sol Ring on turn one comes down on turn 1, got %d", turn)
	}
	// A three-drop waits for turn two: the ring leaves two mana on turn one.
	hand = []simCard{land, land, land, land, {kind: kindRock, mv: 1, mana: 2}, {kind: kindSpell, mv: 5}, {kind: kindSpell, mv: 5}}
	if turn, _ = play(hand, library, simInput{commander: true, commanderMV: 3}); turn != 2 {
		t.Errorf("a three-drop commander comes down on turn 2, got %d", turn)
	}
	// A dork adds from the next turn, and a tapped land gives nothing
	// the turn it lands.
	hand = []simCard{{kind: kindLand, tapped: true}, {kind: kindLand, tapped: true}, {kind: kindDork, mv: 1, mana: 1}}
	turn, mana = play(hand, library, simInput{commander: true, commanderMV: 9})
	// Turn one: a tapped land, no mana. Turn two: two lands, one
	// untapped now, one tapped, one mana, the dork is cast. Turn three: two
	// lands and the dork, three. Turn four: draws are spells, so three.
	if mana != 3 || turn != lastTurn+1 {
		t.Errorf("tapped and dork: mana %d turn %d", mana, turn)
	}
}
