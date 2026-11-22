# JDK-checked document base and Flash bgcolor preflight slices

**Historical base/color checkpoint.** The [later five-member differential](jvm-five-blocker-preflight.md) records the present original `jsr` byte-355 preflight refusal and **237/277, 52** static inventory. Earlier offsets here are checkpoint evidence, not today's selected barrier.

**Scope:** these are three *host/library compatibility slices*, not an
executed Java Flash player or a movie. They target the next members actually
referenced by the original, SHA-pinned 1998 `FlashApi.ZIP`; an owned Java 8
fixture and JDK 11 `-Xverify:all` provide the behavioral comparison.

## `Applet.getDocumentBase()` without giving applets network access

The independent `AppletBasesJdk.java`
attaches a real `java.applet.AppletStub` to the owned
[`AppletBaseResource.class`](../test/fixtures/AppletBaseResource.java). Its
**document base** is a local `file:` URL for `document/page.html`; its
**code base** is a separate `file:` URL for `assets/`. Both directories contain
an identically named `movie.bin` with *different bytes*. The original fixture
method calls the real JDK `Applet.getDocumentBase()`/`getCodeBase()`, resolves
`new URL(base, "movie.bin")`, opens its real file stream and reads two bytes.
The independent JDK output is:

```text
document
2
68
49
code
2
67
50
```

Magnolia now admits only the exact instance method descriptor
`java/applet/Applet.getDocumentBase()Ljava/net/URL;`. It returns an **opaque
base token**, distinct from `getCodeBase()`. `new URL(base, basename)` can open
only explicitly supplied in-memory `options.documentResources[basename]` for
the document base, or `options.resources[basename]` for the code base. The
bases **never** implicitly share a resource map. Both maps combined allow at
most **four** single-component ASCII names; each supplied value is at most
**524,288 bytes**, total at most **1,048,576 bytes**. Paths, URL schemes,
`..`, percent escapes, slashes, query strings and unsupplied resources are
refused; the host never opens a socket or local file on bytecode's request.
These limits intentionally model neither arbitrary JDK URL resolution nor an
`AppletContext` or web page. Builder's **generated Java** backend refuses
this URL island before emission; native, standalone Oak and Builder JS are the
checked targets.

[`applet-base-resource-oracle.sh`](../applet-base-resource-oracle.sh)
rebuilds the owned fixture byte-for-byte, compares the JDK's *actual file
reads* with Magnolia's separate in-memory maps across native/Oak/Builder JS,
and verifies missing/cross-base/unsafe-name/budget guards, a JS refused
remote URL, and generated Java's **no-artifact, pre-emission refusal** for the
reachable resource island. This focused oracle **exited 0**; it did not feed a
third-party movie to the Flash Applet.

## Historical bgcolor helper members

The original `Flash.ň(Z)V` also calls `String.replace('#','0')` before
`Integer.valueOf(text, 16)`. Magnolia now admits the **exact** descriptors
`String.replace(CC)Ljava/lang/String;` and
`Integer.valueOf(Ljava/lang/String;I)Ljava/lang/Integer;`, not arbitrary
String/Integer APIs. The bounded replace accepts at most **4,096 ASCII
characters**, preserves the original JVM String identity for unchanged
results, and returns a new object for changed results. The bounded numeric
parser accepts a signed, at-most-**4,096-character ASCII** number in bases
**2–36**, checks int32 overflow *before* multiplication, raises a catchable
`NumberFormatException` for invalid digits/radix, and retains the explicitly
checked `-128..127` boxed-reference cache. Java-valid UTF-16 digits and
larger inputs fail the declared Magnolia policy, rather than being silently
misparsed. Both helpers also run through generated Java source, unlike URL
streams.

[`applet-string-replace-oracle.py`](../applet-string-replace-oracle.py)
and [`applet-int-radix-oracle.py`](../applet-int-radix-oracle.py)
each independently rebuild their Java 8 source/class bytes, run JDK 11
`-Xverify:all`, compare native Oak, standalone Oak, Builder JS/Node and
Builder-generated Java output, and check malformed-member/Unicode/bound
refusals. **Both exited 0.** The numeric test exercises the actual
`#00Ff → 000Ff → 255` method sequence on an **owned fixture**; it does *not*
run the original Flash player's color/image/timeline code.

**At this three-slice checkpoint**, the selected original
`com/macromedia/flash/Flash` entry still failed preflight before construction
or playback, at `Flash.ň(Z)V` **byte 593**:
`java/awt/image/ColorModel.getRGBdefault()Ljava/awt/image/ColorModel;`.
The then-current Java SE *static first-blocker inventory* was **231/277
decoded bodies and 58 first-blocker entries**—not a list of remaining runtime
defects. A later [JDK-oracled headless RGB pixel island](jvm-headless-rgb-preflight.md)
models both color-model branches and real, bounded `int[]` image pixels; the
next selected refusal was `Component.getGraphics()` byte 91. A further
[peerless-Graphics differential](jvm-peerless-graphics-preflight.md) admits
only the detached Component's null result; the **then-current** selected refusal
was instead `System.currentTimeMillis()` in `Flash.run()V` byte 279; the
later [bounded one-array read](jvm-bulk-stream-preflight.md) changes static
coverage to **233/277, 56** without changing that selected refusal. That still does **not** construct or play the
Flash Applet. The Java player needs a Flash-2-compatible movie and whole
lifecycle/graphics/audio/input reference before any playback claim.
The 2020 Adobe ActiveX [CAB analysis](adobe-swflash-cab.md) is a *separate*
Windows runtime track, not a substitute for these Java host contracts.

Reproduce these three document/bgcolor comparisons from the repository root:

```sh
MAGNOLIA=/path/to/magnolia bash applet-base-resource-oracle.sh
python3 applet-string-replace-oracle.py --magnolia /path/to/magnolia
python3 applet-int-radix-oracle.py --magnolia /path/to/magnolia
# Only with the separately acquired and SHA-checked original SDK ZIP:
/path/to/magnolia applet.oak run --entry /path/to/FlashApi.ZIP \
  --class com/macromedia/flash/Flash --max-steps 2000
# The last command intentionally exits 1 (current Sound.stop() byte 36,
# unsupported sun/audio/AudioPlayer.player, before construction).
```

For a single **opt-in focused regression** of those three oracles **plus the
new detached RGB pixel oracle**, and the exact original Flash refusal/static
inventory *when a separately acquired original ZIP is provided*, run:

```sh
MAGNOLIA=/path/to/magnolia FLASH_ZIP=/path/to/FlashApi.ZIP \
  bash applet-flash-preflight-regression.sh
```

Without `FLASH_ZIP`, that script runs the four JDK differentials and reports
a **skip** for the original-archive check, not an original-player pass. It is
not part of `test/main.oak applet`: the JDK 11, Xvfb and Node requirements,
external proprietary input, and multi-minute runtime do not belong in the
ordinary focused Oak stdlib suite. The [signed CAB inspector](adobe-swflash-cab.md)
is a distinct, opt-in **offline artifact** review, not a JVM regression or a
player run.
