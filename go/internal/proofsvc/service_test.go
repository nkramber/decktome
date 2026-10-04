package proofsvc

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
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/mail"
	"github.com/nkramber/decktome/go/internal/prooflink"
)

type fakeAccounts struct {
	accts   map[string]Account
	proved  []string
	err     error
	noToken bool
}

func (f *fakeAccounts) Account(_ context.Context, uid string) (Account, error) {
	if f.err != nil {
		return Account{}, f.err
	}
	a, ok := f.accts[uid]
	if !ok {
		return Account{}, ErrNoAccount
	}
	return a, nil
}

func (f *fakeAccounts) Prove(_ context.Context, uid string) error {
	f.proved = append(f.proved, uid)
	a := f.accts[uid]
	a.Proved = true
	f.accts[uid] = a
	return nil
}

func (f *fakeAccounts) Token(_ context.Context, uid string) (string, error) {
	if f.noToken {
		return "", errors.New("iam.serviceAccounts.signBlob denied")
	}
	return "token-" + uid, nil
}

// fakeLinks keeps the codes in memory, with the cooldown of the store.
type fakeLinks struct {
	byCode map[string]prooflink.Link
}

func (f *fakeLinks) Put(_ context.Context, code string, link prooflink.Link) error {
	for c, l := range f.byCode {
		if l.UID != link.UID {
			continue
		}
		if link.CreatedAt.Sub(l.CreatedAt) < prooflink.Cooldown {
			return prooflink.ErrTooSoon
		}
		delete(f.byCode, c)
	}
	f.byCode[code] = link
	return nil
}

// Take follows the store: a used link stays until it expires, and a
// second take names the uid alone (D-1119).
func (f *fakeLinks) Take(_ context.Context, code string, now time.Time) (prooflink.Link, error) {
	l, ok := f.byCode[code]
	if !ok {
		return prooflink.Link{}, prooflink.ErrUnknown
	}
	expired := !now.Before(l.ExpiresAt)
	if expired {
		delete(f.byCode, code)
	}
	switch {
	case !l.UsedAt.IsZero():
		return prooflink.Link{UID: l.UID, UsedAt: l.UsedAt}, prooflink.ErrUsed
	case expired:
		return prooflink.Link{}, prooflink.ErrExpired
	}
	f.byCode[code] = prooflink.Link{UID: l.UID, CreatedAt: l.CreatedAt, ExpiresAt: l.ExpiresAt, UsedAt: now}
	return l, nil
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

type fakeList map[string]bool

func (f fakeList) Allowed(_ context.Context, email string) (bool, error) { return f[email], nil }

type fakeClosed map[string]bool

func (f fakeClosed) Closed(_ context.Context, uid string) (bool, error) { return f[uid], nil }

type rig struct {
	accounts *fakeAccounts
	links    *fakeLinks
	mailer   *fakeMailer
	now      time.Time
	server   *Server
}

func newRig(opts ...Option) *rig {
	r := &rig{
		accounts: &fakeAccounts{accts: map[string]Account{
			"u-ann": {Email: "ann@example.com"},
			"u-bob": {Email: "bob@example.com", Proved: true},
			"u-eve": {Email: "eve@example.com"},
		}},
		links:  &fakeLinks{byCode: map[string]prooflink.Link{}},
		mailer: &fakeMailer{},
		now:    time.Date(2026, 10, 3, 12, 0, 0, 0, time.UTC),
	}
	base := []Option{
		WithMailer(r.mailer),
		WithClock(func() time.Time { return r.now }),
		WithAllowlist(fakeList{"ann@example.com": true, "bob@example.com": true}),
	}
	r.server = New(r.accounts, r.links, slog.New(slog.NewTextHandler(io.Discard, nil)), append(base, opts...)...)
	return r
}

func (r *rig) send(uid string) (*mtgv1.SendLinkResponse, error) {
	res, err := r.server.SendLink(auth.WithUserID(context.Background(), uid), connect.NewRequest(&mtgv1.SendLinkRequest{}))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

func (r *rig) open(code string) (string, error) {
	res, err := r.openFull(code)
	if err != nil {
		return "", err
	}
	return res.GetCustomToken(), nil
}

func (r *rig) openFull(code string) (*mtgv1.OpenLinkResponse, error) {
	res, err := r.server.OpenLink(context.Background(), connect.NewRequest(&mtgv1.OpenLinkRequest{Code: code}))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// lastCode reads the code of the last email.
func (r *rig) lastCode(t *testing.T) string {
	t.Helper()
	if len(r.mailer.sent) == 0 {
		t.Fatal("no email sent")
	}
	text := r.mailer.sent[len(r.mailer.sent)-1].Text
	at := strings.Index(text, LinkBase)
	if at < 0 {
		t.Fatalf("the email holds no link: %q", text)
	}
	return text[at+len(LinkBase) : at+len(LinkBase)+prooflink.CodeLen]
}

func codeOf(err error) connect.Code {
	var ce *connect.Error
	if errors.As(err, &ce) {
		return ce.Code()
	}
	return 0
}

// TestSendThenOpen is D-1081 and D-1082: the email carries the text of
// the owner and a short link, and the link proves the email and signs
// in one time.
func TestSendThenOpen(t *testing.T) {
	r := newRig()
	if _, err := r.send("u-ann"); err != nil {
		t.Fatal(err)
	}
	m := r.mailer.sent[0]
	code := r.lastCode(t)
	want := "Hello,\n\nFollow this link to verify your email address.\n\n" + LinkBase + code +
		"\n\nIf you didn't ask to verify this address, you can ignore this email.\n\nThanks,\n\nThe Deck Tome team"
	if m.To != "ann@example.com" || m.Text != want {
		t.Fatalf("email = %+v", m)
	}
	if n := len(LinkBase + code); n > 40 {
		t.Errorf("the link has %d characters", n)
	}
	tok, err := r.open(code)
	if err != nil || tok != "token-u-ann" {
		t.Fatalf("open: %q, %v", tok, err)
	}
	if len(r.accounts.proved) != 1 || !r.accounts.accts["u-ann"].Proved {
		t.Fatalf("proved = %v", r.accounts.proved)
	}
	again, err := r.openFull(code)
	if err != nil || !again.GetAlreadyProved() || again.GetCustomToken() != "" {
		t.Fatalf("a second open: %v, %v, want already_proved and no token", again, err)
	}
	if len(r.accounts.proved) != 1 {
		t.Fatalf("a second open proved again: %v", r.accounts.proved)
	}
}

// TestUsedLink is D-1119: a used link signs in no one. It says that the
// email is proved while the account is proved, and it works as an
// unknown code otherwise and after its expiry.
func TestUsedLink(t *testing.T) {
	used := func(r *rig, uid string) string {
		code, err := prooflink.NewCode()
		if err != nil {
			t.Fatal(err)
		}
		r.links.byCode[code] = prooflink.Link{UID: uid, CreatedAt: r.now, ExpiresAt: r.now.Add(prooflink.Life), UsedAt: r.now}
		return code
	}
	t.Run("a proved account", func(t *testing.T) {
		r := newRig()
		res, err := r.openFull(used(r, "u-bob"))
		if err != nil || !res.GetAlreadyProved() || res.GetCustomToken() != "" {
			t.Fatalf("open: %v, %v", res, err)
		}
	})
	t.Run("an account that is not proved", func(t *testing.T) {
		r := newRig()
		if _, err := r.open(used(r, "u-ann")); codeOf(err) != connect.CodeNotFound {
			t.Fatalf("err = %v, want NotFound", err)
		}
		if len(r.accounts.proved) != 0 {
			t.Fatalf("a used link proved %v", r.accounts.proved)
		}
	})
	t.Run("an account that is gone", func(t *testing.T) {
		r := newRig()
		if _, err := r.open(used(r, "u-gone")); codeOf(err) != connect.CodeNotFound {
			t.Fatalf("err = %v, want NotFound", err)
		}
	})
	t.Run("after the expiry", func(t *testing.T) {
		r := newRig()
		code := used(r, "u-bob")
		r.now = r.now.Add(prooflink.Life)
		if res, err := r.openFull(code); err != nil || !res.GetAlreadyProved() {
			t.Fatalf("the last open: %v, %v", res, err)
		}
		if _, err := r.open(code); codeOf(err) != connect.CodeNotFound {
			t.Fatalf("err = %v, want NotFound", err)
		}
	})
}

func TestSendRefusals(t *testing.T) {
	r := newRig()
	if res, err := r.send("u-bob"); err != nil || !res.GetAlreadyProved() || len(r.mailer.sent) != 0 {
		t.Fatalf("a proved account: %v, %v, sent %d", res, err, len(r.mailer.sent))
	}
	if _, err := r.send(""); codeOf(err) != connect.CodeUnauthenticated {
		t.Fatalf("no user: %v", err)
	}
	if _, err := r.send("u-ann"); err != nil {
		t.Fatal(err)
	}
	r.now = r.now.Add(30 * time.Second)
	if _, err := r.send("u-ann"); codeOf(err) != connect.CodeResourceExhausted {
		t.Fatalf("a send within the cooldown: %v", err)
	}
	first := r.lastCode(t)
	r.now = r.now.Add(prooflink.Cooldown)
	if _, err := r.send("u-ann"); err != nil {
		t.Fatal(err)
	}
	if _, err := r.open(first); codeOf(err) != connect.CodeNotFound {
		t.Fatalf("the replaced link: %v, want NotFound", err)
	}

	r.mailer.err = errors.New("status 500")
	r.now = r.now.Add(prooflink.Cooldown)
	if _, err := r.send("u-ann"); codeOf(err) != connect.CodeUnavailable {
		t.Fatalf("a failed send: %v", err)
	}

	none := New(r.accounts, r.links, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if _, err := none.SendLink(auth.WithUserID(context.Background(), "u-ann"), connect.NewRequest(&mtgv1.SendLinkRequest{})); codeOf(err) != connect.CodeUnimplemented {
		t.Fatalf("no mailer: %v, want Unimplemented", err)
	}
}

func TestOpenRefusals(t *testing.T) {
	put := func(r *rig, uid, email string) string {
		code, err := prooflink.NewCode()
		if err != nil {
			t.Fatal(err)
		}
		r.links.byCode[code] = prooflink.Link{UID: uid, Email: email, CreatedAt: r.now, ExpiresAt: r.now.Add(prooflink.Life)}
		return code
	}
	for _, tc := range []struct {
		name  string
		setup func(r *rig) string
		code  connect.Code
	}{
		{"a malformed code", func(*rig) string { return "abc/../def" }, connect.CodeNotFound},
		{"an unknown code", func(*rig) string { return "Ab3dEf9kQ2" }, connect.CodeNotFound},
		{"an expired link", func(r *rig) string {
			c := put(r, "u-ann", "ann@example.com")
			r.now = r.now.Add(prooflink.Life)
			return c
		}, connect.CodeFailedPrecondition},
		{"an account that is gone", func(r *rig) string { return put(r, "u-gone", "gone@example.com") }, connect.CodeNotFound},
		{"a changed email", func(r *rig) string { return put(r, "u-ann", "old@example.com") }, connect.CodePermissionDenied},
		{"an email off the invite list", func(r *rig) string { return put(r, "u-eve", "eve@example.com") }, connect.CodePermissionDenied},
		{"a disabled account", func(r *rig) string {
			r.accounts.accts["u-ann"] = Account{Email: "ann@example.com", Disabled: true}
			return put(r, "u-ann", "ann@example.com")
		}, connect.CodePermissionDenied},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := newRig()
			if _, err := r.open(tc.setup(r)); codeOf(err) != tc.code {
				t.Fatalf("err = %v, want %v", err, tc.code)
			}
			if len(r.accounts.proved) != 0 {
				t.Fatalf("a refusal proved %v", r.accounts.proved)
			}
		})
	}
	t.Run("a closed account", func(t *testing.T) {
		r := newRig(WithClosed(fakeClosed{"u-ann": true}))
		if _, err := r.open(put(r, "u-ann", "ann@example.com")); codeOf(err) != connect.CodePermissionDenied {
			t.Fatalf("err = %v", err)
		}
	})
	t.Run("a token that can not be signed still proves the email", func(t *testing.T) {
		r := newRig()
		r.accounts.noToken = true
		tok, err := r.open(put(r, "u-ann", "ann@example.com"))
		if err != nil || tok != "" || !r.accounts.accts["u-ann"].Proved {
			t.Fatalf("open: %q, %v, proved %v", tok, err, r.accounts.proved)
		}
	})
	t.Run("a proved account signs in with no new proof", func(t *testing.T) {
		r := newRig()
		tok, err := r.open(put(r, "u-bob", "BOB@example.com"))
		if err != nil || tok != "token-u-bob" || len(r.accounts.proved) != 0 {
			t.Fatalf("open: %q, %v, proved %v", tok, err, r.accounts.proved)
		}
	})
}
