"""Tests of ruleset_check.py (D-828)."""
import copy
import importlib.util
import json
import os
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("ruleset_check", os.path.join(HERE, "ruleset_check.py"))
rc = importlib.util.module_from_spec(spec)
spec.loader.exec_module(rc)

ROOT = os.path.abspath(os.path.join(HERE, "..", ".."))
with open(os.path.join(ROOT, rc.RULESET), encoding="utf-8") as handle:
    FILE = json.load(handle)
with open(os.path.join(ROOT, rc.SETTINGS), encoding="utf-8") as handle:
    SETTINGS = json.load(handle)


def live():
    """The file as GitHub gives it: other key order, extra fields, and new defaults."""
    body = copy.deepcopy(FILE)
    body.update({"id": rc.RULESET_ID, "source": "nkramber/decktome", "_links": {}})
    for rule in body["rules"]:
        if rule["type"] == "required_status_checks":
            rule["parameters"]["required_status_checks"].reverse()
        if rule["type"] == "pull_request":
            rule["parameters"]["required_reviewers"] = []
    return body


def rule(body, kind):
    return next(r for r in body["rules"] if r["type"] == kind)


class TheFile(unittest.TestCase):
    def contexts(self):
        return {c["context"] for c in rule(FILE, "required_status_checks")["parameters"]["required_status_checks"]}

    def test_the_file_requires_the_gate_the_contract_and_each_pull_request_job_of_verify(self):
        with open(os.path.join(ROOT, ".github/workflows/verify.yml"), encoding="utf-8") as handle:
            names = {line.split('"')[1] for line in handle if line.strip().startswith('name: "verify:')}
        # verify:changes runs on a dispatch alone, so each pull request reports it as skipped.
        self.assertEqual(self.contexts(), (names - {"verify:changes"}) | {"review-gate", "pr-contract"})

    def test_every_check_comes_from_the_actions_app(self):
        checks = rule(FILE, "required_status_checks")["parameters"]["required_status_checks"]
        self.assertEqual({c["integration_id"] for c in checks}, {15368})

    def test_no_bypass_and_squash_alone(self):
        self.assertEqual(FILE["bypass_actors"], [])
        params = rule(FILE, "pull_request")["parameters"]
        self.assertEqual(params["allowed_merge_methods"], ["squash"])
        self.assertTrue(params["required_review_thread_resolution"])
        self.assertEqual(SETTINGS, {"allow_auto_merge": True, "allow_squash_merge": True,
                                    "allow_rebase_merge": False, "allow_merge_commit": False})


class Compare(unittest.TestCase):
    def test_the_live_form_of_the_file_matches(self):
        self.assertEqual(rc.compare(FILE, live()), [])

    def test_a_missing_check_is_a_difference(self):
        body = live()
        rule(body, "required_status_checks")["parameters"]["required_status_checks"].pop()
        self.assertEqual(len(rc.compare(FILE, body)), 1)

    def test_an_extra_live_check_is_a_difference(self):
        body = live()
        rule(body, "required_status_checks")["parameters"]["required_status_checks"].append({"context": "x", "integration_id": 1})
        self.assertIn("the file does not", rc.compare(FILE, body)[0])

    def test_a_bypass_actor_is_a_difference(self):
        body = live()
        body["bypass_actors"] = [{"actor_id": 5, "actor_type": "RepositoryRole", "bypass_mode": "always"}]
        self.assertTrue(rc.compare(FILE, body))

    def test_an_extra_live_rule_is_a_difference(self):
        body = live()
        body["rules"].append({"type": "deletion"})
        self.assertIn("does not name", rc.compare(FILE, body)[0])

    def test_a_disabled_ruleset_is_a_difference(self):
        body = live()
        body["enforcement"] = "disabled"
        self.assertTrue(rc.compare(FILE, body))

    def test_a_rebase_merge_is_a_difference(self):
        repo = dict(SETTINGS, allow_rebase_merge=True, private=False)
        self.assertEqual(len(rc.compare(SETTINGS, repo, "repository")), 1)


if __name__ == "__main__":
    unittest.main()
