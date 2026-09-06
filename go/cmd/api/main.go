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
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"cloud.google.com/go/firestore"
	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/agentsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/allowlist"
	"github.com/nkramber/mtg-deck-builder/go/internal/auth"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
	"github.com/nkramber/mtg-deck-builder/go/internal/collectionsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/decks"
	"github.com/nkramber/mtg-deck-builder/go/internal/decksvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/gcpenv"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/health"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
	"github.com/nkramber/mtg-deck-builder/go/internal/quality"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/ratelimit"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
	"github.com/nkramber/mtg-deck-builder/go/internal/spellbook"
	"github.com/nkramber/mtg-deck-builder/go/internal/usage"
)

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

// defaultReloadSeconds is how often the api looks for a newer snapshot.
// CARDS_RELOAD_SECONDS overrides it (make dev sets 15 for fast seeding).
const defaultReloadSeconds = 600

// sharedReadsPerMinute caps the public deck reads of one client address
// (D-315). Sixty is one a second, far above a person and low enough to
// bound a scan.
const sharedReadsPerMinute = 60

// maxRequestBytes bounds one request body before it enters memory. The
// largest expected body is a ManaBox export, under 5 MiB.
const maxRequestBytes = 8 << 20

// Server timing. A Chat stream holds a connection for minutes, so there
// is no ReadTimeout. IdleTimeout closes keep-alive connections that
// carry nothing. shutdownTimeout gives a stream that is mid-build time
// to end before the process exits: the build limit plus a margin for
// the store writes that follow it (D-303).
const (
	idleTimeout     = 120 * time.Second
	shutdownTimeout = agentsvc.DefaultBuildLimit + time.Minute
)

// Environment the api reads: PORT, ALLOW_DEBUG_USER, DEBUG_USER_ID,
// ALLOWED_ORIGINS, FIREBASE_AUTH_EMULATOR_HOST, CARDS_RELOAD_SECONDS,
// and the LLM_* variables of internal/llm, plus what gcpenv lists.

func main() {
	logger := gcpenv.NewLogger(os.Stdout)
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
	addr := ":" + gcpenv.EnvOr("PORT", "8080")
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	authOpts, err := authOptions(ctx, project, logger)
	if err != nil {
		return err
	}

	cardServer := cardsvc.New()
	store, storageClient, err := gcpenv.SnapshotStore(ctx, project, false, logger)
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
	// The invite list gates every deployed request (D-314, D-420). Local
	// mode allows every emulator user, and ALLOWLIST=1 turns the list on
	// there too, for a test of the refusal.
	if gcpenv.OnCloudRun() || os.Getenv("ALLOWLIST") == "1" {
		authOpts.opts = append(authOpts.opts, auth.WithAllowlist(allowlist.FromFirestore(fs)))
		logger.Info("the invite list gates every request", "document", allowlist.Collection+"/"+allowlist.Doc)
	}
	// The interceptor puts the user id in the context.
	userFn := auth.UserID
	// The repo reads the index for the summary of a collection stored
	// before the summary existed (D-398).
	collectionRepo := collections.NewRepo(fs).WithIndex(cardServer.Current)
	collectionServer := collectionsvc.New(collectionRepo, cardServer, userFn)
	rulesCfg, err := rules.Load()
	if err != nil {
		return fmt.Errorf("rules data: %w", err)
	}
	// The deck store holds what a build produced (D-245). One repo serves
	// the build and the reads.
	deckRepo := decks.NewRepo(fs)
	deckServer := decksvc.New(rulesCfg, cardServer,
		decksvc.WithCollections(collectionRepo),
		decksvc.WithDecks(deckRepo),
		decksvc.WithSessions(sessions.NewRepo(fs)),
		decksvc.WithUser(userFn))
	// The LLM role layer. Building it here proves the config and the
	// keys at startup, not on the first user turn.
	llmClient, err := llm.NewFromEnv(os.Getenv, logger)
	if err != nil {
		return fmt.Errorf("llm config: %w", err)
	}
	preconSrc := newPreconSource(cardServer, logger)
	// The deck quality model loads beside the card snapshot and swaps
	// when the worker fits a new one (PR-14B). No model grades nothing.
	metaStore, err := gcpenv.MetaStore(ctx, project, storageClient)
	if err != nil {
		return fmt.Errorf("meta store init: %w", err)
	}
	scorer := quality.NewScorer(nil)
	cardServer.SetQuality(scorer.CardQualities)
	// The precon table loads beside the model, for the precon exclusion
	// of PR-24 (D-407, D-408). No table excludes nothing, and the turn
	// says so.
	tableSrc := &preconTableSource{}
	agentServer, err := agentService(llmClient, fs, cardServer, collectionRepo, deckRepo, rulesCfg, preconSrc, tableSrc, scorer, userFn, logger)
	if err != nil {
		return fmt.Errorf("agent service: %w", err)
	}
	healthServer := health.New(version, cardServer)

	// The health RPC answers a probe, which carries no token.
	probeOpts := []connect.HandlerOption{connect.WithReadMaxBytes(maxRequestBytes)}
	opts := append([]connect.HandlerOption{connect.WithInterceptors(auth.Interceptor(authOpts.verifier, authOpts.opts...))}, probeOpts...)
	// The shared deck reads need no sign-in, and a limit per client
	// address bounds them (D-315). The limiter reads X-Forwarded-For, so
	// every visitor behind the proxy gets a bucket of their own.
	public := []string{mtgv1connect.DeckServiceGetSharedDeckProcedure, mtgv1connect.DeckServiceExportSharedDeckProcedure}
	limiter := ratelimit.New(sharedReadsPerMinute, time.Minute)
	deckOpts := append([]connect.HandlerOption{connect.WithInterceptors(
		limiter.Interceptor(public...),
		auth.Interceptor(authOpts.verifier, append(append([]auth.Option{}, authOpts.opts...), auth.WithPublic(public...))...),
	)}, probeOpts...)
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewHealthServiceHandler(healthServer, probeOpts...))
	mux.Handle(mtgv1connect.NewCardServiceHandler(cardServer, opts...))
	mux.Handle(mtgv1connect.NewCollectionServiceHandler(collectionServer, opts...))
	mux.Handle(mtgv1connect.NewDeckServiceHandler(deckServer, deckOpts...))
	mux.Handle(mtgv1connect.NewAgentServiceHandler(agentServer, opts...))
	// /healthz is liveness: the process answers. /readyz is readiness:
	// a card index is loaded, so the RPCs can answer.
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeStatus(w, healthServer.Status(), http.StatusOK)
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, _ *http.Request) {
		code := http.StatusServiceUnavailable
		if healthServer.Ready() {
			code = http.StatusOK
		}
		writeStatus(w, healthServer.Status(), code)
	})
	inflight := &inflightCounter{}
	handler := auth.CORS(auth.ParseOrigins(os.Getenv("ALLOWED_ORIGINS")), inflight.wrap(mux))

	// Listen first. The snapshot loads in the background, so the
	// Cloud Run startup probe sees a port inside its window. The card
	// handlers answer Unavailable until the first index lands.
	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       idleTimeout,
	}
	serveErr := make(chan error, 1)
	go func() {
		logger.Info("api listening", "addr", addr, "version", version)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
		}
		close(serveErr)
	}()
	loopCtx, stopLoop := context.WithCancel(ctx)
	loopDone := make(chan struct{})
	go func() {
		defer close(loopDone)
		loaded := loadSnapshot(loopCtx, store, cardServer, "", logger)
		preconSrc.refresh()
		reloadLoop(loopCtx, store, cardServer, loaded, logger, preconSrc.refresh)
	}()
	modelDone := make(chan struct{})
	go func() {
		defer close(modelDone)
		modelLoop(loopCtx, metaStore, scorer, logger)
	}()
	tableDone := make(chan struct{})
	go func() {
		defer close(tableDone)
		preconTableLoop(loopCtx, metaStore, tableSrc, logger)
	}()

	select {
	case err := <-serveErr:
		stopLoop()
		<-loopDone
		<-modelDone
		<-tableDone
		if err != nil {
			return fmt.Errorf("api server: %w", err)
		}
		return nil
	case <-ctx.Done():
	}
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	err = srv.Shutdown(shutdownCtx)
	if n := inflight.n.Load(); n > 0 {
		// The shutdown window ended with streams still open. Each one
		// is a turn the client must send again.
		logger.Warn("shutdown cut open requests", "count", n, "window", shutdownTimeout.String())
	}
	// A snapshot load mid-flight ends with the loop context, and the
	// exit waits for it, so a half-swapped index does not outlive the
	// server.
	stopLoop()
	<-loopDone
	if err != nil {
		return fmt.Errorf("api shutdown: %w", err)
	}
	return nil
}

// inflightCounter counts the requests a handler is serving, so the
// shutdown log can say how many streams it cut.
type inflightCounter struct{ n atomic.Int64 }

func (c *inflightCounter) wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.n.Add(1)
		defer c.n.Add(-1)
		next.ServeHTTP(w, r)
	})
}

// writeStatus writes the health document with the given HTTP code.
func writeStatus(w http.ResponseWriter, res *mtgv1.CheckResponse, code int) {
	body, err := protojson.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

// authSetup is the verifier and the interceptor options the api runs.
type authSetup struct {
	verifier auth.Verifier
	opts     []auth.Option
}

// authOptions picks the verifier and the fallback. Local mode keeps the debug user as
// the fallback for a request with no token. On Cloud Run the fallback
// exists only with ALLOW_DEBUG_USER=1, and a missing or refused token is
// Unauthenticated otherwise.
func authOptions(ctx context.Context, project string, logger *slog.Logger) (authSetup, error) {
	onCloudRun := gcpenv.OnCloudRun()
	allowDebug := os.Getenv("ALLOW_DEBUG_USER") == "1"
	var setup authSetup
	switch {
	case onCloudRun && !allowDebug:
		// No fallback. Every request needs a token.
	case onCloudRun:
		logger.Warn("debug user fallback enabled on Cloud Run by ALLOW_DEBUG_USER=1")
		setup.opts = append(setup.opts, auth.WithFallback(gcpenv.EnvOr("DEBUG_USER_ID", "local-dev")))
	default:
		setup.opts = append(setup.opts, auth.WithFallback(gcpenv.EnvOr("DEBUG_USER_ID", "local-dev")))
	}
	// The Firebase verifier needs a project and either the emulator or
	// application credentials. A local run with neither refuses every
	// token and serves the fallback only.
	haveFirebase := onCloudRun || os.Getenv("FIREBASE_AUTH_EMULATOR_HOST") != "" || os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") != ""
	if !haveFirebase {
		logger.Warn("no Firebase auth source, bearer tokens are refused and the debug user serves local requests")
		setup.verifier = auth.RejectAll()
		return setup, nil
	}
	v, err := auth.NewFirebase(ctx, project)
	if err != nil {
		return authSetup{}, fmt.Errorf("firebase auth: %w", err)
	}
	setup.verifier = v
	return setup, nil
}

// liveCards answers the rules engine from the current index. The index
// lands after the builder is made, so the builder must not hold one.
type liveCards struct{ src *cardsvc.Server }

func (l liveCards) ByOracleID(id string) (*mtgv1.Card, bool) {
	idx := l.src.Current()
	if idx == nil {
		return nil, false
	}
	return idx.ByOracleID(id)
}

// preconSource resolves the precon lists against the current index, once
// per snapshot. Before the first snapshot it answers nil, and the
// upgrade path then runs as an ordinary build.
type preconSource struct {
	src *cardsvc.Server
	log *slog.Logger
	mu  sync.Mutex
	idx *cards.Index
	set *precons.Set
}

func newPreconSource(src *cardsvc.Server, log *slog.Logger) *preconSource {
	return &preconSource{src: src, log: log}
}

// Current implements agentsvc.PreconSource. The resolve runs outside
// the lock, so a turn that arrives during it is not held. Two callers
// that resolve the same index at once both install the same set, and
// the first one wins.
func (p *preconSource) Current() *precons.Set {
	idx := p.src.Current()
	if idx == nil {
		return nil
	}
	p.mu.Lock()
	cachedIdx, cachedSet := p.idx, p.set
	p.mu.Unlock()
	if idx == cachedIdx {
		return cachedSet
	}
	set, err := precons.Load(idx)
	if err != nil {
		p.log.Warn("precon decklists unavailable, an upgrade keeps no share", "err", err)
		set = nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.idx == idx {
		return p.set
	}
	p.idx, p.set = idx, set
	return p.set
}

// refresh resolves the set for a new snapshot ahead of the first turn
// that needs it.
func (p *preconSource) refresh() { p.Current() }

// preconTableSource holds the newest precon table of the meta store. It
// implements agentsvc.PreconTableSource, and preconTableLoop swaps it.
type preconTableSource struct {
	mu    sync.Mutex
	table *precons.Table
}

// Table answers the current table, or nil before the first load.
func (p *preconTableSource) Table() *precons.Table {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.table
}

func (p *preconTableSource) swap(t *precons.Table) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.table = t
}

// preconTableLoop loads the newest precon table and swaps it in when the
// meta job stores a newer version, on the snapshot cadence (D-472).
func preconTableLoop(ctx context.Context, store meta.ObjectStore, src *preconTableSource, logger *slog.Logger) {
	seconds := defaultReloadSeconds
	if v, err := strconv.Atoi(os.Getenv("CARDS_RELOAD_SECONDS")); err == nil && v > 0 {
		seconds = v
	}
	load := func() {
		version, err := meta.LatestPreconsVersion(ctx, store)
		if err != nil {
			logger.Error("precon table version check failed", "err", err)
			return
		}
		current := src.Table()
		if version == "" || (current != nil && version == current.Version) {
			if version == "" && current == nil {
				logger.Warn("no precon table in store yet, a precon exclusion excludes nothing")
			}
			return
		}
		rows, err := meta.ReadPrecons(ctx, store, version)
		if err != nil {
			logger.Error("precon table load failed", "version", version, "err", err)
			return
		}
		src.swap(precons.NewTable(version, rows))
		logger.Info("precon table loaded", "version", version, "products", len(rows))
	}
	load()
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			load()
		}
	}
}

// agentService wires the question workflow and the build. The card
// index and the candidate builder feed the hints, so a question that
// names a value names a real card. The generator reads the live index.
// A missing price table only costs the cost field of the usage event.
func agentService(client *llm.Client, fs *firestore.Client, index *cardsvc.Server,
	cols *collections.Repo, deckRepo *decks.Repo, rulesCfg *rules.Config, preconSrc agentsvc.PreconSource,
	tableSrc agentsvc.PreconTableSource, scorer *quality.Scorer, userFn auth.UserFunc, logger *slog.Logger) (*agentsvc.Server, error) {
	cat, err := questions.Load()
	if err != nil {
		return nil, err
	}
	builder, err := candidates.New()
	if err != nil {
		return nil, err
	}
	// The bracket profile reads the live index's tags and the Commander
	// Spellbook endpoint, at the rate of D-459 (PR-14A).
	prof, err := profile.New(rulesCfg, func() *cards.TagIndex {
		if idx := index.Current(); idx != nil {
			return idx.Tags()
		}
		return nil
	}, spellbook.New(nil, "", logger))
	if err != nil {
		return nil, err
	}
	opts := []agentsvc.Option{
		agentsvc.WithLogger(logger),
		agentsvc.WithCandidates(index, builder),
		agentsvc.WithCollections(cols),
		agentsvc.WithDecks(generate.NewBuilder(client, rulesCfg, liveCards{index}, logger, generate.WithProfiler(prof), generate.WithScorer(scorer))),
		agentsvc.WithDeckStore(deckRepo),
		agentsvc.WithPreconSource(preconSrc),
		agentsvc.WithPreconTable(tableSrc),
		agentsvc.WithScorer(scorer),
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		logger.Warn("llm prices unavailable, the usage event carries no cost", "err", err)
	} else {
		opts = append(opts, agentsvc.WithPrices(prices))
	}
	// The per-user monthly spend cap (D-421). Cloud Run gets the $5
	// default, and local mode gets a cap only when SPEND_CAP_USD names one.
	if capUSD := spendCap(logger); capUSD > 0 {
		opts = append(opts, agentsvc.WithSpendCap(usage.NewRepo(fs), capUSD))
		logger.Info("the monthly spend cap is on", "cap_usd", capUSD)
	}
	return agentsvc.New(cat, client, sessions.NewRepo(fs), userFn, opts...)
}

// defaultSpendCapUSD is the cap of D-421.
const defaultSpendCapUSD = 5

// spendCap reads SPEND_CAP_USD. Unset, Cloud Run gets the default and
// local mode gets no cap. Zero turns the cap off anywhere, and a value
// that is not a number is an error the log names.
func spendCap(logger *slog.Logger) float64 {
	raw := os.Getenv("SPEND_CAP_USD")
	if raw == "" {
		if gcpenv.OnCloudRun() {
			return defaultSpendCapUSD
		}
		return 0
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil || v < 0 {
		logger.Error("SPEND_CAP_USD is not a number, so the cap is off", "value", raw)
		return 0
	}
	return v
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

// reloadLoop polls the store until ctx ends. onLoad runs after every
// version change, so a dependent of the index re-resolves.
func reloadLoop(ctx context.Context, store cards.Store, server *cardsvc.Server, lastVersion string, logger *slog.Logger, onLoad func()) {
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
			next := loadSnapshot(ctx, store, server, lastVersion, logger)
			if next != lastVersion && onLoad != nil {
				onLoad()
			}
			lastVersion = next
		}
	}
}

// modelLoop loads the newest quality model, then polls the store on the
// snapshot cadence and swaps a newer version in (PR-14B). A load error
// keeps the old model, so the next tick tries again.
func modelLoop(ctx context.Context, store meta.ObjectStore, scorer *quality.Scorer, logger *slog.Logger) {
	seconds := defaultReloadSeconds
	if v, err := strconv.Atoi(os.Getenv("CARDS_RELOAD_SECONDS")); err == nil && v > 0 {
		seconds = v
	}
	load := func() {
		version, err := meta.LatestModel(ctx, store)
		if err != nil {
			logger.Error("quality model version check failed", "err", err)
			return
		}
		if version == "" || version == scorer.Version() {
			if version == "" && scorer.Version() == "" {
				logger.Warn("no quality model in store yet, decks carry no grade")
			}
			return
		}
		m, err := quality.Load(ctx, store)
		if err != nil {
			logger.Error("quality model load failed", "version", version, "err", err)
			return
		}
		scorer.Swap(m)
		logger.Info("quality model loaded", "version", m.Version, "formats", len(m.Formats))
	}
	load()
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			load()
		}
	}
}
