"""Finite SWF2 JPEG asset writer: exact owned pins and pre-output refusals."""
from __future__ import annotations

from hashlib import sha256
import importlib.util
from pathlib import Path
import struct
import subprocess
import sys
import tempfile
import unittest

from PIL import Image
from io import BytesIO

ROOT = Path(__file__).resolve().parents[1]


def loaded(name: str, path: Path):
    spec = importlib.util.spec_from_file_location(name, path)
    assert spec and spec.loader
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


writer = loaded('finite_jpeg_writer_test', ROOT/'tools/swf-jpeg2-encode.py')
oracle = loaded('finite_jpeg_oracle_test', ROOT/'tools/sdk-jpeg-functional-oracle.py')


class BitmapAssetTest(unittest.TestCase):
    def test_three_historical_cpp_exact_payloads_and_actual_decoded_pixels(self):
        for name, case in oracle.FIXTURES.items():
            with self.subTest(name=name):
                rgb = oracle.make_fixture(name)
                w, h = case['width'], case['height']
                payload = writer.jpeg2_payload(rgb, w, h)
                movie = writer.swf2(rgb, w, h)
                self.assertEqual((len(payload), sha256(payload).hexdigest()),
                                 (case['payloadBytes'], case['payload']))
                self.assertEqual((len(movie), sha256(movie).hexdigest()),
                                 (case['swfBytes'], case['swf']))
                self.assertEqual(oracle.swf_payload(movie, w, h), payload)
                self.assertEqual(oracle.decoded_rgb(payload, w, h),
                                 (case['decoded'], case['colors']))
                self.assertEqual(struct.unpack_from('<I', movie, 4)[0], len(movie))

    def test_out_of_profile_rgb_dimensions_and_zero_refuse(self):
        for raw, w, h in ((b'', 1, 1), (b'abc', 0, 1),
                          (b'abc', 33, 1), (b'abc', 1, 33),
                          (b'abc', 2, 1), (b'abc', 1, 2),
                          (b'abc', True, 1), (b'abc', 1, False)):
            with self.subTest(size=len(raw), w=w, h=h):
                with self.assertRaises(ValueError):
                    writer.swf2(raw, w, h)

    def test_marker_layout_rejects_corruption_before_payload(self):
        rgb = oracle.make_fixture('odd')
        w, h = 7, 9
        output = BytesIO()
        Image.frombytes('RGB', (w, h), rgb).save(
            output, 'JPEG', quality=100, subsampling=2, optimize=False,
            progressive=False)
        genuine = output.getvalue()
        self.assertEqual(writer._separate(genuine, w, h),
                         writer.jpeg2_payload(rgb, w, h))
        bad = []
        for offset, bit in ((3, 0x10), (20, 0x01),
                            (genuine.find(b'\xff\xc0') + 7, 0x01)):
            mutated = bytearray(genuine)
            mutated[offset] ^= bit
            bad.append(bytes(mutated))
        bad.extend([genuine[:-1], b'\xff\xd8' + genuine,
                    genuine.replace(b'\xff\xdb', b'\xff\xe1', 1)])
        for variant in bad:
            with self.subTest(sha=sha256(variant).hexdigest()[:12]):
                with self.assertRaises(ValueError):
                    writer._separate(variant, w, h)

    def test_movie_tag_and_jpeg_image_corruption_refused(self):
        payload = writer.jpeg2_payload(oracle.make_fixture('odd'), 7, 9)
        swf = writer.swf2(oracle.make_fixture('odd'), 7, 9)
        self.assertEqual(oracle.decoded_rgb(payload, 7, 9)[0],
                         oracle.FIXTURES['odd']['decoded'])
        for mutated in (swf[:-1], b'GWS' + swf[3:],
                        swf[:4] + b'\0\0\0\0' + swf[8:],
                        swf[:-4] + b'\0\0\0\0'):
            with self.assertRaises(ValueError):
                oracle.swf_payload(mutated, 7, 9)
        bad = bytearray(payload)
        bad[4] = 0
        with self.assertRaises(ValueError):
            oracle.decoded_rgb(bytes(bad), 7, 9)

    def test_cli_same_size_wrong_pin_symlink_and_existing_destination(self):
        with tempfile.TemporaryDirectory() as d:
            temp = Path(d)
            rgb = oracle.make_fixture('odd')
            good = temp/'original.rgb'; good.write_bytes(rgb)
            linked = temp/'linked.rgb'; linked.symlink_to(good)
            partial = temp/'partial.rgb'; partial.write_bytes(rgb[:-1])
            existing = temp/'existing.swf'; existing.write_bytes(b'preserve')
            cli = [sys.executable, str(ROOT/'tools/swf-jpeg2-encode.py')]

            def run(path, dest, w=7, h=9, pin=oracle.FIXTURES['odd']['raw']):
                return subprocess.run([*cli, '--rgb', str(path), '--width', str(w),
                    '--height', str(h), '--expect-sha256', pin,
                    '--output', str(dest)], cwd=ROOT, capture_output=True,
                    text=True, timeout=6)

            for label, inp, w, h, pin in (
                ('pin', good, 7, 9, '0'*64),
                ('partial', partial, 7, 9, sha256(rgb[:-1]).hexdigest()),
                ('symlink', linked, 7, 9, oracle.FIXTURES['odd']['raw']),
                ('dimension', good, 33, 9, oracle.FIXTURES['odd']['raw']),
            ):
                with self.subTest(label=label):
                    dest = temp/(label+'.swf')
                    process = run(inp, dest, w, h, pin)
                    self.assertEqual(process.returncode, 2, process.stderr)
                    self.assertIn('JPEG2 refused before output:', process.stderr)
                    self.assertFalse(dest.exists())
            self.assertEqual(run(good, existing).returncode, 2)
            self.assertEqual(existing.read_bytes(), b'preserve')
            valid = temp/'actual.swf'
            process = run(good, valid)
            self.assertEqual(process.returncode, 0, process.stderr)
            self.assertEqual(sha256(valid.read_bytes()).hexdigest(),
                             oracle.FIXTURES['odd']['swf'])


if __name__ == '__main__':
    unittest.main()
