# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\video.oak`

### `clampByte(n)`

> returns `:int`

### `clampUnit(n)`

> returns `:int`

### `_min(a, b)`

### `_max(a, b)`

### `_pixelIndexRaw(width, channels, x, y, ch)`

### `pixelIndex(frame, x, y, ch)`

### `frame(width, height, pixels, channels)`

> returns `:object`

### `blank(width, height, channels, value)`

### `cloneFrame(f)`

> returns `:object`

### `getPixel(f, x, y)`

### `setPixel(f, x, y, pixel)`

> returns `:object`

### `mapPixels(f, mapper)`

> returns `:object`

### `_luma(pixel)`

### `toGrayscale(f)`

### `invert(f, maxValue)`

### `threshold(f, t)`

### `rgbToYuv(pixel)`

> returns `:list`

### `yuvToRgb(pixel)`

> returns `:list`

### `crop(f, x, y, width, height)`

### `resizeNearest(f, newWidth, newHeight)`

### `blend(a, b, alpha)`

### `frameDiff(a, b)`

### `mapFrames(frames, f)`

### `sampleFrame(frames, timeSeconds, fps)`

> returns `?`

### `pmapPixels(f, mapper, numWorkers)`

> returns `:object`

### `presizeNearest(f, newWidth, newHeight, numWorkers)`

### `pmapFrames(frames, f)`

