package generate

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
)

// themeMarkList holds one card of each input of the D-1190 class, and a
// staple. Each themed card takes a staple role, so its job word alone
// hides the theme.
func themeMarkList() *candidates.List {
	card := func(id, name, typeLine string) *mtgv1.Card {
		return &mtgv1.Card{OracleId: id, Name: name, TypeLine: typeLine}
	}
	return &candidates.List{
		Candidates: []candidates.Candidate{
			{Card: card("o-pyro", "Way of the Pyromancer", "Legendary Enchantment"), Role: mtgv1.CardRole_CARD_ROLE_RAMP, Themed: true, OnTheme: true},
			{Card: card("o-well", "Well of Lost Dreams", "Artifact"), Role: mtgv1.CardRole_CARD_ROLE_DRAW, Themed: true, OnTheme: true},
			{Card: card("o-skin", "Skinrender", "Creature — Phyrexian Zombie"), Role: mtgv1.CardRole_CARD_ROLE_REMOVAL, Themed: true, OnTheme: true},
			{Card: card("o-sol", "Sol Ring", "Artifact"), Role: mtgv1.CardRole_CARD_ROLE_RAMP},
		},
		Upgrades: []candidates.Candidate{
			{Card: card("o-garruk", "Garruk, Veiled Butcher", "Legendary Planeswalker — Garruk"), Role: mtgv1.CardRole_CARD_ROLE_REMOVAL, Themed: true, OnTheme: true},
		},
	}
}

// TestShortlistMarksEachThemedCard is D-1190. Live eval dDD9Iav9aGhgGF9dDLNQ
// read "Way of the Pyromancer | Legendary Enchantment | ramp" beside
// "Izzet Signet | Artifact | ramp", and the model played the signet it
// knew. Each themed line now carries the mark, the upgrades included, and
// a staple line carries none.
func TestShortlistMarksEachThemedCard(t *testing.T) {
	b, _, _ := testBuilder(t)
	l := themeMarkList()
	req := testRequest()
	req.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	req.Pool = FromList(l, nil, true)
	req.Roles = Roles(l)
	req.Themed = Themed(l)
	got := b.shortlist(req)
	for _, want := range []string{
		"- Way of the Pyromancer | Legendary Enchantment | ramp | on theme\n",
		"- Well of Lost Dreams | Artifact | draw | on theme\n",
		"- Skinrender | Creature — Phyrexian Zombie | removal | on theme\n",
		"- Garruk, Veiled Butcher | Legendary Planeswalker — Garruk | on theme\n",
		"- Sol Ring | Artifact | ramp\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("the shortlist holds no line %q:\n%s", want, got)
		}
	}
	if n := strings.Count(got, "| on theme"); n != 4 {
		t.Errorf("the shortlist marks %d cards, want 4:\n%s", n, got)
	}
	req.Themed = nil
	if got := b.shortlist(req); strings.Contains(got, "| on theme") {
		t.Errorf("a request with no themed card marks a line:\n%s", got)
	}
}

// TestBuildPromptReadsTheThemeMark is D-1190 end to end: the model input
// of a build holds the mark, and the stable prefix tells the model what
// the mark means.
func TestBuildPromptReadsTheThemeMark(t *testing.T) {
	b, _, _ := testBuilder(t)
	l := themeMarkList()
	req := testRequest()
	req.Pool = FromList(l, nil, true)
	req.Roles = Roles(l)
	req.Themed = Themed(l)
	if in := b.input(req, nil, nil); !strings.Contains(in, "- Well of Lost Dreams | Artifact | draw | on theme") {
		t.Errorf("the model input holds no theme mark:\n%s", in)
	}
	if !strings.Contains(generateInstructions, `says "on theme"`) {
		t.Error("the generate prompt does not say what the theme mark means")
	}
	// The replay of the first wording wrote "theme-marked" in 13 reasons
	// (D-1191). The prompt now forbids a mark in user text.
	if !strings.Contains(generateInstructions, `Never write the words "mark" or "marked"`) {
		t.Error("the generate prompt lets a reason name a mark")
	}
}

// TestThemeMarkSkipsACommanderRateCard is D-1195. At bracket 5 a card that
// the lists of the commander play reads Themed with no theme signal
// (D-839). Its line must not say "on theme", because the prompt tells the
// model that such a card matches the theme.
func TestThemeMarkSkipsACommanderRateCard(t *testing.T) {
	b, _, _ := testBuilder(t)
	l := themeMarkList()
	l.Candidates = append(l.Candidates, candidates.Candidate{
		Card: &mtgv1.Card{OracleId: "o-rhystic", Name: "Rhystic Study", TypeLine: "Enchantment"},
		Role: mtgv1.CardRole_CARD_ROLE_DRAW, Themed: true,
	})
	if Themed(l)["o-rhystic"] {
		t.Error("Themed reads a card that only the commander rate leads")
	}
	req := testRequest()
	req.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	req.Pool = FromList(l, nil, true)
	req.Roles = Roles(l)
	req.Themed = Themed(l)
	got := b.shortlist(req)
	if !strings.Contains(got, "- Rhystic Study | Enchantment | draw\n") {
		t.Errorf("the commander rate card reads a theme mark:\n%s", got)
	}
}

// TestThemedLineCarriesTheCardText is D-1196. Three replays of the mark
// alone left out Garruk, Veiled Butcher, because its line held no text and
// the card is newer than the model. A line on theme now ends with the cost
// and the rules text. A staple line stays short, so the shortlist grows by
// the cards on theme alone.
func TestThemedLineCarriesTheCardText(t *testing.T) {
	b, _, _ := testBuilder(t)
	l := themeMarkList()
	l.Upgrades[0].Card.ManaCost = "{3}{B}{B}"
	l.Upgrades[0].Card.OracleText = "If a creature an opponent controls would die, exile it instead.\n+2: Up to one target creature gets -4/-1 until your next turn."
	l.Candidates[3].Card.ManaCost = "{1}"
	l.Candidates[3].Card.OracleText = "{T}: Add {C}{C}."
	req := testRequest()
	req.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	req.Pool = FromList(l, nil, true)
	req.Roles = Roles(l)
	req.Themed = Themed(l)
	got := b.shortlist(req)
	for _, want := range []string{
		"- Garruk, Veiled Butcher | Legendary Planeswalker — Garruk | on theme | {3}{B}{B} | If a creature an opponent controls would die, exile it instead. +2: Up to one target creature gets -4/-1 until your next turn.\n",
		"- Sol Ring | Artifact | ramp\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("the shortlist holds no line %q:\n%s", want, got)
		}
	}
	if !strings.Contains(generateInstructions, "Judge the card by that text") {
		t.Error("the generate prompt does not say what the text of a line on theme is for")
	}
}
