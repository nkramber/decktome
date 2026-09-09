package harvest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Names answers the two paths one harvest writes, under root. The date
// names them, and a second harvest on one day takes a letter, so a run
// never writes over a document that exists (D-65).
func Names(root string, day time.Time) (doc, jsonl string, err error) {
	base := filepath.Join(root, Dir)
	stamp := day.UTC().Format("2006-01-02")
	for _, suffix := range []string{"", "b", "c", "d", "e", "f", "g", "h"} {
		doc = filepath.Join(base, "feedback-"+stamp+suffix+".md")
		jsonl = filepath.Join(base, "feedback-"+stamp+suffix+".jsonl")
		_, docErr := os.Stat(doc)
		_, runErr := os.Stat(jsonl)
		if os.IsNotExist(docErr) && os.IsNotExist(runErr) {
			return doc, jsonl, nil
		}
	}
	return "", "", fmt.Errorf("harvest: %s already holds eight harvests", stamp)
}

// Write writes the JSONL file and the document. It makes the directory
// when it is absent, and it writes nothing when the harvest is empty:
// an empty document tells a later watermark nothing, and it hides the
// documents that hold something.
func Write(root string, day time.Time, recs []Record) (doc, jsonl string, err error) {
	if len(recs) == 0 {
		return "", "", nil
	}
	sort.SliceStable(recs, func(i, j int) bool { return recs[i].CreatedAt.Before(recs[j].CreatedAt) })
	if err := os.MkdirAll(filepath.Join(root, Dir), 0o755); err != nil {
		return "", "", fmt.Errorf("harvest: %w", err)
	}
	doc, jsonl, err = Names(root, day)
	if err != nil {
		return "", "", err
	}
	var lines strings.Builder
	for _, r := range recs {
		raw, err := json.Marshal(r)
		if err != nil {
			return "", "", fmt.Errorf("harvest %s: %w", r.ID, err)
		}
		lines.Write(raw)
		lines.WriteString("\n")
	}
	if err := os.WriteFile(jsonl, []byte(lines.String()), 0o600); err != nil {
		return "", "", fmt.Errorf("harvest: %w", err)
	}
	if err := os.WriteFile(doc, []byte(Document(day, recs)), 0o600); err != nil {
		return "", "", fmt.Errorf("harvest: %w", err)
	}
	return doc, jsonl, nil
}

// Document writes what a person reads. The STE check skips a dated
// record, and this file holds the words of readers, which no rule of
// ours governs.
func Document(day time.Time, recs []Record) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Feedback harvest %s\n\n", day.UTC().Format("2006-01-02"))
	b.WriteString("CAUTION: this file holds what a reader wrote. Keep it off any shared page, and put no part of it in an issue or a pull request.\n\n")

	down, up, absent := 0, 0, 0
	byKind := map[string]int{}
	for _, r := range recs {
		if r.Verdict == "down" {
			down++
		} else {
			up++
		}
		if r.Context == "absent" {
			absent++
		}
		byKind[r.Kind]++
	}
	fmt.Fprintf(&b, "%d verdicts: %d down and %d up. %d carry no snapshot of the object they name.\n\n", len(recs), down, up, absent)

	kinds := make([]string, 0, len(byKind))
	for k := range byKind {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)
	b.WriteString("| Kind | Verdicts |\n|---|---|\n")
	for _, k := range kinds {
		fmt.Fprintf(&b, "| %s | %d |\n", k, byKind[k])
	}
	b.WriteString("\n## The verdicts\n\n")

	for _, r := range recs {
		fmt.Fprintf(&b, "### %s %s/%s\n\n", r.CreatedAt.Format("2006-01-02 15:04"), r.Kind, r.Verdict)
		fmt.Fprintf(&b, "- id `%s`, user `%s`\n", r.ID, r.UID)
		if r.SessionID != "" {
			fmt.Fprintf(&b, "- session `%s`\n", r.SessionID)
		}
		if r.DeckID != "" {
			fmt.Fprintf(&b, "- deck `%s`\n", r.DeckID)
		}
		if r.OracleID != "" {
			fmt.Fprintf(&b, "- card `%s`\n", r.OracleID)
		}
		if len(r.Reasons) > 0 {
			fmt.Fprintf(&b, "- reasons: %s\n", strings.Join(r.Reasons, ", "))
		}
		if r.Question != "" {
			fmt.Fprintf(&b, "- asked: %s\n", r.Question)
		}
		if r.Answer != "" {
			fmt.Fprintf(&b, "- answered: %s\n", r.Answer)
		}
		if r.Text != "" {
			fmt.Fprintf(&b, "- said: %s\n", r.Text)
		}
		switch r.Context {
		case "snapshot":
			fmt.Fprintf(&b, "- context: the verdict carries the object it names\n")
		default:
			fmt.Fprintf(&b, "- context: **absent**. The verdict predates D-635, or its kind names no object. The triage reads the words alone.\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}
