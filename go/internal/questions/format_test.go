package questions

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The format tests: the word rules, the decline rows, and the turn path
// that fills the format. The pure format rules are in words_test.go.

// TestAcceptedNearestFormatFillsTheSlot is D-112. The decline row offers
// "Yes, use the nearest format", and the format must fill on that
// answer.
func TestAcceptedNearestFormatFillsTheSlot(t *testing.T) {
	cases := []struct {
		name, answer string
	}{
		{"the option text", "Yes, use the nearest format."},
		{"a plain yes", "yes"},
		{"an acceptance phrase", "Fine, treat it as that."},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			unknown := classifyOut{Format: "unknown", PoolRule: "unknown"}
			a, sc := testAgent(t,
				classifyStep(t, unknown), fits(t, "format_unsupported", "theme", "colors"), askStep(t),
				classifyStep(t, unknown), fits(t, "power_commander", "budget"), askStep(t))
			st := NewState(false)
			if _, err := a.Turn(context.Background(), st, "I want a Brawl deck for Arena.", nil); err != nil {
				t.Fatalf("turn 1: %v", err)
			}
			if !st.Ctx.Asked["format_unsupported"] {
				t.Fatal("the decline row did not fire on turn 1")
			}
			res, err := a.Turn(context.Background(), st, tc.answer, nil)
			if err != nil {
				t.Fatalf("turn 2: %v", err)
			}
			if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
				t.Fatalf("format = %v, want Commander from the nearest format", st.Ctx.Format)
			}
			if st.Slots.GetSlotStates()["format"] != mtgv1.SlotState_SLOT_STATE_FILLED {
				t.Errorf("format state = %v, want FILLED", st.Slots.GetSlotStates()["format"])
			}
			if len(res.Questions) == 0 {
				t.Error("the plan did not continue after the user accepted the format")
			}
			if question(res.Questions, "format") != nil {
				t.Error("a format question went out after the user accepted the format")
			}
			// The classify call of turn 2 carried the nearest format, so the
			// model could act on it as well.
			var in struct {
				Nearest string `json:"nearest_format"`
			}
			if err := json.Unmarshal([]byte(sc.Calls[3].Input), &in); err != nil {
				t.Fatal(err)
			}
			if in.Nearest != "Commander" {
				t.Errorf("nearest_format in the classify input = %q, want Commander", in.Nearest)
			}
		})
	}
}

// TestUnsupportedFormatAfterAFilledOne is D-112 after D-125. A user who
// names a format this app does not build after a supported one is filled
// must hear the decline: the format reopens through the D-125 path and
// the decline row fires.
func TestUnsupportedFormatAfterAFilledOne(t *testing.T) {
	first := commanderClassify()
	change := classifyOut{Format: "unknown", PoolRule: "unknown"}
	// Turn 2 asks the budget again beside the decline: the format change
	// retired the open budget question, and a retired question may ask
	// again (D-195).
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "power_commander", "budget"), askStep(t),
		classifyStep(t, change), fits(t, "budget"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Commander lifegain deck, white and black.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	res, err := a.Turn(context.Background(), st, "Actually, make it Brawl.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if !st.Ctx.Asked["format_unsupported"] {
		t.Fatalf("the decline row did not fire: %v", ids2(res.Questions))
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED || st.Ctx.Filled["format"] {
		t.Errorf("the format stayed %v after the user named Brawl", st.Ctx.Format)
	}
	if st.UnsupportedFormatName != "Brawl" || st.NearestFormat != "Commander" {
		t.Errorf("decline = %q/%q, want Brawl/Commander", st.UnsupportedFormatName, st.NearestFormat)
	}
	// The open Commander questions retired with the format.
	if st.Slots.GetSlotStates()["commander"] == mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Error("the commander question is still out after the format reopened")
	}
}

// TestWordRuleReadsTheMessageAlone is D-125. A FormatFromWords that read
// the whole conversation would let one "pauper" on turn 1 turn the safety
// net off for the rest of the session.
func TestWordRuleReadsTheMessageAlone(t *testing.T) {
	unknown := classifyOut{Format: "unknown", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, unknown), fits(t, "format_unsupported_open", "theme", "colors"), askStep(t),
		classifyStep(t, unknown), fits(t, "power_sixty", "budget"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Pauper burn deck.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "Modern then.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_MODERN {
		t.Errorf("format = %v, want Modern from the word rule", st.Ctx.Format)
	}
	if st.Ctx.UnsupportedFormat {
		t.Error("the unsupported fact survived a filled format")
	}
}

// TestLintFormatRuleReadsEachMessage is the linter half of D-125. The
// format_already_named rule must not go blind after one unsupported word.
func TestLintFormatRuleReadsEachMessage(t *testing.T) {
	cases := []struct {
		name     string
		messages []string
		want     bool
	}{
		{"a supported format after an unsupported one", []string{"a pauper burn deck", "modern then", "red"}, true},
		{"an unsupported format after a supported one", []string{"a commander deck", "actually brawl", "red"}, false},
		{"only an unsupported format", []string{"a pauper burn deck", "red", "cheap"}, false},
		{"a supported format alone", []string{"a modern burn deck", "red", "cheap"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			q := LintQuestion{Turn: 3, RowID: "format", Slot: "format", Text: "Which format: Commander, Standard, or Modern?"}
			var got bool
			for _, f := range LintConversation(tc.messages, []LintQuestion{q}) {
				if f.Rule == "format_already_named" {
					got = true
				}
			}
			if got != tc.want {
				t.Errorf("format_already_named = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestRetiredThemeAsksAgain is D-195. After a format change retires the
// theme question, the theme row must ask again, or the build runs with
// the slot empty.
func TestRetiredThemeAsksAgain(t *testing.T) {
	first := classifyOut{Format: "commander", PoolRule: "unknown"}
	change := classifyOut{Format: "modern", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "theme", "power_commander", "colors"), askStep(t),
		classifyStep(t, change), fits(t, "theme", "power_sixty", "colors"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Commander deck.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.Asked["theme"] {
		t.Fatal("the theme row did not fire on turn 1")
	}
	res, err := a.Turn(context.Background(), st, "Actually, Modern.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if question(res.Questions, "theme") == nil {
		t.Errorf("the theme row did not ask again after the format change retired it: %v", ids2(res.Questions))
	}
	if question(res.Questions, "colors") == nil {
		t.Errorf("the color row did not ask again after the format change retired it: %v", ids2(res.Questions))
	}
	if q := question(res.Questions, "power"); q == nil || !strings.Contains(q.GetText(), "casual") {
		t.Errorf("the 60-card power row did not fire: %v", ids2(res.Questions))
	}
}

// TestRetireOutstandingClearsTheAskedMark holds the state half of D-195.
// The row ids of a retired key leave Ctx.Asked, and a nil catalog keeps
// them.
func TestRetireOutstandingClearsTheAskedMark(t *testing.T) {
	c := load(t)
	st := NewState(false)
	st.MarkAsked("theme", "theme", "theme")
	st.MarkAsked("power_commander", "power", "power")
	st.RetireOutstanding(c)
	for _, id := range []string{"theme", "power_commander"} {
		if st.Ctx.Asked[id] {
			t.Errorf("row %q kept its asked mark after its key retired", id)
		}
	}
	if st.Outstanding() {
		t.Error("a question is still out after the retire")
	}
	st.MarkAsked("theme", "theme", "theme")
	st.RetireOutstanding(nil)
	if !st.Ctx.Asked["theme"] {
		t.Error("a nil catalog cleared an asked mark it could not map")
	}
}

// TestWordRuleReadsTheFormatTheClassifierMissed is the safety net of
// D-116. "A land destruction Commander deck" names Commander, and the
// format question must not go out.
func TestWordRuleReadsTheFormatTheClassifierMissed(t *testing.T) {
	out := classifyOut{Format: "unknown", Theme: "land destruction", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out),
		fits(t, "commander", "power_commander", "colors"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A land destruction Commander deck.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Fatalf("format = %v, want Commander", st.Ctx.Format)
	}
	if st.Ctx.Asked["format"] {
		t.Error("the agent asked for a format the user had already named")
	}
}

// TestUnsupportedFormatRowNamesTheNearestFormat is D-112.
func TestUnsupportedFormatRowNamesTheNearestFormat(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out),
		fits(t, "format_unsupported", "theme", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I want a Brawl deck for Arena.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	q := question(res.Questions, "format")
	if q == nil {
		t.Fatal("no format question went out")
	}
	for _, want := range []string{"Brawl", "Commander"} {
		if !strings.Contains(q.GetText(), want) {
			t.Errorf("the question %q does not name %q", q.GetText(), want)
		}
	}
}

// TestFormatDoesNotRevertToAReplacedValue is D-125. The user asks for
// Commander, changes to Modern two turns later, and the classifier
// reports Commander again on a message that names no format. The agent
// must not ask for a commander in a 60-card format.
func TestFormatDoesNotRevertToAReplacedValue(t *testing.T) {
	first := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	first.Colors = []string{"W", "B"}
	change := classifyOut{Format: "modern", Theme: "lifegain", PoolRule: "unknown"}
	// The classifier repeats the value the user replaced.
	stale := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "commander", "power_commander"), askStep(t),
		classifyStep(t, change), fits(t, "power_sixty"), askStep(t),
		classifyStep(t, stale), fits(t))
	st := NewState(false)
	for i, msg := range []string{
		"Build me a white-black lifegain Commander deck.",
		"Actually, make it Modern instead. Sixty cards.",
		"FNM level, and 100 dollars.",
	} {
		if _, err := a.Turn(context.Background(), st, msg, nil); err != nil {
			t.Fatalf("turn %d: %v", i+1, err)
		}
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_MODERN {
		t.Errorf("format = %v, want Modern: the user replaced Commander", st.Ctx.Format)
	}
	if st.Ctx.Asked["commander_pick"] {
		t.Error("the agent asked for a commander in a 60-card format")
	}
}

// TestCEDHNamesBracketFive is D-164. "A cEDH deck" names bracket 5, so
// the agent must not ask which power bracket to target.
func TestCEDHNamesBracketFive(t *testing.T) {
	if !cedhRequest("a cEDH deck") {
		t.Error("cEDH was not read as a power level")
	}
	if cedhRequest("a competitive Commander deck") {
		t.Error("competitive Commander names no bracket, and it fired")
	}
	if id, ok := FormatFromWords("a cEDH deck"); !ok || id != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Errorf("cEDH names format %v/%v, want Commander", id, ok)
	}
	out := classifyOut{Format: "commander", Theme: "combo", PoolRule: "any_card", BudgetUSD: 500}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "commander", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "A cEDH deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if q := question(res.Questions, "power"); q != nil {
		t.Errorf("the bracket row asked a cEDH user: %q", q.GetText())
	}
	if got := st.Slots.GetPower().GetBracket(); got != cedhBracket {
		t.Errorf("power bracket = %d, want %d", got, cedhBracket)
	}
}

// TestRetiredQuestionLeavesTheAskedState is H-5. A changed format retires
// the questions that are out, and the retired key stayed in the asked
// state forever: the only ways out of it were an answer and a decline.
// The session never reported ready, and the classifier was offered the
// dead key every turn.
func TestRetiredQuestionLeavesTheAskedState(t *testing.T) {
	first := commanderClassify()
	first.Facts.WantsSuggestion = true
	change := classifyOut{Format: "modern", Theme: "lifegain", PoolRule: "unknown"}
	last := classifyOut{Format: "unknown", PoolRule: "unknown", BudgetUSD: 100}
	last.Power = "casual"
	a, _ := testAgentHints(t, &fakeHints{commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}},
		classifyStep(t, first), fits(t, "commander_pick", "power_commander"), askStep(t),
		classifyStep(t, change), fits(t, "power_sixty", "budget"), askStep(t),
		classifyStep(t, last), fits(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "Commander lifegain, white and black, suggest a commander", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.Asked["commander_pick"] {
		t.Fatal("the pick row did not fire on turn 1")
	}
	if _, err := a.Turn(context.Background(), st, "Actually make it Modern", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	for _, key := range openKeys(st) {
		if key == "commander_pick" {
			t.Error("the retired pick key is still offered to the classifier")
		}
	}
	res, err := a.Turn(context.Background(), st, "Casual, and 100 dollars", nil)
	if err != nil {
		t.Fatalf("turn 3: %v", err)
	}
	if len(res.Questions) != 0 {
		t.Errorf("turn 3 asked %v, want nothing", ids2(res.Questions))
	}
	if !res.Ready {
		t.Errorf("the session is not ready after every slot filled: states %v", st.Slots.GetSlotStates())
	}
	// The retired row may ask again: the user never answered it, so its
	// asked mark leaves with the key (D-195). It stays silent here
	// because Modern has no commander.
	if st.Ctx.Asked["commander_pick"] {
		t.Error("the retired row kept its asked mark, so a return to Commander could never ask it")
	}
}

// TestUnsupportedSubFormatNamesNoFormat is M-9. "Duel Commander" holds
// the word "commander", and the classifier reported Commander for it.
// The format then filled, and the unsupported-format row never fired.
func TestUnsupportedSubFormatNamesNoFormat(t *testing.T) {
	if namesFormat("I play Duel Commander", mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("Duel Commander read as the Commander format")
	}
	if namesFormat("Pauper Commander please", mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("Pauper Commander read as the Commander format")
	}
	if !namesFormat("a Commander deck", mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("a plain Commander request was not read")
	}
	out := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "format_unsupported", "colors", "budget"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "I play Duel Commander. A lifegain deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		t.Errorf("format = %v, want none: the message names a format this app does not build", st.Ctx.Format)
	}
	if st.UnsupportedFormatName != "Duel Commander" {
		t.Errorf("unsupported format = %q, want Duel Commander", st.UnsupportedFormatName)
	}
	if !st.Ctx.Asked["format_unsupported"] {
		t.Errorf("the unsupported-format row did not fire: %v", ids2(res.Questions))
	}
}

// TestDeclineRowWaitsForAnotherFormat holds the D-210 rule. The user
// names Legacy, hears that this app does not build it, and names Legacy
// again. The row said its sentence, so it stays silent.
//
// The same sentence on two turns is a duplicate (D-210).
func TestDeclineRowWaitsForAnotherFormat(t *testing.T) {
	out := classifyOut{Format: "unknown", PoolRule: "unknown"}
	a, _ := testAgent(t,
		classifyStep(t, out), fits(t, "format_unsupported_open", "colors", "budget"), askStep(t),
		classifyStep(t, out), fits(t, "theme"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Legacy deck for an event.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if !st.Ctx.Asked["format_unsupported_open"] {
		t.Fatal("the decline row did not fire on turn 1")
	}
	res, err := a.Turn(context.Background(), st, "Legacy. A Delver of Secrets tempo deck.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	for _, q := range res.Questions {
		if q.GetSlot() == "format" {
			t.Errorf("the decline row asked again on the same format: %q", q.GetText())
		}
	}
}

// The dropped formats (D-155). Pioneer, Legacy, Vintage, and Pauper
// left the app, and nothing may offer or resolve one.

// supportedFormats is the whole list the app builds (D-155). HOUSE is not
// here: no menu offers it, and it is reached only by a user who names no
// format and asks for no ban list.
var supportedFormats = []mtgv1.FormatId{
	mtgv1.FormatId_FORMAT_ID_COMMANDER,
	mtgv1.FormatId_FORMAT_ID_STANDARD,
	mtgv1.FormatId_FORMAT_ID_MODERN,
}

// droppedFormats are the four D-155 removed. Every one must decline.
var droppedFormats = []string{"pioneer", "legacy", "vintage", "pauper"}

// TestNoRowOffersADroppedFormat is the invariant under D-155. A question
// that names a format the app refuses contradicts the row that declines
// it, which is the D-150 defect one level up.
//
// The rows that exist to decline a format are exempt, because naming the
// format is their whole job.
func TestNoRowOffersADroppedFormat(t *testing.T) {
	c := load(t)
	for _, r := range c.Rows {
		if declinesFormat(r.ID) {
			continue
		}
		texts := append([]string{r.Text, r.Fallback}, r.Options...)
		for _, text := range texts {
			low := strings.ToLower(text)
			for _, f := range droppedFormats {
				if strings.Contains(low, f) {
					t.Errorf("row %q names %q, which the app does not build: %q", r.ID, f, text)
				}
			}
		}
	}
}

// TestDroppedFormatsResolveToNothing keeps the classifier and the word
// rules from filling the format slot with a format the app refuses. A
// filled slot would skip the decline row entirely.
func TestDroppedFormatsResolveToNothing(t *testing.T) {
	for _, f := range droppedFormats {
		if got, ok := formatIDs[slotWord(f)]; ok && got != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
			t.Errorf("the classifier resolved %q to %v, want no answer", f, got)
		}
		if id, ok := FormatFromWords("i want a " + f + " deck"); ok {
			t.Errorf("the word rule read %q as %v, want no answer", f, id)
		}
		name, near, ok := unsupportedFormat("i want a " + f + " deck")
		if !ok {
			t.Errorf("%q is not declined by the unsupported-format rule", f)
			continue
		}
		if near != "" {
			t.Errorf("%q offers %q as a substitute, and D-155 names none", f, near)
		}
		if !strings.EqualFold(name, f) {
			t.Errorf("%q declined under the name %q", f, name)
		}
	}
}

// TestSupportedFormatsAllBuild checks that every format the app offers has
// construction rules and a 60-card answer. A format in the menu with no
// rules behind it is the dead-path class of D-76 and D-77.
func TestSupportedFormatsAllBuild(t *testing.T) {
	raw, err := os.ReadFile("../rules/formats.json")
	if err != nil {
		t.Skipf("formats.json not readable: %v", err)
	}
	var cfg struct {
		Formats map[string]map[string]any `json:"formats"`
	}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		t.Fatal(err)
	}
	for _, f := range append(supportedFormats, mtgv1.FormatId_FORMAT_ID_HOUSE) {
		if _, ok := cfg.Formats[f.String()]; !ok {
			t.Errorf("%v is offered and formats.json holds no rules for it", f)
		}
	}
	for name := range cfg.Formats {
		if name == mtgv1.FormatId_FORMAT_ID_HOUSE.String() {
			continue
		}
		var found bool
		for _, f := range supportedFormats {
			if f.String() == name {
				found = true
			}
		}
		if !found {
			t.Errorf("formats.json holds rules for %s, which nothing can select", name)
		}
	}
	// Every 60-card format must answer sixtyCard, or no power row fires.
	for _, f := range []mtgv1.FormatId{
		mtgv1.FormatId_FORMAT_ID_STANDARD,
		mtgv1.FormatId_FORMAT_ID_MODERN,
		mtgv1.FormatId_FORMAT_ID_HOUSE,
	} {
		if !sixtyCard(f) {
			t.Errorf("%v builds 60 cards and sixtyCard says otherwise, so no power row fires", f)
		}
	}
	if sixtyCard(mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("Commander is not a 60-card format")
	}
}

// TestReopenedFormatClearsThePlainRowMark is D-195 for the format. A
// format this app does not build, named after a format was filled,
// reopens the key, and the plain row must be able to ask again.
func TestReopenedFormatClearsThePlainRowMark(t *testing.T) {
	unknown := classifyOut{Format: "unknown", PoolRule: "unknown"}
	p := play(t, false, nil, []turnScript{
		{"A burn deck.", unknown},
		{"Modern.", classifyOut{Format: "modern", PoolRule: "unknown"}},
		{"Actually, make it Brawl.", unknown},
	})
	if !p.askedOn(1, "format") {
		t.Fatalf("turn 1 did not ask the format: %v", p.rows)
	}
	if !p.askedOn(3, "format_unsupported") {
		t.Fatalf("turn 3 did not decline Brawl: %v", p.rows)
	}
	if p.st.Ctx.Asked["format"] {
		t.Error("the plain format row kept its asked mark through the reopen")
	}
	if !p.st.Ctx.Asked["format_unsupported"] {
		t.Error("the decline row lost its asked mark, so it would repeat on the same format")
	}
	// The decline question answered with no format leaves the key open,
	// and the plain row is the one that asks next.
	st := p.st
	st.RetireOutstanding(load(t))
	st.Ctx.UnsupportedFormat, st.Ctx.BadFormatChanged = false, false
	var ids []string
	for _, r := range load(t).Plan(st.Ctx) {
		ids = append(ids, r.ID)
	}
	if len(ids) == 0 || ids[0] != "format" {
		t.Errorf("plan after the reopen = %v, want the plain format row first", ids)
	}
}
