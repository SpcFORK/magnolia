# Finite historical SDK ADPCM task: real sound payload and SWF writer

**Observed executable result:** [`sdk-adpcm-functional-oracle.py`](../sdk-adpcm-functional-oracle.py) **exited 0**. [`swf-adpcm-encode.py`](../swf-adpcm-encode.py) produced three real FWS1 **DefineSound** movies from *constructed, owned* signed 16-bit PCM: 10-frame mono 22,050 Hz, 12-frame interleaved stereo 44,100 Hz, and 4,098-frame mono 11,025 Hz crossing a real **4,096-sample ADPCM block** with the extended SWF tag header. Each complete 4-bit compressed DefineSound payload was **byte-for-byte identical** to the unchanged original historical SDK `Parser/adpcm.cpp`'s `PutDefineSound`/`PutADPCMData` methods, compiled privately inside a small owned, bounded MSB-first bit-buffer/Windows-type wrapper. The **whole generated FWS1** was then inspected by the *separately compiled original SDK C++ `swfparse.cpp`*: expected version, complete stage/length, frame, sound format/rate/channels/count, ShowFrame and End; on mono, the old parser's audio decode emitted the first sample and **all nine delta samples**. This is a real **PCM → compressed SWF sound asset** task, not just a structural sound-tag counter.

| Owned PCM16 fixture | Frames × channels | Exact original-C++ ADPCM DefineSound payload | Complete generated FWS1 |
| --- | --- | --- | --- |
| mono, 22,050 Hz | 10 × 1 | 15 bytes, SHA-256 `b644017afd6590f7eb89af169326a507b45dc17507fe6c0b2a1dbea11418d415` | 37 bytes, `c10a3f7c4017314a914bfb65967659e7c7f232e49fe6ec9a7ec748e039606062` |
| stereo, 44,100 Hz | 12 × 2 | 24 bytes, `c664c67c584540f1a64d90c7510eeea906e303362fdf79d9a3c3bc89417078a3` | 46 bytes, `230d2b25675cd983a8b3705ce28dc19f306ae9502f3d899e50ef0db54c667803` |
| second 4K block, 11,025 Hz | 4,098 × 1 | 2,061 bytes, `d257610cd13923096e9ff109690b695d7f39c1c60a42f4e38d40c5599c1959c9` | 2,087 bytes, `dcc95324c2d38a4538a23efb8a9da0e8f0903527d44673afbb1531b` |

**Provenance and independent roles:** historical [SDK `Parser/adpcm.cpp`](https://raw.githubusercontent.com/SpcFORK/Macromedia-Flash-SDK/bb8536a8f564510f979eb8e09924bc259651a0a9/Parser/adpcm.cpp) (7,768 bytes, SHA-256 `6e1ed6947a3757431c5b994d842c57bac4e93a6bc009ff1d9d93bfb291a89d15`) is **unmodified external C++**. The [original `Parser/swfparse.cpp`](https://raw.githubusercontent.com/SpcFORK/Macromedia-Flash-SDK/bb8536a8f564510f979eb8e09924bc259651a0a9/Parser/swfparse.cpp) (99,495 bytes, `d06448e8da95c002ddfaeedae0d5160e70c82e3fa82f311aaea16336d007cf30`) needs the same previously verified **portability-only** loop-scope/const-literal build edits. The wrapper in the repository implements buffer primitives and 16-bit `WAVEFORMATEX`-like values solely to compile that original **source fragment**; it is **not** a reconstructed full Windows SDK exporter. Magnolia's independent writer is Python. Neither historical source nor user/third-party PCM is embedded in the deliverables. The old parser is given **only the three precisely generated and SHA-checked small SWFs**, never arbitrary/untrusted media.

The original algorithm reads the **next interleaved sample** to choose a new block index, even when that is the other channel's header. The writer faithfully preserves this original behavior, its sign/index/step/clamp rules and nibble order rather than silently substituting ordinary IMA ADPCM. Safety limits are explicit: **PCM16 little-endian only**, mono/stereo, **exactly** 5,500/11,025/22,050/44,100 Hz, **2–4,200 complete sample frames** and no block ending with the source fragment's out-of-bounds one-sample lookahead. It defines sound ID 1 but **does not issue StartSound or produce a playable timeline/mixer**. Four CLI mutants (bad pin, partial stereo frame, unsafe 4,097-frame block, oversize) each **exited 2 without SWF output**; a separately changed original `adpcm.cpp` SHA **exited 1 before any native build/report output**. Five grouped owned encoder and no-artifact tests passed. Unsupported WAV headers, 8-bit PCM, sample rates, huge files, MP3, ADPCM decode, streaming audio, uploaded SWF5 and audio-device output remain unclaimed.

```sh
cd /home/user/magnolia
# Supply the two separately verified original C++ files OUTSIDE this checkout:
python3 sdk-adpcm-functional-oracle.py \
  --adpcm-source /path/to/original/Parser/adpcm.cpp \
  --parser-source /path/to/original/Parser/swfparse.cpp \
  --output /path/to/a/NEW/reference-report.json
# For rights-cleared raw PCM16 whose source and sample rate you verified:
python3 swf-adpcm-encode.py --pcm /path/to/signed-le16.pcm \
  --channels 1 --rate 22050 --expect-sha256 VERIFIED_PCM_SHA256 \
  --output /path/to/NEW-sound.swf
```

[Metadata-only differential report](sdk-adpcm-functional.json). **Acceptance boundary:** one bounded encoder *source-fragment* task now works and matches historical C++; it does **not** complete the historical SDK's general sound/media utilities, SWF parser, vector-shape production, a reference-audible WAV/SWF player, the Java Flash 2 `sun.audio` blocker, or any ActionScript runtime. A separately [original-C++-matched JPEG bitmap asset writer](sdk-jpeg-functional.md) now produces bounded SWF2 image payloads; it does not display them or turn this sound slice into a complete SDK media utility. The user's **every-media-track completion remains open**.
