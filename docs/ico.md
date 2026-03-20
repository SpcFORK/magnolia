# ICO Image Encoder (ico)

## Overview

`libico` encodes raw 32-bit RGBA pixel data as a Windows ICO file. The output is a single-image ICO with a 32 bpp, ARGB DIB stored bottom-up, as required by the ICO format. Both standard sizes (up to 255×255) and 256×256 icons are supported.

## Import

```oak
icoLib := import('ico')
// or destructure
{ ico: ico } := import('ico')
```

## Functions

### `ico(width, height, pixels)`

Encodes a 32-bit `width × height` icon and returns the complete binary `.ico` file string.

**Parameters**

- `width` — icon width in pixels (1–256).
- `height` — icon height in pixels (1–256).
- `pixels` — flat row-major list of `width * height` pixels (top-to-bottom order; the encoder reverses rows automatically). Each pixel may be:
  - A 4-byte binary string (raw BGRA bytes).
  - A `[b, g, r, a]` integer list.
  - A 32-bit integer in little-endian BGRA byte order.

```oak
// Solid red 16×16 icon
red := [0, 0, 255, 255]  // [B, G, R, A]
pixels := []
fn fill(i) if i < 16 * 16 {
    true -> {
        pixels << red
        fill(i + 1)
    }
}
fill(0)
data := ico(16, 16, pixels)
writeFile('app.ico', data)
```

```oak
// 32-bit integer pixel (BGRA packed)
transparent := 0x00000000
opaque := 0xFF0000FF  // opaque red in BGRA
data := ico(1, 1, [opaque])
```

## Low-Level Helpers

These are exported for tooling purposes but are not normally needed directly.

### `bytes(parts)`

Converts a list of byte integers to a binary string, masking each value to `[0, 255]`.

### `intToBytes(n, width)`

Converts integer `n` to a little-endian byte list of `width` bytes.

### `hexsplit(n)`

Returns a 4-byte little-endian representation of `n` (shorthand for `intToBytes(n, 4)`).

## Format Notes

- The ICO `ICONDIR` indicates one image entry.
- Width/height values of 256 are stored as `0` in the `ICONDIRENTRY`, per the ICO specification.
- Pixel rows are written bottom-up in the DIB data chunk.
- The AND mask (1-bit transparency) is written as all-zeros (full alpha comes from the 32 bpp BGRA channel).
