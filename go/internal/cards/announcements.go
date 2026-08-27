package cards

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// F-1: ban lists drift fast, and decks must be legal on the query date.
// The calendar sets the check cadence: normal on a normal day, fast from
// an announcement date until a legality diff lands (C-2). Coverage is a
// fact about the data, not the clock: a snapshot covers an announcement
// only when its legalities differ from the previous snapshot.

//go:embed announcement_dates.json
var announcementJSON []byte

// AnnouncementCalendar is the parsed, sorted announcement date list.
type AnnouncementCalendar struct {
	dates []time.Time
}

// loadAnnouncements parses the embedded calendar. A bad file fails loudly
// at startup, never silently at refresh time.
func loadAnnouncements() (AnnouncementCalendar, error) {
	var raw struct {
		Dates []string `json:"dates"`
	}
	if err := json.Unmarshal(announcementJSON, &raw); err != nil {
		return AnnouncementCalendar{}, fmt.Errorf("announcement_dates.json: %w", err)
	}
	cal := AnnouncementCalendar{}
	for _, d := range raw.Dates {
		t, err := time.Parse("2006-01-02", d)
		if err != nil {
			return AnnouncementCalendar{}, fmt.Errorf("announcement_dates.json: bad date %q: %w", d, err)
		}
		cal.dates = append(cal.dates, t.UTC())
	}
	cal.sortDates()
	return cal, nil
}

// sortDates puts the dates in ascending order. Pending and LagHours walk
// the list from the end and stop at the first date that decides, so an
// unsorted file would make them skip a newer announcement.
func (c *AnnouncementCalendar) sortDates() {
	sort.Slice(c.dates, func(i, j int) bool { return c.dates[i].Before(c.dates[j]) })
}

// Announcements is the package calendar. The error surfaces on first use.
func Announcements() (AnnouncementCalendar, error) { return loadAnnouncements() }

const (
	// NormalCheckInterval is the refresh cadence on a normal day.
	// One catalog call per hour is far inside the Scryfall limits (F-3).
	NormalCheckInterval = 1 * time.Hour
	// FastCheckInterval applies from an announcement date until a
	// legality diff lands, for at most FastWindow.
	FastCheckInterval = 15 * time.Minute
	// FastWindow caps the fast path. Scryfall has never lagged an
	// announcement by three days. After that the normal cadence resumes.
	FastWindow = 3 * 24 * time.Hour
)

// Pending returns the announcement that waits for a legality diff, and
// true when one exists. lastDiff is the AsOf of the last snapshot whose
// legalities changed. The date is UTC midnight, and Wizards posts in the
// US afternoon, so a diff on the date itself still counts as coverage.
func (c AnnouncementCalendar) Pending(now, lastDiff time.Time) (time.Time, bool) {
	for i := len(c.dates) - 1; i >= 0; i-- {
		d := c.dates[i]
		if d.After(now) {
			continue
		}
		if now.Sub(d) > FastWindow {
			return time.Time{}, false
		}
		if lastDiff.Before(d) {
			return d, true
		}
		return time.Time{}, false
	}
	return time.Time{}, false
}

// CheckInterval picks the refresh cadence for the loop mode (make dev).
// Production cadence is Cloud Scheduler, see cmd/worker.
func (c AnnouncementCalendar) CheckInterval(now, lastDiff time.Time) time.Duration {
	if _, pending := c.Pending(now, lastDiff); pending {
		return FastCheckInterval
	}
	return NormalCheckInterval
}

// LagHours measures M-2: the hours between an announcement and the
// snapshot whose legality diff covers it. The caller calls it only when
// a diff landed. It returns false when no announcement precedes the
// snapshot within the last 7 days.
func (c AnnouncementCalendar) LagHours(snapshotAsOf time.Time) (announcement time.Time, hours float64, ok bool) {
	for i := len(c.dates) - 1; i >= 0; i-- {
		d := c.dates[i]
		if snapshotAsOf.Before(d) {
			continue
		}
		if snapshotAsOf.Sub(d) <= 7*24*time.Hour {
			return d, snapshotAsOf.Sub(d).Hours(), true
		}
		return time.Time{}, 0, false
	}
	return time.Time{}, 0, false
}
