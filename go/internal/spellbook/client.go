// Package spellbook calls the Commander Spellbook bracket endpoint.
//
// The endpoint takes a deck list and answers a flag per card for Game
// Changer, mass land denial, and extra turn, and a flag per combo for
// two-card, speed, lock, and extra turn (roadmap PR-14A). The owner read
// the terms, and they allow the call at 90 requests a minute, limited on
// this side (D-459). Every request sends a named agent.
//
// The endpoint is the authority on the content rules of a bracket. The
// profile reads its answer, and no prompt does.
package spellbook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// UserAgent identifies this app to Commander Spellbook.
const UserAgent = "mtg-deck-builder/0.1 (github.com/nkramber/mtg-deck-builder)"

// DefaultBaseURL is the production backend. Its OpenAPI schema is at
// /schema/?format=json, read 2026-09-02.
const DefaultBaseURL = "https://backend.commanderspellbook.com"

// RequestsPerMinute is the rate the owner set (D-459).
const RequestsPerMinute = 90

// DefaultRetryAfter is the wait after a 429 with no Retry-After header.
const DefaultRetryAfter = 30 * time.Second

// MaxMain is the largest main list the endpoint takes, from its schema.
const MaxMain = 600

// Client calls the endpoint under the rate limit.
type Client struct {
	http    *http.Client
	baseURL string
	logger  *slog.Logger
	limiter *limiter
	// retryAfter is the wait after a 429 with no Retry-After header.
	// Tests shorten it.
	retryAfter time.Duration
}

// New returns a Client. A nil httpClient gets a transport with a 30 s
// header timeout. An empty baseURL is the production backend.
func New(httpClient *http.Client, baseURL string, logger *slog.Logger) *Client {
	if httpClient == nil {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.ResponseHeaderTimeout = 30 * time.Second
		httpClient = &http.Client{Transport: transport}
	}
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	if logger == nil {
		logger = slog.Default()
	}
	return &Client{
		http:       httpClient,
		baseURL:    baseURL,
		logger:     logger,
		limiter:    newLimiter(time.Minute / RequestsPerMinute),
		retryAfter: DefaultRetryAfter,
	}
}

// Result is the endpoint's answer for one deck.
type Result struct {
	// BracketTag is the endpoint's own bracket word: E exhibition, C
	// core, O oddball, P powerful, S spicy, R ruthless, B banned.
	BracketTag string            `json:"bracketTag"`
	Cards      []ClassifiedCard  `json:"cards"`
	Combos     []ClassifiedCombo `json:"combos"`
}

// ClassifiedCard is one card the endpoint knew, with its flags.
type ClassifiedCard struct {
	Card           CardRef `json:"card"`
	Quantity       int     `json:"quantity"`
	Banned         bool    `json:"banned"`
	GameChanger    bool    `json:"gameChanger"`
	MassLandDenial bool    `json:"massLandDenial"`
	ExtraTurn      bool    `json:"extraTurn"`
}

// CardRef names a card the endpoint knows. OracleID can be empty.
type CardRef struct {
	Name     string `json:"name"`
	OracleID string `json:"oracleId"`
}

// ClassifiedCombo is one combo the deck holds, with the endpoint's
// reading of it. Speed runs 1 to 6: 5 needs no mana, 4 needs four or
// less, 3 six or less, 2 eight or less, 1 more, and an uncertain
// minimum adds one (variant.py, read 2026-09-02).
type ClassifiedCombo struct {
	Combo               ComboRef `json:"combo"`
	Relevant            bool     `json:"relevant"`
	BorderlineRelevant  bool     `json:"borderlineRelevant"`
	ArguablyTwoCard     bool     `json:"arguablyTwoCard"`
	DefinitelyTwoCard   bool     `json:"definitelyTwoCard"`
	Speed               int      `json:"speed"`
	MassLandDenial      bool     `json:"massLandDenial"`
	ExtraTurn           bool     `json:"extraTurn"`
	Lock                bool     `json:"lock"`
	SkipTurns           bool     `json:"skipTurns"`
	ControlAllOpponents bool     `json:"controlAllOpponents"`
}

// ComboRef is the combo the flags describe: its id, its bracket word,
// and the cards it uses.
type ComboRef struct {
	ID         string     `json:"id"`
	BracketTag string     `json:"bracketTag"`
	Uses       []ComboUse `json:"uses"`
}

// ComboUse is one card of a combo.
type ComboUse struct {
	Card CardRef `json:"card"`
}

// Names lists the cards of a combo, in the endpoint's order.
func (c ComboRef) Names() []string {
	out := make([]string, 0, len(c.Uses))
	for _, u := range c.Uses {
		out = append(out, u.Card.Name)
	}
	return out
}

// BracketOf maps the endpoint's tag onto a bracket number, the way its
// own generated field does: R 4, S 3, P 3, O 2, C 2, E 1. A banned or
// unknown tag gives 0 (variant.py, read 2026-09-02).
func BracketOf(tag string) int32 {
	switch tag {
	case "R":
		return 4
	case "S", "P":
		return 3
	case "O", "C":
		return 2
	case "E":
		return 1
	}
	return 0
}

type cardIn struct {
	Card     string `json:"card"`
	Quantity int    `json:"quantity,omitempty"`
}

type deckIn struct {
	Commanders []cardIn `json:"commanders"`
	Main       []cardIn `json:"main"`
}

// EstimateBracket sends one deck and returns the endpoint's reading.
// commanders and main are card names, one entry per copy or one entry
// with a count folded by the caller. A main list over MaxMain is an
// error before any call.
func (c *Client) EstimateBracket(ctx context.Context, commanders, main []string) (*Result, error) {
	if len(main) > MaxMain {
		return nil, fmt.Errorf("spellbook: %d cards is over the %d the endpoint takes", len(main), MaxMain)
	}
	body := deckIn{Commanders: []cardIn{}, Main: []cardIn{}}
	for _, n := range commanders {
		body.Commanders = append(body.Commanders, cardIn{Card: n})
	}
	for _, n := range main {
		body.Main = append(body.Main, cardIn{Card: n})
	}
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("spellbook: %w", err)
	}
	if err := c.limiter.wait(ctx); err != nil {
		return nil, fmt.Errorf("spellbook: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/estimate-bracket", bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("spellbook: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.do(ctx, req, raw)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()
	var out Result
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("spellbook: estimate-bracket answer: %w", err)
	}
	return &out, nil
}

// do sends one request. On a 429 it waits for Retry-After, or the
// default wait, then retries once, as the Scryfall client does.
func (c *Client) do(ctx context.Context, req *http.Request, body []byte) (*http.Response, error) {
	req.Header.Set("User-Agent", UserAgent)
	for attempt := 0; ; attempt++ {
		r := req.Clone(ctx)
		r.Body = io.NopCloser(bytes.NewReader(body))
		res, err := c.http.Do(r)
		if err != nil {
			return nil, fmt.Errorf("spellbook: estimate-bracket: %w", err)
		}
		if res.StatusCode == http.StatusTooManyRequests && attempt == 0 {
			wait := c.retryWait(res.Header.Get("Retry-After"))
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
			c.logger.Warn("spellbook rate limited, retry once", "wait", wait.String())
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("spellbook: %w", ctx.Err())
			case <-time.After(wait):
			}
			continue
		}
		if res.StatusCode != http.StatusOK {
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
			return nil, fmt.Errorf("spellbook: estimate-bracket: status %d", res.StatusCode)
		}
		return res, nil
	}
}

// retryWait reads a Retry-After header: a delay in seconds, or an
// HTTP-date. Neither form gives the default wait.
func (c *Client) retryWait(header string) time.Duration {
	if secs, err := strconv.Atoi(header); err == nil && secs > 0 {
		return time.Duration(secs) * time.Second
	}
	if at, err := http.ParseTime(header); err == nil {
		if wait := time.Until(at); wait > 0 {
			return wait
		}
	}
	return c.retryAfter
}

// limiter spaces calls by one interval. It holds the rate with no burst,
// so a run of builds never sends more than the owner allowed (D-459).
type limiter struct {
	mu       sync.Mutex
	interval time.Duration
	next     time.Time
	now      func() time.Time
	sleep    func(context.Context, time.Duration) error
}

func newLimiter(interval time.Duration) *limiter {
	return &limiter{
		interval: interval,
		now:      time.Now,
		sleep: func(ctx context.Context, d time.Duration) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(d):
				return nil
			}
		},
	}
}

// wait blocks until the next slot, then takes it.
func (l *limiter) wait(ctx context.Context) error {
	l.mu.Lock()
	now := l.now()
	at := l.next
	if at.Before(now) {
		at = now
	}
	l.next = at.Add(l.interval)
	l.mu.Unlock()
	if d := at.Sub(now); d > 0 {
		return l.sleep(ctx, d)
	}
	return nil
}
