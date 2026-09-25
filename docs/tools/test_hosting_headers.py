"""Tests of the browser security headers of firebase.json (D-923)."""
import json
import os
import unittest

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


def headers_of(source):
    with open(os.path.join(ROOT, "firebase.json"), encoding="utf-8") as handle:
        config = json.load(handle)
    for rule in config["hosting"]["headers"]:
        if rule["source"] == source:
            return {h["key"]: h["value"] for h in rule["headers"]}
    return {}


class HostingHeaders(unittest.TestCase):
    def test_every_path_sends_the_security_headers(self):
        headers = headers_of("**")
        self.assertEqual(headers.get("X-Content-Type-Options"), "nosniff")
        self.assertEqual(headers.get("Referrer-Policy"), "strict-origin-when-cross-origin")
        self.assertEqual(headers.get("Cache-Control"), "no-cache")

    def test_no_frame_rule_blocks_the_auth_iframe(self):
        # The auth domain frames its helper page into the app from another
        # origin, and this site serves both. A frame rule waits for a check
        # on a preview channel (D-923).
        headers = headers_of("**")
        self.assertNotIn("X-Frame-Options", headers)
        self.assertNotIn("Content-Security-Policy", headers)


if __name__ == "__main__":
    unittest.main()
