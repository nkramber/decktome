package questions

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

type stubHints struct {
	colors     string
	commanders []string
	owned      int
}

func (s stubHints) ThemeColors(string) string  { return s.colors }
func (s stubHints) Commanders(string) []string { return s.commanders }
func (s stubHints) OwnedThemeCount(string) int { return s.owned }

func themedState(theme string) *State {
	st := NewState(true)
	st.Slots.Theme = theme
	return st
}

// TestNoPlaceholderReachesTheModel is the rule the live run of 2026-08-24
// broke: the agent sent "{theme} is strongest in {colors}" to the ask role,
// and the model turned the clause into a second question for the user.
func TestNoPlaceholderReachesTheModel(t *testing.T) {
	c := load(t)
	for _, r := range c.Rows {
		for _, h := range []Hints{nil, stubHints{}} {
			text, _ := resolve(r, themedState("lifegain"), h)
			if strings.ContainsAny(text, "{}") {
				t.Errorf("row %q resolved to %q, which still holds a placeholder", r.ID, text)
			}
			if !usable(text) {
				t.Errorf("row %q resolved to %q, which can not stand as a question", r.ID, text)
			}
		}
	}
}

func TestResolveUsesHints(t *testing.T) {
	c := load(t)
	row, ok := c.Row("colors")
	if !ok {
		t.Fatal("no colors row")
	}
	got, _ := resolve(row, themedState("lifegain"), stubHints{colors: "white and black"})
	want := "Any color preference? lifegain is strongest in white and black."
	if got != want {
		t.Errorf("resolve = %q, want %q", got, want)
	}
}

// TestResolveDropsTheClauseWithNoValue keeps the short, correct question.
func TestResolveDropsTheClauseWithNoValue(t *testing.T) {
	c := load(t)
	row, _ := c.Row("colors")
	got, _ := resolve(row, themedState("lifegain"), nil)
	if got != "Any color preference?" {
		t.Errorf("resolve = %q, want the first sentence alone", got)
	}
}

// TestResolveFallsBackWhenNothingSurvives covers a row whose first clause
// carries the meaning.
func TestResolveFallsBackWhenNothingSurvives(t *testing.T) {
	c := load(t)
	row, _ := c.Row("theme_card_named")
	got, _ := resolve(row, themedState(""), nil)
	if got != row.Fallback {
		t.Errorf("resolve = %q, want the fallback %q", got, row.Fallback)
	}
}

func TestResolveNamedCard(t *testing.T) {
	c := load(t)
	row, _ := c.Row("named_card_role")
	st := themedState("sacrifice")
	st.NamedCards = []string{"Grist, the Hunger Tide"}
	got, _ := resolve(row, st, nil)
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
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := guard(tc.phrased, resolved); got != tc.want {
				t.Errorf("guard = %q, want %q", got, tc.want)
			}
		})
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
	if h.ThemeColors("lifegain") != "" || h.Commanders("lifegain") != nil || h.OwnedThemeCount("lifegain") != 0 {
		t.Error("a nil hint source answered something")
	}
	empty := &CandidateHints{}
	if empty.ThemeColors("lifegain") != "" {
		t.Error("an empty hint source answered something")
	}
}

func TestColorWords(t *testing.T) {
	cases := []struct {
		in   []mtgv1.Color
		want string
	}{
		{nil, ""},
		{[]mtgv1.Color{mtgv1.Color_COLOR_W}, "white"},
		{[]mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}, "white and black"},
		{[]mtgv1.Color{mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_R}, "blue, black, and red"},
	}
	for _, tc := range cases {
		if got := colorWords(tc.in); got != tc.want {
			t.Errorf("colorWords(%v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

var _ = llm.RoleAsk
