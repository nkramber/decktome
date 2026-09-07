package questions

import (
	"io"
	"log/slog"
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// leaderHints answers the two card facts the quality rules read.
type leaderHints struct {
	Hints
	// canLead and onlyCommander are keyed by card name. A name in
	// neither map is one the index does not hold.
	canLead       map[string]bool
	onlyCommander map[string]bool
	known         map[string]bool
}

func (h leaderHints) CanLead(name string) (bool, bool) {
	return h.canLead[name], h.known[name]
}

func (h leaderHints) OnlyCommander(name string) (bool, bool) {
	return h.onlyCommander[name], h.known[name]
}

func qualityAgent(t *testing.T) (*Agent, *State) {
	t.Helper()
	cat, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	h := leaderHints{
		canLead:       map[string]bool{"Atraxa, Praetor's Voice": true, "Sheoldred, the Apocalypse": true},
		onlyCommander: map[string]bool{"Atraxa, Praetor's Voice": true},
		known:         map[string]bool{"Atraxa, Praetor's Voice": true, "Sheoldred, the Apocalypse": true, "Lightning Bolt": true},
	}
	a := &Agent{cat: cat, threshold: DefaultFitThreshold, hints: h,
		log: slog.New(slog.NewTextHandler(io.Discard, nil))}
	return a, NewState(false)
}

// TestABudgetOfZeroAnswersTheRow is D-387. The phrase list held only the
// unlimited half, so a reader who said "I can not spend anything" was
// asked for a budget anyway. Eval run 31 flagged it.
func TestABudgetOfZeroAnswersTheRow(t *testing.T) {
	for _, tc := range []struct {
		message string
		want    bool
	}{
		{"I can not spend anything", true},
		{"I cannot spend anything on this", true},
		{"zero budget please", true},
		{"I have no money for cards", true},
		{"money is no object", true},
		{"spend nothing", true},
		// A real number is an answer with a value, and another rule
		// reads it. This rule must not take it.
		{"about 50 dollars", false},
		{"I want a lifegain deck", false},
	} {
		if got := noSpendingLimit(tc.message); got != tc.want {
			t.Errorf("noSpendingLimit(%q) = %v, want %v", tc.message, got, tc.want)
		}
	}
}

// TestADelegationClosesTheColors is D-388. "Surprise me" hands back
// every open key. The classifier read the commander half of it and
// missed the colors, so the color row asked anyway.
func TestADelegationClosesTheColors(t *testing.T) {
	a, st := qualityAgent(t)
	a.applyWords(st, turnWords{Message: "a commander deck, surprise me"})
	if !st.Ctx.Filled["colors"] {
		t.Error("a delegation did not close the color slot")
	}
	if st.Slots.GetSlotStates()["colors"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("colors = %v, want SKIPPED", st.Slots.GetSlotStates()["colors"])
	}
}

// TestColorsSurviveAWhateverRequest is the judge suggestion this repo
// refuses. The corpus routes "whatever" nowhere, after gate runs 11 to
// 13 read "whatever is winning" as house rules six times (D-111).
func TestColorsSurviveAWhateverRequest(t *testing.T) {
	a, st := qualityAgent(t)
	a.applyWords(st, turnWords{Message: "build whatever is winning right now"})
	if st.Ctx.Filled["colors"] {
		t.Error("\"whatever is winning\" closed the color slot, and D-111 routes it nowhere")
	}
}

// TestAColorAnswerStillWins: a delegation must not close a slot the
// reader already answered with a value.
func TestAColorAnswerStillWins(t *testing.T) {
	a, st := qualityAgent(t)
	st.Slots.Colors = []mtgv1.Color{mtgv1.Color_COLOR_W}
	a.applyWords(st, turnWords{Message: "you decide the rest"})
	if st.Slots.GetSlotStates()["colors"] == mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Error("a delegation skipped a color slot that already holds a value")
	}
}

// TestTheFormatComesFromACommanderOnlyCard is D-388. A legendary
// creature names no format on its own, and one that plays in neither
// other format this app builds leaves Commander alone.
func TestTheFormatComesFromACommanderOnlyCard(t *testing.T) {
	a, st := qualityAgent(t)
	st.NamedCards = []string{"Atraxa, Praetor's Voice"}
	a.applyWords(st, turnWords{Message: "build around atraxa praetors voice"})
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("format = %v, want Commander", st.Ctx.Format)
	}
}

// TestACardThatPlaysElsewhereNamesNoFormat holds the other half of the
// rule. Sheoldred leads a Commander deck and plays in Standard, so it
// proves nothing about the format.
func TestACardThatPlaysElsewhereNamesNoFormat(t *testing.T) {
	a, st := qualityAgent(t)
	st.NamedCards = []string{"Sheoldred, the Apocalypse"}
	a.applyWords(st, turnWords{Message: "build around sheoldred the apocalypse"})
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		t.Errorf("format = %v, want none", st.Ctx.Format)
	}
}

// TestAnUnsupportedFormatBeatsTheCard is D-112. A format this app does
// not build is declined by its own row, and a named card must never
// override that.
func TestAnUnsupportedFormatBeatsTheCard(t *testing.T) {
	a, st := qualityAgent(t)
	st.NamedCards = []string{"Atraxa, Praetor's Voice"}
	a.applyWords(st, turnWords{Message: "a Brawl deck with atraxa"})
	if st.Ctx.Format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Error("a named card set the format on a message that names one this app does not build")
	}
}

// TestTheColorRowWaitsForTheRoleOfANamedLeader is D-388. Atraxa fixes
// the deck's color identity when it leads, so a color question before
// the role is settled asks for a value the commander decides.
func TestTheColorRowWaitsForTheRoleOfANamedLeader(t *testing.T) {
	c, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	ctx := Context{Filled: map[string]bool{}, Asked: map[string]bool{}, Outstanding: map[string]string{}}
	ctx.NamedLeader, ctx.NamedCard = true, true
	if slices.Contains(ids(c.Plan(ctx)), "colors") {
		t.Error("the color row fired while a named leader had no role")
	}
	// The role settles, and the row asks again.
	ctx.NamedLeader = false
	if !slices.Contains(ids(c.Plan(ctx)), "colors") {
		t.Error("the color row did not fire after the role was settled")
	}
}

// TestTheFormatRowDropsCommanderForASixtyCardDeck is D-388. Commander is
// a 100-card format, so it can not answer a 60-card request. Eval run 31
// called that question inaccurate.
func TestTheFormatRowDropsCommanderForASixtyCardDeck(t *testing.T) {
	c, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	row, ok := c.Row("format")
	if !ok {
		t.Fatal("no format row")
	}
	st := NewState(false)
	st.Ctx.Words = "i want a 60 card anything goes deck"
	_, opts, _ := resolve(row, st, nil)
	if slices.Contains(opts, "Commander") {
		t.Errorf("options = %v, want no Commander for a 60-card request", opts)
	}
	if !slices.Contains(opts, "Standard") || !slices.Contains(opts, "Modern") {
		t.Errorf("options = %v, want the two 60-card formats", opts)
	}
	// Every other reader still gets all three.
	plain := NewState(false)
	_, all, _ := resolve(row, plain, nil)
	if !slices.Contains(all, "Commander") {
		t.Errorf("options = %v, want Commander for a reader who named no size", all)
	}
}

// TestADelegationAboutTheCommanderLeavesTheColors holds the other half
// of the rule. "You pick" against the pick row chooses a commander, and
// it says nothing about the colors.
func TestADelegationAboutTheCommanderLeavesTheColors(t *testing.T) {
	a, st := qualityAgent(t)
	a.applyWords(st, turnWords{
		Message: "you pick", Declined: []string{"commander_pick"}, Open: []string{"commander_pick"},
	})
	if st.Ctx.Filled["colors"] {
		t.Error("a delegation the classifier read as a commander answer closed the colors")
	}
}
