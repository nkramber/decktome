package triage

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

// The manifest is the hand-off from the triage of PR-28b to the fix
// cycle of PR-28c (D-645). The triage writes the cases into the gate
// files, and the cycle has to know which ids it just added, on which
// gate, and what the reader was unhappy about.
//
// It is a file and never a git diff. A diff reads the lines a commit
// changed, and it can not tell a case of this run from a case a person
// wrote by hand in the same commit.

// Gate names the suite that measures a case. The words are the suite
// names of the eval run files, so a run file and a manifest agree.
const (
	GateQuestions = "questions"
	GateDecks     = "decks"
	GateBrackets  = "brackets"
)

// GateOf answers the gate that owns a target file.
func GateOf(target string) string {
	switch target {
	case TargetConversations:
		return GateQuestions
	case TargetDeckPrompts:
		return GateDecks
	case TargetBracketPrompts:
		return GateBrackets
	}
	return ""
}

// Entry is one case of the manifest.
type Entry struct {
	Class string `json:"class"`
	// From is the verdict the case came from, so the cycle can read the
	// reader's own words out of the harvest again.
	From   string `json:"from"`
	Gate   string `json:"gate"`
	Target string `json:"target"`
	ID     int    `json:"id"`
	// Where names the code or the prompt the fix lives in, from the class
	// table. The fixer reads it as a starting point and never as an
	// instruction: the class is a guess about the fault, not about the
	// cause.
	Where string `json:"where,omitempty"`
	// Said is what the reader checked and wrote, in their words.
	Said string `json:"said,omitempty"`
	// Gaps are what the case could not fill. A case with one measures
	// less than the reader reported.
	Gaps []string `json:"gaps,omitempty"`
}

// Manifest names every case one triage wrote.
type Manifest struct {
	WrittenAt time.Time `json:"written_at"`
	// Harvest is the JSONL file the triage read.
	Harvest string  `json:"harvest,omitempty"`
	Cases   []Entry `json:"cases"`
}

// ManifestOf reads the cases of a finished run. A result with no case,
// or one no gate measures, is left out: the manifest is what the cycle
// can run a gate on.
// The case list starts empty and never nil. A nil slice marshals as
// null, and the cycle reads the length of that list to decide whether
// there is anything to fix. A null would raise instead of reading zero,
// and the cycle would go on to a fixer with no case (D-645).
func ManifestOf(harvest string, day time.Time, rs []Result) Manifest {
	m := Manifest{WrittenAt: day.UTC(), Harvest: harvest, Cases: []Entry{}}
	for _, r := range rs {
		gate := GateOf(r.Case.Target)
		if r.Err != nil || gate == "" || r.Case.ID == 0 {
			continue
		}
		m.Cases = append(m.Cases, Entry{
			Class:  r.Case.Class,
			From:   r.Route.Record.ID,
			Gate:   gate,
			Target: r.Case.Target,
			ID:     r.Case.ID,
			Where:  r.Route.Class.Where,
			Said:   readerSaid(r.Route.Record),
			Gaps:   r.Case.Gaps,
		})
	}
	return m
}

// IDs lists the ids of one gate, in order, as the -only flag reads them.
func (m Manifest) IDs(gate string) []int {
	var out []int
	for _, c := range m.Cases {
		if c.Gate == gate {
			out = append(out, c.ID)
		}
	}
	sort.Ints(out)
	return out
}

// Gates lists the gates the manifest names, in a stable order.
func (m Manifest) Gates() []string {
	seen := map[string]bool{}
	var out []string
	for _, g := range []string{GateQuestions, GateDecks, GateBrackets} {
		for _, c := range m.Cases {
			if c.Gate == g && !seen[g] {
				seen[g], out = true, append(out, g)
			}
		}
	}
	return out
}

// WriteManifest writes the manifest to path.
func WriteManifest(path string, m Manifest) error {
	raw, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("triage: manifest: %w", err)
	}
	// The manifest is a file of this repo and not a secret.
	if err := os.WriteFile(path, append(raw, '\n'), 0o644); err != nil { // #nosec G306
		return fmt.Errorf("triage: %s: %w", path, err)
	}
	return nil
}

// ReadManifest reads a manifest the triage wrote.
func ReadManifest(path string) (Manifest, error) {
	raw, err := os.ReadFile(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return Manifest{}, fmt.Errorf("triage: %s: %w", path, err)
	}
	var m Manifest
	if err := json.Unmarshal(raw, &m); err != nil {
		return Manifest{}, fmt.Errorf("triage: %s: %w", path, err)
	}
	return m, nil
}
