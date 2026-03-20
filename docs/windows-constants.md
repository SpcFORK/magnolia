# Windows Constants Library (windows-constants)

## Overview

`windows-constants` stores DLL names and dynamic-link family prefixes used by
the split `windows-*` modules.

## Import

```oak
wc := import('windows-constants')
```

## Exports

- DLL constants such as `Kernel32`, `User32`, `Gdi32`, `Advapi32`, `Ws2_32`,
  `Wininet`, `OpenGL32`, `Vulkan1`, `D3d9`, and many others.
- Prefix constants:
  `ApiSetPrefix`, `D3dx9Prefix`, `MsvcpPrefix`, `VcruntimePrefix`,
  `AtlPrefix`, `MfcPrefix`, `VcompPrefix`.

## Notes

- `windows` re-exports behavior APIs; this module is primarily internal data.
