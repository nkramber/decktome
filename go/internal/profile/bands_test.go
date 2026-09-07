package profile

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func bracket(n int32) *mtgv1.PowerLevel {
	return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: n}}
}

func TestLoadBandsHoldsEveryBracketAndNarrowsWithPower(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	if b.VerifiedAt == "" {
		t.Error("verified_at is empty")
	}
	for n := int32(1); n <= 5; n++ {
		table, got := b.For(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(n))
		if got != n {
			t.Errorf("bracket %d read as %d", n, got)
		}
		for _, key := range []string{KeyLand, KeyRamp, KeyDraw, KeyRemoval, KeyWipe, KeyInteraction, KeyAvgManaValue,
			KeyTutor, KeyFastMana, KeyTappedLand, KeyColorlessLand, KeyColorSources, KeyManaTurnFour, KeyHandsTwoToFourLands, KeyCommanderTurnOverMV} {
			if _, ok := table[key]; !ok {
				t.Errorf("bracket %d has no band for %s", n, key)
			}
		}
	}
	// The mana floor rises and the tapped cap falls with the bracket.
	prev, _ := b.For(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(1))
	for n := int32(2); n <= 5; n++ {
		cur, _ := b.For(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(n))
		if cur[KeyManaTurnFour].Low < prev[KeyManaTurnFour].Low {
			t.Errorf("bracket %d wants less mana than bracket %d", n, n-1)
		}
		if *cur[KeyTappedLand].High > *prev[KeyTappedLand].High {
			t.Errorf("bracket %d allows more tapped lands than bracket %d", n, n-1)
		}
		prev = cur
	}
}

func TestBandsForDefaults(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	if _, got := b.For(mtgv1.FormatId_FORMAT_ID_COMMANDER, nil); got != DefaultBracket {
		t.Errorf("no bracket reads %d, want %d", got, DefaultBracket)
	}
	if _, got := b.For(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(9)); got != DefaultBracket {
		t.Errorf("an unknown bracket reads %d, want %d", got, DefaultBracket)
	}
	table, got := b.For(mtgv1.FormatId_FORMAT_ID_MODERN, nil)
	if got != 0 || table[KeyTappedLand].High == nil || *table[KeyTappedLand].High != 12 {
		t.Errorf("a 60-card deck with no step reads casual: %d %v", got, table[KeyTappedLand])
	}
	if _, ok := table[KeyLand]; ok {
		t.Error("a 60-card deck has no land band")
	}
}

func TestBandLinesAndMidpoints(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Join(b.Lines(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(3)), "\n")
	for _, want := range []string{"average mana value of the nonland cards: 2 to 3.5", "lands that enter tapped: at most 9", "tutors, cards that search the library for a card: at most 5", "fast mana"} {
		if !strings.Contains(lines, want) {
			t.Errorf("lines lack %q:\n%s", want, lines)
		}
	}
	four := strings.Join(b.Lines(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(4)), "\n")
	if !strings.Contains(four, "tutors, cards that search the library for a card: no limit") {
		t.Errorf("bracket 4 tutors:\n%s", four)
	}
	mid := b.Midpoints(mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(3))
	if mid[KeyLand] != 36 || mid[KeyRamp] != 10 || mid[KeyWipe] != 3 {
		t.Errorf("midpoints %v", mid)
	}
	if _, ok := mid[KeyAvgManaValue]; ok {
		t.Error("a midpoint is a role count alone")
	}
	sixty := b.Lines(mtgv1.FormatId_FORMAT_ID_STANDARD, nil)
	if len(sixty) != 1 || !strings.Contains(sixty[0], "lands that enter tapped: at most 12") {
		t.Errorf("sixty lines %v", sixty)
	}
}

func TestBandHolds(t *testing.T) {
	hi := 5.0
	b := Band{Low: 2, High: &hi}
	if !b.Holds(2) || !b.Holds(5) || b.Holds(1.9) || b.Holds(5.1) {
		t.Error("closed band")
	}
	open := Band{Low: 0.8}
	if !open.Holds(100) || open.Holds(0.7) {
		t.Error("open band")
	}
	if bandWords(KeyColorSources, open) != "0.8 or more" || bandWords(KeyTutor, Band{Low: 0}) != "no limit" {
		t.Error("band words")
	}
}
