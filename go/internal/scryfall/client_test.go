package scryfall

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestBulkFilesAndDownload(t *testing.T) {
	var gotUA, gotAccept string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUA = r.Header.Get("User-Agent")
		switch r.URL.Path {
		case "/bulk-data":
			_, _ = w.Write([]byte(`{"data":[{"type":"oracle_cards","updated_at":"2026-08-24T21:00:00Z","jsonl_download_uri":"` + "http://" + r.Host + `/file"}]}`))
		case "/file":
			gotAccept = r.Header.Get("Accept")
			_, _ = w.Write([]byte("line1\n"))
		default:
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()

	c := New(srv.Client(), srv.URL, slog.Default())
	files, err := c.BulkFiles(context.Background())
	if err != nil {
		t.Fatalf("BulkFiles: %v", err)
	}
	f, ok := files["oracle_cards"]
	if !ok {
		t.Fatal("oracle_cards missing")
	}
	if gotUA != UserAgent {
		t.Errorf("user agent = %q, want %q", gotUA, UserAgent)
	}
	body, err := c.Download(context.Background(), f.DownloadURI)
	if err != nil {
		t.Fatalf("Download: %v", err)
	}
	defer func() { _ = body.Close() }()
	data, _ := io.ReadAll(body)
	if string(data) != "line1\n" {
		t.Errorf("download = %q", data)
	}
	if gotAccept == "" {
		t.Error("download sent no Accept header")
	}
}

func TestRateLimit(t *testing.T) {
	tests := []struct {
		name       string
		retryAfter string
		limitAll   bool
		wantErr    bool
		wantCalls  int32
		minElapsed time.Duration
	}{
		{name: "retry-after honored", retryAfter: "1", wantCalls: 2, minElapsed: time.Second},
		{name: "default wait when header absent", wantCalls: 2, minElapsed: 50 * time.Millisecond},
		{name: "second 429 fails", retryAfter: "1", limitAll: true, wantErr: true, wantCalls: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var calls atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				n := calls.Add(1)
				if n == 1 || tt.limitAll {
					if tt.retryAfter != "" {
						w.Header().Set("Retry-After", tt.retryAfter)
					}
					w.WriteHeader(http.StatusTooManyRequests)
					return
				}
				_, _ = w.Write([]byte(`{"data":[]}`))
			}))
			defer srv.Close()
			c := New(srv.Client(), srv.URL, slog.Default())
			c.retryAfter = 50 * time.Millisecond
			start := time.Now()
			_, err := c.BulkFiles(context.Background())
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got := calls.Load(); got != tt.wantCalls {
				t.Errorf("calls = %d, want %d", got, tt.wantCalls)
			}
			if el := time.Since(start); el < tt.minElapsed {
				t.Errorf("elapsed %v, want at least %v", el, tt.minElapsed)
			}
		})
	}
}

func TestRateLimitContextCancel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()
	c := New(srv.Client(), srv.URL, slog.Default())
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if _, err := c.BulkFiles(ctx); err == nil {
		t.Fatal("want error on cancelled wait")
	}
}
