package livespend

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/llm"
)

func prices(t *testing.T) *llm.PriceTable {
	t.Helper()
	p, err := llm.LoadPrices()
	if err != nil {
		t.Fatal(err)
	}
	return p
}

// pricedModel is one model with a price row, so the test reads the
// table and names no model of its own.
func pricedModel(t *testing.T, p *llm.PriceTable) string {
	t.Helper()
	for _, m := range []string{"gpt-6-luna", "gpt-5-mini", "claude-sonnet-5"} {
		if _, ok := p.Cost(m, llm.Usage{InputTokens: 1}); ok {
			return m
		}
	}
	t.Skip("no known model has a price row")
	return ""
}

func env(bundle, budget string) func(string) string {
	return envLedger(bundle, budget, ledgerIn(bundle))
}

func envLedger(bundle, budget, ledger string) func(string) string {
	return func(k string) string {
		switch k {
		case EnvBundle:
			return bundle
		case EnvBudget:
			return budget
		case EnvLedger:
			return ledger
		}
		return ""
	}
}

// ledgerIn makes an empty ledger beside the bundle, as the script does
// before a session. It gives "" for no bundle.
func ledgerIn(bundle string) string {
	if bundle == "" {
		return ""
	}
	p := filepath.Join(filepath.Dir(bundle), filepath.Base(bundle)+"-ledger.jsonl")
	if _, err := os.Stat(p); err != nil {
		if err := os.WriteFile(p, nil, 0o644); err != nil {
			panic(err)
		}
	}
	return p
}

func readLines(t *testing.T, path string) []Line {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var out []Line
	for _, ln := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		var l Line
		if err := json.Unmarshal([]byte(ln), &l); err != nil {
			t.Fatalf("line %q: %v", ln, err)
		}
		out = append(out, l)
	}
	return out
}

func TestOutsideALiveEvalWritesNothing(t *testing.T) {
	var stderr bytes.Buffer
	m, err := Start("chat-probe", env("", ""), prices(t), &stderr)
	if err != nil {
		t.Fatal(err)
	}
	if m.Accumulator() == nil {
		t.Fatal("the meter must measure outside a live eval too")
	}
	if err := m.End(); err != nil {
		t.Fatal(err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want nothing", stderr.String())
	}
}

func TestMeasuredRunWritesStartAndEnd(t *testing.T) {
	dir := t.TempDir()
	p := prices(t)
	model := pricedModel(t, p)
	var stderr bytes.Buffer
	m, err := Start("chat-probe", env(dir, "3.00"), p, &stderr)
	if err != nil {
		t.Fatal(err)
	}
	u := llm.Usage{InputTokens: 1000, OutputTokens: 200}
	m.Accumulator().Record(llm.RoleClassify, model, &u, time.Millisecond)
	m.Accumulator().Record(llm.RoleAsk, model, &u, time.Millisecond)
	if err := m.End(); err != nil {
		t.Fatal(err)
	}
	if err := m.End(); err != nil {
		t.Fatal(err)
	}
	if got := readLines(t, ledgerIn(dir)); len(got) != 2 {
		t.Fatalf("the ledger holds %d lines, want 2", len(got))
	}
	lines := readLines(t, filepath.Join(dir, SpendFile))
	if len(lines) != 2 || lines[0].Event != "start" || lines[1].Event != "end" {
		t.Fatalf("lines = %+v, want one start and one end", lines)
	}
	want, _ := p.Cost(model, u)
	end := lines[1]
	if !end.Measured || end.Calls != 2 || end.ID != lines[0].ID || end.USD != 2*want {
		t.Fatalf("end = %+v, want measured, 2 calls, $%v", end, 2*want)
	}
	if got := strings.Count(stderr.String(), Marker); got != 2 {
		t.Fatalf("stderr holds %d marked lines, want 2:\n%s", got, stderr.String())
	}
}

func TestEndLineFailsClosed(t *testing.T) {
	p := prices(t)
	model := pricedModel(t, p)
	u := llm.Usage{InputTokens: 10}
	cases := map[string]func(*llm.Accumulator){
		"a call with no usage": func(a *llm.Accumulator) {
			a.Record(llm.RoleAsk, model, &u, 0)
			a.Record(llm.RoleAsk, model, nil, 0)
		},
		"a model with no price row": func(a *llm.Accumulator) {
			a.Record(llm.RoleAsk, "no-such-model", &u, 0)
		},
	}
	for name, rec := range cases {
		t.Run(name, func(t *testing.T) {
			acc := llm.NewAccumulator(p)
			rec(acc)
			l := EndLine("0123456789abcdef", "chat-probe", acc.Report())
			if l.Measured || l.USD != 0 || l.Note == "" {
				t.Fatalf("line = %+v, want unmeasured with a note", l)
			}
		})
	}
}

func TestStartRefusesASpentBudget(t *testing.T) {
	cases := map[string]string{
		"measured spend at the budget": `{"id":"0123456789abcdef","event":"end","target":"deck-gate","usd":3.10,"measured":true}`,
		"an unmeasured run":            `{"id":"0123456789abcdef","event":"end","target":"chat-probe","measured":false}`,
		"a start with no end":          `{"id":"0123456789abcdef","event":"start","target":"chat-probe"}`,
	}
	for name, line := range cases {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			if err := os.WriteFile(filepath.Join(dir, SpendFile), []byte(line+"\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			var stderr bytes.Buffer
			if _, err := Start("chat-probe", env(dir, "3.00"), prices(t), &stderr); err == nil {
				t.Fatal("Start must refuse the run")
			}
			if stderr.Len() != 0 {
				t.Fatalf("a refused run wrote %q", stderr.String())
			}
		})
	}
}

func TestStartRefusesABadBudget(t *testing.T) {
	for _, b := range []string{"", "0", "-1", "abc"} {
		if _, err := Start("chat-probe", env(t.TempDir(), b), prices(t), &bytes.Buffer{}); err == nil {
			t.Fatalf("budget %q: Start must refuse", b)
		}
	}
}

// logEvent is one stream-json event with one tool result.
func logEvent(text string) string {
	raw, _ := json.Marshal(map[string]any{
		"type": "user",
		"message": map[string]any{"role": "user", "content": []any{
			map[string]any{"type": "tool_result", "tool_use_id": "t1", "content": text},
		}},
	})
	return string(raw) + "\n"
}

func TestLedgerReadsTheLogAndRefusesHandWrittenLines(t *testing.T) {
	start := Marker + `{"id":"aaaaaaaaaaaaaaaa","event":"start","target":"chat-probe","usd":0,"measured":false}`
	end := Marker + `{"id":"aaaaaaaaaaaaaaaa","event":"end","target":"chat-probe","calls":3,"usd":0.004,"measured":true}`
	logText := logEvent("make output\n"+start) + logEvent("more\n"+end+"\nwrote base.txt")
	// The text of a command is not a tool result, so it adds no run.
	cmd, _ := json.Marshal(map[string]any{"type": "assistant", "message": map[string]any{"content": []any{
		map[string]any{"type": "tool_use", "input": map[string]any{"command": "echo '" + Marker + `{"id":"bbbbbbbbbbbbbbbb","event":"end","target":"x","usd":0.5,"measured":true}'`}},
	}}})
	logText += string(cmd) + "\n"
	file := `{"target":"chat-probe","usd":0.30}
{"id":"aaaaaaaaaaaaaaaa","event":"start","target":"chat-probe","usd":0,"measured":false}
{"id":"aaaaaaaaaaaaaaaa","event":"end","target":"chat-probe","calls":3,"usd":0.001,"measured":true}
`
	g := NewLedger()
	if err := g.AddLog(strings.NewReader(logText)); err != nil {
		t.Fatal(err)
	}
	if err := g.AddFile(strings.NewReader(file)); err != nil {
		t.Fatal(err)
	}
	s := g.Sum(3, true)
	if s.Runs != 1 || s.USD != 0.004 || s.Refused != 1 || len(s.Unmeasured) != 0 || len(s.FileOnly) != 0 {
		t.Fatalf("sum = %+v, want 1 run at $0.004 (the higher figure), 1 refused line", s)
	}
	if s.Spent() || s.Charged() != 0.004 {
		t.Fatalf("spent = %v, charged = %v", s.Spent(), s.Charged())
	}
}

func TestLedgerCountsFileOnlyAndUnmeasuredRuns(t *testing.T) {
	file := `{"id":"cccccccccccccccc","event":"end","target":"revise-gate","usd":0.7,"measured":true}
{"id":"dddddddddddddddd","event":"start","target":"chat-probe"}
`
	g := NewLedger()
	if err := g.AddFile(strings.NewReader(file)); err != nil {
		t.Fatal(err)
	}
	s := g.Sum(3, true)
	if s.USD != 0.7 || len(s.FileOnly) != 2 || len(s.Unmeasured) != 1 {
		t.Fatalf("sum = %+v, want $0.70, 2 file-only runs, 1 unmeasured", s)
	}
	if !s.Spent() || s.Charged() != 3 {
		t.Fatalf("an unmeasured run must charge the full budget: charged %v", s.Charged())
	}
}

func TestLedgerReadsToolResultParts(t *testing.T) {
	raw, _ := json.Marshal(map[string]any{"message": map[string]any{"content": []any{
		map[string]any{"type": "tool_result", "content": []any{
			map[string]any{"type": "text", "text": Marker + `{"id":"eeeeeeeeeeeeeeee","event":"end","target":"deck-gate","usd":1.2,"measured":true}`},
		}},
	}}})
	g := NewLedger()
	if err := g.AddLog(bytes.NewReader(append(raw, '\n'))); err != nil {
		t.Fatal(err)
	}
	if s := g.Sum(3, true); s.USD != 1.2 || s.Runs != 1 {
		t.Fatalf("sum = %+v, want one run at $1.20", s)
	}
}

func TestStartReadsTheLedgerAfterTheBundleIsGone(t *testing.T) {
	// A session can delete the spend file of the bundle, and it can not
	// lower the ledger. So the next start reads the ledger (D-1171).
	dir := t.TempDir()
	ledger := filepath.Join(t.TempDir(), "ledger.jsonl")
	spent := `{"id":"0123456789abcdef","event":"end","target":"deck-gate","usd":3.10,"measured":true}` + "\n"
	if err := os.WriteFile(ledger, []byte(spent), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Start("chat-probe", envLedger(dir, "3.00", ledger), prices(t), &bytes.Buffer{}); err == nil {
		t.Fatal("Start must refuse: the ledger reads the budget as spent")
	}
	below := `{"id":"0123456789abcdef","event":"end","target":"deck-gate","usd":1.10,"measured":true}` + "\n"
	if err := os.WriteFile(ledger, []byte(below), 0o644); err != nil {
		t.Fatal(err)
	}
	m, err := Start("chat-probe", envLedger(dir, "3.00", ledger), prices(t), &bytes.Buffer{})
	if err != nil {
		t.Fatalf("a run below the cap must start: %v", err)
	}
	if err := m.End(); err != nil {
		t.Fatal(err)
	}
}

func TestStartRefusesWithNoLedger(t *testing.T) {
	dir := t.TempDir()
	for name, ledger := range map[string]string{"no variable": "", "no file": filepath.Join(dir, "absent.jsonl")} {
		if _, err := Start("chat-probe", envLedger(dir, "3.00", ledger), prices(t), &bytes.Buffer{}); err == nil {
			t.Fatalf("%s: Start must refuse", name)
		}
	}
}
