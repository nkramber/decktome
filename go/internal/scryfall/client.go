// Package scryfall downloads Scryfall bulk data.
//
// F-3: the live API has hard rate limits. This package touches only the
// bulk-data endpoint (once per refresh) and the file origin on
// *.scryfall.io, which has no rate limit. Every request sends a real
// User-Agent, as the Scryfall terms require.
package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// UserAgent identifies this app to Scryfall.
const UserAgent = "mtg-deck-builder/0.1 (github.com/nkramber/mtg-deck-builder)"

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
}

// New returns a Client. baseURL is overridable for tests.
func New(httpClient *http.Client, baseURL string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 5 * time.Minute}
	}
	if baseURL == "" {
		baseURL = "https://api.scryfall.com"
	}
	return &Client{http: httpClient, baseURL: baseURL}
}

// BulkFiles returns the bulk-data catalog keyed by type.
func (c *Client) BulkFiles(ctx context.Context) (map[string]BulkFile, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/bulk-data", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("scryfall bulk-data: %w", err)
	}
	defer func() { _ = res.Body.Close() }()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("scryfall bulk-data: status %d", res.StatusCode)
	}
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

// Download streams one bulk file. The caller closes the reader.
func (c *Client) Download(ctx context.Context, uri string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("scryfall download: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		_ = res.Body.Close()
		return nil, fmt.Errorf("scryfall download: status %d", res.StatusCode)
	}
	return res.Body, nil
}
