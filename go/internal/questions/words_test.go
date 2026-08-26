package questions

import (
	"context"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// TestNegationStopsATriggerWord is D-111. Probe 38 of gate run 11 wrote
// "no proxies", and the plain substring test read it as a proxy user.
func TestNegationStopsATriggerWord(t *testing.T) {
	cases := []struct {
		text, phrase string
		want         bool
	}{
		{"we proxy everything at our table", "proxy", true},
		{"edh gruul dino stompy pls, no proxies", "proxies", false},
		{"i do not want proxies", "proxies", false},
		{"any card, no ban list", "no ban list", true},
		{"i want a commander deck", "commander", true},
		// "Not as my commander" denies the role of a card and not the
		// format. The user still wants a Commander deck.
		{"build around grist, but not as my commander", "commander", true},
		{"i don't want commander", "commander", false},
	}
	for _, tc := range cases {
		if got := hasPhrase(tc.text, tc.phrase); got != tc.want {
			t.Errorf("hasPhrase(%q, %q) = %v, want %v", tc.text, tc.phrase, got, tc.want)
		}
	}
}

// TestFormatFromWords covers the safety net under the classifier. Every
// case below is a first message of a gate conversation.
func TestFormatFromWords(t *testing.T) {
	const none = mtgv1.FormatId_FORMAT_ID_UNSPECIFIED
	cases := map[string]mtgv1.FormatId{
		// Conversation 23 of gate run 11 asked the format after this.
		"a land destruction commander deck": mtgv1.FormatId_FORMAT_ID_COMMANDER,
		// Conversation 27 asked the format after this, in all four runs.
		"build around grist, the hunger tide, but not as my commander":  mtgv1.FormatId_FORMAT_ID_COMMANDER,
		"i want the strongest modern deck, money is no object":          mtgv1.FormatId_FORMAT_ID_MODERN,
		"upgrade my atraxa, praetors' voice precon. we play bracket 2":  mtgv1.FormatId_FORMAT_ID_COMMANDER,
		"edh gruul dino stompy pls, no proxies":                         mtgv1.FormatId_FORMAT_ID_COMMANDER,
		"i already have a modern burn deck. i only need sideboard help": mtgv1.FormatId_FORMAT_ID_MODERN,
		"a pauper burn deck, as cheap as possible":                      mtgv1.FormatId_FORMAT_ID_PAUPER,
		// Two formats is a two-deck request, and the one-deck row reads it.
		"i want two decks, one commander and one modern": none,
		// A format this app does not build answers nothing here.
		"i want a brawl deck for arena":   none,
		"build me a lifegain deck":        none,
		"make me something fun and janky": none,
	}
	for text, want := range cases {
		got, ok := FormatFromWords(text)
		if want == none {
			if ok {
				t.Errorf("FormatFromWords(%q) = %v, want no answer", text, got)
			}
			continue
		}
		if !ok || got != want {
			t.Errorf("FormatFromWords(%q) = %v (%v), want %v", text, got, ok, want)
		}
	}
}

// TestUnsupportedFormat is D-112. Gate run 13 offered Brawl to probe 46,
// and this app can not build it.
func TestUnsupportedFormat(t *testing.T) {
	cases := []struct{ text, name, near string }{
		{"i want a brawl deck for arena", "Brawl", "Commander"},
		{"a duel commander deck please", "Duel Commander", "Commander"},
		{"an oathbreaker deck", "Oathbreaker", "Commander"},
		{"a commander deck with lightning bolt", "", ""},
		{"a modern burn deck", "", ""},
	}
	for _, tc := range cases {
		name, near, ok := UnsupportedFormat(tc.text)
		if tc.name == "" {
			if ok {
				t.Errorf("UnsupportedFormat(%q) = %q, want none", tc.text, name)
			}
			continue
		}
		if !ok || name != tc.name || near != tc.near {
			t.Errorf("UnsupportedFormat(%q) = %q/%q, want %q/%q", tc.text, name, near, tc.name, tc.near)
		}
	}
}

// TestOneDeckRequestReadsOneMessage is D-112. A user who changes the
// format across two turns has not asked for two decks, and probe 31 does
// exactly that.
func TestOneDeckRequestReadsOneMessage(t *testing.T) {
	cases := map[string]bool{
		"i want two decks, one commander and one modern": true,
		"i want a second deck as well":                   true,
		"start with the commander one. a dragon deck":    false,
		"actually, make it modern instead. sixty cards":  false,
		"i want a commander deck":                        false,
	}
	for text, want := range cases {
		if got := OneDeckRequest(text); got != want {
			t.Errorf("OneDeckRequest(%q) = %v, want %v", text, got, want)
		}
	}
}

// TestSameCardMergesOnlyAShortName guards the fix for the stuttered name.
// Two names that both hold a comma are two different cards.
func TestSameCardMergesOnlyAShortName(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"Grist", "Grist, the Hunger Tide", true},
		{"Grist, the Hunger Tide", "grist", true},
		{"Brago, King Eternal", "Brago", true},
		{"Toph, Hardheaded Teacher", "Toph, the Blind Bandit", false},
		{"Sanguine Bond", "Grist", false},
		{"", "Grist", false},
	}
	for _, tc := range cases {
		if got := sameCard(tc.a, tc.b); got != tc.want {
			t.Errorf("sameCard(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

// TestLockedNamesDoNotStutter is the defect gate runs 10 to 12 recorded.
// The locked row read "Must the deck keep Grist, the Hunger Tide and
// Grist, or may I cut a card that does not fit the plan?"
func TestLockedNamesDoNotStutter(t *testing.T) {
	st := NewState(true)
	st.AddLocked("Grist, the Hunger Tide")
	st.AddLocked("Grist")
	if got := englishList(st.LockedCards()); got != "Grist, the Hunger Tide" {
		t.Errorf("locked cards = %q, want the full name once", got)
	}
	// A second, different card still joins the list.
	st.AddLocked("Sanguine Bond")
	if got := englishList(st.LockedCards()); got != "Grist, the Hunger Tide and Sanguine Bond" {
		t.Errorf("locked cards = %q, want both cards", got)
	}
}

// TestCardInThe99ClosesTheRoleRow is D-70. Conversation 27 opens with
// "Build around Grist, the Hunger Tide, but not as my commander", and all
// four runs of 2026-08-25 then asked whether Grist should be the
// commander.
func TestCardInThe99ClosesTheRoleRow(t *testing.T) {
	out := classifyOut{Format: "commander", PoolRule: "unknown"}
	out.LockedNames = []string{"Grist, the Hunger Tide"}
	out.Facts.NamedCard = true
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "theme_card_named", "colors"), askStep(t))
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

// TestCompetitiveRequestInfersTheTournamentStep is D-107. Conversation 16
// opens with "I want the strongest Modern deck, money is no object", and
// every run then asked how strong the deck should be.
func TestCompetitiveRequestInfersTheTournamentStep(t *testing.T) {
	out := classifyOut{Format: "modern", Theme: "best deck", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Facts.PowerCompetitive = true
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "power_sixty_confirm", "meta"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I want the strongest Modern deck, money is no object.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	// The slot must hold a value whatever the user answers next (D-90).
	if got := st.Slots.GetPower().GetSixtyStep(); got != mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT {
		t.Errorf("power = %v, want the tournament step", got)
	}
	if !st.Ctx.Filled["power"] {
		t.Error("the power slot did not close")
	}
	if !st.Ctx.Asked["power_sixty_confirm"] {
		t.Error("the agent did not ask the user to confirm the inferred step")
	}
	if st.Ctx.Asked["power_sixty"] {
		t.Error("the open power question went out beside the confirm question")
	}
	if q := question(res.Questions, "power"); q == nil {
		t.Fatal("no power question went out")
	}
}

// TestWordRuleReadsTheFormatTheClassifierMissed is the safety net of
// D-116. Conversation 23 of gate run 11 named Commander in the first
// message and still got the format question.
func TestWordRuleReadsTheFormatTheClassifierMissed(t *testing.T) {
	out := classifyOut{Format: "unknown", Theme: "land destruction", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out),
		fits(t, "commander", "power_commander", "colors"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A land destruction Commander deck.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Fatalf("format = %v, want Commander", st.Ctx.Format)
	}
	if st.Ctx.Asked["format"] {
		t.Error("the agent asked for a format the user had already named")
	}
}

// TestUnsupportedFormatRowNamesTheNearestFormat is D-112.
func TestUnsupportedFormatRowNamesTheNearestFormat(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out),
		fits(t, "format_unsupported", "theme", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I want a Brawl deck for Arena.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	q := question(res.Questions, "format")
	if q == nil {
		t.Fatal("no format question went out")
	}
	for _, want := range []string{"Brawl", "Commander"} {
		if !strings.Contains(q.GetText(), want) {
			t.Errorf("the question %q does not name %q", q.GetText(), want)
		}
	}
}

// TestTwoDeckRequestAsksNothingElse is D-112. Probe 50 chose a deck
// silently in gate runs 11 to 13.
func TestTwoDeckRequestAsksNothingElse(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "one_deck"))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I want two decks, one Commander and one Modern.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(res.Questions) != 1 {
		t.Fatalf("%d questions went out, want 1", len(res.Questions))
	}
	if !st.Ctx.Asked["one_deck"] {
		t.Error("the agent did not say that it builds one deck at a time")
	}
}

// TestFixedRowIsNeverReplaced covers the rows that state what this app
// does or does not do. The smoke run of 2026-08-26 replaced "I build one
// deck at a time. Which deck do you want first?" with a question that
// never said the app builds one deck at a time.
func TestFixedRowIsNeverReplaced(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	// No ask step: a fixed row never reaches the ask role, so the turn
	// costs two model calls and not three.
	a, _ := testAgent(t,
		classifyStep(t, out),
		scoreStep(t, scored{RowID: "one_deck", Fit: 0.05, Reason: "test",
			CustomText: "Which deck would you like to work on first: Commander or Modern?"}))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I want two decks, one Commander and one Modern.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(res.Questions) != 1 {
		t.Fatalf("%d questions went out, want 1", len(res.Questions))
	}
	q := res.Questions[0]
	if q.GetInvented() {
		t.Errorf("a row that states an app limit was replaced: %q", q.GetText())
	}
	if !strings.Contains(q.GetText(), "one deck at a time") {
		t.Errorf("the question does not say what the app does: %q", q.GetText())
	}
	if res.Coverage.NearCopies != 1 {
		t.Errorf("near copies = %d, want the refusal counted", res.Coverage.NearCopies)
	}
}

// TestEveryLimitRowIsFixed keeps a new limit row from missing the flag.
func TestEveryLimitRowIsFixed(t *testing.T) {
	c := load(t)
	for _, id := range []string{"out_of_scope", "one_deck", "format_unsupported", "commander_illegal"} {
		row, ok := c.Row(id)
		if !ok {
			t.Fatalf("no row %q", id)
		}
		if !row.Fixed {
			t.Errorf("row %q states an app limit and the model may replace it", id)
		}
	}
}

// TestAnswerThatRepeatsAnOptionClosesTheKey is D-119. Conversation 5 asks
// "When you say anything goes, do you mean any card with no ban list, or
// Vintage rules?" The user answers "Any card, no ban list." The
// classifier left the key open in gate run 13 and in the batch run.
func TestAnswerThatRepeatsAnOptionClosesTheKey(t *testing.T) {
	first := classifyOut{Format: "unknown", PoolRule: "unknown", Theme: "dragons"}
	// Turn 2 names the format and the theme, and reports no closed key.
	second := classifyOut{Format: "modern", PoolRule: "unknown", Theme: "dragons"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "format", "theme", "house_rules"), askStep(t),
		classifyStep(t, second), fits(t, "power_sixty", "colors"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "I want a 60-card deck, and anything goes at our table.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.Asked["house_rules"] {
		t.Fatal("the house-rules question never went out")
	}
	if _, err := a.Turn(context.Background(), st, "Any card, no ban list. Call it Modern. A dragon deck.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !st.Ctx.Filled["house_rules"] {
		t.Error("the user repeated an option and the key stayed open")
	}
}

// TestAShortOptionClosesNothing keeps the net from firing on a common
// word. "Yes" answers any question, and this one answers none of them.
func TestAShortOptionClosesNothing(t *testing.T) {
	first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	second := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "power_commander", "colors"), askStep(t),
		classifyStep(t, second), fits(t), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a lifegain commander deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	before := len(st.Ctx.Filled)
	if _, err := a.Turn(context.Background(), st, "yes", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if len(st.Ctx.Filled) != before {
		t.Errorf("a bare \"yes\" closed a key: %v", st.Ctx.Filled)
	}
}

// TestOptionAnsweredReadsThePrefix covers the shape of a real answer. An
// option reads "Yes, the normal limits" and the user writes "The normal
// limits hold".
func TestOptionAnsweredReadsThePrefix(t *testing.T) {
	cases := []struct {
		msg, opt string
		want     bool
	}{
		{"any card, no ban list. call it modern. a dragon deck.", "Any card, no ban list", true},
		{"the normal limits hold. casual power, red and green.", "Yes, the normal limits", true},
		{"grist goes in the 99. bracket 3.", "In the 99", true},
		// The classifier owns this one. A commander name closes the row by
		// value, and the option net never sees it (D-83).
		{"karlov is my commander", "As my commander", false},
		// A bare yes answers any question, so it answers none of them.
		{"yes", "Yes, the normal limits", false},
		{"no", "No, I will say the limits", false},
		{"bracket 3, and build from my library first", "Vintage rules", false},
	}
	for _, tc := range cases {
		if got := optionAnswered(tc.msg, tc.opt); got != tc.want {
			t.Errorf("optionAnswered(%q, %q) = %v, want %v", tc.msg, tc.opt, got, tc.want)
		}
	}
}

// TestNoneRepeatsThePickRowWithNewNames is D-73 and D-120. Conversation
// 14 is named "the user says none, then picks". The classifier closed the
// pick row on "None of those." in gate runs 12 and 13, so the session
// called itself complete with a commander nobody chose.
func TestNoneRepeatsThePickRowWithNewNames(t *testing.T) {
	base := commanderClassify()
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	// Turn 3 refuses, and the classifier tries to close the row by name.
	refuse := commanderClassify()
	// The classifier reported the refusal through both channels in gate
	// runs 12 and 13. Neither may close the row.
	refuse.ClosedKeys = []string{"commander_pick"}
	refuse.DeclinedKeys = []string{"commander_pick"}
	h := &offerHints{
		first:  []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, base), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, wants), fits(t, "commander_pick"),
		classifyStep(t, refuse), fits(t, "commander_pick"))
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
	for _, old := range h.first {
		if strings.Contains(q.GetText(), old) {
			t.Errorf("the agent offered %q again after the user refused it", old)
		}
	}
	if st.Ready(a.cat) {
		t.Error("the session called itself complete with no commander chosen")
	}
}

// offerHints names one set of commanders, then another once the first set
// is retired.
type offerHints struct {
	first, second []string
}

func (o *offerHints) ThemeColors(string) string  { return "" }
func (o *offerHints) OwnedThemeCount(string) int { return 0 }
func (o *offerHints) Commanders(_ string, skip []string) []string {
	if len(skip) > 0 {
		return o.second
	}
	return o.first
}

// TestCommanderChosenByPlace is D-121. Conversation 14 ends with "The
// first of the new three is good", and the classifier can not map that
// onto a name, because it never sees the names.
func TestCommanderChosenByPlace(t *testing.T) {
	base := commanderClassify()
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	pick := commanderClassify()
	h := &offerHints{
		first:  []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, base), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, wants), fits(t, "commander_pick"),
		classifyStep(t, pick), fits(t))
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
	if len(st.CommanderNames) == 0 || st.CommanderNames[0] != h.first[1] {
		t.Errorf("commander = %v, want %q", st.CommanderNames, h.first[1])
	}
}

// TestOfferedPickReadsOnlyAClearOrdinal keeps "three more" out. A user
// who refuses every name writes the word "three", and that is not a pick.
func TestOfferedPickReadsOnlyAClearOrdinal(t *testing.T) {
	cases := map[string]int{
		"the first of the new three is good":         0,
		"The second one is good.":                    1,
		"give me the 3rd":                            2,
		"none of those, name three more":             -1,
		"bracket 3, and build from my library first": -1,
	}
	for msg, want := range cases {
		got, ok := OfferedPick(msg)
		if want < 0 {
			if ok {
				t.Errorf("OfferedPick(%q) = %d, want no pick", msg, got)
			}
			continue
		}
		if !ok || got != want {
			t.Errorf("OfferedPick(%q) = %d (%v), want %d", msg, got, ok, want)
		}
	}
}

// TestAMessageThatAnswersItsOwnTriggerAsksNothing is D-122. Conversation
// 21 writes "Any card, no ban list. Call it Vintage. An artifact prison
// deck." The words "no ban list" raise the house-rules row, and the same
// message answers it. The agent asked it anyway.
func TestAMessageThatAnswersItsOwnTriggerAsksNothing(t *testing.T) {
	out := classifyOut{Format: "vintage", Theme: "artifact prison", PoolRule: "unknown"}
	out.Facts.HouseFormat = true
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "power_sixty", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st,
		"Any card, no ban list. Call it Vintage. An artifact prison deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if st.Ctx.Asked["house_rules"] {
		t.Error("the agent asked what the same message had answered")
	}
	if !st.Ctx.Filled["house_rules"] {
		t.Error("the house-rules key stayed open although the user answered it")
	}
	// The house-format row shares the slot, and it is the right next
	// question: a closed key frees the row that waited on it.
	if !st.Ctx.Asked["house_format_limits"] {
		t.Error("the freed row did not fire in the same turn")
	}
	_ = res
}

// TestSuggestionDoesNotSwapTheNames is D-123, which restores D-80. The
// classifier set wants_suggestion again in conversation 23 of the batch
// run, on a message that refused nothing. The agent swapped all three
// commanders under the user.
func TestSuggestionDoesNotSwapTheNames(t *testing.T) {
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	again := commanderClassify()
	again.Facts.WantsSuggestion = true
	h := &offerHints{
		first:  []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, wants), fits(t, "commander_pick", "power_commander"), askStep(t),
		classifyStep(t, again), fits(t, "commander_pick"))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a lifegain commander deck, you pick the commander", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "Bracket 3, and build from my library first.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	q := question(res.Questions, "commander")
	if q == nil {
		t.Fatal("the pick row did not ask again")
	}
	for _, want := range h.first {
		if !strings.Contains(q.GetText(), want) {
			t.Errorf("the agent dropped %q although the user refused nothing: %q", want, q.GetText())
		}
	}
}

// TestHintsReadTheColorsOfThisTurn is D-124. The caller builds the hint
// source before the turn, so a color the classifier fills inside the turn
// was invisible. Conversation 23 offered a five-color commander for a
// red-green deck.
func TestHintsReadTheColorsOfThisTurn(t *testing.T) {
	h := &slotHints{}
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

// slotHints records the colors the agent handed it inside the turn.
type slotHints struct {
	saw []mtgv1.Color
}

func (s *slotHints) UseSlots(_ mtgv1.FormatId, colors []mtgv1.Color, _ mtgv1.PoolRule) {
	s.saw = colors
}
func (s *slotHints) ThemeColors(string) string  { return "" }
func (s *slotHints) OwnedThemeCount(string) int { return 0 }
func (s *slotHints) Commanders(string, []string) []string {
	return []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}
}

// TestFormatDoesNotRevertToAReplacedValue is D-125. Probe 31 asks for
// Commander, changes to Modern two turns later, and the classifier
// reported Commander again on a message that named no format. The agent
// then asked for a commander in a 60-card format.
func TestFormatDoesNotRevertToAReplacedValue(t *testing.T) {
	first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	first.Colors = []string{"W", "B"}
	change := classifyOut{Format: "modern", Theme: "lifegain", PoolRule: "unknown"}
	// The classifier repeats the value the user replaced.
	stale := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, change), fits(t, "power_sixty"), askStep(t),
		classifyStep(t, stale), fits(t))
	st := NewState(false)
	for i, msg := range []string{
		"Build me a white-black lifegain Commander deck.",
		"Actually, make it Modern instead. Sixty cards.",
		"FNM level, and 100 dollars.",
	} {
		if _, err := a.Turn(context.Background(), st, msg, nil); err != nil {
			t.Fatalf("turn %d: %v", i+1, err)
		}
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_MODERN {
		t.Errorf("format = %v, want Modern: the user replaced Commander", st.Ctx.Format)
	}
	if st.Ctx.Asked["commander_pick"] {
		t.Error("the agent asked for a commander in a 60-card format")
	}
}

// TestOneRowPerKeyWhileAQuestionIsOut is D-126. Probe 33 asked the theme
// through the competitive row, got no answer, and asked it again through
// the general row one turn later.
func TestOneRowPerKeyWhileAQuestionIsOut(t *testing.T) {
	c := load(t)
	ctx := Context{
		Format: mtgv1.FormatId_FORMAT_ID_MODERN,
		Filled: map[string]bool{"format": true, "colors": true, "power": true},
		Asked:  map[string]bool{"theme_competitive": true},
		// The competitive theme question is out with no answer.
		Outstanding:      map[string]string{"theme": "theme"},
		PowerCompetitive: true,
	}
	for _, r := range c.Plan(ctx) {
		if r.Slot == "theme" {
			t.Errorf("row %q asked the theme again while a theme question was out", r.ID)
		}
	}
	// Once the key closes, a later row may ask it again.
	ctx.Outstanding = map[string]string{}
	ctx.Filled["theme"] = true
	for _, r := range c.Plan(ctx) {
		if r.Slot == "theme" {
			t.Errorf("row %q asked a filled theme", r.ID)
		}
	}
}

// TestPickRowWithNoNamesAsksNothing is D-127. Probe 35 answers every
// question with "you pick", so the theme stays empty and PR-6 can name no
// commander. The pick row then asked "Which commander would you like, or
// should I suggest three more?" twice, with nothing to suggest.
func TestPickRowWithNoNamesAsksNothing(t *testing.T) {
	out := commanderClassify()
	out.Theme = ""
	out.Facts.WantsSuggestion = true
	// The hint source names no commander, which is what an empty theme
	// gives.
	a, _ := testAgentHints(t, stubHints{}, classifyStep(t, out), fits(t, "power_commander"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "Make me a good deck. I dunno, you pick.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
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

// TestCommanderSwapReopensTheChoice is D-130. Probe 49 chooses Karlov of
// the Ghost Council, then writes "Actually use a different commander,
// suggest one". Every run before 2026-08-26 asked nothing after it,
// because a chosen commander closes every commander row.
func TestCommanderSwapReopensTheChoice(t *testing.T) {
	named := commanderClassify()
	named.CommanderNames = []string{"Karlov of the Ghost Council"}
	swap := commanderClassify()
	swap.Facts.WantsSuggestion = true
	h := &offerHints{
		first:  []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second: []string{"Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim", "Liesa, Shroud of Dusk"},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, named), fits(t, "power_commander"), askStep(t),
		classifyStep(t, swap), fits(t, "commander_pick"))
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

// TestCanLeadStaysSilentOnALegendaryCard is D-140. Gate run 14 told a
// user "Grist, the Hunger Tide can not lead a deck". Grist is a legendary
// planeswalker, and it is a legal commander: a characteristic-defining
// ability makes it a creature card everywhere except the battlefield
// (Scryfall ruling, 2021-06-18). The gate passed and the linter found
// nothing, and the claim was false.
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
		{"a background", "Legendary Enchantment — Background", false, true, true, true},
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
