package agentsvc

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/users"
)

// commanderDeck is a deck the fake generator answers, led by one
// commander.
func commanderDeck(commander string) *mtgv1.Deck {
	return &mtgv1.Deck{
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		CommanderOracleIds: []string{commander},
		Validation:         &mtgv1.ValidationResult{},
		Cards:              []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 30}},
	}
}

// TestARebuildKeepsTheDeckItRevises is D-1118. A rebuild after a change
// of a setting, with the format and the commander of the last deck, is a
// revision of that deck. The record counts a revision, and the web app
// reads the deck it came from. A rebuild with another commander is a new
// deck.
func TestARebuildKeepsTheDeckItRevises(t *testing.T) {
	for _, tc := range []struct {
		name      string
		commander string
		from      string
		counter   users.Counter
	}{
		{name: "the same commander", commander: "o-karlov", from: "deck-1", counter: users.DeckRevisions},
		{name: "another commander", commander: "o-other", from: "", counter: users.DecksCreated},
	} {
		t.Run(tc.name, func(t *testing.T) {
			store := newFakeStore()
			ds := &fakeDeckStore{}
			noter := &fakeNoter{}
			fd := &fakeDecks{res: &generate.Result{Deck: commanderDeck("o-karlov")}}
			steps := append(builtSteps(t), classifyJSON(t, map[string]any{"power": "bracket 2"}))
			opts := append(buildOpts(t, fd), WithDeckStore(ds), WithUsers(noter))
			client, _ := testServerOpts(t, store, opts, steps...)
			first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
			chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
			fd.res = &generate.Result{Deck: commanderDeck(tc.commander)}
			third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Make it bracket 2"})
			if fd.runs != 2 || fd.got.Revision != nil {
				t.Fatalf("runs=%d revision=%v, want a second full build", fd.runs, fd.got.Revision)
			}
			if third.deck == nil || third.deck.GetRevisedFromDeckId() != tc.from {
				t.Fatalf("revised_from = %q, want %q", third.deck.GetRevisedFromDeckId(), tc.from)
			}
			got := noter.decksOnly()
			if len(got) != 2 || got[0] != users.DecksCreated || got[1] != tc.counter {
				t.Errorf("counted %v, want a first deck and one %s", got, tc.counter)
			}
		})
	}
}

// TestASettingChangedInAQuestionTurnStillRebuilds is D-1118. The theme
// changed in a turn that asked the theme row, and the reply changed no
// setting. The reply compared the slots of its own turn alone, so it
// revised the old deck under the old theme. It now compares the slots of
// the last deck, so it builds again, and the status names the word of the
// theme that matches no card (D-1116).
func TestASettingChangedInAQuestionTurnStillRebuilds(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	fd := &fakeDecks{res: &generate.Result{Deck: commanderDeck("o-karlov")}}
	steps := append(builtSteps(t),
		classifyJSON(t, map[string]any{"theme": "lifegain zzzz"}),
		classifyJSON(t, nil))
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Make it lifegain zzzz"})
	if len(third.questions) != 1 || third.deck != nil {
		t.Fatalf("the theme change asked %d questions and built %v, want the theme row alone: %v",
			len(third.questions), third.deck != nil, third.order)
	}
	q := third.questions[0]
	if !strings.Contains(q.GetText(), `"zzzz"`) {
		t.Errorf("the theme row does not name the dead word: %q", q.GetText())
	}
	fd.res = &generate.Result{Deck: commanderDeck("o-karlov")}
	fourth := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started,
		Answers: []*mtgv1.Answer{{QuestionId: q.GetId(), Text: "keep it as it is"}}})
	if fd.runs != 2 || fd.got.Revision != nil {
		t.Fatalf("runs=%d revision=%v, want a full build on the reply: %v", fd.runs, fd.got.Revision, fourth.order)
	}
	statuses := strings.Join(fourth.statuses, "|")
	for _, want := range []string{"a deck setting changed", `No card I know matches "zzzz", so the deck is built without that word`} {
		if !strings.Contains(statuses, want) {
			t.Errorf("statuses lack %q: %v", want, fourth.statuses)
		}
	}
	if fourth.deck == nil || fourth.deck.GetRevisedFromDeckId() != "deck-1" {
		t.Errorf("deck = %v", fourth.deck)
	}
}

// TestThePlanReadsTheLaterMessages is D-1122. The plan sent only the first
// message, so the model never read "too many artifacts" or "heavy on the
// cantrips". It now carries the later messages, the newest ones when the
// chat is long, and what the user wants less of.
func TestThePlanReadsTheLaterMessages(t *testing.T) {
	session := &mtgv1.Session{Turns: []*mtgv1.Turn{
		{UserMessage: "Build a mono-u voltron deck"},
		{UserMessage: ""},
		{UserMessage: "There are a lot of artifacts. Heavy on cantrips, please."},
	}}
	slots := &mtgv1.Slots{Theme: "voltron, cantrips", Avoid: "artifacts"}
	got := plan(session, slots)
	for _, want := range []string{
		"Build a mono-u voltron deck\nLater messages, oldest first:\n- There are a lot of artifacts. Heavy on cantrips, please.",
		"\nTheme: voltron, cantrips",
		"\nLess of: artifacts",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("plan lacks %q:\n%s", want, got)
		}
	}
	if strings.Count(got, "\n- ") != 1 {
		t.Errorf("an empty message reached the plan:\n%s", got)
	}
	for i := 0; i < planLaterTurns+3; i++ {
		session.Turns = append(session.Turns, &mtgv1.Turn{UserMessage: "message " + string(rune('a'+i))})
	}
	got = plan(session, slots)
	if n := strings.Count(got, "\n- "); n != planLaterTurns {
		t.Errorf("the plan holds %d later messages, want %d", n, planLaterTurns)
	}
	if strings.Contains(got, "artifacts. Heavy") || !strings.Contains(got, "message i") {
		t.Errorf("the plan did not keep the newest messages:\n%s", got)
	}
	if !strings.HasPrefix(got, "Build a mono-u voltron deck") {
		t.Errorf("the request left the plan:\n%s", got)
	}
}

// TestAWishForLessIsNoNewSetting is D-1122 beside D-283. A message after
// a build that only fills the avoid slot is a change request against the
// deck, so it revises the deck and does not build it again.
func TestAWishForLessIsNoNewSetting(t *testing.T) {
	before := &mtgv1.Slots{Theme: "voltron"}
	after := &mtgv1.Slots{Theme: "voltron", Avoid: "artifacts"}
	if slotsChanged(before, after) {
		t.Error("the avoid slot alone counts as a new setting")
	}
	if !slotsChanged(before, &mtgv1.Slots{Theme: "cantrips"}) {
		t.Error("a new theme counts as no new setting")
	}
}
