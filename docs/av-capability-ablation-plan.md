# Audio, rendering and scheduling: controlled ablation plan

## Question and safeguards

Which *specific implemented boundary* is necessary for a selected, unchanged program to reach an independently checked observable outcome? Disable one boundary per run on an isolated, instrumented checkout; do not edit originals or use altered builds as release binaries. Run an unmodified, freshly built control under the same input and budgets. A refusal, timeout, output mismatch, unexercised boundary and successful full outcome are different measurements. Missing implementation is not an ablation result, and suppression of an output is not proof that the underlying computation ran.

## Finite first experiment

1. Pin the unchanged original Fliplay class, helper and actual `GEMALOGO.FLI`; re-run the separate JDK 11 headless reference on those bytes. Compare all five displayed frames plus ring, the final 320×200 ARGB raster, and lifecycle counts. Retain its real resource, step and array caps. This file's displayed frames happen to be identical, so this is not a visual-motion test.
2. On separate one-change builds, explicitly refuse only (a) the offscreen `Graphics.drawImage` raster-blit boundary and (b) the finite JVM `Thread.start` scheduling boundary. Record whether original execution reaches each boundary and whether a final artifact exists. Do not silently replace rendering with black pixels or a successful thread start with a no-op. These switches are **not** all renderers or all schedulers.
3. On a third one-change build, explicitly disable the high-level PCM WAV writer only. Verify independently using Python `wave` on an owned short PCM fixture. Run the original silent Fliplay case on the audio-disabled build as a negative control; it should not depend on this WAV boundary. Run the same WAV probe under the rendering/scheduling variants to detect unintended cross-effects. This is encoding, **not physical sound output** or Flash audio.
4. Require exact source-edit anchors, same original hashes and limits, isolated builds with recorded binary hashes, bounded processes, explicit negative refusals and no output on refusal. Do not count a different capability's pre-existing failure as causal. Keep third-party originals and transient executable builds outside the checkout; persist only source, tests and sanitized metadata.
5. Report original-program results and owned fixtures separately, with raw observed statuses, control outcome, changed boundary, outputs, limits and weaknesses. The first controlled switch set does **not** establish whole-release audio, device graphics, true threads, Flash timeline/input, AVM2, MIDP gameplay or browser audio.

## Follow-up matrix (not completed by the first experiment)

- Run all relevant combinations after single-switch results, including interactions such as rendering plus scheduling. Capture latency and steps, but use repeated trials and variance before making performance or statistical claims.
- Isolate input delivery, local-resource access, timing/clocks, network/navigation, codecs, audio output and display-device presentation independently. Separate physical hosts from virtual guest calls and distinguish no-op, explicit denial and resource starvation.
- Repeat on different original cohorts: a changing-frame FLI/Fliplay input, the uploaded SWF5 against an independently pinned Flash player, and the original Acotel MIDlet against MicroEmulator. These last two currently lack the required Magnolia full baselines, so attempted removals cannot yet be interpreted as effects on functioning playback/gameplay.
- Re-run native and JS suites and validate original digests, pre-output refusals, time/memory/output budgets and independent pixel/audio/event references after any runtime change. Broader TODOs stay unchecked until each original reaches its actual requested outcome.
