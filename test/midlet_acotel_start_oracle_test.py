"""Original startApp-return gate: worker scheduled but not run in Magnolia."""
from __future__ import annotations

import importlib.util
import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / 'tools/midlet-acotel-start-oracle.py'
SPEC = importlib.util.spec_from_file_location('midlet_acotel_start_oracle', SCRIPT)
assert SPEC and SPEC.loader
oracle = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(oracle)
EMU = Path('/var/tmp/magnolia-microemulator-2.0.4.jar')
GAME = Path('/var/tmp/magnolia-j2me-acotel/ACO_XmasGifts_s.jar')
MAGNOLIA = Path(os.environ.get('MAGNOLIA_ACOTEL_START_TEST_EXECUTABLE',
                              '/var/tmp/magnolia-midp-generic-refactor/magnolia'))


def reference_line(colors: str = '317') -> str:
    record = {**oracle.REFERENCE, 'workerDrewColors': colors}
    return 'headless emulator\nSTART:' + ';'.join(
        key + '=' + value for key, value in record.items()) + '\n'


def native_record() -> dict:
    return {**oracle.NATIVE, 'steps': oracle.STEPS}


class OriginalAcotelScheduledStart(unittest.TestCase):
    def test_reference_asserts_real_original_worker_but_no_pixel_parity(self) -> None:
        record = oracle.parse_reference(reference_line())
        self.assertEqual(record['originalStartAppReturned'], 'true')
        self.assertEqual(record['workerStarted'], 'true')
        self.assertEqual(record['workerDrewColors'], '317')
        self.assertEqual(record['noPixelParityWithMagnolia'], 'true')
        self.assertEqual(oracle.parse_reference(reference_line('357'))[
            'workerDrewColors'], '357')
        self.assertEqual(oracle.parse_native(json.dumps(native_record())),
                         native_record())
        self.assertFalse(native_record()['workerRunExecuted?'])
        self.assertFalse(native_record()['modelDisplaySelected?'])

    def test_reference_wrong_state_colors_order_or_fake_result_refuses(self) -> None:
        valid = reference_line()
        for old, new in [
            ('original=ACO_XmasGifts_s', 'original=FakeGame'),
            ('canvas=DS_Effects', 'canvas=FakeCanvas'),
            ('preBlack=true', 'preBlack=false'),
            ('preSelected=false', 'preSelected=true'),
            ('preWorker=false', 'preWorker=true'),
            ('registered=Fine:1:0', 'registered=Fake:1:0'),
            ('listener=same-object', 'listener=other-object'),
            ('originalStartAppReturned=true', 'originalStartAppReturned=false'),
            ('workerStarted=true', 'workerStarted=false'),
            ('eventuallySelected=true', 'eventuallySelected=false'),
            ('deviceInputDelivered=false', 'deviceInputDelivered=true'),
            ('noPixelParityWithMagnolia=true', 'noPixelParityWithMagnolia=false'),
        ]:
            with self.subTest(old=old), self.assertRaises(ValueError):
                oracle.parse_reference(valid.replace(old, new))
        for colors in ('1', '0', '-1', '30977', 'not-a-number', '1e3', '١١'):
            with self.subTest(colors=colors), self.assertRaises(ValueError):
                oracle.parse_reference(reference_line(colors))
        for bad in (valid + valid.splitlines()[-1] + '\n',
                    valid.replace('START:', 'START:extra=x;'),
                    valid.replace(';workerDrewColors=317',
                                  ';workerDrewColors=317;workerDrewColors=317'),
                    'START:failed java.lang.VerifyError\n'):
            with self.subTest(bad=bad[-110:]), self.assertRaises(ValueError):
                oracle.parse_reference(bad)

    def test_native_cannot_forge_worker_paint_or_device_input(self) -> None:
        good = native_record()
        for key, value in [
            ('steps', 18_927), ('steps', 18_958), ('steps', True),
            ('type', 'error'), ('phase', 'active'),
            ('actions', ['construct']), ('output', 'fake game'),
            ('requestedDisplay', None), ('registeredCommandLabel', 'Quit'),
            ('originalStartAppExecuted?', False),
            ('workerThreadStartCalled?', False),
            ('workerRunnableDeferred?', False), ('workerRunExecuted?', True),
            ('modelDisplaySelected?', True), ('modelPaintCalled?', True),
            ('deviceInputDelivered?', True), ('gameplayOrVisibleFrame?', True),
            ('originalOffscreenPixels', 0),
            ('originalOffscreenArgbSha256', '0' * 64),
        ]:
            with self.subTest(key=key, value=value), self.assertRaises(ValueError):
                oracle.parse_native(json.dumps({**good, key: value}))
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps({**good, 'fakeSuccess': True}))
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps(good) + '\n' + json.dumps(good))
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps(good).replace('"type": "ok"',
                '"type": "ok", "type": "ok"'))

    def test_saved_gate_is_start_return_not_frame_or_gameplay(self) -> None:
        report = ROOT / 'reports/midlet-acotel-start-return.json'
        if not report.is_file():
            self.skipTest('live start-return metadata gate not yet generated')
        data = json.loads(report.read_text())
        self.assertEqual(data['magnolia']['chargedStepsIncludingConstructor'],
                         oracle.STEPS)
        self.assertTrue(data['magnolia']['originalConstructorAndStartAppExecuted'])
        self.assertTrue(data['magnolia']['workerThreadStartCalledButRunnableDeferred'])
        self.assertFalse(data['magnolia']['originalRunnableRunExecuted'])
        self.assertFalse(data['magnolia']['originalPaintExecuted'])
        self.assertFalse(data['magnolia']['deviceDisplaySelected'])
        self.assertFalse(data['magnolia']['deviceKeypadEventDelivered'])
        self.assertFalse(data['magnolia']['gameplayOrVisibleFrame'])
        self.assertTrue(data['reference']['workerThreadRunningInReference'])
        self.assertTrue(data['reference']['deviceDisplayEventuallySelected'])
        self.assertGreater(data['reference']['offscreenDistinctColorsAfterWorker'], 1)
        self.assertFalse(data['comparison']['sameWorkerExecutionOrVisibleFrame'])
        self.assertFalse(data['comparison']['sameActualDevicePixels'])
        self.assertTrue(all(data['refusals'].values()))

    def test_live_original_changed_jar_and_no_report_overwrite(self) -> None:
        if not all(path.is_file() for path in (EMU, GAME, MAGNOLIA)):
            self.skipTest('separately pinned game/emulator and built Magnolia absent')
        with tempfile.TemporaryDirectory() as temp:
            report = Path(temp) / 'start-return.json'
            finished = subprocess.run([sys.executable, str(SCRIPT),
                '--magnolia', str(MAGNOLIA), '--jar', str(GAME),
                '--microemu', str(EMU), '--output', str(report)],
                cwd=ROOT, capture_output=True, text=True, timeout=315)
            self.assertEqual(finished.returncode, 0, finished.stderr[-1000:])
            self.assertTrue(json.loads(finished.stdout)['passed'])
            good = report.read_bytes()
            data = json.loads(good)
            self.assertEqual(data['magnolia']['chargedStepsIncludingConstructor'],
                             oracle.STEPS)
            self.assertFalse(data['magnolia']['originalRunnableRunExecuted'])
            self.assertGreater(data['reference']['offscreenDistinctColorsAfterWorker'], 1)
            altered = Path(temp) / 'same-size-valid-altered.jar'
            oracle.base.changed_valid_jar(GAME, altered)
            rejected = subprocess.run([sys.executable, str(SCRIPT),
                '--magnolia', str(MAGNOLIA), '--jar', str(altered),
                '--microemu', str(EMU), '--output', str(report)],
                cwd=ROOT, capture_output=True, text=True, timeout=6)
            self.assertEqual(rejected.returncode, 2)
            self.assertIn('SHA-256 differs', rejected.stderr)
            self.assertEqual(report.read_bytes(), good)


if __name__ == '__main__':
    unittest.main()
