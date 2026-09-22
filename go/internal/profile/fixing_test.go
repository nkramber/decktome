package profile

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/spellbook"
)

// The fixing floor reads the power and the count of deck colors (F-33,
// D-799). A deck of one color and the casual step hold no floor.
func TestFixingFloorReadsThePowerAndTheColors(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	commander, modern := mtgv1.FormatId_FORMAT_ID_COMMANDER, mtgv1.FormatId_FORMAT_ID_MODERN
	for _, tc := range []struct {
		name   string
		format mtgv1.FormatId
		power  *mtgv1.PowerLevel
		colors int
		want   float64
	}{
		{"bracket 1, two colors", commander, bracket(1), 2, 7},
		{"bracket 3, two colors", commander, bracket(3), 2, 11},
		{"no bracket reads bracket 3", commander, nil, 2, 11},
		{"bracket 5, five colors", commander, bracket(5), 5, 21},
		{"one color", commander, bracket(3), 1, 0},
		{"casual", modern, step(mtgv1.SixtyStep_SIXTY_STEP_CASUAL), 2, 0},
		{"fnm, three colors", modern, step(mtgv1.SixtyStep_SIXTY_STEP_FNM), 3, 10},
		{"tournament, two colors", modern, step(mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT), 2, 7},
	} {
		if got := b.FixingFloor(tc.format, tc.power, tc.colors); got != tc.want {
			t.Errorf("%s: floor %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestFixingRowRefusesABadColorCount(t *testing.T) {
	for _, row := range []map[string]float64{{"1": 3}, {"6": 3}, {"two": 3}, {"2": -1}} {
		if _, err := fixingRow(row); err == nil {
			t.Errorf("row %v parsed", row)
		}
	}
}

// The prompt names the floor of each count of colors, because a 60-card
// prompt goes out before the model picks the colors.
func TestFixingLineNamesEveryColorCount(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Join(b.Lines(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(3)), "\n")
	for _, want := range []string{"fixing lands", "at least 11 in a two-color deck", "and 17 in a five-color deck", "A basic land counts as none."} {
		if !strings.Contains(lines, want) {
			t.Errorf("the bracket 3 lines miss %q:\n%s", want, lines)
		}
	}
	casual := strings.Join(b.Lines(mtgv1.FormatId_FORMAT_ID_MODERN, step(mtgv1.SixtyStep_SIXTY_STEP_CASUAL)), "\n")
	if strings.Contains(casual, "fixing lands") {
		t.Errorf("the casual step holds no floor, and its lines name one:\n%s", casual)
	}
}

// A two-color deck of basics alone reads a fixing row under its floor,
// and the finding names the count and the floor.
func TestReadFindsAShortFixingCount(t *testing.T) {
	rows := shaped()
	rows[0].n, rows[3].n = 20, 0
	p := newProfiler(t, &fakeClassifier{res: &spellbook.Result{BracketTag: "C"}})
	prof, findings := p.Read(context.Background(), deckOf(3, rows...), source(testCards))
	f := feature(t, prof, KeyFixingLand)
	if f.GetValue() != 3 || f.GetLow() != 11 || !f.GetOffBand() {
		t.Errorf("fixing row %v", f)
	}
	var found bool
	for _, x := range findings {
		if x.GetCode() == CodeOffBand && strings.Contains(x.GetMessage(), "two or more of the deck's colors is 3, and bracket 3 wants 11 or more") {
			found = true
		}
	}
	if !found {
		t.Errorf("want the fixing finding, got %v", findings)
	}
}

// A deck of one color needs no fixing, so it reads no row.
func TestReadGivesAOneColorDeckNoFixingRow(t *testing.T) {
	d := deckOf(3, row{plains, 36, mtgv1.CardRole_CARD_ROLE_LAND}, row{angel, 63, mtgv1.CardRole_CARD_ROLE_THREAT})
	d.CommanderOracleIds = []string{angel.OracleId}
	prof := newProfiler(t, &fakeClassifier{res: &spellbook.Result{}}).Measure(d, source(testCards))
	for _, f := range prof.GetFeatures() {
		if f.GetKey() == KeyFixingLand {
			t.Errorf("a one-color deck reads a fixing row: %v", f)
		}
	}
}
