// Package scryfall downloads Scryfall bulk data.
//
// F-3: the live API has hard rate limits. This package touches only the
// bulk-data endpoint (once per refresh) and the file origin on
// *.scryfall.io, which has no rate limit. Every request sends a real
// User-Agent and an Accept header, as the Scryfall terms require.
package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// UserAgent identifies this app to Scryfall.
const UserAgent = "mtg-deck-builder/0.1 (github.com/nkramber/mtg-deck-builder)"

// DefaultRetryAfter is the wait after a 429 with no Retry-After header.
const DefaultRetryAfter = 30 * time.Second

// BulkFile describes one bulk file from the bulk-data endpoint.
type BulkFile struct {
	Type        string    `json:"type"`
	UpdatedAt   time.Time `json:"updated_at"`
	DownloadURI string    `json:"jsonl_download_uri"`
}

// Client fetches bulk metadata and files.
type Client struct {
	http    *http.Client
	baseURL string
	logger  *slog.Logger
	// retryAfter is the wait after a 429 with no Retry-After header.
	// Tests shorten it.
	retryAfter time.Duration
}

// New returns a Client. baseURL is overridable for tests. A nil httpClient
// gets a transport with a 60 s header timeout and no whole-body timeout:
// the caller bounds a download with its context (C-14).
func New(httpClient *http.Client, baseURL string, logger *slog.Logger) *Client {
	if httpClient == nil {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.ResponseHeaderTimeout = 60 * time.Second
		httpClient = &http.Client{Transport: transport}
	}
	if baseURL == "" {
		baseURL = "https://api.scryfall.com"
	}
	if logger == nil {
		logger = slog.Default()
	}
	return &Client{http: httpClient, baseURL: baseURL, logger: logger, retryAfter: DefaultRetryAfter}
}

// do sends one request. On a 429 it waits for Retry-After (seconds), or
// the default wait, then retries once (C-13). Scryfall asks clients to
// honor 429 and never to ignore it.
func (c *Client) do(ctx context.Context, req *http.Request, what string) (*http.Response, error) {
	req.Header.Set("User-Agent", UserAgent)
	for attempt := 0; ; attempt++ {
		res, err := c.http.Do(req.Clone(ctx))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", what, err)
		}
		if res.StatusCode == http.StatusTooManyRequests && attempt == 0 {
			wait := c.retryWait(res.Header.Get("Retry-After"))
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
			c.logger.Warn("scryfall rate limited, retry once", "what", what, "wait", wait.String())
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("%s: %w", what, ctx.Err())
			case <-time.After(wait):
			}
			continue
		}
		if res.StatusCode != http.StatusOK {
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
			return nil, fmt.Errorf("%s: status %d", what, res.StatusCode)
		}
		return res, nil
	}
}

func (c *Client) retryWait(header string) time.Duration {
	if secs, err := strconv.Atoi(header); err == nil && secs > 0 {
		return time.Duration(secs) * time.Second
	}
	return c.retryAfter
}

// BulkFiles returns the bulk-data catalog keyed by type.
func (c *Client) BulkFiles(ctx context.Context) (map[string]BulkFile, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/bulk-data", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.do(ctx, req, "scryfall bulk-data")
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()
	var payload struct {
		Data []BulkFile `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("scryfall bulk-data: decode: %w", err)
	}
	out := make(map[string]BulkFile, len(payload.Data))
	for _, f := range payload.Data {
		out[f.Type] = f
	}
	return out, nil
}

// Download streams one bulk file. The caller closes the reader and bounds
// the whole transfer with ctx.
func (c *Client) Download(ctx context.Context, uri string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	res, err := c.do(ctx, req, "scryfall download")
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
