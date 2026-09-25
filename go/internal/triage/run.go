package triage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/nkramber/decktome/go/internal/harvest"
)

// Run routes every verdict, asks the judge where a route needs one, and
// writes the case of each (PR-28b).
//
// A nil judge is the dry lane: it calls no model, so a route that needs
// one keeps its need and writes no case. That is the free gate, and it
// moves between no two runs.
//
// One bad verdict never stops the run. Its Result carries the error, the
// document names it, and the verdict of the run fails.
func Run(ctx context.Context, recs []harvest.Record, judge Judger, name Namer, nextID func(target string) (int, error)) []Result {
	ids := map[string]int{}
	out := make([]Result, 0, len(recs))
	for _, rec := range recs {
		r := RouteOf(rec)
		if r.Need == NeedJudge && judge != nil {
			v, err := judge(ctx, rec, r)
			if err != nil {
				out = append(out, Result{Route: r, Err: err})
				continue
			}
			r = r.Apply(v)
		}
		res := Result{Route: r}
		if r.Need == NeedJudge || (r.Class.ID == "" && !r.Keep && !r.Owner) {
			out = append(out, res)
			continue
		}
		target := targetOf(r)
		id := 0
		if target != "" && nextID != nil {
			n, err := nextID(target)
			if err != nil {
				out = append(out, Result{Route: r, Err: err})
				continue
			}
			// A second case for one file takes the next number, so two
			// cases of one run never share an id.
			if ids[target] >= n {
				n = ids[target]
			}
			ids[target] = n + 1
			id = n
		}
		c, err := CaseOf(r, id, name)
		if err != nil {
			out = append(out, Result{Route: r, Err: err})
			continue
		}
		res.Case = c
		out = append(out, res)
	}
	return out
}

// targetOf names the file a route's case joins, before the case exists.
// The id of the case comes from that file, so the driver needs the name
// first.
func targetOf(r Route) string {
	art := r.Class.Artifact
	if r.Keep {
		switch r.Record.Kind {
		case "question", "chat":
			art = AConversation
		default:
			art = ADeckPrompt
		}
	}
	if r.Owner {
		return ""
	}
	switch art {
	case AConversation:
		return TargetConversations
	case ADeckPrompt, ASummaryCase:
		return TargetDeckPrompts
	case ABracketPrompt:
		return TargetBracketPrompts
	case AParseFixture:
		if f, err := FaultOf(r.Record); err == nil {
			return fixtureTarget(f)
		}
	}
	return ""
}

// ReadHarvest reads one JSONL file of the harvest.
func ReadHarvest(path string) ([]harvest.Record, error) {
	raw, err := os.ReadFile(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, fmt.Errorf("triage: %s: %w", path, err)
	}
	var out []harvest.Record
	for i, line := range strings.Split(string(raw), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var rec harvest.Record
		if err := json.Unmarshal([]byte(line), &rec); err != nil {
			return nil, fmt.Errorf("triage: %s line %d: %w", path, i+1, err)
		}
		out = append(out, rec)
	}
	return out, nil
}

// TriagedRecord is the file under harvest.Dir that names each harvest a
// live triage applied, one base name a line. It is no JSONL file, so the
// harvest watermark never reads it.
const TriagedRecord = "triaged.txt"

// PendingHarvests names every JSONL file of the harvest under root that
// no live triage applied, oldest first. Two harvests before one triage
// leave two files, and the triage reads both (REV-082).
func PendingHarvests(root string) ([]string, error) {
	paths, err := filepath.Glob(filepath.Join(root, harvest.Dir, "*.jsonl"))
	if err != nil {
		return nil, err
	}
	done, err := triaged(root)
	if err != nil {
		return nil, err
	}
	// The names carry the date and a letter, so name order is age order
	// (harvest.Names).
	sort.Strings(paths)
	var out []string
	for _, p := range paths {
		if !done[filepath.Base(p)] {
			out = append(out, p)
		}
	}
	return out, nil
}

// MarkTriaged adds each path to the record, so the next triage does not
// write its cases again.
func MarkTriaged(root string, paths []string) error {
	if len(paths) == 0 {
		return nil
	}
	done, err := triaged(root)
	if err != nil {
		return err
	}
	var add strings.Builder
	for _, p := range paths {
		if name := filepath.Base(p); !done[name] {
			done[name] = true
			add.WriteString(name + "\n")
		}
	}
	if err := os.MkdirAll(filepath.Join(root, harvest.Dir), 0o750); err != nil {
		return fmt.Errorf("triage: %w", err)
	}
	path := filepath.Join(root, harvest.Dir, TriagedRecord)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600) // #nosec G304 -- a fixed name under the repo root.
	if err != nil {
		return fmt.Errorf("triage: %w", err)
	}
	if _, err := f.WriteString(add.String()); err != nil {
		_ = f.Close()
		return fmt.Errorf("triage: %s: %w", path, err)
	}
	return f.Close()
}

// triaged reads the record. An absent record names no file.
func triaged(root string) (map[string]bool, error) {
	path := filepath.Join(root, harvest.Dir, TriagedRecord)
	raw, err := os.ReadFile(path) // #nosec G304 -- a fixed name under the repo root.
	if errors.Is(err, fs.ErrNotExist) {
		return map[string]bool{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("triage: %s: %w", path, err)
	}
	done := map[string]bool{}
	for _, line := range strings.Split(string(raw), "\n") {
		if name := strings.TrimSpace(line); name != "" {
			done[name] = true
		}
	}
	return done, nil
}
