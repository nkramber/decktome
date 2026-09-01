package cards

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"slices"
	"testing"
	"time"
)

// TestLiveSetTable drives the set table against the stored snapshot, so
// a session can measure the real numbers instead of guessing them.
// Guarded: set LIVE_SNAPSHOT=1 and CARDS_SNAPSHOT_DIR to run it. Every
// paid gate reads the same variable, so one convention covers them all.
func TestLiveSetTable(t *testing.T) {
	root := os.Getenv("CARDS_SNAPSHOT_DIR")
	if os.Getenv("LIVE_SNAPSHOT") != "1" || root == "" {
		t.Skip("set LIVE_SNAPSHOT=1 and CARDS_SNAPSHOT_DIR to run")
	}
	var before runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	start := time.Now()
	idx, err := LoadIndex(context.Background(), DirStore{Root: root}, slog.New(slog.NewTextHandler(os.Stderr, nil)))
	if err != nil || idx == nil {
		t.Fatalf("load: %v", err)
	}
	took := time.Since(start)
	var after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&after)
	t.Logf("load %s, cards %d, sets %d, derived %v, heap %.1f MB",
		took, idx.Len(), idx.Sets().Len(), idx.Sets().Derived(),
		float64(after.HeapAlloc-before.HeapAlloc)/(1<<20))

	pairs, most, withSets := 0, 0, 0
	for _, c := range idx.All() {
		n := len(c.GetSetCodes())
		pairs += n
		if n > 0 {
			withSets++
		}
		if n > most {
			most = n
		}
	}
	t.Logf("cards with a paper set %d, card-set pairs %d, most sets on one card %d", withSets, pairs, most)

	tbl := idx.Sets()
	for _, tc := range []struct {
		phrase string
		want   []string
	}{
		{"The Hobbit", []string{"hob", "hoc"}},
		{"the hobbit set", []string{"hob", "hoc"}},
		{"Final Fantasy", []string{"fca", "fic", "fin", "pfin", "pss5", "rfin"}},
		{"Lord of the Rings", []string{"ltc", "ltr", "pltc", "pltr"}},
		{"Bloomburrow", []string{"blb", "blc", "pblb"}},
		{"Duskmourn", []string{"dsc", "dsk", "pdsk"}},
	} {
		got := tbl.Resolve(tc.phrase)
		if got.Kind != ResolveOne || !slices.Equal(got.Codes, tc.want) {
			t.Errorf("Resolve(%q) = %v %v, want ResolveOne %v", tc.phrase, got.Kind, got.Codes, tc.want)
		}
	}
	for _, phrase := range []string{"Strixhaven", "Kamigawa", "Ravnica", "Tarkir"} {
		got := tbl.Resolve(phrase)
		if got.Kind != ResolveMany {
			t.Errorf("Resolve(%q) = %v, want ResolveMany", phrase, got.Kind)
			continue
		}
		var names []string
		for _, s := range got.Candidates {
			names = append(names, s.Code+" "+s.Name)
		}
		t.Logf("Resolve(%q) asks: %v", phrase, names)
	}
	// The three commanders the owner named on 2026-08-31 are in the
	// family, and two of them are in hob while one is in hoc.
	family := CodeSet(tbl.Family("hob"))
	for _, name := range []string{"Smaug the Impenetrable", "Thranduil, the Elvenking", "Smaug, Wicked Worm"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Errorf("the snapshot holds no card named %q", name)
			continue
		}
		if !InSets(c, family) {
			t.Errorf("%q is not in the Hobbit family: %v", name, c.GetSetCodes())
		}
		t.Logf("%-26s can lead %v  sets %v", name, c.GetCanBeCommander(), c.GetSetCodes())
	}
}
