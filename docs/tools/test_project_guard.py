"""Tests of the PROJECT_ID guard of the targets of the deployed project (REV-090).

A shell that works on more than one project exports PROJECT_ID for
another one. `make allow`, `make disallow`, and `make mark-verified` must
refuse that value before they reach any project. Each test runs a target
that the guard stops, so no test touches a project.

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import os
import subprocess
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(os.path.dirname(HERE))


def run(target, env_project):
    env = dict(os.environ, PROJECT_ID=env_project)
    return subprocess.run(["make", "-s", target, "EMAIL=ann@example.com"], cwd=ROOT, env=env,
                          capture_output=True, text=True, timeout=60)


class ProjectGuardTest(unittest.TestCase):
    def test_an_exported_project_is_refused(self):
        for target in ("allow", "disallow", "mark-verified"):
            with self.subTest(target=target):
                res = run(target, "some-other-project")
                self.assertNotEqual(res.returncode, 0)
                self.assertIn("on the command line", res.stdout + res.stderr)


if __name__ == "__main__":
    unittest.main()
