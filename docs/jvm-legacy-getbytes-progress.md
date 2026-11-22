# Legacy JVM String byte-copy: bounded DS preflight progress

**Status: an owned JVM fixture passes; the original DS applet is not running.** The original SHA-pinned `DS_BlobsMenu` selected `run()` previously refused before construction at transitive `_mth2c()` byte 65, `String.getBytes(II[BI)V`. Magnolia now admits this exact deprecated method for **ASCII strings only**. Original selected preflight moves to byte **759**, `String.<init>([BI)V`. The separate all-method static inventory remains **24/32 decoded bodies, nine first-blocker entries**. No original DS worker, image, input, or frame ran.

## Contract

The supported method copies the selected ASCII code units into a JVM `byte[]`, with an upper bound of 4,096 source bytes. The JDK 11-checked subset preserves source-range checks before null-destination checks, throws catchable `StringIndexOutOfBoundsException` for tested source/destination bounds and `NullPointerException` for a null destination, and leaves tested destinations unchanged on a failed bounds check. Scan work and the whole copy are charged before destination mutation. This is **not** a complete Java String model: Java's actual method writes the low byte of each UTF-16 code unit; non-ASCII text is deliberately refused instead of being misinterpreted as UTF-8 bytes. Other `getBytes` overloads remain unsupported. Generated Java explicitly refuses this method before source emission.

## Evidence

[`applet-legacy-getbytes-oracle.py`](../applet-legacy-getbytes-oracle.py) freshly compiles four owned Java 8 fixtures and runs JDK 11 `-Xverify:all`. The **40-line** positive/exception trace matched native Magnolia, standalone Oak, Node JS and offline HTML in a **simulated DOM**. A valid JDK non-ASCII example returns `-23` but is explicitly refused by native and Node. Two generated JDK-valid source fixtures verify that native and Node accept the 4,096-byte limit and explicitly refuse 4,097 bytes. Wrong-overload attempts refuse before all four output targets; generated Java refuses the supported Oak/JS method before emission. Separate low budgets exercise both scan and copy refusal on native and Node without guest output. A same-length modified copy of the external archive was separately refused before guest execution/report creation.

The oracle verifies the external archive, nested JAR and class hashes/CRC before the original selected preflight; the unchanged class's SHA-256 is `d3cbfdbcc7514506e7935f91ba1b40c59c2381a01b344f965da7eaf6ec7a8f25`. The [expanded machine report](jvm-legacy-getbytes-boundaries.json) has SHA-256 `c51bed10fa08ee529430d38aa74b63f8e8b3159ef8d4d55181718e6a33a7cd5e`; the tested executable SHA-256 is `d2b00d4864fdd9af50adc04166c26a84c15e29555ffb76915ebcd2f5e4da9b34`.

Fresh native and bundled Node applet suites: **531/531 each** ([native](jvm-getbytes-native-suite.log), [Node](jvm-getbytes-node-suite.log)). The owned MalformedURLException JDK/native/Oak/JS/generated-Java oracle and updated original-DS preflight also passed. Other updated historical DS scripts were not rerun: this environment lacks Xvfb/Chrome. One concurrent `go test ./...` run failed `TestAsyncEventBusAsyncEmit`: it observed `[1, 2, 1]` rather than the test's expected immediate `[1, 0, -1]`. An isolated full `go test ./...` rerun and an initial 50-test repetition passed; a **second 50-test repetition failed three times** with the same unexpected state. This remains an open scheduling/ordering issue, not proof of deterministic behavior. No real-browser check was run for this new slice in the current environment.

## Reproduce

JDK 11, Node, Go and a separately obtained copy of the pinned 26,093-byte DS archive are needed; the tool does not download originals. The report path must be new.

```sh
go build -o /var/tmp/magnolia-getbytes .
python3 applet-legacy-getbytes-oracle.py \
  --magnolia /var/tmp/magnolia-getbytes \
  --ds-zip /external/DS_BlobsMenu_10J.zip \
  --output /var/tmp/jvm-legacy-getbytes-new.json
```

The next observed DS preflight issue is the distinct deprecated `String(byte[], int hibyte)` constructor, which needs its own JDK differential and bounded UTF-16 policy. Clearing that first refusal alone will not establish DS lifecycle or rendering compatibility.
