"""Tests of the context budget check (D-749).

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import importlib.util
import os
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("context_budget", os.path.join(HERE, "context_budget.py"))
cb = importlib.util.module_from_spec(spec)
spec.loader.exec_module(cb)

PAID_LINE = "Two targets and one loop script spend money: `make deck-gate`, `make bracket-gate`, and `scripts/autotune.sh`. Ask the owner first."

CLAUDE = f"# Rules\n\n## Commands that cost money\n\n{PAID_LINE}\n"
HANDOFF = f"""# Session hand-off

## RESUME HERE (2026-09-16)

The state.

## How to resume

- A trap.
- {PAID_LINE}
- Another trap.

## The three most recent sessions

### 2026-09-16c: one

Text.

### 2026-09-16b: two

Text.

## The archive

Pointer.
"""
PAID = f"# Paid targets\n\n{PAID_LINE}\n"
MAKEFILE = "deck-gate: ## paid\n\t@true\nbracket-gate:\n\t@true\n"


def files(**override):
    base = {cb.CLAUDE: CLAUDE, cb.HANDOFF: HANDOFF, cb.PAID: PAID, "Makefile": MAKEFILE}
    base.update({k.replace("__", "/"): v for k, v in override.items()})
    return base


def run(texts, scripts=("scripts/autotune.sh",), skills=()):
    return cb.check(lambda path: texts.get(path), lambda path: path in scripts, skills)


class ContextBudgetTest(unittest.TestCase):
    def test_small_files_pass(self):
        report, errors = run(files())
        self.assertEqual(errors, [])
        self.assertIn("paid targets: 3 named in 3 files", report)

    def test_file_over_limit_fails(self):
        texts = files()
        texts[cb.CLAUDE] = CLAUDE + "x" * cb.FILE_LIMITS[cb.CLAUDE]
        _, errors = run(texts)
        self.assertTrue(any(e.startswith("CLAUDE.md holds") for e in errors), errors)

    def test_resume_section_over_limit_fails(self):
        texts = files()
        texts[cb.HANDOFF] = HANDOFF.replace("The state.", "y" * (cb.RESUME_LIMIT + 1))
        _, errors = run(texts)
        self.assertTrue(any(e.startswith("the resume section holds") for e in errors), errors)

    def test_resume_limit_reads_its_own_section_alone(self):
        texts = files()
        texts[cb.HANDOFF] = HANDOFF.replace("Pointer.", "z" * (cb.RESUME_LIMIT + 1))
        _, errors = run(texts)
        self.assertFalse(any(e.startswith("the resume section") for e in errors), errors)

    def test_too_many_session_records_fail(self):
        texts = files()
        extra = "### 2026-09-15: three\n\nText.\n\n### 2026-09-14: four\n\nText.\n\n## The archive"
        texts[cb.HANDOFF] = HANDOFF.replace("## The archive", extra)
        _, errors = run(texts)
        self.assertTrue(any("4 session records" in e for e in errors), errors)

    def test_missing_sections_fail(self):
        texts = files()
        texts[cb.HANDOFF] = "# Session hand-off\n\nNo sections.\n"
        _, errors = run(texts)
        self.assertTrue(any("RESUME HERE" in e for e in errors), errors)
        self.assertTrue(any("three most recent sessions" in e for e in errors), errors)

    def test_paid_lists_that_differ_fail(self):
        texts = files()
        texts[cb.CLAUDE] = CLAUDE.replace("`make bracket-gate`, ", "")
        _, errors = run(texts)
        self.assertTrue(any("differ" in e and "make bracket-gate" in e for e in errors), errors)

    def test_paid_target_that_is_no_makefile_target_fails(self):
        line = PAID_LINE.replace("`make deck-gate`", "`make deck-gat`")
        texts = files()
        for path in (cb.CLAUDE, cb.HANDOFF, cb.PAID):
            texts[path] = texts[path].replace(PAID_LINE, line)
        _, errors = run(texts)
        self.assertIn("the paid target `make deck-gat` is no Makefile target", errors)

    def test_missing_paid_script_fails(self):
        _, errors = run(files(), scripts=())
        self.assertIn("the paid script `scripts/autotune.sh` does not exist", errors)

    def test_paid_names_stop_at_the_end_of_the_sentence(self):
        text = "Two spend money: `make a-b`. The free `make c-d` stays out.\n"
        self.assertEqual(cb.paid_names(text), {"make a-b"})


class SkillSizeTest(unittest.TestCase):
    SMALL = ".claude/skills/one-pr-one-session/SKILL.md"
    BIG = ".claude/skills/mtg-corpus/SKILL.md"

    def test_a_small_skill_passes(self):
        texts = files()
        texts[self.SMALL] = "x"
        report, errors = run(texts, skills=[self.SMALL])
        self.assertEqual(errors, [])
        self.assertIn(f"skill files: 1 under {cb.SKILL_LIMIT} bytes, 0 over, 0 exempt", report)

    def test_a_skill_over_the_limit_fails(self):
        texts = files()
        texts[self.SMALL] = "x" * (cb.SKILL_LIMIT + 1)
        _, errors = run(texts, skills=[self.SMALL])
        self.assertTrue(any(e.startswith(self.SMALL) for e in errors), errors)

    def test_the_exempt_skill_passes_over_the_limit(self):
        texts = files()
        texts[self.BIG] = "x" * (cb.SKILL_LIMIT + 1)
        report, errors = run(texts, skills=[self.BIG])
        self.assertEqual(errors, [])
        self.assertIn("0 over, 1 exempt", report[-1])


if __name__ == "__main__":
    unittest.main()
