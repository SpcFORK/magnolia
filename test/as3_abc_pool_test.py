"""Independent Python ABC pool-reader adversaries, with no historical bytes."""
from __future__ import annotations

import importlib.util
from pathlib import Path
import unittest

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location('abc_pool_oracle', ROOT / 'tools/swf-as3-abc-pool-oracle.py')
assert SPEC and SPEC.loader
ABC = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(ABC)


def smallest() -> bytes:
    # ABC16/46 and seven empty constant-pool count fields; tail deliberately
    # NOT verified as a valid method/script/trait/code section.
    return b'\x10\x00\x2e\x00' + b'\x00' * 7 + b'\x00'


class BoundedPoolTests(unittest.TestCase):
    def test_minimal_pool_remains_explicitly_unverified(self) -> None:
        observed = ABC.PoolReader(smallest()).scan()
        self.assertEqual(observed['constantPoolEndOffset'], 11)
        self.assertEqual(observed['postPoolUnverifiedBytes'], 1)
        self.assertEqual(observed['poolCounts']['multiname'], 0)
        self.assertFalse(observed['methodTraitOrOpcodeVerified'])
        self.assertFalse(observed['AVM2Executed'])

    def test_version_missing_tail_overlong_or_huge_count_refuses(self) -> None:
        samples = (
            b'\x10\x00\x2f\x00' + smallest()[4:],
            smallest()[:-1],
            smallest()[:4] + b'\x80\x00' + smallest()[5:],
            smallest()[:4] + b'\x81\x20' + smallest()[5:],
            smallest()[:4] + b'\x81',
        )
        for data in samples:
            with self.subTest(data=data[:12].hex()), self.assertRaises(ValueError):
                ABC.PoolReader(data).scan()

    def test_nonascii_and_invalid_qname_refuse(self) -> None:
        # int/uint/double zero, two strings incl one actual UTF-8 byte;
        # namespace kind 8 names string 1; no ns sets/multinames.
        text = (b'\x10\x00\x2e\x00' + b'\x00\x00\x00' +
                b'\x02\x01A' + b'\x02\x08\x01' + b'\x00\x00' + b'\x00')
        self.assertEqual(ABC.PoolReader(text).scan()['asciiStringBytes'], 1)
        for bad in (text.replace(b'\x01A', b'\x01\xff'),
                    text.replace(b'\x08\x01', b'\x08\x02'),
                    text.replace(b'\x08\x01', b'\x08\x00')):
            with self.subTest(bad=bad.hex()), self.assertRaises(ValueError):
                ABC.PoolReader(bad).scan()


if __name__ == '__main__':
    unittest.main()
