# Finish-pass execution evidence and remaining acceptance

This is an evidence index, not a claim that the broad JVM, Flash or ActionScript TODOs are complete. A fresh bounded census reproduced the pinned [census](media-blockers-current.json) byte-for-byte and remains **21 unchecked JVM entries and 24 unchecked ActionScript entries**. Its first attempt hit a 25-second child cap; a separate run with an explicitly bounded threefold timing allowance exited 0. A checkbox counts only when its own full stated acceptance has been met. Reference-only playback and narrowed fixture/preflight gains are not checklist closures.

## Further unchanged GEMALOGO FLI

The independently reproduced [GEMALOGO original-class gate](fliplay-gemalogo-original.json) runs unchanged Fliplay/flickframe classes and the external pinned `INVEX96/GEMALOGO/POKUS3/GEMALOGO.FLI` (58,427 bytes, SHA-256 `7708e1cea2a134bf76d6dc93ab0480d15a71affa080ba99efbda250d07a4a2b0`). Fresh JDK 11 and native Magnolia headless runs match **6/6** 320×200 RGB states (five movie frames plus the original ring), 12 repaints, eight statuses and the entire 256,000-byte final ARGB. All six RGB states happen to be identical; this is not evidence of visual frame changes. Go and the owned Python decoder independently agree on **5/5** codec frames, but none has the original JDK's RGB hash, so those references must not be conflated. Same-length edits of the original class and movie both refuse before original execution/output. The caller explicitly caps arrays at **68,133** cells because the original player allocates the first 58,133-byte frame payload plus 10,000 bytes; default array and resource limits are unchanged. This does not establish device/interactive input, sound or general Fliplay support. Reproduce with [`fliplay-gemalogo-oracle.py`](../fliplay-gemalogo-oracle.py) and its external pinned inputs.

## Two additional original Fliplay FLI movies

[Fresh direct-execution report](fliplay-two-new-original-fli-fresh.json), [earlier consolidated report](fliplay-two-new-original-fli.json) and [`fliplay-new-fli-oracle.py`](../fliplay-new-fli-oracle.py). These are unchanged true FLI files (magic `0xAF11`) obtained from the externally held [Panda381 repository at the pinned commit](https://github.com/Panda381/DOS-Progs/tree/f6d50182477260b87b5c79df70bb258834c040ce). They are not included here; verify their rights before redistribution. Original 1999 Fliplay and flickframe class files are also external and unmodified.

| Input | Bytes | Movie SHA-256 | Original JDK vs Magnolia displayed RGB, including ring | Repaints / statuses | Complete final 320×200 ARGB SHA-256 |
| --- | ---: | --- | --- | --- | --- |
| `HRY/PUZFLI/GLASS.FLI` | 45,854 | `3aaca4f8541ba73890830687e5c436258caa7e1a4075156f23e249f0a9e6908c` | **21/21** | 42 / 23 | `a8b9bceac44ef55d6b8b7a04266ae1e1cfd50c25accc895b3c2953788000369b` |
| `GRAFIKA/AAPLAY/CITY.FLI` | 28,728 | `2419643ec20381e10201f4e82fbfc5c0afea349b27b8690121d2362aa4d9452a` | **4/4** | 8 / 6 | `0e7885778b1f35fd362f854089b9b974b0887020946d856d1e6b83769a886cd9` |

Both complete **256,000-byte** final ARGB arrays byte-match the fresh JDK reference. Both original-bytecode headless runs exited **0**, consuming 9,494,742 and 3,167,406 charged steps. The earlier consolidated report reran JDK 11 with `-Xverify:all`, Go and Python and verified SHA-pinned logs and final ARGB bytes from separately completed native runs; it did **not** rerun positive native execution in that invocation. A later [fresh report](fliplay-two-new-original-fli-fresh.json) independently reran **both positive native executions** as well as fresh JDK/Go/Python checks and exited **0**. The four same-length class/movie edits refused before original construction and output. For a direct native rerun omit the optional `--glass-observed-*` and `--city-observed-*` arguments; it needs considerably more time.

Independent Go and the owned Python FLI codec agree on all **23/23** non-ring movie frames. Their COLOR64 RGB does **not** match historical Fliplay's JDK output on any of those 23 frames, so the JDK/original bytecode, not the modern decoder, is the display parity reference for these two movies. The input cap was not increased: each input is below the existing 524,288-byte per-resource limit. An unrelated `test2.fli` remains over that limit and its JDK-only result is not a Magnolia pass. These two bounded routes do not establish general Fliplay, GUI device/input, FLI/FLC variants or Flash.

Reproduce with local, independently verified inputs and a pinned build. The oracle checks its own pins, so point to the exact compatible reference source/binary versions recorded in the report:

```sh
python3 fliplay-new-fli-oracle.py \
  --fliplay /external/fliplay.class --flickframe /external/flickframe.class \
  --glass /external/glass.fli --city /external/city.fli \
  --go-source-dir /external/go-flic-refmod \
  --magnolia /external/pinned-magnolia \
  --output /external/new-fliplay-report.json
```

## User's unchanged SWF5: real-player input reference, not Magnolia playback

[Recorded reference](uploaded-swf5-ruffle-postinput.json) and [`uploaded-swf5-ruffle-postinput.py`](../uploaded-swf5-ruffle-postinput.py) load the entire unchanged 2,418,588-byte user upload (SHA-256 `2bde4d96b5b6695f8dd81afe005080c2bc6b9a14545f7cc8f35f49c71ac43cbc`) in pinned Ruffle v0.7.1 under an actual X11 display. At 500×350, a real pointer enters and presses the original visible `play` button at `(245,322)`, then releases it. URL opening, TCP and guest filesystem access are denied, with memory-only storage. Captures are made from the actual Ruffle window; this is neither a screenshot approximation nor a Magnolia screenshot.

The initial screen is stable before input. Hover and press show two distinct screen hashes; after release the displayed screen changes in **18,992/175,000 RGB pixels** versus the initial capture and is stable in the captures at one and three seconds. An exploratory observation at a later point also remained unchanged; do **not** infer that the SWF's 2,140-frame movie ran, that audio played, that all actions worked, or that this static post-click screen is the intended movie behavior. Ruffle's first display reports SWF5, 500×350 and 20 FPS. A same-size changed source refuses before player launch and report creation. No SWF or captured image is redistributed.

This is a **new independent post-input reference** for eventual comparison, not an Oak playback gate. A later, separately checked [virtual button-release slice](avm1-swf-button-prefix.md) executes all six original Play-button AVM1 records and emits one virtual `_root.play()` call after synthetic complete-load/loaded-frame selection; it adds **no actual scheduling, physical hit test, pixels or audio**, and has not been compared to Ruffle's displayed pixels. Full Flash acceptance remains open. Reproduce using external SHA-pinned reference binaries and an available `xvfb-run`/ImageMagick:

```sh
python3 uploaded-swf5-ruffle-postinput.py \
  --swf /external/asdf.swf.txt --ruffle /external/ruffle \
  --xdotool /external/usr/bin/xdotool --libs /external/xkb-libs \
  --output /external/new-swf5-reference.json
```

## Separate SDK media-host gates

These independent bounded results do not run the historical Java Flash player or Oak AVM and do not by themselves finish a whole SDK utility:

* [SDK SWF1 original demo](swf1-sdk-real-player.json): 13 actual Ruffle vs owned SVG/Chrome screen states; 2,848,768/2,860,000 RGB pixels exact across the comparison.
* [Original C++-exact URL substitution](sdk-urlsubst-real-playback.json): 599→551-byte output, exact original C++ bytes and 13 identical real-player pixel hashes; three URL actions denied.
* [Original C++-exact JPEG assets](sdk-jpeg-real-playback.json): placed and displayed in Ruffle; 16/16, 227/256 and 58/63 source-pixel exact matches with maximum channel difference 2; 335/335 pixels differ from unplaced controls.
* [Original C++-exact 4,098-sample ADPCM asset](sdk-adpcm-real-audio.json): bounded Ruffle audio capture contains 16,384 non-silent stereo frames; decoded-reference waveform correlation 0.999953, while controls remain silent.

See individual reports for the exact reference artifacts, controls, caps and malformed-input refusals.

## Current regression status

* The fresh native **focused applet suite passed 530/530**, exit 0, on the earlier build SHA `4a6e0b42dabe6225ff61805a577f95346dc7e448b3d1cc15b56c69760cfe29cd`. After the later exception changes, native and rebuilt Node applet suites each passed **531/531** with executable SHA `b19e80c6d12f4e2c05c2fbab70e94ae907eb9eba6e0d63e34b655b23483adb66`. An earlier native attempt timed out after 420 seconds (exit **124**); the separate longer-capped attempt exited 0. These are applet-focused, not full release suites.
* The original standalone audio run timed out after 900 seconds (exit **124**) because the Ogg page-builder used a non-returning expression as a recursion guard. That defect and similar invalid-data guards were corrected; the final rebuilt native Xvfb runner then passed **2,971/2,971**, including audio, but excluding applet and `mg`. The applet native suite separately passed **530/530** on an earlier build; bundled-JS applet passed **530/530** after the later JS numeric correction. The actual full native suite has **not** passed: `test/mg.test.oak` imports absent `ext/mg/comps/scene-manager`. The Xvfb native result is `/var/tmp/magnolia-finish-2026-11-21/native-rest-audio-fixed3.log`, binary SHA-256 `cabdd44ce90106a44badc87dfc7d076ff0d815b67a871fc6bd7d50711b71d617`. Temporary logs/binaries are not bundled.
* Native audio plus core tests and bundled JS/Node audio plus core tests both passed **405/405**. Fixes include PCM unsigned silence at byte 128 across WAV/raw helpers, correct 10-byte AIFF extended-rate field, using math builtins instead of shadowing them with absent imported functions on JS, and JS `int()` flooring negative halves like native. The [owned PCM reference report](audio-pcm-ogg-aiff-owned.json) independently checks Python `wave`'s byte values, exact AIFF layout, and two encoder-page Ogg CRCs. The Ogg format here is a *custom PCM16-in-Ogg* profile, not Vorbis or Opus, and non-16-bit requests refuse. Native tests include every truncated prefix of a tiny valid Ogg. **The Magnolia Ogg parser does not verify CRC**, does not claim general malicious-input hardening or arbitrary Ogg support, and the AIFF/WAV/JS sample cases are not complete format or browser-audio acceptance.
* During that run, the pack tests were corrected to expect actual MessagePack instead of stale JSON; the focused pack runner passed **12/12**. Stale AI default and Go-codegen preamble expectations were corrected and their focused runners passed **172/172**. GUI tests need Xvfb in this environment. Embedded `GUI` imports were corrected to resolve embedded sibling modules by name; a Go regression imported it from an empty working directory, and a freshly bundled JS GUI/window smoke passed under Node. This is not full browser GUI acceptance. `go test ./...`, **77** focused Python media/format tests, and `git diff --check` passed after those changes. A broad Python discovery attempt timed out and was terminated; it is not a pass.
* An earlier JS binary-literal probe exposed replacement characters for raw non-UTF-8 high bytes. The later [web compatibility pass](web-compatibility-progress.md) fixes the literal boundary and tests all 256 bytes plus UTF-8 in Node and Chrome. The original uploaded SWF's SHA-pinned 32-byte button action now also matches as a raw literal across native/Node/Chrome with six virtual records and legacy-host refusal; no physical event or whole-SWF JS parity is claimed.

## Remaining broad gates

Original DS has only preflight/owned-raster improvements. An owned JDK-differential fixture for the actual `MalformedURLException` hierarchy uncovered and fixed a separate JS handler-key bug: a typed catch was previously treated as catch-all because `catch` was mangled during JS property access. After that fixture passed native, standalone Oak, bundled JS and generated Java, the selected unchanged class still refuses **before construction** at `_mth2c()V` byte 6, unsupported `URL.getHost`; no original worker/image output was established. The [fresh later-corpus static report](later-applet-corpus-current-exceptions.json) records this changed refusal; DS remains **24/32 decoded bodies, nine first-blocker entries**. A separate JDK-valid owned URL-constructor case refuses before guest output and performs no networking. The previous broader exception matrix first timed out at its standalone-Oak build (40-second stage cap), then a second attempt timed out at zero-edit recompile. A third attempt with explicit 160-second build/recompile caps **exited 0**: the 13 JDK lines matched native/Oak/JS/HTML/Java, original class restored byte-exactly, checked edited class matched JDK, and a tampered handler table refused. The earlier timeouts are not passes. See the [owned exception report](jvm-malformed-url-exception-owned.json) and [`applet-malformed-url-exception-oracle.sh`](../applet-malformed-url-exception-oracle.sh). The uploaded Flash5 file has no Magnolia post-click/later-frame/audio rendering. Original Acotel worker/scheduled repaint/input gameplay is not established. The newer AS3JS comparator is **not** Magnolia `.as` or AVM2 execution. Complete original JVM applet/player behavior, AVM1/AVM2, broader media formats/SDK tools and the full cross-target release matrix therefore remain open. Timeouts and incomplete test runs must be reported as such, not counted as passing tests.

## Controlled audio, rendering, scheduling and resource ablations

The finite experiment and its distinctions are recorded in
[`av-capability-ablation-results.md`](av-capability-ablation-results.md) and
[`av-capability-ablation.json`](av-capability-ablation.json).
Fresh independent JDK and native unmodified controls matched all six original
GEMALOGO/Fliplay states and final pixels. Explicitly denying offscreen raster
blit or bounded `Thread.start` stopped the original workload at that named
boundary; high-level WAV encoder denial stopped an independently checked owned
WAV fixture but did not affect the silent original. Omitting the resource
exposed an existing PrintStream object-format failure, not a clean resource
refusal. Audible playback, other renderers/schedulers, original SWF and MIDlet
playback were not demonstrated; no whole-program TODO is closed.

### Changing-frame and external-audio follow-up

[`av-ablation-changing-frames.md`](av-ablation-changing-frames.md) records a
fresh 21/21-state JDK/native original GLASS control (11 distinct state hashes),
independent blit and thread-start denials, their combined first-failure masking,
and a test-only successful-no-op blit rejected on the first RGB mismatch. A
separate fresh Ruffle/ALSA reference validates real mixer mute and unplaced
silence against its audible SDK-exact ADPCM fixture; this is **not** Magnolia
playback or physical-speaker evidence. No whole-program TODO was closed.

### Frame capture versus repaint

[`av-ablation-capture-repaint.md`](av-ablation-capture-repaint.md) records a
fresh original CITY/JDK experiment with one shared adapted caller. Omitting the
optional capture sink preserves final pixels and measured lifecycle counts but
has no intermediate-frame evidence; the instrumented run checks four states.
Explicit repaint and repaint-drain denials stop at their respective boundaries.
A test-only false-success repaint no-op instead reaches a bounded Thread.sleep
refusal; it is not a supported runtime behavior. No production source or guest
original was modified, and no whole-program acceptance item was closed.

### Web byte boundaries and actual-browser smoke

[`web-compatibility-progress.md`](web-compatibility-progress.md) records the
fixed high-byte JS literal/output boundary, related Unicode/UTF-16/Node path and
fake-fetch checks, and a selected seven-line JDK-versus-real-Chrome offline
applet-method test. Final JS/Node suites passed applet 531/531, audio/core 405/405
and builder 21/21; this is not full browser, full JVM, SWF or media acceptance.
The earlier binary-literal finding above is superseded by this scoped fix.
