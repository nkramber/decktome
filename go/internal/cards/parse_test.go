package cards

import (
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

func TestParseCardOracleIDAndArtist(t *testing.T) {
	t.Run("reversible card takes oracle_id from the first face", func(t *testing.T) {
		line := `{"id":"p1","name":"A // B","layout":"reversible_card","artist":"Top",
		  "card_faces":[{"oracle_id":"o1","name":"A","artist":"Face A"},{"oracle_id":"o1","name":"B"}]}`
		c, err := parseCard([]byte(line))
		if err != nil {
			t.Fatal(err)
		}
		if c.OracleId != "o1" {
			t.Errorf("oracle_id = %q, want o1", c.OracleId)
		}
		if c.Faces[0].Artist != "Face A" || c.Faces[1].Artist != "Top" {
			t.Errorf("face artists = %q, %q", c.Faces[0].Artist, c.Faces[1].Artist)
		}
	})
	t.Run("single face takes the card artist", func(t *testing.T) {
		c, err := parseCard([]byte(`{"id":"p2","oracle_id":"o2","name":"Solo","layout":"normal","artist":"Top"}`))
		if err != nil {
			t.Fatal(err)
		}
		if len(c.Faces) != 1 || c.Faces[0].Artist != "Top" {
			t.Errorf("faces = %v", c.Faces)
		}
	})
}

func TestLoadCardsStatsSkipsNoOracleID(t *testing.T) {
	in := strings.Join([]string{
		`{"id":"p1","oracle_id":"o1","name":"Good","layout":"normal"}`,
		`{"id":"p2","name":"Orphan","layout":"normal"}`,
		`{"id":"p3","name":"Token","layout":"token"}`,
	}, "\n") + "\n"
	cardList, stats, err := LoadCardsStats(strings.NewReader(in), "x.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	if len(cardList) != 1 || stats.NoOracleID != 1 || stats.Skipped != 1 {
		t.Fatalf("cards=%d stats=%+v", len(cardList), stats)
	}
}

func TestIndexCollisionsAndPriceAsOf(t *testing.T) {
	asOf := time.Date(2026, 8, 24, 9, 0, 0, 0, time.UTC)
	cardList := []*mtgv1.Card{
		{OracleId: "o1", Name: "Same Name", PriceUsd: 1.5, Faces: []*mtgv1.CardFace{{Name: "Same Name"}}},
		{OracleId: "o2", Name: "Same Name", Faces: []*mtgv1.CardFace{{Name: "Same Name"}}},
		{OracleId: "o3", Name: "Split // Card", Faces: []*mtgv1.CardFace{{Name: "Split"}, {Name: "Card"}}},
		{OracleId: "o4", Name: "Other // Split", Faces: []*mtgv1.CardFace{{Name: "Other"}, {Name: "Split"}}},
	}
	idx := NewIndex(cardList, nil, nil, asOf)
	col := idx.Collisions()
	// o2 repeats the full name. o2's face repeats a name held by o1,
	// and o4's second face repeats a face name held by o3.
	if col.FullNames != 1 || col.FaceNames != 2 {
		t.Errorf("collisions = %+v, want {1 2}", col)
	}
	if c, _ := idx.ByName("Same Name"); c.OracleId != "o1" {
		t.Errorf("first card must keep the full name, got %q", c.OracleId)
	}
	if cardList[0].PriceAsOf != "2026-08-24" {
		t.Errorf("price_as_of = %q, want 2026-08-24", cardList[0].PriceAsOf)
	}
	if cardList[1].PriceAsOf != "" {
		t.Errorf("card with no price got price_as_of %q", cardList[1].PriceAsOf)
	}
}

func TestSearchOrderIsRankSorted(t *testing.T) {
	cardList := []*mtgv1.Card{
		{OracleId: "o1", Name: "Unranked"},
		{OracleId: "o2", Name: "Rank 30", EdhrecRank: 30},
		{OracleId: "o3", Name: "Rank 2", EdhrecRank: 2},
	}
	idx := NewIndex(cardList, nil, nil, time.Now())
	res, total := idx.Search(SearchQuery{})
	if total != 3 || res[0].Name != "Rank 2" || res[1].Name != "Rank 30" || res[2].Name != "Unranked" {
		t.Errorf("order = %v", res)
	}
	// The caller's slice keeps its order.
	if cardList[0].Name != "Unranked" {
		t.Error("NewIndex reordered the caller's slice")
	}
}
