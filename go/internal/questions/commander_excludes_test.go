package questions

import (
	"context"
	"io"
	"log/slog"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/llm"
)

// exclusionIndex holds the cards of the deployed session of D-1217: a
// black and red commander, a white card the reader asked for, and a
// colorless one.
func exclusionIndex() *cards.Index {
	creature := "Legendary Creature — Dragon"
	return cards.NewIndex([]*mtgv1.Card{
		{Name: "Smaug the Impenetrable", TypeLine: creature, CanBeCommander: true,
			ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_R}},
		{Name: "Dwalin, Weaponmaster", TypeLine: "Legendary Creature — Dwarf Warrior", CanBeCommander: true,
			ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W}},
		{Name: "Sauron, the Necromancer", TypeLine: "Legendary Creature — Avatar Horror", CanBeCommander: true,
			ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_B}},
		{Name: "The Arkenstone // Seek the Heart", TypeLine: "Legendary Artifact // Sorcery — Adventure",
			ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W}},
		{Name: "The One Ring", TypeLine: "Legendary Artifact"},
	}, nil, nil, time.Time{})
}

// exclusionTurn is one turn of the exclusion tests. option, when set,
// picks that option of the exclusion row, as the client does.
type exclusionTurn struct {
	message  string
	classify classifyOut
	option   *int
}

// exclusionPlay is played with the text of each question that went out.
type exclusionPlay struct {
	played
	texts []string
}

// playExclusion runs the turns through the agent and a store cycle, as
// play does, and it sends an option answer when a turn names one. The
// reader of the deployed case holds a collection and builds from it.
func playExclusion(t *testing.T, turns []exclusionTurn) exclusionPlay {
	t.Helper()
	fake := &roleFake{t: t}
	for _, tr := range turns {
		fake.classify = append(fake.classify, classifyStep(t, tr.classify).Output)
	}
	c, err := llm.New(fakeConfig(), []llm.Provider{fake}, llm.WithoutJitter())
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	a, err := NewAgent(load(t), c, WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
		WithHints(&CandidateHints{Index: exclusionIndex()}))
	if err != nil {
		t.Fatalf("agent: %v", err)
	}
	st := NewState(true)
	st.SessionID = "exclusion"
	out := exclusionPlay{played: played{fake: fake}}
	for i, tr := range turns {
		st.OptionAnswers = nil
		if tr.option != nil {
			id := ""
			for _, rec := range st.Asks {
				if rec.RowID == SlotCommanderExcludes {
					id = rec.QuestionID
				}
			}
			if id == "" {
				t.Fatalf("turn %d picks an option of a row that never asked", i+1)
			}
			st.OptionAnswers = []OptionAnswer{{QuestionID: id, Index: *tr.option}}
		}
		res, err := a.Turn(context.Background(), st, tr.message, nil)
		if err != nil {
			t.Fatalf("turn %d: %v", i+1, err)
		}
		var ids []string
		for _, rec := range st.Asks {
			if rec.Turn == st.Turn {
				ids = append(ids, rec.RowID)
			}
		}
		for _, q := range res.Questions {
			out.texts = append(out.texts, q.GetText())
		}
		out.rows = append(out.rows, ids)
		out.ready = append(out.ready, res.Ready)
		st = roundTrip(t, st)
	}
	out.st = st
	return out
}

// smaugTurn is the deployed request: a Commander deck that keeps The
// Arkenstone and The One Ring, led by Smaug.
func smaugTurn() exclusionTurn {
	return exclusionTurn{
		message: "A Commander deck with The One Ring and The Arkenstone. Use Smaug the Impenetrable.",
		classify: classifyOut{Format: "commander", Theme: "treasure", PoolRule: "owned_only", Power: "bracket 3",
			CommanderNames: []string{"Smaug the Impenetrable"},
			LockedNames:    []string{"The One Ring", "The Arkenstone // Seek the Heart"}},
	}
}

func option(i int) *int { return &i }

// TestACommanderThatLeavesOutALockedCardAsks is D-1217 on the deployed
// case. The build kept The Arkenstone under Smaug, and the engine blocked
// the deck. The row must name the commander, its colors, and the card,
// and leave the colorless card alone.
func TestACommanderThatLeavesOutALockedCardAsks(t *testing.T) {
	p := playExclusion(t, []exclusionTurn{smaugTurn()})
	if !p.askedOn(1, SlotCommanderExcludes) {
		t.Fatalf("the commander left out a locked card in silence: %v", p.rows)
	}
	if p.ready[0] {
		t.Error("the session called itself ready with the exclusion row out")
	}
	if got := p.st.ExcludedCards; len(got) != 1 || !sameCard(got[0], "The Arkenstone // Seek the Heart") {
		t.Errorf("ExcludedCards = %v, want The Arkenstone alone", got)
	}
	text := strings.Join(p.texts, "\n")
	t.Logf("questions: %s", text)
	for _, want := range []string{"Smaug the Impenetrable is black and red", "The Arkenstone // Seek the Heart (white)", "pick another commander"} {
		if !strings.Contains(text, want) {
			t.Errorf("the question %q does not hold %q", text, want)
		}
	}
}

// The first option keeps the commander, and the card leaves the deck. The
// colorless card stays, and the row does not ask again.
func TestKeepingTheCommanderLeavesTheCardOut(t *testing.T) {
	p := playExclusion(t, []exclusionTurn{
		smaugTurn(),
		{option: option(0), classify: classifyOut{Format: "unknown", PoolRule: "unknown"}},
	})
	if p.askedOn(2, SlotCommanderExcludes) {
		t.Errorf("the row asked again after the reader kept the commander: %v", p.rows)
	}
	locked := p.st.LockedCards()
	if hasName(locked, "The Arkenstone // Seek the Heart") {
		t.Errorf("the build still keeps The Arkenstone: %v", locked)
	}
	if !hasName(locked, "The One Ring") {
		t.Errorf("the colorless card left the deck: %v", locked)
	}
	if !hasName(p.st.CommanderNames, "Smaug the Impenetrable") {
		t.Errorf("the commander is %v, want Smaug", p.st.CommanderNames)
	}
	if !p.ready[1] {
		t.Errorf("the session is not ready after the answer: %v", p.rows)
	}
}

// A reader who types "keep" answers as the first option does.
func TestTypedKeepLeavesTheCardOut(t *testing.T) {
	p := playExclusion(t, []exclusionTurn{
		smaugTurn(),
		{message: "Keep Smaug.", classify: classifyOut{Format: "unknown", PoolRule: "unknown"}},
	})
	if hasName(p.st.LockedCards(), "The Arkenstone // Seek the Heart") {
		t.Errorf("the build still keeps The Arkenstone: %v", p.st.LockedCards())
	}
	if !p.ready[1] {
		t.Errorf("the session is not ready after the answer: %v", p.rows)
	}
}

// The second option reopens the choice. A new commander that holds the
// card closes the row, and the card stays locked.
func TestAnotherCommanderKeepsTheCard(t *testing.T) {
	p := playExclusion(t, []exclusionTurn{
		smaugTurn(),
		{option: option(1), classify: classifyOut{Format: "unknown", PoolRule: "unknown"}},
		{message: "Dwalin, Weaponmaster.", classify: classifyOut{Format: "unknown", PoolRule: "unknown",
			CommanderNames: []string{"Dwalin, Weaponmaster"}}},
	})
	if p.askedOn(2, SlotCommanderExcludes) || p.askedOn(3, SlotCommanderExcludes) {
		t.Errorf("the row asked again: %v", p.rows)
	}
	if !hasName(p.st.CommanderNames, "Dwalin, Weaponmaster") || hasName(p.st.CommanderNames, "Smaug the Impenetrable") {
		t.Errorf("the commander is %v, want Dwalin alone", p.st.CommanderNames)
	}
	if !hasName(p.st.LockedCards(), "The Arkenstone // Seek the Heart") {
		t.Errorf("the card left the deck under a commander that holds it: %v", p.st.LockedCards())
	}
	if !p.ready[2] {
		t.Errorf("the session is not ready: %v", p.rows)
	}
}

// The swap words of D-130 answer the row as the second option does.
func TestTypedSwapReopensTheChoice(t *testing.T) {
	p := playExclusion(t, []exclusionTurn{
		smaugTurn(),
		{message: "Pick another commander.", classify: classifyOut{Format: "unknown", PoolRule: "unknown"}},
	})
	if p.st.Ctx.CommanderSet {
		t.Errorf("the commander stays %v", p.st.CommanderNames)
	}
	if !hasName(p.st.LockedCards(), "The Arkenstone // Seek the Heart") {
		t.Errorf("the card left the deck: %v", p.st.LockedCards())
	}
	if p.st.Slots.GetSlotStates()[SlotCommanderExcludes] == mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Error("the exclusion row is still out")
	}
}

// A second commander that leaves out the card asks again, with its own
// name (D-210).
func TestASecondCommanderThatLeavesOutTheCardAsksAgain(t *testing.T) {
	p := playExclusion(t, []exclusionTurn{
		smaugTurn(),
		{option: option(1), classify: classifyOut{Format: "unknown", PoolRule: "unknown"}},
		{message: "Sauron, the Necromancer.", classify: classifyOut{Format: "unknown", PoolRule: "unknown",
			CommanderNames: []string{"Sauron, the Necromancer"}}},
	})
	if !p.askedOn(3, SlotCommanderExcludes) {
		t.Fatalf("a second commander left out the card in silence: %v", p.rows)
	}
	last := p.texts[len(p.texts)-1]
	if !strings.Contains(last, "Sauron, the Necromancer is black") {
		t.Errorf("the second question %q does not name the new commander", last)
	}
}

// The reader who picks the same commander again has answered for it, so
// the card leaves the deck with no second question.
func TestTheSameCommanderAgainLeavesTheCardOut(t *testing.T) {
	p := playExclusion(t, []exclusionTurn{
		smaugTurn(),
		{option: option(1), classify: classifyOut{Format: "unknown", PoolRule: "unknown"}},
		{message: "Smaug the Impenetrable after all.", classify: classifyOut{Format: "unknown", PoolRule: "unknown",
			CommanderNames: []string{"Smaug the Impenetrable"}}},
	})
	if p.askedOn(3, SlotCommanderExcludes) {
		t.Errorf("the row asked the same question twice: %v", p.rows)
	}
	if hasName(p.st.LockedCards(), "The Arkenstone // Seek the Heart") {
		t.Errorf("the build still keeps The Arkenstone: %v", p.st.LockedCards())
	}
}

// A card the index does not hold is no proof, so no question goes out.
func TestAnUnknownLockedCardAsksNothing(t *testing.T) {
	tr := smaugTurn()
	tr.classify.LockedNames = []string{"Not A Real Card"}
	p := playExclusion(t, []exclusionTurn{tr})
	if p.asked(SlotCommanderExcludes) {
		t.Errorf("the row asked about a card the index does not hold: %v", p.rows)
	}
}

func TestIdentityWords(t *testing.T) {
	cases := []struct {
		in   []mtgv1.Color
		want string
	}{
		{nil, "colorless"},
		{[]mtgv1.Color{mtgv1.Color_COLOR_C}, "colorless"},
		{[]mtgv1.Color{mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_B}, "black and red"},
		{[]mtgv1.Color{mtgv1.Color_COLOR_G, mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U}, "white, blue, and green"},
	}
	for _, c := range cases {
		if got := identityWords(c.in); got != c.want {
			t.Errorf("identityWords(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
