// Package sessions stores deck-building conversations in Firestore
// (D-74). One session is two documents:
//
//	users/{uid}/sessions/{id}          the proto Session, the public contract
//	users/{uid}/sessions/{id}/private/state   the agent's private state
//
// GetSession returns the first one. The second holds the facts the proto
// does not carry: which rows the agent asked, every word the user wrote,
// the card names, and the planner triggers. Without it, a resumed
// conversation repeats a question the user already answered.
package sessions

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/gzstore"
	"github.com/nkramber/decktome/go/internal/questions"
)

// schemaVersion marks the document shape.
const schemaVersion = 1

// ErrNotFound reports a session id no document answers.
var ErrNotFound = errors.New("session not found")

// ErrTooLarge reports a conversation that does not fit one document.
var ErrTooLarge = errors.New("session too large for one document (max 900 KiB gzip)")

// ErrConflict reports a Put whose expected version is stale. Another turn
// stored the session first, and this turn must not overwrite it.
var ErrConflict = errors.New("session changed since it was read")

// ErrLeased means that another build holds the lease of the session
// (D-922).
var ErrLeased = errors.New("another build holds the session")

// Repo stores sessions in Firestore. The caller owns the client.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// storedSession is the public document shape. The proto is stored as
// gzip protojson, so a proto change reads back without a schema
// migration. The flat fields exist for a future list query.
type storedSession struct {
	CollectionID  string    `firestore:"collection_id"`
	Status        string    `firestore:"status"`
	CreatedAt     time.Time `firestore:"created_at"`
	UpdatedAt     time.Time `firestore:"updated_at"`
	SessionGz     []byte    `firestore:"session_gz"`
	SchemaVersion int64     `firestore:"schema_version"`
	// Version counts the writes. Put compares it before it writes, so two
	// overlapping turns can not both land. A document without this field
	// reads as version 0.
	Version int64 `firestore:"version"`
}

// storedState is the private document shape.
type storedState struct {
	StateGz       []byte `firestore:"state_gz"`
	SchemaVersion int64  `firestore:"schema_version"`
}

func (r *Repo) doc(uid, id string) *firestore.DocumentRef {
	return r.client.Collection("users").Doc(uid).Collection("sessions").Doc(id)
}

func (r *Repo) stateDoc(uid, id string) *firestore.DocumentRef {
	return r.doc(uid, id).Collection("private").Doc("state")
}

// leaseDoc holds the build lease of one session. It sits apart from the
// session, so a lease never changes the version that Put compares.
func (r *Repo) leaseDoc(uid, id string) *firestore.DocumentRef {
	return r.doc(uid, id).Collection("private").Doc("lease")
}

// storedLease is the lease shape. The token names the build that holds
// it, and the lease ends at until.
type storedLease struct {
	Token string    `firestore:"token"`
	Until time.Time `firestore:"until"`
}

// readLease reads the lease inside a transaction. A missing lease reads
// as the zero lease, which has ended.
func readLease(tx *firestore.Transaction, ref *firestore.DocumentRef) (storedLease, error) {
	snap, err := tx.Get(ref)
	if status.Code(err) == codes.NotFound {
		return storedLease{}, nil
	}
	if err != nil {
		return storedLease{}, err
	}
	var l storedLease
	if err := snap.DataTo(&l); err != nil {
		return storedLease{}, err
	}
	return l, nil
}

// Lease gives the build lease of a session to token until the given
// time. It returns ErrLeased while another token holds a lease that has
// not ended at now. A build on each instance takes it, so a turn on
// another instance can not start a second paid build (D-922).
func (r *Repo) Lease(ctx context.Context, uid, id, token string, now, until time.Time) error {
	ref := r.leaseDoc(uid, id)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		l, err := readLease(tx, ref)
		if err != nil {
			return err
		}
		if l.Token != token && now.Before(l.Until) {
			return fmt.Errorf("%w: %s", ErrLeased, id)
		}
		return tx.Set(ref, storedLease{Token: token, Until: until})
	})
}

// Leased reports whether a build holds the lease of a session at now.
func (r *Repo) Leased(ctx context.Context, uid, id string, now time.Time) (bool, error) {
	snap, err := r.leaseDoc(uid, id).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	var l storedLease
	if err := snap.DataTo(&l); err != nil {
		return false, err
	}
	return now.Before(l.Until), nil
}

// Release ends the lease that token holds. The lease of another token
// stays, because that build took it after this lease ended.
func (r *Repo) Release(ctx context.Context, uid, id, token string) error {
	ref := r.leaseDoc(uid, id)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		l, err := readLease(tx, ref)
		if err != nil || l.Token != token {
			return err
		}
		return tx.Delete(ref)
	})
}

// NewID reserves a session id without a write.
func (r *Repo) NewID(uid string) string {
	return r.client.Collection("users").Doc(uid).Collection("sessions").NewDoc().ID
}

// Put writes both documents in one transaction, so the public session
// and the private state never disagree.
//
// expected is the version GetState returned, or 0 for a new session. The
// transaction reads the stored version first. When it differs, Put
// returns ErrConflict and writes nothing. On success the stored version
// becomes expected+1.
func (r *Repo) Put(ctx context.Context, uid string, s *mtgv1.Session, snap questions.Snapshot, expected int64) error {
	if s.GetId() == "" {
		return errors.New("sessions: a session needs an id")
	}
	sessionGz, err := gzstore.MarshalProto(s)
	if err != nil {
		return err
	}
	stateGz, err := gzstore.MarshalJSON(snap)
	if err != nil {
		return err
	}
	if len(sessionGz) > gzstore.MaxStoredBytes || len(stateGz) > gzstore.MaxStoredBytes {
		return fmt.Errorf("%w: %d and %d bytes", ErrTooLarge, len(sessionGz), len(stateGz))
	}
	// The service stamps updated_at, so the flat field agrees with the
	// proto.
	now := time.Now().UTC()
	if t := s.GetUpdatedAt(); t != nil {
		now = t.AsTime()
	}
	created := now
	if t := s.GetCreatedAt(); t != nil {
		created = t.AsTime()
	}
	public := storedSession{
		CollectionID:  s.GetCollectionId(),
		Status:        s.GetStatus().String(),
		CreatedAt:     created,
		UpdatedAt:     now,
		SessionGz:     sessionGz,
		SchemaVersion: schemaVersion,
		Version:       expected + 1,
	}
	private := storedState{StateGz: stateGz, SchemaVersion: schemaVersion}
	// One transaction, because a session and its state must never
	// disagree. The read inside it is the version check: Firestore
	// retries the transaction when the document changes under it, so the
	// check and the write are one step.
	err = r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		current, err := storedVersion(tx, r.doc(uid, s.GetId()))
		if err != nil {
			return err
		}
		if current != expected {
			return fmt.Errorf("%w: stored version %d, expected %d", ErrConflict, current, expected)
		}
		if err := tx.Set(r.doc(uid, s.GetId()), public); err != nil {
			return err
		}
		return tx.Set(r.stateDoc(uid, s.GetId()), private)
	})
	if err != nil {
		return fmt.Errorf("store session %s: %w", s.GetId(), err)
	}
	return nil
}

// storedVersion reads the version inside a transaction. A missing
// document is version 0, which is what a new session expects.
func storedVersion(tx *firestore.Transaction, ref *firestore.DocumentRef) (int64, error) {
	snap, err := tx.Get(ref)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return 0, nil
		}
		return 0, err
	}
	var stored storedSession
	if err := snap.DataTo(&stored); err != nil {
		return 0, err
	}
	return stored.Version, nil
}

// Get reads the public session alone. GetSession answers from it.
func (r *Repo) Get(ctx context.Context, uid, id string) (*mtgv1.Session, error) {
	snap, err := r.doc(uid, id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("%w: %s", ErrNotFound, id)
		}
		return nil, err
	}
	s, _, err := decodeSession(id, snap)
	return s, err
}

// decodeSession reads the public document into a proto and its version.
func decodeSession(id string, snap *firestore.DocumentSnapshot) (*mtgv1.Session, int64, error) {
	var stored storedSession
	if err := snap.DataTo(&stored); err != nil {
		return nil, 0, fmt.Errorf("session %s: %w", id, err)
	}
	var out mtgv1.Session
	if err := gzstore.UnmarshalProto(stored.SessionGz, &out); err != nil {
		return nil, 0, fmt.Errorf("session %s: %w", id, err)
	}
	out.Id = id
	return &out, stored.Version, nil
}

// GetState reads the session and its private state in one read, for the
// next turn. The version it returns is the one Put must expect.
func (r *Repo) GetState(ctx context.Context, uid, id string) (*mtgv1.Session, questions.Snapshot, int64, error) {
	snaps, err := r.client.GetAll(ctx, []*firestore.DocumentRef{r.doc(uid, id), r.stateDoc(uid, id)})
	if err != nil {
		return nil, questions.Snapshot{}, 0, err
	}
	if len(snaps) != 2 || !snaps[0].Exists() {
		return nil, questions.Snapshot{}, 0, fmt.Errorf("%w: %s", ErrNotFound, id)
	}
	s, version, err := decodeSession(id, snaps[0])
	if err != nil {
		return nil, questions.Snapshot{}, 0, err
	}
	if !snaps[1].Exists() {
		// The public session exists and the private state does not. A
		// zero snapshot opens a new state, which asks again rather than
		// answer from nothing. It keeps the one fact the session carries.
		var out questions.Snapshot
		out.Ctx.HasCollection = s.GetCollectionId() != ""
		return s, out, version, nil
	}
	var stored storedState
	if err := snaps[1].DataTo(&stored); err != nil {
		return nil, questions.Snapshot{}, 0, fmt.Errorf("session %s state: %w", id, err)
	}
	var out questions.Snapshot
	if err := gzstore.UnmarshalJSON(stored.StateGz, &out); err != nil {
		return nil, questions.Snapshot{}, 0, fmt.Errorf("session %s state: %w", id, err)
	}
	if out.Version > questions.SnapshotVersion {
		return nil, questions.Snapshot{}, 0, fmt.Errorf("session %s state: version %d is newer than %d", id, out.Version, questions.SnapshotVersion)
	}
	return s, out, version, nil
}

// ListScanCap bounds the sessions one list reads, as the deck list does.
// A reader with more has the newest ones, which is what a list of
// conversations is for.
const ListScanCap = 500

// List returns the reader's sessions, newest first by updated_at, as
// summaries (roadmap PR-19). It inflates each stored session, because a
// session is a few kilobytes and the summary needs its first message
// and its deck count. The caller pages over the slice.
func (r *Repo) List(ctx context.Context, uid string) ([]*mtgv1.SessionSummary, error) {
	snaps, err := r.client.Collection("users").Doc(uid).Collection("sessions").
		OrderBy("updated_at", firestore.Desc).Limit(ListScanCap).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	out := make([]*mtgv1.SessionSummary, 0, len(snaps))
	for _, snap := range snaps {
		s, _, err := decodeSession(snap.Ref.ID, snap)
		if err != nil {
			return nil, err
		}
		out = append(out, Summarize(s))
	}
	return out, nil
}

// Summarize reads the list row of one session.
func Summarize(s *mtgv1.Session) *mtgv1.SessionSummary {
	first := ""
	for _, t := range s.GetTurns() {
		if m := strings.TrimSpace(t.GetUserMessage()); m != "" {
			first = m
			break
		}
	}
	return &mtgv1.SessionSummary{
		Id:           s.GetId(),
		Name:         s.GetName(),
		FirstMessage: first,
		CollectionId: s.GetCollectionId(),
		Status:       s.GetStatus(),
		DeckCount:    int32(len(s.GetDeckIds())), //nolint:gosec // a session holds a handful of decks
		Usage:        s.GetUsage(),
		CreatedAt:    s.GetCreatedAt(),
		UpdatedAt:    s.GetUpdatedAt(),
	}
}

// Rename writes the name of one session (roadmap PR-19). It rewrites the
// public document alone, inside the version check, so a turn that lands
// at the same time can not lose the name or the turn. The updated_at
// stays: a rename must not move a conversation up the list.
func (r *Repo) Rename(ctx context.Context, uid, id, name string) (*mtgv1.SessionSummary, error) {
	var out *mtgv1.SessionSummary
	err := r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		snap, err := tx.Get(r.doc(uid, id))
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return fmt.Errorf("%w: %s", ErrNotFound, id)
			}
			return err
		}
		var stored storedSession
		if err := snap.DataTo(&stored); err != nil {
			return fmt.Errorf("session %s: %w", id, err)
		}
		var s mtgv1.Session
		if err := gzstore.UnmarshalProto(stored.SessionGz, &s); err != nil {
			return fmt.Errorf("session %s: %w", id, err)
		}
		s.Id = id
		s.Name = name
		gz, err := gzstore.MarshalProto(&s)
		if err != nil {
			return err
		}
		stored.SessionGz = gz
		stored.Version++
		out = Summarize(&s)
		return tx.Set(r.doc(uid, id), stored)
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Delete removes a session and its private state for good (roadmap
// PR-19). The decks it built stay: a deck carries its own id, and the
// deck library reads it without the session.
func (r *Repo) Delete(ctx context.Context, uid, id string, now time.Time) error {
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		if _, err := tx.Get(r.doc(uid, id)); err != nil {
			if status.Code(err) == codes.NotFound {
				return fmt.Errorf("%w: %s", ErrNotFound, id)
			}
			return err
		}
		// The lease check sits in the transaction of the delete, so a
		// build that took the lease after any earlier read still stops it
		// (D-922).
		l, err := readLease(tx, r.leaseDoc(uid, id))
		if err != nil {
			return err
		}
		if now.Before(l.Until) {
			return fmt.Errorf("%w: %s", ErrLeased, id)
		}
		if err := tx.Delete(r.doc(uid, id)); err != nil {
			return err
		}
		if err := tx.Delete(r.leaseDoc(uid, id)); err != nil {
			return err
		}
		return tx.Delete(r.stateDoc(uid, id))
	})
}
