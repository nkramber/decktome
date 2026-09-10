package questions

import (
	"io"
	"log/slog"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
)

// The F-104 tests. Session vDzEKDPnRZntlyhz8Lqg picked bracket 5 on the
// turn the pick row fired, and the offer read the three most popular
// legends the reader owns. The hints took the bracket of the turn
// before, so the bracket 5 signal of OQ-48 never ordered the offer.

// signalIndex holds three legends. Popularity reads them in name order,
// and the signal reads them the other way round.
func signalIndex() (*cards.Index, func(ids ...string) float64) {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	card := func(name, oracle string, rank int32) *mtgv1.Card {
		return &mtgv1.Card{
			Name: name, OracleId: oracle, EdhrecRank: rank,
			TypeLine: "Legendary Creature — Human", CanBeCommander: true, Legalities: legal,
		}
	}
	idx := cards.NewIndex([]*mtgv1.Card{
		card("Popular Legend", "o-popular", 1),
		card("Middle Legend", "o-middle", 2),
		card("Strong Legend", "o-strong", 3),
	}, nil, nil, time.Time{})
	signal := map[string]float64{"o-strong": 0.9, "o-middle": 0.5, "o-popular": 0.1}
	return idx, func(ids ...string) float64 { return signal[ids[0]] }
}

func signalHints(t *testing.T) *CandidateHints {
	t.Helper()
	b, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	idx, signal := signalIndex()
	return &CandidateHints{Index: idx, Builder: b, CommanderSignal: signal}
}

// TestTheOfferReadsTheBracketPickedOnItsOwnTurn drives the turn as the
// session did: the power row is out, the reader picks bracket 5, and the
// agent reads the facts before it plans. The offer must read the signal
// order, and not the popularity order of a bracket the turn never saw.
func TestTheOfferReadsTheBracketPickedOnItsOwnTurn(t *testing.T) {
	cat, err := Load()
	if err != nil {
		t.Fatalf("catalog: %v", err)
	}
	h := signalHints(t)
	a := &Agent{cat: cat, log: slog.New(slog.NewTextHandler(io.Discard, nil)), hints: h}
	st := NewState(false)
	asked(st, "q1-power_commander", "power_commander", "power", "power")
	st.OptionAnswers = []OptionAnswer{{QuestionID: "q1-power_commander", Index: 4}}
	a.applyOptionAnswers(st)
	if got := st.Slots.GetPower().GetBracket(); got != 5 {
		t.Fatalf("bracket = %d after the pick, want 5: the fifth option is cEDH", got)
	}
	a.readFacts(st)
	got := h.Commanders("", nil)
	want := []string{"Strong Legend", "Middle Legend", "Popular Legend"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("the offer of the turn that picked bracket 5 reads %v, want the signal order %v", got, want)
	}
}

// TestUseBracketReordersTheOffer is the setter alone. An offer under no
// bracket reads popularity, and the same hints under bracket 5 read the
// signal. The cache keys the two apart, as it keys the sets (D-437).
func TestUseBracketReordersTheOffer(t *testing.T) {
	h := signalHints(t)
	if got := h.Commanders("", nil); got[0] != "Popular Legend" {
		t.Fatalf("with no bracket the offer reads %v, want popularity first", got)
	}
	h.UseBracket(5)
	if got := h.Commanders("", nil); got[0] != "Strong Legend" {
		t.Errorf("under bracket 5 the offer reads %v, want the signal first", got)
	}
	// Zero is no bracket, so the one on the hints stands.
	h.UseBracket(0)
	if h.Bracket != 5 {
		t.Errorf("UseBracket(0) moved the bracket to %d", h.Bracket)
	}
	// A bracket under 4 reads popularity again.
	h.UseBracket(3)
	if got := h.Commanders("", nil); got[0] != "Popular Legend" {
		t.Errorf("under bracket 3 the offer reads %v, want popularity first", got)
	}
}
