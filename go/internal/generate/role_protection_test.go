package generate

import (
	"encoding/json"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestProtectionIsAJob is D-1120. A permanent that protects a creature
// takes the role protection, and the shortlist, the schema, and the deck
// list carry the word both ways.
func TestProtectionIsAJob(t *testing.T) {
	if got := roleWord(mtgv1.CardRole_CARD_ROLE_PROTECTION); got != "protection" {
		t.Errorf("shortlist job of the protection role = %q", got)
	}
	if got := cardRole("protection"); got != mtgv1.CardRole_CARD_ROLE_PROTECTION {
		t.Errorf("the job word protection reads as %v", got)
	}
}

// TestSchemaJobsMatchTheRoleWords holds the schema enum and the role
// words together. A shortlist job the schema refuses fails a strict call,
// and a schema word the reader does not know reads as other.
func TestSchemaJobsMatchTheRoleWords(t *testing.T) {
	var schema struct {
		Properties map[string]struct {
			Items struct {
				Properties struct {
					Role struct {
						Enum []string `json:"enum"`
					} `json:"role"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"properties"`
	}
	if err := json.Unmarshal([]byte(deckSchema), &schema); err != nil {
		t.Fatalf("schema: %v", err)
	}
	for _, list := range []string{"cards", "sideboard"} {
		enum := schema.Properties[list].Items.Properties.Role.Enum
		if len(enum) == 0 {
			t.Fatalf("the %s list has no role enum", list)
		}
		seen := map[string]bool{}
		for _, w := range enum {
			seen[w] = true
			if _, ok := roleNames[w]; !ok {
				t.Errorf("the %s enum holds %q, which cardRole does not read", list, w)
			}
		}
		for _, w := range roleWords {
			if !seen[w] {
				t.Errorf("the shortlist job %q is not in the %s enum", w, list)
			}
		}
	}
}
