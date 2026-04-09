# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-draw.oak`

- `windows` · `import(...)`
- `linux` · `import(...)`
- `guiFonts` · `import(...)`
- `guiThread` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_asBool(value)`

> returns `:bool`

### `_fontKey(fontSpec)`

### `_windowsDeleteCachedFont(window)`

> returns `?`

### `_ensureWindowsFont(window)`

> returns `:int`

### `_ensureLinuxFont(window, gcHandle)`

> returns `:int`

### `_webFontString(fontSpec)`

### `_windowsDeleteCachedPens(window)`

> returns `:int`

### `_windowsDeleteCachedBrushes(window)`

> returns `:int`

### `releaseResources(window)`

> returns `:int`

### `invalidateDrawCaches(window)`

> returns `:int`

### `drawText(window, x, y, text, color, defaultColor)`

### `textWidth(window, text)`

### `_estimateTextWidth(window, text)`

### `fillRect(window, x, y, width, height, color, defaultColor, borderColor)`

- `_maxCacheSize` · `256`
### `_evictPenCache(window)`

> returns `?`

### `_evictBrushCache(window)`

> returns `?`

### `_getCachedPen(window, hdcValue, useColor)`

### `_getCachedBrush(window, useColor)`

- `_nullPenHandle` · `0`
### `_getNullPen()`

### `pushMask(window, x, y, w, h)`

### `popMask(window)`

### `drawLine(window, x1, y1, x2, y2, color, defaultColor)`

### `_setupFillGDI(window, hdcValue, fillColor, borderColor)`

> returns `:bool`

### `fillEllipse(window, cx, cy, rx, ry, fillColor, borderColor)`

### `fillRoundedRect(window, x, y, width, height, radius, fillColor, borderColor)`

### `_writeI32(address, value)`

### `_ensurePointBuf(window, n)`

### `fillPolygon(window, pts, fillColor, borderColor)`

### `drawPolylineNative(window, pts, color)`

### `drawLinesBatch(window, segs, color, defaultColor)`

### `DrawQueue()`

> returns `:object`

### `parallelFillRects(window, coords, computeFn, numWorkers)`

