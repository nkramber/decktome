package profile

import (
	"context"
	"errors"
	"sort"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/rules"
	"github.com/nkramber/decktome/go/internal/spellbook"
)

// Floor is the lowest bracket whose rules a Commander deck passes: the
// Game Changer limit and each content rule of brackets.json. The rules
// set no ceiling, so a floor is not the bracket of the deck. The judge
// calibration scores a precon against it (F-162, D-793). Commander
// Spellbook reads the deck once, and every bracket reads that answer.
func (p *Profiler) Floor(ctx context.Context, deck *mtgv1.Deck, src rules.CardSource) (int32, error) {
	out, _, entries, commanders := p.measure(deck, src)
	var changers float64
	for _, row := range out.GetFeatures() {
		if row.GetKey() == KeyGameChanger {
			changers = row.GetValue()
		}
	}
	once := &onceClassifier{inner: p.classify}
	q := *p
	q.classify = once
	brackets := make([]int32, 0, len(p.rules.Brackets))
	for b := range p.rules.Brackets {
		brackets = append(brackets, b)
	}
	sort.Slice(brackets, func(i, j int) bool { return brackets[i] < brackets[j] })
	for _, b := range brackets {
		br := p.rules.Brackets[b]
		if br.MaxGameChangers >= 0 && changers > float64(br.MaxGameChangers) {
			continue
		}
		check, _, forbidden := q.content(ctx, entries, commanders, b, nil)
		if !check.GetChecked() {
			return 0, errors.New("the content check did not run: " + check.GetError())
		}
		if forbidden.Empty() {
			return b, nil
		}
	}
	return 0, errors.New("no bracket allows the deck")
}

// onceClassifier asks its inner classifier one time and answers every
// later call from that answer.
type onceClassifier struct {
	inner Classifier
	done  bool
	res   *spellbook.Result
	err   error
}

func (o *onceClassifier) EstimateBracket(ctx context.Context, commanders, main []string) (*spellbook.Result, error) {
	if !o.done {
		o.done = true
		if o.inner == nil {
			o.err = errors.New("no classifier is wired")
		} else {
			o.res, o.err = o.inner.EstimateBracket(ctx, commanders, main)
		}
	}
	return o.res, o.err
}
