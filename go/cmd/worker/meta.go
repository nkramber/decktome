package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gcpenv"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
	"github.com/nkramber/mtg-deck-builder/go/internal/quality"
)

// metaOptions are the -meta flags.
type metaOptions struct {
	months  int
	pages   int
	reparse bool
}

// metaTimeout bounds one meta run: the page cap at the fetcher's gap,
// the deck files of a new MTGJSON version, and the fit.
const metaTimeout = 3 * time.Hour

// metaCommanders is how many of the most built legends of the card
// index the EDHREC read adds to the top pages, about 1,100 commanders
// with the three top pages, once a week (D-489).
const metaCommanders = 1000

// runMeta is the -meta job: every source, then the fit (PR-14B).
//
// Production shape: a Cloud Run job on a daily cron. The Topdeck.gg
// read needs TOPDECK_API_KEY, and the job reads no tournament without
// it and says so in the log.
func runMeta(ctx context.Context, opts metaOptions, logger *slog.Logger) error {
	ctx, cancel := context.WithTimeout(ctx, metaTimeout)
	defer cancel()
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	store, storageClient, err := gcpenv.SnapshotStore(ctx, project, true, logger)
	if err != nil {
		return fmt.Errorf("snapshot store init: %w", err)
	}
	if storageClient != nil {
		defer func() { _ = storageClient.Close() }()
	}
	objects, err := gcpenv.MetaStore(ctx, project, storageClient)
	if err != nil {
		return fmt.Errorf("meta store init: %w", err)
	}
	idx, err := cards.LoadIndex(ctx, store, logger)
	if err != nil {
		return fmt.Errorf("card index: %w", err)
	}
	if idx == nil {
		return fmt.Errorf("no card snapshot is stored, so the meta job can not resolve a list")
	}
	job := &meta.Job{
		Store: objects, Logger: logger, MTGOMonths: opts.months, MaxPages: opts.pages, Reparse: opts.reparse,
		Commanders: topCommanders(idx, metaCommanders),
	}
	if key := os.Getenv(meta.TopdeckKeyEnv); key != "" {
		job.Topdeck, err = meta.NewTopdeck(nil, key, "", logger)
		if err != nil {
			return err
		}
	} else {
		logger.Warn("no " + meta.TopdeckKeyEnv + ", so the meta job reads no Topdeck.gg tournament")
	}
	rep, err := job.Run(ctx)
	if err != nil {
		return err
	}
	logMetaReport(logger, rep)
	model, fit, err := quality.Refit(ctx, objects, idx, time.Now().UTC(), logger)
	if err != nil {
		return fmt.Errorf("quality fit: %w", err)
	}
	for word, fr := range fit.Formats {
		logger.Info("quality fit", "format", word, "read", fr.Read, "used", fr.Used, "unusable", fr.Unusable, "out_of_pool", fr.OutOfPool,
			"synthetic", fr.Synthetic, "holdout", fr.Holdout.Lists,
			"great_over_baseline", fr.Holdout.GreatOverBaseline.Share(),
			"baseline_over_bad", fr.Holdout.BaselineOverBad.Share(), "accuracy", fr.Holdout.Accuracy)
	}
	logger.Info("meta job done", "model", model.Version)
	return nil
}

// logMetaReport writes one line per source.
func logMetaReport(logger *slog.Logger, rep *meta.Report) {
	for _, source := range []string{meta.SourceMTGO, meta.SourceMTGJSON, meta.SourceCEDHDB, meta.SourceTopdeck, meta.SourceEDHREC} {
		logger.Info("meta source", "source", source, "pages", rep.Pages[source], "failures", rep.Failures[source],
			"fetch_errors", rep.FetchErrors[source], "lists", rep.Lists[source], "skipped", rep.Skipped[source])
	}
	if rep.Precons > 0 {
		logger.Info("precon table stored", "version", rep.PreconsVersion, "products", rep.Precons)
	}
	for _, e := range rep.Errors {
		logger.Error("meta source error", "err", e)
	}
}

// topCommanders names the n most built legends of the index by EDHREC
// rank, so the EDHREC read covers the commanders a request can reach.
func topCommanders(idx *cards.Index, n int) []string {
	type ranked struct {
		name string
		rank int32
	}
	var out []ranked
	for _, c := range idx.All() {
		if !c.GetCanBeCommander() || c.GetEdhrecRank() == 0 {
			continue
		}
		out = append(out, ranked{c.GetName(), c.GetEdhrecRank()})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].rank < out[j].rank })
	names := make([]string, 0, n)
	for i := 0; i < len(out) && i < n; i++ {
		names = append(names, out[i].name)
	}
	return names
}
