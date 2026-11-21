"""Pure bounds checks; executable oracle separately runs real emulator/Magnolia."""
from __future__ import annotations

from hashlib import sha256
import importlib.util
import json
from pathlib import Path
import unittest

ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location(
    'midlet_service_oracle', ROOT / 'tools/midlet-service-oracle.py')
assert spec and spec.loader
oracle = importlib.util.module_from_spec(spec)
spec.loader.exec_module(oracle)


class TwoFrameBounds(unittest.TestCase):
    def test_independent_full_frame_bytes(self):
        self.assertEqual(oracle.frame_sha(10, 12, 16, 9, b'\xff\xff\x00\x00'),
                         oracle.FIRST_SHA)
        self.assertEqual(oracle.frame_sha(32, 41, 13, 7, b'\xff\x00\xff\x00'),
                         oracle.SECOND_SHA)
        self.assertNotEqual(oracle.FIRST_SHA, oracle.SECOND_SHA)
        self.assertNotEqual(oracle.frame_sha(32, 41, 13, 7, b'\xff\xff\x00\x00'),
                            oracle.SECOND_SHA)

    def test_emulator_order_accepts_only_the_actual_race(self):
        prefix = [*oracle.PREFIX, oracle.FIRST_SUMMARY]
        for early in ([oracle.REQUEST, oracle.SECOND],
                      [oracle.SECOND, oracle.REQUEST]):
            correct = [*prefix, *early, oracle.SERVICE, oracle.SECOND_SUMMARY]
            self.assertEqual(oracle.real_trace('\n'.join(correct)), correct)
            for wrong in (correct[:-1],
                          [*prefix, oracle.SECOND, oracle.SECOND, oracle.SERVICE,
                           oracle.SECOND_SUMMARY],
                          [*prefix, oracle.REQUEST, oracle.SERVICE, oracle.SECOND,
                           oracle.SECOND_SUMMARY],
                          [*prefix, *early, oracle.SERVICE,
                           oracle.SECOND_SUMMARY.replace('paints=2', 'paints=1')]):
                with self.assertRaises(ValueError):
                    oracle.real_trace('\n'.join(wrong))

    def test_native_result_rejects_one_frame_or_fabricated_second(self):
        good = {
            'type': 'ok', 'output': '\n'.join([
                *oracle.PREFIX, oracle.REQUEST, oracle.SECOND, oracle.SERVICE]) + '\n',
            'actions': ['start', 'service'], 'requestedDisplay': 'PaintServiceMIDlet$View',
            'painted?': True, 'inputDelivered?': False, 'phase': 'active',
            'virtualFreeBytes': None, 'runtimeGcCalls': 0,
            'registeredCommands': {'count': 0}, 'canvasPixels': 30976,
            'canvasFrameCount': 2, 'canvasFirstArgbSha256': oracle.FIRST_SHA,
            'canvasArgbSha256': oracle.SECOND_SHA, 'steps': 1697,
        }
        self.assertEqual(oracle.native_result(json.dumps(good)), good)
        with self.assertRaises(ValueError):
            oracle.native_result(json.dumps([good]))
        for key, wrong in (('canvasFrameCount', 1),
                           ('canvasFirstArgbSha256', oracle.SECOND_SHA),
                           ('canvasArgbSha256', oracle.FIRST_SHA),
                           ('output', '\n'.join(oracle.PREFIX) + '\n'),
                           ('steps', 10001), ('inputDelivered?', True)):
            fake = {**good, key: wrong}
            with self.assertRaises(ValueError):
                oracle.native_result(json.dumps(fake))

    def test_owned_pins_and_no_shipped_native_stub(self):
        sources = {
            'ServiceCanvas.java': oracle.SERVICE_SRC,
            'PaintServiceMIDlet.java': oracle.FIXTURE_SRC / 'PaintServiceMIDlet.java',
            'PaintServiceUnsupportedMIDlet.java': oracle.FIXTURE_SRC / 'PaintServiceUnsupportedMIDlet.java',
            'PaintServiceTwiceMIDlet.java': oracle.FIXTURE_SRC / 'PaintServiceTwiceMIDlet.java',
            'PaintServiceReference.java': oracle.REFERENCE_SRC,
        }
        self.assertEqual({key: sha256(path.read_bytes()).hexdigest()
                          for key, path in sources.items()}, oracle.SOURCE_SHA)
        self.assertEqual(sha256(oracle.SERVICE_CLASS.read_bytes()).hexdigest(),
                         oracle.MODEL_SHA)
        self.assertEqual({path.name for path in oracle.SERVICE_CLASS.parent.glob('*.class')},
                         {'Canvas.class'})
        self.assertEqual(sha256(oracle.GRAPHICS_CLASS.read_bytes()).hexdigest(),
                         oracle.GRAPHICS_SHA)
        self.assertIn('BufferedImage)', oracle.REFERENCE_SRC.read_text())
        self.assertIn('return backing.getRGB(', oracle.REFERENCE_SRC.read_text())

    def test_saved_report_is_only_a_bounded_owned_fixture_gate(self):
        report = json.loads((ROOT / 'reports/midlet-service-reference.json').read_text())
        self.assertEqual(len(report['reference']['runs']), 3)
        for run in report['reference']['runs']:
            self.assertEqual(oracle.real_trace('\n'.join(run)), run)
        self.assertEqual(report['native']['firstARGBSHA256'], oracle.FIRST_SHA)
        self.assertEqual(report['native']['secondARGBSHA256'], oracle.SECOND_SHA)
        self.assertTrue(report['native']['standaloneClassAndOwnedJarMatched'])
        self.assertFalse(report['originalGameConstructedOrPaintedInMagnolia'])
        self.assertEqual(len(report['negativeFirstRefusals']), 9)
        self.assertIn('triggerService()V', report['negativeFirstRefusals']['originalServiceIsNotGeneric'])


if __name__ == '__main__':
    unittest.main()
