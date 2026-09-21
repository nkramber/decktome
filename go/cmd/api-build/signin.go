package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
)

// identityToolkit is the Firebase sign-in endpoint. The deployed project
// enables the email and password provider alone, so this call needs no
// browser (D-779).
const identityToolkit = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword"

// signInResult holds what the sign-in returns. The id token is what the
// API verifies, and it lasts one hour.
type signInResult struct {
	IDToken   string `json:"idToken"`
	LocalID   string `json:"localId"`
	Email     string `json:"email"`
	ExpiresIn string `json:"expiresIn"`
}

// signInError is the failure shape of the same endpoint.
type signInError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// signIn exchanges an email and a password for a Firebase id token. The
// endpoint is overridable so a test can serve it.
func signIn(ctx context.Context, c *http.Client, endpoint, apiKey, email, password string) (*signInResult, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("no Firebase web API key: set FIREBASE_API_KEY")
	}
	if email == "" || password == "" {
		return nil, fmt.Errorf("no account: set API_BUILD_EMAIL and API_BUILD_PASSWORD to the check account (D-779)")
	}
	body, err := json.Marshal(map[string]any{
		"email":             email,
		"password":          password,
		"returnSecureToken": true,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint+"?key="+url.QueryEscape(apiKey), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sign-in: %w", err)
	}
	defer func() { _ = res.Body.Close() }()
	dec := json.NewDecoder(res.Body)
	if res.StatusCode != http.StatusOK {
		var e signInError
		// The reason stays on one line, and it never carries the password.
		if err := dec.Decode(&e); err != nil || e.Error.Message == "" {
			return nil, fmt.Errorf("sign-in refused with status %d", res.StatusCode)
		}
		return nil, fmt.Errorf("sign-in refused: %s", e.Error.Message)
	}
	var out signInResult
	if err := dec.Decode(&out); err != nil {
		return nil, fmt.Errorf("sign-in answer: %w", err)
	}
	if out.IDToken == "" {
		return nil, fmt.Errorf("sign-in returned no id token")
	}
	return &out, nil
}

// bearer adds the id token to every call, the way the web app does
// (D-268). It wraps the transport, so the streaming Chat call carries
// the header too.
type bearer struct {
	token string
	next  http.RoundTripper
}

func (b bearer) RoundTrip(r *http.Request) (*http.Response, error) {
	r = r.Clone(r.Context())
	r.Header.Set("Authorization", "Bearer "+b.token)
	return b.next.RoundTrip(r)
}

// newHTTPClient builds the client the run uses. A build takes minutes,
// so the timeout covers a whole streamed turn.
func newHTTPClient(token string, timeout time.Duration) *http.Client {
	if token == "" {
		return &http.Client{Transport: http.DefaultTransport, Timeout: timeout}
	}
	return &http.Client{Transport: bearer{token: token, next: http.DefaultTransport}, Timeout: timeout}
}

// readyPoll is the gap between two health checks. A test shortens it.
var readyPoll = 3 * time.Second

// waitReady holds the run until the API loads the card snapshot. A
// Cloud Run service scales to zero, so a cold instance answers an
// import with "card database not loaded yet". HealthService.Check is
// public, and it names the snapshot.
func waitReady(ctx context.Context, h mtgv1connect.HealthServiceClient, limit time.Duration, log func(string, ...any)) error {
	deadline := time.Now().Add(limit)
	said := false
	for {
		// Each check gets its own limit. The client timeout covers a whole
		// streamed build, and a hung check must not outlive the deadline.
		res, err := checkOnce(ctx, h)
		switch {
		case err != nil:
			if time.Now().After(deadline) {
				return fmt.Errorf("the API never answered a health check in %s: %w", limit, err)
			}
		case res.Msg.GetCardSnapshot() != "" && res.Msg.GetCardSnapshot() != "none":
			log("ready        the API holds the card snapshot of %s\n", res.Msg.GetCardSnapshot())
			return nil
		case time.Now().After(deadline):
			return fmt.Errorf("the API loaded no card snapshot in %s: a cold instance needs longer, or the bucket holds none", limit)
		}
		if !said {
			log("waiting      the API loads the card snapshot. A cold instance takes a minute\n")
			said = true
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(readyPoll):
		}
	}
}

// checkOnce reads the health of the API one time, under a short limit.
func checkOnce(ctx context.Context, h mtgv1connect.HealthServiceClient) (*connect.Response[mtgv1.CheckResponse], error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	return h.Check(ctx, connect.NewRequest(&mtgv1.CheckRequest{}))
}
