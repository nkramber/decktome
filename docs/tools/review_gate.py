#!/usr/bin/env python3
"""The review gate of every pull request (D-810 to D-817).

  review_gate.py --event FILE --head REF [--repo DIR]
  review_gate.py --effective-head NUMBER [--base REF] [--head REF]
      Print the effective head. A reviewer records this commit (D-813).

A pull request passes in one of three ways:

  - A Codex review record at `docs/reviews/pr-<N>.md` on the head
    approves the effective head (D-811).
  - The `review-override` label is on, and every changed path is in the
    documentation set (D-812, D-814).
  - Dependabot opened it, and Dependabot wrote every commit (D-817).

The workflow runs this file from `main` on `pull_request_target`. It
reads the head of the pull request as data alone: it runs git on the
head ref, and it never runs a file of the head (D-816).

Exit 0 when the gate passes, 1 on a fault, and 2 on a usage error.
"""
import argparse
import json
import os
import re
import subprocess
import sys

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", ".."))
LABEL = "review-override"
APPROVED = "Ready for owner merge"
VERDICTS = (APPROVED, "Changes required", "Blocked")
SHORTEST_HASH = 7
HANDOFF_FILES = ("docs/SESSION-HANDOFF.md", "docs/reference/session-handoff-archive.md")
DEPENDABOT = "dependabot[bot]"
DEPENDABOT_EMAIL = "49699333+dependabot[bot]@users.noreply.github.com"

# The documentation set of the label (D-814). A folder ends in "/". Each
# refused path comes first, because it sits inside an eligible folder: a
# tool, a hook, and the harness settings run as code.
ELIGIBLE = ("docs/", ".claude/", "CLAUDE.md", "AGENTS.md", "README.md", ".github/pull_request_template.md")
REFUSED = ("docs/tools/", ".claude/hooks/", ".claude/settings.json", ".claude/settings.local.json")

# git lists no path for a clean merge. A merge brings new code into the
# branch, so gather gives it this path, which no metadata set holds.
MERGE = "(merge commit)"

HEAD_FIELD = re.compile(r"^\s*-\s*Head:\s*`([0-9a-fA-F]+)`")
BOLD = re.compile(r"\*\*([^*]+?)\*\*")


def record_path(number):
    return f"docs/reviews/pr-{number}.md"


def metadata_paths(number):
    """A commit that changes these paths alone keeps the effective head (D-813)."""
    return {record_path(number), f"docs/reviews/pr-{number}-response.md", *HANDOFF_FILES}


def matches(path, entries):
    return any(path == e or (e.endswith("/") and path.startswith(e)) for e in entries)


def eligible(path):
    return not matches(path, REFUSED) and matches(path, ELIGIBLE)


def effective_head(commits, number):
    """The newest commit that changes a path outside the metadata set, or None.

    commits holds (sha, files) pairs, oldest first. gather gives a merge
    commit the path MERGE, so a merge always moves the effective head.
    """
    metadata = metadata_paths(number)
    for sha, files in reversed(commits):
        if any(f not in metadata for f in files):
            return sha
    return None


def section(text, heading):
    """The lines under the exact heading, up to the next "## " heading, or None."""
    lines, inside = [], False
    for raw in text.split("\n"):
        line = raw.rstrip("\r")
        if inside and line.startswith("## "):
            break
        if inside:
            lines.append(line)
        elif line.rstrip() == heading:
            inside = True
    return lines if inside else None


def check_verdict(path, text):
    lines = section(text, "## Verdict")
    if lines is None:
        return "FAULT", f"`{path}` holds no `## Verdict` section."
    # One bold span is the verdict. A search of the whole section passes
    # "**Not Ready for owner merge.**", and a read of the first span alone
    # passes a record that holds a second verdict after it.
    bold = [m.group(1).strip().rstrip(".").strip() for line in lines for m in BOLD.finditer(line)]
    if not bold:
        return "FAULT", f"the `## Verdict` section of `{path}` holds no verdict name in bold."
    if len(bold) > 1:
        return "FAULT", f"the `## Verdict` section of `{path}` gives {len(bold)} bold spans: {', '.join(bold)}. It gives one, and an earlier verdict goes under `## Earlier verdicts`."
    if bold[0] not in VERDICTS:
        return "FAULT", f"the verdict of `{path}` is `{bold[0]}`, which is no verdict name of the `pr-review` skill."
    if bold[0] != APPROVED:
        return "FAULT", f"the verdict of `{path}` is `{bold[0]}`, and the gate needs `{APPROVED}`."
    return "PASS", f"the verdict of `{path}` is `{APPROVED}`."


def check_head(path, text, head):
    if head is None:
        return "FAULT", "every commit changes the metadata set alone, so the pull request has no effective head."
    lines = section(text, "## Identity")
    if lines is None:
        return "FAULT", f"`{path}` holds no `## Identity` section."
    recorded = next((m.group(1) for m in map(HEAD_FIELD.match, lines) if m), None)
    if recorded is None:
        return "FAULT", f"the `## Identity` list of `{path}` holds no line `- Head: ` with the hash in backticks."
    if len(recorded) < SHORTEST_HASH:
        return "FAULT", f"the head field of `{path}` is `{recorded}`, shorter than {SHORTEST_HASH} characters."
    if not head.lower().startswith(recorded.lower()):
        return "FAULT", f"the head field of `{path}` is `{recorded}`, and the effective head is `{head}`. Review the new diff, then update the head and the verdict together."
    return "PASS", f"the head field of `{path}` names the effective head `{head}`."


def evaluate(number, author, labels, files, commits, authors, read_record):
    """Give the result of each rule as (rule, state, message).

    files is the list of changed paths. commits holds (sha, files) pairs,
    oldest first. authors is the set of commit author emails. read_record
    gives the text of a path at the head, or None.
    """
    results = []
    if LABEL in labels:
        outside = [f for f in files if not eligible(f)]
        if outside:
            more = f" and {len(outside) - 1} other path(s)" if len(outside) > 1 else ""
            results.append(("RG 1", "FAULT", f"the `{LABEL}` label is on, and the pull request changes `{outside[0]}`{more}. That path takes the Codex review (D-814). Remove the label."))
        else:
            results.append(("RG 1", "PASS", f"the `{LABEL}` label is on, and every changed path is in the documentation set (D-812, D-814)."))
            return results
    else:
        results.append(("RG 1", "SKIP", f"the pull request carries no `{LABEL}` label."))

    if author == DEPENDABOT and authors == {DEPENDABOT_EMAIL}:
        results.append(("RG 2", "PASS", "Dependabot opened the pull request and wrote every commit (D-817)."))
        return results
    if author == DEPENDABOT:
        results.append(("RG 2", "SKIP", "Dependabot opened the pull request, and another author wrote a commit, so the Codex review applies (D-817)."))

    path = record_path(number)
    text = read_record(path)
    if text is None:
        results.append(("RG 3", "FAULT", f"the head holds no review record at `{path}`. A Codex session writes it with the `pr-review` skill (D-811)."))
        return results
    results.append(("RG 3", "PASS", f"the head holds the review record at `{path}`."))
    results.append(("RG 4", *check_verdict(path, text)))
    results.append(("RG 5", *check_head(path, text, effective_head(commits, number))))
    return results


def git(repo, *args, check=True):
    out = subprocess.run(["git", *args], cwd=repo, capture_output=True, text=True)
    if check and out.returncode != 0:
        raise SystemExit(f"review_gate: git {' '.join(args)} failed: {out.stderr.strip()}")
    return out


def gather(repo, base, head):
    """Read the facts of the head through git alone."""
    merge_base = git(repo, "merge-base", base, head).stdout.strip()
    # --no-renames: a code file that moves into docs/ still counts as a change of code.
    files = [p for p in git(repo, "diff", "--name-only", "--no-renames", merge_base, head).stdout.splitlines() if p]
    commits, authors = [], set()
    for line in git(repo, "log", "--reverse", "--format=%H %P%x09%ae", f"{merge_base}..{head}").stdout.splitlines():
        ids, email = line.split("\t", 1)
        sha, *parents = ids.split()
        changed = git(repo, "show", "--no-renames", "--name-only", "--format=", sha).stdout.splitlines()
        files = [p for p in changed if p]
        if len(parents) > 1:
            files.append(MERGE)
        commits.append((sha, files))
        authors.add(email)

    def read_record(path):
        out = git(repo, "show", f"{head}:{path}", check=False)
        return out.stdout if out.returncode == 0 else None

    return files, commits, authors, read_record


def main():
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument("--event", help="the pull_request_target event file")
    parser.add_argument("--head", default="HEAD", help="the ref that holds the head of the pull request")
    parser.add_argument("--repo", default=ROOT)
    parser.add_argument("--effective-head", type=int, metavar="NUMBER", help="print the effective head of this pull request")
    parser.add_argument("--base", default="origin/main", help="the base ref of --effective-head")
    args = parser.parse_args()
    if args.effective_head is not None:
        _, commits, _, _ = gather(args.repo, args.base, args.head)
        head = effective_head(commits, args.effective_head)
        print(head or "none: every commit changes the metadata set alone")
        return 0 if head else 1
    if not args.event:
        print("review_gate: pass --event, or --effective-head.", file=sys.stderr)
        return 2
    with open(args.event, encoding="utf-8") as handle:
        pr = json.load(handle).get("pull_request")
    if not pr:
        print("review_gate: the event holds no pull request.", file=sys.stderr)
        return 2
    number = pr["number"]
    author = (pr.get("user") or {}).get("login", "")
    labels = {label["name"] for label in pr.get("labels") or []}
    files, commits, authors, read_record = gather(args.repo, pr["base"]["sha"], args.head)
    results = evaluate(number, author, labels, files, commits, authors, read_record)
    print(f"review-gate: PR #{number}, {len(files)} changed path(s), {len(commits)} commit(s).")
    for rule, state, message in results:
        print(f"{rule}: {state} - {message}")
    faults = sum(1 for _, state, _ in results if state == "FAULT")
    print("review-gate: pass" if not faults else f"review-gate: {faults} fault(s)")
    return 1 if faults else 0


if __name__ == "__main__":
    sys.exit(main())
