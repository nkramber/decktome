package generate

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

func fakeConfig() *llm.Config {
	roles := map[llm.Role]llm.RoleSpec{}
	for _, r := range llm.Roles {
		roles[r] = llm.RoleSpec{Provider: llm.FakeName, Model: "fake-" + string(r), MaxOutputTokens: 4096}
	}
	return &llm.Config{VerifiedAt: "2026-08-24", Roles: roles}
}

// source answers the rules engine. It holds the same cards as the pool,
// so a card outside the pool is unknown to both.
type source map[string]*mtgv1.Card

func (s source) ByOracleID(id string) (*mtgv1.Card, bool) { c, ok := s[id]; return c, ok }

func step(t *testing.T, out deckOut) llm.Step {
	t.Helper()
	if out.Cards == nil {
		out.Cards = []Entry{}
	}
	if out.Sideboard == nil {
		out.Sideboard = []Entry{}
	}
	raw, err := json.Marshal(out)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return llm.Step{Output: raw}
}

func testBuilder(t *testing.T, steps ...llm.Step) (*Builder, source, *llm.Script) {
	t.Helper()
	sc := llm.NewScript(steps...)
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	cfg, err := rules.Load()
	if err != nil {
		t.Fatalf("rules: %v", err)
	}
	src := source{}
	for _, name := range testPool().Names() {
		c, _ := testPool().Card(name)
		src[c.GetOracleId()] = c
	}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewBuilder(c, cfg, src, log), src, sc
}

func testRequest() Request {
	return Request{
		SessionID:    "s-1",
		Format:       mtgv1.FormatId_FORMAT_ID_MODERN,
		Plan:         "a lifegain deck",
		Pool:         testPool(),
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Limits:       "60 cards minimum, four copies per name.",
		LegalityAsOf: "2026-08-24",
	}
}

// TestBuildRepairsAnInventedName is F-13 end to end. The model names a
// card the shortlist does not hold, the normalizer refuses it, and the
// repair turn gets the miss with the near names. The invented name never
// reaches the deck.
func TestBuildRepairsAnInventedName(t *testing.T) {
	first := step(t, deckOut{
		Summary: "a lifegain deck",
		Cards: []Entry{
			{Name: "Ajani's Welcome", Count: 4, Role: "synergy", Reason: "gains life"},
			{Name: "Craterhoof Behemoth", Count: 1, Role: "wincon", Reason: "ends the game"},
		},
	})
	second := step(t, deckOut{
		Summary: "a lifegain deck",
		Cards: []Entry{
			{Name: "Ajani's Welcome", Count: 4, Role: "synergy", Reason: "gains life"},
			{Name: "Ajani's Pridemate", Count: 4, Role: "threat", Reason: "grows on each gain"},
		},
	})
	b, _, sc := testBuilder(t, first, second)
	got, err := b.Build(context.Background(), testRequest(), nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if !got.Repaired {
		t.Error("the repair turn did not run on an invented name")
	}
	if len(sc.Calls) != 2 {
		t.Fatalf("provider calls = %d, want 2", len(sc.Calls))
	}
	// The repair turn must name the miss, or the model cannot fix it.
	if in := sc.Calls[1].Input; !strings.Contains(in, "Craterhoof Behemoth") {
		t.Error("the repair input did not name the miss")
	}
	for _, c := range got.Deck.GetCards() {
		if c.GetName() == "Craterhoof Behemoth" {
			t.Fatal("an invented name reached the deck")
		}
	}
	if len(got.Notes) != 0 {
		t.Errorf("notes = %v, want none: the repair turn fixed the miss", got.Notes)
	}
}

// TestBuildNotesANameThatMissesTwice covers the second miss. The roadmap
// gives the model one repair turn, and a name it writes again becomes a
// user-visible note.
func TestBuildNotesANameThatMissesTwice(t *testing.T) {
	bad := deckOut{
		Summary: "a lifegain deck",
		Cards: []Entry{
			{Name: "Ajani's Welcome", Count: 4, Role: "synergy", Reason: "gains life"},
			{Name: "Craterhoof Behemoth", Count: 1, Role: "wincon", Reason: "ends the game"},
		},
	}
	b, _, _ := testBuilder(t, step(t, bad), step(t, bad))
	got, err := b.Build(context.Background(), testRequest(), nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(got.Notes) != 1 || !strings.Contains(got.Notes[0], "Craterhoof Behemoth") {
		t.Fatalf("notes = %v, want one note naming the card", got.Notes)
	}
	for _, c := range got.Deck.GetCards() {
		if c.GetName() == "Craterhoof Behemoth" {
			t.Fatal("an invented name reached the deck after two misses")
		}
	}
	// The deck still reaches the caller, with the referee's answer on it.
	if got.Deck.GetValidation() == nil {
		t.Error("the deck carries no validation result")
	}
}

// TestBuildAttachesTheValidationResult is the roadmap rule that the deck
// reaches the user with the referee's answer, whatever the verdict.
func TestBuildAttachesTheValidationResult(t *testing.T) {
	one := step(t, deckOut{Summary: "small", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 4, Role: "synergy", Reason: "gains life"},
	}})
	b, _, sc := testBuilder(t, one, one)
	got, err := b.Build(context.Background(), testRequest(), nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	// A four-card deck fails the size check, so the repair turn runs.
	if !got.Repaired || len(sc.Calls) != 2 {
		t.Errorf("repaired = %v, calls = %d, want a repair turn", got.Repaired, len(sc.Calls))
	}
	if len(blocking(got.Deck.GetValidation())) == 0 {
		t.Error("a four-card Modern deck produced no block finding")
	}
	// One cache key per session, on both turns (roadmap PR-8).
	for i, c := range sc.Calls {
		if c.CacheKey != "s-1" {
			t.Errorf("call %d cache key = %q, want the session id", i, c.CacheKey)
		}
	}
}

// TestSixtyCardDeckCarriesNoCommander is D-233. Only Commander has a
// command zone. A probe of 2026-08-27 built a Modern deck with Karlov in
// the command zone, and the engine refused it as not legal in the format.
func TestSixtyCardDeckCarriesNoCommander(t *testing.T) {
	one := step(t, deckOut{Summary: "burn", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 4, Role: "synergy", Reason: "gains life"},
	}})
	b, _, _ := testBuilder(t, one, one)
	req := testRequest()
	req.Format = mtgv1.FormatId_FORMAT_ID_MODERN
	req.Commanders = []string{"o-karlov"}
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if ids := got.Deck.GetCommanderOracleIds(); len(ids) != 0 {
		t.Errorf("a Modern deck carries commanders %v", ids)
	}
	// Commander keeps its command zone.
	req.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	b2, _, _ := testBuilder(t, one, one)
	got2, err := b2.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(got2.Deck.GetCommanderOracleIds()) != 1 {
		t.Error("a Commander deck lost its commander")
	}
}

// TestLockedCardMustReachTheDeck is D-242. The locked row asks which
// cards the deck must keep, and agentsvc never passed the answer, so a
// deck without the card answered the user's own instruction with silence.
func TestLockedCardMustReachTheDeck(t *testing.T) {
	// The model leaves Sol Ring out on both turns.
	without := deckOut{Summary: "s", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 1, Role: "synergy", Reason: "gains life"},
	}}
	b, _, sc := testBuilder(t, step(t, without), step(t, without))
	req := testRequest()
	req.Locked = []string{"o-solring"}
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	var found *mtgv1.Finding
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeLockedCardMissing {
			found = f
		}
	}
	if found == nil {
		t.Fatal("a deck without the locked card carried no finding")
	}
	if found.GetSeverity() != mtgv1.Severity_SEVERITY_BLOCK {
		t.Errorf("severity = %v, want BLOCK so the repair turn runs", found.GetSeverity())
	}
	if !strings.Contains(found.GetMessage(), "Sol Ring") {
		t.Errorf("the finding does not name the card: %q", found.GetMessage())
	}
	// A block finding buys one repair turn (D-223).
	if len(sc.Calls) != 2 {
		t.Errorf("provider calls = %d, want 2: a block finding runs the repair turn", len(sc.Calls))
	}
	// The prompt must name the card, or the model can not keep it.
	if !strings.Contains(sc.Calls[0].Input, "Sol Ring") {
		t.Error("the first prompt did not name the locked card")
	}
}

// TestLockedCommanderCounts covers the card that became the commander. It
// is in the deck, in the command zone, so it is not missing (D-70).
func TestLockedCommanderCounts(t *testing.T) {
	one := step(t, deckOut{Summary: "s", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 1, Role: "synergy"},
	}})
	b, _, _ := testBuilder(t, one, one)
	req := testRequest()
	req.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	req.Commanders = []string{"o-karlov"}
	req.Locked = []string{"o-karlov"}
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeLockedCardMissing {
			t.Errorf("the commander was reported missing: %s", f.GetMessage())
		}
	}
}
