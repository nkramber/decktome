package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixture = `# PR-7 question gate

## Conversations

### 1. lifegain with a collection

**Turn 1, the user:** build me a lifegain deck

- [catalog slot=format row=format fit=0.90 filled=true] Which format?
- [catalog slot=power row=power_commander fit=0.20 filled=true] Which bracket does your table play?
  - Refused as a reword (D-88): Which bracket does your white-black table play?
- [INVENTED slot=colors row=colors fit=0.20 filled=true] Which colors would you like?
  - It replaced: Any color preference?

**Turn 2, the user:** commander

- [INVENTED slot=budget row=acquisition fit=0.05 filled=false] What is your budget?
  - It replaced: Do you buy in person or online?

### 2. blink

**Turn 1, the user:** blink deck

- [INVENTED slot=budget row=budget_scope fit=0.05 filled=false] Is that a cap on the cards you buy?
`

func writeFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "pr7-question-gate-runX.md")
	if err := os.WriteFile(path, []byte(fixture), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCollectPairsBothTexts(t *testing.T) {
	items, err := collect([]string{writeFixture(t)})
	if err != nil {
		t.Fatalf("collect: %v", err)
	}
	// Three inventions, one refused reword, and one plain catalog
	// question that the model left alone.
	if len(items) != 5 {
		t.Fatalf("collected %d questions, want 5", len(items))
	}
	refused := items[1]
	if !refused.Refused || refused.Row != "power_commander" {
		t.Fatalf("the refused reword was not collected: %+v", refused)
	}
	if refused.Catalog != "Which bracket does your table play?" ||
		refused.Invented != "Which bracket does your white-black table play?" {
		t.Errorf("the refused pair is wrong: %+v", refused)
	}
	first := items[2]
	if first.Row != "colors" || first.Slot != "colors" || first.Fit != 0.20 || !first.Filled {
		t.Errorf("first item = %+v", first)
	}
	if first.Catalog != "Any color preference?" {
		t.Errorf("the catalog text did not attach: %q", first.Catalog)
	}
	if first.Conv != "1. lifegain with a collection" || first.Turn != 1 {
		t.Errorf("the conversation or turn is wrong: %+v", first)
	}
	if items[3].Turn != 2 {
		t.Errorf("turn 2 was read as turn %d", items[3].Turn)
	}
	// A question with no "It replaced" line must not borrow the previous
	// one's catalog text.
	if items[4].Catalog != "" {
		t.Errorf("an unpaired question borrowed a catalog text: %q", items[4].Catalog)
	}
}

// TestScorableDropsWhatTheOwnerCanNotJudge is the D-66 rule for a
// replacement: the rubric compares two questions, so a replacement with
// one text is not scorable. A row the catalog no longer holds is not
// worth a scorer's time either.
//
// A question the model left alone is now kept. It becomes a plain item,
// and version 1 of the sheet could not hold one (D-104).
func TestScorableDropsWhatTheOwnerCanNotJudge(t *testing.T) {
	items, err := collect([]string{writeFixture(t)})
	if err != nil {
		t.Fatalf("collect: %v", err)
	}
	kept, dropped, err := scorable(items)
	if err != nil {
		t.Fatalf("scorable: %v", err)
	}
	reps, plain := 0, 0
	rows := map[string]bool{}
	for _, it := range kept {
		rows[it.Row] = true
		if it.replacement() {
			reps++
		} else {
			plain++
		}
	}
	if reps != 2 {
		t.Errorf("kept %d replacements, want the colors invention and the refused reword: %+v", reps, kept)
	}
	if plain == 0 {
		t.Error("no plain item survived, and D-104 makes every question scorable")
	}
	if !rows["colors"] || !rows["power_commander"] {
		t.Errorf("kept rows = %v, want colors and power_commander", rows)
	}
	if dropped["no catalog text recorded"] != 1 {
		t.Errorf("dropped map = %v, want one item with no catalog text", dropped)
	}
	if dropped["the catalog no longer holds the row"] != 1 {
		t.Errorf("dropped map = %v, want one deleted row (acquisition, D-87)", dropped)
	}
	// A catalog question the model left alone is kept, not dropped.
	if _, ok := dropped["the model offered no replacement"]; ok {
		t.Errorf("dropped map = %v, want no entry for an untouched catalog question", dropped)
	}
}

// TestEveryTenthItemRepeats is the self-consistency check of D-66. The
// repeat is inserted, never written over a real item.
func TestEveryTenthItemRepeats(t *testing.T) {
	var items []item
	for i := 0; i < 12; i++ {
		items = append(items, item{Row: "row" + string(rune('a'+i)), Invented: "q", Catalog: "c"})
	}
	got := pick(items, 0)
	if len(got) != 13 {
		t.Fatalf("picked %d, want 12 real items and 1 repeat", len(got))
	}
	if got[9].Row != got[0].Row {
		t.Errorf("item 10 is %q, want a repeat of item 1 (%q)", got[9].Row, got[0].Row)
	}
	// Every real item survives the insertion.
	seen := map[string]int{}
	for _, it := range got {
		seen[it.Row]++
	}
	for _, it := range items {
		if seen[it.Row] == 0 {
			t.Errorf("item %q was lost to the repeat", it.Row)
		}
	}
}

func TestWriteHoldsEveryField(t *testing.T) {
	var buf bytes.Buffer
	write(&buf, []item{{
		Run: "runX", Conv: "1. lifegain", Turn: 1, Row: "colors", Slot: "colors",
		Fit: 0.2, Filled: true, Invented: "Which colors?", Catalog: "Any color preference?",
	}, {
		Run: "runX", Conv: "1. lifegain", Turn: 2, Row: "power_commander", Slot: "power",
		Fit: 0.2, Invented: "Which bracket for your table?", Catalog: "Which bracket?", Refused: true,
	}}, 5, map[string]int{"no catalog text recorded": 2})
	out := buf.String()
	for _, want := range []string{
		"catalog_enough", "invented_better", "right_slot", "filled_slot", "faults", "catalog_action",
		"warranted", "inaccurate", "omits information", "n/a",
		"Any color preference?", "Which colors?", "| filled_slot | the run | yes |",
		"| catalog_enough | the catalog question | |",
		"| right_slot | the replacement | |",
		"2, because no catalog text recorded.",
		"refused this one as a reword",
		"1 were refused as rewords",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("the sheet does not hold %q", want)
		}
	}
}

// TestPlainItemAsksWhetherItWasWarranted covers the item kind D-104 adds.
// Version 1 could hold only a question the model offered to replace, and
// that is where most defects hid.
func TestPlainItemAsksWhetherItWasWarranted(t *testing.T) {
	var buf bytes.Buffer
	write(&buf, []item{{
		Run: "runX", Conv: "23. land destruction", Turn: 1, Row: "format", Slot: "format",
		Fit: 0.9, Filled: true,
		Asked:   "What format would you like: Commander, Standard, Modern, Pioneer, or Pauper?",
		Catalog: "What format would you like: Commander, Standard, Modern, Pioneer, or Pauper?",
	}}, 1, nil)
	out := buf.String()
	for _, want := range []string{
		"**The question that went out:**",
		"The model offered no replacement here.",
		"| warranted | the question | |",
		"| filled_slot | the run | yes |",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("the sheet does not hold %q", want)
		}
	}
	// A plain item must not carry the fields that compare two questions.
	if strings.Contains(out, "| catalog_enough | the catalog question | |") {
		t.Error("a plain item carries catalog_enough, which reads a question that does not exist")
	}
}

// TestPickKeepsEveryReplacementAndReachesEveryRow is the sampling rule of
// D-104. A replacement decides the D-27 threshold, so none may be lost. A
// rare row must still reach the sheet.
func TestPickKeepsEveryReplacementAndReachesEveryRow(t *testing.T) {
	var items []item
	for i := 0; i < 3; i++ {
		items = append(items, item{Row: "colors", Slot: "colors", Asked: "q", Catalog: "c", Invented: "r"})
	}
	// One noisy row, and two rows with a single question each.
	for i := 0; i < 200; i++ {
		items = append(items, item{Row: "format", Slot: "format", Asked: "q", Catalog: "q"})
	}
	items = append(items,
		item{Row: "locked", Slot: "locked", Asked: "q", Catalog: "q"},
		item{Row: "meta", Slot: "meta", Asked: "q", Catalog: "q"})

	got := pick(items, 20)
	if len(got) != 20 {
		t.Fatalf("picked %d items, want 20", len(got))
	}
	reps, rows := 0, map[string]bool{}
	for _, it := range got {
		rows[it.Row] = true
		if it.replacement() {
			reps++
		}
	}
	if reps < 3 {
		t.Errorf("%d replacements reached the sheet, want all 3", reps)
	}
	for _, row := range []string{"colors", "format", "locked", "meta"} {
		if !rows[row] {
			t.Errorf("row %q never reached the sheet", row)
		}
	}
}
