package decksvc

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/decks"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
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

// fakeDecks answers GetDeck, ListDecks, UpdateDeck, and DeleteDeck. The
// order slice fixes the listing order, because a map has none and the
// paging tests read positions.
type fakeDecks struct {
	decks   map[string]*mtgv1.Deck
	order   []string
	err     error
	scan    int
	filter  decks.Filter
	deleted []string
	updates int
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

func (f *fakeDecks) List(_ context.Context, _ string, filter decks.Filter, scan int) ([]*mtgv1.Deck, error) {
	f.scan = scan
	f.filter = filter
	if f.err != nil {
		return nil, f.err
	}
	var out []*mtgv1.Deck
	for _, id := range f.listOrder() {
		out = append(out, f.decks[id])
	}
	return out, nil
}

// listOrder gives the fixed order, or the map keys sorted when the test
// named none.
func (f *fakeDecks) listOrder() []string {
	if len(f.order) > 0 {
		return f.order
	}
	ids := make([]string, 0, len(f.decks))
	for id := range f.decks {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func (f *fakeDecks) Update(_ context.Context, _, id string, name *string, favorite *bool) (*mtgv1.Deck, error) {
	f.updates++
	if f.err != nil {
		return nil, f.err
	}
	d, ok := f.decks[id]
	if !ok {
		return nil, decks.ErrNotFound
	}
	out := proto.Clone(d).(*mtgv1.Deck) //nolint:errcheck,forcetypeassert // Clone of a Deck is a Deck
	if name != nil {
		out.Name = *name
	}
	if favorite != nil {
		out.Favorite = *favorite
	}
	f.decks[id] = out
	return out, nil
}

func (f *fakeDecks) Delete(_ context.Context, _, id string) error {
	if f.err != nil {
		return f.err
	}
	if _, ok := f.decks[id]; !ok {
		return decks.ErrNotFound
	}
	delete(f.decks, id)
	f.deleted = append(f.deleted, id)
	return nil
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
	t.Run("lists with the scan cap", func(t *testing.T) {
		res, err := list(newServer(t, nil, WithDecks(src), asUser("u1")))
		if err != nil || len(res.Msg.GetDecks()) != 2 {
			t.Errorf("decks = %v, err %v", res, err)
		}
		if src.scan != maxScan {
			t.Errorf("scan = %d, want %d", src.scan, maxScan)
		}
	})
}

func TestExportDeck(t *testing.T) {
	ctx := context.Background()
	idx := loadIndex(t)
	sol, _ := idx.ByName("Sol Ring")
	pl, _ := idx.ByName("Plains")
	owned := &mtgv1.DeckCard{OracleId: pl.OracleId, Name: pl.Name, Count: 30, Owned: true, OwnedCount: 40}
	deck := &mtgv1.Deck{Id: "d1", Name: "Rocks", Cards: []*mtgv1.DeckCard{{OracleId: sol.OracleId, Name: "sol ring", Count: 1}, owned}}
	src := &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": deck}}
	export := func(s *Server, id string, f mtgv1.ExportFormat) (*connect.Response[mtgv1.ExportDeckResponse], error) {
		return s.ExportDeck(ctx, connect.NewRequest(&mtgv1.ExportDeckRequest{DeckId: id, Format: f}))
	}
	t.Run("arena text with the index names and printings", func(t *testing.T) {
		res, err := export(newServer(t, idx, WithDecks(src), asUser("u1")), "d1", mtgv1.ExportFormat_EXPORT_FORMAT_UNSPECIFIED)
		if err != nil {
			t.Fatal(err)
		}
		dp := sol.GetDefaultPrinting()
		if want := "Deck\n1 Sol Ring (" + strings.ToUpper(dp.GetSetCode()) + ") " + dp.GetCollectorNumber() + "\n"; !strings.HasPrefix(res.Msg.GetText(), want) {
			t.Errorf("text %q, want prefix %q", res.Msg.GetText(), want)
		}
		if res.Msg.GetFileName() != "rocks.txt" {
			t.Errorf("file name %q", res.Msg.GetFileName())
		}
	})
	t.Run("buy list", func(t *testing.T) {
		res, err := export(newServer(t, idx, WithDecks(src), asUser("u1")), "d1", mtgv1.ExportFormat_EXPORT_FORMAT_BUY_LIST_TEXT)
		if err != nil || res.Msg.GetText() != "1 Sol Ring\n" || res.Msg.GetFileName() != "rocks-buy-list.txt" {
			t.Errorf("res %v, err %v", res, err)
		}
	})
	t.Run("no index falls back to the stored names", func(t *testing.T) {
		res, err := export(newServer(t, nil, WithDecks(src), asUser("u1")), "d1", mtgv1.ExportFormat_EXPORT_FORMAT_ARENA_TEXT)
		if err != nil || res.Msg.GetText() != "Deck\n1 sol ring\n30 Plains\n" {
			t.Errorf("res %v, err %v", res, err)
		}
	})
	t.Run("the read errors pass through", func(t *testing.T) {
		if _, err := export(newServer(t, idx, WithDecks(src), asUser("u1")), "nope", mtgv1.ExportFormat_EXPORT_FORMAT_ARENA_TEXT); codeOf(t, err) != connect.CodeNotFound {
			t.Errorf("code = %v", codeOf(t, err))
		}
		if _, err := export(newServer(t, idx, WithDecks(src)), "d1", mtgv1.ExportFormat_EXPORT_FORMAT_ARENA_TEXT); codeOf(t, err) != connect.CodeUnauthenticated {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
	t.Run("an unknown format is invalid", func(t *testing.T) {
		if _, err := export(newServer(t, idx, WithDecks(src), asUser("u1")), "d1", mtgv1.ExportFormat(99)); codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
}

// manyDecks builds n decks named d1..dn in a fixed order, newest first.
func manyDecks(n int) *fakeDecks {
	f := &fakeDecks{decks: make(map[string]*mtgv1.Deck, n)}
	for i := 1; i <= n; i++ {
		id := fmt.Sprintf("d%d", i)
		f.decks[id] = &mtgv1.Deck{Id: id, Name: id}
		f.order = append(f.order, id)
	}
	return f
}

func TestListDecksPaging(t *testing.T) {
	ctx := context.Background()
	list := func(src *fakeDecks, req *mtgv1.ListDecksRequest) (*connect.Response[mtgv1.ListDecksResponse], error) {
		return newServer(t, nil, WithDecks(src), asUser("u1")).ListDecks(ctx, connect.NewRequest(req))
	}
	ids := func(res *connect.Response[mtgv1.ListDecksResponse]) []string {
		var out []string
		for _, d := range res.Msg.GetDecks() {
			out = append(out, d.GetId())
		}
		return out
	}

	t.Run("a page of the default size, then the next", func(t *testing.T) {
		src := manyDecks(defaultPageSize + 3)
		first, err := list(src, &mtgv1.ListDecksRequest{})
		if err != nil {
			t.Fatal(err)
		}
		if got := len(first.Msg.GetDecks()); got != defaultPageSize {
			t.Fatalf("first page = %d decks, want %d", got, defaultPageSize)
		}
		if first.Msg.GetNextPageToken() == "" {
			t.Fatal("want a next token")
		}
		second, err := list(src, &mtgv1.ListDecksRequest{PageToken: first.Msg.GetNextPageToken()})
		if err != nil {
			t.Fatal(err)
		}
		if got := ids(second); len(got) != 3 || got[0] != "d25" {
			t.Fatalf("second page = %v", got)
		}
		if second.Msg.GetNextPageToken() != "" {
			t.Error("the last page carries no token")
		}
	})

	t.Run("the size caps at the maximum", func(t *testing.T) {
		src := manyDecks(maxPageSize + 5)
		res, err := list(src, &mtgv1.ListDecksRequest{PageSize: maxPageSize + 50})
		if err != nil {
			t.Fatal(err)
		}
		if got := len(res.Msg.GetDecks()); got != maxPageSize {
			t.Errorf("page = %d decks, want %d", got, maxPageSize)
		}
	})

	t.Run("a token of another filter is an invalid argument", func(t *testing.T) {
		src := manyDecks(defaultPageSize + 1)
		first, err := list(src, &mtgv1.ListDecksRequest{})
		if err != nil {
			t.Fatal(err)
		}
		_, err = list(src, &mtgv1.ListDecksRequest{PageToken: first.Msg.GetNextPageToken(), Query: "elves"})
		if codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("a token of one session does not read another listing", func(t *testing.T) {
		src := manyDecks(defaultPageSize + 1)
		first, err := list(src, &mtgv1.ListDecksRequest{SessionId: "s1"})
		if err != nil {
			t.Fatal(err)
		}
		if _, err := list(src, &mtgv1.ListDecksRequest{PageToken: first.Msg.GetNextPageToken()}); codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
		if _, err := list(src, &mtgv1.ListDecksRequest{PageToken: first.Msg.GetNextPageToken(), SessionId: "s1"}); err != nil {
			t.Errorf("the same session must read its own token: %v", err)
		}
	})

	t.Run("a damaged token is an invalid argument", func(t *testing.T) {
		src := manyDecks(2)
		if _, err := list(src, &mtgv1.ListDecksRequest{PageToken: "not-a-token!"}); codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("an offset past the end gives an empty page", func(t *testing.T) {
		src := manyDecks(defaultPageSize + 1)
		first, _ := list(src, &mtgv1.ListDecksRequest{})
		token := first.Msg.GetNextPageToken()
		// The second page ends the listing, and its own token is empty.
		second, err := list(src, &mtgv1.ListDecksRequest{PageToken: token})
		if err != nil || len(second.Msg.GetDecks()) != 1 {
			t.Fatalf("second page = %v, err %v", ids(second), err)
		}
		src.decks = map[string]*mtgv1.Deck{}
		src.order = nil
		empty, err := list(src, &mtgv1.ListDecksRequest{PageToken: token})
		if err != nil || len(empty.Msg.GetDecks()) != 0 || empty.Msg.GetNextPageToken() != "" {
			t.Errorf("empty page = %v, token %q, err %v", ids(empty), empty.Msg.GetNextPageToken(), err)
		}
	})

	t.Run("the filter reaches the store", func(t *testing.T) {
		src := manyDecks(2)
		yes := true
		if _, err := list(src, &mtgv1.ListDecksRequest{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Favorite: &yes, Query: "  Atraxa  "}); err != nil {
			t.Fatal(err)
		}
		if src.filter.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
			t.Errorf("format = %v", src.filter.Format)
		}
		if src.filter.Favorite == nil || !*src.filter.Favorite {
			t.Errorf("favorite = %v", src.filter.Favorite)
		}
		if src.filter.Query != "Atraxa" {
			t.Errorf("query = %q, want the trimmed text", src.filter.Query)
		}
	})

	t.Run("the power filter reaches the store", func(t *testing.T) {
		src := manyDecks(2)
		if _, err := list(src, &mtgv1.ListDecksRequest{PowerBracket: 3}); err != nil {
			t.Fatal(err)
		}
		if src.filter.PowerBracket != 3 {
			t.Errorf("bracket = %d", src.filter.PowerBracket)
		}
		if _, err := list(src, &mtgv1.ListDecksRequest{PowerSixtyStep: mtgv1.SixtyStep_SIXTY_STEP_FNM}); err != nil {
			t.Fatal(err)
		}
		if src.filter.PowerSixtyStep != mtgv1.SixtyStep_SIXTY_STEP_FNM {
			t.Errorf("step = %v", src.filter.PowerSixtyStep)
		}
	})

	t.Run("a bracket outside 1 to 5 is an invalid argument", func(t *testing.T) {
		src := manyDecks(2)
		if _, err := list(src, &mtgv1.ListDecksRequest{PowerBracket: 6}); codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
		if _, err := list(src, &mtgv1.ListDecksRequest{PowerBracket: -1}); codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("a token of another power filter is an invalid argument", func(t *testing.T) {
		src := manyDecks(defaultPageSize + 1)
		first, err := list(src, &mtgv1.ListDecksRequest{PowerBracket: 3})
		if err != nil {
			t.Fatal(err)
		}
		_, err = list(src, &mtgv1.ListDecksRequest{PageToken: first.Msg.GetNextPageToken(), PowerBracket: 4})
		if codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
}

func TestUpdateDeck(t *testing.T) {
	ctx := context.Background()
	update := func(src *fakeDecks, req *mtgv1.UpdateDeckRequest, opts ...Option) (*connect.Response[mtgv1.UpdateDeckResponse], error) {
		all := append([]Option{WithDecks(src), asUser("u1")}, opts...)
		return newServer(t, nil, all...).UpdateDeck(ctx, connect.NewRequest(req))
	}
	fresh := func() *fakeDecks {
		return &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": {Id: "d1", Name: "a deck"}}}
	}
	name := func(s string) *string { return &s }
	mark := func(b bool) *bool { return &b }

	t.Run("writes the name and the mark", func(t *testing.T) {
		src := fresh()
		res, err := update(src, &mtgv1.UpdateDeckRequest{DeckId: "d1", Name: name("  Elf ball  "), Favorite: mark(true)})
		if err != nil {
			t.Fatal(err)
		}
		if got := res.Msg.GetDeck().GetName(); got != "Elf ball" {
			t.Errorf("name = %q, want the trimmed name", got)
		}
		if !res.Msg.GetDeck().GetFavorite() {
			t.Error("favorite = false")
		}
	})

	t.Run("writes one field alone", func(t *testing.T) {
		src := fresh()
		res, err := update(src, &mtgv1.UpdateDeckRequest{DeckId: "d1", Favorite: mark(true)})
		if err != nil {
			t.Fatal(err)
		}
		if res.Msg.GetDeck().GetName() != "a deck" {
			t.Errorf("name = %q, want it unchanged", res.Msg.GetDeck().GetName())
		}
	})

	t.Run("clears the mark", func(t *testing.T) {
		src := &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": {Id: "d1", Name: "a deck", Favorite: true}}}
		res, err := update(src, &mtgv1.UpdateDeckRequest{DeckId: "d1", Favorite: mark(false)})
		if err != nil {
			t.Fatal(err)
		}
		if res.Msg.GetDeck().GetFavorite() {
			t.Error("favorite = true, want the mark cleared")
		}
	})

	for _, tc := range []struct {
		name string
		req  *mtgv1.UpdateDeckRequest
		want connect.Code
	}{
		{"no id", &mtgv1.UpdateDeckRequest{Name: name("x")}, connect.CodeInvalidArgument},
		{"a bad id", &mtgv1.UpdateDeckRequest{DeckId: "a/b", Name: name("x")}, connect.CodeInvalidArgument},
		{"an empty name", &mtgv1.UpdateDeckRequest{DeckId: "d1", Name: name("   ")}, connect.CodeInvalidArgument},
		{"a long name", &mtgv1.UpdateDeckRequest{DeckId: "d1", Name: name(strings.Repeat("a", maxNameBytes+1))}, connect.CodeInvalidArgument},
		{"nothing to write", &mtgv1.UpdateDeckRequest{DeckId: "d1"}, connect.CodeInvalidArgument},
		{"an unknown deck", &mtgv1.UpdateDeckRequest{DeckId: "d9", Name: name("x")}, connect.CodeNotFound},
	} {
		t.Run(tc.name, func(t *testing.T) {
			src := fresh()
			if _, err := update(src, tc.req); codeOf(t, err) != tc.want {
				t.Errorf("code = %v, want %v", codeOf(t, err), tc.want)
			}
			if tc.want == connect.CodeInvalidArgument && src.updates != 0 {
				t.Errorf("the store saw %d writes, want none", src.updates)
			}
		})
	}

	t.Run("unimplemented without a store", func(t *testing.T) {
		s := newServer(t, nil, asUser("u1"))
		_, err := s.UpdateDeck(ctx, connect.NewRequest(&mtgv1.UpdateDeckRequest{DeckId: "d1", Name: name("x")}))
		if codeOf(t, err) != connect.CodeUnimplemented {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("unauthenticated without a user", func(t *testing.T) {
		s := newServer(t, nil, WithDecks(fresh()))
		_, err := s.UpdateDeck(ctx, connect.NewRequest(&mtgv1.UpdateDeckRequest{DeckId: "d1", Name: name("x")}))
		if codeOf(t, err) != connect.CodeUnauthenticated {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("a store failure is internal", func(t *testing.T) {
		bad := &fakeDecks{err: errors.New("firestore down")}
		if _, err := update(bad, &mtgv1.UpdateDeckRequest{DeckId: "d1", Name: name("x")}); codeOf(t, err) != connect.CodeInternal {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
}

func TestDeleteDeck(t *testing.T) {
	ctx := context.Background()
	del := func(src *fakeDecks, id string, opts ...Option) error {
		all := append([]Option{WithDecks(src), asUser("u1")}, opts...)
		_, err := newServer(t, nil, all...).DeleteDeck(ctx, connect.NewRequest(&mtgv1.DeleteDeckRequest{DeckId: id}))
		return err
	}
	fresh := func() *fakeDecks {
		return &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": {Id: "d1"}}}
	}

	t.Run("removes the deck, and a second delete is not found", func(t *testing.T) {
		src := fresh()
		if err := del(src, "d1"); err != nil {
			t.Fatal(err)
		}
		if len(src.deleted) != 1 || src.deleted[0] != "d1" {
			t.Errorf("deleted = %v", src.deleted)
		}
		if err := del(src, "d1"); codeOf(t, err) != connect.CodeNotFound {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("an unknown deck is not found", func(t *testing.T) {
		if err := del(fresh(), "d9"); codeOf(t, err) != connect.CodeNotFound {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("a bad id is an invalid argument", func(t *testing.T) {
		if err := del(fresh(), "a/b"); codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("no id is an invalid argument", func(t *testing.T) {
		if err := del(fresh(), "  "); codeOf(t, err) != connect.CodeInvalidArgument {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("unimplemented without a store", func(t *testing.T) {
		s := newServer(t, nil, asUser("u1"))
		_, err := s.DeleteDeck(ctx, connect.NewRequest(&mtgv1.DeleteDeckRequest{DeckId: "d1"}))
		if codeOf(t, err) != connect.CodeUnimplemented {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})

	t.Run("the chat and every deck of the chat go with it (D-456)", func(t *testing.T) {
		src := &fakeDecks{decks: map[string]*mtgv1.Deck{
			"d1": {Id: "d1", SessionId: "s1"}, "d2": {Id: "d2", SessionId: "s1", RevisedFromDeckId: "d1"}, "d9": {Id: "d9", SessionId: "s9"},
		}}
		chats := &fakeSessions{sessions: map[string]*mtgv1.Session{"s1": {Id: "s1", DeckIds: []string{"d1", "d2", "gone"}}, "s9": {Id: "s9"}}}
		if err := del(src, "d1", WithSessions(chats)); err != nil {
			t.Fatal(err)
		}
		sort.Strings(src.deleted)
		if got := strings.Join(src.deleted, ","); got != "d1,d2" {
			t.Errorf("deleted = %q, want d1 and d2 and not d9", got)
		}
		if _, ok := chats.sessions["s1"]; ok {
			t.Error("the chat is still stored")
		}
		if _, ok := chats.sessions["s9"]; !ok {
			t.Error("another chat went")
		}
	})

	t.Run("a deck whose chat is already gone still goes", func(t *testing.T) {
		src := &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": {Id: "d1", SessionId: "s-gone"}}}
		if err := del(src, "d1", WithSessions(&fakeSessions{sessions: map[string]*mtgv1.Session{}})); err != nil {
			t.Fatal(err)
		}
		if len(src.deleted) != 1 {
			t.Errorf("deleted = %v", src.deleted)
		}
	})

	t.Run("a store failure is internal", func(t *testing.T) {
		bad := &fakeDecks{err: errors.New("firestore down")}
		if err := del(bad, "d1"); codeOf(t, err) != connect.CodeInternal {
			t.Errorf("code = %v", codeOf(t, err))
		}
	})
}

// fakeSessions is the chat store of the delete test.
type fakeSessions struct {
	sessions map[string]*mtgv1.Session
}

func (f *fakeSessions) Get(_ context.Context, _, id string) (*mtgv1.Session, error) {
	s, ok := f.sessions[id]
	if !ok {
		return nil, sessions.ErrNotFound
	}
	return s, nil
}

func (f *fakeSessions) Delete(_ context.Context, _, id string) error {
	if _, ok := f.sessions[id]; !ok {
		return sessions.ErrNotFound
	}
	delete(f.sessions, id)
	return nil
}
