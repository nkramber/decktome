package revise

import (
	"strconv"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

func deck(cards ...string) *mtgv1.Deck {
	d := &mtgv1.Deck{}
	for _, c := range cards {
		parts := strings.SplitN(c, " ", 2)
		n, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		d.Cards = append(d.Cards, &mtgv1.DeckCard{Name: parts[1], Count: int32(n)})
	}
	return d
}

func TestDiffAndNote(t *testing.T) {
	base := deck("24 Plains", "1 Angel of Serenity", "2 Lyra Dawnbringer", "1 Victory's Herald")
	revised := deck("22 Plains", "2 Lyra Dawnbringer", "2 Seraph Sanctuary", "1 Giada, Font of Hope")
	d := DiffDecks(base, revised)
	if got := strings.Join(d.Removed, "|"); got != "1 Angel of Serenity|1 Victory's Herald" {
		t.Errorf("removed = %q", got)
	}
	if got := strings.Join(d.Added, "|"); got != "1 Giada, Font of Hope|2 Seraph Sanctuary" {
		t.Errorf("added = %q", got)
	}
	if got := strings.Join(d.Changed, "|"); got != "Plains: 24 to 22" {
		t.Errorf("changed = %q", got)
	}
	b := &Brief{Changes: []string{"Remove every nonland card with mana value 6 or more"}, MaxManaValue: 5,
		Declined: []Decline{{Request: "Replace some lands with better options", Reason: "for a casual mono-white deck, all basic lands is fine"}}}
	note := Note(b, d)
	for _, want := range []string{
		"I removed 1 Angel of Serenity and 1 Victory's Herald.",
		"I added 1 Giada, Font of Hope and 2 Seraph Sanctuary.",
		"I changed the count of Plains: 24 to 22.",
		"I did not replace some lands with better options: for a casual mono-white deck, all basic lands is fine.",
	} {
		if !strings.Contains(note, want) {
			t.Errorf("note lacks %q:\n%s", want, note)
		}
	}
}

func TestNoteWhenNothingChanged(t *testing.T) {
	base := deck("24 Plains")
	b := &Brief{Changes: []string{"Add more removal"}}
	if got := Note(b, DiffDecks(base, base)); !strings.Contains(got, "The deck is the same as before") {
		t.Errorf("note = %q", got)
	}
}

func TestDeclineNoteIsNeverSilent(t *testing.T) {
	if got := DeclineNote(&Brief{}); got == "" || !strings.Contains(got, "no change") {
		t.Errorf("empty brief note = %q", got)
	}
	got := DeclineNote(&Brief{Declined: []Decline{{Request: "Make it faster", Reason: ""}}})
	if !strings.Contains(got, "I did not make it faster: no change would help here.") || !strings.Contains(got, "The deck stays as it was.") {
		t.Errorf("decline note = %q", got)
	}
}

func TestOnlyInDeckKeepsDeckSpelling(t *testing.T) {
	d := deck("1 Lyra Dawnbringer", "24 Plains")
	got := onlyInDeck([]string{"lyra dawnbringer", "Nonesuch", "Plains", "plains"}, d)
	if strings.Join(got, "|") != "Lyra Dawnbringer|Plains" {
		t.Errorf("got %v", got)
	}
}

func TestActs(t *testing.T) {
	if (&Brief{Question: "Which lands?"}).Acts() {
		t.Error("a question alone acts")
	}
	if (&Brief{Declined: []Decline{{Request: "x", Reason: "y"}}}).Acts() {
		t.Error("a decline alone acts")
	}
	if !(&Brief{MaxManaValue: 5}).Acts() {
		t.Error("a cap does not act")
	}
}

func TestInputNamesManaValues(t *testing.T) {
	d := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o1", Name: "Angel of Serenity", Count: 1}, {OracleId: "o2", Name: "Plains", Count: 24}}}
	src := fakeCards{"o1": {Name: "Angel of Serenity", TypeLine: "Creature — Angel", ManaValue: 7, CardTypes: []string{"Creature"}},
		"o2": {Name: "Plains", TypeLine: "Basic Land — Plains", CardTypes: []string{"Land"}}}
	got := input(Input{Message: "no 7 drops", Prior: "build angels", Deck: d, Format: "Modern", Power: "casual", Cards: src})
	for _, want := range []string{"Format: Modern. Power: casual.", "- 1 Angel of Serenity (Creature — Angel, mana value 7)", "- 24 Plains (Basic Land — Plains)", "## The user's earlier message\n\nbuild angels", "## The user's message\n\nno 7 drops"} {
		if !strings.Contains(got, want) {
			t.Errorf("input lacks %q:\n%s", want, got)
		}
	}
}

type fakeCards map[string]*mtgv1.Card

func (f fakeCards) ByOracleID(id string) (*mtgv1.Card, bool) { c, ok := f[id]; return c, ok }
