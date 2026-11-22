# Bounded MIDP Font metrics — actual bytecode, not game startup

Magnolia's separate **native-only** MIDlet route now runs an **owned** `javax.microedition.lcdui.Font` model for the three fonts requested by the unchanged Acotel game: `(FACE_SYSTEM=0, STYLE_BOLD=1, SIZE_SMALL=8 / MEDIUM=0 / LARGE=16)`. This is **not** the emulator's implementation copied into Magnolia. The model's `getFont(int,int,int)`, `getHeight()`, `charWidth(char)` and `stringWidth(String)` execute as Java classfile bytecode inside the retained Magnolia VM. `getFont` caches the same object for a repeated exact triplet; a different face/style/size throws instead of guessing. `stringWidth` totals measured ASCII 32–126 glyph widths under a 256-character bound and refuses other characters. Font **rasterization**, other device fonts, other JDK/installed-font environments and output pixels are not modeled.

The matching generic JVM `String.charAt(I)C` intrinsic supports **bounded ASCII only**, raises catchable `StringIndexOutOfBoundsException` for negative/end indices, and refuses unmodeled UTF-16/modified-UTF8 rather than giving a wrong code unit. A separate [all-target JDK matrix](../the string charAt oracle script) passed native Oak, self-contained Oak, JS/Node, offline HTML and generated Java, including catchable bounds and a mutated-member pre-output refusal. This does **not** close the [full Java String TODO](JVM TODO roadmap).

## Independent executed gates

[`midlet-font-metrics-oracle.py`](../midlet-font-metrics-oracle.py) checks two externally cached artifacts **outside** this project: MicroEmulator 2.0.4 JAR (643,996 bytes, SHA-256 `dbd5f3eb8365d3e839d6a203149e0e3776fc1a0585e16ac1fc23f76c9fcae1c6`) and optional original Acotel JAR (50,594 bytes, SHA-256 `c24e61a416616a036b537957ec9b0687ad08f9fceae18e4d06336883ec8c3040`). The JAD's URL is never opened; no third-party JAR, class or resource is committed. Each child has a process-group timeout and output cap. The owned model rebuild must match all ten checked-in classfiles; the fixture compiles **against the real emulator API** and must rebuild byte-for-byte before its identical classfiles run on both VMs.

| Gate | Observed result |
| --- | --- |
| Real default-device metrics vs measured owned tables | **285/285 printable ASCII character widths** across three sizes; metrics are explicitly tied to this JDK 11/default-device/font environment. Changes fail the gate. |
| Same fixture bytes on real MicroEmulator and Magnolia native | **39 output lines exactly match**, including repeated-font identity, heights **12/14/17**, `charWidth`, `stringWidth` on independent texts and actual game text, and valid plus two catchable out-of-bounds `String.charAt` cases; native **7,801 VM steps** in the refreshed ten-model run. This fixture merely requests a Canvas: `painted? = false`, `inputDelivered? = false`. |
| Unchanged Acotel game's actual Font fields in independent emulator | **Reference only:** its `DS_Effects._fld6/_fld7/_fld8` are genuine sizes **8/0/16**; heights **12/14/17**; widths for `Xmas Gifts` **53/67/81**, agreeing with Magnolia's *separate owned fixture*, not with an original-game Magnolia execution. |
| Explicit negative gates | A valid `FACE_MONOSPACE` real-device font is accepted in MicroEmulator but refuses **before output** on Magnolia's narrower profile; changed `Font.getHeigHt()` refuses during preflight before output. |
| Unchanged original JAR on Magnolia | Still **refuses before any original constructor, callback, game output or frame**. After the separately checked [Image/resources](midlet-image-profile.md) and [Random/command/thread/geometry/blit fixtures](midlet-next-slices.md), its current first declaration barrier is `java/lang/Runtime`; earlier declaration barriers were `_fld6` Font, `_fld9` Image, `_fld1b` Random and `_fld3b` Thread. This is **not** startup. |

The [metadata-only result](midlet-font-metrics-reference.json) and [separate original-game offscreen reference](acotel-original-reference.md) remain distinct. This desktop Java 8 classfile 52 fixture does **not** establish CLDC-1.0 preverified 45.3, handset-font parity, full MIDP, original-game resource use/`Image`/`Graphics` drawing, command dispatch, scheduled paint/input, or game playability in Magnolia. A **separate owned Image fixture** now loads explicitly selected original PNG members without starting the original game. A fresh **Java SE Applet static** inventory after the shared Random-library additions is **23/48 bodies, 43 first-blocker entries**; the historical **22/48, 45** ablation snapshot is a different earlier profile and neither count proves MIDlet execution.

## Reproduce from the repository root

```sh
# Both JARs must already be independently hash-checked, and outside checkout.
timeout --signal=TERM --kill-after=2s 35s python3 \
  midlet-font-metrics-oracle.py --magnolia /path/to/magnolia \
  --microemu /path/to/microemulator-2.0.4.jar \
  --original-jar /path/to/ACO_XmasGifts_s.jar \
  --output midlet-font-metrics-reference.json

timeout --signal=TERM --kill-after=2s 110s python3 \
  the string charAt oracle script --magnolia /path/to/magnolia
```

**Next real dependency:** the [bounded Image/PNG/offscreen Graphics fixture](midlet-image-profile.md) and [five follow-on slices](midlet-next-slices.md) now match independent references, but the unchanged original still refuses before construction at `java/lang/Runtime`. Supply an independently checked, non-fabricated Runtime/free-memory contract, remaining LCDUI Graphics/Canvas/command **delivery** and glyph profile, then a continuous metered scheduler, actual device paint/key events and JAD rules; only after **original** Acotel bytecode starts and produces independently matching frames/input/game states in Magnolia can this game's functional gate close. Adding class names or moving a static first blocker without executing and comparing the original output does not count.
