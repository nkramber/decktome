#!/usr/bin/env python3
"""Check that every id and every repository path in a document resolves (D-753).

The documents of this repo cite an id of a register and a path in backticks.
A rename or a split leaves a citation that points at nothing, and a reader
then trusts a dead reference. The checker applies two rules:

- REF 1: a cited D-, F-, M-, PR-, or I- id that no register defines.
- REF 2: a path of this repo in backticks that no file and no folder holds.

The registers: docs/decisions.md defines D- with a table row. The roadmap
docs/design-roadmap.md defines F- with a table row, and it defines M-, PR-,
and I- with a bold entry title. A bare family id resolves when a lettered
variant exists, for example PR-28 against PR-28a.

The checker does not read OQ- ids. This repo deletes an answered row from
docs/owner-questions.md, and the decision then holds the answer (D-753).

REF 2 reads a path with a slash and a first part that names a top-level
entry of the checkout. A bare file name is ambiguous, so the rule skips it.
The last part of the path holds a file type, or the whole path holds
lowercase letters alone. So a placeholder such as `go/X` takes no rule.
The first part can start with one dot, for `.claude` and `.github`.
A path resolves from the root, from the folder of the document, from the
folder above it, or as the one path of the checkout that ends with it.

A dated record is history, and a rewrite of it falsifies the record. So the
checker reads no rule in the hand-off archive and in a file that ends with
a date. Makefile keeps the same set out of the STE check.

Usage: python3 docs/tools/ref_check.py FILE [FILE ...]
"""
import os
import re
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

DECISIONS = "docs/decisions.md"
ROADMAP = "docs/design-roadmap.md"

# The fixer of the loop wrote these five numbers on a branch the owner
# dropped (D-204). They are retired, and docs/decisions.md names them.
RETIRED = {f"D-{n}" for n in range(172, 177)}

# A run writes this folder, and no commit holds it (D-642). A path under
# `.local` is the scratch of a fit, and `.gitignore` holds it.
RUNTIME = {"docs/reference/feedback"}
SCRATCH = "/.local/"

# A dated record is history. A rewrite of it falsifies the record.
DATED = re.compile(r"-\d{4}-\d{2}-\d{2}\.md$|session-handoff-archive\.md$")

ID = re.compile(r"\b((?:D|F|M|PR|I)-\d+[A-Za-z]?)\b")
ROW_ID = re.compile(r"^\| ((?:D|F)-\d+)", re.M)
ENTRY_ID = re.compile(r"\*\*((?:M|PR|I)-\d+[A-Za-z]?):")
CODE = re.compile(r"`([^`\n]+)`")
FENCE = re.compile(r"^```.*?^```", re.M | re.S)
PATH = re.compile(r"^\.?[A-Za-z0-9_][\w./-]*$")


def read(path):
    full = os.path.join(ROOT, path)
    if not os.path.exists(full):
        return ""
    with open(full, encoding="utf-8") as handle:
        return handle.read()


def registers(decisions, roadmap):
    """Return the set of ids that the two registers define."""
    known = set(ROW_ID.findall(decisions)) | set(ROW_ID.findall(roadmap))
    known |= set(ENTRY_ID.findall(roadmap))
    # A bare family id resolves when a lettered variant exists (PR-28a).
    known |= {i[:-1] for i in known if i[-1].isalpha()}
    return known


def repo_paths():
    """Return the paths of the checkout, and the folders above each one."""
    paths, folders = set(), set()
    for base, names, files in os.walk(ROOT):
        rel = os.path.relpath(base, ROOT)
        if rel == ".":
            rel = ""
        names[:] = [n for n in names if n not in {".git", "node_modules", "dist", "build"}]
        for name in names:
            folders.add(os.path.join(rel, name) if rel else name)
        for name in files:
            paths.add(os.path.join(rel, name) if rel else name)
    return paths, folders


def resolves(token, doc, paths, folders):
    here = os.path.dirname(doc)
    above = os.path.dirname(here)
    for base in (".", here, above):
        candidate = os.path.normpath(os.path.join(base, token))
        if candidate in paths or candidate in folders:
            return True
    tail = "/" + token
    return any(p.endswith(tail) for p in paths) or any(f.endswith(tail) for f in folders)


def check(doc, text, known, top, paths, folders):
    """Return the findings of one document as (line, rule, message)."""
    findings = []
    clean = FENCE.sub(lambda m: "\n" * m.group(0).count("\n"), text)
    for number, line in enumerate(clean.splitlines(), start=1):
        for match in ID.finditer(line):
            name = match.group(1)
            if name in known or name in RETIRED:
                continue
            findings.append((number, "REF 1", f"no register defines {name}"))
        for match in CODE.finditer(line):
            token = match.group(1).strip().rstrip("/")
            if "/" not in token or not PATH.match(token):
                continue
            if token.split("/")[0] not in top or token in RUNTIME:
                continue
            if SCRATCH in "/" + token:
                continue
            if "." not in token.rsplit("/", 1)[1] and token != token.lower():
                continue
            if resolves(token, doc, paths, folders):
                continue
            findings.append((number, "REF 2", f"no file and no folder holds `{token}`"))
    return findings


def main(argv):
    decisions, roadmap = read(DECISIONS), read(ROADMAP)
    known = registers(decisions, roadmap)
    if not known:
        print("ref_check: the registers hold no id. Run this from the checkout")
        return 1
    paths, folders = repo_paths()
    top = {name.split("/")[0] for name in paths | folders}
    total = 0
    for doc in argv:
        if DATED.search(doc):
            continue
        for number, rule, message in check(doc, read(doc), known, top, paths, folders):
            print(f"{doc}:{number}: rule {rule}: {message}")
            total += 1
    print(f"{total} finding(s)")
    return 1 if total else 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
