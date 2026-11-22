"""Adversarial verification for the reusable, non-Acotel MIDP start boundary."""
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
SOURCE = ROOT / 'tools/midlet-deferred-generic-oracle.py'
SPEC = importlib.util.spec_from_file_location('midlet_deferred_generic_oracle', SOURCE)
assert SPEC and SPEC.loader
oracle = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(oracle)
EMU = Path('/var/tmp/magnolia-microemulator-2.0.4.jar')
MAGNOLIA = Path(os.environ.get('MAGNOLIA_GENERIC_DEFERRED_TEST_EXECUTABLE',
    '/var/tmp/magnolia-midp-generic-refactor/magnolia'))


def reference_line() -> str:
    return '\n'.join(['TRACE:constructed', 'TRACE:start-entered',
        'TRACE:not-alive-before-start', 'TRACE:worker-ran',
        'TRACE:start-returning', oracle.REFERENCE]) + '\n'


def native_record() -> dict:
    return {'type': 'ok', 'output': oracle.OUTPUT,
        'phase': 'start-returned-worker-deferred',
        'actions': ['construct', 'startApp'],
        'constructedClass': 'DeferredSliceMIDlet',
        'deferredRunnableClass': 'DeferredSliceMIDlet$View',
        'requestedDisplay': 'DeferredSliceMIDlet$View',
        'registeredCommands': {'count': 0, 'labels': [], 'listenerClass': None},
        'workerThreadStartCalled?': True, 'workerRunnableDeferred?': True,
        'workerRunExecuted?': False, 'modelDisplaySelected?': False,
        'modelPaintCalled?': False, 'deviceInputDelivered?': False,
        'virtualFreeBytes': None, 'runtimeGcCalls': 0,
        'steps': 75}


class GenericDeferredStart(unittest.TestCase):
    def test_reference_worker_can_race_return_but_must_run_and_select(self) -> None:
        data = oracle.parse_reference(reference_line())
        self.assertTrue(data['workerBeforeReturnTrace'])
        self.assertTrue(data['startAppReturned'])
        self.assertTrue(data['CanvasEventuallySelected'])
        self.assertEqual(data['workerState'], 7)
        after = reference_line().replace('TRACE:worker-ran\nTRACE:start-returning',
                                         'TRACE:start-returning\nTRACE:worker-ran')
        self.assertFalse(oracle.parse_reference(after)['workerBeforeReturnTrace'])

    def test_fake_reference_worker_state_order_and_duplicate_refuse(self) -> None:
        good = reference_line()
        for wrong in (good.replace('workerState=7', 'workerState=0'),
                      good.replace(';selected=true', ';selected=false'),
                      good.replace('TRACE:worker-ran\n', ''),
                      good.replace('TRACE:constructed', 'BAD:constructed'),
                      good.replace('TRACE:not-alive-before-start',
                                   'BAD:alive-before-start'),
                      good.replace(oracle.REFERENCE, oracle.REFERENCE + '\n' +
                                   oracle.REFERENCE),
                      good.replace('TRACE:constructed\nTRACE:start-entered',
                                   'TRACE:start-entered\nTRACE:constructed')):
            with self.subTest(wrong=wrong[-100:]), self.assertRaises(ValueError):
                oracle.parse_reference(wrong)

    def test_native_no_forged_worker_device_budget_or_extra_fields(self) -> None:
        data = native_record()
        self.assertEqual(oracle.parse_native(json.dumps(data)), data)
        for key, value in [('steps', 0), ('steps', 76), ('steps', True),
                           ('workerRunExecuted?', True),
                           ('modelDisplaySelected?', True),
                           ('modelPaintCalled?', True),
                           ('deviceInputDelivered?', True),
                           ('output', oracle.OUTPUT + 'TRACE:worker-ran\n'),
                           ('deferredRunnableClass', 'DS_Effects'),
                           ('requestedDisplay', 'null'),
                           ('registeredCommands', {'count': 1})]:
            with self.subTest(key=key), self.assertRaises(ValueError):
                oracle.parse_native(json.dumps({**data, key: value}))
        for raw in (json.dumps({**data, 'fakeSuccess': True}),
                    json.dumps(data) + '\n' + json.dumps(data),
                    json.dumps(data).replace('"type": "ok"',
                                             '"type": "ok", "type": "ok"')):
            with self.assertRaises(ValueError):
                oracle.parse_native(raw)

    def test_saved_report_is_owned_start_boundary_not_gameplay(self) -> None:
        path = ROOT / 'reports/midlet-deferred-generic.json'
        if not path.is_file():
            self.skipTest('metadata report not generated')
        saved = json.loads(path.read_text())
        self.assertEqual(saved['unchangedOwnedClassSha256'], oracle.CLASS_SHA)
        self.assertEqual(saved['magnolia']['chargedSteps'], 75)
        self.assertTrue(saved['reference']['realWorkerRan'])
        self.assertTrue(saved['reference']['CanvasEventuallySelected'])
        self.assertFalse(saved['magnolia']['workerRan'])
        self.assertFalse(saved['comparison']['sameWorkerExecutionOrFrame'])
        self.assertFalse(saved['comparison']['realDevicePixelOrGameplayParityCompared'])
        self.assertTrue(all(saved['refusals'].values()))

    def test_worker_started_in_constructor_or_duplicate_startApp_is_not_success(self) -> None:
        javac = Path('/usr/lib/jvm/jdk-11/bin/javac')
        if not all(path.is_file() for path in (EMU, MAGNOLIA, javac)):
            self.skipTest('pinned emulator, local Magnolia or JDK 11 missing')
        cases = (
            ('ConstructorThreadMIDlet', '', 'new Thread(this).start();', '',
             'Thread.start outside bounded monitor handoff scheduler'),
            ('ConstructorEarlyStartMIDlet', 'private int calls;', 'startApp();',
             'if (++calls == 2) new Thread(this).start();',
             'deferred MIDP startApp must be one selected lifecycle call'),
            ('RecursiveStartMIDlet', 'private int calls;', '',
             'if (++calls == 1) startApp(); else new Thread(this).start();',
             'deferred MIDP startApp must be one selected lifecycle call'),
        )
        source = ('import javax.microedition.midlet.MIDlet;\n'
                  'public final class {name} extends MIDlet implements Runnable {{\n'
                  '  {fields}\n'
                  '  public {name}() {{ {constructor} }}\n'
                  '  protected void startApp() {{ {start} }}\n'
                  '  protected void pauseApp() {{}}\n'
                  '  protected void destroyApp(boolean unconditional) {{}}\n'
                  '  public void run() {{ System.out.println("BAD: worker ran"); }}\n'
                  '}}\n')
        with tempfile.TemporaryDirectory(prefix='midp-deferred-escape-') as tmp:
            for name, fields, constructor, start, error in cases:
                with self.subTest(name=name):
                    folder = Path(tmp) / name
                    folder.mkdir()
                    file = folder / (name + '.java')
                    file.write_text(source.format(name=name, fields=fields,
                        constructor=constructor, start=start))
                    compiled = subprocess.run([str(javac), '--release', '8',
                        '-g:none', '-proc:none', '-cp', str(EMU), '-d', str(folder),
                        str(file)], cwd=ROOT, capture_output=True, text=True,
                        timeout=12)
                    self.assertEqual(compiled.returncode, 0, compiled.stderr[-450:])
                    jar = folder / 'isolated.jar'
                    with ZipFile(jar, 'w', compression=ZIP_STORED) as packed:
                        packed.write(folder / (name + '.class'), name + '.class')
                    denied = subprocess.run([str(MAGNOLIA), 'tools/midlet.oak',
                        'run', '--entry', str(jar), '--class', name,
                        '--sequence', 'start-deferred', '--expect-sha256',
                        sha256(jar.read_bytes()).hexdigest(),
                        '--max-steps', '50000', '--json'], cwd=ROOT,
                        capture_output=True, text=True, timeout=10)
                    self.assertEqual(denied.returncode, 1, denied.stderr[-450:])
                    self.assertEqual(denied.stdout.strip(), '[midlet] MIDP probe: ' + error)
                    self.assertNotIn('BAD:', denied.stdout + denied.stderr)

    def test_live_two_host_differential_and_changed_emulator_preserves_report(self) -> None:
        if not EMU.is_file() or not MAGNOLIA.is_file():
            self.skipTest('pinned emulator or built Magnolia unavailable')
        with tempfile.TemporaryDirectory() as scratch:
            root = Path(scratch)
            output = root / 'report.json'
            argv = [sys.executable, str(SOURCE), '--magnolia', str(MAGNOLIA),
                    '--microemu', str(EMU), '--output', str(output)]
            finished = subprocess.run(argv, cwd=ROOT, capture_output=True,
                                      text=True, timeout=100)
            self.assertEqual(finished.returncode, 0, finished.stderr[-650:])
            self.assertTrue(json.loads(finished.stdout)['passed'])
            good = output.read_bytes()
            self.assertEqual(json.loads(good)['magnolia']['chargedSteps'], 75)
            mutant = root / 'changed-emulator.jar'
            content = bytearray(EMU.read_bytes())
            content[10] ^= 1
            mutant.write_bytes(content)
            denied = subprocess.run(argv[:-4] + ['--microemu', str(mutant),
                '--output', str(output)], cwd=ROOT, capture_output=True,
                text=True, timeout=10)
            self.assertEqual(denied.returncode, 2)
            self.assertIn('pinned input missing or changed', denied.stderr)
            self.assertEqual(output.read_bytes(), good)


if __name__ == '__main__':
    unittest.main()
