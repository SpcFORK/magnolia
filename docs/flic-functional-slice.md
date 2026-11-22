# Animator FLI/FLC: a real, bounded pixel player (separate from Flash)

**Observed result:** `flic-functional-oracle.py` **exited 0** against an independently built, commit-pinned Go FLIC decoder. `flic-playback.py` actually decompressed and rendered **482/482 displayed RGB frames / 33,408,000 pixels** of six unchanged externally pinned FLI/FLC files; it also produced six independently pixel-compared PNG frames and an offline HTML animation whose timer/play/pause/step/seek controls executed in Node. This is **not** the original Java Fliplay program, a general FLI/FLC compatibility claim, a SWF/FLA/AVM player, or completion of every media blocker. [Full per-frame metadata and comparison evidence](flic-functional-slice.json).

| External original, original source | Magic, displayed frames | Covered behavior / exact input SHA-256 |
| --- | --- | --- |
| [`a.fli`](https://samples.ffmpeg.org/fli-flc/a.fli) | genuine **FLI** 0xAF11, 384, plus matching ring | 320×200, Color64, ByteRun, DeltaFLI, 5/70 s clock; `0ecd22c21f7d48cbcc3473ec374b00f5b5814d690142e7c3648279cd244a5421` |
| [`num3.fli`](https://samples.ffmpeg.org/fli-flc/num3.fli) | actually **FLC** 0xAF12 despite `.fli` name, 6, plus matching ring | 640×480, Color256, ByteRun, DeltaFLC, 72 ms clock; `5626298e5b3a71ef7fa830701116ec4be115dff3d94654d4d932601f0a77a6d2` |
| [`ship.flc`](https://github.com/samuel/go-flic/blob/002496ead18cbde762080150cf60579f48b440b0/testdata/ship.flc) | FLC, 2, plus matching ring | 640×480, prefix, opaque bounded postage stamp, Color256, ByteRun, DeltaFLC; `983601130540f567fe7e8c28e199fced6ebc8ae2fbb76b7f89fa218383a49466` |
| [`alex.flc`](https://github.com/samuel/go-flic/blob/002496ead18cbde762080150cf60579f48b440b0/testdata/alex.flc) | FLC, **3** real displayed frames though header claims 2 | 128×100, odd-width last-pixel DeltaFLC; the extra frame is **not** silently discarded as a ring; `f61ab650ffb0b81a3e803c8f6e279c82ac0b8d27bcbf81808586efef0a70c212` |
| User's ZIP-extracted `moonclouds.flc` | FLC, 60, plus matching ring | 320×240, Color64, ByteRun, DeltaFLC; raw `812379899d151a1460366249fd794899eeefe3c7bacdc528da22c2624b33cff4`, separate uploaded ZIP `15b5d795fd96ae4cf32be8610745fff21611407803b2d3e59b3a430f4971c564` |
| [`2422.FLC`](https://samples.ffmpeg.org/fli-flc/2422.FLC) | FLC, 27, plus matching ring | 320×200, 171 ms clock, bounded prefix, Color256/ByteRun/DeltaFLC; `2dfe6613b3ca0d5bbfbf82010b7e4234c7f26e8eb44d67071fef20ffa3a494c5`. **Standalone only**: original JDK Fliplay rejects its prefix as invalid frame magic. |

**Reference independence:** the oracle builds [Go FLIC at commit `002496e`](https://github.com/samuel/go-flic/tree/002496ead18cbde762080150cf60579f48b440b0) with Go 1.26 in a temporary module *outside Magnolia*, SHA-checking its unmodified `flic.go` (`f62e689a…`), `render.go` (`e0cb9c64…`), and MIT `LICENSE` (`3a8946fe…`). Only Magnolia's **owned 1,288-byte Go adapter** (`flic-go-reference-main.go.txt`) converts that reference's rendered `image.Paletted` frames to RGB hashes; it is not part of the production decoder. The independently built reference, Magnolia CLI, input originals and JDK/Fliplay methods remain distinct. The oracle compares all 482 per-frame RGB SHA-256 digests and independently checks every PNG pixel against the corresponding reference frame. A second independent Pillow FLI plugin agrees on all six `num3` frames and two displayed `alex` frames, **not** on `a.fli` (its 6-bit Color64 palette scaling differs from the Go reference) and cannot decode `ship` or `2422` frame 0; we do not claim three-way parity for these cases.

`flic-playback.py` validates the **entire file and every frame before writing**: bounded file size (2 MB), 8-bit FLI/FLC magic/depth/flags, 320×200 FLI and finite FLC geometry (≤307,200 pixels), 512 declared frames / 40 million frame-pixel work, subchunk length/type/count and counted palette/scanline operations. Only Color64/Color256, ByteRun, DeltaFLI/DeltaFLC, Black and FLICopy affect pixels. A bounded prefix and postage-stamp envelope may be inspected without substituting its thumbnail for a frame. Known historical writers leave at most two unused bytes after a counted delta; the report **counts** these bounded suffix bytes. The decoder rejects audio/script/unknown/extended FLIC chunks, 16/24-bit variants, damaged headers, oversized work, missing/unknown chunks, malformed runs and ambiguous file tails before output. ZIP extraction is **not** implicit in the CLI: the user-uploaded Moon ZIP was separately byte/CRC/entry/hash checked, and its raw original matched the pinned input; no arbitrary archive member or path is followed.

**Real browser check (private derivative):** [`flic-browser-player-oracle.py`](../flic-browser-player-oracle.py) opened the user's separately pinned 60-frame Moon Clouds offline player (SHA-256 `91c16ecb3335ff5aa9f587bef0ee2608a1041109f1924a49b19d636cfc91e4b5`) in external **Google Chrome for Testing 142.0.7444.175** (binary SHA-256 `b0faf712c0f41d033305ae5759cec9a0ee9db66f571e54b364d96c022028fe6a`), with no external resources. Chrome **exited 0**, decoded a real **320×240** embedded frame, actually advanced the playback counter to **4/60** under a virtual 220-ms delay, then paused, stepped and sought **60/60**. The test script is injected only into a private temporary *copy*; the user's player is unchanged. A wrong player SHA pin **exited 1 before browser launch/report output**. [Metadata-only browser report](flic-moon-browser.json). This checks one browser run, **not physical wall-clock fidelity, every frame's browser compositing, or general FLIC**.

The **separate original Java Fliplay host** has its own [second-original-movie JDK differential](fliplay-alex-original-jdk.md): `alex.flc` reaches three displayed original-class states that Magnolia matches to JDK byte-for-byte, even though Fliplay's later pixels *disagree* with the correct independent Go FLIC decoding. `num3.fli` also has its own [third, finite original-class JDK/Magnolia gate](fliplay-num3-original-jdk.md): **7/7** 640×480 JDK RGB states (six movie plus ring) and final ARGB match, via a separately SHA-pinned caller-only 307,840-cell array / 307,200-pixel capture opt-in; the interpreter's 65,536-cell default array cap and separate generated-Java 131,072-cell admission ceiling remain unchanged. `2422.FLC` is rejected by original JDK Fliplay at its prefix, despite the standalone decoder matching Go. These outcomes must not be attributed to the standalone decoder or silently promoted into general original-JVM support. The Moon Clouds 61-state original gate was rerun after the new capture bound and again matched **61/61** JDK RGB states and the complete final ARGB; it does not prove 61 physical X11 captures.

The real original [`phasers.fli`](https://github.com/samuel/go-flic/blob/002496ead18cbde762080150cf60579f48b440b0/testdata/phasers.fli) (SHA-256 `a5255ae219d292bbd3d1e0fa4a3f139735a7fe3e7fdd4ba2a58fec8eaa72cf5c`) is **damaged/inconsistent**: it was not rescued by resync or reported as successful playback. It and four other mutations (header, unsupported chunk, depth, pin) **exited 2 before any HTML/report output**; seven grouped owned codec/output/refusal tests separately passed. Unknown visuals, timing on a real device, sound, unusual palettes/rings, higher depth, other FLIC extensions, signed Windows output and the JVM Fliplay broad-host gate remain open. The new offline HTML is intentionally a playback of cached **decoded pixels**, not original JVM execution or a SWF/ActionScript timeline.

## Run it with original media supplied separately

```sh
cd /home/user/magnolia
# Paths are to unchanged external files; none are committed or downloaded by these tools.
python3 flic-functional-oracle.py \
  --a-fli /path/to/a.fli --num3-flc /path/to/num3.fli \
  --ship-flc /path/to/ship.flc --alex-flc /path/to/alex.flc \
  --moonclouds-flc /path/to/extracted/moonclouds.flc \
  --2422-flc /path/to/2422.FLC \
  --go-source-dir /path/to/pinned/go-flic/ \
  --damaged-phasers /path/to/phasers.fli \
  --output flic-functional-slice.json
# A new, private self-contained player or PNG directory for rights-cleared input:
python3 flic-playback.py --input /path/to/file.flc \
  --expect-sha256 VERIFIED_INPUT_SHA256 --html /path/to/new-player.html
node flic-html-player-oracle.js /path/to/new-player.html FRAME_COUNT
# Optional actual Chromium image/timer/control check (external browser binary):
python3 flic-browser-player-oracle.py --html /path/to/new-player.html \
  --expect-sha256 VERIFIED_HTML_SHA256 --chrome /path/to/chrome \
  --frames FRAME_COUNT --width PIXEL_WIDTH --height PIXEL_HEIGHT \
  --output /path/to/new-browser-check.json
```

The actual user-uploaded Moon FLC also generated a **60-frame single-file HTML player** with embedded pixels, no network requests and separate play/pause/seek controls. Keep any derivative private unless its underlying artwork rights are cleared. The report and source patch contain only hashes/metadata/owned code, **no original FLI/FLC or Go-library source**. This slice is separate from the three original-Fliplay/JDK movie gates and does not turn the Flash-2 Java player, Java ME game, SWF5/AVM1, AS3/AVM2, FLA authoring, other SDK utilities, Falstad, SJG, Jmol, Jalview or archive games into functioning media.
