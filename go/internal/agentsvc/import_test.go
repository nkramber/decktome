package agentsvc

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/users"
)

// ReadImport makes fakeDecks an Importer. It reads bracket 3 for a
// Commander list, or the floor alone when estimate is set, and it records
// one judge call when record is set.
func (f *fakeDecks) ReadImport(_ context.Context, deck *mtgv1.Deck, owned map[string]int32, acc *llm.Accumulator) {
	f.imports++
	f.owned = owned
	if deck.GetFormat().GetId() != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		return
	}
	if f.record && acc != nil {
		acc.Record(llm.RoleJudge, "fake-judge", &llm.Usage{InputTokens: 4000, OutputTokens: 200}, 0)
	}
	deck.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}}
	deck.BracketEstimated = f.estimate
	deck.Validation = &mtgv1.ValidationResult{Passed: true}
}

func importCardIndex() *cards.Index {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	wb := []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}
	w := []mtgv1.Color{mtgv1.Color_COLOR_W}
	return cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-karlov", Name: "Karlov of the Ghost Council", TypeLine: "Legendary Creature - Spirit Advisor",
			CanBeCommander: true, OracleText: karlovText, ColorIdentity: wb, Legalities: legal},
		{OracleId: "o-welcome", Name: "Ajani's Welcome", TypeLine: "Enchantment", OracleText: welcomeText, ColorIdentity: w, Legalities: legal},
		{OracleId: "o-plains", Name: "Plains", TypeLine: "Basic Land - Plains", Legalities: legal},
		{OracleId: "o-bolt", Name: "Lightning Bolt", TypeLine: "Instant", ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_R}, Legalities: legal},
	}, nil, nil, time.Unix(1000, 0).UTC())
}

func importServer(t *testing.T, fd *fakeDecks, ds *fakeDeckStore, noter *fakeNoter, steps ...llm.Step) (mtgv1connect.AgentServiceClient, *fakeStore) {
	t.Helper()
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	store := newFakeStore()
	opts := []Option{WithDecks(fd), WithCandidates(fixedIndex{importCardIndex()}, cb), WithDeckStore(ds), WithUsers(noter)}
	client, _ := testServerOpts(t, store, opts, steps...)
	return client, store
}

func importList(t *testing.T, c mtgv1connect.AgentServiceClient, req *mtgv1.ImportDeckRequest) *mtgv1.ImportDeckResponse {
	t.Helper()
	res, err := c.ImportDeck(context.Background(), connect.NewRequest(req))
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	return res.Msg
}

// markedList is a Commander list in the Archidekt shape of D-847: the
// category Commander{top} marks the commander.
const markedList = "1x Karlov of the Ghost Council (mkm) 1 [Commander{top}]\n1x Ajani's Welcome (m19) 1 [Lifegain]\n97x Plains (neo) 294 [Land]\n1x No Such Card (xyz) 1 [Draw]\n"

// TestImportStoresTheDeckAndASession is D-845, D-851, and D-852: the
// import keeps the list as a deck, writes a built session that names it,
// and counts one import on the user record.
func TestImportStoresTheDeckAndASession(t *testing.T) {
	fd, ds, noter := &fakeDecks{}, &fakeDeckStore{}, &fakeNoter{}
	client, store := importServer(t, fd, ds, noter)
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList, Name: "Living Weapon"})
	deck := res.GetDeck()
	if deck == nil || !deck.GetImported() || deck.GetName() != "Living Weapon" {
		t.Fatalf("deck = %v", deck)
	}
	if got := deck.GetCommanderOracleIds(); len(got) != 1 || got[0] != "o-karlov" {
		t.Errorf("commanders = %v", got)
	}
	if n := len(deck.GetCards()); n != 2 || deck.GetCards()[1].GetCount() != 97 {
		t.Errorf("cards = %v", deck.GetCards())
	}
	for _, c := range deck.GetCards() {
		if c.GetRole() == mtgv1.CardRole_CARD_ROLE_UNSPECIFIED {
			t.Errorf("%s has no role (D-848)", c.GetName())
		}
	}
	if len(res.GetUnresolved()) != 1 || !strings.Contains(res.GetUnresolved()[0].GetRaw(), "No Such Card") {
		t.Errorf("unresolved = %v", res.GetUnresolved())
	}
	if fd.imports != 1 || len(ds.put) != 1 || ds.put[0].GetId() != deck.GetId() {
		t.Errorf("reads = %d, stored = %v", fd.imports, ds.put)
	}
	s := store.sessions[res.GetSessionId()]
	if s == nil || s.GetStatus() != mtgv1.SessionStatus_SESSION_STATUS_BUILT || len(s.GetDeckIds()) != 1 || s.GetDeckIds()[0] != deck.GetId() {
		t.Fatalf("session = %v", s)
	}
	if deck.GetSessionId() != s.GetId() || s.GetSlots().GetPower().GetBracket() != 3 {
		t.Errorf("deck session = %q, slots = %v", deck.GetSessionId(), s.GetSlots())
	}
	want := []users.Counter{users.DecksImported, users.SessionsStarted}
	if fmt.Sprint(noter.counts) != fmt.Sprint(want) {
		t.Errorf("counters = %v, want %v", noter.counts, want)
	}
}

// TestImportThenReviseReadsTheImportedDeck is D-851: the first message on
// the session of an import revises the imported deck, as a message after
// a build does.
func TestImportThenReviseReadsTheImportedDeck(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	revised := &mtgv1.Deck{
		Name: "x", Summary: "a leaner deck", Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 2}, {OracleId: "o-plains", Name: "Plains", Count: 96}},
	}
	fd.res = &generate.Result{Deck: revised}
	client, store := importServer(t, fd, ds, &fakeNoter{},
		classifyJSON(t, nil),
		reviseJSON(t, map[string]any{"changes": []string{"Cut one Plains and add a card that draws cards"}}))
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList})
	ev := chat(t, client, &mtgv1.ChatRequest{SessionId: res.GetSessionId(), Message: "Add a card that draws cards."})
	if fd.runs != 1 || fd.got.Revision == nil {
		t.Fatalf("runs = %d, revision = %v, order = %v, failure = %v", fd.runs, fd.got.Revision, ev.order, ev.failure)
	}
	if fd.got.Revision.BaseDeckID != res.GetDeck().GetId() {
		t.Errorf("base = %q, want %q", fd.got.Revision.BaseDeckID, res.GetDeck().GetId())
	}
	if ev.deck == nil || ev.deck.GetRevisedFromDeckId() != res.GetDeck().GetId() || ev.deck.GetImported() {
		t.Errorf("revised deck = %v", ev.deck)
	}
	if n := len(store.sessions[res.GetSessionId()].GetDeckIds()); n != 2 {
		t.Errorf("deck ids = %d, want 2", n)
	}
}

// TestImportAsksForTheCommander is D-847: a 100-card list with no mark
// offers the cards that can lead, and the pick stores the deck.
func TestImportAsksForTheCommander(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	client, _ := importServer(t, fd, ds, &fakeNoter{})
	list := "1 Karlov of the Ghost Council\n1 Ajani's Welcome\n98 Plains\n"
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: list})
	if res.GetDeck() != nil || len(res.GetCommanderOptions()) != 1 || res.GetCommanderOptions()[0].GetOracleId() != "o-karlov" {
		t.Fatalf("response = %v", res)
	}
	if len(ds.put) != 0 {
		t.Errorf("a question stored a deck: %v", ds.put)
	}
	res = importList(t, client, &mtgv1.ImportDeckRequest{Text: list, CommanderOracleIds: []string{"o-karlov"}})
	deck := res.GetDeck()
	if deck == nil || deck.GetCommanderOracleIds()[0] != "o-karlov" {
		t.Fatalf("deck = %v", deck)
	}
	for _, c := range deck.GetCards() {
		if c.GetOracleId() == "o-karlov" {
			t.Errorf("the commander stayed in the 99")
		}
	}
	_, err := client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: list, CommanderOracleIds: []string{"o-welcome"}}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("a card that can not lead was taken: %v", err)
	}
}

// TestImportAsksForTheFormat is D-857: a list that is not Commander asks
// Standard, Modern, or neither, and neither stores the house format.
func TestImportAsksForTheFormat(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	client, store := importServer(t, fd, ds, &fakeNoter{})
	list := "Deck\n4 Lightning Bolt\n56 Plains\n\nSideboard\n2 Ajani's Welcome\n"
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: list})
	if !res.GetNeedsFormat() || res.GetDeck() != nil {
		t.Fatalf("response = %v", res)
	}
	res = importList(t, client, &mtgv1.ImportDeckRequest{Text: list, Format: mtgv1.FormatId_FORMAT_ID_HOUSE})
	deck := res.GetDeck()
	if deck.GetFormat().GetId() != mtgv1.FormatId_FORMAT_ID_HOUSE || len(deck.GetSideboard()) != 1 || deck.GetPower() != nil {
		t.Fatalf("deck = %v", deck)
	}
	slots := store.sessions[res.GetSessionId()].GetSlots()
	// The sideboard holds the white card, so the main deck reads red.
	if got := slots.GetColors(); len(got) != 1 || got[0] != mtgv1.Color_COLOR_R {
		t.Errorf("colors = %v, want red", got)
	}
}

// TestImportReadsThePickedCollection is D-849.
func TestImportReadsThePickedCollection(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	col := &fakeCollections{counts: map[string]int32{"o-welcome": 1}}
	opts := []Option{WithDecks(fd), WithCandidates(fixedIndex{importCardIndex()}, cb), WithDeckStore(ds), WithCollections(col)}
	client, _ := testServerOpts(t, newFakeStore(), opts)
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList, CollectionId: "c1"})
	if fd.owned["o-welcome"] != 1 {
		t.Errorf("owned = %v", fd.owned)
	}
	if res.GetDeck() == nil {
		t.Fatal("no deck")
	}
}

// TestReadImportBracketAsksAgain is D-854: an estimated bracket asks the
// judge again, and a judged one comes back as it is.
func TestReadImportBracketAsksAgain(t *testing.T) {
	fd, ds := &fakeDecks{estimate: true}, &fakeDeckStore{}
	client, _ := importServer(t, fd, ds, &fakeNoter{})
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList})
	if !res.GetDeck().GetBracketEstimated() {
		t.Fatal("the fake set no estimate")
	}
	fd.estimate = false
	got, err := client.ReadImportBracket(context.Background(), connect.NewRequest(&mtgv1.ReadImportBracketRequest{DeckId: res.GetDeck().GetId()}))
	if err != nil {
		t.Fatal(err)
	}
	if got.Msg.GetDeck().GetBracketEstimated() || fd.imports != 2 {
		t.Errorf("estimated = %v, reads = %d", got.Msg.GetDeck().GetBracketEstimated(), fd.imports)
	}
	if _, err := client.ReadImportBracket(context.Background(), connect.NewRequest(&mtgv1.ReadImportBracketRequest{DeckId: res.GetDeck().GetId()})); err != nil || fd.imports != 2 {
		t.Errorf("a judged bracket was read again: reads = %d, err = %v", fd.imports, err)
	}
}

// TestImportRefusesALongList is P2-1 of the review of #220: a list over
// the line cap stores no deck and no session (D-846).
func TestImportRefusesALongList(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	client, store := importServer(t, fd, ds, &fakeNoter{})
	text := "1 Karlov of the Ghost Council\n" + strings.Repeat("1 Plains\n", 1000)
	_, err := client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: text}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Fatalf("err = %v, want InvalidArgument", err)
	}
	if len(ds.put) != 0 || len(store.sessions) != 0 || fd.imports != 0 {
		t.Errorf("a refused list stored %d decks and %d sessions", len(ds.put), len(store.sessions))
	}
}
