# rMQR — Rectangular Micro QR Code Encoder

`import('rmqr')` generates ISO/IEC 23941:2022 **Rectangular Micro QR (rMQR)** symbols. rMQR is designed for narrow spaces such as labels, receipts, and product packaging. Magnolia produces the unscaled module matrix; render it directly, inspect it in a terminal, or export a PBM bitmap for the existing `image-ppm` workflow.

The encoder implements all 32 standardized sizes (R7x43 through R17x139), both error-correction levels, BCH format information, the rMQR mask, and GF(256) Reed–Solomon correction. It supports Numeric, Alphanumeric, and UTF-8 Byte segments. Kanji requires Shift-JIS conversion and is intentionally not encoded by this module.

## Quick start

```oak
rmqr := import('rmqr')
{ writeFile: writeFile } := import('fs')

// Select a size and segment modes automatically.
code := rmqr.fit('https://magnolia-lang.org/rMQR', rmqr.H, rmqr.Balanced)

if code != ? {
    // A PBM P4 bitmap, six pixels per module with a two-module quiet zone.
    writeFile('code.pbm', rmqr.pbm(code, 6, true))
    print(rmqr.ascii(code))
}
```

`fit` and `make` return `?` when the input is invalid or cannot fit the requested symbol. A successful result is:

```oak
{
    version: 'R13x99'
    ecc: :H
    width: 99
    height: 13
    matrix: [[...], ...]  // row-major: 1 is black, 0 is white
    segments: [...]
}
```

## API

### Constants

| Constant | Meaning |
|---|---|
| `rmqr.M` | Medium error correction (approximately 15% restoration) |
| `rmqr.H` | High error correction (approximately 30% restoration) |
| `rmqr.Numeric` | Digits `0`–`9` only |
| `rmqr.Alphanumeric` | `0`–`9`, `A`–`Z`, space, and `$%*+-./:` |
| `rmqr.Byte` | Raw UTF-8 bytes |
| `rmqr.Kanji` | Reserved mode; currently unsupported and therefore returns `?` |
| `rmqr.MinimizeWidth` | Choose the narrowest fitting symbol |
| `rmqr.MinimizeHeight` | Choose the shortest fitting symbol |
| `rmqr.Balanced` | Default `fit` strategy, balancing height and width |

### `segment(data, mode?)`

Creates a segment for manual construction. `mode` defaults to `rmqr.Byte`.

```oak
parts := [
    rmqr.segment('20261018', rmqr.Numeric)
    rmqr.segment(' ORDER-A7', rmqr.Alphanumeric)
]
```

### `make(versionName, ecc?, segments)`

Encodes segments at a named standard version. Use this when physical label space fixes the symbol dimensions.

```oak
code := rmqr.make('R11x59', rmqr.H, [
    rmqr.segment('123456', rmqr.Numeric)
    rmqr.segment('ABCD', rmqr.Alphanumeric)
])
```

The mode constraints and character-count limits are checked. Unknown version names, invalid ECC values, unsupported Kanji, empty segment arrays, and oversized content return `?`.

### `fit(data, ecc?, strategy?)`

Optimizes an input string into Numeric, Alphanumeric, and Byte segments, then selects a standardized size.

```oak
narrow := rmqr.fit('INV-2026-0042', rmqr.M, rmqr.MinimizeWidth)
short  := rmqr.fit('INV-2026-0042', rmqr.M, rmqr.MinimizeHeight)
defaultChoice := rmqr.fit('INV-2026-0042') // M + Balanced
```

Segmentation works over the UTF-8 bytes of an Oak string. Non-ASCII characters stay in Byte mode; character-count limits therefore count encoded bytes in that mode.

### `matrix(code, quietZone?)`

Returns the row-major `0`/`1` matrix. The optional quiet zone is disabled by default; when enabled, it adds the rMQR-recommended two white modules on each edge.

```oak
modules := rmqr.matrix(code)        // exact symbol dimensions
withBorder := rmqr.matrix(code, true)
```

### `ascii(code, quietZone?)`

Returns a terminal preview. It uses two-character-wide black cells to keep square modules visually proportioned. Its quiet zone defaults to `true`.

### `pbm(code, scale?, quietZone?)`

Returns a binary PBM P4 image. `scale` defaults to `1`; `quietZone` defaults to `true`. This is a convenient lossless interchange format for image tools and the Magnolia `image-ppm` module.

```oak
fileData := rmqr.pbm(code, 4, true)
```

### `versions()`

Returns metadata records for the available sizes:

```oak
{ name: 'R7x43', width: 43, height: 7, dataBitsM: 48, dataBitsH: 24 }
```

## Standard symbol sizes

All dimensions below are width × height in modules. At the largest size, an M-level numeric segment can hold the standard maximum of 361 digits.

| Height | Available widths |
|---:|---|
| 7 | 43, 59, 77, 99, 139 |
| 9 | 43, 59, 77, 99, 139 |
| 11 | 27, 43, 59, 77, 99, 139 |
| 13 | 27, 43, 59, 77, 99, 139 |
| 15 | 43, 59, 77, 99, 139 |
| 17 | 43, 59, 77, 99, 139 |

## Implementation references

Magnolia's encoder is a self-contained Oak implementation; it has no Python runtime dependency. Its development was cross-checked against the MIT-licensed [OUDON/rmqrcode-python](https://github.com/OUDON/rmqrcode-python) reference repository, including its version/capacity tables, segment encoding, format placement, masking, and Reed–Solomon workflow. The ISO specification remains the governing behavior where a reference implementation differs.

## Notes

- rMQR readers are less ubiquitous than ordinary QR Code readers. Test the final printed image with the scanners used in your deployment.
- Use the default two-module quiet zone unless another system explicitly supplies suitable clearance around the symbol.
- See the DENSO WAVE [rMQR overview](https://www.qrcode.com/en/codes/rmqr.html) and the ISO [ISO/IEC 23941:2022 record](https://www.iso.org/standard/77404.html) for the standard and ecosystem context.
