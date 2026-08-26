package questions

import (
	"context"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
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
		// D-155: Pauper is not a format this app builds, so the word rule
		// answers nothing and the unsupported row declines it.
		"a pauper burn deck, as cheap as possible": none,
		"a pioneer aggro deck":                     none,
		"a legacy delver deck":                     none,
		"a vintage shops deck":                     none,
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
//
// The message asked for a suggestion before D-147. It now asks without
// the words that hand the choice over, because a delegation closes the
// pick row instead of repeating it. D-123 is about the names on the
// table, and this test still measures only that.
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
	for _, want := range h.first {
		if !hasName(st.CurrentOffer, want) {
			t.Errorf("the agent dropped %q although the user refused nothing: %v", want, st.CurrentOffer)
		}
	}
}

// TestDelegationClosesTheCommanderPick is D-147. "You pick the commander"
// appears in 18 of the 100 gate conversations, and no rule read it. The
// pick row carries "repeat": true, so it asked again every turn until the
// messages ran out. Eval run 14 refused 18 of its 43 bad questions on
// that row, more than the next four rows together.
//
// A delegation is a decline (D-93): the key closes, it takes no value,
// and the generator picks the best commander of the pool.
func TestDelegationClosesTheCommanderPick(t *testing.T) {
	wants := commanderClassify()
	wants.Facts.WantsSuggestion = true
	again := commanderClassify()
	again.Facts.WantsSuggestion = true
	h := &offerHints{
		first:  []string{"Vito, Thorn of the Dusk Rose", "Heliod, Sun-Crowned", "Haliya, Guided by Light"},
		second: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"},
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

// TestDelegationNeedsTheCommanderInScope guards D-147 against a general
// "up to you". The phrase closes the commander choice only while a
// commander question is out, or while the message names one.
func TestDelegationNeedsTheCommanderInScope(t *testing.T) {
	if !DelegatesChoice("You pick the commander.") {
		t.Error("a plain delegation was not read")
	}
	if !DelegatesChoice("I dunno, you pick.") {
		t.Error("conversation 35 answers this, and it was not read")
	}
	if !DelegatesChoice("You decide.") {
		t.Error("probe 47 answers this, and it was not read")
	}
	if DelegatesChoice("I will pick the commander myself.") {
		t.Error("the user kept the choice, and the rule took it")
	}
	if !NamesCommander("You pick the commander.") {
		t.Error("the scope guard missed the word")
	}
	if NamesCommander("Up to you.") {
		t.Error("the scope guard read a word that is not there")
	}
}

// colorOfferHints names commanders and knows each one's color identity.
// It is the offerHints of D-123 with an IdentityChecker on top.
type colorOfferHints struct {
	first, second []string
	identity      map[string][]mtgv1.Color
}

func (o *colorOfferHints) ThemeColors(string) string  { return "" }
func (o *colorOfferHints) OwnedThemeCount(string) int { return 0 }
func (o *colorOfferHints) Commanders(_ string, skip []string) []string {
	if len(skip) > 0 {
		return o.second
	}
	return o.first
}
func (o *colorOfferHints) FitsColors(name string, colors []mtgv1.Color) (bool, bool) {
	id, ok := o.identity[name]
	if !ok {
		return false, false
	}
	return candidates.IdentityMatches(id, colors), true
}

// TestOffColorOfferLeavesTheTable is D-153. The pick row keeps the names
// on the table until the user refuses them (D-80, D-123). Nothing checked
// them again when the colors arrived later.
//
// Probe 73 of gate run 16 is the case. Turn 1 named no colors, and the
// row offered Jaheira, Friend of the Forest, which is mono-green. The
// user answered "Red and white" on turn 2, and the same three names went
// out on turns 2 and 3. The eval caught it, and D-148 could not: it
// filters the pool, and these names were already on the table.
func TestOffColorOfferLeavesTheTable(t *testing.T) {
	first := commanderClassify()
	first.Facts.WantsSuggestion = true
	first.Colors = nil
	second := commanderClassify()
	second.Facts.WantsSuggestion = true
	second.Colors = []string{"R", "W"}
	// Two of the three offered names are red-white, so they survive the
	// colors. Jaheira is mono-green and must leave. A mono-red name would
	// leave as well, because D-148 asks a commander to hold every color
	// the user named, so the test uses names that isolate D-153.
	h := &colorOfferHints{
		first:  []string{"Jaheira, Friend of the Forest", "Winota, Joiner of Forces", "Feather, the Redeemed"},
		second: []string{"Aurelia, the Warleader", "Anax and Cymede", "Tajic, Blade of the Legion"},
		identity: map[string][]mtgv1.Color{
			"Jaheira, Friend of the Forest": {mtgv1.Color_COLOR_G},
			"Winota, Joiner of Forces":      {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Feather, the Redeemed":         {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Aurelia, the Warleader":        {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Anax and Cymede":               {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
			"Tajic, Blade of the Legion":    {mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W},
		},
	}
	a, _ := testAgentHints(t, h,
		classifyStep(t, first), fits(t, "commander_pick", "power_commander"), askStep(t),
		classifyStep(t, second), fits(t, "commander_pick"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Commander deck with a Background commander pair.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !contains(st.CurrentOffer, "Jaheira, Friend of the Forest") {
		t.Fatalf("turn 1 did not offer the green commander: %v", st.CurrentOffer)
	}
	if len(st.CurrentOffer) != 3 {
		t.Fatalf("turn 1 offered %d names, want 3: %v", len(st.CurrentOffer), st.CurrentOffer)
	}
	res, err := a.Turn(context.Background(), st, "Red and white, aggressive.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
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
	h := &colorOfferHints{identity: map[string][]mtgv1.Color{"Known Legend": {mtgv1.Color_COLOR_G}}}
	if fits, known := h.FitsColors("A Card Nobody Holds", []mtgv1.Color{mtgv1.Color_COLOR_R}); known || fits {
		t.Errorf("an unknown name answered fits=%v known=%v, want both false", fits, known)
	}
	// An empty color list fits everything: nothing is out of no colors.
	if fits, known := h.FitsColors("Known Legend", nil); !fits || !known {
		t.Errorf("no colors named answered fits=%v known=%v, want both true", fits, known)
	}
}

// TestWantsCommanderPair is D-154. Probe 73 writes "A Commander deck
// with a Background commander pair", and every run before this read it as
// a request for one commander.
func TestWantsCommanderPair(t *testing.T) {
	want := []string{
		"A Commander deck with a Background commander pair.",
		"A Commander deck with Thrasios, Triton Hero and Tymna the Weaver as partners.",
		"I want two commanders.",
		"Build a four-color deck with partners.",
		"Give me a Doctor's companion deck.",
	}
	for _, m := range want {
		if !WantsCommanderPair(m) {
			t.Errorf("a pair request was not read: %q", m)
		}
	}
	// The negation guard applies, and an unrelated message names none.
	notWant := []string{
		"No partners, just one commander.",
		"A lifegain Commander deck from my library.",
		"Bracket 3, and 150 dollars.",
	}
	for _, m := range notWant {
		if WantsCommanderPair(m) {
			t.Errorf("a message that asks for no pair was read as one: %q", m)
		}
	}
	// A Background request narrows the pair, because a Background carries
	// no theme signal and loses on score without this.
	if !WantsBackgroundPair("A Commander deck with a Background commander pair.") {
		t.Error("a Background request was not read")
	}
	if WantsBackgroundPair("A Commander deck with partners.") {
		t.Error("a plain partner request was read as a Background request")
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

// TestHouseLimitsRowGoesOutAsWritten is the half of D-162 the evidence
// supports. The row bundles three limits into one yes-or-no question:
// "do the normal limits hold: ...". The ask role rewrote it as "should
// the deck use a 60-card minimum, four copies per name, and a 15-card
// sideboard?", and the eval read three questions in one (conversations
// 21 and 34 of run 20260826-191225-000). A fixed row never reaches the
// ask role.
//
// The confirm row is not fixed. Its own words assert a premise the user
// may not have given, and the eval refused the fixed text in
// conversations 33, 49, and 68 of run 20260826-191225-001.
func TestHouseLimitsRowGoesOutAsWritten(t *testing.T) {
	c := load(t)
	row, ok := c.Row("house_format_limits")
	if !ok {
		t.Fatal("the house-limits row is gone")
	}
	if !row.Fixed {
		t.Error("the house-limits row is not fixed, so the ask role may split it into three questions")
	}
	if confirm, ok := c.Row("power_sixty_confirm"); !ok || confirm.Fixed {
		t.Error("the confirm row is fixed, and its fixed text was refused three times in run 20260826-191225-001")
	}
	out := classifyOut{Format: "modern", Theme: "dragons", PoolRule: "any_card", BudgetUSD: 100}
	out.Colors = []string{"R"}
	out.Power = "casual"
	out.Facts.HouseFormat = true
	a, _ := testAgent(t,
		classifyStep(t, out), fits(t, "house_rules"), askStep(t),
		classifyStep(t, classifyOut{Format: "unknown", PoolRule: "unknown", ClosedKeys: []string{"house_rules"}}),
		fits(t, "house_format_limits"),
		askStep(t, phrasing{RowID: "house_format_limits", Text: "Should the deck use a 60-card minimum, four copies per name, and a 15-card sideboard?"}))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "a 60-card dragons deck, anything goes at our table", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "any card, no ban list", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	q := question(res.Questions, "house_rules")
	if q == nil {
		t.Fatalf("the house-limits row did not fire: %v", ids2(res.Questions))
	}
	if !strings.Contains(q.GetText(), "do the normal limits hold") {
		t.Errorf("the house-limits question lost the clause that bundles the limits: %q", q.GetText())
	}
}

// TestPickRowNeedsNewNames is D-163. Conversations 1, 77, and 90 of gate
// run 18 each got the same three commanders twice, because the user
// answered some other slot and the row repeats every turn.
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

// TestOfferChangedReadsTheTable proves what the planner reads: the names
// on the table against the names the row sent last.
func TestOfferChangedReadsTheTable(t *testing.T) {
	st := NewState(false)
	names := []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}
	st.SetOffer(names)
	st.RecordAskedOffer(names)
	if st.OfferChanged() {
		t.Error("an unchanged table reads as changed")
	}
	st.RetireOffer()
	if !st.OfferChanged() {
		t.Error("a refusal left the table unchanged")
	}
	// A dropped name is the D-153 case: the colors arrived and one name
	// no longer fits.
	st.SetOffer(names)
	st.RecordAskedOffer(names)
	st.CurrentOffer = names[:2]
	if !st.OfferChanged() {
		t.Error("a dropped name left the table unchanged")
	}
}

// TestLocksCard is D-166. Conversation 13 of gate run 18 opened with
// "keep Sanguine Bond in it", and the agent asked whether the deck must
// keep Sanguine Bond.
func TestLocksCard(t *testing.T) {
	cases := []struct {
		message, name string
		want          bool
	}{
		{"Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it", "Sanguine Bond", true},
		{"lock Sanguine Bond in", "Sanguine Bond", true},
		{"it must include Sanguine Bond", "Sanguine Bond", true},
		{"build around Grist, the Hunger Tide", "Grist, the Hunger Tide", false},
		{"and Sanguine Bond as well", "Sanguine Bond", false},
		{"do not keep Sanguine Bond", "Sanguine Bond", false},
		{"Karlov of the Ghost Council. Keep the buy list under 50 dollars.", "Karlov of the Ghost Council", false},
		// The short form is the one the user writes on the second turn.
		{"keep Grist in", "Grist, the Hunger Tide", true},
	}
	for _, tc := range cases {
		t.Run(tc.message, func(t *testing.T) {
			if got := LocksCard(tc.message, tc.name); got != tc.want {
				t.Errorf("LocksCard(%q, %q) = %v, want %v", tc.message, tc.name, got, tc.want)
			}
		})
	}
}

// TestLockedCardClosesOnAKeepInstruction is the turn D-166 fixes.
func TestLockedCardClosesOnAKeepInstruction(t *testing.T) {
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
		t.Errorf("the locked row asked what the user had just said: %q", q.GetText())
	}
	if !st.Ctx.Filled["locked"] {
		t.Error("the locked slot is still open although the user locked a card in")
	}
	// The card stays in the deck. The slot closed on a value, not on a
	// decline.
	if !hasName(st.LockedCards(), "Sanguine Bond") {
		t.Errorf("locked cards = %v, want Sanguine Bond", st.LockedCards())
	}
}

// TestColorlessClosesTheColorSlot is D-165. The classify schema holds the
// five colors alone, so no model call can report a colorless deck. Probe
// 73 of gate run 18 answered "A colorless Commander deck" and got the
// color question.
func TestColorlessClosesTheColorSlot(t *testing.T) {
	if !ColorlessRequest("a colorless Commander deck built around big artifacts") {
		t.Error("a colorless request was not read")
	}
	if ColorlessRequest("not colorless, I want green") {
		t.Error("a negated colorless request fired")
	}
	out := classifyOut{Format: "commander", Theme: "big artifacts", PoolRule: "any_card", BudgetUSD: 250}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "power_commander"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "A colorless Commander deck built around big artifacts.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if q := question(res.Questions, "colors"); q != nil {
		t.Errorf("the color row asked a user who said colorless: %q", q.GetText())
	}
	if !st.Ctx.Filled["colors"] {
		t.Error("the color slot is still open")
	}
}

// TestCEDHNamesBracketFive is D-164. Probe 75 opens with "A cEDH deck",
// and gate run 18 asked which power bracket to target.
func TestCEDHNamesBracketFive(t *testing.T) {
	if !CEDHRequest("a cEDH deck") {
		t.Error("cEDH was not read as a power level")
	}
	if CEDHRequest("a competitive Commander deck") {
		t.Error("competitive Commander names no bracket, and it fired")
	}
	if id, ok := FormatFromWords("a cEDH deck"); !ok || id != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("cEDH names format %v/%v, want Commander", id, ok)
	}
	out := classifyOut{Format: "commander", Theme: "combo", PoolRule: "any_card", BudgetUSD: 500}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "A cEDH deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if q := question(res.Questions, "power"); q != nil {
		t.Errorf("the bracket row asked a cEDH user: %q", q.GetText())
	}
	if got := st.Slots.GetPower().GetBracket(); got != cedhBracket {
		t.Errorf("power bracket = %d, want %d", got, cedhBracket)
	}
}

// TestNoCollectionAsksTheBudget is D-168. 61 of the 63 sessions of gate
// run 18 that held no collection were never asked about the budget, and
// the eval named the budget in 20 of its unasked slots. A user with no
// library buys every card.
func TestNoCollectionAsksTheBudget(t *testing.T) {
	out := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Power = "fnm"
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "budget", "meta"), askStep(t))
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
	if !NoSpendingLimit("money is no object") {
		t.Error("money is no object was not read")
	}
	if NoSpendingLimit("the best deck under budget") {
		t.Error("a budget request read as a refusal of the cap")
	}
	out := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Power = "tournament"
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "meta"), askStep(t))
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

// TestSuperlativeDelegatesTheCommander is D-167. Conversation 10 of gate
// run 18 wrote "Buy the best lifegain commander" and got three names to
// choose from.
func TestSuperlativeDelegatesTheCommander(t *testing.T) {
	if !DelegatesCommander("Buy the best lifegain commander") {
		t.Error("a superlative commander instruction was not read as a delegation")
	}
	if DelegatesCommander("buy the best lands you can find") {
		t.Error("a message that names no commander handed the commander choice over")
	}
	first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown", BudgetUSD: 60}
	first.Colors = []string{"W", "B"}
	second := classifyOut{Format: "unknown", PoolRule: "owned_first"}
	second.Power = "bracket 3"
	a, _ := testAgentHints(t, stubHints{commanders: []string{"Karlov of the Ghost Council"}},
		classifyStep(t, first), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, second), fits(t))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "Lifegain from my collection. Commander, white and black.", nil); err != nil {
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

// ids2 names the slots of the questions one turn sent.
func ids2(qs []*mtgv1.Question) []string {
	var out []string
	for _, q := range qs {
		out = append(out, q.GetSlot())
	}
	return out
}

// TestRetiredQuestionLeavesTheAskedState is H-5. A changed format retires
// the questions that are out, and the retired key stayed in the asked
// state forever: the only ways out of it were an answer and a decline.
// The session never reported ready, and the classifier was offered the
// dead key every turn.
func TestRetiredQuestionLeavesTheAskedState(t *testing.T) {
	first := commanderClassify()
	first.Facts.WantsSuggestion = true
	change := classifyOut{Format: "modern", Theme: "lifegain", PoolRule: "unknown"}
	last := classifyOut{Format: "unknown", PoolRule: "unknown", BudgetUSD: 100}
	last.Power = "casual"
	a, _ := testAgentHints(t, stubHints{commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}},
		classifyStep(t, first), fits(t, "commander_pick", "power_commander"), askStep(t),
		classifyStep(t, change), fits(t, "power_sixty", "budget"), askStep(t),
		classifyStep(t, last), fits(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "Commander lifegain, white and black, suggest a commander", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.Asked["commander_pick"] {
		t.Fatal("the pick row did not fire on turn 1")
	}
	if _, err := a.Turn(context.Background(), st, "Actually make it Modern", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	for _, key := range openKeys(st) {
		if key == "commander_pick" {
			t.Error("the retired pick key is still offered to the classifier")
		}
	}
	res, err := a.Turn(context.Background(), st, "Casual, and 100 dollars", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if len(res.Questions) != 0 {
		t.Errorf("turn 3 asked %v, want nothing", ids2(res.Questions))
	}
	if !res.Ready {
		t.Errorf("the session is not ready after every slot filled: states %v", st.Slots.GetSlotStates())
	}
	// The retired row does not ask again: Ctx.Asked keeps its id.
	if !st.Ctx.Asked["commander_pick"] {
		t.Error("the retired row lost its asked mark, so it could ask twice")
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

// TestOneDeckClosesOnASingleDeckAnswer is one third of H-6. The one-deck
// row closed only through the classifier's closed_keys, and a user who
// answered by naming one deck left the key open forever.
func TestOneDeckClosesOnASingleDeckAnswer(t *testing.T) {
	first := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "one_deck"),
		classifyStep(t, second), fits(t, "colors", "commander", "power_commander"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "I want two decks, one Commander and one Modern.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if got := st.Slots.GetSlotStates()["deck_count"]; got != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Fatalf("deck_count state = %v, want ASKED", got)
	}
	if _, err := a.Turn(context.Background(), st, "The Commander one first, lifegain.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if got := st.Slots.GetSlotStates()["deck_count"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("deck_count state = %v, want SKIPPED: the user chose one deck", got)
	}
	if !st.Ctx.Filled["deck_count"] {
		t.Error("the one-deck key is still open")
	}
}

// TestOutOfScopeClosesOnADeckRequest is one third of H-6. The scope row
// offers "Yes, a Magic deck", and a user who answers "a Modern burn deck"
// has said the same thing. Nothing closed the key.
func TestOutOfScopeClosesOnADeckRequest(t *testing.T) {
	first := classifyOut{Format: "unknown", PoolRule: "unknown"}
	first.Facts.OutOfScope = true
	second := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	second.Colors = []string{"R"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "out_of_scope"),
		classifyStep(t, second), fits(t, "power_sixty", "budget"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "build me a yu-gi-oh deck", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if got := st.Slots.GetSlotStates()["scope"]; got != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Fatalf("scope state = %v, want ASKED", got)
	}
	if _, err := a.Turn(context.Background(), st, "a Modern red burn deck then", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if got := st.Slots.GetSlotStates()["scope"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("scope state = %v, want SKIPPED: the user asked for a Magic deck", got)
	}
	if !st.Ctx.Filled["scope"] {
		t.Error("the scope key is still open")
	}
}

// factHints answers the PR-6 facts with fixed values, so a test can prove
// what the agent reads and when.
type factHints struct {
	stubHints
	missing bool
	weak    bool
	thin    bool
	count   int
}

func (f factHints) MissingCommander([]string) bool                { return f.missing }
func (f factHints) WeakCommanderPool(string) bool                 { return f.weak }
func (f factHints) ThinTheme(string) (bool, int)                  { return f.thin, f.count }
func (f factHints) OwnedThemeCount(string) int                    { return f.count }
func (f factHints) FitsColors(string, []mtgv1.Color) (bool, bool) { return true, true }

// TestNotOwnedRowFiresForANamedCommander is M-5. The row carried the
// commander key, and any named commander filled that key, so the row
// could never fire. It now carries its own key, and a skip on that key
// leaves the named commander alone.
func TestNotOwnedRowFiresForANamedCommander(t *testing.T) {
	out := commanderClassify()
	out.PoolRule = "owned_first"
	out.CommanderNames = []string{"Karlov of the Ghost Council"}
	h := factHints{missing: true, stubHints: stubHints{commanders: []string{"Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}}}
	a, _ := testAgentHints(t, h, classifyStep(t, out), fits(t, "commander_not_owned", "power_commander"), askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "Karlov lifegain from my library first, white and black", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	q := question(res.Questions, "commander")
	if q == nil {
		t.Fatalf("the not-owned row did not fire: %v", ids2(res.Questions))
	}
	if !strings.Contains(q.GetText(), "You do not own") {
		t.Errorf("the commander question is not the not-owned row: %q", q.GetText())
	}
	if got := st.Slots.GetSlotStates()["commander"]; got != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("commander state = %v, want FILLED: the user named one", got)
	}
	if got := st.Slots.GetSlotStates()["commander_owned"]; got != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Errorf("commander_owned state = %v, want ASKED", got)
	}
}

// TestNotOwnedSkipLeavesTheCommanderFilled is the other half of M-5. When
// the library offers no alternative, D-127 skips the row, and that skip
// used to overwrite the commander the user named.
func TestNotOwnedSkipLeavesTheCommanderFilled(t *testing.T) {
	out := commanderClassify()
	out.PoolRule = "owned_first"
	out.CommanderNames = []string{"Karlov of the Ghost Council"}
	h := factHints{missing: true}
	a, _ := testAgentHints(t, h, classifyStep(t, out), fits(t, "power_commander"), askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "Karlov lifegain from my library first, white and black", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if q := question(res.Questions, "commander"); q != nil {
		t.Errorf("the not-owned row fired with no owned option to offer: %q", q.GetText())
	}
	if got := st.Slots.GetSlotStates()["commander"]; got != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("commander state = %v, want FILLED: the skip must not touch the named commander", got)
	}
	if got := st.Slots.GetSlotStates()["commander_owned"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("commander_owned state = %v, want SKIPPED", got)
	}
}

// TestFactsReadTheSlotsOfThisTurn is M-6. agentsvc read the PR-6 facts
// before the turn, so a one-message answer filled the format and the
// theme after the count was skipped. The user got the plain pool question
// and the {n} count of D-67 never showed.
func TestFactsReadTheSlotsOfThisTurn(t *testing.T) {
	out := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Power = "casual"
	h := factHints{thin: true, count: 12}
	a, _ := testAgentHints(t, h, classifyStep(t, out), fits(t, "pool_thin"), askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "Modern red burn deck from my library, casual", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if !st.Ctx.ThinTheme {
		t.Fatal("the thin-theme fact was not read after the slots filled")
	}
	q := question(res.Questions, "pool_rule")
	if q == nil {
		t.Fatalf("no pool question went out: %v", ids2(res.Questions))
	}
	if !strings.Contains(q.GetText(), "12") {
		t.Errorf("the pool question does not carry the count: %q", q.GetText())
	}
}

// TestUnsupportedSubFormatNamesNoFormat is M-9. "Duel Commander" holds
// the word "commander", and the classifier reported Commander for it.
// The format then filled, and the unsupported-format row never fired.
func TestUnsupportedSubFormatNamesNoFormat(t *testing.T) {
	if namesFormat("I play Duel Commander", mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("Duel Commander read as the Commander format")
	}
	if namesFormat("Pauper Commander please", mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("Pauper Commander read as the Commander format")
	}
	if !namesFormat("a Commander deck", mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("a plain Commander request was not read")
	}
	out := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "format_unsupported", "colors", "budget"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I play Duel Commander. A lifegain deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		t.Errorf("format = %v, want none: the message names a format this app does not build", st.Ctx.Format)
	}
	if st.UnsupportedFormatName != "Duel Commander" {
		t.Errorf("unsupported format = %q, want Duel Commander", st.UnsupportedFormatName)
	}
	if !st.Ctx.Asked["format_unsupported"] {
		t.Errorf("the unsupported-format row did not fire: %v", ids2(res.Questions))
	}
}

// TestMissingScoreKeepsTheCatalogRow covers the score call that omits a
// row. The zero value gave a fit of 0, which entered the M-4 record as
// the worst fit ever measured. A fit outside 0 to 1 is clamped.
func TestMissingScoreKeepsTheCatalogRow(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out),
		scoreStep(t, scored{RowID: "theme", Fit: 1.7, Reason: "over the top"}),
		askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "build me a deck", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	format := question(res.Questions, "format")
	if format == nil {
		t.Fatalf("the format row did not fire: %v", ids2(res.Questions))
	}
	if format.GetInvented() || format.GetGapScore() != DefaultFitThreshold {
		t.Errorf("an unscored row went out with fit %v invented %v, want the threshold and the catalog text", format.GetGapScore(), format.GetInvented())
	}
	theme := question(res.Questions, "theme")
	if theme == nil {
		t.Fatalf("the theme row did not fire: %v", ids2(res.Questions))
	}
	if theme.GetGapScore() != 1 {
		t.Errorf("fit = %v, want 1: the score is clamped", theme.GetGapScore())
	}
}

// TestColorsSurviveAnUnknownColorWord covers a classify answer with a
// word the color map does not hold. The schema refuses such a word
// today, so the test reaches apply directly. The old code wiped the
// colors before it validated the new ones, and a provider that skips the
// schema would have emptied the slot.
func TestColorsSurviveAnUnknownColorWord(t *testing.T) {
	a, _ := testAgent(t)
	st := NewState(false)
	first := commanderClassify()
	a.apply(st, first, nil, "white-black lifegain commander deck")
	second := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second.Colors = []string{"purple"}
	a.apply(st, second, nil, "purple is my favourite colour")
	if got := len(st.Slots.GetColors()); got != 2 {
		t.Errorf("colors = %v, want the two the user gave", st.Slots.GetColors())
	}
	if got := st.Slots.GetSlotStates()["colors"]; got != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("colors state = %v, want FILLED", got)
	}
}

// TestLocksCardAfterTheName extends D-166 to the verb after the name.
// Conversation 74 of run 20260826-191225-000 wrote "Sol Ring goes in it".
func TestLocksCardAfterTheName(t *testing.T) {
	if !LocksCard("A Modern burn deck. Sol Ring goes in it.", "Sol Ring") {
		t.Error("a verb after the name did not lock the card")
	}
	if LocksCard("Sol Ring, if it goes in", "Sol Ring") {
		t.Error("a verb two words after the name locked the card")
	}
}
