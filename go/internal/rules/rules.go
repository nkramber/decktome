// Package rules is the deck referee (roadmap PR-5, guardrail 1).
//
// A pure library: deck plus format plus card index in, findings out.
// Every check is deterministic. No finding is silently dropped, and a
// BLOCK finding means the deck must not reach the user in this state.
package rules

import (
	_ "embed"
	"encoding/json"
	"fmt"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

//go:embed formats.json
var formatsJSON []byte

//go:embed brackets.json
var bracketsJSON []byte

//go:embed companion_bans.json
var companionJSON []byte

// FormatRules is one format's construction rules.
type FormatRules struct {
	ScryfallKey  string `json:"scryfall_key"`
	MinSize      int    `json:"min_size"`
	ExactSize    int    `json:"exact_size"`
	MaxCopies    int    `json:"max_copies"`
	SideboardMax int    `json:"sideboard_max"`
	Commander    bool   `json:"commander"`
}

// Config is the loaded rule data.
type Config struct {
	Formats           map[string]FormatRules `json:"-"`
	MaxGameChangers   map[int32]int          `json:"-"`
	BannedAsCompanion map[string]bool        `json:"-"`
}

// Load parses the embedded rule data. A bad file fails loudly.
func Load() (*Config, error) {
	var f struct {
		Formats map[string]FormatRules `json:"formats"`
	}
	if err := json.Unmarshal(formatsJSON, &f); err != nil {
		return nil, fmt.Errorf("formats.json: %w", err)
	}
	var b struct {
		Brackets map[string]struct {
			MaxGameChangers int `json:"max_game_changers"`
		} `json:"brackets"`
	}
	if err := json.Unmarshal(bracketsJSON, &b); err != nil {
		return nil, fmt.Errorf("brackets.json: %w", err)
	}
	var c struct {
		Banned []string `json:"banned_as_companion"`
	}
	if err := json.Unmarshal(companionJSON, &c); err != nil {
		return nil, fmt.Errorf("companion_bans.json: %w", err)
	}
	cfg := &Config{
		Formats:           f.Formats,
		MaxGameChangers:   map[int32]int{},
		BannedAsCompanion: map[string]bool{},
	}
	for k, v := range b.Brackets {
		var n int32
		if _, err := fmt.Sscanf(k, "%d", &n); err != nil {
			return nil, fmt.Errorf("brackets.json: bad bracket %q", k)
		}
		cfg.MaxGameChangers[n] = v.MaxGameChangers
	}
	for _, name := range c.Banned {
		cfg.BannedAsCompanion[name] = true
	}
	return cfg, nil
}

// CardSource resolves Oracle ids to cards. cards.Index satisfies it.
type CardSource interface {
	ByOracleID(id string) (*mtgv1.Card, bool)
}

// Input is one validation request.
type Input struct {
	Deck *mtgv1.Deck
	// PoolRule decides whether ownership findings apply (D-37).
	PoolRule mtgv1.PoolRule
	// OracleCounts is the owned count per Oracle id. Nil means no
	// collection is attached.
	OracleCounts map[string]int32
	Cards        CardSource
}

// finding codes. The UI and the eval harness key on these.
const (
	CodeDeckSize         = "deck_size"
	CodeSideboardSize    = "sideboard_size"
	CodeCopyLimit        = "copy_limit"
	CodeUnknownCard      = "unknown_card"
	CodeNotLegal         = "not_legal"
	CodeBannedCard       = "banned_card"
	CodeRestrictedCard   = "restricted_card"
	CodeNoCommander      = "no_commander"
	CodeBadCommander     = "bad_commander"
	CodeBadPartner       = "bad_partner"
	CodeOffColor         = "off_color"
	CodeGameChangers     = "game_changer_limit"
	CodeBracketProse     = "bracket_prose_rules"
	CodeBadCompanion     = "bad_companion"
	CodeCompanionBanned  = "companion_banned"
	CodeCompanionUncheck = "companion_condition_unchecked"
	CodeNotOwned         = "not_owned"
	CodeLandCount        = "land_count"
	CodeCurve            = "curve_summary"
	CodeHouseRules       = "house_rules_limited"
)

// Validate runs every check and returns one finding per problem.
func (cfg *Config) Validate(in Input) *mtgv1.ValidationResult {
	res := &mtgv1.ValidationResult{}
	deck := in.Deck
	format := deck.GetFormat().GetId().String()
	fr, ok := cfg.Formats[format]
	if !ok {
		add(res, CodeDeckSize, mtgv1.Severity_SEVERITY_BLOCK, "unknown format "+format, "")
		return finish(res)
	}
	if fr.ScryfallKey == "" {
		add(res, CodeHouseRules, mtgv1.Severity_SEVERITY_INFO,
			"house format: size and copy checks apply, legality checks are off (D-3)", "")
	}
	checkSize(res, deck, fr)
	checkCopies(res, in, fr)
	checkLegality(res, in, fr)
	if fr.Commander {
		checkCommander(res, in)
		checkColorIdentity(res, in)
		checkBracket(cfg, res, in)
	}
	checkCompanion(cfg, res, in)
	checkOwnership(res, in)
	checkManaBase(res, in, fr)
	return finish(res)
}

func add(res *mtgv1.ValidationResult, code string, sev mtgv1.Severity, msg, oracleID string) {
	res.Findings = append(res.Findings, &mtgv1.Finding{Code: code, Severity: sev, Message: msg, OracleId: oracleID})
}

func finish(res *mtgv1.ValidationResult) *mtgv1.ValidationResult {
	res.Passed = true
	for _, f := range res.Findings {
		if f.Severity == mtgv1.Severity_SEVERITY_BLOCK {
			res.Passed = false
			break
		}
	}
	return res
}
