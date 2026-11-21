# Explicit bounded multi-class ZIP32 selection — JDK 11 parity, not Jmol/Minecraft

`lib/applet-zip.oak` now accepts **2–8 explicitly named** classes from one local ZIP32/JAR with a **1 MiB aggregate extracted-class budget checked against central-directory size before each member's inflation**. Each selected member retains the one-class reader's 64 MiB archive, 50,000-entry/8 MiB index, safe path, local-vs-central metadata, CRC, compression ratio, parsed internal class name and version checks. The original eager loader still rejects archives over 4,096 entries; the separate one-class reader retains its 8 MiB per-class ceiling. The new route stages **neither unrequested classes nor resources/manifests**, never embeds the original JAR, and refuses extra unbudgeted `classFiles`, duplicate/absent names, wrong selected entry and mixed `selectClass`/`selectClasses`. Java classpath/linker preflight still rejects missing *reachable* dependencies **before program output**. This is explicit dependency *selection*, not automatic transitive closure, implicit fetching, whole-JAR signature validation or arbitrary untrusted-bytecode isolation.

The independently executed [`multiclass oracle`](../applet-selective-multiclass-oracle.py) produced the [machine report](jvm-selective-multiclass.json), SHA-256 `567ee73ecc8a229d2f1c7b6c8582ea821637ec0f7b75f165af5e4d0e3af7f683`. A **five-class owned interface/JAR fixture** and a temporary archive containing the **three unchanged external, SHA-pinned original** `ArcTest`, `ArcCanvas` and `ArcControls` classes plus an ephemeral Java 8 harness both matched independent **JDK 11 `-Xverify:all`** output on Magnolia native Oak, standalone Oak and bundled JS/Node. Original Arc lifecycle ran under **Xvfb JDK** and Magnolia on this selected four-class JAR; this is **detached state and four direct lifecycle calls**, not attached Canvas pixels, event delivery or general AWT. Every selected class restored byte-for-byte; restoration of the *whole* original JAR correctly refused. **Eleven negative cases** covered missing reachable subclass, duplicates, nonexistent owner, >8 owners, false entry, mixed flags, unbudgeted classfiles, duplicate archive member, a deceptive helper with the wrong internal class name, default whole-archive cap, and the aggregate cap. The prior single-class 4,101-entry/adversarial ZIP32 regression separately exited 0 again. `go test ./...` also passed at this checkpoint.

**Original Jmol** `JmolApplet0.jar` (SHA-256 `58ab0d0917091a383fb136cc73354ff7fc7187911c8d45c92c2e5bc7f30a7d83`, outside the repository) now has an independent **selected preflight barrier sequence**: with `JmolApplet` alone, `entry has no supported Applet lifecycle`; with its actual `AppletWrapper` superclass selected, the inherited `init()` reaches missing `WrappedAppletLoader` at byte 7; selecting it reaches missing `ClassPreloader` at byte 34; adding that class reaches unsupported `org/jmol/i18n/GT._(String)` at constructor byte 12; and selecting the **fifth** pinned class `GT` reaches the next **unmodeled `java/lang/Object.getClass()Ljava/lang/Class;`** at `AppletWrapper.<init>` byte 51. All five attempts refuse **before constructing an Applet or output**. Independent JDK 11 `javap` confirmed the original `JmolApplet extends AppletWrapper extends java.applet.Applet` hierarchy without running its initializers. This is *measurable first-barrier movement* for a separately pinned original, **not execution of Jmol**, verified resource/362-class dependency closure, molecule rendering or Java 25 Minecraft support. The old one-class Jmol **2/26-body static** inventory remains separate. Full dependency graph and real applet host remain open.

Reproduce offline with separately acquired pinned external Arc originals and an optional pinned Jmol archive:

```sh
python3 applet-selective-multiclass-oracle.py \
  --magnolia /path/to/magnolia \
  --cache /external/sha-verified/ArcTest-siblings \
  --jmol /external/JmolApplet0.jar \
  --output jvm-selective-multiclass.json
magnolia applet.oak run --entry test/fixtures/InterfaceMain.jar \
  --select-classes InterfaceMain,Named,Tuned,BaseValue,ChangedValue
magnolia build --entry test/fixtures/InterfaceMain.jar \
  --select-classes InterfaceMain,Named,Tuned,BaseValue,ChangedValue \
  --output /tmp/selected-interface.js
```

The opt-in archive selection is **not** a general JDK/Java 25 verifier, signed-JAR authoring, arbitrary monitor/GUI/audio compatibility, a Jmol completion claim, or closure of the broader whole-archive AppletStub/Context and original-program TODO entries. Exact positive and negative output is checked on each scoped route.
