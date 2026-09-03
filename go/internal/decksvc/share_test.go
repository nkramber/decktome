package decksvc

import (
	"context"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

func sharedDeckFixture() *mtgv1.Deck {
	return &mtgv1.Deck{
		Id: "d1", Name: "Elf test", SessionId: "s-secret", Favorite: true,
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN}, Summary: "A small deck.",
		LegalityAsOf: "2026-09-03", CardCount: 5,
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-elf", Name: "Llanowar Elves", Count: 4, Role: mtgv1.CardRole_CARD_ROLE_RAMP, Reason: "Turn-one mana.", Owned: true, OwnedCount: 4,
				OwnedPrinting: &mtgv1.Printing{ScryfallId: "p-alpha", PriceUsd: 40}},
			{OracleId: "o-forest", Name: "Forest", Count: 1, Role: mtgv1.CardRole_CARD_ROLE_LAND, Owned: true, OwnedCount: 20},
		},
		Validation: &mtgv1.ValidationResult{Passed: true},
	}
}

func shareServer(t *testing.T, f *fakeDecks, uid string) *Server {
	t.Helper()
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-elf", Name: "Llanowar Elves", TypeLine: "Creature — Elf Druid", DefaultPrinting: &mtgv1.Printing{ScryfallId: "p-elf", SetCode: "m19", CollectorNumber: "314", Artist: "Anson Maddocks"}},
	}, nil, nil, time.Date(2026, 9, 3, 9, 2, 27, 0, time.UTC))
	opts := []Option{WithDecks(f)}
	if uid != "" {
		opts = append(opts, asUser(uid))
	}
	return New(nil, fixedIndex{idx}, opts...)
}

// TestShareRevokeAndRead is the PR-21 gate: a share makes a token, the
// token reads the deck with no sign-in, and a revoked link is NotFound.
func TestShareRevokeAndRead(t *testing.T) {
	ctx := context.Background()
	f := &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": sharedDeckFixture()}}
	owner := shareServer(t, f, "u1")
	res, err := owner.ShareDeck(ctx, connect.NewRequest(&mtgv1.ShareDeckRequest{DeckId: "d1"}))
	if err != nil {
		t.Fatal(err)
	}
	token := res.Msg.Token
	if len(token) != 43 || strings.ContainsAny(token, "+/=") {
		t.Fatalf("token = %q, want 43 URL-safe characters", token)
	}
	if f.shares[hashToken(token)] != "u1/d1" || f.decks["d1"].Shared != true {
		t.Fatalf("the store holds %v, and the deck reads shared %v", f.shares, f.decks["d1"].Shared)
	}
	// Anyone with the token reads the deck, with no user.
	visitor := shareServer(t, f, "")
	got, err := visitor.GetSharedDeck(ctx, connect.NewRequest(&mtgv1.GetSharedDeckRequest{Token: token}))
	if err != nil {
		t.Fatal(err)
	}
	d := got.Msg.Deck
	if d.Name != "Elf test" || d.Summary != "A small deck." || d.CardCount != 5 || len(d.Cards) != 2 {
		t.Errorf("shared deck = %+v", d)
	}
	if d.Cards[0].Card == nil || d.Cards[0].Card.Name != "Llanowar Elves" || d.Cards[0].Reason != "Turn-one mana." {
		t.Errorf("the card data rides along: %+v", d.Cards[0])
	}
	if d.Cards[1].Card != nil {
		t.Error("a card the index lacks carries no data")
	}
	// The export names the default paper printing, never the owned one.
	text, err := visitor.ExportSharedDeck(ctx, connect.NewRequest(&mtgv1.ExportSharedDeckRequest{Token: token}))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text.Msg.Text, "4 Llanowar Elves (M19) 314") || strings.Contains(text.Msg.Text, "p-alpha") {
		t.Errorf("export = %q", text.Msg.Text)
	}
	// A second share replaces the link.
	again, err := owner.ShareDeck(ctx, connect.NewRequest(&mtgv1.ShareDeckRequest{DeckId: "d1"}))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := visitor.GetSharedDeck(ctx, connect.NewRequest(&mtgv1.GetSharedDeckRequest{Token: token})); connect.CodeOf(err) != connect.CodeNotFound {
		t.Errorf("the old link still opens: %v", err)
	}
	// Revoke ends it.
	if _, err := owner.RevokeShare(ctx, connect.NewRequest(&mtgv1.RevokeShareRequest{DeckId: "d1"})); err != nil {
		t.Fatal(err)
	}
	if _, err := visitor.GetSharedDeck(ctx, connect.NewRequest(&mtgv1.GetSharedDeckRequest{Token: again.Msg.Token})); connect.CodeOf(err) != connect.CodeNotFound {
		t.Errorf("a revoked link answers %v, want NotFound", err)
	}
	if f.decks["d1"].Shared {
		t.Error("the deck still reads shared")
	}
	// A token of the wrong shape, and an unknown one, read NotFound too.
	for _, bad := range []string{"short", strings.Repeat("a", 43)} {
		if _, err := visitor.GetSharedDeck(ctx, connect.NewRequest(&mtgv1.GetSharedDeckRequest{Token: bad})); connect.CodeOf(err) != connect.CodeNotFound {
			t.Errorf("token %q: %v, want NotFound", bad, err)
		}
	}
	if _, err := visitor.GetSharedDeck(ctx, connect.NewRequest(&mtgv1.GetSharedDeckRequest{})); connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("no token: %v, want InvalidArgument", err)
	}
	// The writes need a user.
	if _, err := visitor.ShareDeck(ctx, connect.NewRequest(&mtgv1.ShareDeckRequest{DeckId: "d1"})); connect.CodeOf(err) != connect.CodeUnauthenticated {
		t.Errorf("share with no user: %v", err)
	}
	if _, err := owner.ShareDeck(ctx, connect.NewRequest(&mtgv1.ShareDeckRequest{DeckId: "d-none"})); connect.CodeOf(err) != connect.CodeNotFound {
		t.Errorf("share of an unknown deck: %v", err)
	}
}

// TestSharedDeckHoldsNoUserField is guardrail 13: the proto text of
// SharedDeck and SharedCard names no session, collection, owned mark,
// owned printing, user, or token field.
func TestSharedDeckHoldsNoUserField(t *testing.T) {
	raw, err := os.ReadFile("../../../proto/mtg/v1/deck.proto")
	if err != nil {
		t.Skipf("proto source not readable: %v", err)
	}
	text := string(raw)
	fieldRe := regexp.MustCompile(`(?m)^\s+(?:repeated\s+)?[A-Za-z0-9_.]+\s+([a-z_]+)\s*=\s*\d+;`)
	forbidden := regexp.MustCompile(`session|collection|owned|user|uid|token|favorite|revised`)
	var seen int
	for _, name := range []string{"SharedDeck", "SharedCard"} {
		i := strings.Index(text, "message "+name+" {")
		if i < 0 {
			t.Fatalf("deck.proto holds no message %s", name)
		}
		j := strings.Index(text[i:], "\n}")
		for _, m := range fieldRe.FindAllStringSubmatch(text[i:i+j], -1) {
			seen++
			if forbidden.MatchString(m[1]) {
				t.Errorf("%s.%s is a user field on a public message", name, m[1])
			}
		}
	}
	if seen < 10 {
		t.Fatalf("read %d fields, want the two messages whole", seen)
	}
	// The public export clears every user field of the deck copy.
	d := publicDeck(sharedDeckFixture())
	if d.SessionId != "" || d.Favorite || d.Validation != nil || d.Cards[0].Owned || d.Cards[0].OwnedCount != 0 || d.Cards[0].OwnedPrinting != nil {
		t.Errorf("publicDeck kept a user field: %+v", d)
	}
}
