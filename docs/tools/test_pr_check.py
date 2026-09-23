"""Tests of the one-pr-one-session contract check (D-747, D-748).

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import importlib.util
import os
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("pr_check", os.path.join(HERE, "pr_check.py"))
pc = importlib.util.module_from_spec(spec)
spec.loader.exec_module(pc)

SESSION = """## Session

- Role: author
- Branch: `pr60-example`
- Base: `372d912`
"""

ROWS = {
    "hand-off": "Changed: `docs/SESSION-HANDOFF.md` records the completed state and the next step",
    "decisions": "Changed: `docs/decisions.md` adds the owner answer that this change carries",
    "roadmap": "Changed: `docs/design-roadmap.md` marks the item merged as this pull request",
    "questions": "Reviewed; no change needed: `docs/open-questions.md` holds no question that this change answers or opens",
    "root guidance": "Reviewed; no change needed: `CLAUDE.md` names no command, rule, or file that this change moves",
    "skills and hooks": "Reviewed; no change needed: `.claude/skills/mtg-corpus/SKILL.md` states no rule that the theme code changes",
    "operations": "Reviewed; no change needed: `docs/deploy-and-rollback.md` reads the same deploy triggers after this change",
    "reference documents": "Changed: `docs/reference/pr60-example-2026-09-17.md` holds every count of the free replay",
}

CHANGED = [
    "go/internal/themes/match.go",
    "go/internal/themes/match_test.go",
    "docs/SESSION-HANDOFF.md",
    "docs/decisions.md",
    "docs/design-roadmap.md",
    "docs/reference/pr60-example-2026-09-17.md",
]


def body(rows=None, extra=""):
    rows = dict(ROWS, **(rows or {}))
    table = "\n".join(f"| {k} | {v} |" for k, v in rows.items() if v is not None)
    return f"Summary.\n\n{SESSION}\n## Documentation impact\n\n| Category | Entry |\n|---|---|\n{table}\n\n{extra}"


def exists(path):
    return True


def check(text, changed=CHANGED, title="PR-60: a theme word reads its plural", author="nkramber", handoff_added=""):
    return pc.check_pr(title, text, author, "pr60-example", changed, exists, lambda sha: True, handoff_added)


class PullRequestContract(unittest.TestCase):
    def test_implementation_with_every_document_passes(self):
        self.assertEqual(check(body()), [])

    def test_specific_no_change_reason_passes(self):
        self.assertEqual(check(body({"questions": "Reviewed; no change needed: `docs/owner-questions.md` holds no question that the theme match answers"})), [])

    def test_missing_handoff_change_fails(self):
        changed = [p for p in CHANGED if p != "docs/SESSION-HANDOFF.md"]
        errors = check(body(), changed=changed)
        self.assertTrue(any("does not change docs/SESSION-HANDOFF.md" in e for e in errors), errors)

    def test_handoff_marked_unchanged_fails(self):
        errors = check(body({"hand-off": "Reviewed; no change needed: `docs/SESSION-HANDOFF.md` still reads the right next step"}))
        self.assertTrue(any("every pull request changes" in e for e in errors), errors)

    def test_missing_category_fails(self):
        errors = check(body({"operations": None}))
        self.assertTrue(any("no row for the category 'operations'" in e for e in errors), errors)

    def test_generic_reason_fails(self):
        errors = check(body({"questions": "Reviewed; no change needed: `docs/open-questions.md` no documentation impact"}))
        self.assertTrue(any("generic" in e for e in errors), errors)

    def test_short_reason_fails(self):
        errors = check(body({"questions": "Not applicable: none"}))
        self.assertTrue(any("generic" in e or "words" in e for e in errors), errors)

    def test_placeholder_fails(self):
        errors = check(body({"questions": "Reviewed; no change needed: `docs/open-questions.md` <reason and path> TODO later today"}))
        self.assertTrue(any("placeholder" in e for e in errors), errors)

    def test_pending_status_fails(self):
        errors = check(body({"roadmap": "Pending: a clean author session completes this row (D-747)"}))
        self.assertTrue(any("starts with none of" in e for e in errors), errors)

    def test_unchanged_claim_on_a_changed_document_fails(self):
        errors = check(body({"decisions": "Reviewed; no change needed: `docs/decisions.md` holds every answer this change needs"}))
        self.assertTrue(any("the diff changes docs/decisions.md" in e for e in errors), errors)

    def test_changed_claim_with_no_change_fails(self):
        errors = check(body({"questions": "Changed: `docs/open-questions.md` closes the question that this change answers"}))
        self.assertTrue(any("changes no file of this category" in e for e in errors), errors)

    def test_deferred_documentation_in_the_body_fails(self):
        errors = check(body(extra="The roadmap will update after merge in a follow-up PR."))
        self.assertTrue(any("defers documentation" in e for e in errors), errors)

    def test_negated_follow_up_passes(self):
        self.assertEqual(check(body(extra="No follow-up pull request records the documents.")), [])

    def test_deferred_documentation_in_the_handoff_fails(self):
        errors = check(body(), handoff_added="A separate PR updates the roadmap after the merge.")
        self.assertTrue(any("hand-off defers" in e for e in errors), errors)

    def test_merge_record_pull_request_fails(self):
        changed = ["docs/SESSION-HANDOFF.md", "docs/design-roadmap.md", "CLAUDE.md"]
        rows = {
            "decisions": "Reviewed; no change needed: `docs/decisions.md` holds no new decision for this record",
            "roadmap": "Changed: `docs/design-roadmap.md` marks the merge of the earlier pull request",
            "root guidance": "Changed: `CLAUDE.md` reads the merge of the earlier pull request",
            "reference documents": "Reviewed; no change needed: `docs/reference/` holds no gate document of this record",
        }
        errors = check(body(rows), changed=changed, title="The documents read the merge and the deploy of #182")
        self.assertTrue(any("records an earlier merge" in e for e in errors), errors)

    def test_every_past_merge_record_title_matches(self):
        titles = [
            "The documents read the merge and the deploy of #179 (PR-54) before a context reset",
            "The documents read the merge of #177 (M-17) before a context reset",
            "The documents read the merge and the deploy of #173 (PR-51)",
            "The documents read the merge and the deploy of #169 before a context wipe (PR-45b)",
            "The documents read the merge and the deploy of #167 before a context wipe (PR-49)",
            "The documents read the merge and the deploy of #158 (F-122, PR-25)",
            "The documents read the state before a context wipe (D-685 to D-689)",
            "The documents read the PR-44 deploy (F-120, PR-44)",
            "The documents read the PR-43 deploy and the meta run of 2026-09-12 (F-119, PR-43)",
            "The documents read the state before a context reset (F-109, PR-41)",
            "The hand-off and the roadmap read the deployed fix, the bracket 5 session, and M-11 underway (D-673)",
            "The hand-off and the roadmap read the merges of #134, #135, and the dependabot bumps",
        ]
        for title in titles:
            self.assertRegex(title, pc.MERGE_RECORD_TITLE)
        self.assertNotRegex("PR-53: a finisher takes the wincon role", pc.MERGE_RECORD_TITLE)

    def test_session_block_is_required(self):
        errors = check(body().replace(SESSION, ""))
        self.assertTrue(any("no '## Session' section" in e for e in errors), errors)

    def test_session_branch_must_match_the_head(self):
        errors = check(body().replace("`pr60-example`", "`pr61-other`"))
        self.assertTrue(any("pull request head is pr60-example" in e for e in errors), errors)

    def test_base_must_be_an_ancestor(self):
        errors = pc.check_pr("PR-60", body(), "nkramber", "pr60-example", CHANGED, exists, lambda sha: False)
        self.assertTrue(any("not an ancestor" in e for e in errors), errors)

    def test_reviewer_role_passes(self):
        self.assertEqual(check(body().replace("Role: author", "Role: reviewer")), [])

    def test_a_roadmap_change_without_the_merge_mark_fails(self):
        errors = pc.check_pr("PR-63: a gate", body(), "nkramber", "pr60-example", CHANGED, exists,
                             lambda sha: True, number=212, roadmap_added="**PR-63: A gate.** 🔧 built.")
        self.assertTrue(any("merged as #212" in e for e in errors), errors)

    def test_a_roadmap_change_with_the_merge_mark_passes(self):
        errors = pc.check_pr("PR-63: a gate", body(), "nkramber", "pr60-example", CHANGED, exists,
                             lambda sha: True, number=212, roadmap_added="**PR-63: A gate.** ✅ merged as #212.")
        self.assertFalse(any("merged as" in e for e in errors), errors)

    def test_the_mark_of_a_longer_number_fails(self):
        for text in ("✅ merged as #2120.", "✅ merged as #2120", "✅ merged as #21"):
            with self.subTest(text=text):
                errors = pc.check_pr("PR-63: a gate", body(), "nkramber", "pr60-example", CHANGED, exists,
                                     lambda sha: True, number=212, roadmap_added=text)
                self.assertTrue(any("merged as #212" in e for e in errors), errors)

    def test_the_mark_at_the_end_of_the_text_passes(self):
        errors = pc.check_pr("PR-63: a gate", body(), "nkramber", "pr60-example", CHANGED, exists,
                             lambda sha: True, number=212, roadmap_added="✅ merged as #212")
        self.assertFalse(any("merged as" in e for e in errors), errors)

    def test_a_draft_with_no_number_skips_the_merge_mark(self):
        errors = pc.check_pr("PR-63: a gate", body(), "nkramber", "pr60-example", CHANGED, exists,
                             lambda sha: True, number=None, roadmap_added="🔧 built")
        self.assertFalse(any("merged as" in e for e in errors), errors)

    def test_dependabot_is_exempt(self):
        self.assertEqual(check("", changed=["web/package.json"], author="dependabot[bot]"), [])


class RepositoryWiring(unittest.TestCase):
    def test_the_repository_passes_the_skill_check(self):
        def read(path):
            try:
                with open(os.path.join(pc.ROOT, path), encoding="utf-8") as handle:
                    return handle.read()
            except OSError:
                return None

        def listdir(path):
            full = os.path.join(pc.ROOT, path)
            return [d for d in os.listdir(full) if os.path.isdir(os.path.join(full, d))]

        self.assertEqual(pc.check_skills(read, listdir), [])

    def test_a_skill_without_the_block_text_fails(self):
        files = {".claude/skills/one-pr-one-session/SKILL.md": "---\nname: one-pr-one-session\ndescription: x\n---\nbody"}
        errors = pc.check_skills(files.get, lambda path: ["one-pr-one-session"])
        self.assertTrue(any("exact text" in e for e in errors), errors)
        self.assertTrue(any("CLAUDE.md does not require" in e for e in errors), errors)


if __name__ == "__main__":
    unittest.main()
