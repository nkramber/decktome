package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/harvest"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/triage"
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
	pass := report(&buf, rs, llm.NewAccumulator(nil), idx, time.Second, evalrun.New("bracket", "test"))
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

func TestReadDecksParsesAGateDocument(t *testing.T) {
	doc := `# PR-14A bracket gate

Verdict: FAIL.

## Decks

### 7. Bracket 3, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired true (profile_off_band).

Findings:

- WARN ` + "`land_count`" + `: 39 lands

Cards:

- 17 Plains
- 1 Sol Ring

### 13. Bracket 5, Gishath, Sun's Avatar, combo

Cards:

- 1 Sol Ring

`
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "k", Name: "Karlov of the Ghost Council", CardTypes: []string{"Creature"}, Supertypes: []string{"Legendary"}},
		{OracleId: "g", Name: "Gishath, Sun's Avatar", CardTypes: []string{"Creature"}, Supertypes: []string{"Legendary"}},
		{OracleId: "p", Name: "Plains", CardTypes: []string{"Land"}, Supertypes: []string{"Basic"}},
		{OracleId: "s", Name: "Sol Ring", CardTypes: []string{"Artifact"}},
	}, nil, nil, time.Now())
	rs, err := readDecks(strings.NewReader(doc), idx)
	if err != nil {
		t.Fatal(err)
	}
	if len(rs) != 2 || rs[0].prompt.ID != 7 || rs[0].prompt.Bracket != 3 || rs[0].prompt.Theme != "lifegain" || rs[1].prompt.Bracket != 5 {
		t.Errorf("prompts %+v %+v", rs[0].prompt, rs[1].prompt)
	}
	if d := rs[0].deck; len(d.Cards) != 2 || d.Cards[0].Count != 17 || d.Cards[0].OracleId != "p" || d.CommanderOracleIds[0] != "k" || d.GetPower().GetBracket() != 3 {
		t.Errorf("deck %v", d)
	}
	if len(rs[1].deck.Cards) != 1 || rs[1].prompt.Commander != "Gishath, Sun's Avatar" || rs[1].prompt.Theme != "combo" || rs[1].deck.CommanderOracleIds[0] != "g" {
		t.Errorf("second deck %+v %v", rs[1].prompt, rs[1].deck)
	}
	if _, err := readDecks(strings.NewReader(strings.Replace(doc, "Sol Ring", "Not A Card", 1)), idx); err == nil {
		t.Error("an unknown name must fail")
	}
	if _, err := readDecks(strings.NewReader("# nothing\n"), idx); err == nil {
		t.Error("a document with no deck must fail")
	}
}

func TestReportJudge(t *testing.T) {
	var rs []result
	for i := 1; i <= 5; i++ {
		rs = append(rs, good(i, 3))
	}
	rs[0].judged.Bracket = 4
	var buf bytes.Buffer
	idx := cards.NewIndex(nil, nil, nil, time.Now())
	if !reportJudge(&buf, "run1.md", rs, llm.NewAccumulator(nil), idx, time.Second, evalrun.New("bracket-judge", "test")) {
		t.Errorf("4 of 5 must pass:\n%s", buf.String())
	}
	if !strings.Contains(buf.String(), "| 1 | 3 | 4 | no | Karlov of the Ghost Council |") {
		t.Errorf("table:\n%s", buf.String())
	}
	rs[1].judged, rs[1].judgeErr = nil, errors.New("boom")
	buf.Reset()
	if reportJudge(&buf, "run1.md", rs, llm.NewAccumulator(nil), idx, time.Second, evalrun.New("bracket-judge", "test")) {
		t.Error("a judge error must fail")
	}
}

// TestTheCaseShapeMatchesThePromptFile pins the mirror struct of the
// triage against this file's own struct (PR-28b). The triage writes a
// prompt from a reader's power complaint, and it can not import this
// package, so it holds a mirror of the shape. The decoder refuses an
// unknown field, so the mirror may name no field this struct does not
// read.
func TestTheCaseShapeMatchesThePromptFile(t *testing.T) {
	rec := harvest.Record{
		ID: "f1", Kind: "deck", Verdict: "down", Reasons: []string{"wrong_power"},
		Deck: json.RawMessage(`{"id":"d1","sessionId":"s1",` +
			`"format":{"id":"FORMAT_ID_COMMANDER"},"power":{"bracket":3},` +
			`"commanderOracleIds":["o-karlov"]}`),
		Session: json.RawMessage(`{"id":"s1","slots":{"theme":"lifegain"}}`),
	}
	namer := func(id string) (string, bool) { return "Karlov of the Ghost Council", id == "o-karlov" }
	c, err := triage.CaseOf(triage.RouteOf(rec), 999, namer)
	if err != nil {
		t.Fatal(err)
	}
	if c.Target != "go/cmd/bracket-gate/prompts.json" {
		t.Fatalf("the case joins %q, and this test guards another file", c.Target)
	}
	dec := json.NewDecoder(bytes.NewReader(c.Body))
	dec.DisallowUnknownFields()
	var got prompt
	if err := dec.Decode(&got); err != nil {
		t.Fatalf("the triage wrote a field this gate does not read: %v\n%s", err, c.Body)
	}
	if got.ID != 999 || got.Bracket != 3 || got.Theme != "lifegain" || got.Plan == "" {
		t.Errorf("the case lost a field: %+v", got)
	}
	if got.Commander != "Karlov of the Ghost Council" {
		t.Errorf("commander = %q", got.Commander)
	}
}
