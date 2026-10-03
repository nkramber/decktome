package agentsvc

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/archidekt"
	"github.com/nkramber/decktome/go/internal/decklink"
)

// FetchesPerMinute caps the deck reads of one user. A user pastes one
// link for each import, so the cap stops a loop and no real use (D-1100).
const FetchesPerMinute = 10

var (
	errNoFetcher     = errors.New("no deck URL reader is wired")
	errManyFetches   = errors.New("url: too many deck links in one minute. Wait a minute, then try again")
	errBadSourceLink = errors.New("source_url: the link names no Archidekt deck")
)

// FetchDeckList reads a public deck by its link, and answers its list as
// text (PR-120, D-1100). The import form shows the text, and ImportDeck
// reads it as a pasted list. It stores nothing. A site that the app can
// not read gets the steps of an export, and the read of such a site never
// leaves the server (D-1103). The client files a report of a site with no
// steps of its own (D-1104).
func (s *Server) FetchDeckList(ctx context.Context, req *connect.Request[mtgv1.FetchDeckListRequest]) (*connect.Response[mtgv1.FetchDeckListResponse], error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	link, err := decklink.Parse(req.Msg.GetUrl())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	out := &mtgv1.FetchDeckListResponse{Site: link.Name(), KnownSite: link.Site != decklink.Unknown}
	if link.Site != decklink.Archidekt {
		out.ExportSteps = link.Steps()
		return connect.NewResponse(out), nil
	}
	if s.archidekt == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoFetcher)
	}
	if !s.fetches.Allow(uid) {
		return nil, connect.NewError(connect.CodeResourceExhausted, errManyFetches)
	}
	deck, err := s.archidekt.Fetch(ctx, link.DeckID)
	if errors.Is(err, archidekt.ErrNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, archidekt.ErrNotFound)
	}
	if err != nil {
		s.log.WarnContext(ctx, "archidekt read failed", "deck", link.DeckID, "err", err)
		return nil, connect.NewError(connect.CodeUnavailable, archidekt.ErrUnavailable)
	}
	out.Text, out.Name, out.SourceUrl = deck.Text, deck.Name, archidekt.DeckURL(link.DeckID)
	out.LeftOut = int32(min(deck.LeftOut, 1<<30)) //nolint:gosec // capped above
	return connect.NewResponse(out), nil
}

// sourceURL is the link that an imported deck keeps: the deck page made
// again from the id, and never the text of the request (D-1101). An empty
// request keeps no link.
func sourceURL(raw string) (string, error) {
	if raw == "" {
		return "", nil
	}
	id, err := archidekt.ParseURL(raw)
	if err != nil {
		return "", errBadSourceLink
	}
	return archidekt.DeckURL(id), nil
}
