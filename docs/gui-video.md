# GUI Video Bridge (gui-video)

## Overview

`gui-video` bridges the `video` frame format with the GUI shader pixel buffer and rendering system. It enables:

- Converting between video frames and shader pixel buffers
- Rendering video frames directly into GUI windows
- Applying video effects (grayscale, invert, threshold, blend, diff) to shader buffers
- Exporting shader buffers as video frames or BMP-ready pixel data

## Import

```oak
// Direct import
guiVideo := import('gui-video')

// Or through the GUI facade (recommended)
gui := import('GUI')
```

## Format Conversion

### `frameToBuffer(frame)` / `gui.videoFrameToBuffer(frame)`

Converts a video.oak frame (RGB channel bytes) into a shader pixel buffer (`{width, height, data:{packed-BGR-ints}}`).

### `bufferToFrame(buf)` / `gui.videoBufferToFrame(buf)`

Converts a shader pixel buffer back into a 3-channel video.oak frame.

## Rendering

### `renderFrame(drawCtx, window, frame, ox?, oy?)` / `gui.videoRenderFrame(window, frame, ox?, oy?)`

Draws a video frame into the GUI window at offset `(ox, oy)` using 1×1 `fillRect` calls.

### `renderFrameScaled(drawCtx, window, frame, ox?, oy?, scale?)` / `gui.videoRenderFrameScaled(window, frame, ox?, oy?, scale?)`

Renders a video frame scaled by an integer factor.

## Video Effects on Shader Buffers

### `bufferGrayscale(buf)` / `gui.videoBufferGrayscale(buf)`

Applies luma-weighted grayscale conversion.

### `bufferInvert(buf)` / `gui.videoBufferInvert(buf)`

Inverts all color channels.

### `bufferThreshold(buf, t?)` / `gui.videoBufferThreshold(buf, t?)`

Binary black/white conversion based on luma threshold (default 127).

### `bufferBlend(bufA, bufB, alpha?)` / `gui.videoBufferBlend(bufA, bufB, alpha?)`

Alpha-blends two buffers (default alpha 0.5).

### `bufferDiff(bufA, bufB)` / `gui.videoBufferDiff(bufA, bufB)`

Per-channel absolute difference between two buffers.

## Capture and Export

### `captureBuffer(buf)` / `gui.videoCaptureBuffer(buf)`

Captures the current shader pixel buffer state as a video.oak frame for further processing.

### `frameToBmpPixels(frame)` / `gui.videoFrameToBmpPixels(frame)`

Converts a video.oak RGB frame into the BMP pixel list expected by `lib/bmp` (`[B, G, R]` byte lists in bottom-up row order). The result can be passed directly to `bmp.bmp(w, h, pixels)`.

## Example: Shader Post-Processing

```oak
gui := import('GUI')

// Create a shader buffer and render something into it
buf := gui.shaderCreateBuffer(320, 240)
shader := gui.Shader(fn(x, y, w, h, t, u) {
    gui.rgb(int(x * 255 / w), int(y * 255 / h), 128)
})
gui.shaderRenderShaderToBuffer(buf, shader)

// Apply grayscale post-processing
grayBuf := gui.videoBufferGrayscale(buf)

// Render the processed buffer
gui.shaderRenderBuffer(window, grayBuf, 0, 0)
```

## Example: Render a Video Frame

```oak
gui := import('GUI')
video := import('video')

// Create a test frame (red gradient)
f := video.mapPixels(video.blank(100, 100), fn(pixel, x, y) {
    [int(x * 255 / 100), 0, 0]
})

// Render at 2× scale
gui.videoRenderFrameScaled(window, f, 10, 10, 2)
```

## Example: Export to BMP

```oak
gui := import('GUI')
bmpLib := import('bmp')
video := import('video')

f := video.blank(64, 64, 3, 128)
pixels := gui.videoFrameToBmpPixels(f)
data := bmpLib.bmp(64, 64, pixels)
// write data to file...
```
