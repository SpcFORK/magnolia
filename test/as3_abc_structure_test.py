"""Synthetic structural-envelope tests; NO AVM2 opcode validation/execution."""
from __future__ import annotations

import importlib.util
from pathlib import Path
import unittest

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location('structure_oracle',
    ROOT / 'tools/swf-as3-abc-structure-oracle.py')
assert SPEC and SPEC.loader
mod = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(mod)


def fake_abc(opcode: int = 0xff) -> bytes:
    # Only ABC table shape: one string, namespace and QName, one zero-flag
    # signature, one class, one script, one 1-byte method body. 0xff is NOT a
    # valid AVM2 opcode; this deliberately exercises the opaque-code boundary.
    pool = (b'\x00\x00\x00'  # int, uint, double counts
            b'\x02\x01A'       # string pool[1] = A
            b'\x02\x08\x01'   # public namespace indexed by A
            b'\x00'             # namespace-set count
            b'\x02\x07\x01\x01')  # QName(ns1, name1)
    header = b'\x10\x00\x2e\x00' + pool
    method = b'\x01\x00\x00\x00\x00'  # count 1, zero parameters, flag 0
    meta_classes = (b'\x00\x01'      # 0 metadata, 1 class
                    b'\x01\x00\x08\x01\x00\x00\x00'  # one instance
                    b'\x00\x00'      # one class_info
                    b'\x01\x00\x00')  # one script_info
    body = bytes((1, 0, 1, 1, 0, 1, 1, opcode, 0, 0))
    return header + method + meta_classes + body


class StructureTests(unittest.TestCase):
    def test_opaque_invalid_opcode_not_mistaken_for_verified_code(self) -> None:
        raw = fake_abc()
        seen = mod.StructureReader(raw).scan()
        self.assertEqual((seen['classCount'], seen['scriptCount'],
                          seen['methodBodyCount'], seen['opaqueOpcodeBytes']),
                         (1, 1, 1, 1))
        self.assertEqual(seen['ABCStructuralEndOffset'], len(raw))
        self.assertEqual(seen['opcodeBytesDecodedOrVerified'], 0)
        self.assertFalse(seen['methodBodyInstructionsOrScopesVerified'])
        self.assertFalse(seen['AVM2OrOriginalGameExecuted'])
        for bad_byte in (0x00, 0x42, 0xfe):
            with self.subTest(opaque_byte=bad_byte):
                self.assertEqual(mod.StructureReader(fake_abc(bad_byte)).scan(), seen)

    def test_header_and_body_structure_refusals(self) -> None:
        original = fake_abc()
        r = mod.StructureReader(original)
        r.scan()
        def mutant(where: str, substitute: bytes) -> bytes:
            return mod.splice(original, r.offsets[where], 1, substitute)
        bad = {
            'metadata': mutant('metadata-count', b'\x01'),
            'class-count': mutant('class-count', b'\x00'),
            'class-name': mutant('first-instance-name', b'\x00'),
            'class-flag': mutant('first-instance-flags', b'\x00'),
            'protected-namespace': mutant('first-protected-namespace', b'\x00'),
            'script-initializer': mutant('first-script-init', b'\x01'),
            'body-count': mutant('body-count', b'\x00'),
            'body-stack': mutant('first-body-max-stack', b'\x00'),
            'body-length': mutant('first-body-code-length', b'\x81\x10'),
            'exception-count': mutant('first-body-exception-count', b'\x01'),
            'trailing': original + b'\x00',
            'truncated': original[:-1],
        }
        for label, raw in bad.items():
            with self.subTest(label=label), self.assertRaises(ValueError):
                mod.StructureReader(raw).scan()

    def test_trait_tags_references_and_default_kind(self) -> None:
        original = fake_abc()
        probe = mod.StructureReader(original)
        probe.scan()
        count_at = probe.offsets['first-instance-trait-count']
        # Add one slot with QName #1, slot_id 0, type ANY, int #1. Since
        # the int pool has count 0, a nonzero vindex must be refused; valid
        # no-default path can still carry an invalid opcode as opaque bytes.
        added = mod.splice(original, count_at, 1,
                           b'\x01\x01\x00\x00\x00\x00')
        seen = mod.StructureReader(added).scan()
        self.assertEqual(seen['traitKindCounts']['slot'], 1)
        for label, changed in {
            'unknown-kind': mod.splice(added, count_at + 2, 1, b'\x0f'),
            'unknown-attributes': mod.splice(added, count_at + 2, 1, b'\x40'),
            'out-of-range-class-name': mod.splice(added, count_at + 1, 1, b'\x02'),
            'invalid-int-pool-default': mod.splice(added, count_at + 5, 1, b'\x01'),
        }.items():
            with self.subTest(label=label), self.assertRaises(ValueError):
                mod.StructureReader(changed).scan()


if __name__ == '__main__':
    unittest.main()
