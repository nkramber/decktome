package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/gzstore"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/questions"
	"github.com/nkramber/decktome/go/internal/revise"
)

// rerunText is the user line of a rerun turn. The web shows the same
// words, so a reload reads the thread it showed.
const rerunText = "Rerun this deck after the rule change."

// rerunPickMessage is the message the classify call reads when a ban hit
// the commander. It names no card, so the classify call can not fill the
// commander slot with the banned card again (D-1021).
const rerunPickMessage = "Suggest three new commanders for this deck."

// patchChange is the one change of the patch brief (D-1019).
const patchChange = "Replace each removed card with a legal card that fills the same role in the deck."

var (
	errBadRerunID     = fmt.Errorf("rerun_deck_id: %w", gzstore.ErrBadID)
	errRerunNoSession = errors.New("a rerun needs the session of its deck")
	errRerunInput     = errors.New("a rerun carries no message and no answer")
	errRerunOther     = errors.New("the deck is not a deck of this session")
	errNotStale       = errors.New("every card of this deck is legal, so it needs no rerun")
)

// checkRerun refuses a rerun request of the wrong shape, before any state
// of the turn changes.
func checkRerun(msg *mtgv1.ChatRequest) error {
	id := msg.GetRerunDeckId()
	switch {
	case id == "":
		return nil
	case !gzstore.ValidID(id):
		return connect.NewError(connect.CodeInvalidArgument, errBadRerunID)
	case msg.GetSessionId() == "":
		return connect.NewError(connect.CodeInvalidArgument, errRerunNoSession)
	case msg.GetMessage() != "" || len(msg.GetAnswers()) > 0:
		return connect.NewError(connect.CodeInvalidArgument, errRerunInput)
	}
	return nil
}

// rerunDeck reads the stale deck that a rerun names. The deck must be a
// deck of the session, and it must be stale now: the stale pass clears a
// deck after an unban.
func (s *Server) rerunDeck(ctx context.Context, uid string, session *mtgv1.Session, id string) (*mtgv1.Deck, error) {
	if s.deckStore == nil {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errNoDeckStore)
	}
	if !slices.Contains(session.GetDeckIds(), id) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errRerunOther)
	}
	d, err := s.deckStore.Get(ctx, uid, id)
	if errors.Is(err, decks.ErrNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if !d.GetStale() {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errNotStale)
	}
	return d, nil
}

// rerunNeedsCommander says the ban hit the commander, so the rerun asks
// for a new one with the pick row before it builds (D-1021).
func rerunNeedsCommander(d *mtgv1.Deck) bool {
	for _, id := range d.GetCommanderOracleIds() {
		if slices.Contains(d.GetStaleOracleIds(), id) {
			return true
		}
	}
	return false
}

// staleCardNames reads the names of the stale cards of the main list,
// the sideboard, and the command zone. An imported list keeps its
// commander and its companion as ids alone, so a stale id outside the
// lists takes its name from the card index. The revision and the locks
// match cards by name.
func staleCardNames(d *mtgv1.Deck, cards revise.CardSource) []string {
	var out []string
	named := map[string]bool{}
	for _, list := range [][]*mtgv1.DeckCard{d.GetCommanders(), d.GetCards(), d.GetSideboard()} {
		for _, dc := range list {
			if slices.Contains(d.GetStaleOracleIds(), dc.GetOracleId()) && !slices.Contains(out, dc.GetName()) {
				named[dc.GetOracleId()] = true
				out = append(out, dc.GetName())
			}
		}
	}
	if cards == nil {
		return out
	}
	for _, id := range append(slices.Clone(d.GetCommanderOracleIds()), d.GetCompanionOracleId()) {
		if id == "" || named[id] || !slices.Contains(d.GetStaleOracleIds(), id) {
			continue
		}
		if c, ok := cards.ByOracleID(id); ok && !slices.Contains(out, c.GetName()) {
			named[id] = true
			out = append(out, c.GetName())
		}
	}
	return out
}

// rerunCards is the card index of a rerun, or nil before the index loads.
func (s *Server) rerunCards() revise.CardSource {
	if s.index != nil {
		if idx := s.index.Current(); idx != nil {
			return idx
		}
	}
	return nil
}

// prepareRerun unlocks each stale card, so no kept card holds a ban in
// the next build. For a banned commander it clears the commander, and
// the turn then asks with the pick row (D-1021).
func prepareRerun(st *questions.State, d *mtgv1.Deck, cards revise.CardSource) {
	for _, name := range staleCardNames(d, cards) {
		st.Unlock(name)
	}
	if rerunNeedsCommander(d) {
		st.ClearCommander()
	}
}

// sendRerun runs the rerun of a stale deck whose commander is legal. It
// needs no classify call, because the deck names the rerun (D-1008). A
// patch is a revision of the stale deck with its stale cards in Remove
// (D-1019). A rebuild is a new build from the slots of the conversation
// (D-1020). The next snapshot keeps each banned card out of the pool.
func (s *Server) sendRerun(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State,
	version int64, owned map[string]int32, d *mtgv1.Deck, stream *connect.ServerStream[mtgv1.ChatResponse]) error {
	release, err := s.takeLease(ctx, uid, session.GetId())
	if err != nil {
		return err
	}
	defer release()
	key := buildKey(uid, session.GetId())
	s.building.Store(key, struct{}{})
	defer s.building.Delete(key)

	acc := llm.NewAccumulator(s.prices)
	usageBefore := cloneUsage(session.GetUsage())
	defer s.recordSpend(ctx, uid, session, usageBefore)

	cards := s.rerunCards()
	prepareRerun(st, d, cards)
	now := timestamppb.New(s.now())
	turn := &mtgv1.Turn{UserMessage: rerunText, At: now}
	session.Slots = st.Slots
	session.Turns = append(session.Turns, turn)
	session.Status = mtgv1.SessionStatus_SESSION_STATUS_BUILT
	session.UpdatedAt = now
	// The rerun is a paid build, so the write that starts it runs
	// detached from the client (D-303).
	sctx, cancel := detached(ctx, storeLimit)
	err = s.store.Put(sctx, uid, session, st.Snapshot(), version)
	cancel()
	if err != nil {
		return storeError(err)
	}
	if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: d.GetStaleReason()}}); err != nil {
		return err
	}
	// The rerun builds from the slots of this turn, and the next turn
	// compares its slots with them (D-1118).
	st.MarkBuilt()
	// version+1 is what the Put above stored (D-245).
	if d.GetRerunCase() == mtgv1.RerunCase_RERUN_CASE_PATCH {
		brief := &revise.Brief{Changes: []string{patchChange}, Remove: staleCardNames(d, cards)}
		turn.RevisionBrief = brief.JSON()
		err = s.reviseDeck(ctx, uid, session, st, version+1, owned, acc, usageBefore, d, brief, cards, turn, stream)
	} else {
		err = s.sendDeck(ctx, uid, session, st, st.Snapshot(), version+1, owned, acc, usageBefore, stream, nil)
	}
	if err != nil {
		return err
	}
	return stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Usage{Usage: session.GetUsage()}})
}
