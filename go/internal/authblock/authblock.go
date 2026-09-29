// Package authblock answers the beforeCreate blocking function of
// Identity Platform. It refuses an account whose email is not on the
// invite list, before Identity Platform saves the account (OQ-77, D-990).
//
// Without it a caller who drives the Firebase API directly still makes
// an account. That account reads nothing, because the API refuses every
// call (D-314), but the project never invited it.
//
// Identity Platform POSTs {"data":{"jwt":...}} to the registered URI. The
// token is RS256, signed by the securetoken service account, issued for
// the project, and its audience is the registered URI. A 200 lets the
// account through. Any other answer, or no answer in seven seconds,
// fails the sign-up, so this door fails closed.
package authblock

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// Path is the route that Identity Platform calls as beforeCreate.
	Path = "/auth/before-create"
	// Refusal is the message of a refused sign-up. The web form reads it
	// in the auth error to show the invite sentence (D-990).
	Refusal = "not-invited"
	// IssuerPrefix plus the project id is the issuer of every token.
	IssuerPrefix = "https://securetoken.google.com/"
	// EventBeforeCreate is the one event this route answers.
	EventBeforeCreate = "beforeCreate"

	// skew is the clock difference the time claims tolerate.
	skew = time.Minute
	// maxBody bounds the request body. The token is a few kilobytes.
	maxBody = 64 << 10
	// allowTimeout keeps the list read inside the seven-second limit.
	allowTimeout = 5 * time.Second
)

// Allow answers whether an email is on the invite list.
type Allow func(ctx context.Context, email string) (bool, error)

// Keys returns the public key that signed a token, by its key id.
type Keys interface {
	Key(ctx context.Context, kid string) (*rsa.PublicKey, error)
}

// Handler answers the beforeCreate call.
type Handler struct {
	project string
	keys    Keys
	allow   Allow
	logger  *slog.Logger
	now     func() time.Time
}

// New returns the handler for one project.
func New(project string, keys Keys, allow Allow, logger *slog.Logger) *Handler {
	return &Handler{project: project, keys: keys, allow: allow, logger: logger, now: time.Now}
}

// WithClock sets the clock of the time checks, for a test.
func (h *Handler) WithClock(now func() time.Time) *Handler {
	h.now = now
	return h
}

type request struct {
	Data struct {
		JWT string `json:"jwt"`
	} `json:"data"`
}

type claims struct {
	jwt.RegisteredClaims
	EventType  string `json:"event_type"`
	Email      string `json:"email"`
	UserRecord struct {
		Email string `json:"email"`
	} `json:"user_record"`
}

// ServeHTTP checks the token, then the invite list.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "INVALID_ARGUMENT", "method not allowed")
		return
	}
	var req request
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&req); err != nil || req.Data.JWT == "" {
		writeError(w, http.StatusBadRequest, "INVALID_ARGUMENT", "no token")
		return
	}
	c, err := h.verify(r.Context(), req.Data.JWT, r.Host)
	if err != nil {
		h.logger.Warn("blocking token refused", "err", err)
		writeError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "bad token")
		return
	}
	email := c.UserRecord.Email
	if email == "" {
		email = c.Email
	}
	if email == "" {
		h.logger.Info("sign-up refused: no email")
		writeError(w, http.StatusForbidden, "PERMISSION_DENIED", Refusal)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), allowTimeout)
	defer cancel()
	ok, err := h.allow(ctx, email)
	if err != nil {
		h.logger.Error("sign-up check: invite list unavailable", "err", err)
		writeError(w, http.StatusServiceUnavailable, "UNAVAILABLE", "invite list unavailable")
		return
	}
	if !ok {
		h.logger.Info("sign-up refused: not on the invite list")
		writeError(w, http.StatusForbidden, "PERMISSION_DENIED", Refusal)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("{}"))
}

// verify checks the signature, the issuer, the audience, the times, and
// the event type. The audience must name the host that received the
// call, so a token for another endpoint does not pass here.
func (h *Handler) verify(ctx context.Context, token, host string) (*claims, error) {
	var c claims
	parser := jwt.NewParser(jwt.WithValidMethods([]string{"RS256"}), jwt.WithoutClaimsValidation())
	_, err := parser.ParseWithClaims(token, &c, func(t *jwt.Token) (any, error) {
		kid, _ := t.Header["kid"].(string)
		if kid == "" {
			return nil, errors.New("no key id")
		}
		return h.keys.Key(ctx, kid)
	})
	if err != nil {
		return nil, err
	}
	if c.Issuer != IssuerPrefix+h.project {
		return nil, fmt.Errorf("issuer %q", c.Issuer)
	}
	if !audienceNames(c.Audience, host) {
		return nil, fmt.Errorf("audience %v does not name host %q", []string(c.Audience), host)
	}
	now := h.now()
	if c.ExpiresAt == nil || now.After(c.ExpiresAt.Add(skew)) {
		return nil, errors.New("token expired")
	}
	if c.IssuedAt == nil || c.IssuedAt.After(now.Add(skew)) {
		return nil, errors.New("token issued in the future")
	}
	if c.EventType != EventBeforeCreate {
		return nil, fmt.Errorf("event type %q", c.EventType)
	}
	return &c, nil
}

func audienceNames(aud jwt.ClaimStrings, host string) bool {
	for _, a := range aud {
		u, err := url.Parse(a)
		if err == nil && u.Scheme == "https" && u.Host == host {
			return true
		}
	}
	return false
}

func writeError(w http.ResponseWriter, code int, status, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"error": map[string]string{"status": status, "message": message},
	})
}
