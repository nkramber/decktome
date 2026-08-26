package collections

import (
	"context"
	"errors"
	"math"
	"os"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

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
	// Duplicate rows merge into one entry, so count copies, not entries:
	// every input copy is in an entry or in a reported row.
	badLines := map[int32]bool{}
	for _, b := range badResolve {
		badLines[b.Line] = true
	}
	var wantCopies int32
	for _, r := range rows {
		if !badLines[int32(r.Line)] {
			wantCopies += int32(r.Quantity)
		}
	}
	if CardCount(entries) != wantCopies {
		t.Errorf("copies in entries = %d, want %d", CardCount(entries), wantCopies)
	}
	if len(entries)+len(badResolve) > len(rows) {
		t.Errorf("accounted = %d, more than %d rows", len(entries)+len(badResolve), len(rows))
	}
	t.Logf("real export vs test index: %d entries, %d unresolved rows", len(entries), len(badResolve))
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
	// Both lines name the same printing (M21 4 is the default
	// printing), so they merge into one entry of six copies.
	if len(entries) != 1 || len(badResolve) != 0 {
		t.Fatalf("resolve: entries %d bad %d", len(entries), len(badResolve))
	}
	if entries[0].Quantity != 6 {
		t.Errorf("qty = %d, want 6", entries[0].Quantity)
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

// smallIndex builds an index with one real card and one token printing
// that shares the card's name (the Pawpatch Recruit case, C-1).
func smallIndex() *cards.Index {
	card := &mtgv1.Card{
		OracleId: "50ab0194-0000-0000-0000-000000000000",
		Name:     "Pawpatch Recruit",
		DefaultPrinting: &mtgv1.Printing{
			ScryfallId: "aaaaaaaa-0000-0000-0000-000000000001", SetCode: "blb", SetName: "Bloomburrow", CollectorNumber: "187",
		},
	}
	printings := []cards.Printing{
		{ScryfallID: card.DefaultPrinting.ScryfallId, OracleID: card.OracleId, Name: card.Name, SetCode: "blb", CollectorNumber: "187", Layout: "normal"},
		// The token has its own Oracle id, absent from the card list.
		{ScryfallID: "7cdd8679-93a7-4e58-a6f3-b48897697e89", OracleID: "10c3fe2a-0000-0000-0000-000000000000", Name: "Pawpatch Recruit", SetCode: "tblb", CollectorNumber: "21", Layout: "token"},
	}
	return cards.NewIndex([]*mtgv1.Card{card}, printings, nil, time.Now())
}

func manaboxLine(name, set, num, foil, qty, id, cond, lang string) string {
	return `Main,binder,"` + name + `",` + set + `,Whatever,` + num + `,` + foil + `,common,` + qty + `,1,` + id + `,0.10,false,false,` + cond + `,` + lang + `,USD,2026-01-01T00:00:00Z`
}

func parseOne(t *testing.T, line string) ([]Row, []*mtgv1.UnresolvedRow) {
	t.Helper()
	rows, bad, err := ParseManaBoxCSV(strings.NewReader(manaboxHeader + "\n" + line + "\n"))
	if err != nil {
		t.Fatal(err)
	}
	return rows, bad
}

// TestTokenRowNotPlayable is the PR-4 gate regression: fixture line 186
// is a token. It must be reported, never resolved to the real card.
func TestTokenRowNotPlayable(t *testing.T) {
	idx := smallIndex()
	f, err := os.Open("testdata/manabox_collection.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	rows, _, err := ParseManaBoxCSV(f)
	if err != nil {
		t.Fatal(err)
	}
	var tokenRow *Row
	for i := range rows {
		if rows[i].ScryfallID == "7cdd8679-93a7-4e58-a6f3-b48897697e89" {
			tokenRow = &rows[i]
		}
	}
	if tokenRow == nil {
		t.Fatal("fixture has no Pawpatch Recruit token row")
	}
	if tokenRow.Line != 186 {
		t.Errorf("token row line = %d, want 186", tokenRow.Line)
	}
	entries, bad := Resolve([]Row{*tokenRow}, idx)
	if len(entries) != 0 || len(bad) != 1 || bad[0].Reason != mtgv1.UnresolvedReason_UNRESOLVED_REASON_NOT_PLAYABLE {
		t.Fatalf("entries %v bad %v, want one NOT_PLAYABLE", entries, bad)
	}
	if bad[0].Line != 186 {
		t.Errorf("report line = %d, want 186", bad[0].Line)
	}

	// The set/collector path reports the token too, with no id.
	rows2, _ := parseOne(t, manaboxLine("Pawpatch Recruit", "TBLB", "21", "normal", "1", "", "near_mint", "en"))
	_, bad2 := Resolve(rows2, idx)
	if len(bad2) != 1 || bad2[0].Reason != mtgv1.UnresolvedReason_UNRESOLVED_REASON_NOT_PLAYABLE {
		t.Errorf("set/collector token: %v", bad2)
	}
}

// TestNameOnlyMatchStoresDefaultPrinting covers C-5.
func TestNameOnlyMatchStoresDefaultPrinting(t *testing.T) {
	idx := smallIndex()
	rows, _ := parseOne(t, manaboxLine("Pawpatch Recruit", "ZZZ", "999", "normal", "2", "not-a-known-id", "near_mint", "en"))
	entries, bad := Resolve(rows, idx)
	if len(entries) != 1 || len(bad) != 0 {
		t.Fatalf("entries %v bad %v", entries, bad)
	}
	e := entries[0]
	if e.ScryfallId != "aaaaaaaa-0000-0000-0000-000000000001" || e.SetCode != "blb" || e.CollectorNumber != "187" || e.SetName != "Bloomburrow" {
		t.Errorf("name-only entry kept input printing: %v", e)
	}
	if e.Language != "en" {
		t.Errorf("language = %q", e.Language)
	}
}

func TestDuplicateRowsMerge(t *testing.T) {
	idx := smallIndex()
	id := "aaaaaaaa-0000-0000-0000-000000000001"
	rows, _, err := ParseManaBoxCSV(strings.NewReader(manaboxHeader + "\n" +
		manaboxLine("Pawpatch Recruit", "BLB", "187", "normal", "2", id, "near_mint", "en") + "\n" +
		manaboxLine("Pawpatch Recruit", "BLB", "187", "normal", "3", id, "near_mint", "en") + "\n" +
		manaboxLine("Pawpatch Recruit", "BLB", "187", "foil", "1", id, "near_mint", "en") + "\n"))
	if err != nil {
		t.Fatal(err)
	}
	entries, bad := Resolve(rows, idx)
	if len(bad) != 0 || len(entries) != 2 {
		t.Fatalf("entries %v bad %v, want 2 entries", entries, bad)
	}
	if entries[0].Quantity != 5 || entries[1].Quantity != 1 || CardCount(entries) != 6 {
		t.Errorf("merged quantities: %v", entries)
	}
}

// TestMergedQuantityStopsAtCap is M-3 of the 2026-08-26 review. Each row
// stays under maxQuantity, so the parser accepts it, and the merge sum
// wrapped int32 negative on a long enough file. The row that would push
// the entry over the cap is a BAD_ROW, so the row arithmetic holds.
func TestMergedQuantityStopsAtCap(t *testing.T) {
	idx := smallIndex()
	id := "aaaaaaaa-0000-0000-0000-000000000001"
	var sb strings.Builder
	sb.WriteString(manaboxHeader + "\n")
	const n = 5
	for i := 0; i < n; i++ {
		sb.WriteString(manaboxLine("Pawpatch Recruit", "BLB", "187", "normal", "9000", id, "near_mint", "en") + "\n")
	}
	rows, badParse, err := ParseManaBoxCSV(strings.NewReader(sb.String()))
	if err != nil || len(badParse) != 0 || len(rows) != n {
		t.Fatalf("parse: rows %d bad %v err %v", len(rows), badParse, err)
	}
	entries, bad := Resolve(rows, idx)
	if len(entries) != 1 || entries[0].Quantity != 9000 {
		t.Fatalf("entries = %v, want one entry at 9000", entries)
	}
	if len(bad) != n-1 {
		t.Fatalf("bad = %d rows, want %d", len(bad), n-1)
	}
	for _, b := range bad {
		if b.Reason != mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW {
			t.Errorf("reason = %v, want BAD_ROW", b.Reason)
		}
	}
	if got := len(rows) - len(bad); got != len(rows)-n+1 {
		t.Errorf("resolved rows = %d", got)
	}
}

// TestCountsSaturate covers the sums over stored entries, which the
// merge cap does not bound.
func TestCountsSaturate(t *testing.T) {
	entries := []*mtgv1.CollectionEntry{
		{OracleId: "a", Quantity: math.MaxInt32},
		{OracleId: "a", Quantity: 5},
		{OracleId: "b", Quantity: 1},
	}
	if got := CardCount(entries); got != math.MaxInt32 {
		t.Errorf("CardCount = %d, want the int32 maximum", got)
	}
	if got := OracleCounts(entries)["a"]; got != math.MaxInt32 {
		t.Errorf("OracleCounts[a] = %d, want the int32 maximum", got)
	}
}

func TestParseRejections(t *testing.T) {
	tests := []struct {
		name string
		line string
		want mtgv1.UnresolvedReason
	}{
		{"quantity zero", manaboxLine("X", "blb", "1", "normal", "0", "", "near_mint", "en"), mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW},
		{"quantity above cap", manaboxLine("X", "blb", "1", "normal", "10001", "", "near_mint", "en"), mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW},
		{"quantity overflow", manaboxLine("X", "blb", "1", "normal", "99999999999", "", "near_mint", "en"), mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW},
		{"unknown foil", manaboxLine("X", "blb", "1", "glossy", "1", "", "near_mint", "en"), mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE},
		{"unknown condition", manaboxLine("X", "blb", "1", "normal", "1", "", "damaged", "en"), mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows, bad := parseOne(t, tt.line)
			if len(rows) != 0 || len(bad) != 1 || bad[0].Reason != tt.want {
				t.Errorf("rows %v bad %v, want one %v", rows, bad, tt.want)
			}
		})
	}
	rows, bad := parseOne(t, manaboxLine("X", "blb", "1", "normal", "10000", "", "", "en"))
	if len(bad) != 0 || len(rows) != 1 || rows[0].Quantity != 10000 || rows[0].Condition != mtgv1.Condition_CONDITION_NEAR_MINT {
		t.Errorf("cap row: rows %v bad %v", rows, bad)
	}
}

func TestBOMHeader(t *testing.T) {
	rows, bad, err := ParseManaBoxCSV(strings.NewReader("\uFEFFScryfall ID,Quantity\nabc,2\n"))
	if err != nil {
		t.Fatal(err)
	}
	if len(bad) != 0 || len(rows) != 1 || rows[0].ScryfallID != "abc" || rows[0].Quantity != 2 {
		t.Errorf("rows %v bad %v", rows, bad)
	}
}

func TestLineNumbersUsePhysicalLines(t *testing.T) {
	// The first record's name spans two physical lines. The bad row
	// after it is on physical line 4.
	text := "Name,Set code,Quantity\n\"Two\nLines\",blb,1\nX,blb,-1\n"
	rows, bad, err := ParseManaBoxCSV(strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0].Line != 2 {
		t.Errorf("rows = %v", rows)
	}
	if len(bad) != 1 || bad[0].Line != 4 {
		t.Errorf("bad = %v, want line 4", bad)
	}
}

func TestTruncateRawOnRuneBoundary(t *testing.T) {
	// 199 ASCII bytes, then a 3-byte rune that straddles byte 200.
	raw := strings.Repeat("a", 199) + "€" + "tail"
	got := truncateRaw(raw)
	if !utf8.ValidString(got) {
		t.Fatalf("invalid UTF-8: %q", got)
	}
	if got != strings.Repeat("a", 199) {
		t.Errorf("got %q (len %d)", got, len(got))
	}
	if r := unresolved(1, raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW); !utf8.ValidString(r.Raw) {
		t.Errorf("unresolved raw invalid")
	}
	if got := truncateRaw("ok\xffbad"); !utf8.ValidString(got) {
		t.Errorf("invalid input not repaired: %q", got)
	}
}

func TestReasonCounts(t *testing.T) {
	bad := []*mtgv1.UnresolvedRow{
		{Reason: mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW},
		{Reason: mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW},
		{Reason: mtgv1.UnresolvedReason_UNRESOLVED_REASON_NOT_PLAYABLE},
	}
	got := ReasonCounts(bad)
	if got["UNRESOLVED_REASON_BAD_ROW"] != 2 || got["UNRESOLVED_REASON_NOT_PLAYABLE"] != 1 {
		t.Errorf("counts = %v", got)
	}
	if ReasonCounts(nil) != nil {
		t.Error("empty counts should be nil")
	}
}

func TestPutRefusesOversizedPayload(t *testing.T) {
	r := NewRepo(nil)
	// Random-like text does not compress. 1 MiB of it exceeds the guard.
	var entries []*mtgv1.CollectionEntry
	for i := 0; i < 8000; i++ {
		entries = append(entries, &mtgv1.CollectionEntry{OracleId: ContentHash([]byte{byte(i), byte(i >> 8)}) + ContentHash([]byte{byte(i >> 3)}), Name: ContentHash([]byte{byte(i * 7)}), Quantity: 1})
	}
	_, err := r.Put(context.Background(), "u", &mtgv1.Collection{Entries: entries})
	if !errors.Is(err, ErrTooLarge) {
		t.Fatalf("err = %v, want ErrTooLarge", err)
	}
}

func TestStoredToProtoImportedAt(t *testing.T) {
	at := time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC)
	col := storedToProto("id", storedCollection{Name: "n", Source: "IMPORT_SOURCE_MANABOX_CSV", ImportedAt: at, CardCount: 3})
	if col.ImportedAt == nil || !col.ImportedAt.AsTime().Equal(at) || col.Source != mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV {
		t.Errorf("col = %v", col)
	}
}
