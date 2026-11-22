# JDK-checked detached RGB image island — still not Flash playback

**Historical RGB checkpoint.** The [later five-member differential](jvm-five-blocker-preflight.md) recorded an intermediate original `jsr` byte-355 refusal and **237/277, 52** static inventory; [six further bounded gates](jvm-next-blocker-preflight.md) later reported selected `Object.wait()V` byte 4 and **245/277, 44** static; [five additional archive-derived families](jvm-another-flash-blockers.md) left that same selected refusal and moved the **then-current** static inventory to **256/277, 33**. A [subsequent bounded wait handoff](jvm-wait-handoff-preflight.md) moved the selected refusal to `Sound.stop()`'s `sun/audio/AudioPlayer.player` and separate static coverage to **258/277, 31** at that checkpoint. The [later passive Class probe](jvm-image-decoder-class-preflight.md) moved static coverage to **259/277, 30** at that checkpoint; the [subsequent bounded branching-finally fixture](jvm-branching-finally-preflight.md) moves current **static-only** coverage to **260/277, 29** without changing the selected refusal. Earlier offsets/counts below were real, narrower measurements.

**Result at this RGB checkpoint:** the unchanged, SHA-pinned Java Flash 2
`FlashApi.ZIP` still **refused selected-entry preflight before constructing
the Applet**. After a bounded default ARGB color-model and `int[]`
image-production differential, its first refusal was
`com/macromedia/flash/Flash.run()V` byte **91**:

```text
[applet] com/macromedia/flash/Flash.run()V at byte 91: unsupported Java API java/awt/Component.getGraphics()Ljava/awt/Graphics;
```

The *all-method static* inventory at this checkpoint was **232/277 decoded
bodies, 57 first-blocker entries**. One fewer first message is not one fewer
independent runtime defect. A subsequent
[peerless-Graphics JDK differential](jvm-peerless-graphics-preflight.md)
admits only a detached Component's **null** result; that checkpoint had the
same static counts but moved selected preflight to `System.currentTimeMillis()`
byte 279. The later [bounded bulk read](jvm-bulk-stream-preflight.md) changes
static inventory to **233/277, 56**, with the **same selected clock refusal**. **No Flash constructor, movie, window, frame, audio, input or
Flash-5 SWF executed.** The previous `ColorModel.getRGBdefault()` refusal in
`Flash.ň(Z)V` byte 593 and **231/277, 58** were an earlier checkpoint after
[separate document/code-base and bgcolor helpers](jvm-base-and-color-preflight.md).

## What was actually admitted

The original `Flash.ň(Z)V` calls both the static
`java/awt/image/ColorModel.getRGBdefault()Ljava/awt/image/ColorModel;` and the
alternative `java/awt/Component.getColorModel()Ljava/awt/image/ColorModel;`
before passing a model into `DisplayList`. In a **detached, headless** profile,
Magnolia models the JDK's stable default 32-bit ARGB model as an opaque,
run-local singleton. Only its exact `ColorModel.getRGB(I)I` method returns the
unchanged signed 32-bit ARGB pixel. `Component.getColorModel()` returns that
same model **only for a constructed, detached Component with no modeled
parent**; user-defined overrides retain virtual dispatch. No peer or display
configuration is inferred.

A separate exact `MemoryImageSource(IILjava/awt/image/ColorModel;[III)V` path
accepts **only this default model and a JVM `int[]`**. The producer retains the
source array; `Component.createImage(ImageProducer)` makes a metered,
independent ARGB snapshot, with offset/stride bounds checked before copying.
The older 8-bit `IndexColorModel`/`byte[]` Fliplay route remains distinct.
A **later, separate** [JDK-checked five-argument `int[]`/default-model and
`Image.getWidth/getHeight` fixture](jvm-ds-memory-rgb-next.md) now admits only
that additional bounded path; it does not retroactively make this historical
Flash 2 checkpoint an original DS worker/image result.
Existing per-array (default 65,536 cells), cumulative image (4,194,304 cells),
step and reference budgets still apply. Drawing admits opaque and fully
transparent pixels only; a JDK-valid partial-alpha draw **refuses before
output** rather than being approximated. `IndexColorModel.getRGB(int)` and an
`IndexColorModel` paired with the new `int[]` constructor also refuse under
this declared default-model-only policy, despite being valid on a full JDK.
Generated Java **refuses the reachable ColorModel API before emitting source**;
it has no equivalent image host.

This is not general `DirectColorModel`, `ColorModel` subclassing, Java2D
compositing, asynchronous `ImageProducer`/`ImageConsumer`, arbitrary
`Component.getGraphics()`, display/peer color depth, a window or a SWF player.

## Independent JDK oracle and route results

[`DefaultRGBModelApplet.java`](../test/fixtures/DefaultRGBModelApplet.java) is
an **owned** Java 8 fixture; its checked-in class SHA-256 is
`9afa29946b57cad4254d674fcc639ae47be7bdc34bbf651df806f8c87c67173f`.
[`DefaultRGBModelJdk.java`](../test/fixtures/DefaultRGBModelJdk.java) runs it
under JDK 11 `-Xverify:all -Djava.awt.headless=true`. **Only the reference
harness** bypasses the display-bound Applet constructor; it consumes the real
JDK `MemoryImageSource` through an `ImageConsumer`, calls the actual
`ColorModel.getRGB(int)`, and captures `BufferedImage` pixels. Magnolia runs
the fixture's original compiled bytecode with its bounded native image host.
The producer's 2×2 source (nonzero offset and stride) changes *before* image
creation and is then changed again *after* the image snapshot:

```text
source: ff445566 00112233  |  ff00ff80 ff801020
canvas row 0: ff0000ff ff445566 ff0000ff ff0000ff
canvas row 1: ff0000ff ff00ff80 ff801020 ff0000ff
canvas row 2: ff0000ff ff0000ff ff0000ff ff0000ff
```

The fully transparent source pixel leaves the blue canvas unchanged. **All
source and canvas pixels and seven printed fixture lines match** on JDK,
native Oak and the Builder JS bundle under Node. Standalone Oak and JS/Node
Applet-transpile routes, plus offline HTML in a simulated DOM, match the seven
JDK method-output lines; HTML is **not** a browser display-pixel comparison.
Three additional methods **execute on the JDK** but refuse before output in
Magnolia: partial alpha, indexed model with `int[]`, and
`IndexColorModel.getRGB(int)`. A same-width renamed `getRGBdefault` fails
selected preflight on native and all four Builder targets, without artifacts.

The display-dependent exception is independently demonstrated by
[`DefaultRGBDisplayJdk.java`](../test/fixtures/DefaultRGBDisplayJdk.java):
`Canvas.getColorModel()` **is the default singleton** (32 bits, alpha) under
JDK headless mode; under an explicitly configured **24-bit Xvfb display** it
is a **different** 24-bit model without alpha. Magnolia does **not** assert
Xvfb/display parity. An implementation that returned the default model for
every real GUI component would be wrong.

The focused [`applet-default-rgb-oracle.sh`](../applet-default-rgb-oracle.sh)
and the previous [`applet-raster-oracle.sh`](../applet-raster-oracle.sh)
**both exited 0**, including the existing indexed image regression. The
then-combined four-oracle [`applet-flash-preflight-regression.sh`](../applet-flash-preflight-regression.sh)
**exited 0 with the original SHA-pinned `FlashApi.ZIP`** at the byte-91
checkpoint; the later combined regression adds a fifth, separate
[peerless Graphics oracle](jvm-peerless-graphics-preflight.md), **exits 0**
with pinned Flash and checks its byte-279 refusal. The separate SHA-pinned
[cross-media census](media-blocker-census.md) also exited 0. Reproduce from the repository root with JDK 11, Xvfb, Node and a built
Magnolia interpreter:

```sh
MAGNOLIA=/path/to/magnolia bash applet-default-rgb-oracle.sh
# Only with a separately acquired and SHA-checked original Java Flash ZIP:
MAGNOLIA=/path/to/magnolia FLASH_ZIP=/path/to/FlashApi.ZIP \
  bash applet-flash-preflight-regression.sh
```

The 2020 Adobe [ActiveX CAB review](adobe-swflash-cab.md) is a **separate,
offline Windows installer lead**; no CAB member or installer is loaded by this
JVM test. The once-next `Component.getGraphics()` call is window/peer-dependent:
[returning JDK's detached null](jvm-peerless-graphics-preflight.md) is valid
only without a real peer; fabricating a `Graphics` would move a barrier
without proving Java Flash rendering. The *then-current* `System.currentTimeMillis()`
refusal needed a separately oracled clock/scheduler contract; only an [opt-in logical subset](jvm-five-blocker-preflight.md) was later added. A genuine
Flash-2-compatible movie and independently compared whole-player frames,
input and audio remain open acceptance gates.
