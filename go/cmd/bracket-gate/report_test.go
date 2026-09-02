package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
)

func deck(off bool, content bool, checked bool) *mtgv1.Deck {
	d := &mtgv1.Deck{
		Cards:      []*mtgv1.DeckCard{{Name: "Sol Ring", Count: 1}},
		Validation: &mtgv1.ValidationResult{Passed: true},
		Profile: &mtgv1.DeckProfile{
			Bracket:  3,
			Features: []*mtgv1.ProfileFeature{{Key: profile.KeyLand, Value: 36, Low: 34, High: 38, HasHigh: true, OffBand: off}},
			Goldfish: &mtgv1.Goldfish{Hands: 10000, CommanderTurn: 3.2, ManaTurnFour: 4.6, ShareTwoToFourLands: 0.81},
			Content:  &mtgv1.ContentCheck{Checked: checked, SourceTag: "P"},
		},
	}
	if content {
		d.Validation.Findings = append(d.Validation.Findings, &mtgv1.Finding{Code: profile.CodeMassLandDenial, Severity: mtgv1.Severity_SEVERITY_WARN, Message: "Armageddon"})
	}
	return d
}

func runReport(rs []result) (string, bool) {
	var buf bytes.Buffer
	idx := cards.NewIndex(nil, nil, nil, time.Date(2026, 9, 2, 0, 0, 0, 0, time.UTC))
	pass := report(&buf, rs, llm.NewAccumulator(nil), idx, time.Second)
	return buf.String(), pass
}

func good(id int, bracket int32) result {
	return result{prompt: prompt{ID: id, Bracket: bracket, Commander: "Karlov of the Ghost Council", Theme: "lifegain"},
		deck: deck(false, false, true), judged: &generate.BracketJudgement{Bracket: bracket, Why: "It is what it is."}}
}

func TestReportPassesWhenEveryBarHolds(t *testing.T) {
	rs := []result{good(1, 1), good(2, 2), good(3, 3), good(4, 4), good(5, 5)}
	doc, pass := runReport(rs)
	if !pass || !strings.Contains(doc, "Verdict: PASS.") {
		t.Errorf("want PASS:\n%s", doc)
	}
	for _, want := range []string{"# PR-14A bracket gate", "| Decks in every band | 5 |", "| Judge agreed with the bracket | 5 |", "Judge: bracket 3, agrees.", "Goldfish over 10000 hands", "Content: Spellbook tag P."} {
		if !strings.Contains(doc, want) {
			t.Errorf("document lacks %q", want)
		}
	}
}

func TestReportFailsOnEachBar(t *testing.T) {
	cases := []struct {
		name string
		mod  func(*result)
		want string
	}{
		{"off band", func(r *result) { r.deck = deck(true, false, true) }, "| Decks in every band | 4 |"},
		{"content", func(r *result) { r.deck = deck(false, true, true) }, "| Decks with no content violation | 4 |"},
		{"unchecked", func(r *result) { r.deck = deck(false, false, false) }, "1 decks got no content check"},
		{"judge error", func(r *result) { r.judged, r.judgeErr = nil, errors.New("boom") }, "| Judge errors | 1 |"},
		{"no profile", func(r *result) { r.deck.Profile = nil }, "1 decks carry no profile"},
		{"error", func(r *result) { r.deck, r.err = nil, errors.New("no card named X") }, "1 prompts failed before a deck existed"},
		{"block", func(r *result) {
			r.deck.Validation.Findings = append(r.deck.Validation.Findings, &mtgv1.Finding{Code: "deck_size", Severity: mtgv1.Severity_SEVERITY_BLOCK})
		}, "| Decks with no block finding | 4 |"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rs := []result{good(1, 1), good(2, 2), good(3, 3), good(4, 4), good(5, 5)}
			tc.mod(&rs[2])
			doc, pass := runReport(rs)
			if pass || !strings.Contains(doc, "Verdict: FAIL.") {
				t.Errorf("want FAIL:\n%s", doc)
			}
			if !strings.Contains(doc, tc.want) {
				t.Errorf("document lacks %q:\n%s", tc.want, doc)
			}
		})
	}
}

func TestReportJudgeBarIsEightOfTen(t *testing.T) {
	// Twelve of fifteen agree: 80 percent, the bar holds.
	var rs []result
	for i := 1; i <= 15; i++ {
		r := good(i, 3)
		if i <= 3 {
			r.judged.Bracket = 4
		}
		rs = append(rs, r)
	}
	if doc, pass := runReport(rs); !pass {
		t.Errorf("12 of 15 must pass:\n%s", doc)
	}
	rs[3].judged.Bracket = 4
	if doc, pass := runReport(rs); pass {
		t.Errorf("11 of 15 must fail:\n%s", doc)
	}
	if _, pass := runReport(nil); pass {
		t.Error("an empty run must fail")
	}
}

func TestSelectPrompts(t *testing.T) {
	all := []prompt{{ID: 1}, {ID: 2}, {ID: 3}}
	if got, err := selectPrompts(all, ""); err != nil || len(got) != 3 {
		t.Errorf("all: %v %v", got, err)
	}
	if got, err := selectPrompts(all, "2,3"); err != nil || len(got) != 2 || got[0].ID != 2 {
		t.Errorf("some: %v %v", got, err)
	}
	if _, err := selectPrompts(all, "9"); err == nil {
		t.Error("a list that names no prompt must fail")
	}
}
