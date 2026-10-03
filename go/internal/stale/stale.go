// Package stale finds the stored decks that a legality change made
// illegal, and names the rerun that each one takes (D-29, I-1). The
// snapshot job runs the pass after a new legality diff. The rule of the
// rerun is D-1008: a rebuild when the commander, a win condition, or a
// tenth of the nonland cards is no longer legal, and a patch otherwise.
package stale

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// Legalities maps an oracle id to its Scryfall format keys and their
// status, as cards.LoadLegalities reads them from one snapshot.
type Legalities map[string]map[string]string

// ShareLimit is the share of the nonland copies of a deck at which a ban
// asks for a rebuild (D-1008). UNVERIFIED: no measurement backs it yet.
const ShareLimit = 0.10

// Find returns the oracle ids of the deck that are not legal under the
// format key, sorted. It reads the commanders, the main list, the
// sideboard, and the companion. Legal and restricted pass, as the rules
// engine reads them (D-155). A card that the snapshot does not hold is
// not a ban, so it does not count (REV-059).
func Find(d *mtgv1.Deck, key string, legal Legalities) []string {
	if key == "" {
		return nil
	}
	seen := map[string]bool{}
	var out []string
	check := func(id string) {
		if id == "" || seen[id] {
			return
		}
		seen[id] = true
		by, ok := legal[id]
		if !ok {
			return
		}
		switch by[key] {
		case "legal", "restricted":
		default:
			out = append(out, id)
		}
	}
	for _, id := range d.GetCommanderOracleIds() {
		check(id)
	}
	for _, list := range [][]*mtgv1.DeckCard{d.GetCards(), d.GetSideboard()} {
		for _, dc := range list {
			check(dc.GetOracleId())
		}
	}
	check(d.GetCompanionOracleId())
	slices.Sort(out)
	return out
}

// Verdict is the rerun of one stale deck, and the reason the banner
// shows for it.
type Verdict struct {
	Case   mtgv1.RerunCase
	Reason string
}

// Classify names the rerun of a deck from its stale cards (D-1008). A
// banned commander always asks for a rebuild, because no list stands
// without its commander, and the pick row names a new one (D-1021). An
// imported list takes the patch in each other case, because a rebuild
// would replace the list of the user with a generated one (D-1020).
func Classify(d *mtgv1.Deck, ids []string) Verdict {
	if len(ids) == 0 {
		return Verdict{}
	}
	bad := map[string]bool{}
	for _, id := range ids {
		bad[id] = true
	}
	for _, id := range d.GetCommanderOracleIds() {
		if bad[id] {
			return Verdict{
				Case: mtgv1.RerunCase_RERUN_CASE_REBUILD,
				Reason: fmt.Sprintf("%s is no longer legal as your commander. The rerun asks you to pick a new commander, then builds the deck again.",
					commanderName(d, id)),
			}
		}
	}
	names := staleNames(d, bad)
	if d.GetImported() {
		return Verdict{
			Case:   mtgv1.RerunCase_RERUN_CASE_PATCH,
			Reason: fmt.Sprintf("You imported this list, so the rerun replaces %s alone.", names),
		}
	}
	for _, dc := range d.GetCards() {
		if bad[dc.GetOracleId()] && dc.GetRole() == mtgv1.CardRole_CARD_ROLE_WINCON {
			return Verdict{
				Case:   mtgv1.RerunCase_RERUN_CASE_REBUILD,
				Reason: fmt.Sprintf("%s was a win condition of this deck, so the rerun builds the deck again from your conversation.", dc.GetName()),
			}
		}
	}
	if hit, total := nonlandShare(d, bad); total > 0 && float64(hit) >= ShareLimit*float64(total) {
		return Verdict{
			Case: mtgv1.RerunCase_RERUN_CASE_REBUILD,
			Reason: fmt.Sprintf("The change removes %d of the %d nonland cards, so the rerun builds the deck again from your conversation.",
				hit, total),
		}
	}
	return Verdict{
		Case:   mtgv1.RerunCase_RERUN_CASE_PATCH,
		Reason: fmt.Sprintf("The rerun replaces %s alone, and keeps the rest of the deck.", names),
	}
}

// nonlandShare counts the copies of the main list that are not lands, and
// the copies of them that are stale. The command zone and the sideboard
// stay out: the commander has its own rule, and D-1008 reads the deck.
func nonlandShare(d *mtgv1.Deck, bad map[string]bool) (hit, total int) {
	for _, dc := range d.GetCards() {
		if dc.GetRole() == mtgv1.CardRole_CARD_ROLE_LAND {
			continue
		}
		n := max(int(dc.GetCount()), 1)
		total += n
		if bad[dc.GetOracleId()] {
			hit += n
		}
	}
	return hit, total
}

// commanderName reads the name of one commander, from the command zone
// first and the main list next.
func commanderName(d *mtgv1.Deck, id string) string {
	for _, list := range [][]*mtgv1.DeckCard{d.GetCommanders(), d.GetCards()} {
		for _, dc := range list {
			if dc.GetOracleId() == id && dc.GetName() != "" {
				return dc.GetName()
			}
		}
	}
	return "The commander"
}

// staleNames joins the names of the stale cards in the order of the deck,
// with "and" before the last one.
func staleNames(d *mtgv1.Deck, bad map[string]bool) string {
	var names []string
	seen := map[string]bool{}
	for _, list := range [][]*mtgv1.DeckCard{d.GetCards(), d.GetSideboard()} {
		for _, dc := range list {
			if bad[dc.GetOracleId()] && !seen[dc.GetOracleId()] {
				seen[dc.GetOracleId()] = true
				names = append(names, dc.GetName())
			}
		}
	}
	if id := d.GetCompanionOracleId(); bad[id] && !seen[id] {
		names = append(names, "the companion")
	}
	switch len(names) {
	case 0:
		return "the cards that are no longer legal"
	case 1:
		return names[0]
	default:
		return strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]
	}
}

// Apply writes the stale cards and the verdict onto the deck, and clears
// them when ids is empty. It reports whether the deck changed, so the
// pass writes only a deck whose state moved.
func Apply(d *mtgv1.Deck, ids []string, v Verdict) bool {
	stale := len(ids) > 0
	if d.GetStale() == stale && slices.Equal(d.GetStaleOracleIds(), ids) &&
		d.GetRerunCase() == v.Case && d.GetStaleReason() == v.Reason {
		return false
	}
	d.Stale = stale
	d.StaleOracleIds = slices.Clone(ids)
	d.RerunCase = v.Case
	d.StaleReason = v.Reason
	return true
}

// Store is the deck store the pass reads and writes.
type Store interface {
	Scan(ctx context.Context, fn func(uid string, d *mtgv1.Deck) error) error
	Mark(ctx context.Context, uid, id string, fn func(d *mtgv1.Deck) bool) (bool, error)
}

// Result counts one pass.
type Result struct {
	// Read is each deck the scan read.
	Read int
	// Stale is each deck that holds a card no longer legal now.
	Stale int
	// Written is each deck whose stored state changed: newly stale, a
	// new verdict, or cleared after an unban.
	Written int
	// Hit holds, for each user, the decks that this pass made stale or
	// that took a new illegal card. They get the push of D-1088. An
	// unban hits no deck.
	Hit map[string][]*mtgv1.Deck
}

// NewCard says whether after holds an oracle id that before does not.
func NewCard(before, after []string) bool {
	for _, id := range after {
		if !slices.Contains(before, id) {
			return true
		}
	}
	return false
}

// Pass reads every stored deck, and marks or clears its stale state
// against the legalities of one snapshot. keyOf gives the Scryfall key
// of a format, and "" for a format with no legality check (D-3). The
// write reads the deck again inside its transaction (Store.Mark). A
// pass that fails still returns the decks it hit, because their marks
// are stored and the next pass does not hit them again.
func Pass(ctx context.Context, store Store, keyOf func(mtgv1.FormatId) string, legal Legalities, log *slog.Logger) (Result, error) {
	var res Result
	verdict := func(d *mtgv1.Deck) ([]string, Verdict) {
		ids := Find(d, keyOf(d.GetFormat().GetId()), legal)
		return ids, Classify(d, ids)
	}
	err := store.Scan(ctx, func(uid string, d *mtgv1.Deck) error {
		res.Read++
		ids, v := verdict(d)
		if len(ids) > 0 {
			res.Stale++
		}
		if !Apply(d, ids, v) {
			return nil
		}
		var hit *mtgv1.Deck
		changed, err := store.Mark(ctx, uid, d.GetId(), func(cur *mtgv1.Deck) bool {
			before := slices.Clone(cur.GetStaleOracleIds())
			ids, v := verdict(cur)
			hit = nil
			if !Apply(cur, ids, v) {
				return false
			}
			if NewCard(before, ids) {
				hit = cur
			}
			return true
		})
		if err != nil {
			return fmt.Errorf("deck %s/%s: %w", uid, d.GetId(), err)
		}
		if changed {
			res.Written++
			if hit != nil {
				if res.Hit == nil {
					res.Hit = map[string][]*mtgv1.Deck{}
				}
				res.Hit[uid] = append(res.Hit[uid], hit)
			}
			log.InfoContext(ctx, "stale pass: deck marked", "deck", d.GetId(), "stale", len(ids) > 0,
				"cards", len(ids), "case", v.Case.String())
		}
		return nil
	})
	return res, err
}
