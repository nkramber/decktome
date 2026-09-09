package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/triage"
)

// A case is the fault a reader met. Before the fixer it must read fail,
// and after it must read pass, so both sides of the check are fired here
// on purpose (D-234).

func manifest() triage.Manifest {
	return triage.Manifest{Cases: []triage.Entry{
		{Class: "Q1", From: "f01", Gate: triage.GateQuestions, Target: triage.TargetConversations, ID: 110},
		{Class: "C1", From: "f04", Gate: triage.GateDecks, Target: triage.TargetDeckPrompts, ID: 26},
	}}
}

// runOf writes a question-gate shaped run for one item.
func runOf(item string, rows map[string]float64) *evalrun.Run {
	rec := evalrun.New("questions", "test")
	rec.LowerIsBetter("error", "premature", "dead_end", "lint_findings")
	for m, v := range rows {
		rec.Gate(item, m, v, "")
	}
	return rec
}

// TestAFailingCaseReadsFail is the confirm run: the fault the reader met
// is real, so the case fails before the fixer touches anything.
func TestAFailingCaseReadsFail(t *testing.T) {
	rec := runOf("110", map[string]float64{
		"error": 0, "premature": 0, "lint_findings": 0,
		"never_asked_power_commander": 0, "slot_format": 1,
	})
	var b bytes.Buffer
	code, err := check(&b, manifest(), rec, triage.GateQuestions, "fail", "run.jsonl")
	if err != nil || code != 0 {
		t.Fatalf("code = %d, err = %v, want a clean confirm run\n%s", code, err, b.String())
	}
	if !strings.Contains(b.String(), "never_asked_power_commander") {
		t.Errorf("the row does not name the bar that failed:\n%s", b.String())
	}
	// The same run after the fixer is the failure of the cycle.
	b.Reset()
	code, err = check(&b, manifest(), rec, triage.GateQuestions, "pass", "run.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	if code != 1 {
		t.Errorf("code = %d, want 1: the case still fails after the fixer", code)
	}
	if !strings.Contains(b.String(), "1 of 1 cases did not read pass") {
		t.Errorf("the count is missing:\n%s", b.String())
	}
}

// TestAPassingCaseReadsPass is the measure run after the fixer.
func TestAPassingCaseReadsPass(t *testing.T) {
	rec := runOf("110", map[string]float64{
		"error": 0, "premature": 0, "lint_findings": 0,
		"never_asked_power_commander": 1, "slot_format": 1,
	})
	var b bytes.Buffer
	code, err := check(&b, manifest(), rec, triage.GateQuestions, "pass", "run.jsonl")
	if err != nil || code != 0 {
		t.Fatalf("code = %d, err = %v\n%s", code, err, b.String())
	}
	if !strings.Contains(b.String(), "Every one of the 1 cases reads pass") {
		t.Errorf("the summary is missing:\n%s", b.String())
	}
	// A case that already passed on the confirm run never measured the
	// reader's fault, and the cycle must hear about it.
	b.Reset()
	code, _ = check(&b, manifest(), rec, triage.GateQuestions, "fail", "run.jsonl")
	if code != 1 {
		t.Errorf("code = %d, want 1: a case that passes before the fix measured nothing", code)
	}
}

// TestALowerIsBetterBarReadsTheOtherWay keeps the two senses apart. A
// lint finding is good at zero, and a slot bar is good at one.
func TestALowerIsBetterBarReadsTheOtherWay(t *testing.T) {
	rec := runOf("110", map[string]float64{"lint_findings": 2, "slot_format": 1})
	var b bytes.Buffer
	code, _ := check(&b, manifest(), rec, triage.GateQuestions, "pass", "run.jsonl")
	if code != 1 {
		t.Errorf("code = %d, want 1: two lint findings are a failure", code)
	}
	if !strings.Contains(b.String(), "lint_findings") {
		t.Errorf("the failing bar is not named:\n%s", b.String())
	}
	clean := runOf("110", map[string]float64{"lint_findings": 0, "slot_format": 1})
	b.Reset()
	if code, _ := check(&b, manifest(), clean, triage.GateQuestions, "pass", "run.jsonl"); code != 0 {
		t.Errorf("code = %d, want 0: no lint finding is a pass\n%s", code, b.String())
	}
}

// TestARunThatMeasuredAnotherItemIsAFault holds T-12: a tool fault is
// not a failing case. A cycle that read a run of the wrong ids would
// otherwise revert the fixer's work and blame the fixer.
func TestARunThatMeasuredAnotherItemIsAFault(t *testing.T) {
	rec := runOf("7", map[string]float64{"slot_format": 1})
	var b bytes.Buffer
	code, err := check(&b, manifest(), rec, triage.GateQuestions, "pass", "run.jsonl")
	if code != exitFault {
		t.Errorf("code = %d, want %d on a run with no row for the case", code, exitFault)
	}
	if err == nil || !strings.Contains(err.Error(), "measured something else") {
		t.Errorf("err = %v, want the fault named", err)
	}
	if !strings.Contains(b.String(), "no row") {
		t.Errorf("the row does not say the run holds nothing:\n%s", b.String())
	}
}

// TestAGateWithNoCaseIsAFault stops a cycle that would report a pass
// over an empty list.
func TestAGateWithNoCaseIsAFault(t *testing.T) {
	rec := runOf("110", map[string]float64{"slot_format": 1})
	var b bytes.Buffer
	code, err := check(&b, manifest(), rec, triage.GateBrackets, "pass", "run.jsonl")
	if code != exitFault || err == nil {
		t.Errorf("code = %d, err = %v, want a fault on a gate with no case", code, err)
	}
}

// TestTheGateFilterReadsOneSuite keeps the cases of one gate apart from
// the rest. The cycle runs one gate at a time, and a deck case must not
// count against a question run.
func TestTheGateFilterReadsOneSuite(t *testing.T) {
	rec := runOf("110", map[string]float64{"slot_format": 1})
	var b bytes.Buffer
	if code, err := check(&b, manifest(), rec, triage.GateQuestions, "pass", "run.jsonl"); code != 0 || err != nil {
		t.Fatalf("code = %d, err = %v\n%s", code, err, b.String())
	}
	if strings.Contains(b.String(), "| 26 |") {
		t.Errorf("the deck case reached the question run:\n%s", b.String())
	}
	// With no gate named, the deck case has no row and the check faults.
	b.Reset()
	if code, _ := check(&b, manifest(), rec, "", "pass", "run.jsonl"); code != exitFault {
		t.Errorf("code = %d, want a fault when a case of another gate has no row", code)
	}
}
