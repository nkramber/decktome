package triage

import (
	"context"
	"encoding/json"
	"fmt"
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

// NewestHarvest names the newest JSONL file the harvest wrote under
// root. It answers an empty string when no harvest exists.
func NewestHarvest(root string) (string, error) {
	paths, err := filepath.Glob(filepath.Join(root, harvest.Dir, "*.jsonl"))
	if err != nil {
		return "", err
	}
	if len(paths) == 0 {
		return "", nil
	}
	// The names carry the date and a letter, so the last name in order is
	// the newest file (harvest.Names).
	sort.Strings(paths)
	return paths[len(paths)-1], nil
}
