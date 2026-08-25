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
	// not cancel it. Table tolerance informs power, but the bracket answer
	// must not cancel the question.
	Key string `json:"key"`
	// Order is the ask order. Format first, then theme, then commander
	// (corpus section 11, from the dogfood runs of 2026-08-24).
	Order int `json:"order"`
	// Text may hold {placeholders} that the caller fills before the model
	// phrases the question.
	Text string `json:"text"`
	// Options are suggested answers. The user may answer in free text.
	Options []string `json:"options"`
	// Fallback is the wording used when a placeholder has no value. It
	// holds no placeholder itself. Load enforces that.
	Fallback string `json:"fallback"`
	// Repeat exempts a row from the no-repeat rule. Only the commander
	// pick row uses it: a user who answers "none" gets three new names
	// until one fits (D-73). The row still closes when its key closes.
	Repeat bool `json:"repeat"`
	When   When `json:"when"`
}

// When holds the triggers of one row. A nil pointer means the row does
// not care about that fact. Requires names slots that must be filled
// first, which is how D-67 holds the card-pool question back.
type When struct {
	Words             []string `json:"words"`
	Requires          []string `json:"requires"`
	Format            string   `json:"format"`
	PowerCompetitive  *bool    `json:"power_competitive"`
	OutOfScope        *bool    `json:"out_of_scope"`
	NamedCard         *bool    `json:"named_card"`
	LockedCard        *bool    `json:"locked_card"`
	Suggested         *bool    `json:"suggested"`
	OwnedMode         *bool    `json:"owned_mode"`
	CommanderNotOwned *bool    `json:"commander_not_owned"`
	WeakCommanderPool *bool    `json:"weak_commander_pool"`
	SaltyTheme        *bool    `json:"salty_theme"`
	CommanderSet      *bool    `json:"commander_set"`
	HasCollection     *bool    `json:"has_collection"`
	ThinTheme         *bool    `json:"thin_theme"`
	BuyList           *bool    `json:"buy_list"`
	BudgetAmbiguous   *bool    `json:"budget_ambiguous"`
	HouseFormat       *bool    `json:"house_format"`
	TwoPlans          *bool    `json:"two_plans"`
	AfterBuild        *bool    `json:"after_build"`
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
	"scope":  true,
	"format": true, "power": true, "colors": true, "theme": true,
	"commander": true, "pool_rule": true, "budget": true, "locked": true,
	"plan_variant": true, "house_rules": true, "meta": true,
}

// Load reads the embedded catalog and checks it.
func Load() (*Catalog, error) {
	var c Catalog
	if err := json.Unmarshal(catalogJSON, &c); err != nil {
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
		case r.Order <= 0:
			return nil, fmt.Errorf("questions: row %q has no order", r.ID)
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
	}
	sort.SliceStable(c.Rows, func(i, j int) bool { return c.Rows[i].Order < c.Rows[j].Order })
	return &c, nil
}

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
