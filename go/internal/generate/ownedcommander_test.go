package generate

import (
	"fmt"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
)

// commanderCard is a card the rules engine reads as legal in Commander.
func commanderCard(oid, name string, lead bool) *mtgv1.Card {
	return &mtgv1.Card{
		OracleId:       oid,
		Name:           name,
		CardTypes:      []string{"Creature"},
		ManaValue:      2,
		Legalities:     map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
		CanBeCommander: lead,
	}
}

// TestACommanderFindingBuysNoRepairTurn is F-157 and D-762. The reader
// named a commander the collection does not cover, and the owned-only
// check blocked it (D-37, D-226). The block bought a repair turn of 60
// to 90 seconds that can not fix it, because assemble takes the
// commander from the request on every pass.
func TestACommanderFindingBuysNoRepairTurn(t *testing.T) {
	// The shape of the three M-18 builds: one ownership block on the
	// commander, and the profile findings of bracket 5.
	v := &mtgv1.ValidationResult{Findings: []*mtgv1.Finding{
		{
			Code:     rules.CodeNotOwned,
			Severity: mtgv1.Severity_SEVERITY_BLOCK,
			OracleId: "o-grima",
			Message:  "Gríma, Saruman's Footman: the deck needs 1, the collection has 0",
		},
		{Code: profile.CodeOffBand, Severity: mtgv1.Severity_SEVERITY_WARN, Message: "the tutor count is 0"},
		{Code: profile.CodeOffBand, Severity: mtgv1.Severity_SEVERITY_WARN, Message: "the fast mana count is 3"},
		{Code: profile.CodeOffBand, Severity: mtgv1.Severity_SEVERITY_WARN, Message: "the Game Changer count is 1"},
	}}
	all := repairable(v)
	if len(all) != 4 {
		t.Fatalf("the repair reads %d findings, want all four", len(all))
	}
	got := fixable(all, []string{"o-grima"})
	if len(got) != 3 {
		t.Fatalf("the repair input holds %d findings, want the three profile rows", len(got))
	}
	for _, f := range got {
		if f.GetOracleId() == "o-grima" {
			t.Errorf("the repair input holds a finding on the commander: %s", f.GetMessage())
		}
	}
	// The decision of the build loop: the findings that stand are
	// profile findings alone, so no repair turn runs.
	if profileOnly(all) {
		t.Error("the block on the commander must read as more than a profile finding")
	}
	if !profileOnly(got) {
		t.Error("with the commander finding out, the profile findings stand alone and buy no repair turn")
	}
	// A finding on another card still buys the turn, and a deck with no
	// commander keeps every finding.
	other := fixable(all, []string{"o-other"})
	if len(other) != 4 || profileOnly(other) {
		t.Errorf("a finding on another card reads %d findings, want all four", len(other))
	}
	if len(fixable(all, nil)) != 4 {
		t.Error("a deck with no commander keeps every finding")
	}
}

// TestAnUnownedCommanderRunsNoRepairTurn builds the case end to end. The
// deck goes out with its block (D-226, D-300), and the build spends one
// model call and no repair call.
func TestAnUnownedCommanderRunsNoRepairTurn(t *testing.T) {
	lead := commanderCard("o-lead", "Test Commander", true)
	cards := []*mtgv1.Card{lead}
	src := source{lead.GetOracleId(): lead}
	owned := map[string]int32{}
	out := deckOut{Summary: "an owned-only deck"}
	for i := 0; i < 99; i++ {
		c := commanderCard(fmt.Sprintf("o-%02d", i), fmt.Sprintf("Owned Card %02d", i), false)
		cards = append(cards, c)
		src[c.GetOracleId()] = c
		owned[c.GetOracleId()] = 1
		out.Cards = append(out.Cards, Entry{Name: c.GetName(), Count: 1, Role: "threat", Reason: "fills the deck"})
	}
	// Two steps: the build takes the first, and a repair turn would take
	// the second. The call count is the measurement.
	b, sc := testBuilderWith(t, src, step(t, out), step(t, out))
	req := Request{
		SessionID:    "s-f157",
		Format:       mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Plan:         "a deck from the cards the reader owns",
		Pool:         NewPool(cards, owned),
		PoolRule:     mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		OracleCounts: owned,
		Commanders:   []string{lead.GetOracleId()},
		Limits:       "100 cards, one copy of each name.",
		LegalityAsOf: "2026-09-04",
	}
	got, err := b.Build(t.Context(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(sc.Calls) != 1 {
		t.Fatalf("provider calls = %d, want the build alone and no repair turn", len(sc.Calls))
	}
	if got.Repaired {
		t.Error("no repair turn runs for a finding the model can not fix")
	}
	var block *mtgv1.Finding
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == rules.CodeNotOwned {
			block = f
		}
	}
	if block == nil || block.GetOracleId() != lead.GetOracleId() {
		t.Fatalf("the deck must carry the ownership block on its commander: %v",
			got.Deck.GetValidation().GetFindings())
	}
	if block.GetSeverity() != mtgv1.Severity_SEVERITY_BLOCK {
		t.Errorf("the ownership finding reads %s, want a block in owned-only", block.GetSeverity())
	}
	if len(got.Deck.GetCommanderOracleIds()) != 1 {
		t.Errorf("the deck carries %d commanders, want the one the reader named", len(got.Deck.GetCommanderOracleIds()))
	}
}
