package feedbacksvc

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/notify"
)

type fakeNotifier struct{ notices []notify.Notice }

func (f *fakeNotifier) Notify(n notify.Notice) { f.notices = append(f.notices, n) }

// pingFixture is the fixture of service_test.go with a notifier and a
// clock the test moves. User u2 owns deck d1 as well, so two readers can
// judge the same deck.
func pingFixture(clock *time.Time) (*fakeStore, *fakeNotifier, *Server) {
	store := &fakeStore{}
	sess := fakeSessions{"u1": {"s1": &mtgv1.Session{Id: "s1", Turns: []*mtgv1.Turn{{Questions: []*mtgv1.Question{{Id: "q1", Text: "Which format?"}}}}}}}
	deck := &mtgv1.Deck{Id: "d1", Summary: "A deck.", Cards: []*mtgv1.DeckCard{{OracleId: "o-sol"}}}
	deckList := fakeDecks{"u1": {"d1": deck}, "u2": {"d1": deck}}
	n := &fakeNotifier{}
	s := New(store, sess, deckList, auth.UserID, WithClock(func() time.Time { return *clock }), WithNotifier(n))
	return store, n, s
}

func submitAs(s *Server, uid, email string, fb *mtgv1.Feedback) error {
	ctx := auth.WithEmail(auth.WithUserID(context.Background(), uid), email)
	_, err := s.SubmitFeedback(ctx, connect.NewRequest(&mtgv1.SubmitFeedbackRequest{Feedback: fb}))
	return err
}

// TestEveryVerdictPingsTheOwner locks D-897: a thumbs up, a thumbs down,
// and an import report each send one notice.
func TestEveryVerdictPingsTheOwner(t *testing.T) {
	for _, tc := range []struct {
		name  string
		fb    *mtgv1.Feedback
		title string
		lines []string
	}{
		{
			name:  "a thumbs up",
			fb:    &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1"},
			title: "decktome: thumbs up, deck",
			lines: []string{"deck: d1", "verdict: fb1", "time: 2026-09-06 16:00 UTC"},
		},
		{
			name:  "a thumbs down on a card",
			fb:    &mtgv1.Feedback{Kind: kindCard, Verdict: down, DeckId: "d1", OracleId: "o-sol", Reasons: []string{"off_theme", "wrong_power"}},
			title: "decktome: thumbs down, card",
			lines: []string{"reasons: off_theme, wrong_power", "deck: d1", "card: o-sol"},
		},
		{
			name:  "a thumbs down on a question",
			fb:    &mtgv1.Feedback{Kind: kindQuestion, Verdict: down, SessionId: "s1", QuestionId: "q1", Reasons: []string{"unclear"}},
			title: "decktome: thumbs down, question",
			lines: []string{"session: s1", "question: q1"},
		},
		{
			name:  "an import report",
			fb:    &mtgv1.Feedback{Kind: kindImport, Verdict: down, ImportPage: pageDeck, ImportContent: []byte("hello\nworld\n"), Text: "Some app"},
			title: "decktome: thumbs down, import",
			lines: []string{"reasons: parse_fault", "message: Some app", "page: IMPORT_PAGE_DECK, 1 of 1 rows do not parse", "error: no line of the list reads as a card"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			clock := at
			_, n, s := pingFixture(&clock)
			if err := submitAs(s, "u1", "reader@example.com", tc.fb); err != nil {
				t.Fatalf("submit: %v", err)
			}
			if len(n.notices) != 1 {
				t.Fatalf("sent %d notices, want 1", len(n.notices))
			}
			got := n.notices[0]
			if got.Title != tc.title {
				t.Errorf("title = %q, want %q", got.Title, tc.title)
			}
			for _, want := range append(tc.lines, "user: reader@example.com") {
				if !strings.Contains(got.Message, want) {
					t.Errorf("message lacks %q:\n%s", want, got.Message)
				}
			}
		})
	}
}

// TestThePingHoldsFifteenWords locks D-894: the notice holds the first
// fifteen words of the reader, and marks the cut.
func TestThePingHoldsFifteenWords(t *testing.T) {
	clock := at
	_, n, s := pingFixture(&clock)
	text := "one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen sixteen seventeen"
	if err := submitAs(s, "u1", "", &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1", Text: text}); err != nil {
		t.Fatalf("submit: %v", err)
	}
	msg := n.notices[0].Message
	if !strings.Contains(msg, "message: one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen ...") {
		t.Errorf("the words are not the first fifteen:\n%s", msg)
	}
	if strings.Contains(msg, "sixteen") {
		t.Errorf("the notice holds a sixteenth word:\n%s", msg)
	}
	// No email on the request writes no user line.
	if strings.Contains(msg, "user:") {
		t.Errorf("an empty email wrote a user line:\n%s", msg)
	}
}

// TestOneUserPingsOnceAMinute locks the debounce of D-896: one notice for
// each user in a minute, and another user keeps a window of their own.
func TestOneUserPingsOnceAMinute(t *testing.T) {
	clock := at
	store, n, s := pingFixture(&clock)
	fb := &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1", Reasons: []string{"off_spec"}}
	for i := 0; i < 10; i++ {
		if err := submitAs(s, "u1", "", fb); err != nil {
			t.Fatalf("submit %d: %v", i, err)
		}
	}
	if len(n.notices) != 1 {
		t.Fatalf("ten clicks sent %d notices, want 1", len(n.notices))
	}
	if len(store.items) != 10 {
		t.Fatalf("stored %d verdicts, want 10: the debounce drops the notice alone", len(store.items))
	}
	if err := submitAs(s, "u2", "", fb); err != nil {
		t.Fatalf("submit u2: %v", err)
	}
	if len(n.notices) != 2 {
		t.Fatalf("a second user sent %d notices in all, want 2", len(n.notices))
	}
	clock = at.Add(PingWindow)
	if err := submitAs(s, "u1", "", fb); err != nil {
		t.Fatalf("submit after the window: %v", err)
	}
	if len(n.notices) != 3 {
		t.Fatalf("a click after the window sent %d notices in all, want 3", len(n.notices))
	}
}

// TestARefusedVerdictPingsNothing: a verdict that the server refuses or
// that the store fails to write sends no notice.
func TestARefusedVerdictPingsNothing(t *testing.T) {
	clock := at
	store, n, s := pingFixture(&clock)
	if err := submitAs(s, "u1", "", &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1"}); err == nil {
		t.Fatal("a thumbs down with no reason passed")
	}
	store.err = errors.New("firestore down")
	if err := submitAs(s, "u1", "", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1"}); err == nil {
		t.Fatal("a store failure passed")
	}
	if len(n.notices) != 0 {
		t.Fatalf("sent %d notices, want 0", len(n.notices))
	}
}
