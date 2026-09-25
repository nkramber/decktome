package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// unsignedToken is an ID token in the shape the Auth emulator issues. With
// FIREBASE_AUTH_EMULATOR_HOST set, the Admin SDK checks the claims and not
// the signature, so the test needs no key.
func unsignedToken(t *testing.T, project string, claims map[string]any) string {
	t.Helper()
	now := time.Now().Unix()
	body := map[string]any{
		"iss": "https://securetoken.google.com/" + project, "aud": project,
		"sub": "u-42", "iat": now, "exp": now + 3600, "auth_time": now,
	}
	for k, v := range claims {
		body[k] = v
	}
	enc := func(v any) string {
		raw, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		return base64.RawURLEncoding.EncodeToString(raw)
	}
	return enc(map[string]any{"alg": "none", "typ": "JWT"}) + "." + enc(body) + "."
}

// TestFirebaseReadsTheProvedEmail is D-903: the verifier carries the
// email_verified claim, and a token with no such claim is not proved.
func TestFirebaseReadsTheProvedEmail(t *testing.T) {
	// The SDK looks the account up on the emulator, so a fake answers
	// with one account that is not disabled.
	emu := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/accounts:lookup") {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte(`{"users":[{"localId":"u-42","email":"ann@example.com"}]}`))
	}))
	t.Cleanup(emu.Close)
	t.Setenv("FIREBASE_AUTH_EMULATOR_HOST", strings.TrimPrefix(emu.URL, "http://"))
	f, err := NewFirebase(context.Background(), "mtg-test")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name   string
		claims map[string]any
		want   bool
	}{
		{"a proved email", map[string]any{"email": "ann@example.com", "email_verified": true}, true},
		{"an email that is not proved", map[string]any{"email": "ann@example.com", "email_verified": false}, false},
		{"no claim", map[string]any{"email": "ann@example.com"}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			id, err := f.Verify(context.Background(), unsignedToken(t, "mtg-test", tc.claims))
			if err != nil {
				t.Fatalf("verify: %v", err)
			}
			if id.UID != "u-42" || id.Email != "ann@example.com" || id.EmailVerified != tc.want {
				t.Errorf("identity = %+v, want verified %v", id, tc.want)
			}
		})
	}
}
