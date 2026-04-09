# image-ppm — Netpbm Format Encoder & Decoder

`import('image-ppm')` provides encode/decode helpers for the Netpbm family of image formats: PPM (RGB), PGM (grayscale), and PBM (bitmap), in both binary and ASCII variants.

## Quick Start

```oak
ppmLib := import('image-ppm')

// Encode a binary PPM (P6) image
pixels := [[255, 0, 0], [0, 255, 0], [0, 0, 255], [255, 255, 255]]
data := ppmLib.ppm(2, 2, pixels)

// Decode a PPM file
result := ppmLib.decodePPM(data)
// result.width, result.height, result.pixels

// Encode grayscale PGM
grayPixels := [0, 64, 128, 255]
data := ppmLib.pgm(2, 2, grayPixels)

// Encode bitmap PBM
bits := [1, 0, 0, 1]
data := ppmLib.pbm(2, 2, bits)
```

## PPM (RGB Color)

### `ppm(width, height, pixels)`

Encodes a P6 (binary) PPM image. Pixels are a flat row-major list of `[r, g, b]` lists or 3-byte strings.

### `ppmPlain(width, height, pixels)`

Encodes a P3 (ASCII) PPM image. Same pixel format.

### `decodePPM(data)`

Decodes a P3 or P6 PPM file. Returns `{width, height, pixels}` where pixels are `[r, g, b]` lists.

## PGM (Grayscale)

### `pgm(width, height, pixels)`

Encodes a P5 (binary) PGM image. Pixels are a flat row-major list of integers `[0..255]`.

### `pgmPlain(width, height, pixels)`

Encodes a P2 (ASCII) PGM image.

### `decodePGM(data)`

Decodes a P2 or P5 PGM file. Returns `{width, height, pixels}` where pixels are integers.

## PBM (Bitmap)

### `pbm(width, height, pixels)`

Encodes a P4 (binary) PBM image. Pixels are a flat row-major list of `0` (white) or `1` (black).

### `pbmPlain(width, height, pixels)`

Encodes a P1 (ASCII) PBM image.

## Format Summary

| Magic | Format | Type | Channels |
|-------|--------|------|----------|
| P1 | PBM ASCII | Bitmap | 1-bit |
| P2 | PGM ASCII | Grayscale | 8-bit |
| P3 | PPM ASCII | RGB | 3×8-bit |
| P4 | PBM Binary | Bitmap | 1-bit |
| P5 | PGM Binary | Grayscale | 8-bit |
| P6 | PPM Binary | RGB | 3×8-bit |

## Notes

- Maximum sample value is 255 for PPM/PGM.
- Decoders automatically skip comments (`#`) and whitespace in headers.
- Binary formats are more compact; ASCII formats are human-readable.
