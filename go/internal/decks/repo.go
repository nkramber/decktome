// Package decks stores built decks in Firestore (D-245). A kept deck is
// what `GetDeck`, `ListDecks`, `Export`, and the staleness job read, and
// its id on the session is what makes `Context.AfterBuild` true.
//
// The shape follows internal/sessions: one document per deck under the
// user, the proto stored as gzip protojson so a proto change reads back
// with no migration, and flat fields for the list query.
package decks

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
)

// schemaVersion counts the stored shape, not the proto.
const schemaVersion = 1

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

// listFields are the flat fields List reads. The list never inflates a
// deck: it answers from these alone.
var listFields = []string{"session_id", "name", "format_id", "created_at", "legality_as_of", "buy_cost_usd", "stale"}

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
	payload, err := gzstore.MarshalProto(d)
	if err != nil {
		return err
	}
	if len(payload) > gzstore.MaxStoredBytes {
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
	if err := gzstore.UnmarshalProto(sd.DeckGz, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// List reads the user's decks, newest first, from the flat fields alone.
// A listed deck carries its id, name, session, format, cost, legality
// date, and staleness, and no cards: Get reads the whole deck. limit of
// 0 reads them all.
func (r *Repo) List(ctx context.Context, uid string, limit int) ([]*mtgv1.Deck, error) {
	q := r.col(uid).Select(listFields...).OrderBy("created_at", firestore.Desc)
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
			return nil, fmt.Errorf("deck %s: %w", snap.Ref.ID, err)
		}
		out = append(out, storedToProto(snap.Ref.ID, sd))
	}
}

// storedToProto builds the list view of a deck from its flat fields.
func storedToProto(id string, sd storedDeck) *mtgv1.Deck {
	return &mtgv1.Deck{
		Id:           id,
		Name:         sd.Name,
		SessionId:    sd.SessionID,
		Format:       &mtgv1.Format{Id: mtgv1.FormatId(sd.FormatID)}, //nolint:gosec // FormatID was an enum at Put
		LegalityAsOf: sd.LegalityAsOf,
		BuyCostUsd:   sd.BuyCostUSD,
		Stale:        sd.Stale,
		CreatedAt:    timestamppb.New(sd.CreatedAt),
	}
}
