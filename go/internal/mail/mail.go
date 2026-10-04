// Package mail sends one email through the Resend API (D-1077). The
// approval of a request for beta access and the link that proves an
// email use it (D-1081). The API key lives in Secret Manager, never in a
// file (D-639).
package mail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// ResendURL is the send endpoint of the Resend API. Read 2026-10-02 at
// resend.com/docs/api-reference/emails/send-email.
const ResendURL = "https://api.resend.com/emails"

// DefaultFrom is the sender. The root domain sends, and its inbox takes
// each reply (D-1112). MAIL_FROM overrides it.
const DefaultFrom = "Deck Tome <mail@decktome.com>"

// userAgent names the caller. Resend refuses a request with no
// User-Agent header with 403 (read 2026-10-02).
const userAgent = "decktome-api"

// Timeout bounds one send.
const Timeout = 10 * time.Second

// Message is one email to one address. The text is the source, and the
// send adds an HTML part with a real link for each URL of the text, so a
// mail app never decides whether a link works (D-1080).
type Message struct {
	To      string
	Subject string
	Text    string
}

// Sender delivers one email.
type Sender interface {
	Send(ctx context.Context, m Message) error
}

// Resend sends through the Resend API.
type Resend struct {
	key, from, url string
	client         *http.Client
}

// NewResend makes a sender for one API key and one sender address.
func NewResend(key, from string) *Resend {
	return &Resend{key: key, from: from, url: ResendURL, client: &http.Client{Timeout: Timeout}}
}

// WithURL points the sender at another endpoint (tests).
func (r *Resend) WithURL(u string) *Resend {
	r.url = u
	return r
}

// FromEnv reads RESEND_API_KEY and MAIL_FROM. It answers nil with no
// key, so a local run sends nothing.
func FromEnv(getenv func(string) string) *Resend {
	key := getenv("RESEND_API_KEY")
	if key == "" {
		return nil
	}
	from := getenv("MAIL_FROM")
	if from == "" {
		from = DefaultFrom
	}
	return NewResend(key, from)
}

// Send posts one email. A status other than 200 is an error, and the
// error never names the key.
func (r *Resend) Send(ctx context.Context, m Message) error {
	body, err := json.Marshal(map[string]any{
		"from":    r.from,
		"to":      []string{m.To},
		"subject": m.Subject,
		"text":    m.Text,
		"html":    HTML(m.Text),
	})
	if err != nil {
		return fmt.Errorf("mail: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("mail: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+r.key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("mail: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4<<10))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("mail: status %d", resp.StatusCode)
	}
	return nil
}

// link matches a URL of a text after the escape of HTML. The escape
// leaves no angle bracket and no quote, so a URL ends at a space.
var link = regexp.MustCompile(`https://[^\s]+`)

// HTML writes the HTML part of a text. A blank line starts a paragraph,
// a line break stays a line break, and each https URL becomes a link.
func HTML(text string) string {
	var b strings.Builder
	for _, para := range strings.Split(strings.TrimSpace(text), "\n\n") {
		lines := strings.Split(para, "\n")
		for i, line := range lines {
			lines[i] = link.ReplaceAllStringFunc(html.EscapeString(line), func(u string) string {
				return `<a href="` + u + `">` + u + `</a>`
			})
		}
		b.WriteString("<p>" + strings.Join(lines, "<br>") + "</p>\n")
	}
	return b.String()
}
