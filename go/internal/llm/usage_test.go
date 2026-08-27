package llm

import "testing"

// TestZeroAccumulatorRecords covers the zero value. A caller that writes
// &Accumulator{} instead of NewAccumulator panicked on a nil map, and the
// panic reached a real generate call on 2026-08-27.
func TestZeroAccumulatorRecords(t *testing.T) {
	var a Accumulator
	a.Record(RoleGenerate, "m", &Usage{InputTokens: 10, OutputTokens: 5}, 0)
	got := a.Report()
	if got.Calls != 1 {
		t.Errorf("calls = %d, want 1", got.Calls)
	}
	if got.Tokens.InputTokens != 10 || got.Tokens.OutputTokens != 5 {
		t.Errorf("tokens = %+v, want 10 in and 5 out", got.Tokens)
	}
}
