"""Owned finite SWF1 PCM16->DefineSound ADPCM encoder / pre-output refusal tests."""
from __future__ import annotations

import hashlib
import importlib.util
from pathlib import Path
import struct
import subprocess
import sys
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location('sdk_adpcm_writer', ROOT/'tools/swf-adpcm-encode.py')
assert SPEC and SPEC.loader
writer = importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name] = writer
SPEC.loader.exec_module(writer)


def sha(b: bytes) -> str:
    return hashlib.sha256(b).hexdigest()


class SoundWriterTest(unittest.TestCase):
    def test_original_cpp_mono_payload_and_complete_swf(self) -> None:
        samples = [0, 1000, 1250, 3000, 2500, 0, -2000, -2500, 6000, -3000]
        pcm = struct.pack('<10h', *samples)
        payload = writer.define_sound(pcm, 1, 22050)
        swf = writer.swf_sound(pcm, 1, 22050)
        self.assertEqual((sha(payload), sha(swf)), (
            'b644017afd6590f7eb89af169326a507b45dc17507fe6c0b2a1dbea11418d415',
            'c10a3f7c4017314a914bfb65967659e7c7f232e49fe6ec9a7ec748e039606062'))
        self.assertEqual(len(swf), 37)
        self.assertEqual(struct.unpack_from('<I', swf, 4)[0], len(swf))
        self.assertEqual(swf[16:18], struct.pack('<H', (14 << 6) | len(payload)))
        self.assertEqual(swf[18:18 + len(payload)], payload)
        self.assertEqual(swf[-4:], struct.pack('<HH', 64, 0))

    def test_original_cpp_stereo_interleaving(self) -> None:
        pairs = [(500 + (i*1350*((-1)**(i//4)))%15000,
                  -600+(i*-840)%18000) for i in range(12)]
        pcm = struct.pack('<24h', *(v for pair in pairs for v in pair))
        self.assertEqual(sha(writer.define_sound(pcm, 2, 44100)),
            'c664c67c584540f1a64d90c7510eeea906e303362fdf79d9a3c3bc89417078a3')
        self.assertEqual(sha(writer.swf_sound(pcm, 2, 44100)),
            '230d2b25675cd983a8b3705ce28dc19f306ae9502f3d899e50ef0db54c667803')

    def test_original_cpp_block_boundary_and_extended_tag(self) -> None:
        samples = [((i*1049)%64000)-32000 for i in range(4098)]
        pcm = struct.pack('<4098h', *samples)
        payload = writer.define_sound(pcm, 1, 11025)
        swf = writer.swf_sound(pcm, 1, 11025)
        self.assertEqual(sha(payload),
            'd257610cd13923096e9ff109690b695d7f39c1c60a42f4e38d40c5599c1959c9')
        self.assertEqual(sha(swf),
            'dcc95324c2d38a4538a23efb8a9da0e8f0903524be2363527d44673afbb1531b')
        self.assertEqual(struct.unpack_from('<H', swf, 16)[0], (14 << 6) | 63)
        self.assertEqual(struct.unpack_from('<I', swf, 18)[0], len(payload))
        self.assertEqual(swf[22:22 + len(payload)], payload)

    def test_unsupported_pcm_refuses_before_any_movie_bytes(self) -> None:
        for raw, channel, rate in (
            (b'', 1, 22050),
            (b'\0\0\0', 1, 22050),
            (b'\0'*8194, 1, 22050),  # 4097 frames, unsafe look-ahead
            (b'\0'*8, 2, 16000),
            (b'\0'*8, 3, 22050),
            (b'\0'*16802, 2, 22050),
        ):
            with self.subTest(length=len(raw), channel=channel, rate=rate):
                with self.assertRaises(ValueError):
                    writer.swf_sound(raw, channel, rate)

    def test_cli_pin_partial_existing_and_symlink_preoutput_refusal(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            good = root/'good.pcm'
            good.write_bytes(struct.pack('<10h', *range(10)))
            linked = root/'linked.pcm'
            linked.symlink_to(good)
            part = root/'part.pcm'
            part.write_bytes(good.read_bytes()[:-1])
            dest = root/'movie.swf'
            cmd = [sys.executable, str(ROOT/'tools/swf-adpcm-encode.py')]
            for label, source, extra in (
                    ('pin', good, ['--expect-sha256', '0'*64]),
                    ('partial', part, []),
                    ('symlink', linked, [])):
                out = root/(label+'.swf')
                run = subprocess.run([*cmd, '--pcm', str(source), '--channels','1',
                    '--rate','22050', '--output',str(out), *extra],
                    cwd=ROOT, capture_output=True, text=True, timeout=6)
                self.assertEqual(run.returncode, 2, msg=run.stderr)
                self.assertIn('ADPCM refused before output:', run.stderr)
                self.assertFalse(out.exists())
            dest.write_bytes(b'keep unchanged')
            run = subprocess.run([*cmd, '--pcm', str(good), '--channels','1',
                '--rate','22050', '--output',str(dest)],
                cwd=ROOT, capture_output=True, text=True, timeout=6)
            self.assertEqual(run.returncode, 2)
            self.assertEqual(dest.read_bytes(), b'keep unchanged')


if __name__ == '__main__':
    unittest.main()
