package llm

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	anthropicopt "github.com/anthropics/anthropic-sdk-go/option"
	openaiopt "github.com/openai/openai-go/v3/option"
)

func TestAttemptTimeout(t *testing.T) {
	tests := []struct {
		name string
		cap  int
		left time.Duration
		want time.Duration
	}{
		{"small cap", 1024, 3 * time.Minute, 121 * time.Second},
		{"escalated cap", 65536, 3 * time.Minute, 180 * time.Second},
		{"escalated cap with room", 65536, 10 * time.Minute, 185 * time.Second},
		{"little budget left", 1024, 5 * time.Second, 5 * time.Second},
		{"no budget left", 1024, -time.Second, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := attemptTimeout(attemptBase, tt.cap, tt.left); got != tt.want {
				t.Errorf("attemptTimeout(%d, %v) = %v, want %v", tt.cap, tt.left, got, tt.want)
			}
		})
	}
}

// slowOnce is a Provider whose first attempt waits for its context.
type slowOnce struct{ calls int }

func (s *slowOnce) Name() string { return FakeName }

func (s *slowOnce) Complete(ctx context.Context, call Call) (Response, error) {
	s.calls++
	if s.calls == 1 {
		<-ctx.Done()
		return Response{}, newErr(ClassBudget, FakeName, call.Model, 0, ctx.Err())
	}
	return Response{Output: json.RawMessage(`{"format":"x"}`), Model: call.Model}, nil
}

// TestAttemptTimeoutIsTransient proves an attempt the L-4 window ended
// retries while the budget still has time, whatever class the adapter
// gave the cut-off attempt.
func TestAttemptTimeoutIsTransient(t *testing.T) {
	p := &slowOnce{}
	c := newTestClient(t, p)
	c.attemptBase = 20 * time.Millisecond
	res, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.Attempts != 2 || p.calls != 2 {
		t.Errorf("attempts = %d, calls = %d, want 2 and 2", res.Attempts, p.calls)
	}
}

func TestRetryAfterHeader(t *testing.T) {
	now := time.Date(2026, 8, 28, 12, 0, 0, 0, time.UTC)
	tests := []struct {
		name   string
		header string
		want   time.Duration
	}{
		{"absent", "", 0},
		{"seconds", "7", 7 * time.Second},
		{"fraction", "0.5", 500 * time.Millisecond},
		{"zero", "0", 0},
		{"negative", "-3", 0},
		{"http date", now.Add(90 * time.Second).Format(http.TimeFormat), 90 * time.Second},
		{"past date", now.Add(-time.Minute).Format(http.TimeFormat), 0},
		{"garbage", "soon", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{Header: http.Header{}}
			if tt.header != "" {
				resp.Header.Set("Retry-After", tt.header)
			}
			if got := retryAfter(resp, now); got != tt.want {
				t.Errorf("retryAfter(%q) = %v, want %v", tt.header, got, tt.want)
			}
		})
	}
	if got := retryAfter(nil, now); got != 0 {
		t.Errorf("nil response = %v", got)
	}
}

// TestClientHonorsRetryAfter is L-6: the delay is max(schedule, hint),
// capped at the budget that is left.
func TestClientHonorsRetryAfter(t *testing.T) {
	tests := []struct {
		name     string
		hint     time.Duration
		deadline time.Duration
		want     time.Duration
	}{
		{"schedule wins", 200 * time.Millisecond, time.Minute, time.Second},
		{"hint wins", 10 * time.Second, time.Minute, 10 * time.Second},
		{"budget caps the hint", 10 * time.Minute, time.Minute, time.Minute},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newErr(ClassTransient, FakeName, "m", 429, errors.New("rate"))
			e.RetryAfter = tt.hint
			sc := NewScript(Step{Err: e}, Step{Output: json.RawMessage(`{"format":"x"}`)})
			var slept []time.Duration
			c := newTestClient(t, sc,
				WithSleeper(func(_ context.Context, d time.Duration) error { slept = append(slept, d); return nil }),
				WithBudget(Budget{MaxAttempts: 4, Deadline: tt.deadline, BaseDelay: time.Second}))
			if _, err := c.Complete(context.Background(), RoleClassify, Request{Schema: json.RawMessage(testSchema)}, nil); err != nil {
				t.Fatal(err)
			}
			if len(slept) != 1 || slept[0] > tt.want || slept[0] < tt.want-50*time.Millisecond {
				t.Errorf("slept %v, want about %v", slept, tt.want)
			}
		})
	}
}

// TestAdaptersReadRetryAfter proves both adapters carry the header onto
// the transient error.
func TestAdaptersReadRetryAfter(t *testing.T) {
	tests := []struct {
		name    string
		newProv func(url string) Provider
	}{
		{OpenAIName, func(url string) Provider { return NewOpenAI("sk-test", time.Second, openaiopt.WithBaseURL(url)) }},
		{AnthropicName, func(url string) Provider {
			return NewAnthropic("sk-ant-test", time.Second, anthropicopt.WithBaseURL(url))
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Retry-After", "12")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(`{"error":{"type":"rate_limit_error","message":"slow down"}}`))
			}))
			defer srv.Close()
			_, err := tt.newProv(srv.URL).Complete(context.Background(), testCall())
			if ClassOf(err) != ClassTransient {
				t.Fatalf("class = %v, err = %v", ClassOf(err), err)
			}
			if got := retryAfterOf(err); got != 12*time.Second {
				t.Errorf("retry after = %v, want 12s", got)
			}
		})
	}
}

func TestOpenAIStatusMapping(t *testing.T) {
	tests := []struct {
		status string
		want   Class
	}{
		{"queued", ClassTransient},
		{"in_progress", ClassTransient},
		{"cancelled", ClassTerminal},
		{"something_new", ClassTerminal},
	}
	st := &stub{}
	srv := httptest.NewServer(http.HandlerFunc(st.handler))
	defer srv.Close()
	prov := NewOpenAI("sk-test", time.Second, openaiopt.WithBaseURL(srv.URL))
	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			st.set(200, strings.Replace(openaiOK, `"status":"completed"`, `"status":"`+tt.status+`"`, 1))
			_, err := prov.Complete(context.Background(), testCall())
			if ClassOf(err) != tt.want {
				t.Errorf("class = %v, want %v: %v", ClassOf(err), tt.want, err)
			}
		})
	}
}

func TestValidEffort(t *testing.T) {
	tests := []struct {
		provider, effort string
		wantErr          bool
	}{
		{AnthropicName, "", false},
		{AnthropicName, "high", false},
		{AnthropicName, "max", false},
		{AnthropicName, "minimal", true},
		{OpenAIName, "minimal", false},
		{OpenAIName, "none", false},
		{OpenAIName, "extreme", true},
		{FakeName, "anything", false},
	}
	for _, tt := range tests {
		t.Run(tt.provider+"/"+tt.effort, func(t *testing.T) {
			err := validEffort(tt.provider, tt.effort)
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	cfg := testConfig()
	spec := cfg.Roles[RoleClassify]
	spec.Provider, spec.Effort = OpenAIName, "bogus"
	cfg.Roles[RoleClassify] = spec
	if err := cfg.Validate(); err == nil || !strings.Contains(err.Error(), "bogus") {
		t.Errorf("validate accepted a bad effort: %v", err)
	}
}

// TestFakeNamesTheRolesItServes is A-11 and L-17.
func TestFakeNamesTheRolesItServes(t *testing.T) {
	f := NewFake(fstest.MapFS{
		"health.json": {Data: []byte(`{"ok":true}`)},
		"eval.json":   {Data: []byte(`{}`)},
	})
	_, err := f.Complete(context.Background(), Call{Role: RoleGenerate, Model: "m"})
	if err == nil {
		t.Fatal("no error for a missing fixture")
	}
	for _, want := range []string{`role "generate"`, "eval, health"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not name %q", err, want)
		}
	}
	if ClassOf(err) != ClassTerminal {
		t.Errorf("class = %v", ClassOf(err))
	}
}

func TestSchemaCache(t *testing.T) {
	c := newTestClient(t, NewScript(Step{Output: json.RawMessage(`{"format":"x"}`)}, Step{Output: json.RawMessage(`{"format":"y"}`)}))
	raw := json.RawMessage(testSchema)
	for range 2 {
		if _, err := c.Complete(context.Background(), RoleClassify, Request{Schema: raw}, nil); err != nil {
			t.Fatal(err)
		}
	}
	n := 0
	c.schemas.Range(func(_, _ any) bool { n++; return true })
	if n != 1 {
		t.Errorf("cached schemas = %d, want 1", n)
	}
}
