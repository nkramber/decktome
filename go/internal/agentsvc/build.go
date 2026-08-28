package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

// The build runs when every slot is answered (roadmap PR-8). The
// question workflow decided what to build, and nothing here asks the
// user anything: a slot that is still open never reaches this file.

// ErrThinCommanderPool says a Commander session delegated the commander
// and the library holds none for the theme. A build without a commander
// must fail the engine, so none runs, and the user reads why (D-232,
// G-11 of the 2026-08-28 audit).
var ErrThinCommanderPool = errors.New("your library holds no commander for this theme, so no deck was built: name a commander, or allow cards you do not own")

// buildDeck writes the deck for a ready session. It returns nil when the
// server has no generator wired, so a deployment without one keeps the
// PR-7 behavior and says so.
func (s *Server) buildDeck(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State, acc *llm.Accumulator) (*generate.Result, error) {
	if s.decks == nil || s.index == nil || s.builder == nil {
		return nil, nil
	}
	idx := s.index.Current()
	if idx == nil {
		return nil, fmt.Errorf("build: no card index is loaded")
	}
	slots := session.GetSlots()
	format := slots.GetFormat().GetId()

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

	owned := map[string]int32{}
	if s.collections != nil && session.GetCollectionId() != "" {
		if got, err := s.collections.OracleCounts(ctx, uid, session.GetCollectionId()); err == nil {
			owned = got
		} else {
			s.log.WarnContext(ctx, "owned counts unavailable for the build",
				"collection", session.GetCollectionId(), "err", err)
		}
	}

	// A user who says "you pick" delegates the commander, and D-147 and
	// D-208 skip the slot for the generator. The generator must then pick
	// one, or the engine refuses the deck for no commander. No golden
	// prompt covered this path until D-232.
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER && len(commanderIDs) == 0 {
		pool, err := s.builder.CommanderPool(idx, candidates.Request{
			Format:   format,
			Theme:    slots.GetTheme(),
			Colors:   slots.GetColors(),
			PoolRule: slots.GetPoolRule(),
			Owned:    owned,
			Bracket:  slots.GetPower().GetBracket(),
		})
		switch {
		case err != nil:
			s.log.WarnContext(ctx, "the commander pool failed", "err", err)
		case len(pool) == 0:
			// The library holds no commander for this theme. The weak-pool
			// row used to ask about this and could never reach the user
			// who needed it (D-232). A Commander deck with no commander
			// fails the engine, so no model call is spent on one, and the
			// turn says why (G-11).
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

	list, err := s.builder.Build(idx, candidates.Request{
		Format:             format,
		Colors:             slots.GetColors(),
		Theme:              slots.GetTheme(),
		CommanderOracleIDs: commanderIDs,
		PoolRule:           slots.GetPoolRule(),
		Owned:              owned,
		Bracket:            slots.GetPower().GetBracket(),
	})
	if err != nil {
		return nil, fmt.Errorf("build: candidates: %w", err)
	}
	// Upgrades reach the model only when the session may buy cards. A
	// deck must not name a card the user can neither own nor buy (D-37).
	buyList := st.Ctx.BuyList || slots.GetPoolRule() == mtgv1.PoolRule_POOL_RULE_ANY_CARD
	// The shortlist leaves basic lands out on purpose, and every deck
	// needs them (D-225).
	colors := deckColors(format, slots.GetColors(), commanders)
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
	// rule of D-218 measures, and every one must be nameable. This must
	// run before the pool is built: it did not, and the model was given
	// 27 of 93 cards and an instruction it could not meet (D-248).
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
	// The always cards read their owned count from the collection, so a
	// precon card the user holds is never charged as a purchase (G-3).
	pool := generate.FromListOwned(list, always, owned, buyList)

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
		OracleCounts:    owned,
		Roles:           generate.Roles(list),
		Targets:         generate.TargetsFor(format, slots.GetPower()),
		Limits:          generate.LimitsFor(format),
		LegalityAsOf:    idx.AsOf.Format("2006-01-02"),
	}, acc)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// deckColors is the color set the basic lands follow. In Commander it is
// the union of every commander's color identity, so a partner pair gets
// both halves (G-7 of the 2026-08-28 audit). Otherwise it is the chosen
// colors, and a 60-card session that chose none gets every basic.
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

// sendDeck builds the deck and streams it. A build failure must not lose
// the turn: the questions are already stored and already sent, so the
// user reads a status line and can ask again.
func (s *Server) sendDeck(ctx context.Context, uid string, session *mtgv1.Session, st *questions.State,
	snap questions.Snapshot, version int64, acc *llm.Accumulator,
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
	// error instead of holding the stream open (D-235).
	limit := s.buildLimit
	if limit <= 0 {
		limit = DefaultBuildLimit
	}
	bctx, cancel := context.WithTimeout(ctx, limit)
	defer cancel()
	res, err := s.buildDeck(bctx, uid, session, st, acc)
	if err != nil {
		s.log.ErrorContext(ctx, "the build failed", "session", session.GetId(), "err", err)
		msg := "the deck build failed, please ask again"
		switch {
		case errors.Is(err, ErrThinCommanderPool):
			msg = err.Error()
		case errors.Is(bctx.Err(), context.DeadlineExceeded):
			msg = "the deck build ran past its time limit, please ask again"
		}
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: msg},
		})
	}
	if res == nil {
		return stream.Send(&mtgv1.ChatResponse{
			Event: &mtgv1.ChatResponse_Status{Status: "every slot is filled, and no generator is wired"},
		})
	}
	// The deck is kept before it is sent, so a user who reads it can ask
	// for it again (D-245).
	s.storeDeck(ctx, uid, session, snap, version, res.Deck)
	// A name the model wrote twice and the shortlist never held reaches
	// the user as prose, because the card is absent from the deck.
	for _, n := range res.Notes {
		if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_TextDelta{TextDelta: n}}); err != nil {
			return err
		}
	}
	return stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Deck{Deck: res.Deck}})
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

// storeDeck keeps the deck and records its id on the session. A store
// failure must not lose the deck the user is already reading, so it warns
// and the turn goes on (D-245).
func (s *Server) storeDeck(ctx context.Context, uid string, session *mtgv1.Session,
	snap questions.Snapshot, version int64, d *mtgv1.Deck) {
	if s.deckStore == nil || d.GetId() == "" {
		return
	}
	if err := s.deckStore.Put(ctx, uid, d); err != nil {
		s.log.ErrorContext(ctx, "the deck was not stored", "session", session.GetId(), "deck", d.GetId(), "err", err)
		return
	}
	// The session was written before the build, so the deck id needs its
	// own write. Without it AfterBuild stays false and the variance row
	// is dead for the next turn (D-245).
	session.DeckIds = append(session.GetDeckIds(), d.GetId())
	if err := s.store.Put(ctx, uid, session, snap, version); err != nil {
		s.log.ErrorContext(ctx, "the deck id was not recorded on the session",
			"session", session.GetId(), "deck", d.GetId(), "err", err)
	}
}
