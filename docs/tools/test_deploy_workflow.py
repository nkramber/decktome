"""Tests of the guard of the fallback deploy workflow (REV-009, D-579).

Main alone reaches production. A hand start of the workflow from another
ref must skip every job, so each job carries the ref guard.

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import os
import re
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(os.path.dirname(HERE))
WORKFLOW = os.path.join(ROOT, ".github", "workflows", "deploy.yml")
GUARD = "github.ref == 'refs/heads/main'"


def jobs(text):
    """Return each job name with the text of its block."""
    body = ("\n" + text).split("\njobs:\n", 1)[1]
    out = {}
    name = None
    for line in body.splitlines():
        m = re.match(r"^  ([A-Za-z0-9_-]+):\s*$", line)
        if m:
            name = m.group(1)
            out[name] = []
            continue
        if re.match(r"^\S", line):
            break
        if name:
            out[name].append(line)
    return out


def unguarded(text):
    """Return the jobs whose own if line lacks the main guard."""
    bad = []
    for name, lines in jobs(text).items():
        ifs = [ln for ln in lines if re.match(r"^    if:", ln)]
        if not ifs or GUARD not in ifs[0]:
            bad.append(name)
    return bad


class DeployWorkflowTest(unittest.TestCase):
    def test_each_job_runs_on_main_alone(self):
        with open(WORKFLOW, encoding="utf-8") as f:
            text = f.read()
        self.assertTrue(jobs(text), "the workflow holds no job")
        self.assertEqual(unguarded(text), [])

    def test_a_job_with_no_guard_fails(self):
        text = "on:\n  workflow_dispatch:\n\njobs:\n  api:\n    runs-on: ubuntu-latest\n    steps: []\n"
        self.assertEqual(unguarded(text), ["api"])

    def test_a_guard_in_a_step_is_not_a_job_guard(self):
        text = (
            "jobs:\n  api:\n    runs-on: ubuntu-latest\n    steps:\n"
            "      - if: " + GUARD + "\n        run: true\n"
        )
        self.assertEqual(unguarded(text), ["api"])


class CloudBuildImagesTest(unittest.TestCase):
    """Each step of a Cloud Build file names its image by digest (REV-069, D-931)."""

    def test_each_step_image_names_a_digest(self):
        folder = os.path.join(ROOT, "cloudbuild")
        seen = 0
        for name in sorted(os.listdir(folder)):
            if not name.endswith(".yaml"):
                continue
            with open(os.path.join(folder, name), encoding="utf-8") as f:
                for line in f:
                    m = re.match(r"^\s+(?:- )?name:\s*(\S+)\s*$", line)
                    if m:
                        seen += 1
                        self.assertRegex(m.group(1), r"@sha256:[0-9a-f]{64}$", f"{name}: {m.group(1)}")
        self.assertGreater(seen, 0)


if __name__ == "__main__":
    unittest.main()
