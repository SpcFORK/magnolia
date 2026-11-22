# Original ABC16/46 class/trait/body **structure** — opcodes remain opaque

A new **opt-in**, deliberately narrow Oak [`ABC structure reader`](../lib/swf-as3-abc-structure.oak) follows the independently checked [ASCII constant pool](swf-as3-abc-pool-static.md) and [zero-flag method signature headers](swf-as3-abc-signatures-static.md). On the **same two SHA-pinned, unchanged original CWS SWFs**, it validates bounded metadata/class/instance/script/trait and method-body *table shape*, signed-off offsets and index ranges. It accounts for all ABC bytes **only by skipping each declared method's opaque code span**. Neither the Oak reader nor the separate Python reader decodes even one opcode; neither verifies stack/scope transitions, exceptions, trait semantics or the Flash API, and **neither executes AVM2 or a game**.

| External original | Structural class / script tables | Trait headers (slot, method, class, const) | Bounded method-body envelopes; **opaque unverified opcode bytes** | ABC end |
| --- | ---: | ---: | ---: | ---: |
| `SnakeTutorial.swf` | 3 / 3 | **20** (10, 6, 3, 1) | **15; 1,202** | 2,628 / 2,628 |
| `Virus.swf` | 15 / 15 | **199** (110, 49, 15, 25) | **94; 8,677** | 15,916 / 15,916 |

[The independent original-byte differential](../swf-as3-abc-structure-oracle.py) checks the *entire SWF SHA-256 before extraction*, independently extracts the DoABC payload and checks the original ABC digest, then compares a separately implemented Python reader with Oak on counts, bounded indices, offsets, trait-kind tallies, code-span lengths, limits and exact end of the ABC. **25 independently rejected** malformed/out-of-profile table variants exercise metadata counts, noncanonical/overflow U30, missing/extra classes, names/namespace/flags/interfaces, trait kinds/attributes/defaults, initializer indices, body count/indices, stack/local caps, code lengths, exception count, truncation and trailing bytes. A same-length altered original SWF refuses before Oak/report replacement. The oracle explicitly replaces an original code byte with the invalid opcode `0xFF` and confirms both *structural* readers still accept it **solely as opaque bytes**. This is a negative counterexample to any accidental "verified ABC" claim, not a passing executable.

[Machine report](swf-as3-abc-structure-current.json), SHA-256 **`d7548281336ba3c5e00fbf452134ca5a3fb5f174c63762951429938de3ec0410`**. The direct Oak CLI operates on unpinned, strictly local ABC (12–65,536 bytes); the wrapper enforces original provenance. Resource profile: up to 32 classes/scripts, 128 traits per block / 256 total, 1–256 ascending one-to-one method bodies, stack and locals ≤32, per-body code ≤2,048 bytes and total opaque code ≤16,384 bytes. This deliberately rejects other **valid** ABCs: nonzero metadata/trait attributes/exceptions/interfaces, flags outside the two observed class-flag values (8 and 9), unsupported trait kinds, non-integer slot defaults, missing or reordered body methods, and strings/pools/method headers outside the older gates. These rejections reflect this finite profile, not AVM2-invalid input.

Reproduce with the previously pinned SWFs held **outside** this repository:

```sh
python3 swf-as3-abc-structure-oracle.py \
  --magnolia /path/to/magnolia \
  --cache /external/pinned-Snake-and-Virus-SWFs \
  --output swf-as3-abc-structure-current.json
python3 -m unittest -q test.as3_abc_structure_test
```

**Next real gates:** verify *every* encoded opcode, operand index, branch target, stack/local/scope and exception behavior; implement AVM2 with independent runtime-value/error traces, then actual Flash display/event/time/input/audio and frame/pixel/gameplay comparisons. Parsing an `instance_info` or skipping a `code_length` is **not** linking or executing its class. The broad ABC/AVM2, SWF container and platform checkboxes stay open; the cross-media [census](media-blocker-census.md) links this structural checkpoint without rerunning it or counting the 8,677 / 1,202 opaque bytes as verified behavior.
