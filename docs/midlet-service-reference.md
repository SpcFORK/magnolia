# Bounded two-frame MIDP `serviceRepaints()` — owned fixture only

**New narrow positive gate, not a Java ME game:** with **both** `--canvas-paint --canvas-service --sequence start`, Magnolia executes one initial selected `showNotify`/`paint`, then invokes the same original **owned fixture** classfiles' preflighted public `triggerService()V` method. That method calls `repaint()` **once**, then `serviceRepaints()`; a **distinct opt-in** owned Canvas classfile consumes the request and synchronously executes one more preflighted `paint(Graphics)` on the same metered 176×176 frame. The default Canvas still has no `repaint`, and the earlier one-frame opt-in still has **no** `serviceRepaints`. There is no concurrent MIDP event thread or general queue.

The same fixture `.class` bytes were built once by JDK 11 `javac --release 8 -g:none` against the **real**, SHA-pinned MicroEmulator 2.0.4 MIDP API, then executed unmodified through real JDK/MicroEmulator and both Magnolia input routes (local owned JAR and separate class siblings). Three real-emulator runs and both native routes agree on the **entire** first and second 176×176 Canvas frames:

| Frame | Colored rectangle | Colored/black pixels | Big-endian ARGB SHA-256 (30,976 pixels) |
| --- | --- | ---: | --- |
| First | red, x=10..25/y=12..20 | 144 / 30,832 | `321b60d5f3feeb85c1bcf083b8f590712ccbd4d3c43b0c079e4f71bfcc50634e` |
| Second | green, x=32..44/y=41..47 | 91 / 30,885 | `5d4255653aa9125037a98983971e9bcb3eb292b6719810e1cd28488ea8697a74` |

Native Magnolia used **1,697 metered steps**, printed `start → detached → selected → shown → paint-one → request-returned → paint-two → service-returned`, and returned **two distinct frame digests**. Real MicroEmulator sometimes emits `paint-two` **before** the caller prints `request-returned`, and sometimes after: only the partial-order boundary **`paint-two` before `service-returned`** is asserted. Native runs a deterministic *synchronous* subset, not that concurrent schedule. Reference observation separately verifies exactly **one** `showNotify`, two paints, 176-pixel clip and zero paints while detached.

**Pixel-observer correction:** MicroEmulator's J2SE `MutableImage.getData()` retains a `PixelGrabber` from its first read and can return that **stale first-frame array** on the second read. The checked real-emulator observer instead takes a **fresh `BufferedImage.getRGB(...)` snapshot of the device image after each boundary**, checks every canvas pixel, excludes the 22-row device chrome above/below, and hashes actual ARGB bytes. Ten exploratory real-emulator runs and the oracle's three independent pinned runs reached the same two frame digests. The two-frame digest is not derived from a cached/synthetic PNG or from the Magnolia hash.

The [executable oracle](../midlet-service-oracle.py) pins the external MicroEmulator JAR (643,996 bytes; SHA-256 `dbd5f3eb8365d3e839d6a203149e0e3776fc1a0585e16ac1fc23f76c9fcae1c6`) and unchanged Acotel JAR (50,594 bytes; SHA-256 `c24e61a416616a036b537957ec9b0687ad08f9fceae18e4d06336883ec8c3040`) **outside** the repository. It checks owned [fixture sources](../test/fixtures/midlet-service/PaintServiceMIDlet.java), [reference observer](../test/historical/midlet-service/PaintServiceReference.java), [service Canvas source](../lib/midp1-service-src/javax/microedition/lcdui/Canvas.java), newly compiled fixture/model class hashes, the separately shipped Canvas.class, two native input routes, retained one-frame behavior, and **nine refusals**. The service-classes folder ships **only** the owned Canvas.class; the javac-only `ImageBridge` stub class is never shipped. [Full structured trace and hashes](midlet-service-reference.json); [pure bound/negative tests](../test/midlet_service_oracle_test.py).

The negatives refuse an unmodeled `Display.getCurrent()` in the **extra method before MIDlet construction/output**, `--canvas-service` without explicit paint, full lifecycle under paint, a **second** service call, and an insufficient VM work budget. The default route still refuses the fixture's `repaint()`. The original Acotel JAR remains **pre-construction**: default first barrier is `java/lang/Runtime`; the **one-frame paint plus measured-memory** route still stops at missing `Canvas.serviceRepaints()V`. If the **fixture-specific service flag** is applied to Acotel, preflight refuses its lack of the required `triggerService()V`; this is **not** evidence that original gameplay advanced or a new intrinsic game blocker. The snapshot value 31,119,560 bytes comes from an independent real-emulator fixture and is **never** implicit.

**Not claimed:** original Acotel execution; repeated or concurrent event-thread paint, request coalescing under arbitrary schedules, `hideNotify`/selection changes, partial clip/device variants, glyph drawing, physical key/softkey input, full CLDC/MIDP, a Flash/ActionScript runtime, or generated JS/HTML/Java target parity. The next original-game gate is a genuine preflighted callback/thread/display policy plus independently compared original frames/input—not a hard-coded `triggerService` alias. This bounded gate does **not** close an open JVM/ActionScript TODO entry. [Older one-frame gate](midlet-paint-reference.md), [original-reference scheduling probe](midlet-repaint-reference.md), [cross-media census](media-blocker-census.md).

Reproduce from Magnolia root, with separately obtained pinned input files and the existing native Magnolia executable:

```sh
timeout --signal=TERM --kill-after=3s 150s python3 midlet-service-oracle.py \
  --microemu /path/to/microemulator-2.0.4.jar \
  --original-jar /path/to/ACO_XmasGifts_s.jar \
  --magnolia /path/to/magnolia --output midlet-service-reference.json
python3 -m unittest -v test/midlet_service_oracle_test.py
```

No JAD URL, network, third-party executable, PHP/BAT/EXE or original-game JAR content is bundled or followed by this gate.
