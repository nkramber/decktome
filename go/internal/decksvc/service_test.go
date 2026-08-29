package decksvc

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/decks"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

type fixedIndex struct{ idx *cards.Index }

func (f fixedIndex) Current() *cards.Index { return f.idx }

type fakeCollections struct {
	counts map[string]int32
	err    error
	gotID  string
	gotUsr string
}

func (f *fakeCollections) OracleCounts(_ context.Context, userID, collectionID string) (map[string]int32, error) {
	f.gotUsr, f.gotID = userID, collectionID
	return f.counts, f.err
}

func loadIndex(t *testing.T) *cards.Index {
	t.Helper()
	f, err := os.Open("../cards/testdata/cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	list, err := cards.LoadCards(f, "cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	return cards.NewIndex(list, nil, nil, time.Date(2026, 8, 24, 9, 1, 0, 0, time.UTC))
}

func modernDeck(t *testing.T, idx *cards.Index) *mtgv1.Deck {
	t.Helper()
	sw, ok := idx.ByName("Soul Warden")
	if !ok {
		t.Fatal("fixture card missing")
	}
	pl, _ := idx.ByName("Plains")
	return &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN},
		Cards: []*mtgv1.DeckCard{
			{OracleId: sw.OracleId, Name: sw.Name, Count: 4},
			{OracleId: pl.OracleId, Name: pl.Name, Count: 56},
		},
	}
}

func newServer(t *testing.T, idx *cards.Index, opts ...Option) *Server {
	t.Helper()
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	return New(cfg, fixedIndex{idx}, opts...)
}

func codeOf(t *testing.T, err error) connect.Code {
	t.Helper()
	var ce *connect.Error
	if !errors.As(err, &ce) {
		t.Fatalf("want connect error, got %v", err)
	}
	return ce.Code()
}

func TestValidateHappyPath(t *testing.T) {
	idx := loadIndex(t)
	s := newServer(t, idx)
	resp, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{Deck: modernDeck(t, idx)}))
	if err != nil {
		t.Fatal(err)
	}
	res := resp.Msg.Result
	if !res.Passed {
		t.Errorf("want pass, findings %v", res.Findings)
	}
	if res.LegalityAsOf != "2026-08-24" {
		t.Errorf("legality_as_of = %q", res.LegalityAsOf)
	}
	if res.PoolRule != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
		t.Errorf("pool_rule = %v", res.PoolRule)
	}
	if res.Format != mtgv1.FormatId_FORMAT_ID_MODERN {
		t.Errorf("format = %v", res.Format)
	}
}

func TestValidateNoIndex(t *testing.T) {
	s := newServer(t, nil)
	_, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{Deck: &mtgv1.Deck{}}))
	if codeOf(t, err) != connect.CodeUnavailable {
		t.Errorf("code = %v", codeOf(t, err))
	}
}

func TestValidateNilDeck(t *testing.T) {
	s := newServer(t, loadIndex(t))
	_, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{}))
	if codeOf(t, err) != connect.CodeInvalidArgument {
		t.Errorf("code = %v", codeOf(t, err))
	}
}

func TestResolvePoolRule(t *testing.T) {
	cases := []struct {
		in   mtgv1.PoolRule
		id   string
		want mtgv1.PoolRule
	}{
		{mtgv1.PoolRule_POOL_RULE_UNSPECIFIED, "", mtgv1.PoolRule_POOL_RULE_ANY_CARD},
		{mtgv1.PoolRule_POOL_RULE_UNSPECIFIED, "col-1", mtgv1.PoolRule_POOL_RULE_OWNED_FIRST},
		{mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, "col-1", mtgv1.PoolRule_POOL_RULE_OWNED_ONLY},
		{mtgv1.PoolRule_POOL_RULE_ANY_CARD, "col-1", mtgv1.PoolRule_POOL_RULE_ANY_CARD},
	}
	for _, c := range cases {
		if got := resolvePoolRule(c.in, c.id); got != c.want {
			t.Errorf("resolvePoolRule(%v, %q) = %v, want %v", c.in, c.id, got, c.want)
		}
	}
}

func TestValidateOwnedFlow(t *testing.T) {
	idx := loadIndex(t)
	deck := modernDeck(t, idx)

	t.Run("owned rule without collection is invalid", func(t *testing.T) {
		s := newServer(t, idx)
		_, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{
			Deck: deck, PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY}))
		if codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("collection without source is unavailable", func(t *testing.T) {
		s := newServer(t, idx)
		_, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{
			Deck: deck, CollectionId: "col-1"}))
		if codeOf(t, err) != connect.CodeUnavailable {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("collection default is owned-first and warns", func(t *testing.T) {
		src := &fakeCollections{counts: map[string]int32{}}
		s := newServer(t, idx, WithCollections(src), WithUser(func(context.Context) string { return "u-1" }))
		resp, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{
			Deck: deck, CollectionId: "col-1"}))
		if err != nil {
			t.Fatal(err)
		}
		if src.gotUsr != "u-1" || src.gotID != "col-1" {
			t.Errorf("source got user %q collection %q", src.gotUsr, src.gotID)
		}
		res := resp.Msg.Result
		if res.PoolRule != mtgv1.PoolRule_POOL_RULE_OWNED_FIRST || !res.Passed {
			t.Errorf("pool_rule %v passed %v", res.PoolRule, res.Passed)
		}
		var warns int
		for _, f := range res.Findings {
			if f.Code == rules.CodeNotOwned && f.Severity == mtgv1.Severity_SEVERITY_WARN {
				warns++
			}
		}
		if warns == 0 {
			t.Error("want not_owned warn")
		}
	})
	t.Run("owned-only blocks", func(t *testing.T) {
		src := &fakeCollections{counts: map[string]int32{}}
		s := newServer(t, idx, WithCollections(src))
		resp, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{
			Deck: deck, CollectionId: "col-1", PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY}))
		if err != nil {
			t.Fatal(err)
		}
		if resp.Msg.Result.Passed {
			t.Error("owned-only with an empty collection must block")
		}
	})
	t.Run("missing collection is not found", func(t *testing.T) {
		src := &fakeCollections{err: status.Error(codes.NotFound, "no such collection")}
		s := newServer(t, idx, WithCollections(src))
		_, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{
			Deck: deck, CollectionId: "missing"}))
		if codeOf(t, err) != connect.CodeNotFound {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("a store failure is internal", func(t *testing.T) {
		src := &fakeCollections{err: errors.New("firestore down")}
		s := newServer(t, idx, WithCollections(src))
		_, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{
			Deck: deck, CollectionId: "col-1"}))
		if codeOf(t, err) != connect.CodeInternal {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("a path id is refused", func(t *testing.T) {
		src := &fakeCollections{counts: map[string]int32{}}
		s := newServer(t, idx, WithCollections(src))
		_, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{
			Deck: deck, CollectionId: "../x"}))
		if codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
		if src.gotID != "" {
			t.Error("the bad id reached the store")
		}
	})
}

// fakeDecks answers GetDeck and ListDecks.
type fakeDecks struct {
	decks map[string]*mtgv1.Deck
	err   error
	limit int
}

func (f *fakeDecks) Get(_ context.Context, _, id string) (*mtgv1.Deck, error) {
	if f.err != nil {
		return nil, f.err
	}
	d, ok := f.decks[id]
	if !ok {
		return nil, decks.ErrNotFound
	}
	return d, nil
}

func (f *fakeDecks) List(_ context.Context, _ string, limit int) ([]*mtgv1.Deck, error) {
	f.limit = limit
	if f.err != nil {
		return nil, f.err
	}
	var out []*mtgv1.Deck
	for _, d := range f.decks {
		out = append(out, d)
	}
	return out, nil
}

func asUser(uid string) Option { return WithUser(func(context.Context) string { return uid }) }

func TestGetDeck(t *testing.T) {
	ctx := context.Background()
	src := &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": {Id: "d1", Name: "a deck"}}}
	get := func(s *Server, id string) error {
		_, err := s.GetDeck(ctx, connect.NewRequest(&mtgv1.GetDeckRequest{DeckId: id}))
		return err
	}
	t.Run("unimplemented without a store", func(t *testing.T) {
		if err := get(newServer(t, nil, asUser("u1")), "d1"); codeOf(t, err) != connect.CodeUnimplemented {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("unauthenticated without a user", func(t *testing.T) {
		if err := get(newServer(t, nil, WithDecks(src)), "d1"); codeOf(t, err) != connect.CodeUnauthenticated {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("empty and path ids are invalid", func(t *testing.T) {
		s := newServer(t, nil, WithDecks(src), asUser("u1"))
		for _, id := range []string{"", "  ", "a/b", ".."} {
			if err := get(s, id); codeOf(t, err) != connect.CodeInvalidArgument {
				t.Errorf("id %q: code = %v", id, codeOf(t, err))
			}
		}
	})
	t.Run("not found", func(t *testing.T) {
		if err := get(newServer(t, nil, WithDecks(src), asUser("u1")), "nope"); codeOf(t, err) != connect.CodeNotFound {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("a store failure is internal", func(t *testing.T) {
		bad := &fakeDecks{err: errors.New("firestore down")}
		if err := get(newServer(t, nil, WithDecks(bad), asUser("u1")), "d1"); codeOf(t, err) != connect.CodeInternal {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("found", func(t *testing.T) {
		res, err := newServer(t, nil, WithDecks(src), asUser("u1")).GetDeck(ctx, connect.NewRequest(&mtgv1.GetDeckRequest{DeckId: "d1"}))
		if err != nil || res.Msg.GetDeck().GetName() != "a deck" {
			t.Errorf("deck = %v, err %v", res, err)
		}
	})
}

func TestListDecks(t *testing.T) {
	ctx := context.Background()
	src := &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": {Id: "d1"}, "d2": {Id: "d2"}}}
	list := func(s *Server) (*connect.Response[mtgv1.ListDecksResponse], error) {
		return s.ListDecks(ctx, connect.NewRequest(&mtgv1.ListDecksRequest{}))
	}
	t.Run("unimplemented without a store", func(t *testing.T) {
		if _, err := list(newServer(t, nil, asUser("u1"))); codeOf(t, err) != connect.CodeUnimplemented {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("unauthenticated without a user", func(t *testing.T) {
		if _, err := list(newServer(t, nil, WithDecks(src))); codeOf(t, err) != connect.CodeUnauthenticated {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("a store failure is internal", func(t *testing.T) {
		bad := &fakeDecks{err: errors.New("firestore down")}
		if _, err := list(newServer(t, nil, WithDecks(bad), asUser("u1"))); codeOf(t, err) != connect.CodeInternal {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("lists with the cap", func(t *testing.T) {
		res, err := list(newServer(t, nil, WithDecks(src), asUser("u1")))
		if err != nil || len(res.Msg.GetDecks()) != 2 {
			t.Errorf("decks = %v, err %v", res, err)
		}
		if src.limit != listLimit {
			t.Errorf("limit = %d, want %d", src.limit, listLimit)
		}
	})
}
