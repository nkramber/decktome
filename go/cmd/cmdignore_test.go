package cmd_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// `go build ./cmd/X` writes the binary to the working directory when no
// -o is given, so it lands at go/X. Three reached the repo before anyone
// noticed: candidates-review in PR-6, and chat-probe and summary-judge in
// PR-8, for 216 MB together (D-255).
//
// A per-binary list is the wrong shape of fix, because it only covers the
// binaries someone happened to see. This test fails when a command has no
// ignore line, so a new command cannot slip through the same way.
func TestEveryCommandIsIgnored(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("read the command directory: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join("..", "..", ".gitignore"))
	if err != nil {
		t.Fatalf("read .gitignore: %v", err)
	}
	lines := map[string]bool{}
	for _, l := range strings.Split(string(raw), "\n") {
		lines[strings.TrimSpace(l)] = true
	}
	var missing []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if !lines["/go/"+e.Name()] {
			missing = append(missing, e.Name())
		}
	}
	if len(missing) > 0 {
		t.Errorf("commands with no ignore line: %v.\nAdd /go/<name> to .gitignore, or `go build ./cmd/<name>` drops the binary in the repo.", missing)
	}
}

// TestNoCommandBinaryIsPresent catches one that is already there. A
// compiled binary in the module root is 40 to 90 MB, and it rides into
// the next commit that uses git add -A.
func TestNoCommandBinaryIsPresent(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("read the command directory: %v", err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		path := filepath.Join("..", e.Name())
		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			continue
		}
		t.Errorf("a built binary sits at go/%s (%d bytes). Remove it: it is ignored, and it is not source.",
			e.Name(), info.Size())
	}
}
