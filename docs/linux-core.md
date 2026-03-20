# Linux Core Module (linux-core)

## Overview

`linux-core` contains platform guards, C-string helpers, and low-level symbol
resolution/dispatch for Linux shared libraries.

This module is the base layer used by Linux loader, libc, and X11 helpers.

## Import

```oak
core := import('linux-core')
```

## API

### `isLinux?()`

Returns `true` when the current host OS is Linux.

### `cstr(s)`

Returns `s` with a trailing null byte (`\0`) for C ABI calls.

### `_readCString(ptr, maxLen)`

Reads a null-terminated string from memory starting at `ptr`, up to `maxLen`
bytes.

### `_platformError(apiName)`

Returns a standardized Linux-only error object:

```oak
{type: :error, error: apiName + ' is only available on Linux'}
```

### `_resolveFirst(libraries, symbol, i)`

Recursively attempts `sys.resolve` over a candidate library list.

### `_callResolved(resolved, args...)`

Calls `sys.call(resolved.proc, args...)` when `resolved.type` is `:ok`.
Otherwise returns `resolved` unchanged.

### `resolve(symbol)`

Resolves `symbol` using the first working entry in `LibC`.

### `resolveIn(library, symbol)`

Resolves `symbol` from an explicit `library`.

### `call(target, args...)`

Direct pass-through to `sys.call` on Linux hosts.

### `libc(symbol, args...)`

Resolve + call against `LibC` candidates.

### `libdl(symbol, args...)`

Resolve + call against `LibDL` candidates.

## Notes

- All public calls are OS-gated.
- On non-Linux hosts, Linux-only calls return `{type: :error, ...}` objects.
