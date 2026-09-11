package questions

import (
	"context"
	"testing"
)

// TestAnOccasionDoesNotFillTheTheme walks D-670 through the agent. "I need
// a Modern deck for a team event on Saturday." names a happening, and the
// classifier writes it as the theme. The theme stays empty, so the first
// turn asks for the plan.
func TestAnOccasionDoesNotFillTheTheme(t *testing.T) {
	var out classifyOut
	out.Format, out.Theme, out.PoolRule = "modern", "team event", "any_card"
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "theme", "power_sixty", "colors", "budget"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I need a Modern deck for a team event on Saturday.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got := st.Slots.GetTheme(); got != "" {
		t.Errorf("theme = %q, want empty: an occasion is not a theme", got)
	}
	if question(res.Questions, "theme") == nil {
		t.Error("the first turn did not ask for the theme")
	}
}
