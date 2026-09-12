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

	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/tune"
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
	if n := gateCount(f.Conversations); n < GateSize {
		t.Errorf("%d gate conversations, the gate needs %d", n, GateSize)
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
	if n := gateCount(f.Conversations); n < GateSize {
		t.Errorf("%d gate conversations, the gate needs %d", n, GateSize)
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

// TestPartialRunReadsItsOwnItems: a run under -only covers a part of the
// set. The catalog-only bar and the gate size bar need the whole set, so
// the verdict skips them, their suite rows are information, and the
// document says so. The same probe on a whole run fails both bars.
func TestPartialRunReadsItsOwnItems(t *testing.T) {
	cfg, err := llm.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	probe := result{
		conversation: conversation{ID: 109, Name: "a group of sets by franchise", Probe: true,
			Messages: []string{"Build the best possible deck from the sets with Marvel characters."},
			Expect:   map[string]string{"sets": "msc,msh"}},
		Turns: 1, Ready: true, Values: map[string]string{"sets": "msc,msh"},
	}
	probe.Misses = checkExpect(probe.Expect, probe.Values)
	var cov coverages
	cov.add(probe)
	file := gateFile{VerifiedAt: "2026-08-31", Conversations: make([]conversation, 108)}

	rec := evalrun.New("questions", "test")
	rec.Header.Only = "109"
	var buf bytes.Buffer
	err = write(&buf, file, []result{probe}, cov, llm.Report{Calls: 5}, cfg, "no snapshot", time.Second, rec)
	doc := buf.String()
	if err != nil || !strings.Contains(doc, "Verdict: PASS. 0 of 0 counted gate conversations") {
		t.Errorf("a partial run passes on its item bars: %v\n%s", err, doc)
	}
	if !strings.Contains(doc, "This is a partial run over `109`: 1 of 108 conversations ran.") {
		t.Errorf("the document must say it is partial:\n%s", doc)
	}
	rows := map[string]evalrun.Row{}
	for _, r := range rec.Rows {
		rows[r.Item+"/"+r.Metric] = r
	}
	if r := rows["suite/catalog_only"]; r.Kind != evalrun.KindInfo || !strings.HasSuffix(r.Detail, "not read on a partial run") {
		t.Errorf("the catalog-only suite row of a partial run is information: %+v", r)
	}
	if r := rows["suite/gate_size"]; r.Kind != evalrun.KindInfo {
		t.Errorf("the gate size suite row of a partial run is information: %+v", r)
	}
	if r := rows["suite/expectation_misses"]; r.Kind != evalrun.KindGate {
		t.Errorf("the expectation bar still holds: %+v", r)
	}
	if !rec.Gated() {
		t.Error("a partial run still holds gate rows, so a compare by hand can read it")
	}

	rec = evalrun.New("questions", "test")
	buf.Reset()
	err = write(&buf, file, []result{probe}, cov, llm.Report{Calls: 5}, cfg, "no snapshot", time.Second, rec)
	if err == nil || !strings.Contains(buf.String(), "Verdict: FAIL.") || strings.Contains(buf.String(), "partial run") {
		t.Errorf("the same probe on a whole run fails the two bars: %v\n%s", err, buf.String())
	}
}

// TestEveryCountedConversationNamesItsExpectations pins the data of
// slice 4: each counted conversation carries an expectation, every key
// is one slotValues writes, and an after-build conversation carries
// none. A probe may carry one, and its rows are information until the
// probe joins the bar (D-427, D-522, D-525).
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
		if c.HasDeck && len(c.Expect) > 0 {
			t.Errorf("%d. %s: an expectation on a conversation that starts after a build", c.ID, c.Name)
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

// TestAForbiddenRowFailsTheGate reads the "must not ask" expectation
// end to end (PR-28b, D-644). A reader who says "you already had this
// answer" becomes a conversation with the row they saw twice. The row
// that fires is a miss beside every other miss, so it fails the gate,
// the document names it, and the run records one bar for the row.
func TestAForbiddenRowFailsTheGate(t *testing.T) {
	cfg, err := llm.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	build := func(fired bool) result {
		r := result{
			conversation: conversation{ID: 109, Name: "Q1: asked again", Collection: true,
				Messages:   []string{"Bracket 3, lifegain, from my library."},
				MustNotAsk: []string{"power_commander"}},
			Turns: 1, Ready: true, Values: map[string]string{"format": "commander"},
		}
		r.Questions = []asked{{Turn: 1, Row: "theme"}}
		if fired {
			r.Questions = append(r.Questions, asked{Turn: 1, Row: "power_commander"})
		}
		r.Misses = checkNotAsked(r.MustNotAsk, r.Questions)
		return r
	}
	// The run is partial, so the catalog-only bar and the size bar read
	// nothing and the item bars stand alone (D-526).
	render := func(rs []result) (string, *evalrun.Run) {
		var cov coverages
		for _, r := range rs {
			cov.add(r)
		}
		rec := evalrun.New("questions", "test")
		rec.Header.Only = "109"
		var buf bytes.Buffer
		file := gateFile{VerifiedAt: "2026-09-09", Conversations: make([]conversation, 109)}
		_ = write(&buf, file, rs, cov, llm.Report{Calls: 1}, cfg, "no snapshot", time.Second, rec)
		return buf.String(), rec
	}

	doc, rec := render([]result{build(true)})
	if !strings.Contains(doc, "Verdict: FAIL") {
		t.Errorf("a forbidden row that fired must fail the gate:\n%s", doc)
	}
	if !strings.Contains(doc, "Q1: asked again: must_not_ask: power_commander fired at turn 1") {
		t.Errorf("the document must name the row that fired:\n%s", doc)
	}
	if !strings.Contains(doc, "Rows it must never ask: power_commander.") {
		t.Errorf("the conversation section must name the expectation:\n%s", doc)
	}
	rows := map[string]evalrun.Row{}
	for _, r := range rec.Rows {
		rows[r.Item+"/"+r.Metric] = r
	}
	if r := rows["109/never_asked_power_commander"]; r.Value != 0 || r.Kind != evalrun.KindGate {
		t.Errorf("the fired row = %+v, want a gate bar at zero", r)
	}

	// The control. Without it a bar that always fails would look like a
	// working bar (D-234).
	doc, rec = render([]result{build(false)})
	if !strings.Contains(doc, "Verdict: PASS") {
		t.Errorf("a row that never fired must pass:\n%s", doc)
	}
	rows = map[string]evalrun.Row{}
	for _, r := range rec.Rows {
		rows[r.Item+"/"+r.Metric] = r
	}
	if r := rows["109/never_asked_power_commander"]; r.Value != 1 || r.Detail != "the row never fired" {
		t.Errorf("the clean row = %+v, want a gate bar at one", r)
	}
}

// TestAMissedConversationPlaysAgain reads the reruns of D-671. A counted
// conversation with a miss plays MissReruns more times inside the run, and
// the document and the run file carry its miss rate. The first play stands,
// so the miss still fails the run and no rerun adds to the miss count. A
// clean conversation and a probe play once.
func TestAMissedConversationPlaysAgain(t *testing.T) {
	cfg, err := llm.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	build := func(id int, name string, probe bool, got string) result {
		r := result{
			conversation: conversation{ID: id, Name: name, Probe: probe,
				Messages: []string{"A blink deck."}, Expect: map[string]string{"power": "bracket 3"}},
			Turns: 1, Ready: true, Values: map[string]string{"power": got},
		}
		r.Misses = checkExpect(r.Expect, r.Values)
		return r
	}
	plays := map[int]int{}
	// The first rerun misses on another value, and the second meets it.
	replay := func(c conversation) result {
		plays[c.ID]++
		again := result{conversation: c, Turns: 1, Ready: true, Values: map[string]string{"power": "bracket 3"}}
		if plays[c.ID] == 1 {
			again.Values["power"] = ""
		}
		again.Misses = checkExpect(c.Expect, again.Values)
		return again
	}
	results := []result{
		build(1, "lifegain with a collection", false, "bracket 3"),
		build(2, "blink with a thin library", false, "bracket 2"),
		build(60, "terse: a probe that misses", true, "bracket 2"),
	}
	rerunMisses(results, replay)
	if plays[1] != 0 || plays[60] != 0 {
		t.Errorf("a clean conversation and a probe play once: %v", plays)
	}
	if plays[2] != MissReruns || len(results[1].Reruns) != MissReruns {
		t.Fatalf("the missed conversation played %d more times, want %d: %q", plays[2], MissReruns, results[1].Reruns)
	}
	if got := results[1].Reruns[0]; len(got) != 1 || got[0] != "power: want bracket 3, got none" {
		t.Errorf("play 2 keeps its own miss: %q", got)
	}
	if got := results[1].Reruns[1]; len(got) != 0 {
		t.Errorf("play 3 met every expectation: %q", got)
	}
	if missed, n := results[1].missRate(); missed != 2 || n != 3 {
		t.Errorf("miss rate = %d of %d, want 2 of 3", missed, n)
	}

	var cov coverages
	for _, r := range results {
		cov.add(r)
	}
	rec := evalrun.New("questions", "test")
	rec.Header.Only = "1,2,60"
	var buf bytes.Buffer
	file := gateFile{VerifiedAt: "2026-09-11", Conversations: make([]conversation, 108)}
	err = write(&buf, file, results, cov, llm.Report{Calls: 1}, cfg, "no snapshot", time.Second, rec)
	doc := buf.String()
	if err == nil || !strings.Contains(doc, "Verdict: FAIL") {
		t.Errorf("the first play stands, so the miss still fails the run: %v\n%s", err, doc)
	}
	if !strings.Contains(doc, "1 slots ended on a value other than the one the conversation expects") {
		t.Errorf("a rerun adds no miss to the count:\n%s", doc)
	}
	if !strings.Contains(doc, "played 2 more times inside the run, and the first play stands (D-671): blink with a thin library missed in 2 of 3 plays.") {
		t.Errorf("the document must carry the miss rate beside the miss:\n%s", doc)
	}
	if !strings.Contains(doc, "Reruns (D-671): missed in 2 of 3 plays. Play 2: power: want bracket 3, got none. Play 3: every one met.") {
		t.Errorf("the conversation section must name each rerun:\n%s", doc)
	}
	rows := map[string]evalrun.Row{}
	for _, r := range rec.Rows {
		rows[r.Item+"/"+r.Metric] = r
	}
	if r, ok := rows["2/miss_rate"]; !ok || r.Kind != evalrun.KindInfo || r.Value < 0.666 || r.Value > 0.667 || r.Detail != "missed in 2 of 3 plays" {
		t.Errorf("the run file carries the miss rate as information: %+v", r)
	}
	if _, ok := rows["1/miss_rate"]; ok {
		t.Error("a clean conversation writes no miss rate")
	}
	if r := rows["suite/expectation_misses"]; r.Value != 1 {
		t.Errorf("the suite bar reads the first play alone: %+v", r)
	}

	// A play that fails reads its error as a miss of that play.
	failing := []result{build(2, "blink with a thin library", false, "bracket 2")}
	rerunMisses(failing, func(c conversation) result {
		return result{conversation: c, Err: fmt.Errorf("turn 1: provider down")}
	})
	if got := failing[0].Reruns[0]; len(got) != 1 || got[0] != "error: turn 1: provider down" {
		t.Errorf("a failed play = %q, want its error as a miss", got)
	}
	if missed, n := failing[0].missRate(); missed != 3 || n != 3 {
		t.Errorf("miss rate = %d of %d, want 3 of 3", missed, n)
	}
}
