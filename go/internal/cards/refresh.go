package cards

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"sort"
	"time"

	"github.com/nkramber/decktome/go/internal/scryfall"
)

// bulkTypeByFile maps a snapshot file to its Scryfall bulk type.
var bulkTypeByFile = map[string]string{
	"oracle_cards.jsonl.gz":  "oracle_cards",
	"default_cards.jsonl.gz": "default_cards",
	"oracle_tags.jsonl.gz":   "oracle_tags",
	RulingsFile:              "rulings",
}

// KeepVersions is how many complete snapshots stay in the store after a
// refresh (C-12). Three keeps one for the api, one for a legality diff,
// and one spare for a rollback.
const KeepVersions = 3

// DownloadTimeout bounds one bulk-file transfer. The largest file is
// about 80 MB, so 25 minutes covers a slow link (C-14).
const DownloadTimeout = 25 * time.Minute

// Refresh downloads the bulk files into the store when Scryfall has a
// newer oracle_cards file than the newest stored version. The three bulk
// files must share one UTC date, or the cycle skips and tries later.
// After a successful download it prunes to KeepVersions.
// It returns the current version either way.
func Refresh(ctx context.Context, client *scryfall.Client, store Store, logger *slog.Logger) (string, error) {
	files, err := client.BulkFiles(ctx)
	if err != nil {
		return "", err
	}
	current, err := store.LatestVersion(ctx)
	if err != nil {
		return "", err
	}
	var bulk []scryfall.BulkFile
	for _, name := range SnapshotFiles {
		f, ok := files[bulkTypeByFile[name]]
		if !ok {
			return "", fmt.Errorf("scryfall bulk catalog has no %s", bulkTypeByFile[name])
		}
		bulk = append(bulk, f)
	}
	oracle := bulk[0]
	remote := VersionFor(oracle.UpdatedAt)
	if !newerVersion(remote, current) {
		logger.Info("cards refresh: snapshot current", "version", current)
		return current, nil
	}
	if !sameUTCDate(bulk) {
		logger.Warn("cards refresh: bulk files span two dates, wait for the next cycle",
			"oracle_cards", bulk[0].UpdatedAt, "default_cards", bulk[1].UpdatedAt, "oracle_tags", bulk[2].UpdatedAt)
		return current, nil
	}
	logger.Info("cards refresh: downloading", "from", current, "to", remote)
	for i, name := range SnapshotFiles {
		if err := copyBulk(ctx, client, store, remote, name, bulk[i].DownloadURI); err != nil {
			return "", err
		}
	}
	// The set file is not a bulk file. It comes from the /sets endpoint,
	// and it is the only source of a set family (D-377). It goes in
	// before the marker, so a complete version always holds it.
	if err := copySets(ctx, client, store, remote); err != nil {
		return "", err
	}
	if err := store.Finalize(ctx, remote); err != nil {
		return "", fmt.Errorf("finalize %s: %w", remote, err)
	}
	logger.Info("cards refresh: done", "version", remote)
	if err := Prune(ctx, store, KeepVersions, logger); err != nil {
		// The new snapshot is complete. A failed prune costs storage, not
		// correctness, so it logs and does not fail the refresh.
		logger.Error("cards refresh: prune failed", "err", err)
	}
	return remote, nil
}

// copySets fetches the set table and stores it in one version.
func copySets(ctx context.Context, client *scryfall.Client, store Store, version string) error {
	ctx, cancel := context.WithTimeout(ctx, DownloadTimeout)
	defer cancel()
	rows, err := client.Sets(ctx)
	if err != nil {
		return err
	}
	wctx, wcancel := context.WithCancel(ctx)
	defer wcancel()
	w, err := store.Create(wctx, version, SetsFile)
	if err != nil {
		return err
	}
	gz := gzip.NewWriter(w)
	if err := EncodeSets(gz, SetRowsFrom(rows)); err != nil {
		wcancel()
		_ = w.Close()
		return fmt.Errorf("store %s/%s: %w", version, SetsFile, err)
	}
	if err := gz.Close(); err != nil {
		wcancel()
		_ = w.Close()
		return fmt.Errorf("store %s/%s: %w", version, SetsFile, err)
	}
	return w.Close()
}

// BackfillSets writes the set file into the newest complete version when
// that version has none. A snapshot stored before the set file existed
// then gains the family links without a whole re-download (D-377). It
// returns true when it wrote the file.
func BackfillSets(ctx context.Context, client *scryfall.Client, store Store, logger *slog.Logger) (bool, error) {
	version, err := store.LatestVersion(ctx)
	if err != nil || version == "" {
		return false, err
	}
	r, err := store.Open(ctx, version, SetsFile)
	if err == nil {
		_ = r.Close()
		return false, nil
	}
	logger.Info("cards refresh: the newest snapshot holds no set file, fetching it", "version", version)
	if err := copySets(ctx, client, store, version); err != nil {
		return false, err
	}
	logger.Info("cards refresh: set file written", "version", version)
	return true, nil
}

// newerVersion reports whether remote is strictly newer than current.
// It compares parsed times, not strings. An unparsable current version
// counts as older, so a bad store entry never blocks a refresh.
func newerVersion(remote, current string) bool {
	if current == "" {
		return true
	}
	remoteT, err := VersionTime(remote)
	if err != nil {
		return false
	}
	currentT, err := VersionTime(current)
	if err != nil {
		return true
	}
	return remoteT.After(currentT)
}

// sameUTCDate reports whether every bulk file was updated on one UTC day.
// Scryfall regenerates the files at close but not equal times. A mixed
// set could pair new Oracle text with old printings.
func sameUTCDate(files []scryfall.BulkFile) bool {
	if len(files) == 0 {
		return true
	}
	day := files[0].UpdatedAt.UTC().Format("2006-01-02")
	for _, f := range files[1:] {
		if f.UpdatedAt.UTC().Format("2006-01-02") != day {
			return false
		}
	}
	return true
}

func copyBulk(ctx context.Context, client *scryfall.Client, store Store, version, name, uri string) error {
	ctx, cancel := context.WithTimeout(ctx, DownloadTimeout)
	defer cancel()
	body, err := client.Download(ctx, uri)
	if err != nil {
		return err
	}
	defer func() { _ = body.Close() }()
	// The writer gets its own context. A Close after a copy error would
	// commit a truncated object, so the context is canceled first.
	wctx, wcancel := context.WithCancel(ctx)
	defer wcancel()
	w, err := store.Create(wctx, version, name)
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, body); err != nil {
		wcancel()
		_ = w.Close()
		return fmt.Errorf("store %s/%s: %w", version, name, err)
	}
	return w.Close()
}

// Prune deletes every complete version except the newest keep. It also
// deletes every incomplete version that is older than the newest complete
// one, because a download that old is a failed one. An incomplete
// version newer than that stays: it can be a download in progress.
func Prune(ctx context.Context, store Store, keep int, logger *slog.Logger) error {
	versions, err := store.ListVersions(ctx)
	if err != nil {
		return err
	}
	sortByVersionTime(versions)
	if keep < 1 {
		keep = 1
	}
	for len(versions) > keep {
		old := versions[0]
		versions = versions[1:]
		if err := store.DeleteVersion(ctx, old); err != nil {
			return fmt.Errorf("prune %s: %w", old, err)
		}
		logger.Info("cards refresh: pruned old snapshot", "version", old)
	}
	if len(versions) == 0 {
		return nil
	}
	newest, err := VersionTime(versions[len(versions)-1])
	if err != nil {
		return nil
	}
	incomplete, err := store.ListIncompleteVersions(ctx)
	if err != nil {
		return err
	}
	for _, v := range incomplete {
		t, err := VersionTime(v)
		if err != nil || !t.Before(newest) {
			continue
		}
		if err := store.DeleteVersion(ctx, v); err != nil {
			return fmt.Errorf("prune incomplete %s: %w", v, err)
		}
		logger.Info("cards refresh: pruned incomplete snapshot", "version", v)
	}
	return nil
}

// sortByVersionTime orders versions oldest first by their parsed time.
func sortByVersionTime(versions []string) {
	sort.SliceStable(versions, func(i, j int) bool {
		ti, _ := VersionTime(versions[i])
		tj, _ := VersionTime(versions[j])
		return ti.Before(tj)
	})
}

// LegalityDiff counts the cards whose legalities changed between two
// stored versions (C-2). It reads only oracle_id and legalities, so it
// costs a fraction of a full index load. A card that is present in one
// version only counts as changed.
func LegalityDiff(ctx context.Context, store Store, oldVersion, newVersion string) (changed int, err error) {
	before, err := loadLegalities(ctx, store, oldVersion)
	if err != nil {
		return 0, err
	}
	after, err := loadLegalities(ctx, store, newVersion)
	if err != nil {
		return 0, err
	}
	for id, a := range after {
		b, ok := before[id]
		if !ok || !sameLegalities(a, b) {
			changed++
		}
	}
	for id := range before {
		if _, ok := after[id]; !ok {
			changed++
		}
	}
	return changed, nil
}

func sameLegalities(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func loadLegalities(ctx context.Context, store Store, version string) (map[string]map[string]string, error) {
	r, err := store.Open(ctx, version, "oracle_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Close() }()
	out := map[string]map[string]string{}
	err = readLines(r, "oracle_cards.jsonl.gz", func(line []byte) error {
		var raw struct {
			OracleID   string            `json:"oracle_id"`
			Legalities map[string]string `json:"legalities"`
			CardFaces  []struct {
				OracleID string `json:"oracle_id"`
			} `json:"card_faces"`
		}
		if err := json.Unmarshal(line, &raw); err != nil {
			return err
		}
		if raw.OracleID == "" && len(raw.CardFaces) > 0 {
			raw.OracleID = raw.CardFaces[0].OracleID
		}
		if raw.OracleID != "" {
			out[raw.OracleID] = raw.Legalities
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s/%s: %w", version, "oracle_cards.jsonl.gz", err)
	}
	return out, nil
}

// LoadIndex builds an Index from the newest stored snapshot.
// It returns nil with no error when the store is empty.
func LoadIndex(ctx context.Context, store Store, logger *slog.Logger) (*Index, error) {
	version, err := store.LatestVersion(ctx)
	if err != nil {
		return nil, err
	}
	if version == "" {
		return nil, nil
	}
	start := time.Now()
	asOf, err := VersionTime(version)
	if err != nil {
		return nil, fmt.Errorf("bad snapshot version %q: %w", version, err)
	}
	oc, err := store.Open(ctx, version, "oracle_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	defer func() { _ = oc.Close() }()
	parsed, stats, err := LoadCardsStats(oc, "oracle_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	if stats.NoOracleID > 0 {
		logger.Warn("cards index: cards with no oracle_id skipped", "count", stats.NoOracleID)
	}
	dc, err := store.Open(ctx, version, "default_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	defer func() { _ = dc.Close() }()
	printings, err := LoadPrintings(dc, "default_cards.jsonl.gz")
	if err != nil {
		return nil, err
	}
	tf, err := store.Open(ctx, version, "oracle_tags.jsonl.gz")
	if err != nil {
		return nil, err
	}
	defer func() { _ = tf.Close() }()
	tags, err := LoadTags(tf, "oracle_tags.jsonl.gz")
	if err != nil {
		return nil, err
	}
	// A snapshot stored before the set file existed holds none. The
	// index then derives a table from the printings, which carries no
	// family link, and the log says so (D-377).
	var sets []SetInfo
	sf, err := store.Open(ctx, version, SetsFile)
	if err != nil {
		logger.Warn("cards index: the snapshot holds no set file, so no set family resolves",
			"version", version, "file", SetsFile, "err", err)
	} else {
		sets, err = LoadSets(sf, SetsFile)
		_ = sf.Close()
		if err != nil {
			return nil, fmt.Errorf("%s/%s: %w", version, SetsFile, err)
		}
	}
	// A snapshot stored before the rulings file existed holds none, and
	// the card detail then shows no ruling (PR-20).
	var rulings map[string][]Ruling
	rf, err := store.Open(ctx, version, RulingsFile)
	if err != nil {
		logger.Warn("cards index: the snapshot holds no rulings file, so no card shows a ruling",
			"version", version, "file", RulingsFile, "err", err)
	} else {
		rulings, err = LoadRulings(rf, RulingsFile)
		_ = rf.Close()
		if err != nil {
			return nil, fmt.Errorf("%s/%s: %w", version, RulingsFile, err)
		}
	}
	idx := NewIndex(parsed, printings, tags, asOf, WithSets(sets), WithRulings(rulings))
	col := idx.Collisions()
	logger.Info("cards index loaded", "version", version, "cards", idx.Len(),
		"printings", len(printings), "tags", tags.Len(),
		"name_collisions", col.FullNames, "face_name_collisions", col.FaceNames,
		"paper_swaps", idx.PaperSwaps(),
		"sets", idx.Sets().Len(), "sets_derived", idx.Sets().Derived(),
		"rulings", len(rulings),
		"took", time.Since(start).String())
	return idx, nil
}
