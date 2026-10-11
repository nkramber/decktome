package questions

import (
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The reader's own words correct the classify output before it is
// applied (D-1211, D-1213, D-1224). Each rule reads one open question and
// a value the message gives word for word. A rule never replaces a value
// the classifier gave, and it costs no model call.

// skipPhrases are the words that hand one question back. The commander
// rule removes each one from the message.
var skipPhrases = [][]string{
	{"skip", "that", "question"}, {"skip", "this", "question"}, {"skip", "the", "question"},
	{"skip", "that", "one"}, {"skip", "this", "one"},
	{"skip", "that"}, {"skip", "this"}, {"skip", "it"}, {"skip"},
}

// questionSkips are the phrases that name a question. "Skip the expensive
// cards" names cards, so a bare "skip" counts only as the whole message.
var questionSkips = []string{
	"skip that question", "skip this question", "skip the question",
	"skip that one", "skip this one", "skip it",
}

// bareSkips are the whole messages that skip a question.
var bareSkips = map[string]bool{"skip": true, "skip that": true, "skip this": true}

// pickFiller are the words a reader puts around a commander name. Any
// other word leaves the role open, so the role row asks.
var pickFiller = map[string]bool{
	"i": true, "i'll": true, "i'd": true, "want": true, "take": true, "choose": true,
	"pick": true, "go": true, "with": true, "use": true, "let's": true, "the": true,
	"please": true, "my": true, "commander": true, "is": true, "will": true, "be": true,
	"as": true, "like": true, "yes": true, "ok": true, "okay": true, "and": true, "then": true,
}

// readOwnWords applies the rules that correct the classify output.
func (a *Agent) readOwnWords(st *State, out classifyOut, open []string, message string) classifyOut {
	out = a.keepPowerOption(st, out, open, message)
	return a.readTypedCommander(st, out, open, message)
}

// keepPowerOption fills the power from an option the reader typed. The
// power row offers "1 Exhibition", and the reader answered "1 Exhibition.
// Fill the gaps from any card". The classifier left the power empty, and
// it could decline the key. The build then used bracket 3 (D-1211).
func (a *Agent) keepPowerOption(st *State, out classifyOut, open []string, message string) classifyOut {
	if !slices.Contains(open, "power") || power(out.Power) != nil {
		return out
	}
	row, ok := a.askedRow(st, "power")
	if !ok {
		return out
	}
	found := ""
	for _, opt := range row.Options {
		// "Not 1 Exhibition" names no bracket.
		if powerOf(opt) == nil || !hasPhrase(message, opt) {
			continue
		}
		if found != "" {
			// Two options in one message name no single answer.
			return out
		}
		found = opt
	}
	if found == "" {
		return out
	}
	a.log.Info("the reader typed a power option, so the power holds it",
		"session", st.SessionID, "option", found)
	out.Power = powerWord(found)
	out.DeclinedKeys = withoutKey(out.DeclinedKeys, "power")
	return out
}

// readTypedCommander makes a typed name the commander when the commander
// question is out. The reader answered the pick row with "Skip that
// question.. Jace, Multiverse Architect". The classifier put the name in
// named_cards and declined the commander, so the role row asked what the
// reader had answered (D-1211, D-1213).
//
// The name must be the one card the message names, and it must lead a
// deck. The message holds no other word than the skip words and the
// filler words, so "not Jace" or "Jace in the 99" keeps the role row.
func (a *Agent) readTypedCommander(st *State, out classifyOut, open []string, message string) classifyOut {
	if !slices.Contains(open, "commander") && !slices.Contains(open, "commander_pick") {
		return out
	}
	if st.Ctx.CommanderSet || len(out.CommanderNames) > 0 || len(out.LockedNames) > 0 {
		return out
	}
	var names []string
	for _, n := range out.NamedCards {
		if n = strings.TrimSpace(n); n != "" && !hasName(names, n) {
			names = append(names, n)
		}
	}
	if len(names) != 1 {
		return out
	}
	name := names[0]
	ck, ok := a.hints.(CommanderChecker)
	if !ok {
		return out
	}
	if lead, known := ck.CanLead(name); !known || !lead {
		return out
	}
	rest, ok := withoutPhrase(tokens(message), tokens(name))
	if !ok {
		return out
	}
	for _, p := range skipPhrases {
		rest, _ = withoutPhrase(rest, p)
	}
	for _, w := range rest {
		if !pickFiller[w] {
			return out
		}
	}
	a.log.Info("the reader typed a commander under the commander question, so it leads the deck",
		"session", st.SessionID, "card", name)
	out.CommanderNames = []string{name}
	out.NamedCards = withoutName(out.NamedCards, name)
	out.Facts.NamedCard = len(out.NamedCards) > 0
	out.DeclinedKeys = withoutKey(withoutKey(out.DeclinedKeys, "commander"), "commander_pick")
	return out
}

// declineSkippedQuestion declines the one question a skip leaves out. The
// reader answered two questions in one message: "Skip that question..
// Jace, Multiverse Architect". The name answered the pick row, and the
// classifier declined nothing, so the budget question stayed out. The
// agent asks no question twice, so the chat stalled (D-1224).
//
// It runs after every other rule. It acts only when one question of the
// turn is still out, so the skip can name no other.
func (a *Agent) declineSkippedQuestion(st *State, classifierDeclined []string, open []string, message string) {
	if len(classifierDeclined) > 0 || !skipsQuestion(message) {
		return
	}
	var left []string
	for _, k := range open {
		if st.Slots.GetSlotStates()[k] == mtgv1.SlotState_SLOT_STATE_ASKED {
			left = append(left, k)
		}
	}
	if len(left) != 1 {
		return
	}
	a.log.Info("the reader skipped the one question still out", "session", st.SessionID, "key", left[0])
	st.DeclineKey(left[0])
}

// skipsQuestion reports whether a message holds a skip phrase that no
// negator comes before.
func skipsQuestion(message string) bool {
	return anyPhrase(message, questionSkips) || bareSkips[strings.Join(tokens(message), " ")]
}

// powerOf reads the power level of an option of a power row.
func powerOf(opt string) *mtgv1.PowerLevel { return power(powerWord(opt)) }

// powerWord is the word of an option that power reads. "Tournament-meta"
// is the step "tournament".
func powerWord(opt string) string {
	if power(opt) != nil {
		return opt
	}
	head, _, _ := strings.Cut(strings.ToLower(strings.TrimSpace(opt)), "-")
	return head
}

// withoutPhrase removes the first match of phrase from toks. It reports
// whether toks held the phrase.
func withoutPhrase(toks, phrase []string) ([]string, bool) {
	if len(phrase) == 0 {
		return toks, false
	}
	for i := range toks {
		if matchAt(toks, phrase, i) {
			return slices.Concat(toks[:i], toks[i+len(phrase):]), true
		}
	}
	return toks, false
}

// withoutKey drops one key from a list of keys.
func withoutKey(keys []string, key string) []string {
	return slices.DeleteFunc(slices.Clone(keys), func(k string) bool { return strings.TrimSpace(k) == key })
}
