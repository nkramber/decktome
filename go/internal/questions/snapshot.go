package questions

import mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"

// SnapshotVersion marks the shape of a stored snapshot. Restore reads
// every version up to this one. A field a version lacks restores as its
// zero value, and the row that reads it then asks once more, which is
// the safe failure.
//
// Version 1 carried the card lists, the offer, the counters, and the M-4
// records. Version 2 adds the format decline facts, the precon name, the
// illegal commander, and the user's messages. Production restores from
// the snapshot on every turn, so a fact that stayed in memory only worked
// in the gate harness (audit Q-1, 2026-08-28).
const SnapshotVersion = 2

// Snapshot is the private state of one session, as data (D-74). The proto
// Session carries the slots, the turns, and the usage. It carries none of
// the facts below: which rows the agent asked, every word the user wrote,
// the card names, and the triggers the planner reads.
//
// The store holds this beside the proto session. Without it, a resumed
// conversation repeats a question the user already answered.
//
// TestSnapshotRoundTrip reflects over every exported State field. A new
// State field with no counterpart here fails that test.
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
	// The format decline facts (D-112, D-210), the precon name (D-113),
	// and the illegal commander (D-129). Version 2.
	UnsupportedFormatName  string `json:"unsupported_format_name,omitempty"`
	NearestFormat          string `json:"nearest_format,omitempty"`
	UnsupportedFormatAsked string `json:"unsupported_format_asked,omitempty"`
	PreconName             string `json:"precon_name,omitempty"`
	IllegalCommander       string `json:"illegal_commander,omitempty"`
	// Messages are the user's messages, oldest first. The classify call
	// reads the last few as prior_messages. Version 2.
	Messages []string `json:"messages,omitempty"`
	AskCount int      `json:"ask_count"`
	Turn     int      `json:"turn"`
	// Asks are the M-4 records of every question the session sent.
	Asks []Ask `json:"asks"`
}

// Snapshot reads the private state out of a session.
func (s *State) Snapshot() Snapshot {
	return Snapshot{
		Version:                SnapshotVersion,
		Ctx:                    s.Ctx,
		NamedCards:             s.NamedCards,
		CommanderNames:         s.CommanderNames,
		LockedNames:            s.LockedNames,
		OfferedCommanders:      s.OfferedCommanders,
		CurrentOffer:           s.CurrentOffer,
		OfferAsked:             s.OfferAsked,
		UnsupportedFormatName:  s.UnsupportedFormatName,
		NearestFormat:          s.NearestFormat,
		UnsupportedFormatAsked: s.UnsupportedFormatAsked,
		PreconName:             s.PreconName,
		IllegalCommander:       s.IllegalCommander,
		Messages:               s.Messages,
		AskCount:               s.AskCount,
		Turn:                   s.Turn,
		Asks:                   s.Asks,
	}
}

// Restore rebuilds a session from its slots and its snapshot. A zero
// Snapshot gives a new state, so a session stored before the snapshot
// existed still opens. A version 1 snapshot restores with the version 2
// fields empty.
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
	if st.Ctx.Outstanding == nil {
		st.Ctx.Outstanding = map[string]string{}
	}
	st.NamedCards = snap.NamedCards
	st.CommanderNames = snap.CommanderNames
	st.LockedNames = snap.LockedNames
	st.OfferedCommanders = snap.OfferedCommanders
	st.CurrentOffer = snap.CurrentOffer
	st.OfferAsked = snap.OfferAsked
	st.UnsupportedFormatName = snap.UnsupportedFormatName
	st.NearestFormat = snap.NearestFormat
	st.UnsupportedFormatAsked = snap.UnsupportedFormatAsked
	st.PreconName = snap.PreconName
	st.IllegalCommander = snap.IllegalCommander
	st.Messages = snap.Messages
	st.AskCount = snap.AskCount
	st.Turn = snap.Turn
	st.Asks = snap.Asks
	return st
}
