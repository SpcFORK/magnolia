# One bounded JVM `Object.wait()` / `notifyAll()` handoff — not general concurrency

**Historical wait checkpoint.** Its **258/277, 31** inventory below and [ten-oracle log](jvm-flash-ten-oracle.log) are saved evidence from before the [later, separately checked passive ImageDecoder Class probe](jvm-image-decoder-class-preflight.md). After a [further owned branching-finally fixture](jvm-branching-finally-preflight.md), current original static coverage is **260/277, 29**; the selected original `Sound.stop()` refusal before construction is unchanged. Do not read the historical counts/log as a new eleven-oracle run.

**Result (2026-10-26).** The unchanged **85,800-byte original** `FlashApi.ZIP` (SHA-256 `fbb2149179c1c5a32bd9aa7f479a1e760e36e18dbd8e48147729ca67822932c2`) formerly stopped during selected-entry preflight at `Semaphore.Ƅ()V` byte 4, `java/lang/Object.wait()V`. Magnolia now admits the **exact zero-argument `invokevirtual Object.wait()V`** on **one owned monitor, one running waiter and one pre-queued notifier**. The notifier runs only *after* the waiter releases **all** reentrant holds on that same monitor; it must actually own that monitor and call `notifyAll()`. After the notifier exits and releases it, the waiter reacquires its original depth before proceeding. This is a single, finite, interpreter-controlled schedule, **not** an OS/JVM thread scheduler or Java Memory Model.

The updated **selected original `com/macromedia/flash/Flash` entry still refuses before its Applet constructor**, with or without explicit virtual-clock/version options:

```text
[applet] com/macromedia/flash/Sound.stop()V at byte 36: unsupported field sun/audio/AudioPlayer.player
```

The **separate [all-method static inventory](jvm-flash-after-wait-static-coverage.json)** was **258/277 decoded bodies and 31 static first-blocker entries** at this wait checkpoint, up from the previous **256/277 and 33**. The saved raw inventory JSON SHA-256 is `28b4e55e28dd24283105ebefacd5b2d8e1119eae307831285db95ac2144ca1af`. A fresh CLI rerun returned identical metadata and the same **31-entry `(member, message)` multiset**, but its *raw JSON byte order differed* because iteration order is not pinned; do not use the saved raw-file SHA as a claim of byte-identical reruns. Only the two `Semaphore` wait-bearing method bodies newly decode. A third original `Bitmap` body's *first static message* moves from `Object.wait()V` byte 9 to **restricted legacy `jsr`/`ret` byte 80** and **remains blocked**. No original `Semaphore` wait ran, and the static counts are **not original execution or 31 independent remaining bugs**. `sun.audio.AudioPlayer` is unavailable in the JDK 11 reference here. The original `Sound.stop()` bytecode also invokes its `stop(InputStream)` overload, while other original `Sound` methods invoke `start(InputStream)`; simply supplying a fake static `player` field would neither match a historical audio reference nor establish those methods. No fake audio backend was admitted to make preflight advance. There was **no original Flash Applet construction, audio, frame, movie, image peer or playback**.

## Exact owned-fixture contract and JDK reference

The JDK 11 `-Xverify:all` reference executes **the same pinned Java 8 class bytes** as native Magnolia, a separately built Builder JS, standalone Oak/JS, and offline HTML. The owned `FlashWaitHandoffMain` fixture starts its worker; that worker enters one `synchronized (LOCK)` *inside another*, starts a single notifier while still owning the lock, and calls `LOCK.wait()` in a guarded loop. On the real JDK, the new notifier cannot acquire `LOCK` until that call releases it. It sets `phase = 7`, prints `notifier`, calls `LOCK.notifyAll()` and exits the monitor. The waiter resumes, calls `notifyAll()` again to verify it **reacquired ownership**, then prints `after` and `7`. All positive routes agree byte-for-byte:

```text
before
notifier
after
7
```

The VM inspection also checks **one completed wait, two owned notifications, two started tasks**, no pending thread, no retained waiter/task identity and no held monitor after return. Already-supported `notifyAll` ownership, implicit synchronized methods and ordinary single-task Fliplay behavior are *not* reinterpreted as general multi-task execution. `Thread.start()` from inside a running task queues a notifier **only while that task owns a monitor**, with at most one pending notifier; an otherwise ordinary first task still runs synchronously. The worker and notifier share the existing instruction/depth/reference caps; there is no host/native concurrency or background event queue.

**Negative and exception evidence:**

- JDK and Magnolia both catch the real unowned `IllegalMonitorStateException` and null-receiver `NullPointerException` from `Object.wait()`, printing `unowned` and `null` respectively. The bytecode really contains those invocations; their checks do **not** create a wait set.
- The JDK reference verifies that a missing notifier, a notification on the **wrong monitor**, and a notification **before** the wait do **not** unblock it (the *reference harness only* gives the worker a 50-ms join, verifies it has not terminated, then interrupts and joins it). Magnolia instead **refuses** each indefinite/unnotified pattern: it does not silently return or spin forever. One JDK-valid second pending notifier and a nested `Thread.start()` without an owned monitor also refuse before emitting stdout.
- `Object.wait(J)V` succeeds on the JDK with a 1-ms timeout; Magnolia deliberately **rejects the descriptor at preflight** on Oak, JS, HTML and generated Java. `Object.notify()` and wrong `wait()` invocation opcodes also refuse. Generated-Java output explicitly refuses even the newly admitted zero-argument member **before writing a `.java` artifact**; it does not substitute its host JVM's scheduler.
- A short VM-step budget refuses without partial stdout. A notifier that leaves an uncaught exception or leaks its monitor fails closed; the waiter's original lock depth is restored for ordered cleanup, not falsely reported as a successful wakeup.

**Not covered:** arbitrary runnable interleavings, multiple waiters or notifier queues, real-time or timed waits, spurious wakeups, JVM interruption semantics, wakeup fairness, cross-lock contention, starvation or deadlocks, a real Java memory model, general thread joins/daemon priorities, a browser AWT event queue, original Flash wait execution or the historical audio backend. The `Semaphore` shape passing *preflight* is not proof of its runtime behavior in the original player. In particular, an original unnotified wait still fails at runtime outside this handoff.

## Pins, reproduction and independent gates

- Owned sources: `test/fixtures/FlashWaitHandoff{Main,Jdk}.java`. The runner's `join()` and **negative-test-only** interrupts execute **only on JDK**, not in Magnolia. Recompile with `javac --release 8 -g:none -proc:none`; SHA-256: `FlashWaitHandoffMain.class` `2cf46953401cba9be5fedf87b4fc0ff5480c600834c922ecae8cc17bdd4eaa21`; `FlashWaitHandoffJdk.class` `c3baa1d9e2cdac3095ee90881baa6e099cdb1f706d95568c3f0963e48d09895b`.
- [`applet-wait-handoff-oracle.sh`](../applet-wait-handoff-oracle.sh) **exited 0** ([saved focused PASS log](jvm-wait-handoff-oracle.log)): class pins and `javap` exact opcodes/descriptors, JDK/native positive and caught-exception stdout, native/Builder-JS state/refusals, three selected Oak/JS/offline-HTML methods, generated-Java no-artifact refusal, out-of-profile JDK controls, and two SHA-pinned original selected attempts plus **258/277, 31** static counts. The fixture and original archive tests are **separate**.
- The earlier [six-gate oracle](../applet-next-blocker-oracle.sh) and [five-family oracle](../applet-another-flash-oracle.sh) independently **exited 0 again** after this change. The [explicit monitor matrix](../applet-monitor-matrix.sh) also **exited 0**. The cross-media census **exited 0** and its six focused bounds tests passed. These targeted results do **not** claim a broad `test/main.oak applet` suite run; the combined focused Flash preflight has its own independently recorded result below.

```sh
MAGNOLIA=/path/to/magnolia \
FLASH_ZIP=/path/to/separately-obtained/SHA-pinned/FlashApi.ZIP \
  bash applet-wait-handoff-oracle.sh
python3 media-blocker-census.py \
  --magnolia /path/to/magnolia \
  --acotel /path/to/original/ACO_XmasGifts_s.jar \
  --cache /var/tmp/magnolia-media-census \
  --output media-blockers-current.json
```

The **expanded ten-oracle** [`applet-flash-preflight-regression.sh`](../applet-flash-preflight-regression.sh) subsequently **exited 0 independently** ([saved terminal log](jvm-flash-ten-oracle.log), SHA-256 `a9834ab31edd2949882bd1ccf96f848bb17d89b312c80e76e332baadc60a3f9d`). It includes all ten bounded owned-fixture gates and **eight** SHA-pinned selected-original preflight attempts—four default and four with explicit virtual clock/version—all refusing *before construction* at `Sound.stop()`/`sun.audio`, followed by the **258/277, 31** separate static inventory. Its own final line records `COMBINED_REGRESSION_EXIT=0`; this is **not** an original Flash wait, movie or audio run. A historical Flash 2-compatible reference movie and independently observed original frames/input/audio/teardown remain required before any whole-player claim. Third-party original `.class` bytes are **not committed**.
