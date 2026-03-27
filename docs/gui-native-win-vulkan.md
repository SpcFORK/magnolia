# Win32 Vulkan Bootstrap (gui-native-win-vulkan)

## Overview

`gui-native-win-vulkan` is a facade module that re-exports the public Vulkan API from four sub-modules:

- **gui-native-win-vulkan-constants** — Shared Vulkan constants, DIB constants, and low-level helpers
- **gui-native-win-vulkan-init** — Instance, surface, device, and queue discovery
- **gui-native-win-vulkan-swapchain** — Swapchain lifecycle, command resources, staging buffer, and cleanup
- **gui-native-win-vulkan-present** — Per-frame Vulkan presentation

This module is part of the native Windows GUI backend and is not intended for direct use in application code.

## Import

```oak
vk := import('gui-native-win-vulkan')
{ initVulkan2DLayer: initVulkan2DLayer } := import('gui-native-win-vulkan')
```

## Re-exported Functions

| Function | Source Module |
|---|---|
| `initVulkan2DLayer(window)` | gui-native-win-vulkan-init |
| `createVulkanDevice(instance, physicalDevice, queueFamily)` | gui-native-win-vulkan-init |
| `initVulkanSwapchain(window)` | gui-native-win-vulkan-swapchain |
| `destroyVulkanSwapchain(window)` | gui-native-win-vulkan-swapchain |
| `createSwapchain(window)` | gui-native-win-vulkan-swapchain |
| `createVulkanCommandResources(window)` | gui-native-win-vulkan-swapchain |
| `createStagingBuffer(window)` | gui-native-win-vulkan-swapchain |
| `presentFrameVulkan(window)` | gui-native-win-vulkan-present |
