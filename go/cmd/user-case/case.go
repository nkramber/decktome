package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The repository is public, so a case holds the user's words, the slots,
// the counts, and the bars, and never the collection (D-639, D-642). The
// collection sits in a private object, pinned by its hash (D-1124).
//
//go:embed cases/*.json
var embedded embed.FS

// Case is one replay of a real session.
type Case struct {
	Name       string     `json:"name"`
	Decision   string     `json:"decision"`
	About      string     `json:"about"`
	Collection Collection `json:"collection"`
	Snapshot   Pinned     `json:"snapshot"`
	Model      Pinned     `json:"model"`
	Slots      Slots      `json:"slots"`
	Turns      []Turn     `json:"turns"`
	Paid       Paid       `json:"paid"`
	Steps      []Step     `json:"steps"`
}

// Collection names the private object and its hash.
type Collection struct {
	Object  string `json:"object"`
	SHA256  string `json:"sha256"`
	Format  string `json:"format"`
	Entries int    `json:"entries"`
	Cards   int    `json:"cards"`
}

// Pinned is one stored version and its object prefix.
type Pinned struct {
	Version string `json:"version"`
	Object  string `json:"object"`
}

// Slots are the filled slots of the session.
type Slots struct {
	Format    string   `json:"format"`
	Commander string   `json:"commander"`
	Colors    []string `json:"colors"`
	Bracket   int32    `json:"bracket"`
	Pool      string   `json:"pool"`
}

// Turn is one user turn: the words, or the answers to questions.
type Turn struct {
	Turn    int      `json:"turn"`
	Words   string   `json:"words,omitempty"`
	Answers []Answer `json:"answers,omitempty"`
}

// Answer is one answer to a question, by option or by text.
type Answer struct {
	Question    string `json:"question"`
	OptionIndex *int   `json:"option_index,omitempty"`
	Means       string `json:"means,omitempty"`
	Text        string `json:"text,omitempty"`
}

// Paid is the chat-probe lane, which costs money.
type Paid struct {
	Target   string   `json:"target"`
	Command  string   `json:"command"`
	Messages []string `json:"messages"`
	Limits   string   `json:"limits"`
}

// Step is one shortlist build of the session.
type Step struct {
	Name      string    `json:"name"`
	AfterTurn int       `json:"after_turn"`
	Theme     string    `json:"theme"`
	Avoid     string    `json:"avoid"`
	Measures  []Measure `json:"measures"`
	Deck      []Measure `json:"deck"`
}

// Measure is one count with the number of the real session and a bar.
type Measure struct {
	ID       string  `json:"id"`
	Baseline float64 `json:"baseline"`
	Note     string  `json:"note,omitempty"`
	Bar      Bar     `json:"bar"`
}

// Bar is the pass range. Todo marks a bar that is not set yet, and a
// measure with no bound reads TODO, not PASS.
type Bar struct {
	Min  *float64 `json:"min,omitempty"`
	Max  *float64 `json:"max,omitempty"`
	Todo string   `json:"todo,omitempty"`
}

// Verdict words.
const (
	Pass = "PASS"
	Fail = "FAIL"
	Todo = "TODO"
)

// Judge reads one value against the bar.
func (b Bar) Judge(v float64) string {
	if b.Min == nil && b.Max == nil {
		return Todo
	}
	if b.Min != nil && v < *b.Min {
		return Fail
	}
	if b.Max != nil && v > *b.Max {
		return Fail
	}
	return Pass
}

// String writes the bar for the table.
func (b Bar) String() string {
	switch {
	case b.Min == nil && b.Max == nil:
		return Todo
	case b.Min != nil && b.Max != nil:
		return fmt.Sprintf("%s..%s", num(*b.Min), num(*b.Max))
	case b.Min != nil:
		return ">=" + num(*b.Min)
	default:
		return "<=" + num(*b.Max)
	}
}

// LoadCase reads a case by its name from the embedded set, or from a
// path when the argument names a file.
func LoadCase(arg string) (*Case, error) {
	var raw []byte
	var err error
	if strings.HasSuffix(arg, ".json") {
		raw, err = os.ReadFile(arg) // #nosec G304 -- the operator names the file.
	} else {
		raw, err = embedded.ReadFile("cases/" + arg + ".json")
	}
	if err != nil {
		return nil, fmt.Errorf("case %s: %w", arg, err)
	}
	var c Case
	dec := json.NewDecoder(strings.NewReader(string(raw)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&c); err != nil {
		return nil, fmt.Errorf("case %s: %w", arg, err)
	}
	return &c, c.check()
}

func (c *Case) check() error {
	if c.Name == "" || len(c.Collection.SHA256) != 64 || c.Snapshot.Version == "" || len(c.Steps) == 0 {
		return fmt.Errorf("case %q: a name, a 64-digit sha256, a snapshot version, and a step are required", c.Name)
	}
	for _, s := range c.Steps {
		for _, m := range s.Measures {
			if _, ok := measures[m.ID]; !ok && m.ID != unmatchedWords && m.ID != shortlist {
				return fmt.Errorf("case %q step %s: no measure named %q", c.Name, s.Name, m.ID)
			}
		}
		for _, m := range s.Deck {
			if _, ok := measures[m.ID]; !ok && !slices.Contains(deckMeasures, m.ID) {
				return fmt.Errorf("case %q step %s: no deck measure named %q", c.Name, s.Name, m.ID)
			}
		}
	}
	return nil
}

// Measure ids that read the list, not one card.
const (
	shortlist      = "shortlist"
	unmatchedWords = "unmatched_words"
)

// deckMeasures are the ids that read a whole deck, not one card.
var deckMeasures = []string{"lands", "unique_names", "quality_grade", "kept_from_build_1"}

// Tagged answers whether the card holds a tag of the snapshot.
type Tagged func(tag, oracleID string) bool

var (
	// ownTarget is a spell that targets the user's own creature, the way
	// a voltron deck protects or pumps its commander.
	ownTarget = regexp.MustCompile(`target (creature|permanent|artifact or creature|creature or [a-z]+|nonland permanent)( or [a-z]+)? you (control|own)`)
	handSize  = []string{"maximum hand size", "cards in your hand", "card in your hand"}
)

// measures count one card each. The output prints counts alone, never a
// card name, because a name can come from the private collection (D-1124).
var measures = map[string]func(c *mtgv1.Card, tagged Tagged) bool{
	"equipment": func(c *mtgv1.Card, _ Tagged) bool { return slices.Contains(c.GetSubtypes(), "Equipment") },
	"counterspells": func(c *mtgv1.Card, tagged Tagged) bool {
		return tagged("counterspell", c.GetOracleId())
	},
	"cantrips": func(c *mtgv1.Card, tagged Tagged) bool { return tagged("cantrip", c.GetOracleId()) },
	"hand_size": func(c *mtgv1.Card, _ Tagged) bool {
		tx := strings.ToLower(c.GetOracleText())
		return slices.ContainsFunc(handSize, func(p string) bool { return strings.Contains(tx, p) })
	},
	"targeting_spells": func(c *mtgv1.Card, _ Tagged) bool {
		spell := slices.Contains(c.GetCardTypes(), "Instant") || slices.Contains(c.GetCardTypes(), "Sorcery")
		return spell && ownTarget.MatchString(strings.ToLower(c.GetOracleText()))
	},
	"enchant_creature_auras": func(c *mtgv1.Card, _ Tagged) bool {
		return slices.Contains(c.GetSubtypes(), "Aura") && strings.Contains(strings.ToLower(c.GetOracleText()), "enchant creature")
	},
	// The type line covers both faces, so an artifact that transforms
	// into an artifact creature counts.
	"artifact_creatures": func(c *mtgv1.Card, _ Tagged) bool {
		return strings.Contains(c.GetTypeLine(), "Artifact") && strings.Contains(c.GetTypeLine(), "Creature")
	},
}

// Count measures one shortlist: its size, the theme words that matched
// nothing, and each card measure.
func Count(list []*mtgv1.Card, unmatched []string, tagged Tagged) map[string]float64 {
	out := map[string]float64{shortlist: float64(len(list)), unmatchedWords: float64(len(unmatched))}
	for id, pred := range measures {
		n := 0
		for _, c := range list {
			if pred(c, tagged) {
				n++
			}
		}
		out[id] = float64(n)
	}
	return out
}

// DeckCount measures one deck that reached the user: each card measure,
// the lands, the distinct names, the grade, and the names it kept from
// the deck of the step before (D-1124). It answers counts alone.
func DeckCount(d, prev *mtgv1.Deck, card func(oracleID string) (*mtgv1.Card, bool), tagged Tagged) map[string]float64 {
	ids := map[string]bool{}
	var picked []*mtgv1.Card
	lands := 0.0
	for _, dc := range d.GetCards() {
		ids[dc.GetOracleId()] = true
		c, ok := card(dc.GetOracleId())
		if !ok {
			continue
		}
		picked = append(picked, c)
		if slices.Contains(c.GetCardTypes(), "Land") {
			lands += float64(dc.GetCount())
		}
	}
	out := Count(picked, nil, tagged)
	delete(out, shortlist)
	delete(out, unmatchedWords)
	out["lands"] = lands
	out["unique_names"] = float64(len(ids))
	out["quality_grade"] = float64(d.GetQuality().GetScore())
	if prev != nil {
		kept := 0
		for _, dc := range prev.GetCards() {
			if ids[dc.GetOracleId()] {
				kept++
			}
		}
		out["kept_from_build_1"] = float64(kept)
	}
	return out
}

// Row is one line of the table.
type Row struct {
	Step, ID string
	Baseline float64
	Now      float64
	Bar      string
	Verdict  string
}

// Judge reads each measure of a step against its bar.
func Judge(step Step, now map[string]float64) []Row {
	rows := make([]Row, 0, len(step.Measures))
	for _, m := range step.Measures {
		v := now[m.ID]
		rows = append(rows, Row{Step: step.Name, ID: m.ID, Baseline: m.Baseline, Now: v, Bar: m.Bar.String(), Verdict: m.Bar.Judge(v)})
	}
	return rows
}

func num(v float64) string {
	if v == float64(int64(v)) {
		return fmt.Sprintf("%d", int64(v))
	}
	return fmt.Sprintf("%.3f", v)
}
