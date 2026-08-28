// Package auth names the caller. A Connect interceptor reads the Firebase
// ID token from the Authorization header, verifies it, and puts the user
// id in the context (A-12). The services read it with UserID.
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

	"connectrpc.com/connect"
	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"
)

// Verifier checks one ID token and returns the user id it names.
type Verifier interface {
	Verify(ctx context.Context, idToken string) (string, error)
}

// VerifierFunc adapts a function to Verifier.
type VerifierFunc func(ctx context.Context, idToken string) (string, error)

// Verify implements Verifier.
func (f VerifierFunc) Verify(ctx context.Context, idToken string) (string, error) {
	return f(ctx, idToken)
}

// ErrNotConfigured is what RejectAll answers: the server has no verifier.
var ErrNotConfigured = errors.New("token verification is not configured on this server")

// RejectAll refuses every token. A local server with no Firebase project
// uses it, so the debug fallback is the only way in.
func RejectAll() Verifier {
	return VerifierFunc(func(context.Context, string) (string, error) { return "", ErrNotConfigured })
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

// Verify implements Verifier.
func (f *Firebase) Verify(ctx context.Context, idToken string) (string, error) {
	tok, err := f.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}
	if tok.UID == "" {
		return "", errors.New("token carries no uid")
	}
	return tok.UID, nil
}

type ctxKey struct{}

// UserID reads the user id the interceptor stored, or "" when none.
func UserID(ctx context.Context) string {
	uid, _ := ctx.Value(ctxKey{}).(string)
	return uid
}

// WithUserID returns ctx with uid set. Tests and the interceptor use it.
func WithUserID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, ctxKey{}, uid)
}

// Option tunes the interceptor.
type Option func(*interceptor)

// WithFallback names the user a request with no bearer token acts as.
// It is the DEBUG_USER_ID of local mode. Empty means no fallback, and
// then a missing token is Unauthenticated.
func WithFallback(uid string) Option {
	return func(i *interceptor) { i.fallback = uid }
}

type interceptor struct {
	verify   Verifier
	fallback string
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
	uid, err := i.verify.Verify(ctx, token)
	if err != nil {
		// The verifier's reason stays in the log, not on the wire: it can
		// name the project or the key id.
		return nil, connect.NewError(connect.CodeUnauthenticated, errBadToken)
	}
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errBadToken)
	}
	return WithUserID(ctx, uid), nil
}

// bearer splits "Bearer <token>". present is false when the header is
// absent or names another scheme. A bare "Bearer" is present with an
// empty token, because the transport trims the trailing space.
func bearer(authorization string) (token string, present bool) {
	const scheme = "bearer"
	v := strings.TrimSpace(authorization)
	if len(v) < len(scheme) || !strings.EqualFold(v[:len(scheme)], scheme) {
		return "", false
	}
	rest := v[len(scheme):]
	if rest != "" && rest[0] != ' ' {
		return "", false
	}
	return strings.TrimSpace(rest), true
}

func (i *interceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		if req.Spec().IsClient {
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
		ctx, err := i.resolve(ctx, conn.RequestHeader().Get("Authorization"))
		if err != nil {
			return err
		}
		return next(ctx, conn)
	}
}
