# System Interop Wrappers (sys)

## Overview

`sys` provides safer wrappers around Magnolia's low-level `sysproc` and
`syscall` built-ins.

Instead of handling multiple raw return shapes at every call site, `sys`
normalizes results into consistent objects and helper predicates.

## Import

```oak
sys := import('sys')
```

## Functions

### `ok?(result)`

Returns true when `result` is a successful syscall-style result.

Accepted success tags:

- `:ok`
- `:success`

```oak
if sys.ok?(res) -> println('call succeeded')
```

### `error?(result)`

Returns true when `result` is an error result (`type = :error`).

```oak
if sys.error?(res) -> println('call failed')
```

### `resolve(library, symbol)`

Resolves a procedure symbol from a shared library.

Success shape:

```oak
{
    type: :ok
    proc: <proc>
    library: <string>
    symbol: <string>
}
```

Error shape:

```oak
{
    type: :error
    error: <string>
    library: <string>
    symbol: <string>
    detail: <value>
}
```

```oak
pidProc := sys.resolve('kernel32.dll', 'GetCurrentProcessId')
if pidProc.type = :error -> println(pidProc.error)
```

### `call(target, args...)`

Calls a resolved procedure or address and normalizes malformed payloads.

Success shape:

```oak
{
    type: :ok
    r1: <int>
    r2: <int>
    errno?: <int>
}
```

Error shape:

```oak
{
    type: :error
    error: <string>
    detail?: <value>
}
```

```oak
resolved := sys.resolve('kernel32.dll', 'GetCurrentProcessId')
if resolved.type = :ok -> {
    res := sys.call(resolved.proc)
    if sys.ok?(res) -> println('PID: ' + string(res.r1))
}
```

### `resolveAndCall(library, symbol, args...)`

Convenience API that runs `resolve(...)` then `call(...)`.

```oak
res := sys.resolveAndCall('kernel32.dll', 'GetCurrentProcessId')
```

### `valueOr(result, fallback)`

Returns `result.r1` if successful; otherwise returns `fallback`.

```oak
pid := sys.valueOr(res, -1)
```

## Typical Pattern

```oak
sys := import('sys')

fn currentPID {
    res := sys.resolveAndCall('kernel32.dll', 'GetCurrentProcessId')
    sys.valueOr(res, -1)
}

println('PID: ' + string(currentPID()))
```

## Notes

- Keep direct `sysproc` and `syscall` usage inside wrapper code when possible.
- Validate argument types and pointers before native calls.
- For memory helpers and concurrency primitives, see the `go` documentation.

## Related Docs

- [Go Runtime and System Interop](./go.md)
- [GPU Library](./gpu.md)
- [Multi-Backend GPU Helpers](./gpus.md)
