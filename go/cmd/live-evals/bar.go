package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"strings"

	"github.com/nkramber/decktome/go/internal/livespend"
)

// minReplays is the number of replays with a score that each side of
// the bar needs (D-1200).
const minReplays = 3

// passFailFixReplays is the number of fix replays with a score that a
// bar of pass or fail needs (D-1214). A fix that does nothing passes six
// of six with a chance of 1 in 64 when half of the base replays fail.
const passFailFixReplays = 6

// passFail is the scale of a bar of pass or fail. Each score is 1 for a
// pass and 0 for a fail (D-1214).
const passFail = "pass-fail"

// replayTarget is the target of the spend lines of a replay.
const replayTarget = "chat-probe"

// moduleDir is the folder of the Go module in the repository. A replay
// builds it, so its tree names the code of the replay (D-1206).
const moduleDir = "go"

// barFailed is the answer of `bar` for a verdict that does not hold the
// bar. main exits 5 on it.
type barFailed []string

func (b barFailed) Error() string { return "the bar does not hold: " + strings.Join(b, "; ") }

// verdict is verdict.json, which the eval session writes (D-1200). Each
// replay names the id of its spend lines. A replay that never reached
// the fault has no score (D-1205, D-1212). Scale is passFail for a bar
// of pass or fail, and empty for a count (D-1214).
type verdict struct {
	Bar     string    `json:"bar"`
	Metric  string    `json:"metric"`
	Better  string    `json:"better"`
	Target  *float64  `json:"target"`
	Scale   string    `json:"scale"`
	Replays []replayV `json:"replays"`
}

type replayV struct {
	Run   string   `json:"run"`
	Side  string   `json:"side"`
	Score *float64 `json:"score"`
}

// barFaults answers each reason that the verdict does not hold the bar,
// and none when it holds (D-1200, D-1205, D-1206). runs come from the
// session logs alone. baseTree and headTree are the trees of the Go
// module at the base and at the head of the pull request. The scores are
// the word of the session, and the rest is the record of the logs.
func barFaults(v verdict, runs []livespend.Run, baseTree, headTree string) []string {
	var why []string
	if strings.TrimSpace(v.Bar) == "" || strings.TrimSpace(v.Metric) == "" {
		why = append(why, "verdict.json names no bar or no metric")
	}
	higher := v.Better == "higher"
	if !higher && v.Better != "lower" {
		why = append(why, `"better" is not "higher" or "lower"`)
	}
	if v.Target == nil {
		why = append(why, "verdict.json names no target")
	}
	pf := v.Scale == passFail
	if v.Scale != "" && !pf {
		why = append(why, fmt.Sprintf("scale %q is not %q or empty", v.Scale, passFail))
	}
	if pf && (!higher || v.Target == nil || *v.Target != 1) {
		why = append(why, `a bar of pass or fail needs "better" "higher" and the target 1`)
	}
	if baseTree == "" || headTree == "" {
		return append(why, "the tree of the base or of the head does not read")
	}
	if baseTree == headTree {
		return append(why, "the head changes no file of the Go module, so no replay measures the fix")
	}
	byID := map[string]livespend.Run{}
	for _, r := range runs {
		byID[r.ID] = r
	}
	tree := map[string]string{"base": baseTree, "fix": headTree}
	listed := map[string]bool{}
	scores := map[string][]float64{}
	for _, rp := range v.Replays {
		if listed[rp.Run] {
			why = append(why, fmt.Sprintf("run %s is listed two times", rp.Run))
			continue
		}
		listed[rp.Run] = true
		want, ok := tree[rp.Side]
		if !ok {
			why = append(why, fmt.Sprintf("run %s: side %q is not base or fix", rp.Run, rp.Side))
			continue
		}
		r, ok := byID[rp.Run]
		switch {
		case !ok:
			why = append(why, fmt.Sprintf("run %s: no session log names it", rp.Run))
		case r.Target != replayTarget:
			why = append(why, fmt.Sprintf("run %s: target %s is not a replay", rp.Run, r.Target))
		case r.Tree != want:
			why = append(why, fmt.Sprintf("run %s: it built another tree than the %s code", rp.Run, rp.Side))
		case rp.Score != nil && !r.Ended:
			why = append(why, fmt.Sprintf("run %s: it has a score and no end line in a log", rp.Run))
		case rp.Score != nil && pf && *rp.Score != 0 && *rp.Score != 1:
			why = append(why, fmt.Sprintf("run %s: score %g is not 1 or 0 on a bar of pass or fail", rp.Run, *rp.Score))
		case rp.Score != nil:
			scores[rp.Side] = append(scores[rp.Side], *rp.Score)
		}
	}
	for _, r := range runs {
		if r.Target != replayTarget || listed[r.ID] {
			continue
		}
		for side, t := range tree {
			if r.Tree == t {
				why = append(why, fmt.Sprintf("run %s on the %s code is absent from verdict.json", r.ID, side))
			}
		}
	}
	need := map[string]int{"base": minReplays, "fix": minReplays}
	if pf {
		need["fix"] = passFailFixReplays
	}
	for _, side := range []string{"base", "fix"} {
		if n := len(scores[side]); n < need[side] {
			why = append(why, fmt.Sprintf("the %s side has %d replays with a score, and it needs %d", side, n, need[side]))
		}
	}
	if len(why) > 0 || v.Target == nil {
		return why
	}
	if pf {
		return passFailFaults(scores["base"], scores["fix"])
	}
	worstFix, bestBase := worst(scores["fix"], higher), worst(scores["base"], !higher)
	if higher && worstFix <= bestBase || !higher && worstFix >= bestBase {
		why = append(why, fmt.Sprintf("the worst fix replay (%g) does not beat the best base replay (%g)", worstFix, bestBase))
	}
	if higher && worstFix < *v.Target || !higher && worstFix > *v.Target {
		why = append(why, fmt.Sprintf("the worst fix replay (%g) does not reach the target (%g)", worstFix, *v.Target))
	}
	return why
}

// passFailFaults is the counted rule of a bar of pass or fail (D-1214).
// The fault must show in half or more of the base replays, so the base
// proves it. Each fix replay must pass. A fault that shows on some base
// replays alone can never meet the rule of the worst and the best.
func passFailFaults(base, fix []float64) []string {
	var why []string
	if n := fails(base); 2*n < len(base) {
		why = append(why, fmt.Sprintf("the fault shows in %d of %d base replays, and the bar needs half or more", n, len(base)))
	}
	if n := fails(fix); n > 0 {
		why = append(why, fmt.Sprintf("%d of %d fix replays fail, and each one must pass", n, len(fix)))
	}
	return why
}

// fails counts the scores of 0 on a bar of pass or fail.
func fails(s []float64) int {
	n := 0
	for _, x := range s {
		if x == 0 {
			n++
		}
	}
	return n
}

// worst is the lowest score when higher is better, else the highest.
func worst(s []float64, higher bool) float64 {
	w := math.Inf(1)
	if !higher {
		w = math.Inf(-1)
	}
	for _, x := range s {
		if higher && x < w || !higher && x > w {
			w = x
		}
	}
	return w
}

// bar checks verdict.json against the session logs and the trees of the
// base and the head in repo (D-1200). The eval session runs it before
// it pushes, and the loop script runs it again on the result ready.
func bar(verdictPath, logs, repo, base, head string, out io.Writer) error {
	raw, err := os.ReadFile(verdictPath)
	if err != nil {
		return fmt.Errorf("bar: %w", err)
	}
	var v verdict
	if err := json.Unmarshal(raw, &v); err != nil {
		return fmt.Errorf("bar: verdict.json does not read: %w", err)
	}
	g := livespend.NewLedger()
	if err := g.AddLogs(logs); err != nil {
		return err
	}
	baseTree, headTree := treeAt(repo, base), treeAt(repo, head)
	if why := barFaults(v, g.LogRuns(), baseTree, headTree); len(why) > 0 {
		return barFailed(why)
	}
	_, err = fmt.Fprintf(out, "bar holds: %s\n", v.Bar)
	return err
}

// treeAt is the tree of the Go module at rev, or empty.
func treeAt(repo, rev string) string {
	b, err := exec.Command("git", "-C", repo, "rev-parse", "--verify", "--quiet", rev+":"+moduleDir).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}
