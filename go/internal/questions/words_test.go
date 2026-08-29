package questions

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The pure word rules. Every test here calls a rule with a string and
// reads the answer, and none of them runs a turn. The turn tests live by
// subject: commander_test.go, format_test.go, budget_test.go, and
// rows_test.go.

// TestNegationStopsATriggerWord is D-111. A plain substring test reads
// "no proxies" as a proxy user.
func TestNegationStopsATriggerWord(t *testing.T) {
	cases := []struct {
		text, phrase string
		want         bool
	}{
		{"we proxy everything at our table", "proxy", true},
		{"edh gruul dino stompy pls, no proxies", "proxies", false},
		{"i do not want proxies", "proxies", false},
		{"any card, no ban list", "no ban list", true},
		{"i want a commander deck", "commander", true},
		// "Not as my commander" denies the role of a card and not the
		// format. The user still wants a Commander deck.
		{"build around grist, but not as my commander", "commander", true},
		{"i don't want commander", "commander", false},
	}
	for _, tc := range cases {
		if got := hasPhrase(tc.text, tc.phrase); got != tc.want {
			t.Errorf("hasPhrase(%q, %q) = %v, want %v", tc.text, tc.phrase, got, tc.want)
		}
	}
}

// TestFormatFromWords covers the safety net under the classifier. Every
// case below is a first message of a gate conversation.
func TestFormatFromWords(t *testing.T) {
	const none = mtgv1.FormatId_FORMAT_ID_UNSPECIFIED
	cases := map[string]mtgv1.FormatId{
		// The format is named as an adjective (D-116).
		"a land destruction commander deck": mtgv1.FormatId_FORMAT_ID_COMMANDER,
		// "Not as my commander" names the format without the word (D-116).
		"build around grist, the hunger tide, but not as my commander":  mtgv1.FormatId_FORMAT_ID_COMMANDER,
		"i want the strongest modern deck, money is no object":          mtgv1.FormatId_FORMAT_ID_MODERN,
		"upgrade my atraxa, praetors' voice precon. we play bracket 2":  mtgv1.FormatId_FORMAT_ID_COMMANDER,
		"edh gruul dino stompy pls, no proxies":                         mtgv1.FormatId_FORMAT_ID_COMMANDER,
		"i already have a modern burn deck. i only need sideboard help": mtgv1.FormatId_FORMAT_ID_MODERN,
		// D-155: Pauper is not a format this app builds, so the word rule
		// answers nothing and the unsupported row declines it.
		"a pauper burn deck, as cheap as possible": none,
		"a pioneer aggro deck":                     none,
		"a legacy delver deck":                     none,
		"a vintage shops deck":                     none,
		// Two formats is a two-deck request, and the one-deck row reads it.
		"i want two decks, one commander and one modern": none,
		// A format this app does not build answers nothing here.
		"i want a brawl deck for arena": none,
		// A comparison names no format (D-112).
		"it is like commander but 60 cards on arena": none,
		"something similar to modern":                none,
		"the same as standard, more or less":         none,
		"a commander-style deck for arena":           none,
		"based on commander":                         none,
		"build me a lifegain deck":                   none,
		"make me something fun and janky":            none,
	}
	for text, want := range cases {
		got, ok := FormatFromWords(text)
		if want == none {
			if ok {
				t.Errorf("FormatFromWords(%q) = %v, want no answer", text, got)
			}
			continue
		}
		if !ok || got != want {
			t.Errorf("FormatFromWords(%q) = %v (%v), want %v", text, got, ok, want)
		}
	}
}

// TestUnsupportedFormat is D-112. This app can not build Brawl, and it
// must not offer it.
func TestUnsupportedFormat(t *testing.T) {
	cases := []struct{ text, name, near string }{
		{"i want a brawl deck for arena", "Brawl", "Commander"},
		{"a duel commander deck please", "Duel Commander", "Commander"},
		{"an oathbreaker deck", "Oathbreaker", "Commander"},
		{"a commander deck with lightning bolt", "", ""},
		{"a modern burn deck", "", ""},
	}
	for _, tc := range cases {
		name, near, ok := unsupportedFormat(tc.text)
		if tc.name == "" {
			if ok {
				t.Errorf("unsupportedFormat(%q) = %q, want none", tc.text, name)
			}
			continue
		}
		if !ok || name != tc.name || near != tc.near {
			t.Errorf("unsupportedFormat(%q) = %q/%q, want %q/%q", tc.text, name, near, tc.name, tc.near)
		}
	}
}

// TestOneDeckRequestReadsOneMessage is D-112. A user who changes the
// format across two turns has not asked for two decks.
func TestOneDeckRequestReadsOneMessage(t *testing.T) {
	cases := map[string]bool{
		"i want two decks, one commander and one modern": true,
		"i want a second deck as well":                   true,
		"start with the commander one. a dragon deck":    false,
		"actually, make it modern instead. sixty cards":  false,
		"i want a commander deck":                        false,
	}
	for text, want := range cases {
		if got := oneDeckRequest(text); got != want {
			t.Errorf("oneDeckRequest(%q) = %v, want %v", text, got, want)
		}
	}
}

// TestSameCardMergesOnlyAShortName guards the fix for the stuttered name.
// Two names that both hold a comma are two different cards.
func TestSameCardMergesOnlyAShortName(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"Grist", "Grist, the Hunger Tide", true},
		{"Grist, the Hunger Tide", "grist", true},
		{"Brago, King Eternal", "Brago", true},
		{"Toph, Hardheaded Teacher", "Toph, the Blind Bandit", false},
		{"Sanguine Bond", "Grist", false},
		{"", "Grist", false},
	}
	for _, tc := range cases {
		if got := sameCard(tc.a, tc.b); got != tc.want {
			t.Errorf("sameCard(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

// TestLockedNamesDoNotStutter is D-70. Without the merge a question
// reads "Must the deck keep Grist, the Hunger Tide and Grist, or may I
// cut a card that does not fit the plan?"
func TestLockedNamesDoNotStutter(t *testing.T) {
	st := NewState(true)
	st.AddLocked("Grist, the Hunger Tide")
	st.AddLocked("Grist")
	if got := englishList(st.LockedCards()); got != "Grist, the Hunger Tide" {
		t.Errorf("locked cards = %q, want the full name once", got)
	}
	// A second, different card still joins the list.
	st.AddLocked("Sanguine Bond")
	if got := englishList(st.LockedCards()); got != "Grist, the Hunger Tide and Sanguine Bond" {
		t.Errorf("locked cards = %q, want both cards", got)
	}
}

// TestOptionAnsweredReadsThePrefix covers the shape of a real answer. An
// option reads "Yes, the normal limits" and the user writes "The normal
// limits hold".
func TestOptionAnsweredReadsThePrefix(t *testing.T) {
	cases := []struct {
		msg, opt string
		want     bool
	}{
		{"any card, no ban list. call it modern. a dragon deck.", "Any card, no ban list", true},
		{"the normal limits hold. casual power, red and green.", "Yes, the normal limits", true},
		{"grist goes in the 99. bracket 3.", "In the 99", true},
		// The classifier owns this one. A commander name closes the row by
		// value, and the option net never sees it (D-83).
		{"karlov is my commander", "As my commander", false},
		// A bare yes answers any question, so it answers none of them.
		{"yes", "Yes, the normal limits", false},
		{"no", "No, I will say the limits", false},
		{"bracket 3, and build from my library first", "Vintage rules", false},
	}
	for _, tc := range cases {
		if got := optionAnswered(tc.msg, tc.opt); got != tc.want {
			t.Errorf("optionAnswered(%q, %q) = %v, want %v", tc.msg, tc.opt, got, tc.want)
		}
	}
}

// TestOfferedPickReadsOnlyAClearOrdinal keeps "three more" out. A user
// who refuses every name writes the word "three", and that is not a pick.
func TestOfferedPickReadsOnlyAClearOrdinal(t *testing.T) {
	cases := map[string]int{
		"the first of the new three is good":         0,
		"The second one is good.":                    1,
		"give me the 3rd":                            2,
		"none of those, name three more":             -1,
		"bracket 3, and build from my library first": -1,
	}
	for msg, want := range cases {
		got, ok := offeredPick(msg)
		if want < 0 {
			if ok {
				t.Errorf("offeredPick(%q) = %d, want no pick", msg, got)
			}
			continue
		}
		if !ok || got != want {
			t.Errorf("offeredPick(%q) = %d (%v), want %d", msg, got, ok, want)
		}
	}
}

// TestDelegationNeedsTheCommanderInScope guards D-147 against a general
// "up to you". The phrase closes the commander choice only while a
// commander question is out, or while the message names one.
func TestDelegationNeedsTheCommanderInScope(t *testing.T) {
	if !delegatesChoice("You pick the commander.") {
		t.Error("a plain delegation was not read")
	}
	if !delegatesChoice("I dunno, you pick.") {
		t.Error("conversation 35 answers this, and it was not read")
	}
	if !delegatesChoice("You decide.") {
		t.Error("a delegation with no subject was not read")
	}
	if delegatesChoice("I will pick the commander myself.") {
		t.Error("the user kept the choice, and the rule took it")
	}
	if !namesCommander("You pick the commander.") {
		t.Error("the scope guard missed the word")
	}
	if namesCommander("Up to you.") {
		t.Error("the scope guard read a word that is not there")
	}
}

// TestWantsCommanderPair is D-154. "A Commander deck with a Background
// commander pair" is not a request for one commander.
func TestWantsCommanderPair(t *testing.T) {
	want := []string{
		"A Commander deck with a Background commander pair.",
		"A Commander deck with Thrasios, Triton Hero and Tymna the Weaver as partners.",
		"I want two commanders.",
		"Build a four-color deck with partners.",
		"Give me a Doctor's companion deck.",
	}
	for _, m := range want {
		if !wantsCommanderPair(m) {
			t.Errorf("a pair request was not read: %q", m)
		}
	}
	// The negation guard applies, and an unrelated message names none.
	notWant := []string{
		"No partners, just one commander.",
		"A lifegain Commander deck from my library.",
		"Bracket 3, and 150 dollars.",
	}
	for _, m := range notWant {
		if wantsCommanderPair(m) {
			t.Errorf("a message that asks for no pair was read as one: %q", m)
		}
	}
	// A Background request narrows the pair, because a Background carries
	// no theme signal and loses on score without this.
	if !wantsBackgroundPair("A Commander deck with a Background commander pair.") {
		t.Error("a Background request was not read")
	}
	if wantsBackgroundPair("A Commander deck with partners.") {
		t.Error("a plain partner request was read as a Background request")
	}
}

// TestOfferChangedReadsTheTable proves what the planner reads: the names
// on the table against the names the row sent last.
func TestOfferChangedReadsTheTable(t *testing.T) {
	st := NewState(false)
	names := []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}
	st.SetOffer(names)
	st.RecordAskedOffer(names)
	if st.OfferChanged() {
		t.Error("an unchanged table reads as changed")
	}
	st.RetireOffer()
	if !st.OfferChanged() {
		t.Error("a refusal left the table unchanged")
	}
	// A dropped name is the D-153 case: the colors arrived and one name
	// no longer fits.
	st.SetOffer(names)
	st.RecordAskedOffer(names)
	st.CurrentOffer = names[:2]
	if !st.OfferChanged() {
		t.Error("a dropped name left the table unchanged")
	}
}

// TestDeclineRowAsksAgainForAnotherFormat is the other half. A second
// unsupported format is a new sentence, so the row asks again.
func TestDeclineRowAsksAgainForAnotherFormat(t *testing.T) {
	st := NewState(false)
	st.UnsupportedFormatName = "Legacy"
	st.RecordAskedBadFormat()
	if st.BadFormatChanged() {
		t.Fatal("the same format read as a change")
	}
	st.UnsupportedFormatName = "Vintage"
	if !st.BadFormatChanged() {
		t.Fatal("another format did not read as a change")
	}
}

// TestAnEventIsNotACompetitiveRequest is D-215. "A Modern deck for an
// event" carries no power level.
func TestAnEventIsNotACompetitiveRequest(t *testing.T) {
	for _, s := range []string{"a modern deck for an event", "a tempo deck for friday"} {
		if competitiveRequest(s) {
			t.Errorf("%q: an occasion was read as a power level", s)
		}
	}
	for _, s := range []string{
		"the strongest modern deck", "money is no object",
		"as strong as possible", "a competitive modern deck",
	} {
		if !competitiveRequest(s) {
			t.Errorf("%q: a competitive request was missed", s)
		}
	}
}

// TestOccasionIsNotAPowerStep is D-219. The classifier can answer the
// tournament step for "A Modern deck for an event", and the slot must
// not fill from it.
func TestOccasionIsNotAPowerStep(t *testing.T) {
	for _, s := range []string{
		"a modern deck for an event", "a deck for a team event",
		"something for my local game store",
	} {
		if !occasionOnly(s) {
			t.Errorf("%q: an occasion was read as a power step", s)
		}
	}
	// A message that names a step, or asks for a strong deck, keeps it.
	for _, s := range []string{
		"a modern deck for an fnm event", "poison in a tournament",
		"the strongest deck for an event", "a casual deck for game night",
	} {
		if occasionOnly(s) {
			t.Errorf("%q: a named step was dropped with the occasion", s)
		}
	}
}

// TestUserWordsDropsTheQuotedQuestion is D-280. The format question
// names three formats, and the word rules must not read the echo as a
// request for three decks.
func TestUserWordsDropsTheQuotedQuestion(t *testing.T) {
	msg := QuotedQuestionPrefix + "Which format would you like: Commander, Standard, or Modern?\n" + AnswerPrefix + "Modern"
	if got := UserWords(msg); got != "Modern" {
		t.Errorf("UserWords = %q, want Modern", got)
	}
	// The raw echo still reads as three decks. That is why Turn hands the
	// word rules UserWords and never the raw message.
	if !oneDeckRequest(msg) {
		t.Error("the raw echo no longer reads as two decks, so this guard is dead")
	}
	if oneDeckRequest(UserWords(msg)) {
		t.Error("Modern alone read as two decks")
	}
	if got := UserWords("plain words\nmore words"); got != "plain words\nmore words" {
		t.Errorf("plain message changed: %q", got)
	}
}

// TestAddMessageKeepsOnlyTheUserWords is D-292. The power question echoes
// "near a precon" in its options, and the precon rule must not fire on
// the agent's own words.
func TestAddMessageKeepsOnlyTheUserWords(t *testing.T) {
	st := Restore("s", &mtgv1.Slots{}, Snapshot{Version: SnapshotVersion})
	st.AddMessage(QuotedQuestionPrefix + "Which power bracket? 2 core, near a precon, 3 upgraded\n" + AnswerPrefix + "3 upgraded")
	if preconRequest(st.Ctx.Words) {
		t.Errorf("the quoted question reached the words: %q", st.Ctx.Words)
	}
	if !strings.Contains(st.Ctx.Words, "upgraded") {
		t.Errorf("the user's answer left the words: %q", st.Ctx.Words)
	}
	if len(st.Prior()) != 1 || !strings.HasPrefix(st.Prior()[0], QuotedQuestionPrefix) {
		t.Errorf("the classifier's prior messages lost the question: %v", st.Prior())
	}
}

// TestHistoricNeedsAFormatContext is D-146. "Historic" is a card class
// as well as an Arena format, so the word alone declines nothing.
func TestHistoricNeedsAFormatContext(t *testing.T) {
	cases := map[string]bool{
		"a historic deck":                    true,
		"the historic format please":         true,
		"historic on arena":                  true,
		"i play historic on arena, mono red": true,
		"a historic matters deck":            false,
		"a deck full of historic spells":     false,
		"historic permanents, lots of them":  false,
		"historic creatures in commander":    false,
		"a historic moment for my playgroup": false,
		"not historic, standard":             false,
		"a modern deck with a historic feel": false,
		"build me something with historic":   false,
	}
	for text, want := range cases {
		name, _, ok := unsupportedFormat(text)
		if ok != want || (ok && name != "Historic") {
			t.Errorf("unsupportedFormat(%q) = %q/%v, want Historic=%v", text, name, ok, want)
		}
	}
}

// TestOrdinalNeedsTheOrAnOfferWord is D-121 with the guard. A bare
// ordinal orders the sentence, and the pick needs "the" before it or an
// offer word after it.
func TestOrdinalNeedsTheOrAnOfferWord(t *testing.T) {
	cases := map[string]int{
		"first, make it budget":            -1,
		"first make it cheap, then strong": -1,
		"second, no proxies":               -1,
		"the first":                        0,
		"i'll take the second":             1,
		"second one please":                1,
		"the third commander":              2,
		"third option":                     2,
		"the 1st of those":                 0,
	}
	for msg, want := range cases {
		got, ok := offeredPick(msg)
		if want < 0 {
			if ok {
				t.Errorf("offeredPick(%q) = %d, want no pick", msg, got)
			}
			continue
		}
		if !ok || got != want {
			t.Errorf("offeredPick(%q) = %d (%v), want %d", msg, got, ok, want)
		}
	}
}
