# image-ico — ICO Encoder

`import('image-ico')` encodes 32-bit ICO (Windows icon) images from raw pixel data.

## Quick Start

```oak
icoLib := import('image-ico')

// pixels: flat list of [b, g, r, a] lists, 4-byte strings, or 32-bit integers
pixels := []
// ... fill 32x32 BGRA pixels ...
data := icoLib.ico(32, 32, pixels)

{ writeFile: writeFile } := import('fs')
writeFile('icon.ico', data)
```

## API Reference

### `ico(width, height, pixels)`

Encodes a single-image 32-bit icon file and returns the complete file as a byte string.

**Parameters:**
- `width` — icon width in pixels
- `height` — icon height in pixels
- `pixels` — flat row-major list of pixels; each pixel can be:
  - `[b, g, r, a]` — byte list in BGRA order
  - 4-byte string in BGRA order
  - 32-bit integer

## Notes

- Pixel byte order is BGRA (blue, green, red, alpha).
- Pixel rows are written bottom-up per ICO/DIB specification.
- Includes a 1-bit AND mask plane as required by the ICO format.
- Single-image icon files only (no multi-resolution icon support).
- Header structure: ICONDIR (6 bytes) + ICONDIRENTRY (16 bytes) + BITMAPINFOHEADER (40 bytes) + pixel data.
