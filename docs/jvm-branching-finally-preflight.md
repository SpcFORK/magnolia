# One bounded branching `jsr`/`ret` finally — static Bitmap gain, not Flash playback

**Result (2026-10-26).** A separately obtained, **85,800-byte** original `FlashApi.ZIP` is pinned to SHA-256 `fbb2149179c1c5a32bd9aa7f479a1e760e36e18dbd8e48147729ca67822932c2`. `javap` of its JVM **45.3** `com.macromedia.flash.Bitmap` (`Bitmap.class`: 6,019 bytes, SHA-256 `692242b5913a104662658cda798bb935af353a35a0a5709428c088ae8c7b1b85`) shows one synchronized, obfuscated void cleanup/wait method with **two `jsr` sites at bytes 55 and 60**, a common return-address `astore_2` at byte 65, conditional branches at bytes 80/87/118, a **backedge to byte 103**, and `ret 2` at byte 121. Its exception table also covers the wait path. This is archive-derived *structural evidence*, not execution of that original method. JDK 11 `-Xverify:all` loaded the original class without invoking it. The original selected `Flash` Applet remains blocked **before construction** at the same unrelated audio field:

```text
[applet] com/macromedia/flash/Sound.stop()V at byte 36: unsupported field sun/audio/AudioPlayer.player
```

## What the bounded inliner actually admits

`lib/applet-bytecode.oak` now inlines **one** legacy finally per selected method, class major **45–50**, with **1–16 `jsr` callers**, exactly **one `ret`**, an immediately saved, **private** return-address local and a **2–33-decoded-instruction** subroutine. Interior conditional branches/gotos/backedges may point only to **decoded boundaries after its entry `astore` through its `ret`**. It clones each body and *rebases every internal edge to that caller's own copy*, then replaces entry `astore` with a metered `nop` and `ret` with a goto to that caller's actual continuation. The original `jsr` still costs an instruction step. The normal stack, primitive/reference frame, handler and **8,192-instruction expansion** bounds still run *after* lowering; a loop is metered by the ordinary VM instruction budget.

Branches into the subroutine from outside (except `jsr`), branches escaping its body or reentering the return-address `astore`, non-boundary targets, switches within the body, nested subroutines, abrupt returns/throws within the body, frame declarations, overlapping handler cuts, use of the return-address local as a normal JVM value, and **two-slot `long`/`double` local overlap** all refuse *before* Oak/JS/HTML/Java artifact emission. A JDK-valid **`wide ret`** also remains unsupported. This is not a general `jsr`/`ret` verifier, arbitrary finally expansion, monitor scheduler or AWT image producer.

## Independent owned-bytecode differential

The **owned** [`BranchingFinallyMain.java`](../test/fixtures/BranchingFinallyMain.java) is compiled with `javac --release 8 -g:none -proc:none`; the **intermediate** Java 8 class SHA-256 is `cd4a01d7b61f88e3a92379e7811447cfbd5a8a350e407f9521db5edb0e3b8d61`. The reproducible [`applet-legacy-finally-fixture.py`](../applet-legacy-finally-fixture.py) replaces only this owned `checked(int)` Code attribute with a genuine branching/looping Java 5-era `jsr`/`ret` body. The pinned [class-49 fixture](../test/fixtures/BranchingFinallyMain.class) SHA-256 is **`d07b54f9726834e1f5753affc1c9216bf6afb4ce65dd932412e38427595acfc4`**; the same bytes with header version changed to the original archive's **45.3** have SHA-256 `d74aca0f0e5312801590c3a5f9490861b4ca74e34a29dcdd6cf7b37d3f6f3dd6`. The earlier *linear* `LegacyFinallyMain.class` still regenerates byte-identically, SHA-256 `05c7b19cd776131cb4ebc30a8effff071d54b86c0fcb7215452147d2d87a323f`.

JDK 11 **`-Xverify:all`** executed the Java 8 source class, patched class 49, and patched class 45 **separately**; all three produced the same pinned [stdout](../test/fixtures/BranchingFinallyMain.expected) (SHA-256 `f841d45328c0046fa4390fdb2340851032c3ae27c5844df1c3b31aeb68f05c08`):

```text
3
16
8
0
0
1
12
```

The normal path runs the positive-value loop, the exception handler runs the finally once, and `checked(0)` follows an **actual conditional branch directly to `ret`** instead of applying `+10`. Magnolia matches the JDK bytes/output at both class versions, optimized and non-optimized native Oak, bundled Builder JS/Node, and selected standalone Oak/JS/**network-free offline HTML**. Its native and bundled smoke also inspect **two distinct rebased backedges**, **two distinct branch-to-ret copies**, final static tally, absence of a return-address IR opcode and a **28-step** budget refusal before any program stdout. The focused [oracle](../applet-branching-finally-oracle.sh) **exited 0** ([saved log](jvm-branching-finally-oracle.log), SHA-256 `819e5e485dcd6953a2f50be9d015fa40ddfe8fd52df8f44de8685b1372d76575`).

**Refusal is part of the check.** Eight *different* mutations (`branch-escape`, `branch-head`, `branch-entry`, `branch-nonboundary`, `overlap-ret-slot`, `wrong-ret-slot`, `version-51`, `handler-cross`) all cause actual JDK 11 `VerifyError`s and Magnolia selected-preflight refusals. A ninth mutated body with `wide ret` is **JDK-valid and runs with the reference stdout**, but Magnolia deliberately refuses it. Every mutation refuses on **all four emitters** (`.oak`, `.js`, `.html`, `.java`) without producing an artifact. Even the *valid* inlined fixture refuses generated Java: that separate backend does not model return addresses and must not silently delegate this operation to a full host JDK.

## Separate original-archive inventory and remaining gates

An all-method, **static-only** inventory of the unchanged original archive moves from the prior [Class checkpoint](jvm-flash-after-class-probe-static-coverage.json) **259/277 decoded bodies, 30 first-blocker entries** to [this report](jvm-flash-after-branching-finally-static-coverage.json) **260/277, 29** (saved JSON SHA-256 `efc203223ee330ff48ed7af0eac73a3c16b36c14a76e66af5263707021279e98`). Comparing `(member, message)` **multisets**, not JSON order, removes **only** the earlier first message on the obfuscated `Bitmap` cleanup method (`byte 80: bounded legacy jsr/ret: nested/branching/abrupt finally is not admitted`). It adds **no** new first blocker. **This is not execution of the original `Bitmap` method**, image production, decoding, audio or an original movie frame.

The [guarded cross-media census](media-blockers-current.json) **exited 0** with exactly the Flash row's **259/277, 30 → 260/277, 29** change; Fourier **45/87, 79**, StringWave **42/80, 78** and all other rows and original selected refusals remain unchanged. Seven synthetic bounds/multiset tests also pass ([saved census/check log](jvm-branching-finally-cross-media-check.log), SHA-256 `3fdfb9b9e6b27aafe248d994a7c77a73a04aac8721c38b04c7b1e32399da2f4d`). The earlier linear six-family, Class-token, bounded-wait and five-family oracles **each separately exited 0** during this continuation ([their collected output](jvm-branching-finally-affected-oracles.log)); those runs preceded the *final* owned `ifeq`-to-`ret` fixture and two-slot-alias guard, which the subsequent focused oracle, unit checks and static multiset check covered. At this checkpoint the **twelve-oracle combined driver had not been rerun**. A **later, separate** attempt with a freshly built binary and SHA-pinned Flash ZIP ran the combined driver under a 1,200-second outer cap, completing the first seven focused gates and reaching the eighth (`next-blocker`) standalone-output checks before **timing out with exit 124**; no twelve-gate exit 0 is claimed. [Explicit capped partial log](jvm-flash-twelve-oracle-TIMED-OUT.log). The new guarded cross-media census **separately exited 0** again with the same report hash and seven unit checks. The earlier eleven-oracle log ended in PASS but its combined exit code was **not captured**; do not turn that historical terminal line into an observed exit 0.

Original `Bitmap`'s `Image.getSource()`/`ImageProducer` dependencies still block other bodies; the existing headless raster **snapshot** is not a general `ImageProducer` (JDK `ToolkitImage` and `BufferedImage` return different source objects). `ImageSource` still reaches unmodeled `URL.toString()`, and `SCharacterParser` still reaches unmodeled `Math.cos(D)D` (and later `Math.sin(D)D`). The **`sun.audio`** field blocks original selected execution before construction. These are independent further gates; there is no general trigonometry, URL formatter, asynchronous AWT imaging, audio backend, whole Flash-2 movie playback or modern SWF execution here. Original third-party class bytes remain **outside** the Magnolia workspace.

To reproduce with a separately obtained and SHA-pinned original archive and JDK 11:

```sh
MAGNOLIA=/path/to/magnolia \
FLASH_ZIP=/path/to/FlashApi.ZIP \
  bash applet-branching-finally-oracle.sh
python3 -m unittest -v test.media_blocker_census_test
# The other original archives and Acotel JAR are separately pinned inputs:
python3 media-blocker-census.py --magnolia "$MAGNOLIA" \
  --acotel /path/to/ACO_XmasGifts_s.jar --cache /outside/magnolia/cache \
  --output media-blockers-current.json
```
