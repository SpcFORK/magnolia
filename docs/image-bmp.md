# image-bmp — BMP Encoder

`import('image-bmp')` encodes 24-bit BMP images from raw pixel data.

## Quick Start

```oak
bmpLib := import('image-bmp')

// pixels: flat list of [b, g, r] lists, 3-byte strings, or 24-bit integers
pixels := [[255, 0, 0], [0, 255, 0], [0, 0, 255], [255, 255, 255]]
data := bmpLib.bmp(2, 2, pixels)

// Write to file
{ writeFile: writeFile } := import('fs')
writeFile('output.bmp', data)
```

## API Reference

### `bmp(width, height, pixels)`

Encodes a 24-bit BMP file and returns the complete file as a byte string.

**Parameters:**
- `width` — image width in pixels
- `height` — image height in pixels
- `pixels` — flat row-major list of pixels; each pixel can be:
  - `[b, g, r]` — byte list in BGR order
  - 3-byte string in BGR order
  - 24-bit integer

## Notes

- Pixel byte order is BGR (blue, green, red) per BMP specification.
- Rows are automatically padded to 4-byte boundaries.
- BMP header is 54 bytes; pixel data follows immediately.
- No compression (BI_RGB).
