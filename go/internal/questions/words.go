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
	"modern":    mtgv1.FormatId_FORMAT_ID_MODERN,
}

// commanderSigns are phrases that name no format and mean Commander. A
// deck with a commander, a card "in the 99", a bracket, or a precon is a
// Commander deck.
var commanderSigns = []string{
	"my commander", "as my commander", "the commander", "in the 99",
	"one of the 99", "for the 99", "bracket 1", "bracket 2", "bracket 3",
	"bracket 4", "bracket 5", "precon", "preconstructed",
	// cEDH is competitive Commander, so it names the format as well as
	// the power level (corpus section 15, D-164).
	"cedh",
}

// FormatFromWords reads a format the user named or clearly implied. It is
// a safety net under the classifier, not a replacement for it: the agent
// calls it only when the classify role left the format empty.
//
// It reads one message. The whole conversation held every format the
// user ever named, so one "pauper" on turn 1 turned the net off for the
// rest of the session (audit Q-7). The linter reads it per message too.
//
// It answers nothing when the user named more than one format. Two named
// formats is a two-deck request, and oneDeckRequest reads that.
func FormatFromWords(text string) (mtgv1.FormatId, bool) {
	if _, _, ok := unsupportedFormat(text); ok {
		return mtgv1.FormatId_FORMAT_ID_UNSPECIFIED, false
	}
	found := map[mtgv1.FormatId]bool{}
	for word, id := range formatNames {
		if namesFormatWord(text, word) {
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

// namesFormatWord reports whether the text names a format word outright.
// A word inside a comparison names no format: "It is like Commander but
// 60 cards on Arena" describes Brawl, and gate conversation 50 wrote
// exactly that while the decline row was out. The word rule must leave
// that turn to the decline flow (audit follow-up, 2026-08-28).
func namesFormatWord(text, word string) bool {
	toks, want := tokens(text), tokens(word)
	if len(want) == 0 {
		return false
	}
	for i := range toks {
		if matchAt(toks, want, i) && !negatedAt(toks, i) && !comparedAt(toks, i, len(want)) {
			return true
		}
	}
	return false
}

// comparedAt reports whether the format word at index i sits inside a
// comparison: "like X", "similar to X", "the same as X", "based on X",
// "X but ...", "X style".
func comparedAt(toks []string, i, n int) bool {
	before := func(k int) string {
		if i-k < 0 {
			return ""
		}
		return toks[i-k]
	}
	switch {
	case before(1) == "like":
		return true
	case before(1) == "to" && before(2) == "similar":
		return true
	case before(1) == "as" && before(2) == "same":
		return true
	case before(1) == "on" && before(2) == "based":
		return true
	}
	if after := i + n; after < len(toks) {
		switch toks[after] {
		case "but", "style":
			return true
		}
	}
	return false
}

// namesFormat reports whether one message names one format. It reads the
// message alone, and never the conversation: a value the user replaced
// stays in the conversation forever (D-125).
//
// A message that names a sub-format this app does not build names no
// format here. "I play Duel Commander" holds the word "commander", and
// the classifier may report Commander for it. The unsupported-format
// row must decline it instead (M-9, D-112).
func namesFormat(message string, id mtgv1.FormatId) bool {
	if _, _, ok := unsupportedFormat(message); ok {
		return false
	}
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
// Brawl is 100-card singleton with a commander on Arena, and Standard
// Brawl is the 60-card version (owner ruling, 2026-08-26). Commander is
// nearest to both. The owner confirmed that reading: probe 46 answers
// "treat it as Commander". Oathbreaker, Duel Commander, and Canadian Highlander are
// singleton formats with the same shape. Alchemy is Standard with the
// Arena-only rebalanced cards.
//
// Historic and Timeless name no nearest format, which closes OQ-22
// (D-146). Both are Arena formats, and the card pools were measured
// against the snapshot of 2026-08-24 rather than argued from memory.
// Pioneer is the nearest of the seven by Jaccard similarity, at 0.680 for
// Historic and 0.694 for Timeless, against 0.568 and 0.581 for Modern.
// The earlier mapping to Modern and to Legacy matched neither. Historic
// and Timeless are also 0.968 similar to each other, so no measurement
// separates them. The owner chose to name no substitute over naming one
// the data does not support. An empty near sets Context.NoNearFormat, and
// a second row then asks which format to build.
// D-155 narrowed the app to Commander, Standard, and Modern. Pioneer,
// Legacy, Vintage, and Pauper joined this list and name no substitute.
//
// Pool similarity was measured for each one and it was not used. Legacy
// and Vintage read nearest to Commander at 0.994 Jaccard, because
// Commander is also an all-sets format of the same size. That number
// compares two different games: a Legacy player does not want a 100-card
// singleton multiplayer deck. Restricted to the 60-card formats the app
// keeps, all four read Modern, but only Pioneer carries its whole card
// pool (100 percent). Legacy and Vintage carry 71 percent, and Pauper
// loses the commons-only rule that defines it. The owner chose to name no
// substitute for any of the four over a claim the data does not support.
//
// The list is read in order and the first match wins, so a longer phrase
// stands before the word it holds: "pauper commander" before "pauper".
var unsupported = []struct{ phrase, display, near string }{
	{"canadian highlander", "Canadian Highlander", "Commander"},
	{"duel commander", "Duel Commander", "Commander"},
	{"pauper commander", "Pauper Commander", ""},
	{"oathbreaker", "Oathbreaker", "Commander"},
	{"brawl", "Brawl", "Commander"},
	{"alchemy", "Alchemy", "Standard"},
	{"historic", "Historic", ""},
	{"timeless", "Timeless", ""},
	{"pioneer", "Pioneer", ""},
	{"legacy", "Legacy", ""},
	{"vintage", "Vintage", ""},
	{"pauper", "Pauper", ""},
}

// unsupportedFormat reads a format this app does not build. It returns
// the name the user wrote and the nearest format the app does build.
//
// The list is in longest-phrase order, so "pauper commander" wins over
// "pauper". Gate run 13 of 2026-08-25 offered Brawl to a user, which
// the app can not build (probe 46).
func unsupportedFormat(text string) (name, near string, ok bool) {
	for _, u := range unsupported {
		if hasPhrase(text, u.phrase) {
			return u.display, u.near, true
		}
	}
	return "", "", false
}

// twoDeckSigns name a request for more than one deck outright.
var twoDeckSigns = []string{"two decks", "2 decks", "both decks", "second deck"}

// QuotedQuestionPrefix starts a line that echoes an agent question back
// with a structured answer. A client that sends Answer messages gets the
// question text joined in front of the answer, so the classifier can map
// the answer to its slot. The word rules must not read that echo: the
// format question names three formats, and oneDeckRequest read the echo
// as a request for three decks in the browser gate of 2026-08-28.
const QuotedQuestionPrefix = "Q: "

// AnswerPrefix starts the answer line under a quoted question.
const AnswerPrefix = "A: "

// UserWords returns the user's own words of a message: every line that is
// not a quoted question, with the answer prefix removed.
func UserWords(message string) string {
	lines := strings.Split(message, "\n")
	out := lines[:0]
	for _, line := range lines {
		if strings.HasPrefix(line, QuotedQuestionPrefix) {
			continue
		}
		out = append(out, strings.TrimPrefix(line, AnswerPrefix))
	}
	return strings.Join(out, "\n")
}

// oneDeckRequest reports whether one message asks for more than one deck.
//
// It reads one message and never the whole conversation. A user who
// changes the format across two turns has not asked for two decks, and
// probe 31 does exactly that.
func oneDeckRequest(message string) bool {
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

// preconRequest reports whether the user asked to upgrade a precon.
func preconRequest(text string) bool { return anyPhrase(text, preconSigns) }

// proxySigns name a table that plays with proxies.
var proxySigns = []string{"proxy", "proxies", "proxied"}

// proxyUser reports whether the user proxies their cards. Such a user has
// no budget, so the agent asks no budget question (D-111).
func proxyUser(text string) bool { return anyPhrase(text, proxySigns) }

// notCommanderSigns say a named card belongs in the 99 and not in the
// command zone.
var notCommanderSigns = []string{
	"not as my commander", "not as the commander", "not my commander",
	"in the 99", "one of the 99", "for the 99", "goes in the 99",
}

// namedCardNotCommander reports whether the user placed a named card in
// the 99. The role question must not fire after that (D-70).
//
// The negation guard is off here on purpose. Every phrase in the list
// carries its own negative sense, and "not as my commander" would read as
// negated by any general rule.
func namedCardNotCommander(text string) bool {
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

// swapsCommander reports whether the user wants to replace a commander
// they already chose. Probe 49 writes "Actually use a different
// commander, suggest one", and every run before 2026-08-26 asked nothing
// after it (D-130).
func swapsCommander(message string) bool {
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

// refusedOffer reports whether the message rejects the commanders on the
// table. The caller checks that the pick row is out first, so a "none"
// about anything else reaches nothing.
//
// The negation guard is off here. Every phrase carries its own negative
// sense, and a general rule would read each one as negated.
func refusedOffer(message string) bool {
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

// pairSigns name a request for two commanders. The Commander rules allow
// a pair through Partner, Partner with, Friends forever, "choose a
// Background", and Doctor's companion (corpus section 2.2).
//
// A pair is also the only practical way to reach four colors. WUBR, WBRG,
// and UBRG hold exactly one legal single commander each: Breya, Etherium
// Shaper, Saskia the Unyielding, and Yidris, Maelstrom Wielder (measured
// 2026-08-26 against the snapshot of 2026-08-24).
var pairSigns = []string{
	"partner", "partners", "background", "backgrounds",
	"two commanders", "2 commanders", "commander pair", "pair of commanders",
	"doctor's companion", "doctors companion", "friends forever",
	"both as commanders", "two legends",
}

// wantsCommanderPair reports whether the user asked for two commanders.
// Probe 73 writes "A Commander deck with a Background commander pair",
// and every run before D-154 answered it with three single legends.
//
// The negation guard applies: "no partners" asks for one commander.
func wantsCommanderPair(message string) bool {
	toks := tokens(message)
	for _, p := range pairSigns {
		want := tokens(p)
		for i := range toks {
			if matchAt(toks, want, i) && !negatedAt(toks, i) {
				return true
			}
		}
	}
	return false
}

// wantsBackgroundPair reports whether the user named a Background. It
// narrows a pair request: probe 73 asked for a "Background commander
// pair", and no Background ranked among the best pairs for its theme,
// because a Background carries no theme signal of its own (D-154).
func wantsBackgroundPair(message string) bool {
	toks := tokens(message)
	for _, p := range []string{"background", "backgrounds"} {
		want := tokens(p)
		for i := range toks {
			if matchAt(toks, want, i) && !negatedAt(toks, i) {
				return true
			}
		}
	}
	return false
}

// delegateSigns hand a choice to the agent outright. The user does not
// refuse the names on the table, and does not pick one: they ask the
// agent to decide.
//
// The negation guard is off here, as it is for refusalSigns. Every phrase
// carries its own sense, and a general rule would read each one as
// negated.
var delegateSigns = []string{
	"you pick", "you choose", "you decide", "you select",
	"pick for me", "choose for me", "decide for me", "select for me",
	"up to you", "your call", "your choice", "surprise me",
	"whatever you think", "whichever you think", "you know best",
	"i dunno, you pick", "dealer's choice",
}

// delegatesChoice reports whether the message hands the choice to the
// agent. The caller decides which key the answer closes.
//
// Eighteen of the 100 gate conversations hold such a phrase, and no rule
// read one before D-147. The commander pick carries "repeat": true, so
// the row asked again every turn until the messages ran out. Eval run 14
// refused 18 of its 43 bad questions on that row alone, which is more
// than the next four rows together.
func delegatesChoice(message string) bool {
	toks := tokens(message)
	for _, p := range delegateSigns {
		want := tokens(p)
		for i := range toks {
			if matchAt(toks, want, i) {
				return true
			}
		}
	}
	return false
}

// namesCommander reports whether the message uses the word at all. It
// scopes a delegation that would otherwise read as an answer to any open
// question (D-147).
func namesCommander(message string) bool {
	return hasPhrase(message, "commander")
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

// offeredPick reads a commander chosen by its place, such as "the first
// of the new three". The caller checks that the pick row is out, so an
// ordinal about anything else reaches nothing.
func offeredPick(message string) (int, bool) {
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
// The list carries the paraphrases the classifier used to catch, because
// the fact now needs the user's own words behind it (D-215).
var competitiveSigns = []string{
	"strongest", "competitive", "serious", "best deck", "win the event",
	"money is no object", "whatever is winning", "most powerful",
	"as strong as possible", "tier one", "top tier", "tournament",
}

// competitiveRequest reports whether the user asked for a strong deck.
// The agent infers the tournament step from it in a 60-card format, and
// asks the user to confirm (D-107).
func competitiveRequest(text string) bool { return anyPhrase(text, competitiveSigns) }

// occasionSigns name a place or a happening, and not a power level. A
// user who builds "for an event" has said nothing about how strong the
// deck must be.
var occasionSigns = []string{"event", "store", "lgs", "game night"}

// stepSigns name a 60-card power step outright.
var stepSigns = []string{"casual", "fnm", "friday night", "tournament", "kitchen table"}

// occasionOnly reports whether a message names an occasion and no power
// step. The classifier reads such a message as the tournament step.
// Conversation 33 of gate run 19 opened with "A Modern deck for an
// event", and the user answered "FNM level" two turns later (D-219).
func occasionOnly(message string) bool {
	return anyPhrase(message, occasionSigns) &&
		!anyPhrase(message, stepSigns) && !competitiveRequest(message)
}

// cedhSigns name competitive Commander. cEDH is bracket 5 by definition,
// and it is a Commander deck (corpus sections 2.3 and 15).
//
// "Competitive Commander" is not on the list. It says the deck is strong
// and it does not name bracket 5, and the two are three brackets apart.
var cedhSigns = []string{"cedh", "competitive edh"}

// cedhRequest reports whether the user asked for a cEDH deck. Probe 75
// of gate run 18 opens with "A cEDH deck", and the agent asked which
// power bracket to target. The user had named it (D-164).
func cedhRequest(text string) bool { return anyPhrase(text, cedhSigns) }

// buyListSigns name the cards the user must acquire. A budget beside one
// of these caps the buy list, so the scope question has its answer
// already: "Build owned-first with a buy list" says which of the two the
// cap covers (D-253).
var buyListSigns = []string{"buy list", "buylist", "to buy", "cards i buy", "cards to acquire"}

// namesTheBuyList reports whether a message names the cards to buy.
func namesTheBuyList(text string) bool { return anyPhrase(text, buyListSigns) }

// colorlessSigns name a deck with no colors.
//
// No model call can report this answer. The classify schema offers the
// five colors alone, so an empty list means "the user said nothing" and
// "the user said colorless" at the same time. Probe 73 of gate run 18
// opens with "A colorless Commander deck", and the color question went
// out (D-165).
var colorlessSigns = []string{"colorless", "no colors", "no color"}

// colorlessRequest reports whether the user asked for a colorless deck.
// The negation guard applies, so "not colorless" is not such a request.
func colorlessRequest(text string) bool { return anyPhrase(text, colorlessSigns) }

// bestSigns hand a choice to the agent with a superlative. They name no
// card, and they tell the agent to select one.
var bestSigns = []string{"the best", "the strongest", "the top"}

// delegatesCommander reports whether the message asks the agent to pick
// the commander. Conversation 10 of gate run 18 writes "Buy the best
// lifegain commander", and the agent answered with three names to choose
// from. The user had asked the agent to choose (D-167).
//
// The message must name a commander. Without that guard "the best" would
// hand over the commander choice whenever any commander question is out,
// and "buy the best lands" is not that message.
func delegatesCommander(message string) bool {
	return namesCommander(message) && anyPhrase(message, bestSigns)
}

// noBudgetSigns say the user set no spending limit. The negation guard
// is off here, as it is for refusalSigns: every phrase carries its own
// sense, and a general rule would read each one as negated.
var noBudgetSigns = []string{
	"money is no object", "price is no object", "cost is no object",
	"no budget", "no spending limit", "no price limit",
	"spend what you need", "spend whatever", "budget is no issue",
}

// noSpendingLimit reports whether the user refused a budget cap. The
// budget row must not ask a user who has answered it.
//
// It reads one message and never the whole conversation, which is the
// D-125 rule. A cap the user names later still closes the slot on its
// value.
func noSpendingLimit(message string) bool {
	toks := tokens(message)
	for _, p := range noBudgetSigns {
		want := tokens(p)
		for i := range toks {
			if matchAt(toks, want, i) {
				return true
			}
		}
	}
	return false
}

// yesWords open an answer that accepts what the agent offered.
var yesWords = map[string]bool{
	"yes": true, "yeah": true, "yep": true, "sure": true, "ok": true,
	"okay": true, "fine": true, "please": true,
}

// acceptSigns accept the offer inside a longer sentence. "Fine, treat it
// as Commander" and "use that" both take the nearest format.
var acceptSigns = []string{"use that", "that works", "treat it as", "go with that", "the nearest"}

// acceptsOffer reports whether a message says yes to the question that
// is out. It reads a leading yes-word or an acceptance phrase. The caller
// checks which question is out, so a bare "yes" reaches only the row
// that offered something (audit Q-4).
func acceptsOffer(message string) bool {
	toks := tokens(message)
	if len(toks) > 0 && yesWords[toks[0]] {
		return true
	}
	return anyPhrase(message, acceptSigns)
}

// wholeDeckSigns say a budget caps the whole deck value. The scope row
// offers "The whole deck", and the user repeats it (D-238).
var wholeDeckSigns = []string{"whole deck", "entire deck", "total deck", "deck value"}

// namesTheWholeDeck reports whether a message puts the cap on the whole
// deck.
func namesTheWholeDeck(text string) bool { return anyPhrase(text, wholeDeckSigns) }

// commanderRoleSigns give a named card the command zone. "Not as my
// commander" is negated, so it never matches here.
var commanderRoleSigns = []string{"as my commander", "as the commander", "as commander", "is my commander", "leads the deck"}

// namedCardAsCommander reports whether the user gave a named card the
// commander role. The role row offers "As my commander", and closing it
// must set the commander from the card (audit Q-14).
func namedCardAsCommander(message string) bool { return anyPhrase(message, commanderRoleSigns) }

// buysCards reports whether a pool rule lets the deck hold a card the
// user does not own. Only owned-only builds with no purchase.
func buysCards(r mtgv1.PoolRule) bool {
	switch r {
	case mtgv1.PoolRule_POOL_RULE_OWNED_FIRST, mtgv1.PoolRule_POOL_RULE_ANY_CARD:
		return true
	}
	return false
}
