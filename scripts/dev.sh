#!/usr/bin/env bash
# Start the local stack with no cloud credentials (PR-0c, D-9).
# One Ctrl-C stops every process. State persists in .local/.
set -u
cd "$(dirname "$0")/.."

FAKE_GCS_VERSION="v1.56.1"
export PROJECT_ID="mtg-local"
export FIRESTORE_EMULATOR_HOST="127.0.0.1:8281"
export FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:9199"
export STORAGE_EMULATOR_HOST="http://127.0.0.1:4443"

mkdir -p .local/firestore .local/gcs
# Ports deliberately avoid the Wallabee dev stack on this machine:
# its Firestore emulator holds 8181 and 4000, its flash API holds 8080.
pids=()
kill_tree() { pkill -TERM -P "$1" 2>/dev/null; kill -TERM "$1" 2>/dev/null; }
cleanup() {
  for p in "${pids[@]}"; do kill_tree "$p"; done
  sleep 1
  for p in "${pids[@]}"; do pkill -KILL -P "$p" 2>/dev/null; kill -KILL "$p" 2>/dev/null; done
  # go run leaves grandchildren behind. Sweep OUR ports only (Wallabee owns 8080, 8181, 4000, 5173).
  lsof -ti :8281 -ti :9199 -ti :4443 -ti :8090 -ti :5180 -ti :4100 2>/dev/null | xargs kill -KILL 2>/dev/null
  wait 2>/dev/null
  return 0
}
trap cleanup EXIT INT TERM

echo "==> firestore + auth emulators (:8281, :9199, ui :4100)"
firebase emulators:start --only firestore,auth --project "$PROJECT_ID" \
  --import .local/firestore --export-on-exit .local/firestore & pids+=($!)

echo "==> fake-gcs-server (:4443)"
go run github.com/fsouza/fake-gcs-server@"$FAKE_GCS_VERSION" \
  -scheme http -host 127.0.0.1 -port 4443 -filesystem-root .local/gcs & pids+=($!)

export PORT="8090"
echo "==> api on :8090"
(cd go && go run ./cmd/api) & pids+=($!)
echo "==> worker"
(cd go && go run ./cmd/worker) & pids+=($!)
echo "==> web on :5180"
(cd web && pnpm --filter @mtg/web dev) & pids+=($!)
wait
