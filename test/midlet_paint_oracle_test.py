"""No-download adversarial guards for the opt-in one-frame MIDP oracle.

The live JDK/MicroEmulator/Magnolia differential is separately run by
`tools/midlet-paint-oracle.py` with hash-pinned external inputs.
"""
from __future__ import annotations

import importlib.util
import json
from pathlib import Path
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location('midlet_paint_oracle',
                                               ROOT / 'tools/midlet-paint-oracle.py')
assert spec and spec.loader
oracle = importlib.util.module_from_spec(spec)
spec.loader.exec_module(oracle)


class OneFrameBounds(unittest.TestCase):
    def test_independent_frame_bytes_match_pinned_digest(self) -> None:
        self.assertEqual(oracle.expected_frame_sha(), oracle.FRAME_SHA)

    def test_reference_must_have_real_cropped_pixels_and_order(self) -> None:
        valid = '\n'.join([*oracle.TRACE, oracle.SUMMARY]) + '\n'
        self.assertEqual(oracle.real_trace(valid), [*oracle.TRACE, oracle.SUMMARY])
        for wrong in (valid.replace('reds=144', 'reds=143'),
                      valid.replace('box=10,34,25,42', 'box=10,12,25,20'),
                      valid.replace('TRACE:selected\nTRACE:shown',
                                    'TRACE:shown\nTRACE:selected'),
                      valid + 'BAD:unexpected\n'):
            with self.subTest(wrong=wrong[-90:]):
                with self.assertRaisesRegex(ValueError, 'MicroEmulator'):
                    oracle.real_trace(wrong)

    def test_native_result_refuses_metadata_only_or_fake_game(self) -> None:
        data = {'type': 'ok', 'output': '\n'.join(oracle.TRACE) + '\n',
                'requestedDisplay': 'PaintSliceMIDlet$View',
                'actions': ['start'], 'phase': 'active', 'painted?': True,
                'inputDelivered?': False, 'canvasPixels': 176 * 176,
                'canvasArgbSha256': oracle.FRAME_SHA, 'runtimeGcCalls': 0,
                'registeredCommands': {'count': 0}, 'virtualFreeBytes': None,
                'steps': 1112}
        self.assertEqual(oracle.native_result(json.dumps(data)), data)
        for key, value in [('painted?', False), ('canvasPixels', 0),
                           ('canvasArgbSha256', '0' * 64),
                           ('inputDelivered?', True), ('steps', 10001),
                           ('output', 'TRACE:paint\n')]:
            with self.subTest(key=key):
                bad = dict(data, **{key: value})
                with self.assertRaisesRegex(ValueError, 'native one-frame'):
                    oracle.native_result(json.dumps(bad))

    def test_no_external_pin_is_inferred_or_copied_into_repo(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            path = Path(tmp) / 'fake-reference.jar'
            path.write_bytes(b'not a MIDP implementation')
            with self.assertRaisesRegex(ValueError, 'external reference'):
                oracle.pin_external(path, oracle.EMU_SIZE, oracle.EMU_SHA)
        with self.assertRaisesRegex(ValueError, 'external reference'):
            oracle.pin_external(ROOT / 'README.md',
                                (ROOT / 'README.md').stat().st_size,
                                'irrelevant')


if __name__ == '__main__':
    unittest.main()
