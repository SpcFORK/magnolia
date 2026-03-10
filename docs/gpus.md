# Multi-Backend GPU Helpers (gpus)

## Overview

`gpus` is a thin middleware layer over `gpu`.

It solves one practical problem: your code can resolve and call GPU symbols
without branching on CUDA vs OpenCL in every call site.

The module:

- chooses a backend from a preferred order (default: `['cuda', 'opencl']`)
- resolves symbols against the chosen backend
- forwards calls to `gpu.call(...)`
- preserves structured `:ok` / `:error` result objects

## Import

```oak
gpus := import('gpus')
```

## Functions

### `select(order?)`

Chooses the first available backend from `order`.

- `order` may be a list of backend names or a single backend name string
- default order is `['cuda', 'opencl']`

Success shape:

```oak
{
    type: :ok
    backend: { ... }
    fallback: <bool>
    order: ['cuda', 'opencl']
}
```

Error shape:

```oak
{
    type: :error
    error: 'No GPU backend is available on this system'
    order: [...]
    scan: [...]
}
```

Example:

```oak
choice := gpus.select(['opencl', 'cuda'])
if choice.type {
    :ok -> println('Using backend: ' + choice.backend.name)
    :error -> println(choice.error)
}
```

### `resolve(symbol, order?)`

Resolves a symbol using the backend selected by `select(...)`.

Success shape:

```oak
{
    type: :ok
    backend: { ... }
    proc: <proc>
    symbol: <string>
}
```

Error shape:

```oak
{
    type: :error
    error: <string>
    backend?: { ... }
    symbol: <string>
    detail?: <value>
}
```

```oak
resolved := gpus.resolve('cuInit')
if resolved.type = :ok -> {
    println('Resolved on: ' + resolved.backend.name)
}
```

### `call(procOrAddress, args...)`

Pass-through to `gpu.call(...)`.

```oak
result := gpus.call(proc, 0)
```

### `invoke(symbol, args...)`

Convenience API: `resolve(symbol)` then `call(proc, args...)`.

```oak
result := gpus.invoke('cuInit', 0)
```

### `scan()` and `available()`

Pass-through helpers to `gpu.scan()` and `gpu.available()`.

```oak
allBackends := gpus.scan()
usableBackends := gpus.available()
```

## Typical Pattern

```oak
gpus := import('gpus')

fn initGPU {
    // Try OpenCL first, then CUDA
    chosen := gpus.select(['opencl', 'cuda'])
    if chosen.type = :error -> chosen

    // Resolve a backend-specific symbol in one place
    symbol := if chosen.backend.name {
        'cuda' -> 'cuInit'
        'opencl' -> 'clGetPlatformIDs'
    }
    gpus.resolve(symbol, [chosen.backend.name])
}
```

## Notes

- `gpus` does not hide backend API differences; it only centralizes backend
  selection and symbol resolution.
- For raw low-level interop details, see `gpu` and `go` docs.

## Related Docs

- [GPU Library](./gpu.md)
- [Go Runtime and System Interop](./go.md)
- [System Interop Wrappers](./sys.md)
