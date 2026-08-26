// Command questions-eval scores every question of a gate run with the
// eval role, and writes both a document a person reads and a summary a
// script reads (D-133).
//
// It replaces the hand scoring of the M-5 sheet for the fast lane. The
// owner still scores by hand when the rubric itself is in question: this
// command applies a rubric, and it can not decide one.
//
// CAUTION: this calls a real provider and costs money. Set QUESTIONS_EVAL=1.
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
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

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
	flag.Parse()
	if err := run(*in, *out, *jsonOut, *budget, *limit, *holdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(in, out, jsonOut string, budget float64, limit, holdout int) error {
	if os.Getenv("QUESTIONS_EVAL") != "1" {
		return fmt.Errorf("this run calls a real provider and costs money: set QUESTIONS_EVAL=1 to allow it")
	}
	if in == "" {
		return fmt.Errorf("give -in, a gate document")
	}
	gate, err := tune.ReadRun(in)
	if err != nil {
		return err
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	client, err := llm.NewFromEnv(os.Getenv, quiet)
	if err != nil {
		return err
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)
	started := time.Now()

	convs := gate.Conversations
	if limit > 0 && limit < len(convs) {
		convs = convs[:limit]
	}
	var verdicts []tune.Verdict
	missed := map[string][]string{}
	stopped := ""
	for i, conv := range convs {
		if len(conv.Questions) == 0 {
			continue
		}
		// Every nth conversation is held back. The fixer never reads its
		// failures, so a ratio that falls here is a real gain and not a
		// reworded test set (D-134).
		held := holdout > 0 && (i+1)%holdout == 0
		// The budget is a hard stop, checked before every call. A run that
		// costs more than the owner allowed is worse than a short run.
		if spent := costOf(acc); spent >= budget {
			stopped = fmt.Sprintf("the budget of $%.2f stopped the run after %d conversations", budget, len(missed))
			break
		}
		vs, miss, err := score(client, acc, gate.Name, conv)
		if err != nil {
			return fmt.Errorf("%s: %w", conv.Name, err)
		}
		for j := range vs {
			vs[j].Holdout = held
		}
		verdicts = append(verdicts, vs...)
		if len(miss) > 0 {
			missed[conv.Name] = miss
		} else {
			missed[conv.Name] = nil
		}
	}
	model := "unknown"
	if spec, ok := client.Config().Roles[llm.RoleEval]; ok {
		model = spec.Model
	}
	sum := tune.Summarize(gate.Name, model, costOf(acc), gate.Metrics, verdicts)

	w := os.Stdout
	if out != "" {
		f, err := os.Create(out) // #nosec G304 -- the caller names the output.
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()
		w = f
	}
	write(w, gate, sum, missed, stopped, acc.Report(), time.Since(started))
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
func score(client *llm.Client, acc *llm.Accumulator, run string, conv tune.Conversation) ([]tune.Verdict, []string, error) {
	type qIn struct {
		Turn    int      `json:"turn"`
		Row     string   `json:"row"`
		Slot    string   `json:"slot"`
		Text    string   `json:"text"`
		Options []string `json:"options,omitempty"`
	}
	// The transcript is interleaved, turn by turn. A flat list of messages
	// beside a flat list of questions makes the reader hold the order in
	// its head, and the cost tier did not. The calibration of 2026-08-26
	// found seven questions marked as duplicates of an answer the user
	// gave one or two turns later (D-141).
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
		"transcript":                 turns,
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

func costOf(acc *llm.Accumulator) float64 {
	if c := acc.Report().CostUSD; c != nil {
		return *c
	}
	return 0
}

func write(w io.Writer, _ *tune.Run, s tune.Summary, missed map[string][]string, stopped string, rep llm.Report, elapsed time.Duration) {
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
	if stopped != "" {
		p("**%s.** The ratio above reads the conversations that ran.\n\n", stopped)
	}

	p("## The counters the ratio can not see\n\n")
	p("A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.\n\n")
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
		p("Catalog action: %s.\n\n", orNone(v.CatalogAction))
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

	p("## Run\n\n")
	p("- Calls: %d. Time: %.1f seconds.\n", rep.Calls, elapsed.Seconds())
	if rep.Tokens != nil {
		p("- Tokens: %d input (%d cached), %d output.\n", rep.Tokens.InputTokens, rep.Tokens.CachedInputTokens, rep.Tokens.OutputTokens)
	}
	p("- Cost: $%.4f.\n", s.CostUSD)
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

func orNone(s string) string {
	if s == "" {
		return "none"
	}
	return s
}
