package generate

import (
	"testing"

	"github.com/nkramber/decktome/go/internal/llm"
)

// TestBuildMetricsShape is D-602. The build carries what it cost and how
// it went, so a reader counts the repair turns over many decks.
func TestBuildMetricsShape(t *testing.T) {
	// spendBetween is the build's own spend, and not the turn's.
	cost := func(v float64) *float64 { return &v }
	before := llm.Report{Calls: 2, Tokens: &llm.Usage{InputTokens: 100}, CostUSD: cost(1.0)}
	after := llm.Report{Calls: 5, Tokens: &llm.Usage{InputTokens: 400}, CostUSD: cost(4.0)}
	u := spendBetween(before, after)
	if u.GetCalls() != 3 {
		t.Errorf("calls = %d, want 3: the build made three of the five", u.GetCalls())
	}
	if u.GetInputTokens() != 300 {
		t.Errorf("input tokens = %d, want 300", u.GetInputTokens())
	}
	if u.GetCostUsd() != 3.0 {
		t.Errorf("cost = %v, want 3", u.GetCostUsd())
	}
	if !u.GetPriced() {
		t.Error("a priced pair answers a priced total")
	}
}
