# Web compatibility: byte boundaries, loopback HTTP and original movie

**Status: improved, not full compatibility.** The native Oak string model is byte-oriented. The prior JS generator embedded literal high bytes directly in a UTF-8 JavaScript file, so `00 7f 80 ff` became `[0,127,65533,65533]` under Node. Generated literals now escape non-ASCII and other unsafe bytes as ASCII `\\xNN`, keeping one Oak byte per JS code unit. The separately tested ordinary UTF-8 text remains correctly printable.

The associated runtime changes preserve byte streams on Node stdout and filesystem paths; convert host Unicode arguments, environment values and browser prompt input into UTF-8 byte strings; decode Oak UTF-8 for browser console presentation and request URLs; send fetch bodies as bytes; and make `utf16()` use Go-compatible decoding, including BOM, supplementary characters and malformed/truncated input. Node subprocess arguments are converted as host text and stdin remains bytes. Embedded browser VFS listing now handles the `.` root and child directories; browser `exit` halts the Oak invocation using a tagged `OakExit` exception that hosts can catch, rather than silently continuing. The JS `std` module now exposes the native `keys` and `len` builtins needed by the self-hosted web builder.

The HTTP runtime now uses manual redirects on Node async fetch, avoids launching an extra async request for synchronous Node calls, preserves response bytes in the synchronous child-process fallback and encodes Unicode URL paths. Browser fetch refuses opaque redirects rather than following them; that refusal is **not** equivalent to a native 302 response.

A further browser-side fix ensures a synchronous request (unsupported in this host) returns an error **before** starting `fetch`. Before the fix, the extended real-loopback negative control observed both an OPTIONS preflight and a POST despite the returned error; after the fix, neither request occurs. Browser synchronous HTTP remains unsupported rather than silently made equivalent to native synchronous HTTP.

A further request-boundary fix validates absolute HTTP(S) URLs before any JS request. Previously Node's async `fetch` accepted a `data:` URL (returning 200 where native returned an error), while its synchronous fallback sent an `ftp:` URL as HTTP to the loopback server (returning 404 rather than rejecting it). The extended oracle first reproduced both failures, then confirmed `data:`, `ftp:` and a malformed URL return errors on native and Node sync/async and Chrome async; the FTP target receives no request. This is an HTTP transport boundary, **not** a general Java URL implementation or proof of every URL parsing edge case.

A further Node request fix handles a bounded URL-credential case: the synchronous fallback previously omitted native's Basic authorization, while Node fetch rejected credential-bearing URLs. Node now converts UTF-8 URL userinfo to Basic authorization, respects an explicit `Authorization` header case-insensitively, and removes the userinfo from the fetch URL. Non-UTF-8 userinfo without an explicit overriding header is deliberately refused before network access, unlike native; actual Chrome refuses credential URLs. **Basic over plain HTTP is not confidential**; the oracle uses dummy credentials on loopback only, not real credentials or TLS.

## Evidence

- [`web-http-auth-oracle.py`](../web-http-auth-oracle.py) passed **20** native/Node sync/async dummy-credential cases, including UTF-8 and empty userinfo, explicit header precedence, and deliberate no-network refusal for non-UTF-8 Node userinfo; actual Chrome refused a credential URL without sending a request. The prior build failed the Node sync Basic-auth control. [Report](web-http-auth-extended.json), SHA-256 `24e3d379aa1f0be160b70943b16cd4ed242ad3cb05b9e3c33e6270b134236410`. The same new executable (SHA-256 `8a82b674555804cae2371453a3989f1b2bba1898ccc1f52e303bb6cd331dea9e`) passed the [loopback](web-http-auth-loopback-regression.json), [byte](web-byte-compat-auth.json), and [original CITY](web-original-fliplay-city-auth.json) regressions; Node suites passed **21/21**, **405/405**, and **531/531**. `go test ./...` and `git diff --check` passed. This is a narrow HTTP profile, not browser credential parity or general authentication support.

- [Extended HTTP scheme report](web-http-schemes-final.json), SHA-256 `98f730207d7a2988cd21b3f03a7631a5d3fc35b051d4cae797b606c70b26d3b0`, includes the new negative controls plus the real-loopback HTTP, CORS and redirect checks. It and the passing [byte regression](web-byte-compat-scheme-validation.json) use the rebuilt executable SHA-256 `820e6e843e485347868ac04f185089b79a36148fe29d1d4a1c4dab604fcb0a4b`. On this same executable the rebuilt Node suites passed **21/21**, **405/405** and **531/531**, and the [original CITY raster comparison](web-original-fliplay-city-schemes.json) again passed JDK/Node/Chrome. The scheme-negative tests are limited to these three URLs and one Chrome/Linux version; TLS, authentication beyond the selected loopback Basic/override cases above, cookies and general header/URL parity remain unverified.

- The extended loopback oracle reproduced the unwanted browser sync network side effect on the previous build and passed its no-OPTIONS/no-POST control after the fix. [New report](web-http-sync-refusal.json), SHA-256 `8a13f8e3ae240c5be899fc941872000007818a8e2b1119d363d9c544922fd8b0`. With the post-fix executable (SHA-256 `744ed2f6fa86e538b66f8d415060419207448b22ff7478d1ea255f55da3a3adc`), the [byte/browser regression](web-byte-compat-sync-refusal.json) and [original CITY JDK/Node/Chrome raster regression](web-original-fliplay-city-sync-refusal.json) also passed; fresh Node suites passed builder **21/21**, audio/core **405/405** and applet **531/531**. The fresh native applet suite also **exited 0 at 531/531** ([log](web-native-applet-sync-refusal.log), SHA-256 `fa13b016902af70dd6c212a4f35bfae7415f5a05ea4b9c8ffbc9b2ccf9c4e916`). `go test ./...` and `git diff --check` passed. These do not upgrade the documented scope of those checks.

- [`web-original-fliplay-city.py`](../web-original-fliplay-city.py) executed unchanged SHA-pinned original Fliplay classes and CITY.FLI through the generated web VFS. On a finite headless lifecycle, Node and real Chrome each matched four fresh JDK 11 RGB state hashes (three distinct images), eight repaints, six status events and the final 256,000-byte ARGB hash. The same-length altered movie was refused before guest execution. [Report](web-original-fliplay-city.json), SHA-256 `c8abd8ac684347c5e593fb200ebaef107ec2adfba5fa1894eac1f544a928aee4`. This is a software-raster hash check in one original applet profile, not a canvas display, interactive playback or audio test; the owned caller was adapted to VFS paths.
- [`web-http-loopback-oracle.py`](../web-http-loopback-oracle.py) exercised two private loopback HTTP servers: native and Node sync/async and real Chrome async received matching status 201 and binary response bytes, while the server observed the exact binary POST body, probe header and Unicode percent-encoded URL path. Native and Node returned the 302 without following; Chrome reported an opaque-redirect error without following. A denied CORS preflight blocked the POST. [Report](web-http-loopback.json), SHA-256 `9e1eebb67df400b174165bbf0a53294bdbb9a02af5ec838a35563563cfe5ff0a`. This does not validate TLS, external servers, cookies, streaming, general headers or other browsers.
- After the HTTP changes, the byte/fake-fetch oracle and all three Node suites below were rerun successfully with the same newly built executable used for the two reports above. [Byte regression report](web-byte-compat-http2.json), SHA-256 `7e639e7970c477ffd9f66c58a0ba7b85db3bf6ae84cf5736e833638b3e7d449f`; executable SHA-256 `35465bc61ea0a33b8020198df4dd96cddcf6f4aa0520706938b5941d8cfeadb2`.

- [`web-byte-compat-oracle.py`](../web-byte-compat-oracle.py) exited 0: exact native/Node **256-byte literal vector and 256-byte raw stdout**, valid UTF-8 text, UTF-16 BOM/supplementary/malformed cases, Unicode file/directory names, argv/environment and local subprocess bytes. A real headless **Chrome 142 on Linux** also checked the 256-value vector, visible Unicode, a controlled prompt, embedded VFS bytes and Unicode directory names, and tagged exit termination. A **fake fetch** in Node and Chrome checked a Unicode URL and exact binary request/response bytes; this does not test actual remote HTTP, CORS or servers. [Report](web-byte-compat-current.json), SHA-256 `510a85b5037fe0503a0ed4d9892d5927c431ba9f7233a2141677679b9bd19523`.
- [`web-applet-html-smoke.py`](../web-applet-html-smoke.py) compiled a fresh JDK 11 reference and loaded generated offline HTML in **real Chrome**: all seven selected owned RGB-model method output lines matched. This is **not** a canvas pixel, original applet, sound or input gate. [Report](web-applet-html-smoke.json), SHA-256 `bb93e6dc811b8100d2c63e0f7e9e0944308ce0dcd87afd070698f5d50a137d1f`.
- The original uploaded SWF5's SHA-pinned 32-byte button action, encoded as an Oak byte literal, produced the same six-record virtual trace in native, Node and **real Chrome**. The legacy host refused it as expected. [Report](web-original-button-literal.json), SHA-256 `3b60aa8fd4f96cdcfcdea5c76ad8b826d916240db5a63ffd262a295c87b64873`. This is a direct synthetic action, not movie playback or a physical click.
- On freshly built web bundles run under **Node**, applet fixtures passed **531/531**, audio/core fixtures **405/405**, and self-hosted builder fixtures **21/21**. Native builder fixtures passed **21/21**; native applet fixtures passed **531/531** on the immediately preceding build, before the last JS-only VFS/exit changes. These Node suites are not full-browser execution. `go test ./...` and `git diff --check` passed.

## Reproduce

Provide an installed Chrome/Chromium binary and JDK 11 at the path used by the applet oracle. Output paths must be new. The examples use the pinned Chrome for Testing binary from this workspace; Node is also required.

```sh
go build -o /var/tmp/magnolia-web .
python3 web-byte-compat-oracle.py \
  --magnolia /var/tmp/magnolia-web \
  --chrome /var/tmp/oak-chrome/chrome-linux64/chrome \
  --output /var/tmp/web-byte-compat-new.json
python3 web-http-loopback-oracle.py \
  --magnolia /var/tmp/magnolia-web \
  --chrome /var/tmp/oak-chrome/chrome-linux64/chrome \
  --output /var/tmp/web-http-new.json
python3 web-http-auth-oracle.py \
  --magnolia /var/tmp/magnolia-web \
  --chrome /var/tmp/oak-chrome/chrome-linux64/chrome \
  --output /var/tmp/web-auth-new.json
python3 web-original-fliplay-city.py \
  --fliplay /var/tmp/magnolia-gemalogo-probe/classes/fliplay.class \
  --flickframe /var/tmp/magnolia-gemalogo-probe/classes/flickframe.class \
  --movie /var/tmp/magnolia-fliplay-new-originals/city.fli \
  --magnolia /var/tmp/magnolia-web \
  --chrome /var/tmp/oak-chrome/chrome-linux64/chrome \
  --output /var/tmp/web-city-new.json
python3 web-applet-html-smoke.py \
  --magnolia /var/tmp/magnolia-web \
  --chrome /var/tmp/oak-chrome/chrome-linux64/chrome \
  --output /var/tmp/web-applet-html-new.json
python3 web-original-button-literal.py \
  --swf /home/user/uploads/asdf.swf.txt \
  --magnolia /var/tmp/magnolia-web \
  --chrome /var/tmp/oak-chrome/chrome-linux64/chrome \
  --output /var/tmp/web-original-button-new.json
for suite in WebBuildSuite WebAudioStdSuite WebAppletSuite; do
  /var/tmp/magnolia-web build --entry "test/fixtures/$suite.oak" \
    --web --output "/var/tmp/$suite.js"
  node "/var/tmp/$suite.js"
done
```

## Remaining compatibility work

Full compatibility cannot be asserted from these slices. Required work includes real-browser original-program interactive playback and canvas pixel/input/scheduling/audio comparisons, DOM and file/network byte-to-text boundaries beyond the exercised paths, HTTP/TLS and CORS beyond the controlled loopback cases, browser redirect parity and browser storage, unsupported raw-socket behavior, keyboard/pointer/accessibility and autoplay policies, security/CSP behavior, performance/resource limits, and Firefox/Safari/Windows/mobile coverage. Browser `print` is console text presentation, not an arbitrary binary output API; output split across calls is not proven as a streaming UTF-8 decoder. JS runtime host text beyond the explicitly tested boundaries and non-UTF-8 OS subprocess arguments remain unverified. Existing JVM/AVM1/AVM2 and MIDlet limitations remain, including the original SWF5 and Acotel whole-program gates. The broad acceptance record is still **21 JVM / 24 ActionScript unchecked**.
