package agentsvc

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/gzstore"
	"github.com/nkramber/decktome/go/internal/sessions"
)

// The sessions list of PR-19 (D-433): the conversations a reader can
// resume, rename, and delete. The list reads summaries, never turns.

// MaxSessionNameBytes caps a session name, as a deck name is capped.
const MaxSessionNameBytes = 200

// DefaultSessionPage and MaxSessionPage bound one list answer.
const (
	DefaultSessionPage = 50
	MaxSessionPage     = 200
)

var (
	errEmptySessionName = errors.New("a session name can not be empty")
	errLongSessionName  = fmt.Errorf("name is longer than %d bytes", MaxSessionNameBytes)
	errBadSessionToken  = errors.New("the page token belongs to another listing")
	errDeleteMidBuild   = errors.New("a build is in progress, so the session waits")
)

// sessionRef reads the user and the session id of a request, and it
// refuses an empty or a malformed id.
func (s *Server) sessionRef(ctx context.Context, id string) (string, string, error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return "", "", connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return "", "", connect.NewError(connect.CodeInvalidArgument, errNoSessionID)
	}
	if !gzstore.ValidID(id) {
		return "", "", connect.NewError(connect.CodeInvalidArgument, errBadSessionID)
	}
	return uid, id, nil
}

// ListSessions lists the reader's conversations, newest first (D-433).
// The store answers every session as a summary, and the page is a slice
// of that answer. The token carries the offset and a fingerprint of the
// listing, so a token of another reader's list is refused.
func (s *Server) ListSessions(ctx context.Context, req *connect.Request[mtgv1.ListSessionsRequest]) (*connect.Response[mtgv1.ListSessionsResponse], error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	offset, err := decodeSessionToken(req.Msg.GetPageToken(), uid)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	all, err := s.store.List(ctx, uid)
	if err != nil {
		return nil, storeError(err)
	}
	if offset > len(all) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadSessionToken)
	}
	size := int(req.Msg.GetPageSize())
	switch {
	case size <= 0:
		size = DefaultSessionPage
	case size > MaxSessionPage:
		size = MaxSessionPage
	}
	next := ""
	page := all[offset:]
	if end := offset + size; end < len(all) {
		page = all[offset:end]
		next = encodeSessionToken(end, uid)
	}
	return connect.NewResponse(&mtgv1.ListSessionsResponse{Sessions: page, NextPageToken: next}), nil
}

func sessionFingerprint(uid string) string {
	sum := sha256.Sum256([]byte("sessions\x00" + uid))
	return hex.EncodeToString(sum[:6])
}

func encodeSessionToken(offset int, uid string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%d.%s", offset, sessionFingerprint(uid))))
}

func decodeSessionToken(token, uid string) (int, error) {
	if token == "" {
		return 0, nil
	}
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return 0, errBadSessionToken
	}
	offsetText, fingerprint, ok := strings.Cut(string(raw), ".")
	if !ok || fingerprint != sessionFingerprint(uid) {
		return 0, errBadSessionToken
	}
	n, err := strconv.Atoi(offsetText)
	if err != nil || n < 0 {
		return 0, errBadSessionToken
	}
	return n, nil
}

// UpdateSession writes the name of one conversation (D-433). The name is
// the one field it writes, so a rename never touches a turn, and the
// updated_at stays: a rename must not move a conversation up the list.
func (s *Server) UpdateSession(ctx context.Context, req *connect.Request[mtgv1.UpdateSessionRequest]) (*connect.Response[mtgv1.UpdateSessionResponse], error) {
	uid, id, err := s.sessionRef(ctx, req.Msg.GetSessionId())
	if err != nil {
		return nil, err
	}
	name := strings.TrimSpace(req.Msg.GetName())
	if name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errEmptySessionName)
	}
	if len(name) > MaxSessionNameBytes {
		return nil, connect.NewError(connect.CodeInvalidArgument, errLongSessionName)
	}
	sum, err := s.store.Rename(ctx, uid, id, name)
	if err != nil {
		return nil, storeError(err)
	}
	return connect.NewResponse(&mtgv1.UpdateSessionResponse{Session: sum}), nil
}

// SetPageHidden records that the page of a session went to the
// background, or came back. A phone can suspend the page and keep its
// stream open, so the server can not see the page leave. The build reads
// this record when it stores its deck, and it sends the push of a
// finished build to a page that is hidden (D-1033, F-192).
func (s *Server) SetPageHidden(ctx context.Context, req *connect.Request[mtgv1.SetPageHiddenRequest]) (*connect.Response[mtgv1.SetPageHiddenResponse], error) {
	uid, id, err := s.sessionRef(ctx, req.Msg.GetSessionId())
	if err != nil {
		return nil, err
	}
	if err := s.store.SetHidden(ctx, uid, id, req.Msg.GetHidden(), s.now()); err != nil {
		return nil, storeError(err)
	}
	return connect.NewResponse(&mtgv1.SetPageHiddenResponse{}), nil
}

// DeleteSession removes a conversation and the decks it built, for good
// (D-433, D-456). A session with a build in flight waits: the build writes
// its deck id onto the session when it ends, and a delete under it
// would bring the session back half-written.
func (s *Server) DeleteSession(ctx context.Context, req *connect.Request[mtgv1.DeleteSessionRequest]) (*connect.Response[mtgv1.DeleteSessionResponse], error) {
	uid, id, err := s.sessionRef(ctx, req.Msg.GetSessionId())
	if err != nil {
		return nil, err
	}
	if _, busy := s.building.Load(buildKey(uid, id)); busy {
		return nil, connect.NewError(connect.CodeAborted, errDeleteMidBuild)
	}
	// A chat and its decks are one thing (D-456). The chat goes first,
	// so a deck that outlives a failure is visible and can be deleted
	// again from its tile.
	session, err := s.store.Get(ctx, uid, id)
	if err != nil {
		return nil, storeError(err)
	}
	// A build on another instance holds the lease in the store, and the
	// delete reads it in its own transaction (D-922).
	if err := s.store.Delete(ctx, uid, id, s.now()); err != nil {
		if errors.Is(err, sessions.ErrLeased) {
			return nil, connect.NewError(connect.CodeAborted, errDeleteMidBuild)
		}
		return nil, storeError(err)
	}
	if s.deckStore != nil {
		for _, did := range session.GetDeckIds() {
			if err := s.deckStore.Delete(ctx, uid, did); err != nil && !errors.Is(err, decks.ErrNotFound) {
				return nil, connect.NewError(connect.CodeInternal, err)
			}
		}
	}
	return connect.NewResponse(&mtgv1.DeleteSessionResponse{}), nil
}

// sendPhase streams where the turn stands (D-435). A client that left
// can not receive it, and the turn goes on.
//
// The lease of a running build records the step too, so a page that
// reloads during the build shows the same stepper (F-200, D-1042). The
// write runs detached, because a client that left is the one that needs
// it. The end of a turn records nothing, as the release of the lease
// ends the record.
func (s *Server) sendPhase(ctx context.Context, stream *connect.ServerStream[mtgv1.ChatResponse], uid string, session *mtgv1.Session, phase mtgv1.BuildPhase) {
	s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Phase{Phase: phase}})
	if phase == mtgv1.BuildPhase_BUILD_PHASE_DONE {
		return
	}
	s.recordStep(ctx, uid, session, phase, "")
}

// sendStep streams the working line of a build and records it with the
// lease, so a reloaded page shows the same line (D-1042).
func (s *Server) sendStep(ctx context.Context, stream *connect.ServerStream[mtgv1.ChatResponse], uid string, session *mtgv1.Session, status string) {
	s.sendOrLog(ctx, stream, session, &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: status}})
	s.recordStep(ctx, uid, session, mtgv1.BuildPhase_BUILD_PHASE_UNSPECIFIED, status)
}

// recordStep writes the step to the lease. The write runs detached,
// because a client that left is the one that needs it.
func (s *Server) recordStep(ctx context.Context, uid string, session *mtgv1.Session, phase mtgv1.BuildPhase, status string) {
	if session.GetId() == "" {
		return
	}
	wctx, cancel := detached(ctx, storeLimit)
	defer cancel()
	if err := s.store.SetStep(wctx, uid, session.GetId(), phase, status); err != nil {
		s.log.WarnContext(ctx, "the build step was not recorded, so a reloaded page shows no step",
			"session", session.GetId(), "phase", phase.String(), "status", status, "err", err)
	}
}
