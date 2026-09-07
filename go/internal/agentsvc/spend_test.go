package agentsvc

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// fakeLedger is the spend ledger of the tests: one total per user and
// month, and the adds it received.
type fakeLedger struct {
	mu    sync.Mutex
	spent map[string]float64
	adds  []fakeAdd
	err   error
}

type fakeAdd struct {
	uid, month string
	cost       float64
	calls      int64
}

func (f *fakeLedger) Spent(_ context.Context, uid, month string) (float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.err != nil {
		return 0, f.err
	}
	return f.spent[uid+"/"+month], nil
}

func (f *fakeLedger) Add(_ context.Context, uid, month string, cost float64, calls int64, _ time.Time) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.adds = append(f.adds, fakeAdd{uid, month, cost, calls})
	if f.spent == nil {
		f.spent = map[string]float64{}
	}
	f.spent[uid+"/"+month] += cost
	return nil
}

// TestSpendCapRefusesATurnAtTheCap is D-421: a user at the cap gets
// ResourceExhausted, and the refusal names the day the cap resets. The
// test clock is 1970-01-01, so the reset is 1970-02-01.
func TestSpendCapRefusesATurnAtTheCap(t *testing.T) {
	ledger := &fakeLedger{spent: map[string]float64{"u1/1970-01": 5}}
	store := newFakeStore()
	client, _ := testServerOpts(t, store, []Option{WithSpendCap(ledger, 5)}, firstTurn(t)...)
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"}))
	if err != nil {
		t.Fatalf("chat: %v", err)
	}
	for stream.Receive() {
	}
	err = stream.Err()
	if connect.CodeOf(err) != connect.CodeResourceExhausted || !strings.Contains(err.Error(), "resets on 1970-02-01") {
		t.Fatalf("a turn at the cap: %v", err)
	}
	if len(ledger.adds) != 0 {
		t.Errorf("a refused turn added to the ledger: %+v", ledger.adds)
	}
	if len(store.sessions) != 0 {
		t.Errorf("a refused turn stored a session: %d", len(store.sessions))
	}
}

// TestSpendCapRecordsTheTurn: a turn under the cap runs, and the ledger
// gets its calls and its cost once the turn ends. A ledger that can not
// be read refuses the turn with Unavailable.
func TestSpendCapRecordsTheTurn(t *testing.T) {
	ledger := &fakeLedger{spent: map[string]float64{"u1/1970-01": 4.99}}
	store := newFakeStore()
	client, _ := testServerOpts(t, store, []Option{WithSpendCap(ledger, 5)}, firstTurn(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	if first.started == "" {
		t.Fatal("the turn under the cap did not run")
	}
	if len(ledger.adds) != 1 || ledger.adds[0].uid != "u1" || ledger.adds[0].month != "1970-01" || ledger.adds[0].calls == 0 {
		t.Fatalf("adds = %+v, want one add for u1 in 1970-01 with the turn's calls", ledger.adds)
	}

	down := &fakeLedger{err: errors.New("firestore down")}
	client, _ = testServerOpts(t, newFakeStore(), []Option{WithSpendCap(down, 5)}, firstTurn(t)...)
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"}))
	if err != nil {
		t.Fatalf("chat: %v", err)
	}
	for stream.Receive() {
	}
	if connect.CodeOf(stream.Err()) != connect.CodeUnavailable {
		t.Errorf("a ledger that can not be read: %v", stream.Err())
	}
}

// TestNoCapWithoutALedger: WithSpendCap with a nil ledger or a zero cap
// sets nothing, and every turn runs.
func TestNoCapWithoutALedger(t *testing.T) {
	store := newFakeStore()
	client, _ := testServerOpts(t, store, []Option{WithSpendCap(nil, 5), WithSpendCap(&fakeLedger{}, 0)}, firstTurn(t)...)
	if first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"}); first.started == "" {
		t.Fatal("the turn did not run")
	}
}
