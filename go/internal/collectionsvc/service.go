// Package collectionsvc serves CollectionService.
package collectionsvc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/auth"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
)

// maxUpload bounds an uploaded file. A 2,500-row ManaBox export is
// about 470 KB. Ten times that is generous.
const maxUpload = 5 << 20

// MaxNameBytes caps a collection name.
const MaxNameBytes = 200

// Repo is the storage the service needs. *collections.Repo satisfies it.
type Repo interface {
	Put(ctx context.Context, uid string, col *mtgv1.Collection) (string, error)
	// Get returns a collection the caller owns. Paging slices the entry
	// list in place, so an implementation that shares one object across
	// calls hands the next reader a collection the last page truncated.
	Get(ctx context.Context, uid, id string) (*mtgv1.Collection, error)
	// GetHead reads one collection without its entries (D-392).
	GetHead(ctx context.Context, uid, id string) (*mtgv1.Collection, error)
	List(ctx context.Context, uid string) ([]*mtgv1.Collection, error)
	FindByHash(ctx context.Context, uid, hash string) (string, error)
	Delete(ctx context.Context, uid, id string) error
	// Rename writes the name of one collection (roadmap PR-18).
	Rename(ctx context.Context, uid, id, name string) (*mtgv1.Collection, error)
	// Replace writes new entries and keeps the id, so every deck and
	// chat that names the collection still works (D-393).
	Replace(ctx context.Context, uid, id string, col *mtgv1.Collection) (*mtgv1.Collection, error)
}

// Server answers CollectionService requests.
type Server struct {
	mtgv1connect.UnimplementedCollectionServiceHandler
	repo  Repo
	index cardsvc.IndexSource
	user  auth.UserFunc
}

// New wires the service.
func New(repo Repo, index cardsvc.IndexSource, user auth.UserFunc) *Server {
	return &Server{repo: repo, index: index, user: user}
}

var (
	errTooLarge  = errors.New("upload larger than 5 MiB")
	errNoIndex   = errors.New("card database not loaded yet")
	errBadSource = errors.New("source must be MANABOX_CSV or ARENA_TEXT")
	errEmptyBody = errors.New("content is empty")
	errNoID      = errors.New("collection_id is required")
	errBadID     = fmt.Errorf("collection_id: %w", gzstore.ErrBadID)
	errNoUser    = errors.New("no user in the request context")
	errLongName  = fmt.Errorf("name is longer than %d bytes", MaxNameBytes)
)

// ImportCollection parses, resolves, and stores an upload. Every input
// row lands in the collection or in the report. An upload where no row
// resolves stores nothing and answers with an empty collection and the
// full report, so the user reads why every row failed.
func (s *Server) ImportCollection(ctx context.Context, req *connect.Request[mtgv1.ImportCollectionRequest]) (*connect.Response[mtgv1.ImportCollectionResponse], error) {
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	content := req.Msg.Content
	if len(content) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errEmptyBody)
	}
	if len(content) > maxUpload {
		return nil, connect.NewError(connect.CodeInvalidArgument, errTooLarge)
	}
	name := strings.TrimSpace(req.Msg.GetName())
	if len(name) > MaxNameBytes {
		return nil, connect.NewError(connect.CodeInvalidArgument, errLongName)
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
	unresolved := append(append([]*mtgv1.UnresolvedRow{}, badParse...), badResolve...)
	// ResolvedCount counts rows, not merged entries, so resolved plus
	// unresolved equals the input row count.
	report := &mtgv1.ImportReport{
		Unresolved:         unresolved,
		ResolvedCount:      int32(len(rows) - len(badResolve)), //nolint:gosec // bounded by maxUpload
		UnresolvedByReason: collections.ReasonCounts(unresolved),
	}
	col := &mtgv1.Collection{
		Name:        name,
		Source:      req.Msg.Source,
		ContentHash: collections.ContentHash(content),
	}
	if len(entries) == 0 {
		return connect.NewResponse(&mtgv1.ImportCollectionResponse{Collection: col, Report: report}), nil
	}
	collections.SortEntries(entries)
	col.Entries = entries
	col.CardCount = collections.CardCount(entries)
	// The binder head reads this and never the entries (D-392). The card
	// index answers the color counts, which no entry carries.
	col.Summary = collections.Summarize(entries, idx)
	// The reader read the diff and said to replace, so the entries go
	// into that same collection and every deck and chat that names it
	// still works (D-393). The name stays theirs.
	if target := strings.TrimSpace(req.Msg.GetReplaceCollectionId()); target != "" {
		if !gzstore.ValidID(target) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errBadID)
		}
		head, err := s.repo.Replace(ctx, uid, target, col)
		if err != nil {
			return nil, s.storeError(target, err)
		}
		col.Id, col.Name = head.GetId(), head.GetName()
		return connect.NewResponse(&mtgv1.ImportCollectionResponse{Collection: col, Report: report}), nil
	}
	// An identical re-upload updates the existing document (D-16). The
	// new request name wins: the user chose it. A lookup failure is an
	// error, not a new document.
	existing, err := s.repo.FindByHash(ctx, uid, col.ContentHash)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	col.Id = existing
	id, err := s.repo.Put(ctx, uid, col)
	if err != nil {
		if errors.Is(err, collections.ErrTooLarge) {
			return nil, connect.NewError(connect.CodeResourceExhausted, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	col.Id = id
	return connect.NewResponse(&mtgv1.ImportCollectionResponse{Collection: col, Report: report}), nil
}

// GetCollection returns one collection with entries. A missing document
// is NotFound, and every other store failure is Internal.
func (s *Server) GetCollection(ctx context.Context, req *connect.Request[mtgv1.GetCollectionRequest]) (*connect.Response[mtgv1.GetCollectionResponse], error) {
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	id := strings.TrimSpace(req.Msg.GetCollectionId())
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoID)
	}
	if !gzstore.ValidID(id) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadID)
	}
	// The binder head asks for no entry at all, and it reads the summary
	// the import stored (D-392). That is the whole point of the field:
	// the head moved about a megabyte before it.
	if req.Msg.GetEntriesOmitted() {
		col, err := s.repo.GetHead(ctx, uid, id)
		if err != nil {
			return nil, s.storeError(id, err)
		}
		return connect.NewResponse(&mtgv1.GetCollectionResponse{Collection: col}), nil
	}
	offset, err := pageOffset(req.Msg.GetPageToken())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	col, err := s.repo.Get(ctx, uid, id)
	if err != nil {
		return nil, s.storeError(id, err)
	}
	// The grid reads a page at a time. The whole document is one read
	// whatever the page, so the paging bounds the wire and not the store.
	// Sharding across documents is a later step (repo.go).
	entries := col.GetEntries()
	if offset > len(entries) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadPageToken)
	}
	size := int(req.Msg.GetPageSize())
	switch {
	case size <= 0:
		size = DefaultPageSize
	case size > MaxPageSize:
		size = MaxPageSize
	}
	next := ""
	if end := offset + size; end < len(entries) {
		col.Entries = entries[offset:end]
		next = pageToken(end)
	} else {
		col.Entries = entries[offset:]
	}
	s.decorate(col.GetEntries())
	return connect.NewResponse(&mtgv1.GetCollectionResponse{Collection: col, NextPageToken: next}), nil
}

// decorate fills the display fields of one page: the colors, the card
// types, and the price of the printing the reader owns (D-396). No
// document holds them. The index of the day answers, so a price is
// never stale and an older collection needs no rewrite. A row the
// index does not know keeps its empty fields, and the binder reads it
// as a row no filter matches.
func (s *Server) decorate(entries []*mtgv1.CollectionEntry) {
	if s.index == nil {
		return
	}
	idx := s.index.Current()
	if idx == nil {
		return
	}
	for _, e := range entries {
		c, ok := idx.ByOracleID(e.GetOracleId())
		if !ok {
			continue
		}
		e.Colors = c.GetColors()
		e.CardTypes = c.GetCardTypes()
		// The reader owns one printing, and its price is the one that
		// counts. The card price is the default printing's (D-231), so
		// it stands in only when the printing is unknown.
		e.PriceUsd = c.GetPriceUsd()
		if p, ok := idx.Printing(e.GetScryfallId()); ok && p.GetPriceUsd() > 0 {
			e.PriceUsd = p.GetPriceUsd()
		}
	}
}

// DefaultPageSize and MaxPageSize bound the entries of one answer
// (D-392). The binder grid draws a few rows of art at a time, and it
// asks for the next page as the reader scrolls.
const (
	DefaultPageSize = 200
	MaxPageSize     = 1000
)

// errBadPageToken reports a token that names a row this collection does
// not hold. A token of another collection reads as one of these.
var errBadPageToken = errors.New("the page token does not belong to this collection")

// pageToken writes the offset of the next page. It is the offset alone:
// a collection is one document, so a page is a slice of one read, and
// the token needs no filter fingerprint as a deck page token does.
func pageToken(offset int) string { return strconv.Itoa(offset) }

// pageOffset reads a page token. An empty token is the first page.
func pageOffset(token string) (int, error) {
	if token == "" {
		return 0, nil
	}
	n, err := strconv.Atoi(token)
	if err != nil || n < 0 {
		return 0, errBadPageToken
	}
	return n, nil
}

// storeError maps a store failure onto a Connect code. A missing
// document is NotFound, and every other failure is Internal.
func (s *Server) storeError(id string, err error) error {
	if status.Code(err) == codes.NotFound {
		return connect.NewError(connect.CodeNotFound, fmt.Errorf("collection %q: %w", id, err))
	}
	return connect.NewError(connect.CodeInternal, err)
}

// DeleteCollection removes one collection for good (D-347). A deck built
// from it keeps every card it holds. Its chat notices the collection is
// gone on the next turn, builds from the whole card database, and says
// so.
func (s *Server) DeleteCollection(ctx context.Context, req *connect.Request[mtgv1.DeleteCollectionRequest]) (*connect.Response[mtgv1.DeleteCollectionResponse], error) {
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	id := strings.TrimSpace(req.Msg.GetCollectionId())
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoID)
	}
	if !gzstore.ValidID(id) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadID)
	}
	if err := s.repo.Delete(ctx, uid, id); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("collection %q: %w", id, err))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.DeleteCollectionResponse{}), nil
}

// ListCollections returns the user's collections without entries.
func (s *Server) ListCollections(ctx context.Context, _ *connect.Request[mtgv1.ListCollectionsRequest]) (*connect.Response[mtgv1.ListCollectionsResponse], error) {
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	cols, err := s.repo.List(ctx, uid)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.ListCollectionsResponse{Collections: cols}), nil
}

// UpdateCollection writes the name of one collection. It is the rename
// of the collections list (roadmap PR-18).
//
// The name is the one field it writes, so a rename never rewrites the
// entries, and it keeps the import time: a rename must not move the
// collection up a list ordered by date.
func (s *Server) UpdateCollection(ctx context.Context, req *connect.Request[mtgv1.UpdateCollectionRequest]) (*connect.Response[mtgv1.UpdateCollectionResponse], error) {
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	id := strings.TrimSpace(req.Msg.GetCollectionId())
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoID)
	}
	if !gzstore.ValidID(id) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadID)
	}
	name := strings.TrimSpace(req.Msg.GetName())
	if name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errEmptyName)
	}
	if len(name) > MaxNameBytes {
		return nil, connect.NewError(connect.CodeInvalidArgument, errLongName)
	}
	col, err := s.repo.Rename(ctx, uid, id, name)
	if err != nil {
		return nil, s.storeError(id, err)
	}
	return connect.NewResponse(&mtgv1.UpdateCollectionResponse{Collection: col}), nil
}

// errEmptyName refuses a rename to nothing. A collection with no name is
// one the reader can not tell from another.
var errEmptyName = errors.New("a collection name can not be empty")

// DiffCollections compares an uploaded file with a stored collection,
// and it stores nothing (D-393). The reader reads what changed before
// they replace anything.
//
// The report rides along, because a reader must not replace a
// collection on a file half of which failed to resolve.
func (s *Server) DiffCollections(ctx context.Context, req *connect.Request[mtgv1.DiffCollectionsRequest]) (*connect.Response[mtgv1.DiffCollectionsResponse], error) {
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	id := strings.TrimSpace(req.Msg.GetCollectionId())
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoID)
	}
	if !gzstore.ValidID(id) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadID)
	}
	entries, report, err := s.parseUpload(req.Msg.GetSource(), req.Msg.GetContent())
	if err != nil {
		return nil, err
	}
	stored, err := s.repo.Get(ctx, uid, id)
	if err != nil {
		return nil, s.storeError(id, err)
	}
	return connect.NewResponse(&mtgv1.DiffCollectionsResponse{
		Diff:   collections.Diff(stored.GetEntries(), entries),
		Report: report,
	}), nil
}

// parseUpload reads an uploaded file into entries and an import report.
// ImportCollection and DiffCollections both read a file the same way,
// and one term per concept means one function does it.
func (s *Server) parseUpload(source mtgv1.ImportSource, content []byte) ([]*mtgv1.CollectionEntry, *mtgv1.ImportReport, error) {
	if len(content) == 0 {
		return nil, nil, connect.NewError(connect.CodeInvalidArgument, errEmptyBody)
	}
	if len(content) > maxUpload {
		return nil, nil, connect.NewError(connect.CodeInvalidArgument, errTooLarge)
	}
	idx := s.index.Current()
	if idx == nil {
		return nil, nil, connect.NewError(connect.CodeUnavailable, errNoIndex)
	}
	var rows []collections.Row
	var badParse []*mtgv1.UnresolvedRow
	var err error
	switch source {
	case mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV:
		rows, badParse, err = collections.ParseManaBoxCSV(bytes.NewReader(content))
	case mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT:
		rows, badParse, err = collections.ParseArenaText(bytes.NewReader(content))
	default:
		return nil, nil, connect.NewError(connect.CodeInvalidArgument, errBadSource)
	}
	if err != nil {
		return nil, nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	entries, badResolve := collections.Resolve(rows, idx)
	unresolved := append(append([]*mtgv1.UnresolvedRow{}, badParse...), badResolve...)
	collections.SortEntries(entries)
	return entries, &mtgv1.ImportReport{
		Unresolved:         unresolved,
		ResolvedCount:      int32(len(rows) - len(badResolve)), //nolint:gosec // bounded by maxUpload
		UnresolvedByReason: collections.ReasonCounts(unresolved),
	}, nil
}
