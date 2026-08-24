package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
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

// DefaultBudget is used when the caller passes a zero Budget.
var DefaultBudget = Budget{MaxAttempts: 4, Deadline: 3 * time.Minute, BaseDelay: 2 * time.Second}

// Escalation bounds for a truncation retry, as in connector-syncer:
// the new cap is min(escalationMax, max(escalationFactor x cap, escalationFloor)).
// A small call can not become a huge call.
const (
	escalationFactor = 8
	escalationFloor  = 8192
	escalationMax    = 65536
)

// Client is the role layer. Call sites hold one Client and pass a Role.
type Client struct {
	cfg       *Config
	providers map[string]Provider
	budget    Budget
	sleep     Sleeper
	log       *slog.Logger
}

// Option tunes a Client.
type Option func(*Client)

// WithBudget sets the default budget.
func WithBudget(b Budget) Option { return func(c *Client) { c.budget = b } }

// WithSleeper replaces the backoff sleep (tests).
func WithSleeper(s Sleeper) Option { return func(c *Client) { c.sleep = s } }

// WithLogger sets the logger.
func WithLogger(l *slog.Logger) Option { return func(c *Client) { c.log = l } }

// New wires a Client. providers must cover every provider the config names.
func New(cfg *Config, providers []Provider, opts ...Option) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	c := &Client{cfg: cfg, providers: map[string]Provider{}, budget: DefaultBudget, sleep: realSleep, log: slog.Default()}
	for _, p := range providers {
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
	return c, nil
}

// Config returns the resolved role map, for logs and eval rows.
func (c *Client) Config() *Config { return c.cfg }

// Complete runs one logical call for role. acc may be nil.
//
// Retry classes: a truncation retries once at once with a higher cap. A
// transient failure retries after an exponential backoff. Every other class
// returns at once. All attempts share one Budget.
func (c *Client) Complete(ctx context.Context, role Role, req Request, acc *Accumulator) (Response, error) {
	spec, ok := c.cfg.Roles[role]
	if !ok {
		return Response{}, newErr(ClassTerminal, "", "", 0, fmt.Errorf("unknown role %q", role))
	}
	prov := c.providers[spec.Provider]
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
		switch ClassOf(err) {
		case ClassTruncation:
			if escalated {
				return Response{}, err
			}
			escalated = true
			call.MaxOutputTokens = escalate(call.MaxOutputTokens)
			c.log.Warn("llm truncation, retry with higher cap", "role", role, "model", spec.Model, "cap", call.MaxOutputTokens)
			continue
		case ClassTransient:
			transient++
			delay := c.budget.BaseDelay * (1 << (transient - 1))
			c.log.Warn("llm transient failure, backoff", "role", role, "model", spec.Model, "attempt", attempts, "delay", delay, "err", err)
			if serr := c.sleep(ctx, delay); serr != nil {
				return Response{}, newErr(ClassBudget, spec.Provider, spec.Model, 0, fmt.Errorf("deadline during backoff: %w", err))
			}
			continue
		default:
			return Response{}, err
		}
	}
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

// compileSchema parses a JSON Schema draft 2020-12 document.
func compileSchema(raw json.RawMessage) (*jsonschema.Schema, error) {
	doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("schema is not JSON: %w", err)
	}
	comp := jsonschema.NewCompiler()
	if err := comp.AddResource("request.json", doc); err != nil {
		return nil, fmt.Errorf("schema: %w", err)
	}
	s, err := comp.Compile("request.json")
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
