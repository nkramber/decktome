package usage

import (
	"testing"
	"time"
)

func TestMonthAndResetDate(t *testing.T) {
	at := time.Date(2026, 9, 6, 23, 30, 0, 0, time.FixedZone("west", -7*3600))
	if got := Month(at); got != "2026-09" {
		t.Errorf("Month = %q, want the UTC month", got)
	}
	if got, err := ResetDate("2026-12"); err != nil || got != "2027-01-01" {
		t.Errorf("ResetDate = %q, %v, want the first of the next month", got, err)
	}
	if _, err := ResetDate("nope"); err == nil {
		t.Error("a bad month must fail")
	}
}
