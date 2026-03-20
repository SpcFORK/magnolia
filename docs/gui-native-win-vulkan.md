# Win32 Vulkan Bootstrap (gui-native-win-vulkan)

## Overview

`gui-native-win-vulkan` creates a Vulkan instance and a `VK_KHR_win32_surface` surface for a Win32 window. It enumerates physical devices, finds a queue family that supports both graphics and presentation, and returns the relevant handles so the caller can proceed to device creation.

This module is part of the native Windows GUI backend and is not intended for direct use in application code.

## Import

```oak
vk := import('gui-native-win-vulkan')
{ initVulkan2DLayer: initVulkan2DLayer } := import('gui-native-win-vulkan')
```

## Constants

| Constant                                | Value        | Description                           |
|-----------------------------------------|--------------|---------------------------------------|
| `VK_SUCCESS`                            | 0            | Vulkan success status code            |
| `VK_API_VERSION_1_0`                    | 4194304      | `VK_MAKE_API_VERSION(0,1,0,0)`       |
| `VK_QUEUE_GRAPHICS_BIT`                 | 1            | Queue family graphics capability flag  |
| `VK_STRUCTURE_TYPE_APPLICATION_INFO`    | 0            | `sType` value for `VkApplicationInfo` |
| `VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO`| 1            | `sType` for `VkInstanceCreateInfo`    |
| `VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR` | 1000009000 | `sType` for Win32 surface  |
| `VK_KHR_SURFACE_EXTENSION_NAME`         | string       | `'VK_KHR_surface'`                    |
| `VK_KHR_WIN32_SURFACE_EXTENSION_NAME`   | string       | `'VK_KHR_win32_surface'`              |

## Functions

### `initVulkan2DLayer(window)`

Creates a Vulkan instance with `VK_KHR_surface` and `VK_KHR_win32_surface` extensions, creates a `VkSurfaceKHR` attached to `window.hwnd`, enumerates physical devices, and selects the first device with a queue family that supports both graphics and surface presentation.

Returns `{ type: :ok, instance, surface, physicalDevice, queueFamily }` on success, or `{ type: :error, error: string, ... }` on failure.

```oak
result := initVulkan2DLayer(window)
if result.type = :ok {
    true -> {
        window.vulkanInstance <- result.instance
        window.vulkanSurface <- result.surface
        window.vulkanPhysicalDevice <- result.physicalDevice
        window.vulkanQueueFamily <- result.queueFamily
    }
    _ -> printf('Vulkan init failed: {{0}}', result.error)
}
```

**Note:** This module only performs instance- and surface-level initialization. Logical device creation, swap chain setup, and command buffer management are the caller's responsibility.
