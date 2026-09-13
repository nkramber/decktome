// Package questions holds the clarifying-question workflow (roadmap PR-7).
// The catalog is the first source of a question (D-25). This package is
// the deterministic half: which slot is empty, which row fits it, and in
// which order. The model phrases the question and maps the answer.
//
// The rows mirror the mtg-corpus skill, section 11. TestCatalogMatchesCorpus
// fails when the two drift apart.
package questions

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

//go:embed catalog.json
var catalogJSON []byte

// Row is one catalog question.
type Row struct {
	// ID is stable. The no-repeat rule and the M-4 log use it.
	ID string `json:"id"`
	// Slot is the proto slot the answer informs. Several rows inform one
	// slot: the commander rows all inform "commander".
	Slot string `json:"slot"`
	// Key is the question's own state key. It defaults to Slot. A
	// refinement question carries its own key, so a filled proto slot does
	// not cancel it. The power confirm row informs power, but the inferred
	// step must not cancel the question.
	Key string `json:"key"`
	// Order is the ask order. Format first, then theme, then commander
	// (corpus section 11, D-294).
	Order int `json:"order"`
	// Text may hold {placeholders} that the caller fills before the model
	// phrases the question.
	Text string `json:"text"`
	// Options are suggested answers. The user may answer in free text.
	Options []string `json:"options"`
	// OptionValues carry the typed value of each option, in the words the
	// classifier writes: "bracket 4", "commander", "owned_only" (D-597).
	// A reader who picks an option sends its index, and the engine sets
	// the slot from this list with no model in the path.
	//
	// Before this the index became the option text, the text went into
	// the message, and the slot filled only when the classifier answered
	// with a string the slot could read. A model that echoed the option
	// left the slot asked (F-70).
	//
	// It is empty for a row whose options carry no typed value. Load
	// refuses a list of a length the options do not match.
	OptionValues []string `json:"option_values"`
	// Fallback is the wording used when a placeholder has no value. It
	// holds no placeholder itself. Load enforces that.
	Fallback string `json:"fallback"`
	// Fixed marks a row whose exact words carry the meaning. The model
	// may neither replace it nor phrase it again.
	//
	// Three kinds of row need it. The first states what this app does or
	// does not do: the sentence that names the limit is the point of it,
	// and a replacement drops that sentence (D-112, D-117). The second
	// bundles several values into one yes-or-no question. The house-limits
	// row asks "do the normal limits hold", and a rephrasing that lists
	// the limits reads as three questions in one (D-162).
	//
	// The third explains an acronym the reader has never seen (D-374).
	// The 60-card power row spells out Friday Night Magic. Gate run 29
	// reworded that row 17 times: it kept the expansion 4 times, cut it
	// to the bare acronym 2 times, and dropped the whole option list 11
	// times. The options still say FNM in every one of those 11. A fixed
	// row can not lose the words that carry the meaning.
	Fixed bool `json:"fixed"`
	// Closed says the options are the whole answer space, so the UI
	// offers no free-text field (D-295).
	Closed bool `json:"closed"`
	// NoDecline says the UI shows no "You decide" control for the row.
	// The commander row carries it, because its own option already asks
	// for a suggestion (D-690).
	NoDecline bool `json:"no_decline"`
	// Repeat exempts a row from the no-repeat rule. Only the commander
	// pick row uses it: a user who answers "none" gets three new names
	// until one fits (D-73). The row still closes when its key closes.
	Repeat bool `json:"repeat"`
	// RepeatOnChange narrows Repeat. Such a row asks again only when its
	// content changed, which for the pick row means three other names.
	// A repeat with the same three names is the same question in the same
	// words (D-163).
	RepeatOnChange bool `json:"repeat_on_change"`
	When           When `json:"when"`
}

// When holds the triggers of one row. A nil pointer means the row does
// not care about that fact. Requires names slots that must be filled
// first, which is how D-67 holds the card-pool question back.
type When struct {
	Words    []string `json:"words"`
	Requires []string `json:"requires"`
	// NotOutstanding names the state keys this row waits for. The row
	// does not fire while one of them holds an unanswered question.
	NotOutstanding []string `json:"not_outstanding"`
	// NotBeside names the slots this row never shares a turn with. The row
	// stays out of a turn that asks a question of one of them, and it
	// waits for no answer: a later turn may ask it (D-669).
	NotBeside []string `json:"not_beside"`
	// Format is "commander", "sixty", or empty. Load refuses another word.
	Format           string `json:"format"`
	PowerCompetitive *bool  `json:"power_competitive"`
	OutOfScope       *bool  `json:"out_of_scope"`
	NamedCard        *bool  `json:"named_card"`
	Suggested        *bool  `json:"suggested"`
	CommanderSet     *bool  `json:"commander_set"`
	HasCollection    *bool  `json:"has_collection"`
	ThinTheme        *bool  `json:"thin_theme"`
	BuyList          *bool  `json:"buy_list"`
	BudgetAmbiguous  *bool  `json:"budget_ambiguous"`
	HouseFormat      *bool  `json:"house_format"`
	// TwoDecks marks a request for more than one deck (D-112).
	TwoDecks *bool `json:"two_decks"`
	// UnsupportedFormat marks a format this app does not build (D-112).
	UnsupportedFormat *bool `json:"unsupported_format"`
	// NoNearFormat marks an unsupported format with no substitute (D-146).
	NoNearFormat *bool `json:"no_near_format"`
	// Precon marks a request to upgrade a preconstructed deck (D-113).
	Precon *bool `json:"precon"`
	// CommanderIllegal marks a named commander that can not lead (D-129).
	CommanderIllegal *bool `json:"commander_illegal"`
	// CommanderUnresolved marks a commander name the card index does not
	// hold (F-75, D-606).
	CommanderUnresolved *bool `json:"commander_unresolved"`
	// CommanderNoMatch narrows that to a name no card holds at all.
	CommanderNoMatch *bool `json:"commander_no_match"`
	// SetLimited marks a deck limited to the sets the reader named
	// (D-376).
	SetLimited *bool `json:"set_limited"`
	// SetUnresolved marks a set name this app can not settle (D-376).
	SetUnresolved *bool `json:"set_unresolved"`
	// ThinSetMana marks named sets that hold too few mana cards (D-382).
	ThinSetMana *bool `json:"thin_set_mana"`
	// NamedLeader marks a card the reader named that can lead a deck,
	// while its role is still unsettled (D-388). Such a card may fix the
	// deck's color identity, so the color row waits for the role.
	NamedLeader *bool `json:"named_leader"`
	// PreconsExcluded marks a deck that uses no card of the reader's
	// precons (D-496).
	PreconsExcluded *bool `json:"precons_excluded"`
	// PreconUnresolved marks a precon name the table can not settle
	// (D-496).
	PreconUnresolved *bool `json:"precon_unresolved"`
}

// Catalog is the loaded table.
type Catalog struct {
	VerifiedAt string `json:"verified_at"`
	Rows       []Row  `json:"rows"`
}

// slots are the slot names the proto documents on Slots.slot_states.
// A row must fill one of them. A question with no slot can not close.
var slots = map[string]bool{
	// scope is not a deck value. It records that the agent said it builds
	// Magic decks only, after the user asked for something else (D-99).
	"scope": true,
	// deck_count is not a deck value either. It records that the agent
	// said it builds one deck at a time, after the user asked for two
	// (D-112).
	"deck_count": true,
	"format":     true, "power": true, "colors": true, "theme": true,
	"commander": true, "pool_rule": true, "budget": true,
	// house_rules holds the user's own words for "anything goes". The
	// build copies it to Format.house_rules (D-3, D-265). locked,
	// plan_variant, and meta left with their rows (D-260).
	"house_rules": true,
	// set holds the sets the reader limited the deck to (D-376), and
	// set_outside_mana holds whether the mana base may reach past them
	// (D-382).
	SlotSet:            true,
	SlotSetOutsideMana: true,
	// precons holds the precon products the deck uses no card of (D-496).
	SlotPrecons: true,
}

// Load reads the embedded catalog and checks it.
func Load() (*Catalog, error) { return parse(catalogJSON) }

// parse reads one catalog document and checks it.
func parse(data []byte) (*Catalog, error) {
	var c Catalog
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("questions: catalog.json: %w", err)
	}
	if c.VerifiedAt == "" {
		return nil, fmt.Errorf("questions: catalog.json has no verified_at")
	}
	seen := map[string]bool{}
	order := map[int]string{}
	for _, r := range c.Rows {
		switch {
		case r.ID == "":
			return nil, fmt.Errorf("questions: a row has no id")
		case seen[r.ID]:
			return nil, fmt.Errorf("questions: duplicate row id %q", r.ID)
		case !slots[r.Slot]:
			return nil, fmt.Errorf("questions: row %q names slot %q, which the proto does not have", r.ID, r.Slot)
		case strings.TrimSpace(r.Text) == "":
			return nil, fmt.Errorf("questions: row %q has no text", r.ID)
		case len(r.OptionValues) > 0 && len(r.OptionValues) != len(r.Options):
			return nil, fmt.Errorf("questions: row %q has %d option values for %d options", r.ID, len(r.OptionValues), len(r.Options))
		case r.Order <= 0:
			return nil, fmt.Errorf("questions: row %q has no order", r.ID)
		}
		if r.RepeatOnChange && !r.Repeat {
			return nil, fmt.Errorf("questions: row %q narrows a repeat it does not have", r.ID)
		}
		switch r.When.Format {
		case "", "commander", "sixty":
		default:
			return nil, fmt.Errorf("questions: row %q names format trigger %q, want commander, sixty, or none", r.ID, r.When.Format)
		}
		if r.Key != "" && !keyWord.MatchString(r.Key) {
			return nil, fmt.Errorf("questions: row %q has key %q, want lowercase words and underscores", r.ID, r.Key)
		}
		if other, ok := order[r.Order]; ok {
			return nil, fmt.Errorf("questions: rows %q and %q share order %d", other, r.ID, r.Order)
		}
		seen[r.ID] = true
		order[r.Order] = r.ID
		if placeholder.MatchString(r.Text) {
			if strings.TrimSpace(r.Fallback) == "" {
				return nil, fmt.Errorf("questions: row %q holds a placeholder and has no fallback", r.ID)
			}
			if placeholder.MatchString(r.Fallback) {
				return nil, fmt.Errorf("questions: the fallback of row %q holds a placeholder", r.ID)
			}
		}
		for _, need := range r.When.Requires {
			if !slots[need] {
				return nil, fmt.Errorf("questions: row %q requires slot %q, which the proto does not have", r.ID, need)
			}
		}
		for _, s := range r.When.NotBeside {
			if !slots[s] {
				return nil, fmt.Errorf("questions: row %q keeps slot %q out of its turn, which the proto does not have", r.ID, s)
			}
		}
	}
	// A row may wait on a key only when some row owns that key.
	keys := map[string]bool{}
	for _, r := range c.Rows {
		keys[r.StateKey()] = true
	}
	for _, r := range c.Rows {
		for _, k := range r.When.NotOutstanding {
			if !keys[k] {
				return nil, fmt.Errorf("questions: row %q waits on key %q, which no row owns", r.ID, k)
			}
		}
	}
	sort.SliceStable(c.Rows, func(i, j int) bool { return c.Rows[i].Order < c.Rows[j].Order })
	// The planner reads the rows in order, so a row sees only the slots of
	// the rows before it. A slot a row keeps out of its turn must ask
	// first, or the row shares the turn anyway.
	for _, r := range c.Rows {
		for _, s := range r.When.NotBeside {
			for _, other := range c.Rows {
				if other.Slot == s && other.Order > r.Order {
					return nil, fmt.Errorf("questions: row %q keeps slot %q out of its turn, and row %q of that slot asks after it", r.ID, s, other.ID)
				}
			}
		}
	}
	return &c, nil
}

// keyWord is the shape of a state key.
var keyWord = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

// StateKey is the key the planner tracks for this row.
func (r Row) StateKey() string {
	if r.Key != "" {
		return r.Key
	}
	return r.Slot
}

// Row returns one row by id.
func (c *Catalog) Row(id string) (Row, bool) {
	for _, r := range c.Rows {
		if r.ID == id {
			return r, true
		}
	}
	return Row{}, false
}
