package questions

import (
	"os"
	"regexp"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
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
// mtg-corpus skill, section 11. The skill is the document the owner reads.
// The JSON is the file the code reads. They must hold the same rows.
func TestCatalogMatchesCorpus(t *testing.T) {
	byName := map[string]string{
		"Out of scope": "out_of_scope",
		"Format":       "format", "Format (store event)": "format_store",
		"Theme or plan": "theme", "Theme (competitive)": "theme_competitive",
		"Theme (card named)": "theme_card_named", "Named card role": "named_card_role",
		"Commander": "commander", "Commander (pick)": "commander_pick",
		"Commander not owned": "commander_not_owned", "Weak commander pool": "commander_weak_pool",
		"Power (Commander)": "power_commander", "Power (60-card)": "power_sixty",
		"Colors": "colors", "Card pool": "pool", "Card pool (thin theme)": "pool_thin",
		"Budget": "budget", "Budget scope": "budget_scope",
		"House rules": "house_rules", "House format limits": "house_format_limits",
		"Jank or fun": "jank", "Meta": "meta",
		"One deck at a time": "one_deck", "Format (not supported)": "format_unsupported",
		"Format (no substitute)":       "format_unsupported_open",
		"Power (60-card, competitive)": "power_sixty_confirm", "Card pool (precon)": "pool_precon", "Commander (can not lead)": "commander_illegal",
		"Plan choice": "plan_choice", "Variance": "variance", "Locked cards": "locked",
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

// TestFormatFirst holds the ask order the dogfood runs asked for: nothing
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
// One cell held both questions before 2026-08-24, which made the no-repeat
// check unsafe.
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

// TestFrozenSessionAsksNothing is D-68.
func TestFrozenSessionAsksNothing(t *testing.T) {
	c := load(t)
	frozen := ctx(mtgv1.FormatId_FORMAT_ID_UNSPECIFIED)
	frozen.Frozen = true
	if got := c.Plan(frozen); len(got) != 0 {
		t.Fatalf("a frozen session planned %v", ids(got))
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
// Conversation 80 of gate run 20260826-220840-000 is the evidence. The
// user answered "TOURNAMENT", and the next turn asked "Should I build the
// deck for tournament-level competition?" Conversation 71 answered
// "FNM." and got the same row one turn later.
func TestPowerConfirmNeedsAnInferredStep(t *testing.T) {
	c := load(t)
	named := ctx(mtgv1.FormatId_FORMAT_ID_MODERN, "power")
	named.PowerCompetitive = true
	if got := ids(c.Plan(named)); has(got, "power_sixty_confirm") {
		t.Fatalf("the confirm row fired on a step the user named: %v", got)
	}
	inferred := ctx(mtgv1.FormatId_FORMAT_ID_MODERN, "power")
	inferred.PowerCompetitive, inferred.PowerInferred = true, true
	if got := ids(c.Plan(inferred)); !has(got, "power_sixty_confirm") {
		t.Fatalf("the confirm row stayed silent on an inferred step: %v", got)
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

// TestMetaRowAsksOneThing holds D-211. The row offers the general
// sideboard first, so one choice goes out. The old wording asked for a
// list of decks and for a yes-or-no answer in the same sentence, and the
// eval refused it in conversations 44 and 76 of gate run
// 20260826-220840-000.
func TestMetaRowAsksOneThing(t *testing.T) {
	c := load(t)
	r, ok := c.Row("meta")
	if !ok {
		t.Fatal("no meta row")
	}
	if !strings.HasPrefix(r.Text, "Should I keep the sideboard general") {
		t.Errorf("the meta row does not offer the general sideboard first: %q", r.Text)
	}
	if strings.Count(r.Text, "?") != 1 {
		t.Errorf("the meta row holds more than one question: %q", r.Text)
	}
	if len(r.Options) == 0 || r.Options[0] != "Keep the sideboard general" {
		t.Errorf("options = %v, want the general sideboard first", r.Options)
	}
}
