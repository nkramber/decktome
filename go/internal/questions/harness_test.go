package questions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"testing"

	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// The Turn harness (D-74). The 34 conversations of
// conversations_test.go drive Plan with a hand-edited context, so they
// can not see a defect in Turn, in the word rules, or in the snapshot.
// The conversations here run through Agent.Turn with a scripted
// classifier, and the state goes through Snapshot and Restore between
// turns, which is what production does on every turn.

// turnScript is one scripted turn: the user's message and what the
// classify role answers for it.
type turnScript struct {
	message  string
	classify classifyOut
}

// roleFake answers every model role. The classify role reads the script
// in order. The score role gives every row a good catalog fit, and the
// ask role keeps the catalog wording, so the harness measures the agent
// and not the model.
type roleFake struct {
	t        *testing.T
	mu       sync.Mutex
	classify []json.RawMessage
	calls    []llm.Call
}

func (f *roleFake) Name() string { return llm.FakeName }

func (f *roleFake) Complete(_ context.Context, call llm.Call) (llm.Response, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = append(f.calls, call)
	switch call.SchemaName {
	case "slot_fill":
		if len(f.classify) == 0 {
			return llm.Response{}, fmt.Errorf("the script holds no classify answer for call %d", len(f.calls))
		}
		out := f.classify[0]
		f.classify = f.classify[1:]
		return llm.Response{Output: out, Model: call.Model}, nil
	case "score_questions":
		var in struct {
			Rows []struct {
				ID string `json:"id"`
			} `json:"rows"`
		}
		if err := json.Unmarshal([]byte(call.Input), &in); err != nil {
			return llm.Response{}, err
		}
		ss := make([]scored, 0, len(in.Rows))
		for _, r := range in.Rows {
			ss = append(ss, scored{RowID: r.ID, Fit: 0.9, Reason: "the catalog row fits"})
		}
		raw, err := json.Marshal(map[string]any{"scores": ss})
		return llm.Response{Output: raw, Model: call.Model}, err
	case "phrase_questions":
		return llm.Response{Output: json.RawMessage(`{"questions":[]}`), Model: call.Model}, nil
	}
	return llm.Response{}, fmt.Errorf("unknown schema %q", call.SchemaName)
}

// played is what one scripted conversation produced.
type played struct {
	// rows are the row ids each turn asked, by turn from 1.
	rows [][]string
	// ready is the Ready flag of each turn.
	ready []bool
	// st is the state after the last turn, restored from its snapshot.
	st *State
	// fake holds every model call.
	fake *roleFake
}

// asked reports whether any turn asked the row.
func (p played) asked(rowID string) bool {
	for _, turn := range p.rows {
		for _, id := range turn {
			if id == rowID {
				return true
			}
		}
	}
	return false
}

// askedOn reports whether the turn asked the row. Turns count from 1.
func (p played) askedOn(turn int, rowID string) bool {
	if turn < 1 || turn > len(p.rows) {
		return false
	}
	for _, id := range p.rows[turn-1] {
		if id == rowID {
			return true
		}
	}
	return false
}

// play runs one conversation through Agent.Turn. Between turns the state
// is written to a snapshot, encoded, decoded, and restored, and the
// slots are cloned, so a field the snapshot drops is lost the way
// production loses it.
func play(t *testing.T, hasCollection bool, hints Hints, turns []turnScript) played {
	t.Helper()
	fake := &roleFake{t: t}
	for _, tr := range turns {
		fake.classify = append(fake.classify, classifyStep(t, tr.classify).Output)
	}
	c, err := llm.New(fakeConfig(), []llm.Provider{fake}, llm.WithoutJitter())
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	opts := []AgentOption{WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))}
	if hints != nil {
		opts = append(opts, WithHints(hints))
	}
	a, err := NewAgent(load(t), c, opts...)
	if err != nil {
		t.Fatalf("agent: %v", err)
	}
	st := NewState(hasCollection)
	st.SessionID = "harness"
	out := played{fake: fake}
	for i, tr := range turns {
		res, err := a.Turn(context.Background(), st, tr.message, nil)
		if err != nil {
			t.Fatalf("turn %d: %v", i+1, err)
		}
		var ids []string
		for _, rec := range st.Asks {
			if rec.Turn == st.Turn {
				ids = append(ids, rec.RowID)
			}
		}
		out.rows = append(out.rows, ids)
		out.ready = append(out.ready, res.Ready)
		st = roundTrip(t, st)
	}
	out.st = st
	return out
}

// roundTrip stores and restores a state, as agentsvc does between turns.
func roundTrip(t *testing.T, st *State) *State {
	t.Helper()
	raw, err := json.Marshal(st.Snapshot())
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var snap Snapshot
	if err := json.Unmarshal(raw, &snap); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	slots, err := proto.Marshal(st.Slots)
	if err != nil {
		t.Fatalf("marshal slots: %v", err)
	}
	var back mtgv1.Slots
	if err := proto.Unmarshal(slots, &back); err != nil {
		t.Fatalf("unmarshal slots: %v", err)
	}
	return Restore(st.SessionID, &back, snap)
}

// TestHarnessNearestFormat is gate conversation 50, "a format we do not
// support". The user names Brawl, hears the decline, accepts the nearest
// format, and the plan continues (D-74, D-112).
func TestHarnessNearestFormat(t *testing.T) {
	unknown := classifyOut{Format: "unknown", PoolRule: "unknown"}
	dragons := classifyOut{Format: "unknown", Theme: "dragons", PoolRule: "unknown", Colors: []string{"R"}}
	bracket := classifyOut{Format: "unknown", PoolRule: "unknown", Power: "bracket 2"}
	p := play(t, false, nil, []turnScript{
		{"I want a Brawl deck for Arena.", unknown},
		{"Yes, use the nearest format.", unknown},
		{"A dragon deck, red.", dragons},
		{"Bracket 2.", bracket},
	})
	if !p.askedOn(1, "format_unsupported") {
		t.Fatalf("turn 1 did not decline Brawl: %v", p.rows)
	}
	if p.st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("format = %v, want Commander after the user accepted it", p.st.Ctx.Format)
	}
	if len(p.rows[1]) == 0 {
		t.Error("turn 2 asked nothing after the user accepted the nearest format")
	}
	for turn := 2; turn <= 4; turn++ {
		if p.askedOn(turn, "format_unsupported") || p.askedOn(turn, "format") {
			t.Errorf("turn %d asked the format again: %v", turn, p.rows[turn-1])
		}
	}
	if p.st.Slots.GetPower().GetBracket() != 2 {
		t.Errorf("power = %v, want bracket 2", p.st.Slots.GetPower())
	}
}

// TestHarnessDelegationThenColorChange: the user hands the commander
// choice over, then changes the colors. The delegation stands (D-153).
func TestHarnessDelegationThenColorChange(t *testing.T) {
	first := commanderClassify()
	change := classifyOut{Format: "unknown", PoolRule: "unknown", Colors: []string{"W", "G"}}
	bracket := classifyOut{Format: "unknown", PoolRule: "unknown", Power: "bracket 3", BudgetUSD: 100}
	h := &fakeHints{
		commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
		identity:   map[string][]mtgv1.Color{"Karlov of the Ghost Council": {mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}},
	}
	p := play(t, false, h, []turnScript{
		{"A Commander lifegain deck, white and black, you pick the commander.", first},
		{"Actually make it white and green instead.", change},
		{"Bracket 3, and 100 dollars.", bracket},
	})
	if p.asked("commander") || p.asked("commander_pick") {
		t.Errorf("a commander question went out after the user handed the choice over: %v", p.rows)
	}
	states := p.st.Slots.GetSlotStates()
	if states["commander_pick"] != mtgv1.SlotState_SLOT_STATE_SKIPPED || states["commander"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("the delegation did not survive the color change: %v", states)
	}
	if len(p.st.CurrentOffer) != 0 {
		t.Errorf("names stayed on the table after a delegation: %v", p.st.CurrentOffer)
	}
	if !p.ready[len(p.ready)-1] {
		t.Errorf("the session is not ready: %v", states)
	}
}

// TestHarnessKeepACard is gate conversation 13, "a card to keep that is
// not the commander". The locked card and the commander survive every
// store cycle, and the locked row never asks what the user said (D-70,
// D-166).
func TestHarnessKeepACard(t *testing.T) {
	first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown",
		CommanderNames: []string{"Karlov of the Ghost Council"}, LockedNames: []string{"Sanguine Bond"}}
	first.Facts.NamedCard = true
	second := classifyOut{Format: "unknown", PoolRule: "unknown", Colors: []string{"W", "B"}, Power: "bracket 3"}
	third := classifyOut{Format: "unknown", PoolRule: "owned_first", LockedNames: []string{"Sanguine Bond"}}
	p := play(t, true, nil, []turnScript{
		{"Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.", first},
		{"Commander, white and black, bracket 3.", second},
		{"Build from my library first, and lock Sanguine Bond in.", third},
	})
	if p.asked("named_card_role") || p.asked("commander") {
		t.Errorf("a role or commander question went out for a named commander: %v", p.rows)
	}
	if got := p.st.LockedCards(); len(got) != 1 || got[0] != "Sanguine Bond" {
		t.Errorf("locked cards = %v, want Sanguine Bond", got)
	}
	if len(p.st.CommanderNames) != 1 || p.st.CommanderNames[0] != "Karlov of the Ghost Council" {
		t.Errorf("commander = %v, want Karlov", p.st.CommanderNames)
	}
}

// TestHarnessFormatSwitch is "the user changes the format". The switch
// retires the Commander questions, the 60-card rows take over, and the
// theme the user gave survives (D-125, D-195).
func TestHarnessFormatSwitch(t *testing.T) {
	first := classifyOut{Format: "unknown", Theme: "lifegain", PoolRule: "unknown"}
	second := classifyOut{Format: "commander", PoolRule: "unknown", Colors: []string{"W", "B"}}
	third := classifyOut{Format: "modern", PoolRule: "unknown"}
	fourth := classifyOut{Format: "unknown", PoolRule: "unknown", Power: "fnm", BudgetUSD: 100}
	p := play(t, false, nil, []turnScript{
		{"Build me a lifegain deck.", first},
		{"Commander, white and black.", second},
		{"Actually, make it Modern instead. Sixty cards.", third},
		{"FNM level, and 100 dollars.", fourth},
	})
	if p.st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_MODERN {
		t.Fatalf("format = %v, want Modern", p.st.Ctx.Format)
	}
	for turn := 3; turn <= 4; turn++ {
		for _, id := range p.rows[turn-1] {
			if id == "commander" || id == "commander_pick" || id == "power_commander" {
				t.Errorf("turn %d asked a Commander row in Modern: %v", turn, p.rows[turn-1])
			}
		}
	}
	if !p.askedOn(3, "power_sixty") {
		t.Errorf("turn 3 did not ask the 60-card power row: %v", p.rows[2])
	}
	if p.st.Slots.GetPower().GetSixtyStep() != mtgv1.SixtyStep_SIXTY_STEP_FNM {
		t.Errorf("power = %v, want FNM", p.st.Slots.GetPower())
	}
	if p.st.Slots.GetTheme() != "lifegain" {
		t.Errorf("theme = %q, want lifegain", p.st.Slots.GetTheme())
	}
}

// TestHarnessOutOfScopeRepeat sends a second request for another game.
// The decline goes out twice, and the third message gets the normal
// questions (D-99).
func TestHarnessOutOfScopeRepeat(t *testing.T) {
	other := classifyOut{Format: "unknown", PoolRule: "unknown"}
	other.Facts.OutOfScope = true
	magic := classifyOut{Format: "commander", Theme: "dragons", PoolRule: "unknown"}
	p := play(t, false, nil, []turnScript{
		{"Can you build me a Yu-Gi-Oh deck?", other},
		{"What about a Pokemon deck?", other},
		{"Fine, Magic then. Commander, a dragon deck.", magic},
	})
	if !p.askedOn(1, "out_of_scope") || !p.askedOn(2, "out_of_scope") {
		t.Errorf("the decline did not go out on both turns: %v", p.rows)
	}
	if len(p.rows[0]) != 1 || len(p.rows[1]) != 1 {
		t.Errorf("an out-of-scope turn asked more than the decline: %v", p.rows)
	}
	if p.askedOn(3, "out_of_scope") || len(p.rows[2]) == 0 {
		t.Errorf("turn 3 did not get the normal questions: %v", p.rows[2])
	}
	if got := p.st.Slots.GetSlotStates()["scope"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("scope state = %v, want SKIPPED after the user asked for a Magic deck", got)
	}
}

// TestHarnessNearestFormatByComparison is gate conversation 50 word for
// word. Turn 2 describes Brawl by comparison, so nothing fills the
// format and the decline stays silent (D-210). Turn 3 accepts Commander
// by name, and the plan continues.
func TestHarnessNearestFormatByComparison(t *testing.T) {
	unknown := classifyOut{Format: "unknown", PoolRule: "unknown"}
	accept := classifyOut{Format: "commander", Theme: "dragons", PoolRule: "unknown", Colors: []string{"R"}}
	bracket := classifyOut{Format: "unknown", PoolRule: "unknown", Power: "bracket 2"}
	p := play(t, false, nil, []turnScript{
		{"I want a Brawl deck for Arena.", unknown},
		{"It is like Commander but 60 cards on Arena.", unknown},
		{"Fine, treat it as Commander. A dragon deck, red.", accept},
		{"Bracket 2.", bracket},
	})
	if !p.askedOn(1, "format_unsupported") {
		t.Fatalf("turn 1 did not decline Brawl: %v", p.rows)
	}
	if p.askedOn(2, "format_unsupported") || p.askedOn(2, "format") {
		t.Errorf("turn 2 repeated the format question: %v", p.rows[1])
	}
	if p.askedOn(3, "format_unsupported") || p.askedOn(3, "format") {
		t.Errorf("turn 3 asked the format after the user accepted it: %v", p.rows[2])
	}
	if p.st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("format = %v, want Commander", p.st.Ctx.Format)
	}
	if got := p.st.Slots.GetSlotStates()["format"]; got != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("format state = %v, want FILLED", got)
	}
}
