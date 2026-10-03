package main

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"slices"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/newcards"
)

// oneDeck is the deck store of the test: one Commander deck of user u1.
type oneDeck struct{ deck *mtgv1.Deck }

func (o *oneDeck) Scan(_ context.Context, fn func(uid string, d *mtgv1.Deck) error) error {
	return fn("u1", &mtgv1.Deck{Id: o.deck.GetId(), Format: o.deck.GetFormat(), SessionId: o.deck.GetSessionId(),
		CommanderOracleIds: o.deck.GetCommanderOracleIds(), NewOracleIds: o.deck.GetNewOracleIds()})
}

func (o *oneDeck) Mark(_ context.Context, _, _ string, fn func(d *mtgv1.Deck) bool) (bool, error) {
	return fn(o.deck), nil
}

// allFit scores every card 1 for every theme.
type allFit struct{}

func (allFit) ThemeScores(_ string, _ *cards.Index, cs []*mtgv1.Card) []float64 {
	out := make([]float64, len(cs))
	for i := range out {
		out[i] = 1
	}
	return out
}

func (allFit) DeckTheme(*cards.Index, []*mtgv1.Card) string { return "" }

// TestNewCardsPass runs the pass on the newest new-cards marker, sends
// one push, records the end in the marker, and opens no deck store on
// the next run (D-1091).
func TestNewCardsPass(t *testing.T) {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	store := cards.DirStore{Root: t.TempDir()}
	putVersion(t, store, "20261113T090000")
	if err := cards.WriteNewCards(ctx, store, "20261113T090000", cards.NewCardsRecord{
		SnapshotAsOf: "2026-11-13T09:00:00Z", Cards: []string{"new"},
	}); err != nil {
		t.Fatal(err)
	}
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "cmd", Name: "Commander", Legalities: legal},
		{OracleId: "new", Name: "New Card", Legalities: legal},
	}, nil, nil, time.Now())
	load := func(context.Context, string) (*cards.Index, error) { return idx, nil }
	deck := &mtgv1.Deck{Id: "d1", SessionId: "s1", CommanderOracleIds: []string{"cmd"},
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}}
	var sent []string
	opens := 0
	open := func(context.Context) (newcards.Store, themeOf, notifyStale, func(), error) {
		opens++
		theme := func(context.Context, string, string) (string, error) { return "dinosaurs", nil }
		return &oneDeck{deck: deck}, theme, recordPush(&sent), func() {}, nil
	}
	now := func() time.Time { return time.Date(2026, 11, 13, 10, 0, 0, 0, time.UTC) }

	if err := newCardsPass(ctx, store, load, open, allFit{}, log, now); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(deck.GetNewOracleIds(), []string{"new"}) {
		t.Errorf("deck new cards = %v, want [new]", deck.GetNewOracleIds())
	}
	if len(sent) != 1 {
		t.Errorf("pushes = %v, want one for u1", sent)
	}
	rec, ok, err := cards.ReadNewCards(ctx, store, "20261113T090000")
	if err != nil || !ok || rec.Pass != "2026-11-13T10:00:00Z" || rec.Decks != 1 {
		t.Fatalf("marker = %+v ok=%v err=%v, want an ended pass with one deck", rec, ok, err)
	}
	if err := newCardsPass(ctx, store, load, open, allFit{}, log, now); err != nil {
		t.Fatal(err)
	}
	if opens != 1 || len(sent) != 1 {
		t.Errorf("an ended pass opened the store %d times and sent %v, want once and one push", opens, sent)
	}
}

// TestNewCardsPassNoMarker opens nothing and loads no index when no
// snapshot made a card legal for the first time.
func TestNewCardsPassNoMarker(t *testing.T) {
	ctx := context.Background()
	store := cards.DirStore{Root: t.TempDir()}
	putVersion(t, store, "20261113T090000")
	load := func(context.Context, string) (*cards.Index, error) { return nil, errors.New("loaded") }
	open := func(context.Context) (newcards.Store, themeOf, notifyStale, func(), error) {
		return nil, nil, nil, nil, errors.New("opened")
	}
	if err := newCardsPass(ctx, store, load, open, allFit{}, slog.New(slog.NewTextHandler(io.Discard, nil)), time.Now); err != nil {
		t.Fatalf("no marker = %v, want nil", err)
	}
}
