// Package livespend records the provider spend of a paid command that a
// live-eval session runs, and sums it for the budget (D-1169, D-1170).
//
// A paid command writes two lines for each run: a start line before its
// first provider call, and an end line with the measured calls, tokens,
// and cost. It appends each line to the ledger and to the spend file of
// the bundle, and it prints each line on stderr after Marker, so the
// session log holds it. No session writes a line by hand.
//
// The ledger is a file outside the run folder. The script sets its
// append-only flag, and the profile of the session allows only an append
// to it. So a session can not delete, truncate, or rewrite it, and each
// start reads the spend from it (D-1171).
package livespend

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/nkramber/decktome/go/internal/llm"
)

// The variables that the live-eval script gives a session.
const (
	EnvBundle = "LIVE_EVAL_BUNDLE"
	EnvBudget = "LIVE_EVAL_BUDGET_USD"
	EnvLedger = "LIVE_EVAL_LEDGER"
)

// SpendFile is the name of the spend file in the bundle.
const SpendFile = "spend.jsonl"

// Marker starts each spend line on stderr. The script reads the lines
// after it from the session log (D-1170).
const Marker = "LIVE-EVAL-SPEND "

// Line is one record of a paid run. Event is "start" or "end". An end
// line with Measured false names its cause in Note.
type Line struct {
	ID         string     `json:"id"`
	Event      string     `json:"event"`
	Target     string     `json:"target"`
	Calls      int        `json:"calls,omitempty"`
	Unreported int        `json:"unreported,omitempty"`
	Tokens     *llm.Usage `json:"tokens,omitempty"`
	USD        float64    `json:"usd"`
	Measured   bool       `json:"measured"`
	Note       string     `json:"note,omitempty"`
}

// Meter is the spend record of one paid run. Outside a live eval it
// meters the calls and writes no line. A nil *Meter records nothing.
type Meter struct {
	id     string
	target string
	// paths are the ledger and the spend file of the bundle. Each line
	// goes to each path.
	paths  []string
	acc    *llm.Accumulator
	stderr io.Writer
	once   sync.Once
}

// Start opens the meter of a paid run. When getenv names a bundle, it
// refuses the run when the ledger or the spend file reads the budget as
// spent, and it writes the start line before any provider call. A live
// eval with no ledger, or with a ledger that does not read, starts no
// paid run (D-1171).
func Start(target string, getenv func(string) string, prices *llm.PriceTable, stderr io.Writer) (*Meter, error) {
	bundle := strings.TrimSpace(getenv(EnvBundle))
	if bundle == "" {
		return &Meter{target: target, acc: llm.NewAccumulator(prices)}, nil
	}
	budget, err := ParseBudget(getenv(EnvBudget))
	if err != nil {
		return nil, err
	}
	ledger := strings.TrimSpace(getenv(EnvLedger))
	if ledger == "" {
		return nil, fmt.Errorf("a live eval needs %s, the append-only ledger, to start a paid run (D-1171)", EnvLedger)
	}
	if _, err := os.Stat(ledger); err != nil {
		return nil, fmt.Errorf("the ledger of the live eval: %w (D-1171)", err)
	}
	path := filepath.Join(bundle, SpendFile)
	sum, err := SumFiles(budget, ledger, path)
	if err != nil {
		return nil, err
	}
	if sum.Spent() {
		return nil, fmt.Errorf("the live-eval budget of $%.2f is spent (%s): no paid run can start (D-1169)", budget, sum.Text())
	}
	if prices == nil {
		return nil, errors.New("a live-eval run needs prices.json to measure its spend (D-1169)")
	}
	raw := make([]byte, 8)
	if _, err := rand.Read(raw); err != nil {
		return nil, err
	}
	m := &Meter{id: hex.EncodeToString(raw), target: target, paths: []string{ledger, path}, acc: llm.NewAccumulator(prices), stderr: stderr}
	if err := m.write(Line{ID: m.id, Event: "start", Target: target}); err != nil {
		return nil, err
	}
	return m, nil
}

// Accumulator is the meter of the client. It is nil on a nil *Meter.
func (m *Meter) Accumulator() *llm.Accumulator {
	if m == nil {
		return nil
	}
	return m.acc
}

// End writes the end line one time. A run with no provider call writes
// a measured zero.
func (m *Meter) End() error {
	if m == nil {
		return nil
	}
	var err error
	m.once.Do(func() { err = m.write(EndLine(m.id, m.target, m.acc.Report())) })
	return err
}

// EndLine measures a report. A call with no usage, or a model with no
// price row, marks the line unmeasured (D-1169).
func EndLine(id, target string, rep llm.Report) Line {
	l := Line{ID: id, Event: "end", Target: target, Calls: rep.Calls, Unreported: rep.Unreported, Tokens: rep.Tokens}
	switch {
	case rep.Calls == 0:
		l.Measured = true
	case rep.Unreported > 0:
		l.Note = fmt.Sprintf("%d of %d calls gave no usage", rep.Unreported, rep.Calls)
	case rep.CostUSD == nil:
		l.Note = "a model has no price row in prices.json"
	default:
		l.USD = *rep.CostUSD
		l.Measured = true
	}
	return l
}

func (m *Meter) write(l Line) error {
	if len(m.paths) == 0 {
		return nil
	}
	raw, err := json.Marshal(l)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(m.stderr, "%s%s\n", Marker, raw)
	for _, p := range m.paths {
		if err := appendLine(p, raw); err != nil {
			return fmt.Errorf("the spend record of the live eval: %w", err)
		}
	}
	return nil
}

// appendLine opens p for an append alone. The ledger takes no other
// write (D-1171).
func appendLine(p string, raw []byte) error {
	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	if _, err := f.Write(append(raw, '\n')); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

// ParseBudget reads the budget in dollars. It must be a number above
// zero.
func ParseBudget(raw string) (float64, error) {
	v, err := strconv.ParseFloat(strings.TrimSpace(raw), 64)
	if err != nil || !(v > 0) {
		return 0, fmt.Errorf("%s %q is not a number above zero", EnvBudget, raw)
	}
	return v, nil
}

// Sum is the spend of one session or one item. USD is the sum of the
// measured runs. Unmeasured names each run with no measured end line.
type Sum struct {
	Budget     float64  `json:"budget"`
	USD        float64  `json:"usd"`
	Runs       int      `json:"runs"`
	Unmeasured []string `json:"unmeasured,omitempty"`
	// Refused counts the lines of the spend file that no paid command
	// wrote. They add no cost.
	Refused int `json:"refused,omitempty"`
	// FileOnly names each run that the spend file holds and no log holds.
	// It still counts.
	FileOnly []string `json:"file_only,omitempty"`
}

// Spent says whether the budget is used. An unmeasured run counts as
// the full budget (D-1169).
func (s Sum) Spent() bool { return len(s.Unmeasured) > 0 || s.USD >= s.Budget }

// Charged is the dollar figure the budget reads.
func (s Sum) Charged() float64 {
	if len(s.Unmeasured) > 0 && s.USD < s.Budget {
		return s.Budget
	}
	return s.USD
}

// Text is a short account of the sum for a notice.
func (s Sum) Text() string {
	parts := []string{fmt.Sprintf("measured $%.4f in %d runs", s.USD, s.Runs)}
	if len(s.Unmeasured) > 0 {
		parts = append(parts, fmt.Sprintf("%d unmeasured, counted as the full budget", len(s.Unmeasured)))
	}
	if s.Refused > 0 {
		parts = append(parts, fmt.Sprintf("%d hand-written lines refused", s.Refused))
	}
	if len(s.FileOnly) > 0 {
		parts = append(parts, fmt.Sprintf("%d runs absent from the logs", len(s.FileOnly)))
	}
	return strings.Join(parts, ", ")
}

// Ledger collects the lines of the logs and of the spend file. For one
// id, the highest cost wins, and one unmeasured end line makes the run
// unmeasured. So a forged line can not lower the spend (D-1170).
type Ledger struct {
	runs    map[string]*run
	refused int
}

type run struct {
	target   string
	ended    bool
	measured bool
	usd      float64
	inLog    bool
	inFile   bool
}

// NewLedger makes an empty ledger.
func NewLedger() *Ledger { return &Ledger{runs: map[string]*run{}} }

// add takes one line. fromLog says whether the line came from a log.
func (g *Ledger) add(l Line, fromLog bool) bool {
	if len(l.ID) != 16 || !isHex(l.ID) || (l.Event != "start" && l.Event != "end") {
		return false
	}
	r := g.runs[l.ID]
	if r == nil {
		r = &run{target: l.Target, measured: true}
		g.runs[l.ID] = r
	}
	if fromLog {
		r.inLog = true
	} else {
		r.inFile = true
	}
	if l.Event == "end" {
		r.ended = true
		if !l.Measured {
			r.measured = false
		}
		if l.USD > r.usd {
			r.usd = l.USD
		}
	}
	return true
}

func isHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

// AddFile reads a spend file. A line that is not a record of a paid
// command counts as refused.
func (g *Ledger) AddFile(r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		raw, err := br.ReadBytes('\n')
		if t := bytes.TrimSpace(raw); len(t) > 0 {
			var l Line
			if json.Unmarshal(t, &l) != nil || !g.add(l, false) {
				g.refused++
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

// AddLog reads one session log of stream-json events. It takes the
// marked lines of the tool results alone, so the text of a command or a
// message adds no run.
func (g *Ledger) AddLog(r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		raw, err := br.ReadBytes('\n')
		if len(bytes.TrimSpace(raw)) > 0 {
			for _, text := range toolResults(raw) {
				g.addMarked(text)
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func (g *Ledger) addMarked(text string) {
	for _, ln := range strings.Split(text, "\n") {
		i := strings.Index(ln, Marker)
		if i < 0 {
			continue
		}
		var l Line
		dec := json.NewDecoder(strings.NewReader(ln[i+len(Marker):]))
		if dec.Decode(&l) == nil {
			g.add(l, true)
		}
	}
}

// toolResults gives the text of each tool result of one event.
func toolResults(raw []byte) []string {
	var ev struct {
		Message struct {
			Content json.RawMessage `json:"content"`
		} `json:"message"`
	}
	if json.Unmarshal(raw, &ev) != nil {
		return nil
	}
	var items []struct {
		Type    string          `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if json.Unmarshal(ev.Message.Content, &items) != nil {
		return nil
	}
	var out []string
	for _, it := range items {
		if it.Type != "tool_result" {
			continue
		}
		var s string
		if json.Unmarshal(it.Content, &s) == nil {
			out = append(out, s)
			continue
		}
		var parts []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}
		if json.Unmarshal(it.Content, &parts) == nil {
			for _, p := range parts {
				if p.Type == "text" {
					out = append(out, p.Text)
				}
			}
		}
	}
	return out
}

// Sum totals the ledger. A run with no measured end line is unmeasured.
func (g *Ledger) Sum(budget float64, logs bool) Sum {
	s := Sum{Budget: budget, Refused: g.refused}
	ids := make([]string, 0, len(g.runs))
	for id := range g.runs {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		r := g.runs[id]
		s.Runs++
		s.USD += r.usd
		if !r.ended || !r.measured {
			s.Unmeasured = append(s.Unmeasured, r.target+" "+id)
		}
		if logs && !r.inLog {
			s.FileOnly = append(s.FileOnly, r.target+" "+id)
		}
	}
	return s
}

// SumFiles sums spend files alone, the ledger and the spend file of the
// bundle. An absent file adds no spend.
func SumFiles(budget float64, paths ...string) (Sum, error) {
	g := NewLedger()
	for _, p := range paths {
		if err := g.AddPath(p); err != nil {
			return Sum{}, err
		}
	}
	return g.Sum(budget, false), nil
}

// AddPath reads one spend file. An absent file adds nothing.
func (g *Ledger) AddPath(p string) error {
	f, err := os.Open(p)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	return g.AddFile(f)
}
