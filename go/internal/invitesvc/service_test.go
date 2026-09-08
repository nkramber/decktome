package invitesvc

import (
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
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
