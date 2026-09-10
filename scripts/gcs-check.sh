#!/usr/bin/env bash
# gcs-check runs TestLiveFakeGCS against a fake-gcs-server seeded from the
# trimmed card snapshot of the repo (D-521). The CI step "fake gcs tests"
# and `make gcs-check` both run this script, so the two can not drift.
#
# The script fails unless the test passes by name. A filter that matches no
# test, and a test that skips, both exit 0 from `go test` (D-658).
set -euo pipefail

root="$(cd "$(dirname "$0")/.." && pwd)"
# Same pin as scripts/dev.sh and compose.yaml.
version="v1.56.1"
# 4453 keeps clear of the fake GCS of `make dev` on 4443.
port="${GCS_CHECK_PORT:-4453}"
snapshot="$root/go/cmd/deck-gate/testdata/snapshot/scryfall"

work="$(mktemp -d)"
server=""
cleanup() {
	if [ -n "$server" ]; then
		kill "$server" 2>/dev/null || true
	fi
	rm -rf "$work"
}
trap cleanup EXIT

# The bucket is the first folder under -data, and every file under it is an
# object named by its path.
mkdir -p "$work/seed/mtg-local-cards"
cp -R "$snapshot" "$work/seed/mtg-local-cards/"

# The binary runs directly, so the trap stops the server and not a go run
# wrapper that leaves it on the port.
GOBIN="$work/bin" go install "github.com/fsouza/fake-gcs-server@$version"
"$work/bin/fake-gcs-server" -scheme http -host 127.0.0.1 -port "$port" \
	-backend memory -data "$work/seed" >"$work/server.log" 2>&1 &
server=$!

for _ in $(seq 1 30); do
	if curl -sf "http://127.0.0.1:$port/storage/v1/b" >/dev/null; then
		break
	fi
	sleep 1
done
if ! curl -sf "http://127.0.0.1:$port/storage/v1/b" >/dev/null; then
	echo "gcs-check: fake-gcs-server did not answer on port $port"
	cat "$work/server.log"
	exit 1
fi

STORAGE_EMULATOR_HOST="http://127.0.0.1:$port" FAKE_GCS_LIVE=1 \
	go -C "$root/go" test -race -count=1 -v -run '^TestLiveFakeGCS$' ./internal/cards/ | tee "$work/test.log"

if ! grep -q -- '--- PASS: TestLiveFakeGCS' "$work/test.log"; then
	echo "gcs-check: TestLiveFakeGCS did not pass. A skip or a filter that matches no test reads as a failure here."
	exit 1
fi
echo "gcs-check: TestLiveFakeGCS passed against the seeded fake GCS."
