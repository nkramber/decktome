package generate

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/profile"
)

// The balance phase of F-102. The bands can hold while one color sits far
// under its need, because the band reads the worst color against one
// floor. The pass trades basics toward the color that falls short.

// balanceDeck is a white-black Commander deck of 36 basics and 63 spells.
// The white spells cost {3}{W} and need 16 white sources, and the black
// spells cost {2}{B} and need 18 black sources, by the Karsten table.
func balanceDeck(t *testing.T, plainsCount, swampCount int32) (*Builder, *mtgv1.Deck, Request) {
	t.Helper()
	src := source{}
	commander := manaCard("o-cmdr", "Orzhov Captain", 4, []string{"Creature"}, "")
	commander.Supertypes = []string{"Legendary"}
	commander.ColorIdentity = []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}
	plains := manaCard("o-plains", "Plains", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	plains.Supertypes = []string{"Basic"}
	swamp := manaCard("o-swamp", "Swamp", 0, []string{"Land"}, "", mtgv1.Color_COLOR_B)
	swamp.Supertypes = []string{"Basic"}
	white := manaCard("o-white", "White Spell", 4, []string{"Creature"}, "")
	white.ManaCost, white.Colors = "{3}{W}", []mtgv1.Color{mtgv1.Color_COLOR_W}
	black := manaCard("o-black", "Black Spell", 3, []string{"Creature"}, "")
	black.ManaCost, black.Colors = "{2}{B}", []mtgv1.Color{mtgv1.Color_COLOR_B}
	for _, c := range []*mtgv1.Card{commander, plains, swamp, white, black} {
		src[c.GetOracleId()] = c
	}
	deck := &mtgv1.Deck{
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:              &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 4}},
		CommanderOracleIds: []string{"o-cmdr"},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-plains", Name: "Plains", Count: plainsCount, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-swamp", Name: "Swamp", Count: swampCount, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-white", Name: "White Spell", Count: 14, Role: mtgv1.CardRole_CARD_ROLE_THREAT},
			{OracleId: "o-black", Name: "Black Spell", Count: 49, Role: mtgv1.CardRole_CARD_ROLE_THREAT},
		},
	}
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Power:  deck.GetPower(),
		Pool:   NewPool([]*mtgv1.Card{plains, swamp, white, black}, nil),
	}
	return manaBuilder(t, src), deck, req
}

// worstRatio is the lowest sources-over-need ratio of the deck's colors.
func worstRatio(b *Builder, deck *mtgv1.Deck) float64 {
	worst := 0.0
	for i, r := range profile.ColorSources(deck, b.cards) {
		if i == 0 || r.Ratio() < worst {
			worst = r.Ratio()
		}
	}
	return worst
}

// TestTheBalancePhaseMovesBasicsToTheColorThatFallsShort is F-102 on a
// deck of 26 Plains and 10 Swamps. White needs 16 sources and black 18,
// so the pass trades Plains for Swamps until both colors meet their need.
// The land count and every spell stay.
func TestTheBalancePhaseMovesBasicsToTheColorThatFallsShort(t *testing.T) {
	b, deck, req := balanceDeck(t, 26, 10)
	before := worstRatio(b, deck)
	steps := b.balanceBasics(req, deck, b.basicsOf(deck), b.manaScore(deck))
	if steps == 0 {
		t.Fatal("the balance phase made no trade on a deck of 10 Swamps for a black need of 18")
	}
	if got := countOf(deck, "o-swamp"); got < 18 {
		t.Errorf("the deck holds %d Swamps after the phase, and black needs 18", got)
	}
	if got := countOf(deck, "o-plains"); got < 16 {
		t.Errorf("the deck holds %d Plains after the phase, and white needs 16", got)
	}
	if after := worstRatio(b, deck); after <= before {
		t.Errorf("the worst color read %.2f before the phase and %.2f after, want a rise", before, after)
	}
	if lands := countOf(deck, "o-plains") + countOf(deck, "o-swamp"); lands != 36 {
		t.Errorf("the deck holds %d basics after the phase, want the 36 it had", lands)
	}
	if countOf(deck, "o-white") != 14 || countOf(deck, "o-black") != 49 {
		t.Error("the balance phase moved a spell")
	}
}

// TestTheBalancePhaseKeepsASplitThatMeetsEveryNeed is the cap of F-102. A
// deck whose colors all meet their need keeps the split the model chose.
func TestTheBalancePhaseKeepsASplitThatMeetsEveryNeed(t *testing.T) {
	b, deck, req := balanceDeck(t, 18, 18)
	if steps := b.balanceBasics(req, deck, b.basicsOf(deck), b.manaScore(deck)); steps != 0 {
		t.Errorf("the balance phase made %d trades on a deck whose colors meet their need", steps)
	}
	if countOf(deck, "o-plains") != 18 || countOf(deck, "o-swamp") != 18 {
		t.Error("the balance phase moved a basic of a split that meets every need")
	}
}
