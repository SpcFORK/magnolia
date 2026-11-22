# New uploaded AS3JS: pinned source-to-Node reference for future Magnolia AS3 work

**As of 2026-10-26.** Per the user's direction, the *new* `as3js.zip (2).txt` is treated as a **reference implementation for a finite AS3 source-to-JavaScript subset**, not as a substitute for Magnolia's future AS3 runtime and **not** as an authoritative AVM2/Flash player. The older `as3js.zip.txt` and its separate [10/16 original → 19/19 patched baseline](as3js-baseline.md) remain pinned **unchanged**; this new upload has its **own** reproducible oracle and report.

| External archive | Bytes / SHA-256 | What is actually checked |
| --- | --- | --- |
| **New upload** `/home/user/uploads/as3js.zip (2).txt` | **483,481** / `1b029bf75077edbde997b08156aa072d6a45f87af83c88f04238485094521288` | 202 ZIP entries, 770,202 expanded bytes, CRC/membership/path/budget checks; **only 30** licensed source/runtime/test files extracted to a throwaway directory **outside** Magnolia. All `.git`, config and caches are ignored. |
| Older upload `/home/user/uploads/as3js.zip.txt` | **361,050** / `da7aa29e3741ebf1bc69e54919d54f0c5395d68a233a9e100fee1c09a1e8f60e` | Earlier independent baseline preserved; in this oracle, only the two *new* diagnostic tests are copied into the unchanged older stage as a negative comparator. |

Both contain the same pinned original MIT notice, SHA-256 `8b8d387531c43263d87bd098b37e0da4aa10ab96f0e3fe48e61a240a9ef30e99`. The new archive adds `test/diagnostics.test.js` and changes **five allowlisted files**: `AS3Parser.as`, `Main.as`, `runtime.js`, `runtime-compiled.js`, and `test/syntax.test.js`. Additional ZIP changes are `.git` metadata/objects and are neither extracted nor executed. The new parser supplies source-file and line/column diagnostics, while new syntax code fixes the bounded multiline-comment/regex behavior and adjusts an unsupported `is/as` expression diagnostic. These changes are **in the uploaded reference**, not fixes made inside Magnolia.

## Independently reproducible reference-side observations

[`as3js-runtime-reference-oracle.py`](../as3js-runtime-reference-oracle.py) **exited 0** on a full-file hash and license check. The unchanged new archive's Node v20.20.2 suite passes **18/18**; the two newly added diagnostic tests fail **0/2** against the older unchanged archive. The older original suite's previously recorded **10/16** and separately patched **19/19** are distinct historical cohorts, **not** an 18-versus-19 regression conclusion. The oracle recompiles the new `.as` compiler sources twice: both resulting compiler texts are byte-identical (**130,452 bytes**, SHA-256 `4de240101c73ca50f39a51dec2c413997b0270fe7c901221a88b37467476c875`). The uploaded runtime and **both** generated compilers each independently pass **18/18** tests and the same exported-API/entry-mode probes. This is source-to-Node self-bootstrap, **not** self-hosting in Magnolia/AVM2.

A separate owned `reference-cases.js` (SHA-256 `7ba498cd3f02bdf5c687a74ce094a8245d8c4365a3119b8977830d3639239130`) passes unchanged on all three compilers. Its exact **persisted `.as` inputs**, `Item.as` (SHA-256 `2c82c8b1c1e0d9098a964447e28c4934188012a4dfb17393e1dd0b4620037490`) and `Main.as` (SHA-256 `d0f9d16eeb2c2e316c914023d411be554ee07f7f972982211d5b2602c04a14f0`), are the **same files a future Magnolia AS3 runtime must accept** for a cross-runtime gate. It gives machine-readable **future comparator** outputs from those real `.as` source files: class initial field `17`; typed multi-declaration values `[1,"two",3]`; regex `true`; division `4`; `as String` yields `"hello"` or `null`; `is int` yields `true` for `3` and `false` for `3.5`; imported user-class `is` yields `true` for an instance and `false` for `{}`. Its emitted JavaScript is **1,543 bytes**, SHA-256 `d6d4e84400dc4adb8033099da66d9873362f3a9e206d626bf0c0a0bac8f8f35d`. A compound `is` expression is rejected rather than yielding broken output; malformed raw source reports **line 4, column 25**, and a malformed `Broken.as` reports `Broken.as:4:25`. These observations are **the candidate reference's own behavior**, not independently confirmed Flash AS3 semantics.

The oracle stages and runs only allowlisted files in temporary directories; Node children are bounded (≤8 s CPU, ≤80 MiB old-space, ≤2 GiB address space, ≤15 s wall per child) with libseccomp-denied sockets/connect, and a real `EPERM` network probe. No `npm install`, Git command, archive hook, or external script is invoked. A changed full upload hash refuses **before extraction/execution**. The [full machine-readable manifest, value/diagnostic traces and test outcomes](as3js-runtime-reference-new.json) have SHA-256 `6a266ee8d63bfaade3a3ac0919a6d12a6b164505c8b62112224612f11cf0237a`.

**How to use this reference next:** pin the same owned `.as` sources in a future Magnolia AS3 frontend/runtime, require the same outputs/errors with a clear supported-construct boundary, then use an **independent Flash/AVM2 implementation** to validate which of those source-to-JS behaviors are genuinely AS3-conformant. AS3JS's `lib/as3.js` helpers and Node tests are **not** an AVM2 bytecode verifier, display list, media clock, input host, Flash 5/AVM1 runtime or whole game. Magnolia still accepts **zero `.as` source bytes** and executes **none** of the original Snake/Virus AVM2 code. The ActionScript TODO remains **24 unchecked**; no broad checklist entry closed.

Reproduce without copying third-party files into the checkout:

```sh
python3 as3js-runtime-reference-oracle.py \
  --upload '/home/user/uploads/as3js.zip (2).txt' \
  --old-upload /home/user/uploads/as3js.zip.txt \
  --output as3js-runtime-reference-new.json
```
