# Reusable MIDP checkpoints, with original-specific results kept in oracles

The opt-in, bounded [MIDlet host](../midlet.oak) no longer has Acotel-specific execution sequences, class names, pixel hashes, key codes or JAR hashes. The original-only `acotel-construct`, `acotel-direct-keys` and `acotel-start-scheduled` names now exist **only in the verification driver** [`midlet-acotel-checkpoints.oak`](../midlet-acotel-checkpoints.oak). Its [shared test profile](../midlet-acotel-verify.oak) pins the unchanged original JAR, entry, resources, measured **constant** free-memory snapshot and expected original fields/steps. The native compiler requires an opt-in method image to carry a caller-supplied local JAR SHA/size pin **and exactly the class bytes in that JAR**; runtime gates independently refuse wrong host capabilities, profile combinations, budgets and retained deferred-worker continuations. Ordinary MIDP and Java SE Applet profiles remain separate.

Three separate gates on **two different, owned two-class MIDlets** independently exercise the reusable paths against SHA-pinned MicroEmulator 2.0.4. Neither fixture substitutes behavior for the game. The explicit default-device geometry is centralized in [`midp-device-profile.oak`](../lib/midp-device-profile.oak) and independently checked against the owned guest `Displayable` model; it is **not universal handset geometry**:

| Reusable gate | Same checked classfiles in Magnolia and real MicroEmulator | Magnolia result | What it does **not** establish |
| --- | --- | --- | --- |
| [`construct-only`](../midlet-constructor-generic-oracle.py), [metadata](midlet-constructor-generic.json) | `DirectSliceMIDlet`'s original constructor leaves its Canvas unselected in both hosts | **46 charged steps**, no retained session; its `startApp`, key callbacks and `paint` are unlinked | Any target `startApp`, worker, direct/device key, paint, device frame or gameplay |
| [`start-deferred`](../midlet-deferred-generic-oracle.py), [metadata](midlet-deferred-generic.json) | `DeferredSliceMIDlet` calls its real `startApp`; real emulator later runs `View.run` and selects the Canvas | Returns at **75 charged steps** after one logical `Thread.start`; `run`/`paint` unlinked, no retained session | Worker execution, actual selected device, frame parity or gameplay in Magnolia; no scheduling order is assumed |
| [Direct Canvas callback plan](../midlet-direct-generic-oracle.py), [metadata](midlet-direct-generic.json) | `DirectSliceMIDlet` constructs an unselected Canvas; both hosts invoke `keyPressed(-5)`, `keyReleased(-5)`, `keyRepeated(49)` directly | Same input-field trace, **46** constructor/**139** final charged steps, three bounded methods on one object | **No device-delivered key event**, `startApp`, worker, paint, selected screen, visible frame or gameplay |

The generic constructor CLI requires an explicit SHA-pinned local JAR whose class inputs match its member bytes. The host binds the caller-supplied virtual-memory snapshot, 1–50,000-step/depth/array limits and no retained lifecycle/Canvas callback. A changed-but-ZIP-valid class JAR, wrong pin, paint combination or over-budget partial construction is **not a completed constructor**. Negative guest fixtures that explicitly call `startApp`, `run` or Canvas `paint` from their constructors now refuse at link time, before their constructor's `BAD:` output can run; the deferred profile likewise refuses a linked `run` or Canvas `paint` body. Separately, `Thread.start` in the deferred profile is admitted **only during the one host-selected `startApp` call**: an adversarial constructor-started worker or early/recursive `startApp` refuses rather than masquerading as a successful start-return boundary. The unchanged Acotel original also completed via this reusable CLI in **18,927 steps**. Its generic record reports **zero commands on the logically requested display because the constructor did not request a display**; the original `Fine` command registered on its **unselected Canvas** is checked separately by the pinned [field-specific verification oracle](midlet-acotel-constructor.md), not erased or inferred absent.

The direct-key compiler prelinks only the explicitly ordered, at-most-three `(I)V` Canvas methods. It rejects a missing/wrong source pin, changed or mismatched JAR class bytes, paint/deferred-start combinations, a non-constructor image, invalid/nonconcrete methods and four-callback plans **before guest execution**. The host requires the matching free-memory snapshot, a 1–50,000-step shared budget, depth/array limits, an exact retained constructor capability and an unselected owned Canvas. Wrong callback order/code, a fourth call and forged host settings refuse without advancing guest steps; an insufficient budget **during** a callback is not a completed success. These are direct method calls, **not** handset event dispatch.

The three original Acotel differentials still pass *separately* under the unchanged original-JAR pin: [constructor](midlet-acotel-constructor.md) **18,927** steps, [direct keys](midlet-acotel-direct-keys.md) **19,084** and [`startApp` return](midlet-acotel-start-return.md) **18,957**. The latter does **not** run the original worker even though the real emulator does. The [cross-media census](media-blockers-current.json) **links all three owned and all three original reports** without rerunning the expensive differentials. Its **23 JVM + 25 ActionScript** unchecked checklist entries overlap; the independent remaining-runtime-blocker total remains **unknown**.

Reproduce with a built Magnolia executable and a separately obtained, exactly pinned MicroEmulator JAR outside this checkout:

```sh
python3 midlet-constructor-generic-oracle.py \
  --magnolia /path/to/magnolia --microemu /path/outside/microemulator-2.0.4.jar \
  --output midlet-constructor-generic.json
python3 midlet-deferred-generic-oracle.py \
  --magnolia /path/to/magnolia --microemu /path/outside/microemulator-2.0.4.jar \
  --output midlet-deferred-generic.json
python3 midlet-direct-generic-oracle.py \
  --magnolia /path/to/magnolia --microemu /path/outside/microemulator-2.0.4.jar \
  --output midlet-direct-generic.json
python3 -m unittest -v test.midlet_constructor_generic_oracle_test \
  test.midlet_deferred_generic_oracle_test \
  test.midlet_direct_generic_oracle_test test.media_blocker_census_test
```

**Next genuine acceptance:** independently model and test the original worker's bounded scheduling, `Canvas.repaint/serviceRepaints`, rendering/Font/glyphs, an actual device-frame comparison and delivered keypad/softkey interaction. Neither constructor-only execution, direct key-field transitions nor `startApp()` returning proves those behaviors or a playable game.
