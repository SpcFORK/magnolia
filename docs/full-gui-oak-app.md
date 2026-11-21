# Oak-built original Fliplay app: full GUI + headless contract

This is a **bounded original-bytecode player slice**, not completion of [the JVM TODO roadmap](JVM TODO roadmap). It runs the hash-pinned original `fliplay.class`, `flickframe.class` and Moon Clouds FLC under Magnolia's metered JVM. The selected `java.awt.Event` bridge is not a historical AWT window. The original Java Flash 2 player, arbitrary applets, complete JVM verification and a continuous event scheduler remain open.

## Architecture and control contract

- [`lib/fliplay-control.oak`](../lib/fliplay-control.oak): one explicit `openHeadless`, `openHeadlessInteractive`, `openGui`, `status`, `input`, `command`, `redraw`, `close` API. Classes/movie are hash/size checked before linking; the same selected bytecode and 40-million-step / 16-event limits apply in each mode. Plain `openHeadless` uses no GUI, independently captures 61 original states and **refuses** post-pause input. The **opt-in** `openHeadlessInteractive` retains the original JVM object and delivers the same bounded, purely validated keyboard/mouse/command records after the finite run **without any window, DOM, or browser queue**; it does not turn the headless oracle into an unbounded event scheduler. `openGui` likewise retains the original object after the finite run.
- [`lib/applet-gui.oak`](../lib/applet-gui.oak) and [`lib/GUI.oak`](../lib/GUI.oak): full facade dispatches bounded opaque packed-ARGB images to Web `gui-web` commands, native X11 exact `XFillRectangles` batches, or GDI exact-rectangle fallback. The native X11/Win32 normalizer recognizes printable keys, left-button clicks/releases, motion/drag and damage; it refuses unsupported or out-of-bounds input before the original `handleEvent` gateway. Web input crosses GUI's own queue and is revalidated. The native X11 backend needs Linux `cgo` to resolve shared-library functions; raw kernel `syscall.Syscall` is not a valid Xlib call.
- [`fliplay-oak-app.oak`](../fliplay-oak-app.oak): **the page structure, style, Canvas, accessibility labels, local timeline and original-input controls are Oak code**, compiled by Magnolia. [`fliplay-oak-worker.oak`](../fliplay-oak-worker.oak) runs the same control contract in a Blob Worker, checks source bytes again, transfers original image pixels and returns only sanitized control state. The Python builder only checks/normalizes supplied bytes, invokes Oak builds and packages two files; it does **not** hand-author a page.
- The generated `index.html` is **164 bytes**: a title and **one external `<script src="./app.js">`**, no inline script, inline stylesheet, Worker script tag or hand-written UI. `app.js` embeds the checked original class/FLC bytes **as data**, compiled Oak page code, and the compiled Oak Worker source; it embeds **zero pre-rendered frames**. The pair works offline in a normal browser; the Arena file viewer cannot load its external script because that preview iframe blocks network requests, so use the live server preview or download both files together.
- The app distinguishes **captured timeline replay** (seek/loop/speed/previous/next of 60 images computed during the finite original run) from **original Java controls** (space, comma, period, R, plus, minus, zero, focused printable keys, left mouse/drag). The latter call original `handleEvent` on the retained object, retain step/event budgets, and repaint only if the VM does. A timeline scrub is **not a Java seek**. `input` is one bounded event at a time, not an unbounded desktop input queue.

## Reproduce without committing historical bytes

Use a Go 1.26 compiler. On Linux the native X11 path requires **`CGO_ENABLED=1`**, a C compiler and X11 shared libraries. Paths below are to your separately obtained, original, hash-checked historical inputs.

```sh
cd /path/to/magnolia
CGO_ENABLED=1 go build -o ./magnolia .
python3 fliplay-oak-app-build.py \
  --fliplay /path/to/fliplay.class \
  --flickframe /path/to/flickframe.class \
  --movie /path/to/moonclouds.flc \
  --magnolia ./magnolia --output-dir /path/to/oak-moonclouds
# An exact original one-entry Moon Clouds ZIP is also accepted as --movie.
python3 -m http.server 8103 --bind 0.0.0.0 --directory /path/to/oak-moonclouds
# Or open index.html beside app.js locally in a browser permitting Blob Workers.
./magnolia fliplay-oak-cli.oak --mode headless \
  --fliplay /path/to/fliplay.class --flickframe /path/to/flickframe.class \
  --movie /path/to/moonclouds.flc
./magnolia fliplay-oak-cli.oak --mode native \
  --fliplay /path/to/fliplay.class --flickframe /path/to/flickframe.class \
  --movie /path/to/moonclouds.flc
# Optional finite headless input after the same independently checked lifecycle:
./magnolia fliplay-oak-cli.oak --mode headless \
  --fliplay /path/to/fliplay.class --flickframe /path/to/flickframe.class \
  --movie /path/to/moonclouds.flc \
  --commands faster,slower,reverse,next,resetSpeed
# Alternatively: --input-file events.json, a 1–16-entry <=4096-byte JSON list.
# Example: [{"type":"keydown","key":114},{"type":"mousedown","x":36,"y":12,"button":0}]
# The CLI refuses a command list and JSON input file supplied together.
```

The command and JSON inputs are prevalidated before original-bytecode execution; only the explicitly opted-in retained session accepts them, using the same 320×240 GUI-coordinate and printable-ASCII/left-mouse boundary. `control.command` also supports play/pause, previous/next, faster/slower, reset speed, reverse/forward/backward; the plain headless run remains independent and noninteractive. `--mode native` opens a native GUI (X11 on Linux, Win32/GDI on Windows), pumps selected native input, redraws after X11 Expose and closes with the window. `--oracle yes --argb FILE` is an X11 test-only option for the eight JDK-input cases and captured displayed screen. The bounded original-class X11 oracle **exited 0** after 61 image states and eight selected input callbacks, and its final displayed 320×240 ARGB was byte-identical to the independent JDK final screen. Intermediate native screens were not individually screenshot-compared. **No physical Windows device/visual parity has been tested**; GDI exact-rectangle fallback is slower than an optimized DIB upload. Only 24-bit TrueColor Xvfb pixels have been compared; other X11 visuals are **not yet validated or fail-closed** and may render incorrect colors. Display resizing, IME/complex keys and accessibility of the original JVM are also unverified.

## Evidence and honest boundaries

| Gate | Result |
| --- | --- |
| independent plain headless original run | Standalone CLI **exit 0**: 61/61 original states, 122 repaints, 63 statuses; input on the non-retained session refused |
| **opt-in retained headless control** | **exit 0 without a GUI**: 61 captured states; 8/8 original JDK event outcomes/repaint/status/fields; final 307,200-byte ARGB byte-equals JDK; 16-event ceiling, malformed input and closed-session input refused. The separate standalone ZIP + JSON CLI run is a further gate. |
| new Oak full-GUI Worker under Node vs independent JDK 11 | 61/61 RGB images, 8/8 original event results/pixels, three original event redraws, tampered-class and malformed-input refusal |
| Oak-built DOM fake-DOM check | 61 ordered frames, cached transport does not send JVM events, original command and rerun wire correctly |
| **actual Chrome** DOM/Canvas/input | single script/style from Oak; Canvas frames 0 and 1 match independent JDK; focused key reaches original reverse callback; physical mouse press/release reaches original input; no page exceptions |
| Xvfb native X11 facade | **6/6 displayed pixels exact** via real `XGetImage`; separately injected printable key, click, held drag and release all reach `GUI.pollBoundedInput` with exact bounded records |
| Linux cgo entry | `getpid`, `strlen`, missing-symbol and over-arity Go unit gate passed; rebuilt runner opens a real X11 window |
| **full original native X11 GUI** | **exit 0**: 61/61 original image states displayed; 8/8 selected original event outcomes/fields; final `frame=60`, `repaints=125`, `statuses=68`; real X11 screenshot byte-equals the independent JDK **307,200-byte ARGB** (`e054c58cf2ff7e5273fef9e8d738e431f6b79e79a4abb299e520aa097a248ddc`). Not 61 individually captured native screen comparisons. |

The bounded native/JDK oracle from the **previous** `gui-web`-only milestone is separate; this gate displayed pixels on a real X11 device. A later **self-contained packed-native** executable independently exited 0 with 61/61 original states, eight inputs and a JDK-byte-identical final actual X11 screenshot; its newer optional intermediate-device audit reached matching physical frames 0/1/30 but was interrupted before frame 60/completed exit. [Packed gate and explicit incomplete audit](fliplay-packed-native-x11.md). The **post-facade/FFI** native focused suite and independently rebuilt bundled-JS suite each exited 0 at **433/433**; these remain focused-subset tests, not whole-JVM certification. A later opt-in headless-control edit changed the shared adapter after those suites; its dedicated original-class headless/JDK gate and rebuilt final Worker gate exited 0, while a full repeat of both focused suites after that last adapter edit is not yet counted.

## Why full GUI rather than only `gui-web`?

**Advantages:** one public GUI facade and player contract across Windows, Linux, Web and headless; original JVM state/input stays host-independent; Oak page code replaces a large hand-authored HTML template; a single external loader script makes the generated app inspectable/cacheable and Blob Worker keeps long JVM execution off the UI thread; native-screen evidence can be compared to JDK and browser Canvas independently.

**Costs/trade-offs:** the full-GUI Oak Worker imports about **109 modules** and is larger/slower to build than `gui-web` alone; Linux native calls need a supported cgo build and matching X11 TrueColor visual; Windows GDI per-run fallback is unoptimized and needs physical test coverage; offline browser Blob Workers may be blocked by some embedding policies; frame caching uses memory and is not an original-Java seek; the finite 16-event/40M-step ceilings intentionally stop broad interactivity. The split lets a simpler headless oracle stay independently verifiable instead of pretending a browser Canvas is AWT.

## Remaining-work plan (ordered acceptance gates)

1. **Extend native backend evidence:** the selected 61-state/eight-event/final-screen X11/JDK oracle and real injected key/click/drag/release probe now pass. Add separate X11 screenshot comparisons for *intermediate* frames, resize/Expose, nondefault visual-mask refusal, and display teardown; test actual devices beyond injected XEvents. Add optimized GDI DIB upload and physical Win32 input/render parity. Fail closed for unsupported visuals, frame surfaces and event variants. Re-run focused native/JS suites after any further adapter edits.
2. **Continuous interaction only with a proven scheduler:** replace the selected post-pause 16-event gate only after metered event scheduling, repaint/focus lifecycle, `stop`/restart and multi-task contention are designed. Preserve the finite headless oracle and reject unsupported window/network/audio input explicitly. Add independent JDK and real-device event timelines.
3. **Close JVM Stage B/C TODO items independently:** full verifier/uninitialized/StackMap/exception edges, remaining long/F/D conversions, UTF-16 strings, `jsr`/`ret`, implicit errors, thread/monitor services. For each accepted opcode/library slice, compare a small JDK-verified fixture on native Oak, JS and optional generated Java; do not infer whole-player support from static decode counts.
4. **Broaden libraries and player gates:** AWT fonts/images/audio/resource/URL permissions; run whole original Flash 2 only with a Flash-2-compatible movie and an independent reference JDK. Tackle selective multiple-class modern-JAR linking and its resource/size budgets separately from the old Fliplay artifact; signed-JAR/class-relocation work requires a real assembler/frame strategy.
5. **Release discipline:** record exact source hashes, preserve class/JAR original-byte restoration, build raw and ZIP artifact byte-identically, run clean replay/delta checks and avoid redistributing historical binaries without reviewing rights. The earlier hand-written pages remain *legacy regression fixtures*, not the new deliverable.
