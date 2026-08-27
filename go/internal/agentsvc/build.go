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
	var thinPool bool
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
			// who needed it (D-232). The deck reports it instead.
			thinPool = true
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
	// needs them (D-225). The deck colors are the commander's identity in
	// Commander, and the chosen colors otherwise.
	colors := slots.GetColors()
	if len(commanders) > 0 {
		colors = commanders[0].GetColorIdentity()
	}
	always := append([]*mtgv1.Card(nil), commanders...)
	always = append(always, generate.BasicLands(idx.ByName, colors)...)
	pool := generate.FromList(list, always, buyList)

	res, err := s.decks.Build(ctx, generate.Request{
		ThinCommanderPool: thinPool,
		BudgetUSD:         slots.GetBudgetUsd(),
		SessionID:         session.GetId(),
		Format:            format,
		Power:             slots.GetPower(),
		Plan:              plan(session, slots),
		Pool:              pool,
		Commanders:        commanderIDs,
		PoolRule:          slots.GetPoolRule(),
		OracleCounts:      owned,
		Roles:             generate.Roles(list),
		Targets:           generate.TargetsFor(format, slots.GetPower()),
		Limits:            generate.LimitsFor(format),
		LegalityAsOf:      idx.AsOf.Format("2006-01-02"),
	}, acc)
	if err != nil {
		return nil, err
	}
	return res, nil
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
	acc *llm.Accumulator, stream *connect.ServerStream[mtgv1.ChatResponse]) error {
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
		if errors.Is(bctx.Err(), context.DeadlineExceeded) {
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
	// A name the model wrote twice and the shortlist never held reaches
	// the user as prose, because the card is absent from the deck.
	for _, n := range res.Notes {
		if err := stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_TextDelta{TextDelta: n}}); err != nil {
			return err
		}
	}
	return stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Deck{Deck: res.Deck}})
}
