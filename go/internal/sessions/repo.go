// Package sessions stores deck-building conversations in Firestore
// (roadmap PR-7, D-74). One session is two documents:
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
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

// schemaVersion marks the document shape.
const schemaVersion = 1

// maxStoredBytes bounds one stored payload. Firestore caps a document at
// 1 MiB, and the other fields use the rest. A conversation of 40 turns
// stays far under it.
const maxStoredBytes = 900 << 10

// maxInflatedBytes bounds the inflated read of a stored payload.
const maxInflatedBytes = 8 << 20

// ErrNotFound reports a session id no document answers.
var ErrNotFound = errors.New("session not found")

// ErrTooLarge reports a conversation that does not fit one document.
var ErrTooLarge = errors.New("session too large for one document (max 900 KiB gzip)")

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

// Put writes both documents in one batch, so the public session and the
// private state never disagree.
func (r *Repo) Put(ctx context.Context, uid string, s *mtgv1.Session, snap questions.Snapshot) error {
	if s.GetId() == "" {
		return errors.New("sessions: a session needs an id")
	}
	sessionGz, err := gzProto(s)
	if err != nil {
		return err
	}
	stateGz, err := gzJSON(snap)
	if err != nil {
		return err
	}
	if len(sessionGz) > maxStoredBytes || len(stateGz) > maxStoredBytes {
		return fmt.Errorf("%w: %d and %d bytes", ErrTooLarge, len(sessionGz), len(stateGz))
	}
	now := time.Now().UTC()
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
	}
	private := storedState{StateGz: stateGz, SchemaVersion: schemaVersion}
	// One transaction, because a session and its state must never
	// disagree. A half write would repeat a question the user answered.
	err = r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
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

// Get reads the public session alone. GetSession answers from it.
func (r *Repo) Get(ctx context.Context, uid, id string) (*mtgv1.Session, error) {
	snap, err := r.doc(uid, id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("%w: %s", ErrNotFound, id)
		}
		return nil, err
	}
	var stored storedSession
	if err := snap.DataTo(&stored); err != nil {
		return nil, fmt.Errorf("session %s: %w", id, err)
	}
	var out mtgv1.Session
	if err := ungzProto(stored.SessionGz, &out); err != nil {
		return nil, fmt.Errorf("session %s: %w", id, err)
	}
	out.Id = id
	return &out, nil
}

// GetState reads the session and its private state, for the next turn.
func (r *Repo) GetState(ctx context.Context, uid, id string) (*mtgv1.Session, questions.Snapshot, error) {
	s, err := r.Get(ctx, uid, id)
	if err != nil {
		return nil, questions.Snapshot{}, err
	}
	snap, err := r.stateDoc(uid, id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			// The public session exists and the private state does not.
			// A zero snapshot opens a new state, which asks again rather
			// than answer from nothing.
			return s, questions.Snapshot{}, nil
		}
		return nil, questions.Snapshot{}, err
	}
	var stored storedState
	if err := snap.DataTo(&stored); err != nil {
		return nil, questions.Snapshot{}, fmt.Errorf("session %s state: %w", id, err)
	}
	var out questions.Snapshot
	if err := ungzJSON(stored.StateGz, &out); err != nil {
		return nil, questions.Snapshot{}, fmt.Errorf("session %s state: %w", id, err)
	}
	if out.Version > questions.SnapshotVersion {
		return nil, questions.Snapshot{}, fmt.Errorf("session %s state: version %d is newer than %d", id, out.Version, questions.SnapshotVersion)
	}
	return s, out, nil
}

func gzProto(m *mtgv1.Session) ([]byte, error) {
	raw, err := protojson.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("sessions: encode: %w", err)
	}
	return gzBytes(raw)
}

func ungzProto(payload []byte, m *mtgv1.Session) error {
	raw, err := ungzBytes(payload)
	if err != nil {
		return err
	}
	return protojson.Unmarshal(raw, m)
}

func gzJSON(v any) ([]byte, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("sessions: encode: %w", err)
	}
	return gzBytes(raw)
}

func ungzJSON(payload []byte, v any) error {
	raw, err := ungzBytes(payload)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)
}

func gzBytes(raw []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(raw); err != nil {
		return nil, fmt.Errorf("sessions: compress: %w", err)
	}
	if err := zw.Close(); err != nil {
		return nil, fmt.Errorf("sessions: compress: %w", err)
	}
	return buf.Bytes(), nil
}

func ungzBytes(payload []byte) ([]byte, error) {
	if len(payload) == 0 {
		return []byte("{}"), nil
	}
	zr, err := gzip.NewReader(bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("sessions: read: %w", err)
	}
	defer func() { _ = zr.Close() }()
	raw, err := io.ReadAll(io.LimitReader(zr, maxInflatedBytes))
	if err != nil {
		return nil, fmt.Errorf("sessions: read: %w", err)
	}
	return raw, nil
}
