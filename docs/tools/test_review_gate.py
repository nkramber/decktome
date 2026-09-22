"""Tests of the review gate (D-803, D-804).

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
spec = importlib.util.spec_from_file_location("review_gate", os.path.join(HERE, "review_gate.py"))
rg = importlib.util.module_from_spec(spec)
spec.loader.exec_module(rg)

CODE = "a" * 40
META = "b" * 40

RECORD = """# Pull request 9 review

## Identity

- Pull request: 9
- Head: `{head}`

## Verdict

**{verdict}.** This verdict applies to head `{head}`.
"""


def facts(labels=(), files=("go/a.go",), commits=None, decisions=""):
    return {
        "number": 9,
        "labels": list(labels),
        "files": list(files),
        "commits": commits if commits is not None else [
            {"sha": CODE, "files": ["go/a.go"]},
            {"sha": META, "files": ["docs/reviews/pr-9.md", "docs/SESSION-HANDOFF.md"]},
        ],
        "decisionsDiff": decisions,
    }


def reader(text):
    return lambda path: text if path == "docs/reviews/pr-9.md" else None


def results(value, text):
    return {rule: (result, reason) for rule, result, reason in rg.check(value, reader(text))}


class EffectiveHeadTest(unittest.TestCase):
    def test_a_metadata_commit_does_not_move_the_head(self):
        self.assertEqual(rg.effective_head(facts()["commits"], 9)["sha"], CODE)

    def test_the_record_of_another_pull_request_moves_the_head(self):
        commits = [{"sha": CODE, "files": ["go/a.go"]}, {"sha": META, "files": ["docs/reviews/pr-8.md"]}]
        self.assertEqual(rg.effective_head(commits, 9)["sha"], META)

    def test_the_archive_and_the_response_are_metadata(self):
        commits = [{"sha": CODE, "files": ["go/a.go"]},
                   {"sha": META, "files": ["docs/reference/session-handoff-archive.md",
                                           "docs/reviews/pr-9-response.md"]}]
        self.assertEqual(rg.effective_head(commits, 9)["sha"], CODE)

    def test_metadata_alone_has_no_head(self):
        self.assertIsNone(rg.effective_head([{"sha": META, "files": ["docs/SESSION-HANDOFF.md"]}], 9))


class RecordTest(unittest.TestCase):
    def test_an_approved_record_on_the_effective_head_passes(self):
        got = results(facts(), RECORD.format(head=CODE[:7], verdict="Ready for owner merge"))
        self.assertEqual([got[r][0] for r in ("RG 1", "RG 2", "RG 3", "RG 4", "RG 5")], ["pass"] * 5)

    def test_no_record_is_a_fault(self):
        got = results(facts(), None)
        self.assertEqual((got["RG 3"][0], got["RG 4"][0], got["RG 5"][0]), ("fault", "skip", "skip"))

    def test_changes_required_is_a_fault(self):
        got = results(facts(), RECORD.format(head=CODE, verdict="Changes required"))
        self.assertEqual(got["RG 4"][0], "fault")

    def test_an_unknown_verdict_is_a_fault(self):
        got = results(facts(), RECORD.format(head=CODE, verdict="Looks fine"))
        self.assertIn("no verdict", got["RG 4"][1])

    def test_two_bold_names_are_a_fault(self):
        text = RECORD.format(head=CODE, verdict="Ready for owner merge") + "\n**Not Ready for owner merge**\n"
        self.assertEqual(results(facts(), text)["RG 4"][0], "fault")

    def test_a_bold_name_under_earlier_verdicts_does_not_count(self):
        text = RECORD.format(head=CODE, verdict="Ready for owner merge") + "\n## Earlier verdicts\n\n**Blocked.**\n"
        self.assertEqual(results(facts(), text)["RG 4"][0], "pass")

    def test_no_verdict_section_is_a_fault(self):
        self.assertEqual(results(facts(), "## Identity\n\n- Head: `aaaaaaa`\n")["RG 4"][0], "fault")

    def test_a_stale_head_is_a_fault(self):
        got = results(facts(), RECORD.format(head="c" * 7, verdict="Ready for owner merge"))
        self.assertEqual(got["RG 5"][0], "fault")

    def test_the_tip_is_not_the_effective_head(self):
        got = results(facts(), RECORD.format(head=META, verdict="Ready for owner merge"))
        self.assertEqual(got["RG 5"][0], "fault")

    def test_a_short_hash_is_a_fault(self):
        got = results(facts(), RECORD.format(head=CODE[:6], verdict="Ready for owner merge"))
        self.assertIn("shorter", got["RG 5"][1])

    def test_a_head_field_outside_identity_does_not_count(self):
        text = "## Identity\n\n- Pull request: 9\n\n## Notes\n\n- Head: `%s`\n\n## Verdict\n\n**Ready for owner merge.**\n" % CODE
        self.assertEqual(results(facts(), text)["RG 5"][0], "fault")

    def test_metadata_alone_is_a_fault(self):
        value = facts(commits=[{"sha": META, "files": ["docs/SESSION-HANDOFF.md"]}])
        got = results(value, RECORD.format(head=META, verdict="Ready for owner merge"))
        self.assertIn("no effective head", got["RG 5"][1])


class LabelTest(unittest.TestCase):
    def test_the_label_covers_a_docs_pull_request(self):
        got = results(facts(labels=["review-override"], files=["docs/a.md", "CLAUDE.md"]), None)
        self.assertEqual([got[r][0] for r in ("RG 1", "RG 2", "RG 3", "RG 4", "RG 5")],
                         ["pass", "pass", "skip", "skip", "skip"])

    def test_the_label_never_covers_code_tools_hooks_or_workflows(self):
        for path in ("go/a.go", "docs/tools/pr_check.py", ".claude/hooks/session_bind.py",
                     ".claude/settings.json", ".github/workflows/verify.yml"):
            got = results(facts(labels=["review-override"], files=["docs/a.md", path]), None)
            self.assertEqual(got["RG 1"][0], "fault", path)
            self.assertEqual(got["RG 3"][0], "fault", path)

    def test_the_label_never_covers_a_decision_row(self):
        diff = "@@ -1 +1 @@\n+| D-900 | 2026-09-22 | q | a |\n"
        got = results(facts(labels=["review-override"], files=["docs/decisions.md"], decisions=diff), None)
        self.assertEqual(got["RG 2"][0], "fault")

    def test_prose_of_the_register_is_no_decision_row(self):
        diff = "@@ -1 +1 @@\n+Note on the numbers.\n"
        got = results(facts(labels=["review-override"], files=["docs/decisions.md"], decisions=diff), None)
        self.assertEqual(got["RG 2"][0], "pass")


class ReadTest(unittest.TestCase):
    def test_modes(self):
        self.assertEqual(rg.read_mode("enforced\n"), "enforced")
        self.assertEqual(rg.read_mode("advisory"), "advisory")
        for text in (None, "", "off"):
            with self.assertRaises(rg.FactError):
                rg.read_mode(text)

    def test_facts_need_every_field(self):
        value = facts()
        del value["labels"]
        with self.assertRaises(rg.FactError):
            rg.read_facts(json.dumps(value))

    def test_a_bool_is_no_number(self):
        value = facts()
        value["number"] = True
        with self.assertRaises(rg.FactError):
            rg.read_facts(json.dumps(value))


class MainTest(unittest.TestCase):
    def setUp(self):
        quiet = contextlib.ExitStack()
        quiet.enter_context(contextlib.redirect_stdout(io.StringIO()))
        quiet.enter_context(contextlib.redirect_stderr(io.StringIO()))
        self.addCleanup(quiet.close)
        folder = tempfile.TemporaryDirectory()
        self.addCleanup(folder.cleanup)
        self.root = folder.name
        self.head = os.path.join(self.root, "head")
        os.makedirs(os.path.join(self.head, "docs", "reviews"))
        self.facts = os.path.join(self.root, "facts.json")
        with open(self.facts, "w", encoding="utf-8") as handle:
            json.dump(facts(), handle)

    def write(self, name, text):
        path = os.path.join(self.root, name)
        with open(path, "w", encoding="utf-8") as handle:
            handle.write(text)
        return path

    def run_gate(self, mode):
        mode_file = self.write("mode", mode)
        return rg.main(["--facts", self.facts, "--head-files", self.head, "--mode-file", mode_file])

    def test_enforced_mode_fails_a_fault(self):
        self.assertEqual(self.run_gate("enforced"), 1)

    def test_advisory_mode_passes_a_fault(self):
        self.assertEqual(self.run_gate("advisory"), 0)

    def test_an_unknown_mode_stops(self):
        self.assertEqual(self.run_gate("maybe"), 2)

    def test_an_approved_record_passes_enforced_mode(self):
        path = os.path.join(self.head, "docs", "reviews", "pr-9.md")
        with open(path, "w", encoding="utf-8") as handle:
            handle.write(RECORD.format(head=CODE, verdict="Ready for owner merge"))
        self.assertEqual(self.run_gate("enforced"), 0)

    def test_a_symlink_out_of_the_head_is_no_record(self):
        outside = self.write("outside.md", RECORD.format(head=CODE, verdict="Ready for owner merge"))
        os.symlink(outside, os.path.join(self.head, "docs", "reviews", "pr-9.md"))
        self.assertEqual(self.run_gate("enforced"), 1)


if __name__ == "__main__":
    unittest.main()
