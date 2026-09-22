#!/usr/bin/env python3
"""The review gate of every pull request (D-803, D-804).

The provider that wrote a pull request does not review it. The other
provider writes the review record docs/reviews/pr-<number>.md, and this
gate reads that record from the head of the pull request, as data alone.

The rules:

  RG 1  A pull request with the review-override label changes only paths
        of the eligible set.
  RG 2  A pull request with the label adds or removes no decision row.
  RG 3  The head holds docs/reviews/pr-<number>.md.
  RG 4  The Verdict section of the record gives one bold verdict, and it
        is "Ready for owner merge".
  RG 5  The Head field of the Identity section names the effective head.

When the label passes RG 1 and RG 2, RG 3 to RG 5 skip. The effective head
is the newest commit that changes a path outside the metadata set of the
pull request. A commit of the record, the response, or the two hand-off
files alone never makes a review stale.

The mode file .github/review-gate-mode comes from main. It reads
"enforced" or "advisory". In advisory mode the gate reports each fault and
exits 0. An absent or unknown mode is a fault in both modes.

Run:
  python3 docs/tools/review_gate.py --facts pr.json --head-files DIR --mode-file .github/review-gate-mode
"""
import argparse
import json
import os
import re
import sys

LABEL = "review-override"
DECISIONS = "docs/decisions.md"
MODES = ("enforced", "advisory")
APPROVED = "Ready for owner merge"
VERDICTS = (APPROVED, "Changes required", "Blocked")
SHORTEST_HASH = 7
ELIGIBLE_FOLDERS = ("docs/", ".claude/")
ELIGIBLE_FILES = ("CLAUDE.md", "AGENTS.md", "README.md", ".github/pull_request_template.md")
# Each of these runs as code: a CI tool, a hook of the harness, or the
# settings that start a hook. So the label never covers one of them.
INELIGIBLE = ("docs/tools/", ".claude/hooks/", ".claude/settings.json", ".claude/settings.local.json")
HEAD_FIELD = re.compile(r"^\s*-\s*Head:\s*`([0-9a-fA-F]+)`")
BOLD = re.compile(r"\*\*([^*]+?)\*\*")
DECISION_ROW = re.compile(r"^[+-]\|\s*D-\d+")

PASS, FAULT, SKIP = "pass", "fault", "skip"


class FactError(Exception):
    """A fact of the pull request is absent or has the wrong form."""


def record_path(number):
    return f"docs/reviews/pr-{number}.md"


def metadata_paths(number):
    """The paths whose change never moves the effective head (D-752, D-803)."""
    return {
        record_path(number),
        f"docs/reviews/pr-{number}-response.md",
        "docs/SESSION-HANDOFF.md",
        "docs/reference/session-handoff-archive.md",
    }


def effective_head(commits, number):
    """Return the newest commit that changes a path outside the metadata set, or None."""
    metadata = metadata_paths(number)
    for commit in reversed(commits):
        if any(path not in metadata for path in commit["files"]):
            return commit
    return None


def eligible(path):
    if path.startswith(INELIGIBLE):
        return False
    return path.startswith(ELIGIBLE_FOLDERS) or path in ELIGIBLE_FILES


def section_lines(text, heading):
    """Return the lines under an exact level-2 heading, or None."""
    lines = text.splitlines()
    for index, line in enumerate(lines):
        if line.strip() == heading:
            body = []
            for rest in lines[index + 1:]:
                if rest.startswith("## ") or rest.startswith("# "):
                    break
                body.append(rest)
            return body
    return None


def check_paths(facts):
    if LABEL not in facts["labels"]:
        return ("RG 1", PASS, f"the pull request carries no `{LABEL}` label")
    outside = [p for p in facts["files"] if not eligible(p)]
    if outside:
        return ("RG 1", FAULT, f"the `{LABEL}` label covers docs alone, and the pull request changes "
                               f"{', '.join(outside)}")
    return ("RG 1", PASS, "each changed path is in the eligible set of the label")


def check_decision_rows(facts):
    if LABEL not in facts["labels"]:
        return ("RG 2", PASS, f"the pull request carries no `{LABEL}` label")
    rows = [line for line in facts["decisionsDiff"].splitlines() if DECISION_ROW.match(line)]
    if rows:
        return ("RG 2", FAULT, f"the `{LABEL}` label covers no decision, and the pull request changes "
                               f"{len(rows)} decision row(s) of `{DECISIONS}`")
    return ("RG 2", PASS, f"the pull request changes no decision row of `{DECISIONS}`")


def check_verdict(path, text):
    lines = section_lines(text, "## Verdict")
    if lines is None:
        return ("RG 4", FAULT, f"`{path}` holds no `## Verdict` section")
    # One bold span only. A search of the whole section passes a record
    # that holds a bold negation or a second verdict beside the first one.
    bold = [m.group(1).strip().rstrip(".").strip() for line in lines for m in BOLD.finditer(line)]
    if not bold:
        return ("RG 4", FAULT, f"the `## Verdict` section of `{path}` holds no verdict name in bold")
    if len(bold) > 1:
        return ("RG 4", FAULT, f"the `## Verdict` section of `{path}` gives {len(bold)} bold names: "
                               f"{', '.join(bold)}. Put an earlier verdict under `## Earlier verdicts`")
    if bold[0] not in VERDICTS:
        return ("RG 4", FAULT, f"the verdict of `{path}` is `{bold[0]}`, which is no verdict of the "
                               "`pr-review` skill")
    if bold[0] != APPROVED:
        return ("RG 4", FAULT, f"the verdict of `{path}` is `{bold[0]}`, and the gate needs `{APPROVED}`")
    return ("RG 4", PASS, f"the verdict of `{path}` is `{APPROVED}`")


def check_head(path, text, head):
    if head is None:
        return ("RG 5", FAULT, "every commit changes the metadata set alone, so the pull request "
                               "has no effective head (D-747)")
    lines = section_lines(text, "## Identity")
    if lines is None:
        return ("RG 5", FAULT, f"`{path}` holds no `## Identity` section")
    recorded = next((m.group(1) for m in map(HEAD_FIELD.match, lines) if m), None)
    if recorded is None:
        return ("RG 5", FAULT, f"the `## Identity` section of `{path}` holds no line `- Head: ` "
                               "with the hash in backticks")
    if len(recorded) < SHORTEST_HASH:
        return ("RG 5", FAULT, f"the head field of `{path}` is `{recorded}`, shorter than "
                               f"{SHORTEST_HASH} letters")
    if not head["sha"].lower().startswith(recorded.lower()):
        return ("RG 5", FAULT, f"the head field of `{path}` is `{recorded}`, and the effective head is "
                               f"`{head['sha']}`")
    return ("RG 5", PASS, f"the head field of `{path}` names the effective head `{head['sha']}`")


def check(facts, read_head_file):
    """Return the result of each rule, in rule order."""
    paths = check_paths(facts)
    rows = check_decision_rows(facts)
    results = [paths, rows]
    if LABEL in facts["labels"] and paths[1] == PASS and rows[1] == PASS:
        reason = f"the `{LABEL}` label covers this pull request, so it needs no review record"
        return results + [("RG 3", SKIP, reason), ("RG 4", SKIP, reason), ("RG 5", SKIP, reason)]
    path = record_path(facts["number"])
    text = read_head_file(path)
    if text is None:
        return results + [("RG 3", FAULT, f"the head holds no review record at `{path}` (D-803)"),
                          ("RG 4", SKIP, "RG 3 found no review record"),
                          ("RG 5", SKIP, "RG 3 found no review record")]
    head = effective_head(facts["commits"], facts["number"])
    return results + [("RG 3", PASS, f"the head holds the review record at `{path}`"),
                      check_verdict(path, text), check_head(path, text, head)]


def read_facts(text):
    facts = json.loads(text)
    if not isinstance(facts, dict):
        raise FactError("the facts file is not an object")
    for field, kind in (("number", int), ("labels", list), ("files", list), ("commits", list),
                        ("decisionsDiff", str)):
        if field not in facts:
            raise FactError(f"the facts file holds no field '{field}'")
        if not isinstance(facts[field], kind) or (kind is int and isinstance(facts[field], bool)):
            raise FactError(f"the field '{field}' of the facts file is not a {kind.__name__}")
    for commit in facts["commits"]:
        if not isinstance(commit, dict) or not isinstance(commit.get("sha"), str) \
                or not isinstance(commit.get("files"), list):
            raise FactError("a commit of the facts file needs a sha and a list of files")
    return facts


def read_mode(text):
    if text is None:
        raise FactError("main holds no .github/review-gate-mode file")
    mode = text.strip()
    if mode not in MODES:
        raise FactError(f"the mode '{mode}' is not one of: {', '.join(MODES)}")
    return mode


def report(mode, facts, results):
    faults = sum(1 for _, result, _ in results if result == FAULT)
    lines = [f"review-gate: pull request #{facts['number']}, mode {mode}, "
             f"{len(facts['files'])} changed path(s), {len(facts['commits'])} commit(s)"]
    lines += [f"{rule} {result}: {reason}." for rule, result, reason in results]
    if faults == 0:
        lines.append("review-gate: pass")
        return lines, 0
    if mode == "advisory":
        lines.append(f"review-gate: {faults} fault(s), and the advisory mode passes the check")
        return lines, 0
    lines.append(f"review-gate: {faults} fault(s)")
    return lines, 1


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    parser.add_argument("--facts", required=True)
    parser.add_argument("--head-files", required=True)
    parser.add_argument("--mode-file", required=True)
    args = parser.parse_args(argv)

    def read(path):
        if not os.path.isfile(path):
            return None
        with open(path, encoding="utf-8") as handle:
            return handle.read()

    root = os.path.realpath(args.head_files)

    def read_head_file(path):
        full = os.path.realpath(os.path.join(root, path))
        # A symlink in the head must not point the gate at a file outside it.
        if not full.startswith(root + os.sep):
            return None
        return read(full)

    try:
        if not os.path.isdir(root):
            raise FactError(f"the folder '{args.head_files}' does not exist")
        mode = read_mode(read(args.mode_file))
        facts = read_facts(read(args.facts) or "")
        lines, code = report(mode, facts, check(facts, read_head_file))
    except (FactError, ValueError) as fault:
        print(f"review-gate: stopped. {fault}", file=sys.stderr)
        return 2
    print("\n".join(lines))
    return code


if __name__ == "__main__":
    sys.exit(main())
