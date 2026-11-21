# Acotel Xmas Gifts: separate Java ME intake, not a Java SE applet

The supplied [`ACO_XmasGifts_s.jad.txt`](../../uploads/ACO_XmasGifts_s.jad.txt) is a **511-byte, CRLF Java ME application descriptor (JAD)**, SHA-256 `f60ea572dcc249290a5c62dc5391bccdff609f513b4c0d7599e838e4c056b055`. It declares **CLDC-1.0 / MIDP-1.0**, Acotel's `ACO_XmasGifts_s` MIDlet, and a **50,594-byte companion JAR**. It contains no class bytecode and is **not** an Applet, an executable JAR, a game movie or a Java SE browser-plugin descriptor. The uploaded JAD's obsolete IP-address `MIDlet-Jar-URL` was **never followed**. Oracle's [MIDP application model](https://www.oracle.com/technetwork/java/ch08-159481.pdf) distinguishes MIDlets from the Java SE Applet model.

An online application to try with both files is [J2ME Player](https://kinoflix.github.io/plus/j2me/): its Add new game UI offers **Select JAR file** and **Select JAD file**. The uploaded file here is a JAD only; this upload does **not** by itself provide the game's JAR. The separate archived JAR documented below was fetched and pinned during the independent intake, not uploaded by the user. We **did not upload the game to that player, install it, or establish that it runs there**. Availability of its JAD picker does not prove that supplying JAD is required, optional, or sufficient for this particular game; check the player's current behavior separately. This Magnolia ablation is not a test of that online player.

## Independent bytes and scope

The [Internet Archive's `1000JavaGames.zip` listing](https://ia600600.us.archive.org/view_archive.php?archive=%2F18%2Fitems%2F1000_J2ME_Games_Pack%2F1000JavaGames.zip) has paired `All Java Mobile Phone Games/A/ACO Xmas Gifts/ACO_XmasGifts_s.jad` and `.jar` members. `midlet-acotel-ablation.py` opt-in downloads **those named members only**, under 1,024-byte/65,536-byte transfer caps, then enforces pinned lengths and SHA-256 before inspection. The archived JAD is **byte-identical** to the user-supplied JAD. The JAR is **50,594 bytes**, SHA-256 `c24e61a416616a036b537957ec9b0687ad08f9fceae18e4d06336883ec8c3040`. This proves agreement with this archive copy and the descriptor's size/attributes; it is not an independent software-signature or redistribution-rights verification.

The original JAR passes bounded ZIP member/path/compression/size/CRC checks: **38 entries, 88,379 inflated bytes**, including **two classes** and **35 non-class files** (one is the manifest; one directory entry is excluded). All six shared JAD/manifest attributes match, including `MIDlet-1`'s entry class. The classes are Java classfile **45.3**, not modern Java SE applets:

| Original JAR class | Uncompressed bytes | SHA-256 | Superclass |
| --- | ---: | --- | --- |
| `ACO_XmasGifts_s.class` | 1,870 | `67fa7392200f9ef5af900501079a5179249c6993581b86426ec820cc7f2de38c` | `javax/microedition/midlet/MIDlet` |
| `DS_Effects.class` | 38,935 | `b66e5fc085e6e8de4af61ddb680db04e7668723e01b883fb096bcbd406bf5903` | `javax/microedition/lcdui/Canvas` |

The manifest names `ACO_XmasGifts_s.png` as its icon; **that path is not among this JAR's members**. This is an inventory observation, not a claim that a particular emulator succeeds or fails. The JAR is **not committed to the source tree or the delivery archive**. The independent user upload is also outside the project archive.

## Historical static ablation and ordinary selected-entry refusals — not game playback

The applet CLI now explicitly rejects `.jad` / `.jad.txt` as metadata. The applet compiler recognizes direct or supplied transitive `javax/microedition/midlet/MIDlet` ancestry and refuses it as a **Java ME entry**, instead of misreporting it as a generic Java SE Applet or pretending it has an Applet lifecycle. A separately selected original helper method remains a distinct experiment and is not called a MIDlet run. **The table and seven-variant ablation below preserve their earlier pre-Random library snapshot** (22/48, 45). A fresh current Java SE *static* coverage run now yields **23/48, 43**, still not a profile-aware MIDlet execution or an independent-runtime-defect count; see the [current MIDP checkpoint](midlet-next-slices.md).

| Input/gate | Result | What it does **not** show |
| --- | --- | --- |
| One selected original class from JAR | **0/5** decoded bodies, **11** first-blocker entries | No MIDP library graph or executable game |
| Both original JAR classes | **22/48** decoded bodies, **45** first-blocker entries | Static decoding is not linked/verified execution |
| `inspect --check` original entry | Explicit `Java ME MIDlet requires a CLDC/MIDP runtime` | Not a Java ME verifier decision |
| Native `run`, max 2,000 steps | Refuses **before game output** | No `startApp`, `paint`, input or game state |
| Builder `oak`, `js`, `html`, `java` | Each refuses **before writing a file** | No generated executable, browser game or JDK-hosted MIDlet |

The whole-JAR blockers include missing `MIDlet`, `CommandListener`, `Canvas`, `Display`, `Command`, `Graphics`, `Font`, `Image` and `java.util.Random` contracts; `DS_Effects.run()` also hits unsupported `Runtime.getRuntime`. A new **separately JDK-oracled** `String.valueOf(Object)` slice handles only a verifier-known `String` or null: in the unchanged original `DS_Effects._mth0c()V`, its first blocker moves from `String.valueOf(Object)` at byte **636** to MIDP `Graphics.setColor(I)V` at byte **658**. At that older ablation checkpoint, the whole-JAR decoded-body and first-blocker totals stayed **22/48 and 45**; this was a *shift in the first blocker*, not another decoded body or a runnable MIDlet. Some decoded methods still refer to unsupported APIs via other methods or fields. Java SE AWT `Applet` is **not** a substitute for Java ME LCDUI. The installed JDK `javap` can disassemble the original class bytes but has no Java ME runtime; this static ablation did **not** execute the original game. A **later, separate** [pinned MicroEmulator original-game reference](acotel-original-reference.md) now **does** load the unchanged original JAR, start its genuine MIDlet and capture multicolor offscreen animation state; it is **reference-side only**, not Magnolia gameplay or a differential parity pass. Full original-JAR diagnostics and per-method first blockers remain in [`acotel-midlet-ablation.json`](acotel-midlet-ablation.json).

## Separate finite MIDlet lifecycle implementation — owned fixtures only

A new [native-only Java ME probe](midlet-finite-runtime.md) explicitly isolates MIDlet behavior from the Java SE Applet/Builder host. It executes **one owned fixture MIDlet at a time**, using a small owned MIDP classfile model for `MIDlet` lifecycle notifications, `Display.getDisplay` reference identity, and `Display.setCurrent` **requested-state only**, under a cumulative **10,000-step/four-callback** cap. A real, externally pinned MicroEmulator 2.0.4 headless `MIDletAccess` reference and the new native route match the **seven-line constructor/start/pause/resume/destroy** trace on identical Java 8 fixture bytes, including local class/JAR routes; notified pause and early self-destruction are separate gates. The reference artifact stays outside the checkout. The CLI checks exact allowed MIDP and CLDC-facing types, denies Java SE desktop/network and untrusted `javax.microedition` package shadows, and refuses unknown APIs before any constructor output. A later narrow [finite synchronous Thread/Runnable fixture](midlet-next-slices.md) does **not** admit general Java SE threading or a continuous MIDP paint scheduler. The normal applet route still refuses MIDlet entries.

On **this unchanged original Acotel JAR**, the **ordinary lifecycle CLI** still refuses **before the constructor or game output** (the separate constructor-only and direct-key-method opt-ins do construct the original **unstarted** object; the [start-return opt-in](midlet-acotel-start-return.md) additionally invokes its original `startApp()` but defers the worker): after independently checked [bounded Font](midlet-font-profile.md), [Image/resources](midlet-image-profile.md) and [Random/command/thread/size/blit fixtures](midlet-next-slices.md), the first declaration barrier is now unmodeled `java/lang/Runtime` (previously `_fld3b` Thread, `_fld1b` Random, `_fld9` Image and `_fld6` Font). The nine owned profile/link/JAD/CRC/no-emission gates and pinned original JAR SHA-256 are recorded in the refreshed [`midlet-finite-reference.json`](midlet-finite-reference.json); [the Font fixture/original reference-only result](midlet-font-metrics-reference.json) is separate. Neither a requested Canvas nor a JDK/MicroEmulator *owned-fixture* callback is a rendered Acotel frame. The **ordinary** Magnolia CLI still has **not started the original game** or delivered its inputs, sprites or gameplay; the separate opt-in `startApp` return **does not run** its worker or paint; the absent icon is unresolved. A **different independent reference-side run** below *does* start the genuine original MIDlet and observe offscreen frames on MicroEmulator, but adds **no Magnolia CLDC-1.0/MIDP-1.0 handset parity**. The historical whole-JAR ablation rows below remain **22/48 and 45** at their recorded snapshot; a fresh current static inventory is **23/48 and 43**.

## Separate independent original-game reference (not Magnolia)

A newer, bounded [original Acotel MicroEmulator reference](acotel-original-reference.md) **exited 0** after starting the unchanged SHA-256-pinned game JAR in the independently pinned MicroEmulator 2.0.4, without consulting the JAD URL. It observed `ACO_XmasGifts_s` constructed, `DS_Effects` selected as current LCDUI Canvas and an actual **176×176 original-game offscreen image** with **317–357 colors** and two observed animation hashes. The wrapper verifies both JAR pins, seven observations and nonblank original pixels before writing [metadata](acotel-original-reference.json). The emulator's separate backing buffer is **not** the game's offscreen image; no actual on-screen screenshot or game input/score was checked. **Magnolia still refuses original construction on its ordinary lifecycle route** at `java/lang/Runtime` (formerly Thread, Random, Image, then Font); the separate [constructor-only](midlet-acotel-constructor.md) and [direct original key-method](midlet-acotel-direct-keys.md) opt-ins execute original bytecode **without original `startApp` or device input**; the [third original `startApp`-return opt-in](midlet-acotel-start-return.md) invokes that method but stops at a deferred-worker boundary, **without `run`, paint, device selection or gameplay**. The bounded same-bytecode fixture slices and PNG decoder are **not** original-game playback or gameplay parity.

## Isolated bytecode-function ablations

At the historical ablation checkpoint, the ablator copied the then-current Oak applet libraries to a temporary cache **outside the checkout**. Each variant disabled exactly one admission family in its copy; the verified original JAR and repository libraries remained unchanged. The saved report measures **whole-JAR all-method decoder inventory at that checkpoint**, not current profile-aware admission or feature necessity for a running game. Blocker entries can include class/field declarations as well as methods, so their count is not a count of executable functions.

| Isolated admission removal | Decoded bodies (normal **22/48** → variant) | First-blocker entries (normal **45** → variant) | Example original method newly blocked |
| --- | ---: | ---: | --- |
| Array allocation (`newarray`/`anewarray`/`multianewarray`) | **21/48** | **46** | `DS_Effects._mth7()V`, byte 4 (`newarray int`) |
| Integer array load (`iaload`) | **16/48** | **51** | `DS_Effects._mth0c()V`, byte 153 |
| Integer array store (`iastore`) | **17/48** | **50** | `DS_Effects._mth2b(IIIII)V`, byte 6 |
| Integer addition (`iadd`) | **5/48** | **62** | `DS_Effects._mth0d()V`, byte 91 |
| Integer multiplication (`imul`) | **20/48** | **47** | `DS_Effects._mth6b(I)Z`, byte 31 |
| `invokeinterface` (negative control) | **22/48** | **45** | No *new first* blocker under this already-limited inventory |
| `String.valueOf(Object)` admission for proven String/null | **22/48** | **45** | `DS_Effects._mth0c()V` returns to its original byte-636 `String.valueOf(Object)` first blocker; the normal tree reaches byte-658 MIDP Graphics instead |

The no-`invokeinterface` row **does not** establish that interface dispatch is unnecessary: a body may be blocked earlier for unrelated MIDP APIs. Conversely, a decoder delta establishes a static dependence in some original bodies, **not** that an entire MIDlet plays or that this *static ablation* implemented a Java ME API. The separate finite owned-fixture lifecycle probe above does not fill the game's library/platform deficit; falsely mapping LCDUI to AWT would not repair it.

## Reproduce, and remaining work

From the Magnolia repository root, with a separately built Magnolia binary; all commands have their own short child/network bounds. The first command downloads **only** the archive's pinned JAD/JAR; drop `--download` for subsequent offline runs:

```sh
timeout --signal=TERM --kill-after=2s 10s python3 -m unittest -v test/midlet_acotel_ablation_test.py
timeout --signal=TERM --kill-after=2s 65s python3 \
  applet-string-object-value-of-oracle.py \
  --magnolia /path/to/magnolia --routes
timeout --signal=TERM --kill-after=2s 12s /path/to/magnolia \
  midlet-entry-refusal-smoke.oak
timeout --signal=TERM --kill-after=2s 120s python3 midlet-acotel-ablation.py \
  --cache /tmp/magnolia-acotel-midlet \
  --jad /home/user/uploads/ACO_XmasGifts_s.jad.txt \
  --magnolia /path/to/magnolia \
  --download --phase all \
  --output acotel-midlet-ablation.json
# Separate owned-fixture runtime, independently compared against MicroEmulator:
timeout --signal=TERM --kill-after=2s 65s python3 midlet-finite-oracle.py \
  --magnolia /path/to/magnolia \
  --reference-jar /tmp/microemulator-2.0.4.jar --download-reference \
  --acotel-jar /tmp/magnolia-acotel-midlet/ACO_XmasGifts_s.jar \
  --output midlet-finite-reference.json
```

The original-byte ablator uses per-child **25-second** caps, per-HTTP-response **20-second** socket cap plus byte limits, and keeps the original bytes and variant copies outside the repository. A refreshed `--phase all` run with **seven** isolated ablations exited **0**; the six Python format/policy tests and all four small synthetic Oak entry-boundary checks passed. The separate `String.valueOf(Object)` oracle again exited **0** with JDK/native/standalone Oak/Node/offline HTML/generated Java equality for the exact String/null slice, and with a JDK-valid wider-`Object` fixture refused before output/emission by Magnolia (not a JDK `VerifyError`). The new owned-fixture real-MIDP/native lifecycle oracle also exited **0** and verified the original game refusal. A new **80-second native focused applet-suite attempt exited 124 (timeout)** after two progress markers; it is **not a pass**. The earlier interrupted native attempt also has no completed exit. The prior 484/484 JS exit predates this MIDlet work; a fresh bundled-JS suite and broad Builder integration result are **not** claimed. These are bounded, separate gates, not progress on the old Arc action/event, packed X11 or Java Flash player gates.

For a real game-capable MIDlet host, still implement and oracle **CLDC-1.0/MIDP-1.0** platform behavior beyond this finite, owned-fixture callback probe: JAD provisioning and protected-package policy, `Display`/`Displayable`/`CommandListener`, commands/softkeys and LCDUI `Canvas`/`Graphics`/font glyphs beyond the measured three-size `Font` profile/`Image`, resource loading, keypad, bounded scheduling/repaint, memory/work/time budgets and explicit network capabilities. Obtain independent **MIDP-1 device/reference input and display traces** before treating an original-game startup or frame as supported; MicroEmulator's Java 8 classfile fixture alone does not prove CLDC-1.0 conformance. The absent icon should be resolved or explicitly tolerated by that reference. Never download the original JAD's old IP-address URL implicitly, and keep Java ME work **separate from Java SE applet TODOs**.
