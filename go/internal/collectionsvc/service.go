// Package collectionsvc serves CollectionService (roadmap PR-4).
package collectionsvc

import (
	"bytes"
	"context"
	"errors"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
)

// maxUpload bounds an uploaded file. The owner's 2,548-row export is
// 471 KB. Ten times that is generous.
const maxUpload = 5 << 20

// IndexSource hands out the current card index.
type IndexSource interface {
	Current() *cards.Index
}

// UserResolver names the acting user. PR-11 replaces the debug
// implementation with Firebase Auth.
type UserResolver func(ctx context.Context) string

// Server answers CollectionService requests.
type Server struct {
	mtgv1connect.UnimplementedCollectionServiceHandler
	repo  *collections.Repo
	index IndexSource
	user  UserResolver
}

// New wires the service.
func New(repo *collections.Repo, index IndexSource, user UserResolver) *Server {
	return &Server{repo: repo, index: index, user: user}
}

var (
	errTooLarge   = errors.New("upload larger than 5 MiB")
	errNoIndex    = errors.New("card database not loaded yet")
	errBadSource  = errors.New("source must be MANABOX_CSV or ARENA_TEXT")
	errEmptyBody  = errors.New("content is empty")
	errNoResolved = errors.New("no row resolved to a card: the report lists every row")
)

// ImportCollection parses, resolves, and stores an upload.
// Every input row lands in the collection or in the report (PR-4 gate).
func (s *Server) ImportCollection(ctx context.Context, req *connect.Request[mtgv1.ImportCollectionRequest]) (*connect.Response[mtgv1.ImportCollectionResponse], error) {
	content := req.Msg.Content
	if len(content) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errEmptyBody)
	}
	if len(content) > maxUpload {
		return nil, connect.NewError(connect.CodeInvalidArgument, errTooLarge)
	}
	idx := s.index.Current()
	if idx == nil {
		return nil, connect.NewError(connect.CodeUnavailable, errNoIndex)
	}

	var rows []collections.Row
	var badParse []*mtgv1.UnresolvedRow
	var err error
	switch req.Msg.Source {
	case mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV:
		rows, badParse, err = collections.ParseManaBoxCSV(bytes.NewReader(content))
	case mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT:
		rows, badParse, err = collections.ParseArenaText(bytes.NewReader(content))
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadSource)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	entries, badResolve := collections.Resolve(rows, idx)
	report := &mtgv1.ImportReport{
		Unresolved:    append(badParse, badResolve...),
		ResolvedCount: int32(len(entries)),
	}
	if len(entries) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoResolved)
	}
	collections.SortEntries(entries)

	col := &mtgv1.Collection{
		Name:        req.Msg.Name,
		Source:      req.Msg.Source,
		ContentHash: collections.ContentHash(content),
		Entries:     entries,
		CardCount:   collections.CardCount(entries),
	}
	uid := s.user(ctx)
	// An identical re-upload updates the existing document (D-16).
	if existing, err := s.repo.FindByHash(ctx, uid, col.ContentHash); err == nil && existing != "" {
		col.Id = existing
	}
	id, err := s.repo.Put(ctx, uid, col)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	col.Id = id
	return connect.NewResponse(&mtgv1.ImportCollectionResponse{Collection: col, Report: report}), nil
}

// GetCollection returns one collection with entries.
func (s *Server) GetCollection(ctx context.Context, req *connect.Request[mtgv1.GetCollectionRequest]) (*connect.Response[mtgv1.GetCollectionResponse], error) {
	col, err := s.repo.Get(ctx, s.user(ctx), req.Msg.CollectionId)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	return connect.NewResponse(&mtgv1.GetCollectionResponse{Collection: col}), nil
}

// ListCollections returns the user's collections without entries.
func (s *Server) ListCollections(ctx context.Context, _ *connect.Request[mtgv1.ListCollectionsRequest]) (*connect.Response[mtgv1.ListCollectionsResponse], error) {
	cols, err := s.repo.List(ctx, s.user(ctx))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.ListCollectionsResponse{Collections: cols}), nil
}
