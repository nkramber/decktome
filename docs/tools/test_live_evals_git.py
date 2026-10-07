"""Tests of session_credential in live-evals.sh: no lower helper stores the token (D-1209)."""

import os
import re
import subprocess
import tempfile
import unittest

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
SCRIPT = os.path.join(ROOT, "scripts", "live-evals.sh")


def function_text():
    with open(SCRIPT, encoding="utf-8") as handle:
        text = handle.read()
    match = re.search(r"^session_credential\(\) \{.*?^\}$", text, re.M | re.S)
    if match is None:
        raise AssertionError("scripts/live-evals.sh has no session_credential function")
    return match.group(0)


class SessionCredentialTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        base = self.tmp.name
        self.repo = os.path.join(base, "repo")
        self.log = os.path.join(base, "helpers.log")
        bindir = os.path.join(base, "bin")
        os.makedirs(bindir)
        # The fake gh logs each call and answers a get like gh auth git-credential.
        gh = os.path.join(bindir, "gh")
        with open(gh, "w") as handle:
            handle.write(
                "#!/bin/sh\n"
                f'echo "gh $3" >> "{self.log}"\n'
                'cat > /dev/null\n'
                'if [ "$3" = get ]; then printf "username=x-access-token\\npassword=fake\\n"; fi\n'
            )
        os.chmod(gh, 0o755)
        # A lower config holds a helper in the place of osxkeychain. The
        # test sets GIT_CONFIG_NOSYSTEM, so it never reaches a real keychain.
        lower = os.path.join(base, "global.gitconfig")
        with open(lower, "w") as handle:
            handle.write(f"[credential]\n\thelper = \"!f() {{ cat > /dev/null; echo lower $1 >> '{self.log}'; }}; f\"\n")
        self.env = {
            "PATH": bindir + os.pathsep + "/usr/bin:/bin",
            "HOME": base,
            "GIT_CONFIG_NOSYSTEM": "1",
            "GIT_CONFIG_GLOBAL": lower,
            "GIT_TERMINAL_PROMPT": "0",
        }
        subprocess.run(["git", "init", "-q", self.repo], check=True, env=self.env)
        subprocess.run(["bash", "-c", function_text() + '\nsession_credential "$1"', "bash", self.repo], check=True, env=self.env)

    def tearDown(self):
        self.tmp.cleanup()

    def credential(self, action, text):
        return subprocess.run(
            ["git", "-C", self.repo, "credential", action],
            input=text, capture_output=True, text=True, check=True, env=self.env,
        ).stdout

    def calls(self):
        if not os.path.exists(self.log):
            return []
        with open(self.log) as handle:
            return handle.read().split()

    def test_helper_list_resets_then_names_gh(self):
        out = subprocess.run(
            ["git", "-C", self.repo, "config", "--local", "--get-all", "credential.helper"],
            capture_output=True, text=True, check=True, env=self.env,
        ).stdout
        self.assertEqual(out.split("\n")[:2], ["", "!gh auth git-credential"])

    def test_store_reaches_gh_alone(self):
        got = self.credential("fill", "protocol=https\nhost=github.com\n\n")
        self.assertIn("username=x-access-token", got)
        self.credential("approve", got)
        self.assertNotIn("lower", self.calls())
        self.assertEqual(self.calls(), ["gh", "get", "gh", "store"])


if __name__ == "__main__":
    unittest.main()
