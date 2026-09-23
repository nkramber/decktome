package candidates

import (
	"fmt"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestTotalCutKeepsTheManaHalf is F-165. The staple penalty puts an
// untapped dual at the score of the cut line, and the total cut dropped
// 11 of the 20 mana lands of a five-color shortlist. The mana half of the
// land cap goes before the score cut, inside the total (D-450, D-835).
func TestTotalCutKeepsTheManaHalf(t *testing.T) {
	land := func(name string, score, pop float64, rank int, themed bool) Candidate {
		return Candidate{Card: &mtgv1.Card{OracleId: name, Name: name}, Role: mtgv1.CardRole_CARD_ROLE_LAND,
			Score: score, Pop: pop, LandRank: rank, Themed: themed}
	}
	threat := func(i int) Candidate {
		name := fmt.Sprintf("Threat %d", i)
		return Candidate{Card: &mtgv1.Card{OracleId: name, Name: name}, Role: mtgv1.CardRole_CARD_ROLE_THREAT,
			Score: 0.5 - float64(i)/100}
	}
	in := []Candidate{
		land("Theme Land A", 0.6, 0.2, 9, true), land("Theme Land B", 0.59, 0.2, 9, true),
		land("Watery Grave", 0.15, 0.99, 1, false), land("City of Brass", 0.149, 0.98, 1, false),
		land("Minor Dual", 0.148, 0.5, 1, false),
	}
	for i := range 5 {
		in = append(in, threat(i))
	}
	sortCandidates(in)
	lim := Limits{Total: 6, PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_LAND: 4, mtgv1.CardRole_CARD_ROLE_THREAT: 10}}
	got := capByRole(in, lim)
	if len(got) != 6 {
		t.Fatalf("the cut kept %d cards, want the total of 6", len(got))
	}
	kept := map[string]Candidate{}
	for _, c := range got {
		kept[c.Card.GetName()] = c
	}
	for _, name := range []string{"Watery Grave", "City of Brass"} {
		c, ok := kept[name]
		if !ok {
			t.Errorf("the total cut dropped %s, a land of the mana half", name)
			continue
		}
		if !c.ManaHalf {
			t.Errorf("%s lost its mark of the mana half", name)
		}
	}
	for _, name := range []string{"Theme Land A", "Theme Land B", "Threat 0", "Threat 1"} {
		if _, ok := kept[name]; !ok {
			t.Errorf("the cut dropped %s over a lower score", name)
		}
	}
	if c, ok := kept["Theme Land A"]; ok && c.ManaHalf {
		t.Error("a theme land reads the mark of the mana half")
	}

	// A total under the mana half keeps its best lands in the mana order.
	lim.Total = 1
	got = capByRole(in, lim)
	if len(got) != 1 || got[0].Card.GetName() != "Watery Grave" {
		t.Errorf("a total of 1 kept %v, want Watery Grave", names(got))
	}
}
