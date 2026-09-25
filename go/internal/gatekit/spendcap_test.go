package gatekit

import (
	"testing"

	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/llm"
)

// TestSpendCap: a gate of the fix cycle stops before an item once it
// spent the room under the cap of the cycle (REV-078, D-939).
func TestSpendCap(t *testing.T) {
	usd := func(v float64) *float64 { return &v }
	tests := []struct {
		name     string
		raw      string
		rep      llm.Report
		wantErr  bool
		wantStop bool
	}{
		{name: "no cap never stops", raw: "", rep: llm.Report{Calls: 9, CostUSD: usd(90)}},
		{name: "under the cap", raw: "0.50", rep: llm.Report{Calls: 3, CostUSD: usd(0.49)}},
		{name: "at the cap", raw: "0.50", rep: llm.Report{Calls: 3, CostUSD: usd(0.50)}, wantStop: true},
		{name: "over the cap", raw: "0.50", rep: llm.Report{Calls: 4, CostUSD: usd(0.73)}, wantStop: true},
		{name: "unpriced calls stop", raw: "0.50", rep: llm.Report{Calls: 2}, wantStop: true},
		{name: "no call yet", raw: "0.50", rep: llm.Report{}},
		{name: "a word", raw: "lots", wantErr: true},
		{name: "zero", raw: "0", wantErr: true},
		{name: "under zero", raw: "-1", wantErr: true},
		{name: "infinite", raw: "Inf", wantErr: true},
		{name: "not a number", raw: "NaN", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewSpendCap(func(k string) string {
				if k == EnvMaxUSD {
					return tt.raw
				}
				return ""
			})
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewSpendCap(%q) err = %v, wantErr %v", tt.raw, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			rec := evalrun.New("decks", "r1")
			if got := c.Stop(rec, tt.rep); got != tt.wantStop {
				t.Fatalf("Stop = %v, want %v", got, tt.wantStop)
			}
			if tt.wantStop != (rec.Header.Stopped != "") {
				t.Errorf("Stopped = %q, want set %v", rec.Header.Stopped, tt.wantStop)
			}
			if tt.wantStop && !rec.Header.Partial() {
				t.Error("a stopped run stands for its suite")
			}
		})
	}
}
