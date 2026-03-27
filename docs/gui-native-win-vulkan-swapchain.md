# Vulkan Swapchain Lifecycle (gui-native-win-vulkan-swapchain)

## Overview

`gui-native-win-vulkan-swapchain` manages the full swapchain lifecycle: surface capability/format queries, swapchain creation, command pool/buffer/semaphore/fence allocation, staging buffer for pixel uploads, and teardown. It also provides the orchestrator function `initVulkanSwapchain` that wires everything together.

This module is part of the native Windows GUI backend and is not intended for direct use in application code.

## Import

```oak
{ initVulkanSwapchain: initVulkanSwapchain, destroyVulkanSwapchain: destroyVulkanSwapchain } := import('gui-native-win-vulkan-swapchain')
```

## Functions

### `initVulkanSwapchain(window)`

Full swapchain init: creates device, swapchain, command resources, and staging buffer. Returns `{ type: :ok }` on success. On failure, cleans up any partially created resources.

### `destroyVulkanSwapchain(window)`

Tears down all Vulkan resources: staging buffer, sync objects, command pool, swapchain, and logical device.

### `createSwapchain(window)`

Creates the swapchain, queries surface capabilities/formats, and retrieves swapchain images.

### `createVulkanCommandResources(window)`

Allocates command pool, command buffer, semaphores, and fence for frame rendering.

### `createStagingBuffer(window)`

Creates a host-visible staging buffer for copying GDI backbuffer pixels to Vulkan swapchain images.
