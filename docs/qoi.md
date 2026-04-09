# QOI Image Format (image-qoi)

## Overview

`libqoi` implements the [Quite OK Image Format](https://qoiformat.org/) — a lossless image compression format that is simple, fast, and achieves compression ratios competitive with PNG. It supports RGB (3-channel) and RGBA (4-channel) images.

## Import

```oak
qoiLib := import('image-qoi')
{ qoi: qoi, decodeQOI: decodeQOI } := import('image-qoi')
// or via facade
img := import('image')
img.qoi.qoi(w, h, pixels, 3)
```

## Encoder

### `qoi(width, height, pixels, channels)`

Encodes a QOI image and returns the full file string. `channels` is 3 (RGB, default) or 4 (RGBA).

Each pixel is a `[r, g, b]` or `[r, g, b, a]` integer list.

```oak
pixels := [[255, 0, 0], [0, 255, 0], [0, 0, 255], [255, 255, 255]]
data := qoi(2, 2, pixels, 3)
```

QOI uses several encoding strategies automatically:
- **Index** — reference a recently seen pixel by hash
- **Diff** — encode small channel differences (±2)
- **Luma** — encode medium differences via green-channel delta
- **Run** — encode runs of identical pixels (up to 62)
- **RGB/RGBA** — full literal pixel when no compression applies

```oak
// Solid-color images compress very efficiently
red := [255, 0, 0]
pixels := range(10000) |> map(fn(_) red)
data := qoi(100, 100, pixels, 3)
// Much smaller than 100*100*3 raw bytes
```

## Decoder

### `decodeQOI(data)`

Decodes a QOI file string into `{ width, height, channels, pixels }`. Returns `?` if the magic bytes are invalid.

```oak
result := decodeQOI(fileData)
// result.width, result.height, result.channels
// result.pixels — list of [r,g,b] or [r,g,b,a] lists
```
