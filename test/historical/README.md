# Executed historical JVM `.class` slices (not full players)

These four small **new harness sources** invoke methods on **original, hash-checked JVM 45.3 binaries**, not reimplementations of their algorithms. The Macromedia and Fliplay class files are **not checked in**: redistribution rights are unresolved. A JDK compiles the harnesses and supplies an independent `java -Xverify:all` oracle. Magnolia itself links and executes those original bytes, and Builder emits self-contained Oak/JS plus optional Java source from the selected verified image. The generated runnable outputs exist only in a temporary directory during this test; do not distribute outputs embedding third-party classes without permissions.

From the Magnolia repository root, using separately and lawfully obtained artifacts:

```sh
FLIPLAY_CLASS=/path/to/fliplay.class \
FLICKFRAME_CLASS=/path/to/flickframe.class \
FLASH_ZIP=/path/to/FlashApi.ZIP \
MAGNOLIA=/path/to/magnolia \
    bash tools/historical-core-matrix.sh
```

To **explicitly** fetch matching copies from their recorded public locations instead:

```sh
MAGNOLIA=/path/to/magnolia bash tools/historical-core-matrix.sh --download
```

Requires Bash, Python 3, Node.js, `sha256sum`, `javac`/`java` (JDK 11 tested), and the Magnolia interpreter. The script verifies SHA-256 **before** compiling/running, writes no proprietary bytes to the repository, and removes all temporary artifacts. Its independent JVM oracle runs Fliplay with `-Djava.awt.headless=false` to construct the original Applet object; this is **not** a working graphical host. The Oak and Node executables use Magnolia's verified VM subset, without a runtime JVM or browser Java plugin.

| Original methods actually invoked | JDK 11 output | Evidence/limit |
| --- | --- | --- |
| `fliplay.<init>` + `get_ubyte`, `get_ushort`, `get_ulong` on signed byte arrays | `255`, `128`, `33023`, `13440`, `305430783`, `-1` | Unmodified [`fliplay.class` and `flickframe.class`](https://www.oortman3d.com/fliplay/) hashes: `bc315872dafebbf3c634e64c0fc251b38031ed49b7c4698d39c94990de9fe20b`, `12feb5321cd8c9aeb0bf285f055596d981e3d45ccfb8c62d029f2f6d95ccdf9f`. No file/image/thread playback. |
| `com/macromedia/flash/Matrix.<clinit>` + obfuscated `ŷ(I)I` and `Ÿ(II)I` | `23`, `-2147483648`, `2147483647`, `11`, `63`, `0` | Original [`FlashApi.ZIP`](https://raw.githubusercontent.com/macromedia-os/Macromedia-Flash-SDK/main/JavaFlash/FlashApi.ZIP) SHA-256 `fbb2149179c1c5a32bd9aa7f479a1e760e36e18dbd8e48147729ca67822932c2`; `Matrix.class` SHA-256 `df934e7dce8756e79abe3360d1cea5a29818f4ed25f813ad0d1739abe8fca6c3`. Integer geometry and a real 66-element lookup-table initializer, **not** graphical playback. |
| `com/macromedia/flash/SParser.<init>` + obfuscated `Á()I`, `Â()I`, `Ã()I` and input/cursor fields | `33023`, `-1`, `-2147478988`, `8` | Original `SParser.class` SHA-256 `9fc21c5ec64e9daa32f163588f95d5ad18aae6cf1f73288c2932e11d882338db`. Real little-endian unsigned/signed reads through the Macromedia parser, **not** loading or rendering an SWF. |
| `com/macromedia/flash/Rect.<init>` + original `ſ(Point)` union, `ƀ(Rect)` intersection and `Ɓ(Point)` containment | `true`, `false`, `true`, `false`, `true`, `true`, `false`, `true`, `false`, `true`, `false` | Original `Rect.class` SHA-256 `36e7b476cec00972cc31929a7dfb32d7ba5abb729e4356288dec8e7487110f74`; `java.awt.Point` constructor and mutable `x/y` fields are modeled by Oak, not stubbed. `Rect` has **7/7 decoded bodies**, but this row exercises only its selected geometry path, not the full player. |

Every row is compared against *native Magnolia execution*, *runnable standalone Oak*, *Builder JS under Node*, and *the generated self-contained Java interpreter source*, compiled and run independently. The original `SParser` **and `Rect`** also run in Builder's **self-contained HTML**: a network-free simulated DOM executes two inline scripts per page and matches the independent JDK oracle. This tests the selected browser delivery path, **not** AWT events or graphics. Exact-byte **original class restoration** is checked separately. One edit to the authentic `SParser.Á()` method (`sipush 255` → `sipush 127`) assembles a **different JDK-verified class**; edited Oak and JDK both print `32895`, `-129`, `-2147478988`, `8`. A separate same-width edit to the original `Rect.Ɓ(Point)` changes its first `if_icmpgt` to `if_icmpge`, making the third and fifth boolean lines false instead of true; the new JDK-verified class, edited Oak, Builder JS and generated Java match. Restoring either edited Oak yields **original**, unmodified class bytes; restoration is not reassembly.

The Macromedia obfuscated JVM identifiers are modified-UTF-8 in the class file. `lib/applet-java.oak` decodes those bytes to distinct, ASCII-only Java `\uXXXX` escapes instead of collapsing them or turning each input byte into a character; malformed input is rejected. This improves **the generated Java source**; it does *not* add general Java UTF-16 semantics to Magnolia's Oak/JS VM. Unmodeled Unicode `String` operations are still rejected.

**Not demonstrated:** Fliplay `init`/frame animation, the Macromedia `Flash` Applet constructor/lifecycle, Flash 2 SWF playback, complete JVM verification, complete AWT/image/audio/network/threads, or any Rapper Rhythm executable. Whole-input coverage is still **14/35 Fliplay methods with 33 blockers** and **199/277 Flash methods with 122 blockers**. The Flash selected path still stops at `Flash.<init>()V` byte 98 (`dconst_1`, followed by a double field); Fliplay `init` still stops at `Component.createImage(II)`. A newer lineage additionally reports 218/277 Flash bodies and 101 first blockers after exact-limb long arithmetic, and then **219/277 with 100 first blockers after bounded `multianewarray`** (the one newly decoded body being the original `Bitmap.<clinit>`); the additionally decoded methods and the multianewarray-era rerun were **not executed and not performed** in this environment — those figures are the patch author's, unverified here. A bounded F/D slice ported into this tree on 2026-11-16 is author-reported to decode that constructor, and a fresh selected-entry `run --class com/macromedia/flash/Flash` is reported (author-reported, not reproduced in this environment — no JDK) to stop before execution at `Flash.start()V` byte 6 `monitorenter` (0xc2); the author-reported post-exception inventory is **221/277 Flash methods with 91 first blockers** (96 at the preceding checkpoint) and **18/35 Fliplay decoded bodies with 26 first blockers** (17/35 with 30 at the preceding checkpoint), not whole execution. See [the compatibility plan](../../docs/applet.md) and [Flash artifact note](../../docs/flash-artifacts.md).

## Original Moon Clouds all-frame load/decode (new bounded route)

[`tools/fliplay-all-frames-oracle.py`](../../tools/fliplay-all-frames-oracle.py) verifies the uploaded ZIP/FLC and both archived `.class` SHA-256 values, validates every FLC frame/chunk and the 16-byte ring header, and temporarily compiles the small [`FliplayFrameLink.java`](FliplayFrameLink.java) reachability fixture against **external** original classes. Its checked-in Java 8 `.class` is independently rebuilt byte-for-byte. The JDK 11 [`FliplayAllFramesJdk.java`](FliplayAllFramesJdk.java) harness opens a single continuous FLC stream and executes the genuine `load_header`, `read_loop`, `load_frame`, `extract_frame`, palette/BRUN/DELTA_FLC code on all **60 display frames plus ring** under `-Xverify:all`. A distinct [Oak runner](../../tools/fliplay-all-frames.oak) calls those original methods with separately bounded in-memory header/frame streams, capped at **10,000,000 VM steps**, **81,920 cells per array** and **1,048,576 total array cells**. It uses zero-initialized `flickframe` objects: the original `load_frame` sets each `chunks` field; the unused AWT patch-array constructor is not compiled.

**Result:** all **60/60 indexed images and palettes** match JDK original-method output byte-for-byte; every RGB frame additionally matches Pillow's independent FLC decoder; the ring frame is unchanged on both. **6,831,437** JVM steps and **469,697** charged cells (including external input arrays). The 60-image concatenated indexed SHA-256 is `dd06653c96de9da8cf3b066c72a9885ef3c19d5a8c5de662b42cb4b6a4c26102`; RGB is `dfb9f2fbde80dafbee976ae43ec2c04c1019d7ffccb13807400324465f89fe16`. The Python oracle optionally writes per-frame hashes and a self-contained [offline viewer built from Oak pixels](../../tools/fliplay-original-viewer-template.html); [`fliplay-viewer-oracle.js`](../../tools/fliplay-viewer-oracle.js) exercises that viewer's playback controls without external modules. **Original `process_lc_chunk` is linked but not triggered** (no type-12 chunks in this file). These tests do not execute the full original Fliplay Applet: its URL opener, AWT image host, patch rendering, controls and animation thread remain missing. The new HTML displays precomputed original-decoder frames with newly written playback code; it does not execute Fliplay's JVM bytecode in the browser.

```sh
python3 tools/fliplay-all-frames-oracle.py \\
  --movie /path/to/moonclouds.zip.txt \\
  --fliplay /path/to/fliplay.class --flickframe /path/to/flickframe.class \\
  --magnolia /path/to/magnolia \\
  --html /tmp/moonclouds-original-fliplay-frames.html \\
  --report /tmp/moonclouds-original-fliplay-report.json
node tools/fliplay-viewer-oracle.js /tmp/moonclouds-original-fliplay-frames.html
```

The historical class files and movie are **not in Git**. This original all-frame Fliplay result does not imply Flash Java player, SWF5/AVM1, Rapper Rhythm provenance, or native projector support.
