package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
)

// rungSize is the count of lists a rung takes: five for the dev split and
// twenty for the test split (D-867).
const rungSize = 25

// devEvery puts every fifth list of a rung in the dev split, so the dev
// lists spread over the dates of the rung (D-867).
const devEvery = 5

// The tournament rung takes 13 MTGO lists and 12 MTGTop8 lists, and at
// most two lists of one event, so the rung spans more events and more
// decks (D-871). An RCQ counts from 32 players: test run 1 read the top 8
// of RCQs of 16 to 24 players at the label 14 of 30 times (D-876).
const (
	mtgoLists     = 13
	top8Lists     = 12
	eventLists    = 2
	minRCQPlayers = 32
)

// casualTypes are the product types of the precon table that label the
// casual rung (D-863).
var casualTypes = map[string]bool{"Theme Deck": true, "Intro Pack": true, "Planeswalker Deck": true}

// labelsJSON holds the steps of 60 MTGGoldfish user decks from a model
// outside the family of the judge. No product type marks the FNM line,
// so the FNM rung reads the decks labeled fnm (D-874, D-875).
//
//go:embed goldfish_labels.json
var labelsJSON []byte

// fnmKeys answers the key of each user deck labeled fnm.
func fnmKeys() (map[string]bool, error) {
	var doc struct {
		Labels []struct {
			ID   string `json:"id"`
			Step string `json:"step"`
		} `json:"labels"`
	}
	if err := json.Unmarshal(labelsJSON, &doc); err != nil {
		return nil, fmt.Errorf("sixty-gate: labels: %w", err)
	}
	out := map[string]bool{}
	for _, l := range doc.Labels {
		if l.Step == "fnm" {
			out[meta.SourceGoldfish+" "+l.ID] = true
		}
	}
	return out, nil
}

// rcq matches the event name of a Regional Championship Qualifier on
// MTGTop8. StarCityGames names its own qualifier "ReCQ".
var rcq = regexp.MustCompile(`(?i)\bRe?CQ\b`)

// calList is one labeled list of the calibration set.
type calList struct {
	label  mtgv1.SixtyStep
	id     string
	name   string
	date   string
	format string
	// event names one tournament, for the cap of D-871. A precon holds
	// none.
	event string
	dev   bool
	deck  *mtgv1.Deck
}

// picked is the calibration set and the count of lists that did not
// resolve whole, for each rung.
type picked struct {
	lists   []calList
	skipped map[mtgv1.SixtyStep]int
}

// pick builds the three rungs of lists that resolve whole against the
// card index, newest first by date, then by id (D-863, D-867, D-869). The
// FNM rung reads the user decks whose key fnm names (D-875). The
// tournament rung takes each source apart, with a cap for each event
// (D-871).
//
// A key in exclude leaves a tournament list out, so a new run reads lists
// that no earlier run read (D-876).
func pick(idx *cards.Index, precons []meta.Precon, modern, standard []meta.List, fnm, exclude map[string]bool) (picked, error) {
	out := picked{skipped: map[mtgv1.SixtyStep]int{}}
	var casual, users, mtgo, top8 []calList
	for _, p := range precons {
		if preconCount(p.Cards) < 60 {
			continue
		}
		if !casualTypes[p.Type] {
			continue
		}
		// Each casual product was sold for Standard (D-869).
		casual = append(casual, calList{
			label: mtgv1.SixtyStep_SIXTY_STEP_CASUAL, id: p.Code + " " + p.Name,
			name: p.Name, date: p.ReleaseDate, format: "Standard",
		})
	}
	byID := map[string]meta.Precon{}
	for _, p := range precons {
		byID[p.Code+" "+p.Name] = p
	}
	resolvePrecon := func(l *calList) bool {
		p := byID[l.id]
		deck, ok := resolve(idx, preconCards(p.Cards), preconCards(p.Sideboard))
		l.deck = deck
		return ok
	}
	for _, set := range []struct {
		format string
		lists  []meta.List
	}{{"Modern", modern}, {"Standard", standard}} {
		for _, m := range set.lists {
			if key := m.Source + " " + m.ID; fnm[key] {
				users = append(users, calList{
					label: mtgv1.SixtyStep_SIXTY_STEP_FNM, id: key, name: "user deck", date: m.Date, format: set.format,
				})
				continue
			}
			if !topEight(m) || exclude[m.Source+" "+m.ID] {
				continue
			}
			l := calList{
				label: mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT, id: m.Source + " " + m.ID,
				name: m.Event, date: m.Date, format: set.format, event: m.Source + " " + m.Event + " " + m.Date,
			}
			if m.Source == meta.SourceMTGO {
				mtgo = append(mtgo, l)
			} else {
				top8 = append(top8, l)
			}
		}
	}
	lists := map[string]meta.List{}
	for _, set := range [][]meta.List{modern, standard} {
		for _, m := range set {
			lists[m.Source+" "+m.ID] = m
		}
	}
	resolveList := func(l *calList) bool {
		m := lists[l.id]
		deck, ok := resolve(idx, m.Cards, m.Sideboard)
		l.deck = deck
		return ok
	}
	for _, rung := range [][]part{
		{{casual, resolvePrecon, rungSize, 0}},
		{{users, resolveList, rungSize, 0}},
		{{mtgo, resolveList, mtgoLists, eventLists}, {top8, resolveList, top8Lists, eventLists}},
	} {
		var taken []calList
		for _, pt := range rung {
			got, err := take(pt, out.skipped)
			if err != nil {
				return picked{}, err
			}
			taken = append(taken, got...)
		}
		newestFirst(taken)
		for i := range taken {
			taken[i].dev = i%devEvery == 0
		}
		out.lists = append(out.lists, taken...)
	}
	return out, nil
}

// part is one source of a rung: its lists, the count it gives, and the
// most lists of one event, or 0 for no cap.
type part struct {
	lists    []calList
	resolve  func(*calList) bool
	want     int
	perEvent int
}

// topEight is a top 8 finish in an MTGO Challenge, or in an MTGTop8 RCQ
// of 32 players or more.
func topEight(m meta.List) bool {
	if m.Placement < 1 || m.Placement > 8 {
		return false
	}
	switch m.Source {
	case meta.SourceMTGO:
		return strings.Contains(m.Event, "Challenge")
	case meta.SourceMTGTop8:
		return rcq.MatchString(m.Event) && m.Players >= minRCQPlayers
	}
	return false
}

// take sorts the lists of a part newest first, then by id, and takes the
// first lists that resolve whole, under the cap of each event.
func take(pt part, skipped map[mtgv1.SixtyStep]int) ([]calList, error) {
	newestFirst(pt.lists)
	var out []calList
	events := map[string]int{}
	for i := range pt.lists {
		if len(out) == pt.want {
			break
		}
		l := pt.lists[i]
		if pt.perEvent > 0 && events[l.event] >= pt.perEvent {
			continue
		}
		if !pt.resolve(&l) {
			skipped[l.label]++
			continue
		}
		events[l.event]++
		out = append(out, l)
	}
	if len(out) < pt.want {
		label := "an empty"
		if len(pt.lists) > 0 {
			label = "the " + stepWord(pt.lists[0].label)
		}
		return nil, fmt.Errorf("sixty-gate: %s rung holds %d lists that resolve, want %d", label, len(out), pt.want)
	}
	return out, nil
}

func newestFirst(lists []calList) {
	sort.Slice(lists, func(i, j int) bool {
		if lists[i].date != lists[j].date {
			return lists[i].date > lists[j].date
		}
		return lists[i].id < lists[j].id
	})
}

func preconCount(cs []meta.PreconCard) int {
	n := 0
	for _, c := range cs {
		n += c.Count
	}
	return n
}

func preconCards(cs []meta.PreconCard) []meta.Card {
	out := make([]meta.Card, 0, len(cs))
	for _, c := range cs {
		out = append(out, meta.Card{Name: c.Name, Count: c.Count, OracleID: c.OracleID})
	}
	return out
}

// resolve turns a list into a deck. A card the index does not know makes
// the list unusable, because the judge would read a list that is not the
// labeled one.
func resolve(idx *cards.Index, main, side []meta.Card) (*mtgv1.Deck, bool) {
	deck := &mtgv1.Deck{}
	for _, part := range []struct {
		in  []meta.Card
		out *[]*mtgv1.DeckCard
	}{{main, &deck.Cards}, {side, &deck.Sideboard}} {
		for _, lc := range part.in {
			c, ok := listCard(idx, lc)
			if !ok {
				return nil, false
			}
			*part.out = append(*part.out, &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: int32(lc.Count)})
		}
	}
	return deck, len(deck.GetCards()) > 0
}

func listCard(idx *cards.Index, lc meta.Card) (*mtgv1.Card, bool) {
	if lc.OracleID != "" {
		if c, ok := idx.ByOracleID(lc.OracleID); ok {
			return c, true
		}
	}
	return idx.ByName(lc.Name)
}

func stepWord(s mtgv1.SixtyStep) string {
	switch s {
	case mtgv1.SixtyStep_SIXTY_STEP_CASUAL:
		return "casual"
	case mtgv1.SixtyStep_SIXTY_STEP_FNM:
		return "fnm"
	case mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT:
		return "tournament"
	}
	return "none"
}
