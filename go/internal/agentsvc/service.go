// Package agentsvc serves AgentService (roadmap PR-7). It is the call
// site of internal/questions: one user message in, the next questions
// out, and the whole conversation stored.
//
// The build itself lands with PR-8. A session that reaches a full slot
// set reports SESSION_STATUS_READY and stops there.
package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// Store holds the conversations (D-74).
type Store interface {
	// NewID reserves a session id without a write.
	NewID(uid string) string
	// Put writes the public session and the private state together. It
	// returns sessions.ErrConflict when the stored version is not the
	// expected one, and then writes nothing (H-7).
	Put(ctx context.Context, uid string, s *mtgv1.Session, snap questions.Snapshot, expected int64) error
	// Get reads the public session alone.
	Get(ctx context.Context, uid, id string) (*mtgv1.Session, error)
	// GetState reads the session, its private state, and the version
	// that the next Put must expect.
	GetState(ctx context.Context, uid, id string) (*mtgv1.Session, questions.Snapshot, int64, error)
}

// UserFunc reads the caller's user id from the request context.
type UserFunc func(ctx context.Context) string

// IndexSource hands out the current card index.
type IndexSource interface {
	Current() *cards.Index
}

// CollectionSource gives the owned count per Oracle id (D-37).
type CollectionSource interface {
	OracleCounts(ctx context.Context, userID, collectionID string) (map[string]int32, error)
}

// Server answers AgentService requests.
type Server struct {
	mtgv1connect.UnimplementedAgentServiceHandler
	cat         *questions.Catalog
	client      *llm.Client
	store       Store
	userFn      UserFunc
	index       IndexSource
	builder     *candidates.Builder
	decks       DeckBuilder
	buildLimit  time.Duration
	collections CollectionSource
	prices      *llm.PriceTable
	now         func() time.Time
	log         *slog.Logger
}

// Option configures the server.
type Option func(*Server)

// WithCandidates wires the PR-6 hints. Without it, every question that
// names a value drops that clause and falls back.
func WithCandidates(index IndexSource, b *candidates.Builder) Option {
	return func(s *Server) { s.index, s.builder = index, b }
}

// DeckBuilder writes the deck for a ready session (roadmap PR-8).
// internal/generate holds the one implementation, and the interface keeps
// agentsvc testable without a provider.
type DeckBuilder interface {
	Build(ctx context.Context, req generate.Request, acc *llm.Accumulator) (*generate.Result, error)
}

// WithDecks wires the generator. Without it, a ready session reports that
// every slot is filled and builds nothing, which is the PR-7 behavior.
func WithDecks(b DeckBuilder) Option {
	return func(s *Server) { s.decks = b }
}

// DefaultBuildLimit caps one build. The llm client already caps each call
// at three minutes, so a generate and a repair together can hold the
// stream for six. A build is about two minutes when it goes well, and a
// user waiting longer than this is better served by an error than by a
// stream that does not end (D-235).
const DefaultBuildLimit = 4 * time.Minute

// WithBuildTimeout caps one build. Zero keeps DefaultBuildLimit.
func WithBuildTimeout(d time.Duration) Option {
	return func(s *Server) { s.buildLimit = d }
}

// WithCollections wires the owned counts, which the hints read.
func WithCollections(src CollectionSource) Option {
	return func(s *Server) { s.collections = src }
}

// WithPrices makes the usage event carry a cost (M-1).
func WithPrices(p *llm.PriceTable) Option { return func(s *Server) { s.prices = p } }

// WithLogger sets the logger that carries the M-4 rows.
func WithLogger(l *slog.Logger) Option { return func(s *Server) { s.log = l } }

// WithClock replaces time.Now (tests).
func WithClock(f func() time.Time) Option { return func(s *Server) { s.now = f } }

// New wires the service.
func New(cat *questions.Catalog, client *llm.Client, store Store, userFn UserFunc, opts ...Option) (*Server, error) {
	if cat == nil || client == nil || store == nil || userFn == nil {
		return nil, errors.New("agentsvc: New needs a catalog, a client, a store, and a user function")
	}
	s := &Server{
		cat: cat, client: client, store: store, userFn: userFn,
		now: func() time.Time { return time.Now().UTC() },
		log: slog.Default(),
	}
	for _, o := range opts {
		o(s)
	}
	return s, nil
}

var (
	errNoUser    = errors.New("no user in the request context")
	errNoMessage = errors.New("message or answers are required")
)

// GetSession returns one stored conversation.
func (s *Server) GetSession(ctx context.Context, req *connect.Request[mtgv1.GetSessionRequest]) (*connect.Response[mtgv1.GetSessionResponse], error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	if req.Msg.GetSessionId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("session_id is required"))
	}
	session, err := s.store.Get(ctx, uid, req.Msg.GetSessionId())
	if err != nil {
		return nil, storeError(err)
	}
	return connect.NewResponse(&mtgv1.GetSessionResponse{Session: session}), nil
}

// Chat runs one turn and streams what it produced.
//
// Event order: session_started (first message only), then one question
// event for each question, then the slots, then the usage. A model
// failure ends the turn with a failure event, and the session keeps every
// slot the turn already filled.
//
// Two overlapping calls on one session both run the turn, and only the
// first Put lands. The second one gets CodeAborted and stored nothing, so
// the asked rows of the first turn survive (H-7).
func (s *Server) Chat(ctx context.Context, req *connect.Request[mtgv1.ChatRequest], stream *connect.ServerStream[mtgv1.ChatResponse]) error {
	uid := s.userFn(ctx)
	if uid == "" {
		return connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	if strings.TrimSpace(req.Msg.GetMessage()) == "" && len(req.Msg.GetAnswers()) == 0 {
		return connect.NewError(connect.CodeInvalidArgument, errNoMessage)
	}

	session, snap, version, err := s.load(ctx, uid, req.Msg)
	if err != nil {
		return err
	}
	if session.GetCreatedAt() == nil {
		if err := stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_SessionStarted{SessionStarted: session.GetId()},
		}); err != nil {
			return err
		}
		session.CreatedAt = timestamppb.New(s.now())
	}

	st := questions.Restore(session.GetId(), session.GetSlots(), snap)
	message := withAnswers(req.Msg.GetMessage(), req.Msg.GetAnswers(), session)

	hints := s.hints(ctx, uid, session, st)
	s.facts(session, st, hints)
	agent, err := s.agent(hints)
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(s.prices)
	res, turnErr := agent.Turn(ctx, st, message, acc)
	// The turn is stored either way. A failed turn keeps the slots the
	// classify call already filled, so a retry does not start again.
	session.Slots = st.Slots
	session.Usage = addUsage(session.GetUsage(), acc.Report())
	session.UpdatedAt = timestamppb.New(s.now())
	turn := &mtgv1.Turn{
		UserMessage: req.Msg.GetMessage(),
		Answers:     req.Msg.GetAnswers(),
		At:          timestamppb.New(s.now()),
	}
	if turnErr == nil {
		turn.Questions = res.Questions
		session.Status = mtgv1.SessionStatus_SESSION_STATUS_ASKING
		if res.Ready {
			session.Status = mtgv1.SessionStatus_SESSION_STATUS_READY
		}
	}
	session.Turns = append(session.Turns, turn)
	if err := s.store.Put(ctx, uid, session, st.Snapshot(), version); err != nil {
		return storeError(err)
	}

	if turnErr != nil {
		s.log.ErrorContext(ctx, "turn failed", "session", session.GetId(), "err", turnErr)
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Failure{Failure: failure(turnErr)},
		})
	}
	for _, q := range res.Questions {
		if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Question{Question: q}}); err != nil {
			return err
		}
	}
	if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Slots{Slots: session.GetSlots()}}); err != nil {
		return err
	}
	if res.Ready {
		if err := s.sendDeck(ctx, uid, session, st, acc, stream); err != nil {
			return err
		}
	}
	return stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Usage{Usage: session.GetUsage()}})
}

// load reads the named session, or makes a new one. The version is the
// one Put must expect, and a new session expects 0.
func (s *Server) load(ctx context.Context, uid string, msg *mtgv1.ChatRequest) (*mtgv1.Session, questions.Snapshot, int64, error) {
	if id := msg.GetSessionId(); id != "" {
		session, snap, version, err := s.store.GetState(ctx, uid, id)
		if err != nil {
			return nil, questions.Snapshot{}, 0, storeError(err)
		}
		return session, snap, version, nil
	}
	id := s.store.NewID(uid)
	session := &mtgv1.Session{
		Id:           id,
		CollectionId: msg.GetCollectionId(),
		Status:       mtgv1.SessionStatus_SESSION_STATUS_ASKING,
	}
	// A user with no collection never gets the card-pool question (D-37).
	snap := questions.Snapshot{Version: questions.SnapshotVersion}
	snap.Ctx.HasCollection = msg.GetCollectionId() != ""
	return session, snap, 0, nil
}

// agent builds the turn's agent.
func (s *Server) agent(hints *questions.CandidateHints) (*questions.Agent, error) {
	opts := []questions.AgentOption{questions.WithLogger(s.log)}
	if hints != nil {
		opts = append(opts, questions.WithHints(hints))
	}
	agent, err := questions.NewAgent(s.cat, s.client, opts...)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return agent, nil
}

// facts sets the planner triggers this package owns. The classifier can
// not answer them: they come from the stored session and from PR-6.
//
// Each one gates a catalog row. Without them the row is dead code, which
// is the defect the live run of 2026-08-24 found on Context.Suggested.
//
// The PR-6 facts read the slots, and the agent reads them again after
// the classify call fills the slots of this turn (M-6). This call covers
// the turn's first plan with the stored slots.
func (s *Server) facts(session *mtgv1.Session, st *questions.State, hints *questions.CandidateHints) {
	// A deck exists, so the user may ask for another version (PR-9).
	st.Ctx.AfterBuild = len(session.GetDeckIds()) > 0
	if hints == nil {
		return
	}
	questions.RefreshFacts(st, hints)
}

// hints answers the placeholder values from the card index. It returns
// nil when no index is loaded, and every clause that needs a value is
// then dropped.
func (s *Server) hints(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State) *questions.CandidateHints {
	if s.index == nil || s.builder == nil {
		return nil
	}
	idx := s.index.Current()
	if idx == nil {
		return nil
	}
	h := &questions.CandidateHints{
		Index:   idx,
		Builder: s.builder,
		Format:  st.Slots.GetFormat().GetId(),
		Colors:  st.Slots.GetColors(),
		Pool:    st.Slots.GetPoolRule(),
		Log:     s.log,
	}
	if s.collections != nil && session.GetCollectionId() != "" {
		owned, err := s.collections.OracleCounts(ctx, uid, session.GetCollectionId())
		if err != nil {
			// A missing collection must not end the turn. The questions
			// that need owned counts fall back to their short wording.
			s.log.WarnContext(ctx, "owned counts unavailable", "collection", session.GetCollectionId(), "err", err)
		} else {
			h.Owned = owned
		}
	}
	return h
}

// withAnswers folds the structured replies into the message the
// classifier reads. Each line names the question and the answer, so the
// classifier maps the answer to the right slot.
//
// The Answer contract follows the proto: text wins when it is not empty,
// and option_index names an option only when text is empty. The proto
// says -1 means free text. A client that omits option_index sends 0,
// which is a valid option. So a client must send -1, or text, for a
// free-text reply. A free-text reply with option_index 0 and text set
// uses the text.
func withAnswers(message string, answers []*mtgv1.Answer, session *mtgv1.Session) string {
	if len(answers) == 0 {
		return message
	}
	asked := map[string]*mtgv1.Question{}
	for _, turn := range session.GetTurns() {
		for _, q := range turn.GetQuestions() {
			asked[q.GetId()] = q
		}
	}
	lines := make([]string, 0, len(answers)+1)
	for _, a := range answers {
		text := strings.TrimSpace(a.GetText())
		q := asked[a.GetQuestionId()]
		if text == "" && q != nil {
			if i := int(a.GetOptionIndex()); i >= 0 && i < len(q.GetOptions()) {
				text = q.GetOptions()[i]
			}
		}
		if text == "" {
			continue
		}
		if q == nil {
			lines = append(lines, text)
			continue
		}
		lines = append(lines, fmt.Sprintf("%s %s", q.GetText(), text))
	}
	if message = strings.TrimSpace(message); message != "" {
		lines = append(lines, message)
	}
	return strings.Join(lines, "\n")
}

// addUsage sums one turn's report into the session total (M-1). A turn
// whose provider reported nothing leaves priced false, which is not the
// same fact as a zero cost. A turn with no calls at all, such as a
// frozen session, reports no cost either, and it must not flip the
// flag (M-8).
func addUsage(total *mtgv1.Usage, r llm.Report) *mtgv1.Usage {
	if total == nil {
		total = &mtgv1.Usage{Priced: true}
	}
	total.Calls += int32(r.Calls)
	if r.Tokens != nil {
		total.InputTokens += r.Tokens.InputTokens
		total.CachedInputTokens += r.Tokens.CachedInputTokens
		total.OutputTokens += r.Tokens.OutputTokens
		total.ReasoningTokens += r.Tokens.ReasoningTokens
	}
	switch {
	case r.CostUSD != nil:
		total.CostUsd += *r.CostUSD
	case r.Calls > 0:
		total.Priced = false
	}
	return total
}

// failure maps a model error onto the code the UI acts on.
func failure(err error) *mtgv1.AgentError {
	class := llm.ClassOf(err)
	out := &mtgv1.AgentError{
		Code:      "llm_" + class.String(),
		Message:   "The model call failed. This one does not succeed on a retry.",
		Retryable: class == llm.ClassTransient || class == llm.ClassBudget,
	}
	switch class {
	case llm.ClassTransient, llm.ClassBudget:
		out.Message = "The model call failed. Send the message again."
	case llm.ClassRefusal:
		out.Message = "The model declined that request. Say it in other words."
	}
	return out
}

// storeError maps a store failure onto a Connect code.
func storeError(err error) error {
	var connectErr *connect.Error
	if errors.As(err, &connectErr) {
		return err
	}
	if errors.Is(err, sessions.ErrNotFound) {
		return connect.NewError(connect.CodeNotFound, err)
	}
	if errors.Is(err, sessions.ErrConflict) {
		return connect.NewError(connect.CodeAborted,
			fmt.Errorf("the session is busy with another turn, send the message again: %w", err))
	}
	return connect.NewError(connect.CodeInternal, err)
}
