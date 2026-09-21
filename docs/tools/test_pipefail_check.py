"""Tests of the pipefail check (F-160).

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import importlib.util
import os
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(os.path.dirname(HERE))
spec = importlib.util.spec_from_file_location("pipefail_check", os.path.join(HERE, "pipefail_check.py"))
pc = importlib.util.module_from_spec(spec)
spec.loader.exec_module(pc)

PREAMBLE = "SHELL := bash\n.SHELLFLAGS := -o pipefail -c\n"


def fake_runner(fails=2, passes=0, bare=2, version="GNU Make 4.3\n"):
    """Return a runner that answers with canned exit codes."""

    def runner(_preamble):
        def run(args, text=False):
            if args == ["--version"]:
                return version if text else 0
            return {"guarded-fails": fails, "guarded-passes": passes, "bare": bare}[args[0]]

        return run

    return runner


def check(body, runner=None):
    return pc.check(lambda _path: PREAMBLE + body, runner=runner or fake_runner())


class TestPipeline(unittest.TestCase):
    def test_a_pipe_is_a_pipeline(self):
        self.assertTrue(pc.has_pipeline("\t@go run ./cmd/probe | tee out.md"))

    def test_a_quoted_pipe_is_data(self):
        self.assertFalse(pc.has_pipeline("\t@go test -run 'TestOne|TestTwo' ./..."))

    def test_an_or_is_no_pipeline(self):
        self.assertFalse(pc.has_pipeline("\t@test -f out.md || { echo no; exit 1; }"))

    def test_a_redirect_is_no_pipeline(self):
        self.assertFalse(pc.has_pipeline("\t@go run ./cmd/gate > out.md"))


class TestRecipeLines(unittest.TestCase):
    def test_a_backslash_joins_the_next_line(self):
        lines = pc.recipe_lines("t:\n\t@set -a && \\\n\t\tgo run x | tee out.md\n")
        self.assertEqual(len(lines), 1)
        number, text, _marker = lines[0]
        self.assertEqual(number, 2)
        self.assertIn("tee out.md", text)

    def test_the_marker_reaches_one_recipe_line(self):
        lines = pc.recipe_lines("t:\n# pipefail-ok: a reason\n\t@a | b\n\t@c | d\n")
        self.assertEqual([bool(pc.MARKER in marker) for _n, _t, marker in lines], [True, False])


class TestCheck(unittest.TestCase):
    def test_a_bare_pipeline_is_a_finding(self):
        _report, errors = check("probe:\n\t@go run ./cmd/probe | tee out.md\n")
        self.assertEqual(len(errors), 1)
        self.assertIn("Makefile:4 pipes", errors[0])

    def test_a_guarded_pipeline_passes(self):
        report, errors = check("probe:\n\t@set -o pipefail; go run ./cmd/probe | tee out.md\n")
        self.assertEqual(errors, [])
        self.assertIn("1 guarded", report[0])

    def test_a_marker_exempts_the_line(self):
        report, errors = check("help:\n# pipefail-ok: an empty list is no fault\n\t@grep x Makefile | awk '{print}'\n")
        self.assertEqual(errors, [])
        self.assertIn("1 exempt", report[0])

    def test_a_marker_without_a_pipeline_is_a_finding(self):
        _report, errors = check("t:\n# pipefail-ok: a reason\n\t@echo one\n")
        self.assertEqual(len(errors), 1)
        self.assertIn("holds no pipeline", errors[0])

    def test_no_shell_is_a_finding(self):
        _report, errors = pc.check(lambda _path: "t:\n\t@echo one\n", runner=fake_runner())
        self.assertEqual(len(errors), 1)
        self.assertIn("sets no SHELL", errors[0])

    def test_an_absent_makefile_is_a_finding(self):
        _report, errors = pc.check(lambda _path: None, runner=fake_runner())
        self.assertEqual(errors, ["Makefile is absent"])

    def test_a_make_that_hides_the_failure_is_a_finding(self):
        _report, errors = check("t:\n\t@echo one\n", runner=fake_runner(fails=0))
        self.assertEqual(len(errors), 1)
        self.assertIn("does not fail a guarded pipeline", errors[0])

    def test_a_shell_without_pipefail_is_a_finding(self):
        _report, errors = check("t:\n\t@echo one\n", runner=fake_runner(passes=1))
        self.assertEqual(len(errors), 1)
        self.assertIn("fails 'set -o pipefail' itself", errors[0])

    def test_the_report_names_a_make_that_ignores_shellflags(self):
        report, errors = check("t:\n\t@echo one\n", runner=fake_runner(bare=0))
        self.assertEqual(errors, [])
        self.assertIn("ignores .SHELLFLAGS", report[1])


class TestThisMachine(unittest.TestCase):
    """The live proof. It reads the make and the Makefile of this machine."""

    def test_the_makefile_of_the_repository_holds_the_rule(self):
        def read(path):
            with open(os.path.join(ROOT, path), encoding="utf-8") as handle:
                return handle.read()

        report, errors = pc.check(read)
        self.assertEqual(errors, [])
        self.assertIn("GNU Make", report[1])


if __name__ == "__main__":
    unittest.main()
