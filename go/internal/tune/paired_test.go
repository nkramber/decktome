package tune

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func v(conv string, turn int, row, text, warranted string) Verdict {
	return Verdict{Conversation: conv, Turn: turn, Row: row, Text: text, Warranted: warranted}
}

func sum(vs ...Verdict) *Summary {
	s := Summarize("r", "m", 0, Metrics{Questions: 10, CatalogFilled: 8}, vs)
	return &s
}

// TestPairSeparatesTheJudgeFromTheFixer is the reason the comparison is
// paired. A verdict that flips on identical text is the judge's noise,
// and a verdict that flips on a changed text is the fixer's work.
func TestPairSeparatesTheJudgeFromTheFixer(t *testing.T) {
	prev := sum(
		v("1. a", 1, "colors", "Any color preference?", "yes"),
		v("2. b", 1, "meta", "What decks do you expect to face?", "no"),
		v("3. c", 1, "power", "How strong?", "no"),
		v("4. d", 1, "locked", "Keep it?", "no"),
	)
	next := sum(
		v("1. a", 1, "colors", "Any color preference?", "no"),
		v("2. b", 1, "meta", "Which decks, or keep the sideboard general?", "yes"),
		v("3. c", 1, "power", "How strong: casual or tournament?", "no"),
		v("5. e", 1, "theme", "What theme?", "no"),
	)
	p := Pair(prev, next)
	if p.Shared != 3 {
		t.Errorf("shared = %d, want 3", p.Shared)
	}
	if len(p.JudgeFlips) != 1 || p.JudgeFlips[0].Row != "colors" {
		t.Errorf("judge flips = %+v, want the colors flip alone", p.JudgeFlips)
	}
	if len(p.Better) != 1 || p.Better[0].Row != "meta" {
		t.Errorf("better = %+v, want the meta row", p.Better)
	}
	if len(p.Worse) != 0 {
		t.Errorf("worse = %+v, want none: the power row stayed bad", p.Worse)
	}
	if len(p.NewBad) != 1 || p.NewBad[0].Row != "theme" {
		t.Errorf("new bad = %+v, want the theme row", p.NewBad)
	}
	if len(p.GoneBad) != 1 || p.GoneBad[0].Row != "locked" {
		t.Errorf("gone bad = %+v, want the locked row", p.GoneBad)
	}
	if p.For() != 2 || p.Against() != 1 {
		t.Errorf("for = %d against = %d, want 2 and 1", p.For(), p.Against())
	}
}

// TestAttributeKeepsTheHelpfulChange charges each delta to the change
// that declared its row. One change helped, one hurt, and the loop keeps
// the first.
func TestAttributeKeepsTheHelpfulChange(t *testing.T) {
	prev := sum(
		v("1. a", 1, "meta", "old meta", "no"),
		v("2. b", 1, "meta", "old meta", "no"),
		v("3. c", 1, "power_sixty_confirm", "old confirm", "yes"),
		v("4. d", 1, "power_sixty_confirm", "old confirm", "yes"),
		v("5. e", 1, "colors", "colors?", "yes"),
	)
	next := sum(
		v("1. a", 1, "meta", "new meta", "yes"),
		v("2. b", 1, "meta", "new meta", "yes"),
		v("3. c", 1, "power_sixty_confirm", "new confirm", "no"),
		v("4. d", 1, "power_sixty_confirm", "new confirm", "no"),
		v("5. e", 1, "colors", "colors?", "no"),
	)
	changes := []Change{
		{Commit: "aaa", Subject: "meta offers general", Rows: []string{"meta"}},
		{Commit: "bbb", Subject: "confirm as written", Rows: []string{"power_sixty_confirm"}},
	}
	d, p, a, err := Decide(prev, next, changes, DefaultNoise)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.JudgeFlips) != 1 {
		t.Errorf("judge flips = %d, want 1: the colors row kept its text", len(p.JudgeFlips))
	}
	if !a.Changes[0].Keep || a.Changes[1].Keep {
		t.Errorf("keep = %v/%v, want the meta change kept and the confirm change dropped: %s / %s",
			a.Changes[0].Keep, a.Changes[1].Keep, a.Changes[0].Reason, a.Changes[1].Reason)
	}
	if !d.Partial || len(d.Keep) != 1 || d.Keep[0] != "aaa" || len(d.Drop) != 1 || d.Drop[0] != "bbb" {
		t.Errorf("decision = %+v, want a partial keep of aaa and drop of bbb", d)
	}
	if len(d.Remeasure) != 2 || d.Remeasure[0] != 1 || d.Remeasure[1] != 2 {
		t.Errorf("remeasure = %v, want the two meta conversations", d.Remeasure)
	}
}

// TestDecideWithNoDeclaredRowsFallsBackToTheWholeRun covers a fixer that
// made one commit with no trailers. Nothing can be attributed, so the
// whole-run rules decide.
func TestDecideWithNoDeclaredRowsFallsBackToTheWholeRun(t *testing.T) {
	prev := sum(v("1. a", 1, "meta", "old", "no"))
	next := sum(v("1. a", 1, "meta", "new", "yes"))
	d, _, _, err := Decide(prev, next, nil, DefaultNoise)
	if err != nil {
		t.Fatal(err)
	}
	if !d.Accept || d.Partial {
		t.Errorf("decision = %+v, want a plain accept", d)
	}
}

// TestAttributeRules is the charge table of T-4 and T-21. A commit with
// no rows is dropped, one row has one owner, a flip on a declared row is
// the change's whatever the text did, and a reworded flip on an
// undeclared row is churn.
func TestAttributeRules(t *testing.T) {
	prev := sum(
		v("1. a", 1, "meta", "old meta", "no"),
		v("2. b", 1, "meta", "same meta", "no"),
		v("3. c", 1, "colors", "old colors", "yes"),
		v("4. d", 1, "colors", "same colors", "yes"),
		v("5. e", 1, "power", "gone", "no"),
	)
	next := sum(
		v("1. a", 1, "meta", "new meta", "yes"),
		v("2. b", 1, "meta", "same meta", "yes"),
		v("3. c", 1, "colors", "new colors", "no"),
		v("4. d", 1, "colors", "same colors", "no"),
		v("6. f", 1, "budget", "new question", "no"),
	)
	cases := []struct {
		name    string
		changes []Change
		wantErr bool
		keep    []bool
		reasons []string
		churn   int
		unattr  int
	}{
		{
			name:    "a declared row takes both its flips",
			changes: []Change{{Commit: "aaa", Rows: []string{"meta"}}},
			keep:    []bool{true},
			reasons: []string{"2 questions got better and 0 got worse on its rows"},
			churn:   1, // the reworded colors flip
			unattr:  3, // the identical colors flip, the gone power, the new budget
		},
		{
			name:    "no rows declared is dropped",
			changes: []Change{{Commit: "bbb"}},
			keep:    []bool{false},
			reasons: []string{"no rows declared"},
			churn:   2,
			unattr:  4, // both identical flips, the gone power, the new budget
		},
		{
			name:    "one row has one owner",
			changes: []Change{{Commit: "aaa", Rows: []string{"meta"}}, {Commit: "bbb", Rows: []string{"meta"}}},
			wantErr: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := Attribute(Pair(prev, next), tc.changes)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, want error %v", err, tc.wantErr)
			}
			if err != nil {
				return
			}
			for i, cv := range a.Changes {
				if cv.Keep != tc.keep[i] || cv.Reason != tc.reasons[i] {
					t.Errorf("change %d: keep %v (%s), want %v (%s)", i, cv.Keep, cv.Reason, tc.keep[i], tc.reasons[i])
				}
			}
			if len(a.Churn) != tc.churn || len(a.Unattributed) != tc.unattr {
				t.Errorf("churn %d unattributed %d, want %d and %d", len(a.Churn), len(a.Unattributed), tc.churn, tc.unattr)
			}
		})
	}
}

// TestDecideStopsOnADoubleOwner is the tool fault the loop stops on.
func TestDecideStopsOnADoubleOwner(t *testing.T) {
	prev := sum(v("1. a", 1, "meta", "old", "no"))
	next := sum(v("1. a", 1, "meta", "new", "yes"))
	changes := []Change{{Commit: "aaa", Rows: []string{"meta"}}, {Commit: "bbb", Rows: []string{"meta"}}}
	if _, _, _, err := Decide(prev, next, changes, DefaultNoise); err == nil {
		t.Error("two owners of one row were accepted")
	}
}

// TestHardFailureDropsEveryChange covers T-15 and T-16. A hard failure
// keeps nothing, and the drop list holds every commit, the ones that
// helped included.
func TestHardFailureDropsEveryChange(t *testing.T) {
	prev := sum(v("1. a", 1, "meta", "old", "no"), v("2. b", 1, "power", "old", "yes"))
	next := sum(v("1. a", 1, "meta", "new", "yes"), v("2. b", 1, "power", "new", "no"))
	next.Metrics.LintFindings = 1
	changes := []Change{{Commit: "aaa", Rows: []string{"meta"}}, {Commit: "bbb", Rows: []string{"power"}}}
	d, _, _, err := Decide(prev, next, changes, DefaultNoise)
	if err != nil {
		t.Fatal(err)
	}
	if !d.Hard || d.Accept || len(d.Keep) != 0 {
		t.Errorf("decision = %+v, want a hard reject with nothing kept", d)
	}
	if len(d.Drop) != 2 {
		t.Errorf("drop = %v, want both commits", d.Drop)
	}
}

// TestErroredConversationCountsForNobody is T-5. A conversation that
// ended on an error asked nothing after it, and its missing questions are
// neither gone nor new.
func TestErroredConversationCountsForNobody(t *testing.T) {
	prev := sum(v("1. a", 1, "meta", "old", "no"), v("2. b", 1, "meta", "old", "no"))
	next := sum(v("1. a", 1, "meta", "new", "yes"))
	next.Errored = []string{"2. b"}
	p := Pair(prev, next)
	if p.Errored != 1 || len(p.GoneBad) != 0 || p.For() != 1 {
		t.Errorf("paired = %+v, want one errored conversation and one better question", p)
	}
	// The other direction skips too.
	prev.Errored = []string{"1. a"}
	next.Errored = nil
	next.Verdicts = append(next.Verdicts, v("2. b", 1, "meta", "new", "no"))
	p = Pair(prev, next)
	if p.Errored != 1 || p.For() != 0 || p.Against() != 0 {
		t.Errorf("paired = %+v, want nothing charged", p)
	}
}

// TestNoiseMarginAbsorbsTheJudge is D-183. Two runs of identical code
// differed by three bad questions on the holdout. A rise of that size is
// not a regression.
func TestNoiseMarginAbsorbsTheJudge(t *testing.T) {
	mk := func(holdBad int) *Summary {
		var vs []Verdict
		for i := 0; i < 40; i++ {
			w := "yes"
			if i < holdBad {
				w = "no"
			}
			vs = append(vs, Verdict{Conversation: "3. h", Row: "colors", Warranted: w, Holdout: true})
		}
		for i := 0; i < 80; i++ {
			vs = append(vs, Verdict{Conversation: "1. t", Row: "colors", Warranted: "yes"})
		}
		s := Summarize("r", "m", 0, Metrics{Questions: 100, CatalogFilled: 80}, vs)
		return &s
	}
	prev := mk(10)
	if d := CompareWith(prev, mk(13), 3); !d.Accept {
		t.Errorf("a rise of three was rejected: %v", d.Reasons)
	}
	if d := CompareWith(prev, mk(14), 3); d.Accept {
		t.Error("a rise of four was accepted")
	}
	if d := CompareWith(prev, mk(13), 0); d.Accept {
		t.Error("with no margin a rise of three was accepted")
	}
	// The old message said "although the holdout fell" when it rose.
	prev2 := mk(10)
	next2 := mk(14)
	next2.Bad = prev2.Bad + 10
	d := CompareWith(prev2, next2, 3)
	for _, r := range d.Reasons {
		if strings.Contains(r, "every judged question") && strings.Contains(r, "although the holdout fell") {
			t.Errorf("the whole-set reason claims the holdout fell while it rose: %s", r)
		}
	}
}

// TestHeldOutKeysOnTheConversationNumber is what makes a partial run
// hold out the same conversations as a full one.
func TestHeldOutKeysOnTheConversationNumber(t *testing.T) {
	cases := map[string]bool{
		"3. lifegain": true, "6. x (probe)": true, "4. y": false, "no number": false, "": false,
	}
	for name, want := range cases {
		if got := HeldOut(name, 3); got != want {
			t.Errorf("HeldOut(%q) = %v, want %v", name, got, want)
		}
	}
	if HeldOut("3. x", 0) {
		t.Error("every 0 held a conversation out")
	}
}

// TestMetricsOfMatchesTheDocument proves the recomputed counters agree
// with the M-4 table the gate wrote, on every real gate document.
func TestMetricsOfMatchesTheDocument(t *testing.T) {
	matches, err := filepath.Glob("../../../docs/reference/pr7-question-gate-run1*.md")
	if err != nil || len(matches) == 0 {
		t.Skip("no gate document on disk")
	}
	for _, path := range matches {
		run, err := ReadRun(path)
		if err != nil {
			t.Fatalf("%s: %v", path, err)
		}
		got := MetricsOf(run)
		want := run.Metrics
		if got.Questions != want.Questions || got.Catalog != want.Catalog ||
			got.Invented != want.Invented || got.CatalogFilled != want.CatalogFilled ||
			got.CatalogOnly != want.CatalogOnly || got.QuestionsSeen != want.QuestionsSeen ||
			got.FilledSeen != want.FilledSeen || got.Premature != want.Premature {
			t.Errorf("%s: recomputed %+v, document %+v", filepath.Base(path), got, want)
		}
	}
}

// TestMergeReplacesTheMeasuredConversations covers the partial re-measure.
func TestMergeReplacesTheMeasuredConversations(t *testing.T) {
	baseRun := &Run{Conversations: []Conversation{
		{Name: "1. a", Questions: []Question{{Turn: 1, Row: "meta", Filled: true}}},
		{Name: "2. b", Questions: []Question{{Turn: 1, Row: "colors", Filled: true}, {Turn: 1, Row: "x", Invented: true}}},
		{Name: "3. c (probe)", Probe: true, Questions: []Question{{Turn: 1, Row: "meta"}}},
	}}
	base := sum(v("1. a", 1, "meta", "old", "no"), v("2. b", 1, "colors", "c", "yes"), v("3. c (probe)", 1, "meta", "old", "no"))
	base.Conversations = CountConversations(baseRun)
	partRun := &Run{Conversations: []Conversation{
		{Name: "1. a", Questions: []Question{{Turn: 1, Row: "meta", Filled: true}}},
		{Name: "3. c (probe)", Probe: true, Questions: []Question{{Turn: 1, Row: "meta", Filled: true}}},
	}}
	part := sum(v("1. a", 1, "meta", "new", "yes"), v("3. c (probe)", 1, "meta", "new", "yes"))
	part.Conversations = CountConversations(partRun)
	m, err := Merge(base, part)
	if err != nil {
		t.Fatal(err)
	}
	if m.Judged != 3 || m.Bad != 0 {
		t.Errorf("merged judged=%d bad=%d, want 3 and 0", m.Judged, m.Bad)
	}
	if m.Metrics.Questions != 3 || m.Metrics.Invented != 1 || m.Metrics.CatalogOnly != 1 ||
		m.Metrics.QuestionsSeen != 4 || m.Metrics.FilledSeen != 3 {
		t.Errorf("merged metrics = %+v", m.Metrics)
	}
	if len(m.Conversations) != 3 {
		t.Errorf("merged carries %d conversation counts, want 3", len(m.Conversations))
	}
	// A merged summary can be a base again.
	if _, err := Merge(&m, part); err != nil {
		t.Errorf("a merged summary refused a second merge: %v", err)
	}
	bad := sum(v("no number", 1, "meta", "x", "no"))
	bad.Conversations = []ConvCount{{Name: "no number"}}
	if _, err := Merge(base, bad); err == nil {
		t.Error("a partial conversation with no number was merged")
	}
	old := sum(v("1. a", 1, "meta", "x", "no"))
	if _, err := Merge(old, part); err == nil {
		t.Error("a base with no per-conversation counts was merged")
	}
}

// TestIdenticalCodeIsAccepted pins DefaultDriftNoise to the measurement
// of D-217. Run 18 and run 20260826-191225-000 ran the same agent code on
// the same 104 conversations. On rows no change declared, 15 questions
// got worse and 8 got better, because the ask role rewrites its wording
// every run.
//
// The old guard compared the raw 15 against DefaultNoise, which is 3, so
// it rejected identical code. This test fails if that comes back.
func TestIdenticalCodeIsAccepted(t *testing.T) {
	// One change declares the meta row, so every other row is undeclared
	// and its movement is drift.
	var prevV, nextV []Verdict
	prevV = append(prevV, v("1. a", 1, "meta", "old meta", "no"))
	nextV = append(nextV, v("1. a", 1, "meta", "new meta", "yes"))
	// Three verdicts flip on identical text. They are the judge alone, and
	// they hold the whole-set count at the +3 the two runs really moved.
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("%d. flip", i+70)
		prevV = append(prevV, v(name, 1, "budget", "same budget", "no"))
		nextV = append(nextV, v(name, 1, "budget", "same budget", "yes"))
	}
	// 15 undeclared questions get worse, and 8 get better. The ask role
	// rewrote each one, so every text differs.
	for i := 0; i < 15; i++ {
		name := fmt.Sprintf("%d. worse", i+10)
		prevV = append(prevV, v(name, 1, "colors", "old colors", "yes"))
		nextV = append(nextV, v(name, 1, "colors", "new colors", "no"))
	}
	for i := 0; i < 8; i++ {
		name := fmt.Sprintf("%d. better", i+40)
		prevV = append(prevV, v(name, 1, "pool", "old pool", "no"))
		nextV = append(nextV, v(name, 1, "pool", "new pool", "yes"))
	}
	changes := []Change{{Commit: "aaa", Subject: "meta offers general", Rows: []string{"meta"}}}
	d, _, a, err := Decide(sum(prevV...), sum(nextV...), changes, DefaultNoise)
	if err != nil {
		t.Fatal(err)
	}
	// The 23 reworded flips are churn, charged to nobody (T-21). The three
	// identical-text flips went from bad to good, so they are help.
	if len(a.Churn) != 23 {
		t.Fatalf("churn = %d, want 23", len(a.Churn))
	}
	if unBad, unGood := a.Drift(); unBad != 0 || unGood != 3 {
		t.Fatalf("drift = %d worse and %d better, want 0 and 3", unBad, unGood)
	}
	for _, r := range d.Reasons {
		if strings.Contains(r, "no change declared") {
			t.Errorf("the drift guard rejected identical code: %s", r)
		}
	}
	if !d.Accept {
		t.Errorf("decision = %+v, want an accept: the only declared row got better", d)
	}
}
