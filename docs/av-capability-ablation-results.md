# Controlled host-boundary ablation: results

The [plan](av-capability-ablation-plan.md) separates runtime boundary denial, API admission and scheduling. The completed finite experiment uses unchanged original `fliplay.class`, `flickframe.class` and `GEMALOGO.FLI`; a fresh independent JDK 11 headless reference; an owned PCM WAV fixture; and separately built, temporary one-intervention copies. The runner does not change the original files or working runtime source. A fresh control uses the same step, array and frame caps as the interventions.

| Variant | Original Fliplay outcome | Owned PCM WAV fixture |
| --- | --- | --- |
| Unmodified control | **Pass:** all six display states, final ARGB and lifecycle counts match fresh JDK. | **Pass:** Python `wave` reads 8-bit mono `[0,128,255]`. |
| Deny offscreen `Graphics.drawImage` raster-blit | **Explicit refusal reached**, exit 1; zero completed state checks and no final raster artifact. | Pass, unchanged bytes. |
| Deny bounded JVM `Thread.start` | **Explicit refusal reached**, exit 1; zero completed state checks and no final raster artifact. | Pass, unchanged bytes. |
| Disable high-level WAV encoder | **Pass:** same six states, final ARGB and lifecycle counts as control. This is a silent movie and does not exercise audio playback. | Explicit unavailable refusal, exit 2; no WAV artifact. |
| Omit local movie resource from caller | Exit 1 before any completed state checks or final raster artifact; fails at existing `PrintStream.print(Object) cannot model this toString` while handling the missing resource. **Not** a clean resource-denial diagnostic. | Pass, unchanged bytes. |

A same-length edited FLI is also refused by the original input-digest gate before guest execution or output. The two explicit denials demonstrate that the unchanged original workload reached those *specific* boundaries: blocking them prevents its successful finite lifecycle. They do not show the effects of silently dropping frames or suppressing thread execution while returning success, nor do they establish behavior of other rendering or scheduling paths. The six movie states happen to have identical hashes, so they cannot measure animation sensitivity.

**Audio limitation:** WAV *encoding* is an exercised audio-data boundary, not a speaker/output sink. The control confirms that the tested silent original does not depend on that encoder; it says nothing about audible movies. Magnolia does not currently have a demonstrated full SWF audio-playback baseline in this cohort, so a valid audio-device or SWF-sound ablation remains **unavailable**, not passed. Existing Ruffle ADPCM playback is a separately owned external reference and cannot be attributed to Magnolia. Interactions, physical display, input, network and wider applet/SWF/MIDlet playback remain future work; 21 JVM and 24 ActionScript acceptance items remain unchecked.

## Reproduce

The external originals are not distributed in the checkout. Provide byte-identical originals; the runner checks pinned sizes and SHA-256 before execution and refuses an existing report path. JDK 11 must be at `/usr/lib/jvm/jdk-11/bin`, with Go and Python 3 available. Builds and JDK reference run in temporary trees under `/var/tmp`, which are removed after completion. The previous GEMALOGO run used:

```sh
python3 av-capability-ablation.py \
  --fliplay /var/tmp/magnolia-gemalogo-probe/classes/fliplay.class \
  --flickframe /var/tmp/magnolia-gemalogo-probe/classes/flickframe.class \
  --movie /var/tmp/magnolia-fliplay-new-originals/gemalogo.fli \
  --output /var/tmp/av-capability-ablation-new.json
```

A full run can take minutes; each child has an explicit time cap. The checked [machine-readable report](av-capability-ablation.json) records source and executable hashes, individual outcomes and limits; SHA-256 `f61f300d5dd33416c703360b6cd702dc15bf132edb43578293de9e77bb42b31e`. The run exited 0 after all controls. Single-run elapsed times are diagnostic only and are not performance claims. During development, the initial WAV fixture had an Oak syntax error and an earlier completed attempt expected the wrong missing-resource diagnostic; neither is counted as a pass. The corrected run recorded the actual pre-existing failure instead.
