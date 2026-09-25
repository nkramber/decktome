"""Tests for deploy_order.py (REV-072)."""
import io
import os
import shutil
import subprocess
import tempfile
import unittest
from types import SimpleNamespace

import deploy_order as d


def git(repo, *args):
    return subprocess.run(["git", "-C", repo, *args], check=True, capture_output=True, text=True).stdout.strip()


def commit(repo, path):
    full = os.path.join(repo, path)
    os.makedirs(os.path.dirname(full), exist_ok=True)
    with open(full, "a", encoding="utf-8") as f:
        f.write("x\n")
    git(repo, "add", "-A")
    git(repo, "-c", "user.name=t", "-c", "user.email=t@example.invalid", "commit", "-q", "-m", path)
    return git(repo, "rev-parse", "HEAD")


class Clock:
    """A clock that each sleep moves forward."""

    def __init__(self):
        self.now = 0.0
        self.sleeps = 0

    def __call__(self):
        return self.now

    def sleep(self, seconds):
        self.sleeps += 1
        self.now += seconds


class DeployOrderTest(unittest.TestCase):
    def setUp(self):
        self.repo = tempfile.mkdtemp()
        git(self.repo, "init", "-q", "-b", "main")
        self.base = commit(self.repo, "README.md")
        self.api = commit(self.repo, "go/a.go")
        self.web = commit(self.repo, "web/b.ts")
        self.docs = commit(self.repo, "docs/c.md")
        self.history = d.History(self.repo)
        self.skip_file = os.path.join(tempfile.mkdtemp(), "skip")
        for path in (self.repo, self.history.dir, os.path.dirname(self.skip_file)):
            self.addCleanup(shutil.rmtree, path, True)

    def args(self, mine, **kw):
        base = dict(commit=mine, readyz="api", web="web", skip_file=self.skip_file, interval=15, wait_seconds=60)
        base.update(kw)
        return SimpleNamespace(**base)

    def fetch(self, api=None, web=None):
        """Answer /readyz from the api list in turn, and /version.json with web."""
        api = list(api or [])

        def get(url):
            if url == "api":
                v = api.pop(0) if len(api) > 1 else (api[0] if api else None)
                return None if v is None else {"status": "ok", "version": v}
            return None if web is None else {"commit": web}

        return get

    def test_api_guard(self):
        cases = [
            ("a later live API skips the deploy", self.api, self.docs, True),
            ("an older live API deploys", self.docs, self.api, False),
            ("the same commit deploys again", self.api, self.api, False),
            ("a local build names no commit", self.api, "dev", False),
            ("no answer deploys", self.api, None, False),
            ("a commit that git lacks deploys", self.api, "f" * 40, False),
        ]
        for name, mine, live, skipped in cases:
            with self.subTest(name):
                if os.path.exists(self.skip_file):
                    os.remove(self.skip_file)
                out = io.StringIO()
                rc = d.api_guard(self.args(mine), self.history, fetch=self.fetch(api=[live]), out=out)
                self.assertEqual(rc, 0)
                self.assertEqual(os.path.exists(self.skip_file), skipped, out.getvalue())

    def test_web_of_a_web_change_does_not_wait(self):
        clock = Clock()
        out = io.StringIO()
        rc = d.web_guard(self.args(self.web), self.history, fetch=self.fetch(api=[self.base]), sleep=clock.sleep, clock=clock, out=out)
        self.assertEqual(rc, 0)
        self.assertEqual(clock.sleeps, 0)
        self.assertFalse(os.path.exists(self.skip_file))

    def test_web_of_an_api_change_waits_for_its_api(self):
        clock = Clock()
        out = io.StringIO()
        fetch = self.fetch(api=[self.base, self.base, self.api])
        rc = d.web_guard(self.args(self.api), self.history, fetch=fetch, sleep=clock.sleep, clock=clock, out=out)
        self.assertEqual(rc, 0, out.getvalue())
        self.assertEqual(clock.sleeps, 2)
        self.assertIn("so the web can follow", out.getvalue())

    def test_web_of_an_api_change_follows_a_later_api(self):
        clock = Clock()
        rc = d.web_guard(self.args(self.api), self.history, fetch=self.fetch(api=[self.docs]), sleep=clock.sleep, clock=clock, out=io.StringIO())
        self.assertEqual(rc, 0)
        self.assertEqual(clock.sleeps, 0)

    def test_web_fails_when_its_api_never_arrives(self):
        clock = Clock()
        out = io.StringIO()
        rc = d.web_guard(self.args(self.api), self.history, fetch=self.fetch(api=[self.base]), sleep=clock.sleep, clock=clock, out=out)
        self.assertEqual(rc, 1)
        self.assertIn("does not go live before its API", out.getvalue())
        self.assertFalse(os.path.exists(self.skip_file))

    def test_web_skips_when_the_live_web_is_later(self):
        clock = Clock()
        rc = d.web_guard(self.args(self.web), self.history, fetch=self.fetch(web=self.docs), sleep=clock.sleep, clock=clock, out=io.StringIO())
        self.assertEqual(rc, 0)
        self.assertTrue(os.path.exists(self.skip_file))

    def test_web_releases_over_an_older_live_web(self):
        clock = Clock()
        rc = d.web_guard(self.args(self.docs), self.history, fetch=self.fetch(web=self.web), sleep=clock.sleep, clock=clock, out=io.StringIO())
        self.assertEqual(rc, 0)
        self.assertFalse(os.path.exists(self.skip_file))

    def test_a_merge_after_the_fetch_is_placed(self):
        later = commit(self.repo, "go/d.go")
        self.assertTrue(self.history.at_or_after(self.api, later))
        self.assertFalse(self.history.at_or_after(later, self.api))

    def test_api_ready(self):
        cases = [
            ("ok with the commit", [{"status": "starting", "version": self.api}, {"status": "ok", "version": self.api}], 0),
            ("ok with a later commit", [{"status": "ok", "version": self.docs}], 0),
            ("starting with a later commit", [{"status": "starting", "version": self.docs}], 1),
            ("ok with an older commit", [{"status": "ok", "version": self.base}], 1),
            ("ok with a local build", [{"status": "ok", "version": "dev"}], 1),
            ("no answer", [None], 1),
        ]
        for name, bodies, want in cases:
            with self.subTest(name):
                clock = Clock()
                seq = list(bodies)

                def fetch(_url, seq=seq):
                    return seq.pop(0) if len(seq) > 1 else seq[0]

                rc = d.api_ready(self.args(self.api), self.history, fetch=fetch, sleep=clock.sleep, clock=clock, out=io.StringIO())
                self.assertEqual(rc, want)

    def test_web_ready(self):
        cases = [
            ("the commit", self.web, 0),
            ("a later commit", self.docs, 0),
            ("an older commit", self.api, 1),
            ("no commit", None, 1),
        ]
        for name, live, want in cases:
            with self.subTest(name):
                clock = Clock()
                rc = d.web_ready(self.args(self.web), self.history, fetch=self.fetch(web=live), sleep=clock.sleep, clock=clock, out=io.StringIO())
                self.assertEqual(rc, want)

    def test_fetch_json_reads_the_body_of_a_503(self):
        import http.server
        import threading

        class Handler(http.server.BaseHTTPRequestHandler):
            def do_GET(self):
                self.send_response(503)
                self.end_headers()
                self.wfile.write(b'{"status": "starting", "version": "dev"}')

            def log_message(self, *a):
                pass

        srv = http.server.HTTPServer(("127.0.0.1", 0), Handler)
        threading.Thread(target=srv.serve_forever, daemon=True).start()
        try:
            data = d.fetch_json(f"http://127.0.0.1:{srv.server_port}/readyz", timeout=5)
        finally:
            srv.shutdown()
            srv.server_close()
        self.assertEqual(data, {"status": "starting", "version": "dev"})

    def test_no_git_deploys(self):
        def run(*_a, **_kw):
            raise FileNotFoundError("git")

        history = d.History(self.repo, run=run)
        self.addCleanup(shutil.rmtree, history.dir, True)
        rc = d.api_guard(self.args(self.api), history, fetch=self.fetch(api=[self.docs]), out=io.StringIO())
        self.assertEqual(rc, 0)
        self.assertFalse(os.path.exists(self.skip_file))
        clock = Clock()
        rc = d.web_guard(self.args(self.api), history, fetch=self.fetch(api=[self.base]), sleep=clock.sleep, clock=clock, out=io.StringIO())
        self.assertEqual(rc, 0)
        self.assertEqual(clock.sleeps, 0)

    def test_main_refuses_a_short_commit(self):
        self.assertEqual(d.main(["api-ready", "--commit", "abc", "--readyz", "x"]), 2)


ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


def steps(path):
    """Return the id and the text of each step of a Cloud Build file, in order."""
    out = []
    with open(os.path.join(ROOT, path), encoding="utf-8") as f:
        for line in f:
            if line.startswith("  - id: "):
                out.append([line.split(":", 1)[1].strip(), ""])
            elif line and not line[0].isspace() and not line.startswith("#"):
                if out and not line.startswith("steps:"):
                    break
            elif out:
                out[-1][1] += line
    return out


class BuildFilesTest(unittest.TestCase):
    """Each deploy step after the guard honors its mark, and each build stamps its commit."""

    def check(self, path, first):
        got = steps(path)
        ids = [i for i, _ in got]
        self.assertIn("guard", ids, path)
        after = got[ids.index("guard") + 1:]
        self.assertEqual(after[0][0], first, path)
        for sid, text in after:
            with self.subTest(path=path, step=sid):
                self.assertIn("/workspace/.deploy-skip", text)
        return dict(got)

    def test_api(self):
        got = self.check("cloudbuild/api.yaml", "deploy-api")
        self.assertIn("--build-arg=VERSION=$COMMIT_SHA", got["build-api"])
        self.assertIn("deploy_order.py api-ready", got["check"])

    def test_web(self):
        got = self.check("cloudbuild/web.yaml", "release")
        self.assertIn("dist/version.json", got["build"])
        self.assertIn("deploy_order.py web-ready", got["check"])


if __name__ == "__main__":
    unittest.main()
