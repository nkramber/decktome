// Command worker runs scheduled jobs. The first job is the daily
// Scryfall snapshot refresh (roadmap PR-2, F-3: bulk files only).
package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/storage"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/scryfall"
)

var version = "dev"

// The check cadence comes from the announcement calendar (PR-3, F-1):
// hourly on a normal day, every 15 minutes while an announcement is not
// yet covered by the stored snapshot.

func main() {
	once := flag.Bool("once", false, "run one refresh and exit (make dev-seed)")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	store, err := snapshotStore(ctx, logger)
	if err != nil {
		logger.Error("snapshot store init failed", "err", err)
		stop()
		os.Exit(1) //nolint:gocritic // deliberate: stop() already ran
	}
	client := scryfall.New(nil, os.Getenv("SCRYFALL_BASE_URL"))

	calendar, err := cards.Announcements()
	if err != nil {
		logger.Error("announcement calendar broken", "err", err)
		stop()
		os.Exit(1) //nolint:gocritic // deliberate: stop() already ran
	}

	var snapshotAsOf time.Time
	refresh := func() error {
		refreshCtx, cancel := context.WithTimeout(ctx, 30*time.Minute)
		defer cancel()
		version, err := cards.Refresh(refreshCtx, client, store, logger)
		if err != nil {
			return err
		}
		asOf, err := cards.VersionTime(version)
		if err != nil {
			return err
		}
		if asOf.After(snapshotAsOf) && !snapshotAsOf.IsZero() {
			// M-2: report the announcement lag when a new snapshot lands.
			if ann, hours, ok := calendar.LagHours(asOf); ok {
				logger.Info("legality_lag", "announcement", ann.Format("2006-01-02"), "lag_hours", hours, "snapshot", version)
			}
		}
		snapshotAsOf = asOf
		return nil
	}

	if *once {
		if err := refresh(); err != nil {
			logger.Error("refresh failed", "err", err)
			os.Exit(1)
		}
		return
	}

	logger.Info("worker started", "version", version)
	if err := refresh(); err != nil {
		logger.Error("refresh failed", "err", err)
	}
	for {
		interval := calendar.CheckInterval(time.Now().UTC(), snapshotAsOf)
		logger.Debug("next refresh check", "in", interval.String())
		select {
		case <-ctx.Done():
			logger.Info("worker stopped")
			return
		case <-time.After(interval):
			if err := refresh(); err != nil {
				logger.Error("refresh failed", "err", err)
			}
		}
	}
}

// snapshotStore mirrors the api's backend choice, and also makes sure the
// bucket exists. fake-gcs-server starts empty in local mode.
func snapshotStore(ctx context.Context, logger *slog.Logger) (cards.Store, error) {
	if dir := os.Getenv("CARDS_SNAPSHOT_DIR"); dir != "" {
		return cards.DirStore{Root: dir}, nil
	}
	// WithJSONReads: the SDK's default XML download path 404s on
	// fake-gcs-server's filesystem backend (encoded-slash object names).
	// JSON reads work on fake-gcs and on real GCS.
	client, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		return nil, err
	}
	project := envOr("PROJECT_ID", "mtg-local")
	bucket := envOr("CARDS_BUCKET", project+"-cards")
	if err := client.Bucket(bucket).Create(ctx, project, nil); err != nil {
		// An existing bucket lands here. A real permission problem
		// surfaces on the first write instead.
		logger.Info("bucket create skipped", "bucket", bucket, "note", err.Error())
	}
	return cards.NewGCSStore(client, bucket), nil
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
