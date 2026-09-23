"""Tests that the pause notes and the pause file of D-838 agree.

Each pause note is one paragraph or bullet that starts with NOTE. The
notes exist while PAUSE exists, and the end of the pause deletes both.
"""
import os
import re
import subprocess
import unittest

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
PAUSE = "docs/reference/gitar-pause.md"
NOTE = "**Gitar pause (D-838).**"
START = re.compile(r"^\s*(- )?" + re.escape(NOTE))
SKIP = {PAUSE, "docs/tools/test_gitar_pause.py"}
# The files that the rules of each session read. Each one holds a note while the pause lasts.
MUST = ["CLAUDE.md", "AGENTS.md", ".claude/skills/gitar-review/SKILL.md", ".claude/skills/pr-review/SKILL.md",
        ".claude/skills/pr-review/references/answer-review.md", ".claude/skills/one-pr-one-session/SKILL.md"]


def notes():
    """Each tracked file and line that holds NOTE, as (path, line)."""
    out = subprocess.run(["git", "grep", "-n", "-F", NOTE], cwd=ROOT, capture_output=True, text=True)
    if out.returncode not in (0, 1):
        raise RuntimeError(out.stderr)
    found = []
    for row in out.stdout.splitlines():
        path, _, line = row.split(":", 2)
        if path not in SKIP:
            found.append((path, line))
    return found


class GitarPause(unittest.TestCase):
    def test_the_notes_and_the_pause_file_agree(self):
        found = notes()
        if os.path.exists(os.path.join(ROOT, PAUSE)):
            paths = {path for path, _ in found}
            for path in MUST:
                self.assertIn(path, paths, f"{path} holds no pause note, and {PAUSE} exists")
        else:
            self.assertEqual(found, [], f"{PAUSE} is gone, so delete each pause note")

    def test_each_note_starts_its_paragraph(self):
        # A note that starts its line is one deletion at the end of the pause.
        for path, line in notes():
            self.assertRegex(line, START, f"{path}: the pause note does not start its paragraph")


if __name__ == "__main__":
    unittest.main()
