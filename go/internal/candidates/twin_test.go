package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// twinCard is a card as Scryfall printed it on 2026-10-13.
func twinCard(id, name, cost, typeLine, text, power, toughness string) *mtgv1.Card {
	return &mtgv1.Card{OracleId: id, Name: name, ManaCost: cost, TypeLine: typeLine, OracleText: text, Power: power, Toughness: toughness}
}

var (
	grizzlyBears   = twinCard("grizzly", "Grizzly Bears", "{1}{G}", "Creature — Bear", "", "2", "2")
	balduvianBears = twinCard("balduvian", "Balduvian Bears", "{1}{G}", "Creature — Bear", "", "2", "2")
	runeclawBear   = twinCard("runeclaw", "Runeclaw Bear", "{1}{G}", "Creature — Bear", "", "2", "2")
	spectralBears  = twinCard("spectral", "Spectral Bears", "{1}{G}", "Creature — Bear Spirit",
		"Whenever this creature attacks, if defending player controls no black nontoken permanents, it doesn't untap during your next untap step.", "3", "3")
	ashcoatBear    = twinCard("ashcoat", "Ashcoat Bear", "{1}{G}", "Creature — Bear", "Flash (You may cast this spell any time you could cast an instant.)", "2", "2")
	llanowarElves  = twinCard("llanowar", "Llanowar Elves", "{G}", "Creature — Elf Druid", "{T}: Add {G}.", "1", "1")
	fyndhornElves  = twinCard("fyndhorn", "Fyndhorn Elves", "{G}", "Creature — Elf Druid", "{T}: Add {G}.", "1", "1")
	alpineGrizzly  = twinCard("alpine", "Alpine Grizzly", "{2}{G}", "Creature — Bear", "", "4", "2")
	smallerGrizzly = twinCard("smaller", "Smaller Grizzly", "{2}{G}", "Creature — Bear", "", "3", "2")
)

// TestTwinBeats reads the rule of D-1226 on the cards of F-232.
func TestTwinBeats(t *testing.T) {
	cases := []struct {
		owned, bought *mtgv1.Card
		want          bool
		why           string
	}{
		{balduvianBears, grizzlyBears, true, "the bears of F-232 are one card under two names"},
		{runeclawBear, grizzlyBears, true, "a third name of the same card"},
		{fyndhornElves, llanowarElves, true, "the elves of base replay 1 are one card under two names"},
		{grizzlyBears, grizzlyBears, false, "a card is no twin of itself"},
		{spectralBears, grizzlyBears, false, "another type line and a drawback make no twin"},
		{ashcoatBear, grizzlyBears, false, "other rules text makes no twin, even a better one"},
		{smallerGrizzly, alpineGrizzly, false, "less power makes no twin"},
		{alpineGrizzly, smallerGrizzly, true, "more power at the same cost is a twin"},
		{balduvianBears, alpineGrizzly, false, "another mana cost makes no twin"},
	}
	for _, c := range cases {
		if got := TwinBeats(c.owned, c.bought); got != c.want {
			t.Errorf("TwinBeats(%s, %s) = %v, want %v: %s", c.owned.Name, c.bought.Name, got, c.want, c.why)
		}
	}
}

// TestTwinKeyFoldsTheOwnName reads a card that names itself. Two cards
// with the same text under their own names are twins.
func TestTwinKeyFoldsTheOwnName(t *testing.T) {
	a := twinCard("a", "Kin of Ayula, the First", "{G}", "Legendary Creature — Bear", "When Kin enters, draw a card.", "1", "1")
	a.OracleText = "When Kin of Ayula, the First enters, draw a card. Kin of Ayula gets +1/+1."
	b := twinCard("b", "Kin of Beorn, the Second", "{G}", "Legendary Creature — Bear", "", "1", "1")
	b.OracleText = "When Kin of Beorn, the Second enters, draw a card. Kin of Beorn gets +1/+1."
	if !TwinBeats(a, b) {
		t.Error("two cards that differ in their own names alone are no twins")
	}
}

// TestTwinKeyRefuses reads the cards that have no twin.
func TestTwinKeyRefuses(t *testing.T) {
	forest := twinCard("forest", "Forest", "", "Basic Land — Forest", "({T}: Add {G}.)", "", "")
	forest.Supertypes, forest.CardTypes = []string{"Basic"}, []string{"Land"}
	two := twinCard("mdfc", "Two Faces", "{G}", "Creature — Bear // Land", "", "2", "2")
	two.Faces = []*mtgv1.CardFace{{Name: "A"}, {Name: "B"}}
	for _, c := range []*mtgv1.Card{forest, two, {OracleId: "x", Name: "No Type"}} {
		if _, ok := TwinKey(c); ok {
			t.Errorf("TwinKey(%s) gave a key, want none", c.Name)
		}
	}
}

// TestAddOwnedTwins proves that an owned twin of an upgrade joins the
// main list past the caps, with the role of the upgrade (F-232). Base
// replays 2 to 4 bought Llanowar Elves, and the role cap kept the owned
// Fyndhorn Elves off the list.
func TestAddOwnedTwins(t *testing.T) {
	ramp, synergy := mtgv1.CardRole_CARD_ROLE_RAMP, mtgv1.CardRole_CARD_ROLE_SYNERGY
	cand := func(c *mtgv1.Card, role mtgv1.CardRole, owned int32, score float64) Candidate {
		return Candidate{Card: c, Role: role, Owned: owned, Score: score}
	}
	mystic := twinCard("mystic", "Elvish Mystic", "{G}", "Creature — Elf Druid", "{T}: Add {G}.", "1", "1")
	scored := []Candidate{
		cand(llanowarElves, ramp, 0, 0.9),
		cand(mystic, ramp, 1, 0.8),
		cand(fyndhornElves, ramp, 1, 0.3),
		cand(grizzlyBears, synergy, 0, 0.5),
		cand(balduvianBears, synergy, 2, 0.4),
		cand(ashcoatBear, synergy, 1, 0.4),
	}
	main := []Candidate{scored[1], scored[4], scored[5]}
	upgrades := []Candidate{scored[0], scored[3]}
	got, added := addOwnedTwins(main, upgrades, scored)
	if added != 1 {
		t.Fatalf("added %d twins, want 1 (Fyndhorn Elves): %v", added, names(got))
	}
	tw, ok := find(got, "Fyndhorn Elves")
	if !ok {
		t.Fatalf("the owned twin of Llanowar Elves is not on the main list: %v", names(got))
	}
	if tw.Role != ramp || tw.Score != 0.9 {
		t.Errorf("Fyndhorn Elves has role %v and score %v, want the role and the score of Llanowar Elves", tw.Role, tw.Score)
	}
	// Elvish Mystic is a twin too, and it is on the list already. Balduvian
	// Bears is on the list, so no second copy joins.
	n := 0
	for _, c := range got {
		if c.Card.GetName() == "Balduvian Bears" || c.Card.GetName() == "Elvish Mystic" {
			n++
		}
	}
	if n != 2 {
		t.Errorf("a card already on the list joined it again: %v", names(got))
	}
	if _, added := addOwnedTwins(main, nil, scored); added != 0 {
		t.Error("a list with no card to buy took a twin")
	}
	// The staple floor puts an unowned card on the main list, and its
	// twin joins too. The live shortlist of F-232 held Llanowar Elves on
	// the main list and no upgrade of it.
	got, added = addOwnedTwins([]Candidate{scored[0], scored[1]}, nil, scored)
	if _, ok := find(got, "Fyndhorn Elves"); !ok || added != 1 {
		t.Errorf("the owned twin of a card to buy on the main list did not join: %v", names(got))
	}
}
