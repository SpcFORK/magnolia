"""Actual FWS URL-rewrite behavior and fail-closed boundaries, no archived code.

The tiny independent SWF fixtures are synthesized, not downloaded; the pinned
SDK demo oracle/output hashes in the separate test are from a bounded reference
run of the old C++ utility and can be checked without shipping its binary.
"""
from __future__ import annotations

from hashlib import sha256
import importlib.util
import json
from pathlib import Path
import struct
import subprocess
import sys
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / 'tools/swf-urlsubst.py'
SPEC = importlib.util.spec_from_file_location('swf_urlsubst', SCRIPT)
assert SPEC and SPEC.loader
# dataclasses looks up its containing module during annotation handling.
sys.modules[SPEC.name] = urlsubst = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(urlsubst)


def tag(code: int, data: bytes = b'') -> bytes:
    if len(data) < 63:
        return struct.pack('<H', code << 6 | len(data)) + data
    return struct.pack('<HI', code << 6 | 63, len(data)) + data


def action(url: bytes, window: bytes = b'_self') -> bytes:
    data = url + b'\0' + window + b'\0'
    return b'\x83' + struct.pack('<H', len(data)) + data + b'\0'


def movie(tags: bytes, *, version: int = 5) -> bytes:
    # 1-bit signed zero RECT, 20 fps (8.8), one ShowFrame and End.
    body = b'\x08\0' + struct.pack('<HH', 20 * 256, 1)
    body += tags + tag(1) + tag(0)
    return b'FWS' + bytes((version,)) + struct.pack('<I', 8 + len(body)) + body


def button2(actions: bytes, offset: int = 3) -> bytes:
    # No button records; offset measured from its own field (index 3).
    return struct.pack('<HBH', 99, 0, offset) + b'\0' + struct.pack('<HH', 0, 1) + actions


def synthetic_movie() -> bytes:
    direct = tag(12, action(b'direct:$A$', b'$B$'))
    button1 = tag(7, struct.pack('<H', 5) + b'\0' + action(b'button:$A$'))
    button_v2 = tag(34, button2(action(b'button2:$A$')))
    sprite_stream = tag(12, action(b'sprite:$A$')) + tag(1) + tag(0)
    sprite = tag(39, struct.pack('<HH', 9, 1) + sprite_stream)
    # Opaque tag with a long header despite short payload; historical SDK
    # canonicalizes this tag header without touching its payload.
    opaque = struct.pack('<HI', 9 << 6 | 63, 3) + b'\x01\x02\x03'
    return movie(opaque + direct + button1 + button_v2 + sprite)


class URLSubstitutionTests(unittest.TestCase):
    def test_actual_geturl_button_and_sprite_changes_output(self):
        raw = synthetic_movie()
        result, counts = urlsubst.rewrite_swf(raw,
                          urlsubst.variables([('a', 'value'), ('B', 'window')]))
        self.assertNotEqual(result, raw)
        self.assertEqual(len(result), int.from_bytes(result[4:8], 'little'))
        self.assertEqual((counts.get_urls, counts.substituted_strings,
                          counts.dropped_suffixes), (4, 5, 0))
        # Exact independently compiled SDK C++ oracle on this synthetic input.
        self.assertEqual((len(raw), sha256(raw).hexdigest()),
                         (141, '37a11b6eeff1d19293462aa5afacdded4c53522cb1c15560400195e47e0d539a'))
        self.assertEqual((len(result), sha256(result).hexdigest()),
                         (168, '9360dcb4034f16e0a397f42c289bc67732f2a3b8052c52a2858aa7162ef49601'))
        for expected in (b'direct:value\0window\0', b'button:value',
                         b'button2:value', b'sprite:value'):
            self.assertIn(expected, result)
        self.assertNotIn(b'$A$', result)
        self.assertNotIn(b'$B$', result)
        self.assertEqual(urlsubst.inventory.inspect_swf(result)['tagCounts'],
                         urlsubst.inventory.inspect_swf(raw)['tagCounts'])
        self.assertEqual(result.find(b'\x43\x02\x01\x02\x03'), 14)

    def test_two_conditional_button_actions_relocate_offsets(self):
        first = action(b'first-$A$')
        second = action(b'second-$A$', b'$B$')
        records = (struct.pack('<HH', 4 + len(first), 1) + first +
                   struct.pack('<HH', 0, 2) + second)
        source = movie(tag(34, struct.pack('<HBH', 77, 0, 3) + b'\0' + records))
        result, count = urlsubst.rewrite_swf(source, urlsubst.variables([
            ('A', 'much-longer'), ('B', 'frame')]))
        self.assertEqual((count.get_urls, count.substituted_strings), (2, 3))
        self.assertIn(b'first-much-longer\0_self\0', result)
        self.assertIn(b'second-much-longer\0frame\0', result)
        # A separate two-condition executable C++ SDK run produced these
        # *whole SWF* bytes; the test never invokes or vendors that binary.
        self.assertEqual((len(source), sha256(source).hexdigest()),
                         (73, '54cb21eadc63cbe44dc257913145cadf5a041e227eb63a665691f1f0ee2d6442'))
        self.assertEqual((len(result), sha256(result).hexdigest()),
                         (95, 'a6dfb9f0c49975a562e27f8d1cfdd27c92afdd551a174a1f10eb661ab09da709'))

    def test_historical_sdk_suffix_behavior_is_explicit(self):
        raw = movie(tag(12, action(b'prefix$A$suffix')))
        result, counts = urlsubst.rewrite_swf(raw, urlsubst.variables([('A', 'X')]))
        self.assertIn(b'prefixX\0_self\0', result)
        self.assertNotIn(b'suffix', result)
        self.assertEqual(counts.dropped_suffixes, 1)

    def test_bad_variables_fail_instead_of_erasing_urls(self):
        for pairs in ([], [('A', '')], [('A', 'x'), ('a', 'y')],
                      [('A', '$B$')], [('A', 'hi\n')], [('A-', 'x')],
                      [('A', 'x' * 1025)]):
            with self.subTest(pairs=pairs), self.assertRaises(ValueError):
                urlsubst.variables(pairs)
        for string in (b'$MISSING$', b'half$A', b'$A$$B$', b'$-A$'):
            with self.subTest(string=string), self.assertRaises(ValueError):
                urlsubst.rewrite_swf(movie(tag(12, action(string))),
                                     urlsubst.variables([('A', 'x'), ('B', 'y')]))

    def test_action_and_tag_boundary_refusals(self):
        bad = [
            tag(12, b'\x83\x03\x00x\0\0'),  # missing action End
            tag(12, b'\x83\x08\x00a\0b\0\0'),  # payload exceeds tag
            tag(12, b'\x83\x04\x00a\0b\0\0\0'),  # bytes after End
            tag(12, b'\x83\x06\x00a\0b\0x\0'),  # extra GetURL data
            tag(34, button2(action(b'$A$'), offset=4)),
            tag(7, struct.pack('<H', 1) + b'\x01\0'),  # short record
            tag(26, b'\x02'),  # hidden clip actions cannot be copied as safe
            tag(59, b'\0\0'),
            tag(70, b'\0'),
        ]
        for tags in bad:
            with self.subTest(tags=tags.hex()), self.assertRaises(ValueError):
                urlsubst.rewrite_swf(movie(tags), urlsubst.variables([('A', 'X')]))

    def test_changed_url_refuses_branch_and_dynamic_url(self):
        target = action(b'$A$')[:-1]
        branches = [b'\x99\x02\0\0\0', b'\x9d\x02\0\0\0',
                    b'\x8e\x00\0', b'\x96\x00\0']
        for code in branches:
            data = target + code + b'\0'
            with self.subTest(code=code.hex()), self.assertRaisesRegex(
                    ValueError, 'unsupported control-flow'):
                urlsubst.rewrite_swf(movie(tag(12, data)),
                                     urlsubst.variables([('A', 'extended')]))
        with self.assertRaisesRegex(ValueError, 'dynamic GetURL2'):
            urlsubst.rewrite_swf(movie(tag(12, b'\x9a\x01\0\0\0')),
                                 urlsubst.variables([('A', 'X')]))

    def test_swf_format_and_resource_bounds_refuse(self):
        base = synthetic_movie()
        for bad in (b'CWS' + base[3:], b'ZWS' + base[3:],
                    base[:3] + b'\x08' + base[4:], base + b'\0',
                    base[:4] + b'\xff\xff\xff\xff' + base[8:]):
            with self.subTest(start=bad[:8].hex()), self.assertRaises(ValueError):
                urlsubst.rewrite_swf(bad, urlsubst.variables([('A', 'X')]))
        with self.assertRaises(ValueError):
            urlsubst.rewrite_swf(base + b'\0' * urlsubst.MAX_INPUT,
                                 urlsubst.variables([('A', 'X')]))

    def test_cli_no_output_on_pin_error_and_malformed_bytes(self):
        with tempfile.TemporaryDirectory() as tmp:
            src, dst = Path(tmp) / 'a.swf', Path(tmp) / 'b.swf'
            src.write_bytes(synthetic_movie())
            # A pre-existing output must remain byte-exact on failure.
            dst.write_bytes(b'sentinel')
            cmd = [sys.executable, str(SCRIPT), '-i', str(src), '-o', str(dst),
                   '-v', 'A', 'x', '-v', 'B', 'y']
            denied = subprocess.run(cmd + ['--expect-sha256', '0' * 64],
                                     capture_output=True, timeout=4)
            self.assertEqual(denied.returncode, 2)
            self.assertEqual(dst.read_bytes(), b'sentinel')
            src.write_bytes(movie(tag(12, b'\x83\x01\0x\0')))
            denied = subprocess.run(cmd, capture_output=True, timeout=4)
            self.assertEqual(denied.returncode, 2)
            self.assertEqual(dst.read_bytes(), b'sentinel')
            src.write_bytes(synthetic_movie())
            passed = subprocess.run(cmd + ['--expect-sha256', sha256(src.read_bytes()).hexdigest()],
                                    capture_output=True, timeout=4)
            self.assertEqual(passed.returncode, 0, passed.stderr.decode())
            report = json.loads(passed.stdout)
            self.assertEqual((report['urlActions'], report['substitutedStrings']), (4, 5))
            self.assertEqual(report['outputSha256'], sha256(dst.read_bytes()).hexdigest())

    def test_pinned_historic_sdk_demo_if_external_original_is_present(self):
        original = Path('/var/tmp/magnolia-sdk-src.swf')
        if not original.is_file():
            self.skipTest('outside-checkout SDK demo SWF not present')
        raw = original.read_bytes()
        self.assertEqual((len(raw), sha256(raw).hexdigest()),
                         (599, 'eaff967749b061e9ae8505f2582bf784c6274786d5764d4e9b3582499b625f24'))
        result, counts = urlsubst.rewrite_swf(raw, urlsubst.variables([
            ('A', '57859;0;0;120316'), ('B', 'http://www.macromedia.com')]))
        self.assertEqual((len(result), sha256(result).hexdigest()),
                         (551, '7ca1878c4989357dbb9cb50b683e1dd95a72d84bd6c7d07ebfd97f9fab3a4dc6'))
        self.assertEqual((counts.get_urls, counts.substituted_strings,
                          counts.dropped_suffixes), (3, 3, 2))
        reference = Path('/var/tmp/magnolia-sdk-dst.swf')
        if reference.is_file():
            self.assertEqual(result, reference.read_bytes())


if __name__ == '__main__':
    unittest.main()
