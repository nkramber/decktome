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

	"github.com/nkramber/mtg-deck-builder/go/internal/agentsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/auth"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
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
	srv, err := agentService(client, nil, server, nil, nil, rulesCfg, src, func(context.Context) string { return "u" }, quiet)
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
		{name: "cloud run refuses the fallback", env: map[string]string{"K_SERVICE": "api"}, wantFallback: false},
		{name: "cloud run with the override", env: map[string]string{"K_SERVICE": "api", "ALLOW_DEBUG_USER": "1"}, wantFallback: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range []string{"K_SERVICE", "ALLOW_DEBUG_USER", "DEBUG_USER_ID", "FIREBASE_AUTH_EMULATOR_HOST", "GOOGLE_APPLICATION_CREDENTIALS"} {
				t.Setenv(k, "")
			}
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			if tt.env["K_SERVICE"] != "" {
				// The Firebase client needs credentials on Cloud Run, which
				// the test host has not. The fallback decision is what the
				// test reads, so it stops before the verifier.
				t.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "127.0.0.1:1")
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
