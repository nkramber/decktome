package collections

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestArenaFinishMarker is D-246. A deck export marks a foil with "*F*"
// after the collector number. The marker comes off before the set and
// number pair is read, and an unresolved row carries its raw text.
func TestArenaFinishMarker(t *testing.T) {
	const deck = `// COMMANDER
1 Captain America, Team Leader (MSC) 5 *F*

1 Sol Ring (MSC) 211
1 Arcane Signet (MSC) 191 *E*
2 Island (MSH) 290
`
	rows, bad, err := ParseArenaText(strings.NewReader(deck))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(bad) != 0 {
		t.Fatalf("unparsed rows: %+v", bad)
	}
	if len(rows) != 4 {
		t.Fatalf("rows = %d, want 4", len(rows))
	}
	want := []struct {
		name      string
		set, coll string
		finish    mtgv1.Finish
		qty       int
	}{
		{"Captain America, Team Leader", "MSC", "5", mtgv1.Finish_FINISH_FOIL, 1},
		{"Sol Ring", "MSC", "211", mtgv1.Finish_FINISH_NORMAL, 1},
		{"Arcane Signet", "MSC", "191", mtgv1.Finish_FINISH_ETCHED, 1},
		{"Island", "MSH", "290", mtgv1.Finish_FINISH_NORMAL, 2},
	}
	for i, w := range want {
		got := rows[i]
		if got.Name != w.name || got.SetCode != w.set || got.Collector != w.coll {
			t.Errorf("row %d = %q (%s) %s, want %q (%s) %s",
				i, got.Name, got.SetCode, got.Collector, w.name, w.set, w.coll)
		}
		if got.Finish != w.finish {
			t.Errorf("row %d finish = %v, want %v", i, got.Finish, w.finish)
		}
		if got.Quantity != w.qty {
			t.Errorf("row %d quantity = %d, want %d", i, got.Quantity, w.qty)
		}
		// The raw line must name the row, or an unresolved card says
		// nothing about which line failed.
		if got.Raw == "" {
			t.Errorf("row %d carries no raw line", i)
		}
	}
}

// TestArenaLineOfOnlyAFinishMarker keeps a malformed line out of the
// deck. A quantity and a marker name no card.
func TestArenaLineOfOnlyAFinishMarker(t *testing.T) {
	rows, bad, err := ParseArenaText(strings.NewReader("1 *F*\n"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(rows) != 0 || len(bad) != 1 {
		t.Errorf("rows = %d and bad = %d, want 0 and 1", len(rows), len(bad))
	}
}

// TestBrokenCSVRecordNamesItsLine: a record encoding/csv can not parse
// has no fields to echo, so the report names the line instead.
func TestBrokenCSVRecordNamesItsLine(t *testing.T) {
	const csvText = "Name,Set code\nGood Card,abc\n\"broken,abc\nAnother,abc\n"
	_, bad, err := ParseManaBoxCSV(strings.NewReader(csvText))
	if err != nil {
		t.Fatal(err)
	}
	if len(bad) != 1 {
		t.Fatalf("bad rows = %v, want one", bad)
	}
	if bad[0].GetReason() != mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW {
		t.Errorf("reason = %v", bad[0].GetReason())
	}
	if !strings.HasPrefix(bad[0].GetRaw(), "unparseable record at line ") || bad[0].GetLine() == 0 {
		t.Errorf("raw = %q line %d, want the line named", bad[0].GetRaw(), bad[0].GetLine())
	}
}
