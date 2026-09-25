"""Tests of the traffic check of the API deploy (REV-013).

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import importlib.util
import io
import json
import os
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(os.path.dirname(HERE))
spec = importlib.util.spec_from_file_location("traffic_check", os.path.join(HERE, "traffic_check.py"))
tc = importlib.util.module_from_spec(spec)
spec.loader.exec_module(tc)


def run(status):
    out = io.StringIO()
    code = tc.main(io.StringIO(json.dumps({"status": status})), out)
    return code, out.getvalue()


class TrafficCheckTest(unittest.TestCase):
    def test_the_new_revision_serves_all(self):
        code, _ = run({"latestReadyRevisionName": "mtg-api-00080", "traffic": [
            {"revisionName": "mtg-api-00080", "percent": 100, "latestRevision": True}]})
        self.assertEqual(code, 0)

    def test_a_pin_on_the_old_revision_fails(self):
        code, msg = run({"latestReadyRevisionName": "mtg-api-00080", "traffic": [
            {"revisionName": "mtg-api-00077", "percent": 100}]})
        self.assertEqual(code, 1)
        self.assertIn("0 percent", msg)

    def test_a_split_fails(self):
        code, _ = run({"latestReadyRevisionName": "mtg-api-00080", "traffic": [
            {"revisionName": "mtg-api-00080", "percent": 50},
            {"revisionName": "mtg-api-00077", "percent": 50}]})
        self.assertEqual(code, 1)

    def test_no_ready_revision_fails(self):
        code, _ = run({"traffic": []})
        self.assertEqual(code, 1)

    def test_the_deploy_runs_the_check(self):
        with open(os.path.join(ROOT, "cloudbuild", "api.yaml"), encoding="utf-8") as f:
            self.assertIn("docs/tools/traffic_check.py", f.read())


if __name__ == "__main__":
    unittest.main()
