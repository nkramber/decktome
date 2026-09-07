// Package precons holds the preconstructed decks a user can ask to
// upgrade (D-113, D-218). The share rule of D-218 reads these lists.
//
// A decklist is one file in decks/, in the deck export format the corpus
// names in section 9. The file name is a slug of the product name, in
// lower case with hyphens, and displayNames maps a slug to the name a
// person uses. Adding a precon is adding a file (D-247).
package precons

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"slices"
	"sort"
	"strings"

	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/collections"
)

//go:embed decks/*.txt
var files embed.FS

// Precon is one preconstructed deck.
type Precon struct {
	// Slug is the file name without its suffix, for example
	// "avengers-assemble".
	Slug string
	// Name is what a person calls it, for example "Avengers Assemble".
	Name string
	// OracleIDs are the cards of the deck, the commander included. A card
	// the snapshot does not know is left out and reported.
	OracleIDs []string
	// Unresolved counts the rows the card index could not answer. A
	// precon with any is not trustworthy for the share rule.
	Unresolved int
	// Cards counts the copies the deck holds, before resolution. A
	// sideboard is not counted: it is not part of the hundred.
	Cards int
	// Sideboard counts the bonus cards the file lists after the deck.
	Sideboard int
	// Lands counts the land cards the deck holds, copies included. The
	// upgrade prompt names it so the mana base survives, and a guess at
	// how many basics a precon runs is not good enough: the list says
	// (D-251).
	Lands int
}

// Set is every precon the build can read.
type Set struct {
	byName map[string]*Precon
	list   []*Precon
}

// Load reads every decklist and resolves it against the card index. A
// precon that did not fully resolve stays in the Set with its Unresolved
// count, and Set.Unresolved lists them, so a caller can refuse to trust
// one (D-247).
func Load(idx *cards.Index) (*Set, error) {
	if idx == nil {
		return nil, fmt.Errorf("precons: no card index")
	}
	entries, err := fs.ReadDir(files, "decks")
	if err != nil {
		return nil, err
	}
	s := &Set{byName: map[string]*Precon{}}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".txt") {
			continue
		}
		raw, err := files.ReadFile(path.Join("decks", e.Name()))
		if err != nil {
			return nil, err
		}
		// A sideboard is not part of the hundred, so the loader splits it
		// off before the shared parser drops the headers (D-247).
		main, side := splitSideboard(string(raw))
		rows, bad, err := collections.ParseArenaText(strings.NewReader(main))
		if err != nil {
			return nil, fmt.Errorf("precons: %s: %w", e.Name(), err)
		}
		slug := strings.TrimSuffix(e.Name(), ".txt")
		p := &Precon{Slug: slug, Name: titleOf(slug), Unresolved: len(bad), Sideboard: side}
		for _, r := range rows {
			p.Cards += r.Quantity
		}
		resolved, unresolved := collections.Resolve(rows, idx)
		p.Unresolved += len(unresolved)
		for _, entry := range resolved {
			c, ok := idx.ByOracleID(entry.GetOracleId())
			if ok && slices.Contains(c.GetCardTypes(), "Land") {
				p.Lands += int(entry.GetQuantity())
			}
		}
		seen := map[string]bool{}
		for _, entry := range resolved {
			id := entry.GetOracleId()
			if id == "" || seen[id] {
				continue
			}
			seen[id] = true
			p.OracleIDs = append(p.OracleIDs, id)
		}
		sort.Strings(p.OracleIDs)
		s.byName[strings.ToLower(p.Name)] = p
		s.byName[slug] = p
		s.list = append(s.list, p)
	}
	sort.Slice(s.list, func(i, j int) bool { return s.list[i].Slug < s.list[j].Slug })
	return s, nil
}

// All lists every precon, by slug.
func (s *Set) All() []*Precon {
	if s == nil {
		return nil
	}
	return append([]*Precon(nil), s.list...)
}

// Unresolved lists the precons with a row the card index could not
// answer. The share rule of D-218 must not trust one: a list that lost
// cards gives a wrong share.
func (s *Set) Unresolved() []*Precon {
	var out []*Precon
	for _, p := range s.All() {
		if p.Unresolved > 0 {
			out = append(out, p)
		}
	}
	return out
}

// Find reads the user's words and returns the precon they named. A
// precon matches when its name is inside the message, or when every
// word of the message is a word of the name and at least two words
// match, which is how a short phrase such as "riders of rohan" reaches
// a longer product name. One word alone names nothing: "power" and "of"
// are title words of many things. The longest name wins, so "Avengers
// Assemble" beats a precon called "Avengers". The words are the source,
// because the classifier reports a card name for a product (D-240).
func (s *Set) Find(words string) (*Precon, bool) {
	if s == nil {
		return nil, false
	}
	low := strings.ToLower(words)
	phrase := strings.Fields(low)
	var best *Precon
	for _, p := range s.list {
		name := strings.ToLower(p.Name)
		if !strings.Contains(low, name) && !wordsOf(phrase, strings.Fields(name)) {
			continue
		}
		if best == nil || len(name) > len(strings.ToLower(best.Name)) {
			best = p
		}
	}
	return best, best != nil
}

// wordsOf reports whether every word of phrase is one of the title
// words, and at least two distinct title words match. A one-word phrase
// names nothing unless it is the whole title.
func wordsOf(phrase, title []string) bool {
	if len(phrase) == 0 {
		return false
	}
	matched := map[string]bool{}
	for _, w := range phrase {
		w = strings.Trim(w, ",.!?'\"")
		if !slices.Contains(title, w) {
			return false
		}
		matched[w] = true
	}
	return len(matched) >= 2 || len(matched) == len(title)
}

// Get reads one precon by slug or by name.
func (s *Set) Get(key string) (*Precon, bool) {
	if s == nil {
		return nil, false
	}
	p, ok := s.byName[strings.ToLower(strings.TrimSpace(key))]
	return p, ok
}

// splitSideboard divides a deck export at its sideboard header and
// returns the deck text and how many cards the sideboard holds. The
// header is "// SIDEBOARD" in the ManaBox export, and the bare
// "Sideboard" that collections.ParseArenaText also reads. The sideboard
// ends at the next section header, so a later "Commander" section is
// not counted as sideboard.
func splitSideboard(text string) (deck string, sideboard int) {
	lines := strings.Split(text, "\n")
	cut := len(lines)
	for i, l := range lines {
		if sectionHeader(l) == "sideboard" {
			cut = i
			break
		}
	}
	for _, l := range lines[min(cut+1, len(lines)):] {
		if h := sectionHeader(l); h != "" && h != "sideboard" {
			break
		}
		var n int
		if _, err := fmt.Sscanf(strings.TrimSpace(l), "%d ", &n); err == nil {
			sideboard += n
		}
	}
	return strings.Join(lines[:cut], "\n"), sideboard
}

// sectionHeader reads a deck export section header in lower case, with
// or without the "//" prefix, or "" for any other line.
func sectionHeader(line string) string {
	h := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "//")))
	switch h {
	case "deck", "sideboard", "commander", "about":
		return h
	}
	return ""
}

// displayNames maps a file slug to the product name a person uses. A
// slug that is not here reads in Title Case, and a guessed name would be
// a mistake the user sees, so a name joins once it is verified. The
// MTGJSON table and the Wizards decklist page verified "Limit Break" and
// "Blight Curse" on 2026-09-03 (D-498).
var displayNames = map[string]string{
	"avengers-assemble":                 "Avengers Assemble",
	"ff-cloud":                          "Limit Break",
	"from-cute-to-brute":                "From Cute to Brute",
	"lorwyn-blight-curse":               "Blight Curse",
	"goblin-storm":                      "Goblin Storm",
	"living-energy":                     "Living Energy",
	"lotr-riders-of-rohan":              "Riders of Rohan",
	"tricky-terrain-collectors-edition": "Tricky Terrain",
	"turtle-power":                      "Turtle Power",
}

// titleOf gives the display name of a slug. A slug displayNames does not
// hold reads in Title Case: "ff-cloud" reads "Ff Cloud".
func titleOf(slug string) string {
	if name, ok := displayNames[slug]; ok {
		return name
	}
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, " ")
}
