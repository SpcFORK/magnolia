# image-tga — TGA Encoder & Decoder

`import('image-tga')` provides encode/decode helpers for Truevision TGA (TARGA) images, supporting uncompressed and RLE-compressed variants with 8-bit grayscale, 24-bit RGB, and 32-bit RGBA.

## Quick Start

```oak
tgaLib := import('image-tga')

// Encode uncompressed 24-bit RGB
pixels := [[255, 0, 0], [0, 255, 0], [0, 0, 255], [255, 255, 0]]
data := tgaLib.tga(2, 2, pixels, 24)

// Encode RLE-compressed 32-bit RGBA
data := tgaLib.tgaRLE(320, 240, rgbaPixels, 32)

// Decode any supported TGA
result := tgaLib.decodeTGA(data)
// result.width, result.height, result.bpp, result.pixels
```

## API Reference

### `tga(width, height, pixels, bpp?)`

Encodes an uncompressed TGA image. Returns the complete file as a byte string.

**Parameters:**
- `width` — image width
- `height` — image height
- `pixels` — flat row-major list of pixel values
- `bpp` — bits per pixel: `8` (grayscale), `24` (RGB), or `32` (RGBA); default `24`

### `tgaRLE(width, height, pixels, bpp?)`

Encodes an RLE-compressed TGA image. Same parameters as `tga()`. Uses image type 10 for color or type 11 for grayscale.

### `decodeTGA(data)`

Decodes a TGA file (uncompressed or RLE). Returns `{width, height, bpp, pixels}`. Handles image types 2, 3, 10, and 11. Auto-flips rows based on the descriptor byte.

## Image Types

| Type | Compression | Color |
|------|------------|-------|
| 2 | None | RGB/RGBA |
| 3 | None | Grayscale |
| 10 | RLE | RGB/RGBA |
| 11 | RLE | Grayscale |

## Notes

- Pixel rows are stored bottom-to-top by default; descriptor bit 5 (value 32) indicates top-down ordering.
- RLE runs are capped at 128 pixels per run header.
- Variable-length image ID fields (0–255 bytes) are skipped during decode.
- Pixel byte order for 24/32-bit is BGR/BGRA per TGA specification.
