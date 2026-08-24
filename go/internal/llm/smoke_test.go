package llm

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"
)

// TestSmokeLiveProviders calls the real providers. It runs only with
// LLM_SMOKE=1 and the API keys in the environment (make test-smoke).
// It exercises one OpenAI role (classify) and the Anthropic judge role,
// so both adapters, strict schema output, and usage accounting are proven
// against the live APIs.
func TestSmokeLiveProviders(t *testing.T) {
	if os.Getenv("LLM_SMOKE") != "1" {
		t.Skip("set LLM_SMOKE=1 and the API keys to run the live smoke")
	}
	env := func(k string) string {
		if k == EnvRequireKeys {
			return "1"
		}
		return os.Getenv(k)
	}
	c, err := NewFromEnv(env, discardLogger())
	if err != nil {
		t.Fatal(err)
	}
	prices, err := LoadPrices()
	if err != nil {
		t.Fatal(err)
	}
	acc := NewAccumulator(prices)

	schema := json.RawMessage(`{
	  "type":"object",
	  "properties":{
	    "format":{"type":"string","enum":["commander","standard","modern","pauper","unknown"]},
	    "theme":{"type":"string"}
	  },
	  "required":["format","theme"],
	  "additionalProperties":false
	}`)
	req := Request{
		Instructions: "You classify a Magic: The Gathering deck request. Reply only with the JSON the schema asks for.",
		Input:        "Build me a 100-card lifegain deck with Heliod as my commander.",
		Schema:       schema,
		CacheKey:     "smoke",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()

	for _, role := range []Role{RoleClassify, RoleJudge} {
		res, err := c.Complete(ctx, role, req, acc)
		if err != nil {
			t.Fatalf("%s: %v", role, err)
		}
		var out struct {
			Format string `json:"format"`
			Theme  string `json:"theme"`
		}
		if err := json.Unmarshal(res.Output, &out); err != nil {
			t.Fatalf("%s: output %s: %v", role, res.Output, err)
		}
		if out.Format != "commander" {
			t.Errorf("%s: format = %q, want commander (output %s)", role, out.Format, res.Output)
		}
		if res.Usage == nil || res.Usage.OutputTokens == 0 {
			t.Errorf("%s: usage not reported: %+v", role, res.Usage)
		}
		t.Logf("%s: provider=%s model=%s attempts=%d latency=%s usage=%+v output=%s",
			role, res.Provider, res.Model, res.Attempts, res.Latency, *res.Usage, res.Output)
	}
	rep := acc.Report()
	if rep.CostUSD == nil {
		t.Fatalf("cost is null: report %+v", rep)
	}
	t.Logf("session report: calls=%d tokens=%+v cost_usd=%.6f", rep.Calls, *rep.Tokens, *rep.CostUSD)
}
