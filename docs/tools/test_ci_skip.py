"""Tests of ci_skip.py (D-818 to D-820)."""
import importlib.util
import os
import subprocess
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("ci_skip", os.path.join(HERE, "ci_skip.py"))
cs = importlib.util.module_from_spec(spec)
spec.loader.exec_module(cs)

GREEN = {"id": 1, "status": "completed", "conclusion": "success"}
RED = {"id": 2, "status": "completed", "conclusion": "failure"}
RUNNING = {"id": 3, "status": "in_progress", "conclusion": None}
DOCS = ["docs/decisions.md", "docs/SESSION-HANDOFF.md", "CLAUDE.md"]
CODE = ["go/internal/generate/manapass.go", "docs/decisions.md"]


class Decide(unittest.TestCase):
    def test_rule_1_skips_a_documentation_pull_request_after_a_green_code_pull_request(self):
        skip, reason = cs.decide(DOCS, None, 211, GREEN, None)
        self.assertTrue(skip, reason)
        self.assertIn("#211", reason)

    def test_rule_1_runs_after_a_red_or_unfinished_code_pull_request(self):
        for run in (RED, RUNNING, None):
            with self.subTest(run=run):
                skip, reason = cs.decide(DOCS, None, 211, run, None)
                self.assertFalse(skip, reason)

    def test_rule_1_runs_when_no_pull_request_holds_the_last_code_change(self):
        skip, _ = cs.decide(DOCS, None, None, None, None)
        self.assertFalse(skip)

    def test_rule_2_skips_a_documentation_push_after_a_green_previous_head(self):
        skip, reason = cs.decide(CODE, ["docs/SESSION-HANDOFF.md"], None, None, GREEN)
        self.assertTrue(skip, reason)
        self.assertIn("rule 2", reason)

    def test_rule_2_runs_after_a_red_cancelled_or_absent_previous_run(self):
        for run in (RED, RUNNING, None, {"id": 4, "status": "completed", "conclusion": "cancelled"}):
            with self.subTest(run=run):
                skip, _ = cs.decide(CODE, ["docs/a.md"], None, None, run)
                self.assertFalse(skip)

    def test_rule_2_saves_a_documentation_pull_request_whose_rule_1_fails(self):
        skip, reason = cs.decide(DOCS, ["docs/a.md"], 211, RED, GREEN)
        self.assertTrue(skip, reason)

    def test_a_push_of_code_runs(self):
        skip, reason = cs.decide(CODE, ["web/apps/web/src/a.ts"], None, None, GREEN)
        self.assertFalse(skip)
        self.assertIn("web/apps/web/src/a.ts", reason)

    def test_a_refused_path_is_code(self):
        for path in ("docs/tools/ci_skip.py", ".claude/hooks/session_bind.py", ".github/workflows/verify.yml", "Makefile"):
            with self.subTest(path=path):
                skip, _ = cs.decide(["docs/a.md", path], None, 211, GREEN, None)
                self.assertFalse(skip)

    def test_unknown_paths_run(self):
        skip, _ = cs.decide(None, None, 211, GREEN, GREEN)
        self.assertFalse(skip)


class LastCodeCommit(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.repo = self.tmp.name
        self.git("init", "-q", "-b", "main")

    def tearDown(self):
        self.tmp.cleanup()

    def git(self, *args):
        env = {**os.environ, "GIT_AUTHOR_NAME": "a", "GIT_AUTHOR_EMAIL": "a@example.com",
               "GIT_COMMITTER_NAME": "a", "GIT_COMMITTER_EMAIL": "a@example.com"}
        return subprocess.run(["git", *args], cwd=self.repo, capture_output=True, text=True, env=env, check=True).stdout.strip()

    def commit(self, path, message):
        full = os.path.join(self.repo, path)
        os.makedirs(os.path.dirname(full) or self.repo, exist_ok=True)
        with open(full, "a", encoding="utf-8") as handle:
            handle.write(message + "\n")
        self.git("add", path)
        self.git("commit", "-q", "-m", message)
        return self.git("rev-parse", "HEAD")

    def test_the_search_passes_documentation_commits(self):
        code = self.commit("go/a.go", "code")
        self.commit("docs/a.md", "docs")
        self.commit("CLAUDE.md", "rules")
        self.assertEqual(cs.last_code_commit(self.repo, "HEAD"), code)

    def test_a_tool_under_docs_is_code(self):
        self.commit("go/a.go", "code")
        tool = self.commit("docs/tools/x.py", "tool")
        self.commit("docs/a.md", "docs")
        self.assertEqual(cs.last_code_commit(self.repo, "HEAD"), tool)

    def test_a_history_of_documents_alone_gives_none(self):
        self.commit("docs/a.md", "docs")
        self.assertIsNone(cs.last_code_commit(self.repo, "HEAD"))

    def test_the_push_paths_read_the_previous_head(self):
        self.commit("go/a.go", "code")
        before = self.git("rev-parse", "HEAD")
        head = self.commit("docs/a.md", "docs")
        calls = []
        real = cs.gh_json
        cs.gh_json = lambda repo, path, jq: calls.append(path)
        try:
            facts = cs.gather(self.repo, "o/r", before, head, "synchronize", before)
        finally:
            cs.gh_json = real
        self.assertEqual(facts[1], ["docs/a.md"])
        self.assertEqual(facts[4], None)
        self.assertTrue(calls, "gather reads the runs through gh_json alone")

    def test_a_missing_binary_is_an_unknown_fact(self):
        self.assertIsNone(cs.run(self.repo, "decktome-no-such-binary"))


class Workflow(unittest.TestCase):
    def test_each_heavy_job_needs_the_skip_job_and_the_light_jobs_do_not(self):
        with open(os.path.join(cs.ROOT, ".github/workflows/verify.yml"), encoding="utf-8") as handle:
            text = handle.read()
        jobs = {}
        name = None
        for line in text.splitlines():
            if line.startswith("  ") and not line.startswith("   ") and line.rstrip().endswith(":"):
                name = line.strip()[:-1]
                jobs[name] = ""
            elif name:
                jobs[name] += line + "\n"
        for job in ("go", "vuln", "emulator", "web", "proto", "docker"):
            with self.subTest(job=job):
                self.assertIn("needs.skip.outputs.skip != 'true'", jobs[job])
        for job in ("shell", "eval"):
            with self.subTest(job=job):
                self.assertNotIn("needs.skip.outputs.skip", jobs[job])
        self.assertIn("git show \"$BASE_SHA:docs/tools/ci_skip.py\"", jobs["skip"])


if __name__ == "__main__":
    unittest.main()
