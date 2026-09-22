package cardsvc

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

func testIndex(n int) *cards.Index {
	list := make([]*mtgv1.Card, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, &mtgv1.Card{
			OracleId:        fmt.Sprintf("o%d", i),
			Name:            fmt.Sprintf("Card %d", i),
			EdhrecRank:      int32(i + 1),
			DefaultPrinting: &mtgv1.Printing{ScryfallId: fmt.Sprintf("p%d", i)},
		})
	}
	return cards.NewIndex(list, nil, nil, time.Now())
}

func code(err error) connect.Code {
	var ce *connect.Error
	if errors.As(err, &ce) {
		return ce.Code()
	}
	return connect.CodeUnknown
}

func TestUnavailableBeforeSwap(t *testing.T) {
	s := New()
	s.SetIndexWait(0)
	ctx := context.Background()
	if _, err := s.Lookup(ctx, connect.NewRequest(&mtgv1.LookupRequest{Key: &mtgv1.LookupRequest_Name{Name: "x"}})); code(err) != connect.CodeUnavailable {
		t.Errorf("Lookup before swap: %v", err)
	}
	if _, err := s.Search(ctx, connect.NewRequest(&mtgv1.SearchRequest{})); code(err) != connect.CodeUnavailable {
		t.Errorf("Search before swap: %v", err)
	}
	if s.Current() != nil {
		t.Error("Current must be nil before swap")
	}
}

func TestLookup(t *testing.T) {
	s := New()
	s.Swap(testIndex(3))
	ctx := context.Background()
	tests := []struct {
		name     string
		req      *mtgv1.LookupRequest
		wantCode connect.Code
		wantName string
	}{
		{"by name", &mtgv1.LookupRequest{Key: &mtgv1.LookupRequest_Name{Name: "card 1"}}, 0, "Card 1"},
		{"by printing", &mtgv1.LookupRequest{Key: &mtgv1.LookupRequest_ScryfallId{ScryfallId: "p2"}}, 0, "Card 2"},
		{"by oracle", &mtgv1.LookupRequest{Key: &mtgv1.LookupRequest_OracleId{OracleId: "o0"}}, 0, "Card 0"},
		{"missing", &mtgv1.LookupRequest{Key: &mtgv1.LookupRequest_Name{Name: "nope"}}, connect.CodeNotFound, ""},
		{"no key", &mtgv1.LookupRequest{}, connect.CodeInvalidArgument, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := s.Lookup(ctx, connect.NewRequest(tt.req))
			if tt.wantCode != 0 {
				if code(err) != tt.wantCode {
					t.Fatalf("code = %v, want %v", code(err), tt.wantCode)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if res.Msg.Card.Name != tt.wantName {
				t.Errorf("name = %q, want %q", res.Msg.Card.Name, tt.wantName)
			}
		})
	}
}

func TestSearchPaging(t *testing.T) {
	s := New()
	s.Swap(testIndex(450))
	ctx := context.Background()
	t.Run("default page size", func(t *testing.T) {
		res, err := s.Search(ctx, connect.NewRequest(&mtgv1.SearchRequest{}))
		if err != nil {
			t.Fatal(err)
		}
		if len(res.Msg.Cards) != defaultPageSize || res.Msg.NextPageToken != "50" {
			t.Errorf("cards=%d next=%q", len(res.Msg.Cards), res.Msg.NextPageToken)
		}
	})
	t.Run("size above max clamps to max", func(t *testing.T) {
		res, err := s.Search(ctx, connect.NewRequest(&mtgv1.SearchRequest{PageSize: 1000}))
		if err != nil {
			t.Fatal(err)
		}
		if len(res.Msg.Cards) != maxPageSize || res.Msg.NextPageToken != "200" {
			t.Errorf("cards=%d next=%q", len(res.Msg.Cards), res.Msg.NextPageToken)
		}
	})
	t.Run("page token walks to the end", func(t *testing.T) {
		token, seen := "", 0
		for {
			res, err := s.Search(ctx, connect.NewRequest(&mtgv1.SearchRequest{PageSize: 200, PageToken: token}))
			if err != nil {
				t.Fatal(err)
			}
			seen += len(res.Msg.Cards)
			token = res.Msg.NextPageToken
			if token == "" {
				break
			}
		}
		if seen != 450 {
			t.Errorf("walked %d cards, want 450", seen)
		}
	})
	t.Run("bad token", func(t *testing.T) {
		_, err := s.Search(ctx, connect.NewRequest(&mtgv1.SearchRequest{PageToken: "x"}))
		if code(err) != connect.CodeInvalidArgument {
			t.Errorf("err = %v", err)
		}
		_, err = s.Search(ctx, connect.NewRequest(&mtgv1.SearchRequest{PageToken: "-1"}))
		if code(err) != connect.CodeInvalidArgument {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("offset past the end", func(t *testing.T) {
		res, err := s.Search(ctx, connect.NewRequest(&mtgv1.SearchRequest{PageToken: "9999"}))
		if err != nil || len(res.Msg.Cards) != 0 || res.Msg.NextPageToken != "" {
			t.Errorf("res=%v err=%v", res, err)
		}
	})
}

func TestGetCards(t *testing.T) {
	s := New()
	s.SetIndexWait(0)
	ctx := context.Background()
	if _, err := s.GetCards(ctx, connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: []string{"o0"}})); code(err) != connect.CodeUnavailable {
		t.Errorf("GetCards before swap: %v", err)
	}
	s.Swap(testIndex(130))
	t.Run("request order, dedup, and the missing list", func(t *testing.T) {
		res, err := s.GetCards(ctx, connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: []string{"o2", "nope", "o0", "o2", ""}}))
		if err != nil {
			t.Fatal(err)
		}
		if len(res.Msg.Cards) != 2 || res.Msg.Cards[0].Name != "Card 2" || res.Msg.Cards[1].Name != "Card 0" {
			t.Errorf("cards = %v", res.Msg.Cards)
		}
		if len(res.Msg.MissingOracleIds) != 1 || res.Msg.MissingOracleIds[0] != "nope" {
			t.Errorf("missing = %v", res.Msg.MissingOracleIds)
		}
	})
	t.Run("empty request", func(t *testing.T) {
		res, err := s.GetCards(ctx, connect.NewRequest(&mtgv1.GetCardsRequest{}))
		if err != nil || len(res.Msg.Cards) != 0 {
			t.Errorf("res=%v err=%v", res, err)
		}
	})
	t.Run("120 distinct ids pass, 121 refuse", func(t *testing.T) {
		ids := make([]string, 0, 121)
		for i := 0; i < 121; i++ {
			ids = append(ids, fmt.Sprintf("o%d", i))
		}
		res, err := s.GetCards(ctx, connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: ids[:120]}))
		if err != nil || len(res.Msg.Cards) != 120 {
			t.Errorf("120: len=%d err=%v", len(res.Msg.GetCards()), err)
		}
		if _, err := s.GetCards(ctx, connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: ids})); code(err) != connect.CodeInvalidArgument {
			t.Errorf("121: err=%v", err)
		}
		// Duplicates do not count against the cap.
		dup := append(append([]string{}, ids[:120]...), "o0", "o1")
		if _, err := s.GetCards(ctx, connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: dup})); err != nil {
			t.Errorf("120 plus repeats: %v", err)
		}
	})
}

// TestReadyWaitsForFirstIndex: the API listens before the snapshot loads,
// so a request of the first seconds after a cold start meets no index
// (F-164). Such a request waits for the first Swap, and it reads the
// cards. A request that ends first reads Unavailable and waits no more.
func TestReadyWaitsForFirstIndex(t *testing.T) {
	s := New()
	s.SetIndexWait(5 * time.Second)
	go func() {
		time.Sleep(20 * time.Millisecond)
		s.Swap(testIndex(1))
	}()
	start := time.Now()
	res, err := s.GetCards(context.Background(), connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: []string{"o0"}}))
	if err != nil {
		t.Fatalf("GetCards during the load: %v", err)
	}
	if len(res.Msg.Cards) != 1 {
		t.Errorf("cards: %d, want 1", len(res.Msg.Cards))
	}
	if time.Since(start) < 20*time.Millisecond {
		t.Error("GetCards answered before the swap")
	}
}

// TestReadyStopsOnCanceledRequest: a reader who leaves the page ends the
// request, and the wait of F-164 ends with it.
func TestReadyStopsOnCanceledRequest(t *testing.T) {
	s := New()
	s.SetIndexWait(time.Minute)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	start := time.Now()
	if _, err := s.GetCards(ctx, connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: []string{"o0"}})); code(err) != connect.CodeUnavailable {
		t.Errorf("GetCards on a canceled request: %v, want Unavailable", err)
	}
	if time.Since(start) > 5*time.Second {
		t.Error("GetCards waited past the end of the request")
	}
}
