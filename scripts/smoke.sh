#!/usr/bin/env bash
# Run the Playwright smoke flow of PR-23 over the emulators and the fake
# provider (D-313, D-552). The flow signs in, uploads the fixture export,
# builds a deck, opens it, and exports it. It calls no provider, and it
# costs nothing. One Ctrl-C stops every process. Nothing persists: the
# emulators start empty and export nothing.
set -u
# Job control puts each background job in its own process group, so the
# cleanup kills a whole group, and the binary that `go run` starts dies
# with `go run`. Each job reads stdin from /dev/null on purpose.
set -m
cd "$(dirname "$0")/.." || exit 1

export PROJECT_ID="mtg-local"
export FIRESTORE_EMULATOR_HOST="127.0.0.1:8281"
export FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:9199"
# The trimmed snapshot of the repo serves the card index (D-521), so the
# flow needs no bucket, no worker, and no download.
export CARDS_SNAPSHOT_DIR="$PWD/go/cmd/deck-gate/testdata/snapshot/scryfall"
# The fake provider serves every role from its fixtures. The keys leave
# the environment, and every role names the fake, so a key in the shell
# or in .env can not turn a smoke run into a paid one. This script reads
# .env on purpose never.
export LLM_REQUIRE_KEYS="0"
unset OPENAI_API_KEY ANTHROPIC_API_KEY
for role in CLASSIFY ASK GENERATE REPAIR JUDGE EVAL REVISE; do
  export "LLM_${role}_PROVIDER=fake"
done
export PORT="8090"

# The stack takes the ports of make dev (scripts/dev.sh), so both can not
# run at once. 4100, 4490, and 4590 are the emulator UI, hub, and logging
# ports of firebase.json.
for port in 8281 9199 8090 5180 4100 4490 4590; do
  if nc -z 127.0.0.1 "$port" 2>/dev/null; then
    echo "!! port $port is in use. Stop make dev, then run make smoke again." >&2
    exit 1
  fi
done

logs=$(mktemp -d)
pids=()
# The two traps below call it.
# shellcheck disable=SC2329
cleanup() {
  for p in "${pids[@]}"; do kill -TERM "$p" 2>/dev/null; done
  for _ in $(seq 1 15); do
    alive=0
    for p in "${pids[@]}"; do kill -0 "$p" 2>/dev/null && alive=1; done
    [ "$alive" -eq 0 ] && break
    sleep 1
  done
  for p in "${pids[@]}"; do kill -KILL -- "-$p" 2>/dev/null; kill -KILL "$p" 2>/dev/null; done
  lsof -ti :8281 -ti :9199 -ti :8090 -ti :5180 -ti :4100 -ti :4490 -ti :4590 2>/dev/null | xargs kill -KILL 2>/dev/null
  rmdir firebase-export-* 2>/dev/null
  wait 2>/dev/null
  rm -rf "$logs"
  return 0
}
trap cleanup EXIT
trap 'cleanup; trap - EXIT; exit 130' INT TERM

# wait_port NAME PORT SECONDS waits for a TCP port, and fails the run
# with the log of the process when it never opens.
wait_port() {
  local name="$1" port="$2" limit="$3"
  for i in $(seq 1 "$limit"); do
    if nc -z 127.0.0.1 "$port" 2>/dev/null; then return 0; fi
    if [ "$i" -eq "$limit" ]; then
      echo "!! $name did not open :$port within $limit s" >&2
      cat "$logs/$name.log" >&2
      exit 1
    fi
    sleep 1
  done
}

echo "==> firestore + auth emulators (:8281, :9199), empty"
firebase emulators:start --only firestore,auth --project "$PROJECT_ID" </dev/null >"$logs/emulators.log" 2>&1 & pids+=($!)
wait_port emulators 8281 60
wait_port emulators 9199 60

echo "==> api on :8090 over the trimmed snapshot, every role on the fake"
(cd go && go run ./cmd/api) </dev/null >"$logs/api.log" 2>&1 & pids+=($!)
# /readyz answers 200 once the card index is loaded. The first go run
# compiles the API, so the wait covers a cold build cache.
for i in $(seq 1 180); do
  if curl -sf http://127.0.0.1:8090/readyz >/dev/null 2>&1; then break; fi
  if [ "$i" -eq 180 ]; then
    echo "!! the api was not ready within 180 s" >&2
    cat "$logs/api.log" >&2
    exit 1
  fi
  sleep 1
done

# The dev server binds 127.0.0.1 on purpose: plain localhost resolves
# to ::1 on macOS, and the port wait and the flow both name 127.0.0.1.
echo "==> web on :5180"
(cd web && pnpm --filter @mtg/web exec vite --host 127.0.0.1 --strictPort) </dev/null >"$logs/web.log" 2>&1 & pids+=($!)
wait_port web 5180 60

echo "==> playwright"
if (cd web/apps/web && pnpm exec playwright test "$@"); then
  exit 0
fi
echo "==> the flow failed. The api log, last 80 lines:" >&2
tail -80 "$logs/api.log" >&2
exit 1
