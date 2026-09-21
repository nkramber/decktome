package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/precons"
	"github.com/nkramber/decktome/go/internal/questions"
)

// heroIndex holds Karlov and heroes white Hobbit creatures, "o-hero-00"
// and up. The Hobbit subtype is the only theme signal of the pool.
func heroIndex(heroes int) *cards.Index {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	list := []*mtgv1.Card{{
		OracleId: "o-karlov", Name: "Karlov of the Ghost Council", TypeLine: "Legendary Creature — Spirit Advisor",
		CanBeCommander: true, ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
		CardTypes: []string{"Creature"}, Subtypes: []string{"Spirit", "Advisor"}, Legalities: legal,
	}}
	for i := range heroes {
		list = append(list, &mtgv1.Card{
			OracleId: heroID(i), Name: fmt.Sprintf("Hobbit Hero %02d", i), TypeLine: "Creature — Hobbit",
			ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
			CardTypes:     []string{"Creature"}, Subtypes: []string{"Hobbit"}, Legalities: legal,
		})
	}
	printings := make([]cards.Printing, 0, len(list))
	for _, c := range list {
		printings = append(printings, cards.Printing{ScryfallID: c.GetOracleId() + "-p", OracleID: c.GetOracleId(), SetCode: "m19", CollectorNumber: "1"})
	}
	return cards.NewIndex(list, printings, nil, time.Unix(1000, 0).UTC())
}

func heroID(i int) string { return fmt.Sprintf("o-hero-%02d", i) }

// heroPrecon is a product that holds the first n heroes.
func heroPrecon(n int) *precons.Table {
	var in []meta.PreconCard
	for i := range n {
		in = append(in, meta.PreconCard{Name: fmt.Sprintf("Hobbit Hero %02d", i), Count: 1, OracleID: heroID(i), ScryfallID: heroID(i) + "-p", SetCode: "m19", Number: "1"})
	}
	return precons.NewTable("v1", []meta.Precon{{
		Name: "Hero Precon", Code: "HRO", Type: "Commander Deck", ReleaseDate: "2026-08-14", Cards: in,
	}})
}

// heroServer wires a server over heroIndex and heroPrecon, and a session
// for a Hobbit deck from the library, with the precon excluded.
func heroServer(t *testing.T, fd *fakeDecks, heroes, inPrecon int) (*Server, *mtgv1.Session, *questions.State, map[string]int32) {
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
		WithLogger(quiet), WithDecks(fd), WithCandidates(fixedIndex{heroIndex(heroes)}, cb), WithPreconTable(fakeTable{heroPrecon(inPrecon)}))
	if err != nil {
		t.Fatal(err)
	}
	owned := map[string]int32{"o-karlov": 1}
	for i := range heroes {
		owned[heroID(i)] = 1
	}
	session, st := setSession(nil, "Karlov of the Ghost Council")
	session.Slots.PoolRule = mtgv1.PoolRule_POOL_RULE_OWNED_ONLY
	session.Slots.ExcludePreconKeys = []string{"HeroPrecon_HRO"}
	return s, session, st, owned
}

// TestExcludedThemeEndsTheTurnWithAReason is F-37: the library holds 40
// Hobbits, the excluded precon holds 25 of them, and 15 stay. The build
// stops before the model call, and the reason names the counts and the
// ways out that exist.
func TestExcludedThemeEndsTheTurnWithAReason(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	s, session, st, owned := heroServer(t, fd, 40, 25)
	_, err := s.buildDeck(context.Background(), "u1", session, st, owned, nil, nil)
	var thin *ErrThinTheme
	if !errors.As(err, &thin) {
		t.Fatalf("err = %v, want ErrThinTheme", err)
	}
	if fd.runs != 0 {
		t.Errorf("the generator ran %d times, want 0: a starved theme spends no model call", fd.runs)
	}
	if thin.Have != 15 || thin.Whole != 40 || thin.Want != candidates.ThinThemeFloor {
		t.Errorf("counts = %d of %d, floor %d, want 15 of 40, floor %d", thin.Have, thin.Whole, thin.Want, candidates.ThinThemeFloor)
	}
	msg := thin.Error()
	for _, want := range []string{"Hero Precon", `"hobbits"`, "15", "40", "name another theme", "start a new chat"} {
		if !strings.Contains(msg, want) {
			t.Errorf("the reason does not say %s: %s", want, msg)
		}
	}
}

// TestExcludedThemeFloor: the build goes on when enough of the theme
// stays after the exclusion, and when the library held too few before
// it, because the pool question of D-63 owns that case.
func TestExcludedThemeFloor(t *testing.T) {
	for _, tc := range []struct {
		name             string
		heroes, inPrecon int
	}{
		{"enough stay", 40, 5},
		{"thin before the exclusion", 25, 20},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
			s, session, st, owned := heroServer(t, fd, tc.heroes, tc.inPrecon)
			if _, err := s.buildDeck(context.Background(), "u1", session, st, owned, nil, nil); err != nil {
				t.Fatalf("err = %v, want a build", err)
			}
			if fd.runs != 1 {
				t.Errorf("the generator ran %d times, want 1", fd.runs)
			}
		})
	}
}

// TestExcludedThemeSparesARevision: a revision keeps the deck the reader
// has, so the floor does not stop it.
func TestExcludedThemeSparesARevision(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	s, session, st, owned := heroServer(t, fd, 40, 25)
	rev := &generate.Revision{BaseDeckID: "d1"}
	if _, err := s.buildDeckFrom(context.Background(), "u1", session, st, owned, nil, rev, nil); err != nil {
		t.Fatalf("err = %v, want a revision", err)
	}
}
