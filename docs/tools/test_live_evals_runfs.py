"""Tests of live_evals_runfs.py: a planted link never leads a step out (D-1141)."""

import io
import os
import sys
import tempfile
import unittest
from contextlib import redirect_stderr, redirect_stdout

sys.path.insert(0, os.path.dirname(__file__))
import live_evals_runfs as runfs  # noqa: E402


class RunFSTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        base = self.tmp.name
        self.home = os.path.join(base, "home")
        self.owner = os.path.join(base, "owner")
        self.run = os.path.join(self.home, "deck1")
        os.makedirs(os.path.join(self.run, "bundle"))
        os.makedirs(self.owner)
        self.secret = os.path.join(self.owner, "secret")
        with open(self.secret, "w") as handle:
            handle.write("owner secret")

    def tearDown(self):
        self.tmp.cleanup()

    def call(self, *args):
        out, err = io.BytesIO(), io.StringIO()
        wrapper = io.TextIOWrapper(out, encoding="utf-8")
        with redirect_stdout(wrapper), redirect_stderr(err):
            code = runfs.main(["runfs", *args])
            wrapper.flush()
        return code, out.getvalue().decode(), err.getvalue()

    def write(self, rel, text):
        with open(os.path.join(self.run, rel), "w") as handle:
            handle.write(text)

    def test_get_reads_a_regular_file(self):
        self.write("bundle/result.json", '{"status": "ready"}')
        code, out, _ = self.call("get", self.home, "deck1", "bundle/result.json", "100")
        self.assertEqual(code, 0)
        self.assertEqual(out, '{"status": "ready"}')

    def test_get_refuses_a_linked_file(self):
        os.symlink(self.secret, os.path.join(self.run, "bundle", "spend.jsonl"))
        code, out, _ = self.call("get", self.home, "deck1", "bundle/spend.jsonl", "100")
        self.assertEqual(code, 1)
        self.assertNotIn("owner secret", out)

    def test_get_refuses_a_linked_folder(self):
        os.rmdir(os.path.join(self.run, "bundle"))
        os.symlink(self.owner, os.path.join(self.run, "bundle"))
        code, out, _ = self.call("get", self.home, "deck1", "bundle/secret", "100")
        self.assertEqual(code, 1)
        self.assertNotIn("owner secret", out)

    def test_get_refuses_a_linked_run_folder(self):
        other = os.path.join(self.home, "deck2")
        os.symlink(self.owner, other)
        code, out, _ = self.call("get", self.home, "deck2", "secret", "100")
        self.assertEqual(code, 1)
        self.assertNotIn("owner secret", out)

    def test_get_refuses_a_fifo_and_a_big_file(self):
        os.mkfifo(os.path.join(self.run, "bundle", "fix.json"))
        self.assertEqual(self.call("get", self.home, "deck1", "bundle/fix.json", "100")[0], 1)
        self.write("bundle/findings.md", "x" * 200)
        self.assertEqual(self.call("get", self.home, "deck1", "bundle/findings.md", "100")[0], 1)

    def test_put_replaces_a_linked_folder_and_keeps_the_owner_file(self):
        os.rmdir(os.path.join(self.run, "bundle"))
        os.symlink(self.owner, os.path.join(self.run, "bundle"))
        src = os.path.join(self.tmp.name, "build")
        os.makedirs(os.path.join(src, "chain"))
        with open(os.path.join(src, "deck.json"), "w") as handle:
            handle.write("{}")
        with open(os.path.join(src, "chain", "01.json"), "w") as handle:
            handle.write("{}")
        code, _, _ = self.call("put", self.home, "deck1", "bundle", src)
        self.assertEqual(code, 0)
        self.assertFalse(os.path.islink(os.path.join(self.run, "bundle")))
        self.assertTrue(os.path.isfile(os.path.join(self.run, "bundle", "chain", "01.json")))
        with open(self.secret) as handle:
            self.assertEqual(handle.read(), "owner secret")

    def test_rm_removes_links_and_never_their_targets(self):
        repo = os.path.join(self.run, "repo")
        os.makedirs(repo)
        os.symlink(self.owner, os.path.join(repo, "link"))
        code, _, _ = self.call("rm", self.home, "deck1", "repo")
        self.assertEqual(code, 0)
        self.assertFalse(os.path.exists(repo))
        self.assertTrue(os.path.isfile(self.secret))

    def test_rm_refuses_through_a_linked_run_folder(self):
        os.symlink(self.owner, os.path.join(self.home, "deck2"))
        self.assertEqual(self.call("rm", self.home, "deck2", "secret")[0], 1)
        self.assertTrue(os.path.isfile(self.secret))

    def test_mkdir_refuses_a_linked_part(self):
        os.symlink(self.owner, os.path.join(self.run, "tmp"))
        self.assertEqual(self.call("mkdir", self.home, "deck1", "tmp/x")[0], 1)
        self.assertFalse(os.path.exists(os.path.join(self.owner, "x")))
        self.assertEqual(self.call("mkdir", self.home, "deck1", "logs", "replay")[0], 0)

    def test_clear_removes_a_left_run_folder(self):
        os.makedirs(os.path.join(self.run, "repo", ".git"))
        self.assertEqual(self.call("clear", self.home, "deck1")[0], 0)
        self.assertFalse(os.path.exists(self.run))
        self.assertEqual(self.call("clear", self.home, "deck1")[0], 0)

    def test_clear_removes_a_linked_run_folder_and_never_its_target(self):
        os.symlink(self.owner, os.path.join(self.home, "deck2"))
        self.assertEqual(self.call("clear", self.home, "deck2")[0], 0)
        self.assertFalse(os.path.lexists(os.path.join(self.home, "deck2")))
        self.assertTrue(os.path.isfile(self.secret))
        self.assertEqual(self.call("clear", self.home, "..")[0], 1)

    def test_paths_with_a_parent_part_are_refused(self):
        self.assertEqual(self.call("get", self.home, "deck1", "../deck1/bundle/x", "10")[0], 1)
        self.assertEqual(self.call("rm", self.home, "..", "x")[0], 1)


if __name__ == "__main__":
    unittest.main()
