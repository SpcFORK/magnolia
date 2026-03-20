# BMP Image Encoder (bmp)

## Overview

`libbmp` encodes raw pixel data as a 24-bit Windows BMP image file. Pixels are written row-major (left-to-right, top-to-bottom) and each row is padded to a 4-byte boundary as required by the BMP specification.

## Import

```oak
bmpLib := import('bmp')
// or destructure
{ bmp: bmp } := import('bmp')
```

## Functions

### `bmp(width, height, pixels)`

Encodes a 24-bit BMP file and returns the complete binary file string.

**Parameters**

- `width` — image width in pixels.
- `height` — image height in pixels.
- `pixels` — flat row-major list of `width * height` pixels. Each pixel may be:
  - A 3-byte binary string (raw BGR bytes).
  - A `[b, g, r]` integer list (blue, green, red order matching the BMP format).
  - A 24-bit integer in little-endian BGR byte order.

```oak
// 2×2 red image
red := [0, 0, 255]  // [B, G, R]
pixels := [red, red, red, red]
data := bmp(2, 2, pixels)
writeFile('red.bmp', data)
```

```oak
// integer pixel (BGR packed as 24-bit int)
white := 0xFFFFFF
data := bmp(1, 1, [white])
```

## Low-Level Helpers

These are exported for tooling and testing purposes but are not normally needed directly.

### `bytes(parts)`

Converts a list of byte integers to a binary string, masking each value to the range `[0, 255]`.

```oak
bytes([72, 101, 108]) // => 'Hel'
```

### `intToBytes(n, width)`

Converts integer `n` to a little-endian byte list of `width` bytes.

```oak
intToBytes(256, 2) // => [0, 1]
intToBytes(1, 4)   // => [1, 0, 0, 0]
```

### `hexsplit(n)`

Shorthand for `intToBytes(n, 4)` — returns a 4-byte little-endian representation.

```oak
hexsplit(54) // => [54, 0, 0, 0]
```
