package feedback

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// TestNamesAndReasons pins the short names and the reason keys of
// section 2.4 of the design note. The web test of the feedback feature
// pins the same keys, so a drift on either side fails a test.
func TestNamesAndReasons(t *testing.T) {
	if got := KindName(mtgv1.FeedbackKind_FEEDBACK_KIND_CARD); got != "card" {
		t.Errorf("KindName(CARD) = %q", got)
	}
	if got := KindName(mtgv1.FeedbackKind_FEEDBACK_KIND_UNSPECIFIED); got != "" {
		t.Errorf("KindName(UNSPECIFIED) = %q, want none", got)
	}
	if got := VerdictName(mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN); got != "down" {
		t.Errorf("VerdictName(DOWN) = %q", got)
	}
	if got := VerdictName(mtgv1.FeedbackVerdict(9)); got != "" {
		t.Errorf("VerdictName(9) = %q, want none", got)
	}
	want := map[mtgv1.FeedbackKind]int{
		mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION: 4,
		mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY:  3,
		mtgv1.FeedbackKind_FEEDBACK_KIND_CARD:     5,
		mtgv1.FeedbackKind_FEEDBACK_KIND_DECK:     5,
	}
	for k, n := range want {
		keys := Reasons(k)
		if len(keys) != n {
			t.Errorf("Reasons(%v) = %v, want %d keys", k, keys, n)
		}
		seen := map[string]bool{}
		for _, key := range keys {
			if seen[key] {
				t.Errorf("Reasons(%v) repeats %q", k, key)
			}
			seen[key] = true
			if !KnownReason(k, key) {
				t.Errorf("KnownReason(%v, %q) = false", k, key)
			}
		}
	}
	if KnownReason(mtgv1.FeedbackKind_FEEDBACK_KIND_DECK, "already_answered") {
		t.Error("a question reason must not pass as a deck reason")
	}
	if len(Reasons(mtgv1.FeedbackKind_FEEDBACK_KIND_UNSPECIFIED)) != 0 {
		t.Error("an unknown kind has no reasons")
	}
	// The caller must not reach the table through the copy.
	Reasons(mtgv1.FeedbackKind_FEEDBACK_KIND_DECK)[0] = "changed"
	if Reasons(mtgv1.FeedbackKind_FEEDBACK_KIND_DECK)[0] != "off_spec" {
		t.Error("Reasons must answer a copy")
	}
}
