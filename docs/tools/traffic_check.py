"""Check that the newest ready revision of a Cloud Run service serves all traffic (REV-013).

A rollback by traffic pin sends 100 percent to an old revision. A later
deploy then makes a new revision that takes no traffic, and the health
check reads the old revision, so the build reads SUCCESS. The deploy step
of `cloudbuild/api.yaml` pipes `gcloud run services describe --format=json`
into this script, and a non-zero exit fails the build.

Run: gcloud run services describe mtg-api --region R --format=json | python3 docs/tools/traffic_check.py
"""
import json
import sys


def served(service):
    """Return the newest ready revision and the percent of traffic it serves."""
    status = service.get("status", {})
    latest = status.get("latestReadyRevisionName", "")
    percent = 0
    for t in status.get("traffic", []):
        if t.get("revisionName") == latest or (t.get("latestRevision") and not t.get("revisionName")):
            percent += int(t.get("percent", 0))
    return latest, percent


def main(stdin=sys.stdin, out=sys.stdout):
    latest, percent = served(json.load(stdin))
    if not latest:
        print("traffic_check: the service names no ready revision", file=out)
        return 1
    if percent != 100:
        print(f"traffic_check: {latest} serves {percent} percent of the traffic, not 100."
              " A traffic pin holds the old revision. Run"
              " `gcloud run services update-traffic mtg-api --to-latest` after the cause is fixed.", file=out)
        return 1
    print(f"traffic_check: {latest} serves 100 percent", file=out)
    return 0


if __name__ == "__main__":
    sys.exit(main())
