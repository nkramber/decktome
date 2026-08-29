package cards

import (
	"testing"
	"time"
)

func date(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t.UTC()
}

func stamp(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t.UTC()
}

func calFor(dates ...string) AnnouncementCalendar {
	c := AnnouncementCalendar{}
	for _, d := range dates {
		c.dates = append(c.dates, date(d))
	}
	return c
}

// TestCalendarSortsOnLoad covers the sort the doc promised. Pending
// walks from the end and stops at the first date that decides, so an
// unsorted list hid a newer announcement behind an older one.
func TestCalendarSortsOnLoad(t *testing.T) {
	cal := calFor("2026-08-20", "2026-06-01", "2026-07-15")
	cal.sortDates()
	for i := 1; i < len(cal.dates); i++ {
		if cal.dates[i].Before(cal.dates[i-1]) {
			t.Fatalf("dates not ascending: %v", cal.dates)
		}
	}
	got, ok := cal.Pending(date("2026-08-21"), date("2026-08-01"))
	if !ok || !got.Equal(date("2026-08-20")) {
		t.Errorf("Pending = %v, %v, want 2026-08-20", got, ok)
	}
	loaded, err := loadAnnouncements()
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i < len(loaded.dates); i++ {
		if loaded.dates[i].Before(loaded.dates[i-1]) {
			t.Fatalf("embedded calendar not ascending after load: %v", loaded.dates)
		}
	}
}

// TestEmbeddedCalendarLoads also fails when every date is in the past.
// The calendar then covers no announcement, and the refresh loop never
// enters its fast interval. The fix is a new date in the file.
func TestEmbeddedCalendarLoads(t *testing.T) {
	cal, err := loadAnnouncements()
	if err != nil {
		t.Fatalf("embedded calendar: %v", err)
	}
	if len(cal.dates) == 0 {
		t.Fatal("embedded calendar is empty")
	}
	now := time.Now().UTC()
	future := false
	for _, d := range cal.dates {
		if d.After(now) {
			future = true
		}
	}
	if !future {
		t.Fatalf("internal/cards/announcement_dates.json holds no date after %s: add the next Wizards announcement date",
			now.Format("2006-01-02"))
	}
}

// TestCheckInterval covers C-2: only a legality diff counts as coverage
// (D-47). A diff on the announcement date itself covers it, because the
// date is UTC midnight and Wizards posts later that day.
func TestCheckInterval(t *testing.T) {
	cal := calFor("2026-10-12")
	tests := []struct {
		name     string
		now      string
		lastDiff string
		want     time.Duration
	}{
		{"before announcement day", "2026-10-11T12:00:00Z", "2026-10-01T09:00:00Z", NormalCheckInterval},
		{"announcement day, no diff yet", "2026-10-12T20:00:00Z", "2026-10-01T09:00:00Z", FastCheckInterval},
		{"day after, still no diff", "2026-10-13T12:00:00Z", "2026-10-01T09:00:00Z", FastCheckInterval},
		{"diff landed the day after", "2026-10-13T12:00:00Z", "2026-10-13T09:01:00Z", NormalCheckInterval},
		{"diff landed on the date itself", "2026-10-12T20:00:00Z", "2026-10-12T09:01:00Z", NormalCheckInterval},
		{"window over, no diff", "2026-10-16T12:00:00Z", "2026-10-01T09:00:00Z", NormalCheckInterval},
		{"no diff ever, before any date", "2026-10-11T12:00:00Z", "0001-01-01T00:00:00Z", NormalCheckInterval},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cal.CheckInterval(stamp(tt.now), stamp(tt.lastDiff))
			if got != tt.want {
				t.Errorf("CheckInterval = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLagHours(t *testing.T) {
	cal := calFor("2026-08-10", "2026-10-12")
	t.Run("snapshot two days after announcement", func(t *testing.T) {
		ann, hours, ok := cal.LagHours(date("2026-10-14"))
		if !ok || !ann.Equal(date("2026-10-12")) || hours != 48 {
			t.Errorf("got ann=%v hours=%v ok=%v", ann, hours, ok)
		}
	})
	t.Run("snapshot on announcement day at 09:01Z", func(t *testing.T) {
		_, hours, ok := cal.LagHours(stamp("2026-10-12T09:01:00Z"))
		if !ok || hours < 9 || hours > 9.1 {
			t.Errorf("got hours=%v ok=%v", hours, ok)
		}
	})
	t.Run("snapshot far from any announcement", func(t *testing.T) {
		if _, _, ok := cal.LagHours(date("2026-09-20")); ok {
			t.Error("want ok=false far from announcements")
		}
	})
	t.Run("snapshot before all announcements", func(t *testing.T) {
		if _, _, ok := cal.LagHours(date("2026-08-01")); ok {
			t.Error("want ok=false before all announcements")
		}
	})
}
