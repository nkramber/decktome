package gatekit

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
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
