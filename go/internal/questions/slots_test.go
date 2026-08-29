package questions

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// commanderClassify is one classify answer for a Commander session whose
// format, theme, and colors are already clear.
func commanderClassify() classifyOut {
	var out classifyOut
	out.Format, out.Theme, out.PoolRule = "commander", "lifegain", "any_card"
	out.Colors = []string{"W", "B"}
	return out
}

// TestNamedCommanderIsNotALockedCard is D-70. A locked list that holds
// the commander asks the user to keep or cut their own commander. The
// locked row is retired (D-260), and the list it read still feeds the
// build, so the list must stay right: the commander is never in it, and
// a card named later is.
func TestNamedCommanderIsNotALockedCard(t *testing.T) {
	first := commanderClassify()
	first.CommanderNames = []string{"Karlov of the Ghost Council"}
	first.Facts.NamedCard = true
	second := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second.LockedNames = []string{"Sanguine Bond"}
	second.Facts.NamedCard = true

	a, _ := testAgent(t,
		classifyStep(t, first),
		fits(t, "power_commander"),
		askStep(t),
		classifyStep(t, second),
		fits(t),
		askStep(t))
	st := NewState(false)

	if _, err := a.Turn(context.Background(), st, "karlov lifegain deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if len(st.LockedCards()) != 0 {
		t.Errorf("locked cards = %v, want none", st.LockedCards())
	}
	// A named commander closes every commander row, so the pick row and the
	// role row can not ask the same thing in other words.
	if !st.Ctx.CommanderSet {
		t.Error("the commander is not marked set")
	}
	for _, key := range commanderKeys {
		if !st.Ctx.Filled[key] {
			t.Errorf("key %q is still open although the user named a commander", key)
		}
	}

	res, err := a.Turn(context.Background(), st, "and sanguine bond as well", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if got := st.LockedCards(); len(got) != 1 || !sameCard(got[0], "Sanguine Bond") {
		t.Errorf("locked cards = %v, want Sanguine Bond alone", got)
	}
	if q := question(res.Questions, "locked"); q != nil {
		t.Errorf("a retired row asked about the locked card: %q", q.GetText())
	}
}

// TestNameThatBecomesTheCommanderStopsBeingLocked covers the order the
// user writes in: the card to keep arrives first, the role second.
func TestNameThatBecomesTheCommanderStopsBeingLocked(t *testing.T) {
	st := NewState(false)
	st.AddLocked("Karlov of the Ghost Council")
	if len(st.LockedCards()) != 1 {
		t.Fatal("a named card is not locked")
	}
	st.SetCommander("karlov of the ghost council")
	if len(st.LockedCards()) != 0 {
		t.Errorf("locked cards = %v, want none after the card became the commander", st.LockedCards())
	}
}

// TestSuggestionFiresThePickRow is D-71. A commander row that asks two
// things at once scores badly. Split, the base row asks whether the user
// has one, and the pick row carries the names. The Suggested fact is what
// fires the pick row.
func TestSuggestionFiresThePickRow(t *testing.T) {
	hints := &fakeHints{commanders: []string{
		"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Trelasarra, Moon Dancer",
		"Liesa, Shroud of Dusk", "Ayli, Eternal Pilgrim", "Vito, Thorn of the Dusk Rose",
	}}
	wants := classifyOut{Format: "unknown", PoolRule: "unknown"}
	wants.Facts.WantsSuggestion = true
	// "None of those" asks for other names. An unrelated message does not.
	none := wants
	other := classifyOut{Format: "unknown", PoolRule: "unknown"}

	// The pick row is fixed (D-131), so a turn that asks it alone makes
	// the classify call and no other.
	a, _ := testAgentHints(t, hints,
		classifyStep(t, commanderClassify()),
		fits(t, "commander", "power_commander"),
		askStep(t),
		classifyStep(t, wants),
		classifyStep(t, none),
		classifyStep(t, other))
	st := NewState(false)

	res, err := a.Turn(context.Background(), st, "lifegain commander deck", nil)
	if err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	base := question(res.Questions, "commander")
	if base == nil {
		t.Fatal("the commander row did not fire")
	}
	if strings.Count(base.Text, "?") != 1 {
		t.Errorf("the commander row asks two things at once: %q", base.Text)
	}
	for _, name := range hints.commanders {
		if strings.Contains(base.Text, name) {
			t.Errorf("the base commander row still names candidates: %q", base.Text)
		}
	}

	res, err = a.Turn(context.Background(), st, "no, suggest one", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !st.Ctx.Suggested {
		t.Fatal("the wants_suggestion fact did not reach the planner")
	}
	pick := question(res.Questions, "commander")
	if pick == nil {
		t.Fatal("the pick row did not fire after the user asked for a suggestion")
	}
	for _, name := range hints.commanders[:3] {
		if !strings.Contains(pick.Text, name) {
			t.Errorf("the pick question does not name %q: %q", name, pick.Text)
		}
	}

	// D-73: "none" gets three other names, not the same three again.
	res, err = a.Turn(context.Background(), st, "none of those", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	second := question(res.Questions, "commander")
	if second == nil {
		t.Fatal("the pick row did not repeat after a none answer")
	}
	for _, name := range hints.commanders[:3] {
		if strings.Contains(second.Text, name) {
			t.Errorf("the second round names %q again: %q", name, second.Text)
		}
	}
	for _, name := range hints.commanders[3:] {
		if !strings.Contains(second.Text, name) {
			t.Errorf("the second round does not name %q: %q", name, second.Text)
		}
	}
	if pick.Id == second.Id {
		t.Errorf("two questions share the id %q", pick.Id)
	}

	// A message that does not ask for other names keeps the same three,
	// and it gets no second copy of the same question. Three other names
	// on every turn read as an ignored answer (D-73), and the same three
	// twice is a duplicate (D-163).
	res, err = a.Turn(context.Background(), st, "my table accepts stax", nil)
	if err != nil {
		t.Fatalf("turn 4: %v", err)
	}
	if third := question(res.Questions, "commander"); third != nil {
		t.Errorf("the pick row asked again with the same names: %q", third.Text)
	}
	for _, name := range hints.commanders[3:] {
		if !hasName(st.CurrentOffer, name) {
			t.Errorf("an unrelated answer changed the names, %q is gone: %v", name, st.CurrentOffer)
		}
	}
	// The question is out with no answer, so the session is not ready.
	if st.Ready(load(t)) {
		t.Error("the session calls itself complete with the commander unanswered")
	}
}

// TestSessionIDIsTheCacheKey proves every call of one session carries the
// provider cache key (D-72).
func TestSessionIDIsTheCacheKey(t *testing.T) {
	a, sc := testAgent(t,
		classifyStep(t, commanderClassify()),
		fits(t, "commander", "power_commander"),
		askStep(t))
	st := NewState(false)
	st.SessionID = "sess-42"
	if _, err := a.Turn(context.Background(), st, "lifegain commander deck", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(sc.Calls) != 3 {
		t.Fatalf("a turn made %d calls, want 3", len(sc.Calls))
	}
	for i, call := range sc.Calls {
		if call.CacheKey != "sess-42" {
			t.Errorf("call %d cache key = %q, want the session id", i+1, call.CacheKey)
		}
	}
}

// question returns the first question for one slot, or nil.
func question(qs []*mtgv1.Question, slot string) *mtgv1.Question {
	for _, q := range qs {
		if q.GetSlot() == slot {
			return q
		}
	}
	return nil
}

// TestAskedKeyCanClose covers the advisory keys, which carry no typed
// value. The no-repeat rule drops an asked row from the plan, so without
// this path the house-limits question could never close and the session
// was never ready.
func TestAskedKeyCanClose(t *testing.T) {
	first := commanderClassify()
	first.Power = "bracket 3"
	first.HouseRules = "any card, no ban list"
	first.Facts.HouseFormat = true
	second := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second.ClosedKeys = []string{"house_format_limits"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "house_format_limits"), askStep(t),
		classifyStep(t, second), fits(t), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "any card, no ban list, but a commander deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.Asked["house_format_limits"] {
		t.Fatalf("the house-limits row did not fire: %v", st.Ctx.Asked)
	}
	if _, err := a.Turn(context.Background(), st, "I will say the limits myself", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !st.Ctx.Filled["house_format_limits"] {
		t.Error("the house-limits key did not close although the agent asked it last turn")
	}
}

// TestOnlyAnOpenQuestionCloses is D-83. The classifier can retire power
// and the pool rule by name from "Brago blink deck from my library".
// Neither word appears in that message, and neither question is out, so
// the session must not end after one question.
func TestOnlyAnOpenQuestionCloses(t *testing.T) {
	first := commanderClassify()
	// The classifier claims power and pool rule are done, and gives no
	// value for either.
	first.PoolRule = "unknown"
	first.ClosedKeys = []string{"power", "pool_rule"}
	third := classifyOut{Format: "unknown", PoolRule: "unknown", Power: "bracket 3"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, third), fits(t), askStep(t))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "brago blink deck from my library", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	for _, key := range []string{"power", "pool_rule"} {
		if st.Ctx.Filled[key] {
			t.Errorf("the classifier closed typed slot %q by name, with no value", key)
		}
	}
	if st.Ready(load(t)) {
		t.Error("the session called itself complete with power and the card pool unanswered")
	}
	// The same slot closes as soon as a value arrives.
	if _, err := a.Turn(context.Background(), st, "bracket 3", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !st.Ctx.Filled["power"] {
		t.Error("a typed power value did not close the power slot")
	}
}

// TestTypedSlotNeverClosesWithoutAValue is the invariant the third
// version of D-83 rests on. A format closed by name stays empty, every
// row that triggers on the format stops firing, and the session ends
// with no power level. A deck can not be built from a slot that says
// "answered" and holds nothing.
func TestTypedSlotNeverClosesWithoutAValue(t *testing.T) {
	// The budget carries a value from the first message. Without one the
	// budget row fires, because a session with no collection buys every
	// card, and this test is about the format alone (D-168).
	first := classifyOut{Format: "unknown", PoolRule: "unknown", BudgetUSD: 40}
	// Turn 2 answers the format question by name alone, with no value.
	byName := classifyOut{Format: "unknown", PoolRule: "unknown"}
	byName.ClosedKeys = []string{"format"}
	// Turn 3 gives the value.
	byValue := classifyOut{Format: "commander", PoolRule: "unknown"}

	// Turn 2 asks nothing. The format question is out with no answer, so
	// no other row may ask the format in other words (D-126).
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "format", "theme"), askStep(t),
		classifyStep(t, byName),
		classifyStep(t, byValue), fits(t, "power_commander"), askStep(t))
	st := NewState(false)
	// No message below names a format. The word rules read a format the
	// classifier missed (D-116), and this test must isolate the classify
	// path that D-83 is about.
	if _, err := a.Turn(context.Background(), st, "i need a deck for fnm on friday", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "yes, that one", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if st.Ctx.Filled["format"] {
		t.Error("the format closed on a name alone, so no value reached the deck generator")
	}
	if _, err := a.Turn(context.Background(), st, "the first option", nil); err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if !st.Ctx.Filled["format"] {
		t.Fatal("a typed format value did not close the format slot")
	}
	// The value, not only the state, must reach the slots.
	if st.Slots.GetFormat().GetId() != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("the format value is %v, want Commander", st.Slots.GetFormat().GetId())
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Error("the planner does not know the format, so every row that triggers on it stops firing")
	}
}

// TestRestoreWithoutSnapshot opens a session stored before the private
// state existed. It must not panic, and it must ask again.
func TestRestoreWithoutSnapshot(t *testing.T) {
	st := Restore("sess-2", nil, Snapshot{})
	if st.Slots == nil || st.Slots.SlotStates == nil || st.Ctx.Filled == nil || st.Ctx.Asked == nil {
		t.Fatalf("restore gave an unusable state: %+v", st)
	}
	if len(load(t).Plan(st.Ctx)) == 0 {
		t.Error("a session with no snapshot asks nothing")
	}
}

// TestInventedQuestionKeepsTheOptions is D-37 through the back door.
// A replacement of the card-pool question that offers two of the three
// pool modes loses one. A user who never sees "only my library" can not
// choose it.
func TestInventedQuestionKeepsTheOptions(t *testing.T) {
	out := commanderClassify()
	out.PoolRule = "unknown"
	a, _ := testAgent(t,
		classifyStep(t, out),
		scoreStep(t,
			scored{RowID: "pool", Fit: 0.05, CustomText: "Which cards may I use: only what you own, or anything legal?", Reason: "test"},
			scored{RowID: "commander", Fit: 0.9, Reason: "fits"}),
		askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "lifegain commander deck from my library", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	q := question(res.Questions, "pool_rule")
	if q == nil {
		t.Fatal("the card-pool question did not fire")
	}
	if !q.GetInvented() {
		t.Fatal("the replacement was not used")
	}
	row, _ := load(t).Row("pool")
	if len(q.GetOptions()) != len(row.Options) {
		t.Errorf("the replacement offers %d options, the row offers %d: %v",
			len(q.GetOptions()), len(row.Options), q.GetOptions())
	}
}

// TestRewordIsRefused is D-88. A replacement that only adds the user's
// colors, format, or card name repeats the row, and the ask role adds
// those words anyway.
func TestRewordIsRefused(t *testing.T) {
	row := "Must the deck keep Sanguine Bond, or may I cut a card that does not fit the plan?"
	cases := []struct {
		name, replacement string
		wantRefused       bool
	}{
		{"a reword is refused",
			"Must the deck keep Sanguine Bond, or may I cut other cards that do not fit the lifegain plan?", true},
		{"an added format word is refused",
			"Must the deck keep Sanguine Bond, or may I cut a card that does not fit the tempo plan?", true},
		{"a real question passes",
			"Which cards from your collection should I treat as locked in?", false},
		{"a different subject passes",
			"Just to confirm, is the tournament using Modern format?", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := nearCopy(tc.replacement, row); got != tc.wantRefused {
				t.Errorf("nearCopy = %v, want %v (overlap %.2f)", got, tc.wantRefused, overlap(tc.replacement, row))
			}
		})
	}
}

// TestRewordKeepsTheCatalogQuestion runs the refusal through a turn.
func TestRewordKeepsTheCatalogQuestion(t *testing.T) {
	out := commanderClassify()
	catalog := "Which power bracket should the deck target? 2 is the core level, near a precon, 3 is upgraded, 4 is high power."
	a, _ := testAgent(t,
		classifyStep(t, out),
		scoreStep(t,
			scored{RowID: "power_commander", Fit: 0.20, Reason: "test",
				CustomText: "Which power bracket should the white-black deck target? 2 is the core level, near a precon, 3 is upgraded, 4 is high power."},
			scored{RowID: "commander", Fit: 0.9, Reason: "fits"}),
		askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "lifegain commander deck", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	q := question(res.Questions, "power")
	if q == nil {
		t.Fatal("the bracket question did not fire")
	}
	if q.GetInvented() {
		t.Errorf("a reword was counted as an invention: %q", q.GetText())
	}
	if q.GetText() != catalog {
		t.Errorf("the catalog text did not go out: %q", q.GetText())
	}
	if res.Coverage.NearCopies != 1 {
		t.Errorf("near copies = %d, want 1", res.Coverage.NearCopies)
	}
	if res.Invented != 0 {
		t.Errorf("invented = %d, want 0", res.Invented)
	}
}

// TestFormatWordIsNormalized is D-92. The classifier writes the format in
// free text now, because a schema enum suppressed the field: the model
// answered "unknown" for a message that named the format outright. Go
// owns the vocabulary, so it must accept what the model actually writes.
func TestFormatWordIsNormalized(t *testing.T) {
	cases := []struct {
		word string
		want mtgv1.FormatId
	}{
		{"Commander", mtgv1.FormatId_FORMAT_ID_COMMANDER},
		{"commander", mtgv1.FormatId_FORMAT_ID_COMMANDER},
		{"  MODERN ", mtgv1.FormatId_FORMAT_ID_MODERN},
		{"Standard", mtgv1.FormatId_FORMAT_ID_STANDARD},
		// D-155 removed these four. The classifier must not resolve them,
		// so the unsupported-format row can decline them instead.
		{"Pauper", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
		{"Pioneer", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
		{"Legacy", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
		{"Vintage", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
		{"unknown", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
		{"", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
		{"nonsense", mtgv1.FormatId_FORMAT_ID_UNSPECIFIED},
	}
	for _, tc := range cases {
		word, want := tc.word, tc.want
		t.Run(word, func(t *testing.T) {
			out := classifyOut{Format: word, PoolRule: "unknown"}
			a, _ := testAgent(t, classifyStep(t, out), fits(t), askStep(t))
			st := NewState(false)
			if _, err := a.Turn(context.Background(), st, "a deck please", nil); err != nil {
				t.Fatalf("turn: %v", err)
			}
			if got := st.Slots.GetFormat().GetId(); got != want {
				t.Errorf("format %q gave %v, want %v", word, got, want)
			}
			if want != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED && st.Ctx.Format != want {
				t.Errorf("the planner did not learn the format from %q", word)
			}
		})
	}
}

// TestPoolRuleWordIsNormalized covers the other field that lost its enum.
func TestPoolRuleWordIsNormalized(t *testing.T) {
	cases := []struct {
		word string
		want mtgv1.PoolRule
	}{
		{"owned_first", mtgv1.PoolRule_POOL_RULE_OWNED_FIRST},
		{"Owned First", mtgv1.PoolRule_POOL_RULE_OWNED_FIRST},
		{"any-card", mtgv1.PoolRule_POOL_RULE_ANY_CARD},
		{"OWNED_ONLY", mtgv1.PoolRule_POOL_RULE_OWNED_ONLY},
		{"unknown", mtgv1.PoolRule_POOL_RULE_UNSPECIFIED},
		{"nonsense", mtgv1.PoolRule_POOL_RULE_UNSPECIFIED},
	}
	for _, tc := range cases {
		word, want := tc.word, tc.want
		t.Run(word, func(t *testing.T) {
			out := classifyOut{Format: "unknown", PoolRule: word}
			a, _ := testAgent(t, classifyStep(t, out), fits(t), askStep(t))
			st := NewState(true)
			if _, err := a.Turn(context.Background(), st, "a deck please", nil); err != nil {
				t.Fatalf("turn: %v", err)
			}
			if got := st.Slots.GetPoolRule(); got != want {
				t.Errorf("pool rule %q gave %v, want %v", word, got, want)
			}
		})
	}
}

// TestDeclineClosesASlot is D-93. A user who answers "any colors are
// fine" has handed the choice back, and the color slot must close.
func TestDeclineClosesASlot(t *testing.T) {
	first := classifyOut{Format: "commander", Theme: "sacrifice", PoolRule: "unknown"}
	second := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second.DeclinedKeys = []string{"colors"}

	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "power_commander", "colors"), askStep(t),
		classifyStep(t, second), fits(t), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "something fun and janky", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if st.Slots.GetSlotStates()["colors"] != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Fatalf("the color question is not out: %v", st.Slots.GetSlotStates())
	}
	if _, err := a.Turn(context.Background(), st, "any colors are fine", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !st.Ctx.Filled["colors"] {
		t.Error("a decline left the color slot open, so the agent never finishes")
	}
	// Skipped, not filled. The generator reads the difference and applies
	// the default the corpus names.
	if got := st.Slots.GetSlotStates()["colors"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("color state = %v, want SKIPPED", got)
	}
	if len(st.Slots.GetColors()) != 0 {
		t.Error("a decline invented a value, and the user chose none")
	}
	// The planner must not raise it again.
	for _, r := range load(t).Plan(st.Ctx) {
		if r.Slot == "colors" {
			t.Errorf("the agent plans the color question again after a decline: %s", r.ID)
		}
	}
}

// TestDeclineNeedsAnOpenQuestion keeps the D-83 hazard closed. Nobody can
// decline a question they never saw, and a classifier that names an
// unasked key must not empty the session.
func TestDeclineNeedsAnOpenQuestion(t *testing.T) {
	out := commanderClassify()
	out.PoolRule = "unknown"
	out.DeclinedKeys = []string{"power", "pool_rule", "budget"}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "power_commander"), askStep(t))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "lifegain commander deck", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	for _, key := range []string{"power", "pool_rule", "budget"} {
		if st.Ctx.Filled[key] {
			t.Errorf("key %q was declined although its question was never asked", key)
		}
	}
}

// TestDeclineRecordsTheAskAsAnswered keeps the M-4 report honest. A
// declined question did its work: the slot is closed.
func TestDeclineRecordsTheAskAsAnswered(t *testing.T) {
	first := classifyOut{Format: "commander", Theme: "sacrifice", PoolRule: "unknown"}
	second := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second.DeclinedKeys = []string{"colors"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "power_commander", "colors"), askStep(t),
		classifyStep(t, second), fits(t), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "something fun and janky", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "any colors are fine", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	var found bool
	for _, ask := range st.Asks {
		if ask.Key == "colors" {
			found = true
			if !ask.Filled {
				t.Error("the declined question is recorded as unanswered")
			}
		}
	}
	if !found {
		t.Fatal("no M-4 record for the color question")
	}
	if res.Coverage.CatalogFilled == 0 {
		t.Error("the coverage report counts no closed question")
	}
}

// TestDeclinedFormatTakesTheDefault is D-98. A declined format left
// empty gives the planner nothing to route on: every power row triggers
// on the format, so none could fire, and the session would finish with
// no power level.
func TestDeclinedFormatTakesTheDefault(t *testing.T) {
	first := classifyOut{Format: "unknown", PoolRule: "unknown"}
	declined := classifyOut{Format: "unknown", PoolRule: "unknown"}
	declined.DeclinedKeys = []string{"format"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "format", "theme", "colors"), askStep(t),
		classifyStep(t, declined), fits(t, "theme", "colors", "power_commander"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "make me a good deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "i dunno, you pick", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	// The planner can route again.
	if st.Ctx.Format != DefaultFormat {
		t.Errorf("the planner format is %v, want the corpus default", st.Ctx.Format)
	}
	if st.Slots.GetFormat().GetId() != DefaultFormat {
		t.Errorf("the slot holds %v, want the corpus default", st.Slots.GetFormat().GetId())
	}
	// Skipped, not filled: the user chose nothing, and PR-8 must know.
	if got := st.Slots.GetSlotStates()["format"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("format state = %v, want SKIPPED", got)
	}
	// The power question can now fire, which is the whole point.
	if question(res.Questions, "power") == nil {
		t.Errorf("no power question after the format was declined: %+v", res.Questions)
	}
}

// TestOutOfScopeAsksOneThing is D-99. A Yu-Gi-Oh request must get a
// decline and not "Which Yu-Gi-Oh format would you like?".
func TestOutOfScopeAsksOneThing(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	out.Facts.OutOfScope = true
	// A fixed row is neither scored nor phrased, so the turn costs one
	// call.
	a, sc := testAgent(t, classifyStep(t, out))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "can you build me a yu-gi-oh deck", nil)
	if len(sc.Calls) != 1 {
		t.Errorf("a fixed-only turn made %d model calls, want 1", len(sc.Calls))
	}
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(res.Questions) != 1 {
		t.Fatalf("an out-of-scope request got %d questions, want one: %+v", len(res.Questions), res.Questions)
	}
	q := res.Questions[0]
	if q.GetSlot() != "scope" {
		t.Errorf("question slot = %q, want scope", q.GetSlot())
	}
	if !strings.Contains(q.GetText(), "Magic") {
		t.Errorf("the question does not say what the app builds: %q", q.GetText())
	}
}

// TestBackInScopeAsksNormally proves the fact is not sticky. A user who
// asks for a Magic deck after the agent declines gets the usual flow.
func TestBackInScopeAsksNormally(t *testing.T) {
	first := classifyOut{Format: "unknown", PoolRule: "unknown"}
	first.Facts.OutOfScope = true
	second := classifyOut{Format: "commander", Theme: "dragons", PoolRule: "unknown"}
	// Turn 1 needs no score or ask step. The out-of-scope row states what
	// this app builds, so it goes out as written and reaches no model.
	a, _ := testAgent(t,
		classifyStep(t, first),
		classifyStep(t, second), fits(t, "colors", "commander", "power_commander"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "build me a yu-gi-oh deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "fine, magic then. commander, dragons", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if st.Ctx.OutOfScope {
		t.Error("the out-of-scope fact stuck after the user came back in scope")
	}
	if len(res.Questions) == 0 {
		t.Fatal("the agent asked nothing after the user came back in scope")
	}
	for _, q := range res.Questions {
		if q.GetSlot() == "scope" {
			t.Errorf("the agent declined again: %q", q.GetText())
		}
	}
}

// TestTruncationIsRefused is D-103. Overlap is symmetric, so a
// replacement that deletes half the row scores low and used to pass,
// although it says strictly less: the row explains what the answer is
// for, and the replacement drops that sentence.
func TestTruncationIsRefused(t *testing.T) {
	row := "What do people play at your event? I tune the 15 sideboard cards to it."
	cases := []struct {
		name, replacement string
		wantRefused       bool
	}{
		{"a truncation is refused",
			"What do people play at your Modern event?", true},
		{"a strict subset is refused",
			"What do people play at your event?", true},
		{"an addition is still refused as a reword",
			"What do people play at your Modern event? I tune the 15 sideboard cards to it, exactly.", true},
		{"a different question passes",
			"Which decks beat you most often last season?", false},
		{"a question about another subject passes",
			"How much are you willing to spend?", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := nearCopy(tc.replacement, row)
			if got != tc.wantRefused {
				borrowed, covered := share(tc.replacement, row)
				t.Errorf("nearCopy = %v, want %v (overlap %.2f, borrowed %.2f, covered %.2f)",
					got, tc.wantRefused, overlap(tc.replacement, row), borrowed, covered)
			}
		})
	}
}

// TestTruncationKeepsARealQuestion is the guard on the guard. A short
// replacement is not a truncation when it says something the row does
// not.
func TestTruncationKeepsARealQuestion(t *testing.T) {
	row := "Must the deck keep Sanguine Bond, or may I cut a card that does not fit the plan?"
	for _, replacement := range []string{
		"Which cards from your collection should I treat as locked in?",
		"How many cards do you want to lock in?",
	} {
		if nearCopy(replacement, row) {
			borrowed, covered := share(replacement, row)
			t.Errorf("a real question was refused: %q (borrowed %.2f, covered %.2f)",
				replacement, borrowed, covered)
		}
	}
}
