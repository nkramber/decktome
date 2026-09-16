package profile

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// The PR-53 tests of the finisher count (F-138, D-726). M-17 counts the
// nine parent finisher tags with no child tag, the child tag
// blood-artist-ability, and the evasive creatures of power 5 or more.

// finisherTags builds a tag index of two trees and the evasion tag.
// drain-life holds the child blood-artist-ability, and mill-opponent
// holds the child mill-any, which carries the cards that win no game.
func finisherTags(t *testing.T) *cards.TagIndex {
	t.Helper()
	lines := strings.Join([]string{
		`{"id":"t1","slug":"drain-life","child_ids":["t2"],"taggings":[{"oracle_id":"oid-drainer"}]}`,
		`{"id":"t2","slug":"blood-artist-ability","taggings":[{"oracle_id":"oid-blood-artist"}]}`,
		`{"id":"t3","slug":"mill-opponent","child_ids":["t4"],"taggings":[{"oracle_id":"oid-miller"}]}`,
		`{"id":"t4","slug":"mill-any","taggings":[{"oracle_id":"oid-pilferer"}]}`,
		`{"id":"t5","slug":"evasion","taggings":[{"oracle_id":"oid-big-flier"},{"oracle_id":"oid-small-flier"},{"oracle_id":"oid-flying-rock"}]}`,
	}, "\n")
	idx, err := cards.LoadTags(strings.NewReader(lines), "tags")
	if err != nil {
		t.Fatal(err)
	}
	return idx
}

// TestFinisherSetReadsTheParentTagsAlone is D-726. The tree of
// mill-opponent holds cards that win no game, such as Ragavan, Nimble
// Pilferer, and 47 percent of the top-cut lists play it.
func TestFinisherSetReadsTheParentTagsAlone(t *testing.T) {
	set := FinisherSet(finisherTags(t))
	for _, id := range []string{"oid-drainer", "oid-blood-artist", "oid-miller"} {
		if !set[id] {
			t.Errorf("the finisher set holds no %s", id)
		}
	}
	if set["oid-pilferer"] {
		t.Error("the finisher set holds a card of a child tag")
	}
}

// TestAnEvasiveCreatureOfPowerFiveIsAFinisher is D-726. A creature of
// power 5 or more under the evasion tree closes a game on its own, and no
// finisher tag holds it.
func TestAnEvasiveCreatureOfPowerFiveIsAFinisher(t *testing.T) {
	tags := finisherTags(t)
	big := card(spec{name: "Big Flier", types: []string{"Creature"}})
	big.OracleId, big.Power = "oid-big-flier", "5"
	small := card(spec{name: "Small Flier", types: []string{"Creature"}})
	small.OracleId, small.Power = "oid-small-flier", "4"
	rock := card(spec{name: "Flying Rock", types: []string{"Artifact"}})
	rock.OracleId, rock.Power = "oid-flying-rock", "5"
	drainer := card(spec{name: "Drainer", types: []string{"Creature"}})
	drainer.OracleId, drainer.Power = "oid-drainer", "1"
	of := PowerOf(tags)
	for _, tc := range []struct {
		card *mtgv1.Card
		want bool
	}{{big, true}, {small, false}, {rock, false}, {drainer, true}} {
		got := false
		for _, key := range of(tc.card) {
			got = got || key == KeyFinisher
		}
		if got != tc.want {
			t.Errorf("%s reads finisher %v, want %v", tc.card.GetName(), got, tc.want)
		}
	}
	// The id set of a snapshot holds the evasive creature too.
	ids := FinisherIDs([]*mtgv1.Card{big, small, rock, drainer}, tags)
	if !ids["oid-big-flier"] || ids["oid-small-flier"] || ids["oid-flying-rock"] {
		t.Errorf("the id set of the snapshot reads %v", ids)
	}
}

// TestTheFinisherFeatureCountsTheDeck is D-726. The stored profile holds
// the count, so the check and the gap note read the same number.
func TestTheFinisherFeatureCountsTheDeck(t *testing.T) {
	tags := finisherTags(t)
	drainer := card(spec{name: "Drainer", types: []string{"Creature"}})
	drainer.OracleId = "oid-drainer"
	big := card(spec{name: "Big Flier", types: []string{"Creature"}})
	big.OracleId, big.Power = "oid-big-flier", "5"
	pilferer := card(spec{name: "Pilferer", types: []string{"Creature"}})
	pilferer.OracleId, pilferer.Power = "oid-pilferer", "2"
	f := &features{}
	f.finishers([]entry{{card: drainer, count: 1}, {card: big, count: 2}, {card: pilferer, count: 1}}, tags)
	if got := f.values[KeyFinisher].value; got != 3 {
		t.Errorf("the finisher count reads %v, want 3", got)
	}
	// Without tags the row is absent, as the tutor row is.
	empty := &features{}
	empty.finishers([]entry{{card: drainer, count: 1}}, nil)
	if _, ok := empty.values[KeyFinisher]; ok {
		t.Error("a snapshot with no tag holds a finisher row")
	}
}

// TestTheFinisherFloorOfEachBracket is D-726. Brackets 1 to 4 take a
// floor of 2, and bracket 5 takes 1.
func TestTheFinisherFloorOfEachBracket(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	for bracket, want := range map[int32]float64{1: 2, 2: 2, 3: 2, 4: 2, 5: 1} {
		power := &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bracket}}
		floors := b.PowerFloors(mtgv1.FormatId_FORMAT_ID_COMMANDER, power)
		if got := floors[KeyFinisher]; got != want {
			t.Errorf("bracket %d reads a finisher floor of %v, want %v", bracket, got, want)
		}
		lines := strings.Join(b.Lines(mtgv1.FormatId_FORMAT_ID_COMMANDER, power), "\n")
		if !strings.Contains(lines, "finishers") || !strings.Contains(lines, `marks each one "finisher"`) {
			t.Errorf("bracket %d writes no finisher line for the prompt", bracket)
		}
	}
}

// TestTheFinisherFindingNamesItselfAndItsCode is F-150, F-152, and
// D-743. The finding read " is 0, and bracket 3 wants 2 or more": no
// feature word named the finisher, so the sentence held no subject. It
// also carried the off-band code, so it reached the repair input, and a
// pool that holds too few finishers can not close the gap.
func TestTheFinisherFindingNamesItselfAndItsCode(t *testing.T) {
	if got := Word(KeyFinisher); got != "the finisher count" {
		t.Errorf("the finisher feature word reads %q", got)
	}
	row := &mtgv1.ProfileFeature{Key: KeyFinisher, Value: 0, Low: 2}
	msg := offBandMessage(row, 3, mtgv1.FormatId_FORMAT_ID_COMMANDER, nil)
	if !strings.HasPrefix(msg, "the finisher count is 0") {
		t.Errorf("the finding reads %q", msg)
	}
	// The bracket counts no finisher, so the finding names the deck plan.
	if strings.Contains(msg, "bracket") {
		t.Errorf("the finisher finding names the bracket: %q", msg)
	}
	if !strings.Contains(msg, "this deck plan wants 2 or more") {
		t.Errorf("the finding reads %q", msg)
	}
	// Every other feature keeps the bracket in its finding.
	land := offBandMessage(&mtgv1.ProfileFeature{Key: KeyLand, Value: 30, Low: 34, High: 38, HasHigh: true},
		3, mtgv1.FormatId_FORMAT_ID_COMMANDER, nil)
	if !strings.Contains(land, "bracket 3 wants 34 to 38") {
		t.Errorf("the land finding reads %q", land)
	}
	if CodeFinisherShort == CodeOffBand {
		t.Error("the finisher finding must carry its own code, so no repair input holds it")
	}
}
