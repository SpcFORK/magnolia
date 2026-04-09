# PPM / PGM / PBM Image Formats (image-ppm)

## Overview

`libppm` provides encode and decode helpers for the Netpbm family of image formats:

- **PPM** — Portable Pixmap (RGB color), P3 (plain) and P6 (binary)
- **PGM** — Portable Graymap (grayscale), P2 (plain) and P5 (binary)
- **PBM** — Portable Bitmap (black & white), P1 (plain) and P4 (binary)

Binary variants are compact; plain variants are human-readable ASCII.

## Import

```oak
ppmLib := import('image-ppm')
// or destructure
{ ppm: ppm, pgm: pgm, pbm: pbm, decodePPM: decodePPM } := import('image-ppm')
// or via facade
img := import('image')
img.ppm.ppm(w, h, pixels)
```

## Encoders

### `ppm(width, height, pixels)` — P6 binary RGB

Encodes a P6 (binary) PPM image. Each pixel is a `[r, g, b]` integer list or a 3-byte string.

```oak
data := ppm(2, 2, [[255, 0, 0], [0, 255, 0], [0, 0, 255], [255, 255, 255]])
```

### `ppmPlain(width, height, pixels)` — P3 plain RGB

Same as `ppm()` but outputs human-readable ASCII format.

### `pgm(width, height, pixels)` — P5 binary grayscale

Each pixel is an integer in `[0, 255]`.

```oak
data := pgm(4, 4, range(16) |> map(fn(i) i * 17))
```

### `pgmPlain(width, height, pixels)` — P2 plain grayscale

ASCII variant of `pgm()`.

### `pbm(width, height, pixels)` — P4 binary bitmap

Each pixel is `0` (white) or `1` (black). Bits are packed 8 per byte, rows padded to byte boundaries.

```oak
data := pbm(8, 1, [1, 0, 1, 0, 1, 0, 1, 0])
```

### `pbmPlain(width, height, pixels)` — P1 plain bitmap

ASCII variant of `pbm()`.

## Decoders

### `decodePPM(data)`

Decodes a P3 or P6 PPM file string. Returns `{ width, height, pixels }` where pixels is a list of `[r, g, b]` lists. Returns `?` on unrecognized format.

### `decodePGM(data)`

Decodes a P2 or P5 PGM file string. Returns `{ width, height, pixels }` where pixels is a flat list of integers `[0..255]`.
