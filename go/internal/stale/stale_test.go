package stale

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// card is one main-list entry of a test deck.
func card(id string, n int32, role mtgv1.CardRole) *mtgv1.DeckCard {
	return &mtgv1.DeckCard{OracleId: id, Name: "Card " + id, Count: n, Role: role}
}

// commanderDeck is a Commander deck of 37 lands and 62 nonland cards,
// with commander "cmd". The nonland cards are n00 to n61.
func commanderDeck() *mtgv1.Deck {
	d := &mtgv1.Deck{
		Id:                 "d1",
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		CommanderOracleIds: []string{"cmd"},
		Commanders:         []*mtgv1.DeckCard{{OracleId: "cmd", Name: "Commander Card", Count: 1}},
		Cards:              []*mtgv1.DeckCard{card("land", 37, mtgv1.CardRole_CARD_ROLE_LAND)},
	}
	for i := range 62 {
		d.Cards = append(d.Cards, card(fmt.Sprintf("n%02d", i), 1, mtgv1.CardRole_CARD_ROLE_SYNERGY))
	}
	return d
}

// sixtyDeck is a 60-card deck of 24 lands and nine 4-of nonland cards,
// 36 nonland copies in all.
func sixtyDeck() *mtgv1.Deck {
	d := &mtgv1.Deck{
		Id:     "d2",
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN},
		Cards:  []*mtgv1.DeckCard{card("land", 24, mtgv1.CardRole_CARD_ROLE_LAND)},
	}
	for i := range 9 {
		d.Cards = append(d.Cards, card(fmt.Sprintf("s%d", i), 4, mtgv1.CardRole_CARD_ROLE_THREAT))
	}
	return d
}

func legalAll(d *mtgv1.Deck, key string) Legalities {
	out := Legalities{}
	for _, id := range append(slices.Clone(d.GetCommanderOracleIds()), "land") {
		out[id] = map[string]string{key: "legal"}
	}
	for _, dc := range d.GetCards() {
		out[dc.GetOracleId()] = map[string]string{key: "legal"}
	}
	return out
}

func TestFind(t *testing.T) {
	d := commanderDeck()
	d.Sideboard = []*mtgv1.DeckCard{card("side", 1, mtgv1.CardRole_CARD_ROLE_OTHER)}
	d.CompanionOracleId = "comp"
	legal := legalAll(d, "commander")
	legal["n03"]["commander"] = "banned"
	legal["n01"]["commander"] = "not_legal"
	legal["n02"]["commander"] = "restricted"
	legal["side"] = map[string]string{"commander": "banned"}
	legal["comp"] = map[string]string{"commander": "banned"}
	delete(legal, "n04") // a card the snapshot does not hold is not a ban (REV-059)
	got := Find(d, "commander", legal)
	want := []string{"comp", "n01", "n03", "side"}
	if !slices.Equal(got, want) {
		t.Fatalf("Find = %v, want %v", got, want)
	}
	if got := Find(d, "", legal); got != nil {
		t.Fatalf("a house format has no legality check (D-3), Find = %v", got)
	}
}

// TestClassify holds the rule of D-1008 on its numbers: 7 of 62 nonland
// cards in Commander, a 4-of and a 3-of of 36 nonland copies in a
// 60-card deck, the commander, a win condition, and an imported list.
func TestClassify(t *testing.T) {
	patch, rebuild := mtgv1.RerunCase_RERUN_CASE_PATCH, mtgv1.RerunCase_RERUN_CASE_REBUILD
	ids := func(n int) []string {
		var out []string
		for i := range n {
			out = append(out, fmt.Sprintf("n%02d", i))
		}
		return out
	}
	cases := []struct {
		name   string
		deck   func() *mtgv1.Deck
		ids    []string
		want   mtgv1.RerunCase
		reason string
	}{
		{"no stale card", commanderDeck, nil, mtgv1.RerunCase_RERUN_CASE_UNSPECIFIED, ""},
		{"one filler card", commanderDeck, ids(1), patch, "replaces Card n00 alone"},
		{"6 of 62 nonland", commanderDeck, ids(6), patch, "replaces"},
		{"7 of 62 nonland", commanderDeck, ids(7), rebuild, "removes 7 of the 62 nonland cards"},
		{"the commander", commanderDeck, []string{"cmd"}, rebuild, "Commander Card is no longer legal as your commander"},
		{"a win condition", func() *mtgv1.Deck {
			d := commanderDeck()
			d.Cards[5].Role = mtgv1.CardRole_CARD_ROLE_WINCON
			return d
		}, []string{"n04"}, rebuild, "Card n04 was a win condition"},
		{"a 3-of of 36", func() *mtgv1.Deck {
			d := sixtyDeck()
			d.Cards[1].Count = 3
			return d
		}, []string{"s0"}, patch, "replaces Card s0 alone"},
		{"a 4-of of 36", sixtyDeck, []string{"s0"}, rebuild, "removes 4 of the 36 nonland cards"},
		{"an imported list", func() *mtgv1.Deck {
			d := commanderDeck()
			d.Imported = true
			return d
		}, ids(10), patch, "You imported this list"},
		{"an imported list without its commander", func() *mtgv1.Deck {
			d := commanderDeck()
			d.Imported = true
			return d
		}, []string{"cmd"}, rebuild, "pick a new commander"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			v := Classify(c.deck(), c.ids)
			if v.Case != c.want {
				t.Fatalf("case = %v, want %v (%s)", v.Case, c.want, v.Reason)
			}
			if !strings.Contains(v.Reason, c.reason) {
				t.Fatalf("reason = %q, want it to hold %q", v.Reason, c.reason)
			}
		})
	}
}

func TestStaleNames(t *testing.T) {
	d := commanderDeck()
	v := Classify(d, []string{"n00", "n01", "n02"})
	if want := "Card n00, Card n01 and Card n02"; !strings.Contains(v.Reason, want) {
		t.Fatalf("reason = %q, want %q", v.Reason, want)
	}
}

func TestApply(t *testing.T) {
	d := commanderDeck()
	v := Classify(d, []string{"n00"})
	if !Apply(d, []string{"n00"}, v) {
		t.Fatal("a new stale card must change the deck")
	}
	if !d.GetStale() || d.GetRerunCase() != mtgv1.RerunCase_RERUN_CASE_PATCH || d.GetStaleReason() == "" {
		t.Fatalf("deck = stale %v case %v reason %q", d.GetStale(), d.GetRerunCase(), d.GetStaleReason())
	}
	if Apply(d, []string{"n00"}, v) {
		t.Fatal("the same state must not count as a change")
	}
	if !Apply(d, nil, Verdict{}) || d.GetStale() || len(d.GetStaleOracleIds()) > 0 || d.GetRerunCase() != 0 || d.GetStaleReason() != "" {
		t.Fatal("an empty list must clear the stale state")
	}
}

// fakeStore holds decks in memory. Mark counts its writes.
type fakeStore struct {
	decks  map[string]*mtgv1.Deck
	writes int
}

func (f *fakeStore) Scan(_ context.Context, fn func(uid string, d *mtgv1.Deck) error) error {
	for _, id := range slices.Sorted(func(yield func(string) bool) {
		for k := range f.decks {
			if !yield(k) {
				return
			}
		}
	}) {
		if err := fn("u1", cloneDeck(f.decks[id])); err != nil {
			return err
		}
	}
	return nil
}

func (f *fakeStore) Mark(_ context.Context, _, id string, fn func(d *mtgv1.Deck) bool) (bool, error) {
	cur := cloneDeck(f.decks[id])
	if !fn(cur) {
		return false, nil
	}
	f.decks[id] = cur
	f.writes++
	return true, nil
}

func cloneDeck(d *mtgv1.Deck) *mtgv1.Deck {
	out := &mtgv1.Deck{}
	*out = mtgv1.Deck{
		Id: d.GetId(), Format: d.GetFormat(), CommanderOracleIds: d.GetCommanderOracleIds(), Commanders: d.GetCommanders(),
		Cards: d.GetCards(), Imported: d.GetImported(), Stale: d.GetStale(), StaleOracleIds: d.GetStaleOracleIds(),
		RerunCase: d.GetRerunCase(), StaleReason: d.GetStaleReason(),
	}
	return out
}

func keyOf(f mtgv1.FormatId) string {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		return "commander"
	case mtgv1.FormatId_FORMAT_ID_MODERN:
		return "modern"
	}
	return ""
}

// TestPass marks the deck a ban hits, writes nothing for a deck that
// stays legal, and clears a deck after an unban.
func TestPass(t *testing.T) {
	cmd, sixty := commanderDeck(), sixtyDeck()
	legal := legalAll(cmd, "commander")
	for id, by := range legalAll(sixty, "modern") {
		if legal[id] == nil {
			legal[id] = map[string]string{}
		}
		for k, v := range by {
			legal[id][k] = v
		}
	}
	legal["n00"]["commander"] = "banned"
	store := &fakeStore{decks: map[string]*mtgv1.Deck{"d1": cmd, "d2": sixty}}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	res, err := Pass(context.Background(), store, keyOf, legal, log)
	if err != nil {
		t.Fatal(err)
	}
	if res != (Result{Read: 2, Stale: 1, Written: 1}) {
		t.Fatalf("first pass = %+v", res)
	}
	if d := store.decks["d1"]; !d.GetStale() || !slices.Equal(d.GetStaleOracleIds(), []string{"n00"}) {
		t.Fatalf("d1 = stale %v ids %v", d.GetStale(), d.GetStaleOracleIds())
	}

	res, err = Pass(context.Background(), store, keyOf, legal, log)
	if err != nil {
		t.Fatal(err)
	}
	if res.Written != 0 || store.writes != 1 {
		t.Fatalf("a second pass on the same snapshot must write nothing, %+v, writes %d", res, store.writes)
	}

	legal["n00"]["commander"] = "legal"
	res, err = Pass(context.Background(), store, keyOf, legal, log)
	if err != nil {
		t.Fatal(err)
	}
	if res.Written != 1 || store.decks["d1"].GetStale() {
		t.Fatalf("an unban must clear the deck, %+v", res)
	}
}
