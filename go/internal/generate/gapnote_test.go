package generate

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
)

func changer(id, name string) *mtgv1.Card {
	return &mtgv1.Card{OracleId: id, Name: name, CardTypes: []string{"Enchantment"}, ManaValue: 3, GameChanger: true}
}

func bracketPower(n int32) *mtgv1.PowerLevel {
	return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: n}}
}

// gapFixture is a bracket 4 deck that holds one Game Changer. Its pool
// holds two more, and its reserve holds two outside the pool.
func gapFixture(t *testing.T) (*Builder, Request, *mtgv1.Deck) {
	t.Helper()
	b, _, _ := testBuilder(t)
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	prof, err := profile.New(cfg, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b.profiler = prof
	held, a, bc, c, d := changer("o-held", "Held Changer"), changer("o-a", "Changer A"), changer("o-b", "Changer B"),
		changer("o-c", "Changer C"), changer("o-d", "Changer D")
	list := &candidates.List{
		Candidates: []candidates.Candidate{{Card: held, Rate: 0.9}, {Card: a, Rate: 0.3, Owned: 1}, {Card: bc, Rate: 0.5}},
		// Changer B is on the list as well, so the reserve does not name it
		// a second time.
		Reserve: []candidates.Candidate{{Card: c, Rate: 0.7, Owned: 1}, {Card: bc, Rate: 0.5}, {Card: d, Rate: 0.2}},
	}
	req := Request{
		Format:       mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Power:        bracketPower(4),
		Pool:         FromListOwned(list, nil, map[string]int32{"o-a": 1}, true),
		OracleCounts: map[string]int32{"o-a": 1, "o-c": 1},
	}
	deck := &mtgv1.Deck{
		Cards:   []*mtgv1.DeckCard{{OracleId: "o-held", Name: "Held Changer", Count: 1}},
		Profile: &mtgv1.DeckProfile{Bracket: 4, Features: []*mtgv1.ProfileFeature{{Key: profile.KeyGameChanger, Value: 1}}},
	}
	return b, req, deck
}

// TestGapNoteNamesPoolCardsThenReserveCards is D-709. A bracket 4 deck with
// one Game Changer is 3 under the floor of 4. The note names the pool cards
// the deck lacks by rate, then the reserve by rate, and with a collection
// each card reads owned or to buy. A card a revision removes leaves the
// reserve, and with no collection no card carries an ownership word.
func TestGapNoteNamesPoolCardsThenReserveCards(t *testing.T) {
	b, req, deck := gapFixture(t)
	want := "Bracket 4 wants 4 or more Game Changers, and the deck holds 1. " +
		"To close the gap, add Changer B (to buy), Changer A (owned), and Changer C (owned)."
	if got := b.gapNote(req, deck); got != want {
		t.Errorf("note %q\nwant %q", got, want)
	}
	req.Pool = req.Pool.Filter(func(c *mtgv1.Card) bool { return c.GetOracleId() != "o-c" })
	if got := b.gapNote(req, deck); !strings.HasSuffix(got, "add Changer B (to buy), Changer A (owned), and Changer D (to buy).") {
		t.Errorf("after the filter, note %q", got)
	}
	req.OracleCounts = nil
	if got := b.gapNote(req, deck); !strings.HasSuffix(got, "add Changer B, Changer A, and Changer D.") {
		t.Errorf("with no collection, note %q", got)
	}
}

// TestGapNoteStaysSilentWithoutAMiss is D-709: a deck at its floor, and a
// bracket with no floor, read no note.
func TestGapNoteStaysSilentWithoutAMiss(t *testing.T) {
	b, req, deck := gapFixture(t)
	deck.Profile.Features[0].Value = 4
	if got := b.gapNote(req, deck); got != "" {
		t.Errorf("a deck at its floor reads %q", got)
	}
	deck.Profile.Features[0].Value = 1
	req.Power = bracketPower(3)
	if got := b.gapNote(req, deck); got != "" {
		t.Errorf("a bracket 3 deck reads %q", got)
	}
}

// TestShortlistMarksThePowerCards is D-704: at bracket 4 a Game Changer
// line carries its mark, so the model counts what the check counts. A
// bracket 3 list holds no floor, and a revision reads no deck shape, so
// neither marks a card.
func TestShortlistMarksThePowerCards(t *testing.T) {
	b, req, _ := gapFixture(t)
	if got := b.shortlist(req); !strings.Contains(got, "- Changer A | Game Changer | owned 1") {
		t.Errorf("bracket 4 shortlist:\n%s", got)
	}
	req.Power = bracketPower(3)
	if got := b.shortlist(req); strings.Contains(got, "| Game Changer") {
		t.Errorf("bracket 3 shortlist marks a card:\n%s", got)
	}
	req.Power = bracketPower(4)
	req.Revision = &Revision{}
	if got := b.shortlist(req); strings.Contains(got, "| Game Changer") {
		t.Errorf("a revision shortlist marks a card:\n%s", got)
	}
}
