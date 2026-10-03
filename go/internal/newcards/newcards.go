// Package newcards finds the cards of a new set that fit each saved deck,
// the third event of PR-26 (D-1090, D-1091, D-1095). The snapshot job runs the
// pass after a snapshot makes cards legal for the first time. Each deck
// takes at most PerDeck cards, and the pass replaces the cards of the
// last pass.
package newcards

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// PerDeck is the most new cards that one deck takes (D-1091).
const PerDeck = 3

// Store reads and marks the stored decks. The deck repo holds both, as
// the stale pass reads them.
type Store interface {
	// Scan calls fn for every stored deck of every user.
	Scan(ctx context.Context, fn func(uid string, d *mtgv1.Deck) error) error
	// Mark changes one stored deck through fn, inside a transaction. It
	// writes only when fn reports a change.
	Mark(ctx context.Context, uid, id string, fn func(d *mtgv1.Deck) bool) (bool, error)
}

// Fitter scores cards against a theme. The candidates builder is one.
type Fitter interface {
	// ThemeScores returns the score of each card for a theme in words.
	ThemeScores(theme string, idx *cards.Index, cs []*mtgv1.Card) []float64
	// DeckTheme names the theme that most cards of a deck fit, or "".
	DeckTheme(idx *cards.Index, deck []*mtgv1.Card) string
}

// Input is what one pass reads.
type Input struct {
	// Index is the newest snapshot.
	Index *cards.Index
	// Version names the snapshot version of the new-cards marker. A deck
	// that a pass of this version wrote is not written again (D-1095).
	Version string
	// New are the oracle ids of the new cards (cards.NewlyLegal).
	New []string
	// KeyOf gives the Scryfall key of a format, or "" for a format with
	// no legality check (D-3).
	KeyOf func(mtgv1.FormatId) string
	// ThemeOf reads the theme of the chat of a deck. It returns "" when
	// the chat is gone or holds no theme.
	ThemeOf func(ctx context.Context, uid, sessionID string) (string, error)
	// Fit scores the new cards.
	Fit Fitter
	// Floor is the least score of a card that fits.
	Floor float64
}

// Result counts one pass. Hit holds, for each user, the decks that took
// new cards in this pass. The push reads it (D-1088).
type Result struct {
	Read    int
	Fit     int
	Written int
	Hit     map[string][]*mtgv1.Deck
}

// Pass reads every stored deck and writes the new cards that fit it. A
// deck that takes no card loses the cards of the last pass (D-1095). A
// deck that a pass of the same version wrote is not written and not hit,
// so a pass that runs again after a failure sends no second push, and a
// dismiss stays. A deck whose cards do not change takes the version and
// no push.
func Pass(ctx context.Context, store Store, in Input, log *slog.Logger) (Result, error) {
	var res Result
	if in.Version == "" {
		return res, errors.New("newcards: a pass needs the version of its marker")
	}
	fresh := make([]*mtgv1.Card, 0, len(in.New))
	for _, id := range in.New {
		if c, ok := in.Index.ByOracleID(id); ok {
			fresh = append(fresh, c)
		}
	}
	scores := map[string][]float64{}
	scoresFor := func(theme string) []float64 {
		s, ok := scores[theme]
		if !ok {
			s = in.Fit.ThemeScores(theme, in.Index, fresh)
			scores[theme] = s
		}
		return s
	}
	themes := map[string]string{}
	err := store.Scan(ctx, func(uid string, d *mtgv1.Deck) error {
		res.Read++
		theme, err := deckTheme(ctx, in, themes, uid, d)
		if err != nil {
			return fmt.Errorf("deck %s/%s: %w", uid, d.GetId(), err)
		}
		var picks []string
		if theme != "" {
			picks = Pick(d, DeckColors(in.Index, d), fresh, scoresFor(theme), in.KeyOf(d.GetFormat().GetId()), in.Floor)
		}
		if len(picks) > 0 {
			res.Fit++
		}
		if !needsWrite(d, in.Version, picks) {
			return nil
		}
		var hit *mtgv1.Deck
		changed, err := store.Mark(ctx, uid, d.GetId(), func(cur *mtgv1.Deck) bool {
			hit = nil
			if !needsWrite(cur, in.Version, picks) {
				return false
			}
			if len(picks) > 0 && !slices.Equal(picks, cur.GetNewOracleIds()) {
				hit = cur
			}
			cur.NewOracleIds = picks
			cur.NewCardsVersion = in.Version
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
			log.InfoContext(ctx, "new cards pass: deck marked", "deck", d.GetId(), "cards", len(picks), "theme", theme)
		}
		return nil
	})
	return res, err
}

// needsWrite reports whether a pass of version writes picks to a deck. A
// pass of the same version wrote the deck before, and a deck with no
// cards that takes none needs no write.
func needsWrite(d *mtgv1.Deck, version string, picks []string) bool {
	if d.GetNewCardsVersion() == version {
		return false
	}
	return len(picks) > 0 || len(d.GetNewOracleIds()) > 0
}

// deckTheme reads the theme of the chat of the deck, and the theme of its
// cards when the chat is gone or holds none (D-1091). It keeps each read
// for the pass, because a chat can hold several decks.
func deckTheme(ctx context.Context, in Input, themes map[string]string, uid string, d *mtgv1.Deck) (string, error) {
	if sid := d.GetSessionId(); sid != "" && in.ThemeOf != nil {
		key := uid + "/" + sid
		theme, ok := themes[key]
		if !ok {
			var err error
			if theme, err = in.ThemeOf(ctx, uid, sid); err != nil {
				return "", err
			}
			themes[key] = theme
		}
		if theme != "" {
			return theme, nil
		}
	}
	return in.Fit.DeckTheme(in.Index, deckCards(in.Index, d)), nil
}

// deckCards reads the card data of each card of the deck that the
// snapshot holds.
func deckCards(idx *cards.Index, d *mtgv1.Deck) []*mtgv1.Card {
	var out []*mtgv1.Card
	for _, id := range deckIDs(d) {
		if c, ok := idx.ByOracleID(id); ok {
			out = append(out, c)
		}
	}
	return out
}

// deckIDs lists each oracle id of the deck once: the commanders, the main
// list, the sideboard, and the companion.
func deckIDs(d *mtgv1.Deck) []string {
	seen := map[string]bool{}
	var out []string
	add := func(id string) {
		if id != "" && !seen[id] {
			seen[id] = true
			out = append(out, id)
		}
	}
	for _, id := range d.GetCommanderOracleIds() {
		add(id)
	}
	for _, list := range [][]*mtgv1.DeckCard{d.GetCards(), d.GetSideboard()} {
		for _, dc := range list {
			add(dc.GetOracleId())
		}
	}
	add(d.GetCompanionOracleId())
	return out
}

// Pick returns the oracle ids of the new cards that fit one deck, best
// first, at most PerDeck. A card fits when it reaches the floor, it is
// legal or restricted under the format key, the deck does not hold it,
// and its color identity sits inside the colors of the deck. A tie keeps
// the name order. A format with no key takes no card, because the pass
// can not read its legality.
func Pick(d *mtgv1.Deck, colors map[mtgv1.Color]bool, fresh []*mtgv1.Card, scores []float64, key string, floor float64) []string {
	if key == "" {
		return nil
	}
	held := map[string]bool{}
	for _, id := range deckIDs(d) {
		held[id] = true
	}
	type scored struct {
		card  *mtgv1.Card
		score float64
	}
	var fits []scored
	for i, c := range fresh {
		if scores[i] < floor || held[c.GetOracleId()] || !legal(c, key) || !inside(c.GetColorIdentity(), colors) {
			continue
		}
		fits = append(fits, scored{c, scores[i]})
	}
	slices.SortFunc(fits, func(a, b scored) int {
		if c := cmp.Compare(b.score, a.score); c != 0 {
			return c
		}
		return cmp.Compare(a.card.GetName(), b.card.GetName())
	})
	var out []string
	for _, f := range fits[:min(len(fits), PerDeck)] {
		out = append(out, f.card.GetOracleId())
	}
	return out
}

func legal(c *mtgv1.Card, key string) bool {
	switch c.GetLegalities()[key] {
	case mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL, mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED:
		return true
	}
	return false
}

// DeckColors is the color identity of the deck: that of its commanders
// when it has any, and else that of each card it holds. A deck stores no
// card data, so the snapshot gives the colors.
func DeckColors(idx *cards.Index, d *mtgv1.Deck) map[mtgv1.Color]bool {
	ids := d.GetCommanderOracleIds()
	if len(ids) == 0 {
		ids = deckIDs(d)
	}
	out := map[mtgv1.Color]bool{}
	for _, id := range ids {
		if c, ok := idx.ByOracleID(id); ok {
			for _, col := range c.GetColorIdentity() {
				out[col] = true
			}
		}
	}
	return out
}

func inside(identity []mtgv1.Color, colors map[mtgv1.Color]bool) bool {
	for _, c := range identity {
		if !colors[c] {
			return false
		}
	}
	return true
}
