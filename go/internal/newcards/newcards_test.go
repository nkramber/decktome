package newcards

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"slices"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

var (
	legalCmd = map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	green    = []mtgv1.Color{mtgv1.Color_COLOR_G}
	red      = []mtgv1.Color{mtgv1.Color_COLOR_R}
)

func card(id, name string, identity []mtgv1.Color, legal map[string]mtgv1.LegalityStatus) *mtgv1.Card {
	return &mtgv1.Card{OracleId: id, Name: name, ColorIdentity: identity, Legalities: legal}
}

func cmdKey(f mtgv1.FormatId) string {
	if f == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		return "commander"
	}
	return ""
}

func commanderDeck(id, session string, commander string) *mtgv1.Deck {
	return &mtgv1.Deck{
		Id: id, SessionId: session,
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		CommanderOracleIds: []string{commander},
		Cards:              []*mtgv1.DeckCard{{OracleId: "held", Count: 1}},
	}
}

// TestPick holds each rule of a fit (D-1091): the floor, a card the deck
// holds, the legality, the colors, the order, and the cap.
func TestPick(t *testing.T) {
	fresh := []*mtgv1.Card{
		card("a", "Alpha", green, legalCmd),
		card("b", "Beta", green, legalCmd),
		card("c", "Gamma", green, legalCmd),
		card("d", "Delta", green, legalCmd),
		card("low", "Low", green, legalCmd),
		card("held", "Held", green, legalCmd),
		card("ban", "Banned", green, map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_BANNED}),
		card("off", "Off Color", red, legalCmd),
	}
	scores := []float64{0.5, 0.9, 0.5, 0.4, 0.1, 1, 1, 1}
	d := commanderDeck("d1", "", "cmd")
	colors := map[mtgv1.Color]bool{mtgv1.Color_COLOR_G: true}
	got := Pick(d, colors, fresh, scores, "commander", 0.32)
	if want := []string{"b", "a", "c"}; !slices.Equal(got, want) {
		t.Fatalf("Pick = %v, want %v", got, want)
	}
	if got := Pick(d, colors, fresh, scores, "", 0.32); got != nil {
		t.Fatalf("Pick with no format key = %v, want none", got)
	}
}

// TestDeckColors reads the commanders, and every card of a deck with no
// commander.
func TestDeckColors(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		card("cmd", "Commander", green, legalCmd),
		card("r", "Red Card", red, legalCmd),
	}, nil, nil, time.Now())
	d := &mtgv1.Deck{CommanderOracleIds: []string{"cmd"}, Cards: []*mtgv1.DeckCard{{OracleId: "r"}}}
	if got := DeckColors(idx, d); len(got) != 1 || !got[mtgv1.Color_COLOR_G] {
		t.Fatalf("commander deck colors = %v, want green alone", got)
	}
	d.CommanderOracleIds = nil
	d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: "cmd"})
	if got := DeckColors(idx, d); len(got) != 2 {
		t.Fatalf("sixty-card deck colors = %v, want green and red", got)
	}
}

type fakeStore struct {
	decks map[string]*mtgv1.Deck // uid/id -> deck
	order []string
	marks int
}

func (s *fakeStore) put(uid string, d *mtgv1.Deck) {
	if s.decks == nil {
		s.decks = map[string]*mtgv1.Deck{}
	}
	s.decks[uid+"/"+d.GetId()] = d
	s.order = append(s.order, uid+"/"+d.GetId())
}

func (s *fakeStore) Scan(_ context.Context, fn func(uid string, d *mtgv1.Deck) error) error {
	for _, k := range s.order {
		uid := k[:slices.Index([]byte(k), '/')]
		if err := fn(uid, proto.Clone(s.decks[k]).(*mtgv1.Deck)); err != nil {
			return err
		}
	}
	return nil
}

func (s *fakeStore) Mark(_ context.Context, uid, id string, fn func(d *mtgv1.Deck) bool) (bool, error) {
	cur := proto.Clone(s.decks[uid+"/"+id]).(*mtgv1.Deck)
	if !fn(cur) {
		return false, nil
	}
	s.marks++
	s.decks[uid+"/"+id] = cur
	return true, nil
}

// fakeFit scores a card 1 when the theme names its oracle id, and names
// the theme "deck" for a deck with no chat.
type fakeFit struct{ deckTheme string }

func (f fakeFit) ThemeScores(theme string, _ *cards.Index, cs []*mtgv1.Card) []float64 {
	out := make([]float64, len(cs))
	for i, c := range cs {
		if theme == c.GetOracleId() || theme == "deck" && c.GetOracleId() == "b" {
			out[i] = 1
		}
	}
	return out
}

func (f fakeFit) DeckTheme(*cards.Index, []*mtgv1.Card) string { return f.deckTheme }

// TestPass is D-1091 and D-1092: a deck takes the cards of the theme of
// its chat, a deck with no chat takes the theme of its cards, a rerun
// writes and pushes nothing, a rerun after a dismiss keeps the dismiss,
// and a pass with no fit clears the last cards with no push.
func TestPass(t *testing.T) {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	idx := cards.NewIndex([]*mtgv1.Card{
		card("cmd", "Commander", green, legalCmd),
		card("a", "Alpha", green, legalCmd),
		card("b", "Beta", green, legalCmd),
	}, nil, nil, time.Now())
	store := &fakeStore{}
	store.put("u1", commanderDeck("chat", "s1", "cmd"))
	store.put("u1", commanderDeck("gone", "s-gone", "cmd"))
	stale := commanderDeck("old", "s-none", "cmd")
	stale.NewOracleIds = []string{"z"}
	store.put("u2", stale)
	themes := map[string]string{"u1/s1": "a"}
	in := Input{
		Index: idx, Version: "v1", New: []string{"a", "b", "missing"}, KeyOf: cmdKey, Fit: fakeFit{deckTheme: "deck"}, Floor: 0.32,
		ThemeOf: func(_ context.Context, uid, sid string) (string, error) {
			return themes[uid+"/"+sid], nil
		},
	}
	res, err := Pass(ctx, store, in, log)
	if err != nil {
		t.Fatal(err)
	}
	if got := store.decks["u1/chat"].GetNewOracleIds(); !slices.Equal(got, []string{"a"}) {
		t.Errorf("chat deck = %v, want [a] from the theme of its chat", got)
	}
	if got := store.decks["u1/gone"].GetNewOracleIds(); !slices.Equal(got, []string{"b"}) {
		t.Errorf("deck with no chat theme = %v, want [b] from its cards", got)
	}
	if got := store.decks["u2/old"].GetNewOracleIds(); !slices.Equal(got, []string{"b"}) {
		t.Errorf("deck with old cards = %v, want [b]", got)
	}
	if res.Read != 3 || res.Fit != 3 || res.Written != 3 || len(res.Hit["u1"]) != 2 || len(res.Hit["u2"]) != 1 {
		t.Errorf("first pass = %+v", res)
	}

	again, err := Pass(ctx, store, in, log)
	if err != nil {
		t.Fatal(err)
	}
	if again.Written != 0 || len(again.Hit) != 0 {
		t.Errorf("a rerun wrote %d decks and hit %d users, want none", again.Written, len(again.Hit))
	}

	// The pass of v1 failed after it marked u1/chat, and the user then
	// dismissed the panel. The rerun of v1 keeps the dismiss.
	store.decks["u1/chat"].NewOracleIds = nil
	rerun, err := Pass(ctx, store, in, log)
	if err != nil {
		t.Fatal(err)
	}
	if rerun.Written != 0 || len(rerun.Hit) != 0 || len(store.decks["u1/chat"].GetNewOracleIds()) != 0 {
		t.Errorf("a rerun after a dismiss = %+v, cards %v, want no write and no push", rerun, store.decks["u1/chat"].GetNewOracleIds())
	}

	// A new marker with the same cards writes its version and sends no push.
	in.Version = "v2"
	same, err := Pass(ctx, store, in, log)
	if err != nil {
		t.Fatal(err)
	}
	if same.Written != 3 || len(same.Hit["u1"]) != 1 || len(same.Hit["u2"]) != 0 {
		t.Errorf("a new marker with the same cards = %+v, want three writes and one push for the dismissed deck", same)
	}

	in.Version, in.New = "v3", nil
	cleared, err := Pass(ctx, store, in, log)
	if err != nil {
		t.Fatal(err)
	}
	if cleared.Written != 3 || len(cleared.Hit) != 0 || len(store.decks["u1/chat"].GetNewOracleIds()) != 0 {
		t.Errorf("a pass with no fit = %+v, want three cleared decks and no push", cleared)
	}

	boom := errors.New("read failed")
	in.ThemeOf = func(context.Context, string, string) (string, error) { return "", boom }
	in.Version = "v4"
	if _, err := Pass(ctx, store, in, log); !errors.Is(err, boom) {
		t.Errorf("a failed theme read = %v, want the error", err)
	}
	in.Version = ""
	if _, err := Pass(ctx, store, in, log); err == nil {
		t.Error("a pass with no version ran, want an error")
	}
}
