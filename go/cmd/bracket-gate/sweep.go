package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math"
	"slices"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/profile"
)

// sweepFlag asks for the free sweep of PR-45b (D-707, D-710).
var sweepFlag = flag.Bool("sweep", false, "build each shortlist at each point of the rate grid, print the table, and call no provider (D-707)")

// The axes of the grid. A rate is at most 1, so a keep rate of 2 keeps no
// card, and the first point is the shortlist before PR-45b.
var (
	sweepWeights = []float64{0.1, 0.2, 0.3, 0.5, 1}
	sweepKeeps   = []float64{2, 0.5, 0.3, 0.2, 0.1, 0.05}
)

// sweepPoints are the points of the grid: each weight at each keep rate,
// with no pin and with a pin. A pin reads nothing with no keep rate, so no
// such point runs (D-710).
func sweepPoints() []candidates.PowerRate {
	var out []candidates.PowerRate
	for _, wt := range sweepWeights {
		for _, k := range sweepKeeps {
			out = append(out, candidates.PowerRate{Weight: wt, Keep: k})
			if k < 1 {
				out = append(out, candidates.PowerRate{Weight: wt, Keep: k, Pin: true})
			}
		}
	}
	return out
}

// detailPoints are the points the card table reads.
var detailPoints = []candidates.PowerRate{
	{Weight: 0.1, Keep: 0.3, Pin: true},
	{Weight: 0.3, Keep: 0.3, Pin: true},
	{Weight: 0.3, Keep: 0.3},
	{Weight: 0.5, Keep: 0.3},
	{Weight: 1, Keep: 0.3},
}

// detailCards is how many power cards the card table lists.
const detailCards = 15

// staples are the roles a card with no theme signal filled before PR-45b.
var staples = map[mtgv1.CardRole]bool{
	mtgv1.CardRole_CARD_ROLE_LAND: true, mtgv1.CardRole_CARD_ROLE_RAMP: true, mtgv1.CardRole_CARD_ROLE_DRAW: true,
	mtgv1.CardRole_CARD_ROLE_REMOVAL: true, mtgv1.CardRole_CARD_ROLE_WIPE: true, mtgv1.CardRole_CARD_ROLE_INTERACTION: true,
}

// sweepPoint is what one point of the grid builds for one prompt. A fixing
// land makes two or more colors of the deck.
type sweepPoint struct {
	cards    int
	in       int
	themeOut int
	kept     int
	pinned   int
	other    int
	lands    int
	fixing   int
	power    map[string]int
}

// sweepTotal sums one point of the grid over the prompts.
type sweepTotal struct {
	prompts, of, met, want, in, themeOut, fixingOut int
}

// runSweep builds the shortlist of each prompt at each point of the grid,
// and writes the table that sets DefaultPowerRate (D-707, D-710). It calls
// no provider.
func runSweep(ctx context.Context, prompts []prompt, w io.Writer) error {
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(ctx, quiet)
	if err != nil {
		return err
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	scorer, err := gatekit.Scorer(ctx)
	if err != nil {
		return err
	}
	format := mtgv1.FormatId_FORMAT_ID_COMMANDER
	if scorer.MetaBoost(format) == nil {
		return fmt.Errorf("sweep: no stored quality model, so no top-list rate")
	}
	bands, err := profile.LoadBands()
	if err != nil {
		return err
	}
	of := profile.PowerOf(idx.Tags())
	_, _ = fmt.Fprintf(w, "# PR-45b: the rate sweep\n\nRun date: %s. Card snapshot: `%s`. Quality model: `%s`.\n",
		time.Now().UTC().Format("2006-01-02"), idx.AsOf.UTC().Format("20060102T150405"), scorer.Model().Version)
	totals := map[candidates.PowerRate]*sweepTotal{}
	points := sweepPoints()
	for _, p := range prompts {
		_, _ = fmt.Fprintf(w, "\n## Prompt %d: bracket %d %s (%s)\n\n", p.ID, p.Bracket, p.Commander, p.Theme)
		floors := bands.PowerFloors(format, gatekit.PowerLevel(p.Bracket, ""))
		if len(floors) == 0 {
			_, _ = fmt.Fprintln(w, "No power floor, so the grid moves nothing.")
			continue
		}
		c, ok := idx.ByName(p.Commander)
		if !ok {
			return fmt.Errorf("no card named %q", p.Commander)
		}
		// The request of build, at the rate before PR-45b. The shortlist
		// reads the top-list rate as the app does (F-129, D-706).
		req := candidates.Request{
			MetaBoost:          scorer.MetaBoost(format),
			PowerRate:          candidates.PowerRate{Weight: 0.1, Keep: 2},
			Format:             format,
			Colors:             c.GetColorIdentity(),
			Theme:              p.Theme,
			CommanderOracleIDs: []string{c.GetOracleId()},
			PoolRule:           mtgv1.PoolRule_POOL_RULE_ANY_CARD,
			Bracket:            p.Bracket,
		}
		base, err := cb.Build(idx, req)
		if err != nil {
			return err
		}
		wide, err := buildAt(cb, idx, req, candidates.PowerRate{Weight: 0.1, Keep: 1e-9}, wideLimits())
		if err != nil {
			return err
		}
		reach, rates := reachOf(wide, of)
		basePt := measurePoint(base, base, of)
		_, _ = fmt.Fprintf(w, "Floors: %s. Reach with a rate: %s.\n\n", floorWords(floors), countWords(reach))
		_, _ = fmt.Fprintf(w, "Rates of the power cards with a rate: %s.\n\n", quantileWords(rates))
		if err := writeTopPower(w, cb, idx, req, wide, base, of); err != nil {
			return err
		}
		_, _ = fmt.Fprintln(w, "| Weight | Keep | Pin | Cards | Lands | Fixing | In | Theme out | Kept | Pinned | Other | Tutors | Fast mana | Game Changers | Floors met |")
		_, _ = fmt.Fprintln(w, "|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|")
		for _, rate := range points {
			l, err := buildAt(cb, idx, req, rate, candidates.Limits{})
			if err != nil {
				return err
			}
			pt := measurePoint(l, base, of)
			met, want := floorsMet(pt.power, floors, reach)
			_, _ = fmt.Fprintf(w, "| %g | %s | %s | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d of %d |\n",
				rate.Weight, keepWords(rate.Keep), yes(rate.Pin), pt.cards, pt.lands, pt.fixing, pt.in, pt.themeOut, pt.kept, pt.pinned, pt.other,
				pt.power[profile.KeyTutor], pt.power[profile.KeyFastMana], pt.power[profile.KeyGameChanger], met, want)
			t := totals[rate]
			if t == nil {
				t = &sweepTotal{}
				totals[rate] = t
			}
			t.of++
			t.met += met
			t.want += want
			t.in += pt.in
			t.themeOut += pt.themeOut
			t.fixingOut += max(basePt.fixing-pt.fixing, 0)
			if met == want {
				t.prompts++
			}
		}
	}
	if len(totals) == 0 {
		return nil
	}
	_, _ = fmt.Fprintln(w, "\n## Summary")
	_, _ = fmt.Fprintln(w, "\n| Weight | Keep | Pin | Prompts at every floor | Floors met | Cards in | Theme out | Fixing lands out |")
	_, _ = fmt.Fprintln(w, "|---|---|---|---|---|---|---|---|")
	for _, r := range points {
		t := totals[r]
		_, _ = fmt.Fprintf(w, "| %g | %s | %s | %d of %d | %d of %d | %d | %d | %d |\n",
			r.Weight, keepWords(r.Keep), yes(r.Pin), t.prompts, t.of, t.met, t.want, t.in, t.themeOut, t.fixingOut)
	}
	return nil
}

// writeTopPower writes the power cards with the highest rates: the role
// each one takes, whether the theme matched it, and which lists hold it.
// The lists are the list before PR-45b and each detail point.
func writeTopPower(w io.Writer, cb *candidates.Builder, idx *cards.Index, req candidates.Request,
	wide, base *candidates.List, of func(*mtgv1.Card) []string) error {
	lists := []*candidates.List{base}
	header := "| Card | Rate | Role | Theme | Counts toward | Before |"
	for _, p := range detailPoints {
		l, err := buildAt(cb, idx, req, p, candidates.Limits{})
		if err != nil {
			return err
		}
		lists = append(lists, l)
		header += " " + pointWords(p) + " |"
	}
	var top []candidates.Candidate
	for _, c := range wide.Candidates {
		if c.Rate > 0 && len(of(c.Card)) > 0 {
			top = append(top, c)
		}
	}
	sort.SliceStable(top, func(i, j int) bool { return top[i].Rate > top[j].Rate })
	top = top[:min(len(top), detailCards)]
	_, _ = fmt.Fprintln(w, header)
	_, _ = fmt.Fprintln(w, "|---|---|---|---|---|---|"+strings.Repeat("---|", len(detailPoints)))
	for _, c := range top {
		row := fmt.Sprintf("| %s | %.3f | %s | %s | %s |", c.Card.GetName(), c.Rate, candidates.RoleName(c.Role), yes(c.Themed), keyWords(of(c.Card)))
		for _, l := range lists {
			row += " " + yes(holds(l, c.Card.GetOracleId())) + " |"
		}
		_, _ = fmt.Fprintln(w, row)
	}
	_, _ = fmt.Fprintln(w)
	return nil
}

// buildAt builds the shortlist of req at one rate and one set of limits.
func buildAt(cb *candidates.Builder, idx *cards.Index, req candidates.Request, rate candidates.PowerRate, lim candidates.Limits) (*candidates.List, error) {
	req.PowerRate = rate
	req.Limits = lim
	return cb.Build(idx, req)
}

// wideLimits lifts every cap, so the list holds every card the keep rate
// lets through.
func wideLimits() candidates.Limits {
	per := map[mtgv1.CardRole]int{}
	for v := range mtgv1.CardRole_name {
		per[mtgv1.CardRole(v)] = 1 << 20
	}
	return candidates.Limits{Total: 1 << 20, Upgrades: 1, PerRole: per}
}

// reachOf counts the power cards with a rate that the colors reach, and
// gathers their rates.
func reachOf(l *candidates.List, of func(*mtgv1.Card) []string) (map[string]int, []float64) {
	reach := map[string]int{}
	var rates []float64
	for _, c := range l.Candidates {
		keys := of(c.Card)
		if c.Rate <= 0 || len(keys) == 0 {
			continue
		}
		rates = append(rates, c.Rate)
		for _, key := range keys {
			reach[key]++
		}
	}
	return reach, rates
}

// measurePoint reads one shortlist against the shortlist before PR-45b.
func measurePoint(l, base *candidates.List, of func(*mtgv1.Card) []string) sweepPoint {
	pt := sweepPoint{power: map[string]int{}}
	inBase := map[string]bool{}
	var themed []string
	for _, c := range base.Candidates {
		id := c.Card.GetOracleId()
		inBase[id] = true
		if c.Themed {
			themed = append(themed, id)
		}
	}
	held := map[string]bool{}
	for _, c := range l.Candidates {
		id := c.Card.GetOracleId()
		held[id] = true
		pt.cards++
		if !inBase[id] {
			pt.in++
		}
		if !c.Themed && !staples[c.Role] && slices.Contains(c.Signals, "top-list rate") {
			pt.kept++
		}
		if c.Pinned {
			pt.pinned++
		}
		switch c.Role {
		case mtgv1.CardRole_CARD_ROLE_OTHER:
			pt.other++
		case mtgv1.CardRole_CARD_ROLE_LAND:
			pt.lands++
			if c.Fix >= 2 {
				pt.fixing++
			}
		}
		for _, key := range of(c.Card) {
			pt.power[key]++
		}
	}
	for _, id := range themed {
		if !held[id] {
			pt.themeOut++
		}
	}
	return pt
}

// floorsMet counts the floors a shortlist meets. A floor counts as met
// when the list holds the floor, or every power card with a rate that
// the colors reach.
func floorsMet(power map[string]int, floors map[string]float64, reach map[string]int) (met, want int) {
	for key, floor := range floors {
		want++
		if float64(power[key]) >= math.Min(floor, float64(reach[key])) {
			met++
		}
	}
	return met, want
}

func holds(l *candidates.List, id string) bool {
	for _, c := range l.Candidates {
		if c.Card.GetOracleId() == id {
			return true
		}
	}
	return false
}

func yes(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func pointWords(p candidates.PowerRate) string {
	s := fmt.Sprintf("%g, keep %s", p.Weight, keepWords(p.Keep))
	if p.Pin {
		s += ", pin"
	}
	return s
}

var powerNouns = map[string]string{
	profile.KeyTutor:       "tutors",
	profile.KeyFastMana:    "fast mana",
	profile.KeyGameChanger: "Game Changers",
}

func keyWords(keys []string) string {
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, powerNouns[key])
	}
	return strings.Join(parts, ", ")
}

func floorWords(floors map[string]float64) string {
	var parts []string
	for _, key := range profile.PowerKeys {
		if f, ok := floors[key]; ok {
			parts = append(parts, fmt.Sprintf("%s %g", powerNouns[key], f))
		}
	}
	return strings.Join(parts, ", ")
}

func countWords(counts map[string]int) string {
	var parts []string
	for _, key := range profile.PowerKeys {
		parts = append(parts, fmt.Sprintf("%s %d", powerNouns[key], counts[key]))
	}
	return strings.Join(parts, ", ")
}

func quantileWords(rates []float64) string {
	if len(rates) == 0 {
		return "none"
	}
	sorted := slices.Clone(rates)
	sort.Sort(sort.Reverse(sort.Float64Slice(sorted)))
	at := func(q float64) float64 { return sorted[int(q*float64(len(sorted)-1))] }
	return fmt.Sprintf("highest %.3f, upper quarter %.3f, median %.3f, lower quarter %.3f, over %d cards",
		sorted[0], at(0.25), at(0.5), at(0.75), len(sorted))
}

func keepWords(k float64) string {
	if k >= 1 {
		return "none"
	}
	return fmt.Sprintf("%g", k)
}
