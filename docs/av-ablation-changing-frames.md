# Changing-frame rendering and scheduling ablations

This is a follow-up to the [initial audio/rendering/scheduling experiment](av-capability-ablation-results.md). It uses a different unchanged original movie, `GLASS.FLI` (SHA-256 `3aaca4f8541ba73890830687e5c436258caa7e1a4075156f23e249f0a9e6908c`), with the same pinned historical Fliplay and flickframe classes. A fresh JDK 11 reference and a fresh, independently built Magnolia control execute 20 movie frames plus the original ring. This cohort is sensitive to motion: the 21 JDK states contain **11 distinct RGB hashes**, with **1,227–4,363 pixels changed** between each consecutive pair. Both references match all 21 states, all 256,000 final ARGB bytes, 42 repaints and 23 status events.

| Fresh isolated variant | Observation on unchanged original Fliplay |
| --- | --- |
| Unmodified control | **Pass:** 21/21 JDK RGB states, full final ARGB and lifecycle agree; 9,494,742 guest steps. |
| Deny offscreen raster blit | Explicit denial reached after 1,052,681 guest steps; no completed matching capture or final raster artifact. |
| Deny bounded `Thread.start` | Explicit denial reached after 7,669 guest steps; no completed matching capture or final raster artifact. |
| Deny both | `Thread.start` denial is observed first, at 7,669 guest steps. The prior failure masks the blit boundary; this does **not** measure both effects independently. |
| Drop blits but falsely return Java success | The original caller continues until the independent frame check rejects **RGB state 0**; no final raster artifact. This is a deliberately dishonest, test-only no-op, not a safe or supported implementation. It bypasses normal validation and resource charging. |

Each changed copy has a pinned, exactly matched edit and its own rebuilt executable; neither working runtime source nor original inputs were changed. The identical-length movie mutation negative is rejected by the caller's hash gate before guest start and artifact creation. Explicit denials measure reached, named boundaries—not all drawing or all scheduling. The no-op outcome demonstrates that claiming successful `drawImage` does not ensure correct pixels; because the oracle stops at the first mismatch, it does **not** demonstrate what happens during the rest of that modified lifecycle. No physical screen or interaction is tested. The experiment's single measured times are diagnostics, not a performance or statistical conclusion.

## Actual audio mute, independent reference only

A fresh run of [`sdk-adpcm-real-audio-oracle.py`](../sdk-adpcm-real-audio-oracle.py) separately compares an unchanged original SDK C++-exact 4,098-sample ADPCM payload and independently decoded reference with Ruffle 0.7.1's bounded ALSA FIFO capture. The **same playable SWF** produced 16,384 non-silent stereo frames with volume 1, waveform correlation **0.999953**, and silence with volume 0; an asset-only (not started) control was also silent. Network, URL and guest filesystem access were denied, and capture was paced and capped below 900 KB. This is a real **Ruffle mixer mute control**, not proof of physical speaker output or Magnolia playback. It remains separate from the Magnolia WAV-encoding ablation, which cannot be promoted to an audio-output pass. The user-uploaded SWF5 was not tested by this ADPCM experiment.

## Reproduce and inspect

Supply byte-identical external original files; they are not bundled, and the output path must not already exist. JDK 11 is expected at `/usr/lib/jvm/jdk-11/bin`. The full native control may take several minutes and has a 1,200-second child cap.

```sh
python3 av-ablation-glass.py \
  --fliplay /var/tmp/magnolia-gemalogo-probe/classes/fliplay.class \
  --flickframe /var/tmp/magnolia-gemalogo-probe/classes/flickframe.class \
  --movie /var/tmp/magnolia-fliplay-new-originals/glass.fli \
  --output /var/tmp/av-ablation-glass-new.json

python3 sdk-adpcm-real-audio-oracle.py \
  --adpcm-source /var/tmp/magnolia-sdk-adpcm-reference/adpcm.cpp \
  --parser-source /var/tmp/magnolia-every-probe/sdk-swfparse.cpp \
  --ruffle /var/tmp/magnolia-flash-reference-v0.7.1/ruffle \
  --libs /var/tmp/magnolia-flash-reference-v0.7.1/xkb-lib/root/usr/lib/x86_64-linux-gnu \
  --output /var/tmp/sdk-adpcm-audio-mute-new.json
```

- [Changing-frame report](av-ablation-glass.json): SHA-256 `5ad38712addff2f64566b4487262bfc5071dccde9f797d8f81b930b9d1e0c400`.
- [Fresh external-audio report](sdk-adpcm-audio-mute-fresh.json): SHA-256 `a39300809daef068e1ba4c58b14d230acc5ec5c424f4526299da918adf66c143`. Capture length/hash may vary with timing; correlation and silence checks are the outcomes.
- Both runs exited 0. Python compilation, both JSON parses, `git diff --check` and `go test ./...` passed. The broad acceptance counts remain **21 JVM / 24 ActionScript unchecked**.

Further work: exercise audio output on an actual Magnolia-capable audible original once that baseline exists; independently test repaint/capture, input, clocks, resource and network boundaries; repeat treatments for performance measurements; and distinguish original SWF/MIDlet external references from full Magnolia outcomes.
