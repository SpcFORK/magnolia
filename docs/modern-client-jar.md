# Bounded Java 25 client-JAR investigation (not client execution)

**Input:** the exact Mojang URL supplied for this investigation, `https://piston-data.mojang.com/v1/objects/e877b6a07acd633fb3bb475002175cec036e7b87/client.jar`. Downloaded outside the repository; URL SHA-1 `e877b6a07acd633fb3bb475002175cec036e7b87` matched the bytes. Size: 41,483,720 bytes; SHA-256: `4508d006323f24fa02876310c192d739af56516eb259000ac50f0909a68c9a2d`. Do **not** add this third-party JAR or its classes to the source tree, examples, or patch.

Static ZIP inventory: **34,227 entries**, including **11,383 `.class` files**, 22,844 other entries and **zero `.java` source files**; 59,950,721 uncompressed class bytes. Every inspected class header reports **JVM 69.0 (Java 25)**. The manifest declares `Main-Class: net.minecraft.client.Main`, but that exact class path was not present; `net/minecraft/client/main/Main.class` was present. Neither a manifest nor a decoded body proves the declared application is runnable.

Magnolia's ordinary `lib/applet-zip.oak` whole-archive loader rejects this JAR: it caps the archive at **4,096 ZIP entries**, **256 classes** and **32 MiB total extracted class bytes**. Earlier bounded extraction of six CRC-checked classes used Python's standard ZIP reader *only for selecting input bytes*; Magnolia's own `inspect` and `coverage` processed the resulting original class bytes. In addition, the Oak `selectClass` route now reads central-directory metadata without extracting the other entries: at most **64 MiB input**, **50,000 entries**, **8 MiB central directory**, and **8 MiB for the one selected class**, with checked path, selected local header, ZIP32 size, 256:1 expansion ratio, DEFLATE/stored payload and CRC32. No remote classpath, resources or original whole JAR are embedded. This is **one-class selection**, not lazy linking, full ZIP integrity verification of unused members, or full client execution. Validate untrusted inputs under OS isolation; these budgets are not a security sandbox.

| Selected original class | Methods | Decoded bodies | Coverage blockers | Observed first boundaries |
| --- | ---: | ---: | ---: | --- |
| `package-info` | 0 | 0 | 0 | Metadata-only. |
| `Blaze3D` | 7 | 1 | 10 | External SLF4J/LWJGL dependencies and a double descriptor. |
| `ProjectionType$LayeringTransform` | 1 | 0 | 1 | Float descriptor unsupported. |
| `net/minecraft/client/ClientBrandRetriever` | 2 | 2 | 0 | `getClientModName()Ljava/lang/String;` returned `vanilla`. |
| `net/minecraft/client/main/Main` | 12 | 1 | 18 | Exception handlers block its `main` method. |
| `net/minecraft/client/Minecraft` | 238 | 82 | 389 | Missing superclass/interface/external API types; decoded bodies are not a linkable client. |

The original `ClientBrandRetriever.getClientModName()` returned `vanilla` under **Temurin JDK 25.0.5 with `-Xverify:all`** and matched Magnolia's native Oak VM when selected directly from the Mojang JAR. A small original-class extraction then produced the same output in standalone Oak, bundled JS/Node and generated Java; the standalone Oak archive restored that **class byte for byte**. This is now a reproducible **single-method Java 25 behavioral oracle**, not a Java 25 JVM-wide verification result. A package-info class with no methods has nothing to execute. The client entrypoint and render/game services were **not executed**, and the presence of 82 decoded Minecraft bodies does not imply client compatibility.

To reproduce selection **without importing the whole client as an executable JAR** (the download belongs outside the checkout):

```sh
mkdir -p /tmp/magnolia-client
curl --fail --location --max-time 120 \
  'https://piston-data.mojang.com/v1/objects/e877b6a07acd633fb3bb475002175cec036e7b87/client.jar' \
  --output /tmp/magnolia-client/client.jar
printf '%s  %s\n' e877b6a07acd633fb3bb475002175cec036e7b87 /tmp/magnolia-client/client.jar | sha1sum -c -
magnolia applet.oak inspect --entry /tmp/magnolia-client/client.jar \
  --select-class net/minecraft/client/ClientBrandRetriever
magnolia applet.oak coverage --entry /tmp/magnolia-client/client.jar \
  --select-class net/minecraft/client/main/Main --json  # blockers: nonzero exit
magnolia applet.oak run --entry /tmp/magnolia-client/client.jar \
  --select-class net/minecraft/client/ClientBrandRetriever \
  --method getClientModName --desc '()Ljava/lang/String;'
```

With the JAR outside the repository and a separately installed **JDK 25** (Temurin 25.0.5+8 Linux x64 archive SHA-256 `68982529869257d0ee6b824ec5aa270ff8b118bb438670da0b97ebdab178d7fb` was verified against the vendor's published metadata), run `MAGNOLIA=/path/to/magnolia JDK25_BIN=/path/to/jdk-25/bin bash applet-modern-brand-oracle.sh /tmp/magnolia-client/client.jar`. This script rechecks the exact JAR's SHA-1/SHA-256, compiles `test/modern/BrandOracle.java` against it, runs `java -Xverify:all`, compares the **original method** with direct Oak selective-JAR execution and the three generated interpreters, and verifies class-byte restoration. No compiled Mojang bytecode or JDK is committed. For synthetic adversarial ZIP32 cases and a 4,101-entry archive, run `MAGNOLIA=/path/to/magnolia bash applet-selective-jar-tests.sh`; it checks failures for duplicates, unsafe paths, CRC, size/ratio, ZIP64, local-header name/flag/checksum/size mismatches, encryption, unsupported compression, bad later central entries and inconsistent Builder entry selection.

Programmatically, `import('./lib/jvm-bridge.oak').load({entry: jarPath}, {selectClass: 'net/minecraft/client/ClientBrandRetriever'})` yields **one** validated class and `jar: ?`. `api.inspect`/`api.coverage` accept the same `selectClass` setting. Magnolia Builder accepts `magnolia build --entry JAR --select-class JVM/name --method METHOD --desc DESC --output FILE.oak` (or `.js`/`.html`/`.java`); `config.applet.selectClass` is its programmatic equivalent. An explicit `entryClass` must match the selected owner. Extra helper classes must be supplied deliberately; absent dependencies fail preflight. Selected standalone Oak restores the original selected **class bytes only**, not the original giant JAR; do not infer that it can recompile an edited full JAR.

**Next:** complete the type-aware verifier (including modern frames and exception edges), float/double, observed external APIs and library/host contracts, and build independent Java 25 fixtures/reference oracles for any modern semantic changes. There is no claim here of a full JVM, a functioning Mojang client, or a runtime Java source compiler.
