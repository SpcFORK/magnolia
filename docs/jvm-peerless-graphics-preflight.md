# Peerless `Component.getGraphics()` — null is not a Flash drawing surface

**Historical Graphics checkpoint.** The [later five-member differential](jvm-five-blocker-preflight.md) first superseded the selected clock refusal below with intermediate legacy `jsr` byte 355 and **237/277, 52** static entries; [six further bounded gates](jvm-next-blocker-preflight.md) later reported selected `Object.wait()V` byte 4 and **245/277, 44** static; [five later archive-derived families](jvm-another-flash-blockers.md) left that refusal and moved **then-current** static coverage to **256/277, 33**. The [subsequent bounded wait handoff](jvm-wait-handoff-preflight.md) moved selected preflight to `Sound.stop()`/`sun.audio` and static coverage to **258/277, 31** at that checkpoint. The [later passive Class probe](jvm-image-decoder-class-preflight.md) moved static coverage to **259/277, 30** at that checkpoint; the [subsequent bounded branching-finally fixture](jvm-branching-finally-preflight.md) moves current **static-only** coverage to **260/277, 29**; the selected refusal remains. No Flash construction/playback occurred.

**Result at this checkpoint:** the unchanged, SHA-pinned original Java Flash 2 Applet
still **refuses selected-entry preflight before constructing any original
object**. The first reported refusal has moved from `Component.getGraphics()`
in `Flash.run()V` byte 91 to `java/lang/System.currentTimeMillis()J` in
`Flash.run()V` byte **279**:

```text
[applet] com/macromedia/flash/Flash.run()V at byte 279: unsupported Java API java/lang/System.currentTimeMillis()J
```

At **this Graphics checkpoint**, the separate all-method static result was
**232/277 decoded bodies and 57 first-blocker entries**. The later bounded
[bulk-read admission](jvm-bulk-stream-preflight.md) changed that *static*
inventory to **233/277, 56**, while the selected refusal remained at byte
279. Neither a later offset nor changed static totals demonstrate Flash
construction, rendering, playback or remaining independent bugs. There is still no Flash-2-compatible reference movie or
verified original-player frame/input/audio timeline. The earlier
[default-RGB pixel island](jvm-headless-rgb-preflight.md) and
[document/code-base helpers](jvm-base-and-color-preflight.md) have narrower
contracts; neither grants a display peer.

## The exact JDK comparison

Original `Flash.<init>()V` stores `this` in its `Component` field at byte 47.
Its `run()` calls `Component.getGraphics()` at byte 91 and checks for **null**
before `Graphics.drawImage(...)` at byte 110. These are static call-site
observations from the unmodified SHA-pinned archive, not an original method
executed by this gate. Null is an ordinary branch of the original code; it
is **not** a reason to synthesize a writable `Graphics` object.

The **owned** Java 8 fixture `PeerlessGraphicsApplet.java` compiles to pinned
`PeerlessGraphicsApplet.class` (SHA-256
`2a019f6ca3013d642b1143f7d4a52e7e80c911acc4de28f10588526df35a669b`),
`PeerlessGraphicsOverrideApplet.class`
(`f0388bdbfafba7c83d406ae4614faa879ac6b8987989c5dc4c87acb357a2dad7`)
and `PeerlessGraphicsSurfaceApplet.class`
(`daab73b255146c8b84ab4b907f15944535738e8fde5309d2c5e8c558214ceeb6`).
[`PeerlessGraphicsJdk.java`](../test/fixtures/PeerlessGraphicsJdk.java) runs
them on JDK 11 with `-Xverify:all` on a **24-bit Xvfb** display; a separate
constructed `Canvas` is checked with `-Djava.awt.headless=true`. It does not
allocate an Applet unsafely or run the original Flash Applet.

| Independent JDK observation | Bounded Magnolia result |
| --- | --- |
| Detached, constructed Applet and Canvas: `getGraphics() == null`; the plain Applet prints **five `true` lines** including an inherited symbolic-owner call. | Same original owned method bytes print the same five lines on native Oak, standalone Oak and Builder JS/Node; offline HTML matches method output, not display pixels. |
| User subclass overrides `getGraphics()` with a printed marker and null result: **three** virtual markers occur before its boolean lines. | Native, standalone Oak, Builder JS and offline HTML match every line. The linker preflights concrete overrides even when the call-site owner is built-in `Component`. A selected `getColorModel()` override from the earlier RGB island also prints `override-model, true` on JDK/native/Builder JS; it is not silently replaced by the default model. |
| Another **JDK-valid** subclass override constructs its own `BufferedImage` `Graphics`: detached does **not** imply all virtual calls return null. It prints `true, false, false, false, true`. | Refused at selected preflight on unsupported `BufferedImage`, **before the fixture can print**. The override cannot be erased by a built-in null fallback. |
| A constructed Applet attached to an **unpeer'd** Panel still gets null on the JDK; a Frame-hosted displayable Applet returns a **non-null** `Graphics` (three `true` assertions for parent/peer/graphics). | The attached case deliberately **refuses before output**; there is no modeled parent/peer graphics surface. No display-pixel or window parity is claimed. |

Only exact `java/awt/Component.getGraphics()Ljava/awt/Graphics;` via
`invokevirtual` is admitted. On a constructed, **detached** VM Component with
no modeled parent it returns null. `Image.getGraphics()` remains a **different**
existing bounded raster operation; unknown member names/owners/descriptors
and other call modes are refused. The generated-Java target cannot mirror this
host contract and **refuses before source emission**. A same-width unknown
member edit is rejected by native and all four Builder targets without
creating an output artifact.

[`applet-peerless-graphics-oracle.sh`](../applet-peerless-graphics-oracle.sh)
**exited 0** on JDK/Xvfb, native Oak, Builder JS, standalone Oak/JS/HTML
and refusal checks. The original-archive check and this new fixture are also
in [`applet-flash-preflight-regression.sh`](../applet-flash-preflight-regression.sh).
At this checkpoint the combined five-oracle regression **exited 0** with
the pinned original archive and checked the exact selected refusal across
**eight independent processes**. Initially, unordered receiver traversal
could race that refusal against `ScriptPlayer.run()V` byte 11
(`InputStream.read([B)I`); sorted linker traversal yielded **20/20 further
independent byte-279 preflights**. Both messages were pre-construction, and
ordering was not player progress. After the separate bulk-read admission the
combined gate has **six focused oracles** and **also exited 0**, including
eight fresh original preflights; the pin-checked
[cross-media census](media-blocker-census.md) separately records the now
**233/277, 56** static inventory. Reproduce from
the repository root with a JDK 11, Xvfb, Node and built Magnolia executable:

```sh
MAGNOLIA=/path/to/magnolia bash applet-peerless-graphics-oracle.sh
# Only if the separately acquired original Flash archive matches its SHA-256:
MAGNOLIA=/path/to/magnolia FLASH_ZIP=/path/to/FlashApi.ZIP \
  bash applet-flash-preflight-regression.sh
```

**Next gates identified at this historical checkpoint:** original `Flash.run()` stores the time returned at byte 279,
then adds a `ScriptPlayer` frame delay at bytes 284–295 to form a deadline.
An independently checked deterministic wall-clock/virtual-task contract is
needed before admitting `System.currentTimeMillis()`; a hard-coded zero or an
unoracled host clock would not establish that scheduling behavior. The sorted first preflight message did not itself remove the separately
observed `ScriptPlayer.run()` byte-11 bulk-read blocker; the later,
[independently checked one-array admission](jvm-bulk-stream-preflight.md)
removes it only from the static inventory. A *separate* real graphics peer
with JDK/host pixel comparisons is needed for non-null
`Component.getGraphics()`. Silently returning null for an attached Applet or
**instead of invoking** a user override, or returning fabricated `Graphics`
for a detached one, would mask missing Java/host behavior. This JVM step does
not decode a SWF or play a Flash movie.
