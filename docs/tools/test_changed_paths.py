"""Tests of the docs-only skip of the verify workflow (D-805).

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import contextlib
import importlib.util
import io
import json
import os
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("changed_paths", os.path.join(HERE, "changed_paths.py"))
cp = importlib.util.module_from_spec(spec)
spec.loader.exec_module(cp)

GREEN = [{"name": "verify:gate", "status": "completed", "conclusion": "success", "app": "github-actions"}]


def run(name, status="completed", conclusion="success", app="github-actions"):
    return {"name": name, "status": status, "conclusion": conclusion, "app": app}


class DocsSetTest(unittest.TestCase):
    def test_docs_and_skills_and_root_guides_are_in_the_set(self):
        for path in ("docs/decisions.md", "docs/tools/pr_check.py", ".claude/skills/x/SKILL.md",
                     ".claude/settings.json", "CLAUDE.md", "AGENTS.md", "README.md",
                     ".github/pull_request_template.md"):
            self.assertTrue(cp.in_docs_set(path), path)

    def test_code_read_folders_and_workflows_are_outside_the_set(self):
        # Go tests glob docs/reference/pr7-*, and the eval check reads docs/reference/eval/.
        for path in ("docs/reference/eval/base.jsonl", "docs/reference/pr7-question-gate-run1.md",
                     ".github/workflows/verify.yml", "go/go.mod", "Makefile", "scripts/where.sh",
                     "docsx/a.md", "CLAUDE.md.bak"):
            self.assertFalse(cp.in_docs_set(path), path)

    def test_an_empty_path_is_a_fault(self):
        with self.assertRaises(cp.FactError):
            cp.in_docs_set("")


class DecideTest(unittest.TestCase):
    def test_another_event_runs_every_job(self):
        self.assertEqual(cp.decide("workflow_dispatch", None, None, None)[0], False)

    def test_another_event_with_a_path_file_is_a_fault(self):
        with self.assertRaises(cp.FactError):
            cp.decide("schedule", ["docs/a.md"], None, None)

    def test_a_pull_request_needs_its_paths(self):
        with self.assertRaises(cp.FactError):
            cp.decide("pull_request", None, None, None)

    def test_push_paths_and_checks_go_together(self):
        with self.assertRaises(cp.FactError):
            cp.decide("pull_request", ["go/a.go"], ["docs/a.md"], None)

    def test_an_empty_change_runs_every_job(self):
        self.assertFalse(cp.decide("pull_request", [], None, None)[0])

    def test_a_docs_only_pull_request_skips_with_no_previous_head(self):
        self.assertTrue(cp.decide("pull_request", ["docs/a.md", "CLAUDE.md"], None, None)[0])

    def test_a_code_pull_request_with_no_previous_head_runs(self):
        alone, reason = cp.decide("pull_request", ["docs/a.md", "go/a.go"], None, None)
        self.assertFalse(alone)
        self.assertIn("go/a.go", reason)

    def test_a_docs_push_after_a_green_head_skips(self):
        self.assertTrue(cp.decide("pull_request", ["go/a.go", "docs/a.md"], ["docs/a.md"], GREEN)[0])

    def test_a_code_push_runs(self):
        self.assertFalse(cp.decide("pull_request", ["go/a.go"], ["go/a.go", "docs/a.md"], GREEN)[0])

    def test_an_empty_push_runs(self):
        self.assertFalse(cp.decide("pull_request", ["go/a.go"], [], GREEN)[0])

    def test_a_docs_push_after_a_red_head_runs(self):
        red = [run("verify:gate", conclusion="failure")]
        alone, reason = cp.decide("pull_request", ["go/a.go"], ["docs/a.md"], red)
        self.assertFalse(alone)
        self.assertIn("failure", reason)

    def test_a_docs_push_after_a_cancelled_or_running_head_runs(self):
        for checks in ([run("verify:gate", conclusion="cancelled")],
                       [run("verify:gate", status="in_progress", conclusion=None)]):
            self.assertFalse(cp.decide("pull_request", ["go/a.go"], ["docs/a.md"], checks)[0])

    def test_a_docs_push_with_no_gate_run_runs(self):
        other = [run("verify:go"), run("verify:gate", app="some-other-app")]
        alone, reason = cp.decide("pull_request", ["go/a.go"], ["docs/a.md"], other)
        self.assertFalse(alone)
        self.assertIn("no run", reason)

    def test_an_old_failure_does_not_hide_under_a_new_pass(self):
        checks = [run("verify:gate"), run("verify:gate", conclusion="failure")]
        self.assertFalse(cp.decide("pull_request", ["go/a.go"], ["docs/a.md"], checks)[0])

    def test_a_reference_push_runs(self):
        self.assertFalse(cp.decide("pull_request", ["go/a.go"], ["docs/reference/eval/x.jsonl"], GREEN)[0])


class ReadCheckRunsTest(unittest.TestCase):
    def test_reads_lines_and_skips_blank_ones(self):
        lines = [json.dumps(r) for r in GREEN] + [""]
        self.assertEqual(cp.read_check_runs(lines), GREEN)

    def test_an_absent_field_is_a_fault(self):
        with self.assertRaises(cp.FactError):
            cp.read_check_runs(['{"name": "verify:gate", "status": "completed", "app": "github-actions"}'])

    def test_a_line_that_is_not_an_object_is_a_fault(self):
        with self.assertRaises(cp.FactError):
            cp.read_check_runs(["[1]"])


def needs(documents_alone, results=None, paths_result="success"):
    value = {"paths": {"result": paths_result, "outputs": {"documents-alone": documents_alone}}}
    for job in cp.CODE_JOBS:
        value[job] = {"result": (results or {}).get(job, "success"), "outputs": {}}
    return value


class GateTest(unittest.TestCase):
    def test_every_job_green_passes(self):
        self.assertTrue(cp.gate(needs("false"))[0])

    def test_a_docs_only_skip_passes(self):
        skipped = {job: "skipped" for job in cp.CODE_JOBS}
        self.assertTrue(cp.gate(needs("true", skipped))[0])

    def test_a_skip_with_no_docs_decision_fails(self):
        self.assertFalse(cp.gate(needs("false", {"web": "skipped"}))[0])

    def test_a_failed_job_fails_on_a_docs_only_change_too(self):
        self.assertFalse(cp.gate(needs("true", {"go": "failure"}))[0])

    def test_a_failed_paths_job_fails(self):
        self.assertFalse(cp.gate(needs("", paths_result="failure"))[0])

    def test_an_absent_job_fails(self):
        value = needs("false")
        del value["docker"]
        self.assertFalse(cp.gate(value)[0])


class MainTest(unittest.TestCase):
    def setUp(self):
        quiet = contextlib.ExitStack()
        quiet.enter_context(contextlib.redirect_stdout(io.StringIO()))
        quiet.enter_context(contextlib.redirect_stderr(io.StringIO()))
        self.addCleanup(quiet.close)

    def test_decide_writes_the_output_line(self):
        with tempfile.TemporaryDirectory() as folder:
            paths = os.path.join(folder, "pr.txt")
            output = os.path.join(folder, "out")
            with open(paths, "w", encoding="utf-8") as handle:
                handle.write("docs/a.md\n\n")
            self.assertEqual(cp.main(["decide", "--event-name", "pull_request", "--pr-paths", paths,
                                      "--output", output]), 0)
            with open(output, encoding="utf-8") as handle:
                self.assertEqual(handle.read(), "documents-alone=true\n")

    def test_a_fault_exits_2_and_writes_no_output(self):
        with tempfile.TemporaryDirectory() as folder:
            output = os.path.join(folder, "out")
            self.assertEqual(cp.main(["decide", "--event-name", "pull_request", "--output", output]), 2)
            self.assertFalse(os.path.exists(output))

    def test_gate_exit_codes(self):
        self.assertEqual(cp.main(["gate", "--needs", json.dumps(needs("false"))]), 0)
        self.assertEqual(cp.main(["gate", "--needs", json.dumps(needs("false", {"go": "failure"}))]), 1)


if __name__ == "__main__":
    unittest.main()
