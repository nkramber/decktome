// Command questions-eval scores every question of a gate run with the
// eval role, and writes both a document a person reads and a summary a
// script reads (D-133).
//
// It replaces the hand scoring of the M-5 sheet for the fast lane. The
// owner still scores by hand when the rubric itself is in question: this
// command applies a rubric, and it can not decide one.
//
// CAUTION: this calls a real provider and costs money. Set QUESTIONS_EVAL=1.
// The command refuses an -out or -json file that exists before it makes
// the first call: a paid result is never overwritten (D-65).
//
// Usage:
//
//	QUESTIONS_EVAL=1 go run ./cmd/questions-eval \
//	  -in ../docs/reference/pr7-question-gate-run14.md \
//	  -out ../docs/reference/pr7-question-eval-run14.md \
//	  -json ../.local/tune/run14.json -budget 0.50
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/tune"
)

func main() {
	in := flag.String("in", "", "the gate document to score")
	out := flag.String("out", "", "write the eval document here (default stdout)")
	jsonOut := flag.String("json", "", "write the machine summary here")
	budget := flag.Float64("budget", 0.50, "stop before the run costs more than this, in USD")
	limit := flag.Int("n", 0, "score only the first n conversations (0 scores all)")
	holdout := flag.Int("holdout", 3, "hold every nth conversation back from the fixer (0 holds none)")
	runOut := flag.String("run-out", "", "write the run header and the rows as JSONL here (PR-15)")
	flag.Parse()
	if err := run(*in, *out, *jsonOut, *runOut, *budget, *limit, *holdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(in, out, jsonOut, runOut string, budget float64, limit, holdout int) error {
	if err := gatekit.SpendGuard("QUESTIONS_EVAL"); err != nil {
		return err
	}
	if in == "" {
		return fmt.Errorf("give -in, a gate document")
	}
	if err := gatekit.RefuseExisting(out, jsonOut, runOut); err != nil {
		return err
	}
	gate, err := tune.ReadRun(in)
	if err != nil {
		return err
	}
	// The eval decides nothing on its own, tune-check does, so the run
	// holds information rows alone and no verdict (PR-15).
	rec := evalrun.New("question-eval", evalrun.RunID(runOut))
	rec.Header.Prompts["eval"] = evalVersion
	rec.LowerIsBetter("bad", "bad_ratio", "bad_ratio_tune", "bad_ratio_holdout", "holdout_bad", "partial", "unjudged", "unsure")
	rec.Header.Versions["gate_document"] = gate.Name
	quiet := gatekit.Quiet()
	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	rec.SetRoles(client.Config(), llm.RoleEval)
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	model := "unknown"
	if spec, ok := client.Config().Roles[llm.RoleEval]; ok {
		model = spec.Model
	}
	acc := llm.NewAccumulator(prices)
	started := time.Now()
	// The eval invents card facts, and a false one changes a verdict. The
	// snapshot refutes the claim for nothing (D-149).
	checker, checkNote := newCardChecker()
	corrected := map[string]string{}

	convs := gate.Conversations
	// A cut run is marked as partial, so it can never stand as a baseline
	// or a candidate (T-2).
	stopped := ""
	if limit > 0 && limit < len(convs) {
		convs = convs[:limit]
		stopped = fmt.Sprintf("-n %d cut the run to %d of %d conversations", limit, limit, len(gate.Conversations))
	}
	var verdicts []tune.Verdict
	missed := map[string][]string{}
	unjudged := map[string]string{}
	var errored []string
	// scored counts the conversations the eval was asked about, so the
	// budget-stop message reports a count and not a map size.
	scored := 0
	for _, conv := range convs {
		// A conversation that errored in the gate asked nothing after the
		// error. It is named, and it counts for nobody (T-5).
		if conv.Errored() {
			errored = append(errored, conv.Name)
			continue
		}
		if len(conv.Questions) == 0 {
			continue
		}
		// Every nth conversation is held back. The fixer never reads its
		// failures, so a ratio that falls here is a real gain and not a
		// reworded test set (D-134). The split keys on the conversation
		// number and not on the position, so a partial run holds out the
		// same conversations as a full one (D-181).
		held := tune.HeldOut(conv.Name, holdout)
		// The budget is a hard stop, checked before every call. A run that
		// costs more than -budget allows is worse than a short run.
		spent, err := costOf(acc, model)
		if err != nil {
			return err
		}
		if spent >= budget {
			stopped = fmt.Sprintf("the budget of $%.2f stopped the run after %d conversations", budget, scored)
			break
		}
		vs, miss, err := score(client, acc, gate.Name, conv, checker.facts(conv))
		if err != nil {
			return fmt.Errorf("%s: %w", conv.Name, err)
		}
		scored++
		// One verdict per question asked, or the conversation is unjudged
		// and the document says so (T-9).
		if why := mismatch(conv, vs); why != "" {
			unjudged[conv.Name] = why
			continue
		}
		for j := range vs {
			vs[j].Holdout = held
			// A refusal that rests on a card the eval calls unreal is
			// dropped when the snapshot holds that card (D-149).
			if fixed, name, ok := checker.correct(vs[j]); ok {
				vs[j] = fixed
				corrected[conv.Name+": "+name] = vs[j].Row
			}
		}
		verdicts = append(verdicts, vs...)
		if len(miss) > 0 {
			missed[conv.Name] = miss
		} else {
			missed[conv.Name] = nil
		}
	}
	cost, err := costOf(acc, model)
	if err != nil {
		return err
	}
	sum := tune.Summarize(gate.Name, model, cost, gate.Metrics, verdicts)
	// The per-conversation counts let a partial run be folded into this
	// one later (D-181).
	sum.Conversations = tune.CountConversations(gate)
	sum.Unjudged = sortedKeys(counts(unjudged))
	sum.Errored = errored
	if stopped != "" {
		sum.Partial, sum.StoppedReason = true, stopped
	}

	w := os.Stdout
	if out != "" {
		f, err := os.Create(out) // #nosec G304 -- the caller names the output.
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()
		w = f
	}
	recordRows(rec, sum)
	rec.Finish(acc.Report(), time.Since(started), "")
	write(w, gate, sum, missed, unjudged, acc.Report(), rec)
	writeCardCheck(w, checkNote, corrected)
	if err := evalrun.WriteFile(runOut, rec); err != nil {
		return err
	}
	if jsonOut != "" {
		if err := os.MkdirAll(filepath.Dir(jsonOut), 0o750); err != nil {
			return err
		}
		raw, err := json.MarshalIndent(sum, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(jsonOut, append(raw, '\n'), 0o600); err != nil {
			return err
		}
	}
	if stopped != "" {
		return fmt.Errorf("%s", stopped)
	}
	return nil
}

// evalOut is the eval role's structured answer for one conversation.
type evalOut struct {
	Verdicts []struct {
		Turn          int      `json:"turn"`
		Row           string   `json:"row"`
		Warranted     string   `json:"warranted"`
		Faults        []string `json:"faults"`
		CatalogAction string   `json:"catalog_action"`
		Reason        string   `json:"reason"`
	} `json:"verdicts"`
	Missed []string `json:"missed"`
}

// score judges one conversation in one call. A call per conversation, and
// not a call per question, is what keeps a whole run inside a few cents:
// the instructions cache, and the conversation is read once.
func score(client *llm.Client, acc *llm.Accumulator, run string, conv tune.Conversation,
	facts []cardFact) ([]tune.Verdict, []string, error) {
	type qIn struct {
		Turn    int      `json:"turn"`
		Row     string   `json:"row"`
		Slot    string   `json:"slot"`
		Text    string   `json:"text"`
		Options []string `json:"options,omitempty"`
	}
	// The transcript is interleaved, turn by turn. A flat list of messages
	// beside a flat list of questions makes the reader hold the order in
	// its head, and a question then reads as a duplicate of an answer the
	// user gave a turn later (D-141).
	type turnIn struct {
		Turn int `json:"turn"`
		// User is the message that arrived on this turn.
		User string `json:"user"`
		// Asked are the questions the agent sent in reply to it. Nothing
		// below this point in the transcript was known when they went out.
		Asked []qIn `json:"agent_asked_in_reply"`
	}
	byTurn := map[int][]qIn{}
	for _, q := range conv.Questions {
		byTurn[q.Turn] = append(byTurn[q.Turn],
			qIn{Turn: q.Turn, Row: q.Row, Slot: q.Slot, Text: q.Text, Options: q.Options})
	}
	turns := make([]turnIn, 0, len(conv.Messages))
	for i, msg := range conv.Messages {
		turns = append(turns, turnIn{Turn: i + 1, User: msg, Asked: byTurn[i+1]})
	}
	input, err := json.Marshal(map[string]any{
		"conversation": conv.Name,
		// The session facts the transcript does not show. Without this
		// one, the card-pool question reads as a presumption (D-143).
		"user_has_a_card_collection": conv.Collection,
		// Every card the questions name, as the snapshot holds it. The
		// eval invented facts about four cards of crossover sets, and it
		// called each one unreal or inapplicable (D-152).
		"cards_named_in_questions": facts,
		"transcript":               turns,
	})
	if err != nil {
		return nil, nil, err
	}
	res, err := client.Complete(context.Background(), llm.RoleEval, llm.Request{
		Instructions: evalInstructions,
		Input:        string(input),
		SchemaName:   "question_verdicts",
		Schema:       json.RawMessage(evalSchema),
		// One key per run, so the instruction prefix caches across the
		// whole document.
		CacheKey: run,
	}, acc)
	if err != nil {
		return nil, nil, err
	}
	var got evalOut
	if err := json.Unmarshal(res.Output, &got); err != nil {
		return nil, nil, err
	}
	byTurnRow := map[string]tune.Question{}
	for _, q := range conv.Questions {
		byTurnRow[key(q.Turn, q.Row)] = q
	}
	out := make([]tune.Verdict, 0, len(got.Verdicts))
	for _, v := range got.Verdicts {
		q := byTurnRow[key(v.Turn, v.Row)]
		out = append(out, tune.Verdict{
			Conversation: conv.Name, Turn: v.Turn, Row: v.Row, Slot: q.Slot, Text: q.Text,
			Warranted: word(v.Warranted), Faults: v.Faults,
			CatalogAction: word(v.CatalogAction), Reason: strings.TrimSpace(v.Reason),
		})
	}
	return out, got.Missed, nil
}

func key(turn int, row string) string { return fmt.Sprintf("%d/%s", turn, row) }

// mismatch says why the verdicts do not cover the questions asked, or
// nothing when they do. The eval answers one verdict per question, keyed
// by turn and row. A merged verdict covers two questions and leaves one
// with none, and a verdict on a question never asked keys to nothing. A
// run that counted such a conversation would score a question nobody
// judged (T-9).
func mismatch(conv tune.Conversation, vs []tune.Verdict) string {
	if len(vs) != len(conv.Questions) {
		return fmt.Sprintf("%d verdicts for %d questions", len(vs), len(conv.Questions))
	}
	want := map[string]int{}
	for _, q := range conv.Questions {
		want[key(q.Turn, q.Row)]++
	}
	for _, v := range vs {
		k := key(v.Turn, v.Row)
		if want[k] == 0 {
			return fmt.Sprintf("a verdict on turn %d row %s, which asked no such question", v.Turn, v.Row)
		}
		want[k]--
	}
	return ""
}

// counts turns a set of names into the map sortedKeys reads.
func counts(m map[string]string) map[string]int {
	out := make(map[string]int, len(m))
	for k := range m {
		out[k] = 1
	}
	return out
}

// word normalizes a free-text answer. The schema leaves these fields free
// on purpose: an enum suppressed a field once already (D-92).
func word(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	for _, cut := range []string{" ", "-", ",", "."} {
		if i := strings.Index(s, cut); i > 0 {
			s = s[:i]
		}
	}
	return s
}

// costOf reads the spend so far. A run with calls and no price is an
// error, never $0: the loop's budget would count nothing (T-1).
func costOf(acc *llm.Accumulator, model string) (float64, error) {
	rep := acc.Report()
	if rep.CostUSD != nil {
		return *rep.CostUSD, nil
	}
	if rep.Calls == 0 {
		return 0, nil
	}
	return 0, fmt.Errorf("no price for model %s in prices.json, so the run can not be charged", model)
}

func write(w io.Writer, _ *tune.Run, s tune.Summary, missed map[string][]string, unjudged map[string]string, rep llm.Report, rec *evalrun.Run) {
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	p("# PR-7 question eval\n\n")
	p("Run: `%s`. Eval model: `%s`. Prompt version %d.\n\n", s.Run, s.Model, evalVersion)
	p("Scored %d questions. %d were not warranted, and %d are unsure.\n\n", s.Judged, s.Bad, s.Unsure)
	scored, split := s.Scored()
	p("**Bad-question ratio: %.1f%% on the %s.**\n\n", scored*100, split)
	if s.HoldoutJudged > 0 {
		p("The tune split holds %d questions at %.1f%%, and the holdout holds %d at %.1f%%. ",
			s.TuneJudged, s.TuneRatio*100, s.HoldoutJudged, s.HoldoutRatio*100)
		p("A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).\n\n")
	}
	if s.Partial {
		p("**PARTIAL: %s.** The ratio above reads the conversations that ran, and this run is no baseline (T-2).\n\n", s.StoppedReason)
	}
	if len(s.Errored) > 0 {
		p("## Conversations that errored in the gate\n\n")
		p("These asked nothing after the error. They hold no verdict, and a paired comparison skips them (T-5).\n\n")
		for _, name := range s.Errored {
			p("- %s\n", name)
		}
		p("\n")
	}
	if len(s.Unjudged) > 0 {
		p("## Conversations the eval left unjudged\n\n")
		p("The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).\n\n")
		for _, name := range s.Unjudged {
			p("- %s: %s\n", name, unjudged[name])
		}
		p("\n")
	}

	p("## The counters the ratio can not see\n\n")
	p("A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.\n\n")
	p("| Counter | Value |\n|---|---|\n")
	p("| Questions asked | %d |\n", s.Metrics.Questions)
	p("| Questions that closed a slot | %d |\n", s.Metrics.CatalogFilled)
	p("| Invented by the model | %d |\n", s.Metrics.Invented)
	p("| Catalog-only conversations | %d |\n", s.Metrics.CatalogOnly)
	p("| Premature sessions | %d |\n", s.Metrics.Premature)
	p("| Linter findings | %d |\n\n", s.Metrics.LintFindings)

	if len(s.ByRow) > 0 {
		p("## Where the bad questions came from\n\n")
		p("| Row | Bad questions |\n|---|---|\n")
		for _, row := range s.TopRows(0) {
			p("| `%s` | %d |\n", row, s.ByRow[row])
		}
		p("\n")
	}
	if len(s.ByFault) > 0 {
		p("## Faults\n\n")
		p("| Fault | Count |\n|---|---|\n")
		for _, f := range sortedKeys(s.ByFault) {
			p("| %s | %d |\n", f, s.ByFault[f])
		}
		p("\n")
	}
	if len(s.Actions) > 0 {
		p("## What the catalog needs\n\n")
		p("| Action | Count |\n|---|---|\n")
		for _, a := range sortedKeys(s.Actions) {
			p("| %s | %d |\n", a, s.Actions[a])
		}
		p("\n")
	}

	p("## Every question the eval refused\n\n")
	if s.HoldoutJudged > 0 {
		p("The holdout conversations are left out of this list on purpose. %d of their questions failed, and the fixer may not read which (D-134).\n\n", s.HoldoutBad)
	}
	bad := 0
	for _, v := range s.Verdicts {
		if !v.Bad() || v.Holdout {
			continue
		}
		bad++
		p("### %s, turn %d, row `%s`\n\n", v.Conversation, v.Turn, v.Row)
		p("**Asked:** %s\n\n", v.Text)
		if len(v.Faults) > 0 {
			p("Faults: %s. ", strings.Join(v.Faults, ", "))
		}
		p("Catalog action: %s.\n\n", gatekit.OrNone(v.CatalogAction))
		p("%s\n\n", v.Reason)
	}
	if bad == 0 {
		p("None.\n\n")
	}

	var names []string
	for name, m := range missed {
		if len(m) > 0 {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	if len(names) > 0 {
		p("## Slots the eval says nobody asked about\n\n")
		p("These do not move the ratio. They are the guard against a quiet agent.\n\n")
		for _, name := range names {
			p("- %s: %s\n", name, strings.Join(missed[name], ", "))
		}
		p("\n")
	}

	rec.Markdown(w)
	if rep.Tokens != nil {
		p("- Tokens: %d input (%d cached), %d output.\n", rep.Tokens.InputTokens, rep.Tokens.CachedInputTokens, rep.Tokens.OutputTokens)
	}
	p("- Eval cost: $%.4f.\n", s.CostUSD)
}

// recordRows writes the counters of the eval into the run as
// information rows (PR-15). A conversation row counts its bad questions
// on the tune split alone: the fixer may never read which holdout
// question failed (D-134), and the run file is a committed record.
func recordRows(rec *evalrun.Run, s tune.Summary) {
	byConv := map[string]*[2]int{}
	var order []string
	for _, v := range s.Verdicts {
		if v.Holdout || v.Unsure() {
			continue
		}
		c, ok := byConv[v.Conversation]
		if !ok {
			c = &[2]int{}
			byConv[v.Conversation] = c
			order = append(order, v.Conversation)
		}
		c[0]++
		if v.Bad() {
			c[1]++
		}
	}
	for _, name := range order {
		item := fmt.Sprintf("%d", tune.ConversationNumber(name))
		rec.Info(item, "judged", float64(byConv[name][0]), name)
		rec.Info(item, "bad", float64(byConv[name][1]), "")
	}
	rec.Info("suite", "judged", float64(s.Judged), "")
	rec.Info("suite", "bad", float64(s.Bad), "")
	rec.Info("suite", "unsure", float64(s.Unsure), "")
	rec.Info("suite", "bad_ratio", s.Ratio, "every judged question")
	rec.Info("suite", "bad_ratio_tune", s.TuneRatio, fmt.Sprintf("%d judged", s.TuneJudged))
	rec.Info("suite", "bad_ratio_holdout", s.HoldoutRatio, fmt.Sprintf("%d judged", s.HoldoutJudged))
	rec.Info("suite", "holdout_bad", float64(s.HoldoutBad), "")
	partial := 0.0
	if s.Partial {
		partial = 1
	}
	rec.Info("suite", "partial", partial, s.StoppedReason)
	rec.Info("suite", "unjudged", float64(len(s.Unjudged)), "")
}

func sortedKeys(m map[string]int) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Slice(out, func(i, j int) bool {
		if m[out[i]] != m[out[j]] {
			return m[out[i]] > m[out[j]]
		}
		return out[i] < out[j]
	})
	return out
}
