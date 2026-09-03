package cards

import (
	"encoding/json"
	"io"
	"sort"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// RulingsFile is the rulings bulk file of a snapshot (PR-20). A snapshot
// stored before it existed holds none, and the index then answers no
// ruling, which the log says.
const RulingsFile = "rulings.jsonl.gz"

// Ruling is one Scryfall ruling of a card. The file keys it by Oracle
// id, so every printing shares it.
type Ruling struct {
	// PublishedAt is the ISO date of the ruling.
	PublishedAt string
	Comment     string
	// Source is "wotc" or "scryfall".
	Source string
}

// LoadRulings parses the rulings JSONL stream into a map by Oracle id,
// each list in file order, which is by date.
func LoadRulings(r io.Reader, name string) (map[string][]Ruling, error) {
	out := map[string][]Ruling{}
	err := readLines(r, name, func(line []byte) error {
		var raw struct {
			OracleID    string `json:"oracle_id"`
			Source      string `json:"source"`
			PublishedAt string `json:"published_at"`
			Comment     string `json:"comment"`
		}
		if err := json.Unmarshal(line, &raw); err != nil {
			return err
		}
		if raw.OracleID == "" || raw.Comment == "" {
			return nil
		}
		out[raw.OracleID] = append(out[raw.OracleID], Ruling{PublishedAt: raw.PublishedAt, Comment: raw.Comment, Source: raw.Source})
		return nil
	})
	for _, list := range out {
		sort.SliceStable(list, func(i, j int) bool { return list[i].PublishedAt < list[j].PublishedAt })
	}
	return out, err
}

// WithRulings supplies the rulings of the snapshot's rulings file.
func WithRulings(rulings map[string][]Ruling) IndexOption {
	return func(o *indexOpts) { o.rulings = rulings }
}

// Rulings answers the rulings of one card, oldest first, as a copy. A
// card with none, or a snapshot with no rulings file, answers nil.
func (x *Index) Rulings(oracleID string) []Ruling {
	list := x.rulings[oracleID]
	if len(list) == 0 {
		return nil
	}
	return append([]Ruling(nil), list...)
}

// HasRulings says whether the snapshot carried a rulings file.
func (x *Index) HasRulings() bool { return x.rulings != nil }

// PrintingsOf answers every playable printing of one card, as a copy in
// file order. A digital printing is among them, with its mark (D-306).
func (x *Index) PrintingsOf(oracleID string) []*mtgv1.Printing {
	ids := x.printingsOf[oracleID]
	if len(ids) == 0 {
		return nil
	}
	out := make([]*mtgv1.Printing, 0, len(ids))
	for _, id := range ids {
		if p, ok := x.printings[id]; ok {
			out = append(out, p)
		}
	}
	return out
}
