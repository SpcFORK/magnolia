# Frame capture versus repaint: original-workload ablation

This experiment separates an **optional frame-observation sink**, a guest's **repaint request**, and the runtime's **repaint-drain hook**. It uses unchanged, SHA-pinned historical Fliplay and flickframe classes and original `CITY.FLI`, against a fresh JDK 11 headless reference. The JDK produces four states (three distinct RGB hashes), eight repaint requests, six statuses and a final 320×200 ARGB screen.

The two capture runs use the **same freshly built executable and same test caller**. Their only intentional difference is that the caller either supplies its validating sink or passes no sink. The caller is a tightly anchored, documented adaptation of the pinned existing CITY driver: it reports capture counts, accepts zero observations only when explicitly requested and still requires the same final raster and lifecycle counts. It does not change guest bytecode or movie bytes. For repaint interventions, exactly one runtime site is altered in a separate temporary build.

| Variant | Observed result |
| --- | --- |
| Capture on | **Pass:** 4/4 JDK RGB states, all 256,000 final ARGB bytes, eight repaints, six statuses, one started task; 256,000 captured cells. |
| Capture off | **Pass:** identical final JDK ARGB bytes and the same reported lifecycle counts, with zero captured cells or intermediate state comparisons. |
| Explicitly deny bounded repaint | Reached the named denial after 1,156,885 metered steps; no final artifact. |
| Suppress repaint while returning normal void | A deliberately dishonest **test-only** no-op; the run later refused `Thread.sleep outside bounded virtual-clock task` after 3,101,097 metered steps, before completing the lifecycle or writing a final artifact. |
| Explicitly deny repaint-drain hook | First JDK frame checked, then the named denial was reached after 2,153,097 metered steps; no final artifact. The hook intervention is unconditional even if a queue might be empty. |

Removing observation is **not** removing rendering: the sink-off run has no evidence for intermediate frame parity. Its final-frame parity does not establish that all unseen frames matched. It used 2,399,406 metered steps versus 3,167,406 with capture; the 768,000 difference equals three charged steps per captured cell in this specific run. Those are runtime budget charges, **not** a statistically valid speed benchmark. The repaint no-op is not a supported implementation and does not prove what would happen with every possible missing-event or scheduler policy.

The same-length edited original FLI was refused before guest execution or final output. A first experimental attempt expected the no-op repaint to reach a completed lifecycle and therefore failed its test; the corrected run above instead asserted the actual bounded-clock refusal. Neither the initial attempt nor the temporary experimental builds are release functionality. The broader 21 JVM / 24 ActionScript acceptance items remain unchecked.

## Reproduction

External originals are not bundled. The input digests and the owned CITY driver are pinned by the runner; JDK 11 is required at `/usr/lib/jvm/jdk-11/bin`. Each child has a 900-second wall cap. The output path must be new.

```sh
python3 av-ablation-capture-repaint.py \
  --fliplay /var/tmp/magnolia-gemalogo-probe/classes/fliplay.class \
  --flickframe /var/tmp/magnolia-gemalogo-probe/classes/flickframe.class \
  --movie /var/tmp/magnolia-fliplay-new-originals/city.fli \
  --output /var/tmp/av-ablation-capture-repaint-new.json
```

[Machine-readable report](av-ablation-capture-repaint.json), including observed log tails, source/executable hashes and scope: SHA-256 `19dead8cd534a05d4eb138d47c345432383972ed75eacd51309bafd71d8f3024`. The corrected experiment exited 0; Python compilation, JSON parsing, `git diff --check` and `go test ./...` passed. No physical screen, user input or audio output is exercised.
