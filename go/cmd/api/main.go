// Command api serves the Connect-RPC API.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"

	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
	"github.com/nkramber/mtg-deck-builder/go/internal/collectionsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/decksvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/health"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

// defaultReloadSeconds is how often the api looks for a newer snapshot.
// CARDS_RELOAD_SECONDS overrides it (make dev sets 15 for fast seeding).
const defaultReloadSeconds = 600

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	addr := ":" + envOr("PORT", "8080")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cardServer := cardsvc.New()
	store, err := snapshotStore(ctx)
	if err != nil {
		logger.Error("snapshot store init failed", "err", err)
		stop()
		os.Exit(1) //nolint:gocritic // deliberate: stop() already ran
	}
	// Load whatever snapshot exists. An empty store is not fatal: the
	// worker fills it and the reload loop picks it up.
	loadedVersion := loadSnapshot(ctx, store, cardServer, "", logger)
	go reloadLoop(ctx, store, cardServer, loadedVersion, logger)

	// Firestore for user data. FIRESTORE_EMULATOR_HOST routes it to the
	// emulator in local mode.
	fs, err := firestore.NewClient(ctx, envOr("PROJECT_ID", "mtg-local"))
	if err != nil {
		logger.Error("firestore init failed", "err", err)
		stop()
		os.Exit(1) //nolint:gocritic // deliberate: stop() already ran
	}
	// Debug user until Firebase Auth lands (PR-11): every request acts as
	// one local user. Never ship this beyond local mode.
	debugUser := func(context.Context) string { return envOr("DEBUG_USER_ID", "local-dev") }
	collectionServer := collectionsvc.New(collections.NewRepo(fs), cardServer, debugUser)
	rulesCfg, err := rules.Load()
	if err != nil {
		logger.Error("rules data broken", "err", err)
		stop()
		os.Exit(1) //nolint:gocritic // deliberate: stop() already ran
	}
	deckServer := decksvc.New(rulesCfg, cardServer)
	// The LLM role layer (PR-10). No call site exists until PR-7 and PR-8.
	// Building it here proves the config and the keys at startup, not on
	// the first user turn. Without keys the fixture fake stands in.
	if _, err := llm.NewFromEnv(os.Getenv, logger); err != nil {
		logger.Error("llm config broken", "err", err)
		stop()
		os.Exit(1) //nolint:gocritic // deliberate: stop() already ran
	}

	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewHealthServiceHandler(health.New(version)))
	mux.Handle(mtgv1connect.NewCardServiceHandler(cardServer))
	mux.Handle(mtgv1connect.NewCollectionServiceHandler(collectionServer))
	mux.Handle(mtgv1connect.NewDeckServiceHandler(deckServer))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		snapshot, age := "none", -1.0
		if idx := cardServer.Current(); idx != nil {
			snapshot = idx.AsOf.UTC().Format(time.RFC3339)
			age = time.Since(idx.AsOf).Hours()
		}
		_, _ = fmt.Fprintf(w, `{"status":"ok","version":%q,"card_snapshot":%q,"card_snapshot_age_hours":%.1f}`,
			version, snapshot, age)
	})

	srv := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		logger.Info("api listening", "addr", addr, "version", version)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("api server failed", "err", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("api shutdown failed", "err", err)
	}
	logger.Info("api stopped")
}

// snapshotStore picks the snapshot backend. CARDS_SNAPSHOT_DIR selects a
// local directory (offline mode, tests). Default: the GCS bucket, which
// STORAGE_EMULATOR_HOST routes to fake-gcs-server in local mode.
func snapshotStore(ctx context.Context) (cards.Store, error) {
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
	bucket := envOr("CARDS_BUCKET", envOr("PROJECT_ID", "mtg-local")+"-cards")
	return cards.NewGCSStore(client, bucket), nil
}

func loadSnapshot(ctx context.Context, store cards.Store, server *cardsvc.Server, lastVersion string, logger *slog.Logger) string {
	current, err := store.LatestVersion(ctx)
	if err != nil {
		logger.Error("snapshot version check failed", "err", err)
		return lastVersion
	}
	if current == "" || current == lastVersion {
		if current == "" {
			logger.Warn("no card snapshot in store yet")
		}
		return lastVersion
	}
	idx, err := cards.LoadIndex(ctx, store, logger)
	if err != nil {
		logger.Error("snapshot load failed", "version", current, "err", err)
		return lastVersion
	}
	if idx != nil {
		server.Swap(idx)
	}
	return current
}

func reloadLoop(ctx context.Context, store cards.Store, server *cardsvc.Server, lastVersion string, logger *slog.Logger) {
	seconds := defaultReloadSeconds
	if v, err := strconv.Atoi(os.Getenv("CARDS_RELOAD_SECONDS")); err == nil && v > 0 {
		seconds = v
	}
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			lastVersion = loadSnapshot(ctx, store, server, lastVersion, logger)
		}
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
