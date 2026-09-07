package gatekit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
)

func TestSpendGuard(t *testing.T) {
	cases := []struct {
		name  string
		value string
		allow bool
	}{
		{"unset refuses", "", false},
		{"zero refuses", "0", false},
		{"one allows", "1", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("GATEKIT_TEST_GUARD", tc.value)
			err := SpendGuard("GATEKIT_TEST_GUARD")
			if (err == nil) != tc.allow {
				t.Errorf("err = %v, want allow %v", err, tc.allow)
			}
			if err != nil && !strings.Contains(err.Error(), "GATEKIT_TEST_GUARD=1") {
				t.Errorf("the refusal does not name the variable: %v", err)
			}
		})
	}
}

func TestEnvForcesRealKeys(t *testing.T) {
	t.Setenv(llm.EnvRequireKeys, "0")
	t.Setenv("GATEKIT_TEST_OTHER", "x")
	cases := []struct{ key, want string }{
		{llm.EnvRequireKeys, "1"},
		{"GATEKIT_TEST_OTHER", "x"},
	}
	for _, tc := range cases {
		if got := Env(tc.key); got != tc.want {
			t.Errorf("Env(%q) = %q, want %q", tc.key, got, tc.want)
		}
	}
}

func TestWordMaps(t *testing.T) {
	formats := []struct {
		in   string
		want mtgv1.FormatId
	}{
		{"Commander", mtgv1.FormatId_FORMAT_ID_COMMANDER},
		{" modern ", mtgv1.FormatId_FORMAT_ID_MODERN},
		{"pauper", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
	}
	for _, tc := range formats {
		if got := FormatID(tc.in); got != tc.want {
			t.Errorf("FormatID(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
	if got := Colors([]string{"w", "X", "G"}); len(got) != 2 || got[0] != mtgv1.Color_COLOR_W || got[1] != mtgv1.Color_COLOR_G {
		t.Errorf("Colors = %v", got)
	}
	pools := []struct {
		in   string
		want mtgv1.PoolRule
	}{
		{"owned_only", mtgv1.PoolRule_POOL_RULE_OWNED_ONLY},
		{"owned_first", mtgv1.PoolRule_POOL_RULE_OWNED_FIRST},
		{"", mtgv1.PoolRule_POOL_RULE_ANY_CARD},
	}
	for _, tc := range pools {
		if got := PoolRuleID(tc.in); got != tc.want {
			t.Errorf("PoolRuleID(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
	if OrNone("") != "none" || OrNone("x") != "x" {
		t.Error("OrNone")
	}
}

func TestSnapshotDirNeedsTheVariable(t *testing.T) {
	t.Setenv("CARDS_SNAPSHOT_DIR", "")
	if _, err := SnapshotDir(); err == nil {
		t.Error("an empty CARDS_SNAPSHOT_DIR was accepted")
	}
}

func TestRefuseExisting(t *testing.T) {
	dir := t.TempDir()
	taken := filepath.Join(dir, "taken.md")
	if err := os.WriteFile(taken, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := RefuseExisting("", filepath.Join(dir, "free.md"), os.DevNull); err != nil {
		t.Errorf("a free path was refused: %v", err)
	}
	err := RefuseExisting(filepath.Join(dir, "free.md"), taken)
	if err == nil || !strings.Contains(err.Error(), "D-65") {
		t.Errorf("err = %v, want a D-65 refusal of %s", err, taken)
	}
}

func TestParseIDs(t *testing.T) {
	got, err := ParseIDs(" 3, 1 ,12")
	if err != nil || len(got) != 3 || got[0] != 3 || got[1] != 1 || got[2] != 12 {
		t.Errorf("ParseIDs = %v, %v", got, err)
	}
	if got, err := ParseIDs(""); err != nil || got != nil {
		t.Errorf("empty = %v, %v", got, err)
	}
	if _, err := ParseIDs("1,x"); err == nil {
		t.Error("a word was accepted as an id")
	}
}

func TestCostWord(t *testing.T) {
	if got := CostWord(llm.Report{}); got != "unpriced" {
		t.Errorf("nil cost = %q, want unpriced", got)
	}
	c := 0.12345
	if got := CostWord(llm.Report{CostUSD: &c}); got != "$0.1235" {
		t.Errorf("cost = %q", got)
	}
	zero := 0.0
	if got := CostWord(llm.Report{CostUSD: &zero}); got != "$0.0000" {
		t.Errorf("measured zero = %q, want $0.0000", got)
	}
}

func TestDeckHelpers(t *testing.T) {
	d := &mtgv1.Deck{
		Cards:     []*mtgv1.DeckCard{{Name: "a", Count: 4}, {Name: "b", Count: 1}},
		Sideboard: []*mtgv1.DeckCard{{Name: "c", Count: 15}},
		Validation: &mtgv1.ValidationResult{Findings: []*mtgv1.Finding{
			{Code: "deck_size", Severity: mtgv1.Severity_SEVERITY_BLOCK},
			{Code: "note", Severity: mtgv1.Severity_SEVERITY_INFO},
		}},
	}
	if CountCards(d) != 5 || CountSideboard(d) != 15 {
		t.Errorf("counts = %d main, %d side", CountCards(d), CountSideboard(d))
	}
	if got := BlockFindings(d); len(got) != 1 || got[0].GetCode() != "deck_size" {
		t.Errorf("blocks = %v", got)
	}
	if CountCards(nil) != 0 || len(BlockFindings(nil)) != 0 {
		t.Error("a nil deck must count as empty")
	}
	if p := PowerLevel(3, "casual"); p.GetBracket() != 3 {
		t.Errorf("bracket wins: %v", p)
	}
	if p := PowerLevel(0, "Tournament"); p.GetSixtyStep() != mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT {
		t.Errorf("step = %v", p)
	}
	if PowerLevel(0, "") != nil || PowerLevel(0, "mythic") != nil {
		t.Error("no bracket and no step must give nil")
	}
	letters := ColorLetters([]mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_UNSPECIFIED, mtgv1.Color_COLOR_G})
	if strings.Join(letters, "") != "WG" {
		t.Errorf("letters = %v", letters)
	}
}
