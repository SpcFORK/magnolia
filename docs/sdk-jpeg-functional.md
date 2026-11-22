# Finite historical SDK JPEG bitmap asset — not Flash playback

**Executable result:** [`sdk-jpeg-functional-oracle.py`](../sdk-jpeg-functional-oracle.py) **exited 0**. The owned [`swf-jpeg2-encode.py`](../swf-jpeg2-encode.py) takes SHA-pinned, raw RGB24 pixels and creates a complete **FWS2/DefineBitsJPEG2** SWF. Its JPEG2 payloads were **byte-for-byte identical** to three outputs of the **unchanged original historical C++ SDK** `Parser/jpeg.cpp`, privately compiled with libjpeg 6.2 inside a [bounded, owned image/buffer wrapper](../sdk-jpeg-native-wrapper.cpp.txt). The original C++ method `CExportSwf::PutDefineBitsJPEG2` really calls `CSwfJpeg::PutInterchangeBits`, writes two JPEG substreams, and returns a compressed asset; our wrapper implements the surrounding historical type/output-buffer methods and prepares an already-RGB image. It is **not** the complete historical import/export program. No source from the original `jpeg.cpp` is stored in Magnolia.

The older `jpeg.cpp` first calls `jpeg_write_tables`, then suppresses tables in the image stream. The Python writer independently uses Pillow's JPEG encoder at quality 100, baseline 4:2:0, and validates the **exact nine-marker** JFIF/DQT/SOF/DHT/SOS profile before separating its normal interchange JPEG into the historical **tables JPEG + image JPEG** payload. It writes a complete bounded SWF2 header, character ID 1, one extended DefineBitsJPEG2 tag, ShowFrame and End. **No bitmap shape is placed on stage**: the asset is encoded and independently decodable, but this movie does **not** display a bitmap or have a playable animation. JPEG compression is lossy, so decoded RGB hashes are deliberately **not** the original raw-input hashes.

| Constructed owned RGB input | Original C++-identical DefineBitsJPEG2 payload | Whole FWS2 | Independently decoded JPEG |
| --- | --- | --- | --- |
| Flat **4×4** RGB, 48 source bytes | **641 bytes**, SHA-256 `b573ee51207f4813f40eb0dff973cb975d131034af6b0cea3c0fe28e23aaf0a5` | **668 bytes**, `67c024c4ea83800926f27002345347bc43896ab051cdbbefbb248c7adf23c241` | 16 pixels, 1 color; RGB SHA-256 `3ec6dec9f1ac404a48a181ec35a76dbbf4bc78f04fff33deaf8d2bd6fc5a9b5b` |
| Mixed **16×16** RGB, 768 bytes | **1,122 bytes**, `1ec205b8bb7e91a395dd17c8ca0022ebf6119bc2ea22aa0d4ed449a59f9e1f53` | **1,150 bytes**, `1c7e20095a050acfb37b72d4e81151c7139de915e6ea1a99c60ecee3e83ebab9` | 256 pixels, 256 colors; `fbbcb6c428b5d86f416577713d791355e0c806bb1351325096c452040ff5055b` |
| Odd **7×9** RGB, 189 bytes | **854 bytes**, `ef8ad1339905be5c4b112a710302f585ffde72613a56e09f6bbc8f574ec895a3` | **882 bytes**, `1712afd52c9b486f4742f9ea73a336b50c3e2dd52d22386a2991785b304cab71` | 63 pixels, 63 colors; `fc881e48dad7fc32d5ab9f97be9895edd2bd6a63bb588d394bfbfe54af7dac77` |

**Independent checks and safety:** The same **whole, SHA-checked** SWF2 outputs went to the separately compiled historical `Parser/swfparse.cpp`, which reported FWS version 2, exact stage geometry/file length, `tagDefineBitsJPEG2 tagid 1`, ShowFrame and End. Its JPEG2 parser prints the *tag*, **not JPEG pixel validation**. A separate decoder reconstructs a conventional interchange JPEG by moving the emitted DQT/DHT markers into the marker order required by Pillow; Pillow genuinely decodes all pixel buffers and matches the listed hashes and color counts. The native JPEG reference receives only the **three constructed SHA-pinned, ≤16×16 RGB fixtures**, **never arbitrary input**; the original source's JPEG destination has no bounds checks. The owned CLI admits **1..32×1..32 raw RGB only**, ≤8,192 JPEG bytes, ≤16,384 SWF bytes and a required input SHA-256. Four CLI negatives (wrong pin, partial RGB, wrong geometry and >32-pixel width) each exited **2 without an output file**. Two **same-size altered original C++ source** cases exited **1 before native build/report output**. Five owned tests passed, including marker corruption, dimensions, symlink and destination-preservation refusals. [Machine-readable metadata](sdk-jpeg-functional.json).

**Provenance:** the external unchanged [SDK `Parser/jpeg.cpp`](https://raw.githubusercontent.com/SpcFORK/Macromedia-Flash-SDK/bb8536a8f564510f979eb8e09924bc259651a0a9/Parser/jpeg.cpp) is **4,586 bytes**, SHA-256 `3842ad6af89914ef864ea529a73e2cb73124a81c0b6c5cf4102a9135066b1b18`. The separate [original `Parser/swfparse.cpp`](https://raw.githubusercontent.com/SpcFORK/Macromedia-Flash-SDK/bb8536a8f564510f979eb8e09924bc259651a0a9/Parser/swfparse.cpp) is **99,495 bytes**, SHA-256 `d06448e8da95c002ddfaeedae0d5160e70c82e3fa82f311aaea16336d007cf30` and requires the previously documented [portability-only compiler fixes](sdk-swfparse-reference.md). All original sources and private native builds stay **outside** the source patch. The constructed raw RGB cases are owned test data. This reference exactness was observed under **Pillow 13.0.0** and **libjpeg ABI 6.2**; other library versions may differ and are not silently accepted by the saved pins.

Reproduce with separately obtained original sources and installed C++/libjpeg/Pillow:

```sh
cd /home/user/magnolia
python3 sdk-jpeg-functional-oracle.py \
  --jpeg-source /path/outside/to/original/Parser/jpeg.cpp \
  --parser-source /path/outside/to/original/Parser/swfparse.cpp \
  --output /path/to/a/NEW/report.json
python3 -m unittest -v test.swf_jpeg2_encode_test
# For rights-cleared raw RGB24 <=32x32, with a verified input SHA:
python3 swf-jpeg2-encode.py --rgb /path/to/pixels.rgb \
  --width 7 --height 9 --expect-sha256 EXACT_RGB_SHA256 \
  --output /path/to/NEW-bitmap-asset.swf
```

**Acceptance boundary:** a finite, actually encoded and independently decoded **JPEG bitmap asset** source-fragment task works; this does **not** close the SDK's whole `swf-media` or `swf-shape` utility, prove a historical importer on real third-party images, display the bitmap on a stage, play the user's SWF5, run AVM1/AVM2/AS3, or unblock original Java Flash audio and Acotel gameplay. The [ADPCM source-fragment writer](sdk-adpcm-functional.md), [SWF1 semantic parser](sdk-swfparse-functional.md), and [literal URL writer](swf-urlsubst.md) are separate finite slices. The user's every-media completion remains open.
