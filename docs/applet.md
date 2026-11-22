# Java applet files in Magnolia: Builder integration, Oak first

Magnolia Builder accepts **verified subsets** of JVM `.class`, `.jar`/`.zip`, local applet launch packages **and `.java` source**. The implementation is in Magnolia/Oak libraries, not a Go applet converter. It emits **self-contained, runnable Oak source** as the primary target; JavaScript/HTML are bundled through Magnolia Builder's Oak→JS pipeline; standalone Java *source* is another supported output. A reusable [Oak VM-library transition layer](jvm-vm-libraries.md) connects the input frontends, linker, host-library contracts and bounded interpreter. `.java` input may use either isolated **build-time `javac`** or an explicitly smaller **Oak-only Java subset**; emitted Oak/JS/HTML needs neither. **Original-byte restoration** and **edited Oak → new `.class` or unsigned `.jar` recompilation** are distinct paths: the latter accepts only checked, same-layout bytecode edits, not arbitrary Oak compilation.

**Compatibility is intentionally bounded.** This is not a full JVM, general Java compiler, AWT/Swing implementation, or universal applet runner. An unsupported reachable instruction, exception handler, or Java API produces a specific error instead of a fake execution result. No external JVM or browser-side Java runtime is required for Oak/JS/HTML output; Oak requires Magnolia to run, while emitted Java source requires a JDK. For the mediated Transpile/Eval/Build/Compile route contract, quantified whole-class audit and staged full-coverage gates, see the [JVM compatibility plan](applet.md).

Run examples from the Magnolia repository root with a rebuilt `magnolia` executable. [Runnable outputs](../samples/applet/README.md) include an interface-heavy class/JAR example across standalone Oak, JS/HTML, Java source and a changed executable JAR. An installed binary running elsewhere can find the adapter by walking upward from an input inside the repository, by `--applet-adapter /path/to/magnolia/lib/build-applet.oak`, or by setting `MAGNOLIA_APPLET_ADAPTER` to that path. Its companion Oak libraries must be present. Ordinary Oak Builder builds do not load this adapter.

## Build with Magnolia Builder

```sh
# Local HTML launch descriptor supplies <applet code=...> and <param name="name" value="Ada">.
magnolia build --entry test/fixtures/CounterApplet.html --output Counter.oak
magnolia Counter.oak
# Hello from Magnolia / 8 / 14 / Ada

# Browser/Node output, with Builder compiling the verified in-memory Oak entry.
magnolia build --entry test/fixtures/CounterApplet.html --output Counter.js
node Counter.js
magnolia build --entry test/fixtures/GraphicsApplet.class --paint \
  --width 100 --height 60 --output Graphics.html
# Open Graphics.html in a browser: bytecode Graphics calls paint the Canvas.

# Application, not an Applet: JNLP supplies a local JAR and an argument.
magnolia build --entry test/fixtures/HelloMain.jnlp --output Main.oak
magnolia Main.oak  # 1 / Ada / 14

# Oak-only Java *input*: no javac even during conversion (single-file subset).
magnolia build --entry test/fixtures/HelloMain.java --java-backend oak \
  --args '["Ada"]' --output HelloFromJava.oak
magnolia HelloFromJava.oak  # 1 / Ada / 14

# Or build-time javac; generated JS runs under Node/browser without a JVM.
magnolia build --entry test/fixtures/LibraryMain.java --java-backend javac \
  --args '["Oak"]' --output LibraryMain.js
node LibraryMain.js  # -2147483648 / -1 / 6 / -4 / 5 / true / -2147483648 / -13

# Optional Java *output*. Requires a JDK only to compile/run the output.
magnolia build --entry test/fixtures/HelloMain.class \
  --args '["Ada"]' --output MainConverted.java
javac MainConverted.java && java MainConverted
```

Output extension selects `.oak`, `.js`, `.html`, or `.java`. `--web` and `--java` may be passed, but must match the extension. Other Builder targets (WASM, native, TypeScript, etc.) have **no verified JVM adapter** and are rejected rather than pretending to work. A `.class` file is not Java *source*; `.java` input is compiled into verified class bytes by the selected frontend **before** the same Oak JVM linker runs. `auto` (the default) uses `javac` when available and falls back to the Oak subset only if `javac` is unavailable; it does not silently treat every Java program as supported.

The adapter is dispatched inside `lib/build.oak`. For JS/HTML it supplies `config.source` as a **virtual Oak entrypoint**, lets the normal Builder parse/import/transpile/bundle it in memory, and writes only the requested final output. Ordinary Oak input still follows the previous Builder path. Native Oak uses `lib/applet-oak.oak`'s readable standalone emitter inside the Builder route: the general-purpose Builder's Oak renderer currently misrenders some generated nested case branches (`Unexpected token ->`), so replacing the working emitter with its native bundle would make applets unrunnable. The Java target uses the Magnolia-authored applet-specific emitter, with its preamble **inline in `lib/applet-java.oak`**; the generic Oak→Java renderer is not a substitute.

### Supported local input packages and entrypoints

| Input | Selection | Scope |
| --- | --- | --- |
| `.class` | Applet/JApplet ancestor or `public static main(String[])`; explicit `--entry-class` / `--method` where needed | Additional class dependencies via `--class-files helper.class[,more.class]` |
| `.java` | Resulting class's applet or `public static main`; `--entry-class` when ambiguous | `--java-backend oak` compiles one small Java 8-style subset in Magnolia; `--java-backend javac` compiles explicit sources with isolated build-time JDK, `--java-files` and explicitly staged `--class-files` helpers |
| `.jar` / `.zip` | Valid folded-manifest `Main-Class` wins when supplied; otherwise prefer Applet subclasses and then static `main`; explicit `--entry-class` overrides | Stored/raw-DEFLATE classes, CRC32 and central-directory validation; no ZIP64/encryption, remote manifest Class-Path not fetched. `--select-class JVM/name` scans a bounded larger JAR's index and extracts **one** class, without the other classes or original JAR. |
| `.applet.json` | Verified archived entry and optional original JAR | Class names/SHA-256 are checked on extraction |
| local `.html` / `.htm` | First `<applet code="...">`, optional `archive`, `codebase`, `<param>`, dimensions | Simple bounded tag reader; **not** a general HTML/object/embed parser |
| local `.jnlp` | `<applet-desc main-class>` or `<application-desc main-class>`; local `<jar href>`; applet `<param>` / application `<argument>` | Local relative paths only; remote codebases, DTDs/entities, full JNLP policies not implemented |

Build-time `javac` is invoked with `-proc:none`, `--release` (default 8), an isolated sourcepath/classpath and explicitly supplied `.java` files or staged helper `.class` files; compiled classes are parsed and bounded before Oak linking. Select `--java-backend oak` for no-JDK builds; it handles one ASCII source/class with limited integer/string methods, branches, `while`, `main` and simple applet lifecycle (not arbitrary Java). An absent `javac` in `auto` mode causes an Oak-subset attempt, not an unbounded compiler fallback. `--javac-path` and `--java-release` are optional build-time settings. Neither path adds a JVM dependency to the final Oak/JS/HTML.

Oracle's [1](https://docs.oracle.com/javase/8/docs/technotes/guides/javaws/developersguide/syntax.html) JNLP syntax describes `resources/jar`, `applet-desc`, `application-desc`, parameters, and arguments; this reader implements only the local forms described above. It does **not** fetch remote archives, load JAR resources through Java APIs, or run arbitrary installers. If an archive has several unrelated applet entrypoints, specify `--entry-class internal/ClassName`; selection is deterministic, not a guess about user intent.

```sh
magnolia build --entry test/fixtures/HelperFirst.jar \
  --param name=Ada --output Applet.oak
magnolia build --entry test/fixtures/InheritedPublicApplet.class \
  --class-files test/fixtures/PublicBaseApplet.class \
  --paint --output Inherited.html
magnolia build --entry test/fixtures/CounterApplet.jnlp --output Counter.js
# javac sees Helper only because it is explicitly staged; implicit discovery is off.
magnolia build --entry test/fixtures/UsesHelper.java --java-backend javac \
  --class-files test/fixtures/Helper.class --output UsesHelper.oak
# Additional Java files are likewise explicit; SourceMain alone cannot find SourceUtil.
magnolia build --entry test/fixtures/SourceMain.java --java-backend javac \
  --java-files test/fixtures/SourceUtil.java --output SourceMain.oak
# A JAR can contain interfaces and inherited/overridden implementations.
magnolia build --entry test/fixtures/InterfaceMain.jar --output InterfaceMain.oak
magnolia InterfaceMain.oak  # 3 / 8 / 7 / true / true / 7
```

Builder accepts `--param name=value`, `--params params.json`, `--args '["Ada"]'`, `--paint`, `--width`, `--height`, `--max-steps`, `--entry-class`, `--select-class` (JAR/ZIP only), `--class-files`, `--java-files`, `--java-backend auto|javac|oak`, `--javac-path`, `--java-release`, `--no-optimize`, and an inline `--applet '{"options":{...},"params":{...}}'` configuration. HTML/JNLP parameters and dimensions are used by default; CLI values override them. The `build` subcommands (`inspect`, `coverage`, `build`, `run`, `transpile`, `restore`, `recompile`) provide inspection, runtime, packaging, restoration and Builder-conversion helpers:

```sh
magnolia build inspect --entry test/fixtures/CounterApplet.html --check
magnolia build run --entry test/fixtures/CounterApplet.html
magnolia build transpile --entry test/fixtures/CounterApplet.html --output Counter.oak
magnolia build build --entry test/fixtures/Applets.jar --output Counter.applet.json
magnolia build restore --entry Counter.oak --output restored-classes
```

## Programmatic Oak API and round-trip

`import('build').run(config)` accepts an `applet` options object. `inMemory?` returns the result without writing a file. The higher-level `lib/applet-api.oak` exposes `load`, `inspect`, **`coverage`**, `compile`, `run`, `convert`, `archive`, `restoreOak`, **`recompileOak`**, **`recompileJar`**, **`recompileTo`**, `restoreArchive`, `restoreFile`, `restoreTo`, `compare`, `session`, and `command`; these use the same verifier/Builder rather than a separate applet builder. For extending the JVM host API, see the [shared `jvm-vm-libs.oak` contract and tests](jvm-vm-libraries.md).

```oak
api := import('./lib/applet-api.oak')
info := api.inspect('test/fixtures/CounterApplet.html', {check?: true})
run := api.run('test/fixtures/CounterApplet.html', {})
built := api.convert('test/fixtures/CounterApplet.html', 'Counter.oak', {})
// api.convert defaults to in-memory; built.value is complete runnable Oak text.
restored := api.restoreOak(built.value)  // original byte strings + archive metadata
comparison := api.compare('test/fixtures/CounterApplet.html', built.value, {})
// comparison.equal? checks bounded stdout, return value, drawing.
// comparison.bytesIdentical? checks exact original class-file bytes.
```

For **every supplied class and method**, including code excluded by reachable-entry compilation, use `api.coverage('test/fixtures/InitMain.class', {classFiles: ['test/fixtures/InitBase.class', 'test/fixtures/InitHelper.class', 'test/fixtures/UnreachableInit.class']})` or:

```sh
magnolia build coverage --entry test/fixtures/InitMain.class \
  --include test/fixtures/InitBase.class,test/fixtures/InitHelper.class,test/fixtures/UnreachableInit.class --json
# Nonzero exit: UnreachableInit.<clinit>()V references unsupported java/awt/Font.
```

The structured result counts classes, fields, methods, decoded bodies, abstract declarations and blockers with member names/messages. Unlike `inspect --check`, it also reports an unreachable unsupported initializer, missing superclasses/interfaces and unmodeled signature reference types. A clean report means only that this **bounded decoder and type inventory** admitted these bodies: it is **not** Java bytecode verification, full linking, proof of equivalent execution, or a promise of all applet host APIs. The report refuses oversized/invalid input and more than 1024 blockers rather than silently truncating.

The standalone Oak source includes named readable instruction functions, the Magnolia interpreter logic **inline** (with the shared `lib/ieee754.oak` and `lib/int64.oak` primitives inlined alongside it from the same directory), and a SHA-256-checked archive comment. Altering the generated Oak instructions changes *Oak execution*. `restoreOak` still retrieves the embedded **original** class/JAR bytes. The **separate** `recompileOak` route statically parses the method-literal AST and patches verified JVM instructions to create **new class files**; it never runs supplied Oak code, and it checks that the inlined interpreter was not edited. Older Oak files missing either the inline runtime **or the invocation integrity marker** must be retranspiled before using `recompileOak`; original-byte `restoreOak` still works. `api.restoreArchive` verifies the `.applet.json` returned by browser/Node `restoreArchive()`. Both `restoreTo` and `recompileTo` refuse to overwrite an existing output. The CLI can restore original classes into a directory:

```sh
magnolia build restore --entry Counter.oak --output Counter-original.class
cmp Counter-original.class test/fixtures/CounterApplet.class
magnolia build restore --entry Applet.oak --output Applet-original.jar
```

### Edited Oak → genuinely new `.class` or `.jar` bytes (bounded reassembly)

Use the **same Magnolia Builder** in reverse. For an ordinary application:

```sh
magnolia build --entry test/fixtures/HelloMain.class \
  --args '["Ada"]' --no-optimize --output /tmp/HelloMain.oak
# In the generated applet_method_... list change [16, 7, ?] to [16, 11, ?].
mkdir -p /tmp/edited-classes
magnolia build --entry /tmp/HelloMain.oak \
  --output /tmp/edited-classes/HelloMain.class
# Optional independent check: java -Xverify:all -cp /tmp/edited-classes HelloMain Ada
# Prints 1, Ada, 22 (original class printed 1, Ada, 14).
```

`build recompile --entry Edited.oak --output New.class` does the same for a single class. For an Oak file made from **multiple classes**, use `--output DIRECTORY` in that tool (or `api.recompileTo(sourceText, directory)`), which writes each package-relative `.class` without overwriting originals. Programmatically `api.recompileOak(sourceText)` returns `{type: :ok, value: [newClassBytes, ...], changedInstructions: n}`; Builder accepts `{entry: 'virtual.oak', source: sourceText, output: 'New.class', inMemory?: true}` for **one** class. A new ASCII `ldc` string can be appended to the constant pool if its index still fits the existing instruction width. For example, the multi-class `InitMain` fixture was edited from `prepared` to `retuned`: its recompiled classes passed Java 11 `-Xverify:all` and produced the same output as edited standalone Oak, including parent-first static initialization.

If the Oak archive embeds an **original JAR**, the same verified edit can produce a **new unsigned JAR** in Oak without invoking Java's `jar` tool:

```sh
magnolia build --entry test/fixtures/EditableMain.jar --args '["Ada"]' \
  --output /tmp/Editable.oak
# Edit the named HelloMain instruction [16, 7, ?] to [16, 11, ?] in /tmp/Editable.oak.
magnolia build --entry /tmp/Editable.oak --output /tmp/EditableNew.jar
java -Xverify:all -jar /tmp/EditableNew.jar Ada  # 1 / Ada / 22
# Or: magnolia build recompile --entry /tmp/Editable.oak --output /tmp/EditableNew.jar
```

`api.recompileJar(sourceText)` returns `{type: :ok, value: newJarBytes, changedInstructions: n}` without staging output, and Builder accepts a virtual Oak source plus `.jar` output with `inMemory?: true`. It verifies every copied JAR member (including unused resources), repacks bounded ZIP32 entries, preserves the manifest, **nonclass resource contents**, and all original class names, but does **not** preserve compressed container bytes, timestamps or signatures. Edited signed/digested JARs fail closed rather than emitting an invalid signature. Zero-edit `recompileJar` returns the exact original JAR; `restoreOak` always returns original bytes even after edits. Edited `.zip` and `.applet.json` output remain unsupported.

**Recompiler limits:** original bytecode instruction count and widths, member/field declarations, max stack/locals, branch offsets, exception tables and stack maps must stay unchanged. Accepted edits are bounded `bipush`/`sipush`, `iinc` increments, small `iconst` variants, integer arithmetic/comparison opcodes of the same stack shape, an `if_acmpeq` ↔ `if_acmpne` reference-identity switch, and same-type `ldc`/`ldc_w` constants (new **ASCII** strings or 32-bit integers only when pool/index limits permit). Switch opcodes (`tableswitch`, `lookupswitch`) may be left unchanged while an allowed equal-width instruction **elsewhere in the same method** is edited; the reverse compiler validates their variable-length layout and rejects edits to their cases/targets instead of discarding them. Other changes fail explicitly; `--no-optimize` avoids folded instruction listings that cannot be mapped back to original byte offsets. This is **not** arbitrary Oak→Java bytecode compilation, complete JAR authoring, or proof that any unsupported `.class` runs. The Oak JAR writer is limited to an edited standalone Oak source containing the original JAR. A JVM is **not** needed to recompile; it was used only as an independent test oracle. `restoreOak` remains byte-exact original recovery, even after edits. The five-class `InterfaceMain.jar` fixture exercises a transitive interface call, subclass override and `checkcast`; editing its `score(Named)` constant **after** a five-byte `invokeinterface` generated a verified, runnable new JAR that printed `4\n9\n7\ntrue\ntrue\n7\n` and kept its nonclass resource. A prior zero-edit `MathMain` byte-10 discrepancy was caused by comparing **different javac builds** (the fixture with debug/source metadata versus an isolated `-g:none` build), not by the recompiler: each source round-trips byte-exactly against **its own embedded class**, with both optimized and unoptimized Oak; JDK 11 `-Xverify:all` ran the recompiled class.

Exact JAR container preservation is limited to original ZIP data of at most 8 MiB and a 48 MiB textual archive; `--no-original-jar` / `applet.preserveJar? = false` keeps the classes without the container. Builder-generated JS/HTML embeds the verified original archive by default so its host API can return it. `applet.embedOriginal? = false` omits that extra JS/HTML payload if desired. Browser code returns the archive **for verification by Magnolia**; it does not claim to compile modified Java bytecode.

## Browser and Magnolia GUI host interactions

Browser/Node bundles run the applet once at load and expose `globalThis.MagnoliaApplet`. For a loaded HTML page:

```js
const host = window.MagnoliaApplet;
console.log(host.inspect());
host.on("drawing", frame => console.log(frame.commands.length));
host.setParameter("name", "Bob");
const result = host.run();                  // re-execute from a fresh bounded state
const original = host.restoreArchive();    // .applet.json text, verify with Magnolia
console.log(host.dispatch({type: "event", name: "keydown"}));
// -> {type: "error", message: "AWT input event dispatch is not implemented"}
```

`dispatch` also supports `inspect`, `setParameter`, `run`, `paint` (only if that method was verified into the build), `invoke` (only methods already in the verified bundle), and `restore`. A noncompiled method or unimplemented event produces an error. **Repeated runs are fresh**, not a persistent JVM applet object; keyboard/mouse dispatch, repaint scheduling, threads and complete AWT interaction remain unimplemented.

The self-contained HTML uses Canvas 2D to draw bytecode-generated `Color`/`Graphics` commands (`setColor`, `fillRect`, `drawLine`, `drawString`); the Oak CLI can save corresponding SVG, and the Java target paints to a PNG using real JDK AWT. Programmatically `lib/applet-host.oak` maps recorded drawing commands to Magnolia's `gui-web.oak` command queue via `createWeb`/`replayWeb`, with a `webBridge` callback and **exact rectangle coordinates**. `replayNative` accepts an externally created `GUI.oak` Windows/X11 window, but was not exercised with a native display here and its text/rectangle pixels are not guaranteed to match AWT. `sendEvent` returns an explicit unsupported error. This is an incremental host adapter, **not a complete AWT widget/event model**.

## Historical input research and reproducible blockers

- The Menu Tree Applet v1.00 archive ([11](https://archive.org/details/mtree100_zip)) includes `index.html`, `menutree.jar` and images. Magnolia reads its local HTML descriptor, extracts six JAR classes and selects `MenuTreeApplet` even though helper classes precede it. An earlier `inspect --check` stopped at the exception table in `MenuTreeApplet.init()`; bounded handler support has since been ported to this tree (2026-11-16), but this separate archive has **not** been rerun through selected execution after that change (no JDK here). Its AWT components, images, URL/context access and event handling still need implementation. Its redistribution terms have not been verified.
- The Xstreamer archive ([1](https://archive.org/details/xstreamr_zip)) includes `Default.html` and `Xstreamer.class`. Magnolia reads 13 `<param>` values; preflight stops at `Xstreamer.init()V at byte 198: unsupported or missing class java/awt/Font`. Accurate font metrics, AWT colors/components and animation scheduling require additional work. Its redistribution terms have not been verified.
- The [Fliplay reference](https://www.oortman3d.com/fliplay/) supplies the original `fliplay.class`/`flickframe.class` (JVM 45.3). A hash-checked runnable harness executes its **original** `get_ubyte`/`get_ushort`/`get_ulong` bytecode and constructor against a JDK oracle in native Oak, standalone Oak, Node and generated Java. That is not the graphical player: `fliplay.init()V` still stops at byte 534, `unsupported Java API java/awt/Component.createImage(II)Ljava/awt/Image;`. Its FLI/FLC inputs, threads, imaging, URL/stream I/O and AWT event model are not implemented. The [archived Fliplay README](https://web.archive.org/web/20080209003513/http://vento.pi.tu-berlin.de/STROEMUNGSAKUSTIK/SOFTWARE/fliplay/README.html) documents JDK 1.1.8 and controls. Obtaining its Java source is **not required** for bytecode work.
- The Applet Password Wizard 3.0 archive ([13](https://archive.org/details/passwd30_zip)) was inspected, not executed: its ZIP contains a Windows installer (`setup.exe`) rather than a directly loadable class/JAR. Do not conflate a product called an "applet" with an immediately runnable Java applet file; installer execution and rights review are outside this workflow.

The linked [DededeDrip.NET](https://www.dedededrip.net/main.php) example directs visitors to a [Flash portal](https://www.dedededrip.net/flash.php); its [Rapper Rhythm page](https://www.dedededrip.net/viewflash.php?id=rapperrhythm) links `.FLA` and `.EXE` downloads. It is **not evidence of a Java `.class` applet**. The linked game EXE ZIP returned 404 after the site's JavaScript challenge; it has **not** been inspected. An independently archived Adobe Flash projector EXE and the Linux PPAPI plugin were statically identified as **native** binaries. In contrast, an archived Macromedia **Flash Player Java Edition** SDK contains 20 genuine JVM `.class` files (including `Flash.class`, an `Applet`), **limited by Macromedia to Flash 2 features**. Magnolia actually executes selected *original* `SParser`, `Matrix` and Point-based `Rect` bytecode with JDK parity, including separately edited and JDK-verified parser/geometry classes; this is **not** player execution. The patch author reports that the constructor now decodes under bounded F/D support, but a fresh selected-entry Flash run stops **before execution** at `Flash.start()V` byte 6 (`monitorenter`, opcode `0xc2`); that run was **not reproduced in this tree** (no JDK, and this tree cannot currently build). The whole player has not been converted. See [Flash artifact inspection, hashes and compatibility status](flash-artifacts.md). Ruffle's [Flash emulator](https://github.com/ruffle-rs/ruffle) and [Flashpoint Archive](https://flashpointarchive.org/source) are preservation references, **not Java JVMs**. [Jmol-SwingJS](https://github.com/BobHanson/Jmol-SwingJS) illustrates a browser migration approach; it is not embedded into Magnolia output. A J2ME MIDlet JAR, Flash SWF, or Windows Control Panel applet is not interchangeable with a Java browser applet.

## Limits, safety and next milestones

Current bytecode support covers reachable integer/reference methods, bounded primitive/reference arrays including exact **`long[]`** (`newarray` type 11, `laload`/`lastore`, default zero, nested `[J` references built with `anewarray`), metered **`multianewarray`** (fully or partially allocated boolean/byte/char/short/int/long and modeled reference arrays; up to 16 dimensions), **unsigned char and signed short storage**, `i2b/i2c/i2s`, category-one `dup_x1/dup_x2/dup2` and **`wide` int/reference locals and signed `iinc`**, aligned **`tableswitch`/`lookupswitch`**, `if_acmpeq`/`if_acmpne` with **reference identity** and other selected branches, limited fields, **JVM `ConstantValue` preparation and lazy parent-first `<clinit>`**, inherited field/static-method resolution, receiver-based `invokevirtual` **and bounded `invokeinterface`** dispatch over reachable instantiated classes, transitive user-defined interface membership for `instanceof`/`checkcast`, static main, inherited applet lifecycle, parameter lookup and the drawing subset above. The [shared VM-library boundary](jvm-vm-libraries.md) admits exact descriptors/opcodes for basic `System.out`, ASCII `String`, `StringBuilder`/`StringBuffer`, integer `Math` and `Integer` (`parseInt/toString/compare/sum/max/min`), plus selected AWT `Color`/`Graphics` calls, a mutable, identity-bearing `Point` geometry island (`x/y` int fields, three constructors, `move`, `setLocation(II)`, `getLocation`, `equals`), and exact `System.arraycopy(Object,int,Object,int,int)` on bounded primitive/reference arrays with overlap, per-element covariance and metered host work. `Math.abs(INT_MIN)` and signed 32-bit `Integer.sum` match Java. JVM objects, arrays, selected host singletons and String references have bounded run-local IDs (at most 1,048,576), with `ldc` literals interned and host-created/argument strings given separate identities. `Color.WHITE` and `Color.white` share one cached reference; `System.out` is stable across reads. Every reachable method passes a bounded **I/A/J/F/D frame gate** (int/boolean vs reference vs two-slot long vs float/double, across locals, stack effects, fields, calls, returns and control-flow merges, plus array-opcode compatibility when the component descriptor is statically known) before execution; this is a subset check — **not** full JVM verification, a complete heap, or Java UTF-16 semantics. Three further **bounded verifier slices were hand-ported into this tree on 2026-11-17 (not built or executed here; no JDK)**: a **known-reference-subtype slice** that retains class/interface/array descriptors through calls, fields, returns and branch joins (JDK-invalid statically known mismatches fail before output; JDK-valid interface dispatch and `aastore` failures stay *dynamic*), a **bounded selected-entry allocation-site/constructor slice** that carries uninitialized `new` tokens through locals and stack aliases, distinguishes uninitialized constructor `this` with a separate readiness state, initializes matching aliases only after a matching `invokespecial <init>`, and refuses a constructor return without a valid `this()`/`super()` call (JDK-valid identity/null tests of uninitialized values remain legal), and a **narrow pre-super own-field `putfield` exception** admitting a store on uninitialized constructor `this` **only** for an instance field directly declared in the current class before `this()`/`super()` (inherited fields, fresh `new` objects, `getfield`, and static fields do not qualify). A **version-aware bounded `StackMapTable` structural gate** for version 50+ Code attributes parses named frame/type families in a length-limited reader and rejects reserved tags, invalid object-pool indexes, invalid `new` offsets, frame offsets beyond Code, duplicate tables and inconsistent lengths; selected reachable methods additionally require declared frame PCs and uninitialized `new` sites to be decoded instruction boundaries, and frame targets are preserved against folding. Version 49 retains the JDK-accepted legacy ignore behavior. A further **required-target frame slice** (hand-ported 2026-11-17) requires a declared `StackMapTable` frame at every supported conditional/`goto`/switch target and exception-handler entry for **selected reachable methods of version 51+**, including a backward edge to bytecode PC 0 (the JDK does not treat the implicit initial frame as a substitute); version 50 keeps JDK 11's inference fallback and version 49 its legacy inference, so a missing table is not rejected indiscriminately. This still does **not** compare declared/inferred-frame types (assignability), cover every block start, verify otherwise-unselected methods, implement full access/handler/uninitialized-object rules, or deliver full JVM verification or whole-player execution. The optional Java-source emitter decodes modified-UTF-8 identifiers/literals into ASCII-only Unicode escapes (as verified on original Macromedia class names); this does **not** change the Oak/JS VM's limited Java `String` model. IEEE-754 `float`/`double` arithmetic and conversions, F/D constant pools, two-slot doubles, category-two `dup2`-family forms (`dup_x2`/`dup2_x1`/`dup2_x2`) and `swap`, and bounded catchable Java exceptions (typed/ordered handlers, `athrow`, initialization-failure wrapping) have been **ported into this tree on 2026-11-16**, along with the corresponding verifier shape/gate changes. They have **not** been built or executed in this environment (no JDK available; the pre-port suite was 300/300 native). `lmul`/`ldiv`/`lrem` and long shifts in the Oak runtime are **not** a validated port: unvetted assistance fills that gap pending an explicit keep/revert decision (the implementation now lives in the shared `lib/int64.oak` primitive imported by the Oak runtime; its `lmul` carry bug was repaired on 2026-11-16 and is covered only by a Python limb model, not a JDK), and the generated-Java emitter still lacks the long-arithmetic slice altogether. Unsupported exception tables, interface default methods and dynamic proxies, most host APIs, synchronization, threads, complete array component/subtype verification, remaining multidimensional reference kinds, dynamic class loading/resources, **general** Java source compilation, non-ASCII String semantics, complete AWT/Swing, native event loops and Fliplay animation all remain open. A full JVM compatibility layer would require each of these, test suites across real class versions, resource/security policy and behavioral parity against a reference JVM. None is implicitly delivered by packaging, inspection or original-byte restoration.

Class parsing is capped at 8 MiB/class; ZIP input at 64 MiB, 256 classes and 32 MiB extracted class data; runtime instruction/array/depth/draw budgets, a 1,048,576-reference ceiling and a 4 MiB output ceiling limit accidental runaway work. Oak-only Java source is limited to 1 MiB/8192 tokens, and `javac` source is explicitly staged into an isolated build directory. **These checks are not a secure sandbox for untrusted bytecode or untrusted build tools.** Review inputs and use OS isolation as appropriate. Local HTML/JNLP launch readers do not fetch remote URLs.

## Verification

```sh
magnolia test/main.oak applet
# 300/300 focused native tests in this tree (2026-11-16, after the long[]/multianewarray
# hand-port plus a from-scratch exact-limb land/lor/lxor, which this tree lacked). The
# author's later checkpoints report 383/383 native at the structural StackMapTable gate and
# 393/393 bundled-JS at the required-target frame slice; the target-frame slice has native
# targeted/JDK oracles but no completed updated full native focused rerun is claimed for it.
magnolia build --entry test/main.oak --output /tmp/applet-test.js --web \
  --include applet.test:test/applet.test.oak
node /tmp/applet-test.js applet
# NOT confirmed in this environment: the bundled-JS suite aborts on a web-target codegen
# bug (assignment expressions return the assigned value instead of the mutated target),
# unrelated to applet logic; the historical bundled-JS figures were never reproduced here
# JDK-required oracle and matrix harnesses (long-array, multi-array, stack-forms,
# constructor-verifier, preinit-field, stackmap-structure, stackmap-target-frames,
# long-float, html-escape, route-matrix and the historical-core matrix) were
# present during development but have been ablated from this tree. Their
# author-reported results are summarized below; none was rerun here (no JDK).
```

## Moon Clouds: all frames through original Fliplay loader/decoder (not the whole Applet)

During development, an all-frame differential oracle verified the uploaded Moon Clouds ZIP/FLC and archived `fliplay.class`/`flickframe.class` by SHA-256 before execution, invoking unmodified original `fliplay.load_header → read_loop → load_frame → extract_frame` and chunk decoders through the native Magnolia JVM subset. A Java 8 link fixture supplied method reachability and bounded VM array/stream constructors. The JDK 11 `-Xverify:all` harness independently called those original methods on a continuous FLC file. Author-reported: on all **60/60** display frames every indexed pixel and palette byte agreed with the JDK, and all RGB pixels agreed with Pillow's separate FLC decoder. The oracle and viewer tools have been ablated from this tree; the historical fixtures remain for reference.

**Still incomplete:** Fliplay's original `init`/`load_fli` URL-opening and AWT `createImage`/`MemoryImageSource`/`Graphics` rendering, patch caching, applet controls and threaded/reverse playback; the historical Java Flash 2 Applet (`Flash.start()` still blocks on `monitorenter`), AVM1 execution of the separate Flash-5 SWF, and native projector compatibility.

The `NarrowIntMain`, `PrimitiveArrayMain`, `StackShuffleMain`, `SwitchMain`, `WideLocalMain`, [`ReferenceIdentityMain`](../test/fixtures/ReferenceIdentityMain.java), [`PointGeometryMain`](../test/fixtures/PointGeometryMain.java), [`ArrayCopyMain`](../test/fixtures/ArrayCopyMain.java), [`LongArrayMain`](../test/fixtures/LongArrayMain.java) and [`MultiArrayMain`](../test/fixtures/MultiArrayMain.java) Java 8 fixtures were compared with JDK 11 `-Xverify:all` across the original ten fixtures' routes (native Oak, standalone Oak, Builder JS/Node and generated Java); the two new fixtures added on this branch have not been compared with a JDK. The LongArray fixture checks exact-limb `long[]` (`newarray` type 11, default zero, two-slot `laload`/`lastore`, nested `[J` references, overlapping typed-primitive copies); the MultiArray fixture checks bounded `multianewarray` (outer-zero/negative dimensions rejected, allocation budgeted before construction). Their focused checks pass in this tree (300/300 native, including exact original-byte restoration via the standalone-Oak route); their JDK comparison was **not** run here. The identity fixture checks both `if_acmpeq`/`if_acmpne`, distinct equal-valued objects/arrays/strings versus aliases, null, interned literals, separate equal arguments, `String.concat("")`, stable `System.out`, and `Color.WHITE`/`Color.white`. The Point fixture checks mutable `x/y` fields, constructors/copies, method calls, `equals` versus identity, array references and subtype tests. The arraycopy fixture checks five primitive array kinds, overlapping copies, reference identity/covariance and zero-length boundaries; null, mismatched arrays and bad ranges fail explicitly; after the 2026-11-16 port, incompatible reference stores raise a catchable modeled `java/lang/ArrayStoreException` (not executed here). All three have self-contained HTML that matches the JDK in a network-free simulated DOM. Preflight rejects malformed switch padding/keys/targets, wide operands and reference-branch operands/targets; the VM host rejects unregistered Point fields, wrong receivers and unsupported AWT methods. All **eight** fixtures restore their original class bytes exactly without edits. **Frame preflight (2026-11-16):** all reachable methods now pass the bounded I/A/J gate — including `StackHeightMain` (typed `iconst_0`→`arraylength` and reference/int merge rejections, covered by two new focused tests) and the long/height mutation classes The JDK-required oracle and route-matrix harnesses were **not** rerun in this environment; see the [compatibility plan](applet.md) for the exact status.

**Later bounded checkpoints (author-reported; none of these were reproduced in this tree):** The exception matrix matched 13 JDK lines for ordered handlers, `athrow`, propagation and a checked edit across Oak/JS/HTML/Java; The clinit matrix matched five JDK lines for three-class failed initialization and byte-exact restoration. Targeted malformed-exception probes **5/5** and the corrected verifier matrix **27/27** matched JDK rejection. The stack-forms oracle passed all **nine** legal `dup_x2`/`dup2_x1`/`dup2_x2`/`swap` category forms on native Oak, standalone Oak, JS and generated Java (seven malformed width/category forms rejected, including wide `swap`). [`LongFloatRoundingMain`](../test/fixtures/LongFloatRoundingMain.java) exposed a real `l2f` double-rounding bug for which the author reports a direct-limb fix matching **36** JDK boundary output lines across Oak/JS/Java and **387** seeded JDK/native-Oak/JS midpoint values The author further reports a fully logged **337/337 native** focused suite pass (128s under a 175s hard cap) and **337/337 bundled JS**, plus a fresh selected Flash attempt stopping at `Flash.start()` `monitorenter` (0xc2) and Fliplay `init()` stopping at `Component.createImage(II)`. **In this tree, none of these was run: there is no JDK here and the post-port tree has not been built.** The tools and fixtures are present and the Oak-side/exception/F-D changes were ported on 2026-11-16, but their test evidence is the author's, not this environment's. The author further reports the 2026-11-17 slices: a **known-reference-subtype slice** (`RefSubtypeMain` 13-line oracle across native class/JAR, standalone Oak, bundled JS, offline HTML and compiled generated Java, original JAR restored byte-exact; five JDK `VerifyError` mutants refused before output, two JDK-valid mutations failing dynamically; Applet→Panel/Container/Component ancestry regression; **346/346** focused native and bundled JS); the **constructor-frame slice** (`ConstructorFrameMain`, eight JDK `VerifyError` mutants, **360/360**); the **pre-super own-field slice** (`PreSuperFieldMain`/`PreSuperFieldBase`, **371/371**); and the **version-aware `StackMapTable` structural gate** (`StackMapFrameMain`/`StackMapExtendedMain`/`FrameCtor`, 16 JDK `ClassFormatError`/`VerifyError` mutants and four Builder no-emission targets, **383/383** focused native and bundled JS). The author's next slice is the **required-target frame gate** for version 51+ selected reachable methods (`StackMapEntryLoopMain` plus a synthetic `FrameTarget` class in the focused tests, and the new target-frames oracle): nine JDK `VerifyError`/pre-output refusals across missing conditional/switch/handler/PC 0 frames, nine versioned JDK-valid positives, four Builder no-emission targets, and one unselected-method distinction; the updated structural oracle now reports **17** refusals and retains a single declared-type non-parity. The author reports **393/393** bundled-JS focused tests at this checkpoint; a complete new native focused-suite rerun is **not** claimed for this slice. The new fixture classes, expected outputs and oracle scripts were added to this tree on 2026-11-17 alongside the verifier/class-reader code, and the focused test file gained the corresponding assertions (constructor, pre-super, stack-map, reference-subtype); **none of that has been built or executed here** — no JDK, and the interpreter has not been rebuilt. The later [ActionScript plan](../docs/applet.md) is separate from this JVM work.

Editing the identity fixture's `if_acmpeq` to `if_acmpne`, `SwitchMain`'s `bipush 20` to `21`, and `WideLocalMain`'s signed `wide iinc 200` to `201` separately produces **different** JDK-verified classes with matching edited Oak; an unsafe reference→integer branch is refused. An edited JAR with the identity class preserves its manifest/resource and runs with `java -Xverify:all -jar`, while restoring its Oak recovers the original JAR bytes. Isolated `javac` `.java` input also executes identity, Point and arraycopy sources; the smaller Oak-only Java parser rejects unsupported object/reference syntax rather than pretending to compile it. `coverage` inventories methods beyond the selected execution path: on the independently obtained Flash ZIP it reports 20 classes, 277 methods, 321 fields, **199 decoded bodies and 122 blockers** at the width-gate checkpoint; the patch author later reports 218/277 with 101 blockers (after exact-limb long arithmetic), **219/277 with 100 blockers after bounded multidimensional arrays** (the newly decoded body being original `Bitmap.<clinit>`), and after the IEEE/exception/shuffle checkpoint **221/277 with 96 first blockers** plus two-class Fliplay **17/35 with 30 blockers** — none of these was rerun in this environment. Fliplay is 14/35 with 33 blockers at this tree's last verified state. **These are not runnable-player coverage figures** (see [plan](applet.md)). The separate historical core matrix actually executes original Flash `SParser`, `Matrix` and now `Rect` bytecode with JDK parity (parser and Rect also as inline HTML), as well as selected original Fliplay methods. The newly compiled Point fixture is distinct from the original `Rect` geometry that now executes; decoding other original methods still does not mean they were run.

Generated JS/HTML was also run under Node and a simulated DOM: the latter recorded a red rectangle `(4,5,30,12)`, blue text and a line, and replayed the page through `window.MagnoliaApplet.dispatch({type: 'paint'})`. The `LibraryMain.java` fixture compares `Integer`/ASCII `String` calls, including signed overflow and boundary parsing, across reference Java 11 (`-Xverify:all`), the native Oak bridge, standalone Oak, Builder JS/Node and the optional generated Java backend. Generated Java output also passed JDK checks for static `main`, parent-first class initialization, receiver-based virtual dispatch (`DispatchMain`) and inherited painting. Edited Oak→class tests produced real changed Java 8 class files: `HelloMain` ran with `-Xverify:all` and returned `22` rather than `14`; `InitMain`'s new constant-pool string `retuned` and `DispatchMain`'s overridden value changed to `3` produced matching Oak/JVM output. An edited `EditableMain.jar` was independently run with `java -Xverify:all -jar`, printing `1\nAda\n22\n` while preserving its manifest and nonclass resource contents; signed edits were rejected. Historical Menu Tree/Xstreamer/Fliplay and the Macromedia Java Flash Player produced the explicit blockers above; selected *original* Fliplay and Macromedia core methods execute with JDK parity, but **none of these complete players was claimed to be runnable**. The baseline repository still has an unrelated `test/pack.test.oak` failure in the unfiltered test suite. The only Go change is `lib.go`: the applet/jvm stdlib modules are embedded and registered in the stdlibs map so `TestStdlibRegistryCoversLibDirectory` stays green (the repo invariant that every `lib/*.oak` is registered). Nothing in the JVM/applet pipeline depends on Go code; registration only makes the modules importable by name.
