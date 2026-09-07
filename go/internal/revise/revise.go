// Package revise reads a message after a build and turns it into a
// revision brief (D-283). A request has three outcomes
// and never silence: a question when it is unclear, a change, or a
// decline with a reason (D-284).
package revise

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
)

// Brief is what the revise call returns.
type Brief struct {
	// Changes are the requests the agent will act on, one short line
	// each, for the generator and for the reply.
	Changes []string `json:"changes"`
	// Remove names cards the user wants out. Keep names cards the user
	// wants in. Both copy names from the deck list, exactly.
	Remove []string `json:"remove"`
	Keep   []string `json:"keep"`
	// MaxManaValue caps every nonland card. Zero means no cap.
	MaxManaValue float64 `json:"max_mana_value"`
	// SwapBasics is how many basic lands the user wants replaced with
	// nonbasic lands. Zero means no land swap. LandKinds says what kind,
	// in plain words for the generator (D-448).
	SwapBasics int    `json:"swap_basics"`
	LandKinds  string `json:"land_kinds"`
	// Question is one clarifying question, or empty. A question means
	// no build this turn.
	Question string `json:"question"`
	// Declined lists the requests the agent will not act on, each with
	// a reason the user can read.
	Declined []Decline `json:"declined"`
}

// Decline is one request the agent turned down, with its reason.
type Decline struct {
	Request string `json:"request"`
	Reason  string `json:"reason"`
}

// Acts says the brief changes the deck. A brief with only a question or
// only declines runs no build.
func (b *Brief) Acts() bool {
	return len(b.Changes) > 0 || len(b.Remove) > 0 || len(b.Keep) > 0 || b.MaxManaValue > 0 || b.SwapBasics > 0
}

// JSON is the brief as the turn stores it, so a wrong revision is
// readable after the fact (D-449). A brief that will not marshal is
// impossible, and an empty string marks it.
func (b *Brief) JSON() string {
	raw, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(raw)
}

// Input is what the call reads.
type Input struct {
	SessionID string
	Message   string
	// Prior is the user's message of the turn before, when the agent
	// asked a question and this message answers it.
	Prior string
	Deck  *mtgv1.Deck
	// Format and Power are words for the prompt, for example "Modern"
	// and "casual".
	Format string
	Power  string
	// Cards answers the mana value and the type of each deck card.
	Cards CardSource
}

// CardSource answers one Oracle id.
type CardSource interface {
	ByOracleID(id string) (*mtgv1.Card, bool)
}

// Call runs the revise role once.
func Call(ctx context.Context, client *llm.Client, in Input, acc *llm.Accumulator) (*Brief, error) {
	if in.Deck == nil {
		return nil, fmt.Errorf("revise: no deck to revise")
	}
	res, err := client.Complete(ctx, llm.RoleRevise, llm.Request{
		Instructions: instructions,
		Input:        input(in),
		SchemaName:   "revision_brief",
		Schema:       json.RawMessage(schema),
		CacheKey:     in.SessionID,
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("revise: %w", err)
	}
	var b Brief
	if err := json.Unmarshal(res.Output, &b); err != nil {
		return nil, fmt.Errorf("revise: output: %w", err)
	}
	b.Question = strings.TrimSpace(b.Question)
	b.LandKinds = strings.TrimSpace(b.LandKinds)
	if b.SwapBasics < 0 {
		b.SwapBasics = 0
	}
	b.Remove = onlyInDeck(b.Remove, in.Deck)
	b.Keep = onlyInDeck(b.Keep, in.Deck)
	return &b, nil
}

// onlyInDeck drops a name the deck does not hold. The model copies names
// from the list, and a name it invented has nothing to remove or keep.
func onlyInDeck(names []string, deck *mtgv1.Deck) []string {
	have := map[string]string{}
	for _, c := range deck.GetCards() {
		have[fold(c.GetName())] = c.GetName()
	}
	var out []string
	seen := map[string]bool{}
	for _, n := range names {
		name, ok := have[fold(n)]
		if !ok || seen[name] {
			continue
		}
		seen[name] = true
		out = append(out, name)
	}
	return out
}

func input(in Input) string {
	var s strings.Builder
	fmt.Fprintf(&s, "## The deck\n\nFormat: %s.", in.Format)
	if in.Power != "" {
		fmt.Fprintf(&s, " Power: %s.", in.Power)
	}
	s.WriteString("\n\n")
	for _, c := range in.Deck.GetCards() {
		mv, kind := "", ""
		if in.Cards != nil {
			if card, ok := in.Cards.ByOracleID(c.GetOracleId()); ok {
				kind = card.GetTypeLine()
				if !isLand(card) {
					mv = fmt.Sprintf(", mana value %g", card.GetManaValue())
				}
			}
		}
		fmt.Fprintf(&s, "- %d %s (%s%s)\n", c.GetCount(), c.GetName(), kind, mv)
	}
	if sum := strings.TrimSpace(in.Deck.GetSummary()); sum != "" {
		fmt.Fprintf(&s, "\nThe deck's summary: %s\n", sum)
	}
	if p := strings.TrimSpace(in.Prior); p != "" {
		fmt.Fprintf(&s, "\n## The user's earlier message\n\nYou asked a question about this message, and the message below answers it. Act on both.\n\n%s\n", p)
	}
	fmt.Fprintf(&s, "\n## The user's message\n\n%s\n", strings.TrimSpace(in.Message))
	return s.String()
}

func isLand(c *mtgv1.Card) bool {
	for _, t := range c.GetCardTypes() {
		if t == "Land" {
			return true
		}
	}
	return false
}

func fold(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

// Diff is what changed between the base deck and the revised one.
type Diff struct {
	Added   []string
	Removed []string
	// Changed names a card whose count moved, as "name: 2 to 4".
	Changed []string
}

// DiffDecks compares two card lists by name.
func DiffDecks(base, revised *mtgv1.Deck) Diff {
	before := counts(base)
	after := counts(revised)
	var d Diff
	for name, n := range after {
		m, ok := before[name]
		switch {
		case !ok:
			d.Added = append(d.Added, fmt.Sprintf("%d %s", n, name))
		case m != n:
			d.Changed = append(d.Changed, fmt.Sprintf("%s: %d to %d", name, m, n))
		}
	}
	for name, m := range before {
		if _, ok := after[name]; !ok {
			d.Removed = append(d.Removed, fmt.Sprintf("%d %s", m, name))
		}
	}
	sort.Strings(d.Added)
	sort.Strings(d.Removed)
	sort.Strings(d.Changed)
	return d
}

func counts(d *mtgv1.Deck) map[string]int32 {
	out := map[string]int32{}
	for _, c := range d.GetCards() {
		out[c.GetName()] += c.GetCount()
	}
	return out
}

// Note writes the reply the user reads. It comes from the diff and the
// brief, not from the model, so it can not claim a change that did not
// happen (D-283). A decline is named with its reason (D-284).
func Note(b *Brief, d Diff) string {
	var lines []string
	if len(d.Removed) > 0 {
		lines = append(lines, "I removed "+list(d.Removed)+".")
	}
	if len(d.Added) > 0 {
		lines = append(lines, "I added "+list(d.Added)+".")
	}
	if len(d.Changed) > 0 {
		lines = append(lines, "I changed the count of "+list(d.Changed)+".")
	}
	if b != nil && b.Acts() && len(lines) == 0 {
		lines = append(lines, "The deck is the same as before. The rebuild found nothing to change for that request.")
	}
	for _, dec := range declines(b) {
		lines = append(lines, fmt.Sprintf("I did not %s: %s", trimDot(lower1(dec.Request)), trimDot(dec.Reason)+"."))
	}
	return strings.Join(lines, " ")
}

// DeclineNote is the reply when the brief acts on nothing and asks
// nothing. Silence is never an outcome.
func DeclineNote(b *Brief) string {
	decs := declines(b)
	if len(decs) == 0 {
		return "I read your message, and I found no change to make to the deck. Tell me what you want changed, and I will do it."
	}
	var lines []string
	for _, dec := range decs {
		lines = append(lines, fmt.Sprintf("I did not %s: %s", trimDot(lower1(dec.Request)), trimDot(dec.Reason)+"."))
	}
	return strings.Join(lines, " ") + " The deck stays as it was."
}

func declines(b *Brief) []Decline {
	if b == nil {
		return nil
	}
	var out []Decline
	for _, d := range b.Declined {
		if strings.TrimSpace(d.Request) == "" {
			continue
		}
		if strings.TrimSpace(d.Reason) == "" {
			d.Reason = "no change would help here"
		}
		out = append(out, d)
	}
	return out
}

func list(items []string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return items[0]
	case 2:
		return items[0] + " and " + items[1]
	default:
		return strings.Join(items[:len(items)-1], ", ") + ", and " + items[len(items)-1]
	}
}

func trimDot(s string) string { return strings.TrimRight(strings.TrimSpace(s), ".") }

func lower1(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
