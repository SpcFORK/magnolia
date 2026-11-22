"""Adversarial boundaries for non-executing historic media reconnaissance."""
from __future__ import annotations

import importlib.util
from io import BytesIO
from pathlib import Path
import struct
import unittest
from unittest.mock import patch
from zipfile import ZIP_DEFLATED, ZipFile, ZipInfo
import zlib

ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location(
    'historic_media_inventory', ROOT / 'tools/historic-media-inventory.py')
media = importlib.util.module_from_spec(spec)
spec.loader.exec_module(media)


def swf() -> bytes:
    rect = b'\x08\x00'  # 1-bit signed rectangle, four zero coordinates
    frame = rect + struct.pack('<HH', 20 * 256, 1)
    tags = struct.pack('<HH', 1 << 6, 0)  # ShowFrame, then End
    body = frame + tags
    return b'FWS' + bytes([5]) + struct.pack('<I', 8 + len(body)) + body


def swf_tag(code: int, payload: bytes = b'') -> bytes:
    length = len(payload)
    if length < 63:
        return struct.pack('<H', (code << 6) | length) + payload
    return struct.pack('<HI', (code << 6) | 63, length) + payload


def swf_with_tags(tags: bytes, frames: int = 1) -> bytes:
    rect = b'\x08\x00'
    body = rect + struct.pack('<HH', 20 * 256, frames) + tags
    return b'FWS' + bytes([5]) + struct.pack('<I', 8 + len(body)) + body


def sprite(sprite_id: int, frames: int, tags: bytes) -> bytes:
    return swf_tag(39, struct.pack('<HH', sprite_id, frames) + tags)


def flic(magic=0xaf12, ring=False) -> bytes:
    header = bytearray(128)
    frame = struct.pack('<IHH', 16, 0xf1fa, 0) + b'\0' * 8
    tail = frame if ring else b''
    struct.pack_into('<IHHHHH', header, 0, 128 + len(frame) + len(tail),
                     magic, 1, 2, 2, 8)
    return bytes(header) + frame + tail


def bundle(names: list[str]) -> bytes:
    output = BytesIO()
    with ZipFile(output, 'w', ZIP_DEFLATED) as z:
        for name in names:
            z.writestr(name, b'a')
    return output.getvalue()


class HistoricMediaInventoryTests(unittest.TestCase):
    def assert_refused(self, fn, *args, part: str | None = None):
        with self.assertRaises(ValueError) as raised:
            fn(*args)
        if part:
            self.assertIn(part, str(raised.exception))

    def test_swf_fws_and_cws_minimal(self):
        raw = swf()
        fws = media.inspect_swf(raw)
        self.assertEqual((fws['signature'], fws['swfVersion'], fws['frames'],
                          fws['fps'], fws['tagCounts']),
                         ('FWS', 5, 1, 20.0, {'0': 1, '1': 1}))
        compressed = b'CWS' + raw[3:8] + zlib.compress(raw[8:])
        self.assertEqual(media.inspect_swf(compressed)['tagCounts'], fws['tagCounts'])
        self.assertEqual(media.inspect_swf(compressed)['topLevelAvm1ActionTags'], 0)
        self.assertEqual(media.inspect_swf(compressed)['avm2DoABC'], [])

    def test_swf_refusals(self):
        raw = swf()
        self.assert_refused(media.inspect_swf, raw[:7])
        self.assert_refused(media.inspect_swf, b'ZWS' + raw[3:], part='ZWS')
        self.assert_refused(media.inspect_swf, raw + b'\x00', part='length')
        self.assert_refused(media.inspect_swf, raw[:-2], part='length')
        self.assert_refused(media.inspect_swf, raw[:-2] + b'\x80\x00', part='End')
        self.assert_refused(media.inspect_swf, raw[:-2] + b'\xff\xff', part='dword')
        changed = bytearray(raw)
        changed[8 + 2 + 2] = 2  # header claims two frames, only one present
        self.assert_refused(media.inspect_swf, bytes(changed), part='ShowFrame')
        compressed = b'CWS' + raw[3:8] + zlib.compress(raw[8:])
        self.assert_refused(media.inspect_swf, compressed + b'junk', part='trailing')
        self.assert_refused(media.inspect_swf, compressed[:-1] + b'0', part='CWS')

    def test_nested_swf_sprite_actions_are_counted_but_not_executed(self):
        inner = swf_tag(12, b'\0') + swf_tag(1) + swf_tag(0)
        raw = swf_with_tags(sprite(11, 1, inner) + swf_tag(1) + swf_tag(0))
        item = media.inspect_swf(raw)
        self.assertEqual((item['topLevelAvm1ActionTags'],
                          item['nestedSpriteCount'], item['nestedAvm1ActionTags'],
                          item['totalTagsIncludingSprites']), (0, 1, 1, 6))
        self.assertEqual(item['nestedTagCounts'], {'0': 1, '1': 1, '12': 1})
        self.assertEqual(item['spriteSummaries'], [
            {'id': 11, 'depth': 1, 'frames': 1, 'bytes': len(inner) + 4,
             'tags': 3, 'avm1ActionTags': 1}])
        cws = b'CWS' + raw[3:8] + zlib.compress(raw[8:])
        self.assertEqual(media.inspect_swf(cws)['nestedAvm1ActionTags'], 1)

    def test_nested_swf_second_level_and_long_tag_are_bounded(self):
        show, end = swf_tag(1), swf_tag(0)
        # A long tag is valid when its extended length fits precisely within
        # the nested stream. Counts below distinguish parent/child timelines.
        grandchild = swf_tag(12, b'\0' * 63) + show + end
        child = sprite(22, 1, grandchild) + show + end
        item = media.inspect_swf(swf_with_tags(sprite(11, 1, child) +
                                               show + end))
        self.assertEqual(item['nestedSpriteCount'], 2)
        self.assertEqual(item['nestedAvm1ActionTags'], 1)
        self.assertEqual(item['nestedTagCounts']['39'], 1)
        self.assertEqual(item['totalTagsIncludingSprites'], 9)
        self.assertEqual([s['depth'] for s in item['spriteSummaries']], [1, 2])
        self.assertEqual([s['avm1ActionTags'] for s in item['spriteSummaries']],
                         [0, 1])

    def test_nested_swf_bounds_and_malformed_sprites_fail_closed(self):
        show = swf_tag(1)
        end = swf_tag(0)
        def movie(tags: bytes) -> bytes:
            return swf_with_tags(tags + show + end)
        self.assert_refused(media.inspect_swf,
                            movie(sprite(1, 1, show)), part='End tag')
        self.assert_refused(media.inspect_swf,
                            movie(sprite(1, 1, end)), part='sprite ShowFrame')
        self.assert_refused(media.inspect_swf,
                            movie(swf_tag(39, b'\x01\x00\x01')),
                            part='DefineSprite header')
        # A nested tag must not read bytes from the containing SWF timeline.
        self.assert_refused(media.inspect_swf,
                            movie(sprite(1, 1, struct.pack('<H', (12 << 6) | 62) +
                                         b'\0' + show + end)),
                            part='overflows enclosing sprite')
        self.assert_refused(media.inspect_swf,
                            movie(sprite(1, 1, end + show)), part='End tag')
        self.assert_refused(media.inspect_swf,
                            movie(sprite(1, media.MAX_FRAMES + 1, show + end)),
                            part='FrameCount')
        deep = show + end
        for i in range(media.MAX_SPRITE_DEPTH + 1):
            deep = sprite(i + 1, 1, deep) + show + end
        self.assert_refused(media.inspect_swf, swf_with_tags(deep),
                            part='depth')
        with patch.object(media, 'MAX_TAGS', 3):
            self.assert_refused(media.inspect_swf,
                                movie(sprite(1, 1, show + end)),
                                part='combined tag count')

    def test_fli_flc_headers_and_ring(self):
        self.assertEqual(media.inspect_flic(flic(0xaf11))['format'], 'FLI')
        item = media.inspect_flic(flic(0xaf12, True))
        self.assertEqual((item['format'], item['frames'], item['ringFrameBytes']),
                         ('FLC', 1, 16))
        self.assert_refused(media.inspect_flic, flic()[:-1], part='size')
        changed = bytearray(flic())
        struct.pack_into('<I', changed, 128, 1000)
        self.assert_refused(media.inspect_flic, bytes(changed), part='frame')
        changed = bytearray(flic())
        struct.pack_into('<H', changed, 8, 0)
        self.assert_refused(media.inspect_flic, bytes(changed), part='dimensions')
        self.assert_refused(media.inspect_flic, flic() + b'!', part='size')

    def test_fla_ole_not_swf(self):
        source = bytearray(512)
        source[:8] = media.OLE
        struct.pack_into('<HH', source, 30, 9, 6)
        self.assertEqual(media.inspect_fla(bytes(source))['sectorBytes'], 512)
        self.assert_refused(media.inspect_swf, bytes(source))
        struct.pack_into('<H', source, 30, 15)
        self.assert_refused(media.inspect_fla, bytes(source), part='sector')

    def test_zip_provenance_and_traversal(self):
        inspected = media.inspect_zip(bundle(['ChartLoader.as', 'run.bat']), 'source ZIP')
        self.assertEqual(inspected['entries'], 2)
        self.assertEqual(inspected['members'][1]['format'],
                         'not executed (out-of-scope helper or binary)')
        self.assert_refused(media.inspect_zip, bundle(['../breakout.as']), 'ZIP',
                            part='unsafe')
        self.assert_refused(media.inspect_zip, bundle(['a//b.as']), 'ZIP',
                            part='unsafe')
        self.assert_refused(media.inspect_zip, bundle(['a.as', 'a.as']), 'ZIP',
                            part='duplicate')
        output = BytesIO()
        with ZipFile(output, 'w') as z:
            info = ZipInfo('alias.as')
            info.create_system = 3
            info.external_attr = (0o120777 << 16)
            z.writestr(info, '../outside')
        self.assert_refused(media.inspect_zip, output.getvalue(), 'ZIP',
                            part='symlinks')


if __name__ == '__main__':
    unittest.main()
