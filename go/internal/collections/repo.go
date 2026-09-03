package collections

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
)

// Repo stores collections in Firestore under
// users/{uid}/collections/{id}. One document per collection (D-16):
// metadata, a gzip JSON entry array, and a gzip count map. A 5,000-card
// binder stays far under the 1 MiB document limit. Put refuses a payload
// near that limit (ErrTooLarge). Sharding across documents is a later
// step.
type Repo struct {
	client *firestore.Client
	// index answers the card index of the day, or nil before the first
	// snapshot loads. A document stored before the summary existed
	// computes one from its entries, and the type counts need the
	// index (D-398).
	index func() *cards.Index
}

// ErrTooLarge reports a collection that does not fit one document.
var ErrTooLarge = errors.New("collection too large for one document (max 900 KiB gzip): sharding is a later step")

// NewRepo wraps a Firestore client. The caller owns the client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// WithIndex names where the repo reads the card index of the day. The
// summary of a document stored before schema version 2 reads it.
func (r *Repo) WithIndex(index func() *cards.Index) *Repo {
	r.index = index
	return r
}

// cardSource is the index as a CardSource, or nil when none is loaded.
// A nil *cards.Index must not reach the interface, because a typed nil
// is not a nil interface.
func (r *Repo) cardSource() CardSource {
	if r.index == nil {
		return nil
	}
	if idx := r.index(); idx != nil {
		return idx
	}
	return nil
}

// storedCollection is the Firestore document shape.
type storedCollection struct {
	Name        string    `firestore:"name"`
	Source      string    `firestore:"source"`
	ImportedAt  time.Time `firestore:"imported_at"`
	ContentHash string    `firestore:"content_hash"`
	CardCount   int64     `firestore:"card_count"`
	EntryCount  int64     `firestore:"entry_count"`
	EntriesGz   []byte    `firestore:"entries_gz"`
	OracleCntGz []byte    `firestore:"oracle_counts_gz"`
	// SummaryGz holds the counts the binder head shows, so the head reads
	// no entry (D-392). A collection stored before schema version 2 holds
	// none, and GetHead computes it from the entries that one time.
	SummaryGz     []byte `firestore:"summary_gz"`
	SchemaVersion int64  `firestore:"schema_version"`
}

// schemaVersion 2 adds the stored summary (D-392). A version 1 document
// still reads: its summary is absent, and the head computes one.
const schemaVersion = 2

func (r *Repo) doc(uid, id string) *firestore.DocumentRef {
	return r.client.Collection("users").Doc(uid).Collection("collections").Doc(id)
}

// Put stores a collection and returns its id. A collection with no id
// takes the id its content hash derives, so two identical uploads that
// race land on one document (D-16), and Put never needs a lookup first.
func (r *Repo) Put(ctx context.Context, uid string, col *mtgv1.Collection) (string, error) {
	entriesGz, err := gzstore.MarshalJSON(col.Entries)
	if err != nil {
		return "", err
	}
	countsGz, err := gzstore.MarshalJSON(OracleCounts(col.Entries))
	if err != nil {
		return "", err
	}
	summaryGz, err := gzstore.MarshalJSON(col.GetSummary())
	if err != nil {
		return "", err
	}
	if size := len(entriesGz) + len(countsGz) + len(summaryGz); size > gzstore.MaxStoredBytes {
		return "", fmt.Errorf("%w: %d bytes", ErrTooLarge, size)
	}
	stored := storedCollection{
		Name:          col.Name,
		Source:        col.Source.String(),
		ImportedAt:    time.Now().UTC(),
		ContentHash:   col.ContentHash,
		CardCount:     int64(col.CardCount),
		EntryCount:    int64(len(col.Entries)),
		EntriesGz:     entriesGz,
		OracleCntGz:   countsGz,
		SummaryGz:     summaryGz,
		SchemaVersion: schemaVersion,
	}
	ref := r.client.Collection("users").Doc(uid).Collection("collections").NewDoc()
	switch {
	case col.Id != "":
		ref = r.doc(uid, col.Id)
	case col.ContentHash != "":
		// The derived id belongs to its hash alone (D-399). A Replace
		// keeps an id whose hash moved on, so a later upload of the old
		// file must not land on it. Create refuses an existing document,
		// and the upload then takes a fresh id unless the document holds
		// this same hash, which is the identical re-upload of D-16.
		derived := r.doc(uid, DocID(col.ContentHash))
		_, err := derived.Create(ctx, stored)
		if err == nil {
			return derived.ID, nil
		}
		if status.Code(err) != codes.AlreadyExists {
			return "", fmt.Errorf("store collection: %w", err)
		}
		snap, err := derived.Get(ctx)
		if err != nil {
			return "", fmt.Errorf("store collection: %w", err)
		}
		var have storedCollection
		if err := snap.DataTo(&have); err != nil {
			return "", fmt.Errorf("collection %s: %w", derived.ID, err)
		}
		if have.ContentHash == col.ContentHash {
			ref = derived
		}
	}
	if _, err := ref.Set(ctx, stored); err != nil {
		return "", fmt.Errorf("store collection: %w", err)
	}
	return ref.ID, nil
}

// GetHead loads one collection without its entries (D-392). The binder
// head reads it, and it moves no megabyte.
//
// A document stored before schema version 2 holds no summary. This
// reads its entries that one time and computes one, so an old
// collection still draws a head. A later Put stores the summary.
func (r *Repo) GetHead(ctx context.Context, uid, id string) (*mtgv1.Collection, error) {
	snap, err := r.doc(uid, id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var stored storedCollection
	if err := snap.DataTo(&stored); err != nil {
		return nil, fmt.Errorf("collection %s: %w", id, err)
	}
	col := storedToProto(id, stored)
	if len(stored.SummaryGz) > 0 {
		var sum mtgv1.CollectionSummary
		if err := gzstore.UnmarshalJSON(stored.SummaryGz, &sum); err != nil {
			return nil, fmt.Errorf("collection %s summary: %w", id, err)
		}
		col.Summary = &sum
		return col, nil
	}
	var entries []*mtgv1.CollectionEntry
	if err := gzstore.UnmarshalJSON(stored.EntriesGz, &entries); err != nil {
		return nil, fmt.Errorf("collection %s entries: %w", id, err)
	}
	col.Summary = Summarize(entries, r.cardSource())
	return col, nil
}

// Get loads one collection with its entries.
func (r *Repo) Get(ctx context.Context, uid, id string) (*mtgv1.Collection, error) {
	snap, err := r.doc(uid, id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var stored storedCollection
	if err := snap.DataTo(&stored); err != nil {
		return nil, fmt.Errorf("collection %s: %w", id, err)
	}
	col := storedToProto(id, stored)
	var entries []*mtgv1.CollectionEntry
	if err := gzstore.UnmarshalJSON(stored.EntriesGz, &entries); err != nil {
		return nil, fmt.Errorf("collection %s entries: %w", id, err)
	}
	col.Entries = entries
	if len(stored.SummaryGz) > 0 {
		var sum mtgv1.CollectionSummary
		if err := gzstore.UnmarshalJSON(stored.SummaryGz, &sum); err != nil {
			return nil, fmt.Errorf("collection %s summary: %w", id, err)
		}
		col.Summary = &sum
	} else {
		col.Summary = Summarize(entries, r.cardSource())
	}
	return col, nil
}

// OwnedPrintings maps each Oracle id of a collection to the printing ids
// the user holds of it (D-299).
func (r *Repo) OwnedPrintings(ctx context.Context, uid, id string) (map[string][]string, error) {
	col, err := r.Get(ctx, uid, id)
	if err != nil {
		return nil, err
	}
	return OwnedPrintings(col.GetEntries()), nil
}

// PrintingCounts maps each Scryfall id of a collection to the copies the
// user holds of it (D-408). It reads the entries, so a caller asks for it
// on a turn that needs the precon ownership check and not on every turn.
func (r *Repo) PrintingCounts(ctx context.Context, uid, id string) (map[string]int32, error) {
	col, err := r.Get(ctx, uid, id)
	if err != nil {
		return nil, err
	}
	return PrintingCounts(col.GetEntries()), nil
}

// OracleCounts reads the stored per-Oracle-id count map of one collection.
// It inflates only the count payload, not the entries (D-37 ownership
// check in DeckService.Validate).
func (r *Repo) OracleCounts(ctx context.Context, uid, id string) (map[string]int32, error) {
	snap, err := r.doc(uid, id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var stored storedCollection
	if err := snap.DataTo(&stored); err != nil {
		return nil, fmt.Errorf("collection %s: %w", id, err)
	}
	var counts map[string]int32
	if err := gzstore.UnmarshalJSON(stored.OracleCntGz, &counts); err != nil {
		return nil, fmt.Errorf("collection %s counts: %w", id, err)
	}
	return counts, nil
}

// Delete removes one collection for good (D-347). A deck built from it
// keeps every card it holds, and its chat builds from the whole card
// database from then on.
func (r *Repo) Delete(ctx context.Context, uid, id string) error {
	if _, err := r.doc(uid, id).Get(ctx); err != nil {
		return err
	}
	_, err := r.doc(uid, id).Delete(ctx)
	return err
}

// Rename writes the name of one collection and returns it without its
// entries. It touches one field, so a rename never rewrites the
// megabyte the entries hold, and it keeps the import time: a rename must
// not move the collection up a list ordered by date.
func (r *Repo) Rename(ctx context.Context, uid, id, name string) (*mtgv1.Collection, error) {
	// Update refuses a missing document with NotFound, so no read comes
	// first. The status stays readable through the wrap.
	if _, err := r.doc(uid, id).Update(ctx, []firestore.Update{{Path: "name", Value: name}}); err != nil {
		return nil, fmt.Errorf("rename collection %s: %w", id, err)
	}
	return r.GetHead(ctx, uid, id)
}

// Replace writes new entries into one collection and keeps its id, so
// every deck and chat that names it still works (D-393). The content
// hash, the counts, and the summary all follow the new entries.
func (r *Repo) Replace(ctx context.Context, uid, id string, col *mtgv1.Collection) (*mtgv1.Collection, error) {
	ref := r.doc(uid, id)
	snap, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	var stored storedCollection
	if err := snap.DataTo(&stored); err != nil {
		return nil, fmt.Errorf("collection %s: %w", id, err)
	}
	// The name is the reader's, and a replacement of the cards does not
	// rename their binder.
	col.Id, col.Name = id, stored.Name
	if _, err := r.Put(ctx, uid, col); err != nil {
		return nil, err
	}
	return r.GetHead(ctx, uid, id)
}

// List returns every collection of a user, without entries.
func (r *Repo) List(ctx context.Context, uid string) ([]*mtgv1.Collection, error) {
	snaps, err := r.client.Collection("users").Doc(uid).Collection("collections").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	out := make([]*mtgv1.Collection, 0, len(snaps))
	for _, snap := range snaps {
		var stored storedCollection
		if err := snap.DataTo(&stored); err != nil {
			return nil, fmt.Errorf("collection %s: %w", snap.Ref.ID, err)
		}
		out = append(out, storedToProto(snap.Ref.ID, stored))
	}
	return out, nil
}

// FindByHash returns the id of a stored collection with this content
// hash, or "" when none exists. It detects an identical re-upload
// (D-16). A document stored under the derived id answers by that id, and
// one stored under a random id answers the hash query.
func (r *Repo) FindByHash(ctx context.Context, uid, hash string) (string, error) {
	if hash == "" {
		return "", nil
	}
	snaps, err := r.client.Collection("users").Doc(uid).Collection("collections").
		Where("content_hash", "==", hash).Limit(1).Documents(ctx).GetAll()
	if err != nil {
		return "", err
	}
	if len(snaps) == 0 {
		return "", nil
	}
	return snaps[0].Ref.ID, nil
}

// DocID derives the document id of a collection from its content hash.
// The hash is 64 hex characters, which fits the Firestore id rule.
func DocID(contentHash string) string { return "h-" + contentHash }

func storedToProto(id string, s storedCollection) *mtgv1.Collection {
	return &mtgv1.Collection{
		Id:          id,
		Name:        s.Name,
		Source:      mtgv1.ImportSource(mtgv1.ImportSource_value[s.Source]),
		ContentHash: s.ContentHash,
		CardCount:   int32(s.CardCount), //nolint:gosec // CardCount was an int32 at Put
		ImportedAt:  timestamppb.New(s.ImportedAt),
	}
}
