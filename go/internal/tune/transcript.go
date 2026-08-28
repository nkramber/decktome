// Package tune holds the automated evaluation lane (D-133). One gate run
// writes a document. This package reads that document back, so the eval
// role can score every question and the loop can compare one run with the
// one before it.
//
// The parser lives here and not in a command, because three commands and
// the loop driver all read the same document.
package tune

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Question is one question the agent asked, as the gate document records
// it.
type Question struct {
	Turn     int      `json:"turn"`
	Slot     string   `json:"slot"`
	Row      string   `json:"row"`
	Fit      float64  `json:"fit"`
	Filled   bool     `json:"filled"`
	Invented bool     `json:"invented"`
	Text     string   `json:"text"`
	Options  []string `json:"options,omitempty"`
	// Catalog is the row text an invented question replaced.
	Catalog string `json:"catalog,omitempty"`
	// Refused is a replacement the agent refused as a reword (D-88).
	Refused string `json:"refused,omitempty"`
}

// Conversation is one scripted conversation and what it produced.
type Conversation struct {
	Name      string     `json:"name"`
	Probe     bool       `json:"probe"`
	Messages  []string   `json:"messages"`
	Questions []Question `json:"questions"`
	Premature bool       `json:"premature"`
	// Collection says a card collection was attached to the session. The
	// eval needs it: without it, the card-pool question reads as a
	// presumption that the user owns cards (D-143).
	Collection bool `json:"collection"`
	// Unanswered names the slots the deck needs that no answer filled.
	Unanswered string `json:"unanswered,omitempty"`
	// AfterBuild marks a conversation that starts after a finished build.
	// The catalog-only bar leaves it out, the way it leaves a probe out,
	// because its build slots were closed before the first message (A-9).
	AfterBuild bool `json:"after_build,omitempty"`
	// Failed holds the error a conversation ended on. The gate writes the
	// marker, and a conversation that errored counts for nobody in a
	// paired comparison (T-5).
	Failed string `json:"failed,omitempty"`
}

// Errored reports whether the conversation ended on an error.
func (c Conversation) Errored() bool { return c.Failed != "" }

// Run is one gate document.
type Run struct {
	Name          string         `json:"name"`
	Path          string         `json:"path"`
	Verdict       string         `json:"verdict"`
	Conversations []Conversation `json:"conversations"`
	// Metrics are the M-4 counts the document reports.
	Metrics Metrics `json:"metrics"`
}

// Metrics are the deterministic counters of one run. They are the guard
// against a run that scores well by asking less (D-91). Gate run 7 of
// 2026-08-25 passed both bars and left 26 of 30 conversations with a slot
// unanswered.
type Metrics struct {
	Questions     int `json:"questions"`
	Catalog       int `json:"catalog"`
	Invented      int `json:"invented"`
	CatalogFilled int `json:"catalog_filled"`
	CatalogOnly   int `json:"catalog_only"`
	Premature     int `json:"premature"`
	LintFindings  int `json:"lint_findings"`
	// QuestionsSeen and FilledSeen count every question in the document,
	// probes included. The M-4 table counts the 30 gate conversations
	// alone, and the terse set of D-105 is where the hard cases live. The
	// loop reads these two, so a counter never depends on which half of
	// the document a table meant.
	QuestionsSeen int `json:"questions_seen"`
	FilledSeen    int `json:"filled_seen"`
}

var (
	convRe    = regexp.MustCompile(`^### (.+)$`)
	turnRe    = regexp.MustCompile(`^\*\*Turn (\d+), the user:\*\* (.*)$`)
	questRe   = regexp.MustCompile(`^- \[(catalog|INVENTED) slot=(\S+) row=(\S+) fit=([0-9.]+) filled=(\S+)\] (.+)$`)
	optionRe  = regexp.MustCompile(`^  - Options: (.+)$`)
	replaceRe = regexp.MustCompile(`^  - It replaced: (.+)$`)
	refuseRe  = regexp.MustCompile(`^  - Refused as a reword \(D-88\): (.+)$`)
	metricRe  = regexp.MustCompile(`^\| ([^|]+?) \| (\d+) \|$`)
	verdictRe = regexp.MustCompile(`^Verdict: (\w+)\.`)
	lintRe    = regexp.MustCompile(`^The linter found (\d+) defective`)
	prematRe  = regexp.MustCompile(`^(\d+) conversations called themselves complete`)
	unansRe   = regexp.MustCompile(`^Slots the deck needs and nobody answered: (.+)\.$`)
	// prematureRe reads the same list off a premature session. The old
	// reader matched the plain label alone, so a premature conversation
	// carried no unanswered slots (audit 2026-08-28).
	prematureRe = regexp.MustCompile(`^\*\*PREMATURE\.\*\* It called itself complete with these slots unanswered: (.+)\.$`)
	failedRe    = regexp.MustCompile(`^\*\*Failed: (.+)\*\*$`)
	collectRe   = regexp.MustCompile(`^Collection: (true|false)\.`)
)

// ReadRun parses one gate document.
func ReadRun(path string) (*Run, error) {
	f, err := os.Open(path) // #nosec G304 -- the caller names a gate document.
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	run := &Run{Path: path, Name: strings.TrimSuffix(filepath.Base(path), ".md")}
	var conv *Conversation
	turn := 0
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		switch {
		case convRe.MatchString(line):
			run.flush(conv)
			name := convRe.FindStringSubmatch(line)[1]
			conv = &Conversation{Name: name, Probe: strings.Contains(name, "(probe)"),
				AfterBuild: strings.Contains(name, "(after a build)")}
			turn = 0
		case verdictRe.MatchString(line):
			run.Verdict = verdictRe.FindStringSubmatch(line)[1]
		case lintRe.MatchString(line):
			run.Metrics.LintFindings = atoi(lintRe.FindStringSubmatch(line)[1])
		case prematRe.MatchString(line):
			run.Metrics.Premature = atoi(prematRe.FindStringSubmatch(line)[1])
		case conv == nil && metricRe.MatchString(line):
			m := metricRe.FindStringSubmatch(line)
			run.Metrics.set(strings.TrimSpace(m[1]), atoi(m[2]))
		case conv != nil && collectRe.MatchString(line):
			conv.Collection = collectRe.FindStringSubmatch(line)[1] == "true"
		case conv != nil && unansRe.MatchString(line):
			conv.Unanswered = unansRe.FindStringSubmatch(line)[1]
		case conv != nil && prematureRe.MatchString(line):
			conv.Premature = true
			conv.Unanswered = prematureRe.FindStringSubmatch(line)[1]
		case conv != nil && strings.HasPrefix(line, "**PREMATURE.**"):
			conv.Premature = true
		case conv != nil && failedRe.MatchString(line):
			conv.Failed = failedRe.FindStringSubmatch(line)[1]
		case conv != nil && turnRe.MatchString(line):
			m := turnRe.FindStringSubmatch(line)
			turn = atoi(m[1])
			conv.Messages = append(conv.Messages, strings.TrimSpace(m[2]))
		case conv != nil && questRe.MatchString(line):
			m := questRe.FindStringSubmatch(line)
			fit, _ := strconv.ParseFloat(m[4], 64)
			conv.Questions = append(conv.Questions, Question{
				Turn: turn, Slot: m[2], Row: m[3], Fit: fit,
				Filled: m[5] == "true", Invented: m[1] == "INVENTED",
				Text: strings.TrimSpace(m[6]),
			})
		case conv != nil && len(conv.Questions) > 0:
			last := &conv.Questions[len(conv.Questions)-1]
			switch {
			case optionRe.MatchString(line):
				last.Options = splitOptions(optionRe.FindStringSubmatch(line)[1])
			case replaceRe.MatchString(line):
				last.Catalog = strings.TrimSpace(replaceRe.FindStringSubmatch(line)[1])
			case refuseRe.MatchString(line):
				last.Refused = strings.TrimSpace(refuseRe.FindStringSubmatch(line)[1])
			}
		}
	}
	run.flush(conv)
	if err := sc.Err(); err != nil {
		return nil, err
	}
	for _, c := range run.Conversations {
		for _, q := range c.Questions {
			run.Metrics.QuestionsSeen++
			if q.Filled {
				run.Metrics.FilledSeen++
			}
		}
	}
	if len(run.Conversations) == 0 {
		return nil, fmt.Errorf("tune: %s holds no conversation", path)
	}
	return run, nil
}

// flush keeps a conversation that carried a message. The document has
// other "###" headings, such as the table of invented questions by row,
// and each one read as a 105th conversation with no questions. That
// phantom counted as catalog-only, so every recount was one high (D-181).
func (r *Run) flush(c *Conversation) {
	if c != nil && len(c.Messages) > 0 {
		r.Conversations = append(r.Conversations, *c)
	}
}

// set writes one M-4 table row onto the metrics.
func (m *Metrics) set(label string, n int) {
	switch label {
	case "Questions asked":
		m.Questions = n
	case "From the catalog":
		m.Catalog = n
	case "Invented by the model":
		m.Invented = n
	case "Catalog questions that closed a slot":
		m.CatalogFilled = n
	case "Catalog-only conversations":
		m.CatalogOnly = n
	}
}

// PriorWords joins every message the user sent up to and including one
// turn. The eval role reads it, so a question is judged against what the
// user had actually written when it went out.
func (c Conversation) PriorWords(turn int) string {
	if turn > len(c.Messages) {
		turn = len(c.Messages)
	}
	if turn <= 0 {
		return ""
	}
	return strings.Join(c.Messages[:turn], " ")
}

func splitOptions(s string) []string {
	var out []string
	for _, p := range strings.Split(s, " / ") {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
