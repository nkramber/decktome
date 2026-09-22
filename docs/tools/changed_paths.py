#!/usr/bin/env python3
"""Decide whether the code jobs of the verify workflow skip (D-805).

A pull request whose paths are all in the docs set skips the code jobs.
A push to a pull request skips them too when two things hold: the push
changes docs alone, and the verify:gate check passed on the previous
head. The gate job passes when a docs-only change skips the code jobs.
So a run of docs-only pushes reads back through each skipped head to the
last head that ran the code jobs, and that head passed them.

Two commands:

  decide  reads the facts of the change and writes documents-alone=true
          or documents-alone=false to the output file of the step.
  gate    reads the result of each code job and the decision, and exits 1
          when a code job failed, or skipped with no docs-only decision.

The docs set leaves out docs/reference/, because Go tests and the eval
check read files there. The verify:shell job runs on every head, so the
document checks and the tests of docs/tools and the session hook run on
every docs-only change.

Run:
  python3 docs/tools/changed_paths.py decide --event-name pull_request \\
      --pr-paths pr.txt [--push-paths push.txt --previous-checks checks.jsonl] \\
      --output "$GITHUB_OUTPUT"
  python3 docs/tools/changed_paths.py gate --needs "$NEEDS_JSON"
"""
import argparse
import json
import sys

DOCS_FOLDERS = ("docs/", ".claude/")
DOCS_EXCLUDED_FOLDERS = ("docs/reference/",)
DOCS_FILES = (
    "AGENTS.md",
    "CLAUDE.md",
    "LICENSE",
    "README.md",
    ".github/pull_request_template.md",
)
PULL_REQUEST_EVENT = "pull_request"
OUTPUT_NAME = "documents-alone"
ACTIONS_APP = "github-actions"
PATHS_JOB = "paths"
GATE_CHECK = "verify:gate"
CODE_JOBS = ("go", "vuln", "emulator", "web", "proto", "eval", "docker")


class FactError(Exception):
    """A fact of the change is absent or has the wrong form."""


def in_docs_set(path):
    """Return True when a path from the root of the checkout holds docs alone."""
    if not path:
        raise FactError("an empty path is no path")
    if path.startswith(DOCS_EXCLUDED_FOLDERS):
        return False
    return path.startswith(DOCS_FOLDERS) or path in DOCS_FILES


def first_path_outside(paths):
    """Return the first path outside the docs set, or None."""
    for path in paths:
        if not in_docs_set(path):
            return path
    return None


def read_check_runs(lines):
    """Read check runs from JSON lines: name, status, conclusion, app."""
    runs = []
    for number, line in enumerate(lines, start=1):
        if not line.strip():
            continue
        run = json.loads(line)
        if not isinstance(run, dict):
            raise FactError(f"check line {number} is not an object")
        for field in ("name", "status", "conclusion", "app"):
            if field not in run:
                raise FactError(f"check line {number} holds no field '{field}'")
        for field in ("name", "status", "app"):
            if not isinstance(run[field], str) or not run[field]:
                raise FactError(f"check line {number}: the field '{field}' is not text")
        if run["conclusion"] is not None and not isinstance(run["conclusion"], str):
            raise FactError(f"check line {number}: the field 'conclusion' is not text or null")
        runs.append(run)
    return runs


def gate_fault(runs):
    """Return why verify:gate did not pass on the previous head, or None.

    Each run of the name from GitHub Actions must complete with success.
    A rerun can leave an older run in the list, and an old failure must
    not hide under a new pass.
    """
    found = [r for r in runs if r["name"] == GATE_CHECK and r["app"] == ACTIONS_APP]
    if not found:
        return f"'{GATE_CHECK}' has no run"
    for run in found:
        if run["status"] != "completed":
            return f"'{GATE_CHECK}' has the status '{run['status']}'"
        if run["conclusion"] != "success":
            return f"'{GATE_CHECK}' ended with '{run['conclusion'] or 'no conclusion'}'"
    return None


def decide(event_name, pr_paths, push_paths, previous_checks):
    """Return (documents_alone, reason) for one run of the workflow."""
    if not event_name:
        raise FactError("the event name is empty")
    if event_name != PULL_REQUEST_EVENT:
        if pr_paths is not None or push_paths is not None or previous_checks is not None:
            raise FactError(f"the event '{event_name}' is not a pull request, and it takes no path file")
        return False, f"the event '{event_name}' runs every job it starts"
    if pr_paths is None:
        raise FactError("a pull request needs the file of --pr-paths")
    if (push_paths is None) != (previous_checks is None):
        raise FactError("--push-paths and --previous-checks go together")
    # An empty change skips nothing. An empty commit runs every job again.
    if not pr_paths:
        return False, "the pull request changes no path, so every job runs"
    code_path = first_path_outside(pr_paths)
    if code_path is None:
        return True, "the pull request changes docs alone, so the code jobs skip"
    if push_paths is None:
        return False, (f"the pull request changes '{code_path}', and no previous head is an ancestor "
                       "of this head, so every job runs")
    if not push_paths:
        return False, "the push changes no path, so every job runs"
    push_code_path = first_path_outside(push_paths)
    if push_code_path is not None:
        return False, f"the push changes '{push_code_path}', so every job runs"
    fault = gate_fault(previous_checks)
    if fault is not None:
        return False, f"the push changes docs alone, and on the previous head {fault}, so every job runs"
    return True, f"the push changes docs alone, and '{GATE_CHECK}' passed on the previous head, so the code jobs skip"


def gate(needs):
    """Return (passes, lines) for the verify:gate job.

    needs is the toJSON(needs) object of the workflow: one entry for each
    job, with its result and its outputs.
    """
    lines = []
    paths = needs.get(PATHS_JOB)
    if paths is None:
        return False, [f"the gate needs the job '{PATHS_JOB}'"]
    if paths.get("result") != "success":
        return False, [f"the job '{PATHS_JOB}' ended with '{paths.get('result')}'"]
    documents_alone = (paths.get("outputs") or {}).get(OUTPUT_NAME) == "true"
    passes = True
    for job in CODE_JOBS:
        entry = needs.get(job)
        if entry is None:
            passes = False
            lines.append(f"the gate needs the job '{job}'")
            continue
        result = entry.get("result")
        if result == "success" or (result == "skipped" and documents_alone):
            lines.append(f"'{job}' ended with '{result}'")
        else:
            passes = False
            lines.append(f"'{job}' ended with '{result}', and the gate fails")
    return passes, lines


def read_lines(path):
    with open(path, encoding="utf-8") as handle:
        return handle.read().splitlines()


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    sub = parser.add_subparsers(dest="mode", required=True)
    one = sub.add_parser("decide")
    one.add_argument("--event-name", required=True)
    one.add_argument("--pr-paths")
    one.add_argument("--push-paths")
    one.add_argument("--previous-checks")
    one.add_argument("--output", required=True)
    two = sub.add_parser("gate")
    two.add_argument("--needs", required=True)
    args = parser.parse_args(argv)
    try:
        if args.mode == "gate":
            passes, lines = gate(json.loads(args.needs))
            for line in lines:
                print(f"verify:gate: {line}")
            print(f"verify:gate: {'pass' if passes else 'fail'}")
            return 0 if passes else 1
        pr_paths = [p for p in read_lines(args.pr_paths) if p] if args.pr_paths else None
        push_paths = [p for p in read_lines(args.push_paths) if p] if args.push_paths else None
        checks = read_check_runs(read_lines(args.previous_checks)) if args.previous_checks else None
        alone, reason = decide(args.event_name, pr_paths, push_paths, checks)
    except (FactError, OSError, ValueError) as fault:
        print(f"changed_paths: stopped. {fault}", file=sys.stderr)
        return 2
    with open(args.output, "a", encoding="utf-8") as handle:
        handle.write(f"{OUTPUT_NAME}={'true' if alone else 'false'}\n")
    print(f"changed_paths: {reason}")
    print(f"changed_paths: {OUTPUT_NAME}={'true' if alone else 'false'}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
