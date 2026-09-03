// Command chat-probe drives the real Chat RPC from a first message to a
// finished deck, through the whole production path.
//
// The deck gate calls internal/generate directly, so it never runs
// agentsvc. This probe covers that gap: a fault in the path from the RPC
// to the build shows here and nowhere else (D-232).
//
// CAUTION: this calls the real providers and it costs money. One session
// is a few classify and ask calls plus one generate call. CHAT_PROBE=1 is
// required, so it can not run by accident.
//
// Usage:
//
//	CHAT_PROBE=1 CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/chat-probe -messages "Build me a lifegain Commander deck.|Karlov of the Ghost Council, bracket 3."
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/agentsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// memStore is the session store, in memory. The probe writes nothing.
type memStore struct {
	mu    sync.Mutex
	sess  map[string]*mtgv1.Session
	snaps map[string]questions.Snapshot
	vers  map[string]int64
	n     int
}

func newMemStore() *memStore {
	return &memStore{sess: map[string]*mtgv1.Session{}, snaps: map[string]questions.Snapshot{}, vers: map[string]int64{}}
}

func (m *memStore) NewID(string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.n++
	return fmt.Sprintf("probe-%d", m.n)
}

func (m *memStore) Put(_ context.Context, _ string, s *mtgv1.Session, snap questions.Snapshot, expected int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.vers[s.GetId()] != expected {
		return fmt.Errorf("version %d, expected %d", m.vers[s.GetId()], expected)
	}
	m.sess[s.GetId()], m.snaps[s.GetId()] = s, snap
	m.vers[s.GetId()] = expected + 1
	return nil
}

func (m *memStore) Get(_ context.Context, _, id string) (*mtgv1.Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.sess[id], nil
}

func (m *memStore) GetState(_ context.Context, _, id string) (*mtgv1.Session, questions.Snapshot, int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.sess[id], m.snaps[id], m.vers[id], nil
}

// indexSrc is the loaded snapshot.
type indexSrc struct{ idx *cards.Index }

func (i indexSrc) Current() *cards.Index { return i.idx }

// ownedSrc answers the owned counts of one collection.
type ownedSrc struct{ counts map[string]int32 }

func (o ownedSrc) OwnedPrintings(context.Context, string, string) (map[string][]string, error) {
	return nil, nil
}

func (o ownedSrc) OracleCounts(context.Context, string, string) (map[string]int32, error) {
	return o.counts, nil
}

func (o ownedSrc) PrintingCounts(context.Context, string, string) (map[string]int32, error) {
	return nil, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	collPath := flag.String("collection", "", "a ManaBox CSV, which puts the session in an owned mode")
	msgs := flag.String("messages", "Build me a lifegain Commander deck from any cards.|Karlov of the Ghost Council. Bracket 3, white and black, and no budget.", "the user's turns, separated by |")
	flag.Parse()

	if err := gatekit.SpendGuard("CHAT_PROBE"); err != nil {
		return err
	}
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	cat, err := questions.Load()
	if err != nil {
		return err
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	rcfg, err := rules.Load()
	if err != nil {
		return err
	}
	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	prof, err := gatekit.Profiler(idx, rcfg, quiet)
	if err != nil {
		return err
	}
	// The quality model grades every deck the gate builds, and the
	// summary names the tier (PR-14B). No stored model grades nothing.
	scorer, err := gatekit.Scorer(context.Background())
	if err != nil {
		return err
	}
	opts := []agentsvc.Option{
		agentsvc.WithLogger(quiet),
		agentsvc.WithCandidates(indexSrc{idx}, cb),
		agentsvc.WithDecks(generate.NewBuilder(client, rcfg, idx, quiet, generate.WithProfiler(prof), generate.WithScorer(scorer))),
		agentsvc.WithScorer(scorer),
	}
	if *collPath != "" {
		owned, _, err := gatekit.LoadOwned(*collPath, idx)
		if err != nil {
			return err
		}
		fmt.Printf("collection: %d owned cards\n", len(owned))
		opts = append(opts, agentsvc.WithCollections(ownedSrc{owned}))
	}
	srv, err := agentsvc.New(cat, client, newMemStore(),
		func(context.Context) string { return "probe-user" }, opts...)
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewAgentServiceHandler(srv))
	hs := httptest.NewServer(mux)
	defer hs.Close()
	c := mtgv1connect.NewAgentServiceClient(hs.Client(), hs.URL)

	var sessionID string
	var deck *mtgv1.Deck
	start := time.Now()
	for i, m := range strings.Split(*msgs, "|") {
		m = strings.TrimSpace(m)
		fmt.Printf("\n--- turn %d: %q\n", i+1, m)
		req := &mtgv1.ChatRequest{Message: m}
		if sessionID != "" {
			req.SessionId = sessionID
		}
		if *collPath != "" && sessionID == "" {
			req.CollectionId = "probe-collection"
		}
		stream, err := c.Chat(context.Background(), connect.NewRequest(req))
		if err != nil {
			return err
		}
		for stream.Receive() {
			switch e := stream.Msg().GetEvent().(type) {
			case *mtgv1.ChatResponse_SessionStarted:
				sessionID = e.SessionStarted
			case *mtgv1.ChatResponse_Question:
				fmt.Printf("  Q [%s] %s\n", e.Question.GetSlot(), e.Question.GetText())
			case *mtgv1.ChatResponse_Status:
				fmt.Printf("  status: %s\n", e.Status)
			case *mtgv1.ChatResponse_TextDelta:
				fmt.Printf("  text: %s\n", e.TextDelta)
			case *mtgv1.ChatResponse_Deck:
				deck = e.Deck
			case *mtgv1.ChatResponse_Slots:
				var open []string
				for k, v := range e.Slots.GetSlotStates() {
					if v == mtgv1.SlotState_SLOT_STATE_ASKED {
						open = append(open, k)
					}
				}
				sort.Strings(open)
				fmt.Printf("  slots: %d asked and unanswered: %v\n", len(open), open)
			case *mtgv1.ChatResponse_Failure:
				fmt.Printf("  FAILURE: %s\n", e.Failure.GetMessage())
			}
		}
		if err := stream.Err(); err != nil {
			return fmt.Errorf("turn %d: %w", i+1, err)
		}
		_ = stream.Close()
	}

	if deck == nil {
		return fmt.Errorf("no deck reached the user after %d turns", len(strings.Split(*msgs, "|")))
	}
	report(deck, time.Since(start))
	return nil
}

func report(d *mtgv1.Deck, took time.Duration) {
	fmt.Printf("\n=== the deck reached the user ===\n")
	fmt.Printf("format: %v. Commanders: %d. Cards: %d main, %d sideboard.\n",
		d.GetFormat().GetId(), len(d.GetCommanderOracleIds()), gatekit.CountCards(d), gatekit.CountSideboard(d))
	fmt.Printf("cost: $%.2f to buy, $%.2f the whole deck.\n", generate.BuyCost(d), generate.DeckCost(d))
	fmt.Printf("summary: %s\n", d.GetSummary())
	for _, f := range d.GetValidation().GetFindings() {
		fmt.Printf("  [%s] %s: %s\n", strings.TrimPrefix(f.GetSeverity().String(), "SEVERITY_"), f.GetCode(), f.GetMessage())
	}
	fmt.Printf("block findings: %d. Time: %.0fs\n", len(gatekit.BlockFindings(d)), took.Seconds())
	if claims := generate.LintSummary(d.GetSummary()); len(claims) > 0 {
		fmt.Printf("summary rules claims (F-26): %v\n", claims)
	}
}

// The sessions list of D-433 is not part of a probe. The store answers
// what it holds, so the interface is met.
func (m *memStore) List(_ context.Context, _ string) ([]*mtgv1.SessionSummary, error) {
	out := make([]*mtgv1.SessionSummary, 0, len(m.sess))
	for _, s := range m.sess {
		out = append(out, sessions.Summarize(s))
	}
	return out, nil
}

func (m *memStore) Rename(_ context.Context, _, id, name string) (*mtgv1.SessionSummary, error) {
	s, ok := m.sess[id]
	if !ok {
		return nil, sessions.ErrNotFound
	}
	s.Name = name
	return sessions.Summarize(s), nil
}

func (m *memStore) Delete(_ context.Context, _, id string) error {
	if _, ok := m.sess[id]; !ok {
		return sessions.ErrNotFound
	}
	delete(m.sess, id)
	return nil
}
