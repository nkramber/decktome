package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/revise"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// The build runs when every slot is answered. The question workflow
// decided what to build, and nothing here asks the user anything: a slot
// that is still open never reaches this file.

// ErrThinCommanderPool says a Commander session delegated the commander
// and the library holds none for the theme. A build without a commander
// must fail the engine, so none runs, and the user reads why (D-232).
var ErrThinCommanderPool = errors.New("your library holds no commander for this theme, so no deck was built: name a commander, or allow cards you do not own")

// ErrThinSet says the sets the reader named hold too few cards in the
// deck's colors to build a legal deck (D-380). The build stops before it
// spends a model call, and the reader reads the count and the sets. A
// deck of another set is never the answer.
type ErrThinSet struct {
	// Sets are the set names, as a reader wrote them.
	Sets []string
	// Have is the distinct nonbasic count the sets offer in the deck's
	// colors, and Want is the floor for the format.
	Have, Want int
}

func (e *ErrThinSet) Error() string {
	return fmt.Sprintf("%s hold %d cards in these colors, and a deck of this format needs about %d: "+
		"name another set beside them, drop the set limit, or choose other colors",
		strings.Join(e.Sets, " and "), e.Have, e.Want)
}

// The wiring a build needs, each named so a log says which one is
// missing.
var (
	errNoIndexSource      = errors.New("build: no card index source is wired")
	errNoIndexLoaded      = errors.New("build: no card index is loaded")
	errNoCandidateBuilder = errors.New("build: no candidate builder is wired")
	errNoDeckStore        = errors.New("build: no deck store is wired")
)

// deckIDRetries bounds the retries of the deck id write after a version
// conflict (D-303).
const deckIDRetries = 3

// buildDeck writes the deck for a ready session. It returns nil when the
// server has no generator wired, so a deployment without one reports
// that every slot is filled and says so.
func (s *Server) buildDeck(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State, owned map[string]int32, acc *llm.Accumulator) (*generate.Result, error) {
	return s.buildDeckFrom(ctx, uid, session, st, owned, acc, nil)
}

// buildDeckFrom is buildDeck with an optional revision brief. The base
// deck's cards join the pool so the model can keep them, the brief's
// removed cards and the cards over its cap leave the pool, and the
// request carries the brief (D-283). owned is the collection's count per
// Oracle id, read once per turn.
func (s *Server) buildDeckFrom(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State, owned map[string]int32, acc *llm.Accumulator, rev *generate.Revision) (*generate.Result, error) {
	if s.decks == nil {
		return nil, nil
	}
	if s.index == nil {
		return nil, errNoIndexSource
	}
	if s.builder == nil {
		return nil, errNoCandidateBuilder
	}
	idx := s.index.Current()
	if idx == nil {
		return nil, errNoIndexLoaded
	}
	if owned == nil {
		owned = map[string]int32{}
	}
	slots := session.GetSlots()
	format := slots.GetFormat().GetId()
	setCodes := slots.GetSetCodes()

	// The set floor runs before the commander pool, so a family too thin
	// to build ends the turn with a reason and spends no model call
	// (D-380). It runs on the slot colors here and again on the
	// commander's identity below, because the commander narrows them.
	if err := s.checkSetFloor(ctx, idx, format, slots.GetColors(), setCodes); err != nil {
		return nil, err
	}

	var commanders []*mtgv1.Card
	var commanderIDs []string
	for _, name := range st.CommanderNames {
		c, ok := idx.ByName(name)
		if !ok {
			// The commander came from the card index, so a miss here is a
			// snapshot change and not a user error. The build goes on
			// without it, and the engine reports the missing commander.
			s.log.WarnContext(ctx, "the named commander left the snapshot", "card", name)
			continue
		}
		commanders = append(commanders, c)
		commanderIDs = append(commanderIDs, c.GetOracleId())
	}

	// A user who says "you pick" delegates the commander, and D-147 and
	// D-208 skip the slot for the generator. The generator must then pick
	// one, or the engine refuses the deck for no commander (D-232).
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER && len(commanderIDs) == 0 {
		pool, err := s.builder.CommanderPool(idx, candidates.Request{
			Format:   format,
			Theme:    slots.GetTheme(),
			Colors:   slots.GetColors(),
			PoolRule: slots.GetPoolRule(),
			Owned:    owned,
			Bracket:  slots.GetPower().GetBracket(),
			SetCodes: setCodes,
		})
		switch {
		case err != nil:
			s.log.WarnContext(ctx, "the commander pool failed", "err", err)
		case len(pool) == 0:
			// The library holds no commander for this theme. A Commander
			// deck with no commander fails the engine, so no model call is
			// spent on one, and the turn says why (D-232).
			s.log.WarnContext(ctx, "the library holds no commander for the theme, so no deck is built",
				"session", session.GetId(), "theme", slots.GetTheme())
			return nil, ErrThinCommanderPool
		default:
			c := pool[0]
			commanders = append(commanders, c.Card)
			commanderIDs = append(commanderIDs, c.Card.GetOracleId())
			if c.Partner != nil {
				commanders = append(commanders, c.Partner)
				commanderIDs = append(commanderIDs, c.Partner.GetOracleId())
			}
			s.log.InfoContext(ctx, "the generator picked the commander the user delegated",
				"session", session.GetId(), "card", c.Card.GetName())
		}
	}

	// The shortlist and the basic lands follow the commander's color
	// identity, and the slot colors only when no commander is chosen. The
	// slot can hold every color, and the engine refuses each card outside
	// the identity (D-289).
	colors := deckColors(format, slots.GetColors(), commanders)
	// The commander settles the colors, so the floor runs once more on
	// the identity the deck will actually have (D-380).
	if err := s.checkSetFloor(ctx, idx, format, colors, setCodes); err != nil {
		return nil, err
	}
	// A set family short of mana cards takes them from the whole
	// database, up to the role target and no further. The reader allowed
	// it in the set_outside_mana row, and every such card is marked
	// (D-382).
	var outsideRoles map[mtgv1.CardRole]int
	if len(setCodes) > 0 && slots.GetSlotStates()[questions.SlotSetOutsideMana] == mtgv1.SlotState_SLOT_STATE_FILLED {
		outsideRoles = manaRoles(generate.TargetsFor(format, slots.GetPower()))
	}
	list, err := s.builder.Build(idx, candidates.Request{
		Format:             format,
		Colors:             colors,
		Theme:              slots.GetTheme(),
		CommanderOracleIDs: commanderIDs,
		PoolRule:           slots.GetPoolRule(),
		Owned:              owned,
		Bracket:            slots.GetPower().GetBracket(),
		SetCodes:           setCodes,
		OutsideRoles:       outsideRoles,
	})
	if err != nil {
		return nil, fmt.Errorf("build: candidates: %w", err)
	}
	if len(setCodes) > 0 {
		s.log.InfoContext(ctx, "the deck is limited to the sets the reader named",
			"session", session.GetId(), "sets", strings.Join(setCodes, ","),
			"in_set", list.Stats.InSet, "outside", list.Stats.Outside)
	}
	// Upgrades reach the model only when the session may buy cards. A
	// deck must not name a card the user can neither own nor buy (D-37).
	buyList := st.Ctx.BuyList || slots.GetPoolRule() == mtgv1.PoolRule_POOL_RULE_ANY_CARD
	// The shortlist leaves basic lands out on purpose, and every deck
	// needs them (D-225).
	always := append([]*mtgv1.Card(nil), commanders...)
	always = append(always, generate.BasicLands(idx.ByName, colors)...)
	// A card the user said to keep must be nameable, or the deck can not
	// hold it. The classifier names these (D-70, D-242).
	var lockedIDs []string
	for _, name := range st.LockedCards() {
		c, ok := idx.ByName(name)
		if !ok {
			s.log.WarnContext(ctx, "a locked card left the snapshot", "card", name)
			continue
		}
		always = append(always, c)
		lockedIDs = append(lockedIDs, c.GetOracleId())
	}
	// The precon the user asked to upgrade. Its cards are what the share
	// rule of D-218 measures, and every one must be nameable, so they join
	// the pool before it is built (D-248).
	var preconName string
	var preconIDs []string
	var preconLands int
	if set := s.preconSet(); set != nil && st.Ctx.Precon {
		if p, ok := set.Find(st.Ctx.Words); ok && p.Unresolved > 0 {
			// A list the card index could not fully answer gives a wrong
			// share, so the rule does not run on it.
			s.log.WarnContext(ctx, "the precon has unresolved rows, so the share rule is off",
				"session", session.GetId(), "precon", p.Slug, "unresolved", p.Unresolved)
		} else if ok {
			preconName, preconIDs, preconLands = p.Name, p.OracleIDs, p.Lands
			// The deck must keep a share of these, so the model must be
			// able to name them (D-247).
			for _, id := range preconIDs {
				if c, ok := idx.ByOracleID(id); ok {
					always = append(always, c)
				}
			}
			s.log.InfoContext(ctx, "the user asked to upgrade a precon",
				"session", session.GetId(), "precon", p.Name, "cards", len(p.OracleIDs))
		} else {
			s.log.InfoContext(ctx, "the user named a precon the app does not hold",
				"session", session.GetId())
		}
	}
	// A revision keeps every card the brief does not touch, so the base
	// deck's cards must be nameable (D-283).
	if rev != nil {
		for _, dc := range rev.Base {
			if c, ok := idx.ByOracleID(dc.GetOracleId()); ok {
				always = append(always, c)
			}
		}
	}
	// The always cards read their owned count from the collection, so a
	// precon card the user holds is never charged as a purchase.
	pool := generate.FromListOwned(list, always, owned, buyList)
	if rev != nil {
		// The mana value cap never drops a commander or a locked card:
		// the engine blocks a deck without them, and the model can not
		// put back what it can not name (D-242).
		rev.Exempt = append(append([]string(nil), commanderIDs...), lockedIDs...)
		pool = pool.Filter(func(c *mtgv1.Card) bool { return generate.AllowedByRevision(rev, c) })
	}

	// The deck carries its own id, so the store reserves one first
	// (D-245).
	var deckID string
	if s.deckStore != nil {
		deckID = s.deckStore.NewID(uid)
	}

	res, err := s.decks.Build(ctx, generate.Request{
		Precon:          preconName,
		PreconOracleIDs: preconIDs,
		PreconLands:     preconLands,
		DeckID:          deckID,
		Name:            deckName(slots),
		Now:             s.now,
		BudgetUSD:       slots.GetBudgetUsd(),
		BudgetWholeDeck: slots.GetBudgetScope() == mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK,
		SessionID:       session.GetId(),
		Format:          format,
		HouseRules:      slots.GetHouseRules(),
		Power:           slots.GetPower(),
		Plan:            plan(session, slots),
		Pool:            pool,
		Commanders:      commanderIDs,
		Locked:          lockedIDs,
		PoolRule:        slots.GetPoolRule(),
		SetCodes:        setCodes,
		OracleCounts:    owned,
		Roles:           generate.Roles(list),
		Targets:         generate.TargetsFor(format, slots.GetPower()),
		Limits:          generate.LimitsFor(format),
		LegalityAsOf:    idx.AsOf.Format("2006-01-02"),
		Revision:        rev,
	}, acc)
	if err != nil {
		return nil, err
	}
	s.markOwnedPrintings(ctx, uid, session, idx, res.Deck)
	return res, nil
}

// markOwnedPrintings sets DeckCard.owned_printing to the priciest printing
// the collection holds of each owned card (D-299). A missing collection
// or a printing the snapshot dropped leaves the field empty.
func (s *Server) markOwnedPrintings(ctx context.Context, uid string, session *mtgv1.Session, idx *cards.Index, deck *mtgv1.Deck) {
	if s.collections == nil || session.GetCollectionId() == "" || deck == nil {
		return
	}
	owned, err := s.collections.OwnedPrintings(ctx, uid, session.GetCollectionId())
	if err != nil {
		s.log.WarnContext(ctx, "owned printings unavailable", "collection", session.GetCollectionId(), "err", err)
		return
	}
	for _, list := range [][]*mtgv1.DeckCard{deck.GetCards(), deck.GetSideboard()} {
		for _, dc := range list {
			if dc.GetOwnedCount() == 0 {
				continue
			}
			var best *mtgv1.Printing
			for _, id := range owned[dc.GetOracleId()] {
				p, ok := idx.Printing(id)
				if !ok || p.GetImageUris() == nil {
					continue
				}
				if best == nil || p.GetPriceUsd() > best.GetPriceUsd() {
					best = p
				}
			}
			dc.OwnedPrinting = best
		}
	}
}

// deckColors is the color set the basic lands follow. In Commander it is
// the union of every commander's color identity, so a partner pair gets
// both halves. Otherwise it is the chosen colors, and a 60-card session
// that chose none gets every basic.
func deckColors(format mtgv1.FormatId, chosen []mtgv1.Color, commanders []*mtgv1.Card) []mtgv1.Color {
	if len(commanders) > 0 {
		seen := map[mtgv1.Color]bool{}
		for _, c := range commanders {
			for _, col := range c.GetColorIdentity() {
				seen[col] = true
			}
		}
		var out []mtgv1.Color
		for _, col := range generate.AllColors {
			if seen[col] {
				out = append(out, col)
			}
		}
		return out
	}
	if len(chosen) == 0 && format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		return generate.AllColors
	}
	return chosen
}

// plan is the deck the user asked for, in the user's own words. The
// first message carries the request, and the slots carry what the
// questions settled.
func plan(session *mtgv1.Session, slots *mtgv1.Slots) string {
	var b strings.Builder
	if turns := session.GetTurns(); len(turns) > 0 {
		b.WriteString(strings.TrimSpace(turns[0].GetUserMessage()))
	}
	if t := slots.GetTheme(); t != "" {
		fmt.Fprintf(&b, "\nTheme: %s", t)
	}
	if p := slots.GetPower(); p != nil {
		if br := p.GetBracket(); br > 0 {
			fmt.Fprintf(&b, "\nCommander bracket: %d", br)
		} else if step := p.GetSixtyStep(); step != mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED {
			fmt.Fprintf(&b, "\nPower: %s", strings.ToLower(strings.TrimPrefix(step.String(), "SIXTY_STEP_")))
		}
	}
	if n := slots.GetBudgetUsd(); n > 0 {
		fmt.Fprintf(&b, "\nBudget: %.0f dollars", n)
	}
	return b.String()
}

// sendOrLog sends one event after a paid call. A client that left
// mid-build can not receive it, and the build must still end and store
// its deck, so a failed send is logged and the turn goes on (D-303).
func (s *Server) sendOrLog(ctx context.Context, stream *connect.ServerStream[mtgv1.ChatResponse], session *mtgv1.Session, ev *mtgv1.ChatResponse) {
	if err := stream.Send(ev); err != nil {
		s.log.WarnContext(ctx, "the client left before the event was sent", "session", session.GetId(), "err", err)
	}
}

// sendDeck builds the deck and streams it. A build failure must not lose
// the turn: the questions are already stored and already sent, so the
// user reads a status line and can ask again.
func (s *Server) sendDeck(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State,
	snap questions.Snapshot, version int64, owned map[string]int32, acc *llm.Accumulator,
	stream *connect.ServerStream[mtgv1.ChatResponse]) error {
	if s.decks == nil {
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: "every slot is filled, and no generator is wired"},
		})
	}
	if err := stream.Send(&mtgv1.ChatResponse{
		Event: &mtgv1.ChatResponse_Status{Status: "building the deck"},
	}); err != nil {
		return err
	}
	// The build is bounded, so a slow provider ends the turn with an
	// error instead of a stream that does not end (D-235). It runs
	// detached from the client, so a disconnect after the paid call
	// still stores the deck (D-303).
	bctx, cancel := detached(ctx, s.buildDeadline())
	defer cancel()
	res, err := s.buildDeck(bctx, uid, session, st, owned, acc)
	if err != nil {
		s.log.ErrorContext(ctx, "the build failed", "session", session.GetId(), "err", err)
		msg := "the deck build failed, please ask again"
		var thinSet *ErrThinSet
		switch {
		case errors.Is(err, ErrThinCommanderPool):
			msg = err.Error()
		// A set family too thin for a legal deck ends the turn with the
		// reason, and never with a deck of another set (D-380).
		case errors.As(err, &thinSet):
			msg = thinSet.Error()
		case errors.Is(err, errNoIndexSource), errors.Is(err, errNoIndexLoaded), errors.Is(err, errNoCandidateBuilder):
			msg = "the deck can not be built here: " + strings.TrimPrefix(err.Error(), "build: ")
		case errors.Is(bctx.Err(), context.DeadlineExceeded):
			msg = "the deck build ran past its time limit, please ask again"
		}
		s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: msg}})
		return nil
	}
	if res == nil {
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: "every slot is filled, and no generator is wired"},
		})
	}
	// The deck is kept before it is sent, so a user who reads it can ask
	// for it again (D-245).
	s.storeDeck(ctx, uid, session, snap, version, res.Deck)
	// A name the model wrote twice and the shortlist did not hold reaches
	// the user as prose, because the card is absent from the deck.
	for _, n := range res.Notes {
		s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_TextDelta{TextDelta: n}})
	}
	s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Deck{Deck: res.Deck}})
	return nil
}

// deckName is what the user sees the deck called. The theme and the
// format are what a person would name it by.
func deckName(slots *mtgv1.Slots) string {
	theme := strings.TrimSpace(slots.GetTheme())
	format := generate.FormatWord(slots.GetFormat().GetId())
	if theme == "" {
		return format + " deck"
	}
	return theme + " " + format
}

// storeDeck keeps the deck and records its id on the session. Both
// writes run detached from the client (D-303). A store failure must not
// lose the deck the user is already reading, so it warns and the turn
// goes on (D-245).
func (s *Server) storeDeck(ctx context.Context, uid string, session *mtgv1.Session,
	snap questions.Snapshot, version int64, d *mtgv1.Deck) {
	if s.deckStore == nil || d.GetId() == "" {
		return
	}
	sctx, cancel := detached(ctx, storeLimit)
	defer cancel()
	if err := s.deckStore.Put(sctx, uid, d); err != nil {
		s.log.ErrorContext(ctx, "the deck was not stored", "session", session.GetId(), "deck", d.GetId(), "err", err)
		return
	}
	// The session was written before the build, so the deck id needs its
	// own write. Without it AfterBuild stays false for the next turn
	// (D-245).
	session.DeckIds = append(session.GetDeckIds(), d.GetId())
	err := s.store.Put(sctx, uid, session, snap, version)
	// A version conflict means another write landed since the turn was
	// stored. The deck exists, so the id is appended to the current
	// session instead of lost (D-303).
	for try := 0; errors.Is(err, sessions.ErrConflict) && try < deckIDRetries; try++ {
		var current *mtgv1.Session
		current, _, version, err = s.store.GetState(sctx, uid, session.GetId())
		if err != nil {
			break
		}
		current.DeckIds = append(current.GetDeckIds(), d.GetId())
		current.Status = mtgv1.SessionStatus_SESSION_STATUS_BUILT
		current.UpdatedAt = session.GetUpdatedAt()
		if n := len(current.GetTurns()); n > 0 && n == len(session.GetTurns()) {
			current.Turns[n-1] = session.GetTurns()[n-1]
		}
		err = s.store.Put(sctx, uid, current, snap, version)
		if err == nil {
			session.DeckIds = current.GetDeckIds()
		}
	}
	if err != nil {
		s.log.ErrorContext(ctx, "the deck id was not recorded on the session",
			"session", session.GetId(), "deck", d.GetId(), "err", err)
	}
}

// sendRevision runs the turn after a build: the revise call reads the
// message and the deck the user read, and the outcome is a question, a
// decline with a reason, or a revised deck. Silence is never an outcome
// (D-283, D-284). The reply lands in Turn.agent_message and streams as
// text_delta.
func (s *Server) sendRevision(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State,
	snap questions.Snapshot, version int64, owned map[string]int32, acc *llm.Accumulator, message string, turn *mtgv1.Turn,
	stream *connect.ServerStream[mtgv1.ChatResponse]) error {
	if s.decks == nil {
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: "every slot is filled, and no generator is wired"},
		})
	}
	if s.deckStore == nil {
		s.log.ErrorContext(ctx, "the deck can not be revised", "session", session.GetId(), "err", errNoDeckStore)
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: "the deck can not be revised here, because no deck store is wired"},
		})
	}
	ids := session.GetDeckIds()
	base, err := s.deckStore.Get(ctx, uid, ids[len(ids)-1])
	if err != nil {
		s.log.ErrorContext(ctx, "the base deck could not be read", "session", session.GetId(), "deck", ids[len(ids)-1], "err", err)
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: "the last deck could not be read, please ask again"},
		})
	}
	if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: "reading your request"}}); err != nil {
		return err
	}
	var cards revise.CardSource
	if s.index != nil {
		if idx := s.index.Current(); idx != nil {
			cards = idx
		}
	}
	slots := session.GetSlots()
	brief, err := revise.Call(ctx, s.client, revise.Input{
		SessionID: session.GetId(),
		Message:   message,
		Prior:     priorUserMessage(session),
		Deck:      base,
		Format:    generate.FormatWord(slots.GetFormat().GetId()),
		Power:     powerWord(slots.GetPower()),
		Cards:     cards,
	}, acc)
	if err != nil {
		s.log.ErrorContext(ctx, "the revise call failed", "session", session.GetId(), "err", err)
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: "your request could not be read, please ask again"},
		})
	}
	// The revise call is paid, so the write that records its reply runs
	// detached from the client (D-303).
	record := func(text string) {
		turn.AgentMessage = text
		session.UpdatedAt = timestamppb.New(s.now())
		sctx, cancel := detached(ctx, storeLimit)
		defer cancel()
		if err := s.store.Put(sctx, uid, session, snap, version); err != nil {
			s.log.ErrorContext(ctx, "the reply was not recorded on the session", "session", session.GetId(), "err", err)
		}
	}
	if brief.Question != "" {
		q := &mtgv1.Question{
			Id:       fmt.Sprintf("revise-%d", len(session.GetTurns())),
			Slot:     "revision",
			Text:     brief.Question,
			Invented: true,
		}
		turn.Questions = append(turn.Questions, q)
		// A question was sent, so the session is asking again.
		session.Status = mtgv1.SessionStatus_SESSION_STATUS_ASKING
		record(q.GetText())
		s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Question{Question: q}})
		return nil
	}
	if !brief.Acts() {
		note := revise.DeclineNote(brief)
		record(note)
		s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_TextDelta{TextDelta: note}})
		return nil
	}
	// A card the brief removes is not a card to keep, whatever the
	// classify call read from the same message. The state drops the lock
	// before the build, and the stored state follows (D-301).
	for _, name := range brief.Remove {
		st.Unlock(name)
	}
	snap = st.Snapshot()
	s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: "revising the deck"}})
	bctx, cancel := detached(ctx, s.buildDeadline())
	defer cancel()
	rev := &generate.Revision{
		BaseDeckID:   base.GetId(),
		Base:         base.GetCards(),
		Instructions: brief.Changes,
		Remove:       brief.Remove,
		Keep:         brief.Keep,
		MaxManaValue: brief.MaxManaValue,
	}
	res, err := s.buildDeckFrom(bctx, uid, session, st, owned, acc, rev)
	if err != nil || res == nil {
		s.log.ErrorContext(ctx, "the revision failed", "session", session.GetId(), "err", err)
		msg := "the revision failed, please ask again"
		var thinSet *ErrThinSet
		switch {
		case errors.As(err, &thinSet):
			msg = thinSet.Error()
		case errors.Is(bctx.Err(), context.DeadlineExceeded):
			msg = "the revision ran past its time limit, please ask again"
		}
		s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: msg}})
		return nil
	}
	if base.GetName() != "" {
		res.Deck.Name = base.GetName()
	}
	note := revise.Note(brief, revise.DiffDecks(base, res.Deck))
	res.Deck.RevisionNote = note
	res.Deck.RevisedFromDeckId = base.GetId()
	// storeDeck writes the session with the deck id, and the turn holds
	// the reply, so one write records both.
	turn.AgentMessage = note
	s.storeDeck(ctx, uid, session, snap, version, res.Deck)
	s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_TextDelta{TextDelta: note}})
	for _, n := range res.Notes {
		s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_TextDelta{TextDelta: n}})
	}
	s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Deck{Deck: res.Deck}})
	return nil
}

// priorUserMessage is the user's message of the turn before this one,
// so a revise call that follows its own question reads the request too.
func priorUserMessage(session *mtgv1.Session) string {
	turns := session.GetTurns()
	if len(turns) < 2 {
		return ""
	}
	return questions.UserWords(turns[len(turns)-2].GetUserMessage())
}

// powerWord is the power as the revise prompt names it.
func powerWord(p *mtgv1.PowerLevel) string {
	if br := p.GetBracket(); br > 0 {
		return fmt.Sprintf("bracket %d", br)
	}
	if step := p.GetSixtyStep(); step != mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED {
		return strings.ToLower(strings.TrimPrefix(step.String(), "SIXTY_STEP_"))
	}
	return ""
}

// manaRoles are the roles a set-limited deck may fill from outside the
// named sets: the ramp and the nonbasic lands (D-382). The counts come
// from the role targets of the format, so the fill reaches the wanted
// number and stops.
func manaRoles(targets map[string]int) map[mtgv1.CardRole]int {
	out := map[mtgv1.CardRole]int{}
	if n := targets["ramp"]; n > 0 {
		out[mtgv1.CardRole_CARD_ROLE_RAMP] = n
	}
	if n := targets["land"]; n > 0 {
		out[mtgv1.CardRole_CARD_ROLE_LAND] = n
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

// checkSetFloor refuses a build whose sets can not fill a legal deck in
// these colors (D-380). It counts and scores nothing, so it costs a
// fraction of a shortlist build.
func (s *Server) checkSetFloor(ctx context.Context, idx *cards.Index,
	format mtgv1.FormatId, colors []mtgv1.Color, setCodes []string,
) error {
	if len(setCodes) == 0 {
		return nil
	}
	have := candidates.CountInSets(idx, candidates.Request{
		Format: format, Colors: colors, SetCodes: setCodes,
	})
	want := candidates.SetFloor(format)
	if have >= want {
		return nil
	}
	names := idx.Sets().Names(setCodes)
	s.log.WarnContext(ctx, "the sets hold too few cards to build a deck in these colors",
		"sets", strings.Join(setCodes, ","), "have", have, "want", want)
	return &ErrThinSet{Sets: setNames(names), Have: have, Want: want}
}

// setNames reads a set-name list as a reader would hear it. A family
// reads as its base set, because "The Hobbit and The Hobbit Eternal" is
// how a reader names one product, and a list of six promo sets is not.
func setNames(names []string) []string {
	if len(names) <= 2 {
		return names
	}
	return append(append([]string(nil), names[:2]...), "and the other sets you named")
}
