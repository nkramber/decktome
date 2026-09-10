package main

import (
	"encoding/json"
	"testing"
)

// TestTheAuditWritesAListAndNeverNull holds the shape of the file the
// -audit-out flag writes. A nil slice marshals as the word null, and a
// reader of the file expects a list. A run that walks no own-copy pair
// is the case: a model with no precon in the holdout.
func TestTheAuditWritesAListAndNeverNull(t *testing.T) {
	if auditRows == nil {
		t.Fatal("auditRows starts nil, so an empty run writes null")
	}
	raw, err := json.Marshal(auditRows)
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != "[]" {
		t.Errorf("an empty audit marshals as %s, want []", raw)
	}
}
