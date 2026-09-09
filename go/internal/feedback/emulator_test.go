package feedback

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The store talks to Firestore, and the write and the read back can not
// be proven without it. These tests need the local emulator, so CI skips
// them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/feedback -count=1

func emulatorRepo(t *testing.T) *Repo {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the store against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return NewRepo(client)
}

// TestFeedbackRoundTrip is D-558: a verdict reads back whole under its
// user, with the reasons, the prompt versions, and the time. Another
// user's read of the same id is ErrNotFound, as is an unknown id.
func TestFeedbackRoundTrip(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := "u-" + time.Now().UTC().Format("150405.000000")
	at := time.Date(2026, 9, 6, 15, 30, 0, 0, time.UTC)
	item := Item{
		Kind:      "card",
		Verdict:   "down",
		DeckID:    "d1",
		OracleID:  "o-sol",
		Reasons:   []string{"off_theme", "wrong_power"},
		Text:      "Sol Ring is banned at my table.",
		Prompts:   map[string]int64{"questions": 18, "generate": 12},
		CreatedAt: at,
	}
	id, err := r.Add(ctx, uid, item)
	if err != nil || id == "" {
		t.Fatalf("Add = %q, %v", id, err)
	}
	got, err := r.Get(ctx, uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if got.Kind != "card" || got.Verdict != "down" || got.DeckID != "d1" || got.OracleID != "o-sol" || got.Text != item.Text {
		t.Errorf("Get = %+v, want the item back", got)
	}
	if len(got.Reasons) != 2 || got.Reasons[0] != "off_theme" || got.Reasons[1] != "wrong_power" {
		t.Errorf("reasons = %v", got.Reasons)
	}
	if got.Prompts["questions"] != 18 || got.Prompts["generate"] != 12 {
		t.Errorf("prompts = %v", got.Prompts)
	}
	if !got.CreatedAt.Equal(at) {
		t.Errorf("created_at = %v, want %v", got.CreatedAt, at)
	}
	if _, err := r.Get(ctx, uid+"-other", id); !errors.Is(err, ErrNotFound) {
		t.Errorf("another user's read = %v, want ErrNotFound", err)
	}
	if _, err := r.Get(ctx, uid, "nope"); !errors.Is(err, ErrNotFound) {
		t.Errorf("an unknown id = %v, want ErrNotFound", err)
	}
	// A thumbs up carries no reason and no text, and it reads back so.
	up, err := r.Add(ctx, uid, Item{Kind: "deck", Verdict: "up", DeckID: "d1", Prompts: item.Prompts, CreatedAt: at})
	if err != nil {
		t.Fatal(err)
	}
	if got, err := r.Get(ctx, uid, up); err != nil || len(got.Reasons) != 0 || got.Text != "" {
		t.Errorf("the thumbs up = %+v, %v", got, err)
	}
}

// TestDownReadsEveryUser is D-596. The documents sit under each user,
// and the owner asked to read every negative verdict without a walk of
// the user list. A collection group query answers them all at once, so
// no caller iterates users and no second store is needed.
func TestDownReadsEveryUser(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	stamp := time.Now().UTC().Format("150405.000000")
	at := time.Date(2026, 9, 7, 15, 30, 0, 0, time.UTC)

	// Two users, one down verdict each, and one up that must not show.
	for i, u := range []string{"g1-" + stamp, "g2-" + stamp} {
		if _, err := r.Add(ctx, u, Item{
			Kind: "chat", Verdict: "down", SessionID: "s1",
			Reasons: []string{"stuck"}, CreatedAt: at.Add(time.Duration(i) * time.Minute),
		}); err != nil {
			t.Fatalf("Add down: %v", err)
		}
	}
	upUser := "g3-" + stamp
	if _, err := r.Add(ctx, upUser, Item{Kind: "deck", Verdict: "up", DeckID: "d1", CreatedAt: at}); err != nil {
		t.Fatalf("Add up: %v", err)
	}

	items, err := r.Down(ctx, "down", 100)
	if err != nil {
		t.Fatalf("Down: %v", err)
	}
	seen := map[string]bool{}
	for _, it := range items {
		if it.Verdict != "down" {
			t.Errorf("Down answered a %q verdict", it.Verdict)
		}
		seen[it.UID] = true
	}
	// Both users of this run are in one answer, and neither was named.
	for _, u := range []string{"g1-" + stamp, "g2-" + stamp} {
		if !seen[u] {
			t.Errorf("the query missed the verdict of %s", u)
		}
	}
	if seen[upUser] {
		t.Error("an up verdict came back from the down query")
	}
}

// TestSinceReadsTenSeededItems is the emulator gate item of PR-28a. Ten
// verdicts go in, both up and down, and the harvest reads the ones after
// a watermark, oldest first. The snapshot of D-635 must survive the
// round trip, or the triage of PR-28b has no deck list to read.
func TestSinceReadsTenSeededItems(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := "u-" + time.Now().UTC().Format("150405.000000")
	base := time.Date(2026, 9, 9, 12, 0, 0, 0, time.UTC)

	for i := range 10 {
		item := Item{
			Kind:      "card",
			Verdict:   "down",
			UID:       uid,
			DeckID:    "d1",
			OracleID:  "o-sol",
			Reasons:   []string{"off_theme"},
			CreatedAt: base.Add(time.Duration(i) * time.Minute),
		}
		if i%2 == 0 {
			item.Verdict = "up"
		}
		// Half carry the snapshot, and half read as a verdict from before
		// D-635. Both must harvest.
		if i >= 5 {
			item.Deck = &mtgv1.Deck{Id: "d1", Cards: []*mtgv1.DeckCard{{OracleId: "o-sol"}}}
		}
		if _, err := r.Add(ctx, uid, item); err != nil {
			t.Fatalf("add %d: %v", i, err)
		}
	}

	all, err := r.Since(ctx, time.Time{}, 0)
	if err != nil {
		t.Fatalf("Since: %v", err)
	}
	mine := make([]Item, 0, 10)
	for _, it := range all {
		if it.UID == uid {
			mine = append(mine, it)
		}
	}
	if len(mine) != 10 {
		t.Fatalf("read %d verdicts, want the 10 seeded", len(mine))
	}
	// Oldest first, so a document tells the story in order.
	for i := 1; i < len(mine); i++ {
		if mine[i].CreatedAt.Before(mine[i-1].CreatedAt) {
			t.Fatalf("verdict %d came before %d", i, i-1)
		}
	}
	ups := 0
	snaps := 0
	for _, it := range mine {
		if it.Verdict == "up" {
			ups++
		}
		if it.Deck != nil {
			snaps++
			if len(it.Deck.GetCards()) != 1 {
				t.Error("the snapshot lost the deck list")
			}
		}
		if it.ID == "" {
			t.Error("a verdict read back with no id, so no harvest can name it")
		}
	}
	if ups != 5 {
		t.Errorf("read %d thumbs up, want 5: the harvest reads both verdicts", ups)
	}
	if snaps != 5 {
		t.Errorf("read %d snapshots, want 5", snaps)
	}

	// A watermark inside the run reads the newer verdicts alone.
	after, err := r.Since(ctx, base.Add(4*time.Minute), 0)
	if err != nil {
		t.Fatalf("Since after: %v", err)
	}
	n := 0
	for _, it := range after {
		if it.UID == uid {
			n++
		}
	}
	if n != 5 {
		t.Errorf("a watermark at the 5th verdict read %d, want 5", n)
	}
}

// TestALimitedHarvestLeavesNoGap locks F-88. The watermark of a harvest
// is the newest time it wrote. So a chunk must hold the oldest rows
// after the floor: a chunk of the newest rows moves the floor past every
// row under it, and no later harvest reads those again.
//
// It walks chunks the way a harvest does, and it asks one thing of the
// walk: every verdict of this run reaches it. The store holds the rows
// of every other test as well, so a chunk carries what it carries. The
// invariant holds either way, and it fails on the code F-88 names.
func TestALimitedHarvestLeavesNoGap(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := "u-" + time.Now().UTC().Format("150405.000000")
	base := time.Date(2026, 9, 9, 18, 0, 0, 0, time.UTC)
	// The rows go away with the test. Two runs seed the same instants
	// under two users, and a floor that lands on one of them skips the
	// other, because the read takes what is after it and not what is on
	// it. A store that keeps the rows of the last run fails the next one.
	t.Cleanup(func() {
		docs, err := r.col(uid).Documents(context.Background()).GetAll()
		if err != nil {
			t.Logf("cleanup: %v", err)
			return
		}
		for _, d := range docs {
			if _, err := d.Ref.Delete(context.Background()); err != nil {
				t.Logf("cleanup %s: %v", d.Ref.ID, err)
			}
		}
	})

	const seeded = 10
	for i := range seeded {
		item := Item{
			Kind: "deck", Verdict: "down", UID: uid, DeckID: "d1",
			Reasons:   []string{"off_spec"},
			CreatedAt: base.Add(time.Duration(i) * time.Minute),
		}
		if _, err := r.Add(ctx, uid, item); err != nil {
			t.Fatalf("add %d: %v", i, err)
		}
	}

	// The floor starts under the first verdict, the way a watermark from
	// an earlier harvest would.
	floor := base.Add(-time.Second)
	seen := map[string]bool{}
	// Every chunk moves the floor forward, so the walk ends. The bound
	// keeps a fault from running it for ever.
	for chunk := 0; chunk < 50; chunk++ {
		got, err := r.Since(ctx, floor, 5)
		if err != nil {
			t.Fatalf("chunk %d: %v", chunk, err)
		}
		if len(got) == 0 {
			break
		}
		for i := 1; i < len(got); i++ {
			if got[i].CreatedAt.Before(got[i-1].CreatedAt) {
				t.Fatalf("chunk %d is not oldest first", chunk)
			}
		}
		for _, it := range got {
			if it.UID == uid {
				seen[it.ID] = true
			}
		}
		next := got[len(got)-1].CreatedAt
		if !next.After(floor) {
			t.Fatalf("chunk %d did not move the floor", chunk)
		}
		floor = next
	}
	if len(seen) != seeded {
		t.Errorf("the walk read %d of the %d verdicts, and the rest fall under the watermark for good", len(seen), seeded)
	}
}
