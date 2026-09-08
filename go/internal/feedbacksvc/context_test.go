package feedbacksvc

import (
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/feedback"
)

// withAnswered builds a session whose question q1 the reader answered by
// option index, and q2 the reader declined.
func withAnswered() (*fakeStore, *Server) {
	store := &fakeStore{}
	sess := fakeSessions{"u1": {"s1": &mtgv1.Session{
		Id: "s1",
		Turns: []*mtgv1.Turn{
			{Questions: []*mtgv1.Question{
				{Id: "q1", Text: "Which power bracket should the deck target?", Options: []string{"1 exhibition", "2 core", "3 upgraded", "4 optimized", "5 cEDH"}},
				{Id: "q2", Text: "What should the deck focus on?"},
			}},
			{Answers: []*mtgv1.Answer{
				{QuestionId: "q1", OptionIndex: proto32(3)},
				{QuestionId: "q2", Declined: true},
			}},
		},
	}}}
	s := New(store, sess, fakeDecks{}, auth.UserID, WithClock(func() time.Time { return at }))
	return store, s
}

func proto32(v int32) *int32 { return &v }

// last is the item the fake store wrote last.
func (f *fakeStore) last() feedback.Item { return f.items[len(f.items)-1] }

// TestFeedbackCarriesTheQuestionAndTheAnswer is D-596. A question id
// alone reads nothing after a prompt changes, so the store holds the
// exact wording the reader saw and what the reader answered. Both come
// from the stored session, so neither is the client's word.
func TestFeedbackCarriesTheQuestionAndTheAnswer(t *testing.T) {
	store, s := withAnswered()
	ctx := auth.WithUserID(t.Context(), "u1")
	_, err := s.SubmitFeedback(ctx, connect.NewRequest(&mtgv1.SubmitFeedbackRequest{Feedback: &mtgv1.Feedback{
		Kind:       mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION,
		Verdict:    mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN,
		SessionId:  "s1",
		QuestionId: "q1",
		Reasons:    []string{"bad_options"},
	}}))
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	item := store.last()
	if item.QuestionText != "Which power bracket should the deck target?" {
		t.Errorf("question text = %q", item.QuestionText)
	}
	// The reader answered by option index, and the store holds the words.
	if item.AnswerText != "4 optimized" {
		t.Errorf("answer text = %q, want the option the reader picked", item.AnswerText)
	}
	// The repo writes the uid field from the user it stores under, so the
	// two can not drift. The service passes that user, and this is where
	// it is readable.
	if got := store.uids[len(store.uids)-1]; got != "u1" {
		t.Errorf("stored under %q, want u1", got)
	}
}

func TestFeedbackReadsADeclinedAnswer(t *testing.T) {
	store, s := withAnswered()
	ctx := auth.WithUserID(t.Context(), "u1")
	_, err := s.SubmitFeedback(ctx, connect.NewRequest(&mtgv1.SubmitFeedbackRequest{Feedback: &mtgv1.Feedback{
		Kind:       mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION,
		Verdict:    mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN,
		SessionId:  "s1",
		QuestionId: "q2",
		Reasons:    []string{"unclear"},
	}}))
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	item := store.last()
	if item.QuestionText != "What should the deck focus on?" {
		t.Errorf("question text = %q", item.QuestionText)
	}
	if item.AnswerText != declinedAnswer {
		t.Errorf("answer text = %q, want %q", item.AnswerText, declinedAnswer)
	}
}

// A chat verdict names no question, so it carries no wording and no
// answer. It still carries the user and the session.
func TestChatFeedbackCarriesNoQuestion(t *testing.T) {
	store, s := withAnswered()
	ctx := auth.WithUserID(t.Context(), "u1")
	_, err := s.SubmitFeedback(ctx, connect.NewRequest(&mtgv1.SubmitFeedbackRequest{Feedback: &mtgv1.Feedback{
		Kind:      mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT,
		Verdict:   mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN,
		SessionId: "s1",
		Reasons:   []string{"stuck"},
	}}))
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	item := store.last()
	if item.QuestionText != "" || item.AnswerText != "" {
		t.Errorf("a chat verdict named a question: %q %q", item.QuestionText, item.AnswerText)
	}
	if item.SessionID != "s1" || item.Kind != "chat" {
		t.Errorf("chat verdict = %+v", item)
	}
}
