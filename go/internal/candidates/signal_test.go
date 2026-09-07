package candidates

import (
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// TestCommanderPoolSignal: a bracket 4 or 5 request ranks its
// commanders on the cEDH signal first, and a lower bracket keeps the
// popularity order (PR-14B, OQ-48).
func TestCommanderPoolSignal(t *testing.T) {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	legend := func(id, name string, rank int32) *mtgv1.Card {
		return &mtgv1.Card{
			OracleId: id, Name: name, CanBeCommander: true, EdhrecRank: rank, Legalities: legal,
			CardTypes: []string{"Creature"}, Supertypes: []string{"Legendary"},
			Colors: []mtgv1.Color{mtgv1.Color_COLOR_G}, ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_G},
			DefaultPrinting: &mtgv1.Printing{ScryfallId: "p-" + id, SetCode: "cmr"},
			SetCodes:        []string{"cmr"},
		}
	}
	idx := cards.NewIndex([]*mtgv1.Card{
		legend("a", "Popular Legend", 1),
		legend("b", "Middle Legend", 2),
		legend("c", "Tournament Legend", 3),
	}, nil, nil, time.Now())
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	signal := func(ids ...string) float64 {
		if len(ids) == 1 && ids[0] == "c" {
			return 0.8
		}
		return 0
	}
	names := func(req Request) []string {
		pool, err := b.CommanderPool(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		var out []string
		for _, c := range pool {
			out = append(out, c.Card.GetName())
		}
		return out
	}
	low := names(Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Bracket: 3, CommanderSignal: signal})
	if len(low) != 3 || low[0] != "Popular Legend" {
		t.Errorf("bracket 3 = %v, want the popularity order", low)
	}
	high := names(Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Bracket: 5, CommanderSignal: signal})
	if len(high) != 3 || high[0] != "Tournament Legend" || high[1] != "Popular Legend" {
		t.Errorf("bracket 5 = %v, want the signal first and popularity after", high)
	}
	none := names(Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Bracket: 5})
	if len(none) != 3 || none[0] != "Popular Legend" {
		t.Errorf("bracket 5 with no signal = %v, want the popularity order", none)
	}
}

func TestFoldNameQuotes(t *testing.T) {
	if FoldName("Commander\u2019s Sphere") != "commander's sphere" {
		t.Fatalf("fold = %q", FoldName("Commander\u2019s Sphere"))
	}
}
