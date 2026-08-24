#!/usr/bin/env bash
# Check the developer machine. Print each tool, its version, and a fix command when absent.
set -u
ok=0
check() {
  local name="$1" cmd="$2" fix="$3"
  if command -v "${cmd%% *}" >/dev/null 2>&1; then
    printf '  ok      %-14s %s\n' "$name" "$($cmd 2>&1 | head -1)"
  else
    printf '  MISSING %-14s fix: %s\n' "$name" "$fix"
    ok=1
  fi
}
echo "mtg-deck-builder doctor"
check "go"       "go version"          "https://go.dev/dl/ (see go/go.mod for the version)"
check "node"     "node --version"      "nvm install (see .nvmrc)"
check "pnpm"     "pnpm --version"      "corepack enable && corepack prepare pnpm@9.2.0 --activate"
check "git"      "git --version"       "xcode-select --install"
check "firebase" "firebase --version"  "npm i -g firebase-tools (PR-0c)"
check "java"     "java -version"       "brew install openjdk@17 (Firestore emulator, PR-0c)"
check "docker"   "docker --version"    "https://docs.docker.com/desktop/ (PR-0b, D-10)"
check "gcloud"   "gcloud --version"    "https://cloud.google.com/sdk/docs/install"
echo "  (buf, protoc-gen-go, protoc-gen-connect-go come from go/go.mod tool directives: go -C go tool buf --version)"
exit $ok
