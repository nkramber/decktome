"""Tests of the reference check (D-753).

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import importlib.util
import os
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("ref_check", os.path.join(HERE, "ref_check.py"))
rc = importlib.util.module_from_spec(spec)
spec.loader.exec_module(rc)

DECISIONS = "| # | Date |\n| D-1 | a |\n| D-2 (amended by D-3) | b |\n| D-3 | c |\n"
ROADMAP = "| F-1 | a finding |\n\n**M-2: A metric.** The count.\n\n**PR-28a: A change.** ✅ merged\n"

PATHS = {"docs/decisions.md", "docs/design-roadmap.md", "go/internal/decks/store.go",
         ".claude/skills/one-pr-one-session/SKILL.md"}
FOLDERS = {"docs", "go", "go/internal", "go/internal/decks", ".claude", ".claude/skills"}
TOP = {"docs", "go", ".claude"}


def run(text, doc="docs/note.md"):
    known = rc.registers(DECISIONS, ROADMAP)
    return rc.check(doc, text, known, TOP, PATHS, FOLDERS)


class RefCheckTest(unittest.TestCase):
    def test_a_defined_id_passes(self):
        self.assertEqual(run("The rule of D-2 and F-1 and M-2 and PR-28a."), [])

    def test_an_undefined_id_fails(self):
        findings = run("The rule of D-9.")
        self.assertEqual([(1, "REF 1", "no register defines D-9")], findings)

    def test_a_family_id_resolves_against_a_lettered_entry(self):
        self.assertEqual(run("PR-28 splits into three."), [])

    def test_a_retired_number_passes(self):
        self.assertEqual(run("The branch held D-172 to D-176."), [])

    def test_a_dead_path_fails(self):
        findings = run("Read `go/internal/store` for the schema.")
        self.assertEqual(1, len(findings))
        self.assertEqual("REF 2", findings[0][1])

    def test_a_live_path_passes(self):
        self.assertEqual(run("Read `go/internal/decks/store.go` and `go/internal`."), [])

    def test_a_path_resolves_from_the_folder_of_the_document(self):
        self.assertEqual(run("Read `decisions.md`.", doc="docs/note.md"), [])

    def test_a_bare_name_takes_no_rule(self):
        self.assertEqual(run("The file `themes.json` holds the rows."), [])

    def test_a_dotted_path_reads_the_rule(self):
        findings = run("Load `.claude/skills/absent/SKILL.md`.")
        self.assertEqual(1, len(findings))
        self.assertEqual("REF 2", findings[0][1])

    def test_a_live_dotted_path_passes(self):
        self.assertEqual(run("Load `.claude/skills/one-pr-one-session/SKILL.md`."), [])

    def test_a_parent_path_takes_no_rule(self):
        self.assertEqual(run("The folder `../docs/decisions.md` sits above."), [])

    def test_a_placeholder_takes_no_rule(self):
        self.assertEqual(run("The build lands at `go/X`."), [])

    def test_a_path_outside_the_repository_takes_no_rule(self):
        self.assertEqual(run("Read `backend/spellbook/models/variant.py`."), [])

    def test_a_scratch_path_takes_no_rule(self):
        self.assertEqual(run("The fit wrote `go/.local/tune/run18.json`."), [])

    def test_a_runtime_folder_takes_no_rule(self):
        self.assertEqual(run("The harvest writes `docs/reference/feedback/`."), [])

    def test_a_fenced_block_takes_no_rule(self):
        self.assertEqual(run("Run this:\n\n```\nD-9 `go/internal/store`\n```\n"), [])

    def test_the_line_number_reads_the_source(self):
        findings = run("A line.\n\nThe rule of D-9.\n")
        self.assertEqual(3, findings[0][0])

    def test_a_dated_record_is_exempt(self):
        self.assertTrue(rc.DATED.search("docs/reference/note-2026-09-17.md"))
        self.assertTrue(rc.DATED.search("docs/reference/session-handoff-archive.md"))
        self.assertIsNone(rc.DATED.search("docs/decisions.md"))



class RepoPathsTest(unittest.TestCase):
    def test_the_git_file_of_a_worktree_is_no_top_level_entry(self):
        with tempfile.TemporaryDirectory() as root:
            with open(os.path.join(root, ".git"), "w", encoding="utf-8") as handle:
                handle.write("gitdir: /elsewhere/.git/worktrees/x\n")
            os.makedirs(os.path.join(root, "docs"))
            with open(os.path.join(root, "docs", "a.md"), "w", encoding="utf-8") as handle:
                handle.write("a\n")
            saved, rc.ROOT = rc.ROOT, root
            try:
                paths, folders = rc.repo_paths()
            finally:
                rc.ROOT = saved
        self.assertEqual(paths, {"docs/a.md"})
        self.assertEqual(folders, {"docs"})


if __name__ == "__main__":
    unittest.main()
