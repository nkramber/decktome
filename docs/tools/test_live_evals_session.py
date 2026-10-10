"""Tests of the one-turn rule of a live-eval session: no background task, and the last words in a notice (D-1223)."""

import json
import os
import re
import subprocess
import tempfile
import unittest

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
SCRIPT = os.path.join(ROOT, "scripts", "live-evals.sh")
PROMPTS = os.path.join(ROOT, "scripts", "live-evals")


def function_text(name):
    with open(SCRIPT, encoding="utf-8") as handle:
        text = handle.read()
    match = re.search(r"^" + name + r"\(\) \{.*?^\}$", text, re.M | re.S)
    if match is None:
        raise AssertionError(f"scripts/live-evals.sh has no {name} function")
    return match.group(0)


class SessionEnvTest(unittest.TestCase):
    def test_run_claude_turns_off_background_tasks(self):
        body = function_text("run_claude")
        env = body[body.index("env -i"):body.index('"$CLAUDE_BIN" -p')]
        self.assertIn("CLAUDE_CODE_DISABLE_BACKGROUND_TASKS=1", env)
        self.assertIn("BASH_MAX_TIMEOUT_MS=$(( TIMEOUT * 1000 ))", env)
        self.assertIn("BASH_DEFAULT_TIMEOUT_MS=", env)

    def test_setup_refuses_a_binary_without_the_switch(self):
        with open(SCRIPT, encoding="utf-8") as handle:
            text = handle.read()
        self.assertRegex(text, r'grep -a -q CLAUDE_CODE_DISABLE_BACKGROUND_TASKS "\$CLAUDE_BIN" \\\n\s+\|\| die')

    def test_prompts_name_the_one_turn_rule(self):
        with open(os.path.join(PROMPTS, "eval-prompt.md"), encoding="utf-8") as handle:
            prompt = handle.read()
        self.assertIn("## One turn, no background command", prompt)
        self.assertIn("`make codex-review PR=<number>` in the foreground", prompt)
        with open(os.path.join(PROMPTS, "restack-prompt.md"), encoding="utf-8") as handle:
            self.assertIn('section "One turn, no background command"', handle.read())


class LastWordsTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()

    def tearDown(self):
        self.tmp.cleanup()

    def last_words(self, lines):
        log = os.path.join(self.tmp.name, "session-2.log")
        with open(log, "w", encoding="utf-8") as handle:
            handle.write("\n".join(lines) + "\n")
        out = subprocess.run(["bash", "-c", function_text("last_words") + '\nlast_words "$1"', "bash", log],
                             check=True, capture_output=True, text=True)
        return out.stdout

    def test_reads_the_last_result_past_a_line_that_is_not_json(self):
        lines = [
            json.dumps({"type": "result", "result": "first"}),
            "a line of stderr",
            json.dumps({"type": "assistant", "message": {}}),
            json.dumps({"type": "result", "result": "I will wait for its completion notice.\nMore."}),
            json.dumps({"type": "system", "subtype": "task_updated"}),
        ]
        self.assertEqual(self.last_words(lines), "I will wait for its completion notice. More.\n")

    def test_cuts_a_long_result(self):
        self.assertEqual(self.last_words([json.dumps({"type": "result", "result": "é" * 400})]), "é" * 300 + "\n")

    def test_names_none_with_no_result(self):
        self.assertEqual(self.last_words(["not json", json.dumps({"type": "system"})]), "none\n")
        self.assertEqual(subprocess.run(["bash", "-c", function_text("last_words") + "\nlast_words /no/such/log"],
                                        check=True, capture_output=True, text=True).stdout, "none\n")


if __name__ == "__main__":
    unittest.main()
