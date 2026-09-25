"""Tests of the tree and branch guard of scripts/feedback-review.sh (REV-074).

A round resets the tree and pushes it. A change of a person in the same
checkout must survive a round, and a round must never reset or push a
branch it did not start on. Each test runs the script in a temporary git
repository with a fake gh and a fake fixer, so no test reaches GitHub,
a provider, or the git state of this checkout.

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import os
import shutil
import stat
import subprocess
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(os.path.dirname(HERE))

FAKE_GH = """#!/usr/bin/env bash
case "$1 $2" in
  "repo view") echo owner/repo ;;
  "pr view") git rev-parse HEAD ;;
  "pr checks") echo '[{"name":"verify","bucket":"pass"},{"name":"Gitar","bucket":"pass"}]' ;;
  "api graphql")
    if printf '%s\\n' "$@" | grep -q reviewThreads; then
      echo '{"id":"T1","path":"f.txt","line":1,"author":"gitar-bot","body":"a finding"}'
    fi ;;
esac
exit 0
"""


def write(path, text, mode=None):
    os.makedirs(os.path.dirname(path), exist_ok=True)
    with open(path, "w", encoding="utf-8") as f:
        f.write(text)
    if mode:
        os.chmod(path, os.stat(path).st_mode | mode)


class FeedbackReviewGuardTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.mkdtemp()
        self.addCleanup(shutil.rmtree, self.tmp, True)
        self.repo = os.path.join(self.tmp, "repo")
        self.remote = os.path.join(self.tmp, "remote.git")
        self.env = {k: v for k, v in os.environ.items() if not k.startswith("GIT_")}
        self.env.update(GIT_CONFIG_GLOBAL=os.devnull, GIT_CONFIG_NOSYSTEM="1", GOTOOLCHAIN="local", GOFLAGS="")
        bindir = os.path.join(self.tmp, "bin")
        write(os.path.join(bindir, "gh"), FAKE_GH, stat.S_IXUSR)
        self.env["PATH"] = bindir + os.pathsep + self.env["PATH"]

    def git(self, *args, cwd=None):
        res = subprocess.run(["git", *args], cwd=cwd or self.repo, env=self.env,
                             capture_output=True, text=True, timeout=60)
        if res.returncode != 0:
            raise AssertionError(f"git {' '.join(args)}: {res.stderr}")
        return res.stdout.strip()

    def make_repo(self, fixer):
        """A repository on branch work, with origin, a fake fixer, and a tiny Go module."""
        self.git("init", "-q", "--bare", "-b", "main", self.remote, cwd=self.tmp)
        os.makedirs(self.repo)
        self.git("init", "-q", "-b", "main")
        self.git("config", "user.name", "t")
        self.git("config", "user.email", "t@example.com")
        self.git("config", "commit.gpgsign", "false")
        os.makedirs(os.path.join(self.repo, "scripts"))
        for name in ("feedback-review.sh", "feedback-loop.sh"):
            shutil.copy(os.path.join(ROOT, "scripts", name), os.path.join(self.repo, "scripts", name))
        write(os.path.join(self.repo, "scripts", "autotune-fix.sh"), "#!/usr/bin/env bash\n" + fixer, stat.S_IXUSR)
        write(os.path.join(self.repo, "go", "go.mod"), "module x\n\ngo 1.21\n")
        write(os.path.join(self.repo, "go", "x.go"), "package x\n")
        write(os.path.join(self.repo, "Makefile"), "lint-go:\n\t@true\n")
        write(os.path.join(self.repo, "f.txt"), "base\n")
        self.git("add", "-A")
        self.git("commit", "-qm", "base")
        self.git("remote", "add", "origin", self.remote)
        self.git("push", "-q", "origin", "main")
        for branch in ("other", "work"):
            self.git("switch", "-q", "-c", branch, "main")
            self.git("push", "-q", "-u", "origin", branch)

    def run_round(self):
        env = dict(self.env, FEEDBACK_REVIEW_POLL="0", FEEDBACK_REVIEW_WAIT="5", FEEDBACK_CI_WAIT="5",
                   GITAR_PAUSE_FILE=os.path.join(self.tmp, "no-pause"), CODEX_REVIEW_CMD="false")
        return subprocess.run(["bash", "scripts/feedback-review.sh", "7", "1", os.path.join(self.tmp, "state")],
                              cwd=self.repo, env=env, capture_output=True, text=True, timeout=300)

    def read(self, name):
        path = os.path.join(self.repo, name)
        if not os.path.exists(path):
            return None
        with open(path, encoding="utf-8") as f:
            return f.read()

    def test_a_tracked_change_is_refused_and_kept(self):
        self.make_repo("exit 1\n")
        write(os.path.join(self.repo, "f.txt"), "mine\n")
        res = self.run_round()
        self.assertNotEqual(res.returncode, 0)
        self.assertEqual(self.read("f.txt"), "mine\n", res.stderr)
        self.assertIn("uncommitted", res.stderr)

    def test_an_untracked_file_is_refused_and_kept(self):
        # A fixer that commits with git add -A takes the file, and the revert of its round deletes it.
        self.make_repo("git add -A && git commit -qm fix\nexit 1\n")
        write(os.path.join(self.repo, "notes.txt"), "mine\n")
        res = self.run_round()
        self.assertNotEqual(res.returncode, 0)
        self.assertEqual(self.read("notes.txt"), "mine\n", res.stderr)
        self.assertIn("untracked", res.stderr)

    def test_a_branch_switch_stops_the_reset(self):
        self.make_repo("git switch -q other\necho x >> f.txt\ngit commit -qam x\nexit 1\n")
        res = self.run_round()
        self.assertNotEqual(res.returncode, 0)
        self.assertEqual(self.git("log", "-1", "--format=%s", "other"), "x", res.stderr)
        self.assertIn("'other'", res.stderr)
        self.assertIn("'work'", res.stderr)

    def test_a_branch_switch_stops_the_push(self):
        self.make_repo("git switch -q other\necho y >> f.txt\ngit commit -qam y\nexit 0\n")
        pushed = self.git("rev-parse", "other", cwd=self.remote)
        res = self.run_round()
        self.assertNotEqual(res.returncode, 0)
        self.assertEqual(self.git("rev-parse", "other", cwd=self.remote), pushed, res.stderr)
        self.assertIn("'other'", res.stderr)

    def test_a_fix_on_the_start_branch_is_pushed(self):
        self.make_repo("echo z >> f.txt\ngit commit -qam z\nexit 0\n")
        res = self.run_round()
        self.assertEqual(self.git("log", "-1", "--format=%s", "work", cwd=self.remote), "z", res.stderr)
        self.assertIn("pushed 1 commit(s)", res.stderr)

    def test_a_fixer_that_edits_the_frozen_list_is_refused(self):
        # D-930: the round reads the list of its start commit. A fixer that
        # removes two lines of the list, then edits one of the two paths,
        # must still fail the round.
        fixer = ("sed -i.bak -e '/^scripts\\/feedback-loop.sh$/d' -e '/^docs\\/owner-questions.md$/d' scripts/feedback-loop.sh\n"
                 "rm -f scripts/feedback-loop.sh.bak\n"
                 "mkdir -p docs && echo edited > docs/owner-questions.md\n"
                 "git add -A && git commit -qm fix\nexit 0\n")
        self.make_repo(fixer)
        pushed = self.git("rev-parse", "work", cwd=self.remote)
        res = self.run_round()
        self.assertEqual(self.git("rev-parse", "work", cwd=self.remote), pushed, res.stderr)
        self.assertIn("frozen path", res.stderr)


class FrozenListTest(unittest.TestCase):
    """The list covers the gates, the review, the checks, and the deploy (D-930)."""

    def frozen(self):
        with open(os.path.join(ROOT, "scripts", "feedback-loop.sh"), encoding="utf-8") as f:
            text = f.read()
        block = text.split('\nFROZEN="\n', 1)[1].split('\n"\n', 1)[0]
        return [line.strip() for line in block.splitlines() if line.strip()]

    def test_each_listed_path_exists(self):
        for path in self.frozen():
            self.assertTrue(os.path.exists(os.path.join(ROOT, path)), path)

    def test_the_gate_review_check_and_deploy_paths_are_frozen(self):
        frozen = self.frozen()
        gates = sorted(d for d in os.listdir(os.path.join(ROOT, "go", "cmd")) if d.endswith("-gate"))
        self.assertTrue(gates)
        for path in [*(f"go/cmd/{d}" for d in gates), "go/internal/gatekit", "go/internal/evalrun",
                     "scripts/feedback-loop.sh", "scripts/feedback-review.sh", "scripts/autotune-fix.sh",
                     "docs/tools", "docs/reviews", ".github", ".claude", "cloudbuild", "Makefile",
                     "firebase.json", "firestore.rules", "firestore.indexes.json"]:
            self.assertIn(path, frozen)


if __name__ == "__main__":
    unittest.main()
