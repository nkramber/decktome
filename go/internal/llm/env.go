package llm

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"time"
)

//go:embed fixtures/*.json
var fixturesFS embed.FS

// Fixtures is the embedded fixture set for the fixture Fake.
func Fixtures() fs.FS {
	sub, err := fs.Sub(fixturesFS, "fixtures")
	if err != nil {
		panic(err) // the embed directive guarantees the directory
	}
	return sub
}

// Environment variables the layer reads.
const (
	EnvOpenAIKey    = "OPENAI_API_KEY"
	EnvAnthropicKey = "ANTHROPIC_API_KEY"
	// EnvRequireKeys makes a missing key fatal and refuses the fake
	// provider. That is the default (D-3). Only the value "0" opts out:
	// local dev and tests set it and get the fixture Fake instead.
	EnvRequireKeys = "LLM_REQUIRE_KEYS"
)

// attemptTimeout bounds one provider attempt. The Client's Budget bounds
// the whole logical call.
const attemptTimeout = 120 * time.Second

// NewFromEnv builds the production Client. It loads roles.json, applies
// LLM_<ROLE>_* overrides, and wires one adapter per provider the config
// names. A provider whose key is absent is fatal, unless LLM_REQUIRE_KEYS=0.
// Then the fixture Fake stands in with a warning.
func NewFromEnv(getenv func(string) string, log *slog.Logger) (*Client, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	cfg.RequireKeys = getenv(EnvRequireKeys) != "0"
	if err := cfg.ApplyEnv(getenv); err != nil {
		return nil, err
	}
	warnUnpricedOverrides(cfg, getenv, log)
	fake := NewFake(Fixtures())
	var providers []Provider
	for _, name := range cfg.Providers() {
		var p Provider
		switch name {
		case OpenAIName:
			if key := getenv(EnvOpenAIKey); key != "" {
				p = NewOpenAI(key, attemptTimeout)
			}
		case AnthropicName:
			if key := getenv(EnvAnthropicKey); key != "" {
				p = NewAnthropic(key, attemptTimeout)
			}
		case FakeName:
			p = fake
		default:
			return nil, fmt.Errorf("llm: config names unknown provider %q", name)
		}
		if p == nil {
			if cfg.RequireKeys {
				return nil, fmt.Errorf("llm: provider %q has no API key and %s is not 0", name, EnvRequireKeys)
			}
			log.Warn("llm provider has no API key, fixture fake stands in", "provider", name)
			p = &renamed{Provider: fake, name: name}
		}
		providers = append(providers, p)
	}
	c, err := New(cfg, providers, WithLogger(log))
	if err != nil {
		return nil, err
	}
	for _, r := range Roles {
		s := cfg.Roles[r]
		log.Info("llm role", "role", r, "provider", s.Provider, "model", s.Model, "effort", s.Effort, "verified_at", cfg.VerifiedAt)
	}
	return c, nil
}

// warnUnpricedOverrides logs one warning per LLM_<ROLE>_MODEL override
// whose model has no price row. The session cost then reports null (M-1).
func warnUnpricedOverrides(cfg *Config, getenv func(string) string, log *slog.Logger) {
	prices, err := LoadPrices()
	if err != nil {
		log.Warn("llm price table failed to load, cost reports null", "err", err)
		return
	}
	for _, r := range Roles {
		if getenv(envKey(r, "MODEL")) == "" {
			continue
		}
		model := cfg.Roles[r].Model
		if _, ok := prices.Models[model]; !ok {
			log.Warn("llm model override has no price row, cost reports null", "role", r, "model", model)
		}
	}
}

// renamed lets the fixture Fake stand in under another provider's name.
type renamed struct {
	Provider
	name string
}

func (r *renamed) Name() string { return r.name }
