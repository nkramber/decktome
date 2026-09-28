package meta

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"time"
)

// Job is one refresh of every source. Each source runs on its own, a
// failed source logs and the next one runs, and the report counts the
// pages and the failures per source (M-6).
type Job struct {
	Store  ObjectStore
	Fetch  *Fetcher
	Logger *slog.Logger
	// Now is the clock, for the tests.
	Now func() time.Time
	// MTGOMonths is how many months back the MTGO read goes, this month
	// included. The default is 12: a format's lists of a year ago say
	// little about its lists today (session call, 2026-09-02).
	MTGOMonths int
	// MaxPages caps the event pages one run fetches, newest month first.
	// The default is 200. On Cloud Run the MTGO site took about 14
	// seconds a request, so 200 pages and 141 redirects took 82 minutes
	// (F-183).
	MaxPages int
	// Topdeck is the API client, nil without a key. The job then reads
	// no tournament and says so.
	Topdeck *Topdeck
	// TopdeckDays is the window of the cEDH read, default 90.
	TopdeckDays int
	// Commanders names commanders to read on EDHREC beside the top
	// pages, for example the most built legends of the card index.
	Commanders []string
	// EDHRECDays is how old a commander read can be before the job
	// reads again, default 7.
	EDHRECDays int
	// Reparse re-reads every stored MTGO page in place of a fetch. It is
	// the M-6 repair: a parser fix re-reads the raw pages and fetches
	// nothing.
	Reparse bool
	// MTGTop8Pages caps the requests of the MTGTop8 lane in one run:
	// format pages, event pages, and deck exports together. The default
	// is 300, five minutes at the fetcher's gap. GoldfishPages caps the
	// requests of the MTGGoldfish lane the same way, default 200, and
	// GoldfishListPages caps the listing pages it walks per format,
	// default 100 (PR-14C).
	MTGTop8Pages, GoldfishPages, GoldfishListPages int
	// MTGOBase, MTGJSONBase, EDHRECBase, CEDHDBURL, MTGTop8Base, and
	// GoldfishBase override the sites, for the tests.
	MTGOBase, MTGJSONBase, EDHRECBase, CEDHDBURL, MTGTop8Base, GoldfishBase string
	// MTGORetryWait is the least time between two reads of an empty
	// MTGO month page, default DefaultMTGORetryWait (F-179, D-974,
	// D-977, D-984).
	MTGORetryWait time.Duration
	// MTGOBudget is the time from the start of the MTGO lane after
	// which the lane starts no fetch and no wait, default
	// DefaultMTGOBudget. The pages it does not read wait for the next
	// run, and the quality fit after the sources keeps its time (F-184,
	// D-982).
	MTGOBudget time.Duration

	// mtgo holds the month pages of this run, for the retry and the
	// event fetch after the other sources.
	mtgo *mtgoRun
}

// DefaultMTGORetryWait is the default of Job.MTGORetryWait. The site
// served an empty deck list for each older month for 4.5 minutes on
// 2026-09-27 (F-179).
const DefaultMTGORetryWait = 5 * time.Minute

// MTGORetryPasses is how many times the job reads an empty MTGO
// month page again. The site served full and empty reads of one page
// seconds apart on 2026-09-27, so one more read can get the same empty
// page (F-179, D-977).
const MTGORetryPasses = 3

// DefaultMTGOBudget is the default of Job.MTGOBudget. On 2026-09-27 the
// MTGO lane of mtg-meta-4nps4 ended 150 minutes after its start, with
// the 302 retry. So the task timeout of 150 minutes stopped the quality
// fit 14 seconds after its start (F-184).
const DefaultMTGOBudget = 2 * time.Hour

// mtgoRun is the state of the MTGO lane of one run, from the read of the
// month pages to the fetch of the event pages (D-979).
type mtgoRun struct {
	// slugs holds the covered slugs of each month, newest month first.
	slugs [][]string
	// months holds the month pages with no event link.
	months []mtgoMonth
	seen   map[string]bool
	stamp  string
	// deadline is the end of the MTGO budget of the run, zero for none.
	// held tells that a step stopped at it.
	deadline time.Time
	held     bool
}

// inBudget answers whether a step that starts after wait starts before
// the deadline of the run. The first step that does not names the
// budget in the report and the log (F-184, D-982).
func (j *Job) inBudget(run *mtgoRun, rep *Report, wait time.Duration, stage string) bool {
	if run.deadline.IsZero() || time.Now().Add(wait).Before(run.deadline) {
		return true
	}
	if !run.held {
		run.held = true
		j.Logger.Warn("mtgo time budget held", "stage", stage, "budget", j.MTGOBudget.String())
	}
	rep.Skipped[SourceMTGO] = fmt.Sprintf("the time budget of %s held, and more pages wait", j.MTGOBudget)
	return false
}

// mtgoMonth is one month page that held no event link on its first read.
// index is its place in mtgoRun.slugs. first is the time of that read,
// and read is the time of the last one.
type mtgoMonth struct {
	key, url    string
	index       int
	first, read time.Time
}

// Report counts what one run did.
type Report struct {
	// Pages counts the fetched pages per source, and Failures the pages
	// that did not parse (M-6). FetchErrors counts the pages a source
	// could not fetch after the retry, which the next run reads again.
	Pages       map[string]int
	Failures    map[string]int
	FetchErrors map[string]int
	// Lists counts the new lists per source.
	Lists map[string]int
	// Precons is the size of the table the run wrote, 0 when the stored
	// version was current.
	Precons        int
	PreconsVersion string
	// Commanders counts the commander reads of the run.
	Commanders int
	// Skipped names the sources the run did not read, with the reason.
	Skipped map[string]string
	Errors  []string
}

func newReport() *Report {
	return &Report{Pages: map[string]int{}, Failures: map[string]int{}, FetchErrors: map[string]int{}, Lists: map[string]int{}, Skipped: map[string]string{}}
}

// FailureRate is the M-6 number of one source: failures over pages.
func (r *Report) FailureRate(source string) float64 {
	if r.Pages[source] == 0 {
		return 0
	}
	return float64(r.Failures[source]) / float64(r.Pages[source])
}

func (j *Job) defaults() {
	if j.Logger == nil {
		j.Logger = slog.Default()
	}
	if j.Now == nil {
		j.Now = time.Now
	}
	if j.Fetch == nil {
		j.Fetch = NewFetcher(nil, 0, j.Logger)
	}
	if j.MTGOMonths <= 0 {
		j.MTGOMonths = 12
	}
	if j.MaxPages <= 0 {
		j.MaxPages = 200
	}
	if j.TopdeckDays <= 0 {
		j.TopdeckDays = 90
	}
	if j.EDHRECDays <= 0 {
		j.EDHRECDays = 7
	}
	if j.MTGOBase == "" {
		j.MTGOBase = MTGOBase
	}
	if j.MTGJSONBase == "" {
		j.MTGJSONBase = MTGJSONBase
	}
	if j.EDHRECBase == "" {
		j.EDHRECBase = EDHRECBase
	}
	if j.CEDHDBURL == "" {
		j.CEDHDBURL = CEDHDBURL
	}
	if j.MTGTop8Base == "" {
		j.MTGTop8Base = MTGTop8Base
	}
	if j.GoldfishBase == "" {
		j.GoldfishBase = GoldfishBase
	}
	if j.MTGTop8Pages <= 0 {
		j.MTGTop8Pages = 300
	}
	if j.GoldfishPages <= 0 {
		j.GoldfishPages = 200
	}
	if j.GoldfishListPages <= 0 {
		j.GoldfishListPages = 100
	}
	if j.MTGORetryWait <= 0 {
		j.MTGORetryWait = DefaultMTGORetryWait
	}
	if j.MTGOBudget <= 0 {
		j.MTGOBudget = DefaultMTGOBudget
	}
}

// Run reads every source. The error is the context's alone: a source
// failure is a report row.
func (j *Job) Run(ctx context.Context) (*Report, error) {
	j.defaults()
	rep := newReport()
	steps := []struct {
		name string
		run  func(context.Context, *Report) error
	}{
		{SourceMTGO, j.runMTGO},
		{SourceMTGJSON, j.runMTGJSON},
		{SourceCEDHDB, j.runCEDHDB},
		{SourceTopdeck, j.runTopdeck},
		{SourceEDHREC, j.runEDHREC},
		{SourceMTGTop8, j.runMTGTop8},
		{SourceGoldfish, j.runGoldfish},
	}
	for _, s := range steps {
		if err := s.run(ctx, rep); err != nil {
			if ctx.Err() != nil {
				return rep, ctx.Err()
			}
			rep.Errors = append(rep.Errors, s.name+": "+err.Error())
			j.Logger.Error("meta source failed", "source", s.name, "err", err)
		}
	}
	// The other sources are the delay of the MTGO month retry, and the
	// event pages come after it (D-974, D-979).
	if err := j.finishMTGO(ctx, rep); err != nil {
		if ctx.Err() != nil {
			return rep, ctx.Err()
		}
		rep.Errors = append(rep.Errors, SourceMTGO+" events: "+err.Error())
		j.Logger.Error("meta source failed", "source", SourceMTGO+" events", "err", err)
	}
	return rep, nil
}

// runMTGO reads the month pages back MTGOMonths and keeps the covered
// slugs of each month. finishMTGO fetches the event pages after the
// other sources. A month page with no event link waits for the retry
// there (D-974, D-979, D-984).
func (j *Job) runMTGO(ctx context.Context, rep *Report) error {
	j.mtgo = nil
	if j.Reparse {
		return j.reparseMTGO(ctx, rep)
	}
	now := j.Now().UTC()
	run := &mtgoRun{seen: map[string]bool{}, stamp: now.Format("20060102T150405Z")}
	if j.MTGOBudget > 0 {
		run.deadline = time.Now().Add(j.MTGOBudget)
	}
	j.mtgo = run
	for i := 0; i < j.MTGOMonths; i++ {
		month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, -i, 0)
		url := fmt.Sprintf("%s/decklists/%04d/%02d", j.MTGOBase, month.Year(), int(month.Month()))
		key := month.Format("2006-01")
		run.slugs = append(run.slugs, nil)
		page, err := j.Fetch.Get(ctx, url)
		if err != nil {
			if NotFound(err) {
				j.Logger.Info("mtgo month page", "month", key, "found", false)
				continue
			}
			return err
		}
		links := ParseMTGOMonth(page)
		covered := run.add(i, links)
		// A month page with no event link can not tell a quiet site from
		// a page the site served in place of the listing, so the page
		// stays for a reader under its own prefix (D-962).
		j.Logger.Info("mtgo month page", "month", key, "found", true, "bytes", len(page), "links", len(links), "covered", covered)
		if len(links) == 0 {
			if err := PutRaw(ctx, j.Store, SourceMTGO+"-month-empty", key+"-"+run.stamp, page); err != nil {
				return err
			}
			// The current month reads again too, because the site can
			// serve it empty after a full read (F-185, D-984).
			at := time.Now()
			run.months = append(run.months, mtgoMonth{key: key, url: url, index: i, first: at, read: at})
		}
	}
	return nil
}

// add keeps each covered slug of links that no month of the run listed
// before, under the month at index i, and answers how many it kept.
func (r *mtgoRun) add(i int, links []string) int {
	covered := 0
	for _, slug := range links {
		if MTGOSlugFormat(slug) != "" && !r.seen[slug] {
			r.seen[slug] = true
			r.slugs[i] = append(r.slugs[i], slug)
			covered++
		}
	}
	return covered
}

// finishMTGO ends the MTGO lane after the other sources. It reads each
// empty month page again, then fetches the event pages of every
// month, newest month first, up to MaxPages. So a month that the retry
// reads full keeps its place before the older months (F-179, D-979). It
// then reads each event page that answered 302 again (F-183, D-980).
// Each step starts before the MTGO budget ends, or not at all (F-184,
// D-982).
func (j *Job) finishMTGO(ctx context.Context, rep *Report) error {
	run := j.mtgo
	j.mtgo = nil
	if run == nil {
		return nil
	}
	if err := j.retryMTGOMonths(ctx, rep, run); err != nil {
		return err
	}
	var slugs []string
	for _, s := range run.slugs {
		slugs = append(slugs, s...)
	}
	redirected, err := j.fetchMTGOEvents(ctx, rep, run, slugs, 0)
	if err != nil {
		return err
	}
	return j.retryMTGOEvents(ctx, rep, run, redirected)
}

// retryMTGOMonths reads each empty month page of runMTGO again, up
// to MTGORetryPasses times. Each pass comes at least MTGORetryWait after
// the reads of the pass before it, and it reads only the months that are
// still empty. The store keeps each page that is empty again under its
// own name (F-179, D-974, D-977).
func (j *Job) retryMTGOMonths(ctx context.Context, rep *Report, run *mtgoRun) error {
	months := run.months
	pass := 1
	for ; pass <= MTGORetryPasses && len(months) > 0; pass++ {
		d := j.MTGORetryWait - time.Since(months[len(months)-1].read)
		if !j.inBudget(run, rep, max(d, 0), "months") {
			break
		}
		if d > 0 {
			j.Logger.Info("mtgo month pages wait", "months", len(months), "pass", pass, "wait", d.Round(time.Second).String())
			if err := sleep(ctx, d); err != nil {
				return err
			}
		}
		var empty []mtgoMonth
		for k, m := range months {
			if !j.inBudget(run, rep, 0, "months") {
				empty = append(empty, months[k:]...)
				break
			}
			page, err := j.Fetch.Get(ctx, m.url)
			m.read = time.Now()
			if err != nil {
				if ctx.Err() != nil {
					return err
				}
				j.Logger.Warn("mtgo month page again did not fetch", "month", m.key, "pass", pass, "err", err)
				empty = append(empty, m)
				continue
			}
			links := ParseMTGOMonth(page)
			covered := run.add(m.index, links)
			j.Logger.Info("mtgo month page again", "month", m.key, "pass", pass, "after", m.read.Sub(m.first).Round(time.Second).String(),
				"bytes", len(page), "links", len(links), "covered", covered)
			if len(links) == 0 {
				if err := PutRaw(ctx, j.Store, SourceMTGO+"-month-empty", fmt.Sprintf("%s-%s-again%d", m.key, run.stamp, pass), page); err != nil {
					return err
				}
				empty = append(empty, m)
			}
		}
		months = empty
	}
	if len(months) > 0 {
		keys := make([]string, len(months))
		for i, m := range months {
			keys[i] = m.key
		}
		j.Logger.Warn("mtgo month pages stay empty", "months", strings.Join(keys, " "), "passes", pass-1)
	}
	return nil
}

// retryMTGOEvents reads each event page that answered 302 again, up to
// MTGORetryPasses times. Each pass comes at least MTGORetryWait after the
// end of the fetch before it. The site answered 302 for event pages that
// it served in full to the Mac later, so a 302 is most often a fault of
// the site at that time (F-183, D-980). A page that stays redirected is
// a fetch error, and the next run reads it again.
func (j *Job) retryMTGOEvents(ctx context.Context, rep *Report, run *mtgoRun, slugs []string) error {
	last := time.Now()
	pass := 1
	for ; pass <= MTGORetryPasses && len(slugs) > 0; pass++ {
		d := j.MTGORetryWait - time.Since(last)
		if !j.inBudget(run, rep, max(d, 0), "events again") {
			break
		}
		if d > 0 {
			j.Logger.Info("mtgo event pages wait", "slugs", len(slugs), "pass", pass, "wait", d.Round(time.Second).String())
			if err := sleep(ctx, d); err != nil {
				return err
			}
		}
		var err error
		slugs, err = j.fetchMTGOEvents(ctx, rep, run, slugs, pass)
		last = time.Now()
		if err != nil {
			return err
		}
	}
	if len(slugs) > 0 {
		rep.FetchErrors[SourceMTGO] += len(slugs)
		j.Logger.Warn("mtgo event pages stay redirected", "slugs", len(slugs), "passes", pass-1)
	}
	return nil
}

// sleep waits d, or until ctx ends.
func sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// fetchMTGOEvents fetches every event page of slugs the store lacks, in
// order, and answers the slugs whose page answered 302. Pass 0 is the
// first fetch, and pass 1 to MTGORetryPasses read a 302 again. A slug
// that answered 302 holds its place under MaxPages until its last read,
// so a newer event keeps its place before an older one (D-979, D-980).
// The fetch stops at the MTGO budget. On a retry pass, a slug that it
// did not read again stays redirected (F-184, D-982).
func (j *Job) fetchMTGOEvents(ctx context.Context, rep *Report, run *mtgoRun, slugs []string, pass int) ([]string, error) {
	var redirected []string
	stored, fetched, fetchErrors := 0, 0, 0
	defer func() {
		msg := "mtgo event slugs"
		if pass > 0 {
			msg = "mtgo event pages again"
		}
		j.Logger.Info(msg, "pass", pass, "listed", len(slugs), "stored", stored, "fetched", fetched,
			"redirected", len(redirected), "fetch_errors", fetchErrors)
	}()
	for i, slug := range slugs {
		if rep.Pages[SourceMTGO]+len(redirected) >= j.MaxPages {
			rep.Skipped[SourceMTGO] = fmt.Sprintf("the page cap of %d held, and more pages wait", j.MaxPages)
			break
		}
		if ok, err := HasRaw(ctx, j.Store, SourceMTGO, slug); err != nil {
			return nil, err
		} else if ok {
			stored++
			continue
		}
		stage := "events"
		if pass > 0 {
			stage = "events again"
		}
		if !j.inBudget(run, rep, 0, stage) {
			if pass > 0 {
				redirected = append(redirected, slugs[i:]...)
			}
			break
		}
		page, err := j.Fetch.Get(ctx, j.MTGOBase+"/decklist/"+slug)
		if err != nil {
			if ctx.Err() != nil {
				return nil, err
			}
			if NotFound(err) {
				j.Logger.Warn("mtgo event page is gone", "slug", slug)
				continue
			}
			if Redirected(err) {
				j.Logger.Warn("mtgo event page answered 302", "slug", slug, "pass", pass)
				redirected = append(redirected, slug)
				continue
			}
			// One page the site would not serve is not the whole month.
			// The store lacks it still, so the next run reads it again.
			rep.FetchErrors[SourceMTGO]++
			fetchErrors++
			j.Logger.Warn("mtgo event page did not fetch", "slug", slug, "err", err)
			continue
		}
		fetched++
		rep.Pages[SourceMTGO]++
		lists, err := ParseMTGOEvent(slug, page)
		if err != nil {
			// The page stays for a reader to inspect, under its own
			// prefix, so the next run fetches the event again: the site
			// answers a listing page in place of an event page at times.
			rep.Failures[SourceMTGO]++
			j.Logger.Warn("mtgo page did not parse", "slug", slug, "err", err)
			if err := PutRaw(ctx, j.Store, SourceMTGO+"-failed", slug, page); err != nil {
				return nil, err
			}
			continue
		}
		if err := PutRaw(ctx, j.Store, SourceMTGO, slug, page); err != nil {
			return nil, err
		}
		n, err := MergeLists(ctx, j.Store, lists)
		if err != nil {
			return nil, err
		}
		rep.Lists[SourceMTGO] += n
	}
	return redirected, nil
}

// reparseMTGO re-reads every stored page.
func (j *Job) reparseMTGO(ctx context.Context, rep *Report) error {
	names, err := j.Store.List(ctx, RawPrefix+SourceMTGO+"/")
	if err != nil {
		return err
	}
	for _, name := range names {
		slug := strings.TrimSuffix(strings.TrimPrefix(name, RawPrefix+SourceMTGO+"/"), ".gz")
		page, ok, err := GetRaw(ctx, j.Store, SourceMTGO, slug)
		if err != nil || !ok {
			return err
		}
		rep.Pages[SourceMTGO]++
		lists, err := ParseMTGOEvent(slug, page)
		if err != nil {
			rep.Failures[SourceMTGO]++
			continue
		}
		n, err := MergeLists(ctx, j.Store, lists)
		if err != nil {
			return err
		}
		rep.Lists[SourceMTGO] += n
	}
	return nil
}

// runMTGJSON reads the deck list, and the deck files of a version the
// store lacks. The table is written whole, so the run reads every file
// of the version before it writes.
func (j *Job) runMTGJSON(ctx context.Context, rep *Report) error {
	data, err := j.Fetch.Get(ctx, j.MTGJSONBase+"/DeckList.json")
	if err != nil {
		return err
	}
	rep.Pages[SourceMTGJSON]++
	version, entries, err := ParseDeckList(data)
	if err != nil {
		rep.Failures[SourceMTGJSON]++
		return err
	}
	if _, ok, err := j.Store.Get(ctx, PreconsName(version)); err != nil {
		return err
	} else if ok {
		rep.PreconsVersion = version
		rep.Skipped[SourceMTGJSON] = "the table of version " + version + " is stored"
		return nil
	}
	// The version stamp carries the build day, so a new stamp comes
	// every day while the products stay the same. The deck list stored
	// with the newest table says whether the products changed, and the
	// run reads no deck file when they did not.
	if latest, current, err := j.currentPrecons(ctx, version, entries); err != nil {
		return err
	} else if current {
		rep.PreconsVersion = latest
		rep.Skipped[SourceMTGJSON] = "the deck list of version " + version + " names the products of the stored table " + latest
		return nil
	}
	if err := PutRaw(ctx, j.Store, SourceMTGJSON, version+"/DeckList", data); err != nil {
		return err
	}
	var table []Precon
	var lists []List
	for _, e := range entries {
		if _, ok := KeepPrecon(e.Type); !ok {
			continue
		}
		key := version + "/" + e.FileName
		page, ok, err := GetRaw(ctx, j.Store, SourceMTGJSON, key)
		if err != nil {
			return err
		}
		if !ok {
			page, err = j.Fetch.Get(ctx, j.MTGJSONBase+"/decks/"+e.FileName+".json")
			if err != nil {
				if NotFound(err) {
					j.Logger.Warn("mtgjson deck file is absent", "file", e.FileName)
					continue
				}
				return err
			}
			rep.Pages[SourceMTGJSON]++
			if err := PutRaw(ctx, j.Store, SourceMTGJSON, key, page); err != nil {
				return err
			}
		}
		p, err := ParsePrecon(page)
		if err != nil {
			rep.Failures[SourceMTGJSON]++
			j.Logger.Warn("mtgjson deck did not parse", "file", e.FileName, "err", err)
			continue
		}
		table = append(table, *p)
		if l := p.List(); l != nil {
			lists = append(lists, *l)
		}
	}
	if len(table) == 0 {
		return errors.New("the deck list named no product the table keeps")
	}
	SortPrecons(table)
	if err := WritePrecons(ctx, j.Store, version, table); err != nil {
		return err
	}
	n, err := MergeLists(ctx, j.Store, lists)
	if err != nil {
		return err
	}
	rep.Lists[SourceMTGJSON] += n
	rep.Precons = len(table)
	rep.PreconsVersion = version
	return nil
}

// currentPrecons reports whether the newest stored table holds the
// products the deck list names. A table older than PreconsMaxAge is
// never current, so a corrected deck file reaches the store within the
// month.
func (j *Job) currentPrecons(ctx context.Context, version string, entries []DeckEntry) (latest string, current bool, err error) {
	latest, err = LatestPreconsVersion(ctx, j.Store)
	if err != nil || latest == "" {
		return latest, false, err
	}
	stored, ok, err := GetRaw(ctx, j.Store, SourceMTGJSON, latest+"/DeckList")
	if err != nil || !ok {
		return latest, false, err
	}
	_, storedEntries, err := ParseDeckList(stored)
	if err != nil || !SameProducts(storedEntries, entries) {
		return latest, false, nil
	}
	newDay, okNew := VersionDay(version)
	oldDay, okOld := VersionDay(latest)
	if okNew && okOld && newDay.Sub(oldDay) >= PreconsMaxAge {
		return latest, false, nil
	}
	return latest, true, nil
}

// runCEDHDB reads the database page once a day and marks the
// competitive commanders in the day's commander file.
func (j *Job) runCEDHDB(ctx context.Context, rep *Report) error {
	day := j.Now().UTC().Format("2006-01-02")
	if ok, err := HasRaw(ctx, j.Store, SourceCEDHDB, day); err != nil {
		return err
	} else if ok {
		rep.Skipped[SourceCEDHDB] = "read today"
		return nil
	}
	page, err := j.Fetch.Get(ctx, j.CEDHDBURL)
	if err != nil {
		return err
	}
	rep.Pages[SourceCEDHDB]++
	if err := PutRaw(ctx, j.Store, SourceCEDHDB, day, page); err != nil {
		return err
	}
	entries, err := ParseCEDHDB(page)
	if err != nil {
		rep.Failures[SourceCEDHDB]++
		return err
	}
	var reads []Commander
	seen := map[string]bool{}
	for _, e := range entries {
		if !e.Competitive() {
			continue
		}
		slugs := [][]string{e.Commanders}
		if len(e.Commanders) > 1 {
			for _, name := range e.Commanders {
				slugs = append(slugs, []string{name})
			}
		}
		for _, names := range slugs {
			slug := EDHRECSlug(names...)
			if seen[slug] {
				continue
			}
			seen[slug] = true
			reads = append(reads, Commander{Name: strings.Join(names, " + "), Slug: slug, Competitive: true, ReadAt: day})
		}
	}
	rep.Commanders += len(reads)
	return j.mergeCommanders(ctx, day, reads)
}

// runTopdeck reads the cEDH tournaments of the window and counts the
// top cuts per commander.
func (j *Job) runTopdeck(ctx context.Context, rep *Report) error {
	if j.Topdeck == nil {
		rep.Skipped[SourceTopdeck] = "no " + TopdeckKeyEnv
		return nil
	}
	ts, err := j.Topdeck.Tournaments(ctx, TopdeckFormatEDH, j.TopdeckDays)
	if err != nil {
		return err
	}
	rep.Pages[SourceTopdeck]++
	day := j.Now().UTC().Format("2006-01-02")
	counts := map[string]*Commander{}
	var lists []List
	for _, t := range ts {
		for _, l := range t.Lists() {
			lists = append(lists, l)
			if len(l.Commanders) == 0 {
				continue
			}
			slug := EDHRECSlug(l.Commanders...)
			c, ok := counts[slug]
			if !ok {
				c = &Commander{Name: strings.Join(l.Commanders, " + "), Slug: slug, ReadAt: day}
				counts[slug] = c
			}
			c.Entries++
			if l.Tier == TierGreat {
				c.TopCuts++
			}
		}
	}
	n, err := MergeLists(ctx, j.Store, lists)
	if err != nil {
		return err
	}
	rep.Lists[SourceTopdeck] += n
	reads := make([]Commander, 0, len(counts))
	for _, c := range counts {
		reads = append(reads, *c)
	}
	sort.Slice(reads, func(a, b int) bool { return reads[a].Slug < reads[b].Slug })
	rep.Commanders += len(reads)
	return j.mergeCommanders(ctx, day, reads)
}

// runEDHREC reads the top commanders and the named ones, each with its
// average deck, when the last complete read is older than EDHRECDays.
func (j *Job) runEDHREC(ctx context.Context, rep *Report) error {
	now := j.Now().UTC()
	day := now.Format("2006-01-02")
	last, err := j.latestEDHRECRead(ctx)
	if err != nil {
		return err
	}
	if last == day {
		rep.Skipped[SourceEDHREC] = "read today"
		return nil
	}
	if last != "" {
		t, err := time.Parse("2006-01-02", last)
		if err == nil && now.Sub(t) < time.Duration(j.EDHRECDays)*24*time.Hour {
			rep.Skipped[SourceEDHREC] = "read on " + last
			return nil
		}
	}
	slugs := map[string]string{}
	var order []string
	add := func(slug, name string) {
		if slug == "" || slugs[slug] != "" {
			return
		}
		slugs[slug] = name
		order = append(order, slug)
	}
	for _, window := range []string{"year", "month", "week"} {
		page, err := j.Fetch.Get(ctx, j.EDHRECBase+"/commanders/"+window+".json")
		if err != nil {
			return err
		}
		rep.Pages[SourceEDHREC]++
		top, err := ParseEDHRECTop(page)
		if err != nil {
			rep.Failures[SourceEDHREC]++
			return err
		}
		for _, c := range top {
			add(c.Slug, c.Name)
		}
	}
	for _, name := range j.Commanders {
		add(EDHRECSlug(name), name)
	}
	// The read is bounded by its list (D-480), so no page cap applies.
	// A slug today's file already holds with a deck count was read by a
	// run that stopped short, and it does not read again.
	done := map[string]bool{}
	if have, err := ReadCommanders(ctx, j.Store, day); err == nil {
		for _, c := range have {
			if c.NumDecks > 0 {
				done[c.Slug] = true
			}
		}
	}
	var reads []Commander
	var lists []List
	for _, slug := range order {
		if done[slug] {
			continue
		}
		page, err := j.Fetch.Get(ctx, j.EDHRECBase+"/commanders/"+slug+".json")
		if err != nil {
			if ctx.Err() != nil {
				return err
			}
			// A page the site will not serve is one commander, not the
			// read: the site answered 403 to one slug on 2026-09-02.
			rep.FetchErrors[SourceEDHREC]++
			j.Logger.Warn("edhrec commander page did not fetch", "slug", slug, "err", err)
			continue
		}
		rep.Pages[SourceEDHREC]++
		c, err := ParseEDHRECCommander(slug, page)
		if err != nil {
			rep.Failures[SourceEDHREC]++
			continue
		}
		c.ReadAt = day
		if err := PutRaw(ctx, j.Store, SourceEDHREC, day+"/"+slug, page); err != nil {
			return err
		}
		reads = append(reads, *c)
		avg, err := j.Fetch.Get(ctx, j.EDHRECBase+"/average-decks/"+slug+".json")
		if err != nil {
			if ctx.Err() != nil {
				return err
			}
			if !NotFound(err) {
				rep.FetchErrors[SourceEDHREC]++
			}
			continue
		}
		rep.Pages[SourceEDHREC]++
		l, err := ParseEDHRECAverageDeck(slug, avg)
		if err != nil {
			rep.Failures[SourceEDHREC]++
			continue
		}
		l.Date = day
		lists = append(lists, *l)
	}
	n, err := MergeLists(ctx, j.Store, lists)
	if err != nil {
		return err
	}
	rep.Lists[SourceEDHREC] += n
	rep.Commanders += len(reads)
	if err := j.mergeCommanders(ctx, day, reads); err != nil {
		return err
	}
	return j.Store.Put(ctx, RawName(SourceEDHREC, day+"/read"), []byte(day))
}

// latestEDHRECRead names the newest day whose EDHREC read completed.
// The newest commanders day is not that mark: the tournament lane
// writes the day's file before this lane runs (F-50).
func (j *Job) latestEDHRECRead(ctx context.Context) (string, error) {
	names, err := j.Store.List(ctx, CommandersPrefix)
	if err != nil {
		return "", err
	}
	for i := len(names) - 1; i >= 0; i-- {
		day := strings.TrimSuffix(strings.TrimPrefix(names[i], CommandersPrefix), ".jsonl.gz")
		has, err := j.edhrecReadOn(ctx, day)
		if err != nil {
			return "", err
		}
		if has {
			return day, nil
		}
	}
	return "", nil
}

// edhrecReadOn says whether the EDHREC read of a day completed.
func (j *Job) edhrecReadOn(ctx context.Context, day string) (bool, error) {
	_, ok, err := j.Store.Get(ctx, RawName(SourceEDHREC, day+"/read"))
	return ok, err
}

// mergeCommanders folds reads into the day's file, field by field, so
// the database, the tournaments, and EDHREC each add their signal to
// one row per slug.
func (j *Job) mergeCommanders(ctx context.Context, day string, reads []Commander) error {
	have, err := ReadCommanders(ctx, j.Store, day)
	if err != nil {
		return err
	}
	index := map[string]int{}
	for i, c := range have {
		index[c.Slug] = i
	}
	for _, r := range reads {
		i, ok := index[r.Slug]
		if !ok {
			index[r.Slug] = len(have)
			have = append(have, r)
			continue
		}
		c := &have[i]
		if r.NumDecks > 0 {
			c.NumDecks, c.Rank, c.BracketCounts, c.Name = r.NumDecks, r.Rank, r.BracketCounts, r.Name
		}
		if r.Competitive {
			c.Competitive = true
		}
		if r.Entries > 0 {
			c.Entries, c.TopCuts = r.Entries, r.TopCuts
		}
		if c.ReadAt == "" {
			c.ReadAt = r.ReadAt
		}
	}
	return WriteCommanders(ctx, j.Store, day, have)
}

// runMTGTop8 reads the paper events of the covered formats (D-504): the
// format pages, then every event page the store lacks, then every deck
// export the store lacks, up to MTGTop8Pages requests a run. An event
// page names mtgo.com as its source when the MTGO lane holds the same
// lists, and the reader stores it and reads no deck of it.
func (j *Job) runMTGTop8(ctx context.Context, rep *Report) error {
	if j.Reparse {
		return j.reparseMTGTop8(ctx, rep)
	}
	fetched := 0
	held := func() bool {
		if fetched < j.MTGTop8Pages {
			return false
		}
		rep.Skipped[SourceMTGTop8] = fmt.Sprintf("the page cap of %d held, and more pages wait", j.MTGTop8Pages)
		return true
	}
	seenPage := map[string]bool{}
	for _, code := range MTGTop8FormatCodes {
		pages := []string{MTGTop8FormatURL(j.MTGTop8Base, code)}
		seenPage[pages[0]] = true
		for i := 0; i < len(pages); i++ {
			if held() {
				return nil
			}
			page, err := j.Fetch.Get(ctx, pages[i])
			if err != nil {
				if ctx.Err() != nil {
					return err
				}
				rep.FetchErrors[SourceMTGTop8]++
				j.Logger.Warn("mtgtop8 format page did not fetch", "url", pages[i], "err", err)
				continue
			}
			fetched++
			rep.Pages[SourceMTGTop8]++
			events, next := ParseMTGTop8Format(page)
			if len(events) == 0 {
				rep.Failures[SourceMTGTop8]++
				j.Logger.Warn("mtgtop8 format page holds no event", "url", pages[i])
				continue
			}
			for _, href := range next {
				u := MTGTop8PageURL(j.MTGTop8Base, href)
				if !seenPage[u] {
					seenPage[u] = true
					pages = append(pages, u)
				}
			}
			for _, ev := range events {
				if !ev.Paper {
					continue
				}
				done, err := j.readMTGTop8Event(ctx, rep, code, ev, &fetched)
				if err != nil {
					return err
				}
				if !done {
					held()
					return nil
				}
			}
		}
	}
	return nil
}

// readMTGTop8Event reads one event: its page from the store or the site,
// then the deck exports the store lacks. done is false when the cap held
// before the event was whole, and the next run reads the rest.
func (j *Job) readMTGTop8Event(ctx context.Context, rep *Report, code string, ev MTGTop8Event, fetched *int) (bool, error) {
	key := "event_" + ev.ID
	page, ok, err := GetRaw(ctx, j.Store, SourceMTGTop8, key)
	if err != nil {
		return false, err
	}
	if !ok {
		if *fetched >= j.MTGTop8Pages {
			return false, nil
		}
		page, err = j.Fetch.Get(ctx, MTGTop8EventURL(j.MTGTop8Base, ev.ID, code))
		if err != nil {
			if ctx.Err() != nil {
				return false, err
			}
			rep.FetchErrors[SourceMTGTop8]++
			j.Logger.Warn("mtgtop8 event page did not fetch", "event", ev.ID, "err", err)
			return true, nil
		}
		*fetched++
		rep.Pages[SourceMTGTop8]++
		if _, _, perr := ParseMTGTop8Event(page); perr != nil {
			// The page stays apart for a reader to inspect, and the next
			// run fetches the event again (M-6).
			rep.Failures[SourceMTGTop8]++
			j.Logger.Warn("mtgtop8 event page did not parse", "event", ev.ID, "err", perr)
			return true, PutRaw(ctx, j.Store, SourceMTGTop8+"-failed", key, page)
		}
		if err := PutRaw(ctx, j.Store, SourceMTGTop8, key, page); err != nil {
			return false, err
		}
	}
	info, rows, err := ParseMTGTop8Event(page)
	if err != nil {
		rep.Failures[SourceMTGTop8]++
		return true, nil
	}
	if info.Online {
		return true, nil
	}
	var lists []List
	whole := true
	for _, row := range rows {
		deckKey := "deck_" + row.DeckID
		has, err := HasRaw(ctx, j.Store, SourceMTGTop8, deckKey)
		if err != nil {
			return false, err
		}
		if has {
			continue
		}
		if *fetched >= j.MTGTop8Pages {
			whole = false
			break
		}
		text, err := j.Fetch.Get(ctx, MTGTop8DeckURL(j.MTGTop8Base, row.DeckID))
		if err != nil {
			if ctx.Err() != nil {
				return false, err
			}
			rep.FetchErrors[SourceMTGTop8]++
			j.Logger.Warn("mtgtop8 deck export did not fetch", "deck", row.DeckID, "err", err)
			continue
		}
		*fetched++
		rep.Pages[SourceMTGTop8]++
		l, ok := MTGTop8List(code, ev, info, row, text)
		if !ok {
			rep.Failures[SourceMTGTop8]++
			j.Logger.Warn("mtgtop8 deck export did not parse", "deck", row.DeckID)
			if err := PutRaw(ctx, j.Store, SourceMTGTop8+"-failed", deckKey, text); err != nil {
				return false, err
			}
			continue
		}
		if err := PutRaw(ctx, j.Store, SourceMTGTop8, deckKey, text); err != nil {
			return false, err
		}
		lists = append(lists, l)
	}
	n, err := MergeLists(ctx, j.Store, lists)
	if err != nil {
		return false, err
	}
	rep.Lists[SourceMTGTop8] += n
	return whole, nil
}

// reparseMTGTop8 re-reads every stored event page and its deck exports,
// and fetches nothing (M-6). The listing row is gone, so the tier reads
// the field of the page alone.
func (j *Job) reparseMTGTop8(ctx context.Context, rep *Report) error {
	prefix := RawPrefix + SourceMTGTop8 + "/"
	names, err := j.Store.List(ctx, prefix)
	if err != nil {
		return err
	}
	for _, name := range names {
		key := strings.TrimSuffix(strings.TrimPrefix(name, prefix), ".gz")
		if !strings.HasPrefix(key, "event_") {
			continue
		}
		page, ok, err := GetRaw(ctx, j.Store, SourceMTGTop8, key)
		if err != nil || !ok {
			return err
		}
		rep.Pages[SourceMTGTop8]++
		info, rows, err := ParseMTGTop8Event(page)
		if err != nil {
			rep.Failures[SourceMTGTop8]++
			continue
		}
		if info.Online {
			continue
		}
		ev := MTGTop8Event{ID: strings.TrimPrefix(key, "event_")}
		var lists []List
		for _, row := range rows {
			text, ok, err := GetRaw(ctx, j.Store, SourceMTGTop8, "deck_"+row.DeckID)
			if err != nil {
				return err
			}
			if !ok {
				continue
			}
			if l, ok := MTGTop8List(info.Format, ev, info, row, text); ok {
				lists = append(lists, l)
			}
		}
		n, err := MergeLists(ctx, j.Store, lists)
		if err != nil {
			return err
		}
		rep.Lists[SourceMTGTop8] += n
	}
	return nil
}

// runGoldfish reads the user decks of the 60-card formats as the typical
// rung (D-490, D-503): the listing pages of a format, newest first, and
// every deck page the store lacks, up to GoldfishPages requests a run.
// The walk goes past a listing page whose decks are all stored, so a
// later run reaches older decks, up to GoldfishListPages a format.
func (j *Job) runGoldfish(ctx context.Context, rep *Report) error {
	if j.Reparse {
		return j.reparseGoldfish(ctx, rep)
	}
	// Each format gets its share of the cap, so a run reads Standard
	// decks too and not the newest Modern decks alone.
	share := j.GoldfishPages / len(GoldfishFormatWords)
	if share < 1 {
		share = 1
	}
	for _, word := range GoldfishFormatWords {
		fetched := 0
		held := func() bool {
			if fetched < share {
				return false
			}
			rep.Skipped[SourceGoldfish] = fmt.Sprintf("the page cap of %d held, and more decks wait", j.GoldfishPages)
			return true
		}
		for page := 1; page <= j.GoldfishListPages; page++ {
			if held() {
				break
			}
			url := GoldfishListingURL(j.GoldfishBase, word, page)
			listing, err := j.Fetch.Get(ctx, url)
			if err != nil {
				if ctx.Err() != nil {
					return err
				}
				rep.FetchErrors[SourceGoldfish]++
				j.Logger.Warn("mtggoldfish listing did not fetch", "url", url, "err", err)
				break
			}
			fetched++
			rep.Pages[SourceGoldfish]++
			ids, more := ParseGoldfishListing(listing)
			if len(ids) == 0 {
				rep.Failures[SourceGoldfish]++
				j.Logger.Warn("mtggoldfish listing holds no deck", "url", url)
				break
			}
			var lists []List
			for _, id := range ids {
				key := "deck_" + id
				has, err := HasRaw(ctx, j.Store, SourceGoldfish, key)
				if err != nil {
					return err
				}
				if has || fetched >= share {
					continue
				}
				deck, err := j.Fetch.Get(ctx, GoldfishDeckURL(j.GoldfishBase, id))
				if err != nil {
					if ctx.Err() != nil {
						return err
					}
					rep.FetchErrors[SourceGoldfish]++
					j.Logger.Warn("mtggoldfish deck page did not fetch", "deck", id, "err", err)
					continue
				}
				fetched++
				rep.Pages[SourceGoldfish]++
				l, err := ParseGoldfishDeck(id, deck)
				if err != nil {
					rep.Failures[SourceGoldfish]++
					j.Logger.Warn("mtggoldfish deck page did not parse", "deck", id, "err", err)
					if err := PutRaw(ctx, j.Store, SourceGoldfish+"-failed", key, deck); err != nil {
						return err
					}
					continue
				}
				if err := PutRaw(ctx, j.Store, SourceGoldfish, key, deck); err != nil {
					return err
				}
				if l != nil {
					lists = append(lists, *l)
				}
			}
			n, err := MergeLists(ctx, j.Store, lists)
			if err != nil {
				return err
			}
			rep.Lists[SourceGoldfish] += n
			if !more {
				break
			}
		}
	}
	return nil
}

// reparseGoldfish re-reads every stored deck page and fetches nothing
// (M-6).
func (j *Job) reparseGoldfish(ctx context.Context, rep *Report) error {
	prefix := RawPrefix + SourceGoldfish + "/"
	names, err := j.Store.List(ctx, prefix)
	if err != nil {
		return err
	}
	var lists []List
	flush := func() error {
		n, err := MergeLists(ctx, j.Store, lists)
		rep.Lists[SourceGoldfish] += n
		lists = nil
		return err
	}
	for _, name := range names {
		key := strings.TrimSuffix(strings.TrimPrefix(name, prefix), ".gz")
		page, ok, err := GetRaw(ctx, j.Store, SourceGoldfish, key)
		if err != nil || !ok {
			return err
		}
		rep.Pages[SourceGoldfish]++
		l, err := ParseGoldfishDeck(strings.TrimPrefix(key, "deck_"), page)
		if err != nil {
			rep.Failures[SourceGoldfish]++
			continue
		}
		if l != nil {
			lists = append(lists, *l)
		}
		if len(lists) >= 50 {
			if err := flush(); err != nil {
				return err
			}
		}
	}
	return flush()
}
