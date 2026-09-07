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
	"strings"
	"unicode"

	"connectrpc.com/connect"
	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"
)

// Identity is what a verified token names: the user id, and the email
// the allowlist reads (D-314). The email is empty for a token with none.
type Identity struct {
	UID   string
	Email string
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

// Verify implements Verifier. The email comes from the token's claims,
// and Firebase sets it for an email and password account.
func (f *Firebase) Verify(ctx context.Context, idToken string) (Identity, error) {
	tok, err := f.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return Identity{}, err
	}
	if tok.UID == "" {
		return Identity{}, errors.New("token carries no uid")
	}
	email, _ := tok.Claims["email"].(string)
	return Identity{UID: tok.UID, Email: email}, nil
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

// Option tunes the interceptor.
type Option func(*interceptor)

// WithFallback names the user a request with no bearer token acts as.
// It is the DEBUG_USER_ID of local mode. Empty means no fallback, and
// then a missing token is Unauthenticated.
func WithFallback(uid string) Option {
	return func(i *interceptor) { i.fallback = uid }
}

// WithAllowlist refuses every verified user whose email is not on the
// list, with PermissionDenied and one sentence (D-314). The deployed API
// sets it, and local mode does not, so every emulator user gets in.
func WithAllowlist(a Allowlist) Option {
	return func(i *interceptor) { i.allow = a }
}

type interceptor struct {
	verify   Verifier
	fallback string
	allow    Allowlist
	// public names the procedures that need no sign-in, the shared deck
	// reads of D-315. Such a call carries no user in its context.
	public map[string]bool
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
)

// resolve reads the Authorization header and returns the context that
// carries the user, or the Connect error to answer with.
func (i *interceptor) resolve(ctx context.Context, authorization string) (context.Context, error) {
	token, present := bearer(authorization)
	if !present {
		if i.fallback == "" {
			return nil, connect.NewError(connect.CodeUnauthenticated, errNoToken)
		}
		return WithUserID(ctx, i.fallback), nil
	}
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoToken)
	}
	id, err := i.verify.Verify(ctx, token)
	if err != nil {
		// The verifier's reason stays in the log, not on the wire: it can
		// name the project or the key id.
		return nil, connect.NewError(connect.CodeUnauthenticated, errBadToken)
	}
	if id.UID == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errBadToken)
	}
	if i.allow != nil {
		ok, err := i.allow.Allowed(ctx, id.Email)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, errors.New("the invite list could not be read"))
		}
		if !ok {
			return nil, connect.NewError(connect.CodePermissionDenied, errNotInvited)
		}
	}
	return WithEmail(WithUserID(ctx, id.UID), id.Email), nil
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
		ctx, err := i.resolve(ctx, req.Header().Get("Authorization"))
		if err != nil {
			return nil, err
		}
		return next(ctx, req)
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
		ctx, err := i.resolve(ctx, conn.RequestHeader().Get("Authorization"))
		if err != nil {
			return err
		}
		return next(ctx, conn)
	}
}
