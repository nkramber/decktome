package cards

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/scryfall"
)

// bulkTypeByFile maps a snapshot file to its Scryfall bulk type.
var bulkTypeByFile = map[string]string{
	"oracle_cards.jsonl.gz":  "oracle_cards",
	"default_cards.jsonl.gz": "default_cards",
	"oracle_tags.jsonl.gz":   "oracle_tags",
}

// Refresh downloads the bulk files into the store when Scryfall has a
// newer oracle_cards file than the newest stored version.
// It returns the current version either way.
func Refresh(ctx context.Context, client *scryfall.Client, store Store, logger *slog.Logger) (string, error) {
	files, err := client.BulkFiles(ctx)
	if err != nil {
		return "", err
	}
	oracle, ok := files["oracle_cards"]
	if !ok {
		return "", fmt.Errorf("scryfall bulk catalog has no oracle_cards")
	}
	remote := VersionFor(oracle.UpdatedAt)
	current, err := store.LatestVersion(ctx)
	if err != nil {
		return "", err
	}
	if current >= remote {
		logger.Info("cards refresh: snapshot current", "version", current)
		return current, nil
	}
	logger.Info("cards refresh: downloading", "from", current, "to", remote)
	for _, name := range SnapshotFiles {
		bulk, ok := files[bulkTypeByFile[name]]
		if !ok {
			return "", fmt.Errorf("scryfall bulk catalog has no %s", bulkTypeByFile[name])
		}
		if err := copyBulk(ctx, client, store, remote, name, bulk.DownloadURI); err != nil {
			return "", err
		}
	}
	if err := store.Finalize(ctx, remote); err != nil {
		return "", fmt.Errorf("finalize %s: %w", remote, err)
	}
	logger.Info("cards refresh: done", "version", remote)
	return remote, nil
}

func copyBulk(ctx context.Context, client *scryfall.Client, store Store, version, name, uri string) error {
	body, err := client.Download(ctx, uri)
	if err != nil {
		return err
	}
	defer func() { _ = body.Close() }()
	w, err := store.Create(ctx, version, name)
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, body); err != nil {
		_ = w.Close()
		return fmt.Errorf("store %s/%s: %w", version, name, err)
	}
	return w.Close()
}

// LoadIndex builds an Index from the newest stored snapshot.
// It returns nil with no error when the store is empty.
func LoadIndex(ctx context.Context, store Store, logger *slog.Logger) (*Index, error) {
	version, err := store.LatestVersion(ctx)
	if err != nil {
		return nil, err
	}
	if version == "" {
		return nil, nil
	}
	start := time.Now()
	asOf, err := VersionTime(version)
	if err != nil {
		return nil, fmt.Errorf("bad snapshot version %q: %w", version, err)
	}
	oc, err := store.Open(ctx, version, "oracle_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	defer func() { _ = oc.Close() }()
	parsed, err := LoadCards(oc, "oracle_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	dc, err := store.Open(ctx, version, "default_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	defer func() { _ = dc.Close() }()
	printings, err := LoadPrintings(dc, "default_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	tf, err := store.Open(ctx, version, "oracle_tags.jsonl.gz")
	if err != nil {
		return nil, err
	}
	defer func() { _ = tf.Close() }()
	tags, err := LoadTags(tf, "oracle_tags.jsonl.gz")
	if err != nil {
		return nil, err
	}
	idx := NewIndex(parsed, printings, tags, asOf)
	logger.Info("cards index loaded", "version", version, "cards", idx.Len(),
		"printings", len(printings), "tags", tags.Len(), "took", time.Since(start).String())
	return idx, nil
}
