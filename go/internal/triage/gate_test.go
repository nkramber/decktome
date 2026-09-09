package triage

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/harvest"
)

// The PR-28b gate: a dry triage over a fixture of ten items names the
// class of each. The fixture is the owner's own past complaints, in the
// shapes the dialog of PR-27 sends (D-557, D-643).
//
// The dry lane calls no model, so this test costs nothing and gives the
// same answer on every run. A verdict the keys can not place is expected
// to want the judge, and the fixture names that too.

// want is the route the fixture expects for one verdict.
type want struct {
	id       string
	class    string
	need     Need
	artifact Artifact
	// decision is the rule the reader argues with, when the class meets
	// one.
	decision string
}

var fixtureWants = []want{
	{"f01", "Q1", NeedNothing, AConversation, ""},
	{"f02", "Q3", NeedNothing, AConversation, ""},
	{"f03", "S1", NeedNothing, ASummaryCase, ""},
	{"f04", "C1", NeedNothing, ADeckPrompt, ""},
	{"f05", "C2", NeedNothing, ADefect, ""},
	// The reader argues with owned-first and wrote words, so the judge
	// says whether it is a case or a question for the owner.
	{"f06", "C3", NeedJudge, "", "D-37"},
	// The same class of dispute with no words takes no judge call: there
	// is nothing to read but the checked reason.
	{"f07", "D3", NeedNothing, ABracketPrompt, "D-459"},
	// Two reasons of two classes: the judge picks one.
	{"f08", "", NeedJudge, "", ""},
	{"f09", "X1", NeedNothing, ADefect, ""},
	{"f10", "", NeedNothing, ADeckPrompt, ""},
}

// TestTheDryTriageNamesTheClassOfEveryFixtureItem is the gate.
func TestTheDryTriageNamesTheClassOfEveryFixtureItem(t *testing.T) {
	recs, err := ReadHarvest("testdata/fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != len(fixtureWants) {
		t.Fatalf("the fixture holds %d verdicts and the gate names %d", len(recs), len(fixtureWants))
	}
	results := Run(context.Background(), recs, nil, testNamer, testNextID)
	for i, w := range fixtureWants {
		r := results[i]
		if r.Route.Record.ID != w.id {
			t.Fatalf("result %d is verdict %s, want %s", i, r.Route.Record.ID, w.id)
		}
		if r.Err != nil {
			t.Errorf("%s: %v", w.id, r.Err)
			continue
		}
		if got := r.Route.Class.ID; got != w.class {
			t.Errorf("%s: class = %q, want %q. Why: %s", w.id, got, w.class, r.Route.Why)
		}
		if got := r.Route.Need; got != w.need {
			t.Errorf("%s: need = %q, want %q. Why: %s", w.id, got, w.need, r.Route.Why)
		}
		if got := r.Route.Decision; got != w.decision {
			t.Errorf("%s: decision = %q, want %q", w.id, got, w.decision)
		}
		if got := r.Case.Artifact; got != w.artifact {
			t.Errorf("%s: artifact = %q, want %q", w.id, got, w.artifact)
		}
		// Every route says how it got there, so a person can argue with
		// it.
		if strings.TrimSpace(r.Route.Why) == "" {
			t.Errorf("%s: the route says nothing about why", w.id)
		}
	}
}

// TestTheDryTriageCallsNoModel pins the free lane. A nil judge is the
// whole guard, and a route that wants one keeps its need.
func TestTheDryTriageCallsNoModel(t *testing.T) {
	recs, err := ReadHarvest("testdata/fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	results := Run(context.Background(), recs, nil, testNamer, testNextID)
	judged := 0
	for _, r := range results {
		if r.Route.Need == NeedJudge {
			judged++
			if len(r.Case.Body) > 0 {
				t.Errorf("%s wrote a case with no judge answer", r.Route.Record.ID)
			}
		}
	}
	if judged != 2 {
		t.Errorf("%d verdicts want the judge, want 2", judged)
	}
	var b bytes.Buffer
	if pass := Report(&b, results, time.Now(), true, "nothing"); !pass {
		t.Errorf("the dry verdict is FAIL:\n%s", b.String())
	}
	if !strings.Contains(b.String(), "called no model") {
		t.Error("the document does not say the run was dry")
	}
}

// TestTheAskedAgainCaseForbidsTheRow reads the headline artifact. A
// reader who says "you already had this answer" becomes a conversation
// with the row they saw twice, and the question gate fails when it
// fires.
func TestTheAskedAgainCaseForbidsTheRow(t *testing.T) {
	rec := readFixture(t, "f01")
	r := RouteOf(rec)
	c, err := CaseOf(r, 109, testNamer)
	if err != nil {
		t.Fatal(err)
	}
	var body conversationCase
	if err := json.Unmarshal(c.Body, &body); err != nil {
		t.Fatal(err)
	}
	if len(body.MustNotAsk) != 1 || body.MustNotAsk[0] != "power_commander" {
		t.Errorf("must_not_ask = %v, want [power_commander]", body.MustNotAsk)
	}
	if len(body.Messages) != 3 {
		t.Errorf("messages = %v, want the reader's three", body.Messages)
	}
	// The expectation reads the slots of the session in the gate's own
	// words, so the case can not name a value the gate does not read.
	for key, val := range map[string]string{
		"format": "commander", "colors": "WB", "power": "bracket 3",
		"pool_rule": "owned_first", "theme": "lifegain", "budget": "50 to buy",
		"commander": "Karlov of the Ghost Council",
	} {
		if body.Expect[key] != val {
			t.Errorf("expect[%s] = %q, want %q", key, body.Expect[key], val)
		}
	}
	if body.ID != 109 {
		t.Errorf("id = %d, want the number the file gave", body.ID)
	}
	if !strings.Contains(body.Note, "already told you the bracket") {
		t.Errorf("the note drops the reader's words: %q", body.Note)
	}
}

// TestTheOffThemeCaseForbidsTheCard reads the deck-prompt artifact. The
// card the reader called off theme must not come back, and the prompt
// carries the theme and the pool rule of the session that built the
// deck (D-643).
func TestTheOffThemeCaseForbidsTheCard(t *testing.T) {
	rec := readFixture(t, "f04")
	c, err := CaseOf(RouteOf(rec), 26, testNamer)
	if err != nil {
		t.Fatal(err)
	}
	var body deckPromptCase
	if err := json.Unmarshal(c.Body, &body); err != nil {
		t.Fatal(err)
	}
	if len(body.MustNotInclude) != 1 || body.MustNotInclude[0] != "Blood Artist" {
		t.Errorf("must_not_include = %v, want [Blood Artist]", body.MustNotInclude)
	}
	if body.Theme != "lifegain" || body.Pool != "owned_first" || body.Format != "commander" {
		t.Errorf("prompt = %+v, want the slots of the session", body)
	}
	if body.Bracket != 3 {
		t.Errorf("bracket = %d, want 3 from the deck", body.Bracket)
	}
	if body.Commander != "Karlov of the Ghost Council" {
		t.Errorf("commander = %q, want the name of the deck's commander", body.Commander)
	}
}

// TestTheUnwantedBuyCaseAsksForAnOwnedDeck reads the ownership
// assertion. The judge routes this verdict, so the test applies an
// answer that keeps the class.
func TestTheUnwantedBuyCaseAsksForAnOwnedDeck(t *testing.T) {
	rec := readFixture(t, "f06")
	r := RouteOf(rec)
	if r.Need != NeedJudge {
		t.Fatalf("need = %q, want the judge", r.Need)
	}
	r = r.Apply(Verdict{Class: "C3", Why: "the app broke its own owned-first rule"})
	if r.Need != NeedNothing || r.Class.ID != "C3" {
		t.Fatalf("after the judge: class %q, need %q", r.Class.ID, r.Need)
	}
	c, err := CaseOf(r, 26, testNamer)
	if err != nil {
		t.Fatal(err)
	}
	var body deckPromptCase
	if err := json.Unmarshal(c.Body, &body); err != nil {
		t.Fatal(err)
	}
	if !body.MustOwnAll || !body.Collection {
		t.Errorf("prompt = %+v, want an owned-whole assertion with a collection", body)
	}
}

// TestTheOwnerQuestionWritesNoCase holds the rule of D-558: a reader who
// argues with a decision raises a question and never a fix.
func TestTheOwnerQuestionWritesNoCase(t *testing.T) {
	rec := readFixture(t, "f06")
	r := RouteOf(rec).Apply(Verdict{OwnerQuestion: true, Decision: "D-37",
		Why: "the reader wants owned-first to buy nothing"})
	if !r.Owner || r.Need != NeedNothing {
		t.Fatalf("route = %+v, want an owner question", r)
	}
	c, err := CaseOf(r, 26, testNamer)
	if err != nil {
		t.Fatal(err)
	}
	if len(c.Body) != 0 {
		t.Errorf("an owner question wrote a case: %s", c.Body)
	}
	question, whyYou, blocks := OwnerRow(r)
	if !strings.Contains(question, "D-37") || !strings.Contains(question, "still wants me to buy") {
		t.Errorf("the row drops the decision or the reader's words: %q", question)
	}
	if whyYou == "" || blocks == "" {
		t.Error("the row leaves a cell empty")
	}
}

// TestAJudgeThatInventsAClassWritesNothing keeps a bad judge answer out
// of the gate files. A class this build does not hold leaves the route
// where it was, and the route then writes no case.
func TestAJudgeThatInventsAClassWritesNothing(t *testing.T) {
	rec := readFixture(t, "f08")
	r := RouteOf(rec).Apply(Verdict{Class: "Z9", Why: "made up"})
	if r.Class.ID != "" || r.Need != NeedJudge {
		t.Fatalf("route = %+v, want the route unmoved", r)
	}
	if !strings.Contains(r.Why, "Z9") {
		t.Errorf("why = %q, want the invented class named", r.Why)
	}
	if _, err := CaseOf(r, 1, testNamer); err == nil {
		t.Error("a route that still needs the judge wrote a case")
	}
}

// TestEveryReasonKeyNamesOneClass is the completeness bar. A reason key
// with no class would fall through the triage without a word, and D-594
// added a whole kind after the ten classes of the design note.
func TestEveryReasonKeyNamesOneClass(t *testing.T) {
	seen := map[string]string{}
	for _, c := range classes {
		key := c.Kind + "/" + c.Reason
		if other, ok := seen[key]; ok {
			t.Errorf("%s names both %s and %s", key, other, c.ID)
		}
		seen[key] = c.ID
	}
	// The keys of the store, from feedback.reasonKeys. The triage reads
	// the harvest and not the proto, so the list is repeated here as
	// words, and TestTheClassKeysMatchTheStore pins the two together.
	for _, want := range []string{
		"question/already_answered", "question/not_applicable", "question/bad_options", "question/unclear",
		"summary/false_claim", "summary/misses_plan", "summary/too_long_or_vague",
		"card/off_theme", "card/illegal", "card/unwanted_buy", "card/wrong_printing", "card/wrong_power",
		"deck/off_spec", "deck/bad_mana", "deck/too_little_interaction", "deck/wrong_power", "deck/too_many_to_buy",
		"chat/stuck", "chat/ignored_request", "chat/wrong_questions", "chat/no_deck", "chat/error",
	} {
		if seen[want] == "" {
			t.Errorf("reason %s names no class", want)
		}
	}
	if len(classes) != 22 {
		t.Errorf("%d classes, want 22: one per reason key", len(classes))
	}
}

// TestRowOfReadsTheQuestionID pins the one fact the "must not ask"
// artifact rests on: a question id reads "q<n>-<row>" (agent.go).
func TestRowOfReadsTheQuestionID(t *testing.T) {
	for id, want := range map[string]string{
		"q3-power_commander":     "power_commander",
		"q0-format":              "format",
		"q12-format_unsupported": "format_unsupported",
		"":                       "",
		"nodash":                 "",
	} {
		if got := RowOf(id); got != want {
			t.Errorf("RowOf(%q) = %q, want %q", id, got, want)
		}
	}
}

// readFixture reads one verdict of the fixture by its id.
func readFixture(t *testing.T, id string) harvest.Record {
	t.Helper()
	recs, err := ReadHarvest("testdata/fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range recs {
		if r.ID == id {
			return r
		}
	}
	t.Fatalf("the fixture holds no verdict %s", id)
	return harvest.Record{}
}

// testNamer names the cards the fixture uses. The real triage reads the
// card snapshot, and a test needs no snapshot to prove the writers.
func testNamer(oracleID string) (string, bool) {
	names := map[string]string{
		"o-karlov": "Karlov of the Ghost Council",
		"o-c0":     "Ajani's Welcome",
		"o-c1":     "Sanguine Bond",
		"o-c2":     "Blood Artist",
	}
	n, ok := names[oracleID]
	return n, ok
}

// testNextID gives each target file a starting number, the way the real
// files do.
func testNextID(target string) (int, error) {
	switch target {
	case TargetConversations:
		return 109, nil
	case TargetDeckPrompts:
		return 26, nil
	case TargetBracketPrompts:
		return 16, nil
	}
	return 1, nil
}

// TestAVerdictWithNothingToTriageIsCountedApart keeps a verdict that
// holds neither a reason this build knows nor words out of the classed
// count. PR-27 asks for one or the other on every thumbs down, so this
// is rare, and a count that hides it inside the classed ones would read
// as work the triage did.
func TestAVerdictWithNothingToTriageIsCountedApart(t *testing.T) {
	rec := harvest.Record{ID: "x1", Kind: "deck", Verdict: "down",
		Reasons: []string{"a_key_this_build_renamed"}}
	r := RouteOf(rec)
	if r.Need != NeedNothing || r.Class.ID != "" || r.Keep || r.Owner {
		t.Fatalf("route = %+v, want nothing to triage", r)
	}
	results := Run(context.Background(), []harvest.Record{rec}, nil, testNamer, testNextID)
	if len(results) != 1 || len(results[0].Case.Body) != 0 {
		t.Fatalf("a verdict with nothing to triage wrote a case")
	}
	var b bytes.Buffer
	Report(&b, results, time.Now(), true, "nothing")
	doc := b.String()
	if !strings.Contains(doc, "0 reached a class") {
		t.Errorf("it counted as a classed verdict:\n%s", doc)
	}
	if !strings.Contains(doc, "1 verdicts held nothing to triage") {
		t.Errorf("the document does not name it:\n%s", doc)
	}
	if !strings.Contains(doc, "| Nothing to triage | 1 |") {
		t.Errorf("the summary table does not count it:\n%s", doc)
	}
}
