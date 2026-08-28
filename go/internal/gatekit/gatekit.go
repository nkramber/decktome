// Package gatekit holds the helpers every gate and probe command shares:
// the spend guard, the quiet logger, the environment that requires real
// keys, the snapshot and collection loaders, and the word-to-enum maps.
// Five commands carried their own copy of each one (audit 2026-08-28).
package gatekit

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// SpendGuard refuses a paid run unless the named variable is "1". Every
// command that calls a real provider checks it first, so no run spends
// by accident (D-65).
func SpendGuard(name string) error {
	if os.Getenv(name) != "1" {
		return fmt.Errorf("this run calls a real provider and costs money: set %s=1 to allow it", name)
	}
	return nil
}

// Quiet is a logger that writes nothing. The gate documents are the
// record, and a log line on stdout would land inside one.
func Quiet() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }

// Env reads the process environment and forces LLM_REQUIRE_KEYS=1, so a
// gate never runs on the fake provider by mistake.
func Env(k string) string {
	if k == llm.EnvRequireKeys {
		return "1"
	}
	return os.Getenv(k)
}

// SnapshotDir reads CARDS_SNAPSHOT_DIR. An empty value is an error, so a
// command that needs the snapshot says so before it does anything else.
func SnapshotDir() (string, error) {
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		return "", fmt.Errorf("set CARDS_SNAPSHOT_DIR to a snapshot store root (for example .local/gcs/mtg-local-cards/scryfall)")
	}
	return dir, nil
}

// LoadSnapshot loads the local card index from CARDS_SNAPSHOT_DIR.
func LoadSnapshot(ctx context.Context, log *slog.Logger) (*cards.Index, error) {
	dir, err := SnapshotDir()
	if err != nil {
		return nil, err
	}
	idx, err := cards.LoadIndex(ctx, cards.DirStore{Root: dir}, log)
	if err != nil {
		return nil, fmt.Errorf("cards: %w", err)
	}
	if idx == nil {
		return nil, fmt.Errorf("no complete snapshot under %s", dir)
	}
	return idx, nil
}

// LoadOwned reads a ManaBox export into owned counts per oracle id. The
// note says what was read and how many rows resolved to nothing.
func LoadOwned(path string, idx *cards.Index) (map[string]int32, string, error) {
	f, err := os.Open(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, "", err
	}
	defer func() { _ = f.Close() }()
	rows, bad, err := collections.ParseManaBoxCSV(f)
	if err != nil {
		return nil, "", fmt.Errorf("collection: %w", err)
	}
	entries, unresolved := collections.Resolve(rows, idx)
	note := fmt.Sprintf("%d entries, %d cards, %d rows unresolved",
		len(entries), collections.CardCount(entries), len(bad)+len(unresolved))
	return collections.OracleCounts(entries), note, nil
}

// OrNone shows an empty string as the word none.
func OrNone(s string) string {
	if s == "" {
		return "none"
	}
	return s
}

// FormatID maps a format word onto the enum. The gate prompts use
// commander, standard, and modern.
func FormatID(s string) mtgv1.FormatId {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "commander":
		return mtgv1.FormatId_FORMAT_ID_COMMANDER
	case "standard":
		return mtgv1.FormatId_FORMAT_ID_STANDARD
	case "modern":
		return mtgv1.FormatId_FORMAT_ID_MODERN
	}
	return mtgv1.FormatId_FORMAT_ID_UNSPECIFIED
}

var colorNames = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G,
}

// Colors maps WUBRG letters onto the enum. An unknown letter is dropped.
func Colors(in []string) []mtgv1.Color {
	var out []mtgv1.Color
	for _, s := range in {
		if c, ok := colorNames[strings.ToUpper(strings.TrimSpace(s))]; ok {
			out = append(out, c)
		}
	}
	return out
}

// PoolRuleID maps a pool word onto the enum. Any other word is any-card.
func PoolRuleID(s string) mtgv1.PoolRule {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "owned_first":
		return mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
	case "owned_only":
		return mtgv1.PoolRule_POOL_RULE_OWNED_ONLY
	}
	return mtgv1.PoolRule_POOL_RULE_ANY_CARD
}
