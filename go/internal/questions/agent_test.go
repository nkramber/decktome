package questions

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// fakeConfig puts every role on the fake provider. Validate allows it
// while keys are not required.
func fakeConfig() *llm.Config {
	roles := map[llm.Role]llm.RoleSpec{}
	for _, r := range llm.Roles {
		roles[r] = llm.RoleSpec{Provider: llm.FakeName, Model: "fake-" + string(r), MaxOutputTokens: 1024}
	}
	return &llm.Config{VerifiedAt: "2026-08-24", Roles: roles}
}

func testAgent(t *testing.T, steps ...llm.Step) (*Agent, *llm.Script) {
	t.Helper()
	sc := llm.NewScript(steps...)
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	a, err := NewAgent(load(t), c, WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))))
	if err != nil {
		t.Fatalf("agent: %v", err)
	}
	return a, sc
}

func classifyStep(t *testing.T, out classifyOut) llm.Step {
	t.Helper()
	// The schema requires arrays, so a nil slice would fail validation the
	// same way a provider's null would.
	if out.Colors == nil {
		out.Colors = []string{}
	}
	if out.CommanderNames == nil {
		out.CommanderNames = []string{}
	}
	if out.LockedNames == nil {
		out.LockedNames = []string{}
	}
	if out.ClosedKeys == nil {
		out.ClosedKeys = []string{}
	}
	raw, err := json.Marshal(out)
	if err != nil {
		t.Fatal(err)
	}
	return llm.Step{Output: raw}
}

func askStep(t *testing.T, ps ...phrasing) llm.Step {
	t.Helper()
	for i := range ps {
		if ps[i].Options == nil {
			ps[i].Options = []string{}
		}
	}
	raw, err := json.Marshal(map[string]any{"questions": ps})
	if err != nil {
		t.Fatal(err)
	}
	return llm.Step{Output: raw}
}

// scoreStep is the second classify call of a turn: the gap score.
func scoreStep(t *testing.T, ss ...scored) llm.Step {
	t.Helper()
	raw, err := json.Marshal(map[string]any{"scores": ss})
	if err != nil {
		t.Fatal(err)
	}
	return llm.Step{Output: raw}
}

// fits scores every row a good catalog fit, which is the common case.
func fits(t *testing.T, rowIDs ...string) llm.Step {
	t.Helper()
	ss := make([]scored, 0, len(rowIDs))
	for _, id := range rowIDs {
		ss = append(ss, scored{RowID: id, Fit: 0.9, Reason: "the catalog row fits"})
	}
	return scoreStep(t, ss...)
}

func TestTurnFillsSlotsAndAsks(t *testing.T) {
	var out classifyOut
	out.Format, out.Theme, out.PoolRule, out.Power = "commander", "lifegain", "unknown", ""
	a, sc := testAgent(t,
		classifyStep(t, out),
		fits(t, "colors", "commander", "power_commander"),
		askStep(t,
			phrasing{RowID: "colors", Text: "Any color preference? Lifegain is strongest in white and black."},
			phrasing{RowID: "commander", Text: "Do you have a commander in mind?"},
			phrasing{RowID: "power_commander", Text: "Which bracket does your table play?"},
		))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "build me a lifegain deck", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(sc.Calls) != 3 {
		t.Fatalf("a turn made %d model calls, want 3", len(sc.Calls))
	}
	if res.Ready {
		t.Error("ready with empty slots")
	}
	if res.Invented != 0 {
		t.Errorf("invented %d questions from a good fit", res.Invented)
	}
	if st.Slots.Theme != "lifegain" || st.Slots.Format.GetId() != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("slots = %v", st.Slots)
	}
	if st.Slots.SlotStates["theme"] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("theme state = %v", st.Slots.SlotStates["theme"])
	}
	if len(res.Questions) == 0 {
		t.Fatal("no questions")
	}
	for _, q := range res.Questions {
		if q.Id == "" || q.Slot == "" || q.Text == "" {
			t.Errorf("incomplete question %+v", q)
		}
		if q.Invented {
			t.Errorf("question %q is marked invented", q.Id)
		}
	}
}

// TestInventedNeedsLowFit holds the conservative rule: a custom question
// counts only when the catalog fit is under the threshold (D-27).
func TestInventedNeedsLowFit(t *testing.T) {
	cases := []struct {
		name string
		fit  float64
		want bool
	}{
		{"poor fit invents", 0.10, true},
		{"fit at the threshold keeps the catalog", DefaultFitThreshold, false},
		{"good fit keeps the catalog", 0.90, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out classifyOut
			out.Format, out.Theme, out.PoolRule = "commander", "lifegain", "unknown"
			a, _ := testAgent(t,
				classifyStep(t, out),
				scoreStep(t, scored{RowID: "colors", Fit: tc.fit, CustomText: "Which two colors do you enjoy most?", Reason: "test"}),
				askStep(t, phrasing{RowID: "colors", Text: "Which two colors do you enjoy most?"}))
			st := NewState(false)
			res, err := a.Turn(context.Background(), st, "build me a lifegain deck", nil)
			if err != nil {
				t.Fatalf("turn: %v", err)
			}
			var got bool
			for _, q := range res.Questions {
				if q.Slot == "colors" {
					got = q.Invented
					if q.GapScore != tc.fit {
						t.Errorf("gap score = %v, want %v", q.GapScore, tc.fit)
					}
				}
			}
			if got != tc.want {
				t.Errorf("invented = %v, want %v", got, tc.want)
			}
			if got != (res.Invented == 1) {
				t.Errorf("invented count = %d, question invented = %v", res.Invented, got)
			}
		})
	}
}

// TestFrozenTurnMakesNoCall is D-68: a run has started, so the agent asks
// nothing and spends nothing.
func TestFrozenTurnMakesNoCall(t *testing.T) {
	a, sc := testAgent(t)
	st := NewState(true)
	st.Freeze()
	res, err := a.Turn(context.Background(), st, "switch to owned-only", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(sc.Calls) != 0 {
		t.Errorf("a frozen turn made %d model calls", len(sc.Calls))
	}
	if len(res.Questions) != 0 || !res.Ready {
		t.Errorf("a frozen turn asked %d questions", len(res.Questions))
	}
}

// TestNoRepeatAcrossTurns is the gate rule, with the model in the loop.
func TestNoRepeatAcrossTurns(t *testing.T) {
	var first classifyOut
	first.Format, first.Theme, first.PoolRule = "commander", "lifegain", "unknown"
	var second classifyOut
	second.Format, second.Theme, second.PoolRule, second.Power = "commander", "lifegain", "owned_first", "bracket 3"
	second.Colors = []string{"W", "B"}
	a, _ := testAgent(t,
		classifyStep(t, first),
		fits(t, "colors", "commander", "power_commander"),
		askStep(t,
			phrasing{RowID: "colors", Text: "Colors?"},
			phrasing{RowID: "commander", Text: "Commander?"},
			phrasing{RowID: "power_commander", Text: "Bracket?"}),
		classifyStep(t, second),
		fits(t, "pool"),
		askStep(t, phrasing{RowID: "pool", Text: "Library first?"}))
	st := NewState(true)
	seen := map[string]bool{}
	for turn, msg := range []string{"build me a lifegain deck", "white and black, bracket 3"} {
		res, err := a.Turn(context.Background(), st, msg, nil)
		if err != nil {
			t.Fatalf("turn %d: %v", turn+1, err)
		}
		for _, q := range res.Questions {
			if seen[q.Slot+q.Text] {
				t.Errorf("turn %d repeated %q", turn+1, q.Text)
			}
			seen[q.Slot+q.Text] = true
		}
	}
}

// TestUsageAccumulates proves a turn reports its tokens, which is what M-1
// and the owner's cost review read.
func TestUsageAccumulates(t *testing.T) {
	var out classifyOut
	out.Format, out.Theme, out.PoolRule = "commander", "lifegain", "any_card"
	steps := []llm.Step{
		classifyStep(t, out),
		fits(t, "commander"),
		askStep(t, phrasing{RowID: "commander", Text: "Commander?"}),
	}
	steps[0].Usage = &llm.Usage{InputTokens: 400, OutputTokens: 80}
	steps[1].Usage = &llm.Usage{InputTokens: 250, OutputTokens: 40}
	steps[2].Usage = &llm.Usage{InputTokens: 300, OutputTokens: 60}
	a, _ := testAgent(t, steps...)
	acc := llm.NewAccumulator(nil)
	if _, err := a.Turn(context.Background(), NewState(false), "lifegain deck", acc); err != nil {
		t.Fatalf("turn: %v", err)
	}
	rep := acc.Report()
	if rep.Calls != 3 {
		t.Errorf("calls = %d, want 3", rep.Calls)
	}
	if rep.Tokens == nil || rep.Tokens.InputTokens != 950 || rep.Tokens.OutputTokens != 180 {
		t.Errorf("tokens = %+v", rep.Tokens)
	}
}
