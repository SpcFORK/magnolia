# Win32 DirectDraw Helpers (gui-native-win-ddraw)

## Overview

`gui-native-win-ddraw` provides low-level DirectDraw 7 COM wrappers used by `gui-native-win` to initialize a primary surface, present a GDI back-buffer via `BitBlt`, and release DirectDraw resources. It targets `DDSCL_NORMAL` cooperative-level (windowed) rendering.

This module is part of the native Windows GUI backend and is not intended for direct use in application code. Use the high-level `GUI` module instead.

## Import

```oak
ddraw := import('gui-native-win-ddraw')
{ initDdrawLayer: initDdrawLayer, presentFrameViaDdraw: presentFrameViaDdraw, releaseDdrawLayer: releaseDdrawLayer } := import('gui-native-win-ddraw')
```

## Constants

| Constant                       | Value      | Description                              |
|--------------------------------|------------|------------------------------------------|
| `COM_RELEASE`                  | 2          | vtable index for `IUnknown::Release`     |
| `IDirectDraw7_CreateSurface`   | 6          | vtable index                             |
| `IDirectDraw7_SetCooperativeLevel` | 20      | vtable index                             |
| `IDirectDrawSurface7_GetDC`    | 17         | vtable index                             |
| `IDirectDrawSurface7_ReleaseDC`| 26         | vtable index                             |
| `DDSCL_NORMAL`                 | 8          | Windowed cooperative level               |
| `DDSD_CAPS`                    | 1          | Surface descriptor caps flag             |
| `DDSCAPS_PRIMARYSURFACE`       | 512        | Primary surface caps flag                |
| `SRCCOPY`                      | 13369376   | `BitBlt` raster-operation code           |

## Functions

### `initDdrawLayer(window, ddrawHandle)`

Initializes a DirectDraw 7 primary surface for `window` using the already-loaded `ddrawHandle` (the DLL handle returned by `windows.loadDll`). Sets the cooperative level to `DDSCL_NORMAL`, creates a primary surface, and stores the result on `window`.

Returns `{ type: :ok, ... }` on success, or `{ type: :error, error: string, ... }` on failure.

```oak
result := initDdrawLayer(window, ddrawHandle)
if result.type = :ok -> printf('DirectDraw ready')
```

### `presentFrameViaDdraw(window)`

Presents the current GDI back-buffer (`window.frameHdc`) to the primary DirectDraw surface using `GetDC`/`BitBlt`/`ReleaseDC`. Falls back to GDI mode on failure.

```oak
presentFrameViaDdraw(window)
```

### `releaseDdrawLayer(window)`

Releases the DirectDraw primary surface and DirectDraw object attached to `window`, resetting all related fields.

```oak
releaseDdrawLayer(window)
```
