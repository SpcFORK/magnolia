# Original Acotel direct key-method differential — still no game startup

**Narrow original-program progress, not a working Java ME game:** the SHA-pinned, unchanged 50,594-byte Acotel JAR (SHA-256 `c24e61a416616a036b537957ec9b0687ad08f9fceae18e4d06336883ec8c3040`) now executes its actual `DS_Effects.keyPressed`, `keyReleased` and `keyRepeated` bytecode in Magnolia **after** its independently checked constructor. The [original-only verification driver](../midlet-acotel-checkpoints.oak) selects `--sequence acotel-direct-keys` using the [reusable bounded direct-callback plan](midlet-reusable-checkpoints.md) to invoke exactly three callbacks, on a still-**unselected** Canvas, in one shared metered state. They are **direct method calls, not MIDP device-delivered key events**. No original `startApp`, worker loop, `paint`, display selection, visible game frame, score or gameplay has run in Magnolia. The ordinary `start/full` routes still refuse **before** construction.

| Original code invoked directly | Original `_fld2h` action | Original `_fld3h` key-down flag |
| --- | ---: | --- |
| After original constructor, before direct calls | 0 | false |
| `keyPressed(50)` (numeric-key fallback) | 1 | true |
| `keyReleased(50)` | 1 | false |
| `keyRepeated(-4)` (pinned default-device right action) | 5 | true |

An **independent** `-Xverify:all` JDK 11 run uses a real MicroEmulator 2.0.4 (643,996 bytes; SHA-256 `dbd5f3eb8365d3e839d6a203149e0e3776fc1a0585e16ac1fc23f76c9fcae1c6`) and the **unchanged original JAR**. The owned reference proxy constructs the original game via reflection inside the *proxy's own* `startApp()`, then invokes those same original Canvas methods directly and observes their fields. It **never invokes the original game's `startApp()`**. Both hosts reproduce **all three actual original method state transitions** in the table. At the end, their 176×176 offscreen image remains uniformly opaque black (30,976 ARGB pixels; SHA-256 `582076772019e80f7184d6b85a32c546dbb60b88c72dbd4ca84effec1a6647f0`), Display unselected and worker absent. Magnolia charges **19,084 steps total** for constructor **and** three calls, versus 18,927 for the constructor-only route; this is **not a displayed pixel oracle**.

The [differential](../midlet-acotel-direct-key-oracle.py) pins both external JARs before launching either host, compiles only the owned observer, enforces separate process/step/output bounds, and emits [metadata](midlet-acotel-direct-keys.json) only after all checks. The original-only verification CLI pins the exact JAR, entry class, `a,b1,s0` resources and caller-supplied 31,119,560-byte **constant** virtual-memory snapshot. The separate reusable linker binds source/class bytes to a SHA-pinned local JAR and preflights a bounded Canvas callback plan; the runtime refuses mismatched host snapshot, paint/start profiles or unretained callbacks. No original class or game-specific key script is compiled into the VM. A same-size ZIP/CRC-valid altered JAR, wrong snapshot, incompatible paint flag, and default original lifecycle each exit **1 before construction**; they are **not** successes. The 18,930-step cap refuses **during the first direct callback after construction**, with no success record. A separate [low-level native smoke](../midlet-acotel-direct-key-smoke.oak) exercises wrong order, code, method name, and an extra call on the actual retained original Canvas: each refuses **before additional guest instructions or state changes**; attempts to omit the retained direct-key capability or turn the linked image into `startApp` refuse **before construction**. The three correct calls still complete with the same image/step count. This is a checked host boundary, not a claim of device event scheduling. A **single captured 29/29-test combined exit 0** covers five direct-key Python tests (including a fresh live differential and changed-JAR pre-child/no-report-overwrite negative), fifteen prior constructor/reference/service tests, and nine census tests. The full cross-media census separately exited 0 and **linked, did not re-execute**, this narrowly measured original-method checkpoint. A fresh `CGO_ENABLED=1 go build` (after registering 36 previously omitted opt-in Oak library modules in `lib.go`) produced a new executable; `go test ./...` exited 0, and the new executable independently reran **both** original Acotel differentials, the direct-key low-level smoke, and the full census with unchanged report hashes. Embedding the modules does not expand the selected game's allowed API or start its lifecycle.

Reproduce from the Magnolia root using separately acquired, hash-identical original/emulator JARs **outside** the checkout and an already built Magnolia executable:

```sh
timeout --signal=TERM --kill-after=2s 360s python3 \
  midlet-acotel-direct-key-oracle.py \
  --magnolia /path/to/magnolia \
  --jar /path/outside/ACO_XmasGifts_s.jar \
  --microemu /path/outside/microemulator-2.0.4.jar \
  --output midlet-acotel-direct-keys.json
python3 -m unittest -v test.midlet_acotel_direct_key_oracle_test \
  test.midlet_acotel_constructor_oracle_test \
  test.acotel_original_reference_test test.midlet_service_oracle_test \
  test.media_blocker_census_test
```

**Next acceptance gate:** a [separate original `startApp`-return checkpoint](midlet-acotel-start-return.md) now reaches its `Thread.start()` call, but **does not run the worker**. Run the original worker/paint/frame path with independently compared actual device pixels, bounded scheduling, resource/Font/glyph semantics, delivered keys/softkeys and gameplay/exit—not merely direct input-field changes in an unstarted object. The [constructor-only checkpoint](midlet-acotel-constructor.md) remains a separate first step. All **23 JVM + 25 ActionScript = 48 overlapping unchecked checklist entries** remain; independent remaining runtime-blocker count is still **unknown**.
