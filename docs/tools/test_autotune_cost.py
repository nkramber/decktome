"""Tests of cost_of in scripts/autotune.sh (REV-075).

The loop charges each paid run to its ledger with the value cost_of reads.
A run document of today writes "- Calls: N. Cost: $X. Time: ...", a JSON
summary writes "cost_usd", and an older document writes "- Cost: $X.".
Each test runs the function alone in bash against a document read-only.

Run: python3 -m unittest discover -s docs/tools -p 'test_*.py'
"""
import json
import os
import re
import subprocess
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(os.path.dirname(HERE))
SCRIPT = os.path.join(ROOT, "scripts", "autotune.sh")
REF = os.path.join(ROOT, "docs", "reference")


def cost_of_source():
    with open(SCRIPT, encoding="utf-8") as f:
        text = f.read()
    m = re.search(r"^cost_of\(\) \{\n.*?^\}\n", text, re.M | re.S)
    if not m:
        raise AssertionError("scripts/autotune.sh defines no cost_of")
    return m.group(0)


def cost_of(path):
    res = subprocess.run(["bash", "-c", cost_of_source() + 'cost_of "$1"\n', "cost_of", path],
                         capture_output=True, text=True, timeout=30)
    return res.stdout


class CostOfTest(unittest.TestCase):
    def test_a_gate_document_of_today(self):
        self.assertEqual(cost_of(os.path.join(REF, "pr7-question-gate-run52.md")), "0.195100")

    def test_an_eval_document_of_today(self):
        self.assertEqual(cost_of(os.path.join(REF, "pr7-question-eval-run35.md")), "0.095900")

    def test_an_older_document(self):
        self.assertEqual(cost_of(os.path.join(REF, "pr7-question-eval-20260826-191225-000.md")), "0.095900")

    def test_a_json_summary(self):
        with tempfile.TemporaryDirectory() as tmp:
            for value, want in ((0.0959, "0.095900"), (1e-05, "0.000010")):
                path = os.path.join(tmp, f"sum-{want}.json")
                with open(path, "w", encoding="utf-8") as f:
                    json.dump({"run": "x", "cost_usd": value}, f, indent=2)
                with self.subTest(value=value):
                    self.assertEqual(cost_of(path), want)

    def test_an_unpriced_run_reads_empty(self):
        # charge stops the loop on an empty cost, so no unpriced run counts as $0.
        with tempfile.TemporaryDirectory() as tmp:
            doc = os.path.join(tmp, "gate.md")
            with open(doc, "w", encoding="utf-8") as f:
                f.write("## Run\n\n- Calls: 3. Cost: unpriced. Time: 9 seconds.\n")
            self.assertEqual(cost_of(doc), "")
            js = os.path.join(tmp, "sum.json")
            with open(js, "w", encoding="utf-8") as f:
                json.dump({"run": "x"}, f)
            self.assertEqual(cost_of(js), "")

    def test_a_cost_in_a_conversation_is_not_read(self):
        with tempfile.TemporaryDirectory() as tmp:
            doc = os.path.join(tmp, "gate.md")
            with open(doc, "w", encoding="utf-8") as f:
                f.write("- Calls: 3. Cost: $0.0100. Time: 9 seconds.\n\n"
                        "**Turn 1, the user:** Keep it cheap. Cost: $5 a card.\n")
            self.assertEqual(cost_of(doc), "0.010000")


if __name__ == "__main__":
    unittest.main()
