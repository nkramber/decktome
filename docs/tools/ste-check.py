#!/usr/bin/env python3
"""Check Markdown files against the ASD-STE100 rules that a script can test.

Checks: sentence length (25 words, rule 6.3), semicolons (8.1), contractions (4.2),
"-ing" verb forms after helper words (3.5), and paragraphs over six sentences (6.6).
Tables, code blocks, headings, front matter, and URLs are skipped. A
comma list of technical names (rule 4.3, 8.6) is not held to the length
rule.
Usage: python3 docs/tools/ste-check.py FILE [FILE ...]
"""
import re
import sys

CONTRACTIONS = re.compile(r"\b(\w+n't|\w+'(re|ve|ll|d|m)|it's|let's|that's|there's|what's|here's)\b", re.I)
ING_AFTER = re.compile(r"\b(is|are|was|were|be|been|being|by|of|for|before|after|while|when|without|start|starts|keep|keeps|stop|stops|avoid|avoids|allow|allows)\s+(\w+ing)\b", re.I)
ING_ALLOW = {"thing", "nothing", "something", "anything", "everything", "during", "string", "ring", "king", "bring", "spring", "ping", "wing", "sing", "morning", "evening", "meaning", "setting", "settings", "building", "landing", "ranking", "ordering", "training", "testing", "logging", "streaming", "caching", "matching", "writing", "reading", "learning", "processing", "rendering", "hosting", "routing", "scoring", "tagging", "pricing", "fixing", "sibling", "engineering", "missing", "loading", "warning", "scaling", "sizing"}
MAX_WORDS = 25


def strip_md(line: str) -> str:
    # Curly quotes count like straight ones (rule 8.6), and a curly
    # apostrophe hides a contraction from the check.
    line = line.replace("\u2019", "'").replace("\u2018", "'").replace("\u201c", '"').replace("\u201d", '"')
    line = re.sub(r"`[^`]*`", "X", line)
    line = re.sub(r"\[([^\]]*)\]\([^)]*\)", r"\1", line)
    line = re.sub(r"https?://\S+", "URL", line)
    line = re.sub(r'"[^"]*"', "QUOTE", line)  # rule 8.6: quoted text counts as one word
    line = re.sub(r"\([^)]*\)", "(X)", line)  # rule 8.5: parentheses count as one word
    line = re.sub(r"[*_>#]+", "", line)
    return line


def is_name_list(s: str) -> bool:
    """A run of short comma-separated items is a list of names, not a sentence."""
    items = [i.strip() for i in s.split(",")]
    return len(items) >= 6 and sum(len(i.split()) for i in items) / len(items) <= 3


def sentences(text: str):
    parts = re.split(r"(?<=[.!?])\s+(?=[A-Z0-9\"'(])", text)
    return [p.strip() for p in parts if p.strip()]


def check(path: str):
    findings = []
    in_code = in_front = False
    para = []
    held = []  # (line number, text) of the paragraph so far, for wrapped sentences
    n = 0

    def flush_held():
        # A sentence may span hard-wrapped lines. Join the held lines and
        # split into sentences once, so a wrap is not a sentence end.
        if not held:
            return
        first = held[0][0]
        joined = " ".join(t for _, t in held)
        for s in sentences(joined):
            words = len(s.split())
            if words > MAX_WORDS and not is_name_list(s):
                findings.append((first, "6.3", f"{words} words: {s[:70]}..."))
            para.append(s)
        held.clear()
    with open(path, encoding="utf-8") as fh:
        for n, raw in enumerate(fh, 1):
            line = raw.rstrip("\n")
            if n == 1 and line.strip() == "---":
                in_front = True
                continue
            if in_front:
                if line.strip() == "---":
                    in_front = False
                continue
            if line.strip().startswith("```"):
                in_code = not in_code
                continue
            if in_code or line.strip().startswith("|") or line.strip().startswith("#"):
                continue
            if not line.strip() or line.strip() == ">":
                flush_held()
                if len(para) > 6:
                    findings.append((n, "6.6", f"paragraph has {len(para)} sentences"))
                para = []
                continue
            if line.startswith(">") or line.startswith("**"):
                # a block quote or a bold entry title starts a new paragraph
                flush_held()
                if len(para) > 6:
                    findings.append((n, "6.6", f"paragraph has {len(para)} sentences"))
                para = []
            txt = strip_md(line)
            if ";" in txt:
                findings.append((n, "8.1", "semicolon"))
            for m in CONTRACTIONS.finditer(txt):
                findings.append((n, "4.2", f"contraction '{m.group(0)}'"))
            for m in ING_AFTER.finditer(txt):
                w = m.group(2).lower()
                if w not in ING_ALLOW:
                    findings.append((n, "3.5", f"-ing form '{m.group(0)}'"))
            if re.match(r"^\s*[-*\d]", line):
                # A list item is one unit. It is not part of the paragraph count.
                flush_held()
                for s in sentences(txt):
                    words = len(s.split())
                    if words > MAX_WORDS and not is_name_list(s):
                        findings.append((n, "6.3", f"{words} words: {s[:70]}..."))
                continue
            held.append((n, txt))
    # The last paragraph of a file ends with no blank line after it.
    flush_held()
    if len(para) > 6:
        findings.append((n, "6.6", f"paragraph has {len(para)} sentences"))
    return findings


if __name__ == "__main__":
    total = 0
    for p in sys.argv[1:]:
        for n, rule, msg in check(p):
            print(f"{p}:{n}: rule {rule}: {msg}")
            total += 1
    print(f"{total} finding(s)")
    sys.exit(1 if total else 0)
