package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
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
		reloadLoop(ctx, store, server, "", slog.Default())
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

func TestProjectID(t *testing.T) {
	for _, k := range []string{"PROJECT_ID", "GOOGLE_CLOUD_PROJECT", "FIRESTORE_EMULATOR_HOST", "STORAGE_EMULATOR_HOST"} {
		t.Setenv(k, "")
	}
	if _, err := projectID(); err == nil {
		t.Error("want error with no project and no emulator")
	}
	t.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8081")
	if p, err := projectID(); err != nil || p != "mtg-local" {
		t.Errorf("emulator mode: %q %v", p, err)
	}
	t.Setenv("GOOGLE_CLOUD_PROJECT", "real-proj")
	if p, _ := projectID(); p != "real-proj" {
		t.Errorf("GOOGLE_CLOUD_PROJECT: %q", p)
	}
	t.Setenv("PROJECT_ID", "explicit")
	if p, _ := projectID(); p != "explicit" {
		t.Errorf("PROJECT_ID wins: %q", p)
	}
}
