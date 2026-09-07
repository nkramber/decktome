package questions

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func load(t *testing.T) *Catalog {
	t.Helper()
	c, err := Load()
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	return c
}

func ctx(format mtgv1.FormatId, filled ...string) Context {
	c := Context{Format: format, Filled: map[string]bool{}, Asked: map[string]bool{}}
	for _, s := range filled {
		c.Filled[s] = true
	}
	return c
}

func ids(rows []Row) []string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, r.ID)
	}
	return out
}

func TestLoadValidates(t *testing.T) {
	c := load(t)
	if len(c.Rows) == 0 {
		t.Fatal("no rows")
	}
	for i := 1; i < len(c.Rows); i++ {
		if c.Rows[i-1].Order > c.Rows[i].Order {
			t.Fatalf("rows are not in ask order at %d", i)
		}
	}
}

// TestCatalogMatchesCorpus guards the drift between the data and the
// mtg-corpus skill, section 11. The skill is the document people read.
// The JSON is the file the code reads. They must hold the same rows.
func TestCatalogMatchesCorpus(t *testing.T) {
	byName := map[string]string{
		"Out of scope": "out_of_scope",
		"Format":       "format", "Format (store event)": "format_store",
		"Theme or plan": "theme", "Theme (competitive)": "theme_competitive",
		"Named card role": "named_card_role",
		"Commander":       "commander", "Commander (pick)": "commander_pick",
		"Power (Commander)": "power_commander", "Power (60-card)": "power_sixty",
		"Colors": "colors", "Card pool": "pool", "Card pool (thin theme)": "pool_thin",
		"Budget": "budget", "Budget scope": "budget_scope",
		"House rules": "house_rules", "House format limits": "house_format_limits",
		"One deck at a time": "one_deck", "Format (not supported)": "format_unsupported",
		"Format (no substitute)": "format_unsupported_open",
		"Card pool (precon)":     "pool_precon", "Commander (can not lead)": "commander_illegal",
		"Set (not resolved)": "set_unresolved", "Set (mana from outside)": "set_outside_mana",
		"Precon (not resolved)": "precon_unresolved",
	}
	raw, err := os.ReadFile("../../../.claude/skills/mtg-corpus/SKILL.md")
	if err != nil {
		t.Skipf("corpus skill not readable: %v", err)
	}
	text := string(raw)
	start := strings.Index(text, "## 11. Clarifying-question catalog")
	end := strings.Index(text, "## 12. Validation checklist")
	if start < 0 || end < 0 {
		t.Fatal("corpus section 11 not found")
	}
	row := regexp.MustCompile(`(?m)^\| ([^|]+) \|`)
	var names []string
	for _, m := range row.FindAllStringSubmatch(text[start:end], -1) {
		name := strings.TrimSpace(m[1])
		if name == "Slot" || strings.HasPrefix(name, "---") {
			continue
		}
		names = append(names, name)
	}
	c := load(t)
	if len(names) != len(c.Rows) {
		t.Errorf("corpus holds %d rows, catalog.json holds %d", len(names), len(c.Rows))
	}
	for _, name := range names {
		id, ok := byName[name]
		if !ok {
			t.Errorf("corpus row %q has no catalog row. Add it to catalog.json and to this map", name)
			continue
		}
		if _, ok := c.Row(id); !ok {
			t.Errorf("corpus row %q maps to id %q, which catalog.json does not hold", name, id)
		}
	}
}

// TestFormatFirst holds the ask order of corpus section 11: nothing
// before the format, because every other row depends on it.
func TestFormatFirst(t *testing.T) {
	c := load(t)
	got := ids(c.Plan(Context{Filled: map[string]bool{}, Asked: map[string]bool{}}))
	if len(got) == 0 || got[0] != "format" {
		t.Fatalf("first question = %v, want format first", got)
	}
	for _, id := range got {
		if id == "power_commander" || id == "power_sixty" {
			t.Errorf("power asked before the format is known: %v", got)
		}
	}
}

// TestPoolWaitsForFormatColorsTheme is D-67. PR-6 needs the format, the
// colors, and the theme before it can count on-theme owned cards, so the
// card-pool question can not come first.
func TestPoolWaitsForFormatColorsTheme(t *testing.T) {
	c := load(t)
	early := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format")
	early.HasCollection = true
	for _, id := range ids(c.Plan(early)) {
		if id == "pool" || id == "pool_thin" {
			t.Fatalf("pool asked before colors and theme: %v", ids(c.Plan(early)))
		}
	}
	ready := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format", "colors", "theme")
	ready.HasCollection = true
	if !contains(ids(c.Plan(ready)), "pool") {
		t.Fatalf("pool not asked after format, colors, and theme: %v", ids(c.Plan(ready)))
	}
}

// TestThinThemeReplacesPool checks that the two pool rows never both fire.
// One cell that holds both questions makes the no-repeat check unsafe.
func TestThinThemeReplacesPool(t *testing.T) {
	c := load(t)
	thin := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format", "colors", "theme")
	thin.HasCollection, thin.ThinTheme = true, true
	got := ids(c.Plan(thin))
	if contains(got, "pool") {
		t.Errorf("the generic pool row fired although the theme is thin: %v", got)
	}
	if !contains(got, "pool_thin") {
		t.Errorf("the thin-theme row did not fire: %v", got)
	}

	fat := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format", "colors", "theme")
	fat.HasCollection = true
	got = ids(c.Plan(fat))
	if !contains(got, "pool") || contains(got, "pool_thin") {
		t.Errorf("with a rich theme, want the generic pool row only: %v", got)
	}
}

// TestNoCollectionNeverAsksPool is D-37.
func TestNoCollectionNeverAsksPool(t *testing.T) {
	c := load(t)
	none := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format", "colors", "theme")
	none.HasCollection = false
	for _, id := range ids(c.Plan(none)) {
		if strings.HasPrefix(id, "pool") {
			t.Fatalf("pool asked with no collection: %v", id)
		}
	}
}

// TestColorsSkippedWhenCommanderSet keeps the agent from spending an ask
// on a slot the commander already fills.
func TestColorsSkippedWhenCommanderSet(t *testing.T) {
	c := load(t)
	with := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format", "theme")
	with.CommanderSet = true
	if contains(ids(c.Plan(with)), "colors") {
		t.Fatalf("colors asked although a commander is set: %v", ids(c.Plan(with)))
	}
}

// TestNoRepeat is the PR-7 gate rule.
func TestNoRepeat(t *testing.T) {
	c := load(t)
	first := ctx(mtgv1.FormatId_FORMAT_ID_UNSPECIFIED)
	asked := c.Plan(first)
	if len(asked) == 0 {
		t.Fatal("no first question")
	}
	for _, r := range asked {
		first.Asked[r.ID] = true
	}
	for _, r := range c.Plan(first) {
		if first.Asked[r.ID] {
			t.Fatalf("row %q asked twice", r.ID)
		}
	}
}

// TestMaxThreePerTurn is the corpus rule.
func TestMaxThreePerTurn(t *testing.T) {
	c := load(t)
	full := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	full.HasCollection, full.NamedCard, full.PowerCompetitive = true, true, true
	if n := len(c.Plan(full)); n > MaxPerTurn {
		t.Fatalf("planned %d questions in one turn, max is %d", n, MaxPerTurn)
	}
}

// TestOneRowPerSlotPerTurn stops the agent from asking two commander
// questions in one turn.
func TestOneRowPerSlotPerTurn(t *testing.T) {
	c := load(t)
	both := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format")
	both.NamedCard, both.Suggested = true, true
	seen := map[string]bool{}
	for _, r := range c.Plan(both) {
		if seen[r.Slot] {
			t.Fatalf("two rows for slot %q in one turn", r.Slot)
		}
		seen[r.Slot] = true
	}
}

func TestRoute(t *testing.T) {
	cases := map[string]string{
		"60-card anything goes":    "house_rules",
		"we play with no ban list": "house_rules",
		// D-111. A proxy user has no budget, and the word says nothing
		// about which cards are legal.
		"I will proxy the expensive cards": "",
		// D-111. A negation stops a trigger word.
		"edh gruul dino stompy pls, no proxies": "",
		"make me something fun and janky":       "power",
		"build the strongest deck possible":     "power",
		"I want to win the event":               "power",
		"build me a lifegain deck":              "",
	}
	for in, want := range cases {
		if got := Route(in); got != want {
			t.Errorf("Route(%q) = %q, want %q", in, got, want)
		}
	}
}

func contains(list []string, want string) bool {
	for _, s := range list {
		if s == want {
			return true
		}
	}
	return false
}

// TestPowerConfirmNeedsAnInferredStep holds the D-209 rule. The confirm
// row asks about a step the agent filled in. A step the user named needs
// no confirmation, and a question about it repeats the answer.
//
// A user who answers "TOURNAMENT" must not get "Should I build the deck
// for tournament-level competition?" one turn later. D-216 then removed
// the row: a user who asks for the strongest deck has already given the
// answer.
func TestNoRowConfirmsAnInferredStep(t *testing.T) {
	c := load(t)
	for _, tc := range []struct {
		name     string
		inferred bool
	}{{"a step the user named", false}, {"a step the agent inferred", true}} {
		got := ctx(mtgv1.FormatId_FORMAT_ID_MODERN, "power")
		got.PowerCompetitive = true
		_ = tc.inferred
		if rows := ids(c.Plan(got)); has(rows, "power_sixty_confirm") {
			t.Fatalf("%s: a confirm row went out: %v", tc.name, rows)
		}
	}
	if _, ok := c.Row("power_sixty_confirm"); ok {
		t.Fatal("the confirm row is still in the catalog (D-216)")
	}
}

func has(list []string, want string) bool {
	for _, s := range list {
		if s == want {
			return true
		}
	}
	return false
}

// TestHouseLimitsRowNamesNoList holds D-212. The row asks one yes-or-no
// question. A list of limits inside it reads as one question for each
// item, even when the row goes out word for word.
func TestHouseLimitsRowNamesNoList(t *testing.T) {
	c := load(t)
	r, ok := c.Row("house_format_limits")
	if !ok {
		t.Fatal("no house_format_limits row")
	}
	if !r.Fixed {
		t.Error("the house-limits row lost its fixed flag (D-162)")
	}
	if strings.Contains(r.Text, ":") {
		t.Errorf("the house-limits row names a list of limits: %q", r.Text)
	}
	if strings.Count(r.Text, "?") != 1 {
		t.Errorf("the house-limits row holds more than one question: %q", r.Text)
	}
}

// TestStoreFormatRowNamesItsOwnLimit holds D-213. The three formats are
// the ones this app builds. An FNM can run Pioneer, so a question that
// offers the list as the event's own formats omits information.
func TestStoreFormatRowNamesItsOwnLimit(t *testing.T) {
	c := load(t)
	r, ok := c.Row("format_store")
	if !ok {
		t.Fatal("no format_store row")
	}
	if !strings.HasPrefix(r.Text, "I build Standard, Modern, and Commander.") {
		t.Errorf("the store row does not name the formats it builds: %q", r.Text)
	}
	if strings.Contains(r.Text, "run:") {
		t.Errorf("the store row offers the list as the event's formats: %q", r.Text)
	}
	if f := LintCatalog(c); len(f) > 0 {
		t.Errorf("the linter refused the catalog: %v", f)
	}
}

// TestLoadRefusesABadTrigger holds the catalog checks: a format trigger
// outside the two values, a key that is not a word, and a wait on a key
// no row owns.
func TestLoadRefusesABadTrigger(t *testing.T) {
	const good = `{"verified_at":"2026-08-28","rows":[
		{"id":"format","slot":"format","order":1,"text":"Which format?"},
		{"id":"theme","slot":"theme","order":2,"text":"Which theme?"%s}]}`
	cases := []struct {
		name, extra string
		wantErr     bool
	}{
		{"a clean catalog", "", false},
		{"a known format trigger", `,"when":{"format":"sixty"}`, false},
		{"an unknown format trigger", `,"when":{"format":"pauper"}`, true},
		{"a key that is a word", `,"key":"theme_plan"`, false},
		{"a key with a space", `,"key":"theme plan"`, true},
		{"a wait on an owned key", `,"when":{"not_outstanding":["format"]}`, false},
		{"a wait on a key nobody owns", `,"when":{"not_outstanding":["sideboard"]}`, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parse([]byte(fmt.Sprintf(good, tc.extra)))
			if (err != nil) != tc.wantErr {
				t.Errorf("parse error = %v, want error %v", err, tc.wantErr)
			}
		})
	}
}
