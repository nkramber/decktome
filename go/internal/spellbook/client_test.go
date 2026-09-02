package spellbook

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

const answer = `{"bracketTag":"R","cards":[
 {"card":{"name":"Armageddon","oracleId":"c9ed8b01"},"quantity":1,"banned":false,"gameChanger":false,"massLandDenial":true,"extraTurn":false},
 {"card":{"name":"Demonic Tutor","oracleId":"82004860"},"quantity":1,"banned":false,"gameChanger":true,"massLandDenial":false,"extraTurn":false}],
 "templates":[],
 "combos":[{"combo":{"id":"742-1295","bracketTag":"R","uses":[{"card":{"name":"Demonic Consultation"}},{"card":{"name":"Thassa's Oracle"}}]},
  "relevant":true,"borderlineRelevant":true,"arguablyTwoCard":true,"definitelyTwoCard":true,"speed":5,
  "massLandDenial":false,"extraTurn":false,"lock":false,"skipTurns":false,"controlAllOpponents":false,"controlSomeOpponents":false}]}`

func TestEstimateBracketSendsTheDeckAndReadsTheFlags(t *testing.T) {
	var got deckIn
	var ua string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/estimate-bracket" || r.Method != http.MethodPost {
			t.Errorf("request %s %s", r.Method, r.URL.Path)
		}
		ua = r.Header.Get("User-Agent")
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Errorf("body: %v", err)
		}
		_, _ = w.Write([]byte(answer))
	}))
	defer srv.Close()
	c := New(srv.Client(), srv.URL, nil)
	res, err := c.EstimateBracket(context.Background(), []string{"Ezuri, Renegade Leader"}, []string{"Armageddon", "Demonic Tutor"})
	if err != nil {
		t.Fatal(err)
	}
	if ua != UserAgent {
		t.Errorf("user agent %q", ua)
	}
	if len(got.Commanders) != 1 || got.Commanders[0].Card != "Ezuri, Renegade Leader" || len(got.Main) != 2 {
		t.Errorf("sent %+v", got)
	}
	if res.BracketTag != "R" || len(res.Cards) != 2 || !res.Cards[0].MassLandDenial || !res.Cards[1].GameChanger {
		t.Errorf("cards %+v", res.Cards)
	}
	if len(res.Combos) != 1 || !res.Combos[0].DefinitelyTwoCard || res.Combos[0].Speed != 5 {
		t.Errorf("combos %+v", res.Combos)
	}
	if names := res.Combos[0].Combo.Names(); len(names) != 2 || names[1] != "Thassa's Oracle" {
		t.Errorf("combo names %v", names)
	}
}

func TestEstimateBracketRetriesOnceAfter429(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if calls.Add(1) == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		_, _ = w.Write([]byte(answer))
	}))
	defer srv.Close()
	c := New(srv.Client(), srv.URL, nil)
	c.retryAfter = time.Millisecond
	if _, err := c.EstimateBracket(context.Background(), nil, []string{"Armageddon"}); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 2 {
		t.Errorf("calls %d, want 2", calls.Load())
	}
}

func TestEstimateBracketRefusesAStatusAndAnOversizedList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer srv.Close()
	c := New(srv.Client(), srv.URL, nil)
	if _, err := c.EstimateBracket(context.Background(), nil, []string{"Armageddon"}); err == nil {
		t.Error("a 502 must be an error")
	}
	if _, err := c.EstimateBracket(context.Background(), nil, make([]string, MaxMain+1)); err == nil {
		t.Error("a list over MaxMain must be refused before the call")
	}
}

func TestLimiterSpacesCallsByTheInterval(t *testing.T) {
	now := time.Unix(1000, 0)
	var slept []time.Duration
	l := newLimiter(time.Second)
	l.now = func() time.Time { return now }
	l.sleep = func(_ context.Context, d time.Duration) error {
		slept = append(slept, d)
		now = now.Add(d)
		return nil
	}
	for range 3 {
		if err := l.wait(context.Background()); err != nil {
			t.Fatal(err)
		}
	}
	// The first call goes at once, and each next one waits one interval.
	if len(slept) != 2 || slept[0] != time.Second || slept[1] != time.Second {
		t.Errorf("sleeps %v, want two of one second", slept)
	}
	// After a long pause the next call goes at once.
	now = now.Add(time.Minute)
	if err := l.wait(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(slept) != 2 {
		t.Errorf("a call after a pause must not sleep, got %v", slept)
	}
}

func TestBracketOf(t *testing.T) {
	want := map[string]int32{"R": 4, "S": 3, "P": 3, "O": 2, "C": 2, "E": 1, "B": 0, "": 0}
	for tag, n := range want {
		if got := BracketOf(tag); got != n {
			t.Errorf("BracketOf(%q) = %d, want %d", tag, got, n)
		}
	}
}
