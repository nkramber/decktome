package tune

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// A run and the run before it share their conversations, so the loop can
// compare them question by question instead of ratio by ratio. Two runs
// of the same code differ by three holdout questions, and half of the
// bad set is swapped for another half (run 18 against 20260826-191225-000).
// A ratio can not see that. A paired comparison can: it counts only the
// questions whose verdict changed AND whose text changed, because a
// verdict that flips on identical text is the judge and not the fixer.

// numberRe reads the number a conversation name starts with.
var numberRe = regexp.MustCompile(`^\s*(\d+)\.`)

// ConversationNumber reads the number of a conversation from its name,
// which the gate writes as "12. the user says none, then picks". A name
// with no number gives 0.
func ConversationNumber(name string) int {
	m := numberRe.FindStringSubmatch(name)
	if m == nil {
		return 0
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

// HeldOut reports whether a conversation belongs to the holdout split.
// The split keys on the conversation number, not on the position in the
// document, so a partial run holds out the same conversations as a full
// one (D-134).
func HeldOut(name string, every int) bool {
	if every <= 0 {
		return false
	}
	n := ConversationNumber(name)
	return n > 0 && n%every == 0
}

// PairKey names one question across runs: the conversation, the turn,
// and the row.
func PairKey(v Verdict) string {
	return fmt.Sprintf("%d|%d|%s", ConversationNumber(v.Conversation), v.Turn, v.Row)
}

// Delta is one question that changed between two runs.
type Delta struct {
	Row          string  `json:"row"`
	Conversation string  `json:"conversation"`
	Turn         int     `json:"turn"`
	Holdout      bool    `json:"holdout"`
	Before       Verdict `json:"before"`
	After        Verdict `json:"after"`
	// Kind is worse, better, new_bad, gone_bad, or judge_flip.
	Kind string `json:"kind"`
}

// Paired is the question-by-question comparison of two runs.
type Paired struct {
	// Shared counts the questions both runs asked.
	Shared int `json:"shared"`
	// Worse are shared questions that went from good to bad with a
	// changed text, and NewBad are bad questions the new run asked that
	// the old run did not. Both count against the change.
	Worse  []Delta `json:"worse"`
	NewBad []Delta `json:"new_bad"`
	// Better are shared questions that went from bad to good with a
	// changed text, and GoneBad are bad questions the old run asked that
	// the new run does not. Both count for the change.
	Better  []Delta `json:"better"`
	GoneBad []Delta `json:"gone_bad"`
	// JudgeFlips are verdicts that changed on identical text. They are
	// the judge's own noise, and no change is charged for them.
	JudgeFlips []Delta `json:"judge_flips"`
}

// Pair compares two runs question by question.
func Pair(prev, next *Summary) Paired {
	var p Paired
	before := map[string]Verdict{}
	for _, v := range prev.Verdicts {
		before[PairKey(v)] = v
	}
	after := map[string]Verdict{}
	for _, v := range next.Verdicts {
		after[PairKey(v)] = v
	}
	for _, v := range next.Verdicts {
		k := PairKey(v)
		old, ok := before[k]
		if !ok {
			if v.Bad() {
				p.NewBad = append(p.NewBad, delta(old, v, "new_bad"))
			}
			continue
		}
		p.Shared++
		if old.Bad() == v.Bad() {
			continue
		}
		if sameText(old.Text, v.Text) {
			p.JudgeFlips = append(p.JudgeFlips, delta(old, v, "judge_flip"))
			continue
		}
		if v.Bad() {
			p.Worse = append(p.Worse, delta(old, v, "worse"))
		} else {
			p.Better = append(p.Better, delta(old, v, "better"))
		}
	}
	for _, v := range prev.Verdicts {
		if _, ok := after[PairKey(v)]; !ok && v.Bad() {
			p.GoneBad = append(p.GoneBad, delta(v, Verdict{}, "gone_bad"))
		}
	}
	for _, list := range []*[]Delta{&p.Worse, &p.NewBad, &p.Better, &p.GoneBad, &p.JudgeFlips} {
		sortDeltas(*list)
	}
	return p
}

func delta(before, after Verdict, kind string) Delta {
	d := Delta{Before: before, After: after, Kind: kind}
	src := after
	if kind == "gone_bad" {
		src = before
	}
	d.Row, d.Conversation, d.Turn, d.Holdout = src.Row, src.Conversation, src.Turn, src.Holdout
	return d
}

func sortDeltas(ds []Delta) {
	sort.Slice(ds, func(i, j int) bool {
		if ds[i].Row != ds[j].Row {
			return ds[i].Row < ds[j].Row
		}
		if a, b := ConversationNumber(ds[i].Conversation), ConversationNumber(ds[j].Conversation); a != b {
			return a < b
		}
		return ds[i].Turn < ds[j].Turn
	})
}

func sameText(a, b string) bool {
	return strings.EqualFold(strings.TrimSpace(a), strings.TrimSpace(b))
}

// Against counts every delta that counts against a change.
func (p Paired) Against() int { return len(p.Worse) + len(p.NewBad) }

// For counts every delta that counts for a change.
func (p Paired) For() int { return len(p.Better) + len(p.GoneBad) }

// Change is one independent edit the fixer made, as one commit with two
// trailers: the rows it means to move, and the hypothesis behind it.
type Change struct {
	Commit     string   `json:"commit"`
	Subject    string   `json:"subject"`
	Rows       []string `json:"rows"`
	Hypothesis string   `json:"hypothesis"`
}

// ChangeVerdict is the loop's answer for one change.
type ChangeVerdict struct {
	Change  Change  `json:"change"`
	For     []Delta `json:"for"`
	Against []Delta `json:"against"`
	Keep    bool    `json:"keep"`
	Reason  string  `json:"reason"`
}

// Attribution is the paired comparison, charged to the changes that
// declared the rows.
type Attribution struct {
	Changes []ChangeVerdict `json:"changes"`
	// Unattributed are deltas on rows no change declared. A change that
	// moved a row it did not name is charged nothing, so a large count
	// here says the declarations were wrong, and the loop falls back to
	// the whole-run rules.
	Unattributed []Delta `json:"unattributed"`
}

// Attribute charges each delta to the change that declared its row. A
// change is kept when more questions got better than worse on its rows.
// A change with no delta on its rows is kept: nothing spoke against it.
// A tie with activity is dropped: the change moved questions and did not
// improve them.
func Attribute(p Paired, changes []Change) Attribution {
	owner := map[string]int{}
	for i, c := range changes {
		for _, r := range c.Rows {
			owner[strings.TrimSpace(r)] = i
		}
	}
	out := Attribution{Changes: make([]ChangeVerdict, len(changes))}
	for i, c := range changes {
		out.Changes[i].Change = c
	}
	charge := func(ds []Delta, against bool) {
		for _, d := range ds {
			i, ok := owner[d.Row]
			if !ok {
				out.Unattributed = append(out.Unattributed, d)
				continue
			}
			if against {
				out.Changes[i].Against = append(out.Changes[i].Against, d)
			} else {
				out.Changes[i].For = append(out.Changes[i].For, d)
			}
		}
	}
	charge(p.Worse, true)
	charge(p.NewBad, true)
	charge(p.Better, false)
	charge(p.GoneBad, false)
	for i := range out.Changes {
		cv := &out.Changes[i]
		f, a := len(cv.For), len(cv.Against)
		switch {
		case f == 0 && a == 0:
			cv.Keep = true
			cv.Reason = "no judged question moved on its rows"
		case f > a:
			cv.Keep = true
			cv.Reason = fmt.Sprintf("%d questions got better and %d got worse on its rows", f, a)
		default:
			cv.Keep = false
			cv.Reason = fmt.Sprintf("%d questions got worse and %d got better on its rows", a, f)
		}
	}
	sortDeltas(out.Unattributed)
	return out
}

// Conversations lists the conversation numbers a set of rows touched in
// either run. The loop re-measures those alone after a partial keep, so
// the kept subset is scored on the conversations it can move (D-181).
func Conversations(rows []string, runs ...*Summary) []int {
	want := map[string]bool{}
	for _, r := range rows {
		want[strings.TrimSpace(r)] = true
	}
	seen := map[int]bool{}
	for _, s := range runs {
		if s == nil {
			continue
		}
		for _, v := range s.Verdicts {
			if want[v.Row] {
				if n := ConversationNumber(v.Conversation); n > 0 {
					seen[n] = true
				}
			}
		}
	}
	out := make([]int, 0, len(seen))
	for n := range seen {
		out = append(out, n)
	}
	sort.Ints(out)
	return out
}

// ConvCount is the M-4 count of one conversation. A summary carries one
// per conversation, so a partial run can be folded into it without the
// gate document (D-181).
type ConvCount struct {
	Name          string `json:"name"`
	Probe         bool   `json:"probe,omitempty"`
	Questions     int    `json:"questions"`
	Invented      int    `json:"invented,omitempty"`
	CatalogFilled int    `json:"catalog_filled"`
	Filled        int    `json:"filled"`
	Premature     bool   `json:"premature,omitempty"`
}

// CountConversations reads the per-conversation counts of a gate document.
func CountConversations(run *Run) []ConvCount {
	out := make([]ConvCount, 0, len(run.Conversations))
	for _, c := range run.Conversations {
		cc := ConvCount{Name: c.Name, Probe: c.Probe, Premature: c.Premature}
		for _, q := range c.Questions {
			cc.Questions++
			if q.Filled {
				cc.Filled++
			}
			if q.Invented {
				cc.Invented++
				continue
			}
			if q.Filled {
				cc.CatalogFilled++
			}
		}
		out = append(out, cc)
	}
	return out
}

// MetricsFromCounts recomputes the M-4 counters. The table in the
// document counts the gate conversations alone, and the two Seen counters
// read every conversation, probes included.
// TestMetricsOfMatchesTheDocument proves both readings agree with what
// the gate wrote.
func MetricsFromCounts(counts []ConvCount) Metrics {
	var m Metrics
	for _, c := range counts {
		m.QuestionsSeen += c.Questions
		m.FilledSeen += c.Filled
		// A premature session fails the gate wherever it sits, probes
		// included. Conversation 35 of run 11 is a probe, and it counts.
		if c.Premature {
			m.Premature++
		}
		if c.Probe {
			continue
		}
		m.Questions += c.Questions
		m.Invented += c.Invented
		m.Catalog += c.Questions - c.Invented
		m.CatalogFilled += c.CatalogFilled
		if c.Invented == 0 {
			m.CatalogOnly++
		}
	}
	return m
}

// MetricsOf recomputes the M-4 counters from a gate document.
func MetricsOf(run *Run) Metrics { return MetricsFromCounts(CountConversations(run)) }

// Merge folds a partial run into a full one. Every conversation the
// partial run measured replaces the same conversation of the base, and
// the rest of the base carries over. The counters come from the
// per-conversation counts both summaries carry, so they are counts and
// never estimates. A summary from before D-181 carries none, and it can
// not be a base.
func Merge(base, part *Summary) (Summary, error) {
	if len(base.Conversations) == 0 {
		return Summary{}, fmt.Errorf("tune: %s carries no per-conversation counts, so nothing can be folded into it", base.Run)
	}
	if len(part.Conversations) == 0 {
		return Summary{}, fmt.Errorf("tune: %s carries no per-conversation counts", part.Run)
	}
	replaced := map[int]bool{}
	for _, c := range part.Conversations {
		n := ConversationNumber(c.Name)
		if n == 0 {
			return Summary{}, fmt.Errorf("tune: conversation %q has no number, so it can not replace one", c.Name)
		}
		replaced[n] = true
	}
	var vs []Verdict
	for _, v := range base.Verdicts {
		if !replaced[ConversationNumber(v.Conversation)] {
			vs = append(vs, v)
		}
	}
	vs = append(vs, part.Verdicts...)
	var counts []ConvCount
	for _, c := range base.Conversations {
		if !replaced[ConversationNumber(c.Name)] {
			counts = append(counts, c)
		}
	}
	counts = append(counts, part.Conversations...)
	sort.Slice(counts, func(i, j int) bool {
		return ConversationNumber(counts[i].Name) < ConversationNumber(counts[j].Name)
	})
	m := MetricsFromCounts(counts)
	// The linter count is per document, and a partial document counts
	// its own conversations alone. The sum is the honest upper bound.
	m.LintFindings = base.Metrics.LintFindings + part.Metrics.LintFindings
	s := Summarize(part.Run, part.Model, base.CostUSD+part.CostUSD, m, vs)
	s.Conversations = counts
	return s, nil
}
