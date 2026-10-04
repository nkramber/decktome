#!/usr/bin/env python3
"""Symlink-safe file steps of scripts/live-evals.sh in a run folder (D-1141).

A live-eval session writes its run folder under a Seatbelt profile, but
the script reads and writes the same folder outside the profile. A
session can put a symlink at any path of the folder, or put a link in
place of the folder itself. So the script never names a path in the run
folder to the shell. It calls this helper, and the helper opens each
part of a path from a trusted root, with O_NOFOLLOW, through a folder
file descriptor. A link at any part stops the step.

  live_evals_runfs.py get   HOME DECK REL MAX   print one regular file
  live_evals_runfs.py put   HOME DECK REL SRC   replace folder REL with a copy of SRC
  live_evals_runfs.py rm    HOME DECK REL       remove REL, links not followed
  live_evals_runfs.py mkdir HOME DECK REL...    make each folder REL
  live_evals_runfs.py clear HOME DECK          remove the run folder, links not followed

HOME is the folder of the live evals, which no session can write. DECK
is the name of the run folder in it. REL is a relative path in the run
folder. Exit 1 means that the step refused or failed.
"""

import os
import stat
import sys

DIR_FLAGS = os.O_RDONLY | os.O_DIRECTORY | os.O_NOFOLLOW
COPY_LIMIT = 64 << 20


class Refused(Exception):
    pass


def parts(rel):
    out = [p for p in rel.split("/") if p not in ("", ".")]
    if not out or rel.startswith("/") or ".." in out:
        raise Refused(f"bad relative path {rel!r}")
    return out


def name(value):
    if not value or "/" in value or value in (".", ".."):
        raise Refused(f"bad name {value!r}")
    return value


def open_dir(fd, part):
    try:
        return os.open(part, DIR_FLAGS, dir_fd=fd)
    except OSError as err:
        raise Refused(f"{part}: not a plain folder ({err.strerror})") from err


def run_fd(home, deck):
    """The folder of the run, opened with no link at the run name."""
    root = os.open(home, os.O_RDONLY | os.O_DIRECTORY)
    try:
        return open_dir(root, name(deck))
    finally:
        os.close(root)


def walk(fd, names):
    """Open each folder of names under fd. The caller closes the result."""
    cur = os.dup(fd)
    for part in names:
        nxt = open_dir(cur, part)
        os.close(cur)
        cur = nxt
    return cur


def get(home, deck, rel, limit):
    names = parts(rel)
    run = run_fd(home, deck)
    try:
        folder = walk(run, names[:-1])
    finally:
        os.close(run)
    try:
        try:
            fd = os.open(names[-1], os.O_RDONLY | os.O_NOFOLLOW | os.O_NONBLOCK, dir_fd=folder)
        except OSError as err:
            raise Refused(f"{rel}: {err.strerror}") from err
    finally:
        os.close(folder)
    try:
        info = os.fstat(fd)
        if not stat.S_ISREG(info.st_mode):
            raise Refused(f"{rel}: not a regular file")
        if info.st_size > limit:
            raise Refused(f"{rel}: {info.st_size} bytes, over the limit {limit}")
        data = os.read(fd, limit + 1)
    finally:
        os.close(fd)
    sys.stdout.buffer.write(data[:limit])


def remove_in(fd, entry):
    """Remove entry of folder fd. A link is removed, and never followed."""
    try:
        info = os.stat(entry, dir_fd=fd, follow_symlinks=False)
    except FileNotFoundError:
        return
    if not stat.S_ISDIR(info.st_mode):
        os.unlink(entry, dir_fd=fd)
        return
    sub = open_dir(fd, entry)
    try:
        for child in os.listdir(sub):
            remove_in(sub, child)
    finally:
        os.close(sub)
    os.rmdir(entry, dir_fd=fd)


def rm(home, deck, rel):
    names = parts(rel)
    run = run_fd(home, deck)
    try:
        folder = walk(run, names[:-1])
    finally:
        os.close(run)
    try:
        remove_in(folder, names[-1])
    finally:
        os.close(folder)


def clear(home, deck):
    """Remove the run folder. A link at the run name is removed, never followed."""
    root = os.open(home, os.O_RDONLY | os.O_DIRECTORY)
    try:
        remove_in(root, name(deck))
    finally:
        os.close(root)


def copy_into(src, dst_fd, total):
    """Copy the regular files and folders of src into folder dst_fd."""
    for entry in sorted(os.listdir(src)):
        path = os.path.join(src, entry)
        info = os.lstat(path)
        if stat.S_ISDIR(info.st_mode):
            os.mkdir(entry, 0o700, dir_fd=dst_fd)
            sub = open_dir(dst_fd, entry)
            try:
                total = copy_into(path, sub, total)
            finally:
                os.close(sub)
        elif stat.S_ISREG(info.st_mode):
            total += info.st_size
            if total > COPY_LIMIT:
                raise Refused(f"the copy passes {COPY_LIMIT} bytes")
            out = os.open(entry, os.O_WRONLY | os.O_CREAT | os.O_EXCL | os.O_NOFOLLOW, 0o600, dir_fd=dst_fd)
            try:
                with open(path, "rb") as handle:
                    os.write(out, handle.read())
            finally:
                os.close(out)
        else:
            raise Refused(f"{path}: not a file or a folder")
    return total


def put(home, deck, rel, src):
    names = parts(rel)
    run = run_fd(home, deck)
    try:
        folder = walk(run, names[:-1])
    finally:
        os.close(run)
    try:
        remove_in(folder, names[-1])
        os.mkdir(names[-1], 0o700, dir_fd=folder)
        dst = open_dir(folder, names[-1])
        try:
            copy_into(src, dst, 0)
        finally:
            os.close(dst)
    finally:
        os.close(folder)


def mkdir(home, deck, rels):
    run = run_fd(home, deck)
    try:
        for rel in rels:
            cur = os.dup(run)
            try:
                for part in parts(rel):
                    try:
                        os.mkdir(part, 0o700, dir_fd=cur)
                    except FileExistsError:
                        pass
                    nxt = open_dir(cur, part)
                    os.close(cur)
                    cur = nxt
            finally:
                os.close(cur)
    finally:
        os.close(run)


def main(argv):
    if len(argv) < 4:
        print(__doc__, file=sys.stderr)
        return 2
    cmd, home, deck = argv[1], argv[2], argv[3]
    try:
        if cmd == "get" and len(argv) == 6:
            get(home, deck, argv[4], int(argv[5]))
        elif cmd == "put" and len(argv) == 6:
            put(home, deck, argv[4], argv[5])
        elif cmd == "rm" and len(argv) == 5:
            rm(home, deck, argv[4])
        elif cmd == "mkdir" and len(argv) >= 5:
            mkdir(home, deck, argv[4:])
        elif cmd == "clear" and len(argv) == 4:
            clear(home, deck)
        else:
            print(__doc__, file=sys.stderr)
            return 2
    except (Refused, OSError) as err:
        print(f"live_evals_runfs: {cmd} refused: {err}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
