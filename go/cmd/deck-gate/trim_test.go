package main

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// TestTrimmedSnapshotBuildsEveryShortlist is the free dry-run lane of the
// deck gate (D-521). It loads the trimmed snapshot of the repo, and it
// runs the dry build of every prompt over it: every set resolves, every
// precon resolves, every collection reads, and every shortlist holds a
// card. It calls no provider. A prompt that needs a card the fixture
// lacks fails here, and `make deck-gate-trim` rewrites the fixture.
func TestTrimmedSnapshotBuildsEveryShortlist(t *testing.T) {
	ctx := context.Background()
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	idx, err := cards.LoadIndex(ctx, cards.DirStore{Root: filepath.Join(trimRoot, "scryfall")}, quiet)
	if err != nil {
		t.Fatalf("load the trimmed snapshot: %v", err)
	}
	if idx == nil {
		t.Fatalf("no complete snapshot under %s: run make deck-gate-trim", trimRoot)
	}
	var file struct {
		Prompts []prompt `json:"prompts"`
	}
	if err := json.Unmarshal(promptsJSON, &file); err != nil {
		t.Fatal(err)
	}
	collPath := filepath.Join("..", "..", "internal", "collections", "testdata", "manabox_collection.csv")
	binders, err := loadBinders(collPath, file.Prompts, idx)
	if err != nil {
		t.Fatalf("binders: %v", err)
	}
	store := meta.DirObjects{Root: trimRoot}
	version, err := meta.LatestPreconsVersion(ctx, store)
	if err != nil || version == "" {
		t.Fatalf("the fixture holds no precon table: %v", err)
	}
	rows, err := meta.ReadPrecons(ctx, store, version)
	if err != nil {
		t.Fatal(err)
	}
	tbl := precons.NewTable(version, rows)
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	preconSet, err := precons.Load(idx)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := rules.Load(); err != nil {
		t.Fatal(err)
	}
	for _, p := range file.Prompts {
		r := build(ctx, nil, cb, idx, binders, p, nil, true, preconSet, tbl)
		if r.err != nil {
			t.Errorf("prompt %d, %s: %v", p.ID, p.Name, r.err)
			continue
		}
		if r.shortlist == 0 {
			t.Errorf("prompt %d, %s: the theme search kept no card", p.ID, p.Name)
		}
		if len(p.Sets) > 0 && r.inSet == 0 {
			t.Errorf("prompt %d, %s: the set family holds no pool card", p.ID, p.Name)
		}
		if len(p.ExcludePrecons) > 0 && len(r.excluded) == 0 {
			t.Errorf("prompt %d, %s: the exclusion took no card out", p.ID, p.Name)
		}
	}
}

// TestTrimmedSnapshotStaysUnderTheBudget is D-521: the fixture holds at
// most 10 MB, or it leaves the repo.
func TestTrimmedSnapshotStaysUnderTheBudget(t *testing.T) {
	var total int64
	err := filepath.Walk(trimRoot, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			total += info.Size()
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", trimRoot, err)
	}
	if total == 0 || total > trimBudget {
		t.Errorf("the fixture holds %d bytes, the budget is %d (D-521)", total, trimBudget)
	}
}

// TestStripKeepsTheLoaderKeys: a stripped line drops the images and the
// shop links, keeps what the loader reads, and still parses as a card.
func TestStripKeepsTheLoaderKeys(t *testing.T) {
	line := `{"object":"card","id":"p1","oracle_id":"o1","name":"Sol Ring","mana_cost":"{1}","cmc":1,"type_line":"Artifact","oracle_text":"{T}: Add {C}{C}.","layout":"normal","legalities":{"commander":"legal"},"set":"cmm","collector_number":"1","rarity":"uncommon","image_uris":{"small":"https://x/s.jpg"},"purchase_uris":{"tcgplayer":"https://x"},"related_uris":{"edhrec":"https://x"},"prices":{"usd":"1.00"},"card_faces":[{"oracle_id":"o1","name":"Sol Ring","image_uris":{"small":"https://x"}}]}`
	var obj map[string]any
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		t.Fatal(err)
	}
	card := stripCard(obj)
	for _, gone := range []string{"image_uris", "purchase_uris", "related_uris", "object"} {
		if _, ok := card[gone]; ok {
			t.Errorf("stripCard kept %q", gone)
		}
	}
	faces := card["card_faces"].([]any)
	if _, ok := faces[0].(map[string]any)["image_uris"]; ok {
		t.Error("stripCard kept the face images")
	}
	raw, err := json.Marshal(card)
	if err != nil {
		t.Fatal(err)
	}
	parsed, _, err := cards.LoadCardsStats(strings.NewReader(string(raw)+"\n"), "oracle_cards.jsonl")
	if err != nil || len(parsed) != 1 || parsed[0].GetName() != "Sol Ring" || parsed[0].GetOracleId() != "o1" {
		t.Fatalf("the loader did not read the stripped line: %v %+v", err, parsed)
	}
	printing := stripPrinting(obj)
	if _, ok := printing["image_uris"]; ok || printing["id"] != "p1" || printing["set"] != "cmm" {
		t.Errorf("stripPrinting = %v", printing)
	}
	tag := stripTag(map[string]any{"id": "t1", "slug": "ramp", "parent_ids": []any{}, "child_ids": []any{},
		"taggings": []any{map[string]any{"oracle_id": "o1"}, map[string]any{"oracle_id": "o2"}}}, map[string]bool{"o1": true})
	if kept := tag["taggings"].([]any); len(kept) != 1 {
		t.Errorf("stripTag kept %d taggings, want the needed card alone", len(kept))
	}
}
