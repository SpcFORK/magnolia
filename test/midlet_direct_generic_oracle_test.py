"""Fail-closed, two-host tests for unrelated owned MIDP direct Canvas keys."""
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
from zipfile import ZIP_STORED, ZipFile, ZipInfo

ROOT = Path(__file__).resolve().parents[1]
SOURCE = ROOT / 'tools/midlet-direct-generic-oracle.py'
SPEC = importlib.util.spec_from_file_location('midlet_direct_generic_oracle', SOURCE)
assert SPEC and SPEC.loader
oracle = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(oracle)
EMU = Path('/var/tmp/magnolia-microemulator-2.0.4.jar')
MAGNOLIA = Path(os.environ.get('MAGNOLIA_GENERIC_DIRECT_TEST_EXECUTABLE',
    '/var/tmp/magnolia-midp-generic-refactor/magnolia'))


class UnrelatedDirectCanvasMethods(unittest.TestCase):
    def test_real_reference_is_unselected_and_no_device_keys(self) -> None:
        record = oracle.parse_reference('Headless start\n' + oracle.REFERENCE + '\n')
        self.assertEqual(record['before'], '0:false:0')
        self.assertEqual(record['trace'], oracle.TRACE)
        for key in ('selected', 'deviceInput', 'startApp', 'paint'):
            self.assertEqual(record[key], 'false')

    def test_fake_or_duplicate_reference_refuses(self) -> None:
        good = 'Headless start\n' + oracle.REFERENCE + '\n'
        for wrong in (good.replace('keyPressed(-5)=8:true:1',
                                   'keyPressed(-5)=7:true:1'),
                      good.replace('before=0:false:0', 'before=0:true:0'),
                      good.replace('keyPressed(-5)=8', 'keyPressed(50)=8'),
                      good.replace('keyReleased(-5)', 'keyRepeated(-5)'),
                      good.replace(';selected=false', ';selected=true'),
                      good.replace(';deviceInput=false', ';deviceInput=true'),
                      good.replace(';startApp=false', ';startApp=true'),
                      good.replace(';paint=false', ';paint=true'),
                      good.replace('REFERENCE:game=', 'BAD:game='),
                      good + oracle.REFERENCE + '\n',
                      'REFERENCE:failed java.lang.VerifyError\n'):
            with self.subTest(wrong=wrong[-90:]), self.assertRaises(ValueError):
                oracle.parse_reference(wrong)

    def test_no_forged_native_action_step_or_extra_results(self) -> None:
        line = 'DIRECT-GENERIC:' + ';'.join(
            key + '=' + value for key, value in oracle.NATIVE.items())
        self.assertEqual(oracle.parse_native(line + '\n'), oracle.NATIVE)
        for key, value in [('initialSteps', '45'), ('finalSteps', '46'),
                           ('trace', oracle.TRACE.replace('8:true:1', '7:true:1')),
                           ('constructed', 'ACO_XmasGifts_s'),
                           ('canvas', 'DS_Effects'), ('startApp', 'true'),
                           ('deviceInput', 'true'), ('paint', 'true')]:
            with self.subTest(key=key), self.assertRaises(ValueError):
                oracle.parse_native('DIRECT-GENERIC:' + ';'.join(
                    name + '=' + (value if name == key else item)
                    for name, item in oracle.NATIVE.items()))
        for bad in (line + '\n' + line, line + ';extra=ok',
                    line + ';trace=' + oracle.TRACE,
                    line.replace('DIRECT-GENERIC:', 'FAIL:'),
                    line + '\nBAD:paint'):
            with self.subTest(bad=bad[-80:]), self.assertRaises(ValueError):
                oracle.parse_native(bad)

    def test_changed_class_jar_is_valid_crc_different_checked_bytes(self) -> None:
        with tempfile.TemporaryDirectory() as scratch:
            before, changed = Path(scratch) / 'before.jar', Path(scratch) / 'after.jar'
            with ZipFile(before, 'w', compression=ZIP_STORED) as jar:
                for name in sorted(oracle.CLASS_SHA):
                    member = ZipInfo(name, date_time=(1980, 1, 1, 0, 0, 0))
                    member.compress_type = ZIP_STORED
                    jar.writestr(member, (oracle.COMPILED / name).read_bytes())
            oracle.changed_class_jar(before, changed)
            with ZipFile(before) as original, ZipFile(changed) as different:
                self.assertIsNone(different.testzip())
                self.assertEqual(original.namelist(), different.namelist())
                self.assertEqual(len(original.read('DirectSliceMIDlet$View.class')),
                                 len(different.read('DirectSliceMIDlet$View.class')))
                self.assertNotEqual(original.read('DirectSliceMIDlet$View.class'),
                                    different.read('DirectSliceMIDlet$View.class'))
            self.assertNotEqual(sha256(before.read_bytes()).hexdigest(),
                                sha256(changed.read_bytes()).hexdigest())

    def test_saved_gate_limits_claims(self) -> None:
        path = ROOT / 'reports/midlet-direct-generic.json'
        if not path.is_file():
            self.skipTest('reusable direct-key report not generated')
        data = json.loads(path.read_text())
        self.assertEqual(data['unchangedOwnedClassSha256'], oracle.CLASS_SHA)
        self.assertEqual(data['reference']['trace'], oracle.TRACE)
        self.assertEqual(data['magnolia']['chargedConstructorSteps'], 46)
        self.assertEqual(data['magnolia']['chargedFinalSteps'], 139)
        self.assertFalse(data['magnolia']['deviceKeyEventsDelivered'])
        self.assertFalse(data['reference']['deviceKeyEventsDelivered'])
        self.assertFalse(data['comparison']['realDeviceKeypadOrFrameParityCompared'])
        self.assertTrue(all(data['refusals'].values()))

    def test_live_hosts_changed_emulator_cannot_overwrite_report(self) -> None:
        if not EMU.is_file() or not MAGNOLIA.is_file():
            self.skipTest('pinned independent emulator or built Magnolia unavailable')
        with tempfile.TemporaryDirectory() as scratch:
            output = Path(scratch) / 'report.json'
            argv = [sys.executable, str(SOURCE), '--magnolia', str(MAGNOLIA),
                    '--microemu', str(EMU), '--output', str(output)]
            finished = subprocess.run(argv, cwd=ROOT, capture_output=True,
                                      text=True, timeout=80)
            self.assertEqual(finished.returncode, 0, finished.stderr[-650:])
            self.assertTrue(json.loads(finished.stdout)['passed'])
            good = output.read_bytes()
            self.assertEqual(json.loads(good)['magnolia']['trace'], oracle.TRACE)
            mutant = Path(scratch) / 'changed-emulator.jar'
            raw = bytearray(EMU.read_bytes())
            raw[10] ^= 1
            mutant.write_bytes(raw)
            denied = subprocess.run([sys.executable, str(SOURCE),
                '--magnolia', str(MAGNOLIA), '--microemu', str(mutant),
                '--output', str(output)], cwd=ROOT, capture_output=True,
                text=True, timeout=10)
            self.assertEqual(denied.returncode, 2)
            self.assertIn('pinned input missing or changed', denied.stderr)
            self.assertEqual(output.read_bytes(), good)


if __name__ == '__main__':
    unittest.main()
