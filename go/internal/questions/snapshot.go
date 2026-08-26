package questions

import mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"

// SnapshotVersion marks the shape of a stored snapshot. A reader refuses
// a version it does not know, so an old document can not become a wrong
// conversation.
const SnapshotVersion = 1

// Snapshot is the private state of one session, as data (D-74). The proto
// Session carries the slots, the turns, and the usage. It carries none of
// the facts below: which rows the agent asked, every word the user wrote,
// the card names, and the triggers the planner reads.
//
// The store holds this beside the proto session. Without it, a resumed
// conversation repeats a question the user already answered.
type Snapshot struct {
	Version           int      `json:"version"`
	Ctx               Context  `json:"ctx"`
	NamedCards        []string `json:"named_cards"`
	CommanderNames    []string `json:"commander_names"`
	LockedNames       []string `json:"locked_names"`
	OfferedCommanders []string `json:"offered_commanders"`
	CurrentOffer      []string `json:"current_offer"`
	// OfferAsked are the names the pick row sent last. A snapshot written
	// before D-163 holds none, and the row then asks once more with the
	// names on the table. That is the safe failure: the alternative is a
	// resumed session that never asks again.
	OfferAsked []string `json:"offer_asked"`
	AskCount   int      `json:"ask_count"`
	Turn       int      `json:"turn"`
	// Asks are the M-4 records of every question the session sent.
	Asks []Ask `json:"asks"`
}

// Snapshot reads the private state out of a session.
func (s *State) Snapshot() Snapshot {
	return Snapshot{
		Version:           SnapshotVersion,
		Ctx:               s.Ctx,
		NamedCards:        s.NamedCards,
		CommanderNames:    s.CommanderNames,
		LockedNames:       s.LockedNames,
		OfferedCommanders: s.OfferedCommanders,
		CurrentOffer:      s.CurrentOffer,
		OfferAsked:        s.OfferAsked,
		AskCount:          s.AskCount,
		Turn:              s.Turn,
		Asks:              s.Asks,
	}
}

// Restore rebuilds a session from its slots and its snapshot. A zero
// Snapshot gives a new state, so a session stored before the snapshot
// existed still opens.
func Restore(id string, slots *mtgv1.Slots, snap Snapshot) *State {
	st := NewState(snap.Ctx.HasCollection)
	st.SessionID = id
	if slots != nil {
		st.Slots = slots
	}
	if st.Slots.SlotStates == nil {
		st.Slots.SlotStates = map[string]mtgv1.SlotState{}
	}
	if snap.Version == 0 {
		return st
	}
	st.Ctx = snap.Ctx
	if st.Ctx.Filled == nil {
		st.Ctx.Filled = map[string]bool{}
	}
	if st.Ctx.Asked == nil {
		st.Ctx.Asked = map[string]bool{}
	}
	st.NamedCards = snap.NamedCards
	st.CommanderNames = snap.CommanderNames
	st.LockedNames = snap.LockedNames
	st.OfferedCommanders = snap.OfferedCommanders
	st.CurrentOffer = snap.CurrentOffer
	st.OfferAsked = snap.OfferAsked
	st.AskCount = snap.AskCount
	st.Turn = snap.Turn
	st.Asks = snap.Asks
	st.Refresh()
	return st
}
