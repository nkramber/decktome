package sessions

import (
	"errors"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

// The store talks to Firestore, and no unit test can reach that. Before
// 2026-08-25 every Firestore path here was at zero coverage: NewID, Get,
// GetState, and the document paths never ran, and Put ran only far enough
// to reject a session with no id. A wrong collection path, a bad
// `firestore` struct tag, or a misused transaction would have passed
// every green test in the repo.
//
// These tests need the local emulator, so CI skips them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/sessions -count=1

func isNotFound(err error) bool { return errors.Is(err, ErrNotFound) }

// emulatorRepo opens a repo against the emulator, or skips.
func emulatorRepo(t *testing.T) (*Repo, func()) {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the store against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	return NewRepo(client), func() { _ = client.Close() }
}

func sampleState() questions.Snapshot {
	st := questions.NewState(true)
	st.Slots.Theme = "lifegain"
	st.Close("theme")
	st.MarkAsked("commander", "commander")
	st.AddLocked("Sanguine Bond")
	st.SetCommander("Karlov of the Ghost Council")
	st.Ctx.Words = "build me a lifegain deck"
	return st.Snapshot()
}

func TestEmulatorPutGetRoundTrip(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()

	id := repo.NewID("u1")
	if id == "" {
		t.Fatal("NewID gave an empty id")
	}
	want := sampleSession()
	want.Id = id
	snap := sampleState()

	if err := repo.Put(ctx, "u1", want, snap); err != nil {
		t.Fatalf("put: %v", err)
	}

	// The public half.
	got, err := repo.Get(ctx, "u1", id)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if !proto.Equal(want, got) {
		t.Errorf("session did not survive Firestore:\n got %v\nwant %v", got, want)
	}

	// The private half, and the public half together.
	gotSession, gotSnap, err := repo.GetState(ctx, "u1", id)
	if err != nil {
		t.Fatalf("get state: %v", err)
	}
	if !proto.Equal(want, gotSession) {
		t.Error("GetState returned a different session from Get")
	}
	if gotSnap.Version != questions.SnapshotVersion {
		t.Errorf("snapshot version = %d, want %d", gotSnap.Version, questions.SnapshotVersion)
	}
	back := questions.Restore(id, gotSession.GetSlots(), gotSnap)
	if !back.Ctx.Asked["commander"] {
		t.Error("the asked rows did not survive, so a resumed session repeats a question")
	}
	if back.Ctx.Words != "build me a lifegain deck" {
		t.Errorf("the words did not survive: %q", back.Ctx.Words)
	}
	if len(back.CommanderNames) != 1 || len(back.LockedNames) != 1 {
		t.Errorf("the card lists did not survive: %+v", back)
	}
}

// TestEmulatorNotFound proves the error maps to something the service can
// turn into a 404 rather than a 500.
func TestEmulatorNotFound(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	if _, err := repo.Get(t.Context(), "u1", "no-such-session"); err == nil {
		t.Fatal("reading a missing session gave no error")
	} else if !isNotFound(err) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
	if _, _, err := repo.GetState(t.Context(), "u1", "no-such-session"); !isNotFound(err) {
		t.Errorf("GetState err = %v, want ErrNotFound", err)
	}
}

// TestEmulatorUpdateKeepsOneDocument proves a second turn overwrites the
// session in place instead of making a new one.
func TestEmulatorUpdateKeepsOneDocument(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()
	id := repo.NewID("u2")

	first := sampleSession()
	first.Id = id
	if err := repo.Put(ctx, "u2", first, sampleState()); err != nil {
		t.Fatalf("put 1: %v", err)
	}
	second := sampleSession()
	second.Id = id
	second.Status = mtgv1.SessionStatus_SESSION_STATUS_READY
	second.Turns = append(second.Turns, &mtgv1.Turn{UserMessage: "bracket 3"})
	second.UpdatedAt = timestamppb.Now()
	if err := repo.Put(ctx, "u2", second, sampleState()); err != nil {
		t.Fatalf("put 2: %v", err)
	}
	got, err := repo.Get(ctx, "u2", id)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.GetStatus() != mtgv1.SessionStatus_SESSION_STATUS_READY || len(got.GetTurns()) != 2 {
		t.Errorf("the update did not land: status %v, %d turns", got.GetStatus(), len(got.GetTurns()))
	}
}

// TestEmulatorUsersAreSeparate keeps one user out of another's session.
func TestEmulatorUsersAreSeparate(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()
	id := repo.NewID("owner")
	s := sampleSession()
	s.Id = id
	if err := repo.Put(ctx, "owner", s, sampleState()); err != nil {
		t.Fatalf("put: %v", err)
	}
	if _, err := repo.Get(ctx, "someone-else", id); !isNotFound(err) {
		t.Errorf("another user read the session: %v", err)
	}
}

// TestEmulatorStatePartlyMissing covers a session stored before the
// private document existed. It must open, not fail.
func TestEmulatorStatePartlyMissing(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()
	id := repo.NewID("u3")
	s := sampleSession()
	s.Id = id
	if err := repo.Put(ctx, "u3", s, sampleState()); err != nil {
		t.Fatalf("put: %v", err)
	}
	// Remove the private half, as an old session would have it.
	if _, err := repo.stateDoc("u3", id).Delete(ctx); err != nil {
		t.Fatalf("delete state: %v", err)
	}
	got, snap, err := repo.GetState(ctx, "u3", id)
	if err != nil {
		t.Fatalf("a session with no private state failed to open: %v", err)
	}
	if got.GetId() != id {
		t.Errorf("session id = %q", got.GetId())
	}
	if snap.Version != 0 {
		t.Errorf("snapshot version = %d, want the zero value", snap.Version)
	}
}

// TestEmulatorPrivateStateIsNotPublic is the D-74 layout invariant, and
// a round trip can not prove it. Put and GetState agree with each other
// whatever path they use, so a test that only writes and reads back
// passes even when the private state sits inside the public document.
// GetSession returns that document to the client.
func TestEmulatorPrivateStateIsNotPublic(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()
	id := repo.NewID("u4")
	s := sampleSession()
	s.Id = id
	if err := repo.Put(ctx, "u4", s, sampleState()); err != nil {
		t.Fatalf("put: %v", err)
	}
	snap, err := repo.doc("u4", id).Get(ctx)
	if err != nil {
		t.Fatalf("read the public document: %v", err)
	}
	for field := range snap.Data() {
		if field == "state_gz" {
			t.Error("the private state sits in the public document, which GetSession returns")
		}
	}
	// And the private document really holds it.
	priv, err := repo.stateDoc("u4", id).Get(ctx)
	if err != nil {
		t.Fatalf("read the private document: %v", err)
	}
	if _, ok := priv.Data()["state_gz"]; !ok {
		t.Errorf("the private document holds no state: %v", priv.Data())
	}
}

// TestEmulatorStateSurvivesAPathChange is the mutation guard. It reads
// the private document by its literal path, so a write that lands
// somewhere else fails here instead of passing a round trip.
func TestEmulatorStateSurvivesAPathChange(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()
	id := repo.NewID("u5")
	s := sampleSession()
	s.Id = id
	if err := repo.Put(ctx, "u5", s, sampleState()); err != nil {
		t.Fatalf("put: %v", err)
	}
	literal := repo.client.Collection("users").Doc("u5").
		Collection("sessions").Doc(id).
		Collection("private").Doc("state")
	got, err := literal.Get(ctx)
	if err != nil {
		t.Fatalf("the private state is not at users/{uid}/sessions/{id}/private/state: %v", err)
	}
	var stored storedState
	if err := got.DataTo(&stored); err != nil {
		t.Fatalf("private document: %v", err)
	}
	var out questions.Snapshot
	if err := ungzJSON(stored.StateGz, &out); err != nil {
		t.Fatalf("private payload: %v", err)
	}
	if out.Version != questions.SnapshotVersion {
		t.Errorf("snapshot version = %d", out.Version)
	}
}
