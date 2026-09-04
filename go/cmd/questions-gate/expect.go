package main

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

// The golden expectations of PR-15 (slice 4). A conversation may name
// the values its slots must end with, and the gate reads each one as a
// row. A wrong slot is then a named flip, and a counted conversation
// with a miss fails the gate. The vocabulary is the one slotValues
// writes, and TestConversationExpectations pins it.
//
// Two words are special. "*" asks for any value, for a pick the pool
// made. "delegated" on the commander asks for a skipped slot with no
// commander, which is the pick the user handed to the agent (D-232). A
// theme matches when the expected words are inside the slot's text, and
// a budget with no scope word matches any scope.

// expectKeys are the keys an expectation may name, in report order.
var expectKeys = []string{"format", "colors", "power", "pool_rule", "commander", "budget", "theme", "locked", "sets"}

// slotValues renders the settled slots of a conversation as words.
// names maps an Oracle id to a card name.
func slotValues(st *questions.State, names func(string) string) map[string]string {
	out := map[string]string{}
	s := st.Slots
	if s == nil {
		return out
	}
	switch s.GetFormat().GetId() {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		out["format"] = "commander"
	case mtgv1.FormatId_FORMAT_ID_STANDARD:
		out["format"] = "standard"
	case mtgv1.FormatId_FORMAT_ID_MODERN:
		out["format"] = "modern"
	case mtgv1.FormatId_FORMAT_ID_HOUSE:
		out["format"] = "house"
	}
	out["colors"] = colorLetters(s.GetColors())
	switch p := s.GetPower().GetLevel().(type) {
	case *mtgv1.PowerLevel_Bracket:
		out["power"] = fmt.Sprintf("bracket %d", p.Bracket)
	case *mtgv1.PowerLevel_SixtyStep:
		out["power"] = strings.ToLower(strings.TrimPrefix(p.SixtyStep.String(), "SIXTY_STEP_"))
		if out["power"] == "unspecified" {
			out["power"] = ""
		}
	}
	switch s.GetPoolRule() {
	case mtgv1.PoolRule_POOL_RULE_OWNED_FIRST:
		out["pool_rule"] = "owned_first"
	case mtgv1.PoolRule_POOL_RULE_OWNED_ONLY:
		out["pool_rule"] = "owned_only"
	case mtgv1.PoolRule_POOL_RULE_ANY_CARD:
		out["pool_rule"] = "any_card"
	}
	out["commander"] = joinNames(s.GetCommanderOracleIds(), names, " + ")
	if out["commander"] == "" && s.GetSlotStates()["commander"] == mtgv1.SlotState_SLOT_STATE_SKIPPED {
		out["commander"] = "delegated"
	}
	if b := s.GetBudgetUsd(); b > 0 {
		out["budget"] = fmt.Sprintf("%g", b)
		switch s.GetBudgetScope() {
		case mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY:
			out["budget"] += " to buy"
		case mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK:
			out["budget"] += " whole deck"
		}
	}
	out["theme"] = strings.TrimSpace(s.GetTheme())
	out["locked"] = joinNames(s.GetLockedOracleIds(), names, ", ")
	out["sets"] = strings.Join(s.GetSetCodes(), ",")
	return out
}

// colorLetters writes colors in WUBRG order, and C for colorless.
func colorLetters(colors []mtgv1.Color) string {
	order := []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G, mtgv1.Color_COLOR_C}
	letters := map[mtgv1.Color]string{mtgv1.Color_COLOR_W: "W", mtgv1.Color_COLOR_U: "U", mtgv1.Color_COLOR_B: "B", mtgv1.Color_COLOR_R: "R", mtgv1.Color_COLOR_G: "G", mtgv1.Color_COLOR_C: "C"}
	have := map[mtgv1.Color]bool{}
	for _, c := range colors {
		have[c] = true
	}
	var s strings.Builder
	for _, c := range order {
		if have[c] {
			s.WriteString(letters[c])
		}
	}
	return s.String()
}

func joinNames(ids []string, names func(string) string, sep string) string {
	var out []string
	for _, id := range ids {
		if n := names(id); n != "" {
			out = append(out, n)
		} else {
			out = append(out, id)
		}
	}
	return strings.Join(out, sep)
}

// checkExpect reads the values against the expectation and names every
// miss as "key: want X, got Y", in key order.
func checkExpect(expect, values map[string]string) []string {
	var misses []string
	keys := make([]string, 0, len(expect))
	for k := range expect {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		want, got := expect[key], values[key]
		if !expectMatches(key, want, got) {
			misses = append(misses, fmt.Sprintf("%s: want %s, got %s", key, orNone(want), orNone(got)))
		}
	}
	return misses
}

func expectMatches(key, want, got string) bool {
	w, g := strings.ToLower(strings.TrimSpace(want)), strings.ToLower(strings.TrimSpace(got))
	switch {
	case w == "*":
		return g != "" && g != "delegated"
	case key == "theme":
		return w != "" && strings.Contains(g, w)
	case key == "budget":
		if g == w {
			return true
		}
		// A number with no scope word matches any scope.
		return !strings.Contains(w, " ") && strings.HasPrefix(g, w+" ")
	}
	return g == w
}

func orNone(s string) string {
	if s == "" {
		return "none"
	}
	return s
}
