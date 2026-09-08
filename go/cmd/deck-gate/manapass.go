package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/profile"
)

// The free mana-pass lane of PR-33 (F-78, Part 3).
//
// A gate document holds the deck list of every prompt it ran. This lane
// reads those lists, rebuilds the shortlist each one was built from, and
// runs the mana pass over the stored deck. It calls no provider, so it
// costs nothing and it answers the question the plan asks: does the pass
// close the off-band findings of real decks?

// deckLine reads one row of a stored deck list: "- 13 Plains | land | a
// reason". The reason may hold anything, so the row is read from the
// left.
var deckLine = regexp.MustCompile(`^- (\d+) ([^|]+)\|\s*([a-z_]+)\s*\|`)

// promptHead reads the heading of one prompt: "### 1. lifegain
// Commander, any card".
var promptHead = regexp.MustCompile(`^### (\d+)\. `)

// parseDecks reads the deck list of every prompt of a gate document, by
// prompt id.
func parseDecks(path string) (map[int][]generate.Entry, error) {
	f, err := os.Open(path) //nolint:gosec // the caller names a document of this repo
	if err != nil {
		return nil, err
	}
	defer f.Close() //nolint:errcheck // a read-only close
	out := map[int][]generate.Entry{}
	id := 0
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if m := promptHead.FindStringSubmatch(line); m != nil {
			id, _ = strconv.Atoi(m[1])
			continue
		}
		m := deckLine.FindStringSubmatch(line)
		if m == nil || id == 0 {
			continue
		}
		n, err := strconv.Atoi(m[1])
		if err != nil {
			continue
		}
		out[id] = append(out[id], generate.Entry{
			Name:  strings.TrimSpace(m[2]),
			Count: int32(n), //nolint:gosec // a deck list holds small counts
			Role:  m[3],
		})
	}
	return out, sc.Err()
}

// runManaPass reports what the pass does to the stored decks. It returns
// an error only when the document holds no deck at all: a deck the pass
// can not fix is a finding of the report and not a failure of the lane.
func runManaPass(w io.Writer, path string, results []result, idx *cards.Index, prof *profile.Profiler) error {
	stored, err := parseDecks(path)
	if err != nil {
		return err
	}
	if len(stored) == 0 {
		return fmt.Errorf("manapass: %s holds no deck list", path)
	}
	b := generate.NewBuilder(nil, nil, idx, nil, generate.WithProfiler(prof))
	out := func(format string, args ...any) {
		fmt.Fprintf(w, format, args...) //nolint:errcheck // a report to stdout
	}
	out("# The mana pass over %s\n\n", path)
	out("Every deck of the document, rebuilt from its own shortlist. No provider call ran.\n\n")
	out("| # | Prompt | Off band before | Off band after | Steps | What is left |\n")
	out("|---|---|---|---|---|---|\n")
	var before, after, fixed, moved int
	for _, r := range results {
		entries, ok := stored[r.prompt.ID]
		if !ok || r.pool == nil {
			continue
		}
		norm := generate.Normalize(r.pool, entries)
		deck := &mtgv1.Deck{
			Format:             &mtgv1.Format{Id: r.format},
			Power:              r.power,
			CommanderOracleIds: r.commanderIDs,
			Cards:              norm.Cards,
		}
		was := offBand(prof, deck, idx)
		// The request carries the precon, so the lane skips an upgrade
		// exactly as the build does. A precon is a working deck, and a
		// pass over it is a second author (D-249).
		steps := b.FixManaForCheck(generate.Request{
			Format: r.format, Power: r.power, Pool: r.pool, Precon: r.prompt.Precon,
			SessionID: fmt.Sprintf("manapass-%d", r.prompt.ID),
		}, deck)
		now := offBand(prof, deck, idx)
		before += len(was)
		after += len(now)
		if len(now) < len(was) {
			fixed++
		}
		if steps > 0 {
			moved++
		}
		out("| %d | %s | %d | %d | %d | %s |\n",
			r.prompt.ID, r.prompt.Name, len(was), len(now), steps, orNone(strings.Join(now, ", ")))
	}
	out("\n%d off-band features before the pass, %d after. It moved %d decks and improved %d.\n",
		before, after, moved, fixed)
	out("\nAn upgrade keeps its own mana base, so the pass makes no step on one (D-249).\n")
	return nil
}

// offBand names the features of a deck that sit outside their band.
func offBand(prof *profile.Profiler, deck *mtgv1.Deck, idx *cards.Index) []string {
	var out []string
	for _, f := range prof.Measure(deck, idx).GetFeatures() {
		if f.GetOffBand() {
			out = append(out, f.GetKey())
		}
	}
	return out
}

func orNone(s string) string {
	if s == "" {
		return "nothing"
	}
	return s
}
