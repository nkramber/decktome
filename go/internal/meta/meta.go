// Package meta reads the published deck lists the quality model of
// PR-14B learns from (D-413 to D-417).
//
// The sources are the ones the owner allowed (D-5, D-417): the MTGO
// decklist pages, the MTGJSON deck products, EDHREC, the cEDH Decklist
// Database, and the Topdeck.gg API. Each reader turns one raw page or
// answer into List values, and the worker stores the raw page beside
// the lists. A parse that breaks on a markup change then re-reads the
// stored page and fetches nothing (roadmap PR-14B, M-6).
//
// The package resolves no card name. A List carries names, and the
// quality package resolves them against the card index at fit time.
package meta

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// UserAgent identifies this app to every source. MTGJSON answers 403 to
// a request with no agent (precon-data-2026-09-01).
const UserAgent = "mtg-deck-builder/0.1 (github.com/nkramber/decktome)"

// The tier words of the ladder (D-414). A great list is a top-8 finish
// or a competitive cEDH list. A good list is a league finish or the
// rest of a challenge. A typical list is the EDHREC average deck, a
// baseline list is a precon, and a bad list is synthetic.
const (
	TierGreat    = "great"
	TierGood     = "good"
	TierTypical  = "typical"
	TierBaseline = "baseline"
	TierBad      = "bad"
)

// Tiers lists the ladder, worst first. The ordinal fit reads the order.
var Tiers = []string{TierBad, TierBaseline, TierTypical, TierGood, TierGreat}

// TierLevel answers the rung of a tier, 0 for the worst and -1 for a
// word that is not a tier.
func TierLevel(tier string) int {
	for i, t := range Tiers {
		if t == tier {
			return i
		}
	}
	return -1
}

// The format words of a List. FormatSixty marks a 60-card product with
// no format of its own, which the Standard and the Modern fits both
// read as a baseline (D-414).
const (
	FormatCommander = "commander"
	FormatStandard  = "standard"
	FormatModern    = "modern"
	FormatSixty     = "sixty"
)

// FormatWord answers the format word of a format id, or "" for one the
// model does not cover.
func FormatWord(f mtgv1.FormatId) string {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		return FormatCommander
	case mtgv1.FormatId_FORMAT_ID_STANDARD:
		return FormatStandard
	case mtgv1.FormatId_FORMAT_ID_MODERN:
		return FormatModern
	default:
		return ""
	}
}

// FormatID answers the format id of a format word, or UNSPECIFIED.
func FormatID(word string) mtgv1.FormatId {
	switch word {
	case FormatCommander:
		return mtgv1.FormatId_FORMAT_ID_COMMANDER
	case FormatStandard:
		return mtgv1.FormatId_FORMAT_ID_STANDARD
	case FormatModern:
		return mtgv1.FormatId_FORMAT_ID_MODERN
	default:
		return mtgv1.FormatId_FORMAT_ID_UNSPECIFIED
	}
}

// Covers says whether a list of format word serves the fit of a format.
// A 60-card product serves Standard and Modern both.
func Covers(word string, f mtgv1.FormatId) bool {
	if word == FormatWord(f) {
		return true
	}
	return word == FormatSixty && (f == mtgv1.FormatId_FORMAT_ID_STANDARD || f == mtgv1.FormatId_FORMAT_ID_MODERN)
}

// The source words.
const (
	SourceMTGO    = "mtgo"
	SourceMTGJSON = "mtgjson"
	SourceEDHREC  = "edhrec"
	SourceCEDHDB  = "cedhdb"
	SourceTopdeck = "topdeck"
	// SourceMTGTop8 is the paper tournament site, and SourceGoldfish the
	// user decks of MTGGoldfish (PR-14C, D-502 to D-504).
	SourceMTGTop8  = "mtgtop8"
	SourceGoldfish = "mtggoldfish"
	// SourceSynthetic marks a bad list the engine made (D-414).
	SourceSynthetic = "synthetic"
)

// Card is one row of a list: the name as the source wrote it, and the
// count. The ids are set when the source gives them, and the fit
// resolves the rest by name.
type Card struct {
	Name       string `json:"name"`
	Count      int    `json:"count"`
	OracleID   string `json:"oracle_id,omitempty"`
	ScryfallID string `json:"scryfall_id,omitempty"`
}

// List is one published deck with its label.
type List struct {
	Source string `json:"source"`
	// ID is unique inside the source, for example the MTGO slug plus
	// the player login.
	ID     string `json:"id"`
	Format string `json:"format"`
	// Event names the tournament or the product.
	Event string `json:"event,omitempty"`
	// Date is the event or the release date, YYYY-MM-DD.
	Date string `json:"date"`
	// Players counts the field, 0 when unknown.
	Players int `json:"players,omitempty"`
	// Placement is the final rank, 0 when unknown.
	Placement int `json:"placement,omitempty"`
	Wins      int `json:"wins,omitempty"`
	Losses    int `json:"losses,omitempty"`
	// Tier is the label of D-414.
	Tier string `json:"tier"`
	// Defect names the broken axis of a synthetic list, empty otherwise.
	Defect     string   `json:"defect,omitempty"`
	Commanders []string `json:"commanders,omitempty"`
	Cards      []Card   `json:"cards"`
	Sideboard  []Card   `json:"sideboard,omitempty"`
}

// Key is the list's store key: source and id.
func (l *List) Key() string { return l.Source + ":" + l.ID }

// Size sums the main deck counts.
func (l *List) Size() int {
	n := 0
	for _, c := range l.Cards {
		n += c.Count
	}
	return n
}

// Month is the year and month of the list's date, "YYYY-MM", or "" when
// the date is malformed.
func (l *List) Month() string {
	if len(l.Date) < 7 {
		return ""
	}
	return l.Date[:7]
}

// Fetcher gets pages under a named agent, one host at a time, with a
// wait between two requests to the same host. Every source is a shared
// site with no rate published for a page reader, so the fetcher stays
// slow: the default is one request a second per host.
type Fetcher struct {
	http   *http.Client
	logger *slog.Logger
	gap    time.Duration
	// MaxBody bounds one answer. The default is DefaultMaxBody, and the
	// Topdeck.gg client raises it: a two-week window of cEDH tournaments
	// answered 43 MB on 2026-09-02.
	MaxBody int64

	mu   sync.Mutex
	last map[string]time.Time
}

// DefaultGap is the wait between two requests to one host.
const DefaultGap = time.Second

// NewFetcher makes a Fetcher. A nil client gets a 60 s timeout.
func NewFetcher(httpClient *http.Client, gap time.Duration, logger *slog.Logger) *Fetcher {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 60 * time.Second}
	}
	// A redirect is an answer, not a page: the MTGO site answers 302 to
	// the month listing for an event page it no longer serves, and a
	// reader that follows it stores the listing as the event.
	httpClient.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	if gap <= 0 {
		gap = DefaultGap
	}
	if logger == nil {
		logger = slog.Default()
	}
	return &Fetcher{http: httpClient, logger: logger, gap: gap, last: map[string]time.Time{}}
}

// Get fetches one URL. A status other than 200 is an error that names
// the status, so a caller can tell a 404 from a 403.
func (f *Fetcher) Get(ctx context.Context, rawURL string) ([]byte, error) {
	return f.do(ctx, http.MethodGet, rawURL, nil, nil)
}

// RetryWait is the wait before the one retry of a transport error or
// a 5xx. A first live run of 2026-09-02 lost a whole source to one
// connection reset after 557 good pages.
var RetryWait = 2 * time.Second

// do runs one request under the per-host gap. A transport error or a
// 5xx gets one retry after RetryWait, and a 4xx does not.
func (f *Fetcher) do(ctx context.Context, method, rawURL string, body io.Reader, headers map[string]string) ([]byte, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("meta: bad url %q: %w", rawURL, err)
	}
	var payload []byte
	if body != nil {
		if payload, err = io.ReadAll(body); err != nil {
			return nil, err
		}
	}
	var data []byte
	for attempt := 0; ; attempt++ {
		data, err = f.once(ctx, method, rawURL, u.Host, payload, headers)
		wait, ok := retryAfter(err)
		if err == nil || !ok || ctx.Err() != nil {
			return data, err
		}
		// A transport error or a 5xx retries once. A 429 retries up to
		// MaxRateLimitRetries times, each after the wait the site named:
		// the Topdeck.gg bulk endpoint answered two 429s in a row with
		// waits of 14 and 35 seconds on 2026-09-02.
		limit := 1
		if isRateLimit(err) {
			limit = MaxRateLimitRetries
		}
		if attempt >= limit {
			return nil, err
		}
		f.logger.Warn("meta fetch retries", "url", rawURL, "attempt", attempt+1, "wait", wait.String(), "err", err)
		t := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			t.Stop()
			return nil, ctx.Err()
		case <-t.C:
		}
	}
}

// MaxRateLimitRetries is how many times a 429 retries.
const MaxRateLimitRetries = 4

// isRateLimit says whether an error is a 429.
func isRateLimit(err error) bool {
	var se *StatusError
	return errors.As(err, &se) && se.Status == http.StatusTooManyRequests
}

// retryAfter answers the wait before the one retry, and false for an
// error no retry helps. A transport error waits RetryWait: a client
// timeout reads as a deadline error, so the caller's context decides
// and not the error. A 5xx waits the same. A 429 waits the Retry-After
// header, else RateLimitWait. A 4xx other than 429 gets no retry.
func retryAfter(err error) (time.Duration, bool) {
	var se *StatusError
	if !errors.As(err, &se) {
		return RetryWait, true
	}
	switch {
	case se.Status == http.StatusTooManyRequests:
		if se.RetryAfter > 0 {
			return min(se.RetryAfter, MaxRateLimitWait), true
		}
		return RateLimitWait, true
	case se.Status >= 500:
		return RetryWait, true
	default:
		return 0, false
	}
}

// once runs one request.
func (f *Fetcher) once(ctx context.Context, method, rawURL, host string, payload []byte, headers map[string]string) ([]byte, error) {
	if err := f.wait(ctx, host); err != nil {
		return nil, err
	}
	var body io.Reader
	if payload != nil {
		body = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, method, rawURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json, text/html;q=0.9, */*;q=0.5")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := f.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("meta: %s %s: %w", method, rawURL, err)
	}
	defer func() { _ = res.Body.Close() }()
	limit := f.MaxBody
	if limit <= 0 {
		limit = DefaultMaxBody
	}
	data, err := io.ReadAll(io.LimitReader(res.Body, limit+1))
	if err != nil {
		return nil, fmt.Errorf("meta: read %s: %w", rawURL, err)
	}
	if res.StatusCode != http.StatusOK {
		se := &StatusError{URL: rawURL, Status: res.StatusCode}
		if secs, err := strconv.Atoi(strings.TrimSpace(res.Header.Get("Retry-After"))); err == nil && secs > 0 {
			se.RetryAfter = time.Duration(secs) * time.Second
		}
		return nil, se
	}
	if int64(len(data)) > limit {
		return nil, fmt.Errorf("meta: %s answered more than %d bytes", rawURL, limit)
	}
	return data, nil
}

// DefaultMaxBody bounds one page. An MTGO event page is about 350 KB,
// and the cEDH database page is about 1.1 MB (read 2026-09-02).
const DefaultMaxBody = 16 << 20

// RateLimitWait is the wait after a 429 with no Retry-After header, and
// MaxRateLimitWait caps the header's value. Topdeck.gg limits by the
// minute (docs read 2026-09-02).
var (
	RateLimitWait    = 60 * time.Second
	MaxRateLimitWait = 5 * time.Minute
)

// wait sleeps until the host's gap has passed.
func (f *Fetcher) wait(ctx context.Context, host string) error {
	f.mu.Lock()
	next := f.last[host].Add(f.gap)
	now := time.Now()
	if next.Before(now) {
		next = now
	}
	f.last[host] = next
	f.mu.Unlock()
	d := time.Until(next)
	if d <= 0 {
		return nil
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// StatusError is a response with a status other than 200. RetryAfter
// is the header's wait on a 429, 0 when the header is absent.
type StatusError struct {
	URL        string
	Status     int
	RetryAfter time.Duration
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("meta: %s answered %d", e.URL, e.Status)
}

// NotFound says whether an error is a 404.
func NotFound(err error) bool {
	var se *StatusError
	return errors.As(err, &se) && se.Status == http.StatusNotFound
}

// cleanName trims a card name the way a source writes it: whitespace,
// and a set code in parentheses at the end.
func cleanName(s string) string {
	s = strings.TrimSpace(s)
	if i := strings.LastIndex(s, " ("); i > 0 && strings.HasSuffix(s, ")") {
		s = strings.TrimSpace(s[:i])
	}
	return s
}
