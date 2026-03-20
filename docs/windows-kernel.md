# Windows Kernel Library (windows-kernel)

## Overview

`windows-kernel` exposes process, module, memory, pointer, and error helpers.

## Import

```oak
wk := import('windows-kernel')
```

## Export groups

- String/encoding helpers: `_utf16leToString`, `wstr`, `cstr`
- Pointer helpers: `ptrSize`, `writePtr`, `ptrInt`
- Utility helpers: `statusOk?`, `callValueOrZero`, `_zeros`
- Error helpers: `getLastError`, `formatMessage`, `lastErrorMessage`
- Process/module helpers:
  `currentProcessId`, `currentProcess`, `moduleHandle`, `imageBase`,
  `loadLibrary`, `freeLibrary`, `procAddress`, `openProcess`, `closeHandle`
- Virtual memory helpers:
  `virtualAlloc`, `virtualAllocEx`, `virtualFree`, `virtualFreeEx`,
  `virtualProtect`, `virtualQuery`, `virtualQueryEx`
- Process memory helpers:
  `readProcessMemory`, `writeProcessMemory`
