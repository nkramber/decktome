package main

import (
	"sort"

	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/questions"
)

// A stalled turn moved nothing at all (D-357). It sent no question, it
// did not report ready, and it left every slot exactly as it found it,
// while a question of an earlier turn was still out.
//
// Such a turn can not lead anywhere. The agent never repeats a question
// it asked, so the question that is out stays out. Readiness needs no
// key in the asked state, so no build starts. Every later turn does the
// same nothing. Session 0EqqY19J6A4BxCVgsmAE died that way on
// 2026-08-31, and the gate of run 27 passed with no sign of it.
//
// The premature check is the opposite defect: a session that calls
// itself complete with a slot unanswered. Between the two, a turn that
// ends a conversation and a turn that goes nowhere are both caught.
type stall struct {
	// Turn is the one-based turn that moved nothing.
	Turn int
	// Message is what the user said on that turn.
	Message string
	// Open are the keys whose question was out and stayed out.
	Open []string
}

// slotsOf copies the slots, so a later turn can not change what the
// comparison reads.
func slotsOf(st *questions.State) *mtgv1.Slots {
	if st == nil || st.Slots == nil {
		return nil
	}
	return proto.Clone(st.Slots).(*mtgv1.Slots)
}

// openKeys lists the keys whose question is out with no answer.
func openKeys(st *questions.State) []string {
	var out []string
	for key, state := range st.Slots.GetSlotStates() {
		if state == mtgv1.SlotState_SLOT_STATE_ASKED {
			out = append(out, key)
		}
	}
	sort.Strings(out)
	return out
}

// stalledTurn reports whether one turn moved nothing.
//
// Every argument comes from the turn that just ran. `before` and `after`
// are the slots on each side of it. A turn is stalled only when all four
// hold:
//
//   - it sent no question, so nothing new is on the table,
//   - it did not report ready, so no build follows,
//   - the slots are identical, so it filled, skipped, and changed nothing,
//   - a question of an earlier turn is still out, so the conversation is
//     waiting on an answer it can never take.
//
// A turn that fills a slot from the message alone sends no question and
// is not stalled, because the slots changed. A turn that asks something
// new is not stalled. A conversation whose script simply ends is not
// stalled either: the check reads a turn that ran, never the end of the
// message list.
func stalledTurn(questionsSent int, ready bool, before, after *mtgv1.Slots, open []string) bool {
	if questionsSent > 0 || ready {
		return false
	}
	if len(open) == 0 {
		return false
	}
	return proto.Equal(before, after)
}
