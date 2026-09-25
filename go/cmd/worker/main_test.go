package main

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/cards"
)

func TestSkipRefresh(t *testing.T) {
	now := time.Date(2026, 8, 24, 10, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		latestAt time.Time
		pending  bool
		want     bool
	}{
		{"fresh snapshot, nothing pending", now.Add(-30 * time.Minute), false, true},
		{"fresh snapshot, announcement pending", now.Add(-30 * time.Minute), true, false},
		{"snapshot exactly 60 minutes old", now.Add(-60 * time.Minute), false, false},
		{"old snapshot", now.Add(-25 * time.Hour), false, false},
		{"empty store", time.Time{}, false, false},
		{"empty store, pending", time.Time{}, true, false},
		{"future snapshot counts as fresh", now.Add(5 * time.Minute), false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skipRefresh(now, tt.latestAt, tt.pending); got != tt.want {
				t.Errorf("skipRefresh = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSnapshotNotice is REV-011 of the review of 2026-09-24. A refresh
// that failed again and again reached no one, so a ban could miss the app
// with no signal. A stale snapshot and an unreadable store now alert, at
// the first run of every sixth UTC hour alone.
func TestSnapshotNotice(t *testing.T) {
	at := time.Date(2026, 10, 12, 9, 1, 0, 0, time.UTC)
	version := cards.VersionFor(at)
	failed := errors.New("storage: permission denied")
	tests := []struct {
		name   string
		latest string
		now    time.Time
		err    error
		want   string
	}{
		{"fresh", version, at.Add(5 * time.Hour), nil, ""},
		{"26 hours at an alert hour, a failed run", version, time.Date(2026, 10, 13, 12, 0, 5, 0, time.UTC), failed, ""},
		{"stale at an alert hour", version, time.Date(2026, 10, 13, 18, 0, 5, 0, time.UTC), nil, "32 hours old"},
		{"stale with the error", version, time.Date(2026, 10, 13, 18, 0, 5, 0, time.UTC), failed, "permission denied"},
		{"stale between alert hours", version, time.Date(2026, 10, 13, 19, 0, 5, 0, time.UTC), nil, ""},
		{"stale at minute 15", version, time.Date(2026, 10, 13, 18, 15, 0, 0, time.UTC), nil, ""},
		{"no store read", "", time.Date(2026, 10, 13, 0, 0, 5, 0, time.UTC), failed, "No card snapshot age"},
		{"empty store and no error", "", time.Date(2026, 10, 13, 0, 0, 5, 0, time.UTC), nil, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, ok := snapshotNotice(tt.latest, tt.now, tt.err)
			if ok != (tt.want != "") {
				t.Fatalf("sent = %v, want %v: %+v", ok, tt.want != "", n)
			}
			if ok && !strings.Contains(n.Message, tt.want) {
				t.Errorf("message = %q, want %q in it", n.Message, tt.want)
			}
		})
	}
}
