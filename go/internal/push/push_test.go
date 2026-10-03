package push

import (
	"context"
	"errors"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func TestValidID(t *testing.T) {
	for _, id := range []string{"cYk1lZqT0kq6sGmU3yX_-a", "a", strings.Repeat("A", maxIDBytes)} {
		if !ValidID(id) {
			t.Errorf("ValidID(%q) = false, want true", id)
		}
	}
	for _, id := range []string{"", "a/b", "a.b", "..", "a b", "é", strings.Repeat("A", maxIDBytes+1)} {
		if ValidID(id) {
			t.Errorf("ValidID(%q) = true, want false", id)
		}
	}
}

func TestDeckMessage(t *testing.T) {
	m := DeckMessage(&mtgv1.Deck{Id: "d 1", Name: "Karlov lifegain"})
	if m.Title != "Your deck is ready!" || m.Body != "" || m.URL != "/decks/d%201" {
		t.Errorf("first build = %+v", m)
	}
	m = DeckMessage(&mtgv1.Deck{Id: "d2", RevisedFromDeckId: "d1"})
	if m.Title != "Your revised deck is ready!" || m.Body != "" || m.URL != "/decks/d2" {
		t.Errorf("revision = %+v", m)
	}
}

type fakeStore struct {
	ids     map[string][]string
	removed []string
	err     error
}

func (f *fakeStore) IDs(_ context.Context, uid string) ([]string, error) { return f.ids[uid], f.err }

func (f *fakeStore) Remove(_ context.Context, uid, id string) error {
	f.removed = append(f.removed, uid+"/"+id)
	return nil
}

type fakeSender struct {
	calls [][]string
	msgs  []Message
	gone  []string
	err   error
}

func (f *fakeSender) Send(_ context.Context, ids []string, m Message) ([]string, error) {
	f.calls = append(f.calls, ids)
	f.msgs = append(f.msgs, m)
	return f.gone, f.err
}

// TestDeckReadyReachesEachDeviceAndPrunesTheGone is the gate of PR-26
// in a unit: each device of the user gets the push, and a device that
// Cloud Messaging calls gone leaves the store.
func TestDeckReadyReachesEachDeviceAndPrunesTheGone(t *testing.T) {
	st := &fakeStore{ids: map[string][]string{"u1": {"a", "b"}}}
	se := &fakeSender{gone: []string{"b"}}
	NewNotifier(st, se, nil).DeckReady(t.Context(), "u1", &mtgv1.Deck{Id: "d1", Name: "Burn"})
	if len(se.calls) != 1 || !slices.Equal(se.calls[0], []string{"a", "b"}) {
		t.Fatalf("sends = %v, want one send to a and b", se.calls)
	}
	if se.msgs[0].URL != "/decks/d1" {
		t.Errorf("url = %q", se.msgs[0].URL)
	}
	if !slices.Equal(st.removed, []string{"u1/b"}) {
		t.Errorf("removed = %v, want u1/b", st.removed)
	}
}

// TestDeckReadySendsNothingWithNoDevice: a user who never opted in gets
// no push, and the sender is never called.
func TestDeckReadySendsNothingWithNoDevice(t *testing.T) {
	st := &fakeStore{ids: map[string][]string{"u1": {"a"}}}
	se := &fakeSender{}
	NewNotifier(st, se, nil).DeckReady(t.Context(), "u2", &mtgv1.Deck{Id: "d1"})
	if len(se.calls) != 0 {
		t.Errorf("sends = %v, want none", se.calls)
	}
}

// TestDeckReadySurvivesFailures: a read or a send failure goes to the log
// and never panics, and a failed send still prunes the gone devices.
func TestDeckReadySurvivesFailures(t *testing.T) {
	st := &fakeStore{err: errors.New("firestore down")}
	se := &fakeSender{}
	NewNotifier(st, se, nil).DeckReady(t.Context(), "u1", &mtgv1.Deck{Id: "d1"})
	if len(se.calls) != 0 {
		t.Errorf("a failed read still sent: %v", se.calls)
	}
	st = &fakeStore{ids: map[string][]string{"u1": {"a", "b"}}}
	se = &fakeSender{gone: []string{"a"}, err: errors.New("unavailable")}
	NewNotifier(st, se, nil).DeckReady(t.Context(), "u1", &mtgv1.Deck{Id: "d1"})
	if !slices.Equal(st.removed, []string{"u1/a"}) {
		t.Errorf("removed = %v, want u1/a", st.removed)
	}
}

func TestStaleMessage(t *testing.T) {
	m := StaleMessage([]*mtgv1.Deck{{Id: "d 1", Name: "Karlov's \"lifegain\""}})
	if m.Title != `Your deck "Karlov's "lifegain"" is no longer legal.` || m.Body != "" || m.URL != "/decks/d%201" {
		t.Errorf("one deck = %+v", m)
	}
	m = StaleMessage([]*mtgv1.Deck{{Id: "d2", Name: "  "}})
	if m.Title != "One of your decks is no longer legal." || m.URL != "/decks/d2" {
		t.Errorf("no name = %+v", m)
	}
	m = StaleMessage([]*mtgv1.Deck{{Id: "d1"}, {Id: "d2"}, {Id: "d3"}})
	if m.Title != "3 of your decks are no longer legal." || m.URL != "/decks" {
		t.Errorf("three decks = %+v", m)
	}
}

// TestDecksStaleSendsOnePushForEachUser: the decks of one pass go out as
// one push to each device, a gone device leaves the store, and no deck
// sends nothing (D-1088).
func TestDecksStaleSendsOnePushForEachUser(t *testing.T) {
	st := &fakeStore{ids: map[string][]string{"u1": {"a", "b"}}}
	se := &fakeSender{gone: []string{"a"}}
	n := NewNotifier(st, se, nil)
	n.DecksStale(t.Context(), "u1", []*mtgv1.Deck{{Id: "d1"}, {Id: "d2"}})
	if len(se.calls) != 1 || !slices.Equal(se.calls[0], []string{"a", "b"}) || se.msgs[0].URL != "/decks" {
		t.Fatalf("sends = %v %+v, want one send of /decks to a and b", se.calls, se.msgs)
	}
	if !slices.Equal(st.removed, []string{"u1/a"}) {
		t.Errorf("removed = %v, want u1/a", st.removed)
	}
	n.DecksStale(t.Context(), "u1", nil)
	n.DecksStale(t.Context(), "u2", []*mtgv1.Deck{{Id: "d9"}})
	if len(se.calls) != 1 {
		t.Errorf("no deck, or no device, must send nothing: %v", se.calls)
	}
}
