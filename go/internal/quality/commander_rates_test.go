package quality

import (
	"fmt"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
)

// rated makes one resolved Commander list of a source for the commander
// ids, with one copy of each card.
func rated(source string, n int, commanders []string, held ...*mtgv1.Card) *Resolved {
	r := &Resolved{
		List: &meta.List{Source: source, ID: fmt.Sprintf("%s/%d", source, n), Format: meta.FormatCommander, Tier: meta.TierGood},
		Deck: &mtgv1.Deck{Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}, CommanderOracleIds: commanders},
	}
	for _, c := range held {
		r.Deck.Cards = append(r.Deck.Cards, &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: 1})
	}
	return r
}

// TestCommanderRatesReadTopDeckLists is D-839 (F-166). A commander with
// MinCommanderLists TopDeck lists gets the share of its lists that hold
// each card, lands included. A basic land, the commander itself, a card
// under CommanderRateFloor, and a list of another source never count. A
// commander with fewer lists gets no row, and a pair keys on both ids.
func TestCommanderRatesReadTopDeckLists(t *testing.T) {
	combo, filler, rare := spell("Combo Piece", 2, W), spell("Filler", 3, W), spell("Rare Card", 3, W)
	fetch, plains := land("Fetch Land", false, W), land("Plains", true, W)
	cmd := spell("The Commander", 3, W)
	idx := cards.NewIndex([]*mtgv1.Card{combo, filler, rare, fetch, plains, cmd}, nil, nil, time.Date(2026, 9, 23, 0, 0, 0, 0, time.UTC))
	one := []string{cmd.GetOracleId()}

	var reals []*Resolved
	for i := 0; i < 20; i++ {
		held := []*mtgv1.Card{combo, plains, cmd}
		if i < 10 {
			held = append(held, fetch)
		}
		if i < 2 {
			held = append(held, filler)
		}
		if i < 1 {
			held = append(held, rare)
		}
		reals = append(reals, rated(meta.SourceTopdeck, i, one, held...))
	}
	// Lists of another source hold the filler, and never count.
	for i := 0; i < 30; i++ {
		reals = append(reals, rated(meta.SourceMTGTop8, i, one, filler))
	}
	// A pair with MinCommanderLists lists, and a commander with one fewer.
	pair := []string{"oid-b", "oid-a"}
	for i := 0; i < MinCommanderLists; i++ {
		reals = append(reals, rated(meta.SourceTopdeck, 100+i, pair, combo))
	}
	for i := 0; i < MinCommanderLists-1; i++ {
		reals = append(reals, rated(meta.SourceTopdeck, 200+i, []string{"oid-short"}, combo))
	}

	rates := CommanderRates(reals, idx)
	got, ok := rates[CommanderKey(one...)]
	if !ok {
		t.Fatalf("no row for the commander: %v", rates)
	}
	want := map[string]float64{combo.GetOracleId(): 1, fetch.GetOracleId(): 0.5, filler.GetOracleId(): 0.1}
	if got.Lists != 20 || len(got.Cards) != len(want) {
		t.Fatalf("row %+v, want 20 lists and the cards %v", got, want)
	}
	for id, v := range want {
		if got.Cards[id] != v {
			t.Errorf("rate of %s = %v, want %v", id, got.Cards[id], v)
		}
	}
	if r, ok := rates[CommanderKey("oid-a", "oid-b")]; !ok || r.Lists != MinCommanderLists {
		t.Errorf("pair row = %+v, %v; want %d lists under the ordered key", r, ok, MinCommanderLists)
	}
	if _, ok := rates["oid-short"]; ok {
		t.Errorf("a commander under %d lists has a row", MinCommanderLists)
	}
}

// TestTopdeckListsReadCommanderAlone is D-839: the fit passes the TopDeck
// lists of a Commander fit to the rates, before the tier cap, and none in
// another format.
func TestTopdeckListsReadCommanderAlone(t *testing.T) {
	var reals []*Resolved
	for i := 0; i < MaxTierLists+5; i++ {
		reals = append(reals, rated(meta.SourceTopdeck, i, []string{"oid-x"}))
	}
	reals = append(reals, rated(meta.SourceMTGTop8, 0, []string{"oid-x"}))
	if got := len(topdeckLists(reals, mtgv1.FormatId_FORMAT_ID_COMMANDER)); got != MaxTierLists+5 {
		t.Errorf("topdeck lists = %d, want %d", got, MaxTierLists+5)
	}
	if got := topdeckLists(reals, mtgv1.FormatId_FORMAT_ID_MODERN); got != nil {
		t.Errorf("a Modern fit read %d TopDeck lists", len(got))
	}
}

// TestScorerCommanderRate is D-839: the scorer answers the row of a
// commander after a round trip of the model, and nil with no row or no
// model.
func TestScorerCommanderRate(t *testing.T) {
	m := &Model{Version: "v", Formats: map[string]*FormatModel{meta.FormatCommander: {
		Keys: []string{"k"}, Means: []float64{0}, Stds: []float64{1}, Weights: []float64{1},
		Tiers: []string{meta.TierBaseline, meta.TierGreat}, Thresholds: []float64{0},
		CommanderRates: map[string]CommanderRate{"a+b": {Lists: 12, Cards: map[string]float64{"card": 0.75}}},
	}}}
	data, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	back, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	s := NewScorer(back)
	rate := s.CommanderRate("b", "a")
	if rate == nil || rate("card") != 0.75 || rate("other") != 0 {
		t.Fatalf("pair rate after the round trip is wrong")
	}
	if s.CommanderRate("a") != nil {
		t.Error("one partner of a pair reads the row of the pair")
	}
	var none *Scorer
	if none.CommanderRate("a", "b") != nil {
		t.Error("a nil scorer answers a rate")
	}
}
