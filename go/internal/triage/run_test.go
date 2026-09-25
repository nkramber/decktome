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
