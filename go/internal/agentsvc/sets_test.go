package agentsvc

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/questions"
)

// setIndex builds an index whose cards carry set codes, with a Hobbit
// family of two sets and one card outside it.
func setIndex(t *testing.T, inSet int) *cards.Index {
	t.Helper()
	var list []*mtgv1.Card
	var printings []cards.Printing
	add := func(id, name, typeLine, set string, commander bool, text ...string) {
		c := &mtgv1.Card{
			OracleId: id, Name: name, TypeLine: typeLine, CanBeCommander: commander,
			ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
			CardTypes:     []string{strings.Split(typeLine, " ")[0]},
			Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
		}
		if len(text) > 0 {
			c.OracleText = text[0]
		}
		list = append(list, c)
		printings = append(printings, cards.Printing{
			ScryfallID: id + "-p", OracleID: id, SetCode: set, CollectorNumber: "1",
		})
	}
	add("o-thranduil", "Thranduil, the Elvenking", "Legendary Creature", "hob", true)
	add("o-smaug", "Smaug the Impenetrable", "Legendary Creature", "hoc", true)
	add("o-karlov", "Karlov of the Ghost Council", "Legendary Creature", "m19", true)
	// A mana rock the named sets do not hold. It reaches the pool only
	// when the reader allowed the mana fill (D-382).
	add("o-solring", "Sol Ring", "Artifact", "m19", false, "{T}: Add {C}{C}.")
	// A basic land outside the named sets, for the precon checks that
	// skip basics (D-37, D-523).
	list = append(list, &mtgv1.Card{
		OracleId: "o-plains", Name: "Plains", TypeLine: "Basic Land — Plains",
		Supertypes: []string{"Basic"}, CardTypes: []string{"Land"},
		Legalities: map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	})
	printings = append(printings, cards.Printing{ScryfallID: "o-plains-p", OracleID: "o-plains", SetCode: "m19", CollectorNumber: "2"})
	for i := 0; i < inSet; i++ {
		add("o-in-"+string(rune('a'+i%26))+string(rune('a'+i/26)), "Hobbit Card "+string(rune('a'+i%26))+string(rune('a'+i/26)),
			"Enchantment", "hob", false)
	}
	return cards.NewIndex(list, printings, nil, time.Unix(1000, 0).UTC(),
		cards.WithSets([]cards.SetInfo{
			{Code: "hob", Name: "The Hobbit", Type: "expansion", ReleasedAt: "2026-08-14"},
			{Code: "hoc", Name: "The Hobbit Eternal", Type: "eternal", ReleasedAt: "2026-08-14", ParentCode: "hob"},
			{Code: "m19", Name: "Core Set 2019", Type: "core", ReleasedAt: "2018-07-13"},
		}))
}

// setServer wires a server with a set-carrying index and a fake generator.
func setServer(t *testing.T, fd *fakeDecks, inSet int) *Server {
	t.Helper()
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	client, _ := fakeClient(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	s, err := New(cat, client, newFakeStore(), func(context.Context) string { return "u1" },
		WithLogger(quiet), WithDecks(fd), WithCandidates(fixedIndex{setIndex(t, inSet)}, cb))
	if err != nil {
		t.Fatal(err)
	}
	return s
}

// setSession is a ready session limited to the Hobbit family.
func setSession(codes []string, commander string) (*mtgv1.Session, *questions.State) {
	session := &mtgv1.Session{
		Id: "s1",
		Slots: &mtgv1.Slots{
			Format:   &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
			Theme:    "hobbits",
			Colors:   []mtgv1.Color{mtgv1.Color_COLOR_W},
			SetCodes: codes,
			SlotStates: map[string]mtgv1.SlotState{
				"format": mtgv1.SlotState_SLOT_STATE_FILLED,
			},
		},
		Turns: []*mtgv1.Turn{{UserMessage: "a Hobbit-set deck"}},
	}
	st := questions.Restore("s1", session.Slots, questions.Snapshot{Version: questions.SnapshotVersion})
	if commander != "" {
		st.CommanderNames = []string{commander}
	}
	return session, st
}

// TestBuildPassesTheSetsToTheGenerator is D-373: the build carries the
// set limit through to the deck.
func TestBuildPassesTheSetsToTheGenerator(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	s := setServer(t, fd, 80)
	session, st := setSession([]string{"hob", "hoc"}, "Thranduil, the Elvenking")
	if _, err := s.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(fd.got.SetCodes, ","); got != "hob,hoc" {
		t.Errorf("generate request set codes = %q, want hob,hoc", got)
	}
	// The pool holds no card the sets do not have, apart from the ones
	// the build always adds.
	for _, name := range fd.got.Pool.Names() {
		c, _ := fd.got.Pool.Card(name)
		if c.GetOracleId() == "o-karlov" {
			t.Error("the pool holds a card outside the named sets")
		}
	}
}

// TestThinSetEndsTheTurnWithAReason is D-380: a family under the floor
// builds nothing, and it never builds a deck of another set.
func TestThinSetEndsTheTurnWithAReason(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	s := setServer(t, fd, 5)
	session, st := setSession([]string{"hob", "hoc"}, "Thranduil, the Elvenking")
	_, err := s.buildDeck(context.Background(), "u1", session, st, nil, nil, nil)
	var thin *ErrThinSet
	if !errors.As(err, &thin) {
		t.Fatalf("err = %v, want ErrThinSet", err)
	}
	if fd.runs != 0 {
		t.Errorf("the generator ran %d times, want 0: a thin set spends no model call", fd.runs)
	}
	if thin.Want != candidates.SetFloor(mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Errorf("want = %d, want the Commander floor", thin.Want)
	}
	// The message names the sets and the counts, so the reader can act.
	msg := thin.Error()
	for _, want := range []string{"The Hobbit", "The Hobbit Eternal", "drop the set limit"} {
		if !strings.Contains(msg, want) {
			t.Errorf("the reason does not say %q: %s", want, msg)
		}
	}
}

// TestNoSetLimitSkipsTheFloor: a deck with no set limit is never refused
// by a rule about sets.
func TestNoSetLimitSkipsTheFloor(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	s := setServer(t, fd, 2)
	session, st := setSession(nil, "Karlov of the Ghost Council")
	if _, err := s.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); err != nil {
		t.Fatalf("a build with no set limit failed: %v", err)
	}
	if fd.runs != 1 {
		t.Errorf("build runs = %d, want 1", fd.runs)
	}
}

// TestTheCommanderPoolStaysInTheSets is the second gate line: a set
// build that delegates the commander picks one from the named sets.
func TestTheCommanderPoolStaysInTheSets(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	s := setServer(t, fd, 80)
	session, st := setSession([]string{"hob", "hoc"}, "")
	if _, err := s.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if len(fd.got.Commanders) == 0 {
		t.Fatal("the build chose no commander")
	}
	if fd.got.Commanders[0] == "o-karlov" {
		t.Error("the build chose a commander outside the named sets")
	}
}

// TestManaFillNeedsThePermission is D-382: the fill runs only after the
// reader says yes.
func TestManaFillNeedsThePermission(t *testing.T) {
	for _, tc := range []struct {
		name  string
		state mtgv1.SlotState
		want  bool
	}{
		{"no answer", mtgv1.SlotState_SLOT_STATE_UNSPECIFIED, false},
		{"asked", mtgv1.SlotState_SLOT_STATE_ASKED, false},
		{"refused", mtgv1.SlotState_SLOT_STATE_SKIPPED, false},
		{"allowed", mtgv1.SlotState_SLOT_STATE_FILLED, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
			s := setServer(t, fd, 80)
			session, st := setSession([]string{"hob", "hoc"}, "Thranduil, the Elvenking")
			if tc.state != mtgv1.SlotState_SLOT_STATE_UNSPECIFIED {
				session.Slots.SlotStates[questions.SlotSetOutsideMana] = tc.state
			}
			if _, err := s.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); err != nil {
				t.Fatal(err)
			}
			_, got := fd.got.Pool.ByOracleID("o-solring")
			if got != tc.want {
				t.Errorf("the pool holds the outside mana rock = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestManaRolesFollowTheTargets is D-382: the fill reaches the wanted
// count of each mana role and no further.
func TestManaRolesFollowTheTargets(t *testing.T) {
	got := manaRoles(generate.TargetsFor(mtgv1.FormatId_FORMAT_ID_COMMANDER, nil))
	if got[mtgv1.CardRole_CARD_ROLE_RAMP] != 10 || got[mtgv1.CardRole_CARD_ROLE_LAND] != 36 {
		t.Errorf("mana roles = %v, want ramp 10 and land 36", got)
	}
	if len(got) != 2 {
		t.Errorf("the fill names %d roles, want the two mana roles alone", len(got))
	}
	// A tournament 60-card deck runs no ramp target, so the fill covers
	// the lands alone.
	sixty := manaRoles(generate.TargetsFor(mtgv1.FormatId_FORMAT_ID_MODERN,
		&mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT}}))
	if _, ok := sixty[mtgv1.CardRole_CARD_ROLE_RAMP]; ok {
		t.Error("a tournament 60-card deck has no ramp target, so the fill must name no ramp role")
	}
}

// TestTheSetNoteNamesEverySet is D-390. A reader who wrote one product
// name has no way to know it became two sets, and the deck marks every
// card the sets do not hold. A mark explains nothing until the reader
// knows the limit.
func TestTheSetNoteNamesEverySet(t *testing.T) {
	for _, tc := range []struct {
		names []string
		want  string
	}{
		{[]string{"The Hobbit"}, "I will build from The Hobbit only"},
		{[]string{"The Hobbit", "The Hobbit Eternal"},
			"I will build from The Hobbit and The Hobbit Eternal only"},
		{[]string{"Bloomburrow", "Bloomburrow Commander", "Bloomburrow Promos"},
			"I will build from Bloomburrow, Bloomburrow Commander, and Bloomburrow Promos only"},
	} {
		if got := setNote(tc.names); got != tc.want {
			t.Errorf("setNote(%v) = %q, want %q", tc.names, got, tc.want)
		}
	}
}
