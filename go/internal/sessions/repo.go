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
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
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
