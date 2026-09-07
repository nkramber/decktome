package quality

import (
	"context"
	"fmt"
	"hash/fnv"
	"log/slog"
	"math"
	"sort"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/profile"
)

// FitInput is everything a fit reads.
type FitInput struct {
	Index    *cards.Index
	Profiler *profile.Profiler
	Roles    Roles
	// Lists are every stored list, all formats together.
	Lists []meta.List
	// Commanders are the newest commander reads.
	Commanders []meta.Commander
	Now        time.Time
	Logger     *slog.Logger
	// Formats limits the fit, for the tests. Empty means the three.
	Formats []mtgv1.FormatId
}

// FitReport is what the gate document reads.
type FitReport struct {
	Formats map[string]*FormatReport
}

// FormatReport counts one format's fit.
type FormatReport struct {
	// Read, Unusable, and Used count the lists. OutOfPool counts the
	// 60-card products older than the format's pool (D-478).
	Read, Unusable, Used, OutOfPool int
	// Synthetic counts the bad lists the engine made.
	Synthetic int
	// Dropped names the features with no spread.
	Dropped []string
	Holdout Holdout
	// Iterations and Loss are the fit's own numbers.
	Iterations int
	Loss       float64
}

// The fit constants.
const (
	// HoldoutEvery holds out one list in five by the hash of its key.
	HoldoutEvery = 5
	// Iterations, LearningRate, and L2 are the gradient descent.
	Iterations   = 3000
	LearningRate = 0.05
	L2           = 0.01
	// MaxPairChecks caps the holdout pairs per bar.
	MaxPairChecks = 20000
	// FitHands is the goldfish hand count of the fit, a tenth of the
	// profiler's default, because the fit reads thousands of lists.
	FitHands = 1000
)

// Fit fits one scorer per format.
func Fit(ctx context.Context, in FitInput) (*Model, *FitReport, error) {
	if in.Index == nil || in.Profiler == nil {
		return nil, nil, fmt.Errorf("quality: the fit needs the card index and the profiler")
	}
	if in.Logger == nil {
		in.Logger = slog.Default()
	}
	if in.Now.IsZero() {
		in.Now = time.Now()
	}
	formats := in.Formats
	if len(formats) == 0 {
		formats = []mtgv1.FormatId{mtgv1.FormatId_FORMAT_ID_COMMANDER, mtgv1.FormatId_FORMAT_ID_STANDARD, mtgv1.FormatId_FORMAT_ID_MODERN}
	}
	in.Profiler.SetHands(FitHands)
	model := &Model{
		Version:      in.Now.UTC().Format("20060102T150405Z"),
		SnapshotAsOf: in.Index.AsOf.Format("2006-01-02"),
		Formats:      map[string]*FormatModel{},
	}
	rep := &FitReport{Formats: map[string]*FormatReport{}}
	for _, f := range formats {
		if err := ctx.Err(); err != nil {
			return nil, nil, err
		}
		word := meta.FormatWord(f)
		fm, fr, err := fitFormat(ctx, in, f)
		if err != nil {
			in.Logger.Warn("quality fit skipped a format", "format", word, "err", err)
			rep.Formats[word] = fr
			continue
		}
		model.Formats[word] = fm
		rep.Formats[word] = fr
	}
	if len(model.Formats) == 0 {
		return nil, rep, fmt.Errorf("quality: no format had enough lists to fit")
	}
	return model, rep, nil
}

// ModernFloor is the release date of Eighth Edition, the start of the
// Modern card pool (mtg-corpus section 2.1).
const ModernFloor = "2003-07-28"

// StandardYears is how far back a 60-card product serves the Standard
// baseline: the sets of the last three years, near the pool of a
// format that rotates once a year (D-478).
const StandardYears = 3

// PoolFloor is the earliest release date of a 60-card product that
// serves a format's baseline (D-478). Commander has no floor, because a
// Commander product is a Commander deck whatever its year.
func PoolFloor(f mtgv1.FormatId, now time.Time) string {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_STANDARD:
		return now.AddDate(-StandardYears, 0, 0).Format("2006-01-02")
	case mtgv1.FormatId_FORMAT_ID_MODERN:
		return ModernFloor
	default:
		return ""
	}
}

// MaxTierLists caps the lists of one tier in a fit, the newest first.
// Ninety days of cEDH held 21,146 lists on 2026-09-02, and the pair
// count of a fit grows with the square of a list's size times the
// list count (D-483).
const MaxTierLists = 4000

// capTiers keeps the newest MaxTierLists of each tier, and the order
// of the rest.
func capTiers(reals []*Resolved, limit int) []*Resolved {
	byTier := map[string][]*Resolved{}
	for _, r := range reals {
		byTier[r.List.Tier] = append(byTier[r.List.Tier], r)
	}
	keep := map[*Resolved]bool{}
	for _, rs := range byTier {
		if len(rs) <= limit {
			for _, r := range rs {
				keep[r] = true
			}
			continue
		}
		sort.SliceStable(rs, func(i, j int) bool {
			if rs[i].List.Date != rs[j].List.Date {
				return rs[i].List.Date > rs[j].List.Date
			}
			return rs[i].List.Key() < rs[j].List.Key()
		})
		for _, r := range rs[:limit] {
			keep[r] = true
		}
	}
	out := reals[:0:0]
	for _, r := range reals {
		if keep[r] {
			out = append(out, r)
		}
	}
	return out
}

// holdout says whether a key sits in the holdout split.
func holdout(key string) bool {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return h.Sum32()%HoldoutEvery == 0
}

type sample struct {
	r        *Resolved
	features map[string]float64
	level    int
	hold     bool
}

func fitFormat(ctx context.Context, in FitInput, f mtgv1.FormatId) (*FormatModel, *FormatReport, error) {
	fr := &FormatReport{}
	src := in.Index
	pl := newPool(in.Index, f)
	var reals []*Resolved
	floor := PoolFloor(f, in.Now)
	for i := range in.Lists {
		l := &in.Lists[i]
		if !meta.Covers(l.Format, f) || l.Tier == meta.TierBad {
			continue
		}
		if l.Format == meta.FormatSixty && l.Date < floor {
			fr.OutOfPool++
			continue
		}
		fr.Read++
		r := Resolve(l, in.Index, in.Roles, f)
		if !r.Usable() {
			fr.Unusable++
			continue
		}
		reals = append(reals, r)
	}
	reals = capTiers(reals, MaxTierLists)
	fr.Used = len(reals)
	if len(reals) < 20 {
		return nil, fr, fmt.Errorf("%d usable lists, and the fit wants 20", len(reals))
	}
	// A break draws its cards from the ones the real lists play (D-488).
	pl.setSeen(reals)
	// The bad rung sits under the precons by construction: every
	// baseline and typical list breaks on every axis, and a top list
	// breaks on none. A broken top list still holds the staples and the
	// pairs the top lists play, and it outranked a precon on the first
	// full fit (D-484).
	all := make([]*Resolved, 0, 6*len(reals))
	for _, r := range reals {
		all = append(all, r)
		if r.List.Tier != meta.TierBaseline && r.List.Tier != meta.TierTypical {
			continue
		}
		for _, axis := range Defects {
			all = append(all, Synthesize(r, axis, pl, in.Roles))
			fr.Synthetic++
		}
	}
	// The tiers the lists hold, worst first.
	present := map[string]bool{}
	for _, r := range all {
		present[r.List.Tier] = true
	}
	fm := &FormatModel{Format: meta.FormatWord(f), TrainCounts: map[string]int{}, HoldoutCounts: map[string]int{}}
	for _, t := range meta.Tiers {
		if present[t] {
			fm.Tiers = append(fm.Tiers, t)
		}
	}
	if len(fm.Tiers) < 2 {
		return nil, fr, fmt.Errorf("the lists hold one tier alone")
	}
	level := map[string]int{}
	for i, t := range fm.Tiers {
		level[t] = i
	}
	// The profiles, once per list.
	profiles := map[string]*mtgv1.DeckProfile{}
	for _, r := range all {
		if err := ctx.Err(); err != nil {
			return nil, fr, err
		}
		profiles[r.List.Key()] = in.Profiler.Measure(r.Deck, src)
	}
	// The corpus reads the training split alone, so the holdout says
	// how the scorer does on lists it never saw.
	var train []*Resolved
	for _, r := range reals {
		if !holdout(r.List.Key()) {
			train = append(train, r)
		}
	}
	buildCorpus(fm, train, profiles, in.Index)
	if f == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		commanderSignals(fm, in.Commanders, in.Index)
	}
	samples := make([]sample, 0, len(all))
	for _, r := range all {
		base := r.List.Key()
		if r.List.Source == meta.SourceSynthetic {
			base = base[:len(base)-len("/"+r.List.Defect)]
			base = base[len(meta.SourceSynthetic)+1:]
		}
		samples = append(samples, sample{
			r: r, level: level[r.List.Tier], hold: holdout(base),
			features: Features(Input{Deck: r.Deck, Profile: profiles[r.List.Key()], Cards: src}, fm),
		})
	}
	// The scaler, from the training split, and the keys with spread.
	var trainX, holdX [][]float64
	var trainY, holdY []int
	var holdDefects []string
	means, stds := map[string]float64{}, map[string]float64{}
	n := 0.0
	for _, s := range samples {
		if s.hold {
			continue
		}
		n++
		for _, k := range Keys {
			means[k] += s.features[k]
		}
	}
	for _, k := range Keys {
		means[k] /= n
	}
	for _, s := range samples {
		if s.hold {
			continue
		}
		for _, k := range Keys {
			d := s.features[k] - means[k]
			stds[k] += d * d
		}
	}
	for _, k := range Keys {
		sd := math.Sqrt(stds[k] / n)
		if sd < 1e-9 {
			fr.Dropped = append(fr.Dropped, k)
			continue
		}
		fm.Keys = append(fm.Keys, k)
		fm.Means = append(fm.Means, round4(means[k]))
		fm.Stds = append(fm.Stds, round4(sd))
	}
	if len(fm.Keys) == 0 {
		return nil, fr, fmt.Errorf("no feature has spread")
	}
	// The detector's rows: every real list as 0, the broken copies as 1.
	// The group word weighs the rows: the precons and the average decks
	// are the negatives the bar reads, and the rest keep the boundary
	// off the top lists.
	var defectTrainX, defectHoldX [][]float64
	var defectTrainY, defectHoldY []int
	var defectTrainGroup []int
	for _, s := range samples {
		z := fm.vector(s.features)
		if s.hold {
			holdX, holdY = append(holdX, z), append(holdY, s.level)
			holdDefects = append(holdDefects, s.r.List.Defect)
			fm.HoldoutCounts[s.r.List.Tier]++
		} else {
			trainX, trainY = append(trainX, z), append(trainY, s.level)
			fm.TrainCounts[s.r.List.Tier]++
		}
		label, ok := defectLabel(s.r.List)
		if !ok {
			continue
		}
		if s.hold {
			defectHoldX, defectHoldY = append(defectHoldX, z), append(defectHoldY, label)
		} else {
			defectTrainX, defectTrainY = append(defectTrainX, z), append(defectTrainY, label)
			defectTrainGroup = append(defectTrainGroup, defectGroup(s.r.List))
		}
	}
	w, th, loss := fitOrdinal(trainX, trainY, len(fm.Tiers))
	if len(defectTrainX) > 0 {
		dw, db := fitLogistic(defectTrainX, defectTrainY, defectTrainGroup)
		fm.DefectWeights = make([]float64, len(dw))
		for i, v := range dw {
			fm.DefectWeights[i] = round4(v)
		}
		fm.DefectBias = round4(db)
		fm.DefectThreshold = defectThreshold(fm, defectTrainX, defectTrainGroup)
	}
	fm.Weights = make([]float64, len(w))
	for i, v := range w {
		fm.Weights[i] = round4(v)
	}
	fm.Thresholds = make([]float64, len(th))
	for i, v := range th {
		fm.Thresholds[i] = round4(v)
	}
	fr.Iterations, fr.Loss = Iterations, round4(loss)
	fm.Holdout = evaluate(fm, holdX, holdY, holdDefects, level)
	if len(defectHoldX) > 0 {
		right := 0
		for i, z := range defectHoldX {
			if fm.flagged(z) == (defectHoldY[i] == 1) {
				right++
			}
		}
		fm.Holdout.DefectAccuracy = round4(float64(right) / float64(len(defectHoldX)))
	}
	fr.Holdout = fm.Holdout
	return fm, fr, nil
}

// defectLabel answers the detector's label of a list: 1 for a synthetic
// copy, and 0 for every real list. A detector that saw the precons
// alone read a cEDH list with 29 lands as a broken precon and cut the
// great-over-precon bar to 0.75 (gate run 2), so every real tier is a
// negative (D-485).
func defectLabel(l *meta.List) (int, bool) {
	if l.Source == meta.SourceSynthetic {
		return 1, true
	}
	return 0, true
}

// The detector's groups: the broken copies, the precons and the
// average decks, and the other real lists. Each group weighs a third
// of the fit whatever its count, because the real lists outnumber the
// copies four to one and the top lists outnumber the precons.
const (
	groupSynthetic = iota
	groupBase
	groupOther
	groupCount
)

// defectGroup answers the group of a list.
func defectGroup(l *meta.List) int {
	switch {
	case l.Source == meta.SourceSynthetic:
		return groupSynthetic
	case l.Tier == meta.TierBaseline || l.Tier == meta.TierTypical:
		return groupBase
	default:
		return groupOther
	}
}

// fitLogistic fits the defect detector by gradient descent, each group
// at a third of the weight.
func fitLogistic(x [][]float64, y, groups []int) (w []float64, bias float64) {
	dims := len(x[0])
	w = make([]float64, dims)
	n := float64(len(x))
	counts := make([]float64, groupCount)
	for _, g := range groups {
		counts[g]++
	}
	weight := make([]float64, groupCount)
	present := 0.0
	for _, c := range counts {
		if c > 0 {
			present++
		}
	}
	for g, c := range counts {
		if c > 0 {
			weight[g] = n / (present * c)
		}
	}
	for it := 0; it < Iterations; it++ {
		gw := make([]float64, dims)
		gb := 0.0
		for i, z := range x {
			eta := bias
			for d, v := range z {
				eta += w[d] * v
			}
			g := (sigmoid(eta) - float64(y[i])) * weight[groups[i]]
			for d, v := range z {
				gw[d] += g * v
			}
			gb += g
		}
		for d := range w {
			w[d] -= LearningRate * (gw[d]/n + L2*w[d])
		}
		bias -= LearningRate * gb / n
	}
	return w, bias
}

// defectThreshold picks the cut between the precons and their copies:
// the probability, in steps of a fiftieth, that gets the most of the
// two groups right on the training rows, each group at half the
// weight. The other real lists take no part, because the bar reads
// the precons.
func defectThreshold(fm *FormatModel, x [][]float64, groups []int) float64 {
	best, bestScore := 0.5, -1.0
	probs := make([]float64, len(x))
	for i, z := range x {
		probs[i] = fm.defect(z)
	}
	for step := 10; step <= 45; step++ {
		t := float64(step) / 50
		var baseRight, baseAll, synRight, synAll float64
		for i := range x {
			switch groups[i] {
			case groupBase:
				baseAll++
				if probs[i] < t {
					baseRight++
				}
			case groupSynthetic:
				synAll++
				if probs[i] >= t {
					synRight++
				}
			}
		}
		if baseAll == 0 || synAll == 0 {
			return 0.5
		}
		score := baseRight/baseAll + synRight/synAll
		if score > bestScore {
			best, bestScore = t, score
		}
	}
	return best
}

// flagged says whether the detector reads a standardized vector as a
// broken deck.
func (fm *FormatModel) flagged(z []float64) bool {
	t := fm.DefectThreshold
	if t <= 0 {
		t = 0.5
	}
	return len(fm.DefectWeights) == len(z) && fm.defect(z) >= t
}

// defect answers the detector's probability that a standardized vector
// is a broken deck, 0 with no detector.
func (fm *FormatModel) defect(z []float64) float64 {
	if len(fm.DefectWeights) != len(z) {
		return 0
	}
	eta := fm.DefectBias
	for i, v := range z {
		eta += fm.DefectWeights[i] * v
	}
	return sigmoid(eta)
}

// DefectFloor is the share of the score range a broken deck stays
// under. A deck the detector flags scores inside the bottom band, and
// every other deck above it, so the grade orders a precon over its
// broken copy whenever the detector reads both right (D-485).
const DefectFloor = 0.15

// grade answers the tier probabilities and the score of a standardized
// vector. The ladder gives the probabilities, the defect probability
// moves mass to the bottom rung, and the score is the expected rung
// over the ladder, placed above DefectFloor for a deck the detector
// passes and under it for one it flags (D-485).
func (fm *FormatModel) grade(z []float64) (p []float64, score float64) {
	p = levelProbabilities(fm.Weights, fm.Thresholds, z)
	d := fm.defect(z)
	for k := range p {
		p[k] *= 1 - d
	}
	p[0] += d
	exp := 0.0
	for k, v := range p {
		exp += float64(k) * v
	}
	ladder := 0.0
	if len(fm.Tiers) > 1 {
		ladder = exp / float64(len(fm.Tiers)-1)
	}
	if fm.flagged(z) {
		return p, DefectFloor * (1 - d)
	}
	return p, DefectFloor + (1-DefectFloor)*ladder
}

// sigmoid is the logistic function.
func sigmoid(x float64) float64 { return 1 / (1 + math.Exp(-x)) }

// thresholds reads the cut points from the free parameters: the first
// cut, then a positive gap per further cut, so the cuts stay ordered.
func thresholds(params []float64) []float64 {
	out := make([]float64, len(params))
	out[0] = params[0]
	for i := 1; i < len(params); i++ {
		out[i] = out[i-1] + math.Exp(params[i])
	}
	return out
}

// levelProbabilities answers the probability of each rung for one
// standardized vector.
func levelProbabilities(w, th, z []float64) []float64 {
	eta := 0.0
	for i, v := range z {
		eta += w[i] * v
	}
	levels := len(th) + 1
	p := make([]float64, levels)
	prev := 0.0
	for k := 0; k < levels; k++ {
		cum := 1.0
		if k < len(th) {
			cum = sigmoid(th[k] - eta)
		}
		p[k] = math.Max(cum-prev, 1e-12)
		prev = cum
	}
	return p
}

// fitOrdinal fits a proportional odds model by gradient descent. It
// answers the weights, the thresholds, and the final mean loss. Every
// rung weighs the same whatever its count: 4,000 great and 4,000 good
// Commander lists against 195 typical ones let the tournament lists
// set the whole ladder, and every casual deck fell to the bottom
// (D-488).
func fitOrdinal(x [][]float64, y []int, levels int) (w, th []float64, loss float64) {
	if len(x) == 0 {
		return nil, nil, 0
	}
	dims := len(x[0])
	counts := make([]float64, levels)
	for _, k := range y {
		counts[k]++
	}
	present := 0.0
	for _, c := range counts {
		if c > 0 {
			present++
		}
	}
	rowWeight := make([]float64, levels)
	for k, c := range counts {
		if c > 0 {
			rowWeight[k] = float64(len(x)) / (present * c)
		}
	}
	w = make([]float64, dims)
	params := make([]float64, levels-1)
	for i := 1; i < len(params); i++ {
		params[i] = math.Log(1.0)
	}
	// The first cut starts below zero, so the rungs spread over the
	// standardized range from the start.
	params[0] = -float64(levels-1) / 2
	n := float64(len(x))
	for it := 0; it < Iterations; it++ {
		th = thresholds(params)
		gw := make([]float64, dims)
		gp := make([]float64, len(params))
		loss = 0
		for i, z := range x {
			eta := 0.0
			for d, v := range z {
				eta += w[d] * v
			}
			k := y[i]
			var upper, lower, dUpper, dLower float64
			if k < len(th) {
				s := sigmoid(th[k] - eta)
				upper, dUpper = s, s*(1-s)
			} else {
				upper = 1
			}
			if k > 0 {
				s := sigmoid(th[k-1] - eta)
				lower, dLower = s, s*(1-s)
			}
			p := math.Max(upper-lower, 1e-12)
			rw := rowWeight[k]
			loss -= rw * math.Log(p)
			// d(-log p)/d eta = (dUpper - dLower) / p, and the cuts take
			// the other sign.
			dEta := rw * (dUpper - dLower) / p
			for d, v := range z {
				gw[d] += dEta * v
			}
			// A cut j moves every threshold at or above j through the
			// gap parameters, and the first parameter moves them all.
			if k < len(th) {
				addCutGradient(gp, params, k, -rw*dUpper/p)
			}
			if k > 0 {
				addCutGradient(gp, params, k-1, rw*dLower/p)
			}
		}
		for d := range w {
			gw[d] = gw[d]/n + L2*w[d]
			w[d] -= LearningRate * gw[d]
		}
		for j := range params {
			params[j] -= LearningRate * gp[j] / n
		}
		loss /= n
	}
	th = thresholds(params)
	return w, th, loss
}

// addCutGradient adds the gradient of a loss term in threshold j to the
// free parameters. Threshold j is the first parameter plus the gaps 1
// to j, and a gap is the exp of its parameter, so the chain rule
// scales the term by the gap.
func addCutGradient(gp, params []float64, j int, g float64) {
	gp[0] += g
	for m := 1; m <= j; m++ {
		gp[m] += g * math.Exp(params[m])
	}
}

// evaluate measures the holdout: the pair bars of the gate, the
// accuracy, and the confusion table.
func evaluate(fm *FormatModel, x [][]float64, y []int, defects []string, level map[string]int) Holdout {
	h := Holdout{Lists: len(x)}
	levels := len(fm.Tiers)
	h.Confusion = make([][]int, levels)
	for i := range h.Confusion {
		h.Confusion[i] = make([]int, levels)
	}
	scores := make([]float64, len(x))
	right := 0
	for i, z := range x {
		p, score := fm.grade(z)
		best := 0
		for k, v := range p {
			if v > p[best] {
				best = k
			}
		}
		scores[i] = score
		h.Confusion[y[i]][best]++
		if best == y[i] {
			right++
		}
	}
	if len(x) > 0 {
		h.Accuracy = round4(float64(right) / float64(len(x)))
	}
	pairBar := func(upper, lower string) PairShare {
		u, uok := level[upper]
		l, lok := level[lower]
		var ps PairShare
		if !uok || !lok {
			return ps
		}
		var ui, li []int
		for i, lv := range y {
			switch lv {
			case u:
				ui = append(ui, i)
			case l:
				li = append(li, i)
			}
		}
		sort.Ints(ui)
		sort.Ints(li)
		for _, a := range ui {
			for _, b := range li {
				if ps.Pairs >= MaxPairChecks {
					return ps
				}
				ps.Pairs++
				if scores[a] > scores[b] {
					ps.Wins++
				}
			}
		}
		return ps
	}
	h.GreatOverBaseline = pairBar(meta.TierGreat, meta.TierBaseline)
	h.BaselineOverBad = pairBar(meta.TierBaseline, meta.TierBad)
	// The precon bar per broken axis, so a reader sees which defect the
	// detector misses.
	if u, ok := level[meta.TierBaseline]; ok {
		if l, ok := level[meta.TierBad]; ok {
			h.BadByDefect = map[string]PairShare{}
			for _, axis := range Defects {
				var ps PairShare
				for a, la := range y {
					if la != u {
						continue
					}
					for b, lb := range y {
						if lb != l || defects[b] != axis || ps.Pairs >= MaxPairChecks {
							continue
						}
						ps.Pairs++
						if scores[a] > scores[b] {
							ps.Wins++
						}
					}
				}
				if ps.Pairs > 0 {
					h.BadByDefect[axis] = ps
				}
			}
		}
	}
	return h
}
