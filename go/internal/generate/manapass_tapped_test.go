package generate

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The F-165 tests of the fixing fill (D-835). Each land text is the
// Oracle text of the card.

// tappedCards are the swap cards, two lands that enter tapped and fix no
// color, and one colorless land that enters untapped.
type tappedCards struct {
	swapCards
	lair, crypt, tower *mtgv1.Card
}

func newTappedCards() tappedCards {
	colorless := func(oid, name, text string) *mtgv1.Card {
		return manaCard(oid, name, 0, []string{"Land"}, text, mtgv1.Color_COLOR_C)
	}
	return tappedCards{
		swapCards: newSwapCards(),
		lair:      colorless("o-lair", "Tapped Lair", "This land enters tapped.\n{T}: Add {C}."),
		crypt:     colorless("o-crypt", "Tapped Crypt", "This land enters tapped.\n{T}: Add {C}."),
		tower:     colorless("o-tower", "Reliquary Tower", "You have no maximum hand size.\n{T}: Add {C}."),
	}
}

// deck is the given basic lands, the three colorless lands, and spells of
// two mana to 99 cards, at bracket 3.
func (s tappedCards) deck(islands int32) *mtgv1.Deck {
	entry := func(c *mtgv1.Card, n int32, role mtgv1.CardRole) *mtgv1.DeckCard {
		return &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: n, Role: role}
	}
	land := mtgv1.CardRole_CARD_ROLE_LAND
	spells := 99 - islands - 4
	return &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:  &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}},
		Cards: []*mtgv1.DeckCard{
			entry(s.island, islands, land), entry(s.swamp, 1, land),
			entry(s.lair, 1, land), entry(s.crypt, 1, land), entry(s.tower, 1, land),
			entry(s.blue, spells/2, mtgv1.CardRole_CARD_ROLE_THREAT),
			entry(s.black, spells-spells/2, mtgv1.CardRole_CARD_ROLE_THREAT),
		},
	}
}

// fill runs the fixing fill over the deck with a pool of the given lands.
func (s tappedCards) fill(t *testing.T, deck *mtgv1.Deck, pool ...*mtgv1.Card) int {
	t.Helper()
	src := source{}
	for _, c := range append(s.all(), s.lair, s.crypt, s.tower) {
		src[c.GetOracleId()] = c
	}
	b := manaBuilder(t, src)
	pool = append(pool, s.island, s.swamp)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(),
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, Pool: NewPool(pool, ownedAll(pool))}
	return b.fillFixing(req, deck)
}

// TestTheFillTradesATappedLandWhenNoBasicIsSpare is F-165 and D-835. A
// deck under its fixing floor with one basic land of each color trades
// each land that enters tapped and fixes no color. It keeps the basic
// lands and the colorless land that enters untapped, and the entry names
// the trade.
func TestTheFillTradesATappedLandWhenNoBasicIsSpare(t *testing.T) {
	s := newTappedCards()
	deck := s.deck(1)
	if n := s.fill(t, deck, s.grave, s.isle); n != 2 {
		t.Fatalf("the fill kept %d swaps, want 2", n)
	}
	for _, c := range []*mtgv1.Card{s.lair, s.crypt} {
		if countOf(deck, c.GetOracleId()) != 0 {
			t.Errorf("the fill kept %s, a land that enters tapped and fixes no color", c.GetName())
		}
	}
	for _, c := range []*mtgv1.Card{s.island, s.swamp, s.tower} {
		if countOf(deck, c.GetOracleId()) != 1 {
			t.Errorf("the fill traded %s", c.GetName())
		}
	}
	for _, dc := range deck.GetCards() {
		if dc.GetOracleId() == s.grave.GetOracleId() && dc.GetReason() != tapReason {
			t.Errorf("Watery Grave reads reason %q, want the reason of a tapped land", dc.GetReason())
		}
	}
	if countOf(deck, s.grave.GetOracleId()) != 1 || countOf(deck, s.isle.GetOracleId()) != 1 {
		t.Error("the fill did not take both fixing lands of the pool")
	}
	if n := deckCount(deck); n != 99 {
		t.Errorf("the deck holds %d cards after the fill, want 99", n)
	}
}

// TestTheFillTradesASpareBasicBeforeATappedLand: a spare basic land goes
// first, so a pool of one fixing land leaves each tapped land in place.
func TestTheFillTradesASpareBasicBeforeATappedLand(t *testing.T) {
	s := newTappedCards()
	deck := s.deck(3)
	if n := s.fill(t, deck, s.grave); n != 1 {
		t.Fatalf("the fill kept %d swaps, want 1", n)
	}
	if got := countOf(deck, s.island.GetOracleId()); got != 2 {
		t.Errorf("the deck holds %d Islands, want 2", got)
	}
	for _, c := range []*mtgv1.Card{s.lair, s.crypt} {
		if countOf(deck, c.GetOracleId()) != 1 {
			t.Errorf("the fill traded %s while a basic land was spare", c.GetName())
		}
	}
	for _, dc := range deck.GetCards() {
		if dc.GetOracleId() == s.grave.GetOracleId() && dc.GetReason() != fixReason {
			t.Errorf("Watery Grave reads reason %q, want the fixing reason", dc.GetReason())
		}
	}
}
