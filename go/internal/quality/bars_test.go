package quality

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/nkramber/decktome/go/internal/meta"
)

// passing answers a fit that clears every bar, with the share of one
// pair changed by edit.
func passing(edit func(r *FitReport, m *Model)) (*Model, *FitReport) {
	m := &Model{Version: "20260925T000000Z", Formats: map[string]*FormatModel{}}
	r := &FitReport{Formats: map[string]*FormatReport{}}
	for _, w := range BarFormats {
		m.Formats[w] = &FormatModel{}
		fr := &FormatReport{}
		fr.Holdout.GreatOverBaseline = PairShare{Pairs: 100, Wins: 97}
		fr.Holdout.OwnByDefect = map[string]PairShare{
			DefectLands: {Pairs: 50, Wins: 49}, DefectCurve: {Pairs: 50, Wins: 48},
			DefectSynergy: {Pairs: 100, Wins: 80},
		}
		r.Formats[w] = fr
	}
	if edit != nil {
		edit(r, m)
	}
	return m, r
}

// TestPublishKeepsTheModelOnABarFailure: a daily fit that fails a bar
// or drops a format never becomes the newest model (D-927).
func TestPublishKeepsTheModelOnABarFailure(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	tests := []struct {
		name   string
		edit   func(r *FitReport, m *Model)
		stored bool
	}{
		{name: "every bar passes", stored: true},
		{name: "the bar values pass", edit: func(r *FitReport, _ *Model) {
			r.Formats[meta.FormatModern].Holdout.GreatOverBaseline = PairShare{Pairs: 10, Wins: 9}
			r.Formats[meta.FormatModern].Holdout.OwnByDefect = map[string]PairShare{
				DefectLands: {Pairs: 20, Wins: 19}, DefectSynergy: {Pairs: 20, Wins: 15},
			}
		}, stored: true},
		// The Commander fit of 2026-09-26 read 0.94 of 827 over every axis,
		// under the one bar of 0.95. Each axis but synergy read 572 of 572,
		// and synergy 205 of 255 (D-957).
		{name: "the commander fit of 2026-09-26", edit: func(r *FitReport, _ *Model) {
			r.Formats[meta.FormatCommander].Holdout.OwnByDefect = map[string]PairShare{
				DefectLands: {Pairs: 196, Wins: 196}, DefectCurve: {Pairs: 196, Wins: 196},
				DefectColors: {Pairs: 180, Wins: 180}, DefectSynergy: {Pairs: 255, Wins: 205},
			}
		}, stored: true},
		{name: "great over precon under its bar", edit: func(r *FitReport, _ *Model) {
			r.Formats[meta.FormatCommander].Holdout.GreatOverBaseline = PairShare{Pairs: 100, Wins: 89}
		}},
		{name: "precon over own copy under its bar", edit: func(r *FitReport, _ *Model) {
			r.Formats[meta.FormatStandard].Holdout.OwnByDefect[DefectCurve] = PairShare{Pairs: 50, Wins: 45}
		}},
		{name: "precon over own synergy copy under its bar", edit: func(r *FitReport, _ *Model) {
			r.Formats[meta.FormatCommander].Holdout.OwnByDefect[DefectSynergy] = PairShare{Pairs: 100, Wins: 74}
		}},
		{name: "a bar with no pair", edit: func(r *FitReport, _ *Model) {
			r.Formats[meta.FormatModern].Holdout.OwnByDefect = nil
		}},
		{name: "no synergy pair", edit: func(r *FitReport, _ *Model) {
			delete(r.Formats[meta.FormatCommander].Holdout.OwnByDefect, DefectSynergy)
		}},
		{name: "a format drops out of the lists", edit: func(r *FitReport, _ *Model) {
			delete(r.Formats, meta.FormatStandard)
		}},
		{name: "a format drops out of the model", edit: func(_ *FitReport, m *Model) {
			delete(m.Formats, meta.FormatModern)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			store := meta.DirObjects{Root: t.TempDir()}
			m, r := passing(tt.edit)
			err := publish(ctx, store, m, r, quiet)
			if tt.stored != (err == nil) {
				t.Fatalf("publish err = %v, want stored %v", err, tt.stored)
			}
			if !tt.stored && !errors.Is(err, ErrBarFailed) {
				t.Errorf("err = %v, want ErrBarFailed", err)
			}
			got, err := meta.LatestModel(ctx, store)
			if err != nil {
				t.Fatal(err)
			}
			if (got == m.Version) != tt.stored {
				t.Errorf("newest model = %q, want stored %v", got, tt.stored)
			}
		})
	}
}
