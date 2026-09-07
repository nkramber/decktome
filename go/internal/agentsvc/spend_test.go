package agentsvc

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/questions"
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

// capServer builds a server with a cap of 5 and the overrides given, for
// the checkSpendCap tests of D-576.
func capServer(t *testing.T, ledger Ledger, overrides map[string]float64) *Server {
	t.Helper()
	cat, err := questions.Load()
	if err != nil {
		t.Fatalf("catalog: %v", err)
	}
	client, _ := fakeClient(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	srv, err := New(cat, client, newFakeStore(), func(context.Context) string { return "u1" },
		WithLogger(quiet),
		WithClock(func() time.Time { return time.Unix(1000, 0).UTC() }),
		WithSpendCap(ledger, 5),
		WithSpendCapOverrides(overrides))
	if err != nil {
		t.Fatalf("server: %v", err)
	}
	return srv
}

// TestSpendCapOverrideTurnsTheCapOffForOneEmail is D-576: an override of
// zero lets the named caller past a spend over the cap of every user.
func TestSpendCapOverrideTurnsTheCapOffForOneEmail(t *testing.T) {
	ledger := &fakeLedger{spent: map[string]float64{"u1/1970-01": 500}}
	srv := capServer(t, ledger, map[string]float64{"owner@example.com": 0})
	ctx := auth.WithEmail(context.Background(), "owner@example.com")
	if err := srv.checkSpendCap(ctx, "u1"); err != nil {
		t.Fatalf("the override did not turn the cap off: %v", err)
	}
}

// TestSpendCapOverrideKeepsTheCapForEveryOtherEmail: an override on one
// email never lifts the cap of another caller.
func TestSpendCapOverrideKeepsTheCapForEveryOtherEmail(t *testing.T) {
	ledger := &fakeLedger{spent: map[string]float64{"u1/1970-01": 5}}
	srv := capServer(t, ledger, map[string]float64{"owner@example.com": 0})
	ctx := auth.WithEmail(context.Background(), "guest@example.com")
	err := srv.checkSpendCap(ctx, "u1")
	if connect.CodeOf(err) != connect.CodeResourceExhausted {
		t.Fatalf("a guest at the cap: %v", err)
	}
}

// TestSpendCapOverrideReadsItsOwnNumber: an override that names a number
// caps the caller at that number, over and under it.
func TestSpendCapOverrideReadsItsOwnNumber(t *testing.T) {
	srv := capServer(t, &fakeLedger{spent: map[string]float64{"u1/1970-01": 19.99}},
		map[string]float64{"big@example.com": 20})
	ctx := auth.WithEmail(context.Background(), "big@example.com")
	if err := srv.checkSpendCap(ctx, "u1"); err != nil {
		t.Fatalf("under the override: %v", err)
	}
	srv = capServer(t, &fakeLedger{spent: map[string]float64{"u1/1970-01": 20}},
		map[string]float64{"big@example.com": 20})
	if err := srv.checkSpendCap(ctx, "u1"); connect.CodeOf(err) != connect.CodeResourceExhausted {
		t.Fatalf("at the override: %v", err)
	}
}

// TestSpendCapOverrideIgnoresTheCaseOfTheEmail: the token can carry a
// capital letter, and the override still applies.
func TestSpendCapOverrideIgnoresTheCaseOfTheEmail(t *testing.T) {
	ledger := &fakeLedger{spent: map[string]float64{"u1/1970-01": 500}}
	srv := capServer(t, ledger, map[string]float64{"Owner@Example.com": 0})
	ctx := auth.WithEmail(context.Background(), "OWNER@example.COM")
	if err := srv.checkSpendCap(ctx, "u1"); err != nil {
		t.Fatalf("the override missed on case: %v", err)
	}
}

// TestSpendCapWithNoOverrideReadsTheCapOfEveryUser: a caller with no
// email, which is local mode, keeps the cap of every user.
func TestSpendCapWithNoOverrideReadsTheCapOfEveryUser(t *testing.T) {
	ledger := &fakeLedger{spent: map[string]float64{"u1/1970-01": 5}}
	srv := capServer(t, ledger, map[string]float64{"owner@example.com": 0})
	err := srv.checkSpendCap(context.Background(), "u1")
	if connect.CodeOf(err) != connect.CodeResourceExhausted {
		t.Fatalf("a caller with no email: %v", err)
	}
}
