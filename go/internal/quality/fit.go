package quality

import (
	"context"
	"fmt"
	"hash/fnv"
	"log/slog"
	"math"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"
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
	// Folds is how many folds the bars read, FoldCount when zero.
	Folds int
	// SynergyFloor is the fall a synergy break needs, in standard
	// deviations, DefaultSynergyFloor when zero. A negative floor keeps
	// every copy, which is the fit before the check (D-652).
	SynergyFloor float64
	// SynergyFormats names the formats the synergy check reads,
	// DefaultSynergyFormats when empty (D-653).
	SynergyFormats []mtgv1.FormatId
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
	// Synthetic counts the bad lists the engine made, and SynergyCopies
	// the ones among them that broke on synergy.
	Synthetic, SynergyCopies int
	// Immaterial counts the synergy copies the check dropped over the
	// folds, each copy once, by the axis the fit asked for. SynergyFloor
	// is the floor the check read, negative in a format the check does not
	// read (D-653), and SynergyUnit its standard deviation on fold 0
	// (D-652).
	Immaterial   map[string]int
	SynergyFloor float64
	SynergyUnit  float64
	// Dropped names the features with no spread.
	Dropped []string
	Holdout Holdout
	// Iterations and Loss are the fit's own numbers.
	Iterations int
	Loss       float64
	// Folds is the holdout read over every fold, each list once, and
	// FoldCount how many folds ran. The bars read Folds (M-7).
	Folds     Holdout
	FoldCount int
	// Diagnostic reads the misses of the precon bar per axis and per
	// precon over the folds.
	Diagnostic *Diagnostic
}

// The fit constants.
const (
	// HoldoutEvery holds out one list in five by the hash of its key.
	// FoldCount is how many such folds the bars read: every list is
	// holdout in one fold, so a rung of eight precons reads eight and
	// not one (F-52). The stored model is the fit of fold 0.
	HoldoutEvery = 5
	FoldCount    = 5
	// Iterations, LearningRate, and L2 are the gradient descent.
	Iterations   = 3000
	LearningRate = 0.05
	L2           = 0.01
	// MaxPairChecks caps the holdout pairs per bar. The pairs sample
	// evenly under it: every upper list meets the same number of lower
	// lists, drawn by a seeded permutation, so a bar never reads the
	// first lists alone (F-51).
	MaxPairChecks = 20000
	// FitHands is the goldfish hand count of the fit, a tenth of the
	// profiler's default, because the fit reads thousands of lists.
	FitHands = 1000
	// WorstPrecons is how many precons the diagnostic names.
	WorstPrecons = 10
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
	if in.SynergyFloor == 0 {
		in.SynergyFloor = DefaultSynergyFloor
	}
	if len(in.SynergyFormats) == 0 {
		in.SynergyFormats = DefaultSynergyFormats
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

// holdoutFold says whether a key sits in the holdout split of a fold.
// The folds partition the keys, so every list is holdout in one fold.
func holdoutFold(key string, fold int) bool {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32()%HoldoutEvery) == fold
}

// holdout is the main split, fold 0, the one the stored model reads.
func holdout(key string) bool { return holdoutFold(key, 0) }

// baseKey is the key of the real list a copy came from, and the key
// itself for a real list. A copy's id is the real key, a slash, and
// the axis the fit asked for (F-55).
func baseKey(l *meta.List) string {
	if l.Source != meta.SourceSynthetic {
		return l.Key()
	}
	id := l.ID
	if i := strings.LastIndex(id, "/"); i >= 0 {
		id = id[:i]
	}
	return id
}

// listName names a list for a reader: the product or the event, else
// the id.
func listName(l *meta.List) string {
	if l.Event != "" {
		return l.Event
	}
	return l.ID
}

type sample struct {
	r        *Resolved
	features map[string]float64
	level    int
	hold     bool
}

// holdRow is one holdout list of a fold, as the bars and the
// diagnostic read it.
type holdRow struct {
	z      []float64
	level  int
	defect string
	key    string
	base   string
	name   string
	date   string
}

// prepared is one format's rows, resolved, capped, broken, and
// profiled once, so every fold reads the same lists (M-7).
type prepared struct {
	f        mtgv1.FormatId
	reals    []*Resolved
	all      []*Resolved
	tiers    []string
	level    map[string]int
	profiles map[string]*mtgv1.DeckProfile
}

// AxisRead reads the precon bar of one axis over every fold (M-7).
type AxisRead struct {
	// Pairs and Wins are the sampled pairs of the bar on this axis.
	Pairs, Wins int
	// The misses by kind. BothPassed: the detector passed both and the
	// ladder put the copy at or above the precon. PreconFlagged: it
	// flagged the precon and passed the copy. BothFlagged: it flagged
	// both and read the precon as the more broken.
	BothPassed, PreconFlagged, BothFlagged int
	// OwnPairs and OwnWins read each precon against its own copy, and
	// the own misses split by the same kinds.
	OwnPairs, OwnWins                               int
	OwnBothPassed, OwnPreconFlagged, OwnBothFlagged int
	// LeastMoved and MostMoved name three features each, by the mean
	// absolute standardized delta between a precon and its own copy.
	LeastMoved, MostMoved []string
	moved                 map[string]float64
	movedN                int
	// Audit reads every own-copy pair of this axis, and it answers
	// whether the bar asks for a true ordering (M-8, D-649).
	Audit []PairAudit
}

// PairAudit is one precon against its own broken copy, with what the
// break did to it (M-8, D-649).
//
// The bar asks the ladder to score a precon above its copy. The break
// replaces half the spells with cards the real lists play (D-488), and
// for a weak precon that is arguably the better pile of cards. This row
// carries what a reader needs to judge that: the two scores, the two
// grades, and the signed move of the corpus features. A positive move
// means the copy reads higher than the precon on that feature.
type PairAudit struct {
	Precon string  `json:"precon"`
	Axis   string  `json:"axis"`
	Won    bool    `json:"won"`
	Score  float64 `json:"precon_score"`
	Copy   float64 `json:"copy_score"`
	// Grade and CopyGrade are the tier each one graded, as words.
	Grade     string `json:"precon_grade"`
	CopyGrade string `json:"copy_grade"`
	// Rank is the precon's place among the holdout precons by score,
	// 0 the weakest. It answers whether the failures are the weak ones.
	Rank int `json:"rank"`
	// Moves are the signed standardized deltas, copy minus precon, of
	// the corpus features. A positive card_rate means the copy holds
	// cards the top lists play more than the precon does.
	Moves map[string]float64 `json:"moves"`
}

// PreconRead is one holdout precon over the sampled pairs of the
// precon bar.
type PreconRead struct {
	Name   string
	Date   string
	Grade  string
	Defect float64
	Pairs  int
	Misses int
}

// Diagnostic is the read of the misses, per axis and per precon, over
// every fold (M-7).
type Diagnostic struct {
	Axes map[string]*AxisRead
	// Precons names the ones that lose most, by miss share.
	Precons []PreconRead
	precons map[string]*PreconRead
}

func newDiagnostic() *Diagnostic {
	return &Diagnostic{Axes: map[string]*AxisRead{}, precons: map[string]*PreconRead{}}
}

func (d *Diagnostic) axis(name string) *AxisRead {
	ar := d.Axes[name]
	if ar == nil {
		ar = &AxisRead{moved: map[string]float64{}}
		d.Axes[name] = ar
	}
	return ar
}

// add folds another read into this one. A precon is holdout in one
// fold alone, so its row arrives once.
func (d *Diagnostic) add(o *Diagnostic) {
	for name, oa := range o.Axes {
		ar := d.axis(name)
		ar.Pairs += oa.Pairs
		ar.Wins += oa.Wins
		ar.BothPassed += oa.BothPassed
		ar.PreconFlagged += oa.PreconFlagged
		ar.BothFlagged += oa.BothFlagged
		ar.OwnPairs += oa.OwnPairs
		ar.OwnWins += oa.OwnWins
		ar.OwnBothPassed += oa.OwnBothPassed
		ar.OwnPreconFlagged += oa.OwnPreconFlagged
		ar.OwnBothFlagged += oa.OwnBothFlagged
		ar.movedN += oa.movedN
		for k, v := range oa.moved {
			ar.moved[k] += v
		}
		ar.Audit = append(ar.Audit, oa.Audit...)
	}
	for key, pr := range o.precons {
		have := d.precons[key]
		if have == nil {
			c := *pr
			d.precons[key] = &c
			continue
		}
		have.Pairs += pr.Pairs
		have.Misses += pr.Misses
	}
}

// finish names the features per axis and the precons that lose most.
func (d *Diagnostic) finish() *Diagnostic {
	for _, ar := range d.Axes {
		type moved struct {
			key string
			v   float64
		}
		var ms []moved
		for k, v := range ar.moved {
			ms = append(ms, moved{k, v})
		}
		sort.Slice(ms, func(i, j int) bool {
			if ms[i].v != ms[j].v {
				return ms[i].v < ms[j].v
			}
			return ms[i].key < ms[j].key
		})
		ar.LeastMoved, ar.MostMoved = nil, nil
		for i := 0; i < len(ms) && i < 3; i++ {
			ar.LeastMoved = append(ar.LeastMoved, ms[i].key)
			ar.MostMoved = append(ar.MostMoved, ms[len(ms)-1-i].key)
		}
	}
	d.Precons = d.Precons[:0]
	for _, pr := range d.precons {
		if pr.Pairs > 0 {
			d.Precons = append(d.Precons, *pr)
		}
	}
	share := func(pr PreconRead) float64 { return float64(pr.Misses) / float64(pr.Pairs) }
	sort.Slice(d.Precons, func(i, j int) bool {
		a, b := share(d.Precons[i]), share(d.Precons[j])
		if a != b {
			return a > b
		}
		return d.Precons[i].Name < d.Precons[j].Name
	})
	if len(d.Precons) > WorstPrecons {
		d.Precons = d.Precons[:WorstPrecons]
	}
	return d
}

// MovedShare is the mean absolute standardized delta of a feature
// between a precon and its own copy on this axis.
func (a *AxisRead) MovedShare(key string) float64 {
	if a.movedN == 0 {
		return 0
	}
	return a.moved[key] / float64(a.movedN)
}

// add folds another holdout read into this one: the pairs and the
// confusion sum, and the accuracies weigh by their rows.
func (h *Holdout) add(o Holdout) {
	total := h.Lists + o.Lists
	if total > 0 {
		h.Accuracy = round4((h.Accuracy*float64(h.Lists) + o.Accuracy*float64(o.Lists)) / float64(total))
	}
	dtotal := h.DefectLists + o.DefectLists
	if dtotal > 0 {
		h.DefectAccuracy = round4((h.DefectAccuracy*float64(h.DefectLists) + o.DefectAccuracy*float64(o.DefectLists)) / float64(dtotal))
	}
	h.Lists, h.DefectLists = total, dtotal
	h.GreatOverBaseline.Pairs += o.GreatOverBaseline.Pairs
	h.GreatOverBaseline.Wins += o.GreatOverBaseline.Wins
	h.BaselineOverBad.Pairs += o.BaselineOverBad.Pairs
	h.BaselineOverBad.Wins += o.BaselineOverBad.Wins
	h.BaselineOverOwn.Pairs += o.BaselineOverOwn.Pairs
	h.BaselineOverOwn.Wins += o.BaselineOverOwn.Wins
	if len(h.Confusion) == 0 {
		h.Confusion = make([][]int, len(o.Confusion))
		for i := range o.Confusion {
			h.Confusion[i] = append([]int(nil), o.Confusion[i]...)
		}
	} else {
		for i := range o.Confusion {
			for j := range o.Confusion[i] {
				h.Confusion[i][j] += o.Confusion[i][j]
			}
		}
	}
	if h.BadByDefect == nil {
		h.BadByDefect = map[string]PairShare{}
	}
	for axis, ps := range o.BadByDefect {
		c := h.BadByDefect[axis]
		c.Pairs += ps.Pairs
		c.Wins += ps.Wins
		h.BadByDefect[axis] = c
	}
	if h.OwnByDefect == nil {
		h.OwnByDefect = map[string]PairShare{}
	}
	for axis, ps := range o.OwnByDefect {
		c := h.OwnByDefect[axis]
		c.Pairs += ps.Pairs
		c.Wins += ps.Wins
		h.OwnByDefect[axis] = c
	}
}

// prepareFormat resolves the lists of a format, keeps the newest of
// each tier, breaks the baseline and the typical lists, and profiles
// every row once.
func prepareFormat(ctx context.Context, in FitInput, f mtgv1.FormatId) (*prepared, *FormatReport, error) {
	fr := &FormatReport{Immaterial: map[string]int{}, SynergyFloor: -1}
	if slices.Contains(in.SynergyFormats, f) {
		fr.SynergyFloor = in.SynergyFloor
	}
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
			c := Synthesize(r, axis, pl, in.Roles)
			all = append(all, c)
			fr.Synthetic++
			if c.List.Defect == DefectSynergy {
				fr.SynergyCopies++
			}
		}
	}
	// The tiers the lists hold, worst first.
	present := map[string]bool{}
	for _, r := range all {
		present[r.List.Tier] = true
	}
	prep := &prepared{f: f, reals: reals, all: all, level: map[string]int{}, profiles: map[string]*mtgv1.DeckProfile{}}
	for _, t := range meta.Tiers {
		if present[t] {
			prep.tiers = append(prep.tiers, t)
		}
	}
	if len(prep.tiers) < 2 {
		return nil, fr, fmt.Errorf("the lists hold one tier alone")
	}
	for i, t := range prep.tiers {
		prep.level[t] = i
	}
	// The profiles, once per row. A copy keeps its own key, so no copy
	// reads the profile of another (F-55).
	for _, r := range all {
		if err := ctx.Err(); err != nil {
			return nil, fr, err
		}
		prep.profiles[r.List.Key()] = in.Profiler.Measure(r.Deck, in.Index)
	}
	return prep, fr, nil
}

// fitFold fits one fold: the corpus over its training split, the
// scaler, the ladder, and the detector. It reads the holdout of the
// fold and answers the model and the diagnostic.
func fitFold(ctx context.Context, in FitInput, prep *prepared, fr *FormatReport, fold int) (*FormatModel, *Diagnostic, error) {
	src := in.Index
	fm := &FormatModel{Format: meta.FormatWord(prep.f), Tiers: prep.tiers, TrainCounts: map[string]int{}, HoldoutCounts: map[string]int{}}
	// The corpus reads the training split alone, so the holdout says
	// how the scorer does on lists it never saw.
	var train []*Resolved
	for _, r := range prep.reals {
		if !holdoutFold(r.List.Key(), fold) {
			train = append(train, r)
		}
	}
	buildCorpus(fm, train, prep.profiles, in.Index)
	if prep.f == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		commanderSignals(fm, in.Commanders, in.Index)
	}
	samples := make([]sample, 0, len(prep.all))
	for _, r := range prep.all {
		if err := ctx.Err(); err != nil {
			return nil, nil, err
		}
		samples = append(samples, sample{
			r: r, level: prep.level[r.List.Tier], hold: holdoutFold(baseKey(r.List), fold),
			features: Features(Input{Deck: r.Deck, Profile: prep.profiles[r.List.Key()], Cards: src}, fm),
		})
	}
	// The synergy check reads the corpus of this fold, and it drops a copy
	// before the scaler, the ladder, and the detector see it (D-652).
	samples, immaterial, unit := dropImmaterial(samples, fr.SynergyFloor)
	for axis, n := range immaterial {
		fr.Immaterial[axis] += n
	}
	if fold == 0 {
		fr.SynergyUnit = round4(unit)
	}
	// The scaler, from the training split, and the keys with spread.
	var trainX [][]float64
	var trainY []int
	var rows []holdRow
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
	var dropped []string
	for _, k := range Keys {
		sd := math.Sqrt(stds[k] / n)
		if sd < 1e-9 {
			dropped = append(dropped, k)
			continue
		}
		fm.Keys = append(fm.Keys, k)
		fm.Means = append(fm.Means, round4(means[k]))
		fm.Stds = append(fm.Stds, round4(sd))
	}
	if fold == 0 {
		fr.Dropped = dropped
	}
	if len(fm.Keys) == 0 {
		return nil, nil, fmt.Errorf("no feature has spread")
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
			rows = append(rows, holdRow{z: z, level: s.level, defect: s.r.List.Defect, key: s.r.List.Key(), base: baseKey(s.r.List), name: listName(s.r.List), date: s.r.List.Date})
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
	if fold == 0 {
		fr.Iterations, fr.Loss = Iterations, round4(loss)
	}
	var diag *Diagnostic
	fm.Holdout, diag = evaluate(fm, rows, prep.level, MaxPairChecks)
	if len(defectHoldX) > 0 {
		right := 0
		for i, z := range defectHoldX {
			if fm.flagged(z) == (defectHoldY[i] == 1) {
				right++
			}
		}
		fm.Holdout.DefectAccuracy = round4(float64(right) / float64(len(defectHoldX)))
		fm.Holdout.DefectLists = len(defectHoldX)
	}
	return fm, diag, nil
}

// fitFormat fits one format: the model on fold 0, and the bars over
// every fold, so each list is holdout once (M-7).
func fitFormat(ctx context.Context, in FitInput, f mtgv1.FormatId) (*FormatModel, *FormatReport, error) {
	prep, fr, err := prepareFormat(ctx, in, f)
	if err != nil {
		return nil, fr, err
	}
	folds := in.Folds
	if folds <= 0 {
		folds = FoldCount
	}
	diag := newDiagnostic()
	var fm *FormatModel
	for fold := 0; fold < folds; fold++ {
		m, fd, err := fitFold(ctx, in, prep, fr, fold)
		if err != nil {
			return nil, fr, err
		}
		if fold == 0 {
			fm = m
		}
		fr.Folds.add(m.Holdout)
		diag.add(fd)
	}
	fr.FoldCount = folds
	fr.Diagnostic = diag.finish()
	fr.Holdout = fr.Folds
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

// evaluate measures the holdout of a fold: the pair bars, the
// accuracy, the confusion table, and the diagnostic of the misses. A
// bar samples its pairs evenly under the limit: every upper list
// meets the same number of lower lists, drawn by a seeded permutation
// (F-51).
func evaluate(fm *FormatModel, rows []holdRow, level map[string]int, limit int) (Holdout, *Diagnostic) {
	h := Holdout{Lists: len(rows)}
	diag := newDiagnostic()
	levels := len(fm.Tiers)
	h.Confusion = make([][]int, levels)
	for i := range h.Confusion {
		h.Confusion[i] = make([]int, levels)
	}
	scores := make([]float64, len(rows))
	grades := make([]int, len(rows))
	flagged := make([]bool, len(rows))
	right := 0
	for i, r := range rows {
		p, score := fm.grade(r.z)
		best := 0
		for k, v := range p {
			if v > p[best] {
				best = k
			}
		}
		scores[i], grades[i], flagged[i] = score, best, fm.flagged(r.z)
		h.Confusion[r.level][best]++
		if best == r.level {
			right++
		}
	}
	if len(rows) > 0 {
		h.Accuracy = round4(float64(right) / float64(len(rows)))
	}
	byLevel := func(lv int, axis string) []int {
		var out []int
		for i, r := range rows {
			if r.level == lv && (axis == "" || r.defect == axis) {
				out = append(out, i)
			}
		}
		return out
	}
	bar := func(ui, li []int, seed string, visit func(a, b int, win bool)) PairShare {
		var ps PairShare
		if len(ui) == 0 || len(li) == 0 {
			return ps
		}
		per := limit / len(ui)
		if per < 1 {
			per = 1
		}
		if per > len(li) {
			per = len(li)
		}
		hash := fnv.New64a()
		_, _ = hash.Write([]byte(seed))
		rng := rand.New(rand.NewPCG(hash.Sum64(), 11))
		for _, a := range ui {
			lower := li
			if per < len(li) {
				perm := rng.Perm(len(li))[:per]
				lower = make([]int, per)
				for i, j := range perm {
					lower[i] = li[j]
				}
			}
			for _, b := range lower {
				ps.Pairs++
				win := scores[a] > scores[b]
				if win {
					ps.Wins++
				}
				if visit != nil {
					visit(a, b, win)
				}
			}
		}
		return ps
	}
	u, uok := level[meta.TierBaseline]
	g, gok := level[meta.TierGreat]
	l, lok := level[meta.TierBad]
	if gok && uok {
		h.GreatOverBaseline = bar(byLevel(g, ""), byLevel(u, ""), "great|baseline", nil)
	}
	if !uok || !lok {
		return h, diag
	}
	precons := byLevel(u, "")
	// The rank of each precon among the holdout precons by score, 0 the
	// weakest. M-8 asks whether the pairs that fail are the weak precons
	// (D-649).
	rankOf := map[int]int{}
	{
		byScore := append([]int(nil), precons...)
		sort.Slice(byScore, func(i, j int) bool { return scores[byScore[i]] < scores[byScore[j]] })
		for r, idx := range byScore {
			rankOf[idx] = r
		}
	}
	for _, a := range precons {
		diag.precons[rows[a].key] = &PreconRead{Name: rows[a].name, Date: rows[a].date, Grade: fm.Tiers[grades[a]], Defect: round4(fm.defect(rows[a].z))}
	}
	h.BaselineOverBad = bar(precons, byLevel(l, ""), "baseline|bad", func(a, _ int, win bool) {
		pr := diag.precons[rows[a].key]
		pr.Pairs++
		if !win {
			pr.Misses++
		}
	})
	// The precon bar per broken axis, so a reader sees which defect the
	// detector misses, and the misses by kind. Each precon also meets
	// its own copy, which says how far the break moved each feature.
	own := map[string][]int{}
	for i, r := range rows {
		if r.level == l {
			own[r.base] = append(own[r.base], i)
		}
	}
	h.BadByDefect = map[string]PairShare{}
	h.OwnByDefect = map[string]PairShare{}
	for _, axis := range Defects {
		li := byLevel(l, axis)
		if len(li) == 0 {
			continue
		}
		ar := diag.axis(axis)
		ps := bar(precons, li, "baseline|bad|"+axis, func(a, b int, win bool) {
			if win {
				return
			}
			switch {
			case !flagged[a] && !flagged[b]:
				ar.BothPassed++
			case flagged[a] && !flagged[b]:
				ar.PreconFlagged++
			case flagged[a] && flagged[b]:
				ar.BothFlagged++
			}
		})
		ar.Pairs, ar.Wins = ps.Pairs, ps.Wins
		h.BadByDefect[axis] = ps
		for _, a := range precons {
			for _, b := range own[rows[a].key] {
				if rows[b].defect != axis {
					continue
				}
				ar.OwnPairs++
				h.BaselineOverOwn.Pairs++
				switch {
				case scores[a] > scores[b]:
					ar.OwnWins++
					h.BaselineOverOwn.Wins++
				case !flagged[a] && !flagged[b]:
					ar.OwnBothPassed++
				case flagged[a] && !flagged[b]:
					ar.OwnPreconFlagged++
				case flagged[a] && flagged[b]:
					ar.OwnBothFlagged++
				}
				for i, k := range fm.Keys {
					ar.moved[k] += math.Abs(rows[b].z[i] - rows[a].z[i])
				}
				ar.movedN++
				// The audit of M-8: every own-copy pair, with the signed
				// move of the corpus features. The bar asks the ladder to
				// score the precon above the copy, and this row says
				// whether that ordering is the true one (D-649).
				au := PairAudit{
					Precon: rows[a].key, Axis: axis, Won: scores[a] > scores[b], Rank: rankOf[a],
					Score: round4(scores[a]), Copy: round4(scores[b]),
					Grade: fm.Tiers[grades[a]], CopyGrade: fm.Tiers[grades[b]],
					Moves: map[string]float64{},
				}
				for i, k := range fm.Keys {
					switch k {
					case KeyCardRate, KeyUnseenShare, KeySynergy:
						au.Moves[k] = round4(rows[b].z[i] - rows[a].z[i])
					}
				}
				ar.Audit = append(ar.Audit, au)
			}
		}
		if ar.OwnPairs > 0 {
			h.OwnByDefect[axis] = PairShare{Pairs: ar.OwnPairs, Wins: ar.OwnWins}
		}
	}
	return h, diag
}
