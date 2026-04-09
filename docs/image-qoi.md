# image-qoi — QOI Encoder & Decoder

`import('image-qoi')` implements the Quite OK Image (QOI) format for lossless image compression with simple encoding logic. Supports RGB (3-channel) and RGBA (4-channel) images.

## Quick Start

```oak
qoiLib := import('image-qoi')

// Encode an RGBA image
pixels := [[255, 0, 0, 255], [0, 255, 0, 255], [0, 0, 255, 255]]
data := qoiLib.qoi(3, 1, pixels, 4)

// Encode an RGB image (default)
data := qoiLib.qoi(320, 240, rgbPixels)

// Decode a QOI file
result := qoiLib.decodeQOI(data)
// result.width, result.height, result.channels, result.pixels
```

## API Reference

### `qoi(width, height, pixels, channels?)`

Encodes a QOI image and returns the complete file as a byte string.

**Parameters:**
- `width` — image width
- `height` — image height
- `pixels` — flat row-major list of `[r, g, b]` or `[r, g, b, a]` lists
- `channels` — 3 (RGB) or 4 (RGBA), default 3

### `decodeQOI(data)`

Decodes a QOI file string. Returns `{width, height, channels, pixels}` where pixels are `[r, g, b]` or `[r, g, b, a]` lists.

## Compression Modes

| Op Code | Tag | Description |
|---------|-----|-------------|
| `OP_INDEX` | 0x00 | Pixel matches a previously seen color (64-entry hash table) |
| `OP_DIFF` | 0x40 | Small RGB delta (-2..1 per channel) |
| `OP_LUMA` | 0x80 | Luma delta with chroma offsets |
| `OP_RUN` | 0xC0 | Run-length encoding (1–62 identical pixels) |
| `OP_RGB` | 0xFE | Full RGB pixel (no alpha change) |
| `OP_RGBA` | 0xFF | Full RGBA pixel |

## Notes

- QOI achieves compression ratios comparable to PNG with much simpler encode/decode logic.
- The 64-entry running hash table enables efficient index lookups for repeated colors.
- Lossless — pixel-perfect round-trip guaranteed.
