package evalrun

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// The verdicts of a comparison.
const (
	VerdictPass = "PASS"
	VerdictFail = "FAIL"
	// VerdictNotEvaluated says the next run holds no gate row or no
	// verdict of its own. Observe-only is not pass.
	VerdictNotEvaluated = "NOT EVALUATED"
)

// Flip is one row that moved between two runs.
type Flip struct {
	Item   string  `json:"item"`
	Metric string  `json:"metric"`
	Kind   string  `json:"kind"`
	Before float64 `json:"before"`
	After  float64 `json:"after"`
	// Worse says the row moved against its sense: up on a metric where
	// lower is better, or down on one where higher is better.
	Worse        bool   `json:"worse"`
	BeforeDetail string `json:"before_detail,omitempty"`
	AfterDetail  string `json:"after_detail,omitempty"`
}

// Comparison is the read of two runs of one suite.
type Comparison struct {
	Base, Next Header
	// Epoch names the fingerprint differences: a model, a prompt version,
	// a snapshot, or a stored version that changed. The compare still
	// runs, and the reader decides what the differences mean.
	Epoch []string
	// Flips are the rows that moved, gate rows first, worst first.
	Flips []Flip
	// Gone and New name the items one run holds and the other lacks, for
	// example the 24 items a -only run leaves out.
	Gone, New []string
	Verdict   string
	Reasons   []string
}

// LowerIsBetter marks the metrics of a run whose sense is down, so a
// compare knows which way is worse. Every other metric reads up as
// better.
func (r *Run) LowerIsBetter(metrics ...string) {
	for _, m := range metrics {
		if !contains(r.Header.Lower, m) {
			r.Header.Lower = append(r.Header.Lower, m)
		}
	}
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// Compare reads next against base. margin is how far a gate row may
// move against its sense before the compare calls it a regression, in
// the metric's own unit. A base of nil is the first run: every gate row
// of next is reported and none is a flip.
func Compare(base, next *Run, margin float64) Comparison {
	c := Comparison{Next: next.Header}
	if base != nil {
		c.Base = base.Header
		c.Epoch = epochNotes(base.Header, next.Header)
	}
	if !next.Gated() || next.Header.Verdict == "" {
		c.Verdict = VerdictNotEvaluated
		c.Reasons = append(c.Reasons, "the run holds no gate row or no verdict of its own, so it is observed and not passed")
		return c
	}
	lower := map[string]bool{}
	for _, m := range next.Header.Lower {
		lower[m] = true
	}
	if base != nil {
		for _, m := range base.Header.Lower {
			lower[m] = true
		}
		c.Flips, c.Gone, c.New = flips(base, next, lower, margin)
	}
	worse := 0
	for _, f := range c.Flips {
		if f.Worse && f.Kind == KindGate {
			worse++
		}
	}
	switch {
	case next.Header.Verdict == VerdictFail:
		c.Verdict = VerdictFail
		c.Reasons = append(c.Reasons, "the run fails its own bars")
	case worse > 0:
		c.Verdict = VerdictFail
		c.Reasons = append(c.Reasons, fmt.Sprintf("%d gate rows moved against their sense past the margin of %g", worse, margin))
	default:
		c.Verdict = VerdictPass
		if base == nil {
			c.Reasons = append(c.Reasons, "the first run of the suite, so nothing to compare")
		} else {
			c.Reasons = append(c.Reasons, "no gate row moved against its sense past the margin")
		}
	}
	if worse > 0 && next.Header.Verdict != VerdictFail {
		c.Reasons = append(c.Reasons, "the run passes its own bars, and the flips name what got worse")
	}
	return c
}

// flips pairs the rows of two runs by item and metric.
func flips(base, next *Run, lower map[string]bool, margin float64) (out []Flip, gone, added []string) {
	type key struct{ item, metric string }
	baseRows := map[key]Row{}
	baseItems := map[string]bool{}
	for _, r := range base.Rows {
		baseRows[key{r.Item, r.Metric}] = r
		baseItems[r.Item] = true
	}
	nextItems := map[string]bool{}
	for _, r := range next.Rows {
		nextItems[r.Item] = true
		b, ok := baseRows[key{r.Item, r.Metric}]
		if !ok {
			continue
		}
		if b.Value == r.Value {
			continue
		}
		delta := r.Value - b.Value
		if lower[r.Metric] {
			delta = -delta
		}
		// delta is now positive when the row got better.
		worse := delta < 0 && math.Abs(delta) > margin
		out = append(out, Flip{
			Item: r.Item, Metric: r.Metric, Kind: r.Kind, Before: b.Value, After: r.Value,
			Worse: worse, BeforeDetail: b.Detail, AfterDetail: r.Detail,
		})
	}
	for item := range baseItems {
		if !nextItems[item] {
			gone = append(gone, item)
		}
	}
	for item := range nextItems {
		if !baseItems[item] {
			added = append(added, item)
		}
	}
	sortItems(gone)
	sortItems(added)
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Kind != b.Kind {
			return a.Kind == KindGate
		}
		if a.Worse != b.Worse {
			return a.Worse
		}
		if a.Item != b.Item {
			return itemLess(a.Item, b.Item)
		}
		return a.Metric < b.Metric
	})
	return out, gone, added
}

// itemLess orders items by number when both are numbers, else by text,
// so "2" sorts before "10" and "suite" sorts last.
func itemLess(a, b string) bool {
	var na, nb int
	_, errA := fmt.Sscanf(a, "%d", &na)
	_, errB := fmt.Sscanf(b, "%d", &nb)
	switch {
	case errA == nil && errB == nil:
		if na != nb {
			return na < nb
		}
		return a < b
	case errA == nil:
		return true
	case errB == nil:
		return false
	}
	return a < b
}

func sortItems(items []string) {
	sort.Slice(items, func(i, j int) bool { return itemLess(items[i], items[j]) })
}

// epochNotes names what differs between two fingerprints.
func epochNotes(base, next Header) []string {
	var notes []string
	if base.Suite != next.Suite {
		notes = append(notes, fmt.Sprintf("the suite differs: %s against %s", base.Suite, next.Suite))
	}
	for _, role := range sortedKeys(union(base.Roles, next.Roles)) {
		b, n := base.Roles[role], next.Roles[role]
		if b != n {
			notes = append(notes, fmt.Sprintf("the %s role moved from %s to %s", role, modelWord(b), modelWord(n)))
		}
	}
	for _, name := range sortedKeys(union(base.Prompts, next.Prompts)) {
		if base.Prompts[name] != next.Prompts[name] {
			notes = append(notes, fmt.Sprintf("the %s prompt moved from version %d to %d", name, base.Prompts[name], next.Prompts[name]))
		}
	}
	if base.Snapshot != next.Snapshot {
		notes = append(notes, fmt.Sprintf("the card snapshot moved from %s to %s", orWord(base.Snapshot, "none"), orWord(next.Snapshot, "none")))
	}
	for _, name := range sortedKeys(union(base.Versions, next.Versions)) {
		if base.Versions[name] != next.Versions[name] {
			notes = append(notes, fmt.Sprintf("%s moved from %s to %s", name, orWord(base.Versions[name], "none"), orWord(next.Versions[name], "none")))
		}
	}
	return notes
}

func modelWord(m Model) string {
	if m.Model == "" {
		return "none"
	}
	if m.Effort != "" {
		return m.Model + " (" + m.Provider + ", effort " + m.Effort + ")"
	}
	return m.Model + " (" + m.Provider + ")"
}

func union[V any](a, b map[string]V) map[string]bool {
	out := map[string]bool{}
	for k := range a {
		out[k] = true
	}
	for k := range b {
		out[k] = true
	}
	return out
}

// Merge folds runs into one: the header of the last run, and the rows
// of every run with a later run's items in place of an earlier run's.
// It reads a run and its rerun as one, the way a gate document and its
// -only rerun are read together.
func Merge(runs ...*Run) *Run {
	if len(runs) == 0 {
		return nil
	}
	out := &Run{Header: runs[len(runs)-1].Header}
	out.Header.Lower = nil
	byItem := map[string][]Row{}
	var order []string
	for _, r := range runs {
		out.LowerIsBetter(r.Header.Lower...)
		seen := map[string]bool{}
		for _, row := range r.Rows {
			if !seen[row.Item] {
				seen[row.Item] = true
				if _, ok := byItem[row.Item]; !ok {
					order = append(order, row.Item)
				}
				byItem[row.Item] = nil
			}
			byItem[row.Item] = append(byItem[row.Item], row)
		}
	}
	for _, item := range order {
		out.Rows = append(out.Rows, byItem[item]...)
	}
	if len(runs) > 1 {
		var ids []string
		for _, r := range runs {
			ids = append(ids, r.Header.RunID)
		}
		out.Header.RunID = strings.Join(ids, "+")
	}
	return out
}
