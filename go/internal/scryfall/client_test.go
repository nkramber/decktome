package scryfall

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBulkFilesAndDownload(t *testing.T) {
	var gotUA string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUA = r.Header.Get("User-Agent")
		switch r.URL.Path {
		case "/bulk-data":
			_, _ = w.Write([]byte(`{"data":[{"type":"oracle_cards","updated_at":"2026-08-24T21:00:00Z","jsonl_download_uri":"` + "http://" + r.Host + `/file"}]}`))
		case "/file":
			_, _ = w.Write([]byte("line1\n"))
		default:
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()

	c := New(srv.Client(), srv.URL)
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
}
