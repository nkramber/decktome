package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// TestBuildDropsTheExcludedCards is D-408: a card of an excluded precon
// with no copy to spare never reaches the shortlist, whatever its rank.
func TestBuildDropsTheExcludedCards(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := setFixture(t)
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain",
		Colors: []mtgv1.Color{W}, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
	}
	before, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if got := idSet(before.Candidates); !got["outrock"] || !got["insoul"] {
		t.Fatalf("the fixture must offer Sol Ring and Soul Warden before the exclusion: %v", got)
	}
	req.ExcludeOracleIDs = []string{"outrock", "insoul"}
	after, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	got := idSet(after.Candidates)
	for _, id := range req.ExcludeOracleIDs {
		if got[id] {
			t.Errorf("the shortlist holds %q, which the reader excluded", id)
		}
	}
	if !got["outrock2"] || !got["inangel"] {
		t.Errorf("the exclusion dropped a card it did not name: %v", got)
	}
}

// TestCommanderPoolDropsAnExcludedCommander: the commander of an excluded
// precon can not lead the deck (D-408).
func TestCommanderPoolDropsAnExcludedCommander(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain"}
	pool, err := b.CommanderPool(setFixture(t), req)
	if err != nil {
		t.Fatal(err)
	}
	if !idSet(pool)["incmdr"] {
		t.Fatal("the fixture must offer Thranduil before the exclusion")
	}
	req.ExcludeOracleIDs = []string{"incmdr"}
	pool, err = b.CommanderPool(setFixture(t), req)
	if err != nil {
		t.Fatal(err)
	}
	got := idSet(pool)
	if got["incmdr"] {
		t.Error("the pool offers a commander the reader excluded")
	}
	if !got["outcmdr"] {
		t.Error("the exclusion dropped a commander it did not name")
	}
}
