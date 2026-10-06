package main

import (
	"encoding/json"
	"io"

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
// session logs in logs and the append-only ledger are the sources, and
// file is the spend file of the bundle, a cross-check. The session can
// edit the bundle, and it can not edit a log or lower the ledger (D-1170,
// D-1171). With logs, a run with no end line in a log is unmeasured
// (D-1173). With no logs, a session reads its own spend file.
func spend(logs, ledger string, budget float64, file io.Reader, out io.Writer) error {
	g := livespend.NewLedger()
	if logs != "" {
		if err := g.AddLogs(logs); err != nil {
			return err
		}
	}
	if ledger != "" {
		if err := g.AddPath(ledger); err != nil {
			return err
		}
	}
	if err := g.AddFile(file); err != nil {
		return err
	}
	s := g.Sum(budget, logs != "")
	return json.NewEncoder(out).Encode(spendReport{Sum: s, Charged: s.Charged(), Text: s.Text()})
}
