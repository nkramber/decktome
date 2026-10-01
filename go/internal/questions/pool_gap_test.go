package questions

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
)

// gapIndex holds a white legend, a red card, a Plains, and n white
// commons. The owned map holds the legend when ownLegend is set, the red
// card, and the first owned commons.
func gapHints(t *testing.T, commons, owned int, ownLegend bool) *CandidateHints {
	t.Helper()
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	white := []mtgv1.Color{mtgv1.Color_COLOR_W}
	list := []*mtgv1.Card{
		{Name: "Karlov of the Ghost Council", OracleId: "o-karlov", ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B},
			TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true, Legalities: legal},
		{Name: "Heliod, Sun-Crowned", OracleId: "o-heliod", ColorIdentity: white,
			TypeLine: "Legendary Enchantment Creature — God", CanBeCommander: true, Legalities: legal},
		{Name: "Lightning Bolt", OracleId: "o-bolt", ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_R},
			TypeLine: "Instant", Legalities: legal},
		{Name: "Plains", OracleId: "o-plains", ColorIdentity: nil, TypeLine: "Basic Land — Plains",
			Supertypes: []string{"Basic"}, CardTypes: []string{"Land"}, Legalities: legal},
	}
	have := map[string]int32{"o-bolt": 1, "o-heliod": 1}
	if ownLegend {
		have["o-karlov"] = 1
	}
	for i := range commons {
		id := fmt.Sprintf("o-common-%d", i)
		list = append(list, &mtgv1.Card{Name: fmt.Sprintf("White Common %d", i), OracleId: id,
			ColorIdentity: white, TypeLine: "Creature — Soldier", Legalities: legal})
		if i < owned {
			have[id] = 1
		}
	}
	b, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	return &CandidateHints{
		Index:   cards.NewIndex(list, nil, nil, time.Time{}),
		Builder: b,
		Format:  mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Owned:   have,
	}
}

// TestPoolGapChecks is D-1027: three checks, in order, and the first that
// holds names the gap.
func TestPoolGapChecks(t *testing.T) {
	t.Run("a named commander the collection lacks", func(t *testing.T) {
		h := gapHints(t, 80, 80, false)
		g := h.PoolGap("", []string{"Plains", "Karlov of the Ghost Council"}, []string{"Karlov of the Ghost Council"})
		if g.Kind != GapUnowned || !slices.Equal(g.Names, []string{"Karlov of the Ghost Council"}) {
			t.Errorf("gap = %+v, want Karlov alone: a basic land is never unowned", g)
		}
		if want := "Your collection holds no copy of Karlov of the Ghost Council, so the buy list names it in each case."; g.Sentence() != want {
			t.Errorf("sentence = %q, want %q", g.Sentence(), want)
		}
	})

	t.Run("a theme under the floor", func(t *testing.T) {
		h := gapHints(t, 80, 80, true)
		g := h.PoolGap("elves", nil, []string{"Heliod, Sun-Crowned"})
		if g.Kind != GapTheme || g.Have != 0 || g.Want != 30 {
			t.Errorf("gap = %+v, want 0 elves cards of a floor of 30", g)
		}
	})

	// The commander sets the colors when the colors slot is empty. The
	// red card is owned, and it does not count for a white deck.
	t.Run("too few owned cards in the commander's colors", func(t *testing.T) {
		h := gapHints(t, 80, 68, true)
		g := h.PoolGap("", nil, []string{"Heliod, Sun-Crowned"})
		if g.Kind != GapColors || g.Have != 69 || g.Want != 70 {
			t.Errorf("gap = %+v, want 69 of 70: 68 commons and Heliod, and no red card", g)
		}
	})

	t.Run("a collection that meets the request", func(t *testing.T) {
		h := gapHints(t, 80, 80, true)
		if g := h.PoolGap("", nil, []string{"Heliod, Sun-Crowned"}); g.Kind != "" {
			t.Errorf("gap = %+v, want none", g)
		}
	})

	t.Run("a 60-card format halves the floors", func(t *testing.T) {
		h := gapHints(t, 80, 30, true)
		h.Format = mtgv1.FormatId_FORMAT_ID_MODERN
		h.Colors = []mtgv1.Color{mtgv1.Color_COLOR_W}
		if g := h.PoolGap("", nil, nil); g.Kind != GapColors || g.Want != 35 {
			t.Errorf("gap = %+v, want the floor of 35", g)
		}
		if got := candidates.ThemeFloor(mtgv1.FormatId_FORMAT_ID_MODERN); got != 15 {
			t.Errorf("theme floor = %d, want 15 (D-1032)", got)
		}
	})
}

// gapSource answers one gap for every call.
type gapSource struct{ gap Gap }

func (g gapSource) ThinTheme(string) (bool, int)           { return false, 0 }
func (g gapSource) PoolGap(string, []string, []string) Gap { return g.gap }

// TestTheGapQuestionAsksOnTheReaderChoice is D-1011 and D-1027. The
// reader's choice fills the pool key on the first turn, so the gap row
// asks on its own key. An answer writes the pool rule.
func TestTheGapQuestionAsksOnTheReaderChoice(t *testing.T) {
	gap := Gap{Kind: GapTheme, Theme: "elves", Have: 12, Want: 30}
	chosen := func() *State {
		st := NewState(true)
		st.Ctx.HasCollection, st.Ctx.PoolFromReader = true, true
		st.Slots.PoolRule = mtgv1.PoolRule_POOL_RULE_OWNED_ONLY
		st.Close("pool_rule")
		st.Slots.Format = &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}
		st.Slots.Theme = "elves"
		st.Close("format")
		st.Close("theme")
		return st
	}

	t.Run("the facts read the gap and the sentence", func(t *testing.T) {
		st := chosen()
		RefreshFacts(st, gapSource{gap})
		if !st.Ctx.PoolGap || st.Ctx.PoolGapReason != "Your collection holds 12 elves cards, and I want 30 or more." {
			t.Errorf("facts = %v %q", st.Ctx.PoolGap, st.Ctx.PoolGapReason)
		}
	})

	t.Run("no gap question without the reader's choice", func(t *testing.T) {
		st := chosen()
		st.Ctx.PoolFromReader = false
		RefreshFacts(st, gapSource{gap})
		if st.Ctx.PoolGap {
			t.Error("the gap fact holds for a pool the reader did not choose")
		}
	})

	t.Run("an answered gap question stops the checks", func(t *testing.T) {
		st := chosen()
		st.Slots.PoolRule = mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
		st.Close("pool_gap")
		RefreshFacts(st, gapSource{gap})
		if st.Ctx.PoolGap {
			t.Error("the gap fact holds after the answer")
		}
	})

	t.Run("the row plans on the filled pool key", func(t *testing.T) {
		cat, err := Load()
		if err != nil {
			t.Fatal(err)
		}
		st := chosen()
		st.Close("colors")
		RefreshFacts(st, gapSource{gap})
		var ids []string
		for _, r := range cat.Plan(st.Ctx) {
			ids = append(ids, r.ID)
		}
		if !slices.Contains(ids, "pool_gap") {
			t.Errorf("plan = %v, want pool_gap", ids)
		}
		row, _ := cat.Row("pool_gap")
		text, opts, _ := resolve(row, st, nil)
		if text != "Your collection holds 12 elves cards, and I want 30 or more. Should I fill the gaps from any card, with a buy list, or use only your cards?" {
			t.Errorf("text = %q", text)
		}
		if !slices.Equal(opts, []string{"Fill the gaps from any card (buy list)", "Only my cards"}) {
			t.Errorf("options = %v", opts)
		}
	})

	t.Run("a typed answer writes the rule, and any card reads as a fill", func(t *testing.T) {
		for word, want := range map[string]mtgv1.PoolRule{
			"owned_first": mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
			"any_card":    mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
			"owned_only":  mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		} {
			a := &Agent{log: slog.New(slog.NewTextHandler(io.Discard, nil))}
			st := chosen()
			st.MarkAsked("pool_gap", "pool_gap", "pool_rule")
			a.apply(context.Background(), st, classifyOut{PoolRule: word}, nil, "a message", nil)
			if got := st.Slots.GetPoolRule(); got != want || !st.Ctx.Filled["pool_gap"] {
				t.Errorf("%q: rule = %v, filled = %v, want %v and filled", word, got, st.Ctx.Filled["pool_gap"], want)
			}
		}
	})
}
