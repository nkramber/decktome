package cards

import (
	"bytes"
	"compress/gzip"
	"slices"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// hobbitSets is the shape the 2026-08-31 snapshot holds around The
// Hobbit. hoc is a companion product and thob is its token set.
func hobbitSets() []SetInfo {
	return []SetInfo{
		{Code: "hob", Name: "The Hobbit", Type: "expansion", ReleasedAt: "2026-08-14"},
		{Code: "hoc", Name: "The Hobbit Eternal", Type: "eternal", ReleasedAt: "2026-08-14", ParentCode: "hob"},
		{Code: "thob", Name: "The Hobbit Tokens", Type: "token", ReleasedAt: "2026-08-14", ParentCode: "hob"},
		{Code: "ltr", Name: "The Lord of the Rings: Tales of Middle-earth", Type: "draft_innovation", ReleasedAt: "2023-06-23"},
		{Code: "ltc", Name: "Tales of Middle-earth Commander", Type: "commander", ReleasedAt: "2023-06-23", ParentCode: "ltr"},
		{Code: "ths", Name: "Theros", Type: "expansion", ReleasedAt: "2013-09-27"},
		{Code: "thb", Name: "Theros Beyond Death", Type: "expansion", ReleasedAt: "2020-01-24"},
		{Code: "pths", Name: "Theros Promos", Type: "promo", ReleasedAt: "2013-09-27", ParentCode: "ths"},
		{Code: "ala", Name: "Alchemy: Innistrad", Type: "alchemy", ReleasedAt: "2021-12-09", Digital: true},
	}
}

// TestFamilyHoldsTheCompanionProducts is D-376: a base set and every
// product Scryfall names under it are one family. The token child stays
// out, because it holds no playable card.
func TestFamilyHoldsTheCompanionProducts(t *testing.T) {
	tbl := NewSetTable(hobbitSets())
	for _, tc := range []struct {
		code string
		want []string
	}{
		{"hob", []string{"hob", "hoc"}},
		{"hoc", []string{"hoc"}},
		{"ltr", []string{"ltc", "ltr"}},
		{"ths", []string{"pths", "ths"}},
		{"thb", []string{"thb"}},
		// A code the table does not hold filters on itself alone.
		{"zzz", []string{"zzz"}},
	} {
		if got := tbl.Family(tc.code); !slices.Equal(got, tc.want) {
			t.Errorf("Family(%q) = %v, want %v", tc.code, got, tc.want)
		}
	}
}

// TestResolveReadsAProductName is D-376. A name prefix alone can not do
// this: "Tales of Middle-earth Commander" shares no prefix with "The
// Lord of the Rings: Tales of Middle-earth".
func TestResolveReadsAProductName(t *testing.T) {
	tbl := NewSetTable(hobbitSets())
	for _, tc := range []struct {
		phrase string
		want   []string
	}{
		{"The Hobbit", []string{"hob", "hoc"}},
		{"the hobbit set", []string{"hob", "hoc"}},
		{"hobbit", []string{"hob", "hoc"}},
		{"hob", []string{"hob", "hoc"}},
		{"HOC", []string{"hoc"}},
		{"Lord of the Rings", []string{"ltc", "ltr"}},
		{"the lord of the rings", []string{"ltc", "ltr"}},
		// An exact name wins over a phrase match, so "Theros" is Theros
		// and never Theros Beyond Death.
		{"Theros", []string{"pths", "ths"}},
		{"Theros Beyond Death", []string{"thb"}},
	} {
		got := tbl.Resolve(tc.phrase)
		if got.Kind != ResolveOne {
			t.Errorf("Resolve(%q) kind = %v, want ResolveOne", tc.phrase, got.Kind)
			continue
		}
		if !slices.Equal(got.Codes, tc.want) {
			t.Errorf("Resolve(%q) = %v, want %v", tc.phrase, got.Codes, tc.want)
		}
	}
}

// TestResolveAsksWhenTwoSetsAnswer is D-376: a phrase that names two
// base sets is a question, never a guess.
func TestResolveAsksWhenTwoSetsAnswer(t *testing.T) {
	tbl := NewSetTable([]SetInfo{
		{Code: "ktk", Name: "Khans of Tarkir", Type: "expansion", ReleasedAt: "2014-09-26"},
		{Code: "dtk", Name: "Dragons of Tarkir", Type: "expansion", ReleasedAt: "2015-03-27"},
		{Code: "tdm", Name: "Tarkir: Dragonstorm", Type: "expansion", ReleasedAt: "2025-04-11"},
	})
	got := tbl.Resolve("Tarkir")
	if got.Kind != ResolveMany {
		t.Fatalf("Resolve(Tarkir) kind = %v, want ResolveMany", got.Kind)
	}
	if len(got.Candidates) != 3 {
		t.Fatalf("Resolve(Tarkir) named %d sets, want 3", len(got.Candidates))
	}
	// Newest first, so the question offers the set a reader most likely
	// means before the twelve-year-old one.
	if got.Candidates[0].Code != "tdm" {
		t.Errorf("first candidate = %q, want tdm", got.Candidates[0].Code)
	}
}

// TestResolveNamesNothingForAnUnknownPhrase keeps the resolver from
// guessing. An unknown name is a question (D-376).
func TestResolveNamesNothingForAnUnknownPhrase(t *testing.T) {
	tbl := NewSetTable(hobbitSets())
	for _, phrase := range []string{"", "  ", "Yu-Gi-Oh", "a", "the pokemon set"} {
		if got := tbl.Resolve(phrase); got.Kind != ResolveNone {
			t.Errorf("Resolve(%q) = %v, want ResolveNone", phrase, got.Kind)
		}
	}
}

// TestResolveSkipsADigitalSet is D-306: this app offers paper cards only.
func TestResolveSkipsADigitalSet(t *testing.T) {
	tbl := NewSetTable(hobbitSets())
	if got := tbl.Resolve("Alchemy: Innistrad"); got.Kind != ResolveNone {
		t.Errorf("Resolve(Alchemy: Innistrad) = %v, want ResolveNone", got.Kind)
	}
	if got := tbl.Resolve("ala"); got.Kind != ResolveNone {
		t.Errorf("Resolve(ala) = %v, want ResolveNone", got.Kind)
	}
}

// TestResolveBoundsTheCandidateList: "commander" names over a hundred
// sets, and a question can not list them.
func TestResolveBoundsTheCandidateList(t *testing.T) {
	var rows []SetInfo
	for _, c := range []string{"c13", "c14", "c15", "c16", "c17", "c18"} {
		rows = append(rows, SetInfo{Code: c, Name: "Commander 20" + c[1:], Type: "commander", ReleasedAt: "20" + c[1:] + "-01-01"})
	}
	got := NewSetTable(rows).Resolve("commander")
	if got.Kind != ResolveMany {
		t.Fatalf("kind = %v, want ResolveMany", got.Kind)
	}
	if len(got.Candidates) > maxCandidates {
		t.Errorf("named %d sets, want at most %d", len(got.Candidates), maxCandidates)
	}
}

// TestSetCodesHoldEveryPaperSet is D-373: a card is not in one set, and
// a digital printing is not a paper set.
func TestSetCodesHoldEveryPaperSet(t *testing.T) {
	card := &mtgv1.Card{OracleId: "o1", Name: "Elvish Archdruid",
		DefaultPrinting: &mtgv1.Printing{ScryfallId: "p1", SetCode: "hoc"}}
	idx := NewIndex([]*mtgv1.Card{card}, []Printing{
		{ScryfallID: "p1", OracleID: "o1", SetCode: "HOC", CollectorNumber: "204"},
		{ScryfallID: "p2", OracleID: "o1", SetCode: "m19", CollectorNumber: "12"},
		{ScryfallID: "p3", OracleID: "o1", SetCode: "hoc", CollectorNumber: "300"},
		{ScryfallID: "p4", OracleID: "o1", SetCode: "mtgo", CollectorNumber: "1", Digital: true},
		{ScryfallID: "p5", OracleID: "o1", SetCode: "thob", CollectorNumber: "1", Layout: "token"},
	}, nil, time.Time{})
	got, _ := idx.ByOracleID("o1")
	want := []string{"hoc", "m19"}
	if !slices.Equal(got.GetSetCodes(), want) {
		t.Errorf("set_codes = %v, want %v", got.GetSetCodes(), want)
	}
}

// TestInSetsReadsTheWholeList is the filter every stage applies. An
// empty code list is no limit at all.
func TestInSetsReadsTheWholeList(t *testing.T) {
	card := &mtgv1.Card{SetCodes: []string{"hoc", "m19"}}
	for _, tc := range []struct {
		codes []string
		want  bool
	}{
		{nil, true},
		{[]string{"hob"}, false},
		{[]string{"hob", "hoc"}, true},
		{[]string{"m19"}, true},
	} {
		if got := InSets(card, CodeSet(tc.codes)); got != tc.want {
			t.Errorf("InSets(%v) = %v, want %v", tc.codes, got, tc.want)
		}
	}
	// A card with no paper printing is in no set.
	if InSets(&mtgv1.Card{}, CodeSet([]string{"hob"})) {
		t.Error("a card with no set codes passed a set filter")
	}
}

// TestDerivedTableHasNoFamily is D-377: a snapshot stored before the set
// file existed still filters by one set, and it resolves no family.
func TestDerivedTableHasNoFamily(t *testing.T) {
	idx := NewIndex([]*mtgv1.Card{{OracleId: "o1", Name: "A"}}, []Printing{
		{ScryfallID: "p1", OracleID: "o1", SetCode: "hob", SetName: "The Hobbit", CollectorNumber: "1"},
		{ScryfallID: "p2", OracleID: "o1", SetCode: "hoc", SetName: "The Hobbit Eternal", CollectorNumber: "1"},
	}, nil, time.Time{})
	tbl := idx.Sets()
	if !tbl.Derived() {
		t.Fatal("a build with no set rows must give a derived table")
	}
	if got := tbl.Family("hob"); !slices.Equal(got, []string{"hob"}) {
		t.Errorf("Family(hob) on a derived table = %v, want [hob]", got)
	}
	// The name still resolves, so one named set is still a filter.
	if got := tbl.Resolve("The Hobbit"); got.Kind != ResolveOne || !slices.Equal(got.Codes, []string{"hob"}) {
		t.Errorf("Resolve on a derived table = %v %v", got.Kind, got.Codes)
	}
}

// TestSetFileRoundTrip proves the stored shape reads back.
func TestSetFileRoundTrip(t *testing.T) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if err := EncodeSets(gz, hobbitSets()); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	got, err := LoadSets(&buf, SetsFile)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(hobbitSets()) {
		t.Fatalf("read %d rows, want %d", len(got), len(hobbitSets()))
	}
	if got[1].ParentCode != "hob" || got[1].Code != "hoc" {
		t.Errorf("row 1 = %+v", got[1])
	}
}

// TestNamesReadsTheCodes: a message names the sets, not their codes.
func TestNamesReadsTheCodes(t *testing.T) {
	tbl := NewSetTable(hobbitSets())
	got := strings.Join(tbl.Names([]string{"hob", "hoc", "zzz"}), ", ")
	want := "The Hobbit, The Hobbit Eternal, zzz"
	if got != want {
		t.Errorf("Names = %q, want %q", got, want)
	}
}

// TestResolveNeverLeadsWithABonusSheet is D-518: "Marvel" matches the
// expansion and the bonus sheet, and the sheet never leads a family
// beside a product. The sheet still resolves by its exact name, and the
// possessive of "Marvel's Spider-Man" is a different word on purpose.
func TestResolveNeverLeadsWithABonusSheet(t *testing.T) {
	tbl := NewSetTable([]SetInfo{
		{Code: "msh", Name: "Marvel Super Heroes", Type: "expansion", ReleasedAt: "2026-06-26"},
		{Code: "msc", Name: "Marvel Super Heroes Commander", Type: "commander", ReleasedAt: "2026-06-26", ParentCode: "msh"},
		{Code: "mar", Name: "Marvel Universe", Type: "masterpiece", ReleasedAt: "2025-09-26"},
		{Code: "spm", Name: "Marvel's Spider-Man", Type: "expansion", ReleasedAt: "2025-09-26"},
	})
	got := tbl.Resolve("Marvel")
	if got.Kind != ResolveOne || !slices.Equal(got.Codes, []string{"msc", "msh"}) {
		t.Fatalf("Resolve(Marvel) = %v %v, want the Marvel Super Heroes family alone", got.Kind, got.Codes)
	}
	exact := tbl.Resolve("Marvel Universe")
	if exact.Kind != ResolveOne || !slices.Equal(exact.Codes, []string{"mar"}) {
		t.Errorf("Resolve(Marvel Universe) = %v %v, want mar by its exact name", exact.Kind, exact.Codes)
	}
	sheetOnly := NewSetTable([]SetInfo{
		{Code: "mar", Name: "Marvel Universe", Type: "masterpiece", ReleasedAt: "2025-09-26"},
	})
	if got := sheetOnly.Resolve("Marvel"); got.Kind != ResolveOne || !slices.Equal(got.Codes, []string{"mar"}) {
		t.Errorf("a sheet that is the one match still resolves: %v %v", got.Kind, got.Codes)
	}
}

// TestResolveGroupReadsEveryFamilyOfAFranchise is D-525: a group word
// reaches every family whose name holds it, the possessive drops before
// the match, and the bonus sheet and its inserts stay out. "Marvel"
// alone still resolves to one product (D-518).
func TestResolveGroupReadsEveryFamilyOfAFranchise(t *testing.T) {
	tbl := NewSetTable([]SetInfo{
		{Code: "msh", Name: "Marvel Super Heroes", Type: "expansion", ReleasedAt: "2026-06-26"},
		{Code: "msc", Name: "Marvel Super Heroes Commander", Type: "commander", ReleasedAt: "2026-06-26", ParentCode: "msh"},
		{Code: "tmsh", Name: "Marvel Super Heroes Tokens", Type: "token", ReleasedAt: "2026-06-26", ParentCode: "msh"},
		{Code: "mar", Name: "Marvel Universe", Type: "masterpiece", ReleasedAt: "2025-09-26"},
		{Code: "lmar", Name: "Marvel Legends Series Inserts", Type: "promo", ReleasedAt: "2025-09-30", ParentCode: "mar"},
		{Code: "spm", Name: "Marvel's Spider-Man", Type: "expansion", ReleasedAt: "2025-09-26"},
		{Code: "spe", Name: "Marvel's Spider-Man Eternal", Type: "eternal", ReleasedAt: "2025-09-26", ParentCode: "spm"},
		{Code: "hob", Name: "The Hobbit", Type: "expansion", ReleasedAt: "2026-08-14"},
	})
	want := []string{"msc", "msh", "spe", "spm"}
	if got := tbl.ResolveGroup("Marvel"); !slices.Equal(got, want) {
		t.Errorf("ResolveGroup(Marvel) = %v, want %v", got, want)
	}
	if got := tbl.ResolveGroup("Marvel's"); !slices.Equal(got, want) {
		t.Errorf("ResolveGroup(Marvel's) = %v, want %v", got, want)
	}
	if got := tbl.ResolveGroup("Star Wars"); got != nil {
		t.Errorf("an unknown group = %v, want nothing", got)
	}
	if got := tbl.Resolve("Marvel"); got.Kind != ResolveOne || !slices.Equal(got.Codes, []string{"msc", "msh"}) {
		t.Errorf("Resolve(Marvel) = %v %v, want one product", got.Kind, got.Codes)
	}
}

// TestResolveAllReadsSeveralNamedSets is F-64: a phrase that names two
// sets answers the union of their families. The whole phrase answers
// first, so a set name that holds a connector word keeps its meaning.
func TestResolveAllReadsSeveralNamedSets(t *testing.T) {
	tbl := NewSetTable(hobbitSets())
	for _, tc := range []struct {
		phrase string
		want   []string
	}{
		{"Hobbit and Lord of the Rings", []string{"hob", "hoc", "ltc", "ltr"}},
		{"Hobbit, Lord of the Rings", []string{"hob", "hoc", "ltc", "ltr"}},
		{"hob + ltr", []string{"hob", "hoc", "ltc", "ltr"}},
		{"The Hobbit & the lord of the rings", []string{"hob", "hoc", "ltc", "ltr"}},
		// A single name still answers as it did, and the order of the
		// union never repeats a code.
		{"Hobbit", []string{"hob", "hoc"}},
		{"Hobbit and the hobbit set", []string{"hob", "hoc"}},
	} {
		got := tbl.ResolveAll(tc.phrase)
		if got.Kind != ResolveOne {
			t.Errorf("ResolveAll(%q) kind = %v, want ResolveOne", tc.phrase, got.Kind)
			continue
		}
		if !slices.Equal(got.Codes, tc.want) {
			t.Errorf("ResolveAll(%q) = %v, want %v", tc.phrase, got.Codes, tc.want)
		}
	}
}

// TestResolveAllKeepsTheAnswerOfTheWholePhrase is F-64: a part this app
// can not settle never becomes a guess. The answer of the whole phrase
// stands, so the agent asks exactly as it did.
func TestResolveAllKeepsTheAnswerOfTheWholePhrase(t *testing.T) {
	tbl := NewSetTable(hobbitSets())
	for _, phrase := range []string{
		"Hobbit and LOTR",
		"LOTR",
		"Hobbit and nothing anyone printed",
		"",
	} {
		if got := tbl.ResolveAll(phrase); got.Kind == ResolveOne {
			t.Errorf("ResolveAll(%q) = %v, want no answer", phrase, got.Codes)
		}
	}
}
