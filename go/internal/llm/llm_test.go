package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"strings"
	"testing"
	"testing/fstest"
	"time"
)

const testSchema = `{"type":"object","properties":{"format":{"type":"string"}},"required":["format"],"additionalProperties":false}`

func testConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	for r, s := range cfg.Roles {
		s.Provider = FakeName
		cfg.Roles[r] = s
	}
	return cfg
}

func newTestClient(t *testing.T, p Provider, opts ...Option) *Client {
	t.Helper()
	var slept []time.Duration
	opts = append([]Option{
		WithSleeper(func(_ context.Context, d time.Duration) error { slept = append(slept, d); return nil }),
		WithBudget(Budget{MaxAttempts: 4, Deadline: time.Minute, BaseDelay: time.Second}),
		WithoutJitter(),
	}, opts...)
	c, err := New(testConfig(), []Provider{p}, opts...)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestLoadConfigDefaults(t *testing.T) {
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Roles[RoleGenerate].Provider == cfg.Roles[RoleJudge].Provider {
		t.Error("judge shares the generator's provider (D-22)")
	}
	if cfg.Roles[RoleGenerate].Provider != OpenAIName {
		t.Errorf("generate provider = %q, want openai (D-21)", cfg.Roles[RoleGenerate].Provider)
	}
	if cfg.Roles[RoleJudge].Provider != AnthropicName {
		t.Errorf("judge provider = %q, want anthropic (D-38)", cfg.Roles[RoleJudge].Provider)
	}
	prices, err := LoadPrices()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range Roles {
		if _, ok := prices.Models[cfg.Roles[r].Model]; !ok {
			t.Errorf("role %s model %q has no price row (M-1)", r, cfg.Roles[r].Model)
		}
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name string
		edit func(*Config)
	}{
		{"missing role", func(c *Config) { delete(c.Roles, RoleAsk) }},
		{"empty model", func(c *Config) { s := c.Roles[RoleAsk]; s.Model = ""; c.Roles[RoleAsk] = s }},
		{"zero cap", func(c *Config) { s := c.Roles[RoleAsk]; s.MaxOutputTokens = 0; c.Roles[RoleAsk] = s }},
		{"judge equals generate", func(c *Config) { s := c.Roles[RoleJudge]; s.Provider = OpenAIName; c.Roles[RoleJudge] = s }},
		{"unknown role", func(c *Config) { c.Roles["oracle"] = c.Roles[RoleAsk] }},
		{"no date", func(c *Config) { c.VerifiedAt = "" }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, _ := LoadConfig()
			tt.edit(cfg)
			if err := cfg.Validate(); err == nil {
				t.Error("want error, got nil")
			}
		})
	}
}

func TestApplyEnv(t *testing.T) {
	cfg, _ := LoadConfig()
	env := map[string]string{
		"LLM_GENERATE_MODEL":  "gpt-5.6-sol",
		"LLM_GENERATE_EFFORT": "high",
		"LLM_JUDGE_PROVIDER":  FakeName,
	}
	if err := cfg.ApplyEnv(func(k string) string { return env[k] }); err != nil {
		t.Fatal(err)
	}
	if got := cfg.Roles[RoleGenerate]; got.Model != "gpt-5.6-sol" || got.Effort != "high" {
		t.Errorf("generate = %+v", got)
	}
	if cfg.Roles[RoleJudge].Provider != FakeName {
		t.Errorf("judge provider = %q", cfg.Roles[RoleJudge].Provider)
	}
	// An override that breaks D-22 is refused.
	bad := map[string]string{"LLM_JUDGE_PROVIDER": OpenAIName}
	cfg2, _ := LoadConfig()
	if err := cfg2.ApplyEnv(func(k string) string { return bad[k] }); err == nil {
		t.Error("want D-22 error, got nil")
	}
}

func TestNewNeedsEveryProvider(t *testing.T) {
	cfg, _ := LoadConfig()
	_, err := New(cfg, []Provider{NewScript()})
	if err == nil {
		t.Error("want missing-provider error, got nil")
	}
}

func TestFakeFixtures(t *testing.T) {
	fsys := fstest.MapFS{
		"classify.json": {Data: []byte(`{"format":"commander"}`)},
		"ask.json":      {Data: []byte(`{not json`)},
	}
	tests := []struct {
		name string
		role Role
		want string
		err  bool
	}{
		{"fixture found", RoleClassify, `{"format":"commander"}`, false},
		{"fixture absent", RoleGenerate, "", true},
		{"fixture invalid", RoleAsk, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := NewFake(fsys).Complete(context.Background(), Call{Role: tt.role})
			if tt.err {
				if err == nil {
					t.Fatal("want error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if string(res.Output) != tt.want {
				t.Errorf("output = %s, want %s", res.Output, tt.want)
			}
		})
	}
}

func TestCompleteHappyPath(t *testing.T) {
	sc := NewScript(Step{Output: json.RawMessage(`{"format":"commander"}`), Usage: &Usage{InputTokens: 10, OutputTokens: 5}})
	c := newTestClient(t, sc)
	acc := NewAccumulator(nil)
	res, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, acc)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attempts != 1 || res.Provider != FakeName || res.Model != "gpt-5.6-luna" {
		t.Errorf("res = %+v", res)
	}
	call := sc.Calls[0]
	if call.Role != RoleClassify || call.Effort != "none" || call.MaxOutputTokens != 1024 || call.SchemaName != "classify_output" {
		t.Errorf("call = %+v", call)
	}
	rep := acc.Report()
	if rep.Calls != 1 || rep.Tokens == nil || rep.Tokens.InputTokens != 10 || rep.CostUSD != nil {
		t.Errorf("report = %+v", rep)
	}
	if res.Latency <= 0 {
		t.Errorf("latency = %v, want > 0", res.Latency)
	}
}

func TestCompleteSchemaMismatch(t *testing.T) {
	sc := NewScript(Step{Output: json.RawMessage(`{"format":7}`)})
	c := newTestClient(t, sc)
	_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassSchema {
		t.Fatalf("class = %v, err = %v", ClassOf(err), err)
	}
	if len(sc.Calls) != 1 {
		t.Errorf("schema errors must not retry, got %d calls", len(sc.Calls))
	}
}

func TestCompleteTruncationEscalatesOnce(t *testing.T) {
	trunc := newErr(ClassTruncation, FakeName, "m", 0, errors.New("cap"))
	sc := NewScript(Step{Err: trunc, Usage: &Usage{OutputTokens: 1024}}, Step{Output: json.RawMessage(`{"format":"modern"}`)})
	c := newTestClient(t, sc)
	acc := NewAccumulator(nil)
	res, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, acc)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attempts != 2 {
		t.Errorf("attempts = %d", res.Attempts)
	}
	if sc.Calls[0].MaxOutputTokens != 1024 || sc.Calls[1].MaxOutputTokens != 8192 {
		t.Errorf("caps = %d, %d; want 1024 then 8192", sc.Calls[0].MaxOutputTokens, sc.Calls[1].MaxOutputTokens)
	}
	// The failed attempt's tokens are real spend and are counted.
	if rep := acc.Report(); rep.Calls != 2 || rep.Tokens.OutputTokens != 1024 {
		t.Errorf("report = %+v", rep)
	}

	// A second truncation is terminal.
	sc = NewScript(Step{Err: trunc}, Step{Err: trunc}, Step{Output: json.RawMessage(`{"format":"modern"}`)})
	c = newTestClient(t, sc)
	_, err = c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassTruncation || len(sc.Calls) != 2 {
		t.Errorf("class = %v, calls = %d", ClassOf(err), len(sc.Calls))
	}
}

func TestEscalate(t *testing.T) {
	tests := []struct{ in, want int }{{100, 8192}, {1024, 8192}, {2048, 16384}, {16384, 65536}, {65536, 65536}, {70000, 70000}}
	for _, tt := range tests {
		if got := escalate(tt.in); got != tt.want {
			t.Errorf("escalate(%d) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestCompleteTransientBackoff(t *testing.T) {
	tr := newErr(ClassTransient, FakeName, "m", 429, errors.New("rate"))
	sc := NewScript(Step{Err: tr}, Step{Err: tr}, Step{Output: json.RawMessage(`{"format":"pauper"}`)})
	var slept []time.Duration
	c := newTestClient(t, sc, WithSleeper(func(_ context.Context, d time.Duration) error { slept = append(slept, d); return nil }))
	res, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attempts != 3 {
		t.Errorf("attempts = %d", res.Attempts)
	}
	if len(slept) != 2 || slept[0] != time.Second || slept[1] != 2*time.Second {
		t.Errorf("backoff = %v, want 1s then 2s", slept)
	}
}

func TestCompleteAttemptBudget(t *testing.T) {
	tr := newErr(ClassTransient, FakeName, "m", 503, errors.New("down"))
	sc := NewScript(Step{Err: tr}, Step{Err: tr}, Step{Err: tr}, Step{Err: tr}, Step{Err: tr})
	c := newTestClient(t, sc)
	_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassBudget {
		t.Errorf("class = %v, err = %v", ClassOf(err), err)
	}
	if len(sc.Calls) != 4 {
		t.Errorf("calls = %d, want 4", len(sc.Calls))
	}
}

func TestCompleteTerminalClasses(t *testing.T) {
	for _, class := range []Class{ClassTerminal, ClassRefusal} {
		sc := NewScript(Step{Err: newErr(class, FakeName, "m", 400, errors.New("no"))}, Step{Output: json.RawMessage(`{"format":"x"}`)})
		c := newTestClient(t, sc)
		_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
		if ClassOf(err) != class || len(sc.Calls) != 1 {
			t.Errorf("%v: class = %v, calls = %d", class, ClassOf(err), len(sc.Calls))
		}
	}
	// A plain error from a provider is terminal too.
	sc := NewScript(Step{Err: errors.New("plain")})
	c := newTestClient(t, sc)
	_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassTerminal || len(sc.Calls) != 1 {
		t.Errorf("plain: class = %v, calls = %d", ClassOf(err), len(sc.Calls))
	}
}

func TestCompleteNoSchema(t *testing.T) {
	c := newTestClient(t, NewScript())
	_, err := c.Complete(context.Background(), RoleClassify, Request{}, nil)
	if ClassOf(err) != ClassTerminal {
		t.Errorf("class = %v", ClassOf(err))
	}
	_, err = c.Complete(context.Background(), "oracle", Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassTerminal {
		t.Errorf("unknown role class = %v", ClassOf(err))
	}
}

func TestAccumulatorNullSemantics(t *testing.T) {
	prices, _ := LoadPrices()
	acc := NewAccumulator(prices)
	if rep := acc.Report(); rep.Calls != 0 || rep.Tokens != nil || rep.CostUSD != nil {
		t.Errorf("empty report = %+v", rep)
	}
	acc.Record(RoleClassify, "gpt-5.6-luna", nil, time.Millisecond)
	if rep := acc.Report(); rep.Calls != 1 || rep.Tokens != nil || rep.CostUSD != nil {
		t.Errorf("uninstrumented report = %+v", rep)
	}
	acc.Record(RoleClassify, "gpt-5.6-luna", &Usage{InputTokens: 1_000_000, CachedInputTokens: 500_000, OutputTokens: 1_000_000}, time.Millisecond)
	rep := acc.Report()
	if rep.Tokens == nil || rep.CostUSD == nil {
		t.Fatalf("report = %+v", rep)
	}
	// 0.5M fresh at $0.20 + 0.5M cached at $0.02 + 1M out at $1.20.
	if want := 0.10 + 0.01 + 1.20; *rep.CostUSD < want-1e-9 || *rep.CostUSD > want+1e-9 {
		t.Errorf("cost = %v, want %v", *rep.CostUSD, want)
	}
	if rep.ByRole[RoleClassify].Calls != 2 || rep.ByRole[RoleClassify].Tokens.OutputTokens != 1_000_000 {
		t.Errorf("by role = %+v", rep.ByRole[RoleClassify])
	}
	// Anthropic cache writes cost 1.25 x input: 1M in = 0.2M fresh at $3
	// + 0.5M read at $0.30 + 0.3M write at $3.75, plus 1M out at $15.
	acc.Record(RoleJudge, "claude-sonnet-5", &Usage{InputTokens: 1_000_000, CachedInputTokens: 500_000, CacheWriteTokens: 300_000, OutputTokens: 1_000_000}, 2*time.Millisecond)
	rep = acc.Report()
	if rep.CostUSD == nil {
		t.Fatalf("report = %+v", rep)
	}
	if want := 1.31 + 0.60 + 0.15 + 1.125 + 15.0; *rep.CostUSD < want-1e-9 || *rep.CostUSD > want+1e-9 {
		t.Errorf("cost = %v, want %v", *rep.CostUSD, want)
	}
	if rep.Tokens.CacheWriteTokens != 300_000 || rep.LatencyMS != 4 {
		t.Errorf("report = %+v", rep)
	}
	acc.Record(RoleJudge, "unknown-model", &Usage{InputTokens: 1}, time.Millisecond)
	if rep := acc.Report(); rep.CostUSD != nil || rep.Tokens.InputTokens != 2_000_001 {
		t.Errorf("unpriced report = %+v", rep)
	}
	var nilAcc *Accumulator
	nilAcc.Record(RoleAsk, "m", nil, 0)
	if rep := nilAcc.Report(); rep.Calls != 0 {
		t.Error("nil accumulator must be inert")
	}
}

func TestNewFromEnvFallsBackToFake(t *testing.T) {
	env := map[string]string{EnvRequireKeys: "0"}
	c, err := NewFromEnv(func(k string) string { return env[k] }, discardLogger())
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range c.Config().Providers() {
		if _, ok := c.providers[name]; !ok {
			t.Errorf("provider %q not wired", name)
		}
	}
	env["LLM_JUDGE_PROVIDER"] = FakeName
	if _, err := NewFromEnv(func(k string) string { return env[k] }, discardLogger()); err != nil {
		t.Errorf("LLM_REQUIRE_KEYS=0 must allow the fake judge: %v", err)
	}
}

func TestNewFromEnvRequiresKeysByDefault(t *testing.T) {
	// Unset means required (D-3): a missing key is fatal.
	_, err := NewFromEnv(func(string) string { return "" }, discardLogger())
	if err == nil {
		t.Error("no keys and LLM_REQUIRE_KEYS unset must fail")
	}
	// A fake role is refused even when every key is present.
	env := map[string]string{
		EnvOpenAIKey:         "sk-test",
		EnvAnthropicKey:      "sk-ant-test",
		"LLM_JUDGE_PROVIDER": FakeName,
	}
	_, err = NewFromEnv(func(k string) string { return env[k] }, discardLogger())
	if err == nil {
		t.Error("LLM_JUDGE_PROVIDER=fake with keys required must fail")
	}
	// LLM_REQUIRE_KEYS=1 is the same as unset.
	env[EnvRequireKeys] = "1"
	if _, err := NewFromEnv(func(k string) string { return env[k] }, discardLogger()); err == nil {
		t.Error("LLM_REQUIRE_KEYS=1 with a fake role must fail")
	}
	// With both keys and no fake role, the real adapters are wired.
	delete(env, "LLM_JUDGE_PROVIDER")
	c, err := NewFromEnv(func(k string) string { return env[k] }, discardLogger())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := c.providers[OpenAIName].(*OpenAI); !ok {
		t.Errorf("openai provider = %T", c.providers[OpenAIName])
	}
	if _, ok := c.providers[AnthropicName].(*Anthropic); !ok {
		t.Errorf("anthropic provider = %T", c.providers[AnthropicName])
	}
}

func TestNewFromEnvWarnsOnUnpricedOverride(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	env := map[string]string{
		EnvRequireKeys:       "0",
		"LLM_CLASSIFY_MODEL": "gpt-99-unknown",
		"LLM_GENERATE_MODEL": "gpt-5.6-sol",
	}
	if _, err := NewFromEnv(func(k string) string { return env[k] }, log); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "no price row") || !strings.Contains(out, "gpt-99-unknown") {
		t.Errorf("want a price-row warning for gpt-99-unknown, got:\n%s", out)
	}
	if strings.Contains(out, "model=gpt-5.6-sol level=WARN") || strings.Count(out, "no price row") != 1 {
		t.Errorf("a priced override must not warn, got:\n%s", out)
	}
}

func TestErrorString(t *testing.T) {
	e := newErr(ClassTransient, "openai", "m", 429, errors.New("slow down"))
	if got := e.Error(); got != "llm openai/m: transient (http 429): slow down" {
		t.Errorf("got %q", got)
	}
	if got := Class(99).String(); got != "class(99)" {
		t.Errorf("got %q", got)
	}
}
