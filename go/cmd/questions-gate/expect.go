package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nkramber/decktome/go/internal/questions"
)

// The golden expectations of PR-15 (slice 4). A conversation may name
// the values its slots must end with, and the gate reads each one as a
// row. A wrong slot is then a named flip, and a counted conversation
// with a miss fails the gate. The vocabulary is the one slotValues
// writes, and TestSlotValuesWriteTheVocabulary pins it.
//
// Three words are special. "*" asks for any value, for a pick the pool
// made. "delegated" on the commander asks for a skipped slot with no
// commander, which is the pick the user handed to the agent (D-232).
// "none" asks for an empty slot, which is how a colorless deck ends: the
// color slot closes with no color (D-165). A theme matches when the
// expected words are inside the slot's text, and a budget with no scope
// word matches any scope.

// expectKeys are the keys an expectation may name, in report order.
var expectKeys = []string{"format", "colors", "power", "pool_rule", "commander", "budget", "theme", "locked", "sets"}

// slotValues renders the settled slots of a conversation as words. The
// vocabulary lives in questions.SlotWords, so the gate and the triage of
// PR-28b read one renderer (D-643).
func slotValues(st *questions.State) map[string]string {
	if st == nil {
		return map[string]string{}
	}
	return questions.SlotWords(st.Slots, st.CommanderNames, st.LockedCards(), st.Ctx.HasCollection)
}

// checkExpect reads the values against the expectation and names every
// miss as "key: want X, got Y", in key order.
func checkExpect(expect, values map[string]string) []string {
	var misses []string
	keys := make([]string, 0, len(expect))
	for k := range expect {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		want, got := expect[key], values[key]
		if !expectMatches(key, want, got) {
			misses = append(misses, fmt.Sprintf("%s: want %s, got %s", key, orNone(want), orNone(got)))
		}
	}
	return misses
}

func expectMatches(key, want, got string) bool {
	w, g := strings.ToLower(strings.TrimSpace(want)), strings.ToLower(strings.TrimSpace(got))
	switch {
	case w == "*":
		return g != "" && g != "delegated"
	case w == "none":
		return g == ""
	case key == "theme":
		return w != "" && strings.Contains(g, w)
	case key == "budget":
		if g == w {
			return true
		}
		// A number with no scope word matches any scope.
		return !strings.Contains(w, " ") && strings.HasPrefix(g, w+" ")
	}
	return g == w
}

func orNone(s string) string {
	if s == "" {
		return "none"
	}
	return s
}

// A "must not ask" expectation names a catalog row the conversation must
// never send (PR-28b, D-643). Expect names the value a slot must end
// with, and this names a row that must not fire. A reader who says "you
// already had this answer" becomes a conversation with the row they saw
// twice, and the gate fails when that row fires again.
//
// The row id is the key and not the question text, because the ask role
// rephrases the text on every run. An invented question keeps the id of
// the row it replaced (agent.go, the M-4 record), so a replacement fires
// the expectation the same way the catalog wording does.

// checkNotAsked names every named row that fired, in ask order. A miss
// reads like an expectation miss, so the two join one list.
func checkNotAsked(rows []string, asks []asked) []string {
	want := map[string]bool{}
	for _, r := range rows {
		want[strings.TrimSpace(r)] = true
	}
	var misses []string
	seen := map[string]bool{}
	for _, a := range asks {
		if !want[a.Row] || seen[a.Row] {
			continue
		}
		seen[a.Row] = true
		misses = append(misses, fmt.Sprintf("must_not_ask: %s fired at turn %d", a.Row, a.Turn))
	}
	return misses
}

// checkRowIDs names every row a conversation forbids that the catalog
// does not hold. A typo would make an expectation that can never fail,
// so the gate refuses the file before it spends anything, and a free
// test reads the same check.
func checkRowIDs(convs []conversation, cat *questions.Catalog) error {
	var bad []string
	for _, c := range convs {
		for _, id := range c.MustNotAsk {
			if _, ok := cat.Row(strings.TrimSpace(id)); !ok {
				bad = append(bad, fmt.Sprintf("conversation %d names row %q", c.ID, id))
			}
		}
	}
	if len(bad) > 0 {
		return fmt.Errorf("must_not_ask names %d row(s) the catalog does not hold: %s", len(bad), strings.Join(bad, "; "))
	}
	return nil
}
