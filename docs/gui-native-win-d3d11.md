# gui-native-win-d3d11 — Direct3D 11 Presenter

`import('gui-native-win-d3d11')` provides a Direct3D 11 2D presenter backend using `D3D11CreateDeviceAndSwapChain` with automatic WARP software fallback and pixel upload via `UpdateSubresource`.

## Overview

This module is an alternative presentation backend to GDI, DirectDraw, or OpenGL. It creates a D3D11 device and swap chain, uploads the GDI frame buffer pixels to a D3D11 texture, and presents to the screen. If hardware acceleration is unavailable, it falls back to the WARP software rasterizer.

## API Reference

### `initD3d112DLayer(window)`

Initializes the D3D11 rendering layer for a window. Creates the device, device context, swap chain, and back buffer. Returns an initialization result.

### `presentFrameViaD3d11(window)`

Uploads the current GDI frame buffer pixels to the D3D11 back buffer texture and presents to the screen.

### `releaseD3d11(window)`

Releases all D3D11 resources (device, context, swap chain, pixel buffer).

## Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `DXGI_FORMAT_B8G8R8A8_UNORM` | 87 | 32-bit BGRA color format |
| `DXGI_SWAP_EFFECT_DISCARD` | 0 | Swap chain discard mode |
| `D3D11_SDK_VERSION` | 7 | SDK version |
| `D3D_DRIVER_TYPE_HARDWARE` | 1 | Hardware driver |
| `D3D_DRIVER_TYPE_WARP` | 5 | WARP software fallback |
| `D3D11_CREATE_DEVICE_BGRA_SUPPORT` | 32 | BGRA device flag |

## Notes

- The module uses COM vtable indices directly for D3D11 interface calls.
- WARP fallback ensures the presenter works even on systems without D3D11-capable hardware.
- Pixel format is always BGRA (`DXGI_FORMAT_B8G8R8A8_UNORM`).
