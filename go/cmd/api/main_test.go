package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/agentsvc"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/cardsvc"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/precons"
	"github.com/nkramber/decktome/go/internal/rules"
)

func gz(t *testing.T, s string) []byte {
	t.Helper()
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, _ = w.Write([]byte(s))
	_ = w.Close()
	return buf.Bytes()
}

// writeVersion stores one complete snapshot. broken makes oracle_cards
// unreadable so LoadIndex fails.
func writeVersion(t *testing.T, store cards.DirStore, version string, broken bool) {
	t.Helper()
	ctx := context.Background()
	for _, name := range cards.SnapshotFiles {
		w, err := store.Create(ctx, version, name)
		if err != nil {
			t.Fatal(err)
		}
		body := gz(t, "")
		if name == "oracle_cards.jsonl.gz" {
			if broken {
				body = []byte("not gzip")
			} else {
				body = gz(t, `{"id":"p1","oracle_id":"o1","name":"Card","layout":"normal"}`+"\n")
			}
		}
		_, _ = w.Write(body)
		_ = w.Close()
	}
	if err := store.Finalize(ctx, version); err != nil {
		t.Fatal(err)
	}
}

func TestLoadSnapshot(t *testing.T) {
	ctx := context.Background()
	logger := slog.Default()
	t.Run("empty store keeps nil index", func(t *testing.T) {
		store := cards.DirStore{Root: t.TempDir()}
		server := cardsvc.New()
		if v := loadSnapshot(ctx, store, server, "", logger); v != "" || server.Current() != nil {
			t.Fatalf("v=%q current=%v", v, server.Current())
		}
	})
	t.Run("first load swaps", func(t *testing.T) {
		store := cards.DirStore{Root: t.TempDir()}
		writeVersion(t, store, "20260824T090000", false)
		server := cardsvc.New()
		v := loadSnapshot(ctx, store, server, "", logger)
		if v != "20260824T090000" || server.Current() == nil {
			t.Fatalf("v=%q current=%v", v, server.Current())
		}
	})
	t.Run("version unchanged keeps the index", func(t *testing.T) {
		store := cards.DirStore{Root: t.TempDir()}
		writeVersion(t, store, "20260824T090000", false)
		server := cardsvc.New()
		v := loadSnapshot(ctx, store, server, "", logger)
		first := server.Current()
		v2 := loadSnapshot(ctx, store, server, v, logger)
		if v2 != v || server.Current() != first {
			t.Fatal("unchanged version must not reload")
		}
	})
	t.Run("load error keeps old index and old version", func(t *testing.T) {
		store := cards.DirStore{Root: t.TempDir()}
		writeVersion(t, store, "20260824T090000", false)
		server := cardsvc.New()
		v := loadSnapshot(ctx, store, server, "", logger)
		first := server.Current()
		writeVersion(t, store, "20260825T090000", true)
		v2 := loadSnapshot(ctx, store, server, v, logger)
		if v2 != v || server.Current() != first {
			t.Fatalf("broken snapshot replaced the index: v2=%q", v2)
		}
	})
	// REV-012: a new instance with a broken newest version served no card
	// data until the next version came.
	t.Run("a cold start serves the version before a broken one", func(t *testing.T) {
		store := cards.DirStore{Root: t.TempDir()}
		writeVersion(t, store, "20260824T090000", false)
		writeVersion(t, store, "20260825T090000", true)
		server := cardsvc.New()
		v := loadSnapshot(ctx, store, server, "", logger)
		if v != "20260824T090000" || server.Current() == nil {
			t.Fatalf("v=%q current=%v, want the older version served", v, server.Current())
		}
		first := server.Current()
		if v2 := loadSnapshot(ctx, store, server, v, logger); v2 != v || server.Current() != first {
			t.Errorf("the next tick moved to %q", v2)
		}
	})
}

func TestReloadLoop(t *testing.T) {
	t.Setenv("CARDS_RELOAD_SECONDS", "1")
	store := cards.DirStore{Root: t.TempDir()}
	server := cardsvc.New()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		reloadLoop(ctx, store, server, "", slog.Default(), nil)
		close(done)
	}()
	writeVersion(t, store, "20260824T090000", false)
	deadline := time.After(5 * time.Second)
	for server.Current() == nil {
		select {
		case <-deadline:
			cancel()
			t.Fatal("reload loop never picked up the snapshot")
		case <-time.After(50 * time.Millisecond):
		}
	}
	cancel()
	<-done
}

// TestPreconSourceFollowsTheIndex: the agent service is built before the
// first snapshot lands, so the precon set must resolve late and again on
// every snapshot change.
func TestPreconSourceFollowsTheIndex(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := cardsvc.New()
	src := newPreconSource(server, quiet)
	if src.Current() != nil {
		t.Fatal("a nil index gave a precon set")
	}

	store := cards.DirStore{Root: t.TempDir()}
	writeVersion(t, store, "20260824T090000", false)
	loadSnapshot(context.Background(), store, server, "", quiet)
	if server.Current() == nil {
		t.Fatal("the snapshot did not load")
	}
	first := src.Current()
	if first == nil {
		t.Fatal("the swapped index gave no precon set")
	}
	if src.Current() != first {
		t.Error("the same index resolved the set twice")
	}

	// The reload loop re-resolves on a version change.
	t.Setenv("CARDS_RELOAD_SECONDS", "1")
	writeVersion(t, store, "20260825T090000", false)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		reloadLoop(ctx, store, server, "20260824T090000", quiet, src.refresh)
		close(done)
	}()
	deadline := time.After(5 * time.Second)
	for src.Current() == first {
		select {
		case <-deadline:
			cancel()
			<-done
			t.Fatal("the reload loop did not re-resolve the precons")
		case <-time.After(50 * time.Millisecond):
		}
	}
	cancel()
	<-done
}

// TestAgentServiceBuildsWithoutAnIndex proves the wiring needs no
// snapshot at startup: the generator and the precons bind late.
func TestAgentServiceBuildsWithoutAnIndex(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	rulesCfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	client, err := llm.NewFromEnv(func(k string) string {
		if k == llm.EnvRequireKeys {
			return "0"
		}
		return ""
	}, quiet)
	if err != nil {
		t.Fatal(err)
	}
	server := cardsvc.New()
	src := newPreconSource(server, quiet)
	srv, err := agentService(client, nil, server, nil, nil, rulesCfg, src, &preconTableSource{}, nil, func(context.Context) string { return "u" }, quiet)
	if err != nil {
		t.Fatalf("agent service with a nil index: %v", err)
	}
	if srv == nil {
		t.Fatal("no server")
	}
	if lc := (liveCards{server}); func() bool { _, ok := lc.ByOracleID("x"); return ok }() {
		t.Error("a nil index answered a card")
	}
}

func TestAuthOptions(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	tests := []struct {
		name         string
		env          map[string]string
		wantFallback bool
		wantReject   bool
	}{
		{name: "local, no firebase", wantFallback: true, wantReject: true},
		{name: "local with a debug id", env: map[string]string{"DEBUG_USER_ID": "nate"}, wantFallback: true, wantReject: true},
		{name: "local with the emulator", env: map[string]string{"FIREBASE_AUTH_EMULATOR_HOST": "127.0.0.1:1"}, wantFallback: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range []string{"K_SERVICE", "ALLOW_DEBUG_USER", "DEBUG_USER_ID", "FIREBASE_AUTH_EMULATOR_HOST", "GOOGLE_APPLICATION_CREDENTIALS"} {
				t.Setenv(k, "")
			}
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			got, err := authOptions(context.Background(), "p", quiet)
			if err != nil {
				t.Fatalf("authOptions: %v", err)
			}
			if (len(got.opts) > 0) != tt.wantFallback {
				t.Errorf("fallback set = %v, want %v", len(got.opts) > 0, tt.wantFallback)
			}
			if tt.wantReject {
				if _, err := got.verifier.Verify(context.Background(), "tok"); !errors.Is(err, auth.ErrNotConfigured) {
					t.Errorf("local verifier accepted a token: %v", err)
				}
			}
		})
	}
}

// TestCloudRunRefusesLocalSwitches: on Cloud Run the debug user and the
// emulator host stop the start, and never open the API (D-925).
func TestCloudRunRefusesLocalSwitches(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	tests := []struct {
		name string
		env  map[string]string
	}{
		{name: "the debug user", env: map[string]string{"ALLOW_DEBUG_USER": "1"}},
		{name: "the emulator host", env: map[string]string{"FIREBASE_AUTH_EMULATOR_HOST": "127.0.0.1:1"}},
		{name: "any debug value", env: map[string]string{"ALLOW_DEBUG_USER": "true", "FIREBASE_AUTH_EMULATOR_HOST": "127.0.0.1:1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range []string{"ALLOW_DEBUG_USER", "DEBUG_USER_ID", "FIREBASE_AUTH_EMULATOR_HOST", "GOOGLE_APPLICATION_CREDENTIALS"} {
				t.Setenv(k, "")
			}
			t.Setenv("K_SERVICE", "api")
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			if _, err := authOptions(context.Background(), "p", quiet); err == nil {
				t.Fatal("authOptions started on Cloud Run with a local switch")
			}
		})
	}
	t.Run("a clean Cloud Run env passes the guard", func(t *testing.T) {
		if err := cloudRunAuthGuard(true, func(string) string { return "" }); err != nil {
			t.Fatalf("guard: %v", err)
		}
	})
	t.Run("local mode keeps both switches", func(t *testing.T) {
		env := map[string]string{"ALLOW_DEBUG_USER": "1", "FIREBASE_AUTH_EMULATOR_HOST": "127.0.0.1:1"}
		if err := cloudRunAuthGuard(false, func(k string) string { return env[k] }); err != nil {
			t.Fatalf("guard: %v", err)
		}
	})
}

// TestSpendCap: a bad value stops the start, and never turns the cap
// off (D-925).
func TestSpendCap(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		cloud   bool
		want    float64
		wantErr bool
	}{
		{name: "unset on Cloud Run takes the default", cloud: true, want: defaultSpendCapUSD},
		{name: "unset locally is off", want: 0},
		{name: "zero is off", raw: "0", cloud: true, want: 0},
		{name: "a number", raw: "7.5", cloud: true, want: 7.5},
		{name: "a word", raw: "five", cloud: true, wantErr: true},
		{name: "under zero", raw: "-1", cloud: true, wantErr: true},
		{name: "not a number", raw: "NaN", cloud: true, wantErr: true},
		{name: "infinite", raw: "Inf", cloud: true, wantErr: true},
		{name: "a word locally", raw: "five", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := spendCap(tt.cloud, func(k string) string {
				if k == "SPEND_CAP_USD" {
					return tt.raw
				}
				return ""
			})
			if (err != nil) != tt.wantErr {
				t.Fatalf("spendCap(%q) err = %v, wantErr %v", tt.raw, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("spendCap(%q) = %v, want %v", tt.raw, got, tt.want)
			}
		})
	}
}

// TestUnpricedRoles: a role whose model has no price row books no cost,
// so the cap needs a row for each real role (D-925).
func TestUnpricedRoles(t *testing.T) {
	cfg, err := llm.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		t.Fatal(err)
	}
	if got := unpricedRoles(cfg, prices); len(got) != 0 {
		t.Errorf("the shipped config has unpriced roles %v", got)
	}
	override := cfg.Clone()
	spec := override.Roles[llm.RoleGenerate]
	spec.Model = "a-model-with-no-row"
	override.Roles[llm.RoleGenerate] = spec
	if got := unpricedRoles(override, prices); len(got) != 1 || got[0] != llm.RoleGenerate {
		t.Errorf("unpricedRoles = %v, want [%s]", got, llm.RoleGenerate)
	}
	if got := unpricedRoles(cfg, nil); len(got) == 0 {
		t.Error("no price table left every role priced")
	}
	fake := cfg.Clone()
	for r, s := range fake.Roles {
		s.Provider = llm.FakeName
		fake.Roles[r] = s
	}
	if got := unpricedRoles(fake, nil); len(got) != 0 {
		t.Errorf("the fake provider needs no row, got %v", got)
	}
}

// TestPreconSourceResolvesOutsideTheLock: two callers that see a new
// index at once both get one set, and the resolve holds no lock.
func TestPreconSourceResolvesOutsideTheLock(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	server := cardsvc.New()
	src := newPreconSource(server, quiet)
	store := cards.DirStore{Root: t.TempDir()}
	writeVersion(t, store, "20260824T090000", false)
	loadSnapshot(context.Background(), store, server, "", quiet)
	var wg sync.WaitGroup
	sets := make([]*precons.Set, 4)
	for i := range sets {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sets[i] = src.Current()
		}()
	}
	wg.Wait()
	for i, set := range sets {
		if set == nil || set != sets[0] {
			t.Errorf("caller %d got %p, want the one set %p", i, set, sets[0])
		}
	}
}

func TestInflightCounter(t *testing.T) {
	c := &inflightCounter{}
	var seen int64
	h := c.wrap(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		seen = c.n.Load()
		w.WriteHeader(http.StatusOK)
	}))
	h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/x", nil))
	if seen != 1 || c.n.Load() != 0 {
		t.Errorf("in flight during = %d, after = %d", seen, c.n.Load())
	}
}

func TestShutdownWindowCoversABuild(t *testing.T) {
	if shutdownTimeout <= agentsvc.DefaultBuildLimit {
		t.Errorf("shutdown window %v does not cover a build of %v", shutdownTimeout, agentsvc.DefaultBuildLimit)
	}
}

// TestSpendCapOverrides is D-576: SPEND_CAP_OVERRIDES names a cap per
// email, and a malformed entry never becomes a cap of zero.
func TestSpendCapOverrides(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	tests := []struct {
		name string
		raw  string
		want map[string]float64
	}{
		{name: "unset", raw: "", want: nil},
		{name: "one email at zero", raw: "ann@example.com:0", want: map[string]float64{"ann@example.com": 0}},
		{name: "two emails", raw: "ann@example.com:0, bo@example.com:20", want: map[string]float64{"ann@example.com": 0, "bo@example.com": 20}},
		{name: "the case of the email drops", raw: "Ann@Example.COM:0", want: map[string]float64{"ann@example.com": 0}},
		{name: "no colon drops the entry", raw: "ann@example.com", want: map[string]float64{}},
		{name: "a cap that is not a number drops the entry", raw: "ann@example.com:many", want: map[string]float64{}},
		{name: "a cap under zero drops the entry", raw: "ann@example.com:-1", want: map[string]float64{}},
		{name: "no email drops the entry", raw: ":5", want: map[string]float64{}},
		{name: "a good entry survives a bad one", raw: "bad,ann@example.com:0", want: map[string]float64{"ann@example.com": 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SPEND_CAP_OVERRIDES", tt.raw)
			got := spendCapOverrides(quiet)
			if len(got) != len(tt.want) {
				t.Fatalf("overrides = %v, want %v", got, tt.want)
			}
			for email, capUSD := range tt.want {
				if got[email] != capUSD {
					t.Errorf("cap of %q = %v, want %v", email, got[email], capUSD)
				}
			}
		})
	}
}
