package questions

import (
	"context"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// The commander tests: the pick row, the delegation, the offer, and the
// named-card role. The pure commander rules are in words_test.go.

// TestColorChangeKeepsTheDelegation is D-153. A color change after the
// user delegated the commander choice must not reopen the pick row: the
// offer leaves, and the delegation stands.
func TestColorChangeKeepsTheDelegation(t *testing.T) {
	first := commanderClassify()
	change := classifyOut{Format: "unknown", PoolRule: "unknown", Colors: []string{"W", "G"}}
	h := &fakeHints{
		commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
		identity:   map[string][]mtgv1.Color{"Karlov of the Ghost Council": {mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, first), fits(t, "power_commander", "budget"), askStep(t),
		classifyStep(t, change))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A lifegain Commander deck, white and black, you pick the commander.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if st.Slots.GetSlotStates()["commander_pick"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Fatalf("the delegation did not close the pick row: %v", st.Slots.GetSlotStates())
	}
	if len(st.CurrentOffer) != 0 {
		t.Errorf("names stayed on the table after the delegation: %v", st.CurrentOffer)
	}
	res, err := a.Turn(context.Background(), st, "White and green instead.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if q := question(res.Questions, "commander"); q != nil {
		t.Errorf("a color change reopened the commander choice: %q", q.GetText())
	}
	if st.Slots.GetSlotStates()["commander_pick"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("commander_pick = %v after the color change, want SKIPPED", st.Slots.GetSlotStates()["commander_pick"])
	}
}

// TestColorChangeDropsTheOfferOfASetCommander is the other half of D-153.
// A named commander settles the choice, and a color change leaves the
// old offer alone.
func TestColorChangeDropsTheOfferOfASetCommander(t *testing.T) {
	a, _ := testAgentHints(t, &fakeHints{identity: map[string][]mtgv1.Color{}})
	st := NewState(false)
	st.SetOffer([]string{"Oloro, Ageless Ascetic"})
	st.SetCommander("Karlov of the Ghost Council")
	st.Slots.Colors = []mtgv1.Color{mtgv1.Color_COLOR_G}
	a.dropOffColorOffers(st)
	if len(st.CurrentOffer) != 0 {
		t.Errorf("the offer stayed after a commander was set: %v", st.CurrentOffer)
	}
	if !st.Ctx.Filled["commander_pick"] {
		t.Error("the pick row reopened under a set commander")
	}
}

// TestDelegationFollowsTheClassifier is D-147. "Any colors, you pick"
// delegates the colors, and the word rule must not hand the commander
// choice over because a commander question is out. The classifier says
// which question the message answered.
func TestDelegationFollowsTheClassifier(t *testing.T) {
	cases := []struct {
		name     string
		message  string
		declined []string
		wantSkip bool
	}{
		{"the classifier filed it under the colors", "Any colors, you pick.", []string{"colors"}, false},
		{"the classifier filed it under the commander", "Whatever, you pick.", []string{"commander"}, true},
		{"no verdict and other questions out", "Up to you.", nil, false},
		{"the message names the commander", "You pick the commander, any colors.", []string{"colors"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// A named cap closes the budget row, so the commander row sits
			// in the first three open rows (D-294).
			first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown", BudgetUSD: 50}
			second := classifyOut{Format: "unknown", PoolRule: "unknown", DeclinedKeys: tc.declined}
			a, _ := testAgent(t,
				classifyStep(t, first), fits(t, "commander", "power_commander", "colors"), askStep(t),
				classifyStep(t, second), fits(t), askStep(t))
			st := NewState(false)
			if _, err := a.Turn(context.Background(), st, "A lifegain Commander deck, 50 dollars.", nil); err != nil {
				t.Fatalf("turn 1: %v", err)
			}
			for _, key := range []string{"commander", "power", "colors"} {
				if st.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_ASKED {
					t.Fatalf("turn 1 did not ask %s: %v", key, st.Slots.GetSlotStates())
				}
			}
			if _, err := a.Turn(context.Background(), st, tc.message, nil); err != nil {
				t.Fatalf("turn 2: %v", err)
			}
			got := st.Slots.GetSlotStates()["commander"] == mtgv1.SlotState_SLOT_STATE_SKIPPED
			if got != tc.wantSkip {
				t.Errorf("commander skipped = %v, want %v: states %v", got, tc.wantSkip, st.Slots.GetSlotStates())
			}
		})
	}
}

// TestDelegationAloneClosesTheOnlyOpenQuestion keeps the D-147 case: a
// bare "up to you" answers the commander when nothing else is out.
func TestDelegationAloneClosesTheOnlyOpenQuestion(t *testing.T) {
	st := NewState(false)
	st.MarkAsked("commander", "commander", "commander")
	if !delegationIsAboutTheCommander(st, turnWords{Message: "Up to you."}) {
		t.Error("a delegation with the commander as the only open question was not read")
	}
	st.MarkAsked("colors", "colors", "colors")
	if delegationIsAboutTheCommander(st, turnWords{Message: "Up to you."}) {
		t.Error("a delegation with two questions out took the commander")
	}
	if !delegationIsAboutTheCommander(st, turnWords{Message: "Up to you.", Declined: []string{"commander_pick"}}) {
		t.Error("the classifier filed it under the pick and the rule did not read that")
	}
}

// TestNamedCardRoleSetsTheCommander is D-118 and D-83. "As my commander"
// must set the commander, "In the 99" must lock the card, and a name
// alone closes nothing.
func TestNamedCardRoleSetsTheCommander(t *testing.T) {
	cases := []struct {
		name, answer  string
		closed        []string
		wantCommander bool
		wantLocked    bool
		wantOpen      bool
	}{
		{"as my commander", "As my commander.", nil, true, false, false},
		{"in the 99", "In the 99.", nil, false, true, false},
		{"a name alone closes nothing", "sure, that one", []string{"named_card_role"}, false, false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			first := classifyOut{Format: "commander", Theme: "sacrifice", PoolRule: "unknown", NamedCards: []string{"Grist, the Hunger Tide"}}
			first.Facts.NamedCard = true
			second := classifyOut{Format: "unknown", PoolRule: "unknown", ClosedKeys: tc.closed}
			a, _ := testAgent(t,
				classifyStep(t, first), fits(t, "named_card_role", "power_commander", "colors"), askStep(t),
				classifyStep(t, second), fits(t), askStep(t))
			st := NewState(false)
			if _, err := a.Turn(context.Background(), st, "Build around Grist, the Hunger Tide. A sacrifice deck.", nil); err != nil {
				t.Fatalf("turn 1: %v", err)
			}
			if st.Slots.GetSlotStates()["named_card_role"] != mtgv1.SlotState_SLOT_STATE_ASKED {
				t.Fatalf("the role row did not fire: %v", st.Slots.GetSlotStates())
			}
			if _, err := a.Turn(context.Background(), st, tc.answer, nil); err != nil {
				t.Fatalf("turn 2: %v", err)
			}
			if got := hasName(st.CommanderNames, "Grist, the Hunger Tide"); got != tc.wantCommander {
				t.Errorf("Grist is the commander = %v, want %v", got, tc.wantCommander)
			}
			if got := hasName(st.LockedCards(), "Grist, the Hunger Tide"); got != tc.wantLocked {
				t.Errorf("Grist is locked = %v, want %v", got, tc.wantLocked)
			}
			open := st.Slots.GetSlotStates()["named_card_role"] == mtgv1.SlotState_SLOT_STATE_ASKED
			if open != tc.wantOpen {
				t.Errorf("role question open = %v, want %v", open, tc.wantOpen)
			}
		})
	}
}

// TestCardInThe99ClosesTheRoleRow is D-70. "Build around Grist, the
// Hunger Tide, but not as my commander" answers the role question, so
// the row must not ask whether Grist should be the commander.
func TestCardInThe99ClosesTheRoleRow(t *testing.T) {
	out := classifyOut{Format: "commander", PoolRule: "unknown"}
	out.LockedNames = []string{"Grist, the Hunger Tide"}
	out.Facts.NamedCard = true
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "theme", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st,
		"Build around Grist, the Hunger Tide, but not as my commander.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if !st.Ctx.Filled["named_card_role"] {
		t.Error("the role question is still open after the user answered it")
	}
	if st.Ctx.Asked["named_card_role"] {
		t.Error("the agent asked the role question the user had answered")
	}
	if q := question(res.Questions, "commander"); q != nil {
		t.Errorf("a commander question went out: %q", q.GetText())
	}
}

// TestNoneRepeatsThePickRowWithNewNames is D-73 and D-120. "None of
// those" is a refusal, and a pick row closed on it leaves a commander
// nobody chose.
func TestNoneRepeatsThePickRowWithNewNames(t *testing.T) {
	base := commanderClassify()
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	// The offer waits for the power (D-630), and the turn text names it.
	wants.Power = "bracket 3"
	// Turn 3 refuses, and the classifier tries to close the row by name.
	refuse := commanderClassify()
	// The classifier can report the refusal through both channels.
	// Neither may close the row.
	refuse.ClosedKeys = []string{"commander_pick"}
	refuse.DeclinedKeys = []string{"commander_pick"}
	h := &fakeHints{
		commanders: []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second:     []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, base), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, wants),
		classifyStep(t, refuse))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a lifegain commander deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "Bracket 3. I have no commander in mind, so suggest one.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !st.Ctx.Asked["commander_pick"] {
		t.Fatal("the pick row never went out")
	}
	res, err := a.Turn(context.Background(), st, "None of those.", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if st.Ctx.Filled["commander_pick"] {
		t.Fatal("a refusal closed the pick row, so the session ends with no commander")
	}
	q := question(res.Questions, "commander")
	if q == nil {
		t.Fatal("the pick row did not ask again after the refusal")
	}
	for _, old := range h.commanders {
		if strings.Contains(q.GetText(), old) {
			t.Errorf("the agent offered %q again after the user refused it", old)
		}
	}
	if st.Ready(a.cat) {
		t.Error("the session called itself complete with no commander chosen")
	}
}

// TestCommanderChosenByPlace is D-121. "The first of the new three is
// good" names a place, and the classifier can not map that onto a name,
// because it never sees the names.
func TestCommanderChosenByPlace(t *testing.T) {
	base := commanderClassify()
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	// The offer waits for the power (D-630).
	wants.Power = "bracket 3"
	pick := commanderClassify()
	h := &fakeHints{
		commanders: []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second:     []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, base), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, wants),
		classifyStep(t, pick))
	st := NewState(false)
	for i, msg := range []string{
		"a lifegain commander deck",
		"Bracket 3. I have no commander in mind, so suggest one.",
		"The second one is good.",
	} {
		if _, err := a.Turn(context.Background(), st, msg, nil); err != nil {
			t.Fatalf("turn %d: %v", i+1, err)
		}
	}
	if !st.Ctx.Filled["commander"] {
		t.Fatal("the commander slot stayed open after the user chose one")
	}
	if len(st.CommanderNames) == 0 || st.CommanderNames[0] != h.commanders[1] {
		t.Errorf("commander = %v, want %q", st.CommanderNames, h.commanders[1])
	}
}

// TestSuggestionDoesNotSwapTheNames is D-123, which restores D-80. The
// classifier can set wants_suggestion again on a message that refused
// nothing, and the agent must not swap the names under the user.
//
// The message asked for a suggestion before D-147. It now asks without
// the words that hand the choice over, because a delegation closes the
// pick row instead of repeating it. D-123 is about the names on the
// table, and this test still measures only that.
func TestSuggestionDoesNotSwapTheNames(t *testing.T) {
	// The offer waits for the power (D-630), so both turns name it.
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	wants.Power = "bracket 3"
	again := commanderClassify()
	again.Facts.WantsSuggestion = true
	again.Power = "bracket 3"
	h := &fakeHints{
		commanders: []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second:     []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, wants), fits(t, "commander_pick", "power_commander"), askStep(t),
		classifyStep(t, again))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a lifegain commander deck, please suggest a commander", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "Bracket 3, and build from my library first.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	// The names stay on the table. D-163 stops the second copy of the
	// same question, so the test reads the table and not a new question.
	if q := question(res.Questions, "commander"); q != nil {
		t.Errorf("the pick row asked again on a message that refused nothing: %q", q.GetText())
	}
	for _, want := range h.commanders {
		if !hasName(st.CurrentOffer, want) {
			t.Errorf("the agent dropped %q although the user refused nothing: %v", want, st.CurrentOffer)
		}
	}
}

// TestDelegationClosesTheCommanderPick is D-147. "You pick the commander"
// is a delegation. Unread, it leaves the pick row to ask again every
// turn, because the row carries "repeat": true.
//
// A delegation is a decline (D-93): the key closes, it takes no value,
// and the generator picks the best commander of the pool.
func TestDelegationClosesTheCommanderPick(t *testing.T) {
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	again := commanderClassify()
	again.Facts.WantsSuggestion = true
	h := &fakeHints{
		commanders: []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second:     []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, wants), fits(t, "power_commander"), askStep(t),
		classifyStep(t, again))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a lifegain commander deck, you pick the commander", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if st.Slots.GetSlotStates()["commander_pick"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("the delegation left commander_pick in state %v, want SKIPPED",
			st.Slots.GetSlotStates()["commander_pick"])
	}
	res, err := a.Turn(context.Background(), st, "Bracket 3, and build from my library first.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if q := question(res.Questions, "commander"); q != nil {
		t.Errorf("the pick row asked again after the user handed over the choice: %q", q.GetText())
	}
}

// TestOffColorOfferLeavesTheTable is D-153. The pick row keeps the names
// on the table until the user refuses them (D-80, D-123). Nothing checked
// them again when the colors arrived later.
//
// Turn 1 asks the colors, and turn 2 declines them, so the row offers a
// mono-green commander (D-669). The user answers "Red and white" on
// turn 3, and the same three names
// must not go out again. D-148 can not catch it: it filters the pool,
// and these names are already on the table.
func TestOffColorOfferLeavesTheTable(t *testing.T) {
	// The offer waits for the power (D-630), and it never shares a turn
	// with the color question (D-669). Turn 1 asks the colors, turn 2
	// declines them, and turn 2 makes the offer.
	first := commanderClassify()
	first.Facts.WantsSuggestion = true
	first.Colors = nil
	first.BudgetUSD = 50
	first.Power = "bracket 3"
	decline := first
	decline.DeclinedKeys = []string{"colors"}
	second := commanderClassify()
	second.Facts.WantsSuggestion = true
	second.Colors = []string{"R", "W"}
	second.Power = "bracket 3"
	// Two of the three offered names are red-white, so they survive the
	// colors. Jaheira is mono-green and must leave. A mono-red name would
	// leave as well, because D-148 asks a commander to hold every color
	// the user named, so the test uses names that isolate D-153.
	h := &fakeHints{
		commanders: []string{"Jaheira, Friend of the Forest", "Winota, Joiner of Forces", "Feather, the Redeemed"},
		second:     []string{"Aurelia, the Warleader", "Anax and Cymede", "Tajic, Blade of the Legion"},
		identity: map[string][]mtgv1.Color{
			"Jaheira, Friend of the Forest": {mtgv1.Color_COLOR_G},
			"Winota, Joiner of Forces":      {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Feather, the Redeemed":         {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Aurelia, the Warleader":        {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Anax and Cymede":               {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Tajic, Blade of the Legion":    {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
		},
	}
	// Turn 2 plans the pick row alone, which is fixed, so it makes no score
	// call and no ask call (D-131).
	a, _ := testAgentHints(t, h,
		classifyStep(t, first), fits(t, "colors"), askStep(t),
		classifyStep(t, decline),
		classifyStep(t, second), fits(t, "commander_pick"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Commander deck with a Background commander pair, 50 dollars.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if len(st.CurrentOffer) != 0 {
		t.Fatalf("turn 1 offered %v beside the color question", st.CurrentOffer)
	}
	if _, err := a.Turn(context.Background(), st, "Any colors are fine.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !contains(st.CurrentOffer, "Jaheira, Friend of the Forest") {
		t.Fatalf("turn 2 did not offer the green commander: %v", st.CurrentOffer)
	}
	if len(st.CurrentOffer) != 3 {
		t.Fatalf("turn 2 offered %d names, want 3: %v", len(st.CurrentOffer), st.CurrentOffer)
	}
	res, err := a.Turn(context.Background(), st, "Red and white, aggressive.", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if contains(st.CurrentOffer, "Jaheira, Friend of the Forest") {
		t.Errorf("the mono-green commander stayed on the table after the user named red and white: %v",
			st.CurrentOffer)
	}
	for _, keep := range []string{"Winota, Joiner of Forces", "Feather, the Redeemed"} {
		if !contains(st.CurrentOffer, keep) {
			t.Errorf("%q fits red-white and it left the table: %v", keep, st.CurrentOffer)
		}
	}
	if q := question(res.Questions, "commander"); q != nil {
		if strings.Contains(q.GetText(), "Jaheira") {
			t.Errorf("the question still names the green commander: %q", q.GetText())
		}
	}
}

// TestOffColorDropKeepsAnUnknownName guards D-153. An unknown name proves
// nothing, so the agent drops nothing on it. This is the D-140 rule for a
// card the index can not confirm.
func TestOffColorDropKeepsAnUnknownName(t *testing.T) {
	h := &fakeHints{identity: map[string][]mtgv1.Color{"Known Legend": {mtgv1.Color_COLOR_G}}}
	if fits, known := h.FitsColors("A Card Nobody Holds", []mtgv1.Color{mtgv1.Color_COLOR_R}); known || fits {
		t.Errorf("an unknown name answered fits=%v known=%v, want both false", fits, known)
	}
	// An empty color list fits everything: nothing is out of no colors.
	if fits, known := h.FitsColors("Known Legend", nil); !fits || !known {
		t.Errorf("no colors named answered fits=%v known=%v, want both true", fits, known)
	}
}

// TestHintsReadTheColorsOfThisTurn is D-124. The caller builds the hint
// source before the turn, so a color the classifier fills inside the turn
// would be invisible, and the offer could name a five-color commander for
// a red-green deck.
func TestHintsReadTheColorsOfThisTurn(t *testing.T) {
	h := &fakeHints{commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}}
	out := commanderClassify()
	out.Facts.WantsSuggestion = true
	a, _ := testAgentHints(t, h, classifyStep(t, out), fits(t, "commander_pick", "power_commander"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a white-black lifegain commander deck, you pick", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(h.saw) != 2 {
		t.Fatalf("the hint source saw %d colors, want the two the classifier filled", len(h.saw))
	}
}

// TestPickRowWithNoNamesAsksNothing is D-127. A user who answers every
// question with "you pick" leaves the theme empty, so PR-6 can name no
// commander, and the pick row must not ask with nothing to suggest.
func TestPickRowWithNoNamesAsksNothing(t *testing.T) {
	out := commanderClassify()
	out.Theme = ""
	out.Facts.WantsSuggestion = true
	out.BudgetUSD = 50
	// The offer waits for the power (D-630), so the reader names it. It
	// never shares a turn with the theme question (D-669), so turn 1 asks
	// the theme, and turn 2 declines it.
	out.Power = "bracket 3"
	decline := out
	decline.DeclinedKeys = []string{"theme"}
	// The hint source names no commander, which is what an empty theme
	// gives.
	a, _ := testAgentHints(t, &fakeHints{},
		classifyStep(t, out), fits(t, "theme"), askStep(t),
		classifyStep(t, decline))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "Make me a good deck for 50 dollars.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "I dunno, you pick.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	for _, q := range res.Questions {
		if strings.Contains(strings.ToLower(q.GetText()), "commander") {
			t.Errorf("a commander question went out with nothing to offer: %q", q.GetText())
		}
	}
	for _, key := range []string{"commander_pick", "commander"} {
		if st.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
			t.Errorf("%s = %v, want skipped so the agent chooses",
				key, st.Slots.GetSlotStates()[key])
		}
	}
}

// TestCommanderSwapReopensTheChoice is D-130. The user chooses Karlov of
// the Ghost Council, then writes "Actually use a different commander,
// suggest one". A chosen commander closes every commander row, so the
// swap must reopen the choice.
func TestCommanderSwapReopensTheChoice(t *testing.T) {
	named := commanderClassify()
	named.CommanderNames = []string{"Karlov of the Ghost Council"}
	named.Power = "bracket 3"
	swap := commanderClassify()
	swap.Facts.WantsSuggestion = true
	// The offer waits for the power (D-630).
	swap.Power = "bracket 3"
	h := &fakeHints{
		commanders: []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second:     []string{"Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim", "Liesa, Shroud of Dusk"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, named), fits(t, "power_commander"), askStep(t),
		classifyStep(t, swap))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st,
		"A white-black lifegain Commander deck. Karlov of the Ghost Council is the commander.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.CommanderSet {
		t.Fatal("the commander did not stick")
	}
	res, err := a.Turn(context.Background(), st, "Actually use a different commander, suggest one.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if st.Ctx.CommanderSet {
		t.Error("the old commander survived the swap")
	}
	q := question(res.Questions, "commander")
	if q == nil {
		t.Fatal("the agent asked nothing after the user asked for another commander")
	}
	if strings.Contains(q.GetText(), "Karlov") {
		t.Errorf("the agent offered the commander the user replaced: %q", q.GetText())
	}
}

// TestALegendaryArtifactCanNotLead is F-82, found on the deployed app on
// 2026-09-08. The reader asked to add The Arkenstone, a Legendary
// Artifact, and the app asked "Should The Arkenstone be your commander,
// or one card in the 99?".
//
// The card index derives the answer from the front face against CR
// 903.3, and CanLead threw it away: it read the type line again, saw the
// word "legendary", and said nothing. The role question then went out
// for a card that can not lead any deck.
func TestALegendaryArtifactCanNotLead(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{Name: "The Arkenstone // Seek the Heart", TypeLine: "Legendary Artifact // Sorcery — Adventure",
			Faces: []*mtgv1.CardFace{
				{Name: "The Arkenstone", TypeLine: "Legendary Artifact"},
				{Name: "Seek the Heart", TypeLine: "Sorcery — Adventure"},
			}},
		{Name: "Sigil of the Empty Throne", TypeLine: "Legendary Enchantment"},
		{Name: "Karlov of the Ghost Council", TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true},
	}, nil, nil, time.Time{})
	h := &CandidateHints{Index: idx}

	for _, tc := range []struct {
		name             string
		wantLead, wantOK bool
	}{
		{"The Arkenstone // Seek the Heart", false, true},
		{"Sigil of the Empty Throne", false, true},
		{"Karlov of the Ghost Council", true, true},
	} {
		lead, known := h.CanLead(tc.name)
		if lead != tc.wantLead || known != tc.wantOK {
			t.Errorf("CanLead(%q) = %v/%v, want %v/%v", tc.name, lead, known, tc.wantLead, tc.wantOK)
		}
	}
}

// TestALegendaryPlaneswalkerStaysUnknown keeps D-269. A characteristic
// -defining ability can make a planeswalker a creature card everywhere
// except the battlefield, and Grist is the card that proved it. The
// engine says nothing about that class.
func TestALegendaryPlaneswalkerStaysUnknown(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{Name: "Some Walker", TypeLine: "Legendary Planeswalker — Test"},
	}, nil, nil, time.Time{})
	h := &CandidateHints{Index: idx}
	if lead, known := h.CanLead("Some Walker"); lead || known {
		t.Errorf("CanLead(planeswalker) = %v/%v, want false/false", lead, known)
	}
}

// TestCanLeadStaysSilentOnALegendaryCard is D-140. Grist, the Hunger
// Tide is a legendary planeswalker, and it is a legal commander: a
// characteristic-defining ability makes it a creature card everywhere
// except the battlefield (D-269). The engine must not claim that it can
// not lead a deck.
func TestCanLeadStaysSilentOnALegendaryCard(t *testing.T) {
	cases := []struct {
		name             string
		typeLine         string
		commander, back  bool
		wantLead, wantOK bool
	}{
		{"Grist, the Hunger Tide", "Legendary Planeswalker — Grist", false, false, false, false},
		{"Jace, the Mind Sculptor", "Legendary Planeswalker — Jace", false, false, false, false},
		{"Lightning Bolt", "Instant", false, false, false, true},
		{"Sol Ring", "Artifact", false, false, false, true},
		{"Karlov of the Ghost Council", "Legendary Creature — Spirit Advisor", true, false, true, true},
		{"a background", "Legendary Enchantment — Background", false, true, false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := &CandidateHints{Index: cards.NewIndex([]*mtgv1.Card{{
				Name: tc.name, TypeLine: tc.typeLine,
				CanBeCommander: tc.commander, IsBackground: tc.back,
			}}, nil, nil, time.Time{})}
			lead, known := h.CanLead(tc.name)
			if lead != tc.wantLead || known != tc.wantOK {
				t.Errorf("CanLead = %v/%v, want %v/%v", lead, known, tc.wantLead, tc.wantOK)
			}
		})
	}
	// An unknown name claims nothing.
	h := &CandidateHints{Index: cards.NewIndex(nil, nil, nil, time.Time{})}
	if _, known := h.CanLead("Nonesuch"); known {
		t.Error("an unknown card name produced a claim")
	}
}

// TestPickRowNeedsNewNames is D-163. A user who answers some other slot
// must not get the same three commanders twice.
func TestPickRowNeedsNewNames(t *testing.T) {
	c := load(t)
	row, ok := c.Row("commander_pick")
	if !ok {
		t.Fatal("the pick row is gone")
	}
	if !row.Repeat || !row.RepeatOnChange {
		t.Fatalf("the pick row repeats %v and narrows it %v", row.Repeat, row.RepeatOnChange)
	}
	base := Context{
		Format:      mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Suggested:   true,
		Filled:      map[string]bool{"format": true, "theme": true, "colors": true, "power": true},
		Asked:       map[string]bool{"commander_pick": true},
		Outstanding: map[string]string{},
	}
	if got := ids(c.Plan(base)); len(got) != 0 {
		t.Errorf("the pick row asked again with the same names: %v", got)
	}
	base.OfferChanged = true
	if got := ids(c.Plan(base)); len(got) != 1 || got[0] != "commander_pick" {
		t.Errorf("the pick row did not ask again with new names: %v", got)
	}
}

// TestLockedCardStaysWithNoQuestion is D-166 and D-260. The locked row
// is retired, so no row asks whether the card may be cut, and the name
// still reaches the build.
func TestLockedCardStaysWithNoQuestion(t *testing.T) {
	out := commanderClassify()
	out.BudgetUSD = 60
	out.LockedNames = []string{"Sanguine Bond"}
	out.Facts.NamedCard = true
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "power_commander"), askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st,
		"Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if q := question(res.Questions, "locked"); q != nil {
		t.Errorf("a retired row asked about the locked card: %q", q.GetText())
	}
	// The card stays in the deck.
	if !hasName(st.LockedCards(), "Sanguine Bond") {
		t.Errorf("locked cards = %v, want Sanguine Bond", st.LockedCards())
	}
}

// TestSuperlativeDelegatesTheCommander is D-167. "Buy the best lifegain
// commander" asks the agent to choose, and three names to choose from do
// not answer it.
func TestSuperlativeDelegatesTheCommander(t *testing.T) {
	if !delegatesCommander("Buy the best lifegain commander") {
		t.Error("a superlative commander instruction was not read as a delegation")
	}
	if delegatesCommander("buy the best lands you can find") {
		t.Error("a message that names no commander handed the commander choice over")
	}
	first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown", BudgetUSD: 60}
	first.Colors = []string{"W", "B"}
	second := classifyOut{Format: "unknown", PoolRule: "owned_first"}
	second.Power = "bracket 3"
	a, _ := testAgentHints(t, &fakeHints{commanders: []string{"Karlov of the Ghost Council"}},
		classifyStep(t, first), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, second), fits(t))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "Lifegain from my collection. Commander, white and black, 60 dollars.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st,
		"Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if q := question(res.Questions, "commander"); q != nil {
		t.Errorf("the agent asked the user to choose after they handed the choice over: %q", q.GetText())
	}
	for _, key := range []string{"commander", "commander_pick"} {
		if !st.Ctx.Filled[key] {
			t.Errorf("key %q is still open although the user asked the agent to choose", key)
		}
	}
}

// TestNamedCommanderClosesTheIllegalRow is one third of H-6. The illegal
// row asks for a replacement, and the replacement arrived through
// SetCommander, which closed every commander key except this one.
func TestNamedCommanderClosesTheIllegalRow(t *testing.T) {
	st := NewState(false)
	st.Ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	st.Ctx.CommanderIllegal, st.IllegalCommander = true, "Lightning Bolt"
	st.MarkAsked("commander_illegal", "commander_illegal", "commander")
	if !st.Outstanding() {
		t.Fatal("the illegal question is not out")
	}
	st.SetCommander("Karlov of the Ghost Council")
	if st.Outstanding() {
		t.Errorf("a question is still out after the user named a commander: %v", st.Slots.GetSlotStates())
	}
	if !st.Ctx.Filled["commander_illegal"] {
		t.Error("the illegal key is still open")
	}
	if st.Ctx.CommanderIllegal {
		t.Error("the illegal fact survived a legal commander")
	}
}

// TestNotOwnedRowIsRetired is D-226. The row asked "You do not own
// {card}. Add it to the buy list, or pick from your library?", and it
// scored badly on every firing.
//
// The rules engine answers the same question after the build, per card
// and with the exact count: a warning in owned-first and a block in
// owned-only (D-37). That costs no turn and names the card.
func TestNotOwnedRowIsRetired(t *testing.T) {
	c := load(t)
	if _, ok := c.Row("commander_not_owned"); ok {
		t.Fatal("the not-owned row is back in the catalog (D-226)")
	}
	// A named commander the collection does not hold still reaches the
	// build, and the commander slot stays filled.
	out := commanderClassify()
	out.PoolRule = "owned_first"
	out.CommanderNames = []string{"Karlov of the Ghost Council"}
	h := &fakeHints{}
	a, _ := testAgentHints(t, h, classifyStep(t, out), fits(t, "power_commander"), askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "Karlov lifegain from my library first, white and black", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	for _, q := range res.Questions {
		if strings.Contains(q.GetText(), "You do not own") {
			t.Errorf("a not-owned question went out: %q", q.GetText())
		}
	}
	if got := st.Slots.GetSlotStates()["commander"]; got != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("commander state = %v, want FILLED: the user named one", got)
	}
}

// TestDeclinedPickClosesTheCommanderSlot is D-208. Turn 2 answers "I
// dunno, you pick", the classifier sets wants_suggestion, and the pick
// row asks. Turn 3 answers "Whatever you think is best", and the
// classifier declines the pick key. The decline closes commander_pick
// alone. The D-147 rule runs after the decline, its guard reads an
// outstanding pick question that the decline already removed, and the
// commander slot stays UNSPECIFIED. The session then reports ready, and
// the gate calls it premature: "commander (never asked)".
//
// A delegation is a decline (D-93, D-147): it closes commander_pick and
// commander, and the generator picks. The gate's required() accepts a
// SKIPPED commander for that reason.
func TestDeclinedPickClosesTheCommanderSlot(t *testing.T) {
	var vague classifyOut
	var delegate classifyOut
	delegate.DeclinedKeys = []string{"format", "colors"}
	delegate.Facts.WantsSuggestion = true
	// The offer waits for the power (D-630), and the reader names it
	// with the delegation. The turn then plans the pick row alone, which
	// is fixed, so it makes no score call and no ask call (D-131).
	delegate.Power = "bracket 3"
	var decline classifyOut
	decline.DeclinedKeys = []string{"commander_pick", "power", "pool_rule"}
	h := &fakeHints{
		commanders: []string{"Peregrin Took", "Joshua, Phoenix's Dominant // Phoenix, Warden of Fire", "Éowyn, Shieldmaiden"},
		second:     []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, vague), fits(t), askStep(t),
		classifyStep(t, delegate),
		classifyStep(t, decline))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "Make me a good deck.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "I dunno, you pick.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if st.Slots.GetSlotStates()["commander_pick"] != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Fatalf("the pick row did not ask on turn 2: %v", st.Slots.GetSlotStates())
	}
	res, err := a.Turn(context.Background(), st, "Whatever you think is best.", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	states := st.Slots.GetSlotStates()
	if states["commander_pick"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("the decline left commander_pick in state %v, want SKIPPED", states["commander_pick"])
	}
	if states["commander"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("the declined pick left the commander slot in state %v, want SKIPPED (D-147)", states["commander"])
	}
	if res.Ready && states["commander"] == mtgv1.SlotState_SLOT_STATE_UNSPECIFIED {
		t.Errorf("the session reported ready with the commander never asked and never chosen")
	}
}

// TestNamedCardThatCanNotLeadSettlesItsRole is D-220. A named Sol Ring
// must not get "Should Sol Ring be your commander or one of the 99
// cards?". D-129 read the commander list alone, and Sol Ring is never on
// it.
func TestNamedCardThatCanNotLeadSettlesItsRole(t *testing.T) {
	h := &CandidateHints{Index: cards.NewIndex([]*mtgv1.Card{
		{Name: "Sol Ring", TypeLine: "Artifact"},
		{Name: "Karlov of the Ghost Council", TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true},
	}, nil, nil, time.Time{})}
	out := commanderClassify()
	out.NamedCards = []string{"Sol Ring"}
	a, _ := testAgentHints(t, h, classifyStep(t, out),
		fits(t, "commander", "power_commander"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Commander deck with Sol Ring in it.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if st.Ctx.Asked["named_card_role"] {
		t.Error("the agent asked whether Sol Ring should be the commander")
	}
	if len(st.LockedCards()) == 0 {
		t.Error("Sol Ring did not reach the 99")
	}
}

// TestSecondIllegalCommanderGetsTheQuestion is D-129 across two turns.
// The illegal row carries no repeat, so a second illegal name needs its
// asked mark cleared, or the agent accepts it in silence (D-195).
func TestSecondIllegalCommanderGetsTheQuestion(t *testing.T) {
	h := &CandidateHints{Index: cards.NewIndex([]*mtgv1.Card{
		{Name: "Lightning Bolt", TypeLine: "Instant"},
		{Name: "Sol Ring", TypeLine: "Artifact"},
	}, nil, nil, time.Time{})}
	first := classifyOut{Format: "commander", Theme: "burn", PoolRule: "unknown", CommanderNames: []string{"Lightning Bolt"}}
	second := classifyOut{Format: "unknown", PoolRule: "unknown", CommanderNames: []string{"Sol Ring"}}
	p := play(t, false, h, []turnScript{
		{"A Commander burn deck with Lightning Bolt as my commander.", first},
		{"Sol Ring as my commander then.", second},
	})
	if !p.askedOn(1, "commander_illegal") {
		t.Fatalf("turn 1 did not refuse Lightning Bolt: %v", p.rows)
	}
	if !p.askedOn(2, "commander_illegal") {
		t.Errorf("turn 2 accepted Sol Ring in silence: %v", p.rows)
	}
	if p.st.IllegalCommander != "Sol Ring" {
		t.Errorf("IllegalCommander = %q, want Sol Ring", p.st.IllegalCommander)
	}
	if p.st.Ctx.CommanderSet {
		t.Error("an illegal card became the commander")
	}
}

// TestABackgroundAloneCanNotLead is D-154 read the other way. A
// Background joins a creature that chooses one, and it never leads a
// deck on its own.
func TestABackgroundAloneCanNotLead(t *testing.T) {
	h := &CandidateHints{Index: cards.NewIndex([]*mtgv1.Card{
		{Name: "Guild Artisan", TypeLine: "Legendary Enchantment — Background", IsBackground: true},
	}, nil, nil, time.Time{})}
	lead, known := h.CanLead("Guild Artisan")
	if lead || !known {
		t.Errorf("CanLead(Background) = %v/%v, want false/true", lead, known)
	}
	out := classifyOut{Format: "commander", Theme: "treasure", PoolRule: "unknown", CommanderNames: []string{"Guild Artisan"}}
	p := play(t, false, h, []turnScript{
		{"A Commander treasure deck with Guild Artisan as my commander.", out},
	})
	if !p.askedOn(1, "commander_illegal") {
		t.Errorf("a sole Background was accepted as the commander: %v", p.rows)
	}
}

// TestCardNamedAfterTheCommanderGetsTheRoleQuestion is D-118. The
// commander closed the role key, and a card named later with no role
// reopens it, so the user says where that card goes.
func TestCardNamedAfterTheCommanderGetsTheRoleQuestion(t *testing.T) {
	h := &CandidateHints{Index: cards.NewIndex([]*mtgv1.Card{
		{Name: "Karlov of the Ghost Council", TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true},
		{Name: "Grist, the Hunger Tide", TypeLine: "Legendary Planeswalker — Grist"},
	}, nil, nil, time.Time{})}
	first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown", CommanderNames: []string{"Karlov of the Ghost Council"}}
	second := classifyOut{Format: "unknown", PoolRule: "unknown", NamedCards: []string{"Grist, the Hunger Tide"}}
	second.Facts.NamedCard = true
	p := play(t, false, h, []turnScript{
		{"A Commander lifegain deck with Karlov of the Ghost Council as my commander.", first},
		{"Also build around Grist, the Hunger Tide.", second},
	})
	if p.askedOn(1, "named_card_role") {
		t.Errorf("turn 1 asked the role of the commander: %v", p.rows[0])
	}
	if !p.askedOn(2, "named_card_role") {
		t.Errorf("turn 2 did not ask where Grist goes: %v", p.rows)
	}
	if !p.st.Ctx.CommanderSet || !hasName(p.st.CommanderNames, "Karlov of the Ghost Council") {
		t.Error("the commander did not survive the second card")
	}
}

// TestBareOrdinalPicksNoCommander is D-121 with the guard. "First, make
// it budget" orders the sentence and chooses no commander.
func TestBareOrdinalPicksNoCommander(t *testing.T) {
	h := &fakeHints{commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Trelasarra, Moon Dancer"}}
	first := commanderClassify()
	first.Facts.WantsSuggestion = true
	// The offer ranks on the bracket, so the pick row waits for the
	// power (D-630). The reader names it with the request.
	first.Power = "bracket 3"
	p := play(t, false, h, []turnScript{
		{"A Commander lifegain deck, white and black. Suggest a commander.", first},
		{"First, make it budget. 50 dollars.", classifyOut{Format: "unknown", PoolRule: "unknown", BudgetUSD: 50}},
		{"The second one.", classifyOut{Format: "unknown", PoolRule: "unknown"}},
	})
	if !p.askedOn(1, "commander_pick") {
		t.Fatalf("turn 1 offered no commanders: %v", p.rows)
	}
	if hasName(p.st.CommanderNames, "Karlov of the Ghost Council") {
		t.Error(`"First, make it budget" picked the first commander`)
	}
	if !hasName(p.st.CommanderNames, "Oloro, Ageless Ascetic") {
		t.Errorf(`"The second one" did not pick the second commander: %v`, p.st.CommanderNames)
	}
}

// The production shape of D-73 and D-120, from session
// 5A1p1rznS2vap8UQg5D2 on 2026-08-31. The reader clicked the option
// "None, name three more", and the classifier reported nothing at all
// for the key: no close, no decline. The word rule alone must reopen the
// row. The test above drives the same refusal through the classifier,
// and this one drives it through the words.
func TestRefusalByOptionAloneRepeatsThePickRow(t *testing.T) {
	base := commanderClassify()
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	// The reader names the bracket in the same message, and the offer
	// ranks on it (D-630). The turn text says so already.
	wants.Power = "bracket 3"
	// The classifier says nothing about the key. The words are all there is.
	silent := commanderClassify()
	h := &fakeHints{
		commanders: []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second:     []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, base), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, wants),
		classifyStep(t, silent))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a lifegain commander deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "Bracket 3. I have no commander in mind, so suggest one.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}

	res, err := a.Turn(context.Background(), st, "None, name three more", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if st.Ready(a.cat) {
		t.Fatal("the session called itself ready after a refusal, so it built with a commander nobody chose")
	}
	q := question(res.Questions, "commander")
	if q == nil {
		t.Fatal("the pick row did not ask again after the refusal")
	}
	for _, old := range h.commanders {
		if strings.Contains(q.GetText(), old) {
			t.Errorf("the agent offered %q again after the user refused it", old)
		}
	}
}

// What happens when the pool has no more names to offer. Session
// 5A1p1rznS2vap8UQg5D2 refused three and the turn then asked nothing at
// all, so the session went on to build with a commander nobody chose.
func TestRefusalWithNoMoreNames(t *testing.T) {
	base := commanderClassify()
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	silent := commanderClassify()
	h := &fakeHints{
		commanders: []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second:     nil, // the pool is out of names
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, base), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, wants),
		classifyStep(t, silent))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a lifegain commander deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "Bracket 3. I have no commander in mind, so suggest one.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "None, name three more", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	t.Logf("questions asked = %d, ready = %v, offer = %v, asked-mark = %v",
		len(res.Questions), st.Ready(a.cat), st.CurrentOffer, st.Ctx.Asked["commander_pick"])
	for _, q := range res.Questions {
		t.Logf("  asked slot=%s text=%q", q.GetSlot(), q.GetText())
	}
}
