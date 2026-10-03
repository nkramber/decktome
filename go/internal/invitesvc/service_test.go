package invitesvc

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/notify"
)

type fakeList struct {
	emails map[string]bool
	err    error
}

func (f fakeList) Allowed(_ context.Context, email string) (bool, error) {
	if f.err != nil {
		return false, f.err
	}
	return f.emails[email], nil
}

func check(t *testing.T, s *Server, email string) (*mtgv1.CheckInviteResponse, error) {
	t.Helper()
	res, err := s.CheckInvite(context.Background(), connect.NewRequest(&mtgv1.CheckInviteRequest{Email: email}))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// TestCheckInvite is D-592. The create-account form asks the list before
// it makes an account, so an email off the list never becomes one.
func TestCheckInvite(t *testing.T) {
	list := fakeList{emails: map[string]bool{"ann@example.com": true}}

	t.Run("an email on the list is allowed", func(t *testing.T) {
		msg, err := check(t, New(list), "ann@example.com")
		if err != nil || !msg.GetAllowed() {
			t.Errorf("allowed = %v, err = %v, want true", msg.GetAllowed(), err)
		}
	})

	t.Run("an email off the list is not", func(t *testing.T) {
		msg, err := check(t, New(list), "bob@example.com")
		if err != nil || msg.GetAllowed() {
			t.Errorf("allowed = %v, err = %v, want false", msg.GetAllowed(), err)
		}
	})

	t.Run("an empty email is not", func(t *testing.T) {
		msg, err := check(t, New(list), "")
		if err != nil || msg.GetAllowed() {
			t.Errorf("allowed = %v, err = %v, want false", msg.GetAllowed(), err)
		}
	})

	// Local mode wires no list, and every emulator user gets in
	// (guardrail 9). A nil interface and a nil pointer both read as none.
	t.Run("a server with no list allows every email", func(t *testing.T) {
		msg, err := check(t, New(nil), "anyone@example.com")
		if err != nil || !msg.GetAllowed() {
			t.Errorf("allowed = %v, err = %v, want true", msg.GetAllowed(), err)
		}
	})

	// A read failure must never answer false. A false sends a person away
	// who is on the list.
	t.Run("a list that can not be read answers Unavailable", func(t *testing.T) {
		_, err := check(t, New(fakeList{err: errors.New("firestore down")}), "ann@example.com")
		if connect.CodeOf(err) != connect.CodeUnavailable {
			t.Errorf("code = %v, want Unavailable", connect.CodeOf(err))
		}
	})
}

type fakeRequests struct {
	puts map[string]string
}

func (f *fakeRequests) Put(_ context.Context, email, note string, _ time.Time) (bool, error) {
	_, seen := f.puts[email]
	f.puts[email] = note
	return !seen, nil
}

type fakeNotifier struct{ notices []notify.Notice }

func (f *fakeNotifier) Notify(n notify.Notice) { f.notices = append(f.notices, n) }

func request(s *Server, email, note string) (*mtgv1.RequestAccessResponse, error) {
	res, err := s.RequestAccess(context.Background(), connect.NewRequest(&mtgv1.RequestAccessRequest{Email: email, Note: note}))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// TestRequestAccess is D-1075: a new request is stored in its clean form
// and pings the owner, a repeat pings no one, an invited email stores
// nothing, and a bad email is refused.
func TestRequestAccess(t *testing.T) {
	store := &fakeRequests{puts: map[string]string{}}
	pings := &fakeNotifier{}
	s := New(fakeList{emails: map[string]bool{"ann@example.com": true}}, WithRequests(store), WithNotifier(pings))

	msg, err := request(s, " Bob@Example.com ", " elves ")
	if err != nil || msg.GetAlreadyInvited() {
		t.Fatalf("new request: %v, %v", msg, err)
	}
	if store.puts["bob@example.com"] != "elves" {
		t.Errorf("stored = %v", store.puts)
	}
	if len(pings.notices) != 1 || !strings.Contains(pings.notices[0].Message, "bob@example.com") {
		t.Errorf("notices = %v", pings.notices)
	}
	if _, err := request(s, "bob@example.com", ""); err != nil {
		t.Fatal(err)
	}
	if len(pings.notices) != 1 {
		t.Errorf("a repeat pinged the owner: %d notices", len(pings.notices))
	}

	msg, err = request(s, "Ann@example.com", "")
	if err != nil || !msg.GetAlreadyInvited() {
		t.Fatalf("invited email: %v, %v", msg, err)
	}
	if _, stored := store.puts["ann@example.com"]; stored {
		t.Error("an invited email was stored")
	}

	if _, err := request(s, "not an email", ""); connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("bad email: code %v, want InvalidArgument", connect.CodeOf(err))
	}
}

func TestRequestAccessCapsTheNotices(t *testing.T) {
	store := &fakeRequests{puts: map[string]string{}}
	pings := &fakeNotifier{}
	now := time.Date(2026, 10, 2, 0, 0, 0, 0, time.UTC)
	s := New(nil, WithRequests(store), WithNotifier(pings), WithClock(func() time.Time { return now }))
	for i := range NoticesPerHour + 5 {
		if _, err := request(s, fmt.Sprintf("p%d@example.com", i), ""); err != nil {
			t.Fatal(err)
		}
	}
	if len(store.puts) != NoticesPerHour+5 || len(pings.notices) != NoticesPerHour {
		t.Errorf("stored %d, notices %d, want %d and %d", len(store.puts), len(pings.notices), NoticesPerHour+5, NoticesPerHour)
	}
}

func TestRequestAccessWithNoStore(t *testing.T) {
	if _, err := request(New(nil), "bob@example.com", ""); connect.CodeOf(err) != connect.CodeUnimplemented {
		t.Errorf("code %v, want Unimplemented", connect.CodeOf(err))
	}
}
