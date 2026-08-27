#!/usr/bin/env bash
# Start the local stack with no cloud credentials (PR-0c, D-9).
# One Ctrl-C stops every process. State persists in .local/.
set -u
# Job control puts each background job in its own process group. The
# cleanup then kills a whole group, so the binary that `go run` starts
# dies with `go run`. Each job reads stdin from /dev/null on purpose:
# with job control a background read from the terminal stops the job.
set -m
cd "$(dirname "$0")/.." || exit 1

FAKE_GCS_VERSION="v1.56.1"
export PROJECT_ID="mtg-local"
export FIRESTORE_EMULATOR_HOST="127.0.0.1:8281"
export FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:9199"
export STORAGE_EMULATOR_HOST="http://127.0.0.1:4443"
export CARDS_BUCKET="mtg-local-cards"
export CARDS_RELOAD_SECONDS="15"
# Local dev opts out of the key requirement. Without keys the API uses the
# fixture fake for every LLM role. .env (gitignored) can still set keys.
export LLM_REQUIRE_KEYS="0"
if [ -f .env ]; then set -a; . ./.env; set +a; fi

mkdir -p .local/firestore .local/gcs
# Ports deliberately avoid the Wallabee dev stack on this machine:
# its Firestore emulator holds 8181 and 4000, its flash API holds 8080.
pids=()
cleanup() {
  # Phase 1: TERM the top-level processes only. The firebase CLI must keep
  # its java emulator alive to serve the --export-on-exit request; a TERM
  # to the whole tree kills the emulator first and the export fails.
  for p in "${pids[@]}"; do kill -TERM "$p" 2>/dev/null; done
  # The firestore emulator writes its --export-on-exit data during shutdown.
  # It stages into firebase-export-* in the cwd, then moves to .local/firestore.
  # A kill during that window leaves empty firebase-export-* litter. Wait up
  # to 30 seconds for clean exits before the hard sweep.
  for _ in $(seq 1 30); do
    alive=0
    for p in "${pids[@]}"; do kill -0 "$p" 2>/dev/null && alive=1; done
    [ "$alive" -eq 0 ] && break
    sleep 1
  done
  # Hard sweep. Each pid leads its own process group (set -m), so a kill
  # of the group also reaches the compiled binary under `go run`.
  for p in "${pids[@]}"; do kill -KILL -- "-$p" 2>/dev/null; kill -KILL "$p" 2>/dev/null; done
  # A process that left its group survives the sweep. Sweep OUR ports only (Wallabee owns 8080, 8181, 4000, 5173).
  # 4490 is the emulator hub and 4590 is the emulator logging port (firebase.json).
  lsof -ti :8281 -ti :9199 -ti :4443 -ti :8090 -ti :5180 -ti :4100 -ti :4490 -ti :4590 2>/dev/null | xargs kill -KILL 2>/dev/null
  # An abandoned export staging dir is always empty. Remove it.
  rmdir firebase-export-* 2>/dev/null
  wait 2>/dev/null
  return 0
}
# INT/TERM: run cleanup once, drop the EXIT trap so it does not run twice,
# then exit with the conventional 128+SIGINT code.
trap cleanup EXIT
trap 'cleanup; trap - EXIT; exit 130' INT TERM

echo "==> firestore + auth emulators (:8281, :9199, ui :4100)"
firebase emulators:start --only firestore,auth --project "$PROJECT_ID" \
  --import .local/firestore --export-on-exit .local/firestore </dev/null & pids+=($!)

# Wait for the Firestore port. The API and worker need it at start.
for i in $(seq 1 30); do
  if nc -z 127.0.0.1 8281 2>/dev/null; then break; fi
  if [ "$i" -eq 30 ]; then
    echo "!! firestore emulator did not open :8281 within 30 s. Check java and firebase-tools (make doctor)." >&2
    exit 1
  fi
  sleep 1
done

echo "==> fake-gcs-server (:4443)"
go run github.com/fsouza/fake-gcs-server@"$FAKE_GCS_VERSION" \
  -scheme http -host 127.0.0.1 -port 4443 -filesystem-root .local/gcs </dev/null & pids+=($!)

export PORT="8090"
echo "==> api on :8090"
(cd go && go run ./cmd/api) </dev/null & pids+=($!)
echo "==> worker"
(cd go && go run ./cmd/worker) </dev/null & pids+=($!)
echo "==> web on :5180"
(cd web && pnpm --filter @mtg/web dev) </dev/null & pids+=($!)
wait
