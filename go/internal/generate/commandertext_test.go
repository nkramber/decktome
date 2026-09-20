package generate

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// grimaCard is the commander of the Gríma review, as the snapshot of
// 2026-09-04 holds it. Its text states a payoff shape that the name alone
// never states: the trigger fires one time for each hit (D-771).
func grimaCard() *mtgv1.Card {
	return &mtgv1.Card{
		OracleId:  "o-grima",
		Name:      "Gríma, Saruman's Footman",
		TypeLine:  "Legendary Creature — Human Advisor",
		ManaCost:  "{2}{U}{B}",
		Power:     "1",
		Toughness: "4",
		OracleText: "Gríma can't be blocked.\n" +
			"Whenever Gríma deals combat damage to a player, that player exiles cards from the top of their library.",
	}
}

func grimaSource() source {
	src := source{}
	for _, name := range testPool().Names() {
		c, _ := testPool().Card(name)
		src[c.GetOracleId()] = c
	}
	src["o-grima"] = grimaCard()
	return src
}

// TestCommanderTextReachesThePrompt is D-771. The prompt carried the name
// of the commander alone, so the model had to recall the card.
func TestCommanderTextReachesThePrompt(t *testing.T) {
	b, _ := testBuilderWith(t, grimaSource())
	req := testRequest()
	req.Format, req.Commanders = mtgv1.FormatId_FORMAT_ID_COMMANDER, []string{"o-grima"}
	in := b.input(req, nil, nil)
	for _, want := range []string{
		"The commander reads:",
		"\nGríma, Saruman's Footman | Legendary Creature — Human Advisor | {2}{U}{B} | 1/4\n",
		"Gríma can't be blocked.",
		"Whenever Gríma deals combat damage to a player",
	} {
		if !strings.Contains(in, want) {
			t.Errorf("the prompt misses %q", want)
		}
	}
}

// TestCommanderTextIsCommanderOnly extends D-233 to the card text. Only a
// Commander session has a command zone.
func TestCommanderTextIsCommanderOnly(t *testing.T) {
	b, _ := testBuilderWith(t, grimaSource())
	for _, format := range []mtgv1.FormatId{
		mtgv1.FormatId_FORMAT_ID_MODERN,
		mtgv1.FormatId_FORMAT_ID_STANDARD,
	} {
		req := testRequest()
		req.Format, req.Commanders = format, []string{"o-grima"}
		if in := b.input(req, nil, nil); strings.Contains(in, "The commander reads:") {
			t.Errorf("%s: the prompt reads the commander text", format)
		}
	}
}

// TestCommanderTextReadsBothCommanders covers a pair. Each commander
// states its own payoff shape (D-771).
func TestCommanderTextReadsBothCommanders(t *testing.T) {
	src := grimaSource()
	src["o-second"] = &mtgv1.Card{
		OracleId:   "o-second",
		Name:       "Second Commander",
		TypeLine:   "Legendary Creature — Human Advisor",
		ManaCost:   "{1}{B}",
		Power:      "2",
		Toughness:  "2",
		OracleText: "Partner",
	}
	b, _ := testBuilderWith(t, src)
	req := testRequest()
	req.Format, req.Commanders = mtgv1.FormatId_FORMAT_ID_COMMANDER, []string{"o-grima", "o-second"}
	in := b.input(req, nil, nil)
	if !strings.Contains(in, "\nGríma, Saruman's Footman |") {
		t.Error("the prompt misses the first commander")
	}
	if !strings.Contains(in, "\nSecond Commander | Legendary Creature — Human Advisor | {1}{B} | 2/2\n") {
		t.Error("the prompt misses the second commander")
	}
}

// TestCommanderTextWithNoOracleText covers a card the index holds with no
// text. The block names the card and adds no empty line (D-771).
func TestCommanderTextWithNoOracleText(t *testing.T) {
	b, _ := testBuilderWith(t, grimaSource())
	req := testRequest()
	req.Format, req.Commanders = mtgv1.FormatId_FORMAT_ID_COMMANDER, []string{"o-karlov"}
	in := b.input(req, nil, nil)
	if !strings.Contains(in, "The commander is Karlov of the Ghost Council.") {
		t.Fatal("the prompt misses the commander sentence")
	}
	if !strings.Contains(in, "\nKarlov of the Ghost Council\n") {
		t.Error("the prompt misses the commander line")
	}
}

// TestCommanderTextSkipsAnUnknownCard covers an id the card index does not
// hold. The prompt then carries no commander block at all (D-771).
func TestCommanderTextSkipsAnUnknownCard(t *testing.T) {
	b, _ := testBuilderWith(t, grimaSource())
	req := testRequest()
	req.Format, req.Commanders = mtgv1.FormatId_FORMAT_ID_COMMANDER, []string{"o-missing"}
	if in := b.input(req, nil, nil); strings.Contains(in, "The commander reads:") {
		t.Error("the prompt reads a commander the index does not hold")
	}
}
