package generate

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// testBuilderWith is testBuilder over a card source of the test's own.
func testBuilderWith(t *testing.T, src source, steps ...llm.Step) (*Builder, *llm.Script) {
	t.Helper()
	sc := llm.NewScript(steps...)
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	cfg, err := rules.Load()
	if err != nil {
		t.Fatalf("rules: %v", err)
	}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewBuilder(c, cfg, src, log), sc
}

// housePool is a pool of n distinct cards, enough for a 60-card house
// deck with four copies of each of fifteen names. A house format checks
// the size and the copies and no legality (D-3), so a made-up card
// passes the referee.
func housePool(n int) (*Pool, source) {
	var list []*mtgv1.Card
	src := source{}
	for i := 0; i < n; i++ {
		c := card(fmt.Sprintf("o-%02d", i), fmt.Sprintf("House Card %02d", i))
		list = append(list, c)
		src[c.GetOracleId()] = c
	}
	return NewPool(list, nil), src
}

func houseDeck(names int) deckOut {
	out := deckOut{Summary: "a house deck"}
	for i := 0; i < names; i++ {
		out.Cards = append(out.Cards, Entry{Name: fmt.Sprintf("House Card %02d", i), Count: 4, Role: "threat", Reason: "fills the deck"})
	}
	return out
}

func houseRequest(pool *Pool) Request {
	return Request{
		SessionID: "s-house", Format: mtgv1.FormatId_FORMAT_ID_HOUSE, HouseRules: "any card, no ban list",
		Plan: "a house deck", Pool: pool, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Limits: "60 cards minimum, four copies per name.", LegalityAsOf: "2026-09-03",
	}
}

// TestPreconShareCountsThePoolAlone is the finding of deck gate run 14:
// the bracket cut of D-468 drops a precon card the power level forbids,
// and the share rule then asks for a count the pool can not meet. The
// count reads the precon names the pool still holds.
func TestPreconShareCountsThePoolAlone(t *testing.T) {
	pool, src := housePool(20)
	req := houseRequest(pool)
	req.Precon = "Test Precon"
	for i := 0; i < 10; i++ {
		req.PreconOracleIDs = append(req.PreconOracleIDs, fmt.Sprintf("o-%02d", i))
	}
	// Two precon cards the cut dropped, and one basic land.
	req.PreconOracleIDs = append(req.PreconOracleIDs, "o-cut-1", "o-cut-2", "o-plains")
	src["o-cut-1"] = card("o-cut-1", "Cut One")
	src["o-cut-2"] = card("o-cut-2", "Cut Two")
	src["o-plains"] = basic("o-plains", "Plains")
	in := preconNonbasics(req, src)
	if len(in) != 10 {
		t.Fatalf("precon nonbasics = %d, want the 10 the pool holds", len(in))
	}
	if in["o-cut-1"] || in["o-plains"] {
		t.Error("a card the pool lacks, or a basic land, counted")
	}
	// With no pool the count reads the index, as the gate document does.
	req.Pool = nil
	if got := len(preconNonbasics(req, src)); got != 12 {
		t.Errorf("with no pool = %d, want 12: every nonbasic of the precon", got)
	}
}

// TestBuildKeepsTheLegalDeckWhenTheRepairFails is deck gate run 14: the
// repair turn gave up on a precon upgrade and answered no cards, and the
// reader got an empty deck with a block. The deck before the repair
// stands, with its warnings and a note.
func TestBuildKeepsTheLegalDeckWhenTheRepairFails(t *testing.T) {
	pool, src := housePool(20)
	req := houseRequest(pool)
	req.Precon = "Test Precon"
	for i := 0; i < 20; i++ {
		req.PreconOracleIDs = append(req.PreconOracleIDs, fmt.Sprintf("o-%02d", i))
	}
	// The first deck holds 15 of the 20 precon names, under the share
	// the rule asks for, so the repair turn runs. The repair answers no
	// cards at all.
	b, sc := testBuilderWith(t, src, step(t, houseDeck(15)), step(t, deckOut{Summary: "no legal deck"}))
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(sc.Calls) != 2 {
		t.Fatalf("provider calls = %d, want the build and one repair", len(sc.Calls))
	}
	if n := len(got.Deck.GetCards()); n != 15 {
		t.Fatalf("the deck holds %d entries, want the 15 of the first pass", n)
	}
	if !got.Deck.GetValidation().GetPassed() {
		t.Errorf("the first deck must pass the block checks: %v", got.Deck.GetValidation().GetFindings())
	}
	if !got.Repaired {
		t.Error("the repair turn ran, and the result must say so")
	}
	var kept, share bool
	for _, f := range got.Deck.GetValidation().GetFindings() {
		switch f.GetCode() {
		case CodeRepairKept:
			kept = f.GetSeverity() == mtgv1.Severity_SEVERITY_INFO
		case CodePreconShare:
			share = true
		}
	}
	if !kept {
		t.Errorf("no %s note on the deck: %v", CodeRepairKept, got.Deck.GetValidation().GetFindings())
	}
	if !share {
		t.Error("the precon share warning must stay on the deck the reader gets")
	}
}

// TestWorseRepairReadsTheBlocks: a repair with a miss or a block is
// worse than a clean pass, and never worse than a pass that failed too.
func TestWorseRepairReadsTheBlocks(t *testing.T) {
	clean := pass{deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{Passed: true}}}
	blocked := pass{deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{Passed: false}}}
	missed := pass{deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{Passed: true}}, misses: []Miss{{Name: "x"}}}
	if !worseRepair(clean, blocked) || !worseRepair(clean, missed) {
		t.Error("a block or a miss after a clean pass is worse")
	}
	if worseRepair(blocked, blocked) || worseRepair(missed, clean) || worseRepair(clean, clean) {
		t.Error("a repair of a failed pass, or a clean repair, is never worse")
	}
}
