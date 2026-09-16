package generate

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/profile"
)

// gapNouns names each power floor for the gap note.
var gapNouns = map[string]string{
	profile.KeyTutor:       "tutors",
	profile.KeyFastMana:    "fast mana",
	profile.KeyGameChanger: "Game Changers",
	profile.KeyFinisher:    "finishers, the cards that win the game",
}

// gapNote writes one sentence for each power floor the deck misses, and
// names the cards that close each gap (D-704, D-709). A profile finding
// buys no repair turn (D-613), so the note is what the reader gets. The
// pool cards the deck does not hold come first, then the reserve cards
// outside the pool, each by rate. With a collection, each card reads
// owned or to buy. It reads the profile the check stored, so it names the
// value the finding names.
func (b *Builder) gapNote(req Request, deck *mtgv1.Deck) string {
	if b.profiler == nil || req.Pool == nil {
		return ""
	}
	floors := b.profiler.Bands().PowerFloors(req.Format, req.Power)
	if len(floors) == 0 {
		return ""
	}
	values := map[string]float64{}
	for _, row := range deck.GetProfile().GetFeatures() {
		values[row.GetKey()] = row.GetValue()
	}
	held := map[string]bool{}
	for _, dc := range deck.GetCards() {
		held[dc.GetOracleId()] = true
	}
	for _, id := range deck.GetCommanderOracleIds() {
		held[id] = true
	}
	of := b.profiler.Power()
	var out []string
	for _, key := range profile.PowerKeys {
		floor, ok := floors[key]
		value, measured := values[key]
		if !ok || !measured || value >= floor {
			continue
		}
		picks := gapCards(req.Pool, of, key, held, int(math.Ceil(floor-value)))
		out = append(out, gapSentence(req, key, deck.GetProfile().GetBracket(), value, floor, picks))
	}
	return strings.Join(out, " ")
}

// gapCards picks up to need cards that count toward key and that the deck
// does not hold: the pool cards by rate first, then the reserve by rate.
func gapCards(pool *Pool, of func(*mtgv1.Card) []string, key string, held map[string]bool, need int) []*mtgv1.Card {
	counts := func(c *mtgv1.Card) bool { return !held[c.GetOracleId()] && slices.Contains(of(c), key) }
	var picks []*mtgv1.Card
	for _, name := range pool.Names() {
		if c, ok := pool.Card(name); ok && counts(c) {
			picks = append(picks, c)
		}
	}
	// Names sorts by the alphabet, so a tie of rate keeps that order.
	sort.SliceStable(picks, func(i, j int) bool {
		return pool.Rate(picks[i].GetOracleId()) > pool.Rate(picks[j].GetOracleId())
	})
	for _, c := range pool.Reserve() {
		if counts(c) {
			picks = append(picks, c)
		}
	}
	return picks[:min(need, len(picks))]
}

// gapSentence writes one missed floor and the cards that close it.
func gapSentence(req Request, key string, bracket int32, value, floor float64, picks []*mtgv1.Card) string {
	// A bracket names its own floor for tutors, fast mana, and Game
	// Changers, because the bracket system counts those three (D-704).
	// It counts no finisher, so the finisher sentence names the deck plan
	// and not the bracket. The judge reads a bracket that "wants" a
	// finisher count as a false rule of the game (F-151, D-743).
	s := fmt.Sprintf("Bracket %d wants %g or more %s, and the deck holds %g.", bracket, floor, gapNouns[key], value)
	if key == profile.KeyFinisher {
		s = fmt.Sprintf("This deck plan aims for %g or more %s, and the deck holds %g.", floor, gapNouns[key], value)
	}
	if len(picks) == 0 {
		return s
	}
	names := make([]string, 0, len(picks))
	for _, c := range picks {
		name := c.GetName()
		if req.OracleCounts != nil {
			if req.Pool.OwnedCount(c.GetOracleId()) > 0 {
				name += " (owned)"
			} else {
				name += " (to buy)"
			}
		}
		names = append(names, name)
	}
	return s + " To close the gap, add " + gapList(names) + "."
}

// gapList joins names the way a reader writes a list: "A", "A and B", or
// "A, B, and C".
func gapList(names []string) string {
	switch len(names) {
	case 1:
		return names[0]
	case 2:
		return names[0] + " and " + names[1]
	}
	return strings.Join(names[:len(names)-1], ", ") + ", and " + names[len(names)-1]
}

// powerMarks answers the marks a shortlist line carries: the power floors
// of the bracket that the card counts toward (D-704). A deck with no floor
// marks nothing, and neither does a revision, which reads no deck shape
// (D-283).
func (b *Builder) powerMarks(req Request) func(*mtgv1.Card) []string {
	none := func(*mtgv1.Card) []string { return nil }
	if b.profiler == nil || req.Revision != nil {
		return none
	}
	floors := b.profiler.Bands().PowerFloors(req.Format, req.Power)
	if len(floors) == 0 {
		return none
	}
	of := b.profiler.Power()
	return func(c *mtgv1.Card) []string {
		var out []string
		for _, key := range of(c) {
			if _, ok := floors[key]; ok {
				out = append(out, profile.Mark(key))
			}
		}
		return out
	}
}
