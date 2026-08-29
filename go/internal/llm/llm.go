// Package llm is the one door for language-model calls (D-1, guardrail 3).
//
// A call site names a Role. The Client resolves the role to a provider and a
// model from the frozen config (roles.json), owns the request budget and the
// retry classes, validates the structured output against the caller's JSON
// Schema, and records usage. No call site ever names a model.
//
// The package holds the seam, the fixture Fake, the role layer, the
// OpenAI and Anthropic adapters (D-21, D-22), and usage accounting (M-1).
package llm

import (
	"context"
	"encoding/json"
	"time"
)

// Role names a call site's job. The config maps each role to a provider,
// a model, a reasoning effort, and an output cap.
type Role string

// The roles of the agent (D-133, D-283). Roles lists them all.
const (
	RoleClassify Role = "classify"
	RoleAsk      Role = "ask"
	RoleGenerate Role = "generate"
	RoleRepair   Role = "repair"
	RoleJudge    Role = "judge"
	// RoleEval scores the questions the agent asked, for the automated
	// M-5 lane. It is not RoleJudge: the judge rates a deck (D-4), and
	// D-22 keeps it off the generator's provider. The eval role rates a
	// question and runs on the cost tier (D-133).
	RoleEval Role = "eval"
	// RoleRevise reads a message after a build and returns the revision
	// brief: what to remove, what to keep, the limits, one question when
	// the request is unclear, and what it declines with a reason
	// (D-283, D-284).
	RoleRevise Role = "revise"
)

// Roles lists every role in config order. Config validation requires all.
var Roles = []Role{RoleClassify, RoleAsk, RoleGenerate, RoleRepair, RoleJudge, RoleEval, RoleRevise}

// Request is one structured-output call as the call site writes it.
type Request struct {
	// Instructions is the system prompt. Keep it stable across calls of the
	// same role: both providers cache a stable prefix.
	Instructions string
	// Input is the user-turn content, serialized by the caller.
	Input string
	// SchemaName labels the output schema for the provider (a-z, 0-9, _ , -).
	SchemaName string
	// Schema is the JSON Schema the output must satisfy. Strict rules apply:
	// every object needs "additionalProperties": false and a full "required"
	// list. The Client validates the output against it before it returns.
	Schema json.RawMessage
	// CacheKey is an optional routing hint for provider-side prompt caching.
	// Use one key per session so a session's turns land on the same cache.
	CacheKey string
}

// Call is what a provider adapter receives: the request plus the resolved
// model settings. Only the Client builds a Call.
type Call struct {
	Request
	Role            Role
	Model           string
	Effort          string
	MaxOutputTokens int
}

// Usage is the token count of one provider call, as the provider reported
// it. A nil *Usage means the provider reported nothing. Zero means measured
// zero. These are different facts (M-1).
type Usage struct {
	// InputTokens is the full input count: fresh, cache read, and cache
	// write tokens together.
	InputTokens int64 `json:"input_tokens"`
	// CachedInputTokens is the part of InputTokens read from a cache.
	CachedInputTokens int64 `json:"cached_input_tokens"`
	// CacheWriteTokens is the part of InputTokens written to a cache.
	// Anthropic bills it at 1.25 x input. OpenAI reports none.
	CacheWriteTokens int64 `json:"cache_write_tokens"`
	OutputTokens     int64 `json:"output_tokens"`
	// ReasoningTokens is the part of OutputTokens spent on reasoning.
	ReasoningTokens int64 `json:"reasoning_tokens"`
}

// Add sums o into u.
func (u *Usage) Add(o Usage) {
	u.InputTokens += o.InputTokens
	u.CachedInputTokens += o.CachedInputTokens
	u.CacheWriteTokens += o.CacheWriteTokens
	u.OutputTokens += o.OutputTokens
	u.ReasoningTokens += o.ReasoningTokens
}

// Response carries the validated structured output.
type Response struct {
	// Output is a JSON document that satisfies the request's Schema.
	Output json.RawMessage
	// Provider and Model are the resolved names, for eval rows and logs.
	Provider string
	Model    string
	// Usage is the provider's token report for the final attempt, or nil.
	Usage *Usage
	// Attempts counts provider calls made for this response, retries included.
	Attempts int
	// Latency is the wall-clock time of all attempts.
	Latency time.Duration
}

// Provider executes one Call against one vendor. Adapters return a *Error
// with a retry Class for every failure they can classify.
type Provider interface {
	// Name is the provider key used in roles.json ("openai", "anthropic", "fake").
	Name() string
	Complete(ctx context.Context, call Call) (Response, error)
}

// Sleeper waits, or returns early when ctx ends. Tests replace it.
type Sleeper func(ctx context.Context, d time.Duration) error

func realSleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
