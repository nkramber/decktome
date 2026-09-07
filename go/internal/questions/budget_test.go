package questions

import (
	"context"
	"io"
	"log/slog"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The budget tests: the cap, the buy list, the proxy user, and the scope
// of a cap.

// TestBudgetScopeIsTyped is D-238. The scope is a typed value, so an
// option match or a key name must not close it with the value
// UNSPECIFIED (D-83).
func TestBudgetScopeIsTyped(t *testing.T) {
	cases := []struct {
		name, answer string
		closed       []string
		want         mtgv1.BudgetScope
		wantState    mtgv1.SlotState
	}{
		{"a key name closes nothing", "sure, the second option", []string{"budget_scope"},
			mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED, mtgv1.SlotState_SLOT_STATE_ASKED},
		{"the whole deck carries its value", "The whole deck.", nil,
			mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK, mtgv1.SlotState_SLOT_STATE_FILLED},
		{"the cards I buy carries its value", "The cards I buy.", nil,
			mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY, mtgv1.SlotState_SLOT_STATE_FILLED},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			first := commanderClassify()
			first.CommanderNames = []string{"Karlov of the Ghost Council"}
			first.BudgetUSD = 100
			first.Facts.BudgetAmbiguous = true
			second := classifyOut{Format: "unknown", PoolRule: "unknown", ClosedKeys: tc.closed}
			a, _ := testAgent(t,
				classifyStep(t, first), fits(t, "power_commander", "pool", "budget_scope"), askStep(t),
				classifyStep(t, second), fits(t), askStep(t))
			st := NewState(true)
			if _, err := a.Turn(context.Background(), st, "Karlov lifegain, white and black, 100 dollars.", nil); err != nil {
				t.Fatalf("turn 1: %v", err)
			}
			if st.Slots.GetSlotStates()["budget_scope"] != mtgv1.SlotState_SLOT_STATE_ASKED {
				t.Fatalf("the scope row did not fire: %v", st.Slots.GetSlotStates())
			}
			if _, err := a.Turn(context.Background(), st, tc.answer, nil); err != nil {
				t.Fatalf("turn 2: %v", err)
			}
			if got := st.Slots.GetBudgetScope(); got != tc.want {
				t.Errorf("scope = %v, want %v", got, tc.want)
			}
			if got := st.Slots.GetSlotStates()["budget_scope"]; got != tc.wantState {
				t.Errorf("scope state = %v, want %v", got, tc.wantState)
			}
			if st.Slots.GetSlotStates()["budget_scope"] == mtgv1.SlotState_SLOT_STATE_FILLED &&
				st.Slots.GetBudgetScope() == mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED {
				t.Error("the scope closed with the value UNSPECIFIED")
			}
		})
	}
}

// TestProxyUserGetsNoBudgetQuestion is D-111. A user who proxies every
// card has no budget.
func TestProxyUserGetsNoBudgetQuestion(t *testing.T) {
	out := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Facts.BuyList = true
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "power_sixty"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "We proxy everything at our table. A Modern burn deck, mono red.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if q := question(res.Questions, "budget"); q != nil {
		t.Errorf("a budget question went out to a proxy user: %q", q.GetText())
	}
	if st.Slots.GetSlotStates()["budget"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("budget state = %v, want skipped", st.Slots.GetSlotStates()["budget"])
	}
}

// TestNoCollectionAsksTheBudget is D-168. A user with no library buys
// every card, so the budget row must ask.
func TestNoCollectionAsksTheBudget(t *testing.T) {
	out := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Power = "fnm"
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "budget"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "A Modern burn deck, mono red, FNM level.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if !st.Ctx.BuyList {
		t.Fatal("a session with no collection carries no buy list")
	}
	if q := question(res.Questions, "budget"); q == nil {
		t.Errorf("the budget row did not fire: %v", ids2(res.Questions))
	}
}

// TestNoSpendingLimitClosesTheBudget keeps the budget row off a user who
// has answered it (D-168).
func TestNoSpendingLimitClosesTheBudget(t *testing.T) {
	if !noSpendingLimit("money is no object") {
		t.Error("money is no object was not read")
	}
	if noSpendingLimit("the best deck under budget") {
		t.Error("a budget request read as a refusal of the cap")
	}
	out := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Power = "tournament"
	a, _ := testAgent(t, classifyStep(t, out), fits(t), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st,
		"The strongest Modern burn deck, mono red, tournament level, money is no object.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if q := question(res.Questions, "budget"); q != nil {
		t.Errorf("the budget row asked a user who named no cap: %q", q.GetText())
	}
	if !st.Ctx.Filled["budget"] {
		t.Error("the budget slot is still open")
	}
}

// TestBudgetScopeIsStored is D-238. The row asked which the cap covers,
// and the answer reached no slot, so the agent asked and discarded it.
func TestBudgetScopeIsStored(t *testing.T) {
	for _, tc := range []struct {
		word string
		want mtgv1.BudgetScope
	}{
		{"buy", mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY},
		{"deck", mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK},
		{"unknown", mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED},
		{"", mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED},
	} {
		if got := budgetScope(tc.word); got != tc.want {
			t.Errorf("budgetScope(%q) = %v, want %v", tc.word, got, tc.want)
		}
	}
	out := commanderClassify()
	out.BudgetUSD, out.BudgetScope = 100, "deck"
	a, _ := testAgentHints(t, nil, classifyStep(t, out), fits(t, "power_commander"), askStep(t))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "A lifegain deck, 100 dollars for the whole deck.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got := st.Slots.GetBudgetScope(); got != mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK {
		t.Errorf("scope = %v, want the whole deck", got)
	}
	// The answer closes its own row, so the agent does not ask again.
	if st.Slots.GetSlotStates()["budget_scope"] == mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Error("the scope row is still outstanding after the user answered it")
	}
}

// TestBudgetScopeIsNotAskedWhenTheWordsAnswerIt is D-253. The row must
// not ask what the message has already said.
func TestBudgetScopeIsNotAskedWhenTheWordsAnswerIt(t *testing.T) {
	for _, s := range []string{
		"build owned-first with a buy list", "40 dollars for the cards to buy",
	} {
		if !namesTheBuyList(s) {
			t.Errorf("%q: the buy list was not read", s)
		}
	}
	for _, s := range []string{"a lifegain deck for 40 dollars", "no more than 40 dollars"} {
		if namesTheBuyList(s) {
			t.Errorf("%q: a plain budget was read as a buy list", s)
		}
	}

	// A message that names the buy list answers the scope row.
	out := commanderClassify()
	out.BudgetUSD = 40
	a, _ := testAgentHints(t, nil, classifyStep(t, out), fits(t, "power_commander"), askStep(t))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "Build owned-first with a buy list, no more than 40 dollars.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got := st.Slots.GetBudgetScope(); got != mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY {
		t.Errorf("scope = %v, want the cards to buy", got)
	}
	if st.Slots.GetSlotStates()["budget_scope"] == mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Error("the scope row asked what the message had already said")
	}

	// A proxy user has no budget, so there is no scope to ask about.
	out2 := commanderClassify()
	a2, _ := testAgentHints(t, nil, classifyStep(t, out2), fits(t, "power_commander"), askStep(t))
	st2 := NewState(true)
	if _, err := a2.Turn(context.Background(), st2, "I proxy anything over 20 dollars.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got := st2.Slots.GetSlotStates()["budget_scope"]; got == mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Error("a proxy user was asked what their budget covers")
	}
}

// TestBudgetNumberAnswersTheCardsToBuyQuestion is D-288. The budget row
// names the scope in its text, so "$100" as its answer is a cap on the
// cards to buy, and the scope row must not ask again.
func TestBudgetNumberAnswersTheCardsToBuyQuestion(t *testing.T) {
	a := &Agent{log: slog.New(slog.NewTextHandler(io.Discard, nil))}
	st := Restore("s", &mtgv1.Slots{BudgetUsd: 100}, Snapshot{Version: SnapshotVersion})
	ruleBudgetScope(a, st, turnWords{Message: "$100", Open: []string{"budget"}})
	if st.Slots.GetBudgetScope() != mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY {
		t.Errorf("scope = %v, want cards to buy", st.Slots.GetBudgetScope())
	}
	if st.Slots.GetSlotStates()["budget_scope"] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("budget_scope state = %v, want filled", st.Slots.GetSlotStates()["budget_scope"])
	}
	// The words win: "$100 for the whole deck" as the same answer is the deck value.
	st = Restore("s", &mtgv1.Slots{BudgetUsd: 100, SlotStates: map[string]mtgv1.SlotState{"budget_scope": mtgv1.SlotState_SLOT_STATE_ASKED}}, Snapshot{Version: SnapshotVersion})
	ruleBudgetScope(a, st, turnWords{Message: "$100 for the whole deck", Open: []string{"budget", "budget_scope"}})
	if st.Slots.GetBudgetScope() != mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK {
		t.Errorf("scope = %v, want the whole deck", st.Slots.GetBudgetScope())
	}
	// No budget number, no scope.
	st = Restore("s", &mtgv1.Slots{}, Snapshot{Version: SnapshotVersion})
	ruleBudgetScope(a, st, turnWords{Message: "no", Open: []string{"budget"}})
	if st.Slots.GetBudgetScope() != mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED {
		t.Errorf("scope = %v, want unspecified", st.Slots.GetBudgetScope())
	}
}
