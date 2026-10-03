package cards

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"slices"

	"cloud.google.com/go/storage"
)

// newCardsFile is the marker of the new cards of one version. Prune
// keeps it with its version, because it lives under the version prefix.
const newCardsFile = "new_cards.json"

// NewCardsRecord is the marker of the cards that one snapshot made legal
// for the first time (D-1091). The new-cards pass reads it, and it writes
// the end of the pass into it, so a pass that failed or stopped runs
// again on the next run of the job.
type NewCardsRecord struct {
	// SnapshotAsOf is the snapshot time, RFC 3339.
	SnapshotAsOf string `json:"snapshot_as_of"`
	// Cards are the oracle ids of the new cards, sorted.
	Cards []string `json:"cards"`
	// Pass is the time the new-cards pass ended, RFC 3339, or "" when no
	// pass ended yet.
	Pass string `json:"pass,omitempty"`
	// Decks counts the decks that the pass gave new cards.
	Decks int `json:"decks,omitempty"`
}

// NewlyLegal returns the oracle ids, sorted, that the after legalities
// hold as legal in a format of the app, and that the before legalities do
// not hold, or hold as legal in no format at all (D-1091). Scryfall holds
// a preview card as not_legal in each format until its set releases. So a
// card enters the snapshot at its preview, and it becomes new here at its
// release. A reprint was legal before. An unban was legal before in some
// format, Legacy or Vintage at the least, so it is not new either.
func NewlyLegal(before, after map[string]map[string]string) []string {
	var out []string
	for id, a := range after {
		if !legalIn(a, DiffFormats) {
			continue
		}
		if b, ok := before[id]; ok && legalIn(b, nil) {
			continue
		}
		out = append(out, id)
	}
	slices.Sort(out)
	return out
}

// legalIn reports whether a card is legal or restricted in one of the
// formats, or in any format when formats is nil.
func legalIn(l map[string]string, formats []string) bool {
	for k, s := range l {
		if (s == "legal" || s == "restricted") && (formats == nil || slices.Contains(formats, k)) {
			return true
		}
	}
	return false
}

// CompareVersions reads the legalities of two stored versions one time,
// and it returns the count of LegalityDiff and the cards of NewlyLegal.
// The worker calls it once for each new version, so the snapshot job
// reads each version once.
func CompareVersions(ctx context.Context, store Store, oldVersion, newVersion string) (changed int, newCards []string, err error) {
	before, err := loadLegalities(ctx, store, oldVersion)
	if err != nil {
		return 0, nil, err
	}
	after, err := loadLegalities(ctx, store, newVersion)
	if err != nil {
		return 0, nil, err
	}
	return countChanged(before, after), NewlyLegal(before, after), nil
}

// WriteNewCards stores the new-cards marker of one version.
func WriteNewCards(ctx context.Context, s Store, version string, rec NewCardsRecord) error {
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	w, err := s.Create(ctx, version, newCardsFile)
	if err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		_ = w.Close()
		return err
	}
	return w.Close()
}

// ReadNewCards returns the new-cards marker of one version, or false when
// the version has none.
func ReadNewCards(ctx context.Context, s Store, version string) (NewCardsRecord, bool, error) {
	r, err := s.Open(ctx, version, newCardsFile)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) || errors.Is(err, storage.ErrObjectNotExist) {
			return NewCardsRecord{}, false, nil
		}
		return NewCardsRecord{}, false, err
	}
	defer func() { _ = r.Close() }()
	var rec NewCardsRecord
	if err := json.NewDecoder(r).Decode(&rec); err != nil {
		return NewCardsRecord{}, false, fmt.Errorf("%s/%s: %w", version, newCardsFile, err)
	}
	return rec, true, nil
}

// PendingMarker is a new-cards marker with no ended pass, and its version.
type PendingMarker struct {
	Version string
	Record  NewCardsRecord
}

// PendingNewCards returns each new-cards marker with no ended pass,
// oldest first. A newer marker holds only the cards that are new since
// the version before it, so an older marker that a stopped pass left
// must still run (D-1091).
func PendingNewCards(ctx context.Context, s Store) ([]PendingMarker, error) {
	versions, err := s.ListVersions(ctx)
	if err != nil {
		return nil, err
	}
	var out []PendingMarker
	for _, v := range versions {
		rec, ok, err := ReadNewCards(ctx, s, v)
		if err != nil {
			return nil, err
		}
		if ok && rec.Pass == "" {
			out = append(out, PendingMarker{Version: v, Record: rec})
		}
	}
	return out, nil
}
