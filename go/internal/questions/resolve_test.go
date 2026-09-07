package questions

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
)

func themedState(theme string) *State {
	st := NewState(true)
	st.Slots.Theme = theme
	return st
}

// TestNoPlaceholderReachesTheModel is D-82. A brace such as "{theme} is
// strongest in {colors}" that reaches the ask role comes back as a
// second question aimed at the user.
func TestNoPlaceholderReachesTheModel(t *testing.T) {
	c := load(t)
	for _, r := range c.Rows {
		for _, h := range []Hints{nil, &fakeHints{}} {
			text, _, _ := resolve(r, themedState("lifegain"), h)
			if strings.ContainsAny(text, "{}") {
				t.Errorf("row %q resolved to %q, which still holds a placeholder", r.ID, text)
			}
			if !usable(text) {
				t.Errorf("row %q resolved to %q, which can not stand as a question", r.ID, text)
			}
		}
	}
}

// colorsStubRow is the row the catalog held before D-108. No catalog row
// names {colors} now, and no hint answers it, so the clause drops.
func colorsStubRow() Row {
	return Row{ID: "colors_stub", Slot: "colors",
		Text:     "Any color preference? A {theme} deck is strongest in {colors}.",
		Fallback: "Any color preference?"}
}

// TestColorsRowStatesNoFact is D-108. A row that asserts which colors a
// theme is strongest in can be wrong: "Black Lotus decks are strongest
// in black", and Black Lotus is a colorless card.
func TestColorsRowStatesNoFact(t *testing.T) {
	c := load(t)
	row, ok := c.Row("colors")
	if !ok {
		t.Fatal("no colors row")
	}
	if placeholder.MatchString(row.Text) {
		t.Errorf("the colors row holds a placeholder again: %q", row.Text)
	}
	for _, word := range []string{"strongest", "best in", "strong in"} {
		if strings.Contains(strings.ToLower(row.Text), word) {
			t.Errorf("the colors row states a fact again: %q", row.Text)
		}
	}
	got, _, _ := resolve(row, themedState("lifegain"), &fakeHints{})
	if got != "Any color preference?" {
		t.Errorf("resolve = %q, want the question alone", got)
	}
}

// TestResolveDropsTheClauseWithNoValue keeps the short, correct question.
func TestResolveDropsTheClauseWithNoValue(t *testing.T) {
	row := colorsStubRow()
	got, _, _ := resolve(row, themedState("lifegain"), nil)
	if got != "Any color preference?" {
		t.Errorf("resolve = %q, want the first sentence alone", got)
	}
}

// TestResolveFallsBackWhenNothingSurvives covers a row whose first clause
// carries the meaning. The illegal-commander row opens with the card
// name, and a state with no such card leaves nothing to fill.
func TestResolveFallsBackWhenNothingSurvives(t *testing.T) {
	c := load(t)
	row, _ := c.Row("commander_illegal")
	got, _, _ := resolve(row, themedState(""), nil)
	if got != row.Fallback {
		t.Errorf("resolve = %q, want the fallback %q", got, row.Fallback)
	}
}

func TestResolveNamedCard(t *testing.T) {
	c := load(t)
	row, _ := c.Row("named_card_role")
	st := themedState("sacrifice")
	st.NamedCards = []string{"Grist, the Hunger Tide"}
	got, _, _ := resolve(row, st, nil)
	if !strings.Contains(got, "Grist, the Hunger Tide") {
		t.Errorf("resolve = %q, want the card name", got)
	}
}

// TestGuard is the check on the model's phrasing.
func TestGuard(t *testing.T) {
	resolved := "Any color preference?"
	cases := []struct {
		name, phrased, want string
	}{
		{"good phrasing passes", "Do you have a color preference?", "Do you have a color preference?"},
		{"compound question falls back", "Do you have a color preference? Lifegain is strongest in which colors?", resolved},
		{"a brace falls back", "Any color preference? {theme} is strongest?", resolved},
		{"empty falls back", "   ", resolved},
		{"no question mark falls back", "Tell me your colors.", resolved},
		{"a long ramble falls back", strings.Repeat("word ", 200) + "?", resolved},
		// D-150. "What should the Oathbreaker deck focus on?" one line
		// under "I do not build Oathbreaker" is a contradiction.
		{"an unsupported format falls back",
			"What should the Oathbreaker deck focus on: a creature type, a mechanic, or a play style?", resolved},
		{"Historic falls back too",
			"What should the Historic deck focus on: a creature type or a play style?", resolved},
		{"a supported format passes",
			"What should the Modern deck focus on: a creature type or a play style?",
			"What should the Modern deck focus on: a creature type or a play style?"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := guard("theme", tc.phrased, resolved); got != tc.want {
				t.Errorf("guard = %q, want %q", got, tc.want)
			}
		})
	}
}

// TestGuardLetsTheDecliningRowNameTheFormat is D-150. The row that
// declines an unsupported format must name it, which is the whole job of
// that row. Those rows carry `fixed` and never reach the ask role today,
// so this is a guard against a later change (D-112, D-117).
func TestGuardLetsTheDecliningRowNameTheFormat(t *testing.T) {
	phrased := "I do not build Oathbreaker. Shall I use Commander instead?"
	resolved := "I do not build Oathbreaker. The nearest format I build is Commander. Shall I use that?"
	for _, row := range []string{"format_unsupported", "format_unsupported_open"} {
		if got := guard(row, phrased, resolved); got != phrased {
			t.Errorf("%s: guard refused the row that must name the format: %q", row, got)
		}
	}
}

// TestClosedKeysMustBeOffered is the second live defect: the classifier
// named keys it was never offered, and the session ended with three slots
// still empty.
func TestClosedKeysMustBeOffered(t *testing.T) {
	var out classifyOut
	out.Format, out.Theme, out.PoolRule = "commander", "lifegain", "unknown"
	// The classifier claims the commander, power, and pool slots are done.
	out.ClosedKeys = []string{"commander", "power", "pool_rule"}
	a, _ := testAgent(t,
		classifyStep(t, out),
		fits(t, "colors"),
		askStep(t, phrasing{RowID: "colors", Text: "Colors?"}))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "build me a lifegain deck", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	for _, key := range []string{"commander", "power", "pool_rule"} {
		if st.Ctx.Filled[key] {
			t.Errorf("key %q closed although the agent never offered it", key)
		}
		if st.Slots.SlotStates[key] == mtgv1.SlotState_SLOT_STATE_FILLED {
			t.Errorf("slot %q marked filled although the agent never offered it", key)
		}
	}
	if st.Ready(load(t)) {
		t.Error("the session called itself complete with three slots open")
	}
}

// TestPhrasedBraceNeverReachesTheUser is the end-to-end version of the
// live defect: the model returns a mangled question and the guard drops it.
func TestPhrasedBraceNeverReachesTheUser(t *testing.T) {
	var out classifyOut
	out.Format, out.Theme, out.PoolRule = "commander", "lifegain", "any_card"
	a, _ := testAgent(t,
		classifyStep(t, out),
		fits(t, "colors"),
		askStep(t, phrasing{RowID: "colors", Text: "Any color preference? Lifegain is strongest in which colors?"}))
	res, err := a.Turn(context.Background(), NewState(false), "lifegain deck", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	for _, q := range res.Questions {
		if strings.Count(q.Text, "?") != 1 {
			t.Errorf("question %q asks more than one thing", q.Text)
		}
	}
}

// TestHintsNilSafety keeps a missing index from breaking a turn.
func TestHintsNilSafety(t *testing.T) {
	var h *CandidateHints
	if h.Commanders("lifegain", nil) != nil || h.OwnedThemeCount("lifegain") != 0 {
		t.Error("a nil hint source answered something")
	}
	empty := &CandidateHints{}
	if empty.Commanders("lifegain", nil) != nil {
		t.Error("an empty hint source answered something")
	}
}

var _ = llm.RoleAsk

// TestOrListJoinsTheChoices is D-518: a question that offers choices
// joins them with "or", and a list of things a deck holds keeps "and".
func TestOrListJoinsTheChoices(t *testing.T) {
	if got := orList([]string{"Marvel Super Heroes", "Marvel Universe"}); got != "Marvel Super Heroes or Marvel Universe" {
		t.Errorf("orList of two = %q", got)
	}
	if got := orList([]string{"Khans of Tarkir", "Dragons of Tarkir", "Tarkir: Dragonstorm"}); got != "Khans of Tarkir, Dragons of Tarkir, or Tarkir: Dragonstorm" {
		t.Errorf("orList of three = %q", got)
	}
	if got := englishList([]string{"Sol Ring", "Sanguine Bond"}); got != "Sol Ring and Sanguine Bond" {
		t.Errorf("englishList of two = %q", got)
	}
	if got := orList([]string{" ", "one"}); got != "one" {
		t.Errorf("orList drops blanks: %q", got)
	}
}
