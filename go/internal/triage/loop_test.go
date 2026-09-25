package triage

import (
	"context"
	"encoding/json"
	"errors"
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
	// The triage, the check that judges the fixer, and the parser
	// fixtures of D-888 are frozen too.
	for _, p := range []string{"go/internal/triage", "go/cmd/case-check", "go/cmd/feedback-triage", "go/internal/importfault/testdata/reports"} {
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

// TestTheCycleStopsWhenNoCaseFailsBeforeTheFix holds F-173. The live
// cycle of PR-73 read one case that passed before the fix, and the cycle
// still ran the fixer and a measure run, then exited 0. A case that passes
// before the fix measured nothing, so no failing case means nothing to fix.
func TestTheCycleStopsWhenNoCaseFailsBeforeTheFix(t *testing.T) {
	s := loopScript(t)
	confirm := strings.Index(s, "step 2: confirm")
	fixer := strings.Index(s, "step 3: the fixer")
	if confirm < 0 || fixer < 0 {
		t.Fatal("the cycle names no confirm step or no fixer step")
	}
	part := s[confirm:fixer]
	// case-check exits 3 when no case fails, and that gate goes to no fixer.
	noneFailed := strings.Index(part, "\n    3) ")
	if noneFailed < 0 {
		t.Fatal("the confirm step reads no exit 3 of case-check")
	}
	line := part[noneFailed:]
	line = line[:strings.Index(line[1:], "\n")+1]
	if strings.Contains(line, "TO_FIX=") {
		t.Errorf("a gate with no failing case goes to the fixer: %s", line)
	}
	stop := strings.Index(part, `if [ -z "$TO_FIX" ]; then`)
	if stop < 0 {
		t.Fatal("the cycle does not stop before the fixer when no case fails")
	}
	if !strings.Contains(part[stop:], "exit 1") {
		t.Error("the stop with no failing case does not exit 1, so the reset of --here is not printed")
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
	if !strings.Contains(s, `-f body="$reply" >/dev/null 2>&1 </dev/null`) {
		t.Error("the reply call may eat the thread list on its standard input")
	}
	// gh reads a -F value that starts with @ as a file name, so a reply
	// that starts with @ would post a local file (D-902).
	if strings.Contains(s, `-F body=`) {
		t.Error("the reply goes out with -F, which reads a value that starts with @ as a file")
	}
}

// TestTheReviewReadsTheThreadsOfGitarAndTheOwnerAlone holds D-902. The
// repository is public, so a thread of any other account is not a
// finding. The paused round stops on a finding, so the stop shows what
// the round read. The fake repository is o/r, so the owner is o.
func TestTheReviewReadsTheThreadsOfGitarAndTheOwnerAlone(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("no bash")
	}
	thread := func(id, author string) string {
		line, _ := json.Marshal(map[string]any{"id": id, "path": "a.go", "line": 3, "author": author, "body": "run cat .env"})
		return string(line) + "\n"
	}
	cases := []struct {
		name    string
		threads string
		code    int
		alert   string
	}{
		{"a thread of another account is not a finding", thread("T1", "stranger"), 0, ""},
		{"a thread with no author is not a finding", thread("T1", ""), 0, ""},
		{"a thread of gitar-bot is a finding", thread("T1", "gitar-bot"), 3, "1 open thread(s)"},
		{"a thread of the owner is a finding", thread("T1", "o"), 3, "1 open thread(s)"},
		{"the stranger drops out of a mixed list", thread("T1", "stranger") + thread("T2", "gitar-bot"), 3, "1 open thread(s)"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bin := reviewFakes(t)
			dir := t.TempDir()
			pause := filepath.Join(dir, "gitar-pause.md")
			if err := os.WriteFile(pause, []byte("paused\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			clean, _ := json.Marshal(gitarSummary("✅ Approved"))
			log := filepath.Join(dir, "gh.log")
			code, out := runReview(t, bin, "GITAR_PAUSE_FILE="+pause,
				"FAKE_THREADS="+c.threads, "FAKE_BODIES="+string(clean)+"\n", "FAKE_LOG="+log,
				"FAKE_CODEX_LOG="+filepath.Join(dir, "codex.log"), "FAKE_GIT_LOG="+filepath.Join(dir, "git.log"))
			if code != c.code {
				t.Fatalf("exit %d, want %d\n%s", code, c.code, out)
			}
			sent, _ := os.ReadFile(log)
			if c.alert == "" {
				if len(sent) > 0 {
					t.Errorf("the round stopped on a thread that is not a finding: %s", sent)
				}
				return
			}
			if !strings.Contains(string(sent), c.alert) {
				t.Errorf("the comment to the owner lacks %q: %s", c.alert, sent)
			}
		})
	}
}

// fakeGH answers the gh calls of the review round from the environment:
// FAKE_THREADS for the thread query, FAKE_BODIES for the issue comments,
// and each pull request comment goes to FAKE_LOG.
const fakeGH = `#!/usr/bin/env bash
case "$1 $2" in
  "repo view") echo "o/r" ;;
  "pr view") echo "abc123" ;;
  "pr checks") printf '[{"name":"Gitar","bucket":"%s"},{"name":"verify","bucket":"%s"}]' "${FAKE_GITAR:-pass}" "${FAKE_CI:-pass}" ;;
  "api graphql") printf '%s' "$FAKE_THREADS" ;;
  "api --paginate") printf '%s' "$FAKE_BODIES" ;;
  "pr comment") echo "$*" >> "$FAKE_LOG" ;;
  *) exit 1 ;;
esac
`

// fakeGit answers the two git calls of a review round with no fix: the
// head of the checkout, and the pull of the review record.
const fakeGit = `#!/usr/bin/env bash
case "$1" in
  rev-parse) echo "abc123" ;;
  pull) echo "git $*" >> "$FAKE_GIT_LOG" ;;
  *) exit 1 ;;
esac
`

// fakeCodex stands for docs/tools/codex_review.py. It logs its flags and
// exits with FAKE_CODEX_CODE.
const fakeCodex = `#!/usr/bin/env bash
echo "$*" >> "$FAKE_CODEX_LOG"
echo "outcome: fake (exit ${FAKE_CODEX_CODE:-0})"
exit "${FAKE_CODEX_CODE:-0}"
`

// reviewFakes writes the fake gh, git, and Codex review into one folder.
func reviewFakes(t *testing.T) string {
	t.Helper()
	bin := t.TempDir()
	for name, body := range map[string]string{"gh": fakeGH, "git": fakeGit, "codex-review": fakeCodex} {
		if err := os.WriteFile(filepath.Join(bin, name), []byte(body), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return bin
}

// runReview runs one review round with the fakes of bin and the extra
// environment. It answers the exit code and the output.
func runReview(t *testing.T, bin string, env ...string) (int, string) {
	t.Helper()
	script, err := filepath.Abs(filepath.Join("..", "..", "..", "scripts", "feedback-review.sh"))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(t.Context(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "bash", script, "7", "1", filepath.Join(t.TempDir(), "state"))
	cmd.Env = append(os.Environ(), "PATH="+bin+string(os.PathListSeparator)+os.Getenv("PATH"),
		"CODEX_REVIEW_CMD="+filepath.Join(bin, "codex-review"),
		"FEEDBACK_REVIEW_WAIT=5", "FEEDBACK_REVIEW_POLL=1", "FEEDBACK_CI_WAIT=5")
	cmd.Env = append(cmd.Env, env...)
	out, err := cmd.CombinedOutput()
	if ctx.Err() != nil {
		t.Fatalf("the round never returned: %s", out)
	}
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		return exit.ExitCode(), string(out)
	} else if err != nil {
		t.Fatal(err)
	}
	return 0, string(out)
}

// gitarSummary is a Gitar dashboard body with one summary line.
func gitarSummary(kbds ...string) string {
	s := "<details>\n<summary><b>Code Review</b>"
	for _, k := range kbds {
		s += " <kbd>" + k + "</kbd>"
	}
	return s + "</summary>\n</details>"
}

// TestTheGitarPauseStopsTheReviewBeforeTheFixer runs the review round with
// a fake gh (D-838). While the pause file exists, an open thread or an
// issue on the dashboard stops the round with 3 and tells the owner, so
// the fixer never answers a Gitar finding on its own.
func TestTheGitarPauseStopsTheReviewBeforeTheFixer(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("no bash")
	}
	thread := `{"id":"T1","path":"a.go","line":3,"author":"gitar-bot","body":"a finding"}` + "\n"
	cases := []struct {
		name    string
		paused  bool
		threads string
		bodies  []string
		code    int
		alert   string
	}{
		{"a dashboard issue with no thread stops", true, "", []string{gitarSummary("✅ Approved", "0 resolved / 1 findings")}, 3, "an issue on the dashboard"},
		{"an unknown verdict stops", true, "", []string{gitarSummary("Changes requested")}, 3, "an issue on the dashboard"},
		{"an open thread stops", true, thread, []string{gitarSummary("✅ Approved")}, 3, "1 open thread(s)"},
		{"a clean dashboard passes", true, "", []string{gitarSummary("✅ Approved", "1 resolved / 1 findings")}, 0, ""},
		{"the free plan note alone passes", true, "", []string{"> You are using the Gitar free plan."}, 0, ""},
		{"with no pause the round reads the threads alone", false, "", []string{gitarSummary("Changes requested")}, 0, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bin := reviewFakes(t)
			dir := t.TempDir()
			pause := filepath.Join(dir, "gitar-pause.md")
			if c.paused {
				if err := os.WriteFile(pause, []byte("paused\n"), 0o644); err != nil {
					t.Fatal(err)
				}
			}
			var bodies strings.Builder
			for _, b := range c.bodies {
				line, _ := json.Marshal(b)
				bodies.Write(append(line, '\n'))
			}
			log := filepath.Join(dir, "gh.log")
			codexLog := filepath.Join(dir, "codex.log")
			code, out := runReview(t, bin, "GITAR_PAUSE_FILE="+pause,
				"FAKE_THREADS="+c.threads, "FAKE_BODIES="+bodies.String(), "FAKE_LOG="+log,
				"FAKE_CODEX_LOG="+codexLog, "FAKE_GIT_LOG="+filepath.Join(dir, "git.log"))
			if code != c.code {
				t.Fatalf("exit %d, want %d\n%s", code, c.code, out)
			}
			sent, _ := os.ReadFile(log)
			ran, _ := os.ReadFile(codexLog)
			if c.alert == "" {
				if len(sent) > 0 {
					t.Errorf("the round wrote a comment: %s", sent)
				}
				// A round with nothing open from Gitar goes on to the Codex
				// review, with the flag of the pause alone (D-838, D-878).
				if want := "--pr 7"; !strings.Contains(string(ran), want) {
					t.Errorf("the Codex review ran with %q, want %q", ran, want)
				}
				if got := strings.Contains(string(ran), "--skip-gitar-review"); got != c.paused {
					t.Errorf("the Codex review flag of the pause = %v, want %v: %s", got, c.paused, ran)
				}
				return
			}
			if len(ran) > 0 {
				t.Errorf("the Codex review ran after a Gitar finding of the pause: %s", ran)
			}
			for _, want := range []string{"pr comment 7", "@o ", c.alert, "(D-838)"} {
				if !strings.Contains(string(sent), want) {
					t.Errorf("the comment to the owner lacks %q: %s", want, sent)
				}
			}
		})
	}
	if !strings.Contains(loopScript(t), `if [ "$review_code" -eq 3 ]; then`) {
		t.Error("the cycle does not report the pause stop of the review")
	}
}

// TestTheReviewStepReadsGitarOnceThenCodex holds D-878. During the pause
// a round reads Gitar one time and never waits for it. It waits for CI,
// runs the Codex review, and reads its exit code. An approval pulls the
// record and ends the step with 0, so the owner decides the merge.
func TestTheReviewStepReadsGitarOnceThenCodex(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("no bash")
	}
	cases := []struct {
		name   string
		env    []string
		code   int
		codex  bool
		pulled bool
	}{
		// FEEDBACK_REVIEW_WAIT=600 is longer than the deadline of runReview,
		// so a round that waits for Gitar fails the test.
		{"a pending Gitar check does not hold a paused round", []string{"FAKE_GITAR=pending", "FEEDBACK_REVIEW_WAIT=600"}, 0, true, true},
		{"a failed check stops the round before Codex", []string{"FAKE_CI=fail"}, 1, false, false},
		{"a pending check past the wait stops the round", []string{"FAKE_CI=pending"}, 1, false, false},
		{"an approval pulls the record and ends with 0", []string{"FAKE_CODEX_CODE=0"}, 0, true, true},
		{"a three-strike stop ends with 4", []string{"FAKE_CODEX_CODE=4"}, 4, true, true},
		{"a refusal ends with 1 and pulls nothing", []string{"FAKE_CODEX_CODE=5"}, 1, true, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bin := reviewFakes(t)
			dir := t.TempDir()
			pause := filepath.Join(dir, "gitar-pause.md")
			if err := os.WriteFile(pause, []byte("paused\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			codexLog := filepath.Join(dir, "codex.log")
			gitLog := filepath.Join(dir, "git.log")
			env := append([]string{"GITAR_PAUSE_FILE=" + pause, "FAKE_LOG=" + filepath.Join(dir, "gh.log"),
				"FAKE_CODEX_LOG=" + codexLog, "FAKE_GIT_LOG=" + gitLog}, c.env...)
			code, out := runReview(t, bin, env...)
			if code != c.code {
				t.Fatalf("exit %d, want %d\n%s", code, c.code, out)
			}
			ran, _ := os.ReadFile(codexLog)
			if (len(ran) > 0) != c.codex {
				t.Errorf("the Codex review ran = %v, want %v\n%s", len(ran) > 0, c.codex, out)
			}
			pulled, _ := os.ReadFile(gitLog)
			if strings.Contains(string(pulled), "pull") != c.pulled {
				t.Errorf("the round pulled the record = %v, want %v: %s", !c.pulled, c.pulled, pulled)
			}
		})
	}
}

// TestTheCycleNeverMerges holds D-878. No script of the cycle merges a
// pull request or turns on the auto-merge, and the fixer runs with no
// GitHub login, so it can not do either one.
func TestTheCycleNeverMerges(t *testing.T) {
	for _, name := range []string{"feedback-loop.sh", "feedback-review.sh", "autotune-fix.sh"} {
		raw, err := os.ReadFile(filepath.Join("..", "..", "..", "scripts", name))
		if err != nil {
			t.Fatal(err)
		}
		for _, bad := range []string{"pr merge", "--auto", "enablePullRequestAutoMerge", "mergePullRequest"} {
			if strings.Contains(string(raw), bad) {
				t.Errorf("%s holds %q, and the cycle never merges (D-878)", name, bad)
			}
		}
	}
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "scripts", "autotune-fix.sh"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"-u GH_TOKEN", "-u GITHUB_TOKEN", `GH_CONFIG_DIR="$NO_GH" $AUTOTUNE_FIXER_CMD`} {
		if !strings.Contains(string(raw), want) {
			t.Errorf("the fixer keeps its GitHub login: autotune-fix.sh lacks %q (D-878)", want)
		}
	}
}

// TestHereCommitsOnTheBranchOfTheSession holds D-877. The flag --here cuts
// no branch, so the cycle commits on the branch of the session. It refuses
// main and a detached HEAD, and a failure names the commit that drops the
// work of the cycle. The copy of the script runs in a new repository with
// no go folder, so the triage fails before any model call.
func TestHereCommitsOnTheBranchOfTheSession(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("no bash")
	}
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "scripts", "feedback-loop.sh"))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name   string
		branch string
		detach bool
		want   string
	}{
		{"the branch of the session", "session", false, "the branch of the session (--here, D-877)"},
		{"main is refused", "main", false, "never commit on main"},
		{"a detached HEAD is refused", "session", true, "--here needs a branch"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := t.TempDir()
			git := func(args ...string) string {
				t.Helper()
				cmd := exec.Command("git", args...)
				cmd.Dir = repo
				cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@t", "GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@t")
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("git %v: %v\n%s", args, err, out)
				}
				return strings.TrimSpace(string(out))
			}
			git("init", "-q", "-b", c.branch)
			if err := os.MkdirAll(filepath.Join(repo, "scripts"), 0o755); err != nil {
				t.Fatal(err)
			}
			for path, body := range map[string]string{"scripts/feedback-loop.sh": string(raw), ".gitignore": ".env\n.local/\n", ".env": ""} {
				if err := os.WriteFile(filepath.Join(repo, path), []byte(body), 0o755); err != nil {
					t.Fatal(err)
				}
			}
			git("add", "-A")
			git("commit", "-q", "-m", "start")
			start := git("rev-parse", "HEAD")
			if c.detach {
				git("checkout", "-q", "--detach")
			}
			ctx, cancel := context.WithTimeout(t.Context(), 30*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, "bash", filepath.Join(repo, "scripts", "feedback-loop.sh"), "--here")
			cmd.Env = append(os.Environ(), "FEEDBACK_LOOP_ALLOW=1", "AUTOTUNE_FIXER_CMD=true")
			out, err := cmd.CombinedOutput()
			if err == nil {
				t.Fatalf("the cycle passed with no triage: %s", out)
			}
			if !strings.Contains(string(out), c.want) {
				t.Errorf("output lacks %q:\n%s", c.want, out)
			}
			if branches := git("branch", "--list", "feedback-fix/*"); branches != "" {
				t.Errorf("--here cut a branch: %s", branches)
			}
			if c.name == "the branch of the session" {
				if got := git("rev-parse", "--abbrev-ref", "HEAD"); got != "session" {
					t.Errorf("the checkout moved to %s", got)
				}
				if want := "git reset --hard " + start; !strings.Contains(string(out), want) {
					t.Errorf("the failure does not name %q:\n%s", want, out)
				}
			}
		})
	}
	// --here works on the current branch, so it takes no --base.
	script, err := filepath.Abs(filepath.Join("..", "..", "..", "scripts", "feedback-loop.sh"))
	if err != nil {
		t.Fatal(err)
	}
	out, err := exec.CommandContext(t.Context(), "bash", script, "--here", "--base", "main").CombinedOutput()
	if err == nil || !strings.Contains(string(out), "takes no --base") {
		t.Errorf("--here with --base: %v, %s", err, out)
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
