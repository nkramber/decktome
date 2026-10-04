package deckreads

import (
	"testing"
	"time"
)

var t0 = time.Date(2026, 10, 3, 12, 0, 0, 0, time.UTC)

// TestMatchesNeedsTheUserTheDeckAndTheText is D-1107: a link stays only
// after a read of the same deck and the same text by the same user, less
// than Life ago.
func TestMatchesNeedsTheUserTheDeckAndTheText(t *testing.T) {
	read := New("u1", 42, "Deck\n4 Lightning Bolt\n", t0)
	for _, tc := range []struct {
		name string
		uid  string
		deck int64
		text string
		at   time.Time
		want bool
	}{
		{"the read", "u1", 42, "Deck\n4 Lightning Bolt\n", t0.Add(time.Minute), true},
		{"another user", "u2", 42, "Deck\n4 Lightning Bolt\n", t0.Add(time.Minute), false},
		{"another deck", "u1", 43, "Deck\n4 Lightning Bolt\n", t0.Add(time.Minute), false},
		{"another text", "u1", 42, "Deck\n4 Shock\n", t0.Add(time.Minute), false},
		{"an old read", "u1", 42, "Deck\n4 Lightning Bolt\n", t0.Add(Life), false},
	} {
		if got := read.Matches(tc.uid, tc.deck, tc.text, tc.at); got != tc.want {
			t.Errorf("%s: Matches = %v, want %v", tc.name, got, tc.want)
		}
	}
	if !read.ExpireAt.Equal(t0.Add(Life)) {
		t.Errorf("expire_at = %v, want the read time and Life", read.ExpireAt)
	}
}

// TestEachReadHasItsOwnKey: a read of another user, deck, or text never
// replaces a read, so no count of reads drops one before its hour ends.
func TestEachReadHasItsOwnKey(t *testing.T) {
	keys := map[string]bool{}
	for _, k := range []string{
		Key("u1", 42, "a"), Key("u2", 42, "a"), Key("u1", 43, "a"), Key("u1", 42, "b"), Key("u14", 2, "a"),
	} {
		if keys[k] {
			t.Fatalf("key %s repeats", k)
		}
		keys[k] = true
	}
	first, again := Key("u1", 42, "a"), Key("u1", 42, "a")
	if first != again {
		t.Error("the same read gets two keys")
	}
}
