package triage

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// The fix cycle of PR-28c is a shell script, and shellcheck reads its
// shape. These tests read its content: a frozen path that no longer
// exists protects nothing, and a gate file the list forgets is a file
// the fixer may edit to make its own measurement agree with it (D-645).

// loopScript reads the cycle script.
func loopScript(t *testing.T) string {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "scripts", "feedback-loop.sh"))
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

// frozenPaths reads the FROZEN block of the script.
func frozenPaths(t *testing.T) []string {
	t.Helper()
	s := loopScript(t)
	start := strings.Index(s, "FROZEN=\"\n")
	if start < 0 {
		t.Fatal("the cycle names no FROZEN block")
	}
	rest := s[start+len("FROZEN=\"\n"):]
	end := strings.Index(rest, "\n\"")
	if end < 0 {
		t.Fatal("the FROZEN block does not close")
	}
	var out []string
	for _, line := range strings.Split(rest[:end], "\n") {
		if p := strings.TrimSpace(line); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// TestEveryFrozenPathExists keeps the list honest. A path that moved
// still reads in the list, and the check that names it then guards a
// file nobody has.
func TestEveryFrozenPathExists(t *testing.T) {
	paths := frozenPaths(t)
	if len(paths) < 10 {
		t.Fatalf("the FROZEN block holds %d paths, which is too few to be the real list", len(paths))
	}
	for _, p := range paths {
		if _, err := os.Stat(filepath.Join("..", "..", "..", p)); err != nil {
			t.Errorf("the cycle freezes %s and no such path exists: %v", p, err)
		}
	}
}

// TestTheCaseFilesAreFrozen is the guard that matters most. A case is
// the measurement, and a fixer that edits one makes the gate agree with
// the code instead of with the reader.
func TestTheCaseFilesAreFrozen(t *testing.T) {
	paths := frozenPaths(t)
	held := map[string]bool{}
	for _, p := range paths {
		held[p] = true
	}
	for _, target := range []string{TargetConversations, TargetDeckPrompts, TargetBracketPrompts} {
		if !held[target] {
			t.Errorf("the cycle does not freeze %s, so the fixer may edit its own measurement", target)
		}
	}
	// The triage and the check that judges the fixer are frozen too.
	for _, p := range []string{"go/internal/triage", "go/cmd/case-check", "go/cmd/feedback-triage"} {
		if !held[p] {
			t.Errorf("the cycle does not freeze %s", p)
		}
	}
}

// TestTheFixerPromptAndTheCycleAgreeOnTheFrozenList keeps the two lists
// in step. The prompt tells the fixer what it may not touch, and the
// cycle enforces it. A path in one and not the other either reverts a
// fixer that was never warned, or warns about a file nothing checks.
func TestTheFixerPromptAndTheCycleAgreeOnTheFrozenList(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "docs", "reference", "feedback-fixer-prompt.md"))
	if err != nil {
		t.Fatal(err)
	}
	prompt := string(raw)
	for _, p := range frozenPaths(t) {
		if !strings.Contains(prompt, "`"+p+"`") && !strings.Contains(prompt, "`"+p+"/`") {
			t.Errorf("the cycle freezes %s and the fixer prompt does not name it", p)
		}
	}
}

// TestTheCycleNeedsItsOwnPermission holds the unattended guard. The
// cycle edits code, commits, pushes, and writes to a public pull
// request, so it may never start by accident.
func TestTheCycleNeedsItsOwnPermission(t *testing.T) {
	s := loopScript(t)
	for _, want := range []string{
		"FEEDBACK_LOOP_ALLOW",
		"AUTOTUNE_FIXER_CMD",
		// The cap of D-559, and the default the design note names.
		`CAP="2.00"`,
	} {
		if !strings.Contains(s, want) {
			t.Errorf("the cycle does not name %s", want)
		}
	}
	// A dry run must reach neither guard, so a person can read the plan
	// with no permission and no agent.
	if !strings.Contains(s, `if [ "$DRY_RUN" != "1" ]; then`) {
		t.Error("the cycle asks for its permission on a dry run too")
	}
}

// TestTheCycleStopsWhenTheCapLeavesACaseUnmeasured holds the rule under
// the ledger. A gate the cap stopped leaves cases nobody read, and those
// cases are in the gate files already. A cycle that went on would open a
// pull request carrying a case no run ever measured.
func TestTheCycleStopsWhenTheCapLeavesACaseUnmeasured(t *testing.T) {
	s := loopScript(t)
	if !strings.Contains(s, "capped=1") {
		t.Error("the confirm loop does not record a gate the cap stopped")
	}
	if !strings.Contains(s, `if [ "$capped" != "0" ]; then`) {
		t.Error("the cycle does not stop before the fixer on a capped confirm run")
	}
	// The stop comes before the fixer, so the cases stay and the fixer
	// never reads a part of them.
	capIdx := strings.Index(s, `if [ "$capped" != "0" ]; then`)
	fixIdx := strings.Index(s, "step 3: the fixer")
	if capIdx < 0 || fixIdx < 0 || capIdx > fixIdx {
		t.Error("the capped check does not sit before the fixer")
	}
}

// TestTheReviewCountsAnEmptyThreadListAsNone reads the fault that made
// this test: grep -c prints 0 and exits 1 on an empty file, so an
// "|| echo 0" wrote a second line and the count never read 0. The review
// then handed the fixer a findings file with no finding in it.
func TestTheReviewCountsAnEmptyThreadListAsNone(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "scripts", "feedback-review.sh"))
	if err != nil {
		t.Fatal(err)
	}
	s := string(raw)
	if strings.Contains(s, `grep -c . "$threads"`) {
		t.Error("the review counts threads with grep -c, which prints a second line on an empty file")
	}
	if !strings.Contains(s, `awk 'NF' "$threads"`) {
		t.Error("the review does not count the thread lines")
	}
	// A command inside the reply loop must not read the thread list.
	if !strings.Contains(s, `-F body="$reply" >/dev/null 2>&1 </dev/null`) {
		t.Error("the reply call may eat the thread list on its standard input")
	}
}

// TestAnEmptyManifestMarshalsAsAList holds the fault gitar found on
// pull request #122. A nil slice marshals as null, and the cycle reads
// the length of that list to decide whether there is anything to fix.
// A null raised instead of reading zero, so the cycle went on to a fixer
// with no case at all.
func TestAnEmptyManifestMarshalsAsAList(t *testing.T) {
	// Every verdict of this run writes a defect or wants the judge, so no
	// gate measures one. The gate document names it: four of the ten
	// fixture verdicts reach no gate.
	m := ManifestOf("h.jsonl", time.Now(), nil)
	if m.Cases == nil {
		t.Fatal("the case list is nil, and it marshals as null")
	}
	raw, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"cases":[]`) {
		t.Errorf("an empty manifest reads %s, want an empty list", raw)
	}
	// The cycle reads the file, so the round trip has to answer zero.
	path := filepath.Join(t.TempDir(), "cases.json")
	if err := WriteManifest(path, m); err != nil {
		t.Fatal(err)
	}
	got, err := ReadManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Cases) != 0 {
		t.Errorf("%d cases read back, want none", len(got.Cases))
	}
	if got.IDs(GateQuestions) != nil || len(got.Gates()) != 0 {
		t.Error("an empty manifest names an id or a gate")
	}
}

// TestAValueFlagWithNoValueStopsTheCycle holds the other fault of #122.
// "shift 2" fails with one argument left, nothing shifts, and the parse
// loop spins forever. The cycle must refuse the flag and stop.
func TestAValueFlagWithNoValueStopsTheCycle(t *testing.T) {
	script, err := filepath.Abs(filepath.Join("..", "..", "..", "scripts", "feedback-loop.sh"))
	if err != nil {
		t.Fatal(err)
	}
	for _, flag := range []string{"--cap", "--in", "--base", "--rounds"} {
		t.Run(flag, func(t *testing.T) {
			// The deadline is the test: the old code never returned.
			ctx, cancel := context.WithTimeout(t.Context(), 15*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, "bash", script, flag)
			out, err := cmd.CombinedOutput()
			if ctx.Err() != nil {
				t.Fatalf("%s with no value never returned, so the parser spins: %s", flag, out)
			}
			if err == nil {
				t.Errorf("%s with no value was accepted: %s", flag, out)
			}
			if !strings.Contains(string(out), "needs a value") {
				t.Errorf("%s: output = %s, want the refusal", flag, out)
			}
		})
	}
}
