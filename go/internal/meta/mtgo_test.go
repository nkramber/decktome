package meta

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func fixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestMTGOMonthURL(t *testing.T) {
	if got := MTGOMonthURL(2026, time.September); got != "https://www.mtgo.com/decklists/2026/09" {
		t.Fatalf("month url = %s", got)
	}
}

func TestParseMTGOMonth(t *testing.T) {
	slugs := ParseMTGOMonth(fixture(t, "mtgo_month.html"))
	if len(slugs) != 6 {
		t.Fatalf("slugs = %d, want 6 (the repeated link counts once): %v", len(slugs), slugs)
	}
	if slugs[0] != "standard-challenge-16-2026-09-0212853229" {
		t.Errorf("first slug = %s", slugs[0])
	}
}

func TestMTGOSlugFormat(t *testing.T) {
	tests := map[string]string{
		"modern-challenge-32-2026-09-0212853228": FormatModern,
		"standard-league-2026-09-0211015":        FormatStandard,
		"duel-commander-league-2026-09-0210931":  "",
		"legacy-challenge-32-2026-09-0112853222": "",
		"modern-super-qualifier-2026-08-3012800": FormatModern,
	}
	for slug, want := range tests {
		if got := MTGOSlugFormat(slug); got != want {
			t.Errorf("%s: format = %q, want %q", slug, got, want)
		}
	}
}

// TestParseMTGOEventChallenge reads a challenge page: the standings give
// the placement, and a top-8 finish is great (D-414).
func TestParseMTGOEventChallenge(t *testing.T) {
	slug := "modern-challenge-32-2026-09-0212853228"
	lists, err := ParseMTGOEvent(slug, fixture(t, "mtgo_challenge.html"))
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 3 {
		t.Fatalf("lists = %d, want 3", len(lists))
	}
	byID := map[string]List{}
	for _, l := range lists {
		byID[l.ID] = l
		if l.Format != FormatModern || l.Date != "2026-09-02" || l.Source != SourceMTGO {
			t.Errorf("%s: format %s, date %s, source %s", l.ID, l.Format, l.Date, l.Source)
		}
		if l.Players != 83 || l.Event != "Modern Challenge 32" {
			t.Errorf("%s: players %d, event %q", l.ID, l.Players, l.Event)
		}
		if l.Size() < 60 {
			t.Errorf("%s: size %d", l.ID, l.Size())
		}
	}
	winner := byID[slug+"/3227644"]
	if winner.Placement != 1 || winner.Tier != TierGreat {
		t.Errorf("winner: placement %d, tier %s", winner.Placement, winner.Tier)
	}
	if winner.Wins == 0 {
		t.Errorf("winner carries no record")
	}
	rest := byID[slug+"/30930"]
	if rest.Placement != 27 || rest.Tier != TierGood {
		t.Errorf("27th: placement %d, tier %s", rest.Placement, rest.Tier)
	}
	if len(rest.Sideboard) == 0 {
		t.Errorf("27th carries no sideboard")
	}
}

// TestParseMTGOEventLeague reads a league page: no standings, a record
// per list, and every list is a finish, so good.
func TestParseMTGOEventLeague(t *testing.T) {
	slug := "modern-league-2026-09-0210983"
	lists, err := ParseMTGOEvent(slug, fixture(t, "mtgo_league.html"))
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 2 {
		t.Fatalf("lists = %d, want 2", len(lists))
	}
	for _, l := range lists {
		if l.Tier != TierGood || l.Placement != 0 || l.Wins != 5 || l.Losses != 0 {
			t.Errorf("%s: tier %s, placement %d, record %d-%d", l.ID, l.Tier, l.Placement, l.Wins, l.Losses)
		}
		if l.Event != slug {
			t.Errorf("%s: event %q, want the slug when the page names none", l.ID, l.Event)
		}
	}
}

func TestParseMTGOEventSkipsOtherFormats(t *testing.T) {
	lists, err := ParseMTGOEvent("legacy-challenge-32-2026-09-0112853222", fixture(t, "mtgo_challenge.html"))
	if err != nil || lists != nil {
		t.Fatalf("legacy: lists %v, err %v, want none", lists, err)
	}
}

func TestParseMTGOEventNoObject(t *testing.T) {
	if _, err := ParseMTGOEvent("modern-challenge-32-2026-09-0212853228", []byte("<html></html>")); err == nil {
		t.Fatal("a page with no decklist object must be an error, because M-6 counts it")
	}
}
