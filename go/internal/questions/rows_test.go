package questions

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The row tests: the scope row, the one-deck row, the option net, the
// house rules, and the rule table.

// TestWordRulesRunInOrder holds the order of the word rules (audit
// Q-15). A rule that moves changes what a later rule reads, so the order
// is explicit here and a change to it must change this list.
func TestWordRulesRunInOrder(t *testing.T) {
	want := []string{
		"format_from_words", "accept_nearest_format", "unsupported_format",
		"one_deck", "precon", "proxy_user", "budget_scope", "no_spending_limit",
		"buy_list", "colorless", "cedh", "house_rules", "card_in_the_99",
		"named_card_as_commander", "swap_commander", "commander_pair",
		"delegate_commander", "pick_by_place", "refuse_offer", "infer_power",
	}
	if len(wordRules) != len(want) {
		t.Fatalf("%d word rules, want %d", len(wordRules), len(want))
	}
	for i, r := range wordRules {
		if r.name != want[i] {
			t.Errorf("rule %d is %q, want %q", i, r.name, want[i])
		}
		if r.apply == nil {
			t.Errorf("rule %q has no function", r.name)
		}
	}
}

// TestOutOfScopeRepeats is audit Q-10. A user who asks for another game
// a second time heard nothing: the row was gated on its asked mark. The
// fact is raised again on the second message, so the row fires again.
func TestOutOfScopeRepeats(t *testing.T) {
	other := classifyOut{Format: "unknown", PoolRule: "unknown"}
	other.Facts.OutOfScope = true
	magic := classifyOut{Format: "commander", Theme: "dragons", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, other),
		classifyStep(t, magic), fits(t, "colors", "commander", "power_commander"), askStep(t),
		classifyStep(t, other))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "Can you build me a Yu-Gi-Oh deck?", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "Fine, Magic then. Commander, a dragon deck.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if got := st.Slots.GetSlotStates()["scope"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Fatalf("scope = %v after the user came back in scope, want SKIPPED", got)
	}
	res, err := a.Turn(context.Background(), st, "Actually, a Pokemon deck instead.", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if len(res.Questions) != 1 || res.Questions[0].GetSlot() != "scope" {
		t.Fatalf("turn 3 asked %v, want the decline alone", ids2(res.Questions))
	}
	if got := st.Slots.GetSlotStates()["scope"]; got != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Errorf("scope = %v after the second request, want ASKED", got)
	}
	if res.Ready {
		t.Error("the session reported ready with the decline out")
	}
}

// TestOneDeckRepeats is the one-deck half of audit Q-10. A user who asks
// for a second deck again hears the sentence again.
func TestOneDeckRepeats(t *testing.T) {
	unknown := classifyOut{Format: "unknown", PoolRule: "unknown"}
	one := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, unknown),
		classifyStep(t, one), fits(t, "colors", "commander", "power_commander"), askStep(t),
		classifyStep(t, unknown))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "I want two decks, one Commander and one Modern.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "The Commander one first, lifegain.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if got := st.Slots.GetSlotStates()["deck_count"]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Fatalf("deck_count = %v after the user chose one deck, want SKIPPED", got)
	}
	res, err := a.Turn(context.Background(), st, "And I want a second deck as well, a Modern one.", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if len(res.Questions) != 1 || res.Questions[0].GetSlot() != "deck_count" {
		t.Fatalf("turn 3 asked %v, want the one-deck row alone", ids2(res.Questions))
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("format = %v, want the Commander deck the user chose first", st.Ctx.Format)
	}
}

// TestPlanRepeatsTheLimitRowsOnTheFact holds the planner half of Q-10.
// The asked mark no longer gates the two limit rows.
func TestPlanRepeatsTheLimitRowsOnTheFact(t *testing.T) {
	c := load(t)
	scope := ctx(mtgv1.FormatId_FORMAT_ID_UNSPECIFIED)
	scope.OutOfScope, scope.Asked["out_of_scope"] = true, true
	if got := ids(c.Plan(scope)); len(got) != 1 || got[0] != "out_of_scope" {
		t.Errorf("a repeated out-of-scope request planned %v, want the decline", got)
	}
	decks := ctx(mtgv1.FormatId_FORMAT_ID_UNSPECIFIED)
	decks.TwoDecks, decks.Asked["one_deck"] = true, true
	if got := ids(c.Plan(decks)); len(got) != 1 || got[0] != "one_deck" {
		t.Errorf("a repeated two-deck request planned %v, want the one-deck row", got)
	}
	decks.Filled["deck_count"] = true
	if got := ids(c.Plan(decks)); len(got) == 0 || got[0] == "one_deck" {
		t.Errorf("a filled deck count still planned the one-deck row: %v", got)
	}
}

// TestCompetitiveRequestInfersTheTournamentStep is D-107. Conversation 16
// opens with "I want the strongest Modern deck, money is no object", and
// every run then asked how strong the deck should be.
func TestCompetitiveRequestInfersTheTournamentStep(t *testing.T) {
	out := classifyOut{Format: "modern", Theme: "best deck", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Facts.PowerCompetitive = true
	a, _ := testAgent(t, classifyStep(t, out), fits(t), askStep(t))
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
	// The agent states the step and asks nothing. A user who asked for the
	// strongest deck has given the answer, and the eval refused the
	// confirm row on every such conversation (D-216).
	if st.Ctx.Asked["power_sixty_confirm"] {
		t.Error("the agent asked the user to confirm a step they had given")
	}
	if st.Ctx.Asked["power_sixty"] {
		t.Error("the open power question went out on an answered slot")
	}
	// The slot is filled and closed, so no power question goes out at
	// all. The plan states the step instead (D-216).
	if q := question(res.Questions, "power"); q != nil {
		t.Fatalf("a power question went out: %q", q.GetText())
	}
}

// TestTwoDeckRequestAsksNothingElse is D-112. Probe 50 chose a deck
// silently in gate runs 11 to 13.
func TestTwoDeckRequestAsksNothingElse(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out))
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
	// No score and no ask step: a fixed row goes out as written, so a
	// turn of fixed rows alone costs one model call and not three.
	a, sc := testAgent(t, classifyStep(t, out))
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
	if len(sc.Calls) != 1 {
		t.Errorf("a fixed-only turn made %d model calls, want 1", len(sc.Calls))
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
// classifier left the key open in gate run 13 and in the batch run. The
// slot is typed since A-6 of the 2026-08-28 audit, so the repeated
// option is also the value the slot holds.
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
	if got := st.Slots.GetHouseRules(); got != "Any card, no ban list" {
		t.Errorf("house rules = %q, want the option the user repeated", got)
	}
}

// TestHouseRulesCloseOnTheUsersWords is A-6 of the 2026-08-28 audit. The
// house-rules row asked its question for 24 gate runs, and nothing
// stored the answer. The classifier now reports the user's own words,
// the slot holds them, and a bare key name can not close the slot.
func TestHouseRulesCloseOnTheUsersWords(t *testing.T) {
	first := classifyOut{Format: "modern", Theme: "dragons", PoolRule: "unknown"}
	byName := classifyOut{Format: "unknown", PoolRule: "unknown", ClosedKeys: []string{"house_rules"}}
	byValue := classifyOut{Format: "unknown", PoolRule: "unknown", HouseRules: "Modern card pool, but proxies are fine"}
	byValue.Facts.HouseFormat = true
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "house_rules", "colors"), askStep(t),
		classifyStep(t, byName), fits(t), askStep(t),
		classifyStep(t, byValue), fits(t, "house_format_limits"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Modern dragons deck, and anything goes at our table.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.Asked["house_rules"] {
		t.Fatal("the house-rules question never went out")
	}
	if _, err := a.Turn(context.Background(), st, "hmm", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if st.Ctx.Filled["house_rules"] {
		t.Error("a key name closed the house-rules slot with no value")
	}
	res, err := a.Turn(context.Background(), st, "I mean the Modern card pool, but proxies are fine.", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if got := st.Slots.GetHouseRules(); got != "Modern card pool, but proxies are fine" {
		t.Errorf("house rules = %q, want the user's words", got)
	}
	if !st.Ctx.Filled["house_rules"] {
		t.Error("the slot holds a value and stayed open")
	}
	// The limits row waits on the house-rules key, and the value frees it.
	if question(res.Questions, "house_rules") == nil {
		t.Errorf("the house-limits row did not fire: %v", ids2(res.Questions))
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
	if _, ok := c.Row("power_sixty_confirm"); ok {
		t.Error("the confirm row is back, and the eval refused it on every inference (D-216)")
	}
	out := classifyOut{Format: "modern", Theme: "dragons", PoolRule: "any_card", BudgetUSD: 100}
	out.Colors = []string{"R"}
	out.Power = "casual"
	out.Facts.HouseFormat = true
	a, _ := testAgent(t,
		classifyStep(t, out), fits(t, "house_rules"), askStep(t),
		classifyStep(t, classifyOut{Format: "unknown", PoolRule: "unknown", HouseRules: "any card, no ban list"}))
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
	if !strings.Contains(q.GetText(), "do the normal 60-card deck limits hold") {
		t.Errorf("the house-limits question did not go out as written: %q", q.GetText())
	}
}

// TestColorlessClosesTheColorSlot is D-165. The classify schema holds the
// five colors alone, so no model call can report a colorless deck. Probe
// 73 of gate run 18 answered "A colorless Commander deck" and got the
// color question.
func TestColorlessClosesTheColorSlot(t *testing.T) {
	if !colorlessRequest("a colorless Commander deck built around big artifacts") {
		t.Error("a colorless request was not read")
	}
	if colorlessRequest("not colorless, I want green") {
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

// TestOneDeckClosesOnASingleDeckAnswer is one third of H-6. The one-deck
// row closed only through the classifier's closed_keys, and a user who
// answered by naming one deck left the key open forever.
func TestOneDeckClosesOnASingleDeckAnswer(t *testing.T) {
	first := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, first),
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
		classifyStep(t, first),
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

// TestFactsReadTheSlotsOfThisTurn is M-6. agentsvc read the PR-6 facts
// before the turn, so a one-message answer filled the format and the
// theme after the count was skipped. The user got the plain pool question
// and the {n} count of D-67 never showed.
func TestFactsReadTheSlotsOfThisTurn(t *testing.T) {
	out := classifyOut{Format: "modern", Theme: "burn", PoolRule: "unknown"}
	out.Colors = []string{"R"}
	out.Power = "casual"
	h := &fakeHints{thin: true, count: 12}
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
