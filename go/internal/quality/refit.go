package quality

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
)

// Refit reads every stored list and the newest commander reads, fits
// the model against the index, and stores it as a new version. The
// worker runs it after each meta refresh. It answers the model and the
// report.
func Refit(ctx context.Context, store meta.ObjectStore, idx *cards.Index, now time.Time, logger *slog.Logger) (*Model, *FitReport, error) {
	if logger == nil {
		logger = slog.Default()
	}
	model, rep, err := FitStore(ctx, store, idx, now, logger)
	if err != nil {
		return nil, rep, err
	}
	data, err := model.Encode()
	if err != nil {
		return nil, rep, err
	}
	if err := meta.WriteModel(ctx, store, model.Version, data); err != nil {
		return nil, rep, fmt.Errorf("quality: write model: %w", err)
	}
	logger.Info("quality model stored", "version", model.Version, "bytes", len(data))
	return model, rep, nil
}

// FitStore fits the model over the store's lists and stores nothing.
// The gate reads it, so a gate run leaves no model behind.
func FitStore(ctx context.Context, store meta.ObjectStore, idx *cards.Index, now time.Time, logger *slog.Logger) (*Model, *FitReport, error) {
	if logger == nil {
		logger = slog.Default()
	}
	var lists []meta.List
	for _, word := range []string{meta.FormatCommander, meta.FormatStandard, meta.FormatModern, meta.FormatSixty} {
		got, err := meta.AllLists(ctx, store, word)
		if err != nil {
			return nil, nil, fmt.Errorf("quality: read %s lists: %w", word, err)
		}
		lists = append(lists, got...)
	}
	day, err := meta.LatestCommandersDay(ctx, store)
	if err != nil {
		return nil, nil, err
	}
	var reads []meta.Commander
	if day != "" {
		if reads, err = meta.ReadCommanders(ctx, store, day); err != nil {
			return nil, nil, err
		}
	}
	cfg, err := rules.Load()
	if err != nil {
		return nil, nil, err
	}
	prof, err := profile.New(cfg, func() *cards.TagIndex { return idx.Tags() }, nil)
	if err != nil {
		return nil, nil, err
	}
	builder, err := candidates.New()
	if err != nil {
		return nil, nil, err
	}
	model, rep, err := Fit(ctx, FitInput{
		Index: idx, Profiler: prof, Roles: builder.Roles(idx),
		Lists: lists, Commanders: reads, Now: now, Logger: logger,
	})
	if err != nil {
		return nil, rep, err
	}
	logger.Info("quality model fitted", "version", model.Version, "lists", len(lists), "commanders", len(reads))
	return model, rep, nil
}

// Load reads the newest stored model, or nil when none is stored.
func Load(ctx context.Context, store meta.ObjectStore) (*Model, error) {
	version, err := meta.LatestModel(ctx, store)
	if err != nil || version == "" {
		return nil, err
	}
	data, err := meta.ReadModel(ctx, store, version)
	if err != nil {
		return nil, err
	}
	return Decode(data)
}
