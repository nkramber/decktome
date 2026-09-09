package questions

import (
	"fmt"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// SlotWords renders settled slots as the words a golden expectation uses
// (PR-15). The question gate reads them from a live state, and the
// triage of PR-28b reads them from a stored session, so the vocabulary
// has one home and the two can not drift.
//
// Three words are special. "delegated" on the commander is a skipped
// slot with no commander, which is the pick the reader handed to the
// agent (D-232). An empty color slot is how a colorless deck ends
// (D-165). A reader with no collection never gets the pool question, and
// the build reads any-card for them (D-37), so the renderer writes that
// word.
//
// The commander and the locked cards are names and not ids: the
// classifier records names, and the build resolves them to cards. The
// caller passes the locked names with the commander already out of them
// (D-70).
func SlotWords(s *mtgv1.Slots, commanderNames, lockedNames []string, hasCollection bool) map[string]string {
	out := map[string]string{}
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
	out["colors"] = ColorLetters(s.GetColors())
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
	default:
		if !hasCollection {
			out["pool_rule"] = "any_card"
		}
	}
	out["commander"] = strings.Join(commanderNames, " + ")
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
	out["locked"] = strings.Join(lockedNames, ", ")
	out["sets"] = strings.Join(s.GetSetCodes(), ",")
	return out
}

// ColorLetters writes colors in WUBRG order, and C for colorless.
func ColorLetters(colors []mtgv1.Color) string {
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
