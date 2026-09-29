package authblock

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	testProject = "demo-project"
	testHost    = "api.example.run.app"
	testKid     = "kid-1"
)

var testNow = time.Date(2026, 9, 29, 12, 0, 0, 0, time.UTC)

type staticKeys map[string]*rsa.PublicKey

func (s staticKeys) Key(_ context.Context, kid string) (*rsa.PublicKey, error) {
	if k, ok := s[kid]; ok {
		return k, nil
	}
	return nil, fmt.Errorf("unknown key id %q", kid)
}

func newKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	return k
}

// goodClaims are the claims of a real beforeCreate call.
func goodClaims(email string) jwt.MapClaims {
	return jwt.MapClaims{
		"iss":         IssuerPrefix + testProject,
		"aud":         "https://" + testHost + Path,
		"iat":         testNow.Add(-time.Second).Unix(),
		"exp":         testNow.Add(time.Hour).Unix(),
		"event_type":  EventBeforeCreate,
		"user_record": map[string]any{"uid": "u1", "email": email},
	}
}

func sign(t *testing.T, key *rsa.PrivateKey, kid string, c jwt.MapClaims) string {
	t.Helper()
	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	if kid != "" {
		tok.Header["kid"] = kid
	}
	s, err := tok.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func allowOnly(emails ...string) Allow {
	return func(_ context.Context, email string) (bool, error) {
		for _, e := range emails {
			if e == email {
				return true, nil
			}
		}
		return false, nil
	}
}

func call(t *testing.T, h http.Handler, method, body string) (int, map[string]map[string]string) {
	t.Helper()
	req := httptest.NewRequest(method, "https://"+testHost+Path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	var out map[string]map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &out)
	return rec.Code, out
}

func bodyOf(token string) string {
	b, _ := json.Marshal(map[string]any{"data": map[string]string{"jwt": token}})
	return string(b)
}

func newHandler(key *rsa.PrivateKey, allow Allow) *Handler {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	return New(testProject, staticKeys{testKid: &key.PublicKey}, allow, logger).WithClock(func() time.Time { return testNow })
}

func TestInvitedEmailPasses(t *testing.T) {
	key := newKey(t)
	h := newHandler(key, allowOnly("invited@example.com"))
	code, out := call(t, h, http.MethodPost, bodyOf(sign(t, key, testKid, goodClaims("invited@example.com"))))
	if code != http.StatusOK || out["error"] != nil {
		t.Fatalf("got %d %v, want 200 and no error", code, out)
	}
}

func TestEmailOffTheListIsRefused(t *testing.T) {
	key := newKey(t)
	h := newHandler(key, allowOnly("invited@example.com"))
	code, out := call(t, h, http.MethodPost, bodyOf(sign(t, key, testKid, goodClaims("stranger@example.com"))))
	if code != http.StatusForbidden || out["error"]["message"] != Refusal || out["error"]["status"] != "PERMISSION_DENIED" {
		t.Fatalf("got %d %v, want 403 PERMISSION_DENIED %q", code, out, Refusal)
	}
}

func TestTopLevelEmailIsRead(t *testing.T) {
	key := newKey(t)
	c := goodClaims("")
	delete(c, "user_record")
	c["email"] = "invited@example.com"
	code, _ := call(t, newHandler(key, allowOnly("invited@example.com")), http.MethodPost, bodyOf(sign(t, key, testKid, c)))
	if code != http.StatusOK {
		t.Fatalf("got %d, want 200", code)
	}
}

func TestNoEmailIsRefused(t *testing.T) {
	key := newKey(t)
	code, out := call(t, newHandler(key, allowOnly("")), http.MethodPost, bodyOf(sign(t, key, testKid, goodClaims(""))))
	if code != http.StatusForbidden || out["error"]["message"] != Refusal {
		t.Fatalf("got %d %v, want 403 %q", code, out, Refusal)
	}
}

func TestListFailureFailsClosed(t *testing.T) {
	key := newKey(t)
	down := func(context.Context, string) (bool, error) { return false, errors.New("firestore down") }
	code, _ := call(t, newHandler(key, down), http.MethodPost, bodyOf(sign(t, key, testKid, goodClaims("invited@example.com"))))
	if code != http.StatusServiceUnavailable {
		t.Fatalf("got %d, want 503", code)
	}
}

// Every bad token answers 401, and the list is never read.
func TestBadTokensAreRefused(t *testing.T) {
	key := newKey(t)
	other := newKey(t)
	cases := map[string]func() string{
		"wrong issuer": func() string {
			c := goodClaims("invited@example.com")
			c["iss"] = IssuerPrefix + "other-project"
			return sign(t, key, testKid, c)
		},
		"audience of another host": func() string {
			c := goodClaims("invited@example.com")
			c["aud"] = "https://other.example.run.app" + Path
			return sign(t, key, testKid, c)
		},
		"audience over plain http": func() string {
			c := goodClaims("invited@example.com")
			c["aud"] = "http://" + testHost + Path
			return sign(t, key, testKid, c)
		},
		"expired": func() string {
			c := goodClaims("invited@example.com")
			c["exp"] = testNow.Add(-2 * time.Minute).Unix()
			return sign(t, key, testKid, c)
		},
		"no expiry": func() string {
			c := goodClaims("invited@example.com")
			delete(c, "exp")
			return sign(t, key, testKid, c)
		},
		"issued in the future": func() string {
			c := goodClaims("invited@example.com")
			c["iat"] = testNow.Add(2 * time.Minute).Unix()
			return sign(t, key, testKid, c)
		},
		"sign-in event": func() string {
			c := goodClaims("invited@example.com")
			c["event_type"] = "beforeSignIn"
			return sign(t, key, testKid, c)
		},
		"signed by another key": func() string {
			return sign(t, other, testKid, goodClaims("invited@example.com"))
		},
		"unknown key id": func() string {
			return sign(t, key, "kid-9", goodClaims("invited@example.com"))
		},
		"no key id": func() string {
			return sign(t, key, "", goodClaims("invited@example.com"))
		},
		"HS256 with the public key as secret": func() string {
			tok := jwt.NewWithClaims(jwt.SigningMethodHS256, goodClaims("invited@example.com"))
			tok.Header["kid"] = testKid
			s, err := tok.SignedString([]byte("secret"))
			if err != nil {
				t.Fatal(err)
			}
			return s
		},
		"alg none": func() string {
			tok := jwt.NewWithClaims(jwt.SigningMethodNone, goodClaims("invited@example.com"))
			tok.Header["kid"] = testKid
			s, err := tok.SignedString(jwt.UnsafeAllowNoneSignatureType)
			if err != nil {
				t.Fatal(err)
			}
			return s
		},
	}
	for name, token := range cases {
		t.Run(name, func(t *testing.T) {
			read := false
			allow := func(context.Context, string) (bool, error) { read = true; return true, nil }
			code, _ := call(t, newHandler(key, allow), http.MethodPost, bodyOf(token()))
			if code != http.StatusUnauthorized || read {
				t.Fatalf("got %d, list read %v; want 401 and no read", code, read)
			}
		})
	}
}

func TestSmallSkewPasses(t *testing.T) {
	key := newKey(t)
	c := goodClaims("invited@example.com")
	c["iat"] = testNow.Add(30 * time.Second).Unix()
	c["exp"] = testNow.Add(-30 * time.Second).Unix()
	code, _ := call(t, newHandler(key, allowOnly("invited@example.com")), http.MethodPost, bodyOf(sign(t, key, testKid, c)))
	if code != http.StatusOK {
		t.Fatalf("got %d, want 200 inside the skew", code)
	}
}

func TestBadRequests(t *testing.T) {
	key := newKey(t)
	h := newHandler(key, allowOnly("invited@example.com"))
	if code, _ := call(t, h, http.MethodGet, ""); code != http.StatusMethodNotAllowed {
		t.Errorf("GET: got %d, want 405", code)
	}
	for _, body := range []string{"", "not json", `{"data":{}}`, `{"jwt":"x"}`} {
		if code, _ := call(t, h, http.MethodPost, body); code != http.StatusBadRequest {
			t.Errorf("body %q: got %d, want 400", body, code)
		}
	}
}
