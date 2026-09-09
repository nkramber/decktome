// Command feedback-harvest writes every verdict since the last harvest
// into a dated document and a JSONL file (PR-28a).
//
// The watermark comes from the JSONL files themselves: the newest time
// any harvest wrote is the floor of this one. So the documents are the
// only record of what a harvest read, and no second store can disagree
// with them.
//
// Usage:
//
//	PROJECT_ID=decktome-prod go run ./cmd/feedback-harvest
//	PROJECT_ID=decktome-prod go run ./cmd/feedback-harvest -since 2026-09-01
//	PROJECT_ID=decktome-prod go run ./cmd/feedback-harvest -dry
//
// CAUTION: the files hold what a reader wrote. Keep them off any shared
// page, and put no part of them in an issue or a pull request.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/nkramber/decktome/go/internal/feedback"
	"github.com/nkramber/decktome/go/internal/gcpenv"
	"github.com/nkramber/decktome/go/internal/harvest"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	root := flag.String("root", "..", "the repo root the harvest writes under")
	since := flag.String("since", "", `read every verdict after this day, as 2006-01-02. Empty reads the watermark`)
	limit := flag.Int("limit", 0, "read at most this many verdicts, oldest first. 0 reads every one")
	dry := flag.Bool("dry", false, "read and count, and write no file")
	flag.Parse()

	floor, err := floorOf(*root, *since)
	if err != nil {
		return err
	}

	ctx := context.Background()
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	client, err := firestore.NewClient(ctx, project)
	if err != nil {
		return fmt.Errorf("firestore: %w", err)
	}
	defer func() { _ = client.Close() }()

	// The project is named on every run, so a read of the wrong one is
	// never a silent one.
	fmt.Fprintf(os.Stderr, "project      %s\n", project)
	if floor.IsZero() {
		fmt.Fprintln(os.Stderr, "watermark    none, so this reads every verdict")
	} else {
		fmt.Fprintf(os.Stderr, "watermark    %s\n", floor.Format(time.RFC3339))
	}

	items, err := feedback.NewRepo(client).Since(ctx, floor, *limit)
	if err != nil {
		return err
	}
	recs := make([]harvest.Record, 0, len(items))
	for _, item := range items {
		rec, err := harvest.RecordOf(item)
		if err != nil {
			return err
		}
		recs = append(recs, rec)
	}
	if len(recs) == 0 {
		fmt.Fprintln(os.Stderr, "no verdict since the watermark, so this wrote nothing")
		return nil
	}
	if *dry {
		fmt.Fprintf(os.Stderr, "%d verdict(s) to harvest, and -dry wrote nothing\n", len(recs))
		return nil
	}
	doc, jsonl, err := harvest.Write(*root, time.Now().UTC(), recs)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "wrote %s and %s\n", doc, jsonl)
	return nil
}

// floorOf reads the day the caller named, or the watermark of the files.
func floorOf(root, since string) (time.Time, error) {
	if since == "" {
		return harvest.Watermark(root)
	}
	day, err := time.Parse("2006-01-02", since)
	if err != nil {
		return time.Time{}, fmt.Errorf("-since reads a day as 2006-01-02: %w", err)
	}
	return day.UTC(), nil
}
