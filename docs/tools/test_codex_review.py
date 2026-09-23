"""Tests of codex_review.py (D-823 to D-832)."""
import contextlib
import importlib.util
import io
import json
import os
import subprocess
import tempfile
import unittest

HERE = os.path.dirname(os.path.abspath(__file__))
spec = importlib.util.spec_from_file_location("codex_review", os.path.join(HERE, "codex_review.py"))
cr = importlib.util.module_from_spec(spec)
spec.loader.exec_module(cr)

N = 213
R1, R2, R3, R4 = "1111111aaaa", "2222222bbbb", "3333333cccc", "4444444dddd"
GITAR = cr.GITAR
PUSHED = "2026-09-23T10:00:00Z"


def finding(fid, status="open", open_at=()):
    lines = [f"### {fid}: a defect", "", f"Status: {status}.", ""]
    if open_at:
        lines += ["Open at: " + ", ".join(f"`{h}`" for h in open_at) + ".", ""]
    return lines + ["File: `go/x.go:1-2`.", ""]


def record(head, verdict=cr.rg.APPROVED, items=()):
    lines = ["# Pull request 213 review", "", "## Identity", "", f"- PR: {N}", f"- Head: `{head}`", "", "## Findings", ""]
    for item in items:
        lines += item
    if not items:
        lines += ["No finding.", ""]
    lines += ["## Verdict", "", f"**{verdict}.** This verdict applies to head `{head}`.", ""]
    return "\n".join(lines)


class Version(unittest.TestCase):
    def test_a_release_sorts_above_its_pre_release(self):
        self.assertLess(cr.version("codex-cli 0.157.0-alpha.11"), cr.version("codex-cli 0.157.0"))

    def test_the_bundled_alpha_is_below_the_minimum(self):
        self.assertLess(cr.version("codex-cli 0.155.0-alpha.9.2"), cr.version(cr.MIN_VERSION))

    def test_the_homebrew_build_is_below_the_minimum(self):
        self.assertLess(cr.version("codex-cli 0.39.0"), cr.version(cr.MIN_VERSION))

    def test_text_with_no_version_gives_none(self):
        self.assertIsNone(cr.version("command not found"))


class ThreeStrikes(unittest.TestCase):
    """The owner's rule: the same id, counted once for each head at which it is open (D-826)."""

    def strike(self, items, head):
        return [f["id"] for f in cr.strikes(cr.findings(record(head, "Changes required", items)), head)]

    def test_a_finding_open_in_rounds_1_and_2_passes(self):
        self.assertEqual(self.strike([finding("P1-1", open_at=[R1, R2])], R2), [])

    def test_a_finding_open_in_round_3_stops(self):
        self.assertEqual(self.strike([finding("P1-1", open_at=[R1, R2, R3])], R3), ["P1-1"])

    def test_the_current_head_counts_when_the_line_omits_it(self):
        self.assertEqual(self.strike([finding("P1-1", open_at=[R1, R2])], R3), ["P1-1"])

    def test_a_fixed_then_reopened_finding_keeps_its_earlier_rounds(self):
        # Open at R1, fixed at R2, open again at R3: two rounds, so the loop goes on.
        self.assertEqual(self.strike([finding("P2-1", open_at=[R1, R3])], R3), [])
        # Open again at R4: the third round stops the loop.
        self.assertEqual(self.strike([finding("P2-1", open_at=[R1, R3, R4])], R4), ["P2-1"])

    def test_a_second_review_of_one_head_counts_once(self):
        self.assertEqual(self.strike([finding("P1-1", open_at=[R1, R2[:7], R2])], R2), [])

    def test_a_p3_never_stops_the_loop(self):
        self.assertEqual(self.strike([finding("P3-1", open_at=[R1, R2, R3])], R3), [])

    def test_a_closed_finding_never_stops_the_loop(self):
        items = [finding("P1-1", status=f"fixed in `{R3}`", open_at=[R1, R2, R3]),
                 finding("P1-2", status="withdrawn", open_at=[R1, R2, R3])]
        self.assertEqual(self.strike(items, R4), [])

    def test_a_record_with_no_open_at_line_counts_the_head_alone(self):
        items = cr.findings(record(R1, "Changes required", [finding("P0-1")]))
        self.assertEqual(len(cr.rounds(items[0], R1)), 1)


class Outcome(unittest.TestCase):
    def test_an_approval_exits_0(self):
        code, lines = cr.outcome(R1, cr.rg.APPROVED, [])
        self.assertEqual(code, cr.EXIT_APPROVE)
        self.assertIn("open findings: none", lines)

    def test_changes_required_and_blocked_exit_3(self):
        items = cr.findings(record(R1, "Changes required", [finding("P1-1")]))
        self.assertEqual(cr.outcome(R1, "Changes required", items)[0], cr.EXIT_CHANGES)
        self.assertEqual(cr.outcome(R1, "Blocked", [])[0], cr.EXIT_CHANGES)

    def test_a_third_round_exits_4_and_names_the_id(self):
        items = cr.findings(record(R3, "Changes required", [finding("P1-1", open_at=[R1, R2])]))
        code, lines = cr.outcome(R3, "Changes required", items)
        self.assertEqual(code, cr.EXIT_STRIKE)
        self.assertTrue(any(line.startswith("three-strike: P1-1") for line in lines))

    def test_the_codes_are_distinct(self):
        codes = [cr.EXIT_APPROVE, cr.EXIT_FAULT, cr.EXIT_USAGE, cr.EXIT_CHANGES, cr.EXIT_STRIKE, cr.EXIT_REFUSAL]
        self.assertEqual(len(set(codes)), len(codes))

    def test_two_bold_spans_are_a_fault(self):
        text = record(R1).replace("This verdict", "**Blocked.** This verdict")
        with self.assertRaises(cr.Stop) as caught:
            cr.verdict(text)
        self.assertEqual(caught.exception.code, cr.EXIT_FAULT)


def comment(login, body, created, updated=None):
    return {"user": {"login": login}, "body": body, "created_at": created, "updated_at": updated or created}


DASH = f"<details><summary>{cr.DASHBOARD}</summary></details>"


class GitarPass(unittest.TestCase):
    def test_a_dashboard_after_the_push_and_no_open_thread_pass(self):
        comments = [comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:02:00Z")]
        self.assertEqual(cr.gitar_problems(PUSHED, comments, [{"status": "completed"}], [{"isResolved": True}]), [])

    def test_a_dashboard_older_than_the_push_is_stale(self):
        comments = [comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T09:30:00Z")]
        problems = cr.gitar_problems(PUSHED, comments, [], [])
        self.assertTrue(any("before the push" in p for p in problems))

    def test_the_newest_dashboard_counts(self):
        comments = [comment(GITAR, DASH, "2026-09-23T08:00:00Z", "2026-09-23T10:05:00Z"),
                    comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T09:00:00Z")]
        self.assertTrue(cr.gitar_problems(PUSHED, comments, [], []))

    def test_no_dashboard_fails(self):
        self.assertTrue(any("no Gitar dashboard" in p for p in cr.gitar_problems(PUSHED, [], [], [])))

    def test_a_running_check_fails(self):
        comments = [comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:02:00Z")]
        self.assertTrue(cr.gitar_problems(PUSHED, comments, [{"status": "in_progress"}], []))

    def test_an_open_thread_fails(self):
        comments = [comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:02:00Z")]
        problems = cr.gitar_problems(PUSHED, comments, [], [{"isResolved": False, "path": "a.py", "line": 3}])
        self.assertEqual(problems, ["1 review thread(s) are not resolved: a.py:3."])

    def test_a_request_with_no_reply_fails(self):
        comments = [comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:02:00Z"),
                    comment("nkramber", "Gitar review", "2026-09-23T10:04:00Z")]
        self.assertTrue(any("no reply" in p for p in cr.gitar_problems(PUSHED, comments, [], [])))

    def test_a_refused_request_fails(self):
        comments = [comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:09:00Z"),
                    comment("nkramber", "Gitar review", "2026-09-23T10:04:00Z"),
                    comment(GITAR, "> Gitar review\n\nYou've sent several Gitar comments in a short window", "2026-09-23T10:05:00Z")]
        self.assertTrue(any("refused" in p for p in cr.gitar_problems(PUSHED, comments, [], [])))

    def test_a_request_needs_a_dashboard_after_the_reply(self):
        reply = comment(GITAR, "> Gitar review\n\nOn it", "2026-09-23T10:05:00Z")
        ask = comment("nkramber", "gitar review", "2026-09-23T10:04:00Z")
        before = comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:04:30Z")
        after = comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:08:00Z")
        self.assertTrue(cr.gitar_problems(PUSHED, [before, ask, reply], [], []))
        self.assertEqual(cr.gitar_problems(PUSHED, [after, ask, reply], [], []), [])


class Fake:
    """A runner that answers each command from a table of prefixes."""

    def __init__(self, table):
        self.table = table
        self.calls = []
        self.envs = []

    def __call__(self, cmd, cwd=None, timeout=None, stdout=None, stderr=None, env=None):
        self.calls.append(cmd)
        self.envs.append(env)
        for prefix, answer in self.table:
            if cmd[:len(prefix)] == prefix:
                return answer(cmd) if callable(answer) else answer
        return 1, "", f"unexpected: {cmd}"


class Refusals(unittest.TestCase):
    def pr(self, state="OPEN", cross=False):
        return 0, json.dumps({"state": state, "headRefName": "codex-review", "headRefOid": R1 * 4, "isCrossRepository": cross}), ""

    def test_a_merged_pull_request_refuses(self):
        with self.assertRaises(cr.Stop) as caught:
            cr.check_pr(Fake([(["gh", "pr", "view"], self.pr("MERGED"))]), N)
        self.assertEqual(caught.exception.code, cr.EXIT_REFUSAL)
        self.assertIn("MERGED", str(caught.exception))

    def test_a_fork_refuses(self):
        with self.assertRaises(cr.Stop):
            cr.check_pr(Fake([(["gh", "pr", "view"], self.pr(cross=True))]), N)

    def checkout(self, local, remote, status=""):
        return Fake([(["git", "fetch"], (0, "", "")),
                     (["git", "rev-parse", "HEAD"], (0, local, "")),
                     (["git", "rev-parse", "origin/b"], (0, remote, "")),
                     (["git", "status"], (0, status, ""))])

    def test_a_local_branch_that_differs_from_origin_refuses(self):
        with self.assertRaises(cr.Stop) as caught:
            cr.check_checkout(self.checkout(R2, R1), ".", "b", R1)
        self.assertIn("Push or pull", str(caught.exception))

    def test_a_stale_fetch_refuses(self):
        with self.assertRaises(cr.Stop):
            cr.check_checkout(self.checkout(R1, R2), ".", "b", R1)

    def test_a_dirty_tree_refuses_and_names_a_path(self):
        with self.assertRaises(cr.Stop) as caught:
            cr.check_checkout(self.checkout(R1, R1, " M go/x.go\n?? y.txt\n"), ".", "b", R1)
        self.assertIn("2 uncommitted", str(caught.exception))
        self.assertIn("go/x.go", str(caught.exception))

    def test_a_clean_head_passes(self):
        cr.check_checkout(self.checkout(R1, R1), ".", "b", R1)

    def cli(self, installed, latest, install=(0, "", "")):
        return Fake([(["npm", "install"], install),
                     (["npm", "view"], (0, latest + "\n", "")),
                     (["npm", "prefix"], (0, "/n\n", "")),
                     (["/n/bin/codex", "--version"], (0, f"codex-cli {installed}\n", ""))])

    def test_the_cli_is_updated_before_each_review(self):
        run = self.cli(cr.MIN_VERSION, cr.MIN_VERSION)
        self.assertEqual(cr.update_cli(run)[0], "/n/bin/codex")
        self.assertEqual(run.calls[0][:3], ["npm", "install", "-g"])
        self.assertIn(f"{cr.NPM_PACKAGE}@latest", run.calls[0])

    def test_a_failed_update_refuses(self):
        with self.assertRaises(cr.Stop):
            cr.update_cli(self.cli(cr.MIN_VERSION, cr.MIN_VERSION, install=(1, "", "EACCES")))

    def test_a_cli_behind_the_newest_release_refuses(self):
        with self.assertRaises(cr.Stop):
            cr.update_cli(self.cli("0.156.1", "0.157.0"))

    def test_a_cli_below_the_minimum_refuses(self):
        with self.assertRaises(cr.Stop) as caught:
            cr.update_cli(self.cli("0.100.0", "0.100.0"))
        self.assertIn("below the minimum", str(caught.exception))

    def probe(self, code, answer):
        def run(cmd, cwd=None, **_):
            if answer is not None:
                with open(cmd[cmd.index("-o") + 1], "w", encoding="utf-8") as handle:
                    handle.write(answer)
            return code, "", "ERROR: The 'x' model is not supported"
        return run

    def test_a_probe_that_answers_ok_passes(self):
        cr.probe(self.probe(0, "OK\n"), "codex")

    def test_a_rejected_model_refuses(self):
        with self.assertRaises(cr.Stop) as caught:
            cr.probe(self.probe(1, None), "codex")
        self.assertIn("not supported", str(caught.exception))

    def test_a_wrong_answer_refuses(self):
        with self.assertRaises(cr.Stop):
            cr.probe(self.probe(0, "I can not do that."), "codex")


class Invocation(unittest.TestCase):
    def test_each_call_names_the_model_and_the_effort(self):
        args = cr.model_args()
        self.assertEqual(args[:2], ["-m", "gpt-6-luna"])
        self.assertIn('model_reasoning_effort="medium"', args)
        self.assertIn('approval_policy="never"', args)

    def test_the_prompt_loads_the_skill_and_pushes_to_the_branch(self):
        text = cr.prompt(N, "o/r", "codex-review", R1)
        self.assertIn(".claude/skills/pr-review/SKILL.md", text)
        self.assertIn("references/commit-and-end.md", text)
        self.assertIn("git push origin HEAD:codex-review", text)

    def test_no_codex_call_sees_an_api_key(self):
        # D-833: a review bills the ChatGPT plan, and never the API.
        os.environ["OPENAI_API_KEY"] = "sk-test"
        os.environ["CODEX_API_KEY"] = "sk-test"
        try:
            env = cr.codex_env()
        finally:
            del os.environ["OPENAI_API_KEY"], os.environ["CODEX_API_KEY"]
        self.assertNotIn("OPENAI_API_KEY", env)
        self.assertNotIn("CODEX_API_KEY", env)
        self.assertIn("PATH", env)

    def test_the_probe_and_the_login_check_run_with_no_key(self):
        os.environ["OPENAI_API_KEY"] = "sk-test"
        try:
            run = Fake([(["codex", "login"], (0, "Logged in using ChatGPT\n", "")),
                        (["codex", "exec"], (1, "", "ERROR"))])
            cr.check_login(run, "codex")
            with self.assertRaises(cr.Stop):
                cr.probe(run, "codex")
        finally:
            del os.environ["OPENAI_API_KEY"]
        self.assertEqual(len(run.envs), 2)
        for env in run.envs:
            self.assertIsNotNone(env)
            self.assertNotIn("OPENAI_API_KEY", env)

    def test_an_api_key_login_refuses(self):
        run = Fake([(["codex", "login"], (0, "Logged in using an API key - sk-***\n", ""))])
        with self.assertRaises(cr.Stop) as caught:
            cr.check_login(run, "codex")
        self.assertEqual(caught.exception.code, cr.EXIT_REFUSAL)
        self.assertIn("ChatGPT", str(caught.exception))

    def test_a_missing_pr_number_is_a_usage_error(self):
        with contextlib.redirect_stderr(io.StringIO()):
            self.assertEqual(cr.main([], run=Fake([])), cr.EXIT_USAGE)


def git(repo, *args):
    return subprocess.run(["git", *args], cwd=repo, check=True, capture_output=True, text=True).stdout.strip()


class ReadResult(unittest.TestCase):
    """The read of part 3, over a real origin and two clones."""

    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        root = self.tmp.name
        self.origin = os.path.join(root, "origin.git")
        git(root, "init", "-q", "--bare", "-b", "main", self.origin)
        self.author = self.clone("author")
        self.write(self.author, "README.md", "base\n")
        self.commit(self.author, "base")
        git(self.author, "push", "-q", "origin", "main")
        git(self.author, "switch", "-q", "-c", "b")
        self.write(self.author, "go/x.go", "package x\n")
        self.commit(self.author, "code")
        git(self.author, "push", "-q", "origin", "b")
        self.head = git(self.author, "rev-parse", "HEAD")
        self.reviewer = self.clone("reviewer")
        git(self.reviewer, "switch", "-q", "--detach", "origin/b")

    def tearDown(self):
        self.tmp.cleanup()

    def clone(self, name):
        path = os.path.join(self.tmp.name, name)
        git(self.tmp.name, "clone", "-q", self.origin, path)
        git(path, "config", "user.email", "t@example.com")
        git(path, "config", "user.name", "t")
        return path

    def write(self, repo, path, text):
        os.makedirs(os.path.dirname(os.path.join(repo, path)) or repo, exist_ok=True)
        with open(os.path.join(repo, path), "w", encoding="utf-8") as handle:
            handle.write(text)

    def commit(self, repo, message):
        git(repo, "add", "-A")
        git(repo, "commit", "-q", "-m", message)

    def push_review(self, text, extra=None):
        self.write(self.reviewer, cr.rg.record_path(N), text)
        self.write(self.reviewer, "docs/SESSION-HANDOFF.md", "state\n")
        if extra:
            self.write(self.reviewer, extra, "x\n")
        self.commit(self.reviewer, "The review record")
        git(self.reviewer, "push", "-q", "origin", "HEAD:b")

    def read(self):
        return cr.read_result(cr.sh, self.author, N, "b", self.head)

    def fault(self):
        with self.assertRaises(cr.Stop) as caught:
            self.read()
        self.assertEqual(caught.exception.code, cr.EXIT_FAULT)
        return str(caught.exception)

    def test_an_approved_record_of_the_effective_head_reads(self):
        self.push_review(record(self.head[:7]))
        head, name, items = self.read()
        self.assertEqual((head, name, items), (self.head, cr.rg.APPROVED, []))

    def test_a_record_of_another_head_is_a_fault(self):
        self.push_review(record("abcdef0"))
        self.assertIn("effective head", self.fault())

    def test_no_push_is_a_fault(self):
        self.assertIn("pushed no record", self.fault())

    def test_a_push_outside_the_metadata_set_is_a_fault(self):
        self.push_review(record(self.head[:7]), extra="go/y.go")
        self.assertIn("go/y.go", self.fault())

    def test_a_push_with_no_record_is_a_fault(self):
        self.write(self.reviewer, "docs/SESSION-HANDOFF.md", "state\n")
        self.commit(self.reviewer, "The hand-off")
        git(self.reviewer, "push", "-q", "origin", "HEAD:b")
        self.assertIn("holds no", self.fault())


if __name__ == "__main__":
    unittest.main()
