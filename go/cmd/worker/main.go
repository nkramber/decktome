// Command worker refreshes the Scryfall snapshot from the bulk files
// only.
//
// Production shape: a Cloud Run job. One Cloud Scheduler cron
// starts it with -once:
//
//	"*/15 * * * *"   every 15 minutes, every day
//
// A run exits 0 early ("skip") when the newest stored version is less
// than 60 minutes old and no announcement is pending. Otherwise it
// refreshes, then exits 0, or exits 1 on a failure. The announcement
// calendar and the store decide, not process memory.
//
// The loop mode (no -once) exists for make dev only. It polls on the
// calendar cadence in-process, because no Scheduler exists locally.
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gcpenv"
	"github.com/nkramber/mtg-deck-builder/go/internal/scryfall"
)

var version = "dev"

// refreshTimeout bounds one whole refresh: three downloads plus the
// legality diff.
const refreshTimeout = 90 * time.Minute

func main() {
	once := flag.Bool("once", false, "run one refresh and exit (Cloud Run job, make dev-seed)")
	flag.Parse()

	logger := gcpenv.NewLogger(os.Stdout)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := run(ctx, *once, logger)
	stop()
	if err != nil {
		logger.Error("worker failed", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, once bool, logger *slog.Logger) error {
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	// The worker creates the bucket in emulator mode, because
	// fake-gcs-server starts empty.
	store, storageClient, err := gcpenv.SnapshotStore(ctx, project, true, logger)
	if err != nil {
		return fmt.Errorf("snapshot store init: %w", err)
	}
	if storageClient != nil {
		defer func() { _ = storageClient.Close() }()
	}
	client := scryfall.New(nil, os.Getenv("SCRYFALL_BASE_URL"), logger)
	calendar, err := cards.Announcements()
	if err != nil {
		return fmt.Errorf("announcement calendar: %w", err)
	}

	// lastDiff is the AsOf of the last snapshot whose legalities changed.
	// It comes from the markers in the store, not from memory.
	lastDiff, err := cards.LastLegalityDiff(ctx, store)
	if err != nil {
		return fmt.Errorf("read legality markers: %w", err)
	}
	refresh := func() error {
		refreshCtx, cancel := context.WithTimeout(ctx, refreshTimeout)
		defer cancel()
		previous, err := store.LatestVersion(refreshCtx)
		if err != nil {
			return err
		}
		current, err := cards.Refresh(refreshCtx, client, store, logger)
		if err != nil {
			return err
		}
		// A snapshot stored before the set file existed gains it here,
		// so the set family resolves without a whole re-download
		// (D-377). A failure costs the families, not the cards, so it
		// logs and the cycle goes on.
		if _, err := cards.BackfillSets(refreshCtx, client, store, logger); err != nil {
			logger.Error("cards refresh: the set file could not be backfilled", "err", err)
		}
		if current == previous || previous == "" {
			return nil
		}
		asOf, err := cards.VersionTime(current)
		if err != nil {
			return err
		}
		// A snapshot covers an announcement only when a legality
		// changed. The calendar date is UTC midnight and the 09:00Z
		// snapshot on that day can predate the post.
		changed, err := cards.LegalityDiff(refreshCtx, store, previous, current)
		if err != nil {
			return fmt.Errorf("legality diff %s to %s: %w", previous, current, err)
		}
		logger.Info("legality_diff", "from", previous, "to", current, "changed_cards", changed)
		if changed == 0 {
			return nil
		}
		lastDiff = asOf
		rec := cards.LegalityDiffRecord{SnapshotAsOf: asOf.Format(time.RFC3339), ChangedCards: changed}
		if ann, hours, ok := calendar.LagHours(asOf); ok {
			// The hours from the announcement date to the snapshot
			// that changed legalities.
			rec.Announcement = ann.Format("2006-01-02")
			rec.LagHours = hours
			logger.Info("announcement_covered", "announcement", rec.Announcement,
				"lag_hours", hours, "snapshot", current, "changed_cards", changed)
		}
		if err := store.WriteLegalityDiff(refreshCtx, current, rec); err != nil {
			return fmt.Errorf("write legality marker %s: %w", current, err)
		}
		return nil
	}

	if once {
		now := time.Now().UTC()
		latest, err := store.LatestVersion(ctx)
		if err != nil {
			return err
		}
		var latestAt time.Time
		if latest != "" {
			if latestAt, err = cards.VersionTime(latest); err != nil {
				return err
			}
		}
		_, pending := calendar.Pending(now, lastDiff)
		if skipRefresh(now, latestAt, pending) {
			logger.Info("refresh skipped: snapshot fresh and no announcement pending", "version", latest)
			return nil
		}
		return refresh()
	}

	logger.Info("worker started in loop mode", "version", version)
	if err := refresh(); err != nil {
		logger.Error("refresh failed", "err", err)
	}
	timer := time.NewTimer(0)
	defer timer.Stop()
	for {
		interval := calendar.CheckInterval(time.Now().UTC(), lastDiff)
		logger.Debug("next refresh check", "in", interval.String())
		timer.Reset(interval)
		select {
		case <-ctx.Done():
			logger.Info("worker stopped")
			return nil
		case <-timer.C:
			if err := refresh(); err != nil {
				logger.Error("refresh failed", "err", err)
			}
		}
	}
}

// freshWindow is how long a stored snapshot counts as fresh for the
// -once skip decision.
const freshWindow = 60 * time.Minute

// skipRefresh decides the -once early exit. It skips only when the
// newest stored version is younger than freshWindow and no announcement
// waits for a legality diff. A zero latestAt means an empty store.
func skipRefresh(now, latestAt time.Time, pending bool) bool {
	if pending || latestAt.IsZero() {
		return false
	}
	return now.Sub(latestAt) < freshWindow
}
