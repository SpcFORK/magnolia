# Vulkan Frame Presentation (gui-native-win-vulkan-present)

## Overview

`gui-native-win-vulkan-present` handles per-frame Vulkan presentation. It acquires a swapchain image, copies the GDI backbuffer into the staging buffer via `GetDIBits`, records command buffer operations (image layout transitions + buffer-to-image copy), submits to the queue, and presents.

This module is part of the native Windows GUI backend and is not intended for direct use in application code.

## Import

```oak
{ presentFrameVulkan: presentFrameVulkan } := import('gui-native-win-vulkan-present')
```

## Functions

### `presentFrameVulkan(window)`

Presents one frame via Vulkan. Returns `{ presented: true }` on success, or `{ presented: false, error: string }` on failure. No-ops if the swapchain is not ready.
