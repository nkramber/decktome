package cardsvc

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestGetCardsQualityRows: the rows ride on a copy, the index card stays
// bare, and a nil source removes them (PR-14B).
func TestGetCardsQualityRows(t *testing.T) {
	s := New()
	idx := testIndex(3)
	s.Swap(idx)
	ctx := context.Background()
	req := connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: []string{"o1", "o2"}})
	res, err := s.GetCards(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Msg.GetCards()) != 2 || res.Msg.GetCards()[0].GetQuality() != nil {
		t.Fatalf("before a source: %+v", res.Msg.GetCards())
	}
	s.SetQuality(func(id string) []*mtgv1.CardQuality {
		if id != "o1" {
			return nil
		}
		return []*mtgv1.CardQuality{{Format: mtgv1.FormatId_FORMAT_ID_MODERN, Inclusion: 0.25}}
	})
	res, err = s.GetCards(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	got := res.Msg.GetCards()
	if len(got[0].GetQuality()) != 1 || got[0].GetQuality()[0].GetInclusion() != 0.25 || got[1].GetQuality() != nil {
		t.Errorf("with a source: %+v", got)
	}
	if orig, _ := idx.ByOracleID("o1"); orig.GetQuality() != nil {
		t.Errorf("the index card carries the rows")
	}
	s.SetQuality(nil)
	res, err = s.GetCards(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.GetCards()[0].GetQuality() != nil {
		t.Errorf("after the source left: %+v", res.Msg.GetCards()[0])
	}
}

// TestGetCardsStopsAtTheCap is REV-085 of the review of 2026-09-24. The
// handler sized its map and its list from the whole request before the
// cap check, so a request of a million ids allocated for a million.
func TestGetCardsStopsAtTheCap(t *testing.T) {
	s := New()
	s.Swap(testIndex(3))
	ids := make([]string, 1_000_000)
	for i := range ids {
		ids[i] = fmt.Sprintf("o-%d", i)
	}
	req := connect.NewRequest(&mtgv1.GetCardsRequest{OracleIds: ids})
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	_, err := s.GetCards(context.Background(), req)
	runtime.ReadMemStats(&after)
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Fatalf("err = %v, want InvalidArgument", err)
	}
	if grew := after.TotalAlloc - before.TotalAlloc; grew > 1<<20 {
		t.Errorf("the refused call allocated %d bytes, want under 1 MiB", grew)
	}
}
