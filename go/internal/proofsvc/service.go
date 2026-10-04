// Package proofsvc sends the email that proves an address, and it opens
// the link of that email (D-1081, D-1082).
//
// The Firebase template can not carry the text of the owner, and its
// link is long. So the API sends the email through Resend, with a short
// link to the web app. The link proves the email and signs in the
// browser that opens it, one time, for 3 days.
package proofsvc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"connectrpc.com/connect"
	fbauth "firebase.google.com/go/v4/auth"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/allowlist"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/mail"
	"github.com/nkramber/decktome/go/internal/prooflink"
)

// LinkBase is the start of each link. The web app opens the path /v/
// and sends the code to OpenLink.
const LinkBase = "https://decktome.com/v/"

// Account is the state of one Firebase account.
type Account struct {
	Email    string
	Proved   bool
	Disabled bool
}

// ErrNoAccount is a uid that names no account.
var ErrNoAccount = errors.New("proofsvc: no account")

// Accounts reads and changes the Firebase accounts.
type Accounts interface {
	Account(ctx context.Context, uid string) (Account, error)
	// Prove marks the email of the account proved.
	Prove(ctx context.Context, uid string) error
	// Token makes a custom token that signs in to the account.
	Token(ctx context.Context, uid string) (string, error)
}

// Links is the store of the links.
type Links interface {
	Put(ctx context.Context, code string, link prooflink.Link) error
	Take(ctx context.Context, code string, now time.Time) (prooflink.Link, error)
}

// Allowlist is the invite list (D-420). A custom token starts no
// blocking function, so the open reads the list itself (D-1082).
type Allowlist interface {
	Allowed(ctx context.Context, email string) (bool, error)
}

// Closed reads whether the owner closed an account (D-941).
type Closed interface {
	Closed(ctx context.Context, uid string) (bool, error)
}

// Server answers SendLink and OpenLink. A server with no mailer answers
// SendLink with Unimplemented, and the web app then sends the email of
// Firebase. That is local mode, and a deploy with no Resend secret.
type Server struct {
	accounts Accounts
	links    Links
	mailer   mail.Sender
	list     Allowlist
	closed   Closed
	logger   *slog.Logger
	now      func() time.Time
}

// Option tunes the server.
type Option func(*Server)

// WithMailer sends the email.
func WithMailer(m mail.Sender) Option { return func(s *Server) { s.mailer = m } }

// WithAllowlist reads the invite list before the open signs in.
func WithAllowlist(l Allowlist) Option { return func(s *Server) { s.list = l } }

// WithClosed refuses the open of a closed account.
func WithClosed(c Closed) Option { return func(s *Server) { s.closed = c } }

// WithClock replaces time.Now (tests).
func WithClock(now func() time.Time) Option { return func(s *Server) { s.now = now } }

// New builds the server.
func New(accounts Accounts, links Links, logger *slog.Logger, opts ...Option) *Server {
	s := &Server{accounts: accounts, links: links, logger: logger, now: time.Now}
	for _, o := range opts {
		o(s)
	}
	return s
}

// The sentences of a refusal. The web app shows each one as it is.
var (
	errNoMailer = errors.New("this server sends no email")
	errTooSoon  = errors.New("wait one minute, then send the link again")
	errNoSend   = errors.New("the email could not be sent, try again soon")
	errUnknown  = errors.New("this link does not work. It worked before, or a newer link replaced it. Sign in to send a new link")
	errExpired  = errors.New("this link expired. Sign in to send a new link")
	errRefused  = errors.New("this account can not sign in")
	errListRead = errors.New("the invite list could not be read")
)

// SendLink sends the email that proves the address of the caller. An
// account that proved its email gets no email.
func (s *Server) SendLink(ctx context.Context, _ *connect.Request[mtgv1.SendLinkRequest]) (*connect.Response[mtgv1.SendLinkResponse], error) {
	uid := auth.UserID(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("sign in first"))
	}
	if s.mailer == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoMailer)
	}
	acct, err := s.accounts.Account(ctx, uid)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if acct.Proved {
		return connect.NewResponse(&mtgv1.SendLinkResponse{AlreadyProved: true}), nil
	}
	code, err := prooflink.NewCode()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	now := s.now()
	err = s.links.Put(ctx, code, prooflink.Link{UID: uid, Email: acct.Email, CreatedAt: now, ExpiresAt: now.Add(prooflink.Life)})
	if errors.Is(err, prooflink.ErrTooSoon) {
		return nil, connect.NewError(connect.CodeResourceExhausted, errTooSoon)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if err := s.mailer.Send(ctx, Message(acct.Email, LinkBase+code)); err != nil {
		s.logger.Error("the proof email failed", "uid", uid, "err", err)
		return nil, connect.NewError(connect.CodeUnavailable, errNoSend)
	}
	return connect.NewResponse(&mtgv1.SendLinkResponse{}), nil
}

// OpenLink takes the link of one code. It proves the email of the
// account, and it answers a custom token that signs in to the account.
// The link works one time, also when a later step refuses. A token that
// can not be signed leaves the email proved, and the answer carries no
// token, so the person signs in with the password. A second open of a
// link whose account is proved answers already_proved and no token, so
// a mail scanner that opens the link first leaves the person a clear
// next step (D-1119).
func (s *Server) OpenLink(ctx context.Context, req *connect.Request[mtgv1.OpenLinkRequest]) (*connect.Response[mtgv1.OpenLinkResponse], error) {
	code := req.Msg.GetCode()
	if !prooflink.Valid(code) {
		return nil, connect.NewError(connect.CodeNotFound, errUnknown)
	}
	link, err := s.links.Take(ctx, code, s.now())
	switch {
	case errors.Is(err, prooflink.ErrUsed):
		return s.usedLink(ctx, link.UID)
	case errors.Is(err, prooflink.ErrUnknown):
		return nil, connect.NewError(connect.CodeNotFound, errUnknown)
	case errors.Is(err, prooflink.ErrExpired):
		return nil, connect.NewError(connect.CodeFailedPrecondition, errExpired)
	case err != nil:
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	acct, err := s.accounts.Account(ctx, link.UID)
	if errors.Is(err, ErrNoAccount) {
		return nil, connect.NewError(connect.CodeNotFound, errUnknown)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	// A changed email makes the link stale: it proved the old address.
	if acct.Disabled || allowlist.Normalize(acct.Email) != allowlist.Normalize(link.Email) {
		return nil, connect.NewError(connect.CodePermissionDenied, errRefused)
	}
	if s.list != nil {
		ok, err := s.list.Allowed(ctx, acct.Email)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, errListRead)
		}
		if !ok {
			return nil, connect.NewError(connect.CodePermissionDenied, errRefused)
		}
	}
	if s.closed != nil {
		closed, err := s.closed.Closed(ctx, link.UID)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, errors.New("the account state could not be read"))
		}
		if closed {
			return nil, connect.NewError(connect.CodePermissionDenied, errRefused)
		}
	}
	if !acct.Proved {
		if err := s.accounts.Prove(ctx, link.UID); err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}
	token, err := s.accounts.Token(ctx, link.UID)
	if err != nil {
		s.logger.Error("the custom token failed", "uid", link.UID, "err", err)
		token = ""
	}
	return connect.NewResponse(&mtgv1.OpenLinkResponse{CustomToken: token}), nil
}

// usedLink answers the second open of a link (D-1119). A proved account
// gets already_proved and no token. Any other state gets the answer of
// an unknown code, as before.
func (s *Server) usedLink(ctx context.Context, uid string) (*connect.Response[mtgv1.OpenLinkResponse], error) {
	acct, err := s.accounts.Account(ctx, uid)
	if errors.Is(err, ErrNoAccount) {
		return nil, connect.NewError(connect.CodeNotFound, errUnknown)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if !acct.Proved {
		return nil, connect.NewError(connect.CodeNotFound, errUnknown)
	}
	s.logger.Info("a used proof link opened again", "uid", uid)
	return connect.NewResponse(&mtgv1.OpenLinkResponse{AlreadyProved: true}), nil
}

// Message writes the email of one link with the text of the owner
// (D-1081).
func Message(to, link string) mail.Message {
	return mail.Message{
		To:      to,
		Subject: "Verify your email for Deck Tome",
		Text: "Hello,\n\n" +
			"Follow this link to verify your email address.\n\n" +
			link + "\n\n" +
			"If you didn't ask to verify this address, you can ignore this email.\n\n" +
			"Thanks,\n\n" +
			"The Deck Tome team",
	}
}

// Firebase reads and changes the accounts through the Admin SDK. A
// custom token needs the permission iam.serviceAccounts.signBlob on the
// account of the API (D-1082).
type Firebase struct {
	client *fbauth.Client
}

// NewFirebase wraps one Admin SDK client.
func NewFirebase(client *fbauth.Client) *Firebase { return &Firebase{client: client} }

// Account implements Accounts.
func (f *Firebase) Account(ctx context.Context, uid string) (Account, error) {
	u, err := f.client.GetUser(ctx, uid)
	if fbauth.IsUserNotFound(err) {
		return Account{}, ErrNoAccount
	}
	if err != nil {
		return Account{}, fmt.Errorf("proofsvc: %w", err)
	}
	return Account{Email: u.Email, Proved: u.EmailVerified, Disabled: u.Disabled}, nil
}

// Prove implements Accounts.
func (f *Firebase) Prove(ctx context.Context, uid string) error {
	if _, err := f.client.UpdateUser(ctx, uid, (&fbauth.UserToUpdate{}).EmailVerified(true)); err != nil {
		return fmt.Errorf("proofsvc: %w", err)
	}
	return nil
}

// Token implements Accounts.
func (f *Firebase) Token(ctx context.Context, uid string) (string, error) {
	tok, err := f.client.CustomToken(ctx, uid)
	if err != nil {
		return "", fmt.Errorf("proofsvc: %w", err)
	}
	return tok, nil
}
