package llm

import (
	"sync"
	"time"
)

// Accumulator sums usage over one flow, normally one session (M-1).
// Thread it through every Complete call of the flow. It never fails.
type Accumulator struct {
	prices *PriceTable

	mu       sync.Mutex
	calls    int
	reported int
	tokens   Usage
	cost     float64
	unpriced int
	latency  time.Duration
	byRole   map[Role]*RoleUsage
}

// RoleUsage is the per-role slice of a report.
type RoleUsage struct {
	Calls  int    `json:"calls"`
	Tokens *Usage `json:"tokens"`
}

// Report is the flow total. A nil pointer means "not instrumented": no
// call reported that value. Zero means "measured zero".
type Report struct {
	// Calls counts provider attempts, retries included.
	Calls int `json:"calls"`
	// Tokens is nil when no attempt reported usage.
	Tokens *Usage `json:"tokens"`
	// CostUSD is nil when no attempt reported usage, or when any reported
	// attempt used a model with no price row.
	CostUSD *float64 `json:"cost_usd"`
	// LatencyMS is the summed wall-clock time of all attempts, in
	// milliseconds.
	LatencyMS int64               `json:"latency_ms"`
	ByRole    map[Role]*RoleUsage `json:"by_role"`
}

// NewAccumulator makes an empty accumulator. prices may be nil: then the
// report never carries a cost.
func NewAccumulator(prices *PriceTable) *Accumulator {
	return &Accumulator{prices: prices, byRole: map[Role]*RoleUsage{}}
}

// Record adds one attempt. u is nil when the provider gave no usage.
func (a *Accumulator) Record(role Role, model string, u *Usage, latency time.Duration) {
	if a == nil {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	// The zero value must work: a caller may write &Accumulator{}
	// instead of NewAccumulator, and the map is made on first use.
	if a.byRole == nil {
		a.byRole = map[Role]*RoleUsage{}
	}
	a.calls++
	a.latency += latency
	ru := a.byRole[role]
	if ru == nil {
		ru = &RoleUsage{}
		a.byRole[role] = ru
	}
	ru.Calls++
	if u == nil {
		return
	}
	a.reported++
	a.tokens.Add(*u)
	if ru.Tokens == nil {
		ru.Tokens = &Usage{}
	}
	ru.Tokens.Add(*u)
	if a.prices == nil {
		a.unpriced++
		return
	}
	usd, ok := a.prices.Cost(model, *u)
	if !ok {
		a.unpriced++
		return
	}
	a.cost += usd
}

// Report snapshots the totals. The result is a copy: a later Record does
// not change it.
func (a *Accumulator) Report() Report {
	if a == nil {
		return Report{}
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	r := Report{Calls: a.calls, LatencyMS: a.latency.Milliseconds(), ByRole: map[Role]*RoleUsage{}}
	for role, ru := range a.byRole {
		cp := &RoleUsage{Calls: ru.Calls}
		if ru.Tokens != nil {
			t := *ru.Tokens
			cp.Tokens = &t
		}
		r.ByRole[role] = cp
	}
	if a.reported == 0 {
		return r
	}
	t := a.tokens
	r.Tokens = &t
	if a.unpriced == 0 {
		c := a.cost
		r.CostUSD = &c
	}
	return r
}
