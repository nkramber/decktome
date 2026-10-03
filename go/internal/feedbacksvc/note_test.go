package feedbacksvc

import (
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/auth"
)

const kindGeneral = mtgv1.FeedbackKind_FEEDBACK_KIND_GENERAL

// TestAGeneralNoteIsStored is D-1078: a note keeps its text, its screen,
// and the deck of a deck screen with its snapshot (D-635). It moves no
// counter of the user record, and it pings the owner.
func TestAGeneralNoteIsStored(t *testing.T) {
	store := &fakeStore{}
	deckList := fakeDecks{"u1": {"d1": &mtgv1.Deck{Id: "d1", Imported: true}}}
	noter := &fakeNoter{}
	pings := &fakeNotifier{}
	s := New(store, fakeSessions{}, deckList, auth.UserID, WithClock(func() time.Time { return at }), WithUsers(noter), WithNotifier(pings))
	err := submitAs(s, "u1", "ann@example.com", &mtgv1.Feedback{Kind: kindGeneral, Text: "  The deck screen is slow.  ", Screen: "deck", DeckId: "d1"})
	if err != nil {
		t.Fatalf("submit: %v", err)
	}
	if len(store.items) != 1 {
		t.Fatalf("stored %d items", len(store.items))
	}
	got := store.items[0]
	if got.Kind != "general" || got.Verdict != "" || got.Text != "The deck screen is slow." || got.Screen != "deck" || got.DeckID != "d1" || got.Deck.GetId() != "d1" {
		t.Errorf("item = %+v", got)
	}
	if len(noter.counts) != 0 {
		t.Errorf("a note moved a counter: %v", noter.counts)
	}
	if len(pings.notices) != 1 || !strings.Contains(pings.notices[0].Title, "feedback, general") || !strings.Contains(pings.notices[0].Message, "screen: deck") {
		t.Errorf("notices = %+v", pings.notices)
	}
}

func TestAGeneralNoteRefuses(t *testing.T) {
	for _, tc := range []struct {
		name string
		fb   *mtgv1.Feedback
		want connect.Code
	}{
		{"no text", &mtgv1.Feedback{Kind: kindGeneral, Screen: "build"}, connect.CodeInvalidArgument},
		{"a verdict", &mtgv1.Feedback{Kind: kindGeneral, Verdict: up, Text: "x", Screen: "build"}, connect.CodeInvalidArgument},
		{"no screen", &mtgv1.Feedback{Kind: kindGeneral, Text: "x"}, connect.CodeInvalidArgument},
		{"an unknown screen", &mtgv1.Feedback{Kind: kindGeneral, Text: "x", Screen: "kitchen"}, connect.CodeInvalidArgument},
		{"a reason", &mtgv1.Feedback{Kind: kindGeneral, Text: "x", Screen: "build", Reasons: []string{"stuck"}}, connect.CodeInvalidArgument},
		{"a question", &mtgv1.Feedback{Kind: kindGeneral, Text: "x", Screen: "chat", SessionId: "s1", QuestionId: "q1"}, connect.CodeInvalidArgument},
		{"a deck of another user", &mtgv1.Feedback{Kind: kindGeneral, Text: "x", Screen: "deck", DeckId: "d1"}, connect.CodePermissionDenied},
		{"a screen on a verdict", &mtgv1.Feedback{Kind: kindDeck, Verdict: up, DeckId: "d1", Screen: "deck"}, connect.CodeInvalidArgument},
	} {
		t.Run(tc.name, func(t *testing.T) {
			store, s := fixture()
			uid := "u1"
			if tc.want == connect.CodePermissionDenied {
				uid = "u2"
			}
			_, err := submit(s, uid, tc.fb)
			if err == nil {
				t.Fatal("submit passed")
			}
			if got := codeOf(t, err); got != tc.want {
				t.Errorf("code = %v, want %v (%v)", got, tc.want, err)
			}
			if len(store.items) != 0 {
				t.Error("a refused note was stored")
			}
		})
	}
}
