"""Order the deploys of the API and the web, and skip a stale deploy (REV-072).

Two Cloud Build triggers deploy the API and the web (D-584). Cloud Build
has no queue, so the build of an older merge can finish last, and the web
of a merge can go live before its API. Each build runs this script before
its deploy step:

- `api-guard`: the API build skips its deploy when the live API runs a
  newer commit.
- `web-guard`: when the merge changed `go/` or `docker/`, the web build
  waits until the API runs the commit of the merge or a newer one. Then it
  skips its release when the live web runs a newer commit.
- `api-ready`: the API build waits until `/readyz` reads ok and names the
  commit of the build.

The API reports its commit in the `version` field of `/readyz`, and the
web in `/version.json`. A live commit that no read gives, or that git can
not place, never stops a deploy. The script says so in the log, and the
deploy runs as before.

Run: python3 docs/tools/deploy_order.py api-guard --commit SHA --readyz URL --skip-file F
"""
import argparse
import json
import re
import subprocess
import sys
import tempfile
import time
import urllib.error
import urllib.request

# The repository is public (D-639), so a build reads its history with no
# credential.
REPO = "https://github.com/nkramber/decktome.git"

# A change under these paths starts the API trigger (D-584).
API_PATHS = ("go/", "docker/")

SHA = re.compile(r"^[0-9a-f]{40}$")


def fetch_json(url, timeout=60):
    """Return the JSON object at url, or None. A 503 of /readyz still holds a body."""
    try:
        with urllib.request.urlopen(url, timeout=timeout) as res:
            body = res.read()
    except urllib.error.HTTPError as err:
        with err:
            body = err.read()
    except (urllib.error.URLError, OSError):
        return None
    try:
        data = json.loads(body)
    except ValueError:
        return None
    return data if isinstance(data, dict) else None


def live_commit(data, key):
    """Return the commit that a status body names, or "" when it names none."""
    value = (data or {}).get(key, "")
    return value if isinstance(value, str) and SHA.match(value) else ""


class History:
    """The history of `main`, fetched with no file content."""

    def __init__(self, repo=REPO, run=subprocess.run):
        self.repo = repo
        self.run = run
        self.dir = tempfile.mkdtemp(prefix="deploy-order-")
        self.git("init", "-q")
        self.fetch()

    def git(self, *args):
        cmd = ["git", "-C", self.dir, *args]
        try:
            return self.run(cmd, capture_output=True, text=True)
        except OSError as err:
            # An image with no git places no commit, and the deploy runs.
            return subprocess.CompletedProcess(cmd, 127, "", str(err))

    def fetch(self):
        return self.git("fetch", "-q", "--no-tags", "--filter=blob:none", self.repo, "main").returncode == 0

    def known(self, sha):
        return self.git("cat-file", "-e", f"{sha}^{{commit}}").returncode == 0

    def at_or_after(self, mine, live):
        """Return True when live is mine or a later commit, False when not, and None when git can not tell."""
        if mine == live:
            return True
        # A merge after the first fetch gives a commit that the fetch lacks.
        if not self.known(live) and not (self.fetch() and self.known(live)):
            return None
        if not self.known(mine):
            return None
        rc = self.git("merge-base", "--is-ancestor", mine, live).returncode
        return {0: True, 1: False}.get(rc)

    def changed(self, sha):
        """Return the paths that the commit changed against its first parent, or None."""
        res = self.git("diff", "--name-only", f"{sha}^1", sha)
        if res.returncode != 0:
            return None
        return [line for line in res.stdout.splitlines() if line]


def newer_live(history, mine, live, name, out):
    """Report whether the live commit is later than mine. An unknown commit is not later."""
    if not live:
        print(f"deploy_order: the live {name} names no commit, so the deploy runs", file=out)
        return False
    later = history.at_or_after(mine, live)
    if later is None:
        print(f"deploy_order: git can not place the live {name} commit {live}, so the deploy runs", file=out)
        return False
    if later and live != mine:
        print(f"deploy_order: the live {name} runs {live}, which is later than {mine}. A later build deployed it", file=out)
        return True
    return False


def skip(skip_file, out):
    with open(skip_file, "w", encoding="utf-8") as f:
        f.write("a later build deployed\n")
    print(f"deploy_order: the deploy steps skip, and {skip_file} marks it", file=out)


def api_guard(args, history, fetch=fetch_json, out=sys.stdout):
    live = live_commit(fetch(args.readyz), "version")
    if newer_live(history, args.commit, live, "API", out):
        skip(args.skip_file, out)
    else:
        print(f"deploy_order: the API deploys {args.commit}", file=out)
    return 0


def wait_for_api(args, history, fetch, sleep, clock, out):
    """Wait until the API runs the commit of the merge or a later one. Return False at the time limit."""
    end = clock() + args.wait_seconds
    while True:
        live = live_commit(fetch(args.readyz), "version")
        if live and history.at_or_after(args.commit, live):
            print(f"deploy_order: the API runs {live}, so the web can follow", file=out)
            return True
        if clock() >= end:
            return False
        print(f"deploy_order: the API runs {live or 'no known commit'}, and the web waits for {args.commit}", file=out)
        sleep(args.interval)


def web_guard(args, history, fetch=fetch_json, sleep=time.sleep, clock=time.monotonic, out=sys.stdout):
    changed = history.changed(args.commit)
    if changed is None:
        print(f"deploy_order: git can not read the change of {args.commit}, so the web does not wait", file=out)
    elif any(p.startswith(API_PATHS) for p in changed):
        if not wait_for_api(args, history, fetch, sleep, clock, out):
            print(f"deploy_order: the API never ran {args.commit} in {args.wait_seconds} seconds."
                  " The web does not go live before its API. Read the deploy-api build of this commit", file=out)
            return 1
    else:
        print("deploy_order: the merge changed no API path, so the web does not wait", file=out)
    # The wait can be long, so the check of the live web comes after it.
    live = live_commit(fetch(args.web), "commit")
    if newer_live(history, args.commit, live, "web", out):
        skip(args.skip_file, out)
    else:
        print(f"deploy_order: the web releases {args.commit}", file=out)
    return 0


def api_ready(args, fetch=fetch_json, sleep=time.sleep, clock=time.monotonic, out=sys.stdout):
    end = clock() + args.wait_seconds
    while True:
        data = fetch(args.readyz) or {}
        status, version = data.get("status", ""), data.get("version", "")
        print(f"deploy_order: /readyz status {status or 'none'}, version {version or 'none'}", file=out)
        if status == "ok" and version == args.commit:
            return 0
        if clock() >= end:
            print(f"deploy_order: the API never read ok with {args.commit}", file=out)
            return 1
        sleep(args.interval)


def parse(argv):
    p = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    sub = p.add_subparsers(dest="cmd", required=True)
    for name in ("api-guard", "web-guard", "api-ready"):
        s = sub.add_parser(name)
        s.add_argument("--commit", required=True)
        s.add_argument("--readyz", required=True)
        s.add_argument("--repo", default=REPO)
        s.add_argument("--interval", type=float, default=15)
        if name != "api-ready":
            s.add_argument("--skip-file", required=True)
        if name == "web-guard":
            s.add_argument("--web", required=True)
        s.add_argument("--wait-seconds", type=float, default=1500 if name == "web-guard" else 300)
    return p.parse_args(argv)


def main(argv=None):
    args = parse(sys.argv[1:] if argv is None else argv)
    if not SHA.match(args.commit):
        print(f"deploy_order: {args.commit!r} is not a full commit id", file=sys.stdout)
        return 2
    if args.cmd == "api-ready":
        return api_ready(args)
    history = History(args.repo)
    if args.cmd == "api-guard":
        return api_guard(args, history)
    return web_guard(args, history)


if __name__ == "__main__":
    sys.exit(main())
