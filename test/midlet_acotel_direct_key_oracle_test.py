"""Original (unchanged) Acotel direct key callback gate, no device events."""
from __future__ import annotations

import importlib.util
import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

from test.midlet_acotel_constructor_oracle_test import native_record

ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / 'tools/midlet-acotel-direct-key-oracle.py'
SPEC = importlib.util.spec_from_file_location('midlet_acotel_direct_key_oracle', SCRIPT)
assert SPEC and SPEC.loader
oracle = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(oracle)
EMU = Path('/var/tmp/magnolia-microemulator-2.0.4.jar')
GAME = Path('/var/tmp/magnolia-j2me-acotel/ACO_XmasGifts_s.jar')
MAGNOLIA = Path(os.environ.get('MAGNOLIA_ACOTEL_TEST_EXECUTABLE',
    '/var/tmp/magnolia-midp-generic-refactor/magnolia'))


def reference_line() -> str:
    return ('headless emulator\nDIRECT:' + ';'.join(
        name + '=' + value for name, value in oracle.EXPECTED_REFERENCE.items()) + '\n')


def native_direct() -> dict:
    constructor = native_record()
    constructor.update({
        'phase': 'constructed-direct-keys-only',
        'steps': 19_084,
        'actions': oracle.ACTIONS.copy(),
        'originalKeyCallbacksExecuted?': True,
        'directKeyTrace': [dict(item) for item in oracle.TRACE],
    })
    return constructor


class OriginalAcotelDirectKeys(unittest.TestCase):
    def test_exact_original_method_states_not_device_key_delivery(self) -> None:
        self.assertEqual(oracle.parse_reference(reference_line()),
                         oracle.EXPECTED_REFERENCE)
        self.assertEqual(oracle.parse_native(json.dumps(native_direct())),
                         native_direct())
        self.assertEqual([item['callback'] for item in oracle.TRACE],
                         ['keyPressed', 'keyReleased', 'keyRepeated'])
        self.assertEqual([item['action'] for item in oracle.TRACE], [1, 1, 5])
        self.assertFalse(native_direct()['inputDelivered?'])
        self.assertFalse(native_direct()['modelStartAppCalled?'])

    def test_reference_wrong_order_or_fake_original_result_refuses(self) -> None:
        valid = reference_line()
        cases = [
            valid.replace('ACO_XmasGifts_s', 'Other'),
            valid.replace('canvas=DS_Effects', 'canvas=FakeCanvas'),
            valid.replace('before=0:false', 'before=1:true'),
            valid.replace('keyPressed(50)=1:true', 'keyPressed(50)=2:true'),
            valid.replace('keyReleased(50)=1:false', 'keyReleased(50)=1:true'),
            valid.replace('keyRepeated(-4)=5:true', 'keyRepeated(-4)=6:true'),
            valid.replace('keyRepeated(-4)', 'keyRepeated(-3)'),
            valid.replace('offscreen=176x176', 'offscreen=50x50'),
            valid.replace('argbSHA256=' + oracle.base.BLACK_SHA,
                          'argbSHA256=' + '0' * 64),
            valid.replace('workerStarted=false', 'workerStarted=true'),
            valid.replace('gameStartAppCalled=false', 'gameStartAppCalled=true'),
            valid.replace('deviceInputDelivered=false', 'deviceInputDelivered=true'),
            valid + valid.splitlines()[-1] + '\n',
            valid.replace(';before=0:false', ';before=0:false;before=0:false'),
            valid.replace(';before=0:false', ''),
            valid.replace(';before=0:false', ';before=0:false;extra=unknown'),
            'DIRECT:failed java.lang.VerifyError\n',
        ]
        for case in cases:
            with self.subTest(log=case[-135:]), self.assertRaises(ValueError):
                oracle.parse_reference(case)

    def test_native_wrong_actions_keys_state_worker_or_early_budget_refuses(self) -> None:
        good = native_direct()
        for key, value in [
            ('type', 'error'), ('phase', 'active'), ('actions', ['construct']),
            ('originalKeyCallbacksExecuted?', False), ('steps', 18_927),
            ('steps', 50_001), ('modelWorkerStarted?', True),
            ('modelDisplaySelected?', True), ('modelStartAppCalled?', True),
            ('modelPaintCalled?', True), ('inputDelivered?', True),
            ('constructorCanvasArgbSha256', '0' * 64),
        ]:
            mutated = dict(good, **{key: value})
            with self.subTest(key=key), self.assertRaises(ValueError):
                oracle.parse_native(json.dumps(mutated))
        for index, name, bad in [(0, 'action', 2), (1, 'down?', True),
                                 (2, 'callback', 'keyPressed'), (2, 'key', -3),
                                 (0, 'down?', 1)]:
            mutated = native_direct()
            mutated['directKeyTrace'][index][name] = bad
            with self.subTest(index=index, name=name), self.assertRaises(ValueError):
                oracle.parse_native(json.dumps(mutated))
        mutated = native_direct()
        mutated['directKeyTrace'].append(dict(oracle.TRACE[0]))
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps(mutated))
        mutated = native_direct()
        mutated['directKeyTrace'][0]['fakeGameStarted'] = True
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps(mutated))
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps(good) + '\n' + json.dumps(good))

    def test_saved_direct_gate_remains_narrower_than_gameplay(self) -> None:
        report = ROOT / 'reports/midlet-acotel-direct-keys.json'
        if not report.is_file():
            self.skipTest('live metadata gate not yet generated')
        data = json.loads(report.read_text())
        self.assertTrue(data['magnolia']['unchangedOriginalConstructorExecuted'])
        self.assertEqual(data['magnolia']['directKeyTrace'], oracle.TRACE)
        self.assertFalse(data['magnolia']['deviceKeypadEventDelivered'])
        self.assertFalse(data['magnolia']['originalStartAppExecuted'])
        self.assertFalse(data['magnolia']['gameplayOrVisibleFrame'])
        self.assertTrue(all(data['refusals'].values()))

    def test_live_original_and_changed_jar_if_external_artifacts_present(self) -> None:
        if not all(path.is_file() for path in (GAME, EMU, MAGNOLIA)):
            self.skipTest('external pinned original/emulator and built Magnolia absent')
        with tempfile.TemporaryDirectory() as tmp:
            report = Path(tmp) / 'direct-keys.json'
            result = subprocess.run([sys.executable, str(SCRIPT),
                '--magnolia', str(MAGNOLIA), '--jar', str(GAME),
                '--microemu', str(EMU), '--output', str(report)],
                text=True, capture_output=True, timeout=360)
            self.assertEqual(result.returncode, 0, result.stderr[-900:])
            self.assertTrue(json.loads(result.stdout)['passed'])
            good = report.read_bytes()
            data = json.loads(good)
            self.assertEqual(data['magnolia']['chargedStepsIncludingConstructor'], 19_084)
            self.assertEqual(data['reference']['directKeyTrace'],
                oracle.EXPECTED_REFERENCE['trace'])
            self.assertEqual(data['magnolia']['directKeyTrace'], oracle.TRACE)
            self.assertFalse(data['comparison']['independentGameplayOrDeviceEventCompared'])
            self.assertTrue(data['refusals']['wrongOrderCodeMethodAndExtraCallBeforeEffects'])
            self.assertTrue(data['refusals']['missingDirectKeyRetentionCapabilityBeforeConstruction'])
            # Corrupt only ZIP timestamps; it stays same-size, CRC-valid and
            # readable. The script must refuse on the SHA before either host
            # launches, and must not replace successful metadata.
            changed = Path(tmp) / 'altered.jar'
            oracle.base.changed_valid_jar(GAME, changed)
            refused = subprocess.run([sys.executable, str(SCRIPT),
                '--magnolia', str(MAGNOLIA), '--jar', str(changed),
                '--microemu', str(EMU), '--output', str(report)],
                text=True, capture_output=True, timeout=5)
            self.assertEqual(refused.returncode, 2)
            self.assertIn('SHA-256 differs', refused.stderr)
            self.assertEqual(report.read_bytes(), good)


if __name__ == '__main__':
    unittest.main()
