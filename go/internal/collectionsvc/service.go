// Package collectionsvc serves CollectionService.
package collectionsvc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
	Get(ctx context.Context, uid, id string) (*mtgv1.Collection, error)
	List(ctx context.Context, uid string) ([]*mtgv1.Collection, error)
	FindByHash(ctx context.Context, uid, hash string) (string, error)
	Delete(ctx context.Context, uid, id string) error
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
	col, err := s.repo.Get(ctx, uid, id)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("collection %q: %w", id, err))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.GetCollectionResponse{Collection: col}), nil
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
