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
	// Conversations holds the M-4 count of every conversation, so a
	// partial run can be folded into this one (D-181).
	Conversations []ConvCount `json:"conversations,omitempty"`
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

// DefaultNoise is how many bad questions two runs of identical code may
// differ by. Run 18 and run 20260826-191225-000 ran the same agent code
// on the same 104 conversations: the holdout moved by three bad
// questions and the whole set by three, and half of the bad set was a
// different half. A ratio guard below that level judges the judge, and
// not the fixer (D-183).
const DefaultNoise = 3

// DefaultDriftNoise is how many questions may get worse on rows no change
// declared before the loop calls it damage. It is a different quantity
// from DefaultNoise, and it has a far larger floor.
//
// Measured on 2026-08-26. Run 18 and run 20260826-191225-000 ran the same
// agent code on the same 104 conversations. On rows no change declared,
// 15 questions got worse and 8 got better, so the harm minus the help was
// 7. The ask role rewrites its wording every run, so a question moves
// with no edit behind it, and 12 of those 15 moved on changed text alone.
//
// The old guard compared the raw harm against DefaultNoise, which is 3.
// It therefore rejected identical code, and it rejected every iteration
// the loop ever ran (D-217).
//
// CAUTION: one same-code pair sets this number. Measure a second pair
// before you lower it, and keep TestIdenticalCodeIsAccepted green.
const DefaultDriftNoise = 8

// Decision is the accept or reject of one iteration.
type Decision struct {
	Accept  bool     `json:"accept"`
	Reasons []string `json:"reasons"`
	// Partial says some changes are kept and some are dropped. Keep and
	// Drop hold the commits, and Remeasure the conversations the kept
	// rows touch, which the loop runs again before the merge (D-181).
	Partial   bool     `json:"partial,omitempty"`
	Keep      []string `json:"keep,omitempty"`
	Drop      []string `json:"drop,omitempty"`
	Remeasure []int    `json:"remeasure,omitempty"`
}

// Compare decides whether a new run may be kept. prev may be nil, which
// is the first iteration: then only the hard rules apply.
//
// The hard rules stand on their own. The linter must find nothing, and no
// session may call itself complete with a needed slot unanswered. The
// soft rules read the run before: the agent must not go quiet, and it
// must not close fewer slots.
//
// The ratio guards carry a margin of DefaultNoise bad questions. Use
// CompareWith to set another margin.
func Compare(prev, next *Summary) Decision { return CompareWith(prev, next, DefaultNoise) }

// CompareWith is Compare with a named noise margin, in bad questions.
func CompareWith(prev, next *Summary, noise int) Decision {
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
	// The guards count bad questions, not points. A margin of noise bad
	// questions is what two runs of the same code differ by, so a rise
	// inside it is not evidence against the change (D-183).
	scoredBad := func(s *Summary) int {
		if s.HoldoutJudged > 0 {
			return s.HoldoutBad
		}
		return s.Bad
	}
	if got, was := scoredBad(next), scoredBad(prev); got > was+noise {
		fail("the bad-question ratio on the %s rose from %.1f%% to %.1f%%, which is %d more bad questions and the margin is %d",
			split, prevRatio*100, nextRatio*100, got-was, noise)
	}
	// The holdout decides, and the whole set is the sanity check under it.
	// The holdout holds about 120 questions, so two of them move it by
	// more than a point. Iteration 1 of 2026-08-26 improved the holdout by
	// two bad questions while every judged question got worse by three,
	// and the loop called that progress. A ratio that falls on the holdout
	// while it rises over the whole set is noise, and not a gain (D-178).
	if next.HoldoutJudged > 0 && next.Bad > prev.Bad+noise {
		shape := "and the holdout rose too"
		if nextRatio < prevRatio {
			shape = "although the holdout fell"
		} else if nextRatio == prevRatio {
			shape = "and the holdout stood still"
		}
		fail("the bad-question ratio over every judged question rose from %.1f%% to %.1f%%, %s",
			prev.Ratio*100, next.Ratio*100, shape)
	}
	// A fall on the tune split alone is the shape of an overfit. It does
	// not reject the iteration, and the loop says so out loud.
	if next.HoldoutJudged > 0 && next.TuneRatio < prev.TuneRatio && nextRatio >= prevRatio {
		d.Reasons = append(d.Reasons,
			fmt.Sprintf("WARNING: the tune split improved and the holdout did not (%.1f%% to %.1f%%)",
				prevRatio*100, nextRatio*100))
	}
	if d.Accept {
		d.Reasons = append(d.Reasons, fmt.Sprintf("the %s ratio moved from %.1f%% to %.1f%%, inside the margin of %d, and no counter fell",
			split, prevRatio*100, nextRatio*100, noise))
	}
	return d
}

// Decide is Compare with the fixer's changes in hand. When the changes
// declare their rows, every question that moved is charged to the change
// that owns its row, and the loop keeps the changes that helped and drops
// the ones that hurt (D-181). The whole-run rules still stand over it:
// a hard failure, a quiet agent, or a rise past the noise margin on rows
// no change declared rejects everything.
func Decide(prev, next *Summary, changes []Change, noise int) (Decision, Paired, Attribution) {
	d := CompareWith(prev, next, noise)
	if prev == nil || len(changes) == 0 {
		return d, Paired{}, Attribution{}
	}
	p := Pair(prev, next)
	a := Attribute(p, changes)
	// Unattributed movement means a question moved on a row no change
	// declared. Most of it is the ask role, which rewrites its wording
	// every run, so the raw harm is not damage on its own (D-217). The
	// guard reads the harm against the help, because pure churn moves
	// both and real damage moves one.
	unBad, unGood := 0, 0
	for _, u := range a.Unattributed {
		switch u.Kind {
		case "worse", "new_bad":
			unBad++
		case "better", "gone_bad":
			unGood++
		}
	}
	if drift := unBad - unGood; drift > DefaultDriftNoise {
		d.Accept = false
		d.Reasons = append(d.Reasons, fmt.Sprintf("on rows no change declared, %d questions got worse and %d got better, so the drift of %d passed the margin of %d",
			unBad, unGood, drift, DefaultDriftNoise))
		return d, p, a
	}
	var keepRows []string
	for _, cv := range a.Changes {
		if cv.Keep {
			d.Keep = append(d.Keep, cv.Change.Commit)
			keepRows = append(keepRows, cv.Change.Rows...)
		} else {
			d.Drop = append(d.Drop, cv.Change.Commit)
		}
	}
	hard := false
	for _, r := range d.Reasons {
		if strings.Contains(r, "linter") || strings.Contains(r, "complete with a slot") ||
			strings.Contains(r, "floor") {
			hard = true
		}
	}
	switch {
	case hard:
		// A hard failure is the whole run's, and no change is kept.
		d.Keep, d.Drop = nil, d.Keep
		d.Accept = false
	case len(d.Drop) == 0:
		// Every change helped or stood still: the whole-run decision holds.
	case len(d.Keep) == 0:
		d.Accept = false
		d.Reasons = append(d.Reasons, "every change made its own rows worse")
	default:
		d.Accept = false
		d.Partial = true
		d.Remeasure = Conversations(keepRows, prev, next)
		d.Reasons = append(d.Reasons, fmt.Sprintf("%d changes are kept and %d are dropped, so the kept rows are measured again on %d conversations",
			len(d.Keep), len(d.Drop), len(d.Remeasure)))
	}
	return d, p, a
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
