package main

import (
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/rules"
)

// smokeSlots is what the classify fixture of the fake provider fills.
// The smoke flow of PR-23 builds from these slots over the trimmed
// snapshot (D-552).
type smokeSlots struct {
	Format         string   `json:"format"`
	Theme          string   `json:"theme"`
	CommanderNames []string `json:"commander_names"`
	Power          string   `json:"power"`
	PoolRule       string   `json:"pool_rule"`
}

// TestSmokeFixturesBuildFromTheTrimmedSnapshot proves the fixtures of
// the fake provider against the trimmed snapshot (PR-23, D-552). The
// smoke flow fills the slots of the classify fixture, and the generate
// fixture names the cards. Every name must sit in the pool the
// candidates builder makes for those slots over the fixture, and the
// deck must pass the rules engine. A miss or a block buys a repair turn
// the fake answers with the same list, and the flow ends with a deck
// the reader can not trust.
func TestSmokeFixturesBuildFromTheTrimmedSnapshot(t *testing.T) {
	ctx := context.Background()
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	idx, err := cards.LoadIndex(ctx, cards.DirStore{Root: filepath.Join(trimRoot, "scryfall")}, quiet)
	if err != nil {
		t.Fatalf("load the trimmed snapshot: %v", err)
	}
	if idx == nil {
		t.Fatalf("no complete snapshot under %s: run make deck-gate-trim", trimRoot)
	}
	fixtures := llm.Fixtures()
	var slots smokeSlots
	readFixture(t, fixtures, "classify.slot_fill.json", &slots)
	var deck struct {
		Cards []generate.Entry `json:"cards"`
	}
	readFixture(t, fixtures, "generate.json", &deck)
	// The repair turn answers the same deck, so a repair changes nothing
	// and the deck before it stands.
	gen, _ := fs.ReadFile(fixtures, "generate.json")
	rep, _ := fs.ReadFile(fixtures, "repair.json")
	if string(gen) != string(rep) {
		t.Error("repair.json differs from generate.json, so a repair turn of the smoke flow answers another deck")
	}

	format := gatekit.FormatID(slots.Format)
	if format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		t.Fatalf("the smoke slots name format %q, and this test reads a Commander deck", slots.Format)
	}
	if len(slots.CommanderNames) != 1 {
		t.Fatalf("the smoke slots name %d commanders, want 1", len(slots.CommanderNames))
	}
	commander, ok := idx.ByName(slots.CommanderNames[0])
	if !ok {
		t.Fatalf("the trimmed snapshot has no card named %q", slots.CommanderNames[0])
	}
	var bracket int32
	if _, err := parseBracket(slots.Power, &bracket); err != nil {
		t.Fatalf("power %q: %v", slots.Power, err)
	}
	poolRule := gatekit.PoolRuleID(slots.PoolRule)
	if poolRule == mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		t.Fatalf("the smoke slots name pool rule %q, which the gate does not know", slots.PoolRule)
	}

	// The flow uploads the fixture export first, so the build reads its
	// owned counts, as the agent service does.
	collPath := filepath.Join("..", "..", "internal", "collections", "testdata", "manabox_collection.csv")
	coll, err := gatekit.LoadCollection(collPath, idx)
	if err != nil {
		t.Fatalf("collection: %v", err)
	}
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	colors := commander.GetColorIdentity()
	list, err := cb.Build(idx, candidates.Request{
		Format:             format,
		Colors:             colors,
		Theme:              slots.Theme,
		CommanderOracleIDs: []string{commander.GetOracleId()},
		PoolRule:           poolRule,
		Owned:              coll.Oracle,
		Bracket:            bracket,
	})
	if err != nil {
		t.Fatalf("candidates: %v", err)
	}
	always := append([]*mtgv1.Card{commander}, generate.BasicLands(idx.ByName, colors)...)
	pool := generate.FromListOwned(list, always, coll.Oracle, poolRule == mtgv1.PoolRule_POOL_RULE_ANY_CARD)
	norm := generate.Normalize(pool, deck.Cards)
	for _, m := range norm.Misses {
		t.Errorf("the generate fixture names %q, and the pool of %d names does not hold it", m.Name, pool.Size())
	}
	have := 1
	for _, c := range norm.Cards {
		have += int(c.GetCount())
	}
	if want := generate.DeckSize(format); have != want {
		t.Errorf("the generate fixture holds %d cards with the commander, want %d", have, want)
	}

	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	built := &mtgv1.Deck{
		Format:             &mtgv1.Format{Id: format},
		Power:              gatekit.PowerLevel(bracket, ""),
		CommanderOracleIds: []string{commander.GetOracleId()},
		Cards:              norm.Cards,
	}
	v := cfg.Validate(rules.Input{Deck: built, PoolRule: poolRule, OracleCounts: coll.Oracle, Cards: idx})
	if !v.GetPassed() {
		var codes []string
		for _, f := range v.GetFindings() {
			if f.GetSeverity() == mtgv1.Severity_SEVERITY_BLOCK {
				codes = append(codes, f.GetCode()+": "+f.GetMessage())
			}
		}
		t.Errorf("the rules engine refuses the smoke deck: %s", strings.Join(codes, "; "))
	}
}

// parseBracket reads "bracket 3" the way the classify apply does: the
// last word is the number.
func parseBracket(power string, out *int32) (bool, error) {
	words := strings.Fields(strings.ToLower(power))
	if len(words) != 2 || words[0] != "bracket" {
		return false, nil
	}
	n, err := strconv.Atoi(words[1])
	if err != nil {
		return false, err
	}
	*out = int32(n) // #nosec G115 -- a bracket is 1 to 5.
	return true, nil
}

func readFixture(t *testing.T, fsys fs.FS, name string, into any) {
	t.Helper()
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		t.Fatalf("the fake provider has no fixture %s: %v", name, err)
	}
	if err := json.Unmarshal(data, into); err != nil {
		t.Fatalf("fixture %s: %v", name, err)
	}
}
