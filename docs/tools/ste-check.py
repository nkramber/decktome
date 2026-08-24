#!/usr/bin/env python3
"""Check Markdown files against the ASD-STE100 rules that a script can test.

Checks: sentence length (25 words, rule 6.3), semicolons (8.1), contractions (4.2),
"-ing" verb forms after helper words (3.5), and paragraphs over six sentences (6.6).
Tables, code blocks, headings, front matter, and URLs are skipped.
Usage: python3 docs/tools/ste-check.py FILE [FILE ...]
"""
import re
import sys

CONTRACTIONS = re.compile(r"\b(\w+n't|\w+'(re|ve|ll|d|m)|it's|let's|that's|there's|what's|here's)\b", re.I)
ING_AFTER = re.compile(r"\b(is|are|was|were|be|been|being|by|of|for|before|after|while|when|without|start|starts|keep|keeps|stop|stops|avoid|avoids|allow|allows)\s+(\w+ing)\b", re.I)
ING_ALLOW = {"thing", "nothing", "something", "anything", "everything", "during", "string", "ring", "king", "bring", "spring", "ping", "wing", "sing", "morning", "evening", "meaning", "setting", "settings", "building", "landing", "ranking", "ordering", "training", "testing", "logging", "streaming", "caching", "matching", "writing", "reading", "learning", "processing", "rendering", "hosting", "routing", "scoring", "tagging", "pricing", "fixing", "sibling", "engineering", "missing", "loading", "warning", "scaling", "sizing"}
MAX_WORDS = 25


def strip_md(line: str) -> str:
    line = re.sub(r"`[^`]*`", "X", line)
    line = re.sub(r"\[([^\]]*)\]\([^)]*\)", r"\1", line)
    line = re.sub(r"https?://\S+", "URL", line)
    line = re.sub(r'"[^"]*"', "QUOTE", line)  # rule 8.6: quoted text counts as one word
    line = re.sub(r"\([^)]*\)", "(X)", line)  # rule 8.5: parentheses count as one word
    line = re.sub(r"[*_>#]+", "", line)
    return line


def sentences(text: str):
    parts = re.split(r"(?<=[.!?])\s+(?=[A-Z0-9\"'(])", text)
    return [p.strip() for p in parts if p.strip()]


def check(path: str):
    findings = []
    in_code = in_front = False
    para = []
    n = 0
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
                if len(para) > 6:
                    findings.append((n, "6.6", f"paragraph has {len(para)} sentences"))
                para = []
                continue
            if line.startswith(">") or line.startswith("**"):
                # a block quote or a bold entry title starts a new paragraph
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
            for s in sentences(txt):
                words = len(s.split())
                if words > MAX_WORDS:
                    findings.append((n, "6.3", f"{words} words: {s[:70]}..."))
                if not re.match(r"^\s*[-*\d]", line):
                    para.append(s)
    return findings


if __name__ == "__main__":
    total = 0
    for p in sys.argv[1:]:
        for n, rule, msg in check(p):
            print(f"{p}:{n}: rule {rule}: {msg}")
            total += 1
    print(f"{total} finding(s)")
    sys.exit(1 if total else 0)
