# image — Unified Image Format Facade

`import('image')` re-exports all image format libraries under a single namespace.

## Quick Start

```oak
img := import('image')

// BMP
bmpData := img.bmp.bmp(320, 240, pixels)

// ICO
icoData := img.ico.ico(32, 32, pixels)

// PPM / PGM / PBM
ppmData := img.ppm.ppm(320, 240, pixels)
decoded := img.ppm.decodePPM(data)

// TGA (uncompressed and RLE)
tgaData := img.tga.tga(320, 240, pixels, 24)
tgaRle := img.tga.tgaRLE(320, 240, pixels, 32)
decoded := img.tga.decodeTGA(data)

// QOI (lossless compression)
qoiData := img.qoi.qoi(320, 240, pixels, 3)
decoded := img.qoi.decodeQOI(data)
```

## Available Modules

| Property | Module | Formats |
|----------|--------|---------|
| `bmp` | `image-bmp` | 24-bit BMP |
| `ico` | `image-ico` | 32-bit ICO |
| `ppm` | `image-ppm` | PPM (P3/P6), PGM (P2/P5), PBM (P1/P4) |
| `tga` | `image-tga` | TGA (uncompressed & RLE, 8/24/32-bit) |
| `qoi` | `image-qoi` | QOI (RGB/RGBA lossless) |

## See Also

- [image-bmp](image-bmp.md)
- [image-ico](image-ico.md)
- [image-ppm](image-ppm.md)
- [image-tga](image-tga.md)
- [image-qoi](image-qoi.md)
