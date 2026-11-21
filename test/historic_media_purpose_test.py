"""Adversarial, source-only gates for bounded historic-media purpose parsing."""
from __future__ import annotations

from hashlib import sha256
from io import BytesIO
from pathlib import Path
import importlib.util
import json
import struct
import subprocess
import sys
import tempfile
import unittest
from zipfile import ZipFile

ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location(
    'historic_media_purpose', ROOT / 'tools/historic-media-purpose.py')
assert spec and spec.loader
purpose = importlib.util.module_from_spec(spec)
spec.loader.exec_module(purpose)


def tag(kind: int, payload: bytes = b'') -> bytes:
    if len(payload) < 63:
        return struct.pack('<H', (kind << 6) | len(payload)) + payload
    return struct.pack('<HI', (kind << 6) | 63, len(payload)) + payload


def swf_with_sprite(action: bytes) -> bytes:
    inner = tag(12, action) + tag(1) + tag(0)
    sprite = tag(39, struct.pack('<HH', 11, 1) + inner)
    sound = tag(18, b'\x0f\x2f' + struct.pack('<Hh', 2205, 0))
    body = b'\x08\x00' + struct.pack('<HH', 20 * 256, 1)
    body += sprite + sound + tag(19, b'\0') + tag(1) + tag(0)
    return b'FWS\x05' + struct.pack('<I', 8 + len(body)) + body


def cfb_fixture() -> bytes:
    """Small independent v3 OLE: FAT, directory, miniFAT, root and 4KiB stream."""
    raw = bytearray(512 * 13)  # header plus sectors 0..11
    raw[:8] = purpose.inventory.OLE
    struct.pack_into('<HHHH', raw, 24, 0x3e, 3, 0xfffe, 9)
    struct.pack_into('<H', raw, 32, 6)
    for off, value in [(40, 0), (44, 1), (48, 1), (56, 4096),
                       (60, 2), (64, 1), (68, purpose.END), (72, 0)]:
        struct.pack_into('<I', raw, off, value)
    struct.pack_into('<109I', raw, 76, 0, *([purpose.FREE] * 108))
    fat = [purpose.FREE] * 128
    fat[:4] = [purpose.FAT_SECTOR, purpose.END, purpose.END, purpose.END]
    for i in range(4, 11):
        fat[i] = i + 1
    fat[11] = purpose.END
    struct.pack_into('<128I', raw, 512, *fat)
    struct.pack_into('<128I', raw, 512 * 3, purpose.END,
                     *([purpose.FREE] * 127))

    def entry(slot: int, name: str, kind: int, *, right: int = purpose.FREE,
              child: int = purpose.FREE, start: int = purpose.END,
              size: int = 0) -> None:
        offset = 512 * 2 + slot * 128
        text = name.encode('utf-16-le') + b'\0\0'
        raw[offset:offset + len(text)] = text
        struct.pack_into('<H', raw, offset + 64, len(text))
        raw[offset + 66] = kind
        struct.pack_into('<III', raw, offset + 68, purpose.FREE, right, child)
        struct.pack_into('<I', raw, offset + 116, start)
        struct.pack_into('<Q', raw, offset + 120, size)

    entry(0, 'Root Entry', 5, child=1, start=3, size=64)
    entry(1, 'Small', 2, right=2, start=0, size=3)
    entry(2, 'Large', 2, start=4, size=4096)
    return bytes(raw)


def one_chart() -> str:
    lines = [purpose.CHART_INIT,
             "_global.charts['tutorial'] = {easy:[], normal:[], hard:[], bpm:100};"]
    for diff in ('easy', 'normal', 'hard'):
        lines.extend([
            f"_global.charts['tutorial'].{diff}.noteCount = 1;",
            f"_global.charts['tutorial'].{diff}.push({{t:100, s:'opponent', d:'up'}});",
            f"_global.charts['tutorial'].{diff}.push({{t:101, s:'player', d:'left'}});",
        ])
    return '\n'.join(lines) + '\n'


class PurposeParserTests(unittest.TestCase):
    def reject(self, fn, *args, has: str | None = None):
        with self.assertRaises(ValueError) as raised:
            fn(*args)
        if has:
            self.assertIn(has, str(raised.exception))

    def test_action_envelope_and_valid_synthetic_sprite_sound(self):
        pool = b'\x88' + struct.pack('<H', 4) + b'\x01\0A\0'
        pushed = b'\x96' + struct.pack('<H', 2) + b'\x08\0'
        cond = b'\x9d\x02\0\0\0'
        actions = pool + pushed + cond + b'\0'
        decoded = purpose.action_records(actions)
        self.assertEqual([r['name'] for r in decoded['records']],
                         ['ConstantPool', 'Push', 'If', 'End'])
        self.assertEqual(decoded['records'][1]['values'],
                         [{'type': 'constant', 'index': 0, 'value': 'A'}])
        self.assertEqual(decoded['records'][2]['branchTarget'], len(actions) - 1)
        meta = purpose.swf_actions(swf_with_sprite(actions))
        self.assertEqual((meta['nestedSprites'], meta['soundStreamBlocks'],
                          meta['oneSoundBlockForEveryFrame'],
                          meta['opaqueActionRecords']), (1, 1, True, 0))
        self.assertEqual(meta['audioHeads'][0]['codec'], 'MP3')

    def test_action_invalid_envelope_offsets_and_value_types(self):
        self.reject(purpose.action_records, b'\x06', has='missing End')
        self.reject(purpose.action_records, b'\0\x06', has='not final')
        self.reject(purpose.action_records, b'\x96\x02', has='header')
        self.reject(purpose.action_records, b'\x96\x06\0\0', has='crosses')
        self.reject(purpose.action_records,
                    b'\x88\x03\0\x01\0a\0\0', has='missing or oversized')
        self.reject(purpose.action_records,
                    b'\x96\x02\0\x08\0\0', has='constant index')
        self.reject(purpose.action_records,
                    b'\x96\x01\0\xff\0', has='unknown AVM1 Push')
        self.reject(purpose.action_records,
                    b'\x9d\x02\0\xfc\xff\0', has='boundary')
        self.reject(purpose.action_records,
                    b'\x99\x01\0\0\0', has='branch operand length')

    def test_swf_invalid_sprite_and_compressed_profile_refuse(self):
        valid = swf_with_sprite(b'\0')
        bad = valid[:-2] + b'\x01\x00'
        self.reject(purpose.swf_actions, bad, has='overflows')
        self.reject(purpose.swf_actions, valid[:3] + valid[3:8] + b'!',
                    has='outside bounded')

    def test_cfb_small_regular_mini_streams_and_broken_chains(self):
        valid = cfb_fixture()
        parsed = purpose.ole_directory(valid)
        self.assertEqual((parsed['streams'], parsed['reachableEntries'],
                          parsed['fatSectors'], parsed['miniFatSectors']),
                         (2, 3, 1, 1))
        self.assertEqual({s['allocation'] for s in parsed['largestStreams']},
                         {'FAT', 'miniFAT'})
        self.reject(purpose.ole_directory, valid[:-1], has='alignment')
        alias = bytearray(valid)
        struct.pack_into('<I', alias, 512 * 2 + 128 + 116, 4)
        struct.pack_into('<Q', alias, 512 * 2 + 128 + 120, 4096)
        self.reject(purpose.ole_directory, bytes(alias), has='sector collision')
        for off, value, part in [
            (512 + 4, 1, 'directory chain'),
            (512 * 3, 0, 'mini stream chain'),
            (512 * 2 + 128 + 116, 1, 'mini stream chain'),
            (512 * 2 + 256 + 116, 1, 'regular stream chain length'),
            (512, purpose.END, 'sector type marker'),
            (512 * 2 + 76, 3, 'directory pointer'),
        ]:
            altered = bytearray(valid)
            struct.pack_into('<I', altered, off, value)
            self.reject(purpose.ole_directory, bytes(altered), has=part)

    def test_cfb_difat_budget_and_sector_outside_bounds(self):
        raw = bytearray(cfb_fixture())
        struct.pack_into('<I', raw, 72, 65)
        self.reject(purpose.ole_directory, bytes(raw), has='count')
        raw = bytearray(cfb_fixture())
        struct.pack_into('<I', raw, 48, 512)
        self.reject(purpose.ole_directory, bytes(raw), has='outside')
        raw = bytearray(cfb_fixture())
        struct.pack_into('<I', raw, 512 + 4 * 5, 512)
        self.reject(purpose.ole_directory, bytes(raw), has='FAT next')

    def test_literal_chart_and_unparsed_code_refusal(self):
        result = purpose.chart_source(one_chart(), 'tutorial')
        self.assertEqual(result['difficulties']['hard']['playerNotes'], 1)
        self.assertEqual(result['difficulties']['hard']['opponentNotes'], 1)
        self.reject(purpose.chart_source, one_chart().replace('noteCount = 1',
                    'noteCount = 2', 1), 'tutorial', has='declared noteCount')
        self.reject(purpose.chart_source, one_chart() + 'eval(_root.foo);\n',
                    'tutorial', has='outside literal-data')
        self.reject(purpose.chart_source, one_chart().replace('t:101', 't:0', 1),
                    'tutorial', has='time/order')
        self.reject(purpose.chart_source, one_chart(), 'level 1',
                    has='level/bpm')

    def test_real_inputs_hashes_static_results_and_no_claimed_link(self):
        base = Path('/home/user/uploads')
        swf = purpose.inventory.checked(base / 'asdf.swf.txt',
                purpose.inventory.MAX_MEDIA, purpose.inventory.PINNED['legacy_swf'])
        bundle = purpose.inventory.checked(base / 'rapper_rhythm.zip.txt',
                purpose.inventory.MAX_INPUT, purpose.inventory.PINNED['fla_sources'])
        flash = purpose.swf_actions(swf)
        self.assertEqual((len(flash['actionTags']), flash['totalActionRecords'],
                          flash['opaqueActionRecords'], flash['soundStreamBlocks']),
                         (2, 36, 0, 2140))
        self.assertEqual(flash['actionTags'][0]['spritePath'], [11])
        self.assertEqual(flash['actionTags'][0]['records'][-2]['label'], 'loaded')
        source = purpose.authoring_source(bundle)
        self.assertEqual((source['projectDirectory']['streams'],
                          source['projectDirectory']['reachableEntries']),
                         (325, 326))
        self.assertEqual(source['relationToSeparatelyUploadedSWF'], 'not established')
        self.assertEqual(source['eventSource']['literalTriggerFunctions'], 287)
        self.assertEqual(source['chartLevels']['level 3']['difficulties']
                         ['hard']['playerNotes'], 387)
        with ZipFile(BytesIO(bundle)) as z:
            self.assertEqual(sha256(z.read('rapper_rhythm/rapper_rhythm.fla'))
                             .hexdigest(),
                             '3e5ef093065cedc618d14d8d26146bb992381e50b1b6f4d300ccafe7986b566a')

    def test_cli_bad_hash_does_not_emit_or_overwrite_report(self):
        with tempfile.TemporaryDirectory(prefix='purpose-negative-') as folder:
            folder = Path(folder)
            wrong = folder / 'wrong.swf'
            wrong.write_bytes(b'FWS\x05' + b'\0' * 8)
            output = folder / 'keep.json'
            output.write_text('unchanged\n')
            argv = [sys.executable, str(ROOT / 'tools/historic-media-purpose.py'),
                    '--swf', str(wrong), '--fla-bundle',
                    '/home/user/uploads/rapper_rhythm.zip.txt', '--flic-zip',
                    '/home/user/uploads/moonclouds.zip.txt', '--output', str(output)]
            result = subprocess.run(argv, capture_output=True, text=True,
                                    timeout=8, cwd=ROOT, check=False)
            self.assertEqual(result.returncode, 1)
            self.assertIn('SHA-256 differs', result.stderr)
            self.assertEqual(output.read_text(), 'unchanged\n')
            output.unlink()
            result = subprocess.run(argv, capture_output=True, text=True,
                                    timeout=8, cwd=ROOT, check=False)
            self.assertEqual(result.returncode, 1)
            self.assertFalse(output.exists())


if __name__ == '__main__':
    unittest.main()
