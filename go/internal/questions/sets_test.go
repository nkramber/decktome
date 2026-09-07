package questions

import (
	"io"
	"log/slog"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// fakeSets answers the SetResolver contract from a fixed table.
type fakeSets struct {
	Hints
	one   map[string][]string
	many  map[string][]string
	group map[string][]string
}

func (f fakeSets) ResolveSetGroup(phrase string) (codes, names []string, ok bool) {
	if c, hit := f.group[phrase]; hit {
		return c, c, true
	}
	return nil, nil, false
}

func (f fakeSets) ResolveSet(phrase string) (codes, names, options []string, ok bool) {
	if c, hit := f.one[phrase]; hit {
		return c, c, nil, true
	}
	if o, hit := f.many[phrase]; hit {
		return nil, nil, o, false
	}
	return nil, nil, nil, false
}

func setAgent(t *testing.T) (*Agent, *State) {
	t.Helper()
	cat, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	a := &Agent{cat: cat, threshold: DefaultFitThreshold, log: slog.New(slog.NewTextHandler(io.Discard, nil)), hints: fakeSets{
		one:   map[string][]string{"the Hobbit set": {"hob", "hoc"}, "Bloomburrow": {"blb", "blc"}},
		many:  map[string][]string{"Tarkir": {"Tarkir: Dragonstorm", "Dragons of Tarkir"}},
		group: map[string][]string{"Marvel": {"msc", "msh", "spe", "spm"}},
	}}
	return a, NewState(false)
}

// TestApplySetsFillsTheSlot is D-376: a phrase that names one set family
// closes the set row and writes the codes.
func TestApplySetsFillsTheSlot(t *testing.T) {
	a, st := setAgent(t)
	a.applySets(st, classifyOut{SetNames: []string{"the Hobbit set"}})
	if got := st.Slots.GetSetCodes(); !slices.Equal(got, []string{"hob", "hoc"}) {
		t.Errorf("set codes = %v, want hob hoc", got)
	}
	if !st.Ctx.SetLimited {
		t.Error("the planner does not know the deck is limited to a set")
	}
	if st.Slots.GetSlotStates()[SlotSet] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Error("the set slot did not close")
	}
	if st.Ctx.SetUnresolved {
		t.Error("a resolved phrase must leave nothing to ask")
	}
}

// TestApplySetsAsksAboutAnAmbiguousName is D-376: "Tarkir" names three
// base sets, so the row asks and names them.
func TestApplySetsAsksAboutAnAmbiguousName(t *testing.T) {
	a, st := setAgent(t)
	a.applySets(st, classifyOut{SetNames: []string{"Tarkir"}})
	if !st.Ctx.SetUnresolved {
		t.Fatal("an ambiguous set name must open the row")
	}
	if st.UnresolvedSet != "Tarkir" {
		t.Errorf("UnresolvedSet = %q", st.UnresolvedSet)
	}
	if len(st.SetOptions) != 2 {
		t.Errorf("options = %v, want the two sets", st.SetOptions)
	}
	if len(st.Slots.GetSetCodes()) != 0 {
		t.Error("an unresolved name must set no code")
	}
	// The row fires, and it names the phrase and the sets.
	rows := a.cat.Plan(st.Ctx)
	var found bool
	for _, r := range rows {
		if r.ID == "set_unresolved" {
			found = true
			text, opts, _ := resolve(r, st, nil)
			if text == r.Fallback {
				t.Fatalf("the row fell back although it had every value: %q", text)
			}
			for _, want := range []string{"Tarkir", "Tarkir: Dragonstorm", "Dragons of Tarkir"} {
				if !strings.Contains(text, want) {
					t.Errorf("the question does not name %q: %s", want, text)
				}
			}
			// The reader picks a set, or drops the limit.
			if len(opts) != 3 || opts[2] != everySetOption {
				t.Errorf("options = %v, want the two sets and a way out", opts)
			}
		}
	}
	if !found {
		t.Fatalf("the set row did not fire: %v", ids(rows))
	}
}

// TestApplySetsKeepsBothHalves is the hole a diff review found. A
// message that names two sets and resolves one must apply that one and
// still ask about the other. Dropping the second in silence is the
// defect this project keeps fixing (D-376).
func TestApplySetsKeepsBothHalves(t *testing.T) {
	a, st := setAgent(t)
	a.applySets(st, classifyOut{SetNames: []string{"Bloomburrow", "Tarkir"}})
	if got := st.Slots.GetSetCodes(); !slices.Equal(got, []string{"blb", "blc"}) {
		t.Errorf("set codes = %v, want the resolved half", got)
	}
	if !st.Ctx.SetUnresolved || st.UnresolvedSet != "Tarkir" {
		t.Fatalf("the unresolved half was dropped: %q", st.UnresolvedSet)
	}
	// The set slot is filled and the ask row still fires, because the row
	// carries its own key.
	var found bool
	for _, r := range a.cat.Plan(st.Ctx) {
		if r.ID == "set_unresolved" {
			found = true
		}
	}
	if !found {
		t.Fatal("the row did not fire beside a filled set slot")
	}
	// The answer arrives, and it joins the sets already named.
	a.applySets(st, classifyOut{SetNames: []string{"the Hobbit set"}})
	if got := st.Slots.GetSetCodes(); !slices.Equal(got, []string{"blb", "blc", "hob", "hoc"}) {
		t.Errorf("set codes after the answer = %v, want both families", got)
	}
	if st.Ctx.SetUnresolved {
		t.Error("the row must close when every phrase names a set")
	}
}

// TestApplySetsIgnoresAnEmptyList: a message that names no set changes
// nothing, so a set the reader gave before stays.
func TestApplySetsIgnoresAnEmptyList(t *testing.T) {
	a, st := setAgent(t)
	a.applySets(st, classifyOut{SetNames: []string{"the Hobbit set"}})
	a.applySets(st, classifyOut{})
	if got := st.Slots.GetSetCodes(); !slices.Equal(got, []string{"hob", "hoc"}) {
		t.Errorf("set codes = %v, want the sets to survive a message about something else", got)
	}
}

// TestManaPermissionReadsTheAnswer is D-382: only a yes fills the key,
// and every other answer leaves the mana base inside the sets.
func TestManaPermissionReadsTheAnswer(t *testing.T) {
	for _, tc := range []struct {
		answer string
		want   mtgv1.SlotState
	}{
		{"Yes, mana from any set", mtgv1.SlotState_SLOT_STATE_FILLED},
		{"yes please", mtgv1.SlotState_SLOT_STATE_FILLED},
		{"No, only these sets", mtgv1.SlotState_SLOT_STATE_SKIPPED},
		{"no", mtgv1.SlotState_SLOT_STATE_SKIPPED},
		// A message about something else leaves the question out.
		{"bracket 3", mtgv1.SlotState_SLOT_STATE_ASKED},
	} {
		t.Run(tc.answer, func(t *testing.T) {
			a, st := setAgent(t)
			row, ok := a.cat.Row("set_outside_mana")
			if !ok {
				t.Fatal("no set_outside_mana row")
			}
			st.MarkAsked(row.ID, row.StateKey(), row.Slot)
			st.Asks = append(st.Asks, Ask{QuestionID: "q1", RowID: row.ID, Key: row.StateKey(), Slot: row.Slot})
			a.applyManaPermission(st, tc.answer)
			if got := st.Slots.GetSlotStates()[SlotSetOutsideMana]; got != tc.want {
				t.Errorf("state = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestTheSetRowAsksOncePerName is the D-210 rule for the set row. Gate
// run 29 showed it asking about Tarkir on two turns in a row, because
// nothing recorded the phrase it had already named.
func TestTheSetRowAsksOncePerName(t *testing.T) {
	a, st := setAgent(t)
	a.applySets(st, classifyOut{SetNames: []string{"Tarkir"}})
	st.Ctx.SetChanged = st.BadSetChanged()
	if !st.Ctx.SetChanged {
		t.Fatal("a new set name must let the row ask")
	}
	// The row goes out and records the phrase.
	st.RecordAskedSet()
	st.Ctx.Asked["set_unresolved"] = true
	st.Ctx.SetChanged = st.BadSetChanged()
	if st.Ctx.SetChanged {
		t.Error("the same set name must not read as a change")
	}
	for _, r := range a.cat.Plan(st.Ctx) {
		if r.ID == "set_unresolved" {
			t.Error("the row asked twice about one name")
		}
	}
	// A second name the app can not settle asks again.
	st.SetUnresolved("Kamigawa", []string{"Kamigawa: Neon Dynasty", "Champions of Kamigawa"})
	st.Ctx.SetChanged = st.BadSetChanged()
	var found bool
	for _, r := range a.cat.Plan(st.Ctx) {
		if r.ID == "set_unresolved" {
			found = true
		}
	}
	if !found {
		t.Error("a new name the app can not settle must ask again")
	}
}

// TestTheTurnCarriesTheSetsItApplied is D-390. The turn names the sets
// it read, so the chat can tell the reader. It is turn state: the next
// turn starts with it clear, the way the commander mark of D-366 does.
func TestTheTurnCarriesTheSetsItApplied(t *testing.T) {
	a, st := setAgent(t)
	a.applySets(st, classifyOut{SetNames: []string{"the Hobbit set"}})
	if !slices.Equal(st.setsThisTurn, []string{"hob", "hoc"}) {
		t.Errorf("setsThisTurn = %v, want the two set names", st.setsThisTurn)
	}
	// A message that names no set carries nothing, so the chat says
	// nothing twice.
	st.setsThisTurn = nil
	a.applySets(st, classifyOut{})
	if len(st.setsThisTurn) != 0 {
		t.Errorf("setsThisTurn = %v, want none on a turn that named no set", st.setsThisTurn)
	}
	// The set limit itself survives, so only the note is per turn.
	if !slices.Equal(st.Slots.GetSetCodes(), []string{"hob", "hoc"}) {
		t.Errorf("set codes = %v, want the limit to survive", st.Slots.GetSetCodes())
	}
}

// TestApplySetsReadsAGroup is D-525: a franchise word the classifier
// marks as a group reaches every family, and the set row closes. An
// unknown group asks the set row, as an unknown name does.
func TestApplySetsReadsAGroup(t *testing.T) {
	a, st := setAgent(t)
	a.applySets(st, classifyOut{SetGroups: []string{"Marvel"}})
	if got := st.Slots.GetSetCodes(); !slices.Equal(got, []string{"msc", "msh", "spe", "spm"}) {
		t.Errorf("set codes = %v, want every Marvel family", got)
	}
	if st.Slots.GetSlotStates()[SlotSet] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("the set row is not filled: %v", st.Slots.GetSlotStates())
	}
	a2, st2 := setAgent(t)
	a2.applySets(st2, classifyOut{SetGroups: []string{"Star Wars"}})
	if st2.UnresolvedSet != "Star Wars" || len(st2.SetOptions) != 0 {
		t.Errorf("an unknown group must ask with no option: %q %v", st2.UnresolvedSet, st2.SetOptions)
	}
}
