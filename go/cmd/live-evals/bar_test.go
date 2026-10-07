package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/livespend"
)

const (
	baseTree = "1111111111111111111111111111111111111111"
	headTree = "2222222222222222222222222222222222222222"
	oldTree  = "3333333333333333333333333333333333333333"
)

func score(x float64) *float64 { return &x }

// barCase is the verdict of #294 in this shape: six base replays of 17
// to 23 themed cards, five fix replays of 31, and a target of 29.
func barCase() (verdict, []livespend.Run) {
	v := verdict{Bar: "the worst fix deck plays 29 or more of the 31 themed nonland cards",
		Metric: "themed nonland cards played", Better: "higher", Target: score(29)}
	var runs []livespend.Run
	add := func(side, tree string, s *float64) {
		id := fmt.Sprintf("%016x", len(runs)+1)
		runs = append(runs, livespend.Run{ID: id, Target: replayTarget, Tree: tree, Ended: true})
		v.Replays = append(v.Replays, replayV{Run: id, Side: side, Score: s})
	}
	for _, s := range []float64{17, 20, 23, 19, 22, 18} {
		add("base", baseTree, score(s))
	}
	for range 5 {
		add("fix", headTree, score(31))
	}
	return v, runs
}

func TestBarHoldsForTheReplaysOf294(t *testing.T) {
	v, runs := barCase()
	if why := barFaults(v, runs, baseTree, headTree); len(why) > 0 {
		t.Fatalf("faults = %v, want none", why)
	}
}

func TestBarRefusesAWeakVerdict(t *testing.T) {
	cases := map[string]struct {
		edit func(*verdict, *[]livespend.Run)
		want string
	}{
		"one replay of each side, as the eval of #294 did": {func(v *verdict, r *[]livespend.Run) {
			v.Replays = []replayV{v.Replays[0], v.Replays[6]}
			*r = []livespend.Run{(*r)[0], (*r)[6]}
		}, "needs 3"},
		"a fix replay no better than the best base replay": {func(v *verdict, _ *[]livespend.Run) {
			v.Replays[6].Score = score(23)
		}, "does not beat the best base replay (23)"},
		"a fix that beats the base and misses the target": {func(v *verdict, _ *[]livespend.Run) {
			v.Replays[6].Score = score(26)
		}, "does not reach the target (29)"},
		"a fix replay left out of the verdict": {func(v *verdict, _ *[]livespend.Run) {
			v.Replays = v.Replays[:10]
		}, "on the fix code is absent"},
		"a base replay left out of the verdict": {func(v *verdict, _ *[]livespend.Run) {
			v.Replays = v.Replays[1:]
		}, "on the base code is absent"},
		"a fix replay of an older head": {func(_ *verdict, r *[]livespend.Run) {
			(*r)[7].Tree = oldTree
		}, "built another tree than the fix code"},
		"a replay of a change that no commit holds": {func(_ *verdict, r *[]livespend.Run) {
			(*r)[8].Tree = ""
		}, "built another tree than the fix code"},
		"a run that no session log names": {func(_ *verdict, r *[]livespend.Run) {
			*r = (*r)[:10]
		}, "no session log names it"},
		"a score for a run with no end line": {func(_ *verdict, r *[]livespend.Run) {
			(*r)[9].Ended = false
		}, "no end line in a log"},
		"a run listed two times": {func(v *verdict, _ *[]livespend.Run) {
			v.Replays = append(v.Replays, v.Replays[6])
		}, "listed two times"},
		"a run of another paid target": {func(_ *verdict, r *[]livespend.Run) {
			(*r)[6].Target = "deck-gate"
		}, "is not a replay"},
		"no target":    {func(v *verdict, _ *[]livespend.Run) { v.Target = nil }, "no target"},
		"no direction": {func(v *verdict, _ *[]livespend.Run) { v.Better = "" }, `"better"`},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			v, runs := barCase()
			c.edit(&v, &runs)
			why := barFaults(v, runs, baseTree, headTree)
			if !strings.Contains(strings.Join(why, "; "), c.want) {
				t.Fatalf("faults = %v, want one with %q", why, c.want)
			}
		})
	}
}

func TestBarRefusesAHeadThatChangesNoGoFile(t *testing.T) {
	v, runs := barCase()
	why := barFaults(v, runs, baseTree, baseTree)
	if len(why) != 1 || !strings.Contains(why[0], "changes no file of the Go module") {
		t.Fatalf("faults = %v", why)
	}
}

func TestBarDropsAReplayWithNoDeck(t *testing.T) {
	// OQ-98: a replay of the question phase can build no deck. Such a run
	// stays in the verdict with no score, and it counts on no side (D-1205).
	v, runs := barCase()
	v.Replays[0].Score = nil
	runs[0].Ended = false
	v.Replays[6].Score = nil
	if why := barFaults(v, runs, baseTree, headTree); len(why) > 0 {
		t.Fatalf("faults = %v, want none", why)
	}
	v.Replays[7].Score, v.Replays[8].Score = nil, nil
	if why := barFaults(v, runs, baseTree, headTree); !strings.Contains(strings.Join(why, "; "), "the fix side has 2 replays") {
		t.Fatalf("faults = %v, want too few fix replays", why)
	}
}

func TestBarReadsALowerIsBetterMetric(t *testing.T) {
	v, runs := barCase()
	v.Better, v.Target = "lower", score(2)
	for i := range v.Replays {
		if v.Replays[i].Side == "base" {
			v.Replays[i].Score = score(8)
		} else {
			v.Replays[i].Score = score(1)
		}
	}
	if why := barFaults(v, runs, baseTree, headTree); len(why) > 0 {
		t.Fatalf("faults = %v, want none", why)
	}
	v.Replays[10].Score = score(3)
	if why := barFaults(v, runs, baseTree, headTree); !strings.Contains(strings.Join(why, "; "), "does not reach the target (2)") {
		t.Fatalf("faults = %v, want the target missed", why)
	}
}

// gitRun runs one git command in dir for a test.
func gitRun(t *testing.T, dir string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@example.com",
		"GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@example.com")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func TestBarReadsTheLogsAndTheTreesOfTheRepository(t *testing.T) {
	repo := t.TempDir()
	mod := filepath.Join(repo, moduleDir)
	write := func(p, body string) {
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write(filepath.Join(mod, "go.mod"), "module x\n")
	gitRun(t, repo, "init", "-q")
	gitRun(t, repo, "add", ".")
	gitRun(t, repo, "commit", "-q", "-m", "base")
	base := gitRun(t, repo, "rev-parse", "HEAD")
	write(filepath.Join(mod, "fix.go"), "package x\n")
	gitRun(t, repo, "add", ".")
	gitRun(t, repo, "commit", "-q", "-m", "fix")
	// A commit of documents alone keeps the tree of the module.
	write(filepath.Join(repo, "docs", "handoff.md"), "x\n")
	gitRun(t, repo, "add", ".")
	gitRun(t, repo, "commit", "-q", "-m", "docs")
	trees := map[string]string{"base": gitRun(t, repo, "rev-parse", base+":go"), "fix": gitRun(t, repo, "rev-parse", "HEAD:go")}

	v := verdict{Bar: "b", Metric: "m", Better: "higher", Target: score(5)}
	var log strings.Builder
	for i, s := range []float64{1, 2, 3, 6, 7, 8} {
		side := "base"
		if i >= 3 {
			side = "fix"
		}
		id := fmt.Sprintf("%016x", i+1)
		v.Replays = append(v.Replays, replayV{Run: id, Side: side, Score: score(s)})
		for _, l := range []string{
			fmt.Sprintf(`{"id":%q,"event":"start","target":"chat-probe","tree":%q,"usd":0,"measured":false}`, id, trees[side]),
			fmt.Sprintf(`{"id":%q,"event":"end","target":"chat-probe","usd":0.03,"measured":true}`, id),
		} {
			raw, _ := json.Marshal(map[string]any{"type": "user", "message": map[string]any{"content": []any{
				map[string]any{"type": "tool_result", "content": livespend.Marker + l},
			}}})
			log.Write(append(raw, '\n'))
		}
	}
	logs := t.TempDir()
	write(filepath.Join(logs, "session-1.log"), log.String())
	raw, _ := json.Marshal(v)
	vp := filepath.Join(t.TempDir(), "verdict.json")
	write(vp, string(raw))

	var out bytes.Buffer
	if err := bar(vp, logs, repo, base, "HEAD", &out); err != nil {
		t.Fatalf("bar: %v", err)
	}
	if !strings.HasPrefix(out.String(), "bar holds") {
		t.Fatalf("out = %q", out.String())
	}
	// The base as the head measures no fix.
	var bf barFailed
	if err := bar(vp, logs, repo, base, base, &out); !errors.As(err, &bf) {
		t.Fatalf("err = %v, want barFailed", err)
	}
}
