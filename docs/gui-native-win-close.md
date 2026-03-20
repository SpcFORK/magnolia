# Win32 Window Close/Teardown Helpers (gui-native-win-close)

## Overview

`gui-native-win-close` provides Win32 resource cleanup routines for `gui-native-win`. It safely releases GDI back-buffer objects, OpenGL rendering contexts, and Vulkan instance resources when a window is destroyed.

This module is part of the native Windows GUI backend and is not intended for direct use in application code. Use the high-level `GUI` module instead.

## Import

```oak
close := import('gui-native-win-close')
{ cleanupFrameBuffers: cleanupFrameBuffers, cleanupOpenGL: cleanupOpenGL, cleanupVulkan: cleanupVulkan } := import('gui-native-win-close')
```

## Functions

### `cleanupFrameBuffers(window)`

Releases the GDI back-buffer associated with `window`. Selects the original bitmap back into the DC, deletes the compatible bitmap, and deletes the memory DC. Resets all frame buffer fields on `window` to `0`. Returns `0`.

```oak
cleanupFrameBuffers(window)
```

Affected `window` fields: `frameHdc`, `frameBitmap`, `framePrevBitmap`, `frameWidth`, `frameHeight`.

### `cleanupOpenGL(window)`

If an OpenGL context is allocated on `window`, calls `wglMakeCurrent(0, 0)` and then `wglDeleteContext` to release it. Resets `window.openglContext` and `window.openglPixelFormat` to `0`. Returns `0`.

```oak
cleanupOpenGL(window)
```

### `cleanupVulkan(window)`

If a Vulkan instance is allocated on `window`, destroys any surface via `vkDestroySurfaceKHR` and then the instance via `vkDestroyInstance`. Resets `window.vulkanInstance`, `window.vulkanSurface`, `window.vulkanPhysicalDevice`, and `window.vulkanQueueFamily`. Returns `0`.

```oak
cleanupVulkan(window)
```
