package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/decklist"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/gzstore"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/questions"
	"github.com/nkramber/decktome/go/internal/rules"
	"github.com/nkramber/decktome/go/internal/users"
)

// Importer reads a deck that a user brought, as the build reads a deck it
// made (PR-70). generate.Builder is the one implementation.
type Importer interface {
	ReadImport(ctx context.Context, deck *mtgv1.Deck, owned map[string]int32, acc *llm.Accumulator)
}

// The builder of the app is the importer.
var _ Importer = (*generate.Builder)(nil)

// maxImportBytes caps one list. A Commander list of 100 lines is about
// 5 KB, and an Archidekt line with its category about 60 bytes.
const maxImportBytes = 128 * 1024

// commanderSize is the size of a Commander deck, the commander included.
// A list of this size with no commander mark reads as Commander (D-857).
const commanderSize = 100

var (
	errNoImportText  = errors.New("text: paste a deck list or choose a file")
	errLongImport    = fmt.Errorf("text: a deck list takes at most %d bytes", maxImportBytes)
	errLongDeckName  = fmt.Errorf("name: a deck name takes at most %d bytes", maxDeckNameBytes)
	errNoImporter    = errors.New("no deck reader is wired")
	errNoCardMatched = errors.New("no line of the list names a card")
	errNoLeader      = errors.New("the list holds no card that can lead a Commander deck")
	errBadPick       = errors.New("commander_oracle_ids: pick one or two cards of the list that can lead a deck")
	errBadPickPair   = errors.New("commander_oracle_ids: those two cards can not lead a deck together")
	errBadFormatPick = errors.New("format: pick Commander, Standard, Modern, or the house format")
	errNotImported   = errors.New("deck_id: that deck is no imported list")
	errBadDeckID     = fmt.Errorf("deck_id: %w", gzstore.ErrBadID)
)

// maxDeckNameBytes caps a deck name, as UpdateDeck caps it.
const maxDeckNameBytes = 200

// ImportDeck stores a deck list that a user brought, and a session that
// the revise turn reads (PR-70, D-845, D-851). A list that needs an
// answer comes back with the question and stores nothing: a list that is
// not Commander asks for the format (D-857), and a Commander list with no
// mark asks for the commander (D-847).
func (s *Server) ImportDeck(ctx context.Context, req *connect.Request[mtgv1.ImportDeckRequest]) (*connect.Response[mtgv1.ImportDeckResponse], error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	msg := req.Msg
	switch {
	case strings.TrimSpace(msg.GetText()) == "":
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoImportText)
	case len(msg.GetText()) > maxImportBytes:
		return nil, connect.NewError(connect.CodeInvalidArgument, errLongImport)
	case len(msg.GetName()) > maxDeckNameBytes:
		return nil, connect.NewError(connect.CodeInvalidArgument, errLongDeckName)
	case msg.GetCollectionId() != "" && !gzstore.ValidID(msg.GetCollectionId()):
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadCollectionID)
	}
	importer, ok := s.decks.(Importer)
	if !ok || s.deckStore == nil || s.index == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoImporter)
	}
	idx := s.index.Current()
	if idx == nil {
		return nil, connect.NewError(connect.CodeUnavailable, errIndexNotLoaded)
	}

	list, err := decklist.Parse(strings.NewReader(msg.GetText()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	entries, bad := decklist.Resolve(list, idx)
	if len(entries) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoCardMatched)
	}
	out := &mtgv1.ImportDeckResponse{Unresolved: bad}

	format, err := importFormat(msg.GetFormat(), list, entries)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		out.NeedsFormat = true
		return connect.NewResponse(out), nil
	}
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER && !list.Marked {
		if len(msg.GetCommanderOracleIds()) == 0 {
			out.CommanderOptions = leaders(entries)
			if len(out.CommanderOptions) == 0 {
				return nil, connect.NewError(connect.CodeInvalidArgument, errNoLeader)
			}
			return connect.NewResponse(out), nil
		}
		if entries, err = pickCommanders(entries, msg.GetCommanderOracleIds()); err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
	}

	owned, err := s.ownedCounts(ctx, uid, msg.GetCollectionId())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("collection %q: %w", msg.GetCollectionId(), err))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	// The judge reads a Commander list, so the spend cap applies (D-421).
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		if err := s.checkSpendCap(ctx, uid); err != nil {
			return nil, err
		}
	}

	now := timestamppb.New(s.now())
	session := &mtgv1.Session{
		Id:           s.store.NewID(uid),
		CollectionId: msg.GetCollectionId(),
		Status:       mtgv1.SessionStatus_SESSION_STATUS_BUILT,
		CreatedAt:    now,
		UpdatedAt:    now,
		Name:         importName(msg.GetName(), list.Name),
	}
	deck := importedDeck(entries, s.roles(idx), format)
	deck.Id, deck.SessionId, deck.Name, deck.CreatedAt = s.deckStore.NewID(uid), session.GetId(), session.GetName(), now

	acc := llm.NewAccumulator(s.prices)
	importer.ReadImport(ctx, deck, owned, acc)
	s.markOwnedPrintings(ctx, uid, session, idx, deck, nil)
	session.Usage = addUsage(&mtgv1.Usage{}, acc.Report())
	defer s.recordSpend(ctx, uid, session, &mtgv1.Usage{})

	st := s.importState(idx, deck, msg.GetCollectionId() != "", owned)
	session.Slots = st.Slots
	session.DeckIds = []string{deck.GetId()}

	sctx, cancel := detached(ctx, storeLimit)
	defer cancel()
	if err := s.deckStore.Put(sctx, uid, deck); err != nil {
		return nil, storeError(err)
	}
	if err := s.store.Put(sctx, uid, session, st.Snapshot(), 0); err != nil {
		return nil, storeError(err)
	}
	users.NoteQuietly(sctx, s.users, uid, auth.Email(ctx), users.DecksImported, s.now())
	users.NoteQuietly(sctx, s.users, uid, auth.Email(ctx), users.SessionsStarted, s.now())
	out.Deck, out.SessionId = deck, session.GetId()
	return connect.NewResponse(out), nil
}

// ReadImportBracket asks the judge again for an imported deck whose
// bracket is the floor alone (D-854). Any other deck comes back as it is.
func (s *Server) ReadImportBracket(ctx context.Context, req *connect.Request[mtgv1.ReadImportBracketRequest]) (*connect.Response[mtgv1.ReadImportBracketResponse], error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	if !gzstore.ValidID(req.Msg.GetDeckId()) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadDeckID)
	}
	importer, ok := s.decks.(Importer)
	if !ok || s.deckStore == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoImporter)
	}
	deck, err := s.deckStore.Get(ctx, uid, req.Msg.GetDeckId())
	if err != nil {
		return nil, storeError(err)
	}
	if !deck.GetImported() {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errNotImported)
	}
	if !deck.GetBracketEstimated() {
		return connect.NewResponse(&mtgv1.ReadImportBracketResponse{Deck: deck}), nil
	}
	if err := s.checkSpendCap(ctx, uid); err != nil {
		return nil, err
	}
	session, err := s.store.Get(ctx, uid, deck.GetSessionId())
	if err != nil {
		return nil, storeError(err)
	}
	owned, err := s.ownedCounts(ctx, uid, session.GetCollectionId())
	if err != nil {
		s.log.WarnContext(ctx, "owned counts unavailable", "collection", session.GetCollectionId(), "err", err)
	}
	acc := llm.NewAccumulator(s.prices)
	importer.ReadImport(ctx, deck, owned, acc)
	if report := acc.Report(); report.Calls > 0 {
		before := cloneUsage(session.GetUsage())
		session.Usage = addUsage(cloneUsage(before), report)
		s.recordSpend(ctx, uid, session, before)
	}
	sctx, cancel := detached(ctx, storeLimit)
	defer cancel()
	if err := s.deckStore.Put(sctx, uid, deck); err != nil {
		return nil, storeError(err)
	}
	return connect.NewResponse(&mtgv1.ReadImportBracketResponse{Deck: deck}), nil
}

// importFormat reads the format of a list. The pick of the user wins. A
// list with a commander mark, or with 100 cards, reads Commander. Any
// other list answers UNSPECIFIED, and the client asks Standard, Modern,
// or neither (D-857).
func importFormat(pick mtgv1.FormatId, list *decklist.List, entries []decklist.Entry) (mtgv1.FormatId, error) {
	switch pick {
	case mtgv1.FormatId_FORMAT_ID_UNSPECIFIED:
	case mtgv1.FormatId_FORMAT_ID_COMMANDER, mtgv1.FormatId_FORMAT_ID_STANDARD,
		mtgv1.FormatId_FORMAT_ID_MODERN, mtgv1.FormatId_FORMAT_ID_HOUSE:
		return pick, nil
	default:
		return 0, errBadFormatPick
	}
	if list.Marked {
		return mtgv1.FormatId_FORMAT_ID_COMMANDER, nil
	}
	n := 0
	for _, e := range entries {
		if e.Section == decklist.Main || e.Section == decklist.Commander {
			n += e.Count
		}
	}
	if n == commanderSize {
		return mtgv1.FormatId_FORMAT_ID_COMMANDER, nil
	}
	return mtgv1.FormatId_FORMAT_ID_UNSPECIFIED, nil
}

// leaders lists the cards of the main deck that can lead a Commander
// deck, in the order of the list (D-847).
func leaders(entries []decklist.Entry) []*mtgv1.DeckCard {
	var out []*mtgv1.DeckCard
	for _, e := range entries {
		if e.Section == decklist.Main && e.Card.GetCanBeCommander() {
			out = append(out, &mtgv1.DeckCard{OracleId: e.Card.GetOracleId(), Name: e.Card.GetName(), Count: 1})
		}
	}
	return out
}

// pickCommanders moves one copy of each picked card from the main deck
// to the command zone. The pick names one card, or two that can lead
// together.
func pickCommanders(entries []decklist.Entry, ids []string) ([]decklist.Entry, error) {
	if len(ids) > 2 || (len(ids) == 2 && ids[0] == ids[1]) {
		return nil, errBadPick
	}
	out := slices.Clone(entries)
	var picked []*mtgv1.Card
	for _, id := range ids {
		i := slices.IndexFunc(out, func(e decklist.Entry) bool {
			return e.Section == decklist.Main && e.Card.GetOracleId() == id && e.Card.GetCanBeCommander()
		})
		if i < 0 {
			return nil, errBadPick
		}
		picked = append(picked, out[i].Card)
		out = append(out, decklist.Entry{Card: out[i].Card, Count: 1, Section: decklist.Commander})
		if out[i].Count--; out[i].Count == 0 {
			out = slices.Delete(out, i, i+1)
		}
	}
	if len(picked) == 2 && !rules.ValidPair(picked[0], picked[1]) {
		return nil, errBadPickPair
	}
	return out, nil
}

// importedDeck builds the deck of a resolved list. The roles come from
// the role reader alone (D-848).
func importedDeck(entries []decklist.Entry, role func(*mtgv1.Card) mtgv1.CardRole, format mtgv1.FormatId) *mtgv1.Deck {
	deck := &mtgv1.Deck{Imported: true, Format: &mtgv1.Format{Id: format}}
	for _, e := range entries {
		card := &mtgv1.DeckCard{OracleId: e.Card.GetOracleId(), Name: e.Card.GetName(), Count: int32(e.Count)}
		if role != nil {
			card.Role = role(e.Card)
		}
		switch e.Section {
		case decklist.Commander:
			deck.CommanderOracleIds = append(deck.CommanderOracleIds, e.Card.GetOracleId())
		case decklist.Sideboard:
			deck.Sideboard = append(deck.Sideboard, card)
		case decklist.Companion:
			if deck.CompanionOracleId == "" {
				deck.CompanionOracleId = e.Card.GetOracleId()
			}
		default:
			deck.Cards = append(deck.Cards, card)
		}
	}
	return deck
}

// roles is the role reader of the candidate builder over one index, or
// nil when no builder is wired.
func (s *Server) roles(idx *cards.Index) func(*mtgv1.Card) mtgv1.CardRole {
	if s.builder == nil {
		return nil
	}
	return s.builder.Roles(idx)
}

// importName is the name of an imported deck: the name of the form, the
// name of the list, or a plain default.
func importName(form, list string) string {
	if n := strings.TrimSpace(form); n != "" {
		return n
	}
	if n := strings.TrimSpace(list); n != "" && len(n) <= maxDeckNameBytes {
		return n
	}
	return "Imported deck"
}

// importState is the private state of an import session. The slots hold
// what the list gives: the format, the power, the colors, and the
// commander. The facts of a turn refresh, and every row the planner would
// still ask closes as a decline, the budget row included. So the next
// message after the import revises the deck (D-851).
func (s *Server) importState(idx *cards.Index, deck *mtgv1.Deck, hasCollection bool, owned map[string]int32) *questions.State {
	st := questions.NewState(hasCollection)
	st.Slots.Format = deck.GetFormat()
	st.Ctx.Format = deck.GetFormat().GetId()
	st.Close("format")
	if deck.GetPower() != nil {
		st.Slots.Power = deck.GetPower()
		st.Close("power")
	}
	if ids := deck.GetCommanderOracleIds(); len(ids) > 0 {
		st.Slots.CommanderOracleIds = ids
		st.Close("commander")
	}
	// The colors of the list bound the cards a revision adds.
	if colors := listColors(deck, idx); len(colors) > 0 {
		st.Slots.Colors = colors
		st.Close("colors")
	}
	if hasCollection {
		st.Slots.PoolRule = mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
		st.Close("pool_rule")
	}
	st.Ctx.AfterBuild = true
	// An imported list holds cards the user may not own, so the rule of
	// D-168 sees a buy list, and the budget row closes as a decline. A
	// revision of an import then reads no cap.
	st.Ctx.BuyList = true
	if hints := s.hints(nil, st, owned); hints != nil {
		questions.RefreshFacts(st, hints)
	}
	// Each pass closes the rows the plan names. A closed row can open
	// another one, so the loop runs until the plan is empty, with a cap.
	for pass := 0; pass < 20 && s.cat != nil; pass++ {
		rows := s.cat.Plan(st.Ctx)
		if len(rows) == 0 {
			break
		}
		for _, r := range rows {
			st.Skip(r.StateKey())
		}
	}
	return st
}

// listColors is the union of the color identities of the commanders and
// the main deck, in WUBRG order.
func listColors(deck *mtgv1.Deck, idx *cards.Index) []mtgv1.Color {
	seen := map[mtgv1.Color]bool{}
	add := func(id string) {
		if c, ok := idx.ByOracleID(id); ok {
			for _, col := range c.GetColorIdentity() {
				seen[col] = true
			}
		}
	}
	for _, id := range deck.GetCommanderOracleIds() {
		add(id)
	}
	for _, c := range deck.GetCards() {
		add(c.GetOracleId())
	}
	var out []mtgv1.Color
	for _, col := range generate.AllColors {
		if seen[col] {
			out = append(out, col)
		}
	}
	return out
}
