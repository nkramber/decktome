package main

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

func names(id string) string {
	return map[string]string{"o-karlov": "Karlov of the Ghost Council", "o-bond": "Sanguine Bond"}[id]
}

// TestSlotValuesWriteTheVocabulary pins the words an expectation uses.
func TestSlotValuesWriteTheVocabulary(t *testing.T) {
	st := questions.NewState(true)
	st.Slots = &mtgv1.Slots{
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Colors:             []mtgv1.Color{mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_W},
		Power:              &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}},
		PoolRule:           mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
		CommanderOracleIds: []string{"o-karlov"},
		BudgetUsd:          50,
		BudgetScope:        mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY,
		Theme:              "lifegain",
		LockedOracleIds:    []string{"o-bond"},
		SetCodes:           []string{"hob", "hoc"},
	}
	got := slotValues(st, names)
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
	got = slotValues(sixty, names)
	if got["format"] != "modern" || got["power"] != "tournament" || got["pool_rule"] != "any_card" || got["budget"] != "300" || got["commander"] != "delegated" || got["colors"] != "" {
		t.Errorf("sixty = %v", got)
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
}
