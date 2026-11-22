# Original ABC method signature headers — still **not** ABC verification

The opt-in [`ABC16/46 ASCII constant-pool gate`](swf-as3-abc-pool-static.md) is followed by a **separate bounded `method_info` *signature-header* parser** for exactly the observed **zero-flag** entries in two pinned original games. It validates a finite count, parameter count/type indices, return type, optional name index and one-byte flags, then **stops before metadata and all class/script/trait/method-body/exception/opcode bytes**. Nonzero flags (default/optional/rest/parameter-name semantics) refuse; this does not imply AVM2 method semantics, verified ABC or a source frontend.

| SHA-pinned CWS original | ABC bytes | Method signatures counted; parameter-type references | Parsed signature end → **UNVERIFIED tail** |
| --- | ---: | ---: | ---: |
| `SnakeTutorial.swf` | 2,628 | **15; 3**, maximum 1 per method | 1,175 → **1,453 bytes** |
| `Virus.swf` | 15,916 | **94; 29**, maximum 5 per method | 5,272 → **10,644 bytes** |

[Independent original-byte differential](../swf-as3-abc-signatures-oracle.py) rechecks each CWS SHA and extracted ABC digest before Oak, independently parses the seven pool sections and every method signature in Python, and compares the opt-in Oak CLI's counts and end offsets. It **exited 0**, checked 12 malformed/out-of-profile method-table mutations on both independent readers (count, noncanonical/overflow U30, parameter/type/name indices, unmodeled flags, truncation and missing unverified tail), and refused a changed original SWF before extraction/report overwrite. A distinct small Python unit test checks that a **deliberately invalid, unverified tail** cannot be mistaken for valid bytecode. [Machine report](swf-as3-abc-signatures-current.json), SHA-256 `45905665836c98699d0ba9687aa77fc15f60b3739f05bb5ab0e17c77cc201875`.

Only ABC input **12–65,536 bytes**, at most **256 method signatures**, **32 parameters per method / 2,048 total**, canonical U30 and zero-flag method_info are admitted. Direct CLI use is **not pinned**; use the wrapper for original provenance. Reproduce offline:

```sh
python3 swf-as3-abc-signatures-oracle.py \
  --magnolia /path/to/magnolia \
  --cache /external/pinned-Snake-and-Virus-SWFs \
  --output swf-as3-abc-signatures-current.json
python3 -m unittest -q test.as3_abc_signatures_test
```

**No original ActionScript opcode ran; no frames, sprites, text, input, sound, event, timer or gameplay were compared.** Default FWS-only and separate opt-in CWS/ABC-pool gates keep their older scopes; no broad SWF-container, ABC verifier, AVM2, Flash platform, or media-completion checkbox is closed by this signature preflight.
