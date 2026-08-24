package collections

import (
	"os"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

func fixtureIndex(t *testing.T) *cards.Index {
	t.Helper()
	cf, err := os.Open("../cards/testdata/cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cf.Close() }()
	cardList, err := cards.LoadCards(cf, "cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	printings := printingsForTest(t)
	return cards.NewIndex(cardList, printings, nil, time.Now())
}

func printingsForTest(t *testing.T) []cards.Printing {
	t.Helper()
	pf, err := os.Open("../cards/testdata/printings_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = pf.Close() }()
	printings, err := cards.LoadPrintings(pf, "printings_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	return printings
}

const manaboxHeader = "Binder Name,Binder Type,Name,Set code,Set name,Collector number,Foil,Rarity,Quantity,ManaBox ID,Scryfall ID,Purchase price,Misprint,Altered,Condition,Language,Purchase price currency,Added"

func TestManaBoxParseAndResolve(t *testing.T) {
	idx := fixtureIndex(t)
	prints := printingsForTest(t)
	if len(prints) < 3 {
		t.Fatal("need printings in fixture")
	}
	p0, p1 := prints[0], prints[1]
	csvText := manaboxHeader + "\n" +
		`Main,binder,"` + p0.Name + `",XXX,Whatever,999,foil,rare,3,1,` + p0.ScryfallID + `,0.10,false,false,near_mint,en,USD,2026-01-01T00:00:00Z` + "\n" +
		`Main,binder,"` + p1.Name + `",` + p1.SetCode + `,Whatever,` + p1.CollectorNumber + `,normal,common,2,2,not-a-real-id,0.10,false,false,played,en,USD,2026-01-01T00:00:00Z` + "\n" +
		`Main,binder,Braingeyser,XXX,Whatever,1,normal,rare,1,3,,0.10,false,false,near_mint,ja,USD,2026-01-01T00:00:00Z` + "\n" +
		`Main,binder,Definitely Not A Card,XXX,Whatever,1,normal,rare,1,4,,0.10,false,false,near_mint,en,USD,2026-01-01T00:00:00Z` + "\n" +
		`Main,binder,Whatever,XXX,Whatever,1,normal,rare,-2,5,,0.10,false,false,near_mint,en,USD,2026-01-01T00:00:00Z`

	rows, badParse, err := ParseManaBoxCSV(strings.NewReader(csvText))
	if err != nil {
		t.Fatal(err)
	}
	entries, badResolve := Resolve(rows, idx)
	bad := append(append([]*mtgv1.UnresolvedRow{}, badParse...), badResolve...)

	if len(entries) != 2 {
		t.Fatalf("entries = %d, want 2 (got bad: %v)", len(entries), bad)
	}
	if entries[0].Quantity != 3 || entries[0].Finish != mtgv1.Finish_FINISH_FOIL {
		t.Errorf("entry0 = %v", entries[0])
	}
	if entries[1].Condition != mtgv1.Condition_CONDITION_PLAYED {
		t.Errorf("entry1 condition = %v", entries[1].Condition)
	}
	wantReasons := map[mtgv1.UnresolvedReason]int{
		mtgv1.UnresolvedReason_UNRESOLVED_REASON_NON_ENGLISH:  1,
		mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_CARD: 1,
		mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW:      1,
	}
	got := map[mtgv1.UnresolvedReason]int{}
	for _, b := range bad {
		got[b.Reason]++
	}
	for reason, n := range wantReasons {
		if got[reason] != n {
			t.Errorf("reason %v = %d, want %d (all: %v)", reason, got[reason], n, bad)
		}
	}
	if len(entries)+len(bad) != 5 {
		t.Errorf("accounted rows = %d, want 5", len(entries)+len(bad))
	}
}

func TestHeaderVariants(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		wantErr bool
	}{
		{"full export header", manaboxHeader, false},
		{"minimal scryfall id", "Scryfall ID,Quantity", false},
		{"name and set code", "Name,Set code,Quantity", false},
		{"name and set name", "Name,Set name", false},
		{"unusable", "Foo,Bar", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := ParseManaBoxCSV(strings.NewReader(tt.header + "\n"))
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestRealExportNoSilentDrops is the PR-4 gate on the owner's real file:
// every row parses, and every row lands in entries or in the report.
// Resolution uses the small test index, so most rows report unknown-card
// here. The resolution RATE gate runs against the full snapshot (e2e).
func TestRealExportNoSilentDrops(t *testing.T) {
	f, err := os.Open("testdata/manabox_collection.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	rows, badParse, err := ParseManaBoxCSV(f)
	if err != nil {
		t.Fatal(err)
	}
	if len(badParse) != 0 {
		t.Errorf("real export: %d rows failed to parse, first: %v", len(badParse), badParse[0])
	}
	// The export has 2,548 data rows (header excluded, no trailing newline).
	if len(rows) != 2548 {
		t.Errorf("rows = %d, want 2548", len(rows))
	}
	idx := fixtureIndex(t)
	entries, badResolve := Resolve(rows, idx)
	if len(entries)+len(badResolve) != len(rows) {
		t.Errorf("accounted = %d, want %d", len(entries)+len(badResolve), len(rows))
	}
	t.Logf("real export vs test index: %d resolved, %d unresolved", len(entries), len(badResolve))
}

func TestArenaText(t *testing.T) {
	idx := fixtureIndex(t)
	text := "Deck\n4 Ajani's Pridemate (M21) 4\n2 Ajani's Pridemate\n\n// comment\nnot a line\n"
	rows, bad, err := ParseArenaText(strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fatalf("rows = %d, want 2 (bad %v)", len(rows), bad)
	}
	if len(bad) != 1 {
		t.Errorf("bad = %d, want 1", len(bad))
	}
	entries, badResolve := Resolve(rows, idx)
	if len(entries) != 2 || len(badResolve) != 0 {
		t.Fatalf("resolve: entries %d bad %d", len(entries), len(badResolve))
	}
	if entries[0].Quantity != 4 {
		t.Errorf("qty = %d", entries[0].Quantity)
	}
}

func TestHashAndCounts(t *testing.T) {
	h1 := ContentHash([]byte("abc"))
	h2 := ContentHash([]byte("abc"))
	h3 := ContentHash([]byte("abd"))
	if h1 != h2 || h1 == h3 || len(h1) != 64 {
		t.Errorf("hash: %q %q %q", h1, h2, h3)
	}
	entries := []*mtgv1.CollectionEntry{
		{OracleId: "a", Quantity: 2, Name: "B"},
		{OracleId: "a", Quantity: 1, Name: "A"},
		{OracleId: "b", Quantity: 4, Name: "C"},
	}
	counts := OracleCounts(entries)
	if counts["a"] != 3 || counts["b"] != 4 {
		t.Errorf("counts = %v", counts)
	}
	if CardCount(entries) != 7 {
		t.Errorf("card count = %d", CardCount(entries))
	}
	SortEntries(entries)
	if entries[0].Name != "A" {
		t.Errorf("sort: first = %s", entries[0].Name)
	}
}
