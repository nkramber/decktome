package questions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// TestLiveConversation runs one whole question phase against the real
// providers: the words one user writes to build one deck. It runs only
// with QUESTIONS_LIVE=1 and the API keys in the environment.
//
// The test prints every question with its source and fit score, then the
// usage report. The owner approved one run of this size on 2026-08-24, to
// measure the token cost of a turn.
func TestLiveConversation(t *testing.T) {
	if os.Getenv("QUESTIONS_LIVE") != "1" {
		t.Skip("set QUESTIONS_LIVE=1 and the API keys to run the live conversation")
	}
	env := func(k string) string {
		if k == llm.EnvRequireKeys {
			return "1"
		}
		return os.Getenv(k)
	}
	client, err := llm.NewFromEnv(env, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		t.Fatalf("prices: %v", err)
	}
	opts := []AgentOption{WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))}
	// With a local snapshot the hints are real: PR-6 names the theme's
	// colors and its commanders, so the placeholder clauses resolve
	// instead of being dropped.
	if dir := os.Getenv("CARDS_SNAPSHOT_DIR"); dir != "" {
		idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, slog.New(slog.NewTextHandler(io.Discard, nil)))
		if err != nil {
			t.Fatalf("index: %v", err)
		}
		builder, err := candidates.New()
		if err != nil {
			t.Fatalf("candidates: %v", err)
		}
		opts = append(opts, WithHints(&CandidateHints{
			Index: idx, Builder: builder,
			Format: mtgv1.FormatId_FORMAT_ID_COMMANDER,
			Log:    slog.New(slog.NewTextHandler(io.Discard, nil)),
		}))
		t.Log("hints: PR-6 index loaded")
	}
	agent, err := NewAgent(load(t), client, opts...)
	if err != nil {
		t.Fatalf("agent: %v", err)
	}

	// One user, one deck. The messages answer the questions a real turn
	// would ask, in the order the planner asks them.
	messages := []string{
		"Build me a lifegain deck. I have a collection.",
		"Commander, and white and black is right.",
		"Bracket 3. Suggest a commander from my library, and build from my library first.",
		"Karlov of the Ghost Council. Keep the buy list under 50 dollars.",
	}
	acc := llm.NewAccumulator(prices)
	st := NewState(true)
	for i, msg := range messages {
		res, err := agent.Turn(context.Background(), st, msg, acc)
		if err != nil {
			t.Fatalf("turn %d: %v", i+1, err)
		}
		t.Logf("turn %d user: %s", i+1, msg)
		for _, q := range res.Questions {
			source := "catalog"
			if q.Invented {
				source = "INVENTED"
			}
			t.Logf("   [%s slot=%s fit=%.2f] %s", source, q.Slot, q.GapScore, q.Text)
		}
		if res.Ready {
			t.Logf("   slots complete after %d turns", i+1)
			break
		}
	}
	t.Logf("ready=%v outstanding=%v slot_states=%v", st.Ready(load(t)), st.Outstanding(), st.Slots.GetSlotStates())
	rep := acc.Report()
	raw, _ := json.MarshalIndent(rep, "", "  ")
	t.Logf("usage report:\n%s", raw)
	fmt.Printf("LIVE USAGE: %s\n", raw)
	if rep.Calls == 0 {
		t.Fatal("no calls were made")
	}
}
