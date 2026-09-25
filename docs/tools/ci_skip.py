#!/usr/bin/env python3
"""Decide whether the heavy jobs of `verify` skip a change of documents alone (D-818 to D-820).

  ci_skip.py --repo-slug OWNER/NAME --base SHA --head SHA --action ACTION
             [--before SHA] [--output FILE] [--repo DIR]

The heavy jobs skip when one of two rules holds:

  Rule 1: every path of the pull request is a document, and the newest
          `verify` run on the head of the pull request that last changed
          code on the base passed.
  Rule 2: every path of the push is a document, and the newest `verify`
          run on the previous head passed.

A fact that the script can not read runs every job, and never skips one.
The documentation set is the set of the `review-override` label (D-814),
less each document that a Go test reads (REV-065).
The script writes `skip=true` or `skip=false` to the output file.
"""
import argparse
import fnmatch
import json
import os
import subprocess
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from review_gate import eligible  # noqa: E402

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", ".."))
WORKFLOW = "verify.yml"
# The search for the last change of code on the base stops here, and the
# jobs then run.
MAX_COMMITS = 300

# A Go test reads these documents, so a change to one runs the Go job
# (REV-065). A pattern takes fnmatch wildcards. test_ci_skip.py finds
# each read in the Go tests and refuses one that this list misses.
TEST_READS = (
    ".claude/skills/mtg-corpus/SKILL.md",
    "docs/reference/feedback-fixer-prompt.md",
    "docs/reference/pr7-question-gate-run1*.md",
)


def document(path):
    return eligible(path) and not any(fnmatch.fnmatchcase(path, p) for p in TEST_READS)


def passed(run):
    return bool(run) and run.get("status") == "completed" and run.get("conclusion") == "success"


def describe(run):
    if not run:
        return "no run"
    return f"run {run.get('id')} with the status '{run.get('status')}' and the conclusion '{run.get('conclusion')}'"


def decide(pr_paths, push_paths, code_pr, code_run, previous_run):
    """Give (skip, reason).

    pr_paths and push_paths are lists of paths, or None when the fact is
    absent. code_pr is the number of the pull request that last changed
    code on the base, or None. code_run and previous_run are the newest
    `verify` runs, as dicts, or None.
    """
    if pr_paths is None:
        return False, "the paths of the pull request are unknown."
    code = [p for p in pr_paths if not document(p)]
    if not code:
        if code_pr is not None and passed(code_run):
            return True, f"rule 1: each of the {len(pr_paths)} path(s) is a document, and #{code_pr}, the last change of code on the base, passed verify ({describe(code_run)})."
        rule1 = f"rule 1 does not hold: #{code_pr} has {describe(code_run)}" if code_pr is not None else "rule 1 does not hold: no pull request of the last change of code on the base"
    else:
        rule1 = f"rule 1 does not hold: the pull request changes `{code[0]}`"
    if push_paths is None:
        return False, f"{rule1}, and the push has no previous head that is an ancestor of the head."
    pushed = [p for p in push_paths if not document(p)]
    if pushed:
        return False, f"{rule1}, and the push changes `{pushed[0]}`."
    if passed(previous_run):
        return True, f"rule 2: each of the {len(push_paths)} path(s) of the push is a document, and the previous head passed verify ({describe(previous_run)})."
    return False, f"{rule1}, and the previous head has {describe(previous_run)}."


def run(repo, *args):
    try:
        out = subprocess.run(list(args), cwd=repo, capture_output=True, text=True)
    except OSError:
        return None
    return out.stdout if out.returncode == 0 else None


def paths(repo, *diff):
    out = run(repo, "git", "diff", "--name-only", "--no-renames", *diff)
    return None if out is None else [p for p in out.splitlines() if p]


def last_code_commit(repo, base):
    out = run(repo, "git", "log", "--no-renames", "--name-only", "--format=@%H", f"-n{MAX_COMMITS}", base)
    if out is None:
        return None
    sha = None
    for line in out.splitlines():
        if line.startswith("@"):
            sha = line[1:]
        elif line and not document(line):
            return sha
    return None


def gh_json(repo, path, jq):
    out = run(repo, "gh", "api", path, "--jq", jq)
    if not out or out.strip() == "null":
        return None
    try:
        return json.loads(out)
    except json.JSONDecodeError:
        return None


def newest_run(repo, slug, sha):
    return gh_json(repo, f"repos/{slug}/actions/workflows/{WORKFLOW}/runs?head_sha={sha}&event=pull_request&per_page=100",
                   ".workflow_runs | sort_by(.created_at, .id) | last | if . == null then null else {id, status, conclusion} end")


def gather(repo, slug, base, head, action, before):
    merge_base = (run(repo, "git", "merge-base", base, head) or "").strip()
    pr_paths = paths(repo, merge_base, head) if merge_base else None
    code_pr = code_run = None
    if pr_paths is not None and all(document(p) for p in pr_paths):
        sha = last_code_commit(repo, merge_base)
        pr = gh_json(repo, f"repos/{slug}/commits/{sha}/pulls",
                     '[.[] | select(.merged_at != null)] | first | {number, sha: .head.sha}') if sha else None
        if pr:
            code_pr, code_run = pr["number"], newest_run(repo, slug, pr["sha"])
    push_paths = previous_run = None
    zero = "0" * 40
    if action == "synchronize" and before and before != zero:
        ancestor = subprocess.run(["git", "merge-base", "--is-ancestor", before, head], cwd=repo, capture_output=True)
        full = (run(repo, "git", "rev-parse", f"{before}^{{commit}}") or "").strip()
        if ancestor.returncode == 0 and full:
            push_paths = paths(repo, before, head)
            previous_run = newest_run(repo, slug, full)
    return pr_paths, push_paths, code_pr, code_run, previous_run


def main():
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument("--repo-slug", required=True)
    parser.add_argument("--base", required=True)
    parser.add_argument("--head", required=True)
    parser.add_argument("--action", required=True)
    parser.add_argument("--before", default="")
    parser.add_argument("--output")
    parser.add_argument("--repo", default=ROOT)
    args = parser.parse_args()
    skip, reason = decide(*gather(args.repo, args.repo_slug, args.base, args.head, args.action, args.before))
    print(f"ci-skip: {'skip' if skip else 'run'} the heavy jobs. {reason[0].upper()}{reason[1:]}")
    if args.output:
        with open(args.output, "a", encoding="utf-8") as handle:
            handle.write(f"skip={'true' if skip else 'false'}\n")
    return 0


if __name__ == "__main__":
    sys.exit(main())
