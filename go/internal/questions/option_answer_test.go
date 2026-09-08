package questions

import (
	"io"
	"log/slog"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func optionAgent(t *testing.T) *Agent {
	t.Helper()
	cat, err := Load()
	if err != nil {
		t.Fatalf("catalog: %v", err)
	}
	return &Agent{cat: cat, log: slog.New(slog.NewTextHandler(io.Discard, nil))}
}

// asked puts one question of a row out, as the planner does.
func asked(st *State, questionID, rowID, slot, key string) {
	st.Asks = append(st.Asks, Ask{QuestionID: questionID, RowID: rowID, Slot: slot, Key: key})
	if st.Slots.SlotStates == nil {
		st.Slots.SlotStates = map[string]mtgv1.SlotState{}
	}
	st.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_ASKED
}

// TestOptionAnswerFillsItsSlot is D-597, and the session that named it
// is oUZMC0F2vHe7GGl24LIP. The reader picked "4 optimized", the fourth
// option of the power row, and the slot stayed asked. The index named
// one value exactly, and the engine threw that away: it made the option
// text, folded it into the message, and waited for the classifier to
// write a string the slot could read. The engine reads the index now.
func TestOptionAnswerFillsItsSlot(t *testing.T) {
	a := optionAgent(t)

	t.Run("the power option of the stuck session", func(t *testing.T) {
		st := NewState(false)
		asked(st, "q2-power_commander", "power_commander", "power", "power")
		st.OptionAnswers = []OptionAnswer{{QuestionID: "q2-power_commander", Index: 3}}
		a.applyOptionAnswers(st)
		if got := st.Slots.GetPower().GetBracket(); got != 4 {
			t.Errorf("bracket = %d, want 4: the reader picked \"4 optimized\"", got)
		}
		if st.Slots.GetSlotStates()["power"] == mtgv1.SlotState_SLOT_STATE_ASKED {
			t.Error("the power slot is still asked after the reader answered it")
		}
	})

	t.Run("every option of every typed row fills its slot", func(t *testing.T) {
		cat, err := Load()
		if err != nil {
			t.Fatal(err)
		}
		rows := 0
		for _, row := range cat.Rows {
			if len(row.OptionValues) == 0 {
				continue
			}
			rows++
			for i := range row.OptionValues {
				st := NewState(true)
				st.Ctx.HasCollection = true
				key := row.StateKey()
				asked(st, "q1-"+row.ID, row.ID, row.Slot, key)
				st.OptionAnswers = []OptionAnswer{{QuestionID: "q1-" + row.ID, Index: i}}
				a.applyOptionAnswers(st)
				if st.Slots.GetSlotStates()[key] == mtgv1.SlotState_SLOT_STATE_ASKED {
					t.Errorf("row %q option %d (%q) left the slot asked", row.ID, i, row.Options[i])
				}
			}
		}
		if rows == 0 {
			t.Fatal("no row carries option values, so this test proves nothing")
		}
		t.Logf("checked %d rows with typed options", rows)
	})

	t.Run("a row with no option values is left to the classifier", func(t *testing.T) {
		st := NewState(false)
		asked(st, "q1-house_rules", "house_rules", "house_rules", "house_rules")
		st.OptionAnswers = []OptionAnswer{{QuestionID: "q1-house_rules", Index: 0}}
		a.applyOptionAnswers(st)
		if st.Slots.GetSlotStates()["house_rules"] != mtgv1.SlotState_SLOT_STATE_ASKED {
			t.Error("a row with no typed values closed a slot on a name alone (D-83)")
		}
	})

	t.Run("an index outside the options changes nothing", func(t *testing.T) {
		st := NewState(false)
		asked(st, "q1-power_commander", "power_commander", "power", "power")
		for _, i := range []int{-1, 5, 99} {
			st.OptionAnswers = []OptionAnswer{{QuestionID: "q1-power_commander", Index: i}}
			a.applyOptionAnswers(st)
		}
		if st.Slots.GetPower() != nil {
			t.Errorf("an index outside the options set a value: %v", st.Slots.GetPower())
		}
	})

	// An owned rule needs a collection behind it (D-371). The option path
	// takes the same rule as the classifier path.
	t.Run("an owned pool option with no collection reads as any card", func(t *testing.T) {
		st := NewState(false)
		st.Ctx.HasCollection = false
		asked(st, "q1-pool", "pool", "pool_rule", "pool_rule")
		st.OptionAnswers = []OptionAnswer{{QuestionID: "q1-pool", Index: 1}}
		a.applyOptionAnswers(st)
		if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
			t.Errorf("pool rule = %v, want ANY_CARD with no collection", got)
		}
	})
}

// TestEveryOptionRowIsTypedOrLeftToTheClassifier keeps the catalog
// honest. A row with options on a typed slot has no deterministic exit
// unless it carries option values, and that is the shape of F-70. This
// names the rows that still lack them, so the count only goes down.
func TestEveryOptionRowIsTypedOrLeftToTheClassifier(t *testing.T) {
	cat, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	var untyped []string
	for _, row := range cat.Rows {
		if len(row.Options) == 0 || len(row.OptionValues) > 0 {
			continue
		}
		if typedSlots[row.Slot] {
			untyped = append(untyped, row.ID)
		}
	}
	// These rows carry no typed value yet. Each one still rests on the
	// classifier, with CloseStalled as the net under it (D-351). Take a
	// row off this list when it gains option values, and never add one.
	const want = 9
	if len(untyped) > want {
		t.Errorf("%d rows with options on a typed slot carry no option values, want at most %d: %v", len(untyped), want, untyped)
	}
	t.Logf("rows still left to the classifier: %v", untyped)
}

// TestReaskStalledAsksOnceThenGivesUp is D-599. A question the reader
// answered, whose answer never reached the slot, goes out once more. The
// second stall re-opens nothing, so CloseStalled skips the key and the
// build takes the default. The reader is never asked the same question
// a third time, and the chat never stops.
func TestReaskStalledAsksOnceThenGivesUp(t *testing.T) {
	st := NewState(false)
	asked(st, "q1-power_commander", "power_commander", "power", "power")
	st.Ctx.Asked = map[string]bool{"power_commander": true}
	st.Ctx.Filled = map[string]bool{"power": true}
	// The reader replied to this question, and the reply missed the slot.
	// A question nobody answered is not stalled.
	st.AnsweredQuestions = []string{"q1-power_commander"}

	// A question with no reply stays where it is.
	quiet := NewState(false)
	asked(quiet, "q1-power_commander", "power_commander", "power", "power")
	if got := quiet.ReaskStalled(); len(got) != 0 {
		t.Errorf("a question nobody answered re-opened %v, want none", got)
	}

	first := st.ReaskStalled()
	if len(first) != 1 || first[0] != "power" {
		t.Fatalf("the first stall re-opened %v, want [power]", first)
	}
	// The slot is open again, and the no-repeat rule no longer holds the
	// row, so the planner may ask it.
	if got := st.Slots.GetSlotStates()["power"]; got != mtgv1.SlotState_SLOT_STATE_UNSPECIFIED {
		t.Errorf("power = %v, want UNSPECIFIED after the re-ask", got)
	}
	if st.Ctx.Asked["power_commander"] {
		t.Error("the row is still marked asked, so the planner will not ask it again")
	}
	if st.Ctx.Filled["power"] {
		t.Error("the key is still marked filled, so the planner will not ask it again")
	}

	// The reader answers again, and the answer misses again.
	st.Slots.SlotStates["power"] = mtgv1.SlotState_SLOT_STATE_ASKED
	st.AnsweredQuestions = []string{"q1-power_commander"}
	if second := st.ReaskStalled(); len(second) != 0 {
		t.Errorf("the second stall re-opened %v, want none: the net skips it now", second)
	}
	if got := st.Slots.GetSlotStates()["power"]; got != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Errorf("power = %v, want it left for CloseStalled", got)
	}

	// CloseStalled is the way on, and it skips the key.
	st.Turn += StallGrace
	closed, _ := st.CloseStalled()
	if len(closed) != 1 || closed[0] != "power" {
		t.Errorf("CloseStalled closed %v, want [power]", closed)
	}
}
