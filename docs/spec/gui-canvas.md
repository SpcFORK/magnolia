# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-canvas.oak`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `linux` · `import(...)`
- `_OK` · `{1 entries}`
- `SRCCOPY` · `13369376`
- `_parentMap` · `{}`
- `_nextId` · `{1 entries}`
### `_getParent(canvas)`

### `create(parentWindow, options)`

### `_ensureCanvasSurface(canvas)`

### `beginCanvas(canvas)`

### `endCanvas(canvas)`

### `_sortedVisibleCanvases(window)`

### `_compositeCanvasWindows(window, canvas)`

### `_compositeCanvasLinux(window, canvas)`

### `_compositeCanvasWeb(window, canvas)`

### `compositeAll(window)`

### `move(canvas, x, y)`

### `resize(canvas, w, h)`

### `setVisible(canvas, vis)`

### `setZIndex(canvas, z)`

### `setOpacity(canvas, alpha)`

### `setTransparentColor(canvas, color)`

### `_releaseCanvasSurface(canvas)`

### `destroy(canvas)`

### `isCanvas?(obj)`

### `canvases(window)`

### `canvasCount(window)`

### `hitTest?(canvas, px, py)`

> returns `:bool`

### `canvasAt(window, px, py)`

### `_findTopmost(sorted, i, px, py)`

> returns `?`

### `toLocal(canvas, px, py)`

> returns `:object`

### `destroyAll(window)`

