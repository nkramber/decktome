package questions

import (
	"io"
	"log/slog"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// fakePrecons answers the PreconResolver and the OwnedPreconSource
// contracts from a fixed table.
type fakePrecons struct {
	Hints
	one     map[string]PreconMatch
	options map[string][]string
	owned   []PreconRef
	// noCollection makes OwnedPrecons answer as a session with no
	// collection does.
	noCollection bool
}

func (f fakePrecons) ResolvePrecon(phrase string) PreconMatch {
	if m, hit := f.one[phrase]; hit {
		return m
	}
	return PreconMatch{Options: f.options[phrase]}
}

func (f fakePrecons) OwnedPrecons() ([]PreconRef, bool) {
	if f.noCollection {
		return nil, false
	}
	return f.owned, true
}

var (
	avengers = PreconRef{Key: "AvengersAssemble_MSC", Name: "Avengers Assemble"}
	turtles  = PreconRef{Key: "TurtlePower_TMC", Name: "Turtle Power!"}
	blight   = PreconRef{Key: "BlightCurse_ECC", Name: "Blight Curse"}
)

func preconAgent(t *testing.T, f fakePrecons) (*Agent, *State) {
	t.Helper()
	cat, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if f.one == nil {
		f.one = map[string]PreconMatch{
			"Avengers Assemble": {Products: []PreconRef{avengers}, OK: true},
			"my Turtle Power precon": {Products: []PreconRef{turtles}, OK: true,
				Partial: []string{"Turtle Power!"}},
		}
	}
	if f.options == nil {
		f.options = map[string][]string{"Deck A": {"Deck A (AAA)", "Deck A (BBB)"}}
	}
	a := &Agent{cat: cat, threshold: DefaultFitThreshold, log: slog.New(slog.NewTextHandler(io.Discard, nil)), hints: f}
	return a, NewState(true)
}

// TestApplyPreconsFillsTheSlot is D-496: a named precon closes the precon
// row and writes the product key, and the turn names the product.
func TestApplyPreconsFillsTheSlot(t *testing.T) {
	a, st := preconAgent(t, fakePrecons{})
	a.applyPrecons(st, classifyOut{PreconNames: []string{"Avengers Assemble"}})
	if got := st.Slots.GetExcludePreconKeys(); !slices.Equal(got, []string{avengers.Key}) {
		t.Errorf("keys = %v, want the Avengers Assemble key", got)
	}
	if !st.Ctx.PreconsExcluded {
		t.Error("the planner does not know the deck excludes a precon")
	}
	if st.Slots.GetSlotStates()[SlotPrecons] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Error("the precon slot did not close")
	}
	if !slices.Equal(st.preconsThisTurn, []string{"Avengers Assemble"}) {
		t.Errorf("preconsThisTurn = %v", st.preconsThisTurn)
	}
	if len(st.preconsPartialThisTurn) != 0 || st.preconsNoneThisTurn {
		t.Error("a product the collection holds whole is neither partial nor none")
	}
	if st.Ctx.PreconUnresolved {
		t.Error("a resolved phrase must leave nothing to ask")
	}
	// A second turn that names another product adds it, and the first
	// survives.
	a.applyPrecons(st, classifyOut{PreconNames: []string{"my Turtle Power precon"}})
	if got := st.Slots.GetExcludePreconKeys(); !slices.Equal(got, []string{turtles.Key, avengers.Key}) {
		t.Errorf("keys = %v, want both products", got)
	}
	if !slices.Equal(st.preconsThisTurn, []string{"Turtle Power!", "Avengers Assemble"}) {
		t.Errorf("preconsThisTurn = %v, want every excluded product", st.preconsThisTurn)
	}
}

// TestApplyPreconsMarksAPartialProduct is D-497: a named product the
// collection does not hold whole is excluded anyway, and the turn says so.
func TestApplyPreconsMarksAPartialProduct(t *testing.T) {
	a, st := preconAgent(t, fakePrecons{})
	a.applyPrecons(st, classifyOut{PreconNames: []string{"my Turtle Power precon"}})
	if got := st.Slots.GetExcludePreconKeys(); !slices.Equal(got, []string{turtles.Key}) {
		t.Errorf("keys = %v, want the product excluded anyway", got)
	}
	if !slices.Equal(st.preconsPartialThisTurn, []string{"Turtle Power!"}) {
		t.Errorf("partial = %v", st.preconsPartialThisTurn)
	}
}

// TestApplyPreconsReadsTheOwnedPrecons is D-408: "not from my precons"
// excludes every product the collection holds whole.
func TestApplyPreconsReadsTheOwnedPrecons(t *testing.T) {
	out := classifyOut{}
	out.Facts.ExcludePrecons = true
	a, st := preconAgent(t, fakePrecons{owned: []PreconRef{avengers, blight}})
	a.applyPrecons(st, out)
	if got := st.Slots.GetExcludePreconKeys(); !slices.Equal(got, []string{avengers.Key, blight.Key}) {
		t.Errorf("keys = %v, want both owned products", got)
	}
	if st.PreconPhrase != "my precons" {
		t.Errorf("phrase = %q", st.PreconPhrase)
	}
	// A collection with no whole precon excludes nothing, and the turn
	// says so. The row still closes: the request has its answer.
	a, st = preconAgent(t, fakePrecons{})
	a.applyPrecons(st, out)
	if len(st.Slots.GetExcludePreconKeys()) != 0 || !st.preconsNoneThisTurn {
		t.Errorf("keys = %v none = %v, want none", st.Slots.GetExcludePreconKeys(), st.preconsNoneThisTurn)
	}
	if st.Slots.GetSlotStates()[SlotPrecons] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Error("the precon slot must close with nothing to exclude")
	}
	if st.Ctx.PreconsExcluded {
		t.Error("nothing is excluded, so the fact stays off")
	}
	// No collection, or no table: the turn says the check is unavailable.
	a, st = preconAgent(t, fakePrecons{noCollection: true})
	a.applyPrecons(st, out)
	if !st.preconsUnavailableThisTurn || st.preconsNoneThisTurn {
		t.Error("with no collection the turn must say the check is unavailable")
	}
}

// TestApplyPreconsAsksAboutAnUnknownName mirrors the set row (D-376): a
// phrase the table can not settle opens the precon row with its options.
func TestApplyPreconsAsksAboutAnUnknownName(t *testing.T) {
	a, st := preconAgent(t, fakePrecons{})
	a.applyPrecons(st, classifyOut{PreconNames: []string{"Deck A"}})
	if !st.Ctx.PreconUnresolved || st.UnresolvedPrecon != "Deck A" {
		t.Fatalf("an unknown name must open the row: %v %q", st.Ctx.PreconUnresolved, st.UnresolvedPrecon)
	}
	if len(st.Slots.GetExcludePreconKeys()) != 0 {
		t.Error("an unresolved name must exclude nothing")
	}
	rows := a.cat.Plan(st.Ctx)
	var found bool
	for _, r := range rows {
		if r.ID == "precon_unresolved" {
			found = true
			text, opts, _ := resolve(r, st, nil)
			if text == r.Fallback {
				t.Fatalf("the row fell back although it had every value: %q", text)
			}
			for _, want := range []string{"Deck A", "Deck A (AAA)", "Deck A (BBB)"} {
				if !strings.Contains(text, want) {
					t.Errorf("text %q lacks %q", text, want)
				}
			}
			if len(opts) != 0 {
				t.Errorf("a fixed row carries its options in the text, got %v", opts)
			}
		}
	}
	if !found {
		t.Fatalf("the precon row did not fire: %v", ids(rows))
	}
	// A message that names two products and settles one keeps the other
	// open, both halves run (D-376).
	a, st = preconAgent(t, fakePrecons{})
	a.applyPrecons(st, classifyOut{PreconNames: []string{"Avengers Assemble", "Deck A"}})
	if got := st.Slots.GetExcludePreconKeys(); !slices.Equal(got, []string{avengers.Key}) {
		t.Errorf("keys = %v, want the resolved half", got)
	}
	if !st.Ctx.PreconUnresolved {
		t.Error("the unresolved half must still ask")
	}
	// The answer names a product, and the row closes.
	a.applyPrecons(st, classifyOut{PreconNames: []string{"Avengers Assemble"}})
	if st.Ctx.PreconUnresolved || st.UnresolvedPrecon != "" {
		t.Error("a turn where every phrase resolves closes the row")
	}
}

// TestThePreconRowAsksOncePerName is the D-210 rule for the precon row.
func TestThePreconRowAsksOncePerName(t *testing.T) {
	a, st := preconAgent(t, fakePrecons{})
	a.applyPrecons(st, classifyOut{PreconNames: []string{"Deck A"}})
	st.UnresolvedPreconAsked = "Deck A"
	st.Ctx.PreconChanged = st.BadPreconChanged()
	if st.Ctx.PreconChanged {
		t.Error("the same phrase must not read as a change")
	}
	st.Ctx.Asked["precon_unresolved"] = true
	if slices.Contains(ids(a.cat.Plan(st.Ctx)), "precon_unresolved") {
		t.Error("the row asked again with the same phrase")
	}
	a.applyPrecons(st, classifyOut{PreconNames: []string{"Fluffy"}})
	st.Ctx.PreconChanged = st.BadPreconChanged()
	if !st.Ctx.PreconChanged {
		t.Error("another phrase must read as a change")
	}
}

// TestTheTurnCarriesThePreconsItApplied: the result carries the marks of
// the turn, and the next turn starts clear.
func TestTheTurnCarriesThePreconsItApplied(t *testing.T) {
	a, st := preconAgent(t, fakePrecons{})
	a.applyPrecons(st, classifyOut{PreconNames: []string{"my Turtle Power precon"}})
	var res Result
	st.fillPrecons(&res)
	if !slices.Equal(res.PreconsApplied, []string{"Turtle Power!"}) || !slices.Equal(res.PreconsPartial, []string{"Turtle Power!"}) {
		t.Errorf("result = %+v", res)
	}
	var next Result
	st.fillPrecons(&next)
	if len(next.PreconsApplied) != 0 || len(next.PreconsPartial) != 0 || next.PreconsNone || next.PreconsUnavailable {
		t.Errorf("the marks must clear after one result: %+v", next)
	}
	if !slices.Equal(st.Slots.GetExcludePreconKeys(), []string{turtles.Key}) {
		t.Error("the exclusion itself survives, only the note is per turn")
	}
}

// TestApplyPreconsWithoutAResolver: a hint source with no precon table
// says so, and excludes nothing in silence.
func TestApplyPreconsWithoutAResolver(t *testing.T) {
	cat, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	a := &Agent{cat: cat, threshold: DefaultFitThreshold, log: slog.New(slog.NewTextHandler(io.Discard, nil)), hints: nil}
	st := NewState(true)
	a.applyPrecons(st, classifyOut{PreconNames: []string{"Avengers Assemble"}})
	if !st.preconsUnavailableThisTurn {
		t.Error("with no table the turn must say the exclusion is unavailable")
	}
	if len(st.Slots.GetExcludePreconKeys()) != 0 || st.Ctx.PreconUnresolved {
		t.Error("with no table nothing resolves and nothing asks")
	}
}
