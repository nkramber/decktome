#!/usr/bin/env python3
"""Check the size of the files a session reads at start (D-749).

Claude Code loads CLAUDE.md into every call, and every session reads
docs/SESSION-HANDOFF.md first. Each byte of them sits in the context of
every later call. The audit of 2026-09-16 measured the hand-off at 5.5
percent of the carried context, and it grew from 35 KB to 59 KB in five
days. docs/reference/context-budget-2026-09-16.md holds the measurement.

The check fails when:
- a file passes its byte limit,
- the resume section of the hand-off passes its byte limit,
- the hand-off holds more session records than its limit,
- the three lists of paid targets differ, or a listed target does not exist,
- a .md file of .claude/skills passes its byte limit (D-753).

Run: python3 docs/tools/context_budget.py
"""
import os
import re
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

CLAUDE = "CLAUDE.md"
HANDOFF = "docs/SESSION-HANDOFF.md"
PAID = "docs/reference/paid-targets.md"
README = "README.md"
AGENTS = "AGENTS.md"

FILE_LIMITS = {CLAUDE: 11000, HANDOFF: 24000}

# A skill loads whole into the context of the session that reads it (D-753).
SKILLS = ".claude/skills"
SKILL_LIMIT = 36864
# D-750 parked the split of this skill until the measurement of five
# sessions. The exemption goes when the owner decides that rule.
SKILL_EXEMPT = ".claude/skills/mtg-corpus/"
RESUME_HEADING = "## RESUME HERE"
RESUME_LIMIT = 6000
SESSIONS_HEADING = "## The three most recent sessions"
SESSIONS_LIMIT = 3

# Each of these files names the paid targets in one sentence with this marker.
PAID_MARKER = "spend money:"
PAID_FILES = [CLAUDE, HANDOFF, PAID, README, AGENTS]
PAID_NAME = re.compile(r"`(make [a-z0-9-]+|scripts/[a-z0-9_-]+\.sh)`")


def section(text, heading):
    """Return the text from a '## ' heading that starts with `heading` to the next '## ' heading."""
    lines = text.splitlines(keepends=True)
    out, inside = [], False
    for line in lines:
        if line.startswith("## "):
            if inside:
                break
            inside = line.startswith(heading)
        if inside:
            out.append(line)
    return "".join(out) if out else None


def paid_names(text):
    """Return the paid target names of the first paragraph or list item that holds the marker."""
    for block in re.split(r"\n\s*\n|\n(?=- )", text):
        if PAID_MARKER in block:
            head = block.split(PAID_MARKER, 1)[1]
            head = re.split(r"(?<=[.])\s", head, maxsplit=1)[0]
            return set(PAID_NAME.findall(head))
    return None


def makefile_targets(text):
    return set(re.findall(r"^([a-z0-9-]+):", text, flags=re.M))


def check(read, exists, skill_files=()):
    """Return (report lines, errors). `read` maps a path to its text or None, and `exists` tests a path.

    `skill_files` names every `.md` file of the skills folder.
    """
    report, errors = [], []
    texts = {path: read(path) for path in {CLAUDE, HANDOFF, PAID, README, AGENTS, "Makefile"}}
    for path, limit in FILE_LIMITS.items():
        text = texts[path]
        if text is None:
            errors.append(f"{path} does not exist")
            continue
        size = len(text.encode("utf-8"))
        report.append(f"{path}: {size} of {limit} bytes")
        if size > limit:
            errors.append(f"{path} holds {size} bytes, over its limit of {limit}. Move history to the archive or a reference document")
    handoff = texts[HANDOFF] or ""
    resume = section(handoff, RESUME_HEADING)
    if resume is None:
        errors.append(f"{HANDOFF} holds no '{RESUME_HEADING}' section")
    else:
        size = len(resume.encode("utf-8"))
        report.append(f"{HANDOFF} resume section: {size} of {RESUME_LIMIT} bytes")
        if size > RESUME_LIMIT:
            errors.append(f"the resume section holds {size} bytes, over its limit of {RESUME_LIMIT}. Keep the state and the next step, and link the detail")
    sessions = section(handoff, SESSIONS_HEADING)
    if sessions is None:
        errors.append(f"{HANDOFF} holds no '{SESSIONS_HEADING}' section")
    else:
        count = len(re.findall(r"^### ", sessions, flags=re.M))
        report.append(f"{HANDOFF} session records: {count} of {SESSIONS_LIMIT}")
        if count > SESSIONS_LIMIT:
            errors.append(f"the hand-off holds {count} session records, over its limit of {SESSIONS_LIMIT}. Move the oldest to the archive, word for word")
    lists = {}
    for path in PAID_FILES:
        names = paid_names(texts[path] or "")
        if not names:
            errors.append(f"{path} holds no sentence with '{PAID_MARKER}' that names the paid targets")
            continue
        lists[path] = names
    if lists:
        first_path, first = next(iter(lists.items()))
        for path, names in lists.items():
            if names != first:
                missing = sorted(first - names)
                extra = sorted(names - first)
                errors.append(f"the paid targets of {path} differ from {first_path}: missing {missing}, extra {extra}")
        targets = makefile_targets(texts["Makefile"] or "")
        for name in sorted(set().union(*lists.values())):
            if name.startswith("make ") and name[5:] not in targets:
                errors.append(f"the paid target `{name}` is no Makefile target")
            if name.startswith("scripts/") and not exists(name):
                errors.append(f"the paid script `{name}` does not exist")
        report.append(f"paid targets: {len(first)} named in {len(lists)} files")
    over = []
    for path in sorted(skill_files):
        if path.startswith(SKILL_EXEMPT):
            continue
        size = len((read(path) or "").encode("utf-8"))
        if size > SKILL_LIMIT:
            over.append(path)
            errors.append(f"{path} holds {size} bytes, over its limit of {SKILL_LIMIT}. Move the detail to a file of the references folder")
    counted = [p for p in skill_files if not p.startswith(SKILL_EXEMPT)]
    report.append(f"skill files: {len(counted)} under {SKILL_LIMIT} bytes, {len(over)} over, {len(skill_files) - len(counted)} exempt")
    return report, errors


def main():
    def read(path):
        full = os.path.join(ROOT, path)
        if not os.path.exists(full):
            return None
        with open(full, encoding="utf-8") as handle:
            return handle.read()

    skills = []
    for base, _, files in os.walk(os.path.join(ROOT, SKILLS)):
        for name in files:
            if name.endswith(".md"):
                skills.append(os.path.relpath(os.path.join(base, name), ROOT))

    report, errors = check(read, lambda path: os.path.exists(os.path.join(ROOT, path)), skills)
    for line in report:
        print(f"context_budget: {line}")
    for error in errors:
        print(f"context_budget: {error}")
    return 1 if errors else 0


if __name__ == "__main__":
    sys.exit(main())
