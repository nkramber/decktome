package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/collections"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/gcpenv"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/precons"
)

// The calibration mode writes decks whose bracket comes from outside the
// builder, in the shape readDecks parses, and it calls no model (M-15,
// D-698). A precon stands for bracket 2, the judge's "near the strength
// of a preconstructed deck". A top-finish cEDH list stands for bracket 5.
// The re-judge mode then reads the document.

const (
	calibrationLists      = 12
	calibrationMinPlayers = 32
	calibrationMaxPlace   = 4
)

// calibrationDeck is one deck of known bracket.
type calibrationDeck struct {
	bracket   int32
	commander *mtgv1.Card
	theme     string
	cards     []*mtgv1.DeckCard
}

// add counts copies of a card, one line for each oracle id.
func (d *calibrationDeck) add(c *mtgv1.Card, n int32) {
	for _, dc := range d.cards {
		if dc.GetOracleId() == c.GetOracleId() {
			dc.Count += n
			return
		}
	}
	d.cards = append(d.cards, &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: n})
}

func runCalibrate(out string) error {
	if os.Getenv("CARDS_SNAPSHOT_DIR") == "" {
		return fmt.Errorf("calibrate: set CARDS_SNAPSHOT_DIR, because the cEDH lists come from the local meta store beside the snapshot")
	}
	if _, err := os.Stat(out); err == nil {
		return fmt.Errorf("calibrate: %s exists, so write to a new file (D-65)", out)
	}
	ctx := context.Background()
	idx, err := gatekit.LoadSnapshot(ctx, gatekit.Quiet())
	if err != nil {
		return err
	}
	lists, err := precons.Decklists()
	if err != nil {
		return err
	}
	store, err := gcpenv.MetaStore(ctx, "", nil)
	if err != nil {
		return err
	}
	all, err := meta.AllLists(ctx, store, "commander")
	if err != nil {
		return err
	}
	decks, err := calibrationDecks(idx, lists, all, calibrationLists)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	writeCalibration(&b, idx.AsOf, decks)
	back, err := readDecks(bytes.NewReader(b.Bytes()), idx)
	if err != nil {
		return fmt.Errorf("calibrate: the document does not read back: %w", err)
	}
	if len(back) != len(decks) {
		return fmt.Errorf("calibrate: %d decks read back, want %d", len(back), len(decks))
	}
	if err := os.WriteFile(out, b.Bytes(), 0o600); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "wrote %d calibration decks to %s\n", len(decks), out)
	return nil
}

// calibrationDecks resolves the precons and picks the cEDH lists. A
// precon that does not resolve whole is an error, because the set is
// small and fixed. A cEDH list with a name the index does not know is
// skipped, and the next list takes its place. Fewer qualifying lists
// than the lane wants is an error, so a thin set never writes.
func calibrationDecks(idx *cards.Index, lists []precons.Decklist, all []meta.List, wantCEDH int) ([]calibrationDeck, error) {
	if len(lists) == 0 {
		return nil, fmt.Errorf("calibrate: no precon decklist")
	}
	var out []calibrationDeck
	for _, d := range lists {
		commander, ok := idx.ByName(d.Commander)
		if !ok {
			return nil, fmt.Errorf("calibrate: precon %s: no card named %q", d.Slug, d.Commander)
		}
		entries, unresolved := collections.Resolve(d.Rows, idx)
		if len(unresolved) > 0 {
			return nil, fmt.Errorf("calibrate: precon %s: %d rows do not resolve", d.Slug, len(unresolved))
		}
		deck := calibrationDeck{bracket: 2, commander: commander, theme: "precon " + d.Name}
		for _, e := range entries {
			c, ok := idx.ByOracleID(e.GetOracleId())
			if !ok {
				return nil, fmt.Errorf("calibrate: precon %s: no card for oracle id %s", d.Slug, e.GetOracleId())
			}
			deck.add(c, e.GetQuantity())
		}
		out = append(out, deck)
	}
	cedh := pickCEDH(idx, all, wantCEDH)
	if len(cedh) < wantCEDH {
		return nil, fmt.Errorf("calibrate: %d cEDH lists qualify, and the lane wants %d", len(cedh), wantCEDH)
	}
	return append(out, cedh...), nil
}

// pickCEDH takes the newest top-finish cEDH list of each commander: tier
// great, placed calibrationMaxPlace or better in a field of
// calibrationMinPlayers or more, one commander, and 99 cards.
func pickCEDH(idx *cards.Index, all []meta.List, want int) []calibrationDeck {
	var pool []meta.List
	for _, l := range all {
		if l.Source == meta.SourceTopdeck && l.Tier == "great" &&
			l.Placement >= 1 && l.Placement <= calibrationMaxPlace &&
			l.Players >= calibrationMinPlayers && len(l.Commanders) == 1 && l.Size() == 99 {
			pool = append(pool, l)
		}
	}
	sort.SliceStable(pool, func(i, j int) bool {
		if pool[i].Date != pool[j].Date {
			return pool[i].Date > pool[j].Date
		}
		if pool[i].Placement != pool[j].Placement {
			return pool[i].Placement < pool[j].Placement
		}
		return pool[i].ID < pool[j].ID
	})
	seen := map[string]bool{}
	var out []calibrationDeck
	for _, l := range pool {
		if len(out) == want {
			break
		}
		commander, ok := idx.ByName(l.Commanders[0])
		if !ok || seen[commander.GetOracleId()] {
			continue
		}
		deck := calibrationDeck{bracket: 5, commander: commander,
			theme: fmt.Sprintf("cedh %s, place %d of %d", l.Date, l.Placement, l.Players)}
		whole := true
		for _, lc := range l.Cards {
			c, ok := listCard(idx, lc)
			if !ok || c.GetOracleId() == commander.GetOracleId() {
				whole = false
				break
			}
			deck.add(c, int32(lc.Count))
		}
		if !whole {
			continue
		}
		seen[commander.GetOracleId()] = true
		out = append(out, deck)
	}
	return out
}

func listCard(idx *cards.Index, lc meta.Card) (*mtgv1.Card, bool) {
	if lc.OracleID != "" {
		if c, ok := idx.ByOracleID(lc.OracleID); ok {
			return c, true
		}
	}
	return idx.ByName(lc.Name)
}

// writeCalibration writes the decks as readDecks parses them: a header
// with the id, the bracket, the commander, and the theme, and a card list
// under "Cards:".
func writeCalibration(w io.Writer, asOf time.Time, decks []calibrationDeck) {
	precon, cedh := 0, 0
	for _, d := range decks {
		if d.bracket == 2 {
			precon++
			continue
		}
		cedh++
	}
	_, _ = fmt.Fprintf(w, "# M-15 judge calibration decks\n\n")
	_, _ = fmt.Fprintf(w, "Card snapshot: %s. %d precons of the repository stand for bracket 2, and %d cEDH lists of the local meta store stand for bracket 5 (D-698).\n\n",
		asOf.Format("2006-01-02"), precon, cedh)
	_, _ = fmt.Fprintf(w, "A cEDH list is a Topdeck.gg list of tier great, with one commander and 99 cards. It placed %d or better in a field of %d or more. The newest list of each commander stands. The pick skips a list with a name that the card index does not know.\n\n",
		calibrationMaxPlace, calibrationMinPlayers)
	_, _ = fmt.Fprintf(w, "The re-judge mode of `bracket-gate` reads this document. No model call wrote it.\n")
	for i, d := range decks {
		_, _ = fmt.Fprintf(w, "\n### %d. Bracket %d, %s, %s\n\nCards:\n\n", i+1, d.bracket, d.commander.GetName(), d.theme)
		for _, c := range d.cards {
			_, _ = fmt.Fprintf(w, "- %d %s\n", c.GetCount(), c.GetName())
		}
	}
}
