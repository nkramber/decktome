package cards

import (
	"context"
	"io"
	"os"
	"testing"

	"cloud.google.com/go/storage"
)

// TestLiveFakeGCS drives GCSStore against a running fake-gcs-server.
// Guarded: set FAKE_GCS_LIVE=1 and STORAGE_EMULATOR_HOST to run it.
func TestLiveFakeGCS(t *testing.T) {
	if os.Getenv("FAKE_GCS_LIVE") != "1" {
		t.Skip("set FAKE_GCS_LIVE=1 to run")
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		t.Fatal(err)
	}
	s := NewGCSStore(client, "mtg-local-cards")
	v, err := s.LatestVersion(ctx)
	if err != nil {
		t.Fatalf("LatestVersion: %v", err)
	}
	t.Logf("latest version: %q", v)
	if v == "" {
		t.Fatal("no complete version visible")
	}
	for _, f := range SnapshotFiles {
		r, err := s.Open(ctx, v, f)
		if err != nil {
			t.Fatalf("Open %s: %v", f, err)
		}
		n, err := io.CopyN(io.Discard, r, 100)
		_ = r.Close()
		if err != nil || n != 100 {
			t.Fatalf("read %s: n=%d err=%v", f, n, err)
		}
		t.Logf("open+read ok: %s", f)
	}
}
