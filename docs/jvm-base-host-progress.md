# Bounded Applet URL host metadata

**Later update:** the subsequent [JDK-checked bounded legacy String.getBytes slice](jvm-legacy-getbytes-progress.md) moved the unchanged original DS selected preflight from the byte-65 barrier described below to `_mth2c()V` byte **759**, `String.<init>([BI)V`. The original worker still does not execute.

**Status: one independently tested JVM host slice, not a running DS applet or full URL support.** The earlier original `DS_BlobsMenu.run()` selected-entry preflight stopped at transitive `_mth2c()V` byte 6, `java/net/URL.getHost()`. The JVM host now admits that exact instance descriptor for its already-existing opaque Applet document/code-base tokens and safe relative-resource tokens. The unchanged original selected entry now stops **before construction** at `_mth2c()V` byte **65**, `java/lang/String.getBytes(II[BI)V`. This is a changed *static first refusal*, not original DS worker or frame execution; the original class's separate static inventory remains **24/32 decoded bodies and nine first-blocker entries**.

## Contract and boundaries

`app.run` can explicitly supply `documentBaseHost` and `codeBaseHost` independently. Each must be a string of 0–253 ASCII characters drawn from letters, digits, dot and hyphen. Empty text models a file-like host. The host is an **explicit metadata snapshot**, not DNS resolution, a general URL parser or permission to use the network. Missing configuration is refused on `getHost`; invalid configuration is refused before guest execution. A safe relative-resource token inherits its base token's configured host, while the existing distinct in-memory resource maps and their restrictions remain unchanged. Calls on the same token retain the same returned String reference in the tested profile.

This does not parse or validate DNS names, implement IPv6/IDN/Unicode URL rules, arbitrary URL constructors, host-based access control, network access, general AppletStub semantics or generated-Java URL APIs. Generated Java explicitly refuses this selected URL program before source emission. Browser evidence covers a **real offline Chrome/Linux run** of the owned fixture, not general browsers, DS pixels, network or input.

## Evidence

[`applet-base-host-oracle.sh`](../applet-base-host-oracle.sh) compiles the owned Java 8 fixture and JDK 11/Xvfb AppletStub reference. Two independent document/code host settings, including empty and case-preserved values and safe relative URLs, matched JDK on native Magnolia. One profile also matched standalone Oak, generated Node JS and actual Chrome output. Invalid/unconfigured hosts were refused, generated Java left no artifact, and the SHA-pinned external original was checked before JDK or guest execution. A separately run same-length edited original was rejected before execution and produced no report.

[Machine report](jvm-base-host.json), SHA-256 `1eb8289c778bb2e20df1e5b2fa10806ea7bb14be44b753dad912524dada3f609`; Magnolia executable SHA-256 `3377c1eb38b149cce84fc85bc708f19cba5962bd312d4bbeeb05ede4e3bd101e`. The original SHA-256 is `d3cbfdbcc7514506e7935f91ba1b40c59c2381a01b344f965da7eaf6ec7a8f25`; it is not included here.

The original `_mth4b` color-method JDK/native/Oak/Node/simulated-HTML regression exited 0 after the change; the separate [original CITY finite raster regression](web-original-fliplay-city-base-host.json) also matched JDK/Node/real Chrome. Fresh native and bundled Node applet tests each **exited 0 at 531/531** ([native log](jvm-base-host-native-suite.log), [Node log](jvm-base-host-node-suite.log)). The initially concurrent base-resource and malformed-URL oracle runs **timed out (exit 124)** and are not passes; subsequent isolated reruns of **both exited 0**, including the updated original DS refusal. `go test ./...`, Python/shell syntax and `git diff --check` passed.

## Reproduce

Requires JDK 11, Xvfb, Node, real Chrome and the separately acquired SHA-pinned external class. The report path must not exist.

```sh
go build -o /var/tmp/magnolia-host .
MAGNOLIA=/var/tmp/magnolia-host CHROME=/path/to/chrome \
  bash applet-base-host-oracle.sh /external/DS_BlobsMenu.class \
  /var/tmp/base-host-new.json
```

Next DS work is an independent JDK oracle for the observed legacy `String.getBytes(II[BI)V` semantics and its explicit byte/string limits, followed by the next selected barriers. Full initialization, graphics/worker scheduling, original pixels and interactions must be checked independently before claiming the applet's intended behavior works.
