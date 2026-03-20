# Linux Loader Module (linux-loader)

## Overview

`linux-loader` provides dynamic library loading helpers (`dlopen`, `dlsym`,
`dlclose`), memoized handle caching, and convenience dispatch helpers for
arbitrary Linux shared libraries.

## Import

```oak
loader := import('linux-loader')
```

## API

### `_loadedLibraries`

Module-level cache map of `library -> handle` values.

### `_libraryKey(library)`

Creates a stable cache key for either:

- a single library name/path string, or
- a candidate list (uses the first element or `[]` for empty lists)

### `_normalizeHandleResult(result, apiName, library)`

Normalizes `sys.call` style results into `{type: :ok|:error, ...}` objects.

### `dlopen(path, flags)`
### `dlsym(handle, symbol)`
### `dlclose(handle)`

Thin wrappers over `libdl` entry points.

### `_loadDllCandidate(candidates, i, flags)`

Attempts `dlopen` for each candidate path/name in order.

### `loadDll(library)`

Loads a shared library and caches the handle.

- `library` can be a string or a candidate list.
- success shape: `{type: :ok, handle: <int>, library: <string>}`

### `resolveInLoaded(library, symbol)`

Ensures `library` is loaded, then resolves `symbol` with `dlsym`.

### `callIn(library, symbol, args...)`

Resolve + call helper.

- first tries cached loaded-handle dispatch
- if that fails, falls back to `resolveIn` / `_resolveFirst`

### `x11(symbol, args...)`

Convenience wrapper for dispatching through `LibX11` candidates.

## Example

```oak
loader := import('linux-loader')

res := loader.callIn(['libc.so.6', 'libc.so'], 'getpid')
println(string(res.r1))
```
