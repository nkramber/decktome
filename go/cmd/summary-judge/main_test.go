package main

import "testing"

// TestDeckBlocksReadsPerDeck covers the parser. A deck with no summary
// is reported as such, and it never takes the next deck's summary.
func TestDeckBlocksReadsPerDeck(t *testing.T) {
	doc := `# PR-8 deck gate

### 1. first deck

Cards: 100 main.

**Summary:** the first summary

### 2. no summary here

ERROR: the prompt failed

### 3. third deck

**Summary:** the third summary
`
	got := deckBlocks(doc)
	if len(got) != 3 {
		t.Fatalf("blocks = %d, want 3: %+v", len(got), got)
	}
	want := []deckBlock{
		{ID: "1", Name: "first deck", Summary: "the first summary"},
		{ID: "2", Name: "no summary here", Summary: ""},
		{ID: "3", Name: "third deck", Summary: "the third summary"},
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("block %d = %+v, want %+v", i, got[i], want[i])
		}
	}
	if len(deckBlocks("no headings")) != 0 {
		t.Error("a document with no heading gave a block")
	}
}
