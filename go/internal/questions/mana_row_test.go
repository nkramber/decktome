package questions

import (
	"io"
	"log/slog"
	"slices"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
)

// fakeMana answers the counts of the mana row from fixed tables, keyed
// by the colors it reads, and it knows the color identity of each card
// it names. A color list it does not hold counts every color.
type fakeMana struct {
	identity map[string][]mtgv1.Color
	mana     map[string][2]int
	fixing   map[string][2]int
	read     [][]mtgv1.Color
}

func (f *fakeMana) Commanders(string, []string) []string { return nil }
func (f *fakeMana) OwnedThemeCount(string) int           { return 0 }

func (f *fakeMana) ColorIdentity(name string) ([]mtgv1.Color, bool) {
	c, ok := f.identity[name]
	return c, ok
}

func (f *fakeMana) ThinSetMana(_ []string, _ mtgv1.FormatId, colors []mtgv1.Color, _ *mtgv1.PowerLevel) (bool, int, int) {
	f.read = append(f.read, colors)
	n := f.mana[colorKey(colors)]
	return n[0] < n[1], n[0], n[1]
}

func (f *fakeMana) ThinSetFixing(_ []string, _ mtgv1.FormatId, colors []mtgv1.Color, _ *mtgv1.PowerLevel) (bool, int, int) {
	n, ok := f.fixing[colorKey(colors)]
	if !ok || n[1] == 0 {
		return false, 0, 0
	}
	return n[0] < n[1], n[0], n[1]
}

var (
	cW, cU, cB = mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B
	cR, cG     = mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G
)

// hobbitMana holds the counts of the Hobbit family on the snapshot of
// 2026-09-04, read 2026-10-31: 57 mana cards in all colors, 48 in white,
// blue, red, and green, and 16 in blue and black. The family holds 5
// fixing lands for the four colors, and 1 for blue and black. The floor
// of bracket 3 is 17 for four colors and 11 for two (D-799).
func hobbitMana() *fakeMana {
	return &fakeMana{
		identity: map[string][]mtgv1.Color{
			"Aragorn, the Uniter":      {cR, cG, cW, cU},
			"Gríma, Saruman's Footman": {cU, cB},
			"Smaug the Magnificent":    {cR},
		},
		mana: map[string][2]int{
			"":                                      {57, 46},
			colorKey([]mtgv1.Color{cR, cG, cW, cU}): {48, 46},
			colorKey([]mtgv1.Color{cU, cB}):         {16, 46},
			colorKey([]mtgv1.Color{cR}):             {30, 46},
			colorKey([]mtgv1.Color{cB, cR}):         {40, 26},
		},
		fixing: map[string][2]int{
			colorKey([]mtgv1.Color{cR, cG, cW, cU}): {5, 17},
			colorKey([]mtgv1.Color{cU, cB}):         {1, 11},
			colorKey([]mtgv1.Color{cB, cR}):         {2, 7},
		},
	}
}

// manaState is a set-limited deck of the Hobbit family. The theme is a
// decline, as in the deployed session of D-1227.
func manaState(format mtgv1.FormatId) *State {
	st := NewState(true)
	st.SessionID = "mana-row"
	st.Slots.Format = &mtgv1.Format{Id: format}
	st.Ctx.Format = format
	st.Slots.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}}
	st.Close("format")
	st.Skip("theme")
	st.SetLimit("the Hobbit set", []string{"hob", "hoc"}, []string{"The Hobbit", "The Hobbit Eternal"})
	return st
}

func manaAgent(t *testing.T, h Hints) *Agent {
	t.Helper()
	cat, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	return &Agent{cat: cat, threshold: DefaultFitThreshold, hints: h,
		log: slog.New(slog.NewTextHandler(io.Discard, nil))}
}

// manaRowText is the text of the mana row as it goes out, or "" when
// the planner does not ask it.
func manaRowText(a *Agent, st *State) string {
	for _, r := range a.cat.Plan(st.Ctx) {
		if r.ID == "set_outside_mana" {
			text, _ := substitute(r.Text, st, a.hints)
			return text
		}
	}
	return ""
}

// TestTheManaRowReadsTheFixingLands is D-1227 on the deployed case.
// Aragorn, the Uniter leads a deck of the Hobbit family: 48 mana cards
// passed the count of 46, and 5 fixing lands of 17 asked nothing. The
// deck shipped with 0.67 of its color sources and the grade bad.
func TestTheManaRowReadsTheFixingLands(t *testing.T) {
	h := hobbitMana()
	a := manaAgent(t, h)
	st := manaState(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	st.SetCommander("Aragorn, the Uniter")
	a.readFacts(st)
	if !st.Ctx.ThinSetMana {
		t.Fatal("5 fixing lands against a floor of 17 left the mana row silent")
	}
	text := manaRowText(a, st)
	want := "The Hobbit and The Hobbit Eternal hold 5 lands that make two or more colors of this deck, and this deck wants 17."
	if !strings.HasPrefix(text, want) {
		t.Errorf("the mana row reads %q, want it to start %q", text, want)
	}
}

// TestTheManaRowReadsTheCommanderColors is D-1227 for a commander that
// sets the colors. The row read the color slot, which a commander never
// writes, so a blue and black deck of the Hobbit family counted the 57
// mana cards of every color in place of its 16.
func TestTheManaRowReadsTheCommanderColors(t *testing.T) {
	h := hobbitMana()
	a := manaAgent(t, h)
	st := manaState(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	st.SetCommander("Gríma, Saruman's Footman")
	a.readFacts(st)
	if len(h.read) == 0 || !slices.Equal(h.read[len(h.read)-1], []mtgv1.Color{cU, cB}) {
		t.Errorf("the mana count read the colors %v, want the identity of the commander", h.read)
	}
	if !st.Ctx.ThinSetMana {
		t.Fatal("16 mana cards against 46 left the mana row silent")
	}
	want := "The Hobbit and The Hobbit Eternal hold 16 mana cards, and this deck wants 46."
	if text := manaRowText(a, st); !strings.HasPrefix(text, want) {
		t.Errorf("the mana row reads %q, want it to start %q", text, want)
	}
}

// TestTheManaRowReadsTheFixingOfASixtyCardDeck is D-1227 in a 60-card
// format. The slot names the colors, the mana cards suffice, and the
// sets hold 2 fixing lands against the floor of 7.
func TestTheManaRowReadsTheFixingOfASixtyCardDeck(t *testing.T) {
	a := manaAgent(t, hobbitMana())
	st := manaState(mtgv1.FormatId_FORMAT_ID_STANDARD)
	st.Slots.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: mtgv1.SixtyStep_SIXTY_STEP_FNM}}
	st.Slots.Colors = []mtgv1.Color{cB, cR}
	a.readFacts(st)
	if !st.Ctx.ThinSetMana {
		t.Fatal("2 fixing lands against a floor of 7 left the mana row silent")
	}
	if text := manaRowText(a, st); !strings.Contains(text, "hold 2 lands that make two or more colors of this deck, and this deck wants 7") {
		t.Errorf("the mana row reads %q", text)
	}
}

// TestTheManaRowLeavesAMonoColoredDeckAlone: a deck of one color needs
// no fixing, and 30 mana cards of 46 still ask.
func TestTheManaRowLeavesAMonoColoredDeckAlone(t *testing.T) {
	h := hobbitMana()
	h.mana[colorKey([]mtgv1.Color{cR})] = [2]int{50, 46}
	a := manaAgent(t, h)
	st := manaState(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	st.SetCommander("Smaug the Magnificent")
	a.readFacts(st)
	if st.Ctx.ThinSetMana {
		t.Errorf("a mono-red deck with 50 mana cards of 46 asked: %q", manaRowText(a, st))
	}
}

// TestThinSetFixingReadsTheFloor is D-1227 on the card index: the count
// is the fixing class of the profile, and the floor is the one of
// D-799 for the power and the count of colors.
func TestThinSetFixingReadsTheFloor(t *testing.T) {
	b, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	land := func(id, name, text string, colors ...mtgv1.Color) *mtgv1.Card {
		return &mtgv1.Card{OracleId: id, Name: name, TypeLine: "Land", CardTypes: []string{"Land"},
			OracleText: text, ColorIdentity: colors, ProducedMana: colors, Legalities: legal}
	}
	list := []*mtgv1.Card{
		land("gu", "Elvenking's Halls", "{T}: Add {G} or {U}.", cG, cU),
		land("rw", "Iron Hills", "{T}: Add {R} or {W}.", cR, cW),
		land("uw", "Lake-town", "{T}: Add {U} or {W}.", cU, cW),
		land("fetch", "Hobbit Hole", "{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle."),
		land("mono", "Rivendell", "{T}: Add {U}.", cU),
	}
	var printings []cards.Printing
	for _, c := range list {
		printings = append(printings, cards.Printing{ScryfallID: c.OracleId + "-1", OracleID: c.OracleId, Name: c.Name, SetCode: "hob", CollectorNumber: "1"})
	}
	h := &CandidateHints{Index: cards.NewIndex(list, printings, nil, time.Time{}), Builder: b}
	bracket3 := &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}}
	cmdr := mtgv1.FormatId_FORMAT_ID_COMMANDER
	codes := []string{"hob", "hoc"}
	if thin, have, want := h.ThinSetFixing(codes, cmdr, []mtgv1.Color{cW, cU, cR, cG}, bracket3); !thin || have != 4 || want != 17 {
		t.Errorf("four colors: thin %v, have %d, want %d, and the test wants true, 4, 17", thin, have, want)
	}
	if thin, have, want := h.ThinSetFixing(codes, cmdr, []mtgv1.Color{cU, cB}, bracket3); !thin || have != 1 || want != 11 {
		t.Errorf("blue and black: thin %v, have %d, want %d, and the test wants true, 1, 11", thin, have, want)
	}
	if thin, _, want := h.ThinSetFixing(codes, cmdr, []mtgv1.Color{cU}, bracket3); thin || want != 0 {
		t.Errorf("one color: thin %v, want %d, and the test wants no floor", thin, want)
	}
}
