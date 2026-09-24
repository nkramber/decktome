package main

import (
	"fmt"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
)

func testIndex() *cards.Index {
	return cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-bolt", Name: "Lightning Bolt", TypeLine: "Instant"},
		{OracleId: "o-mountain", Name: "Mountain", TypeLine: "Basic Land - Mountain"},
	}, nil, nil, time.Unix(1000, 0).UTC())
}

func precon(kind string, n int, date string) meta.Precon {
	return meta.Precon{
		Name: fmt.Sprintf("%s %d", kind, n), Code: "P", Type: kind, ReleaseDate: date,
		Cards: []meta.PreconCard{{Name: "Lightning Bolt", Count: 4, OracleID: "o-bolt"}, {Name: "Mountain", Count: 56}},
	}
}

func list(source, event, date string, n, place int) meta.List {
	return meta.List{
		Source: source, ID: fmt.Sprintf("%s-%s-%d", event, date, n), Event: event, Date: date, Placement: place,
		Cards: []meta.Card{{Name: "Lightning Bolt", Count: 4}, {Name: "Mountain", Count: 56}},
	}
}

// fixture holds 30 lists of each precon rung, a Starter deck that labels
// no rung, a 40-card Intro pack, and tournament lists in many events.
func fixture() ([]meta.Precon, []meta.List) {
	var ps []meta.Precon
	for i := range 30 {
		date := fmt.Sprintf("2020-01-%02d", i+1)
		ps = append(ps, precon("Theme Deck", i, date), precon("Challenger Deck", i, date))
	}
	ps = append(ps, precon("Starter Deck", 0, "2030-01-01"))
	small := precon("Intro Pack", 99, "2030-01-01")
	small.Cards = small.Cards[:1]
	ps = append(ps, small)
	var ls []meta.List
	for day := 1; day <= 20; day++ {
		date := fmt.Sprintf("2026-09-%02d", day)
		for n := range 4 {
			ls = append(ls, list(meta.SourceMTGO, "Modern Challenge 32", date, n, n+1))
			ls = append(ls, list(meta.SourceMTGTop8, "RCQ @ A Store", date, n, n+1))
		}
		// A league list and a ninth place label no rung.
		ls = append(ls, list(meta.SourceMTGO, "Modern League", date, 9, 1))
		ls = append(ls, list(meta.SourceMTGO, "Modern Challenge 64", date, 9, 9))
	}
	return ps, ls
}

// TestPickBuildsTheRungs is D-863, D-867, and D-871: three rungs of 25,
// every fifth list in the dev split, the newest lists first, and the
// tournament rung split by source with two lists of one event at most.
func TestPickBuildsTheRungs(t *testing.T) {
	ps, ls := fixture()
	set, err := pick(testIndex(), ps, ls, nil)
	if err != nil {
		t.Fatal(err)
	}
	count := map[mtgv1.SixtyStep]int{}
	dev := map[mtgv1.SixtyStep]int{}
	sources := map[string]int{}
	events := map[string]int{}
	for _, l := range set.lists {
		count[l.label]++
		if l.dev {
			dev[l.label]++
		}
		if l.label == mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT {
			sources[strings.Fields(l.id)[0]]++
			events[l.event]++
			if strings.Contains(l.name, "League") || strings.Contains(l.name, "64") {
				t.Errorf("the rung took %s", l.id)
			}
		}
		if strings.Contains(l.name, "Starter") || strings.Contains(l.name, "99") {
			t.Errorf("the rung took %s", l.id)
		}
	}
	for _, s := range steps {
		if count[s] != 25 || dev[s] != 5 {
			t.Errorf("%s: lists = %d, dev = %d, want 25 and 5", stepWord(s), count[s], dev[s])
		}
	}
	if sources[meta.SourceMTGO] != 13 || sources[meta.SourceMTGTop8] != 12 {
		t.Errorf("sources = %v, want 13 MTGO and 12 MTGTop8", sources)
	}
	for e, n := range events {
		if n > 2 {
			t.Errorf("event %s gave %d lists, want 2 at most", e, n)
		}
	}
	if first := set.lists[0]; first.date != "2020-01-30" || first.format != "Standard" {
		t.Errorf("first casual list = %s %s, want the newest", first.date, first.format)
	}
}

// TestPickSkipsAListThatDoesNotResolve reads a newer list with an unknown
// card as a skip, and the rung takes the next list in its place.
func TestPickSkipsAListThatDoesNotResolve(t *testing.T) {
	ps, ls := fixture()
	odd := precon("Theme Deck", 50, "2021-01-01")
	odd.Cards = append(odd.Cards, meta.PreconCard{Name: "No Such Card", Count: 1})
	ps = append(ps, odd)
	set, err := pick(testIndex(), ps, ls, nil)
	if err != nil {
		t.Fatal(err)
	}
	if set.skipped[mtgv1.SixtyStep_SIXTY_STEP_CASUAL] != 1 {
		t.Errorf("skipped = %v, want one casual list", set.skipped)
	}
}

// TestPickRefusesAThinRung stops before any provider call when a rung
// holds fewer lists than the split wants.
func TestPickRefusesAThinRung(t *testing.T) {
	ps, ls := fixture()
	if _, err := pick(testIndex(), ps[:10], ls, nil); err == nil {
		t.Error("a rung of 5 lists passed")
	}
}

// TestPickReadsThePioneerFormat is D-869: a Pioneer Challenger deck
// reads the format of its release.
func TestPickReadsThePioneerFormat(t *testing.T) {
	ps, ls := fixture()
	ps = append(ps, precon("Pioneer Challenger Deck", 0, "2029-01-01"))
	set, err := pick(testIndex(), ps, ls, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range set.lists {
		if l.label == mtgv1.SixtyStep_SIXTY_STEP_FNM {
			if l.format != "Pioneer" {
				t.Errorf("newest FNM list = %s, format %s", l.id, l.format)
			}
			return
		}
	}
	t.Error("no FNM list")
}

// TestVerdict is D-868: 80 percent of the reads name the label, and no
// read sits two steps off. A judge error counts against the rate.
func TestVerdict(t *testing.T) {
	casual, fnm, top := mtgv1.SixtyStep_SIXTY_STEP_CASUAL, mtgv1.SixtyStep_SIXTY_STEP_FNM, mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT
	none := mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED
	res := func(label mtgv1.SixtyStep, reads ...mtgv1.SixtyStep) result {
		return result{list: calList{label: label}, steps: reads}
	}
	for _, c := range []struct {
		name string
		in   []result
		pass bool
	}{
		{"all exact", []result{res(casual, casual, casual), res(top, top, top)}, true},
		{"four of five", []result{res(fnm, fnm, fnm, fnm, fnm, casual)}, true},
		{"three of five", []result{res(fnm, fnm, fnm, fnm, casual, top)}, false},
		{"two steps off", []result{res(casual, casual, casual, casual, casual, casual, casual, casual, casual, casual, top)}, false},
		{"errors count", []result{res(fnm, fnm, fnm, fnm, none, none)}, false},
		{"no reads", nil, false},
	} {
		if got := tally(c.in).pass(); got != c.pass {
			t.Errorf("%s: pass = %v, want %v", c.name, got, c.pass)
		}
	}
	tt := tally([]result{res(casual, casual, fnm, top, none)})
	if r := tt.rungs[casual]; r.exact != 1 || r.one != 1 || r.two != 1 || r.errs != 1 || r.reads != 4 {
		t.Errorf("casual tally = %+v", *r)
	}
}

func TestCell(t *testing.T) {
	if got := cell("a | b\nc"); got != "a / b c" {
		t.Errorf("cell = %q", got)
	}
}

func TestReasonNamesTheError(t *testing.T) {
	if got := reason(result{err: "judge step: timeout"}); got != "error: judge step: timeout" {
		t.Errorf("reason = %q", got)
	}
	if got := reason(result{why: "w", err: "e"}); got != "w" {
		t.Errorf("reason = %q", got)
	}
}
