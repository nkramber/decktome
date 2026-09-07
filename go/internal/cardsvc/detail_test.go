package cardsvc

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

func detailIndex(rulings map[string][]cards.Ruling) *cards.Index {
	list := []*mtgv1.Card{{OracleId: "o-sol", Name: "Sol Ring", DefaultPrinting: &mtgv1.Printing{ScryfallId: "p-c21"}}}
	printings := []cards.Printing{
		{ScryfallID: "p-c21", OracleID: "o-sol", SetCode: "c21", CollectorNumber: "263", PriceUSD: 1.5},
		{ScryfallID: "p-msc-10", OracleID: "o-sol", SetCode: "msc", CollectorNumber: "10", PriceUSD: 2},
		{ScryfallID: "p-msc-9", OracleID: "o-sol", SetCode: "msc", CollectorNumber: "9", PriceUSD: 2.5},
		{ScryfallID: "p-mtga", OracleID: "o-sol", SetCode: "mtga", CollectorNumber: "1", Digital: true},
	}
	sets := []cards.SetInfo{
		{Code: "c21", Name: "Commander 2021", ReleasedAt: "2021-04-23"},
		{Code: "msc", Name: "Marvel Super Heroes Commander", ReleasedAt: "2026-06-26"},
	}
	opts := []cards.IndexOption{cards.WithSets(sets)}
	if rulings != nil {
		opts = append(opts, cards.WithRulings(rulings))
	}
	return cards.NewIndex(list, printings, nil, time.Date(2026, 9, 3, 9, 2, 27, 0, time.UTC), opts...)
}

// TestGetRulingsCarriesTheSnapshotDate is the PR-20 gate line: the
// rulings show their dates and the snapshot date.
func TestGetRulingsCarriesTheSnapshotDate(t *testing.T) {
	s := New()
	ctx := context.Background()
	s.Swap(detailIndex(map[string][]cards.Ruling{"o-sol": {
		{PublishedAt: "2004-10-04", Comment: "Sol Ring is an artifact.", Source: "wotc"},
		{PublishedAt: "2020-11-10", Comment: "Still an artifact.", Source: "scryfall"},
	}}))
	res, err := s.GetRulings(ctx, connect.NewRequest(&mtgv1.GetRulingsRequest{OracleId: "o-sol"}))
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.AsOf != "2026-09-03" || !res.Msg.HasRulings || len(res.Msg.Rulings) != 2 {
		t.Fatalf("response = %+v", res.Msg)
	}
	if res.Msg.Rulings[0].PublishedAt != "2004-10-04" || res.Msg.Rulings[1].Source != "scryfall" {
		t.Errorf("rulings = %+v", res.Msg.Rulings)
	}
	if _, err := s.GetRulings(ctx, connect.NewRequest(&mtgv1.GetRulingsRequest{OracleId: "o-none"})); code(err) != connect.CodeNotFound {
		t.Errorf("unknown card: %v, want NotFound", err)
	}
	// A snapshot with no rulings file says so.
	s.Swap(detailIndex(nil))
	res, err = s.GetRulings(ctx, connect.NewRequest(&mtgv1.GetRulingsRequest{OracleId: "o-sol"}))
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.HasRulings || len(res.Msg.Rulings) != 0 {
		t.Errorf("no rulings file: %+v", res.Msg)
	}
}

// TestGetPrintingsOrdersNewestSetFirst: every playable printing with its
// price, the newest set first, then the collector number as a number.
func TestGetPrintingsOrdersNewestSetFirst(t *testing.T) {
	s := New()
	ctx := context.Background()
	if _, err := s.GetPrintings(ctx, connect.NewRequest(&mtgv1.GetPrintingsRequest{OracleId: "o-sol"})); code(err) != connect.CodeUnavailable {
		t.Errorf("before a swap: %v, want Unavailable", err)
	}
	s.Swap(detailIndex(nil))
	res, err := s.GetPrintings(ctx, connect.NewRequest(&mtgv1.GetPrintingsRequest{OracleId: "o-sol"}))
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, p := range res.Msg.Printings {
		got = append(got, p.ScryfallId)
	}
	want := []string{"p-msc-9", "p-msc-10", "p-c21", "p-mtga"}
	if len(got) != len(want) {
		t.Fatalf("printings = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("printings = %v, want %v", got, want)
			break
		}
	}
	if res.Msg.PriceAsOf != "2026-09-03" || res.Msg.Printings[0].PriceUsd != 2.5 || !res.Msg.Printings[3].Digital {
		t.Errorf("response = %+v", res.Msg)
	}
	res.Msg.Printings[0].PriceUsd = 0
	if again, _ := s.GetPrintings(ctx, connect.NewRequest(&mtgv1.GetPrintingsRequest{OracleId: "o-sol"})); again.Msg.Printings[0].PriceUsd != 2.5 {
		t.Error("the answer must carry copies of the index rows")
	}
	if _, err := s.GetPrintings(ctx, connect.NewRequest(&mtgv1.GetPrintingsRequest{OracleId: "o-none"})); code(err) != connect.CodeNotFound {
		t.Errorf("unknown card: %v, want NotFound", err)
	}
}
