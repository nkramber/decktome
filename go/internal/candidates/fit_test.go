package candidates

import (
	"fmt"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

func fitCreature(id, subtype string) *mtgv1.Card {
	return &mtgv1.Card{
		OracleId: id, Name: "Creature " + id, CardTypes: []string{"Creature"}, Subtypes: []string{subtype},
		TypeLine: "Creature — " + subtype,
	}
}

// TestThemeScores is D-1091: a new card of the theme reaches FitFloor,
// and a card off the theme scores 0.
func TestThemeScores(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	dino, human := fitCreature("d", "Dinosaur"), fitCreature("h", "Human")
	idx := cards.NewIndex([]*mtgv1.Card{dino, human}, nil, nil, time.Now())
	got := b.ThemeScores("dinosaurs", idx, []*mtgv1.Card{dino, human})
	if got[0] < FitFloor || got[1] != 0 {
		t.Fatalf("scores = %v, want the Dinosaur at %v or more and the Human at 0", got, FitFloor)
	}
	if none := b.ThemeScores("the best possible deck", idx, []*mtgv1.Card{dino}); none[0] != 0 {
		t.Fatalf("an empty theme scored %v", none)
	}
}

// TestDeckTheme names the plan of a deck from its cards, when its chat
// is gone (D-1091), and names nothing for a deck with no plan.
func TestDeckTheme(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	var deck []*mtgv1.Card
	for i := range 12 {
		deck = append(deck, fitCreature(fmt.Sprintf("d%d", i), "Dinosaur"))
	}
	for i := range 4 {
		deck = append(deck, fitCreature(fmt.Sprintf("h%d", i), "Human"))
	}
	for i := range 20 {
		deck = append(deck, &mtgv1.Card{OracleId: fmt.Sprintf("l%d", i), CardTypes: []string{"Land"}})
	}
	idx := cards.NewIndex(deck, nil, nil, time.Now())
	theme := b.DeckTheme(idx, deck)
	if theme == "" {
		t.Fatal("DeckTheme named nothing for twelve Dinosaurs")
	}
	scores := b.ThemeScores(theme, idx, []*mtgv1.Card{fitCreature("x", "Dinosaur"), fitCreature("y", "Human")})
	if scores[0] < FitFloor || scores[0] <= scores[1] {
		t.Fatalf("theme %q scores a new Dinosaur %v and a Human %v", theme, scores[0], scores[1])
	}

	var mixed []*mtgv1.Card
	for i, st := range []string{"Human", "Elf", "Goblin", "Merfolk", "Zombie", "Cat", "Bird", "Snake"} {
		mixed = append(mixed, fitCreature(fmt.Sprintf("m%d", i), st))
	}
	mixed = append(mixed, &mtgv1.Card{OracleId: "art", CardTypes: []string{"Artifact"}},
		&mtgv1.Card{OracleId: "ench", CardTypes: []string{"Enchantment"}})
	if got := b.DeckTheme(cards.NewIndex(mixed, nil, nil, time.Now()), mixed); got != "" {
		t.Fatalf("DeckTheme = %q for a deck with no plan, want none", got)
	}
	if got := b.DeckTheme(idx, deck[16:]); got != "" {
		t.Fatalf("DeckTheme = %q for lands alone, want none", got)
	}
}
