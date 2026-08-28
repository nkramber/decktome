package health

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

type fakeSource struct{ idx *cards.Index }

func (f fakeSource) Current() *cards.Index { return f.idx }

func TestCheck(t *testing.T) {
	asOf := time.Date(2026, 8, 24, 9, 0, 0, 0, time.UTC)
	now := asOf.Add(90 * time.Minute)
	tests := []struct {
		name     string
		version  string
		source   IndexSource
		wantSnap string
		wantAge  float64
		wantStat string
		ready    bool
	}{
		{name: "dev build, nil source", version: "dev", wantSnap: "none", wantAge: -1, wantStat: "starting"},
		{name: "tagged build, no index yet", version: "v0.1.0", source: fakeSource{}, wantSnap: "none", wantAge: -1, wantStat: "starting"},
		{name: "index loaded", version: "v0.1.0", source: fakeSource{idx: cards.NewIndex(nil, nil, nil, asOf)},
			wantSnap: "2026-08-24T09:00:00Z", wantAge: 1.5, wantStat: "ok", ready: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.version, tt.source)
			s.now = func() time.Time { return now }
			res, err := s.Check(context.Background(), connect.NewRequest(&mtgv1.CheckRequest{}))
			if err != nil {
				t.Fatalf("Check returned error: %v", err)
			}
			if got := res.Msg.GetStatus(); got != tt.wantStat {
				t.Errorf("status = %q, want %q", got, tt.wantStat)
			}
			if got := s.Ready(); got != tt.ready {
				t.Errorf("ready = %v, want %v", got, tt.ready)
			}
			if got := res.Msg.GetVersion(); got != tt.version {
				t.Errorf("version = %q, want %q", got, tt.version)
			}
			if got := res.Msg.GetCardSnapshot(); got != tt.wantSnap {
				t.Errorf("card_snapshot = %q, want %q", got, tt.wantSnap)
			}
			if got := res.Msg.GetCardSnapshotAgeHours(); got != tt.wantAge {
				t.Errorf("card_snapshot_age_hours = %v, want %v", got, tt.wantAge)
			}
		})
	}
}
