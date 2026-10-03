package adminsvc

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/access"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/mail"
)

type fakeRequests struct {
	byEmail map[string]access.Request
}

func (f *fakeRequests) List(_ context.Context, state string) ([]access.Request, error) {
	var out []access.Request
	for _, r := range f.byEmail {
		if r.Status == state {
			out = append(out, r)
		}
	}
	return out, nil
}

func (f *fakeRequests) Get(_ context.Context, email string) (access.Request, error) {
	r, ok := f.byEmail[email]
	if !ok {
		return access.Request{}, access.ErrNotFound
	}
	return r, nil
}

func (f *fakeRequests) Decide(_ context.Context, email, state string, _ time.Time) error {
	r, ok := f.byEmail[email]
	if !ok {
		return access.ErrNotFound
	}
	r.Status = state
	f.byEmail[email] = r
	return nil
}

type fakeMailer struct {
	sent []mail.Message
	err  error
}

func (f *fakeMailer) Send(_ context.Context, m mail.Message) error {
	if f.err != nil {
		return f.err
	}
	f.sent = append(f.sent, m)
	return nil
}

func setup() (*fakeRequests, *[]string, *fakeMailer, *Server) {
	store := &fakeRequests{byEmail: map[string]access.Request{
		"bob@example.com": {Email: "bob@example.com", Status: access.Pending},
		"cat@example.com": {Email: "cat@example.com", Status: access.Pending},
	}}
	invited := &[]string{}
	invite := func(_ context.Context, email string) error {
		*invited = append(*invited, email)
		return nil
	}
	mailer := &fakeMailer{}
	return store, invited, mailer, New(store, invite, slog.New(slog.NewTextHandler(io.Discard, nil)), WithMailer(mailer))
}

var adminCtx = auth.WithAdmin(context.Background(), true)

// TestEveryCallNeedsTheAdminClaim is D-1076.
func TestEveryCallNeedsTheAdminClaim(t *testing.T) {
	store, invited, mailer, s := setup()
	ctx := auth.WithUserID(context.Background(), "u-1")
	_, err1 := s.ListAccessRequests(ctx, connect.NewRequest(&mtgv1.ListAccessRequestsRequest{}))
	_, err2 := s.ApproveAccessRequest(ctx, connect.NewRequest(&mtgv1.ApproveAccessRequestRequest{Email: "bob@example.com"}))
	_, err3 := s.DismissAccessRequest(ctx, connect.NewRequest(&mtgv1.DismissAccessRequestRequest{Email: "bob@example.com"}))
	for i, err := range []error{err1, err2, err3} {
		if connect.CodeOf(err) != connect.CodePermissionDenied {
			t.Errorf("call %d: code %v, want PermissionDenied", i+1, connect.CodeOf(err))
		}
	}
	if len(*invited) != 0 || len(mailer.sent) != 0 || store.byEmail["bob@example.com"].Status != access.Pending {
		t.Error("a refused call changed state")
	}
}

// TestApproveInvitesMarksAndMails is D-1076 and D-1077.
func TestApproveInvitesMarksAndMails(t *testing.T) {
	store, invited, mailer, s := setup()
	res, err := s.ApproveAccessRequest(adminCtx, connect.NewRequest(&mtgv1.ApproveAccessRequestRequest{Email: "bob@example.com"}))
	if err != nil || !res.Msg.GetEmailSent() {
		t.Fatalf("approve: %v, %v", res, err)
	}
	if len(*invited) != 1 || (*invited)[0] != "bob@example.com" {
		t.Errorf("invited = %v", *invited)
	}
	if store.byEmail["bob@example.com"].Status != access.Approved {
		t.Errorf("status = %q", store.byEmail["bob@example.com"].Status)
	}
	if len(mailer.sent) != 1 || mailer.sent[0].To != "bob@example.com" || !strings.Contains(mailer.sent[0].Text, SignInURL) {
		t.Errorf("sent = %+v", mailer.sent)
	}
}

func TestApproveKeepsTheApprovalWhenTheMailFails(t *testing.T) {
	store, invited, mailer, s := setup()
	mailer.err = errors.New("down")
	res, err := s.ApproveAccessRequest(adminCtx, connect.NewRequest(&mtgv1.ApproveAccessRequestRequest{Email: "bob@example.com"}))
	if err != nil || res.Msg.GetEmailSent() {
		t.Fatalf("approve: %v, %v", res, err)
	}
	if len(*invited) != 1 || store.byEmail["bob@example.com"].Status != access.Approved {
		t.Error("the approval did not hold")
	}
}

func TestApproveOfNoRequestIsNotFound(t *testing.T) {
	_, invited, _, s := setup()
	_, err := s.ApproveAccessRequest(adminCtx, connect.NewRequest(&mtgv1.ApproveAccessRequestRequest{Email: "nobody@example.com"}))
	if connect.CodeOf(err) != connect.CodeNotFound || len(*invited) != 0 {
		t.Errorf("code %v, invited %v", connect.CodeOf(err), *invited)
	}
}

func TestDismissAndList(t *testing.T) {
	store, invited, mailer, s := setup()
	if _, err := s.DismissAccessRequest(adminCtx, connect.NewRequest(&mtgv1.DismissAccessRequestRequest{Email: "cat@example.com"})); err != nil {
		t.Fatal(err)
	}
	if store.byEmail["cat@example.com"].Status != access.Dismissed || len(*invited) != 0 || len(mailer.sent) != 0 {
		t.Error("dismiss did more than mark the request")
	}
	res, err := s.ListAccessRequests(adminCtx, connect.NewRequest(&mtgv1.ListAccessRequestsRequest{}))
	if err != nil {
		t.Fatal(err)
	}
	if got := res.Msg.GetRequests(); len(got) != 1 || got[0].GetEmail() != "bob@example.com" || got[0].GetStatus() != mtgv1.AccessStatus_ACCESS_STATUS_PENDING {
		t.Errorf("pending = %v", got)
	}
}
