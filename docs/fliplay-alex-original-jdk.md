# Second original Fliplay movie: `alex.flc` in the bounded JVM host

**Observed executable gate:** [`fliplay-alex-oracle.py`](../fliplay-alex-oracle.py) **exited 0**. Two unchanged original JVM 45.3 classes (`fliplay.class` SHA-256 `bc315872dafebbf3c634e64c0fc251b38031ed49b7c4698d39c94990de9fe20b`, `flickframe.class` SHA-256 `12feb5321cd8c9aeb0bf285f055596d981e3d45ccfb8c62d029f2f6d95ccdf9f`) ran against **a second original movie**: externally supplied `alex.flc` (37,297 bytes, SHA-256 `f61ab650ffb0b81a3e803c8f6e279c82ac0b8d27bcbf81808586efef0a70c212`). Neither third-party classes nor media are in the source/reports. Redistribution rights are unconfirmed.

The original `fliplay` bytecode actually executed `<init>`, `init`, `start`, its worker, `load_fli`, palette and pixel updates, `repaint`/`update`, `next_frame`, `paint`, and paused at frame **2**. Against an independently executed **JDK 11 `-Xverify:all`** instance of the same unmodified classes, Magnolia's finite headless JVM matched **all three** displayed **128×100** RGB states (frame 0/1/2, including the extra last original FLC frame), **6** repaint requests, **5** status messages, and every byte of the **51,200-byte** final ARGB image, SHA-256 `54f5ac9a770a0df9cc4679fb01c95d82552a666b1c44ed3ca059fefe4c507225`. Its run took **515,525 charged steps**. The owned test harness is a uniquely anchored, temporary adaptation of the SHA-pinned existing Moon Clouds headless JDK fixture, not a modification of either original class. JDK headless image creation and synchronous repaint dispatch are test adapters; this is **not** an X11 device or browser gate.

| Original JDK displayed state | SHA-256 of 38,400 RGB bytes | Independent Go FLIC decoder |
| --- | --- | --- |
| 0 | `c76fa3fcbd57f5890087e224d8bcd4df8be70a32d01dc0f29bc32696513f2452` | Same |
| 1 | `0a8ffeaf31d1d1c79677f3b3217356c858d9e4690ef1f2c2b929977c4cbb5a0a` | **Different** (9,895 of 12,800 pixels differ) |
| 2 | `60ed570bab36ac41dd5f4b252ff32a7493e297371d31645ffe34e62c4d671901` | **Different** (8,652 of 12,800 pixels differ) |

**Do not confuse original-program fidelity with codec correctness.** The [separate independently Go-checked FLIC CLI](flic-functional-slice.md) decodes all three `alex.flc` frames as FLIC; original 1999 Fliplay is not pixel-equivalent on later odd-width DeltaFLC frames. The second gate explicitly reproduces the *original JDK renderer's* frames in Magnolia instead of silently replacing them with the CLI's corrected output. The per-pixel divergence was calculated against the CLI's PNG RGB, whose three frame SHA-256 values were each separately matched to Go. No source/media bytes are written into this report.

**Fail closed:** a same-length changed original movie and a same-length changed original class each **exited 1** before original construction, resource access, or image/report output; neither mutation was passed to the old JVM code. `lib/applet-runtime.oak` charges a frame-sink capture against its **declared viewport** (default ≤81,920 pixels/frame; ≤61 frames, ≤5 million total captured pixels), rather than silently hard-coding 320×240. A separately pinned `num3.fli` caller now [opts in to a larger 307,200-pixel viewport](fliplay-num3-original-jdk.md) without changing this default. The changed bound was tested by repeating the *original Moon Clouds JVM capture* under a 460-second cap: **exit 0, 61/61 captured RGB frames byte-for-byte against separately executed JDK frames and the entire 307,200-byte final ARGB equal to JDK**, with the existing 122 repaints and 63 status messages. This does not recapture all 61 images on a physical X11 device; the prior [four-sample X11 device audit](fliplay-packed-native-x11.md) remains separate.

**Additional honest variants:** `num3.fli` has since [passed a separate **third** original-JDK/Magnolia gate](fliplay-num3-original-jdk.md): all **seven** 640×480 RGB display states and the final **1,228,800-byte ARGB** match the unchanged JDK, with a SHA-pinned, caller-only array/capture opt-in; the interpreter's **default** per-array cap remains 65,536 (the generated-Java target ceiling remains 131,072). Its separately regenerated JDK and recorded-native verifier exited 0; the preceding direct native per-frame run itself also exited 0. Conversely, original JDK `fliplay` **exited 1 before playback** on `2422.FLC` at `Invalid Frame Magic: 61696` (its prefix is not handled by this player), while the **separate** standalone FLIC decoder rendered **27/27** frames against independent Go. The rejected movie is **not** an original-Magnolia completion gate.

## Reproduce the bounded second-movie differential

```sh
cd /home/user/magnolia
# Use unchanged verified original classfiles, movie, a newly built Magnolia
# executable, and the separately rerun six-file Go differential report.
python3 fliplay-alex-oracle.py \
  --fliplay /path/to/original/fliplay.class \
  --flickframe /path/to/original/flickframe.class \
  --alex-flc /path/to/unchanged/alex.flc \
  --go-report flic-functional-slice.json \
  --magnolia /path/to/built/magnolia \
  --output /path/to/a/NEW/report.json
```

[Metadata-only machine report](fliplay-alex-original-jdk.json). **Limits:** three particular original movies now have separate finite JVM/JDK comparisons, but only Moon Clouds has four sampled physical X11 frames. Unknown media/palettes/corrupt files, unpinned image geometry, resize/expose, physical Win32 and timing remain open. This does not complete the Java Flash 2, Java ME, SWF5/AVM1, AS3/AVM2, SDK utilities, FLA, or other named applets and games.
