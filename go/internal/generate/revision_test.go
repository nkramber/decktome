package generate

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

type cardMap map[string]*mtgv1.Card

func (m cardMap) ByOracleID(id string) (*mtgv1.Card, bool) { c, ok := m[id]; return c, ok }

func TestCheckRevision(t *testing.T) {
	cards := cardMap{
		"o-serenity": {OracleId: "o-serenity", Name: "Angel of Serenity", ManaValue: 7, CardTypes: []string{"Creature"}},
		"o-plains":   {OracleId: "o-plains", Name: "Plains", ManaValue: 0, CardTypes: []string{"Land"}},
		"o-lyra":     {OracleId: "o-lyra", Name: "Lyra Dawnbringer", ManaValue: 5, CardTypes: []string{"Creature"}},
	}
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-serenity", Name: "Angel of Serenity", Count: 1},
		{OracleId: "o-plains", Name: "Plains", Count: 24},
	}}
	rev := &Revision{Remove: []string{"angel of serenity"}, Keep: []string{"Lyra Dawnbringer"}, MaxManaValue: 5}
	got := CheckRevision(deck, rev, cards)
	codes := map[string]mtgv1.Severity{}
	for _, f := range got {
		codes[f.GetCode()] = f.GetSeverity()
	}
	if codes[CodeRevisionRemovedPresent] != mtgv1.Severity_SEVERITY_BLOCK {
		t.Errorf("removed present: %v", got)
	}
	if codes[CodeRevisionKeptMissing] != mtgv1.Severity_SEVERITY_BLOCK {
		t.Errorf("kept missing: %v", got)
	}
	if codes[CodeRevisionOverManaValue] != mtgv1.Severity_SEVERITY_WARN {
		t.Errorf("over mana value: %v", got)
	}
	for _, f := range got {
		if f.GetCode() == CodeRevisionOverManaValue && !strings.Contains(f.GetMessage(), "Angel of Serenity (7)") {
			t.Errorf("message = %q", f.GetMessage())
		}
	}
	// A clean deck has no finding.
	clean := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-lyra", Name: "Lyra Dawnbringer", Count: 2}, {OracleId: "o-plains", Name: "Plains", Count: 24}}}
	if got := CheckRevision(clean, rev, cards); len(got) != 0 {
		t.Errorf("clean deck: %v", got)
	}
}

func TestAllowedByRevisionAndPoolFilter(t *testing.T) {
	serenity := &mtgv1.Card{OracleId: "o-serenity", Name: "Angel of Serenity", ManaValue: 7, CardTypes: []string{"Creature"}}
	plains := &mtgv1.Card{OracleId: "o-plains", Name: "Plains", CardTypes: []string{"Land"}}
	lyra := &mtgv1.Card{OracleId: "o-lyra", Name: "Lyra Dawnbringer", ManaValue: 5, CardTypes: []string{"Creature"}}
	rev := &Revision{Remove: []string{"Lyra Dawnbringer"}, MaxManaValue: 5}
	if AllowedByRevision(rev, serenity) {
		t.Error("a 7-drop passed a cap of 5")
	}
	if !AllowedByRevision(rev, plains) {
		t.Error("a land failed the cap")
	}
	if AllowedByRevision(rev, lyra) {
		t.Error("a removed card passed")
	}
	if !AllowedByRevision(nil, serenity) {
		t.Error("no revision refused a card")
	}
	pool := NewPool([]*mtgv1.Card{serenity, plains, lyra}, map[string]int32{"o-plains": 30}).Filter(func(c *mtgv1.Card) bool { return AllowedByRevision(rev, c) })
	if pool.Size() != 1 {
		t.Errorf("pool = %v", pool.Names())
	}
	if pool.OwnedCount("o-plains") != 30 {
		t.Error("the filter lost the owned count")
	}
}

func TestRepairableReadsTheManaCap(t *testing.T) {
	v := &mtgv1.ValidationResult{Findings: []*mtgv1.Finding{{Code: CodeRevisionOverManaValue, Severity: mtgv1.Severity_SEVERITY_WARN}}}
	if len(repairable(v)) != 1 {
		t.Error("the mana cap warning did not buy the repair turn")
	}
}
