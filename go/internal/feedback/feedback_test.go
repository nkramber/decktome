package feedback

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
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
		mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT:     5,
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
	if got := KindName(mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT); got != "chat" {
		t.Errorf("KindName(CHAT) = %q", got)
	}
	if KnownReason(mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT, "off_spec") {
		t.Error("a deck reason must not pass as a chat reason")
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

// TestTheSnapshotSurvivesTheStore locks the round trip of D-635. The
// verdict outlives the object it names, so the snapshot must read back
// whole. A document written before D-635 holds no blob, and it must
// still open.
func TestTheSnapshotSurvivesTheStore(t *testing.T) {
	item := Item{
		Kind:    "card",
		Verdict: "down",
		Session: &mtgv1.Session{Id: "s1"},
		Deck: &mtgv1.Deck{
			Id:      "d1",
			Summary: "A lifegain deck.",
			Cards:   []*mtgv1.DeckCard{{OracleId: "o-sol"}, {OracleId: "o-karlov"}},
		},
	}
	sessionGz, deckGz, err := snapshotOf(item)
	if err != nil {
		t.Fatalf("snapshotOf: %v", err)
	}
	if len(sessionGz) == 0 || len(deckGz) == 0 {
		t.Fatal("the snapshot wrote no blob")
	}

	got := itemOf("f1", stored{Kind: "card", Verdict: "down", SessionGz: sessionGz, DeckGz: deckGz})
	// The harvest names the row it wrote, so a read must fill the id.
	if got.ID != "f1" {
		t.Errorf("id = %q, want f1: a harvest can not name the verdict", got.ID)
	}
	if got.Session.GetId() != "s1" {
		t.Errorf("session id = %q, want s1", got.Session.GetId())
	}
	if got.Deck.GetSummary() != "A lifegain deck." {
		t.Errorf("summary = %q, want the stored one", got.Deck.GetSummary())
	}
	if len(got.Deck.GetCards()) != 2 {
		t.Errorf("the deck list holds %d cards, want 2", len(got.Deck.GetCards()))
	}

	// A document from before D-635, and the two real items of 2026-09-09.
	old := itemOf("f1", stored{Kind: "card", Verdict: "down", DeckID: "u8FV7fc98qzNvRfsuJ5q"})
	if old.Session != nil || old.Deck != nil {
		t.Error("a document written before the snapshot read one anyway")
	}
	if old.DeckID != "u8FV7fc98qzNvRfsuJ5q" {
		t.Error("a document written before the snapshot lost its deck id")
	}

	// A blob that does not open leaves the field nil and hides no verdict.
	bad := itemOf("f1", stored{Kind: "card", Verdict: "down", DeckGz: []byte("not gzip")})
	if bad.Deck != nil {
		t.Error("a bad blob read as a deck")
	}
	if bad.Verdict != "down" {
		t.Error("a bad blob hid the verdict")
	}
}
