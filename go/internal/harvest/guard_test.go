package harvest

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTheHarvestNeverReadsTheUserRecord holds the guard of D-638. The
// user record carries the reader's email, and no harvest file may hold
// one (D-559). The written-bytes test catches a leak after something
// writes it. This one refuses the join that would write it.
//
// A later session that wants a name or an address on a harvest reads
// this test first, and takes the question to the owner.
func TestTheHarvestNeverReadsTheUserRecord(t *testing.T) {
	for _, name := range []string{"internal/harvest", "cmd/feedback-harvest"} {
		dir := filepath.Join("..", "..", name)
		if _, err := os.Stat(dir); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		fset := token.NewFileSet()
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
				continue
			}
			path := filepath.Join(dir, e.Name())
			file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
			if err != nil {
				t.Fatalf("parse %s: %v", path, err)
			}
			for _, imp := range file.Imports {
				if strings.Contains(imp.Path.Value, "internal/users") {
					t.Errorf("%s imports the user record. A harvest file must hold no email (D-559, D-638). "+
						"Take the question to the owner before you join them.", path)
				}
			}
		}
	}
}
