package quality

import (
	"context"
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
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

// fitWorld fits the test world with the synergy check on at the default
// floor. The world is a Standard world, and the check reads Commander
// alone unless a fit names another format (D-653).
func fitWorld(t *testing.T) (*world, *Model, *FitReport) {
	t.Helper()
	return fitWorldAt(t, 0, mtgv1.FormatId_FORMAT_ID_STANDARD)
}

// fitWorldAt fits the test world at a synergy floor, the default at zero,
// with the check on in the named formats.
func fitWorldAt(t *testing.T, floor float64, checked ...mtgv1.FormatId) (*world, *Model, *FitReport) {
	t.Helper()
	w := newWorld(t)
	model, rep, err := Fit(context.Background(), FitInput{
		Index: w.idx, Profiler: w.profiler, Lists: w.corpus(),
		Now:            time.Date(2026, 9, 2, 15, 0, 0, 0, time.UTC),
		Formats:        []mtgv1.FormatId{mtgv1.FormatId_FORMAT_ID_STANDARD},
		SynergyFloor:   floor,
		SynergyFormats: checked,
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
	// Thirty precons break on five axes each (D-484). A precon of basics
	// breaks on synergy for the colors axis too, so 60 copies read synergy.
	if fr.Used != 100 || fr.Synthetic != 150 || fr.SynergyCopies != 60 || fr.Unusable != 0 || fr.OutOfPool != 1 {
		t.Errorf("report = %+v", fr)
	}
	if fr.SynergyFloor != DefaultSynergyFloor {
		t.Errorf("synergy floor = %.2f, want the default %.2f", fr.SynergyFloor, DefaultSynergyFloor)
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

// TestPairBarSamplesEvenly: every upper list meets the same number of
// lower lists under the limit, so a bar never reads the first lists
// alone (F-51).
func TestPairBarSamplesEvenly(t *testing.T) {
	fm := &FormatModel{
		Tiers: []string{meta.TierBad, meta.TierBaseline, meta.TierGreat},
		Keys:  []string{"a"}, Means: []float64{0}, Stds: []float64{1}, Weights: []float64{1}, Thresholds: []float64{-1, 1},
	}
	level := map[string]int{meta.TierBad: 0, meta.TierBaseline: 1, meta.TierGreat: 2}
	var rows []holdRow
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("p%d", i)
		rows = append(rows, holdRow{z: []float64{float64(i)}, level: 1, key: key, base: key, name: key})
	}
	for i := 0; i < 100; i++ {
		rows = append(rows, holdRow{z: []float64{-5}, level: 0, defect: DefectLands, key: fmt.Sprintf("b%d", i), base: fmt.Sprintf("p%d", i%5)})
	}
	h, diag := evaluate(fm, rows, level, 50)
	diag.finish()
	if h.BaselineOverBad.Pairs != 50 || h.BaselineOverBad.Wins != 50 {
		t.Errorf("baseline over bad = %+v, want 50 pairs of 5 precons at 10 each", h.BaselineOverBad)
	}
	if len(diag.Precons) != 5 {
		t.Fatalf("precons = %d, want every precon in the sample", len(diag.Precons))
	}
	for _, pr := range diag.Precons {
		if pr.Pairs != 10 || pr.Misses != 0 {
			t.Errorf("%s met %d lists with %d misses, want 10 and 0", pr.Name, pr.Pairs, pr.Misses)
		}
	}
	ar := diag.Axes[DefectLands]
	if ar == nil || ar.Pairs != 50 || ar.OwnPairs != 100 || ar.OwnWins != 100 || len(ar.LeastMoved) != 1 {
		t.Errorf("lands axis = %+v", ar)
	}
	if h.BaselineOverOwn.Pairs != 100 || h.BaselineOverOwn.Wins != 100 || h.OwnByDefect[DefectLands] != (PairShare{Pairs: 100, Wins: 100}) {
		t.Errorf("baseline over own = %+v, per axis %+v, want every precon over its copies", h.BaselineOverOwn, h.OwnByDefect)
	}
	// A limit above the pairs reads every pair once.
	h, _ = evaluate(fm, rows, level, 100000)
	if h.BaselineOverBad.Pairs != 500 {
		t.Errorf("unlimited pairs = %d, want 500", h.BaselineOverBad.Pairs)
	}
}

// TestFoldsPartition: every key is holdout in one fold alone, and the
// main split is fold 0 (F-52).
func TestFoldsPartition(t *testing.T) {
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("mtgjson:precon/%d", i)
		n := 0
		for f := 0; f < FoldCount; f++ {
			if holdoutFold(key, f) {
				n++
			}
		}
		if n != 1 {
			t.Fatalf("%s is holdout in %d folds", key, n)
		}
		if holdout(key) != holdoutFold(key, 0) {
			t.Fatalf("%s: the main split is not fold 0", key)
		}
	}
}

// TestSynthesizeKeepsTwoKeys: the copies-axis copy of a list that
// falls to synergy keeps its own key beside the synergy copy, and both
// name the real list as their base (F-55).
func TestSynthesizeKeepsTwoKeys(t *testing.T) {
	w := newWorld(t)
	l := meta.List{Source: meta.SourceMTGJSON, ID: "precon/two", Format: meta.FormatStandard, Date: "2026-09-01", Tier: meta.TierBaseline}
	l.Cards = append(l.Cards, meta.Card{Name: "Plains", Count: 12}, meta.Card{Name: "Island", Count: 12})
	l.Cards = append(l.Cards, meta.Card{Name: w.staples[0].GetName(), Count: 4}, meta.Card{Name: w.staples[1].GetName(), Count: 4})
	for _, c := range w.staples[2:30] {
		l.Cards = append(l.Cards, meta.Card{Name: c.GetName(), Count: 1})
	}
	r := Resolve(&l, w.idx, nil, mtgv1.FormatId_FORMAT_ID_STANDARD)
	if !r.Usable() {
		t.Fatalf("unusable: %v", r.Missing)
	}
	pl := newPool(w.idx, mtgv1.FormatId_FORMAT_ID_STANDARD)
	copies := Synthesize(r, DefectCopies, pl, nil)
	synergy := Synthesize(r, DefectSynergy, pl, nil)
	if copies.List.Defect != DefectSynergy || synergy.List.Defect != DefectSynergy {
		t.Errorf("defects = %s and %s, want synergy for both", copies.List.Defect, synergy.List.Defect)
	}
	if copies.List.Key() == synergy.List.Key() {
		t.Errorf("the two copies share the key %s", copies.List.Key())
	}
	if baseKey(copies.List) != l.Key() || baseKey(synergy.List) != l.Key() || baseKey(&l) != l.Key() {
		t.Errorf("bases = %s, %s, %s, want %s", baseKey(copies.List), baseKey(synergy.List), baseKey(&l), l.Key())
	}
}

// TestFoldsReadEveryList: the bars over the folds read every row as
// holdout once, and the diagnostic's kinds sum to the misses (M-7).
func TestFoldsReadEveryList(t *testing.T) {
	_, _, rep := fitWorld(t)
	fr := rep.Formats[meta.FormatStandard]
	if fr.FoldCount != FoldCount {
		t.Fatalf("folds = %d", fr.FoldCount)
	}
	// The folds count each copy the synergy check drops once, in the fold
	// that holds it out (D-652).
	kept := fr.Synthetic - fr.ImmaterialCount()
	if fr.Folds.Lists != fr.Used+kept {
		t.Errorf("the folds read %d rows, want %d", fr.Folds.Lists, fr.Used+kept)
	}
	if fr.Holdout.BaselineOverOwn.Pairs != kept {
		t.Errorf("baseline over own read %d pairs, want the %d copies the check kept", fr.Holdout.BaselineOverOwn.Pairs, kept)
	}
	ownSum := 0
	for _, ps := range fr.Holdout.OwnByDefect {
		ownSum += ps.Pairs
	}
	if ownSum != fr.Holdout.BaselineOverOwn.Pairs {
		t.Errorf("the own pairs per axis sum to %d, and the bar reads %d", ownSum, fr.Holdout.BaselineOverOwn.Pairs)
	}
	if fr.Holdout.BaselineOverBad.Pairs == 0 || fr.Holdout.BaselineOverBad.Share() < 0.95 {
		t.Errorf("baseline over bad over the folds = %+v", fr.Holdout.BaselineOverBad)
	}
	d := fr.Diagnostic
	if d == nil || len(d.Axes) == 0 {
		t.Fatal("no diagnostic")
	}
	for axis, ar := range d.Axes {
		if ar.Pairs-ar.Wins != ar.BothPassed+ar.PreconFlagged+ar.BothFlagged {
			t.Errorf("%s: %d misses, and the kinds sum to %d", axis, ar.Pairs-ar.Wins, ar.BothPassed+ar.PreconFlagged+ar.BothFlagged)
		}
		if ar.OwnPairs == 0 || len(ar.LeastMoved) != 3 || len(ar.MostMoved) != 3 {
			t.Errorf("%s: own pairs %d, least %v, most %v", axis, ar.OwnPairs, ar.LeastMoved, ar.MostMoved)
		}
		if ar.OwnPairs-ar.OwnWins != ar.OwnBothPassed+ar.OwnPreconFlagged+ar.OwnBothFlagged {
			t.Errorf("%s: %d own misses, and the kinds sum to %d", axis, ar.OwnPairs-ar.OwnWins, ar.OwnBothPassed+ar.OwnPreconFlagged+ar.OwnBothFlagged)
		}
	}
	if len(d.Precons) == 0 || len(d.Precons) > WorstPrecons {
		t.Errorf("precons named = %d", len(d.Precons))
	}
	for _, pr := range d.Precons {
		if pr.Pairs == 0 || pr.Name == "" || pr.Grade == "" {
			t.Errorf("precon row = %+v", pr)
		}
	}
}

// TestSynergyCheckShapesTheBadRung: the check reads Commander alone by
// default, so a Standard fit keeps every copy (D-653). At a floor no break
// reaches, every synergy copy leaves the fit, so a request that passes no
// check makes no copy at all (D-652).
func TestSynergyCheckShapesTheBadRung(t *testing.T) {
	_, _, unchecked := fitWorldAt(t, 0)
	fr := unchecked.Formats[meta.FormatStandard]
	if fr.SynergyFloor >= 0 || fr.ImmaterialCount() != 0 || fr.SynergyUnit != 0 || fr.Holdout.BaselineOverOwn.Pairs != 150 {
		t.Errorf("a Standard fit read the floor %.2f and dropped %d copies at a unit of %.4f over %d own pairs, want no check and 150 pairs", fr.SynergyFloor, fr.ImmaterialCount(), fr.SynergyUnit, fr.Holdout.BaselineOverOwn.Pairs)
	}
	_, _, all := fitWorldAt(t, 1000, mtgv1.FormatId_FORMAT_ID_STANDARD)
	fr = all.Formats[meta.FormatStandard]
	if fr.SynergyCopies == 0 {
		t.Fatal("the test world made no synergy copy, so the check reads nothing")
	}
	if fr.ImmaterialCount() != fr.SynergyCopies || fr.SynergyUnit <= 0 {
		t.Errorf("a floor of 1000 dropped %d of %d synergy copies at a unit of %.4f, want every one", fr.ImmaterialCount(), fr.SynergyCopies, fr.SynergyUnit)
	}
	if _, ok := fr.Holdout.OwnByDefect[DefectSynergy]; ok {
		t.Errorf("the synergy axis still reads own pairs: %+v", fr.Holdout.OwnByDefect)
	}
	if fr.Holdout.BaselineOverOwn.Pairs != fr.Synthetic-fr.SynergyCopies {
		t.Errorf("own pairs = %d, want the %d copies of the other axes", fr.Holdout.BaselineOverOwn.Pairs, fr.Synthetic-fr.SynergyCopies)
	}
}
