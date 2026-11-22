"""Original Java ME game reference evidence: narrow parsing, pins and optional live run.

The live game and emulator JARs remain outside this source tree. No JAD or
remote URL is read; this is independent MicroEmulator execution, not Magnolia.
"""
from __future__ import annotations

from hashlib import sha256
import importlib.util
import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / 'tools/acotel-original-reference.py'
SPEC = importlib.util.spec_from_file_location('acotel_original_reference', SCRIPT)
assert SPEC and SPEC.loader
reference = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(reference)
EMU = Path('/var/tmp/magnolia-microemulator-2.0.4.jar')
GAME = Path('/var/tmp/magnolia-j2me-acotel-run/ACO_XmasGifts_s.jar')
FRAME = 'bac65c402bb433e005edfae87c2b5d47eb90ef39796ef561a281c2f9cde3dc89'


def log() -> str:
    return ('\n'.join('REF:tick=' + str(i) + ',midlet=ACO_XmasGifts_s,'
        'current=DS_Effects,device=176x176,pixels=38720,'
        'argbSHA256=' + 'f' * 64 + ',gameOffscreen=176x176,colors=317,sha256=' + FRAME
        for i in range(7)) + '\nREF:observed-original-startup\n')


class OriginalAcotelReferenceTests(unittest.TestCase):
    def test_strict_observations_not_a_gameplay_claim(self):
        result = reference.interpret(log())
        self.assertEqual(result['currentCanvas'], 'DS_Effects')
        self.assertEqual(result['gameOffscreen'], '176x176')
        self.assertEqual(result['distinctReferenceFrameHashes'], [FRAME])
        self.assertFalse(result['frameTimingDeterministic'])
        self.assertFalse(result['deviceDisplayImageEquivalentToGameOffscreen'])
        self.assertEqual(result['nonBlankReferenceTicks'], 7)

    def test_fabricated_game_or_blank_or_bad_device_refuses(self):
        valid = log()
        for mutated in (valid.replace('DS_Effects', 'OtherClass'),
                        valid.replace('colors=317', 'colors=1'),
                        valid.replace('device=176x176', 'device=100x100'),
                        valid.replace('tick=0', 'tick=9'),
                        valid.replace('REF:observed-original-startup', 'REF:FAILED'),
                        valid.replace('midlet=ACO_XmasGifts_s', 'midlet=FakeGame'),
                        valid.replace('pixels=38720', 'pixels=176'),
                        valid.replace('sha256=' + FRAME, 'sha256=invalid')):
            with self.subTest(mutated=mutated[:80]), self.assertRaises(ValueError):
                reference.interpret(mutated)

    def test_wrong_hash_and_internal_paths_refuse(self):
        with tempfile.TemporaryDirectory() as tmp:
            dummy = Path(tmp) / 'game.jar'
            dummy.write_bytes(b'JAR')
            with self.assertRaises(ValueError):
                reference.pinned(dummy, 3, reference.GAME_SHA)
        with self.assertRaises(ValueError):
            reference.pinned(ROOT / 'README.md',
                             (ROOT / 'README.md').stat().st_size,
                             sha256((ROOT / 'README.md').read_bytes()).hexdigest())

    def test_live_original_game_if_external_jars_present(self):
        if not EMU.is_file() or not GAME.is_file():
            self.skipTest('separately cached original/emulator JARs not present')
        with tempfile.TemporaryDirectory() as tmp:
            report = Path(tmp) / 'report.json'
            result = subprocess.run([sys.executable, str(SCRIPT), '--jar', str(GAME),
                '--microemu', str(EMU), '--output', str(report)],
                text=True, capture_output=True, timeout=24)
            self.assertEqual(result.returncode, 0, result.stderr[-800:])
            self.assertTrue(report.is_file())
            gate = json.loads(result.stdout)
            self.assertTrue(gate['originalGameRenderedInEmulator'])
            self.assertFalse(gate['magnoliaOriginalGameExecuted'])
            self.assertGreaterEqual(gate['reference']['nonBlankReferenceTicks'], 1)
            data = json.loads(report.read_text())
            self.assertFalse(data['magnoliaOriginalGamePainted'])
            self.assertEqual(data['reference']['originalGameSha256'], reference.GAME_SHA)
            # Same-length corrupted JAR must refuse without launching the
            # reference child or replacing the successful metadata report.
            original_report = report.read_bytes()
            corrupted = Path(tmp) / 'bad.jar'
            raw = bytearray(GAME.read_bytes())
            raw[12] ^= 1
            corrupted.write_bytes(raw)
            denied = subprocess.run([sys.executable, str(SCRIPT), '--jar', str(corrupted),
                '--microemu', str(EMU), '--output', str(report)],
                text=True, capture_output=True, timeout=4)
            self.assertEqual(denied.returncode, 2)
            self.assertIn('SHA-256 differs', denied.stderr)
            self.assertEqual(report.read_bytes(), original_report)


if __name__ == '__main__':
    unittest.main()
