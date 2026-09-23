"""Tests of codex_review.py (D-823 to D-832)."""
import contextlib
import importlib.util
import io
import json
import os
import subprocess
import tempfile
import unittest
import unittest.mock

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


def thread_page(threads, more):
    return {"data": {"repository": {"pullRequest": {"reviewThreads": {
        "nodes": threads, "pageInfo": {"hasNextPage": more, "endCursor": "c1" if more else None}}}}}}


class Threads(unittest.TestCase):
    def test_an_open_thread_on_a_later_page_fails_the_pass(self):
        first = [{"isResolved": True, "path": "a.py", "line": i} for i in range(100)]
        later = [{"isResolved": False, "path": "b.py", "line": 7}]
        run = Fake([(["gh", "api", "graphql"], (0, json.dumps([thread_page(first, True), thread_page(later, False)]), ""))])
        threads = cr.review_threads(run, "o/r", N)
        self.assertEqual(len(threads), 101)
        self.assertIn("--paginate", run.calls[0])
        self.assertIn("after: $endCursor", run.calls[0][-1])
        problems = cr.gitar_problems(PUSHED, [comment(GITAR, DASH, "2026-09-23T09:00:00Z", "2026-09-23T10:02:00Z")], [], threads)
        self.assertEqual(problems, ["1 review thread(s) are not resolved: b.py:7."])


def summary(*kbds):
    """A Gitar dashboard body with one summary line of the Code Review block."""
    return "<details>\n<summary><b>Code Review</b> " + " ".join(f"<kbd>{k}</kbd>" for k in kbds) + "</summary>\n</details>"


class Dashboard(unittest.TestCase):
    """dashboard_issue reads the newest Gitar dashboard, and each unknown form is an issue (D-838)."""

    def issue(self, *bodies):
        return cr.dashboard_issue([comment(GITAR, b, f"2026-09-23T10:0{i}:00Z") for i, b in enumerate(bodies)])

    def test_each_clean_form_of_the_record_passes(self):
        # The forms of 90 dashboards of #120 to #216, read 2026-09-23.
        for kbds in [("\u2705 Approved",), ("\u2705 No issues found",), ("\u2705 Approved", "1 resolved / 1 findings"),
                     ("\u2705 Approved", "2 closed / 2 findings"), ("\u2705 No issues found", "2 closed / 2 findings")]:
            self.assertIsNone(self.issue(summary(*kbds)), kbds)

    def test_an_open_finding_is_an_issue(self):
        self.assertTrue(self.issue(summary("\u2705 Approved", "1 resolved / 2 findings")))

    def test_an_unknown_verdict_is_an_issue(self):
        self.assertTrue(self.issue(summary("Changes requested")))
        self.assertTrue(self.issue(summary("\u2705 Approved with suggestions")))

    def test_an_unknown_tally_is_an_issue(self):
        self.assertTrue(self.issue(summary("\u2705 Approved", "1 open")))
        self.assertTrue(self.issue(summary("\u2705 Approved", "1 resolved / 1 findings", "new")))

    def test_a_dashboard_with_no_summary_line_is_an_issue(self):
        self.assertTrue(self.issue("<b>Code Review</b> in a new shape"))

    def test_the_newest_dashboard_counts(self):
        self.assertIsNone(self.issue(summary("Changes requested"), summary("\u2705 Approved")))
        self.assertTrue(self.issue(summary("\u2705 Approved"), summary("Changes requested")))

    def test_the_free_plan_note_alone_is_no_issue(self):
        self.assertIsNone(self.issue("> [!IMPORTANT]\n> You are using the Gitar free plan."))
        self.assertIsNone(cr.dashboard_issue([comment("nkramber", summary("Changes requested"), PUSHED)]))


class SkipGitar(unittest.TestCase):
    """--skip-gitar-review reads no Gitar pass, and still refuses an open thread (D-838)."""

    def threads(self, *nodes, comments=()):
        return Fake([(["gh", "api", "graphql"], (0, json.dumps([thread_page(list(nodes), False)]), "")),
                     (["gh", "api", "--paginate"], (0, json.dumps([list(comments)]), ""))])

    def test_an_open_thread_refuses_and_names_the_owner(self):
        run = self.threads({"isResolved": True, "path": "a.py", "line": 1}, {"isResolved": False, "path": "b.py", "line": 7})
        with self.assertRaises(cr.Stop) as caught:
            cr.check_threads(run, "o/r", N)
        self.assertEqual(caught.exception.code, cr.EXIT_REFUSAL)
        self.assertIn("1 review thread(s) are not resolved: b.py:7.", str(caught.exception))
        self.assertIn("tell the owner (D-838)", str(caught.exception))

    def test_resolved_threads_and_a_clean_dashboard_pass(self):
        run = self.threads({"isResolved": True, "path": "a.py", "line": 1}, comments=[comment(GITAR, summary("\u2705 Approved"), PUSHED)])
        cr.check_threads(run, "o/r", N)
        self.assertEqual(len(run.calls), 2)

    def test_a_dashboard_issue_with_no_thread_refuses(self):
        run = self.threads(comments=[comment(GITAR, summary("\u2705 Approved", "0 resolved / 1 findings"), PUSHED)])
        with self.assertRaises(cr.Stop) as caught:
            cr.check_threads(run, "o/r", N)
        self.assertEqual(caught.exception.code, cr.EXIT_REFUSAL)
        self.assertIn("the Gitar dashboard reports an issue", str(caught.exception))
        self.assertIn("tell the owner (D-838)", str(caught.exception))

    def run_main(self, argv):
        """Run main up to the CLI update, and give the Gitar calls and the thread calls."""
        seen = {"gitar": 0, "threads": 0}

        def gitar(*_):
            seen["gitar"] += 1

        def threads(*_):
            seen["threads"] += 1

        def stop(_):
            raise cr.refuse("the CLI update stops the test.")

        patches = [
            unittest.mock.patch.object(cr, "check_pr", return_value=("b", R1 * 4)),
            unittest.mock.patch.object(cr, "check_checkout"),
            unittest.mock.patch.object(cr.rg, "gather", return_value=(None, [], None, None)),
            unittest.mock.patch.object(cr.rg, "effective_head", return_value=R1 * 4),
            unittest.mock.patch.object(cr, "check_gitar", side_effect=gitar),
            unittest.mock.patch.object(cr, "check_threads", side_effect=threads),
            unittest.mock.patch.object(cr, "update_cli", side_effect=stop),
        ]
        with contextlib.ExitStack() as stack:
            for patch in patches:
                stack.enter_context(patch)
            stack.enter_context(contextlib.redirect_stdout(io.StringIO()))
            code = cr.main(argv, run=Fake([(["gh", "repo", "view"], (0, "o/r\n", ""))]))
        return code, seen

    def test_the_flag_skips_the_gitar_pass_and_reads_the_threads(self):
        code, seen = self.run_main(["--pr", str(N), "--skip-gitar-review"])
        self.assertEqual(code, cr.EXIT_REFUSAL)
        self.assertEqual(seen, {"gitar": 0, "threads": 1})

    def test_with_no_flag_the_gitar_pass_runs(self):
        code, seen = self.run_main(["--pr", str(N)])
        self.assertEqual(code, cr.EXIT_REFUSAL)
        self.assertEqual(seen, {"gitar": 1, "threads": 0})


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

    def test_the_review_prepares_a_worktree_and_runs_codex_there_with_no_key(self):
        os.environ["OPENAI_API_KEY"] = "sk-test"
        try:
            with tempfile.TemporaryDirectory() as repo:
                run = Fake([(["git", "worktree"], (0, "", "")), (["pnpm"], (0, "", "")), (["/n/codex", "exec"], (0, "", ""))])
                code, tree, base = cr.review(run, repo, "/n/codex", N, "o/r", "b", R1, "20260923T000000Z")
                os.rmdir(tree)
        finally:
            del os.environ["OPENAI_API_KEY"]
        self.assertEqual(code, 0)
        self.assertEqual([c[0] for c in run.calls], ["git", "pnpm", "/n/codex"])
        self.assertEqual(run.calls[0][-2:], [tree, R1])
        exec_call = run.calls[2]
        self.assertEqual(exec_call[exec_call.index("-s") + 1], cr.SANDBOX)
        self.assertEqual(exec_call[exec_call.index("-C") + 1], tree)
        self.assertNotIn("OPENAI_API_KEY", run.envs[2])
        self.assertTrue(base.endswith(f"pr-{N}-20260923T000000Z"))

    def test_a_failed_install_removes_the_worktree(self):
        with tempfile.TemporaryDirectory() as repo:
            run = Fake([(["git", "worktree", "add"], (0, "", "")), (["pnpm"], (1, "", "ERR_PNPM_OUTDATED_LOCKFILE")),
                        (["git", "worktree", "remove"], (0, "", ""))])
            with self.assertRaises(cr.Stop) as caught:
                cr.review(run, repo, "/n/codex", N, "o/r", "b", R1, "20260923T000000Z")
        tree = run.calls[0][-2]
        self.assertEqual(caught.exception.code, cr.EXIT_FAULT)
        self.assertEqual(run.calls[-1], ["git", "worktree", "remove", "--force", tree])
        self.assertFalse(os.path.exists(tree))
        self.assertNotIn("/n/codex", [c[0] for c in run.calls])

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

    def test_the_push_walk_of_a_merge_head_skips_the_commits_of_main(self):
        # The branch merges a newer main. The merge is the effective head.
        git(self.author, "switch", "-q", "main")
        self.write(self.author, "go/main.go", "package main\n")
        self.commit(self.author, "main1")
        main1 = git(self.author, "rev-parse", "HEAD")
        git(self.author, "switch", "-q", "b")
        git(self.author, "merge", "-q", "--no-edit", "main")
        merge = git(self.author, "rev-parse", "HEAD")
        self.write(self.author, "docs/SESSION-HANDOFF.md", "state\n")
        self.commit(self.author, "hand-off")
        tip = git(self.author, "rev-parse", "HEAD")
        shas = cr.later_commits(cr.sh, self.author, merge, tip)
        self.assertEqual(shas, [merge, tip])
        self.assertNotIn(main1, shas)

    def test_a_push_with_no_record_is_a_fault(self):
        self.write(self.reviewer, "docs/SESSION-HANDOFF.md", "state\n")
        self.commit(self.reviewer, "The hand-off")
        git(self.reviewer, "push", "-q", "origin", "HEAD:b")
        self.assertIn("holds no", self.fault())


if __name__ == "__main__":
    unittest.main()
