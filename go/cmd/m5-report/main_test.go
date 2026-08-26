package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func fixture(t *testing.T) []scored {
	t.Helper()
	raw, err := os.ReadFile("testdata/scored.md")
	if err != nil {
		t.Fatal(err)
	}
	return parse(string(raw))
}

func TestParseReadsEveryField(t *testing.T) {
	items := fixture(t)
	if len(items) != 5 {
		t.Fatalf("parsed %d items, want 5", len(items))
	}
	first := items[0]
	if first.Row != "pool" || first.Fit != 0.20 || first.Refused {
		t.Errorf("item 1 = %+v", first)
	}
	if first.Enough != "yes" || first.Better != "worse" || first.Slot != "yes" || first.Action != "none" {
		t.Errorf("item 1 fields = %+v", first)
	}
	if !items[3].Refused {
		t.Error("item 4 is a refused reword and was not read as one")
	}
	if items[1].Row != "locked" || items[1].Fit != 0.05 {
		t.Errorf("item 2 = %+v", items[1])
	}
}

// TestFirstWordKeepsFreeText is D-100. The owner may write a keyword, a
// dash, and then prose. The count reads the keyword and leaves the prose
// in the sheet.
func TestFirstWordKeepsFreeText(t *testing.T) {
	items := fixture(t)
	if got := items[1].Action; got != "reword" {
		t.Errorf("catalog_action = %q, want the keyword alone", got)
	}
	// The prose survives in the faults field, which is not counted.
	if !strings.Contains(items[0].Faults, "owned-only") {
		t.Errorf("the free text was lost: %q", items[0].Faults)
	}
	cases := map[string]string{
		"reword - presumes a list": "reword",
		"Add, no row covers this":  "add",
		"  NONE  ":                 "none",
		"better: much clearer":     "better",
		"":                         "",
	}
	for in, want := range cases {
		if got := firstWord(in); got != want {
			t.Errorf("firstWord(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestWarrantedRule is the D-66 definition, and the whole threshold rests
// on it.
func TestWarrantedRule(t *testing.T) {
	cases := []struct {
		enough, better string
		want, unsure   bool
	}{
		{"no", "better", true, false},
		{"no", "same", true, false},
		{"no", "worse", false, false},
		{"yes", "better", false, false},
		{"yes", "same", false, false},
		{"unsure", "better", false, true},
	}
	for _, tc := range cases {
		s := scored{Enough: tc.enough, Better: tc.better}
		if got := s.warranted(); got != tc.want {
			t.Errorf("warranted(%s,%s) = %v, want %v", tc.enough, tc.better, got, tc.want)
		}
		if got := s.unsure(); got != tc.unsure {
			t.Errorf("unsure(%s) = %v, want %v", tc.enough, got, tc.unsure)
		}
	}
}

// TestReportPicksTheHighestThresholdThatHolds walks the fixture. Two
// items sit at 0.05 and both are warranted. The 0.20 group adds an
// unwarranted one, which drops the precision under the floor.
func TestReportPicksTheHighestThresholdThatHolds(t *testing.T) {
	var buf bytes.Buffer
	report(&buf, fixture(t), PrecisionFloor)
	out := buf.String()
	if !strings.Contains(out, "Set the fit threshold to 0.06") {
		t.Errorf("the report did not pick 0.06:\n%s", out)
	}
	if !strings.Contains(out, "| 0.21 | 3 | 2 | 67% | no |") {
		t.Errorf("the 0.21 row is wrong:\n%s", out)
	}
	// The unsure item is left out of the fit (D-66).
	if !strings.Contains(out, "1 item is unsure") {
		t.Errorf("the unsure item was not reported:\n%s", out)
	}
	// A refused reword never reached a user, so it can not tell us what a
	// threshold would have allowed.
	if strings.Contains(out, "| 0.21 | 4 ") {
		t.Error("a refused reword was counted toward the fit threshold")
	}
	for _, want := range []string{"**add** (1 item): `house_rules`", "**reword** (1 item): `locked`"} {
		if !strings.Contains(out, want) {
			t.Errorf("the catalog actions are wrong, want %q:\n%s", want, out)
		}
	}
}

// TestReportSaysWhenNothingHolds covers the answer the owner most needs
// to see: every threshold lets through too much.
func TestReportSaysWhenNothingHolds(t *testing.T) {
	items := []scored{
		{Fit: 0.05, Enough: "yes", Better: "worse", Complete: true},
		{Fit: 0.20, Enough: "yes", Better: "same", Complete: true},
	}
	var buf bytes.Buffer
	report(&buf, items, PrecisionFloor)
	if !strings.Contains(buf.String(), "No candidate meets the floor") {
		t.Errorf("the report did not say the floor is unreachable:\n%s", buf.String())
	}
}

// TestReportCountsAnUnfinishedSheet keeps a half-scored sheet honest.
func TestReportCountsAnUnfinishedSheet(t *testing.T) {
	items := []scored{
		{Fit: 0.05, Enough: "no", Better: "better", Complete: true},
		{Fit: 0.20},
	}
	var buf bytes.Buffer
	report(&buf, items, PrecisionFloor)
	out := buf.String()
	if !strings.Contains(out, "Not scored yet: 1") {
		t.Errorf("the unscored item was not counted:\n%s", out)
	}
	if !strings.Contains(out, "The sheet is not finished") {
		t.Errorf("the report did not warn that the sheet is unfinished:\n%s", out)
	}
}

// TestReportSeparatesNoDataFromNoThreshold keeps two different answers
// apart. A sheet with no scored sent invention has nothing to say about
// the threshold, and saying "no candidate meets the floor" would read as
// a verdict on the catalog.
func TestReportSeparatesNoDataFromNoThreshold(t *testing.T) {
	// One scored item, and it is a refused reword.
	items := []scored{{Fit: 0.05, Refused: true, Enough: "yes", Better: "same", Complete: true}}
	var buf bytes.Buffer
	report(&buf, items, PrecisionFloor)
	out := buf.String()
	if strings.Contains(out, "No candidate meets the floor") {
		t.Errorf("a sheet with no sent invention read as a verdict on the catalog:\n%s", out)
	}
	if !strings.Contains(out, "No sent invention is scored yet") {
		t.Errorf("the report did not say the data is missing:\n%s", out)
	}
}
