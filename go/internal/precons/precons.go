// Package precons holds the preconstructed decks a user can ask to
// upgrade (D-113, D-218, OQ-40).
//
// D-218 sets the share of a precon a built deck must keep, and the rule
// could not run: nothing in the repo listed the cards of a precon, and
// the question workflow held only the first card the user named (D-240).
//
// A decklist is one file in decks/, in the deck export format the corpus
// names in section 9. The file name is the product name, in lower case
// with hyphens. Adding a precon is adding a file (D-247).
package precons

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
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

// Load reads every decklist and resolves it against the card index. It
// returns a Set and the precons that did not fully resolve, so a caller
// can refuse to trust one (D-247).
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
		// A precon file may carry a sideboard, and it is not part of the
		// hundred. From Cute to Brute lists five Secret Lair cards there,
		// and Tricky Terrain lists an alternate commander. The share rule
		// measures the deck, so the loader splits the sections. The shared
		// parser drops the headers, because a 60-card import wants both
		// halves in one list (D-247).
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
			if ok && strings.Contains(strings.ToLower(c.GetTypeLine()), "land") {
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

// Find reads the user's words and returns the precon they named. It
// matches the whole product name inside the message, longest name first,
// so "Avengers Assemble" wins over a precon called "Avengers".
//
// A user says "upgrade my Avengers Assemble precon", and the classifier
// reports a card name for that phrase, not a product (D-240). The words
// are the reliable source.
func (s *Set) Find(words string) (*Precon, bool) {
	if s == nil {
		return nil, false
	}
	low := strings.ToLower(words)
	var best *Precon
	for _, p := range s.list {
		name := strings.ToLower(p.Name)
		if !strings.Contains(low, name) {
			continue
		}
		if best == nil || len(name) > len(strings.ToLower(best.Name)) {
			best = p
		}
	}
	return best, best != nil
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
// returns the deck text and how many cards the sideboard holds.
func splitSideboard(text string) (deck string, sideboard int) {
	lines := strings.Split(text, "\n")
	cut := len(lines)
	for i, l := range lines {
		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(l)), "// sideboard") {
			cut = i
			break
		}
	}
	for _, l := range lines[min(cut+1, len(lines)):] {
		var n int
		if _, err := fmt.Sscanf(strings.TrimSpace(l), "%d ", &n); err == nil {
			sideboard += n
		}
	}
	return strings.Join(lines[:cut], "\n"), sideboard
}

// titleOf turns a slug into a product name: "avengers-assemble" reads
// "Avengers Assemble".
func titleOf(slug string) string {
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, " ")
}

// Kept counts how many of the precon's cards a deck still holds.
func (p *Precon) Kept(deck *mtgv1.Deck) int {
	if p == nil {
		return 0
	}
	have := map[string]bool{}
	for _, c := range deck.GetCards() {
		have[c.GetOracleId()] = true
	}
	for _, id := range deck.GetCommanderOracleIds() {
		have[id] = true
	}
	n := 0
	for _, id := range p.OracleIDs {
		if have[id] {
			n++
		}
	}
	return n
}
