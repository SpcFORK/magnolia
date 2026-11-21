# Oak VM-library transition boundary for JVM input

This is a **Magnolia/Oak implementation of a bounded JVM compatibility layer**, not a complete JVM or Java standard library. It lets independently acquired `.class`, `.jar` and `.java` inputs enter one verified class image, and lets that image call explicitly admitted Java host-library methods without relying on a JVM in the resulting Oak/JS/HTML. The primary output is readable, runnable standalone Oak. Java *source* output is optional and needs a JDK to execute.

## One class image, several frontends

```text
.class ────────────────┐
.jar/.zip → ZIP/manifest ├→ bounded class bytes → Oak class parser → reachable
.java → isolated javac ┤                                       bytecode verifier
.java → Oak subset ────┘                                               │
                                     resolved, predecoded instruction image
                                          │              │             │
                                  native Oak VM      standalone Oak    Builder JS/HTML
                                                           │
                                                original bytes / checked edits
                                                         → new .class / .jar
```

`lib/jvm-bridge.oak` exposes `load({entry, source?}, options)`, **`coverage(loaded)`**, `link(loaded, selection)`, `execute(image, params, options)` and `run(...)`. `lib/applet-api.oak` exposes the higher-level `load`, **`coverage(entry, settings)`**, `compile`, `run`, `convert`, `restoreOak`, `recompileOak`, `recompileJar` and `recompileTo` routes. Builder invokes the same bridge through `lib/build-applet.oak`; it does **not** need a separate applet builder. JAR manifest selection, verified local ZIP extraction, the isolated `javac` adapter and the Oak-only Java subset all feed class bytes to the **same** linker. Coverage inventories every supplied class/field/method, including methods outside the selected reachability graph, but it does not certify JVM verification, dynamic host calls or behavioral parity. The [staged compatibility plan](applet.md) describes how to extend that boundary without treating a parsed class as a runnable one. The generated standalone Oak inlines its runtime **and** its VM-library code; the runtime's SHA-256 marker prevents an altered host implementation from silently being turned back into a JVM class.

| Layer | Oak source | Responsibility |
| --- | --- | --- |
| Input/selection | `jvm-bridge.oak`, `jvm-java-input.oak`, `jvm-java-subset.oak`, `applet-zip.oak`, `jvm-jar-manifest.oak` | Bounded, explicit input loading; no network classpath or implicit source discovery |
| JVM instruction linkage | `applet-class.oak`, `applet-bytecode.oak`, `applet.oak` | Parse class bytes and interface metadata; resolve reachable methods/fields, concrete receiver overrides and inherited interface declarations; reject unsupported opcodes, descriptors and APIs before emission |
| Host-library contract | **`jvm-vm-libs.oak`** | `jvmMethod(owner, name, descriptor, opcode)`, `jvmField(...)`, `jvmNew(owner)`: exact external member admission. Pure integer/String operations use `jvmInvoke(spec, args)` and `jvmString(spec, receiver, args)`; `jvmColorName(name)` maps uppercase Color aliases to historical singleton names. |
| Stateful interpreter | `applet-runtime.oak` | Recheck the same member contract during execution; execute bounded instructions, initialization, output, applet parameters, reference identity/ASCII String wrapping, a mutable Point geometry host and the limited AWT drawing host |
| Output/reverse | `applet-oak.oak`, `applet-java.oak`, `applet-recompile.oak`, `jvm-jar-write.oak`, `build.oak` | Inline standalone Oak, Builder JS/HTML, optional Java source, original-byte restoration and verified *new* class/JAR editing |

For example, `jvmMethod('java/lang/Integer', 'sum', '(II)I', 184)` returns `{op: 'isum', argc: 2}`; a mismatched descriptor or opcode returns `?`. `jvmInvoke` performs signed 32-bit wraparound, so `Integer.sum(2147483647, 1)` is `-2147483648` in native Oak, standalone Oak and bundled JS. A hand-edited instruction image also passes through the runtime's contract: it cannot use an unregistered host call by changing just its descriptor. This is **an extensible allowlist**, not dynamic access to the host's Java classes.

The reference-branch layer (`if_acmpeq`/`if_acmpne`) gives each admitted JVM object/array/host value a unique run-local numeric identity, wraps JVM String values with identities, interns class literals/`ConstantValue` strings and keeps `System.out`/named Color fields stable. It compares **IDs**, never Magnolia object structure, JS deep equality or raw String characters; a hand-supplied integer reference operand fails instead of being compared. At most 1,048,576 references may be created in one run. Host String operations still receive plain text through the VM-library contract; unsupported Java UTF-16 operations continue to fail. The optional Java emitter instead uses real Java `==` for reference comparisons, non-interned `new String(...)` for external arguments/parameters, and cached host singleton objects. [The differential fixture](../test/fixtures/ReferenceIdentityMain.java) checks both branches, object/array/String aliases and distinctions, null, literal interning and upper/lowercase Color constants; it does **not** verify arbitrary JVM control-flow frames.

The audited **`java.awt.Point` geometry island** admits `new Point()`, `Point(int,int)`, `Point(Point)`, only `getfield`/`putfield` on `x:I` and `y:I`, and exact instance calls `move(II)V`, `setLocation(II)V`, `getLocation()Ljava/awt/Point;` and `equals(Ljava/lang/Object;)Z`. Coordinates are signed int32 fields on a mutable, run-local identity: copying or `getLocation()` creates a **different** reference; equal coordinates make `Point.equals` true but not `==`. The bounded subtype walk recognizes `Point2D`, `Serializable`, `Cloneable` and `Object`. Preflight rejects other descriptors/opcodes and unsupported methods such as overflow-aware `translate` or double-returning `getX/getY`; runtime rejects an edited field reference, wrong receiver or null copy argument instead of returning an Oak object property. The optional generated Java interpreter mirrors these semantics. The [26-line JDK fixture](../test/fixtures/PointGeometryMain.java) passes native Oak, standalone Oak, Node JS, inline HTML and generated Java, while the [historical Rect harness](../test/historical/README.md) invokes **original** Flash 2 `Rect` methods on Point values with JDK parity. This is geometry only, not AWT widget rendering, input or a complete Point2D API.

The exact `System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V` static contract copies bounded JVM primitive or reference arrays, handles in-place overlap through a snapshot, checks primitive component types even for zero elements, validates each covariant destination element in order, and charges each moved element to the instruction budget. The [23-line JDK fixture](../test/fixtures/ArrayCopyMain.java) matches native Oak, standalone Oak, JS, generated Java and offline HTML; five invalid-input modes report explicit VM errors, with incompatible stores also refused in JS and Java output. Java's catchable `NullPointerException`, `ArrayStoreException` and bounds exceptions are **not yet modeled** by this interpreter. This addition advances four original Flash method-body inventory reports, but those methods have not been executed against a JDK.

The linker separately handles **user-defined interface calls**: it decodes `invokeinterface` (`0xb9`) as five bytes, checks its constant-pool tag, argument-count byte and reserved zero byte, resolves an inherited declaration through packaged superinterfaces, and verifies reachable concrete implementations for instantiated classes. The runtime dispatches against the actual receiver, and `instanceof`/`checkcast` follow transitive interface membership. Abstract interface declarations are never executed as methods. Interface **default** methods and interfaces missing from the package fail explicitly. Both native and bundled JS use a bounded subtype walk. Builder JS can represent the reserved Oak property `super` in two ways; `applet-runtime.oak`'s `classParent` handles both, and numeric depth increments use subtraction to avoid JS treating dynamic `+` as string concatenation.

### Adding a JVM library method

1. Admit the exact owner/name/descriptor/**JVM invocation opcode** in `jvm-vm-libs.oak`, with its argument count. For an external field or constructible host type, use `jvmField` or `jvmNew` instead. Do not infer methods from their names alone.
2. Implement the behavior in `jvmInvoke` or `jvmString` if pure, or in `applet-runtime.oak`'s stateful host adapter if it needs output, drawing or applet state. Use the existing instruction/depth/array/drawing budgets and the 4 MiB output ceiling; reject unimplemented effects explicitly.
3. Mirror the semantics in `applet-java.oak` if the optional Java output should support the method. Standalone Oak inlines the library source via `applet-oak.oak`; Builder JS bundles that same Oak source. A non-ASCII String operation currently fails explicitly rather than producing a false Java UTF-16 result.
4. Add a reference-JDK fixture and bounded native/bundled-JS tests in `test/applet.test.oak`. Check the Java backend under `javac/java -Xverify:all`, and test wrong opcodes/descriptors as well as successful calls. Run the class through `.java`/`.class`/`.jar` frontends where relevant.

**Implemented boundary:** `Math.abs/max/min` for integers; `Integer.parseInt/toString/compare/sum/max/min`; ASCII `String.length/equals/equalsIgnoreCase/indexOf/concat`; enumerated `Object` and applet initialization/parameter lookup, `PrintStream.print/println`, `StringBuilder`/`StringBuffer`, AWT `Color`/`Graphics` drawing members, the bounded `Point` constructors/coordinates/move/copy/equals geometry island, `System.arraycopy` on bounded JVM arrays, and `System.out`/named color fields (including JDK uppercase Color aliases). Integer overflow and `Math.abs(Integer.MIN_VALUE)` match Java. Invalid `parseInt` reports a VM error instead of modeling a catchable `NumberFormatException`; inputs longer than 4096 characters are rejected. `PrintStream` output, ASCII concatenation and text-builder results are capped at 4 MiB; an oversized operation fails before appending the new content. The string reader retains modified-UTF-8 class bytes, not decoded Java UTF-16; non-ASCII String operations are deliberately rejected. These declarations are **not a claim that the entire corresponding Java classes work**.

**Blocked/unimplemented:** exception handlers and catchable exceptions, dynamic/remote class loading, JAR resource APIs and manifest classpaths, interface default methods and the full interface API surface, threads/synchronization, reflection, long/float/double and many opcodes (the authentic Macromedia Java Flash player's `dconst_1`, opcode `0x0f`, is a current blocker), a complete Java compiler, full AWT/Swing/browsers, and animation/input event loops. Full JVM/Flash compatibility remains a goal. Parsing a class is not evidence that it runs. The Rapper Rhythm ZIP was not obtained, so its EXE contents cannot be classified.

## Reproduce behavior and limits

From the Magnolia repository root, with the rebuilt `magnolia` executable and a JDK 11 for comparison:

```sh
javac --release 8 -g:none -d /tmp test/fixtures/LibraryMain.java
java -Xverify:all -cp /tmp LibraryMain Oak
magnolia build --entry test/fixtures/LibraryMain.class --args '["Oak"]' --output /tmp/LibraryMain.oak
magnolia /tmp/LibraryMain.oak
magnolia build --entry test/fixtures/LibraryMain.java --java-backend javac --args '["Oak"]' --output /tmp/LibraryMain.js
node /tmp/LibraryMain.js
magnolia build run --entry test/fixtures/LibraryMain.jar --args '["Oak"]'
# All print -2147483648, -1, 6, -4, 5, true, -2147483648, -13 on separate lines.

java -Xverify:all -jar test/fixtures/InterfaceMain.jar
magnolia build --entry test/fixtures/InterfaceMain.jar --output /tmp/InterfaceMain.oak
magnolia /tmp/InterfaceMain.oak
# Both print 3, 8, 7, true, true, 7 on separate lines.
# Editing the score(Named) instruction [4, 0, ?] to [5, 0, ?]
# and building EDITED.oak -> NEW.jar yields 4, 9, 7, true, true, 7.

java -Xverify:all -cp test/fixtures PointGeometryMain
magnolia build run --entry test/fixtures/PointGeometryMain.class
magnolia build --entry test/fixtures/PointGeometryMain.java --java-backend javac \
  --output /tmp/PointGeometryMain.js
node /tmp/PointGeometryMain.js
# All print the same 26 lines; Oak-only Java source syntax remains smaller.
```

`--java-backend oak` instead compiles a **single, smaller Java subset** entirely in Oak. `--java-backend javac` uses the configured JDK **only at conversion time** with `-proc:none`, a restricted sourcepath/classpath and explicitly staged helpers. `auto` selects `javac` if available, otherwise the Oak subset; it does not promise arbitrary Java acceptance. Neither frontend makes the emitted Oak/JS/HTML depend on a JVM. Use OS isolation for untrusted inputs: the resource and instruction limits are **not** a security sandbox. See the [applet/Builder workflow](applet.md) for targets, round-trips, JAR edits and specific compatibility diagnostics.
