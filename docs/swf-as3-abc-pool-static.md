# Original AS3 DoABC constant-pool *structure* — no traits, methods or AVM2

A separate opt-in [`ABC 16/46 diagnostic`](../lib/swf-as3-abc-pool.oak) now inspects seven constant-pool sections inside the **two SHA-pinned original** CWS games. The independently implemented [Python oracle](../swf-as3-abc-pool-oracle.py) verifies the entire original SWF hash **before** extracting a `DoABC` payload to a temporary external directory, checks its CWS/tag envelope independently, and compares the resulting ABC pool counts and end offset with the Oak CLI. No original ABC is committed. [Machine-readable differential](swf-as3-abc-pools-current.json), SHA-256 `c20050e7db0139f92bcb0150bc4fb098130445ed9086b1f7904ad733a149ff4e`.

| Pinned game | Exact ABC bytes (SHA in report) | Int/uint/double/string | Namespace / set / multiname | QName / Multiname / MultinameL | Pool-end offset → **unverified tail** |
| --- | ---: | ---: | ---: | ---: | ---: |
| `SnakeTutorial.swf` | 2,628 | 2 / 0 / 0 / 84 | 21 / 3 / 71 | 64 / 5 / 1 | 1,111 → **1,517 bytes** |
| `Virus.swf` | 15,916 | 11 / 0 / 13 / 321 | 40 / 6 / 273 | 239 / 31 / 2 | 4,866 → **11,050 bytes** |

The named `pool-end offset` is **not** a verified method offset. All bytes after it remain unparsed and may be invalid or dangerous: method signatures, class/script/trait tables, exceptions, code bodies and opcodes have **not** been validated or run. Pool strings are explicitly **ASCII only** (751 / 3,345 checked bytes); non-ASCII UTF-8, more multiname variants, namespace-resolution semantics, TypeName, proxy/dynamic access, Unicode, full AS3 source syntax and all Flash APIs are **outside this profile**. The parser checks canonical U30/U32, 12–65,536 input bytes, at most 4,096 entries per pool, 4,096 bytes per string / 32,768 total string bytes, 256 namespaces per set and bounded local indices; it never invents classes or interprets bytecode. The oracle passed **ten separately mutated**, independently refused original-pool variants (version, counts, overlong varint, length, non-ASCII, namespace kind/index, multiname kind/index, missing tail), plus a changed original-SWF hash refusal before extraction or report overwrite. Separate synthetic Python unit tests cover the no-method-verification boundary.

```sh
python3 swf-as3-abc-pool-oracle.py \
  --magnolia /path/to/magnolia \
  --cache /external/pinned-Snake-and-Virus-SWFs \
  --output swf-as3-abc-pools-current.json
python3 -m unittest -q test.as3_abc_pool_test
```

The prior opt-in [CWS zlib/tag/DoABC **header** gate](swf-as3-cws-static.md) and existing default FWS-only scanner are separate. **Only the ABC constant-pool boundary advanced** here; no true ABC verification, AVM2 interpreter, Snake/Virus rendering, audio, input, gameplay or actual ActionScript source execution was established. Neither the broad ABC/AVM2 nor SWF container/playback checkbox may be checked on this evidence; [census](media-blocker-census.md) keeps first barriers and runtime blockers non-additive.
