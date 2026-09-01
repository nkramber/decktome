package cards

import (
	"encoding/json"
	"io"
	"sort"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/scryfall"
)

// SetInfo is one Magic set, as the app reads it. The fields come from the
// Scryfall /sets endpoint, which is the only source of ParentCode: the
// bulk card files carry the set code and the set name and no link between
// two products of one release (D-377).
type SetInfo struct {
	Code string `json:"code"`
	Name string `json:"name"`
	// Type is the Scryfall set_type, for example "expansion" or
	// "commander". Empty in the fallback table, which is built from the
	// printings and reads no type.
	Type string `json:"set_type"`
	// ReleasedAt is the release date, ISO 8601. Empty when unknown.
	ReleasedAt string `json:"released_at"`
	// ParentCode names the base set of a companion product. Empty in the
	// fallback table and on every base set.
	ParentCode string `json:"parent_set_code"`
	Digital    bool   `json:"digital"`
}

// familySkipTypes are the child products a set family leaves out. A token
// set holds no playable card, and memorabilia and minigame products hold
// no card a deck can run (D-376).
var familySkipTypes = map[string]bool{
	"token": true, "memorabilia": true, "minigame": true,
}

// SetTable is every set, with the family links. Build it once per
// snapshot. It is read-only after that.
type SetTable struct {
	byCode map[string]*SetInfo
	// children maps a base set code to the codes of its products.
	children map[string][]string
	list     []*SetInfo
	// derived marks a table built from the printings, with no set type
	// and no parent link. A resolver on such a table names one set and
	// never a family (D-377).
	derived bool
}

// NewSetTable builds the table from the /sets rows. Nil rows give an
// empty table, which resolves nothing.
func NewSetTable(rows []SetInfo) *SetTable {
	t := &SetTable{
		byCode:   make(map[string]*SetInfo, len(rows)),
		children: map[string][]string{},
	}
	for i := range rows {
		s := rows[i]
		s.Code = strings.ToLower(strings.TrimSpace(s.Code))
		s.ParentCode = strings.ToLower(strings.TrimSpace(s.ParentCode))
		if s.Code == "" {
			continue
		}
		t.byCode[s.Code] = &s
		t.list = append(t.list, &s)
	}
	for _, s := range t.list {
		if s.ParentCode != "" && t.byCode[s.ParentCode] != nil {
			t.children[s.ParentCode] = append(t.children[s.ParentCode], s.Code)
		}
	}
	for k := range t.children {
		sort.Strings(t.children[k])
	}
	// Newest first, then by code, so a listing reads in release order and
	// two runs read the same.
	sort.SliceStable(t.list, func(i, j int) bool {
		if t.list[i].ReleasedAt != t.list[j].ReleasedAt {
			return t.list[i].ReleasedAt > t.list[j].ReleasedAt
		}
		return t.list[i].Code < t.list[j].Code
	})
	return t
}

// Len returns the set count.
func (t *SetTable) Len() int {
	if t == nil {
		return 0
	}
	return len(t.list)
}

// Derived reports whether the table came from the printings rather than
// from the set file. Such a table holds no family link.
func (t *SetTable) Derived() bool { return t != nil && t.derived }

// Get returns one set by code.
func (t *SetTable) Get(code string) (*SetInfo, bool) {
	if t == nil {
		return nil, false
	}
	s, ok := t.byCode[strings.ToLower(strings.TrimSpace(code))]
	return s, ok
}

// All returns every set, newest first. The slice is shared: do not
// modify it.
func (t *SetTable) All() []*SetInfo {
	if t == nil {
		return nil
	}
	return t.list
}

// Family returns the code and every product Scryfall names under it,
// sorted. It walks the whole child tree, because a product can carry a
// product of its own. A token, memorabilia, or minigame child stays out,
// and so does a digital one (D-376).
//
// An unknown code returns itself alone, so a reader who names a set code
// this snapshot does not hold still gets a filter and not a crash.
func (t *SetTable) Family(code string) []string {
	code = strings.ToLower(strings.TrimSpace(code))
	if code == "" {
		return nil
	}
	if t == nil {
		return []string{code}
	}
	seen := map[string]bool{code: true}
	stack := []string{code}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, kid := range t.children[cur] {
			s := t.byCode[kid]
			if s == nil || seen[kid] || s.Digital || familySkipTypes[s.Type] {
				continue
			}
			seen[kid] = true
			stack = append(stack, kid)
		}
	}
	out := make([]string, 0, len(seen))
	for c := range seen {
		out = append(out, c)
	}
	sort.Strings(out)
	return out
}

// FamilyOf expands every code into its family and returns the union,
// sorted and without a repeat.
func (t *SetTable) FamilyOf(codes []string) []string {
	seen := map[string]bool{}
	for _, c := range codes {
		for _, f := range t.Family(c) {
			seen[f] = true
		}
	}
	out := make([]string, 0, len(seen))
	for c := range seen {
		out = append(out, c)
	}
	sort.Strings(out)
	return out
}

// ResolveKind says what the resolver found.
type ResolveKind int

const (
	// ResolveNone says the phrase names no set this snapshot holds.
	ResolveNone ResolveKind = iota
	// ResolveOne says one set family answers the phrase. Codes holds it.
	ResolveOne
	// ResolveMany says two or more base sets answer the phrase.
	// Candidates names them, and the agent asks which one (D-376).
	ResolveMany
)

// Resolution is one answer from Resolve.
type Resolution struct {
	Kind ResolveKind
	// Codes is the whole set family, sorted. Filled for ResolveOne.
	Codes []string
	// Candidates are the base sets a ResolveMany phrase matched, newest
	// first. The question names them.
	Candidates []*SetInfo
}

// maxCandidates bounds the sets a question names. A phrase such as
// "commander" matches over a hundred, and a question can not list them.
const maxCandidates = 4

// minPhrase is the shortest phrase that may match a set name by its
// words. A shorter one is read as a set code alone.
const minPhrase = 3

// tailWords are the words a reader adds after a product name. They carry
// no set, so the resolver drops them before it matches.
var tailWords = map[string]bool{
	"set": true, "sets": true, "block": true, "expansion": true,
	"edition": true, "product": true,
}

// normSetPhrase lowercases a phrase and reduces it to its words. Every
// punctuation mark becomes a space, because a set name holds a colon
// ("Strixhaven: School of Mages") and a reader writes none.
func normSetPhrase(s string) []string {
	var b strings.Builder
	for _, r := range strings.ToLower(strings.TrimSpace(s)) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '\'':
			b.WriteRune(r)
		default:
			b.WriteRune(' ')
		}
	}
	words := strings.Fields(b.String())
	for len(words) > 1 && tailWords[words[len(words)-1]] {
		words = words[:len(words)-1]
	}
	return words
}

// containsWords reports whether needle appears in haystack as a run of
// whole words. "lord of the rings" is inside "the lord of the rings
// tales of middle earth", and "ring" is inside neither.
func containsWords(haystack, needle []string) bool {
	if len(needle) == 0 || len(needle) > len(haystack) {
		return false
	}
	for i := 0; i+len(needle) <= len(haystack); i++ {
		hit := true
		for j, w := range needle {
			if haystack[i+j] != w {
				hit = false
				break
			}
		}
		if hit {
			return true
		}
	}
	return false
}

// Resolve maps the words a reader wrote onto a set family (D-376).
//
// The order is exact set code, then exact set name, then a whole-word
// phrase match over every set name. A candidate whose parent is also a
// candidate leaves the list, because the family expansion brings it back.
// One base set left resolves to its family. Two or more ask.
//
// A digital set is never a candidate: this app offers paper cards only
// (D-306).
func (t *SetTable) Resolve(phrase string) Resolution {
	if t == nil {
		return Resolution{}
	}
	words := normSetPhrase(phrase)
	if len(words) == 0 {
		return Resolution{}
	}
	// A set code is one word, and it never has a space.
	if len(words) == 1 {
		if s, ok := t.byCode[words[0]]; ok && !s.Digital {
			return Resolution{Kind: ResolveOne, Codes: t.Family(s.Code)}
		}
	}
	if len(words) == 1 && len(words[0]) < minPhrase {
		return Resolution{}
	}
	var exact, loose []*SetInfo
	for _, s := range t.list {
		if s.Digital {
			continue
		}
		name := normSetPhrase(s.Name)
		switch {
		case len(name) == len(words) && containsWords(name, words):
			exact = append(exact, s)
		case containsWords(name, words):
			loose = append(loose, s)
		}
	}
	cands := exact
	if len(cands) == 0 {
		cands = loose
	}
	if len(cands) == 0 {
		return Resolution{}
	}
	inList := make(map[string]bool, len(cands))
	for _, s := range cands {
		inList[s.Code] = true
	}
	var roots []*SetInfo
	for _, s := range cands {
		// A product whose base set also matched is not its own answer.
		if s.ParentCode != "" && inList[s.ParentCode] {
			continue
		}
		roots = append(roots, s)
	}
	// A root that holds no playable card of its own is still a root, but
	// a token or memorabilia product never leads a family.
	var kept []*SetInfo
	for _, s := range roots {
		if !familySkipTypes[s.Type] {
			kept = append(kept, s)
		}
	}
	if len(kept) > 0 {
		roots = kept
	}
	switch {
	case len(roots) == 0:
		return Resolution{}
	case len(roots) == 1:
		return Resolution{Kind: ResolveOne, Codes: t.Family(roots[0].Code)}
	}
	if len(roots) > maxCandidates {
		roots = roots[:maxCandidates]
	}
	return Resolution{Kind: ResolveMany, Candidates: roots}
}

// Names reads a code list as set names, for a message the reader sees.
// A code the table does not hold reads as the code itself.
func (t *SetTable) Names(codes []string) []string {
	out := make([]string, 0, len(codes))
	for _, c := range codes {
		if s, ok := t.Get(c); ok {
			out = append(out, s.Name)
			continue
		}
		out = append(out, c)
	}
	return out
}

// SetsFile is the fourth snapshot file. It holds the /sets answer, which
// is the only source of a set family (D-377).
const SetsFile = "sets.json.gz"

// LoadSets parses a stored set file. The file is one JSON array, not
// JSONL: the endpoint answers about a thousand rows, and one array is
// smaller and simpler than a line per row.
func LoadSets(r io.Reader, name string) ([]SetInfo, error) {
	rc, err := maybeGunzip(r, name)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rc.Close() }()
	var out []SetInfo
	if err := json.NewDecoder(rc).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// EncodeSets writes the set rows as the stored file holds them.
func EncodeSets(w io.Writer, rows []SetInfo) error {
	return json.NewEncoder(w).Encode(rows)
}

// SetRowsFrom converts the client rows into the stored shape. One term
// per concept: the store and the resolver read one struct.
func SetRowsFrom(rows []scryfall.SetRow) []SetInfo {
	out := make([]SetInfo, 0, len(rows))
	for _, r := range rows {
		out = append(out, SetInfo{
			Code: strings.ToLower(strings.TrimSpace(r.Code)), Name: r.Name,
			Type: r.SetType, ReleasedAt: r.ReleasedAt,
			ParentCode: strings.ToLower(strings.TrimSpace(r.ParentSetCode)),
			Digital:    r.Digital,
		})
	}
	return out
}

// derivedSetTable builds the fallback table from the printings. It holds
// the code, the name, and the release date, and no set type and no
// parent link. A snapshot stored before the set file existed loads with
// it, so the app still filters by one named set (D-377).
func derivedSetTable(printings []Printing) *SetTable {
	seen := map[string]SetInfo{}
	for _, p := range printings {
		code := strings.ToLower(strings.TrimSpace(p.SetCode))
		if code == "" {
			continue
		}
		s, ok := seen[code]
		if !ok {
			s = SetInfo{Code: code, Name: p.SetName, ReleasedAt: p.ReleasedAt, Digital: true}
		}
		// A set is paper as soon as one paper printing carries it.
		if !p.Digital {
			s.Digital = false
		}
		if s.Name == "" {
			s.Name = p.SetName
		}
		seen[code] = s
	}
	rows := make([]SetInfo, 0, len(seen))
	for _, s := range seen {
		rows = append(rows, s)
	}
	t := NewSetTable(rows)
	t.derived = true
	return t
}
