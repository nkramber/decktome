// Package cardsvc serves CardService from the in-memory card index.
package cardsvc

import (
	"context"
	"strconv"
	"sync/atomic"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// Server answers CardService requests from the current index.
// The index pointer swaps atomically on refresh. Requests never block.
type Server struct {
	mtgv1connect.UnimplementedCardServiceHandler
	index atomic.Pointer[cards.Index]
}

// New returns a Server with no index. Swap installs one.
func New() *Server { return &Server{} }

// Swap installs a new index. The old one serves in-flight requests.
func (s *Server) Swap(idx *cards.Index) { s.index.Store(idx) }

// Current returns the current index, or nil before the first Swap.
func (s *Server) Current() *cards.Index { return s.index.Load() }

func (s *Server) ready() (*cards.Index, error) {
	idx := s.index.Load()
	if idx == nil {
		return nil, connect.NewError(connect.CodeUnavailable, errNoSnapshot)
	}
	return idx, nil
}

var errNoSnapshot = &snapshotError{}

type snapshotError struct{}

func (*snapshotError) Error() string {
	return "card database not loaded yet: no snapshot available"
}

// Lookup finds one card by exact name, printing id, or Oracle id.
func (s *Server) Lookup(_ context.Context, req *connect.Request[mtgv1.LookupRequest]) (*connect.Response[mtgv1.LookupResponse], error) {
	idx, err := s.ready()
	if err != nil {
		return nil, err
	}
	var c *mtgv1.Card
	var ok bool
	switch key := req.Msg.Key.(type) {
	case *mtgv1.LookupRequest_Name:
		c, ok = idx.ByName(key.Name)
	case *mtgv1.LookupRequest_ScryfallId:
		c, ok = idx.ByPrintingID(key.ScryfallId)
	case *mtgv1.LookupRequest_OracleId:
		c, ok = idx.ByOracleID(key.OracleId)
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoKey)
	}
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, errNotFound)
	}
	return connect.NewResponse(&mtgv1.LookupResponse{Card: c}), nil
}

var (
	errNoKey    = &keyError{"lookup needs a name, scryfall_id, or oracle_id"}
	errNotFound = &keyError{"card not found"}
)

type keyError struct{ msg string }

func (e *keyError) Error() string { return e.msg }

const defaultPageSize = 50

// Search filters the card database with structured filters.
func (s *Server) Search(_ context.Context, req *connect.Request[mtgv1.SearchRequest]) (*connect.Response[mtgv1.SearchResponse], error) {
	idx, err := s.ready()
	if err != nil {
		return nil, err
	}
	size := int(req.Msg.PageSize)
	if size <= 0 || size > 200 {
		size = defaultPageSize
	}
	offset := 0
	if req.Msg.PageToken != "" {
		offset, err = strconv.Atoi(req.Msg.PageToken)
		if err != nil || offset < 0 {
			return nil, connect.NewError(connect.CodeInvalidArgument, errBadToken)
		}
	}
	res, total := idx.Search(cards.SearchQuery{
		ColorsWithin: req.Msg.ColorsWithin,
		TypeContains: req.Msg.TypeContains,
		Keywords:     req.Msg.Keywords,
		OracleTags:   req.Msg.OracleTags,
		LegalIn:      req.Msg.LegalIn,
		Limit:        size,
		Offset:       offset,
	})
	next := ""
	if offset+len(res) < total {
		next = strconv.Itoa(offset + len(res))
	}
	return connect.NewResponse(&mtgv1.SearchResponse{Cards: res, NextPageToken: next}), nil
}

var errBadToken = &keyError{"page_token must be a non-negative integer"}
