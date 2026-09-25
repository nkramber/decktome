package main

import (
	"testing"

	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/llm"
)

// TestASpentCapAfterTheLastConversationKeepsTheRunWhole is the Gitar
// finding of #230: the rerun guard marked a run stopped after it played
// every conversation, and the fix cycle then threw its measure away.
func TestASpentCapAfterTheLastConversationKeepsTheRunWhole(t *testing.T) {
	c, err := gatekit.NewSpendCap(func(k string) string {
		if k == gatekit.EnvMaxUSD {
			return "0.50"
		}
		return ""
	})
	if err != nil {
		t.Fatal(err)
	}
	usd := func(v float64) *float64 { return &v }
	tests := []struct {
		name string
		rep  llm.Report
		skip bool
	}{
		{name: "under the cap plays the reruns", rep: llm.Report{Calls: 4, CostUSD: usd(0.30)}},
		{name: "at the cap skips them", rep: llm.Report{Calls: 6, CostUSD: usd(0.50)}, skip: true},
		{name: "unpriced calls skip them", rep: llm.Report{Calls: 6}, skip: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := evalrun.New("questions", "r1")
			if got := skipReruns(rec, c, tt.rep); got != tt.skip {
				t.Errorf("skipReruns = %v, want %v", got, tt.skip)
			}
			if rec.Header.Stopped != "" || rec.Header.Partial() {
				t.Errorf("a run that played every conversation reads stopped: %q", rec.Header.Stopped)
			}
		})
	}
}
