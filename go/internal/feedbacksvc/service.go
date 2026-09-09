// Package feedbacksvc serves FeedbackService (PR-27, D-557 to D-559).
// SubmitFeedback checks that the object a verdict names belongs to the
// caller, then writes one item under the caller.
package feedbacksvc

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/feedback"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/gzstore"
	"github.com/nkramber/decktome/go/internal/questions"
	"github.com/nkramber/decktome/go/internal/sessions"
)

// MaxTextBytes caps the "Other" text. It is the message cap of the chat,
// agentsvc.MaxMessageBytes, and a test holds the two equal.
const MaxTextBytes = 8 << 10

// Store writes one verdict under a user.
type Store interface {
	Add(ctx context.Context, uid string, item feedback.Item) (string, error)
}

// SessionSource reads the caller's sessions, for a verdict on a question.
type SessionSource interface {
	Get(ctx context.Context, uid, id string) (*mtgv1.Session, error)
}

// DeckSource reads the caller's decks, for a verdict on a summary, a
// card, or a deck.
type DeckSource interface {
	Get(ctx context.Context, uid, id string) (*mtgv1.Deck, error)
}

// Option tunes the server.
type Option func(*Server)

// WithClock replaces time.Now (tests).
func WithClock(now func() time.Time) Option {
	return func(s *Server) { s.now = now }
}

// Server answers FeedbackService requests.
type Server struct {
	mtgv1connect.UnimplementedFeedbackServiceHandler
	store    Store
	sessions SessionSource
	decks    DeckSource
	userFn   auth.UserFunc
	now      func() time.Time
}

// New wires the service. Every source is needed: a question verdict
// reads the session, and the other three read the deck.
func New(store Store, sessionSrc SessionSource, deckSrc DeckSource, userFn auth.UserFunc, opts ...Option) *Server {
	s := &Server{store: store, sessions: sessionSrc, decks: deckSrc, userFn: userFn, now: time.Now}
	for _, o := range opts {
		o(s)
	}
	return s
}

// Prompts names the prompt versions of the moment, so the harvest of
// PR-28 knows which prompt the user judged (D-558). The keys are the
// ones of the eval run files: questions covers the classify and the ask
// roles, and generate covers the build.
func Prompts() map[string]int64 {
	return map[string]int64{"questions": questions.PromptVersion, "generate": generate.PromptVersion}
}

var (
	errNoUser          = errors.New("no user in the request context")
	errNoFeedback      = errors.New("feedback is required")
	errBadKind         = errors.New("kind: name a question, a summary, a card, or a deck")
	errBadVerdict      = errors.New("verdict: name up or down")
	errLongText        = fmt.Errorf("text: the text takes at most %d bytes", MaxTextBytes)
	errUpWithReason    = errors.New("a thumbs up carries no reason and no text")
	errDownNeedsReason = errors.New("a thumbs down needs one reason or a text")
	errNoQuestion      = errors.New("question_id: no question of the session has this id")
	errNoSummary       = errors.New("deck_id: the deck has no description")
	errNoCard          = errors.New("oracle_id: the deck holds no card with this id")
	errNotYourSession  = errors.New("the session is not yours, or there is no such session")
	errNotYourDeck     = errors.New("the deck is not yours, or there is no such deck")
)

func invalid(err error) error { return connect.NewError(connect.CodeInvalidArgument, err) }

// SubmitFeedback stores one verdict under the caller.
func (s *Server) SubmitFeedback(ctx context.Context, req *connect.Request[mtgv1.SubmitFeedbackRequest]) (*connect.Response[mtgv1.SubmitFeedbackResponse], error) {
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	fb := req.Msg.GetFeedback()
	if fb == nil {
		return nil, invalid(errNoFeedback)
	}
	item, err := itemOf(fb)
	if err != nil {
		return nil, invalid(err)
	}
	sess, deck, err := s.checkOwner(ctx, uid, fb)
	if err != nil {
		return nil, err
	}
	// The exact wording the reader saw, and what the reader answered
	// (D-596). A question id alone reads nothing after a prompt changes.
	// Both come from the stored session, so neither is the client's word.
	item.QuestionText, item.AnswerText = questionContext(sess, item.QuestionID)
	// The object the verdict names, as it stood at this moment (D-635).
	// A reader deletes a session or a deck, and the verdict outlives it.
	item.Session, item.Deck = sess, deck
	item.Prompts = Prompts()
	item.CreatedAt = s.now()
	id, err := s.store.Add(ctx, uid, item)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.SubmitFeedbackResponse{FeedbackId: id}), nil
}

// itemOf reads the fields of a verdict and refuses a bad one. The ids
// the kind does not use must stay empty, so a stored item names one
// object and nothing else.
func itemOf(fb *mtgv1.Feedback) (feedback.Item, error) {
	kind := fb.GetKind()
	item := feedback.Item{
		Kind:       feedback.KindName(kind),
		Verdict:    feedback.VerdictName(fb.GetVerdict()),
		SessionID:  strings.TrimSpace(fb.GetSessionId()),
		QuestionID: strings.TrimSpace(fb.GetQuestionId()),
		DeckID:     strings.TrimSpace(fb.GetDeckId()),
		OracleID:   strings.TrimSpace(fb.GetOracleId()),
		Text:       strings.TrimSpace(fb.GetText()),
	}
	if item.Kind == "" {
		return item, errBadKind
	}
	if item.Verdict == "" {
		return item, errBadVerdict
	}
	if len(item.Text) > MaxTextBytes {
		return item, errLongText
	}
	seen := map[string]bool{}
	for _, key := range fb.GetReasons() {
		key = strings.TrimSpace(key)
		if !feedback.KnownReason(kind, key) {
			return item, fmt.Errorf("reasons: %q is not a reason for this kind", key)
		}
		if !seen[key] {
			seen[key] = true
			item.Reasons = append(item.Reasons, key)
		}
	}
	down := fb.GetVerdict() == mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN
	if !down && (len(item.Reasons) > 0 || item.Text != "") {
		return item, errUpWithReason
	}
	if down && len(item.Reasons) == 0 && item.Text == "" {
		return item, errDownNeedsReason
	}
	question := kind == mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION
	chat := kind == mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT
	card := kind == mtgv1.FeedbackKind_FEEDBACK_KIND_CARD
	// A chat verdict names the session and no question, and it names no
	// deck: a chat that stops has none (D-594).
	if err := wantID("session_id", item.SessionID, question || chat); err != nil {
		return item, err
	}
	if err := wantID("question_id", item.QuestionID, question); err != nil {
		return item, err
	}
	if err := wantID("deck_id", item.DeckID, !question && !chat); err != nil {
		return item, err
	}
	if err := wantID("oracle_id", item.OracleID, card); err != nil {
		return item, err
	}
	return item, nil
}

// wantID checks one id field: a kind that uses it needs a valid one, and
// a kind that does not use it needs it empty.
func wantID(field, id string, used bool) error {
	switch {
	case used && id == "":
		return fmt.Errorf("%s: give the id", field)
	case used && !gzstore.ValidID(id):
		return fmt.Errorf("%s: %w", field, gzstore.ErrBadID)
	case !used && id != "":
		return fmt.Errorf("%s: this kind takes no %s", field, field)
	}
	return nil
}

// checkOwner reads the object the verdict names under the caller. The
// stores answer per user, so another user's id and an unknown id read
// the same, and neither is the caller's: PermissionDenied for both.
//
// It answers the object it read, because the verdict stores a snapshot
// of it (D-635). The read happens either way, so the snapshot costs no
// call.
func (s *Server) checkOwner(ctx context.Context, uid string, fb *mtgv1.Feedback) (*mtgv1.Session, *mtgv1.Deck, error) {
	kind := fb.GetKind()
	if kind == mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION || kind == mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT {
		sess, err := s.sessions.Get(ctx, uid, strings.TrimSpace(fb.GetSessionId()))
		if errors.Is(err, sessions.ErrNotFound) {
			return nil, nil, connect.NewError(connect.CodePermissionDenied, errNotYourSession)
		}
		if err != nil {
			return nil, nil, connect.NewError(connect.CodeInternal, err)
		}
		// A chat verdict reads the whole session, so it names no question.
		if kind == mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT {
			return sess, nil, nil
		}
		if !hasQuestion(sess, strings.TrimSpace(fb.GetQuestionId())) {
			return nil, nil, invalid(errNoQuestion)
		}
		return sess, nil, nil
	}
	deck, err := s.decks.Get(ctx, uid, strings.TrimSpace(fb.GetDeckId()))
	if errors.Is(err, decks.ErrNotFound) {
		return nil, nil, connect.NewError(connect.CodePermissionDenied, errNotYourDeck)
	}
	if err != nil {
		return nil, nil, connect.NewError(connect.CodeInternal, err)
	}
	switch fb.GetKind() {
	case mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY:
		if strings.TrimSpace(deck.GetSummary()) == "" {
			return nil, nil, invalid(errNoSummary)
		}
	case mtgv1.FeedbackKind_FEEDBACK_KIND_CARD:
		if !hasCard(deck, strings.TrimSpace(fb.GetOracleId())) {
			return nil, nil, invalid(errNoCard)
		}
	}
	return nil, deck, nil
}

// questionContext reads the exact wording of one question and the answer
// the reader gave it, from the stored session (D-596). It answers two
// empty strings for a verdict that names no question, and for a question
// the reader has not answered.
func questionContext(sess *mtgv1.Session, questionID string) (text, answer string) {
	if sess == nil || questionID == "" {
		return "", ""
	}
	for _, turn := range sess.GetTurns() {
		for _, q := range turn.GetQuestions() {
			if q.GetId() == questionID {
				text = q.GetText()
			}
		}
		for _, a := range turn.GetAnswers() {
			if a.GetQuestionId() != questionID {
				continue
			}
			switch {
			case a.GetDeclined():
				answer = declinedAnswer
			case strings.TrimSpace(a.GetText()) != "":
				answer = strings.TrimSpace(a.GetText())
			default:
				answer = optionText(sess, questionID, a)
			}
		}
	}
	return text, answer
}

// declinedAnswer is what the store holds for a reader who declined. The
// thread reads "You decide", and the store reads one term.
const declinedAnswer = "(declined)"

// optionText reads the option an answer named by index, from the
// question the session holds.
func optionText(sess *mtgv1.Session, questionID string, a *mtgv1.Answer) string {
	if a.OptionIndex == nil {
		return ""
	}
	for _, turn := range sess.GetTurns() {
		for _, q := range turn.GetQuestions() {
			if q.GetId() != questionID {
				continue
			}
			if i := int(a.GetOptionIndex()); i >= 0 && i < len(q.GetOptions()) {
				return q.GetOptions()[i]
			}
		}
	}
	return ""
}

func hasQuestion(sess *mtgv1.Session, id string) bool {
	for _, turn := range sess.GetTurns() {
		for _, q := range turn.GetQuestions() {
			if q.GetId() == id {
				return true
			}
		}
	}
	return false
}

// hasCard reports whether the deck holds the card: in the main deck, the
// sideboard, the upgrades, or the command zone.
func hasCard(deck *mtgv1.Deck, oracleID string) bool {
	for _, id := range deck.GetCommanderOracleIds() {
		if id == oracleID {
			return true
		}
	}
	for _, list := range [][]*mtgv1.DeckCard{deck.GetCards(), deck.GetSideboard(), deck.GetUpgrades()} {
		for _, c := range list {
			if c.GetOracleId() == oracleID {
				return true
			}
		}
	}
	return false
}

func (s *Server) user(ctx context.Context) string {
	if s.userFn == nil {
		return ""
	}
	return s.userFn(ctx)
}
