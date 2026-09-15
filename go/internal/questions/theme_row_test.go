package questions

import (
	"context"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
)

// themeRowAsked is a Commander session whose theme row is out for the
// theme, as turn 1 leaves it.
func themeRowAsked(theme string) *State {
	st := NewState(true)
	st.Slots.Format = &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}
	st.Slots.Theme = theme
	st.Close("format")
	st.Close("theme")
	st.Turn, st.AskCount = 1, 1
	st.Asks = []Ask{{QuestionID: "q1-theme_unmatched", RowID: "theme_unmatched", Slot: "theme", Key: SlotThemeUnmatched, Turn: 1}}
	st.MarkAsked("theme_unmatched", SlotThemeUnmatched, "theme")
	st.RecordAskedTheme()
	return st
}

// TestThemeRowAsksBeforeTheBuild is F-143. The reader wrote "opponent
// milling cards", no word matched a card, and the build ran with no word
// to the reader. The row now asks for the theme again, and the session is
// not ready (D-725).
func TestThemeRowAsksBeforeTheBuild(t *testing.T) {
	out := classifyOut{Format: "commander", Theme: "opponent milling cards", PoolRule: "unknown"}
	a, _ := testAgentHints(t, &fakeHints{unmatched: true}, classifyStep(t, out),
		fits(t, "theme_unmatched", "power_commander", "colors"), askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "Build a deck focused around opponent milling cards.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if res.Ready {
		t.Error("a theme that matches no card reached the build")
	}
	asked := false
	for _, q := range res.Questions {
		asked = asked || strings.HasPrefix(q.GetText(), "No card I know matches the theme opponent milling cards.")
	}
	if !asked {
		t.Errorf("the turn did not ask the theme row: %v", res.Questions)
	}
	if st.UnmatchedThemeAsked != "opponent milling cards" {
		t.Errorf("the row recorded %q as the theme it named", st.UnmatchedThemeAsked)
	}
}

// TestThemeRowClosesOnAThemeThatMatches is D-725. The reader answers with
// a theme that matches cards, and the key closes.
func TestThemeRowClosesOnAThemeThatMatches(t *testing.T) {
	a, _ := testAgentHints(t, &fakeHints{})
	st := themeRowAsked("opponent milling cards")
	st.Slots.Theme = "mill"
	a.readThemeMatch(st)
	if got := st.Slots.GetSlotStates()[SlotThemeUnmatched]; got != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("slot state = %v, want filled", got)
	}
	if st.Ctx.ThemeUnmatched {
		t.Error("a theme that matches cards still reads as unmatched")
	}
}

// TestThemeRowSkipsAKeptTheme is D-725 beside D-599. The reader replied to
// the row and kept the theme, so the build reads its words as they are.
// Before this, the re-ask of D-599 sent the same question again. With no
// reply the question stays out, for the net of D-351.
func TestThemeRowSkipsAKeptTheme(t *testing.T) {
	a, _ := testAgentHints(t, &fakeHints{unmatched: true})
	st := themeRowAsked("opponent milling cards")
	st.AnsweredQuestions = []string{"q1-theme_unmatched"}
	a.readThemeMatch(st)
	if got := st.Slots.GetSlotStates()[SlotThemeUnmatched]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("a reply that keeps the theme left slot state %v, want skipped", got)
	}
	st = themeRowAsked("opponent milling cards")
	a.readThemeMatch(st)
	if got := st.Slots.GetSlotStates()[SlotThemeUnmatched]; got != mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Errorf("no reply left slot state %v, want asked", got)
	}
}

// TestThemeRowAsksAgainOnlyForAnotherMiss is the D-210 rule of the row.
// The same theme gets no second question, and another theme that matches
// no card gets one.
func TestThemeRowAsksAgainOnlyForAnotherMiss(t *testing.T) {
	a, _ := testAgentHints(t, &fakeHints{unmatched: true})
	c := load(t)
	st := themeRowAsked("opponent milling cards")
	st.Skip(SlotThemeUnmatched)
	a.readThemeMatch(st)
	if contains(ids(c.Plan(st.Ctx)), "theme_unmatched") {
		t.Error("the row asked again for the same theme")
	}
	st.Slots.Theme = "zzzz"
	a.readThemeMatch(st)
	if got := ids(c.Plan(st.Ctx)); !contains(got, "theme_unmatched") {
		t.Errorf("another theme that matches no card got no question: %v", got)
	}
}

// TestPoolRowsStayOutOfTheThemeRowTurn: the card pool rows read the theme,
// so none of them shares a turn with the row that asks for it again. The
// thin-theme row would name a count of zero cards beside it.
func TestPoolRowsStayOutOfTheThemeRowTurn(t *testing.T) {
	c := load(t)
	base := ctx(mtgv1.FormatId_FORMAT_ID_COMMANDER, "format", "theme", "colors", "power", "commander")
	base.HasCollection = true
	if got := ids(c.Plan(base)); !contains(got, "pool") {
		t.Fatalf("the pool row must ask in this context: %v", got)
	}
	base.ThemeUnmatched = true
	got := ids(c.Plan(base))
	if !contains(got, "theme_unmatched") {
		t.Errorf("the theme row did not ask: %v", got)
	}
	for _, id := range []string{"pool", "pool_thin", "pool_precon"} {
		if contains(got, id) {
			t.Errorf("row %s shares a turn with the theme row: %v", id, got)
		}
	}
}

// TestThemeUnmatchedReadsTheCardDatabase is F-142 through the real hint
// source. "opponent milling cards" matched no card before D-724, and a
// word no card holds still reads as unmatched.
func TestThemeUnmatchedReadsTheCardDatabase(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{{
		OracleId: "millstone", Name: "Millstone", TypeLine: "Artifact", CardTypes: []string{"Artifact"},
		OracleText: "{2}, {T}: Target player mills two cards.", ManaValue: 2,
		Legalities: map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}}, nil, nil, time.Date(2026, 9, 14, 0, 0, 0, 0, time.UTC))
	b, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	h := &CandidateHints{Index: idx, Builder: b, Format: mtgv1.FormatId_FORMAT_ID_COMMANDER}
	if h.ThemeUnmatched("opponent milling cards") {
		t.Error("milling finds the mill row, and Millstone mills")
	}
	if !h.ThemeUnmatched("zzzz") {
		t.Error("a word no card holds must read as unmatched")
	}
	if h.ThemeUnmatched("") {
		t.Error("an empty theme is no miss")
	}
	var none *CandidateHints
	if none.ThemeUnmatched("zzzz") {
		t.Error("a nil hint source can not tell, so it answers false")
	}
}
