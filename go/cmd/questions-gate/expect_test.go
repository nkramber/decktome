package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/harvest"
	"github.com/nkramber/decktome/go/internal/questions"
	"github.com/nkramber/decktome/go/internal/triage"
)

// TestSlotValuesWriteTheVocabulary pins the words an expectation uses.
func TestSlotValuesWriteTheVocabulary(t *testing.T) {
	st := questions.NewState(true)
	st.Slots = &mtgv1.Slots{
		Format:      &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Colors:      []mtgv1.Color{mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_W},
		Power:       &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}},
		PoolRule:    mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
		BudgetUsd:   50,
		BudgetScope: mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY,
		Theme:       "lifegain",
		SetCodes:    []string{"hob", "hoc"},
	}
	// The names live on the state, and a locked name that is also the
	// commander is not a locked card (D-70).
	st.CommanderNames = []string{"Karlov of the Ghost Council"}
	st.LockedNames = []string{"Karlov of the Ghost Council", "Sanguine Bond"}
	got := slotValues(st)
	want := map[string]string{
		"format": "commander", "colors": "WB", "power": "bracket 3", "pool_rule": "owned_first",
		"commander": "Karlov of the Ghost Council", "budget": "50 to buy", "theme": "lifegain",
		"locked": "Sanguine Bond", "sets": "hob,hoc",
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s = %q, want %q", k, got[k], v)
		}
	}
	sixty := questions.NewState(false)
	sixty.Slots = &mtgv1.Slots{
		Format:     &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN},
		Power:      &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT}},
		PoolRule:   mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		BudgetUsd:  300,
		SlotStates: map[string]mtgv1.SlotState{"commander": mtgv1.SlotState_SLOT_STATE_SKIPPED},
	}
	got = slotValues(sixty)
	if got["format"] != "modern" || got["power"] != "tournament" || got["pool_rule"] != "any_card" || got["budget"] != "300" || got["commander"] != "delegated" || got["colors"] != "" {
		t.Errorf("sixty = %v", got)
	}
	// A reader with no collection never gets the pool question, and the
	// build reads any-card (D-37). A reader with one and no answer reads
	// none.
	if got := slotValues(questions.NewState(false)); got["pool_rule"] != "any_card" {
		t.Errorf("no collection = %q, want any_card", got["pool_rule"])
	}
	if got := slotValues(questions.NewState(true)); got["pool_rule"] != "" {
		t.Errorf("a collection and no answer = %q, want none", got["pool_rule"])
	}
	partners := questions.NewState(true)
	partners.CommanderNames = []string{"Thrasios, Triton Hero", "Tymna the Weaver"}
	if got := slotValues(partners); got["commander"] != "Thrasios, Triton Hero + Tymna the Weaver" {
		t.Errorf("partners = %q", got["commander"])
	}
}

// TestCheckExpectNamesTheMiss: an exact key misses on a different value,
// a theme matches inside the text, a budget with no scope matches any
// scope, "*" wants any pick, and "delegated" wants the skipped slot.
func TestCheckExpectNamesTheMiss(t *testing.T) {
	values := map[string]string{
		"format": "commander", "colors": "WB", "power": "bracket 3", "commander": "Karlov of the Ghost Council",
		"budget": "50 to buy", "theme": "a lifegain deck with drain",
	}
	misses := checkExpect(map[string]string{
		"format": "commander", "colors": "WB", "power": "bracket 2", "commander": "*",
		"budget": "50", "theme": "lifegain", "pool_rule": "owned_first",
	}, values)
	if len(misses) != 2 || !strings.HasPrefix(misses[0], "pool_rule: want owned_first, got none") || misses[1] != "power: want bracket 2, got bracket 3" {
		t.Errorf("misses = %v", misses)
	}
	if m := checkExpect(map[string]string{"budget": "50 whole deck"}, values); len(m) != 1 {
		t.Errorf("a scope word must match: %v", m)
	}
	delegated := map[string]string{"commander": "delegated"}
	if m := checkExpect(map[string]string{"commander": "delegated"}, delegated); len(m) != 0 {
		t.Errorf("delegated matches delegated: %v", m)
	}
	if m := checkExpect(map[string]string{"commander": "*"}, delegated); len(m) != 1 {
		t.Errorf("a star wants a pick, and delegated is none: %v", m)
	}
	if m := checkExpect(nil, values); len(m) != 0 {
		t.Errorf("no expectation, no miss: %v", m)
	}
	// "none" wants an empty slot, the end of a colorless deck (D-165).
	if m := checkExpect(map[string]string{"colors": "none"}, map[string]string{"colors": ""}); len(m) != 0 {
		t.Errorf("none matches an empty slot: %v", m)
	}
	if m := checkExpect(map[string]string{"colors": "none"}, values); len(m) != 1 || m[0] != "colors: want none, got WB" {
		t.Errorf("none misses a filled slot: %v", m)
	}
}

// TestAForbiddenRowThatFiresIsAMiss reads the "must not ask" expectation
// of PR-28b. A row the conversation forbids fails the gate when it
// fires, and one that never fires reads clean.
func TestAForbiddenRowThatFiresIsAMiss(t *testing.T) {
	asks := []asked{
		{Turn: 1, Row: "format"},
		{Turn: 2, Row: "power_commander"},
		{Turn: 3, Row: "power_commander"},
	}
	got := checkNotAsked([]string{"power_commander"}, asks)
	if len(got) != 1 {
		t.Fatalf("misses = %v, want one", got)
	}
	// The first firing is the miss, and a row that fires twice is still
	// one miss: the report names the row and not the count.
	if !strings.Contains(got[0], "power_commander") || !strings.Contains(got[0], "turn 2") {
		t.Errorf("miss = %q, want the row and turn 2", got[0])
	}
	if got := checkNotAsked([]string{"budget"}, asks); len(got) != 0 {
		t.Errorf("a row that never fired gave %v, want none", got)
	}
	if got := checkNotAsked(nil, asks); len(got) != 0 {
		t.Errorf("no expectation gave %v, want none", got)
	}
}

// TestAForbiddenRowMustExist refuses a row id the catalog does not hold.
// A typo would make an expectation that can never fail, and the gate
// would spend a run to learn nothing.
func TestAForbiddenRowMustExist(t *testing.T) {
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	good := []conversation{{ID: 1, MustNotAsk: []string{"power_commander", "budget"}}}
	if err := checkRowIDs(good, cat); err != nil {
		t.Errorf("real rows were refused: %v", err)
	}
	bad := []conversation{{ID: 7, MustNotAsk: []string{"no_such_row"}}}
	err = checkRowIDs(bad, cat)
	if err == nil {
		t.Fatal("a row the catalog does not hold was accepted")
	}
	if !strings.Contains(err.Error(), "no_such_row") || !strings.Contains(err.Error(), "7") {
		t.Errorf("error = %v, want the row and the conversation", err)
	}
}

// TestEveryForbiddenRowOfTheGateFileExists reads the shipped file, so a
// bad row id fails a free test and never a paid run.
func TestEveryForbiddenRowOfTheGateFileExists(t *testing.T) {
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	var file gateFile
	if err := json.Unmarshal(conversationsJSON, &file); err != nil {
		t.Fatal(err)
	}
	if err := checkRowIDs(file.Conversations, cat); err != nil {
		t.Error(err)
	}
}

// TestTheCaseShapeMatchesTheGateFile pins the mirror struct of the
// triage against this file's own struct (PR-28b). The triage writes a
// conversation from a reader's verdict, and it can not import this
// package, so it holds a mirror of the shape. A field renamed here and
// not there would write the old name, and the case would lose that
// field with nothing to say so.
//
// The decoder refuses an unknown field, so the mirror may name no field
// this struct does not read.
func TestTheCaseShapeMatchesTheGateFile(t *testing.T) {
	rec := harvest.Record{
		ID: "f1", Kind: "question", Verdict: "down",
		Reasons: []string{"already_answered"}, Text: "You asked me that already.",
		QuestionID: "q3-power_commander",
		Session: json.RawMessage(`{"id":"s1","collectionId":"c1",` +
			`"slots":{"format":{"id":"FORMAT_ID_COMMANDER"},"theme":"lifegain"},` +
			`"turns":[{"userMessage":"Build me a lifegain deck."}]}`),
	}
	c, err := triage.CaseOf(triage.RouteOf(rec), 999, nil)
	if err != nil {
		t.Fatal(err)
	}
	if c.Target != "go/cmd/questions-gate/conversations.json" {
		t.Fatalf("the case joins %q, and this test guards another file", c.Target)
	}
	dec := json.NewDecoder(bytes.NewReader(c.Body))
	dec.DisallowUnknownFields()
	var got conversation
	if err := dec.Decode(&got); err != nil {
		t.Fatalf("the triage wrote a field this gate does not read: %v\n%s", err, c.Body)
	}
	// Every field the mirror fills must land, so a renamed tag fails here
	// and not in a paid run.
	if got.ID != 999 || got.Name == "" || got.Note == "" || !got.Collection {
		t.Errorf("the case lost a field: %+v", got)
	}
	if len(got.Messages) != 1 || got.Messages[0] != "Build me a lifegain deck." {
		t.Errorf("messages = %v", got.Messages)
	}
	if got.Expect["theme"] != "lifegain" || got.Expect["format"] != "commander" {
		t.Errorf("expect = %v", got.Expect)
	}
	if len(got.MustNotAsk) != 1 || got.MustNotAsk[0] != "power_commander" {
		t.Errorf("must_not_ask = %v", got.MustNotAsk)
	}
}
