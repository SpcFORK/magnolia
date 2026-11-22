"""Adversarial two-host tests for generic, pinned MIDP constructor-only run."""
from __future__ import annotations

import importlib.util
from hashlib import sha256
import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest
from zipfile import ZIP_STORED, ZipFile

ROOT = Path(__file__).resolve().parents[1]
SOURCE = ROOT / 'tools/midlet-constructor-generic-oracle.py'
SPEC = importlib.util.spec_from_file_location('midlet_constructor_generic_oracle', SOURCE)
assert SPEC and SPEC.loader
oracle = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(oracle)
EMU = Path('/var/tmp/magnolia-microemulator-2.0.4.jar')
MAGNOLIA = Path(os.environ.get('MAGNOLIA_GENERIC_CONSTRUCTOR_TEST_EXECUTABLE',
    '/var/tmp/magnolia-midp-generic-refactor/magnolia'))


class OwnedConstructorOnly(unittest.TestCase):
    def test_reference_has_real_target_constructor_but_no_target_start(self) -> None:
        data = oracle.parse_reference('Headless proxy\n' + oracle.REFERENCE + '\n')
        self.assertEqual(data['state'], '0:false:0')
        self.assertEqual(data['game'], 'DirectSliceMIDlet')
        for key in ('selected', 'targetStartApp', 'callbacks', 'paint', 'deviceInput'):
            self.assertEqual(data[key], 'false')

    def test_reference_cannot_forge_target_start_or_select_canvas(self) -> None:
        good = 'Headless proxy\n' + oracle.REFERENCE + '\n'
        for wrong in (good.replace('DirectSliceMIDlet$View', 'DS_Effects'),
                      good.replace('state=0:false:0', 'state=8:true:1'),
                      good.replace('displaySame=true', 'displaySame=false'),
                      good.replace('selected=false', 'selected=true'),
                      good.replace('targetStartApp=false', 'targetStartApp=true'),
                      good.replace('callbacks=false', 'callbacks=true'),
                      good.replace('paint=false', 'paint=true'),
                      good.replace('deviceInput=false', 'deviceInput=true'),
                      good + oracle.REFERENCE + '\n',
                      'CONSTRUCTOR:failed java.lang.VerifyError\n'):
            with self.subTest(wrong=wrong[-110:]), self.assertRaises(ValueError):
                oracle.parse_reference(wrong)

    def test_native_cannot_fake_lifecycle_worker_or_result_shape(self) -> None:
        good = dict(oracle.EXPECTED)
        self.assertEqual(oracle.parse_native(json.dumps(good)), good)
        for key, value in [
            ('steps', 45), ('steps', 47), ('steps', True),
            ('constructedClass', 'ACO_XmasGifts_s'),
            ('phase', 'active'), ('actions', ['construct', 'startApp']),
            ('output', 'BAD:unexpected-startApp'),
            ('requestedDisplay', 'DirectSliceMIDlet$View'),
            ('modelStartAppCalled?', True), ('modelWorkerStarted?', True),
            ('modelPaintCalled?', True), ('modelDisplaySelected?', True),
            ('deviceInputDelivered?', True),
            ('registeredCommands', {'count': 1, 'labels': [], 'listenerClass': None}),
            ('virtualFreeBytes', None), ('runtimeGcCalls', 1),
        ]:
            with self.subTest(key=key), self.assertRaises(ValueError):
                oracle.parse_native(json.dumps({**good, key: value}))
        for raw in (json.dumps({**good, 'gameplay': True}),
                    json.dumps(good) + '\n' + json.dumps(good),
                    json.dumps(good).replace('"type": "ok"',
                        '"type": "ok", "type": "ok"'),
                    'not JSON'):
            with self.assertRaises(ValueError):
                oracle.parse_native(raw)

    def test_saved_metadata_only_is_not_a_game_frame(self) -> None:
        path = ROOT / 'reports/midlet-constructor-generic.json'
        if not path.is_file():
            self.skipTest('generic constructor report not generated')
        data = json.loads(path.read_text())
        self.assertEqual(data['unchangedOwnedClassSha256'], oracle.direct.CLASS_SHA)
        self.assertEqual(data['magnolia']['chargedSteps'], 46)
        self.assertFalse(data['magnolia']['retainedSessionOrWorkerOrPaint'])
        self.assertFalse(data['magnolia']['targetStartAppOrDeviceInput'])
        self.assertFalse(data['reference']['targetStartAppCalled'])
        self.assertFalse(data['comparison']['sameDeviceFrameOrGameplayCompared'])
        self.assertTrue(all(data['refusals'].values()))

    def test_constructor_that_calls_lifecycle_worker_or_paint_refuses_before_guest(self) -> None:
        javac = Path('/usr/lib/jvm/jdk-11/bin/javac')
        if not all(path.is_file() for path in (EMU, MAGNOLIA, javac)):
            self.skipTest('pinned emulator, local Magnolia or JDK 11 missing')
        common = ('import javax.microedition.midlet.MIDlet;\n'
                  'public final class {name} extends MIDlet {interface} {{\n'
                  '  public {name}() {{ System.out.println("BAD: constructor ran"); {call} }}\n'
                  '  protected void startApp() {{ {start} }}\n'
                  '  protected void pauseApp() {{}}\n'
                  '  protected void destroyApp(boolean unconditional) {{}}\n'
                  '  {extra}\n'
                  '}}\n{companion}')
        cases = {
            'ConstructorEscapeMIDlet': dict(interface='', call='startApp();',
                start='System.out.println("BAD: startApp ran");', extra='', companion='',
                error='isolated MIDP constructor links a guest lifecycle callback'),
            'ConstructorRunsMIDlet': dict(interface=' implements Runnable', call='run();',
                start='', extra='public void run() { System.out.println("BAD: run ran"); }',
                companion='', error='isolated MIDP image links a Runnable body'),
            'ConstructorPaintMIDlet': dict(interface='',
                call='new ConstructorPaintCanvas().paint(null);', start='', extra='',
                companion=('final class ConstructorPaintCanvas extends '
                    'javax.microedition.lcdui.Canvas {\n'
                    '  protected void paint(javax.microedition.lcdui.Graphics g) '
                    '{ System.out.println("BAD: paint ran"); }\n}\n'),
                error='isolated MIDP image links a Canvas paint body'),
        }
        with tempfile.TemporaryDirectory(prefix='midp-constructor-refusals-') as tmp:
            for name, parts in cases.items():
                with self.subTest(name=name):
                    folder = Path(tmp) / name
                    folder.mkdir()
                    source = folder / (name + '.java')
                    source.write_text(common.format(name=name, **parts))
                    compiled = subprocess.run([str(javac), '--release', '8',
                        '-g:none', '-proc:none', '-cp', str(EMU), '-d', str(folder),
                        str(source)], cwd=ROOT, capture_output=True, text=True,
                        timeout=12)
                    self.assertEqual(compiled.returncode, 0, compiled.stderr[-450:])
                    classes = sorted(folder.glob('*.class'))
                    self.assertIn(name + '.class', [c.name for c in classes])
                    jar = folder / 'isolated.jar'
                    with ZipFile(jar, 'w', compression=ZIP_STORED) as packed:
                        for item in classes:
                            packed.write(item, item.name)
                    command = [str(MAGNOLIA), 'tools/midlet.oak', 'run',
                        '--entry', str(jar), '--class', name,
                        '--sequence', 'construct-only', '--expect-sha256',
                        sha256(jar.read_bytes()).hexdigest(),
                        '--max-steps', '50000', '--json']
                    denied = subprocess.run(command, cwd=ROOT, capture_output=True,
                        text=True, timeout=10)
                    self.assertEqual(denied.returncode, 1, denied.stderr[-450:])
                    self.assertEqual(denied.stdout.strip(),
                        '[midlet] MIDP probe: ' + parts['error'])
                    self.assertNotIn('BAD:', denied.stdout + denied.stderr)
                    if name != 'ConstructorEscapeMIDlet':
                        # The deferred worker boundary must not let an
                        # explicitly reached run()/paint() slip through.
                        command[command.index('construct-only')] = 'start-deferred'
                        refused = subprocess.run(command, cwd=ROOT,
                            capture_output=True, text=True, timeout=10)
                        self.assertEqual(refused.returncode, 1, refused.stderr[-450:])
                        self.assertEqual(refused.stdout.strip(),
                            '[midlet] MIDP probe: ' + parts['error'])
                        self.assertNotIn('BAD:', refused.stdout + refused.stderr)

    def test_live_two_host_differential_and_no_overwrite_by_changed_emulator(self) -> None:
        if not EMU.is_file() or not MAGNOLIA.is_file():
            self.skipTest('separately pinned emulator or local Magnolia executable missing')
        with tempfile.TemporaryDirectory() as scratch:
            report = Path(scratch) / 'constructor-only.json'
            argv = [sys.executable, str(SOURCE), '--magnolia', str(MAGNOLIA),
                    '--microemu', str(EMU), '--output', str(report)]
            run = subprocess.run(argv, cwd=ROOT, capture_output=True, text=True,
                                 timeout=90)
            self.assertEqual(run.returncode, 0, run.stderr[-650:])
            self.assertTrue(json.loads(run.stdout)['passed'])
            good = report.read_bytes()
            self.assertEqual(json.loads(good)['magnolia']['chargedSteps'], 46)
            altered = Path(scratch) / 'changed-emulator.jar'
            raw = bytearray(EMU.read_bytes())
            raw[10] ^= 1
            altered.write_bytes(raw)
            denied = subprocess.run([sys.executable, str(SOURCE),
                '--magnolia', str(MAGNOLIA), '--microemu', str(altered),
                '--output', str(report)], cwd=ROOT, capture_output=True,
                text=True, timeout=10)
            self.assertEqual(denied.returncode, 2)
            self.assertIn('pinned input missing or changed', denied.stderr)
            self.assertEqual(report.read_bytes(), good)


if __name__ == '__main__':
    unittest.main()
