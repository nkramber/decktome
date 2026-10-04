// Package auth names the caller. A Connect interceptor reads the Firebase
// ID token from the Authorization header, verifies it, and puts the user
// id in the context. The services read it with UserID.
//
// Local mode keeps the debug user: a request with no bearer token falls
// back to the fallback id when one is set. A request that carries a
// token is always verified, so a bad token never becomes the debug user.
package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
	"unicode"

	"connectrpc.com/connect"
	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"
)

// Identity is what a verified token names: the user id, and the email
// the allowlist reads (D-314). The email is empty for a token with none.
// EmailVerified says the holder proved the email, and the allowlist
// trusts a proved email alone (D-903).
//
// Admin is the custom claim admin: true, which opens the admin screen
// of the owner (D-1076).
type Identity struct {
	UID           string
	Email         string
	EmailVerified bool
	Admin         bool
}

// Verifier checks one ID token and returns the identity it names.
type Verifier interface {
	Verify(ctx context.Context, idToken string) (Identity, error)
}

// VerifierFunc adapts a function to Verifier.
type VerifierFunc func(ctx context.Context, idToken string) (Identity, error)

// Verify implements Verifier.
func (f VerifierFunc) Verify(ctx context.Context, idToken string) (Identity, error) {
	return f(ctx, idToken)
}

// Allowlist says whether an email may use the deployed app (D-314). The
// interceptor asks it for every verified token, and a fallback user of
// local mode never reaches it.
type Allowlist interface {
	Allowed(ctx context.Context, email string) (bool, error)
}

// Closed says whether the owner closed the account of a user (D-941).
// The interceptor asks it for every verified token, so a closed account
// ends at once and not when its token expires.
type Closed interface {
	Closed(ctx context.Context, uid string) (bool, error)
}

// Visits records each verified call of a user, so every user who signs
// in has a record (D-1092). It answers no error, and it never fails the
// call.
type Visits interface {
	Seen(ctx context.Context, uid, email string)
}

// ErrNotConfigured is what RejectAll answers: the server has no verifier.
var ErrNotConfigured = errors.New("token verification is not configured on this server")

// RejectAll refuses every token. A local server with no Firebase project
// uses it, so the debug fallback is the only way in.
func RejectAll() Verifier {
	return VerifierFunc(func(context.Context, string) (Identity, error) { return Identity{}, ErrNotConfigured })
}

// Firebase verifies tokens with the Firebase Admin SDK. With
// FIREBASE_AUTH_EMULATOR_HOST set, the SDK trusts the emulator's unsigned
// tokens, so the local stack needs no service account.
type Firebase struct {
	client *fbauth.Client
}

// NewFirebase builds the verifier for one project.
func NewFirebase(ctx context.Context, projectID string) (*Firebase, error) {
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID})
	if err != nil {
		return nil, fmt.Errorf("auth: firebase app: %w", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("auth: firebase auth client: %w", err)
	}
	return &Firebase{client: client}, nil
}

// Client answers the Admin SDK client, which the proof service shares
// (D-1082).
func (f *Firebase) Client() *fbauth.Client { return f.client }

// Verify implements Verifier. The email and its email_verified claim come
// from the token, and Firebase sets both for an email and password
// account.
func (f *Firebase) Verify(ctx context.Context, idToken string) (Identity, error) {
	tok, err := f.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return Identity{}, err
	}
	if tok.UID == "" {
		return Identity{}, errors.New("token carries no uid")
	}
	email, _ := tok.Claims["email"].(string)
	verified, _ := tok.Claims["email_verified"].(bool)
	admin, _ := tok.Claims[AdminClaim].(bool)
	return Identity{UID: tok.UID, Email: email, EmailVerified: verified, Admin: admin}, nil
}

type ctxKey struct{}

// UserFunc reads the caller's user id from the request context. UserID is
// the one implementation, and every service takes this type.
type UserFunc func(ctx context.Context) string

// UserID reads the user id the interceptor stored, or "" when none.
func UserID(ctx context.Context) string {
	uid, _ := ctx.Value(ctxKey{}).(string)
	return uid
}

// WithUserID returns ctx with uid set. Tests and the interceptor use it.
func WithUserID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, ctxKey{}, uid)
}

type emailKey struct{}

// Email reads the verified email the interceptor stored, or "" when none.
// The spend cap reads it for a per-user override (D-576).
func Email(ctx context.Context) string {
	email, _ := ctx.Value(emailKey{}).(string)
	return email
}

// WithEmail returns ctx with email set. Tests and the interceptor use it.
func WithEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, emailKey{}, email)
}

// AdminClaim is the custom claim that names the admin (D-1076). A make
// target sets it on the account of the owner.
const AdminClaim = "admin"

type adminKey struct{}

// IsAdmin reads whether the verified token carries the admin claim. A
// fallback user of local mode is never the admin.
func IsAdmin(ctx context.Context) bool {
	admin, _ := ctx.Value(adminKey{}).(bool)
	return admin
}

// WithAdmin returns ctx with the admin mark set. Tests and the
// interceptor use it.
func WithAdmin(ctx context.Context, admin bool) context.Context {
	return context.WithValue(ctx, adminKey{}, admin)
}

// Option tunes the interceptor.
type Option func(*interceptor)

// WithFallback names the user a request with no bearer token acts as.
// It is the DEBUG_USER_ID of local mode. Empty means no fallback, and
// then a missing token is Unauthenticated.
func WithFallback(uid string) Option {
	return func(i *interceptor) { i.fallback = uid }
}

// WithAllowlist refuses every verified user whose email is not on the
// list, with PermissionDenied and one sentence (D-314). It also refuses
// an email on the list that its holder has not proved, because anyone
// can make an account with an invited address (D-903). The deployed API
// sets it, and local mode does not, so every emulator user gets in.
func WithAllowlist(a Allowlist) Option {
	return func(i *interceptor) { i.allow = a }
}

// WithClosed refuses every verified user whose account the owner closed,
// with PermissionDenied (D-941).
func WithClosed(c Closed) Option {
	return func(i *interceptor) { i.closed = c }
}

// WithVisits records each call of a user with a proved email, after
// every check admits it (D-1092).
func WithVisits(v Visits) Option {
	return func(i *interceptor) { i.visits = v }
}

// WithCallLog writes one line for each call that the check reads: the
// procedure, the Connect code ("ok" for success), the latency, the uid,
// and the trace (D-1117). A refused call gets a line too, with the uid
// when the token named one. The email never enters a line (D-559,
// D-596, D-1093). trace reads the trace attribute from the request
// header, and nil means no trace.
func WithCallLog(log *slog.Logger, trace func(http.Header) slog.Attr) Option {
	return func(i *interceptor) {
		i.log = log
		i.trace = trace
	}
}

type interceptor struct {
	verify   Verifier
	fallback string
	allow    Allowlist
	closed   Closed
	visits   Visits
	log      *slog.Logger
	trace    func(http.Header) slog.Attr
	// public names the procedures that need no sign-in, the shared deck
	// reads of D-315. Such a call carries no user in its context.
	public map[string]bool
	// unproved names the procedures that an invited caller can call
	// before it proves the email, the send of the proof link (D-1081).
	unproved map[string]bool
}

// WithPublic names the procedures that need no sign-in (D-315). Every
// other procedure of the same service keeps the check.
func WithPublic(procedures ...string) Option {
	return func(i *interceptor) {
		if i.public == nil {
			i.public = map[string]bool{}
		}
		for _, p := range procedures {
			i.public[p] = true
		}
	}
}

// WithUnproved names the procedures that an invited caller can call
// before it proves the email (D-1081). The invite list still applies.
func WithUnproved(procedures ...string) Option {
	return func(i *interceptor) {
		if i.unproved == nil {
			i.unproved = map[string]bool{}
		}
		for _, p := range procedures {
			i.unproved[p] = true
		}
	}
}

// Interceptor returns the Connect interceptor. It covers unary and
// streaming handlers. The client side is untouched.
func Interceptor(v Verifier, opts ...Option) connect.Interceptor {
	i := &interceptor{verify: v}
	for _, o := range opts {
		o(i)
	}
	return i
}

var (
	errNoToken  = errors.New("a bearer token is required")
	errBadToken = errors.New("the bearer token was refused")
	// errNotInvited is the one sentence a user off the list reads (D-314).
	errNotInvited = errors.New("this app is open to invited users alone, and your email is not on the list")
	// errUnverified is the one sentence an invited user with an email
	// that is not proved reads (D-903).
	errUnverified = errors.New("open the link in the email that Deck Tome sent to prove your address, then sign in again")
	// errClosed is the one sentence a user of a closed account reads
	// (D-941).
	errClosed = errors.New("this account is closed")
)

// RefusalHeader names the state behind a refusal (F-59). RefusalNotInvited
// is an email off the list, and RefusalUnverified is an email on the list
// that its holder has not proved (D-903). A client reads the state from
// this header alone, and never from the sentence, so the wording of the
// message stays free to change.
const (
	RefusalHeader     = "Deck-Tome-Refusal"
	RefusalNotInvited = "not-invited"
	RefusalUnverified = "email-unverified"
	RefusalClosed     = "account-closed"
)

// notInvited is the refusal a caller off the list reads (D-314). The
// metadata carries the state, and the message carries the sentence.
func notInvited() *connect.Error {
	err := connect.NewError(connect.CodePermissionDenied, errNotInvited)
	err.Meta().Set(RefusalHeader, RefusalNotInvited)
	return err
}

// unverified is the refusal an invited caller reads before it proves the
// email (D-903).
func unverified() *connect.Error {
	err := connect.NewError(connect.CodePermissionDenied, errUnverified)
	err.Meta().Set(RefusalHeader, RefusalUnverified)
	return err
}

// resolve reads the Authorization header and returns the context that
// carries the user, or the Connect error to answer with. The uid is the
// one the token named, also on a refusal, and "" when none is known.
func (i *interceptor) resolve(ctx context.Context, procedure, authorization string) (context.Context, string, error) {
	token, present := bearer(authorization)
	if !present {
		if i.fallback == "" {
			return nil, "", connect.NewError(connect.CodeUnauthenticated, errNoToken)
		}
		return WithUserID(ctx, i.fallback), i.fallback, nil
	}
	if token == "" {
		return nil, "", connect.NewError(connect.CodeUnauthenticated, errNoToken)
	}
	id, err := i.verify.Verify(ctx, token)
	if err != nil {
		// The verifier's reason stays in the log, not on the wire: it can
		// name the project or the key id.
		return nil, "", connect.NewError(connect.CodeUnauthenticated, errBadToken)
	}
	if id.UID == "" {
		return nil, "", connect.NewError(connect.CodeUnauthenticated, errBadToken)
	}
	if i.allow != nil {
		ok, err := i.allow.Allowed(ctx, id.Email)
		if err != nil {
			return nil, id.UID, connect.NewError(connect.CodeUnavailable, errors.New("the invite list could not be read"))
		}
		if !ok {
			return nil, id.UID, notInvited()
		}
		if !id.EmailVerified && !i.unproved[procedure] {
			return nil, id.UID, unverified()
		}
	}
	if i.closed != nil {
		closed, err := i.closed.Closed(ctx, id.UID)
		if err != nil {
			return nil, id.UID, connect.NewError(connect.CodeUnavailable, errors.New("the account state could not be read"))
		}
		if closed {
			err := connect.NewError(connect.CodePermissionDenied, errClosed)
			err.Meta().Set(RefusalHeader, RefusalClosed)
			return nil, id.UID, err
		}
	}
	if i.visits != nil && id.EmailVerified {
		i.visits.Seen(ctx, id.UID, id.Email)
	}
	return WithAdmin(WithEmail(WithUserID(ctx, id.UID), id.Email), id.Admin), id.UID, nil
}

// logCall writes the line of WithCallLog (D-1117). A refusal carries
// its state from RefusalHeader. The line holds no email.
func (i *interceptor) logCall(ctx context.Context, msg, procedure, uid string, header http.Header, start time.Time, err error) {
	if i.log == nil {
		return
	}
	code := "ok"
	if err != nil {
		code = connect.CodeOf(err).String()
	}
	attrs := []slog.Attr{
		slog.String("procedure", procedure),
		slog.String("code", code),
		slog.Int64("latency_ms", time.Since(start).Milliseconds()),
		slog.String("uid", uid),
	}
	var ce *connect.Error
	if errors.As(err, &ce) {
		if state := ce.Meta().Get(RefusalHeader); state != "" {
			attrs = append(attrs, slog.String("refusal", state))
		}
	}
	if i.trace != nil {
		if a := i.trace(header); a.Key != "" {
			attrs = append(attrs, a)
		}
	}
	i.log.LogAttrs(ctx, slog.LevelInfo, msg, attrs...)
}

// bearer splits "Bearer <token>". present is false when the header is
// absent or names another scheme. A bare "Bearer" is present with an
// empty token, because the transport trims the trailing space. A scheme
// followed by a tab or another separator that is not a space is a
// malformed bearer: present with an empty token, so it is refused and
// never falls back to the debug user.
func bearer(authorization string) (token string, present bool) {
	const scheme = "bearer"
	v := strings.TrimSpace(authorization)
	if len(v) < len(scheme) || !strings.EqualFold(v[:len(scheme)], scheme) {
		return "", false
	}
	rest := v[len(scheme):]
	if rest == "" {
		return "", true
	}
	if rest[0] != ' ' {
		if unicode.IsSpace(rune(rest[0])) {
			return "", true
		}
		return "", false
	}
	return strings.TrimSpace(rest), true
}

func (i *interceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		if req.Spec().IsClient || i.public[req.Spec().Procedure] {
			return next(ctx, req)
		}
		start := time.Now()
		procedure := req.Spec().Procedure
		authed, uid, err := i.resolve(ctx, procedure, req.Header().Get("Authorization"))
		if err != nil {
			i.logCall(ctx, "rpc refused", procedure, uid, req.Header(), start, err)
			return nil, err
		}
		res, err := next(authed, req)
		i.logCall(authed, "rpc", procedure, uid, req.Header(), start, err)
		return res, err
	}
}

func (i *interceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (i *interceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		if i.public[conn.Spec().Procedure] {
			return next(ctx, conn)
		}
		start := time.Now()
		procedure := conn.Spec().Procedure
		authed, uid, err := i.resolve(ctx, procedure, conn.RequestHeader().Get("Authorization"))
		if err != nil {
			i.logCall(ctx, "rpc refused", procedure, uid, conn.RequestHeader(), start, err)
			return err
		}
		err = next(authed, conn)
		i.logCall(authed, "rpc", procedure, uid, conn.RequestHeader(), start, err)
		return err
	}
}
