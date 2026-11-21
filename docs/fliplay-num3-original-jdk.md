# Third original Fliplay movie: `num3.fli` at 640×480

**Observed, bounded original-bytecode result:** the unchanged `fliplay.class` and `flickframe.class` execute `num3.fli` in Magnolia's headless JVM from `init` through the actual original worker, original `load_fli`/palette/deltas, `repaint`/`update`, `next_frame`, `paint`, and finite pause. Its **7/7 displayed RGB states** (six movie frames and the ring display) match separately executed **JDK 11 `-Xverify:all`** original classes; all **1,228,800 bytes** of the final ARGB match the JDK file (`47339666281caeaaac52bd6a5e3b92e0cbd827fc05e849c808959a3ef7eaebeb`). The direct native command exited **0** in about **665 seconds**, with **16,917,685 charged VM steps**, 14 repaint requests and nine statuses. It captured **2,150,400 RGB pixels** from seven 640×480 original display states. No physical window, native timing, input, Flash or general media compatibility is inferred.

**Unchanged external originals:** `fliplay.class` (12,732 bytes; SHA-256 `bc315872dafebbf3c634e64c0fc251b38031ed49b7c4698d39c94990de9fe20b`), `flickframe.class` (830 bytes; `12feb5321cd8c9aeb0bf285f055596d981e3d45ccfb8c62d029f2f6d95ccdf9f`), and [`num3.fli`](https://samples.ffmpeg.org/fli-flc/num3.fli) (31,700 bytes; `5626298e5b3a71ef7fa830701116ec4be115dff3d94654d4d932601f0a77a6d2`). The movie has **FLC magic despite `.fli` extension**. Original bytes remain outside this repo; redistribution rights have not been established.

| Displayed state | JDK RGB SHA-256 of 921,600 bytes | Independent Go FLIC |
| ---: | --- | --- |
| 0 | `a6f7af72f8786a0ae4b544b7e11ad89ee95be643eac198b0ec1a7c0c74d4a01a` | Same |
| 1 | `8fd5543adca8054793196a7f567842f2c49c7827a533bce709245402e2d6e716` | Same |
| 2 | `04e7970301b2807c788fb6af49674c15153fbb9cee6ea97bef92a4c4c11f1b6a` | Same |
| 3 | `398a89b72c008d691d47455acb7ac5026a04828dbc0a693c3171d942f7496965` | Same |
| 4 | `8ce0e9761bf1a691cc472325bb6d99d9d4751a75f147cae4a856ffeefa83ac08` | Same |
| 5 | `965699ac9f7466d8917a6f1b652a52c9c3bb6ee6b9bdd9a2acb616d7be789616` | Same |
| 6 (ring) | `a6f7af72f8786a0ae4b544b7e11ad89ee95be643eac198b0ec1a7c0c74d4a01a` | Repeats 0 |

**Independent comparator order:** [`fliplay-num3-oracle.py`](../fliplay-num3-oracle.py) validates hashes/sizes and the separately passed [482-frame independent Go report](flic-functional-slice.json), adapts **only the pinned owned JDK headless test harness**, recompiles that harness against the *unchanged original* classes, runs JDK, and checks each fresh exported RGB file and final ARGB. The native runner [`fliplay-num3-original.oak`](../fliplay-num3-original.oak) checks the same originals **before** construction and checks each JDK RGB digest *inside the bounded frame sink*, not from the Python standalone player's frames. The first six also independently agree with Go; Go agreement is an extra codec cross-check, not a substitute for original JDK fidelity. A fresh Python verification **exited 0** using the previously captured **exit-0 native log/output**, regenerated JDK 7/7 and full final comparison, and separately checked two same-size changed movie/class negatives, both **exit 1 before original JVM construction or output**. Its [metadata-only report](fliplay-num3-original-jdk.json) labels this native evidence **recorded**, not a fresh second 665-second native run.

**Exact admission and limits:** no change to the interpreter's default **65,536-element array cap** (or the separate generated-Java target's 131,072-element admission ceiling). This SHA-pinned caller alone sets `maxArray: 307840` for the original movie's 640×(480+1) allocation. The separate frame-sink bound defaults to **81,920 viewport pixels per capture**; this caller explicitly sets `frameSinkMaxPixels: 307200` for 640×480. The runtime rejects an invalid/oversize cap, an actual captured viewport larger than its active array budget, viewport mismatch, >61 captures or >5,000,000 aggregate captured pixels; it charges RGB work before calling the sink. No other original movie acquires a larger allocation or capture budget by default. The finite run sets 40,000,000 charged steps, no post-pause control, and uses the headless cooperative adapter, **not** actual X11 or historical browser/OS timing.

Reproduce from the Magnolia root, providing **separately acquired and verified** originals and an existing compiled Magnolia native executable. Fresh native mode can take about 11 minutes; the Python wall cap is 780 seconds. The recorded-evidence mode is explicitly labelled and should only be used with a *separately captured exit-0* native run, not an interrupted log:

```sh
cd /home/user/magnolia
python3 fliplay-num3-oracle.py \
  --fliplay /path/to/original/fliplay.class \
  --flickframe /path/to/original/flickframe.class \
  --num3 /path/to/unchanged/num3.fli \
  --magnolia /path/to/compiled/magnolia \
  --output /path/to/a/NEW/report.json
# OR add --observed-log /path/to/finished-native.log \
#        --observed-final /path/to/finished-native.argb
```

**Post-change regressions:** `alex.flc` again **exited 0** at 3/3 original-JDK states and the same final image; an additional run with the interpreter's **65,536-cell default array budget** and unchanged **81,920-pixel default capture ceiling** also exited 0, showing that a 12,800-pixel viewport need not increase its array budget just because the capture *ceiling* is higher. An invalid `307201`-pixel capture ceiling **exited 1 before any sink callback or ARGB output**. A newer full Moon Clouds child reached its terminal paused **61-capture** log, and a *separate completed comparison* byte-matched **all 61** exported RGB files plus final JDK ARGB; the **user-interrupted parallel shell did not capture that child wrapper's exit status**, so this latest attempt is **not** reported as a fresh shell exit 0 (an earlier independent full gate did capture exit 0). **22/22** selected media/census unit tests passed. [Explicit regression metadata and exit-status caveat](fliplay-three-movie-regression.json).

**Boundary:** this closes one explicitly opted-in **third** original Fliplay *movie/headless profile*, not generic 640×480 arrays, whole AWT hosting, physical X11 (only [Moon Clouds has four sampled device screens](fliplay-packed-native-x11.md)), a general FLI/FLC player, or any ActionScript/Java Flash game. `alex.flc` has a [separate original-JDK gate](fliplay-alex-original-jdk.md), with later original pixels that disagree with the independent Go standard-oriented decoder. Original JDK `fliplay` still rejects `2422.FLC`'s prefix before playback, even though the separate standalone utility decodes it.
