# Win32 Frame/Backbuffer Helpers (gui-native-win-frame)

## Overview

`gui-native-win-frame` manages the GDI back-buffer lifecycle for `gui-native-win`. It queries the window's client area size, lazily creates a compatible memory DC, and resizes the back-buffer bitmap when needed.

This module is part of the native Windows GUI backend and is not intended for direct use in application code. Use the high-level `GUI` module instead.

## Import

```oak
frame := import('gui-native-win-frame')
{ beginFrame: beginFrame } := import('gui-native-win-frame')
```

## Functions

### `beginFrame(window)`

Prepares the off-screen back-buffer for the next frame. It:

1. Queries the current client area via `GetClientRect` and updates `window.width` / `window.height` if the size has changed.
2. Creates a compatible `memDC` (memory device context) the first time it is called.
3. Creates or recreates a compatible bitmap when the window is first opened or has been resized.
4. Selects the bitmap into the memory DC and records the previous bitmap for cleanup.

Returns `{ type: :ok }` on success or `{ type: :error, error: string, detail: ... }` on Win32 failure.

```oak
result := beginFrame(window)
if result.type != :ok -> printf('beginFrame error: {{0}}', result.error)
// now draw into window.frameHdc ...
```

**Fields managed on `window`**

| Field               | Description                                    |
|---------------------|------------------------------------------------|
| `window.frameHdc`       | Memory DC handle for off-screen drawing    |
| `window.frameBitmap`    | Compatible bitmap handle                   |
| `window.framePrevBitmap`| Previously selected bitmap (restored on cleanup) |
| `window.frameWidth`     | Width of the current back-buffer           |
| `window.frameHeight`    | Height of the current back-buffer          |
| `window.frameTargetHdc` | HDC to present to (set by caller)          |
