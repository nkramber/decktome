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
	// MaxPages caps the event pages one run fetches, newest first. The
	// default is 200, about four minutes at the fetcher's gap.
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
	// MTGOBase, MTGJSONBase, EDHRECBase, and CEDHDBURL override the
	// sites, for the tests.
	MTGOBase, MTGJSONBase, EDHRECBase, CEDHDBURL string
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
	return rep, nil
}

// runMTGO reads the month pages back MTGOMonths, then every event page
// of a covered format the store lacks, newest first, up to MaxPages.
func (j *Job) runMTGO(ctx context.Context, rep *Report) error {
	if j.Reparse {
		return j.reparseMTGO(ctx, rep)
	}
	now := j.Now().UTC()
	var slugs []string
	seen := map[string]bool{}
	for i := 0; i < j.MTGOMonths; i++ {
		month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, -i, 0)
		url := fmt.Sprintf("%s/decklists/%04d/%02d", j.MTGOBase, month.Year(), int(month.Month()))
		page, err := j.Fetch.Get(ctx, url)
		if err != nil {
			if NotFound(err) {
				continue
			}
			return err
		}
		for _, slug := range ParseMTGOMonth(page) {
			if MTGOSlugFormat(slug) != "" && !seen[slug] {
				seen[slug] = true
				slugs = append(slugs, slug)
			}
		}
	}
	fetched := 0
	for _, slug := range slugs {
		if fetched >= j.MaxPages {
			rep.Skipped[SourceMTGO] = fmt.Sprintf("the page cap of %d held, and more pages wait", j.MaxPages)
			break
		}
		if ok, err := HasRaw(ctx, j.Store, SourceMTGO, slug); err != nil {
			return err
		} else if ok {
			continue
		}
		page, err := j.Fetch.Get(ctx, j.MTGOBase+"/decklist/"+slug)
		if err != nil {
			if ctx.Err() != nil {
				return err
			}
			if NotFound(err) {
				j.Logger.Warn("mtgo event page is gone", "slug", slug)
				continue
			}
			// One page the site would not serve is not the whole month.
			// The store lacks it still, so the next run reads it again.
			rep.FetchErrors[SourceMTGO]++
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
				return err
			}
			continue
		}
		if err := PutRaw(ctx, j.Store, SourceMTGO, slug, page); err != nil {
			return err
		}
		n, err := MergeLists(ctx, j.Store, lists)
		if err != nil {
			return err
		}
		rep.Lists[SourceMTGO] += n
	}
	return nil
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
// average deck, when the last read is older than EDHRECDays.
func (j *Job) runEDHREC(ctx context.Context, rep *Report) error {
	now := j.Now().UTC()
	day := now.Format("2006-01-02")
	last, err := LatestCommandersDay(ctx, j.Store)
	if err != nil {
		return err
	}
	if last != "" && last != day {
		t, err := time.Parse("2006-01-02", last)
		if err == nil && now.Sub(t) < time.Duration(j.EDHRECDays)*24*time.Hour {
			if has, _ := j.edhrecReadOn(ctx, last); has {
				rep.Skipped[SourceEDHREC] = "read on " + last
				return nil
			}
		}
	}
	if has, err := j.edhrecReadOn(ctx, day); err != nil {
		return err
	} else if has {
		rep.Skipped[SourceEDHREC] = "read today"
		return nil
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
