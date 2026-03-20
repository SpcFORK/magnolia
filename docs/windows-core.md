# Windows Core Library (windows-core)

## Overview

`windows-core` provides the root platform checks and low-level resolve/call
helpers used by the full `windows` facade.

## Import

```oak
core := import('windows-core')
```

## Exports

- `_platformError(apiName)`
- `isWindows?()`
- `makeWord(low, high)`
- `resolve(symbol)`
- `resolveIn(library, symbol)`
- `call(target, args...)`
- `kernel32(symbol, args...)`
- `ntdll(symbol, args...)`
- `ntNative(symbol, args...)`
- `psapi(symbol, args...)`
