package main

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
)

// The trimmed snapshot of D-521. A dry run over the full snapshot reaches
// a few thousand cards, and a snapshot cut to those cards serves the same
// dry run from the repo, so CI builds every shortlist for nothing. The
// trim keeps every card the run touched, every card of the collection
// fixtures, every card of the embedded precon lists, every card of the
// excluded products, every printing of the set families the prompts
// name, and every basic land. Each kept line holds the keys the loaders
// read and nothing else, so the images and the shop links stay out.
//
// The trimmed index is not the full index. The theme search reads the
// whole corpus for its noise rule (D-411) and for the popularity share,
// so a shortlist over the trimmed index can differ from the same prompt
// over the full one. The dry lane proves that every prompt builds a
// shortlist, resolves its sets and its precons, and reads its
// collection. It proves nothing about the cards a paid run picks.

// trimCommanderPool is how many commanders of a delegated pick the trim
// keeps, so the pick still has a pool.
const trimCommanderPool = 25

// trimRoot is the fixture root, relative to the package. The scryfall
// store sits under it, and the meta store beside the scryfall store, the
// way gcpenv.MetaStore reads them.
const trimRoot = "testdata/snapshot"

// trimBudget is the size the owner allowed for the fixture (D-521).
const trimBudget = 10_000_000

// cardKeys are the keys the card loader reads from an oracle card line.
var cardKeys = keySet("id", "oracle_id", "name", "mana_cost", "cmc", "colors", "color_identity", "type_line",
	"oracle_text", "keywords", "layout", "card_faces", "power", "toughness", "loyalty", "produced_mana",
	"legalities", "game_changer", "edhrec_rank", "set", "set_name", "collector_number", "rarity", "artist",
	"digital", "prices", "released_at")

// faceKeys are the keys the card loader reads from a face.
var faceKeys = keySet("oracle_id", "name", "artist", "mana_cost", "type_line", "oracle_text", "power", "toughness", "loyalty")

// printingKeys are the keys the printing loader reads from a default
// card line.
var printingKeys = keySet("id", "oracle_id", "name", "set", "set_name", "collector_number", "layout", "rarity",
	"artist", "digital", "released_at", "prices", "card_faces")

func keySet(keys ...string) map[string]bool {
	out := make(map[string]bool, len(keys))
	for _, k := range keys {
		out[k] = true
	}
	return out
}

// basicLandNames are the cards every deck pads with.
var basicLandNames = []string{"Plains", "Island", "Swamp", "Mountain", "Forest", "Wastes",
	"Snow-Covered Plains", "Snow-Covered Island", "Snow-Covered Swamp", "Snow-Covered Mountain", "Snow-Covered Forest"}

// writeTrimmed writes the trimmed snapshot under root: the scryfall store
// with the source version, and the precon table beside it. It reads the
// source store from CARDS_SNAPSHOT_DIR.
func writeTrimmed(ctx context.Context, log io.Writer, root string, idx *cards.Index, results []result,
	binders map[string]*gatekit.Collection, preconSet *precons.Set, tbl *precons.Table) error {
	srcDir, err := gatekit.SnapshotDir()
	if err != nil {
		return err
	}
	src := cards.DirStore{Root: srcDir}
	version, err := src.LatestVersion(ctx)
	if err != nil {
		return err
	}
	needed := map[string]bool{}
	printings := map[string]bool{}
	sets := map[string]bool{}
	productKeys := map[string]bool{}
	for _, r := range results {
		if r.err != nil {
			return fmt.Errorf("prompt %d failed, so the trim would miss its cards: %w", r.prompt.ID, r.err)
		}
		for _, id := range r.touched {
			needed[id] = true
		}
		for _, code := range r.setCodes {
			sets[code] = true
		}
		for _, k := range r.productKeys {
			productKeys[k] = true
		}
	}
	for _, b := range binders {
		for id := range b.Oracle {
			needed[id] = true
		}
		for id := range b.Printings {
			printings[id] = true
		}
	}
	for _, pc := range preconSet.All() {
		for _, id := range pc.OracleIDs {
			needed[id] = true
		}
	}
	if tbl != nil {
		for key := range productKeys {
			p, ok := tbl.Get(key)
			if !ok {
				continue
			}
			for id := range p.Counts() {
				needed[id] = true
			}
			for id := range p.Printings() {
				printings[id] = true
			}
		}
	}
	for _, name := range basicLandNames {
		if c, ok := idx.ByName(name); ok {
			needed[c.GetOracleId()] = true
		}
	}

	out := cards.DirStore{Root: filepath.Join(root, "scryfall")}
	sizes := map[string]int64{}
	// The printings come first: a printing of a kept set or a kept
	// product names a card the run never touched, and that card must
	// join the oracle file too.
	n, err := trimPrintings(ctx, src, out, version, needed, printings, sets)
	if err != nil {
		return err
	}
	sizes["default_cards.jsonl.gz"] = n
	if sizes["oracle_cards.jsonl.gz"], err = trimLines(ctx, src, out, version, "oracle_cards.jsonl.gz", func(obj map[string]any) (map[string]any, bool) {
		if !oracleNeeded(obj, needed) {
			return nil, false
		}
		return stripCard(obj), true
	}); err != nil {
		return err
	}
	if sizes["oracle_tags.jsonl.gz"], err = trimLines(ctx, src, out, version, "oracle_tags.jsonl.gz", func(obj map[string]any) (map[string]any, bool) {
		return stripTag(obj, needed), true
	}); err != nil {
		return err
	}
	if sizes[cards.RulingsFile], err = trimLines(ctx, src, out, version, cards.RulingsFile, func(obj map[string]any) (map[string]any, bool) {
		id, _ := obj["oracle_id"].(string)
		return obj, needed[id]
	}); err != nil {
		return err
	}
	if sizes[cards.SetsFile], err = copyFile(ctx, src, out, version, cards.SetsFile); err != nil {
		return err
	}
	if rec, ok, err := src.ReadLegalityDiff(ctx, version); err != nil {
		return err
	} else if ok {
		if err := out.WriteLegalityDiff(ctx, version, rec); err != nil {
			return err
		}
	}
	if err := out.Finalize(ctx, version); err != nil {
		return err
	}
	if tbl != nil {
		n, err := trimPrecons(ctx, filepath.Dir(srcDir), root, tbl.Version, productKeys)
		if err != nil {
			return err
		}
		sizes["meta/precons"] = n
	}
	var total int64
	names := make([]string, 0, len(sizes))
	for name := range sizes {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		_, _ = fmt.Fprintf(log, "  %-24s %9d bytes\n", name, sizes[name])
		total += sizes[name]
	}
	_, _ = fmt.Fprintf(log, "trimmed snapshot %s under %s: %d cards, %d printings kept, %d bytes in all, the budget is %d\n",
		version, root, len(needed), len(printings), total, trimBudget)
	if total > trimBudget {
		return fmt.Errorf("the trimmed snapshot holds %d bytes, over the budget of %d (D-521)", total, trimBudget)
	}
	return nil
}

// trimPrintings writes the default cards the fixture keeps: every
// printing of the collections and the products, every printing of the
// kept sets, and the newest paper printing of every other card the run
// touched. The cards of the kept sets and products join needed.
func trimPrintings(ctx context.Context, src, out cards.DirStore, version string, needed, printings, sets map[string]bool) (int64, error) {
	// Pass one picks the lines. The first paper printing of a card wins,
	// and a card with paper printings none keeps its first digital one.
	chosen := map[string]string{}
	chosenPaper := map[string]bool{}
	err := scanLines(ctx, src, version, "default_cards.jsonl.gz", func(obj map[string]any) error {
		id, _ := obj["id"].(string)
		set, _ := obj["set"].(string)
		oracle := oracleOf(obj)
		digital, _ := obj["digital"].(bool)
		switch {
		case printings[id], sets[set]:
			printings[id] = true
			if oracle != "" {
				needed[oracle] = true
			}
		case oracle != "" && needed[oracle]:
			if _, ok := chosen[oracle]; !ok || (!digital && !chosenPaper[oracle]) {
				chosen[oracle] = id
				chosenPaper[oracle] = !digital
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	for _, id := range chosen {
		printings[id] = true
	}
	return trimLines(ctx, src, out, version, "default_cards.jsonl.gz", func(obj map[string]any) (map[string]any, bool) {
		id, _ := obj["id"].(string)
		if !printings[id] {
			return nil, false
		}
		return stripPrinting(obj), true
	})
}

// oracleOf reads the Oracle id of a line, from the card or its first
// face.
func oracleOf(obj map[string]any) string {
	if id, _ := obj["oracle_id"].(string); id != "" {
		return id
	}
	faces, _ := obj["card_faces"].([]any)
	for _, f := range faces {
		if m, ok := f.(map[string]any); ok {
			if id, _ := m["oracle_id"].(string); id != "" {
				return id
			}
		}
	}
	return ""
}

// oracleNeeded reports whether the card or one of its faces is needed.
func oracleNeeded(obj map[string]any, needed map[string]bool) bool {
	if id, _ := obj["oracle_id"].(string); needed[id] {
		return true
	}
	faces, _ := obj["card_faces"].([]any)
	for _, f := range faces {
		if m, ok := f.(map[string]any); ok {
			if id, _ := m["oracle_id"].(string); needed[id] {
				return true
			}
		}
	}
	return false
}

// stripCard keeps the keys the card loader reads.
func stripCard(obj map[string]any) map[string]any {
	out := keep(obj, cardKeys)
	if faces, ok := obj["card_faces"].([]any); ok {
		kept := make([]any, 0, len(faces))
		for _, f := range faces {
			if m, ok := f.(map[string]any); ok {
				kept = append(kept, keep(m, faceKeys))
			}
		}
		out["card_faces"] = kept
	}
	return out
}

// stripPrinting keeps the keys the printing loader reads. A face keeps
// its Oracle id alone.
func stripPrinting(obj map[string]any) map[string]any {
	out := keep(obj, printingKeys)
	if faces, ok := obj["card_faces"].([]any); ok {
		kept := make([]any, 0, len(faces))
		for _, f := range faces {
			if m, ok := f.(map[string]any); ok {
				kept = append(kept, keep(m, keySet("oracle_id")))
			}
		}
		out["card_faces"] = kept
	}
	return out
}

// stripTag keeps the tag with its hierarchy, and the taggings of the
// needed cards alone.
func stripTag(obj map[string]any, needed map[string]bool) map[string]any {
	out := keep(obj, keySet("id", "slug", "parent_ids", "child_ids", "taggings"))
	taggings, _ := obj["taggings"].([]any)
	kept := make([]any, 0)
	for _, t := range taggings {
		if m, ok := t.(map[string]any); ok {
			if id, _ := m["oracle_id"].(string); needed[id] {
				kept = append(kept, m)
			}
		}
	}
	out["taggings"] = kept
	return out
}

func keep(obj map[string]any, keys map[string]bool) map[string]any {
	out := make(map[string]any, len(keys))
	for k, v := range obj {
		if keys[k] {
			out[k] = v
		}
	}
	return out
}

// scanLines runs fn over every JSON object of a gzip JSONL file.
func scanLines(ctx context.Context, src cards.DirStore, version, file string, fn func(obj map[string]any) error) error {
	r, err := src.Open(ctx, version, file)
	if err != nil {
		return err
	}
	defer func() { _ = r.Close() }()
	gz, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("%s: %w", file, err)
	}
	defer func() { _ = gz.Close() }()
	sc := bufio.NewScanner(gz)
	sc.Buffer(make([]byte, 0, 1<<20), 64<<20)
	for sc.Scan() {
		line := sc.Bytes()
		if len(strings.TrimSpace(string(line))) == 0 {
			continue
		}
		var obj map[string]any
		if err := json.Unmarshal(line, &obj); err != nil {
			return fmt.Errorf("%s: %w", file, err)
		}
		if err := fn(obj); err != nil {
			return err
		}
	}
	return sc.Err()
}

// trimLines writes the objects fn keeps into the same file of the out
// store, gzip JSONL, and returns the size written.
func trimLines(ctx context.Context, src, out cards.DirStore, version, file string, fn func(obj map[string]any) (map[string]any, bool)) (int64, error) {
	w, err := out.Create(ctx, version, file)
	if err != nil {
		return 0, err
	}
	counter := &countWriter{w: w}
	gz := gzip.NewWriter(counter)
	enc := json.NewEncoder(gz)
	enc.SetEscapeHTML(false)
	err = scanLines(ctx, src, version, file, func(obj map[string]any) error {
		kept, ok := fn(obj)
		if !ok {
			return nil
		}
		return enc.Encode(kept)
	})
	if err != nil {
		_ = gz.Close()
		_ = w.Close()
		return 0, err
	}
	if err := gz.Close(); err != nil {
		_ = w.Close()
		return 0, err
	}
	if err := w.Close(); err != nil {
		return 0, err
	}
	return counter.n, nil
}

// copyFile copies one file of the version as it is.
func copyFile(ctx context.Context, src, out cards.DirStore, version, file string) (int64, error) {
	r, err := src.Open(ctx, version, file)
	if err != nil {
		return 0, err
	}
	defer func() { _ = r.Close() }()
	w, err := out.Create(ctx, version, file)
	if err != nil {
		return 0, err
	}
	n, err := io.Copy(w, r)
	if err != nil {
		_ = w.Close()
		return 0, err
	}
	return n, w.Close()
}

// trimPrecons writes the rows of the products the prompts resolved into
// the meta store beside the trimmed scryfall store, under the source
// table's version.
func trimPrecons(ctx context.Context, srcRoot, root, version string, productKeys map[string]bool) (int64, error) {
	rows, err := meta.ReadPrecons(ctx, meta.DirObjects{Root: srcRoot}, version)
	if err != nil {
		return 0, err
	}
	var kept []meta.Precon
	for _, r := range rows {
		if productKeys[r.Key()] {
			kept = append(kept, r)
		}
	}
	if len(kept) != len(productKeys) {
		return 0, fmt.Errorf("the precon table %s holds %d of the %d products the prompts name", version, len(kept), len(productKeys))
	}
	store := meta.DirObjects{Root: root}
	if err := meta.WritePrecons(ctx, store, version, kept); err != nil {
		return 0, err
	}
	info, err := os.Stat(filepath.Join(root, filepath.FromSlash(meta.PreconsName(version))))
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

type countWriter struct {
	w io.Writer
	n int64
}

func (c *countWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.n += int64(n)
	return n, err
}
