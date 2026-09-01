// Package agentsvc serves AgentService. It is the call site of
// internal/questions: one user message in, the next questions out, and
// the whole conversation stored.
//
// A session that reaches a full slot set reports SESSION_STATUS_READY,
// builds, and then reports SESSION_STATUS_BUILT.
package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/auth"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// Store holds the conversations (D-74).
type Store interface {
	// NewID reserves a session id without a write.
	NewID(uid string) string
	// Put writes the public session and the private state together. It
	// returns sessions.ErrConflict when the stored version is not the
	// expected one, and then writes nothing.
	Put(ctx context.Context, uid string, s *mtgv1.Session, snap questions.Snapshot, expected int64) error
	// Get reads the public session alone.
	Get(ctx context.Context, uid, id string) (*mtgv1.Session, error)
	// GetState reads the session, its private state, and the version
	// that the next Put must expect.
	GetState(ctx context.Context, uid, id string) (*mtgv1.Session, questions.Snapshot, int64, error)
}

// PreconSource hands out the precon set for the current card index, or
// nil before the first snapshot loads. The api resolves it late, because
// the precon lists need an index and the index lands after the server
// starts.
type PreconSource interface {
	Current() *precons.Set
}

// MaxMessageBytes caps one message, one answer text, and the message
// and the answers together. A deck request is a few sentences, and a
// whole ManaBox export goes through ImportCollection, not Chat.
const MaxMessageBytes = 8 << 10

// MaxAnswers caps the answers of one turn. A turn asks at most a few
// questions, so more answers than this is not a reply to them.
const MaxAnswers = 16

// maxFoldedBytes caps the classify message after the answers are folded
// in with their question texts.
const maxFoldedBytes = 2 * MaxMessageBytes

// DefaultChatLimit caps the Chat turns one process runs at once. A turn
// holds a model call open for minutes, and every turn is billed.
const DefaultChatLimit = 8

// storeLimit bounds one store write that runs after a paid model call.
// The write runs detached from the client, so it needs its own clock
// (D-303).
const storeLimit = 30 * time.Second

// CollectionSource gives the owned count per Oracle id (D-37).
type CollectionSource interface {
	OracleCounts(ctx context.Context, userID, collectionID string) (map[string]int32, error)
	// OwnedPrintings maps each Oracle id to the printing ids the user
	// holds, so a deck card can show the printing the user owns (D-299).
	OwnedPrintings(ctx context.Context, userID, collectionID string) (map[string][]string, error)
}

// Server answers AgentService requests.
type Server struct {
	mtgv1connect.UnimplementedAgentServiceHandler
	cat        *questions.Catalog
	client     *llm.Client
	store      Store
	userFn     auth.UserFunc
	index      cardsvc.IndexSource
	builder    *candidates.Builder
	decks      DeckBuilder
	deckStore  DeckStore
	preconSrc  PreconSource
	buildLimit time.Duration
	// turns is the concurrency gate: one token per running Chat turn.
	turns chan struct{}
	// building marks each session with a build in flight, keyed by
	// uid and session id. A turn during a build is refused (D-303).
	building    sync.Map
	collections CollectionSource
	prices      *llm.PriceTable
	now         func() time.Time
	log         *slog.Logger
}

// Option configures the server.
type Option func(*Server)

// WithCandidates wires the candidate hints. Without it, every question
// that names a value drops that clause and falls back.
func WithCandidates(index cardsvc.IndexSource, b *candidates.Builder) Option {
	return func(s *Server) { s.index, s.builder = index, b }
}

// DeckBuilder writes the deck for a ready session. internal/generate
// holds the one implementation, and the interface keeps agentsvc
// testable without a provider.
type DeckBuilder interface {
	Build(ctx context.Context, req generate.Request, acc *llm.Accumulator) (*generate.Result, error)
}

// WithDecks wires the generator. Without it, a ready session reports that
// every slot is filled and builds nothing.
func WithDecks(b DeckBuilder) Option {
	return func(s *Server) { s.decks = b }
}

// DeckStore holds the decks a build produced. Without it a deck streams
// to the user and is gone: session.deck_ids stays empty and AfterBuild
// is never true (D-245).
type DeckStore interface {
	// NewID reserves a deck id without a write. The deck carries its own
	// id, so the build needs one before it runs.
	NewID(uid string) string
	Put(ctx context.Context, uid string, d *mtgv1.Deck) error
	// Get reads one kept deck. A revision starts from the deck the user
	// read (D-283).
	Get(ctx context.Context, uid, id string) (*mtgv1.Deck, error)
}

// WithPreconSource wires the late-bound precon set. Without it the
// precon share of D-218 does not run, and an upgrade request is served
// as an ordinary owned-first build (D-247).
func WithPreconSource(src PreconSource) Option {
	return func(s *Server) { s.preconSrc = src }
}

// WithChatLimit caps the Chat turns that run at once. Zero keeps
// DefaultChatLimit. A refused turn answers ResourceExhausted.
func WithChatLimit(n int) Option {
	return func(s *Server) {
		if n > 0 {
			s.turns = make(chan struct{}, n)
		}
	}
}

// WithDeckStore wires the deck store. Without it the build still returns
// a deck, and nothing keeps it.
func WithDeckStore(s DeckStore) Option {
	return func(srv *Server) { srv.deckStore = s }
}

// DefaultBuildLimit caps one build. The llm client caps each call at
// three minutes, so a generate and a repair together can hold the stream
// for six. A build is about two minutes when it goes well, and a user
// waiting longer than this is better served by an error than by a
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

// WithPrices makes the usage event carry a cost.
func WithPrices(p *llm.PriceTable) Option { return func(s *Server) { s.prices = p } }

// WithLogger sets the logger.
func WithLogger(l *slog.Logger) Option { return func(s *Server) { s.log = l } }

// WithClock replaces time.Now (tests).
func WithClock(f func() time.Time) Option { return func(s *Server) { s.now = f } }

// New wires the service.
func New(cat *questions.Catalog, client *llm.Client, store Store, userFn auth.UserFunc, opts ...Option) (*Server, error) {
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
	if s.turns == nil {
		s.turns = make(chan struct{}, DefaultChatLimit)
	}
	return s, nil
}

// preconSet returns the precon set for the current index, or nil.
func (s *Server) preconSet() *precons.Set {
	if s.preconSrc == nil {
		return nil
	}
	return s.preconSrc.Current()
}

// buildDeadline is the cap on one build.
func (s *Server) buildDeadline() time.Duration {
	if s.buildLimit > 0 {
		return s.buildLimit
	}
	return DefaultBuildLimit
}

// detached returns a context that outlives the client with its own
// clock. A build and every store write after a paid model call run on
// one, so a disconnect still stores the turn and the deck (D-303).
func detached(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.WithoutCancel(ctx), d)
}

var (
	errNoUser          = errors.New("no user in the request context")
	errNoMessage       = errors.New("message or answers are required")
	errTooLong         = fmt.Errorf("the message and the answers are longer than %d bytes together, or one is alone", MaxMessageBytes)
	errTooManyAnswers  = fmt.Errorf("more than %d answers in one turn", MaxAnswers)
	errBusy            = errors.New("the server runs its limit of turns at once, send the message again in a moment")
	errSessionBig      = errors.New("this conversation is too long to continue, start a new session")
	errBuildInProgress = errors.New("a build is in progress")
	errNoSessionID     = errors.New("session_id is required")
	errBadSessionID    = fmt.Errorf("session_id: %w", gzstore.ErrBadID)
	errBadCollectionID = fmt.Errorf("collection_id: %w", gzstore.ErrBadID)
)

// tooLong reports whether the message, any answer text, or the message
// and the answers together pass the cap, or the answers pass their
// count. It returns the error to answer with, or nil.
func tooLong(req *mtgv1.ChatRequest) error {
	if len(req.GetAnswers()) > MaxAnswers {
		return errTooManyAnswers
	}
	total := len(req.GetMessage())
	if total > MaxMessageBytes {
		return errTooLong
	}
	for _, a := range req.GetAnswers() {
		if len(a.GetText()) > MaxMessageBytes {
			return errTooLong
		}
		total += len(a.GetText())
	}
	if total > MaxMessageBytes {
		return errTooLong
	}
	return nil
}

// GetSession returns one stored conversation.
func (s *Server) GetSession(ctx context.Context, req *connect.Request[mtgv1.GetSessionRequest]) (*connect.Response[mtgv1.GetSessionResponse], error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	id := req.Msg.GetSessionId()
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoSessionID)
	}
	if !gzstore.ValidID(id) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadSessionID)
	}
	session, err := s.store.Get(ctx, uid, id)
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
// the asked rows of the first turn survive. A turn that arrives while a
// build runs is refused before any model call (D-303).
func (s *Server) Chat(ctx context.Context, req *connect.Request[mtgv1.ChatRequest], stream *connect.ServerStream[mtgv1.ChatResponse]) error {
	uid := s.userFn(ctx)
	if uid == "" {
		return connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	if strings.TrimSpace(req.Msg.GetMessage()) == "" && len(req.Msg.GetAnswers()) == 0 {
		return connect.NewError(connect.CodeInvalidArgument, errNoMessage)
	}
	if err := tooLong(req.Msg); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	if id := req.Msg.GetSessionId(); id != "" && !gzstore.ValidID(id) {
		return connect.NewError(connect.CodeInvalidArgument, errBadSessionID)
	}
	if id := req.Msg.GetCollectionId(); id != "" && !gzstore.ValidID(id) {
		return connect.NewError(connect.CodeInvalidArgument, errBadCollectionID)
	}
	if id := req.Msg.GetSessionId(); id != "" {
		if _, busy := s.building.Load(buildKey(uid, id)); busy {
			return connect.NewError(connect.CodeAborted, errBuildInProgress)
		}
	}
	// The gate is non-blocking: a caller past the limit hears it at once
	// instead of a queue that holds the connection open.
	select {
	case s.turns <- struct{}{}:
		defer func() { <-s.turns }()
	default:
		return connect.NewError(connect.CodeResourceExhausted, errBusy)
	}

	session, snap, version, owned, collectionGone, err := s.load(ctx, uid, req.Msg)
	if err != nil {
		return err
	}
	if collectionGone {
		if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{
			Status: "the collection this deck was built from is gone, so this chat builds from the whole card database now"}}); err != nil {
			return err
		}
	}
	if req.Msg.GetSessionId() == "" {
		if err := stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_SessionStarted{SessionStarted: session.GetId()},
		}); err != nil {
			return err
		}
	}
	if session.GetCreatedAt() == nil {
		session.CreatedAt = timestamppb.New(s.now())
	}

	st := questions.Restore(session.GetId(), session.GetSlots(), snap)
	// A pool rule the request carried needs no question (D-359).
	if session.GetSlots().GetPoolRule() != mtgv1.PoolRule_POOL_RULE_UNSPECIFIED &&
		st.Slots.GetSlotStates()["pool_rule"] == mtgv1.SlotState_SLOT_STATE_UNSPECIFIED {
		st.Close("pool_rule")
	}
	// A bare "no" to a yes-or-no question is a whole answer, and the
	// classifier reads it as neither a value nor a decline (D-352). The
	// pairing is exact here, because the answer names its question.
	for _, key := range declineNegatives(st, req.Msg.GetAnswers(), session) {
		s.log.InfoContext(ctx, "the user answered a yes-or-no question with a bare negative, so the key closed",
			"session", session.GetId(), "key", key)
	}
	message := withAnswers(req.Msg.GetMessage(), req.Msg.GetAnswers(), session)
	if len(message) > maxFoldedBytes {
		return connect.NewError(connect.CodeInvalidArgument, errTooLong)
	}
	// The slots before the turn. A turn after a build that changes none
	// of them is a revision of the deck, and one that changes any is a
	// full rebuild (D-241).
	slotsBefore := proto.Clone(session.GetSlots()).(*mtgv1.Slots)

	hints := s.hints(session, st, owned)
	s.facts(session, st, hints)
	agent, err := s.agent(hints)
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(s.prices)
	var stalled []string
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
		// An option that names a card carries its Oracle id, so the UI
		// can show the art and the rules text of a commander offer (D-287).
		if s.index != nil {
			cardOptions(res.Questions, s.index.Current())
		}
		turn.Questions = res.Questions
		// A turn that asks nothing new and is not ready can never move
		// again: the agent does not repeat a question it already asked,
		// so the questions that are out stay out for good, and no build
		// ever starts. The reader sees a chat that does nothing. Close
		// those questions with no value and build (D-351).
		if !res.Ready && len(res.Questions) == 0 {
			if closed, _ := st.CloseStalled(); len(closed) > 0 {
				s.log.WarnContext(ctx, "a turn asked nothing and was not ready, so the open questions were closed",
					"session", session.GetId(), "keys", closed)
				res.Ready = st.Ready(s.cat)
				session.Slots = st.Slots
				stalled = closed
			}
		}
		session.Status = mtgv1.SessionStatus_SESSION_STATUS_ASKING
		if res.Ready {
			session.Status = mtgv1.SessionStatus_SESSION_STATUS_READY
		}
	}
	session.Turns = append(session.Turns, turn)
	// The classify call is paid, so the write that records it runs
	// detached from the client (D-303).
	sctx, cancel := detached(ctx, storeLimit)
	err = s.store.Put(sctx, uid, session, st.Snapshot(), version)
	cancel()
	if err != nil {
		return storeError(err)
	}

	if turnErr != nil {
		s.log.ErrorContext(ctx, "turn failed", "session", session.GetId(), "err", turnErr)
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Failure{Failure: failure(turnErr)},
		})
	}
	// The pool offered no commander, so the agent took the choice (D-127).
	// Silence there reads as a bug, and the reader must hear it (D-366).
	if res.ChoseCommander {
		if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{
			Status: "I have no more commanders that fit this deck, so I chose one for you"}}); err != nil {
			return err
		}
	}
	if len(stalled) > 0 {
		if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{
			Status: "I did not read an answer to every question, so I am building with what I have"}}); err != nil {
			return err
		}
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
		// The session was stored before the build, so the deck id needs a
		// second write. version+1 is what that Put stored (D-245). That
		// write carries the post-turn state, or the next turn would lose
		// this turn's asked rows and repeat a question. It also carries
		// the built status, and storeDeck writes it only when a deck was
		// kept.
		key := buildKey(uid, session.GetId())
		s.building.Store(key, struct{}{})
		defer s.building.Delete(key)
		before := len(session.GetDeckIds())
		session.Status = mtgv1.SessionStatus_SESSION_STATUS_BUILT
		var err error
		switch {
		case before > 0 && !slotsChanged(slotsBefore, session.GetSlots()):
			// A message after a build with no slot change asks for a
			// change to the deck the user read (D-283).
			err = s.sendRevision(ctx, uid, session, st, st.Snapshot(), version+1, owned, acc, message, turn, stream)
		case before > 0:
			if err = stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{
				Status: "a deck setting changed, so the deck is built again from the start"}}); err != nil {
				return err
			}
			err = s.sendDeck(ctx, uid, session, st, st.Snapshot(), version+1, owned, acc, stream)
		default:
			err = s.sendDeck(ctx, uid, session, st, st.Snapshot(), version+1, owned, acc, stream)
		}
		if err != nil {
			return err
		}
	}
	return stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Usage{Usage: session.GetUsage()}})
}

// buildKey names one session in the in-flight build map.
func buildKey(uid, sessionID string) string { return uid + "/" + sessionID }

// load reads the named session, or makes a new one, and reads the owned
// counts of its collection once for the whole turn. The version is the
// one Put must expect, and a new session expects 0. A new session that
// names a collection the store does not hold is NotFound.
func (s *Server) load(ctx context.Context, uid string, msg *mtgv1.ChatRequest) (*mtgv1.Session, questions.Snapshot, int64, map[string]int32, bool, error) {
	if id := msg.GetSessionId(); id != "" {
		session, snap, version, err := s.store.GetState(ctx, uid, id)
		if err != nil {
			return nil, questions.Snapshot{}, 0, nil, false, storeError(err)
		}
		owned, err := s.ownedCounts(ctx, uid, session.GetCollectionId())
		if err != nil {
			// A collection that left after the session started must not
			// end the turn. The questions that need owned counts fall
			// back to their short wording.
			s.log.WarnContext(ctx, "owned counts unavailable", "collection", session.GetCollectionId(), "err", err)
		}
		// A deleted collection is not a failure of this turn. The chat
		// builds from the whole card database from now on, and it says
		// so once (D-347). The session forgets the collection, so every
		// later turn reads no collection at all.
		gone := err != nil && status.Code(err) == codes.NotFound
		if gone {
			session.CollectionId = ""
			snap.Ctx.HasCollection = false
			if session.GetSlots().GetPoolRule() != mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
				session.Slots.PoolRule = mtgv1.PoolRule_POOL_RULE_ANY_CARD
			}
		}
		return session, snap, version, owned, gone, nil
	}
	owned, err := s.ownedCounts(ctx, uid, msg.GetCollectionId())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, questions.Snapshot{}, 0, nil, false, connect.NewError(connect.CodeNotFound,
				fmt.Errorf("collection %q: %w", msg.GetCollectionId(), err))
		}
		return nil, questions.Snapshot{}, 0, nil, false, connect.NewError(connect.CodeInternal, err)
	}
	id := s.store.NewID(uid)
	session := &mtgv1.Session{
		Id:           id,
		CollectionId: msg.GetCollectionId(),
		Status:       mtgv1.SessionStatus_SESSION_STATUS_ASKING,
	}
	// The chat screen knows the card pool already: the reader named a
	// collection and said whether the deck may reach past it. The agent
	// takes that as the answer and asks nothing (D-359).
	if rule := msg.GetPoolRule(); rule != mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		session.Slots = &mtgv1.Slots{PoolRule: rule}
	}
	// A user with no collection never gets the card-pool question (D-37).
	snap := questions.Snapshot{Version: questions.SnapshotVersion}
	snap.Ctx.HasCollection = msg.GetCollectionId() != ""
	return session, snap, 0, owned, false, nil
}

// ownedCounts reads the owned counts of one collection. It answers nil
// with no error when the session has no collection or no source is
// wired.
func (s *Server) ownedCounts(ctx context.Context, uid, collectionID string) (map[string]int32, error) {
	if s.collections == nil || collectionID == "" {
		return nil, nil
	}
	return s.collections.OracleCounts(ctx, uid, collectionID)
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
// not answer them: they come from the stored session and from the
// candidate hints. Each one gates a catalog row.
//
// The hint facts read the slots, and the agent reads them again after
// the classify call fills the slots of this turn. This call covers the
// turn's first plan with the stored slots.
func (s *Server) facts(session *mtgv1.Session, st *questions.State, hints *questions.CandidateHints) {
	// A deck exists, so the user may ask for a change to it (D-283).
	st.Ctx.AfterBuild = len(session.GetDeckIds()) > 0
	if hints == nil {
		return
	}
	questions.RefreshFacts(st, hints)
}

// hints answers the placeholder values from the card index. It returns
// nil when no index is loaded, and every clause that needs a value is
// then dropped.
func (s *Server) hints(_ *mtgv1.Session, st *questions.State, owned map[string]int32) *questions.CandidateHints {
	if s.index == nil || s.builder == nil {
		return nil
	}
	idx := s.index.Current()
	if idx == nil {
		return nil
	}
	return &questions.CandidateHints{
		Index:   idx,
		Builder: s.builder,
		Format:  st.Slots.GetFormat().GetId(),
		Colors:  st.Slots.GetColors(),
		Pool:    st.Slots.GetPoolRule(),
		Owned:   owned,
		Log:     s.log,
	}
}

// declineNegatives closes every key whose question the user answered
// with a bare negative (D-352). It returns the keys it closed.
func declineNegatives(st *questions.State, answers []*mtgv1.Answer, session *mtgv1.Session) []string {
	if len(answers) == 0 {
		return nil
	}
	asked := map[string]*mtgv1.Question{}
	for _, turn := range session.GetTurns() {
		for _, q := range turn.GetQuestions() {
			asked[q.GetId()] = q
		}
	}
	var closed []string
	for _, a := range answers {
		q := asked[a.GetQuestionId()]
		if q == nil {
			continue
		}
		// A declined answer says outright that the user named no value
		// (D-353). It needs no reading of the words.
		if a.GetDeclined() {
			if key, ok := st.Decline(q.GetId()); ok {
				closed = append(closed, key)
			}
			continue
		}
		// An option index names a value, never a negative.
		if a.OptionIndex != nil {
			continue
		}
		if key, ok := st.DeclineNegative(q.GetId(), q.GetText(), a.GetText()); ok {
			closed = append(closed, key)
		}
	}
	return closed
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
			// An unset option_index is free text. The field has explicit
			// presence, so option 0 and "no option" are distinct.
			if i := int(a.GetOptionIndex()); a.OptionIndex != nil && i >= 0 && i < len(q.GetOptions()) {
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
		// The question goes on its own quoted line, so the classifier
		// sees it and the word rules skip it (questions.UserWords).
		lines = append(lines, questions.QuotedQuestionPrefix+q.GetText(), questions.AnswerPrefix+text)
	}
	if message = strings.TrimSpace(message); message != "" {
		lines = append(lines, message)
	}
	return strings.Join(lines, "\n")
}

// addUsage sums one turn's report into the session total. A turn whose
// provider reported nothing leaves priced false, which is not the same
// fact as a zero cost. A turn with no calls at all, such as a frozen
// session, reports no cost either, and it must not flip the flag.
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

// failure maps a model error onto the code the UI acts on. A schema
// miss is sampled, so a retry can succeed. A terminal fault is either
// the operator's (auth, config) or the user's.
func failure(err error) *mtgv1.AgentError {
	class := llm.ClassOf(err)
	out := &mtgv1.AgentError{
		Code:      "llm_" + class.String(),
		Message:   "The model call failed. This one does not succeed on a retry.",
		Retryable: class == llm.ClassTransient || class == llm.ClassBudget || class == llm.ClassSchema,
	}
	switch class {
	case llm.ClassTransient, llm.ClassBudget:
		out.Message = "The model call failed. Send the message again."
	case llm.ClassRefusal:
		out.Message = "The model declined that request. Say it in other words."
	case llm.ClassSchema:
		out.Message = "The model answered in the wrong shape. Send the message again, a retry usually succeeds."
	case llm.ClassTerminal:
		if operatorFault(err) {
			out.Message = "The model provider refused the setup. The operator must fix the provider setup."
		}
	}
	return out
}

// operatorFault reports a terminal error the user cannot fix: a refused
// key, a forbidden model, or a role with no provider wired.
func operatorFault(err error) bool {
	var e *llm.Error
	if !errors.As(err, &e) {
		return false
	}
	if e.Status == 401 || e.Status == 403 || e.Status == 404 {
		return true
	}
	if e.Err == nil {
		return false
	}
	msg := e.Err.Error()
	return strings.Contains(msg, "no provider") || strings.Contains(msg, "unknown role") ||
		strings.Contains(msg, "fixture") || strings.Contains(msg, "no fixture")
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
	if errors.Is(err, sessions.ErrTooLarge) {
		// The document limit is a hard stop. The user hears it as a clear
		// message, not as a session that fails every later turn.
		return connect.NewError(connect.CodeResourceExhausted, fmt.Errorf("%w: %w", errSessionBig, err))
	}
	return connect.NewError(connect.CodeInternal, err)
}

// slotsChanged compares the deck settings of two slot sets. The fill
// states are not settings: a turn that marks a question asked changes
// no deck. A nil side counts as empty.
func slotsChanged(before, after *mtgv1.Slots) bool {
	if before == nil {
		before = &mtgv1.Slots{}
	}
	if after == nil {
		after = &mtgv1.Slots{}
	}
	a := proto.Clone(before).(*mtgv1.Slots)
	b := proto.Clone(after).(*mtgv1.Slots)
	a.SlotStates, b.SlotStates = nil, nil
	return !proto.Equal(a, b)
}

// pairSeparator joins the two names of a commander pair in one option.
const pairSeparator = " + "

// cardOptions fills Question.option_oracle_ids for every option that is
// an exact card name. A question with no card option keeps the field
// empty, so a client can tell the two apart.
//
// A commander pair reads as "A + B" in one option, and neither half is
// the whole option, so a lookup of the option text finds nothing and the
// tile showed no card at all (D-361). Each half is resolved on its own:
// the first into option_oracle_ids, the second into the partner list.
func cardOptions(qs []*mtgv1.Question, idx *cards.Index) {
	if idx == nil {
		return
	}
	for _, q := range qs {
		ids := make([]string, len(q.GetOptions()))
		partners := make([]string, len(q.GetOptions()))
		found, anyPartner := false, false
		for i, opt := range q.GetOptions() {
			if c, ok := idx.ByName(opt); ok {
				ids[i] = c.GetOracleId()
				found = true
				continue
			}
			first, second, ok := strings.Cut(opt, pairSeparator)
			if !ok {
				continue
			}
			a, aok := idx.ByName(strings.TrimSpace(first))
			b, bok := idx.ByName(strings.TrimSpace(second))
			if !aok || !bok {
				continue
			}
			ids[i], partners[i] = a.GetOracleId(), b.GetOracleId()
			found, anyPartner = true, true
		}
		if found {
			q.OptionOracleIds = ids
		}
		if anyPartner {
			q.OptionPartnerOracleIds = partners
		}
	}
}
