package cards

import (
	"os"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// TestLoadRulingsKeysByOracleID reads three rows of the Scryfall rulings
// file of 2026-09-03, two of one card, and answers them oldest first.
func TestLoadRulingsKeysByOracleID(t *testing.T) {
	lines := strings.Join([]string{
		`{"object":"ruling","oracle_id":"o-energy","source":"wotc","published_at":"2025-02-07","comment":"{E} is the energy symbol. It represents one energy counter."}`,
		`{"object":"ruling","oracle_id":"o-energy","source":"wotc","published_at":"2024-01-01","comment":"Energy counters are a kind of counter that a player may have."}`,
		`{"object":"ruling","oracle_id":"o-other","source":"scryfall","published_at":"2026-01-01","comment":"Another card."}`,
		`{"object":"ruling","oracle_id":"","source":"wotc","published_at":"2026-01-01","comment":"no id"}`,
	}, "\n")
	rulings, err := LoadRulings(strings.NewReader(lines), "rulings.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	if len(rulings) != 2 || len(rulings["o-energy"]) != 2 || len(rulings["o-other"]) != 1 {
		t.Fatalf("rulings = %+v", rulings)
	}
	if rulings["o-energy"][0].PublishedAt != "2024-01-01" || rulings["o-energy"][1].Source != "wotc" {
		t.Errorf("the rulings of a card read oldest first: %+v", rulings["o-energy"])
	}
	idx := NewIndex([]*mtgv1.Card{{OracleId: "o-energy", Name: "Energy Card"}}, nil, nil, time.Now(), WithRulings(rulings))
	got := idx.Rulings("o-energy")
	if len(got) != 2 || !idx.HasRulings() {
		t.Fatalf("index rulings = %+v", got)
	}
	got[0].Comment = "changed"
	if idx.Rulings("o-energy")[0].Comment == "changed" {
		t.Error("Rulings must answer a copy")
	}
	if idx.Rulings("o-none") != nil {
		t.Error("a card with no ruling answers nil")
	}
	bare := NewIndex(nil, nil, nil, time.Now())
	if bare.HasRulings() || bare.Rulings("o-energy") != nil {
		t.Error("a snapshot with no rulings file answers none")
	}
}

// TestPrintingsOfListsEveryPlayablePrinting: the card detail shows every
// printing with its price, digital ones marked (PR-20).
func TestPrintingsOfListsEveryPlayablePrinting(t *testing.T) {
	card := &mtgv1.Card{OracleId: "o-sol", Name: "Sol Ring", DefaultPrinting: &mtgv1.Printing{ScryfallId: "p-1"}}
	printings := []Printing{
		{ScryfallID: "p-1", OracleID: "o-sol", SetCode: "c21", CollectorNumber: "1", PriceUSD: 1.5},
		{ScryfallID: "p-2", OracleID: "o-sol", SetCode: "msc", CollectorNumber: "5", PriceUSD: 2},
		{ScryfallID: "p-3", OracleID: "o-sol", SetCode: "mtga", CollectorNumber: "9", Digital: true},
		{ScryfallID: "p-4", OracleID: "o-sol", SetCode: "tst", CollectorNumber: "1", Layout: "art_series"},
	}
	idx := NewIndex([]*mtgv1.Card{card}, printings, nil, time.Now())
	got := idx.PrintingsOf("o-sol")
	if len(got) != 3 {
		t.Fatalf("printings = %d, want the three playable ones", len(got))
	}
	if got[0].ScryfallId != "p-1" || got[1].PriceUsd != 2 || !got[2].Digital {
		t.Errorf("printings = %+v", got)
	}
	if idx.PrintingsOf("o-none") != nil {
		t.Error("an unknown card answers nil")
	}
}

// TestLoadRulingsRealFile reads the Scryfall rulings bulk file when
// RULINGS_FILE names one, for example the 5.4 MB file of 2026-09-03 with
// 78,949 rows. It proves the parser on the real shape and skips without
// the file.
func TestLoadRulingsRealFile(t *testing.T) {
	path := os.Getenv("RULINGS_FILE")
	if path == "" {
		t.Skip("set RULINGS_FILE to a rulings.jsonl.gz to read the real file")
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	rulings, err := LoadRulings(f, path)
	if err != nil {
		t.Fatal(err)
	}
	rows := 0
	for _, list := range rulings {
		rows += len(list)
	}
	if len(rulings) < 10000 || rows < 50000 {
		t.Fatalf("%d cards with %d rulings, want tens of thousands", len(rulings), rows)
	}
	t.Logf("%d cards with %d rulings", len(rulings), rows)
	for id, list := range rulings {
		for i := 1; i < len(list); i++ {
			if list[i].PublishedAt < list[i-1].PublishedAt {
				t.Fatalf("%s: rulings out of date order", id)
			}
		}
	}
}
