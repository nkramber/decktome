package cards

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/scryfall"
)

// fakeScryfall serves a bulk catalog and gzip JSONL files.
type fakeScryfall struct {
	srv       *httptest.Server
	updatedAt map[string]time.Time
	bodies    map[string]string
	// downloads is written from the handler goroutine and read by the
	// test, so it is atomic.
	downloads atomic.Int64
}

func newFakeScryfall(t *testing.T, at time.Time) *fakeScryfall {
	t.Helper()
	f := &fakeScryfall{
		updatedAt: map[string]time.Time{"oracle_cards": at, "default_cards": at, "oracle_tags": at},
		bodies:    map[string]string{"oracle_cards": "", "default_cards": "", "oracle_tags": ""},
	}
	f.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/bulk-data" {
			var items []string
			for typ, ts := range f.updatedAt {
				items = append(items, fmt.Sprintf(`{"type":%q,"updated_at":%q,"jsonl_download_uri":"http://%s/file/%s"}`,
					typ, ts.Format(time.RFC3339), r.Host, typ))
			}
			_, _ = fmt.Fprintf(w, `{"data":[%s]}`, strings.Join(items, ","))
			return
		}
		typ := strings.TrimPrefix(r.URL.Path, "/file/")
		body, ok := f.bodies[typ]
		if !ok {
			w.WriteHeader(404)
			return
		}
		f.downloads.Add(1)
		_, _ = w.Write(gzipBytes(t, body))
	}))
	t.Cleanup(f.srv.Close)
	return f
}

func (f *fakeScryfall) client() *scryfall.Client {
	return scryfall.New(f.srv.Client(), f.srv.URL, slog.Default())
}

func gzipBytes(t *testing.T, s string) []byte {
	t.Helper()
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write([]byte(s)); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func cardLine(oracleID string, legal map[string]string) string {
	var parts []string
	for k, v := range legal {
		parts = append(parts, fmt.Sprintf("%q:%q", k, v))
	}
	return fmt.Sprintf(`{"oracle_id":%q,"name":"Card %s","layout":"normal","legalities":{%s}}`,
		oracleID, oracleID, strings.Join(parts, ","))
}

func TestRefresh(t *testing.T) {
	ctx := context.Background()
	day1 := time.Date(2026, 8, 23, 9, 1, 0, 0, time.UTC)
	day2 := day1.Add(24 * time.Hour)

	t.Run("empty store downloads all three files", func(t *testing.T) {
		f := newFakeScryfall(t, day1)
		store := DirStore{Root: t.TempDir()}
		v, err := Refresh(ctx, f.client(), store, slog.Default())
		if err != nil {
			t.Fatal(err)
		}
		if v != VersionFor(day1) || f.downloads.Load() != 3 {
			t.Fatalf("version %q downloads %d", v, f.downloads.Load())
		}
		latest, _ := store.LatestVersion(ctx)
		if latest != v {
			t.Fatalf("latest %q, want %q", latest, v)
		}
	})
	t.Run("same or older remote skips", func(t *testing.T) {
		f := newFakeScryfall(t, day1)
		store := DirStore{Root: t.TempDir()}
		if _, err := Refresh(ctx, f.client(), store, slog.Default()); err != nil {
			t.Fatal(err)
		}
		f.downloads.Store(0)
		v, err := Refresh(ctx, f.client(), store, slog.Default())
		if err != nil || f.downloads.Load() != 0 || v != VersionFor(day1) {
			t.Fatalf("second run: v=%q downloads=%d err=%v", v, f.downloads.Load(), err)
		}
		for typ := range f.updatedAt {
			f.updatedAt[typ] = day1.Add(-time.Hour)
		}
		v, err = Refresh(ctx, f.client(), store, slog.Default())
		if err != nil || f.downloads.Load() != 0 || v != VersionFor(day1) {
			t.Fatalf("older remote: v=%q downloads=%d err=%v", v, f.downloads.Load(), err)
		}
	})
	t.Run("newer remote by one second downloads", func(t *testing.T) {
		f := newFakeScryfall(t, day1)
		store := DirStore{Root: t.TempDir()}
		if _, err := Refresh(ctx, f.client(), store, slog.Default()); err != nil {
			t.Fatal(err)
		}
		for typ := range f.updatedAt {
			f.updatedAt[typ] = day1.Add(time.Second)
		}
		f.downloads.Store(0)
		v, err := Refresh(ctx, f.client(), store, slog.Default())
		if err != nil || f.downloads.Load() != 3 || v != VersionFor(day1.Add(time.Second)) {
			t.Fatalf("v=%q downloads=%d err=%v", v, f.downloads.Load(), err)
		}
	})
	t.Run("bulk files on two dates skip the cycle", func(t *testing.T) {
		f := newFakeScryfall(t, day2)
		f.updatedAt["oracle_tags"] = day1
		store := DirStore{Root: t.TempDir()}
		v, err := Refresh(ctx, f.client(), store, slog.Default())
		if err != nil || v != "" || f.downloads.Load() != 0 {
			t.Fatalf("v=%q downloads=%d err=%v", v, f.downloads.Load(), err)
		}
	})
	t.Run("download failure leaves no complete version", func(t *testing.T) {
		f := newFakeScryfall(t, day1)
		delete(f.bodies, "oracle_tags")
		store := DirStore{Root: t.TempDir()}
		if _, err := Refresh(ctx, f.client(), store, slog.Default()); err == nil {
			t.Fatal("want error")
		}
		if latest, _ := store.LatestVersion(ctx); latest != "" {
			t.Fatalf("incomplete version visible: %q", latest)
		}
	})
	t.Run("stored file is the gzip body", func(t *testing.T) {
		f := newFakeScryfall(t, day1)
		f.bodies["oracle_cards"] = cardLine("a", map[string]string{"modern": "legal"}) + "\n"
		store := DirStore{Root: t.TempDir()}
		v, err := Refresh(ctx, f.client(), store, slog.Default())
		if err != nil {
			t.Fatal(err)
		}
		r, err := store.Open(ctx, v, "oracle_cards.jsonl.gz")
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = r.Close() }()
		cardList, err := LoadCards(r, "oracle_cards.jsonl.gz")
		if err != nil || len(cardList) != 1 {
			t.Fatalf("cards=%d err=%v", len(cardList), err)
		}
	})
}

func TestRefreshPrunes(t *testing.T) {
	ctx := context.Background()
	base := time.Date(2026, 8, 20, 9, 0, 0, 0, time.UTC)
	f := newFakeScryfall(t, base)
	store := DirStore{Root: t.TempDir()}
	for i := 0; i < KeepVersions+2; i++ {
		at := base.Add(time.Duration(i) * 24 * time.Hour)
		for typ := range f.updatedAt {
			f.updatedAt[typ] = at
		}
		if _, err := Refresh(ctx, f.client(), store, slog.Default()); err != nil {
			t.Fatal(err)
		}
	}
	versions, err := store.ListVersions(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(versions) != KeepVersions {
		t.Fatalf("versions after prune = %v, want %d", versions, KeepVersions)
	}
	if versions[len(versions)-1] != VersionFor(base.Add(time.Duration(KeepVersions+1)*24*time.Hour)) {
		t.Fatalf("newest version wrong: %v", versions)
	}
	entries, _ := os.ReadDir(store.Root)
	if len(entries) != KeepVersions {
		t.Fatalf("directories on disk = %d, want %d", len(entries), KeepVersions)
	}
}

// TestPruneKeepsIncompleteAndNewest covers the incomplete versions. A
// failed download older than the newest complete version goes, also
// when it is the newest incomplete one. An incomplete version newer than
// every complete one stays, because it can be a download in progress.
func TestPruneKeepsIncompleteAndNewest(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name           string
		complete       []string
		incomplete     []string
		keep           int
		wantComplete   []string
		wantIncomplete []string
	}{
		{
			name:           "stale incomplete versions go, the newest stays",
			complete:       []string{"20260820T090000", "20260821T090000", "20260822T090000"},
			incomplete:     []string{"20260818T090000", "20260819T090000", "20260823T090000"},
			keep:           1,
			wantComplete:   []string{"20260822T090000"},
			wantIncomplete: []string{"20260823T090000"},
		},
		{
			name:           "an incomplete version newer than every complete one stays",
			complete:       []string{"20260820T090000"},
			incomplete:     []string{"20260821T090000", "20260822T090000"},
			keep:           3,
			wantComplete:   []string{"20260820T090000"},
			wantIncomplete: []string{"20260821T090000", "20260822T090000"},
		},
		{
			name:         "the newest incomplete version goes when a complete one is newer",
			complete:     []string{"20260822T090000"},
			incomplete:   []string{"20260820T090000", "20260821T090000"},
			keep:         3,
			wantComplete: []string{"20260822T090000"},
		},
		{
			name:           "no complete version leaves every incomplete one",
			incomplete:     []string{"20260818T090000", "20260819T090000"},
			keep:           1,
			wantIncomplete: []string{"20260818T090000", "20260819T090000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := DirStore{Root: t.TempDir()}
			for _, v := range tt.complete {
				w, _ := store.Create(ctx, v, "oracle_cards.jsonl.gz")
				_ = w.Close()
				if err := store.Finalize(ctx, v); err != nil {
					t.Fatal(err)
				}
			}
			for _, v := range tt.incomplete {
				w, _ := store.Create(ctx, v, "oracle_cards.jsonl.gz")
				_ = w.Close()
			}
			if err := Prune(ctx, store, tt.keep, slog.Default()); err != nil {
				t.Fatal(err)
			}
			complete, _ := store.ListVersions(ctx)
			if strings.Join(complete, ",") != strings.Join(tt.wantComplete, ",") {
				t.Errorf("complete = %v, want %v", complete, tt.wantComplete)
			}
			incomplete, _ := store.ListIncompleteVersions(ctx)
			if strings.Join(incomplete, ",") != strings.Join(tt.wantIncomplete, ",") {
				t.Errorf("incomplete = %v, want %v", incomplete, tt.wantIncomplete)
			}
			for _, v := range tt.wantIncomplete {
				if _, err := os.Stat(filepath.Join(store.Root, v)); err != nil {
					t.Errorf("incomplete version %s removed: %v", v, err)
				}
			}
		})
	}
}

func writeSnapshot(t *testing.T, store DirStore, version, oracleCards string) {
	t.Helper()
	ctx := context.Background()
	for _, name := range SnapshotFiles {
		w, err := store.Create(ctx, version, name)
		if err != nil {
			t.Fatal(err)
		}
		body := ""
		if name == "oracle_cards.jsonl.gz" {
			body = oracleCards
		}
		if _, err := w.Write(gzipBytes(t, body)); err != nil {
			t.Fatal(err)
		}
		_ = w.Close()
	}
	if err := store.Finalize(ctx, version); err != nil {
		t.Fatal(err)
	}
}

// TestLegalityDiff covers C-2: coverage is a data fact, not a date.
func TestLegalityDiff(t *testing.T) {
	ctx := context.Background()
	legalA := cardLine("a", map[string]string{"modern": "legal", "legacy": "legal"})
	bannedA := cardLine("a", map[string]string{"modern": "banned", "legacy": "legal"})
	legalB := cardLine("b", map[string]string{"modern": "legal"})
	faceOnly := `{"name":"Reversible","layout":"reversible_card","legalities":{"modern":"legal"},"card_faces":[{"oracle_id":"r1"}]}`
	tests := []struct {
		name        string
		before      string
		after       string
		wantChanged int
	}{
		{"same data", legalA + "\n" + legalB + "\n", legalA + "\n" + legalB + "\n", 0},
		{"one ban", legalA + "\n" + legalB + "\n", bannedA + "\n" + legalB + "\n", 1},
		{"new card", legalA + "\n", legalA + "\n" + legalB + "\n", 1},
		{"removed card", legalA + "\n" + legalB + "\n", legalA + "\n", 1},
		{"reversible card keyed by face oracle id", faceOnly + "\n", faceOnly + "\n", 0},
		{"blank lines ignored", legalA + "\n\n", "\n" + legalA + "\n", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := DirStore{Root: t.TempDir()}
			writeSnapshot(t, store, "20260823T090000", tt.before)
			writeSnapshot(t, store, "20260824T090000", tt.after)
			changed, err := LegalityDiff(ctx, store, "20260823T090000", "20260824T090000")
			if err != nil {
				t.Fatal(err)
			}
			if changed != tt.wantChanged {
				t.Errorf("changed = %d, want %d", changed, tt.wantChanged)
			}
		})
	}
	t.Run("missing version is an error", func(t *testing.T) {
		store := DirStore{Root: t.TempDir()}
		writeSnapshot(t, store, "20260824T090000", legalA+"\n")
		if _, err := LegalityDiff(ctx, store, "20260823T090000", "20260824T090000"); err == nil {
			t.Fatal("want error for missing old version")
		}
	})
}

func TestReadLines(t *testing.T) {
	t.Run("gzip stream", func(t *testing.T) {
		data := gzipBytes(t, "one\n\ntwo\n")
		var got []string
		err := readLines(bytes.NewReader(data), "x.jsonl.gz", func(line []byte) error {
			got = append(got, string(line))
			return nil
		})
		if err != nil || strings.Join(got, ",") != "one,two" {
			t.Fatalf("got %v err %v", got, err)
		}
	})
	t.Run("truncated gzip fails", func(t *testing.T) {
		data := gzipBytes(t, strings.Repeat("line\n", 5000))
		cut := data[:len(data)/2]
		err := readLines(bytes.NewReader(cut), "x.jsonl.gz", func([]byte) error { return nil })
		if err == nil {
			t.Fatal("want error on truncated gzip")
		}
	})
	t.Run("plain stream", func(t *testing.T) {
		n := 0
		err := readLines(strings.NewReader("a\nb\n"), "x.jsonl", func([]byte) error { n++; return nil })
		if err != nil || n != 2 {
			t.Fatalf("n=%d err=%v", n, err)
		}
	})
	t.Run("callback error names the line", func(t *testing.T) {
		err := readLines(strings.NewReader("a\nb\n"), "x.jsonl", func(line []byte) error {
			if string(line) == "b" {
				return io.ErrUnexpectedEOF
			}
			return nil
		})
		if err == nil || !strings.Contains(err.Error(), "line 2") {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestDirStoreSafety(t *testing.T) {
	ctx := context.Background()
	store := DirStore{Root: t.TempDir()}
	bad := []string{"../x", "a/b", "..", "", `a\b`, "x..y"}
	for _, v := range bad {
		if _, err := store.Open(ctx, v, "f"); err == nil {
			t.Errorf("Open accepted version %q", v)
		}
		if _, err := store.Create(ctx, "20260824T000000", v); err == nil {
			t.Errorf("Create accepted file %q", v)
		}
		if err := store.Finalize(ctx, v); err == nil {
			t.Errorf("Finalize accepted version %q", v)
		}
		if err := store.DeleteVersion(ctx, v); err == nil {
			t.Errorf("DeleteVersion accepted version %q", v)
		}
	}
	if entries, _ := os.ReadDir(store.Root); len(entries) != 0 {
		t.Errorf("bad names created entries: %v", entries)
	}
}

func TestDirStoreListAndDelete(t *testing.T) {
	ctx := context.Background()
	store := DirStore{Root: t.TempDir()}
	versions, err := store.ListVersions(ctx)
	if err != nil || len(versions) != 0 {
		t.Fatalf("empty: %v %v", versions, err)
	}
	for _, v := range []string{"20260822T090000", "20260820T090000", "20260821T090000"} {
		w, _ := store.Create(ctx, v, "oracle_cards.jsonl.gz")
		_ = w.Close()
		if err := store.Finalize(ctx, v); err != nil {
			t.Fatal(err)
		}
	}
	versions, _ = store.ListVersions(ctx)
	if strings.Join(versions, ",") != "20260820T090000,20260821T090000,20260822T090000" {
		t.Fatalf("versions = %v", versions)
	}
	if err := store.DeleteVersion(ctx, "20260820T090000"); err != nil {
		t.Fatal(err)
	}
	versions, _ = store.ListVersions(ctx)
	if len(versions) != 2 || versions[0] != "20260821T090000" {
		t.Fatalf("after delete: %v", versions)
	}
	if err := store.DeleteVersion(ctx, "20260820T090000"); err != nil {
		t.Fatalf("delete of a missing version must be idempotent: %v", err)
	}
}

func TestNewerVersion(t *testing.T) {
	tests := []struct {
		remote, current string
		want            bool
	}{
		{"20260824T090000", "", true},
		{"20260824T090000", "20260824T090000", false},
		{"20260824T090001", "20260824T090000", true},
		{"20260824T090000", "20260824T090001", false},
		{"20260824T090000", "garbage", true},
		{"garbage", "20260824T090000", false},
	}
	for _, tt := range tests {
		if got := newerVersion(tt.remote, tt.current); got != tt.want {
			t.Errorf("newerVersion(%q, %q) = %v, want %v", tt.remote, tt.current, got, tt.want)
		}
	}
}

func TestLegalityDiffMarker(t *testing.T) {
	ctx := context.Background()
	store := DirStore{Root: t.TempDir()}
	writeSnapshot(t, store, "20260823T090000", "")
	writeSnapshot(t, store, "20260824T090000", "")

	if _, ok, err := store.ReadLegalityDiff(ctx, "20260824T090000"); err != nil || ok {
		t.Fatalf("no marker yet: ok=%v err=%v", ok, err)
	}
	if last, err := LastLegalityDiff(ctx, store); err != nil || !last.IsZero() {
		t.Fatalf("no markers: last=%v err=%v", last, err)
	}
	rec := LegalityDiffRecord{Announcement: "2026-08-23", SnapshotAsOf: "2026-08-24T09:00:00Z", LagHours: 33, ChangedCards: 4}
	if err := store.WriteLegalityDiff(ctx, "20260824T090000", rec); err != nil {
		t.Fatal(err)
	}
	got, ok, err := store.ReadLegalityDiff(ctx, "20260824T090000")
	if err != nil || !ok || got != rec {
		t.Fatalf("read back: %+v ok=%v err=%v", got, ok, err)
	}
	raw, err := os.ReadFile(filepath.Join(store.Root, "20260824T090000", legalityDiffFile))
	if err != nil {
		t.Fatal(err)
	}
	want := `{"announcement":"2026-08-23","snapshot_as_of":"2026-08-24T09:00:00Z","lag_hours":33,"changed_cards":4}`
	if string(raw) != want {
		t.Errorf("marker json = %s", raw)
	}
	last, err := LastLegalityDiff(ctx, store)
	if err != nil || last.Format(time.RFC3339) != "2026-08-24T09:00:00Z" {
		t.Fatalf("last = %v err=%v", last, err)
	}
	// A marker on an older version does not shadow the newer one.
	if err := store.WriteLegalityDiff(ctx, "20260823T090000", LegalityDiffRecord{SnapshotAsOf: "2026-08-23T09:00:00Z"}); err != nil {
		t.Fatal(err)
	}
	if last, _ = LastLegalityDiff(ctx, store); last.Format(time.RFC3339) != "2026-08-24T09:00:00Z" {
		t.Errorf("last = %v, want the newest marker", last)
	}
	// Prune keeps the marker with its version and drops it with its version.
	writeSnapshot(t, store, "20260825T090000", "")
	if err := Prune(ctx, store, 2, slog.Default()); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := store.ReadLegalityDiff(ctx, "20260824T090000"); !ok {
		t.Error("marker of a kept version was lost")
	}
	if _, ok, _ := store.ReadLegalityDiff(ctx, "20260823T090000"); ok {
		t.Error("marker of a pruned version survived")
	}
	if _, err := os.Stat(filepath.Join(store.Root, "20260823T090000")); !os.IsNotExist(err) {
		t.Error("pruned version directory still exists")
	}
	if err := store.WriteLegalityDiff(ctx, "../evil", rec); err == nil {
		t.Error("marker write accepted a bad version")
	}
}

// ctxWriter records the context its store got and fails a Close after a
// cancel, the way a GCS writer does.
type ctxWriter struct {
	ctx   context.Context
	buf   bytes.Buffer
	saved *[]byte
}

func (w *ctxWriter) Write(p []byte) (int, error) { return w.buf.Write(p) }

func (w *ctxWriter) Close() error {
	if err := w.ctx.Err(); err != nil {
		return err
	}
	*w.saved = w.buf.Bytes()
	return nil
}

type ctxStore struct {
	DirStore
	saved []byte
}

func (s *ctxStore) Create(ctx context.Context, _, _ string) (io.WriteCloser, error) {
	return &ctxWriter{ctx: ctx, saved: &s.saved}, nil
}

// TestCopyBulkCancelsTheWriterOnError: a copy that fails must not commit a
// truncated object, so the writer's context is canceled before Close.
func TestCopyBulkCancelsTheWriterOnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", "100")
		_, _ = w.Write([]byte("abc"))
	}))
	t.Cleanup(srv.Close)
	client := scryfall.New(srv.Client(), srv.URL, slog.Default())
	store := &ctxStore{DirStore: DirStore{Root: t.TempDir()}}
	err := copyBulk(context.Background(), client, store, "20260823T090000", "oracle_cards.jsonl.gz", srv.URL+"/file")
	if err == nil {
		t.Fatal("a truncated download must fail")
	}
	if store.saved != nil {
		t.Errorf("the truncated object was committed: %q", store.saved)
	}
}
