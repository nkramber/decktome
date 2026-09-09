package questions

import (
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// State is one session's question state. Slots hold the values the deck
// generator reads. Ctx holds what the planner needs: which keys are
// closed, which rows were asked, and the facts the rows trigger on.
//
// slot_states carries both proto slot names and the keys of the
// refinement rows (budget_scope, house_format_limits, named_card_role,
// commander_pick, commander_illegal). The map is free-form by design.
type State struct {
	Slots *mtgv1.Slots
	Ctx   Context
	// OptionAnswers are the options the reader picked this turn, by
	// question id and index (D-597). The caller sets them before Turn,
	// and Turn applies them after the classifier, so the reader's own
	// choice stands. They live for one turn, and no snapshot holds them.
	OptionAnswers []OptionAnswer
	// AnsweredQuestions are the ids of the questions the reader replied
	// to this turn, whatever the shape of the reply (D-599). A key
	// re-opens only when its question got a reply that did not reach the
	// slot. A question nobody answered is not stalled, and the grace of
	// CloseStalled holds it.
	AnsweredQuestions []string
	// SessionID is the conversation id. It goes out as the provider cache
	// key, so every call of one session routes together.
	SessionID string
	// NamedCards are the cards the user named, newest first. The {card}
	// placeholder reads the first one.
	NamedCards []string
	// CommanderNames are the cards the user named as the commander.
	CommanderNames []string
	// LockedNames are the cards the user named to keep, oldest first. A
	// name that is also a commander name is not a locked card (D-70).
	LockedNames []string
	// OfferedCommanders are the names the agent has retired. A retired
	// name never comes back (D-73).
	OfferedCommanders []string
	// UnsupportedFormatName is the format the user asked for that this app
	// does not build, such as "Brawl". NearestFormat is the format the app
	// builds that is closest to it (D-112).
	UnsupportedFormatName string
	NearestFormat         string
	// UnsupportedFormatAsked is the format the decline row named last. The
	// row asks again only when UnsupportedFormatName differs from it
	// (D-210).
	UnsupportedFormatAsked string
	// PreconName is the precon the user wants to upgrade, named by its
	// commander (D-113).
	PreconName string
	// SetPhrase is the set the reader named, in the reader's own words,
	// for example "the Hobbit set" (D-376). It is empty when the reader
	// named no set.
	SetPhrase string
	// UnresolvedSet is a set phrase the resolver could not settle: an
	// unknown name, or one that names two base sets. The set row asks
	// about it, and it clears when the reader answers.
	UnresolvedSet string
	// UnresolvedSetAsked is the phrase the set row named last. The row
	// asks again only when the phrase differs, which is the D-210 rule
	// for the format decline rows.
	UnresolvedSetAsked string
	// SetOptions are the set names the row offers when a phrase names two
	// or more base sets. Empty for an unknown name.
	SetOptions []string
	// SetNames are the names of the sets the deck is limited to, in
	// Slots.set_codes order. A message names them, so the reader reads
	// "The Hobbit" and never "hob".
	SetNames []string
	// PreconPhrase is what the reader wrote for the precons the deck must
	// use no card of, for example "my Avengers Assemble precon" (D-496).
	// It is empty when the reader excluded no precon.
	PreconPhrase string
	// ExcludedPreconNames names the products of Slots.exclude_precon_keys
	// in that order, so a message reads "Avengers Assemble" and never a
	// key.
	ExcludedPreconNames []string
	// UnresolvedPrecon is a precon phrase the table could not settle: a
	// name it does not hold, or one that names two products with
	// different cards. The precon row asks about it (D-496).
	UnresolvedPrecon string
	// UnresolvedPreconAsked is the phrase the precon row named last. The
	// row asks again only when the phrase differs (D-210).
	UnresolvedPreconAsked string
	// PreconOptions are the product names the precon row offers.
	PreconOptions []string
	// IllegalCommander is a card the user named as the commander that can
	// not lead a deck (D-129).
	IllegalCommander string
	// UnresolvedCommander is a commander name the card index does not
	// hold, such as "Aragorn". The commander row asks which card the
	// reader means, and it clears when the reader answers (F-75, D-606).
	UnresolvedCommander string
	// UnresolvedCommanderAsked is the name that row named last. The row
	// asks again only when the name differs, the D-210 rule of the set
	// row.
	UnresolvedCommanderAsked string
	// CommanderOptions are the card names that row offers, best first.
	// Empty for a name that matches no commander.
	CommanderOptions []string
	// CurrentOffer are the names on the table now. The pick row repeats
	// with these same names until the user asks for others, so the user
	// never reads a new list as an ignored answer (D-80).
	CurrentOffer []string
	// OfferAsked are the names the pick row sent last. The row asks again
	// only when CurrentOffer differs from these, so a turn that changes
	// nothing gets no second copy of the same question (D-163).
	OfferAsked []string
	// AskCount counts the questions the session has sent. It makes the
	// question id unique, which a repeated row would otherwise break.
	AskCount int
	// Turn counts the turns the session has run, from 1.
	Turn int
	// Messages are the user's messages, oldest first. The classify call
	// reads the last few of them, so the model sees what the user wrote
	// before and never has to guess at it (D-90).
	Messages []string
	// Asks are the M-4 records, oldest first.
	Asks []Ask
	// setsThisTurn names the sets this turn read out of the reader's
	// words. It is turn state, not session state: the Result carries it
	// out and the next turn starts with it clear, the way the commander
	// mark of D-366 does. The snapshot therefore does not hold it.
	setsThisTurn []string
	// preconsThisTurn names the products the deck uses no card of after
	// this turn, and preconsPartialThisTurn the named ones the collection
	// does not hold whole (D-497). preconsNoneThisTurn says the reader
	// excluded their precons and the collection holds none whole.
	// preconsUnavailableThisTurn says no precon table is loaded. All four
	// are turn state, as setsThisTurn is.
	preconsThisTurn            []string
	preconsPartialThisTurn     []string
	preconsNoneThisTurn        bool
	preconsUnavailableThisTurn bool
}

// PriorMessages is how many earlier messages the classify call sees.
const PriorMessages = 5

// Prior returns the last PriorMessages user messages, oldest first.
func (s *State) Prior() []string {
	if len(s.Messages) <= PriorMessages {
		return s.Messages
	}
	return s.Messages[len(s.Messages)-PriorMessages:]
}

// AddMessage keeps one user message and its words. The agent calls it
// after the classify call succeeds, so a failed turn leaves no trace.
func (s *State) AddMessage(text string) {
	s.Messages = append(s.Messages, text)
	// The words hold what the user wrote, and never a quoted question.
	// A quoted option such as "near a precon" is the agent's text, and
	// the word rules must not read it as the user's (D-280, D-292).
	s.AddWords(UserWords(text))
}

// SetOffer records the commander names now on the table.
func (s *State) SetOffer(names []string) {
	if len(names) == 0 {
		return
	}
	s.CurrentOffer = append([]string(nil), names...)
}

// RecordAskedOffer keeps the names a repeat-on-change row just sent. An
// empty list changes nothing, which is the rule SetOffer follows.
func (s *State) RecordAskedOffer(names []string) {
	if len(names) == 0 {
		return
	}
	s.OfferAsked = append([]string(nil), names...)
}

// OfferChanged reports whether the names on the table differ from the
// ones the pick row sent last. A refusal empties the table, and the
// color check of D-153 drops the names the colors exclude. Both cases
// give the user three names they have not seen.
func (s *State) OfferChanged() bool {
	if len(s.CurrentOffer) != len(s.OfferAsked) {
		return true
	}
	for i, name := range s.CurrentOffer {
		if !sameCard(name, s.OfferAsked[i]) {
			return true
		}
	}
	return false
}

// RecordAskedBadFormat keeps the format the decline row just named, so
// the next turn can tell a new format from the same one.
func (s *State) RecordAskedBadFormat() {
	s.UnsupportedFormatAsked = s.UnsupportedFormatName
}

// BadFormatChanged reports whether the format the decline row would name
// differs from the one it named last. A user who repeats the same
// unsupported format hears the same sentence, so the row stays silent
// (D-163, D-210).
func (s *State) BadFormatChanged() bool {
	now := strings.TrimSpace(s.UnsupportedFormatName)
	last := strings.TrimSpace(s.UnsupportedFormatAsked)
	return !strings.EqualFold(now, last)
}

// RetireOffer drops the names on the table and never offers them again.
// The agent calls it when the user asks for other names (D-73).
func (s *State) RetireOffer() {
	for _, name := range s.CurrentOffer {
		if name = strings.TrimSpace(name); name != "" && !hasName(s.OfferedCommanders, name) {
			s.OfferedCommanders = append(s.OfferedCommanders, name)
		}
	}
	s.CurrentOffer = nil
}

// SlotSet is the state key of the set limit (D-376).
const SlotSet = "set"

// SlotSetUnresolved is the state key of the row that asks which set a
// name means. It is its own key, not the set slot: a message can name
// two sets, resolve one, and leave the other open. The resolved set then
// fills the slot, and the row still asks about the rest (D-376).
const SlotSetUnresolved = "set_unresolved"

// SlotSetOutsideMana is the state key of the mana-fill row: may the deck
// take ramp and lands from outside the named sets (D-382)?
const SlotSetOutsideMana = "set_outside_mana"

// SlotPrecons is the state key of the precon exclusion (D-496).
const SlotPrecons = "precons"

// SlotPreconUnresolved is the state key of the row that asks which precon
// a name means. It is its own key, as SlotSetUnresolved is, so a message
// that names two products and settles one keeps the other open.
const SlotPreconUnresolved = "precon_unresolved"

// PreconRef is one product of the precon table, by key and by name.
type PreconRef struct {
	Key  string
	Name string
}

// ExcludePrecons records the products the deck uses no card of, and
// closes the precon row (D-496). An empty list closes the row too: the
// reader excluded their precons, and the collection holds none whole.
func (s *State) ExcludePrecons(phrase string, refs []PreconRef) {
	s.PreconPhrase = strings.TrimSpace(phrase)
	keys := make([]string, 0, len(refs))
	names := make([]string, 0, len(refs))
	for _, r := range refs {
		keys = append(keys, r.Key)
		names = append(names, r.Name)
	}
	s.ExcludedPreconNames = names
	s.Slots.ExcludePreconKeys = keys
	s.Ctx.PreconsExcluded = len(keys) > 0
	s.Close(SlotPrecons)
}

// PreconResolved closes the row that asks which precon a name means.
func (s *State) PreconResolved() {
	s.UnresolvedPrecon, s.PreconOptions = "", nil
	s.Ctx.PreconUnresolved = false
	s.Close(SlotPreconUnresolved)
}

// PreconUnresolved records a precon phrase the table could not settle.
// The precon row asks about it, and options names the products it offers.
func (s *State) PreconUnresolved(phrase string, options []string) {
	phrase = strings.TrimSpace(phrase)
	if phrase == "" {
		return
	}
	s.UnresolvedPrecon, s.PreconOptions = phrase, options
	s.Ctx.PreconUnresolved = true
}

// BadPreconChanged reports whether the precon row would name another
// phrase than it named last (D-210).
func (s *State) BadPreconChanged() bool {
	return !strings.EqualFold(strings.TrimSpace(s.UnresolvedPrecon), strings.TrimSpace(s.UnresolvedPreconAsked))
}

// SetLimit records the sets the reader named, and closes the set row.
// codes is the whole family, and names is what a message calls them.
func (s *State) SetLimit(phrase string, codes, names []string) {
	s.SetPhrase = strings.TrimSpace(phrase)
	s.SetNames = names
	s.Slots.SetCodes = codes
	s.Ctx.SetLimited = len(codes) > 0
	s.Close(SlotSet)
}

// SetResolved closes the row that asks which set a name means. Every
// phrase the reader gave now names a set, so the question is answered.
func (s *State) SetResolved() {
	s.UnresolvedSet, s.SetOptions = "", nil
	s.Ctx.SetUnresolved = false
	s.Close(SlotSetUnresolved)
}

// SetUnresolved records a set phrase the resolver could not settle. The
// set row asks about it, and options names the sets it may offer.
func (s *State) SetUnresolved(phrase string, options []string) {
	phrase = strings.TrimSpace(phrase)
	if phrase == "" {
		return
	}
	s.UnresolvedSet, s.SetOptions = phrase, options
	s.Ctx.SetUnresolved = true
}

// SlotCommanderUnresolved is the state key of the row that asks which
// card a commander name means. It is its own key, as SlotSetUnresolved
// is: the name the reader wrote is not the commander until the reader
// picks one card (F-75, D-606).
const SlotCommanderUnresolved = "commander_unresolved"

// CommanderResolvedName closes the row that asks which card a commander
// name means, and it records the card the reader picked.
func (s *State) CommanderResolvedName(name string) {
	s.UnresolvedCommander, s.CommanderOptions = "", nil
	s.Ctx.CommanderUnresolved, s.Ctx.CommanderNoMatch = false, false
	s.Close(SlotCommanderUnresolved)
	s.SetCommander(name)
}

// DropUnresolvedCommander closes that row with no commander. The reader
// answered "None of these", so the app offers its own three names
// (D-607).
func (s *State) DropUnresolvedCommander() {
	s.UnresolvedCommander, s.CommanderOptions = "", nil
	s.Ctx.CommanderUnresolved, s.Ctx.CommanderNoMatch = false, false
	s.Close(SlotCommanderUnresolved)
}

// CommanderUnresolved records a commander name the card index does not
// hold. options are the cards the row offers, best first, and an empty
// list marks a name that matches no commander.
func (s *State) CommanderUnresolved(name string, options []string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	}
	s.UnresolvedCommander, s.CommanderOptions = name, options
	s.Ctx.CommanderUnresolved = true
	s.Ctx.CommanderNoMatch = len(options) == 0
}

// RecordAskedCommander keeps the name that row just named.
func (s *State) RecordAskedCommander() { s.UnresolvedCommanderAsked = s.UnresolvedCommander }

// BadCommanderChanged reports whether that row would name another name
// than it named last (D-210).
func (s *State) BadCommanderChanged() bool {
	return !strings.EqualFold(strings.TrimSpace(s.UnresolvedCommander), strings.TrimSpace(s.UnresolvedCommanderAsked))
}

// RecordAskedSet keeps the set phrase the row just named.
func (s *State) RecordAskedSet() { s.UnresolvedSetAsked = s.UnresolvedSet }

// BadSetChanged reports whether the set phrase the row would name differs
// from the one it named last. A reader who repeats a name this app can
// not resolve hears the same sentence once, and not twice (D-210).
func (s *State) BadSetChanged() bool {
	return !strings.EqualFold(strings.TrimSpace(s.UnresolvedSet), strings.TrimSpace(s.UnresolvedSetAsked))
}

// commanderKeys are the state keys of the rows that ask for the commander.
// A named commander closes all of them: the pick row and the role row ask
// the same thing in other words, and the illegal row asked for a
// replacement that has now arrived (D-196).
var commanderKeys = []string{"commander", "commander_pick", "named_card_role", "commander_illegal"}

// CommanderKeys are the state keys of the rows that ask for the
// commander. A caller that reads whether the commander is settled must
// read all four: the pick row carries its own key, so a session that
// answered it leaves the plain "commander" key empty for good.
func CommanderKeys() []string { return append([]string(nil), commanderKeys...) }

// RefreshFacts reads the planner facts a FactSource answers, from the
// slots as they stand now. agentsvc calls it before the turn, and the
// agent calls it again after the classify call (D-198).
func RefreshFacts(s *State, src FactSource) {
	if src == nil || !s.Ctx.HasCollection {
		return
	}
	// The not-owned row is retired (D-226). The rules engine reports
	// ownership per card after the build, which names the exact card
	// and count and costs no turn.
	// The weak-pool row is retired (D-232). It sat on the commander key,
	// and a delegated commander fills that key, so the row could not
	// reach the user who needed it. PR-8 reports the thin pool instead.
	// PR-6 counts the on-theme owned cards (D-63). The count needs the
	// format and the theme, and it only matters while the pool key is
	// open.
	if s.Ctx.Filled["pool_rule"] || s.Slots.GetTheme() == "" ||
		s.Slots.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return
	}
	thin, _ := src.ThinTheme(s.Slots.GetTheme())
	s.Ctx.ThinTheme = thin
}

// SetCommander records a commander the user named. It closes every
// commander row, and it drops the name from the locked list.
func (s *State) SetCommander(name string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	}
	s.CommanderNames = mergeName(s.CommanderNames, name)
	s.Ctx.CommanderSet = true
	// A legal commander settles the question the illegal one raised.
	s.Ctx.CommanderIllegal, s.IllegalCommander = false, ""
	for _, key := range commanderKeys {
		s.Close(key)
	}
	// The commander fills the color slot: its color identity is the
	// deck's color identity (corpus section 11). The agent must not ask
	// for colors, and the card-pool question must not wait for an answer
	// that never comes (D-67). PR-8 reads the identity from the card.
	s.Close("colors")
}

// Reopen undoes a closed key, so its row may ask again. The slot state
// goes back to unspecified, which is what the planner reads.
func (s *State) Reopen(key string) {
	delete(s.Ctx.Filled, key)
	delete(s.Ctx.Skipped, key)
	delete(s.Ctx.Outstanding, key)
	delete(s.Slots.SlotStates, key)
}

// ReopenRows reopens a key and clears the asked mark of every row that
// owns it. A reopened key with the mark in place can not ask again: the
// planner skips an asked row that carries no repeat, so a second illegal
// commander or a reopened format got silence (D-129, D-195). A repeat
// row keeps its mark, because its own change rule decides when it asks
// again (D-163, D-210). A nil catalog keeps every mark.
func (s *State) ReopenRows(c *Catalog, key string) {
	s.Reopen(key)
	if c == nil {
		return
	}
	for _, r := range c.Rows {
		if r.StateKey() == key && !r.Repeat {
			delete(s.Ctx.Asked, r.ID)
		}
	}
}

// ClearCommander drops the commander the user chose and reopens every
// commander row. The user asked for another one, and a closed key would
// otherwise leave the agent with nothing to ask (D-130).
//
// The color slot stays closed. The user gave the colors, or the old
// commander did, and a new commander comes from inside them.
func (s *State) ClearCommander() {
	s.CommanderNames = nil
	s.Ctx.CommanderSet = false
	s.Slots.CommanderOracleIds = nil
	// The role row stays closed. The user has just said that this card
	// does not lead the deck, and reopening the row would ask about the
	// card they replaced.
	for _, key := range []string{"commander", "commander_pick"} {
		s.Reopen(key)
	}
	s.RetireOffer()
	s.Ctx.Suggested = true
}

// AddLocked records a card the user wants in the deck, and not as the
// commander. The role of that card is then settled, so the role question
// closes with it.
//
// "Build around Grist, but not as my commander" answers the role
// question before it goes out, so the row must not fire (D-70).
func (s *State) AddLocked(name string) {
	if name = strings.TrimSpace(name); name == "" {
		return
	}
	s.LockedNames = mergeName(s.LockedNames, name)
	s.Close("named_card_role")
}

// Unlock drops a card from the locked names. A revision that removes a
// card unlocks it, whatever the classifier read from the message (D-301).
func (s *State) Unlock(name string) {
	var kept []string
	for _, n := range s.LockedNames {
		if !sameCard(n, name) {
			kept = append(kept, n)
		}
	}
	s.LockedNames = kept
}

// LockedCards are the named cards that are not the commander. The build
// keeps every one of them (D-242). The locked row that once asked to
// keep or cut them is retired: it could name a list that held only the
// commander (D-70), and no function read the answer (D-260).
func (s *State) LockedCards() []string {
	var out []string
	for _, name := range s.LockedNames {
		if !hasName(s.CommanderNames, name) {
			out = append(out, name)
		}
	}
	return out
}

// hasName reports whether list holds the same card as name.
func hasName(list []string, name string) bool {
	for _, n := range list {
		if sameCard(n, name) {
			return true
		}
	}
	return false
}

// normName reads a card name for comparison, without case or edge space.
func normName(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

// baseName is the part of a card name before the first comma. For
// "Grist, the Hunger Tide" it is "grist".
func baseName(s string) string {
	if i := strings.Index(s, ","); i >= 0 {
		return strings.TrimSpace(s[:i])
	}
	return s
}

// sameCard reports whether two strings name one card. A user writes the
// full name once and the short name afterwards, and the classifier
// reports both.
//
// Only a bare first name merges with a full name. Two names that both
// hold a comma stay apart, because "Toph, Hardheaded Teacher" and "Toph,
// the Blind Bandit" are two different cards.
//
// Without the merge, the full name and the short name both reach the
// locked list, and a question stutters the card (D-70).
func sameCard(a, b string) bool {
	x, y := normName(a), normName(b)
	switch {
	case x == "" || y == "":
		return false
	case x == y:
		return true
	case !strings.Contains(x, ",") && baseName(y) == x:
		return true
	case !strings.Contains(y, ",") && baseName(x) == y:
		return true
	}
	return false
}

// mergeName adds a card name to a list. It keeps the longer form when the
// list already holds the same card, because the full name is the one the
// user can act on.
func mergeName(list []string, name string) []string {
	name = strings.TrimSpace(name)
	if name == "" {
		return list
	}
	for i, n := range list {
		if sameCard(n, name) {
			if len(name) > len(n) {
				list[i] = name
			}
			return list
		}
	}
	return append(list, name)
}

// mergeFront adds a card name to a newest-first list. The {card}
// placeholder reads the first entry, so a repeated name moves to the
// front instead of growing the list.
func mergeFront(list []string, name string) []string {
	name = strings.TrimSpace(name)
	if name == "" {
		return list
	}
	var out []string
	for _, n := range list {
		if sameCard(n, name) {
			if len(n) > len(name) {
				name = n
			}
			continue
		}
		out = append(out, n)
	}
	return append([]string{name}, out...)
}

// NewState makes an empty state for a new session.
func NewState(hasCollection bool) *State {
	return &State{
		Slots: &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{}},
		Ctx: Context{
			Filled:      map[string]bool{},
			Skipped:     map[string]bool{},
			Asked:       map[string]bool{},
			Outstanding: map[string]string{},
			// A user with no collection never gets the pool question. The
			// default is any-card (D-37).
			HasCollection: hasCollection,
		},
	}
}

// Close marks a key filled and records it in the proto. It also closes
// the newest open M-4 record for that key, which is how the report knows
// whether an invented question filled its slot.
func (s *State) Close(key string) {
	s.Ctx.Filled[key] = true
	delete(s.Ctx.Outstanding, key)
	s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_FILLED
	s.fillAsk(key)
}

// Skip marks a key the user declined. A default applies, and the agent
// does not ask again. The M-4 record counts a decline as filled: the
// question did its work, and the slot is closed.
func (s *State) Skip(key string) {
	s.Ctx.Filled[key] = true
	s.markSkipped(key)
	delete(s.Ctx.Outstanding, key)
	s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_SKIPPED
	s.fillAsk(key)
}

// markSkipped records a key that closed with no value from the reader.
// A stored snapshot from before D-631 holds no map, so it makes one.
func (s *State) markSkipped(key string) {
	if s.Ctx.Skipped == nil {
		s.Ctx.Skipped = map[string]bool{}
	}
	s.Ctx.Skipped[key] = true
}

// fillAsk closes the newest open M-4 record for one key.
func (s *State) fillAsk(key string) {
	for i := len(s.Asks) - 1; i >= 0; i-- {
		if s.Asks[i].Key == key && !s.Asks[i].Filled {
			s.Asks[i].Filled = true
			return
		}
	}
}

// MarkAsked records a question the agent sent. The gate forbids a repeat.
func (s *State) MarkAsked(rowID, key, slot string) {
	s.Ctx.Asked[rowID] = true
	if s.Slots.SlotStates[key] == mtgv1.SlotState_SLOT_STATE_UNSPECIFIED {
		s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_ASKED
	}
	if !s.Ctx.Filled[key] {
		if s.Ctx.Outstanding == nil {
			s.Ctx.Outstanding = map[string]string{}
		}
		s.Ctx.Outstanding[key] = slot
	}
}

// RetireOutstanding drops every key whose question is out with no
// answer. A changed format makes those questions meaningless, and the
// one-row-per-key rule must not block the questions that replace them
// (D-125, D-126).
//
// A retired key leaves the asked state. Ready reads that state, and the
// only other ways out of it are an answer and a decline, so a retired
// question kept the session from ever reporting ready, and it offered
// the dead key to the classifier every turn (D-195).
//
// The rows of a retired key lose their asked mark too. The theme and the
// colors have one row each, so a retired theme question could never be
// asked again, and the build ran with the slot empty. A row may ask a
// retired question once more, because the user never answered it
// (D-195). The catalog maps the keys onto the rows, and a nil catalog
// keeps the marks.
func (s *State) RetireOutstanding(c *Catalog) {
	for key := range s.Ctx.Outstanding {
		if s.Slots.GetSlotStates()[key] == mtgv1.SlotState_SLOT_STATE_ASKED {
			delete(s.Slots.SlotStates, key)
		}
		if c == nil {
			continue
		}
		for _, r := range c.Rows {
			if r.StateKey() == key {
				delete(s.Ctx.Asked, r.ID)
			}
		}
	}
	s.Ctx.Outstanding = map[string]string{}
}

// StallGrace is how many turns the reader may take to answer a question
// before the net of D-351 closes it. Two means the reader gets a whole
// turn beyond the first reply (D-386).
//
// A grace of one closed a question on the turn the reader first replied
// to it. Gate run 30 ended 32 of 107 conversations on turn 2 that way,
// and 31 conversations lost an answer the reader had given. Run 29 built
// the same conversations without the net and played them out.
const StallGrace = 2

// CloseStalled closes a question that has been out for StallGrace turns
// or more, with no value (D-351, D-386). It is the way out of a dead
// conversation: a turn that asks nothing new and is not ready can never
// move again, because the agent does not repeat a question it already
// asked. The slot takes the skipped state, not the asked state, so no
// later turn offers it again.
//
// A question that is still inside the grace period stays out. The reader
// answers it on the next turn, or the net closes it on the turn after.
//
// It returns the keys it closed and the keys it left, for the log, for
// the reader, and for the gate.
func (s *State) CloseStalled() (closed, waiting []string) {
	for key, state := range s.Slots.GetSlotStates() {
		if state != mtgv1.SlotState_SLOT_STATE_ASKED {
			continue
		}
		if s.askAge(key) < StallGrace {
			waiting = append(waiting, key)
			continue
		}
		s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_SKIPPED
		s.markSkipped(key)
		closed = append(closed, key)
		delete(s.Ctx.Outstanding, key)
	}
	sort.Strings(closed)
	sort.Strings(waiting)
	return closed, waiting
}

// ReaskStalled asks one more time for a key the reader answered and the
// classifier could not read (D-599).
//
// The caller runs it only when the planner has nothing to ask and the
// session is not ready. That pair proves the turn is stuck: the question
// is out, the no-repeat rule holds it out, and no build can start.
//
// The key re-opens once. A second stall on the same key falls to
// CloseStalled, which skips it and lets the build take the default. So
// the reader answers, hears the question once more in the same turn, and
// never waits on a chat that says nothing.
//
// It returns the keys it re-opened, for the log and the gate.
func (s *State) ReaskStalled() (reasked []string) {
	if s.Ctx.Reasked == nil {
		s.Ctx.Reasked = map[string]bool{}
	}
	answered := map[string]bool{}
	for _, id := range s.AnsweredQuestions {
		if key := s.keyOfQuestion(id); key != "" {
			answered[key] = true
		}
	}
	for key, state := range s.Slots.GetSlotStates() {
		if state != mtgv1.SlotState_SLOT_STATE_ASKED || s.Ctx.Reasked[key] || !answered[key] {
			continue
		}
		s.Ctx.Reasked[key] = true
		s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_UNSPECIFIED
		delete(s.Ctx.Filled, key)
		delete(s.Ctx.Skipped, key)
		delete(s.Ctx.Outstanding, key)
		// The no-repeat rule reads the row ids, so every row that asked
		// this key may ask it again.
		for _, a := range s.Asks {
			if a.Key == key {
				delete(s.Ctx.Asked, a.RowID)
			}
		}
		reasked = append(reasked, key)
	}
	sort.Strings(reasked)
	return reasked
}

// askAge is how many turns the reader has had to answer the newest open
// question for one key. A question sent this turn has age 0, and the
// reader's first reply reads it at age 1.
//
// A key with no M-4 record has no age this can read. It answers the
// grace, so the net still closes it: a session restored without its
// records must not stall for good.
func (s *State) askAge(key string) int {
	for i := len(s.Asks) - 1; i >= 0; i-- {
		if s.Asks[i].Key == key && !s.Asks[i].Filled {
			return s.Turn - s.Asks[i].Turn
		}
	}
	return StallGrace
}

// AddWords keeps every word the user has written. The routing rules and
// the word triggers read it.
func (s *State) AddWords(text string) {
	s.Ctx.Words = strings.TrimSpace(s.Ctx.Words + " " + strings.ToLower(text))
}

// Ready reports whether the session can build. Two things must hold: the
// planner has nothing left to ask, and no question is still out. The
// no-repeat rule empties the plan as soon as a question goes out, so the
// plan alone would call an unanswered session complete.
func (s *State) Ready(c *Catalog) bool {
	if len(c.Plan(s.Ctx)) > 0 {
		return false
	}
	return !s.Outstanding()
}

// Outstanding reports whether a question is out with no answer mapped yet.
func (s *State) Outstanding() bool {
	for _, state := range s.Slots.GetSlotStates() {
		if state == mtgv1.SlotState_SLOT_STATE_ASKED {
			return true
		}
	}
	return false
}

// slotWord normalizes a value the classifier wrote in free text. The
// format and pool-rule fields carry no schema enum, because an enum
// suppressed the field: the model answered "unknown" for a message that
// named the format outright (D-92). Go owns the vocabulary instead.
func slotWord(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	return strings.ReplaceAll(strings.ReplaceAll(s, "-", "_"), " ", "_")
}

// DefaultFormat is the format a declined format slot takes. The corpus
// names it under "default answers": Commander, the most played format in
// 2026. The slot keeps the SKIPPED state, so the deck generator can tell
// a default from a choice (D-98).
const DefaultFormat = mtgv1.FormatId_FORMAT_ID_COMMANDER

// formatIDs maps the classifier's format word to the proto enum.
var formatIDs = map[string]mtgv1.FormatId{
	"commander": mtgv1.FormatId_FORMAT_ID_COMMANDER,
	"standard":  mtgv1.FormatId_FORMAT_ID_STANDARD,
	"modern":    mtgv1.FormatId_FORMAT_ID_MODERN,
	"house":     mtgv1.FormatId_FORMAT_ID_HOUSE,
}

var poolRules = map[string]mtgv1.PoolRule{
	"owned_first": mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
	"owned_only":  mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
	"any_card":    mtgv1.PoolRule_POOL_RULE_ANY_CARD,
}

var colorIDs = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G,
}

// cedhBracket is the bracket cEDH names. Bracket 5 is competitive
// Commander, and only the ban list limits it (corpus section 2.3).
const cedhBracket = 5

var sixtySteps = map[string]mtgv1.SixtyStep{
	"casual":     mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
	"fnm":        mtgv1.SixtyStep_SIXTY_STEP_FNM,
	"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}
