package triage

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/harvest"
	"github.com/nkramber/decktome/go/internal/importfault"
)

// importRecord is the harvest record of an import report on the deck
// page, with the fault the server read from the file.
func importRecord(t *testing.T, content string) harvest.Record {
	t.Helper()
	fault, err := importfault.Read(mtgv1.ImportPage_IMPORT_PAGE_DECK, []byte(content))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	raw, err := protojson.Marshal(fault)
	if err != nil {
		t.Fatal(err)
	}
	return harvest.Record{ID: "f1", Kind: "import", Verdict: "down", Reasons: []string{"parse_fault"}, Text: "archidekt", Context: "fault", Import: raw}
}

// An import report routes to I1 with no judge, and its case is a parser
// fixture in the folder of its page (D-888).
func TestAnImportReportWritesAParserFixture(t *testing.T) {
	rec := importRecord(t, "1 Sol Ring\nthis reads as no card\n1 Island\n")
	results := Run(context.Background(), []harvest.Record{rec}, nil, nil, func(string) (int, error) { return 7, nil })
	if len(results) != 1 || results[0].Err != nil {
		t.Fatalf("results = %+v", results)
	}
	c := results[0].Case
	if c.Class != "I1" || c.Artifact != AParseFixture || c.Target != TargetDeckReports || c.ID != 7 {
		t.Fatalf("case = %s %s %s %d", c.Class, c.Artifact, c.Target, c.ID)
	}
	var text string
	if err := json.Unmarshal(c.Body, &text); err != nil {
		t.Fatal(err)
	}
	if text != "1 Sol Ring\nthis reads as no card\n" {
		t.Errorf("fixture = %q, want the header and the bad row", text)
	}
	if len(c.Gaps) != 0 {
		t.Errorf("gaps = %v, want none", c.Gaps)
	}
}

// The fixture the triage writes fails the parser test until the parser
// reads it, so the fix has a test before it starts (D-888, D-891).
func TestTheWrittenFixtureStillHoldsTheFault(t *testing.T) {
	rec := importRecord(t, "hello\nworld\n")
	c, err := CaseOf(RouteOf(rec), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(t.TempDir(), "deck")
	if n, err := NextFixtureID(dir); err != nil || n != 1 {
		t.Fatalf("NextFixtureID of no folder = %d, %v, want 1", n, err)
	}
	path, err := WriteFixture(dir, 1, c.Body)
	if err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := importfault.Read(mtgv1.ImportPage_IMPORT_PAGE_DECK, content); err != nil {
		t.Fatalf("the fixture reads as %v, want the fault it came from", err)
	}
	if n, err := NextFixtureID(dir); err != nil || n != 2 {
		t.Errorf("NextFixtureID = %d, %v, want 2", n, err)
	}
	if _, err := WriteFixture(dir, 1, c.Body); err == nil {
		t.Error("WriteFixture wrote over a fixture")
	}
}

// A report that kept fewer rows than do not parse names the gap, and a
// report with no fault writes no case.
func TestAFixtureNamesWhatItLeftOut(t *testing.T) {
	rec := importRecord(t, "hello\nworld\n")
	rec.Text = ""
	var f mtgv1.ImportFault
	if err := protojson.Unmarshal(rec.Import, &f); err != nil {
		t.Fatal(err)
	}
	f.BadRowCount = 900
	rec.Import, _ = protojson.Marshal(&f)
	c, err := CaseOf(RouteOf(rec), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(c.Gaps) != 2 {
		t.Errorf("gaps = %v, want the cut rows and the missing service", c.Gaps)
	}
	rec.Import = nil
	if _, err := CaseOf(RouteOf(rec), 1, nil); err == nil {
		t.Error("a report with no fault wrote a case")
	}
}
