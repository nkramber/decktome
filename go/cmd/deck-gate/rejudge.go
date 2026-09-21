package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
)

// The rejudge lane reads the decks of a whole gate run back and runs the
// two judges over them, and never the build (D-789). A change to the
// input of a judge alone then costs the judge calls and no build. The
// decks stay fixed, so a grade that moves reads the judge alone.

var (
	commanderLine = regexp.MustCompile(`^Commander: (.+)\.$`)
	summaryLine   = regexp.MustCompile(`^\*\*Summary:\*\* (.*)$`)
)

// judgeMetrics are the rows the two judges write. The lane copies every
// other row of the source run, because the decks did not change.
var judgeMetrics = map[string]bool{
	"false_rules": true, "rules_claims": true, "judge_error": true,
	"plan_score": true, "plan_reasons_empty": true, "plan_judge_error": true,
}

func isJudgeMetric(m string) bool {
	return judgeMetrics[m] || strings.HasPrefix(m, "plan_")
}

// storedDeck is one deck as a gate document holds it.
type storedDeck struct {
	commander string
	summary   string
	cards     []*mtgv1.DeckCard
}

// readStored reads the commander line, the summary, and the deck list of
// every prompt of a gate document, by prompt id. A name the index does
// not know is an error, so the judge never reads a deck with a hole.
func readStored(r io.Reader, idx *cards.Index) (map[int]*storedDeck, error) {
	out := map[int]*storedDeck{}
	var cur *storedDeck
	id := 0
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if m := promptHead.FindStringSubmatch(line); m != nil {
			id, _ = strconv.Atoi(m[1])
			cur = &storedDeck{}
			out[id] = cur
			continue
		}
		if cur == nil {
			continue
		}
		if m := commanderLine.FindStringSubmatch(line); m != nil {
			cur.commander = m[1]
			continue
		}
		if m := summaryLine.FindStringSubmatch(line); m != nil {
			cur.summary = m[1]
			continue
		}
		m := deckLine.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		n, err := strconv.Atoi(m[1])
		if err != nil {
			continue
		}
		name := strings.TrimSpace(m[2])
		c, ok := idx.ByName(name)
		if !ok {
			return nil, fmt.Errorf("deck %d: no card named %q", id, name)
		}
		role := mtgv1.CardRole(mtgv1.CardRole_value["CARD_ROLE_"+strings.ToUpper(m[3])])
		cur.cards = append(cur.cards, &mtgv1.DeckCard{
			OracleId: c.GetOracleId(), Name: c.GetName(),
			Count: int32(n), //nolint:gosec // a deck list holds small counts
			Role:  role,
		})
	}
	return out, sc.Err()
}

// commanderIDs resolves a commander line. Two partners share the line as
// "A, B", and a name can hold a comma of its own, so the whole line is
// tried first and then each split point.
func commanderIDs(line string, idx *cards.Index) ([]string, bool) {
	if line == "" {
		return nil, true
	}
	if c, ok := idx.ByName(line); ok {
		return []string{c.GetOracleId()}, true
	}
	parts := strings.Split(line, ", ")
	for i := 1; i < len(parts); i++ {
		a, okA := idx.ByName(strings.Join(parts[:i], ", "))
		b, okB := idx.ByName(strings.Join(parts[i:], ", "))
		if okA && okB {
			return []string{a.GetOracleId(), b.GetOracleId()}, true
		}
	}
	return nil, false
}

// sourceRun names the run file of a gate document: the same name under
// the eval folder beside it.
func sourceRun(doc string) string {
	base := strings.TrimSuffix(filepath.Base(doc), filepath.Ext(doc))
	return filepath.Join(filepath.Dir(doc), "eval", base+".jsonl")
}

// rejudgeDecks turns the stored decks into results for the judges, in
// prompt order. A whole run needs every deck, so a missing one is an
// error.
func rejudgeDecks(prompts []prompt, stored map[int]*storedDeck, idx *cards.Index) ([]result, error) {
	var out []result
	for _, p := range prompts {
		st, ok := stored[p.ID]
		if !ok || len(st.cards) == 0 {
			return nil, fmt.Errorf("prompt %d: the document holds no deck list", p.ID)
		}
		ids, ok := commanderIDs(st.commander, idx)
		if !ok {
			return nil, fmt.Errorf("prompt %d: no commander named %q", p.ID, st.commander)
		}
		setCodes, err := resolveSets(idx, p.Sets)
		if err != nil {
			return nil, fmt.Errorf("prompt %d: %w", p.ID, err)
		}
		out = append(out, result{
			prompt: p,
			deck: &mtgv1.Deck{
				Format:             &mtgv1.Format{Id: gatekit.FormatID(p.Format)},
				Summary:            st.summary,
				CommanderOracleIds: ids,
				Cards:              st.cards,
			},
			setCodes: setCodes,
		})
	}
	return out, nil
}

// copyBuildRows copies every row of the source run that no judge wrote,
// and reads the build bars of the verdict from them: each deck built,
// with no block, no invented name, and no missed case assertion.
func copyBuildRows(run, src *evalrun.Run) (buildPass bool) {
	buildPass = true
	for _, row := range src.Rows {
		if isJudgeMetric(row.Metric) {
			continue
		}
		run.Rows = append(run.Rows, row)
		if row.Kind != evalrun.KindGate {
			continue
		}
		switch row.Metric {
		case "built":
			if row.Value != 1 {
				buildPass = false
			}
		case "blocks", "invented_names", "case_assertions":
			if row.Value != 0 {
				buildPass = false
			}
		}
	}
	return buildPass
}

// keptRows reads the judge rows of an earlier rejudge of the same run,
// by item. An item keeps its rows when both judges answered it. An item
// with a judge error is judged again, so a provider fault costs the
// failed calls alone and never the whole lane (D-789).
func keptRows(path string, src *evalrun.Run) (map[string][]evalrun.Row, error) {
	if path == "" {
		return nil, nil
	}
	kept, err := evalrun.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("-keep: %w", err)
	}
	if kept.Header.Versions["rejudge_of"] != src.Header.RunID {
		return nil, fmt.Errorf("-keep: %s is not a rejudge of %s", kept.Header.RunID, src.Header.RunID)
	}
	for _, k := range []struct {
		name string
		want int
	}{{"plan_rubric", generate.PlanRubricVersion}, {"summary_judge", generate.SummaryJudgeVersion}} {
		if kept.Header.Prompts[k.name] != k.want {
			return nil, fmt.Errorf("-keep: %s read %s %d, and this lane reads %d", kept.Header.RunID, k.name, kept.Header.Prompts[k.name], k.want)
		}
	}
	rows := map[string][]evalrun.Row{}
	failed := map[string]bool{}
	answered := map[string]bool{}
	for _, row := range kept.Rows {
		if !isJudgeMetric(row.Metric) {
			continue
		}
		switch row.Metric {
		case "judge_error", "plan_judge_error":
			failed[row.Item] = true
		case "false_rules":
			answered[row.Item] = true
		}
		rows[row.Item] = append(rows[row.Item], row)
	}
	out := map[string][]evalrun.Row{}
	for item, rs := range rows {
		if answered[item] && !failed[item] {
			out[item] = rs
		}
	}
	return out, nil
}

func runRejudge(doc, runOut, keep string, prompts []prompt) error {
	if err := gatekit.SpendGuard("DECK_GATE"); err != nil {
		return err
	}
	if runOut == "" {
		return errors.New("-rejudge needs -run-out: the lane writes a run of the decks suite")
	}
	if err := gatekit.RefuseExisting(runOut); err != nil {
		return err
	}
	src, err := evalrun.ReadFile(sourceRun(doc))
	if err != nil {
		return fmt.Errorf("rejudge: the run of %s: %w", doc, err)
	}
	if src.Header.Suite != "decks" || src.Header.Partial() {
		return fmt.Errorf("rejudge: %s is not a whole run of the decks suite", src.Header.RunID)
	}
	kept, err := keptRows(keep, src)
	if err != nil {
		return err
	}
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	// The judges read the card facts, so they read the card data the
	// decks were built on.
	if got := idx.AsOf.Format("2006-01-02"); got != src.Header.Snapshot {
		return fmt.Errorf("rejudge: the snapshot reads %s, and %s read %s", got, src.Header.RunID, src.Header.Snapshot)
	}
	f, err := os.Open(doc) // #nosec G304 -- the operator names the file.
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	stored, err := readStored(f, idx)
	if err != nil {
		return fmt.Errorf("rejudge %s: %w", doc, err)
	}
	results, err := rejudgeDecks(prompts, stored, idx)
	if err != nil {
		return fmt.Errorf("rejudge %s: %w", doc, err)
	}
	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)

	run := evalrun.New("decks", evalrun.RunID(runOut))
	for k, v := range src.Header.Roles {
		run.Header.Roles[k] = v
	}
	for k, v := range src.Header.Prompts {
		run.Header.Prompts[k] = v
	}
	for k, v := range src.Header.Versions {
		run.Header.Versions[k] = v
	}
	run.Header.Overrides = append(run.Header.Overrides, src.Header.Overrides...)
	run.SetRoles(client.Config(), llm.RoleJudge)
	run.Header.Prompts["plan_rubric"] = generate.PlanRubricVersion
	run.Header.Prompts["summary_judge"] = generate.SummaryJudgeVersion
	run.Header.Versions["rejudge_of"] = src.Header.RunID
	run.Header.Snapshot = src.Header.Snapshot
	run.Header.Lower = append(run.Header.Lower, src.Header.Lower...)
	run.Header.Note = fmt.Sprintf("a rejudge of %s: the decks and the build rows are copied, and the judge rows are new (D-789)", src.Header.RunID)
	buildPass := copyBuildRows(run, src)
	if keep != "" {
		run.Header.Versions["kept_from"] = evalrun.RunID(keep)
		run.Header.Note += fmt.Sprintf(". The judge rows of %d decks are kept from %s, and its document holds their lines", len(kept), evalrun.RunID(keep))
	}

	start := time.Now()
	for i := range results {
		r := &results[i]
		if _, ok := kept[strconv.Itoa(r.prompt.ID)]; ok {
			continue
		}
		r.judged, r.judgeErr = judge(context.Background(), client, r.prompt.Name, r.deck, idx, acc)
		r.plan, r.planErr = generate.JudgePlan(context.Background(), client, r.prompt.Plan, r.deck, idx, r.setCodes, acc)
		fmt.Fprintf(os.Stderr, "  %2d. %-38s %s\n", r.prompt.ID, r.prompt.Name, rejudgeWord(*r))
	}
	falseRules, judgeErrs := 0, 0
	for _, r := range results {
		item := strconv.Itoa(r.prompt.ID)
		if rows, ok := kept[item]; ok {
			run.Rows = append(run.Rows, rows...)
			for _, row := range rows {
				if row.Metric == "false_rules" && row.Value != 0 {
					falseRules++
				}
			}
			continue
		}
		judgeRows(run, item, r)
		if r.judgeErr != nil {
			judgeErrs++
		} else if r.judged != nil && r.judged.StatesAFalseRule() {
			falseRules++
		}
	}
	pass := buildPass && falseRules == 0 && judgeErrs == 0
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	took := time.Since(start)
	run.Finish(acc.Report(), took, verdict)
	reportRejudge(os.Stdout, doc, src, results, kept, buildPass, falseRules, judgeErrs, acc.Report(), took, run)
	if err := evalrun.WriteFile(runOut, run); err != nil {
		return err
	}
	if !pass {
		return errGateFailed
	}
	return nil
}

func rejudgeWord(r result) string {
	switch {
	case r.judgeErr != nil:
		return "judge error"
	case r.judged != nil && r.judged.StatesAFalseRule():
		return "FALSE rule"
	}
	return "judged"
}

// reportRejudge writes the rejudge document. The judge lines keep the
// form of the source document, so one reader counts both.
func reportRejudge(w io.Writer, doc string, src *evalrun.Run, rs []result, kept map[string][]evalrun.Row, buildPass bool, falseRules, judgeErrs int, rep llm.Report, took time.Duration, run *evalrun.Run) {
	_, _ = fmt.Fprintf(w, "# PR-8 deck gate, a rejudge of %s\n\n", src.Header.RunID)
	_, _ = fmt.Fprintf(w, "Run date: %s. Card snapshot: %s. Decks read from `%s`.\n\n", time.Now().UTC().Format("2006-01-02"), src.Header.Snapshot, filepath.Base(doc))
	build := "every deck of the source passed its build bars"
	if !buildPass {
		build = "a deck of the source failed a build bar"
	}
	_, _ = fmt.Fprintf(w, "Verdict: %s. %s, %d summaries stated a false rule of the game, and %d judge calls failed. The build rows are the rows of the source run, and no deck was built again (D-789).\n\n",
		run.Header.Verdict, build, falseRules, judgeErrs)
	_, _ = fmt.Fprintf(w, "- Calls: %d. Cost: %s. Time: %.0f seconds.\n\n", rep.Calls, gatekit.CostWord(rep), took.Seconds())
	run.Markdown(w)
	_, _ = fmt.Fprintf(w, "\n## Decks\n\n")
	for _, r := range rs {
		_, _ = fmt.Fprintf(w, "### %d. %s\n\n", r.prompt.ID, r.prompt.Name)
		if _, ok := kept[strconv.Itoa(r.prompt.ID)]; ok {
			_, _ = fmt.Fprintf(w, "Kept from %s, and its document holds the judge lines.\n\n", run.Header.Versions["kept_from"])
			continue
		}
		_, _ = fmt.Fprintf(w, "**Summary:** %s\n\n", r.deck.GetSummary())
		if r.judgeErr != nil {
			_, _ = fmt.Fprintf(w, "- JUDGE ERROR: %v\n", r.judgeErr)
		}
		if r.judged != nil {
			for _, c := range r.judged.Claims {
				_, _ = fmt.Fprintf(w, "- JUDGE [%s]: %q. %s\n", c.Truth, c.Text, c.Why)
			}
		}
		if r.planErr != nil {
			_, _ = fmt.Fprintf(w, "- PLAN JUDGE ERROR: %v\n", r.planErr)
		}
		if r.plan != nil {
			for _, f := range generate.PlanFields {
				g := r.plan.Grade(f)
				_, _ = fmt.Fprintf(w, "- PLAN %s=%s: %s\n", f, g.Grade, g.Why)
			}
		}
		_, _ = fmt.Fprintf(w, "\n")
	}
}
