package triage

import (
	"testing"

	"github.com/nkramber/decktome/go/internal/harvest"
)

// TestAGeneralNoteCallsNoJudge is D-1078: a note judges the app and not
// a build, so the triage spends nothing on it and writes no case.
func TestAGeneralNoteCallsNoJudge(t *testing.T) {
	r := RouteOf(harvest.Record{Kind: "general", Text: "The deck screen is slow."})
	if r.Need != NeedNothing || r.Keep || r.Class.ID != "" {
		t.Errorf("route = %+v", r)
	}
}
