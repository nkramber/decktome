package collections

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/firestore"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// Repo stores collections in Firestore under
// users/{uid}/collections/{id}. One document per collection (D-16,
// OQ-2): metadata, a gzip JSON entry array, and a gzip count map.
// A 5,000-card binder stays far under the 1 MiB document limit.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client. The caller owns the client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// storedCollection is the Firestore document shape.
type storedCollection struct {
	Name          string    `firestore:"name"`
	Source        string    `firestore:"source"`
	ImportedAt    time.Time `firestore:"imported_at"`
	ContentHash   string    `firestore:"content_hash"`
	CardCount     int64     `firestore:"card_count"`
	EntryCount    int64     `firestore:"entry_count"`
	EntriesGz     []byte    `firestore:"entries_gz"`
	OracleCntGz   []byte    `firestore:"oracle_counts_gz"`
	SchemaVersion int64     `firestore:"schema_version"`
}

const schemaVersion = 1

func (r *Repo) doc(uid, id string) *firestore.DocumentRef {
	return r.client.Collection("users").Doc(uid).Collection("collections").Doc(id)
}

// Put stores a collection and returns its id.
func (r *Repo) Put(ctx context.Context, uid string, col *mtgv1.Collection) (string, error) {
	entriesGz, err := gzJSON(col.Entries)
	if err != nil {
		return "", err
	}
	countsGz, err := gzJSON(OracleCounts(col.Entries))
	if err != nil {
		return "", err
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
		SchemaVersion: schemaVersion,
	}
	ref := r.client.Collection("users").Doc(uid).Collection("collections").NewDoc()
	if col.Id != "" {
		ref = r.doc(uid, col.Id)
	}
	if _, err := ref.Set(ctx, stored); err != nil {
		return "", fmt.Errorf("store collection: %w", err)
	}
	return ref.ID, nil
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
	if err := ungzJSON(stored.EntriesGz, &entries); err != nil {
		return nil, fmt.Errorf("collection %s entries: %w", id, err)
	}
	col.Entries = entries
	return col, nil
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
// hash, or "" when none exists. It detects an identical re-upload (D-16).
func (r *Repo) FindByHash(ctx context.Context, uid, hash string) (string, error) {
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

func storedToProto(id string, s storedCollection) *mtgv1.Collection {
	return &mtgv1.Collection{
		Id:          id,
		Name:        s.Name,
		Source:      mtgv1.ImportSource(mtgv1.ImportSource_value[s.Source]),
		ContentHash: s.ContentHash,
		CardCount:   int32(s.CardCount),
	}
}

func gzJSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if err := json.NewEncoder(gz).Encode(v); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ungzJSON(data []byte, v any) error {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer func() { _ = gz.Close() }()
	raw, err := io.ReadAll(gz)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)
}
