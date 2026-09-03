// Package cardsvc serves CardService from the in-memory card index.
package cardsvc

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// IndexSource hands out the current card index, or nil before the first
// snapshot loads. *Server is the one implementation, and every service
// that reads cards takes this interface.
type IndexSource interface {
	Current() *cards.Index
}

// Server answers CardService requests from the current index.
// The index pointer swaps atomically on refresh. Requests do not block.
type Server struct {
	mtgv1connect.UnimplementedCardServiceHandler
	index atomic.Pointer[cards.Index]
	// quality answers the quality rows of a card, nil with no model
	// (PR-14B).
	quality atomic.Pointer[QualitySource]
}

// QualitySource answers the quality rows of one card by Oracle id. The
// quality scorer's CardQualities is the one implementation.
type QualitySource func(oracleID string) []*mtgv1.CardQuality

// SetQuality installs the quality source. GetCards then carries the
// rows. A nil source removes them.
func (s *Server) SetQuality(src QualitySource) {
	if src == nil {
		s.quality.Store(nil)
		return
	}
	s.quality.Store(&src)
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

var errNoSnapshot = errors.New("card database not loaded yet: no snapshot available")

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
	errNoKey    = errors.New("lookup needs a name, scryfall_id, or oracle_id")
	errNotFound = errors.New("card not found")
)

const (
	defaultPageSize = 50
	maxPageSize     = 200
)

// Search filters the card database with structured filters.
func (s *Server) Search(_ context.Context, req *connect.Request[mtgv1.SearchRequest]) (*connect.Response[mtgv1.SearchResponse], error) {
	idx, err := s.ready()
	if err != nil {
		return nil, err
	}
	size := int(req.Msg.PageSize)
	if size <= 0 {
		size = defaultPageSize
	}
	if size > maxPageSize {
		size = maxPageSize
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

var errBadToken = errors.New("page_token must be a non-negative integer")

// MaxGetCards caps one GetCards call. A Commander deck holds 100 cards,
// and 120 leaves room for a sideboard and two commanders.
const MaxGetCards = 120

// GetCards returns the cards of up to MaxGetCards Oracle ids, in request
// order. A repeated id counts once and comes back once. An unknown id
// lands in missing_oracle_ids, so a client never reads a short list as
// a full one.
func (s *Server) GetCards(_ context.Context, req *connect.Request[mtgv1.GetCardsRequest]) (*connect.Response[mtgv1.GetCardsResponse], error) {
	idx, err := s.ready()
	if err != nil {
		return nil, err
	}
	seen := make(map[string]struct{}, len(req.Msg.GetOracleIds()))
	ids := make([]string, 0, len(req.Msg.GetOracleIds()))
	for _, id := range req.Msg.GetOracleIds() {
		if _, dup := seen[id]; dup || id == "" {
			continue
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}
	if len(ids) > MaxGetCards {
		return nil, connect.NewError(connect.CodeInvalidArgument, errTooManyIDs)
	}
	res := &mtgv1.GetCardsResponse{}
	quality := s.quality.Load()
	for _, id := range ids {
		if c, ok := idx.ByOracleID(id); ok {
			// The index card is shared, so the quality rows go on a copy
			// (PR-14B).
			if quality != nil {
				if rows := (*quality)(id); len(rows) > 0 {
					c = proto.Clone(c).(*mtgv1.Card)
					c.Quality = rows
				}
			}
			res.Cards = append(res.Cards, c)
		} else {
			res.MissingOracleIds = append(res.MissingOracleIds, id)
		}
	}
	return connect.NewResponse(res), nil
}

var errTooManyIDs = errors.New("oracle_ids holds more than 120 distinct ids")
