# Archived applet and SWF corpus: bounded compatibility ablations

**Scope.** Eight authentic JVM inputs (one historical Java Flash player SDK, the Fliplay pair, four Falstad Java applet JARs, a hosted Jmol JAR, and a small compiled applet) plus two *different-format* ActionScript 3 SWFs. They are independently hash-pinned and kept **outside the repository**. The results distinguish class parsing, decoder inventory, reference-JDK loading, Magnolia's reachable-entry preflight, actual selected-method execution, and target emission. None establishes a complete JVM, graphical applet host, Flash player, or SWF/ActionScript runtime. The two SWFs were only parsed by a bounded **forensic SWF metadata reader**, not by Magnolia.

## Provenance and immutable identities

Download URLs are in [`applet-archived-ablation.py`](../applet-archived-ablation.py); the rows below give the original hosting/preservation page and full SHA-256 for the downloaded binary. The `fetch` mode refuses a changed digest or ZIP CRC and caps individual downloads and extracted ZIP sizes. A ZIP and its inner JAR have **different** identities; both are recorded.

| Original source | Downloaded binary (bytes); SHA-256 | Extracted JVM binary; SHA-256 |
| --- | --- | --- |
| [Macromedia Java Flash Player SDK preservation](https://github.com/macromedia-os/Macromedia-Flash-SDK/tree/main/JavaFlash), [archive](https://raw.githubusercontent.com/macromedia-os/Macromedia-Flash-SDK/main/JavaFlash/FlashApi.ZIP) | `FlashApi.ZIP` (85,800); `fbb2149179c1c5a32bd9aa7f479a1e760e36e18dbd8e48147729ca67822932c2` | 20 original class files; selected original `Rect.class`: `36e7b476cec00972cc31929a7dfb32d7ba5abb729e4356288dec8e7487110f74` |
| [Oortman Fliplay](https://www.oortman3d.com/fliplay/) | [`fliplay.class`](https://www.oortman3d.com/fliplay/fliplay.class) (12,732); `bc315872dafebbf3c634e64c0fc251b38031ed49b7c4698d39c94990de9fe20b` | [`flickframe.class`](https://www.oortman3d.com/fliplay/flickframe.class) (830); `12feb5321cd8c9aeb0bf285f055596d981e3d45ccfb8c62d029f2f6d95ccdf9f` |
| [Falstad Dot Product original Java applet](https://www.falstad.com/dotproduct.old/) | [`dotproduct.zip`](https://www.falstad.com/dotproduct.old/dotproduct.zip) (9,334); `928e9f5126239b4e4747655178cf1b440ab1bee48ec00aba94f5af9319ded3e9` | `dotproduct.jar` (5,793); `758c9722afe2e81c89942dc3337988953e9c78ec67cedf055b8e48635ff115ac` |
| [Falstad Ripple original Java applet](https://www.falstad.com/ripple-java/) | [`ripple.zip`](https://www.falstad.com/ripple-java/ripple.zip) (131,056); `4ab0132e44e2af744ee4ee6ec121bd4698c6233810114a21f33adf009f15b8a5` | `ripple.jar` (106,701); `86c3127a031d91859d822067f1408fee752e172e50c09c3c0a70d7cbad3e811b` |
| [Falstad Fourier original Java applet](https://www.falstad.com/fourier-java/) | [`fourier.zip`](https://www.falstad.com/fourier-java/fourier.zip) (52,994); `d7bbdee85ed69bd2fb6673949e5a071d830a6fef95e2d79cc876d0d86a18ccc4` | `fourier.jar` (24,063); `712d1d09f14aace449e8049cb49517bce19749ffa0401e06d0e33ba3ae8135d7` |
| [Falstad Loaded String original Java applet](https://www.falstad.com/loadedstring-java/) | [`loadedstring.zip`](https://www.falstad.com/loadedstring-java/loadedstring.zip) (37,202); `46a942fbb9627168dd1722d37f0070d8817418e2923f34ab53224ea94220d812` | `loadedstring.jar` (23,507); `795610792899a9e5d9956be8ec4384c635c47f3cd5a53a0f31050448c5a47d80` |
| [Hosted Jmol applet archive](https://www.chem.uwec.edu/Chem352_Resources/jmol/jsmol/java/) | [`JmolApplet0.jar`](https://www.chem.uwec.edu/Chem352_Resources/jmol/jsmol/java/JmolApplet0.jar) (1,123,429); `58ab0d0917091a383fb136cc73354ff7fc7187911c8d45c92c2e5bc7f30a7d83` | 362 class files; **only `JmolApplet` selected** for Magnolia, because the whole JAR exceeds its 256-class cap |
| [Cyberster Java Applet Demo](https://github.com/Cyberster/Java-Applet-Demo) | [`MyApplet.class`](https://raw.githubusercontent.com/Cyberster/Java-Applet-Demo/master/MyApplet.class) (1,003); `03ed1b621b7cc4c0def43eeb20f27bd6c7252c75f745f16614c6deddb79ffc7d` | Standalone original Java 8 class (JVM 52.0) |

Additional format/provenance controls (not JVM inputs): [author-shared AS3 Snake project](https://github.com/AthulDilip/Snake-Game-ActionScript3) [`SnakeTutorial.swf`](https://raw.githubusercontent.com/AthulDilip/Snake-Game-ActionScript3/master/SnakeTutorial.swf) (2,461 bytes; SHA-256 `f07cc3422d04bfa49713cb71a2db8a55c92a6c258f85fc94c782d6ff76c7f10c`) and its `Main.as` source (4,339 bytes; `8e3feddb0833fa5f640798a92bda49ff42126440b409cd675fd7ef1588a99a81`); [preserved 2012 game directory](https://github.com/noahsug/game-archive/tree/master/virus) [`Virus.swf`](https://raw.githubusercontent.com/noahsug/game-archive/master/virus/Virus.swf) (171,197 bytes; `0332ebd3b8f7110af6298633d2d87504f4bba8156b1b61b449498f6e31a36120`). A [commit-pinned No Hats Java **source**](https://raw.githubusercontent.com/chvid/sjg/81164a19f3e81d394541740e6d9c59cbb75eb336/games/nohats/src/noHats.java) was also preserved (`noHats.java`, 24,682 bytes; `cdbefa19b78714e1514e8dfae30bf50cdbf5faaf070e7aad724784b420d8d0c9`), **not** counted as a binary, verified class, or executed applet. J2ME game JARs were not substituted for browser applets.

**Rights:** the historical Macromedia Java source carries a proprietary/confidential notice even though the preserving repository has a modern MIT license. Neither these third-party binaries nor converted outputs embedding them are included in the Magnolia patch. Check rights before redistribution.

## Parse, JDK load, Magnolia preflight, and actual execution

`inspect` successfully parsed every class in the listed original JVM inputs **except** Jmol, where it parsed only the explicitly selected `JmolApplet`. `coverage` inventories all input fields/methods and *first* decoder blockers, including unreachable code. `decoded bodies` is **not** a verifier-pass or an execution count. The reference JDK 11 independently loaded **each selected entry class** under `java -Xverify:all`, via `Class.forName(name, false, loader)` without initialization; that does **not** verify all classes in a JAR or run its lifecycle. The local probe uses a separate package so it does not collide with signed default-package JAR classes, and independently asserts the JDK's `Applet`→`Panel`→`Container`→`Component` ancestry before each load.

| Original JVM input / selected entry | Magnolia classes parsed | Decoded bodies / methods; first blocker entries | JDK 11 `-Xverify:all` selected entry load | Magnolia **reachable entry** preflight result |
| --- | ---: | ---: | --- | --- |
| Macromedia `Flash` | 20 | 221/277; 96 | Loaded, no init | `Flash.start()V` byte 6: `monitorenter` (`0xc2`) unsupported |
| Fliplay `fliplay` + `flickframe` | 2 | 17/35; 30 | Loaded, no init | `fliplay.init()V` byte 534: `Component.createImage(II)` unsupported |
| Falstad `DotProduct` | 3 | 15/36; 53 | Loaded, no init | `DotProduct.init()V` byte 8: `Container.setLayout(LayoutManager)` unsupported |
| Falstad `Ripple` | 81 | 312/432; 195 | Loaded, no init | `Ripple.addComponentListener(ComponentListener)` is unmodeled |
| Falstad `Fourier` | 7 | 41/87; 115 | Loaded, no init | `Fourier.addComponentListener(ComponentListener)` is unmodeled |
| Falstad `StringWave` / Loaded String | 6 | 34/80; 105 | Loaded, no init | `StringWave.addComponentListener(ComponentListener)` is unmodeled |
| Jmol `JmolApplet` | **1 selected / 362 in JAR** | 2/26; 28 (selected class only) | Loaded from full JAR, no init | No supported entry lifecycle/main with wrapper excluded; selected class's superclass/interface appear unmodeled because they were **not linked**, not because the JAR lacks them |
| Cyberster `MyApplet` | 1 | 5/6; 1 (unmodeled `Font` in `paint`) | Loaded, no init | Inherited `setBackground(Color)` host method is unmodeled |

**Historical snapshot, not current preflight:** The table above records the original eight-input ablation before later host changes. On the later tree, original `MyApplet` executes constructor/`init`/`start` with JDK-matched color state **and**, after a separately JDK-oracled bounded Font island, its original `paint` produces a JDK-matched **logical draw-call trace**. Generated Java on the controlled JDK/Xvfb host produces a byte-identical PNG; Oak/Canvas **font pixels are not certified** (see [later AWT/paint gates](later-applet-corpus.md)). The bounded original Fliplay/Moon Clouds headless and GUI gates are separate later experiments, and the Java Flash whole player still refuses: the [current selected barrier](media-blocker-census.md) is `Sound.stop()V` byte 36 (`sun/audio/AudioPlayer.player`) after [detached/headless RGB](jvm-headless-rgb-preflight.md), [peerless Graphics-null](jvm-peerless-graphics-preflight.md), [five earlier member-family](jvm-five-blocker-preflight.md) [six further bounded](jvm-next-blocker-preflight.md), [five archive-derived](jvm-another-flash-blockers.md) and [single-notifier wait](jvm-wait-handoff-preflight.md) oracles, not the earlier wait, `jsr`, `Component.getGraphics`, `ColorModel.getRGBdefault` or `Applet.getDocumentBase` barriers. Do not apply the older eight full-entry refusals or decoded-body counts to the later tree. The **new three-input cohort** has its own [non-executing report and AWT research](later-applet-corpus.md), not a ninth/eleventh row in this baseline.

**Route-by-route:** each of those eight *full applet entry* attempts produced `Unsupported:` in `inspect --check`; `inspect` itself still exits 0, so diagnostics were explicitly checked. Native `run` (2,000-step bound) rejected **before any application output**. Four *independent* Builder attempts per entry—standalone `.oak`, bundled `.js`, offline `.html`, and generated `.java`—each rejected at the **same reachable-entry preflight** with nonzero status and **without writing a file**. Thus there is **no standalone Oak execution, no Node/JS execution, no HTML DOM execution, and no compiled generated-Java execution for these full entries**. Their JDK load success is not proof of Magnolia compatibility; Magnolia's rejection is generally *unsupported subset/host behavior*, **not** a Java `VerifyError`.

There are successful, strictly narrower *original-method* routes. The hash-checked Macromedia `Rect.class` plus a **new local harness** executed its original `Point.x/y`-based constructor/union/intersection/containment methods and produced **11 identical lines** on JDK `-Xverify:all`, Magnolia native Oak, standalone Oak, bundled JS/Node, offline HTML's two inline scripts in a simulated DOM, and compiled generated Java. The Macromedia `Matrix` and `SParser` core harnesses and Fliplay's unsigned-byte readers also matched JDK/native/standalone Oak/JS/generated Java; only `SParser` and `Rect` were tested as inline offline HTML. Original-byte restoration and separate same-width edits to real `SParser` and `Rect` classes were checked against newly JDK-verified classes. These are **selected method paths**, not a Flash 2 player, SWF renderer or Fliplay animation. See [`historical-core-matrix.sh`](../historical-core-matrix.sh) and its [scope/readme](../test/historical/README.md).

A **local** `RefSubtypeMain` fixture (not an archive) produced **13 JDK-matched lines** on native class/JAR, standalone Oak, JS, offline HTML and compiled generated Java; the original JAR restored byte-for-byte. Five one-byte mutations were rejected by JDK `VerifyError` and Magnolia preflight before output; two JDK-*valid* mutations instead threw dynamic `ArrayStoreException` and `IncompatibleClassChangeError` after the matching prefix of output. The [reference oracle](../applet-reference-verifier-oracle.sh) passed across these routes; it does not certify full JVMS reference verification.

## Isolated feature ablations, not production edits

Variants copy `lib/` and the CLI to a temporary cache and change **one** admission/check in that copy; the checkout is not ablated. First-blocker entries include class/interface/field declarations as well as method bodies, so their counts need not equal the number of undecoded methods. A method already blocked by something else may gain a new first blocker without changing the decoded-body count.

| Original input / isolated removal | Decoded bodies (baseline → variant) | First-blocker entries | Specific observation |
| --- | ---: | ---: | --- |
| Flash / array-allocation opcodes | 221 → **197** | 96 → **120**; 28 new array first messages | `Matrix.<clinit>` now stops on array allocation; no claim the remaining 197 execute |
| Flash / selected `Math` host calls | 221 → **209** | 96 → **108**; 13 new `Math` first messages | E.g. `Math.abs(I)I` blocked in original `Bitmap` method; unrelated `Math.atan2(DD)D` already unsupported |
| Flash / JDK-verified Applet→Panel→Container→Component ancestry | 221 → **220** | 96 → **97**; one new subtype blocker | Original `Flash.<init>` falsely fails at byte 47: expected `java/awt/Component`, found `Flash` |
| Flash / `invokeinterface` admission | 221 → **221** | 96 → **96**; no new first message | Negative control for this inventory only; does **not** imply interface dispatch is unnecessary |
| Ripple / array-allocation opcodes | 312 → **310** | 195 → **197**; 3 new array first messages | `reinit` and `getTimeMillis` newly blocked; another method was already blocked elsewhere |
| Ripple / selected `Math` host calls | 312 → **312** | 195 → **195**; no new first message | No demonstrated first-blocker effect in this input |

**Behavioral ablations:** removing just the AWT `Point.x/y` field admission from an isolated engine copy makes the original `Rect` harness fail *before any of its 11 JDK lines*, and Builder refuses to emit a JS bundle. A separate one-byte `RefSubtypeMain.class` mutation passes a sibling to a `RefParent` call *after* the first print site: JDK `-Xverify:all` and the normal Magnolia gate reject before output, whereas an isolated weaker reference gate prints `4` and then fails on `RefOther.score()I`. This demonstrates why preflight matters; the deliberately weakened variant is **not** shipped.

## Different format: two AS3 SWFs

A bounded CWS inflate (declared-size cap 8 MiB), RECT/frame header and tag walk found valid End tags and a DoABC header in both files; this is **container/header inspection only**, not an ActionScript ABC parser/verifier/player. Both set the AS3 FileAttributes bit, each contains **one** DoABC tag with an ABC header `minor 16 / major 46`, and neither has a DoAction tag.

| Original SWF | Container bytes → declared expanded bytes | Version; fps/frames; tags | DoABC payload after tag name | Magnolia parse / verify / execute |
| --- | ---: | --- | ---: | --- |
| `SnakeTutorial.swf` | 2,461 → 4,732 | CWS v15; 24 fps / 1 frame; 22 tags | 2,628 bytes | **No AS3/SWF route**; not attempted as a JVM class |
| `Virus.swf` | 171,197 → 236,578 | CWS v9; 30 fps / 2 frames; 303 tags | 15,916 bytes | **No AS3/SWF route**; not attempted as a JVM class |

A SWF game is **not** the archived Java Flash player implementation, a PE projector, or a JVM class. **Later separate user uploads are not part of this earlier corpus:** `moonclouds.zip.txt` is a 60-frame Autodesk FLC (for Fliplay), `rapper_rhythm.zip.txt` contains an OLE `.fla` authoring project but no exported SWF/EXE, and `asdf.swf.txt` is Flash v5 (not Flash-2-player compatible). Their bounded, hash-checked inspection and the original Fliplay *first-frame* execution at this earlier checkpoint are documented in [the Flash artifact status](flash-artifacts.md) and [the JVM plan](jvm-compatibility-plan.md). A [later bounded full-Moon-Clouds Fliplay gate](current-media-status.md) does run, but the Java Flash Applet still does not. [the ActionScript TODO roadmap](ActionScript TODO roadmap) is a roadmap, not working support. The separately supplied **unchanged** AS3JS ZIP remains 10/16 on Node. A **later**, separate MIT [patched reference fork](as3js-baseline.md) passes 19/19 and produces two stable JS compiler generations, neither of which is a Magnolia `.as`/AVM/SWF engine.

## Reproduce with bounded phases

From the Magnolia repository root (Python 3, `curl`, Node, JDK `javac`/`java`, and an already built Magnolia interpreter):

```sh
export MAGNOLIA=/path/to/magnolia
cache=/tmp/magnolia-archived-corpus
# Fetch only matching digest/capped bytes. No third-party binary is committed.
python3 applet-archived-ablation.py fetch --cache "$cache"
python3 applet-archived-ablation.py inventory --cache "$cache"
python3 applet-archived-ablation.py jdk-load --cache "$cache"
python3 applet-archived-ablation.py routes --cache "$cache"
python3 applet-archived-ablation.py flash --cache "$cache"
python3 applet-archived-ablation.py ripple --cache "$cache"
python3 applet-archived-ablation.py rect --cache "$cache"
python3 applet-archived-ablation.py verifier --cache "$cache"
python3 applet-archived-ablation.py swf --cache "$cache"
# Individual operations have child-process/network bounds; cap the whole run too.
MAGNOLIA="$MAGNOLIA" timeout --signal=TERM --kill-after=3s 240s \
    bash applet-reference-verifier-oracle.sh
```

Each phase is independently invocable and writes diagnostics/JSON under `$cache/logs` and `$cache/reports`. A whole-inventory `coverage` command exits **1 on any blocker**, which is expected here; a `routes` invocation exits 0 only when the **expected** rejection occurs on every checked target. The local reference/selected-method matrices additionally passed `bash applet-verifier-oracles.sh` (**27/27** malformed-class comparisons) and the historical matrix. The focused subset passed **346/346 native** and **346/346 bundled JS** (build with explicit `--include 'applet.test:test/applet.test.oak'` alias), including a new regression for the JDK-verified Applet ancestry. These results do not supersede unrelated repository-wide Go/pack test failures, do not promise a sandbox for hostile classes, and do not make either complete Java Flash player or either AS3 game runnable.
