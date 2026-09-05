package precons

import (
	"sort"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
)

// Product is one row of the MTGJSON precon table, indexed for the precon
// exclusion (D-407, D-408). The commanders count as cards of the product.
type Product struct {
	// Key is the product's id in the table, the MTGJSON file name.
	Key string
	// Name is the product name, for example "Avengers Assemble".
	Name        string
	Code        string
	Type        string
	ReleaseDate string
	// counts is the copies per Oracle id, printings merged.
	counts map[string]int32
	// printings is the copies per Scryfall id.
	printings map[string]int32
	// oracle is the Oracle id of each printing, so the ownership check
	// can ask whether a printing is a basic land (D-523).
	oracle map[string]string
}

// Counts is the copies per Oracle id, the commanders included.
func (p *Product) Counts() map[string]int32 {
	out := make(map[string]int32, len(p.counts))
	for k, v := range p.counts {
		out[k] = v
	}
	return out
}

// OwnedWhole reports whether a collection holds every nonbasic printing
// of the product with its count (D-408, D-523). printings is the
// collection's copies per Scryfall id. isBasic says which Oracle ids are
// basic lands, and the check skips their printings: a ManaBox deck binder
// can omit the basic lands, and the exclusion never removes one (D-37),
// so the check loses nothing. A nil isBasic reads every printing. A
// product with no nonbasic printing is never owned.
func (p *Product) OwnedWhole(printings map[string]int32, isBasic func(oracleID string) bool) bool {
	checked := 0
	for id, n := range p.printings {
		if isBasic != nil && isBasic(p.oracle[id]) {
			continue
		}
		checked++
		if printings[id] < n {
			return false
		}
	}
	return checked > 0
}

// SameCards reports whether two products hold the same cards by Oracle
// id and count, as a deck and its Collector's Edition do.
func (p *Product) SameCards(q *Product) bool {
	if len(p.counts) != len(q.counts) {
		return false
	}
	for id, n := range p.counts {
		if q.counts[id] != n {
			return false
		}
	}
	return true
}

// Table is the indexed precon table of one MTGJSON version.
type Table struct {
	Version string
	byKey   map[string]*Product
	list    []*Product
	lower   []string
}

// NewTable indexes the rows of a stored table. The list keeps the
// table's order, release date then name.
func NewTable(version string, rows []meta.Precon) *Table {
	t := &Table{Version: version, byKey: map[string]*Product{}}
	for i := range rows {
		r := &rows[i]
		p := &Product{
			Key: r.Key(), Name: r.Name, Code: r.Code, Type: r.Type, ReleaseDate: r.ReleaseDate,
			counts: map[string]int32{}, printings: map[string]int32{}, oracle: map[string]string{},
		}
		for _, list := range [][]meta.PreconCard{r.Commanders, r.Cards} {
			for _, c := range list {
				if c.Count <= 0 {
					continue
				}
				if c.OracleID != "" {
					p.counts[c.OracleID] += int32(c.Count)
				}
				if c.ScryfallID != "" {
					p.printings[c.ScryfallID] += int32(c.Count)
					p.oracle[c.ScryfallID] = c.OracleID
				}
			}
		}
		t.byKey[p.Key] = p
		t.list = append(t.list, p)
		t.lower = append(t.lower, plainWords(p.Name))
	}
	return t
}

// Get reads one product by key.
func (t *Table) Get(key string) (*Product, bool) {
	if t == nil {
		return nil, false
	}
	p, ok := t.byKey[key]
	return p, ok
}

// Len counts the products.
func (t *Table) Len() int {
	if t == nil {
		return 0
	}
	return len(t.list)
}

// All lists every product in table order.
func (t *Table) All() []*Product {
	if t == nil {
		return nil
	}
	return append([]*Product(nil), t.list...)
}

// Owned lists the products a collection holds whole (D-408), in table
// order. isBasic is the basic land test of OwnedWhole (D-523).
func (t *Table) Owned(printings map[string]int32, isBasic func(oracleID string) bool) []*Product {
	if t == nil || len(printings) == 0 {
		return nil
	}
	var out []*Product
	for _, p := range t.list {
		if p.OwnedWhole(printings, isBasic) {
			out = append(out, p)
		}
	}
	return out
}

// Match is the answer to Resolve.
type Match struct {
	// Products are the products the phrase named. A deck and its
	// Collector's Edition hold the same cards, so both answer, and the
	// exclusion counts them once.
	Products []*Product
	// Options are the names the question offers when the phrase names
	// nothing, or names products with different cards.
	Options []string
}

// OK reports whether the phrase named a product.
func (m Match) OK() bool { return len(m.Products) > 0 }

// Names lists the product names once each, in table order.
func (m Match) Names() []string {
	var out []string
	seen := map[string]bool{}
	for _, p := range m.Products {
		if !seen[p.Name] {
			seen[p.Name] = true
			out = append(out, p.Name)
		}
	}
	return out
}

// Keys lists the product keys.
func (m Match) Keys() []string {
	out := make([]string, 0, len(m.Products))
	for _, p := range m.Products {
		out = append(out, p.Key)
	}
	return out
}

// maxOptions caps the names a question offers for an unknown phrase.
const maxOptions = 5

// Resolve maps the words a reader wrote onto the products they name.
// The whole phrase is tried first, so "Deck A" names the product of that
// name. Then the words such as "my", "the", "precon", and "deck" around
// the name are dropped, so "my Avengers Assemble precon" reads as
// "Avengers Assemble". The exact name wins, then the longest name inside
// the phrase, then a name of which the phrase is a part. One word alone
// is a part of nothing: "Power" is a word of many names.
func (t *Table) Resolve(phrase string) Match {
	if t == nil {
		return Match{}
	}
	full := plainWords(phrase)
	stripped := stripFiller(full)
	tried := ""
	for _, p := range []string{full, stripped} {
		if p == "" || p == tried {
			continue
		}
		tried = p
		if m, ok := t.resolveWords(p); ok {
			return m
		}
	}
	if stripped == "" {
		return Match{}
	}
	return Match{Options: t.nearNames(stripped)}
}

// resolveWords answers one normalized phrase. ok is false when no name
// matches at all.
func (t *Table) resolveWords(p string) (Match, bool) {
	var exact, within, holding []*Product
	longest := 0
	several := strings.Contains(p, " ")
	for i, name := range t.lower {
		switch {
		case name == p:
			exact = append(exact, t.list[i])
		case containsWords(p, name):
			if len(name) > longest {
				longest, within = len(name), nil
			}
			if len(name) == longest {
				within = append(within, t.list[i])
			}
		case several && containsWords(name, p):
			holding = append(holding, t.list[i])
		}
	}
	for _, hits := range [][]*Product{exact, within, holding} {
		if len(hits) > 0 {
			return group(hits), true
		}
	}
	return Match{}, false
}

// group answers one phrase's hits. Products with the same cards are one
// answer. Products with different cards are a question.
func group(hits []*Product) Match {
	for _, h := range hits[1:] {
		if !h.SameCards(hits[0]) {
			var options []string
			for _, p := range hits {
				options = append(options, p.Name+" ("+p.Code+")")
			}
			return Match{Options: options}
		}
	}
	return Match{Products: hits}
}

// nearNames lists product names that share a word with the phrase, for
// the question about a name the table does not hold.
func (t *Table) nearNames(p string) []string {
	words := map[string]bool{}
	for _, w := range strings.Fields(p) {
		if len(w) >= 4 && !fillerWords[w] {
			words[w] = true
		}
	}
	if len(words) == 0 {
		return nil
	}
	var out []string
	seen := map[string]bool{}
	for i, name := range t.lower {
		for _, w := range strings.Fields(name) {
			if words[w] && !seen[t.list[i].Name] {
				seen[t.list[i].Name] = true
				out = append(out, t.list[i].Name)
				break
			}
		}
		if len(out) == maxOptions {
			break
		}
	}
	return out
}

// fillerWords are the words a reader puts around a product name.
var fillerWords = map[string]bool{
	"my": true, "the": true, "a": true, "an": true, "our": true,
	"precon": true, "precons": true, "deck": true, "decks": true,
	"commander": true, "preconstructed": true, "starter": true, "own": true,
}

// plainWords lowers a phrase, strips its punctuation, and joins its
// words with one space. A product name goes through it too, so
// "Collector's Edition" and "collectors edition" read the same.
func plainWords(phrase string) string {
	low := strings.ToLower(strings.TrimSpace(phrase))
	low = strings.Map(func(r rune) rune {
		switch r {
		case '"', '\'', '\u2019', '.', ',', '!', '?', '(', ')', '[', ']', ':':
			return -1
		}
		return r
	}, low)
	return strings.Join(strings.Fields(low), " ")
}

// stripFiller drops the filler words at both ends of a plain phrase.
func stripFiller(p string) string {
	words := strings.Fields(p)
	for len(words) > 0 && fillerWords[words[0]] {
		words = words[1:]
	}
	for len(words) > 0 && fillerWords[words[len(words)-1]] {
		words = words[:len(words)-1]
	}
	return strings.Join(words, " ")
}

// containsWords reports whether needle is inside hay on word borders.
func containsWords(hay, needle string) bool {
	return strings.Contains(" "+hay+" ", " "+needle+" ")
}

// Exclude subtracts the products' copies from the owned counts per
// Oracle id, so a surplus copy stays usable (D-408). A card with no copy
// left is excluded from the pool, and so is every card of a product when
// owned is nil, which is the any-card pool. Products with the same cards
// count once. A basic land is never excluded (D-37), and isBasic says
// which Oracle ids are basics. The excluded ids come back sorted.
func Exclude(products []*Product, owned map[string]int32, isBasic func(oracleID string) bool) (reduced map[string]int32, excluded []string) {
	reduced = make(map[string]int32, len(owned))
	for k, v := range owned {
		reduced[k] = v
	}
	take := map[string]int32{}
	var counted []*Product
	for _, p := range products {
		dup := false
		for _, c := range counted {
			if c.SameCards(p) {
				dup = true
				break
			}
		}
		if dup {
			continue
		}
		counted = append(counted, p)
		for id, n := range p.counts {
			take[id] += n
		}
	}
	for id, n := range take {
		if isBasic != nil && isBasic(id) {
			continue
		}
		left := reduced[id] - n
		if left <= 0 {
			delete(reduced, id)
			excluded = append(excluded, id)
			continue
		}
		reduced[id] = left
	}
	sort.Strings(excluded)
	return reduced, excluded
}
