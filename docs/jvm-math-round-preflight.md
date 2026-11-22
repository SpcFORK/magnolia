# Bounded JVM `Math.round(float)` and `Math.round(double)` — exact-primitive gate

This is a **completed, narrowly scoped primitive feature**, not a completed historical program or a checklist closure. Magnolia admits only `java/lang/Math.round(F)I` and `round(D)J` on the exact static JVM opcode. The interpreter uses tagged IEEE raw bits, an exact significand/remainder comparison for half ties toward positive infinity, and its existing signed-long limb conversion for integral values/saturation. Neither native Oak nor the bundled JavaScript coerces the double result through a lossy JS Number. The generated-Java route uses the matching real JDK overloads. Other `Math` methods, overloads, virtual invocation, unbounded work and general floating-point application semantics are **not** implied.

## Independent executable checks

- [`MathRoundMain.java`](../test/fixtures/MathRoundMain.java) is Java 8 source compiled by **JDK 11** (`javac --release 8 -g:none -proc:none`), yielding the SHA-256 pinned [`MathRoundMain.class`](../test/fixtures/MathRoundMain.class), `fbfa7fcce327f963d3f00efb75d3c87fe4935c5f5d30e78d3c2b0d42186b4160`. Its 56-line [`JDK -Xverify:all` oracle](../test/fixtures/MathRoundMain.expected) hashes to `e9b31332b6fa0facfc8d9178df3a014315830ba473dd97d1de8ae5f5a4925269`. It includes signed zero, subnormals, NaN/infinities, half ties, adjacent representable values, and saturation near both integer and long limits.
- The [live differential](../applet-math-round-oracle.py) **exited 0**: it recompiled the exact checked class and re-ran its JDK oracle; the **same class bytes** produced the same 56 output lines in Magnolia native, standalone Oak, bundled JS/Node, offline HTML (both scripts evaluated without network) and generated Java/JDK. The standalone Oak route restored the exact original class bytes. A deterministic seeded Java 8 raw-bit class compared an additional **396 float and 592 double cases** across JDK, Magnolia native and bundled JS. [Machine-readable report](jvm-math-round.json) SHA-256 `25bd180003b3836f1325e1546b6213d1fbf9107d97c2ba04dbf10f392a5994c1` records checksums, targets and scope; it does not embed either third-party archive or an unexecuted program.
- [`applet-math-round-smoke.oak`](../applet-math-round-smoke.oak) independently checks exact owner/descriptor/static opcode, 56 native results and finite partial-budget refusal. An unchanged class with one static invocation opcode changed to virtual is rejected by both JDK and Magnolia **before fixture output**; unknown round overloads and a five-step run do not claim a completed result. The shared MIDP device profile is now inlined into standalone Oak builds rather than importing an adjacent source file; this fixes packaging, **not** MIDP device behavior.

Reproduction (requires a locally built Magnolia, JDK 11 `java`/`javac`/`javap` and Node; original external jar **must be separately obtained** at the pinned SHA and kept outside the checkout):

```sh
python3 applet-math-round-oracle.py \
  --magnolia /path/to/magnolia \
  --fourier /path/outside/checkout/fourier.jar \
  --output jvm-math-round.json
python3 media-blocker-census.py \
  --magnolia /path/to/magnolia \
  --acotel /path/outside/checkout/ACO_XmasGifts_s.jar \
  --cache /path/outside/checkout/pinned-census-cache \
  --output media-blockers-current.json
```

## Regression results on the same built executable

The focused Oak Applet suite **exited 0: 497/497**. Two older assertions initially failed because they still expected the previously added `Rectangle` constructor and `System.getProperty` to be unsupported; those test assertions were corrected to check the existing narrow admissions and unrelated negative owner/descriptor cases, then the full suite reran green. The combined owned/original MIDP constructor, deferred-start and direct-callback suite **exited 0: 35/35** with every live reference pointed to the same rebuilt Magnolia; its worker/frame/device-input limits are unchanged. A separate census/ablation/one-frame/two-frame metadata and adversarial suite passed **29/29**. `go test ./...` and `git diff --check` passed.

The pinned historical Flash/Fliplay core-method matrix also **exited 0**, matching the actual original JVM bytes and independently executed JDK values through native, standalone Oak, JS, HTML where applicable and generated Java; its old assumption that Fliplay's **35/35 static decode** must fail was corrected to the observed zero-blocker static inventory. This does **not** establish general Fliplay movies. The pre-existing long→float regression **exited 0** on native/Oak/JS/HTML/Java with its 387 seeded midpoint samples and invalid conversion refusal. Its former **23-second build-only deadline** was too short for the shared-profile-inlined standalone source (measured **29 seconds**); builder-only limits were raised to **75 seconds**, with **no VM execution, guest-step or admission budget widened**. Historical checks run separately from the new `Math.round` oracle rather than implying that any Flash or Falstad movie ran.

## What the unchanged original proves — and does not

The separate SHA-256-pinned **original** Falstad Fourier JAR (`712d1d09f14aace449e8049cb49517bce19749ffa0401e06d0e33ba3ae8135d7`) was inspected by all-method Java SE **static** coverage. At the immediately preceding Class/finally checkpoint its first-entry inventory was **45/87 decoded bodies; 79 first-blocker entries**. After the bounded round addition the **pinned census exited 0** at **46/87; 78**: `FourierFrame.doQuantize()V` now *decodes*, not executes. That one-body/one-first-entry change is a static diagnostic, not a count of independent remaining runtime bugs. Other bodies still encounter `Math.sin(D)D` and `Math.cos(D)D`; separate earlier diagnostic observations **43/87; 83** predated multiple intervening changes and cannot all be credited to rounding. The original selected `Fourier` entry still exits **1 before construction**, with `Fourier.addComponentListener(Ljava/awt/event/ComponentListener;)V` missing. **No Fourier transform, drawing, user interaction or original Fourier `Math.round` method was executed in Magnolia.**

Flash 2 Java still refuses before construction at `sun/audio/AudioPlayer.player`; Fliplay and Java ME retain distinct profiles and oracles; AS3 CWS/AVM2 and SWF5/AVM1 are not affected by this Java SE primitive. See [all-media non-additive census](media-blocker-census.md) and [functional closure ledger](historic-functional-gates.md). The **23 JVM and 25 ActionScript open checklist entries** remain open; the independent remaining-runtime-blocker total remains **unknown** rather than the sum of static entries.
