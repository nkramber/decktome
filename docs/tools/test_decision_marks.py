"""Tests of the marks of amended rows in docs/decisions.md (D-924)."""
import os
import re
import unittest

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
ROW = re.compile(r"^\| (D-\d+)((?: \([a-z]+ by [^)]*\))*) \|")
CHANGE = re.compile(r"\b(amends|ends|supersedes|replaces|revises)\s+((?:D-\d+(?:,\s*|\s+and\s+|\s*))+)")


def unmarked(text):
    """List (old, new) pairs where row new changes row old and old names no new."""
    rows = {}
    for line in text.splitlines():
        m = ROW.match(line)
        if m:
            cells = line.split("|")
            rows[m.group(1)] = (m.group(2), cells[3] + cells[4] if len(cells) > 4 else "")
    missing = []
    for new, (_, text_of_row) in rows.items():
        for m in CHANGE.finditer(text_of_row):
            for old in re.findall(r"D-\d+", m.group(2)):
                if old in rows and int(old[2:]) < int(new[2:]) and new not in rows[old][0]:
                    missing.append((old, new))
    return missing


class DecisionMarks(unittest.TestCase):
    def test_each_changed_row_names_the_row_that_changed_it(self):
        with open(os.path.join(ROOT, "docs/decisions.md"), encoding="utf-8") as handle:
            self.assertEqual(unmarked(handle.read()), [])

    def test_a_row_with_no_mark_fails(self):
        text = "| D-1 | 2026-01-01 | A rule | x |\n| D-2 | 2026-01-02 | A new rule (amends D-1) | y |\n"
        self.assertEqual(unmarked(text), [("D-1", "D-2")])

    def test_a_marked_row_passes(self):
        text = "| D-1 (amended by D-2) | 2026-01-01 | A rule | x |\n| D-2 | 2026-01-02 | A new rule (amends D-1) | y |\n"
        self.assertEqual(unmarked(text), [])


if __name__ == "__main__":
    unittest.main()
