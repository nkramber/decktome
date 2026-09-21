#!/usr/bin/env python3
"""Check that every pipeline of the Makefile fails its target (F-160).

GNU Make 3.82 added the .SHELLFLAGS variable. GNU Make 3.81 ignores it,
and the Mac of the owner runs 3.81. So a recipe that reads pipefail from
.SHELLFLAGS alone runs without pipefail there, and a pipeline reports the
status of its last command. On 2026-09-20 the first live run of
make api-build failed with "card database not loaded yet" and the target
exited 0, because tee succeeded.

The rule: every recipe line that holds a pipeline sets pipefail itself,
with "set -o pipefail;" before the pipeline. A recipe whose exit code
carries no fault is exempt. It carries a make comment line directly above
it that starts with "pipefail-ok:" and gives the reason.

The check fails when:
- the Makefile sets no SHELL,
- a recipe line holds a pipeline and neither sets pipefail nor is exempt,
- an exempt marker names no recipe line with a pipeline,
- the make of this machine does not fail a guarded pipeline.

The last check is a live proof. It runs the make of this machine over a
fixture makefile that carries the SHELL lines of the real Makefile. So
the proof reads GNU Make 3.81 on the Mac of the owner and GNU Make 4 on
the runner of the verify workflow.

Run: python3 docs/tools/pipefail_check.py
"""
import os
import re
import subprocess
import sys
import tempfile

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

MAKEFILE = "Makefile"
GUARD = "set -o pipefail"
MARKER = "pipefail-ok:"
SHELL_LINE = re.compile(r"^(SHELL|\.SHELLFLAGS)\s*[:?+]?=")


def recipe_lines(text):
    """Return one entry for each recipe line: its number, text, and marker.

    A backslash joins a line to the next one, because both reach one
    shell. The marker is the comment line directly above the recipe line.
    """
    lines = text.split("\n")
    out = []
    marker = ""
    index = 0
    while index < len(lines):
        first = index
        joined = lines[index]
        while joined.endswith("\\") and index + 1 < len(lines):
            index += 1
            joined = joined[:-1] + " " + lines[index]
        bare = lines[first].lstrip()
        if lines[first].startswith("\t"):
            out.append((first + 1, joined, marker))
            marker = ""
        elif bare.startswith("#"):
            marker = joined.strip()
        elif bare:
            marker = ""
        index += 1
    return out


def unquoted(line):
    """Return the line without the text inside single or double quotes.

    A quoted pipe is data, such as the -run regex of a go test call.
    """
    out = []
    quote = ""
    for char in line:
        if quote:
            if char == quote:
                quote = ""
            continue
        if char in "'\"":
            quote = char
            continue
        out.append(char)
    return "".join(out)


def has_pipeline(line):
    """Say whether the line holds a pipe that bash reads as a pipeline."""
    return "|" in unquoted(line).replace("||", " ")


def shell_preamble(text):
    """Return the SHELL and .SHELLFLAGS lines of the makefile."""
    return [line for line in text.split("\n") if SHELL_LINE.match(line)]


FIXTURE = """\
%s

guarded-fails:
\t@set -o pipefail; false | cat

guarded-passes:
\t@set -o pipefail; true | cat

bare:
\t@false | cat
"""


def prove(preamble, run):
    """Run the make of this machine over the fixture, and read the result.

    The fixture carries the SHELL lines of the real Makefile. A guarded
    pipeline must fail its target, and a guarded pipeline that succeeds
    must not fail it. The bare target reads whether this make honors
    .SHELLFLAGS, which is a record and not a rule.
    """
    report = []
    errors = []
    version = run(["--version"], text=True)
    name = version.splitlines()[0].strip() if version else "unknown make"
    fails = run(["guarded-fails"])
    passes = run(["guarded-passes"])
    bare = run(["bare"])
    if fails == 0:
        errors.append(f"{name} does not fail a guarded pipeline. The recipe rule of F-160 does not hold here")
    if passes != 0:
        errors.append(f"{name} fails 'set -o pipefail' itself. The Makefile needs SHELL := bash: {preamble}")
    honors = "honors" if bare != 0 else "ignores"
    report.append(f"{name} {honors} .SHELLFLAGS, and it fails a guarded pipeline")
    return report, errors


def make_runner(preamble):
    """Return a function that runs the make of this machine on the fixture."""
    environment = {k: v for k, v in os.environ.items() if not k.startswith(("MAKE", "MFLAGS"))}
    binary = os.environ.get("MAKE") or "make"

    def run(args, text=False):
        with tempfile.TemporaryDirectory() as work:
            path = os.path.join(work, "Makefile")
            with open(path, "w", encoding="utf-8") as handle:
                handle.write(FIXTURE % "\n".join(preamble))
            done = subprocess.run(
                [binary, "-f", path] + args,
                cwd=work,
                env=environment,
                capture_output=True,
                text=True,
                timeout=120,
                check=False,
            )
        return done.stdout if text else done.returncode

    return run


def check(read, runner=make_runner):
    """Read the Makefile, hold the rule, and prove it on this machine."""
    report = []
    errors = []
    text = read(MAKEFILE)
    if text is None:
        return report, [f"{MAKEFILE} is absent"]
    preamble = shell_preamble(text)
    if not any(line.startswith("SHELL") for line in preamble):
        errors.append(f"{MAKEFILE} sets no SHELL. Every recipe then runs under /bin/sh, which has no pipefail")
    guarded = 0
    exempt = 0
    for number, line, marker in recipe_lines(text):
        piped = has_pipeline(line)
        marked = marker.startswith("#") and MARKER in marker
        if piped and GUARD in line:
            guarded += 1
        elif piped and marked:
            exempt += 1
        elif piped:
            errors.append(
                f"{MAKEFILE}:{number} pipes and sets no pipefail. Start the recipe line with 'set -o pipefail;' (F-160)"
            )
        elif marked:
            errors.append(f"{MAKEFILE}:{number} carries a {MARKER} marker and holds no pipeline. Remove the marker")
    report.append(f"{MAKEFILE} pipelines: {guarded} guarded, {exempt} exempt")
    if preamble:
        proof, failures = prove(preamble, runner(preamble))
        report += proof
        errors += failures
    return report, errors


def main():
    def read(path):
        full = os.path.join(ROOT, path)
        if not os.path.exists(full):
            return None
        with open(full, encoding="utf-8") as handle:
            return handle.read()

    report, errors = check(read)
    for line in report:
        print(f"pipefail_check: {line}")
    for error in errors:
        print(f"pipefail_check: {error}")
    return 1 if errors else 0


if __name__ == "__main__":
    sys.exit(main())
