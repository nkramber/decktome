package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/tune"
)

func load(t *testing.T) gateFile {
	t.Helper()
	var f gateFile
	if err := json.Unmarshal(conversationsJSON, &f); err != nil {
		t.Fatalf("conversations.json: %v", err)
	}
	return f
}

func TestConversationsFile(t *testing.T) {
	f := load(t)
	if f.VerifiedAt == "" {
		t.Error("conversations.json has no verified_at")
	}
	if f.Note == "" {
		t.Error("conversations.json has no note")
	}
	if n := gateCount(f.Conversations); n < questions.MinGateSize {
		t.Errorf("%d gate conversations, the gate needs %d", n, questions.MinGateSize)
	}
	seenID := map[int]bool{}
	seenName := map[string]bool{}
	lastID := 0
	for i, c := range f.Conversations {
		switch {
		// The ids climb and never repeat. A gap is allowed: 67 left the
		// set and the numbers stay, because the holdout split and the
		// paired comparison key on them (A-9, D-134).
		case c.ID <= lastID:
			t.Errorf("conversation %d has id %d after id %d", i+1, c.ID, lastID)
		case seenID[c.ID]:
			t.Errorf("duplicate id %d", c.ID)
		case c.Name == "":
			t.Errorf("conversation %d has no name", c.ID)
		case seenName[c.Name]:
			t.Errorf("duplicate name %q", c.Name)
		case len(c.Messages) == 0:
			t.Errorf("conversation %q has no message", c.Name)
		case len(c.Messages) > 4:
			t.Errorf("conversation %q has %d messages, the gate allows four turns", c.Name, len(c.Messages))
		}
		seenID[c.ID], seenName[c.Name] = true, true
		lastID = c.ID
		for j, m := range c.Messages {
			if strings.TrimSpace(m) == "" {
				t.Errorf("conversation %q message %d is empty", c.Name, j+1)
			}
		}
	}
}

// TestRunNeedsApproval keeps the command from spending money by accident.
func TestRunNeedsApproval(t *testing.T) {
	t.Setenv("QUESTIONS_GATE", "")
	err := run("", 0, "", "", io.Discard)
	if err == nil || !strings.Contains(err.Error(), "QUESTIONS_GATE=1") {
		t.Errorf("err = %v, want a refusal without QUESTIONS_GATE=1", err)
	}
}

// TestSizeTestCountsLikeRun is T-20. The size test and run() count the
// gate conversations the same way, and both leave the probes out.
func TestSizeTestCountsLikeRun(t *testing.T) {
	f := load(t)
	if n := gateCount(f.Conversations); n < questions.MinGateSize {
		t.Errorf("%d gate conversations, the gate needs %d", n, questions.MinGateSize)
	}
	cases := []struct {
		name  string
		convs []conversation
		want  int
	}{
		{"probes leave the count", []conversation{{Probe: true}, {}, {HasDeck: true}}, 2},
		{"no conversations", nil, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := gateCount(tc.convs); got != tc.want {
				t.Errorf("gateCount = %d, want %d", got, tc.want)
			}
		})
	}
}

// TestConversationSetOfA9 pins the conversation set rule (A-9): no
// two conversations share a script, and every after-build conversation
// carries the flag.
func TestConversationSetOfA9(t *testing.T) {
	f := load(t)
	scripts := map[string]int{}
	for _, c := range f.Conversations {
		key := strings.Join(c.Messages, "\x1f")
		if prev, dup := scripts[key]; dup {
			t.Errorf("conversation %d repeats the script of %d", c.ID, prev)
		}
		scripts[key] = c.ID
	}
	for _, id := range []int{11, 12, 30} {
		found := false
		for _, c := range f.Conversations {
			if c.ID == id {
				found = true
				if !c.HasDeck {
					t.Errorf("conversation %d starts after a build and lacks has_deck", id)
				}
			}
		}
		if !found {
			t.Errorf("conversation %d is missing", id)
		}
	}
	for _, c := range f.Conversations {
		if c.ID == 67 {
			t.Error("conversation 67 was a copy of 38 and should be gone")
		}
	}
}

// TestDocumentRoundTrip is the round-trip lock: the document this
// command writes is what tune.ReadRun reads. A changed heading or label
// would leave the loop reading nothing, and nothing said so before.
func TestDocumentRoundTrip(t *testing.T) {
	cfg, err := llm.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	q := func(turn int, slot, row string, filled, invented bool, text string) asked {
		return asked{Turn: turn, Slot: slot, Row: row, Source: source(invented), Fit: 0.9,
			Text: text, Filled: filled, Invented: invented, Catalog: "the catalog text"}
	}
	results := []result{
		{
			conversation: conversation{ID: 1, Name: "lifegain with a collection", Collection: true,
				Messages: []string{"Build me a lifegain deck.", "Commander."}},
			Questions: []asked{q(1, "format", "format", true, false, "What format?"),
				q(2, "colors", "colors", true, true, "Which colors?")},
			Turns: 2, Ready: true,
		},
		{
			conversation: conversation{ID: 11, Name: "another version after a build", Collection: true,
				HasDeck: true, Messages: []string{"Give me another version."}},
			Questions: []asked{q(1, "theme", "theme", false, false, "Same plan?")},
			Turns:     1,
		},
		{
			conversation: conversation{ID: 38, Name: "everything in one message", Probe: true,
				Messages: []string{"Commander, bracket 3, lifegain."}},
			Turns: 1, Ready: true, Unanswered: []string{"colors (never asked)"}, Premature: true,
		},
		{
			conversation: conversation{ID: 5, Name: "an error", Messages: []string{"Hello.", "Never sent."}},
			Turns:        1, Err: fmt.Errorf("turn 1: the provider timed out"),
		},
	}
	var cov coverages
	for _, r := range results {
		cov.add(r)
	}
	var buf bytes.Buffer
	_ = write(&buf, gateFile{VerifiedAt: "2026-08-28"}, results, cov, llm.Report{Calls: 3}, cfg, "no snapshot", time.Second, evalrun.New("questions", "test"))
	path := filepath.Join(t.TempDir(), "pr7-question-gate-roundtrip.md")
	if err := os.WriteFile(path, buf.Bytes(), 0o600); err != nil {
		t.Fatal(err)
	}
	run, err := tune.ReadRun(path)
	if err != nil {
		t.Fatalf("ReadRun: %v\n%s", err, buf.String())
	}
	if len(run.Conversations) != 4 {
		t.Fatalf("ReadRun gave %d conversations, want 4:\n%s", len(run.Conversations), buf.String())
	}
	got := run.Conversations
	checks := []struct {
		name string
		ok   bool
	}{
		{"names carry the id", got[0].Name == "1. lifegain with a collection"},
		{"the collection flag", got[0].Collection && !got[2].Collection},
		{"both messages", len(got[0].Messages) == 2},
		{"both questions with their rows", len(got[0].Questions) == 2 && got[0].Questions[1].Row == "colors" && got[0].Questions[1].Turn == 2},
		{"the invented flag and the replaced text", got[0].Questions[1].Invented && got[0].Questions[1].Catalog == "the catalog text"},
		{"the after-build suffix", got[1].AfterBuild && !got[1].Probe},
		{"the probe suffix", got[2].Probe},
		{"the premature flag and its slots", got[2].Premature && got[2].Unanswered == "colors (never asked)"},
		{"the failed marker", got[3].Errored() && got[3].Failed == "turn 1: the provider timed out"},
		{"the messages that went out alone", len(got[3].Messages) == 1},
		{"the verdict", run.Verdict == "FAIL"},
		{"the premature count", run.Metrics.Premature == 1},
	}
	for _, c := range checks {
		if !c.ok {
			t.Errorf("%s: not read back\n%+v", c.name, got)
		}
	}
	// The M-4 table and the recount agree, with the after-build
	// conversation out of both (A-9).
	if m := tune.MetricsOf(run); m.Questions != run.Metrics.Questions || m.CatalogOnly != run.Metrics.CatalogOnly ||
		m.Invented != run.Metrics.Invented || m.CatalogFilled != run.Metrics.CatalogFilled {
		t.Errorf("recount %+v, table %+v", m, run.Metrics)
	}
	if run.Metrics.CatalogOnly != 1 || run.Metrics.Questions != 2 {
		t.Errorf("table = %+v, want the one counted gate conversation", run.Metrics)
	}
	if !strings.Contains(buf.String(), "1 of 2 counted gate conversations") {
		t.Errorf("the verdict line does not read the counted set:\n%s", buf.String())
	}
}

// TestExpectationsAreTheFourthBar: a counted conversation whose slots
// missed an expectation fails the gate, the document names the miss,
// and the run holds one gate row per expected key. A probe's rows are
// information, and a conversation with no misses reads as met.
func TestExpectationsAreTheFourthBar(t *testing.T) {
	cfg, err := llm.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	met := result{
		conversation: conversation{ID: 1, Name: "lifegain with a collection", Collection: true,
			Messages: []string{"Build me a lifegain deck."}, Expect: map[string]string{"format": "commander", "theme": "lifegain"}},
		Turns: 1, Ready: true, Values: map[string]string{"format": "commander", "theme": "a lifegain deck"},
	}
	met.Misses = checkExpect(met.Expect, met.Values)
	missed := result{
		conversation: conversation{ID: 2, Name: "blink with a thin library", Collection: true,
			Messages: []string{"A blink deck."}, Expect: map[string]string{"power": "bracket 3"}},
		Turns: 1, Ready: true, Values: map[string]string{"power": "bracket 2"},
	}
	missed.Misses = checkExpect(missed.Expect, missed.Values)
	probe := result{
		conversation: conversation{ID: 60, Name: "terse: a format we do not build", Probe: true,
			Messages: []string{"Brawl."}, Expect: map[string]string{"format": "commander"}},
		Turns: 1, Ready: true, Values: map[string]string{"format": "modern"},
	}
	probe.Misses = checkExpect(probe.Expect, probe.Values)
	render := func(rs []result) (string, *evalrun.Run) {
		var cov coverages
		for _, r := range rs {
			cov.add(r)
		}
		rec := evalrun.New("questions", "test")
		var buf bytes.Buffer
		_ = write(&buf, gateFile{VerifiedAt: "2026-08-31"}, rs, cov, llm.Report{Calls: 1}, cfg, "no snapshot", time.Second, rec)
		return buf.String(), rec
	}
	doc, rec := render([]result{met, missed, probe})
	if !strings.Contains(doc, "Verdict: FAIL") || !strings.Contains(doc, "1 slots ended on a value other than the one the conversation expects") {
		t.Errorf("a counted miss must fail the gate:\n%s", doc)
	}
	if !strings.Contains(doc, "blink with a thin library: power: want bracket 3, got bracket 2") {
		t.Errorf("the document must name the miss:\n%s", doc)
	}
	if !strings.Contains(doc, "Expected slots: every one met.") || !strings.Contains(doc, "Expected slots: 1 misses. power: want bracket 3, got bracket 2.") {
		t.Errorf("the per-conversation lines:\n%s", doc)
	}
	rows := map[string]evalrun.Row{}
	for _, r := range rec.Rows {
		rows[r.Item+"/"+r.Metric] = r
	}
	if r := rows["1/slot_theme"]; r.Value != 1 || r.Kind != evalrun.KindGate || r.Detail != "want lifegain, got a lifegain deck" {
		t.Errorf("met row = %+v", r)
	}
	if r := rows["2/slot_power"]; r.Value != 0 || r.Kind != evalrun.KindGate {
		t.Errorf("missed row = %+v", r)
	}
	if r := rows["60/slot_format"]; r.Value != 0 || r.Kind != evalrun.KindInfo {
		t.Errorf("a probe's row is information: %+v", r)
	}
	if r := rows["suite/expectation_misses"]; r.Value != 1 || r.Kind != evalrun.KindGate {
		t.Errorf("suite row = %+v", r)
	}
	doc, _ = render([]result{met, probe})
	if strings.Contains(doc, "which fails the gate (PR-15)") || !strings.Contains(doc, "Every slot of the 1 counted conversations that name expectations ended on the expected value") {
		t.Errorf("a probe's miss moves no verdict:\n%s", doc)
	}
}

// TestEveryCountedConversationNamesItsExpectations pins the data of
// slice 4: each counted conversation carries an expectation, every key
// is one slotValues writes, and no probe or after-build conversation
// carries one yet (OQ-61).
func TestEveryCountedConversationNamesItsExpectations(t *testing.T) {
	f := load(t)
	known := map[string]bool{}
	for _, k := range expectKeys {
		known[k] = true
	}
	for _, c := range f.Conversations {
		counted := !c.Probe && !c.HasDeck
		if counted && len(c.Expect) == 0 {
			t.Errorf("%d. %s: a counted conversation with no expectation", c.ID, c.Name)
		}
		if !counted && len(c.Expect) > 0 {
			t.Errorf("%d. %s: an expectation on a conversation the bar does not count", c.ID, c.Name)
		}
		for k, v := range c.Expect {
			if !known[k] {
				t.Errorf("%d. %s: expectation key %q is not one slotValues writes", c.ID, c.Name, k)
			}
			if strings.TrimSpace(v) == "" {
				t.Errorf("%d. %s: expectation %q is empty", c.ID, c.Name, k)
			}
		}
	}
}
