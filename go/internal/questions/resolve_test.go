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

func (s stubHints) ThemeColors(string) string { return s.colors }

// Commanders answers the names the agent has not offered yet, which is
// what the real hint source does (D-73).
func (s stubHints) Commanders(_ string, skip []string) []string {
	var out []string
	for _, name := range s.commanders {
		if hasName(skip, name) {
			continue
		}
		out = append(out, name)
		if len(out) == 3 {
			break
		}
	}
	return out
}
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
// names {colors} now, and the placeholder path still runs for PR-8.
func colorsStubRow() Row {
	return Row{ID: "colors_stub", Slot: "colors",
		Text:     "Any color preference? A {theme} deck is strongest in {colors}.",
		Fallback: "Any color preference?"}
}

// TestColorsRowStatesNoFact is D-108. The row asserted which colors a
// theme is strongest in. Gate run 13 of 2026-08-25 read "Extra-turns
// decks are strongest in blue and green", and "Black Lotus decks are
// strongest in black". Black Lotus is a colorless card.
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
	got, _, _ := resolve(row, themedState("lifegain"), stubHints{colors: "white and black"})
	if got != "Any color preference?" {
		t.Errorf("resolve = %q, want the question alone", got)
	}
}

func TestResolveUsesHints(t *testing.T) {
	row := colorsStubRow()
	got, _, _ := resolve(row, themedState("lifegain"), stubHints{colors: "white and black"})
	want := "Any color preference? A lifegain deck is strongest in white and black."
	if got != want {
		t.Errorf("resolve = %q, want %q", got, want)
	}
}

// TestVagueThemeDropsTheColorClause is the D-79 side effect that gate
// run 3 of 2026-08-25 found. "The best deck under budget" is a theme
// value, and it reads as nonsense inside a statement about colors. The
// model replaced three color questions for that reason.
func TestVagueThemeDropsTheColorClause(t *testing.T) {
	row := colorsStubRow()
	for _, theme := range []string{"the best deck under budget", "the strongest Modern deck", "a named tier-one deck"} {
		got, _, _ := resolve(row, themedState(theme), stubHints{colors: "white and black"})
		if got != "Any color preference?" {
			t.Errorf("theme %q gave %q, want the short question", theme, got)
		}
	}
	// A real archetype keeps the clause.
	got, _, _ := resolve(row, themedState("lifegain"), stubHints{colors: "white and black"})
	if !strings.Contains(got, "strongest in white and black") {
		t.Errorf("a real theme lost its color clause: %q", got)
	}
}

// TestTooManyColorsDropsTheClause keeps a useless statement out. Four
// colors name no preference at all.
func TestTooManyColorsDropsTheClause(t *testing.T) {
	row := colorsStubRow()
	got, _, _ := resolve(row, themedState("lifegain"), stubHints{colors: "white, blue, black, and green"})
	if got != "Any color preference?" {
		t.Errorf("four colors gave %q, want the short question", got)
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
// carries the meaning.
func TestResolveFallsBackWhenNothingSurvives(t *testing.T) {
	c := load(t)
	row, _ := c.Row("theme_card_named")
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
	if h.ThemeColors("lifegain") != "" || h.Commanders("lifegain", nil) != nil || h.OwnedThemeCount("lifegain") != 0 {
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
