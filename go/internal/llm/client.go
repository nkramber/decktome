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

// Escalation bounds for a truncation retry, as in connector-syncer:
// the new cap is min(escalationMax, max(escalationFactor x cap, escalationFloor)).
// A small call can not become a huge call.
const (
	escalationFactor = 8
	escalationFloor  = 8192
	escalationMax    = 65536
)

// Backoff bounds. A transient delay never exceeds maxBackoff. Jitter moves
// each delay by up to jitterFraction in either direction, so many callers
// do not retry in step.
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
		cfg:       cfg,
		providers: map[string]Provider{},
		budget:    DefaultBudget,
		sleep:     realSleep,
		log:       slog.Default(),
		rng:       rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
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
// Retry classes: a truncation retries once at once with a higher cap. A
// transient failure retries after an exponential backoff. Every other class
// returns at once. All attempts share one Budget. When the caller's ctx
// ends during an attempt, the result is ClassBudget.
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
	schema, err := compileSchema(req.Schema)
	if err != nil {
		return Response{}, newErr(ClassTerminal, spec.Provider, spec.Model, 0, err)
	}
	if req.SchemaName == "" {
		req.SchemaName = string(role) + "_output"
	}

	ctx, cancel := context.WithTimeout(ctx, c.budget.Deadline)
	defer cancel()

	call := Call{Request: req, Role: role, Model: spec.Model, Effort: spec.Effort, MaxOutputTokens: spec.MaxOutputTokens}
	start := time.Now()
	attempts := 0
	escalated := false
	transient := 0
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
		res, err := prov.Complete(ctx, call)
		acc.Record(role, spec.Model, res.Usage, time.Since(t0))
		if err == nil {
			if verr := schema.Validate(mustAny(res.Output)); verr != nil {
				return Response{}, newErr(ClassSchema, spec.Provider, spec.Model, 0, verr)
			}
			res.Provider, res.Model = spec.Provider, spec.Model
			res.Attempts, res.Latency = attempts, time.Since(start)
			return res, nil
		}
		// The caller's cancel or the deadline ended the attempt. The
		// provider's own error class does not matter then.
		if cerr := ctx.Err(); cerr != nil {
			return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0, fmt.Errorf("%w (attempt %d: %w)", cerr, attempts, err))
		}
		switch ClassOf(err) {
		case ClassTruncation:
			if escalated {
				return Response{}, err
			}
			next := escalate(call.MaxOutputTokens)
			if next == call.MaxOutputTokens {
				// The cap is at the maximum. The same cap gives the same truncation.
				return Response{}, err
			}
			escalated = true
			call.MaxOutputTokens = next
			c.log.Warn("llm truncation, retry with higher cap", "role", role, "model", spec.Model, "cap", call.MaxOutputTokens)
			continue
		case ClassTransient:
			transient++
			delay := c.backoff(transient)
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

// backoff returns the delay before transient retry n (1-based): BaseDelay
// doubled per retry, capped at maxBackoff, then moved by the jitter.
func (c *Client) backoff(n int) time.Duration {
	delay := c.budget.BaseDelay
	for i := 1; i < n && delay < maxBackoff; i++ {
		delay *= 2
	}
	if delay > maxBackoff {
		delay = maxBackoff
	}
	c.rngMu.Lock()
	defer c.rngMu.Unlock()
	if c.rng == nil {
		return delay
	}
	// f is in [-jitterFraction, +jitterFraction).
	f := (c.rng.Float64()*2 - 1) * jitterFraction
	return delay + time.Duration(float64(delay)*f)
}

func escalate(limit int) int {
	n := limit * escalationFactor
	if n < escalationFloor {
		n = escalationFloor
	}
	if n > escalationMax {
		n = escalationMax
	}
	if n <= limit {
		return limit
	}
	return n
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
