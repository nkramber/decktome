package tune

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixture = "# PR-7 question gate\n\n" +
	"Verdict: FAIL. 27 of 30 gate conversations used catalog questions only. The bar is 25.\n\n" +
	"2 conversations called themselves complete with a slot still unanswered, which fails the gate: a; b.\n\n" +
	"The linter found 3 defective questions, which fails the gate (D-115).\n\n" +
	"## M-4 report\n\n" +
	"| Measure | Value |\n|---|---|\n" +
	"| Conversations | 30 |\n" +
	"| Questions asked | 136 |\n" +
	"| From the catalog | 133 |\n" +
	"| Invented by the model | 3 |\n" +
	"| Catalog questions that closed a slot | 89 |\n" +
	"| Catalog-only conversations | 27 |\n\n" +
	"## Conversations\n\n" +
	"### 23. land destruction\n\n" +
	"Collection: true. Catalog: 2. Invented: 0.\n\n" +
	"Slots the deck needs and nobody answered: commander (asked, no answer).\n\n" +
	"**Turn 1, the user:** A land destruction Commander deck.\n\n" +
	"- [catalog slot=format row=format fit=0.90 filled=true] What format would you like?\n" +
	"  - Options: Commander / Standard\n\n" +
	"**Turn 2, the user:** Commander. Red and green.\n\n" +
	"- [INVENTED slot=colors row=colors fit=0.20 filled=true] Which colors?\n" +
	"  - It replaced: Any color preference?\n\n" +
	"### 44. sideboard help only (probe)\n\n" +
	"Collection: false. Catalog: 1. Invented: 0.\n\n" +
	"**PREMATURE.** It called itself complete with these slots unanswered: colors (never asked).\n\n" +
	"**Turn 1, the user:** I already have a Modern burn deck.\n\n" +
	"- [catalog slot=power row=power_sixty fit=0.90 filled=false] How strong?\n" +
	"  - Refused as a reword (D-88): How strong should it be?\n"

func writeFixture(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "pr7-question-gate-runX.md")
	if err := os.WriteFile(path, []byte(fixture), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestReadRun(t *testing.T) {
	run, err := ReadRun(writeFixture(t))
	if err != nil {
		t.Fatalf("ReadRun: %v", err)
	}
	if run.Name != "pr7-question-gate-runX" || run.Verdict != "FAIL" {
		t.Errorf("name = %q, verdict = %q", run.Name, run.Verdict)
	}
	want := Metrics{Questions: 136, Catalog: 133, Invented: 3, CatalogFilled: 89,
		CatalogOnly: 27, Premature: 2, LintFindings: 3, QuestionsSeen: 3, FilledSeen: 2}
	if !run.Conversations[0].Collection || run.Conversations[1].Collection {
		t.Errorf("the collection flag was not read: %v, %v",
			run.Conversations[0].Collection, run.Conversations[1].Collection)
	}
	if run.Metrics != want {
		t.Errorf("metrics = %+v, want %+v", run.Metrics, want)
	}
	if len(run.Conversations) != 2 {
		t.Fatalf("%d conversations, want 2", len(run.Conversations))
	}
	c := run.Conversations[0]
	if c.Name != "23. land destruction" || c.Probe {
		t.Errorf("first conversation = %q, probe = %v", c.Name, c.Probe)
	}
	if len(c.Messages) != 2 || c.Messages[0] != "A land destruction Commander deck." {
		t.Errorf("messages = %v", c.Messages)
	}
	if c.Unanswered == "" {
		t.Error("the unanswered slots were not read")
	}
	if len(c.Questions) != 2 {
		t.Fatalf("%d questions, want 2", len(c.Questions))
	}
	if q := c.Questions[0]; q.Turn != 1 || q.Row != "format" || q.Fit != 0.90 || !q.Filled || len(q.Options) != 2 {
		t.Errorf("first question = %+v", q)
	}
	if q := c.Questions[1]; !q.Invented || q.Turn != 2 || q.Catalog != "Any color preference?" {
		t.Errorf("second question = %+v", q)
	}
	probe := run.Conversations[1]
	if !probe.Probe || !probe.Premature {
		t.Errorf("probe = %v, premature = %v", probe.Probe, probe.Premature)
	}
	if probe.Questions[0].Refused == "" {
		t.Error("the refused reword was not read")
	}
	// A question is judged against the words the user had written then.
	if got := c.PriorWords(1); got != "A land destruction Commander deck." {
		t.Errorf("PriorWords(1) = %q", got)
	}
}

func TestSummarizeCountsOnlyDecidedVerdicts(t *testing.T) {
	vs := []Verdict{
		{Row: "colors", Warranted: "yes"},
		{Row: "format", Warranted: "no", Faults: []string{"duplicate"}, CatalogAction: "reword"},
		{Row: "format", Warranted: "no", Faults: []string{"Duplicate"}, CatalogAction: "reword"},
		{Row: "meta", Warranted: "unsure"},
	}
	s := Summarize("runX", "gpt-5.6-luna", 0.11, Metrics{Questions: 4}, vs)
	if s.Judged != 3 || s.Bad != 2 || s.Unsure != 1 {
		t.Fatalf("judged = %d, bad = %d, unsure = %d", s.Judged, s.Bad, s.Unsure)
	}
	if s.Ratio < 0.66 || s.Ratio > 0.67 {
		t.Errorf("ratio = %v, want two of three", s.Ratio)
	}
	if s.ByRow["format"] != 2 {
		t.Errorf("by row = %v", s.ByRow)
	}
	// A fault is counted once, whatever case the model wrote.
	if s.ByFault["duplicate"] != 2 {
		t.Errorf("by fault = %v", s.ByFault)
	}
	if got := s.TopRows(1); len(got) != 1 || got[0] != "format" {
		t.Errorf("top rows = %v", got)
	}
}

// TestCompareRefusesAQuietRun is the guard the owner's own history asks
// for. Gate run 7 of 2026-08-25 passed both bars and left 26 of 30
// conversations with a slot unanswered. Fewer questions means a better
// ratio and a worse product.
func TestCompareRefusesAQuietRun(t *testing.T) {
	prev := &Summary{Judged: 100, Bad: 20, Ratio: 0.20,
		Metrics: Metrics{Questions: 200, CatalogFilled: 150}}
	cases := []struct {
		name   string
		next   *Summary
		accept bool
	}{
		{"a better ratio with the same work", &Summary{Judged: 100, Bad: 10, Ratio: 0.10,
			Metrics: Metrics{Questions: 198, CatalogFilled: 149}}, true},
		{"a better ratio because the agent went quiet", &Summary{Judged: 40, Bad: 2, Ratio: 0.05,
			Metrics: Metrics{Questions: 90, CatalogFilled: 60}}, false},
		{"fewer slots closed", &Summary{Judged: 100, Bad: 5, Ratio: 0.05,
			Metrics: Metrics{Questions: 200, CatalogFilled: 100}}, false},
		{"a worse ratio", &Summary{Judged: 100, Bad: 30, Ratio: 0.30,
			Metrics: Metrics{Questions: 200, CatalogFilled: 150}}, false},
		{"a defective question", &Summary{Judged: 100, Bad: 5, Ratio: 0.05,
			Metrics: Metrics{Questions: 200, CatalogFilled: 150, LintFindings: 1}}, false},
		{"a premature session", &Summary{Judged: 100, Bad: 5, Ratio: 0.05,
			Metrics: Metrics{Questions: 200, CatalogFilled: 150, Premature: 1}}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := Compare(prev, tc.next)
			if d.Accept != tc.accept {
				t.Errorf("accept = %v, want %v: %v", d.Accept, tc.accept, d.Reasons)
			}
			if !d.Accept && len(d.Reasons) == 0 {
				t.Error("a reject named no reason")
			}
		})
	}
}

// TestCompareFirstRunHasNothingToBeat covers the first iteration.
func TestCompareFirstRunHasNothingToBeat(t *testing.T) {
	next := &Summary{Judged: 10, Bad: 9, Ratio: 0.90, Metrics: Metrics{Questions: 10}}
	if d := Compare(nil, next); !d.Accept {
		t.Errorf("the first run was rejected: %v", d.Reasons)
	}
	bad := &Summary{Judged: 10, Metrics: Metrics{LintFindings: 2}}
	if d := Compare(nil, bad); d.Accept {
		t.Error("a first run with a defective question was accepted")
	}
}

// TestHoldoutSplitDrivesTheDecision is D-134. A fixer reads the tune
// split alone. A ratio that falls only there is a reworded test set, not
// a better product.
func TestHoldoutSplitDrivesTheDecision(t *testing.T) {
	mk := func(tuneBad, tuneOK, holdBad, holdOK int) *Summary {
		var vs []Verdict
		add := func(n int, warranted string, hold bool) {
			for i := 0; i < n; i++ {
				vs = append(vs, Verdict{Row: "colors", Warranted: warranted, Holdout: hold})
			}
		}
		add(tuneBad, "no", false)
		add(tuneOK, "yes", false)
		add(holdBad, "no", true)
		add(holdOK, "yes", true)
		s := Summarize("runX", "m", 0, Metrics{Questions: 100, CatalogFilled: 80}, vs)
		return &s
	}
	prev := mk(10, 30, 10, 30)
	if got, split := prev.Scored(); split != "holdout" || got != 0.25 {
		t.Fatalf("scored = %v on %q, want the holdout at 0.25", got, split)
	}
	// The fixer reworded what it was shown and nothing else.
	overfit := mk(2, 38, 10, 30)
	d := Compare(prev, overfit)
	if !d.Accept {
		t.Errorf("an equal holdout was rejected: %v", d.Reasons)
	}
	var warned bool
	for _, r := range d.Reasons {
		if strings.Contains(r, "WARNING") {
			warned = true
		}
	}
	if !warned {
		t.Errorf("the loop did not warn about the overfit: %v", d.Reasons)
	}
	// A real gain moves the holdout too.
	genuine := mk(2, 38, 2, 38)
	if d := Compare(prev, genuine); !d.Accept {
		t.Errorf("a real gain was rejected: %v", d.Reasons)
	}
	// A holdout that gets worse is rejected, whatever the tune split says.
	worse := mk(0, 40, 20, 20)
	if d := Compare(prev, worse); d.Accept {
		t.Error("a worse holdout was accepted")
	}
}

// TestSummarizeWithNoHoldoutFallsBackToEveryConversation covers -holdout 0.
func TestSummarizeWithNoHoldoutFallsBackToEveryConversation(t *testing.T) {
	s := Summarize("runX", "m", 0, Metrics{}, []Verdict{
		{Row: "a", Warranted: "no"}, {Row: "b", Warranted: "yes"},
	})
	got, split := s.Scored()
	if split != "every conversation" || got != 0.5 {
		t.Errorf("scored = %v on %q, want 0.5 on every conversation", got, split)
	}
}

// TestReadRunOnARealDocument guards the parser against format drift. The
// gate writer and this reader are separate files, and a changed heading
// would leave the loop reading nothing.
func TestReadRunOnARealDocument(t *testing.T) {
	matches, err := filepath.Glob("../../../docs/reference/pr7-question-gate-run1*.md")
	if err != nil || len(matches) == 0 {
		t.Skip("no gate document on disk")
	}
	run, err := ReadRun(matches[len(matches)-1])
	if err != nil {
		t.Fatalf("ReadRun(%s): %v", matches[len(matches)-1], err)
	}
	if len(run.Conversations) < 10 {
		t.Errorf("%s gave %d conversations", run.Name, len(run.Conversations))
	}
	if run.Metrics.Questions == 0 {
		t.Errorf("%s gave no question count", run.Name)
	}
	var questions, withMessages int
	for _, c := range run.Conversations {
		questions += len(c.Questions)
		if len(c.Messages) > 0 {
			withMessages++
		}
	}
	if questions == 0 || withMessages == 0 {
		t.Errorf("%s gave %d questions over %d conversations with messages",
			run.Name, questions, withMessages)
	}
	// The transcript counters read every conversation, probes included.
	// The M-4 table reads the 30 gate conversations alone, so the two
	// numbers differ on purpose.
	if run.Metrics.QuestionsSeen != questions {
		t.Errorf("QuestionsSeen = %d, and the test counted %d", run.Metrics.QuestionsSeen, questions)
	}
	if run.Metrics.QuestionsSeen < run.Metrics.Questions {
		t.Errorf("the transcript holds %d questions and the table claims %d",
			run.Metrics.QuestionsSeen, run.Metrics.Questions)
	}
}

// TestQuietRunIsCaughtByTheTranscriptCount proves the guard reads the
// whole document. The M-4 table counts the 30 gate conversations alone,
// and a fixer could quieten the terse set without moving it.
func TestQuietRunIsCaughtByTheTranscriptCount(t *testing.T) {
	prev := &Summary{Judged: 100, Bad: 20, Ratio: 0.20,
		Metrics: Metrics{Questions: 130, CatalogFilled: 90, QuestionsSeen: 220, FilledSeen: 150}}
	// The table is untouched, and half the terse questions are gone.
	quiet := &Summary{Judged: 60, Bad: 3, Ratio: 0.05,
		Metrics: Metrics{Questions: 130, CatalogFilled: 90, QuestionsSeen: 130, FilledSeen: 90}}
	if d := Compare(prev, quiet); d.Accept {
		t.Error("a run that went quiet outside the gate set was accepted")
	}
}
