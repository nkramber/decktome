package cards

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"
)

// F-1: ban lists drift fast, and decks must be legal on the query date.
// The refresh loop checks the bulk catalog on a normal cadence, and a
// tight cadence around an announcement until the new data lands.

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
	return cal, nil
}

// Announcements is the package calendar. The error surfaces on first use.
func Announcements() (AnnouncementCalendar, error) { return loadAnnouncements() }

const (
	// NormalCheckInterval is the refresh cadence on a normal day.
	// One catalog call per hour is far inside the Scryfall limits (F-3).
	NormalCheckInterval = 1 * time.Hour
	// FastCheckInterval applies while an announcement is not yet in the
	// stored snapshot.
	FastCheckInterval = 15 * time.Minute
)

// pendingAnnouncement returns the newest announcement date that is in the
// past but not yet covered by the snapshot, and true when one exists.
func (c AnnouncementCalendar) pendingAnnouncement(now, snapshotAsOf time.Time) (time.Time, bool) {
	for i := len(c.dates) - 1; i >= 0; i-- {
		d := c.dates[i]
		if d.After(now) {
			continue
		}
		// Covered when the snapshot was published on or after the date.
		if snapshotAsOf.Before(d) {
			return d, true
		}
		return time.Time{}, false
	}
	return time.Time{}, false
}

// CheckInterval picks the refresh cadence. Fast while an announcement has
// happened and the snapshot does not yet reflect a same-day-or-later
// publish. Normal otherwise.
func (c AnnouncementCalendar) CheckInterval(now, snapshotAsOf time.Time) time.Duration {
	if _, pending := c.pendingAnnouncement(now, snapshotAsOf); pending {
		return FastCheckInterval
	}
	return NormalCheckInterval
}

// LagHours measures M-2: the hours between an announcement and the
// snapshot that covers it. It returns false when the snapshot is not the
// first one after an announcement within the last 7 days.
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
