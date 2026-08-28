package llm

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"testing"
	"time"
)

// blocking is a Provider that waits for ctx to end, then returns ctx.Err()
// wrapped as a transient error. It models an attempt cut off mid-flight.
type blocking struct {
	mu    sync.Mutex
	calls int
}

func (b *blocking) Name() string { return FakeName }

func (b *blocking) Complete(ctx context.Context, call Call) (Response, error) {
	b.mu.Lock()
	b.calls++
	b.mu.Unlock()
	<-ctx.Done()
	return Response{}, newErr(ClassTransient, FakeName, call.Model, 0, ctx.Err())
}

func TestCompleteCallerCancelDuringCall(t *testing.T) {
	b := &blocking{}
	c := newTestClient(t, b)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()
	_, err := c.Complete(ctx, RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassBudget || !errors.Is(err, context.Canceled) {
		t.Errorf("class = %v, err = %v", ClassOf(err), err)
	}
	if b.calls != 1 {
		t.Errorf("calls = %d, want 1", b.calls)
	}
}

func TestCompleteDeadlineDuringCall(t *testing.T) {
	b := &blocking{}
	c := newTestClient(t, b, WithBudget(Budget{Deadline: 20 * time.Millisecond}))
	_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassBudget || !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("class = %v, err = %v", ClassOf(err), err)
	}
	if b.calls != 1 {
		t.Errorf("calls = %d, want 1", b.calls)
	}
}

func TestCompleteSleeperError(t *testing.T) {
	tr := newErr(ClassTransient, FakeName, "m", 503, errors.New("down"))
	sc := NewScript(Step{Err: tr}, Step{Output: json.RawMessage(`{"format":"x"}`)})
	stop := errors.New("sleep cut short")
	c := newTestClient(t, sc, WithSleeper(func(context.Context, time.Duration) error { return stop }))
	_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassBudget || !errors.Is(err, stop) {
		t.Errorf("class = %v, err = %v", ClassOf(err), err)
	}
	if len(sc.Calls) != 1 {
		t.Errorf("calls = %d, want 1", len(sc.Calls))
	}
}

func TestCompleteEscalatedCapSurvivesTransient(t *testing.T) {
	trunc := newErr(ClassTruncation, FakeName, "m", 0, errors.New("cap"))
	tr := newErr(ClassTransient, FakeName, "m", 429, errors.New("rate"))
	sc := NewScript(Step{Err: trunc}, Step{Err: tr}, Step{Output: json.RawMessage(`{"format":"x"}`)})
	c := newTestClient(t, sc)
	res, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attempts != 3 {
		t.Errorf("attempts = %d", res.Attempts)
	}
	caps := []int{sc.Calls[0].MaxOutputTokens, sc.Calls[1].MaxOutputTokens, sc.Calls[2].MaxOutputTokens}
	if caps[0] != 1024 || caps[1] != 8192 || caps[2] != 8192 {
		t.Errorf("caps = %v, want 1024, 8192, 8192", caps)
	}
}

func TestCompleteTruncationAtMaxCapNoRetry(t *testing.T) {
	trunc := newErr(ClassTruncation, FakeName, "m", 0, errors.New("cap"))
	sc := NewScript(Step{Err: trunc}, Step{Output: json.RawMessage(`{"format":"x"}`)})
	cfg := testConfig()
	s := cfg.Roles[RoleClassify]
	s.MaxOutputTokens = escalationMax
	cfg.Roles[RoleClassify] = s
	c, err := New(cfg, []Provider{sc}, WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassTruncation || len(sc.Calls) != 1 {
		t.Errorf("class = %v, calls = %d; want truncation after 1 call", ClassOf(err), len(sc.Calls))
	}
}

func TestCompleteConcurrent(t *testing.T) {
	// Every call gets its own step. The Script serializes its own state.
	const n = 16
	steps := make([]Step, n)
	for i := range steps {
		steps[i] = Step{Output: json.RawMessage(`{"format":"x"}`), Usage: &Usage{InputTokens: 1, OutputTokens: 1}}
	}
	c := newTestClient(t, NewScript(steps...))
	acc := NewAccumulator(nil)
	var wg sync.WaitGroup
	errs := make(chan error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, acc)
			errs <- err
			_ = c.Config()
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Error(err)
		}
	}
	if rep := acc.Report(); rep.Calls != n || rep.Tokens.InputTokens != n {
		t.Errorf("report = %+v", rep)
	}
}

func TestReportIsolation(t *testing.T) {
	acc := NewAccumulator(nil)
	acc.Record(RoleAsk, "m", &Usage{InputTokens: 1}, time.Millisecond)
	rep := acc.Report()
	acc.Record(RoleAsk, "m", &Usage{InputTokens: 5}, time.Millisecond)
	acc.Record(RoleJudge, "m", &Usage{InputTokens: 5}, time.Millisecond)
	if rep.Calls != 1 || rep.Tokens.InputTokens != 1 || rep.LatencyMS != 1 {
		t.Errorf("report changed after a later Record: %+v", rep)
	}
	if len(rep.ByRole) != 1 || rep.ByRole[RoleAsk].Tokens.InputTokens != 1 {
		t.Errorf("by-role slice changed after a later Record: %+v", rep.ByRole[RoleAsk])
	}
	rep.Tokens.InputTokens = 99
	if acc.Report().Tokens.InputTokens != 11 {
		t.Error("a write to the report reached the accumulator")
	}
}

func TestBackoffCapAndJitter(t *testing.T) {
	c := newTestClient(t, NewScript(), WithBudget(Budget{BaseDelay: 10 * time.Second}))
	for i, want := range []time.Duration{10 * time.Second, 20 * time.Second, 30 * time.Second, 30 * time.Second} {
		if got := c.backoff(i + 1); got != want {
			t.Errorf("backoff(%d) = %v, want %v", i+1, got, want)
		}
	}
	c = newTestClient(t, NewScript(), WithBudget(Budget{BaseDelay: 10 * time.Second}), WithJitterSeed(7))
	same := 0
	// From retry 3 on, the base is at the 30 s cap.
	for i := 3; i <= 42; i++ {
		got := c.backoff(i)
		if got < 24*time.Second || got > 36*time.Second {
			t.Errorf("backoff(%d) = %v, outside 30s +/- 20%%", i, got)
		}
		if got == 30*time.Second {
			same++
		}
	}
	if same == 40 {
		t.Error("seeded jitter never moved the delay")
	}
	// A fixed seed gives a fixed sequence.
	a := newTestClient(t, NewScript(), WithJitterSeed(3))
	b := newTestClient(t, NewScript(), WithJitterSeed(3))
	if a.backoff(1) != b.backoff(1) || a.backoff(2) != b.backoff(2) {
		t.Error("same seed gave different delays")
	}
}

func TestNewFillsZeroBudget(t *testing.T) {
	c, err := New(testConfig(), []Provider{NewScript()}, WithBudget(Budget{MaxAttempts: 2}))
	if err != nil {
		t.Fatal(err)
	}
	if c.budget.MaxAttempts != 2 || c.budget.Deadline != DefaultBudget.Deadline || c.budget.BaseDelay != DefaultBudget.BaseDelay {
		t.Errorf("budget = %+v", c.budget)
	}
	c, err = New(testConfig(), []Provider{NewScript()}, WithBudget(Budget{}))
	if err != nil {
		t.Fatal(err)
	}
	if c.budget != DefaultBudget {
		t.Errorf("budget = %+v, want %+v", c.budget, DefaultBudget)
	}
	// A zero Budget must not fail at once.
	sc := NewScript(Step{Output: json.RawMessage(`{"format":"x"}`)})
	c, _ = New(testConfig(), []Provider{sc}, WithBudget(Budget{}))
	if _, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil); err != nil {
		t.Error(err)
	}
}

func TestConfigIsACopy(t *testing.T) {
	c := newTestClient(t, NewScript())
	cp := c.Config()
	s := cp.Roles[RoleClassify]
	s.Model = "changed"
	cp.Roles[RoleClassify] = s
	delete(cp.Roles, RoleAsk)
	if got := c.cfg.Roles[RoleClassify].Model; got == "changed" {
		t.Error("a write to the copy reached the client")
	}
	if _, ok := c.cfg.Roles[RoleAsk]; !ok {
		t.Error("a delete on the copy reached the client")
	}
}

func TestCompleteNilProvider(t *testing.T) {
	c := &Client{cfg: testConfig(), providers: map[string]Provider{}, budget: DefaultBudget, sleep: realSleep}
	_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if ClassOf(err) != ClassTerminal {
		t.Errorf("class = %v, err = %v", ClassOf(err), err)
	}
	if _, err := New(testConfig(), []Provider{nil}); err == nil {
		t.Error("New must refuse a nil provider")
	}
}

func TestCompileSchemaRules(t *testing.T) {
	bad := []struct{ name, schema string }{
		{"not an object", `[1]`},
		{"root not object type", `{"type":"string"}`},
		{"root has no type", `{"properties":{}}`},
		{"file ref", `{"type":"object","properties":{"a":{"$ref":"file:///etc/passwd"}}}`},
		{"http ref", `{"type":"object","properties":{"a":{"$ref":"http://example.com/s.json"}}}`},
	}
	for _, tt := range bad {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestClient(t, NewScript(Step{Output: json.RawMessage(`{}`)}))
			_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(tt.schema)}, nil)
			if ClassOf(err) != ClassTerminal {
				t.Errorf("class = %v, err = %v", ClassOf(err), err)
			}
		})
	}
	// format keywords are asserted.
	schema := `{"type":"object","properties":{"when":{"type":"string","format":"date"}},"required":["when"],"additionalProperties":false}`
	// Two misses, because a schema miss retries once (L-7).
	notDate := Step{Output: json.RawMessage(`{"when":"not a date"}`)}
	c := newTestClient(t, NewScript(notDate, notDate))
	_, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(schema)}, nil)
	if ClassOf(err) != ClassSchema {
		t.Errorf("format: class = %v, err = %v", ClassOf(err), err)
	}
	c = newTestClient(t, NewScript(Step{Output: json.RawMessage(`{"when":"2026-08-24"}`)}))
	if _, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(schema)}, nil); err != nil {
		t.Errorf("valid date: %v", err)
	}
}
