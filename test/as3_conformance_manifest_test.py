"""No-download integrity and false-claim negatives for the AS3JS manifest."""
from __future__ import annotations

from copy import deepcopy
from hashlib import sha256
import importlib.util
import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / 'tools/as3-conformance-manifest-check.py'
spec = importlib.util.spec_from_file_location('as3_manifest', SCRIPT)
assert spec and spec.loader
checker = importlib.util.module_from_spec(spec)
spec.loader.exec_module(checker)
MANIFEST = json.loads((ROOT / 'reports/as3-conformance-manifest.json').read_text())
BASELINE = json.loads((ROOT / 'reports/as3js-baseline-reference.json').read_text())
ARCHIVE = Path('/home/user/uploads/as3js.zip.txt')


class ConformanceManifestIntegrity(unittest.TestCase):
    def test_pinned_source_and_reference_match_without_implementation_claim(self):
        if not ARCHIVE.is_file():
            self.skipTest('external user upload absent')
        self.assertEqual(sha256(ARCHIVE.read_bytes()).hexdigest(), checker.UPLOAD_SHA)
        self.assertEqual(sha256((ROOT / 'reports/as3-conformance-manifest.json').read_bytes()).hexdigest(),
                         checker.MANIFEST_SHA)
        self.assertIsNone(checker.audit(MANIFEST, BASELINE, ARCHIVE.read_bytes()))
        self.assertFalse(MANIFEST['referenceOracles']['magnoliaAS3ProgramExecuted'])
        self.assertEqual(MANIFEST['maximumInputAndWork']['magnoliaAS3AcceptedBytes'], 0)

    def test_forged_runtime_compiler_semantics_or_input_budget_refuse(self):
        if not ARCHIVE.is_file():
            self.skipTest('external user upload absent')
        for section, key, value in [
            ('sourceSyntax', 'noMagnoliaLexerParserOrTypedIR', False),
            ('typeSystem', 'independentAS3SemanticReference', True),
            ('builtinsAndLibraries', 'abcOrAvm2Runtime', True),
            ('emittedTargets', 'magnoliaSourceToJs', True),
            ('emittedTargets', 'swfWriter', True),
            ('maximumInputAndWork', 'magnoliaAS3AcceptedBytes', 1024),
            ('maximumInputAndWork', 'generalAS3SourceLimitEstablished', True),
            ('maximumInputAndWork', 'arbitraryHostileJsSandboxClaimed', True),
            ('referenceOracles', 'independentRealAS3OrAvm2ReferenceCompared', True),
            ('referenceOracles', 'magnoliaAS3ProgramExecuted', True),
            ('referenceArchive', 'sha256', '0' * 64),
        ]:
            candidate = deepcopy(MANIFEST)
            candidate[section][key] = value
            with self.subTest(section=section, key=key), self.assertRaises(ValueError):
                checker.audit(candidate, BASELINE, ARCHIVE.read_bytes())
        changed = deepcopy(MANIFEST)
        changed['referenceOracles']['originalNodeSuite']['passed'] = 16
        with self.assertRaises(ValueError):
            checker.audit(changed, BASELINE, ARCHIVE.read_bytes())

    def test_external_changed_archive_refuses_before_manifest_output(self):
        if not ARCHIVE.is_file():
            self.skipTest('external user upload absent')
        with tempfile.TemporaryDirectory() as tmp:
            damaged = Path(tmp) / 'as3js.zip'
            raw = bytearray(ARCHIVE.read_bytes())
            raw[50] ^= 1
            damaged.write_bytes(raw)
            result = subprocess.run([sys.executable, str(SCRIPT), '--upload', str(damaged)],
                                    cwd=ROOT, text=True, capture_output=True, timeout=10)
            self.assertEqual(result.returncode, 2)
            self.assertEqual(result.stdout, '')
            self.assertIn('hash differs', result.stderr)

    def test_unverified_manifest_edit_refuses_preflight(self):
        if not ARCHIVE.is_file():
            self.skipTest('external user upload absent')
        with tempfile.TemporaryDirectory() as tmp:
            changed = Path(tmp) / 'manifest.json'
            original = (ROOT / 'reports/as3-conformance-manifest.json').read_bytes()
            self.assertIn(b'false', original)
            changed.write_bytes(original.replace(b'false', b'true ', 1))
            result = subprocess.run([sys.executable, str(SCRIPT), '--upload', str(ARCHIVE),
                                     '--manifest', str(changed)], cwd=ROOT,
                                    text=True, capture_output=True, timeout=10)
            self.assertEqual(result.returncode, 2)
            self.assertEqual(result.stdout, '')
            self.assertIn('hash differs', result.stderr)


if __name__ == '__main__':
    unittest.main()
