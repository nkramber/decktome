package scryfall

import (
	"context"
	"fmt"
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

// TestRetryWaitForms covers both Retry-After forms of RFC 9110: a delay
// in seconds and an HTTP-date. The date form was ignored before.
func TestRetryWaitForms(t *testing.T) {
	c := New(nil, "", nil)
	c.retryAfter = 50 * time.Millisecond
	if got := c.retryWait("3"); got != 3*time.Second {
		t.Errorf("seconds form = %v, want 3s", got)
	}
	future := time.Now().Add(5 * time.Second).UTC().Format(http.TimeFormat)
	if got := c.retryWait(future); got < 3*time.Second || got > 5*time.Second {
		t.Errorf("date form = %v, want about 4s", got)
	}
	past := time.Now().Add(-5 * time.Second).UTC().Format(http.TimeFormat)
	if got := c.retryWait(past); got != c.retryAfter {
		t.Errorf("past date = %v, want the default %v", got, c.retryAfter)
	}
	if got := c.retryWait("soon"); got != c.retryAfter {
		t.Errorf("garbage = %v, want the default", got)
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

// TestSetsReadsTheEndpoint is D-377: the set family comes from /sets,
// because no bulk card file carries a parent link.
func TestSetsReadsTheEndpoint(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/sets" {
			w.WriteHeader(404)
			return
		}
		_, _ = fmt.Fprint(w, `{"object":"list","has_more":false,"data":[
			{"code":"hob","name":"The Hobbit","set_type":"expansion","released_at":"2026-08-14","digital":false,"card_count":321},
			{"code":"hoc","name":"The Hobbit Eternal","set_type":"eternal","released_at":"2026-08-14","parent_set_code":"hob"}]}`)
	}))
	defer srv.Close()
	rows, err := New(srv.Client(), srv.URL, slog.Default()).Sets(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fatalf("read %d rows, want 2", len(rows))
	}
	if rows[0].Code != "hob" || rows[0].CardCount != 321 {
		t.Errorf("row 0 = %+v", rows[0])
	}
	if rows[1].ParentSetCode != "hob" {
		t.Errorf("row 1 parent = %q, want hob", rows[1].ParentSetCode)
	}
}

// TestSetsFollowsThePages covers a paged answer. The endpoint answered
// has_more false on 2026-08-31, and the loop must still work if it does
// not.
func TestSetsFollowsThePages(t *testing.T) {
	var srv *httptest.Server
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") == "2" {
			_, _ = fmt.Fprint(w, `{"has_more":false,"data":[{"code":"hoc","name":"The Hobbit Eternal"}]}`)
			return
		}
		_, _ = fmt.Fprintf(w, `{"has_more":true,"next_page":%q,"data":[{"code":"hob","name":"The Hobbit"}]}`,
			srv.URL+"/sets?page=2")
	}))
	defer srv.Close()
	rows, err := New(srv.Client(), srv.URL, slog.Default()).Sets(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 || rows[1].Code != "hoc" {
		t.Fatalf("rows = %+v", rows)
	}
}

// TestSetsRefusesAnEmptyAnswer: an empty set table would silently turn
// every set filter off.
func TestSetsRefusesAnEmptyAnswer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, `{"has_more":false,"data":[]}`)
	}))
	defer srv.Close()
	if _, err := New(srv.Client(), srv.URL, slog.Default()).Sets(t.Context()); err == nil {
		t.Fatal("an empty /sets answer must be an error")
	}
}
