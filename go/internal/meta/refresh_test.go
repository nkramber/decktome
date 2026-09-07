package meta

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// fakeSites serves every source from the fixtures. It counts the
// requests, so a test can say what a second run fetched.
func fakeSites(t *testing.T) (*httptest.Server, *atomic.Int32) {
	t.Helper()
	var hits atomic.Int32
	serve := func(w http.ResponseWriter, name string) {
		data, err := os.ReadFile(filepath.Join("testdata", name))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write(data)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		p := r.URL.Path
		switch {
		case strings.HasPrefix(p, "/decklists/2026/09"):
			serve(w, "mtgo_month.html")
		case strings.HasPrefix(p, "/decklists/"):
			w.WriteHeader(http.StatusNotFound)
		case strings.Contains(p, "/decklist/modern-challenge-32-2026-09-0212853228"):
			serve(w, "mtgo_challenge.html")
		case strings.Contains(p, "/decklist/modern-league-"):
			serve(w, "mtgo_league.html")
		case strings.HasPrefix(p, "/decklist/standard-challenge-16"):
			// A page the site will not serve: a fetch error, and the next
			// run reads it again.
			w.WriteHeader(http.StatusBadGateway)
		case strings.HasPrefix(p, "/decklist/standard-league-"):
			// The site answers a redirect to the listing for an event it
			// no longer serves: a fetch error, never a stored page.
			w.Header().Set("Location", "/decklists")
			w.WriteHeader(http.StatusFound)
		case strings.HasPrefix(p, "/decklist/standard-"):
			// A standard page with no object: the parse failure M-6 counts.
			_, _ = w.Write([]byte("<html></html>"))
		case strings.HasPrefix(p, "/decklist/"):
			w.WriteHeader(http.StatusNotFound)
		case p == "/DeckList.json":
			serve(w, "mtgjson_decklist.json")
		case strings.HasPrefix(p, "/decks/"):
			serve(w, "mtgjson_commander.json")
		case p == "/cedh/":
			serve(w, "cedhdb.html")
		case strings.HasPrefix(p, "/commanders/year.json"), strings.HasPrefix(p, "/commanders/month.json"), strings.HasPrefix(p, "/commanders/week.json"):
			serve(w, "edhrec_top.json")
		case strings.HasPrefix(p, "/commanders/"):
			serve(w, "edhrec_commander.json")
		case strings.HasPrefix(p, "/average-decks/"):
			serve(w, "edhrec_average.json")
		case p == "/v2/tournaments":
			_, _ = w.Write([]byte(topdeckAnswer))
		case p == "/format":
			serve(w, "mtgtop8-format-mo.html")
		case p == "/event":
			serve(w, "mtgtop8-event-90258.html")
		case p == "/mtgo":
			serve(w, "mtgtop8-deck-885253.txt")
		case strings.HasPrefix(p, "/deck/custom/"):
			serve(w, "mtggoldfish-custom-modern.html")
		case strings.HasPrefix(p, "/deck/"):
			serve(w, "mtggoldfish-deck-7938202.html")
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(srv.Close)
	return srv, &hits
}

func testJob(t *testing.T, srv *httptest.Server, store ObjectStore) *Job {
	t.Helper()
	fetch := NewFetcher(srv.Client(), time.Millisecond, nil)
	td, err := NewTopdeck(fetch, "key", srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	td.MinPlayers = 2
	td.Now = func() time.Time { return time.Date(2026, 9, 2, 12, 0, 0, 0, time.UTC) }
	RetryWait = time.Millisecond
	return &Job{
		Store: store, Fetch: fetch, Topdeck: td, TopdeckDays: 7,
		Now:        func() time.Time { return time.Date(2026, 9, 2, 12, 0, 0, 0, time.UTC) },
		MTGOMonths: 2, MTGOBase: srv.URL, MTGJSONBase: srv.URL, EDHRECBase: srv.URL, CEDHDBURL: srv.URL + "/cedh/",
		MTGTop8Base: srv.URL, GoldfishBase: srv.URL, GoldfishListPages: 2,
		Commanders: []string{"Kinnan, Bonder Prodigy"},
	}
}

func TestJobRun(t *testing.T) {
	ctx := context.Background()
	srv, hits := fakeSites(t)
	store := DirObjects{Root: t.TempDir()}
	job := testJob(t, srv, store)
	rep, err := job.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(rep.Errors) != 0 {
		t.Fatalf("errors: %v", rep.Errors)
	}
	// The month page names six events: two of a covered format, and the
	// standard one does not fetch, so the run goes on without it.
	if rep.Pages[SourceMTGO] != 1 || rep.Failures[SourceMTGO] != 0 || rep.FetchErrors[SourceMTGO] != 1 {
		t.Errorf("mtgo pages %d, failures %d, fetch errors %d", rep.Pages[SourceMTGO], rep.Failures[SourceMTGO], rep.FetchErrors[SourceMTGO])
	}
	if rep.Lists[SourceMTGO] != 3 {
		t.Errorf("mtgo lists = %d, want the 3 of the challenge", rep.Lists[SourceMTGO])
	}
	if rep.FailureRate(SourceMTGO) != 0 {
		t.Errorf("failure rate = %.2f", rep.FailureRate(SourceMTGO))
	}
	if ok, _ := HasRaw(ctx, store, SourceMTGO, "standard-challenge-16-2026-09-0212853229"); ok {
		t.Errorf("the page that did not fetch is stored")
	}
	if rep.Precons != 8 || rep.PreconsVersion != "5.3.0+20260902" || rep.Lists[SourceMTGJSON] != 1 {
		t.Errorf("precons %d of %s, lists %d (eight files serve one fixture, so one key)", rep.Precons, rep.PreconsVersion, rep.Lists[SourceMTGJSON])
	}
	if rep.Lists[SourceTopdeck] != 3 || rep.Lists[SourceEDHREC] == 0 {
		t.Errorf("topdeck lists %d, edhrec lists %d", rep.Lists[SourceTopdeck], rep.Lists[SourceEDHREC])
	}
	if rep.Commanders == 0 {
		t.Errorf("no commander read")
	}
	// Every paper event of the format page serves the same fixture, so
	// the first one yields its 64 lists and the others find every deck
	// stored. The three format pages and the three later pages count,
	// and the site's own pagination hrefs repeat across formats.
	events, _ := ParseMTGTop8Format(readTestdata(t, "mtgtop8-format-mo.html"))
	paper := 0
	for _, ev := range events {
		if ev.Paper {
			paper++
		}
	}
	if paper == 0 {
		t.Fatal("the format fixture names no paper event")
	}
	if rep.Lists[SourceMTGTop8] != 64 || rep.Pages[SourceMTGTop8] != 6+paper+64 || rep.Failures[SourceMTGTop8] != 0 {
		t.Errorf("mtgtop8 lists %d, pages %d (want %d), failures %d", rep.Lists[SourceMTGTop8], rep.Pages[SourceMTGTop8], 6+paper+64, rep.Failures[SourceMTGTop8])
	}
	// Two formats, two listing pages each, and the three decks of the
	// fixture once.
	if rep.Lists[SourceGoldfish] != 3 || rep.Pages[SourceGoldfish] != 7 || rep.Failures[SourceGoldfish] != 0 {
		t.Errorf("mtggoldfish lists %d, pages %d, failures %d", rep.Lists[SourceGoldfish], rep.Pages[SourceGoldfish], rep.Failures[SourceGoldfish])
	}
	cs, err := ReadCommanders(ctx, store, "2026-09-02")
	if err != nil {
		t.Fatal(err)
	}
	bySlug := map[string]Commander{}
	for _, c := range cs {
		bySlug[c.Slug] = c
	}
	kinnan := bySlug["kinnan-bonder-prodigy"]
	if kinnan.Entries != 1 || kinnan.TopCuts != 1 || kinnan.NumDecks != 21106 {
		t.Errorf("kinnan = %+v, want the tournament count and the EDHREC read on one row", kinnan)
	}
	// One top cut in one entry shrinks to (1 + 1.25) / 6 (D-485).
	if sig := kinnan.CEDHSignal(); sig < 0.37 || sig > 0.38 {
		t.Errorf("kinnan signal = %.3f", sig)
	}
	if _, ok := bySlug["thrasios-triton-hero-tymna-the-weaver"]; !ok {
		t.Errorf("the pair row is absent: %v", bySlug)
	}
	// The database fixture holds brews alone, so no competitive row.
	for _, c := range cs {
		if c.Competitive {
			t.Errorf("%s is competitive, and the fixture holds no competitive entry", c.Slug)
		}
	}
	first := hits.Load()

	// A second run the same day fetches the month pages, the deck list,
	// and the page that did not fetch, and reads nothing else again.
	rep, err = job.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Pages[SourceMTGO] != 0 || rep.FetchErrors[SourceMTGO] != 1 || rep.Precons != 0 || rep.Skipped[SourceMTGJSON] == "" || rep.Skipped[SourceCEDHDB] == "" || rep.Skipped[SourceEDHREC] == "" {
		t.Errorf("second run: %+v", rep)
	}
	// The MTGTop8 lane reads its six format pages again, and the
	// MTGGoldfish lane its four listing pages, and both find every event
	// and every deck stored.
	if rep.Pages[SourceMTGTop8] != 6 || rep.Lists[SourceMTGTop8] != 0 || rep.Pages[SourceGoldfish] != 4 || rep.Lists[SourceGoldfish] != 0 {
		t.Errorf("second run: mtgtop8 pages %d lists %d, mtggoldfish pages %d lists %d",
			rep.Pages[SourceMTGTop8], rep.Lists[SourceMTGTop8], rep.Pages[SourceGoldfish], rep.Lists[SourceGoldfish])
	}
	if hits.Load()-first > 16 {
		t.Errorf("second run made %d requests", hits.Load()-first)
	}

	// The re-parse reads the stored pages and fetches nothing.
	job.Reparse = true
	before := hits.Load()
	rep, err = job.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Pages[SourceMTGO] != 1 || rep.Failures[SourceMTGO] != 0 {
		t.Errorf("reparse: pages %d, failures %d", rep.Pages[SourceMTGO], rep.Failures[SourceMTGO])
	}
	if rep.Pages[SourceMTGTop8] != paper || rep.Pages[SourceGoldfish] != 3 {
		t.Errorf("reparse: mtgtop8 pages %d (want %d), mtggoldfish pages %d (want 3)", rep.Pages[SourceMTGTop8], paper, rep.Pages[SourceGoldfish])
	}
	if hits.Load()-before > 4 {
		t.Errorf("reparse made %d requests", hits.Load()-before)
	}
}

// TestEDHRECSkipsInsideTheWeek: the tournament lane writes the day's
// commanders file before the EDHREC lane, so the newest commanders day
// is today on every run. The skip reads the last completed read
// instead, and it holds for EDHRECDays (F-50).
func TestEDHRECSkipsInsideTheWeek(t *testing.T) {
	ctx := context.Background()
	srv, _ := fakeSites(t)
	store := DirObjects{Root: t.TempDir()}
	job := testJob(t, srv, store)
	if _, err := job.Run(ctx); err != nil {
		t.Fatal(err)
	}
	if has, _ := job.edhrecReadOn(ctx, "2026-09-02"); !has {
		t.Fatal("the first run left no read marker")
	}

	job.Now = func() time.Time { return time.Date(2026, 9, 5, 12, 0, 0, 0, time.UTC) }
	rep, err := job.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if last, _ := LatestCommandersDay(ctx, store); last != "2026-09-05" {
		t.Fatalf("newest commanders day = %q, want today's file from the tournament lane", last)
	}
	if rep.Skipped[SourceEDHREC] != "read on 2026-09-02" || rep.Pages[SourceEDHREC] != 0 {
		t.Errorf("three days on: skipped %q, pages %d", rep.Skipped[SourceEDHREC], rep.Pages[SourceEDHREC])
	}

	job.Now = func() time.Time { return time.Date(2026, 9, 10, 12, 0, 0, 0, time.UTC) }
	rep, err = job.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Skipped[SourceEDHREC] != "" || rep.Pages[SourceEDHREC] == 0 {
		t.Errorf("eight days on: skipped %q, pages %d", rep.Skipped[SourceEDHREC], rep.Pages[SourceEDHREC])
	}
	if has, _ := job.edhrecReadOn(ctx, "2026-09-10"); !has {
		t.Errorf("the read of the eighth day left no marker")
	}
}

func TestJobRunWithoutTopdeck(t *testing.T) {
	srv, _ := fakeSites(t)
	job := testJob(t, srv, DirObjects{Root: t.TempDir()})
	job.Topdeck = nil
	rep, err := job.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rep.Skipped[SourceTopdeck], TopdeckKeyEnv) {
		t.Errorf("skipped = %v", rep.Skipped)
	}
}

// TestFetcherRetriesOnce: a transport error or a 5xx gets one retry,
// and a 404 gets none.
func TestFetcherRetriesOnce(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := hits.Add(1)
		switch r.URL.Path {
		case "/flaky":
			if n == 1 {
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			_, _ = w.Write([]byte("ok"))
		case "/gone":
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()
	RetryWait = time.Millisecond
	f := NewFetcher(srv.Client(), time.Millisecond, nil)
	data, err := f.Get(context.Background(), srv.URL+"/flaky")
	if err != nil || string(data) != "ok" || hits.Load() != 2 {
		t.Fatalf("flaky: %q, %v, %d hits", data, err, hits.Load())
	}
	hits.Store(0)
	if _, err := f.Get(context.Background(), srv.URL+"/gone"); !NotFound(err) || hits.Load() != 1 {
		t.Fatalf("gone: %v, %d hits", err, hits.Load())
	}
}

// TestFetcherWaitsOnRateLimit: a 429 waits the Retry-After header, then
// retries once.
func TestFetcherWaitsOnRateLimit(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if hits.Add(1) <= 2 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()
	MaxRateLimitWait = 5 * time.Millisecond
	f := NewFetcher(srv.Client(), time.Millisecond, nil)
	started := time.Now()
	data, err := f.Get(context.Background(), srv.URL+"/x")
	if err != nil || string(data) != "ok" || hits.Load() != 3 {
		t.Fatalf("rate limit: %q, %v, %d hits (two 429s, then the answer)", data, err, hits.Load())
	}
	if time.Since(started) < 10*time.Millisecond {
		t.Errorf("no wait before the retries")
	}
	// A site that never relents ends after MaxRateLimitRetries.
	hits.Store(-100)
	if _, err := f.Get(context.Background(), srv.URL+"/x"); !isRateLimit(err) || hits.Load() != -100+1+MaxRateLimitRetries {
		t.Errorf("endless 429: %v, %d hits", err, hits.Load()+100)
	}
	wait, ok := retryAfter(&StatusError{Status: 429})
	if !ok || wait != RateLimitWait {
		t.Errorf("bare 429 waits %s, %v", wait, ok)
	}
	if _, ok := retryAfter(&StatusError{Status: 403}); ok {
		t.Errorf("a 403 retries")
	}
}

// TestJobKeepsFailedPageApart: a page that does not parse is stored
// under its own prefix, and the next run fetches the event again.
func TestJobKeepsFailedPageApart(t *testing.T) {
	ctx := context.Background()
	srv, _ := fakeSites(t)
	store := DirObjects{Root: t.TempDir()}
	job := testJob(t, srv, store)
	job.MTGOBase = srv.URL + "/broken"
	// The broken site serves the month page, and every event page is an
	// empty document.
	mux := http.NewServeMux()
	mux.HandleFunc("/broken/decklists/", func(w http.ResponseWriter, _ *http.Request) {
		data, _ := os.ReadFile(filepath.Join("testdata", "mtgo_month.html"))
		_, _ = w.Write(data)
	})
	mux.HandleFunc("/broken/decklist/", func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte("<html></html>")) })
	broken := httptest.NewServer(mux)
	defer broken.Close()
	job.MTGOBase = broken.URL + "/broken"
	rep, err := job.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Pages[SourceMTGO] != 2 || rep.Failures[SourceMTGO] != 2 {
		t.Fatalf("pages %d, failures %d", rep.Pages[SourceMTGO], rep.Failures[SourceMTGO])
	}
	if ok, _ := HasRaw(ctx, store, SourceMTGO, "modern-challenge-32-2026-09-0212853228"); ok {
		t.Errorf("a page that did not parse is stored as the event")
	}
	if ok, _ := HasRaw(ctx, store, SourceMTGO+"-failed", "modern-challenge-32-2026-09-0212853228"); !ok {
		t.Errorf("the failed page is not kept for a reader")
	}
	rep, err = job.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Pages[SourceMTGO] != 2 {
		t.Errorf("the second run fetched %d pages, want the two again", rep.Pages[SourceMTGO])
	}
}

// TestMTGJSONSkipsAnUnchangedDeckList: the version stamp carries the
// build day, so the table of yesterday never matches today's stamp. The
// job compares the products instead. It reads no deck file while they
// hold, and it reads them all again on a new product or after thirty
// days.
func TestMTGJSONSkipsAnUnchangedDeckList(t *testing.T) {
	ctx := context.Background()
	list := fixture(t, "mtgjson_decklist.json")
	deck := fixture(t, "mtgjson_commander.json")
	var mu sync.Mutex
	version, extra := "5.3.0+20260902", ""
	var deckHits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/DeckList.json":
			mu.Lock()
			doc := bytes.Replace(list, []byte("5.3.0+20260902"), []byte(version), 1)
			doc = bytes.Replace(doc, []byte(`"data": [`), []byte(`"data": [`+extra), 1)
			mu.Unlock()
			_, _ = w.Write(doc)
		case strings.HasPrefix(r.URL.Path, "/decks/"):
			deckHits.Add(1)
			_, _ = w.Write(deck)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(srv.Close)
	store := DirObjects{Root: t.TempDir()}
	job := testJob(t, srv, store)
	job.Logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	run := func(v, add string) *Report {
		t.Helper()
		mu.Lock()
		version, extra = v, add
		mu.Unlock()
		rep := newReport()
		if err := job.runMTGJSON(ctx, rep); err != nil {
			t.Fatalf("%s: %v", v, err)
		}
		return rep
	}

	// The first read stores the table and its deck list.
	rep := run("5.3.0+20260902", "")
	if rep.Precons != 8 || rep.PreconsVersion != "5.3.0+20260902" || deckHits.Load() != 8 {
		t.Fatalf("first run: precons %d of %s, %d deck fetches", rep.Precons, rep.PreconsVersion, deckHits.Load())
	}

	// The next day's stamp names the same products: no deck fetch, and
	// the report names the stored table.
	rep = run("5.3.0+20260903", "")
	if rep.Precons != 0 || rep.PreconsVersion != "5.3.0+20260902" || rep.Skipped[SourceMTGJSON] == "" || deckHits.Load() != 8 {
		t.Errorf("same products: precons %d of %s, skipped %q, %d deck fetches", rep.Precons, rep.PreconsVersion, rep.Skipped[SourceMTGJSON], deckHits.Load())
	}
	if _, ok, _ := store.Get(ctx, PreconsName("5.3.0+20260903")); ok {
		t.Error("the skip wrote a table of the new stamp")
	}

	// A new product reads every file again under the new stamp.
	newDeck := `{"code": "NEW", "fileName": "NewDeck_NEW", "name": "New Deck", "releaseDate": "2026-09-04", "type": "Commander Deck"}, `
	rep = run("5.3.0+20260904", newDeck)
	if rep.Precons != 9 || rep.PreconsVersion != "5.3.0+20260904" || deckHits.Load() != 17 {
		t.Errorf("new product: precons %d of %s, %d deck fetches", rep.Precons, rep.PreconsVersion, deckHits.Load())
	}

	// The same products thirty days on read whole again, and the day
	// after that skips on the fresh table.
	rep = run("5.3.0+20261010", newDeck)
	if rep.Precons != 9 || rep.PreconsVersion != "5.3.0+20261010" || deckHits.Load() != 26 {
		t.Errorf("thirty days: precons %d of %s, %d deck fetches", rep.Precons, rep.PreconsVersion, deckHits.Load())
	}
	rep = run("5.3.0+20261011", newDeck)
	if rep.Precons != 0 || rep.PreconsVersion != "5.3.0+20261010" || deckHits.Load() != 26 {
		t.Errorf("day after: precons %d of %s, %d deck fetches", rep.Precons, rep.PreconsVersion, deckHits.Load())
	}
}

func TestSameProductsReadsTheKeptTypesAlone(t *testing.T) {
	a := []DeckEntry{
		{Code: "WHO", FileName: "TimeyWimey_WHO", Name: "Timey-Wimey", ReleaseDate: "2023-10-13", Type: "Commander Deck"},
		{Code: "J25", FileName: "Pack_J25", Name: "A Jumpstart pack", ReleaseDate: "2024-11-15", Type: "Jumpstart"},
	}
	b := []DeckEntry{a[0]}
	if !SameProducts(a, b) {
		t.Error("a Jumpstart pack is not a product of the table, so it must not count")
	}
	c := []DeckEntry{{Code: "WHO", FileName: "TimeyWimey_WHO", Name: "Timey-Wimey", ReleaseDate: "2023-10-14", Type: "Commander Deck"}}
	if SameProducts(a, c) {
		t.Error("a changed release date is a changed product")
	}
	if day, ok := VersionDay("5.3.0+20260903"); !ok || day.Format("2006-01-02") != "2026-09-03" {
		t.Errorf("version day = %v, %v", day, ok)
	}
	if _, ok := VersionDay("5.3.0"); ok {
		t.Error("a stamp with no day has no day")
	}
}
