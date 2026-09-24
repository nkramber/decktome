package notify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unicode/utf8"
)

// TestPushoverPostsTheForm pins the four fields of the Pushover API and
// the clip of each limit.
func TestPushoverPostsTheForm(t *testing.T) {
	var got map[string]string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse: %v", err)
		}
		got = map[string]string{}
		for k := range r.PostForm {
			got[k] = r.PostForm.Get(k)
		}
		_, _ = w.Write([]byte(`{"status":1}`))
	}))
	defer srv.Close()
	long := strings.Repeat("é", 2000)
	p := NewPushover("tok", "usr").WithURL(srv.URL)
	if err := p.Send(context.Background(), Notice{Title: long, Message: long}); err != nil {
		t.Fatalf("send: %v", err)
	}
	if got["token"] != "tok" || got["user"] != "usr" {
		t.Errorf("token %q, user %q", got["token"], got["user"])
	}
	if n := utf8.RuneCountInString(got["title"]); n != maxTitle {
		t.Errorf("title holds %d characters, want %d", n, maxTitle)
	}
	if n := utf8.RuneCountInString(got["message"]); n != maxMessage {
		t.Errorf("message holds %d characters, want %d", n, maxMessage)
	}
}

// TestPushoverErrorNamesNoSecret: a refused send is an error, and the
// error holds neither the token nor the key (D-639).
func TestPushoverErrorNamesNoSecret(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"status":0,"errors":["application token is invalid"]}`))
	}))
	defer srv.Close()
	err := NewPushover("secret-token", "secret-user").WithURL(srv.URL).Send(context.Background(), Notice{Message: "hi"})
	if err == nil {
		t.Fatal("a 400 passed")
	}
	if strings.Contains(err.Error(), "secret") {
		t.Errorf("the error names a secret: %v", err)
	}
}

func TestFromEnvNeedsBoth(t *testing.T) {
	for _, tc := range []struct {
		env  map[string]string
		want bool
	}{
		{map[string]string{"PUSHOVER_APP_TOKEN": "t", "PUSHOVER_USER_KEY": "u"}, true},
		{map[string]string{"PUSHOVER_APP_TOKEN": "t"}, false},
		{map[string]string{"PUSHOVER_USER_KEY": "u"}, false},
		{map[string]string{}, false},
	} {
		got := FromEnv(func(k string) string { return tc.env[k] }) != nil
		if got != tc.want {
			t.Errorf("FromEnv(%v) set = %v, want %v", tc.env, got, tc.want)
		}
	}
}

func TestClip(t *testing.T) {
	if got := Clip("short", 10); got != "short" {
		t.Errorf("Clip kept %q", got)
	}
	if got := Clip("abcdefghij", 6); got != "abc..." {
		t.Errorf("Clip cut to %q, want abc...", got)
	}
}

type countSender struct {
	mu      sync.Mutex
	sent    []Notice
	running atomic.Int32
	most    atomic.Int32
	err     error
}

func (c *countSender) Send(_ context.Context, n Notice) error {
	now := c.running.Add(1)
	defer c.running.Add(-1)
	for {
		most := c.most.Load()
		if now <= most || c.most.CompareAndSwap(most, now) {
			break
		}
	}
	time.Sleep(time.Millisecond)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sent = append(c.sent, n)
	return c.err
}

// TestBackgroundSendsOneAtATime: Notify returns at once, every notice
// sends, and no two sends overlap, because Pushover asks for at most two
// connections at once.
func TestBackgroundSendsOneAtATime(t *testing.T) {
	c := &countSender{}
	b := NewBackground(c, nil)
	for i := 0; i < 5; i++ {
		b.Notify(Notice{Message: "x"})
	}
	b.Wait()
	if len(c.sent) != 5 {
		t.Fatalf("sent %d, want 5", len(c.sent))
	}
	if m := c.most.Load(); m != 1 {
		t.Errorf("%d sends overlapped, want 1", m)
	}
}

// TestBackgroundSurvivesAFailure: a failed send ends quietly, with no
// logger as well.
func TestBackgroundSurvivesAFailure(t *testing.T) {
	c := &countSender{err: errors.New("down")}
	b := NewBackground(c, nil)
	b.Notify(Notice{Message: "x"})
	b.Wait()
	if len(c.sent) != 1 {
		t.Fatalf("sent %d, want 1", len(c.sent))
	}
}
