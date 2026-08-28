// Package decks stores built decks in Firestore (roadmap PR-8).
//
// PR-8 built a deck, streamed it, and let it go. Nothing held it, so
// `session.deck_ids` stayed empty, `Context.AfterBuild` was never true,
// and the variance row was dead for every real user. `GetDeck`,
// `ListDecks`, and `Export` had nothing to read, and the staleness job of
// PR-3 had nothing to re-validate (D-245).
//
// The shape follows internal/sessions: one document per deck under the
// user, the proto stored as gzip protojson so a proto change reads back
// with no migration, and flat fields for the list query.
package decks

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// schemaVersion counts the stored shape, not the proto.
const schemaVersion = 1

// maxInflatedBytes bounds the inflated read of a stored payload. A deck
// is far smaller than a conversation, and the bound is the same.
const maxInflatedBytes = 8 << 20

// maxStoredBytes is the Firestore document limit with room to spare.
const maxStoredBytes = 900 << 10

// ErrNotFound reports a deck id no document answers.
var ErrNotFound = errors.New("deck not found")

// ErrTooLarge reports a deck that does not fit one document.
var ErrTooLarge = errors.New("deck too large for one document (max 900 KiB gzip)")

// Repo stores decks in Firestore. The caller owns the client.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// storedDeck is the document shape. The flat fields answer the list
// query without inflating every deck.
type storedDeck struct {
	SessionID     string    `firestore:"session_id"`
	Name          string    `firestore:"name"`
	FormatID      int64     `firestore:"format_id"`
	CreatedAt     time.Time `firestore:"created_at"`
	LegalityAsOf  string    `firestore:"legality_as_of"`
	BuyCostUSD    float64   `firestore:"buy_cost_usd"`
	Stale         bool      `firestore:"stale"`
	DeckGz        []byte    `firestore:"deck_gz"`
	SchemaVersion int64     `firestore:"schema_version"`
}

func (r *Repo) col(uid string) *firestore.CollectionRef {
	return r.client.Collection("users").Doc(uid).Collection("decks")
}

func (r *Repo) doc(uid, id string) *firestore.DocumentRef { return r.col(uid).Doc(id) }

// NewID reserves a deck id without a write. The build needs the id before
// it stores the deck, because the deck carries its own id.
func (r *Repo) NewID(uid string) string { return r.col(uid).NewDoc().ID }

// Put writes one deck. A deck is written once and never updated, so
// there is no version to compare: a re-roll is a new deck with a new id
// (D-18).
func (r *Repo) Put(ctx context.Context, uid string, d *mtgv1.Deck) error {
	if d.GetId() == "" {
		return errors.New("decks: a deck needs an id")
	}
	if uid == "" {
		return errors.New("decks: a deck needs a user")
	}
	payload, err := gzProto(d)
	if err != nil {
		return err
	}
	if len(payload) > maxStoredBytes {
		return ErrTooLarge
	}
	created := d.GetCreatedAt().AsTime()
	if d.GetCreatedAt() == nil {
		created = time.Now().UTC()
	}
	_, err = r.doc(uid, d.GetId()).Set(ctx, storedDeck{
		SessionID:     d.GetSessionId(),
		Name:          d.GetName(),
		FormatID:      int64(d.GetFormat().GetId()),
		CreatedAt:     created,
		LegalityAsOf:  d.GetLegalityAsOf(),
		BuyCostUSD:    d.GetBuyCostUsd(),
		Stale:         d.GetStale(),
		DeckGz:        payload,
		SchemaVersion: schemaVersion,
	})
	return err
}

// Get reads one deck.
func (r *Repo) Get(ctx context.Context, uid, id string) (*mtgv1.Deck, error) {
	snap, err := r.doc(uid, id).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	var sd storedDeck
	if err := snap.DataTo(&sd); err != nil {
		return nil, err
	}
	var d mtgv1.Deck
	if err := ungzProto(sd.DeckGz, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// List reads the user's decks, newest first. limit of 0 reads them all.
func (r *Repo) List(ctx context.Context, uid string, limit int) ([]*mtgv1.Deck, error) {
	q := r.col(uid).OrderBy("created_at", firestore.Desc)
	if limit > 0 {
		q = q.Limit(limit)
	}
	it := q.Documents(ctx)
	defer it.Stop()
	var out []*mtgv1.Deck
	for {
		snap, err := it.Next()
		if errors.Is(err, iterator.Done) {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
		var sd storedDeck
		if err := snap.DataTo(&sd); err != nil {
			return nil, err
		}
		var d mtgv1.Deck
		if err := ungzProto(sd.DeckGz, &d); err != nil {
			return nil, err
		}
		out = append(out, &d)
	}
}

func gzProto(m *mtgv1.Deck) ([]byte, error) {
	raw, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(raw); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ungzProto(payload []byte, m *mtgv1.Deck) error {
	zr, err := gzip.NewReader(bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer func() { _ = zr.Close() }()
	raw, err := io.ReadAll(io.LimitReader(zr, maxInflatedBytes))
	if err != nil {
		return err
	}
	if len(raw) == maxInflatedBytes {
		return fmt.Errorf("decks: stored deck is larger than %d bytes", maxInflatedBytes)
	}
	return protojson.Unmarshal(raw, m)
}
