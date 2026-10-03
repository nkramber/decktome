package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/newcards"
	"github.com/nkramber/decktome/go/internal/push"
	"github.com/nkramber/decktome/go/internal/sessions"
)

// newCardsPassTimeout bounds one new-cards pass. A pass that stops leaves
// the marker without an end, so the next run of the job runs it again.
const newCardsPassTimeout = 10 * time.Minute

// themeOf reads the theme of the chat of a deck, or "" when the chat is
// gone (D-1091).
type themeOf func(ctx context.Context, uid, sessionID string) (string, error)

// openNewCards opens the deck store of the new-cards pass, the theme of
// a chat, and the push of the decks it hits. The account of the job
// reads the chats under the role of D-1018. A nil notify sends no push.
type openNewCards func(ctx context.Context) (newcards.Store, themeOf, notifyStale, func(), error)

// loadIndex loads the card index of one stored version.
type loadIndex func(ctx context.Context, version string) (*cards.Index, error)

// firestoreNewCards opens the stores of the pass on the Firestore of the
// project. A nil sender sends no push, as in the local stack.
func firestoreNewCards(project string, sender push.Sender, logger *slog.Logger) openNewCards {
	return func(ctx context.Context) (newcards.Store, themeOf, notifyStale, func(), error) {
		client, err := firestore.NewClient(ctx, project)
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("firestore: %w", err)
		}
		chats := sessions.NewRepo(client)
		theme := func(ctx context.Context, uid, id string) (string, error) {
			s, err := chats.Get(ctx, uid, id)
			if errors.Is(err, sessions.ErrNotFound) {
				return "", nil
			}
			if err != nil {
				return "", err
			}
			return s.GetSlots().GetTheme(), nil
		}
		var notify notifyStale
		if sender != nil {
			notify = push.NewNotifier(push.NewRepo(client), sender, logger).NewCards
		}
		return decks.NewRepo(client), theme, notify, func() { _ = client.Close() }, nil
	}
}

// newCardsPass runs the pass of D-1091 when the newest new-cards marker
// has no ended pass. It loads the whole index of the newest version only
// then, about once for each set. The end of the pass goes into the
// marker.
func newCardsPass(ctx context.Context, store cards.Store, load loadIndex, open openNewCards, fit newcards.Fitter,
	logger *slog.Logger, now func() time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, newCardsPassTimeout)
	defer cancel()
	marked, rec, ok, err := cards.LatestNewCards(ctx, store)
	if err != nil {
		return fmt.Errorf("read new-cards markers: %w", err)
	}
	if !ok || rec.Pass != "" {
		return nil
	}
	latest, err := store.LatestVersion(ctx)
	if err != nil {
		return err
	}
	idx, err := load(ctx, latest)
	if err != nil {
		return fmt.Errorf("index %s: %w", latest, err)
	}
	keyOf, err := formatKeys()
	if err != nil {
		return fmt.Errorf("rules config: %w", err)
	}
	deckStore, theme, notify, closeStore, err := open(ctx)
	if err != nil {
		return err
	}
	defer closeStore()
	res, err := newcards.Pass(ctx, deckStore, newcards.Input{
		Index: idx, Version: marked, New: rec.Cards, KeyOf: keyOf, ThemeOf: theme, Fit: fit, Floor: candidates.FitFloor,
	}, logger)
	// The cards of the hit decks are stored, so a later pass does not hit
	// them again. Their push goes out now, also after a failed pass.
	if notify != nil && len(res.Hit) > 0 {
		pctx, pcancel := context.WithTimeout(context.WithoutCancel(ctx), pushNotifyTimeout)
		for _, uid := range slices.Sorted(maps.Keys(res.Hit)) {
			notify(pctx, uid, res.Hit[uid])
		}
		pcancel()
	}
	if err != nil {
		return fmt.Errorf("new cards pass for %s: %w", marked, err)
	}
	rec.Pass = now().UTC().Format(time.RFC3339)
	rec.Decks = res.Fit
	if err := cards.WriteNewCards(ctx, store, marked, rec); err != nil {
		return fmt.Errorf("write new-cards marker %s: %w", marked, err)
	}
	logger.Info("new cards pass ended", "marker", marked, "index", latest, "new cards", len(rec.Cards),
		"read", res.Read, "fit", res.Fit, "written", res.Written, "users hit", len(res.Hit))
	return nil
}
