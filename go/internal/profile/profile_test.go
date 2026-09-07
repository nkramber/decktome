package profile

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/rules"
	"github.com/nkramber/decktome/go/internal/spellbook"
)

// The test cards. Each one carries the parsed types, the produced mana,
// and the text the readers look at, and nothing else.

type spec struct {
	name     string
	cost     string
	mv       float64
	types    []string
	supers   []string
	produced []mtgv1.Color
	text     string
	colors   []mtgv1.Color
	changer  bool
}

func card(s spec) *mtgv1.Card {
	return &mtgv1.Card{
		OracleId: "oid-" + strings.ToLower(strings.ReplaceAll(s.name, " ", "-")), Name: s.name,
		ManaCost: s.cost, ManaValue: s.mv, CardTypes: s.types, Supertypes: s.supers,
		ProducedMana: s.produced, OracleText: s.text, Colors: s.colors, ColorIdentity: s.colors,
		GameChanger: s.changer,
	}
}

var (
	W = mtgv1.Color_COLOR_W
	B = mtgv1.Color_COLOR_B
	C = mtgv1.Color_COLOR_C
)

var testCards = map[string]*mtgv1.Card{}

func add(s spec) *mtgv1.Card {
	c := card(s)
	testCards[c.OracleId] = c
	return c
}

var (
	plains     = add(spec{name: "Plains", types: []string{"Land"}, supers: []string{"Basic"}, produced: []mtgv1.Color{W}})
	swamp      = add(spec{name: "Swamp", types: []string{"Land"}, supers: []string{"Basic"}, produced: []mtgv1.Color{B}})
	tapDual    = add(spec{name: "Scoured Barrens", types: []string{"Land"}, produced: []mtgv1.Color{W, B}, text: "Scoured Barrens enters tapped.\nWhen it enters, you gain 1 life."})
	shock      = add(spec{name: "Godless Shrine", types: []string{"Land"}, produced: []mtgv1.Color{W, B}, text: "As Godless Shrine enters, you may pay 2 life. If you don't, it enters tapped."})
	fetch      = add(spec{name: "Marsh Flats", types: []string{"Land"}, text: "{T}, Pay 1 life, Sacrifice Marsh Flats: Search your library for a Plains or Swamp card, put it onto the battlefield, then shuffle."})
	rogue      = add(spec{name: "Rogue's Passage", types: []string{"Land"}, produced: []mtgv1.Color{C}, text: "{T}: Add {C}."})
	solRing    = add(spec{name: "Sol Ring", cost: "{1}", mv: 1, types: []string{"Artifact"}, produced: []mtgv1.Color{C}, text: "{T}: Add {C}{C}."})
	signet     = add(spec{name: "Orzhov Signet", cost: "{2}", mv: 2, types: []string{"Artifact"}, produced: []mtgv1.Color{W, B}, text: "{1}, {T}: Add {W}{B}."})
	darkRitual = add(spec{name: "Dark Ritual", cost: "{B}", mv: 1, types: []string{"Instant"}, produced: []mtgv1.Color{B}, text: "Add {B}{B}{B}.", colors: []mtgv1.Color{B}})
	tutor      = add(spec{name: "Demonic Tutor", cost: "{1}{B}", mv: 2, types: []string{"Sorcery"}, text: "Search your library for a card, put that card into your hand, then shuffle.", colors: []mtgv1.Color{B}, changer: true})
	wrath      = add(spec{name: "Wrath of God", cost: "{2}{W}{W}", mv: 4, types: []string{"Sorcery"}, text: "Destroy all creatures. They can't be regenerated.", colors: []mtgv1.Color{W}})
	knight     = add(spec{name: "Knight of the White Orchid", cost: "{W}{W}", mv: 2, types: []string{"Creature"}, text: "First strike", colors: []mtgv1.Color{W}})
	angel      = add(spec{name: "Serra Angel", cost: "{3}{W}{W}", mv: 5, types: []string{"Creature"}, text: "Flying, vigilance", colors: []mtgv1.Color{W}})
	karlov     = add(spec{name: "Karlov of the Ghost Council", cost: "{W}{B}", mv: 2, types: []string{"Creature"}, supers: []string{"Legendary"}, text: "Whenever you gain life, put two +1/+1 counters on Karlov.", colors: []mtgv1.Color{W, B}})
)

type source map[string]*mtgv1.Card

func (s source) ByOracleID(id string) (*mtgv1.Card, bool) {
	c, ok := s[id]
	return c, ok
}

// deckOf builds a Commander deck from counts and roles.
func deckOf(bracket int32, rows ...row) *mtgv1.Deck {
	d := &mtgv1.Deck{
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		CommanderOracleIds: []string{karlov.OracleId},
		Validation:         &mtgv1.ValidationResult{},
	}
	if bracket > 0 {
		d.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bracket}}
	}
	for _, r := range rows {
		d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: r.c.OracleId, Name: r.c.Name, Count: r.n, Role: r.role})
	}
	return d
}

type row struct {
	c    *mtgv1.Card
	n    int32
	role mtgv1.CardRole
}

// shaped is a 99-card bracket 3 deck that sits in every band: 36 lands
// of which 2 enter tapped, 10 ramp, 10 draw, 8 removal, 3 wipes, 6
// interaction, and filler at mana value two to four.
func shaped() []row {
	return []row{
		{plains, 17, mtgv1.CardRole_CARD_ROLE_LAND}, {swamp, 15, mtgv1.CardRole_CARD_ROLE_LAND},
		{tapDual, 2, mtgv1.CardRole_CARD_ROLE_LAND}, {shock, 1, mtgv1.CardRole_CARD_ROLE_LAND}, {fetch, 1, mtgv1.CardRole_CARD_ROLE_LAND},
		{solRing, 1, mtgv1.CardRole_CARD_ROLE_RAMP}, {signet, 9, mtgv1.CardRole_CARD_ROLE_RAMP},
		{knight, 10, mtgv1.CardRole_CARD_ROLE_DRAW}, {knight, 8, mtgv1.CardRole_CARD_ROLE_REMOVAL},
		{wrath, 3, mtgv1.CardRole_CARD_ROLE_WIPE}, {knight, 6, mtgv1.CardRole_CARD_ROLE_INTERACTION},
		{angel, 10, mtgv1.CardRole_CARD_ROLE_THREAT}, {knight, 16, mtgv1.CardRole_CARD_ROLE_SYNERGY},
	}
}

type fakeClassifier struct {
	res   *spellbook.Result
	err   error
	sent  []string
	cmdrs []string
}

func (f *fakeClassifier) EstimateBracket(_ context.Context, commanders, main []string) (*spellbook.Result, error) {
	f.cmdrs, f.sent = commanders, main
	return f.res, f.err
}

func newProfiler(t *testing.T, classify Classifier) *Profiler {
	t.Helper()
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	p, err := New(cfg, func() *cards.TagIndex { return nil }, classify)
	if err != nil {
		t.Fatal(err)
	}
	p.SetHands(400)
	return p
}

func feature(t *testing.T, p *mtgv1.DeckProfile, key string) *mtgv1.ProfileFeature {
	t.Helper()
	for _, f := range p.GetFeatures() {
		if f.GetKey() == key {
			return f
		}
	}
	t.Fatalf("no feature %q in %v", key, p.GetFeatures())
	return nil
}

func TestReadMeasuresTheShapedDeckInBand(t *testing.T) {
	fc := &fakeClassifier{res: &spellbook.Result{BracketTag: "C"}}
	p := newProfiler(t, fc)
	prof, findings := p.Read(context.Background(), deckOf(3, shaped()...), source(testCards))
	if prof.GetBracket() != 3 || prof.GetBandsVerifiedAt() == "" {
		t.Errorf("profile head %v", prof)
	}
	want := map[string]float64{
		KeyLand: 36, KeyTappedLand: 2, KeyColorlessLand: 0, KeyRamp: 10, KeyDraw: 10,
		KeyRemoval: 8, KeyWipe: 3, KeyInteraction: 6, KeyFastMana: 1, KeyGameChanger: 0,
	}
	for key, v := range want {
		if f := feature(t, prof, key); f.GetValue() != v || f.GetOffBand() {
			t.Errorf("%s = %v off %v, want %v in band", key, f.GetValue(), f.GetOffBand(), v)
		}
	}
	if f := feature(t, prof, KeyAvgManaValue); f.GetValue() < 2.5 || f.GetValue() > 3.5 || f.GetOffBand() {
		t.Errorf("avg mana value %v", f)
	}
	if f := feature(t, prof, KeyColorSources); f.GetValue() < 0.85 || f.GetOffBand() {
		t.Errorf("color sources %v", f)
	}
	if g := prof.GetGoldfish(); g.GetHands() != 400 || g.GetManaTurnFour() < 4 || g.GetCommanderTurn() > 3 {
		t.Errorf("goldfish %v", g)
	}
	if !prof.GetContent().GetChecked() || prof.GetContent().GetSourceTag() != "C" {
		t.Errorf("content %v", prof.GetContent())
	}
	if len(findings) != 0 {
		t.Errorf("findings %v", findings)
	}
	// The classifier got the commander and every nonbasic name once.
	if len(fc.cmdrs) != 1 || fc.cmdrs[0] != "Karlov of the Ghost Council" {
		t.Errorf("commanders sent %v", fc.cmdrs)
	}
	for _, n := range fc.sent {
		if n == "Plains" || n == "Swamp" {
			t.Errorf("a basic land was sent: %v", fc.sent)
		}
	}
}

func TestReadFindsAnOffBandFeature(t *testing.T) {
	rows := shaped()
	// Cut five lands for five more spells, and swap the wipes for tutors.
	rows[0].n = 12
	rows = append(rows, row{knight, 5, mtgv1.CardRole_CARD_ROLE_SYNERGY})
	p := newProfiler(t, &fakeClassifier{res: &spellbook.Result{}})
	prof, findings := p.Read(context.Background(), deckOf(3, rows...), source(testCards))
	if f := feature(t, prof, KeyLand); f.GetValue() != 31 || !f.GetOffBand() || f.GetLow() != 34 || f.GetHigh() != 38 {
		t.Errorf("land %v", f)
	}
	var found bool
	for _, f := range findings {
		if f.GetCode() == CodeOffBand && strings.Contains(f.GetMessage(), "the land count is 31, and bracket 3 wants 34 to 38") {
			found = true
		}
		if f.GetSeverity() != mtgv1.Severity_SEVERITY_WARN && f.GetSeverity() != mtgv1.Severity_SEVERITY_INFO {
			t.Errorf("a profile finding must never block: %v", f)
		}
	}
	if !found {
		t.Errorf("want the land finding, got %v", findings)
	}
}

func TestReadMarksTheGameChangerBandWithoutASecondFinding(t *testing.T) {
	rows := append(shaped(), row{tutor, 1, mtgv1.CardRole_CARD_ROLE_DRAW})
	rows[0].n--
	p := newProfiler(t, &fakeClassifier{res: &spellbook.Result{}})
	prof, findings := p.Read(context.Background(), deckOf(2, rows...), source(testCards))
	f := feature(t, prof, KeyGameChanger)
	if f.GetValue() != 1 || !f.GetOffBand() || f.GetHigh() != 0 || !f.GetHasHigh() {
		t.Errorf("game changer %v", f)
	}
	for _, fd := range findings {
		if strings.Contains(fd.GetMessage(), "Game Changer") {
			t.Errorf("the rules engine owns the Game Changer finding, got %v", fd)
		}
	}
	if f := feature(t, prof, KeyFastMana); f.GetValue() != 1 || f.GetOffBand() {
		t.Errorf("fast mana at bracket 2 %v", f)
	}
}

func TestReadContentRulesByBracket(t *testing.T) {
	combo := spellbook.ClassifiedCombo{
		Combo: spellbook.ComboRef{ID: "1", Uses: []spellbook.ComboUse{
			{Card: spellbook.CardRef{Name: "Demonic Consultation"}}, {Card: spellbook.CardRef{Name: "Thassa's Oracle"}}}},
		Relevant: true, DefinitelyTwoCard: true, Speed: 4,
	}
	res := &spellbook.Result{BracketTag: "R",
		Cards: []spellbook.ClassifiedCard{
			{Card: spellbook.CardRef{Name: "Armageddon"}, Quantity: 1, MassLandDenial: true},
			{Card: spellbook.CardRef{Name: "Time Warp"}, Quantity: 1, ExtraTurn: true},
			{Card: spellbook.CardRef{Name: "Temporal Manipulation"}, Quantity: 1, ExtraTurn: true},
		},
		Combos: []spellbook.ClassifiedCombo{combo},
	}
	cases := []struct {
		bracket int32
		want    []string
	}{
		{1, []string{CodeMassLandDenial, CodeExtraTurns, CodeTwoCardCombo}},
		{2, []string{CodeMassLandDenial, CodeExtraTurns, CodeTwoCardCombo}},
		{3, []string{CodeMassLandDenial, CodeExtraTurns, CodeTwoCardCombo}},
		{4, nil},
		{5, nil},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprint("bracket ", tc.bracket), func(t *testing.T) {
			p := newProfiler(t, &fakeClassifier{res: res})
			prof, findings := p.Read(context.Background(), deckOf(tc.bracket, shaped()...), source(testCards))
			var got []string
			for _, f := range findings {
				if f.GetCode() != CodeOffBand {
					got = append(got, f.GetCode())
				}
			}
			if strings.Join(got, ",") != strings.Join(tc.want, ",") {
				t.Errorf("codes %v, want %v", got, tc.want)
			}
			c := prof.GetContent()
			if len(c.GetMassLandDenial()) != 1 || len(c.GetExtraTurns()) != 2 || len(c.GetCombos()) != 1 || !c.GetCombos()[0].GetTwoCard() {
				t.Errorf("content %v", c)
			}
		})
	}
	// A slow combo passes bracket 2 and 3, and fails bracket 1.
	slow := res
	slow.Combos[0].Speed = 2
	for _, tc := range []struct {
		bracket int32
		hit     bool
	}{{1, true}, {2, false}, {3, false}} {
		p := newProfiler(t, &fakeClassifier{res: slow})
		_, findings := p.Read(context.Background(), deckOf(tc.bracket, shaped()...), source(testCards))
		var hit bool
		for _, f := range findings {
			if f.GetCode() == CodeTwoCardCombo {
				hit = true
			}
		}
		if hit != tc.hit {
			t.Errorf("bracket %d slow combo hit %v, want %v", tc.bracket, hit, tc.hit)
		}
	}
}

func TestReadReportsAnUncheckedContentRun(t *testing.T) {
	p := newProfiler(t, &fakeClassifier{err: errors.New("status 502")})
	prof, findings := p.Read(context.Background(), deckOf(2, shaped()...), source(testCards))
	if prof.GetContent().GetChecked() || !strings.Contains(prof.GetContent().GetError(), "502") {
		t.Errorf("content %v", prof.GetContent())
	}
	var content []*mtgv1.Finding
	for _, f := range findings {
		if f.GetCode() != CodeOffBand {
			content = append(content, f)
		}
	}
	if len(content) != 1 || content[0].GetCode() != CodeContentUnchecked || content[0].GetSeverity() != mtgv1.Severity_SEVERITY_INFO {
		t.Errorf("findings %v", findings)
	}
	// No classifier at all says so too.
	p = newProfiler(t, nil)
	prof, _ = p.Read(context.Background(), deckOf(2, shaped()...), source(testCards))
	if prof.GetContent().GetChecked() || prof.GetContent().GetError() == "" {
		t.Errorf("content with no classifier %v", prof.GetContent())
	}
}

func TestReadSixtyCardDeckUsesTheStepBands(t *testing.T) {
	d := &mtgv1.Deck{
		Format:     &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN},
		Power:      &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT}},
		Validation: &mtgv1.ValidationResult{},
	}
	for _, r := range []row{{plains, 14, 0}, {tapDual, 10, 0}, {knight, 20, 0}, {wrath, 8, 0}, {angel, 8, 0}} {
		d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: r.c.OracleId, Name: r.c.Name, Count: r.n})
	}
	p := newProfiler(t, &fakeClassifier{res: &spellbook.Result{}})
	prof, findings := p.Read(context.Background(), d, source(testCards))
	if prof.GetBracket() != 0 || prof.GetContent() != nil {
		t.Errorf("a 60-card deck has no bracket and no content check: %v", prof)
	}
	if f := feature(t, prof, KeyTappedLand); f.GetValue() != 10 || !f.GetOffBand() || f.GetHigh() != 4 {
		t.Errorf("tapped %v", f)
	}
	if f := feature(t, prof, KeyLand); f.GetHasHigh() || f.GetOffBand() {
		t.Errorf("a 60-card land count has no band: %v", f)
	}
	for _, f := range prof.GetFeatures() {
		if f.GetKey() == KeyCommanderTurnOverMV {
			t.Error("a 60-card deck has no commander turn")
		}
	}
	var found bool
	for _, f := range findings {
		if f.GetCode() == CodeOffBand && strings.Contains(f.GetMessage(), "the tournament level wants 4 at most") {
			found = true
		}
	}
	if !found {
		t.Errorf("want the tapped finding, got %v", findings)
	}
}

func TestReaders(t *testing.T) {
	if !entersTapped(tapDual) || entersTapped(shock) || entersTapped(plains) {
		t.Error("enters tapped")
	}
	if !isColorlessLand(rogue) || isColorlessLand(fetch) || isColorlessLand(plains) {
		t.Error("colorless land")
	}
	if !isFastMana(solRing) || !isFastMana(darkRitual) || isFastMana(signet) || isFastMana(knight) {
		t.Error("fast mana")
	}
	if !isFetch(fetch) || isFetch(tapDual) {
		t.Error("fetch")
	}
	if manaMade(solRing) != 2 || manaMade(signet) != 1 {
		t.Error("mana made")
	}
	pips := colorPips("{2}{W}{W}{W/U}{B/P}")
	if pips[W] != 2 || pips[mtgv1.Color_COLOR_U] != 0 || pips[B] != 0 {
		t.Errorf("pips %v", pips)
	}
	if got := colorPips(""); len(got) != 0 {
		t.Errorf("pips of no cost %v", got)
	}
}

func TestOffBandMessageShapes(t *testing.T) {
	power := &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 2}}
	cases := []struct {
		row  *mtgv1.ProfileFeature
		want string
	}{
		{&mtgv1.ProfileFeature{Key: KeyLand, Value: 31, Low: 35, High: 40, HasHigh: true}, "the land count is 31, and bracket 2 wants 35 to 40"},
		{&mtgv1.ProfileFeature{Key: KeyTutor, Value: 4, Low: 0, High: 2, HasHigh: true, Note: "Demonic Tutor"}, "the tutor count is 4, and bracket 2 wants 2 at most (Demonic Tutor)"},
		{&mtgv1.ProfileFeature{Key: KeyManaTurnFour, Value: 3.1, Low: 3.8}, "the mana available on turn four is 3.1, and bracket 2 wants 3.8 or more"},
	}
	for _, tc := range cases {
		if got := offBandMessage(tc.row, 2, mtgv1.FormatId_FORMAT_ID_COMMANDER, power); got != tc.want {
			t.Errorf("got %q\nwant %q", got, tc.want)
		}
	}
}

func TestReadNearTwoCardComboReadsOneStepSlower(t *testing.T) {
	// Sanguine Bond and Exquisite Blood: the endpoint reads the pair as a
	// near two-card combo at speed 5, which is speed 4 for the cap.
	near := &spellbook.Result{Combos: []spellbook.ClassifiedCombo{{
		Combo: spellbook.ComboRef{ID: "2", Uses: []spellbook.ComboUse{
			{Card: spellbook.CardRef{Name: "Sanguine Bond"}}, {Card: spellbook.CardRef{Name: "Exquisite Blood"}}}},
		Relevant: true, ArguablyTwoCard: true, DefinitelyTwoCard: false, Speed: 5,
	}}}
	for _, tc := range []struct {
		bracket int32
		hit     bool
	}{{1, true}, {2, true}, {3, true}, {4, false}} {
		p := newProfiler(t, &fakeClassifier{res: near})
		prof, findings := p.Read(context.Background(), deckOf(tc.bracket, shaped()...), source(testCards))
		var hit bool
		for _, f := range findings {
			if f.GetCode() == CodeTwoCardCombo {
				hit = true
				if !strings.Contains(f.GetMessage(), "near two-card") {
					t.Errorf("the finding must say near two-card: %s", f.GetMessage())
				}
			}
		}
		if hit != tc.hit {
			t.Errorf("bracket %d near combo hit %v, want %v", tc.bracket, hit, tc.hit)
		}
		if c := prof.GetContent().GetCombos(); len(c) != 1 || !c[0].GetTwoCard() || c[0].GetSpeed() != 5 {
			t.Errorf("combo hit %v", c)
		}
	}
	// A near combo at speed 4 reads as 3, which bracket 3 allows.
	near.Combos[0].Speed = 4
	p := newProfiler(t, &fakeClassifier{res: near})
	_, findings := p.Read(context.Background(), deckOf(3, shaped()...), source(testCards))
	for _, f := range findings {
		if f.GetCode() == CodeTwoCardCombo {
			t.Errorf("bracket 3 allows a near combo at speed 4: %s", f.GetMessage())
		}
	}
}

func TestCutShortlistDropsWhatTheBracketForbids(t *testing.T) {
	res := &spellbook.Result{Cards: []spellbook.ClassifiedCard{
		{Card: spellbook.CardRef{Name: "Armageddon"}, MassLandDenial: true},
		{Card: spellbook.CardRef{Name: "Time Warp"}, ExtraTurn: true},
		{Card: spellbook.CardRef{Name: "Sol Ring"}},
	}}
	names := []string{"Armageddon", "Time Warp", "Sol Ring"}
	cases := []struct {
		bracket int32
		want    []string
	}{
		{1, []string{"Armageddon", "Time Warp"}},
		{2, []string{"Armageddon"}},
		{3, []string{"Armageddon"}},
		{4, nil},
		{5, nil},
		{9, nil},
	}
	for _, tc := range cases {
		fc := &fakeClassifier{res: res}
		p := newProfiler(t, fc)
		got, err := p.CutShortlist(context.Background(), mtgv1.FormatId_FORMAT_ID_COMMANDER, tc.bracket, []string{"Karlov of the Ghost Council"}, names)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Join(got, ",") != strings.Join(tc.want, ",") {
			t.Errorf("bracket %d: drop %v, want %v", tc.bracket, got, tc.want)
		}
		if tc.want == nil && fc.sent != nil {
			t.Errorf("bracket %d needs no call", tc.bracket)
		}
	}
	// A 60-card format never calls, and a failed call drops nothing.
	fc := &fakeClassifier{res: res}
	p := newProfiler(t, fc)
	if got, err := p.CutShortlist(context.Background(), mtgv1.FormatId_FORMAT_ID_MODERN, 1, nil, names); got != nil || err != nil || fc.sent != nil {
		t.Errorf("modern: %v %v sent %v", got, err, fc.sent)
	}
	p = newProfiler(t, &fakeClassifier{err: errors.New("status 502")})
	if got, err := p.CutShortlist(context.Background(), mtgv1.FormatId_FORMAT_ID_COMMANDER, 1, nil, names); got != nil || err == nil {
		t.Errorf("a failed call must drop nothing and say so: %v %v", got, err)
	}
	p = newProfiler(t, nil)
	if got, err := p.CutShortlist(context.Background(), mtgv1.FormatId_FORMAT_ID_COMMANDER, 1, nil, names); got != nil || err != nil {
		t.Errorf("no classifier: %v %v", got, err)
	}
}

func TestReadWithNoTagSource(t *testing.T) {
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	p, err := New(cfg, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	p.SetHands(50)
	prof, _ := p.Read(context.Background(), deckOf(2, shaped()...), source(testCards))
	for _, f := range prof.GetFeatures() {
		if f.GetKey() == KeyTutor {
			t.Error("no tag source means no tutor row")
		}
	}
	if prof.GetGoldfish() == nil {
		t.Error("the rest of the profile must still read")
	}
}
