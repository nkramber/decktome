package tune

import (
	"fmt"
	"sort"
	"strings"
)

// Verdict is the eval role's score for one question. The fields mirror
// the owner's rubric, so an automated score and a hand score can sit in
// one table (D-66, D-104, D-114).
type Verdict struct {
	Conversation string `json:"conversation"`
	Turn         int    `json:"turn"`
	Row          string `json:"row"`
	Slot         string `json:"slot"`
	Text         string `json:"text"`
	// Warranted says whether the question deserved to be asked, to this
	// user, at this point. Values: yes, no, unsure.
	Warranted string `json:"warranted"`
	// Faults come from the D-114 vocabulary.
	Faults []string `json:"faults,omitempty"`
	// CatalogAction is none, reword, add, or remove (D-25).
	CatalogAction string `json:"catalog_action"`
	// Reason is one sentence. It is what a person reads first.
	Reason string `json:"reason"`
	// Holdout marks a verdict the fixer never reads. The loop measures
	// its own overfit with it: a ratio that falls on the tune split and
	// stands still on the holdout says the fixer learned the test set and
	// not the product (D-134).
	Holdout bool `json:"holdout,omitempty"`
}

// Bad reports whether this verdict counts against the run.
func (v Verdict) Bad() bool { return strings.EqualFold(v.Warranted, "no") }

// Unsure reports whether this verdict is left out of the ratio.
func (v Verdict) Unsure() bool {
	return !v.Bad() && !strings.EqualFold(v.Warranted, "yes")
}

// Summary is one judged run.
type Summary struct {
	Run     string  `json:"run"`
	Model   string  `json:"model"`
	CostUSD float64 `json:"cost_usd"`
	Metrics Metrics `json:"metrics"`
	Judged  int     `json:"judged"`
	Bad     int     `json:"bad"`
	Unsure  int     `json:"unsure"`
	Ratio   float64 `json:"bad_ratio"`
	// TuneRatio reads the conversations the fixer sees. HoldoutRatio reads
	// the ones it never sees (D-134).
	TuneJudged    int     `json:"tune_judged"`
	TuneBad       int     `json:"tune_bad"`
	TuneRatio     float64 `json:"tune_ratio"`
	HoldoutJudged int     `json:"holdout_judged"`
	HoldoutBad    int     `json:"holdout_bad"`
	HoldoutRatio  float64 `json:"holdout_ratio"`

	ByRow    map[string]int `json:"bad_by_row"`
	ByFault  map[string]int `json:"bad_by_fault"`
	Actions  map[string]int `json:"catalog_actions"`
	Verdicts []Verdict      `json:"verdicts"`
}

// Summarize counts the verdicts of one run.
func Summarize(run string, model string, cost float64, m Metrics, vs []Verdict) Summary {
	s := Summary{
		Run: run, Model: model, CostUSD: cost, Metrics: m, Verdicts: vs,
		ByRow: map[string]int{}, ByFault: map[string]int{}, Actions: map[string]int{},
	}
	for _, v := range vs {
		switch {
		case v.Unsure():
			s.Unsure++
			continue
		case v.Bad():
			s.Bad++
			s.ByRow[v.Row]++
			for _, f := range v.Faults {
				s.ByFault[strings.ToLower(strings.TrimSpace(f))]++
			}
		}
		s.Judged++
		if v.Holdout {
			s.HoldoutJudged++
			if v.Bad() {
				s.HoldoutBad++
			}
		} else {
			s.TuneJudged++
			if v.Bad() {
				s.TuneBad++
			}
		}
		if a := strings.ToLower(strings.TrimSpace(v.CatalogAction)); a != "" && a != "none" {
			s.Actions[a]++
		}
	}
	s.Ratio = ratio(s.Bad, s.Judged)
	s.TuneRatio = ratio(s.TuneBad, s.TuneJudged)
	s.HoldoutRatio = ratio(s.HoldoutBad, s.HoldoutJudged)
	return s
}

func ratio(bad, judged int) float64 {
	if judged == 0 {
		return 0
	}
	return float64(bad) / float64(judged)
}

// Scored is the ratio the loop reads. It prefers the holdout split, which
// the fixer never sees. A fixer that only rewords the questions it was
// shown moves the tune ratio and leaves this one alone (D-134).
func (s Summary) Scored() (float64, string) {
	if s.HoldoutJudged > 0 {
		return s.HoldoutRatio, "holdout"
	}
	return s.Ratio, "every conversation"
}

// asked is the question count the loop reads. It prefers the count taken
// from the transcript, which holds every conversation.
func (m Metrics) asked() int {
	if m.QuestionsSeen > 0 {
		return m.QuestionsSeen
	}
	return m.Questions
}

// filled is the count of questions whose key closed, read the same way.
func (m Metrics) filled() int {
	if m.FilledSeen > 0 {
		return m.FilledSeen
	}
	return m.CatalogFilled
}

// Tolerance is how far a counter may fall before the loop calls it a
// regression. A run that asks fewer questions can score better while it
// serves the user worse, which gate run 7 of 2026-08-25 proved.
const Tolerance = 0.90

// Decision is the accept or reject of one iteration.
type Decision struct {
	Accept  bool     `json:"accept"`
	Reasons []string `json:"reasons"`
}

// Compare decides whether a new run may be kept. prev may be nil, which
// is the first iteration: then only the hard rules apply.
//
// The hard rules stand on their own. The linter must find nothing, and no
// session may call itself complete with a needed slot unanswered. The
// soft rules read the run before: the agent must not go quiet, and it
// must not close fewer slots.
func Compare(prev, next *Summary) Decision {
	d := Decision{Accept: true}
	fail := func(format string, a ...any) {
		d.Accept = false
		d.Reasons = append(d.Reasons, fmt.Sprintf(format, a...))
	}
	if next == nil {
		return Decision{Accept: false, Reasons: []string{"no run to judge"}}
	}
	if next.Metrics.LintFindings > 0 {
		fail("the linter found %d defective questions", next.Metrics.LintFindings)
	}
	if next.Metrics.Premature > 0 {
		fail("%d sessions called themselves complete with a slot unanswered", next.Metrics.Premature)
	}
	if prev == nil {
		if d.Accept {
			d.Reasons = append(d.Reasons, "first run, so nothing to compare")
		}
		return d
	}
	if got, was := next.Metrics.asked(), prev.Metrics.asked(); got < int(float64(was)*Tolerance) {
		fail("the agent asked %d questions, down from %d: below the %.0f%% floor",
			got, was, Tolerance*100)
	}
	if got, was := next.Metrics.filled(), prev.Metrics.filled(); got < int(float64(was)*Tolerance) {
		fail("%d questions closed a slot, down from %d: below the %.0f%% floor",
			got, was, Tolerance*100)
	}
	nextRatio, split := next.Scored()
	prevRatio, _ := prev.Scored()
	if nextRatio > prevRatio {
		fail("the bad-question ratio on the %s rose from %.1f%% to %.1f%%",
			split, prevRatio*100, nextRatio*100)
	}
	// A fall on the tune split alone is the shape of an overfit. It does
	// not reject the iteration, and the loop says so out loud.
	if next.HoldoutJudged > 0 && next.TuneRatio < prev.TuneRatio && nextRatio >= prevRatio {
		d.Reasons = append(d.Reasons,
			fmt.Sprintf("WARNING: the tune split improved and the holdout did not (%.1f%% to %.1f%%)",
				prevRatio*100, nextRatio*100))
	}
	if d.Accept {
		d.Reasons = append(d.Reasons, fmt.Sprintf("the %s ratio moved from %.1f%% to %.1f%%, and no counter fell",
			split, prevRatio*100, nextRatio*100))
	}
	return d
}

// TopRows names the rows with the most bad questions, worst first.
func (s Summary) TopRows(n int) []string {
	rows := make([]string, 0, len(s.ByRow))
	for r := range s.ByRow {
		rows = append(rows, r)
	}
	sort.Slice(rows, func(i, j int) bool {
		if s.ByRow[rows[i]] != s.ByRow[rows[j]] {
			return s.ByRow[rows[i]] > s.ByRow[rows[j]]
		}
		return rows[i] < rows[j]
	})
	if n > 0 && len(rows) > n {
		rows = rows[:n]
	}
	return rows
}
