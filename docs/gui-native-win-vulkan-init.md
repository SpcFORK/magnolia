# Vulkan Instance & Device Init (gui-native-win-vulkan-init)

## Overview

`gui-native-win-vulkan-init` creates a Vulkan instance and `VK_KHR_win32_surface` surface for a Win32 window. It enumerates physical devices, finds a queue family supporting both graphics and presentation, and creates a logical device.

This module is part of the native Windows GUI backend and is not intended for direct use in application code.

## Import

```oak
{ initVulkan2DLayer: initVulkan2DLayer, createVulkanDevice: createVulkanDevice } := import('gui-native-win-vulkan-init')
```

## Functions

### `initVulkan2DLayer(window)`

Creates a Vulkan instance with surface extensions, creates a Win32 surface, enumerates physical devices, and selects a queue family. Returns `{ type: :ok, instance, surface, selectedPhysicalDevice, selectedQueueFamily, ... }` on success.

### `createVulkanDevice(instance, physicalDevice, queueFamily)`

Creates a logical Vulkan device with `VK_KHR_swapchain` enabled and retrieves the graphics/present queue. Returns `{ type: :ok, device, queue }` on success.
