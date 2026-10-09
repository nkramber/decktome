#!/usr/bin/env python3
"""Check that each build of Go names the version of go/go.mod (D-1216).

The go line of go/go.mod is the one source of the Go version. The verify
workflow reads it through setup-go. A local go command in go/ switches to
it by itself. Two copies can not read it, so this check holds them to it:

- the FROM golang line of each Dockerfile in docker/,
- the go pin of scripts/live-evals/pins.sh.

Run: python3 docs/tools/go_version_check.py
"""
import pathlib
import re
import sys

ROOT = pathlib.Path(__file__).resolve().parents[2]


def go_mod_version(root):
    m = re.search(r"^go (\d+\.\d+\.\d+)$", (root / "go/go.mod").read_text(), re.M)
    return m.group(1) if m else None


def faults(root):
    want = go_mod_version(root)
    if not want:
        return ["go/go.mod has no go line of the form 1.X.Y"]
    out = []
    files = sorted((root / "docker").glob("*.Dockerfile"))
    for f in files:
        for tag in re.findall(r"^FROM golang:([^@\s]+)", f.read_text(), re.M):
            if tag != want:
                out.append(f"{f.relative_to(root)}: golang:{tag}, and go/go.mod reads {want}")
    pins = (root / "scripts/live-evals/pins.sh").read_text()
    m = re.search(r'^\s*"go\|go version\|go version go([^ ]+) ', pins, re.M)
    if not m:
        out.append("scripts/live-evals/pins.sh: no go pin")
    elif m.group(1) != want:
        out.append(f"scripts/live-evals/pins.sh: go{m.group(1)}, and go/go.mod reads {want}")
    return out


def main():
    found = faults(ROOT)
    for f in found:
        print(f"go-version-check: {f}")
    print(f"go-version-check: {len(found)} finding(s)")
    return 1 if found else 0


if __name__ == "__main__":
    sys.exit(main())
