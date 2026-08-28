// Package gcpenv holds the environment reads the api and the worker
// share: the GCP project, the snapshot store, and the logger shape that
// Cloud Logging reads.
//
// Variables it reads: PROJECT_ID, GOOGLE_CLOUD_PROJECT,
// FIRESTORE_EMULATOR_HOST, STORAGE_EMULATOR_HOST, CARDS_SNAPSHOT_DIR,
// CARDS_BUCKET, and K_SERVICE.
package gcpenv

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"os"

	"cloud.google.com/go/storage"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// LocalProject is the project name every emulator mode uses.
const LocalProject = "mtg-local"

// ProjectID resolves the GCP project. Cloud Run sets neither PROJECT_ID
// nor GOOGLE_CLOUD_PROJECT by default, so a missing value is fatal
// unless an emulator host or a snapshot directory marks local mode (C-17).
func ProjectID() (string, error) {
	if p := EnvOr("PROJECT_ID", os.Getenv("GOOGLE_CLOUD_PROJECT")); p != "" {
		return p, nil
	}
	if os.Getenv("FIRESTORE_EMULATOR_HOST") != "" || os.Getenv("STORAGE_EMULATOR_HOST") != "" ||
		os.Getenv("CARDS_SNAPSHOT_DIR") != "" {
		return LocalProject, nil
	}
	return "", errors.New("PROJECT_ID or GOOGLE_CLOUD_PROJECT must be set outside emulator mode")
}

// OnCloudRun reports whether the process runs on Cloud Run, which sets
// K_SERVICE.
func OnCloudRun() bool { return os.Getenv("K_SERVICE") != "" }

// SnapshotStore picks the snapshot backend. CARDS_SNAPSHOT_DIR selects a
// local directory (offline mode, tests). Default: the GCS bucket, which
// STORAGE_EMULATOR_HOST routes to fake-gcs-server in local mode. The
// returned client is nil for the directory backend, and the caller
// closes it otherwise.
//
// createBucket makes the bucket in emulator mode, because fake-gcs-server
// starts empty. In production the bucket is infrastructure, so only the
// worker asks for it.
func SnapshotStore(ctx context.Context, project string, createBucket bool, logger *slog.Logger) (cards.Store, *storage.Client, error) {
	if dir := os.Getenv("CARDS_SNAPSHOT_DIR"); dir != "" {
		return cards.DirStore{Root: dir}, nil, nil
	}
	// WithJSONReads: the SDK's default XML download path 404s on
	// fake-gcs-server's filesystem backend (encoded-slash object names).
	// JSON reads work on fake-gcs and on real GCS.
	client, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		return nil, nil, err
	}
	bucket := EnvOr("CARDS_BUCKET", project+"-cards")
	if createBucket && os.Getenv("STORAGE_EMULATOR_HOST") != "" {
		if err := client.Bucket(bucket).Create(ctx, project, nil); err != nil {
			// An existing bucket lands here.
			logger.Info("bucket create skipped", "bucket", bucket, "note", err.Error())
		}
	}
	return cards.NewGCSStore(client, bucket), client, nil
}

// EnvOr reads key, or returns fallback when it is empty.
func EnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// NewLogger builds the JSON logger. On Cloud Run the level and message
// keys become severity and message, which Cloud Logging reads (L-18).
// Elsewhere the slog defaults stay, so a local log stays plain.
func NewLogger(w io.Writer) *slog.Logger {
	opts := &slog.HandlerOptions{}
	if OnCloudRun() {
		opts.ReplaceAttr = cloudLoggingAttr
	}
	return slog.New(slog.NewJSONHandler(w, opts))
}

// cloudLoggingAttr renames the two keys Cloud Logging parses. The level
// names map onto the LogSeverity enum.
func cloudLoggingAttr(groups []string, a slog.Attr) slog.Attr {
	if len(groups) > 0 {
		return a
	}
	switch a.Key {
	case slog.LevelKey:
		level, _ := a.Value.Any().(slog.Level)
		return slog.String("severity", severity(level))
	case slog.MessageKey:
		return slog.String("message", a.Value.String())
	}
	return a
}

func severity(l slog.Level) string {
	switch {
	case l >= slog.LevelError:
		return "ERROR"
	case l >= slog.LevelWarn:
		return "WARNING"
	case l >= slog.LevelInfo:
		return "INFO"
	}
	return "DEBUG"
}
