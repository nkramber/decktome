package generate

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestCheaperOfRoleKeepsTheJob is F-133. The step that swaps the costliest
// spell for a cheaper one offers only cards of the same job, so the added
// card does the job its label names. The step first offered every cheaper
// card, and a mana rock came in with the job removal.
func TestCheaperOfRoleKeepsTheJob(t *testing.T) {
	rock := &mtgv1.Card{OracleId: "o-rock", Name: "Cheap Rock", CardTypes: []string{"Artifact"}, ManaValue: 0}
	bolt := &mtgv1.Card{OracleId: "o-bolt", Name: "Cheap Bolt", CardTypes: []string{"Instant"}, ManaValue: 1}
	study := &mtgv1.Card{OracleId: "o-study", Name: "Cheap Study", CardTypes: []string{"Enchantment"}, ManaValue: 2}
	// A card with no shortlist job, such as a commander, is no answer.
	legend := &mtgv1.Card{OracleId: "o-legend", Name: "Cheap Legend", CardTypes: []string{"Creature"}, ManaValue: 1}
	req := Request{
		Pool:  NewPool([]*mtgv1.Card{rock, bolt, study, legend}, nil),
		Roles: map[string]string{"o-rock": "ramp", "o-bolt": "removal", "o-study": "draw"},
	}
	got := (&Builder{}).cheaperOfRole(req, map[string]bool{}, mtgv1.CardRole_CARD_ROLE_REMOVAL, 4)
	if len(got) != 1 || got[0].GetName() != "Cheap Bolt" {
		var names []string
		for _, c := range got {
			names = append(names, c.GetName())
		}
		t.Errorf("a removal spell of mana value 4 is offered %v, want Cheap Bolt alone", names)
	}
}
