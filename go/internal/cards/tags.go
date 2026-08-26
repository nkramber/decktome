package cards

import (
	"encoding/json"
	"io"
)

// Tag is one Scryfall Tagger Oracle tag with its direct card list.
// F-5: tags are community data. They seed candidate lists. They never
// gate a card.
type Tag struct {
	Slug      string
	ParentIDs []string
	ChildIDs  []string
	OracleIDs []string
	id        string
}

// TagIndex resolves a tag slug to Oracle ids, children included.
type TagIndex struct {
	bySlug map[string]*Tag
	byID   map[string]*Tag
}

// LoadTags parses the oracle-tags JSONL stream.
func LoadTags(r io.Reader, name string) (*TagIndex, error) {
	idx := &TagIndex{bySlug: map[string]*Tag{}, byID: map[string]*Tag{}}
	err := readLines(r, name, func(line []byte) error {
		var raw struct {
			ID        string   `json:"id"`
			Slug      string   `json:"slug"`
			ParentIDs []string `json:"parent_ids"`
			ChildIDs  []string `json:"child_ids"`
			Taggings  []struct {
				OracleID string `json:"oracle_id"`
			} `json:"taggings"`
		}
		if err := json.Unmarshal(line, &raw); err != nil {
			return err
		}
		t := &Tag{Slug: raw.Slug, ParentIDs: raw.ParentIDs, ChildIDs: raw.ChildIDs, id: raw.ID}
		for _, g := range raw.Taggings {
			if g.OracleID != "" {
				t.OracleIDs = append(t.OracleIDs, g.OracleID)
			}
		}
		idx.bySlug[t.Slug] = t
		idx.byID[t.ID()] = t
		return nil
	})
	return idx, err
}

// ID returns the Tagger id of the tag.
func (t *Tag) ID() string { return t.id }

// Resolve returns the Oracle ids of a slug, descendants included.
// An unknown slug returns nil: the caller treats it as an empty theme,
// not an error (F-5).
func (x *TagIndex) Resolve(slug string) []string {
	if x == nil {
		return nil
	}
	root, ok := x.bySlug[slug]
	if !ok {
		return nil
	}
	seen := map[string]bool{}
	var out []string
	var walk func(t *Tag)
	walk = func(t *Tag) {
		if seen[t.id] {
			return
		}
		seen[t.id] = true
		out = append(out, t.OracleIDs...)
		for _, child := range t.ChildIDs {
			if c, ok := x.byID[child]; ok {
				walk(c)
			}
		}
	}
	walk(root)
	return out
}

// Len returns the tag count.
func (x *TagIndex) Len() int {
	if x == nil {
		return 0
	}
	return len(x.bySlug)
}

// Has reports whether a slug exists.
func (x *TagIndex) Has(slug string) bool {
	if x == nil {
		return false
	}
	_, ok := x.bySlug[slug]
	return ok
}
