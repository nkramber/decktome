package questions

import (
	"context"
	"testing"
)

// TestStrongestWords is D-1153. A superlative fills bracket 5 in
// Commander. "Competitive" and a budget phrase name no bracket, so they
// fill only the tournament step of a 60-card format (D-216).
func TestStrongestWords(t *testing.T) {
	for _, m := range []string{
		"Marvel set only, best possible deck",
		"the strongest deck you can make",
		"a top tier Commander deck",
	} {
		if !strongestRequest(m) {
			t.Errorf("%q is a request for the strongest deck, and it did not fire", m)
		}
		if !competitiveRequest(m) {
			t.Errorf("%q fills the tournament step in a 60-card format, and it did not fire", m)
		}
	}
	for _, m := range []string{
		"a competitive Commander deck",
		"money is no object",
		"a serious deck",
		"not the strongest deck, just a fun one",
	} {
		if strongestRequest(m) {
			t.Errorf("%q names no bracket, and it fired", m)
		}
	}
}

// TestStrongestCommanderFillsBracketFive is D-1153. Session
// wBrsxouAndrjDXEJ8dDw asked for "Marvel set only, best possible deck",
// and the agent asked the bracket. The words now fill bracket 5 with no
// power question.
func TestStrongestCommanderFillsBracketFive(t *testing.T) {
	out := classifyOut{Format: "commander", Theme: "marvel", PoolRule: "any_card", BudgetUSD: 500}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "A Commander deck. Marvel set only, best possible deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got := st.Slots.GetPower().GetBracket(); got != cedhBracket {
		t.Errorf("power bracket = %d, want %d", got, cedhBracket)
	}
	if !st.Ctx.Filled["power"] {
		t.Error("the power slot did not close")
	}
	if q := question(res.Questions, "power"); q != nil {
		t.Errorf("the bracket row asked a reader who asked for the best deck: %q", q.GetText())
	}
}

// TestCompetitiveCommanderAsksTheBracket keeps D-164: "competitive
// Commander" is three brackets from cEDH, so it fills no bracket.
func TestCompetitiveCommanderAsksTheBracket(t *testing.T) {
	out := classifyOut{Format: "commander", Theme: "dragons", PoolRule: "any_card", BudgetUSD: 500}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "colors"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A competitive Commander deck with dragons.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if p := st.Slots.GetPower(); p != nil {
		t.Errorf("competitive Commander filled the power %v", p)
	}
}

// TestANamedBracketBeatsTheSuperlative is D-209. A reader who names a
// bracket and asks for the best deck wants the best deck inside it.
func TestANamedBracketBeatsTheSuperlative(t *testing.T) {
	out := classifyOut{Format: "commander", Theme: "marvel", PoolRule: "any_card", BudgetUSD: 500, Power: "bracket 3"}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "colors"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "The best possible Commander deck at bracket 3.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got := st.Slots.GetPower().GetBracket(); got != 3 {
		t.Errorf("power bracket = %d, want the named bracket 3", got)
	}
}

// TestStrongestWaitsForTheFormat is D-1153. The superlative comes before
// the format, and the fill follows on the turn that names Commander.
func TestStrongestWaitsForTheFormat(t *testing.T) {
	first := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second := classifyOut{Format: "commander", Theme: "marvel", PoolRule: "any_card", BudgetUSD: 500}
	a, _ := testAgent(t, classifyStep(t, first), fits(t, "format", "theme"), askStep(t),
		classifyStep(t, second), fits(t, "commander", "colors"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "The best possible Marvel deck.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if p := st.Slots.GetPower(); p != nil {
		t.Fatalf("the power filled before the format: %v", p)
	}
	if _, err := a.Turn(context.Background(), st, "Commander.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if got := st.Slots.GetPower().GetBracket(); got != cedhBracket {
		t.Errorf("power bracket = %d after the format, want %d", got, cedhBracket)
	}
}
