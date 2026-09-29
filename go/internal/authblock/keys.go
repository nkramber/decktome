package authblock

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// CertURL holds the certificates of the securetoken service account,
// keyed by key id. The Firebase Admin SDK reads the same URL for a
// blocking token.
const CertURL = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

const (
	// defaultTTL holds the keys when the answer names no max-age.
	defaultTTL = time.Hour
	// refetchGap bounds the fetches that an unknown key id starts.
	refetchGap = time.Minute
	fetchLimit = 1 << 20
)

// GoogleKeys fetches the certificates and keeps them for the max-age of
// the answer.
type GoogleKeys struct {
	url    string
	client *http.Client
	now    func() time.Time

	mu      sync.Mutex
	keys    map[string]*rsa.PublicKey
	expires time.Time
	fetched time.Time
}

// NewGoogleKeys returns the cache for the securetoken certificates.
func NewGoogleKeys(client *http.Client) *GoogleKeys {
	return &GoogleKeys{url: CertURL, client: client, now: time.Now}
}

// Key returns the key for one key id. An unknown id fetches again at
// most once a minute, because Google rotates the keys.
func (g *GoogleKeys) Key(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	now := g.now()
	stale := g.keys == nil || now.After(g.expires)
	if k, ok := g.keys[kid]; ok && !stale {
		return k, nil
	}
	if stale || now.Sub(g.fetched) >= refetchGap {
		if err := g.fetch(ctx, now); err != nil {
			return nil, err
		}
	}
	if k, ok := g.keys[kid]; ok {
		return k, nil
	}
	return nil, fmt.Errorf("unknown key id %q", kid)
}

func (g *GoogleKeys) fetch(ctx context.Context, now time.Time) error {
	g.fetched = now
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, g.url, nil)
	if err != nil {
		return err
	}
	resp, err := g.client.Do(req)
	if err != nil {
		return fmt.Errorf("fetch certificates: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetch certificates: status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, fetchLimit))
	if err != nil {
		return err
	}
	keys, err := parseCerts(body)
	if err != nil {
		return err
	}
	g.keys = keys
	g.expires = now.Add(maxAge(resp.Header.Get("Cache-Control")))
	return nil
}

func parseCerts(body []byte) (map[string]*rsa.PublicKey, error) {
	var certs map[string]string
	if err := json.Unmarshal(body, &certs); err != nil {
		return nil, fmt.Errorf("parse certificates: %w", err)
	}
	keys := map[string]*rsa.PublicKey{}
	for kid, p := range certs {
		block, _ := pem.Decode([]byte(p))
		if block == nil {
			return nil, fmt.Errorf("certificate %q: no PEM block", kid)
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("certificate %q: %w", kid, err)
		}
		k, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("certificate %q: not an RSA key", kid)
		}
		keys[kid] = k
	}
	if len(keys) == 0 {
		return nil, errors.New("no certificates")
	}
	return keys, nil
}

func maxAge(cacheControl string) time.Duration {
	for _, part := range strings.Split(cacheControl, ",") {
		v, ok := strings.CutPrefix(strings.TrimSpace(part), "max-age=")
		if !ok {
			continue
		}
		if s, err := strconv.Atoi(v); err == nil && s > 0 {
			return time.Duration(s) * time.Second
		}
	}
	return defaultTTL
}
