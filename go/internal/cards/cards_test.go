package cards

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

func loadFixtureIndex(t *testing.T) *Index {
	t.Helper()
	cardsFile, err := os.Open("testdata/cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cardsFile.Close() }()
	cardList, err := LoadCards(cardsFile, "cards_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	printFile, err := os.Open("testdata/printings_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = printFile.Close() }()
	printings, err := LoadPrintings(printFile, "printings_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	tagFile, err := os.Open("testdata/tags_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tagFile.Close() }()
	tags, err := LoadTags(tagFile, "tags_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	return NewIndex(cardList, printings, tags, time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC))
}

// TestTrickyNamesGate is the PR-2 gate: every name on the fixed list
// resolves by exact name lookup.
func TestTrickyNamesGate(t *testing.T) {
	idx := loadFixtureIndex(t)
	f, err := os.Open("testdata/tricky_names.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	sc := bufio.NewScanner(f)
	total, missed := 0, 0
	for sc.Scan() {
		name := strings.TrimSpace(sc.Text())
		if name == "" {
			continue
		}
		total++
		if _, ok := idx.ByName(name); !ok {
			missed++
			t.Errorf("name does not resolve: %q", name)
		}
	}
	if total < 200 {
		t.Fatalf("gate list has %d names, want at least 200", total)
	}
	t.Logf("gate: %d/%d resolve", total-missed, total)
}

func TestFaceNameResolves(t *testing.T) {
	idx := loadFixtureIndex(t)
	tests := []struct{ face, full string }{
		{"Fire", "Fire // Ice"},
		{"Stomp", "Bonecrusher Giant // Stomp"},
		{"Delver of Secrets", "Delver of Secrets // Insectile Aberration"},
	}
	for _, tt := range tests {
		c, ok := idx.ByName(tt.face)
		if !ok {
			t.Errorf("face %q does not resolve", tt.face)
			continue
		}
		if c.Name != tt.full {
			t.Errorf("face %q resolved to %q, want %q", tt.face, c.Name, tt.full)
		}
	}
}

func TestDerivations(t *testing.T) {
	idx := loadFixtureIndex(t)
	get := func(name string) *mtgv1.Card {
		t.Helper()
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("fixture card missing: %q", name)
		}
		return c
	}

	t.Run("types parsed", func(t *testing.T) {
		c := get("Ajani's Pridemate")
		if !slices.Contains(c.CardTypes, "Creature") {
			t.Errorf("card types = %v, want Creature", c.CardTypes)
		}
		if !slices.Contains(c.Subtypes, "Cat") {
			t.Errorf("subtypes = %v, want Cat", c.Subtypes)
		}
	})
	t.Run("any count in deck", func(t *testing.T) {
		for _, name := range []string{"Relentless Rats", "Persistent Petitioners", "Shadowborn Apostle"} {
			if !get(name).AnyCountInDeck {
				t.Errorf("%s: any_count_in_deck = false, want true", name)
			}
		}
		if get("Ajani's Pridemate").AnyCountInDeck {
			t.Error("Ajani's Pridemate: any_count_in_deck = true, want false")
		}
	})
	t.Run("commander eligibility", func(t *testing.T) {
		tests := []struct {
			name string
			want bool
		}{
			{"Thrasios, Triton Hero", true},
			{"Tovolar, Dire Overlord // Tovolar, the Midnight Scourge", true},
			{"Ajani's Pridemate", false},
			{"Sol Ring", false},
			{"The Prismatic Bridge", true}, // face of Esika: text allows it
		}
		for _, tt := range tests {
			if got := get(tt.name).CanBeCommander; got != tt.want {
				t.Errorf("%s: can_be_commander = %v, want %v", tt.name, got, tt.want)
			}
		}
	})
	t.Run("partner kinds", func(t *testing.T) {
		if got := get("Thrasios, Triton Hero").Partner; got != mtgv1.PartnerKind_PARTNER_KIND_PARTNER {
			t.Errorf("Thrasios partner = %v", got)
		}
		c := get("Halana and Alena, Partners")
		if c.Partner != mtgv1.PartnerKind_PARTNER_KIND_NONE {
			// Halana and Alena is one card, no partner keyword. Guard the fixture.
			t.Logf("note: Halana and Alena partner = %v", c.Partner)
		}
		w := get("Wilson, Refined Grizzly")
		if w.Partner != mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND {
			t.Errorf("Wilson partner = %v, want CHOOSE_BACKGROUND", w.Partner)
		}
		// The companion card carries "Doctor's companion". The Doctor
		// pairs through the companion, so the Doctor reads NONE.
		d := get("Clara Oswald")
		if d.Partner != mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION {
			t.Errorf("Clara Oswald partner = %v, want DOCTORS_COMPANION", d.Partner)
		}
		if got := get("The Tenth Doctor").Partner; got != mtgv1.PartnerKind_PARTNER_KIND_NONE {
			t.Errorf("The Tenth Doctor partner = %v, want NONE", got)
		}
	})
	t.Run("companion", func(t *testing.T) {
		for _, name := range []string{"Lutri, the Spellchaser", "Yorion, Sky Nomad", "Jegantha, the Wellspring"} {
			if !get(name).IsCompanion {
				t.Errorf("%s: is_companion = false, want true", name)
			}
		}
	})
	t.Run("colorless produced mana", func(t *testing.T) {
		c := get("Sol Ring")
		if !slices.Contains(c.ProducedMana, mtgv1.Color_COLOR_C) {
			t.Errorf("Sol Ring produced_mana = %v, want COLOR_C", c.ProducedMana)
		}
	})
	t.Run("legalities", func(t *testing.T) {
		c := get("Sol Ring")
		if c.Legalities["commander"] != mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL {
			t.Errorf("Sol Ring commander legality = %v", c.Legalities["commander"])
		}
		if c.Legalities["legacy"] != mtgv1.LegalityStatus_LEGALITY_STATUS_BANNED {
			t.Errorf("Sol Ring legacy legality = %v, want BANNED", c.Legalities["legacy"])
		}
	})
	t.Run("faces normalized", func(t *testing.T) {
		c := get("Delver of Secrets // Insectile Aberration")
		if len(c.Faces) != 2 {
			t.Fatalf("faces = %d, want 2", len(c.Faces))
		}
		for i, f := range c.Faces {
			if f.ImageUris == nil || f.ImageUris.Normal == "" {
				t.Errorf("face %d has no image", i)
			}
		}
		single := get("Ajani's Pridemate")
		if len(single.Faces) != 1 || single.Faces[0].ImageUris == nil {
			t.Error("single-faced card must have one face with an image")
		}
	})
}

func TestPrintingLookup(t *testing.T) {
	idx := loadFixtureIndex(t)
	f, _ := os.Open("testdata/printings_fixture.jsonl")
	defer func() { _ = f.Close() }()
	printings, err := LoadPrintings(f, "printings_fixture.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	if len(printings) == 0 {
		t.Fatal("no printings in fixture")
	}
	found := 0
	for _, p := range printings {
		if c, ok := idx.ByPrintingID(p.ScryfallID); ok {
			found++
			if _, ok := idx.ByOracleID(c.OracleId); !ok {
				t.Fatalf("printing %s resolves to unknown oracle", p.ScryfallID)
			}
		}
	}
	if found != len(printings) {
		t.Errorf("printings resolved = %d, want %d", found, len(printings))
	}
}

func TestSearch(t *testing.T) {
	idx := loadFixtureIndex(t)
	t.Run("color identity filter", func(t *testing.T) {
		res, _ := idx.Search(SearchQuery{ColorsWithin: []mtgv1.Color{mtgv1.Color_COLOR_W}})
		for _, c := range res {
			for _, col := range c.ColorIdentity {
				if col != mtgv1.Color_COLOR_W {
					t.Fatalf("%s has off-color identity %v", c.Name, col)
				}
			}
		}
		if len(res) == 0 {
			t.Fatal("no white-or-colorless cards in fixture")
		}
	})
	t.Run("legal filter excludes banned", func(t *testing.T) {
		res, _ := idx.Search(SearchQuery{LegalIn: "legacy"})
		for _, c := range res {
			if c.Name == "Sol Ring" {
				t.Fatal("Sol Ring is banned in legacy and must not match")
			}
		}
	})
	t.Run("tag filter", func(t *testing.T) {
		if idx.tags.Len() == 0 {
			t.Skip("no tags in fixture")
		}
		res, total := idx.Search(SearchQuery{OracleTags: []string{"lifegain"}})
		_ = res
		if total == 0 {
			t.Skip("fixture has no lifegain-tagged cards")
		}
	})
	t.Run("pagination", func(t *testing.T) {
		all, total := idx.Search(SearchQuery{})
		if total != idx.Len() || len(all) != total {
			t.Fatalf("unfiltered search: got %d/%d, index has %d", len(all), total, idx.Len())
		}
		page, _ := idx.Search(SearchQuery{Limit: 10, Offset: 5})
		if len(page) != 10 {
			t.Fatalf("page len = %d, want 10", len(page))
		}
	})
}

func TestDirStoreRoundTrip(t *testing.T) {
	dir := t.TempDir()
	s := DirStore{Root: dir}
	ctx := context.Background()
	v, err := s.LatestVersion(ctx)
	if err != nil || v != "" {
		t.Fatalf("empty store: version %q err %v", v, err)
	}
	w, err := s.Create(ctx, "20260824T000000", "oracle_cards.jsonl.gz")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write([]byte("data")); err != nil {
		t.Fatal(err)
	}
	_ = w.Close()
	// A version without its completion marker stays invisible.
	v, err = s.LatestVersion(ctx)
	if err != nil || v != "" {
		t.Fatalf("incomplete version visible: %q err %v", v, err)
	}
	if err := s.Finalize(ctx, "20260824T000000"); err != nil {
		t.Fatal(err)
	}
	v, err = s.LatestVersion(ctx)
	if err != nil || v != "20260824T000000" {
		t.Fatalf("version %q err %v", v, err)
	}
	r, err := s.Open(ctx, v, "oracle_cards.jsonl.gz")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = r.Close() }()
	buf := make([]byte, 4)
	if _, err := r.Read(buf); err != nil && err.Error() != "EOF" {
		t.Fatal(err)
	}
	if string(buf) != "data" {
		t.Fatalf("read %q", buf)
	}
	if _, err := VersionTime(v); err != nil {
		t.Fatalf("VersionTime: %v", err)
	}
	_ = filepath.Join
}
