package evalrun

import (
	"strings"
	"testing"
)

func gated(runID, verdict string, rows ...Row) *Run {
	r := New("decks", runID)
	r.Header.Verdict = verdict
	r.LowerIsBetter("blocks", "repaired")
	r.Rows = rows
	return r
}

func row(item, metric string, value float64, kind string) Row {
	return Row{Item: item, Metric: metric, Value: value, Kind: kind}
}

// TestCompareNamesTheFlip: a gate row that moves against its sense is a
// named flip, and the verdict fails on it. A row that moves with its
// sense is a flip and not a regression.
func TestCompareNamesTheFlip(t *testing.T) {
	base := gated("run13", VerdictPass,
		row("1", "blocks", 0, KindGate), row("17", "blocks", 0, KindGate), row("17", "built", 1, KindGate),
		row("17", "repaired", 0, KindInfo), row("2", "cards", 99, KindInfo))
	next := gated("run14", VerdictPass,
		row("1", "blocks", 0, KindGate), row("17", "blocks", 1, KindGate), row("17", "built", 1, KindGate),
		row("17", "repaired", 1, KindInfo), row("2", "cards", 99, KindInfo))
	c := Compare(base, next, 0)
	if c.Verdict != VerdictFail {
		t.Errorf("verdict = %s, want FAIL on a block that appeared", c.Verdict)
	}
	if len(c.Flips) != 2 || c.Flips[0].Item != "17" || c.Flips[0].Metric != "blocks" || !c.Flips[0].Worse {
		t.Errorf("flips = %+v, want the block first and worse", c.Flips)
	}
	if c.Flips[1].Metric != "repaired" || !c.Flips[1].Worse || c.Flips[1].Kind != KindInfo {
		t.Errorf("the repaired row is information that got worse: %+v", c.Flips[1])
	}
	better := gated("run15", VerdictPass, row("17", "blocks", 0, KindGate))
	c = Compare(next, better, 0)
	if c.Verdict != VerdictPass || len(c.Flips) != 1 || c.Flips[0].Worse {
		t.Errorf("a block that went away is a better flip: %s %+v", c.Verdict, c.Flips)
	}
	if len(c.Gone) != 2 || strings.Join(c.Gone, ",") != "1,2" {
		t.Errorf("gone = %v, want the items the -only run left out, in number order", c.Gone)
	}
}

// TestCompareReadsTheSense: a higher value is better unless the header
// says otherwise, and the margin holds a small move back.
func TestCompareReadsTheSense(t *testing.T) {
	base := gated("a", VerdictPass, row("suite", "catalog_only", 27, KindGate), row("suite", "judge_agreement", 87, KindGate))
	next := gated("b", VerdictPass, row("suite", "catalog_only", 26, KindGate), row("suite", "judge_agreement", 93, KindGate))
	c := Compare(base, next, 0)
	if c.Verdict != VerdictFail || len(c.Flips) != 2 || !c.Flips[0].Worse || c.Flips[1].Worse {
		t.Errorf("catalog_only fell and judge_agreement rose: %s %+v", c.Verdict, c.Flips)
	}
	if c = Compare(base, next, 1); c.Verdict != VerdictPass {
		t.Errorf("a fall of one inside a margin of one is no regression: %s %v", c.Verdict, c.Reasons)
	}
}

// TestObserveOnlyIsNotPass: a run with no gate row, or no verdict of its
// own, is not evaluated. A run that fails its own bars fails.
func TestObserveOnlyIsNotPass(t *testing.T) {
	info := New("question-eval", "e1")
	info.Info("suite", "bad", 22, "")
	if c := Compare(nil, info, 0); c.Verdict != VerdictNotEvaluated {
		t.Errorf("verdict = %s, want not evaluated for information rows alone", c.Verdict)
	}
	noVerdict := gated("x", "", row("1", "blocks", 0, KindGate))
	if c := Compare(nil, noVerdict, 0); c.Verdict != VerdictNotEvaluated {
		t.Errorf("verdict = %s, want not evaluated with no verdict of its own", c.Verdict)
	}
	failed := gated("y", VerdictFail, row("1", "blocks", 1, KindGate))
	if c := Compare(nil, failed, 0); c.Verdict != VerdictFail {
		t.Errorf("verdict = %s, want FAIL when the run fails its own bars", c.Verdict)
	}
	first := gated("z", VerdictPass, row("1", "blocks", 0, KindGate))
	if c := Compare(nil, first, 0); c.Verdict != VerdictPass || len(c.Flips) != 0 {
		t.Errorf("the first run passes with nothing to compare: %s", c.Verdict)
	}
}

// TestEpochNotesNameWhatChanged pins the fingerprint diff.
func TestEpochNotesNameWhatChanged(t *testing.T) {
	base := gated("a", VerdictPass, row("1", "blocks", 0, KindGate))
	base.Header.Roles["generate"] = Model{Provider: "openai", Model: "gpt-5.6-terra"}
	base.Header.Prompts["generate"] = 12
	base.Header.Snapshot = "2026-09-03"
	base.Header.Versions["precons"] = "5.3.0+20260903"
	next := gated("b", VerdictPass, row("1", "blocks", 0, KindGate))
	next.Header.Roles["generate"] = Model{Provider: "anthropic", Model: "claude-opus-5", Effort: "high"}
	next.Header.Prompts["generate"] = 13
	next.Header.Snapshot = "2026-09-04"
	next.Header.Versions["precons"] = "5.3.0+20260903"
	c := Compare(base, next, 0)
	want := []string{
		"the generate role moved from gpt-5.6-terra (openai) to claude-opus-5 (anthropic, effort high)",
		"the generate prompt moved from version 12 to 13",
		"the card snapshot moved from 2026-09-03 to 2026-09-04",
	}
	if strings.Join(c.Epoch, "\n") != strings.Join(want, "\n") {
		t.Errorf("epoch notes:\n%s\nwant:\n%s", strings.Join(c.Epoch, "\n"), strings.Join(want, "\n"))
	}
	if c.Verdict != VerdictPass {
		t.Errorf("a new epoch still compares: %s", c.Verdict)
	}
}

// TestMergeReadsARunWithItsRerun: the rerun's items replace the run's,
// and the rest stay, so 14 with 14b reads as one run.
func TestMergeReadsARunWithItsRerun(t *testing.T) {
	run := gated("run14", VerdictFail, row("1", "blocks", 0, KindGate), row("17", "blocks", 1, KindGate), row("17", "cards", 0, KindInfo))
	rerun := gated("run14b", VerdictPass, row("17", "blocks", 0, KindGate), row("17", "cards", 99, KindInfo))
	m := Merge(run, rerun)
	if m.Header.RunID != "run14+run14b" || m.Header.Verdict != VerdictPass {
		t.Errorf("header = %+v", m.Header)
	}
	if len(m.Rows) != 3 || m.Rows[0].Item != "1" || m.Rows[1].Value != 0 || m.Rows[2].Value != 99 {
		t.Errorf("rows = %+v, want deck 1 kept and deck 17 replaced", m.Rows)
	}
	if !contains(m.Header.Lower, "blocks") {
		t.Error("the merge keeps the senses")
	}
	if Merge() != nil {
		t.Error("a merge of nothing is nothing")
	}
}
