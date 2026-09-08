#!/usr/bin/env bash
# Read one chat session of the deployed project, for debugging.
#
# The owner operates the project and answers for every session in it, so
# this reads the session of any user (D-591). It prints the fields that
# explain what the agent did: the card pool, the slots, and each turn.
#
# Usage:
#   scripts/read-session.sh <session-id> [uid]
#   SESSION_PROJECT=other-project scripts/read-session.sh <session-id>
#   RAW=1 scripts/read-session.sh <session-id>
#
# With no uid it reads the user list and looks for the session under each
# one. RAW=1 prints the whole document. The credentials come from
# `gcloud auth login`, and `use_decktome` puts the shell on the right
# project.
#
# CAUTION: the output holds what a reader wrote. Keep it off any shared
# page, and put no part of it in an issue or a pull request.
set -uo pipefail

session="${1:-}"
uid="${2:-}"
# The project comes from SESSION_PROJECT, then from the active gcloud
# configuration, and it never comes from PROJECT_ID. A shell that works
# on more than one project exports PROJECT_ID for another one, and this
# script read `wallabee-dev` from it on 2026-09-08.
project="${SESSION_PROJECT:-$(gcloud config get-value project 2>/dev/null)}"
[ -z "$project" ] && project="decktome-prod"

if [ -z "$session" ]; then
  echo "usage: scripts/read-session.sh <session-id> [uid]" >&2
  exit 2
fi

token=$(gcloud auth print-access-token 2>/dev/null)
if [ -z "$token" ]; then
  echo "no access token. Run: gcloud auth login" >&2
  exit 1
fi

# The project is named on every run, so a read of the wrong one is never
# a silent one.
echo "project      ${project}" >&2

base="https://firestore.googleapis.com/v1/projects/${project}/databases/(default)/documents"

get() { curl -s -H "Authorization: Bearer ${token}" "$1"; }

# users prints every user id of the project, one per line.
#
# showMissing is what makes this work. The app writes
# users/<uid>/sessions/<id> and never gives users/<uid> a field of its
# own, so that parent is a missing document. Firestore leaves a missing
# document out of a plain list, and the first version of this script
# read an empty list and found nothing.
users() {
  local page="" out
  while :; do
    out=$(get "${base}/users?pageSize=300&showMissing=true${page}")
    printf '%s' "$out" | python3 -c '
import json, sys
try:
    d = json.load(sys.stdin)
except ValueError:
    print("read-session: the user list is not JSON", file=sys.stderr)
    sys.exit(0)
# An error answer is JSON too, and it holds no documents. Without this
# line a refusal reads the same as a project with no user.
if "error" in d:
    e = d["error"]
    print("read-session: {} {}".format(e.get("status", ""), e.get("message", "")), file=sys.stderr)
    sys.exit(0)
for doc in d.get("documents", []):
    print(doc["name"].rsplit("/", 1)[-1])
'
    page=$(printf '%s' "$out" | python3 -c '
import json, sys
t = json.load(sys.stdin).get("nextPageToken")
print("&pageToken=" + t if t else "")
')
    [ -z "$page" ] && break
  done
}

# show prints the readable part of one session document. It exits 1 when
# the answer holds no document, so the caller can try the next user.
show() {
  python3 -c '
import json, sys

raw = sys.argv[1] == "1"
doc = json.load(sys.stdin)
if "fields" not in doc:
    sys.exit(1)


def val(v):
    # bytesValue arrives base64 in the REST answer, and session_gz is the
    # one field that uses it.
    for k in ("stringValue", "booleanValue", "timestampValue", "nullValue", "bytesValue"):
        if k in v:
            return v[k]
    if "integerValue" in v:
        return int(v["integerValue"])
    if "doubleValue" in v:
        return v["doubleValue"]
    if "mapValue" in v:
        return {k: val(x) for k, x in v["mapValue"].get("fields", {}).items()}
    if "arrayValue" in v:
        return [val(x) for x in v["arrayValue"].get("values", [])]
    return v


f = {k: val(v) for k, v in doc["fields"].items()}

# The document holds the index fields in snake_case, and the session
# itself in session_gz: gzip protojson of the Session message. The slots,
# the pool rule, and every turn live in there and nowhere else. The first
# version of this script read camelCase names off the document and
# reported an empty session that was not empty.
inner = {}
gz = f.get("session_gz")
if isinstance(gz, str):
    import base64, gzip
    try:
        inner = json.loads(gzip.decompress(base64.b64decode(gz)).decode("utf-8"))
    except Exception as err:
        print("session_gz did not decode: {}".format(err), file=sys.stderr)

if raw:
    print(json.dumps({"document": f, "session": inner}, indent=2, default=str))
    sys.exit(0)


def show_line(label, value):
    print("{:<14}{}".format(label, value))


show_line("session", doc["name"].rsplit("/", 1)[-1])
show_line("user", doc["name"].split("/users/")[1].split("/")[0])
for key in ("collection_id", "status", "created_at", "updated_at", "version", "schema_version"):
    if key in f:
        show_line(key, f[key] if f[key] != "" else "(empty)")
if "collection_id" not in f:
    show_line("collection_id", "(no field: the session names no collection)")

slots = inner.get("slots")
if isinstance(slots, dict):
    print("slots")
    for k in sorted(slots):
        print("  {:<22}{}".format(k, slots[k]))
else:
    print("slots         (none)")

if inner.get("deckIds"):
    show_line("deckIds", ", ".join(inner["deckIds"]))

def clip(s, n):
    s = str(s or "").replace("\n", " ").strip()
    return s[:n] + "..." if len(s) > n else s


# The Turn message names user_message and agent_message, so protojson
# writes userMessage and agentMessage. The first version of this script
# looked for "message" and printed an empty line for every turn.
turns = inner.get("turns") or []
show_line("turns", len(turns))
for i, t in enumerate(turns, 1):
    if not isinstance(t, dict):
        continue
    print("  {:>2}. user   {}".format(i, clip(t.get("userMessage"), 300)))
    if t.get("agentMessage"):
        print("      agent  {}".format(clip(t.get("agentMessage"), 300)))
    for q in t.get("questions") or []:
        print("      asked  {}".format(clip(q.get("prompt") or q.get("text") or q.get("key"), 160)))
    for a in t.get("answers") or []:
        print("      answer {}".format(clip(json.dumps(a, default=str), 160)))
' "${RAW:-0}"
}

try_one() {
  get "${base}/users/${1}/sessions/${session}" | show
}

if [ -n "$uid" ]; then
  if try_one "$uid"; then exit 0; fi
  echo "no session ${session} under user ${uid} in ${project}" >&2
  exit 1
fi

found=1
scanned=0
while read -r u; do
  [ -z "$u" ] && continue
  scanned=$((scanned + 1))
  if try_one "$u"; then
    found=0
    break
  fi
done < <(users)

if [ "$found" -ne 0 ]; then
  # The count separates a session that is not there from a user list that
  # came back empty. The two failures read the same without it.
  echo "no session ${session} in ${project}, over ${scanned} user(s)" >&2
  if [ "$scanned" -eq 0 ]; then
    echo "the user list is empty. Check the account: gcloud config get-value account" >&2
  fi
  exit 1
fi
