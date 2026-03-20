# Win32 Backend Probe (gui-native-win-probe)

## Overview

`gui-native-win-probe` detects which GPU/graphics backends are available on the current system by loading DLLs at runtime. It supports a Vulkan → OpenGL → DirectDraw → GDI fallback chain and caches the result so subsequent `createWindow` calls pay no probe cost.

This module is part of the native Windows GUI backend and is not intended for direct use in application code. Backend selection is triggered automatically by `gui-native-win` when a window is created.

## Import

```oak
probe := import('gui-native-win-probe')
{ probe2DGraphicsStack: probe2DGraphicsStack, probe2DGdiOnly: probe2DGdiOnly, probeTimings: probeTimings } := import('gui-native-win-probe')
```

## Functions

### `probe2DGraphicsStack(layer2D?)`

Probes available 2D graphics backends and returns the best available one. Backends are tested in this order: Vulkan → OpenGL → DirectDraw → GDI. Stops at the first working backend.

When `layer2D` is `'auto'`, `?`, or omitted, the full fallback chain runs. Pass an explicit backend atom (`:vulkan`, `:opengl`, `:ddraw`, `:gdi`) to skip straight to that backend.

Results are cached after the first call; subsequent calls return the cached value immediately.

Returns a probe result object:

```oak
{
    backend: :vulkan | :opengl | :ddraw | :gdi
    available: true | false
    library: string
    handle: int
    // backend-specific extra fields
}
```

```oak
result := probe2DGraphicsStack()
printf('selected backend: {{0}}', result.backend)
```

### `probe2DGdiOnly()`

Returns a GDI-only probe result without loading any GPU DLL. Use this when the caller has explicitly set `layer2D = 'gdi'` to skip all probe overhead.

```oak
result := probe2DGdiOnly()
// => { backend: :gdi, available: true }
```

### `probeTimings()`

Returns an object with the wall-clock time in seconds spent on each individual DLL probe:

```oak
{
    ddraw:  float   // time spent probing DirectDraw
    d3d9:   float   // time spent probing Direct3D 9
    opengl: float   // time spent probing OpenGL
    vulkan: float   // time spent probing Vulkan
}
```

```oak
t := probeTimings()
printf('vulkan probe: {{0}}s', t.vulkan)
```
