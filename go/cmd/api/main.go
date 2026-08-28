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
	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/agentsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
	"github.com/nkramber/mtg-deck-builder/go/internal/collectionsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/decks"
	"github.com/nkramber/mtg-deck-builder/go/internal/decksvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/health"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

// defaultReloadSeconds is how often the api looks for a newer snapshot.
// CARDS_RELOAD_SECONDS overrides it (make dev sets 15 for fast seeding).
const defaultReloadSeconds = 600

// maxRequestBytes bounds one request body before it enters memory
// (C-10). The largest expected body is a ManaBox export, under 5 MiB.
const maxRequestBytes = 8 << 20

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := run(ctx, logger)
	stop()
	if err != nil {
		logger.Error("api failed", "err", err)
		os.Exit(1)
	}
	logger.Info("api stopped")
}

// run starts the server and blocks until ctx ends or the server fails.
// Every fatal path returns here, so main has one exit.
func run(ctx context.Context, logger *slog.Logger) error {
	addr := ":" + envOr("PORT", "8080")
	project, err := projectID()
	if err != nil {
		return err
	}
	// C-3: the debug user must never reach Cloud Run by accident.
	// Firebase Auth lands in PR-11 and replaces the debug user.
	if os.Getenv("K_SERVICE") != "" && os.Getenv("ALLOW_DEBUG_USER") != "1" {
		return errors.New("refuse to start on Cloud Run with the debug user: " +
			"Firebase Auth lands in PR-11, set ALLOW_DEBUG_USER=1 to override")
	}
	if os.Getenv("K_SERVICE") != "" {
		logger.Warn("debug user enabled on Cloud Run by ALLOW_DEBUG_USER=1, Firebase Auth lands in PR-11")
	}

	cardServer := cardsvc.New()
	store, storageClient, err := snapshotStore(ctx, project)
	if err != nil {
		return fmt.Errorf("snapshot store init: %w", err)
	}
	if storageClient != nil {
		defer func() { _ = storageClient.Close() }()
	}
	// Firestore for user data. FIRESTORE_EMULATOR_HOST routes it to the
	// emulator in local mode.
	fs, err := firestore.NewClient(ctx, project)
	if err != nil {
		return fmt.Errorf("firestore init: %w", err)
	}
	defer func() { _ = fs.Close() }()
	// Debug user until Firebase Auth lands (PR-11): every request acts as
	// one local user.
	debugUser := func(context.Context) string { return envOr("DEBUG_USER_ID", "local-dev") }
	collectionRepo := collections.NewRepo(fs)
	collectionServer := collectionsvc.New(collectionRepo, cardServer, debugUser)
	rulesCfg, err := rules.Load()
	if err != nil {
		return fmt.Errorf("rules data: %w", err)
	}
	// The deck store holds what a build produced (D-245). Without it a
	// deck streams to the user and is gone, GetDeck and ListDecks have
	// nothing to read, and the variance row is dead.
	deckRepo := decks.NewRepo(fs)
	deckServer := decksvc.New(rulesCfg, cardServer,
		decksvc.WithCollections(collectionRepo, debugUser),
		decksvc.WithDecks(deckRepo, debugUser))
	// The LLM role layer (PR-10). Building it here proves the config and
	// the keys at startup, not on the first user turn. Without keys the
	// fixture fake stands in.
	llmClient, err := llm.NewFromEnv(os.Getenv, logger)
	if err != nil {
		return fmt.Errorf("llm config: %w", err)
	}
	agentServer, err := agentService(llmClient, fs, cardServer, collectionRepo, debugUser, logger)
	if err != nil {
		return fmt.Errorf("agent service: %w", err)
	}
	healthServer := health.New(version, cardServer)

	opts := []connect.HandlerOption{connect.WithReadMaxBytes(maxRequestBytes)}
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewHealthServiceHandler(healthServer, opts...))
	mux.Handle(mtgv1connect.NewCardServiceHandler(cardServer, opts...))
	mux.Handle(mtgv1connect.NewCollectionServiceHandler(collectionServer, opts...))
	mux.Handle(mtgv1connect.NewDeckServiceHandler(deckServer, opts...))
	mux.Handle(mtgv1connect.NewAgentServiceHandler(agentServer, opts...))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, err := protojson.Marshal(healthServer.Status())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(body)
	})

	// Listen first (C-17). The snapshot loads in the background, so the
	// Cloud Run startup probe sees a port inside its window. The card
	// handlers answer Unavailable until the first index lands.
	srv := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	serveErr := make(chan error, 1)
	go func() {
		logger.Info("api listening", "addr", addr, "version", version)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
		}
		close(serveErr)
	}()
	go func() {
		loaded := loadSnapshot(ctx, store, cardServer, "", logger)
		reloadLoop(ctx, store, cardServer, loaded, logger)
	}()

	select {
	case err := <-serveErr:
		if err != nil {
			return fmt.Errorf("api server: %w", err)
		}
		return nil
	case <-ctx.Done():
	}
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("api shutdown: %w", err)
	}
	return nil
}

// agentService wires the question workflow (PR-7). The card index and
// the candidate builder feed the PR-6 hints, so a question that names a
// value names a real card. A missing price table only costs the cost
// field of the usage event.
func agentService(client *llm.Client, fs *firestore.Client, index *cardsvc.Server,
	cols *collections.Repo, userFn agentsvc.UserFunc, logger *slog.Logger) (*agentsvc.Server, error) {
	cat, err := questions.Load()
	if err != nil {
		return nil, err
	}
	builder, err := candidates.New()
	if err != nil {
		return nil, err
	}
	opts := []agentsvc.Option{
		agentsvc.WithLogger(logger),
		agentsvc.WithCandidates(index, builder),
		agentsvc.WithCollections(cols),
		agentsvc.WithDeckStore(decks.NewRepo(fs)),
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		logger.Warn("llm prices unavailable, the usage event carries no cost", "err", err)
	} else {
		opts = append(opts, agentsvc.WithPrices(prices))
	}
	return agentsvc.New(cat, client, sessions.NewRepo(fs), userFn, opts...)
}

// projectID resolves the GCP project. Cloud Run sets neither PROJECT_ID
// nor GOOGLE_CLOUD_PROJECT by default, so a missing value is fatal
// unless an emulator host marks local mode (C-17).
func projectID() (string, error) {
	if p := envOr("PROJECT_ID", os.Getenv("GOOGLE_CLOUD_PROJECT")); p != "" {
		return p, nil
	}
	if os.Getenv("FIRESTORE_EMULATOR_HOST") != "" || os.Getenv("STORAGE_EMULATOR_HOST") != "" {
		return "mtg-local", nil
	}
	return "", errors.New("PROJECT_ID or GOOGLE_CLOUD_PROJECT must be set outside emulator mode")
}

// snapshotStore picks the snapshot backend. CARDS_SNAPSHOT_DIR selects a
// local directory (offline mode, tests). Default: the GCS bucket, which
// STORAGE_EMULATOR_HOST routes to fake-gcs-server in local mode. The
// returned client is nil for the directory backend.
func snapshotStore(ctx context.Context, project string) (cards.Store, *storage.Client, error) {
	if dir := os.Getenv("CARDS_SNAPSHOT_DIR"); dir != "" {
		return cards.DirStore{Root: dir}, nil, nil
	}
	// WithJSONReads: the SDK's default XML download path 404s on
	// fake-gcs-server's filesystem backend (encoded-slash object names).
	// JSON reads work on fake-gcs and on real GCS.
	client, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		return nil, nil, err
	}
	bucket := envOr("CARDS_BUCKET", project+"-cards")
	return cards.NewGCSStore(client, bucket), client, nil
}

// loadSnapshot installs the newest stored snapshot when its version
// differs from lastVersion. A load error keeps the old index and the
// old version, so the next tick tries again.
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
