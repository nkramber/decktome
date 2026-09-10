package collections

import (
	"os"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The fixture is a real Moxfield export, written 2026-09-10. The
// documentation of the format disagrees with itself, and a column list
// from a help page is not the file the app receives (F-91). So every
// test below reads the real file or a row copied out of it.

const moxfieldFixture = "testdata/moxfield_collection.csv"

func moxfieldFile(t *testing.T) *os.File {
	t.Helper()
	f, err := os.Open(moxfieldFixture)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = f.Close() })
	return f
}

// TestTheWholeMoxfieldExportParses is the bar that matters. A parser
// that reads the header and then rejects the rows is worse than no
// parser: the reader uploads their collection and reads that every card
// of it failed.
func TestTheWholeMoxfieldExportParses(t *testing.T) {
	rows, bad, err := ParseMoxfieldCSV(moxfieldFile(t))
	if err != nil {
		t.Fatalf("the real export does not parse: %v", err)
	}
	if len(bad) != 0 {
		t.Errorf("%d rows of the real export were rejected, want none. First: %v", len(bad), bad[0])
	}
	// The file holds 3193 rows under its header, counted 2026-09-10.
	if len(rows) != 3193 {
		t.Errorf("%d rows parsed, want 3193", len(rows))
	}
	// Every row carries what the resolver needs: a name, a set code, and
	// a quantity. A row with none of those resolves by luck alone.
	for i, r := range rows {
		if r.Name == "" || r.SetCode == "" || r.Quantity < 1 {
			t.Fatalf("row %d is short: %+v", i, r)
		}
		if r.Line == 0 {
			t.Fatalf("row %d carries no line number, so a report can not name it", i)
		}
	}
}

// TestTheMoxfieldLanguageIsACode holds the fault that would have lost
// every row. Resolve refuses a row whose language is not "en" (D-23),
// and Moxfield writes "English".
func TestTheMoxfieldLanguageIsACode(t *testing.T) {
	rows, _, err := ParseMoxfieldCSV(moxfieldFile(t))
	if err != nil {
		t.Fatal(err)
	}
	english, other := 0, 0
	for _, r := range rows {
		switch r.Language {
		case "en":
			english++
		case "ja":
			other++
		default:
			t.Fatalf("row %d carries language %q, which Resolve does not read", r.Line, r.Language)
		}
	}
	// The export holds 3192 English cards and one Japanese, counted
	// 2026-09-10.
	if english != 3192 || other != 1 {
		t.Errorf("%d English and %d Japanese, want 3192 and 1", english, other)
	}
	// The whole point: a row that kept "English" would come back as
	// non-English from the resolver.
	if _, ok := moxfieldLanguageOf("English"); !ok {
		t.Fatal("English does not map")
	}
	if code, _ := moxfieldLanguageOf("English"); code != "en" {
		t.Errorf("English maps to %q, want en", code)
	}
}

// TestTheMoxfieldEditionIsASetCode holds the other ambiguity the real
// file settled. One guide describes a set name in this column, and that
// guide describes the import format and not the export.
func TestTheMoxfieldEditionIsASetCode(t *testing.T) {
	rows, _, err := ParseMoxfieldCSV(moxfieldFile(t))
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, r := range rows {
		seen[r.SetCode] = true
		if len(r.SetCode) > 5 || strings.Contains(r.SetCode, " ") {
			t.Fatalf("row %d holds set %q, which reads as a name and not a code", r.Line, r.SetCode)
		}
		if r.SetCode != strings.ToLower(r.SetCode) {
			t.Fatalf("row %d holds set %q, which is not lowercase", r.Line, r.SetCode)
		}
		// The parser fills no set name: the export carries none, and a
		// field nothing sets must not read as a value (T-19).
		if r.SetName != "" {
			t.Fatalf("row %d holds a set name %q, and the export writes none", r.Line, r.SetName)
		}
	}
	// The export holds 125 distinct sets, counted 2026-09-10.
	if len(seen) != 125 {
		t.Errorf("%d distinct sets, want 125", len(seen))
	}
}

// TestTheMoxfieldFinishAndConditionRead reads the two vocabularies the
// export exercises.
func TestTheMoxfieldFinishAndConditionRead(t *testing.T) {
	rows, _, err := ParseMoxfieldCSV(moxfieldFile(t))
	if err != nil {
		t.Fatal(err)
	}
	byFinish := map[mtgv1.Finish]int{}
	byCondition := map[mtgv1.Condition]int{}
	for _, r := range rows {
		byFinish[r.Finish]++
		byCondition[r.Condition]++
	}
	// Counted from the export on 2026-09-10: an empty Foil cell 2526
	// times, "foil" 663, and "etched" 4.
	for finish, want := range map[mtgv1.Finish]int{
		mtgv1.Finish_FINISH_NORMAL: 2526,
		mtgv1.Finish_FINISH_FOIL:   663,
		mtgv1.Finish_FINISH_ETCHED: 4,
	} {
		if byFinish[finish] != want {
			t.Errorf("%s = %d, want %d", finish, byFinish[finish], want)
		}
	}
	// "Near Mint" 3192 times and "Played" once.
	if byCondition[mtgv1.Condition_CONDITION_NEAR_MINT] != 3192 {
		t.Errorf("near mint = %d, want 3192", byCondition[mtgv1.Condition_CONDITION_NEAR_MINT])
	}
	if byCondition[mtgv1.Condition_CONDITION_PLAYED] != 1 {
		t.Errorf("played = %d, want 1", byCondition[mtgv1.Condition_CONDITION_PLAYED])
	}
}

// TestTheMoxfieldCollectorNumberIsAString reads the numbers the export
// actually holds. Four of them are not numbers at all, and a parser
// that read an integer would drop those rows.
func TestTheMoxfieldCollectorNumberIsAString(t *testing.T) {
	rows, _, err := ParseMoxfieldCSV(moxfieldFile(t))
	if err != nil {
		t.Fatal(err)
	}
	odd := map[string]bool{}
	for _, r := range rows {
		if r.Collector == "" {
			t.Fatalf("row %d carries no collector number", r.Line)
		}
		if strings.ContainsAny(r.Collector, "abcdefghijklmnopqrstuvwxyz-★") {
			odd[r.Collector] = true
		}
	}
	// "ml219sb", "2025-21", "2026-1", and "14★" are in the export.
	for _, want := range []string{"ml219sb", "2025-21", "2026-1", "14★"} {
		if !odd[want] {
			t.Errorf("the export holds collector number %q and the parser lost it", want)
		}
	}
}

// TestMoxfieldRejectsARowItCanNotRead: a bad count or an unknown value
// is reported, and it never defaults in silence (D-23).
func TestMoxfieldRejectsARowItCanNotRead(t *testing.T) {
	const header = `"Count","Name","Edition","Condition","Language","Foil","Collector Number"` + "\n"
	for name, tc := range map[string]struct {
		row    string
		reason mtgv1.UnresolvedReason
	}{
		"a count of zero":              {`"0","Sol Ring","c21","Near Mint","English","","263"`, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW},
		"a count that is not a number": {`"many","Sol Ring","c21","Near Mint","English","","263"`, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW},
		"an unknown condition":         {`"1","Sol Ring","c21","Chewed","English","","263"`, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE},
		"an unknown finish":            {`"1","Sol Ring","c21","Near Mint","English","rainbow","263"`, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE},
		"an unknown language":          {`"1","Sol Ring","c21","Near Mint","Klingon","","263"`, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE},
	} {
		t.Run(name, func(t *testing.T) {
			rows, bad, err := ParseMoxfieldCSV(strings.NewReader(header + tc.row + "\n"))
			if err != nil {
				t.Fatal(err)
			}
			if len(rows) != 0 {
				t.Fatalf("a bad row parsed: %+v", rows)
			}
			if len(bad) != 1 {
				t.Fatalf("%d rows reported, want 1", len(bad))
			}
			if bad[0].GetReason() != tc.reason {
				t.Errorf("reason = %s, want %s", bad[0].GetReason(), tc.reason)
			}
			// The report echoes the row, so a person can read what failed.
			if bad[0].GetRaw() == "" {
				t.Error("the report echoes nothing")
			}
		})
	}
}

// TestMoxfieldRefusesAHeaderItDoesNotKnow: a file with no key columns
// fails whole, and it never reports every row as bad.
func TestMoxfieldRefusesAHeaderItDoesNotKnow(t *testing.T) {
	_, _, err := ParseMoxfieldCSV(strings.NewReader("First,Last\nAnn,Lee\n"))
	if err == nil {
		t.Fatal("a header with no key columns parsed")
	}
	if !strings.Contains(err.Error(), "moxfield") {
		t.Errorf("err = %v, want the format named", err)
	}
}

// TestTheRealExportDetectsAsMoxfield closes the loop: the reader drops
// this file and never names the app it came from (D-647).
func TestTheRealExportDetectsAsMoxfield(t *testing.T) {
	content, err := os.ReadFile(moxfieldFixture)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Detect(content)
	if err != nil {
		t.Fatalf("the real export does not detect: %v", err)
	}
	if got != mtgv1.ImportSource_IMPORT_SOURCE_MOXFIELD_CSV {
		t.Errorf("Detect = %s, want the Moxfield source", got)
	}
	// And the real ManaBox export still detects as ManaBox, so the new
	// signature took nothing from the old one.
	manabox, err := os.ReadFile("testdata/manabox_collection.csv")
	if err != nil {
		t.Fatal(err)
	}
	if got, err := Detect(manabox); err != nil || got != mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV {
		t.Errorf("the ManaBox export detects as %s (%v), want ManaBox", got, err)
	}
}

// TestAFormatWithNoScryfallIdStillCarriesItsPrinting holds the gap PR-34
// opened and closed. Moxfield writes no Scryfall id column, so a row of
// it resolves on the set code and the collector number alone.
//
// The entry then carried no printing at all. The binder reads that field
// for the art and the price of what the reader owns (D-299),
// OwnedPrintings drops an entry without one, and the binder tile keys
// itself on it, so every tile of a Moxfield collection shared one key.
func TestAFormatWithNoScryfallIdStillCarriesItsPrinting(t *testing.T) {
	idx := fixtureIndex(t)
	// One printing of the fixture, named the way Moxfield names it: a
	// set code and a collector number, and no id.
	var set, num, wantID string
	for _, p := range printingsForTest(t) {
		if p.SetCode != "" && p.CollectorNumber != "" && p.ScryfallID != "" {
			set, num, wantID = p.SetCode, p.CollectorNumber, p.ScryfallID
			break
		}
	}
	if wantID == "" {
		t.Fatal("the printing fixture holds no printing with a set and a number")
	}
	rows := []Row{{Line: 2, Raw: "row", Name: "", SetCode: set, Collector: num,
		Quantity: 1, Finish: mtgv1.Finish_FINISH_NORMAL,
		Condition: mtgv1.Condition_CONDITION_NEAR_MINT, Language: "en"}}
	entries, bad := Resolve(rows, idx)
	if len(bad) != 0 {
		t.Fatalf("the row did not resolve: %v", bad[0])
	}
	if len(entries) != 1 {
		t.Fatalf("%d entries, want 1", len(entries))
	}
	if got := entries[0].GetScryfallId(); got != wantID {
		t.Errorf("scryfall id = %q, want %q from the printing the pair named", got, wantID)
	}
	// The two readers that a missing id silently disabled.
	if _, ok := idx.Printing(entries[0].GetScryfallId()); !ok {
		t.Error("the entry names no printing the index holds, so the binder shows no art")
	}
	if n := len(OwnedPrintings(entries)); n != 1 {
		t.Errorf("OwnedPrintings covers %d oracle ids, want 1", n)
	}
}

// TestAScryfallIdOnTheRowIsKept: a format that writes the id keeps it,
// and the set-and-number path never rewrites it.
func TestAScryfallIdOnTheRowIsKept(t *testing.T) {
	idx := fixtureIndex(t)
	var set, num, id string
	for _, p := range printingsForTest(t) {
		if p.SetCode != "" && p.CollectorNumber != "" && p.ScryfallID != "" {
			set, num, id = p.SetCode, p.CollectorNumber, p.ScryfallID
			break
		}
	}
	rows := []Row{{Line: 2, Raw: "row", ScryfallID: id, SetCode: set, Collector: num,
		Quantity: 1, Language: "en"}}
	entries, bad := Resolve(rows, idx)
	if len(bad) != 0 || len(entries) != 1 {
		t.Fatalf("the row did not resolve: %v", bad)
	}
	if got := entries[0].GetScryfallId(); got != id {
		t.Errorf("scryfall id = %q, want the row's own %q", got, id)
	}
}
