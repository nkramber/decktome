package generate

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
)

// CodeOwnedTwin reports the cards the builder played in place of a card
// to buy, because the user owns a twin of it (D-1226). It is an INFO:
// the deck does the same thing, and it costs less.
const CodeOwnedTwin = "owned_twin"

// swapOwnedTwins plays an owned twin in place of each copy the deck must
// buy, under owned-first, and returns one line for each swap (F-232,
// D-1226). The model chose Grizzly Bears and Bear Cub over the owned
// Balduvian Bears, and Llanowar Elves over the owned Fyndhorn Elves. The
// shortlist held both cards, and no line of the prompt said which to
// prefer. The pass needs no model call, and its rule reads the printed
// card alone (candidates.TwinBeats).
//
// A card the user named stays: a commander, a locked card, a card a
// revision keeps, and a precon card, whose share D-218 counts. A twin
// comes from the pool, so the model can name it in a repair turn, and it
// takes no more copies than the user owns or the format allows.
func swapOwnedTwins(deck *mtgv1.Deck, req Request) []string {
	if req.PoolRule != mtgv1.PoolRule_POOL_RULE_OWNED_FIRST || req.Pool == nil {
		return nil
	}
	keep := map[string]bool{}
	for _, ids := range [][]string{req.Commanders, req.Locked, req.PreconOracleIDs} {
		for _, id := range ids {
			keep[id] = true
		}
	}
	keepNames := map[string]bool{}
	if req.Revision != nil {
		for _, n := range req.Revision.Keep {
			keepNames[candidates.FoldName(n)] = true
		}
	}
	twins := map[string][]*mtgv1.Card{}
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || req.Pool.OwnedCount(c.GetOracleId()) <= 0 {
			continue
		}
		if k, ok := candidates.TwinKey(c); ok {
			twins[k] = append(twins[k], c)
		}
	}
	if len(twins) == 0 {
		return nil
	}
	for _, list := range twins {
		sort.Slice(list, func(i, j int) bool { return list[i].GetName() < list[j].GetName() })
	}
	limit := int32(sixtyCopies)
	if req.Format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		limit = 1
	}
	var lines []string
	swapList := func(list, other []*mtgv1.DeckCard) []*mtgv1.DeckCard {
		for i := 0; i < len(list); i++ {
			dc := list[i]
			id := dc.GetOracleId()
			buy := dc.GetCount() - req.Pool.OwnedCount(id)
			if buy <= 0 || keep[id] || keepNames[candidates.FoldName(dc.GetName())] {
				continue
			}
			bought, ok := req.Pool.ByOracleID(id)
			if !ok {
				continue
			}
			k, ok := candidates.TwinKey(bought)
			if !ok {
				continue
			}
			last := ""
			for _, tw := range twins[k] {
				if buy <= 0 {
					break
				}
				if !candidates.TwinBeats(tw, bought) {
					continue
				}
				held := listCopies(list, tw.GetOracleId()) + listCopies(other, tw.GetOracleId())
				n := min(buy, req.Pool.OwnedCount(tw.GetOracleId())-held, limit-held)
				if n <= 0 {
					continue
				}
				reason := strings.ReplaceAll(dc.GetReason(), bought.GetName(), tw.GetName())
				if j := indexOf(list, tw.GetOracleId()); j >= 0 {
					list[j].Count += n
				} else {
					list = append(list, deckCard(req.Pool, tw, n, dc.GetRole(), reason))
				}
				dc.Count -= n
				buy -= n
				last = tw.GetName()
				lines = append(lines, fmt.Sprintf("%s for %s", copiesName(n, tw.GetName()), bought.GetName()))
			}
			// A summary that names a card the deck does not hold misleads
			// the reader, so it names the twin that took the last copy.
			if dc.GetCount() == 0 {
				deck.Summary = strings.ReplaceAll(deck.GetSummary(), bought.GetName(), last)
				list = append(list[:i], list[i+1:]...)
				i--
			}
		}
		return list
	}
	deck.Cards = swapList(deck.GetCards(), deck.GetSideboard())
	deck.Sideboard = swapList(deck.GetSideboard(), deck.GetCards())
	if len(lines) > 0 {
		markOwned(append(append([]*mtgv1.DeckCard(nil), deck.GetCards()...), deck.GetSideboard()...))
	}
	return lines
}

// indexOf is the place of a card in a list, or -1.
func indexOf(list []*mtgv1.DeckCard, id string) int {
	for i, dc := range list {
		if dc.GetOracleId() == id {
			return i
		}
	}
	return -1
}

// listCopies is how many copies of a card a list holds.
func listCopies(list []*mtgv1.DeckCard, id string) int32 {
	var n int32
	for _, dc := range list {
		if dc.GetOracleId() == id {
			n += dc.GetCount()
		}
	}
	return n
}

// copiesName is "Fyndhorn Elves" for one copy and "2 Balduvian Bears"
// for more.
func copiesName(n int32, name string) string {
	if n == 1 {
		return name
	}
	return fmt.Sprintf("%d %s", n, name)
}
