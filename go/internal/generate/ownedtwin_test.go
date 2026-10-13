package generate

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// twinCard is a card as Scryfall printed it on 2026-10-13.
func twinCard(id, name, cost, typeLine, text, power, toughness string) *mtgv1.Card {
	return &mtgv1.Card{OracleId: id, Name: name, ManaCost: cost, TypeLine: typeLine, OracleText: text, Power: power, Toughness: toughness, PriceUsd: 0.5}
}

var (
	twGrizzly   = twinCard("grizzly", "Grizzly Bears", "{1}{G}", "Creature — Bear", "", "2", "2")
	twBearCub   = twinCard("bearcub", "Bear Cub", "{1}{G}", "Creature — Bear", "", "2", "2")
	twBalduvian = twinCard("balduvian", "Balduvian Bears", "{1}{G}", "Creature — Bear", "", "2", "2")
	twRuneclaw  = twinCard("runeclaw", "Runeclaw Bear", "{1}{G}", "Creature — Bear", "", "2", "2")
	twLlanowar  = twinCard("llanowar", "Llanowar Elves", "{G}", "Creature — Elf Druid", "{T}: Add {G}.", "1", "1")
	twFyndhorn  = twinCard("fyndhorn", "Fyndhorn Elves", "{G}", "Creature — Elf Druid", "{T}: Add {G}.", "1", "1")
	twAlpine    = twinCard("alpine", "Alpine Grizzly", "{2}{G}", "Creature — Bear", "", "4", "2")
	twBeorn     = twinCard("beorn", "Beorn the Fierce", "{3}{G}{G}", "Legendary Creature — Bear Shapeshifter Warrior", "Trample", "5", "5")
)

// twinRequest is an owned-first request over the cards of F-232. The
// collection holds two Balduvian Bears, two Runeclaw Bears, and one
// Fyndhorn Elves, as the collection of the case does.
func twinRequest(format mtgv1.FormatId) Request {
	cards := []*mtgv1.Card{twGrizzly, twBearCub, twBalduvian, twRuneclaw, twLlanowar, twFyndhorn, twAlpine, twBeorn}
	owned := map[string]int32{"balduvian": 2, "runeclaw": 2, "fyndhorn": 1}
	return Request{
		Format:       format,
		Pool:         NewPool(cards, owned),
		PoolRule:     mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
		OracleCounts: owned,
	}
}

func twinDeck(cards ...*mtgv1.DeckCard) *mtgv1.Deck {
	return &mtgv1.Deck{Cards: cards, Summary: "Grizzly Bears and Bear Cub join Beorn the Fierce."}
}

func bought(c *mtgv1.Card, n int32) *mtgv1.DeckCard {
	return &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: n, Role: mtgv1.CardRole_CARD_ROLE_SYNERGY,
		Reason: c.GetName() + " grows the Bear count.", PriceUsd: c.GetPriceUsd()}
}

func countIn(deck *mtgv1.Deck, name string) int32 {
	var n int32
	for _, dc := range append(append([]*mtgv1.DeckCard(nil), deck.GetCards()...), deck.GetSideboard()...) {
		if dc.GetName() == name {
			n += dc.GetCount()
		}
	}
	return n
}

// TestSwapOwnedTwinsCommander is the case of F-232: the deck bought
// Grizzly Bears and Bear Cub, and the collection holds Balduvian Bears and
// Runeclaw Bear. A Commander deck holds one copy of each name, so each
// bought card takes a twin of its own.
func TestSwapOwnedTwinsCommander(t *testing.T) {
	req := twinRequest(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	deck := twinDeck(bought(twGrizzly, 1), bought(twBearCub, 1), bought(twAlpine, 1))
	lines := swapOwnedTwins(deck, req)
	if len(lines) != 2 {
		t.Fatalf("swaps = %v, want 2", lines)
	}
	for _, name := range []string{"Grizzly Bears", "Bear Cub"} {
		if n := countIn(deck, name); n != 0 {
			t.Errorf("the deck still buys %d %s while the collection holds a twin", n, name)
		}
	}
	for _, name := range []string{"Balduvian Bears", "Runeclaw Bear"} {
		if n := countIn(deck, name); n != 1 {
			t.Errorf("the deck holds %d %s, want 1", n, name)
		}
	}
	if n := countIn(deck, "Alpine Grizzly"); n != 1 {
		t.Error("a card with no owned twin left the deck")
	}
	for _, dc := range deck.GetCards() {
		if dc.GetName() == "Balduvian Bears" {
			if !dc.GetOwned() || dc.GetRole() != mtgv1.CardRole_CARD_ROLE_SYNERGY {
				t.Errorf("Balduvian Bears: owned %v, role %v, want owned and the role of the bought card", dc.GetOwned(), dc.GetRole())
			}
			if !strings.Contains(dc.GetReason(), "Balduvian Bears") {
				t.Errorf("the reason %q names the card that left", dc.GetReason())
			}
		}
	}
	if strings.Contains(deck.GetSummary(), "Grizzly Bears") || strings.Contains(deck.GetSummary(), "Bear Cub") {
		t.Errorf("the summary names a card the deck does not hold: %q", deck.GetSummary())
	}
}

// TestSwapOwnedTwinsElves is the class: base replays 2 to 4 bought
// Llanowar Elves, and the collection holds Fyndhorn Elves.
func TestSwapOwnedTwinsElves(t *testing.T) {
	req := twinRequest(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	deck := twinDeck(bought(twLlanowar, 1))
	swapOwnedTwins(deck, req)
	if countIn(deck, "Llanowar Elves") != 0 || countIn(deck, "Fyndhorn Elves") != 1 {
		t.Errorf("deck = %v, want Fyndhorn Elves in place of Llanowar Elves", deck.GetCards())
	}
}

// TestSwapOwnedTwinsKeepsAnOwnedCopy proves that a twin already in a
// Commander deck takes no second copy.
func TestSwapOwnedTwinsKeepsAnOwnedCopy(t *testing.T) {
	req := twinRequest(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	deck := twinDeck(bought(twFyndhorn, 1), bought(twLlanowar, 1))
	if lines := swapOwnedTwins(deck, req); len(lines) != 0 {
		t.Errorf("swaps = %v, want none: the only owned twin is in the deck", lines)
	}
	if countIn(deck, "Fyndhorn Elves") != 1 || countIn(deck, "Llanowar Elves") != 1 {
		t.Errorf("deck = %v, want both elves once", deck.GetCards())
	}
}

// TestSwapOwnedTwinsSixty proves the copy count in a 60-card deck: four
// bought Llanowar Elves and one owned Fyndhorn Elves swap one copy, and
// three stay on the buy list. Two owned Balduvian Bears take two of four
// Grizzly Bears.
func TestSwapOwnedTwinsSixty(t *testing.T) {
	req := twinRequest(mtgv1.FormatId_FORMAT_ID_STANDARD)
	deck := twinDeck(bought(twLlanowar, 4), bought(twGrizzly, 4))
	swapOwnedTwins(deck, req)
	want := map[string]int32{"Llanowar Elves": 3, "Fyndhorn Elves": 1, "Balduvian Bears": 2, "Runeclaw Bear": 2, "Grizzly Bears": 0}
	for name, n := range want {
		if got := countIn(deck, name); got != n {
			t.Errorf("%s: %d copies, want %d", name, got, n)
		}
	}
}

// TestSwapOwnedTwinsGuards proves the cards that stay: a card the user
// asked to keep, and every card of a session that may not prefer the
// collection.
func TestSwapOwnedTwinsGuards(t *testing.T) {
	t.Run("a locked card stays", func(t *testing.T) {
		req := twinRequest(mtgv1.FormatId_FORMAT_ID_COMMANDER)
		req.Locked = []string{"grizzly"}
		deck := twinDeck(bought(twGrizzly, 1))
		if lines := swapOwnedTwins(deck, req); len(lines) != 0 {
			t.Errorf("swaps = %v, want none", lines)
		}
	})
	t.Run("a card a revision keeps stays", func(t *testing.T) {
		req := twinRequest(mtgv1.FormatId_FORMAT_ID_COMMANDER)
		req.Revision = &Revision{Keep: []string{"grizzly bears"}}
		deck := twinDeck(bought(twGrizzly, 1))
		if lines := swapOwnedTwins(deck, req); len(lines) != 0 {
			t.Errorf("swaps = %v, want none", lines)
		}
	})
	t.Run("a precon card stays", func(t *testing.T) {
		req := twinRequest(mtgv1.FormatId_FORMAT_ID_COMMANDER)
		req.PreconOracleIDs = []string{"grizzly"}
		deck := twinDeck(bought(twGrizzly, 1))
		if lines := swapOwnedTwins(deck, req); len(lines) != 0 {
			t.Errorf("swaps = %v, want none", lines)
		}
	})
	for _, rule := range []mtgv1.PoolRule{mtgv1.PoolRule_POOL_RULE_ANY_CARD, mtgv1.PoolRule_POOL_RULE_UNSPECIFIED} {
		t.Run(rule.String(), func(t *testing.T) {
			req := twinRequest(mtgv1.FormatId_FORMAT_ID_COMMANDER)
			req.PoolRule = rule
			deck := twinDeck(bought(twGrizzly, 1))
			if lines := swapOwnedTwins(deck, req); len(lines) != 0 {
				t.Errorf("swaps = %v, want none: the reader asked for no collection first", lines)
			}
		})
	}
}

// TestBuildSwapsOwnedTwins proves that the build runs the pass and says
// so: the model returns four Grizzly Bears, and the deck plays the two
// owned Balduvian Bears and the two owned Runeclaw Bears in their place.
func TestBuildSwapsOwnedTwins(t *testing.T) {
	out := deckOut{Summary: "a Bear deck", Cards: []Entry{{Name: "Grizzly Bears", Count: 4, Role: "synergy", Reason: "a Bear"}}}
	b, src, _ := testBuilder(t, step(t, out), step(t, out))
	req := testRequest()
	tw := twinRequest(req.Format)
	req.Pool, req.PoolRule, req.OracleCounts = tw.Pool, tw.PoolRule, tw.OracleCounts
	for _, c := range []*mtgv1.Card{twGrizzly, twBalduvian, twRuneclaw} {
		src[c.GetOracleId()] = c
	}
	got, err := b.Build(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if n := countIn(got.Deck, "Grizzly Bears"); n != 0 {
		t.Errorf("the deck buys %d Grizzly Bears while the collection holds four twins", n)
	}
	var note string
	for _, f := range got.Deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeOwnedTwin {
			note = f.GetMessage()
		}
	}
	if !strings.Contains(note, "2 Balduvian Bears for Grizzly Bears") || !strings.Contains(note, "2 Runeclaw Bear for Grizzly Bears") {
		t.Errorf("the owned_twin finding reads %q, want both swaps", note)
	}
}
