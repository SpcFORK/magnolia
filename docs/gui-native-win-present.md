# Win32 Frame Present Helpers (gui-native-win-present)

## Overview

`gui-native-win-present` blits the off-screen back-buffer to the visible window DC at the end of each frame. It supports two presentation paths: OpenGL-composited (GDI-based BitBlt) and plain GDI. The OpenGL path automatically falls back to GDI if the blit fails at runtime.

This module is part of the native Windows GUI backend and is not intended for direct use in application code.

## Import

```oak
present := import('gui-native-win-present')
{ presentFrameViaOpenGL: presentFrameViaOpenGL, presentFrameViaGdi: presentFrameViaGdi } := import('gui-native-win-present')
```

## Constants

| Constant  | Value      | Description             |
|-----------|------------|-------------------------|
| `SRCCOPY` | 13369376   | `BitBlt` raster-operation for a direct pixel copy |

## Functions

### `presentFrameViaOpenGL(window)`

Presents when `window.presenterBackend = :opengl` and `window.openglContext > 0`. Uses `BitBlt` to copy from `window.frameHdc` (the GDI back-buffer) to `window.frameTargetHdc`.

On `BitBlt` failure the backend is silently downgraded to `:gdi` and the layer selection is updated to reflect the fallback.

Returns `{ presented: true, detail }` on success or `{ presented: false, detail? }` on failure or when the preconditions are not met.

```oak
result := presentFrameViaOpenGL(window)
if !result.presented -> printf('frame not presented')
```

### `presentFrameViaGdi(window)`

Unconditionally blits `window.frameHdc` to `window.frameTargetHdc` using `BitBlt`/`SRCCOPY`. Returns the raw Win32 call result.

```oak
presentFrameViaGdi(window)
```

**Usage note:** `beginFrame` (from `gui-native-win-frame`) must be called before the first draw call, and one of these present functions must be called after the last draw call to make the frame visible.
