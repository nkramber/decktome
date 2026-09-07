package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/nkramber/decktome/go/internal/evalrun"
)

// baselinesFile names the accepted run files per suite, relative to the
// run directory. A suite may name a run and its rerun, and the check
// reads them as one (evalrun.Merge).
const baselinesFile = "baselines.json"

type baselines map[string][]string

func readBaselines(dir string) (baselines, error) {
	raw, err := os.ReadFile(filepath.Join(dir, baselinesFile)) // #nosec G304 -- the operator names the directory.
	if errors.Is(err, os.ErrNotExist) {
		return baselines{}, nil
	}
	if err != nil {
		return nil, err
	}
	var bl baselines
	if err := json.Unmarshal(raw, &bl); err != nil {
		return nil, fmt.Errorf("%s: %w", baselinesFile, err)
	}
	return bl, nil
}

func writeBaselines(dir string, bl baselines) error {
	raw, err := json.MarshalIndent(bl, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, baselinesFile), append(raw, '\n'), 0o600)
}

func sortedSuites(bl baselines) []string {
	out := make([]string, 0, len(bl))
	for s := range bl {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

// setBaseline records the run files of one suite. The merged run must
// pass its own bars, or -force says the owner accepts it as it stands.
func setBaseline(w io.Writer, dir, suite string, files []string, force bool) error {
	if suite == "" || len(files) == 0 {
		return errors.New("give -suite and -run, the run files of the baseline")
	}
	merged, err := readRuns(prefixed(dir, files))
	if err != nil {
		return err
	}
	if merged.Header.Suite != suite {
		return fmt.Errorf("the run files are suite %s, not %s", merged.Header.Suite, suite)
	}
	if !merged.Gated() {
		return fmt.Errorf("the run holds no gate row, so it can be no baseline")
	}
	if merged.Header.Verdict != evalrun.VerdictPass && !force {
		return fmt.Errorf("the run reads %s, and a baseline passes its own bars: pass -force to record it as it stands", orUnnamed(merged.Header.Verdict))
	}
	bl, err := readBaselines(dir)
	if err != nil {
		return err
	}
	bl[suite] = files
	if err := writeBaselines(dir, bl); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(w, "baseline of %s: %s (%s, %d rows)\n", suite, strings.Join(files, " + "), merged.Header.Verdict, len(merged.Rows))
	return nil
}

// fileHeader is one run file of the directory with its header and the
// count of its items.
type fileHeader struct {
	name   string
	header evalrun.Header
	items  int
}

// readHeaders reads the header of every run file in the directory.
func readHeaders(dir string) ([]fileHeader, error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var out []fileHeader
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".jsonl" {
			continue
		}
		r, err := evalrun.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		out = append(out, fileHeader{name: e.Name(), header: r.Header, items: countItems(r)})
	}
	return out, nil
}

// newestRun names the newest whole run file of a suite that is newer
// than every run of the baseline, or "" when none. A run older than the
// baseline is history, and the check never reads it. A partial run, one
// of -only or a count, never stands for the suite, so the check skips it
// and partialRuns lists it. Newest reads the date first, then the number
// at the end of the run id, then the id.
func newestRun(headers []fileHeader, suite string, baseline []string) string {
	best := ""
	var bestHeader evalrun.Header
	for _, fh := range sinceBaseline(headers, suite, baseline) {
		if fh.header.Partial() {
			continue
		}
		if best == "" || runLess(bestHeader, fh.header) {
			best, bestHeader = fh.name, fh.header
		}
	}
	return best
}

// partialRuns names the partial runs of a suite newer than the baseline,
// oldest first. The check lists them and compares none.
func partialRuns(headers []fileHeader, suite string, baseline []string) []fileHeader {
	var out []fileHeader
	for _, fh := range sinceBaseline(headers, suite, baseline) {
		if fh.header.Partial() {
			out = append(out, fh)
		}
	}
	sort.Slice(out, func(i, j int) bool { return runLess(out[i].header, out[j].header) })
	return out
}

// sinceBaseline keeps the runs of a suite that are newer than every run
// of the baseline and not in it.
func sinceBaseline(headers []fileHeader, suite string, baseline []string) []fileHeader {
	inBase := map[string]bool{}
	for _, f := range baseline {
		inBase[f] = true
	}
	var floor evalrun.Header
	hasFloor := false
	for _, fh := range headers {
		if inBase[fh.name] && (!hasFloor || runLess(floor, fh.header)) {
			floor, hasFloor = fh.header, true
		}
	}
	var out []fileHeader
	for _, fh := range headers {
		if fh.header.Suite != suite || inBase[fh.name] {
			continue
		}
		if hasFloor && !runLess(floor, fh.header) {
			continue
		}
		out = append(out, fh)
	}
	return out
}

var trailingNumber = regexp.MustCompile(`(\d+)[a-z]?$`)

// runLess orders two headers: by date, then by the number the run id
// ends with, then by the id.
func runLess(a, b evalrun.Header) bool {
	if a.Date != b.Date {
		return a.Date < b.Date
	}
	na, nb := idNumber(a.RunID), idNumber(b.RunID)
	if na != nb {
		return na < nb
	}
	return a.RunID < b.RunID
}

func idNumber(id string) int {
	m := trailingNumber.FindStringSubmatch(id)
	if m == nil {
		return 0
	}
	n, _ := strconv.Atoi(m[1])
	return n
}
