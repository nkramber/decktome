package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"cloud.google.com/go/firestore"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/rules"
	"github.com/nkramber/decktome/go/internal/stale"
)

// stalePassTimeout bounds one stale pass. A pass that stops leaves the
// marker without an end, so the next run of the job runs it again.
const stalePassTimeout = 10 * time.Minute

// openDecks opens the deck store of the pass. The job opens Firestore
// only when a pass waits, so an hour with no legality change reads no
// deck. The account of the job holds roles/datastore.user (D-1018).
type openDecks func(ctx context.Context) (stale.Store, func(), error)

// firestoreDecks opens the deck store on the Firestore of the project.
func firestoreDecks(project string) openDecks {
	return func(ctx context.Context) (stale.Store, func(), error) {
		client, err := firestore.NewClient(ctx, project)
		if err != nil {
			return nil, nil, fmt.Errorf("firestore: %w", err)
		}
		return decks.NewRepo(client), func() { _ = client.Close() }, nil
	}
}

// formatKeys maps each format of the app to its Scryfall legality key,
// from the rules config. A house format has no key, so the pass reads
// no legality for it (D-3).
func formatKeys() (func(mtgv1.FormatId) string, error) {
	cfg, err := rules.Load()
	if err != nil {
		return nil, err
	}
	return func(f mtgv1.FormatId) string { return cfg.Formats[f.String()].ScryfallKey }, nil
}

// stalePass runs the pass of I-1 when the newest legality diff has no
// ended pass (D-29, D-1018). It reads the legalities of the newest
// version, which holds that change and each later one. The end of the
// pass goes into the marker of the diff.
func stalePass(ctx context.Context, store cards.Store, open openDecks, logger *slog.Logger, now func() time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, stalePassTimeout)
	defer cancel()
	diff, rec, ok, err := cards.LatestLegalityDiff(ctx, store)
	if err != nil {
		return fmt.Errorf("read legality markers: %w", err)
	}
	if !ok || rec.StalePass != "" {
		return nil
	}
	latest, err := store.LatestVersion(ctx)
	if err != nil {
		return err
	}
	legal, err := cards.LoadLegalities(ctx, store, latest)
	if err != nil {
		return fmt.Errorf("legalities %s: %w", latest, err)
	}
	keyOf, err := formatKeys()
	if err != nil {
		return fmt.Errorf("rules config: %w", err)
	}
	deckStore, closeStore, err := open(ctx)
	if err != nil {
		return err
	}
	defer closeStore()
	res, err := stale.Pass(ctx, deckStore, keyOf, stale.Legalities(legal), logger)
	if err != nil {
		return fmt.Errorf("stale pass for %s: %w", diff, err)
	}
	rec.StalePass = now().UTC().Format(time.RFC3339)
	rec.StaleDecks = res.Stale
	if err := store.WriteLegalityDiff(ctx, diff, rec); err != nil {
		return fmt.Errorf("write legality marker %s: %w", diff, err)
	}
	logger.Info("stale pass ended", "diff", diff, "legalities", latest,
		"read", res.Read, "stale", res.Stale, "written", res.Written)
	return nil
}
