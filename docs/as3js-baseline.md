# Separate AS3JS 0.3.3 source-compiler reference checkpoint

A **later, distinct user upload** has its own [pinned AS3 source/runtime comparator](as3js-runtime-reference-new.md) (18/18 unchanged Node tests, two new old-red diagnostics, exact owned source values/errors). It is **not** substituted into this older archive's 10/16→19/19 reference or treated as an AVM2/Flash player.

**This is not Magnolia ActionScript support.** The uploaded ZIP is a MIT-licensed AS3→JS self-transpiler, not a Flash/ABC/AVM2 player or an independent AS3 language implementation. Its unchanged Node suite remains **10 passed, 6 failed out of 16**; that red result is recorded, not concealed. A [small review patch](../patches/as3js-baseline.patch), applied only to a disposable outside-checkout extraction, makes those **16/16 pass** and adds three tests that are **0/3 on the original and 3/3 on the fork**. The fork is **19/19**. The [source of truth and 29-file hash manifest](as3js-baseline-reference.json) are emitted by the [bounded reproduction oracle](../as3js-baseline-oracle.py), which verifies uploaded ZIP SHA-256 `da7aa29e3741ebf1bc69e54919d54f0c5395d68a233a9e100fee1c09a1e8f60e` before extracting only allowlisted files. It excludes `.git`, config, caches, archive hooks and arbitrary dependencies; the [original MIT notice](../licenses/AS3JS-MIT.txt) accompanies the patch. **No third-party implementation or generated bundle is vendored into Magnolia.**

## Why the original six tests failed

- Both the AS3 `AS3Parser.as` source and its two existing JS runtime copies advanced one byte at a multiline comment terminator `*/`. The trailing `/` could be rescanned as a regex start, throwing `unterminated regular expression literal` while compiling the compiler's own source and preventing the CLI's intended missing-entry diagnostic. The separate declaration-end scan had a similar two-byte skip. The patch preserves the terminator in the token buffer and advances past both characters. A new class-body multiline-comment fixture fails on the original, succeeds on the fork.
- The archived `syntax.test.js` encoded **two** backslashes before a regex delimiter in its JavaScript string, yielding an invalid generated regex with `b` treated as a flag (`Invalid regular expression flags`). The patched fixture encodes a single escaped slash in the AS3 source. A second new fixture exercises that slash alongside division and a numeric type check. This is a *test-fixture correction*, not a claim to support all ActionScript regex syntax.
- The parser's `is/as` pattern did not recognize a numeric literal to the left of the operator; `value + 1 is Item` reached a generic error rather than the existing test's precise compound-expression refusal. The patch recognizes this bounded numeric island; a third new fixture rejects `2 * 3 is int` before emitting broken JS. This is not a general type-expression or AS3-coercion implementation.

The archived README is not a conformance specification: its language-limit claims lag tests that already attempt regexes, typed multi-declarations and `as/is`. The fork and the matched generated compilers pass **only the tested JS behavior**; these tests do not replace a real AS3/AVM2 reference or establish Flash value/namespace semantics. The original `runtime.js` and `runtime-compiled.js` differed; both existing copies were patched consistently, but neither was silently substituted for a genuine Flash runtime.

## Bounded bootstrap, not Magnolia self-hosting

Using pinned **Linux + Node v20.20.2**, the patched prior JS compiler compiled the archived `.as` source into a new JS compiler, which was **loaded and executed** to compile the same source again. The two resulting raw texts are **exactly identical**, **127,861 bytes**, SHA-256 `cfccac560c90d5162005d81af0c831607419c761750ec71e7bc1328415d06ab6`. Each independently passes **19/19 Node tests**. The owned bootstrap driver and entry-mode/API comparator additionally check constructor/loader exports and both `static` and `instance` entry behavior against the patched prior compiler, not only hashes.

The Python oracle uses caps on archive/member sizes, CRC, selected paths, child wall time, CPU, Node old-space/address-space, file size and output. A Linux libseccomp filter denies `socket`/`connect` and related network syscalls to **each Node child and its descendants**; an actual Node `127.0.0.1` connection probe must report `NETWORK_BLOCKED:EPERM`. No npm install or script hook runs. This is bounded **research isolation**, not a security sandbox for arbitrary hostile JavaScript. An environment without the pinned Node/libseccomp setup refuses the gate rather than claiming reproducibility. No Node package is called by Magnolia at runtime.

Reproduce from Magnolia root with the user-uploaded ZIP (the oracle itself creates/deletes all third-party extraction and generated output *outside* the checkout):

```sh
timeout --signal=TERM --kill-after=3s 120s python3 as3js-baseline-oracle.py \
  --upload /home/user/uploads/as3js.zip.txt \
  --output as3js-baseline-reference.json
```

**Separately completed documentation only:** the [bounded conformance manifest](as3-conformance-manifest.md) now accurately inventories source syntax, rights, limits, reference oracles and distinct profiles against a re-executed identical Node report. **Still open:** a genuine native Oak `.as` frontend/typed IR, AS3 semantics and independent reference; ABC/AVM2, separate AVM1, SWF container/graphics/audio/input, the uploaded SWF5/FLA and pinned Snake/Virus game playback. A Node source-compiler fixed point does **not** satisfy the later *Magnolia-native* self-hosting checkbox.
