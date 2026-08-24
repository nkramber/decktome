#!/usr/bin/env bash
# Start the local stack. One Ctrl-C stops every process.
# PR-0c adds the Firestore and Auth emulators and fake-gcs-server here.
set -u
cd "$(dirname "$0")/.."
pids=()
cleanup() { for p in "${pids[@]}"; do kill "$p" 2>/dev/null; done; wait 2>/dev/null; }
trap cleanup EXIT INT TERM

echo "==> api on :8080"
(cd go && go run ./cmd/api) & pids+=($!)
echo "==> worker"
(cd go && go run ./cmd/worker) & pids+=($!)
echo "==> web on :5180"
(cd web && pnpm --filter @mtg/web dev) & pids+=($!)
wait
