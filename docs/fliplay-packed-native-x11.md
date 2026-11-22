# Original Fliplay — packed native X11 device pixels (bounded checkpoint)

This is an artifact-specific real-device check, **not** completion of every Magnolia media track, general AWT, or general FLI/FLC. The separately obtained unchanged `fliplay.class` (SHA-256 `bc315872dafebbf3c634e64c0fc251b38031ed49b7c4698d39c94990de9fe20b`), `flickframe.class` (`12feb5321cd8c9aeb0bf285f055596d981e3d45ccfb8c62d029f2f6d95ccdf9f`) and raw Moon Clouds FLC (`812379899d151a1460366249fd794899eeefe3c7bacdc528da22c2624b33cff4`) remain outside the checkout. Do not redistribute the resulting executable without reviewing their rights.

## Older completed final-device gate, separate from the newer completed four-sample gate

The previously built **self-contained** 14,766,405-byte native executable at `/var/tmp/oak-packed-isolated/player` (SHA-256 `27eb9ba532a6fe40e4148776072702f72a4e0614d39efbede4a1ca642230f008`) ran via `xvfb-run` with a 540-second child cap and **exited 0 after 433 seconds**. It executed the original JVM `fliplay` + `flickframe` over the full FLC, presented all **61/61** states, delivered eight selected original callbacks and reported frame 60, 125 repaints, 68 statuses. Its actual `XGetImage` 320×240 **307,200-byte device ARGB** capture was byte-identical to the separately executed JDK 11 `-Xverify:all` original lifecycle's post-event screen (SHA-256 `e054c58cf2ff7e5273fef9e8d738e431f6b79e79a4abb299e520aa097a248ddc`). This closes the previously *uncaptured packed final-screen result* for this one movie and 24-bit Xvfb visual; it does **not** demonstrate 61 device-backed intermediate frame captures. This older packed executable predates the new optional audit callbacks; do not pass it off as a result from the changed code.

A **fresh, separate unchanged-class JDK reference run** also exited 0 twice: once with `--full` exporting 61 original RGB frames and once with `--events-full` exporting the final event screen. The JDK harness source historical fixtures: FliplayLifecycleJdk.java` has SHA-256 `97918bff5263507fbc083aa8bbb39a7d3c6a3b5cca7403310adc1421bdfa9707`. Its SHA-256 results for frames 0, 1, 30, 60 are respectively `979353685a444821dc47857965156d1dd45dfe6f55495f2c591653584f9c504f`, `a045f7d74823cd175adbc7ec5ab4b25f6c3f7cef35636ab64b938f639924f21c`, `d229b03b1a139805a131abda24b97561b5a1507437e708812ceb6b9dc2c00879`, and the frame-0 digest again. Its aggregate first-60 RGB digest was `dfb9f2fbde80dafbee976ae43ec2c04c1019d7ffccb13807400324465f89fe16`. The JDK post-event raw ARGB file byte-matched the **older, completed packed final-screen run**.

## New optional audit; completed run after the earlier interruption

`lib/applet-gui.oak` now passes the actual window as an optional **third** argument to a trusted frame observer after `endFrame`. Existing two-argument observers ignore it. `fliplay-oak-cli.oak --mode native --oracle yes --oracle-frames-dir DIR --argb FILE` accepts exactly four geometry- and SHA-pinned JDK RGB references **before original construction**, samples the *displayed X11 window* with `XSync`/`XGetImage` after frames 0/1/30/60, checks 24-bit TrueColor masks and byte order, compares all 230,400 sampled RGB bytes to their independent JDK files, and separately exports the selected post-event final ARGB. It refuses the audit without an explicit native oracle; no network, media rewrite, or unverified X11 visual fallback is added. The source pack includes no third-party classes or movie in the Git tree; the explicitly packaged *external* executable embeds them for local testing.

The newly packaged executable (SHA-256 `b4450692e0b1a857a645e8f4623da293395bab77fde0c31d6fed97203ad98bac`) **now completed a subsequent full four-sample Xvfb run with observed exit 0** under a 960-second child cap (560.7 seconds wall). Four *physical X11 device* RGB captures at frames **0, 1, 30 and 60** each matched the SHA-pinned independently generated JDK 11 original-class RGB reference byte-for-byte; it drove all 61 original image states, eight original post-pause callbacks, and reported final frame 60, 125 repaints and 68 statuses. Its final `XGetImage` 307,200-byte ARGB file (SHA-256 `e054c58cf2ff7e5273fef9e8d738e431f6b79e79a4abb299e520aa097a248ddc`) was `cmp`-identical to the JDK event reference. A metadata-only [machine-readable verified evidence report](fliplay-x11-four-frame-audit.json) records the sample pins, exact independent/reference equalities, completed exit, and negative. The new run log is `/var/tmp/magnolia-every-probe/packed-audit-completed-run.log` (SHA-256 `d344f5fd9e1231aa37c03237b873f84ed08367e06aecdd834c3b34fc6cd61882`); raw captures and original media remain external.

The earlier same-binary run **was interrupted** after matching 0, 1 and 30. It was *not* a second completed gate. An earlier 12-second corrupt-reference attempt **timed out**; a new same-length **byte-tampered `000.rgb`** with the other three true references ran under a 160-second cap and **exited 1 in ~15 seconds** with `independent JDK RGB reference refused before construction: 000.rgb`, and no ARGB or original-image lines. Missing-reference and two option-policy negatives separately refused before original construction; the six-pixel Xvfb/XEvent GUI smoke exited 0 earlier. No whole cross-media or twelve-oracle regression ran with this X11 audit.

### Reproduce the bounded device comparison

Use the SHA-pinned external classes/FLC and a cgo-enabled native Magnolia. Run the owned JDK reference (with the classes copied into the same temporary class directory) under `-Xverify:all`, `-Djava.awt.headless=true`, `FliplayLifecycleJdk raw RAW_FLC JDK.argb --full JDK_FRAMES`, and a second time with `--events-full` for `JDK_EVENTS.argb`. Compile the harness with JDK 11 `javac --release 8 -g:none -proc:none -cp CLASS_DIR -d CLASS_DIR FliplayLifecycleJdk.java`. Check all class/movie/harness and reference frame digests above *before* running the packed program. Build the current source pack with:

```sh
CGO_ENABLED=1 go build -o /tmp/magnolia-x11-interp .
python3 fliplay-oak-native-pack.py \
  --interp /tmp/magnolia-x11-interp --fliplay /path/to/fliplay.class \
  --flickframe /path/to/flickframe.class --movie /path/to/moonclouds.flc \
  --output /tmp/magnolia-packed-fliplay
# This is intentionally a slow, explicitly capped real-X11 original-bytecode gate.
timeout --signal=TERM --kill-after=5s 900s \
  xvfb-run -a -s '-screen 0 640x480x24' \
  /tmp/magnolia-packed-fliplay --mode native --oracle yes \
  --oracle-frames-dir JDK_FRAMES --argb /tmp/packed-final.argb
cmp /tmp/packed-final.argb JDK_EVENTS.argb
```

The **observed exit 0** with all four expected `displayed X11 JDK frame` lines, the final `cmp` exit 0, and the separate corrupted-reference refusal was recorded above. This closes only the **four sampled device-frame checks for this pinned original movie**; it does not claim all 61 intermediate physical device captures. Resized/Expose, nondefault visuals, physical Win32 and broader original Fliplay functionality are open; the separate [standalone FLI/FLC decoder](flic-functional-slice.md) is not original JVM execution. The user-requested **every-media-track** completion remains open in `docs/historic-functional-gates.md` and `docs/completion-pass.md`.
