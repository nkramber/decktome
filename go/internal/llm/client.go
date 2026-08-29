package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Budget bounds one logical call. Every retry class draws from it.
type Budget struct {
	// MaxAttempts caps provider calls, transient retries included.
	MaxAttempts int
	// Deadline caps wall-clock time for all attempts together.
	Deadline time.Duration
	// BaseDelay is the first transient backoff. Each retry doubles it.
	BaseDelay time.Duration
}

// DefaultBudget fills every zero field of the Budget a caller passes.
var DefaultBudget = Budget{MaxAttempts: 4, Deadline: 3 * time.Minute, BaseDelay: 2 * time.Second}

// withDefaults returns b with every zero field set from DefaultBudget.
func (b Budget) withDefaults() Budget {
	if b.MaxAttempts <= 0 {
		b.MaxAttempts = DefaultBudget.MaxAttempts
	}
	if b.Deadline <= 0 {
		b.Deadline = DefaultBudget.Deadline
	}
	if b.BaseDelay <= 0 {
		b.BaseDelay = DefaultBudget.BaseDelay
	}
	return b
}

// Escalation bounds for a truncation retry. The new cap is
// min(escalationMax, max(escalationFactor x cap, escalationFloor)), and
// never more than the provider's own output limit (providerMaxOutput).
// A small call can not become a huge call.
const (
	escalationFactor = 8
	escalationFloor  = 8192
	escalationMax    = 65536
)

// providerMaxOutput is the largest max_tokens value an escalation may
// send to each provider. The values are the smallest "max output" of the
// models roles.json names, read on 2026-08-29 from
// https://platform.claude.com/docs/en/docs/about-claude/models/overview
// (Claude Sonnet 5: 128K, Claude Haiku 4.5: 64K) and from
// https://developers.openai.com/api/docs/models (gpt-5.6-luna, -terra,
// and -sol: 128K). Haiku 4.5 sets the Anthropic floor, so an override to
// it stays legal. A provider not listed here takes escalationMax.
var providerMaxOutput = map[string]int{
	AnthropicName: 65536,
	OpenAIName:    131072,
}

// Attempt timeout rule (L-4). One attempt gets attemptBase plus
// attemptPerKTokens for each 1,000 output tokens of its cap, and never
// more than what is left of the Budget deadline. The cap-based part is
// what lets an escalated retry finish inside its own attempt.
const (
	attemptBase       = 120 * time.Second
	attemptPerKTokens = time.Second
)

// Backoff bounds. A transient delay never exceeds maxBackoff, jitter
// included, unless the provider's Retry-After hint asks for more. Jitter
// moves each delay by up to jitterFraction in either direction, so many
// callers do not retry in step.
const (
	maxBackoff     = 30 * time.Second
	jitterFraction = 0.2
)

// Client is the role layer. Call sites hold one Client and pass a Role.
type Client struct {
	cfg       *Config
	providers map[string]Provider
	budget    Budget
	sleep     Sleeper
	log       *slog.Logger

	// rng feeds the backoff jitter. nil means no jitter. The mutex guards
	// it: rand.Rand is not safe for concurrent use.
	rngMu sync.Mutex
	rng   *rand.Rand

	// attemptBase is the per-attempt base timeout. Tests shorten it.
	attemptBase time.Duration

	// schemas caches compiled schemas by their bytes. A role's schema is
	// one constant document, so the compile runs once per process.
	schemas sync.Map
}

// Option tunes a Client.
type Option func(*Client)

// WithBudget sets the default budget. Zero fields take DefaultBudget values.
func WithBudget(b Budget) Option { return func(c *Client) { c.budget = b } }

// WithSleeper replaces the backoff sleep (tests).
func WithSleeper(s Sleeper) Option { return func(c *Client) { c.sleep = s } }

// WithLogger sets the logger.
func WithLogger(l *slog.Logger) Option { return func(c *Client) { c.log = l } }

// WithJitterSeed seeds the backoff jitter. Tests use it for a fixed sequence.
func WithJitterSeed(seed uint64) Option {
	return func(c *Client) { c.rng = rand.New(rand.NewPCG(seed, seed)) }
}

// WithoutJitter turns the backoff jitter off (tests).
func WithoutJitter() Option { return func(c *Client) { c.rng = nil } }

// New wires a Client. providers must cover every provider the config names.
func New(cfg *Config, providers []Provider, opts ...Option) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	c := &Client{
		cfg:         cfg,
		providers:   map[string]Provider{},
		budget:      DefaultBudget,
		attemptBase: attemptBase,
		sleep:       realSleep,
		log:         slog.Default(),
		rng:         rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}
	for _, p := range providers {
		if p == nil {
			return nil, errors.New("llm: nil provider given")
		}
		c.providers[p.Name()] = p
	}
	for _, name := range cfg.Providers() {
		if _, ok := c.providers[name]; !ok {
			return nil, fmt.Errorf("llm: config needs provider %q but none was given", name)
		}
	}
	for _, o := range opts {
		o(c)
	}
	c.budget = c.budget.withDefaults()
	return c, nil
}

// Config returns a copy of the resolved role map, for logs and eval rows.
func (c *Client) Config() *Config { return c.cfg.Clone() }

// Complete runs one logical call for role. acc may be nil.
//
// Retry classes: a truncation retries once at once with a higher cap,
// and the higher cap stays for the rest of the call: a transient failure
// on the escalated attempt does not change what the output needs. A
// transient failure retries after an exponential
// backoff, or after the provider's Retry-After hint when that is
// longer. A schema miss retries once, because the output is sampled and
// a second sample usually fits. Every other class returns at once. All
// attempts share one Budget, and each attempt gets its own timeout
// (see attemptTimeout). When the caller's ctx ends during an attempt,
// the result is ClassBudget.
func (c *Client) Complete(ctx context.Context, role Role, req Request, acc *Accumulator) (Response, error) {
	spec, ok := c.cfg.Roles[role]
	if !ok {
		return Response{}, newErr(ClassTerminal, "", "", 0, fmt.Errorf("unknown role %q", role))
	}
	prov := c.providers[spec.Provider]
	if prov == nil {
		return Response{}, newErr(ClassTerminal, spec.Provider, spec.Model, 0, fmt.Errorf("no provider %q wired", spec.Provider))
	}
	if len(req.Schema) == 0 {
		return Response{}, newErr(ClassTerminal, spec.Provider, spec.Model, 0, errors.New("request has no schema"))
	}
	schema, err := c.schema(req.Schema)
	if err != nil {
		return Response{}, newErr(ClassTerminal, spec.Provider, spec.Model, 0, err)
	}
	if req.SchemaName == "" {
		req.SchemaName = string(role) + "_output"
	}

	start := time.Now()
	ctx, cancel := context.WithDeadline(ctx, start.Add(c.budget.Deadline))
	defer cancel()

	call := Call{Request: req, Role: role, Model: spec.Model, Effort: spec.Effort, MaxOutputTokens: spec.MaxOutputTokens}
	attempts := 0
	escalated := false
	transient := 0
	schemaMisses := 0
	for {
		if attempts >= c.budget.MaxAttempts {
			return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0,
				fmt.Errorf("attempt budget of %d used", c.budget.MaxAttempts))
		}
		if ctx.Err() != nil {
			return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0, ctx.Err())
		}
		attempts++
		t0 := time.Now()
		actx, acancel := context.WithTimeout(ctx, attemptTimeout(c.attemptBase, call.MaxOutputTokens, c.budget.Deadline-time.Since(start)))
		res, err := prov.Complete(actx, call)
		attemptEnded := actx.Err() != nil
		acancel()
		acc.Record(role, spec.Model, res.Usage, time.Since(t0))
		if err == nil {
			verr := schema.Validate(mustAny(res.Output))
			if verr == nil {
				res.Provider, res.Model = spec.Provider, spec.Model
				res.Attempts, res.Latency = attempts, time.Since(start)
				return res, nil
			}
			err = newErr(ClassSchema, spec.Provider, spec.Model, 0, verr)
		}
		// The caller's cancel or the deadline ended the attempt. The
		// provider's own error class does not matter then.
		if cerr := ctx.Err(); cerr != nil {
			return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0, fmt.Errorf("%w (attempt %d: %w)", cerr, attempts, err))
		}
		// The attempt's own timeout ended it, and the budget still has
		// time. That is a slow provider, so a retry can help. A schema
		// miss or a truncation is a complete answer that arrived as the
		// timer fired, and each keeps its own class and its own retry.
		if class := ClassOf(err); attemptEnded && class != ClassSchema && class != ClassTruncation {
			err = newErr(ClassTransient, spec.Provider, spec.Model, 0, fmt.Errorf("attempt %d timed out: %w", attempts, err))
		}
		switch ClassOf(err) {
		case ClassTruncation:
			if escalated {
				return Response{}, err
			}
			next := escalate(call.MaxOutputTokens, spec.Provider)
			if next == call.MaxOutputTokens {
				// The cap is at the maximum. The same cap gives the same truncation.
				return Response{}, err
			}
			escalated = true
			call.MaxOutputTokens = next
			c.log.Warn("llm truncation, retry with higher cap", "role", role, "model", spec.Model, "cap", call.MaxOutputTokens)
			continue
		case ClassSchema:
			if schemaMisses > 0 {
				return Response{}, err
			}
			schemaMisses++
			c.log.Warn("llm schema miss, one retry", "role", role, "model", spec.Model, "err", err)
			continue
		case ClassTransient:
			// No sleep before an attempt the budget forbids: the caller
			// learns the outcome now instead of after a full backoff.
			if attempts >= c.budget.MaxAttempts {
				return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0,
					fmt.Errorf("attempt budget of %d used: %w", c.budget.MaxAttempts, err))
			}
			transient++
			delay := c.backoff(transient)
			left := c.budget.Deadline - time.Since(start)
			// A hint longer than the budget can not be honored, and a sleep
			// that ends at the deadline buys nothing.
			if hint := retryAfterOf(err); hint > delay {
				if hint > left {
					return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0,
						fmt.Errorf("retry-after hint of %v exceeds the %v left of the deadline: %w", hint, left.Round(time.Millisecond), err))
				}
				delay = hint
			}
			if delay > left {
				delay = left
			}
			c.log.Warn("llm transient failure, backoff", "role", role, "model", spec.Model, "attempt", attempts, "delay", delay, "err", err)
			if serr := c.sleep(ctx, delay); serr != nil {
				return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0, fmt.Errorf("deadline during backoff: %w", serr))
			}
			continue
		default:
			return Response{}, err
		}
	}
}

// attemptTimeout applies the L-4 rule: base plus attemptPerKTokens per
// 1,000 tokens of cap, and never more than the budget that is left.
func attemptTimeout(base time.Duration, maxOutputTokens int, left time.Duration) time.Duration {
	d := base + time.Duration(maxOutputTokens/1000)*attemptPerKTokens
	if d > left {
		d = left
	}
	if d < 0 {
		d = 0
	}
	return d
}

// backoff returns the delay before transient retry n (1-based): BaseDelay
// doubled per retry, moved by the jitter, then capped at maxBackoff.
func (c *Client) backoff(n int) time.Duration {
	delay := c.budget.BaseDelay
	for i := 1; i < n && delay < maxBackoff; i++ {
		delay *= 2
	}
	c.rngMu.Lock()
	defer c.rngMu.Unlock()
	if c.rng != nil {
		// f is in [-jitterFraction, +jitterFraction).
		f := (c.rng.Float64()*2 - 1) * jitterFraction
		delay += time.Duration(float64(delay) * f)
	}
	if delay > maxBackoff {
		delay = maxBackoff
	}
	return delay
}

// retryAfterOf reads the provider's Retry-After hint from err, or 0.
func retryAfterOf(err error) time.Duration {
	var e *Error
	if errors.As(err, &e) {
		return e.RetryAfter
	}
	return 0
}

// escalate returns the next cap after a truncation, or limit when the
// cap can not grow. The result never exceeds the provider's own limit.
func escalate(limit int, provider string) int {
	n := limit * escalationFactor
	if n < escalationFloor {
		n = escalationFloor
	}
	top := escalationMax
	if m, ok := providerMaxOutput[provider]; ok && m < top {
		top = m
	}
	if n > top {
		n = top
	}
	if n <= limit {
		return limit
	}
	return n
}

// schema returns the compiled form of raw, from the cache when the same
// bytes were compiled before.
func (c *Client) schema(raw json.RawMessage) (*jsonschema.Schema, error) {
	key := string(raw)
	if s, ok := c.schemas.Load(key); ok {
		return s.(*jsonschema.Schema), nil
	}
	s, err := compileSchema(raw)
	if err != nil {
		return nil, err
	}
	got, _ := c.schemas.LoadOrStore(key, s)
	return got.(*jsonschema.Schema), nil
}

// compileSchema parses a JSON Schema draft 2020-12 document. The root must
// be an object schema: both providers require that. No loader is set, so a
// $ref to a file or URL fails instead of a read from disk.
func compileSchema(raw json.RawMessage) (*jsonschema.Schema, error) {
	doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("schema is not JSON: %w", err)
	}
	root, ok := doc.(map[string]any)
	if !ok {
		return nil, errors.New("schema must be a JSON object")
	}
	if t, _ := root["type"].(string); t != "object" {
		return nil, errors.New(`schema root must have "type": "object"`)
	}
	comp := jsonschema.NewCompiler()
	comp.UseLoader(jsonschema.SchemeURLLoader{})
	comp.AssertFormat()
	const name = "mem://request.json"
	if err := comp.AddResource(name, doc); err != nil {
		return nil, fmt.Errorf("schema: %w", err)
	}
	s, err := comp.Compile(name)
	if err != nil {
		return nil, fmt.Errorf("schema: %w", err)
	}
	return s, nil
}

// mustAny decodes provider output for the validator. Invalid JSON becomes
// a string so the validator reports a type error instead of a panic.
func mustAny(raw json.RawMessage) any {
	v, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))
	if err != nil {
		return string(raw)
	}
	return v
}
