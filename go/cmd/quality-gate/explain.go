package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/quality"
	"github.com/nkramber/decktome/go/internal/rules"
)

// runExplain reads the decks of a deck gate document and prints how the
// stored model grades each one: the detector's probability against its
// cut, the rules read of a Commander deck, the ladder, and the six
// largest contributions. It is free, and it is how a session reads a
// grade it does not believe.
func runExplain(path, promptsPath string) error {
	ctx := context.Background()
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(ctx, quiet)
	if err != nil {
		return err
	}
	commanders, err := promptCommanders(promptsPath)
	if err != nil {
		return err
	}
	f, err := os.Open(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	decks, err := readDeckGate(f, idx, commanders)
	if err != nil {
		return err
	}
	scorer, err := gatekit.Scorer(ctx)
	if err != nil {
		return err
	}
	cfg, err := rules.Load()
	if err != nil {
		return err
	}
	prof, err := profile.New(cfg, func() *cards.TagIndex { return idx.Tags() }, nil)
	if err != nil {
		return err
	}
	builder, err := candidates.New()
	if err != nil {
		return err
	}
	roles := builder.Roles(idx)
	for _, d := range decks {
		for _, dc := range d.deck.GetCards() {
			if c, ok := idx.ByOracleID(dc.GetOracleId()); ok {
				dc.Role = roles(c)
			}
		}
		e := scorer.Explain(quality.Input{Deck: d.deck, Profile: prof.Measure(d.deck, idx), Cards: idx})
		if e == nil {
			fmt.Printf("%d. %s: no model\n", d.id, d.title)
			continue
		}
		fmt.Printf("%d. %s: %s, score %.2f, defect %.2f (cut %.2f, flagged %v), ladder %v\n", d.id, d.title, e.Tier, e.Score, e.Defect, e.Threshold, e.Flagged, rounded(e.Ladder))
		if r := e.Rules; r.Checked {
			fmt.Printf("   rules: lands %.0f of a need of %.1f (short %.1f, cut %.0f, cheap %d), curve %.2f (cut %.1f), colors %.2f (cut %.2f), flagged %v\n",
				r.LandCount, r.Need, r.Shortfall(), quality.RuleLandShortfall, r.Cheap, r.AvgManaValue, quality.RuleCurveCeiling, r.ColorSources, quality.RuleColorFloor, r.Flagged())
		}
		cs := e.Contributions
		sort.Slice(cs, func(i, j int) bool {
			return math.Abs(cs[i].Defect)+math.Abs(cs[i].Ladder) > math.Abs(cs[j].Defect)+math.Abs(cs[j].Ladder)
		})
		for i := 0; i < len(cs) && i < 6; i++ {
			c := cs[i]
			fmt.Printf("   %-24s value %8.3f  z %6.2f  ladder %6.2f  defect %6.2f\n", c.Key, c.Value, c.Z, c.Ladder, c.Defect)
		}
	}
	return nil
}

func rounded(v []float64) []float64 {
	out := make([]float64, len(v))
	for i, x := range v {
		out[i] = math.Round(x*100) / 100
	}
	return out
}
