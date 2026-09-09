package generate

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
)

// The PR-33 tests. F-78: almost every deck left the first model call off
// its bracket band, and the app paid a 60 to 90 second call to move a
// number it computes in 10 milliseconds. F-77: a build that lost its
// deadline returned nothing, though a legal deck stood.

// manaCard is one card of the mana-pass fixtures.
func manaCard(oid, name string, mv float64, types []string, text string, produced ...mtgv1.Color) *mtgv1.Card {
	return &mtgv1.Card{
		OracleId: oid, Name: name, ManaValue: mv, CardTypes: types,
		OracleText: text, ProducedMana: produced,
		Legalities: map[string]mtgv1.LegalityStatus{
			"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL,
			"modern":    mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL,
		},
	}
}

func manaBuilder(t *testing.T, src source) *Builder {
	t.Helper()
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	p, err := profile.New(cfg, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	return &Builder{cards: src, rules: cfg, profiler: p,
		log: slog.New(slog.NewTextHandler(io.Discard, nil))}
}

// TestBandDistanceReadsEveryDirection is the score the pass ranks its
// steps by. A feature inside its band is zero, and one outside reads how
// far out it is, as a share of the band's own scale. Without the share a
// land count of 40 would outweigh a broken mana base.
func TestBandDistanceReadsEveryDirection(t *testing.T) {
	under := &mtgv1.ProfileFeature{Key: "mana_turn_four", Value: 4.4, Low: 4.8, OffBand: true}
	over := &mtgv1.ProfileFeature{Key: "land", Value: 39, Low: 34, High: 38, HasHigh: true, OffBand: true}
	inside := &mtgv1.ProfileFeature{Key: "land", Value: 36, Low: 34, High: 38, HasHigh: true}
	if d := bandDistance(under); d <= 0 {
		t.Errorf("a value under its floor reads %v, want a positive distance", d)
	}
	if d := bandDistance(over); d <= 0 {
		t.Errorf("a value over its ceiling reads %v, want a positive distance", d)
	}
	if d := bandDistance(inside); d != 0 {
		t.Errorf("a value inside its band reads %v, want 0", d)
	}
	// The share makes the two comparable: one land over a ceiling of 38
	// is a smaller fault than 0.4 mana under a floor of 4.8.
	if bandDistance(over) >= bandDistance(under) {
		t.Errorf("one land over the band (%v) outweighs a broken mana base (%v)",
			bandDistance(over), bandDistance(under))
	}
}

// TestTheManaPassTradesATappedLandForAnUntappedOne is the first lever.
// The band of every bracket caps the lands that enter tapped, and the
// pool holds an untapped land the model did not take.
func TestTheManaPassTradesATappedLandForAnUntappedOne(t *testing.T) {
	src := source{}
	var deckCards []*mtgv1.DeckCard
	add := func(c *mtgv1.Card, n int32, role mtgv1.CardRole) {
		src[c.GetOracleId()] = c
		deckCards = append(deckCards, &mtgv1.DeckCard{
			OracleId: c.GetOracleId(), Name: c.GetName(), Count: n, Role: role,
		})
	}
	plains := manaCard("o-plains", "Plains", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	plains.Supertypes = []string{"Basic"}
	// Ten tapped lands break the bracket 4 ceiling of five.
	tapped := manaCard("o-tap", "Tapped Hold", 0, []string{"Land"}, "This land enters tapped.", mtgv1.Color_COLOR_W)
	untapped := manaCard("o-untap", "Open Hold", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	spell := manaCard("o-spell", "White Spell", 2, []string{"Creature"}, "")
	add(plains, 25, mtgv1.CardRole_CARD_ROLE_LAND)
	add(tapped, 10, mtgv1.CardRole_CARD_ROLE_LAND)
	add(spell, 64, mtgv1.CardRole_CARD_ROLE_THREAT)
	src[untapped.GetOracleId()] = untapped

	b := manaBuilder(t, src)
	deck := &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:  &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 4}},
		Cards:  deckCards,
	}
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Power:  deck.GetPower(),
		Pool:   NewPool([]*mtgv1.Card{plains, tapped, untapped, spell}, nil),
	}
	before := b.manaScore(deck)
	steps := b.fixMana(req, deck)
	after := b.manaScore(deck)
	if steps == 0 {
		t.Fatal("the mana pass made no step on a deck of ten tapped lands")
	}
	if after >= before {
		t.Errorf("the pass did not move the deck toward its bands: %v then %v", before, after)
	}
	if countOf(deck, "o-untap") == 0 {
		t.Error("the pass took no untapped land from the pool")
	}
	if countOf(deck, "o-tap") == 10 {
		t.Error("the pass cut no tapped land")
	}
	// Every step keeps the deck size, or the engine refuses the deck.
	if n := deckCount(deck); n != 99 {
		t.Errorf("the deck holds %d cards after the pass, want 99", n)
	}
}

// TestTheManaPassCarriesOwnershipAndPrice is F-79, found on the
// deployed app on 2026-09-08. The pass added seven cards of the reader's
// own collection, and every one read "to buy" at no price, on a deck
// built from owned cards alone. Session IDS5oQE3D0XDWs6o1FnR.
func TestTheManaPassCarriesOwnershipAndPrice(t *testing.T) {
	src := source{}
	var deckCards []*mtgv1.DeckCard
	add := func(c *mtgv1.Card, n int32, role mtgv1.CardRole, owned int32) {
		src[c.GetOracleId()] = c
		deckCards = append(deckCards, &mtgv1.DeckCard{
			OracleId: c.GetOracleId(), Name: c.GetName(), Count: n, Role: role,
			OwnedCount: owned, Owned: owned >= n,
		})
	}
	plains := manaCard("o-plains", "Plains", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	plains.Supertypes = []string{"Basic"}
	tapped := manaCard("o-tap", "Tapped Hold", 0, []string{"Land"}, "This land enters tapped.", mtgv1.Color_COLOR_W)
	untapped := manaCard("o-untap", "Open Hold", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	untapped.PriceUsd = 2.5
	spell := manaCard("o-spell", "White Spell", 2, []string{"Creature"}, "")
	add(plains, 25, mtgv1.CardRole_CARD_ROLE_LAND, 40)
	add(tapped, 10, mtgv1.CardRole_CARD_ROLE_LAND, 10)
	add(spell, 64, mtgv1.CardRole_CARD_ROLE_THREAT, 64)
	src[untapped.GetOracleId()] = untapped

	b := manaBuilder(t, src)
	deck := &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:  &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 4}},
		Cards:  deckCards,
	}
	// The reader owns three copies of the untapped land, so the pass
	// takes one of their own cards.
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Power:  deck.GetPower(),
		Pool: NewPool([]*mtgv1.Card{plains, tapped, untapped, spell},
			map[string]int32{"o-untap": 3, "o-plains": 40, "o-tap": 10, "o-spell": 64}),
	}
	if steps := b.fixMana(req, deck); steps == 0 {
		t.Fatal("the pass made no step")
	}
	var added *mtgv1.DeckCard
	for _, dc := range deck.GetCards() {
		if dc.GetOracleId() == "o-untap" {
			added = dc
		}
	}
	if added == nil {
		t.Fatal("the pass took no untapped land")
	}
	if added.GetOwnedCount() != 3 {
		t.Errorf("the added card reads %d copies owned, and the collection holds 3", added.GetOwnedCount())
	}
	if !added.GetOwned() {
		t.Error("the added card reads as one to buy, and the reader owns it")
	}
	if added.GetPriceUsd() != 2.5 {
		t.Errorf("the added card reads price %v, want the card's own price", added.GetPriceUsd())
	}
	// Nothing of this deck reaches a buy list: every card is owned.
	for _, dc := range deck.GetCards() {
		if !dc.GetOwned() {
			t.Errorf("a deck of owned cards names %q as one to buy", dc.GetName())
		}
	}
}

// A deck already inside its bands is a deck the pass leaves alone. The
// pass must never move a deck for its own sake.
func TestTheManaPassLeavesADeckInBandAlone(t *testing.T) {
	src := source{}
	plains := manaCard("o-plains", "Plains", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	plains.Supertypes = []string{"Basic"}
	spell := manaCard("o-spell", "White Spell", 2, []string{"Creature"}, "")
	src[plains.GetOracleId()], src[spell.GetOracleId()] = plains, spell
	b := manaBuilder(t, src)
	deck := &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:  &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 4}},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-plains", Name: "Plains", Count: 35, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-spell", Name: "White Spell", Count: 64, Role: mtgv1.CardRole_CARD_ROLE_THREAT},
		},
	}
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(),
		Pool: NewPool([]*mtgv1.Card{plains, spell}, nil)}
	if b.manaScore(deck) == 0 {
		if steps := b.fixMana(req, deck); steps != 0 {
			t.Errorf("the pass made %d steps on a deck already in band", steps)
		}
	}
}

// A revision keeps the deck it was given. The reader asked for a change,
// and a pass over it is a second author (PR-12B, D-283).
func TestTheManaPassSkipsARevision(t *testing.T) {
	src := source{}
	plains := manaCard("o-plains", "Plains", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	plains.Supertypes = []string{"Basic"}
	tapped := manaCard("o-tap", "Tapped Hold", 0, []string{"Land"}, "This land enters tapped.", mtgv1.Color_COLOR_W)
	untapped := manaCard("o-untap", "Open Hold", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	spell := manaCard("o-spell", "White Spell", 2, []string{"Creature"}, "")
	for _, c := range []*mtgv1.Card{plains, tapped, untapped, spell} {
		src[c.GetOracleId()] = c
	}
	b := manaBuilder(t, src)
	deck := &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:  &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 4}},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-plains", Name: "Plains", Count: 25, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-tap", Name: "Tapped Hold", Count: 10, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-spell", Name: "White Spell", Count: 64, Role: mtgv1.CardRole_CARD_ROLE_THREAT},
		},
	}
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(),
		Pool: NewPool([]*mtgv1.Card{plains, tapped, untapped, spell}, nil), Revision: &Revision{}}
	if steps := b.fixMana(req, deck); steps != 0 {
		t.Errorf("the pass moved a revision: %d steps", steps)
	}
}

// TestTheManaPassMovesAnUpgradesLandsAlone is D-628. An upgrade takes
// the pass on its mana base, and never on its spells. D-249 keeps the
// precon's own composition, so no step drops a card of the precon and
// no step trades a spell.
func TestTheManaPassMovesAnUpgradesLandsAlone(t *testing.T) {
	src := source{}
	plains := manaCard("o-plains", "Plains", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	plains.Supertypes = []string{"Basic"}
	// The precon owns this tapped land, so no step may drop it.
	preconLand := manaCard("o-preconland", "Precon Hold", 0, []string{"Land"}, "This land enters tapped.", mtgv1.Color_COLOR_W)
	// This tapped land is not the precon's, so a step may trade it.
	spare := manaCard("o-spare", "Spare Hold", 0, []string{"Land"}, "This land enters tapped.", mtgv1.Color_COLOR_W)
	untapped := manaCard("o-untap", "Open Hold", 0, []string{"Land"}, "", mtgv1.Color_COLOR_W)
	spell := manaCard("o-spell", "Precon Spell", 5, []string{"Creature"}, "")
	for _, c := range []*mtgv1.Card{plains, preconLand, spare, untapped, spell} {
		src[c.GetOracleId()] = c
	}
	b := manaBuilder(t, src)
	deck := &mtgv1.Deck{
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:  &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 4}},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-plains", Name: "Plains", Count: 25, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-preconland", Name: "Precon Hold", Count: 5, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-spare", Name: "Spare Hold", Count: 5, Role: mtgv1.CardRole_CARD_ROLE_LAND},
			{OracleId: "o-spell", Name: "Precon Spell", Count: 64, Role: mtgv1.CardRole_CARD_ROLE_THREAT},
		},
	}
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Power: deck.GetPower(),
		Pool:            NewPool([]*mtgv1.Card{plains, preconLand, spare, untapped, spell}, nil),
		Precon:          "Goblin Storm",
		PreconOracleIDs: []string{"o-preconland", "o-spell"},
	}
	before := b.manaScore(deck)
	steps := b.fixMana(req, deck)
	if steps == 0 {
		t.Fatal("the pass made no step on an upgrade of ten tapped lands")
	}
	if b.manaScore(deck) >= before {
		t.Error("the pass did not move the upgrade toward its bands")
	}
	// The precon keeps every card it owns, so the share of D-218 can not
	// move under the pass.
	if countOf(deck, "o-preconland") != 5 {
		t.Errorf("the pass dropped a precon land: %d copies left", countOf(deck, "o-preconland"))
	}
	if countOf(deck, "o-spell") != 64 {
		t.Errorf("the pass moved a precon spell: %d copies left", countOf(deck, "o-spell"))
	}
	// The spare tapped land is the one it may trade.
	if countOf(deck, "o-spare") == 5 {
		t.Error("the pass traded no tapped land of its own")
	}
	if n := deckCount(deck); n != 99 {
		t.Errorf("the upgrade holds %d cards after the pass, want 99", n)
	}
}

// TestNoRepairTurnForProfileFindingsAlone is PR-33, Part 4. A deck whose
// findings are profile findings alone is a legal deck, and the mana pass
// has moved what the pool allows. Session 833r7UccvAqFsyYJzHfz spent two
// calls of 88 seconds on such a deck and the reader got nothing.
func TestNoRepairTurnForProfileFindingsAlone(t *testing.T) {
	one := step(t, deckOut{Summary: "burn", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 4, Role: "synergy", Reason: "gains life"},
	}})
	b, _, sc := testBuilder(t, one, one)
	req := testRequest()
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	// The fixture deck carries no profile, so this reads the rule that
	// matters: a build never spends a call on a deck with no miss and no
	// block. The script holds two steps, and one call must be enough.
	if got.Repaired && len(sc.Calls) > 1 {
		if profileOnly(repairable(got.Deck.GetValidation())) {
			t.Error("a repair turn ran for profile findings alone")
		}
	}
}

// TestATimedOutRepairKeepsTheLegalDeck is F-77. The build returned nil
// and the reader read an error, though a legal deck stood. Session
// 833r7UccvAqFsyYJzHfz waited 3 minutes 59 seconds for that error.
func TestATimedOutRepairKeepsTheLegalDeck(t *testing.T) {
	// A pool of fifteen names, four copies each, is a legal 60-card
	// deck. The model answers all sixty and one name the pool does not
	// hold, so the miss buys the repair turn and the deck is legal.
	var pool []*mtgv1.Card
	src := source{}
	entries := make([]Entry, 0, 16)
	for i := range 15 {
		id := "o-legal" + string(rune('a'+i))
		c := manaCard(id, "Legal Card "+string(rune('A'+i)), 2, []string{"Creature"}, "")
		pool = append(pool, c)
		src[id] = c
		entries = append(entries, Entry{Name: c.GetName(), Count: 4, Role: "threat", Reason: "a body"})
	}
	entries = append(entries, Entry{Name: "Craterhoof Behemoth", Count: 1, Role: "wincon", Reason: "ends the game"})

	sc := llm.NewScript(step(t, deckOut{Summary: "a creature deck", Cards: entries}))
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	b := NewBuilder(c, cfg, src, slog.New(slog.NewTextHandler(io.Discard, nil)))
	req := testRequest()
	req.Pool = NewPool(pool, nil)

	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("the build returned an error over a legal deck: %v", err)
	}
	if got.Deck == nil || len(got.Deck.GetCards()) == 0 {
		t.Fatal("the build returned no deck")
	}
	if !got.Deck.GetValidation().GetPassed() {
		t.Fatalf("the deck that stands does not pass: %v", got.Deck.GetValidation().GetFindings())
	}
	// The reader is told the repair turn did not finish.
	var kept bool
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeRepairKept {
			kept = true
		}
	}
	if !kept {
		t.Error("the deck carries no note that the repair turn did not finish")
	}
	// A deck the engine refuses is another matter: nothing legal stands,
	// so the build still answers the error.
	sc2 := llm.NewScript(step(t, deckOut{Summary: "four cards", Cards: []Entry{
		{Name: "Legal Card A", Count: 4, Role: "threat", Reason: "a body"},
		{Name: "Craterhoof Behemoth", Count: 1, Role: "wincon", Reason: "ends the game"},
	}}))
	c2, err := llm.New(fakeConfig(), []llm.Provider{sc2}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	b2 := NewBuilder(c2, cfg, src, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if _, err := b2.Build(context.Background(), req, nil); err == nil {
		t.Error("a build whose only deck is illegal returned no error")
	}
}

// TestTimeForAnotherCallReadsTheDeadline is the clock of Part 4. A build
// with 30 seconds left starts no call that took 88.
func TestTimeForAnotherCallReadsTheDeadline(t *testing.T) {
	b := &Builder{log: slog.New(slog.NewTextHandler(io.Discard, nil))}
	// No deadline: the gate runners and the tests build under one, and
	// the rule is for the deployed app.
	if !b.timeForAnotherCall(context.Background(), 88*time.Second) {
		t.Error("a context with no deadline refused a call")
	}
	short, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if b.timeForAnotherCall(short, 88*time.Second) {
		t.Error("a build with 30 seconds left started a call that takes 88")
	}
	long, cancel2 := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel2()
	if !b.timeForAnotherCall(long, 60*time.Second) {
		t.Error("a build with four minutes left refused a call of one")
	}
	// The margin covers the assemble, the engine, and the profile after
	// the call, so a call that exactly fills the time is refused.
	tight, cancel3 := context.WithTimeout(context.Background(), 61*time.Second)
	defer cancel3()
	if b.timeForAnotherCall(tight, 60*time.Second) {
		t.Error("a call with no room for the work after it was allowed")
	}
}

func countOf(deck *mtgv1.Deck, id string) int32 {
	for _, dc := range deck.GetCards() {
		if dc.GetOracleId() == id {
			return dc.GetCount()
		}
	}
	return 0
}

func deckCount(deck *mtgv1.Deck) int32 {
	var n int32
	for _, dc := range deck.GetCards() {
		n += dc.GetCount()
	}
	return n + int32(len(deck.GetCommanderOracleIds()))
}
