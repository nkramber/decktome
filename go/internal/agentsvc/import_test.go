package agentsvc

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/importfault"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/users"
)

// ReadImport makes fakeDecks an Importer. It reads bracket 3 for a
// Commander list, or the floor 2 alone when estimate is set, and it
// records one judge call when record is set. A 60-card list reads the
// step of the field step, and no step when it is unset.
func (f *fakeDecks) ReadImport(_ context.Context, deck *mtgv1.Deck, owned map[string]int32, acc *llm.Accumulator) {
	f.imports++
	f.owned = owned
	if f.during != nil {
		f.during()
	}
	if deck.GetFormat().GetId() != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		deck.Power = nil
		if f.step != mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED {
			deck.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: f.step}}
		}
		return
	}
	if f.record && acc != nil {
		acc.Record(llm.RoleJudge, "fake-judge", &llm.Usage{InputTokens: 4000, OutputTokens: 200}, 0)
	}
	bracket := int32(3)
	if f.estimate {
		bracket = 2
	}
	deck.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bracket}}
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

// TestImportTakesAPickedCommanderFormat is REV-027 of the review of
// 2026-09-24. A Commander list of 99 cards reads no format of its own.
// A pick of Commander asks for the leader, and the deck stores as
// Commander.
func TestImportTakesAPickedCommanderFormat(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	client, _ := importServer(t, fd, ds, &fakeNoter{})
	list := "1 Karlov of the Ghost Council\n1 Ajani's Welcome\n97 Plains\n"
	if res := importList(t, client, &mtgv1.ImportDeckRequest{Text: list}); !res.GetNeedsFormat() {
		t.Fatalf("a list of 99 cards read a format of its own: %v", res)
	}
	commander := mtgv1.FormatId_FORMAT_ID_COMMANDER
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: list, Format: commander})
	if len(res.GetCommanderOptions()) != 1 || res.GetCommanderOptions()[0].GetOracleId() != "o-karlov" {
		t.Fatalf("response = %v, want Karlov as the one leader", res)
	}
	res = importList(t, client, &mtgv1.ImportDeckRequest{Text: list, Format: commander, CommanderOracleIds: []string{"o-karlov"}})
	if d := res.GetDeck(); d.GetFormat().GetId() != commander || d.GetCommanderOracleIds()[0] != "o-karlov" {
		t.Errorf("deck format %v, commanders %v, want Commander led by Karlov", d.GetFormat().GetId(), d.GetCommanderOracleIds())
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

// TestReadImportBracketLeavesADeletedDeckDeleted is REV-007 of the
// review of 2026-09-24. The new read rewrites the stored deck. A deck
// that the user deleted during the judge call stays deleted, so its share
// link does not come back with it.
func TestReadImportBracketLeavesADeletedDeckDeleted(t *testing.T) {
	fd, ds := &fakeDecks{estimate: true}, &fakeDeckStore{}
	client, _ := importServer(t, fd, ds, &fakeNoter{})
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList})
	fd.estimate = false
	fd.during = func() { ds.put = nil }
	_, err := client.ReadImportBracket(context.Background(), connect.NewRequest(&mtgv1.ReadImportBracketRequest{DeckId: res.GetDeck().GetId()}))
	if connect.CodeOf(err) != connect.CodeNotFound {
		t.Errorf("err = %v, want NotFound", err)
	}
	if len(ds.put) != 0 {
		t.Errorf("the deleted deck came back: %d decks", len(ds.put))
	}
}

// TestReadImportBracketWritesTheSessionPower is F-172 and D-870: a new
// read that raises the floor also writes the bracket into the power slot
// of the session, because a revise turn reads the slot.
func TestReadImportBracketWritesTheSessionPower(t *testing.T) {
	fd, ds := &fakeDecks{estimate: true}, &fakeDeckStore{}
	client, store := importServer(t, fd, ds, &fakeNoter{})
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList})
	if got := store.sessions[res.GetSessionId()].GetSlots().GetPower().GetBracket(); got != 2 {
		t.Fatalf("session bracket = %d, want the floor 2", got)
	}
	fd.estimate = false
	if _, err := client.ReadImportBracket(context.Background(), connect.NewRequest(&mtgv1.ReadImportBracketRequest{DeckId: res.GetDeck().GetId()})); err != nil {
		t.Fatal(err)
	}
	if got := store.sessions[res.GetSessionId()].GetSlots().GetPower().GetBracket(); got != 3 {
		t.Errorf("session bracket = %d, want the judged 3", got)
	}
}

// TestReadImportStepAsksAgain is D-864, F-172, and D-870: a 60-card
// import with no step asks the judge again at the next open. The step
// then fills the power slot of the session, and a deck with a step is
// not read again.
func TestReadImportStepAsksAgain(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	client, store := importServer(t, fd, ds, &fakeNoter{})
	list := "Deck\n4 Lightning Bolt\n56 Plains\n"
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: list, Format: mtgv1.FormatId_FORMAT_ID_MODERN})
	id := res.GetDeck().GetId()
	if res.GetDeck().GetPower() != nil || store.sessions[res.GetSessionId()].GetSlots().GetPower() != nil {
		t.Fatalf("deck power = %v, session power = %v", res.GetDeck().GetPower(), store.sessions[res.GetSessionId()].GetSlots().GetPower())
	}
	fd.step = mtgv1.SixtyStep_SIXTY_STEP_FNM
	got, err := client.ReadImportBracket(context.Background(), connect.NewRequest(&mtgv1.ReadImportBracketRequest{DeckId: id}))
	if err != nil {
		t.Fatal(err)
	}
	if got.Msg.GetDeck().GetPower().GetSixtyStep() != fd.step || fd.imports != 2 {
		t.Errorf("step = %v, reads = %d", got.Msg.GetDeck().GetPower(), fd.imports)
	}
	slots := store.sessions[res.GetSessionId()].GetSlots()
	if slots.GetPower().GetSixtyStep() != fd.step || slots.GetSlotStates()["power"] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("session power = %v, state = %v", slots.GetPower(), slots.GetSlotStates()["power"])
	}
	if _, err := client.ReadImportBracket(context.Background(), connect.NewRequest(&mtgv1.ReadImportBracketRequest{DeckId: id})); err != nil || fd.imports != 2 {
		t.Errorf("a judged step was read again: reads = %d, err = %v", fd.imports, err)
	}
}

// TestReadImportBracketKeepsConcurrentUsage is P2-1 of the review of
// #221: a turn that writes the session between the new read and its
// write keeps its usage, and the judge usage adds to it (D-447).
func TestReadImportBracketKeepsConcurrentUsage(t *testing.T) {
	fd, ds := &fakeDecks{estimate: true, record: true}, &fakeDeckStore{}
	client, store := importServer(t, fd, ds, &fakeNoter{})
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList})
	id := res.GetSessionId()
	base := proto.Clone(store.sessions[id].GetUsage()).(*mtgv1.Usage)
	fd.estimate = false
	// The first GetState of the write sees the session, then a turn lands
	// before the Put: it adds 7 calls and moves the version.
	fired := false
	store.onGetState = func() {
		if fired {
			return
		}
		fired = true
		u := proto.Clone(store.sessions[id]).(*mtgv1.Session)
		u.Usage = proto.Clone(base).(*mtgv1.Usage)
		u.Usage.Calls += 7
		store.sessions[id] = u
		store.versions[id]++
	}
	if _, err := client.ReadImportBracket(context.Background(), connect.NewRequest(&mtgv1.ReadImportBracketRequest{DeckId: res.GetDeck().GetId()})); err != nil {
		t.Fatal(err)
	}
	got := store.sessions[id]
	if got.GetUsage().GetCalls() != base.GetCalls()+7+1 {
		t.Errorf("calls = %d, want %d: the turn's 7 and the judge's 1", got.GetUsage().GetCalls(), base.GetCalls()+8)
	}
	if got.GetSlots().GetPower().GetBracket() != 3 {
		t.Errorf("session bracket = %d, want 3", got.GetSlots().GetPower().GetBracket())
	}
}

// TestImportRefusesALongList is P2-1 of the review of #220: a list over
// the line cap stores no deck and no session (D-846).
func TestImportRefusesALongList(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	client, store := importServer(t, fd, ds, &fakeNoter{})
	text := "1 Karlov of the Ghost Council\n" + strings.Repeat("// a note\n", 1000)
	_, err := client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: text}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Fatalf("err = %v, want InvalidArgument", err)
	}
	if len(ds.put) != 0 || len(store.sessions) != 0 || fd.imports != 0 {
		t.Errorf("a refused list stored %d decks and %d sessions", len(ds.put), len(store.sessions))
	}
}

// TestImportRefusesAListOfTooManyCards is REV-006 of the review of
// 2026-09-24. About 13 KB of text asked the profile for ten million
// copies. The list fails before the read of the import. With its four
// heading lines, a list of 998 card lines meets the line cap, so this
// list holds 990.
func TestImportRefusesAListOfTooManyCards(t *testing.T) {
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	client, store := importServer(t, fd, ds, &fakeNoter{})
	text := "Commander\n1 Karlov of the Ghost Council\n\nDeck\n" + strings.Repeat("10000 Forest\n", 990)
	_, err := client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: text}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument || importfault.IsUnreadable(err) {
		t.Fatalf("err = %v, want InvalidArgument and no report form", err)
	}
	if len(ds.put) != 0 || len(store.sessions) != 0 || fd.imports != 0 {
		t.Errorf("a refused list stored %d decks and %d sessions, and read %d imports", len(ds.put), len(store.sessions), fd.imports)
	}
}

// A list with no line that reads as a card offers the report form, and a
// list whose lines read but name no known card does not (D-887). A list
// over the line cap is a size limit, and it offers no form either.
func TestImportMarksAListThatDoesNotRead(t *testing.T) {
	cases := []struct {
		name, text string
		unreadable bool
	}{
		{"no card line", "hello\nworld\n", true},
		{"no known card", "1 Nosuch Cardname\n", false},
		{"over the line cap", strings.Repeat("// a note\n", 1001), false},
		{"over the card cap", strings.Repeat("1 Plains\n", 251), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client, _ := importServer(t, &fakeDecks{}, &fakeDeckStore{}, &fakeNoter{})
			_, err := client.ImportDeck(context.Background(), connect.NewRequest(&mtgv1.ImportDeckRequest{Text: c.text}))
			if connect.CodeOf(err) != connect.CodeInvalidArgument {
				t.Fatalf("err = %v, want InvalidArgument", err)
			}
			if got := importfault.IsUnreadable(err); got != c.unreadable {
				t.Errorf("unreadable = %v, want %v (%v)", got, c.unreadable, err)
			}
		})
	}
}

// TestImportAtTheCapRefusesEveryFormat is REV-002 of the review of
// 2026-09-24. The judge reads a 60-card list as it reads a Commander
// list, so a user at the cap gets ResourceExhausted and no judge call
// for each format (D-421).
func TestImportAtTheCapRefusesEveryFormat(t *testing.T) {
	sixty := "Deck\n4 Lightning Bolt\n56 Plains\n"
	for _, tc := range []struct {
		name   string
		req    *mtgv1.ImportDeckRequest
		spent  float64
		refuse bool
	}{
		{"Standard at the cap", &mtgv1.ImportDeckRequest{Text: sixty, Format: mtgv1.FormatId_FORMAT_ID_STANDARD}, 5, true},
		{"Modern at the cap", &mtgv1.ImportDeckRequest{Text: sixty, Format: mtgv1.FormatId_FORMAT_ID_MODERN}, 5, true},
		{"house rules at the cap", &mtgv1.ImportDeckRequest{Text: sixty, Format: mtgv1.FormatId_FORMAT_ID_HOUSE}, 5, true},
		{"Commander at the cap", &mtgv1.ImportDeckRequest{Text: markedList}, 5, true},
		{"Standard under the cap", &mtgv1.ImportDeckRequest{Text: sixty, Format: mtgv1.FormatId_FORMAT_ID_STANDARD}, 4.99, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cb, err := candidates.New()
			if err != nil {
				t.Fatal(err)
			}
			fd, ds := &fakeDecks{}, &fakeDeckStore{}
			ledger := &fakeLedger{spent: map[string]float64{"u1/1970-01": tc.spent}}
			opts := []Option{WithDecks(fd), WithCandidates(fixedIndex{importCardIndex()}, cb), WithDeckStore(ds), WithUsers(&fakeNoter{}), WithSpendCap(ledger, 5)}
			client, _ := testServerOpts(t, newFakeStore(), opts)
			_, err = client.ImportDeck(context.Background(), connect.NewRequest(tc.req))
			if !tc.refuse {
				if err != nil || fd.imports != 1 {
					t.Fatalf("an import under the cap: err = %v, reads = %d", err, fd.imports)
				}
				return
			}
			if connect.CodeOf(err) != connect.CodeResourceExhausted {
				t.Fatalf("err = %v, want ResourceExhausted", err)
			}
			if fd.imports != 0 || len(ds.put) != 0 {
				t.Errorf("a refused import read %d decks and stored %d", fd.imports, len(ds.put))
			}
		})
	}
}

// TestImportThenReviseKeepsTheImportedCommander: the revision of an
// imported Commander deck builds with the commander of the list. The
// index holds a more popular legend of the same colors, and the build
// picked it when the session named no commander (D-851).
func TestImportThenReviseKeepsTheImportedCommander(t *testing.T) {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	idx := cards.NewIndex(append(importCardIndex().All(), &mtgv1.Card{
		OracleId: "o-teysa", Name: "Teysa Karlov", TypeLine: "Legendary Creature - Human Advisor",
		CanBeCommander: true, ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B},
		Legalities: legal, EdhrecRank: 1,
	}), nil, nil, time.Unix(1000, 0).UTC())
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	fd.res = &generate.Result{Deck: &mtgv1.Deck{
		Name: "x", Summary: "a deck that draws", Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 2}, {OracleId: "o-plains", Name: "Plains", Count: 96}},
	}}
	opts := []Option{WithDecks(fd), WithCandidates(fixedIndex{idx}, cb), WithDeckStore(ds), WithUsers(&fakeNoter{})}
	client, _ := testServerOpts(t, newFakeStore(), opts,
		classifyJSON(t, nil),
		reviseJSON(t, map[string]any{"changes": []string{"Cut one Plains and add a card that draws cards"}}))
	res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList})
	ev := chat(t, client, &mtgv1.ChatRequest{SessionId: res.GetSessionId(), Message: "Add more card draw."})
	if fd.runs != 1 {
		t.Fatalf("runs = %d, failure = %v", fd.runs, ev.failure)
	}
	if got := fd.got.Commanders; len(got) != 1 || got[0] != "o-karlov" {
		t.Errorf("the revision built with commanders %v, want the imported [o-karlov]", got)
	}
}
