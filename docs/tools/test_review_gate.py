"""Tests of review_gate.py (D-810 to D-817, D-837)."""
import importlib.util
import json
import os
import subprocess
import sys
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("review_gate", os.path.join(HERE, "review_gate.py"))
rg = importlib.util.module_from_spec(spec)
spec.loader.exec_module(rg)

N = 212
HEAD = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678"
CODE = [(HEAD, ["go/internal/generate/manapass.go"])]
SESSION = "session@example.com"


def record(head=HEAD[:7], verdict=f"**{rg.APPROVED}.** This verdict applies to head `x`.", identity=True):
    lines = ["# PR-212 review", "", "Date: 2026-09-23", ""]
    if identity:
        lines += ["## Identity", "", "- PR: 212", f"- Head: `{head}`", ""]
    lines += ["## Findings", "", "No finding.", "", "## Verdict", "", verdict, ""]
    return "\n".join(lines)


def run(labels=(), files=("go/internal/generate/manapass.go",), commits=CODE, authors=(SESSION,), text=None, author="nkramber"):
    reader = lambda path: text if path == rg.record_path(N) else None
    results = rg.evaluate(N, author, set(labels), list(files), commits, set(authors), reader)
    return [(rule, state) for rule, state, _ in results], results


def faults(results):
    return [r for r in results if r[1] == "FAULT"]


class ReviewRecord(unittest.TestCase):
    def test_an_approved_record_of_the_effective_head_passes(self):
        states, results = run(text=record())
        self.assertEqual(faults(states), [], results)

    def test_no_record_fails(self):
        states, _ = run()
        self.assertIn(("RG 3", "FAULT"), states)

    def test_changes_required_fails(self):
        states, _ = run(text=record(verdict="**Changes required.** Two findings stay open."))
        self.assertIn(("RG 4", "FAULT"), states)

    def test_a_bold_negation_fails(self):
        states, _ = run(text=record(verdict=f"**Not {rg.APPROVED}.**"))
        self.assertIn(("RG 4", "FAULT"), states)

    def test_two_verdicts_fail(self):
        states, _ = run(text=record(verdict=f"**Changes required.** then **{rg.APPROVED}.**"))
        self.assertIn(("RG 4", "FAULT"), states)

    def test_an_earlier_verdict_in_its_own_section_passes(self):
        text = record().replace("## Verdict", "## Earlier verdicts\n\n**Changes required.** Fixed in abc1234.\n\n## Verdict")
        states, results = run(text=text)
        self.assertEqual(faults(states), [], results)

    def test_a_verdict_heading_with_more_words_is_not_the_section(self):
        text = record().replace("## Verdict", "## Verdict history")
        states, _ = run(text=text)
        self.assertIn(("RG 4", "FAULT"), states)

    def test_a_stale_head_fails(self):
        states, _ = run(text=record(head="0000000"))
        self.assertIn(("RG 5", "FAULT"), states)

    def test_a_short_head_fails(self):
        states, _ = run(text=record(head=HEAD[:6]))
        self.assertIn(("RG 5", "FAULT"), states)

    def test_a_head_outside_the_identity_list_fails(self):
        text = record(identity=False) + f"\n## Notes\n\n- Head: `{HEAD}`\n"
        states, _ = run(text=text)
        self.assertIn(("RG 5", "FAULT"), states)

    def test_a_metadata_commit_keeps_the_effective_head(self):
        commits = CODE + [("f" * 40, [rg.record_path(N), "docs/SESSION-HANDOFF.md"])]
        states, results = run(commits=commits, text=record())
        self.assertEqual(faults(states), [], results)

    def test_a_commit_of_another_record_moves_the_effective_head_and_keeps_the_gate(self):
        commits = CODE + [("f" * 40, ["docs/reviews/pr-211.md"])]
        self.assertEqual(rg.effective_head(commits, N), "f" * 40)
        states, results = run(commits=commits, text=record())
        self.assertEqual(faults(states), [], results)

    def test_metadata_commits_alone_have_no_effective_head(self):
        states, _ = run(commits=[(HEAD, ["docs/SESSION-HANDOFF.md"])], text=record())
        self.assertIn(("RG 5", "FAULT"), states)


class DocumentsAfterTheApproval(unittest.TestCase):
    """A commit of documents alone keeps a green gate green (D-837)."""

    def test_a_commit_of_each_kind_of_document_keeps_the_gate(self):
        for path in ["docs/design-roadmap.md", "docs/decisions.md", ".claude/skills/pr-review/SKILL.md",
                     "CLAUDE.md", "AGENTS.md", "README.md", ".github/pull_request_template.md",
                     "docs/reference/eval/baselines.json"]:
            with self.subTest(path=path):
                states, results = run(commits=CODE + [("d" * 40, [path])], text=record())
                self.assertEqual(faults(states), [], results)
                self.assertIn("D-837", results[-1][2])

    def test_many_commits_of_documents_keep_the_gate(self):
        commits = CODE + [("d" * 40, ["docs/decisions.md"]), ("e" * 40, ["docs/SESSION-HANDOFF.md"]),
                          ("f" * 40, ["docs/design-roadmap.md", ".claude/skills/gitar-review/SKILL.md"])]
        states, results = run(commits=commits, text=record())
        self.assertEqual(faults(states), [], results)

    def test_a_code_commit_after_a_commit_of_documents_fails(self):
        commits = CODE + [("d" * 40, ["docs/decisions.md"]), ("e" * 40, ["go/internal/generate/fill.go"])]
        states, _ = run(commits=commits, text=record())
        self.assertIn(("RG 5", "FAULT"), states)

    def test_a_commit_that_mixes_a_document_and_code_fails(self):
        commits = CODE + [("d" * 40, ["docs/decisions.md", "Makefile"])]
        states, _ = run(commits=commits, text=record())
        self.assertIn(("RG 5", "FAULT"), states)

    def test_a_refused_path_after_the_approval_fails(self):
        for path in ["docs/tools/review_gate.py", ".claude/hooks/session_bind.py", ".claude/settings.json",
                     ".claude/settings.local.json", ".github/workflows/review-gate.yml", "docsx/a.md"]:
            with self.subTest(path=path):
                states, _ = run(commits=CODE + [("d" * 40, [path])], text=record())
                self.assertIn(("RG 5", "FAULT"), states)

    def test_a_merge_commit_after_the_approval_fails(self):
        states, _ = run(commits=CODE + [("d" * 40, [rg.MERGE])], text=record())
        self.assertIn(("RG 5", "FAULT"), states)

    def test_a_head_outside_the_branch_fails_even_before_documents(self):
        commits = CODE + [("d" * 40, ["docs/decisions.md"])]
        states, _ = run(commits=commits, text=record(head="0000000"))
        self.assertIn(("RG 5", "FAULT"), states)

    def test_a_documentation_pull_request_keeps_its_approval(self):
        first = "c" * 40
        commits = [(first, ["docs/decisions.md"]), ("d" * 40, ["docs/design-roadmap.md"])]
        states, results = run(files=["docs/decisions.md", "docs/design-roadmap.md"], commits=commits,
                              text=record(head=first[:7]))
        self.assertEqual(faults(states), [], results)

    def test_the_check_without_commits_needs_the_effective_head_itself(self):
        state, _ = rg.check_head(rg.record_path(N), record(), "d" * 40)
        self.assertEqual(state, "FAULT")


class OverrideLabel(unittest.TestCase):
    def test_a_documentation_pull_request_with_the_label_passes(self):
        files = ["docs/decisions.md", "docs/SESSION-HANDOFF.md", ".claude/skills/pr-review/SKILL.md", "CLAUDE.md"]
        states, results = run(labels=[rg.LABEL], files=files)
        self.assertEqual(states, [("RG 1", "PASS")], results)

    def test_the_label_does_not_pass_code(self):
        states, _ = run(labels=[rg.LABEL])
        self.assertIn(("RG 1", "FAULT"), states)
        self.assertIn(("RG 3", "FAULT"), states)

    def test_the_label_does_not_pass_a_refused_path(self):
        for path in ["docs/tools/review_gate.py", ".claude/hooks/session_bind.py", ".claude/settings.json",
                     ".claude/settings.local.json", ".github/workflows/review-gate.yml", "Makefile", "docsx/a.md"]:
            with self.subTest(path=path):
                states, _ = run(labels=[rg.LABEL], files=["docs/decisions.md", path])
                self.assertIn(("RG 1", "FAULT"), states)

    def test_no_label_needs_the_record_even_for_documents(self):
        states, _ = run(files=["docs/decisions.md"])
        self.assertIn(("RG 3", "FAULT"), states)


class Dependabot(unittest.TestCase):
    def test_a_pull_request_of_dependabot_alone_passes(self):
        states, results = run(author=rg.DEPENDABOT, authors=[rg.DEPENDABOT_EMAIL])
        self.assertEqual(faults(states), [], results)

    def test_a_session_commit_on_a_dependabot_branch_needs_the_record(self):
        states, _ = run(author=rg.DEPENDABOT, authors=[rg.DEPENDABOT_EMAIL, SESSION])
        self.assertIn(("RG 3", "FAULT"), states)

    def test_the_dependabot_email_alone_does_not_exempt_another_author(self):
        states, _ = run(authors=[rg.DEPENDABOT_EMAIL])
        self.assertIn(("RG 3", "FAULT"), states)


class GitFacts(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.repo = self.tmp.name
        self.git("init", "-q", "-b", "main")
        self.commit({"README.md": "base\n"}, "base")
        self.base = self.git("rev-parse", "HEAD")
        self.git("checkout", "-q", "-b", "work")

    def tearDown(self):
        self.tmp.cleanup()

    def git(self, *args, email=SESSION):
        env = {**os.environ, "GIT_AUTHOR_NAME": "a", "GIT_AUTHOR_EMAIL": email,
               "GIT_COMMITTER_NAME": "a", "GIT_COMMITTER_EMAIL": email}
        out = subprocess.run(["git", *args], cwd=self.repo, capture_output=True, text=True, env=env, check=True)
        return out.stdout.strip()

    def commit(self, files, message, email=SESSION):
        for path, text in files.items():
            full = os.path.join(self.repo, path)
            os.makedirs(os.path.dirname(full), exist_ok=True)
            with open(full, "w", encoding="utf-8") as handle:
                handle.write(text)
            self.git("add", path)
        self.git("commit", "-q", "-m", message, email=email)
        return self.git("rev-parse", "HEAD")

    def gate(self, labels=(), author="nkramber"):
        event = os.path.join(self.repo, "..", f"event-{os.path.basename(self.repo)}.json")
        with open(event, "w", encoding="utf-8") as handle:
            json.dump({"pull_request": {"number": N, "user": {"login": author}, "base": {"sha": self.base},
                                        "labels": [{"name": name} for name in labels]}}, handle)
        out = subprocess.run([sys.executable, os.path.join(HERE, "review_gate.py"), "--event", event,
                              "--head", "HEAD", "--repo", self.repo], capture_output=True, text=True)
        os.remove(event)
        return out.returncode, out.stdout

    def test_the_command_passes_an_approved_head_and_ignores_the_review_commit(self):
        code = self.commit({"go/a.go": "package a\n"}, "code")
        self.commit({rg.record_path(N): record(head=code[:10])}, "review")
        status, out = self.gate()
        self.assertEqual(status, 0, out)
        self.assertIn(f"effective head `{code}`", out)

    def test_the_command_fails_a_code_push_after_the_review(self):
        code = self.commit({"go/a.go": "package a\n"}, "code")
        self.commit({rg.record_path(N): record(head=code)}, "review")
        self.commit({"go/b.go": "package a\n"}, "more code")
        status, out = self.gate()
        self.assertEqual(status, 1, out)
        self.assertIn("RG 5: FAULT", out)

    def test_the_command_passes_a_commit_of_documents_after_the_review(self):
        code = self.commit({"go/a.go": "package a\n"}, "code")
        self.commit({rg.record_path(N): record(head=code)}, "review")
        self.commit({"docs/design-roadmap.md": "mark\n", "docs/decisions.md": "row\n"}, "documents")
        status, out = self.gate()
        self.assertEqual(status, 0, out)
        self.assertIn("RG 5: PASS", out)
        self.assertIn("D-837", out)

    def test_the_command_fails_code_after_a_commit_of_documents(self):
        code = self.commit({"go/a.go": "package a\n"}, "code")
        self.commit({rg.record_path(N): record(head=code)}, "review")
        self.commit({"docs/decisions.md": "row\n"}, "documents")
        self.commit({"docs/tools/new.py": "print(1)\n"}, "a tool")
        status, out = self.gate()
        self.assertEqual(status, 1, out)
        self.assertIn("RG 5: FAULT", out)

    def test_a_rename_into_docs_counts_as_a_change_of_code(self):
        self.commit({"go/a.go": "package a\n"}, "code")
        self.base = self.git("rev-parse", "HEAD")
        os.makedirs(os.path.join(self.repo, "docs"))
        self.git("mv", "go/a.go", "docs/a.go")
        self.git("commit", "-q", "-m", "move")
        status, out = self.gate(labels=[rg.LABEL])
        self.assertEqual(status, 1, out)
        self.assertIn("`go/a.go`", out)

    def test_a_merge_of_main_after_the_review_moves_the_effective_head(self):
        code = self.commit({"go/a.go": "package a\n"}, "code")
        self.commit({rg.record_path(N): record(head=code)}, "review")
        self.git("checkout", "-q", "main")
        self.commit({"go/c.go": "package c\n"}, "main moves")
        self.git("checkout", "-q", "work")
        self.git("merge", "-q", "--no-edit", "main")
        merge = self.git("rev-parse", "HEAD")
        status, out = self.gate()
        self.assertEqual(status, 1, out)
        self.assertIn(f"effective head is `{merge}`", out)

    def test_the_effective_head_mode_prints_the_rule_of_the_check(self):
        code = self.commit({"go/a.go": "package a\n"}, "code")
        self.commit({"docs/SESSION-HANDOFF.md": "state\n"}, "hand-off")
        out = subprocess.run([sys.executable, os.path.join(HERE, "review_gate.py"), "--effective-head", str(N),
                              "--base", self.base, "--repo", self.repo], capture_output=True, text=True)
        self.assertEqual((out.returncode, out.stdout.strip()), (0, code), out.stderr)

    def test_the_label_reads_every_commit_and_not_the_last_one(self):
        self.commit({"go/a.go": "package a\n"}, "code")
        self.commit({"docs/a.md": "notes\n"}, "docs")
        status, out = self.gate(labels=[rg.LABEL])
        self.assertEqual(status, 1, out)
        self.assertIn("`go/a.go`", out)
        self.assertIn("2 changed path(s)", out)

    def test_the_label_passes_documents(self):
        self.commit({"docs/decisions.md": "| D-1 |\n"}, "docs")
        status, out = self.gate(labels=[rg.LABEL])
        self.assertEqual(status, 0, out)

    def test_dependabot_reads_the_commit_authors(self):
        self.commit({"go/go.mod": "module a\n"}, "bump", email=rg.DEPENDABOT_EMAIL)
        status, out = self.gate(author=rg.DEPENDABOT)
        self.assertEqual(status, 0, out)
        self.commit({"web/x.ts": "x\n"}, "fix")
        status, out = self.gate(author=rg.DEPENDABOT)
        self.assertEqual(status, 1, out)


class RecordTemplate(unittest.TestCase):
    def test_a_filled_skeleton_passes_the_reference_check(self):
        refs = importlib.util.spec_from_file_location("ref_check", os.path.join(HERE, "ref_check.py"))
        rc = importlib.util.module_from_spec(refs)
        refs.loader.exec_module(rc)
        doc = ".claude/skills/pr-review/references/review-record.md"
        with open(os.path.join(rg.ROOT, doc), encoding="utf-8") as handle:
            text = handle.read()
        skeleton = text.split("## The skeleton", 1)[1].split("```markdown\n", 1)[1].split("\n```", 1)[0]
        record_text = skeleton.replace("<number>", "212").replace("PR-<id>", "PR-63")
        known = rc.registers(rc.read(rc.DECISIONS), rc.read(rc.ROADMAP))
        findings = [f for f in rc.check("docs/reviews/pr-212.md", record_text, known, set(), set(), set()) if f[1] == "REF 1"]
        self.assertEqual(findings, [])


class Workflow(unittest.TestCase):
    def test_the_workflow_runs_this_file_from_the_base_and_never_the_head(self):
        with open(os.path.join(rg.ROOT, ".github/workflows/review-gate.yml"), encoding="utf-8") as handle:
            text = handle.read()
        self.assertIn("pull_request_target:", text)
        self.assertIn("python3 docs/tools/review_gate.py", text)
        self.assertIn("name: review-gate", text)
        self.assertNotIn("ref: ${{ github.event.pull_request.head", text)
        self.assertNotIn("write", text.split("permissions:", 1)[1].split("\n\n", 1)[0])


if __name__ == "__main__":
    unittest.main()
