#!/usr/bin/env bash
# Check the developer machine against the pinned versions.
# Prints one line per tool: ok, WRONG VERSION, or MISSING with a fix command.
# Exit 1 when any line is not ok. See docs/setup.md.
set -u
cd "$(dirname "$0")/.." || exit 1

status=0
want_go=$(awk '/^go /{print $2}' go/go.mod)
want_node=$(tr -d '[:space:]' < .nvmrc)
want_pnpm=$(sed -n 's/.*"packageManager": *"pnpm@\([^"]*\)".*/\1/p' web/package.json)
want_buf=$(awk '/github.com\/bufbuild\/buf /{print $2}' go/go.mod | sed 's/^v//')
# Same pin as docker/emulators.Dockerfile and docs/setup.md step 6.
want_firebase="14.14.0"

row() { printf '  %-14s %-14s %s\n' "$1" "$2" "$3"; }

check() {
  # check NAME CMD VERSION_CMD WANT FIX
  local name="$1" cmd="$2" vcmd="$3" want="$4" fix="$5" have
  if ! command -v "$cmd" >/dev/null 2>&1; then
    row "MISSING" "$name" "fix: $fix"
    status=1
    return
  fi
  version_row "$name" "$vcmd" "$want" "$fix"
}

# optional NAME CMD VERSION_CMD WANT FIX
# The same as check, but a missing tool is a warn line and not a failure.
# make dev needs none of these. See docs/setup.md.
optional() {
  local name="$1" cmd="$2" vcmd="$3" want="$4" fix="$5"
  if ! command -v "$cmd" >/dev/null 2>&1; then
    row "warn" "$name" "absent. Not needed for make dev. fix: $fix"
    return
  fi
  version_row "$name" "$vcmd" "$want" "$fix"
}

version_row() {
  local name="$1" vcmd="$2" want="$3" fix="$4" have
  have=$(eval "$vcmd" 2>&1 | head -1)
  if [ -n "$want" ] && ! printf '%s' "$have" | grep -q "$want"; then
    row "WRONG VERSION" "$name" "have $have, want $want. fix: $fix"
    status=1
    return
  fi
  row "ok" "$name" "$have"
}

echo "mtg-deck-builder doctor (macOS $(sw_vers -productVersion 2>/dev/null || echo ?), $(uname -m))"
optional "brew"  brew     "brew --version | awk '{print \$2}'"          ""           "https://brew.sh"
check "git"      git      "git --version | awk '{print \$3}'"           ""           "xcode-select --install"
# scripts/dev.sh waits on ports with nc and sweeps them with lsof.
check "nc"       nc       "echo present"                                ""           "brew install netcat (macOS ships nc)"
check "lsof"     lsof     "echo present"                                ""           "macOS ships lsof"
check "go"       go       "go version | awk '{print \$3}' | sed 's/^go//'" "$want_go" "https://go.dev/dl/ version $want_go"
check "node"     node     "node --version | sed 's/^v//'"               "$want_node" "nvm install (reads .nvmrc)"
check "pnpm"     pnpm     "pnpm --version"                              "$want_pnpm" "corepack enable && corepack prepare pnpm@$want_pnpm --activate"
check "firebase" firebase "firebase --version"                          "$want_firebase" "npm install -g firebase-tools@$want_firebase"
check "java"     java     "java -version 2>&1 | head -1 | sed 's/.*\"\\(.*\\)\".*/\\1/'" "17." "brew install openjdk@17 (docs/setup.md step 6)"
optional "docker" docker  "docker --version | awk '{print \$3}' | tr -d ," ""         "brew install --cask docker (docs/setup.md step 7, D-10). Only make dev-docker needs it."
optional "gcloud" gcloud  "gcloud --version | head -1 | awk '{print \$4}'" ""         "brew install --cask google-cloud-sdk. Only a cloud deploy needs it."
optional "python3" python3 "python3 --version | awk '{print \$2}'"      ""           "brew install python. docs/tools/ste-check.py and the README import example use it."
optional "shellcheck" shellcheck "shellcheck --version | sed -n 's/^version: //p'" "" "brew install shellcheck. CI lints scripts/*.sh with it."

if [ -x .bin/buf ]; then
  check "buf" .bin/buf ".bin/buf --version" "$want_buf" "make buf"
else
  row "MISSING" "buf" "fix: make buf (builds .bin/buf from go/go.mod)"
  status=1
fi

if [ -d web/node_modules ]; then
  row "ok" "web deps" "web/node_modules present"
else
  row "MISSING" "web deps" "fix: cd web && pnpm install"
  status=1
fi

# The daemon is only needed for make dev-docker. A stopped daemon is a warning.
if command -v docker >/dev/null 2>&1; then
  if docker info >/dev/null 2>&1; then
    row "ok" "docker daemon" "running"
  else
    row "warn" "docker daemon" "stopped. Only make dev-docker needs it. Open Docker Desktop to start it."
  fi
fi

if [ "$status" -eq 0 ]; then
  echo "all ok"
else
  echo "some checks failed. See docs/setup.md."
fi
exit $status
