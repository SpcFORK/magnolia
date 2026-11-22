# Uploaded SWF5: verified *first-frame* timeline prefix, not movie playback

**Status (2026-10-26):** Magnolia now accepts the **entire original 2,418,588-byte FWS5 file** as input to a bounded, **unrendered virtual first-frame execution**. The new Oak [`lib/avm1-swf-timeline-prefix.oak`](../lib/avm1-swf-timeline-prefix.oak) first runs the complete finite SWF envelope scanner (**7,394** verified top/nested tags, **2,140** root `ShowFrame`s at **20 fps**, four sprite definitions). It follows the actual initial `PlaceObject2` hierarchy and executes the *original binary* sprite-11 **frame-0 `DoAction`** through the separate bounded AVM1 engine. Unlike the [prior direct-action test](avm1-preloader-original.md), this invocation uses the **whole SWF and its real placement/first-frame tags**, not extracted action files. The direct CLI itself is unpinned; only the independent [Python oracle](../avm1-swf-timeline-prefix-oracle.py) SHA-pins the whole uploaded input (`2bde4d96b5b6695f8dd81afe005080c2bc6b9a14545f7cc8f35f49c71ac43cbc`) **before** invoking it.

| Original first-frame structure actually read | Absolute byte offset | Bounded result |
| --- | ---: | --- |
| Root `SoundStreamHead` | **26** | MP3 format flags, 2,205 stream samples/block; metadata only. |
| Sprite 11 frame-0 original `DoAction` payload | **1966** (201 bytes) | AVM1 runs with injected, finite loaded/total byte counts. |
| Sprite 11 places sprite 6, named `bar`, depth 3 | **2185** | Validates the target of its `_xscale` effect; no image pixels. |
| Sprite 12 places sprite 11 at depth 2 | **2301** | Establishes the nested preloader display-list chain. |
| Root frame 0 places sprite 12 at depth 3 | **2322** | Instantiates the virtual chain from the movie's own placement tags. |
| Root first `SoundStreamBlock` | **2334** | 1,570-byte tag, 3,456 declared samples and 1,673 seek samples; **not decoded or played**. |
| Root first `ShowFrame` | **3910** | First logical frame entered, not rasterized. Root `_root.stop()` event applied; no root frame 1 entered. |
| Sprite 11 frame-1 second action payload | **2226** (7 bytes) | Located but **not executed/scheduled by this prefix**. The earlier *separate direct-action test* does execute it. |

With a **synthetic** byte snapshot **49/200**, original actions write `PercentLoaded=24.5`, `bar._xscale=24.5` and stop the virtual root. At **100/100**, they write `PercentLoaded=100` and *request* the `loaded` label (sprite frame 2) without writing a bar scale; **frame 2 is not entered**. `0/100` and `1/4` provide two additional paths. These numbers do **not** come from real network/file-loading progress. The implementation scans/places all applicable first-frame objects before dispatching the original frame-0 action tag, a deliberately bounded virtual schedule; it has **not** been compared with a version-compatible Flash player's action ordering.

The independent oracle checks the full-file SHA and structural frame/tag counts, separately walks root frame-0 and nested sprites/label/placements, compares *every* output field and original ActionRecord offset against a manual fixed-body AVM1 calculation, and verifies **nine atomic one-line refusals**: altered root/sprite/bar references, first sound block count, bad action opcode, an **action-budget self-loop after earlier guest effects**, bad snapshots/options and oversize. It rejects an independently changed full-file hash **before** asking Magnolia to execute it. [Machine-readable first-frame report](avm1-swf-timeline-first-frame.json), exit **0**, SHA-256 `6f79daddc6f8d496b4519a51c24f54bad75ef8829182155753f6254ad44b0f92`. Its Python expectations are **independent code**, but **not observations from an independent Flash player**.

**Explicit boundary:** this is one virtual first-frame prefix, *not* automatic second-tick/movie scheduling, nor a player. The remaining **2,139 root frames** are structurally counted, **not entered**. No later nested sprite frame is entered, no real bytes-loaded host is measured, no MP3 stream is decoded/audible, no shapes/text are rasterized, and no click/key/input or AVM1 general object model is present. The original project's `.fla` upload is **not proved to export this SWF**. The Java Flash 2 `.class` program and AVM2/AS3 Snake/Virus inputs are separate lanes. No broad ActionScript checklist item is complete.

Reproduce, keeping the upload outside the checkout:

```sh
go build -o /var/tmp/magnolia-avm1-built/magnolia .
python3 avm1-preloader-oracle.py \
  --magnolia /var/tmp/magnolia-avm1-built/magnolia \
  --swf /home/user/uploads/asdf.swf.txt \
  --output avm1-preloader-original.json
python3 avm1-swf-timeline-prefix-oracle.py \
  --magnolia /var/tmp/magnolia-avm1-built/magnolia \
  --swf /home/user/uploads/asdf.swf.txt \
  --output avm1-swf-timeline-first-frame.json
```

**Next hard gates:** obtain a legally usable, version-pinned Flash5/AVM1 reference **before claiming chronological or visual parity**; implement bounded sprite clocks and `stop`/`play`/`gotoFrame`/`gotoLabel` effects across further frames and completed load states; parse actual transforms/shape/text/bitmap geometry and verify reference pixels; decode/compare streamed MP3 audio; deliver real input and complete a full original-media trace. Refuse unsupported semantic paths rather than misreporting a static scan as playback.
