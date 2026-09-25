package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/tune"
)

// twoConversations is a gate document of two conversations, one question
// each, so a fault can land on either call.
const twoConversations = "# question gate\n\n## Conversations\n\n" +
	"### 1. first\n\nCollection: false. Catalog: 1. Invented: 0.\n\n" +
	"**Turn 1, the user:** A Commander deck.\n\n" +
	"- [catalog slot=colors row=colors fit=0.90 filled=true] Which colors?\n\n" +
	"### 2. second\n\nCollection: false. Catalog: 1. Invented: 0.\n\n" +
	"**Turn 1, the user:** A Modern deck.\n\n" +
	"- [catalog slot=budget row=budget fit=0.90 filled=true] What budget?\n\n"

func verdictStep(t *testing.T, row string) llm.Step {
	t.Helper()
	raw, err := json.Marshal(map[string]any{
		"verdicts": []map[string]any{{"turn": 1, "row": row, "warranted": "yes",
			"faults": []string{}, "catalog_action": "none", "reason": "fine"}},
		"missed": []string{},
	})
	if err != nil {
		t.Fatal(err)
	}
	return llm.Step{Output: raw, Usage: &llm.Usage{InputTokens: 1_000_000, OutputTokens: 100_000}}
}

// TestProviderFaultRecordsTheSpend is REV-077: a provider fault part way
// through a run still writes a partial summary with the cost so far.
func TestProviderFaultRecordsTheSpend(t *testing.T) {
	t.Setenv("QUESTIONS_EVAL", "1")
	fault := errors.New("the provider is down")
	cases := []struct {
		name     string
		steps    func(t *testing.T) []llm.Step
		wantSum  bool
		wantCost bool
		wantRuns int
	}{
		{"fault on the first call", func(*testing.T) []llm.Step {
			return []llm.Step{{Err: fault}}
		}, false, false, 0},
		{"fault on the second call", func(t *testing.T) []llm.Step {
			return []llm.Step{verdictStep(t, "colors"), {Err: fault}}
		}, true, true, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			in := filepath.Join(dir, "gate.md")
			if err := os.WriteFile(in, []byte(twoConversations), 0o600); err != nil {
				t.Fatal(err)
			}
			sc := llm.NewScript(tc.steps(t)...)
			roles := map[llm.Role]llm.RoleSpec{}
			for _, r := range llm.Roles {
				roles[r] = llm.RoleSpec{Provider: llm.FakeName, Model: "gpt-5.6-luna", MaxOutputTokens: 4096}
			}
			client, err := llm.New(&llm.Config{VerifiedAt: "2026-08-24", Roles: roles}, []llm.Provider{sc}, llm.WithoutJitter())
			if err != nil {
				t.Fatal(err)
			}
			prev := newClient
			newClient = func() (*llm.Client, error) { return client, nil }
			t.Cleanup(func() { newClient = prev })

			out, jsonOut := filepath.Join(dir, "eval.md"), filepath.Join(dir, "eval.json")
			err = run(in, out, jsonOut, "", 5, 0, 0)
			if !errors.Is(err, fault) {
				t.Fatalf("err = %v, want the provider fault", err)
			}
			raw, rerr := os.ReadFile(jsonOut) // #nosec G304 -- a test temp file.
			if !tc.wantSum {
				// No call reported usage, so there is no spend to record.
				if rerr == nil {
					t.Errorf("a summary with no priced call: %s", raw)
				}
				return
			}
			if rerr != nil {
				t.Fatalf("no partial summary: %v", rerr)
			}
			var sum tune.Summary
			if err := json.Unmarshal(raw, &sum); err != nil {
				t.Fatal(err)
			}
			if !sum.Partial || !strings.Contains(sum.StoppedReason, "provider is down") {
				t.Errorf("partial = %v, stopped = %q, want a partial run that names the fault", sum.Partial, sum.StoppedReason)
			}
			if (sum.CostUSD > 0) != tc.wantCost {
				t.Errorf("cost_usd = %v, want spend recorded = %v", sum.CostUSD, tc.wantCost)
			}
			if len(sum.Verdicts) != tc.wantRuns {
				t.Errorf("verdicts = %d, want %d", len(sum.Verdicts), tc.wantRuns)
			}
			doc, rerr := os.ReadFile(out) // #nosec G304 -- a test temp file.
			if rerr != nil || !strings.Contains(string(doc), "provider is down") {
				t.Errorf("the eval document does not name the fault: %v", rerr)
			}
		})
	}
}
