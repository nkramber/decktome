package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func ptr(v float64) *float64 { return &v }

// fixture is six invented cards, one per measure, plus a plain one.
func fixture() []*mtgv1.Card {
	return []*mtgv1.Card{
		{OracleId: "eq", TypeLine: "Artifact — Equipment", CardTypes: []string{"Artifact"}, Subtypes: []string{"Equipment"}, OracleText: "Equipped creature gets +1/+1.\nEquip {1}"},
		{OracleId: "ac", TypeLine: "Artifact // Artifact Creature — Construct", CardTypes: []string{"Artifact"}, OracleText: "Flying"},
		{OracleId: "au", CardTypes: []string{"Enchantment"}, Subtypes: []string{"Aura"}, OracleText: "Enchant creature\nEnchanted creature has hexproof."},
		{OracleId: "ts", CardTypes: []string{"Instant"}, OracleText: "Target creature you control gains hexproof until end of turn."},
		{OracleId: "hs", CardTypes: []string{"Enchantment"}, OracleText: "You have no maximum hand size."},
		{OracleId: "cs", CardTypes: []string{"Instant"}, OracleText: "Counter target spell. Draw a card."},
		{OracleId: "plain", TypeLine: "Creature — Human", CardTypes: []string{"Creature"}, OracleText: "Target creature an opponent controls gets -1/-0."},
	}
}

func tagged(tag, id string) bool {
	return (tag == "counterspell" && id == "cs") || (tag == "cantrip" && id == "cs")
}

func TestCountEachMeasure(t *testing.T) {
	got := Count(fixture(), []string{"heavy"}, tagged)
	want := map[string]float64{
		shortlist: 7, unmatchedWords: 1, "equipment": 1, "artifact_creatures": 1,
		"enchant_creature_auras": 1, "targeting_spells": 1, "hand_size": 1,
		"counterspells": 1, "cantrips": 1,
	}
	if len(got) != len(want) {
		t.Fatalf("measures %v, want %v", got, want)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s = %v, want %v", k, got[k], v)
		}
	}
}

// A targeting creature is not a spell, and an Aura with no "enchant
// creature" is not a voltron Aura.
func TestCountRefusesNearMisses(t *testing.T) {
	cards := []*mtgv1.Card{
		{OracleId: "a", CardTypes: []string{"Creature"}, OracleText: "When this enters, target creature you control gains flying."},
		{OracleId: "b", CardTypes: []string{"Enchantment"}, Subtypes: []string{"Aura"}, OracleText: "Enchant land"},
	}
	got := Count(cards, nil, tagged)
	if got["targeting_spells"] != 0 || got["enchant_creature_auras"] != 0 {
		t.Errorf("near misses counted: %v", got)
	}
}

func TestBarJudge(t *testing.T) {
	for _, tc := range []struct {
		bar  Bar
		v    float64
		want string
	}{
		{Bar{Todo: "set later"}, 5, Todo},
		{Bar{}, 5, Todo},
		{Bar{Max: ptr(0)}, 0, Pass},
		{Bar{Max: ptr(0)}, 1, Fail},
		{Bar{Min: ptr(3)}, 2, Fail},
		{Bar{Min: ptr(3), Max: ptr(9)}, 9, Pass},
		{Bar{Min: ptr(3), Max: ptr(9)}, 10, Fail},
	} {
		if got := tc.bar.Judge(tc.v); got != tc.want {
			t.Errorf("%+v judges %v as %s, want %s", tc.bar, tc.v, got, tc.want)
		}
	}
	if s := (Bar{Min: ptr(0.4), Max: ptr(12)}).String(); s != "0.400..12" {
		t.Errorf("bar string %q", s)
	}
}

func TestJudgeKeepsTheCaseOrder(t *testing.T) {
	step := Step{Name: "s", Measures: []Measure{
		{ID: unmatchedWords, Baseline: 5, Bar: Bar{Max: ptr(0)}},
		{ID: "equipment", Baseline: 131, Bar: Bar{Todo: "later"}},
	}}
	rows := Judge(step, map[string]float64{unmatchedWords: 2, "equipment": 40})
	if len(rows) != 2 || rows[0].Verdict != Fail || rows[1].Verdict != Todo || rows[1].Now != 40 {
		t.Errorf("rows %+v", rows)
	}
}

func TestCheckSHARefusesOtherBytes(t *testing.T) {
	sum := sha256.Sum256([]byte("a"))
	if err := checkSHA([]byte("a"), hex.EncodeToString(sum[:])); err != nil {
		t.Errorf("same bytes refused: %v", err)
	}
	if err := checkSHA([]byte("b"), hex.EncodeToString(sum[:])); err == nil {
		t.Error("other bytes passed")
	}
}

// The local file passes the hash before it is used, and a mismatch is
// refused with no download, because -no-fetch is set.
func TestCollectionRefusesAMismatch(t *testing.T) {
	dir := t.TempDir()
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, _ = zw.Write([]byte("Name,Set code\n"))
	_ = zw.Close()
	csvPath := filepath.Join(dir, "collection.csv")
	if err := os.WriteFile(csvPath+".gz", buf.Bytes(), 0o600); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(buf.Bytes())
	c := &Case{Collection: Collection{SHA256: hex.EncodeToString(sum[:])}}
	if err := collection(c, csvPath, true); err != nil {
		t.Fatalf("pinned bytes refused: %v", err)
	}
	if raw, _ := os.ReadFile(csvPath); string(raw) != "Name,Set code\n" {
		t.Errorf("csv %q", raw)
	}
	c.Collection.SHA256 = strings.Repeat("0", 64)
	if err := collection(c, csvPath, true); err == nil || !strings.Contains(err.Error(), "sha256") {
		t.Errorf("mismatch: %v", err)
	}
}

// The embedded case loads, names only known measures, and holds no uid.
func TestEmbeddedCaseLoads(t *testing.T) {
	c, err := LoadCase("first-user-ms-marvel")
	if err != nil {
		t.Fatal(err)
	}
	if len(c.Steps) != 2 || c.Steps[1].Avoid != "artifacts" {
		t.Errorf("steps %+v", c.Steps)
	}
	last := c.Steps[1].Measures
	for _, m := range last {
		if m.ID == unmatchedWords && (m.Bar.Max == nil || *m.Bar.Max != 0) {
			t.Error("step 2 must bar the unmatched words at 0")
		}
	}
	raw, _ := embedded.ReadFile("cases/first-user-ms-marvel.json")
	for _, bad := range []string{"users/", "@", "uid"} {
		if strings.Contains(string(raw), bad) {
			t.Errorf("the case holds %q", bad)
		}
	}
}

// TestDeckCountReadsTheDeck counts the copies of a land, each card
// measure, the distinct names, the grade, and the names kept from the
// deck of the step before (D-1124).
func TestDeckCountReadsTheDeck(t *testing.T) {
	cards := append(fixture(), &mtgv1.Card{OracleId: "isl", TypeLine: "Basic Land — Island", CardTypes: []string{"Land"}})
	byID := map[string]*mtgv1.Card{}
	for _, c := range cards {
		byID[c.GetOracleId()] = c
	}
	card := func(id string) (*mtgv1.Card, bool) { c, ok := byID[id]; return c, ok }
	prev := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "cs", Count: 1}, {OracleId: "eq", Count: 1}, {OracleId: "gone", Count: 1}}}
	d := &mtgv1.Deck{
		Cards:   []*mtgv1.DeckCard{{OracleId: "cs", Count: 1}, {OracleId: "eq", Count: 1}, {OracleId: "hs", Count: 1}, {OracleId: "isl", Count: 30}},
		Quality: &mtgv1.DeckQuality{Score: 0.5},
	}
	got := DeckCount(d, prev, card, tagged)
	want := map[string]float64{
		"lands": 30, "unique_names": 4, "quality_grade": 0.5, "kept_from_build_1": 2,
		"counterspells": 1, "cantrips": 1, "equipment": 1, "hand_size": 1,
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s = %v, want %v", k, got[k], v)
		}
	}
	if _, ok := got[shortlist]; ok {
		t.Error("a deck count holds the shortlist size")
	}
	if _, ok := DeckCount(d, nil, card, tagged)["kept_from_build_1"]; ok {
		t.Error("the first deck has no deck before it to keep from")
	}
}

// TestReadDecksByTurn reads the decks file of chat-probe, and deckAt
// answers the newest deck by a turn.
func TestReadDecksByTurn(t *testing.T) {
	path := filepath.Join(t.TempDir(), "decks.jsonl")
	body := `{"turn":4,"deck":{"id":"a"}}` + "\n\n" + `{"turn":6,"deck":{"id":"b"}}` + "\n"
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	decks, err := readDecks(path)
	if err != nil {
		t.Fatal(err)
	}
	for turn, want := range map[int]string{3: "", 4: "a", 5: "a", 6: "b", 9: "b"} {
		if got := deckAt(decks, turn).GetId(); got != want {
			t.Errorf("deckAt(%d) = %q, want %q", turn, got, want)
		}
	}
}
