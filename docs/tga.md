# TGA Image Format (image-tga)

## Overview

`libtga` encodes and decodes Truevision TGA (TARGA) images. Supports uncompressed and RLE-compressed images in 8-bit grayscale, 24-bit RGB, and 32-bit RGBA.

## Import

```oak
tgaLib := import('image-tga')
{ tga: tga, tgaRLE: tgaRLE, decodeTGA: decodeTGA } := import('image-tga')
// or via facade
img := import('image')
img.tga.tga(w, h, pixels, 24)
```

## Encoders

### `tga(width, height, pixels, bpp)`

Encodes an uncompressed TGA image. `bpp` (bits per pixel) is 8, 24, or 32 (default 24).

**Pixel formats:**
- `bpp=24` — `[b, g, r]` byte lists or 3-byte strings (BGR order)
- `bpp=32` — `[b, g, r, a]` byte lists or 4-byte strings
- `bpp=8` — integers `[0..255]`

```oak
// 2×2 blue image (BGR order)
blue := [255, 0, 0]
data := tga(2, 2, [blue, blue, blue, blue], 24)
```

### `tgaRLE(width, height, pixels, bpp)`

Encodes a run-length encoded TGA image (type 10 for color, type 11 for grayscale). Same pixel format as `tga()`. RLE compresses well when the image contains runs of identical pixels.

```oak
// Solid color compresses very well
red := [0, 0, 255]
pixels := range(1000) |> map(fn(_) red)
data := tgaRLE(100, 10, pixels, 24)
```

## Decoder

### `decodeTGA(data)`

Decodes an uncompressed or RLE-compressed TGA file string. Returns `{ width, height, bpp, pixels }` where pixels is a list of values matching the bit depth:
- 24-bit: `[b, g, r]` lists
- 32-bit: `[b, g, r, a]` lists
- 8-bit: integers

Supports image types 2, 3 (uncompressed) and 10, 11 (RLE). Handles both top-down and bottom-up pixel ordering.
