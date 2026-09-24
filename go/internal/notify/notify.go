// Package notify sends the owner a short notice of each verdict a reader
// gives (PR-75, D-892 to D-897). The notice is a notice only: the
// harvest, the triage, and the fix cycle stay the path that acts on a
// verdict (D-890).
package notify

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// PushoverURL is the message endpoint of the Pushover API.
const PushoverURL = "https://api.pushover.net/1/messages.json"

// The Pushover limits of a message, read 2026-09-24 at pushover.net/api:
// a message of at most 1024 UTF-8 characters and a title of at most 250.
const (
	maxMessage = 1024
	maxTitle   = 250
)

// Timeout bounds one send, so a slow Pushover never holds a goroutine
// for long.
const Timeout = 10 * time.Second

// Notice is one message to the owner.
type Notice struct {
	Title   string
	Message string
}

// Sender delivers one notice.
type Sender interface {
	Send(ctx context.Context, n Notice) error
}

// Pushover sends a notice through the Pushover API (D-893). The app token
// and the user key live in Secret Manager, never in a file (D-639).
type Pushover struct {
	token, user, url string
	client           *http.Client
}

// NewPushover makes a sender for one app token and one user key.
func NewPushover(token, user string) *Pushover {
	return &Pushover{token: token, user: user, url: PushoverURL, client: &http.Client{Timeout: Timeout}}
}

// WithURL points the sender at another endpoint (tests).
func (p *Pushover) WithURL(u string) *Pushover {
	p.url = u
	return p
}

// FromEnv reads PUSHOVER_APP_TOKEN and PUSHOVER_USER_KEY. It answers nil
// when either one is empty, so a local run sends nothing.
func FromEnv(getenv func(string) string) *Pushover {
	token, user := getenv("PUSHOVER_APP_TOKEN"), getenv("PUSHOVER_USER_KEY")
	if token == "" || user == "" {
		return nil
	}
	return NewPushover(token, user)
}

// Send posts one notice. A status other than 200 is an error, and the
// error never names the token or the key.
func (p *Pushover) Send(ctx context.Context, n Notice) error {
	form := url.Values{
		"token":   {p.token},
		"user":    {p.user},
		"title":   {Clip(n.Title, maxTitle)},
		"message": {Clip(n.Message, maxMessage)},
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.url, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("pushover: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("pushover: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4<<10))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pushover: status %d", resp.StatusCode)
	}
	return nil
}

// Clip cuts s to at most n characters, and marks a cut with "...".
func Clip(s string, n int) string {
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	r := []rune(s)
	return string(r[:n-3]) + "..."
}

// Background sends each notice off the request path (D-892). One send
// runs at a time, because Pushover asks for at most two connections at
// once. A failure goes to the log, and the verdict stays stored.
type Background struct {
	sender Sender
	log    *slog.Logger
	one    sync.Mutex
	wg     sync.WaitGroup
}

// NewBackground wraps a sender.
func NewBackground(sender Sender, log *slog.Logger) *Background {
	return &Background{sender: sender, log: log}
}

// Notify starts one send and returns at once.
func (b *Background) Notify(n Notice) {
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		b.one.Lock()
		defer b.one.Unlock()
		ctx, cancel := context.WithTimeout(context.Background(), Timeout)
		defer cancel()
		if err := b.sender.Send(ctx, n); err != nil && b.log != nil {
			b.log.Warn("notify: the notice did not send", "err", err)
		}
	}()
}

// Wait blocks until every started send ends (tests, and a clean stop).
func (b *Background) Wait() { b.wg.Wait() }
