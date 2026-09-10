package quality

import (
	"fmt"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/meta"
)

// The casual corpus of PR-29 (F-53). The corpus read the great and the
// good lists alone, so a precon and its synergy-broken copy both sat far
// from a tournament list and the two barely separated. The casual group
// reads the typical and the baseline lists, and there the precon's pairs
// live and the broken copy's do not.

// TestTheTwoGroupsReadDifferentLists holds the split. A list that feeds
// one group must not feed the other, or the casual features would be a
// second copy of the top ones.
func TestTheTwoGroupsReadDifferentLists(t *testing.T) {
	for tier, want := range map[string]struct{ top, casual float64 }{
		meta.TierGreat:    {2, 0},
		meta.TierGood:     {1, 0},
		meta.TierTypical:  {0, 1},
		meta.TierBaseline: {0, 1},
		meta.TierBad:      {0, 0},
	} {
		t.Run(tier, func(t *testing.T) {
			if got := tierWeight(tier); got != want.top {
				t.Errorf("tierWeight(%s) = %g, want %g", tier, got, want.top)
			}
			if got := casualTierWeight(tier); got != want.casual {
				t.Errorf("casualTierWeight(%s) = %g, want %g", tier, got, want.casual)
			}
			// No list may carry weight in both groups.
			if tierWeight(tier) > 0 && casualTierWeight(tier) > 0 {
				t.Errorf("%s feeds both groups", tier)
			}
		})
	}
	// A bad list feeds neither. It is the broken copy the bar scores
	// against, and a corpus that learned from it would score it well.
	if tierWeight(meta.TierBad)+casualTierWeight(meta.TierBad) != 0 {
		t.Error("a bad list feeds a corpus group")
	}
}

// TestTheCasualCorpusSeparatesAPreconFromItsBrokenCopy is the mechanism
// of PR-29, and the reason the item exists. The pairs of a precon sit in
// the casual lists. A copy of that precon with its pairs broken holds
// the same cards and not the same pairs, so casual_synergy separates the
// two where synergy does not.
func TestTheCasualCorpusSeparatesAPreconFromItsBrokenCopy(t *testing.T) {
	w := newWorld(t)
	// A pair lifts only when it beats chance, so a corpus whose lists
	// all hold one set lifts nothing: every card is in every list, and
	// the lift reads exactly 1. Each group therefore holds its core in a
	// third of its lists and other cards in the rest.
	var train []*Resolved
	add := func(tier, id string, spells []*mtgv1.Card) {
		l := w.list(meta.SourceMTGJSON, id, tier, 22, spells)
		r := Resolve(&l, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD)
		if r == nil {
			t.Fatalf("the fixture list %s does not resolve", id)
		}
		train = append(train, r)
	}
	// The casual group pairs the first five staples, in ten of its
	// thirty lists. No great or good list holds that pairing.
	casualCore := w.staples[:5]
	for i := 0; i < 10; i++ {
		add(meta.TierBaseline, fmt.Sprintf("precon/core/%d", i), casualCore)
	}
	for i := 0; i < 20; i++ {
		add(meta.TierBaseline, fmt.Sprintf("precon/other/%d", i), w.staples[10+i%15:15+i%15])
	}
	// The top group pairs five other staples, the same way.
	topCore := w.staples[5:10]
	for i := 0; i < 10; i++ {
		add(meta.TierGreat, fmt.Sprintf("great/core/%d", i), topCore)
	}
	for i := 0; i < 20; i++ {
		add(meta.TierGreat, fmt.Sprintf("great/other/%d", i), w.staples[12+i%14:17+i%14])
	}

	fm := &FormatModel{}
	buildCorpus(fm, train, nil, w.idx)

	if len(fm.CardRates) == 0 || len(fm.Pairs) == 0 {
		t.Fatal("the top group built no tables")
	}
	if len(fm.CasualRates) == 0 || len(fm.CasualPairs) == 0 {
		t.Fatal("the casual group built no tables")
	}

	// The precon holds the casual pairs. Its broken copy holds the same
	// cards with the pairing gone: five cards nobody plays together.
	precon := deckOf(t, w, casualCore)
	broken := deckOf(t, w, append(append([]*mtgv1.Card{}, casualCore[:1]...), w.fillers[:4]...))

	pf := Features(Input{Deck: precon, Cards: w.idx}, fm)
	bf := Features(Input{Deck: broken, Cards: w.idx}, fm)

	if pf[KeyCasualSynergy] <= bf[KeyCasualSynergy] {
		t.Errorf("casual_synergy reads %g for the precon and %g for its broken copy, and the precon must read higher",
			pf[KeyCasualSynergy], bf[KeyCasualSynergy])
	}
	// And the casual rate reads the precon's own cards.
	if pf[KeyCasualRate] <= bf[KeyCasualRate] {
		t.Errorf("casual_rate reads %g for the precon and %g for its broken copy", pf[KeyCasualRate], bf[KeyCasualRate])
	}
	// The point of the item: the top group barely tells them apart,
	// because neither deck is a tournament list. If the top group ever
	// separated them this well on its own, PR-29 would not be needed.
	topGap := pf[KeySynergy] - bf[KeySynergy]
	casualGap := pf[KeyCasualSynergy] - bf[KeyCasualSynergy]
	if casualGap <= topGap {
		t.Errorf("the casual gap is %g and the top gap is %g, so the casual group adds nothing", casualGap, topGap)
	}
}

// TestAModelWithNoCasualTablesStillScores holds the compatibility rule.
// A model fitted before PR-29 carries no casual table, and every casual
// feature then reads zero rather than breaking the score.
func TestAModelWithNoCasualTablesStillScores(t *testing.T) {
	w := newWorld(t)
	fm := &FormatModel{
		CardRates: map[string]float64{}, Pairs: map[string]float64{}, RatePrior: 0.01,
	}
	deck := deckOf(t, w, w.staples[:5])
	f := Features(Input{Deck: deck, Cards: w.idx}, fm)
	if f[KeyCasualRate] != 0 {
		t.Errorf("casual_rate = %g on a model with no casual table, want 0", f[KeyCasualRate])
	}
	if f[KeyCasualSynergy] != 0 {
		t.Errorf("casual_synergy = %g on a model with no casual table, want 0", f[KeyCasualSynergy])
	}
	// The top features still read, so an old model scores as it did.
	if f[KeyCardRate] == 0 {
		t.Error("the top features stopped reading")
	}
	if got := fm.CasualRate("anything"); got != 0 {
		t.Errorf("CasualRate = %g on a model with no casual table, want the prior of 0", got)
	}
}

// TestEveryFeatureKeyHasAReason keeps the explain mode whole. A feature
// with no sentence reaches a reader as a bare key.
func TestEveryFeatureKeyHasAReason(t *testing.T) {
	for _, k := range Keys {
		p, ok := phrases[k]
		if !ok {
			t.Errorf("feature %s has no reason sentence", k)
			continue
		}
		if p.high == "" || p.low == "" {
			t.Errorf("feature %s has half a reason: %+v", k, p)
		}
	}
}

// TestTheCasualKeysJoinTheFit keeps the two new features in the fit. The
// fit reads Keys, so a key that is not in that list is a feature nothing
// ever weighs.
func TestTheCasualKeysJoinTheFit(t *testing.T) {
	want := map[string]bool{KeyCasualRate: false, KeyCasualSynergy: false}
	for _, k := range Keys {
		if _, ok := want[k]; ok {
			want[k] = true
		}
	}
	for k, found := range want {
		if !found {
			t.Errorf("%s is not in Keys, so the fit never weighs it", k)
		}
	}
}

// deckOf builds a Standard deck of the named spells plus lands, so
// Features has copies to read. The fixture world resolves its lists as
// Standard, and the two must agree: the playset and singleton features
// fire for a 60-card format alone, and a Commander deck reads zero on
// both.
func deckOf(t *testing.T, w *world, spells []*mtgv1.Card) *mtgv1.Deck {
	t.Helper()
	d := &mtgv1.Deck{Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_STANDARD}}
	for _, s := range spells {
		d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: s.GetOracleId(), Name: s.GetName(), Count: 4})
	}
	for _, l := range []*mtgv1.Card{w.plains, w.island} {
		d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: l.GetOracleId(), Name: l.GetName(), Count: 12})
	}
	return d
}
