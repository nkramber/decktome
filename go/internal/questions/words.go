package questions

import (
	"strings"
	"unicode"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The deterministic word rules. Every rule here needs no model call, so
// it costs nothing and it can not drift between runs.
//
// The rules exist because gate runs 10 to 13 of 2026-08-25 showed three
// failures that a model call alone did not catch. A word trigger fired on
// a negation ("no proxies" fired the house-rules row, probe 38). The
// classifier missed a format the user had named outright (conversation 23
// of run 11). And no rule read "not as my commander", so the agent asked
// the role question the user had already answered, in all four runs.

// negators stop a word trigger. A user who writes "no proxies" is not a
// user who proxies.
var negators = map[string]bool{
	"no": true, "not": true, "never": true, "without": true, "non": true,
	"dont": true, "doesnt": true, "cant": true, "wont": true, "isnt": true,
	"don't": true, "doesn't": true, "can't": true, "won't": true, "isn't": true,
}

// negatorWindow is how many words before a match may hold a negator.
const negatorWindow = 3

// tokens splits a text into lowercase words. An apostrophe stays, so
// "don't" is one token.
func tokens(text string) []string {
	return strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	})
}

// matchAt reports whether the phrase starts at index i of toks.
func matchAt(toks, phrase []string, i int) bool {
	if i+len(phrase) > len(toks) {
		return false
	}
	for j, w := range phrase {
		if toks[i+j] != w {
			return false
		}
	}
	return true
}

// negatedAt reports whether a negator stands before the match at index i.
//
// One shape is exempt. "Not as my commander" denies the role and not the
// format: the user wants a Commander deck, and they want that card in the
// 99. Conversation 27 asked the format in all four runs of 2026-08-25 for
// exactly this reason.
func negatedAt(toks []string, i int) bool {
	if i >= 2 && toks[i-2] == "as" {
		switch toks[i-1] {
		case "my", "the", "a":
			return false
		}
	}
	for j := i - 1; j >= 0 && j >= i-negatorWindow; j-- {
		if negators[toks[j]] {
			return true
		}
	}
	return false
}

// hasPhrase reports whether the text holds the phrase with no negator
// before it. The phrase may hold more than one word.
func hasPhrase(text, phrase string) bool {
	toks, want := tokens(text), tokens(phrase)
	if len(want) == 0 {
		return false
	}
	for i := range toks {
		if matchAt(toks, want, i) && !negatedAt(toks, i) {
			return true
		}
	}
	return false
}

// anyPhrase reports whether the text holds any of the phrases, with no
// negator before it. It replaces a plain substring test for every trigger
// a user writes, so "no proxies" no longer reads as "proxies".
func anyPhrase(text string, phrases []string) bool {
	for _, p := range phrases {
		if hasPhrase(text, p) {
			return true
		}
	}
	return false
}

// formatNames maps a format word the user may write onto the proto enum.
// "edh" is the common name for Commander.
var formatNames = map[string]mtgv1.FormatId{
	"commander": mtgv1.FormatId_FORMAT_ID_COMMANDER,
	"edh":       mtgv1.FormatId_FORMAT_ID_COMMANDER,
	"standard":  mtgv1.FormatId_FORMAT_ID_STANDARD,
	"pioneer":   mtgv1.FormatId_FORMAT_ID_PIONEER,
	"modern":    mtgv1.FormatId_FORMAT_ID_MODERN,
	"legacy":    mtgv1.FormatId_FORMAT_ID_LEGACY,
	"vintage":   mtgv1.FormatId_FORMAT_ID_VINTAGE,
	"pauper":    mtgv1.FormatId_FORMAT_ID_PAUPER,
}

// commanderSigns are phrases that name no format and mean Commander. A
// deck with a commander, a card "in the 99", a bracket, or a precon is a
// Commander deck.
var commanderSigns = []string{
	"my commander", "as my commander", "the commander", "in the 99",
	"one of the 99", "for the 99", "bracket 1", "bracket 2", "bracket 3",
	"bracket 4", "bracket 5", "precon", "preconstructed",
}

// FormatFromWords reads a format the user named or clearly implied. It is
// a safety net under the classifier, not a replacement for it: the agent
// calls it only when the classify role left the format empty.
//
// It answers nothing when the user named more than one format. Two named
// formats is a two-deck request, and OneDeckRequest reads that.
func FormatFromWords(text string) (mtgv1.FormatId, bool) {
	if _, _, ok := UnsupportedFormat(text); ok {
		return mtgv1.FormatId_FORMAT_ID_UNSPECIFIED, false
	}
	found := map[mtgv1.FormatId]bool{}
	for word, id := range formatNames {
		if hasPhrase(text, word) {
			found[id] = true
		}
	}
	if len(found) == 1 {
		for id := range found {
			return id, true
		}
	}
	if len(found) == 0 && anyPhrase(text, commanderSigns) {
		return mtgv1.FormatId_FORMAT_ID_COMMANDER, true
	}
	return mtgv1.FormatId_FORMAT_ID_UNSPECIFIED, false
}

// namesFormat reports whether one message names one format. It reads the
// message alone, and never the conversation: a value the user replaced
// stays in the conversation forever (D-125).
func namesFormat(message string, id mtgv1.FormatId) bool {
	for word, want := range formatNames {
		if want == id && hasPhrase(message, word) {
			return true
		}
	}
	// A commander phrase names the Commander format without the word.
	return id == mtgv1.FormatId_FORMAT_ID_COMMANDER && anyPhrase(message, commanderSigns)
}

// unsupported is a format this app does not build, and the nearest one it
// does. Verified 2026-08-25 against the format definitions.
//
// Brawl is 60-card singleton with a commander, so Commander is nearest.
// The owner confirmed that reading: probe 46 answers "treat it as
// Commander". Oathbreaker, Duel Commander, and Canadian Highlander are
// singleton formats with the same shape. Alchemy is Standard with the
// Arena-only rebalanced cards.
//
// The Historic and Timeless rows are marked unverified. Both are Arena
// formats with no exact paper equivalent, and the owner has not confirmed
// the nearest format for either (OQ-22).
var unsupported = []struct{ phrase, display, near string }{
	{"canadian highlander", "Canadian Highlander", "Commander"},
	{"duel commander", "Duel Commander", "Commander"},
	{"oathbreaker", "Oathbreaker", "Commander"},
	{"brawl", "Brawl", "Commander"},
	{"alchemy", "Alchemy", "Standard"},
	{"historic", "Historic", "Modern"},
	{"timeless", "Timeless", "Legacy"},
}

// UnsupportedFormat reads a format this app does not build. It returns
// the name the user wrote and the nearest format the app does build.
//
// The list is in longest-phrase order, so "duel commander" wins over
// "commander". Gate run 13 of 2026-08-25 offered Brawl to a user, which
// the app can not build (probe 46).
func UnsupportedFormat(text string) (name, near string, ok bool) {
	for _, u := range unsupported {
		if hasPhrase(text, u.phrase) {
			return u.display, u.near, true
		}
	}
	return "", "", false
}

// twoDeckSigns name a request for more than one deck outright.
var twoDeckSigns = []string{"two decks", "2 decks", "both decks", "second deck"}

// OneDeckRequest reports whether one message asks for more than one deck.
//
// It reads one message and never the whole conversation. A user who
// changes the format across two turns has not asked for two decks, and
// probe 31 does exactly that.
func OneDeckRequest(message string) bool {
	if anyPhrase(message, twoDeckSigns) {
		return true
	}
	found := map[mtgv1.FormatId]bool{}
	for word, id := range formatNames {
		if hasPhrase(message, word) {
			found[id] = true
		}
	}
	return len(found) > 1
}

// preconSigns name a preconstructed deck.
var preconSigns = []string{"precon", "preconstructed"}

// PreconRequest reports whether the user asked to upgrade a precon.
func PreconRequest(text string) bool { return anyPhrase(text, preconSigns) }

// proxySigns name a table that plays with proxies.
var proxySigns = []string{"proxy", "proxies", "proxied"}

// ProxyUser reports whether the user proxies their cards. Such a user has
// no budget, so the agent asks no budget question (D-111).
func ProxyUser(text string) bool { return anyPhrase(text, proxySigns) }

// notCommanderSigns say a named card belongs in the 99 and not in the
// command zone.
var notCommanderSigns = []string{
	"not as my commander", "not as the commander", "not my commander",
	"in the 99", "one of the 99", "for the 99", "goes in the 99",
}

// NamedCardNotCommander reports whether the user placed a named card in
// the 99. The role question must not fire after that (D-70).
//
// The negation guard is off here on purpose. Every phrase in the list
// carries its own negative sense, and "not as my commander" would read as
// negated by any general rule.
func NamedCardNotCommander(text string) bool {
	toks := tokens(text)
	for _, p := range notCommanderSigns {
		want := tokens(p)
		for i := range toks {
			if matchAt(toks, want, i) {
				return true
			}
		}
	}
	return false
}

// swapSigns ask for a commander other than the one already chosen.
var swapSigns = []string{
	"different commander", "another commander", "other commander",
	"change the commander", "swap the commander", "new commander",
	"someone else as my commander", "a different one",
}

// SwapsCommander reports whether the user wants to replace a commander
// they already chose. Probe 49 writes "Actually use a different
// commander, suggest one", and every run before 2026-08-26 asked nothing
// after it (D-130).
func SwapsCommander(message string) bool {
	toks := tokens(message)
	for _, p := range swapSigns {
		want := tokens(p)
		for i := range toks {
			if matchAt(toks, want, i) {
				return true
			}
		}
	}
	return false
}

// refusalSigns reject the names the agent put on the table. They are not
// an answer, and the pick row asks again with three others (D-73).
var refusalSigns = []string{
	"none of those", "none of them", "none of these", "not those",
	"something else", "three more", "other options", "different ones",
	"none", "neither", "no thanks",
}

// RefusedOffer reports whether the message rejects the commanders on the
// table. The caller checks that the pick row is out first, so a "none"
// about anything else reaches nothing.
//
// The negation guard is off here. Every phrase carries its own negative
// sense, and a general rule would read each one as negated.
func RefusedOffer(message string) bool {
	toks := tokens(message)
	for _, p := range refusalSigns {
		want := tokens(p)
		for i := range toks {
			if matchAt(toks, want, i) {
				return true
			}
		}
	}
	return false
}

// ordinals name a commander by its place in the offered list.
//
// Only unambiguous ordinal words are here. "One", "two", and "three" are
// left out: "Say none and I name three more" puts "three" in the answer
// of a user who chose nothing.
var ordinals = map[string]int{
	"first": 0, "1st": 0,
	"second": 1, "2nd": 1,
	"third": 2, "3rd": 2,
}

// notAPick are the words that turn an ordinal into a manner adverb.
// "Build from my library first" chooses no commander, and conversation
// 23 writes exactly that while the pick row is out.
var notAPick = map[string]bool{
	"library": true, "collection": true, "binder": true,
	"precon": true, "deck": true, "cards": true, "pool": true,
}

// OfferedPick reads a commander chosen by its place, such as "the first
// of the new three". The caller checks that the pick row is out, so an
// ordinal about anything else reaches nothing.
func OfferedPick(message string) (int, bool) {
	toks := tokens(message)
	for j, t := range toks {
		i, ok := ordinals[t]
		if !ok {
			continue
		}
		if j > 0 && notAPick[toks[j-1]] {
			continue
		}
		return i, true
	}
	return 0, false
}

// competitiveSigns ask for a strong deck without naming a power step.
var competitiveSigns = []string{
	"strongest", "competitive", "serious", "best deck", "win the event",
	"money is no object", "whatever is winning", "most powerful",
}

// CompetitiveRequest reports whether the user asked for a strong deck.
// The agent infers the tournament step from it in a 60-card format, and
// asks the user to confirm (D-107).
func CompetitiveRequest(text string) bool { return anyPhrase(text, competitiveSigns) }
