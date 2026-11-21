"""Constructor-only Acotel gate: parsers, refusals, and optional live original."""
from __future__ import annotations

from hashlib import sha256
import importlib.util
import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / 'tools/midlet-acotel-constructor-oracle.py'
SPEC = importlib.util.spec_from_file_location('midlet_acotel_constructor_oracle', SCRIPT)
assert SPEC and SPEC.loader
oracle = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(oracle)
EMU = Path('/var/tmp/magnolia-microemulator-2.0.4.jar')
GAME = Path('/var/tmp/magnolia-j2me-acotel/ACO_XmasGifts_s.jar')
MAGNOLIA = Path(os.environ.get('MAGNOLIA_ACOTEL_TEST_EXECUTABLE',
    '/var/tmp/magnolia-midp-generic-refactor/magnolia'))


def reference_line() -> str:
    return ('headless emulator\nREFERENCE:' + ';'.join(
        key + '=' + value for key, value in oracle.REFERENCE.items()) + '\n')


def native_record() -> dict:
    return {
        'type': 'ok', 'phase': 'constructed-only', 'actions': ['construct'],
        'constructedClass': 'ACO_XmasGifts_s', 'canvasClass': 'DS_Effects',
        'output': '', 'steps': 18_927, 'constructorCanvasPixels': 30_976,
        'constructorCanvasArgbSha256': oracle.BLACK_SHA,
        'constructorCanvasDistinctColors': 1,
        'registeredCommandCount': 1, 'registeredCommandLabel': 'Fine',
        'registeredCommandType': 1, 'registeredCommandPriority': 0,
        'commandListenerClass': 'ACO_XmasGifts_s',
        'modelDisplaySelected?': False, 'modelWorkerStarted?': False,
        'modelStartAppCalled?': False, 'modelPaintCalled?': False,
        'inputDelivered?': False,
    }


class AcotelConstructorOracleTests(unittest.TestCase):
    def test_matching_black_constructor_only(self):
        self.assertEqual(oracle.BLACK_SHA,
            '582076772019e80f7184d6b85a32c546dbb60b88c72dbd4ca84effec1a6647f0')
        self.assertEqual(oracle.parse_reference(reference_line()), oracle.REFERENCE)
        self.assertEqual(oracle.parse_native(json.dumps(native_record())), native_record())
        self.assertFalse(native_record()['modelStartAppCalled?'])

    def test_reference_wrong_identity_frame_status_and_ambiguity_refuse(self):
        valid = reference_line()
        changes = [
            valid.replace('ACO_XmasGifts_s', 'not-original'),
            valid.replace('canvas=DS_Effects', 'canvas=Other'),
            valid.replace('offscreen=176x176', 'offscreen=100x100'),
            valid.replace('colors=1', 'colors=2'),
            valid.replace('argbSHA256=' + oracle.BLACK_SHA, 'argbSHA256=' + '0' * 64),
            valid.replace('registered=Fine:1:0', 'registered=Other:1:0'),
            valid.replace('listener=same-object', 'listener=other-object'),
            valid.replace('worker=false', 'worker=true'),
            valid.replace('gameStartAppCalled=false', 'gameStartAppCalled=true'),
            valid.replace('gamePaintCalled=false', 'gamePaintCalled=true'),
            valid.replace('sleep=20', 'sleep=21'),
            valid + valid.splitlines()[-1] + '\n',
            valid.replace(';state=0', ';state=0;state=0'),
            valid.replace(';stop=false', ''),
            valid.replace(';stop=false', ';stop=false;extra=x'),
            'REFERENCE:constructor-failed java.lang.VerifyError\n',
        ]
        for sample in changes:
            with self.subTest(log=sample[-135:]), self.assertRaises(ValueError):
                oracle.parse_reference(sample)

    def test_native_wrong_phase_result_frame_step_and_status_refuse(self):
        good = native_record()
        for key, bad in [
            ('type', 'error'), ('phase', 'active'), ('actions', ['start']),
            ('output', 'BEGIN'), ('constructorCanvasPixels', 30_975),
            ('constructorCanvasArgbSha256', '0' * 64),
            ('constructorCanvasDistinctColors', 2),
            ('registeredCommandLabel', 'No'), ('registeredCommandCount', 0),
            ('modelDisplaySelected?', True), ('modelWorkerStarted?', True),
            ('modelStartAppCalled?', True), ('modelPaintCalled?', True),
            ('inputDelivered?', True), ('steps', 0), ('steps', 50_001),
            ('steps', True), ('steps', 1000.5),
        ]:
            changed = dict(good, **{key: bad})
            with self.subTest(key=key, bad=bad), self.assertRaises(ValueError):
                oracle.parse_native(json.dumps(changed))
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps(dict(good, gameFrame='success')))
        with self.assertRaises(ValueError):
            oracle.parse_native('not JSON')
        with self.assertRaises(ValueError):
            oracle.parse_native(json.dumps(good) + '\n' + json.dumps(good))

    def test_generic_pinned_constructor_preserves_original_only_field_claims(self):
        generic = {
            'type': 'ok', 'phase': 'constructed-only-bounded',
            'actions': ['construct'], 'output': '',
            'constructedClass': oracle.ENTRY, 'steps': oracle.CONSTRUCTOR_STEPS,
            'requestedDisplay': None,
            # No logical selection is NOT absence of the original Fine command.
            'registeredCommands': {'count': 0, 'labels': [], 'listenerClass': None},
            'modelStartAppCalled?': False, 'modelWorkerStarted?': False,
            'modelPaintCalled?': False, 'modelDisplaySelected?': False,
            'deviceInputDelivered?': False,
            'virtualFreeBytes': oracle.SNAPSHOT, 'runtimeGcCalls': 0,
        }
        self.assertEqual(oracle.parse_generic_pinned_constructor(json.dumps(generic)),
                         generic)
        for key, value in [
            ('steps', 18_928), ('constructedClass', 'OtherGame'),
            ('requestedDisplay', 'DS_Effects'),
            ('registeredCommands', {'count': 1, 'labels': ['Fine'],
                                    'listenerClass': oracle.ENTRY}),
            ('modelStartAppCalled?', True), ('modelWorkerStarted?', True),
            ('modelPaintCalled?', True), ('modelDisplaySelected?', True),
            ('deviceInputDelivered?', True), ('virtualFreeBytes', 123),
        ]:
            with self.subTest(key=key), self.assertRaises(ValueError):
                oracle.parse_generic_pinned_constructor(
                    json.dumps({**generic, key: value}))
        for raw in (json.dumps({**generic, 'gameplay': True}),
                    json.dumps(generic) + '\n' + json.dumps(generic),
                    json.dumps(generic).replace('"type": "ok"',
                        '"type": "ok", "type": "ok"')):
            with self.assertRaises(ValueError):
                oracle.parse_generic_pinned_constructor(raw)

    def test_refusals_are_not_counted_as_execution(self):
        message = 'Acotel constructor-only pin/profile mismatch before original construction'
        oracle.check_refusal((1, '[midlet] ' + message + '\n'), message)
        for code, output in [(0, json.dumps(native_record())),
                             (2, '[midlet] ' + message + '\n'),
                             (1, '[midlet] ' + message + '\nBEGIN\n'),
                             (1, 'REFERENCE:constructor=ACO_XmasGifts_s')]:
            with self.subTest(code=code, output=output[:32]), self.assertRaises(ValueError):
                oracle.check_refusal((code, output), message)

    def test_wrong_hash_size_and_internal_original_refuse(self):
        with tempfile.TemporaryDirectory() as tmp:
            bad = Path(tmp) / 'not-the-original.jar'
            bad.write_bytes(b'x' * oracle.GAME_SIZE)
            with self.assertRaisesRegex(ValueError, 'SHA-256'):
                oracle.pinned(bad, oracle.GAME_SIZE, oracle.GAME_SHA)
            bad.write_bytes(b'x')
            with self.assertRaisesRegex(ValueError, 'byte count'):
                oracle.pinned(bad, oracle.GAME_SIZE, oracle.GAME_SHA)
        internal = ROOT / 'README.md'
        with self.assertRaisesRegex(ValueError, 'outside the checkout'):
            oracle.pinned(internal, internal.stat().st_size,
                          sha256(internal.read_bytes()).hexdigest())

    def test_live_pinned_original_if_separately_available(self):
        if not all(path.is_file() for path in (GAME, EMU, MAGNOLIA)):
            self.skipTest('external pinned game/emulator and built Magnolia not present')
        with tempfile.TemporaryDirectory() as tmp:
            report = Path(tmp) / 'constructor.json'
            result = subprocess.run([sys.executable, str(SCRIPT),
                '--magnolia', str(MAGNOLIA), '--jar', str(GAME),
                '--microemu', str(EMU), '--output', str(report)],
                capture_output=True, text=True, timeout=205)
            self.assertEqual(result.returncode, 0, result.stderr[-900:])
            self.assertTrue(json.loads(result.stdout)['passed'])
            good = report.read_bytes()
            data = json.loads(good)
            self.assertTrue(data['reference']['realMicroEmulatorOriginalConstructorExecuted'])
            self.assertTrue(data['magnolia']['unchangedOriginalConstructorExecuted'])
            self.assertFalse(data['magnolia']['originalGameStartAppExecuted'])
            self.assertFalse(data['comparison']['independentOriginalGameStartupCompared'])
            self.assertEqual(data['magnolia']['offscreenArgbSha256'], oracle.BLACK_SHA)
            generic = data['genericPinnedConstructor']
            self.assertTrue(generic['sameConstructorStepsAndIdentityAsOriginalFieldOracle'])
            self.assertEqual(generic['chargedSteps'], oracle.CONSTRUCTOR_STEPS)
            self.assertIsNone(generic['logicalRequestedDisplay'])
            self.assertTrue(generic['unselectedOriginalCommandVerifiedOnlyByFieldOracle'])
            self.assertFalse(generic['originalStartAppWorkerPaintOrDeviceInputExecuted'])
            self.assertTrue(generic['changedJarRefusedBeforeConstruction'])
            self.assertEqual(sum(data['refusals'].values()), 5)
            mutant = Path(tmp) / 'changed-original.jar'
            oracle.changed_valid_jar(GAME, mutant)
            denied = subprocess.run([sys.executable, str(SCRIPT),
                '--magnolia', str(MAGNOLIA), '--jar', str(mutant),
                '--microemu', str(EMU), '--output', str(report)],
                capture_output=True, text=True, timeout=4)
            self.assertEqual(denied.returncode, 2)
            self.assertIn('SHA-256 differs', denied.stderr)
            self.assertEqual(report.read_bytes(), good)


if __name__ == '__main__':
    unittest.main()
