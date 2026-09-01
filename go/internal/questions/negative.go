package questions

import (
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// A question that invites a yes or a no takes a bare "no" as a whole
// answer: the user has no color preference, no budget, no house rule
// (D-352). The classifier reads such an answer as neither a value nor a
// decline, because a decline in its contract is a delegation, as in "you
// decide". The key then stays in the asked state for good, and no later
// turn can reopen it, because the agent never repeats a question it
// asked. Session 0EqqY19J6A4BxCVgsmAE died that way on 2026-08-31.
//
// This net is deterministic and free. It closes a key only when the
// answer is one of the words below, and only when the question that
// answer belongs to invited a yes or a no.

// bareNegatives are the whole answers that decline such a question. Each
// one carries no value: the user named no color, no cap, and no rule.
var bareNegatives = map[string]bool{
	"no": true, "nope": true, "none": true, "no thanks": true, "no thank you": true,
	"no preference": true, "not really": true, "nothing": true, "no idea": true,
	"no budget": true, "no limit": true, "no cap": true, "unlimited": true,
	"doesnt matter": true, "does not matter": true, "dont care": true, "do not care": true,
	"any": true, "anything": true, "either": true, "whatever": true, "n/a": true, "na": true,
}

// yesNoOpeners start a question that invites a yes or a no. "Which
// commander do you want?" is not one of them, and a bare negative
// answers it with nothing.
var yesNoOpeners = []string{"do you ", "did you ", "are there ", "is there ", "have you ", "would you like ", "any ", "may i ", "shall i "}

// noNegativeClose names the keys a bare negative may never close. The
// commander rows offer names, and "none of those" refuses the names
// without handing the choice back (D-120).
var noNegativeClose = map[string]bool{"commander": true, "commander_pick": true}

// normalizeAnswer lowers an answer and drops the punctuation a user
// types around it, so "No." and "Doesn't matter!" read as one word each.
func normalizeAnswer(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.NewReplacer("'", "", "’", "", ".", "", "!", "", ",", "").Replace(s)
	return strings.Join(strings.Fields(s), " ")
}

// BareNegative reports whether an answer names no value at all.
func BareNegative(answer string) bool {
	return bareNegatives[normalizeAnswer(answer)]
}

// InvitesYesNo reports whether a question takes a yes or a no as a whole
// answer.
func InvitesYesNo(question string) bool {
	q := strings.ToLower(strings.TrimSpace(question))
	if !strings.Contains(q, "?") {
		return false
	}
	for _, opener := range yesNoOpeners {
		if strings.HasPrefix(q, opener) {
			return true
		}
	}
	return false
}

// Decline closes the key of one question the user handed back outright
// (D-353). The client says so, so no words have to be read. A commander
// row may be declined: a delegation of the pick is not a refusal of the
// names on the table (D-147).
func (s *State) Decline(questionID string) (string, bool) {
	key := s.keyOfQuestion(questionID)
	if key == "" || s.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_ASKED {
		return "", false
	}
	s.Skip(key)
	return key, true
}

// keyOfQuestion reads the state key of a question the session sent.
func (s *State) keyOfQuestion(questionID string) string {
	for i := len(s.Asks) - 1; i >= 0; i-- {
		if s.Asks[i].QuestionID == questionID {
			return s.Asks[i].Key
		}
	}
	return ""
}

// DeclineNegative closes the key of one question when the user answered
// it with a bare negative (D-352). It returns the key it closed.
//
// The caller pairs the answer with the question by id, so no text has to
// be matched. A key whose question is not out closes nothing, and so
// does a key the commander rows own.
func (s *State) DeclineNegative(questionID, questionText, answer string) (string, bool) {
	if questionID == "" || !BareNegative(answer) || !InvitesYesNo(questionText) {
		return "", false
	}
	key := s.keyOfQuestion(questionID)
	if key == "" || noNegativeClose[key] {
		return "", false
	}
	if s.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_ASKED {
		return "", false
	}
	s.Skip(key)
	return key, true
}
