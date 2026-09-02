// Package gatekit holds the helpers every gate and probe command shares:
// the spend guard, the output guard, the quiet logger, the environment
// that requires real keys, the snapshot and collection loaders, the
// word-to-enum maps, and the deck counters the reports print. A command
// must not carry its own copy of one.
package gatekit

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
	"github.com/nkramber/mtg-deck-builder/go/internal/spellbook"
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

// RefuseExisting returns an error when any named path exists. A paid
// result is never overwritten (D-65), and the check runs before the
// first provider call. An empty path is skipped.
func RefuseExisting(paths ...string) error {
	for _, p := range paths {
		if p == "" || p == os.DevNull {
			continue
		}
		if _, err := os.Stat(p); err == nil {
			return fmt.Errorf("%s exists, and a result is never overwritten (D-65): name a new file", p)
		}
	}
	return nil
}

// ParseIDs reads a comma-separated list of integer ids, as the -only
// flag of a gate takes it. An empty string gives nil.
func ParseIDs(s string) ([]int, error) {
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	var out []int
	for _, part := range strings.Split(s, ",") {
		id, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("-only takes integer ids, comma separated: %w", err)
		}
		out = append(out, id)
	}
	return out, nil
}

// CostWord shows a report's cost as dollars, or the word unpriced when
// no cost is known. A nil cost is not zero (M-1).
func CostWord(rep llm.Report) string {
	if rep.CostUSD == nil {
		return "unpriced"
	}
	return fmt.Sprintf("$%.4f", *rep.CostUSD)
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

// sixtySteps maps the power words of a 60-card format onto the enum. The
// corpus calls the top step tournament-meta.
var sixtySteps = map[string]mtgv1.SixtyStep{
	"casual":     mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
	"fnm":        mtgv1.SixtyStep_SIXTY_STEP_FNM,
	"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}

// PowerLevel builds the power level of a gate prompt. A bracket above
// zero wins, then a 60-card step word. Neither gives nil.
func PowerLevel(bracket int32, word string) *mtgv1.PowerLevel {
	if bracket > 0 {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bracket}}
	}
	if step, ok := sixtySteps[strings.ToLower(strings.TrimSpace(word))]; ok {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: step}}
	}
	return nil
}

// CountCards sums the counts of the main deck.
func CountCards(d *mtgv1.Deck) int {
	n := 0
	for _, c := range d.GetCards() {
		n += int(c.GetCount())
	}
	return n
}

// CountSideboard sums the counts of the sideboard.
func CountSideboard(d *mtgv1.Deck) int {
	n := 0
	for _, c := range d.GetSideboard() {
		n += int(c.GetCount())
	}
	return n
}

// BlockFindings returns the findings that stop the deck.
func BlockFindings(d *mtgv1.Deck) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	for _, f := range d.GetValidation().GetFindings() {
		if f.GetSeverity() == mtgv1.Severity_SEVERITY_BLOCK {
			out = append(out, f)
		}
	}
	return out
}

var colorLetter = map[mtgv1.Color]string{
	mtgv1.Color_COLOR_W: "W", mtgv1.Color_COLOR_U: "U", mtgv1.Color_COLOR_B: "B",
	mtgv1.Color_COLOR_R: "R", mtgv1.Color_COLOR_G: "G",
}

// ColorLetters maps the enum onto WUBRG letters, the way Scryfall writes
// them. An unknown value is dropped.
func ColorLetters(in []mtgv1.Color) []string {
	var out []string
	for _, c := range in {
		if l, ok := colorLetter[c]; ok {
			out = append(out, l)
		}
	}
	return out
}

// Profiler makes the bracket profiler a gate builds with: the bands, the
// snapshot's tags, and the live Commander Spellbook client (PR-14A). The
// endpoint is free and rate limited on this side (D-459), so a gate
// reads it as the app does.
func Profiler(idx *cards.Index, cfg *rules.Config, log *slog.Logger) (*profile.Profiler, error) {
	return profile.New(cfg, func() *cards.TagIndex { return idx.Tags() }, spellbook.New(nil, "", log))
}
