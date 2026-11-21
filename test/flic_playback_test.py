"""Owned FLI/FLC codec, policy, output and preflight tests.

Real external originals and an independently compiled Go reference are exercised
by tools/flic-functional-oracle.py; no third-party media is bundled here.
"""
from __future__ import annotations

import importlib.util
from io import BytesIO
from pathlib import Path
import struct
import tempfile
import unittest

from PIL import Image

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location('flic_playback', ROOT / 'tools/flic-playback.py')
assert SPEC and SPEC.loader
import sys
flic = importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name] = flic
SPEC.loader.exec_module(flic)


def chunk(code: int, data: bytes = b'') -> bytes:
    return struct.pack('<IH', len(data) + 6, code) + data


def frame(*chunks: bytes) -> bytes:
    body = struct.pack('<IHHHHHH', 16 + sum(map(len, chunks)),
                       flic.FRAME, len(chunks), 0, 0, 0, 0)
    return body + b''.join(chunks)


def movie(width: int, height: int, frames: list[bytes], *,
          declared: int | None = None, magic: int = 0xAF12,
          prefix: bytes = b'') -> bytes:
    header = bytearray(128)
    n = len(frames) if declared is None else declared
    raw = bytes(header) + prefix + b''.join(frames)
    struct.pack_into('<IHHHHHHI', header, 0, len(raw), magic, n,
                     width, height, 8, 3, 50)
    return bytes(header) + raw[128:]


class FlicPlaybackTest(unittest.TestCase):
    def setUp(self) -> None:
        self.palette = chunk(4, struct.pack('<HBB', 1, 0, 5) +
            bytes((0, 0, 0, 255, 0, 0, 0, 255, 0,
                   0, 0, 255, 255, 255, 0)))
        self.brun = chunk(15, b'\x01\xfd\x00\x01\x02' + b'\x01\x03\x03')
        self.first = frame(self.palette, self.brun)

    def test_real_rendering_byte_runs_delta_flic_black_and_copy(self) -> None:
        delta = chunk(7, struct.pack('<HHHBB', 1, 0x8004, 1, 0, 1) + b'\x02\x02')
        raw = movie(3, 2, [self.first, frame(delta), frame(chunk(13)),
                           frame(chunk(16, b'\x04\x00\x01\x02\x03\x04'))])
        parsed = flic.decode(raw)
        self.assertEqual((parsed.header.width, parsed.header.height,
                          parsed.header.delay_ms, parsed.ring), (3, 2, 50, False))
        self.assertEqual([idx for idx, _ in parsed.shown], [
            b'\x00\x01\x02\x03\x03\x03',
            b'\x02\x02\x04\x03\x03\x03',
            b'\x00' * 6,
            b'\x04\x00\x01\x02\x03\x04'])
        self.assertEqual(flic.rgb(*parsed.shown[0]), bytes((
            0, 0, 0, 255, 0, 0, 0, 255, 0,
            0, 0, 255, 0, 0, 255, 0, 0, 255)))

    def test_fli_delta_lc_with_literal_and_repeated_pixels(self) -> None:
        head = frame(chunk(16, bytes((1, 1, 1, 1, 1, 1, 1, 1))))
        delta = chunk(12, struct.pack('<HHB', 0, 2, 1) +
                      b'\x01\x02\x02\x03' +
                      b'\x01\x00\xff\x04')
        parsed = flic.decode(movie(4, 2, [head, frame(delta)]))
        self.assertEqual(parsed.shown[1][0], b'\x01\x02\x03\x01' +
                         b'\x04\x01\x01\x01')

    def test_brun_clips_within_row_not_the_next_scanline(self) -> None:
        overshoot = chunk(15, b'\x01\x0a\x01' + b'\x01\xfc\x02\x03\x04\x05')
        parsed = flic.decode(movie(4, 2, [frame(overshoot)]))
        self.assertEqual(parsed.shown[0][0], b'\x01' * 4 + b'\x02\x03\x04\x05')

    def test_palette_64_full_range_and_trailer_is_counted(self) -> None:
        palette64 = chunk(11, struct.pack('<HBB', 1, 0, 1) + b'\x3f\x20\x00\x0a')
        raw = movie(320, 200, [frame(palette64)], magic=0xAF11)
        decoded = flic.decode(raw)
        self.assertEqual(decoded.shown[0][1][:3], bytes((255, 129, 0)))
        self.assertEqual(decoded.bounded_trailer_bytes, 1)

    def test_true_ring_vs_real_extra(self) -> None:
        first = frame(chunk(16, b'\x00\x01'))
        ring = flic.decode(movie(2, 1, [first, frame(chunk(16, b'\x00\x01'))], declared=1))
        self.assertTrue(ring.ring)
        self.assertEqual(len(ring.shown), 1)
        real = flic.decode(movie(2, 1, [first, frame(chunk(16, b'\x01\x00'))], declared=1))
        self.assertFalse(real.ring)
        self.assertTrue(real.extra_not_ring)
        self.assertEqual(len(real.shown), 2)

    def test_offline_html_and_png_are_actual_decoded_frames(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            source = root / 'owned.flc'
            source.write_bytes(movie(3, 2, [self.first]))
            outcome = flic.run(['--input', str(source), '--html', str(root / 'viewer.html'),
                                '--png-dir', str(root / 'frames'),
                                '--report', str(root / 'meta.json')])
            self.assertEqual(outcome['frames'][0]['rgbSHA256'],
                             flic.sha(flic.rgb(*flic.decode(source.read_bytes()).shown[0])))
            html = (root / 'viewer.html').read_text()
            self.assertEqual(html.count('<script>'), 1)
            self.assertIn('data:image/png;base64,', html)
            self.assertNotIn('src="http', html)
            self.assertNotIn('<script src=', html)
            self.assertIn('50.000 ms/frame', html)
            with Image.open(root / 'frames/000.png') as image:
                self.assertEqual(image.size, (3, 2))
                self.assertEqual(image.convert('RGB').tobytes(),
                                 flic.rgb(*flic.decode(source.read_bytes()).shown[0]))

    def test_unsupported_or_corrupt_refuses_before_any_output(self) -> None:
        valid = movie(3, 2, [self.first])
        examples = {
            'bad magic': bytearray(valid),
            'bad length': bytearray(valid),
            'unknown subchunk': movie(3, 2, [frame(chunk(99))]),
            'unknown top': movie(3, 2, [self.first], prefix=
                                 struct.pack('<IH', 16, 0xF200) + b'\x00' * 10),
            'truncated frame': valid[:-1],
            'frame count': movie(3, 2, [self.first], declared=3),
            'palette overflow': movie(3, 2, [frame(chunk(4, b'\x01\x00\xff\x02\x00'*2))]),
            'palette 64 range': movie(320, 200, [frame(chunk(11, b'\x01\x00\x00\x01\x40\x00\x00'))], magic=0xAF11),
            'brun zero step': movie(3, 2, [frame(chunk(15, b'\x01\x00'))]),
            'flc undefined opcode': movie(3, 2, [frame(chunk(7, b'\x01\x00\x00\x40'))]),
            'fli delta line overflow': movie(3, 2, [frame(chunk(12, b'\x01\x00\x03\x00'))]),
            'unknown trailer bytes': movie(3, 2, [frame(chunk(16, b'\x00'*6 + b'\x01\x02'))]),
            'bad postage stamp': movie(3, 2, [frame(chunk(18, b'\x00'*12))]),
        }
        examples['bad magic'][4:6] = b'\x44\xaf'
        examples['bad length'][0] ^= 1
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            for label, data in examples.items():
                with self.subTest(label):
                    source = root / f'{label}.flc'
                    destination = root / f'{label}.json'
                    source.write_bytes(data)
                    with self.assertRaises(ValueError):
                        flic.run(['--input', str(source), '--report', str(destination)])
                    self.assertFalse(destination.exists())
            source = root / 'valid.flc'
            source.write_bytes(valid)
            target = root / 'existing.json'
            target.write_bytes(b'owned and unchanged')
            with self.assertRaises(ValueError):
                flic.run(['--input', str(source), '--report', str(target)])
            self.assertEqual(target.read_bytes(), b'owned and unchanged')
            with self.assertRaises(ValueError):
                flic.run(['--input', str(source), '--expect-sha256', '0'*64,
                          '--report', str(root / 'wrong-hash.json')])
            self.assertFalse((root / 'wrong-hash.json').exists())


if __name__ == '__main__':
    unittest.main()
