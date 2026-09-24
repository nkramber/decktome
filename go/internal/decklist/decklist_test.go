package decklist

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

func parseFile(t *testing.T, path string) *List {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	l, err := Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func sum(l *List, sec Section) int {
	n := 0
	for _, line := range l.Lines {
		if line.Section == sec {
			n += line.Quantity
		}
	}
	return n
}

func find(l *List, name string) (Line, bool) {
	for _, line := range l.Lines {
		if line.Name == name {
			return line, true
		}
	}
	return Line{}, false
}

// TestParseArchidektExport reads the Archidekt export of the owner
// (D-856). The category Commander{top} marks the commander, and every
// other category is dropped (D-847, D-848).
func TestParseArchidektExport(t *testing.T) {
	l := parseFile(t, "testdata/archidekt_living_weapon.txt")
	if len(l.Bad) != 0 {
		t.Fatalf("bad lines: %v", l.Bad)
	}
	if len(l.Lines) != 80 || sum(l, Main)+sum(l, Commander) != 100 {
		t.Fatalf("lines = %d, cards = %d, want 80 and 100", len(l.Lines), sum(l, Main)+sum(l, Commander))
	}
	if !l.Marked || sum(l, Commander) != 1 {
		t.Fatalf("marked = %v, commander cards = %d", l.Marked, sum(l, Commander))
	}
	cmd, _ := find(l, "Ekthi, Contaminator Priest")
	if cmd.Section != Commander || cmd.SetCode != "mbc" || cmd.Collector != "1" {
		t.Errorf("commander line = %+v", cmd)
	}
	plains, _ := find(l, "Plains")
	if plains.Quantity != 21 || plains.Section != Main {
		t.Errorf("Plains = %+v", plains)
	}
	foil, _ := find(l, "Forge Anew")
	if foil.Finish != mtgv1.Finish_FINISH_FOIL || foil.Collector != "7002" {
		t.Errorf("foil line = %+v", foil)
	}
	dfc, ok := find(l, "Dowsing Dagger // Lost Vale")
	if !ok || dfc.SetCode != "plst" || dfc.Collector != "XLN-235" {
		t.Errorf("double-faced line = %+v", dfc)
	}
	for _, line := range l.Lines {
		if strings.ContainsAny(line.Name, "[]") {
			t.Errorf("a category stayed in the name: %q", line.Name)
		}
	}
}

// TestParseArenaExport reads the Arena export of the owner (D-860). The
// line "// COMMANDER" opens the commander section, and a blank line ends
// it with no Deck header after it.
func TestParseArenaExport(t *testing.T) {
	l := parseFile(t, "testdata/arena_turtle_power.txt")
	if len(l.Bad) != 0 {
		t.Fatalf("bad lines: %v", l.Bad)
	}
	if !l.Marked || sum(l, Commander) != 1 || sum(l, Main) != 99 {
		t.Fatalf("marked = %v, commander = %d, main = %d", l.Marked, sum(l, Commander), sum(l, Main))
	}
	cmd, _ := find(l, "Heroes in a Half Shell")
	if cmd.Section != Commander || cmd.Finish != mtgv1.Finish_FINISH_FOIL || cmd.SetCode != "TMC" {
		t.Errorf("commander line = %+v", cmd)
	}
	if don, _ := find(l, "Donatello, the Brains"); don.Section != Main {
		t.Errorf("a foil card of the 99 left the deck: %+v", don)
	}
}

// TestParseArenaSections covers each header of an Arena list, the About
// name, a comment, and a line that reads as no card.
func TestParseArenaSections(t *testing.T) {
	text := "About\nName Rats Galore\n\nCommander\n1 Marrow-Gnawer (CHK) 124\n\nCompanion\n1 Lurrus of the Dream-Den (IKO) 226\n\n" +
		"Deck\n// the rats\n40 Relentless Rats (M11) 106\n59 Swamp\nnot a card line\n\nSideboard\n1 Duress (M19) 94\n"
	l, err := Parse(strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}
	if l.Name != "Rats Galore" {
		t.Errorf("name = %q", l.Name)
	}
	if !l.Marked || sum(l, Commander) != 1 || sum(l, Companion) != 1 || sum(l, Main) != 99 || sum(l, Sideboard) != 1 {
		t.Errorf("sections: commander %d, companion %d, main %d, sideboard %d",
			sum(l, Commander), sum(l, Companion), sum(l, Main), sum(l, Sideboard))
	}
	if len(l.Bad) != 1 || l.Bad[0].GetLine() != 14 {
		t.Errorf("bad = %v", l.Bad)
	}
}

// TestParseUnmarkedList is D-847: a list with no commander mark reads as
// the main deck alone, and the import then asks for the commander.
func TestParseUnmarkedList(t *testing.T) {
	l, err := Parse(strings.NewReader("1x Sol Ring (c21) 263 [Ramp]\n4 Lightning Bolt\n"))
	if err != nil {
		t.Fatal(err)
	}
	if l.Marked || sum(l, Main) != 5 {
		t.Errorf("marked = %v, main = %d", l.Marked, sum(l, Main))
	}
}

// TestResolveSumsWithinASection reads two printings of one card as one
// entry, and keeps the commander apart from the deck.
func TestResolveSumsWithinASection(t *testing.T) {
	bolt := &mtgv1.Card{OracleId: "o-bolt", Name: "Lightning Bolt"}
	leader := &mtgv1.Card{OracleId: "o-lead", Name: "Leader", CanBeCommander: true}
	idx := cards.NewIndex([]*mtgv1.Card{bolt, leader}, []cards.Printing{
		{ScryfallID: "p1", OracleID: "o-bolt", Name: "Lightning Bolt", SetCode: "m11", CollectorNumber: "149", Layout: "normal"},
		{ScryfallID: "p2", OracleID: "o-bolt", Name: "Lightning Bolt", SetCode: "sta", CollectorNumber: "42", Layout: "normal"},
		{ScryfallID: "p3", OracleID: "o-lead", Name: "Leader", SetCode: "abc", CollectorNumber: "1", Layout: "normal"},
	}, nil, time.Now())
	l, err := Parse(strings.NewReader("1x Leader (abc) 1 [Commander{top}]\n2x Lightning Bolt (m11) 149 [Removal]\n1 Lightning Bolt (STA) 42\n1 No Such Card\n"))
	if err != nil {
		t.Fatal(err)
	}
	got, bad := Resolve(l, idx)
	if len(got) != 2 || got[0].Card.GetName() != "Leader" || got[0].Section != Commander {
		t.Fatalf("entries = %+v", got)
	}
	if got[1].Card.GetName() != "Lightning Bolt" || got[1].Count != 3 || got[1].Section != Main {
		t.Errorf("bolt = %+v", got[1])
	}
	if len(bad) != 1 || !strings.Contains(bad[0].GetRaw(), "No Such Card") {
		t.Errorf("bad = %v", bad)
	}
}

var (
	snapshotOnce sync.Once
	snapshotIdx  *cards.Index
	snapshotErr  error
)

// snapshotIndex loads the card snapshot under CARDS_SNAPSHOT_DIR, once.
// The tests skip without the variable.
func snapshotIndex(t *testing.T) *cards.Index {
	t.Helper()
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		t.Skip("set CARDS_SNAPSHOT_DIR to run the snapshot tests")
	}
	snapshotOnce.Do(func() {
		quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
		snapshotIdx, snapshotErr = cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, quiet)
	})
	if snapshotErr != nil || snapshotIdx == nil {
		t.Fatalf("load index: %v", snapshotErr)
	}
	return snapshotIdx
}

// TestResolveRealExports reads both exports of the owner against the
// card snapshot. Every line resolves, the foil line and the collector
// number XLN-235 included.
func TestResolveRealExports(t *testing.T) {
	idx := snapshotIndex(t)
	for _, tc := range []struct{ file, commander string }{
		{"testdata/archidekt_living_weapon.txt", "Ekthi, Contaminator Priest"},
		{"testdata/arena_turtle_power.txt", "Heroes in a Half Shell"},
	} {
		got, bad := Resolve(parseFile(t, tc.file), idx)
		if len(bad) != 0 {
			t.Errorf("%s: unresolved %v", tc.file, bad)
		}
		n := 0
		var cmd []string
		for _, e := range got {
			n += e.Count
			if e.Section == Commander {
				cmd = append(cmd, e.Card.GetName())
			}
		}
		if n != 100 || len(cmd) != 1 || cmd[0] != tc.commander {
			t.Errorf("%s: %d cards, commander %v", tc.file, n, cmd)
		}
	}
}

// TestParseRefusesALongList is P2-1 of the review of #220: a list over
// maxLines fails whole, and no line past the cap goes missing (D-846).
func TestParseRefusesALongList(t *testing.T) {
	text := strings.Repeat("1 Plains\n", maxLines+1)
	l, err := Parse(strings.NewReader(text))
	if !errors.Is(err, ErrTooManyLines) || l != nil {
		t.Fatalf("list is nil = %v, err = %v, want ErrTooManyLines", l == nil, err)
	}
	if _, err := Parse(strings.NewReader(strings.Repeat("1 Plains\n", maxLines))); err != nil {
		t.Errorf("a list of %d lines failed: %v", maxLines, err)
	}
}
