# Vulkan Constants & Utilities (gui-native-win-vulkan-constants)

## Overview

`gui-native-win-vulkan-constants` provides all shared Vulkan API constants, DIB constants, and low-level utility functions used by the other `gui-native-win-vulkan-*` sub-modules.

This module is part of the native Windows GUI backend and is not intended for direct use in application code.

## Import

```oak
c := import('gui-native-win-vulkan-constants')
```

## Constants

Includes all `VK_STRUCTURE_TYPE_*`, `VK_IMAGE_LAYOUT_*`, `VK_PIPELINE_STAGE_*`, `VK_ACCESS_*`, `VK_PRESENT_MODE_*`, format, color space, and DIB constants needed by the Vulkan subsystem.

## Utility Functions

- `_vkZeros(n)` — Create a zero-filled list of length `n`
- `_vkWritePtr(address, value)` — Write a pointer-sized value (32 or 64-bit)
- `_default(v, d)` — Return `d` if `v` is `?`, otherwise `v`
- `_vkGetProc(instance, name)` — Resolve a Vulkan extension function via `vkGetInstanceProcAddr`
- `_vkCall(proc, args...)` — Call a resolved Vulkan function pointer
- `_vkCallOk?(r)` — Check if a Vulkan call succeeded (`callOk?` and `VK_SUCCESS`)
