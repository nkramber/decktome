package quality

import (
	"context"
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// The test world: a 60-card format with thirty staples the top lists
// play, thirty fillers they do not, and a mana base. Every card is
// legal in Standard, so the synthetic pool is the whole world.

var (
	W = mtgv1.Color_COLOR_W
	U = mtgv1.Color_COLOR_U
)

func spell(name string, mv float64, color mtgv1.Color) *mtgv1.Card {
	return &mtgv1.Card{
		OracleId: "oid-" + strings.ToLower(strings.ReplaceAll(name, " ", "-")), Name: name,
		ManaValue: mv, ManaCost: fmt.Sprintf("{%d}{%s}", int(mv)-1, strings.TrimPrefix(color.String(), "COLOR_")),
		CardTypes: []string{"Creature"}, Colors: []mtgv1.Color{color}, ColorIdentity: []mtgv1.Color{color},
		Legalities: map[string]mtgv1.LegalityStatus{"standard": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
}

func land(name string, basic bool, produced ...mtgv1.Color) *mtgv1.Card {
	c := &mtgv1.Card{
		OracleId: "oid-" + strings.ToLower(strings.ReplaceAll(name, " ", "-")), Name: name,
		CardTypes: []string{"Land"}, ProducedMana: produced, ColorIdentity: produced,
		Legalities: map[string]mtgv1.LegalityStatus{"standard": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	if basic {
		c.Supertypes = []string{"Basic"}
	}
	return c
}

type world struct {
	idx      *cards.Index
	staples  []*mtgv1.Card
	fillers  []*mtgv1.Card
	plains   *mtgv1.Card
	island   *mtgv1.Card
	duals    []*mtgv1.Card
	profiler *profile.Profiler
}

func newWorld(t *testing.T) *world {
	t.Helper()
	w := &world{}
	var all []*mtgv1.Card
	for i := 0; i < 30; i++ {
		color := W
		if i%2 == 1 {
			color = U
		}
		c := spell(fmt.Sprintf("Staple %02d", i), float64(1+i%4), color)
		w.staples = append(w.staples, c)
		all = append(all, c)
	}
	for i := 0; i < 30; i++ {
		color := W
		if i%2 == 1 {
			color = U
		}
		c := spell(fmt.Sprintf("Filler %02d", i), float64(3+i%5), color)
		w.fillers = append(w.fillers, c)
		all = append(all, c)
	}
	w.plains = land("Plains", true, W)
	w.island = land("Island", true, U)
	all = append(all, w.plains, w.island)
	for i := 0; i < 4; i++ {
		c := land(fmt.Sprintf("Dual %d", i), false, W, U)
		w.duals = append(w.duals, c)
		all = append(all, c)
	}
	w.idx = cards.NewIndex(all, nil, nil, time.Date(2026, 9, 2, 0, 0, 0, 0, time.UTC))
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	w.profiler, err = profile.New(cfg, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	return w
}

// list makes a 60-card list: lands, then playsets of the named spells.
func (w *world) list(source, id, tier string, lands int, spells []*mtgv1.Card) meta.List {
	l := meta.List{Source: source, ID: id, Format: meta.FormatStandard, Date: "2026-09-01", Tier: tier}
	l.Cards = append(l.Cards, meta.Card{Name: "Plains", Count: lands / 2}, meta.Card{Name: "Island", Count: lands - lands/2})
	for _, s := range spells {
		l.Cards = append(l.Cards, meta.Card{Name: s.GetName(), Count: 4})
	}
	return l
}

// corpus is the ladder: great lists of nine staple playsets, good lists
// with one filler playset among them, and baseline products with four
// filler playsets and two lands fewer.
func (w *world) corpus() []meta.List {
	var out []meta.List
	pick := func(from []*mtgv1.Card, start, n int) []*mtgv1.Card {
		var s []*mtgv1.Card
		for i := 0; i < n; i++ {
			s = append(s, from[(start+i)%len(from)])
		}
		return s
	}
	for i := 0; i < 40; i++ {
		out = append(out, w.list(meta.SourceMTGO, fmt.Sprintf("great/%d", i), meta.TierGreat, 24, pick(w.staples, i, 9)))
	}
	for i := 0; i < 30; i++ {
		spells := append(pick(w.staples, i*2, 8), w.fillers[i%len(w.fillers)])
		out = append(out, w.list(meta.SourceMTGO, fmt.Sprintf("good/%d", i), meta.TierGood, 24, spells))
	}
	for i := 0; i < 30; i++ {
		spells := append(pick(w.staples, i*3, 5), pick(w.fillers, i, 4)...)
		out = append(out, w.list(meta.SourceMTGJSON, fmt.Sprintf("precon/%d", i), meta.TierBaseline, 22, spells))
	}
	// A product older than the pool serves no baseline (D-478).
	old := w.list(meta.SourceMTGJSON, "precon/1999", meta.TierBaseline, 22, w.fillers[:9])
	old.Format, old.Date = meta.FormatSixty, "1999-01-01"
	out = append(out, old)
	return out
}

func fitWorld(t *testing.T) (*world, *Model, *FitReport) {
	t.Helper()
	w := newWorld(t)
	model, rep, err := Fit(context.Background(), FitInput{
		Index: w.idx, Profiler: w.profiler, Lists: w.corpus(),
		Now:     time.Date(2026, 9, 2, 15, 0, 0, 0, time.UTC),
		Formats: []mtgv1.FormatId{mtgv1.FormatId_FORMAT_ID_STANDARD},
	})
	if err != nil {
		t.Fatalf("fit: %v (%+v)", err, rep)
	}
	return w, model, rep
}

func TestFitSeparatesTheLadder(t *testing.T) {
	_, model, rep := fitWorld(t)
	fm := model.Formats[meta.FormatStandard]
	if fm == nil {
		t.Fatalf("no standard model: %+v", rep.Formats)
	}
	fr := rep.Formats[meta.FormatStandard]
	// Thirty precons break on five axes each (D-484).
	if fr.Used != 100 || fr.Synthetic != 150 || fr.Unusable != 0 || fr.OutOfPool != 1 {
		t.Errorf("report = %+v", fr)
	}
	want := []string{meta.TierBad, meta.TierBaseline, meta.TierGood, meta.TierGreat}
	if strings.Join(fm.Tiers, ",") != strings.Join(want, ",") {
		t.Errorf("tiers = %v, want %v", fm.Tiers, want)
	}
	if len(fm.Thresholds) != 3 {
		t.Errorf("thresholds = %v", fm.Thresholds)
	}
	h := fm.Holdout
	if h.Lists == 0 {
		t.Fatal("no holdout list")
	}
	if h.GreatOverBaseline.Share() < 0.9 {
		t.Errorf("great over baseline = %.2f of %d pairs, the bar is 0.9", h.GreatOverBaseline.Share(), h.GreatOverBaseline.Pairs)
	}
	if h.BaselineOverBad.Share() < 0.95 {
		t.Errorf("baseline over bad = %.2f of %d pairs, the bar is 0.95", h.BaselineOverBad.Share(), h.BaselineOverBad.Pairs)
	}
	if h.Accuracy < 0.6 {
		t.Errorf("accuracy = %.2f", h.Accuracy)
	}
	if fm.Rate("oid-staple-00") <= fm.RatePrior || fm.Rate("oid-filler-05") > 0.5 {
		t.Errorf("rates: staple %.4f, filler %.4f, prior %.4f", fm.Rate("oid-staple-00"), fm.Rate("oid-filler-05"), fm.RatePrior)
	}
	if len(fm.TopCards) != TopCardCount || !strings.HasPrefix(fm.TopCards[0], "Staple") {
		t.Errorf("top cards = %v", fm.TopCards)
	}
	if fm.Shape.Lists == 0 || fm.Shape.Lands < 23 || fm.Shape.Lands > 25 {
		t.Errorf("shape = %+v", fm.Shape)
	}
	if len(fm.Pairs) == 0 {
		t.Errorf("no pair lifted")
	}
	for _, k := range []string{KeyCEDHSignal, KeyHighBracket, KeyCommanderDecks} {
		for _, have := range fm.Keys {
			if have == k {
				t.Errorf("%s has spread in a 60-card format", k)
			}
		}
	}
	t.Logf("holdout %+v, weights %v", h, fm.Weights)
}

func TestScoreAndRoundTrip(t *testing.T) {
	w, model, _ := fitWorld(t)
	data, err := model.Encode()
	if err != nil {
		t.Fatal(err)
	}
	back, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	s := NewScorer(back)
	if s.Version() != "20260902T150000Z" {
		t.Errorf("version = %s", s.Version())
	}
	great := w.list(meta.SourceMTGO, "probe", meta.TierGreat, 24, w.staples[:9])
	r := Resolve(&great, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD)
	prof := w.profiler.Measure(r.Deck, w.idx)
	q := s.Score(Input{Deck: r.Deck, Profile: prof, Cards: w.idx})
	if q == nil {
		t.Fatal("no grade")
	}
	if q.GetTier() != meta.TierGreat && q.GetTier() != meta.TierGood {
		t.Errorf("a staple list grades %s: %+v", q.GetTier(), q)
	}
	if len(q.GetReasons()) != ReasonCount || q.GetModelVersion() == "" || len(q.GetProbabilities()) != 4 {
		t.Errorf("grade = %+v", q)
	}
	sum := 0.0
	for _, p := range q.GetProbabilities() {
		sum += p.GetProbability()
	}
	if math.Abs(sum-1) > 0.01 {
		t.Errorf("probabilities sum to %.3f", sum)
	}
	bad := w.list(meta.SourceMTGJSON, "probe-bad", meta.TierBad, 12, append(w.fillers[:10], w.fillers[10:12]...))
	rb := Resolve(&bad, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD)
	qb := s.Score(Input{Deck: rb.Deck, Profile: w.profiler.Measure(rb.Deck, w.idx), Cards: w.idx})
	if qb.GetScore() >= q.GetScore() {
		t.Errorf("a filler list with 12 lands scores %.3f, the staple list %.3f", qb.GetScore(), q.GetScore())
	}
	if !strings.HasPrefix(Summary(q), "The quality model grades this deck") {
		t.Errorf("summary = %q", Summary(q))
	}
	// A format with no model grades nothing, and the shortlist boost
	// reads the rates.
	r.Deck.Format.Id = mtgv1.FormatId_FORMAT_ID_MODERN
	if s.Score(Input{Deck: r.Deck, Profile: prof, Cards: w.idx}) != nil {
		t.Errorf("a format with no model graded a deck")
	}
	boost := s.MetaBoost(mtgv1.FormatId_FORMAT_ID_STANDARD)
	if boost == nil || boost("oid-staple-00") <= boost("oid-filler-05") || boost("oid-staple-00") > 1 {
		t.Errorf("boost: staple %v, filler %v", boost("oid-staple-00"), boost("oid-filler-05"))
	}
	if s.MetaBoost(mtgv1.FormatId_FORMAT_ID_MODERN) != nil || s.CommanderSignal() != nil {
		t.Errorf("a missing model must answer nil")
	}
	rows := s.CardQualities("oid-staple-00")
	if len(rows) != 1 || rows[0].GetFormat() != mtgv1.FormatId_FORMAT_ID_STANDARD || rows[0].GetInclusion() == 0 {
		t.Errorf("card qualities = %+v", rows)
	}
	lines := s.ShapeLines(mtgv1.FormatId_FORMAT_ID_STANDARD)
	if len(lines) != 2 || !strings.Contains(lines[0], "lands") || !strings.Contains(lines[1], "Staple") {
		t.Errorf("shape lines = %v", lines)
	}
}

func TestSynthesizeAxes(t *testing.T) {
	w := newWorld(t)
	l := w.list(meta.SourceMTGO, "base", meta.TierGreat, 24, w.staples[:9])
	r := Resolve(&l, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD)
	pl := newPool(w.idx, mtgv1.FormatId_FORMAT_ID_STANDARD)
	count := func(d *mtgv1.Deck, land bool) (n int) {
		for _, dc := range d.GetCards() {
			c, _ := w.idx.ByOracleID(dc.GetOracleId())
			if (c.GetCardTypes()[0] == "Land") == land {
				n += int(dc.GetCount())
			}
		}
		return n
	}
	avgMV := func(d *mtgv1.Deck) float64 {
		var sum, n float64
		for _, dc := range d.GetCards() {
			c, _ := w.idx.ByOracleID(dc.GetOracleId())
			if c.GetCardTypes()[0] != "Land" {
				sum += c.GetManaValue() * float64(dc.GetCount())
				n += float64(dc.GetCount())
			}
		}
		return sum / n
	}
	lands := Synthesize(r, DefectLands, pl, nil)
	if count(lands.Deck, true) != 16 || count(lands.Deck, false) != 44 {
		t.Errorf("lands defect: %d lands, %d spells", count(lands.Deck, true), count(lands.Deck, false))
	}
	if lands.List.Tier != meta.TierBad || lands.List.Defect != DefectLands || lands.List.Source != meta.SourceSynthetic || lands.List.Key() != "synthetic:mtgo:base/lands" {
		t.Errorf("lands list = %+v", lands.List)
	}
	curve := Synthesize(r, DefectCurve, pl, nil)
	if avgMV(curve.Deck) <= avgMV(r.Deck)+1 {
		t.Errorf("curve defect: mv %.2f from %.2f", avgMV(curve.Deck), avgMV(r.Deck))
	}
	copies := Synthesize(r, DefectCopies, pl, nil)
	playsets := 0
	for _, dc := range copies.Deck.GetCards() {
		if dc.GetCount() >= 4 && !strings.HasPrefix(dc.GetName(), "Plains") && !strings.HasPrefix(dc.GetName(), "Island") {
			playsets++
		}
	}
	if playsets != 0 || count(copies.Deck, false) != 36 {
		t.Errorf("copies defect: %d playsets, %d spells", playsets, count(copies.Deck, false))
	}
	synergy := Synthesize(r, DefectSynergy, pl, nil)
	if count(synergy.Deck, false) != 36 || count(synergy.Deck, true) != 24 {
		t.Errorf("synergy defect: %d spells, %d lands", count(synergy.Deck, false), count(synergy.Deck, true))
	}
	// The same seed makes the same deck.
	again := Synthesize(r, DefectSynergy, pl, nil)
	if len(again.Deck.GetCards()) != len(synergy.Deck.GetCards()) || again.Deck.GetCards()[len(again.Deck.GetCards())-1].GetName() != synergy.Deck.GetCards()[len(synergy.Deck.GetCards())-1].GetName() {
		t.Errorf("the synthesis is not deterministic")
	}
	// A list with basics alone can not break on colors, so it breaks on
	// synergy.
	colors := Synthesize(r, DefectColors, pl, nil)
	if colors.List.Defect != DefectSynergy {
		t.Errorf("colors on basics = %s", colors.List.Defect)
	}
	withDuals := w.list(meta.SourceMTGO, "duals", meta.TierGreat, 20, w.staples[:9])
	for _, d := range w.duals {
		withDuals.Cards = append(withDuals.Cards, meta.Card{Name: d.GetName(), Count: 1})
	}
	rd := Resolve(&withDuals, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD)
	colors = Synthesize(rd, DefectColors, pl, nil)
	if colors.List.Defect != DefectColors {
		t.Fatalf("colors on duals = %s", colors.List.Defect)
	}
	for _, dc := range colors.Deck.GetCards() {
		if strings.HasPrefix(dc.GetName(), "Dual") {
			t.Errorf("a dual survived the colors defect")
		}
	}
}

func TestFitOrdinalSeparable(t *testing.T) {
	var x [][]float64
	var y []int
	for i := 0; i < 90; i++ {
		v := float64(i%30)/10 - 1.5
		x = append(x, []float64{v})
		y = append(y, i%30/10)
	}
	w, th, loss := fitOrdinal(x, y, 3)
	if len(w) != 1 || w[0] <= 0 || len(th) != 2 || th[0] >= th[1] {
		t.Fatalf("w %v, th %v", w, th)
	}
	if loss > 0.5 {
		t.Errorf("loss = %.3f", loss)
	}
	right := 0
	for i, z := range x {
		p := levelProbabilities(w, th, z)
		best := 0
		for k := range p {
			if p[k] > p[best] {
				best = k
			}
		}
		if best == y[i] {
			right++
		}
	}
	if right < 85 {
		t.Errorf("%d of 90 right", right)
	}
}

func TestResolveUsable(t *testing.T) {
	w := newWorld(t)
	l := w.list(meta.SourceMTGO, "x", meta.TierGood, 24, w.staples[:9])
	l.Cards = append(l.Cards, meta.Card{Name: "No Such Card", Count: 1})
	r := Resolve(&l, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD)
	if len(r.Missing) != 1 || !r.Usable() {
		t.Errorf("one miss: missing %v, usable %v", r.Missing, r.Usable())
	}
	short := w.list(meta.SourceMTGO, "y", meta.TierGood, 10, w.staples[:3])
	if Resolve(&short, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD).Usable() {
		t.Errorf("a 22-card list is usable")
	}
	cmd := meta.List{Source: meta.SourceEDHREC, ID: "c", Format: meta.FormatCommander, Tier: meta.TierTypical, Cards: []meta.Card{{Name: "Plains", Count: 99}}}
	if Resolve(&cmd, w.idx, nil, mtgv1.FormatId_FORMAT_ID_COMMANDER).Usable() {
		t.Errorf("a Commander list with no commander is usable")
	}
}

func TestCommanderKeys(t *testing.T) {
	if PairKey("b", "a") != "a|b" || CommanderKey("b", "a") != "a+b" || CommanderKey("x") != "x" {
		t.Fatal("keys are not ordered")
	}
	fm := &FormatModel{Commanders: map[string]CommanderSignal{"a": {CEDH: 0.2}, "b": {CEDH: 0.6}, "a+c": {CEDH: 0.9}}}
	if s, ok := fm.Commander("c", "a"); !ok || s.CEDH != 0.9 {
		t.Errorf("pair row = %+v, %v", s, ok)
	}
	if s, ok := fm.Commander("a", "b"); !ok || s.CEDH != 0.6 {
		t.Errorf("pair with no row reads the stronger partner: %+v, %v", s, ok)
	}
	if _, ok := fm.Commander("z"); ok {
		t.Errorf("an unknown commander has a row")
	}
}

func TestPoolFloor(t *testing.T) {
	now := time.Date(2026, 9, 2, 0, 0, 0, 0, time.UTC)
	if got := PoolFloor(mtgv1.FormatId_FORMAT_ID_STANDARD, now); got != "2023-09-02" {
		t.Errorf("standard floor = %s", got)
	}
	if got := PoolFloor(mtgv1.FormatId_FORMAT_ID_MODERN, now); got != ModernFloor {
		t.Errorf("modern floor = %s", got)
	}
	if got := PoolFloor(mtgv1.FormatId_FORMAT_ID_COMMANDER, now); got != "" {
		t.Errorf("commander floor = %s", got)
	}
}
