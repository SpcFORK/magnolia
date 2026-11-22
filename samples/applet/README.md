# Magnolia Oak-first applet/JVM adapter examples

These examples are generated from the **local test fixtures**, not from an
Adobe or Macromedia player. They are self-contained except that `.oak` needs
the Magnolia interpreter. They demonstrate a *verified subset*, not universal
applet support. The Oak examples were regenerated with both runtime and
invocation integrity markers so the checked zero-edit/edited reverse route can
read them. The earlier Interface, reference-identity and Point examples and the
graphics applet were checked against their JDK/Canvas oracles at the Point
milestone; a later exact-descriptor `System.arraycopy` host added the ArrayCopy
example, verified against its 23-line JDK oracle.

Only the **Oak source** samples are committed here. The `.js`, `.html`,
`.java`, edited `.jar`, `.svg` and `.png` targets are regenerable build
outputs and are **not** checked in — regenerate them with Magnolia Builder
(`magnolia build --entry <sample>.oak --output <target>`) or run the Oak
sample directly.

## Samples

| File | What it demonstrates |
| --- | --- |
| `InterfaceMain.oak` | Readable standalone Oak from a five-class manifest JAR with inherited interfaces, receiver-based `invokeinterface`, `instanceof` and `checkcast` |
| `ReferenceIdentityMain.oak` | Executable `if_acmpeq`/`if_acmpne` sample generated from the [Java 8 fixture](../../test/fixtures/ReferenceIdentityMain.java); equal-but-distinct objects/arrays/strings stay distinct, while aliases, interned literals and `System.out`/AWT Color constants preserve identity |
| `PointGeometryMain.oak` | A [Java 8 fixture](../../test/fixtures/PointGeometryMain.java) for exact AWT `Point.x/y` int fields, constructors/copies, `move`/`setLocation`/`getLocation`, `equals` versus `==`, `Point2D`/`Serializable`/`Cloneable` subtyping and array references. **Not** a complete AWT window or historical player |
| `ArrayCopyMain.oak` | A [Java 8 fixture](../../test/fixtures/ArrayCopyMain.java) for overlap-safe `System.arraycopy`, five primitive array kinds, reference covariance/identity and bounded host work. Matches the JDK's 23-line success oracle; invalid-input cases fail explicitly |
| `GraphicsApplet.oak` | Standalone Oak applet execution, named JVM methods, embedded original-class archive and inlined VM library/runtime; `--paint` records bounded drawing commands |

## Running the Oak samples

From the repository root:

```sh
magnolia samples/applet/InterfaceMain.oak
# Prints 3, 8, 7, true, true, 7 on separate lines.

magnolia samples/applet/ReferenceIdentityMain.oak
# Prints 22 boolean results; compare with:
java -Xverify:all -cp test/fixtures ReferenceIdentityMain same same

magnolia samples/applet/PointGeometryMain.oak
# Prints 26 results; compare with:
java -Xverify:all -cp test/fixtures PointGeometryMain

magnolia samples/applet/ArrayCopyMain.oak
# Prints 23 lines; compare with:
java -Xverify:all -cp test/fixtures ArrayCopyMain

magnolia samples/applet/GraphicsApplet.oak
# Records bounded Canvas drawing commands (--paint).
```

To reproduce an edited JAR, build `test/fixtures/InterfaceMain.jar` to Oak with
Magnolia Builder; in `InterfaceMain.score(LNamed;)I`, change the single literal
`[4, 0, ?]` to `[5, 0, ?]` (after the five-byte `invokeinterface` instruction).
Then `magnolia build --entry EDITED.oak --output NEW.jar`. `restoreOak` and
`--output ORIGINAL.jar` **without edits** preserve the exact original container;
a new edited JAR repacks its class bytes, not the exact ZIP container
representation. Signed/digested edited JARs are rejected.

See [the applet workflow](../../docs/applet.md) and [Oak VM library
contracts](../../docs/jvm-vm-libraries.md). Authentic Java Flash Edition is
**not yet runnable** (the preflight blocker is `dconst_1`); the linked Rapper
Rhythm game EXE ZIP was not obtained and its contents remain unknown.
