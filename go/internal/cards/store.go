package cards

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
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

// gcsPrefix is the object prefix of every snapshot in a bucket.
const gcsPrefix = "scryfall/"

// Store persists snapshots. GCS in the cloud, a directory in tests.
type Store interface {
	// LatestVersion returns the newest stored version, or "" when none.
	LatestVersion(ctx context.Context) (string, error)
	// ListVersions returns every complete version, oldest first.
	ListVersions(ctx context.Context) ([]string, error)
	// ListIncompleteVersions returns every version with files and no
	// completion marker, oldest first. Prune reads it. A download in
	// progress is one of them, so a caller never deletes the newest.
	ListIncompleteVersions(ctx context.Context) ([]string, error)
	// Open reads one file of one version. The caller closes it.
	Open(ctx context.Context, version, file string) (io.ReadCloser, error)
	// Create writes one file of one version. Closing commits it.
	Create(ctx context.Context, version, file string) (io.WriteCloser, error)
	// Finalize marks a version complete. LatestVersion sees it after this.
	Finalize(ctx context.Context, version string) error
	// DeleteVersion removes one version and all its files.
	DeleteVersion(ctx context.Context, version string) error
	// WriteLegalityDiff stores the M-2 marker of one version.
	WriteLegalityDiff(ctx context.Context, version string, rec LegalityDiffRecord) error
	// ReadLegalityDiff returns the marker of one version, or false when
	// the version has none.
	ReadLegalityDiff(ctx context.Context, version string) (LegalityDiffRecord, bool, error)
}

// legalityDiffFile is the M-2 marker name inside a version. Prune keeps
// it with its version, because it lives under the version prefix.
const legalityDiffFile = "legality_diff.json"

// LegalityDiffRecord is the M-2 marker: which announcement a snapshot
// covered, and the lag from the announcement date to the snapshot.
type LegalityDiffRecord struct {
	// Announcement is the covered announcement date, YYYY-MM-DD, or ""
	// when the diff matched no calendar date.
	Announcement string `json:"announcement"`
	// SnapshotAsOf is the snapshot time, RFC 3339.
	SnapshotAsOf string  `json:"snapshot_as_of"`
	LagHours     float64 `json:"lag_hours"`
	ChangedCards int     `json:"changed_cards"`
}

// writeDiffTo encodes the marker through Create, so both backends share
// one code path.
func writeDiffTo(ctx context.Context, s Store, version string, rec LegalityDiffRecord) error {
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	w, err := s.Create(ctx, version, legalityDiffFile)
	if err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		_ = w.Close()
		return err
	}
	return w.Close()
}

func readDiffFrom(r io.ReadCloser, err error, notExist func(error) bool) (LegalityDiffRecord, bool, error) {
	if err != nil {
		if notExist(err) {
			return LegalityDiffRecord{}, false, nil
		}
		return LegalityDiffRecord{}, false, err
	}
	defer func() { _ = r.Close() }()
	var rec LegalityDiffRecord
	if err := json.NewDecoder(r).Decode(&rec); err != nil {
		return LegalityDiffRecord{}, false, fmt.Errorf("%s: %w", legalityDiffFile, err)
	}
	return rec, true, nil
}

// LastLegalityDiff returns the snapshot time of the newest version that
// carries an M-2 marker, or the zero time when none does. The worker
// reads it at start, so no state lives in process memory (C-11).
func LastLegalityDiff(ctx context.Context, s Store) (time.Time, error) {
	versions, err := s.ListVersions(ctx)
	if err != nil {
		return time.Time{}, err
	}
	for i := len(versions) - 1; i >= 0; i-- {
		rec, ok, err := s.ReadLegalityDiff(ctx, versions[i])
		if err != nil {
			return time.Time{}, err
		}
		if !ok {
			continue
		}
		t, err := time.Parse(time.RFC3339, rec.SnapshotAsOf)
		if err != nil {
			return time.Time{}, fmt.Errorf("%s/%s: %w", versions[i], legalityDiffFile, err)
		}
		return t, nil
	}
	return time.Time{}, nil
}

// VersionTime parses a version string back into its time.
func VersionTime(version string) (time.Time, error) {
	return time.Parse(versionFormat, version)
}

// VersionFor formats a snapshot time as a version string.
func VersionFor(t time.Time) string {
	return t.UTC().Format(versionFormat)
}

// latestOf picks the newest version. Versions sort as strings because
// the format is fixed width, so a string sort equals a time sort.
func latestOf(versions []string) string {
	if len(versions) == 0 {
		return ""
	}
	return versions[len(versions)-1]
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
	versions, err := s.ListVersions(ctx)
	if err != nil {
		return "", err
	}
	return latestOf(versions), nil
}

// ListVersions returns every complete version, oldest first.
func (s *GCSStore) ListVersions(ctx context.Context) ([]string, error) {
	complete, _, err := s.listVersions(ctx)
	return complete, err
}

// ListIncompleteVersions returns every version with objects and no
// completion marker, oldest first.
func (s *GCSStore) ListIncompleteVersions(ctx context.Context) ([]string, error) {
	_, incomplete, err := s.listVersions(ctx)
	return incomplete, err
}

// listVersions walks the bucket once and splits the versions by marker.
func (s *GCSStore) listVersions(ctx context.Context) (complete, incomplete []string, err error) {
	it := s.client.Bucket(s.bucket).Objects(ctx, &storage.Query{Prefix: gcsPrefix})
	marked := map[string]bool{}
	seen := map[string]bool{}
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("list snapshots: %w", err)
		}
		// Object names are slash paths on every platform, so path, not
		// filepath, splits them.
		dir, file := path.Split(attrs.Name)
		version := path.Base(path.Clean(dir))
		if version == "." || version == "/" {
			continue
		}
		seen[version] = true
		if file == completeMarker {
			marked[version] = true
		}
	}
	for version := range seen {
		if marked[version] {
			complete = append(complete, version)
		} else {
			incomplete = append(incomplete, version)
		}
	}
	sort.Strings(complete)
	sort.Strings(incomplete)
	return complete, incomplete, nil
}

// Finalize marks a version complete.
func (s *GCSStore) Finalize(ctx context.Context, version string) error {
	w := s.client.Bucket(s.bucket).Object(gcsPrefix + version + "/" + completeMarker).NewWriter(ctx)
	if _, err := w.Write([]byte("ok")); err != nil {
		_ = w.Close()
		return err
	}
	return w.Close()
}

// Open reads one file of one version.
func (s *GCSStore) Open(ctx context.Context, version, file string) (io.ReadCloser, error) {
	return s.client.Bucket(s.bucket).Object(gcsPrefix + version + "/" + file).NewReader(ctx)
}

// Create writes one file of one version. Closing commits it.
func (s *GCSStore) Create(ctx context.Context, version, file string) (io.WriteCloser, error) {
	return s.client.Bucket(s.bucket).Object(gcsPrefix + version + "/" + file).NewWriter(ctx), nil
}

// DeleteVersion removes every object under the version prefix. The
// marker goes first, so a partial delete never leaves a "complete"
// version with missing files.
func (s *GCSStore) DeleteVersion(ctx context.Context, version string) error {
	bucket := s.client.Bucket(s.bucket)
	prefix := gcsPrefix + version + "/"
	if err := bucket.Object(prefix + completeMarker).Delete(ctx); err != nil && !errors.Is(err, storage.ErrObjectNotExist) {
		return fmt.Errorf("delete %s: %w", prefix+completeMarker, err)
	}
	it := bucket.Objects(ctx, &storage.Query{Prefix: prefix})
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("list %s: %w", prefix, err)
		}
		if err := bucket.Object(attrs.Name).Delete(ctx); err != nil && !errors.Is(err, storage.ErrObjectNotExist) {
			return fmt.Errorf("delete %s: %w", attrs.Name, err)
		}
	}
}

// WriteLegalityDiff stores the M-2 marker of one version.
func (s *GCSStore) WriteLegalityDiff(ctx context.Context, version string, rec LegalityDiffRecord) error {
	return writeDiffTo(ctx, s, version, rec)
}

// ReadLegalityDiff returns the marker of one version.
func (s *GCSStore) ReadLegalityDiff(ctx context.Context, version string) (LegalityDiffRecord, bool, error) {
	r, err := s.Open(ctx, version, legalityDiffFile)
	return readDiffFrom(r, err, func(err error) bool { return errors.Is(err, storage.ErrObjectNotExist) })
}

// DirStore stores snapshots under <root>/<version>/<file>.
// Tests and offline local mode use it.
type DirStore struct {
	Root string
}

// errBadComponent rejects a version or file name that could leave Root.
var errBadComponent = errors.New("version or file name must be one plain path component")

// safeComponent rejects names with a separator, "..", or an empty value.
func safeComponent(name string) error {
	if name == "" || name == "." || name == ".." ||
		strings.ContainsAny(name, `/\`) || strings.Contains(name, "..") {
		return fmt.Errorf("%w: %q", errBadComponent, name)
	}
	return nil
}

func (s DirStore) join(version, file string) (string, error) {
	if err := safeComponent(version); err != nil {
		return "", err
	}
	if err := safeComponent(file); err != nil {
		return "", err
	}
	return filepath.Join(s.Root, version, file), nil
}

// LatestVersion returns the newest stored version, or "" when none.
func (s DirStore) LatestVersion(ctx context.Context) (string, error) {
	versions, err := s.ListVersions(ctx)
	if err != nil {
		return "", err
	}
	return latestOf(versions), nil
}

// ListVersions returns every complete version, oldest first.
func (s DirStore) ListVersions(_ context.Context) ([]string, error) {
	return s.listVersions(true)
}

// ListIncompleteVersions returns every version directory with no
// completion marker, oldest first.
func (s DirStore) ListIncompleteVersions(_ context.Context) ([]string, error) {
	return s.listVersions(false)
}

// listVersions lists the version directories that hold a completion
// marker, or the ones that do not.
func (s DirStore) listVersions(complete bool) ([]string, error) {
	entries, err := os.ReadDir(s.Root)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var versions []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		_, err := os.Stat(filepath.Join(s.Root, e.Name(), completeMarker))
		if (err == nil) == complete {
			versions = append(versions, e.Name())
		}
	}
	sort.Strings(versions)
	return versions, nil
}

// Open reads one file of one version.
func (s DirStore) Open(_ context.Context, version, file string) (io.ReadCloser, error) {
	p, err := s.join(version, file)
	if err != nil {
		return nil, err
	}
	return os.Open(p)
}

// Create writes one file of one version.
func (s DirStore) Create(_ context.Context, version, file string) (io.WriteCloser, error) {
	p, err := s.join(version, file)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// Finalize marks a version complete.
func (s DirStore) Finalize(_ context.Context, version string) error {
	p, err := s.join(version, completeMarker)
	if err != nil {
		return err
	}
	return os.WriteFile(p, []byte("ok"), 0o644)
}

// DeleteVersion removes the version directory.
func (s DirStore) DeleteVersion(_ context.Context, version string) error {
	if err := safeComponent(version); err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(s.Root, version))
}

// WriteLegalityDiff stores the M-2 marker of one version.
func (s DirStore) WriteLegalityDiff(ctx context.Context, version string, rec LegalityDiffRecord) error {
	return writeDiffTo(ctx, s, version, rec)
}

// ReadLegalityDiff returns the marker of one version.
func (s DirStore) ReadLegalityDiff(ctx context.Context, version string) (LegalityDiffRecord, bool, error) {
	r, err := s.Open(ctx, version, legalityDiffFile)
	return readDiffFrom(r, err, os.IsNotExist)
}
