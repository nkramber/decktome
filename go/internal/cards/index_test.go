package cards

import (
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// TestDigitalDefaultPrintingIsSwappedForPaper is D-221. The oracle_cards
// file names one printing per card, and for Diamond Valley that printing
// is Masters Edition, an online set. Its paper printing is Arabian
// Nights, and the deck generator reported a paper problem the card does
// not have.
func TestDigitalDefaultPrintingIsSwappedForPaper(t *testing.T) {
	digital := &mtgv1.Card{
		OracleId: "o-dv", Name: "Diamond Valley",
		DefaultPrinting: &mtgv1.Printing{
			ScryfallId: "p-me1", SetCode: "me1", SetName: "Masters Edition",
			Artist: "Old Artist", Digital: true,
		},
	}
	// A card whose default is already paper must not move.
	paperOnly := &mtgv1.Card{
		OracleId: "o-bolt", Name: "Lightning Bolt",
		DefaultPrinting: &mtgv1.Printing{ScryfallId: "p-lea", SetCode: "lea", Digital: false},
	}
	idx := NewIndex([]*mtgv1.Card{digital, paperOnly}, []Printing{
		{ScryfallID: "p-me1", OracleID: "o-dv", SetCode: "me1", Digital: true, ReleasedAt: "2007-09-10"},
		{ScryfallID: "p-arn", OracleID: "o-dv", SetCode: "arn", SetName: "Arabian Nights",
			Artist: "Jesper Myrfors", Digital: false, ReleasedAt: "1993-12-17"},
		{ScryfallID: "p-plst", OracleID: "o-dv", SetCode: "plst", SetName: "The List",
			Artist: "Newer Artist", Digital: false, ReleasedAt: "2024-08-02", PriceUSD: 12.34},
		{ScryfallID: "p-lea", OracleID: "o-bolt", SetCode: "lea", Digital: false, ReleasedAt: "1993-08-05"},
	}, nil, time.Time{})

	got, ok := idx.ByOracleID("o-dv")
	if !ok {
		t.Fatal("the card left the index")
	}
	if got.GetDefaultPrinting().GetDigital() {
		t.Error("the default printing is still digital")
	}
	// The newest paper printing wins: a player buys that one.
	if got.GetDefaultPrinting().GetSetCode() != "plst" {
		t.Errorf("set = %q, want the newest paper printing", got.GetDefaultPrinting().GetSetCode())
	}
	// The artist must follow the printing, or the credit is wrong (D-6).
	if got.GetDefaultPrinting().GetArtist() != "Newer Artist" {
		t.Errorf("artist = %q, want the paper printing's artist", got.GetDefaultPrinting().GetArtist())
	}
	if bolt, _ := idx.ByOracleID("o-bolt"); bolt.GetDefaultPrinting().GetScryfallId() != "p-lea" {
		t.Error("a paper default was replaced")
	}
	if idx.PaperSwaps() != 1 {
		t.Errorf("swaps = %d, want 1", idx.PaperSwaps())
	}
	// The price follows the printing. A digital printing carries no USD
	// price, so the card read as free before D-231.
	if got.GetPriceUsd() != 12.34 {
		t.Errorf("price = %v, want the paper printing's price", got.GetPriceUsd())
	}
	if got.GetPriceAsOf() == "" {
		t.Error("the swapped price carries no date")
	}
}

// TestFullNameTiePrefersPlayableCard covers a full name two cards share.
// A playtest card from set cmb2 is legal nowhere and carries the name of
// a playable card. The playable card wins the name whatever the file order.
// Two cards legal somewhere keep the first-wins rule.
func TestFullNameTiePrefersPlayableCard(t *testing.T) {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	nowhere := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_NOT_LEGAL}
	playtest := &mtgv1.Card{OracleId: "o-playtest", Name: "Patient Turtle", Legalities: nowhere}
	playable := &mtgv1.Card{OracleId: "o-playable", Name: "Patient Turtle", Legalities: legal}
	first := &mtgv1.Card{OracleId: "o-first", Name: "Twin Name", Legalities: legal}
	second := &mtgv1.Card{OracleId: "o-second", Name: "Twin Name", Legalities: legal}
	tests := []struct {
		name  string
		cards []*mtgv1.Card
		look  string
		want  string
	}{
		{"unplayable first", []*mtgv1.Card{playtest, playable}, "Patient Turtle", "o-playable"},
		{"playable first", []*mtgv1.Card{playable, playtest}, "Patient Turtle", "o-playable"},
		{"two playable cards keep the first", []*mtgv1.Card{first, second}, "Twin Name", "o-first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := NewIndex(tt.cards, nil, nil, time.Time{})
			got, ok := idx.ByName(tt.look)
			if !ok {
				t.Fatalf("%q does not resolve", tt.look)
			}
			if got.OracleId != tt.want {
				t.Errorf("%q resolved to %s, want %s", tt.look, got.OracleId, tt.want)
			}
			if idx.Collisions().FullNames != 1 {
				t.Errorf("full name collisions = %d, want 1", idx.Collisions().FullNames)
			}
		})
	}
}

// TestIndexKeepsPrintings is D-299. A playable printing keeps its display
// fields and its price, and a dropped one answers false.
func TestIndexKeepsPrintings(t *testing.T) {
	card := &mtgv1.Card{OracleId: "o1", Name: "Forest", DefaultPrinting: &mtgv1.Printing{ScryfallId: "p-new"}}
	printings := []Printing{
		{ScryfallID: "p-old", OracleID: "o1", Name: "Forest", SetCode: "lea", CollectorNumber: "294", Artist: "Christopher Rush",
			ImageUris: &mtgv1.ImageUris{Normal: "https://x/old.jpg"}, PriceUSD: 40},
		{ScryfallID: "p-token", OracleID: "o-none", Name: "Elf", Layout: "token"},
	}
	idx := NewIndex([]*mtgv1.Card{card}, printings, nil, time.Now())
	p, ok := idx.Printing("p-old")
	if !ok || p.GetPriceUsd() != 40 || p.GetImageUris().GetNormal() != "https://x/old.jpg" || p.GetArtist() != "Christopher Rush" {
		t.Errorf("printing = %v, ok = %v", p, ok)
	}
	if _, ok := idx.Printing("p-token"); ok {
		t.Error("a token printing entered the index")
	}
	if _, ok := idx.Printing("nope"); ok {
		t.Error("an unknown id answered true")
	}
}
