#!/usr/bin/env python3
"""Start the Codex review of one pull request, and read its record (D-823 to D-832).

  codex_review.py --pr NUMBER [--repo DIR]

`make codex-review PR=<n>` runs this file. The run has three parts:

  1. The refusals. The tool refuses to start when the pull request is
     not open, the checkout is not the head of the pull request, the tree
     is dirty, or the Gitar pass of the effective head is not complete.
     It then updates the npm CLI to the newest release, and it refuses a
     CLI below MIN_VERSION, a login that is not ChatGPT, or a model that
     fails the probe (D-824, D-833).
  2. The review. Codex runs the `pr-review` skill in a new git worktree
     at the head, with a detached HEAD, so the checkout of the author
     stays unchanged. The transcript goes to `.local/codex-review/`.
  3. The read. The tool fetches the branch, reads the record from
     origin, and checks its head field against the effective head.

Each outcome has its own exit code:

  0  approve: the record says `Ready for owner merge` for the effective head.
  1  fault: no record, a stale head, a Codex error, or a push of another path.
  2  usage error.
  3  changes: the record says `Changes required` or `Blocked`.
  4  three-strike stop: a blocking finding is open at its third head (D-826).
  5  refusal: a check of part 1 failed, and no review ran.

`make` exits 2 for each code that is not 0. The last line of the output
names the outcome and the code of this file.

Each Codex call runs with OPENAI_API_KEY and CODEX_API_KEY removed from
its environment, so a review bills the ChatGPT plan and never the API
(D-833).
"""
import argparse
import datetime
import json
import os
import re
import shutil
import subprocess
import sys
import tempfile

HERE = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, HERE)
import review_gate as rg  # noqa: E402

ROOT = os.path.abspath(os.path.join(HERE, "..", ".."))
NPM_PACKAGE = "@openai/codex"
MIN_VERSION = "0.156.1"
MODEL = "gpt-6-luna"
EFFORT = "medium"
SANDBOX = "danger-full-access"
STRIKES = 3
BLOCKING = ("P0", "P1", "P2")
TRANSCRIPTS = ".local/codex-review"
GITAR = "gitar-bot[bot]"
GITAR_SLUG = "gitar-bot"
REVIEW_TIMEOUT = 4 * 3600
# A Codex call bills the ChatGPT plan alone, and never the API (D-833).
API_KEYS = ("OPENAI_API_KEY", "CODEX_API_KEY")
CHATGPT_LOGIN = "Logged in using ChatGPT"

EXIT_APPROVE, EXIT_FAULT, EXIT_USAGE, EXIT_CHANGES, EXIT_STRIKE, EXIT_REFUSAL = 0, 1, 2, 3, 4, 5
OUTCOMES = {EXIT_APPROVE: "approve", EXIT_FAULT: "fault", EXIT_CHANGES: "changes",
            EXIT_STRIKE: "three-strike stop", EXIT_REFUSAL: "refusal"}

VERSION = re.compile(r"(\d+)\.(\d+)\.(\d+)(-\S+)?")
FINDING = re.compile(r"^###\s+(P([0-3])-\d+)\s*:")
STATUS = re.compile(r"^\s*Status:\s*(.+?)\s*$")
OPEN_AT = re.compile(r"^\s*Open at:\s*(.+?)\s*$")
HASH = re.compile(r"`([0-9a-fA-F]{7,40})`")
REQUEST = re.compile(r"^\s*gitar review\s*$", re.IGNORECASE)
REPLY = re.compile(r"^> gitar review", re.IGNORECASE)
DASHBOARD = "<b>Code Review</b>"


class Stop(Exception):
    """End the run with an exit code and a message."""

    def __init__(self, code, message):
        super().__init__(message)
        self.code = code


def refuse(message):
    return Stop(EXIT_REFUSAL, message)


def fault(message):
    return Stop(EXIT_FAULT, message)


def sh(cmd, cwd=None, timeout=None, stdout=None, stderr=None, env=None):
    """Run one command, and give (code, stdout, stderr)."""
    try:
        out = subprocess.run(cmd, cwd=cwd, timeout=timeout, text=True, env=env,
                             stdout=stdout or subprocess.PIPE, stderr=stderr or subprocess.PIPE)
    except FileNotFoundError:
        return 127, "", f"{cmd[0]}: command not found"
    except subprocess.TimeoutExpired:
        return 124, "", f"{cmd[0]}: no exit after {timeout} seconds"
    return out.returncode, out.stdout or "", out.stderr or ""


def must(run, cmd, stop, cwd=None, env=None):
    code, out, err = run(cmd, cwd=cwd, env=env)
    if code != 0:
        raise stop(f"`{' '.join(cmd[:4])}` exited {code}: {(err or out).strip()[:300]}")
    return out


def version(text):
    """The version of a `codex --version` line as a sortable tuple, or None.

    A pre-release sorts below its release: 0.157.0-alpha.1 < 0.157.0.
    """
    m = VERSION.search(text or "")
    if not m:
        return None
    return int(m.group(1)), int(m.group(2)), int(m.group(3)), 0 if m.group(4) else 1


# Part 1: the refusals.

def check_pr(run, number):
    out = must(run, ["gh", "pr", "view", str(number), "--json",
                     "state,headRefName,headRefOid,isCrossRepository"], refuse)
    pr = json.loads(out)
    if pr["state"] != "OPEN":
        raise refuse(f"pull request #{number} is {pr['state']}, not OPEN.")
    if pr["isCrossRepository"]:
        raise refuse(f"pull request #{number} comes from a fork, and the review can not push its record there.")
    return pr["headRefName"], pr["headRefOid"]


def check_checkout(run, repo, branch, head):
    must(run, ["git", "fetch", "--quiet", "origin", "main", branch], refuse, cwd=repo)
    local = must(run, ["git", "rev-parse", "HEAD"], refuse, cwd=repo).strip()
    remote = must(run, ["git", "rev-parse", f"origin/{branch}"], refuse, cwd=repo).strip()
    if remote != head:
        raise refuse(f"origin/{branch} is {remote[:7]}, and GitHub gives the head {head[:7]}. Fetch again.")
    if local != head:
        raise refuse(f"the checkout is at {local[:7]}, and the head of `{branch}` on origin is {head[:7]}. Push or pull first.")
    dirty = [line for line in must(run, ["git", "status", "--porcelain"], refuse, cwd=repo).splitlines() if line]
    if dirty:
        raise refuse(f"the tree has {len(dirty)} uncommitted path(s), the first `{dirty[0][3:]}`. Commit and push, or stash.")


def push_time(run, slug, shas):
    """The time GitHub first saw the push that holds shas[0].

    shas lists the effective head and each later commit, oldest first. A
    push of many commits gets check suites on its tip alone, so the first
    commit with a suite names the push.
    """
    for sha in shas:
        out = must(run, ["gh", "api", f"repos/{slug}/commits/{sha}/check-suites",
                         "--jq", "[.check_suites[].created_at] | min"], refuse)
        if out.strip() and out.strip() != "null":
            return out.strip()
    raise refuse(f"GitHub holds no check suite for {shas[0][:7]} or a later commit, so the push time is unknown.")


def gitar_problems(pushed, comments, gitar_runs, threads):
    """Each reason that the Gitar pass is not complete (the `gitar-review` skill).

    pushed is the push time of the effective head. comments are the issue
    comments, gitar_runs the Gitar check runs on the tip, and threads the
    review threads of the pull request.
    """
    problems = []
    running = [r for r in gitar_runs if r.get("status") != "completed"]
    if running:
        problems.append("the Gitar check on the tip is not complete.")
    dashboards = [c for c in comments if c["user"]["login"] == GITAR and DASHBOARD in (c.get("body") or "")]
    if not dashboards:
        problems.append("the pull request holds no Gitar dashboard comment.")
        dashboard = None
    else:
        dashboard = max(dashboards, key=lambda c: c["created_at"])["updated_at"]
        if dashboard <= pushed:
            problems.append(f"the Gitar dashboard changed at {dashboard}, before the push of the effective head at {pushed}. Ask for a review with the `gitar-review` skill.")
    requests = [c for c in comments if REQUEST.match(c.get("body") or "") and c["created_at"] > pushed]
    if requests and dashboard:
        asked = max(c["created_at"] for c in requests)
        replies = [c for c in comments if c["user"]["login"] == GITAR and REPLY.match(c.get("body") or "") and c["created_at"] > asked]
        if not replies:
            problems.append(f"Gitar gave no reply to the request of {asked}.")
        else:
            reply = max(replies, key=lambda c: c["created_at"])
            if "on it" not in reply["body"].lower():
                problems.append(f"Gitar refused the request of {asked}. Wait, then ask again.")
            elif dashboard <= reply["created_at"]:
                problems.append(f"the manual review of {asked} has not changed the dashboard yet.")
    open_threads = [t for t in threads if not t.get("isResolved")]
    if open_threads:
        where = ", ".join(f"{t.get('path')}:{t.get('line')}" for t in open_threads[:3])
        problems.append(f"{len(open_threads)} review thread(s) are not resolved: {where}.")
    return problems


THREADS = """query($owner: String!, $name: String!, $number: Int!, $endCursor: String) {
  repository(owner: $owner, name: $name) { pullRequest(number: $number) {
    reviewThreads(first: 100, after: $endCursor) {
      nodes { isResolved path line } pageInfo { hasNextPage endCursor } } } } }"""


def review_threads(run, slug, number):
    """Every review thread of the pull request. `gh --paginate` follows endCursor to the last page."""
    owner, name = slug.split("/", 1)
    pages = json.loads(must(run, ["gh", "api", "graphql", "--paginate", "--slurp", "-F", f"owner={owner}",
                                  "-F", f"name={name}", "-F", f"number={number}", "-f", f"query={THREADS}"], refuse))
    return [t for page in pages for t in page["data"]["repository"]["pullRequest"]["reviewThreads"]["nodes"]]


def later_commits(run, repo, effective, tip):
    """The effective head and each later commit of the branch, oldest first.

    The walk follows the first parent alone. A merge of `main` as the
    effective head otherwise brings in old commits of `main`, and their
    check suites give a push time long before the merge.
    """
    return must(run, ["git", "rev-list", "--reverse", "--first-parent", f"{effective}^..{tip}"], refuse, cwd=repo).split()


def check_gitar(run, repo, slug, number, effective, tip):
    shas = later_commits(run, repo, effective, tip)
    pushed = push_time(run, slug, shas or [effective])
    pages = json.loads(must(run, ["gh", "api", "--paginate", "--slurp", f"repos/{slug}/issues/{number}/comments"], refuse))
    comments = [c for page in pages for c in page]
    runs = json.loads(must(run, ["gh", "api", f"repos/{slug}/commits/{tip}/check-runs", "--jq",
                                 f"[.check_runs[] | select(.app.slug == \"{GITAR_SLUG}\") | {{status}}]"], refuse))
    threads = review_threads(run, slug, number)
    problems = gitar_problems(pushed, comments, runs, threads)
    if problems:
        raise refuse("the Gitar pass is not complete: " + " ".join(problems))


def update_cli(run):
    """Install the newest npm release, and give the path of its binary (D-824)."""
    must(run, ["npm", "install", "-g", f"{NPM_PACKAGE}@latest", "--no-fund", "--no-audit", "--loglevel=error"], refuse)
    latest = must(run, ["npm", "view", NPM_PACKAGE, "version"], refuse).strip()
    codex = os.path.join(must(run, ["npm", "prefix", "-g"], refuse).strip(), "bin", "codex")
    installed = must(run, [codex, "--version"], refuse, env=codex_env()).strip()
    have, want = version(installed), version(latest)
    if have is None or want is None:
        raise refuse(f"can not read the versions: installed `{installed}`, newest `{latest}`.")
    if have != want:
        raise refuse(f"{codex} reads `{installed}` after the update, and npm gives {latest} as the newest.")
    if have < version(MIN_VERSION):
        raise refuse(f"{codex} is `{installed}`, below the minimum {MIN_VERSION}.")
    return codex, installed


def codex_env():
    """The environment of each Codex call, with no API key in it (D-833)."""
    return {k: v for k, v in os.environ.items() if k not in API_KEYS}


def check_login(run, codex):
    """Refuse a Codex that would bill the API and not the ChatGPT plan (D-833)."""
    code, out, err = run([codex, "login", "status"], env=codex_env())
    text = (out + err).strip()
    if code != 0 or CHATGPT_LOGIN not in text:
        raise refuse(f"`codex login status` gives `{text[:120]}`, and the review needs `{CHATGPT_LOGIN}`. Run `codex login` with the ChatGPT account.")


def model_args():
    """The model, the effort, the approvals, and the sandbox of each call. None comes from the config."""
    return ["-m", MODEL, "-c", f'model_reasoning_effort="{EFFORT}"', "-c", 'approval_policy="never"']


def probe(run, codex):
    with tempfile.TemporaryDirectory() as scratch:
        answer = os.path.join(scratch, "answer.txt")
        code, out, err = run([codex, "exec", *model_args(), "-s", "read-only", "--ephemeral",
                              "--skip-git-repo-check", "-C", scratch, "-o", answer,
                              "Reply with the single word OK."], cwd=scratch, env=codex_env())
        text = ""
        if os.path.exists(answer):
            with open(answer, encoding="utf-8") as handle:
                text = handle.read().strip()
    if code != 0 or text.rstrip(".").upper() != "OK":
        last = (err or out).strip().splitlines()[-1:] or ["no output"]
        raise refuse(f"the probe of model {MODEL} at effort {EFFORT} failed, exit {code}: {last[0][:300]}")


# Part 2: the review.

def prompt(number, slug, branch, head):
    return "\n".join([
        f"Review pull request #{number} of {slug} with the `pr-review` skill.",
        "Load and follow `.claude/skills/pr-review/SKILL.md`. Your role is reviewer.",
        f"This directory is a git worktree at {head}, the head of branch `{branch}`, with a detached HEAD.",
        "Commit and push the record as `.claude/skills/pr-review/references/commit-and-end.md` says,",
        f"with `git push origin HEAD:{branch}`.",
    ])


def review(run, repo, codex, number, slug, branch, head, stamp):
    os.makedirs(os.path.join(repo, TRANSCRIPTS), exist_ok=True)
    base = os.path.join(repo, TRANSCRIPTS, f"pr-{number}-{stamp}")
    tree = tempfile.mkdtemp(prefix=f"decktome-codex-pr{number}-")
    must(run, ["git", "worktree", "add", "--quiet", "--detach", tree, head], fault, cwd=repo)
    # A new worktree holds no web packages, and `make verify` needs them for `buf generate`.
    try:
        must(run, ["pnpm", "--dir", "web", "install", "--frozen-lockfile", "--silent"], fault, cwd=tree)
    except Stop:
        run(["git", "worktree", "remove", "--force", tree], cwd=repo)
        shutil.rmtree(tree, ignore_errors=True)
        raise
    with open(base + ".jsonl", "w", encoding="utf-8") as events, open(base + ".stderr.log", "w", encoding="utf-8") as log:
        code, _, _ = run([codex, "exec", *model_args(), "-s", SANDBOX, "-C", tree, "--json",
                          "-o", base + ".last.md", prompt(number, slug, branch, head)],
                         cwd=tree, timeout=REVIEW_TIMEOUT, stdout=events, stderr=log, env=codex_env())
    return code, tree, base


def remove_tree(run, repo, tree, branch):
    """Remove the worktree when it holds nothing that origin lacks. Give the kept path, or None."""
    status = run(["git", "status", "--porcelain"], cwd=tree)
    unpushed = run(["git", "rev-list", "--count", f"origin/{branch}..HEAD"], cwd=tree)
    if status[0] == 0 and not status[1].strip() and unpushed[0] == 0 and unpushed[1].strip() == "0":
        run(["git", "worktree", "remove", "--force", tree], cwd=repo)
        shutil.rmtree(tree, ignore_errors=True)
        return None
    return tree


# Part 3: the read.

def verdict(text):
    lines = rg.section(text, "## Verdict")
    if lines is None:
        raise fault("the record holds no `## Verdict` section.")
    bold = [m.group(1).strip().rstrip(".").strip() for line in lines for m in rg.BOLD.finditer(line)]
    if len(bold) != 1 or bold[0] not in rg.VERDICTS:
        raise fault(f"the `## Verdict` section gives {bold or 'no bold span'}, and it must give one verdict name.")
    return bold[0]


def findings(text):
    """Each finding of the record as (id, severity, status, open_at), in record order.

    open_at holds the heads of the `Open at:` line, the heads at which a
    review found the finding open (D-826). An older record has no such line.
    """
    result, current = [], None
    for line in rg.section(text, "## Findings") or []:
        m = FINDING.match(line)
        if m:
            current = {"id": m.group(1), "severity": f"P{m.group(2)}", "status": "", "open_at": []}
            result.append(current)
            continue
        if current is None:
            continue
        s = STATUS.match(line)
        if s and not current["status"]:
            current["status"] = s.group(1).rstrip(".").strip().lower()
        o = OPEN_AT.match(line)
        if o:
            current["open_at"] = HASH.findall(o.group(1))
    return result


def is_open(finding):
    return finding["status"].startswith("open")


def rounds(finding, head):
    """The distinct heads at which the finding was open. An open finding counts the head too."""
    heads = {h.lower()[:rg.SHORTEST_HASH] for h in finding["open_at"]}
    if is_open(finding):
        heads.add(head.lower()[:rg.SHORTEST_HASH])
    return heads


def strikes(items, head):
    """The blocking findings open at STRIKES heads or more (D-826)."""
    return [f for f in items if is_open(f) and f["severity"] in BLOCKING and len(rounds(f, head)) >= STRIKES]


def read_result(run, repo, number, branch, before):
    must(run, ["git", "fetch", "--quiet", "origin", "main", branch], fault, cwd=repo)
    after = must(run, ["git", "rev-parse", f"origin/{branch}"], fault, cwd=repo).strip()
    if after == before:
        raise fault(f"origin/{branch} is still {before[:7]}, so the review pushed no record.")
    if run(["git", "merge-base", "--is-ancestor", before, after], cwd=repo)[0] != 0:
        raise fault(f"origin/{branch} at {after[:7]} does not hold the reviewed head {before[:7]}.")
    changed = must(run, ["git", "diff", "--name-only", "--no-renames", before, after], fault, cwd=repo).split()
    outside = [p for p in changed if p not in rg.metadata_paths(number)]
    if outside:
        raise fault(f"the review pushed a change of `{outside[0]}`, outside the metadata set of D-813.")
    _, commits, _, read = rg.gather(repo, "origin/main", f"origin/{branch}")
    head = rg.effective_head(commits, number)
    path = rg.record_path(number)
    text = read(path)
    if text is None:
        raise fault(f"origin/{branch} holds no `{path}`.")
    state, message = rg.check_head(path, text, head)
    if state != "PASS":
        raise fault(message)
    return head, verdict(text), findings(text)


def outcome(head, name, items):
    """The exit code and the report lines of a read record."""
    strike = strikes(items, head)
    open_items = [f for f in items if is_open(f)]
    lines = [f"verdict: {name}, head {head[:7]}"]
    lines.append("open findings: " + (", ".join(f"{f['id']} (round {len(rounds(f, head))})" for f in open_items) or "none"))
    if strike:
        lines.append("three-strike: " + ", ".join(f["id"] for f in strike) + f" open at {STRIKES} heads. Stop the fix loop, and ask the owner (D-826).")
        return EXIT_STRIKE, lines
    if name == rg.APPROVED:
        return EXIT_APPROVE, lines
    return EXIT_CHANGES, lines


def main(argv=None, run=sh):
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument("--pr", type=int, required=True, help="the number of the pull request")
    parser.add_argument("--repo", default=ROOT)
    try:
        args = parser.parse_args(argv)
    except SystemExit:
        return EXIT_USAGE
    code, lines = EXIT_FAULT, []
    try:
        slug = must(run, ["gh", "repo", "view", "--json", "nameWithOwner", "--jq", ".nameWithOwner"], refuse).strip()
        branch, head = check_pr(run, args.pr)
        check_checkout(run, args.repo, branch, head)
        _, commits, _, _ = rg.gather(args.repo, "origin/main", head)
        effective = rg.effective_head(commits, args.pr)
        if effective is None:
            raise refuse("every commit changes the metadata set alone, so the pull request has no effective head.")
        check_gitar(run, args.repo, slug, args.pr, effective, head)
        codex, installed = update_cli(run)
        check_login(run, codex)
        probe(run, codex)
        print(f"codex-review: PR #{args.pr}, head {head[:7]}, effective head {effective[:7]}, {installed}, {MODEL} at {EFFORT}.")
        stamp = datetime.datetime.now(datetime.timezone.utc).strftime("%Y%m%dT%H%M%SZ")
        status, tree, base = review(run, args.repo, codex, args.pr, slug, branch, head, stamp)
        kept = remove_tree(run, args.repo, tree, branch)
        lines.append(f"transcript: {os.path.relpath(base, args.repo)}.jsonl")
        if kept:
            lines.append(f"kept worktree: {kept}")
        if status != 0:
            raise fault(f"codex exited {status}. Read {os.path.relpath(base, args.repo)}.stderr.log.")
        reviewed, name, items = read_result(run, args.repo, args.pr, branch, head)
        code, report = outcome(reviewed, name, items)
        lines = report + lines + [f"the checkout is behind origin/{branch}. Run `git pull --ff-only`."]
    except Stop as stop:
        code = stop.code
        lines.append(f"codex-review: {stop}")
    for line in lines:
        print(line)
    print(f"outcome: {OUTCOMES[code]} (exit {code})")
    return code


if __name__ == "__main__":
    sys.exit(main())
