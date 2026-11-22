# Finite historical SDK `swfparse` task — real semantic parser, not a player

`swf-sdk-dump.py` is an **owned, working Python FWS1 semantic dump utility**. It reads the *unchanged* historical SDK SWF1 demo, checks all supported tag and substructure extents **before output**, and writes ordered JSON with four decoded solid-fill vector shapes (bounds, every move/straight/curved edge in twips, style groups and transitions), three actual button definitions (12 states, symbols/layers/matrices, three inert `ActionGetURL` strings/targets), three placed symbols and transform matrices, background, one frame and End. URL fields are **data**, never followed. It does **not** render shapes, evaluate an AVM1 action, play audio, load remote content or run a Flash timeline. Its host is a Python CLI, **not** an Oak/Builder Flash runtime.

The [`sdk-swfparse-functional-oracle.py`](../sdk-swfparse-functional-oracle.py) gate **exited 0**. It independently compiled and ran the historical original C++ `swfparse.cpp` with [verified portability-only changes](sdk-swfparse-reference.md), the privately staged **599-byte real SDK FWS1** (SHA-256 `eaff967749b061e9ae8505f2582bf784c6274786d5764d4e9b3582499b625f24`), and the new standalone Python parser. Both programs actually performed their **parser/dump task** and exited 0. The gate compared **all 13 ordered top-level tags; 43 shape move/edge coordinates and edge kinds across four actual shapes; 23 solid-fill/stroke/style-state entries; all 12 button state/character/layer/matrix records; the exact six strings in three GetURL actions; all three placement transforms; and the frame/end/background**. The original dump remains private; the [machine-readable differential report](sdk-swfparse-functional.json) contains hashes, counts and conditions, not third-party source/media/dump URLs. The original C++'s 4,316-byte stdout SHA-256 is `cc1bc3aee8503c35d5e21ef91edbfa6763ce524522d19c4b0787b4f1592dcfdc`.

Six same-length or structurally damaged input cases (magic, unsupported tag, unknown AVM action, missing End, header length and wrong SHA pin) independently **exited 2 before any owned-parser output**. Three grouped owned tests also passed a self-made straight-edge shape/frame, signed curved control/anchor coordinates, and four no-artifact refusals. The CLI limits file size to 2 MB, top-level tags to 256, shapes to 64, bitfield/edge records to 4,096 per shape, palette/fill/line to 32 per group, button actions to 64, and validates character/depth references and frame count. Unsupported CWS/ZWS, later SWF versions, gradients/bitmaps, other AVM1 opcodes, sound/image/font/text, nested sprites and SWF5/ABC **fail closed**, rather than being silently counted as successfully parsed. Do **not** feed untrusted bytes to the old 1999 native C++; only the SHA-pinned validated demo is ever passed to it.

```sh
cd /home/user/magnolia
# Original media/source are separately obtained and hash-pinned OUTSIDE the repository.
python3 sdk-swfparse-functional-oracle.py \
  --source /path/to/Parser/swfparse.cpp \
  --demo /path/to/urlsubst/demo/src.swf \
  --output sdk-swfparse-functional.json
# Work with a rights-cleared FWS1 within the declared profile:
python3 swf-sdk-dump.py --input /path/to/file.swf \
  --expect-sha256 VERIFIED_INPUT_SHA256 \
  --output /path/to/new-semantic-dump.json
```

**Acceptance boundary:** this closes an **executable, reference-matched, one-original-input parser/dump slice**, not the requested *whole* historical SDK parser on arbitrary files and target hosts. The user-uploaded FWS5 has 2,140 frames, nested sprite AVM1 preloader actions, audio and graphics outside this utility. The other SDK utility tasks (`urlsubst` has a finite writer slice; `swf-adpcm` has a [finite PCM16→DefineSound writer slice](sdk-adpcm-functional.md); `swf-media` also has a separate [finite RGB→JPEG2 bitmap-asset slice](sdk-jpeg-functional.md) with no visible placement, but full sound/media, `swf-shape` and `swf-media` remain open), the Java Flash 2 player, SWF5/AVM1, AS3/AVM2 and FLA remain separate. No checklist entry is marked complete merely because this bounded input passes.
