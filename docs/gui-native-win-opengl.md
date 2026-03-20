# Win32 OpenGL Bootstrap (gui-native-win-opengl)

## Overview

`gui-native-win-opengl` initializes an OpenGL 2D rendering layer for a Win32 window via `wgl`. It selects a compatible pixel format using a standard `PIXELFORMATDESCRIPTOR`, creates a context with `wglCreateContext`, and verifies it with `wglMakeCurrent`.

This module is part of the native Windows GUI backend and is not intended for direct use in application code. OpenGL layer selection is controlled by the `GUI.createWindow` options object.

## Import

```oak
opengl := import('gui-native-win-opengl')
{ initOpenGL2DLayer: initOpenGL2DLayer } := import('gui-native-win-opengl')
```

## Constants

| Constant            | Value | Description                                  |
|---------------------|-------|----------------------------------------------|
| `PFD_DRAW_TO_WINDOW`| 4     | Pixel format flag: render to window DC       |
| `PFD_SUPPORT_OPENGL`| 32    | Pixel format flag: support OpenGL            |
| `PFD_DOUBLEBUFFER`  | 1     | Pixel format flag: double-buffered           |

## Functions

### `initOpenGL2DLayer(window)`

Attempts to initialize an OpenGL context on `window`. The sequence is:

1. Acquire the window DC with `GetDC`.
2. Build a 40-byte `PIXELFORMATDESCRIPTOR` requesting 32-bit RGBA with 24-bit depth and 8-bit stencil, double-buffered, drawn to a window.
3. Select a pixel format with `ChoosePixelFormat` / `SetPixelFormat`.
4. Create the context with `wglCreateContext` and verify it with `wglMakeCurrent`.
5. Release the resources and return the context handle.

Returns `{ type: :ok, context, pixelFormat, makeCurrentResult }` on success, or `{ type: :error, error: string, ... }` on failure.

```oak
result := initOpenGL2DLayer(window)
if result.type = :ok {
    true -> {
        window.openglContext <- result.context
        window.openglPixelFormat <- result.pixelFormat
    }
    _ -> printf('OpenGL init failed: {{0}}', result.error)
}
```

**Note:** The context is released (`wglMakeCurrent(0,0)`) immediately after creation. Callers are responsible for making it current before issuing draw calls.
