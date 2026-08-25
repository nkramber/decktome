// Command candidates-review writes the PR-6 gate document. It runs the
// gate prompts (prompts.json) against a local snapshot and prints a
// Markdown file for the owner to score: top 40 candidates per prompt,
// with role, owned count, and the signals that put each card there.
//
// Usage:
//
//	CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/candidates-review -collection ../go/internal/collections/testdata/manabox_collection.csv \
//	  > ../docs/reference/pr6-candidate-review.md
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
)

//go:embed prompts.json
var promptsJSON []byte

type prompt struct {
	ID         int      `json:"id"`
	Theme      string   `json:"theme"`
	Format     string   `json:"format"`
	Colors     []string `json:"colors"`
	Collection bool     `json:"collection"`
}

var formats = map[string]mtgv1.FormatId{
	"commander": mtgv1.FormatId_FORMAT_ID_COMMANDER, "standard": mtgv1.FormatId_FORMAT_ID_STANDARD,
	"pioneer": mtgv1.FormatId_FORMAT_ID_PIONEER, "modern": mtgv1.FormatId_FORMAT_ID_MODERN,
	"legacy": mtgv1.FormatId_FORMAT_ID_LEGACY, "vintage": mtgv1.FormatId_FORMAT_ID_VINTAGE,
	"pauper": mtgv1.FormatId_FORMAT_ID_PAUPER,
}

var colors = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G,
}

func main() {
	collectionPath := flag.String("collection", "", "ManaBox CSV for the owned-first prompts")
	top := flag.Int("top", 40, "candidates to print per prompt")
	flag.Parse()
	if err := run(*collectionPath, *top, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(collectionPath string, top int, w io.Writer) error {
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		return fmt.Errorf("set CARDS_SNAPSHOT_DIR to a snapshot store root (for example .local/gcs/mtg-local-cards/scryfall)")
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, logger)
	if err != nil {
		return err
	}
	if idx == nil {
		return fmt.Errorf("no complete snapshot under %s", dir)
	}
	var owned map[string]int32
	var ownedNote string
	if collectionPath != "" {
		f, err := os.Open(collectionPath)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()
		rows, bad, err := collections.ParseManaBoxCSV(f)
		if err != nil {
			return err
		}
		entries, unresolved := collections.Resolve(rows, idx)
		owned = collections.OracleCounts(entries)
		ownedNote = fmt.Sprintf("%d entries, %d cards, %d rows unresolved", len(entries), collections.CardCount(entries), len(bad)+len(unresolved))
	}
	var file struct {
		VerifiedAt string   `json:"verified_at"`
		Prompts    []prompt `json:"prompts"`
	}
	if err := json.Unmarshal(promptsJSON, &file); err != nil {
		return err
	}
	b, err := candidates.New()
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(w, "# PR-6 candidate review\n\n")
	_, _ = fmt.Fprintf(w, "Snapshot: %s. Prompts: %s. Collection: %s.\n\n", idx.AsOf.Format("2006-01-02"), file.VerifiedAt, orNone(ownedNote))
	_, _ = fmt.Fprintf(w, "Gate: for 20 theme prompts, a human confirms the top %d candidates are on theme in at least 18. Ten run with no collection.\n\n", top)
	_, _ = fmt.Fprintf(w, "Score each prompt in the table, then fill the total.\n\n")
	_, _ = fmt.Fprintf(w, "| # | Theme | Mode | On theme (yes/no) | Notes |\n|---|---|---|---|---|\n")
	for _, p := range file.Prompts {
		mode := "any-card"
		if p.Collection {
			mode = "owned-first"
		}
		_, _ = fmt.Fprintf(w, "| %d | %s | %s | | |\n", p.ID, p.Theme, mode)
	}
	_, _ = fmt.Fprintf(w, "\nTotal on theme: __ of 20.\n\n")

	for _, p := range file.Prompts {
		req := candidates.Request{Format: formats[p.Format], Theme: p.Theme}
		for _, c := range p.Colors {
			req.Colors = append(req.Colors, colors[c])
		}
		mode := "any-card"
		if p.Collection {
			if owned == nil {
				_, _ = fmt.Fprintf(w, "## %d. %s (%s) - skipped, no collection given\n\n", p.ID, p.Theme, p.Format)
				continue
			}
			req.Owned = owned
			req.PoolRule = mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
			mode = "owned-first"
		}
		list, err := b.Build(idx, req)
		if err != nil {
			return fmt.Errorf("prompt %d: %w", p.ID, err)
		}
		_, _ = fmt.Fprintf(w, "## %d. %s (%s, %s, %s)\n\n", p.ID, p.Theme, p.Format, strings.Join(p.Colors, ""), mode)
		_, _ = fmt.Fprintf(w, "Theme signals: %s.\n\n", list.Theme.Describe())
		s := list.Stats
		_, _ = fmt.Fprintf(w, "Funnel: %d legal in colors, %d on theme, %d owned, %d on theme and owned, %d returned, %d upgrades.\n\n", s.Pool, s.OnTheme, s.Owned, s.OnThemeOwned, s.Returned, s.UpgradeSize)
		if s.ThinTheme {
			_, _ = fmt.Fprintf(w, "Thin theme: the collection holds under %d on-theme cards. PR-7 asks the pool-mode question again here (D-63).\n\n", candidates.ThinThemeFloor)
		}
		if mode == "owned-first" {
			_, _ = fmt.Fprintf(w, "The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.\n\n")
		}
		_, _ = fmt.Fprintf(w, "| Rank | Card | Role | Owned | Score | Signals |\n|---|---|---|---|---|---|\n")
		merged := append(append([]candidates.Candidate(nil), list.Candidates...), list.Upgrades...)
		for i, c := range byScore(merged, top) {
			_, _ = fmt.Fprintf(w, "| %d | %s | %s | %d | %.2f | %s |\n", i+1, c.Card.Name, candidates.RoleName(c.Role), c.Owned, c.Score, strings.Join(c.Signals, " "))
		}
		if len(list.Upgrades) > 0 {
			_, _ = fmt.Fprintf(w, "\nTop upgrades (unowned):\n\n")
			for i, c := range list.Upgrades {
				if i >= 10 {
					break
				}
				_, _ = fmt.Fprintf(w, "- %s (%s, %.2f)\n", c.Card.Name, candidates.RoleName(c.Role), c.Score)
			}
		}
		_, _ = fmt.Fprintln(w)
	}
	return nil
}

// byScore returns the top n across roles by score, so the reviewer sees
// the strongest theme matches first, not the land group.
func byScore(cs []candidates.Candidate, n int) []candidates.Candidate {
	out := append([]candidates.Candidate(nil), cs...)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j].Score > out[j-1].Score; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	if len(out) > n {
		out = out[:n]
	}
	return out
}

func orNone(s string) string {
	if s == "" {
		return "none"
	}
	return s
}
