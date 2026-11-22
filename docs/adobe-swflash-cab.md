# Signed Adobe `swflash.cab`: what it can and cannot supply

**This CAB is useful evidence, not a playable deliverable.** The locally acquired
`swflash.cab` is a timestamp-signed Adobe Flash **ActiveX installer cabinet**.
It identifies a real 2020 Windows installer route and can be used as a
cryptographically checked lead when planning a *separate*, isolated Windows
reference environment. It does not supply the 1998 Macromedia **Java Flash 2
Applet** with an AWT implementation; it cannot be loaded as a JVM class or a
SWF in Magnolia. Nothing from this CAB was installed or executed.

## Offline findings

- CAB: **230,624 bytes**, SHA-256
  `eea074db3f86fed73a36d9e6c734af8080a4d2364e817eecd5cb37cb9ec9dc0b`.
  A validated `MSCF` header reports one LZX folder, 24 compressed blocks,
  222,832 cabinet bytes and a 7,792-byte PKCS#7 signature overlay.
- The *only* CAB entries, independently decompressed with `cabextract` 1.11,
  are `swflash64.inf` (**218 bytes**, SHA-256
  `97ac49c33b06efc45061441a392a55f04548ee47dc48aa8a916de8d13dabec58`)
  and `FP_AX_CAB_INSTALLER64.exe` (**775,224 bytes**, SHA-256
  `c8c8cff5dc0a3f205e59f0bbfe30b6ade490c10b9ecc7043f264ec67ef9b6a11`).
  The INF contains `run=%EXTRACT_DIR%\FP_AX_CAB_INSTALLER64.exe`: loading
  the CAB in its intended host would **launch** the installer. Our inspection
  did not follow that directive. There is no separately packaged `.ocx` or
  `.swf` **member**. This does *not* rule out data inside the installer binary.
- Despite its `64` filename suffix, the extracted installer has an **x86
  PE32** header. Its UTF-16 version string is **`32.0.0.465`**; static strings
  include Flash Player/ActiveX install and remote update references. An
  installer or downloader is not equivalent to a working player, and the
  presence of a remote URL is not evidence that a 2026 server still serves it.
- A second **offline PE layout/resource inventory** found six mapped sections
  (`.text`, `.rdata`, `.data`, `.gfids`, `.rsrc`, `.reloc`) and **124 resource
  leaves**: seven icons, 112 strings, a group icon, version and manifest,
  plus named `TYPELIB` (**17,892 stored bytes**) and `LZMG` (**9,489 stored
  bytes**) leaves. The EXE's 7,736-byte certificate directory follows the
  final section and ends at **EOF**, with no *additional appended file
  overlay*. The hash-pinned `LZMG` bytes are **opaque** here: they were not
  decompressed or executed. No resource layout or absent appended overlay can
  prove that arbitrary code/resources cannot contain or fetch a plugin. In
  particular, we do **not** claim to have extracted an OCX from the EXE.
- A further **read-only inspection of those exact signed resource bytes**
  found that the 17,892-byte `TYPELIB` leaf begins with the `MSFT` type-library
  magic and contains raw ASCII markers `FlashBrokerLib`, `IFlashBroker`,
  `IFlashBroker6`, `BrokerCreateFile`, `BrokerLMOpenDownload`, `BrokerLMLaunch`,
  `BrokerMarkFileUnsafe` and `BrokerPathToBitmapFill`. The machine-readable
  report records their byte offsets *within the pinned resource*. These are
  broker-labelled **metadata strings**, not executed interfaces, proof of a
  resident plugin, or verified functionality. The 9,489-byte `LZMG` leaf's
  first 16 stored bytes are also recorded, but its format/content remains
  opaque: no decoder, installer, COM registration or network request ran.
- `osslsigncode` 2.9 verified both CAB and EXE payload digests, an Adobe Inc.
  / Flash Player code-signing certificate chain to the local trust store,
  and their 2020 timestamps **offline**. The CAB signed digest was
  `a5ac2338f161f770097561f7a13930745f3bdb86ccc5bf66ad37fcdd22c1eb18`;
  the extracted EXE's was
  `2990999959124664ae302e9fd6f69771879590d91645f87f93eea91330407bdd`.
  **Both checks returned exit 0 with `-ignore-crl -ignore-cdp`.** Online
  certificate revocation was *not* checked; neither this nor a valid
  historical signature establishes present-day installer safety or permission
  to redistribute it. A one-byte-changed CAB failed the pin/signature test.
- The HTTP request for
  `http://download.macromedia.com/pub/shockwave/cabs/flash/swflash.cab`
  redirected once over **unauthenticated HTTP** to `fpdownload2.macromedia.com`;
  HTTPS to that host failed certificate hostname verification. The Adobe
  signature authenticates the checked **file content**, not the HTTP response,
  the redirection, any future download the installer might make, or a claim
  that installing obsolete ActiveX is safe.

[Machine-readable local review](adobe-swflash-cab-review.json) and
[reproducible bounded inspector](../adobe-cab-review.py) keep the two
third-party binaries **outside the repository**. The inspector requires an
explicit CAB path plus trusted, locally supplied `cabextract` and
`osslsigncode` executables. It checks sizes, member names/hashes, CAB layout,
PE format, bounded section/resource layout (including the opaque LZMG leaf
and the `MSFT` TYPELIB marker bytes), INF directive and both signatures before
writing a report. A repeat offline
run produced a **byte-identical** JSON report; in-memory changes to the PE
resource RVA, certificate size and section offset were each rejected by the
bounded inventory reader. These synthetic probes did not bypass the normal
CAB/EXE SHA pins. It extracts to a disposable temporary directory; it never
launches the `.exe` or decompresses its opaque named resources.
For example, using separately obtained distro tools:

```sh
LD_LIBRARY_PATH=/path/to/extracted-deb-usr/lib/x86_64-linux-gnu \
python3 adobe-cab-review.py \
  --cab /path/to/swflash.cab \
  --cabextract /path/to/extracted-deb-usr/bin/cabextract \
  --osslsigncode /path/to/extracted-deb-usr/bin/osslsigncode \
  --report adobe-swflash-cab-review.json
```

The 2020 Windows **ActiveX** installation route is distinct from the archived
Flash 2 **Java** Applet (`FlashApi.ZIP`), a Windows **projector** executable,
Pepper **PPAPI** plugin, and the uploaded **SWF5** game. Using the CAB for
actual reference playback would require a separately validated Windows/IE
host and the installed plugin, compatible SWF input, explicit consent to run
obsolete software in isolation, and independent frames/input/audio tests.
None was attempted here. The separately archived standalone projector is
another possible Windows reference path, *not* evidence that the CAB itself
plays the user's movie. In Magnolia, actual playback remains unimplemented.
