# Bounded `InputStream.read(byte[])` — an independently checked Java media slice

**Historical bulk-read checkpoint, superseded as the latest player gate.** The [five later checked member families](jvm-five-blocker-preflight.md) first moved the selected original preflight to legacy `jsr` byte 355 and **237/277, 52** static entries; [six more bounded gates](jvm-next-blocker-preflight.md) later moved the selected refusal to `Object.wait()V` in `Semaphore.Ƅ()V` byte 4, with **245/277, 44** static; [five later archive-derived families](jvm-another-flash-blockers.md) left that refusal and moved **then-current** static coverage to **256/277, 33**. The [single bounded wait handoff](jvm-wait-handoff-preflight.md) subsequently moved selected preflight to `Sound.stop()`/`sun.audio` before construction and **then-current static coverage to 258/277, 31**. A [later passive Class probe](jvm-image-decoder-class-preflight.md) moved static coverage to **259/277, 30** at that checkpoint; the [subsequent bounded branching-finally fixture](jvm-branching-finally-preflight.md) moves current **static-only** coverage to **260/277, 29** without changing the selected refusal. There is still no original Flash construction or playback. The measurements below remain reproducible state **at this bulk-read checkpoint**, not the current state.

**Original-player gate at this checkpoint:** the unchanged, SHA-pinned `FlashApi.ZIP`
(SHA-256 `fbb2149179c1c5a32bd9aa7f479a1e760e36e18dbd8e48147729ca67822932c2`)
still **refuses selected-entry preflight before constructing the Applet**:

```text
[applet] com/macromedia/flash/Flash.run()V at byte 279: unsupported Java API java/lang/System.currentTimeMillis()J
```

The separate **all-method static** decoder inventory at this checkpoint was **233/277
bodies, 56 first-blocker entries**, versus **232/277, 57** at the earlier
[peerless-Graphics checkpoint](jvm-peerless-graphics-preflight.md). The added
member removes `ScriptPlayer.run()V`'s *static* byte-11
`java/io/InputStream.read([B)I` blocker; it does **not** run `ScriptPlayer`,
advance the selected `Flash` entry, measure remaining independent bugs, or
play a movie. `Flash.run()` uses the refused clock value to construct a frame
deadline at bytes 284–295. A controlled clock and scheduler contract is
needed before admitting that call; a fabricated timestamp is not a fix.

## Exact admitted contract

Only the **symbolic** `java/io/InputStream.read([B)I` with
`invokevirtual` is admitted in `lib/jvm-vm-libs.oak`. At execution its receiver
must be a bounded `ByteArrayInputStream` backed by a real JVM `[B`, or the
opaque `InputStream` returned by `URL.openStream()` for a *base-relative,
explicitly supplied local resource*. The VM opens no real path or socket for
applet code. It does **not** construct `FileInputStream`, replace an arbitrary
subclass's override, grant `InputStream.read([BII)I`, or admit the different
symbolic owner `ByteArrayInputStream.read([B)I`. A `DataInputStream` receiver
through this one-array member deliberately refuses even though that call is
JDK-valid; the separately admitted `DataInputStream.read([BII)I` remains
available for bounded sources.

Null destinations raise a catchable `NullPointerException`; a closed resource
raises a catchable `IOException`. The method preserves untouched target cells
and uses source-specific EOF rules. Before copying **any** byte, it charges the
entire moved count to the finite VM step budget; a short budget leaves the
byte array, stream cursor and work counter unchanged. A bounded snapshot of
the source bytes before stores is necessary: `ByteArrayInputStream` keeps the
original source array by reference, and that array can also be the destination.
The shared snapshot also corrects overlapping copies through the pre-existing
`DataInputStream.read([BII)I` host path. No asynchronous/remote stream,
filesystem access or arbitrary Java I/O is inferred.

Only caller-supplied `options.resources` (code base) and
`options.documentResources` (document base) are eligible: at most **four**
single-component ASCII file names in total, **524,288 bytes per file**, and
**1,048,576 bytes total**. Values can be native raw byte strings or lists of
unsigned `0..255` integers. The list representation avoids silent corruption
when non-UTF-8 binary data is embedded in Builder JS/HTML; the Builder now
converts high-bit raw resource bytes to bounded lists *before* emitting
JS/HTML, and rejects invalid lists/names/sizes without an artifact. The
runtime validates and materializes those lists before Applet construction;
JS host overrides containing non-byte Unicode code units are refused rather
than silently truncated to low-byte values.

## Reference and execution evidence

The **owned** Java 8 fixture
[`BulkStreamApplet.java`](../test/fixtures/BulkStreamApplet.java) compiles
reproducibly to `BulkStreamApplet.class` (SHA-256
`946302eaf0976a0bc96dfba596e108d498ae7eaec918f027d0c12b4fc498b5d8`).
Its [JDK harness](../test/fixtures/BulkStreamJdk.java) uses a real
`AppletStub` with `document/page.html` and a separate `assets/` code directory
under JDK 11 `-Xverify:all` on Xvfb24. **Only the JDK reference** opens real
`file:` URLs. The local binary fixtures are document bytes `ff 00 80` and
code bytes `43 32` (`C2`); Magnolia receives those exact bytes via two
explicit in-memory maps. The independent output is pinned in
[`BulkStreamApplet.expected`](../test/fixtures/BulkStreamApplet.expected).

| JDK observation | Bounded Magnolia result |
| --- | --- |
| File-backed reads of a zero-length array return **0 both before and at EOF**; nonempty reads return **3** document bytes or **2** code bytes, then **-1**. Signed values `-1, 0, -128` and unwritten sentinels are retained; `read(null)` is caught as `NullPointerException`. Closing then reading throws `IOException`. | Native Oak, standalone Oak, Builder JS/Node and offline HTML match the whole `probe` trace; native and Builder JS also match the separately selected closed-file method. The document/code bases remain distinct. |
| `ByteArrayInputStream.read(new byte[0])` returns **0 before EOF** but **-1 at EOF**. Later source-array mutation is visible to the next read. With source and target the *same* `[1,2,3,4]` array after one byte consumed, a bulk read returns **3** and leaves `[2,3,4,4]`. | Native and Builder JS match sharing, signed bytes, both EOF rules and the overlap trace. The shared host helper snapshots before aliasing stores. |
| A `DataInputStream` wrapping the file reads **3**, then its pre-existing `read(byte[],0,0)` returns **0 at EOF**. | Native and Builder JS match. This wrapper route is distinct from admitting one-array calls on all `DataInputStream` receivers. |
| A 256-byte read succeeds on the JDK; wrong symbolic owner and the three-argument overload are JDK-valid. | With a 160-step limit, native and Builder JS refuse the 256-byte read before fixture output. A direct native helper check confirms no partial array/cursor/work mutation. The unsupported owner and overload refuse on native and **all four Builder targets** before artifacts; a one-array `DataInputStream` receiver refuses at runtime without output. |

The [focused reproducible oracle](../applet-bulk-stream-oracle.sh) **exited
0** for the JDK, native VM, standalone Oak, Builder JS and offline HTML, plus
these negative gates. Builder-generated **Java** deliberately refuses the new
reachable one-array member **before writing source**; it does not claim
resource/bulk-read parity. The prior
[`DataInputStream.read([BII)I` oracle](../applet-byte-stream-oracle.sh)
was rerun and **exited 0** on its JDK/native/Oak/JS/HTML/Java routes after the
shared-helper change. The six-oracle
[Flash preflight regression](../applet-flash-preflight-regression.sh)
and pinned [cross-media census](media-blocker-census.md) **also exited 0**,
with eight independent original selected refusals still at byte 279.

Reproduce the focused JVM gate from the repository root with JDK 11, Xvfb,
Node and a built Magnolia executable:

```sh
MAGNOLIA=/path/to/magnolia bash applet-bulk-stream-oracle.sh
# Optional: separately supply the size/SHA-pinned original Flash archive.
MAGNOLIA=/path/to/magnolia FLASH_ZIP=/path/to/FlashApi.ZIP \
  bash applet-flash-preflight-regression.sh
```

This is a **stream member differential on owned bytecode**, plus a separate
original *preflight-only* measurement. It is not a Flash frame, an AWT peer,
a timestamp, a scheduler, audio, or playback.
