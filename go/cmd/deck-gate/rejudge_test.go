package main

import (
	"bytes"
	"path/filepath"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/generate"
)

// TestReadStoredReadsWriteDeckBack is D-789: the rejudge lane reads the
// commander, the summary, and the deck list that writeDeck wrote, with
// each count and role. The cards are test fixtures, and one partner
// name holds a comma of its own.
func TestReadStoredReadsWriteDeckBack(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-a", Name: "Partner, the First"},
		{OracleId: "o-b", Name: "The Second Partner"},
		{OracleId: "o-land", Name: "A Plain Land"},
		{OracleId: "o-kill", Name: "A Removal Spell"},
	}, nil, nil, time.Time{})
	r := result{
		prompt: prompt{ID: 7, Name: "a partner deck", Format: "commander"},
		deck: &mtgv1.Deck{
			Summary:            "A plan in one line, with a | bar.",
			CommanderOracleIds: []string{"o-a", "o-b"},
			Cards: []*mtgv1.DeckCard{
				{OracleId: "o-land", Name: "A Plain Land", Count: 36, Role: mtgv1.CardRole_CARD_ROLE_LAND, Reason: "mana | and more"},
				{OracleId: "o-kill", Name: "A Removal Spell", Count: 1, Role: mtgv1.CardRole_CARD_ROLE_REMOVAL},
			},
		},
	}
	var doc bytes.Buffer
	writeDeck(&doc, r, idx)
	stored, err := readStored(&doc, idx)
	if err != nil {
		t.Fatal(err)
	}
	st := stored[7]
	if st == nil || st.summary != r.deck.GetSummary() || len(st.cards) != 2 {
		t.Fatalf("stored = %+v", st)
	}
	if c := st.cards[0]; c.GetOracleId() != "o-land" || c.GetCount() != 36 || c.GetRole() != mtgv1.CardRole_CARD_ROLE_LAND {
		t.Errorf("first card = %+v", c)
	}
	if c := st.cards[1]; c.GetRole() != mtgv1.CardRole_CARD_ROLE_REMOVAL {
		t.Errorf("second card = %+v", c)
	}
	ids, ok := commanderIDs(st.commander, idx)
	if !ok || len(ids) != 2 || ids[0] != "o-a" || ids[1] != "o-b" {
		t.Errorf("commanders of %q = %v, %v", st.commander, ids, ok)
	}
	if _, err := readStored(bytes.NewBufferString("### 1. x\n\n- 1 A Card Nobody Printed | land | a reason\n"), idx); err == nil {
		t.Error("a name the index does not know reads as a deck with a hole")
	}
}

// TestCopyBuildRowsDropsTheJudgeRows is D-789: the lane keeps every row
// the build wrote, drops each judge row, and reads the build bars.
func TestCopyBuildRowsDropsTheJudgeRows(t *testing.T) {
	src := evalrun.New("decks", "src")
	src.Gate("1", "built", 1, "")
	src.Gate("1", "blocks", 0, "")
	src.Gate("1", "false_rules", 1, "")
	src.Info("1", "rules_claims", 2, "")
	src.Info("1", "plan_theme_fit", 0.5, "")
	src.Info("1", "plan_score", 0.7, "")
	src.Info("1", "deck_cost", 200, "")
	run := evalrun.New("decks", "next")
	if !copyBuildRows(run, src) {
		t.Error("a clean build reads as a failed build bar")
	}
	if len(run.Rows) != 3 {
		t.Fatalf("rows = %+v, want built, blocks, and deck_cost", run.Rows)
	}
	for _, row := range run.Rows {
		if isJudgeMetric(row.Metric) {
			t.Errorf("a judge row stayed: %+v", row)
		}
	}
	src.Gate("2", "blocks", 1, "SIZE")
	if copyBuildRows(evalrun.New("decks", "x"), src) {
		t.Error("a block of the source reads as a pass")
	}
}

// TestKeptRowsKeepsTheAnsweredDecks is D-789: a deck both judges answered
// keeps its rows, and a deck with a judge error is judged again. A run
// that rejudged another source, or read another judge version, is
// refused.
func TestKeptRowsKeepsTheAnsweredDecks(t *testing.T) {
	src := evalrun.New("decks", "src")
	src.Header.RunID = "src"
	kept := evalrun.New("decks", "kept")
	kept.Header.Versions["rejudge_of"] = "src"
	kept.Header.Prompts["plan_rubric"] = generate.PlanRubricVersion
	kept.Header.Prompts["summary_judge"] = generate.SummaryJudgeVersion
	kept.Gate("1", "false_rules", 0, "")
	kept.Info("1", "plan_score", 0.75, "")
	kept.Gate("2", "judge_error", 1, "529")
	kept.Info("2", "plan_score", 0.5, "")
	kept.Gate("3", "false_rules", 0, "")
	kept.Info("3", "plan_judge_error", 1, "529")
	kept.Gate("3", "built", 1, "")
	path := filepath.Join(t.TempDir(), "kept.jsonl")
	if err := evalrun.WriteFile(path, kept); err != nil {
		t.Fatal(err)
	}
	got, err := keptRows(path, src)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || len(got["1"]) != 2 {
		t.Errorf("kept = %+v, want the two rows of deck 1 alone", got)
	}
	src.Header.RunID = "another"
	if _, err := keptRows(path, src); err == nil {
		t.Error("a rejudge of another run is kept")
	}
}
