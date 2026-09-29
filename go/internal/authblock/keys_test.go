package authblock

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func certPEM(t *testing.T, key *rsa.PrivateKey) string {
	t.Helper()
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "securetoken.system.gserviceaccount.com"},
		NotBefore:    testNow.Add(-time.Hour),
		NotAfter:     testNow.Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
}

// certServer serves the certificates in the shape of CertURL and counts
// the fetches.
func certServer(t *testing.T, certs *atomic.Value, fetches *atomic.Int32) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fetches.Add(1)
		w.Header().Set("Cache-Control", "public, max-age=600, must-revalidate")
		_ = json.NewEncoder(w).Encode(certs.Load())
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestGoogleKeysCachesForMaxAge(t *testing.T) {
	key := newKey(t)
	var certs atomic.Value
	certs.Store(map[string]string{testKid: certPEM(t, key)})
	var fetches atomic.Int32
	srv := certServer(t, &certs, &fetches)
	now := testNow
	g := NewGoogleKeys(srv.Client())
	g.url = srv.URL
	g.now = func() time.Time { return now }

	for range 3 {
		k, err := g.Key(context.Background(), testKid)
		if err != nil || !k.Equal(&key.PublicKey) {
			t.Fatalf("got %v %v, want the served key", k, err)
		}
	}
	if n := fetches.Load(); n != 1 {
		t.Fatalf("got %d fetches, want 1 inside max-age", n)
	}
	now = now.Add(11 * time.Minute)
	if _, err := g.Key(context.Background(), testKid); err != nil {
		t.Fatal(err)
	}
	if n := fetches.Load(); n != 2 {
		t.Fatalf("got %d fetches, want 2 after max-age", n)
	}
}

// Google rotates the keys, so an unknown id fetches again, at most once
// a minute.
func TestGoogleKeysRefetchOnRotation(t *testing.T) {
	first, second := newKey(t), newKey(t)
	var certs atomic.Value
	certs.Store(map[string]string{testKid: certPEM(t, first)})
	var fetches atomic.Int32
	srv := certServer(t, &certs, &fetches)
	now := testNow
	g := NewGoogleKeys(srv.Client())
	g.url = srv.URL
	g.now = func() time.Time { return now }

	if _, err := g.Key(context.Background(), testKid); err != nil {
		t.Fatal(err)
	}
	certs.Store(map[string]string{testKid: certPEM(t, first), "kid-2": certPEM(t, second)})
	if _, err := g.Key(context.Background(), "kid-2"); err == nil {
		t.Fatal("an unknown id inside the gap fetched again")
	}
	now = now.Add(refetchGap)
	k, err := g.Key(context.Background(), "kid-2")
	if err != nil || !k.Equal(&second.PublicKey) {
		t.Fatalf("got %v %v, want the rotated key", k, err)
	}
	if n := fetches.Load(); n != 2 {
		t.Fatalf("got %d fetches, want 2", n)
	}
}

func TestGoogleKeysRefusesBadAnswers(t *testing.T) {
	for name, body := range map[string]string{
		"not json": "nope",
		"empty":    "{}",
		"no pem":   `{"k":"not a certificate"}`,
	} {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				_, _ = w.Write([]byte(body))
			}))
			defer srv.Close()
			g := NewGoogleKeys(srv.Client())
			g.url = srv.URL
			if _, err := g.Key(context.Background(), "k"); err == nil {
				t.Fatal("got a key from a bad answer")
			}
		})
	}
}

func TestMaxAge(t *testing.T) {
	for in, want := range map[string]time.Duration{
		"public, max-age=19204, must-revalidate, no-transform": 19204 * time.Second,
		"no-cache":   defaultTTL,
		"":           defaultTTL,
		"max-age=-1": defaultTTL,
	} {
		if got := maxAge(in); got != want {
			t.Errorf("maxAge(%q) = %v, want %v", in, got, want)
		}
	}
}
