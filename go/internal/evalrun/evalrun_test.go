package evalrun

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

func sample() *Run {
	r := New("decks", "pr8-deck-gate-run15")
	r.Header.Date = "2026-09-04"
	r.Header.Commit = "abc1234"
	r.SetRoles(&llm.Config{Roles: map[llm.Role]llm.RoleSpec{
		llm.RoleGenerate: {Provider: "openai", Model: "gpt-5.6-terra", Effort: "medium"},
		llm.RoleJudge:    {Provider: "anthropic", Model: "claude-opus-5"},
	}}, llm.RoleGenerate, llm.RoleJudge, llm.RoleEval)
	r.SetSnapshot(time.Date(2026, 9, 3, 12, 0, 0, 0, time.UTC))
	r.Header.Prompts["generate"] = 12
	r.Header.Versions["precons"] = "5.3.0+20260903"
	r.Gate("17", "blocks", 1, "deck_size")
	r.Info("17", "repaired", 1, "")
	cost := 2.5432
	r.Finish(llm.Report{Calls: 62, CostUSD: &cost}, 2081*time.Second, "FAIL")
	return r
}

// TestRoundTrip: what Write writes, Read reads back whole.
func TestRoundTrip(t *testing.T) {
	r := sample()
	var buf bytes.Buffer
	if err := Write(&buf, r); err != nil {
		t.Fatal(err)
	}
	if lines := strings.Count(buf.String(), "\n"); lines != 3 {
		t.Errorf("wrote %d lines, want the header and two rows", lines)
	}
	got, err := Read(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if got.Header.Suite != "decks" || got.Header.RunID != "pr8-deck-gate-run15" || got.Header.Verdict != "FAIL" {
		t.Errorf("header = %+v", got.Header)
	}
	if got.Header.Roles["generate"].Effort != "medium" || got.Header.Roles["judge"].Provider != "anthropic" {
		t.Errorf("roles = %+v", got.Header.Roles)
	}
	if _, ok := got.Header.Roles["eval"]; ok {
		t.Error("a role the config lacks must not appear")
	}
	if got.Header.Snapshot != "2026-09-03" || got.Header.Prompts["generate"] != 12 || got.Header.Calls != 62 || *got.Header.CostUSD != 2.5432 {
		t.Errorf("facts = %+v", got.Header)
	}
	if len(got.Rows) != 2 || got.Rows[0].Kind != KindGate || got.Rows[1].Kind != KindInfo || got.Rows[0].Detail != "deck_size" {
		t.Errorf("rows = %+v", got.Rows)
	}
	if !got.Gated() {
		t.Error("a run with a gate row is gated")
	}
	if (&Run{}).Gated() {
		t.Error("a run with no row is not evaluated")
	}
}

// TestReadRefusesAHeaderlessFile: the first line must name the suite.
func TestReadRefusesAHeaderlessFile(t *testing.T) {
	if _, err := Read(strings.NewReader(`{"item":"1","metric":"blocks","value":0,"kind":"gate"}` + "\n")); err == nil {
		t.Error("a file that starts with a row read as a run")
	}
	if _, err := Read(strings.NewReader("")); err == nil {
		t.Error("an empty file read as a run")
	}
}

// TestWriteFileNeverOverwrites is D-65 for the run file.
func TestWriteFileNeverOverwrites(t *testing.T) {
	path := filepath.Join(t.TempDir(), "eval", "run1.jsonl")
	if err := WriteFile(path, sample()); err != nil {
		t.Fatal(err)
	}
	if err := WriteFile(path, sample()); err == nil {
		t.Error("the second write did not refuse")
	}
	got, err := ReadFile(path)
	if err != nil || len(got.Rows) != 2 {
		t.Errorf("read back %v, %v", got, err)
	}
	if err := WriteFile("", sample()); err != nil {
		t.Errorf("an empty path writes nothing and fails nothing: %v", err)
	}
	if RunID("/a/b/pr8-deck-gate-run15.jsonl") != "pr8-deck-gate-run15" {
		t.Error("the run id is the file stem")
	}
}

// TestMarkdownIsTheSameBlockEverywhere pins the fingerprint lines.
func TestMarkdownIsTheSameBlockEverywhere(t *testing.T) {
	var buf bytes.Buffer
	sample().Markdown(&buf)
	doc := buf.String()
	for _, want := range []string{
		"## Run\n\n",
		"- Suite `decks`, run `pr8-deck-gate-run15`, on 2026-09-04, commit `abc1234`.\n",
		"- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic).\n",
		"- Versions: card snapshot 2026-09-03, generate prompt version 12, precons `5.3.0+20260903`.\n",
		"- Calls: 62. Cost: $2.5432. Time: 2081 seconds.\n",
	} {
		if !strings.Contains(doc, want) {
			t.Errorf("the block lacks %q:\n%s", want, doc)
		}
	}
	var free bytes.Buffer
	New("quality", "").Markdown(&free)
	if !strings.Contains(free.String(), "run `unnamed`") || !strings.Contains(free.String(), "Roles: none, no provider call.") || !strings.Contains(free.String(), "Cost: unpriced.") {
		t.Errorf("a free run reads wrong:\n%s", free.String())
	}
}
