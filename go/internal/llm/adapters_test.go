package llm

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	anthropicopt "github.com/anthropics/anthropic-sdk-go/option"
	openaiopt "github.com/openai/openai-go/v3/option"
)

// stub is one fake provider endpoint. It records the last request body.
type stub struct {
	mu     sync.Mutex
	body   map[string]any
	path   string
	status int
	reply  string
	sleep  time.Duration
}

func (s *stub) handler(w http.ResponseWriter, r *http.Request) {
	raw, _ := io.ReadAll(r.Body)
	var body map[string]any
	_ = json.Unmarshal(raw, &body)
	s.mu.Lock()
	s.body, s.path = body, r.URL.Path
	status, reply, sleep := s.status, s.reply, s.sleep
	s.mu.Unlock()
	if sleep > 0 {
		select {
		case <-time.After(sleep):
		case <-r.Context().Done():
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = io.WriteString(w, reply)
}

func (s *stub) set(status int, reply string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status, s.reply, s.sleep = status, reply, 0
}

func (s *stub) lastBody() map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.body
}

// dig reads a nested value by key path.
func dig(m map[string]any, keys ...string) any {
	var cur any = m
	for _, k := range keys {
		mm, ok := cur.(map[string]any)
		if !ok {
			return nil
		}
		cur = mm[k]
	}
	return cur
}

func first(v any) map[string]any {
	arr, ok := v.([]any)
	if !ok || len(arr) == 0 {
		return nil
	}
	m, _ := arr[0].(map[string]any)
	return m
}

// adapterCase binds one adapter to its stub replies.
type adapterCase struct {
	name       string
	newProv    func(url string, timeout time.Duration) Provider
	path       string
	ok         string
	noUsage    string
	truncated  string
	refused    string
	apiErr     string
	thinking   string
	checkBody  func(t *testing.T, body map[string]any)
	checkUsage func(t *testing.T, u *Usage)
}

const openaiOK = `{"id":"resp_1","object":"response","status":"completed","model":"gpt-5.6-luna",
 "output":[{"type":"message","id":"msg_1","role":"assistant","status":"completed",
   "content":[{"type":"output_text","text":"{\"format\":\"commander\"}","annotations":[]}]}],
 "usage":{"input_tokens":10,"input_tokens_details":{"cached_tokens":4},"output_tokens":5,
   "output_tokens_details":{"reasoning_tokens":2},"total_tokens":15}}`

const anthropicOK = `{"id":"msg_1","type":"message","role":"assistant","model":"claude-sonnet-5",
 "content":[{"type":"text","text":"{\"format\":\"commander\"}"}],
 "stop_reason":"end_turn","stop_sequence":null,
 "usage":{"input_tokens":10,"cache_read_input_tokens":4,"cache_creation_input_tokens":3,
   "output_tokens":5,"output_tokens_details":{"thinking_tokens":2}}}`

func adapterCases() []adapterCase {
	return []adapterCase{
		{
			name: OpenAIName,
			newProv: func(url string, timeout time.Duration) Provider {
				return NewOpenAI("sk-test", timeout, openaiopt.WithBaseURL(url))
			},
			path:    "/responses",
			ok:      openaiOK,
			noUsage: strings.Replace(openaiOK, `"usage":{`, `"usage":null,"x":{`, 1),
			truncated: strings.Replace(openaiOK, `"status":"completed"`,
				`"status":"incomplete","incomplete_details":{"reason":"max_output_tokens"}`, 1),
			refused: strings.Replace(openaiOK, `{"type":"output_text","text":"{\"format\":\"commander\"}","annotations":[]}`,
				`{"type":"refusal","refusal":"no"}`, 1),
			apiErr: `{"error":{"message":"%s","type":"x","code":null,"param":null}}`,
			checkBody: func(t *testing.T, b map[string]any) {
				if dig(b, "text", "format", "strict") != true || dig(b, "text", "format", "type") != "json_schema" {
					t.Errorf("text.format = %v", dig(b, "text", "format"))
				}
				if dig(b, "store") != false {
					t.Errorf("store = %v, want false", dig(b, "store"))
				}
				if dig(b, "reasoning", "effort") != "low" {
					t.Errorf("reasoning.effort = %v", dig(b, "reasoning", "effort"))
				}
				if dig(b, "prompt_cache_key") != "sess-1" {
					t.Errorf("prompt_cache_key = %v", dig(b, "prompt_cache_key"))
				}
				if dig(b, "max_output_tokens") != float64(1024) {
					t.Errorf("max_output_tokens = %v", dig(b, "max_output_tokens"))
				}
				if dig(b, "instructions") != "sys" {
					t.Errorf("instructions = %v", dig(b, "instructions"))
				}
			},
			checkUsage: func(t *testing.T, u *Usage) {
				want := Usage{InputTokens: 10, CachedInputTokens: 4, OutputTokens: 5, ReasoningTokens: 2}
				if *u != want {
					t.Errorf("usage = %+v, want %+v", *u, want)
				}
			},
		},
		{
			name: AnthropicName,
			newProv: func(url string, timeout time.Duration) Provider {
				return NewAnthropic("sk-ant-test", timeout, anthropicopt.WithBaseURL(url))
			},
			path:      "/v1/messages",
			ok:        anthropicOK,
			noUsage:   strings.Replace(anthropicOK, `"usage":{`, `"usage":null,"x":{`, 1),
			truncated: strings.Replace(anthropicOK, `"stop_reason":"end_turn"`, `"stop_reason":"max_tokens"`, 1),
			refused: strings.Replace(anthropicOK, `"stop_reason":"end_turn"`,
				`"stop_reason":"refusal","stop_details":{"type":"refusal","category":"cyber","explanation":"nope"}`, 1),
			apiErr: `{"type":"error","error":{"type":"x","message":"%s"}}`,
			thinking: strings.Replace(anthropicOK, `"content":[`,
				`"content":[{"type":"thinking","thinking":"let me see","signature":"sig"},`, 1),
			checkBody: func(t *testing.T, b map[string]any) {
				if dig(b, "output_config", "format", "type") != "json_schema" {
					t.Errorf("output_config.format = %v", dig(b, "output_config", "format"))
				}
				if dig(b, "output_config", "effort") != "low" {
					t.Errorf("output_config.effort = %v", dig(b, "output_config", "effort"))
				}
				if sys := first(b["system"]); sys == nil || dig(sys, "cache_control", "type") != "ephemeral" || sys["text"] != "sys" {
					t.Errorf("system = %v", b["system"])
				}
				if dig(b, "max_tokens") != float64(1024) {
					t.Errorf("max_tokens = %v", dig(b, "max_tokens"))
				}
				if _, ok := b["thinking"]; ok {
					t.Errorf("thinking must be omitted, got %v", b["thinking"])
				}
			},
			checkUsage: func(t *testing.T, u *Usage) {
				// input = 10 fresh + 4 read + 3 write.
				want := Usage{InputTokens: 17, CachedInputTokens: 4, CacheWriteTokens: 3, OutputTokens: 5, ReasoningTokens: 2}
				if *u != want {
					t.Errorf("usage = %+v, want %+v", *u, want)
				}
			},
		},
	}
}

func testCall() Call {
	return Call{
		Request: Request{
			Instructions: "sys",
			Input:        "hello",
			SchemaName:   "classify_output",
			Schema:       json.RawMessage(testSchema),
			CacheKey:     "sess-1",
		},
		Role:            RoleClassify,
		Model:           "m",
		Effort:          "low",
		MaxOutputTokens: 1024,
	}
}

func TestAdapters(t *testing.T) {
	for _, tc := range adapterCases() {
		t.Run(tc.name, func(t *testing.T) {
			st := &stub{}
			srv := httptest.NewServer(http.HandlerFunc(st.handler))
			defer srv.Close()
			prov := tc.newProv(srv.URL, 5*time.Second)
			ctx := context.Background()

			t.Run("happy path", func(t *testing.T) {
				st.set(200, tc.ok)
				res, err := prov.Complete(ctx, testCall())
				if err != nil {
					t.Fatal(err)
				}
				if string(res.Output) != `{"format":"commander"}` {
					t.Errorf("output = %s", res.Output)
				}
				if res.Usage == nil {
					t.Fatal("usage is nil")
				}
				tc.checkUsage(t, res.Usage)
				if !strings.HasSuffix(st.path, tc.path) {
					t.Errorf("path = %q, want suffix %q", st.path, tc.path)
				}
				tc.checkBody(t, st.lastBody())
			})

			t.Run("empty instructions omitted", func(t *testing.T) {
				st.set(200, tc.ok)
				call := testCall()
				call.Instructions = ""
				if _, err := prov.Complete(ctx, call); err != nil {
					t.Fatal(err)
				}
				b := st.lastBody()
				if _, ok := b["instructions"]; ok {
					t.Errorf("instructions present: %v", b["instructions"])
				}
				if _, ok := b["system"]; ok {
					t.Errorf("system present: %v", b["system"])
				}
			})

			t.Run("usage null", func(t *testing.T) {
				st.set(200, tc.noUsage)
				res, err := prov.Complete(ctx, testCall())
				if err != nil {
					t.Fatal(err)
				}
				if res.Usage != nil {
					t.Errorf("usage = %+v, want nil", res.Usage)
				}
			})

			t.Run("truncation", func(t *testing.T) {
				st.set(200, tc.truncated)
				res, err := prov.Complete(ctx, testCall())
				if ClassOf(err) != ClassTruncation {
					t.Errorf("class = %v, err = %v", ClassOf(err), err)
				}
				if res.Usage == nil || res.Usage.OutputTokens != 5 {
					t.Errorf("usage on truncation = %+v", res.Usage)
				}
			})

			t.Run("refusal", func(t *testing.T) {
				st.set(200, tc.refused)
				_, err := prov.Complete(ctx, testCall())
				if ClassOf(err) != ClassRefusal {
					t.Errorf("class = %v, err = %v", ClassOf(err), err)
				}
			})

			t.Run("status table", func(t *testing.T) {
				table := []struct {
					status int
					class  Class
				}{
					{429, ClassTransient}, {503, ClassTransient}, {529, ClassTransient},
					{400, ClassTerminal}, {401, ClassTerminal},
				}
				for _, row := range table {
					st.set(row.status, strings.Replace(tc.apiErr, "%s", "boom", 1))
					_, err := prov.Complete(ctx, testCall())
					var e *Error
					if !errors.As(err, &e) || e.Class != row.class || e.Status != row.status {
						t.Errorf("http %d: got %v", row.status, err)
					}
				}
			})

			t.Run("attempt timeout is transient", func(t *testing.T) {
				slow := tc.newProv(srv.URL, 50*time.Millisecond)
				st.set(200, tc.ok)
				st.mu.Lock()
				st.sleep = 500 * time.Millisecond
				st.mu.Unlock()
				_, err := slow.Complete(ctx, testCall())
				if ClassOf(err) != ClassTransient {
					t.Errorf("class = %v, err = %v", ClassOf(err), err)
				}
			})

			t.Run("caller cancel is budget", func(t *testing.T) {
				st.set(200, tc.ok)
				st.mu.Lock()
				st.sleep = 500 * time.Millisecond
				st.mu.Unlock()
				cctx, cancel := context.WithCancel(ctx)
				go func() {
					time.Sleep(20 * time.Millisecond)
					cancel()
				}()
				_, err := prov.Complete(cctx, testCall())
				if ClassOf(err) != ClassBudget {
					t.Errorf("class = %v, err = %v", ClassOf(err), err)
				}
			})

			if tc.thinking != "" {
				t.Run("thinking block ignored", func(t *testing.T) {
					st.set(200, tc.thinking)
					res, err := prov.Complete(ctx, testCall())
					if err != nil {
						t.Fatal(err)
					}
					if string(res.Output) != `{"format":"commander"}` {
						t.Errorf("output = %s", res.Output)
					}
					if res.Usage == nil || res.Usage.ReasoningTokens != 2 || res.Usage.CacheWriteTokens != 3 {
						t.Errorf("usage = %+v", res.Usage)
					}
				})
			}
		})
	}
}

func TestOpenAIFailedStatus(t *testing.T) {
	st := &stub{}
	srv := httptest.NewServer(http.HandlerFunc(st.handler))
	defer srv.Close()
	prov := NewOpenAI("sk-test", time.Second, openaiopt.WithBaseURL(srv.URL))
	st.set(200, strings.Replace(openaiOK, `"status":"completed"`,
		`"status":"failed","error":{"code":"server_error","message":"it broke"}`, 1))
	_, err := prov.Complete(context.Background(), testCall())
	if ClassOf(err) != ClassTerminal || !strings.Contains(err.Error(), "server_error") || !strings.Contains(err.Error(), "it broke") {
		t.Errorf("err = %v", err)
	}
}

func TestAnthropicContextWindowIsTerminal(t *testing.T) {
	st := &stub{}
	srv := httptest.NewServer(http.HandlerFunc(st.handler))
	defer srv.Close()
	prov := NewAnthropic("sk-ant-test", time.Second, anthropicopt.WithBaseURL(srv.URL))
	for _, reason := range []string{"model_context_window_exceeded", "tool_use", "pause_turn", "something_new"} {
		st.set(200, strings.Replace(anthropicOK, `"stop_reason":"end_turn"`, `"stop_reason":"`+reason+`"`, 1))
		res, err := prov.Complete(context.Background(), testCall())
		if ClassOf(err) != ClassTerminal {
			t.Errorf("%s: class = %v, err = %v", reason, ClassOf(err), err)
		}
		if res.Usage == nil {
			t.Errorf("%s: usage dropped", reason)
		}
	}
	// stop_sequence is a normal end.
	st.set(200, strings.Replace(anthropicOK, `"stop_reason":"end_turn"`, `"stop_reason":"stop_sequence"`, 1))
	if _, err := prov.Complete(context.Background(), testCall()); err != nil {
		t.Errorf("stop_sequence: %v", err)
	}
}
