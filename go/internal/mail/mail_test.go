package mail

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestSendPostsTheResendShape is D-1077: one JSON post with the key as a
// bearer token and a User-Agent, which Resend requires.
func TestSendPostsTheResendShape(t *testing.T) {
	var got map[string]any
	var auth, agent string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, agent = r.Header.Get("Authorization"), r.Header.Get("User-Agent")
		_ = json.NewDecoder(r.Body).Decode(&got)
		_, _ = w.Write([]byte(`{"id":"x"}`))
	}))
	t.Cleanup(srv.Close)
	r := NewResend("re_test", DefaultFrom).WithURL(srv.URL)
	if err := r.Send(context.Background(), Message{To: "ann@example.com", Subject: "s", Text: "t"}); err != nil {
		t.Fatal(err)
	}
	if auth != "Bearer re_test" || agent == "" {
		t.Errorf("headers: auth %q, agent %q", auth, agent)
	}
	to, _ := got["to"].([]any)
	if got["from"] != DefaultFrom || len(to) != 1 || to[0] != "ann@example.com" || got["subject"] != "s" || got["text"] != "t" {
		t.Errorf("body = %v", got)
	}
}

func TestSendAnswersAnErrorThatHidesTheKey(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	t.Cleanup(srv.Close)
	err := NewResend("re_secret", DefaultFrom).WithURL(srv.URL).Send(context.Background(), Message{To: "a@b.c"})
	if err == nil || strings.Contains(err.Error(), "re_secret") {
		t.Errorf("err = %v", err)
	}
}

func TestFromEnv(t *testing.T) {
	if FromEnv(func(string) string { return "" }) != nil {
		t.Error("no key must answer nil")
	}
	env := map[string]string{"RESEND_API_KEY": "k"}
	if r := FromEnv(func(k string) string { return env[k] }); r == nil || r.from != DefaultFrom {
		t.Errorf("from = %v", r)
	}
	env["MAIL_FROM"] = "X <x@y.z>"
	if r := FromEnv(func(k string) string { return env[k] }); r.from != "X <x@y.z>" {
		t.Errorf("from = %q", r.from)
	}
}
