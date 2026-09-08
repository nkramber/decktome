package questions

import (
	"context"
	"io"
	"log/slog"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestTheReaderChoiceOfPoolWins is D-591. Session Wh7MdP2MucgpMl5IRTD5
// named a collection and "Only cards I own", and it read "Pool: any
// card". The reader asked for the LOTR and Hobbit sets, the classifier
// took that phrase for a pool answer, and apply wrote over the choice.
// The flip also turns the buy list on, so the deck names cards the
// reader does not own. It is the D-371 collision from the other side.
func TestTheReaderChoiceOfPoolWins(t *testing.T) {
	newAgent := func(hasCollection, fromReader bool, rule mtgv1.PoolRule) (*Agent, *State) {
		a := &Agent{log: slog.New(slog.NewTextHandler(io.Discard, nil))}
		st := NewState(hasCollection)
		st.Ctx.HasCollection = hasCollection
		st.Ctx.PoolFromReader = fromReader
		if rule != mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
			st.Slots.PoolRule = rule
			st.Close("pool_rule")
		}
		return a, st
	}

	// The defect, in one case: a set phrase must not move the pool.
	t.Run("a set phrase does not move the pool the reader chose", func(t *testing.T) {
		a, st := newAgent(true, true, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
		a.apply(context.Background(), st, classifyOut{PoolRule: "any_card"}, nil,
			"build a Commander deck from the LOTR and Hobbit sets only", nil)
		if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_OWNED_ONLY {
			t.Errorf("pool rule = %v, want OWNED_ONLY: the reader chose it", got)
		}
	})

	t.Run("no classifier rule moves a pool the reader chose", func(t *testing.T) {
		for _, word := range []string{"any_card", "owned_first", "owned_only"} {
			a, st := newAgent(true, true, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
			a.apply(context.Background(), st, classifyOut{PoolRule: word}, nil, "a message", nil)
			if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_OWNED_ONLY {
				t.Errorf("%q: pool rule = %v, want OWNED_ONLY", word, got)
			}
		}
	})

	// The guard covers the pool alone. Every later rule of apply must run
	// on the same turn, or a set phrase would also cost the reader the
	// power and the budget of that message.
	t.Run("the rest of the turn still applies", func(t *testing.T) {
		a, st := newAgent(true, true, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
		a.apply(context.Background(), st, classifyOut{PoolRule: "any_card", Theme: "elves"}, nil,
			"an elves deck from the LOTR sets only", nil)
		if st.Slots.GetTheme() == "" {
			t.Error("the theme did not apply, so the guard left the turn early")
		}
	})

	// D-371 still holds where the reader chose nothing.
	t.Run("the classifier still fills a pool the reader never chose", func(t *testing.T) {
		a, st := newAgent(true, false, mtgv1.PoolRule_POOL_RULE_UNSPECIFIED)
		a.apply(context.Background(), st, classifyOut{PoolRule: "owned_only"}, nil, "only cards I own", nil)
		if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_OWNED_ONLY {
			t.Errorf("pool rule = %v, want OWNED_ONLY", got)
		}
	})

	t.Run("an owned rule with no collection still reads as any card", func(t *testing.T) {
		a, st := newAgent(false, false, mtgv1.PoolRule_POOL_RULE_UNSPECIFIED)
		a.apply(context.Background(), st, classifyOut{PoolRule: "owned_only"}, nil, "build only from the Hobbit set", nil)
		if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
			t.Errorf("pool rule = %v, want ANY_CARD", got)
		}
	})

	// A stored snapshot written before D-591 holds no pool_from_reader,
	// and it reads false. Those sessions keep the rule before this one.
	t.Run("a snapshot without the fact reads the rule before it", func(t *testing.T) {
		snap := Snapshot{Version: SnapshotVersion}
		st := Restore("s1", nil, snap)
		if st.Ctx.PoolFromReader {
			t.Error("PoolFromReader must default to false")
		}
	})

	// The fact survives the store, or the second turn of a chat loses the
	// choice that the first one carried.
	t.Run("the fact survives a snapshot and a restore", func(t *testing.T) {
		_, st := newAgent(true, true, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
		back := Restore("s1", st.Slots, st.Snapshot())
		if !back.Ctx.PoolFromReader {
			t.Error("PoolFromReader did not survive the round trip")
		}
	})
}
