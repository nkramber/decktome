package meta

import (
	"bufio"
	"bytes"
	"compress/gzip"
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

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"

	"github.com/nkramber/decktome/go/internal/gzstore"
)

// ObjectStore holds the raw pages, the normalized lists, and the
// fitted model. Names are slash paths under the meta prefix. The GCS
// bucket is the one of the card snapshot, and a directory serves the
// tests and the offline mode.
type ObjectStore interface {
	// Put writes one object whole.
	Put(ctx context.Context, name string, data []byte) error
	// Get reads one object. The second value is false when none exists.
	Get(ctx context.Context, name string) ([]byte, bool, error)
	// List names every object under a prefix, sorted.
	List(ctx context.Context, prefix string) ([]string, error)
}

// Prefix is the bucket prefix of everything the package stores.
const Prefix = "meta/"

// The layout under Prefix.
const (
	// RawPrefix holds one gzipped page per fetch: raw/<source>/<key>.gz.
	RawPrefix = Prefix + "raw/"
	// ListsPrefix holds the normalized lists: lists/<format>/<YYYY-MM>.jsonl.gz.
	ListsPrefix = Prefix + "lists/"
	// PreconsPrefix holds one table per MTGJSON version:
	// precons/<version>/precons.jsonl.gz.
	PreconsPrefix = Prefix + "precons/"
	// CommandersPrefix holds the commander reads per day:
	// commanders/<YYYY-MM-DD>.jsonl.gz.
	CommandersPrefix = Prefix + "commanders/"
	// ModelPrefix holds the fitted model per version:
	// model/<version>/quality.json.gz and its complete marker.
	ModelPrefix = Prefix + "model/"
	// ModelFile is the model object of a version.
	ModelFile = "quality.json.gz"
	// CompleteMarker is written last, as the card snapshot does. A model
	// version without it is mid-write, and LatestModel never returns it.
	CompleteMarker = "complete"
)

// RawName is the object name of one raw page.
func RawName(source, key string) string {
	return RawPrefix + source + "/" + strings.ReplaceAll(key, "/", "_") + ".gz"
}

// PutRaw stores one page gzipped.
func PutRaw(ctx context.Context, s ObjectStore, source, key string, page []byte) error {
	data, err := gzstore.Marshal(page)
	if err != nil {
		return err
	}
	return s.Put(ctx, RawName(source, key), data)
}

// GetRaw reads one stored page.
func GetRaw(ctx context.Context, s ObjectStore, source, key string) ([]byte, bool, error) {
	data, ok, err := s.Get(ctx, RawName(source, key))
	if err != nil || !ok {
		return nil, ok, err
	}
	page, err := gzstore.Unmarshal(data)
	return page, true, err
}

// HasRaw says whether a page is stored.
func HasRaw(ctx context.Context, s ObjectStore, source, key string) (bool, error) {
	_, ok, err := s.Get(ctx, RawName(source, key))
	return ok, err
}

// ListsName is the object of one format and month.
func ListsName(format, month string) string {
	return ListsPrefix + format + "/" + month + ".jsonl.gz"
}

// ReadLists reads the lists of one object. A missing object answers
// nil and no error.
func ReadLists(ctx context.Context, s ObjectStore, name string) ([]List, error) {
	data, ok, err := s.Get(ctx, name)
	if err != nil || !ok {
		return nil, err
	}
	raw, err := inflate(data)
	if err != nil {
		return nil, fmt.Errorf("meta: %s: %w", name, err)
	}
	return DecodeLists(bytes.NewReader(raw))
}

// MaxInflate bounds one list or table object. A month of cEDH lists
// from Topdeck.gg inflates past the 16 MB limit of gzstore (2026-09-02),
// and these objects are the worker's own.
const MaxInflate = 1 << 30

// inflate reads a gzipped object under MaxInflate.
func inflate(data []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer func() { _ = zr.Close() }()
	out, err := io.ReadAll(io.LimitReader(zr, MaxInflate+1))
	if err != nil {
		return nil, err
	}
	if len(out) > MaxInflate {
		return nil, fmt.Errorf("meta: object inflates past %d bytes", MaxInflate)
	}
	return out, nil
}

// DecodeLists reads JSON lines.
func DecodeLists(r io.Reader) ([]List, error) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64<<10), 8<<20)
	var out []List
	for sc.Scan() {
		line := bytes.TrimSpace(sc.Bytes())
		if len(line) == 0 {
			continue
		}
		var l List
		if err := json.Unmarshal(line, &l); err != nil {
			return nil, fmt.Errorf("meta: list row: %w", err)
		}
		out = append(out, l)
	}
	return out, sc.Err()
}

// EncodeLists writes JSON lines.
func EncodeLists(w io.Writer, lists []List) error {
	enc := json.NewEncoder(w)
	for i := range lists {
		if err := enc.Encode(&lists[i]); err != nil {
			return err
		}
	}
	return nil
}

// MergeLists adds lists to their month objects. A list already stored
// under its key is replaced, so a re-read of a page changes nothing.
// It answers the count of new keys.
func MergeLists(ctx context.Context, s ObjectStore, lists []List) (int, error) {
	byName := map[string][]List{}
	for _, l := range lists {
		if l.Format == "" || l.Month() == "" {
			continue
		}
		name := ListsName(l.Format, l.Month())
		byName[name] = append(byName[name], l)
	}
	names := make([]string, 0, len(byName))
	for n := range byName {
		names = append(names, n)
	}
	sort.Strings(names)
	added := 0
	for _, name := range names {
		have, err := ReadLists(ctx, s, name)
		if err != nil {
			return added, err
		}
		index := map[string]int{}
		for i, l := range have {
			index[l.Key()] = i
		}
		for _, l := range byName[name] {
			if i, ok := index[l.Key()]; ok {
				have[i] = l
				continue
			}
			index[l.Key()] = len(have)
			have = append(have, l)
			added++
		}
		sort.SliceStable(have, func(i, j int) bool { return have[i].Key() < have[j].Key() })
		var buf bytes.Buffer
		if err := EncodeLists(&buf, have); err != nil {
			return added, err
		}
		data, err := gzstore.Marshal(buf.Bytes())
		if err != nil {
			return added, err
		}
		if err := s.Put(ctx, name, data); err != nil {
			return added, err
		}
	}
	return added, nil
}

// AllLists reads every stored list of a format word. The caller filters
// by Covers when a fit reads two words.
func AllLists(ctx context.Context, s ObjectStore, format string) ([]List, error) {
	names, err := s.List(ctx, ListsPrefix+format+"/")
	if err != nil {
		return nil, err
	}
	var out []List
	for _, name := range names {
		lists, err := ReadLists(ctx, s, name)
		if err != nil {
			return nil, err
		}
		out = append(out, lists...)
	}
	return out, nil
}

// PreconsName is the table object of one MTGJSON version.
func PreconsName(version string) string { return PreconsPrefix + version + "/precons.jsonl.gz" }

// WritePrecons stores a table.
func WritePrecons(ctx context.Context, s ObjectStore, version string, table []Precon) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for i := range table {
		if err := enc.Encode(&table[i]); err != nil {
			return err
		}
	}
	data, err := gzstore.Marshal(buf.Bytes())
	if err != nil {
		return err
	}
	return s.Put(ctx, PreconsName(version), data)
}

// ReadPrecons reads the table of one version, or nil when none.
func ReadPrecons(ctx context.Context, s ObjectStore, version string) ([]Precon, error) {
	data, ok, err := s.Get(ctx, PreconsName(version))
	if err != nil || !ok {
		return nil, err
	}
	raw, err := inflate(data)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(bytes.NewReader(raw))
	sc.Buffer(make([]byte, 0, 64<<10), 8<<20)
	var out []Precon
	for sc.Scan() {
		if len(bytes.TrimSpace(sc.Bytes())) == 0 {
			continue
		}
		var p Precon
		if err := json.Unmarshal(sc.Bytes(), &p); err != nil {
			return nil, fmt.Errorf("meta: precon row: %w", err)
		}
		out = append(out, p)
	}
	return out, sc.Err()
}

// LatestPreconsVersion names the newest stored table, or "".
func LatestPreconsVersion(ctx context.Context, s ObjectStore) (string, error) {
	names, err := s.List(ctx, PreconsPrefix)
	if err != nil {
		return "", err
	}
	latest := ""
	for _, n := range names {
		rest := strings.TrimPrefix(n, PreconsPrefix)
		v, _, ok := strings.Cut(rest, "/")
		if ok && v > latest {
			latest = v
		}
	}
	return latest, nil
}

// CommandersName is the commander reads of one day.
func CommandersName(day string) string { return CommandersPrefix + day + ".jsonl.gz" }

// WriteCommanders stores the reads of one day, merged with the ones
// already there by slug.
func WriteCommanders(ctx context.Context, s ObjectStore, day string, reads []Commander) error {
	have, err := ReadCommanders(ctx, s, day)
	if err != nil {
		return err
	}
	index := map[string]int{}
	for i, c := range have {
		index[c.Slug] = i
	}
	for _, c := range reads {
		if i, ok := index[c.Slug]; ok {
			have[i] = c
			continue
		}
		index[c.Slug] = len(have)
		have = append(have, c)
	}
	sort.Slice(have, func(i, j int) bool { return have[i].Slug < have[j].Slug })
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for i := range have {
		if err := enc.Encode(&have[i]); err != nil {
			return err
		}
	}
	data, err := gzstore.Marshal(buf.Bytes())
	if err != nil {
		return err
	}
	return s.Put(ctx, CommandersName(day), data)
}

// ReadCommanders reads the reads of one day, or nil when none.
func ReadCommanders(ctx context.Context, s ObjectStore, day string) ([]Commander, error) {
	data, ok, err := s.Get(ctx, CommandersName(day))
	if err != nil || !ok {
		return nil, err
	}
	raw, err := inflate(data)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(bytes.NewReader(raw))
	sc.Buffer(make([]byte, 0, 64<<10), 8<<20)
	var out []Commander
	for sc.Scan() {
		if len(bytes.TrimSpace(sc.Bytes())) == 0 {
			continue
		}
		var c Commander
		if err := json.Unmarshal(sc.Bytes(), &c); err != nil {
			return nil, fmt.Errorf("meta: commander row: %w", err)
		}
		out = append(out, c)
	}
	return out, sc.Err()
}

// LatestCommandersDay names the newest stored day, or "".
func LatestCommandersDay(ctx context.Context, s ObjectStore) (string, error) {
	names, err := s.List(ctx, CommandersPrefix)
	if err != nil {
		return "", err
	}
	if len(names) == 0 {
		return "", nil
	}
	last := names[len(names)-1]
	return strings.TrimSuffix(strings.TrimPrefix(last, CommandersPrefix), ".jsonl.gz"), nil
}

// WriteModel stores one model version, the marker last.
func WriteModel(ctx context.Context, s ObjectStore, version string, model []byte) error {
	data, err := gzstore.Marshal(model)
	if err != nil {
		return err
	}
	if err := s.Put(ctx, ModelPrefix+version+"/"+ModelFile, data); err != nil {
		return err
	}
	return s.Put(ctx, ModelPrefix+version+"/"+CompleteMarker, []byte(version))
}

// LatestModel names the newest complete model version, or "".
func LatestModel(ctx context.Context, s ObjectStore) (string, error) {
	names, err := s.List(ctx, ModelPrefix)
	if err != nil {
		return "", err
	}
	latest := ""
	for _, n := range names {
		rest := strings.TrimPrefix(n, ModelPrefix)
		v, file, ok := strings.Cut(rest, "/")
		if ok && file == CompleteMarker && v > latest {
			latest = v
		}
	}
	return latest, nil
}

// ReadModel reads the model bytes of one version.
func ReadModel(ctx context.Context, s ObjectStore, version string) ([]byte, error) {
	data, ok, err := s.Get(ctx, ModelPrefix+version+"/"+ModelFile)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("meta: model %s is absent", version)
	}
	return inflate(data)
}

// DirObjects stores objects under a directory. The tests and the
// offline mode use it.
type DirObjects struct {
	Root string
}

func (d DirObjects) path(name string) (string, error) {
	clean := path.Clean("/" + name)
	if strings.Contains(name, "..") || clean == "/" {
		return "", fmt.Errorf("meta: bad object name %q", name)
	}
	return filepath.Join(d.Root, filepath.FromSlash(clean)), nil
}

// Put writes one object.
func (d DirObjects) Put(_ context.Context, name string, data []byte) error {
	p, err := d.path(name)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	tmp := p + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, p)
}

// Get reads one object.
func (d DirObjects) Get(_ context.Context, name string) ([]byte, bool, error) {
	p, err := d.path(name)
	if err != nil {
		return nil, false, err
	}
	data, err := os.ReadFile(p)
	if errors.Is(err, os.ErrNotExist) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return data, true, nil
}

// List names every object under a prefix, sorted.
func (d DirObjects) List(_ context.Context, prefix string) ([]string, error) {
	root := filepath.Join(d.Root, filepath.FromSlash(prefix))
	var out []string
	err := filepath.WalkDir(root, func(p string, entry os.DirEntry, err error) error {
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return nil
			}
			return err
		}
		if entry.IsDir() || strings.HasSuffix(p, ".tmp") {
			return nil
		}
		rel, err := filepath.Rel(d.Root, p)
		if err != nil {
			return err
		}
		out = append(out, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(out)
	return out, nil
}

// GCSObjects stores objects in the card bucket.
type GCSObjects struct {
	client *storage.Client
	bucket string
}

// NewGCSObjects wraps an existing client. The caller owns the client.
func NewGCSObjects(client *storage.Client, bucket string) *GCSObjects {
	return &GCSObjects{client: client, bucket: bucket}
}

// Put writes one object.
func (g *GCSObjects) Put(ctx context.Context, name string, data []byte) error {
	w := g.client.Bucket(g.bucket).Object(name).NewWriter(ctx)
	if _, err := w.Write(data); err != nil {
		_ = w.Close()
		return fmt.Errorf("meta: write %s: %w", name, err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("meta: write %s: %w", name, err)
	}
	return nil
}

// Get reads one object.
func (g *GCSObjects) Get(ctx context.Context, name string) ([]byte, bool, error) {
	r, err := g.client.Bucket(g.bucket).Object(name).NewReader(ctx)
	if errors.Is(err, storage.ErrObjectNotExist) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("meta: read %s: %w", name, err)
	}
	defer func() { _ = r.Close() }()
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, false, fmt.Errorf("meta: read %s: %w", name, err)
	}
	return data, true, nil
}

// List names every object under a prefix, sorted.
func (g *GCSObjects) List(ctx context.Context, prefix string) ([]string, error) {
	it := g.client.Bucket(g.bucket).Objects(ctx, &storage.Query{Prefix: prefix})
	var out []string
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("meta: list %s: %w", prefix, err)
		}
		out = append(out, attrs.Name)
	}
	sort.Strings(out)
	return out, nil
}
