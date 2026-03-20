# Windows Loader Library (windows-loader)

## Overview

`windows-loader` handles DLL loading, symbol lookup, and dynamic family helpers
for `windows`.

## Import

```oak
loader := import('windows-loader')
```

## Core exports

- `loadDll(library)`
- `resolveInLoaded(library, symbol)`
- `callIn(library, symbol, args...)`

## Convenience wrappers

- Per-DLL wrappers such as `user32`, `gdi32`, `advapi32`, `shell32`, `ole32`,
  `ws2_32`, `wininet`, `d3d9`, `ddraw`, `vulkan1`, and many more.
- Family helpers:
  `d3dx9Dll`, `d3dx9`, `apiSetDll`, `apiSet`,
  `msvcpDll`, `msvcpFamily`, `vcruntimeDll`, `vcruntimeFamily`,
  `atlDll`, `atlFamily`, `mfcDll`, `mfcFamily`, `vcompDll`, `vcompFamily`.
- DirectDraw/Direct3D bootstrap helpers:
  `directDrawCreate`, `directDrawCreateEx`, `direct3DCreate9`.

## Notes

- This module is large by design and mostly used through `import('windows')`.
