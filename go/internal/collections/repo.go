package collections

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/gzstore"
)

// Repo stores collections in Firestore under
// users/{uid}/collections/{id}. One document per collection (D-16):
// metadata, a gzip JSON entry array, a gzip count map, and the summary.
// A payload over the size of one document goes into parts under
// {id}/parts, and the collection document keeps the summary and the
// count of each kind of part (D-1110). Put refuses a payload over
// MaxPayloadBytes (ErrTooLarge).
type Repo struct {
	client *firestore.Client
	// index answers the card index of the day, or nil before the first
	// snapshot loads. A document stored before the summary existed
	// computes one from its entries, and the type counts need the
	// index (D-398).
	index func() *cards.Index
}

// MaxEntries is the stated limit of one collection: the distinct rows
// by printing, finish, condition, and language (REV-026, D-1110). A
// measure of real printings stored 50,000 rows in 4,146,949 gzip bytes,
// and their entries inflated to 13,906,176 bytes.
const MaxEntries = 50000

// MaxPayloadBytes bounds the gzip payload of one collection. Put writes
// the document and every part in one commit, and Firestore caps one
// request at 10 MiB.
const MaxPayloadBytes = 8 << 20

// maxInflatedBytes bounds the inflated read of one payload of a
// collection. The entries of MaxEntries rows inflate past the shared
// limit of gzstore.
const maxInflatedBytes = 32 << 20

// ErrTooManyEntries refuses a collection over MaxEntries before the
// write.
var ErrTooManyEntries = fmt.Errorf("a collection holds at most %d distinct rows (by printing, finish, condition, and language)", MaxEntries)

// ErrTooLarge reports a collection over MaxPayloadBytes.
var ErrTooLarge = fmt.Errorf("the collection is too large to store: a collection holds at most about %d distinct rows", MaxEntries)

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
	SummaryGz []byte `firestore:"summary_gz"`
	// EntryParts and CountParts count the parts that hold the entries and
	// the counts of a payload over one document (D-1110). Zero means the
	// payload is in EntriesGz and OracleCntGz.
	EntryParts    int64 `firestore:"entry_parts"`
	CountParts    int64 `firestore:"count_parts"`
	SchemaVersion int64 `firestore:"schema_version"`
}

// storedPart is one part of a payload over one document (D-1110).
type storedPart struct {
	Data []byte `firestore:"data"`
}

// schemaVersion 2 adds the stored summary (D-392). Version 3 adds the
// parts (D-1110). A version 1 or 2 document still reads: it holds no
// part, and a version 1 head computes its summary.
const schemaVersion = 3

// Part ids: "e" and the index for the entries, "c" and the index for
// the counts.
const (
	partsName     = "parts"
	entryPartKind = "e"
	countPartKind = "c"
)

func partRef(col *firestore.DocumentRef, kind string, i int) *firestore.DocumentRef {
	return col.Collection(partsName).Doc(fmt.Sprintf("%s%d", kind, i))
}

// split cuts a payload into parts of at most gzstore.MaxStoredBytes.
func split(payload []byte) [][]byte {
	var out [][]byte
	for len(payload) > gzstore.MaxStoredBytes {
		out = append(out, payload[:gzstore.MaxStoredBytes])
		payload = payload[gzstore.MaxStoredBytes:]
	}
	return append(out, payload)
}

func (r *Repo) doc(uid, id string) *firestore.DocumentRef {
	return r.client.Collection("users").Doc(uid).Collection("collections").Doc(id)
}

// Put stores a collection and returns its id. A collection with no id
// takes the id its content hash derives, so two identical uploads that
// race land on one document (D-16).
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
	size := len(entriesGz) + len(countsGz) + len(summaryGz)
	if size > MaxPayloadBytes {
		return "", fmt.Errorf("%w: %d bytes", ErrTooLarge, size)
	}
	stored := storedCollection{
		Name:          col.Name,
		Source:        col.Source.String(),
		ImportedAt:    time.Now().UTC(),
		ContentHash:   col.ContentHash,
		CardCount:     int64(col.CardCount),
		EntryCount:    int64(len(col.Entries)),
		SummaryGz:     summaryGz,
		SchemaVersion: schemaVersion,
	}
	var entryParts, countParts [][]byte
	if size <= gzstore.MaxStoredBytes {
		stored.EntriesGz, stored.OracleCntGz = entriesGz, countsGz
	} else {
		entryParts, countParts = split(entriesGz), split(countsGz)
		stored.EntryParts, stored.CountParts = int64(len(entryParts)), int64(len(countParts))
	}
	// One transaction writes the document and its parts, and deletes each
	// older part, so no read sees a part of another file (D-1110).
	var id string
	err = r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		ref, err := r.target(tx, uid, col)
		if err != nil {
			return err
		}
		old, err := tx.Documents(ref.Collection(partsName)).GetAll()
		if err != nil {
			return err
		}
		if err := tx.Set(ref, stored); err != nil {
			return err
		}
		keep := map[string]bool{}
		for kind, parts := range map[string][][]byte{entryPartKind: entryParts, countPartKind: countParts} {
			for i, data := range parts {
				part := partRef(ref, kind, i)
				keep[part.ID] = true
				if err := tx.Set(part, storedPart{Data: data}); err != nil {
					return err
				}
			}
		}
		for _, snap := range old {
			if !keep[snap.Ref.ID] {
				if err := tx.Delete(snap.Ref); err != nil {
					return err
				}
			}
		}
		id = ref.ID
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("store collection: %w", err)
	}
	return id, nil
}

// target names the document Put writes. A collection with no id takes
// the id its content hash derives, so two identical uploads that race
// land on one document (D-16).
//
// The derived id belongs to its hash alone (D-399). A Replace keeps an
// id whose hash moved on, so a later upload of the old file must not
// land on it. The upload takes a fresh id unless the derived document is
// absent or holds this same hash, which is the identical re-upload of
// D-16.
func (r *Repo) target(tx *firestore.Transaction, uid string, col *mtgv1.Collection) (*firestore.DocumentRef, error) {
	if col.Id != "" {
		return r.doc(uid, col.Id), nil
	}
	fresh := r.client.Collection("users").Doc(uid).Collection("collections").NewDoc()
	if col.ContentHash == "" {
		return fresh, nil
	}
	derived := r.doc(uid, DocID(col.ContentHash))
	snap, err := tx.Get(derived)
	if status.Code(err) == codes.NotFound {
		return derived, nil
	}
	if err != nil {
		return nil, err
	}
	var have storedCollection
	if err := snap.DataTo(&have); err != nil {
		return nil, fmt.Errorf("collection %s: %w", derived.ID, err)
	}
	if have.ContentHash == col.ContentHash {
		return derived, nil
	}
	return fresh, nil
}

// GetHead loads one collection without its entries (D-392). The binder
// head reads it, and it moves no megabyte.
//
// A document stored before schema version 2 holds no summary. This
// reads its entries that one time and computes one, so an old
// collection still draws a head. A later Put stores the summary.
func (r *Repo) GetHead(ctx context.Context, uid, id string) (*mtgv1.Collection, error) {
	stored, _, _, err := r.load(ctx, uid, id, false, false)
	if err != nil {
		return nil, err
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
	_, entriesGz, _, err := r.load(ctx, uid, id, true, false)
	if err != nil {
		return nil, err
	}
	var entries []*mtgv1.CollectionEntry
	if err := gzstore.UnmarshalJSONMax(entriesGz, &entries, maxInflatedBytes); err != nil {
		return nil, fmt.Errorf("collection %s entries: %w", id, err)
	}
	col.Summary = Summarize(entries, r.cardSource())
	return col, nil
}

// Get loads one collection with its entries.
func (r *Repo) Get(ctx context.Context, uid, id string) (*mtgv1.Collection, error) {
	stored, entriesGz, _, err := r.load(ctx, uid, id, true, false)
	if err != nil {
		return nil, err
	}
	col := storedToProto(id, stored)
	var entries []*mtgv1.CollectionEntry
	if err := gzstore.UnmarshalJSONMax(entriesGz, &entries, maxInflatedBytes); err != nil {
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

// load reads one collection document, and the entries or the counts
// payload when asked. A payload in parts reads in a read-only
// transaction with the document again, so a Put between the two reads
// can not mix two files (D-1110).
func (r *Repo) load(ctx context.Context, uid, id string, entries, counts bool) (storedCollection, []byte, []byte, error) {
	ref := r.doc(uid, id)
	var stored storedCollection
	snap, err := ref.Get(ctx)
	if err != nil {
		return stored, nil, nil, err
	}
	if err := snap.DataTo(&stored); err != nil {
		return stored, nil, nil, fmt.Errorf("collection %s: %w", id, err)
	}
	inParts := (entries && stored.EntryParts > 0) || (counts && stored.CountParts > 0)
	if !inParts {
		return stored, stored.EntriesGz, stored.OracleCntGz, nil
	}
	var entriesGz, countsGz []byte
	err = r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		snap, err := tx.Get(ref)
		if err != nil {
			return err
		}
		stored = storedCollection{}
		if err := snap.DataTo(&stored); err != nil {
			return fmt.Errorf("collection %s: %w", id, err)
		}
		entriesGz, countsGz = stored.EntriesGz, stored.OracleCntGz
		if entries && stored.EntryParts > 0 {
			if entriesGz, err = readParts(tx, ref, entryPartKind, stored.EntryParts); err != nil {
				return err
			}
		}
		if counts && stored.CountParts > 0 {
			if countsGz, err = readParts(tx, ref, countPartKind, stored.CountParts); err != nil {
				return err
			}
		}
		return nil
	}, firestore.ReadOnly)
	return stored, entriesGz, countsGz, err
}

// readParts joins the n parts of one kind, in order.
func readParts(tx *firestore.Transaction, col *firestore.DocumentRef, kind string, n int64) ([]byte, error) {
	refs := make([]*firestore.DocumentRef, n)
	for i := range refs {
		refs[i] = partRef(col, kind, i)
	}
	snaps, err := tx.GetAll(refs)
	if err != nil {
		return nil, err
	}
	var out []byte
	for i, snap := range snaps {
		if !snap.Exists() {
			return nil, fmt.Errorf("collection %s: part %s is missing", col.ID, refs[i].ID)
		}
		var part storedPart
		if err := snap.DataTo(&part); err != nil {
			return nil, fmt.Errorf("collection %s part %s: %w", col.ID, refs[i].ID, err)
		}
		out = append(out, part.Data...)
	}
	return out, nil
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
	_, _, countsGz, err := r.load(ctx, uid, id, false, true)
	if err != nil {
		return nil, err
	}
	var counts map[string]int32
	if err := gzstore.UnmarshalJSONMax(countsGz, &counts, maxInflatedBytes); err != nil {
		return nil, fmt.Errorf("collection %s counts: %w", id, err)
	}
	return counts, nil
}

// Delete removes one collection for good (D-347). A deck built from it
// keeps every card it holds, and its chat builds from the whole card
// database from then on. Its parts go in the same transaction (D-1110).
func (r *Repo) Delete(ctx context.Context, uid, id string) error {
	ref := r.doc(uid, id)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		if _, err := tx.Get(ref); err != nil {
			return err
		}
		parts, err := tx.Documents(ref.Collection(partsName)).GetAll()
		if err != nil {
			return err
		}
		for _, snap := range parts {
			if err := tx.Delete(snap.Ref); err != nil {
				return err
			}
		}
		return tx.Delete(ref)
	})
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
