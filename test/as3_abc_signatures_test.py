"""No-execution method-signature-reference tests on owned ABC-shaped data."""
from __future__ import annotations

import importlib.util
from pathlib import Path
import unittest

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location('signature_oracle',
    ROOT / 'tools/swf-as3-abc-signatures-oracle.py')
assert SPEC and SPEC.loader
mod = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(mod)


def tiny() -> bytes:
    # Seven empty pool counts, one (params=0, return=0, name=0, flags=0)
    # method_info, then one byte of UNVERIFIED data. This is not valid ABC.
    return b'\x10\x00\x2e\x00' + b'\x00' * 7 + b'\x01\x00\x00\x00\x00' + b'\x00'


class SignatureTests(unittest.TestCase):
    def test_signature_boundary_without_method_body_validation(self) -> None:
        seen = mod.SignatureReader(tiny()).scan()
        self.assertEqual((seen['methodSignatureCount'], seen['signatureTableEndOffset'],
                          seen['postSignaturesUnverifiedBytes']), (1, 16, 1))
        self.assertFalse(seen['metadataClassScriptTraitOrCodeVerified'])
        self.assertFalse(seen['AVM2Executed'])

    def test_wrong_count_param_index_flags_and_missing_tail_refuse(self) -> None:
        valid = tiny()
        mutants = (
            valid[:11] + b'\x00' + valid[12:],
            valid[:11] + b'\x81\x02' + valid[12:],
            valid[:12] + b'\x21' + valid[13:],
            valid[:13] + b'\x01' + valid[14:],
            valid[:14] + b'\x01' + valid[15:],
            valid[:15] + b'\x08' + valid[16:],
            valid[:-1],
        )
        for sample in mutants:
            with self.subTest(sample=sample.hex()), self.assertRaises(ValueError):
                mod.SignatureReader(sample).scan()


if __name__ == '__main__':
    unittest.main()
