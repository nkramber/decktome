package llm

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed roles.json
var rolesJSON []byte

//go:embed prices.json
var pricesJSON []byte

// RoleSpec is the resolved setting for one role.
type RoleSpec struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	// Effort is the provider's reasoning-effort value, or "" for the default.
	Effort string `json:"effort"`
	// MaxOutputTokens is the first-attempt output cap. A truncation retry
	// raises it once (see escalate).
	MaxOutputTokens int `json:"max_output_tokens"`
}

// Config is the frozen role-to-model map (D-1). roles.json is the default.
// Environment variables override one field of one role each.
type Config struct {
	VerifiedAt string            `json:"verified_at"`
	Note       string            `json:"note"`
	Roles      map[Role]RoleSpec `json:"roles"`
	// RequireKeys is true when every role must run on a real provider.
	// Validate then refuses the fake provider. NewFromEnv sets it from
	// LLM_REQUIRE_KEYS (D-3).
	RequireKeys bool `json:"-"`
}

// Clone returns a deep copy. Changes to the copy do not reach c.
func (c *Config) Clone() *Config {
	cp := *c
	cp.Roles = make(map[Role]RoleSpec, len(c.Roles))
	for r, s := range c.Roles {
		cp.Roles[r] = s
	}
	return &cp
}

// Price is USD per one million tokens for one model.
type Price struct {
	Input       float64 `json:"input"`
	CachedInput float64 `json:"cached_input"`
	// CacheWrite is the price of one cache-write token. Anthropic charges
	// 1.25 x input. OpenAI charges nothing extra, so its rows hold 0.
	CacheWrite float64 `json:"cache_write"`
	Output     float64 `json:"output"`
}

// PriceTable is the dated price list (M-1).
type PriceTable struct {
	VerifiedAt string           `json:"verified_at"`
	Source     string           `json:"source"`
	Models     map[string]Price `json:"models"`
}

// LoadConfig parses the embedded roles.json.
func LoadConfig() (*Config, error) {
	var cfg Config
	if err := json.Unmarshal(rolesJSON, &cfg); err != nil {
		return nil, fmt.Errorf("llm: roles.json: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// LoadPrices parses the embedded prices.json.
func LoadPrices() (*PriceTable, error) {
	var t PriceTable
	if err := json.Unmarshal(pricesJSON, &t); err != nil {
		return nil, fmt.Errorf("llm: prices.json: %w", err)
	}
	if t.VerifiedAt == "" {
		return nil, fmt.Errorf("llm: prices.json has no verified_at")
	}
	return &t, nil
}

// Validate checks that every role is set and that the judge does not share
// the generator's provider (D-22). With RequireKeys, no role may use the
// fake provider (D-3).
func (c *Config) Validate() error {
	if c.VerifiedAt == "" {
		return fmt.Errorf("llm: config has no verified_at")
	}
	for _, r := range Roles {
		spec, ok := c.Roles[r]
		if !ok {
			return fmt.Errorf("llm: config has no role %q", r)
		}
		if spec.Provider == "" || spec.Model == "" {
			return fmt.Errorf("llm: role %q needs a provider and a model", r)
		}
		if spec.MaxOutputTokens <= 0 {
			return fmt.Errorf("llm: role %q needs max_output_tokens > 0", r)
		}
		if err := validEffort(spec.Provider, spec.Effort); err != nil {
			return fmt.Errorf("llm: role %q: %w", r, err)
		}
		if c.RequireKeys && spec.Provider == FakeName {
			return fmt.Errorf("llm: role %q uses provider %q but %s is not 0", r, FakeName, EnvRequireKeys)
		}
	}
	for r := range c.Roles {
		if !knownRole(r) {
			return fmt.Errorf("llm: config names unknown role %q", r)
		}
	}
	gen, judge := c.Roles[RoleGenerate], c.Roles[RoleJudge]
	// The fake exemption only holds when keys are not required.
	if gen.Provider == judge.Provider && (gen.Provider != FakeName || c.RequireKeys) {
		return fmt.Errorf("llm: judge provider %q equals generate provider (D-22)", judge.Provider)
	}
	return nil
}

// ApplyEnv overrides fields from variables of the form LLM_<ROLE>_<FIELD>,
// for example LLM_GENERATE_MODEL or LLM_JUDGE_PROVIDER. FIELD is one of
// PROVIDER, MODEL, EFFORT. getenv is os.Getenv in production.
func (c *Config) ApplyEnv(getenv func(string) string) error {
	for _, r := range Roles {
		spec := c.Roles[r]
		if v := getenv(envKey(r, "PROVIDER")); v != "" {
			spec.Provider = v
		}
		if v := getenv(envKey(r, "MODEL")); v != "" {
			spec.Model = v
		}
		if v := getenv(envKey(r, "EFFORT")); v != "" {
			spec.Effort = v
		}
		c.Roles[r] = spec
	}
	return c.Validate()
}

// envKey builds the override variable name for one role field.
func envKey(r Role, field string) string {
	return "LLM_" + strings.ToUpper(string(r)) + "_" + field
}

// Providers lists the distinct provider names the config needs.
func (c *Config) Providers() []string {
	seen := map[string]bool{}
	var out []string
	for _, r := range Roles {
		p := c.Roles[r].Provider
		if !seen[p] {
			seen[p] = true
			out = append(out, p)
		}
	}
	return out
}

func knownRole(r Role) bool {
	for _, k := range Roles {
		if k == r {
			return true
		}
	}
	return false
}

// Cost prices one usage report in USD. ok is false when the model has no
// price row: the caller reports null, not zero (M-1). Fresh input is the
// input total minus cache reads and cache writes.
func (t *PriceTable) Cost(model string, u Usage) (usd float64, ok bool) {
	p, ok := t.Models[model]
	if !ok {
		return 0, false
	}
	fresh := u.InputTokens - u.CachedInputTokens - u.CacheWriteTokens
	if fresh < 0 {
		fresh = 0
	}
	usd = (float64(fresh)*p.Input +
		float64(u.CachedInputTokens)*p.CachedInput +
		float64(u.CacheWriteTokens)*p.CacheWrite +
		float64(u.OutputTokens)*p.Output) / 1e6
	return usd, true
}

// efforts lists the reasoning-effort values each provider accepts, as
// its SDK names them (anthropic-sdk-go v1.66.0, openai-go v3.52.0).
var efforts = map[string][]string{
	AnthropicName: {"low", "medium", "high", "xhigh", "max"},
	OpenAIName:    {"none", "minimal", "low", "medium", "high", "xhigh", "max"},
}

// validEffort checks an LLM_<ROLE>_EFFORT value against the provider's
// list at load time, so a typo fails at startup and not on the first
// user turn. The fake accepts every value, and "" is the default.
func validEffort(provider, effort string) error {
	if effort == "" {
		return nil
	}
	allowed, ok := efforts[provider]
	if !ok {
		return nil
	}
	for _, v := range allowed {
		if v == effort {
			return nil
		}
	}
	return fmt.Errorf("effort %q is not one of %s for provider %q", effort, strings.Join(allowed, ", "), provider)
}
