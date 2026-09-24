package importfault

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

const (
	collection = mtgv1.ImportPage_IMPORT_PAGE_COLLECTION
	deck       = mtgv1.ImportPage_IMPORT_PAGE_DECK
	manaBox    = "Binder Name,Binder Type,Name,Set code,Set name,Collector number,Foil,Rarity,Quantity,ManaBox ID,Scryfall ID,Purchase price,Misprint,Altered,Condition,Language,Purchase price currency,Added"
)

// row is one ManaBox row with a foil value and a quantity of the test's
// choice.
func row(name, foil, qty string) string {
	return fmt.Sprintf("Main,binder,%s,ECL,Lorwyn Eclipsed,181,%s,common,%s,110386,a79649c4-559e-4306-a102-5fd8750629c7,0.13,false,false,near_mint,en,USD,2026-04-26T05:28:38.780Z", name, foil, qty)
}

func TestAGoodFileHasNoFault(t *testing.T) {
	content := manaBox + "\n" + row("Lys Alana Informant", "normal", "3") + "\n"
	if _, err := Read(collection, []byte(content)); !errors.Is(err, ErrNoFault) {
		t.Fatalf("Read = %v, want ErrNoFault", err)
	}
	if _, err := Read(deck, []byte("1 Sol Ring\n1 Island\n")); !errors.Is(err, ErrNoFault) {
		t.Fatalf("Read deck = %v, want ErrNoFault", err)
	}
}

// A file no format claims reads as a whole-file fault: the error of the
// server, and every row after the header.
func TestAFileNoFormatClaimsKeepsEveryRow(t *testing.T) {
	content := "Title,Count\r\nSol Ring,1\r\n\r\nIsland,4\r\n"
	f, err := Read(collection, []byte(content))
	if err != nil {
		t.Fatal(err)
	}
	if f.GetError() == "" {
		t.Error("the fault holds no error")
	}
	if f.GetHeader() != "Title,Count" {
		t.Errorf("header = %q", f.GetHeader())
	}
	if f.GetRowCount() != 2 || f.GetBadRowCount() != 2 || len(f.GetRows()) != 2 {
		t.Fatalf("rows %d, bad %d, kept %d, want 2 each", f.GetRowCount(), f.GetBadRowCount(), len(f.GetRows()))
	}
	if r := f.GetRows()[1]; r.GetLine() != 4 || r.GetRaw() != "Island,4" {
		t.Errorf("second row = line %d %q, want line 4 %q", r.GetLine(), r.GetRaw(), "Island,4")
	}
	if f.GetByteCount() != int64(len(content)) {
		t.Errorf("byte count = %d", f.GetByteCount())
	}
}

// A file that reads keeps the rows that do not parse alone, as the
// source line and not the text the parser cut (D-885, D-887).
func TestAFileThatReadsKeepsItsBadRows(t *testing.T) {
	long := strings.Repeat("Long Name ", 30)
	content := strings.Join([]string{
		manaBox,
		row("Lys Alana Informant", "normal", "3"),
		row(long, "shiny", "1"),
		row("Deepway Navigator", "normal", "many"),
	}, "\n")
	f, err := Read(collection, []byte(content))
	if err != nil {
		t.Fatal(err)
	}
	if f.GetError() != "" {
		t.Errorf("a file that reads holds the error %q", f.GetError())
	}
	if f.GetRowCount() != 3 || f.GetBadRowCount() != 2 || len(f.GetRows()) != 2 {
		t.Fatalf("rows %d, bad %d, kept %d, want 3, 2, 2", f.GetRowCount(), f.GetBadRowCount(), len(f.GetRows()))
	}
	first, second := f.GetRows()[0], f.GetRows()[1]
	if first.GetLine() != 3 || first.GetReason() != mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE {
		t.Errorf("first row = line %d %v, want line 3 UNKNOWN_VALUE", first.GetLine(), first.GetReason())
	}
	if first.GetRaw() != row(long, "shiny", "1") {
		t.Errorf("the first row is not the source line: %q", first.GetRaw())
	}
	if second.GetLine() != 4 || second.GetReason() != mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW {
		t.Errorf("second row = line %d %v, want line 4 BAD_ROW", second.GetLine(), second.GetReason())
	}
}

func TestADeckListKeepsItsBadLines(t *testing.T) {
	f, err := Read(deck, []byte("1 Sol Ring\nthis reads as no card\n1 Island\n"))
	if err != nil {
		t.Fatal(err)
	}
	if len(f.GetRows()) != 1 || f.GetRows()[0].GetLine() != 2 {
		t.Fatalf("rows = %v, want line 2 alone", f.GetRows())
	}
	if f.GetHeader() != "1 Sol Ring" {
		t.Errorf("header = %q, want the first line", f.GetHeader())
	}
}

func TestADeckListWithNoCardLineDoesNotRead(t *testing.T) {
	f, err := Read(deck, []byte("hello\nworld\n"))
	if err != nil {
		t.Fatal(err)
	}
	if f.GetError() != ErrNoCardLine.Error() {
		t.Errorf("error = %q, want %q", f.GetError(), ErrNoCardLine)
	}
	if len(f.GetRows()) != 1 || f.GetRows()[0].GetRaw() != "world" {
		t.Errorf("rows = %v, want the line after the header", f.GetRows())
	}
}

// The report keeps at most MaxRows rows, and it counts every one (D-885).
func TestTheReportKeepsAtMost500Rows(t *testing.T) {
	var b strings.Builder
	b.WriteString("Title,Count\n")
	for i := range MaxRows + 100 {
		fmt.Fprintf(&b, "Card %d,1\n", i)
	}
	f, err := Read(collection, []byte(b.String()))
	if err != nil {
		t.Fatal(err)
	}
	if len(f.GetRows()) != MaxRows || f.GetBadRowCount() != MaxRows+100 {
		t.Fatalf("kept %d of %d, want %d of %d", len(f.GetRows()), f.GetBadRowCount(), MaxRows, MaxRows+100)
	}
}

func TestACutRowStaysValidText(t *testing.T) {
	s := strings.Repeat("é", MaxRowBytes)
	got := cut(s)
	if len(got) > MaxRowBytes || !strings.HasPrefix(s, got) {
		t.Fatalf("cut is %d bytes, or not a prefix", len(got))
	}
}

func TestReadRefusesABadArgument(t *testing.T) {
	cases := map[string]struct {
		page    mtgv1.ImportPage
		content []byte
	}{
		"no page": {mtgv1.ImportPage_IMPORT_PAGE_UNSPECIFIED, []byte("x")},
		"no file": {collection, nil},
		"too big": {collection, make([]byte, MaxContentBytes+1)},
	}
	for name, c := range cases {
		if _, err := Read(c.page, c.content); err == nil || errors.Is(err, ErrNoFault) {
			t.Errorf("%s: Read = %v, want a bad argument", name, err)
		}
	}
}

func TestUnreadableCarriesItsDetail(t *testing.T) {
	err := Unreadable(errors.New("no format"))
	if err.Code() != connect.CodeInvalidArgument {
		t.Errorf("code = %v", err.Code())
	}
	if !IsUnreadable(err) {
		t.Error("the detail is lost")
	}
	if IsUnreadable(connect.NewError(connect.CodeInvalidArgument, errors.New("too large"))) {
		t.Error("a plain error reads as unreadable")
	}
}

// TestEveryReportReads reads each fixture the triage wrote from a report
// (D-888, D-891). A fixture fails until the parser reads it. The folder
// name is the page, and each file is the header and the kept rows.
func TestEveryReportReads(t *testing.T) {
	for dir, page := range map[string]mtgv1.ImportPage{"collection": collection, "deck": deck} {
		files, err := filepath.Glob(filepath.Join("testdata", "reports", dir, "*.txt"))
		if err != nil {
			t.Fatal(err)
		}
		for _, name := range files {
			content, err := os.ReadFile(name)
			if err != nil {
				t.Fatal(err)
			}
			f, err := Read(page, content)
			if errors.Is(err, ErrNoFault) {
				continue
			}
			if err != nil {
				t.Errorf("%s: %v", name, err)
				continue
			}
			t.Errorf("%s: %q, %d of %d rows do not parse", name, f.GetError(), f.GetBadRowCount(), f.GetRowCount())
		}
	}
}
