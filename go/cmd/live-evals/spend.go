package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/nkramber/decktome/go/internal/livespend"
)

// spendReport is the answer of the spend command. Charged is the dollar
// figure the budget reads, and Text is the account for a notice.
type spendReport struct {
	livespend.Sum
	Charged float64 `json:"charged"`
	Text    string  `json:"text"`
}

// spend sums the provider spend of one item. The marked lines of the
// session logs in logs are the source, and file is the spend file of the
// bundle, a cross-check. The session can edit the bundle, and it can not
// edit a log (D-1170). With no logs, a session reads its own spend file.
func spend(logs string, budget float64, file io.Reader, out io.Writer) error {
	g := livespend.NewLedger()
	var paths []string
	if logs != "" {
		var err error
		if paths, err = filepath.Glob(filepath.Join(logs, "session-*.log")); err != nil {
			return err
		}
	}
	sort.Strings(paths)
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return err
		}
		err = g.AddLog(f)
		_ = f.Close()
		if err != nil {
			return fmt.Errorf("%s: %w", p, err)
		}
	}
	if err := g.AddFile(file); err != nil {
		return err
	}
	s := g.Sum(budget, logs != "")
	return json.NewEncoder(out).Encode(spendReport{Sum: s, Charged: s.Charged(), Text: s.Text()})
}
