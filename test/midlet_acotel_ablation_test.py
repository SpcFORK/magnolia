"""Format/policy unit tests; no archived JAR or uploaded JAD is committed."""
from __future__ import annotations

import importlib.util
from pathlib import Path
import unittest

SCRIPT = Path(__file__).resolve().parent.parent / "tools/midlet-acotel-ablation.py"
SPEC = importlib.util.spec_from_file_location("midlet_acotel_ablation", SCRIPT)
assert SPEC and SPEC.loader
module = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(module)


class MidletAcotelIntakeTests(unittest.TestCase):
    def test_archive_urls_do_not_follow_untrusted_jad_url(self) -> None:
        for kind in ("jad", "jar"):
            url = module.archive_url(kind)
            self.assertTrue(url.startswith("https://archive.org/download/"))
            self.assertIn("ACO_XmasGifts_s." + kind, url)
            self.assertNotIn("124.1.1.52", url)

    def test_jad_properties_are_metadata_not_classfile(self) -> None:
        raw = (b"MicroEdition-Configuration: CLDC-1.0\r\n"
               b"MicroEdition-Profile: MIDP-1.0\r\n"
               b"MIDlet-1: Xmas, Xmas.png, Xmas\r\n")
        props = module.attributes(raw, "test JAD")
        self.assertEqual(props["MicroEdition-Profile"], "MIDP-1.0")
        self.assertEqual(props["MIDlet-1"], "Xmas, Xmas.png, Xmas")
        self.assertFalse(raw.startswith(b"\xca\xfe\xba\xbe"))

    def test_manifest_continuations_and_blank_terminator(self) -> None:
        attrs = module.attributes(b"MIDlet-Name: Acotel\r\n continuation\r\n\r\n", "manifest")
        self.assertEqual(attrs, {"MIDlet-Name": "Acotelcontinuation"})

    def test_duplicate_or_orphan_descriptor_attribute_refuses(self) -> None:
        bad = (b"MIDlet-Name: one\r\nMIDlet-Name: two\r\n",
               b" orphan\r\nMIDlet-Name: one\r\n",
               b"MIDlet-Name=no-colon\r\n",
               b"MIDlet-Name: bad\x00value\r\n")
        for raw in bad:
            with self.subTest(raw=raw), self.assertRaises(ValueError):
                module.attributes(raw, "test")

    def test_pins_reject_wrong_bytes_before_parser_or_vm(self) -> None:
        for kind in ("jad", "jar"):
            with self.subTest(kind=kind), self.assertRaises(ValueError):
                module.pinned(b"\xca\xfe\xba\xbe", kind)

    def test_ablation_anchor_is_unique_and_isolated(self) -> None:
        for _, (path, anchor, replacement) in module.VARIANTS.items():
            source = (module.ROOT / "lib" / path).read_text()
            self.assertIn(path, {"applet-bytecode.oak", "jvm-vm-libs.oak"})
            self.assertNotEqual(anchor, replacement)
            self.assertEqual(source.count(anchor), 1)


if __name__ == "__main__":
    unittest.main()
