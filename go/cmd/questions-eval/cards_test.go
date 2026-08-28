package main

import (
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/tune"
)

// The snapshot loads once per package. Every snapshot test shares it.
var (
	checkerOnce sync.Once
	checkerIdx  cardChecker
	checkerNote string
)

// snapshotChecker returns the card checker over CARDS_SNAPSHOT_DIR. The
// tests skip without the variable, the way the candidates tests do.
func snapshotChecker(t *testing.T) cardChecker {
	t.Helper()
	if os.Getenv("CARDS_SNAPSHOT_DIR") == "" {
		t.Skip("set CARDS_SNAPSHOT_DIR to run the snapshot tests")
	}
	checkerOnce.Do(func() { checkerIdx, checkerNote = newCardChecker() })
	if checkerNote != "" {
		t.Fatalf("snapshot: %s", checkerNote)
	}
	return checkerIdx
}

// TestDeniesExistence is D-149. The two false claims of eval run 14 are
// the first two cases. The rest are reasons that name a card and claim
// nothing about whether it exists, and every one must pass through.
func TestDeniesExistence(t *testing.T) {
	deny := []string{
		`Commander selection is needed, but "Ran and Shaw" is not a real Magic card or commander option.`,
		`"Quina, Qu Gourmet" is not a valid Magic card option; offer only actual legal commander candidates.`,
		`The agent offered "Made Up Legend", which does not exist.`,
	}
	for _, r := range deny {
		if !deniesExistence(r) {
			t.Errorf("a denial was not read: %q", r)
		}
	}
	keep := []string{
		`It offers Birgi and Emrakul, whose color identities do not match the requested blue-red deck.`,
		`The user had already said "Grist is my commander", so asking again is redundant.`,
		`This repeats the same commander-selection question already asked on turn 2.`,
	}
	for _, r := range keep {
		if deniesExistence(r) {
			t.Errorf("a reason that claims nothing about existence was read as a denial: %q", r)
		}
	}
}

// TestQuotedNames reads the card name out of a reason. The eval writes
// straight quotes and curly quotes, and run 14 used both.
func TestQuotedNames(t *testing.T) {
	cases := []struct {
		reason string
		want   string
	}{
		{`"Ran and Shaw" is not a real Magic card.`, "Ran and Shaw"},
		{"“Quina, Qu Gourmet” is not a valid Magic card option.", "Quina, Qu Gourmet"},
		{`The card 'Lightning Bolt' can not lead a deck.`, "Lightning Bolt"},
	}
	for _, c := range cases {
		got := quotedNames(c.reason)
		if len(got) == 0 || got[0] != c.want {
			t.Errorf("quotedNames(%q) = %v, want %q first", c.reason, got, c.want)
		}
	}
}

// TestCorrectKeepsAnotherFault is D-149. A verdict that held a second
// fault stays bad. Run 14 refused probe 46 turn 4 as "duplicate,
// inaccurate", and the duplicate is real even though the card is too.
func TestCorrectKeepsAnotherFault(t *testing.T) {
	c := cardChecker{}
	// With no index the check is off, and every verdict passes through.
	v := tune.Verdict{Warranted: "no", Faults: []string{"duplicate", "inaccurate"},
		Reason: `"Ran and Shaw" is not a real Magic card.`}
	got, _, ok := c.correct(v)
	if ok {
		t.Error("the check ran with no snapshot loaded")
	}
	if len(got.Faults) != 2 {
		t.Errorf("faults = %v, want both kept", got.Faults)
	}
}

// TestCardCheckReportsWhenItDidNotRun keeps a silent skip out of the
// document. A missing snapshot must read as a missing check, and never as
// a clean run (D-149).
func TestCardCheckReportsWhenItDidNotRun(t *testing.T) {
	var b strings.Builder
	writeCardCheck(&b, "no card snapshot: the eval card check did not run", nil)
	if !strings.Contains(b.String(), "did not run") {
		t.Errorf("the report hid a skipped check: %q", b.String())
	}
	b.Reset()
	writeCardCheck(&b, "", nil)
	if !strings.Contains(b.String(), "No refusal rested on a card name") {
		t.Errorf("a clean run was not reported: %q", b.String())
	}
	b.Reset()
	writeCardCheck(&b, "", map[string]string{"46. a format we do not support: Ran and Shaw": "commander_pick"})
	if !strings.Contains(b.String(), "Ran and Shaw") || !strings.Contains(b.String(), "commander_pick") {
		t.Errorf("the correction was not reported: %q", b.String())
	}
}

// TestSnapshotRefutesRunFourteen is the live half of D-149. It runs only
// where the card snapshot is present, and it proves the two claims of
// eval run 14 against it rather than against memory.
func TestSnapshotRefutesRunFourteen(t *testing.T) {
	c := snapshotChecker(t)
	cases := []struct{ reason, card string }{
		{`Commander selection is needed because the user said only "A dragon deck, red", but "Ran and Shaw" is not a real Magic card or commander option.`, "Ran and Shaw"},
		{`The user had not chosen a commander, but "Quina, Qu Gourmet" is not a valid Magic card option.`, "Quina, Qu Gourmet"},
	}
	for _, tc := range cases {
		name, ok := c.deniedRealCard(tc.reason)
		if !ok || name != tc.card {
			t.Errorf("the snapshot did not refute the claim about %s: got %q, %v", tc.card, name, ok)
		}
	}
	// A verdict whose only fault was the false claim is no longer bad.
	v := tune.Verdict{Warranted: "no", Faults: []string{"inaccurate"}, CatalogAction: "reword",
		Reason: `"Ran and Shaw" is not a real Magic card or commander option.`}
	got, _, ok := c.correct(v)
	if !ok {
		t.Fatal("the check did not fire on a refuted claim")
	}
	if got.Bad() {
		t.Errorf("the verdict stayed bad after its only fault was refuted: %+v", got)
	}
	// A real card fact the eval got right must survive untouched.
	keep := tune.Verdict{Warranted: "no", Faults: []string{"duplicate"},
		Reason: "This repeats the same commander-selection question already asked on turn 2."}
	if _, _, fired := c.correct(keep); fired {
		t.Error("the check fired on a reason that claims nothing about a card")
	}
}

// TestFactsNameTheCardsAQuestionOffers is D-152. Every card the eval got
// wrong came from a crossover set and reached the user through the pick
// row's options. The facts block answers all four from the snapshot.
func TestFactsNameTheCardsAQuestionOffers(t *testing.T) {
	c := snapshotChecker(t)
	conv := tune.Conversation{
		Name: "the four cards the eval got wrong",
		Questions: []tune.Question{{
			Turn: 1, Row: "commander_pick",
			Options: []string{
				"Cloak and Dagger, Entwined", "Shadow the Hedgehog",
				"Vincent, Vengeful Atoner", "Quina, Qu Gourmet",
				"None, name three more", "A Card Nobody Holds",
			},
		}},
	}
	got := c.facts(conv)
	by := map[string]cardFact{}
	for _, f := range got {
		by[f.Name] = f
	}
	// Every real card is present, and every one is a legal commander.
	for _, name := range []string{
		"Cloak and Dagger, Entwined", "Shadow the Hedgehog",
		"Vincent, Vengeful Atoner", "Quina, Qu Gourmet",
	} {
		f, ok := by[name]
		if !ok {
			t.Errorf("%q is real and the facts block left it out", name)
			continue
		}
		if !f.CommanderList {
			t.Errorf("%q is a legal commander and the facts say otherwise", name)
		}
		if len(f.ColorIdentity) == 0 {
			t.Errorf("%q has a color identity and the facts carry none", name)
		}
		if f.TypeLine == "" {
			t.Errorf("%q carries no type line", name)
		}
	}
	// An option that names no card is left out, and nothing is claimed
	// about it.
	for _, skip := range []string{"None, name three more", "A Card Nobody Holds"} {
		if _, ok := by[skip]; ok {
			t.Errorf("%q is not a card and it reached the facts block", skip)
		}
	}
	// The color identity is what refutes an applicability claim. Cloak and
	// Dagger, Entwined is white and black, and eval run 16 called it
	// inapplicable to a white-black lifegain deck.
	if id := by["Cloak and Dagger, Entwined"].ColorIdentity; len(id) != 2 {
		t.Errorf("Cloak and Dagger, Entwined color identity = %v, want two colors", id)
	}
}
