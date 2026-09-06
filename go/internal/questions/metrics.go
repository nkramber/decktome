package questions

import "sort"

// MinGateSize is the number of gate conversations the roadmap asks for
// of the offline half, TestConversations. The live half in
// cmd/questions-gate holds 77 since D-522, and GateSize there names it.
const MinGateSize = 30

// Ask is the M-4 record of one question the agent sent. The roadmap asks
// for four facts per session: how many questions came from the catalog,
// how many the model invented, the gap score of each one, and whether an
// invented question filled its slot (M-4, F-17).
//
// The record lives in the private state, so it survives a resume. The
// M-5 scoring lane reads the same fields (D-66).
type Ask struct {
	QuestionID string `json:"question_id"`
	RowID      string `json:"row_id"`
	Slot       string `json:"slot"`
	Key        string `json:"key"`
	Invented   bool   `json:"invented"`
	// CatalogText is the row text an invented question replaced. The M-5
	// rubric asks whether the catalog was enough and whether the invention
	// is better. Neither can be scored without both texts (D-66).
	CatalogText string  `json:"catalog_text,omitempty"`
	Fit         float64 `json:"fit"`
	Threshold   float64 `json:"threshold"`
	// Filled is true when the key closed after this question went out.
	Filled bool `json:"filled"`
	// NearCopy marks a turn where the model offered a replacement and the
	// agent refused it as a reword (D-88).
	NearCopy bool `json:"near_copy,omitempty"`
	// RefusedText is the replacement the agent refused. The M-5 sheet
	// shows it, so the refusal can be judged, which is what decides the
	// reword threshold (D-88).
	RefusedText string `json:"refused_text,omitempty"`
	// ResolvedText is the catalog row after the placeholders are filled,
	// and before the ask role phrases it. The reword guard compares a
	// replacement against this text, and never against the phrasing that
	// went out. The M-5 sheet shows it beside the phrasing, so an exact
	// copy of the row can be told from a new question (D-116).
	ResolvedText string `json:"resolved_text,omitempty"`
	// Turn is the turn number, counted from 1.
	Turn int `json:"turn"`
}

// Coverage is the M-4 report of one session, or of many sessions
// together. Counts are questions, not turns.
type Coverage struct {
	Sessions int `json:"sessions"`
	Asked    int `json:"asked"`
	Catalog  int `json:"catalog"`
	Invented int `json:"invented"`
	// CatalogFilled and InventedFilled count the questions whose key
	// closed. An invented question that fills nothing is the failure the
	// metric looks for.
	CatalogFilled  int `json:"catalog_filled"`
	InventedFilled int `json:"invented_filled"`
	// CatalogOnly counts the sessions that invented nothing. The PR-7
	// gate needs at least 25 of 30.
	CatalogOnly int `json:"catalog_only"`
	// Fits holds the gap score of every question, sorted.
	Fits []float64 `json:"fits"`
	// NearCopies counts the replacements the agent refused as rewords
	// (D-88). A high count means the scorer flags style, not faults.
	NearCopies int `json:"near_copies"`
	// InventedByRow counts the invented questions per catalog row they
	// replaced. A row that repeats here is a catalog change candidate
	// (D-25, PR-15).
	InventedByRow map[string]int `json:"invented_by_row"`
}

// Add folds one session's records into the report.
func (c *Coverage) Add(asks []Ask) {
	if c.InventedByRow == nil {
		c.InventedByRow = map[string]int{}
	}
	c.Sessions++
	invented := 0
	for _, a := range asks {
		c.Asked++
		c.Fits = append(c.Fits, a.Fit)
		if a.NearCopy {
			c.NearCopies++
		}
		if a.Invented {
			invented++
			c.Invented++
			c.InventedByRow[a.RowID]++
			if a.Filled {
				c.InventedFilled++
			}
			continue
		}
		c.Catalog++
		if a.Filled {
			c.CatalogFilled++
		}
	}
	if invented == 0 {
		c.CatalogOnly++
	}
	sort.Float64s(c.Fits)
}

// MedianFit is the middle gap score. It reports 0 for an empty report.
func (c *Coverage) MedianFit() float64 {
	n := len(c.Fits)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return c.Fits[n/2]
	}
	return (c.Fits[n/2-1] + c.Fits[n/2]) / 2
}

// Metrics is the M-4 report of one session.
func (s *State) Metrics() Coverage {
	var c Coverage
	c.Add(s.Asks)
	return c
}
