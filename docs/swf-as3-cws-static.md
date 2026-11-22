# Opt-in AS3 CWS v9/v15 container envelopes — no ABC verification or Flash game

The existing [`FWS` 1–7 scanner](../swf-tag-scan.oak) remains **FWS-only**: it still refuses each unchanged CWS SWF before tag scanning. A **separate opt-in** [`lib/swf-as3-cws.oak`](../lib/swf-as3-cws.oak) and [`inspect` CLI](../swf-as3-cws-scan.oak) now consume independently pinned original CWS bytes with finite zlib/Adler32, sprite, tag, frame and `FileAttributes`/`DoABC` **header** checks. It rejects ZWS, versions outside **9/15**, non-AS3 headers, extra ABC headers, and corrupt/trailing compression before returning metadata. This is not a general SWF container or permission to evaluate ABC.

[Independent byte-pinned differential](../swf-as3-cws-oracle.py) rechecked the two previously obtained, **unchanged** author/preservation-hosted files outside the repository, invoked both the independent Python SWF parser and the Oak scanner, and **exited 0** after comparing all top/nested tags, sprite summaries, frame counts, attributes, exact `DoABC` headers and eight pre-metadata refusals. A deliberately changed original SHA refused **before Oak invocation and without replacing the successful report**. [Machine report](swf-as3-cws-current.json) SHA-256 `4bbea677026e1719e8bacffb1d170351344eaad883d796e88c10b5372d329e43`.

| Unchanged external source | CWS compressed → declared bytes | Top + nested tags; sprites | ABC **header only**, 16/46 | State |
| --- | ---: | ---: | ---: | --- |
| `SnakeTutorial.swf`, SHA-256 `f07cc3422d04bfa49713cb71a2db8a55c92a6c258f85fc94c782d6ff76c7f10c` | 2,461 → 4,732 | **22 + 6 = 28; 2** | **2,628** bytes | CWS v15 envelope, no ABC opcode/trait parse |
| `Virus.swf`, SHA-256 `0332ebd3b8f7110af6298633d2d87504f4bba8156b1b61b449498f6e31a36120` | 171,197 → 236,578 | **303 + 3,867 = 4,170; 58** | **15,916** bytes | CWS v9 envelope, no ABC opcode/trait parse |

The CLI limits **compressed input to 500,000 bytes**, **declared inflation to 1,000,000 bytes**, **10,000 combined tags**, **128 sprites**, **depth 8**, and **4,096 frames**, verifies the zlib dictionary/header, DEFLATE consumption and Adler32, requires exact timeline End/ShowFrame counts and a single top-level AS3 16/46 ABC envelope. The offline oracle additionally enforces original SHA/size, read-only staged copies, bounded children and output. The direct CLI is **not provenance-pinned**; only use the separate oracle for historical claims. Reproduce offline with a built Magnolia and both separately hash-checked originals in a cache **outside** the checkout:

```sh
python3 swf-as3-cws-oracle.py \
  --magnolia /path/to/magnolia \
  --cache /path/outside/checkout/pinned-swf-cache \
  --output swf-as3-cws-current.json
python3 -m unittest -q test.media_blocker_census_test
```

The old FWS-only CLI's first refusal and its original SWF5 scanner results remain unchanged; the new path merely advances **two original CWS *metadata* inspections** to the next dependency. **Still blocking actual games:** bounded ABC pool/trait/opcode parsing and verification, AVM2 execution, Flash display/event/graphics/text/sound/loading, deterministic frame/input oracles, and supported host/export routes. No original ActionScript executed, no Snake/Virus pixels/audio/input were compared, no SWF5/AVM1 or `.fla` authoring project was advanced, and the broad SWF-container/AVM2/Flash checkboxes remain unchecked. Do not add these tag counts to the independent-runtime-blocker total; [census](media-blocker-census.md) links but does not rerun this differential.
