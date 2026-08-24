package main

import (
	"testing"
	"time"
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
