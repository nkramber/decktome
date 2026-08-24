package cards

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// A Snapshot is one day's bulk files, stored under one version prefix.
// The version is the Scryfall updated_at in UTC, formatted 20060102T150405.
const versionFormat = "20060102T150405"

// SnapshotFiles are the files every snapshot holds, in store order.
var SnapshotFiles = []string{"oracle_cards.jsonl.gz", "default_cards.jsonl.gz", "oracle_tags.jsonl.gz"}

// completeMarker is written last. A version without it is mid-write or
// abandoned, and LatestVersion never returns it. Rule: a listed snapshot
// implies complete artifacts, no third state.
const completeMarker = "complete"

// Store persists snapshots. GCS in the cloud, a directory in tests.
type Store interface {
	// LatestVersion returns the newest stored version, or "" when none.
	LatestVersion(ctx context.Context) (string, error)
	// Open reads one file of one version. The caller closes it.
	Open(ctx context.Context, version, file string) (io.ReadCloser, error)
	// Create writes one file of one version. Closing commits it.
	Create(ctx context.Context, version, file string) (io.WriteCloser, error)
	// Finalize marks a version complete. LatestVersion sees it after this.
	Finalize(ctx context.Context, version string) error
}

// VersionTime parses a version string back into its time.
func VersionTime(version string) (time.Time, error) {
	return time.Parse(versionFormat, version)
}

// VersionFor formats a snapshot time as a version string.
func VersionFor(t time.Time) string {
	return t.UTC().Format(versionFormat)
}

// GCSStore stores snapshots under gs://<bucket>/scryfall/<version>/<file>.
// STORAGE_EMULATOR_HOST routes it to fake-gcs-server in local mode.
type GCSStore struct {
	client *storage.Client
	bucket string
}

// NewGCSStore wraps an existing client. The caller owns the client.
func NewGCSStore(client *storage.Client, bucket string) *GCSStore {
	return &GCSStore{client: client, bucket: bucket}
}

// LatestVersion returns the newest stored version, or "" when none.
func (s *GCSStore) LatestVersion(ctx context.Context) (string, error) {
	// Complete versions only: list the markers, not the directories.
	it := s.client.Bucket(s.bucket).Objects(ctx, &storage.Query{Prefix: "scryfall/"})
	var versions []string
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return "", fmt.Errorf("list snapshots: %w", err)
		}
		dir, file := filepath.Split(attrs.Name)
		if file == completeMarker {
			versions = append(versions, filepath.Base(filepath.Clean(dir)))
		}
	}
	if len(versions) == 0 {
		return "", nil
	}
	sort.Strings(versions)
	return versions[len(versions)-1], nil
}

// Finalize marks a version complete.
func (s *GCSStore) Finalize(ctx context.Context, version string) error {
	w := s.client.Bucket(s.bucket).Object("scryfall/" + version + "/" + completeMarker).NewWriter(ctx)
	if _, err := w.Write([]byte("ok")); err != nil {
		_ = w.Close()
		return err
	}
	return w.Close()
}

// Open reads one file of one version.
func (s *GCSStore) Open(ctx context.Context, version, file string) (io.ReadCloser, error) {
	return s.client.Bucket(s.bucket).Object("scryfall/" + version + "/" + file).NewReader(ctx)
}

// Create writes one file of one version. Closing commits it.
func (s *GCSStore) Create(ctx context.Context, version, file string) (io.WriteCloser, error) {
	return s.client.Bucket(s.bucket).Object("scryfall/" + version + "/" + file).NewWriter(ctx), nil
}

// DirStore stores snapshots under <root>/<version>/<file>.
// Tests and offline local mode use it.
type DirStore struct {
	Root string
}

// LatestVersion returns the newest stored version, or "" when none.
func (s DirStore) LatestVersion(_ context.Context) (string, error) {
	entries, err := os.ReadDir(s.Root)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	var versions []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if _, err := os.Stat(filepath.Join(s.Root, e.Name(), completeMarker)); err == nil {
			versions = append(versions, e.Name())
		}
	}
	if len(versions) == 0 {
		return "", nil
	}
	sort.Strings(versions)
	return versions[len(versions)-1], nil
}

// Open reads one file of one version.
func (s DirStore) Open(_ context.Context, version, file string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(s.Root, version, file))
}

// Create writes one file of one version.
func (s DirStore) Create(_ context.Context, version, file string) (io.WriteCloser, error) {
	dir := filepath.Join(s.Root, version)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(dir, file))
}

// Finalize marks a version complete.
func (s DirStore) Finalize(_ context.Context, version string) error {
	return os.WriteFile(filepath.Join(s.Root, version, completeMarker), []byte("ok"), 0o644)
}
