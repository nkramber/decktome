package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/collections"
)

// turns gives the user's turns. A JSON file wins over the text with |,
// because a reader's text can hold a | (D-1144).
func turns(text, jsonPath string) ([]string, error) {
	if jsonPath == "" {
		var out []string
		for _, m := range strings.Split(text, "|") {
			out = append(out, strings.TrimSpace(m))
		}
		return out, nil
	}
	b, err := os.ReadFile(jsonPath) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, err
	}
	var out []string
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, fmt.Errorf("messages %s: %w", jsonPath, err)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("messages %s: no turn", jsonPath)
	}
	return out, nil
}

// ownedFromJSON reads a stored collection as protojson, as
// `live-evals replay-input` writes it, and counts the copies per Oracle
// id (D-1144).
func ownedFromJSON(path string) (map[string]int32, error) {
	b, err := os.ReadFile(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, err
	}
	var c mtgv1.Collection
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("collection %s: %w", path, err)
	}
	if len(c.GetEntries()) == 0 {
		return nil, fmt.Errorf("collection %s: no entry", path)
	}
	return collections.OracleCounts(c.GetEntries()), nil
}
