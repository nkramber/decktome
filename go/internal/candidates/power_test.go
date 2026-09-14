package candidates

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// powerCards are the cards of testCards, a tutor on the Game Changers
// list, and a Game Changer that no top list plays.
func powerCards() []tc {
	return append(testCards(),
		tc{id: "tutor", name: "Demonic Tutor", typeLine: "Sorcery", text: "Search your library for a card, put that card into your hand, then shuffle.",
			identity: []mtgv1.Color{B}, mv: 2, rank: 50, tags: []string{"tutor"}, gameChanger: true},
		tc{id: "gc0", name: "Unplayed Changer", typeLine: "Enchantment", text: "Draw a card.",
			identity: []mtgv1.Color{W}, mv: 3, rank: 9000, gameChanger: true},
	)
}

// fixedRates is a MetaBoost over fixed rates, and 0 for every other card.
func fixedRates(m map[string]float64) func(string) float64 {
	return func(id string) float64 { return m[id] }
}

// TestPowerRateKeepsAHighRateCard is D-707. At brackets 4 and 5 an
// off-theme tutor whose rate reaches the keep rate stays on the list,
// with its full score and the rate signal. Under the keep rate, or below
// bracket 4, the theme cut drops it as before (F-130).
func TestPowerRateKeepsAHighRateCard(t *testing.T) {
	b, _ := New()
	idx := fixture(t, powerCards())
	rate, weight := 0.6, 0.3
	for _, tt := range []struct {
		name    string
		bracket int32
		keep    float64
		want    bool
	}{
		{"bracket 4 over the keep rate", 4, 0.5, true},
		{"bracket 5 at the keep rate", 5, 0.6, true},
		{"bracket 4 under the keep rate", 4, 0.7, false},
		{"bracket 3", 3, 0.5, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain", Bracket: tt.bracket,
				MetaBoost: fixedRates(map[string]float64{"tutor": rate}), PowerRate: PowerRate{Weight: weight, Keep: tt.keep}})
			if err != nil {
				t.Fatal(err)
			}
			c, ok := find(list.Candidates, "Demonic Tutor")
			if ok != tt.want {
				t.Fatalf("Demonic Tutor listed = %v, want %v", ok, tt.want)
			}
			if !ok {
				return
			}
			want := popularity(c.Card, maxRankOf(idx))*0.3 + rate*weight
			if c.Score != want || c.Rate != rate || !slices.Contains(c.Signals, "top-list rate") {
				t.Errorf("score %v, rate %v, signals %v; want score %v, rate %v, and the rate signal", c.Score, c.Rate, c.Signals, want, rate)
			}
		})
	}
}

// TestPowerWeightReadsOnlyAtBracketsFourAndFive is D-704: the rate weighs
// more at brackets 4 and 5, and a bracket 3 request scores as before.
func TestPowerWeightReadsOnlyAtBracketsFourAndFive(t *testing.T) {
	b, _ := New()
	idx := fixture(t, powerCards())
	rate := 0.5
	score := func(bracket int32) float64 {
		t.Helper()
		list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain", Bracket: bracket,
			MetaBoost: fixedRates(map[string]float64{"wrath": rate}), PowerRate: PowerRate{Weight: 0.4, Keep: 0.9}})
		if err != nil {
			t.Fatal(err)
		}
		c, ok := find(list.Candidates, "Wrath of God")
		if !ok {
			t.Fatalf("bracket %d lists no Wrath of God", bracket)
		}
		return c.Score
	}
	wrath, _ := idx.ByOracleID("wrath")
	pop := popularity(wrath, maxRankOf(idx))
	if got, want := score(3), (pop*0.3+rate*0.1)*staplePenalty; got != want {
		t.Errorf("bracket 3 score %v, want %v", got, want)
	}
	if got, want := score(4), (pop*0.3+rate*0.4)*staplePenalty; got != want {
		t.Errorf("bracket 4 score %v, want %v", got, want)
	}
}

// TestReserveListsThePowerCardsByRate is D-709. A bracket 5 list in the
// owned-only mode lists owned cards alone, and its reserve still reads each
// power card in the colors that the top lists play, by rate. A card with
// no rate is no answer, and a bracket 3 list holds no reserve.
func TestReserveListsThePowerCardsByRate(t *testing.T) {
	b, _ := New()
	idx := fixture(t, powerCards())
	req := Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain", Bracket: 5,
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, Owned: map[string]int32{"solring": 1},
		MetaBoost: fixedRates(map[string]float64{"solring": 0.9, "tutor": 0.6, "gc": 0.4})}
	list, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := find(list.Candidates, "Demonic Tutor"); ok {
		t.Error("the owned-only list holds a card the collection lacks")
	}
	var got []string
	for _, c := range list.Reserve {
		got = append(got, fmt.Sprintf("%s %g %d", c.Card.GetName(), c.Rate, c.Owned))
	}
	if want := "Sol Ring 0.9 1,Demonic Tutor 0.6 0,Smothering Tithe 0.4 0"; strings.Join(got, ",") != want {
		t.Errorf("reserve %v, want %q", got, want)
	}
	req.Bracket = 3
	if list, err = b.Build(idx, req); err != nil {
		t.Fatal(err)
	}
	if list.Reserve != nil {
		t.Errorf("a bracket 3 list holds a reserve: %v", names(list.Reserve))
	}
}
