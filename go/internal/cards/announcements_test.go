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

func calFor(dates ...string) AnnouncementCalendar {
	c := AnnouncementCalendar{}
	for _, d := range dates {
		c.dates = append(c.dates, date(d))
	}
	return c
}

func TestEmbeddedCalendarLoads(t *testing.T) {
	cal, err := loadAnnouncements()
	if err != nil {
		t.Fatalf("embedded calendar: %v", err)
	}
	if len(cal.dates) == 0 {
		t.Fatal("embedded calendar is empty")
	}
}

func TestCheckInterval(t *testing.T) {
	cal := calFor("2026-10-12")
	tests := []struct {
		name     string
		now      string
		snapshot string
		want     time.Duration
	}{
		{"before announcement day", "2026-10-11", "2026-10-10", NormalCheckInterval},
		{"announcement day, stale snapshot", "2026-10-12", "2026-10-11", FastCheckInterval},
		{"day after, still stale", "2026-10-13", "2026-10-11", FastCheckInterval},
		{"announcement covered", "2026-10-13", "2026-10-12", NormalCheckInterval},
		{"long after, covered", "2026-11-01", "2026-10-31", NormalCheckInterval},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cal.CheckInterval(date(tt.now), date(tt.snapshot))
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
