# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-filedrop.oak`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `WM_DROPFILES` · `563`
### `enableFileDrop(window)`

### `disableFileDrop(window)`

### `_readUtf16Str(addr)`

### `onFileDrop(window, handler)`

- `DROPEFFECT_NONE` · `0`
- `DROPEFFECT_COPY` · `1`
- `DROPEFFECT_MOVE` · `2`
- `DROPEFFECT_LINK` · `4`
### `enableOleDrop(window)`

> returns `:bool`

### `disableOleDrop(window)`

> returns `:bool`

### `onOleDrop(window, handler)`

### `dragDropState()`

> returns `:object`

### `onDragOver(window, state, handler)`

