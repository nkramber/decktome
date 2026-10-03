package main

import (
	"context"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"time"

	"cloud.google.com/go/firestore"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/push"
	"github.com/nkramber/decktome/go/internal/rules"
	"github.com/nkramber/decktome/go/internal/stale"
)

// stalePassTimeout bounds one stale pass. A pass that stops leaves the
// marker without an end, so the next run of the job runs it again.
const stalePassTimeout = 10 * time.Minute

// openDecks opens the deck store of the pass, and the push of the decks
// it hits. The job opens Firestore only when a pass waits, so an hour
// with no legality change reads no deck. The account of the job holds
// roles/datastore.user (D-1018) and roles/firebasecloudmessaging.admin
// (D-1089). A nil notify sends no push.
type openDecks func(ctx context.Context) (stale.Store, notifyStale, func(), error)

// notifyStale tells one user about the decks that one pass hit (D-1088).
type notifyStale func(ctx context.Context, uid string, decks []*mtgv1.Deck)

// pushNotifyTimeout bounds the push after a pass. The push runs on its
// own context, so a pass that ran out of time still sends its hits.
const pushNotifyTimeout = time.Minute

// firestoreDecks opens the deck store on the Firestore of the project. A
// nil sender sends no push, as in the local stack.
func firestoreDecks(project string, sender push.Sender, logger *slog.Logger) openDecks {
	return func(ctx context.Context) (stale.Store, notifyStale, func(), error) {
		client, err := firestore.NewClient(ctx, project)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("firestore: %w", err)
		}
		var notify notifyStale
		if sender != nil {
			notify = push.NewNotifier(push.NewRepo(client), sender, logger).DecksStale
		}
		return decks.NewRepo(client), notify, func() { _ = client.Close() }, nil
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
	deckStore, notify, closeStore, err := open(ctx)
	if err != nil {
		return err
	}
	defer closeStore()
	res, err := stale.Pass(ctx, deckStore, keyOf, stale.Legalities(legal), logger)
	// The marks of the hit decks are stored, so a later pass does not hit
	// them again. Their push goes out now, also after a failed pass.
	if notify != nil && len(res.Hit) > 0 {
		pctx, pcancel := context.WithTimeout(context.WithoutCancel(ctx), pushNotifyTimeout)
		for _, uid := range slices.Sorted(maps.Keys(res.Hit)) {
			notify(pctx, uid, res.Hit[uid])
		}
		pcancel()
	}
	if err != nil {
		return fmt.Errorf("stale pass for %s: %w", diff, err)
	}
	rec.StalePass = now().UTC().Format(time.RFC3339)
	rec.StaleDecks = res.Stale
	if err := store.WriteLegalityDiff(ctx, diff, rec); err != nil {
		return fmt.Errorf("write legality marker %s: %w", diff, err)
	}
	logger.Info("stale pass ended", "diff", diff, "legalities", latest,
		"read", res.Read, "stale", res.Stale, "written", res.Written, "users hit", len(res.Hit))
	return nil
}
