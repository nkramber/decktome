package generate

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// F-76: the deck carried commander_oracle_ids and no ownership fact for
// the commander, so every reader had to invent one. The web app wrote
// "not owned" and named the commander as a card to buy on every deck
// (D-604). The deck carries the fact now (D-608).

func TestCommanderCardsCarryTheOwnership(t *testing.T) {
	src := source{
		"o-karlov": {OracleId: "o-karlov", Name: "Karlov of the Ghost Council", PriceUsd: 3.5},
		"o-thran":  {OracleId: "o-thran", Name: "Thranduil, Elvenking", PriceUsd: 1.25},
	}
	got := commanderCards([]string{"o-karlov", "o-thran"}, src, map[string]int32{"o-karlov": 1})
	if len(got) != 2 {
		t.Fatalf("commanders = %+v, want two", got)
	}
	if !got[0].GetOwned() || got[0].GetOwnedCount() != 1 || got[0].GetPriceUsd() != 3.5 {
		t.Errorf("the owned commander = %+v", got[0])
	}
	if got[0].GetName() != "Karlov of the Ghost Council" || got[0].GetCount() != 1 {
		t.Errorf("the entry names one copy of the card: %+v", got[0])
	}
	if got[1].GetOwned() || got[1].GetOwnedCount() != 0 || got[1].GetPriceUsd() != 1.25 {
		t.Errorf("the commander the reader does not own = %+v", got[1])
	}
}

// A commander the card source does not hold gets no entry. Nothing
// recorded whether the reader owns it, and an invented answer is F-76.
func TestCommanderCardsSkipACardTheSourceDoesNotHold(t *testing.T) {
	if got := commanderCards([]string{"o-gone"}, source{}, nil); len(got) != 0 {
		t.Errorf("commanders = %+v, want nothing", got)
	}
	if got := commanderCards([]string{"o-karlov"}, nil, nil); len(got) != 0 {
		t.Errorf("commanders = %+v with no card source, want nothing", got)
	}
}

// The built deck carries the entry, so no reader of the deck invents one.
func TestBuiltDeckCarriesItsCommander(t *testing.T) {
	out := step(t, deckOut{Summary: "a lifegain deck", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 4, Role: "synergy", Reason: "gains life"},
	}})
	b, _, _ := testBuilder(t, out, out)
	req := testRequest()
	req.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	req.Commanders = []string{"o-karlov"}
	req.OracleCounts = map[string]int32{"o-solring": 1}
	got, err := b.Build(t.Context(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(got.Deck.GetCommanders()) != 1 {
		t.Fatalf("the deck carries %+v, want one commander entry", got.Deck.GetCommanders())
	}
	e := got.Deck.GetCommanders()[0]
	if e.GetOracleId() != "o-karlov" || e.GetOwned() {
		t.Errorf("the commander entry = %+v, want the unowned card", e)
	}
}
