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

// TestCheckRevisionReadsTheSideboardAndTheCommander: a removed card in
// the sideboard stayed, and a kept card that is the commander is held.
func TestCheckRevisionReadsTheSideboardAndTheCommander(t *testing.T) {
	cards := cardMap{
		"o-karlov": {OracleId: "o-karlov", Name: "Karlov of the Ghost Council", ManaValue: 2, CardTypes: []string{"Creature"}},
		"o-lyra":   {OracleId: "o-lyra", Name: "Lyra Dawnbringer", ManaValue: 5, CardTypes: []string{"Creature"}},
	}
	deck := &mtgv1.Deck{
		CommanderOracleIds: []string{"o-karlov"},
		Sideboard:          []*mtgv1.DeckCard{{OracleId: "o-lyra", Name: "Lyra Dawnbringer", Count: 1}},
	}
	rev := &Revision{Remove: []string{"Lyra Dawnbringer"}, Keep: []string{"Karlov of the Ghost Council"}}
	codes := map[string]bool{}
	for _, f := range CheckRevision(deck, rev, cards) {
		codes[f.GetCode()] = true
	}
	if !codes[CodeRevisionRemovedPresent] {
		t.Error("a removed card in the sideboard was not reported")
	}
	if codes[CodeRevisionKeptMissing] {
		t.Error("the kept commander was reported missing")
	}
}

// TestRevisionCapExemptsKeptLockedAndCommanderCards: a kept name, a
// locked card, and a commander above the cap stay in the pool, or the
// deck must hold a card the model can not name (D-242).
func TestRevisionCapExemptsKeptLockedAndCommanderCards(t *testing.T) {
	serenity := &mtgv1.Card{OracleId: "o-serenity", Name: "Angel of Serenity", ManaValue: 7, CardTypes: []string{"Creature"}}
	titan := &mtgv1.Card{OracleId: "o-titan", Name: "Sun Titan", ManaValue: 6, CardTypes: []string{"Creature"}}
	other := &mtgv1.Card{OracleId: "o-other", Name: "Other Seven", ManaValue: 7, CardTypes: []string{"Creature"}}
	rev := &Revision{Keep: []string{"angel of serenity"}, MaxManaValue: 5, Exempt: []string{"o-titan"}}
	if !AllowedByRevision(rev, serenity) {
		t.Error("a kept 7-drop left the pool under a cap of 5")
	}
	if !AllowedByRevision(rev, titan) {
		t.Error("an exempt 6-drop left the pool under a cap of 5")
	}
	if AllowedByRevision(rev, other) {
		t.Error("a 7-drop that is neither kept nor exempt passed the cap")
	}
	pool := NewPool([]*mtgv1.Card{serenity, titan, other}, nil).Filter(func(c *mtgv1.Card) bool { return AllowedByRevision(rev, c) })
	if pool.Size() != 2 {
		t.Errorf("pool = %v, want the kept card and the exempt card", pool.Names())
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

// TestMissingLockedNamesACardOutsideThePool is D-301. A revision drops a
// locked card from the pool, and the finding must still name it.
func TestMissingLockedNamesACardOutsideThePool(t *testing.T) {
	arena := &mtgv1.Card{OracleId: "o-arena", Name: "Arena of Glory"}
	req := Request{Locked: []string{"o-arena"}, Pool: NewPool(nil, nil)}
	got := missingLocked(&mtgv1.Deck{}, req, cardMap{"o-arena": arena})
	if len(got) != 1 || got[0] != "Arena of Glory" {
		t.Errorf("missing = %v, want the name", got)
	}
	if got := missingLocked(&mtgv1.Deck{}, req, nil); len(got) != 1 || got[0] != "o-arena" {
		t.Errorf("with no card source: %v", got)
	}
}
