// Package evalrun holds the run header and the long-format rows every
// gate writes (PR-15). The header names the run: the suite, the run id,
// the date, the commit, the resolved model and effort per role, the
// snapshot date, the prompt versions, and the cost. A row holds one
// measurement of one item. A gate writes its document as before, and it
// writes the run as JSONL beside it, so a compare reads rows and never
// prose (lessons 3 and 6 of the roadmap).
package evalrun

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"sort"
	"strings"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// The kinds of a row. A gate row decides the verdict, and an info row
// is reported and never gates (the metric-suffix rule of the eval
// patterns note). A suite with no gate row is not evaluated, and a
// compare never calls it a pass.
const (
	KindGate = "gate"
	KindInfo = "info"
)

// Model is the resolved model of one role: the one that ran, not the
// one requested (lesson 3).
type Model struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Effort   string `json:"effort,omitempty"`
}

// Header is the fingerprint of one run. A compare refuses two runs whose
// fingerprints name different epochs.
type Header struct {
	Suite  string `json:"suite"`
	RunID  string `json:"run_id"`
	Date   string `json:"date"`
	Commit string `json:"commit,omitempty"`
	// Roles holds the roles the suite called, by role name.
	Roles map[string]Model `json:"roles,omitempty"`
	// Snapshot is the card snapshot date, empty when the suite reads none.
	Snapshot string `json:"snapshot,omitempty"`
	// Prompts holds the prompt version per prompt name.
	Prompts map[string]int `json:"prompts,omitempty"`
	// Versions holds every other version the suite rests on: the quality
	// model, the precon table, the conversation set.
	Versions map[string]string `json:"versions,omitempty"`
	Calls    int               `json:"calls"`
	CostUSD  *float64          `json:"cost_usd,omitempty"`
	Seconds  float64           `json:"seconds"`
	// Verdict is PASS or FAIL, or empty when the suite decides nothing.
	Verdict string `json:"verdict,omitempty"`
	// Only is the -only flag of a partial run, and empty on a whole run.
	// A partial run reads the bars of its own items alone. It never
	// stands for the suite: the check skips it, and it is no baseline.
	Only string `json:"only,omitempty"`
	// Lower names the metrics where a lower value is better, so a
	// compare of this file knows which way is worse. Every other metric
	// reads a higher value as better.
	Lower []string `json:"lower,omitempty"`
	// Note is provenance a compare never reads, for example the document
	// an imported run came from.
	Note string `json:"note,omitempty"`
}

// Row is one measurement of one item.
type Row struct {
	Item   string  `json:"item"`
	Metric string  `json:"metric"`
	Value  float64 `json:"value"`
	Kind   string  `json:"kind"`
	Detail string  `json:"detail,omitempty"`
}

// Run is a header with its rows.
type Run struct {
	Header Header `json:"header"`
	Rows   []Row  `json:"rows"`
}

// New starts a run of one suite, dated today, on the commit of the tree.
func New(suite, runID string) *Run {
	return &Run{Header: Header{
		Suite: suite, RunID: runID,
		Date:     time.Now().UTC().Format("2006-01-02"),
		Commit:   Commit(),
		Roles:    map[string]Model{},
		Prompts:  map[string]int{},
		Versions: map[string]string{},
	}}
}

// Gate adds a row that decides the verdict.
func (r *Run) Gate(item, metric string, value float64, detail string) {
	r.Rows = append(r.Rows, Row{Item: item, Metric: metric, Value: value, Kind: KindGate, Detail: detail})
}

// Info adds a row that is reported and never gates.
func (r *Run) Info(item, metric string, value float64, detail string) {
	r.Rows = append(r.Rows, Row{Item: item, Metric: metric, Value: value, Kind: KindInfo, Detail: detail})
}

// SetRoles records the resolved model of each named role.
func (r *Run) SetRoles(cfg *llm.Config, roles ...llm.Role) {
	if cfg == nil {
		return
	}
	for _, role := range roles {
		spec, ok := cfg.Roles[role]
		if !ok {
			continue
		}
		r.Header.Roles[string(role)] = Model{Provider: spec.Provider, Model: spec.Model, Effort: spec.Effort}
	}
}

// SetSnapshot records the card snapshot date.
func (r *Run) SetSnapshot(asOf time.Time) {
	if !asOf.IsZero() {
		r.Header.Snapshot = asOf.Format("2006-01-02")
	}
}

// Finish records the cost, the time, and the verdict.
func (r *Run) Finish(rep llm.Report, took time.Duration, verdict string) {
	r.Header.Calls = rep.Calls
	r.Header.CostUSD = rep.CostUSD
	r.Header.Seconds = took.Seconds()
	r.Header.Verdict = verdict
}

// Partial reports whether the run covered a part of its suite, under
// -only or a count. A partial run never stands for the suite.
func (h Header) Partial() bool { return h.Only != "" }

// Gated reports whether the run holds a gate row. A run with none is
// not evaluated.
func (r *Run) Gated() bool {
	for _, row := range r.Rows {
		if row.Kind == KindGate {
			return true
		}
	}
	return false
}

// Write writes the run as JSONL: the header first, then one row a line.
func Write(w io.Writer, r *Run) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(r.Header); err != nil {
		return err
	}
	for i := range r.Rows {
		if err := enc.Encode(r.Rows[i]); err != nil {
			return err
		}
	}
	return nil
}

// Read reads a run written by Write.
func Read(rd io.Reader) (*Run, error) {
	sc := bufio.NewScanner(rd)
	sc.Buffer(make([]byte, 0, 64<<10), 8<<20)
	var out *Run
	for sc.Scan() {
		line := bytes.TrimSpace(sc.Bytes())
		if len(line) == 0 {
			continue
		}
		if out == nil {
			var h Header
			if err := json.Unmarshal(line, &h); err != nil {
				return nil, fmt.Errorf("evalrun: header: %w", err)
			}
			if h.Suite == "" {
				return nil, errors.New("evalrun: the first line names no suite")
			}
			out = &Run{Header: h}
			continue
		}
		var row Row
		if err := json.Unmarshal(line, &row); err != nil {
			return nil, fmt.Errorf("evalrun: row: %w", err)
		}
		out.Rows = append(out.Rows, row)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if out == nil {
		return nil, errors.New("evalrun: the file is empty")
	}
	return out, nil
}

// WriteFile writes the run to a path. An empty path writes nothing. An
// existing file is never overwritten (D-65), and the directory is made.
func WriteFile(path string, r *Run) error {
	if path == "" {
		return nil
	}
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("%s exists, and a result is never overwritten (D-65): name a new file", path)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := Write(&buf, r); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o600)
}

// ReadFile reads a run from a path.
func ReadFile(path string) (*Run, error) {
	f, err := os.Open(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	return Read(f)
}

// RunID is the run id of a path: the file name without its extension.
func RunID(path string) string {
	if path == "" {
		return ""
	}
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// Markdown writes the "## Run" block every gate document carries, the
// same lines in the same order. It ends on the last line, so a caller
// adds its own lines under it and then the blank line.
func (r *Run) Markdown(w io.Writer) {
	h := r.Header
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	p("## Run\n\n")
	commit := h.Commit
	if commit == "" {
		commit = "unknown"
	}
	p("- Suite `%s`, run `%s`, on %s, commit `%s`.\n", h.Suite, orWord(h.RunID, "unnamed"), h.Date, commit)
	if h.Partial() {
		p("- Partial run over `%s`. It reads the bars of its own items, and it never stands for the suite.\n", h.Only)
	}
	if len(h.Roles) > 0 {
		var parts []string
		for _, role := range sortedKeys(h.Roles) {
			m := h.Roles[role]
			effort := ""
			if m.Effort != "" {
				effort = ", effort " + m.Effort
			}
			parts = append(parts, fmt.Sprintf("%s on `%s` (%s%s)", role, m.Model, m.Provider, effort))
		}
		p("- Roles: %s.\n", strings.Join(parts, ", "))
	} else {
		p("- Roles: none, no provider call.\n")
	}
	var facts []string
	if h.Snapshot != "" {
		facts = append(facts, "card snapshot "+h.Snapshot)
	}
	for _, name := range sortedKeys(h.Prompts) {
		facts = append(facts, fmt.Sprintf("%s prompt version %d", name, h.Prompts[name]))
	}
	for _, name := range sortedKeys(h.Versions) {
		facts = append(facts, fmt.Sprintf("%s `%s`", name, h.Versions[name]))
	}
	if len(facts) > 0 {
		p("- Versions: %s.\n", strings.Join(facts, ", "))
	}
	cost := "unpriced"
	if h.CostUSD != nil {
		cost = fmt.Sprintf("$%.4f", *h.CostUSD)
	}
	p("- Calls: %d. Cost: %s. Time: %.0f seconds.\n", h.Calls, cost, h.Seconds)
	if h.Note != "" {
		p("- Note: %s\n", h.Note)
	}
}

func orWord(s, word string) string {
	if s == "" {
		return word
	}
	return s
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Commit names the commit of the tree, short. It asks git first, then
// the build info, and it answers empty when neither knows.
func Commit() string {
	if out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output(); err == nil {
		return strings.TrimSpace(string(out))
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, s := range info.Settings {
			if s.Key == "vcs.revision" && len(s.Value) >= 7 {
				return s.Value[:7]
			}
		}
	}
	return ""
}
