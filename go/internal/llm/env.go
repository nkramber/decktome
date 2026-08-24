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
	// EnvRequireKeys, when "1", makes a missing key fatal. Production sets
	// it. Local mode leaves it unset and gets the fixture Fake instead.
	EnvRequireKeys = "LLM_REQUIRE_KEYS"
)

// attemptTimeout bounds one provider attempt. The Client's Budget bounds
// the whole logical call.
const attemptTimeout = 120 * time.Second

// NewFromEnv builds the production Client. It loads roles.json, applies
// LLM_<ROLE>_* overrides, and wires one adapter per provider the config
// names. A provider whose key is absent falls back to the fixture Fake with
// a warning, unless LLM_REQUIRE_KEYS=1.
func NewFromEnv(getenv func(string) string, log *slog.Logger) (*Client, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	if err := cfg.ApplyEnv(getenv); err != nil {
		return nil, err
	}
	require := getenv(EnvRequireKeys) == "1"
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
			if require {
				return nil, fmt.Errorf("llm: provider %q has no API key and %s=1", name, EnvRequireKeys)
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

// renamed lets the fixture Fake stand in under another provider's name.
type renamed struct {
	Provider
	name string
}

func (r *renamed) Name() string { return r.name }
