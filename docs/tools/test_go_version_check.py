import pathlib
import tempfile
import unittest

import go_version_check as g


def tree(go="1.27.2", docker="1.27.2", pin="1.27.2"):
    d = pathlib.Path(tempfile.mkdtemp())
    (d / "go").mkdir()
    (d / "go/go.mod").write_text(f"module x\n\ngo {go}\n")
    (d / "docker").mkdir()
    for name in ("api", "worker"):
        (d / f"docker/{name}.Dockerfile").write_text(f"FROM golang:{docker}@sha256:00 AS build\nFROM scratch\n")
    (d / "scripts/live-evals").mkdir(parents=True)
    (d / "scripts/live-evals/pins.sh").write_text(f'PINS=(\n  "go|go version|go version go{pin} darwin/arm64"\n)\n')
    return d


class GoVersionCheckTest(unittest.TestCase):
    def test_one_version_passes(self):
        self.assertEqual(g.faults(tree()), [])

    def test_a_dockerfile_that_differs_fails(self):
        found = g.faults(tree(docker="1.27.0"))
        self.assertEqual(len(found), 2)
        self.assertIn("golang:1.27.0, and go/go.mod reads 1.27.2", found[0])

    def test_a_pin_that_differs_fails(self):
        self.assertEqual(g.faults(tree(pin="1.27.1")),
                         ["scripts/live-evals/pins.sh: go1.27.1, and go/go.mod reads 1.27.2"])

    def test_a_go_line_with_no_patch_fails(self):
        self.assertEqual(g.faults(tree(go="1.27")), ["go/go.mod has no go line of the form 1.X.Y"])


if __name__ == "__main__":
    unittest.main()
