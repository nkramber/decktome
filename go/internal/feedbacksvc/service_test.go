package feedbacksvc

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/agentsvc"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/feedback"
	"github.com/nkramber/decktome/go/internal/sessions"
)

type fakeStore struct {
	items []feedback.Item
	uids  []string
	err   error
}

func (f *fakeStore) Add(_ context.Context, uid string, item feedback.Item) (string, error) {
	if f.err != nil {
		return "", f.err
	}
	f.items = append(f.items, item)
	f.uids = append(f.uids, uid)
	return "fb1", nil
}

type fakeSessions map[string]map[string]*mtgv1.Session

func (f fakeSessions) Get(_ context.Context, uid, id string) (*mtgv1.Session, error) {
	if s, ok := f[uid][id]; ok {
		return s, nil
	}
	return nil, sessions.ErrNotFound
}

type fakeDecks map[string]map[string]*mtgv1.Deck

func (f fakeDecks) Get(_ context.Context, uid, id string) (*mtgv1.Deck, error) {
	if d, ok := f[uid][id]; ok {
		return d, nil
	}
	return nil, decks.ErrNotFound
}

const (
	kindQuestion = mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION
	kindSummary  = mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY
	kindCard     = mtgv1.FeedbackKind_FEEDBACK_KIND_CARD
	kindDeck     = mtgv1.FeedbackKind_FEEDBACK_KIND_DECK
	up           = mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_UP
	down         = mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN
)

var at = time.Date(2026, 9, 6, 16, 0, 0, 0, time.UTC)

// fixture: user u1 owns session s1 with question q1 and deck d1 with a
// summary, Sol Ring in the main deck, and Karlov in the command zone.
// User u2 owns nothing.
func fixture() (*fakeStore, *Server) {
	store := &fakeStore{}
	sess := fakeSessions{"u1": {"s1": &mtgv1.Session{Id: "s1", Turns: []*mtgv1.Turn{{Questions: []*mtgv1.Question{{Id: "q1", Text: "Which format?"}}}}}}}
	deckList := fakeDecks{"u1": {
		"d1": &mtgv1.Deck{Id: "d1", Summary: "A lifegain deck.", CommanderOracleIds: []string{"o-karlov"}, Cards: []*mtgv1.DeckCard{{OracleId: "o-sol"}}, Upgrades: []*mtgv1.DeckCard{{OracleId: "o-up"}}},
		"d2": &mtgv1.Deck{Id: "d2"},
	}}
	s := New(store, sess, deckList, auth.UserID, WithClock(func() time.Time { return at }))
	return store, s
}

func submit(s *Server, uid string, fb *mtgv1.Feedback) (*mtgv1.SubmitFeedbackResponse, error) {
	ctx := context.Background()
	if uid != "" {
		ctx = auth.WithUserID(ctx, uid)
	}
	resp, err := s.SubmitFeedback(ctx, connect.NewRequest(&mtgv1.SubmitFeedbackRequest{Feedback: fb}))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

func codeOf(t *testing.T, err error) connect.Code {
	t.Helper()
	var ce *connect.Error
	if !errors.As(err, &ce) {
		t.Fatalf("want connect error, got %v", err)
	}
	return ce.Code()
}

// TestSubmitStoresTheVerdict: the four kinds write one item each under
// the caller, with the short names, the reasons once each, the prompt
// versions, and the time.
func TestSubmitStoresTheVerdict(t *testing.T) {
	store, s := fixture()
	cases := []struct {
		name string
		fb   *mtgv1.Feedback
		want feedback.Item
	}{
		{"a question up", &mtgv1.Feedback{Kind: kindQuestion, Verdict: up, SessionId: "s1", QuestionId: "q1"}, feedback.Item{Kind: "question", Verdict: "up", SessionID: "s1", QuestionID: "q1"}},
		{"a summary down", &mtgv1.Feedback{Kind: kindSummary, Verdict: down, DeckId: "d1", Reasons: []string{"false_claim"}}, feedback.Item{Kind: "summary", Verdict: "down", DeckID: "d1", Reasons: []string{"false_claim"}}},
		{"a card down with a repeated reason and a text", &mtgv1.Feedback{Kind: kindCard, Verdict: down, DeckId: "d1", OracleId: "o-sol", Reasons: []string{"off_theme", " off_theme ", "wrong_power"}, Text: "  Not in a lifegain deck. "}, feedback.Item{Kind: "card", Verdict: "down", DeckID: "d1", OracleID: "o-sol", Reasons: []string{"off_theme", "wrong_power"}, Text: "Not in a lifegain deck."}},
		{"the commander", &mtgv1.Feedback{Kind: kindCard, Verdict: up, DeckId: "d1", OracleId: "o-karlov"}, feedback.Item{Kind: "card", Verdict: "up", DeckID: "d1", OracleID: "o-karlov"}},
		{"an upgrade", &mtgv1.Feedback{Kind: kindCard, Verdict: up, DeckId: "d1", OracleId: "o-up"}, feedback.Item{Kind: "card", Verdict: "up", DeckID: "d1", OracleID: "o-up"}},
		{"a deck down with a text alone", &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1", Text: "Too slow."}, feedback.Item{Kind: "deck", Verdict: "down", DeckID: "d1", Text: "Too slow."}},
	}
	for i, tc := range cases {
		resp, err := submit(s, "u1", tc.fb)
		if err != nil {
			t.Fatalf("%s: %v", tc.name, err)
		}
		if resp.GetFeedbackId() != "fb1" {
			t.Errorf("%s: id = %q", tc.name, resp.GetFeedbackId())
		}
		got := store.items[i]
		if store.uids[i] != "u1" {
			t.Errorf("%s: stored under %q", tc.name, store.uids[i])
		}
		if got.Kind != tc.want.Kind || got.Verdict != tc.want.Verdict || got.SessionID != tc.want.SessionID || got.QuestionID != tc.want.QuestionID || got.DeckID != tc.want.DeckID || got.OracleID != tc.want.OracleID || got.Text != tc.want.Text {
			t.Errorf("%s: stored %+v, want %+v", tc.name, got, tc.want)
		}
		if strings.Join(got.Reasons, ",") != strings.Join(tc.want.Reasons, ",") {
			t.Errorf("%s: reasons = %v, want %v", tc.name, got.Reasons, tc.want.Reasons)
		}
		if got.Prompts["questions"] == 0 || got.Prompts["generate"] == 0 {
			t.Errorf("%s: prompts = %v, want both versions", tc.name, got.Prompts)
		}
		if !got.CreatedAt.Equal(at) {
			t.Errorf("%s: created_at = %v", tc.name, got.CreatedAt)
		}
	}
}

// TestSubmitRefuses: every bad request gets its code, and nothing is
// stored. Another user's object is PermissionDenied (design note 2.5).
func TestSubmitRefuses(t *testing.T) {
	long := strings.Repeat("x", MaxTextBytes+1)
	cases := []struct {
		name string
		uid  string
		fb   *mtgv1.Feedback
		code connect.Code
	}{
		{"no user", "", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1"}, connect.CodeUnauthenticated},
		{"no feedback", "u1", nil, connect.CodeInvalidArgument},
		{"no kind", "u1", &mtgv1.Feedback{Verdict: up, DeckId: "d1"}, connect.CodeInvalidArgument},
		{"no verdict", "u1", &mtgv1.Feedback{Kind: kindDeck, DeckId: "d1"}, connect.CodeInvalidArgument},
		{"a long text", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1", Text: long}, connect.CodeInvalidArgument},
		{"a reason of another kind", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1", Reasons: []string{"already_answered"}}, connect.CodeInvalidArgument},
		{"an unknown reason", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1", Reasons: []string{"meh"}}, connect.CodeInvalidArgument},
		{"a thumbs up with a reason", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1", Reasons: []string{"off_spec"}}, connect.CodeInvalidArgument},
		{"a thumbs up with a text", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1", Text: "nice"}, connect.CodeInvalidArgument},
		{"a thumbs down with nothing", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: down, DeckId: "d1", Text: "  "}, connect.CodeInvalidArgument},
		{"a question with no session", "u1", &mtgv1.Feedback{Kind: kindQuestion, Verdict: up, QuestionId: "q1"}, connect.CodeInvalidArgument},
		{"a question with no question id", "u1", &mtgv1.Feedback{Kind: kindQuestion, Verdict: up, SessionId: "s1"}, connect.CodeInvalidArgument},
		{"a question with a deck id", "u1", &mtgv1.Feedback{Kind: kindQuestion, Verdict: up, SessionId: "s1", QuestionId: "q1", DeckId: "d1"}, connect.CodeInvalidArgument},
		{"a deck with a session id", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1", SessionId: "s1"}, connect.CodeInvalidArgument},
		{"a deck with an oracle id", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1", OracleId: "o-sol"}, connect.CodeInvalidArgument},
		{"a card with no oracle id", "u1", &mtgv1.Feedback{Kind: kindCard, Verdict: up, DeckId: "d1"}, connect.CodeInvalidArgument},
		{"a bad deck id", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "a/b"}, connect.CodeInvalidArgument},
		{"a bad session id", "u1", &mtgv1.Feedback{Kind: kindQuestion, Verdict: up, SessionId: "..", QuestionId: "q1"}, connect.CodeInvalidArgument},
		{"a question the session never asked", "u1", &mtgv1.Feedback{Kind: kindQuestion, Verdict: up, SessionId: "s1", QuestionId: "q9"}, connect.CodeInvalidArgument},
		{"a card the deck does not hold", "u1", &mtgv1.Feedback{Kind: kindCard, Verdict: up, DeckId: "d1", OracleId: "o-none"}, connect.CodeInvalidArgument},
		{"a deck with no summary", "u1", &mtgv1.Feedback{Kind: kindSummary, Verdict: up, DeckId: "d2"}, connect.CodeInvalidArgument},
		{"another user's session", "u2", &mtgv1.Feedback{Kind: kindQuestion, Verdict: up, SessionId: "s1", QuestionId: "q1"}, connect.CodePermissionDenied},
		{"another user's deck", "u2", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1"}, connect.CodePermissionDenied},
		{"another user's card", "u2", &mtgv1.Feedback{Kind: kindCard, Verdict: down, DeckId: "d1", OracleId: "o-sol", Reasons: []string{"illegal"}}, connect.CodePermissionDenied},
		{"an unknown deck", "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d9"}, connect.CodePermissionDenied},
	}
	for _, tc := range cases {
		store, s := fixture()
		_, err := submit(s, tc.uid, tc.fb)
		if err == nil {
			t.Errorf("%s: want %v, got success", tc.name, tc.code)
			continue
		}
		if got := codeOf(t, err); got != tc.code {
			t.Errorf("%s: code = %v, want %v (%v)", tc.name, got, tc.code, err)
		}
		if len(store.items) != 0 {
			t.Errorf("%s: stored %d items, want none", tc.name, len(store.items))
		}
	}
}

// TestSubmitReportsAStoreFailure: a write that fails is Internal.
func TestSubmitReportsAStoreFailure(t *testing.T) {
	store, s := fixture()
	store.err = errors.New("down")
	_, err := submit(s, "u1", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1"})
	if codeOf(t, err) != connect.CodeInternal {
		t.Errorf("code = %v, want Internal", codeOf(t, err))
	}
}

// TestTextCapIsTheMessageCap holds the two caps equal, so the "Other"
// box and the chat refuse the same length.
func TestTextCapIsTheMessageCap(t *testing.T) {
	if MaxTextBytes != agentsvc.MaxMessageBytes {
		t.Errorf("MaxTextBytes = %d, agentsvc.MaxMessageBytes = %d", MaxTextBytes, agentsvc.MaxMessageBytes)
	}
}

// TestTheVerdictKeepsTheObjectItNames locks D-635. A reader deletes a
// session or a deck, and the verdict outlives it: both real items of
// 2026-09-09 named a target the store no longer held. So the server
// stores the object beside the verdict, and it never asks the client
// for it.
func TestTheVerdictKeepsTheObjectItNames(t *testing.T) {
	for _, tc := range []struct {
		name    string
		fb      *mtgv1.Feedback
		session bool
		deck    bool
	}{
		{
			name:    "a question keeps its session",
			fb:      &mtgv1.Feedback{Kind: mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION, Verdict: mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN, SessionId: "s1", QuestionId: "q1", Reasons: []string{"already_answered"}},
			session: true,
		},
		{
			name:    "a chat keeps its session",
			fb:      &mtgv1.Feedback{Kind: mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT, Verdict: mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN, SessionId: "s1", Reasons: []string{"stuck"}},
			session: true,
		},
		{
			name: "a card keeps its deck",
			fb:   &mtgv1.Feedback{Kind: mtgv1.FeedbackKind_FEEDBACK_KIND_CARD, Verdict: mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN, DeckId: "d1", OracleId: "o-sol", Reasons: []string{"off_theme"}},
			deck: true,
		},
		{
			name: "a summary keeps its deck",
			fb:   &mtgv1.Feedback{Kind: mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY, Verdict: mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN, DeckId: "d1", Reasons: []string{"false_claim"}},
			deck: true,
		},
		{
			name: "a deck keeps its deck",
			fb:   &mtgv1.Feedback{Kind: mtgv1.FeedbackKind_FEEDBACK_KIND_DECK, Verdict: mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_UP, DeckId: "d1"},
			deck: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			store, s := fixture()
			if _, err := submit(s, "u1", tc.fb); err != nil {
				t.Fatalf("submit: %v", err)
			}
			if len(store.items) != 1 {
				t.Fatalf("stored %d items, want 1", len(store.items))
			}
			got := store.items[0]
			if tc.session && got.Session.GetId() != "s1" {
				t.Errorf("the verdict kept no session, so a harvest after a delete reads nothing")
			}
			if !tc.session && got.Session != nil {
				t.Errorf("the verdict kept a session it does not name")
			}
			if tc.deck && got.Deck.GetId() != "d1" {
				t.Errorf("the verdict kept no deck, so a harvest after a delete reads nothing")
			}
			if !tc.deck && got.Deck != nil {
				t.Errorf("the verdict kept a deck it does not name")
			}
		})
	}
}

// TestTheSnapshotIsNotTheClientsWord holds the rule of D-596 for the
// snapshot: the server reads the object it already loaded to check the
// owner, so a client that names another deck changes nothing.
func TestTheSnapshotIsNotTheClientsWord(t *testing.T) {
	store, s := fixture()
	fb := &mtgv1.Feedback{
		Kind:    mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY,
		Verdict: mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN,
		DeckId:  "d1",
		Reasons: []string{"false_claim"},
		Text:    "the summary claims lifegain",
	}
	if _, err := submit(s, "u1", fb); err != nil {
		t.Fatalf("submit: %v", err)
	}
	got := store.items[0].Deck
	if got.GetSummary() != "A lifegain deck." {
		t.Errorf("summary = %q, want the stored one", got.GetSummary())
	}
	if len(got.GetCards()) != 1 || got.GetCards()[0].GetOracleId() != "o-sol" {
		t.Errorf("the snapshot does not hold the deck list, so no card class can write its case")
	}
}
