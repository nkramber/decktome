package feedback

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestIndexFileCoversDown: each query form of Down needs an index of
// collection-group scope, and the emulator enforces none. With no
// verdict, the query orders by created_at alone, and that needs a field
// override (REV-083).
func TestIndexFileCoversDown(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "firestore.indexes.json"))
	if err != nil {
		t.Fatal(err)
	}
	type field struct {
		FieldPath   string `json:"fieldPath"`
		Order       string `json:"order"`
		ArrayConfig string `json:"arrayConfig"`
		QueryScope  string `json:"queryScope"`
	}
	var file struct {
		Indexes []struct {
			CollectionGroup string  `json:"collectionGroup"`
			QueryScope      string  `json:"queryScope"`
			Fields          []field `json:"fields"`
		} `json:"indexes"`
		FieldOverrides []struct {
			CollectionGroup string  `json:"collectionGroup"`
			FieldPath       string  `json:"fieldPath"`
			Indexes         []field `json:"indexes"`
		} `json:"fieldOverrides"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	verdictForm := false
	for _, ix := range file.Indexes {
		if ix.CollectionGroup == "feedback" && ix.QueryScope == "COLLECTION_GROUP" && len(ix.Fields) == 2 &&
			ix.Fields[0].FieldPath == "verdict" && ix.Fields[1] == (field{FieldPath: "created_at", Order: "DESCENDING"}) {
			verdictForm = true
		}
	}
	if !verdictForm {
		t.Error("no composite index for Down with a verdict: verdict, then created_at descending")
	}
	groupOrder := false
	var collectionScope []field
	for _, fo := range file.FieldOverrides {
		if fo.CollectionGroup != "feedback" || fo.FieldPath != "created_at" {
			continue
		}
		for _, ix := range fo.Indexes {
			switch ix.QueryScope {
			case "COLLECTION_GROUP":
				if ix.Order == "DESCENDING" {
					groupOrder = true
				}
			case "COLLECTION":
				collectionScope = append(collectionScope, ix)
			}
		}
		// An override replaces the default indexes of the field, so it
		// must keep the three of the collection scope.
		if len(collectionScope) != 3 {
			t.Errorf("the override keeps %d collection-scope indexes of created_at, want 3", len(collectionScope))
		}
	}
	if !groupOrder {
		t.Error("no collection-group index for Down with no verdict: created_at descending")
	}
}
