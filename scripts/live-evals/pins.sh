# shellcheck shell=bash
# The pinned version of each program that a tick runs (D-1165). Each
# line holds the name, the version command, and the first line of its
# output. The macOS build pins the programs of /bin and /usr/bin, and
# the Command Line Tools pin git and python3 apart. Claude Code has its
# own pin (D-1145), and .nvmrc pins Node 22.
#
# scripts/live-evals.sh reads each pin at each tick. The install of
# scripts/live-evals-launchd.sh picks the folder of each program by its
# pin (D-1175). Move a pin in a pull request after the check of the new
# version.
# shellcheck disable=SC2034 # the two scripts that source this file read PINS.
PINS=(
  "macOS|sw_vers -buildVersion|25F84"
  "jq|jq --version|jq-1.7.1-apple"
  "git|git --version|git version 2.50.1 (Apple Git-155)"
  "python3|python3 --version|Python 3.9.6"
  "go|go version|go version go1.27.1 darwin/arm64"
  "gh|gh --version|gh version 2.102.0 (2026-09-30)"
  "codex|codex --version|codex-cli 0.39.0"
  "gcloud|gcloud version|Google Cloud SDK 533.0.0"
  "node|node --version|v20.17.0"
  "pnpm|pnpm --version|9.2.0"
)
