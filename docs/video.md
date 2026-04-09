# Video Library (video)

## Overview

`libvideo` provides frame-based pixel processing helpers for Magnolia.

It is designed for lightweight video-style pipelines where each frame is represented as:

```oak
{
    width: <int>
    height: <int>
    channels: <int>
    pixels: [byte...]
}
```

Pixels are row-major and interleaved by channel.

## Import

```oak
video := import('video')
```

## Frame Construction

### `frame(width, height, pixels?, channels?)`

Creates a frame object. Defaults:
- `channels = 3`
- `pixels =` zero-filled byte buffer

### `blank(width, height, channels?, value?)`

Creates a solid-color frame filled with a single byte value.

### `cloneFrame(frame)`

Clones frame metadata and pixel data.

## Pixel Access and Mapping

### `pixelIndex(frame, x, y, ch?)`

Returns the flat index into `frame.pixels`.

### `getPixel(frame, x, y)`

Reads one pixel and returns channel values as a list.

### `setPixel(frame, x, y, pixel)`

Returns a new frame with one pixel overwritten.

### `mapPixels(frame, mapper)`

Transforms all pixels with `mapper(pixel, x, y)`.

## Color Operations

### `toGrayscale(frame)`

Converts RGB/RGBA to grayscale using luma weighting while preserving frame shape.

### `invert(frame, maxValue?)`

Inverts channels against `maxValue` (default `255`).

### `threshold(frame, t?)`

Converts to black/white via luma threshold (`t`, default `127`).

### `rgbToYuv([r, g, b])`

Converts one RGB pixel to `[y, u, v]`.

### `yuvToRgb([y, u, v])`

Converts one YUV pixel to `[r, g, b]`.

## Geometric Operations

### `crop(frame, x, y, width, height)`

Extracts a clamped rectangle from a frame.

### `resizeNearest(frame, newWidth, newHeight)`

Resizes with nearest-neighbor sampling.

## Blending and Sequences

### `blend(a, b, alpha?)`

Blends two frames with `alpha` in `[0, 1]`. Uses overlapping top-left region if dimensions differ.

### `frameDiff(a, b)`

Computes per-channel absolute difference.

### `mapFrames(frames, f)`

Maps frame sequences with `f(frame, index)`.

### `sampleFrame(frames, timeSeconds, fps?)`

Returns nearest frame for a timestamp (`fps` default `30`).

## Example

```oak
video := import('video')

base := video.blank(320, 180, 3, 32)
overlay := video.blank(320, 180, 3, 200)

mixed := video.blend(base, overlay, 0.25)
grayscale := video.toGrayscale(mixed)
small := video.resizeNearest(grayscale, 160, 90)

edge := video.frameDiff(grayscale, mixed)
frameAtTime := video.sampleFrame([base, mixed, grayscale], 0.08, 30)
```

## Parallel Pixel Processing

### `pmapPixels(frame, mapper, numWorkers?)`

Applies `mapper(pixel, x, y)` in parallel using bounded concurrency. Each row is processed as an independent unit of work. `numWorkers` defaults to 4.

```oak
bright := video.pmapPixels(f, fn(px, x, y) [
    px.0 + 30, px.1 + 30, px.2 + 30
], 4)
```

### `presizeNearest(frame, newWidth, newHeight, numWorkers?)`

Rescales a frame using nearest-neighbor sampling with parallel row processing.

```oak
scaled := video.presizeNearest(f, 1920, 1080, 4)
```

### `pmapFrames(frames, f)`

Applies `f(frame, index)` over a sequence of frames in parallel.

```oak
processed := video.pmapFrames(frames, fn(f, i) video.toGrayscale(f))
```

## Notes

- Channel and byte values are clamped where appropriate.
- Most operations return new frames rather than mutating the input.
- The library targets practical frame processing, not codec/container decoding.
