package triage

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/nkramber/decktome/go/internal/harvest"
)

// TestPendingHarvestsReadsEveryUntriagedFile is REV-082: two harvests
// before one triage leave two files, and the triage reads both. A file a
// triage applied is not read again, so no case is written twice.
func TestPendingHarvestsReadsEveryUntriagedFile(t *testing.T) {
	files := []string{"feedback-2026-09-20.jsonl", "feedback-2026-09-21.jsonl", "feedback-2026-09-21b.jsonl"}
	cases := []struct {
		name    string
		triaged []string
		want    []string
	}{
		{"no triage yet", nil, files},
		{"the first harvest triaged", files[:1], files[1:]},
		{"every harvest triaged", files, nil},
		{"an older harvest left behind", files[1:2], []string{files[0], files[2]}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			dir := filepath.Join(root, harvest.Dir)
			if err := os.MkdirAll(dir, 0o750); err != nil {
				t.Fatal(err)
			}
			for _, f := range files {
				if err := os.WriteFile(filepath.Join(dir, f), []byte("{}\n"), 0o600); err != nil {
					t.Fatal(err)
				}
			}
			var marked []string
			for _, f := range tc.triaged {
				marked = append(marked, filepath.Join(dir, f))
			}
			if err := MarkTriaged(root, marked); err != nil {
				t.Fatal(err)
			}
			got, err := PendingHarvests(root)
			if err != nil {
				t.Fatal(err)
			}
			var names []string
			for _, p := range got {
				names = append(names, filepath.Base(p))
			}
			if !slices.Equal(names, tc.want) {
				t.Errorf("pending = %v, want %v", names, tc.want)
			}
		})
	}
}

// TestSettleKeepsAFailedVerdictPending is REV-082 (Codex, #229). A live
// run marks each verdict that it applied, and a harvest is read only when
// each of its verdicts is. A failed verdict keeps its harvest pending, and
// the next run skips the verdicts that already wrote a case.
func TestSettleKeepsAFailedVerdictPending(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, harvest.Dir)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		t.Fatal(err)
	}
	a, b := filepath.Join(dir, "feedback-2026-09-20.jsonl"), filepath.Join(dir, "feedback-2026-09-21.jsonl")
	for _, p := range []string{a, b} {
		if err := os.WriteFile(p, nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	recs := []harvest.Record{{ID: "v1"}, {ID: "v2"}, {ID: "v3"}}
	from := []string{a, b, b}
	results := []Result{{}, {}, {Err: os.ErrDeadlineExceeded}}
	if err := Settle(root, recs, from, results); err != nil {
		t.Fatal(err)
	}
	pending, err := PendingHarvests(root)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(pending, []string{b}) {
		t.Errorf("pending = %v, want the harvest with the failed verdict alone", pending)
	}
	left, err := Unapplied(root, recs)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) != 1 || left[0].ID != "v3" {
		t.Errorf("the next run reads %v, want v3 alone", left)
	}
}
