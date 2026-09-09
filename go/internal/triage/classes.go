// Package triage reads the feedback harvest and turns each thumbs down
// into a test case (PR-28b, D-557 to D-559, D-636).
//
// The class comes from the reason keys first, because the dialog of
// PR-27 already asked the reader to name the fault and the keys map onto
// the classes almost one to one (D-643). The judge role reads the item
// only when the keys can not answer: a verdict with free text alone, a
// verdict whose reasons cross two classes, and a verdict that argues
// with a decision the owner made. So a harvest of checked reasons costs
// nothing, and the fixture gate runs free and moves between no two runs.
//
// A class that meets an owner decision writes a row to
// docs/owner-questions.md and no case. A reader who disagrees with a
// bracket rule of D-459 raises a question for the owner, and the loop
// makes no fix (D-558).
package triage

// Artifact names what a class writes.
type Artifact string

const (
	// AConversation is a golden conversation of the question gate. A
	// question class carries the reader's own messages and the row that
	// must not fire again.
	AConversation Artifact = "conversation"
	// ADeckPrompt is a prompt of the deck gate, from the slots of the
	// deck the reader judged, with the assertion of the fault they met.
	ADeckPrompt Artifact = "deck_prompt"
	// ASummaryCase is a deck gate prompt whose bar is the summary judge
	// of F-26. The deck gate judges the summary of every deck it builds,
	// so the summary classes need no harness of their own.
	ASummaryCase Artifact = "summary_case"
	// ABracketPrompt is a prompt of the bracket gate, where the judge
	// names the bracket the deck plays at.
	ABracketPrompt Artifact = "bracket_prompt"
	// ADefect is a report row and no case. The fix is a code fix, and no
	// gate of this repo measures it yet, so a person writes the test.
	ADefect Artifact = "defect"
)

// Class is one triage class: the fault a reader reported, the artifact
// it writes, and where the fix lives.
type Class struct {
	// ID is stable. The report and the fixture name it.
	ID string
	// Name is the fault in the reader's terms.
	Name string
	// Kind and Reason are the feedback kind and the reason key that name
	// this class. Together they are unique.
	Kind   string
	Reason string
	// Artifact is what the class writes.
	Artifact Artifact
	// Where names the code or the prompt the fix lives in.
	Where string
	// Disputes names the decision a reader may be arguing with, or is
	// empty. A verdict of a disputing class that carries free text goes
	// to the judge, and the judge says whether it is a case or a
	// question for the owner (D-558).
	Disputes string
}

// classes is the whole set, in report order. The ten classes of the
// design note keep their ids and their words. Five more carry the chat
// kind of D-594, which arrived after the note, and seven carry the
// reason keys the note left unmapped (D-643). Every reason key of every
// kind names exactly one class, so no verdict falls through.
var classes = []Class{
	{"Q1", "asked again", "question", "already_answered", AConversation, "the catalog trigger, the word rules, the classify prompt", ""},
	{"Q2", "off target", "question", "not_applicable", AConversation, "the catalog when clause", ""},
	{"Q3", "bad options", "question", "bad_options", AConversation, "the resolver and the hints", ""},
	{"Q4", "unclear wording", "question", "unclear", AConversation, "the catalog text and the ask prompt", ""},

	{"S1", "false summary", "summary", "false_claim", ASummaryCase, "the generate prompt", ""},
	{"S2", "misses the plan", "summary", "misses_plan", ASummaryCase, "the generate prompt", ""},
	{"S3", "too long or vague", "summary", "too_long_or_vague", ASummaryCase, "the generate prompt", ""},

	{"C1", "off theme", "card", "off_theme", ADeckPrompt, "the candidates ranking, the tags, the lint", ""},
	{"C2", "illegal", "card", "illegal", ADefect, "the rules engine, a code fix and never a prompt", ""},
	{"C3", "unwanted buy", "card", "unwanted_buy", ADeckPrompt, "the pool and the generate prompt", "D-37"},
	{"C4", "wrong printing", "card", "wrong_printing", ADefect, "the printing choice", ""},
	{"C5", "wrong power for one card", "card", "wrong_power", ABracketPrompt, "the profile and the bracket rules", "D-459"},

	{"D1", "off spec", "deck", "off_spec", ADeckPrompt, "the classify prompt and the build", ""},
	{"D2", "bad mana", "deck", "bad_mana", ADeckPrompt, "the targets and the profile", ""},
	{"D3", "wrong power", "deck", "wrong_power", ABracketPrompt, "the profile and the bracket rules", "D-459"},
	{"D4", "too little interaction", "deck", "too_little_interaction", ADeckPrompt, "the targets and the profile", ""},
	{"D5", "too many to buy", "deck", "too_many_to_buy", ADeckPrompt, "the budget and the pool", "D-37"},

	{"X1", "the chat stopped", "chat", "stuck", ADefect, "the chat turn and the net of D-351", ""},
	{"X2", "it ignored the request", "chat", "ignored_request", AConversation, "the classify prompt and the chat turn", ""},
	{"X3", "the wrong questions", "chat", "wrong_questions", AConversation, "the catalog when clause", ""},
	{"X4", "no deck came", "chat", "no_deck", ADefect, "the build path", ""},
	{"X5", "an error", "chat", "error", ADefect, "the error path", ""},
}

// byKey finds a class by its kind and its reason key.
var byKey = func() map[string]Class {
	m := map[string]Class{}
	for _, c := range classes {
		m[c.Kind+"/"+c.Reason] = c
	}
	return m
}()

// byID finds a class by its id.
var byID = func() map[string]Class {
	m := map[string]Class{}
	for _, c := range classes {
		m[c.ID] = c
	}
	return m
}()

// Classes lists every class, in report order.
func Classes() []Class { return append([]Class(nil), classes...) }

// ClassOf answers the class a kind and a reason key name.
func ClassOf(kind, reason string) (Class, bool) {
	c, ok := byKey[kind+"/"+reason]
	return c, ok
}

// ClassByID answers the class of an id, which is how the judge names one.
func ClassByID(id string) (Class, bool) {
	c, ok := byID[id]
	return c, ok
}
